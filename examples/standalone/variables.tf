variable "description" {
  default     = "Managed by Terraform"
  description = "The description of the role assignment."
  type        = string
}

variable "principal_id" {
  description = "The principal id of the user, group or service principal."
  type        = string
}

variable "principal_type" {
  default     = null
  description = "The principal type of the role assignment."
  type        = string
}

variable "role_definition_name" {
  description = "The role definition name of the role assignment."
  type        = string
}

variable "scope" {
  description = "The scope of the role assignment."
  type        = string
}