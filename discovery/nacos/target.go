package nacos

import (
	"github.com/pkg/errors"
	"net/url"
	"strings"
	"time"
)

type target struct {
	//Addr        string        `key:",optional"`
	//User        string        `key:",optional"`
	//Password    string        `key:",optional"`
	Service     string        `key:",optional"`
	GroupName   string        `key:",optional"`
	Clusters    []string      `key:",optional"`
	NamespaceID string        `key:"namespaceid,optional"`
	Timeout     time.Duration `key:"timeout,optional"`

	//LogLevel string `key:",optional"`
	//LogDir   string `key:",optional"`
	//CacheDir string `key:",optional"`
}

// parseURL with parameters
func parseURL(rawURL url.URL) (target, error) {
	if rawURL.Scheme != schemeName ||
		len(rawURL.Host) == 0 || len(strings.TrimLeft(rawURL.Path, "/")) == 0 {
		return target{},
			errors.Errorf("Malformed URL('%s'). Must be in the next format: 'nacos://[user:passwd]@host/service?param=value'", rawURL.String())
	}

	var tgt target
	//TODO decode url.Values to target
	//params := make(map[string]interface{}, len(rawURL.Query()))
	//for name, value := range rawURL.Query() {
	//	params[name] = value[0]
	//}
	//err := mapping.UnmarshalKey(params, &tgt)
	//if err != nil {
	//	return target{}, errors.Wrap(err, "Malformed URL parameters")
	//}

	if tgt.NamespaceID == "" {
		tgt.NamespaceID = "public"
	}

	tgt.Service = strings.TrimLeft(rawURL.Path, "/")

	return tgt, nil
}
