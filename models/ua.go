package models

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/httplib"
)

var ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 SP-engine/2.14.0 main%2F1.0 baiduboxapp/11.18.0.16 (Baidu; P2 13.3.1) NABar/0.0"

func initUserAgent() {
	if Config.UserAgent != "" {
		logs.Info("使用自定义User-Agent")
		ua = Config.UserAgent
	} else {
		logs.Info("更新User-Agent")
		var err error
		ua, err = httplib.Get("https://ghproxy.com/https://raw.githubusercontent.com/cdle/jd_study/main/xdd/ua.txt").String()
		if err != nil {
			logs.Info("更新User-Agent失败")
		}
	}

}

func GetUserAgent() string {
	return ua
}
