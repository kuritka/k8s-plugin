.PHONY: pre-push
pre-push:
	go fmt ./...
	golint -set_exit_status=1 ./...
	go test  ./...
