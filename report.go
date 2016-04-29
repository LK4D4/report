package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"time"
)

type report struct {
	gp               gopath
	oldGoCompileTime time.Duration
	newGoCompileTime time.Duration
	oldResults       []result
	newResults       []result
	oldBenchmark     []byte
	newBenchmark     []byte
}

func getStringPercents(rel float64) string {
	per := rel * 100
	if math.Abs(per) < 0.1 {
		return "~"
	}
	res := fmt.Sprintf("%.2f%%", per)
	if per > 0.0 {
		res = "+" + res
	}
	return res
}

func (r report) writeComparison(w io.Writer) error {
	compare, err := r.gp.BenchCmp(r.oldBenchmark, r.newBenchmark)
	if err != nil {
		return err
	}
	if _, err := w.Write(compare); err != nil {
		return err
	}
	return nil
}

func (r report) Bytes() ([]byte, error) {
	var b bytes.Buffer
	diffs := r.getDiffReport()
	b.WriteString("## Compilation time\n\n")
	for i, d := range diffs.dr {
		line := fmt.Sprintf("* %s: from %v to %v, %s\n", d.name, avgDuration(r.oldResults[i]), avgDuration(r.newResults[i]), getStringPercents(d.compileDiffRel))
		b.WriteString(line)
	}
	b.WriteString("\n")
	b.WriteString("## Binary size:\n\n")
	for i, d := range diffs.dr {
		line := fmt.Sprintf("* %s: from %v to %v, %s\n", d.name, r.oldResults[i].BinarySize, r.newResults[i].BinarySize, getStringPercents(d.sizeDiffRel))
		b.WriteString(line)
	}
	b.WriteString("\n")
	b.WriteString("## Bechmarks:\n\n")
	b.WriteString("```\n")
	if err := r.writeComparison(&b); err != nil {
		return nil, err
	}
	b.WriteString("```")
	b.WriteString("\n")
	b.WriteString("## Highlights: \n\n")
	b.WriteString("\t<-------------------HIGHLIGHTS HERE---------------------->\n")
	b.WriteString("\n")
	return b.Bytes(), nil
}

type diffResult struct {
	name           string
	compileDiffAbs time.Duration
	compileDiffRel float64
	sizeDiffAbs    int64
	sizeDiffRel    float64
}

func avgDuration(r result) time.Duration {
	var sum time.Duration
	for _, d := range r.TimeSamples {
		sum += d
	}
	return sum / time.Duration(len(r.TimeSamples))
}

func diff(old, new result) diffResult {
	if old.Name != new.Name {
		panic("different packages")
	}
	oldAvgCompile := avgDuration(old)
	newAvgCompile := avgDuration(new)
	absCdiff := newAvgCompile - oldAvgCompile
	relCdiff := float64(absCdiff) / float64(oldAvgCompile)
	sizeDiffAbs := new.BinarySize - old.BinarySize
	sizeDiffRel := float64(sizeDiffAbs) / float64(old.BinarySize)
	return diffResult{
		name:           old.Name,
		compileDiffAbs: absCdiff,
		compileDiffRel: relCdiff,
		sizeDiffAbs:    sizeDiffAbs,
		sizeDiffRel:    sizeDiffRel,
	}
}

type diffReport struct {
	dr []diffResult
}

func (r report) getDiffReport() diffReport {
	var d diffReport
	for i, or := range r.oldResults {
		d.dr = append(d.dr, diff(or, r.newResults[i]))
	}
	return d
}
