#!/bin/bash

# shellcheck disable=SC2046
docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q)
docker system prune -f
docker compose build --no-cache
docker compose up -d
