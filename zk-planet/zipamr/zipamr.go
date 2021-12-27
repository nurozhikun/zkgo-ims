package zipamr

import (
	"gitee.com/sienectagv/gozk/zdev"
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-planet/zipamr/amrdb"
)

type ZipAmr struct {
	AddrTcpListener string
	zdev.Vessel
	*amrdb.DB
}

func (amr *ZipAmr) InitDB(cfg *zsql.Cfg) *ZipAmr {
	amr.DB = amrdb.CreateDB(cfg)
	return amr
}

func (amr *ZipAmr) InitDevs() *ZipAmr {
	amr.start()
	//append the device of TcpListener
	return amr
}

func (amr *ZipAmr) FreeAll() {
	amr.stop()
	amr.DB.DeleteDB()
}

func (amr *ZipAmr) start() {
	amr.Vessel.FnCreateCustom = func(code int64, cmd *zdev.Command) zdev.ICustom {
		return nil
	}
	amr.Vessel.FnCreateStream = func(code int64, cmd *zdev.Command) zdev.IStream {
		return nil
	}
	amr.Vessel.Start()
}

func (amr *ZipAmr) stop() {
	amr.Vessel.Close()
}
