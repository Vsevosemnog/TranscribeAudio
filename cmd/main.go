package main

import (
	"context"
	"fmt"

	//"io/ioutil"

	internalModels "transcribeaudio/internal/models"
	utils "transcribeaudio/internal/utils"

	"transcribeaudio/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
)

func main() {
	var (
		err       error
		configure internalModels.AppConfigure
	)

	ctx := context.Background()
	appConf := config.Load()

	endpointCustomResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central1" {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}
		if service == mediaconvert.ServiceID {
			switch region {
			case "eu-central-1":
				return aws.Endpoint{
					URL:           "https://6qbvwvyqc.mediaconvert.eu-central-1.amazonaws.com",
					SigningRegion: region,
				}, nil
			default:
				return aws.Endpoint{}, fmt.Errorf("unknown redion for service mediaconvert")
			}
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	configure.YandexCloud.S3Client = s3.NewFromConfig(aws.Config{
		Region:                      appConf.YC.Region,
		Credentials:                 credentials.NewStaticCredentialsProvider(appConf.YC.StaticKeyID, appConf.YC.StaticKeySecret, ""),
		EndpointResolverWithOptions: endpointCustomResolver,
	})

	yandexCloudServiceAccountKey, err := iamkey.ReadFromJSONBytes(appConf.YC.IamSaKey)
	if err != nil {
		fmt.Println(err)
	}

	yandexCloudServiceAccountCredentials, err := ycsdk.ServiceAccountKey(yandexCloudServiceAccountKey)
	if err != nil {
		fmt.Println(err)
	}

	configure.YandexCloud.SDK, err = ycsdk.Build(ctx, ycsdk.Config{
		Credentials: yandexCloudServiceAccountCredentials,
	})
	if err != nil {
		fmt.Println(err)
	}

	transcribeFilePath := "output/test"

	audioFileUrl := "https://storage.yandexcloud.net/<bucket>/<filePath>"
	err = utils.YandexSttRecognition(ctx, configure.YandexCloud.SDK, audioFileUrl, transcribeFilePath)
	if err != nil {
		fmt.Println(err)
	}
}
