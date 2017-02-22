package core

import (
	"bytes"

	"github.com/lxc/lxd"
	"github.com/mrtc0/lxdexec"
)

func Compile(name string, language string, stdin_text string) map[string]string {
	var stdout string
	var stderr string
	if language == "clang" {
		// Compile
		cmd := []string{"clang", "-o", "/tmp/a.out", "/tmp/code.c"}
		_, uuid := lxdexec.ContainerExec(name, cmd)
		lxdexec.Wait(uuid)
		_, stdout, stderr = lxdexec.ContainerGetStd(name, uuid)

		// Run ELF
		if stderr == "" {
			var cmd []string
			if stdin_text == "" {
				cmd = []string{"/tmp/a.out"}
			} else {
				cmd = []string{"bash", "-c", "echo -e '" + stdin_text + "' | /tmp/a.out"}
			}
			_, uuid := lxdexec.ContainerExec(name, cmd)
			lxdexec.Wait(uuid)
			_, stdout, stderr = lxdexec.ContainerGetStd(name, uuid)
		}
	}
	return map[string]string{"stdout": stdout, "stderr": stderr}
}

func CodePush(name string, code string, extension string) error {
	client, err := lxd.NewClient(&lxd.DefaultConfig, "local")
	if err != nil {
		return err
	}

	extensions := map[string]string{"clang": ".c", "gcc": ".c"}
	f := bytes.NewReader([]byte(code))
	err = client.PushFile(name, "/tmp/code"+extensions[extension], -1, -1, "", f)
	return err
}
