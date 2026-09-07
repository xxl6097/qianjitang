package main

import (
	"fmt"
	"os"

	"github.com/xxl6097/go-service/pkg/ukey"
	"github.com/xxl6097/go-thousand-hub/pkg/qjt"
	_ "github.com/xxl6097/qianjitang/cmd/test/buffer"
)

func main() {
	fmt.Println(len(ukey.GetBuffer()))
	certs, _ := os.ReadFile("./rc-agent/server.crt")
	opts := qjt.Options{
		ServerURL: "wss://103.42.30.173:8080/ws/agent",
		Token:     "zhujiangjiayuan2026",
		IDFile:    "/var/lib/rc-agent/id",
		Name:      "ten",
		CAData:    certs,
	}
	data, _ := ukey.StructToGob(&opts)
	fmt.Println(len(data), string(data))
}
