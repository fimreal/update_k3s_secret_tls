package main

import (
	"log"

	"github.com/fimreal/update_k8s_tls_secret/config"
	"github.com/fimreal/update_k8s_tls_secret/lego"
	"github.com/fimreal/update_k8s_tls_secret/secret"
)

func main() {
	log.Printf("任务启动: 更新证书到 secret[%s/%s]", config.Conf.SecretNamespace, config.Conf.SecretName)

	// 加载配置
	config.InitConfigFile()
	clientset, err := config.NewKubeClient()
	if err != nil {
		log.Fatal("加载 kube 配置遇到错误，退出！", err)
	}

	err = lego.Generator(config.Conf)
	if err != nil {
		log.Fatal("使用 lego 生成证书时遇到错误，退出！", err)
	}

	err = secret.UpdateSecret(clientset, config.Conf)
	if err != nil {
		log.Fatal("更新 secret 时遇到错误，退出！", err)
	}
	log.Printf("任务结束: 创建/更新证书到 secret[%s/%s] -> 执行完毕！", config.Conf.SecretNamespace, config.Conf.SecretName)
}
