# Copyright 2015 The Prometheus Authors
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GO           := GO15VENDOREXPERIMENT=1 go
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
PROMU        := ./promu

PREFIX                  ?= ./
GOFMT_FILES             ?= $$(C:/cygwin64/bin/find . -name '*.go' | grep -v vendor)

all: format build 

linux: format buildlinux
format:
	@echo ">> formatting code"
	@gofmt -w $(GOFMT_FILES)

build:
	@echo ">> building binaries"
	@$(PROMU) build --prefix $(PREFIX)

buildlinux:
	@echo ">> building binaries"
	@$(PROMU) build --prefix $(PREFIX) crossbuild linux/amd64
	
# Will build both the front-end as well as the back-end
build-all: build



.PHONY: all format build build-all
