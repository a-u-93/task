from debian as buildtime
run --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
  apt update && \
  apt dist-upgrade -y
run rm -f /etc/apt/apt.conf.d/docker-clean
run --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
  apt install -y --no-install-recommends \
    curl jq ca-certificates ranger nmap
run \
  curl -s -L 'https://go.dev/dl/go1.23.3.linux-amd64.tar.gz' | \
  tar -xvz -C /opt --strip=1
add . /srv
workdir /srv
run --mount=type=cache,target=/srv/cache \
  GOROOT=/opt GOCACHE=/srv/cache GOOS=linux CGO_ENABLED=0 \
    /opt/bin/go mod tidy
run --mount=type=cache,target=/srv/cache \
  GOROOT=/opt GOCACHE=/srv/cache GOOS=linux CGO_ENABLED=0 \
    /opt/bin/go build -o /usr/bin/task /srv/cmd/api/main.go
run --mount=type=cache,target=/srv/cache \
  GOROOT=/opt GOCACHE=/srv/cache GOOS=linux CGO_ENABLED=0 \
    /opt/bin/go build -o /usr/bin/cron /srv/cmd/cron/main.go

# from scratch as runtime
# copy --from=buildtime /srv/task /task
# copy --from=buildtime /srv/*.gohtml /
# entrypoint ["/task"]
