package backend

import (
	"github.com/obnahsgnaw/api"
	// TODO 增加新增的proto模块
	_ "github.com/obnahsgnaw/appapi/gen/app_backend_api/index/v1"
	"github.com/obnahsgnaw/appapi/service"
)

func Scan(s *api.Server) error {
	if s == nil {
		return nil
	}
	return service.ScanNoAuth(s, "app_backend_api")
}
