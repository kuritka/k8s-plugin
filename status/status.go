package status

import (
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/kubernetes/pkg/util/env"

	"github.com/kuritka/plugin/common/log"

	"os"
	"path/filepath"
)

type Status struct {}

type Options struct {
	Namespace string
}

var logger = log.Log

func New(options Options) *Status {
	return new(Status)
}

func (s *Status) Run() error {

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	logger.Info().Msgf("Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	bytes, err := ioutil.ReadFile(kubeconfig)
	if err != nil {
		return err
	}
	c2,err := clientcmd.NewClientConfigFromBytes(bytes)

	c3, err := c2.RawConfig()

	//cmd.Flags().GetString("namespace")


	for k,_ := range c3.Clusters {
		logger.Info().Msgf(k)
	}

	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	ns, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _,n := range ns.Items {
		logger.Info().Msgf("%s %s",n.ClusterName, n.Name)
	}
	return nil
}

func (s *Status) String() string {
	return "Status"
}