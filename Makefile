.PHONY: gen
gen:
	mkdir -p mocks
	mockgen -destination=mocks/mock_doer.go -package=mocks github.com/lianyz/gomock-demo/doer Doer,DoDo

.PHONY: test
test:
	go test -v github.com/lianyz/gomock-demo/user