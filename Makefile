.PHONY: generate
ak.generate:
	cd api && protoc --go_out=../AK/proto --go_opt=paths=source_relative --go-grpc_out=../AK/proto --go-grpc_opt=paths=source_relative adapterkit.proto && cd -

.PHONY: goconsole.generate
goconsole.generate:
	cd api && protoc --go_out=../AK-interfaces/go-console/proto --go_opt=paths=source_relative --go-grpc_out=../AK-interfaces/go-console/proto --go-grpc_opt=paths=source_relative adapterkit.proto && cd -

.PHONY: grpcweb.generate
grpcwe.generate:
	cd api && protoc -I=. adapterkit.proto  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../AK-interfaces/gRPC-web/src/proto --js_out=import_style=commonjs:../AK-interfaces/gRPC-web/src/proto && cd -

.PHONY: custom.generate
custom.generate:
	cd custom-proto/ && protoc -I=. --go_out=. --js_out=import_style=commonjs:. ./announcement.proto && cd -
