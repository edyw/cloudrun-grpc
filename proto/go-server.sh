protoc \
    --go_out=../go-contact/proto --go_opt=paths=source_relative \
    --go-grpc_out=../go-contact/proto --go-grpc_opt=paths=source_relative \
    contact.proto