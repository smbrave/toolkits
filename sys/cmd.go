package sys

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func CmdOut(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func CmdOutTimeout(timeout time.Duration, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Start()

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		if err = cmd.Process.Kill(); err != nil {
			log.Printf("failed to kill: %s, error: %s", cmd.Path, err)
		}

		return "", fmt.Errorf("cmd[%s] timeout[%d]", cmd.Path, timeout)
	case err = <-done:
		return out.String(), err
	}
}
