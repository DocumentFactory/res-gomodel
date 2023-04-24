package s3client

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const bigChanSize int64 = 1000000

type fileJob struct {
	path string
	info os.FileInfo
}

type MultiCopyInput struct {
	Input *s3.CopyObjectInput

	Bucket *string
}

// copyerror stores the error and the input object that generated the error.  Since it could be one of 4 types of input
// we use pointers rather than having a different error struct for each input type
type copyError struct {
	error    error
	download *s3.GetObjectInput
	upload   *s3manager.UploadInput
	copy     *s3.CopyObjectInput
	multi    *MultiCopyInput
}

type copyErrorList struct {
	errorList []copyError
}

func getAwsSession(endpoint string, accessKeyID string, accessKey string, region string) *session.Session {
	return session.Must(session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Endpoint:    aws.String(endpoint),
				Region:      aws.String(region),
				Credentials: credentials.NewStaticCredentials(accessKeyID, accessKey, ""),
			},
		}))
}

func (ce copyError) Error() string {
	var errString string

	if ce.download != nil {
		errString = *ce.download.Key + " "
	} else if ce.upload != nil {
		errString = *ce.upload.Key + " "
	} else if ce.copy != nil {
		errString = *ce.copy.Key + " "
	} else if ce.multi != nil {
		errString = *ce.multi.Input.Key + " "
	}

	return errString + ce.error.Error()
}

func (cel copyErrorList) Error() string {
	if len(cel.errorList) > 0 {
		out := make([]string, len(cel.errorList))
		for i, err := range cel.errorList {
			out[i] = err.Error()
		}
		return strings.Join(out, "\n")
	}
	return ""
}
