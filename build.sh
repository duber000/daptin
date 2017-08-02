#!/usr/bin/env bash


cd gomsweb
#npm run build
cd ..


export GOPATH=/media/artpar/ddrive/workspace/newgocode
rm -rf rice-box.go
rice embed-go
go build  -ldflags '-linkmode external -extldflags -static -w' main.go
rice append --exec main



rm -rf docker_dir
mkdir docker_dir

cp main docker_dir/main
cp -Rf gomsweb/dist docker_dir/static

cp Dockerfile docker_dir/Dockerfile

cd docker_dir
docker build -t goms/goms  .

cd ..
docker images | grep goms | grep latest