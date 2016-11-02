package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

const samplesCount = 1

var testPackages = []string{
	"github.com/boltdb/bolt/cmd/bolt",
	//"github.com/coreos/etcd",
	//"github.com/gogits/gogs",
	//"github.com/grafana/grafana/pkg/cmd/grafana-server",
	//"github.com/influxdata/influxdb/cmd/influxd",
	"github.com/junegunn/fzf/src/fzf",
	//"github.com/mholt/caddy/caddy",
	//"github.com/monochromegane/the_platinum_searcher/cmd/pt",
	//"github.com/nsqio/nsq/apps/nsqd",
	//"github.com/prometheus/prometheus/cmd/prometheus",
	//"github.com/spf13/hugo",
	//"golang.org/x/tools/cmd/guru",
}

const benchmark = "github.com/alecthomas/go_serialization_benchmarks"

const defaultReportName = "report.md"

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

	commits, err := gr.getGitLog(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalf("get git log: %v", err)
	}

	log.Printf("Switch go to %s", os.Args[1])
	oldDur, err := gr.switchRevision(os.Args[1])
	if err != nil {
		log.Fatalf("Go switch revision: %v", err)
	}
	oldResults, err := getResult(gp, testPackages...)
	if err != nil {
		log.Fatalf("build report: %v", err)
	}

	//oldBench, err := gp.RunBenchmark(benchmark)
	//if err != nil {
	//log.Fatalf("benchmark run: %v", err)
	//}

	gp.CleanPkg()

	log.Printf("Switch go to %s", os.Args[2])
	newDur, err := gr.switchRevision(os.Args[2])
	if err != nil {
		log.Fatalf("Go switch revision: %v", err)
	}
	newResults, err := getResult(gp, testPackages...)
	if err != nil {
		log.Fatalf("build report: %v", err)
	}
	//newBench, err := gp.RunBenchmark(benchmark)
	//if err != nil {
	//log.Fatalf("benchmark run: %v", err)
	//}

	rep := report{
		gp:               gp,
		commits:          commits,
		oldGoCompileTime: oldDur,
		newGoCompileTime: newDur,
		oldResults:       oldResults,
		newResults:       newResults,
		//oldBenchmark:     oldBench,
		//newBenchmark:     newBench,
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
