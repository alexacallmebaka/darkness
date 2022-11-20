package emilia

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

func Untar(reader io.Reader, dirName string) error {
	tarReader := tar.NewReader(reader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(dirName, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(
			filepath.Clean(path),
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
			info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		for {
			_, err = io.CopyN(file, tarReader, 1024)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
		}
	}
	return nil
}
