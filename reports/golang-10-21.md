# October 21, 2016 Report

Number of commits: 156

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 579.644959ms to 601.629093ms, +3.79%
* github.com/gogits/gogs: from 12.652254686s to 12.442620232s, -1.66%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 33.372342201s to 34.832100212s, +4.37%
* github.com/influxdata/influxdb/cmd/influxd: from 6.59264148s to 6.836362096s, +3.70%
* github.com/junegunn/fzf/src/fzf: from 1.097065732s to 1.086517969s, -0.96%
* github.com/mholt/caddy/caddy: from 6.219022269s to 5.788373213s, -6.92%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.421859813s to 1.415578496s, -0.44%
* github.com/nsqio/nsq/apps/nsqd: from 2.100267782s to 2.081137521s, -0.91%
* github.com/prometheus/prometheus/cmd/prometheus: from 39.69239001s to 40.57783638s, +2.23%
* github.com/spf13/hugo: from 7.812349237s to 7.661181821s, -1.93%
* golang.org/x/tools/cmd/guru: from 2.693571838s to 2.969812294s, +10.26%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2542636 to 2538814, -0.15%
* github.com/gogits/gogs: from 23264695 to 23394350, +0.56%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 29292824 to 29456816, +0.56%
* github.com/influxdata/influxdb/cmd/influxd: from 16296754 to 16387710, +0.56%
* github.com/junegunn/fzf/src/fzf: from 3083032 to 3087488, +0.14%
* github.com/mholt/caddy/caddy: from 14715066 to 14810042, +0.65%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4297359 to 4297905, ~
* github.com/nsqio/nsq/apps/nsqd: from 9797434 to 9892209, +0.97%
* github.com/prometheus/prometheus/cmd/prometheus: from 56922356 to 57021583, +0.17%
* github.com/spf13/hugo: from 16209675 to 16304772, +0.59%
* golang.org/x/tools/cmd/guru: from 7332675 to 7341299, +0.12%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               189           189           +0.00%
BenchmarkMsgpUnmarshal-4             373           376           +0.80%
BenchmarkEasyJsonMarshal-4           1436          1466          +2.09%
BenchmarkEasyJsonUnmarshal-4         1792          1422          -20.65%
BenchmarkFlatBuffersMarshal-4        395           625           +58.23%
BenchmarkFlatBuffersUnmarshal-4      271           273           +0.74%
BenchmarkGogoprotobufMarshal-4       155           154           -0.65%
BenchmarkGogoprotobufUnmarshal-4     240           245           +2.08%

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

* [strings: use Index in Count](https://github.com/golang/go/commit/6347367be36df608cce84beb097378f8654dd208)
* [net/http: make Server Handler's Request.Context be done on conn errors](https://github.com/golang/go/commit/faf882d1d427e8c8a9a1be00d8ddcab81d1e848e)
* [syscall: for ForkExec on Linux, always use 32-bit setgroups system call](https://github.com/golang/go/commit/6c295a9a71924478a344e7b447ff3b44b1e94511)
* [sync: throw, not panic, for unlock of unlocked mutex](https://github.com/golang/go/commit/40d81cf061d8a2a277d70446f582a984c1701ff3)


## GIT Log

```
commit 325c2aa5b6f12ccc82e8472beec0e1511e975bd4
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Mon Oct 17 09:28:06 2016 -0700

    database/sql: update the conversion errors to be clearer
    
    There was some ambiguity over which argument was referred to when
    a conversion error was returned. Now refer to the argument by
    either explicit ordinal position or name if present.
    
    Fixes #15676
    
    Change-Id: Id933196b7e648baa664f4121fa3fb1b07b3c4880
    Reviewed-on: https://go-review.googlesource.com/31262
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 930ab0afd787214d379aca230944d0d41c8b90e6
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Oct 14 11:37:55 2016 -0400

    cmd/asm, cmd/internal/obj/s390x: fix VFMA and VFMS encoding
    
    The m5 and m6 fields were the wrong way round.
    
    Fixes #17444.
    
    Change-Id: I10297064f2cd09d037eac581c96a011358f70aae
    Reviewed-on: https://go-review.googlesource.com/31130
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 0f29942489409ccd81619b5f82fce9c7de18165f
Author: David Chase <drchase@google.com>
Date:   Wed Oct 19 11:47:52 2016 -0400

    cmd/compile: Repurpose old sliceopt.go for prove phase.
    
    Adapt old test for prove's bounds check elimination.
    Added missing rule to generic rules that lead to differences
    between 32 and 64 bit platforms on sliceopt test.
    Added debugging to prove.go that was helpful-to-necessary to
    discover that missing rule.
    Lowered debugging level on prove.go from 3 to 1; no idea
    why it was previously 3.
    
    Change-Id: I09de206aeb2fced9f2796efe2bfd4a59927eda0c
    Reviewed-on: https://go-review.googlesource.com/23290
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit c9517b1ffe7c99a3fd2a748fab51645bd674ad75
Author: Klaus Post <klauspost@gmail.com>
Date:   Sun Oct 16 22:18:01 2016 +0200

    compress/gzip, compress/zlib: add HuffmanOnly as compression levels.
    
    This exposes HuffmanOnly in zlib and gzip packages, which is currently
    unavailable.
    
    Change-Id: If5d103bbc8b5fce2f5d740fd103a235c5d1ed7cd
    Reviewed-on: https://go-review.googlesource.com/31186
    Reviewed-by: Nigel Tao <nigeltao@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7ea58297171ef9ba2680abd5e2490d48c1c9f24c
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Oct 20 15:27:26 2016 -0700

    cmd/compile: a couple of minor comment fixes
    
    Change-Id: If1d08a84c9295816489b1cfdd031ba12892ae963
    Reviewed-on: https://go-review.googlesource.com/31598
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 584e3ea2581c858688803d41e207eeaa881a2a85
Author: Nigel Tao <nigeltao@golang.org>
Date:   Thu Oct 20 16:57:55 2016 +1100

    image/color: tweak the formula for converting to gray.
    
    This makes grayModel and gray16Model in color.go use the exact same
    formula as RGBToYCbCr in ycbcr.go. They were the same formula in theory,
    but in practice the color.go versions used a divide by 1000 and the
    ycbcr.go versions used a (presumably faster) shift by 16.
    
    This implies the nice property that converting an image.RGBA to an
    image.YCbCr and then taking only the Y channel is equivalent to
    converting an image.RGBA directly to an image.Gray.
    
    The difference between the two formulae is non-zero, but small:
    https://play.golang.org/p/qG7oe-eqHI
    
    Updates #16251
    
    Change-Id: I288ecb957fd6eceb9626410bd1a8084d2e4f8198
    Reviewed-on: https://go-review.googlesource.com/31538
    Reviewed-by: Rob Pike <r@golang.org>

commit a190f3c8a34b859dd578b1d30b48ecd04c7a99c7
Author: David Chase <drchase@google.com>
Date:   Wed May 11 15:25:17 2016 -0400

    cmd/compile: enable flag-specified dump of specific phase+function
    
    For very large input files, use of GOSSAFUNC to obtain a dump
    after compilation steps can lead to both unwieldy large output
    files and unwieldy larger processes (because the output is
    buffered in a string).  This flag
    
      -d=ssa/<phase>/dump:<function name>
    
    provides finer control of what is dumped, into a smaller
    file, and with less memory overhead in the running compiler.
    The special phase name "build" is added to allow printing
    of the just-built ssa before any transformations are applied.
    
    This was helpful in making sense of the gogo/protobuf
    problems.
    
    The output format was tweaked to remove gratuitous spaces,
    and a crude -d=ssa/help help text was added.
    
    Change-Id: If7516e22203420eb6ed3614f7cee44cb9260f43e
    Reviewed-on: https://go-review.googlesource.com/23044
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 10560afb540b783da568aebe83d0f782e46bb673
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Thu Oct 20 11:24:51 2016 +0200

    runtime/debug: avoid overflow in SetMaxThreads
    
    Fixes #16076
    
    Change-Id: I91fa87b642592ee4604537dd8c3197cd61ec8b31
    Reviewed-on: https://go-review.googlesource.com/31516
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f6f3aef53f7ef6085ea14b6147b2478848778709
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Oct 20 15:22:46 2016 -0400

    go/build: reserve GOOS=zos for IBM z/OS
    
    Closes #17528.
    
    Change-Id: I2ba55ad4e41077808f882ed67a0549f0a00e25d0
    Reviewed-on: https://go-review.googlesource.com/31596
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ca8dd033a0a88946f02a37bdbd107af83afa0e24
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Oct 20 11:10:03 2016 -0700

    cmd/cgo: correct comment on Package.rewriteCall
    
    Account for changes in https://golang.org/cl/31233.
    
    Change-Id: I3311c6850a3c714d18209fdff500dd817e9dfcb2
    Reviewed-on: https://go-review.googlesource.com/31594
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit abcf8457b30ef50b878b2a1add39fa5f5a2f1b60
Author: Adam Langley <agl@golang.org>
Date:   Thu Oct 20 09:35:19 2016 -0700

    vendor/golang_org/x/crypto/poly1305: update to 3ded668c5379f6951fb0de06174442072e5447d3
    
    This change updates the vendored copy of x/crypto/poly1305, specifically
    to include the following changes:
      3ded668 poly1305: enable assembly for ARM in Go 1.6.
      dec8741 poly1305: fix stack handling in sum_arm.s
    
    Fixes #17499.
    
    Change-Id: I8f152da9599bd15bb976f630b0ef602be05143d3
    Reviewed-on: https://go-review.googlesource.com/31592
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 33b71dfa1c993d2e06d7723e3205c792df3b383d
Author: David Chase <drchase@google.com>
Date:   Thu Oct 20 11:05:45 2016 -0400

    cmd/compile: add patterns to improve PPC64 FP comparisons
    
    Uncommented 4 comparison rules of this form:
    (NE (CMPWconst [0] (FLessThan cc)) yes no) -> (FLT cc yes no)
    
    Fixes #17507.
    
    Change-Id: I74f34f13526aeee619711c8281a66652d90a962a
    Reviewed-on: https://go-review.googlesource.com/31612
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit f9bbfe4a093144e48a4774d1388c1e99be64673b
Author: David du Colombier <0intro@gmail.com>
Date:   Thu Oct 20 16:51:03 2016 +0200

    net/http: remove workaround in TestTransportClosesBodyOnError on Plan 9
    
    This issue has been fixed in CL 31390.
    
    Fixes #9554.
    
    Change-Id: Ib8ff4cb1ffcb7cdbf117510b98b4a7e13e4efd2b
    Reviewed-on: https://go-review.googlesource.com/31520
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8589049d5c1cc1ff9856defb6d9097428ad09b82
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Wed Oct 19 20:21:42 2016 +0300

    cmd/compile/internal/amd64: break dependency for CVTS[LQ]2S[DS]
    
    CVTSL2SS, CVTSQ2SS, CVTSL2SD, CVTSQ2SD preserve upper part of xmm register,
    introducing false dependency on a previous value.
    Break it by xoring destination with itself.
    Increases size of go executable by 320 bytes, but shows nice improvement on go1.
    Also fixes performance degradation introduced by 1.7.
    
    name                     old time/op    new time/op    delta
    BinaryTree17-4              2.20s ± 1%     2.19s ± 0%  -0.36%        (p=0.000 n=18+16)
    Fannkuch11-4                2.44s ± 1%     2.45s ± 2%  +0.47%        (p=0.030 n=20+20)
    FmtFprintfEmpty-4          40.9ns ± 7%    40.5ns ± 1%    ~           (p=0.531 n=20+16)
    FmtFprintfString-4          111ns ± 2%     111ns ± 1%    ~           (p=0.510 n=18+19)
    FmtFprintfInt-4            98.3ns ± 3%    99.3ns ± 1%  +1.01%        (p=0.003 n=20+18)
    FmtFprintfIntInt-4          148ns ± 3%     147ns ± 1%    ~           (p=0.919 n=20+17)
    FmtFprintfPrefixedInt-4     149ns ± 1%     152ns ± 0%  +1.73%        (p=0.000 n=19+17)
    FmtFprintfFloat-4           231ns ± 0%     231ns ± 1%    ~           (p=0.678 n=18+19)
    FmtManyArgs-4               667ns ± 1%     672ns ± 1%  +0.73%        (p=0.005 n=20+20)
    GobDecode-4                5.60ms ± 0%    5.61ms ± 0%  +0.24%        (p=0.000 n=20+20)
    GobEncode-4                4.74ms ± 0%    4.73ms ± 1%  -0.20%        (p=0.002 n=20+20)
    Gzip-4                      199ms ± 0%     199ms ± 1%  +0.35%        (p=0.000 n=19+20)
    Gunzip-4                   31.8ms ± 1%    31.5ms ± 1%  -0.89%        (p=0.000 n=20+20)
    HTTPClientServer-4         38.1µs ± 1%    38.0µs ± 1%    ~           (p=0.117 n=19+18)
    JSONEncode-4               14.2ms ± 1%    13.4ms ± 0%  -5.73%        (p=0.000 n=20+20)
    JSONDecode-4               42.7ms ± 0%    42.7ms ± 1%  +0.18%        (p=0.019 n=18+19)
    Mandelbrot200-4            3.26ms ± 0%    2.99ms ± 0%  -8.38%        (p=0.000 n=19+19)
    GoParse-4                  2.76ms ± 1%    2.76ms ± 1%    ~           (p=0.583 n=20+20)
    RegexpMatchEasy0_32-4      69.5ns ± 0%    69.6ns ± 0%  +0.10%        (p=0.017 n=16+17)
    RegexpMatchEasy0_1K-4       703ns ± 0%     708ns ± 3%  +0.65%        (p=0.000 n=17+18)
    RegexpMatchEasy1_32-4      68.2ns ± 1%    68.2ns ± 2%    ~           (p=0.094 n=18+20)
    RegexpMatchEasy1_1K-4       288ns ± 1%     288ns ± 0%    ~           (p=0.403 n=17+18)
    RegexpMatchMedium_32-4      104ns ± 2%     103ns ± 1%    ~           (p=0.110 n=20+16)
    RegexpMatchMedium_1K-4     31.7µs ± 3%    31.7µs ± 3%    ~           (p=0.091 n=19+20)
    RegexpMatchHard_32-4       1.59µs ± 2%    1.58µs ± 2%    ~           (p=0.083 n=20+20)
    RegexpMatchHard_1K-4       48.1µs ± 3%    47.9µs ± 2%    ~           (p=0.461 n=20+19)
    Revcomp-4                   344ms ± 0%     345ms ± 0%  +0.08%        (p=0.009 n=18+17)
    Template-4                 44.8ms ± 1%    44.7ms ± 1%    ~           (p=0.277 n=20+20)
    TimeParse-4                 258ns ± 0%     258ns ± 0%    ~     (all samples are equal)
    TimeFormat-4                275ns ± 0%     273ns ± 0%  -0.64%        (p=0.000 n=20+18)
    
    name                     old speed      new speed      delta
    GobDecode-4               137MB/s ± 0%   137MB/s ± 0%  -0.24%        (p=0.000 n=20+20)
    GobEncode-4               162MB/s ± 0%   162MB/s ± 0%  +0.20%        (p=0.002 n=20+20)
    Gzip-4                   97.6MB/s ± 0%  97.3MB/s ± 1%  -0.35%        (p=0.000 n=19+20)
    Gunzip-4                  610MB/s ± 1%   615MB/s ± 1%  +0.89%        (p=0.000 n=20+20)
    JSONEncode-4              136MB/s ± 1%   145MB/s ± 0%  +6.08%        (p=0.000 n=20+20)
    JSONDecode-4             45.5MB/s ± 0%  45.4MB/s ± 1%  -0.17%        (p=0.017 n=18+19)
    GoParse-4                21.0MB/s ± 1%  21.0MB/s ± 1%    ~           (p=0.578 n=20+20)
    RegexpMatchEasy0_32-4     460MB/s ± 0%   460MB/s ± 0%  -0.09%        (p=0.031 n=16+17)
    RegexpMatchEasy0_1K-4    1.46GB/s ± 0%  1.45GB/s ± 3%  -0.64%        (p=0.000 n=17+18)
    RegexpMatchEasy1_32-4     469MB/s ± 0%   469MB/s ± 2%  +0.06%        (p=0.043 n=18+20)
    RegexpMatchEasy1_1K-4    3.55GB/s ± 1%  3.55GB/s ± 0%    ~           (p=0.057 n=17+18)
    RegexpMatchMedium_32-4   9.61MB/s ± 2%  9.64MB/s ± 2%    ~           (p=0.856 n=20+20)
    RegexpMatchMedium_1K-4   32.3MB/s ± 3%  32.3MB/s ± 3%    ~           (p=0.085 n=19+20)
    RegexpMatchHard_32-4     20.1MB/s ± 2%  20.2MB/s ± 2%    ~           (p=0.086 n=20+20)
    RegexpMatchHard_1K-4     21.3MB/s ± 3%  21.4MB/s ± 2%    ~           (p=0.578 n=20+20)
    Revcomp-4                 738MB/s ± 0%   737MB/s ± 0%  -0.08%        (p=0.009 n=18+17)
    Template-4               43.3MB/s ± 1%  43.4MB/s ± 1%    ~           (p=0.274 n=20+20)
    
    Fixes #16982
    
    Change-Id: If574d66f39f4183a9b1d5ffff0339909cc73f59d
    Reviewed-on: https://go-review.googlesource.com/31490
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 5ce715fdfea635c2cb429166294fb005cefde896
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Wed Oct 19 20:35:04 2016 +0300

    cmd/internal/obj/x86: add some missing AMD64 instructions
    
    Add VBROADCASTSD, BROADCASTSS, MOVDDUP, MOVSHDUP, MOVSLDUP,
    VMOVDDUP, VMOVSHDUP, VMOVSLDUP.
    
    Fixes #16007
    
    Change-Id: I9614e58eed6c1b6f299d9b4f0b1a7750aa7c1725
    Reviewed-on: https://go-review.googlesource.com/31491
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 692df217ca21b6df8e4dc65538fcc90733e8900e
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Wed Oct 19 12:26:55 2016 -0700

    database/sql: add missing unlock when context is expired
    
    Missing the DB mutex unlock on an early return after checking
    if the context has expired.
    
    Fixes #17518
    
    Change-Id: I247cafcef62623d813f534a941f3d5a3744f0738
    Reviewed-on: https://go-review.googlesource.com/31494
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9cefbe9d03918c07451e5512cc4c227fb0d6885a
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Thu Oct 20 08:59:07 2016 +0200

    lib/time: update to IANA release 2016g (September 2016)
    
    Change-Id: Ie6258602554c5bb6685c9de42ccda84d297af1e2
    Reviewed-on: https://go-review.googlesource.com/31515
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ef8e85e8b8176fb096683dac00cf536fd7d228bc
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Oct 18 16:38:54 2016 -0700

    archive/tar: fix parsePAXTime
    
    Issues fixed:
    * Could not handle quantity of seconds greater than 1<<31 on
    32bit machines since strconv.ParseInt did not treat integers as 64b.
    * Did not handle negative timestamps properly if nanoseconds were used.
    Note that "-123.456" should result in a call to time.Unix(-123, -456000000).
    * Incorrectly allowed a '-' right after the '.' (e.g., -123.-456)
    * Did not detect invalid input after the truncation point (e.g., 123.123456789badbadbad).
    
    Note that negative timestamps are allowed by PAX, but are not guaranteed
    to be portable. See the relevant specification:
    <<<
    If pax encounters a file with a negative timestamp in copy or write mode,
    it can reject the file, substitute a non-negative timestamp, or generate
    a non-portable timestamp with a leading '-'.
    >>>
    
    Since the previous behavior already partially supported negative timestamps,
    we are bound by Go's compatibility rules to keep support for them.
    However, we should at least make sure we handle them properly.
    
    Change-Id: I5686997708bfb59110ea7981175427290be737d1
    Reviewed-on: https://go-review.googlesource.com/31441
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f35b8658dca82f317251f7947964fb2878b57a85
Author: Quentin Smith <quentin@golang.org>
Date:   Tue Oct 18 15:54:04 2016 -0400

    doc: update install-source.html for SSA in Go 1.8
    
    Fixes #17491
    
    Change-Id: Ic070cbed60fa893fed568e8fac448b86cd3e0cbc
    Reviewed-on: https://go-review.googlesource.com/31411
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit f1ad4863aae1fd5cd5d0e3e4e6cb6bfae62951a6
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Oct 17 17:10:24 2016 -0400

    runtime: get s390x vector facility availability from AT_HWCAP
    
    This is a more robust method for obtaining the availability of vx.
    Since this variable may be checked frequently I've also now
    padded it so that it will be in its own cache line.
    
    I've kept the other check (in hash/crc32) the same for now until
    I can figure out the best way to update it.
    
    Updates #15403.
    
    Change-Id: I74eed651afc6f6a9c5fa3b88fa6a2b0c9ecf5875
    Reviewed-on: https://go-review.googlesource.com/31149
    Reviewed-by: Austin Clements <austin@google.com>

commit 2be3ab441578413972d93fc1048decc422640b09
Author: Austin Clements <austin@google.com>
Date:   Sun Oct 16 21:22:02 2016 -0400

    runtime: keep gcMarkRootCheck happy with spare Gs
    
    oneNewExtraM creates a spare M and G for use with cgo callbacks. The G
    doesn't run right away, but goes directly into syscall status. For the
    garbage collector, it's marked as "scan valid" and not on the rescan
    list, but I forgot to also mark it as "scan done". As a result,
    gcMarkRootCheck thinks that the goroutine hasn't been scanned and
    panics.
    
    This only affects GODEBUG=gccheckmark=1 mode, since we otherwise skip
    the gcMarkRootCheck.
    
    Fixes #17473.
    
    Change-Id: I94f5671c42eb44bd5ea7dc68fbf85f0c19e2e52c
    Reviewed-on: https://go-review.googlesource.com/31139
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 7bc42a145a03d61b504b18d20f2a4e5d8e2436fd
Author: Austin Clements <austin@google.com>
Date:   Mon Sep 12 17:48:34 2016 -0400

    runtime: don't reserve space for stack barriers if they're off
    
    Change-Id: I79ebccdaefc434c47b77bd545cc3c50723c18b61
    Reviewed-on: https://go-review.googlesource.com/31135
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 5b7497f327f83510193b1ec1de2eabb287a02982
Author: Austin Clements <austin@google.com>
Date:   Sun Sep 11 20:03:14 2016 -0400

    runtime: update heap profile stats after world is started
    
    Updating the heap profile stats is one of the most expensive parts of
    mark termination other than stack rescanning, but there's really no
    need to do this with the world stopped. Move it to right after we've
    started the world back up. This creates a *very* small window where
    allocations from the next cycle can slip into the profile, but the
    exact point where mark termination happens is so non-deterministic
    already that a slight reordering here is unimportant.
    
    Change-Id: I2f76f22c70329923ad6a594a2c26869f0736d34e
    Reviewed-on: https://go-review.googlesource.com/31363
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 9429aab9999e00958abd8b21d06fa4a2253437c2
Author: Austin Clements <austin@google.com>
Date:   Sun Sep 11 16:55:34 2016 -0400

    runtime: remove gcWork flushes in mark termination
    
    The only reason these flushes are still necessary at all is that
    gcmarknewobject doesn't flush its gcWork stats like it's supposed to.
    By changing gcmarknewobject to follow the standard protocol, the
    flushes become completely unnecessary because mark 2 ensures caches
    are flushed (and stay flushed) before we ever enter mark termination.
    
    In the garbage benchmark, this takes roughly 50 µs, which is
    surprisingly long for doing nothing. We still double-check after
    draining that they are in fact empty.
    
    Change-Id: Ia1c7cf98a53f72baa513792eb33eca6a0b4a7128
    Reviewed-on: https://go-review.googlesource.com/31134
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit a16954b8a7d66169760fb60dd7f3d4e400a5e98c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Oct 14 14:53:59 2016 -0700

    cmd/cgo: always use a function literal for pointer checking
    
    The pointer checking code needs to know the exact type of the parameter
    expected by the C function, so that it can use a type assertion to
    convert the empty interface returned by cgoCheckPointer to the correct
    type. Previously this was done by using a type conversion, but that
    meant that the code accepted arguments that were convertible to the
    parameter type, rather than arguments that were assignable as in a
    normal function call. In other words, some code that should not have
    passed type checking was accepted.
    
    This CL changes cgo to always use a function literal for pointer
    checking. Now the argument is passed to the function literal, which has
    the correct argument type, so type checking is performed just as for a
    function call as it should be.
    
    Since we now always use a function literal, simplify the checking code
    to run as a statement by itself. It now no longer needs to return a
    value, and we no longer need a type assertion.
    
    This does have the cost of introducing another function call into any
    call to a C function that requires pointer checking, but the cost of the
    additional call should be minimal compared to the cost of pointer
    checking.
    
    Fixes #16591.
    
    Change-Id: I220165564cf69db9fd5f746532d7f977a5b2c989
    Reviewed-on: https://go-review.googlesource.com/31233
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e32ac7978df02fae0cbbd92bb65d0d50ea4d2df5
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Oct 19 12:47:59 2016 -0700

    cmd/link, cmd/internal/obj: stop exporting various names
    
    Just happened to notice that these names (funcAlign and friends) are
    never referenced outside their package, so no need to export them.
    
    Change-Id: I4bbdaa4b0ef330c3c3ef50a2ca39593977a83545
    Reviewed-on: https://go-review.googlesource.com/31496
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 832082b44e2965c136e53a9b8009d2e860766000
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Oct 19 12:58:16 2016 -0700

    cmd/compile: remove -A flag
    
    mkbuiltin.go now generates builtin.go using go/ast instead of running
    the compiler, so we don't need the -A flag anymore.
    
    Passes toolstash -cmp.
    
    Change-Id: Ifa70f4f3c9feae10c723cbec81a0a47c39610090
    Reviewed-on: https://go-review.googlesource.com/31497
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 42b37819a132fc2e79149643691894f501077161
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 18 16:11:50 2016 -0700

    cmd/compile: rework mkbuiltin.go to generate code
    
    Generating binary export data requires a working Go compiler. Even
    trickier to change the export data format itself requires a careful
    bootstrapping procedure.
    
    Instead, simply generate normal Go code that lets us directly
    construct the builtin runtime declarations.
    
    Passes toolstash -cmp.
    
    Fixes #17508.
    
    Change-Id: I4f6078a3c7507ba40072580695d57c87a5604baf
    Reviewed-on: https://go-review.googlesource.com/31493
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit abdd73cc43f9187e8918879944ec0dacbc912b5c
Author: Edward Muller <edwardam@interlix.com>
Date:   Tue Oct 4 21:24:58 2016 -0700

    net/http/httptrace: add ClientTrace.TLSHandshakeStart & TLSHandshakeDone
    
    Fixes #16965
    
    Change-Id: I3638fe280a5b1063ff589e6e1ff8a97c74b77c66
    Reviewed-on: https://go-review.googlesource.com/30359
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 12c9844cc6b7b9396bad4ceccfe93874b43b3c72
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Oct 18 16:57:02 2016 -0700

    archive/tar: fix parsePAX to be POSIX.1-2001 compliant
    
    Relevant PAX specification:
    <<<
    If the <value> field is zero length, it shall delete any header
    block field, previously entered extended header value, or
    global extended header value of the same name.
    >>>
    
    We don't delete global extender headers since the Reader doesn't
    even support global headers (which the specification admits was
    a controversial feature).
    
    Change-Id: I2125a5c907b23a3dc439507ca90fa5dc47d474a9
    Reviewed-on: https://go-review.googlesource.com/31440
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 04262986a0b112de5f6f0b287447319c31ef15f9
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Thu Oct 13 17:14:24 2016 -0700

    archive/tar: compact slices in tests
    
    Took this opportunity to also embed tables in the functions
    that they are actually used in and other stylistic cleanups.
    
    There was no logical changes to the tests.
    
    Change-Id: Ifa724060532175f6f4407d6cedc841891efd8f7b
    Reviewed-on: https://go-review.googlesource.com/31436
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5378dd77684e2eee5f05aab4b77497bb635fd544
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 19 09:11:16 2016 -0400

    text/template: add support for reflect.Value args, results in funcs
    
    Add support for passing reflect.Values to and returning reflect.Values from
    any registered functions in the FuncMap, much as if they were
    interface{} values. Keeping the reflect.Value instead of round-tripping
    to interface{} preserves addressability of the value, which is important
    for method lookup.
    
    Change index and a few other built-in functions to use reflect.Values,
    making a loop using explicit indexing now match the semantics that
    range has always had.
    
    Fixes #14916.
    
    Change-Id: Iae1a2fd9bb426886a7fcd9204f30a2d6ad4646ad
    Reviewed-on: https://go-review.googlesource.com/31462
    Reviewed-by: Rob Pike <r@golang.org>

commit d2aa8601b5aafa9736f8e49ca713ecc31f9a011f
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Oct 18 17:51:04 2016 -0700

    archive/tar: make Reader handle GNU format properly
    
    The GNU format does not have a prefix field, so we should make
    no attempt to read it. It does however have atime and ctime fields.
    Since Go previously placed incorrect values here, we liberally
    read the atime and ctime fields and ignore errors so that old tar
    files written by Go can at least be partially read.
    
    This fixes half of #12594. The Writer is much harder to fix.
    
    Updates #12594
    
    Change-Id: Ia32845e2f262ee53366cf41dfa935f4d770c7a30
    Reviewed-on: https://go-review.googlesource.com/31444
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 40d81cf061d8a2a277d70446f582a984c1701ff3
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 10:26:07 2016 -0400

    sync: throw, not panic, for unlock of unlocked mutex
    
    The panic leaves the lock in an unusable state.
    Trying to panic with a usable state makes the lock significantly
    less efficient and scalable (see early CL patch sets and discussion).
    
    Instead, use runtime.throw, which will crash the program directly.
    
    In general throw is reserved for when the runtime detects truly
    serious, unrecoverable problems. This problem is certainly serious,
    and, without a significant performance hit, is unrecoverable.
    
    Fixes #13879.
    
    Change-Id: I41920d9e2317270c6f909957d195bd8b68177f8d
    Reviewed-on: https://go-review.googlesource.com/31359
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d2315fdc11ebdf5c0ae94f33cb01ffaab82c00b6
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 19 13:14:16 2016 -0400

    html/template: adjust ambiguous URL context text
    
    Before: ... appears in an ambiguous URL context.
    After:  ... appears in an ambiguous context within a URL.
    
    It's a minor point, but it's confused multiple people.
    Try to make clearer that the ambiguity is "where exactly inside the URL?"
    
    Fixes #17319.
    
    Change-Id: Id834868d1275578036c1b00c2bdfcd733d9d2b7b
    Reviewed-on: https://go-review.googlesource.com/31465
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 13c65723419a41a8b5339e11f40fe2ecf0bbc8ab
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed Oct 19 11:06:36 2016 -0400

    plugin: mention OS X support and concurrency
    
    Change-Id: I4270bf81511a5bf80ed146f5e66e4f8aeede2aa2
    Reviewed-on: https://go-review.googlesource.com/31463
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 023bb034e93363492ef444fefcb1d38cdc61ede1
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Oct 19 09:56:53 2016 -0700

    spec: slightly more realistic example for type assertions
    
    For #17428.
    
    Change-Id: Ia902b50cf0c40e3c2167fb573a39d328331c38c7
    Reviewed-on: https://go-review.googlesource.com/31449
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 1e3dc3d5d43a835a60e0261e343d3a44b5f93db0
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Oct 19 12:39:37 2016 +0000

    syscall: make Utimes on Solaris match all the other geese
    
    Updates #14892
    
    Change-Id: I640c6e1635ccdf611f219521a7d297a9885c4cb3
    Reviewed-on: https://go-review.googlesource.com/31446
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4859f6a416b053d57fcc9d8f43e81e9d218280e9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Oct 19 10:31:15 2016 +0000

    net/http: make NewRequest set empty Body nil, don't peek Read Body in Transport
    
    This CL makes NewRequest set Body nil for known-zero bodies, and makes
    the http1 Transport not peek-Read a byte to determine whether there's
    a body.
    
    Background:
    
    Many fields of the Request struct have different meanings for whether
    they're outgoing (via the Transport) or incoming (via the Server).
    
    For outgoing requests, ContentLength and Body are documented as:
    
            // Body is the request's body.
            //
            // For client requests a nil body means the request has no
            // body, such as a GET request. The HTTP Client's Transport
            // is responsible for calling the Close method.
            Body io.ReadCloser
    
            // ContentLength records the length of the associated content.
            // The value -1 indicates that the length is unknown.
            // Values >= 0 indicate that the given number of bytes may
            // be read from Body.
            // For client requests, a value of 0 with a non-nil Body is
            // also treated as unknown.
            ContentLength int64
    
    Because of the ambiguity of what ContentLength==0 means, the http1 and
    http2 Transports previously Read the first byte of a non-nil Body when
    the ContentLength was 0 to determine whether there was an actual body
    (with a non-zero length) and ContentLength just wasn't populated, or
    it was actually empty.
    
    That byte-sniff has been problematic and gross (see #17480, #17071)
    and was removed for http2 in a previous commit.
    
    That means, however, that users doing:
    
        req, _ := http.NewRequest("POST", url, strings.NewReader(""))
    
    ... would not send a Content-Length header in their http2 request,
    because the size of the reader (even though it was known, being one of
    the three common recognized types from NewRequest) was zero, and so
    the HTTP Transport thought it was simply unset.
    
    To signal explicitly-zero vs unset-zero, this CL changes NewRequest to
    signal explicitly-zero by setting the Body to nil, instead of the
    strings.NewReader("") or other zero-byte reader.
    
    This CL also removes the byte sniff from the http1 Transport, like
    https://golang.org/cl/31326 did for http2.
    
    Updates #17480
    Updates #17071
    
    Change-Id: I329f02f124659bf7d8bc01e2c9951ebdd236b52a
    Reviewed-on: https://go-review.googlesource.com/31445
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4c1995f95b9786006c71537e34609e356efafb4b
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 11:01:54 2016 -0400

    reflect: document DeepEqual(nil map, empty non-nil map) behavior
    
    Fixes #16531.
    
    Change-Id: I41ec8123f2d3fbe063fd3b09a9366e69722793e5
    Reviewed-on: https://go-review.googlesource.com/31355
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6c295a9a71924478a344e7b447ff3b44b1e94511
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 22:56:14 2016 -0400

    syscall: for ForkExec on Linux, always use 32-bit setgroups system call
    
    Fixes #17092.
    
    Change-Id: If203d802a919e00594ddc1282782fc59a083fd63
    Reviewed-on: https://go-review.googlesource.com/31458
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f2f8d58b92cd7fd9616c98fb012467656de5a3cb
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 19 08:48:25 2016 -0400

    reflect: update FieldByNameFunc comment
    
    This was supposed to be in CL 31354
    but was dropped due to a Git usage error.
    
    For #16573.
    
    Change-Id: I3d99087c8efc8cbc016c55e8365d0005f79d1b2f
    Reviewed-on: https://go-review.googlesource.com/31461
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2f7f679c79009137bd34fcc33a6d3a6762f45e75
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 23:22:38 2016 -0400

    html/template, text/template: clarify template redefinition behavior
    
    Make two important points clearer:
    
     - Giving a template definition containing
       nothing but spaces has no effect.
     - Giving a template definition containing
       non-spaces can only be done once per template.
    
    Fixes #16912.
    Fixes #16913.
    Fixes #17360.
    
    Change-Id: Ie3971b83ab148b7c8bb800fe4a21579566378e3e
    Reviewed-on: https://go-review.googlesource.com/31459
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit ae14472af9e9508eab948303b14290ca98d09646
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 22:21:58 2016 -0400

    os: clean up after test
    
    Noted in CL 31358 after submit.
    
    Change-Id: I76ddad9b9e27dd6a03c1c4f49153213747fe0a61
    Reviewed-on: https://go-review.googlesource.com/31365
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9aed16e96dc97a68a00e6358fe05f96ce5c16b35
Author: Ingo Oeser <nightlyone@googlemail.com>
Date:   Tue Oct 18 18:05:43 2016 +0200

    regexp: avoid alloc in QuoteMeta when not quoting
    
    Many users quote literals in regular expressions just in case.
    No need to allocate then.
    
    Note: Also added benchmarks for quoting and not quoting.
    
            name             old time/op    new time/op     delta
            QuoteMetaAll-4      629ns ± 6%      654ns ± 5%    +4.01%        (p=0.001 n=20+19)
            QuoteMetaNone-4    1.02µs ± 6%     0.20µs ± 0%   -80.73%        (p=0.000 n=18+20)
    
            name             old speed      new speed       delta
            QuoteMetaAll-4   22.3MB/s ± 6%   21.4MB/s ± 5%    -3.94%        (p=0.001 n=20+19)
            QuoteMetaNone-4  25.3MB/s ± 3%  131.5MB/s ± 0%  +419.28%        (p=0.000 n=17+19)
    
            name             old alloc/op   new alloc/op    delta
            QuoteMetaAll-4      64.0B ± 0%      64.0B ± 0%      ~     (all samples are equal)
            QuoteMetaNone-4     96.0B ± 0%      0.0B ±NaN%  -100.00%        (p=0.000 n=20+20)
    
            name             old allocs/op  new allocs/op   delta
            QuoteMetaAll-4       2.00 ± 0%       2.00 ± 0%      ~     (all samples are equal)
            QuoteMetaNone-4      2.00 ± 0%      0.00 ±NaN%  -100.00%        (p=0.000 n=20+20)
    
    Change-Id: I38d50f463cde463115d22534f8eb849e54d899af
    Reviewed-on: https://go-review.googlesource.com/31395
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f36e1adaa2c72d74dc669b596ea1c4df5e938def
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Oct 9 22:45:17 2016 -0700

    math/big: implement Float.Scan, type assert fmt interfaces to enforce docs
    
    Implements Float.Scan which satisfies fmt.Scanner interface.
    Also enforces docs' interface implementation claims with compile time
    type assertions, that is:
    + Float always implements fmt.Formatter and fmt.Scanner
    + Int always implements fmt.Formatter and fmt.Scanner
    + Rat always implements fmt.Formatter
    which will ensure that the API claims are strictly matched.
    
    Also note that Float.Scan doesn't handle ±Inf.
    
    Fixes #17391
    
    Change-Id: I3d3dfbe7f602066975c7a7794fe25b4c645440ce
    Reviewed-on: https://go-review.googlesource.com/30723
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit ead08e91f6468ab1c35c250ec487935103c580f6
Author: Caio Marcelo de Oliveira Filho <caio.oliveira@intel.com>
Date:   Wed Apr 20 14:29:30 2016 -0300

    cmd/go, testing: indicate when no tests are run
    
    For example, testing the current directory:
    
            $ go test -run XXX
            testing: warning: no tests to run
            PASS
            ok      testing 0.013s
            $
    
    And in a summary:
    
            $ go test -run XXX testing
            ok      testing 0.013s [no tests to run]
            $
    
    These make it easy to spot when the -run regexp hasn't matched anything
    or there are no tests. Previously the message was printed in the "current directory"
    case when there were no tests at all, but not for no matches, and either way
    was not surfaced in the directory list summary form.
    
    Fixes #15211.
    
    Change-Id: I1c82a423d6bd429fb991c9ca964c9d26c96fd3c5
    Reviewed-on: https://go-review.googlesource.com/22341
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>

commit 95abb5a36aa1a727db512772f632ecaf05df34aa
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 23:05:27 2016 -0400

    cmd/go: make go test -i -o x.test actually write x.test
    
    Fixes #17078.
    
    Change-Id: I1dfb71f64361b575ec461ed44b0779f2d5cf45fc
    Reviewed-on: https://go-review.googlesource.com/31352
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit e6a901ea3ac4b37be94aaf7b0285ba1840354c4e
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 09:37:54 2016 -0400

    cmd/go: disable SSH connection pooling to avoid git hang
    
    Fixes #13453.
    Fixes #16104.
    
    Change-Id: I4e94f606df786af8143f8649c9afde570f346301
    Reviewed-on: https://go-review.googlesource.com/31353
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit e05d0140483e78c36cd03b3f6173e9f23e975645
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 11:00:00 2016 -0400

    reflect: correct Type.FieldByNameFunc docs
    
    Fixes #16573.
    
    Change-Id: I5a26eaa8b258cb1861190f9690086725532b8a0d
    Reviewed-on: https://go-review.googlesource.com/31354
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e65bce7144dbced232df8842ef6825d7e45f094e
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Sun Oct 16 13:25:53 2016 +0900

    os, syscall: fix incorrect offset calculation in Readlink on windows
    
    Current implementation of syscall.Readlink mistakenly calculates
    the end offset of the PrintName field.
    Also, there are some cases that the PrintName field is empty.
    Instead, the CL uses SubstituteName with correct calculation.
    
    Fixes #15978
    Fixes #16145
    
    Change-Id: If3257137141129ac1c552d003726d5b9c08bb754
    Reviewed-on: https://go-review.googlesource.com/31118
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ebf827ded7d6997747c96ae8f0f4871c15090d49
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 10:46:55 2016 -0400

    testing: wrap long comment line
    
    Requested in CL 31324 review.
    
    Change-Id: Ic81410e07cce07c6f3727bc46d86b6c54c15eca0
    Reviewed-on: https://go-review.googlesource.com/31410
    Reviewed-by: Rob Pike <r@golang.org>

commit 321c312d8246dec6889f5fe334b6193c320baf0e
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 12:34:19 2016 -0400

    os: reject Rename("old", "new") where new is a directory
    
    Unix rejects this when new is a non-empty directory.
    Other systems reject this when new is a directory, empty or not.
    Make Unix reject empty directory too.
    
    Fixes #14527.
    
    Change-Id: Ice24b8065264c91c22cba24aa73e142386c29c87
    Reviewed-on: https://go-review.googlesource.com/31358
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8fbfdad28145bfaad2fa2082336128944d5a3543
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 10 11:18:41 2015 -0800

    spec: require 16 bit minimum exponent in constants rather than 32
    
    A 16bit binary exponent permits a constant range covering roughly the range
    from 7e-9865 to 7e9863 which is more than enough for any practical and
    hypothetical constant arithmetic.
    
    Furthermore, until recently cmd/compile could not handle very large exponents
    correctly anyway; i.e., the chance that any real programs (but for tests that
    explore corner cases) are affected are close to zero.
    
    Finally, restricting the minimum supported range significantly reduces the
    implementation complexity in an area that hardly matters in reality for new
    or alternative spec-compliant implementations that don't or cannot rely on
    pre-existing arbitratry precision arithmetic packages that support a 32bit
    exponent range.
    
    This is technically a language change but for the reasons mentioned above
    this is unlikely to affect any real programs, and certainly not programs
    compiled with the gc or gccgo compilers as they currently support up to
    32bit exponents.
    
    Fixes #13572.
    
    Change-Id: I970f919c57fc82c0175844364cf48ea335f17d39
    Reviewed-on: https://go-review.googlesource.com/17711
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 3f2cb493e5d2a2c2beac9f75a3717a56e294d38a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 18 14:17:05 2016 -0700

    cmd/compile: handle unsafe builtins like universal builtins
    
    Reuse the same mechanisms for handling universal builtins like len to
    handle unsafe.Sizeof, etc. Allows us to drop package unsafe's export
    data, and simplifies some code.
    
    Updates #17508.
    
    Change-Id: I620e0617c24e57e8a2d7cccd0e2de34608779656
    Reviewed-on: https://go-review.googlesource.com/31433
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 7eed848a178cbecae7131434eed1eaab81709a85
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Tue Oct 18 14:00:30 2016 +0530

    math: speed up Gamma(+Inf)
    
    Add special case for Gamma(+∞) which speeds it up:
    
    benchmark            old ns/op     new ns/op     delta
    BenchmarkGamma-4     14.5          7.44          -48.69%
    
    The documentation for math.Gamma already specifies it as a special
    case:
    
            Gamma(+Inf) = +Inf
    
    The original C code that has been used as the reference implementation
    (as mentioned in the comments in gamma.go) also treats Gamma(+∞) as a
    special case:
    
    if( x == INFINITY )
            return(x);
    
    Change-Id: Idac36e19192b440475aec0796faa2d2c7f8abe0b
    Reviewed-on: https://go-review.googlesource.com/31370
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 97b04152bc3d9df1155bc97d9e6095f69b6882c7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Oct 18 21:42:09 2016 +0000

    net/http: update test to check Content-Length 0 Body more reliably
    
    The way to send an explicitly-zero Content-Length is to set a nil Body.
    
    Fix this test to do that, rather than relying on type sniffing.
    
    Updates #17480
    Updates #17071
    
    Change-Id: I6a38e20f17013c88ec4ea69d73c507e4ed886947
    Reviewed-on: https://go-review.googlesource.com/31434
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit 75fef5a0f60bcace88a8d7470df2d85d8eee048f
Author: Quentin Smith <quentin@golang.org>
Date:   Mon Oct 17 18:40:18 2016 -0400

    cmd/go: print more env variables in "go env"
    
    "go env" previously only printed a subset of the documented environment
    variables; now it includes everything, such as GO386 and CGO_*.
    
    This also fixes the CGO_CFLAGS environment variable to always have the
    same default. According to iant@ and confirmed by testing, cgo can now
    understand the default value of CGO_CFLAGS.
    
    Fixes #17191.
    
    Change-Id: Icf75055446dd250b6256ef1139e9ce848f4a9d3b
    Reviewed-on: https://go-review.googlesource.com/31330
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Quentin Smith <quentin@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8b3194ac8fd8436bf4bfd252de58ab81154f334d
Author: Dhananjay Nakrani <dhananjaynakrani@gmail.com>
Date:   Mon Oct 17 14:17:46 2016 -0700

    cmd/compile: fix code duplication in race-instrumentation
    
    instrumentnode() accidentally copies parent's already-instrumented nodes
    into child's Ninit block. This generates repeated code in race-instrumentation.
    This case surfaces only when it duplicates inline-labels, because of
    compile time error. In other cases, it silently generates incorrect
    instrumented code. This change prevents it from doing so.
    
    Fixes #17449.
    
    Change-Id: Icddf2198990442166307e176b7e20aa0cf6c171c
    Reviewed-on: https://go-review.googlesource.com/31317
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2b687a7df854f3c88b266b6cec59a207a45c2353
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Oct 17 16:03:27 2016 -0700

    cmd/compile: stop treating interface methods like actual functions
    
    Interface methods can't have function bodies, so there's no need to
    process their parameter lists as variable declarations. The only
    possible reason would be to check for duplicate parameter names and/or
    invalid types, but we do that anyway, and have regression tests for it
    (test/funcdup.go).
    
    Change-Id: Iedb15335467caa5d872dbab829bf32ab8cf6204d
    Reviewed-on: https://go-review.googlesource.com/31430
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2a441d307868ea5b757fb90eeab07bfc308b94c6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Oct 18 14:01:02 2016 +0000

    net/http/internal: don't block unnecessarily in ChunkedReader
    
    Fixes #17355
    
    Change-Id: I5390979cd0081b61a639466377faa46b4221b74a
    Reviewed-on: https://go-review.googlesource.com/31329
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 430b82009c32e854a209e186f011d47f3241e9b4
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Oct 18 11:47:42 2016 -0400

    cmd/internal/obj/{ppc64,s390x}: mark functions with small stacks NOSPLIT
    
    This change omits the stack check on ppc64 and s390x when the size of
    a stack frame is less than obj.StackSmall. This is an optimization
    x86 already performs.
    
    The effect on s390x isn't huge because we were already omitting the
    stack check when the frame size was 0 (it shaves about 1K from the
    size of bin/go). On ppc64 however this change reduces the size of the
    .text section in bin/go by 33K (1%).
    
    Updates #13379 (for ppc64).
    
    Change-Id: I6af0eb987646bea47fcaf0a812db3496bab0f680
    Reviewed-on: https://go-review.googlesource.com/31357
    Reviewed-by: David Chase <drchase@google.com>

commit c1ab165fa6f0744c87135b53f21576d61c74dec4
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Oct 18 15:27:51 2016 +0200

    net/http: enable TestTransportRemovesDeadIdleConnections on Plan 9
    
    This issue has been fixed in CL 31390.
    
    Fixes #15464.
    
    Change-Id: I35e088f37bf3b544100ff131c72690bcfd788e5b
    Reviewed-on: https://go-review.googlesource.com/31393
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8a255cb85cf4a0b20d0f5dc513d3bb0c56518ec2
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Oct 18 15:28:44 2016 +0200

    net: enable TestCancelRequestWithChannelBeforeDo on Plan 9
    
    This issue has been fixed in CL 31390.
    
    Fixes #11476.
    
    Change-Id: I6658bda2e494d3239d62c49d0bd5d34a36b744d0
    Reviewed-on: https://go-review.googlesource.com/31394
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0c12bdf73b1cf60ceb45f2a302cddba17e86a503
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Oct 18 15:21:46 2016 +0200

    net: always wake up the readers on close on Plan 9
    
    Previously, in acceptPlan9 we set netFD.ctl to the listener's
    /net/tcp/*/listen file instead of the accepted connection's
    /net/tcp/*/ctl file.
    
    In netFD.Read, we write "close" to netFD.ctl to close the
    connection and wake up the readers. However, in the
    case of an accepted connection, we got the error
    "write /net/tcp/*/listen: inappropriate use of fd"
    because the /net/tcp/*/listen doesn't handle the "close" message.
    
    In this case, the connection wasn't closed and the readers
    weren't awake.
    
    We modified the netFD structure so that netFD.ctl represents
    the accepted connection and netFD.listen represents the
    listener.
    
    Change-Id: Ie38c7dbaeaf77fe9ff7da293f09e86d1a01b3e1e
    Reviewed-on: https://go-review.googlesource.com/31390
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c1a1328c5f004c62b8c08faf0d0d2845e0be5d37
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 12 22:58:47 2016 -0400

    encoding/xml: add wildcard support for collecting all attributes
    
    - Like ",any" for elements, add ",any,attr" for attributes to allow
      a mop-up field that gets any otherwise unmapped attributes.
    - Map attributes to fields of type slice by extending the slice,
      just like for elements.
    - Allow storing an attribute into an xml.Attr directly, to provide
      a way to record the name.
    
    Combined, these three independent features allow
    
            AllAttrs []Attr `xml:",any,attr"`
    
    to collect all attributes not otherwise spoken for in a particular struct.
    
    Tests based on CL 16292 by Charles Weill.
    
    Fixes #3633.
    
    Change-Id: I2d75817f17ca8752d7df188080a407836af92611
    Reviewed-on: https://go-review.googlesource.com/30946
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit 0794dce07239fad5845b9c77b50d084f19f7278f
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 21:08:48 2016 -0400

    bufio: read from underlying reader at most once in Read
    
    Fixes #17059.
    
    Change-Id: I5c7ee46604399f7dc3c3c49f964cbb1aa6c0d621
    Reviewed-on: https://go-review.googlesource.com/31320
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3d1ae4b75c77a88ba3b20ada874f1027365a8060
Author: David du Colombier <0intro@gmail.com>
Date:   Mon Oct 17 19:15:37 2016 +0200

    net: close the connection gracefully on Plan 9
    
    Previously, we used to write the "hangup" message to
    the TCP connection control file to be able to close
    a connection, while waking up the readers.
    
    The "hangup" message closes the TCP connection with a
    RST message. This is a problem when closing a connection
    consecutively to a write, because the reader may not have
    time to acknowledge the message before the connection is
    closed, resulting in loss of data.
    
    We use a "close" message, newly implemented in the Plan 9
    kernel to be able to close a TCP connection gracefully with a FIN.
    
    Updates #15464.
    
    Change-Id: I2050cc72fdf7a350bc6c9128bae7d14af11e599c
    Reviewed-on: https://go-review.googlesource.com/31271
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a431bdc712c7a404307f38228271d970d9d2c023
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 23:19:10 2016 -0400

    testing: document that Skip cannot undo Error
    
    Fixes #16502.
    
    Change-Id: Id8e117a724d73cd51844c06d47bbeba61f8dc827
    Reviewed-on: https://go-review.googlesource.com/31324
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>

commit 1188569534fb65fb5e9f0e3eea6b20edc996e983
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 23:30:38 2016 -0400

    testing: fix flag usage messages
    
    Fixes #16404.
    
    Change-Id: Iabaeeef3eff2fff6e5ed2d6bc9ef9c2f6d1cb5e7
    Reviewed-on: https://go-review.googlesource.com/31332
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c999108723fb35cec3667f1fcd60933d2608becc
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 23:51:16 2016 -0400

    time: revise Timer comments for Stop, Reset
    
    The comments added for Go 1.7 are very close.
    Make explicit that they only apply if the timer is
    not known to have expired already.
    
    Fixes #14038.
    
    Change-Id: I6a38be7b2015e1571fc477e18444a8cee38aab29
    Reviewed-on: https://go-review.googlesource.com/31350
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 427674fa0e2fd8c44a2fe7002add00cafcc5894b
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 22:14:32 2016 -0400

    io: clarify Pipe docs
    
    Fixes #14139.
    
    Change-Id: I6d2181720c38582b3d2160e94c7593a6cb4fc60f
    Reviewed-on: https://go-review.googlesource.com/31321
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 736443c13a718f0a9c30327ebbf09f58ccbe6d49
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 17:20:48 2016 -0400

    os/exec: allow simultaneous cmd.Wait and Write of cmd.StdinPipe
    
    cmd.StdinPipe returns an io.WriteCloser.
    It's reasonable to expect the caller not to call Write and Close simultaneously,
    but there is an implicit Close in cmd.Wait that's not obvious.
    We already synchronize the implicit Close in cmd.Wait against
    any explicit Close from the caller. Also synchronize that implicit
    Close against any explicit Write from the caller.
    
    Fixes #9307.
    
    Change-Id: I8561e9369d6e5ac88dfbca1175549f6dfa04b8ac
    Reviewed-on: https://go-review.googlesource.com/31148
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5fbf35dc3fadd29785739fcec061d42157ea7861
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 15:47:18 2016 -0400

    time: be consistent about representation of UTC location in Time struct
    
    In the zero Time, the (not user visible) nil *Location indicates UTC.
    In the result of t.UTC() and other ways to create times in specific
    zones, UTC is indicated by a non-nil *Location, specifically &utcLoc.
    This creates a representation ambiguity exposed by comparison with ==
    or reflect.DeepEqual or the like.
    
    Change time.Time representation to use only nil, never &utcLoc,
    to represent UTC. This eliminates the ambiguity.
    
    Fixes #15716.
    
    Change-Id: I7dcc2c20ce6b073e1daae323d3e49d17d1d52802
    Reviewed-on: https://go-review.googlesource.com/31144
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 79c036238d83c8fceb49aa4bab4d6dade1c321f6
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue Oct 18 19:43:04 2016 +0900

    net: update docs on network interface API
    
    This change documents that the InterfaceAddrs function is less usable on
    multi-homed IP nodes because of the lack of network interface
    identification information.
    
    Also updates documentation on exposed network interface API.
    
    Fixes #14518.
    
    Change-Id: I5e86606f8019ab475eb5d385bd797b052cba395d
    Reviewed-on: https://go-review.googlesource.com/31371
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2a85578b0ecd424e95b29d810b7a414a299fd6a7
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Tue Sep 27 13:27:02 2016 -0700

    database/sql: support returning query database types
    
    Creates a ColumnType structure that can be extended in to future.
    Allow drivers to implement what makes sense for the database.
    
    Fixes #16652
    
    Change-Id: Ieb1fd64eac1460107b1d3474eba5201fa300a4ec
    Reviewed-on: https://go-review.googlesource.com/29961
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2ecaaf18f94cd5ad3ccd46937d36c7a68d3e69bf
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sun Oct 9 21:06:03 2016 +0200

    fmt: always handle special methods if print operand is a reflect.Value
    
    Check for and call the special printing and format methods such as String
    at printing depth 0 when printing the concrete value of a reflect.Value.
    
    Fixes: #16015
    
    Change-Id: I23bd2927255b60924e5558321e98dd4a95e11c4c
    Reviewed-on: https://go-review.googlesource.com/30753
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Martin Möhrmann <martisch@uos.de>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit faf882d1d427e8c8a9a1be00d8ddcab81d1e848e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Oct 14 11:45:59 2016 +0100

    net/http: make Server Handler's Request.Context be done on conn errors
    
    This CL changes how the http1 Server reads from the client.
    
    The goal of this change is to make the Request.Context given to Server
    Handlers become done when the TCP connection dies (has seen any read
    or write error). I didn't finish that for Go 1.7 when Context was
    added to http package.
    
    We can't notice the peer disconnect unless we're blocked in a Read
    call, though, and previously we were only doing read calls as needed,
    when reading the body or the next request. One exception to that was
    the old pre-context CloseNotifier mechanism.
    
    The implementation of CloseNotifier has always been tricky. The past
    few releases have contained the complexity and moved the
    reading-from-TCP-conn logic into the "connReader" type. This CL
    extends connReader to make sure that it's always blocked in a Read
    call, at least once the request body has been fully consumed.
    
    In the process, this deletes all the old CloseNotify code and unifies
    it with the context cancelation code. The two notification mechanisms
    are nearly identical, except the CloseNotify path always notifies on
    the arrival of pipelined HTTP/1 requests. We might want to change that
    in a subsequent commit. I left a TODO for that. For now there's no
    change in behavior except that the context now cancels as it was
    supposed to.
    
    As a bonus that fell out for free, a Handler can now use CloseNotifier
    and Hijack together in the same request now.
    
    Fixes #15224 (make http1 Server always in a Read, like http2)
    Fixes #15927 (cancel context when underlying connection closes)
    Updates #9763 (CloseNotifier + Hijack)
    
    Change-Id: I972cf6ecbab7f1230efe8cc971e89f8e6e56196b
    Reviewed-on: https://go-review.googlesource.com/31173
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 35e5fd0c4de88567aefa354ce6613b7d1ec3a4d9
Author: Adam Langley <agl@golang.org>
Date:   Wed Oct 12 10:53:35 2016 -0700

    crypto/tls: enable ChaCha20-Poly1305 cipher suites by default.
    
    This change enables the ChaCha20-Poly1305 cipher suites by default. This
    changes the default ClientHello and thus requires updating all the
    tests.
    
    Change-Id: I6683a2647caaff4a11f9e932babb6f07912cad94
    Reviewed-on: https://go-review.googlesource.com/30958
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit cff3e7587a516933d842c11b68bdd346ae6fc9be
Author: Adam Langley <agl@golang.org>
Date:   Mon Oct 10 15:27:34 2016 -0700

    crypto/tls: add Config.GetConfigForClient
    
    GetConfigForClient allows the tls.Config to be updated on a per-client
    basis.
    
    Fixes #16066.
    Fixes #15707.
    Fixes #15699.
    
    Change-Id: I2c675a443d557f969441226729f98502b38901ea
    Reviewed-on: https://go-review.googlesource.com/30790
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7e2bf952a905f16a17099970392ea17545cdd193
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 22:40:32 2016 -0400

    net/url: add PathEscape, PathUnescape
    
    Fixes #13737.
    
    Change-Id: Ib655dbf06f44709f687f8a2410c80f31e4075f13
    Reviewed-on: https://go-review.googlesource.com/31322
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 59dae58174ef6dd2ca5720fcce97c565979375ce
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 22:50:44 2016 -0400

    net/url: document and add example for ParseQuery("x")
    
    Fixes #16460.
    
    Change-Id: Ie9d5f725d2d7e8210ab6f7604a5a05fc49f707de
    Reviewed-on: https://go-review.googlesource.com/31331
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ac1108bdcbd316aa3ffc0bf70d50caa35e7b785f
Author: Rob Pike <r@golang.org>
Date:   Sun Oct 16 11:25:37 2016 -0700

    fmt: fix documention for %#v on uints
    
    It's the same as %#x not %x.
    
    Just a documentation change; tests already cover it.
    
    Fixes #17322
    
    Change-Id: Ia9db229f781f9042ac5c0bb824e3d7a26fb74ec5
    Reviewed-on: https://go-review.googlesource.com/31254
    Reviewed-by: Russ Cox <rsc@golang.org>

commit f6cdfc7987d9f3ee7380b3e6f52e433608f342c5
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Mon Oct 17 21:59:10 2016 +0200

    math/big: add benchmarks for big.Float String
    
    In addition to the DecimalConversion benchmark, that exercises the
    String method of the internal decimal type on a range of small shifts,
    add a few benchmarks for the big.Float String method. They can be used
    to obtain more realistic data on the real-world performance of
    big.Float printing.
    
    Change-Id: I7ada324e7603cb1ce7492ccaf3382db0096223ba
    Reviewed-on: https://go-review.googlesource.com/31275
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 57666c3fe842c8a210b55e22514834ec724f945d
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 18 00:34:57 2016 -0400

    test: avoid matching file names in errcheck
    
    Fixes #17030.
    
    Change-Id: Ic7f237ac7553ae0176929056e64b01667ed59066
    Reviewed-on: https://go-review.googlesource.com/31351
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5a0d50f4eb7e35a57364ed5d2f7b79c0b1a80b36
Author: David Symonds <dsymonds@golang.org>
Date:   Tue Oct 18 10:43:47 2016 +1100

    cmd/vet: fix formatting of headings in doc.go.
    
    This will cause godoc to correctly render these docs,
    since go/doc.ToHTML requires no punctuation for headings.
    
    Change-Id: Ic95245147d3585f2ccc59d4424fcab17d2a5617b
    Reviewed-on: https://go-review.googlesource.com/31319
    Reviewed-by: Rob Pike <r@golang.org>

commit 4a5b3ef9b90b2a724fbe9557b4cb15a65d2be87b
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Sun Oct 16 14:31:18 2016 +0900

    encoding/asn1: return error instead of dereferencing nil *big.Int in marshaling
    
    Fixes #17461
    
    Change-Id: I9954f6ae46c7e15560d7460841be8f2bc37233a9
    Reviewed-on: https://go-review.googlesource.com/31121
    Reviewed-by: Adam Langley <agl@golang.org>
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 99b7984de7b02494b4f0d67c9ff42a2367148489
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Oct 17 14:08:36 2016 -0700

    cmd/link: remove some unnecessary comments
    
    The comments about pcln functions are obsolete since those functions
    now live in cmd/internal/obj. The copyright header is redundant with
    the existing one at the top of the file.
    
    Change-Id: I568fd3d259253a0d8eb3b0a157d008df1b5de106
    Reviewed-on: https://go-review.googlesource.com/31315
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c7e855658d73b85f345c9a0ac81de42acad7ae9b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Oct 17 13:40:18 2016 -0400

    cmd/link: trampoline support for external linking on ARM
    
    all.bash passes with -debugtramp=2 (except the unavoidable
    disassembly test as we change instructions). And successfully
    build k8s.io/kubernetes/cmd/hyperkube in both internal linking
    and external linking mode.
    
    Fixes #17028.
    
    Change-Id: Ic8fac6a394488155c5eba9215662db1c1086e24b
    Reviewed-on: https://go-review.googlesource.com/31143
    Reviewed-by: David Chase <drchase@google.com>

commit 007c907b8080e484feed373331210f9287c27120
Author: Adam Langley <agl@golang.org>
Date:   Wed Oct 12 11:20:27 2016 -0700

    crypto/tls: only store a single nonce for AES-GCM.
    
    Although an AEAD, in general, can be used concurrently in both the seal
    and open directions, TLS is easier. Since the transport keys are
    different for different directions in TLS, an AEAD will only ever be
    used in one direction. Thus we don't need separate buffers for seal and
    open because they can never happen concurrently.
    
    Also, fix the nonce size to twelve bytes since the fixed-prefix
    construction for AEADs is superseded and will never be used for anything
    else now.
    
    Change-Id: Ibbf6c6b1da0e639f4ee0e3604410945dc7dcbb46
    Reviewed-on: https://go-review.googlesource.com/30959
    Run-TryBot: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d1bfa3c68d11c861f3e3629e525b2515aa1ef68f
Author: Adam Langley <agl@golang.org>
Date:   Mon Oct 17 14:26:57 2016 -0700

    Revert "crypto/tls: add CloseWrite method to Conn"
    
    This reverts commit c6185aa63217c84a1a73c578c155e7d4dec6cec8. That
    commit seems to be causing flaky failures on the builders. See
    discussion on the original thread: https://golang.org/cl/25159.
    
    Change-Id: I26e72d962d4efdcee28a0bc61a53f246b046df77
    Reviewed-on: https://go-review.googlesource.com/31316
    Run-TryBot: Adam Langley <agl@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 67d8226b4862f0f8deb4dc6fa8617017ecb0f32b
Author: Adam Langley <agl@golang.org>
Date:   Wed Oct 12 10:46:43 2016 -0700

    crypto/tls: support ChaCha20-Poly1305.
    
    This change adds support for the ChaCha20-Poly1305 AEAD to crypto/tls,
    as specified in https://tools.ietf.org/html/rfc7905.
    
    Fixes #15499.
    
    Change-Id: Iaa689be90e03f208c40b574eca399e56f3c7ecf1
    Reviewed-on: https://go-review.googlesource.com/30957
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e7edc7e27e13c789b56768b556dfcde767920f10
Author: Adam Langley <agl@golang.org>
Date:   Mon Oct 17 13:35:27 2016 -0700

    vendor: update golang.org/x/crypto/chacha20poly1305
    
    This change updates the vendored chacha20poly1305 package to match
    revision 14f9af67c679edd414f72f13d67c917447113df2 of x/crypto.
    
    Change-Id: I05a4ba86578b0f0cdb1ed7dd50fee3b38bb48cf5
    Reviewed-on: https://go-review.googlesource.com/31312
    Run-TryBot: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bb4c40b21a30c02c069270327c5b34acab19b3f2
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Oct 17 13:34:40 2016 -0700

    doc: document go1.7.2
    
    Change-Id: I34b3650ee9512879ff7528336813a7850c46ea90
    Reviewed-on: https://go-review.googlesource.com/31311
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7dd9c385f6896c7dcb5d76353e52e36c81af2838
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Oct 17 20:23:51 2016 +0000

    Revert "cmd/compile: inline convI2E"
    
    This reverts commit 395d36a67df8d25b35617ec8709f0164ae2f655e.
    
    Appears to be responsible for builder failures.
    
    Change-Id: Ic6c6307f662767e529060b88704a9f074785d99e
    Reviewed-on: https://go-review.googlesource.com/31310
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 9ee21f90d2594412dd60dd821831056db708fa53
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 10 16:18:43 2016 -0400

    math/big: add (*Int).Sqrt
    
    This is needed for some of the more complex primality tests
    (to filter out exact squares), and while the code is simple the
    boundary conditions are not obvious, so it seems worth having
    in the library.
    
    Change-Id: Ica994a6b6c1e412a6f6d9c3cf823f9b653c6bcbd
    Reviewed-on: https://go-review.googlesource.com/30706
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 2190f771d876180fe3fe51d785f0dbc32a5373d0
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Mon Oct 17 14:35:55 2016 -0500

    bytes: fix typo in ppc64le asm for Compare
    
    Correcting a line in asm_ppc64x.s in the cmpbodyLE function
    that originally was R14 but accidentally changed to R4.
    
    Fixes #17488
    
    Change-Id: Id4ca6fb2e0cd81251557a0627e17b5e734c39e01
    Reviewed-on: https://go-review.googlesource.com/31266
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    Run-TryBot: Michael Munday <munday@ca.ibm.com>

commit 1cfb5c3fd5578a3665231a302ef7f03abec78d1d
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Sep 14 10:42:14 2016 -0400

    cmd/compile: merge loads into operations on s390x
    
    Adds the new canMergeLoad function which can be used by rules to
    decide whether a load can be merged into an operation. The function
    ensures that the merge will not reorder the load relative to memory
    operations (for example, stores) in such a way that the block can no
    longer be scheduled.
    
    This new function enables transformations such as:
    
    MOVD 0(R1), R2
    ADD  R2, R3
    
    to:
    
    ADD  0(R1), R3
    
    The two-operand form of the following instructions can now read a
    single memory operand:
    
     - ADD
     - ADDC
     - ADDW
     - MULLD
     - MULLW
     - SUB
     - SUBC
     - SUBE
     - SUBW
     - AND
     - ANDW
     - OR
     - ORW
     - XOR
     - XORW
    
    Improves SHA3 performance by 6-8%.
    
    Updates #15054.
    
    Change-Id: Ibcb9122126cd1a26f2c01c0dfdbb42fe5e7b5b94
    Reviewed-on: https://go-review.googlesource.com/29272
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 92221fe8bc73ec6d487bd479f9739fdddb6fcada
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Oct 14 17:09:54 2016 -0700

    math/big: slightly faster float->decimal conversion
    
    Inspired by Alberto Donizetti's observations in
    https://go-review.googlesource.com/#/c/30099/.
    
    name                 old time/op  new time/op  delta
    DecimalConversion-8   138µs ± 1%   136µs ± 2%  -1.85%  (p=0.000 n=10+10)
    
    10 runs each, measured on a Mac Mini, 2.3 GHz Intel Core i7.
    
    Performance improvements varied between -1.25% to -4.4%; -1.85% is
    about in the middle of the observed improvement. The generated code
    is slightly shorter in the inner loops of the conversion code.
    
    Change-Id: I10fb3b2843da527691c39ad5e5e5bd37ed63e2fa
    Reviewed-on: https://go-review.googlesource.com/31250
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 984753b665875cd62f6144a42f6df58cb5f159a8
Author: Austin Clements <austin@google.com>
Date:   Mon Oct 10 12:18:00 2016 -0400

    runtime: fix GC assist retry path
    
    GC assists retry if preempted or if they fail to park. However, on the
    retry path they currently use stale statistics. In particular, the
    retry can use "debtBytes", but debtBytes isn't updated when the debt
    changes (since other than retries it is only used once). Also, though
    less of a problem, the if the assist ratio has changed while the
    assist was blocked, the retry will still use the old assist ratio.
    
    Fix all of this by simply making the retry jump back to where we
    compute these statistics, rather than just after.
    
    Change-Id: I2ed8b4f0fc9f008ff060aa926f4334b662ac7d3f
    Reviewed-on: https://go-review.googlesource.com/30701
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 81c431a53780d90ebe8a81205db9b3bee7824ffd
Author: Austin Clements <austin@google.com>
Date:   Thu Oct 6 15:12:12 2016 -0400

    runtime: abstract out assist queue management
    
    This puts all of the assist queue-related code together and makes it
    easier to modify how the assist queue works.
    
    Change-Id: Id54e06702bdd5a5dd3fef2ce2c14cd7ca215303c
    Reviewed-on: https://go-review.googlesource.com/30700
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 395d36a67df8d25b35617ec8709f0164ae2f655e
Author: Keith Randall <khr@golang.org>
Date:   Sun Oct 16 23:40:12 2016 -0700

    cmd/compile: inline convI2E
    
    It's pretty simple.  For:
      e = (interface{})(i)
    Do:
      tmp = i.itab
      if tmp != nil {
        tmp = tmp.typ_ // load type from itab
      }
      e = eface{tmp, i.data}
    
    It is smaller and faster than calling the runtime.
    
    Change-Id: I0ad27f62f4ec0b6cd53bc8530e4da0eae3e67a6c
    Reviewed-on: https://go-review.googlesource.com/31260
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d5bd797ee55d36b07b388d0b8ca2393dc186bea1
Author: Austin Clements <austin@google.com>
Date:   Sun Oct 16 18:23:39 2016 -0400

    runtime: fix getArgInfo for deferred reflection calls
    
    getArgInfo for reflect.makeFuncStub and reflect.methodValueCall is
    necessarily special. These have dynamically determined argument maps
    that are stored in their context (that is, their *funcval). These
    functions are written to store this context at 0(SP) when called, and
    getArgInfo retrieves it from there.
    
    This technique works if getArgInfo is passed an active call frame for
    one of these functions. However, getArgInfo is also used in
    tracebackdefers, where the "call" is not a true call with an active
    stack frame, but a deferred call. In this situation, getArgInfo
    currently crashes because tracebackdefers passes a frame with sp set
    to 0. However, the entire approach used by getArgInfo is flawed in
    this situation because the wrapper has not actually executed, and
    hence hasn't saved this metadata to any stack frame.
    
    In the defer case, we know the *funcval from the _defer itself, so we
    can fix this by teaching getArgInfo to use the *funcval context
    directly when its available, and otherwise get it from the active call
    frame.
    
    While we're here, this commit simplifies getArgInfo a bit by making it
    play more nicely with the type system. Rather than decoding the
    *reflect.methodValue that is the wrapper's context as a *[2]uintptr,
    just write out a copy of the reflect.methodValue type in the runtime.
    
    Fixes #16331. Fixes #17471.
    
    Change-Id: I81db4d985179b4a81c68c490cceeccbfc675456a
    Reviewed-on: https://go-review.googlesource.com/31138
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 687d9d5d78f8a2d09b2052e73be0c83740e17fda
Author: Austin Clements <austin@google.com>
Date:   Thu Oct 13 10:44:57 2016 -0400

    runtime: print a message on bad morestack
    
    If morestack runs on the g0 or gsignal stack, it currently performs
    some abort operation that typically produces a signal (e.g., it does
    an INT $3 on x86). This is useful if you're running in a debugger, but
    if you're not, the runtime tries to trap this signal, which is likely
    to send the program into a deeper spiral of collapse and lead to very
    confusing diagnostic output.
    
    Help out people trying to debug without a debugger by making morestack
    print an informative message before blowing up.
    
    Change-Id: I2814c64509b137bfe20a00091d8551d18c2c4749
    Reviewed-on: https://go-review.googlesource.com/31133
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 0ba3c607dfcc90072191375d57c4059be1ae96c7
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 17 13:29:31 2016 -0400

    cmd/dist, go/build: make CGO_ENABLED during make.bash sticky
    
    Per discussion on #12808, it's a bit odd that if you do
    
            CGO_ENABLED=0 ./make.bash
    
    then you get a toolchain that still tries to use cgo.
    So make the CGO_ENABLED setting propagate into
    the resulting toolchain as the default setting for that
    environment variable, like we do with other variables
    like CC and GOROOT.
    
    No reasonable way to test automatically, but I did
    test by hand that after the above command, 'go env'
    shows CGO_ENABLED=0; before it showed CGO_ENABLED=1.
    
    Fixes #12808.
    
    Change-Id: I26a2fa6cc00e73bde8af7469270b27293392ed71
    Reviewed-on: https://go-review.googlesource.com/31141
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1e28dce80ad2ec195d55269266c5cca7ebd845a5
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Thu Oct 13 12:59:07 2016 -0500

    bytes: improve performance for bytes.Compare on ppc64x
    
    This improves the performance for byte.Compare by rewriting
    the cmpbody function in runtime/asm_ppc64x.s.  The previous code
    had a simple loop which loaded a pair of bytes and compared them,
    which is inefficient for long buffers.  The updated function checks
    for 8 or 32 byte chunks and then loads and compares double words where
    possible.
    
    Because the byte.Compare result indicates greater or less than,
    the doubleword loads must take endianness into account, using a
    byte reversed load in the little endian case.
    
    Fixes #17433
    
    benchmark                                   old ns/op     new ns/op     delta
    BenchmarkBytesCompare/8-16                  13.6          7.16          -47.35%
    BenchmarkBytesCompare/16-16                 25.7          7.83          -69.53%
    BenchmarkBytesCompare/32-16                 38.1          7.78          -79.58%
    BenchmarkBytesCompare/64-16                 63.0          10.6          -83.17%
    BenchmarkBytesCompare/128-16                112           13.0          -88.39%
    BenchmarkBytesCompare/256-16                211           28.1          -86.68%
    BenchmarkBytesCompare/512-16                410           38.6          -90.59%
    BenchmarkBytesCompare/1024-16               807           60.2          -92.54%
    BenchmarkBytesCompare/2048-16               1601          103           -93.57%
    
    Change-Id: I121acc74fcd27c430797647b8d682eb0607c63eb
    Reviewed-on: https://go-review.googlesource.com/30949
    Reviewed-by: David Chase <drchase@google.com>

commit 7c46f0349844f950cc811727ded2393cff7e0369
Author: Quentin Smith <quentin@golang.org>
Date:   Fri Oct 14 15:13:30 2016 -0400

    strconv: strip \r in raw strings passed to Unquote
    
    To match the language spec, strconv.Unquote needs to strip carriage
    returns from the raw string.
    
    Also fixes TestUnquote to not be a noop.
    
    Fixes #15997
    
    Change-Id: I2456f50f2ad3830f37e545f4f6774ced9fe609d7
    Reviewed-on: https://go-review.googlesource.com/31210
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 6f4a6faf86238285fd02d2d04cbf3aeb51eb9d37
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Oct 17 14:09:40 2016 -0400

    cmd/dist: disable math/big assembly when using the bootstrap compiler
    
    The assembly in math/big may contain instructions that the bootstrap
    compiler does not support. Disable it using the math_big_pure_go
    build tag.
    
    Fixes #17484.
    
    Change-Id: I766cab6a888721ab4ed76ebdbfc87ad4e919ec41
    Reviewed-on: https://go-review.googlesource.com/31142
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cbf28ff87c0aab519cd87a27c168d433f2404764
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sat Oct 1 14:01:09 2016 +0200

    strconv: make FormatFloat slowpath a little faster
    
    The relevant benchmark (on an Intel i7-4510U machine):
    
    name                      old time/op  new time/op  delta
    FormatFloat/Slowpath64-4  68.6µs ± 0%  44.1µs ± 2%  -35.71%  (p=0.000 n=13+15)
    
    Change-Id: I67eb0e81ce74ed57752d0280059f91419f09e93b
    Reviewed-on: https://go-review.googlesource.com/30099
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 77b6a08e0d3061cbf505394f1d0c5bc9d16fa6cb
Author: Adam Langley <agl@golang.org>
Date:   Wed Oct 12 10:18:37 2016 -0700

    vendor: add golang.org/x/crypto/{chacha20poly1305,poly1305}
    
    This change imports the chacha20poly1305 and poly1305 packages from
    x/crypto at 5f4e837b98443e9e7a65072235205993af565d85. These packages
    will be used to support the ChaCha20-Poly1305 AEAD in crypto/tls.
    
    Change-Id: I1a38d671ef9aeff3bc41e3924655883d465a5617
    Reviewed-on: https://go-review.googlesource.com/30956
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9a97c3bfe41d1ed768ea3bd3d8f0f52b8a51bb62
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 13 13:45:31 2016 -0400

    cmd/go: accept plain file for .vcs (instead of directory)
    
    Sometimes .git is a plain file; maybe others will follow.
    This CL matches CL 21430, made in x/tools/go/vcs.
    
    The change in the Swift test case makes the test case
    pass by changing the test to match current behavior,
    which I assume is better than the reverse.
    (The test only runs locally and without -short, so the
    builders are not seeing this particular failure.)
    
    For #10322.
    
    Change-Id: Iccd08819a01c5609a2880b9d8a99af936e20faff
    Reviewed-on: https://go-review.googlesource.com/30948
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d26b0661c23dedc5c3e2c2ed2ca6250e440bf010
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Tue Oct 11 09:26:40 2016 -0500

    cmd/link: add trampolines for too far calls in ppc64x
    
    This change adds support for trampolines on ppc64x when using
    internal linking, in the case where the offset to the branch
    target is larger than what fits in the field provided by the
    branch instruction.
    
    Fixes #16665
    
    Change-Id: Icfee72910f38c94588d2adce517b64dee6176145
    Reviewed-on: https://go-review.googlesource.com/30850
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 0ce1d79a6a771f7449ec493b993ed2a720917870
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Oct 11 08:34:58 2016 -0700

    database/sql: accept nil pointers to Valuers implemented on value receivers
    
    The driver.Valuer interface lets types map their Go representation to
    a suitable database/sql/driver.Value.
    
    If a user defines the Value method with a value receiver, such as:
    
        type MyStr string
    
        func (s MyStr) Value() (driver.Value, error) {
            return strings.ToUpper(string(s)), nil
        }
    
    Then they can't use (*MyStr)(nil) as an argument to an SQL call via
    database/sql, because *MyStr also implements driver.Value, but via a
    compiler-generated wrapper which checks whether the pointer is nil and
    panics if so.
    
    We now accept (*MyStr)(nil) and map it to "nil" (an SQL "NULL")
    if the Valuer method is implemented on MyStr instead of *MyStr.
    
    If a user implements the driver.Value interface with a pointer
    receiver, they retain full control of what nil means:
    
        type MyStr string
    
        func (s *MyStr) Value() (driver.Value, error) {
            if s == nil {
                return "missing MyStr", nil
            }
            return strings.ToUpper(string(*s)), nil
        }
    
    Adds tests for both cases.
    
    Fixes #8415
    
    Change-Id: I897d609d80d46e2354d2669a8a3e090688eee3ad
    Reviewed-on: https://go-review.googlesource.com/31259
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Daniel Theophanes <kardianos@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 237d7e34bc7579ec499a029191c33106dc5476a1
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 13 15:27:04 2016 -0400

    cmd/dist: use debug/pe directly for cmd/link
    
    Delete vendored copy.
    
    Change-Id: I06e9d3b709553a1a8d06275e99bd8f617aac5788
    Reviewed-on: https://go-review.googlesource.com/31011
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 15040c11b9b8f2e33b2833feb99a48ea9a2dba6c
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 13 15:13:41 2016 -0400

    cmd/dist: copy needed packages from standard library during bootstrap
    
    This allows use of newer math/big (and later debug/pe)
    without maintaining a vendored copy somewhere in cmd.
    
    Use for math/big, deleting cmd/compile/internal/big.
    
    Change-Id: I2bffa7a9ef115015be29fafdb02acc3e7a665d11
    Reviewed-on: https://go-review.googlesource.com/31010
    Reviewed-by: Minux Ma <minux@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f444b48fe419c2d19b7b9a89faad30f0e8b0e474
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 12 16:54:02 2016 -0400

    encoding/json: fix decoding of null into Unmarshaler, TextUnmarshaler
    
    1. Define behavior for Unmarshal of JSON null into Unmarshaler and
    TextUnmarshaler. Specifically, an Unmarshaler will be given the
    literal null and can decide what to do (because otherwise
    json.RawMessage is impossible to implement), and a TextUnmarshaler
    will be skipped over (because there is no text to unmarshal), like
    most other inappropriate types. Document this in Unmarshal, with a
    reminder in UnmarshalJSON about handling null.
    
    2. Test all this.
    
    3. Fix the TextUnmarshaler case, which was returning an unmarshalling
    error, to match the definition.
    
    4. Fix the error that had been used for the TextUnmarshaler, since it
    was claiming that there was a JSON string when in fact the problem was
    NOT having a string.
    
    5. Adjust time.Time and big.Int's UnmarshalJSON to ignore null, as is
    conventional.
    
    Fixes #9037.
    
    Change-Id: If78350414eb8dda712867dc8f4ca35a9db041b0c
    Reviewed-on: https://go-review.googlesource.com/30944
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c6185aa63217c84a1a73c578c155e7d4dec6cec8
Author: Ben Burkert <ben@benburkert.com>
Date:   Sun Jul 24 15:13:56 2016 -0700

    crypto/tls: add CloseWrite method to Conn
    
    The CloseWrite method sends a close_notify alert record to the other
    side of the connection. This record indicates that the sender has
    finished sending on the connection. Unlike the Close method, the sender
    may still read from the connection until it recieves a close_notify
    record (or the underlying connection is closed). This is analogous to a
    TCP half-close.
    
    Updates #8579
    
    Change-Id: I9c6bc193efcb25cc187f7735ee07170afa7fdde3
    Reviewed-on: https://go-review.googlesource.com/25159
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b97b7537bc5cf77b8781bd699b17747ca2f02e1a
Author: Victor Vrantchan <vrancean+github@gmail.com>
Date:   Sun Oct 16 09:55:33 2016 -0400

    encoding/pem: add Decode example
    
    For #16360.
    
    Change-Id: I99d1e5ab1f814f65b3066a498158a442f1bd477f
    Reviewed-on: https://go-review.googlesource.com/31137
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d2951740303587fc0c5d14cb5461e39b099e6695
Author: Martin Möhrmann <martisch@uos.de>
Date:   Fri Sep 2 17:04:41 2016 +0200

    runtime: speed up non-ASCII rune decoding
    
    Copies utf8 constants and EncodeRune implementation from unicode/utf8.
    
    Adds a new decoderune implementation that is used by the compiler
    in code generated for ranging over strings. It does not handle
    ASCII runes since these are handled directly before calls to decoderune.
    
    The DecodeRuneInString implementation from unicode/utf8 is not used
    since it uses a lookup table that would increase the use of cpu caches.
    
    Adds more tests that check decoding of valid and invalid utf8 sequences.
    
    name                              old time/op  new time/op  delta
    RuneIterate/range2/ASCII-4        7.45ns ± 2%  7.45ns ± 1%     ~     (p=0.634 n=16+16)
    RuneIterate/range2/Japanese-4     53.5ns ± 1%  49.2ns ± 2%   -8.03%  (p=0.000 n=20+20)
    RuneIterate/range2/MixedLength-4  46.3ns ± 1%  41.0ns ± 2%  -11.57%  (p=0.000 n=20+20)
    
    new:
    "".decoderune t=1 size=423 args=0x28 locals=0x0
    old:
    "".charntorune t=1 size=666 args=0x28 locals=0x0
    
    Change-Id: I1df1fdb385bb9ea5e5e71b8818ea2bf5ce62de52
    Reviewed-on: https://go-review.googlesource.com/28490
    Run-TryBot: Martin Möhrmann <martisch@uos.de>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fe4307f0607dff7742d047b04df06e721aea7906
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Oct 16 05:00:27 2016 -0700

    net/http: support multiple identical Content-Length headers
    
    Referencing RFC 7230 Section 3.3.2, this CL
    deduplicates multiple identical Content-Length headers
    of a message or rejects the message as invalid if the
    Content-Length values differ.
    
    Fixes #16490
    
    Change-Id: Ia6b0f58ec7d35710b11a36113d2bd9128f693f64
    Reviewed-on: https://go-review.googlesource.com/31252
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 516001f50ff3de9f7de2a0525bb771b509504e93
Author: David du Colombier <0intro@gmail.com>
Date:   Mon Oct 17 10:54:44 2016 +0200

    net: skip TestReadTimeoutUnblocksRead on Plan 9
    
    Deadlines aren't implemented on Plan 9 yet.
    
    Updates #17477.
    
    Change-Id: I44ffdbef97276dfec56547e5189672b7da24bfc1
    Reviewed-on: https://go-review.googlesource.com/31188
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: David du Colombier <0intro@gmail.com>

commit c1cd64d0ac2001ec8fe6f253e4e95561a444d533
Author: Klaus Post <klauspost@gmail.com>
Date:   Sat Oct 15 13:06:22 2016 +0200

    compress/flate: use correct table for size estimation
    
    The incorrect table was used for estimating output size.
    This can give suboptimal selection of entropy encoder in rare cases.
    
    Change-Id: I8b358200f2d1f9a3f9b79a44269d7be704e1d2d9
    Reviewed-on: https://go-review.googlesource.com/31172
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 05471e9ee64a300bd2dcc4582ee1043c055893bb
Author: Yasuhiro Matsumoto <mattn.jp@gmail.com>
Date:   Fri Oct 7 17:10:53 2016 +0900

    crypto/x509: implement SystemCertPool on Windows
    
    Fixes #16736
    
    Change-Id: I335d201e3f6738d838de3881087cb640fc7670e8
    Reviewed-on: https://go-review.googlesource.com/30578
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 707a83341b8c7973f4e0fce731fa279c618f233b
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Mon Oct 3 09:49:25 2016 -0700

    database/sql: add option to use named parameter in query arguments
    
    Modify the new Context methods to take a name-value driver struct.
    This will require more modifications to drivers to use, but will
    reduce the overall number of structures that need to be maintained
    over time.
    
    Fixes #12381
    
    Change-Id: I30747533ce418a1be5991a0c8767a26e8451adbd
    Reviewed-on: https://go-review.googlesource.com/30166
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 99df54f19696e26bea8d6a052d8d91ddb1e4ea65
Author: Martin Möhrmann <martisch@uos.de>
Date:   Tue Sep 6 13:42:49 2016 +0200

    bytes: encode size of rune read by ReadRune into lastRead to speed up UnreadRune
    
    In ReadRune store the size of the rune that was read into lastRead
    to avoid the need to call DecodeRuneLast in UnreadRune.
    
    fmt:
    name        old time/op  new time/op  delta
    ScanInts-4   481µs ± 4%   458µs ± 3%  -4.64%  (p=0.000 n=20+20)
    
    Change-Id: I500848e663a975f426402a4b3d27a541e5cac06c
    Reviewed-on: https://go-review.googlesource.com/28817
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    Run-TryBot: Martin Möhrmann <martisch@uos.de>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1e775fe4a30c78cf6e8e995efe5a0f469b8d9a4d
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Oct 16 10:27:29 2016 +0100

    net/http: add more docs on ErrHijacked
    
    Updates #16456
    
    Change-Id: Ifea651ea3ece2267a6f0c1638181d6ddd9248a9f
    Reviewed-on: https://go-review.googlesource.com/31181
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 7a7ea01c65e8366af277b956dc8ccf0601727172
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Oct 15 16:56:51 2016 +0100

    syscall, net: make deadline changes affect blocked read/write calls on nacl
    
    Flesh out nacl's fake network system to match how all the other
    platforms work: all other systems' SetReadDeadline and
    SetWriteDeadline affect currently-blocked read & write calls.
    This was documented in golang.org/cl/30164 because it was the status
    quo and existing packages relied on it. (notably the net/http package)
    
    And add a test.
    
    Change-Id: I074a1054dcabcedc97b173dad5e827f8babf7cfc
    Reviewed-on: https://go-review.googlesource.com/31178
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit cd2c9df7612795cad5b56cabe5ec29c7771db5fe
Author: Caleb Spare <cespare@gmail.com>
Date:   Fri Oct 14 00:59:19 2016 -0700

    html/template: fix Clone so that t.Lookup(t.Name()) yields t
    
    Template.escape makes the assumption that t.Lookup(t.Name()) is t
    (escapeTemplate looks up the associated template by name and sets
    escapeErr appropriately).
    
    This assumption did not hold for a Cloned template, because the template
    associated with t.Name() was a second copy of the original.
    
    Add a test for the assumption that t.Lookup(t.Name()) == t.
    
    One effect of this broken assumption was #16101: parallel Executes
    racily accessed the template namespace because each Execute call saw
    t.escapeErr == nil and re-escaped the template concurrently with read
    accesses occurring outside the namespace mutex.
    
    Add a test for this race.
    
    Related to #12996 and CL 16104.
    
    Fixes #16101
    
    Change-Id: I59831d0847abbabb4ef9135f2912c6ce982f9837
    Reviewed-on: https://go-review.googlesource.com/31092
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit d8cbc2c918f68e8ca5992e68fed052a0e52a8e67
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Fri Oct 14 17:03:01 2016 +1100

    misc/cgo/testcarchive: do not use same executable name in TestInstall
    
    Fixes #17439
    
    Change-Id: I7caa28519f38692f9ca306f0789cbb975fa1d7c4
    Reviewed-on: https://go-review.googlesource.com/31112
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 42f5ee4cbf15ff02803863d1002259565da0f071
Author: Rob Pike <r@golang.org>
Date:   Sun Oct 16 11:28:45 2016 -0700

    testing: mention in docs for Logf that a final newline is added if needed
    
    Fixes #16423
    
    Change-Id: I9635db295be4d356d427adadd309084e16c4582f
    Reviewed-on: https://go-review.googlesource.com/31255
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ca28620db9392d5aa1002256e41bd066f0cc6377
Author: Rob Pike <r@golang.org>
Date:   Sun Oct 16 11:13:37 2016 -0700

    cmd/go: use normal code 2 for 'no such tool'
    
    Exit code 3 is unprecedented and inconsistent with other failures here,
    such as having no tool directory.
    
    Fixes #17145
    
    Change-Id: Ie7ed56494d4511a600214666ce3a726d63a8fd8e
    Reviewed-on: https://go-review.googlesource.com/31253
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1bcfca0563ce21dce69b2652859d6f0f800c39af
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Oct 15 15:52:57 2016 +0100

    net: enable a test on nacl
    
    No need to skip it. It passes.
    
    Maybe it was fixed at some point.
    
    Change-Id: I9848924aefda44f9b3a574a8705fa549d657f28d
    Reviewed-on: https://go-review.googlesource.com/31177
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit 4d898776ff973056aa906d90be0f23af772631c6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Oct 16 11:27:43 2016 +0100

    doc: update go1.8.txt
    
    Change-Id: Ibae0be046c6a6596d3a98b094ec5f089bb68be7a
    Reviewed-on: https://go-review.googlesource.com/31182
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ad50408fe7d9edfc7a1d9791e7391df132bc58b2
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Tue Sep 20 19:43:34 2016 +0900

    path/filepath: simplify TestToNorm
    
    Change-Id: I8a176ed9c7f59ebdfd39c1e2b88905f977179982
    Reviewed-on: https://go-review.googlesource.com/31119
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0cc400f0e7e73e1ee870d803b8faf1904c0c9ec3
Author: Alex Carol <alex.carol.c@gmail.com>
Date:   Sun Oct 16 00:26:39 2016 +0200

    net/rpc: add missing import to rpc server documentation
    
    Change-Id: Idca6115181960eed7a955027ee77a02decb4e7f2
    Reviewed-on: https://go-review.googlesource.com/31179
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ad5fd2872f6cce199aa4f543801d8793ef0a901f
Author: Austin Clements <austin@google.com>
Date:   Fri Oct 14 17:35:18 2016 -0400

    test: simplify fixedbugs/issue15747.go
    
    The error check patterns in this test are more complex than necessary
    because f2 gets inlined into f1. This behavior isn't important to the
    test, so disable inlining of f2 and simplify the error check patterns.
    
    Change-Id: Ia8aee92a52f9217ad71b89b2931494047e8d2185
    Reviewed-on: https://go-review.googlesource.com/31132
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 9897e4081192abd81fd404774dd492975d7ccedd
Author: Austin Clements <austin@google.com>
Date:   Mon Oct 10 17:14:14 2016 -0400

    runtime: use more go:nowritebarrierrec in proc.go
    
    Currently we use go:nowritebarrier in many places in proc.go.
    go:notinheap and go:yeswritebarrierrec now let us use
    go:nowritebarrierrec (the recursive form of the go:nowritebarrier
    pragma) more liberally. Do so in proc.go
    
    Change-Id: Ia7fcbc12ce6c51cb24730bf835fb7634ad53462f
    Reviewed-on: https://go-review.googlesource.com/30942
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 1bc6be6423e48318451a0faeaae840772137b001
Author: Austin Clements <austin@google.com>
Date:   Tue Oct 11 22:58:21 2016 -0400

    runtime: mark several types go:notinheap
    
    This covers basically all sysAlloc'd, persistentalloc'd, and
    fixalloc'd types.
    
    Change-Id: I0487c887c2a0ade5e33d4c4c12d837e97468e66b
    Reviewed-on: https://go-review.googlesource.com/30941
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 991a85c88944e9cb92c4860c173f49d549a92845
Author: Austin Clements <austin@google.com>
Date:   Tue Oct 11 11:47:14 2016 -0400

    runtime: make mSpanList more go:notinheap-friendly
    
    Currently mspan links to its previous mspan using a **mspan field that
    points to the previous span's next field. This simplifies some of the
    list manipulation code, but is going to make it very hard to convince
    the compiler that mspan list manipulations don't need write barriers.
    
    Fix this by using a more traditional ("boring") linked list that uses
    a simple *mspan pointer to the previous mspan. This complicates some
    of the list manipulation slightly, but it will let us eliminate all
    write barriers from the mspan list manipulation code by marking mspan
    go:notinheap.
    
    Change-Id: I0d0b212db5f20002435d2a0ed2efc8aa0364b905
    Reviewed-on: https://go-review.googlesource.com/30940
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 77527a316b33d6f4c072c0774a1478bb53f42d35
Author: Austin Clements <austin@google.com>
Date:   Tue Oct 11 22:53:27 2016 -0400

    cmd/compile: add go:notinheap type pragma
    
    This adds a //go:notinheap pragma for declarations of types that must
    not be heap allocated. We ensure these rules by disallowing new(T),
    make([]T), append([]T), or implicit allocation of T, by disallowing
    conversions to notinheap types, and by propagating notinheap to any
    struct or array that contains notinheap elements.
    
    The utility of this pragma is that we can eliminate write barriers for
    writes to pointers to go:notinheap types, since the write barrier is
    guaranteed to be a no-op. This will let us mark several scheduler and
    memory allocator structures as go:notinheap, which will let us
    disallow write barriers in the scheduler and memory allocator much
    more thoroughly and also eliminate some problematic hybrid write
    barriers.
    
    This also makes go:nowritebarrierrec and go:yeswritebarrierrec much
    more powerful. Currently we use go:nowritebarrier all over the place,
    but it's almost never what you actually want: when write barriers are
    illegal, they're typically illegal for a whole dynamic scope. Partly
    this is because go:nowritebarrier has been around longer, but it's
    also because go:nowritebarrierrec couldn't be used in situations that
    had no-op write barriers or where some nested scope did allow write
    barriers. go:notinheap eliminates many no-op write barriers and
    go:yeswritebarrierrec makes it possible to opt back in to write
    barriers, so these two changes will let us use go:nowritebarrierrec
    far more liberally.
    
    This updates #13386, which is about controlling pointers from non-GC'd
    memory to GC'd memory. That would require some additional pragma (or
    pragmas), but could build on this pragma.
    
    Change-Id: I6314f8f4181535dd166887c9ec239977b54940bd
    Reviewed-on: https://go-review.googlesource.com/30939
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a9e6cebde21875379ccc05d680b3c3a78adbf089
Author: Austin Clements <austin@google.com>
Date:   Mon Oct 10 16:46:28 2016 -0400

    cmd/compile, runtime: add go:yeswritebarrierrec pragma
    
    This pragma cancels the effect of go:nowritebarrierrec. This is useful
    in the scheduler because there are places where we enter a function
    without a valid P (and hence cannot have write barriers), but then
    obtain a P. This allows us to annotate the function with
    go:nowritebarrierrec and split out the part after we've obtained a P
    into a go:yeswritebarrierrec function.
    
    Change-Id: Ic8ce4b6d3c074a1ecd8280ad90eaf39f0ffbcc2a
    Reviewed-on: https://go-review.googlesource.com/30938
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 6347367be36df608cce84beb097378f8654dd208
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Wed Sep 7 16:35:32 2016 +0300

    strings: use Index in Count
    
    This simplifies code and provides performance iprovments:
    Similar to https://go-review.googlesource.com/#/c/28577
    
    CountHard1-48               1.74ms ±14%  0.17ms ±14%  -90.16%  (p=0.000 n=19+19)
    CountHard2-48               1.78ms ±15%  0.25ms ±13%  -86.10%  (p=0.000 n=19+20)
    CountHard3-48               1.78ms ±12%  0.80ms ±11%  -55.19%  (p=0.000 n=17+20)
    CountTorture-48             13.5µs ±14%  13.6µs ±11%     ~     (p=0.625 n=18+19)
    CountTortureOverlapping-48  6.92ms ±13%  8.42ms ±11%  +21.72%  (p=0.000 n=19+17)
    
    Change-Id: Ief120aee918a66487c76be56e0796871c8502f89
    Reviewed-on: https://go-review.googlesource.com/28586
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 86b2f29676c52774d91dda96e0ba5d4d7bcd3b47
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Thu Oct 6 11:06:21 2016 -0700

    database/sql: add support for multiple result sets
    
    Many database systems allow returning multiple result sets
    in a single query. This can be useful when dealing with many
    intermediate results on the server and there is a need
    to return more then one arity of data to the client.
    
    Fixes #12382
    
    Change-Id: I480a9ac6dadfc8743e0ba8b6d868ccf8442a9ca1
    Reviewed-on: https://go-review.googlesource.com/30592
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit be48aa3f3a16006ab31c424487af352ca374afed
Author: Rob Pike <r@golang.org>
Date:   Wed Oct 12 19:51:02 2016 -0700

    cmd/cover: handle gotos
    
    If a labeled statement is the target of a goto, we must treat it as the
    boundary of a new basic block, but only if it is not already the boundary
    of a basic block such as a labeled for loop.
    
    Fixes #16624
    
    Now reports 100% coverage for the test in the issue.
    
    Change-Id: If118bb6ff53a96c738e169d92c03cb3ce97bad0e
    Reviewed-on: https://go-review.googlesource.com/30977
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 5567b878915f7c2f1e7ee3898125c2cd2b7fe287
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Oct 14 11:27:11 2016 -0700

    spec: fix examples for predeclared function complex
    
    Fixes #17398.
    
    Change-Id: Iac7899031c1bfbadc4f84e5b374eaf1f01dff8c8
    Reviewed-on: https://go-review.googlesource.com/31190
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d13fa4d2256d6dfd030c03a82db258872e3e646c
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu May 5 12:44:28 2016 +1000

    os: use FindFirstFile when GetFileAttributesEx fails in Stat
    
    Fixes #15355
    
    Change-Id: Idbab7a627c5de249bb62d519c5a47f3d2f6c82a7
    Reviewed-on: https://go-review.googlesource.com/22796
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit abbd502d638262fa80e142ad18a89d6c75490672
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Oct 11 09:23:39 2016 +0000

    net/http: allow Handlers to test Hijacked conn without spamming error log
    
    Make a zero-byte write to a hijacked connection not log anything, so handlers
    can test whether a connection is hacked by doing a Write(nil).
    
    Fixes #16456
    
    Change-Id: Id56caf822c8592067bd8422672f0c1aec89e866c
    Reviewed-on: https://go-review.googlesource.com/30812
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 61f1a38bcb52ad5e1753b43c405bb5b144b6966c
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Oct 13 17:08:54 2016 -0400

    crypto/{aes,cipher}: fix panic in CBC on s390x when src length is 0
    
    Adds a test to check that block cipher modes accept a zero-length
    input.
    
    Fixes #17435.
    
    Change-Id: Ie093c4cdff756b5c2dcb79342e167b3de5622389
    Reviewed-on: https://go-review.googlesource.com/31070
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 714318be882d1974ec624272c8d905256d8a5932
Author: Filippo Valsorda <hi@filippo.io>
Date:   Thu Oct 13 17:56:04 2016 +0100

    expvar: add Value methods
    
    Closes #15815
    
    Change-Id: I08154dbff416198cf7787e446b1e00e62c03a972
    Reviewed-on: https://go-review.googlesource.com/30917
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 46276d6b6b14993b4851b77cf48c67ede3360e31
Author: Austin Clements <austin@google.com>
Date:   Mon Sep 26 18:09:54 2016 -0400

    doc: catch go1.8.txt up on runtime changes
    
    This clarifies some of the titles so they're more "news" friendly and
    less implementation-oriented.
    
    Change-Id: Ied02aa1e6824b04db5d32ecdd58e972515b1f588
    Reviewed-on: https://go-review.googlesource.com/29830
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 1da1da3d6ab9ac8287e6b4185afcd1c3396178f5
Author: Alan Donovan <adonovan@google.com>
Date:   Mon May 9 09:40:27 2016 -0400

    go/internal/gcimporter: set Pos attribute of decoded types.Objects
    
    This change is a copy of CL 22788 in x/tools.
    It has no observable effect yet, but brings the two packages in synch.
    
    Change-Id: I266c77547cb46deb69b1a36e1674dfebc430e3a5
    Reviewed-on: https://go-review.googlesource.com/22936
    Reviewed-by: Russ Cox <rsc@golang.org>

commit d08c3d1329994ddba2e0a2dd13b950f9f178ab02
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Fri Apr 29 16:23:55 2016 +1000

    cmd/link/internal/ld: use debug/pe package to rewrite ldpe.go
    
    This CL also includes vendored copy of debug/pe,
    otherwise bootstrapping fails.
    
    Updates #15345
    
    Change-Id: I3a8ac990e3cb12cb4d24ec11b618b68190397fd1
    Reviewed-on: https://go-review.googlesource.com/22603
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2427123d9391134a68c6f60b5acbcb6f60856ff8
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 12 22:42:40 2016 -0400

    encoding/xml: split attribute marshaling into its own method
    
    No functional changes here. Just makes next CL easier to read.
    
    Change-Id: Icf7b2281b4da6cb59ff4edff05943b2ee288576a
    Reviewed-on: https://go-review.googlesource.com/30945
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 26c7b4fb1e0553e8f9abe5fdd9008bb1f5bd3228
Author: Anthony Canino <anthony.canino1@gmail.com>
Date:   Sat Oct 10 15:24:34 2015 -0400

    cmd/compile: "abc"[1] is not an ideal constant
    
    "abc"[1] is not like 'b', in that -"abc"[1] is uint8 math, not ideal constant math.
    Delay the constantification until after ideal constant folding is over.
    
    Fixes #11370.
    
    Change-Id: Iba2fc00ca2455959e7bab8f4b8b4aac14b1f9858
    Reviewed-on: https://go-review.googlesource.com/15740
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0da30d5cbdd092499fe199c212f8799fd0cc676e
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 12 15:55:02 2016 -0400

    encoding/json: handle misspelled JSON literals in ,string
    
    Fixes #15146.
    
    Change-Id: I229611b9cc995a1391681c492c4d742195c787ea
    Reviewed-on: https://go-review.googlesource.com/30943
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3c1e1c30fdfbdaf7cf5f947c53245f1c28e56f91
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Oct 12 17:09:54 2016 -0700

    cmd/cgo: use alias for unsafe rather than separate functions
    
    When we need to generate a call to _cgoCheckPointer, we need to type
    assert the result back to the desired type. That is harder when the type
    is unsafe.Pointer, as the package can have values of unsafe.Pointer
    types without actually importing unsafe, by mixing C void* and :=. We
    used to handle this by generating a special function for each needed
    type, and defining that function in a separate file where we did import
    unsafe.
    
    Simplify the code by not generating those functions, but instead just
    import unsafe under the alias _cgo_unsafe. This is a simplification step
    toward a fix for issue #16591.
    
    Change-Id: I0edb3e04b6400ca068751709fe063397cf960a54
    Reviewed-on: https://go-review.googlesource.com/30973
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
```
