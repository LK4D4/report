# May 9, 2017 Report

Number of commits: 57

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 7.077596857s to 6.138483836s, -13.27%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.842339314s to 20.857190835s, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 20.239437793s to 15.980347356s, -21.04%
* code.gitea.io/gitea: from 8.457533141s to 7.819306942s, -7.55%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29265360 to 29271936, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38475112 to 38486208, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 64081040 to 64081366, ~
* code.gitea.io/gitea: from 32250212 to 32258730, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               149           146           -2.01%
BenchmarkMsgpUnmarshal-8             276           276           +0.00%
BenchmarkEasyJsonMarshal-8           1119          1112          -0.63%
BenchmarkEasyJsonUnmarshal-8         1110          1113          +0.27%
BenchmarkFlatBuffersMarshal-8        292           308           +5.48%
BenchmarkFlatBuffersUnmarshal-8      225           240           +6.67%
BenchmarkGogoprotobufMarshal-8       126           124           -1.59%
BenchmarkGogoprotobufUnmarshal-8     197           188           -4.57%

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

* [bytes: optimize Buffer's Write, WriteString, WriteByte, and WriteRune](https://github.com/golang/go/commit/c08ac36761d3dc03d0a0b0ffb240c4a7c524536b)
* [container/heap: optimization when selecting smaller child](https://github.com/golang/go/commit/f5352a7763c8f96f7f092990d64339eae0623263)
* [cmd/go: add support for concurrent backend compilation](https://github.com/golang/go/commit/f4e5bd483b1c6f731c9925d3d1b66d2aba88980e)
* [cmd/go: enable concurrent backend compilation by default](https://github.com/golang/go/commit/5e0bcb3893c2e54fdb96affcafd2953f20dd64eb)
* [cmd/compile: use a buffered channel for the function queue](https://github.com/golang/go/commit/1213776650486aff30b607a6c6b6ece3e9c0155f)


## GIT Log

```
git log 6e9b6e1d222a4f8ad3d50929ee1d6178fb3c6077..266a3b66ca1f49463a29f047d6fde62eb18025b8
```
