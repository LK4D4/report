# August 8, 2017 Report

Number of commits: 156

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.350105551s to 6.345802079s, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.931790794s to 21.028021607s, +0.46%
* github.com/prometheus/prometheus/cmd/prometheus: from 15.55642627s to 15.278596068s, -1.79%
* code.gitea.io/gitea: from 7.929924679s to 7.922107289s, ~

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29363232 to 29374696, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38358008 to 38370112, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 60953302 to 60957772, ~
* code.gitea.io/gitea: from 31672846 to 31685601, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               145           147           +1.38%
BenchmarkMsgpUnmarshal-8             278           275           -1.08%
BenchmarkEasyJsonMarshal-8           1098          1088          -0.91%
BenchmarkEasyJsonUnmarshal-8         1094          1084          -0.91%
BenchmarkFlatBuffersMarshal-8        295           290           -1.69%
BenchmarkFlatBuffersUnmarshal-8      219           217           -0.91%
BenchmarkGogoprotobufMarshal-8       125           124           -0.80%
BenchmarkGogoprotobufUnmarshal-8     218           182           -16.51%

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

## GIT Log

```
git log 6c4f3a0c16e5da3caa08cb8f368dc7db90bb211d..64bd2c49b4afb80c6f062f52eb3c748970771acf
```
