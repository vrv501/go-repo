export GO_MODULE = github.com/vrv501/go-repo/grpc
protogen:
	@protoc -Iproto --go_out=. --go-grpc_out=. --go_opt=module=${GO_MODULE} --go-grpc_opt=module=${GO_MODULE} proto/*.proto
.PHONY: protogen