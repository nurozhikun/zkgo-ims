package zipbf

import (
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf/protbee"
)

type (
	BeeHeader = protbee.Header
	// BeeTailer = protbee.Tailer
	// BeeThunmbnail = protbee.Thumbnail
)

func SetBeeHeaderError(h *BeeHeader, err error) {
	h.Code = int64(zutils.ErrorCode(err))
	h.Error = err.Error()
}
