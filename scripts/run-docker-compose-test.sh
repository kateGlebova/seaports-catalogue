#!/bin/bash

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
COMPOSE_FILE=docker-compose.test.yml

cd $SCRIPT_DIR/../deployments

docker-compose -f $COMPOSE_FILE stop
docker-compose -f $COMPOSE_FILE rm -f
docker-compose -f $COMPOSE_FILE up -d --build

