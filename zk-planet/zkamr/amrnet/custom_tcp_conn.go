package amrnet

import "github.com/nurozhikun/zkgo-ims/zk-atom/zicustom"

func NewCustomTcpConn() *CustomTcpConn {
	c := &CustomTcpConn{
		CustProtBinaryEgg: zicustom.NewCustProtBinaryEgg(),
	}
	return c
}

type CustomTcpConn struct {
	*zicustom.CustProtBinaryEgg
}
