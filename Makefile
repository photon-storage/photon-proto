.PHONY: default

default:
	$(info 'make env' to install environment for building repo)
	$(info 'make pb' to compile protobuf definitions and generate go files)

env:
	rm -rf .env
	./scripts/update_proto.sh install
	./scripts/update_mockgen.sh install

pb: clean
	./scripts/update_proto.sh
	./scripts/update_mockgen.sh
	go test -v ./...

clean:
	find . \( -name "*.pb.go" -o -name "*.pb.gw.go" \) -not -path "./.env/*" -delete
	find . -name "generated.ssz.go" -delete
	rm -rf ./testing/mockgen
