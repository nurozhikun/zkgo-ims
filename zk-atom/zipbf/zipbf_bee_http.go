package zipbf

import (
	"encoding/base64"
	"fmt"
	"time"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zinet/netas"
)

func BeeHttpPost(c *znet.HttpClient, h *BeeHeader, cmd *Command) (*znet.Response, error) {
	if nil == c || nil == h {
		return nil, fmt.Errorf("Reqeust or BeeHeader is null")
	}
	r := c.R()
	SetBeeHeader(r, h)
	if nil != cmd.FnRequestBody {
		rqbody := BeeReqBodyFromCmd(cmd)
		if nil != rqbody {
			r.SetBody(zproto.MarshalString(rqbody))
		}
	}
	url := c.BaseURL + cmd.Path
	zlogger.Info("Post", url, cmd.Cmd)
	return r.Post(url)
}

func SetBeeHeader(r *znet.Request, h *BeeHeader) {
	h.Timestamp = time.Now().UTC().Unix()
	if netas.UseBase64 {
		js := zproto.MarshalString(h)
		s := base64.StdEncoding.EncodeToString(([]byte(js)))
		r.SetHeader(netas.HKeyHeader, s)
	}
	r.SetHeader(netas.HKeyCmd, zutils.StringFromInterface(h.Cmd)).
		SetHeader(netas.HKeyTimestamp, zutils.StringFromInterface(h.Timestamp)).
		SetHeader(netas.HKeyJwt, h.Jwt)
}
