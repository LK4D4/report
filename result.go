package main

import "time"

type result struct {
	Name        string
	TimeSamples []time.Duration
	BinarySize  int64
}
