!/bin/bash
echo "Atualizando apt-get..."
sudo apt-get update
echo "Instalamdp dependencias..."
sudo apt-get install nginx python3 python3-pip git nginx python3-venv mysql-client -y
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
RDS_ENDPOINT="${aws_db_instance.pc_db_01.address}"
export DBHOST=$RDS_ENDPOINT
export DBPORT=3306
export DBUSER="${aws_db_instance.pc_db_01.username}"
export DBPASS="${aws_db_instance.pc_db_01.password}"
export DBNAME="price_compare_db"
echo $RDS_ENDPOINT

echo "Criando tabela"
export MYSQL_PWD="p4ssw0rd"
echo "Criando banco de dados"
mysql -h $RDS_ENDPOINT -u admin -e "use price_compare_db; GRANT ALL PRIVILEGES ON price_compare_db.* TO 'admin'@'%' WITH GRANT OPTION; FLUSH PRIVILEGES; create table price_compare_db.users (id serial primary key, name text, email text, created_at timestamp default current_timestamp);"

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
git clone https://github.com/evertonmj/price-compare-api-v2.git

echo "Iniciando o servico..."
cd /home/ubuntu/price-compare-api-v2/
python3 index.py &
echo "Instalacao concluida!