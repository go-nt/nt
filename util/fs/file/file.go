package file

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// 是否是文件
func IsFile(path string) (bool, error) {
	fInfo, err := os.Stat(path)
	if err == nil {
		return !fInfo.IsDir(), nil
	} else {
		return false, err
	}
}

// 删除文件
func Remove(path string) error {
	return os.Remove(path)
}

// 拷贝文件
func Copy(src string, dst string) error {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	srcFileInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// 源文件不是目录
	if srcFileInfo.IsDir() {
		return errors.New("source is a directory, does not support copy as file")
	}

	// 目标目录
	dstDir := filepath.Dir(dst)

	// 创建目标目录
	err = os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return err
	}

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 文件模式取源文件
	err = dstFile.Chmod(srcFileInfo.Mode())
	if err != nil {
		return err
	}

	// 拷贝数据
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// 移动文件
func Move(src string, dst string) error {
	return os.Rename(src, dst)
}
