package futugg

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math/rand"
)

type FutuPack struct {
	HeaderFlag    [2]uint8  // [0]"F" [1]"T"
	ProtoId       uint32    // proto id
	ProtoType     uint8     // proto 类型，0 proto，1 json
	ProtoVer      uint8     // proto version 用于迭代 默认0
	SeqNo         uint32    // seq number 用于回包
	BodyLen       uint32    // body length
	BodySha1      [20]uint8 // body sha1
	ReservedField [8]uint8  // 预留8位
	Body          []byte    // body content
}

func (p *FutuPack) SetProto(protoId uint32) {
	// set proto id
	p.ProtoId = protoId
}

func (p *FutuPack) SetBody(c2sBody []byte) {
	p.Body = c2sBody
	p.BodyLen = uint32(len(c2sBody))

	sha := sha1.New()
	sha.Write(p.Body)
	bodySha1 := sha.Sum(nil)
	copy(p.BodySha1[:], bodySha1)
}

func (p *FutuPack) Encode() ([]byte, error) {
	var err error

	// set header
	p.HeaderFlag[0] = uint8('F')
	p.HeaderFlag[1] = uint8('T')

	// proto type default pb
	p.ProtoType = uint8(0)

	p.ProtoVer = 0

	p.SeqNo = uint32(rand.Int31())

	packBuf := new(bytes.Buffer)
	err = binary.Write(packBuf, binary.LittleEndian, &p.HeaderFlag)
	err = binary.Write(packBuf, binary.LittleEndian, &p.ProtoId)
	err = binary.Write(packBuf, binary.LittleEndian, &p.ProtoType)
	err = binary.Write(packBuf, binary.LittleEndian, &p.ProtoVer)
	err = binary.Write(packBuf, binary.LittleEndian, &p.SeqNo)
	err = binary.Write(packBuf, binary.LittleEndian, &p.BodyLen)
	err = binary.Write(packBuf, binary.LittleEndian, &p.BodySha1)
	err = binary.Write(packBuf, binary.LittleEndian, &p.ReservedField)

	err = binary.Write(packBuf, binary.LittleEndian, &p.Body)

	return packBuf.Bytes(), err
}

func (p *FutuPack) Decode(recv []byte) error {
	var err error

	reader := bytes.NewReader(recv)
	err = binary.Read(reader, binary.LittleEndian, &p.HeaderFlag)
	err = binary.Read(reader, binary.LittleEndian, &p.ProtoId)
	err = binary.Read(reader, binary.LittleEndian, &p.ProtoType)
	err = binary.Read(reader, binary.LittleEndian, &p.ProtoVer)
	err = binary.Read(reader, binary.LittleEndian, &p.SeqNo)
	err = binary.Read(reader, binary.LittleEndian, &p.BodyLen)
	err = binary.Read(reader, binary.LittleEndian, &p.BodySha1)
	err = binary.Read(reader, binary.LittleEndian, &p.ReservedField)

	p.Body = make([]byte, p.BodyLen)
	err = binary.Read(reader, binary.LittleEndian, &p.Body)

	return err
}

func (p *FutuPack) ToString() string {
	return fmt.Sprintf("Bodylen: %d body: %s", p.BodyLen, p.Body)
}
