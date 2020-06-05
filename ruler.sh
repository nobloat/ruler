#!/bin/sh
REPO="https://github.com/nobloat/git.css"

[[ -d dir ]] ||git clone --depth 1 $1 src


cp $2 $1-analyze/


docker build -f Dockerfile.ruler --target rulder-dev -t ruler-dev .
docker build -f Dockerfile.ruler -t ruler-prod .
docker history --human=false --format="{{.CreatedBy}}\t{{.Size}}" ruler-dev > dev.layers
docker history --human=false --format="{{.CreatedBy}}\t{{.Size}}" ruler-prod > prod.layers

