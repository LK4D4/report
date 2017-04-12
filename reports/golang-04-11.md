# April 11, 2017 Report

Number of commits: 91

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 1.756231123s to 1.830287273s, +4.22%
* github.com/coreos/etcd/cmd/etcd: from 21.236700801s to 21.068541099s, -0.79%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 1m33.624360428s to 1m41.443431695s, +8.35%
* github.com/junegunn/fzf/src/fzf: from 2.047522313s to 2.233728822s, +9.09%
* github.com/mholt/caddy/caddy: from 11.399929242s to 11.381431322s, -0.16%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 2.466245707s to 2.370034628s, -3.90%
* github.com/nsqio/nsq/apps/nsqd: from 4.000234165s to 4.260929839s, +6.52%
* github.com/prometheus/prometheus/cmd/prometheus: from 1m9.384398736s to 1m15.828466557s, +9.29%
* github.com/spf13/hugo: from 12.448496024s to 12.17841034s, -2.17%
* golang.org/x/tools/cmd/guru: from 6.107231423s to 6.973775622s, +14.19%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2983095 to 2983194, ~
* github.com/coreos/etcd/cmd/etcd: from 29077672 to 29032904, -0.15%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 35955072 to 35901048, -0.15%
* github.com/junegunn/fzf/src/fzf: from 3425864 to 3422042, -0.11%
* github.com/mholt/caddy/caddy: from 16956791 to 16920056, -0.22%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4714139 to 4718509, ~
* github.com/nsqio/nsq/apps/nsqd: from 10847409 to 10827058, -0.19%
* github.com/prometheus/prometheus/cmd/prometheus: from 63880345 to 63757594, -0.19%
* github.com/spf13/hugo: from 18191566 to 18150735, -0.22%
* golang.org/x/tools/cmd/guru: from 8570919 to 8567043, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               375           363           -3.20%
BenchmarkMsgpUnmarshal-4             1135          1019          -10.22%
BenchmarkEasyJsonMarshal-4           4336          3148          -27.40%
BenchmarkEasyJsonUnmarshal-4         6850          5356          -21.81%
BenchmarkFlatBuffersMarshal-4        1561          1411          -9.61%
BenchmarkFlatBuffersUnmarshal-4      685           644           -5.99%
BenchmarkGogoprotobufMarshal-4       595           500           -15.97%
BenchmarkGogoprotobufUnmarshal-4     851           770           -9.52%

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

* [strings: optimize Count for amd64](d206af1e6c53df0c59d9466fe9c50415f9d8dcd5)
* [testing: consider a test failed after race errors](221541ec8c4ec1b0ed0c6f26f5e13ca128e2a3cd)
* [cmd/compile: make iface == iface const evaluation respect !=](b83a916f7186eb98636407c304974db34277aa2f)
	* https://github.com/golang/go/issues/19911

## GIT Log

```
git log 4c1622082e493dea24a936930be8b324aae54505..ab0e9019ea61c1b49572876354af7086f961bc8c
```
