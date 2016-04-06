package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var testPackages = []string{
	"github.com/coreos/etcd",
	"github.com/boltdb/bolt/cmd/bolt",
	"github.com/gogits/gogs",
	"github.com/spf13/hugo",
	"github.com/influxdata/influxdb/cmd/influxd",
	"github.com/nsqio/nsq/apps/nsqd",
	"github.com/mholt/caddy",
}

const defaultReportName = "report.txt"

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <start> <end>", os.Args[0])
	}
	start := time.Now()
	gp, err := initGopath(testPackages...)
	if err != nil {
		log.Fatalf("Gopath init: %v", err)
	}
	gr := newGoroot()

	log.Printf("Switch go to %s", os.Args[1])
	oldDur, err := gr.switchRevision(os.Args[1])
	if err != nil {
		log.Fatalf("Go switch revision: %v", err)
	}
	oldResults, err := getResult(gp, testPackages...)
	if err != nil {
		log.Fatalf("build report: %v", err)
	}

	log.Printf("Switch go to %s", os.Args[2])
	newDur, err := gr.switchRevision(os.Args[2])
	if err != nil {
		log.Fatalf("Go switch revision: %v", err)
	}
	newResults, err := getResult(gp, testPackages...)
	if err != nil {
		log.Fatalf("build report: %v", err)
	}
	report := report{
		oldGoCompileTime: oldDur,
		newGoCompileTime: newDur,
		oldResults:       oldResults,
		newResults:       newResults,
	}
	if err := writeReport(defaultReportName, gr, report); err != nil {
		log.Fatalf("write report file: %v", err)
	}
	log.Println("report build: %v", time.Since(start))
}

func writeReport(name string, gr goroot, report report) error {
	gl, err := gr.getGitLog(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalf("get git log: %v", err)
	}
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Number of commits: %d\n", gl.cnt))
	b.WriteString("\n")
	b.Write(report.Bytes())
	b.WriteString("\n")
	b.Write(gl.log)
	return ioutil.WriteFile(name, b.Bytes(), 0666)
}

func initGopath(pkg ...string) (gopath, error) {
	var testGopath string
	if tg := os.Getenv("TEST_GOPATH"); tg != "" {
		testGopath = tg
	} else {
		tg, err := ioutil.TempDir("", "report-gopath-")
		if err != nil {
			return gopath{}, err
		}
		defer os.RemoveAll(tg)
		testGopath = tg
	}
	gp, err := newGopath(testGopath)
	if err != nil {
		return gp, err
	}
	if err := gp.CleanPkg(); err != nil {
		return gp, err
	}
	log.Println(append([]string{"Download:"}, pkg...))
	if err := gp.GoGet(pkg...); err != nil {
		return gp, err
	}
	return gp, nil
}

const samplesCount = 5

func getResult(gp gopath, pkg ...string) ([]result, error) {
	var results []result
	for _, pkg := range testPackages {
		res, err := gp.RunTest(pkg, samplesCount)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}
