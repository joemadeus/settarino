catalog_proto:
	@ rm catalog/file_format.pb.go
	@ protoc --proto_path=catalog --go_out=catalog --go_opt=paths=source_relative catalog/file_format.proto
