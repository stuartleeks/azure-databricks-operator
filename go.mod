module github.com/microsoft/azure-databricks-operator

go 1.12

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.10.3
	github.com/onsi/gomega v1.7.0
	github.com/prometheus/client_golang v0.9.2
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xinsnake/databricks-sdk-golang v0.1.3
	golang.org/x/crypto v0.0.0-20191112222119-e1110fd1c708 // indirect
	golang.org/x/net v0.0.0-20191112182307-2180aed22343
	golang.org/x/sys v0.0.0-20191112214154-59a1497f0cea // indirect
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
	k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
	k8s.io/klog v1.0.0 // indirect
	sigs.k8s.io/controller-runtime v0.4.0
)
