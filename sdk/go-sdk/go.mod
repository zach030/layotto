module mosn.io/layotto/sdk/go-sdk

go 1.14

require (
	github.com/golang/protobuf v1.5.2
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	mosn.io/layotto/spec v0.0.0-20210707123820-584778d048d3
)

replace mosn.io/layotto/spec v0.0.0-20210707123820-584778d048d3 => ../../spec
