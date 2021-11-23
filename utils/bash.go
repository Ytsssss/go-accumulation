package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// RunBashCommand 执行cmd命令并返回结果
func RunBashCommand(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	var stdout, stderr bytes.Buffer
	// 标准输出
	cmd.Stdout = &stdout
	// 标准错误
	cmd.Stderr = &stderr
	err := cmd.Run()
	result := fmt.Sprintf("base result is: %s and err is: %s", string(stdout.Bytes()), string(stderr.Bytes()))
	return result, err
}
