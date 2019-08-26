VERSION := 0.2

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
	docker build -f ./build/client-api/Dockerfile -t glebova/client-api:$(VERSION) .

repository:
	docker build -f ./build/domain-service/Dockerfile -t glebova/port-domain-service:$(VERSION) .

run:
	./scripts/run-docker-compose.sh