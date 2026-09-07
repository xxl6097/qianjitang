package agent

import (
	"context"
	"fmt"

	"github.com/xxl6097/glog/pkg/z"
	"github.com/xxl6097/go-service/pkg/ukey"
	"github.com/xxl6097/go-service/pkg/utils"
	"github.com/xxl6097/go-thousand-hub/pkg/q"
	"github.com/xxl6097/go-thousand-hub/pkg/qjt"
)

func input() []byte {
	var cfg *qjt.Options
	cfm := getCfg()
	if cfm == nil {
		cfg = &qjt.Options{}
	} else {
		cfg = &cfm.Opts
	}

	if cfg.ServerURL == "" {
		addr := utils.InputString("请输入服务器地址：")
		cfg.ServerURL = fmt.Sprintf("wss://%s/ws/agent", addr)
	}

	if cfg.Token == "" {
		cfg.Token = utils.InputString("请输入token：")
	}

	if cfg.Name == "" {
		cfg.Name = utils.InputString("请输入名称：")
	}
	bb, e := ukey.StructToGob(cfg)
	if e != nil {
		return nil
	}
	return bb
}

func boot(opts *qjt.Options) error {
	if err := q.NewQJT(opts, context.Background()); err != nil {
		z.L().Sugar().Errorf("agent 退出: %v", err)
		return err
	}
	return nil
}
