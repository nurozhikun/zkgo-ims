package amrnet

import (
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apias"
	"github.com/nurozhikun/zkgo-ims/zk-planet/zipamr/amrdb"
)

func NewHttpApi(app *iris.Application, db *amrdb.DB) *AmrHttpApi {
	a := &AmrHttpApi{
		app:  app,
		db:   db,
		apis: &ziapi.IrisBeeApis{},
		// apiMap: ziapi.NewApiMap(db.DB),
	}
	a.apis.InstallApi(ziapi.NewApiMap(db.DB))
	return a
}

type AmrHttpApi struct {
	app *iris.Application
	db  *amrdb.DB
	// apiMap *ziapi.ApiMap
	apis *ziapi.IrisBeeApis
}

func (a *AmrHttpApi) InstallPath(cmds ...int) {
	//install map apis
	party := a.app.Party("/map")
	{
		a.apis.InstallBeeHandleByCmds(party,
			apias.Cmd_MapThumbnails,
			apias.Cmd_MapOneDetail)
		// ziapi.InstallIrisBeeHandles(a.apiMap, party,
		// 	apias.Cmd_MapThumbnails,
		// 	apias.Cmd_MapOneDetail)
	}
}
