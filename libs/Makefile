all:
	protoc --go_out=./go --go_opt=paths=source_relative  --go-grpc_out=./go --go-grpc_opt=paths=source_relative example.proto
	# Update the PATH below if you want to regenerate the libraries
	protoc --plugin="protoc-gen-ts=/Users/sankarp4/.nvm/versions/node/v12.22.1/lib/node_modules/ts-protoc-gen/bin/protoc-gen-ts" --js_out="import_style=commonjs,binary:./js" --ts_out="service=grpc-web:./js" example.proto
