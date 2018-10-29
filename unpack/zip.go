package unpack

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func UnpackZip(filename string, dest string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, zipFile := range reader.Reader.File {
		name := sanitizedName(zipFile.Name)
		fname := filepath.Join(dest, name)
		mode := zipFile.Mode()
		if mode.IsDir() {
			if err = os.MkdirAll(fname, 0755); err != nil {
				return err
			}
		} else {
			if err = unpackZippedFile(fname, zipFile); err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateZip(filename string, src string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	zipper := zip.NewWriter(file)
	defer zipper.Close()

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
		if err := writeFileToZip(zipper, name); err != nil {
			return err
		}
	}
	return nil
}

func writeFileToZip(zipper *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// 获取时间醝与权限标志位
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// 处理一下文件名的问题
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = sanitizedName(filename)
	writer, err := zipper.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, file)
	return err
}

func unpackZippedFile(filename string, zipFile *zip.File) error {
	writer, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer writer.Close()
	reader, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	if _, err := io.Copy(writer, reader); err != nil {
		return err
	}
	if filename == zipFile.Name {
		fmt.Println(filename)
	} else {
		fmt.Printf("%s [%s]\n", filename, zipFile.Name)
	}
	return nil
}

func sanitizedName(filename string) string {
	//if len(filename) >1 && filename[1] == ":" && runtime.GOOS == "windows" {
	//	filename = filename[2:]
	//}
	filename = filepath.ToSlash(filename)
	filename = strings.TrimLeft(filename, "/.")
	return strings.Replace(filename, "../", "", -1)
}
