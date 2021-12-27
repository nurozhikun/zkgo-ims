package zicustom

import (
	"bytes"
	"encoding/binary"

	"gitee.com/sienectagv/gozk/zdev"
	"gitee.com/sienectagv/gozk/zutils"
)

const (
	FieldEggCmd = "egg_cmd"
)

var (
	DefHeadMarker = []byte{0xF0, 0xF1, 0xF2, 0xF3}
	DefTailMarker = []byte{0xFF, 0xFE, 0xFD, 0xFC}
)

type CustProtBinaryEgg struct {
	zdev.CustomBase
	HeadMarker   []byte //everyone is different
	TailMarker   []byte
	CurHead      EggHead
	FnDecodeBody func(body []byte, cmd *zdev.Command) error
	minPackSize  int32
	maxPackSize  int32
	buffers      []byte
}

type EggHead struct {
	Cmd    uint32
	PkgTye uint16
	Len    uint16
}

const headSize = 8

type EggTail struct {
	Check int32
}

const tailsize = 4

func (p *CustProtBinaryEgg) IUnpackToCommand(bin interface{}) (cmd *zdev.Command, unfinished bool, err error) {
	buff, ok := bin.([]byte)
	if !ok {
		return nil, false, zutils.NewError(-1, "it's not a slice of []byte")
	}
	cmd, _, err = p.CustomBase.IUnpackToCommand(bin)
	body, unfinished, err := p.getOnePackage(buff)
	if nil == body || nil != err {
		return nil, unfinished, nil
	}
	err = p.decodeBody(body, cmd)
	return
}

// func (p *CustProtBinaryEgg) SetHeadMarker(headMarker []byte) {
// 	p.HeadMarker = headMarker
// }

// func (p *CustProtBinaryEgg) SetTailMarker(tailMarker []byte) {
// 	p.TailMarker = tailMarker
// }

func (p *CustProtBinaryEgg) getOnePackage(in []byte) (body []byte, unfinished bool, err error) {
	var buff []byte
	if len(in) > 0 {
		if len(p.buffers) == 0 {
			buff = in
		} else {
			p.buffers = append(p.buffers, in...)
			buff = p.buffers
		}
	} else {
		buff = p.buffers
	}
	if len(buff) < int(p.minPackSize) {
		return nil, false, nil
	}
	i := bytes.Index(buff, p.HeadMarker)
	if -1 == i {
		p.buffers = nil
		return nil, false, zutils.NewError(-1, "can't find the HeadMarker")
	}
	if i > 0 { //skip the unknown bytes
		buff = buff[i:]
		i = 0
	}
	if len(buff) < int(p.minPackSize) {
		p.cachedBuff(buff)
		return nil, false, nil
	}
	err = binary.Read(bytes.NewReader(buff[len(p.HeadMarker):]), binary.BigEndian, &p.CurHead)
	if nil != err {
		p.cachedBuff(buff)
		return nil, false, err
	}
	i = len(p.HeadMarker) + headSize + int(p.CurHead.Len) + tailsize + len(p.TailMarker)
	if i > len(buff) {
		p.cachedBuff(buff[len(p.HeadMarker):]) //skip the HeadMarker
		return nil, true, zutils.NewError(-1, "can't find the TailMarker for len")
	}
	if bytes.Compare(p.TailMarker, buff[i-len(p.TailMarker):i]) != 0 {
		p.cachedBuff(buff[len(p.HeadMarker):]) //skip the HeadMarker
		return nil, true, zutils.NewError(-1, "can't find the TailMarker for data")
	}
	p.cachedBuff(buff[i:])
	i = len(p.HeadMarker) + headSize
	return buff[i : i+int(p.CurHead.Len)], true, nil
}

func (p *CustProtBinaryEgg) cachedBuff(buff []byte) {
	if len(buff) == 0 {
		p.buffers = nil
	} else if len(p.buffers) == 0 { //make a copy
		p.buffers = make([]byte, len(buff))
		copy(p.buffers, buff)
	} else {
		p.buffers = buff
	}
}

func (p *CustProtBinaryEgg) decodeBody(body []byte, cmd *zdev.Command) error {
	return nil
}
