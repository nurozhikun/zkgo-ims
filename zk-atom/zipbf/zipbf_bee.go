package zipbf

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zinet/netas"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

type (
	BeeHeader = protbee.Header
	Command   = znet.Command
	// BeeTailer = protbee.Tailer
	// BeeThunmbnail = protbee.Thumbnail
)

func BeeReqBodyFromCmd(cmd *Command) zproto.Message {
	if nil == cmd.FnRequestBody {
		return nil
	}
	r := cmd.FnRequestBody()
	if nil == r {
		return nil
	}
	req, _ := r.(zproto.Message)
	return req
}

func SetBeeHeaderError(h *BeeHeader, err error) {
	h.Code = int64(zutils.ErrorCode(err))
	h.Error = err.Error()
}

func SetRequestBeeHeader(r *http.Request, h *BeeHeader) {
	if nil == r || nil == h {
		return
	}
	h.Timestamp = time.Now().UTC().Unix()
	if netas.UseBase64 {
		js := zproto.MarshalString(h)
		s := base64.StdEncoding.EncodeToString(([]byte(js)))
		r.Header.Add(netas.HKeyHeader, s)
	}
	r.Header.Add(netas.HKeyCmd, zutils.StringFromInterface(h.Cmd))
	r.Header.Add(netas.HKeyTimestamp, zutils.StringFromInterface(h.Timestamp))
	r.Header.Add(netas.HKeyJwt, h.Jwt)
	// r.Header.Add(netas.HKeyCode, zutils.StringFromInterface(h.Code))
	// r.Header.Add(netas.HKeyError, h.Error)
}

func SetResponseBee(r *http.Request, h *BeeHeader, msg zproto.Message) error {
	if nil == r || nil == h {
		return fmt.Errorf("Reqeust or BeeHeader is null")
	}
	SetRequestBeeHeader(r, h)
	// if nil != msg {
	// 	r.Body = bytes.NewBufferString(zproto.MarshalString(msg))
	// }
	return nil
}

func ParseResponseBee(r *http.Response, h *BeeHeader, msg zproto.Message) error {
	if nil != h {
		s := r.Header.Get(netas.HKeyHeader)
		if len(s) > 0 {
			if netas.UseBase64 {
				bs, _ := base64.StdEncoding.DecodeString(s)
				zproto.UnmarshalString(bs, h)
			} else {
				zproto.UnmarshalString([]byte(s), h)
			}
		}
		s = r.Header.Get(netas.HKeyCmd)
		if len(s) > 0 {
			h.Cmd, _ = zutils.InterfaceToInt(s)
		}
		s = r.Header.Get(netas.HKeyTimestamp)
		if len(s) > 0 {
			h.Timestamp, _ = zutils.InterfaceToInt(s)
		}
		s = r.Header.Get(netas.HKeyJwt)
		if len(s) > 0 {
			h.Jwt = s
		}
		s = r.Header.Get(netas.HKeyCode)
		if len(s) > 0 {
			h.Code, _ = zutils.InterfaceToInt(s)
		}
		h.Error = r.Header.Get(netas.HKeyError)
	}
	var err error
	if nil != msg {
		bs := make([]byte, r.ContentLength)
		r.Body.Read(bs)
		err = zproto.UnmarshalString(bs, msg)
	}
	return err
}
