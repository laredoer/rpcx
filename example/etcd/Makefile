build:
#告知Go编译器生成二进制文件的目标环境：amd64CPU的Linux系统
	GOOS=linux GOARCH=amd64 go build -tags etcd server.go
#根据当前目录下的Dockerfile生成名为server的镜像
	docker build -t server .
run:
	./server