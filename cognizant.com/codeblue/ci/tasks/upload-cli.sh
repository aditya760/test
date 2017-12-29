#!/usr/bin/env bash
set -e -x

mkdir -p /gopath/src/cognizant.com
cp -r codeblue /gopath/src/cognizant.com/

export version=$(cat version-semver/number)

pushd /gopath/src/cognizant.com/codeblue
  glide up

  echo $GCS_KEY > gcs_key.json
  gcloud auth activate-service-account --key-file gcs_key.json



  env GOOS=linux GOARCH=amd64 go build
  gsutil cp codeblue gs://codeblue-linux/
  gsutil cp codeblue gs://codeblue-linux/codeblue-linux-amd64-$version
  # gsutil cp codeblue gs://codeblue-cli/linux/

  env GOOS=darwin GOARCH=amd64 go build
  gsutil cp codeblue gs://codeblue-darwin/
  gsutil cp codeblue gs://codeblue-darwin/codeblue-darwin-amd64-$version

  env GOOS=windows GOARCH=amd64 go build
  gsutil cp codeblue.exe gs://codeblue-windows/
  gsutil cp codeblue.exe gs://codeblue-windows/codeblue-windows-amd64-$version.exe

popd
