usage: 
## Directly Use
- ETCD:
```go
package example

import (
    clientv3 "go.etcd.io/etcd/client/v3"
    "google.golang.org/grpc"
    "google.golang.org/grpc/resolver"
    "github.com/cro4k/micro/discovery/etcd"
    "context"
    "strings"
    "fmt"
)

func Discover(ctx context.Context, client *clientv3.Client, name string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
    rsv := etcd.NewBuilder(client)
    if !strings.HasPrefix(name, rsv.Scheme()) {
        name = fmt.Sprintf("%s:///%s", rsv.Scheme(), name)
    }
    var opt = []grpc.DialOption{grpc.WithResolvers(rsv)}
    opt = append(opt,options...)
    return grpc.DialContext(ctx, name, opt...)
}
```
- nacos:

```go
package example

import (
	"context"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/cro4k/micro/discovery/nacos"
	"google.golang.org/grpc"
	"strings"
)

func Discover(ctx context.Context, client naming_client.INamingClient, name string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
    rsv := nacos.NewBuilder(client)
    if !strings.HasPrefix(name, rsv.Scheme()) {
        name = fmt.Sprintf("%s:///%s", rsv.Scheme(), name)
    }
    var opt = []grpc.DialOption{grpc.WithResolvers(rsv)}
    opt = append(opt, options...)
    return grpc.DialContext(ctx, name, opt...)
}

```

## Global Register
1. register global grpc resolver 
```go
package example

import (
    "github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
    clientv3 "go.etcd.io/etcd/client/v3"
    "github.com/cro4k/micro/discovery/nacos"
    "github.com/cro4k/micro/discovery/etcd"
)

func init() {
	// create a new nacos client
	// create a new etcd client
    nacos.RegisterGRPCResolver(nacosClient)
    etcd.RegisterGRPCResolver(etcdClient)
}


```
2. discover service
```go
package example

import (
    "context"
    "google.golang.org/grpc"
)

func Discover(ctx context.Context, name string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
    return grpc.DialContext(ctx, name, options...)
}
```