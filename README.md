# usage

## 编译

1)protobuf使用2.5.0

cd protobuf-2.5.0

./configure --prefix=/usr/local/protobuf

sudo make install

export PATH=$PATH:/usr/local/protobuf/bin

2)protoc-gen-go暂时使用v1.0.0

cd go/src/github.com/golang/protobuf/protoc-gen-go

git checkout v1.0.0

go build

cp protoc-gen-go /usr/local/bin/

3)pb.go hadoop使用branch-3.2.0

protoc --proto_path=/Users/kinderyj/go/src/github.com/apache/hadoop/hadoop-common-project/hadoop-common/src/main/proto/  --proto_path=. --go_out=build *.proto


## 调试&demo

1. 登陆emr集群， su hadoop
2. cd /home/hadoop/go/src/github.com/hortonworks/gohadoop/hadoop_yarn/examples/dist_shell
3. go run emr.go #查询集群内node信息，查询指定了appId的信息

## TODO

为proto文件添加resourceRequests并编译pb文件
