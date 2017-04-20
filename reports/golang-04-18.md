# April 18, 2017 Report

Number of commits: 110

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.948306554s to 6.949451634s, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.945806988s to 21.000798416s, +0.26%
* github.com/prometheus/prometheus/cmd/prometheus: from 20.314839289s to 20.225088531s, -0.44%
* code.gitea.io/gitea: from 7.992406049s to 7.96261295s, -0.37%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29097296 to 29173768, +0.26%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38211296 to 38324360, +0.30%
* github.com/prometheus/prometheus/cmd/prometheus: from 63757674 to 63869473, +0.18%
* code.gitea.io/gitea: from 31469419 to 31581226, +0.36%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               151           151           +0.00%
BenchmarkMsgpUnmarshal-8             276           281           +1.81%
BenchmarkEasyJsonMarshal-8           1126          1128          +0.18%
BenchmarkEasyJsonUnmarshal-8         1105          1110          +0.45%
BenchmarkFlatBuffersMarshal-8        293           292           -0.34%
BenchmarkFlatBuffersUnmarshal-8      225           227           +0.89%
BenchmarkGogoprotobufMarshal-8       128           130           +1.56%
BenchmarkGogoprotobufUnmarshal-8     192           192           +0.00%

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

* [testing: add TB.Helper to better support test helpers](bc2931372243043842161c0a60bd2f86ef9696ee)
* [sync: improve Pool performance](af5c95117b26e22d942a12e15bdc8e25607f738c)

## GIT Log

```
git log ab0e9019ea61c1b49572876354af7086f961bc8c..48163968b2927247213fca7a6f4678d3c93855dc
```
