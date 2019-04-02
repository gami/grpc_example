.PHONY: setup
setup:
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: protoc
protoc:
	protoc \
		-Iproto \
		--go_out=plugins=grpc:api \
		proto/*.proto

.PHONY: run
run:
	cd grpc && fresh