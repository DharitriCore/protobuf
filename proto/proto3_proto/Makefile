include $(shell go list -m -f '{{.Dir}}')/make/proto_path.mk

regenerate:
	go install github.com/gogo/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=github.com/gogo/protobuf/proto/test_proto,\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:. \
	--proto_path=../:$(GOGO_PROTOBUF_PROTO_PATH) proto3.proto
	
