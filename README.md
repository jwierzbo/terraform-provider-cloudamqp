
# terraform-provider-cloudamqp

This terraform provider is able to manage lifecycle of CloudAMQP instances + alarms configuration for them (works only with paid plans).

Instance part base on: https://github.com/cloudamqp/terraform-provider

## Local build and testing

```shell script
make install-local
cd sample
terraform init
TF_LOG=DEBUG terraform plan
```

WARN: provider logs shows up only on DEBUG level (https://github.com/hashicorp/terraform/issues/16752)!

## Usage
To import existing resource to terraform *resource* use: `terraform import RESOURCE.NAME ID`, e.g.:

```shell script
terraform import cloudamqp_alarm_recipient.test_recipient 20124
```

To import existing resource to terraform *module* use: `terraform import module.MODULE_NAME.RESOURCE.NAME ID`, e.g.:

```shell script
terraform import module.test_instance_bunny.cloudamqp_instance.generic_instance 76460
```

Example code is located under [sample](sample) directory, e.g.: [sample/main.tf](sample/main.tf)
