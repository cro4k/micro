package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"google.golang.org/grpc/resolver"
)

func RegisterGRPCResolver(client naming_client.INamingClient) {
	resolver.Register(NewBuilder(client))
}
