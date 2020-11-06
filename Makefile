build_eceRolesToken:
	go build -o cmd/eceRolesToken runners/eceRolesToken.go

startECE:
	VBoxManage startvm "ece-director"
	VBoxManage startvm "ece-allocator-1"
	VBoxManage startvm "ece-proxy"
