# August 18, 2016 Report - 1.7 Release Report

Number of commits: 2618

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 896.422603ms to 591.938888ms, -33.97%
* github.com/coreos/etcd: from 19.972392715s to 13.720111409s, -31.30%
* github.com/gogits/gogs: from 21.321361237s to 13.765160253s, -35.44%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 41.857540153s to 35.21983685s, -15.86%
* github.com/influxdata/influxdb/cmd/influxd: from 9.315015559s to 6.747189682s, -27.57%
* github.com/junegunn/fzf/src/fzf: from 1.502196146s to 1.028185928s, -31.55%
* github.com/mholt/caddy/caddy: from 10.547308591s to 6.205710907s, -41.16%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 2.244452542s to 1.611611546s, -28.20%
* github.com/nsqio/nsq/apps/nsqd: from 3.549434917s to 2.212453609s, -37.67%
* github.com/prometheus/prometheus/cmd/prometheus: from 15.847923625s to 9.975685522s, -37.05%
* github.com/spf13/hugo: from 11.062636479s to 7.876749451s, -28.80%
* golang.org/x/tools/cmd/guru: from 3.824423257s to 2.780342645s, -27.30%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 3502960 to 2679372, -23.51%
* github.com/coreos/etcd: from 30182432 to 24570808, -18.59%
* github.com/gogits/gogs: from 29320424 to 24165738, -17.58%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 32969216 to 27705288, -15.97%
* github.com/influxdata/influxdb/cmd/influxd: from 20228464 to 16692722, -17.48%
* github.com/junegunn/fzf/src/fzf: from 4097752 to 3252216, -20.63%
* github.com/mholt/caddy/caddy: from 18374240 to 15152137, -17.54%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 5542200 to 4598435, -17.03%
* github.com/nsqio/nsq/apps/nsqd: from 12011912 to 10012284, -16.65%
* github.com/prometheus/prometheus/cmd/prometheus: from 28141728 to 23119508, -17.85%
* github.com/spf13/hugo: from 19397912 to 15885876, -18.11%
* golang.org/x/tools/cmd/guru: from 9571344 to 7851170, -17.97%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               284           188           -33.80%
BenchmarkMsgpUnmarshal-4             502           405           -19.32%
BenchmarkEasyJsonMarshal-4           1926          1615          -16.15%
BenchmarkEasyJsonUnmarshal-4         2450          2107          -14.00%
BenchmarkFlatBuffersMarshal-4        835           363           -56.53%
BenchmarkFlatBuffersUnmarshal-4      361           291           -19.39%
BenchmarkGogoprotobufMarshal-4       210           162           -22.86%
BenchmarkGogoprotobufUnmarshal-4     309           254           -17.80%

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

* [Go 1.7 Release Notes](https://golang.org/doc/go1.7)
