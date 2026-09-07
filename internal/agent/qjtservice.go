package agent

import (
	"fmt"

	"github.com/kardianos/service"
	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-service/pkg"
	"github.com/xxl6097/go-service/pkg/gs/igs"
)

type QjtService struct {
	timestamp string
	gs        igs.Service
}

func (q *QjtService) OnStop() {
	z.L().Sugar().Debugln("qjt service OnStop")
}

func (q *QjtService) OnShutdown() {
	z.L().Sugar().Debugln("qjt service OnShutdown")
}
func (q *QjtService) OnFinish() {
	z.L().Sugar().Debugln("qjt service OnFinish")
}

func (q *QjtService) OnConfig() *service.Config {
	z.L().Sugar().Debugln("qjt service OnConfig")
	cfg := service.Config{
		Name:        pkg.AppName,
		DisplayName: fmt.Sprintf("qjt_%s", pkg.AppVersion),
		Description: "a system for qjt service",
	}
	return &cfg
}

func (q *QjtService) OnVersion() string {
	z.L().Sugar().Debugln("qjt service OnVersion")
	pkg.Version()
	return pkg.AppVersion
}

func (q *QjtService) OnRun(service igs.Service) error {
	z.L().Sugar().Debugln("qjt service OnRun")
	q.gs = service
	cfg, err := LoadMemBuffer()
	if err != nil {
		return err
	}
	return boot(cfg)
}

func (q *QjtService) GetAny(s string) ([]byte, []string) {
	z.L().Sugar().Debugln("qjt service GetAny")
	return input(), nil
}
