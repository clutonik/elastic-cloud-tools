build_eceRolesToken:
	go build -o cmd/eceRolesToken runners/eceRolesToken.go

startECE:
	VBoxManage startvm "ece-director"
	VBoxManage startvm "ece-allocator-1"
	VBoxManage startvm "ece-proxy"

generate_go:
	protoc api/proto/v1/*.proto --go_out=plugins=grpc:.
