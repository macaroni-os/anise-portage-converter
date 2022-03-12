module github.com/Luet-lab/luet-portage-converter

go 1.16

replace github.com/mudler/luet => github.com/geaaru/luet v0.22.2-geaaru

replace github.com/containerd/containerd => github.com/containerd/containerd v1.3.1-0.20200227195959-4d242818bf55

replace github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe

replace github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305

replace github.com/opencontainers/runc => github.com/opencontainers/runc v1.0.0-rc9.0.20200221051241-688cf6d43cc4

replace github.com/docker/docker => github.com/Luet-lab/moby v17.12.0-ce-rc1.0.20200605210607-749178b8f80d+incompatible

require (
	github.com/MottainaiCI/lxd-compose v0.16.1
	github.com/MottainaiCI/simplestreams-builder v0.1.1 // indirect
	github.com/geaaru/pkgs-checker v0.12.3
	github.com/hashicorp/go-version v1.4.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-sqlite3 v1.14.12 // indirect
	github.com/mudler/luet v0.0.0-00010101000000-000000000000
	github.com/onsi/ginkgo/v2 v2.1.3
	github.com/onsi/gomega v1.18.1
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cobra v1.4.0
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0
)
