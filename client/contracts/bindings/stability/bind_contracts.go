// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../truffle/contracts/stability/Stability.sol (3.385kB)

package stability

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

var _TruffleContractsStabilityStabilitySol = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4d\x6f\xdb\x48\x0c\xbd\xfb\x57\x10\x3d\xc9\xd9\x44\xde\x2c\x7a\x33\x5c\xa4\x68\x83\x45\x0f\x4d\x82\x6d\x17\x7b\x08\x8c\x60\xac\xa1\xed\x41\x47\x33\x2a\x87\x72\xea\x14\xf9\xef\x8b\x51\xa2\xf9\x90\x95\x26\x8b\x02\x0b\x34\x97\x58\x12\xf9\x48\x3e\x92\x6f\xa6\x21\xb1\xa9\x05\x38\xab\x95\x54\xbc\x87\xdf\xcb\xd7\xe5\x1f\xaf\xe7\x93\x89\xaa\x1b\x4b\x0c\xaf\x6c\x83\xe6\x0e\x9b\x06\xb5\x32\x27\xbd\xd9\xac\xb2\x86\x49\x54\xec\x66\x5a\xad\xb1\xda\x57\x1a\x67\x57\xa2\x75\x62\xa5\xb1\x74\x56\xbf\x9a\x07\x80\x72\x76\x45\xaa\xc2\x2b\xb2\x3b\x25\x91\x06\x5f\xef\xac\x3b\xd1\x6a\x95\x00\xd6\x6a\x43\x82\x95\x35\x6e\xf6\xc1\x28\x56\x42\xab\xbb\x04\x76\x32\x99\x1d\x1d\x4d\xe0\x08\xce\x58\xb1\x46\xf8\xc4\x62\xa5\xb4\x4f\xbd\x87\x00\xd7\x36\x1e\xdc\x81\x41\xbe\xb5\xf4\x05\x5a\xee\x2c\x26\x70\x34\x9b\x04\xab\xe8\xa8\x1c\xf4\xb9\x1f\x43\x16\xf3\xfb\x64\x02\x00\xd0\x2a\xc3\x1e\xde\xb1\x30\x0c\x97\x17\xe7\xb0\x80\x53\x40\xde\x22\xcd\x3b\x83\x68\xd5\xb4\x2b\xad\x2a\xa8\x95\x79\x8f\x8d\x75\x8a\x1f\x0c\x32\x0a\xa0\x49\x9f\x12\x04\xc7\xd4\xfa\xcc\xda\x95\xab\x48\x35\x9e\x04\xf8\xde\x7d\x09\xf8\xca\x48\xfc\x36\x0f\xef\x56\xd6\x6a\xd8\x0a\x97\xba\xcc\x73\x0f\x99\xe6\x71\xff\x50\x50\x2d\x9a\x46\x99\x0d\x14\x42\x4a\x42\xe7\x60\xf1\x26\x8b\x3a\xf5\x29\xee\x04\x23\xb8\xe4\xed\x5f\xb8\x51\x8e\x69\x3f\x7f\x00\x79\xf4\xbd\x5e\x8e\x1a\xbb\x47\xab\xda\x4a\xb5\x56\x48\x60\x8d\xde\x3f\xc6\x58\x21\x25\x75\x11\x7e\x6d\x15\x61\x71\x33\xa8\xa3\xa8\xdd\xa6\x74\x68\x24\xd2\x74\x1a\x8b\xba\xc9\x2b\x49\xe1\xff\x51\xbc\xfd\x18\x98\x1f\x09\xe1\x11\x77\x42\xb7\x08\x6f\x16\x49\x8f\x9e\x47\xbf\xdd\xa2\xe9\x7a\xf8\x27\xa1\x60\xa4\xf3\xaf\xad\xd0\x97\x06\x47\x62\x64\xcd\x2d\xbb\xa7\x62\xea\xe3\x5d\x5e\x9c\x3f\x19\xa8\x9b\x69\xff\x77\x04\xef\xfc\x9c\xf9\x39\xb0\xd4\xbf\x3a\x6b\x04\x89\x1a\x6e\x62\xc6\x3e\x79\x55\xb7\x75\xdf\xdd\x3e\xba\x04\xb6\x7d\x1b\x56\xe8\x1f\x78\x8b\xe0\x90\x76\xaa\xc2\x21\x5c\x96\xe9\x5b\x29\xa9\x6f\x29\xd8\x35\xb8\xbd\x63\xac\x61\x27\x48\xf9\x4d\x70\x61\xbd\x1e\x51\x66\xdd\xff\x75\x6b\xaa\x6e\x4c\xc3\x36\x15\xdd\xd0\x25\xa9\x1e\x07\xd4\xc3\x80\xd3\x7e\x5f\x22\x8d\x49\x8d\x8b\x14\x26\x32\x97\xa1\xc0\x22\xdf\xad\x62\x24\xc8\x53\x4c\xab\x7e\xd5\x31\xd6\xb1\xb6\x04\x57\x64\xbf\xed\xe1\x4a\x30\x23\x99\xf2\x17\x6a\x42\x2c\xe8\x3f\x76\x41\xb9\x20\x7b\x5e\x9e\xfe\xef\x9e\x84\x02\x36\xc8\xa9\x02\xbc\xb3\xad\xe1\x22\x0c\xc9\x4e\xe1\x2d\x10\x72\x4b\xc6\x41\xf1\x28\xca\xad\xe1\x69\xb6\x85\xfe\x73\xae\x44\xa5\x46\xb3\xe1\xed\x8b\x42\xbe\xe5\x0f\x5e\x61\x8b\x28\xb6\x4f\x84\xef\xe9\xac\xac\xc4\xe3\x4c\x69\xd3\x74\xfc\x57\x58\xe4\xe9\x5c\x77\xb0\xcb\x48\x5e\xa6\xf7\xde\x74\xe0\xd1\xeb\xee\xb5\x47\x4b\xfc\x64\xe8\x89\xb7\x2e\xc7\x84\x3e\x94\x79\x20\xae\x7d\xfe\x4a\xa2\x61\xc5\xfb\x28\xf9\x79\x9d\xdd\x09\xe3\xe7\x43\x22\xca\x67\x98\x0e\x89\xf6\xa0\xcb\x72\xf4\x6c\x3a\x48\x4e\x19\x87\x94\xf5\xa1\x88\xf9\xfc\x50\xd9\x5f\xcc\x5d\x3c\x4c\x12\x06\x3b\xde\xba\x76\x0c\x9b\x54\x36\xad\xdb\xa6\x27\x10\x9c\xc0\xe9\xc0\x71\x50\x1b\x2c\x80\xa9\xc5\x81\x51\x6c\x52\x38\x7c\x32\x12\x52\x41\x3a\x93\xb8\x83\x4f\x07\xc2\x31\x72\xc3\xc9\xb4\x64\xa0\x02\x41\x79\xe2\xe6\x34\x62\xef\xf5\xa3\x3b\xc6\x2e\x2c\xfb\xdb\x0e\xca\x84\x45\xb5\x86\x1f\x1f\xbf\x89\xed\xcf\xb3\x7e\xc0\xcd\x6f\x07\xe4\xe4\x43\x16\xdf\xdd\x87\x5f\xf1\x20\x1d\x1b\x9e\xa7\x24\xbf\x63\xf8\x6f\x93\x88\xf3\x81\x32\x0f\xd8\x6c\xcd\x08\x9f\x83\xbb\xcc\x0b\x6e\x07\x3f\xc7\x58\x27\x2f\x64\x6f\x3f\xdb\xf7\xa8\x91\x7b\x49\x29\x07\xb7\xc1\xe8\x5b\x32\x09\xe3\xd6\x48\x45\x4a\xf4\x34\xd5\x8e\x0e\xe7\xf9\x14\x82\xc7\x6c\x06\x84\x8d\x16\xd5\x03\x69\x12\x1b\xc2\x4a\x30\x4a\x20\xac\x2c\x49\xb8\x55\xbc\xed\x3e\x69\xe1\x18\x50\x63\x8d\x86\x83\x7b\xaf\x37\x5f\x70\xff\xd9\x7e\xb4\xbb\x43\x59\x1c\xd3\xec\x93\xd3\xe5\x1c\xb2\x85\x8a\xf6\x09\x1f\x4b\x58\x44\xe0\xf9\xa8\x7d\xa8\x2e\xd8\x2d\xc3\xe6\x27\x48\xe3\xce\x21\x9d\x93\x7e\xb0\xee\xff\x0d\x00\x00\xff\xff\x61\x5e\x36\xcb\x39\x0d\x00\x00")

func TruffleContractsStabilityStabilitySolBytes() ([]byte, error) {
	return bindataRead(
		_TruffleContractsStabilityStabilitySol,
		"../../truffle/contracts/stability/Stability.sol",
	)
}

func TruffleContractsStabilityStabilitySol() (*asset, error) {
	bytes, err := TruffleContractsStabilityStabilitySolBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../truffle/contracts/stability/Stability.sol", size: 3385, mode: os.FileMode(420), modTime: time.Unix(1535712721, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x32, 0x60, 0x16, 0xe2, 0xf1, 0x75, 0x34, 0x68, 0xef, 0x47, 0x85, 0xfc, 0x17, 0x4c, 0xa7, 0x96, 0xe5, 0xfd, 0x1e, 0xb3, 0x2f, 0x70, 0x8c, 0x52, 0x0, 0x74, 0xf8, 0x52, 0x89, 0x7e, 0x40}}
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
	"../../truffle/contracts/stability/Stability.sol": TruffleContractsStabilityStabilitySol,
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
					"stability": &bintree{nil, map[string]*bintree{
						"Stability.sol": &bintree{TruffleContractsStabilityStabilitySol, map[string]*bintree{}},
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
