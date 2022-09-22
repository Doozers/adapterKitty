.PHONY: generate
generate:
	cd api && protoc --go_out=../proto --go_opt=paths=source_relative --go-grpc_out=../proto --go-grpc_opt=paths=source_relative adapterkit.proto && cd -

.PHONY: goconsole.generate
goconsole.generate:
	cd api && protoc --go_out=../AK-interfaces/go-console/proto --go_opt=paths=source_relative --go-grpc_out=../AK-interfaces/go-console/proto --go-grpc_opt=paths=source_relative adapterkit.proto && cd -
