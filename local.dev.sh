#!/bin/bash

podman_spawn_psql() {
  local container_name="$1"
  local psql_password="$2"
  local psql_port="$3"
  local psql_dbname="$4"

  echo "Spawning local postgresql server using podman @\"postgresql://postgres:$psql_password@localhost:$psql_port/$psql_dbname\""

  ## Check if the volume mount path based on the container name does not already
  ## exist, if it does not, then create it the directory.
  local volume_path="$HOME/.local/share/podman/volumes/$container_name"
  if [ ! -d "$volume_path" ]; then
    mkdir -p "$volume_path"
  fi

  /bin/sh -c "podman run --rm --name $container_name -v pgdata:$volume_path -e POSTGRES_DB=$psql_dbname -e POSTGRES_PASSWORD=$psql_password -p $psql_port:5432 -d postgres"
}

podman_spawn_caddy() {
  local container_name="$1"
  local caddyfile_path="$2"
  local caddy_port="$3"

  local volume_path="$HOME/.local/share/podman/volumes/$container_name"

  /bin/sh -c "podman run --rm --network host --name $container_name -v caddy_data:$volume_path/data -v caddy_config:$volume_path/config -v $caddyfile_path:/etc/caddy/Caddyfile -e FRONTEND_URL=localhost:8080 -e BACKEND_URL=localhost:8081 -e PORT=9000 -d caddy"
}

## This is going to spawn a local caddy proxy.
podman_spawn_caddy "caddy.local.dev" "./Caddyfile" "9000"

## This script will be used to stand-up and deploy
## a local postgresql server using podman.
podman_spawn_psql "staircase.local.dev" "password" "5432" "staircase"
