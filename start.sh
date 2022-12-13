#!/bin/bash
docker-compose down --volumes
docker-compose up --build --no-start
if [ "$#" -lt 1 ]; then

    docker-compose run -e SaveMethod=$1 urlcutterapi
fi
docker-compose run -e SaveMethod=$1 urlcutterapi