package define

// 对于subcmd的定义，约定小于1000属于公共subcmd，服务自定义的subcmd必须从1001开始
const (
	RPCRequestCMD = iota
	RPCResponseCMD
)
