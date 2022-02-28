module golang_snippets

go 1.16

require (
	github.com/aws/aws-sdk-go v1.42.20
	github.com/deckarep/golang-set v1.8.0
	github.com/emicklei/go-restful-openapi/v2 v2.2.1
	github.com/emicklei/go-restful/v3 v3.2.0
	github.com/go-openapi/spec v0.19.5
	github.com/jinzhu/copier v0.3.5
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/robfig/cron v1.2.0
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.21.3 // indirect
	k8s.io/klog/v2 v2.8.0
	sigs.k8s.io/controller-runtime v0.9.5
)

replace k8s.io/klog/v2 => k8s.io/klog/v2 v2.40.1
