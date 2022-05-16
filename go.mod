module github.com/innovia/secrets-consumer-env

go 1.15

require (
	cloud.google.com/go v0.75.0
	github.com/aws/aws-sdk-go v1.37.19
	github.com/google/go-cmp v0.5.5
	github.com/googleapis/gax-go/v2 v2.0.5
	github.com/hashicorp/hcl v1.0.1-vault // indirect
	github.com/hashicorp/vault v1.7.6
	github.com/hashicorp/vault-plugin-secrets-kv v0.8.0
	github.com/hashicorp/vault/api v1.1.1
	github.com/hashicorp/vault/sdk v0.2.1
	github.com/magiconair/properties v1.8.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/prometheus/common v0.11.1
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3
	google.golang.org/api v0.36.0
	google.golang.org/genproto v0.0.0-20210113195801-ae06605f4595
	google.golang.org/grpc v1.35.0
)
