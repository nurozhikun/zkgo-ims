package apias

import "gitee.com/sienectagv/gozk/zproto"

type Command struct {
	Cmd        int
	Path       string
	MethodName string
	// 需要实现这个函数去返回具体的类型
	FnBeeReqBody func() zproto.Message
}

func (c *Command) BeeReqBody() zproto.Message {
	if nil == c.FnBeeReqBody {
		return nil
	}
	return c.FnBeeReqBody()
}
