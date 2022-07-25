package itface

type IConnMgr interface {
	// 通过过滤条件cond获取conn
	GetConn(cond map[string]interface{}) (IConnection, error)
	//添加conn到connmgr中
	Add(conn IConnection) error
	// 将conn从connmgr中移除
	Remove(conn IConnection) error
	// 返回connmgr中conn数量
	Count() int
	//清除所有conn
	ClearConn()
	// 通过过滤条件Cond随机选择一个conn
	RandomSelectConnWithCond(cond map[string]interface{}) (IConnection, error)
	// 获取所有满足条件的conns
	GetMeetCondConns(cond map[string]interface{}) ([]IConnection, error)

	SetMsgHandler(func(IRequest))
	GetMsgHandler() func(IRequest)
}
