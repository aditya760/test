

What is CodeBlue AutoPCF :
---------------------------
Command line option to provision Pivotal Cloud Foundry on any Iaas Platform through concourse pipeline

Pre-Requiste:

a) AWS Credentials for launching PCF Instances ( AWS Access Key and Secret Key). <br/>
          http://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html <br/>
b) Register Domain in AWS.<br/>
        http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-register.html <br/>
c) AWS certificate ARN for the above registered Domain.<br/>
      http://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request.html <br/>  
d) Fly should be have installed on the sytem.<br/>
      https://concourse.ci/fly-cli.html <br/>

##  For Direct Download of Command Line Interface:

| Windows (AMD64)       | Darwin (AMD64)           | Linux (AMD64) |
| ------------- |:-------------:| -----:|
|[codeblue.exe](https://s3-ap-southeast-1.amazonaws.com/codeblue-executables/codeblue-windows-64.exe) |[codeblue](https://s3-ap-southeast-1.amazonaws.com/codeblue-executables/codeblue-darwin-64) | [codeblue](https://s3-ap-southeast-1.amazonaws.com/codeblue-executables/codeblue-linux-64) |


How to Setup CodeBlue:
----------------------

a)  Install go on your machine.

      `https://golang.org/dl/`


b) Set GOPATH in your machine to your workspace folder

  export GOPATH=[WORKSPACE]

d) Create Cognizant.com folder under workspace and clone the codeblue autopcf through below command.

    `https://github.com/TheCognizantFoundry/TheCodeBlue.git`

e) Go to Code Blue folder and then install dependencies

    `brew install glide`

    `glide up`

f)  Go to root workspace folder and build code blue

    `go build cognizant.com/codeblue`

    `go install cognizant.com/codeblue`



How to  Run Codeblue command

    `./codeblue autopcf --help`

    Parameter for autopcf are:

    `-n`: Name of pipeline
    `-p`: Param file path
    `-t`: Iaas target

      `./codeblue autopcf -n autopcfdeploy -t AWS`


  List of Input required for setting up PCF:

        -  *AWS Access and Secret Key
        -  PWS Master User Id and Password  (User Id and password for Ops Manager, RDS Manager etc)
        -  AWS Key Pair Name
        -  AWS Bucket NAME
        -  *AWS Certificate ARN_
        -  *AWS Domain Name and Hosted Zone ID
        -  AWS Region NAME
        - AWS Availablity Zones

    * Note: - AWS Services should have setup before launching PCF.*


    How to Test Codeblue AutoPCF


    Install Mockgen to generate mock files.

      https://github.com/golang/mock

    Run the mockgen to generate stubs.
      
    $GOBIN/mockgen -source vendor/github.com/aws/aws-sdk-go/service/ec2/ec2iface/interface.go  -destination test/mocks/ec2_mock.go
