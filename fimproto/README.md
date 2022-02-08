# fimproto

#fimproto
Contains all shared proto files and generated code

To build the proto files:
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I C:\Users\DonTorge\go\src\github.com\envoyproxy\protoc-gen-validate --validate_out="lang=go:." -I . proto/fim.proto 
```