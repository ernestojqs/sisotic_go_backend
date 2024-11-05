package lib

import (
	"github.com/pjmd89/goutils/jsonutils"
	"github.com/pjmd89/goutils/systemutils"
	"github.com/pjmd89/goutils/systemutils/debugmode"
)

var MyConfig SysConf
var Logs systemutils.Logs

type SysConf struct {
	HTTPConfigFile string `json:"httpConfigFile"`
	DBConfigFile   string `json:"dbConfigFile"`
	FilesFolder    string `json:"filesFolder"`
	SystemLog      string `json:"systemlog"`
	AccessLog      string `json:"accesslog"`
}

func Config() (r SysConf) {
	filePath := "etc/config.json"
	if debugmode.Enabled {
		filePath = "etc/config.json"
	}
	jsonutils.GetJson(filePath, &r)
	return
}
