package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fimreal/goutils/file"
	"gopkg.in/yaml.v2"
)

type Config struct {
	SecretName      string   `json:"secretName"`
	SecretNamespace string   `json:"secretNamespace"`
	AcmeEmail       string   `json:"acmeEmail"`
	AcmeDomains     []string `json:"acmeDomains"`
	Kubeconfig      string   `json:"kubeconfig"`
	LegoArgs        string   `json:"lego_args"`
}

var Conf = new(Config)

func InitConfigFile() {
	configFile := flag.String("c", "./config.json", "Path to the config file")
	flag.Parse()
	if err := ReadConf(*configFile); err != nil {
		log.Fatal("读取配置文件时出错，启动失败！", err)
	}
}

func ReadConf(configFile string) error {
	log.Printf("读取配置文件[%s]...", configFile)
	if !file.PathExists(configFile) {
		return fmt.Errorf("找不到配置文件[%s]", configFile)
	}
	parseConfigFile(configFile)
	log.Println("解析成功")
	return nil
}

// 解析配置文件
func parseConfigFile(configFile string) error {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("读取配置文件时出现错误！", err)
		return err
	}
	err = yaml.Unmarshal(content, Conf)
	if err != nil {
		return err
	}
	log.Println("配置文件解析完成：", Conf)
	return nil
}
