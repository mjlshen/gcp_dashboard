package gcloud

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectGenerator(t *testing.T) {
	tests := []struct {
		p []Project
	}{
		{
			p: []Project{
				{
					ProjectId: "abc",
					Number:    123,
				},
				{
					ProjectId: "def",
					Number:    456,
				},
			},
		},
	}

	for _, test := range tests {
		expected := test.p
		done := make(chan interface{})
		defer close(done)

		actualchannel := ProjectGenerator(done, test.p...)

		var actual []Project
		for project := range actualchannel {
			actual = append(actual, project)
		}

		assert.IsType(t, expected, actual)
		assert.Equal(t, expected, actual)
	}
}

func TestBucketGenerator(t *testing.T) {
	tests := []struct {
		buckets []Bucket
	}{
		{
			buckets: []Bucket{
				{
					ProjectId:         "abc",
					Name:              "abc",
					Location:          "abc",
					StorageClass:      "abc",
					VersioningEnabled: true,
				},
				{
					ProjectId:         "def",
					Name:              "def",
					Location:          "def",
					StorageClass:      "def",
					VersioningEnabled: false,
				},
			},
		},
	}

	for _, test := range tests {
		expected := test.buckets
		done := make(chan interface{})
		defer close(done)

		actualchannel := BucketGenerator(done, test.buckets...)

		var actual []Bucket
		for bucket := range actualchannel {
			actual = append(actual, bucket)
		}

		assert.IsType(t, expected, actual)
		assert.Equal(t, expected, actual)
	}
}

func TestBucketAssembler(t *testing.T) {
	tests := []struct {
		buckets []Bucket
	}{
		{
			buckets: []Bucket{
				{
					ProjectId:         "abc",
					Name:              "abc",
					Location:          "abc",
					StorageClass:      "abc",
					VersioningEnabled: true,
				},
				{
					ProjectId:         "def",
					Name:              "def",
					Location:          "def",
					StorageClass:      "def",
					VersioningEnabled: false,
				},
			},
		},
	}

	for _, test := range tests {
		expected := test.buckets
		done := make(chan interface{})
		defer close(done)

		fmt.Println(test.buckets)
		actualchannel := BucketGenerator(done, test.buckets...)
		actual := BucketAssembler(done, actualchannel)

		assert.IsType(t, expected, actual)
		assert.Equal(t, expected, actual)
	}
}
