package agent

import (
	"github.com/xxl6097/go-service/pkg/ukey"
	"github.com/xxl6097/go-thousand-hub/pkg/qjt"
)

var cfgData *CfgModel

type CfgModel struct {
	Opts qjt.Options `json:"opts"`
}

func LoadMemBuffer() (*qjt.Options, error) {
	byteArray, err := ukey.Load()
	if err != nil {
		return nil, err
	}
	var cfg qjt.Options
	err = ukey.GobToStruct(byteArray, &cfg)
	if err != nil {
		return nil, err
	}
	cfgData = &CfgModel{
		Opts: cfg,
	}
	return &cfg, nil
}

func getCfg() *CfgModel {
	if cfgData == nil {
		_, err := LoadMemBuffer()
		if err != nil {
			return nil
		}
	}
	return cfgData
}
