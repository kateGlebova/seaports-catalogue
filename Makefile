vendor:
	@mkdir -p $@

deps: vendor
	GO111MODULE=on go mod vendor