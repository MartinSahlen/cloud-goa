// Code generated by go-bindata.
// sources:
// index.js
// DO NOT EDIT!

package template

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

var _indexJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xdf\x6f\xdb\x46\x0c\x7e\xf7\x5f\xc1\x3e\x49\x42\x1d\xa9\x2f\x03\x86\x04\xde\x80\x02\x1d\xba\x61\xed\x8a\x26\xc3\x1e\x8a\xa2\xb8\xe8\x68\xeb\x9a\xd3\x51\xe5\x51\x51\x0d\xc3\xff\xfb\x70\x3f\xac\xc8\x69\xf2\x30\x60\x4f\xb6\x8e\xd4\xc7\x8f\xe4\x47\x9e\x9a\xe6\xad\x72\xda\x22\xbc\x56\xed\xdd\x8e\x69\x74\x1a\xf0\x1e\x9d\x78\x50\x6d\x4b\xac\x8d\xdb\x81\x10\xf8\x01\xdb\xd5\x76\x74\xad\x18\x72\xe0\x3b\xd3\xa7\xf7\xb8\xd4\x4a\x54\x05\x87\x15\x00\xa3\x8c\xec\xc0\xe1\x04\x1f\x98\x7a\xe3\xb1\x2c\x19\x3d\xd9\x7b\x5c\x03\xe3\x57\x6c\xa5\x82\xcd\x2f\xd1\x17\xa0\x69\xae\x07\x35\x39\x90\x0e\x61\x06\x56\x4e\x83\x71\xc1\x33\x9e\xa3\xbb\x87\x2d\x53\x1f\x1f\x06\xc5\xe8\x04\x06\xa6\x16\xbd\xaf\x23\x48\x4b\xce\x0b\x0c\xb0\x01\xc6\x6f\xa3\x61\x2c\x8b\xb6\x33\x56\x7f\xc9\x5e\x45\x55\xfb\x10\xa5\x2c\xea\xe6\x70\xa8\x7f\xcb\x71\xde\xab\x1e\x8f\xc7\x62\x0d\x9f\x3e\xaf\x33\x1d\x08\xc1\x2e\x67\x74\x74\xf7\xeb\x78\x7e\xac\xae\xe2\xef\xbd\x62\xb0\xca\xcb\x3b\xf4\x5e\xed\x30\x1d\x0e\xb5\x17\x6d\x5c\xed\x51\xde\xb8\x96\x42\xb1\xca\x62\x94\xed\xc5\xcf\x45\x7e\xad\x69\xfe\xa4\x1d\x78\x51\x4e\x2b\xd6\x80\xcc\xd0\x27\x04\x1f\xcb\xba\x30\x3c\x20\x22\x73\x4d\xae\x2c\x42\x69\x8b\x35\x94\xc8\xbc\x28\x5c\xca\x9a\x2c\xd6\xc8\x4c\x1c\xac\xb5\xd0\xb5\x70\x88\x5e\xe5\xb8\xc7\xea\x01\x8e\x46\x39\x83\xa3\x51\x9e\x84\xb3\xb4\x0b\xb6\x1f\xc1\x60\x99\x39\x6c\x80\x46\x79\x14\x25\xc0\xb7\x96\x3c\x06\xfc\x96\x34\x9e\x05\x30\xdb\x74\x08\x2f\x36\x1b\x78\x55\xcd\xe7\xa1\x3c\x37\x9d\xf1\xd0\xa3\x72\x3e\x36\x39\x28\x0b\xb6\xca\x58\xd4\xd0\xc0\xa0\x9c\x69\xef\x50\xd7\x70\x4d\x30\x61\x16\x11\x74\x8a\x75\x3d\x63\xa4\xc3\x72\xe6\x7a\x04\xb4\x1e\xcf\x82\xc0\xc7\x24\xc3\xa4\xa3\xa4\x4d\x98\x8c\x74\xf1\xc0\x2a\x41\x2f\x21\xab\x61\x94\xa4\xb7\x54\xb5\x25\xc2\xef\x0e\x5a\xe5\x11\x68\x1b\x39\xf6\x61\x2e\x3a\x91\x61\x0d\x12\x32\x30\x89\x3e\xa3\x1f\xc8\x05\xb7\xdb\x40\x6a\x49\x32\x12\x28\x17\x85\x7c\x20\x7c\x26\xb4\xa6\xf9\x87\x8d\x24\xaa\x09\xa5\xc9\x8a\x69\x82\xc6\x03\x53\xa1\xb9\x56\x85\x87\x28\xc1\x38\x39\xde\xec\x9c\xb2\x19\xe5\x8d\xd3\x81\xac\x71\xc3\x98\x79\x9c\xd4\x3a\x05\xfc\xf2\x8f\xeb\xbf\xde\xd7\x3e\x36\xda\x6c\xf7\x69\x8c\xab\x73\x59\xa3\xd3\xa9\xac\x81\xdb\x71\xb5\x9a\x97\x45\x48\x1c\x32\x9b\x87\xad\xd0\x45\xe3\x5b\x91\xa1\x64\xfc\x16\x46\xde\xa7\x5e\x87\xd9\xc9\xde\xaf\x49\xef\xaf\x56\x2b\x00\x3f\x19\x69\x3b\x08\x9e\xf5\x0e\xa5\x2c\x5a\x72\x82\x4e\x2e\x64\x3f\x60\x51\x9d\x44\x12\x6b\x5e\xa8\x61\xb0\xa6\x55\x21\x4a\xf3\xd5\x93\x2b\x2e\x57\xa7\xb2\xce\xa8\xb0\x81\x47\x29\x05\xe8\x5b\xd2\xfb\xb9\xd0\xb7\x8c\xea\xee\xea\x19\xdc\xef\x17\xd3\x34\x5d\x6c\x89\xfb\x8b\x91\x2d\x86\x71\x46\x3d\x07\x0a\x42\x45\x08\x68\x61\x0f\x79\x64\xd8\x12\x43\x6b\x69\xd4\xf3\xf6\xf2\xa0\x29\x8c\x75\x67\xfc\x1a\x3c\xc1\xd7\x31\xec\x26\xab\xf6\xa0\x2c\xb9\xdd\x8c\x14\x85\x67\x24\xf8\x30\xef\xa1\x57\xee\x05\xbc\x53\xfb\x5b\x0c\x12\xf7\x1d\x8d\x56\xa7\xc5\xc6\x63\x2b\xe0\xa9\x47\xb8\x33\xa9\x9b\x03\xd3\x80\x3c\x23\x05\xb6\xa7\x1a\x44\x72\xbf\x02\x31\x38\x92\x1a\x2c\x4a\xe1\xe1\x0e\x71\x00\x23\x49\xa3\x93\xda\x47\xd6\x8e\xa6\x35\x28\xbf\x3a\x1b\x41\xe3\xa1\xa3\xe9\x87\x8c\x6e\xb1\x53\xf7\xe8\xeb\x87\x7a\x87\x65\xf7\xa8\x5b\xeb\x27\x3a\x54\xfd\xff\x1d\xa2\x56\x50\x2e\xbc\x30\xaa\xfe\x19\x05\x9c\x00\x9f\xc5\x13\xfc\x2e\xcd\x60\x95\x79\x4e\x43\xcf\x22\x1c\x57\x59\xc9\xdb\xd1\xda\xbf\xd9\x66\xdf\x81\x49\xa8\x25\x0b\x2f\xa1\xb8\x6c\x9a\x02\x5e\xc2\x2c\xe9\x8e\xbc\x14\x55\x3e\x21\x36\x3b\xe3\x54\x78\xf5\xea\x04\x15\xe6\xe8\x63\x6e\xdf\x26\x2b\xbe\x08\xd1\x8b\xcb\x25\xad\x74\x13\x15\x1d\x2a\x8d\xec\x93\xad\xce\x4f\xd9\xd6\xa3\x74\xa4\xb3\x29\x3d\x64\x0b\x63\x4f\x82\x5f\x94\xd6\x9c\xcd\x66\xc8\xa6\x91\x6d\x71\x79\xca\x27\xa4\x98\x26\x73\x71\xb9\x2f\x08\x86\x86\xd6\xd2\xa1\x8b\x97\xfa\x68\x97\x77\x48\x58\x1e\xa7\xee\xc6\xf1\x38\xb9\xa4\x2a\x32\xfa\xda\x8b\x92\xd1\xc7\x35\x93\xff\x7f\x89\xf7\xc4\xc2\x03\x25\x99\x73\x66\xd5\xc2\xe2\x74\x32\xcd\x62\x89\x37\x4f\xdd\x2a\x69\xbb\xb2\x5c\x30\x59\x84\xfa\xe9\xd5\xab\x6a\xb9\xc2\xd2\x06\x3b\x1c\xc2\x75\x54\xdf\xb0\xd9\xed\x90\xdf\xde\xdc\x7c\x80\xe3\x71\x85\xdf\x07\x62\xf1\x9f\x8a\x1f\x3f\x12\x3e\xc3\x66\x9e\x87\x47\x7b\x2d\x7f\xf0\x3c\xb5\xf7\xae\x56\xc7\xa6\x81\xc3\x21\xdd\x45\xff\x21\x42\xfc\xf8\x3a\x83\x5f\xf6\x23\x5a\xeb\xb8\xab\x1f\x22\x38\x1d\x02\xfc\x1b\x00\x00\xff\xff\xe9\xd7\x72\xc5\xca\x09\x00\x00")

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 2506, mode: os.FileMode(420), modTime: time.Unix(1490708828, 0)}
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
	"index.js": indexJs,
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
	"index.js": &bintree{indexJs, map[string]*bintree{}},
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
