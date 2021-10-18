package secret

import (
	"context"
	"io/ioutil"

	"github.com/fimreal/update_k8s_tls_secret/config"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func UpdateSecret(clientset *kubernetes.Clientset, Conf *config.Config) (err error) {

	// create secret struct
	newSecret := v1.Secret{
		TypeMeta:   metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
		ObjectMeta: metav1.ObjectMeta{Name: Conf.SecretName, Namespace: Conf.SecretNamespace},
		Data:       map[string][]byte{},
		StringData: map[string]string{},
		Type:       "kubernetes/tls",
	}
	newSecret.Data["tls.key"], err = ioutil.ReadFile("/certs/certificates/tls.key")
	if err != nil {
		return err
	}
	newSecret.Data["tls.crt"], err = ioutil.ReadFile("/certs/certificates/tls.crt")
	if err != nil {
		return err
	}

	// update secret
	secretOut, err := clientset.CoreV1().Secrets(Conf.SecretNamespace).Update(context.TODO(), &newSecret, metav1.UpdateOptions{})
	logrus.Debug(secretOut)
	if err != nil {
		logrus.Errorln(err)
	} else {
		return
	}

	// create secret
	secretOut, err = clientset.CoreV1().Secrets(Conf.SecretNamespace).Create(context.TODO(), &newSecret, metav1.CreateOptions{})
	logrus.Debug(secretOut)
	if err != nil {
		logrus.Errorln(err)
	}
	return

}
