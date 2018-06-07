# July 5, 2017 Report

Number of commits: 282

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.22480492s to 6.123507841s, -1.63%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.86464471s to 20.988301391s, +0.59%
* github.com/prometheus/prometheus/cmd/prometheus: from 15.479684486s to 15.323795279s, -1.01%
* code.gitea.io/gitea: from 24.057353423s to 23.725619051s, -1.38%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 28896080 to 28892104, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 37782096 to 37793744, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 60959013 to 60940456, ~
* code.gitea.io/gitea: from 36143472 to 36161864, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               146           147           +0.68%
BenchmarkMsgpUnmarshal-8             274           302           +10.22%
BenchmarkEasyJsonMarshal-8           1104          1092          -1.09%
BenchmarkEasyJsonUnmarshal-8         1100          1074          -2.36%
BenchmarkFlatBuffersMarshal-8        295           293           -0.68%
BenchmarkFlatBuffersUnmarshal-8      218           220           +0.92%
BenchmarkGogoprotobufMarshal-8       123           135           +9.76%
BenchmarkGogoprotobufUnmarshal-8     189           189           +0.00%

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

* [cmd/compile/internal/gc: speed-up small array comparison](https://github.com/golang/go/commit/3bdc2f3abf0f9cffc8f4e294ef22a23b82e88415)
* [runtime: avoid division in gc](https://github.com/golang/go/commit/a4ee95c805fb77e594603bcd62d7858dc9e853ab)
* [cmd/link: fix accidentally-quadratic library loading](https://github.com/golang/go/commit/51711d1429cb592c9ddc772e6362e74ac8545dc8)
* [doc: add qualified mention of dep to FAQ](https://github.com/golang/go/commit/dc8b4e65a7a68e102484020efbf80cecd2d515bd)
* [syscall: use CLONE_VFORK safely](https://github.com/golang/go/commit/67e537541c043c701001f002bed0cda70ce72767)
* [runtime, syscall: workaround for bug in Linux's execve](https://github.com/golang/go/commit/91139b87f776a553524b022753981e7909386777)
* [os/signal: avoid race between Stop and receiving on channel](https://github.com/golang/go/commit/8ec7a39fec2acab98ce5e41363dd1c65c03d7479)
* [doc, api: add syscall.SysProcAttr.AmbientCaps change to 1.9 notes, API](https://github.com/golang/go/commit/dc86c9a6afa8b5b998dfa6621d1566d1296f2bf4)

## GIT Log

```
git log 1e0819101b476f807bf8ef3fd50f1ee26691f33e..6c4f3a0c16e5da3caa08cb8f368dc7db90bb211d
```
