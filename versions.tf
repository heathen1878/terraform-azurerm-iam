terraform {
  required_version = ">= 1.5.5, <= 1.9.3"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.74.0, <= 4.15.0"
    }
  }
}