package futugg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"futugg/event"
	"futugg/pb/InitConnect"
	"github.com/golang/protobuf/proto"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type FutuGG struct {
	Host              string
	Port              string
	Pass              string
	Conn              net.Conn
	SyncMode          bool
	ServerVer         int32
	ConnID            uint64
	LoginUserID       uint64
	ConnAESKey        string
	KeepAliveInterval int32
}

func New(host string, port string, pass string) *FutuGG {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Cannot connect server [%s:%s]", host, port))
	}

	gg := &FutuGG{
		Host: host,
		Port: port,
		Pass: pass,
		Conn: conn,
	}
	gg.initConnect()
	return gg
}

func (c *FutuGG) Send(pack *FutuPack) error {
	packData, err := pack.Encode()
	if err != nil {
		return err
	}

	// set timeout
	c.Conn.SetWriteDeadline(time.Now().Add(3 * time.Second))

	// write
	_, err = c.Conn.Write(packData)
	return err
}

func (c *FutuGG) Recv() []byte {
	// scanner
	scanner := bufio.NewScanner(c.Conn)
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*10)

	// split
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'F' {
			if len(data) >= 44 {
				length := uint32(0)
				binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &length)

				if int(length)+44 <= len(data) {
					return int(length) + 44, data[:int(length)+44], nil
				}

			}
		}

		return
	})

	// scan
	for scanner.Scan() {
		// read timeout
		c.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		pack := new(FutuPack)
		err := pack.Decode(scanner.Bytes())
		if err != nil {
			log.Fatalln("unpack error", err)
		}

		protoId := uint32(pack.ProtoId)

		if c.SyncMode {
			return pack.Body
		}

		handlerName, ok := TransHandlerId(protoId)

		if ok {
			recvFunc := "recv." + handlerName
			if !HasHandler(recvFunc) {
				fmt.Println("recvFunc error: ", protoId, recvFunc)
			}
			// Recv
			Cmd(recvFunc, pack.Body)
		} else {
			fmt.Println("trans handler " + strconv.Itoa(int(protoId)) + " error")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("invalid apckage ", err)
	}
	return nil
}

func (c *FutuGG) KeepAlive() {
	go func() {
		tick := time.NewTicker(3 * time.Second)

		for {
			select {
			case <-tick.C:
				fmt.Println("updated")
				Cmd("send.KeepAlive", c)
			}

		}
	}()
}

func (c *FutuGG) Close() {
	c.Conn.Close()
	c.ConnID = 0
}

func (c *FutuGG) initConnect() {
	rand.Seed(int64(time.Now().Nanosecond()))
	clientId := strconv.Itoa(int(rand.Int31()))
	clientVer := int32(100)

	pack := &FutuPack{}
	pack.SetProto(uint32(1001))

	reqData := &InitConnect.Request{
		C2S: &InitConnect.C2S{
			ClientID:  &clientId,
			ClientVer: &clientVer,
		},
	}

	pbData, err := proto.Marshal(reqData)
	if err != nil {
		log.Fatalln("marsha1 error: ", err)
	}

	pack.SetBody(pbData)
	c.Send(pack)

	c.SyncMode = true
	rawPack := c.Recv()

	retData := initConnectData(rawPack)
	c.ServerVer = *retData.S2C.ServerVer
	c.LoginUserID = *retData.S2C.LoginUserID
	c.ConnID = *retData.S2C.ConnID
	c.ConnAESKey = *retData.S2C.ConnAESKey
	c.KeepAliveInterval = *retData.S2C.KeepAliveInterval

	c.SyncMode = false
	return
}

func initConnectData(data []byte) *InitConnect.Response {
	resp := &InitConnect.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		log.Fatal("InitConnect unmarshal error: ", err)
	}
	if resp.GetRetType() != int32(0) {
		log.Fatalln("InitConnectData rettype: ", resp.GetRetType())
	}
	return resp
}

var handlerMap = make(map[uint32]string)

var handlers = event.New()

func On(name string, fn interface{}) error {
	return handlers.On(name, fn)
}

func Cmd(name string, params ...interface{}) error {
	if !HasHandler(name) {
		fmt.Println(name, " handler func is not exist")
	}
	return handlers.Cmd(name, params...)
}

func HasHandler(name string) bool {
	return handlers.Has(name)
}

func SetHandlerId(id uint32, name string) {
	handlerMap[id] = name
}

func TransHandlerId(id uint32) (string, bool) {
	str, ok := handlerMap[id]
	return str, ok
}
