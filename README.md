# GRPC Server Example
This is an basic example of GRPC.

- [GRPC Server Example](#grpc-server-example)
  - [Dependencies](#dependencies)
  - [starting the Example](#starting-the-example)
  - [protobuff](#protobuff)
- [Author](#author)

## Dependencies
To install your dependencies you need to type the following command:

```
go mod download
```
> go version go1.15.2 windows/amd64
## starting the Example
To run the GRPC Server just type in the following command:

```
go run server.go
```

For the client you need to type in this command:

```
go run client.go
```

## protobuff

to translate the `.proto` file into go code you first need to install the binary on your local machine.
the current releases can be found at [ProtocolBuffers Releases](https://github.com/protocolbuffers/protobuf/releases/tag/v3.13.0).

The following command shows you how to translate the `.proto` file into golang code.

> make you sure you have created the folder where you want to but in your go code

```
/path/to/binary/protoc.exe --go_out=plugins=grpc:chat .\chat.proto
```

# Author

[Rafael Lazenhofer](https://github.com/RaLazo)