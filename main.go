package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := context.Background()

	headers := metadata.Pairs(
		"authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"),
		"developer-token", os.Getenv("DEVELOPER_TOKEN"),
		"login-customer-id", os.Getenv("CUSTOMER_ID"),
	)
	ctx = metadata.NewOutgoingContext(ctx, headers)

	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial("googleads.googleapis.com:443", cred)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	glicd := "11ppppp"
	customerId := "aaaaa"
	jobId:= int32(68)
	
	clickConversion := services.ClickConversion{
		Gclid: &glicd,
	}

	conversions:= []*services.ClickConversion{&clickConversion}

	uploadCallConversionsRequest :=  services.UploadClickConversionsRequest{
		CustomerId: customerId,
		Conversions: conversions,
		PartialFailure: true,
		ValidateOnly: false,
		DebugEnabled: true,
		JobId: &jobId,
	}

	conversionUploadServiceClient := services.NewConversionUploadServiceClient(conn)
	
	resp,err := conversionUploadServiceClient.UploadClickConversions(
		ctx,
		&uploadCallConversionsRequest,

	)

	fmt.Println(resp)
	fmt.Println(err)
}

	