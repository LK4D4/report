# May 31, 2017 Report

Number of commits: 104

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.005449465s to 5.948581201s, -0.95%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.586031217s to 20.412286416s, -0.84%
* github.com/prometheus/prometheus/cmd/prometheus: from 15.190143317s to 14.958586758s, -1.52%
* code.gitea.io/gitea: from 23.755945697s to 23.500992123s, -1.07%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 28799736 to 28237600, -1.95%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 37881848 to 37276632, -1.60%
* github.com/prometheus/prometheus/cmd/prometheus: from 61362163 to 60328292, -1.68%
* code.gitea.io/gitea: from 36328488 to 35741976, -1.61%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               146           146           +0.00%
BenchmarkMsgpUnmarshal-8             276           274           -0.72%
BenchmarkEasyJsonMarshal-8           1107          1106          -0.09%
BenchmarkEasyJsonUnmarshal-8         1125          1081          -3.91%
BenchmarkFlatBuffersMarshal-8        323           294           -8.98%
BenchmarkFlatBuffersUnmarshal-8      220           218           -0.91%
BenchmarkGogoprotobufMarshal-8       123           122           -0.81%
BenchmarkGogoprotobufUnmarshal-8     192           188           -2.08%

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

* [syscall: add Conn and RawConn interfaces](https://github.com/golang/go/commit/de5c573baabf925ee7cb868285ed4f14de5f7fe9)
* [cmd/go: Document that -cover causes incorrect line numbers](https://github.com/golang/go/commit/f6f1daa4b807f1164da67c96dc8b07ee84731b52)
* [crypto/rand: use blocking getrandom call on Linux when supported](https://github.com/golang/go/commit/95d991d30c59edc4943bd8baf5c664c5f8b1cebe)
* [strings: simplify indexFunc](https://github.com/golang/go/commit/a5083bbf0784a4664e49207b6a3677986ca69b49)


## GIT Log

```
git log c20e54533ea49ca68640d9a59c9ed935b27da8e5..1e0819101b476f807bf8ef3fd50f1ee26691f33e
```
