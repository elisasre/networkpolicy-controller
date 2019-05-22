OPERATOR_NAME := networkpolicy-controller
IMAGE := elisaoyj/$(OPERATOR_NAME)
ifeq ($(USE_JSON_OUTPUT), 1)
GOTEST_REPORT_FORMAT := -json
endif

.PHONY: clean deps test gofmt run ensure check build build-image build-linux-amd64

clean:
	git clean -Xdf

deps:
	go get -u golang.org/x/lint/golint
	go get -u github.com/golang/dep/cmd/dep

test:
	go test ./... -v -coverprofile=gotest-coverage.out $(GOTEST_REPORT_FORMAT) > gotest-report.out && cat gotest-report.out || (cat gotest-report.out; exit 1)
	golint -set_exit_status cmd/... pkg/... > golint-report.out && cat golint-report.out || (cat golint-report.out; exit 1)
	./hack/gofmt.sh

gofmt:
	./hack/gofmt.sh

check:
	dep check|grep "lock is out of sync"; test $$? -eq 1

ensure:
	dep ensure -v

run: build
	./bin/$(OPERATOR_NAME)

build-linux-amd64:
	rm -f bin/linux/$(OPERATOR_NAME)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -i -o bin/linux/$(OPERATOR_NAME) ./cmd

build:
	rm -f bin/$(OPERATOR_NAME)
	go build -v -i -o bin/$(OPERATOR_NAME) ./cmd

build-image: build-linux-amd64
	docker build -t $(IMAGE):latest .
