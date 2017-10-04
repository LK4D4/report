# October 3, 2017 Report

Number of commits: 224

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.320858819s to 6.274773921s, -0.73%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 21.377521367s to 21.514336434s, +0.64%
* github.com/prometheus/prometheus/cmd/prometheus: from 15.088287443s to 14.9311496s, -1.04%
* code.gitea.io/gitea: from 7.856976891s to 7.766483044s, -1.15%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29788192 to 29920392, +0.44%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 41169528 to 41271096, +0.25%
* github.com/prometheus/prometheus/cmd/prometheus: from 60531827 to 60647769, +0.19%
* code.gitea.io/gitea: from 32105074 to 32271354, +0.52%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               145           143           -1.38%
BenchmarkMsgpUnmarshal-8             274           267           -2.55%
BenchmarkEasyJsonMarshal-8           1090          1093          +0.28%
BenchmarkEasyJsonUnmarshal-8         1088          1077          -1.01%
BenchmarkFlatBuffersMarshal-8        287           288           +0.35%
BenchmarkFlatBuffersUnmarshal-8      214           206           -3.74%
BenchmarkGogoprotobufMarshal-8       122           120           -1.64%
BenchmarkGogoprotobufUnmarshal-8     179           174           -2.79%

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

* [runtime: improve fastrand with a better generator](https://github.com/golang/go/commit/e7e4a4ffa3330518250c4075e1f16a8ba62414df)
* [io: Improve performance of CopyN](https://github.com/golang/go/commit/098eb01600fe0e90aee21d204924c67188fe26d4)
* [cmd/compile/internal/gc: better inliner diagnostics](https://github.com/golang/go/commit/475df0ebccaf0871c86b2c0b55ee841aede324b7)
* [runtime: don't call lockOSThread for every cgo call](https://github.com/golang/go/commit/332719f7cee2abafb3963009d44ad7cc93474707)
* [regexp: make (*bitState).push inlinable](https://github.com/golang/go/commit/1ae67965e47a8a8eb71c92e44134c89cd1c67657)

## GIT Log

```
git log de2231888821add783305e7674bbb43d4d8453dc..273b657b4e970f510afb258aa73dc2e264a701e3
```
