package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const samplesCount = 1

var testPackages = []string{
	"github.com/coreos/etcd/cmd/etcd",
	"github.com/grafana/grafana/pkg/cmd/grafana-server",
	"github.com/prometheus/prometheus/cmd/prometheus",
	"code.gitea.io/gitea",
}

const benchmark = "github.com/alecthomas/go_serialization_benchmarks"

const defaultReportName = "report.md"

func main() {
	revStart := flag.String("start", "",
		`GOROOT git commit hash for "old" Go version`)
	revEnd := flag.String("end", "master",
		`GOROOT git commit hash for "new" Go version`)
	flag.Parse()
	if *revStart == "" || *revEnd == "" {
		log.Fatalf("Empty start/end revision")
	}

	start := time.Now()
	gp, err := initGopath(testPackages...)
	if err != nil {
		log.Fatalf("Gopath init: %v", err)
	}
	gr := newGoroot()

	commits, err := gr.getGitLog(*revStart, *revEnd)
	if err != nil {
		log.Fatalf("get git log: %v", err)
	}

	log.Printf("Switch go to %s", *revStart)
	oldDur, err := gr.switchRevision(*revStart)
	if err != nil {
		log.Fatalf("Go switch revision: %v", err)
	}
	oldResults, err := getResult(gp, testPackages...)
	if err != nil {
		log.Fatalf("build report: %v", err)
	}

	oldBench, err := gp.RunBenchmark(benchmark)
	if err != nil {
		log.Fatalf("benchmark run: %v", err)
	}

	gp.CleanPkg()

	log.Printf("Switch go to %s", *revEnd)
	newDur, err := gr.switchRevision(*revEnd)
	if err != nil {
		log.Fatalf("Go switch revision: %v", err)
	}
	newResults, err := getResult(gp, testPackages...)
	if err != nil {
		log.Fatalf("build report: %v", err)
	}
	newBench, err := gp.RunBenchmark(benchmark)
	if err != nil {
		log.Fatalf("benchmark run: %v", err)
	}

	rep := report{
		gp:               gp,
		commits:          commits,
		oldGoCompileTime: oldDur,
		newGoCompileTime: newDur,
		oldResults:       oldResults,
		newResults:       newResults,
		oldBenchmark:     oldBench,
		newBenchmark:     newBench,
	}
	if err := writeReport(defaultReportName, rep); err != nil {
		log.Fatalf("write report file: %v", err)
	}
	log.Println("report build:", time.Since(start))
}

func writeReport(name string, rep report) error {
	repBts, err := rep.Bytes()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name, repBts, 0666)
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

func getResult(gp gopath, pkgs ...string) ([]result, error) {
	var results []result
	for _, pkg := range pkgs {
		res, err := gp.RunTest(pkg, samplesCount)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}
