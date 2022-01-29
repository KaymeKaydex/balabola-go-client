PWD = $(shell pwd)

# RUN tests
.PHONY: test
test:
	go test $(PWD)/... -coverprofile=cover.out
