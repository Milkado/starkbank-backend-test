output "public_ip" {
  description = "Public IP of the application server"
  value       = aws_instance.app_server.public_ip
}
