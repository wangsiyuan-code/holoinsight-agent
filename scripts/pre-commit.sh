#!/usr/bin/env bash
set -e

cd `dirname $0`/..

pre-commit run --all-files
