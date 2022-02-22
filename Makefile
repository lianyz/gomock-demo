.PHONY: gen
gen:
	mkdir -p mocks
	go generate ./...

.PHONY: test
test:
	go test -v github.com/lianyz/gomock-demo/user