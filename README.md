
## Secrets Consumer Env

Consume secrets from AWS, GCP or Hashicorp Vault

### Synopsis

There are a few secret managers that holds secrets, the problem becomes how to consume these secrets
securely.

The Secrets Consumer Env creates a new shell environment, and fetch the secrets from the secret engine
adding them to the environment variables on the new shell and then calling the `syscall.execv` which will
replace the running process with the given process, that given process will inherit all environment variables.

In the world of containers, its important that the process running in it should get the PID 1 so
that a sig TERM will work properly.

will have access to the env vars, the operating system / docker container will not have any of the
secrets exposed.

This tool can either run as a standalone outside of kubernetes or using the Kubernetes mutation webhook.

This tool works with the following secrets managers:

* GCP Secret Manager
* AWS Secret Manager
* Hashicorp Vault
  * Kubernetes backend login (Default)
  * GCP backend login

### CLI Commands

* `aws`  - enable the AWS Secret Manager
* `gcp`  - enable the GCP Secret Manager
* `vault`  - enable the Vault Secret Manager

**Note: The double dash symbol “–-” is used to separate the arguments you want to pass to the command from the secrets-consumer-env arguments.**

**Important: Do not use double-quotes for your command as it will first be evaluated by your shell and not by the secrets-consumer-env.**

### Options

```
      --config string      config file (default is $HOME/.secrets-consumer-env.yaml)
  -h, --help               help for secrets-consumer-env
  -t, --toggle             Help message for toggle
  -v, --verbosity string   Log level (debug, info, warn, error, fatal, panic (default "info")
```

### SEE ALSO

* [secrets-consumer-env aws](docs/secrets-consumer-env_aws.md)	 - Secrets Consumer for AWS Secret Manager
* [secrets-consumer-env gcp](docs/secrets-consumer-env_gcp.md)	 - Secrets Consumer for GCP Secret Manager
* [secrets-consumer-env vault](docs/secrets-consumer-env_vault.md)	 - Fetch and inject secrets from Vault to a given command
* [secrets-consumer-env version](docs/secrets-consumer-env_version.md)	 - Print the version of Secrets Consumer Env


