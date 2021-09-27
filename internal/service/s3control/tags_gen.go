// Code generated by internal/generate/tags/main.go; DO NOT EDIT.

package s3control

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// []*SERVICE.Tag handling

// Tags returns s3control service tags.
func Tags(tags tftags.KeyValueTags) []*s3control.S3Tag {
	result := make([]*s3control.S3Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &s3control.S3Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from s3control service tags.
func KeyValueTags(tags []*s3control.S3Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(m)
}