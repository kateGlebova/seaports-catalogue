#!/bin/bash

SCRIPTDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

COMPOSE_FILE=docker-compose.yml

cd $SCRIPTDIR/../deployments

docker-compose -f $COMPOSE_FILE stop
docker-compose -f $COMPOSE_FILE rm -f
docker-compose -f $COMPOSE_FILE up --build

retcode=$?

docker-compose -f $COMPOSE_FILE stop
docker-compose -f $COMPOSE_FILE rm -f

exit $retcode
