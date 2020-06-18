### usage

编译：

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

cd /Users/kinderyj/go/src/github.com/hortonworks/gohadoop/hadoop_yarn

protoc --proto_path=../hadoop_common --proto_path=. --go_out=build yarn_protos.proto



