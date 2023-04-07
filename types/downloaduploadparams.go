package types

type DownloadUploadParams struct {
	Bucket      string `json:"bucket"`      // The bucket name
	RelativeUrl string `json:"relativeurl"` // The path to the file or folder to download or upload
}
