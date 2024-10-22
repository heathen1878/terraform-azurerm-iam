resource "random_id" "this" {
  keepers = {
    subscription_id = data.azurerm_client_config.this.subscription_id
    location        = local.location
  }
  byte_length = 4
}