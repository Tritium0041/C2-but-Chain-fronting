package executer

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"text/template/parse"
	"time"

	"github.com/mmcloughlin/addchain/acc/pass"
)




func doCommand(command string, args []string, timeOut int) []byte, error {
	ctx,cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()
	// Execute the command
	cmd := exec.CommandContext(ctx, command, args...)
	
	// 启动命令
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	// 等待命令完成或超时
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	result,err := cmd.Output()
	// 检查上下文是否因超时被取消
	if ctx.Err() == context.DeadlineExceeded {
		panic(err)
	}

	// 返回执行结果
	return result,nil
}