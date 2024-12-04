package executer

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func DoCommand(command string, args []string, timeOut int) ([]byte, error) {
	// 创建上下文，设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()

	// 创建命令
	cmd := exec.CommandContext(ctx, command, args...)

	// 执行命令并获取输出
	result, err := cmd.CombinedOutput()


	// 检查上下文是否因超时被取消
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("command timed out")
	}

	// 返回执行结果和错误
	return result, err
}
