# April 27 Report

Number of commits: 128

## Compilation time

* github.com/coreos/etcd: from 12.518253016s to 11.040618367s, -11.80%
* github.com/boltdb/bolt/cmd/bolt: from 619.12847ms to 541.388933ms, -12.56%
* github.com/gogits/gogs: from 15.174854722s to 12.569516615s, -17.17%
* github.com/spf13/hugo: from 7.488501683s to 6.542196366s, -12.64%
* github.com/influxdata/influxdb/cmd/influxd: from 7.070586468s to 6.155792797s, -12.94%
* github.com/nsqio/nsq/apps/nsqd: from 2.46299607s to 2.089137538s, -15.18%
* github.com/mholt/caddy: from 5.511706049s to 4.681733959s, -15.06%

## Binary size

* github.com/coreos/etcd: from 21522088 to 21090104, -2.01%
* github.com/boltdb/bolt/cmd/bolt: from 2597639 to 2552515, -1.74%
* github.com/gogits/gogs: from 22958438 to 22468436, -2.13%
* github.com/spf13/hugo: from 14926740 to 14625749, -2.02%
* github.com/influxdata/influxdb/cmd/influxd: from 15651864 to 15382665, -1.72%
* github.com/nsqio/nsq/apps/nsqd: from 9805160 to 9619978, -1.89%
* github.com/mholt/caddy: from 13232411 to 12993178, -1.81%

## Benchmark results

```
BenchmarkMsgpMarshal-4                  10000000               184 ns/op             128 B/op          1 allocs/op
BenchmarkMsgpUnmarshal-4                 5000000               401 ns/op             112 B/op          3 allocs/op
BenchmarkEasyJsonMarshal-4               1000000              1520 ns/op             784 B/op          5 allocs/op
BenchmarkEasyJsonUnmarshal-4             1000000              1554 ns/op             160 B/op          4 allocs/op
BenchmarkFlatBuffersMarshal-4            5000000               355 ns/op               0 B/op          0 allocs/op
BenchmarkFlatBuffersUnmarshal-4          5000000               299 ns/op             112 B/op          3 allocs/op
BenchmarkGogoprotobufMarshal-4          10000000               162 ns/op              64 B/op          1 allocs/op
BenchmarkGogoprotobufUnmarshal-4         5000000               251 ns/op              96 B/op          3 allocs/op
PASS
ok      github.com/alecthomas/go_serialization_benchmarks       14.889s

benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               209           184           -11.96%
BenchmarkMsgpUnmarshal-4             413           392           -5.08%
BenchmarkEasyJsonMarshal-4           1508          1439          -4.58%
BenchmarkEasyJsonUnmarshal-4         1571          1526          -2.86%
BenchmarkFlatBuffersMarshal-4        357           352           -1.40%
BenchmarkFlatBuffersUnmarshal-4      292           287           -1.71%
BenchmarkGogoprotobufMarshal-4       181           162           -10.50%
BenchmarkGogoprotobufUnmarshal-4     266           251           -5.64%

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

## Highlights

* [encoding/json: add Encoder.DisableHTMLEscaping](https://github.com/golang/go/commit/ab52ad894f453a02153fb2bc106d08c47ba643b6)
* [api: update next.txt](https://github.com/golang/go/commit/f7d19672f273ecb600d0b0db32990d1a6462a898)
* [strings: use SSE4.2 in strings.Index on AMD64](https://github.com/golang/go/commit/6b02a1924725688b4d264065454ac5287fbed535)
* [unicode: improve SimpleFold performance for ascii](https://github.com/golang/go/commit/e607abbfd6e0550c13f4fa7b666d033eb9b14759)
* [cmd/compile: switch to compact export format by default](https://github.com/golang/go/commit/7538b1db8ec0d82a623847fe5987f1988fe16448)

## GIT Log

```
commit 9faf5cdf9d1f9050a03ae3d420768c846e54646d
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Thu Apr 28 09:41:32 2016 -0300

    net: change type of dnsConfig.timeout from int to time.Duration
    
    Instead of keeping the desired number of seconds and converting to
    time.Duration for every query, convert to time.Duration when
    building the config.
    
    Updates #15473
    
    Change-Id: Ib24c050b593b3109011e359f4ed837a3fb45dc65
    Reviewed-on: https://go-review.googlesource.com/22548
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 22db3c5a62b01dba6122230aa71d35c48107c70c
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Sat Apr 23 21:00:38 2016 +0300

    cmd/vet: improve checking unkeyed fields in composite literals
    
    - Simplified the code.
    
    - Removed types for slice aliases from composite literals' whitelist, since they
    are properly handled by vet.
    
    Fixes #15408
    Updates #9171
    Updates #11041
    
    Change-Id: Ia1806c9eb3f327c09d2e28da4ffdb233b5a159b0
    Reviewed-on: https://go-review.googlesource.com/22318
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 80e9a7f0797c73b27471eb4b371baa1c7ccb427b
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Apr 27 22:04:49 2016 -0700

    cmd/compile: have all or no parameter named in exported signatures
    
    Binary export format only.
    
    Make sure we don't accidentally export an unnamed parameter
    in signatures which expect all named parameters; otherwise
    we crash during import. Appears to happen for _ (blank)
    parameter names, as observed in method signatures such as
    the one at: x/tools/godoc/analysis/analysis.go:76.
    
    Fixes #15470.
    
    TBR=mdempsky
    
    Change-Id: I1b1184bf08c4c09d8a46946539c4b8c341acdb84
    Reviewed-on: https://go-review.googlesource.com/22543
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit e8d4ffb7661a5e2662d93340d88244f9f833b153
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Apr 27 17:30:20 2016 -0700

    cmd/compile: use correct (field/method) node for position info
    
    Position info for fields and methods was based on the wrong node
    in the new export format, leading to position info for empty
    file names and 0 line numbers. Use correct node now.
    
    Due to compact delta encoding, there is no difference in export
    format size. In fact, because encoding of "no line changed" is
    uncommon and thus a bit more expensive, in many cases the data
    is now slightly shorter.
    
    Stats for export data size (pachage, before, after, delta%):
    
                                             archive/tar     5128     5025  -1%
                                             archive/zip     7672     7515  -1%
                                                   bufio     3486     3377  -2%
                                                   bytes     4879     4821   0%
                                           cmd/addr2line       66       66   0%
                                                 cmd/api    14033    13970   0%
                                                 cmd/asm       60       60   0%
                                   cmd/asm/internal/arch    11659    11647   0%
                                    cmd/asm/internal/asm    13858    13766   0%
                                  cmd/asm/internal/flags      306      300  -1%
                                    cmd/asm/internal/lex    15684    15623   0%
                                                 cmd/cgo    15383    15298   0%
                                             cmd/compile       63       63   0%
                              cmd/compile/internal/amd64      838      838   0%
                                cmd/compile/internal/arm     7333     7323   0%
                              cmd/compile/internal/arm64    19953    19951   0%
                                cmd/compile/internal/big     8943     9043   1%
                                 cmd/compile/internal/gc    57465    56801   0%
                             cmd/compile/internal/mips64      645      645   0%
                              cmd/compile/internal/ppc64      695      695   0%
                              cmd/compile/internal/s390x      553      553   0%
                                cmd/compile/internal/ssa    34883    34559   0%
                                cmd/compile/internal/x86      744      744   0%
                                               cmd/cover     4961     4892   0%
                                                cmd/dist      145      145   0%
                                                 cmd/doc     8891     8853   0%
                                             cmd/expdump       67       67   0%
                                                 cmd/fix      422      406  -3%
                                                  cmd/go     9951     9747  -1%
                                               cmd/gofmt       66       66   0%
                                        cmd/internal/bio     6378     6340   0%
                                     cmd/internal/gcprog      684      644  -5%
                                      cmd/internal/goobj     1276     1193  -6%
                                        cmd/internal/obj    12908    12551  -2%
                                    cmd/internal/obj/arm    10074    10053   0%
                                  cmd/internal/obj/arm64    17723    17699   0%
                                   cmd/internal/obj/mips    12573    12530   0%
                                  cmd/internal/obj/ppc64    15352    15330   0%
                                  cmd/internal/obj/s390x    18785    18769   0%
                                    cmd/internal/obj/x86    23586    23551   0%
                                    cmd/internal/objfile    17148    17359   1%
                             cmd/internal/pprof/commands     1948     1930   0%
                               cmd/internal/pprof/driver    11123    11095   0%
                                cmd/internal/pprof/fetch     8931     8907   0%
                               cmd/internal/pprof/plugin    15335    15221   0%
                              cmd/internal/pprof/profile     8493     8370   0%
                               cmd/internal/pprof/report     9273     9214   0%
                                  cmd/internal/pprof/svg     1589     1589   0%
                           cmd/internal/pprof/symbolizer     8737     8727   0%
                              cmd/internal/pprof/symbolz     8277     8346   1%
                             cmd/internal/pprof/tempfile     4319     4317   0%
                                        cmd/internal/sys      622      603  -2%
      cmd/internal/unvendor/golang.org/x/arch/arm/armasm    79231    79148   0%
      cmd/internal/unvendor/golang.org/x/arch/x86/x86asm    11761    11726   0%
                                                cmd/link       60       60   0%
                                 cmd/link/internal/amd64    11190    11178   0%
                                   cmd/link/internal/arm      204      204   0%
                                 cmd/link/internal/arm64      210      210   0%
                                    cmd/link/internal/ld    60670    59758  -1%
                                cmd/link/internal/mips64      213      213   0%
                                 cmd/link/internal/ppc64      211      211   0%
                                 cmd/link/internal/s390x      210      210   0%
                                   cmd/link/internal/x86      203      203   0%
                                                  cmd/nm       57       57   0%
                                             cmd/objdump       64       64   0%
                                                cmd/pack     4968     4908   0%
                                               cmd/pprof       63       63   0%
                                               cmd/trace      828      782  -5%
                                                 cmd/vet    13485    13503   0%
                              cmd/vet/internal/whitelist      109      109   0%
                                                cmd/yacc     1315     1269  -2%
                                          compress/bzip2     2561     2506  -1%
                                          compress/flate     4906     4748  -2%
                                           compress/gzip     7788     7717   0%
                                            compress/lzw      406      402   0%
                                           compress/zlib     4739     4712   0%
                                          container/heap      265      257  -2%
                                          container/list     1506     1450  -3%
                                          container/ring      556      536  -3%
                                                 context     3552     3527   0%
                                                  crypto      864      834  -2%
                                              crypto/aes      313      311   0%
                                           crypto/cipher     1139     1138   0%
                                              crypto/des      317      315   0%
                                              crypto/dsa     5326     5304   0%
                                            crypto/ecdsa     6383     6364   0%
                                         crypto/elliptic     5983     6063   1%
                                             crypto/hmac      258      256   0%
                                              crypto/md5      722      700  -2%
                                             crypto/rand     4996     4993   0%
                                              crypto/rc4      327      317  -2%
                                              crypto/rsa     6763     6722   0%
                                             crypto/sha1      767      744  -2%
                                           crypto/sha256      348      348   0%
                                           crypto/sha512      487      487   0%
                                           crypto/subtle      620      620   0%
                                              crypto/tls    24344    24083   0%
                                             crypto/x509    17473    17524   0%
                                        crypto/x509/pkix     9682     9596   0%
                                            database/sql     8099     7831  -2%
                                     database/sql/driver     1556     1500  -3%
                                             debug/dwarf     9358     9010  -3%
                                               debug/elf    28226    27882   0%
                                             debug/gosym     2472     2333  -5%
                                             debug/macho     9032     8830  -1%
                                                debug/pe     8561     8328  -2%
                                          debug/plan9obj     1347     1295  -3%
                                                encoding      275      261  -4%
                                        encoding/ascii85      775      738  -4%
                                           encoding/asn1     1280     1246  -2%
                                         encoding/base32     1207     1146  -4%
                                         encoding/base64     1471     1407  -3%
                                         encoding/binary     2430     2386  -1%
                                            encoding/csv     4347     4280  -1%
                                            encoding/gob    13488    13387   0%
                                            encoding/hex      665      646  -2%
                                           encoding/json    11763    11592   0%
                                            encoding/pem      283      273  -3%
                                            encoding/xml    13804    13631   0%
                                                  errors      166      162  -1%
                                                  expvar     1193     1139  -4%
                                                    flag     6896     6964   1%
                                                     fmt     1247     1213  -2%
                                                  go/ast    15797    15473  -1%
                                                go/build     6497     6336  -1%
                                             go/constant     1846     1820   0%
                                                  go/doc     3942     3871  -1%
                                               go/format     1854     1850   0%
                                             go/importer     1702     1695   0%
                               go/internal/gccgoimporter     2084     2063   0%
                                  go/internal/gcimporter     3236     3253   1%
                                               go/parser     7377     7371   0%
                                              go/printer     2480     2469   0%
                                              go/scanner     3806     3733  -1%
                                                go/token     3579     3523  -1%
                                                go/types    26514    26117   0%
                                                    hash      323      295  -8%
                                            hash/adler32      568      554  -1%
                                              hash/crc32      843      825  -1%
                                              hash/crc64      758      739  -2%
                                                hash/fnv     1583     1530  -2%
                                                    html      113      113   0%
                                           html/template    16957    16937   0%
                                                   image    11470    11045  -3%
                                             image/color     2566     2503  -1%
                                     image/color/palette      165      163   0%
                                              image/draw     2543     2522   0%
                                               image/gif     3467     3439   0%
                                image/internal/imageutil     3481     3479   0%
                                              image/jpeg     2725     2717   0%
                                               image/png     2702     2689   0%
                                       index/suffixarray     5802     5777   0%
                                           internal/race      274      274   0%
                                   internal/singleflight      756      718  -4%
                         internal/syscall/windows/sysdll      162      162   0%
                                        internal/testenv     5288     5276   0%
                                          internal/trace     1853     1768  -4%
                                                      io     3425     3349  -1%
                                               io/ioutil     4768     4756   0%
                                                     log     4173     4224   1%
                                              log/syslog     5049     4996   0%
                                                    math     4343     4343   0%
                                                math/big     8779     8817   0%
                                              math/cmplx     1580     1580   0%
                                               math/rand      944      982   4%
                                                    mime     2313     2298   0%
                                          mime/multipart     5021     4922  -1%
                                    mime/quotedprintable     2049     2008  -1%
                                                     net    19332    19090   0%
                                                net/http    50404    49542  -1%
                                            net/http/cgi    22533    22637   0%
                                      net/http/cookiejar     5488     5431   0%
                                           net/http/fcgi    20557    20512   0%
                                       net/http/httptest    30350    30255   0%
                                       net/http/httputil    24045    23964   0%
                                       net/http/internal     2579     2550   0%
                                          net/http/pprof    20307    20258   0%
                                   net/internal/socktest     2227     2159  -2%
                                                net/mail     5086     5054   0%
                                                 net/rpc    28365    28208   0%
                                         net/rpc/jsonrpc    12805    12722   0%
                                                net/smtp    19975    19887   0%
                                           net/textproto     4558     4466  -1%
                                                 net/url     1391     1326  -4%
                                                      os    10372    10195  -1%
                                                 os/exec     7814     7928   1%
                                               os/signal      239      237   0%
                                                 os/user      735      717  -1%
                                                    path      391      391   0%
                                           path/filepath     4136     4136   0%
                                                 reflect     6258     5893  -5%
                                                  regexp     5808     5623  -2%
                                           regexp/syntax     3118     3077   0%
                                                 runtime    11685    10912  -6%
                                             runtime/cgo       18       18   0%
                                           runtime/debug     3320     3304   0%
                                 runtime/internal/atomic      728      728   0%
                                    runtime/internal/sys     2287     2287   0%
                                           runtime/pprof      611      587  -3%
                                            runtime/race       19       19   0%
                                           runtime/trace      145      143   0%
                                                    sort     1229     1206  -1%
                                                 strconv     1752     1744   0%
                                                 strings     3809     3775   0%
                                                    sync     1331     1306  -1%
                                             sync/atomic     1135     1130   0%
                                                 syscall    46280    45722   0%
                                                 testing     7558     7284  -3%
                                          testing/iotest     1122     1071  -4%
                                           testing/quick     5656     5609   0%
                                            text/scanner     3367     3312  -1%
                                          text/tabwriter     2810     2755  -1%
                                           text/template    15613    15595   0%
                                     text/template/parse     9499     9040  -4%
                                                    time     5515     5395  -1%
                                                 unicode     4357     4344   0%
                                           unicode/utf16      583      583   0%
                                            unicode/utf8      970      970   0%
                     vendor/golang.org/x/net/http2/hpack     4105     4012  -1%
                                                 average  1524284  1509610   0%
    
    Change-Id: Ibe1ce098c7c575965389c1cad368c62c2cea256a
    Reviewed-on: https://go-review.googlesource.com/22536
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit c04bc70cd1b290436917301846992e492944805f
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Apr 28 11:50:48 2016 +0900

    net: fix misrecongnization of IPv6 zone on Windows
    
    Fixes #15463.
    
    Change-Id: Ic85886861c650ffcb71240d847941534152b92bc
    Reviewed-on: https://go-review.googlesource.com/22540
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit cad04e7e66d51486b9d3b63bf1b9281a0153dd9f
Author: David du Colombier <0intro@gmail.com>
Date:   Thu Apr 28 01:01:15 2016 +0200

    net/http: skip TestTransportRemovesDeadIdleConnections on Plan 9
    
    Updates #15464.
    
    Change-Id: If3221034bb10751c6fcf1fbeba401a879c18079f
    Reviewed-on: https://go-review.googlesource.com/22513
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2da642a917c618c367a2fc01eb5b704769312199
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Apr 27 15:15:47 2016 +1000

    cmd/compile/internal/gc: unexport {J,S,F,H,B,V}conv
    
    Updates #15462
    
    Unexport Jconv, Sconv, Fconv, Hconv, Bconv, and VConv as they are
    not referenced outside internal/gc.
    
    Econv was only called by EType.String, so merge it into that method.
    
    Change-Id: Iad9b06078eb513b85a03a43cd9eb9366477643d1
    Reviewed-on: https://go-review.googlesource.com/22531
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 733f835f307595366a87fd377ba60c5d23841982
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Apr 27 19:34:17 2016 +1000

    cmd/compile/internal/gc: remove all uses of oconv(op, FmtSharp)
    
    Updates #15462
    
    Replace all use of oconv(op, FmtSharp) with fmt.Printf("%#v", op).
    This removes all the callers of oconv.
    
    Change-Id: Ic3bf22495147f8497c8bada01d681428e2405b0e
    Reviewed-on: https://go-review.googlesource.com/22530
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f08f1cd2e97835dbaa7c509b837265774ae39dac
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Apr 27 16:32:37 2016 -0500

    net: clarify DialContext's use of its provided context
    
    Fixes #15325
    
    Change-Id: I60137ecf27e236e97734b1730ce29ab23e9fe07f
    Reviewed-on: https://go-review.googlesource.com/22509
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 06d639e07588ef290ed28ab384d5371e052240b1
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Apr 27 16:56:13 2016 -0500

    runtime: fix SetCgoTraceback doc indentation
    
    It wasn't rendering as HTML nicely.
    
    Change-Id: I5408ec22932a05e85c210c0faa434bd19dce5650
    Reviewed-on: https://go-review.googlesource.com/22532
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 239fb76ea01b88a2b3dff8b4da357e605435b980
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Apr 25 23:31:36 2016 -0400

    crypto/md5: add s390x assembly implementation
    
    Adapted from md5block_amd64.s.
    
    name                 old speed      new speed      delta
    Hash8Bytes           14.0MB/s ± 1%  39.9MB/s ± 0%  +185.52%   (p=0.000 n=9+10)
    Hash1K                176MB/s ± 1%   661MB/s ± 1%  +274.44%  (p=0.000 n=10+10)
    Hash8K                196MB/s ± 0%   742MB/s ± 1%  +278.35%   (p=0.000 n=10+9)
    Hash8BytesUnaligned  14.2MB/s ± 2%  39.8MB/s ± 0%  +180.06%  (p=0.000 n=10+10)
    Hash1KUnaligned       177MB/s ± 1%   651MB/s ± 0%  +267.38%  (p=0.000 n=10+10)
    Hash8KUnaligned       197MB/s ± 1%   731MB/s ± 1%  +271.73%  (p=0.000 n=10+10)
    
    Change-Id: I45ece98ee10f30fcd192b9c3d743ba61c248f36a
    Reviewed-on: https://go-review.googlesource.com/22505
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f4d38a87927f42272c7dfd10283beac8865edeab
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Mon Mar 28 22:27:36 2016 +1300

    cmd/compile: de-dup the gclocals symbols in compiler too
    
    These symbols are de-duplicated in the linker but the compiler generates quite
    many duplicates too: 2425 of 13769 total symbols for runtime.a for example.
    De-duplicating them in the compiler saves the linker a bit of work.
    
    Fixes #14983
    
    Change-Id: I5f18e5f9743563c795aad8f0a22d17a7ed147711
    Reviewed-on: https://go-review.googlesource.com/22293
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit d3c79d324acd7300b6f705e66af8ca711af00d9f
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Apr 27 15:10:10 2016 +1000

    cmd/compile/internal/gc: remove oconv(op, 0) calls
    
    Updates #15462
    
    Automatic refactor with sed -e.
    
    Replace all oconv(op, 0) to string conversion with the raw op value
    which fmt's %v verb can print directly.
    
    The remaining oconv(op, FmtSharp) will be replaced with op.GoString and
    %#v in the next CL.
    
    Change-Id: I5e2f7ee0bd35caa65c6dd6cb1a866b5e4519e641
    Reviewed-on: https://go-review.googlesource.com/22499
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit cbd72318b964bde9d95102571cd22d1919dbd37f
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Sat Apr 2 18:07:24 2016 -0300

    net: search domain from hostname if no search directives
    
    Fixes #14897
    
    Change-Id: Iffe7462983a5623a37aa0dc6f74c8c70e10c3244
    Reviewed-on: https://go-review.googlesource.com/21464
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 4edb40d441b0def61507e65141535de4d86b9edc
Author: Damien Neil <dneil@google.com>
Date:   Wed Apr 27 11:08:58 2016 -0700

    syscall: fix uint64->int cast of control message header
    
    Change-Id: I28980b307d10730b122a4f833809bc400d6aff24
    Reviewed-on: https://go-review.googlesource.com/22525
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 78bcdeb6a36a6d45f93c8ff546fa946e5fbec093
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Apr 27 15:44:22 2016 -0400

    misc/cgo/testcarchive: fix path of libgo.a for darwin/arm
    
    After CL 22461, c-archive build on darwin/arm is by default compiled
    with -shared, so update the install path.
    
    Fix build.
    
    Change-Id: Ie93dbd226ed416b834da0234210f4b98bc0e3606
    Reviewed-on: https://go-review.googlesource.com/22507
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b49b71ae192c72faf699edd321ff0637f90e794c
Author: Austin Clements <austin@google.com>
Date:   Fri Mar 18 11:27:59 2016 -0400

    runtime: don't rescan globals
    
    Currently the runtime rescans globals during mark 2 and mark
    termination. This costs as much as 500µs/MB in STW time, which is
    enough to surpass the 10ms STW limit with only 20MB of globals.
    
    It's also basically unnecessary. The compiler already generates write
    barriers for global -> heap pointer updates and the regular write
    barrier doesn't check whether the slot is a global or in the heap.
    Some less common write barriers do cause problems.
    heapBitsBulkBarrier, which is used by typedmemmove and related
    functions, currently depends on having access to the pointer bitmap
    and as a result ignores writes to globals. Likewise, the
    reflect-related write barriers reflect_typedmemmovepartial and
    callwritebarrier ignore non-heap destinations; though it appears they
    can never be called with global pointers anyway.
    
    This commit makes heapBitsBulkBarrier issue write barriers for writes
    to global pointers using the data and BSS pointer bitmaps, removes the
    inheap checks from the reflection write barriers, and eliminates the
    rescans during mark 2 and mark termination. It also adds a test that
    writes to globals have write barriers.
    
    Programs with large data+BSS segments (with pointers) aren't common,
    but for programs that do have large data+BSS segments, this
    significantly reduces pause time:
    
    name \ 95%ile-time/markTerm              old         new  delta
    LargeBSS/bss:1GB/gomaxprocs:4  148200µs ± 6%  302µs ±52%  -99.80% (p=0.008 n=5+5)
    
    This very slightly improves the go1 benchmarks:
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.62s ± 3%     2.62s ± 4%    ~     (p=0.904 n=20+20)
    Fannkuch11-12                2.15s ± 1%     2.13s ± 0%  -1.29%  (p=0.000 n=18+20)
    FmtFprintfEmpty-12          48.3ns ± 2%    47.6ns ± 1%  -1.52%  (p=0.000 n=20+16)
    FmtFprintfString-12          152ns ± 0%     152ns ± 1%    ~     (p=0.725 n=18+18)
    FmtFprintfInt-12             150ns ± 1%     149ns ± 1%  -1.14%  (p=0.000 n=19+20)
    FmtFprintfIntInt-12          250ns ± 0%     244ns ± 1%  -2.12%  (p=0.000 n=20+18)
    FmtFprintfPrefixedInt-12     219ns ± 1%     217ns ± 1%  -1.20%  (p=0.000 n=19+20)
    FmtFprintfFloat-12           280ns ± 0%     281ns ± 1%  +0.47%  (p=0.000 n=19+19)
    FmtManyArgs-12               928ns ± 0%     923ns ± 1%  -0.53%  (p=0.000 n=19+18)
    GobDecode-12                7.21ms ± 1%    7.24ms ± 2%    ~     (p=0.091 n=19+19)
    GobEncode-12                6.07ms ± 1%    6.05ms ± 1%  -0.36%  (p=0.002 n=20+17)
    Gzip-12                      265ms ± 1%     265ms ± 1%    ~     (p=0.496 n=20+19)
    Gunzip-12                   39.6ms ± 1%    39.3ms ± 1%  -0.85%  (p=0.000 n=19+19)
    HTTPClientServer-12         74.0µs ± 2%    73.8µs ± 1%    ~     (p=0.569 n=20+19)
    JSONEncode-12               15.4ms ± 1%    15.3ms ± 1%  -0.25%  (p=0.049 n=17+17)
    JSONDecode-12               53.7ms ± 2%    53.0ms ± 1%  -1.29%  (p=0.000 n=18+17)
    Mandelbrot200-12            3.97ms ± 1%    3.97ms ± 0%    ~     (p=0.072 n=17+18)
    GoParse-12                  3.35ms ± 2%    3.36ms ± 1%  +0.51%  (p=0.005 n=18+20)
    RegexpMatchEasy0_32-12      72.7ns ± 2%    72.2ns ± 1%  -0.70%  (p=0.005 n=19+19)
    RegexpMatchEasy0_1K-12       246ns ± 1%     245ns ± 0%  -0.60%  (p=0.000 n=18+16)
    RegexpMatchEasy1_32-12      72.8ns ± 1%    72.5ns ± 1%  -0.37%  (p=0.011 n=18+18)
    RegexpMatchEasy1_1K-12       380ns ± 1%     385ns ± 1%  +1.34%  (p=0.000 n=20+19)
    RegexpMatchMedium_32-12      115ns ± 2%     115ns ± 1%  +0.44%  (p=0.047 n=20+20)
    RegexpMatchMedium_1K-12     35.4µs ± 1%    35.5µs ± 1%    ~     (p=0.079 n=18+19)
    RegexpMatchHard_32-12       1.83µs ± 0%    1.80µs ± 1%  -1.76%  (p=0.000 n=18+18)
    RegexpMatchHard_1K-12       55.1µs ± 0%    54.3µs ± 1%  -1.42%  (p=0.000 n=18+19)
    Revcomp-12                   386ms ± 1%     381ms ± 1%  -1.14%  (p=0.000 n=18+18)
    Template-12                 61.5ms ± 2%    61.5ms ± 2%    ~     (p=0.647 n=19+20)
    TimeParse-12                 338ns ± 0%     336ns ± 1%  -0.72%  (p=0.000 n=14+19)
    TimeFormat-12                350ns ± 0%     357ns ± 0%  +2.05%  (p=0.000 n=19+18)
    [Geo mean]                  55.3µs         55.0µs       -0.41%
    
    Change-Id: I57e8720385a1b991aeebd111b6874354308e2a6b
    Reviewed-on: https://go-review.googlesource.com/20829
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 30172f1811a3e08487c6191d1923f8608a338496
Author: Austin Clements <austin@google.com>
Date:   Wed Apr 27 14:30:01 2016 -0400

    runtime: make {add,subtract}{b,1} nosplit
    
    These are used at the bottom level of various GC operations that must
    not be preempted. To be on the safe side, mark them all nosplit.
    
    Change-Id: I8f7360e79c9852bd044df71413b8581ad764380c
    Reviewed-on: https://go-review.googlesource.com/22504
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit bddfc337f9b053740b51e1fd8429f84dafa89205
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed Apr 27 13:10:49 2016 -0400

    reflect: fix strings of SliceOf-created types
    
    The new type was inheriting the tflagExtraStar from its prototype.
    
    Fixes #15467
    
    Change-Id: Ic22c2a55cee7580cb59228d52b97e1c0a1e60220
    Reviewed-on: https://go-review.googlesource.com/22501
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit 217be5b35d8fb0f812ca59bf7dec3aa0fb850c46
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed Apr 27 12:49:27 2016 -0400

    reflect: unnamed interface types have no name
    
    Fixes #15468
    
    Change-Id: I8723171f87774a98d5e80e7832ebb96dd1fbea74
    Reviewed-on: https://go-review.googlesource.com/22524
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>

commit 74a9bad63899ffb02b747678c2c181ffb13983b9
Author: Zhongwei Yao <zhongwei.yao@arm.com>
Date:   Mon Apr 25 11:08:38 2016 +0800

    cmd/compile: enable const division for arm64
    
    performance:
    benchmark                   old ns/op     new ns/op     delta
    BenchmarkDivconstI64-8      8.28          2.70          -67.39%
    BenchmarkDivconstU64-8      8.28          4.69          -43.36%
    BenchmarkDivconstI32-8      8.28          6.39          -22.83%
    BenchmarkDivconstU32-8      8.28          4.43          -46.50%
    BenchmarkDivconstI16-8      5.17          5.17          +0.00%
    BenchmarkDivconstU16-8      5.33          5.34          +0.19%
    BenchmarkDivconstI8-8       3.50          3.50          +0.00%
    BenchmarkDivconstU8-8       3.51          3.50          -0.28%
    
    Fixes #15382
    
    Change-Id: Ibce7b28f0586d593b33c4d4ecc5d5e7e7c905d13
    Reviewed-on: https://go-review.googlesource.com/22292
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: David Chase <drchase@google.com>

commit 7538b1db8ec0d82a623847fe5987f1988fe16448
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Apr 15 14:14:04 2016 -0700

    cmd/compile: switch to compact export format by default
    
    builtin.go was auto-generated via go generate; all other
    changes were manual.
    
    The new format reduces the export data size by ~65% on average
    for the std library packages (and there is still quite a bit of
    room for improvement).
    
    The average time to write export data is reduced by (at least)
    62% as measured in one run over the std lib, it is likely more.
    
    The average time to read import data is reduced by (at least)
    37% as measured in one run over the std lib, it is likely more.
    There is also room to improve this time.
    
    The compiler transparently handles both packages using the old
    and the new format.
    
    Comparing the -S output of the go build for each package via
    the cmp.bash script (added) shows identical assembly code for
    all packages, but 6 files show file:line differences:
    
    The following files have differences because they use cgo
    and cgo uses different temp. directories for different builds.
    Harmless.
    
    	src/crypto/x509
    	src/net
    	src/os/user
    	src/runtime/cgo
    
    The following files have file:line differences that are not yet
    fully explained; however the differences exist w/ and w/o new export
    format (pre-existing condition). See issue #15453.
    
    	src/go/internal/gccgoimporter
    	src/go/internal/gcimporter
    
    In summary, switching to the new export format produces the same
    package files as before for all practical purposes.
    
    How can you tell which one you have (if you care): Open a package
    (.a) file in an editor. Textual export data starts with a $$ after
    the header and is more or less legible; binary export data starts
    with a $$B after the header and is mostly unreadable. A stand-alone
    decoder (for debugging) is in the works.
    
    In case of a problem, please first try reverting back to the old
    textual format to determine if the cause is the new export format:
    
    For a stand-alone compiler invocation:
    - go tool compile -newexport=0 <files>
    
    For a single package:
    - go build -gcflags="-newexport=0" <pkg>
    
    For make/all.bash:
    - (export GO_GCFLAGS="-newexport=0"; sh make.bash)
    
    Fixes #13241.
    
    Change-Id: I2588cb463be80af22446bf80c225e92ab79878b8
    Reviewed-on: https://go-review.googlesource.com/22123
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 70d95a488da89d268d0a61171ec389982a62184d
Author: Michael Matloob <matloob@golang.org>
Date:   Tue Apr 26 15:28:17 2016 -0400

    regexp: add a harder regexp to the benchmarks
    
    This regexp has many parallel alternations
    
    Change-Id: I8044f460aa7d18f20cb0452e9470557b87facd6d
    Reviewed-on: https://go-review.googlesource.com/22471
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9629f55fbbfb4052ea24c247cbd5db49ba2f389e
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Apr 26 15:17:56 2016 -0400

    cmd/link: remove absolute address for c-archive on darwin/arm
    
    Now it is possible to build a c-archive as PIC on darwin/arm (this is
    now the default). Then the system linker can link the binary using
    the archive as PIE.
    
    Fixes #12896.
    
    Change-Id: Iad84131572422190f5fa036e7d71910dc155f155
    Reviewed-on: https://go-review.googlesource.com/22461
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 86c93c989e73e823e9e66f3d3e319b616544c320
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Apr 26 22:31:02 2016 -0700

    cmd/compile: don't write pos info for builtin packages
    
    TestBuiltin will fail if run on Windows and builtin.go was generated
    on a non-Windows machine (or vice versa) because path names have
    different separators. Avoid problem altogether by not writing pos
    info for builtin packages. It's not needed.
    
    Affects -newexport only.
    
    Change-Id: I8944f343452faebaea9a08b5fb62829bed77c148
    Reviewed-on: https://go-review.googlesource.com/22498
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a19e60b2c3c388abd3503da3fc2659bef1e76b46
Author: Keith Randall <khr@golang.org>
Date:   Tue Apr 26 15:22:33 2016 -0700

    cmd/compile: don't use line numbers from ONAME and named OLITERALs
    
    The line numbers of ONAMEs are the location of their
    declaration, not their use.
    
    The line numbers of named OLITERALs are also the location
    of their declaration.
    
    Ignore both of these.  Instead, we will inherit the line number from
    the containing syntactic item.
    
    Fixes #14742
    Fixes #15430
    
    Change-Id: Ie43b5b9f6321cbf8cead56e37ccc9364d0702f2f
    Reviewed-on: https://go-review.googlesource.com/22479
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c9389a101b020b41a954ac35642ec254c8344238
Author: Zhongwei Yao <zhongwei.yao@arm.com>
Date:   Thu Mar 31 15:34:12 2016 +0800

    cmd/asm: fix SIMD register name on arm64
    
    Current V-register range is V32~V63 on arm64. This patch changes it to
    V0~V31.
    
    fix #15465.
    
    Change-Id: I90dab42dea46825ec5d7a8321ec4f6550735feb8
    Reviewed-on: https://go-review.googlesource.com/22520
    Reviewed-by: Aram Hăvărneanu <aram@mgk.ro>
    Run-TryBot: Aram Hăvărneanu <aram@mgk.ro>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6dfba5c7ce867583cbcea9da09dceacd2633bacc
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri Feb 26 13:02:42 2016 +0100

    runtime/race: improve TestNoRaceIOHttp test
    
    TestNoRaceIOHttp does all kinds of bad things:
    1. Binds to a fixed port, so concurrent tests fail.
    2. Registers HTTP handler multiple times, so repeated tests fail.
    3. Relies on sleep to wait for listen.
    
    Fix all of that.
    
    Change-Id: I1210b7797ef5e92465b37dc407246d92a2a24fe8
    Reviewed-on: https://go-review.googlesource.com/19953
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 102cf2ae0321775eef2d36d7c4258b740fe92458
Author: Martin Möhrmann <martisch@uos.de>
Date:   Fri Apr 15 09:38:36 2016 +0200

    image/color: optimize RGBToYCbCr
    
    Apply optimizations used to speed up YCbCrToRGB from
    https://go-review.googlesource.com/#/c/21910/
    to RGBToYCbCr.
    
    name             old time/op  new time/op  delta
    RGBToYCbCr/0-2   6.81ns ± 0%  5.96ns ± 0%  -12.48%  (p=0.000 n=38+50)
    RGBToYCbCr/Cb-2  7.68ns ± 0%  6.13ns ± 0%  -20.21%  (p=0.000 n=50+33)
    RGBToYCbCr/Cr-2  6.84ns ± 0%  6.04ns ± 0%  -11.70%  (p=0.000 n=39+42)
    
    Updates #15260
    
    Change-Id: If3ea5393ae371a955ddf18ab226aae20b48f9692
    Reviewed-on: https://go-review.googlesource.com/22411
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ralph Corderoy <ralph@inputplus.co.uk>

commit 8f2e780e8ac29e47466103998484c0a73df34d51
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Apr 27 14:46:09 2016 +1000

    cmd/compile/internal: unexport gc.Oconv
    
    Updates #15462
    
    Semi automatic change with gofmt -r and hand fixups for callers outside
    internal/gc.
    
    All the uses of gc.Oconv outside cmd/compile/internal/gc were for the
    Oconv(op, 0) form, which is already handled the Op.String method.
    
    Replace the use of gc.Oconv(op, 0) with op itself, which will call
    Op.String via the %v or %s verb. Unexport Oconv.
    
    Change-Id: I84da2a2e4381b35f52efce427b2d6a3bccdf2526
    Reviewed-on: https://go-review.googlesource.com/22496
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 707aed0363c31bfef761a86464a09ecf0817267e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Apr 26 19:11:53 2016 -0700

    cmd/compile: fix opnames
    
    Change-Id: Ief4707747338912216a8509b1adbf655c8ffac56
    Reviewed-on: https://go-review.googlesource.com/22495
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2e30218223a7bf2b560fbaf79bac8d80ea4ece1c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Apr 26 15:43:04 2016 -0700

    net/http: remove idle transport connections from Transport when server closes
    
    Previously the Transport would cache idle connections from the
    Transport for later reuse, but if a peer server disconnected
    (e.g. idle timeout), we would not proactively remove the *persistConn
    from the Transport's idle list, leading to a waste of memory
    (potentially forever).
    
    Instead, when the persistConn's readLoop terminates, remote it from
    the idle list, if present.
    
    This also adds the beginning of accounting for the total number of
    idle connections, which will be needed for Transport.MaxIdleConns
    later.
    
    Updates #15461
    
    Change-Id: Iab091f180f8dd1ee0d78f34b9705d68743b5557b
    Reviewed-on: https://go-review.googlesource.com/22492
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 87bca88c703c1f14fe8473dc2f07dc521cf2b989
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Apr 26 18:54:12 2016 -0700

    context: fix doc typo
    
    Fixes #15449
    
    Change-Id: I8d84d076a05c56694b48f7b84f572b1a6524f522
    Reviewed-on: https://go-review.googlesource.com/22493
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 0b5fbf7098b10da284f48de604e7d2860d2cd9d0
Author: Russ Cox <rsc@golang.org>
Date:   Mon Apr 25 10:51:26 2016 -0400

    cmd/go: add Package.StaleReason for debugging with go list
    
    It comes up every few months that we can't understand why
    the go command is rebuilding some package.
    Add diagnostics so that the go command can explain itself
    if asked.
    
    For #2775, #3506, #12074.
    
    Change-Id: I1c73b492589b49886bf31a8f9d05514adbd6ed70
    Reviewed-on: https://go-review.googlesource.com/22432
    Reviewed-by: Rob Pike <r@golang.org>

commit 525ae3f897bf79fd78f3e693bd65056efc8f9109
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Apr 25 17:58:34 2016 -0400

    crypto/sha256: add s390x assembly implementation
    
    Renames block to blockGeneric so that it can be called when the
    assembly feature check fails. This means making block a var on
    platforms without an assembly implementation (similar to the sha1
    package).
    
    Also adds a test to check that the fallback path works correctly
    when the feature check fails.
    
    name        old speed      new speed       delta
    Hash8Bytes  6.42MB/s ± 1%  27.14MB/s ± 0%  +323.01%  (p=0.000 n=10+10)
    Hash1K      53.9MB/s ± 0%  511.1MB/s ± 0%  +847.57%   (p=0.000 n=10+9)
    Hash8K      57.1MB/s ± 1%  609.7MB/s ± 0%  +967.04%  (p=0.000 n=10+10)
    
    Change-Id: If962b2a5c9160b3a0b76ccee53b2fd809468ed3d
    Reviewed-on: https://go-review.googlesource.com/22460
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2a889b9d931e58166350f785b16edc51e28ef19b
Author: Austin Clements <austin@google.com>
Date:   Fri Mar 4 11:58:26 2016 -0500

    runtime: make stack re-scan O(# dirty stacks)
    
    Currently the stack re-scan during mark termination is O(# stacks)
    because we enqueue a root marking job for every goroutine. It takes
    ~34ns to process this root marking job for a valid (clean) stack, so
    at around 300k goroutines we exceed the 10ms pause goal. A non-trivial
    portion of this time is spent simply taking the cache miss to check
    the gcscanvalid flag, so simply optimizing the path that handles clean
    stacks can only improve this so much.
    
    Fix this by keeping an explicit list of goroutines with dirty stacks
    that need to be rescanned. When a goroutine first transitions to
    running after a stack scan and marks its stack dirty, it adds itself
    to this list. We enqueue root marking jobs only for the goroutines in
    this list, so this improves stack re-scanning asymptotically by
    completely eliminating time spent on clean goroutines.
    
    This reduces mark termination time for 500k idle goroutines from 15ms
    to 238µs. Overall performance effect is negligible.
    
    name \ 95%ile-time/markTerm     old           new         delta
    IdleGs/gs:500000/gomaxprocs:12  15000µs ± 0%  238µs ± 5%  -98.41% (p=0.000 n=10+10)
    
    name              old time/op  new time/op  delta
    XBenchGarbage-12  2.30ms ± 3%  2.29ms ± 1%  -0.43%  (p=0.049 n=17+18)
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.57s ± 3%     2.59s ± 2%    ~     (p=0.141 n=19+20)
    Fannkuch11-12                2.09s ± 0%     2.10s ± 1%  +0.53%  (p=0.000 n=19+19)
    FmtFprintfEmpty-12          45.3ns ± 3%    45.2ns ± 2%    ~     (p=0.845 n=20+20)
    FmtFprintfString-12          129ns ± 0%     127ns ± 0%  -1.55%  (p=0.000 n=16+16)
    FmtFprintfInt-12             123ns ± 0%     119ns ± 1%  -3.24%  (p=0.000 n=19+19)
    FmtFprintfIntInt-12          195ns ± 1%     189ns ± 1%  -3.11%  (p=0.000 n=17+17)
    FmtFprintfPrefixedInt-12     193ns ± 1%     187ns ± 1%  -3.06%  (p=0.000 n=19+19)
    FmtFprintfFloat-12           254ns ± 0%     255ns ± 1%  +0.35%  (p=0.001 n=14+17)
    FmtManyArgs-12               781ns ± 0%     770ns ± 0%  -1.48%  (p=0.000 n=16+19)
    GobDecode-12                7.00ms ± 1%    6.98ms ± 1%    ~     (p=0.563 n=19+19)
    GobEncode-12                5.91ms ± 1%    5.92ms ± 0%    ~     (p=0.118 n=19+18)
    Gzip-12                      219ms ± 1%     215ms ± 1%  -1.81%  (p=0.000 n=18+18)
    Gunzip-12                   37.2ms ± 0%    37.4ms ± 0%  +0.45%  (p=0.000 n=17+19)
    HTTPClientServer-12         76.9µs ± 3%    77.5µs ± 2%  +0.81%  (p=0.030 n=20+19)
    JSONEncode-12               15.0ms ± 0%    14.8ms ± 1%  -0.88%  (p=0.001 n=15+19)
    JSONDecode-12               50.6ms ± 0%    53.2ms ± 2%  +5.07%  (p=0.000 n=17+19)
    Mandelbrot200-12            4.05ms ± 0%    4.05ms ± 1%    ~     (p=0.581 n=16+17)
    GoParse-12                  3.34ms ± 1%    3.30ms ± 1%  -1.21%  (p=0.000 n=15+20)
    RegexpMatchEasy0_32-12      69.6ns ± 1%    69.8ns ± 2%    ~     (p=0.566 n=19+19)
    RegexpMatchEasy0_1K-12       238ns ± 1%     236ns ± 0%  -0.91%  (p=0.000 n=17+13)
    RegexpMatchEasy1_32-12      69.8ns ± 1%    70.0ns ± 1%  +0.23%  (p=0.026 n=17+16)
    RegexpMatchEasy1_1K-12       371ns ± 1%     363ns ± 1%  -2.07%  (p=0.000 n=19+19)
    RegexpMatchMedium_32-12      107ns ± 2%     106ns ± 1%  -0.51%  (p=0.031 n=18+20)
    RegexpMatchMedium_1K-12     33.0µs ± 0%    32.9µs ± 0%  -0.30%  (p=0.004 n=16+16)
    RegexpMatchHard_32-12       1.70µs ± 0%    1.70µs ± 0%  +0.45%  (p=0.000 n=16+17)
    RegexpMatchHard_1K-12       51.1µs ± 2%    51.4µs ± 1%  +0.53%  (p=0.000 n=17+19)
    Revcomp-12                   378ms ± 1%     385ms ± 1%  +1.92%  (p=0.000 n=19+18)
    Template-12                 64.3ms ± 2%    65.0ms ± 2%  +1.09%  (p=0.001 n=19+19)
    TimeParse-12                 315ns ± 1%     317ns ± 2%    ~     (p=0.108 n=18+20)
    TimeFormat-12                360ns ± 1%     337ns ± 0%  -6.30%  (p=0.000 n=18+13)
    [Geo mean]                  51.8µs         51.6µs       -0.48%
    
    Change-Id: Icf8994671476840e3998236e15407a505d4c760c
    Reviewed-on: https://go-review.googlesource.com/20700
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5b765ce310c594276ea919a9cb455cc894fee999
Author: Austin Clements <austin@google.com>
Date:   Fri Mar 11 14:08:10 2016 -0500

    runtime: don't clear gcscanvalid in casfrom_Gscanstatus
    
    Currently we clear gcscanvalid in both casgstatus and
    casfrom_Gscanstatus if the new status is _Grunning. This is very
    important to do in casgstatus. However, this is potentially wrong in
    casfrom_Gscanstatus because in this case the caller doesn't own gp and
    hence the write is racy. Unlike the other _Gscan statuses, during
    _Gscanrunning, the G is still running. This does not indicate that
    it's transitioning into a running state. The scan simply hasn't
    happened yet, so it's neither valid nor invalid.
    
    Conveniently, this also means clearing gcscanvalid is unnecessary in
    this case because the G was already in _Grunning, so we can simply
    remove this code. What will happen instead is that the G will be
    preempted to scan itself, that scan will set gcscanvalid to true, and
    then the G will return to _Grunning via casgstatus, clearing
    gcscanvalid.
    
    This fix will become necessary shortly when we start keeping track of
    the set of G's with dirty stacks, since it will no longer be
    idempotent to simply set gcscanvalid to false.
    
    Change-Id: I688c82e6fbf00d5dbbbff49efa66acb99ee86785
    Reviewed-on: https://go-review.googlesource.com/20669
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c707d8385639dfda22dc06b112f5f7af78006a1f
Author: Austin Clements <austin@google.com>
Date:   Mon Apr 18 18:28:36 2016 -0400

    runtime: fix typos in comment about gcscanvalid
    
    Change-Id: Id4ad7ebf88a21eba2bc5714b96570ed5cfaed757
    Reviewed-on: https://go-review.googlesource.com/22210
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9f263c14edccb564b675ed6c4f12260f333505d5
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 2 17:55:45 2016 -0500

    runtime: remove stack barriers during sweep
    
    This adds a best-effort pass to remove stack barriers immediately
    after the end of mark termination. This isn't necessary for the Go
    runtime, but should help external tools that perform stack walks but
    aren't aware of Go's stack barriers such as GDB, perf, and VTune.
    (Though clearly they'll still have trouble unwinding stacks during
    mark.)
    
    Change-Id: I66600fae1f03ee36b5459d2b00dcc376269af18e
    Reviewed-on: https://go-review.googlesource.com/20668
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 269c969c81774d1579e80a8c35edbd0ebea065a7
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 2 14:52:08 2016 -0500

    runtime: remove stack barriers during concurrent mark
    
    Currently we remove stack barriers during STW mark termination, which
    has a non-trivial per-goroutine cost and means that we have to touch
    even clean stacks during mark termination. However, there's no problem
    with leaving them in during the sweep phase. They just have to be out
    by the time we install new stack barriers immediately prior to
    scanning the stack such as during the mark phase of the next GC cycle
    or during mark termination in a STW GC.
    
    Hence, move the gcRemoveStackBarriers from STW mark termination to
    just before we install new stack barriers during concurrent mark. This
    removes the cost from STW. Furthermore, this combined with concurrent
    stack shrinking means that the mark termination scan of a clean stack
    is a complete no-op, which will make it possible to skip clean stacks
    entirely during mark termination.
    
    This has the downside that it will mess up anything outside of Go that
    tries to walk Go stacks all the time instead of just some of the time.
    This includes tools like GDB, perf, and VTune. We'll improve the
    situation shortly.
    
    Change-Id: Ia40baad8f8c16aeefac05425e00b0cf478137097
    Reviewed-on: https://go-review.googlesource.com/20667
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit efb0c55407cdbef6aa5471f057b8afd1d0303369
Author: Austin Clements <austin@google.com>
Date:   Fri Mar 11 13:54:55 2016 -0500

    runtime: avoid span root marking entirely during mark termination
    
    Currently we enqueue span root mark jobs during both concurrent mark
    and mark termination, but we make the job a no-op during mark
    termination.
    
    This is silly. Instead of queueing them up just to not do them, don't
    queue them up in the first place.
    
    Change-Id: Ie1d36de884abfb17dd0db6f0449a2b7c997affab
    Reviewed-on: https://go-review.googlesource.com/20666
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e8337491aa6b6a5f96f31077764352549dc34159
Author: Austin Clements <austin@google.com>
Date:   Fri Mar 11 17:00:46 2016 -0500

    runtime: free dead G stacks concurrently
    
    Currently we free cached stacks of dead Gs during STW stack root
    marking. We do this during STW because there's no way to take
    ownership of a particular dead G, so attempting to free a dead G's
    stack during concurrent stack root marking could race with reusing
    that G.
    
    However, we can do this concurrently if we take a completely different
    approach. One way to prevent reuse of a dead G is to remove it from
    the free G list. Hence, this adds a new fixed root marking task that
    simply removes all Gs from the list of dead Gs with cached stacks,
    frees their stacks, and then adds them to the list of dead Gs without
    cached stacks.
    
    This is also a necessary step toward rescanning only dirty stacks,
    since it eliminates another task from STW stack marking.
    
    Change-Id: Iefbad03078b284a2e7bf30fba397da4ca87fe095
    Reviewed-on: https://go-review.googlesource.com/20665
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1a2cf91f5e9e3dfb0873e61ed6907cc365857f6c
Author: Austin Clements <austin@google.com>
Date:   Fri Mar 11 16:27:51 2016 -0500

    runtime: split gfree list into with-stacks and without-stacks
    
    Currently all free Gs are added to one list. Split this into two
    lists: one for free Gs with cached stacks and one for Gs without
    cached stacks.
    
    This lets us preferentially allocate Gs that already have a stack, but
    more importantly, it sets us up to free cached G stacks concurrently.
    
    Change-Id: Idbe486f708997e1c9d166662995283f02d1eeb3c
    Reviewed-on: https://go-review.googlesource.com/20664
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3b0efa689ec7a32de30cbda2221452f57abb2532
Author: Keith Randall <khr@golang.org>
Date:   Tue Apr 26 14:09:58 2016 -0700

    cmd/compile: a rule's line number is at its ->
    
    Let's define the line number of a multiline rule as the line
    number on which the -> appears.  This helps make the rule
    cover analysis look a bit nicer.
    
    Change-Id: I4ac4c09f2240285976590ecfd416bc4c05e78946
    Reviewed-on: https://go-review.googlesource.com/22473
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 8d075beeef137455b9dc40f1c724b495f3ceda26
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Apr 26 10:55:32 2016 -0700

    cmd/compile: lazily initialize litbuf
    
    Instead of eagerly creating strings like "literal 2.01" for every
    lexed number in case we need to mention it in an error message, defer
    this work to (*parser).syntax_error.
    
    name      old allocs/op  new allocs/op  delta
    Template      482k ± 0%      482k ± 0%  -0.12%   (p=0.000 n=9+10)
    GoTypes      1.35M ± 0%     1.35M ± 0%  -0.04%  (p=0.015 n=10+10)
    Compiler     5.45M ± 0%     5.44M ± 0%  -0.12%    (p=0.000 n=9+8)
    
    Change-Id: I333b3c80e583864914412fb38f8c0b7f1d8c8821
    Reviewed-on: https://go-review.googlesource.com/22480
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 19912e1d0a0739cd8a3214de994f7e9ecb656e1d
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Apr 26 15:18:48 2016 -0700

    cmd/dist: sort entries in zcgo.go generated file for deterministic build
    
    This simplifies comparison of object files across different builds
    by ensuring that the strings in the zcgo.go always appear in the
    same order.
    
    Change-Id: I3639ea4fd10e0d645b838d1bbb03cd33deca340e
    Reviewed-on: https://go-review.googlesource.com/22478
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e607abbfd6e0550c13f4fa7b666d033eb9b14759
Author: Egon Elbre <egonelbre@gmail.com>
Date:   Tue Nov 17 16:51:23 2015 +0200

    unicode: improve SimpleFold performance for ascii
    
    This change significantly speeds up case-insensitive regexp matching.
    
    benchmark                      old ns/op      new ns/op      delta
    BenchmarkMatchEasy0i_32-8      2690           1473           -45.24%
    BenchmarkMatchEasy0i_1K-8      80404          42269          -47.43%
    BenchmarkMatchEasy0i_32K-8     3272187        2076118        -36.55%
    BenchmarkMatchEasy0i_1M-8      104805990      66503805       -36.55%
    BenchmarkMatchEasy0i_32M-8     3360192200     2126121600     -36.73%
    
    benchmark                      old MB/s     new MB/s     speedup
    BenchmarkMatchEasy0i_32-8      11.90        21.72        1.83x
    BenchmarkMatchEasy0i_1K-8      12.74        24.23        1.90x
    BenchmarkMatchEasy0i_32K-8     10.01        15.78        1.58x
    BenchmarkMatchEasy0i_1M-8      10.00        15.77        1.58x
    BenchmarkMatchEasy0i_32M-8     9.99         15.78        1.58x
    
    Issue #13288
    
    Change-Id: I94af7bb29e75d60b4f6ee760124867ab271b9642
    Reviewed-on: https://go-review.googlesource.com/16943
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6e4a8615f652a2020471622354be6d890404020c
Author: Alan Donovan <adonovan@google.com>
Date:   Mon Apr 25 18:31:36 2016 -0400

    gc: use AbsFileLine for deterministic binary export data
    
    This version of the file name honors the -trimprefix flag,
    which strips off variable parts like $WORK or $PWD.
    The TestCgoConsistentResults test now passes.
    
    Change-Id: If93980b054f9b13582dd314f9d082c26eaac4f41
    Reviewed-on: https://go-review.googlesource.com/22444
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 17db07f9b5034f22851f32f7700649ac61c44e8f
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Apr 26 14:11:38 2016 -0700

    cmd/compile: don't discard inlineable but empty functions with binary export format
    
    Change-Id: I0f016fa000f949d27847d645b4cdebe68a8abf20
    Reviewed-on: https://go-review.googlesource.com/22474
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3a72d626a8bae104c658f361d97f992f609d91e7
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue Apr 26 21:37:49 2016 +1200

    cmd/link: pass -no-pie (if supported) when creating a race-enabled executable.
    
    Fixes #15443
    
    Change-Id: Ia3593104fc1a4255926ae5675c25390604b44b7b
    Reviewed-on: https://go-review.googlesource.com/22453
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 55154cf0b27e3c48e7cf7654c890868a95e7eed6
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Apr 13 13:34:41 2016 -0400

    cmd/link: fix gdb backtrace on architectures using a link register
    
    Also adds TestGdbBacktrace to the runtime package.
    
    Dwarf modifications written by Bryan Chan (@bryanpkc) who is also
    at IBM and covered by the same CLA.
    
    Fixes #14628
    
    Change-Id: I106a1f704c3745a31f29cdadb0032e3905829850
    Reviewed-on: https://go-review.googlesource.com/20193
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 01d5e63faa7cbfe10c6c45a788cd9859da2dfcdb
Author: Russ Cox <rsc@golang.org>
Date:   Tue Apr 26 13:18:14 2016 -0400

    cmd/compile/internal/gc: rewrite comment to avoid automated meaning
    
    The comment says 'DΟ NΟT SUBMIT', and that text being in a file can cause
    automated errors or warnings when trying to check the Go sources into other
    source control systems.
    
    (We reject that string in CL commit messages, which I've avoided here
    by changing the O's to Ο's above.)
    
    Change-Id: I6cdd57a8612ded5208f05a8bd6b137f44424a030
    Reviewed-on: https://go-review.googlesource.com/22434
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit e4355aeedfdd6a68185c4551c889eb13823cd86d
Author: Keith Randall <khr@golang.org>
Date:   Wed Apr 20 11:17:41 2016 -0700

    cmd/compile: more sanity checks on rewrite rules
    
    Make sure ops have the right number of args, set
    aux and auxint only if allowed, etc.
    
    Normalize error reporting format.
    
    Change-Id: Ie545fcc5990c8c7d62d40d9a0a55885f941eb645
    Reviewed-on: https://go-review.googlesource.com/22320
    Reviewed-by: David Chase <drchase@google.com>

commit 24a297286a3032223c432a830a53ebf102e08de4
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Apr 25 16:17:42 2016 -0400

    crypto/sha512: add s390x assembly implementation
    
    Renames block to blockGeneric so that it can be called when the
    assembly feature check fails. This means making block a var on
    platforms without an assembly implementation (similar to the sha1
    package).
    
    Also adds a test to check that the fallback path works correctly
    when the feature check fails.
    
    name        old speed      new speed       delta
    Hash8Bytes  7.13MB/s ± 2%  19.89MB/s ± 1%  +178.82%   (p=0.000 n=9+10)
    Hash1K       121MB/s ± 1%    661MB/s ± 1%  +444.54%   (p=0.000 n=10+9)
    Hash8K       137MB/s ± 0%    918MB/s ± 1%  +569.29%  (p=0.000 n=10+10)
    
    Change-Id: Id65dd6e943f14eeffe39a904dc88065fc6a60179
    Reviewed-on: https://go-review.googlesource.com/22402
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 98b99d561225cc1d140360b217df2acc9aa1f746
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Apr 25 13:09:11 2016 -0700

    net: ignore lame referral responses like libresolv
    
    Fixes #15434.
    
    Change-Id: Ia88b740df5418a6d3af1c29a03756f4234f388b0
    Reviewed-on: https://go-review.googlesource.com/22428
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 96b8f70e22e103c11fbb89ba6df9d229d24cdbc2
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Apr 26 10:53:25 2016 -0400

    cmd/link: correctly decode name length
    
    The linker was incorrectly decoding type name lengths, causing
    typelinks to be sorted out of order and in cases where the name was
    the exact right length, linker panics.
    
    Added a test to the reflect package that causes TestTypelinksSorted
    to fail before this CL. It's not the exact failure seen in #15448
    but it has the same cause: decodetype_name calculating the wrong
    length.
    
    The equivalent decoders in reflect/type.go and runtime/type.go
    have the parenthesis in the right place.
    
    Fixes #15448
    
    Change-Id: I33257633d812b7d2091393cb9d6cc8a73e0138c8
    Reviewed-on: https://go-review.googlesource.com/22403
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0b6332eb54767f916926ae39516ddaed87b26edb
Author: David Chase <drchase@google.com>
Date:   Mon Apr 25 16:24:11 2016 -0400

    cmd/compile: fix another bug in dominator computation
    
    Here, "fix" means "replace".  The new dominator computation
    is the "simple" algorithm from Lengauer and Tarjan's TOPLAS
    paper, with minimal changes.
    
    Also included is a test that tweaks the fixed error.
    
    Change-Id: I0abdf53d5d64df1e67e4e62f55e88957045cd63b
    Reviewed-on: https://go-review.googlesource.com/22401
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 6b02a1924725688b4d264065454ac5287fbed535
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Thu Apr 21 18:24:12 2016 +0300

    strings: use SSE4.2 in strings.Index on AMD64
    
    Use PCMPESTRI instruction if available.
    
    Index-4              21.1ns ± 0%  21.1ns ± 0%     ~     (all samples are equal)
    IndexHard1-4          395µs ± 0%   105µs ± 0%  -73.53%        (p=0.000 n=19+20)
    IndexHard2-4          300µs ± 0%   147µs ± 0%  -51.11%        (p=0.000 n=19+20)
    IndexHard3-4          665µs ± 0%   665µs ± 0%     ~           (p=0.942 n=16+19)
    
    Change-Id: I4f66794164740a2b939eb1c78934e2390b489064
    Reviewed-on: https://go-review.googlesource.com/22337
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit d78c84c419b0ecdd70e85aad22951798c1707f50
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Apr 25 15:59:42 2016 -0700

    cmd/compile: sort import strings for canonical obj files
    
    This is not necessary for reproduceability but it removes
    differences due to imported package order between compiles
    using textual vs binary export format. The packages list
    tends to be very short, so it's ok doing it always for now.
    
    Guarded with a documented (const) flag so it's trivial to
    disable and remove eventually.
    
    Also, use the same flag now to enforce parameter numbering.
    
    Change-Id: Ie05d2490df770239696ecbecc07532ed62ccd5c0
    Reviewed-on: https://go-review.googlesource.com/22445
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 9cb79e9536a2f7977f9139a808f912d216094ecc
Author: Keith Randall <khr@golang.org>
Date:   Mon Apr 25 14:12:26 2016 -0700

    runtime: arm5, fix large-offset floating-point stores
    
    The code sequence for large-offset floating-point stores
    includes adding the base pointer to r11.  Make sure we
    can interpret that instruction correctly.
    
    Fixes build.
    
    Fixes #15440
    
    Change-Id: I7fe5a4a57e08682967052bf77c54e0ec47fcb53e
    Reviewed-on: https://go-review.googlesource.com/22440
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 0b8c0767d0b95066734647edeb5a252c270a4a1a
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Apr 25 14:39:51 2016 -0700

    cmd/compile: for now, keep parameter numbering in binary export format
    
    The numbering is only required for parameters of functions/methods
    with exported inlineable bodies. For now, always export parameter names
    with internal numbering to minimize the diffs between assembly code
    dumps of code compiled with the textual vs the binary format.
    
    To be disabled again once the new export format is default.
    
    Change-Id: I6d14c564e734cc5596c7e995d8851e06d5a35013
    Reviewed-on: https://go-review.googlesource.com/22441
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e48a2958d1cfa4ae75dead9d8e65489b53c70f14
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Apr 25 13:24:48 2016 -0700

    cmd/compile: treat empty and absent struct field tags as identical
    
    Fixes #15439.
    
    Change-Id: I5a32384c46e20f8db6968e5a9e854c45ab262fe4
    Reviewed-on: https://go-review.googlesource.com/22429
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6f3f02f80d23d3bbc2857be147341517d1762fbd
Author: Keith Randall <khr@golang.org>
Date:   Sun Apr 24 17:04:32 2016 -0700

    runtime: zero tmpbuf between len and cap
    
    Zero the entire buffer so we don't need to
    lower its capacity upon return.  This lets callers
    do some appending without allocation.
    
    Zeroing is cheap, the byte buffer requires only
    4 extra instructions.
    
    Fixes #14235
    
    Change-Id: I970d7badcef047dafac75ac17130030181f18fe2
    Reviewed-on: https://go-review.googlesource.com/22424
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8b92397bcdcd5d6de3f72951a5514933fee32eb2
Author: Alexandru Moșoi <brtzsnr@gmail.com>
Date:   Sun Apr 24 21:21:07 2016 +0200

    cmd/compile: introduce bool operations.
    
    Introduce OrB, EqB, NeqB, AndB to handle bool operations.
    
    Change-Id: I53e4d5125a8090d5eeb4576db619103f19fff58d
    Reviewed-on: https://go-review.googlesource.com/22412
    Reviewed-by: Keith Randall <khr@golang.org>

commit 0436a89a2c5afad41356dc1dff7c745cd30636a7
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Apr 22 16:35:29 2016 -0700

    spec: be more explicit about equivalence of empty string and absent field tags
    
    Note that the spec already makes that point with a comment in the very first
    example for struct field tags. This change is simply stating this explicitly
    in the actual spec prose.
    
    - gccgo and go/types already follow this rule
    - the current reflect package API doesn't distinguish between absent tags
      and empty tags (i.e., there is no discoverable difference)
    
    Fixes #15412.
    
    Change-Id: I92f9c283064137b4c8651630cee0343720717a02
    Reviewed-on: https://go-review.googlesource.com/22391
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 9dcbc43f4f299b8ea6546a464d9fdeb5839b5ae9
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Apr 25 10:55:22 2016 -0700

    test: add test for issue 15084
    
    The underlying issues have been fixed.
    All the individual fixes have their own tests,
    but it's still useful to have a plain source test.
    
    Fixes #15084
    
    Change-Id: I06c485a7d0716201bd57d1f3be53668dddd7ec14
    Reviewed-on: https://go-review.googlesource.com/22426
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit f12bd8a5a8f8485f13793f03d4803a924923badb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Apr 21 11:55:33 2016 -0700

    cmd/compile: encapsulate OSLICE* representation
    
    As a nice side-effect, this allows us to
    unify several code paths.
    
    The terminology (low, high, max, simple slice expr,
    full slice expr) is taken from the spec and
    the examples in the spec.
    
    This is a trial run. The plan, probably for Go 1.8,
    is to change slice expressions to use Node.List
    instead of OKEY, and to do some similar
    tree structure changes for other ops.
    
    Passes toolstash -cmp. No performance change.
    all.bash passes with GO_GCFLAGS=-newexport.
    
    Updates #15350
    
    Change-Id: Ic1efdc36e79cdb95ae1636e9817a3ac8f83ab1ac
    Reviewed-on: https://go-review.googlesource.com/22425
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2bf7034d87c051ff3d3fcce9f62d7ef2d2f4108d
Author: Russ Cox <rsc@golang.org>
Date:   Mon Apr 25 10:48:45 2016 -0400

    cmd/go: disable failing TestGoGetInsecure
    
    Update #15410
    
    Change-Id: Iad3f2639aa7a67b11efc35a629e1893f7d87b957
    Reviewed-on: https://go-review.googlesource.com/22430
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 093ac15a14137b4a9454442ae9fea282e5c09180
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sun Apr 24 15:09:00 2016 +1000

    debug/pe: better error messages
    
    Updates #15345
    
    Change-Id: Iae35d3e378cbc8157ba1ff91e4971ed4515a5e5c
    Reviewed-on: https://go-review.googlesource.com/22394
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 758431fe8c2906690a209e33531d8b95e381c8c1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Apr 24 14:09:03 2016 -0700

    cmd/compile: minor cleanup in inl
    
    * Make budget an int32 to avoid needless conversions.
    * Introduce some temporary variables to reduce repetition.
    * If ... args are present, they will be the last argument
      to the function. No need to scan all arguments.
    
    Passes toolstash -cmp.
    
    Change-Id: I55203609f5d2f25a4e238cd48c63214651120cfc
    Reviewed-on: https://go-review.googlesource.com/22421
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1da62afeef1fdfb822afc4af0feb2eece10d8c7d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Apr 24 13:50:26 2016 -0700

    cmd/compile: replace len(Nodes.Slice()) with Nodes.Len()
    
    Generated with eg:
    
    func before(n gc.Nodes) int { return len(n.Slice()) }
    func after(n gc.Nodes) int  { return n.Len() }
    
    Change-Id: Ifdf01915e60069166afe96aa7b1d08720bf62fc5
    Reviewed-on: https://go-review.googlesource.com/22420
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3de87bc4d92d58446d5072d416a29366b83135c1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Apr 24 13:31:55 2016 -0700

    test: add test that required algs get generated
    
    This is a follow-up to CLs 19769 and 19770.
    
    Change-Id: Ia9b71055613b80df4ce62b34fcc4f479f04f72fd
    Reviewed-on: https://go-review.googlesource.com/22399
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit fca0f331c8b99d476c871d8718e296b32ad24073
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Apr 22 08:39:56 2016 -0700

    cmd/compile: use gc.Etype's String method
    
    Passes toolstash -cmp.
    
    Change-Id: I42c962cc5a3ffec2969f223cf238c2fdadbf5857
    Reviewed-on: https://go-review.googlesource.com/22381
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f027241445f3064b41f5d5e68f86370d37bad0be
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Apr 22 07:14:10 2016 -0700

    cmd/compile: give gc.Op a String method, use it
    
    Passes toolstash -cmp.
    
    Change-Id: I915e76374fd64aa2597e6fa47e4fa95ca00ca643
    Reviewed-on: https://go-review.googlesource.com/22380
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit a6abc1cd70bf561d1e4c10d53499733c502c30b5
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 4 15:01:03 2015 -0700

    cmd/compile: don't generate algs for map buckets
    
    Note that this is only safe because
    the compiler generates multiple distinct
    gc.Types. If we switch to having canonical
    gc.Types, then this will need to be updated
    to handle the case in which the user uses both
    map[T]S and also map[[8]T]S. In that case,
    the runtime needs algs for [8]T, but this could
    mark the sole [8]T type as Noalg. This is a general
    problem with having a single bool to represent
    whether alg generation is needed for a type.
    
    Cuts 5k off cmd/go and 22k off golang.org/x/tools/cmd/godoc,
    approx 0.04% and 0.12% respectively.
    
    For #6853 and #9930
    
    Change-Id: I30a15ec72ecb62e2aa053260a7f0f75015fc0ade
    Reviewed-on: https://go-review.googlesource.com/19769
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b6b144bf97744ead3ac51fd1b5648d2e31a8de0e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 4 15:01:29 2015 -0700

    cmd/compile: don't generate algs for ... args
    
    Note that this is only safe because
    the compiler generates multiple distinct
    gc.Types. If we switch to having canonical
    gc.Types, then this will need to be updated
    to handle the case in which the user uses both
    map[[n]T]S and also calls a function f(...T) with n arguments.
    In that case, the runtime needs algs for [n]T, but this could
    mark the sole [n]T type as Noalg. This is a general
    problem with having a single bool to represent
    whether alg generation is needed for a type.
    
    Cuts 17k off cmd/go and 13k off golang.org/x/tools/cmd/godoc,
    approx 0.14% and 0.07% respectively.
    
    For #6853 and #9930
    
    Change-Id: Iccb6b9fd88ade5497d7090528a903816d340bf0a
    Reviewed-on: https://go-review.googlesource.com/19770
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 934c3599648ae841668ec753881134347fc28c29
Author: Keith Randall <khr@golang.org>
Date:   Sat Apr 23 22:59:01 2016 -0700

    cmd/compile: reorder how slicelit initializes a slice
    
      func f(x, y, z *int) {
        a := []*int{x,y,z}
        ...
      }
    
    We used to use:
      var tmp [3]*int
      a := tmp[:]
      a[0] = x
      a[1] = y
      a[2] = z
    
    Now we do:
      var tmp [3]*int
      tmp[0] = x
      tmp[1] = y
      tmp[2] = z
      a := tmp[:]
    
    Doesn't sound like a big deal, but the compiler has trouble
    eliminating write barriers when using the former method because it
    doesn't know that the slice points to the stack.  In the latter
    method, the compiler knows the array is on the stack and as a result
    doesn't emit any write barriers.
    
    This turns out to be extremely common when building ... args, like
    for calls fmt.Printf.
    
    Makes go binaries ~1% smaller.
    
    Doesn't have a measurable effect on the go1 fmt benchmarks,
    unfortunately.
    
    Fixes #14263
    Update #6853
    
    Change-Id: I9074a2788ec9e561a75f3b71c119b69f304d6ba2
    Reviewed-on: https://go-review.googlesource.com/22395
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit c4c182140adedf300800778127840e3b8b9f7423
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Sat Apr 23 21:18:34 2016 +0200

    internal/trace: fix event ordering for coarse timestamps
    
    Arm arch uses coarse-grained kernel timer as cputicks.
    As the result sort.Sort smashes trace entirely. Use sort.Stable instead.
    
    Change-Id: Idfa017a86a489be58cf239f7fe56d7f4b66b52a9
    Reviewed-on: https://go-review.googlesource.com/22317
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 75b844f0d228bda5dea2aabae096909f81355bac
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Thu Apr 7 15:48:15 2016 +0200

    runtime/trace: test detection of broken timestamps
    
    On some processors cputicks (used to generate trace timestamps)
    produce non-monotonic timestamps. It is important that the parser
    distinguishes logically inconsistent traces (e.g. missing, excessive
    or misordered events) from broken timestamps. The former is a bug
    in tracer, the latter is a machine issue.
    
    Test that (1) parser does not return a logical error in case of
    broken timestamps and (2) broken timestamps are eventually detected
    and reported.
    
    Change-Id: Ib4b1eb43ce128b268e754400ed8b5e8def04bd78
    Reviewed-on: https://go-review.googlesource.com/21608
    Reviewed-by: Austin Clements <austin@google.com>

commit 687fe991e42f15fe1f491680c615ef96568f780a
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu Apr 21 16:51:36 2016 +1000

    debug/pe: introduce File.COFFSymbols and (*COFFSymbol).FullName
    
    Reloc.SymbolTableIndex is an index into symbol table. But
    Reloc.SymbolTableIndex cannot be used as index into File.Symbols,
    because File.Symbols slice has Aux lines removed as it is built.
    
    We cannot change the way File.Symbols works, so I propose we
    introduce new File.COFFSymbols that does not have that limitation.
    
    Also unlike File.Symbols, File.COFFSymbols will consist of
    COFFSymbol. COFFSymbol matches PE COFF specification exactly,
    and it is simpler to use.
    
    Updates #15345
    
    Change-Id: Icbc265853a472529cd6d64a76427b27e5459e373
    Reviewed-on: https://go-review.googlesource.com/22336
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d224e98d9ac969e733f5578dce3e1831c5c84f45
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Apr 22 18:49:59 2016 -0700

    cmd/link: add -dumpdep flag to dump linker dependency graph
    
    This is what led to https://golang.org/cl/20763 and
    https://golang.org/cl/20765 to shrink binary sizes.
    
    Change-Id: Id360d474e6153cfe32a525b0a720810fd113195b
    Reviewed-on: https://go-review.googlesource.com/22392
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 9e3c68f1e02021a845c452ae347d06332e4ed79d
Author: Keith Randall <khr@golang.org>
Date:   Fri Apr 22 13:09:18 2016 -0700

    cmd/compile: get rid of most byte and word insns for amd64
    
    Now that we're using 32-bit ops for 8/16-bit logical operations
    (to avoid partial register stalls), there's really no need to
    keep track of the 8/16-bit ops at all.  Convert everything we
    can to 32-bit ops.
    
    This CL is the obvious stuff.  I might think a bit more about
    whether we can get rid of weirder stuff like HMULWU.
    
    The only downside to this CL is that we lose some information
    about constants.  If we had source like:
      var a byte = ...
      a += 128
      a += 128
    We will convert that to a += 256, when we could get rid of the
    add altogether.  This seems like a fairly unusual scenario and
    I'm happy with forgoing that optimization.
    
    Change-Id: Ia7c1e5203d0d110807da69ed646535194a3efba1
    Reviewed-on: https://go-review.googlesource.com/22382
    Reviewed-by: Todd Neal <todd@tneal.org>

commit 217c284995400bb761e5718782c8a90748c75aef
Author: Keith Randall <khr@golang.org>
Date:   Wed Apr 20 15:02:48 2016 -0700

    cmd/compile: combine stores into larger widths
    
    Combine stores into larger widths when it is safe to do so.
    
    Add clobber() function so stray dead uses do not impede the
    above rewrites.
    
    Fix bug in loads where all intermediate values depending on
    a small load (not just the load itself) must have no other uses.
    We really need the small load to be dead after the rewrite..
    
    Fixes #14267
    
    Change-Id: Ib25666cb19777f65082c76238fba51a76beb5d74
    Reviewed-on: https://go-review.googlesource.com/22326
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Todd Neal <todd@tneal.org>

commit a3703618eadeb74b60f2cb9a23fabe178d4b141d
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Tue Apr 5 15:29:14 2016 +0200

    runtime: use per-goroutine sequence numbers in tracer
    
    Currently tracer uses global sequencer and it introduces
    significant slowdown on parallel machines (up to 10x).
    Replace the global sequencer with per-goroutine sequencer.
    
    If we assign per-goroutine sequence numbers to only 3 types
    of events (start, unblock and syscall exit), it is enough to
    restore consistent partial ordering of all events. Even these
    events don't need sequence numbers all the time (if goroutine
    starts on the same P where it was unblocked, then start does
    not need sequence number).
    The burden of restoring the order is put on trace parser.
    Details of the algorithm are described in the comments.
    
    On http benchmark with GOMAXPROCS=48:
    no tracing: 5026 ns/op
    tracing: 27803 ns/op (+453%)
    with this change: 6369 ns/op (+26%, mostly for traceback)
    
    Also trace size is reduced by ~22%. Average event size before: 4.63
    bytes/event, after: 3.62 bytes/event.
    
    Besides running trace tests, I've also tested with manually broken
    cputicks (random skew for each event, per-P skew and episodic random skew).
    In all cases broken timestamps were detected and no test failures.
    
    Change-Id: I078bde421ccc386a66f6c2051ab207bcd5613efa
    Reviewed-on: https://go-review.googlesource.com/21512
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ba966f5d891191450f286859c35bf8a7fa49cde2
Author: Francesc Campoy <campoy@golang.org>
Date:   Fri Apr 22 16:27:34 2016 -0700

    doc: mention security from contribution guidelines
    
    Fixes #15413
    
    Change-Id: I837a391276eed565cf66d3715ec68b7b959ce143
    Reviewed-on: https://go-review.googlesource.com/22390
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit f058ab09fb14afe3a51b880a6895b96aa3e07c85
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Apr 22 15:45:24 2016 -0700

    cmd/compile: remove redundant "// fallthrough" comments
    
    Change-Id: Ia3f262f06592b66447c213e2350402cd5e6e2ccd
    Reviewed-on: https://go-review.googlesource.com/22389
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 889c0a66fc7a43b23cc02ee42cfa17d221fce3c4
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Apr 22 14:50:20 2016 -0700

    cmd/compile: don't export pos info in new export format for now
    
    Exporting filenames as part of the position information can lead
    to different object files which breaks tests.
    
    Change-Id: Ia678ab64293ebf04bf83601e6ba72919d05762a4
    Reviewed-on: https://go-review.googlesource.com/22385
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f7d19672f273ecb600d0b0db32990d1a6462a898
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Apr 22 22:15:17 2016 +0000

    api: update next.txt
    
    Change-Id: I12d5e5d0e74b354f26898bab4ea30eb27ac45cd7
    Reviewed-on: https://go-review.googlesource.com/22387
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3564ec52cda2a3c83aaf41159b26369ca4e7ecee
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Apr 22 22:13:50 2016 +0000

    cmd/api: ignore vendored packages
    
    Fixes #15404
    
    Change-Id: I16f2a34a1e4c3457053a1fc2141f21747cfb22b4
    Reviewed-on: https://go-review.googlesource.com/22386
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 3411d6321979b33291e3b4c6fe79d4dd41bd5fba
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Apr 15 19:19:58 2016 -0700

    net: keep waiting for valid DNS response until timeout
    
    Prevents denial of service attacks from bogus UDP packets.
    
    Fixes #13281.
    
    Change-Id: Ifb51b17a1b0807bfd27b144d6037431701184e7b
    Reviewed-on: https://go-review.googlesource.com/22126
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9f1ccd647fcdb1b703c1042c90434e15aff75013
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Fri Apr 22 00:47:04 2016 +0530

    net/url: validate ports in IPv4 addresses
    
    Fixes #14860
    
    Change-Id: Id55ad942d45a104d560a879d6e8e1aa09671789b
    Reviewed-on: https://go-review.googlesource.com/22351
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ab52ad894f453a02153fb2bc106d08c47ba643b6
Author: Caleb Spare <cespare@gmail.com>
Date:   Sat Apr 9 21:18:22 2016 -0700

    encoding/json: add Encoder.DisableHTMLEscaping
    
    This provides a way to disable the escaping of <, >, and & in JSON
    strings.
    
    Fixes #14749.
    
    Change-Id: I1afeb0244455fc8b06c6cce920444532f229555b
    Reviewed-on: https://go-review.googlesource.com/21796
    Run-TryBot: Caleb Spare <cespare@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 97360096e5e9fdea06be8c97f32bd83741f68adb
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Apr 22 12:27:29 2016 -0700

    cmd/compile: replace Ctype switches with type switches
    
    Instead of switching on Ctype (which internally uses a type switch)
    and then scattering lots of type assertions throughout the CTFOO case
    clauses, just use type switches directly on the underlying constant
    value.
    
    Passes toolstash/buildall.
    
    Change-Id: I9bc172cc67e5f391cddc15539907883b4010689e
    Reviewed-on: https://go-review.googlesource.com/22384
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2d342fba78d9cbddb4c8c71bfc0d1044b2e5c58a
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri Apr 22 22:48:11 2016 +0200

    runtime: fix description of trace events
    
    Change-Id: I037101b1921fe151695d32e9874b50dd64982298
    Reviewed-on: https://go-review.googlesource.com/22314
    Reviewed-by: Austin Clements <austin@google.com>

commit e05b9746ddc6e53864d1ab26fc172b09ccbe321c
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Apr 21 20:58:35 2016 -0700

    cmd/compile: map TSLICE to obj.KindSlice directly
    
    Change-Id: Idab5f603c1743895b8f4edbcc55f7be83419a099
    Reviewed-on: https://go-review.googlesource.com/22383
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit eee20b7244586da70b2bca6fe6346da7dac6be78
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Apr 22 10:17:06 2016 -0700

    cmd/dist: skip misc/cgo/test with internal linking on ppc64le
    
    CL 22372 changed ppc64le to use normal cgo initialization on ppc64le.
    Doing this uncovered a cmd/link error using internal linking.
    Opened issue 15409 for the problem.  This CL disables the test.
    
    Update #15409.
    
    Change-Id: Ia1bb6b874c1b5a4df1a0436c8841c145142c30f7
    Reviewed-on: https://go-review.googlesource.com/22379
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f71a13da93b6ec2da88e2568deacdb6ef002b36c
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Tue Apr 12 20:50:59 2016 +0200

    cmd/trace: generate new pprof profiles
    
    Generate new protobuf pprof profiles with embed symbol info.
    This makes program binary unnecessary.
    
    Change-Id: Ie628439c13c5e34199782031138102c83ea50621
    Reviewed-on: https://go-review.googlesource.com/21873
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>

commit d32229b3b1edd3d3b1e2dbb61bd6ae7cd8400d56
Author: David Chase <drchase@google.com>
Date:   Fri Apr 22 12:15:08 2016 -0400

    cmd/compile: in a Tarjan algorithm, DFS should really be DFS
    
    Replaced incorrect recursion-free rendering of DFS with
    something that was correct.  Enhanced test with all
    permutations of IF successors to ensure that all possible
    DFS traversals are exercised.
    
    Test is improved version of
    https://go-review.googlesource.com/#/c/22334
    
    Update 15084.
    
    Change-Id: I6e944c41244e47fe5f568dfc2b360ff93b94079e
    Reviewed-on: https://go-review.googlesource.com/22347
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: David Chase <drchase@google.com>

commit babd5da61fbaa7a1b3a5413c3c8947d71fa1001d
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Apr 21 08:00:07 2016 -0400

    crypto/aes: use asm for BenchmarkExpand on amd64
    
    This reverses the change to this benchmark made in 9b6bf20.
    
    Change-Id: I79ab88286c3028d3be561957140375bbc413e7ab
    Reviewed-on: https://go-review.googlesource.com/22340
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 5833d843de4d774f6b76e982b170f42b8e94a198
Author: Chris Zou <chriszou@ca.ibm.com>
Date:   Mon Apr 18 19:30:17 2016 -0400

    hash/crc32: use vector instructions on s390x
    
    The input buffer is aligned to a doubleword boundary to
    improve performance of the vector instructions. The pure
    Go implementation is used to align the input data, and is
    also used when the vector instructions are not available
    or the data length is less than 64 bytes.
    
    Change-Id: Ie259a5f2f1562bcc17961c99e5776c99091d6bed
    Reviewed-on: https://go-review.googlesource.com/22201
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7879e9193b39e6455ae03f2baace9c41f6393ee4
Author: Alexandru Moșoi <mosoi@google.com>
Date:   Thu Apr 21 10:11:33 2016 +0200

    cmd/compile: reenable phielim during rewrite
    
    Remove the "optimization" that was causing the issue.
    
    For the following code the "optimization" was
    converting v to (OpCopy x) which is wrong because
    x doesn't dominate v.
    
    b1:
        y = ...
        First .. b3
    b2:
       x = ...
       Goto b3
    b3:
       v = phi x y
       ... use v ...
    
    That "optimization" is likely no longer needed because
    we now have a second opt pass with a dce in between
    which removes blocks of type First.
    
    For pkg/tools/linux_amd64/* the binary size drops
    from 82142886 to 82060034.
    
    Change-Id: I10428abbd8b32c5ca66fec3da2e6f3686dddbe31
    Reviewed-on: https://go-review.googlesource.com/22312
    Run-TryBot: Alexandru Moșoi <alexandru@mosoi.ro>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit caef4496fcdaca8dc5b86f60b07760e5434ca1f3
Author: Alexandru Moșoi <mosoi@google.com>
Date:   Fri Apr 22 12:44:31 2016 +0200

    cmd/compile: convert some Phis into And8.
    
    See discussion at [1]. True value must have a fixed non-zero
    representation meaning that a && b can be implemented as a & b.
    
    [1] https://groups.google.com/forum/#!topic/golang-dev/xV0vPuFP9Vg
    
    This change helps with m := a && b, but it's more common to see
    if a && b { do something } which is not handled.
    
    Change-Id: Ib6f9ff898a0a8c05d12466e2464e4fe781035394
    Reviewed-on: https://go-review.googlesource.com/22313
    Run-TryBot: Alexandru Moșoi <alexandru@mosoi.ro>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3c1a4c1902711c16489ed0c3506df97439ffbd85
Author: Keith Randall <khr@golang.org>
Date:   Tue Apr 19 21:06:53 2016 -0700

    cmd/compile: don't nilcheck newobject and return values from mapaccess{1,2}
    
    They are guaranteed to be non-nil, no point in inserting
    nil checks for them.
    
    Fixes #15390
    
    Change-Id: I3b9a0f2319affc2139dcc446d0a56c6785ae5a86
    Reviewed-on: https://go-review.googlesource.com/22291
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 32302d6289e9721015d5d7ac99bbce30de47746c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Apr 22 07:08:13 2016 -0700

    runtime/cgo: use normal libinit on PPC GNU/Linux
    
    The special case was because PPC did not support external linking, but
    now it does.
    
    Fixes #10410.
    
    Change-Id: I9b024686e0f03da7a44c1c59b41c529802f16ab0
    Reviewed-on: https://go-review.googlesource.com/22372
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c165988360457553ccbfa4a09919de3262a4438a
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Apr 7 21:37:45 2016 -0400

    cmd/compile, etc: use nameOff in uncommonType
    
    linux/amd64 PIE:
    	cmd/go:  -62KB (0.5%)
    	jujud:  -550KB (0.7%)
    
    For #6853.
    
    Change-Id: Ieb67982abce5832e24b997506f0ae7108f747108
    Reviewed-on: https://go-review.googlesource.com/22371
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 1492e7db059ea7903110b0725d5ced3134558e73
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Apr 7 16:29:16 2016 -0400

    cmd/compile, etc: use nameOff for rtype string
    
    linux/amd64:
    	cmd/go:   -8KB (basically nothing)
    
    linux/amd64 PIE:
    	cmd/go: -191KB (1.6%)
    	jujud:  -1.5MB (1.9%)
    
    Updates #6853
    Fixes #15064
    
    Change-Id: I0adbb95685e28be92e8548741df0e11daa0a9b5f
    Reviewed-on: https://go-review.googlesource.com/21777
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit bb52ceafea60dc4688b6c6b71f241752ce597db8
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Apr 7 20:07:09 2016 -0400

    cmd/link: do not align type name data
    
    Now that reflect.name objects contain an offset to pkgPath instead of a
    pointer, there is no need to align the symbol data.
    
    Removes approx. 10KB from the cmd/go binary. The effect becomes more
    important later as more type data is moved into name objects.
    
    For #6853
    
    Change-Id: Idb507fdbdad04f16fc224378f82272cb5c236ab7
    Reviewed-on: https://go-review.googlesource.com/21776
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 63ceeafa308b99e6b7d5480521b83360b4f6b2fd
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Apr 18 17:41:50 2016 -0400

    crypto/sha1: add s390x assembly implementation
    
    Use the compute intermediate message digest (KIMD) instruction
    when possible. Adds test to check fallback code path in case
    KIMD is not available.
    
    Benchmark changes:
    Hash8Bytes  3.4x
    Hash1K      9.3x
    Hash8K      10.9x
    
    Change-Id: Ibcd71a886dfd7b3822042235b4f4eaa7a148036b
    Reviewed-on: https://go-review.googlesource.com/22350
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7085fb31dfb1a6b447a139064a4a692677284132
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon Apr 11 11:14:53 2016 +1000

    html/template, text/template: clarify Parse{Files,Glob} semantics
    
    Document the subtle property that files with equivalent base names
    will overwrite extant templates with those same names.
    
    Fixes golang/go#14320
    
    Change-Id: Ie9ace1b08e6896ea599836e31582123169aa7a25
    Reviewed-on: https://go-review.googlesource.com/21824
    Reviewed-by: Rob Pike <r@golang.org>

commit b563fcfabba559646b42b0ff2066fc98042d8a8a
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon Apr 18 13:25:51 2016 +1000

    cmd/go: write test file to temporary directory
    
    Before this change, a go-vendor-issue-14613 file would be left in the
    working directory after tests run.
    
    Change-Id: If1858421bb287215ab4a19163f489131b2e8912c
    Reviewed-on: https://go-review.googlesource.com/22169
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit 8082828ed0b07225c50a991dbe2a176346fba3b8
Author: Rob Pike <r@golang.org>
Date:   Thu Apr 21 12:43:22 2016 -0700

    encoding/gob: document compatibility
    
    Fixes #13808.
    
    Change-Id: Ifbd5644da995a812438a405485c9e08b4503a313
    Reviewed-on: https://go-review.googlesource.com/22352
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 70184087239482331a9e4a66627e4458ffe48933
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Apr 21 15:55:50 2016 -0700

    flag: update test case (fix build)
    
    Change-Id: I2275dc703be4fda3feedf76483148eab853b43b8
    Reviewed-on: https://go-review.googlesource.com/22360
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 25d95ee918b4b1315cb2fee0fc625d24cd408240
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Apr 22 10:10:08 2016 +1200

    cmd/link: convert Link.Filesyms into a slice
    
    Change-Id: I6490de325b0f4ba962c679503102d30d41dcc384
    Reviewed-on: https://go-review.googlesource.com/22359
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4b175fd23b3bf220e4121ba4986f2d7af1415482
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Apr 22 09:38:41 2016 +1200

    cmd/link: fix Codeblk printing when -a to use Textp as a slice
    
    Does anyone actually pass -a to the linker?
    
    Change-Id: I1d31ea66aa5604b7fd42adf15bdab71e9f52d0ed
    Reviewed-on: https://go-review.googlesource.com/22356
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 0ef041cfa5d51d86c25b039d7ae8aef8a92c085e
Author: Rob Pike <r@golang.org>
Date:   Thu Apr 21 15:09:49 2016 -0700

    doc/go1.7.txt: 0s for zero duration, go doc groups constructors with types
    
    Change-Id: I4fc35649ff5a3510f5667b62e7e84e113e95dffe
    Reviewed-on: https://go-review.googlesource.com/22358
    Reviewed-by: Rob Pike <r@golang.org>

commit 9c4295b574a89bf02294111f811f90ab06b9951b
Author: Rob Pike <r@golang.org>
Date:   Thu Apr 21 14:53:19 2016 -0700

    time: print zero duration as 0s, not 0
    
    There should be a unit, and s is the SI unit name, so use that.
    The other obvious possibility is ns (nanosecond), but the fact
    that durations are measured in nanoseconds is an internal detail.
    
    Fixes #14058.
    
    Change-Id: Id1f8f3c77088224d9f7cd643778713d5cc3be5d9
    Reviewed-on: https://go-review.googlesource.com/22357
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit a33e9cf7ead3d7853546a71796a6c404d02cb474
Author: Rob Pike <r@golang.org>
Date:   Thu Apr 21 13:29:07 2016 -0700

    cmd/doc: group constructors with type in package presentation
    
    Fixes #14004.
    
    $ go doc encoding.gob
    Before:
    func Register(value interface{})
    func RegisterName(name string, value interface{})
    func NewDecoder(r io.Reader) *Decoder
    func NewEncoder(w io.Writer) *Encoder
    type CommonType struct { ... }
    type Decoder struct { ... }
    type Encoder struct { ... }
    type GobDecoder interface { ... }
    type GobEncoder interface { ... }
    
    After:
    func Register(value interface{})
    func RegisterName(name string, value interface{})
    type CommonType struct { ... }
    type Decoder struct { ... }
        func NewDecoder(r io.Reader) *Decoder
    type Encoder struct { ... }
        func NewEncoder(w io.Writer) *Encoder
    type GobDecoder interface { ... }
    type GobEncoder interface { ... }
    
    Change-Id: I021db25bce4a16b3dfa22ab323ca1f4e68d50111
    Reviewed-on: https://go-review.googlesource.com/22354
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 8ad8d7d87edf0aec3b56c2e2d0139bc12531d359
Author: Keith Randall <khr@golang.org>
Date:   Thu Apr 21 13:58:22 2016 -0700

    cmd/compile: Use pre-regalloc value ID in lateSpillUse
    
    The cached copy's ID is sometimes outside the bounds of the orig array.
    
    There's no reason to start at the cached copy and work backwards
    to the original value. We already have the original value ID at
    all the callsites.
    
    Fixes noopt build
    
    Change-Id: I313508a1917e838a87e8cc83b2ef3c2e4a8db304
    Reviewed-on: https://go-review.googlesource.com/22355
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 40f1d0ca9f978376f7db24de3737b58589c8542b
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Apr 18 14:02:08 2016 -0700

    cmd/compile: split TSLICE into separate Type kind
    
    Instead of using TARRAY for both arrays and slices, create a new
    TSLICE kind to handle slices.
    
    Also, get rid of the "DDDArray" distinction. While kinda ugly, it
    seems likely we'll need to defer evaluating the constant bounds
    expressions for golang.org/issue/13890.
    
    Passes toolstash/buildall.
    
    Change-Id: I8e45d4900e7df3a04cce59428ec8b38035d3cc3a
    Reviewed-on: https://go-review.googlesource.com/22329
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5213cd700062917bc98f949479dfc0865751f2e8
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Apr 21 12:47:24 2016 -0700

    spec: fix incorrect comment in shift example
    
    - adjusted example code
    - fixed comments
    
    Fixes #14785.
    
    Change-Id: Ia757dc93b0a69b8408559885ece7f3685a37daaa
    Reviewed-on: https://go-review.googlesource.com/22353
    Reviewed-by: Rob Pike <r@golang.org>

commit c8bd293e56d17c5599ec62aee63fe819366adcab
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 30 17:15:15 2016 -0400

    runtime: eliminate floating garbage estimate
    
    Currently when we compute the trigger for the next GC, we do it based
    on an estimate of the reachable heap size at the start of the GC
    cycle, which is itself based on an estimate of the floating garbage.
    This was introduced by 4655aad to fix a bad feedback loop that allowed
    the heap to grow to many times the true reachable size.
    
    However, this estimate gets easily confused by rapidly allocating
    applications, and, worse it's different than the heap size the trigger
    controller uses to compute the trigger itself. This results in the
    trigger controller often thinking that GC finished before it started.
    Since this would be a pretty great outcome from it's perspective, it
    sets the trigger for the next cycle as close to the next goal as
    possible (which is limited to 95% of the goal).
    
    Furthermore, the bad feedback loop this estimate originally fixed
    seems not to happen any more, suggesting it was fixed more correctly
    by some other change in the mean time. Finally, with the change to
    allocate black, it shouldn't even be theoretically possible for this
    bad feedback loop to occur.
    
    Hence, eliminate the floating garbage estimate and simply consider the
    reachable heap to be the marked heap. This harms overall throughput
    slightly for allocation-heavy benchmarks, but significantly improves
    mutator availability.
    
    Fixes #12204. This brings the average trigger in this benchmark from
    0.95 (the cap) to 0.7 and the active GC utilization from ~90% to ~45%.
    
    Updates #14951. This makes the trigger controller much better behaved,
    so it pulls the trigger lower if assists are consuming a lot of CPU
    like it's supposed to, increasing mutator availability.
    
    name              old time/op  new time/op  delta
    XBenchGarbage-12  2.21ms ± 1%  2.28ms ± 3%  +3.29%  (p=0.000 n=17+17)
    
    Some of this slow down we paid for in earlier commits. Relative to the
    start of the series to switch to allocate-black (the parent of "count
    black allocations toward scan work"), the garbage benchmark is 2.62%
    slower.
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.53s ± 3%     2.53s ± 3%    ~     (p=0.708 n=20+19)
    Fannkuch11-12                2.08s ± 0%     2.08s ± 0%  -0.22%  (p=0.002 n=19+18)
    FmtFprintfEmpty-12          45.3ns ± 2%    45.2ns ± 3%    ~     (p=0.505 n=20+20)
    FmtFprintfString-12          129ns ± 0%     131ns ± 2%  +1.80%  (p=0.000 n=16+19)
    FmtFprintfInt-12             121ns ± 2%     121ns ± 2%    ~     (p=0.768 n=19+19)
    FmtFprintfIntInt-12          186ns ± 1%     188ns ± 3%  +0.99%  (p=0.000 n=19+19)
    FmtFprintfPrefixedInt-12     188ns ± 1%     188ns ± 1%    ~     (p=0.947 n=18+16)
    FmtFprintfFloat-12           254ns ± 1%     255ns ± 1%  +0.30%  (p=0.002 n=19+17)
    FmtManyArgs-12               763ns ± 0%     770ns ± 0%  +0.92%  (p=0.000 n=18+18)
    GobDecode-12                7.00ms ± 1%    7.04ms ± 1%  +0.61%  (p=0.049 n=20+20)
    GobEncode-12                5.88ms ± 1%    5.88ms ± 0%    ~     (p=0.641 n=18+19)
    Gzip-12                      214ms ± 1%     215ms ± 1%  +0.43%  (p=0.002 n=18+19)
    Gunzip-12                   37.6ms ± 0%    37.6ms ± 0%  +0.11%  (p=0.015 n=17+18)
    HTTPClientServer-12         76.9µs ± 2%    78.1µs ± 2%  +1.44%  (p=0.000 n=20+18)
    JSONEncode-12               15.2ms ± 2%    15.1ms ± 1%    ~     (p=0.271 n=19+18)
    JSONDecode-12               53.1ms ± 1%    53.3ms ± 0%  +0.49%  (p=0.000 n=18+19)
    Mandelbrot200-12            4.04ms ± 1%    4.03ms ± 0%  -0.33%  (p=0.005 n=18+18)
    GoParse-12                  3.29ms ± 1%    3.28ms ± 1%    ~     (p=0.146 n=16+17)
    RegexpMatchEasy0_32-12      69.9ns ± 3%    69.5ns ± 1%    ~     (p=0.785 n=20+19)
    RegexpMatchEasy0_1K-12       237ns ± 0%     237ns ± 0%    ~     (p=1.000 n=18+18)
    RegexpMatchEasy1_32-12      69.5ns ± 1%    69.2ns ± 1%  -0.44%  (p=0.020 n=16+19)
    RegexpMatchEasy1_1K-12       372ns ± 1%     371ns ± 2%    ~     (p=0.086 n=20+19)
    RegexpMatchMedium_32-12      108ns ± 3%     107ns ± 1%  -1.00%  (p=0.004 n=19+14)
    RegexpMatchMedium_1K-12     34.2µs ± 4%    34.0µs ± 2%    ~     (p=0.380 n=19+20)
    RegexpMatchHard_32-12       1.77µs ± 4%    1.76µs ± 3%    ~     (p=0.558 n=18+20)
    RegexpMatchHard_1K-12       53.4µs ± 4%    52.8µs ± 2%  -1.10%  (p=0.020 n=18+20)
    Revcomp-12                   359ms ± 4%     377ms ± 0%  +5.19%  (p=0.000 n=20+18)
    Template-12                 63.7ms ± 2%    62.9ms ± 2%  -1.27%  (p=0.005 n=18+20)
    TimeParse-12                 316ns ± 2%     313ns ± 1%    ~     (p=0.059 n=20+16)
    TimeFormat-12                329ns ± 0%     331ns ± 0%  +0.39%  (p=0.000 n=16+18)
    [Geo mean]                  51.6µs         51.7µs       +0.18%
    
    Change-Id: I1dce4640c8205d41717943b021039fffea863c57
    Reviewed-on: https://go-review.googlesource.com/21324
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6002e01e34f5b847eb4c49ca84e9623d7204f5f2
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 30 17:02:23 2016 -0400

    runtime: allocate black during GC
    
    Currently we allocate white for most of concurrent marking. This is
    based on the classical argument that it produces less floating
    garbage, since allocations during GC may not get linked into the heap
    and allocating white lets us reclaim these. However, it's not clear
    how often this actually happens, especially since our write barrier
    shades any pointer as soon as it's installed in the heap regardless of
    the color of the slot.
    
    On the other hand, allocating black has several advantages that seem
    to significantly outweigh this downside.
    
    1) It naturally bounds the total scan work to the live heap size at
    the start of a GC cycle. Allocating white does not, and thus depends
    entirely on assists to prevent the heap from growing faster than it
    can be scanned.
    
    2) It reduces the total amount of scan work per GC cycle by the size
    of newly allocated objects that are linked into the heap graph, since
    objects allocated black never need to be scanned.
    
    3) It reduces total write barrier work since more objects will already
    be black when they are linked into the heap graph.
    
    This gives a slight overall improvement in benchmarks.
    
    name              old time/op  new time/op  delta
    XBenchGarbage-12  2.24ms ± 0%  2.21ms ± 1%  -1.32%  (p=0.000 n=18+17)
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.60s ± 3%     2.53s ± 3%  -2.56%  (p=0.000 n=20+20)
    Fannkuch11-12                2.08s ± 1%     2.08s ± 0%    ~     (p=0.452 n=19+19)
    FmtFprintfEmpty-12          45.1ns ± 2%    45.3ns ± 2%    ~     (p=0.367 n=19+20)
    FmtFprintfString-12          131ns ± 3%     129ns ± 0%  -1.60%  (p=0.000 n=20+16)
    FmtFprintfInt-12             122ns ± 0%     121ns ± 2%  -0.86%  (p=0.000 n=16+19)
    FmtFprintfIntInt-12          187ns ± 1%     186ns ± 1%    ~     (p=0.514 n=18+19)
    FmtFprintfPrefixedInt-12     189ns ± 0%     188ns ± 1%  -0.54%  (p=0.000 n=16+18)
    FmtFprintfFloat-12           256ns ± 0%     254ns ± 1%  -0.43%  (p=0.000 n=17+19)
    FmtManyArgs-12               769ns ± 0%     763ns ± 0%  -0.72%  (p=0.000 n=18+18)
    GobDecode-12                7.08ms ± 2%    7.00ms ± 1%  -1.22%  (p=0.000 n=20+20)
    GobEncode-12                5.88ms ± 0%    5.88ms ± 1%    ~     (p=0.406 n=18+18)
    Gzip-12                      214ms ± 0%     214ms ± 1%    ~     (p=0.103 n=17+18)
    Gunzip-12                   37.6ms ± 0%    37.6ms ± 0%    ~     (p=0.563 n=17+17)
    HTTPClientServer-12         77.2µs ± 3%    76.9µs ± 2%    ~     (p=0.606 n=20+20)
    JSONEncode-12               15.1ms ± 1%    15.2ms ± 2%    ~     (p=0.138 n=19+19)
    JSONDecode-12               53.3ms ± 1%    53.1ms ± 1%  -0.33%  (p=0.000 n=19+18)
    Mandelbrot200-12            4.04ms ± 1%    4.04ms ± 1%    ~     (p=0.075 n=19+18)
    GoParse-12                  3.30ms ± 1%    3.29ms ± 1%  -0.57%  (p=0.000 n=18+16)
    RegexpMatchEasy0_32-12      69.5ns ± 1%    69.9ns ± 3%    ~     (p=0.822 n=18+20)
    RegexpMatchEasy0_1K-12       237ns ± 1%     237ns ± 0%    ~     (p=0.398 n=19+18)
    RegexpMatchEasy1_32-12      69.8ns ± 2%    69.5ns ± 1%    ~     (p=0.090 n=20+16)
    RegexpMatchEasy1_1K-12       371ns ± 1%     372ns ± 1%    ~     (p=0.178 n=19+20)
    RegexpMatchMedium_32-12      108ns ± 2%     108ns ± 3%    ~     (p=0.124 n=20+19)
    RegexpMatchMedium_1K-12     33.9µs ± 2%    34.2µs ± 4%    ~     (p=0.309 n=20+19)
    RegexpMatchHard_32-12       1.75µs ± 2%    1.77µs ± 4%  +1.28%  (p=0.018 n=19+18)
    RegexpMatchHard_1K-12       52.7µs ± 1%    53.4µs ± 4%  +1.23%  (p=0.013 n=15+18)
    Revcomp-12                   354ms ± 1%     359ms ± 4%  +1.27%  (p=0.043 n=20+20)
    Template-12                 63.6ms ± 2%    63.7ms ± 2%    ~     (p=0.654 n=20+18)
    TimeParse-12                 313ns ± 1%     316ns ± 2%  +0.80%  (p=0.014 n=17+20)
    TimeFormat-12                332ns ± 0%     329ns ± 0%  -0.66%  (p=0.000 n=16+16)
    [Geo mean]                  51.7µs         51.6µs       -0.09%
    
    Change-Id: I2214a6a0e4f544699ea166073249a8efdf080dc0
    Reviewed-on: https://go-review.googlesource.com/21323
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 64a26b79ac781118d4fa364f884ce8080ba97870
Author: Austin Clements <austin@google.com>
Date:   Sun Apr 17 11:42:37 2016 -0400

    runtime: simplify/optimize allocate-black a bit
    
    Currently allocating black switches to the system stack (which is
    probably a historical accident) and atomically updates the global
    bytes marked stat. Since we're about to depend on this much more,
    optimize it a bit by putting it back on the regular stack and updating
    the per-P bytes marked stat, which gets lazily folded into the global
    bytes marked stat.
    
    Change-Id: Ibbe16e5382d3fd2256e4381f88af342bf7020b04
    Reviewed-on: https://go-review.googlesource.com/22170
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 479501c14c9d36e27727bc4b4294d57c5ddc29d0
Author: Austin Clements <austin@google.com>
Date:   Sat Apr 16 18:27:38 2016 -0400

    runtime: count black allocations toward scan work
    
    Currently we count black allocations toward the scannable heap size,
    but not toward the scan work we've done so far. This is clearly
    inconsistent (we have, in effect, scanned these allocations and since
    they're already black, we're not going to scan them again). Worse, it
    means we don't count black allocations toward the scannable heap size
    as of the *next* GC because this is based on the amount of scan work
    we did in this cycle.
    
    Fix this by counting black allocations as scan work. Currently the GC
    spends very little time in allocate-black mode, so this probably
    hasn't been a problem, but this will become important when we switch
    to always allocating black.
    
    Change-Id: If6ff693b070c385b65b6ecbbbbf76283a0f9d990
    Reviewed-on: https://go-review.googlesource.com/22119
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a683c385ad874b0066787dc010cacba8aaff894c
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Fri Jan 29 16:16:03 2016 +0100

    testing: add matching of subtest
    
    Allows passing regexps per subtest to --test.run and --test.bench
    
    Note that the documentation explicitly states that the split regular
    expressions match the correpsonding parts (path components) of
    the bench/test identifier. This is intended and slightly different
    from the i'th RE matching the subtest/subbench at the respective
    level.  Picking this semantics allows guaranteeing that a test or
    benchmark identifier as printed by go test can be passed verbatim
    (possibly quoted) to, respectively, -run or -bench: subtests and
    subbenches might have a '/' in their name, causing a misaligment if
    their ID is passed to -run or -bench as is.
    This semantics has other benefits, but this is the main motivation.
    
    Fixes golang.go#15126
    
    Change-Id: If72e6d3f54db1df6bc2729ac6edc7ab3c740e7c3
    Reviewed-on: https://go-review.googlesource.com/19122
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f8fc3710fd4c596adac57048f705a994f199df8c
Author: Keith Randall <khr@golang.org>
Date:   Thu Apr 21 10:02:36 2016 -0700

    cmd/compile: handle mem copies in amd64 backend
    
    Fixes noopt builder.
    
    Change-Id: If13373b2597f0fcc9b1b2f9c860f2bd043e43c6c
    Reviewed-on: https://go-review.googlesource.com/22338
    Reviewed-by: Keith Randall <khr@golang.org>

commit 508a424eedccfe77f64d50c9870988a8c15b46b1
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Apr 18 10:30:20 2016 -0400

    cmd/compile/internal/gc: fix return value offset for SSA backend on ARM
    
    Progress on SSA backend for ARM. Still not complete. It compiles a
    Fibonacci function, but the caller picked the return value from an
    incorrect offset. This CL adjusts it to match the stack frame layout
    for architectures with link register.
    
    Updates #15365.
    
    Change-Id: I01e03c3e95f5503a185e8ac2b6d9caf4faf3d014
    Reviewed-on: https://go-review.googlesource.com/22186
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7c6b48ffba9e0ea8ed846d194fe30189863f17f0
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Apr 18 12:21:51 2016 -0400

    cmd/compile/internal/arm: fix comparison & conditional branch for SSA on ARM
    
    Progress on SSA for ARM. Still not complete. Now Fibonacci function compiles
    and runs correctly.
    
    The old backend swaps the operands for CMP instruction. This CL does the same
    on SSA backend, and uses conditional branch accordingly.
    
    Updates #15365.
    
    Change-Id: I117e17feb22f03d936608bd232f76970e4bbe21a
    Reviewed-on: https://go-review.googlesource.com/22187
    Reviewed-by: Keith Randall <khr@golang.org>
```
