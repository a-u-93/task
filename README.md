# Currency Exchange Task

## Quickstart

```sh
dash quickstart
```

or

```sh
chmod +s quickstart
./quickstart
```

or (do not forget to copy without extratrailing rightmost spaces to craft
proper shell command, if you copy by hand)

```sh
docker compose down

docker compose build

MARIADB_ROOT_PASSWORD=exchange \
MARIADB_DATABASE=exchange \
MARIADB_ADDRESS=127.0.0.1:3306 \
MARIADB_USER=exchange \
MARIADB_PASSWORD=exchange \
MIDDLEWARE_ADDRESS=127.0.0.1:7777 \
UPSTREAM_API='https://api.nbrb.by/exrates/rates' \
  docker compose up -d
```

test

```sh
curl 127.0.0.1:7777/year/month/day | jq .
curl -s 127.0.0.1:7777/ | jq .
```
