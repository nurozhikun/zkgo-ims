package ziclt

import (
	"gitee.com/sienectagv/gozk/znet"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zipbf"
)

type ZiHttpBee struct {
	*znet.HttpClient
	*zipbf.BeeHeader
}

func NewHttpBee() *ZiHttpBee {
	return &ZiHttpBee{HttpClient: znet.NewHttpClient(),
		BeeHeader: &zipbf.BeeHeader{}}
}

// func (zhb *ZiHttpBee) Get(cmd *ziapi.Command) (*http.Response, error) {
// 	r, _ := http.NewRequest("GET", path.Join(zhb.BasePath, cmd.Path), nil)
// 	zhb.BeeHeader.Cmd = int64(cmd.Cmd)
// 	zipbf.SetRequestBeeHeader(r, zhb.BeeHeader)
// 	return zhb.Do(r)
// }

// func (zhb *ZiHttpBee) Post(cmd *ziapi.Command) (*http.Response, error) {
// 	msg := cmd.FnRequestBody()
// 	r, _ := http.NewRequest("POST", path.Join(zhb.BasePath, cmd.Path), nil)
// }
