include $(shell go list -m -f '{{.Dir}}')/make/proto_path.mk

regenerate:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install github.com/gogo/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --proto_path=$(GOGO_PROTOBUF_PROTO_PATH) --gogo_out=. a.proto