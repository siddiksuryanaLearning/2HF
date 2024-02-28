# 2HF

Hunting For Halal Food

//docker
sudo apt-get update
sudo apt-get install docker.io
sudo chmod 777 /var/run/docker.sock
sudo apt-get install python3-pip
pip install docker==6.1.3

//clone
git clone -b mvc <remote-repo-url>

//Go
wget https://golang.org/dl/go1.20.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
echo 'PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

//Deploy
cd app
git pull origin mvc

//Test
go test -v ./...

//Running Docker
docker-compose down && docker-compose up --build -d

//Running Container
docker logs <ContainerID>

//Setup

- Swagger Host
- Script deploy.yml
- Github Secret
- Connection Db (db.go, docker-compose.yaml )
