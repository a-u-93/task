# Currency Exchange Task

## Quickstart

```sh
docker compose down

docker compose build

MARIADB_ROOT_PASSWORD=exchange \
MARIADB_DATABASE=exchange \
MARIADB_ADDRESS=127.0.0.1:3306 \
MARIADB_USER=exchange \
MARIADB_PASSWORD=exchange \
MIDDLEWARE_ADDRESS=127.0.0.1:7777 \
UPSTREAM_API='https://api.nbrb.by/exrates/rates?periodicity=0' \
  docker compose up -d
```

```sh
curl 127.0.0.1:7777/year/month/day | jq .
curl -s 127.0.0.1:7777/ | jq .
```
