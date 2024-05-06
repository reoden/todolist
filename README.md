# todolist

protoc-gen-micro
```shell
$ go get -u github.com/go-micro/generator/cmd/protoc-gen-micro
```
generate pb https://github.com/go-micro/generator.git
```shell
$ protoc --proto_path=. --micro_out=./pb --go_out=./pb userService.proto
```
install protoc-go-inject-tag to ensure @inject generate in .pb.go
https://github.com/favadi/protoc-go-inject-tag
```shell
$ go install github.com/favadi/protoc-go-inject-tag@latest
```

protoc-go-inject-tag usage
```shell
$ protoc-go-inject-tag -input="./pb/*.pb.go"
```