package ziapi

import (
	"gitee.com/sienectagv/gozk/znet"
	"gitee.com/sienectagv/gozk/zsql"
	"github.com/nurozhikun/zkgo-ims/zk-atom/ziapi/apiauth"
	"github.com/nurozhikun/zkgo-ims/zk-atom/zisql"
)

func (iba *IrisBeeApis) InstallBeeAdminApi(party znet.IrisParty, db *zsql.DB, cmds ...Command) error {
	iba.InstallApi(&apiauth.ApiAuth{
		DB: zisql.WrapupDbAuth(db),
	})
	if nil != cmds {
		cmds = AuthCmds
	}
	iba.InstallBeeHandlerOfCmds(party, cmds...)
	return nil
}
