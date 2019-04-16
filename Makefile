vendor:
	@mkdir -p $@

deps: vendor
	GO111MODULE=on go mod vendor

test: deps run-tests rm-test-svc clean

run-tests: run-test-svc
	./scripts/run-tests.sh

run-test-svc:
	./scripts/run-docker-compose-test.sh

rm-test-svc:
	./scripts/rm-docker-compose-test.sh