## secrets-consumer-env aws

Secrets Consumer for AWS Secret Manager

### Synopsis

AWS secret manager can hold secrets in a json format. the secret can be rotated using a lambda function
and the only versions that AWS secret manager knows are CURRENT_VERSION and PREVIOUS_VERSION
you have the option of specifying PREVIOUS_VERSION=true to fetch previous version

```
secrets-consumer-env aws [flags]
```

### Options

```
  -h, --help                      help for aws
      --previous-version string   If using lambda to rotate secrets you can get the previous version (default: current version)
      --region string             AWS Region for the Secret Manager (default: us-east-1) (default "us-east-1")
      --role-arn string           AWS Role ARN with access to the secret, this requires also permissions on the KMS key for that role
      --secret-name string        AWS Secret Name
```

### Options inherited from parent commands

```
      --config string      config file (default is $HOME/.secrets-consumer-env.yaml)
  -v, --verbosity string   Log level (debug, info, warn, error, fatal, panic (default "info")
```

### SEE ALSO

* [secrets-consumer-env](secrets-consumer-env.md)	 - Consume secrets from AWS, GCP or Hashicorp Vault

###### Auto generated by spf13/cobra on 29-Apr-2020
