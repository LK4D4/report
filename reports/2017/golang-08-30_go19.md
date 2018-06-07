# August 29, 2017 Report

Number of commits: 2190

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 8.904411832s to 8.355669793s, -6.16%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 24.155850826s to 24.362854437s, +0.86%
* github.com/prometheus/prometheus/cmd/prometheus: from 23.605404382s to 21.868107331s, -7.36%
* code.gitea.io/gitea: from 10.53319879s to 9.881360545s, -6.19%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29906984 to 29612344, -0.99%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38789616 to 38418360, -0.96%
* github.com/prometheus/prometheus/cmd/prometheus: from 62014986 to 60971320, -1.68%
* code.gitea.io/gitea: from 31819138 to 31773325, -0.14%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               168           158           -5.95%
BenchmarkMsgpUnmarshal-8             339           320           -5.60%
BenchmarkEasyJsonMarshal-8           1319          1282          -2.81%
BenchmarkEasyJsonUnmarshal-8         1325          1178          -11.09%
BenchmarkFlatBuffersMarshal-8        340           304           -10.59%
BenchmarkFlatBuffersUnmarshal-8      284           235           -17.25%
BenchmarkGogoprotobufMarshal-8       146           137           -6.16%
BenchmarkGogoprotobufUnmarshal-8     225           211           -6.22%

benchmark                            old allocs     new allocs     delta
BenchmarkMsgpMarshal-8               1              1              +0.00%
BenchmarkMsgpUnmarshal-8             3              3              +0.00%
BenchmarkEasyJsonMarshal-8           5              5              +0.00%
BenchmarkEasyJsonUnmarshal-8         4              4              +0.00%
BenchmarkFlatBuffersMarshal-8        0              0              +0.00%
BenchmarkFlatBuffersUnmarshal-8      3              3              +0.00%
BenchmarkGogoprotobufMarshal-8       1              1              +0.00%
BenchmarkGogoprotobufUnmarshal-8     3              3              +0.00%

benchmark                            old bytes     new bytes     delta
BenchmarkMsgpMarshal-8               128           128           +0.00%
BenchmarkMsgpUnmarshal-8             112           112           +0.00%
BenchmarkEasyJsonMarshal-8           784           784           +0.00%
BenchmarkEasyJsonUnmarshal-8         160           160           +0.00%
BenchmarkFlatBuffersMarshal-8        0             0             +0.00%
BenchmarkFlatBuffersUnmarshal-8      112           112           +0.00%
BenchmarkGogoprotobufMarshal-8       64            64            +0.00%
BenchmarkGogoprotobufUnmarshal-8     96            96            +0.00%
```
## Highlights: 

https://golang.org/doc/go1.9

## GIT Log

```
git log go1.8.3..go1.9
```
