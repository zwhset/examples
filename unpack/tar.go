package unpack

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateTar(filename string, src string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileWriter io.WriteCloser = file
	if strings.HasSuffix(filename, ".gz") {
		fileWriter = gzip.NewWriter(file)
		defer fileWriter.Close()
	}
	writer := tar.NewWriter(fileWriter)
	defer writer.Close()

	var files []string
	err = filepath.Walk(src, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return err
	}
	for _, name := range files {
		if err := writeFileToTar(writer, name); err != nil {
			return err
		}
	}
	return nil
}

func UnpackTar(filename string, dest string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileReader io.ReadCloser = file
	if strings.HasSuffix(filename, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			return err
		}
		defer fileReader.Close()
	}

	reader := tar.NewReader(fileReader)
	return unpackTarFiles(reader, dest)

}

func unpackTarFiles(reader *tar.Reader, dest string) error {
	for {
		header, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				return nil // OK
			}
			return err
		}
		filename := sanitizedName(header.Name)
		filename = filepath.Join(dest, filename)
		switch header.Typeflag {
		case tar.TypeDir:
			if err = os.MkdirAll(filename, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err = unpackTarFile(filename, header.Name, reader); err != nil {
				return err
			}
		}
	}
	return nil
}

func unpackTarFile(filename, tarFilename string, reader *tar.Reader) error {
	// tar bug
	dir, _ := filepath.Split(filename)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	writer, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer writer.Close()
	if _, err := io.Copy(writer, reader); err != nil {
		return err
	}
	if filename == tarFilename {
		fmt.Println(filename)
	} else {
		fmt.Printf("%s [%s]\n", filename, tarFilename)
	}
	return nil
}

func writeFileToTar(writer *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// 获取时间醝与权限标志位
	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    sanitizedName(filename),
		Mode:    int64(stat.Mode()),
		Uid:     os.Getuid(),
		Gid:     os.Getgid(),
		Size:    stat.Size(),
		ModTime: stat.ModTime(),
	}
	if err := writer.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
