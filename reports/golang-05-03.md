# May 3, 2017 Report

Number of commits: 116

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.984987651s to 6.940754247s, -0.63%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.901351369s to 20.867653921s, -0.16%
* github.com/prometheus/prometheus/cmd/prometheus: from 19.993288473s to 20.166363846s, +0.87%
* code.gitea.io/gitea: from 8.51736895s to 8.427515671s, -1.05%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29214032 to 29228016, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38464216 to 38475112, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 64041132 to 64068247, ~
* code.gitea.io/gitea: from 32206689 to 32209230, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-8               151           148           -1.99%
BenchmarkMsgpUnmarshal-8             277           275           -0.72%
BenchmarkEasyJsonMarshal-8           1119          1112          -0.63%
BenchmarkEasyJsonUnmarshal-8         1124          1111          -1.16%
BenchmarkFlatBuffersMarshal-8        291           292           +0.34%
BenchmarkFlatBuffersUnmarshal-8      226           223           -1.33%
BenchmarkGogoprotobufMarshal-8       146           125           -14.38%
BenchmarkGogoprotobufUnmarshal-8     299           195           -34.78%

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
* [runtime: align mcentral by cache line size](https://github.com/golang/go/commit/259d60995d735523fc25939c35847538eb0d067)
* [context: define behavior for Err before Done is closed](https://github.com/golang/go/commit/6e2c4bc012f8cc262db25d3fee414c5231fea03a)
* [testing: add argument to list tests, benchmarks, and examples](https://github.com/golang/go/commit/ba8ff87dbeb87813a4604e36adb609b1e8fcb7be)
* [sync: import Map from x/sync/syncmap](https://github.com/golang/go/commit/959025c0ac97ec3533ef9f3f70d64453352a7b56)
    * [reflect: use sync.Map instead of RWMutex for type caches](https://github.com/golang/go/commit/33b92cd6ce46c61a4d00a86b88971534773dd4a8)
    * [encoding/gob: replace RWMutex usage with sync.Map](https://github.com/golang/go/commit/c120e449fbc618f9510387d718de0cef5f73af3a)
    * [encoding/xml: replace tinfoMap RWMutex with sync.Map](https://github.com/golang/go/commit/eb6adc27d56687970dd8a49794ca85acc4cf9097)
    * [encoding/json: replace encoderCache RWMutex with a sync.Map](https://github.com/golang/go/commit/d6ce7e4feca75d2833f0790260ea46e194c55170)
    * [mime: use sync.Map instead of RWMutex for type lookups](https://github.com/golang/go/commit/e8d7e5d1fa7d8477b91cb4dffeac57c7c20cb5c5)
* [cmd/compile: add initial backend concurrency support](https://github.com/golang/go/commit/756b9ce3a5a518555114b0e023eb1674084b38e1)

## GIT Log

```
git log e3d7ec006f25385972c89f771d5d577adce3f024..6e9b6e1d222a4f8ad3d50929ee1d6178fb3c6077
```
