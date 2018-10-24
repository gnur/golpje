#!/bin/bash
set -ex

docker build -t gcr.io/dekeijzer-xyz/golpje:localbuild .

docker run --env-file .env gcr.io/dekeijzer-xyz/golpje:localbuild
