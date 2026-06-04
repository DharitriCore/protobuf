include $(shell go list -m -f '{{.Dir}}')/make/proto_path.mk

regenerate:
	go install github.com/gogo/protobuf/protoc-min-version
	go install github.com/gogo/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=$(GOGO_PROTOBUF_PROTO_PATH) \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:. timefail.proto 
