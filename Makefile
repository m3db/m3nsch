SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
include $(SELF_DIR)/.ci/common.mk

SHELL=/bin/bash -o pipefail

html_report := coverage.html
test := .ci/test-cover.sh
test_ci_integration := .ci/test-integration.sh
convert-test-data := .ci/convert-test-data.sh
coverfile := cover.out
coverage_xml := coverage.xml
junit_xml := junit.xml
test_log := test.log
lint_check := .ci/lint.sh

BUILD := $(abspath ./out)
LINUX_AMD64_ENV := GOOS=linux GOARCH=amd64 CGO_ENABLED=0
VENDOR_ENV := GO15VENDOREXPERIMENT=1

SERVICES := \
	m3nsch_server \
	m3nsch_client \

setup:
	mkdir -p $(BUILD)

define SERVICE_RULES

$(SERVICE): setup
	@echo Building $(SERVICE)
	$(VENDOR_ENV) go build -o $(BUILD)/$(SERVICE) ./$(SERVICE)/.

$(SERVICE)-linux-amd64:
	$(LINUX_AMD64_ENV) make $(SERVICE)

endef

services: $(SERVICES)
services-linux-amd64:
	$(LINUX_AMD64_ENV) make services

$(foreach SERVICE,$(SERVICES),$(eval $(SERVICE_RULES)))

all: lint test-ci-unit m3nsch_client m3nsch_server
	@echo "Making all"

lint:
	@which golint > /dev/null || go get -u github.com/golang/lint/golint
	$(VENDOR_ENV) $(lint_check)

test-internal:
	@which go-junit-report > /dev/null || go get -u github.com/sectioneight/go-junit-report
	@$(VENDOR_ENV) $(test) $(coverfile) | tee $(test_log)

test-xml: test-internal
	go-junit-report < $(test_log) > $(junit_xml)
	gocov convert $(coverfile) | gocov-xml > $(coverage_xml)
	@$(convert-test-data) $(coverage_xml)
	@rm $(coverfile) &> /dev/null

test: test-internal
	gocov convert $(coverfile) | gocov report

testhtml: test-internal
	gocov convert $(coverfile) | gocov-html > $(html_report) && open $(html_report)
	@rm -f $(test_log) &> /dev/null

test-ci-unit: test-internal
	@which goveralls > /dev/null || go get -u -f github.com/mattn/goveralls
	goveralls -coverprofile=$(coverfile) -service=travis-ci || echo -e "\x1b[31mCoveralls failed\x1b[m"

test-ci-integration:
	@$(VENDOR_ENV) $(test_ci_integration)

clean:
	echo Cleaning build artifacts...
	go clean
	rm -rf $(BUILD)
	@rm -f *.html *.xml *.out *.test
	echo

.DEFAULT_GOAL := test
.PHONY: test test-xml test-internal testhtml clean all
