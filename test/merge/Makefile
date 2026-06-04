include $(shell go list -m -f '{{.Dir}}')/make/proto_path.mk

regenerate:
	go install github.com/gogo/protobuf/protoc-min-version
	go install github.com/gogo/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=. \
	--proto_path=$(GOGO_PROTOBUF_PROTO_PATH) merge.proto
