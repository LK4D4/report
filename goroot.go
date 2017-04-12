package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type goroot struct {
	path string
}

func newGoroot() goroot {
	return goroot{path: runtime.GOROOT()}
}

func (gr goroot) switchRevision(rev string) (time.Duration, error) {
	checkoutCmd := exec.Command("git", "checkout", rev)
	checkoutCmd.Dir = gr.path
	out, err := checkoutCmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("git checkout: %v, out: %s", err, out)
	}
	cleanCmd := exec.Command("git", "clean", "-f")
	cleanCmd.Dir = gr.path
	out, err = cleanCmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("git clean: %v, out: %s", err, out)
	}
	log.Printf("compile go at %s", rev)
	compileCmd := exec.Command("./make.bash")
	compileCmd.Dir = filepath.Join(gr.path, "src")
	start := time.Now()
	out, err = compileCmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("./make.bash: %v, out: %s", err, out)
	}
	dur := time.Since(start)
	log.Printf("go compilation time: %v", dur)
	return dur, nil
}

type gitLog struct {
	N   int
	Cmd string
}

func (gr goroot) getGitLog(old, new string) (gitLog, error) {
	gl := gitLog{}
	cntCmd := exec.Command("git", "rev-list", "--count", old+".."+new)
	cntCmd.Dir = gr.path
	out, err := cntCmd.CombinedOutput()
	if err != nil {
		return gl, fmt.Errorf("git rev-list: %v, out: %s", err, out)
	}
	cnt, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return gl, err
	}
	gl.N = cnt
	gl.Cmd = "git log " + old + ".." + new
	return gl, nil
}
