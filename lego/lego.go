package lego

import (
	"fmt"
	"log"

	"github.com/fimreal/goutils/cmder"
	"github.com/fimreal/update_k8s_tls_secret/config"
)

func Generator(Conf *config.Config) (err error) {
	cmdStr := fmt.Sprintf("/usr/bin/lego -m %s --dns dnspod %s", Conf.AcmeEmail, Conf.LegoArgs)
	for _, d := range Conf.AcmeDomains {
		cmdStr += " -d " + d
	}
	cmdStr += " -a --path /certs --filename tls run"
	log.Println("使用 lego 申请证书...")
	log.Println("执行命令为：", cmdStr)
	err = cmder.SuperCMD(cmdStr)
	if err != nil {
		return err
	}
	return
}
