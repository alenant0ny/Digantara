#!/bin/sh

host=$1
port=$2
shift 2

echo "Waiting for $host:$port to be available..."

while ! nc -z $host $port; do
  echo "Waiting for $host:$port to be available..."
  sleep 1
done

echo "$host:$port is up, running command: $@"
exec "$@"
