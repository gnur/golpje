#!/bin/bash
set -ex

export GOLPJE_STORAGE="s3://"
export S3_HOST="s3.erwin.land"
export S3_BUCKET="golpje-dev"
export S3_ACCESS_KEY_ID="erwin"
export S3_SECRET_ACCESS_KEY="qJNLayisk1gxf6JRK"
#export GOLPJE_STORAGE="/Users/erwin/code/src/github.com/gnur/golpje/"

GOOS=linux GOARCH=arm go build -o "dist/golpje" ./*.go
scp dist/golpje pit:/tmp/
#ssh pit 'export GOLPJE_STORAGE="s3://"; export S3_HOST="s3.erwin.land"; export S3_BUCKET="golpje"; export S3_ACCESS_KEY_ID="erwin"; export S3_SECRET_ACCESS_KEY="qJNLayisk1gxf6JRK"; /tmp/golpje'
