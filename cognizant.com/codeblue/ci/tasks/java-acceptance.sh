#!/usr/bin/env bash
set -e -x

export TERM=${TERM:-dumb}

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
  codeblue init -a arch-test.yml
  gradle clean assemble

  cf api $CF_API
  cf login -u $CF_USERNAME -p $CF_PASSWORD -o $CF_ORG -s $CF_SPACE
  cf push generated-java-test-app --no-start
  # cf bind-service generated-test-app generated-app-db

  cf start generated-java-test-app

popd
