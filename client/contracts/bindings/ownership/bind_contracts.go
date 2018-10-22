// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../truffle/contracts/ownership/MultiSigWallet.sol (12.489kB)

package ownership

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _TruffleContractsOwnershipMultisigwalletSol = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x5b\x6f\x1b\xb9\xf5\x7f\xcf\xa7\x38\xc9\x83\xff\x52\x22\xdb\xca\x65\xb1\xf8\xc7\xab\xa0\xae\x93\x14\x06\xb2\x4e\x11\x3b\x4d\x01\xc3\x30\xa8\x99\x23\x89\xc8\x88\x54\x49\x8e\x2d\x6d\x36\xdf\xbd\xe0\x65\x66\xc8\x21\x47\x1e\x3b\x4a\xd1\xea\x21\xb1\x48\x9e\x0b\xcf\xf9\x9d\x0b\x49\xad\x04\x99\x2f\x09\x48\x5e\xd0\x9c\xaa\x0d\x8c\x0f\x5e\x1d\xbc\x78\x75\xf4\xe8\x51\xc6\x99\x12\x24\x53\xf0\x7b\x59\x28\x7a\x4e\xe7\x5f\x48\x51\xa0\x82\x6f\x8f\x1e\x01\x00\xe0\x0d\x32\x05\x27\x9c\xcd\xa8\x58\x12\x45\x39\x1b\x90\x3c\x17\x28\x25\x50\x96\xe3\x1a\x73\x90\xc8\x72\x14\x23\x28\x29\x53\xf5\xa0\x12\x84\x49\x92\x69\x82\xd3\x7c\x78\xe4\xf1\xfa\x84\x37\x3c\xdb\x09\xa7\xf3\x72\xba\xa4\x52\x6a\x4e\x3d\x29\xde\xad\x31\x2b\xd5\x43\x08\xde\x13\x5a\x94\x02\xfb\xd2\xbd\xc5\x15\x97\x54\x6d\xdf\xe1\x0d\x29\x4a\x0c\xc8\x3e\xde\x32\x14\xc7\x79\x4e\x93\xe6\xe1\x7a\x36\x5e\xff\x09\x97\xfc\x86\x14\x3d\x96\x7f\xc2\x7f\x95\x54\xe0\x12\x99\x3a\x59\x10\x36\x77\xdb\x11\x76\xb8\xda\x81\xf9\xc7\x4c\x64\x9c\x49\x45\x98\x82\x55\x39\x2d\x68\x06\xbf\x1f\xff\xf3\xfa\xe3\x97\xb3\x77\x9f\xae\x4f\x3e\x7e\x3e\xbb\x80\x09\xfc\x32\x3e\xb2\x38\x59\x92\xd5\x8a\xb2\x39\x58\x8e\x93\x37\x70\xd1\xd8\x66\x58\xd1\x7b\xf6\x92\x47\x69\xb2\x7a\xa0\xda\xcd\xe4\x0d\x4c\x39\x2f\x86\x35\x93\xcc\xc3\x62\x9b\x4b\x9b\xa8\xa2\xa1\xd2\x58\xca\xae\x76\x8b\x2e\xaf\xaa\x59\x63\x28\xc7\xca\xe8\xe1\xc6\x2b\xbb\xc4\x33\xde\x46\x4e\x78\xc9\x94\x33\x82\x54\xa2\xcc\x94\xbf\x75\xf8\x66\x26\x3c\xb1\x90\xa3\x54\x94\x19\xf5\x8f\xea\xc9\x06\x0f\xcd\xd8\x74\xa3\x50\x42\x4e\x14\xf1\xc6\x38\x2f\x00\x0d\x28\x2b\xbd\xbe\x3b\x07\xf0\x9c\xce\x28\x0a\xe0\xac\xd8\xd8\x18\x1e\x0c\x3d\xf1\x6e\x33\x83\xa5\x9c\x1f\x58\x14\xc2\x64\x52\x29\x35\x50\x0b\x2a\x87\xc3\x46\xce\x75\x07\x6f\x6d\xa9\xb7\x1c\xe5\x19\x57\xef\xd6\x54\x36\x00\xb7\x60\x4b\xc8\x7b\xec\x6c\x7f\x69\x56\x5c\xf5\x94\x61\x98\xcb\xbb\xb9\xdf\x97\xb9\xe7\x38\x27\xc2\x58\x3e\x0c\xe4\x84\x1c\x1f\xb9\x97\xc1\xea\xab\x03\xcf\xa1\xf0\x78\x02\xe3\xbb\xb5\x70\x10\xc6\x3c\x21\x7d\x04\x77\x6e\x3a\x88\x80\x96\x36\xbd\x2d\xc1\xb8\x3a\xf9\x31\x35\x1e\xef\x4a\x8f\x77\x0e\xcf\x3d\x7d\xf1\x78\x9b\x33\xaa\xd8\xe8\x25\xf8\xac\x2c\x9a\xb4\x79\xed\xfe\x48\x89\xac\xe6\x7a\xfa\xf7\x86\x14\x34\xf7\x92\xad\xdd\x98\xb1\x88\x49\x17\x2e\xff\x5f\xd7\x89\x37\x21\xb2\x59\x0d\xbf\x4d\xa2\xcc\xbb\xb7\xd7\x50\xeb\x79\x6f\x75\x30\xa5\xf5\xd5\x23\xde\xfc\xd6\x2d\x1c\x1e\x1e\xc2\x5f\x72\xbc\x81\xf7\xa4\x28\xa6\x24\xfb\x0a\xb3\x92\xd9\x4c\x46\x8a\x82\xdf\x4a\x50\x1c\x72\x5b\xdc\xe0\xeb\xe7\xf3\xb7\x07\x86\xac\x5a\x34\x18\xc2\x8a\x6c\xc8\xb4\xc0\x2a\x53\x36\x1b\xa3\x33\x30\xc9\xc7\x64\x39\x78\x03\xe3\x61\x5d\x25\x9b\x9c\x34\x82\x7a\xc9\x30\x54\xec\xa9\x65\xf4\x14\xfe\x6e\x19\x57\x22\xa5\x1b\x3f\x0c\xf5\x3f\xa9\xda\x1a\x53\xc3\x74\x5a\xe6\x02\x24\x2a\x5d\x1e\xa9\xa2\xa4\x70\x59\x1f\x08\xcb\xeb\x44\x0f\xac\x5c\x4e\x75\x0a\x9a\x85\x55\xe6\xa0\xe1\xbd\x22\x82\x2c\xe1\xda\x11\x7f\xa0\x52\xe9\xd5\x21\xcf\x78\x79\x2d\xe0\xac\x16\x50\x0f\x25\x24\xd5\x36\x0f\x7b\xb2\x41\x53\xb9\x9c\x02\x11\x90\x6a\x6b\x5b\xf3\xd7\x5f\x23\x44\x3a\x06\x07\x05\xb2\xb9\x5a\x8c\xda\x2c\x1a\xb7\xcd\xb8\x70\xd5\x99\xc2\x04\xc6\x47\x40\x7f\x0b\x69\x8f\x80\x3e\x7b\xe6\x23\x18\x52\xd9\xdf\xd1\x5c\xd2\xab\x2b\x83\xd0\xfa\x6b\x1b\x8f\x06\x2a\x09\xa2\x09\x28\xe1\x57\xc7\xef\xf5\x5f\xce\x17\x93\x8a\xe9\x51\x3b\x96\x72\x3d\x17\x16\xf3\x36\xdc\x8f\x6b\x70\x93\x3c\x07\x02\x0c\x6f\x2d\xdf\x83\xa0\x9a\x2f\x88\x59\x33\x45\xdd\xc8\x29\x98\x6e\xe0\xd6\x78\x26\xf2\xb8\xa1\x85\xe3\x2a\x85\xce\x3c\x86\xa1\x87\x49\x9e\x9b\xbd\xb6\x4a\x5d\x97\x1f\x9b\xe2\x1e\x6e\x3f\xa8\xc9\x2d\x16\x55\xa2\x6b\x0d\x47\x90\x08\xbc\x0a\xcf\xe0\xf9\x08\xba\x40\x11\x96\xdd\xc8\x37\x8e\xd3\xaa\x94\x8b\x81\xdf\x85\xea\x4f\xd8\xe2\xfa\xb3\xdd\x3e\x11\xba\xc3\x45\x20\x6c\xb7\x3e\x49\xf9\xc3\x8a\xda\x85\x4b\x5c\x7f\xe1\x11\x6f\xb1\xdf\x8c\x14\xd2\x33\x60\x22\xe8\x42\xef\xec\xc3\x73\x1b\x77\x61\xdc\xcc\x60\xd0\x04\xd6\x64\x12\xd7\xee\xd0\x43\x66\x55\xf5\x77\x24\xe0\xea\x28\x22\x9b\x0a\x24\x5f\xc3\xe1\x76\x1c\xd6\x0c\x26\xf0\xfc\x28\x48\xfe\x75\x34\xbe\x09\x97\x86\x7b\xc8\xcc\xc1\xa4\x13\x97\x6d\x24\x55\x87\x9f\xbe\x40\x5a\x15\x24\x6b\x90\x04\xb7\x54\x2d\x7e\x46\xb8\xdb\x01\x4b\xeb\x84\xe6\x11\x15\xc3\xdb\x8f\x7d\xf3\x84\x63\x92\x00\x66\xd3\xa9\x55\xfc\x76\x00\x55\x48\x26\x96\x50\xc0\xd6\x1a\x91\x28\x11\x3b\x80\x6a\xa5\xc0\x3d\x91\x79\x57\xb8\x55\xf3\x15\xfb\x38\xa3\x75\x62\x0d\xa2\x94\x56\x1b\xe9\x2e\x30\x5a\xa4\x83\x5a\xa0\xd7\x77\x74\xb4\x05\x0f\x83\xe5\x83\xfb\x8e\x38\x08\x7b\xf6\x19\x09\x84\x6d\xaf\x33\x5b\x3a\x8f\xee\xea\xad\x3f\xf1\x5d\xc6\x75\xeb\x1e\xa3\xc3\xf2\x75\xec\x2b\x0e\xb2\x9c\x2e\xa9\x32\x3d\xa0\x33\x07\x10\xff\x00\x12\x59\xd4\x3f\xeb\xf9\x2e\x51\x44\xcc\x51\x55\x91\x18\x91\xd9\x9e\xd7\x27\x40\xb5\xb0\x07\x85\x12\x63\x21\x44\x91\x60\xb1\x19\x58\x91\x4d\xc1\x89\x9f\x45\x04\xaa\x52\x30\xf8\x64\xfe\x93\xbe\xde\x70\xda\x6a\xcc\xdd\x4e\x3d\xa6\x83\xc4\x7d\x84\x7f\x2f\x35\xf2\xee\x1f\x3a\xdd\x2d\x9c\xe8\xd4\xc9\xad\xe5\xcd\x60\x12\xcc\xcd\x83\xaf\x4d\xa0\x85\x53\xc0\x88\x6e\x3c\xee\x3c\xe4\x53\xa5\x2e\xe0\x7a\xb8\xbd\x9f\xaf\x43\x8d\x2f\xb6\x59\x37\xa1\x5a\x97\x45\x52\x41\xe3\xe5\xe0\xe6\x28\x34\x4c\x99\xce\xad\xea\xe0\x1b\x9c\xe9\x5b\xc7\xf9\x36\xe3\x6f\x6d\xbb\xa6\x8f\xf1\x0d\x59\x9c\x14\x83\x5b\x61\xff\x0c\x97\x72\x8b\xfe\xb8\xc3\xf9\x0e\x3c\x28\xf0\x86\x7f\x45\x20\x81\xf2\xa6\x0e\xed\xcc\xa7\x56\x44\xb0\xc7\x1d\xfa\x34\xeb\xeb\x27\xe7\xd7\xfa\x8e\x64\x6b\x90\xdd\xc3\x95\xad\x12\xe8\xdd\xca\xdf\xe5\xc9\x4e\xf7\x6c\x38\x43\xed\x1c\xe7\xe6\xc6\x3b\xe1\x55\xf9\x8f\x38\x25\x81\xa0\xff\x7a\xa7\xe8\x56\x87\xca\x8e\xc0\x1c\xb6\x5b\x1e\x7f\xfb\x52\x71\x41\x74\x8b\xb0\x66\x26\xf8\x3a\xaf\xbd\xc2\xc6\x47\xad\x59\x7d\x0f\x16\x05\x6d\xa5\x11\xae\x15\x0a\x46\x8a\xeb\x8c\x14\xc5\x40\x93\x04\x29\x58\x0f\xb8\x34\x6c\xe6\x88\x22\x75\xc5\xae\x06\x86\xc3\xa8\x0d\x6b\x1e\x58\xba\x72\x80\xc9\x03\x85\xc4\x44\xa7\x17\x3d\xb6\x6c\xe3\x91\xd8\x68\x0b\xd2\x10\xf4\x80\xdf\x43\xec\x82\xde\xb6\xe9\xa4\xa6\x88\x0c\x24\x6a\x28\x6a\x36\x94\x29\x0e\x54\x99\xe6\xba\xc1\x1d\x65\xc0\x45\x6e\x73\x8f\x22\x3a\xf3\xe4\x37\x84\x29\x32\xc7\x8a\x1f\x9f\x99\x4e\xee\xdc\x3d\xb0\xfd\x9f\x84\x8c\xe7\x08\x73\x64\x28\x88\xe2\x86\x72\x25\x78\x5e\xea\xe3\x07\x14\x9c\xaf\x40\x2d\x88\x82\x8c\xaf\x28\x4a\x50\x6b\x63\x52\x2b\x7e\x89\x4b\x2e\x36\x11\xf4\x7d\x87\xdd\x59\xbd\xcd\xdf\x9a\xe5\x07\xe7\x35\xaf\x9c\xc3\x4a\xd0\x1b\xa2\xb0\xa9\xdf\xf6\xb5\xa4\x71\x89\x79\x67\x10\x28\xcb\x42\x35\x06\x3d\x7c\x6a\xde\x0f\xcb\xe5\x7e\x4e\x25\x99\x16\xb8\xcf\x70\xad\xf6\x0b\xca\x74\x27\x9a\x95\x82\xaa\xcd\x21\xe3\xfb\x94\xe9\xa1\x7d\x22\x25\x2e\xa7\xc5\xa6\xba\x9b\xd3\x9f\x7a\x2c\xf4\x7e\x81\x0a\xd6\xf0\x7a\x02\x4b\xdd\xe2\x0c\xc6\xeb\x57\xe3\xa1\x35\xeb\x13\x9d\x5d\x32\xa2\xf0\x89\xb3\x8a\x49\xf2\xbc\x54\xab\x52\x81\x59\x08\x54\xc2\xed\x02\x05\xc2\x93\x99\x40\x74\xcb\x9e\xc0\x8a\x53\xa6\x50\xe8\x69\x1d\x47\x98\xeb\x46\x39\xe3\xec\x06\x99\x79\x94\x8a\x14\xc8\xb5\x02\x24\xcf\x07\xda\x46\x23\x78\xf9\x62\xa8\x15\x78\x4f\x85\x54\xf0\xf2\x85\xb3\x1f\x11\xb6\x63\x5f\x91\x3c\xc7\x1c\xdc\x61\x97\xcf\xc0\x12\x49\x9d\xfd\xb2\xa2\xcc\xd1\x78\xb7\x75\x27\xa6\xcd\xa9\x85\x18\x0f\x46\x68\x96\xe5\x74\x30\x27\x72\x04\x2f\x5f\xfd\xfa\x7c\x3c\x1e\x8e\xac\x05\xcc\x37\xbd\x0b\x2d\xd6\x36\x92\x06\x38\xf5\x53\x2e\x95\x90\x95\x42\x20\x53\xc5\x06\x70\x49\x95\xa2\x6c\x1e\x71\x4f\x7c\x0e\x0f\xe1\x54\x01\x65\x46\x5f\x69\xb4\xfa\x1b\x91\x30\xf8\x75\x3c\x1e\xc2\x33\xf3\xfd\x1f\x28\x36\x1f\xf8\x2d\x0c\x5e\x8e\x0c\x80\x89\xb5\xff\xf9\xe7\xbf\xd6\x2b\xb4\x46\x26\x69\xcd\x50\x18\xf2\xff\x1f\x1b\xfa\x9e\x1a\x68\x1e\x67\x78\x7b\x9c\x65\xbc\x64\xca\x30\x78\xf1\xcb\x78\x3c\x1e\xe9\x98\xcb\x88\xb4\xe6\xf6\xfb\xee\x1a\xfa\x1c\xa5\xce\xc2\x80\x3a\xa3\x9b\x26\x9e\x21\xe6\x12\x32\x81\x44\xdb\x20\x4e\x4f\x7e\xb4\x44\x93\x36\x70\x62\x9a\xc4\x90\x17\x56\xcd\x4e\xce\xe9\x1f\x58\xa5\x01\xca\x0c\x3e\x29\xb3\xb0\x19\xc2\x3e\xa8\x05\x95\x16\xab\x44\xc1\x8c\xae\x51\xd6\x40\xa2\x6c\xae\x73\xc3\xb4\xc0\x65\x24\x6b\x1d\x8b\x1f\x27\x0d\xf9\xd1\xc6\x84\x16\x32\x67\x1a\xef\x23\xcd\x5f\xe0\x8c\x3b\xc8\xba\xa0\x91\x5a\x4d\x2a\xe1\x0f\x14\x3c\x60\x3d\x6c\xe5\x4a\xa8\x1b\xfc\x20\x13\xb4\xcb\x7f\x7d\xfc\x58\x60\xd8\x8d\x49\x45\x54\x69\x2e\x35\x7e\xb4\x27\xf3\x4f\x3b\x27\xb1\x88\x56\xa2\xf4\x6b\xed\x3d\x9a\x83\x1b\x8a\xb7\xf1\xc1\xc6\x24\xc6\x56\x41\x77\x6f\xe3\x25\x53\xe6\xb6\xe3\x9e\xd7\x20\xad\xd4\xa7\x6b\xf1\xdd\x0f\x69\xe6\x0e\x3c\xc6\xb3\x55\xe2\x59\x70\xd5\xd6\x30\x35\x0a\x4e\x20\x3e\xb4\xb7\xbc\x9b\xba\x58\x8f\x9e\x5d\x4e\x99\x2d\x3f\x77\x3d\xbc\x1c\xe7\xb9\x74\x37\x6a\xfe\x81\x54\xd7\xcd\x05\x06\x43\xee\xd5\x7e\xa4\xb5\xf5\xc7\x5b\x71\xbd\x49\x5c\x6c\xfc\xcf\x1e\xc3\x5b\xa7\xde\x07\x9d\xc1\xa9\xf3\x84\xdf\x89\x9a\x0b\x7e\x8f\xcb\x70\x17\x47\xf4\xf8\x57\x0e\x89\x95\x6d\xbc\xc2\xc4\xb7\xd9\x20\x04\xbb\xa7\xe1\xeb\xee\x64\x6c\xf6\xfe\x3a\x95\x8f\xb5\x21\x5e\xdb\x52\x1b\x36\x94\xae\x09\x7c\x6d\x7b\xc0\x06\xcb\xc3\xa4\xce\x27\x89\xb0\xf1\x7e\x4c\xb4\xed\xd0\x53\xc7\xc3\x17\x9c\xbe\xb4\x4d\xe4\x1d\x01\x51\xe1\xa2\xe3\x6d\x71\xd7\xf9\xf1\x6c\xdb\x13\x66\x8d\xc3\x39\x2a\x3f\x91\x1a\x83\xfc\x78\xb2\x6c\x12\xe3\x8e\x6e\x88\xfb\x26\xc6\x28\x11\x76\x16\x29\xae\x48\xe1\xb9\xc2\x07\x32\x90\x99\xee\x17\x67\xb4\x30\xef\xc2\x02\x81\xac\x56\x05\x4d\x5c\xdc\xaf\x90\x99\x9a\x7d\x6a\x5b\xa7\xfa\xbb\xcf\x2d\x22\xaa\x0f\x2a\x15\x55\x3d\xd0\x41\xe6\x1c\x7a\xd1\x4b\x65\x95\xd4\xd9\xf7\xf7\x45\x0b\xff\xa6\xb4\x55\xaa\x8f\xc2\x5f\x16\xfd\x4c\xbf\x47\x59\x25\xed\xfa\xca\xa6\x7b\x7b\x10\xfe\xe0\x83\x36\x3f\xf2\x80\x3f\xff\x6c\xac\xb8\xb7\x07\x5d\xeb\xfa\xe3\xa3\x70\x0f\xfa\xd1\x43\xbe\x73\xc6\x07\x7f\xbe\xaa\x33\x98\x88\x2e\xf3\x04\x20\x07\xf7\x34\x64\xfd\xc0\x1f\xdf\x85\x1b\xf1\xfe\x03\x77\xd7\x16\x88\x10\x64\x63\x1f\xb6\x5a\x5a\x8e\xe0\x76\x41\xb3\xc5\xce\x6e\x65\x52\xd5\xcf\x4a\xef\x63\x21\x3f\xff\x24\x7f\x86\xf5\x40\xd3\xc1\x75\x90\x35\xda\x96\x6c\xd6\xb9\xf3\x64\xb0\xfa\x02\x97\x2b\xfb\xce\xd4\x2c\xec\x7c\x80\xec\xea\x03\x2d\xde\x5b\x7d\xe1\xcf\xc9\x7b\xf1\x25\x4a\xb4\x9d\x4b\xa3\xa1\xf7\xd0\x4b\x13\x0f\xbb\x5d\x8d\x64\x73\x12\x08\xad\x1a\xd9\xc8\x46\x7f\xd7\x9e\xb3\x8e\x40\x0f\x99\xda\x37\xbe\x78\x03\x95\xc2\x77\x05\x6d\xd8\x79\x49\x7d\x80\xcc\x71\x46\x19\xe6\x20\x08\x9b\xc7\x1d\xde\x4c\xf0\x25\x9c\xb2\x1c\xd7\xfa\x24\x21\x14\x98\x9f\x23\x69\xf2\x16\x37\x83\xea\x38\x42\xb8\x23\x46\x96\xdf\x93\xf4\x3f\x5a\x41\xa2\xd8\x6c\x19\x6a\x6b\xad\x38\xcd\x5d\x70\x6a\x63\xb9\xe6\x54\x71\x57\x2f\x76\x56\x3d\x74\xd8\x06\x30\x8f\xe2\xd6\xad\x72\x41\x1b\xae\xf5\xa2\xd6\x2e\x1b\xb4\x8b\xcc\x0f\x47\x6d\xcf\xaa\xa5\xf7\xdc\xa3\x70\x45\xf1\xd7\xb7\x90\x05\x84\x71\xf0\xc7\x66\x69\xa2\x9f\x3e\x28\xea\x43\x8e\x2d\x23\x73\xd8\x37\xb0\x48\xc5\xbd\x1e\x37\x86\xe3\xa9\xb8\x0f\xd9\x5e\x52\xc7\xe8\x2a\x3c\x75\x54\x7b\x68\x12\xc0\xf7\x47\xff\x0e\x00\x00\xff\xff\x68\xc2\xb9\x50\xc9\x30\x00\x00")

func TruffleContractsOwnershipMultisigwalletSolBytes() ([]byte, error) {
	return bindataRead(
		_TruffleContractsOwnershipMultisigwalletSol,
		"../../truffle/contracts/ownership/MultiSigWallet.sol",
	)
}

func TruffleContractsOwnershipMultisigwalletSol() (*asset, error) {
	bytes, err := TruffleContractsOwnershipMultisigwalletSolBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../truffle/contracts/ownership/MultiSigWallet.sol", size: 12489, mode: os.FileMode(420), modTime: time.Unix(1535712721, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0x34, 0xf2, 0xe7, 0xec, 0x7c, 0x9d, 0x5d, 0xa9, 0x81, 0x87, 0xb4, 0xd8, 0x96, 0xb1, 0x85, 0x3, 0x12, 0x16, 0x20, 0x3e, 0x5d, 0x67, 0x83, 0x2d, 0x8d, 0xe9, 0xf9, 0xc, 0x78, 0xde, 0x65}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"../../truffle/contracts/ownership/MultiSigWallet.sol": TruffleContractsOwnershipMultisigwalletSol,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"truffle": &bintree{nil, map[string]*bintree{
				"contracts": &bintree{nil, map[string]*bintree{
					"ownership": &bintree{nil, map[string]*bintree{
						"MultiSigWallet.sol": &bintree{TruffleContractsOwnershipMultisigwalletSol, map[string]*bintree{}},
					}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
