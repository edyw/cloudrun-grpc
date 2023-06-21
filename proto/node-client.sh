grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:../node-api/proto \
    --grpc_out=grpc_js:../node-api/proto \
    contact.proto