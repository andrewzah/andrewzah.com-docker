#!/bin/sh

./docker-compose exec miniflux miniflux -migrate
./docker-compose exec miniflux miniflux -create-admin
