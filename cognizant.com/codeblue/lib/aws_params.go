package lib
type AWSParams struct {
  AWS_ACCESS_KEY string
  AWS_SECRET_KEY string
  PCF_USER_NAME  string
  PCF_PASSWORD   string
  AWS_KEY_PAIR   string
  AWS_KEY_MATERIAL string
  BUCKET_NAME    string
  AWS_CERT_ARN   string
  AWS_DOMAIN     string
  AWS_ROUTE_ZONE_ID string
  AWS_REGION string
  AWS_AVAIL_ZONES map[int] string
}
