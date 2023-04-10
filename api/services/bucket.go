package services

import (
	"context"
	"letschat/infrastructure"

	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

// StorageBucketService -> handles the file upload/download functions
type StorageBucketService struct {
	logger infrastructure.Logger
	client *storage.Client
	env    infrastructure.Env
}

// NewStorageBucketService -> initilization for the StorageBucketService struct
func NewStorageBucketService(
	logger infrastructure.Logger,
	client *storage.Client,
	env infrastructure.Env,
) StorageBucketService {
	return StorageBucketService{
		logger: logger,
		client: client,
		env:    env,
	}
}

// UploadFile -> uploads the file to the cloud storage
func (s StorageBucketService) UploadFile(
	ctx context.Context,
	file multipart.File,
	fileName string,
) (string, error) {
	var bucketName = s.env.StorageBucketName

	if bucketName == "" {
		s.logger.Zap.Fatal("Please check your env file for StorageBucketName")
	}

	_, err := s.client.Bucket(bucketName).Attrs(ctx)

	if err == storage.ErrBucketNotExist {
		s.logger.Zap.Fatalf("provided bucket %v doesn't exists", bucketName)
	}

	if err != nil {
		s.logger.Zap.Fatalf("cloud bucket error: %v", err.Error())
	}
	wc := s.client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	wc.ContentType = "application/octet-stream"

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	u, err := url.ParseRequestURI("/" + bucketName + "/" + wc.Attrs().Name)

	if err != nil {
		return "", err
	}
	path := u.EscapedPath()
	path = strings.Replace(path, "/"+bucketName, "", 1)
	path = strings.Replace(path, "/", "", 1)
	return path, nil
}

// GetObjectSignedURL -> gets the signed url for the stored object
func (s StorageBucketService) GetObjectSignedURL(
	object string,
) (string, error) {
	var bucketName = s.env.StorageBucketName
	jsonKey, err := ioutil.ReadFile("serviceAccountKey.json")
	if err != nil {
		return "", nil
	}
	conf, err := google.JWTConfigFromJSON(jsonKey)

	if err != nil {
		return "", err
	}
	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(15 * time.Minute),
	}
	u, err := storage.SignedURL(bucketName, object, opts)
	if err != nil {
		return "", err
	}
	return u, nil
}

func (s StorageBucketService) GetObjectUnsignedURL(
	object string,
) (string, error) {
	ctx := context.Background()
	var bucketName = s.env.StorageBucketName
	bucket := s.client.Bucket(bucketName)
	obj := bucket.Object(object)
	u, err := obj.Attrs(ctx)
	if err != nil {
		return "", err
	}
	return u.MediaLink, nil
}
