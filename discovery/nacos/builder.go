package nacos

import (
	"context"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/pkg/errors"
	"google.golang.org/grpc/resolver"
)

const schemeName = "nacos"

type builder struct {
	client naming_client.INamingClient
}

func (b *builder) Build(url resolver.Target, conn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	tgt, err := parseURL(url.URL)
	if err != nil {
		return nil, errors.Wrap(err, "Wrong nacos URL")
	}

	ctx, cancel := context.WithCancel(context.Background())
	pipe := make(chan []string)

	go b.client.Subscribe(&vo.SubscribeParam{
		ServiceName:       tgt.Service,
		Clusters:          tgt.Clusters,
		GroupName:         tgt.GroupName,
		SubscribeCallback: newWatcher(ctx, cancel, pipe).CallBackHandle(), // required
	})

	go populateEndpoints(ctx, conn, pipe)

	return &resolvr{cancelFunc: cancel}, nil
}

// Scheme returns the scheme supported by this resolver.
// Scheme is defined at https://github.com/grpc/grpc/blob/master/doc/naming.md.
func (b *builder) Scheme() string {
	return schemeName
}

func NewBuilder(c naming_client.INamingClient) resolver.Builder {
	return &builder{client: c}
}
