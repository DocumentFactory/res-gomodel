package s3client

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pnocera/res-gomodel/config"
)

type S3Client struct {
	session *session.Session
	region  string
	cfg     *config.Config
	svc     *s3.S3
}

func LogError(err error) error {
	if err == nil {
		return nil
	}
	log.Printf("Error: %s", err.Error())
	return err
}

func NewS3ClientFromSpec(endpoint string, accessKeyID string, accessKey string, region string, insecureTLS bool) (*S3Client, error) {
	cfg := config.New()
	if insecureTLS {
		endpoint = "http://" + endpoint
	} else {
		endpoint = "https://" + endpoint
	}

	session := getAwsSession(endpoint, accessKeyID, accessKey, region)
	svc := s3.New(session, &aws.Config{
		MaxRetries: aws.Int(30),
		Region:     aws.String(region),
	})

	return &S3Client{
		session: session,
		region:  region,
		cfg:     cfg,
		svc:     svc,
	}, nil
}

func NewS3Client(cfg *config.Config) (*S3Client, error) {
	endpoint := cfg.GetString("MINIO_ENDPOINT", "")
	accessKeyID := cfg.GetString("MINIO_ACCESSKEY_ID", "")
	accessKey := cfg.GetString("MINIO_ACCESSKEY", "")
	insecureTLS := cfg.GetBool("MINIO_INSECURE")
	region := cfg.GetString("MINIO_REGION", "")

	return NewS3ClientFromSpec(endpoint, accessKeyID, accessKey, region, insecureTLS)
}

func (c *S3Client) BucketExists(ctx context.Context, runid string) (bool, error) {

	_, err := c.svc.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: aws.String(runid),
	})
	if err != nil {
		return false, err
	}

	return true, nil

}

func (c *S3Client) GetOrCreateBucket(ctx context.Context, runid string) error {
	found, _ := c.BucketExists(ctx, runid)

	if !found {
		// create new bucket with default ACL in default region
		_, err := c.svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(runid),
		})
		if err != nil {
			log.Printf("Error creating bucket: %s", err.Error())
			return err
		}
		c.BucketExists(ctx, runid)

	}
	return nil

}

func (c *S3Client) DownloadFile(ctx context.Context, runid string, id string, filename string) (int64, error) {

	c.BucketExists(ctx, runid)

	localfoldername := path.Join("/img/", runid)
	localfilename := path.Join(localfoldername, filename)

	err := os.MkdirAll(localfoldername, os.ModePerm)

	if LogError(err) != nil {
		return 0, err
	}

	inputTemplate := s3manager.UploadInput{
		// ACL:          aws.String(s3.ObjectCannedACLPrivate),
		// StorageClass: aws.String(s3.StorageClassStandard),
		//ServerSideEncryption: "AES256",
		Bucket: aws.String(runid),
	}

	// if *cpSSEKMSKeyID != "" {
	// 	inputTemplate.ServerSideEncryption = cpSSEKMSKeyID
	// }

	copier, err := NewBucketCopier(false, "s3://"+runid+"/"+id, "file://"+"/img/"+runid+"/"+filename, 30, true, c.session, inputTemplate, "", false, false)

	if LogError(err) != nil {
		return 0, err
	}

	err = copier.copy()

	if LogError(err) != nil {
		return 0, err
	}

	fi, err := os.Stat(localfilename)

	if LogError(err) != nil {
		return 0, err
	}

	return fi.Size(), nil
}

func (c *S3Client) UploadFile(ctx context.Context, runid string, id string, contenttype string, fullfilename string) (int64, error) {

	fi, err := os.Stat(fullfilename)
	if LogError(err) != nil {
		return 0, err
	}

	err = c.GetOrCreateBucket(ctx, runid)
	if LogError(err) != nil {
		return 0, err
	}

	inputTemplate := s3manager.UploadInput{
		// ACL:          aws.String(s3.ObjectCannedACLPrivate),
		// StorageClass: aws.String(s3.StorageClassStandard),
		ContentType: aws.String(contenttype),
		Bucket:      aws.String(runid),
		//ServerSideEncryption: "AES256",
	}

	// if *cpSSEKMSKeyID != "" {
	// 	inputTemplate.ServerSideEncryption = cpSSEKMSKeyID
	// }

	copier, err := NewBucketCopier(false, "file://"+fullfilename, "s3://"+runid+"/"+id, 30, true, c.session, inputTemplate, "", false, false)

	if LogError(err) != nil {
		return 0, err
	}

	err = copier.copy()

	if LogError(err) != nil {
		return 0, err
	}

	return fi.Size(), nil

}

func (c *S3Client) Delete(ctx context.Context, runid string, id string) error {

	_, err := c.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(runid),
		Key:    aws.String(id),
	})

	if LogError(err) != nil {
		return err
	}

	return nil
}

func (c *S3Client) DeleteBucket(ctx context.Context, runid string) error {

	found, _ := c.BucketExists(ctx, runid)

	if !found {
		return nil
	}

	resp, err := c.svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(runid),
	})
	if LogError(err) != nil {
		return err
	}

	for _, obj := range resp.Contents {
		c.Delete(ctx, runid, *obj.Key)
	}

	_, err = c.svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(runid),
	})

	if LogError(err) != nil {
		return err
	}

	return nil
}

func (c *S3Client) ListBuckets(ctx context.Context) ([]*s3.Bucket, error) {
	resp, err := c.svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return resp.Buckets, nil
}

func (c *S3Client) ListObjects(ctx context.Context, runid string, prefix string, maxkeys int) ([]s3.Object, error) {

	objectsCh := make(chan s3.Object)

	errorchan := make(chan error)

	go func() {
		defer close(objectsCh)

		resp, err := c.svc.ListObjects(&s3.ListObjectsInput{
			Bucket: aws.String(runid),
			Prefix: aws.String(prefix),
		})
		if err != nil {
			errorchan <- err
			return
		}

		for _, obj := range resp.Contents {
			objectsCh <- s3.Object{
				Key:          obj.Key,
				ETag:         obj.ETag,
				LastModified: obj.LastModified,
				Size:         obj.Size,
				StorageClass: obj.StorageClass,
			}
		}
	}()

	var objects []s3.Object

	for object := range objectsCh {
		objects = append(objects, object)
	}

	return objects, nil

}
