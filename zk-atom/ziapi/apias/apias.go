package apias

const (
	//for apimaps
	Cmd_MapThumbnails = 1
	Cmd_MapOneDetail  = 2
)

var CmdPaths = map[int]string{Cmd_MapThumbnails: "/thumbnails"}
