proto:
	cd proto && \
	protoc --go_out=../build/go \
		   --go-grpc_out=../build/go \
		   --go_opt=paths=source_relative \
		   --go-grpc_opt=paths=source_relative \
		   ./*.proto && \
	cd ..