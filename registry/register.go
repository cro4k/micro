package registry

import (
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

// Register 服务注册
// client etcd客户端连接
// service 服务名称
// addr 服务grpc监听地址
// ttl 过期时间（租约时间）
func Register(client *clientv3.Client, service, addr string, ttl ...int64) (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	em, err := endpoints.NewManager(client, service)
	if err != nil {
		return nil, err
	}

	// 设置租约
	if len(ttl) > 0 && ttl[0] > 0 {
		lease, err := client.Grant(client.Ctx(), 300)
		if err != nil {
			return nil, err
		}
		ch, err := client.KeepAlive(client.Ctx(), lease.ID)
		if err != nil {
			return ch, err
		}
		return ch, em.AddEndpoint(client.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID))
	}
	return nil, em.AddEndpoint(client.Ctx(), service+"/"+addr, endpoints.Endpoint{Addr: addr})
}
