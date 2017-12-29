#!/usr/bin/env bash
set -e -x

mkdir -p /gopath/src/cognizant.com
cp -r codeblue /gopath/src/cognizant.com/

pushd /gopath/src/cognizant.com/codeblue
  glide up
  go install
popd

mkdir -p /tmp/test-microservice
cp codeblue-backend/test/acceptance/arch-test.yml /tmp/test-microservice/

pushd /tmp/test-microservice
  codeblue api $TEST_BACKEND_URL
  codeblue init -a arch-test.yml -l nodejs
  npm test

  cf api $CF_API
  cf login -u $CF_USERNAME -p $CF_PASSWORD -o $CF_ORG -s $CF_SPACE
  cf push generated-test-app --no-start
  cf bind-service generated-test-app generated-app-db

  cf set-env people-tech-backend-test NPM_CONFIG_PRODUCTION false
  cf set-env people-tech-backend-test NODE_ENV 'test-acceptance'

  cf start generated-test-app

  npm run test-acceptance
popd
