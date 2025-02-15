terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_vpc" "web_server_vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "pc_api_public_subnet" {
  vpc_id = aws_vpc.web_server_vpc.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "us-east-1a"
  map_public_ip_on_launch = true
}

resource "aws_subnet" "pc_api_private_subnet1" {
  vpc_id = aws_vpc.web_server_vpc.id
  cidr_block = "10.0.2.0/24"
  map_public_ip_on_launch = false
  availability_zone_id = "use1-az2"
}

resource "aws_subnet" "pc_api_private_subnet2" {
  vpc_id = aws_vpc.web_server_vpc.id
  cidr_block = "10.0.4.0/24"
  map_public_ip_on_launch = false
  availability_zone_id = "use1-az3"
}

resource "aws_internet_gateway" "pc_api_ig" {
  vpc_id = aws_vpc.web_server_vpc.id
  tags = {
    Name= "Internet Gateway for public subnet"
  }
}

resource "aws_route_table" "pc_api_public_route_table" {
  vpc_id = aws_vpc.web_server_vpc.id
  tags = {
    Name= "Route table for public subnet"
  }
}

resource "aws_route_table" "pc_api_private_route_table" {
  vpc_id = aws_vpc.web_server_vpc.id
  tags = {
    Name= "Route table for public subnet"
  }
}

resource "aws_route_table_association" "pc_api_public_rta" {
  route_table_id = aws_route_table.pc_api_public_route_table.id
  subnet_id = aws_subnet.pc_api_public_subnet.id
}

resource "aws_route_table_association" "pc_api_private_rta1" {
  route_table_id = aws_route_table.pc_api_private_route_table.id
  subnet_id = aws_subnet.pc_api_private_subnet1.id
}

resource "aws_route_table_association" "pc_api_private_rta2" {
  route_table_id = aws_route_table.pc_api_private_route_table.id
  subnet_id = aws_subnet.pc_api_private_subnet2.id
}

resource "aws_route" "pc_api_public_route" {
  route_table_id = aws_route_table.pc_api_public_route_table.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.pc_api_ig.id
}

# resource "aws_db_subnet_group" "pc_api_rds_sg" {
#   name = "pc-api-rds-db-sg"
#   subnet_ids = [aws_subnet.pc_api_private_subnet1.id, aws_subnet.pc_api_private_subnet2.id]
# }

resource "aws_security_group" "web_server_sg" {
  #Vinculacao deste security group a VPC criada acima
  vpc_id = aws_vpc.web_server_vpc.id
  name = "pc-api-web-server-sg"

  #Liberacao para acesso SSH
  ingress {
    protocol  = "tcp"
    from_port = 22
    to_port   = 22
    cidr_blocks      = ["0.0.0.0/0"]

  }

  #Liberacao de acesso HTTP
  ingress {
    protocol  = "tcp"
    from_port = 80
    to_port   = 80
    cidr_blocks      = ["0.0.0.0/0"]
  }

  #Liberacao de acesso HTTPS
  ingress {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443
    cidr_blocks      = ["0.0.0.0/0"]
  }

  #regra de saida
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
  }
}

# resource "aws_security_group" "rds_db_sg" {
#   name="pc-api-rds-db-sg"
#   vpc_id = aws_vpc.web_server_vpc.id
# }

# resource "aws_security_group" "rds_sg" {
#   name = "rds-sg"
#   vpc_id = aws_vpc.web_server_vpc.id
#   ingress {
#     from_port = 3306
#     to_port = 3306
#     protocol = "tcp"
#     security_groups = [aws_security_group.web_server_sg.id]
#   }
# }

# resource "aws_vpc_security_group_ingress_rule" "ec2_rds_sgir" {
#   security_group_id = aws_security_group.rds_db_sg.id
#   referenced_security_group_id = aws_security_group.web_server_sg.id
#   from_port= 3306
#   to_port=3306
#   ip_protocol="tcp"
# }

#### Criacao do banco de dados RDS
# resource "aws_db_instance" "pc_db_01" {
#   instance_class = "db.t3.micro"
#   allocated_storage = 10
#   db_name = "price_compare_db"
#   engine = "mysql"
#   engine_version = "8.0"
#   username = "admin"
#   password = "p4ssw0rd"
#   parameter_group_name = "default.mysql8.0"
#   skip_final_snapshot = true
#   vpc_security_group_ids = [aws_security_group.rds_db_sg.id]
#   db_subnet_group_name = aws_db_subnet_group.pc_api_rds_sg.id
# }

resource "aws_instance" "web_server" {
  ami         = "ami-04b70fa74e45c3917" #imagem do ubuntu
  instance_type    = "t2.micro" #tipo da instancia
  security_groups = [aws_security_group.web_server_sg.id] #vinculacao ao security group criado acima
  subnet_id = aws_subnet.pc_api_public_subnet.id
  associate_public_ip_address = true
  count = 1
  # depends_on = [aws_db_instance.pc_db_01]
  user_data = <<EOF
#!/bin/bash
echo "Atualizando apt-get..."
sudo apt-get update
echo "Instalamdp dependencias..."
sudo apt-get install nginx python3 python3-pip git nginx python3-venv -y
sudo snap install aws-cli --classic

echo "indo para pasta do usuario"
cd /home/ubuntu
echo "Criando ambiente Python..."
sudo python3 -m venv /home/ubuntu/web_server
source /home/ubuntu/web_server/bin/activate
echo "Instalando dependencias python..."
sudo /home/ubuntu/web_server/bin/pip install flask flask_restful jsonify sqlalchemy pymysql mysql.connector

IP_CUR_EC2=$(curl http://checkip.amazonaws.com)
echo "IP publico da instancia"

echo "Definindo variaveis de ambiente..."

echo "Criando configuracao NGINX..."
echo "server {
listen 80;
listen [::]:80;
server_name $(echo $IP_CUR_EC2);

location / {
proxy_pass http://127.0.0.1:5000;
include proxy_params;
}
}" | sudo tee /etc/nginx/sites-enabled/pc-site

#restart nginx
echo "Reiniciando nginx..."
sudo systemctl restart nginx

echo "Fazendo download do projeto..."
git clone https://github.com/evertonmj/price-compare-api.git

echo "Iniciando o servico..."
cd /home/ubuntu/price-compare-api-v2/
python3 index.py &
echo "Instalacao concluida!"
EOF
  tags = {
    Name = "Ever Instance Test"
  }
}