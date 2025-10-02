#!/usr/bin/env sh

IMAGE_NAME="http-headers-app"
set -e

docker rm "$IMAGE_NAME"

set +e


# Сборка образа
docker build -t "$IMAGE_NAME" .

# Запуск контейнера с пробросом порта 8181
docker run --rm -p 8181:8181 --name "$IMAGE_NAME" "$IMAGE_NAME"