// Code generated by go-bindata.
// sources:
// data/ElasticsearchMetricType.json
// DO NOT EDIT!

package elasticsearch

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataElasticsearchmetrictypeJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9d\xcf\x6e\x24\xb7\x11\xc6\x5f\x65\x31\xc7\x40\x01\x36\xd7\x3d\x3b\x07\x1f\x92\x8b\x93\x1c\x62\x18\x04\xc5\xae\x99\xa1\xc5\x6e\x72\x49\xf6\x48\xb2\x91\x77\x0f\xa6\x67\xb4\x2b\x3b\x88\x86\xd5\x02\xdb\x2c\xee\x07\xf8\xb0\xc6\x72\xa4\xe9\xe6\x6f\xc9\xfa\xf3\x55\xd5\x8f\xbf\xee\x26\x3d\x52\x0a\xda\xd0\xee\xd3\x8f\xbf\xee\xfe\xa5\xdd\x4c\xbb\x4f\x3b\x3b\x65\x72\xbb\xbb\xdd\x77\x94\x4c\xb4\x21\x5b\x3f\xed\x3e\xed\x76\x77\xbb\xbf\xeb\xf1\xfc\xf7\xbb\xff\xdc\x7d\x5d\x4c\x4e\xa7\x6c\x4d\x22\x1d\xcd\xb1\xf4\x43\x93\x1f\xa8\x74\xed\x9f\xfe\x67\xe1\xf7\xdf\x7d\xf0\xfb\x0f\x7f\xfd\xe1\xc3\xd1\xa7\xfc\xf5\x53\xe7\xff\xfb\xf0\xfd\x77\xbf\xf9\x70\x3e\x46\xd2\x83\x0a\xde\x17\x3f\xd0\x81\x26\x8a\xd6\x94\x2e\xbf\xfc\x86\xf4\xe6\xf2\x9f\xee\x76\xe7\xb7\xa4\xf4\x70\xa2\x98\x6d\xa2\x41\x65\xbb\xfc\xdd\xc7\x8f\x1f\xff\xf2\xe7\xe5\xbf\x7f\x7c\xfc\xf8\x69\xf9\xef\xdf\xbb\xbb\xdd\x89\x62\x5a\x7e\xcc\xc7\xbb\x9d\xf1\xd3\xde\x1e\x76\x9f\xa6\xd9\xb9\xbb\xdd\xa0\xb3\x7e\xf9\x73\xd6\x87\xf4\xf2\xe7\x7f\x4e\x36\xab\xcb\xaf\x1d\x7e\xff\x3d\xce\xbf\x2b\x65\x3d\x86\xff\xf7\x0b\xcf\xcf\xd3\x3d\x0a\x76\x1a\xac\xa1\xb7\x37\xea\xd5\xf2\x40\xd1\x78\xa7\x73\xf1\x97\xfb\x3c\x53\xb4\x37\x7e\x3e\x40\x68\x00\x84\x15\x67\x02\xef\x49\x8c\x1f\x83\xa3\x4c\x03\x58\x68\x9d\x05\xe6\xa1\x10\x69\x1f\x29\x15\x3f\x46\xf6\x59\xbf\xfd\xa2\x00\x81\x3c\x08\x12\x1d\x46\x9a\x72\xf1\xfa\xbd\x7d\xa2\x41\xdd\xdb\xac\x12\x65\x35\xd2\xe8\xe3\xb3\xb2\x93\xba\x7f\xce\xb8\x2e\xda\xa7\x63\xc5\x75\xb1\xf7\xd1\x90\x1a\x29\x1e\x38\xd6\xc3\xfc\xf6\x62\xc0\xd0\x00\x0c\xc7\x9c\x03\xeb\xf4\x57\x3e\xd0\x04\x4b\xa0\xfd\x9d\xfd\xf9\x34\x96\xfe\x9e\x91\x8a\x97\x1e\x49\x07\x65\xfc\x38\xda\x9c\x69\xc0\xa9\xdf\x23\x0e\x25\x11\x80\xdf\x78\x07\xf3\x94\x01\x40\xeb\x00\xb0\x3d\x83\xcf\x33\xa5\xac\x8c\x36\xc7\xe2\x2f\x38\xda\x94\x14\x78\x10\xc1\xc3\x7d\x24\xfd\x40\x91\x0b\x44\xe9\x72\x7f\xa2\x78\x24\x0d\x43\xa1\x79\x10\x56\xf8\x03\xec\x58\x62\xa4\x9f\xc9\x20\x80\x24\x80\x06\x6e\xec\x20\xfb\x58\xfc\xc5\x92\xfd\x85\x60\x32\x8a\x41\x81\x7b\x43\xec\x2d\xb9\x61\x79\xdd\x85\x1f\x70\x76\xb4\x59\x81\x0a\x51\x54\xac\xb8\x2e\x9c\x4d\x99\x26\x8a\xc5\xeb\x75\x3c\xdc\x32\x35\xc0\x42\x03\x2c\x30\x9c\xca\xfb\x79\xbf\xa7\xb8\x60\x53\x7c\xa0\x0c\x36\x92\x29\xb6\x38\xe7\x84\x80\x04\xd8\xf9\xe2\x8a\xea\x10\x6e\xd8\x9b\x88\x61\x48\x63\x66\x5f\x7e\x74\x30\xcc\x90\xfc\x1c\x90\xa9\x68\x7e\xef\xeb\xc4\xb3\x59\x67\xca\xb3\x9f\xa7\x03\xae\xa3\x2e\xf1\xaa\x9c\x33\xb7\xd3\x40\x4f\xea\x31\xda\x4c\x11\x29\x73\xc0\xf1\x06\x1c\xfa\x09\x80\x48\x03\xa4\x62\xb2\xf5\xcc\x03\x40\x90\x02\x42\xe5\xac\x0a\xa5\x6c\x47\x9d\x69\x58\xa2\x66\xc0\xa1\x75\x1c\x56\x04\xcb\xb8\x92\x4c\xc8\xf5\x65\xa0\xb0\xe9\xc9\x80\x1b\x43\x0c\x17\x4c\xdb\xf2\x40\xe5\x48\x3c\xd9\x94\xd3\xb2\xa5\x67\x1e\x46\xeb\x9c\x05\x10\xcd\x03\xb1\xe2\xce\xb8\x9f\xdd\x03\x92\x2b\x9d\x71\xc0\x3c\x18\x16\x75\x76\xb9\x68\x6f\x8e\x91\xa6\xac\x06\x6f\x70\x24\x34\x8f\x42\x8e\x7a\x4a\xc1\xc7\xe2\xa3\x3f\x3e\xc1\x0c\x10\xb6\xc5\x2b\xaa\x32\x28\x9b\xa3\x4a\x47\x1d\x07\x95\xb2\x8e\xb7\xf4\x55\xf0\x1a\xe4\x61\x51\x27\xe3\x35\x22\xdd\x29\x61\xf3\xeb\x17\xfb\x6b\x93\xed\x09\x71\xa4\x1e\x49\x48\xf3\xe1\xa6\x9d\x0f\xb7\x40\x1e\x0a\x4c\xb7\x60\xef\x66\x66\x7d\x37\x02\x06\xc2\x88\xa8\xae\xd3\xcd\xd1\xde\x14\x53\x81\x84\x06\x48\x28\xb7\x16\x6f\x77\x72\x78\x6d\x25\x9c\xb4\x75\xfa\xde\xc1\x9f\x94\x83\x42\x7d\xdb\x11\x16\x83\x10\x14\xa4\x29\xe7\x02\xe9\x07\xc8\x1e\x44\x21\x56\x3f\xcf\x0d\x4f\x55\x06\x09\x4c\xf7\x64\x91\xbe\xd9\xf2\xb3\xe1\x22\x95\xdb\x6b\xeb\x60\x8f\x76\x07\xc3\xcd\x64\xd4\xeb\x04\x16\x42\x98\x1d\x12\xc0\x7b\x8a\xcf\x33\xc5\x67\x84\x2b\x84\x21\xe1\x19\x29\xed\x62\x6b\x74\xa9\xb6\x08\x14\x0d\xe1\x54\x68\x1f\x81\x15\xe6\xe2\xa3\x8e\x63\x79\x05\x31\xfa\x4d\x48\x41\x61\x13\x95\x0b\xb4\x10\x3d\x33\xc1\x8e\x67\xd3\xc9\x9a\xf3\x0a\xa0\xd0\x3c\x0a\x0d\xc4\xae\xbc\x2b\x96\xd5\xa0\xe6\x53\x14\x5c\x21\x7a\x43\xa9\x3c\x67\xf2\xe5\xa5\x61\x67\x1b\xdf\x59\x76\xb3\x43\xe3\x4f\x14\x9f\xb9\x76\x85\x4e\x2a\xf9\x39\x1a\xc4\x25\x9b\x27\xa2\x7e\x3e\x0c\x1d\xaf\xfb\x05\x01\x32\xaa\xde\x30\xa8\xd7\x0f\x79\x49\x5f\x22\x3c\x2d\x82\x82\x35\xba\xca\x49\x87\x74\xf4\xc5\x27\x02\xa4\xf6\x42\x58\x60\xda\x8c\x97\xe4\x03\xab\x3d\xf6\xb2\x18\x47\x43\x97\x38\xf0\x9e\x22\x99\xe8\x1d\x94\xb6\xd2\x98\x60\x18\x0d\xc6\xe9\x94\xca\xf1\xb9\x28\xaf\xe7\xc9\x79\x3d\xd0\x80\x33\x42\x06\x0f\xcc\x00\x92\x0f\x34\xa9\xbd\x75\xa4\x5e\xde\xa7\xbf\xa1\xda\xc6\x2e\xcb\xdb\xe5\x51\x3f\x61\x93\xc5\x6d\xf2\x26\xba\x35\xdc\xf7\xb2\xa0\xa8\x7d\xdf\xe3\xb6\x97\x44\xc3\x06\xc1\x02\xa8\x9c\x65\xa0\x50\xd5\x39\x5c\xcc\x44\xe3\xa7\x4c\x4f\x37\x7a\x88\x82\x85\x06\x58\x60\x5c\x12\x87\xe2\x64\x92\xf1\xce\x91\xb9\x69\x3a\xbe\xfa\x04\xab\xc4\xe6\xfa\xe3\xad\x9f\x60\x93\x08\xc3\xad\x7a\xfd\x37\x10\x68\x1d\x81\xeb\x0a\xec\x52\xd3\xbb\xf4\xee\x8e\x4e\x8c\xc9\x6a\xd0\x1e\xf4\x8a\x04\x77\x76\x16\xd2\x8d\x42\x58\x60\xde\xe3\xc6\x8f\xc1\xd1\xb2\xa6\xd4\xe9\x80\xe8\xbd\x63\x1e\xd6\xe5\x1b\x61\xdf\x49\x40\xa1\x7e\x35\xfd\xf5\x34\x41\x81\x54\x97\x30\xb0\xc7\x33\x43\xb5\x28\x04\x86\xa6\x06\x6e\x22\x6d\x21\x82\x99\xaa\x86\xc5\xc5\x63\x45\x0c\x51\x16\x12\x4c\x45\x83\x09\x33\x2f\xaf\x09\x14\xc4\xa0\xb0\x95\xea\x15\xf3\xaa\x3a\xa4\x81\x3b\xe8\x30\x53\x1c\xd5\xe9\x92\xd7\xc2\x90\x43\x69\x70\x54\x69\xdc\xb2\x8f\x84\x28\x95\x1c\x06\xd8\x93\x48\x32\x26\x91\x08\xdb\x62\xe6\x1d\xc0\x6b\xcc\x73\xad\x9c\x06\x05\x9d\x51\xc0\x28\x91\x1d\x6d\x4a\x76\x3a\xc0\x69\xec\x9b\x08\xae\x6d\x38\x78\xa3\x4e\xe7\x3f\xc3\x32\x14\x87\x46\xf5\x09\x04\xce\x8e\x36\xc3\x87\x14\x01\x43\x3d\xd1\xfc\x4b\xdb\x15\xc8\xe6\x25\xf1\xb0\xed\x64\x3b\x64\x36\x7b\x06\x83\x33\xe8\x16\x3d\x40\xa5\x80\xd0\x40\x93\xbf\x1f\xe6\x78\xb2\x27\x5f\xec\xc7\x62\x46\x85\x38\xca\xaa\xc6\x35\x30\x39\xed\x1b\x40\x82\xd7\x7c\x18\x62\x3b\x11\x0c\xd4\x1f\xad\x09\x49\xb6\x10\x14\x98\xc7\x01\x5b\x66\x87\xeb\x41\x16\x0f\x75\x86\x6f\x07\x9d\xdf\x7e\x54\xec\x7d\x03\x7b\xbf\x41\xb1\x3f\x8a\xb6\x7a\x25\x81\x7d\x33\x20\x6a\x25\x06\x87\x7a\x89\xd1\xf3\xfb\xfd\x7e\xfa\x1b\x8c\x83\x6f\x9c\x03\x7a\xb2\x29\x27\x14\x6b\x7d\xeb\x1c\x00\x80\x1e\x01\x58\xa1\xa0\xc6\xbc\x22\x31\x30\x6c\x9b\xec\x44\x8a\x4b\x0a\x17\x5b\x8c\xb9\x03\x05\xad\x53\x20\x30\xd1\x89\x91\x66\xa2\x08\xe3\x2a\xf1\x22\x7d\x9e\x19\xa9\x8c\xaf\x3a\x3c\x30\x21\x86\x89\xfa\x59\x4e\x95\x8f\xd1\xe7\xec\xae\xdb\x8b\x1c\x87\x1c\x38\x1a\xb8\x92\x18\x03\x36\x21\xbb\x01\x60\x7c\xc0\x58\xad\x51\x17\xc4\x60\xf5\x88\x62\x6c\x8d\xd7\xed\xa3\x21\xb5\xdc\x75\x48\xd2\x74\x07\x44\xe5\x3e\x39\xa3\x0e\xa1\x3c\x4a\x83\xc3\xa4\x57\x76\x5a\xec\xe7\x8d\xaa\x18\x11\x94\xad\xb9\xb2\x58\xad\xbc\x21\x3a\xec\x17\x04\xbe\xed\x82\x4c\x81\x14\x1e\x36\x88\xd6\x0c\xde\xe0\x60\x68\x1e\x84\x3a\xea\xd3\x14\x2c\xf2\xc8\xed\x6f\xfe\x8a\x5b\x61\x19\x31\x57\xba\x18\x83\xa6\x64\x70\xc0\x6d\xba\xc1\x19\x1a\xf1\x12\xb4\x47\xcc\x5e\x18\x14\xd5\xdb\x6d\x50\xca\x76\xd4\x99\x06\xb4\xdc\x10\x01\xc4\x8a\xdb\x82\xa3\x43\x84\x2b\x09\x0c\x60\x32\x48\xa1\xa0\x81\x84\x1b\x23\xa3\x8b\x64\xae\x24\xb6\x98\xe6\xe8\xd2\x2d\xd4\xf9\xe2\x30\xb6\x0f\x14\x75\x86\xd2\x59\x02\x0a\x2b\x2e\x9b\x51\x4f\xfa\x40\xe3\x2d\xa9\x2a\x4c\x0f\x79\x34\x6c\x31\xbd\x0a\x32\x67\x21\x30\x6c\x39\xf9\x12\x27\x44\xbf\x50\xf0\x9e\x04\x43\xac\x84\x90\xb0\xc1\x90\x09\xf4\x10\xee\x9d\x8a\x15\xed\xd9\xc0\x40\xeb\x0c\xd4\x49\x79\x62\xc4\x88\x28\x08\x2a\x5f\x0f\xb8\x18\x84\xf1\xb0\xa5\x37\x01\x99\x94\x14\x2a\xea\x77\x5d\x18\x6d\x4a\x50\xd3\x76\x49\x03\xef\x29\x2e\xec\xc0\x84\x14\x41\x42\xdd\x1c\x28\x4a\x80\xc4\x80\xc0\x3c\x12\x16\xf5\x9c\x2d\xaf\xb8\x58\xd6\xe3\x50\x90\xc1\xc2\x8a\x43\x81\x17\x5f\x80\x36\xa2\x5b\x10\x56\x14\x5b\x80\x06\x19\x34\x70\x75\x94\x41\xdf\xcc\x3f\xbe\xce\x47\x60\x66\x99\x18\x12\xaa\xfa\x0f\x97\x68\x04\x4c\x05\x11\x24\x54\x4f\x53\xe2\x72\xe8\x17\x04\xde\x5c\x10\xf4\x7f\x17\x01\x42\x5d\x2b\xc1\x9f\x28\x1e\x49\x23\xa0\xd0\x3c\x07\x95\x03\x0a\xa8\xc6\x12\xca\x85\x2f\x37\x1b\x1f\x75\x28\x5d\x8b\xa6\x31\xa2\x20\x60\x14\x5d\xcc\x01\xff\xc2\x65\x6d\xee\xe5\x63\x9c\x1c\x81\x75\x28\x63\x90\xb1\xb7\x95\xd5\x26\xd7\x2d\x53\xa3\x0e\x90\x24\x76\xce\xc6\x0a\x8d\xc1\xd1\x66\x48\x0c\x64\xc0\xb0\x4d\xea\x00\x22\x76\x21\x38\x54\xae\x84\x44\x8b\x75\x51\x34\xd4\x11\x2e\x5f\xba\x75\x81\x02\x29\x14\xac\xb8\x22\x9c\x4d\x99\xa6\x72\xa1\x01\x62\xc6\x22\x48\xa8\x3d\x49\xfa\xea\x4a\xe0\x96\xe8\x99\x8a\x48\xc6\x9f\x28\x3e\x23\x82\xdc\x39\x17\x75\x3a\xb6\x4c\x7e\x52\x47\xd2\x41\x19\x3f\x8e\x36\x67\xc4\x93\x05\x21\xb1\x49\x77\x0d\x38\x9b\xfd\xd2\xf0\x8e\xa9\x92\x50\xa8\x74\x4b\x05\xdb\xd0\x84\xbf\x21\x02\x85\xba\x1a\x15\xb4\x03\x95\x45\xc3\x26\xc6\x03\x4a\x65\x3b\xc6\x81\xd9\x5d\x03\x3d\x98\x64\x90\x50\x59\xc2\x36\x90\xa3\x4c\xe8\xd3\xd6\x29\x0e\xd7\x59\xc1\xbc\x0c\x37\x9d\xec\x32\x97\x0a\xa7\x43\xf3\x38\x54\x2f\x7d\x80\xc9\x20\x05\x85\x46\x66\xe0\x31\xfa\x48\x63\x02\x9e\x34\xc6\xea\x17\xd8\x20\x90\xd5\x2d\x09\x3c\x07\x05\x81\xee\x7e\x49\x88\xb4\x8f\x54\x3e\x17\x13\x01\x4d\x11\x20\x30\x5d\x13\x46\x0b\x9f\xd1\xa6\x64\xa7\x03\xea\xb1\x65\x80\xb0\x45\x92\x03\xf1\xab\x7e\x61\xe0\x07\xb6\x71\x43\xf4\xca\x02\xbb\xed\x1f\x38\x68\x9f\x83\xf2\xf2\x5b\x86\x76\x0a\xd5\xb7\xa2\x18\xd8\xe4\x5e\x40\x34\x41\x06\x0c\xb5\xb5\xd7\x70\x1c\x44\x60\x50\xe5\x5e\x40\x45\x8e\x2c\x08\x1a\x98\x85\xf9\xec\xe7\xf2\x0c\x3a\xa6\x61\x4a\xa2\xab\x6e\xaf\x69\xf8\x1f\x22\x20\x60\x9a\x1b\x4b\x35\x78\x79\xd9\xf0\x72\xdf\xa4\xec\x43\xb8\x6e\x2e\xea\x7a\xba\x45\x83\x2b\xb7\x9a\xbc\x0f\x6a\x0e\x83\xce\x84\x68\xb6\x0c\x22\xb6\x9c\x65\x03\x7f\x55\x06\x13\xe5\x8e\x8a\xf3\x7a\x50\xfa\x44\x51\xdf\x68\x28\x82\x7d\x6d\x60\x5f\xb7\x11\xdb\xc2\x26\xe8\x9a\x8a\xf7\x74\x06\x00\x0c\xad\xc3\x50\xde\x36\xe6\xb6\x7d\xf7\xfa\xc7\x62\xe0\x65\x9f\x14\x70\x9a\x07\xe9\x93\xb6\x4e\xdf\x3b\x90\x20\x87\x84\xca\x8d\x20\x90\xd5\x14\x45\xc3\x0a\x67\x71\x31\x22\x11\x60\xec\x0a\x03\xc6\xa1\x70\x3f\xef\xf7\x14\x15\x2b\x3f\x31\xea\x10\xca\x7b\x43\x5c\x02\x92\x46\x07\x6d\x6c\x46\x2f\xdb\x2e\x29\xaa\x95\x09\xfb\x61\x8e\x27\x7b\xf2\xc5\xaa\x7d\x24\xc3\x24\x01\x56\xb9\x91\xf6\xe4\xe3\x98\xd0\x42\x5b\x1a\x15\x95\x2f\xaf\xc1\x46\x32\xc5\x19\x55\x5c\x5e\x42\x29\xda\x64\x96\x2f\xda\x16\x08\xa1\x61\x8d\xf2\x7b\xd2\x21\x1d\x3d\x3a\xe3\xf5\xc6\x02\xd7\xea\xe0\x15\x10\x23\xab\x2e\x02\x82\xba\x3a\x2c\x74\xad\xe8\x98\x03\x76\x7e\x0d\xaa\x8a\x6e\x59\xb8\x9f\xdd\x43\xf1\xcd\x80\x92\x51\x19\x1c\x9c\x1f\x10\xbb\xd4\xfa\x2e\xad\xd1\xc5\xb9\x19\x5d\x1f\x3a\xc3\xa0\x6e\x1b\x5b\x67\x47\x9b\x31\x33\x43\x16\x12\x1b\xcc\xd4\x31\x7e\x0c\x8e\x60\xe1\xf7\x49\x03\x37\xf4\x03\x77\x4f\x0a\x0c\x75\x6f\x8b\x1c\xed\xcd\xc4\x38\x30\x68\x00\x03\x66\x08\x70\xf0\xa6\x7c\x2d\xe1\x5e\xe8\x91\x01\x6e\x18\x18\x1a\x7b\x51\x34\x34\xa0\x75\x61\x74\xae\x0d\xa4\x1f\x20\xcd\x94\x45\xd8\xb6\x03\x99\x90\x8d\x14\x82\x45\x95\x9e\x23\x4b\x35\x47\xa0\x68\x20\x4e\x10\x80\x00\xa3\xa4\xe7\xcb\xfb\xc2\xa6\x36\xbe\xa9\x1b\xf4\x78\xd0\x73\xf6\xea\xcb\x04\x57\x58\x02\x52\xd0\xa8\x5d\xa4\x81\xe0\xa4\x18\x14\x98\xa7\xc4\xde\x92\x1b\x38\x15\x7d\x18\xfa\xfd\x0d\x50\xc1\x1d\xfa\x7d\x55\xad\x2a\x9d\x54\x3e\xbb\x09\x30\x11\xfb\x23\x62\xc5\xd8\x2d\x9c\x15\x12\xc9\x58\x25\x64\x9c\x28\x5a\xc6\x0c\x26\x98\x13\x42\x60\xa8\x13\xc5\x44\x21\xb8\x3c\x12\xea\x0d\x43\x41\xcd\x4b\x9f\x08\x7c\x9e\x29\x3e\xf3\x2c\x86\x6b\x71\x1c\x66\xe7\x49\xc0\xa1\xb2\xba\x19\x56\x82\x1c\x12\xa2\x9e\x52\xf0\xb1\xf8\xc8\x4f\x14\x4f\x14\x95\x0f\x34\x61\x77\x5b\xdf\x5d\x86\x0d\x88\x74\x82\x98\x5d\xad\xdc\x2a\x01\x77\xb8\x08\x0a\xd6\x68\x08\x7c\x34\xa4\x96\xf4\x52\xf1\xb1\x80\x42\x25\x19\x38\x94\x6b\x07\xd2\xa3\x7e\xfb\x90\xff\xbd\x78\x00\x9e\xbe\x14\x08\x6a\xd7\x41\xc1\xb0\x17\x83\x02\x57\xdd\x7e\x4d\x13\x94\x2e\xff\x5a\x0c\x05\x14\x5a\x47\x61\x8b\x84\x11\x28\x68\x9d\x82\x10\xbd\xa1\x54\xee\x05\x84\xb9\x74\x29\xc4\x85\x42\x10\x58\x61\x1e\x3c\xea\x38\x96\x97\x42\xa2\x4e\x5a\x04\x07\x97\x8f\x15\x9f\x04\xe7\x7b\x40\xd1\xc9\x9a\xf3\x3a\x78\x01\xcd\x6f\x2f\xd7\xf4\x63\x2b\xc9\xfc\x89\xe2\x91\x34\xbc\x80\xe6\x51\xa8\x9c\xe8\x41\x37\xa3\x6e\x31\x88\xb4\x8f\x54\x1e\x19\x40\x8d\x91\x10\x14\xd8\xc9\x03\xce\x53\x5c\x75\x03\x50\x85\xc8\x60\xa1\xf2\xed\x00\x6f\xa0\x57\x0a\x98\x05\xf0\x08\x1b\x8b\x81\x61\x8b\x71\xc4\x10\x95\x4b\x22\x82\x19\x37\x64\x68\x89\x2f\x38\x9c\x6c\xcc\xb3\x76\x20\x42\x0c\x11\xec\x7c\x02\xcb\x93\xb8\x50\x81\x9e\x29\x5d\x33\xc1\xee\x9b\x0c\x9f\x42\x08\x08\xb5\x0b\x9a\xd1\x5f\x4f\x0a\x09\x0d\xb4\x51\xe2\x8e\x0c\x43\x2f\x25\x79\x98\xad\x71\x5f\x59\x4f\x02\x09\xa4\x10\x12\xaa\xe7\xbd\x28\x65\x3b\xea\x4c\x03\x7c\x58\x59\x64\x6c\x10\xd5\x80\xc3\x22\x0b\x89\x15\xd7\xc6\xa8\x27\x7d\xa0\x91\xd1\x04\x18\x81\x4f\x31\x3c\x54\xcd\x8b\x5d\x9a\x37\xc2\x87\x15\xc2\x42\xfd\x8c\x08\xbc\x58\x29\x2c\xd4\xa9\xab\xb9\xd8\x0c\x30\x20\xa5\x50\x80\xea\x58\x31\x5b\x55\xb9\x3a\x36\x65\x1f\x69\x50\x8b\xef\x88\x81\xe2\xe2\xe8\xd8\x60\x02\x0c\x8c\x7e\x31\x34\xd4\x1e\x0c\x9e\xbe\xb4\xe4\x05\x0c\xcd\xc3\x50\xbb\x78\x16\x36\xbf\x14\x12\xea\x4e\x86\x42\x14\x59\x28\x16\x8c\x84\xe6\x81\xd1\x51\xd3\x39\x32\xd9\x97\xd3\xc6\x98\x0d\x73\xfd\xe1\xd6\x4f\x88\x4f\x0b\x83\xad\xb2\x69\x72\x19\x4d\xa6\x96\x38\x04\x60\x68\x1d\x86\xfa\x8d\x7d\x61\x9c\x74\xcc\x02\x77\xa2\x2d\xca\xfa\xba\x45\x81\xe5\xb1\xa0\xa8\x4f\x0a\x08\xdc\x2e\x9e\xf9\x09\xad\x7a\x65\x6c\x6d\xfd\xce\xcd\x98\xf5\x20\x91\x8b\xea\xbd\x5c\x10\xc6\x16\xc3\x42\xe5\x5e\x6f\x18\x65\x2e\x84\x83\xda\x23\x3f\xd0\xe8\x4b\x08\x08\xef\x9e\x2f\xec\x63\xf1\x17\x85\x93\xd0\x2f\x14\xac\xca\x2c\xd4\x49\xf4\x0b\xc2\x3b\xa6\x8f\xa3\x0b\x88\x08\x28\x8e\x39\x17\xab\x17\x5f\x86\x46\x62\x08\x88\x84\x9d\xad\x7d\xee\xc3\x00\xe8\x17\x04\x6e\xe6\x00\x69\xa4\x8e\x61\x60\xe5\x0e\x90\x43\xea\x96\x03\x66\x43\x17\x9c\x09\x1d\xb3\xb0\x3a\x6c\x80\xf0\xb2\x18\x2c\xca\x4b\xa0\x50\x26\x23\x66\x53\xd9\x29\xe3\x88\x94\xb1\x90\xad\xe5\x56\x40\x31\x9b\x79\xa2\x8a\x59\x06\x06\x7b\x5e\x07\x0b\xd6\x62\x08\x05\xde\x8b\x81\x71\x73\xca\xe5\x8e\xf5\xf9\x59\x06\xe5\x67\xfc\xbb\xdb\xe8\x85\x5f\xfc\x57\x15\xa2\x1d\x75\x7c\xbe\x58\xb8\xa0\x7d\xa3\x97\x6f\x27\x9b\xad\x76\xf6\x17\x3b\x1d\xf0\xea\x37\x7d\xf5\xd3\x3c\xde\x53\x54\x7e\xaf\x02\x4d\xc3\xf9\xfd\x67\x9d\x1e\xf0\xfa\x37\x7f\xfd\x76\x52\x7b\x67\x0f\xc7\xac\x16\x1f\x1b\x1b\xb0\xf5\x06\x9c\x0d\x35\x70\xbf\xed\x75\x8b\xb3\x7e\xcb\x97\x3e\x90\xd3\xcf\x34\xa8\x79\xd2\x29\xd9\xc3\x44\x03\x36\x60\xd3\x0d\xc0\x8b\xff\xa3\xdc\x29\x9d\x1e\xd4\xa8\x9f\xd4\xa3\xb6\xf9\x6c\xe4\xd8\x49\x2d\x42\x15\xd4\xc7\xfe\x71\xc7\xfe\x8b\xbe\x54\xe9\xa4\x2e\xd7\x30\xb6\x61\x9b\x6d\xb8\x2e\x57\xe7\xaf\x8b\x77\xbe\xb5\xa1\x79\x7e\x13\xb0\x36\xb7\x7c\xf7\x91\x9c\x37\x3a\x23\xba\xb0\xf5\x8b\x4f\x59\xe7\x59\xf8\xdb\xfe\xe9\xbf\x01\x00\x00\xff\xff\x85\xbb\x79\x4c\x29\x4c\x02\x00")

func dataElasticsearchmetrictypeJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataElasticsearchmetrictypeJson,
		"data/ElasticsearchMetricType.json",
	)
}

func dataElasticsearchmetrictypeJson() (*asset, error) {
	bytes, err := dataElasticsearchmetrictypeJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/ElasticsearchMetricType.json", size: 150569, mode: os.FileMode(420), modTime: time.Unix(1473655092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/ElasticsearchMetricType.json": dataElasticsearchmetrictypeJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"ElasticsearchMetricType.json": &bintree{dataElasticsearchmetrictypeJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
