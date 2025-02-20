# Before running, install Terraform utils if you don't havem them

# sudo yum install -y yum-utils shadow-utils
# sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/AmazonLinux/hashicorp.repo
# sudo yum -y install terraform

// Define basic configurations
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

// Configure the AWS provider
provider "aws" {
  region = "us-east-1"
}

// Create a VPC
resource "aws_vpc" "web_server_vpc" {
  cidr_block = "10.0.0.0/16"
}

// Create a public subnet
resource "aws_subnet" "pc_api_public_subnet" {
  vpc_id = aws_vpc.web_server_vpc.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "us-east-1a"
  map_public_ip_on_launch = true
}

// Create a private subnet 1
resource "aws_subnet" "pc_api_private_subnet1" {
  vpc_id = aws_vpc.web_server_vpc.id
  cidr_block = "10.0.2.0/24"
  map_public_ip_on_launch = false
  availability_zone_id = "use1-az2"
}

// Create a private subnet 2
resource "aws_subnet" "pc_api_private_subnet2" {
  vpc_id = aws_vpc.web_server_vpc.id
  cidr_block = "10.0.4.0/24"
  map_public_ip_on_launch = false
  availability_zone_id = "use1-az3"
}

// Create an Internet Gateway
resource "aws_internet_gateway" "pc_api_ig" {
  vpc_id = aws_vpc.web_server_vpc.id
  tags = {
    Name= "Internet Gateway for public subnet"
  }
}

// Create a public route table
resource "aws_route_table" "pc_api_public_route_table" {
  vpc_id = aws_vpc.web_server_vpc.id
  tags = {
    Name= "Route table for public subnet"
  }
}

// Create a private route table
resource "aws_route_table" "pc_api_private_route_table" {
  vpc_id = aws_vpc.web_server_vpc.id
  tags = {
    Name= "Route table for public subnet"
  }
}

// Associate public route table with public subnet
resource "aws_route_table_association" "pc_api_public_rta" {
  route_table_id = aws_route_table.pc_api_public_route_table.id
  subnet_id = aws_subnet.pc_api_public_subnet.id
}

// Associate private route table with private subnet 1
resource "aws_route_table_association" "pc_api_private_rta1" {
  route_table_id = aws_route_table.pc_api_private_route_table.id
  subnet_id = aws_subnet.pc_api_private_subnet1.id
}

// Associate private route table with private subnet 2
resource "aws_route_table_association" "pc_api_private_rta2" {
  route_table_id = aws_route_table.pc_api_private_route_table.id
  subnet_id = aws_subnet.pc_api_private_subnet2.id
}

// Create a route in the public route table
resource "aws_route" "pc_api_public_route" {
  route_table_id = aws_route_table.pc_api_public_route_table.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.pc_api_ig.id
}

// Create a security group for the web server
resource "aws_security_group" "web_server_sg" {
  vpc_id = aws_vpc.web_server_vpc.id
  name = "pc-api-web-server-sg"

  // Allow SSH access
  ingress {
    protocol  = "tcp"
    from_port = 22
    to_port   = 22
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Allow HTTP access
  ingress {
    protocol  = "tcp"
    from_port = 80
    to_port   = 80
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Allow HTTPS access
  ingress {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Allow access to service
  ingress {
    protocol  = "tcp"
    from_port = 1323
    to_port   = 1323
    cidr_blocks = ["0.0.0.0/0"]
  }

  // Allow all outbound traffic
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

// Create an EC2 instance for the web server
resource "aws_instance" "web_server" {
  ami         = "ami-04b70fa74e45c3917" // Ubuntu image
  instance_type    = "t2.micro" // Instance type
  security_groups = [aws_security_group.web_server_sg.id] // Associate with security group
  subnet_id = aws_subnet.pc_api_public_subnet.id
  associate_public_ip_address = true
  count = 1

  // User data script to configure the instance
  user_data = <<EOF
#!/bin/bash
echo "Atualizando apt-get..."
sudo apt-get update
echo "Instalando dependencias..."
sudo apt-get install nginx git golang redis-server -y
sudo snap install aws-cli --classic

echo "Indo para pasta do usuario"
sudo chmod -R 777 /home/ubuntu
cd /home/ubuntu

IP_CUR_EC2=$(curl http://checkip.amazonaws.com)
echo "IP publico da instancia"
echo $IP_CUR_EC2

echo "Fazendo download do projeto..."
git clone https://github.com/evertonmj/price-compare-api.git

echo "Iniciando o servico..."

export GOPATH=/home/ubuntu/price-compare-api/app/src
cd /home/ubuntu/price-compare-api/app/src

echo "Compilando projeto..."
sudo go mod init
sudo go mod tidy
sudo go env -w GOCACHE
sudo go build -o price-app
sudo chmod +x price-app

sudo nohup ./price-app &

echo "Instalacao concluida!"
EOF

  tags = {
    Name = "Ever Instance Test"
  }
}