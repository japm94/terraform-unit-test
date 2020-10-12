terraform {
  required_version = ">= 0.13.4"
}

provider "aws" {
  region = var.aws_region
}

# Deploy an EC2 Instance.
resource "aws_instance" "ec2" {
  # Run an Ubuntu Server 20.04 AMI on the EC2 instance.
  ami                    = "ami-06fd8a495a537da8b"
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.instance.id]

  # When the instance boots, start a web server on port 8080 that responds with "Hello, World!".
  user_data = <<EOF
#!/bin/bash
echo "Hello, World!" > index.html
nohup busybox httpd -f -p 8080 &
EOF
}

# Allow the instance to receive requests on port 8080.
resource "aws_security_group" "instance" {
  name        = "Experiments with Terratest"
  description = "Allow the instance to receive requests on port 8080."
  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "terratest_sg"
  }
}