package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	args := append([]string{"get", "-insecure", "-u", "-d", "-t"}, pkg...)
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

func (gp gopath) RunBenchmark(pkg string) ([]byte, error) {
	log.Printf("Run Benchmarks %s", pkg)
	if err := gp.GoGet(pkg); err != nil {
		return nil, err
	}
	cmd := exec.Command("go", "test", "-bench", "(^BenchmarkMsgp)|EasyJson|Gogoprotobuf|Flat", "-benchmem", pkg)
	gp.setGopath(cmd)
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("go test -bench: %v, out: %s", err, stderr.String())
	}
	return stdout.Bytes(), nil
}

func (gp gopath) installBenchCmp() error {
	cmd := exec.Command("go", "get", "golang.org/x/tools/cmd/benchcmp")
	gp.setGopath(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go get: %v, out: %s", err, out)
	}
	return nil
}

func (gp gopath) BenchCmp(old []byte, new []byte) ([]byte, error) {
	if err := gp.installBenchCmp(); err != nil {
		return nil, err
	}
	oldFile, err := ioutil.TempFile("", "report-bench-")
	if err != nil {
		return nil, err
	}
	newFile, err := ioutil.TempFile("", "report-bench-")
	if err != nil {
		return nil, err
	}
	if _, err := oldFile.Write(old); err != nil {
		return nil, err
	}
	oldFile.Close()
	if _, err := newFile.Write(new); err != nil {
		return nil, err
	}
	newFile.Close()

	cmd := exec.Command(filepath.Join(gp.path, "bin", "benchcmp"), oldFile.Name(), newFile.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("benchcmp: %v, out: %s", err, out)
	}
	return out, nil
}
