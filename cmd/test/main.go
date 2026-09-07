package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

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
	gen(data)
}

func gen(data []byte) {
	cfgBuffer := bytes.Repeat([]byte{byte(ukey.B)}, len(ukey.GetBuffer()))
	cfgNewBytes, err := ukey.GenConfig(data, false)
	if err != nil {
		fmt.Println(fmt.Errorf("文件签名失败：%v", err))
		return
	}
	fmt.Println(string(cfgNewBytes))

	filePath := "./dist/agent-linux-amd64"
	tpl, err := os.Open(filePath)
	if err != nil {
		fmt.Println(fmt.Errorf("打开文件失败：%v", err))
		return
	}
	defer func() {
		_ = tpl.Close()
	}()
	fileName := filepath.Base(filePath)
	fmt.Println(fileName)

	dstFile := filepath.Join("./dist", fmt.Sprintf("%s.bak", fileName))
	outFile, err := os.Create(dstFile)
	if err != nil {
		fmt.Println(fmt.Errorf("创建失败：%v", err))
		return
	}
	defer func() {
		_ = outFile.Close()
	}()

	prevBuffer := make([]byte, 0)
	for {
		thisBuffer := make([]byte, 1024)
		n, err1 := tpl.Read(thisBuffer)
		thisBuffer = thisBuffer[:n]
		tempBuffer := append(prevBuffer, thisBuffer...)
		bufIndex := bytes.Index(tempBuffer, cfgBuffer)
		if bufIndex > -1 {
			tempBuffer = bytes.Replace(tempBuffer, cfgBuffer, cfgNewBytes, -1)
		}
		//w.Write(tempBuffer[:len(prevBuffer)])
		_, _ = outFile.Write(tempBuffer[:len(prevBuffer)])
		prevBuffer = tempBuffer[len(prevBuffer):]
		if err1 != nil {
			break
		}
	}
	if len(prevBuffer) > 0 {
		//w.Write(prevBuffer)
		_, _ = outFile.Write(prevBuffer)
		prevBuffer = nil
	}
}
