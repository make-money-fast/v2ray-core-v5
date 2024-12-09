package geo

import (
	_ "embed"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	//go:embed geoip.dat
	geoIP string

	//go:embed geoip-only-cn-private.dat
	geoipCN string

	//go:embed dlc.dat
	dlc string
)

func writeAll(directory string) error {
	if err := ioutil.WriteFile(filepath.Join(directory, "geoip.dat"), []byte(geoIP), 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(directory, "geoip-only-cn-private.dat"), []byte(geoIP), 0755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(directory, "dlc.dat"), []byte(geoIP), 0755); err != nil {
		return err
	}

	return nil
}

func CheckAndInit(directory string) error {
	if !fileExist(filepath.Join(directory, "geoip.dat")) {
		return writeAll(directory)
	}
	if !fileExist(filepath.Join(directory, "geoip-only-cn-private.dat")) {
		return writeAll(directory)
	}
	if !fileExist(filepath.Join(directory, "dlc.dat")) {
		return writeAll(directory)
	}
	return nil
}

func fileExist(fp string) bool {
	_, err := os.Stat(fp)
	if err != nil {
		return false
	}
	return true
}
