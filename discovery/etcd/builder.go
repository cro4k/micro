package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	etcdResolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc/resolver"
)

func NewBuilder(c *clientv3.Client) resolver.Builder {
	b, _ := etcdResolver.NewBuilder(c)
	return b
}
