package utils

import "os"

// directory, 文件夹相关的工具函数

func PathExists(path string) (bool, error) {
	// Stat返回描述命名文件的FileInfo，如果有错误，则为*PathError类型
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// IsNotExist返回一个布尔值，指示是否已知错误以报告文件或目录不存在。ErrNotExist以及一些系统调用错误都能满足该问题。
	// 该函数先于errors.is。它只支持OS包返回的错误。新代码应该使用errors.is(err，os.ErrNotExist)。
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
