
#The wild card function is used to search for any .protofiles
#Otherwise Makefile will not know what to look for

PROTO_FILES := $(wildcard proto/*.proto)

create:
	protoc --proto_path=proto/ --go_out=gen/ $(PROTO_FILES)
	protoc --proto_path=proto/ --go-grpc_out=gen/ $(PROTO_FILES)

clean:
	rm gen/proto/*.go