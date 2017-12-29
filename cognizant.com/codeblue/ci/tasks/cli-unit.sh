#!/usr/bin/env bash
set -e -x

mkdir -p /gopath/src/cognizant.com
cp -r codeblue /gopath/src/cognizant.com/

cd /gopath/src/cognizant.com/codeblue
glide up
ginkgo -r
