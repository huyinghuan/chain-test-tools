package resource

import (
	"embed"
	"io/ioutil"
)

//go:embed asserts/*
var resourceFiles embed.FS

// 避免直接读取配置文件，单元测试运行在/tmp目录，读取配置文件导致单元测试时无法正确找到文件路径
func Get(filepath string) ([]byte, error) {
	fileBody, err := resourceFiles.ReadFile(filepath)
	if err == nil {
		return fileBody, err
	}
	return ioutil.ReadFile(filepath)
}

func ReadFile(filepath string) ([]byte, error) {
	//return resourceFiles.ReadFile(filepath)
	return Get(filepath)
}
