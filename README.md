# Web Scraper

This server application will use to retrieve all the links of a given URL. This server application will be run in both REST and gRPC mode by setting the environment variable `ENABLE_GRPC` `true`. By default, `ENABLE_GRPC` will be `false`, so REST server will be up when run this application.

`ADDRESS` environment variable is the combination of host name and port number. Default value is `localhost:9876`.

We can change the log level by setting the environment variable `LOG_LEVEL`, default value is `info`. Available log levels are `panic`,`fatal`,`error`,`warn`,`info`,`debug`,`trace`.

## Development
To add the appropriate dependencies to the `BUILD.bazel` files, `bazel run :gazelle`

To update the repositories.bzl file,
`bazel run :gazelle -- update-repos -from_file=go.mod -prune=true -build_file_proto_mode=disable_global -to_macro=repositories.bzl%go_repositories`

To build the application, `bazel build //...`

To run the application, `bazel run //cmd`

To run the application with configuration variables,
`ENABLE_GRPC=true bazel run //cmd`

### Note
1. If you face any issue while doing `go mod tidy`, run the following command
   `protoc \
   --go_out=:. \
   --go-grpc_out=:. \
   --go_opt=paths=source_relative \
   --go-grpc_opt=paths=source_relative \
   api/v1/proto/*.proto`

    Reason: Protobuf file will be generated and available in bazel-out folder which is recognizable by go build system. So when run go mod tidy it will throw few files are not found. We have to generate those file by using above command.
