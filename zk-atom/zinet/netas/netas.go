package netas

import "gitee.com/sienectagv/gozk/znet"

const (
	HKeyCmd       = znet.ZkCmd
	HKeyTimestamp = znet.ZkTimestamp
	HKeyJwt       = znet.ZkJwt
	HKeyCode      = znet.ZkCode
	HKeyError     = znet.ZkError
	HKeyHeader    = znet.ZkHeader
	// HKeyTailer    = znet.ZkTailer
)

const (
	ResCode_Success      = 0
	ResCode_UnknownError = -1
)

var (
	UseBase64 = false
)
