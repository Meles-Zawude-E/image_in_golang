package main

// import (
// 	"test/routes"

// 	"github.com/labstack/echo/v4/middleware"
// )

// func main() {
// 	e := routes.Routes()
// 	e.Use(middleware.CORS())
// 	e.Logger.Fatal(e.Start(":8080"))
// }

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )

    if err != nil {
        fmt.Println("Failed to connect to AWS", err)
        return
    }

    svc := s3.New(sess)

    resp, err := svc.ListBuckets(nil)
    if err != nil {
        fmt.Println("Error listing buckets", err)
        return
    }

    fmt.Println("Your Amazon S3 buckets are:")

    for _, b := range resp.Buckets {
        fmt.Println("* " + aws.StringValue(b.Name))
    }
}
