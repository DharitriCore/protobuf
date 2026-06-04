GOGO_PROTOBUF_ROOT ?= $(shell go list -m -f '{{.Dir}}')
GOGO_PROTOBUF_IMPORT_ROOT ?= $(GOGO_PROTOBUF_ROOT)/.proto_path
GOGO_PROTOBUF_PROTO_PATH ?= $(shell \
	mkdir -p "$(GOGO_PROTOBUF_IMPORT_ROOT)/github.com/gogo"; \
	link="$(GOGO_PROTOBUF_IMPORT_ROOT)/github.com/gogo/protobuf"; \
	if [ "$$(readlink "$$link" 2>/dev/null)" != "$(GOGO_PROTOBUF_ROOT)" ]; then \
		rm -rf "$$link"; \
		ln -s "$(GOGO_PROTOBUF_ROOT)" "$$link"; \
	fi; \
	printf '%s:%s:%s' "$(GOGO_PROTOBUF_IMPORT_ROOT)" "$(GOGO_PROTOBUF_ROOT)/protobuf" ".")
