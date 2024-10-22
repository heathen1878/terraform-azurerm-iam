# Azure IAM module

## Examples

[Standalone](./examples/standalone/README.md)
[Nested](./examples/nested/README.md)

```shell
module "IAM" {

    source ="heathen1878/iam/azurerm"
    version = "1.0.0"

    # Three mandatory parameters
}
```

## Version 1.0.0

- Supports AzureRM 3.74.0 up to AzureRM 3.116.0
- Supports Terraform core executable 1.5.5 up to 1.9.3
- Creates IAM assignments at the scope provided
