# golang

# Instalação
git clone do projeto na pasta ~/projects/go/src/

wget https://storage.googleapis.com/golang/go1.13.7.linux-amd64.tar.gz

sudo tar xvfz go1.13.7.linux-amd64.tar.gz --directory /opt

echo "export GOROOT=/opt/go" >> ~/.bashrc
echo "export PATH=$PATH:$GOROOT/bin" >> ~/.bashrc
echo "export GOPATH=$HOME/projects/go" >> ~/.bashrc
source .bashrc
go version
