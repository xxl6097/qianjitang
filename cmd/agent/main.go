package main

import (
	"fmt"

	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-service/pkg"
	"github.com/xxl6097/go-service/pkg/gs"
	"github.com/xxl6097/go-service/pkg/utils"
	"github.com/xxl6097/qianjitang/internal/agent"
	"go.uber.org/zap"
)

func init() {
	if utils.IsMacOs() {
		pkg.AppVersion = "v0.0.3"
		pkg.BinName = "aatest_v0.0.20_darwin_arm64"
		fmt.Println("Hello World...1")
	}
}

func main() {
	pkg.BinName = "tc-agent"
	err := gs.Run(&agent.QjtService{})
	z.L().Debug("程序结束", zap.Error(err))

}
