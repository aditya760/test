package lib

import (
    "fmt"
    "os"
   	"io/ioutil"
	  "strings"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
    "path/filepath"
    "bufio"
    "strconv"
)

func GetAWSparams() *AWSParams{

  	 var aws_access_key string
  	 var aws_secret_key  string

    fmt.Println("Enter AWS_ACCESS_KEY:::")
    fmt.Scanln(&aws_access_key)

    fmt.Println("Enter AWS_SECRET_KEY:::")
    fmt.Scanln(&aws_secret_key)


  	os.Setenv("AWS_ACCESS_KEY_ID", aws_access_key)
  	os.Setenv("AWS_SECRET_ACCESS_KEY", aws_secret_key)

    var userName string
    fmt.Println("Enter PCF Master User Name :::")
    fmt.Scanln(&userName)

    var password string
    fmt.Println("Enter PCF Master Password :::")
    fmt.Scanln(&password)


    var keyPairName string
    fmt.Println("Enter Aws Key Pair:::")
    fmt.Scanln(&keyPairName)

    var bucketName string
    fmt.Println("Enter Bucket Name :::")
    fmt.Scanln(&bucketName)

    var awsCertArn string
    fmt.Println("Enter AWS Certificate ARN::")
    fmt.Scanln(&awsCertArn)

    var domain string
    fmt.Println("Enter AWS Domain::")
    fmt.Scanln(&domain)

    var routeZoneId string
    fmt.Println("Enter Route Zone ID")
    fmt.Scanln(&routeZoneId)

    sess2, err := session.NewSession(&aws.Config{
      Region: aws.String("us-east-1")},
    )
    if err != nil {
      panic(err)
    }
    // Create an EC2 service client.
    svc2 := ec2.New(sess2)
    fmt.Println("Fetch AWS Regions .................")
    regionName := ChooseRegion(svc2)

    fmt.Println("Fetch AWS Availblity Zones for Region" + regionName + ".................")
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(regionName)},
    )
    // Create an EC2 service client.
    svc := ec2.New(sess)

    availablityZones := ChooseAvailablityZones(svc)
    keyMaterial := CreateAWSKeyPair(svc, keyPairName)
    awsParams:= AWSParams{
      AWS_ACCESS_KEY: aws_access_key,
      AWS_SECRET_KEY: aws_secret_key,
      PCF_USER_NAME: userName,
      PCF_PASSWORD: password,
      AWS_KEY_PAIR: keyPairName,
      AWS_KEY_MATERIAL: keyMaterial,
      BUCKET_NAME: bucketName,
      AWS_CERT_ARN: awsCertArn,
      AWS_DOMAIN: domain,
      AWS_ROUTE_ZONE_ID: routeZoneId,
      AWS_REGION: regionName,
      AWS_AVAIL_ZONES: availablityZones,
    }
    return &awsParams
}

func UpdateParams() {

  awsparamvar:= GetAWSparams()
  templateFilePath, _ := filepath.Abs("../src/cognizant.com/codeblue/config/params_template.yml")

  dir, err := os.Getwd()
  filePath := dir + "/" + "params.yml"
  if err != nil {
    panic(err)
  }
  if awsparamvar.BUCKET_NAME != "" {
    createBucket(awsparamvar.BUCKET_NAME)
  }
  WriteAWSParams(awsparamvar,templateFilePath,filePath)
}

func WriteAWSParams(awsparamvar *AWSParams,templateFilePath string,filePath string){

  read, err := ioutil.ReadFile(templateFilePath)
  if err != nil {
    panic(err)
  }

  fmt.Println("*******Genrating Bucket *********")
  params_yml := string(read)
  if awsparamvar.BUCKET_NAME != "" {
    params_yml = strings.Replace(string(read), "$BUCKET_NAME", awsparamvar.BUCKET_NAME, -1)
  }
  if awsparamvar.PCF_USER_NAME != "" {
    params_yml = strings.Replace(string(params_yml), "$user",  awsparamvar.PCF_USER_NAME, -1)
  }
  if awsparamvar.PCF_PASSWORD != "" {
    params_yml = strings.Replace(string(params_yml), "$password", awsparamvar.PCF_PASSWORD, -1)
  }
  if awsparamvar.AWS_ACCESS_KEY != "" {
    params_yml = strings.Replace(string(params_yml), "$AWS_ACCESS_KEY", awsparamvar.AWS_ACCESS_KEY, -1)
  }
  if awsparamvar.AWS_SECRET_KEY != "" {
    params_yml = strings.Replace(string(params_yml), "$AWS_SECRET_KEY", awsparamvar.AWS_SECRET_KEY, -1)
  }
  if awsparamvar. AWS_CERT_ARN != "" {
    params_yml = strings.Replace(string(params_yml), "$AWS_CERT_ARN", awsparamvar. AWS_CERT_ARN, -1)
  }
  if awsparamvar. AWS_DOMAIN != "" {
    params_yml = strings.Replace(string(params_yml), "$DOMAIN", awsparamvar. AWS_DOMAIN, -1)
    params_yml = strings.Replace(string(params_yml), "$OPS_DOMAIN", "opsman." + awsparamvar. AWS_DOMAIN, -1)

  }
  if awsparamvar.AWS_ROUTE_ZONE_ID != "" {
    params_yml = strings.Replace(string(params_yml), "$ROUTE_53_ZONE_ID", awsparamvar.AWS_ROUTE_ZONE_ID, -1)
  }
  if (awsparamvar.AWS_REGION != "") {
    params_yml = strings.Replace(string(params_yml), "$REGION", awsparamvar.AWS_REGION, -1)
  }
  if (awsparamvar.AWS_AVAIL_ZONES[1] != "") {
    params_yml = strings.Replace(string(params_yml), "$AZ1", awsparamvar.AWS_AVAIL_ZONES[1], -1)
  }
  if (awsparamvar.AWS_AVAIL_ZONES[2] != "") {
    params_yml = strings.Replace(string(params_yml), "$AZ2", awsparamvar.AWS_AVAIL_ZONES[2], -1)
  }
  if (awsparamvar.AWS_AVAIL_ZONES[3] != "") {
    params_yml = strings.Replace(string(params_yml), "$AZ3", awsparamvar.AWS_AVAIL_ZONES[3], -1)
  }



  if awsparamvar.AWS_KEY_PAIR != "" {
    //keyMaterial := CreateAWSKeyPair(svc, awsparamvar.AWS_KEY_PAIR)
    params_yml = strings.Replace(string(params_yml), "$KEY_NAME", awsparamvar.AWS_KEY_PAIR, -1)
    params_yml = strings.Replace(string(params_yml), "$PEM", awsparamvar.AWS_KEY_MATERIAL, -1)
  }

  err = ioutil.WriteFile(filePath, []byte(params_yml), 0644)
	if err != nil {
    panic(err)
  }
}

func GetListOfAvailabilityZones(iec2 ec2iface.EC2API) map[int] string {
  resultAvalZones, err := iec2.DescribeAvailabilityZones(nil)
  if err != nil {
    fmt.Println("Error", err)
  }
  //fmt.Println("Success", resultAvalZones.AvailabilityZones)

   avalZoneNumber := 0
    fmt.Printf("\n*************")
    fmt.Printf("\nList of Availabilty Zones")
    fmt.Printf("\n*************\n")
    var avalZones map[int]string
    avalZones = make(map[int]string)
    //var chosenavalZoneNumbers string

    for _, zone := range resultAvalZones.AvailabilityZones {
      avalZoneNumber = avalZoneNumber + 1
      fmt.Print(avalZoneNumber)
      fmt.Print(".  " + string(*zone.ZoneName) + "\n")
      avalZones[avalZoneNumber] = string(*zone.ZoneName)
    }
    return avalZones
}

func SelectAvailabilityZones(chosenavalZoneNumbers string,avalZones map[int] string) map[int] string{


  chosenZoneNumbers := strings.Split(chosenavalZoneNumbers, ",")
  var chosenAvalZones map[int]string
  chosenAvalZones = make(map[int]string)
  i := 0
  if len(chosenZoneNumbers) < 3 {
    var chosenavalZoneNumbers1 string
    fmt.Println("3 Availblity zones are required for setting up PCF Instance")
    fmt.Printf("Chose 3 Availabilty Zones::")
    fmt.Scanln(&chosenavalZoneNumbers1)
    var chosenAvalZones1 map[int]string
    chosenAvalZones1 = make(map[int]string)
    chosenAvalZones1 = SelectAvailabilityZones(chosenavalZoneNumbers1,avalZones)
    return chosenAvalZones1
  }
  for _, chosenZoneNumber := range chosenZoneNumbers {
      i = i + 1
      num, _ := strconv.Atoi(chosenZoneNumber)
      chosenAvalZones[i] = avalZones[num]
  }
  fmt.Println("\n Chosen Availablity Zones :: ")
  fmt.Println(chosenAvalZones[1])
  fmt.Println(chosenAvalZones[2])
  fmt.Println(chosenAvalZones[3])
  return chosenAvalZones
}
func ChooseAvailablityZones(iec2 ec2iface.EC2API) map[int] string{
  zones:= GetListOfAvailabilityZones(iec2)
  var chosenavalZoneNumbers string
  fmt.Printf("Chose 3 Availabilty Zones::")
  fmt.Scanln(&chosenavalZoneNumbers)
  return SelectAvailabilityZones(chosenavalZoneNumbers,zones)
}


func ChooseRegion(iec2 ec2iface.EC2API) string{
  regions:= GetListOfRegions(iec2)
  var chosenRegionNumber int
  fmt.Printf("Chose Region::")
  fmt.Scanln(&chosenRegionNumber)
  return SelectRegion(regions, chosenRegionNumber)
}
func GetListOfRegions(iec2 ec2iface.EC2API) map[int] string {
  resultRegions, err := iec2.DescribeRegions(nil)
  if err != nil {
      fmt.Println("Error", err)
  }
  var regions map[int]string
  regionNumber := 0
  fmt.Printf("\n*************")
  fmt.Printf("\nList of Regions")
  fmt.Printf("\n*************\n")
  regions = make(map[int]string)
  for _, region := range resultRegions.Regions {
    regionNumber = regionNumber + 1
    fmt.Print(regionNumber)
    fmt.Print(".  " + string(*region.RegionName) + "\n")
    regions[regionNumber] = string(*region.RegionName)
  }
  return regions
}

func SelectRegion(regions map[int] string, chosenRegionNumber int)string {
  fmt.Println("Chosen Region Name is ::" +  regions[chosenRegionNumber])
  return regions[chosenRegionNumber]
}
func createBucket(bucketName string) {

  sess, err := session.NewSession(&aws.Config{
      Region: aws.String("us-east-1"),
   })

    // Create S3 service client
    svc := s3.New(sess)

    // Create the S3 Bucket
    _, err = svc.CreateBucket(&s3.CreateBucketInput{
        Bucket: aws.String(bucketName),

    })

    if err != nil {
        exitErrorf("Unable to create bucket %q, %v", bucketName, err)
    }

    // Wait until bucket is created before finishing
    fmt.Printf("Waiting for bucket %q to be created...\n", bucketName)

    err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
        Bucket: aws.String(bucketName),
    })
    if err != nil {
        exitErrorf("Error occurred while waiting for bucket to be created, %v", bucketName)
    }

    fmt.Printf("Bucket %q successfully created\n", bucketName)
    _,err = svc.PutBucketVersioning(&s3.PutBucketVersioningInput{
        Bucket: aws.String(bucketName),
        VersioningConfiguration: &s3.VersioningConfiguration{
            Status: aws.String("Enabled"),
    }})
     if err != nil {
        exitErrorf("Unable to version the bucket %q, %v", bucketName, err)
    }

}



func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}


func Add(a int,b int) int {
  return a+b
}

func CreateAWSKeyPair(iec2 ec2iface.EC2API, pairName string) (string) {

    // Creates a new  key pair with the given name
    result, err := iec2.CreateKeyPair(&ec2.CreateKeyPairInput{
        KeyName: aws.String(pairName),
    })
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "InvalidKeyPair.Duplicate" {
            exitErrorf("Keypair %q already exists.", pairName)
        }
        exitErrorf("Unable to create key pair: %s, %v.", pairName, err)
    }

    fmt.Printf("Created key pair %q %s\n%s\n",
        *result.KeyName, *result.KeyFingerprint,
        *result.KeyMaterial)

    dir, err := os.Getwd()
    pemFilePath := dir + "/" + "pem.txt"
    if err != nil {
      panic(err)
    }

    err = ioutil.WriteFile(pemFilePath, []byte(*result.KeyMaterial), 0644)
    if err != nil {
      panic(err)
    }

    inFile, err := os.Open(pemFilePath)
    if err != nil {
       fmt.Println(err.Error() + `: ` + pemFilePath)
    } else {
       defer inFile.Close()
    }

    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines)
    var pemData string
    for scanner.Scan() {
      pemData = pemData + "    " +  scanner.Text() + "\n"
    }

    err = ioutil.WriteFile(pemFilePath, []byte(pemData), 0644)
    if err != nil {
      panic(err)
    }

    read, err := ioutil.ReadFile(pemFilePath)
    if err != nil {
      panic(err)
    }
    return string(read)
}

func TestStruct() string{
  type testStruct struct {
    BucketName string
    RegionName string
  }
  test := testStruct{}
  test.BucketName = "test"
  return test.BucketName
}
