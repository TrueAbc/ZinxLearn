package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"trueabc.top/zinx/ziface"
)

/*
存储全局配置文件信息
提供给其他模块使用
*/

type GlobalObj struct {
	TcpServer ziface.IServer
	Host      string // 当前的监听IP
	TcpPort   int64
	Name      string

	Version        string // 当前zinx版本号
	MaxConn        int    // 当前服务器的最大链接数
	MaxPackageSize uint32 // 单次数据包的最大值
}

/*
定义一个全局的对外对象
*/

var GlobalObject *GlobalObj

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func (o *GlobalObj) Reload() {
	fmt.Println("当前路径是: ", GetAppPath())
	data, err := ioutil.ReadFile("conf/zinx.json")
	// 将json数据解析到struct

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	// 配置文件未加载的默认值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "v0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}
