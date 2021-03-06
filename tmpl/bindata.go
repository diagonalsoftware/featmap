package tmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _tmpl_bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func tmpl_bindata_go() ([]byte, error) {
	return bindata_read(
		_tmpl_bindata_go,
		"tmpl/bindata.go",
	)
}

var _tmpl_email_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\x41\x0a\xc2\x30\x10\x05\xd0\x7d\x4e\xf1\x97\x0a\xa5\x3d\x83\x0b\x45\xa8\x0b\x51\x3c\xc0\xd8\x8c\x65\xa0\x4d\x42\x92\x2a\x61\x98\xbb\x8b\xb8\x7d\x8b\x77\x96\xce\xb9\xeb\xc2\x54\x18\x6f\xce\xf2\x6a\x68\x71\xcb\x08\xfc\x01\xaf\x24\x0b\xc8\xfb\xcc\xa5\x60\xa7\xda\x1f\x7f\x62\xb6\xc7\xb3\x61\x8e\x12\x66\xd4\x08\xd5\xfe\x90\xd2\x5d\x2a\x3f\x6e\x17\x33\x9a\xa6\xb8\x85\x3a\xfc\xb7\x41\xb5\x1f\xb9\x99\x39\x37\x4a\xf0\xc8\x3c\x53\xf6\xa5\x73\x27\xa6\xba\x52\xfa\x06\x00\x00\xff\xff\xbc\x0d\xdb\x19\x80\x00\x00\x00")

func tmpl_email_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_email_tmpl,
		"tmpl/email.tmpl",
	)
}

var _tmpl_invite_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xb1\x4e\xf4\x30\x10\x84\x7b\x3f\xc5\xfe\x57\xfd\x48\x27\x5f\x4f\x07\x08\xc4\x09\x44\x71\x08\x21\xca\x4d\x32\x97\x33\x24\xbb\x91\xe3\x5c\x14\x59\x7e\x77\xe4\x5c\x22\x51\x50\x50\xce\xae\x67\xf6\xf3\x3c\xba\xad\x31\x1f\x3a\xd0\x89\xcf\xa0\x02\x10\x72\x72\x76\x01\x15\x15\x13\xc5\x68\xf7\x17\x75\x3b\xa5\x44\xff\x7f\xea\xfb\x96\x5d\x93\xd2\x15\x05\xa5\x4f\x75\x42\xe1\x04\x1a\xd5\x7f\xf5\x1d\x97\xa0\x4d\x8c\xf6\x7d\x55\x2f\xdc\x22\xa5\x0d\x71\xa0\x07\x70\x68\xb9\xfb\x77\xb9\x5a\xb2\xfc\x66\x2e\x26\xaa\xd5\x49\x9d\xb3\x63\xb4\x37\x5d\xf7\xea\x02\xde\x0e\xcf\x29\x71\x59\xea\x20\x61\x37\x53\x72\x70\x2a\xbb\x18\xed\x9d\x56\x48\xc9\x98\xfd\x91\x26\x1d\x88\x3d\x48\x30\x66\xfb\x72\x6f\x3b\xcf\x05\xa8\xf2\xb0\xf4\xe0\x00\x62\xa1\x25\xce\xd2\x4a\xb3\xae\x66\xff\xb2\xfd\x0b\x4f\xef\x6a\x19\x3a\x4b\x07\xb4\x68\x0b\xf8\xfc\x74\xe8\x31\x7f\xec\xa8\x4d\xa3\x63\xf6\x23\x97\x76\x9d\x33\x96\xfa\xac\x31\x4f\x4e\x2a\xf2\xa8\xd9\x57\xfd\xd6\x2c\xbc\xdf\x01\x00\x00\xff\xff\x6c\x17\x12\x23\x98\x01\x00\x00")

func tmpl_invite_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_invite_tmpl,
		"tmpl/invite.tmpl",
	)
}

var _tmpl_reset_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\x41\x4e\xc3\x40\x0c\x45\xf7\x73\x8a\xbf\x04\x29\x4a\xcf\xc0\x06\x81\xca\xaa\x15\x07\x70\x1b\x27\xb1\x94\x8c\x87\xb1\x03\x1a\x8d\xe6\xee\x28\x51\xc5\x92\xb5\xfd\xdf\x7b\x6f\xd2\x85\x70\xd5\x95\x6f\x3a\x14\x3c\xcd\x9a\x78\xdc\x96\xa5\xa0\xe8\xf6\x8c\x99\x0c\x99\xbf\x36\x36\xe7\x01\x84\x44\x66\x3f\x9a\x07\x64\x36\x76\x8c\x9a\xf7\xbf\x0c\xba\xdf\x75\x8b\x0e\x89\x78\x65\xf2\x95\x52\x8f\xf7\x71\xbf\x61\xa6\x6f\x46\x54\xff\x87\xd3\x21\x2d\x4c\xc6\x90\x29\x6a\x66\xf8\x2c\x06\x5e\x49\x96\x3e\x84\xcb\x61\x3a\x2c\x7f\xab\x5b\xc1\xa4\x12\x27\xb8\xa2\xd6\xfe\x25\xa5\xab\x38\x7f\x5e\x3e\x5a\x7b\x94\x9c\x0e\xf0\xa9\xd6\xfe\xcc\xa5\xb5\x10\xce\x12\x77\xdb\x44\x79\xb0\x2e\x3c\x22\x7f\x03\x00\x00\xff\xff\x68\xbc\xf4\x21\xfd\x00\x00\x00")

func tmpl_reset_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_reset_tmpl,
		"tmpl/reset.tmpl",
	)
}

var _tmpl_welcome_tmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\x31\x4b\xc4\x40\x10\xc5\xf1\xda\xfd\x14\xaf\x53\xe1\xc8\xf5\x76\x16\xca\xc1\x59\x29\x2a\x96\x73\xd9\xb9\x30\x98\xec\x86\xd9\xc9\x1d\x61\x98\xef\x2e\x11\x6c\x6c\x1f\x8f\x1f\xff\x83\xec\xd2\x81\x95\x6f\x1b\x5a\x9d\x18\x52\xce\x15\x56\x31\xb0\x61\xad\x0b\x9a\x91\x1a\xe7\x87\x94\xbe\xea\xa2\x28\x7c\xc5\xb5\xea\x77\x9b\xa9\x67\x48\x03\x5d\x48\x46\x3a\x8d\x0c\x32\xb8\x77\x8f\xf3\xfc\x26\xc6\xef\xaf\x2f\x11\xee\xdd\xe7\xdf\x37\x22\xdd\xa4\x0f\x56\x39\xaf\x9b\xab\xe0\x89\x64\x04\xe5\xac\xdc\x1a\xee\xdc\xbb\xa7\x6d\x89\xb8\xc7\x69\xc5\x50\xa5\x0c\x5b\xc7\x3f\x92\xfa\xbe\x2e\xc5\xf6\x97\x5f\x69\xef\xde\x1d\x79\x8d\x48\xe9\x28\x25\x43\x79\x20\xcd\x6d\x97\x9e\x99\x6c\xa2\xf9\x27\x00\x00\xff\xff\x16\xec\x4e\x54\xdd\x00\x00\x00")

func tmpl_welcome_tmpl() ([]byte, error) {
	return bindata_read(
		_tmpl_welcome_tmpl,
		"tmpl/welcome.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"tmpl/bindata.go": tmpl_bindata_go,
	"tmpl/email.tmpl": tmpl_email_tmpl,
	"tmpl/invite.tmpl": tmpl_invite_tmpl,
	"tmpl/reset.tmpl": tmpl_reset_tmpl,
	"tmpl/welcome.tmpl": tmpl_welcome_tmpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
		"bindata.go": &_bintree_t{tmpl_bindata_go, map[string]*_bintree_t{
		}},
		"email.tmpl": &_bintree_t{tmpl_email_tmpl, map[string]*_bintree_t{
		}},
		"invite.tmpl": &_bintree_t{tmpl_invite_tmpl, map[string]*_bintree_t{
		}},
		"reset.tmpl": &_bintree_t{tmpl_reset_tmpl, map[string]*_bintree_t{
		}},
		"welcome.tmpl": &_bintree_t{tmpl_welcome_tmpl, map[string]*_bintree_t{
		}},
	}},
}}
