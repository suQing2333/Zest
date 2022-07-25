package sys

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"strings"
	"sync"
	"time"
	"zest/engine/conf"
	"zest/engine/funcmgr"
	"zest/engine/zslog"
)

const (
	Prefix = "service"
)

type etcdDiscover struct {
	ctx        context.Context
	cancel     context.CancelFunc
	etcdClient *clientv3.Client
	ServiceMap map[string]map[int][]byte
	lock       sync.Mutex
}

func newEtcdDiscover() *etcdDiscover {
	endpoints := conf.GetStringSlice("etcd.endpoints")
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: 5 * time.Second})
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	ed := &etcdDiscover{
		ctx:        ctx,
		cancel:     cancel,
		etcdClient: cli,
		ServiceMap: make(map[string]map[int][]byte),
	}
	return ed
}

func (ed *etcdDiscover) Start() {
	go ed.discover()
}

func (ed *etcdDiscover) discover() {
	ctx, cancel := context.WithCancel(ed.ctx)
	defer cancel()

	err := ed.loadService(ctx)
	if err != nil {
		zslog.LogError("etcd discover loadService err :%v", err)
	}

	watch := ed.etcdClient.Watch(ctx, Prefix, clientv3.WithPrefix())
	for {
		select {
		case <-ed.ctx.Done():
			zslog.LogError("etcd discover end ")
			return
		case resp := <-watch:
			if err := resp.Err(); err != nil {
				zslog.LogError("discover watch err :%v", err)
				return
			}
			for _, event := range resp.Events {
				zslog.LogDebug("discover new register :%v", event)
				if event.Kv == nil {
					continue
				}
				switch event.Type {
				case mvccpb.PUT:
					ed.addOrUpdate(event.Kv.Key, event.Kv.Value)
				case mvccpb.DELETE:
					ed.delete(event.Kv.Key)
				}
			}
		}
	}
}

func (ed *etcdDiscover) loadService(ctx context.Context) error {
	resp, err := ed.etcdClient.Get(ctx, Prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, kv := range resp.Kvs {
		ed.addOrUpdate(kv.Key, kv.Value)
	}
	return nil
}

func (ed *etcdDiscover) addOrUpdate(key []byte, value []byte) {
	zslog.LogDebug("etcdDiscover addOrUpdate key:%v,value %v", key, value)
	ed.lock.Lock()
	defer ed.lock.Unlock()

	sname, sid, err := tailKey(key)
	if err != nil {
		zslog.LogError("addOrUpdate tailKey err,key :%v", key)
		return
	}
	res, err := funcmgr.CallFunc("Service.AddOrUpdateConn", sname, sid, value)
	if err != nil {
		zslog.LogError("EtcdDiscover addOrUpdate callfunc Service.AddOrUpdateConn err")
		return
	}
	isSuccess := res[0].(bool)
	if isSuccess {
		if _, ok := ed.ServiceMap[sname]; !ok {
			ed.ServiceMap[sname] = make(map[int][]byte)
		}
		ed.ServiceMap[sname][sid] = value
	}
	zslog.LogDebug("etcdDiscover addOrUpdate Map :%v", ed.ServiceMap)
}

func (ed *etcdDiscover) delete(key []byte) {
	zslog.LogDebug("etcdDiscover delete key:%v", key)
	ed.lock.Lock()
	defer ed.lock.Unlock()

	sname, sid, err := tailKey(key)
	if err != nil {
		zslog.LogError("delete tailKey err,key : %v", key)
		return
	}
	_, err = funcmgr.CallFunc("Service.DeleteConn", sname, sid)
	if err != nil {
		zslog.LogError("EtcdDiscover delete call func Service.DeleteConn err")
		return
	}
	if _, ok := ed.ServiceMap[sname]; ok {
		if _, ok := ed.ServiceMap[sname][sid]; ok {
			delete(ed.ServiceMap[sname], sid)
		}
	}

	zslog.LogDebug("etcdDiscover delete Map : %v", ed.ServiceMap)
}

type etcdRegister struct {
	ttl        int64
	svrInfo    map[string]interface{}
	leaseID    clientv3.LeaseID
	etcdClient *clientv3.Client
}

func newEtcdRegister(svrInfo map[string]interface{}) *etcdRegister {
	endpoints := conf.GetStringSlice("etcd.endpoints")
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: 5 * time.Second})
	if err != nil {
		panic(err)
	}

	er := &etcdRegister{
		svrInfo:    svrInfo,
		ttl:        5,
		leaseID:    0,
		etcdClient: cli,
	}
	return er
}

func (er *etcdRegister) Start() {
	go er.keepAlive()
}

func (er *etcdRegister) keepAlive() {
	duration := time.Duration(er.ttl) * time.Second
	timer := time.NewTimer(duration)
	for {
		select {
		case <-timer.C:
			if er.leaseID > 0 {
				if err := er.leaseRenewal(); err != nil {
					fmt.Println("leaseRenewal err : ", err)
					er.leaseID = 0
				}
			} else {
				if err := er.register(); err != nil {
					fmt.Println("register err : ", err)
				} else {
					fmt.Println("register success,service: ", er.svrInfo["sname"])
				}
			}
			timer.Reset(duration)
		}
	}
}

func (er *etcdRegister) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := er.etcdClient.Delete(ctx, er.registerKey())
	return err
}

func (er *etcdRegister) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := er.etcdClient.Grant(ctx, er.ttl+3)
	if err != nil {
		return err
	}
	_, err = er.etcdClient.Put(ctx, er.registerKey(), er.svrInfo["info"].(string), clientv3.WithLease(resp.ID))
	er.leaseID = resp.ID
	return err
}

func (er *etcdRegister) leaseRenewal() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := er.etcdClient.KeepAliveOnce(ctx, er.leaseID)
	return err
}

func (er *etcdRegister) registerKey() string {
	return fmt.Sprintf("%v_%v_%v", Prefix, er.svrInfo["sname"], er.svrInfo["sid"])
}

func tailKey(key []byte) (string, int, error) {
	keyStr := string(key)
	topicSlice := strings.Split(keyStr, "_")

	if len(topicSlice) != 3 {
		return "", 0, errors.New("tailKey err ,not enough arguments")
	}

	sname := topicSlice[1]
	sid, err := strconv.Atoi(topicSlice[2])
	if err != nil {
		return "", 0, err
	}
	return sname, sid, nil
}
