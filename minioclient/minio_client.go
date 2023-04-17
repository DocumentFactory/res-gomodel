package minioclient

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"path"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pnocera/res-gomodel/config"
)

type MinioClient struct {
	client         *minio.Client
	cfg            *config.Config
	size_treshhold int
}

func LogError(err error) error {
	if err == nil {
		return nil
	}
	log.Printf("Error: %s", err.Error())
	return err
}

// isAccessDenied returns true if the error is caused by Access Denied.
func isAccessDenied(err error) bool {
	log.Printf("isAccessDenied(%T, %#v)", err, err)

	var e minio.ErrorResponse
	return errors.As(err, &e) && e.Code == "AccessDenied"
}

// IsNotExist returns true if the error is caused by a not existing file.
func IsNotExist(err error) bool {
	log.Printf("IsNotExist(%T, %#v)", err, err)

	var e minio.ErrorResponse
	return errors.As(err, &e) && e.Code == "NoSuchKey"
}

// Join combines path components with slashes.
func Join(p ...string) string {
	return path.Join(p...)
}

func NewMinioClient(cfg *config.Config) (*MinioClient, error) {
	endpoint := cfg.GetString("MINIO_ENDPOINT", "")
	accessKeyID := cfg.GetString("MINIO_ACCESSKEY_ID", "")
	accessKey := cfg.GetString("MINIO_ACCESSKEY", "")
	insecureTLS := cfg.GetBool("MINIO_INSECURE")
	region := cfg.GetString("MINIO_REGION", "")

	minio.MaxRetry = cfg.GetInt("MINIO_MAXRETRY", 10)

	creds := credentials.NewChainCredentials([]credentials.Provider{
		&credentials.EnvAWS{},
		&credentials.Static{
			Value: credentials.Value{
				AccessKeyID:     accessKeyID,
				SecretAccessKey: accessKey,
			},
		},
		&credentials.EnvMinio{},
		&credentials.FileAWSCredentials{},
		&credentials.FileMinioClient{},
		&credentials.IAM{
			Client: &http.Client{
				Transport: http.DefaultTransport,
			},
		},
	})

	c, err := creds.Get()
	if LogError(err) != nil {
		return nil, err
	}

	if c.SignerType == credentials.SignatureAnonymous {
		log.Printf("using anonymous access for %#v", endpoint)
	}

	tr, err := Transport(TransportOptions{
		InsecureTLS: insecureTLS,
	})

	if LogError(err) != nil {
		return nil, err
	}

	options := &minio.Options{
		Creds:        creds,
		Secure:       !insecureTLS,
		Region:       region,
		Transport:    tr,
		BucketLookup: minio.BucketLookupAuto,
	}

	client, err := minio.New(endpoint, options)
	if LogError(err) != nil {
		return nil, err
	}

	return &MinioClient{
		client:         client,
		cfg:            cfg,
		size_treshhold: cfg.GetInt("MINIO_SIZE_TRESHHOLD", 30*1024*1024),
	}, nil
}

func (c *MinioClient) BucketExists(ctx context.Context, runid string) (bool, error) {
	found, err := c.client.BucketExists(ctx, runid)
	if err != nil && isAccessDenied(err) {
		err = nil
		found = true
	}

	if LogError(err) != nil {
		return false, err
	}
	return found, nil
}

func (c *MinioClient) Upload(ctx context.Context, runid string, id string, contenttype string, reader io.ReadSeeker) (int64, error) {

	found, err := c.BucketExists(ctx, runid)

	if LogError(err) != nil {
		return 0, err
	}

	if !found {
		// create new bucket with default ACL in default region
		err = c.client.MakeBucket(ctx, runid, minio.MakeBucketOptions{})
		if LogError(err) != nil {
			return 0, err
		}
	}

	if contenttype == "" {
		contenttype = "application/octet-stream"
	}

	size, err := reader.Seek(0, io.SeekEnd)
	if LogError(err) != nil {
		return 0, err
	}

	opts := minio.PutObjectOptions{StorageClass: "STANDARD"}
	opts.ContentType = contenttype
	// the only option with the high-level api is to let the library handle the checksum computation
	opts.SendContentMd5 = true
	// only use multipart uploads for large files
	opts.PartSize = uint64(c.size_treshhold)

	if size > int64(c.size_treshhold) {

		opts.ConcurrentStreamParts = true

		opts.NumThreads = 5
	}

	_, err = reader.Seek(0, io.SeekStart)
	if LogError(err) != nil {
		return 0, err
	}

	info, err := c.client.PutObject(ctx, runid, id, reader, size, opts)

	if LogError(err) != nil {
		return 0, err
	}

	log.Printf("Uploaded %s of size %d to bucket %s with version %s", id, info.Size, runid, info.VersionID)

	return info.Size, nil
}

func (c *MinioClient) Download(ctx context.Context, runid string, id string) (io.ReadCloser, error) {

	found, err := c.BucketExists(ctx, runid)

	if LogError(err) != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("bucket not found")
	}

	obj, err := c.client.GetObject(ctx, runid, id, minio.GetObjectOptions{})

	if LogError(err) != nil {
		return nil, err
	}

	return obj, nil
}

func (c *MinioClient) Delete(ctx context.Context, runid string, id string) error {

	found, err := c.BucketExists(ctx, runid)

	if LogError(err) != nil {
		return err
	}

	if !found {
		return errors.New("bucket not found")
	}

	err = c.client.RemoveObject(ctx, runid, id, minio.RemoveObjectOptions{})

	if LogError(err) != nil {
		return err
	}

	return nil
}

func (c *MinioClient) DeleteBucket(ctx context.Context, runid string) error {

	found, err := c.BucketExists(ctx, runid)

	if LogError(err) != nil {
		return err
	}

	if !found {
		return nil
	}

	objectsCh := make(chan minio.ObjectInfo)

	go func() {
		defer close(objectsCh)

		// List all objects from a bucket-name with a matching prefix.
		for object := range c.client.ListObjects(ctx, runid, minio.ListObjectsOptions{Recursive: true}) {
			if object.Err != nil {
				return
			}

			objectsCh <- object
		}
	}()

	errChn := c.client.RemoveObjects(ctx, runid, objectsCh, minio.RemoveObjectsOptions{})

	for rmObjErr := range errChn {
		if LogError(rmObjErr.Err) != nil {
			return rmObjErr.Err
		}
	}

	err = c.client.RemoveBucket(ctx, runid)

	if LogError(err) != nil {
		return err
	}

	return nil
}

func (c *MinioClient) CopyTo(ctx context.Context, runid string, id string, accesskeyid string, accesskey string, region string, destbucket string, dest string, endpoint string) (int64, error) {

	reader, err := c.Download(ctx, runid, id)

	if LogError(err) != nil {
		return 0, err
	}

	defer reader.Close()

	// create new minio client
	creds := credentials.NewStaticV4(accesskeyid, accesskey, "")
	options := &minio.Options{
		Creds:        creds,
		Secure:       true,
		Region:       region,
		BucketLookup: minio.BucketLookupAuto,
	}

	client, err := minio.New(endpoint, options)
	if LogError(err) != nil {
		return 0, err
	}

	// create new bucket if it does not exist
	found, err := client.BucketExists(ctx, destbucket)

	if LogError(err) != nil {
		return 0, err
	}

	if !found {
		// create new bucket with default ACL in default region
		err = client.MakeBucket(ctx, destbucket, minio.MakeBucketOptions{})
		if LogError(err) != nil {
			return 0, err
		}
	}

	// upload file
	opts := minio.PutObjectOptions{StorageClass: "STANDARD"}
	opts.ContentType = "application/octet-stream"
	// the only option with the high-level api is to let the library handle the checksum computation
	opts.SendContentMd5 = true
	// only use multipart uploads for very large files
	opts.PartSize = 200 * 1024 * 1024

	opts.ConcurrentStreamParts = true

	opts.NumThreads = 4

	info, err := client.PutObject(ctx, destbucket, dest, reader, -1, opts)

	if LogError(err) != nil {
		return 0, err
	}

	log.Printf("Uploaded %s of size %d to bucket %s with version %s", dest, info.Size, destbucket, info.VersionID)

	return info.Size, nil

}
