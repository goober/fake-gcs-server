package fakestorage

import (
	"encoding/xml"
	"strings"
)

type InitiateMultipartUploadResult struct {
	XMLName  xml.Name `xml:"InitiateMultipartUploadResult"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	UploadId string   `xml:"UploadId"`
}

type CompleteMultipartUploadRequest struct {
	XMLName xml.Name `xml:"CompleteMultipartUpload"`
	Part    []Part
}

type CompleteMultipartUploadResult struct {
	XMLName  xml.Name `xml:"CompleteMultipartUploadResult"`
	Location string   `xml:"Location"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	ETag     ETag
}

type Part struct {
	XMLName    xml.Name `xml:"Part"`
	PartNumber int      `xml:"PartNumber"`
	ETag       ETag
}

type ListBucketResult struct {
	XMLName        xml.Name   `xml:"ListBucketResult"`
	Name           string     `xml:"Name"`
	CommonPrefixes []Prefix   `xml:"CommonPrefixes>Prefix"`
	Delimiter      string     `xml:"Delimiter"`
	MaxKeys        int64      `xml:"MaxKeys"`
	Prefix         string     `xml:"Prefix"`
	Marker         string     `xml:"Marker"`
	NextMarker     string     `xml:"NextMarker,omitempty"`
	StartAfter     string     `xml:"StartAfter"`
	IsTruncated    bool       `xml:"IsTruncated"`
	KeyCount       int        `xml:"KeyCount"`
	Contents       []Contents `xml:"Contents"`
}

type Contents struct {
	XMLName      xml.Name `xml:"Contents"`
	Key          string   `xml:"Key"`
	Generation   int64    `xml:"Generation"`
	LastModified string   `xml:"LastModified"`
	ETag         ETag
	Size         int64 `xml:"Size"`
}

type Prefix struct {
	Value string `xml:",innerxml"`
}

type ETag struct {
	Value string `xml:",innerxml"`
}

func (e *ETag) Equals(etag string) bool {
	trim := func(s string) string {
		return strings.TrimPrefix(strings.TrimSuffix(s, "\""), "\"")
	}
	return trim(e.Value) == trim(etag)
}
