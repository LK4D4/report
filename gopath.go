package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type gopath struct {
	path string
}

func newGopath(dir string) (gopath, error) {
	return gopath{path: dir}, os.MkdirAll(filepath.Join(dir, "src"), 0700)
}

func (gp gopath) setGopath(cmd *exec.Cmd) {
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "GOPATH=") {
			continue
		}
		cmd.Env = append(cmd.Env, env)
	}
	cmd.Env = append(cmd.Env, "GOPATH="+gp.path)
}

func (gp gopath) GoGet(pkg ...string) error {
	args := append([]string{"get", "-d"}, pkg...)
	cmd := exec.Command("go", args...)
	gp.setGopath(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go get: %v, out: %s", err, out)
	}
	return nil
}

func (gp gopath) GoInstall(pkg string) error {
	cmd := exec.Command("go", "install", pkg)
	gp.setGopath(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go install: %v, out: %s", err, out)
	}
	return nil
}

func (gp gopath) CleanPkg() error {
	return os.RemoveAll(filepath.Join(gp.path, "pkg"))
}

func (gp gopath) RunTest(pkg string, samples int) (result, error) {
	r := result{
		Name: pkg,
	}
	log.Printf("Run tests %s", pkg)
	for i := 1; i <= samples; i++ {
		start := time.Now()
		if err := gp.GoInstall(pkg); err != nil {
			return r, err
		}
		dur := time.Since(start)
		r.TimeSamples = append(r.TimeSamples, dur)
		if r.BinarySize == 0 {
			fi, err := os.Stat(filepath.Join(gp.path, "bin", filepath.Base(pkg)))
			if err != nil {
				return r, err
			}
			r.BinarySize = fi.Size()
			log.Printf("\tBinary size: %d", r.BinarySize)
		}
		log.Printf("\tRun %d, compilation time: %v", i, dur)
		gp.CleanPkg()
	}
	return r, nil
}
