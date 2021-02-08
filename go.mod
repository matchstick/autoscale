module github.com/matchstick/autoscale

go 1.15

require (
	github.com/golangci/golangci-lint v1.36.0
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	golang.org/x/tools v0.1.0 // indirect
	// As per https://github.com/kubernetes/client-go/issues/551
	// we need to run:
	// go mod init
	// go mod vendor
	// go mod downlaod
	// Then pin the version of:
	// k8s.io/api, k8s.io/apimachinery and k8s.io/client-go
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009 // indirect
)
