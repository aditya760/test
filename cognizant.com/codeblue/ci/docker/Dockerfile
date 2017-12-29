FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y curl

RUN curl -O https://deb.nodesource.com/setup_8.x
RUN chmod +x setup_8.x
RUN ./setup_8.x
RUN apt-get install -y nodejs

RUN curl -O https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
RUN tar -xvf go1.9.2.linux-amd64.tar.gz
RUN mv go /usr/local/
RUN rm go1.9.2.linux-amd64.tar.gz

ENV PATH=/usr/local/go/bin:$PATH
RUN mkdir /gopath
ENV GOPATH=/gopath

RUN apt-get install -y git
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
ENV PATH=/gopath/bin:$PATH

RUN curl https://glide.sh/get | sh

RUN curl -L "https://cli.run.pivotal.io/stable?release=linux64-binary&source=github" | tar -zx
RUN mv cf /usr/local/bin

RUN apt-get install -y software-properties-common python-software-properties
RUN yes | add-apt-repository ppa:webupd8team/java
RUN apt-get update
RUN echo "oracle-java8-installer shared/accepted-oracle-license-v1-1 select true" | debconf-set-selections
RUN apt-get install -y oracle-java8-installer

ENV JAVA_HOME=/usr/lib/jvm/java-8-oracle/

RUN wget http://www-eu.apache.org/dist/maven/maven-3/3.3.9/binaries/apache-maven-3.3.9-bin.tar.gz
RUN tar -xvzf apache-maven-3.3.9-bin.tar.gz
RUN mv apache-maven-3.3.9 /usr/local/maven
RUN ln -s /usr/local/maven/bin/mvn /usr/local/bin/mvn

RUN curl https://sdk.cloud.google.com | bash
RUN mv /root/google-cloud-sdk /usr/local/
ENV PATH=/usr/local/google-cloud-sdk/bin:$PATH

RUN apt-get install -y build-essential

RUN wget https://services.gradle.org/distributions/gradle-3.0-bin.zip
RUN apt install -y unzip
RUN unzip gradle-3.0-bin.zip -d /usr/local
RUN ln -s /usr/local/gradle-3.0/bin/gradle /usr/local/bin/gradle
