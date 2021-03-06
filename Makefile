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
FIRST_GOPATH := ./
PROMU        := ./promu

PREFIX                  ?= ./release/
GOFMT_FILES             ?= $$(find . -name '*.go' | grep -v vendor)
VERSIONINFO 			?= $$(cat VERSION|sed "s/\r//g"|sed "s/\n//g")

all: format build tarzip

linux: format buildlinux
format:
	@echo ">> formatting code"
	@gofmt -w $(GOFMT_FILES)

build:
	@echo ">> building binaries"
	@$(PROMU) build --prefix $(PREFIX)
	
# Will build both the front-end as well as the back-end
build-all: build

tarzip:
	@echo ">> tar zip"
	@echo prometheus-webhook-dingtalk-$(VERSIONINFO).tgz
	@tar -zcvf prometheus-webhook-dingtalk-$(VERSIONINFO).tgz release/*

.PHONY: all format build build-all
