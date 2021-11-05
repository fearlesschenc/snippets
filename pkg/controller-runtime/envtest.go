package controller_runtime

import (
	//clientset "k8s.io/client-go/kubernetes"
	//. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

func Main() {
	testEnv := &envtest.Environment{}

	cfg, err := testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	_ = testEnv.Stop()
}
