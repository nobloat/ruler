#!/bin/sh

git clone $REPO


docker build -f Dockerfile.ruler --target rulder-dev -t ruler-dev .
docker build -f Dockerfile.ruler -t ruler-prod .
docker history --human=false --format="{{.CreatedBy}}\t{{.Size}}" ruler-dev
docker history --human=false --format="{{.CreatedBy}}\t{{.Size}}" ruler-prod

