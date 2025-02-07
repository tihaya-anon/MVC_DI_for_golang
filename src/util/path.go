package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// CopyFile 复制单个文件
func CopyFile(srcFile, dstFile string) error {
	source, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

// CopyDir 复制整个文件夹
func CopyDir(src string, dst string) error {
	// 获取源目录的文件信息
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// 如果目标目录不存在，创建它
	err = os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	// 遍历源目录中的每个文件/文件夹
	for _, file := range files {
		srcFile := filepath.Join(src, file.Name())
		dstFile := filepath.Join(dst, file.Name())

		if file.IsDir() {
			// 如果是文件夹，递归复制
			err := CopyDir(srcFile, dstFile)
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，直接复制
			err := CopyFile(srcFile, dstFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteDir 删除目录及其内容
func DeleteDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			err := DeleteDir(filePath)
			if err != nil {
				return err
			}
		} else {
			err := os.Remove(filePath)
			if err != nil {
				return err
			}
		}
	}

	return os.Remove(dir)
}

func MoveDir(oldPath, newPath string) error {

	// 复制文件夹到新路径
	err := CopyDir(oldPath, newPath)
	if err != nil {
		fmt.Println("Error copying folder:", err)
		return err
	}

	// 删除原文件夹
	err = DeleteDir(oldPath)
	if err != nil {
		fmt.Println("Error deleting original folder:", err)
		return err
	}

	fmt.Println("Folder moved successfully.")
	return nil
}
func CreateDir(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("unable to create dir: %s, error: %v", path, err)
	}
}
