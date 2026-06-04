test:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install github.com/gogo/protobuf/protoc-min-version
	go get golang.org/x/net/context@v0.29.0
	go get google.golang.org/grpc@v1.68.1
	tmp=$$(mktemp -d); \
	mkdir -p "$$tmp/github.com/gogo"; \
	ln -s "$$(go list -m -f '{{.Dir}}')" "$$tmp/github.com/gogo/protobuf"; \
	protoc-min-version --version="3.0.0" --gogo_out=plugins=grpc:. --proto_path="$$tmp" --proto_path=../../protobuf --proto_path=. issue427.proto; \
	rm -rf "$$tmp"
	go test ./...
