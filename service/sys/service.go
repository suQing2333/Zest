package sys

import (
	"fmt"
	"time"
	"zest/engine/conf"
	"zest/engine/funcmgr"
	"zest/engine/itface"
	"zest/engine/netutil"
	"zest/engine/pbmgr"
	"zest/engine/zslog"
	// "zest/service/define"
)

type Service struct {
	sID        int32                  // 服务的id
	sName      string                 // 服务名
	svrInfo    map[string]interface{} // 服务的信息
	OutServer  *netutil.TCPServer     // TCPServer
	er         *etcdRegister
	ed         *etcdDiscover
	ConnMgr    itface.IConnMgr // 用来管理与其他服务的连接
	MsgHandler func(req itface.IRequest)
}

// 单例,一个进程只能有一个service
func NewService(sid int32, sname string) *Service {
	if Svr == nil {
		mu.Lock()
		defer mu.Unlock()
		Svr = &Service{
			sID:     sid,
			sName:   sname,
			svrInfo: GetSvrInfo(sname, sid),
			ConnMgr: NewConnMgr(),
		}
		funcmgr.RegisterFunc(*Svr, Svr)
	}
	return Svr
}

func GetService() *Service {
	return Svr
}

func GetSvrInfo(sname string, sid int32) map[string]interface{} {
	svrInfo := make(map[string]interface{})
	addr := fmt.Sprintf("%v:%v", netutil.GetPulicIP(), sid)
	svrInfo["sname"] = sname
	svrInfo["sid"] = sid
	svrInfo["info"] = addr
	return svrInfo
}

func (s *Service) GetSID() int32 {
	return s.sID
}

func (s *Service) GetSName() string {
	return s.sName
}

func (s *Service) GetSvrInfo() map[string]interface{} {
	return s.svrInfo
}

func (s *Service) GetOutServer() *netutil.TCPServer {
	return s.OutServer
}

func (s *Service) Start() {
	s.SetServiceMshHanler(processRequest)

	isOut := conf.GetBool("Service.isOut")
	s.OutServer = netutil.NewTCPServer(s.sName, int(s.sID), isOut)
	go s.OutServer.Serve()
	s.SetOutServerMsgHandler(processRequest)

	s.ConnMgr.SetMsgHandler(processRequest)

	s.er = newEtcdRegister(s.svrInfo)
	s.er.Start()
	s.ed = newEtcdDiscover()
	s.ed.Start()

}

func (s *Service) SetOutServerMsgHandler(hookFunc func(itface.IRequest)) {
	s.OutServer.SetMsgHandler(hookFunc)
}

func (s *Service) SetServiceMshHanler(hookFunc func(itface.IRequest)) {
	s.MsgHandler = hookFunc
}

func (s *Service) AddOrUpdateConn(sname string, sid int, data []byte) bool {
	isSuccess := false
	// 过滤掉自己的注册
	if s.sName == sname && s.sID == int32(sid) {
		zslog.LogInfo("addOrUpdateConn CreatConn filter Self register")
		return isSuccess
	}
	addr := string(data)
	conn, err := netutil.CreatConn(addr)
	if err != nil {
		zslog.LogError("addOrUpdateConn CreatConn error :%v", err)
		return isSuccess
	}

	// 增加三秒钟的keepalive周期
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(3 * time.Second)
	// 首次添加必定失败,需要再添加一次
	sConn := netutil.NewConnection(conn, s.ConnMgr)
	extraInfo := map[string]interface{}{"sname": sname, "sid": sid}
	sConn.SetExtraInfo(extraInfo)
	s.ConnMgr.Add(sConn)
	go sConn.Start()
	isSuccess = true
	zslog.LogDebug("AddOrUpdateConn sname %v,sid %v", sname, sid)
	return isSuccess
}

func (s *Service) DeleteConn(sname string, sid int) {
	cond := map[string]interface{}{"sname": sname, "sid": sid}
	conn, err := s.ConnMgr.GetConn(cond)
	if err != nil {
		zslog.LogWarn("deleteConn error : %v", err)
		return
	}
	s.ConnMgr.Remove(conn)
}

func (s *Service) SendProtoToServiceWithSID(msgType int, module string, sid int32, Cmd int32, subCmd int32, protoData []byte) {
	data := pbmgr.GetDataByFullName("protoc.BaseProto", Cmd, subCmd, protoData, []byte{})
	s.SendBaseProtoToServiceWithSID(msgType, module, sid, data)
}

func (s *Service) SendBaseProtoToServiceWithSID(msgType int, module string, sid int32, protoData []byte) {
	msg := netutil.NewMsg(uint32(msgType), protoData)
	s.SendMsgToServiceWithSID(module, sid, msg)
}

// 发送消息到指定Service
func (s *Service) SendMsgToServiceWithSID(module string, sid int32, msg itface.IMessage) {
	cond := map[string]interface{}{"sname": module, "sid": int(sid)}
	conn, err := s.ConnMgr.GetConn(cond)
	if err != nil {
		zslog.LogError("SendMsgToServiceWithSID err : %v", err)
		return
	}
	s.SendMsgWithConn(conn, msg)
}

// 发送协议到Service
func (s *Service) SendProtoToService(msgType int, module string, Cmd int32, subCmd int32, protoData []byte) {
	data := pbmgr.GetDataByFullName("protoc.BaseProto", Cmd, subCmd, protoData, []byte{})
	s.SendBaseProtoToService(msgType, module, data)
}

func (s *Service) SendBaseProtoToService(msgType int, module string, protoData []byte) {
	msg := netutil.NewMsg(uint32(msgType), protoData)
	s.SendMsgToService(module, msg)
}

// 发送消息到Service
func (s *Service) SendMsgToService(module string, msg itface.IMessage) {
	var cond = map[string]interface{}{"sname": module}
	conn, err := s.ConnMgr.RandomSelectConnWithCond(cond)
	if err != nil {
		zslog.LogError("SendMsgToService err : %v", err)
		return
	}
	s.SendMsgWithConn(conn, msg)
}

// 通过conn发送协议
func (s *Service) SendProtoWithConn(conn itface.IConnection, msgType int, Cmd int32, subCmd int32, protoData []byte) {
	data := pbmgr.GetDataByFullName("protoc.BaseProto", Cmd, subCmd, protoData, []byte{})
	s.SendBaseProtoWithConn(conn, msgType, data)
}

func (s *Service) SendBaseProtoWithConn(conn itface.IConnection, msgType int, protoData []byte) {
	msg := netutil.NewMsg(uint32(msgType), protoData)
	s.SendMsgWithConn(conn, msg)
}

func (s *Service) SendMsgWithConn(conn itface.IConnection, msg itface.IMessage) {
	err := conn.SendBuffMsg(msg)
	if err != nil {
		zslog.LogError("SendMsgWithConn err : %v", err)
		return
	}
}
