include $(shell go list -m -f '{{.Dir}}')/make/proto_path.mk

regenerate:
	protoc --proto_path=$(GOGO_PROTOBUF_PROTO_PATH) \
	--gogo_out=. issue312.proto 
	(cd events && make regenerate)
