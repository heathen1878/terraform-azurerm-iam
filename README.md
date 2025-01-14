# Azure IAM module

## Tests

[![Terratest](https://github.com/heathen1878/terraform-azurerm-iam/actions/workflows/module_tests.yaml/badge.svg)](https://github.com/heathen1878/terraform-azurerm-iam/actions/workflows/module_tests.yaml)

## Security

[![Dependabot](https://img.shields.io/badge/dependabot-active-brightgreen?style=flat-square&logo=dependabot)](https://github.com/heathen1878/terraform-azurerm-iam/security/dependabot)

## Examples

[Standalone](./examples/standalone/README.md)

## Usage

```shell
# Typically nested within another module to manage IAM
module "IAM" {

    source ="heathen1878/iam/azurerm"
    version = "1.1.0"

    # Three mandatory parameters
}
```

## Version 1.0.0

- Supports AzureRM 3.74.0 up to AzureRM 3.116.0
- Supports Terraform core executable 1.5.5 up to 1.9.3
- Creates IAM assignments at the scope provided


## Version 1.1.0

- Supports AzureRM 3.74.0 up to AzureRM 4.15.0