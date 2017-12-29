#!/usr/bin/env bash
set -e -x

cd codeblue-backend
cf api $CF_API
cf login -u $CF_USERNAME -p $CF_PASSWORD -o $CF_ORG -s $CF_SPACE
cf push codeblue-backend-test --no-start

cf set-env codeblue-backend-test NPM_CONFIG_PRODUCTION false
cf set-env codeblue-backend-test NODE_ENV 'test-acceptance'

cf start codeblue-backend-test
