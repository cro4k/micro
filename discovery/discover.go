package discovery

import (
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Discover 服务发现
// client etcd客户端连接
// service 服务名称
// @return
// *grpc.ClientConn grpc连接
func Discover(client *clientv3.Client, service string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	rsv, err := resolver.NewBuilder(client)
	if err != nil {
		return nil, err
	}
	opt := []grpc.DialOption{grpc.WithResolvers(rsv)}
	if len(options) > 0 {
		opt = append(opt, options...)
	}
	return grpc.DialContext(client.Ctx(), "etcd:///"+service, opt...)
}

func DiscoverInsecure(client *clientv3.Client, service string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	opt := append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return Discover(client, service, opt...)
}
