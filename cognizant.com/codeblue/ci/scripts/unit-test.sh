#/bin/bash
    set -e -x
    
    ls -ltr
    apt-get update
    curl https://glide.sh/get | sh
    ls -ltr
    go env
    pwd
    export GOPATH=$PWD
    mkdir -p src/cognizant.com/codeblue/
    mv $GOPATH/autopcftest/* $GOPATH/src/cognizant.com/codeblue/
    cd src/cognizant.com/codeblue
    glide up
    cd ..
    cd ..
    echo "START TESTING"
    go test cognizant.com/codeblue/test/
    echo "TEST DONE"