package awstest

import (
    "testing"
    "fmt"
    "os"
    "strings"
    "io/ioutil"
    . "github.com/onsi/ginkgo"
    "github.com/aws/aws-sdk-go/aws"
    . "github.com/onsi/gomega"
    "cognizant.com/codeblue/lib"
    "github.com/golang/mock/gomock"
    "github.com/aws/aws-sdk-go/service/ec2"
    "path/filepath"
)

func TestCart(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, " Cart Suite")
}

type GinkgoTestReporter struct {}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
  Fail(fmt.Sprintf(format, args))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
  Fail(fmt.Sprintf(format, args))
}

var _ = Describe("generate aws params", func() {


    Context("initially", func() {
        var (
           t GinkgoTestReporter
           mockCtrl *gomock.Controller
        )
        BeforeEach(func() {
            os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIHSSC6AZ6VMNGFIQ")
            os.Setenv("AWS_SECRET_ACCESS_KEY", "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")
             mockCtrl = gomock.NewController(t)
         })
         AfterEach(func() {
            mockCtrl.Finish()
          })
        It("should return valid key pair", func() {
            keyPairOutput := ec2.CreateKeyPairOutput {
              KeyFingerprint: aws.String("abc"),
              KeyMaterial: aws.String("abc"),
              KeyName: aws.String("abc"),
            }
            ec2aws := NewMockEC2API(mockCtrl)
            ec2aws.EXPECT().CreateKeyPair(&ec2.CreateKeyPairInput{
                KeyName: aws.String("test"),
            }).
            Return(&keyPairOutput, nil)
            keyMaterial:= lib.CreateAWSKeyPair(ec2aws, "test")
            Expect(strings.TrimSpace(keyMaterial)).To(Equal("abc"))

        })
        It("should return list of regions", func() {
          regions:= []*ec2.Region{}
          region1:= new(ec2.Region)
          region1.Endpoint = aws.String("ec2.ap-south-1.amazonaws.com")
          region1.RegionName = aws.String("ap-south-1")

          regions = append(regions, region1)

          descriptionOutput:= ec2.DescribeRegionsOutput{
            Regions: regions,
          }

         ec2 := NewMockEC2API(mockCtrl)
         ec2.EXPECT().DescribeRegions(nil).
                      Return(&descriptionOutput,nil)
          resultRegions:= lib.GetListOfRegions(ec2)
          Expect(resultRegions[1]).To(Equal("ap-south-1"))
        })
        It("should return chosen region", func() {
          var regions map[int]string
          regions = make(map[int]string)
          regions[1] = "ap-south-1"
          regions[2] = "ap-south-2"
          chosenRegion:= lib.SelectRegion(regions, 2)
          Expect(chosenRegion).To(Equal("ap-south-2"))
        })
        It("should return struct", func() {
          test:= lib.TestStruct()
          Expect(test).To(Equal("test"))
        })
        It("test aws params", func() {

          var avalZones map[int]string
          avalZones = make(map[int]string)
          avalZones[1]="abczone0"
          avalZones[2]="abczone1"
          avalZones[3]="abczone2"

          awsParam := lib.AWSParams {
            AWS_ACCESS_KEY : "abcaccessk",
            AWS_SECRET_KEY: "abcsecretk",
            PCF_USER_NAME:  "abcusern",
            PCF_PASSWORD:   "abcpass",
            AWS_KEY_PAIR :  "abckey",
            AWS_KEY_MATERIAL: "abcmat",
            BUCKET_NAME:    "abcbucket",
            AWS_CERT_ARN:   "abccert",
            AWS_DOMAIN:     "abcdomain",
            AWS_ROUTE_ZONE_ID: "abcroute",
            AWS_REGION: "abckey",
            AWS_AVAIL_ZONES: avalZones,
          }
          templateFilePath, _ := filepath.Abs("../config/params_template.yml")

          dir1, err := os.Getwd()
          filePath := dir1 + "/params.yml"
          testDataFilePath:= dir1 + "/data/params.yml"
          if err != nil {
            panic(err)
          }
          lib.WriteAWSParams(&awsParam,templateFilePath,filePath)
          actual, err := ioutil.ReadFile(filePath)
          if err != nil {
            panic(err)
          }
          expected, err := ioutil.ReadFile(testDataFilePath)
          if err != nil {
            panic(err)
          }
          Expect(string(actual)).To(Equal(string(expected)))
        })

        It("should return availablity zones", func() {
          //regionName:= aws.String("us-east-1")
          azs:= []*ec2.AvailabilityZone{}
          az1:= new(ec2.AvailabilityZone)
          az1.ZoneName = aws.String("us-east-1a")
          azs = append(azs, az1)
          availablityZonesOutput:= ec2.DescribeAvailabilityZonesOutput{
            AvailabilityZones: azs,
          }

         ec2s := NewMockEC2API(mockCtrl)
         ec2s.EXPECT().DescribeAvailabilityZones(nil).
                      Return(&availablityZonesOutput,nil)
          test := lib.GetListOfAvailabilityZones(ec2s)
           Expect(test[1]).To(Equal("us-east-1a"))
        })

        It("should return chosen Zone", func() {
          var zones map[int]string
          zones = make(map[int]string)
          zones[1] = "us-south-1"
          zones[2] = "us-south-2"
          zones[3] = "us-south-3"
          zones[4] = "us-south-4"
          zones[5] = "us-south-5"
          chosenZone:= lib.SelectAvailabilityZones("1,3,5",zones)
          Expect(chosenZone[1]).To(Equal("us-south-1"))
        })
    })
})
