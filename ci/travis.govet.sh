#!/usr/bin/env bash

set -e

REPO_ROOT=`git rev-parse --show-toplevel`
cd $REPO_ROOT
DIRS="example"

for subdir in $DIRS; do
  pushd $subdir
  go vet
  popd
done
