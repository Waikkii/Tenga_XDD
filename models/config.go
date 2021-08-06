package models

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"gopkg.in/yaml.v2"
)

type Container struct {
	Type      string
	Name      string
	Default   bool
	Address   string
	Username  string
	Password  string
	Path      string
	Version   string
	Token     string
	Available bool
	Delete    []string
	Weigth    int
	Mode      string
	Reader    *bufio.Reader
	Config    string
}
type Yaml struct {
	Containers []Container
	Qrcode     string
	Master     string
	Mode       string
	Static     string
	Database   string
	QywxKey    string `yaml:"qywx_key"`
	Resident   string
	UserAgent  string `yaml:"user_agent"`
	Theme      string
}

var Balance = "balance"
var Parallel = "parallel"

var Config Yaml

func initConfig() {
	confDir := ExecPath + "/conf"
	if _, err := os.Stat(confDir); err != nil {
		os.MkdirAll(confDir, os.ModePerm)
	}
	for _, name := range []string{"app.conf", "config.yaml"} {
		f, err := os.OpenFile(ExecPath+"/conf/"+name, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			logs.Warn(err)
		}
		s, _ := ioutil.ReadAll(f)
		if len(s) == 0 {
			logs.Info("下载配置%s", name)
			r, err := httplib.Get("https://ghproxy.com/https://raw.githubusercontent.com/cdle/jd_study/main/xdd/conf/" + name).Response()
			if err == nil {
				io.Copy(f, r.Body)
			}
		}
		f.Close()
	}
	content, err := ioutil.ReadFile(ExecPath + "/conf/config.yaml")
	if err != nil {
		logs.Warn("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &Config) != nil {
		logs.Warn("解析config.yaml出错: %v", err)
	}
	if Config.Master == "" {
		Config.Master = "xxxx"
	}
	if Config.Mode != Parallel {
		Config.Mode = Balance
	}
	if Config.Qrcode != "" {
		Config.Theme = Config.Qrcode
	}
}
