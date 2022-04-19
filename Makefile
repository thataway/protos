export GOSUMDB=off
export GO111MODULE=on

$(value $(shell [ ! -d "$(CURDIR)/bin" ] && mkdir -p "$(CURDIR)/bin"))
export GOBIN=$(CURDIR)/bin
GOLANGCI_BIN?=$(GOBIN)/golangci-lint
GOLANGCI_REPO?=https://github.com/golangci/golangci-lint
GOLANGCI_LATEST_VERSION?= $(shell git ls-remote --tags --refs --sort='v:refname' $(GOLANGCI_REPO)|tail -1|egrep -o "v[0-9]+.*")

ifneq ($(wildcard $(GOLANGCI_BIN)),)
	GOLANGCI_CUR_VERSION:=v$(shell $(GOLANGCI_BIN) --version|sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
	GOLANGCI_CUR_VERSION:=
endif

# install linter tool
.PHONY: install-linter
install-linter:
	$(info GOLANGCI-LATEST-VERSION=$(GOLANGCI_LATEST_VERSION))
ifeq ($(filter $(GOLANGCI_CUR_VERSION), $(GOLANGCI_LATEST_VERSION)),)
	$(info Installing GOLANGCI-LINT $(GOLANGCI_LATEST_VERSION)...)
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_LATEST_VERSION)
	@chmod +x $(GOLANGCI_BIN)
else
	@echo "GOLANGCI-LINT is need not install"
endif


# run full lint like in pipeline
.PHONY: lint
lint: | install-linter go-deps
	$(info GOBIN=$(GOBIN))
	$(info GOLANGCI_BIN=$(GOLANGCI_BIN))
	@$(GOLANGCI_BIN) cache clean && \
	$(GOLANGCI_BIN) run --config=$(CURDIR)/.golangci.yaml -v $(CURDIR)/...

# install project dependencies
.PHONY: go-deps
go-deps:
	$(info Check go modules dependencies...)
	@go mod tidy && go mod vendor && go mod verify && echo -=OK=-

.PHONY: bin-tools
bin-tools:
ifeq ($(wildcard $(GOBIN)/protoc-gen-grpc-gateway),)
	@echo "Install 'protoc-gen-grpc-gateway'"
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
endif
ifeq ($(wildcard $(GOBIN)/protoc-gen-openapiv2),)
	@echo "Install 'protoc-gen-openapiv2'"
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
endif
ifeq ($(wildcard $(GOBIN)/protoc-gen-go),)
	@echo "Install 'protoc-gen-go'"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go
endif
ifeq ($(wildcard $(GOBIN)/protoc-gen-go-grpc),)
	@echo "Install 'protoc-gen-go-grpc'"
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
endif
	@echo "" > /dev/null



APIS?=\
	ipvs/api.proto \
	ipruler/api.proto \
	route/api.proto \
	tunnel/api.proto \
	announcer/api.proto \
	healthcheck/api.proto

.PHONY: generate-api
generate-api: | go-deps bin-tools
	@echo "Generate API from proto..." &&\
	PATH=$(PATH):$(GOBIN) && \
	protoc -I $(CURDIR)/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		-I $(CURDIR)/3d-party \
		-I /usr/local/include \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--go_out $(CURDIR)/pkg/api \
		--go-grpc_out $(CURDIR)/pkg/api \
		--proto_path=$(CURDIR)/api \
		--grpc-gateway_out $(CURDIR)/pkg/api \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt standalone=false \
		$(APIS) &&\
	protoc -I $(CURDIR)/vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		-I $(CURDIR)/3d-party \
		-I /usr/local/include \
		--proto_path=$(CURDIR)/api \
		--openapiv2_out $(CURDIR)/pkg/api \
		--openapiv2_opt logtostderr=true \
		$(APIS) &&\
	echo -=OK=-


.PHONY: test
test:
	$(info Running tests...)
	@go clean -testcache && go test -v ./...


