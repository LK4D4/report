# April 25, 2017 Report

Number of commits: 165

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.842432634s to 6.916159253s, +1.08%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.931030524s to 20.901291831s, -0.14%
* github.com/prometheus/prometheus/cmd/prometheus: from 20.452920998s to 20.308057577s, -0.71%
* code.gitea.io/gitea: from 8.32561586s to 8.291286123s, -0.41%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29175856 to 29197488, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38427152 to 38462344, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 63869473 to 63914633, ~
* code.gitea.io/gitea: from 31724422 to 31766853, +0.13%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               150           151           +0.67%
BenchmarkMsgpUnmarshal-8             280           278           -0.71%
BenchmarkEasyJsonMarshal-8           1129          1119          -0.89%
BenchmarkEasyJsonUnmarshal-8         1112          1112          +0.00%
BenchmarkFlatBuffersMarshal-8        291           291           +0.00%
BenchmarkFlatBuffersUnmarshal-8      226           227           +0.44%
BenchmarkGogoprotobufMarshal-8       129           196           +51.94%
BenchmarkGogoprotobufUnmarshal-8     189           214           +13.23%

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

* [runtime: preallocate some overflow buckets](https://github.com/golang/go/commit/94e44a9c8edb64f514b6f3b7f7001db0cfeb2d70)
* [sync: align poolLocal to CPU cache line size](https://github.com/golang/go/commit/8aa31d5dae9644b3e8f6950af58c0cb83e8fc062)
* [os: fix race between file I/O and Close](https://github.com/golang/go/commit/11c7b4491bd2cd1deb7b50433f431be9ced330db)

## GIT Log

```
git log 48163968b2927247213fca7a6f4678d3c93855dc..e3d7ec006f25385972c89f771d5d577adce3f024
```
