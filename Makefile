SHELL=/bin/bash -o pipefail

GO_PKG=github.com/coreos/prometheus-operator
REPO?=quay.io/coreos/prometheus-operator
REPO_PROMETHEUS_CONFIG_RELOADER?=quay.io/coreos/prometheus-config-reloader
TAG?=$(shell git rev-parse --short HEAD)

FIRST_GOPATH:=$(firstword $(subst :, ,$(shell go env GOPATH)))
GOJSONTOYAML_BINARY:=$(FIRST_GOPATH)/bin/gojsontoyaml
PO_DOCGEN_BINARY:=$(FIRST_GOPATH)/bin/po-docgen
EMBEDMD_BINARY:=$(FIRST_GOPATH)/bin/embedmd

TYPES_V1_TARGET:=pkg/apis/monitoring/v1/types.go

K8S_GEN_VERSION:=release-1.8
K8S_GEN_BINARIES:=deepcopy-gen informer-gen lister-gen client-gen
K8S_GEN_ARGS:=--go-header-file $(FIRST_GOPATH)/src/$(GO_PKG)/.header

K8S_GEN_DEPS:=.header
K8S_GEN_DEPS+=$(TYPES_V1_TARGET)
K8S_GEN_DEPS+=$(foreach bin,$(K8S_GEN_BINARIES),$(FIRST_GOPATH)/bin/$(bin))

GOLANG_FILES:=$(shell find . -path ./vendor -prune -o -name \*.go -print)
pkgs = $(shell go list ./... | grep -v /vendor/ | grep -v /test/)

.PHONY: all
all: format generate build test

.PHONY: clean
clean:
	# Remove all files and directories ignored by git.
	git clean -Xfd .

############
# Building #
############

.PHONY: build
build: operator prometheus-config-reloader k8s-gen

.PHONY: operator
operator: $(GOLANG_FILES)
	GOOS=linux CGO_ENABLED=0 go build \
	-ldflags "-X $(GO_PKG)/pkg/version.Version=$(shell cat VERSION)" \
	-o $@ cmd/operator/main.go

.PHONY: prometheus-config-reloader
prometheus-config-reloader:
	GOOS=linux CGO_ENABLED=0 go build \
	-ldflags "-X $(GO_PKG)/pkg/version.Version=$(shell cat VERSION)" \
	-o $@ cmd/$@/main.go

DEEPCOPY_TARGET := pkg/apis/monitoring/v1/zz_generated.deepcopy.go
$(DEEPCOPY_TARGET): $(K8S_GEN_DEPS)
	$(DEEPCOPY_GEN_BINARY) \
	$(K8S_GEN_ARGS) \
	--input-dirs    "$(GO_PKG)/pkg/apis/monitoring/v1" \
	--bounding-dirs "$(GO_PKG)/pkg/apis/monitoring" \
	--output-file-base zz_generated.deepcopy

CLIENT_TARGET := pkg/client/versioned/clientset.go
$(CLIENT_TARGET): $(K8S_GEN_DEPS)
	$(CLIENT_GEN_BINARY) \
	$(K8S_GEN_ARGS) \
	--input-base     "" \
	--clientset-name "versioned" \
	--input	         "$(GO_PKG)/pkg/apis/monitoring/v1" \
	--output-package "$(GO_PKG)/pkg/client"

LISTER_TARGET := pkg/client/listers/monitoring/v1/prometheus.go
$(LISTER_TARGET): $(K8S_GEN_DEPS)
	$(LISTER_GEN_BINARY) \
	$(K8S_GEN_ARGS) \
	--input-dirs     "$(GO_PKG)/pkg/apis/monitoring/v1" \
	--output-package "$(GO_PKG)/pkg/client/listers"

INFORMER_TARGET := pkg/client/informers/externalversions/monitoring/v1/prometheus.go
$(INFORMER_TARGET): $(K8S_GEN_DEPS) $(LISTER_TARGET) $(CLIENT_TARGET)
	$(INFORMER_GEN_BINARY) \
	$(K8S_GEN_ARGS) \
	--versioned-clientset-package "$(GO_PKG)/pkg/client/versioned" \
	--listers-package "$(GO_PKG)/pkg/client/listers" \
	--input-dirs      "$(GO_PKG)/pkg/apis/monitoring/v1" \
	--output-package  "$(GO_PKG)/pkg/client/informers"

.PHONY: k8s-gen
k8s-gen: \
 $(DEEPCOPY_TARGET) \
 $(CLIENT_TARGET) \
 $(LISTER_TARGET) \
 $(INFORMER_TARGET) \

.PHONY: image
image: hack/operator-image hack/prometheus-config-reloader-image

hack/operator-image: Dockerfile operator
# Create empty target file, for the sole purpose of recording when this target
# was last executed via the last-modification timestamp on the file. See
# https://www.gnu.org/software/make/manual/make.html#Empty-Targets
	docker build -t $(REPO):$(TAG) .
	touch $@

hack/prometheus-config-reloader-image: cmd/prometheus-config-reloader/Dockerfile prometheus-config-reloader
# Create empty target file, for the sole purpose of recording when this target
# was last executed via the last-modification timestamp on the file. See
# https://www.gnu.org/software/make/manual/make.html#Empty-Targets
	docker build -t $(REPO_PROMETHEUS_CONFIG_RELOADER):$(TAG) -f cmd/prometheus-config-reloader/Dockerfile .
	touch $@


##############
# Generating #
##############

.PHONY: generate
generate: $(DEEPCOPY_TARGET) $(shell find Documentation -type f)

FULLY_GENERATED_DOCS = Documentation/api.md Documentation/compatibility.md
TO_BE_EXTENDED_DOCS = $(filter-out $(FULLY_GENERATED_DOCS), $(wildcard Documentation/*.md))

Documentation/api.md: $(PO_DOCGEN_BINARY) $(TYPES_V1_TARGET)
	$(PO_DOCGEN_BINARY) api $(TYPES_V1_TARGET) > $@

Documentation/compatibility.md: $(PO_DOCGEN_BINARY) pkg/prometheus/statefulset.go
	$(PO_DOCGEN_BINARY) compatibility > $@

$(TO_BE_EXTENDED_DOCS): $(EMBEDMD_BINARY)
	$(EMBEDMD_BINARY) -w `find Documentation -name "*.md" | grep -v vendor`


##############
# Formatting #
##############

.PHONY: format
format: go-fmt check-license shellcheck

.PHONY: go-fmt
go-fmt:
	go fmt $(pkgs)

.PHONY: check-license
check-license:
	./scripts/check_license.sh

.PHONY: shellcheck
shellcheck:
	docker run -v "${PWD}:/mnt" koalaman/shellcheck:stable $(shell find . -type f -name "*.sh" -not -path "*vendor*")


###########
# Testing #
###########

.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	@go test -race $(TEST_RUN_ARGS) -short $(pkgs)

############
# Binaries #
############

# generate k8s generator variable and target,
# i.e. if $(1)=informer-gen:
#
# INFORMER_GEN_BINARY=/home/user/go/bin/informer-gen
#
# /home/user/go/bin/informer-gen:
#	go get -u -d k8s.io/code-generator/cmd/informer-gen
#	cd /home/user/go/src/k8s.io/code-generator; git checkout release-1.11
#	go install k8s.io/code-generator/cmd/informer-gen
#
define _K8S_GEN_VAR_TARGET_
$(shell echo $(1) | tr '[:lower:]' '[:upper:]' | tr '-' '_')_BINARY:=$(FIRST_GOPATH)/bin/$(1)

$(FIRST_GOPATH)/bin/$(1):
	go get -u -d k8s.io/code-generator/cmd/$(1)
	cd $(FIRST_GOPATH)/src/k8s.io/code-generator; git checkout $(K8S_GEN_VERSION)
	go install k8s.io/code-generator/cmd/$(1)

endef

$(foreach binary,$(K8S_GEN_BINARIES),$(eval $(call _K8S_GEN_VAR_TARGET_,$(binary))))

$(EMBEDMD_BINARY):
	@go get github.com/campoy/embedmd

$(PO_DOCGEN_BINARY): $(shell find cmd/po-docgen -type f) $(TYPES_V1_TARGET)
	go install $(GO_PKG)/cmd/po-docgen

$(GOJSONTOYAML_BINARY):
	go get -u github.com/brancz/gojsontoyaml
