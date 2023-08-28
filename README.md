## Setup instructions

- Install proto compiler [protoc]      
``sudo apt install protobuf-compiler ``


- Install the following go dependencies    

``go get google.golang.org/grpc/cmd/protoc-gen-go-grpc``

``go get google.golang.org/genproto``

- Run the server

``go run server/main.go``


- Run the client

``go run client/main.go``


- To generate additional protocol buffers

``protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
*.proto``