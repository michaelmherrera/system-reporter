// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// index.html
package main

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4f\x6f\xdb\x3e\x0c\xbd\xfb\x53\xb0\x3e\xff\x6c\xfd\x82\x6e\xc0\x50\xd8\x3e\x2c\xd9\xd0\x0e\xeb\xd6\x21\xed\x61\x47\x5a\x66\x6a\xa5\xfa\xe3\x49\x6c\xda\xa0\xc8\x77\x1f\x2c\x25\xcd\x96\x25\xd8\x7c\xb1\xde\xd3\x23\x1f\x45\x82\xd5\xd9\xec\xeb\xf4\xf6\xfb\xcd\x07\xe8\xd9\xe8\x26\xab\xc6\x1f\x68\xb4\xf7\x75\x4e\x36\x6f\xb2\xac\xea\x09\xbb\x26\x03\x00\xa8\x0c\x31\x82\xec\xd1\x07\xe2\x3a\xbf\xbb\xfd\x58\xbc\xcb\x7f\xbd\xb2\x68\xa8\xce\x57\x8a\x9e\x06\xe7\x39\x07\xe9\x2c\x93\xe5\x3a\x7f\x52\x1d\xf7\x75\x47\x2b\x25\xa9\x88\xe0\x3f\x50\x56\xb1\x42\x5d\x04\x89\x9a\xea\x49\xf9\xff\x2e\x15\x2b\xd6\xd4\xcc\x9c\x7c\x34\x64\xb9\x12\x09\x67\x95\x48\x95\x54\x67\x45\x01\x9f\x91\x29\x30\x48\x67\x06\xa5\xa9\x03\xb4\x1d\x18\x65\xd5\x42\x51\x07\xd3\xf9\x1c\x8a\xa2\xc9\x2a\xad\xec\x03\x78\xd2\x75\x1e\x78\xad\x29\xf4\x44\x9c\x43\xef\x69\x51\xe7\x3d\xf3\x10\x2e\x84\x30\xf8\x2c\x3b\x5b\xb6\xce\x71\x60\x8f\xc3\x08\xa4\x33\xe2\x95\x10\xe7\xe5\x9b\x72\x22\x64\x08\x7b\xae\x34\xca\x96\x32\x84\xd8\xa0\xb1\x9e\xe5\xb7\x47\xf2\x6b\xd0\xaa\xf5\xe8\xd7\xc9\x3d\x48\xaf\x06\x86\xe0\xe5\xde\x0d\x97\xf8\x5c\xde\x3b\x77\xaf\x09\x07\x15\xa2\xd3\xc8\x09\xad\xda\x20\x96\x3f\xc6\x2c\xe2\xbc\x7c\x5b\x4e\xb6\x20\x3a\x2d\x43\xde\x54\x22\xe5\xdb\x39\x1e\x76\xe0\x13\xae\x70\x9e\x1c\x4f\xba\xff\xeb\x5b\x97\x87\x4f\x3d\x2c\xa0\x75\xdd\x7a\x3b\xad\x7e\xd2\x5c\xd9\x85\xf3\x06\x59\x39\x0b\x0b\xe7\xe1\xca\x32\x3e\x50\x25\xfa\xc9\x56\x33\x34\x10\x0f\xe3\x37\x27\xaf\x50\xc3\x97\x47\xd3\x92\xbf\x80\x97\x97\x32\x31\x89\xd8\x6c\x52\x44\xeb\x9b\xd7\x88\xe9\xcd\x1d\x5c\xbb\x8e\x74\x54\x4f\x6f\xee\x22\x38\xa5\x9c\x0f\x24\xc3\x4e\x19\xc1\x66\x13\x2f\xff\x94\x5f\x93\x71\x7e\x1d\xb5\xe9\x78\x52\xf9\x1e\x99\xc7\xf9\x5e\x12\x6a\xee\x63\xc4\x96\x4a\xcc\xb1\x62\x2e\xd1\x77\x30\xf3\x6a\x45\x30\xc5\x01\xa5\xe2\x64\x75\x39\xdb\xc1\x63\x51\xb3\xb8\x23\x7f\xed\xd2\x91\x1a\xb7\x91\xfb\x4e\x5d\xa3\xec\x95\xa5\xdf\xba\x25\x86\x71\x91\xd2\xf8\xc6\x8d\x1a\x97\xfe\x67\x00\x00\x00\xff\xff\x0f\x13\xf6\xe4\x04\x04\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 1028, mode: os.FileMode(420), modTime: time.Unix(1608368368, 0)}
	a := &asset{bytes: bytes, info: info}
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
	"index.html": indexHtml,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
