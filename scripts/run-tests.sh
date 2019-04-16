#!/usr/bin/env bash

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

set -ex

mongoUrl=${TEST_MONGO_URL:=127.0.0.1:27017}
mongoDB=${TEST_MONGO_DB:=ports}

test_cmd() {
    MONGO_URL=$mongoUrl MONGO_DB=$mongoDb go test -v
}

for f in $(find . -type f -name '*_test.go' | sed -E 's|/[^/]+$||' |sort -u)
do
    pushd $SCRIPT_DIR/../$f
    test_cmd
    popd
done
