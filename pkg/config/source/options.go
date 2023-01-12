/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/21 18:31:13
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/21 18:31:13
 */

package source

import (
	"io/fs"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Provider string

const (
	FS    Provider = "fs"
	Local Provider = "local"
	S3    Provider = "s3"
	MGDB  Provider = "mgdb"
)

type Option func(*Options)

type Options struct {
	Provider          Provider
	Dir               string
	Region            string
	Bucket            string
	ProjectName       string
	Timeout           time.Duration
	Client            *s3.Client
	FS                fs.ReadFileFS
	MongoDBURL        string
	MongoDBName       string
	MongoDBCollection string
}

// WithMongoDBURL set mongodb url
func WithMongoDBURL(url string) Option {
	return func(args *Options) {
		args.MongoDBURL = url
	}
}

// WithMongoDBName set mongodb name
func WithMongoDBName(name string) Option {
	return func(args *Options) {
		args.MongoDBName = name
	}
}

// WithMongoDBCollection set mongodb collection
func WithMongoDBCollection(collection string) Option {
	return func(args *Options) {
		args.MongoDBCollection = collection
	}
}

// WithProvider set provider
func WithProvider(provider Provider) Option {
	return func(args *Options) {
		args.Provider = provider
	}
}

// WithDir set dir
func WithDir(dir string) Option {
	return func(args *Options) {
		dir = strings.ReplaceAll(dir, "\\", "/")
		args.Dir = dir
	}
}

// WithProjectName set projectName
func WithProjectName(projectName string) Option {
	return func(args *Options) {
		args.ProjectName = projectName
	}
}

// WithRegion set s3 region
func WithRegion(region string) Option {
	return func(args *Options) {
		args.Region = region
	}
}

// WithBucket set s3 bucket
func WithBucket(bucket string) Option {
	return func(args *Options) {
		args.Bucket = bucket
	}
}

// WithTimeout set s3 client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(args *Options) {
		args.Timeout = timeout
	}
}

// WithClient set s3 client
func WithClient(client *s3.Client) Option {
	return func(args *Options) {
		args.Client = client
	}
}

// WithFrom set embed.FS
func WithFrom(fs fs.ReadFileFS) Option {
	return func(args *Options) {
		args.FS = fs
	}
}
