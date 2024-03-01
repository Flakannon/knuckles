variable "handler" {
  type = string
}

variable "function_name" {
  type = string
}

variable "filename" {
  type = string
}

variable "role" {
  type = string
}

variable "app_version" {
  type    = string
  default = "0.0.1"
}


variable "sns_topic_arn" {
  type = string
}
