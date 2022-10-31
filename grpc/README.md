## 安装 

Protocol Buffers v3
```shell
cd ~/Downloads
wget https://github.com/protocolbuffers/protobuf/releases/download/v21.9/protobuf-all-21.9.zip
unzip protobuf-all-21.9.zip
cd protobuf-all-21.9.zip
./configure
make
make install

# check:
protoc --version
```

Go Protoc Plugin
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## 使用 protoc 生成对应代码
```shell
cd proto
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
./*.proto
```