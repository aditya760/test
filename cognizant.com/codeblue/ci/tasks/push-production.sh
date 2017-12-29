#!/usr/bin/env bash
set -e -x

pushd codeblue-backend
  cf api $CF_API
  cf login -u $CF_USERNAME -p $CF_PASSWORD -o $CF_ORG -s $CF_SPACE
  cf push codeblue-backend
popd
