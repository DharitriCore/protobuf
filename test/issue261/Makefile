include $(shell go list -m -f '{{.Dir}}')/make/proto_path.mk

regenerate:
	go install github.com/gogo/protobuf/protoc-min-version
	go install github.com/gogo/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --gogoslick_out=\
	Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
	:. \
	--proto_path=$(GOGO_PROTOBUF_PROTO_PATH) issue261.proto
