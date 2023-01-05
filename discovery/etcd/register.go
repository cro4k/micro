package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

func RegisterGRPCResolver(client *clientv3.Client) {
	resolver.Register(NewBuilder(client))
}
