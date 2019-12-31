package booter

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type OSS struct {
	Client        *oss.Client
	DefaultBucket *oss.Bucket
}

func BootOSS(baseConfig *BaseConfig) *OSS {
	client, err := oss.New(baseConfig.AliyunOSS.Endpoint,
		baseConfig.AliyunOptions.AccessKeyId,
		baseConfig.AliyunOptions.AccessSecret)
	if err != nil {
		panic(err.Error())
	}
	bucket, err := client.Bucket(baseConfig.AliyunOSS.DefaultBucket)
	if err != nil {
		panic(err.Error())
	}
	return &OSS{
		Client:        client,
		DefaultBucket: bucket,
	}
}
