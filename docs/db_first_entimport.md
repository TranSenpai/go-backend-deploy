go get ariga.io/entimport/cmd/entimport

go run ariga.io/entimport/cmd/entimport -dsn "mysql://root:12345@tcp(localhost:3307)/go-backend"

lỗi:
# github.com/golang/protobuf/protoc-gen-go/descriptor
../../../../../../go/pkg/mod/github.com/golang/protobuf@v1.5.2/protoc-gen-go/descriptor/descriptor.pb.go:106:61: undefined: descriptorpb.Default_FileOptions_PhpGenericServices

fix lỗi cập nhật lên version mới:
go get github.com/golang/protobuf

go: upgraded github.com/golang/protobuf v1.5.2 => v1.5.4



