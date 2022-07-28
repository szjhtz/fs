package exec

import (
	"bufio"
	"context"
	"github.com/farseernet/farseer.go/utils/str"
	"io"
	"os/exec"
	"strings"
	"sync"
)

// RunShell 执行shell命令
func RunShell(command string, receiveOutput chan string, environment map[string]string, workingDirectory string, ctx context.Context, done context.CancelFunc) {
	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	cmd.Dir = workingDirectory
	cmd.Env = str.MapToStringList(environment)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		receiveOutput <- "执行失败：" + err.Error()
		return
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go readInputStream(stdout, receiveOutput, &waitGroup)
	go readInputStream(stderr, receiveOutput, &waitGroup)

	go func() {
		err := cmd.Wait()
		if err != nil && !strings.Contains(err.Error(), "exit status") {
			receiveOutput <- "wait:" + err.Error()
		}
		waitGroup.Wait()
		done()
	}()
}

func readInputStream(out io.ReadCloser, receiveOutput chan string, waitGroup *sync.WaitGroup) {
	defer func() {
		waitGroup.Done()
		out.Close()
	}()
	reader := bufio.NewReader(out)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if io.EOF != err {
				receiveOutput <- err.Error()
			}
			break
		}
		receiveOutput <- str.CutRight(line, "\n")
	}
}