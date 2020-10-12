# Output the instance's public IP address.
output "public_ip" {
  value = aws_instance.ec2.public_ip
}