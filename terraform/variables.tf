variable "aws_region" {
  description = "AWS region"
  default     = "us-east-1"
}

variable "instance_type" {
  description = "EC2 instance type"
  default     = "t3.micro"
}

variable "app_port" {
  description = "The port the application listens on"
  default     = 1313
}

variable "key_name" {
  description = "Name of the SSH key pair to use"
  type        = string
}

variable "profile" {
  description = "Profile configured on CLI"
  type = string
}

