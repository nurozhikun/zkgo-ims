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
	header = &zipbf.BeeHeader{}
	s := ctx.GetHeader(netas.HKeyHeader)
	if len(s) > 0 {
		zproto.UnmarshalString([]byte(s), header)
	}
	header.Cmd, _ = zutils.InterfaceToInt(ctx.GetHeader(netas.HKeyCmd))
	header.Timestamp, _ = zutils.InterfaceToInt(ctx.GetHeader(netas.HKeyTimestamp))
	header.Jwt = ctx.GetHeader(netas.HKeyJwt)
	header.Code, _ = zutils.InterfaceToInt(ctx.GetHeader(netas.HKeyCode))
	header.Error = ctx.GetHeader(netas.HKeyError)
	return
}

func SetHeader(ctx znet.IrisCtx, header *zipbf.BeeHeader, err error) {
	if nil != header {
		if nil != err {
			zipbf.SetBeeHeaderError(header, err)
		}
		header.Timestamp = time.Now().UTC().UnixMilli()
		if netas.UseBase64 {
			js := zproto.MarshalString(header)
			s := base64.StdEncoding.EncodeToString([]byte(js))
			ctx.Header(netas.HKeyHeader, s)
		}
		ctx.Header(netas.HKeyCmd, zutils.StringFromInterface(header.Cmd))
		ctx.Header(netas.HKeyTimestamp, zutils.StringFromInterface(header.Timestamp))
		ctx.Header(netas.HKeyJwt, header.Jwt)
		ctx.Header(netas.HKeyCode, zutils.StringFromInterface(header.Code))
		ctx.Header(netas.HKeyError, header.Error)
	} else {
		CopyHeader(ctx)
		if nil != err {
			ctx.Header(netas.HKeyCode, zutils.StringFromInterface(zutils.ErrorCode(err)))
			ctx.Header(netas.HKeyError, err.Error())
		}
	}
}

func CopyHeader(ctx znet.IrisCtx) {
	znet.IrisCopyHeaderKeys(ctx, netas.HKeyCmd, netas.HKeyJwt)
	t := time.Now().UTC().UnixMilli()
	ctx.Header(netas.HKeyTimestamp, zutils.StringFromInterface(t))
}

func ParseBody(ctx znet.IrisCtx, msg zproto.Message) error {
	bs, err := ctx.GetBody()
	if nil != err {
		return err
	}
	return zproto.UnmarshalString(bs, msg)
}
