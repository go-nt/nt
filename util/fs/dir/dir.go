package dir

import (
	"errors"
	"os"
)

// 是否是文件夹
func IsDir(path string) (bool, error) {
	fInfo, err := os.Stat(path)
	if err == nil {
		return fInfo.IsDir(), nil
	} else {
		return false, err
	}
}

// 创建文件夹
func Make(path string) error {
	fInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, os.ModePerm)
		} else {
			// 无权限、文件损坏等
			return err
		}
	} else {
		// 文件已存在，便非文件夹
		if !fInfo.IsDir() {
			return os.MkdirAll(path, os.ModePerm)
		}
	}

	return nil
}

// 删除文件夹
func Remove(path string) error {
	fInfo, err := os.Stat(path)
	if err == nil {
		if fInfo.IsDir() {
			return os.RemoveAll(path)
		} else {
			// 非目录
			return errors.New("the path (" + path + ") is not a directory")
		}
	}

	return nil
}

// 拷贝文件夹
func Copy(src string, dst string) error {

	// 检测目标目录
	dstInfo, err := os.Stat(dst)
	if err == nil {
		if !dstInfo.IsDir() {
			return os.MkdirAll(dst, os.ModePerm)
		}
	} else {
		if os.IsNotExist(err) {
			return os.MkdirAll(dst, os.ModePerm)
		} else {
			// 无权限、文件损坏等
			return err
		}
	}

	// 读取源文件
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// 循环拷贝文件
	for _, F := range files {
		srcPath := src + "/" + f.Name()
		dstPath := dst + "/" + f.Name()

		fInfo, err := os.Stat(srcPath)
		if err != nil {
			return err
		}

		if fInfo.IsDir() {
			if err := Copy(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			f, err := os.ReadFile(src)
			if err != nil {
				return err
			}

			err = os.WriteFile(dst, f, 0644)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

// 移动文件夹
func Move(src string, dst string) error {
	_, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	_, err = os.ReadDir(dst)
	if err != nil {
		// 目标文件平不存在时，自动创建
		err = os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return os.Rename(src, dst)
}
