package netirisbee

import (
	"encoding/base64"
	"time"

	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zinet/netas"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
)

func ParserHeader(ctx znet.IrisCtx) (header *zipbf.BeeHeader, err error) {
	return
}

func SetHeader(ctx znet.IrisCtx, header *zipbf.BeeHeader) {
	if nil != header {
		header.Timestamp = time.Now().UTC().UnixMilli()
		if netas.UseBase64 {
			js := zproto.MarshalString(header)
			s := base64.StdEncoding.EncodeToString([]byte(js))
			ctx.Header(netas.HKeyHeader, s)
		}
		ctx.Header(netas.HKeyCmd, zutils.StringFromInterface(header.Cmd))
		ctx.Header(netas.HKeyTimestamp, zutils.StringFromInterface(header.Timestamp))
		ctx.Header(netas.HKeyJwt, header.Jwt)
	} else {
		CopyHeader(ctx)
	}
}

func CopyHeader(ctx znet.IrisCtx) {
	// ctx.Header(netas.HKeyCmd, ctx.GetHeader(netas.HKeyCmd))
	znet.IrisCopyHeaderKeys(ctx, netas.HKeyCmd, netas.HKeyJwt)
	t := time.Now().UTC().UnixMilli()
	ctx.Header(netas.HKeyTimestamp, zutils.StringFromInterface(t))

}

func ResponseError(ctx znet.IrisCtx, err error, code int) {
	if err != nil {
		if n := zutils.ErrorCode(err); 0 != n {
			code = n
		}
	}
	ctx.Header(netas.HKeyError, err.Error())
	szcode, _ := zutils.InterfaceToString(code)
	ctx.Header(netas.HKeyCode, szcode)
	if netas.UseBase64 {
		tailer := &zipbf.BeeTailer{
			Code:  int64(code),
			Error: err.Error(),
		}
		js := zproto.MarshalString(tailer)
		ctx.Header(netas.HKeyTailer, base64.StdEncoding.EncodeToString([]byte(js)))
	}
}
