// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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

// Mode return file modify time
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

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x4f\x73\x9b\x30\x10\xc5\xef\x9e\xf1\x77\xd0\x4c\xcf\xc1\xf8\x1f\x76\x74\x8a\xe3\xe0\x89\x1b\xbb\x65\x02\xa9\xa7\xa7\xce\x62\xd6\x98\x56\x20\x59\x5a\x12\xd2\x4f\xdf\x41\xe0\x98\x26\xb9\xa1\xdf\x2e\xab\x7d\xef\x29\x44\xfd\x8c\x9a\xf7\x7b\x8c\xd5\x9f\xd9\x1e\xbf\x41\x8e\x9c\xa5\x59\x71\x15\x0b\x99\xd6\x85\xc7\xb2\xd8\xca\x04\x39\x4b\x30\x2e\x2d\xb9\x27\x52\x81\xd4\xc4\xd9\xdc\x9d\xbb\xb6\x07\x21\x89\xb2\x1c\x65\x49\x9c\x79\x16\xed\x74\x46\xf8\x1f\x5b\x28\x65\x6f\xba\xc3\x03\x94\x82\x02\x48\x31\xcc\xfe\x22\x67\x43\xb7\x83\x97\xb2\x20\xac\xe8\xdd\xb4\x2d\x54\xdd\x7e\xcb\x36\x32\x0d\xe1\x19\x03\xa0\x23\x67\x86\xa4\x86\x14\x07\x42\xa6\xa6\x2d\xae\x32\xd1\xca\x01\xa5\x3a\xcc\xaf\x88\x33\xa7\x55\xf7\xa4\x84\x84\xe4\xe3\x9c\xd2\x72\xd3\x69\xb1\x56\x3d\x69\xc1\xd9\x91\x48\xf1\xc1\x60\x38\x9a\x39\xae\xe3\x3a\x43\x5e\xdb\xf0\xf1\x8f\x75\x0e\x29\x6e\xa1\x6a\x96\x9e\xb2\x2f\xdb\xdb\x77\xc5\x85\x10\xf2\xc5\xaf\xc8\x58\x5f\x18\xbb\x62\xce\x6f\x95\x76\xbe\xf1\x72\x50\x45\xda\xef\xdd\x01\x41\x0c\x06\x1b\x1f\x6f\xa3\x57\x85\x9c\xe5\xaf\xe6\x24\xec\x64\x83\xba\xb0\x82\xb5\x94\x54\x93\x00\x8c\x79\x91\x3a\xe1\x6c\x38\x1a\x4f\xa6\x9e\x4d\x4f\x1a\xaa\xcf\xe7\xe5\xc7\x63\xd7\x6b\xc6\x35\x66\x9d\x73\x8f\x20\x16\x18\x68\x3c\x64\x55\x03\x7f\xd5\x74\x79\x04\x6d\x90\x38\x2b\xe9\x30\xcf\xe3\x49\x73\x8b\x36\x36\x6a\xce\x22\x5d\x62\x9b\xd7\x3a\x11\xb8\x94\x45\x61\xce\x01\x6f\xa1\xfa\xae\xb0\x68\xd9\xd8\xed\xf7\xbe\xee\xa2\xf6\xed\xed\x35\x52\x77\xc9\xb5\x31\x25\xea\xcb\x32\x7e\xa5\x32\x8d\x9c\xcd\x46\x75\xf6\x7e\x0e\x99\xe0\x17\x31\x26\x27\xe5\x0c\x47\x9e\xb3\x97\xb9\x5d\xc8\x3e\xce\x89\x37\x3d\xbb\xd2\x28\x13\xb0\xff\x23\x0b\xbc\xe9\x76\xbe\x19\x74\xbf\xf9\xb1\x0c\xfd\x87\x95\xff\x73\xb1\x78\xdc\x3d\xec\x9a\x2d\xc2\x70\xc3\x19\xb5\xa2\x56\x5a\xe6\x9f\x4e\x89\xe4\x5b\x80\x73\x77\x3a\xbf\xbe\x9e\x79\xe3\x9b\xd3\xa9\x2e\xff\x0b\x00\x00\xff\xff\xe0\xd3\xa8\x5d\x62\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 866, mode: os.FileMode(438), modTime: time.Unix(1673425528, 0)}
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
	"configs/config.yaml": configsConfigYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"configs": {nil, map[string]*bintree{
		"config.yaml": {configsConfigYaml, map[string]*bintree{}},
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
