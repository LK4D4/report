# May 16, 2017 Report

Number of commits: 52

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.131619599s to 6.238794828s, +1.75%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.823279563s to 21.130564939s, +1.48%
* github.com/prometheus/prometheus/cmd/prometheus: from 15.749732993s to 15.624541063s, -0.79%
* code.gitea.io/gitea: from 7.844720914s to 7.782880927s, -0.79%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29289000 to 28790552, -1.70%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38491720 to 37865736, -1.63%
* github.com/prometheus/prometheus/cmd/prometheus: from 62515718 to 61312214, -1.93%
* code.gitea.io/gitea: from 32284773 to 31601419, -2.12%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               148           146           -1.35%
BenchmarkMsgpUnmarshal-8             276           276           +0.00%
BenchmarkEasyJsonMarshal-8           1115          1106          -0.81%
BenchmarkEasyJsonUnmarshal-8         1126          1124          -0.18%
BenchmarkFlatBuffersMarshal-8        311           321           +3.22%
BenchmarkFlatBuffersUnmarshal-8      241           220           -8.71%
BenchmarkGogoprotobufMarshal-8       125           123           -1.60%
BenchmarkGogoprotobufUnmarshal-8     191           193           +1.05%

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

* [internal/cpu: new package to detect cpu features](https://github.com/golang/go/commit/69972aea74de6a0397a05281475d1ca006da7bb0)
* [net: allow Resolver to use a custom dialer](https://github.com/golang/go/commit/380aa884b8b2935137eee266d0a44e9083fae71f)
* [container/heap: avoid up() invoke if down() success at heap.Remove()](https://github.com/golang/go/commit/ee57e36dfa6879c05ac6717c29f2df5b546e1256)
* [cmd/compile: eliminate some bounds checks from generated rewrite rules](https://github.com/golang/go/commit/5548f7d5cfadf1319a99b2a17e72ff91fcdd09fc)

## GIT Log

```
git log 266a3b66ca1f49463a29f047d6fde62eb18025b8..c20e54533ea49ca68640d9a59c9ed935b27da8e5
```
