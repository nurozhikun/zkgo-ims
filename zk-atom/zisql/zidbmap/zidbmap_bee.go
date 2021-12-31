package zidbmap

import (
	"fmt"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zproto"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

//@reqBody is nil
//@return &protbee.MapThumbnails{}
func (db *DB) BeeMapThumbnails(header *zipbf.BeeHeader, reqBody zproto.Message) (zproto.Message, error) {
	body := &protbee.MapThumbnails{}
	//TODO read data from zidbmap then fill body
	zlogger.Info("test in BeeMapThumbnails")
	return body, nil
}

//@reqBody is zipbf.BeeThumbnail
func (db *DB) BeeMapOneDetail(header *zipbf.BeeHeader, reqBody zproto.Message) (zproto.Message, error) {
	thumbnail, ok := reqBody.(*protbee.Thumbnail)
	if !ok {
		return nil, zutils.NewError(-1, "the req body is nil")
	}
	fmt.Println(thumbnail)
	return nil, nil
}
