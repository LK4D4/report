# February 16, 2017 Report (Release 1.8)

Number of commits: 2301

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 608.221933ms to 607.683564ms, ~
* github.com/coreos/etcd/cmd/etcd: from 14.493261174s to 12.574284838s, -13.24%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 36.757500578s to 36.256616348s, -1.36%
* github.com/junegunn/fzf/src/fzf: from 1.024158847s to 1.087723734s, +6.21%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.422330519s to 1.2497697s, -12.13%
* github.com/nsqio/nsq/apps/nsqd: from 2.096109458s to 1.960299521s, -6.48%
* github.com/prometheus/prometheus/cmd/prometheus: from 47.221796519s to 41.124079111s, -12.91%
* github.com/spf13/hugo: from 7.479409536s to 6.44479598s, -13.83%
* golang.org/x/tools/cmd/guru: from 2.883006747s to 2.658148815s, -7.80%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2696808 to 3188113, +18.22%
* github.com/coreos/etcd/cmd/etcd: from 27756416 to 26987880, -2.77%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 35954896 to 34414776, -4.28%
* github.com/junegunn/fzf/src/fzf: from 3362675 to 3208398, -4.59%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4615986 to 4362805, -5.48%
* github.com/nsqio/nsq/apps/nsqd: from 9982014 to 10345037, +3.64%
* github.com/prometheus/prometheus/cmd/prometheus: from 67052893 to 62571775, -6.68%
* github.com/spf13/hugo: from 17062116 to 16789673, -1.60%
* golang.org/x/tools/cmd/guru: from 8279573 to 8234173, -0.55%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               193           189           -2.07%
BenchmarkMsgpUnmarshal-4             408           368           -9.80%
BenchmarkEasyJsonMarshal-4           1601          1436          -10.31%
BenchmarkEasyJsonUnmarshal-4         1926          1987          +3.17%
BenchmarkFlatBuffersMarshal-4        539           371           -31.17%
BenchmarkFlatBuffersUnmarshal-4      295           293           -0.68%
BenchmarkGogoprotobufMarshal-4       163           156           -4.29%
BenchmarkGogoprotobufUnmarshal-4     261           253           -3.07%

benchmark                            old allocs     new allocs     delta
BenchmarkMsgpMarshal-4               1              1              +0.00%
BenchmarkMsgpUnmarshal-4             3              3              +0.00%
BenchmarkEasyJsonMarshal-4           5              5              +0.00%
BenchmarkEasyJsonUnmarshal-4         4              4              +0.00%
BenchmarkFlatBuffersMarshal-4        0              0              +0.00%
BenchmarkFlatBuffersUnmarshal-4      3              3              +0.00%
BenchmarkGogoprotobufMarshal-4       1              1              +0.00%
BenchmarkGogoprotobufUnmarshal-4     3              3              +0.00%

benchmark                            old bytes     new bytes     delta
BenchmarkMsgpMarshal-4               128           128           +0.00%
BenchmarkMsgpUnmarshal-4             112           112           +0.00%
BenchmarkEasyJsonMarshal-4           784           784           +0.00%
BenchmarkEasyJsonUnmarshal-4         160           160           +0.00%
BenchmarkFlatBuffersMarshal-4        0             0             +0.00%
BenchmarkFlatBuffersUnmarshal-4      112           112           +0.00%
BenchmarkGogoprotobufMarshal-4       64            64            +0.00%
BenchmarkGogoprotobufUnmarshal-4     96            96            +0.00%
```
## Highlights: 

* [release notes](https://golang.org/doc/go1.8)
