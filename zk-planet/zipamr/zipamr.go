package zipamr

import (
	"gitee.com/sienectagv/gozk/zdev"
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/zkgo-ims/zk-planet/zipamr/amrdb"
	"github.com/nurozhikun/zkgo-ims/zk-planet/zipamr/amrnet"
)

const (
	DeviceID_irisApp     = "iris_app"
	DeviceID_tcpListener = "tcp_listener"
)

type ZipAmr struct {
	// AddrTcpListener string
	zdev.Vessel
	App *iris.Application
	db  *amrdb.DB
	api *amrnet.AmrHttpApi
}

func (amr *ZipAmr) InitDB(cfg *zsql.Cfg) *ZipAmr {
	amr.db = amrdb.CreateDB(cfg)
	return amr
}

func (amr *ZipAmr) InitApi(cmds ...int) *ZipAmr {
	// irisApp := iris.New()
	if nil == amr.App {
		amr.App = iris.New()
	}
	amr.api = amrnet.NewHttpApi(amr.App, amr.db)
	amr.api.InstallPath(cmds...)
	return amr
}

func (amr *ZipAmr) StartDevices() *ZipAmr {
	amr.start()
	//append the deivce of iris application
	if nil != amr.App {
		node := &zdev.VirtualNode{
			ID:     DeviceID_irisApp,
			Addr:   ":9000",
			Stream: zdev.NewStreamIrisApp(amr.App),
		}
		zdev.CreateDevInVessel(node, &amr.Vessel)
	}
	//append the device of TcpListener
	node := &zdev.VirtualNode{
		ID:     DeviceID_tcpListener,
		Addr:   ":8000",
		Stream: zdev.NewStreamTcpListener(),
		Custom: zdev.NewCustomTcpListener(func() zdev.ICustom {
			return amrnet.NewCustomTcpConn()
		}),
	}
	zdev.CreateDevInVessel(node, &amr.Vessel)
	return amr
}

func (amr *ZipAmr) FreeAll() {
	amr.stop()
	if nil != amr.db {
		amr.db.DeleteDB()
	}
}

func (amr *ZipAmr) start() {
	// amr.Vessel.FnCreateCustom = func(code int64, cmd *zdev.Command) zdev.ICustom {
	// 	return nil
	// }
	// amr.Vessel.FnCreateStream = func(code int64, cmd *zdev.Command) zdev.IStream {
	// 	return nil
	// }
	amr.Vessel.Start()
}

func (amr *ZipAmr) stop() {
	amr.Vessel.Close()
}
