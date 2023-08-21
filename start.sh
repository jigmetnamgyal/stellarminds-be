#!/bin/sh

export PATH=$PATH:/usr/local/go/bin

set -e

# Run db migration
./migrate

exec "$@"