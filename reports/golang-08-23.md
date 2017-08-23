# August 23, 2017 Report

Number of commits: 303

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.365659349s to 6.317265843s, -0.76%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.806365114s to 20.54922543s, -1.24%
* github.com/prometheus/prometheus/cmd/prometheus: from 15.309195167s to 15.324423488s, ~
* code.gitea.io/gitea: from 7.930903576s to 7.641791653s, -3.65%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29607512 to 29594360, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38405744 to 38399216, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 60971147 to 60903120, -0.11%
* code.gitea.io/gitea: from 31721052 to 31632687, -0.28%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               147           149           +1.36%
BenchmarkMsgpUnmarshal-8             278           316           +13.67%
BenchmarkEasyJsonMarshal-8           1089          1106          +1.56%
BenchmarkEasyJsonUnmarshal-8         1083          1172          +8.22%
BenchmarkFlatBuffersMarshal-8        291           298           +2.41%
BenchmarkFlatBuffersUnmarshal-8      219           231           +5.48%
BenchmarkGogoprotobufMarshal-8       122           121           -0.82%
BenchmarkGogoprotobufUnmarshal-8     185           185           +0.00%

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

* [crypto/rand: batch large calls to linux getrandom](https://github.com/golang/go/commit/d401c427b29f48d5cbc5092e62c20aa8524ce356)
* [archive/tar: optimize formatPAXRecord() call](https://github.com/golang/go/commit/23cd87eb0a2d49a3208824feaf34d8b852da422f)
* [cmd/go: parallelize fmt](https://github.com/golang/go/commit/1f631a2f9a20d8dc57fb877fb95f807c895d1c40)
* [runtime: simplify memory capacity check in growslice](https://github.com/golang/go/commit/365594ad59873cd8f7fde5ec158067bf695185ee)
* [cmd/compile: combine x*n + y*n into (x+y)*n](https://github.com/golang/go/commit/a0453a180fc3555843185385e9d4ad9d57f1d36a)
* [runtime: improve makechan memory checks and allocation calls](https://github.com/golang/go/commit/455775dae63eb1277227cbde9e99dc67a3fdb0ea)
* [cmd/compile: pass stack allocated bucket to makemap inside hmap](https://github.com/golang/go/commit/06a78b57377ce63c7fca968af5056a3dec0a06bb)
* [strconv: check bitsize range in ParseInt and ParseUint](https://github.com/golang/go/commit/63c428434692bdeab14115a1f70813feca7795e7)

## GIT Log

```
git log 64bd2c49b4afb80c6f062f52eb3c748970771acf..8f1e2a2610765528068107e33ab0d1d2ff224ce3
```
