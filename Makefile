.PHONY: proto data build

GOPATH=/Users/huchenhao/go

proto:
	for d in proto ; do \
		for f in $$d/*/*.proto; do \
			protoc -I$(GOPATH)/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20201016211454-9db36e164668 -I. -I$(GOPATH)/pkg/mod/github.com/protocolbuffers/protobuf@v3.13.0+incompatible/src  --go_out=plugins=grpc:../../../ $$f; \
			echo compiled: $$f; \
		done \
	done


