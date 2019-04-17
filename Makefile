vendor:
	@mkdir -p $@

deps: vendor
	GO111MODULE=on go mod vendor

test: deps run-tests rm-test-svc

run-tests: run-test-svc
	./scripts/run-tests.sh

run-test-svc:
	./scripts/run-docker-compose-test.sh

rm-test-svc:
	./scripts/rm-docker-compose-test.sh

api:
	docker build -f ./build/api/Dockerfile -t glebova/client-api .

repository:
	docker build -f ./build/repository/Dockerfile -t glebova/port-domain-svc .

run:
	./scripts/run-docker-compose.sh