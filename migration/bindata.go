// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migration/0001_initial.sql (6.387kB)
// migration/0002_player_name.sql (133B)

package migration

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _migration0001_initialSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\xc1\x6e\xe3\x36\x10\xbd\xfb\x2b\x78\xb3\x8d\x66\x81\x36\xd8\x00\x05\x8a\x1c\x64\x9b\x49\x8d\xda\x72\x2a\xcb\xc5\xee\x89\x62\xa4\x89\xcc\x46\x22\x55\x92\xf2\xc6\x7f\x5f\x50\x92\x6d\x51\x76\xbc\x92\xb6\x8b\x45\x75\x89\xc5\xf0\x71\x46\x6f\x86\x1c\xbe\xf9\xf0\x01\xfd\x94\xb2\x58\x52\x0d\x68\x93\x0d\x06\xf5\xf7\xb5\xa6\x1a\x52\xe0\x7a\x02\x31\xe3\x83\xa9\x87\x1d\x1f\x23\xdf\x99\x2c\x30\x9a\x3f\x20\x77\xe5\x23\xfc\x69\xbe\xf6\xd7\x28\x4b\xe8\x1e\xe4\x60\x34\x40\x08\x21\x16\xa1\xda\x33\x77\x7d\xe4\x6c\xfc\x15\x99\xbb\x53\x0f\x2f\xb1\xeb\xa3\x27\x6f\xbe\x74\xbc\xcf\xe8\x0f\xfc\xf9\xa6\x40\x28\x0d\x34\x25\x07\xdc\x64\xfe\x68\x40\x1b\x77\xfe\xe7\x06\x17\x56\xdc\xcd\x62\x51\xce\x64\x8a\x3c\x53\xce\xa1\x9c\x3a\x59\xad\x16\xd8\x71\x2b\x4b\xf6\x4c\x29\x12\x50\x75\x2f\xca\x61\x1a\x6e\x19\xec\x8a\xaf\x52\x08\x4d\x16\xab\x49\xe5\x02\x24\x10\x6a\x26\x78\x81\x39\x8d\x6b\x21\x12\x92\x49\x11\x4b\x50\xaa\x36\x1e\x4a\xa0\x1a\x22\x42\xb5\x79\xf3\xe7\x4b\xbc\xf6\x9d\xe5\x93\xe5\x08\x9a\xe1\x07\x67\xb3\xf0\xd1\x74\xe3\x79\xd8\xf5\xc9\x71\x5a\xb9\x46\x9e\x45\xdf\xb0\x06\x5a\xb9\x68\xf3\x34\x33\x21\x39\xfb\xdf\x60\x8c\xa6\xbf\x3b\x9e\x33\xf5\xb1\x87\xd6\xd8\x47\xb9\x7e\xf9\x35\x7d\xfe\x38\x40\x68\xba\x5a\x2c\x0c\xa6\x1a\x21\x39\x67\xa1\x88\x80\x84\xec\xb7\xcb\xc1\xc7\x3c\xea\x9b\x16\xe1\x96\x6a\x92\x88\x58\x9d\x65\x46\x9b\xac\x48\x41\x29\x1a\x83\xa1\x06\x7f\xf2\xaf\x38\xde\x88\x7c\x99\x8c\x26\x9d\x8c\x95\x6b\x8f\x8d\xd3\x2c\x05\xa5\x69\x9a\x21\x43\xaa\xe1\xb2\x1d\xee\x61\xe5\xe1\xf9\xa3\x6b\x1c\x47\xa3\xa3\xf1\x31\xf2\xf0\x03\xf6\xb0\x3b\xc5\x87\xfd\x81\x46\x66\x78\xe5\xa2\x19\x5e\x60\x13\x36\x67\x3d\x75\x66\xb8\x1e\xc8\x72\xe4\x87\x85\x2f\xa8\xdc\x3f\x24\x7c\x50\x05\x2e\x60\x51\x70\x81\x07\xc6\x75\x83\x91\x46\x50\x4b\x82\x82\x23\x29\xc1\x65\xf4\xf1\x39\x24\xfb\x89\xdc\xe0\x95\x25\x09\x09\x45\xce\x75\xd0\xc0\xe6\x5c\xb1\xf8\x70\x14\x9c\xb0\xc3\x9f\x87\x15\x34\x02\xaa\xb7\x67\xd8\x56\xd0\x04\x68\x04\x92\x18\xe3\x2a\xe8\x06\xa5\x4a\xd1\x3c\xd1\x0d\x6c\x2b\x68\x0a\x11\x0b\x9b\x46\xdb\x41\x81\xc7\x8c\x43\xd3\xe5\x56\x50\x95\x67\x99\x90\xbd\x1c\x96\x10\x0a\xde\xcf\xe1\x2f\x8c\x9f\x87\xb5\x65\x70\x84\x82\x9e\x29\xf1\x22\x19\xf0\x28\xd9\x13\xb5\x15\xba\x1b\x4d\x47\x68\x77\x86\x25\xec\xd8\x0e\x9a\x5b\xa8\x0b\x94\x14\x15\x32\xa5\x1a\x0a\xcb\x6d\x13\x91\x29\xad\xfa\x58\xcd\x24\x28\xcd\x62\x08\xba\x43\xc3\x5c\x4a\xe0\x9a\x48\xca\x5f\x3b\xee\x1c\x78\xcb\x2e\x9c\x32\xed\x72\xd8\x84\x93\xbc\x30\x69\xb1\xdc\x01\xba\x65\x3d\x12\x71\x0b\x34\xb2\x12\xa9\x03\x4d\x22\xcd\x12\x30\xd5\x5f\x3c\xff\x6d\x6e\x1e\x3b\x13\xd9\xb6\x56\x13\x88\xc8\x36\x53\x9d\xd3\x5f\x0a\x1a\x9d\x6d\xd7\xb6\xa7\x04\x0b\x59\x04\xaa\x47\x4a\xec\x60\xcb\xc2\x04\x14\x89\x40\x69\x29\xf6\x65\x8c\xba\x40\xc9\x36\x23\x12\x32\x5a\xc5\xb7\xe5\x29\xc1\x63\x50\xe5\xb1\xd6\x31\x11\x4d\xbd\x22\xe6\x3a\x40\x94\x39\xe0\xa2\xf2\xa3\xbb\x14\x8e\xe3\x0a\x41\x7b\xab\x87\xc2\x61\x63\x3b\x14\x8e\x86\xd1\x8e\x85\xc3\x42\x77\x2a\x1c\x3d\x1c\x2e\x0b\x47\x2f\x87\x2b\x86\x55\x28\x24\xf4\x2c\xcd\x36\xb6\x03\xc3\x0d\xa3\x1d\x19\xb6\xd0\x9d\x18\xee\xe1\x70\xc9\x70\x2f\x87\xb5\xd0\x34\xe9\x07\x3d\xc9\xa2\xfa\x39\x71\xba\x59\xb7\x95\x46\xc1\x49\x1b\xf5\x5d\xe8\x9a\x3e\x2a\x8d\xd4\x14\x07\x1a\x99\x2b\xee\xb8\x1c\xaf\x94\xa7\x19\x6e\x5e\x88\x49\xf5\x9e\x73\xf6\x4f\x80\x46\xb5\x9b\x6d\x85\x9d\xae\xdc\xb5\xef\x39\x46\x79\x9c\x61\xd9\xf3\xcb\x2b\xf9\x25\xb0\x05\x43\x7d\x89\xba\x64\xa8\xc6\x83\xca\xb3\x96\xb2\x01\xbb\x8f\x73\x17\xa3\x7b\x34\xe7\x5c\xcc\x26\x83\x53\x8c\x8c\xa0\x30\x52\xe2\xfe\x82\x98\xb8\xef\x28\x27\x7a\xca\x41\x1a\xa5\x8c\x9f\x69\xc1\x36\x52\x90\xd3\x14\x8a\xc9\x7f\x39\x9e\xf9\x90\xd1\xed\xdd\xdd\xb8\x21\xc3\x8e\x4d\x84\xaa\x7f\x50\x3e\xc7\x3c\x29\x83\x5a\x75\x11\xd2\x34\xe7\x4c\xef\x6d\x81\x68\x2f\xf7\x92\xd0\x58\xbd\x6b\xf2\x3f\x57\x68\x3d\x29\x7d\x4e\x44\xf8\x7a\x2e\xaf\xdb\x33\x6b\xb7\x5e\x90\x4d\x5e\xe3\xb1\x09\x2a\x4c\x13\xbd\xcf\xca\xc8\x20\xec\x6e\x96\x68\x34\x9c\x38\xee\xf0\x06\x0d\x1f\x9d\x47\xf3\x67\xb9\xf1\xf1\xb0\x19\x29\x09\x54\x09\x7e\x5a\xd7\x62\xf8\xaa\x4d\x78\xcb\x98\xdc\x13\x73\x34\x94\x07\xd0\x15\xa1\xde\x6c\x1c\xa9\x1c\x24\x29\x52\xf0\x6b\xad\x81\x46\x5b\x80\xca\x18\x34\x61\x59\xc3\xdb\x8f\x77\xe3\x0b\x7a\xb5\x32\x90\x35\xbf\xed\xf2\x6c\xeb\x24\x68\xb8\x68\x9d\x06\xe5\xd6\x29\x1a\x08\x3f\x2e\xf5\x0e\x47\x95\x04\x53\x94\xae\x37\x07\x2e\x88\xfb\x6f\x78\xae\xb7\x17\x4a\x7f\xec\x06\x83\xb1\x7f\xa1\x9b\x50\x4d\x8d\x48\xad\x25\xf1\x7d\x5c\x3d\x1a\x34\xb9\x6e\xf3\xa3\xe1\x4d\x5f\x8e\xe1\xb5\x26\x57\x4b\x83\xc7\x12\x59\xb3\x69\xf6\x8b\x19\xff\x2e\x5f\xa8\x34\xd5\x79\x43\x16\x00\xcf\x53\x34\x1a\x3e\x01\x8f\x18\x8f\x87\x37\x43\x0f\x76\x0c\xbe\x40\x54\xfc\x54\x22\xd9\x15\x3f\x67\x4c\xa5\x4c\x29\x88\x86\xe3\xf6\xed\xbd\xa0\xdc\x20\x5c\xe8\x9a\x18\xe9\x45\xe9\xf5\x5b\x40\xbd\xfc\x57\x19\x7f\xac\xe0\x23\x2b\xe7\xbe\x06\xb8\xad\x01\xea\x99\xf7\xfe\x95\xa1\x69\xcf\xbe\x31\x58\xb6\xdb\xdd\x19\x0c\x21\xc5\x06\x3a\xbb\x34\xb4\xf3\xe1\xf6\x1d\x1f\xac\xcf\xf9\x9f\x5c\x5f\xfe\x0d\x00\x00\xff\xff\xcd\x98\x33\xfa\xf3\x18\x00\x00")

func migration0001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration0001_initialSql,
		"migration/0001_initial.sql",
	)
}

func migration0001_initialSql() (*asset, error) {
	bytes, err := migration0001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/0001_initial.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc5, 0x8b, 0xd5, 0xff, 0x4, 0x71, 0x33, 0xae, 0x53, 0x7d, 0x7e, 0xfd, 0x54, 0x53, 0xbc, 0x9f, 0xf6, 0x39, 0x6d, 0x9, 0x14, 0xbd, 0x27, 0xe5, 0xdd, 0x83, 0xdd, 0x3b, 0x8d, 0xe2, 0xa6, 0x1e}}
	return a, nil
}

var _migration0002_player_nameSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\x31\x0a\x02\x41\x10\x44\xd1\x7c\x4e\x51\xa1\x22\x9b\x08\x1b\x99\x09\x9e\x40\x8c\xa5\x74\xca\x75\xa0\xa7\x5d\xda\x56\xf0\xf6\x62\x66\xb0\xe1\x0f\xfe\x1b\x06\x6c\x7a\x9b\x82\x29\x9c\xe6\x52\xfe\xfb\x98\x4c\x75\x79\xee\x35\x35\x2f\xb4\x54\x20\x79\x31\x61\x36\x7e\x14\x60\xad\x70\x76\xe1\xcd\xb8\xde\x19\xab\xed\x38\xae\xe1\x8f\x84\xbf\xcc\xc0\xdb\xef\x78\xa6\xd8\xcf\xad\xee\x96\xf1\x83\xd7\x6f\x00\x00\x00\xff\xff\xdb\x50\x21\x61\x85\x00\x00\x00")

func migration0002_player_nameSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration0002_player_nameSql,
		"migration/0002_player_name.sql",
	)
}

func migration0002_player_nameSql() (*asset, error) {
	bytes, err := migration0002_player_nameSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/0002_player_name.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x6d, 0x49, 0xfa, 0x7a, 0xa2, 0xaa, 0xc6, 0x30, 0xab, 0x3d, 0x2e, 0xcb, 0xb6, 0x39, 0xdc, 0x8, 0x4f, 0xfa, 0x8b, 0xef, 0xd4, 0x7a, 0x96, 0xc7, 0x1a, 0xec, 0x94, 0x60, 0xa9, 0x7e, 0x41}}
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
	"migration/0001_initial.sql":     migration0001_initialSql,
	"migration/0002_player_name.sql": migration0002_player_nameSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"migration": {nil, map[string]*bintree{
		"0001_initial.sql":     {migration0001_initialSql, map[string]*bintree{}},
		"0002_player_name.sql": {migration0002_player_nameSql, map[string]*bintree{}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
