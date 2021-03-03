package ssh

import (
	"bytes"
	"errors"
	"os/exec"
	"time"
)

type ExecResult struct {
	Id             int
	Host           string
	Command        string
	LocalFilePath  string
	RemoteFilePath string
	Result         string
	StartTime      time.Time
	EndTime        time.Time
	Error          error
}

func CommandLocal(cmd string, to int) (ExecResult, error) {
	timeout := time.After(time.Duration(to) * time.Second)
	execResultCh := make(chan *ExecResult, 1)
	go func() {
		execResult := LocalExec(cmd)
		execResultCh <- &execResult
	}()
	select {
	case res := <-execResultCh:
		sres := *res
		errorText := ""
		if sres.Error != nil {
			errorText += " commond  exec error.\n" + "rsult info :" + sres.Result + "\nerror info :" + sres.Error.Error()
		}
		if errorText != "" {
			return sres, errors.New(errorText)
		} else {
			return sres, nil
		}

	case <-timeout:
		return ExecResult{Command: cmd, Error: errors.New("cmd time out")}, errors.New("cmd time out")
	}

}

func LocalExec(cmd string) ExecResult {
	execResult := ExecResult{}
	execResult.StartTime = time.Now()
	execResult.Command = cmd
	execCommand := exec.Command("/bin/bash", "-c", cmd)
	var b bytes.Buffer
	execCommand.Stdout = &b
	var b1 bytes.Buffer
	execCommand.Stderr = &b1
	err := execCommand.Run()
	if err != nil {
		execResult.Error = err
		// execResult.ErrorInfo = err.Error()
		execResult.Result = b1.String()
		return execResult
	} else {
		execResult.EndTime = time.Now()
		execResult.Result = b.String()
		return execResult
	}
}
