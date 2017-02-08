# September 22, 2016 Report

Number of commits: 136

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 589.338809ms to 565.86678ms, -3.98%
* github.com/coreos/etcd: from 13.087782903s to 12.818164414s, -2.06%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 31.168829013s to 34.865258419s, +11.86%
* github.com/influxdata/influxdb/cmd/influxd: from 7.237644614s to 6.540941147s, -9.63%
* github.com/junegunn/fzf/src/fzf: from 996.470341ms to 1.058238407s, +6.20%
* github.com/mholt/caddy/caddy: from 6.582485492s to 5.812851971s, -11.69%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.553959203s to 1.316896424s, -15.26%
* github.com/nsqio/nsq/apps/nsqd: from 2.047243994s to 2.108817024s, +3.01%
* github.com/prometheus/prometheus/cmd/prometheus: from 12.425055386s to 11.911162941s, -4.14%
* github.com/spf13/hugo: from 7.716383018s to 7.633568132s, -1.07%
* golang.org/x/tools/cmd/guru: from 2.742848441s to 2.84635903s, +3.77%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2590003 to 2573545, -0.64%
* github.com/coreos/etcd: from 23827152 to 23713920, -0.48%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 28449472 to 28308944, -0.49%
* github.com/influxdata/influxdb/cmd/influxd: from 16090932 to 16036333, -0.34%
* github.com/junegunn/fzf/src/fzf: from 3156544 to 3137168, -0.61%
* github.com/mholt/caddy/caddy: from 14647203 to 14600838, -0.32%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4394640 to 4361930, -0.74%
* github.com/nsqio/nsq/apps/nsqd: from 9657057 to 9643381, -0.14%
* github.com/prometheus/prometheus/cmd/prometheus: from 25248623 to 25091499, -0.62%
* github.com/spf13/hugo: from 16094308 to 16023328, -0.44%
* golang.org/x/tools/cmd/guru: from 7570536 to 7496866, -0.97%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               185           182           -1.62%
BenchmarkMsgpUnmarshal-4             398           369           -7.29%
BenchmarkEasyJsonMarshal-4           1447          1412          -2.42%
BenchmarkEasyJsonUnmarshal-4         2211          1549          -29.94%
BenchmarkFlatBuffersMarshal-4        355           479           +34.93%
BenchmarkFlatBuffersUnmarshal-4      287           281           -2.09%
BenchmarkGogoprotobufMarshal-4       154           155           +0.65%
BenchmarkGogoprotobufUnmarshal-4     251           245           -2.39%

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

* [cmd/compile: make ssa compilation unconditional](https://github.com/golang/go/commit/167e381f405d36f71ef152e45bb845b866592c80)
* [cmd/compile: unroll comparisons to short constant strings](https://github.com/golang/go/commit/c9fd997524ce7d531579500218f11b528bab4c88)
* [cmd/compile/internal/syntax: support for alias declarations](https://github.com/golang/go/commit/32db3f2756324616b7c856ac9501deccc2491239)

## GIT Log

```
commit dcbbd319e9cdd44d50314818ec05672b60e8f8e7
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Thu Sep 15 07:39:19 2016 -0700

    compress/gzip: add examples
    
    Updates #16360.
    
    Adds examples uing:
    + Writer, Reader
    + Reader.Multistream to concatenate and then
    individually retrieve multiple gzipped files
    + Reset
    
    Change-Id: I9ad9b92729a5cd58f7368eaf2db05f1cdf21063d
    Reviewed-on: https://go-review.googlesource.com/29218
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a7dc821d6447f9bdfb54e2fad2ab76f6d40873ca
Author: Nigel Tao <nigeltao@golang.org>
Date:   Thu Sep 22 10:58:43 2016 +1000

    crypto/rsa: clarify comment on maximum message length.
    
    See https://groups.google.com/d/topic/golang-nuts/stbum5gZbAc/discussion
    
    Change-Id: I2e78e8d0dadd78c8b0389514cad3c45d061b663b
    Reviewed-on: https://go-review.googlesource.com/29496
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9e028b70ea3c1a81220a63c2ed3a2c1cd216c733
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Sep 21 14:21:11 2016 -0700

    runtime: merge signal[12]_unix.go into signal_unix.go
    
    Requires adding a sigfwd function for Solaris, as previously
    signal2_unix.go was not built for Solaris.
    
    Change-Id: Iea3ff0ddfa15af573813eb075bead532b324a3fc
    Reviewed-on: https://go-review.googlesource.com/29550
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5db80c30e6b360ef18159332b3ff19234f911f36
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Sep 21 12:53:25 2016 +0900

    runtime: revert CL 18835; don't install new signal stack unconditionally on dragonfly
    
    This change reverts CL 18835 which is a workaroud for older DragonFly
    BSD kernels, and fixes #14051, #14052 and #14067 in a more general way
    the same as other platforms except NetBSD.
    
    This change also bumps the minimum required version of DragonFly BSD
    kernel to 4.4.4.
    
    Fixes #16329.
    
    Change-Id: I0b44b6afa675f5ed9523914226bd9ec4809ba5ae
    Reviewed-on: https://go-review.googlesource.com/29491
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e6134702bb63603a3a26b7167dab73d145886c0f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 01:42:58 2016 -0700

    cmd/compile: simplify obj.ProgInfo and extract from obj.Prog
    
    Updates #16357.
    
    Change-Id: Ia837dd44bad76931baa9469e64371bc253d6694b
    Reviewed-on: https://go-review.googlesource.com/29219
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 404ae84aa220943c85fe04e9d7499087aedb0677
Author: Russ Cox <rsc@golang.org>
Date:   Wed Sep 21 15:03:52 2016 -0400

    cmd/link: add time stamp to hostlink print in -v mode
    
    Change-Id: I128b142aee5e1b917e7ba63b48512972f053ea0b
    Reviewed-on: https://go-review.googlesource.com/29531
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b4efd09d1880793e33fbb191ccfe1657bfeba0c9
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Thu Aug 25 11:07:33 2016 -0500

    cmd/link: split large elf text sections on ppc64x
    
    Some applications built with Go on ppc64x with external linking
    can fail to link with relocation truncation errors if the elf
    text section that is generated is larger than 2^26 bytes and that
    section contains a call instruction (bl) which calls a function
    beyond the limit addressable by the 24 bit field in the
    instruction.
    
    This solution consists of generating multiple text sections where
    each is small enough to allow the GNU linker to resolve the calls
    by generating long branch code where needed.  Other changes were added
    to handle differences in processing when multiple text sections exist.
    
    Some adjustments were required to the computation of a method's address
    when using the method offset table when there are multiple text sections.
    
    The number of possible section headers was increased to allow for up
    to 128 text sections.  A test case was also added.
    
    Fixes #15823.
    
    Change-Id: If8117b0e0afb058cbc072258425a35aef2363c92
    Reviewed-on: https://go-review.googlesource.com/27790
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 445f51fb119789e08889abcb8ae54b4eb5029fad
Author: Suyash <dextrous93@gmail.com>
Date:   Sun Sep 11 12:47:47 2016 +0530

    image/png: add Encode and Decode examples
    
    partially addresses #16360
    
    Change-Id: I8274825b9ca6aec46294c8513b4795b0eb3062a2
    Reviewed-on: https://go-review.googlesource.com/28992
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2bc5f1258e049e1d59711046e5c865500eefa159
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 20 19:45:37 2016 +0000

    net: add Resolver type, Dialer.Resolver, and DefaultResolver
    
    The new Resolver type (a struct) has 9 Lookup methods, all taking a
    context.Context.
    
    There's now a new DefaultResolver global, like http's
    DefaultTransport and DefaultClient.
    
    net.Dialer now has an optional Resolver field to set the Resolver.
    
    This also does finishes some resolver cleanup internally, deleting
    lookupIPMerge and renaming lookupIPContext into Resolver.LookupIPAddr.
    
    The Resolver currently doesn't let you tweak much, but it's a struct
    specifically so we can add knobs in the future. Currently I just added
    a bool to force the pure Go resolver. In the future we could let
    people provide an interface to implement the methods, or add a Timeout
    time.Duration, which would wrap all provided contexts in a
    context.WithTimeout.
    
    Fixes #16672
    
    Change-Id: I7ba1f886704f06def7b6b5c4da9809db51bc1495
    Reviewed-on: https://go-review.googlesource.com/29440
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ea143c299040f8a270fb782c5efd3a3a5e6057a4
Author: Thomas de Zeeuw <thomasdezeeuw@gmail.com>
Date:   Thu Sep 1 14:54:08 2016 +0200

    net/http/httptest: fill ContentLength in recorded Response
    
    This change fills the ContentLength field in the http.Response returned by
    ResponseRecorder.Result.
    
    Fixes #16952.
    
    Change-Id: I9c49b1bf83e3719b5275b03a43aff5033156637d
    Reviewed-on: https://go-review.googlesource.com/28302
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e69d63e807b7322e3a7d6c8ad83c251c9f46c9ca
Author: Michal Bohuslávek <mbohuslavek@gmail.com>
Date:   Wed Sep 21 14:49:51 2016 +0100

    net/http/cookiejar: fix typo
    
    Change-Id: I6ea8650927e7946c6fd4659f400fd91ddaae68af
    Reviewed-on: https://go-review.googlesource.com/29510
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c03925edd37b504346c11656a6de5f5e4b791061
Author: Austin Clements <austin@google.com>
Date:   Mon May 9 15:12:07 2016 -0400

    runtime: remove unnecessary atomics from heapBitSetType
    
    These used to be necessary when racing with updates to the mark bit,
    but since the mark bit is no longer in the bitmap and the checkmark is
    only updated with the world stopped, we can now always use regular
    writes to update the type information in the heap bitmap.
    
    Somewhat surprisingly, this has basically no overall performance
    effect beyond the usual noise, but it does clean up the code.
    
    Change-Id: I3933d0b4c0bc1c9bcf6313613515c0b496212105
    Reviewed-on: https://go-review.googlesource.com/29277
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 43d9c29abb57e797075d9cc15f6e21362d4be136
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Sep 21 13:04:38 2016 +0900

    doc: add note about CL 29491 to go1.8.txt
    
    Change-Id: I808fab97076493a95b0b5eb0ad15645099f54aee
    Reviewed-on: https://go-review.googlesource.com/29492
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6072e4d710dc28fe5907d43e59074d5b55de8981
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Sep 21 12:53:25 2016 +0900

    syscall: fix build on dragonfly
    
    This change fixes the broken build caused by CL 23780.
    
    Change-Id: I142cf8a1af033d036d57ac56e9e21ea925d922d4
    Reviewed-on: https://go-review.googlesource.com/29490
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 35d22afb4b8e4f2c0ce06150727dc91a5c54378e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Sep 20 18:12:27 2016 -0700

    cmd/internal/obj: remove unused GOROOT-related fields
    
    Change-Id: I6634f70d6bd1a4eced47eda69a2d9b207d222a1b
    Reviewed-on: https://go-review.googlesource.com/29470
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit e6158b3c469dc4ec5d6571f10fbd739acd493472
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Sep 20 18:02:27 2016 -0700

    cmd/internal/obj: remove unused Textp and Etextp fields
    
    Change-Id: Idcb5a8d6676aa38b4ebd0975edd2068386f5ca83
    Reviewed-on: https://go-review.googlesource.com/29449
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b851ded09a300033849b60ab47a468087ce557a1
Author: Yasuhiro Matsumoto <mattn.jp@gmail.com>
Date:   Wed Aug 24 10:34:16 2016 +0900

    os: use GetConsoleCP() instead of GetACP()
    
    It is possible (and common) for Windows systems to use a different codepage
    for console applications from that used on normal windowed application
    (called ANSI codepage); for instance, most of the western Europe uses
    CP850 for console (for backward compatibility with MS-DOS), while
    windowed applications use a different codepage depending on the country
    (eg: CP1252 aka Latin-1). The usage being changed with this commit is
    specifically related to decoding input coming from the console, so the
    previous usage of the ANSI codepage was wrong.
    
    Also fixes an issue that previous did convert bytes as NFD. Go is
    designed to handle single Unicode code point. This fix change behaivor
    to NFC.
    
    Fixes #16857.
    
    Change-Id: I4f41ae83ece47321b6e9a79a2087ecbb8ac066dd
    Reviewed-on: https://go-review.googlesource.com/27575
    Reviewed-by: Hiroshi Ioka <hirochachacha@gmail.com>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 16f81b617eadc739b9097ce4b7d67e9a00a91c7c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 20 23:43:15 2016 +0000

    image/draw: add FloydSteinberg Drawer example
    
    Updates #16360
    
    Change-Id: I80b981aa291a8e16d2986d4a2dfd84d3819bf488
    Reviewed-on: https://go-review.googlesource.com/29443
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit e7191479ec57e89aa1967e75a87f9da4d1d8b734
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 20 21:52:11 2016 +0000

    doc: add some missing HTML tags in the FAQ
    
    Fixes #17170
    
    Change-Id: I939f087df133710495fdf6f09040051cb9b176d7
    Reviewed-on: https://go-review.googlesource.com/29442
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit dd24b1098ae8923b98a13560c89ae59fc0c49258
Author: Keith Randall <khr@golang.org>
Date:   Wed Sep 7 14:04:31 2016 -0700

    cmd/compile: improve tighten pass
    
    Move a value to the block which is the lowest common ancestor in the
    dominator tree of all of its uses.  Make sure not to move a value into a
    loop.
    
    Makes the tighten pass on average (across go1 benchmarks) 40% slower.
    Still not a big contributor to overall compile time.
    
    Binary size is just a tad smaller.
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.77s ± 9%     2.76s ± 9%     ~     (p=0.878 n=8+8)
    Fannkuch11-12                2.75s ± 1%     2.74s ± 1%     ~     (p=0.232 n=8+7)
    FmtFprintfEmpty-12          48.9ns ± 9%    47.7ns ± 0%     ~     (p=0.431 n=8+8)
    FmtFprintfString-12          143ns ± 8%     142ns ± 1%     ~     (p=0.257 n=8+7)
    FmtFprintfInt-12             123ns ± 1%     122ns ± 1%   -1.04%  (p=0.026 n=7+8)
    FmtFprintfIntInt-12          195ns ± 7%     185ns ± 0%   -5.32%  (p=0.000 n=8+8)
    FmtFprintfPrefixedInt-12     194ns ± 4%     195ns ± 0%   +0.81%  (p=0.015 n=7+7)
    FmtFprintfFloat-12           267ns ± 0%     268ns ± 0%   +0.37%  (p=0.001 n=7+6)
    FmtManyArgs-12               800ns ± 0%     762ns ± 1%   -4.78%  (p=0.000 n=8+8)
    GobDecode-12                7.67ms ± 2%    7.60ms ± 2%     ~     (p=0.234 n=8+8)
    GobEncode-12                6.55ms ± 0%    6.57ms ± 1%     ~     (p=0.336 n=7+8)
    Gzip-12                      237ms ± 0%     238ms ± 0%   +0.40%  (p=0.017 n=7+7)
    Gunzip-12                   40.8ms ± 0%    40.2ms ± 0%   -1.52%  (p=0.000 n=7+8)
    HTTPClientServer-12          208µs ± 3%     209µs ± 3%     ~     (p=0.955 n=8+7)
    JSONEncode-12               16.2ms ± 1%    17.2ms ±11%   +5.80%  (p=0.001 n=7+8)
    JSONDecode-12               57.3ms ±12%    55.5ms ± 3%     ~     (p=0.867 n=8+7)
    Mandelbrot200-12            4.68ms ± 6%    4.46ms ± 1%     ~     (p=0.442 n=8+8)
    GoParse-12                  4.27ms ±44%    3.42ms ± 1%  -19.95%  (p=0.005 n=8+8)
    RegexpMatchEasy0_32-12      75.1ns ± 0%    75.8ns ± 1%   +0.99%  (p=0.002 n=7+7)
    RegexpMatchEasy0_1K-12       963ns ± 0%    1021ns ± 6%   +5.98%  (p=0.001 n=7+7)
    RegexpMatchEasy1_32-12      72.4ns ±11%    70.8ns ± 1%     ~     (p=0.368 n=8+8)
    RegexpMatchEasy1_1K-12       394ns ± 1%     399ns ± 0%   +1.23%  (p=0.000 n=8+7)
    RegexpMatchMedium_32-12      114ns ± 0%     115ns ± 1%   +0.63%  (p=0.021 n=7+7)
    RegexpMatchMedium_1K-12     35.9µs ± 0%    37.6µs ± 1%   +4.72%  (p=0.000 n=7+8)
    RegexpMatchHard_32-12       1.93µs ± 2%    1.91µs ± 0%   -0.91%  (p=0.001 n=7+7)
    RegexpMatchHard_1K-12       60.2µs ± 3%    61.2µs ±10%     ~     (p=0.442 n=8+8)
    Revcomp-12                   404ms ± 1%     406ms ± 1%     ~     (p=0.054 n=8+7)
    Template-12                 64.6ms ± 1%    63.5ms ± 1%   -1.66%  (p=0.000 n=8+8)
    TimeParse-12                 347ns ± 8%     309ns ± 0%  -11.13%  (p=0.000 n=8+7)
    TimeFormat-12                343ns ± 4%     331ns ± 0%   -3.34%  (p=0.000 n=8+7)
    
    Change-Id: Id6da1239ddd4d0cb074ff29cffb06302d1c6d08f
    Reviewed-on: https://go-review.googlesource.com/28712
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit b7426089e597d20bade4b4bbfea1188844a07af8
Author: Keith Randall <khr@golang.org>
Date:   Tue Sep 20 13:59:40 2016 -0700

    cmd/compile: simple cleanups
    
    Change-Id: If2cf3c5a29afc6cf74c3b08b9745e950231ead37
    Reviewed-on: https://go-review.googlesource.com/29441
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f5f7d6e32d5384a3638325ff8393bf94ec8d6971
Author: Damien Neil <dneil@google.com>
Date:   Fri Jun 3 15:04:53 2016 -0700

    syscall: validate ParseDirent inputs
    
    Don't panic, crash, or return references to uninitialized memory when
    ParseDirent is passed invalid input.
    
    Move common dirent parsing to syscall.go with minimal platform-specific
    functions in syscall_$GOOS.go.
    
    Fixes #15653
    
    Change-Id: I5602475e02321fe381064488401c14b33bec6886
    Reviewed-on: https://go-review.googlesource.com/23780
    Run-TryBot: Damien Neil <dneil@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ab5923572984651af05a47755109642bfc529cb5
Author: Austin Clements <austin@google.com>
Date:   Thu Aug 25 11:50:50 2016 -0400

    runtime: consistency check for G rescan position
    
    Issue #17099 shows a failure that indicates we rescanned a stack twice
    concurrently during mark termination, which suggests that the rescan
    list became inconsistent. Add a simple check when we dequeue something
    from the rescan list that it claims to be at the index where we found
    it.
    
    Change-Id: I6a267da4154a2e7b7d430cb4056e6bae978eaf62
    Reviewed-on: https://go-review.googlesource.com/29280
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 39ce6eb9ec53e53e361824feb96cbbc6d236df5a
Author: Austin Clements <austin@google.com>
Date:   Thu Aug 25 11:52:24 2016 -0400

    runtime: report GCSys and OtherSys in heap profile
    
    The comment block at the end of the heap profile includes *almost*
    everything from MemStats. Add the missing fields. These are useful for
    debugging RSS that has gone to GC-internal data structures.
    
    Change-Id: I0ee8a918d49629e28fd8fd2bf6861c4529461c24
    Reviewed-on: https://go-review.googlesource.com/29276
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 38cd79889ece342643b56ad6d496ef8931ca9272
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Sep 19 07:45:08 2016 -0400

    cmd/compile: simplify div/mod on ARM
    
    On ARM, DIV, DIVU, MOD, MODU are pseudo instructions that makes
    runtime calls _div/_udiv/_mod/_umod, which themselves are wrappers
    of udiv. The udiv function does the real thing.
    
    Instead of generating these pseudo instructions, call to udiv
    directly. This removes one layer of wrappers (which has an awkward
    way of passing argument), and also allows combining DIV and MOD
    if both results are needed.
    
    Change-Id: I118afc3986db3a1daabb5c1e6e57430888c91817
    Reviewed-on: https://go-review.googlesource.com/29390
    Reviewed-by: David Chase <drchase@google.com>

commit f964810025ae58b623798a1944c39c06266fb45d
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Mon Sep 19 20:36:45 2016 -0700

    net/http/httptrace: fix bad tracing example
    
    Tracing happens at the http.Trace level. Fix the example to demostrate
    tracing in the lifecycle of a RoundTrip.
    
    Updates #17152.
    
    Change-Id: Ic7d7bcc550176189206185482e8962dbf1504ff1
    Reviewed-on: https://go-review.googlesource.com/29431
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 836a3ae6639c310e1a13834c1f8f84bb982d920d
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue Sep 20 15:57:53 2016 +1200

    cmd/link: remove more unused ctxt parameters
    
    This time in elf.go.
    
    Change-Id: Ifaf71742ebbc9aadc8606c39ea2d417ae5cc7e0d
    Reviewed-on: https://go-review.googlesource.com/29450
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b6324ef5ff212bea011d4d60ff8aeaa0ef7a5bba
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue Sep 20 15:46:37 2016 +1200

    cmd/link: kill off Symbols.Version
    
    Change-Id: Iee8f773355870f2333637a093e51c5fd36e5a6e5
    Reviewed-on: https://go-review.googlesource.com/29349
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 63837092726fa01806ee945dc7b3a2c8c707cd76
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue Sep 20 15:31:26 2016 +1200

    cmd/link: remove now-unused ctxt arguments from a few functions
    
    Specifically Addstring, Addbytes and Symgrow.
    
    Change-Id: Ia74093bfcf9f360bf223accbc8feef54a7f059c9
    Reviewed-on: https://go-review.googlesource.com/29348
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 25d094034bd5d0e4317180a70d41bf774dae1598
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Mon Sep 19 12:12:04 2016 +1200

    cmd/link: remove Linklookup & Linkrlookup
    
    Change-Id: I25d9f74cb52e6fd4f2ad4b1c8b7102efadbc7481
    Reviewed-on: https://go-review.googlesource.com/29344
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 324f6ab48c9558a417876b62fd2f633943ed3bb8
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue Sep 20 15:06:08 2016 +1200

    cmd/link: use ctxt.{Lookup,ROLookup} in favour of function versions of same
    
    Done with two eg templates:
    
    package p
    
    import (
            "cmd/link/internal/ld"
    )
    
    func before(ctxt *ld.Link, name string, v int) *ld.Symbol {
            return ld.Linklookup(ctxt, name, v)
    }
    func after(ctxt *ld.Link, name string, v int) *ld.Symbol {
            return ctxt.Syms.Lookup(name, v)
    }
    
    package p
    
    import (
            "cmd/link/internal/ld"
    )
    
    func before(ctxt *ld.Link, name string, v int) *ld.Symbol {
            return ld.Linkrlookup(ctxt, name, v)
    }
    func after(ctxt *ld.Link, name string, v int) *ld.Symbol {
            return ctxt.Syms.ROLookup(name, v)
    }
    
    Change-Id: I00647dbf62294557bd24c29ad1f108fc786335f1
    Reviewed-on: https://go-review.googlesource.com/29343
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d41a7f77c5ed5114fe81b65ee28013227d20d637
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue Sep 20 14:59:39 2016 +1200

    cmd/link: do not directly embed Symbols in Link
    
    Mostly done with sed.
    
    Change-Id: Ic8c534a3fdd332b5420d062ee85bb77a30ad1efb
    Reviewed-on: https://go-review.googlesource.com/29346
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d284d4ff9204e42c3cfc5424880d74063ed08b23
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Mon Sep 19 12:02:58 2016 +1200

    cmd/link: split "bag of Symbols" functionality out of Link
    
    Mechanical refactorings to follow.
    
    Change-Id: I9b98e69a58c3cba7c7d1d3e3f600d4ed99d4fce2
    Reviewed-on: https://go-review.googlesource.com/29342
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e9fddf8f863c17e7112e8dd0a52490ecf165ef13
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 19 14:11:20 2016 -0400

    cmd/internal/obj, cmd/link: darwin dynlink support
    
    This makes it possible for cmd/compile, when run with -dynlink on
    darwin/amd64, to generate TLS_LE relocations which the linker then
    turns into the appropriate PC-relative GOT load.
    
    Change-Id: I1a71da432608bdb108ff66c22de600100209c873
    Reviewed-on: https://go-review.googlesource.com/29393
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1d3fae461c6302d2d91085e03d50712177af74c0
Author: David Crawshaw <crawshaw@golang.org>
Date:   Sat Sep 17 10:01:17 2016 -0400

    cmd/link: remove Cursym
    
    Change-Id: I58253a6cd2d77a9319c0783afb0d92cd5a88a7f7
    Reviewed-on: https://go-review.googlesource.com/29370
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 5a597e2d2995309c4cc02a4c367157908ad4add2
Author: David Crawshaw <crawshaw@golang.org>
Date:   Sat Sep 17 09:39:33 2016 -0400

    cmd/link: replace ld.Link.Diag with ld.Errorf
    
    Instead of using ctxt.Cursym, Errorf takes an explicit *Symbol
    parameter. This removes most uses of Cursym and means the *Link
    context object is needed in fewer parts of the linker.
    
    All transformations done manually, as wiring Cursym is tricky.
    
    Change-Id: Ief88b00b73904224675c0035684c3a84c19249d7
    Reviewed-on: https://go-review.googlesource.com/29369
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bd3d00e3109335bd66ec6dbede411c19f37df20f
Author: Keith Randall <khr@golang.org>
Date:   Mon Sep 19 15:26:48 2016 -0700

    doc: mention KeepAlive & input args change
    
    Change-Id: Icfb38f492ae71432858b7104fcba18a9ba835192
    Reviewed-on: https://go-review.googlesource.com/29410
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 50644f2e67e846c3641d5a907733ccc4ca5e60d2
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Mon Sep 19 07:13:27 2016 +1200

    cmd/link: move comment somewhere more appropriate
    
    At least, I assume it's meant to be here. It makes no sense at all where it
    currently is.
    
    Change-Id: Ic6a6c112c3dcf1318256d7d602168c3446b55412
    Reviewed-on: https://go-review.googlesource.com/29339
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b625810d28adde8c24da4838525694e74e131520
Author: Tormod Erevik Lea <tormodlea@gmail.com>
Date:   Sat Sep 17 19:04:46 2016 +0200

    cmd/vet: hard-code program name in usage message
    
    Example on linux_amd64 for 'go tool vet -h':
    
    Before:
            Usage of /usr/local/go/pkg/tool/linux_amd64/vet:
    After:
            Usage of vet:
    
    Change-Id: I11cb16b656bd097062d57a8c7441fbe66caaef78
    Reviewed-on: https://go-review.googlesource.com/29294
    Reviewed-by: Rob Pike <r@golang.org>

commit 3cca069220af044c1a36da3f588ffe3abbeab9c5
Author: Rob Pike <r@golang.org>
Date:   Sun Sep 18 14:35:42 2016 +1000

    time: allow long fractions in ParseDuration
    
    The code scanned for an integer after a decimal point, which
    meant things could overflow if the number was very precise
    (0.1234123412341234123412342134s). This fix changes the
    parser to stop adding precision once we run out of bits, rather
    than trigger an erroneous overflow.
    
    We could parse durations using floating-point arithmetic,
    but since the type is int64 and float64 has only has 53 bits
    of precision, that would be imprecise.
    
    Fixes #15011.
    
    Change-Id: If85e22b8f6cef12475e221169bb8f493bb9eb590
    Reviewed-on: https://go-review.googlesource.com/29338
    Reviewed-by: Costin Chirvasuta <costinc@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e94c52933b9c414d3f8fa94ead0d9cc5b7d7d717
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Sep 16 21:42:18 2016 -0400

    cmd/compile: intrinsify Ctz{32,64} and Bswap{32,64} on s390x
    
    Also adds the 'find leftmost one' instruction (FLOGR) and replaces the
    WORD-encoded use of FLOGR in math/big with it.
    
    Change-Id: I18e7cd19e75b8501a6ae8bd925471f7e37ded206
    Reviewed-on: https://go-review.googlesource.com/29372
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f1973fca717f2c3f3f10bcc2bc3512a4c549710b
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Tue Apr 12 18:38:00 2016 -0300

    cmd/asm, cmd/internal/obj/ppc64: add ppc64 vector registers and instructions
    
    The current implementation for Power architecture does not include the vector
    (Altivec) registers.  This adds the 32 VMX registers and the most commonly used
    instructions: X-form loads/stores; VX-form logical operations, add/sub,
    rotate/shift, count, splat, SHA Sigma and AES cipher; VC-form compare; and
    VA-form permute, shift, add/sub and select.
    
    Fixes #15619
    
    Change-Id: I544b990631726e8fdfcce8ecca0aeeb72faae9aa
    Reviewed-on: https://go-review.googlesource.com/25600
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>
    Reviewed-by: David Chase <drchase@google.com>

commit 31ba855014c62ed8ea2a19208d43318d99948e5b
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Sep 18 16:51:02 2016 -0700

    crypto/md5, crypto/sha1, crypto/sha256: add examples for checksumming a file
    
    Updates #16360.
    
    Change-Id: I75714d2b5f095fe39fd81edfa6dd9e44d7c44da1
    Reviewed-on: https://go-review.googlesource.com/29375
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ca4089ad62b806db7d3f32335d3f20865a75edcd
Author: Keith Randall <khr@golang.org>
Date:   Wed Aug 31 15:17:02 2016 -0700

    cmd/compile: args no longer live until end-of-function
    
    We're dropping this behavior in favor of runtime.KeepAlive.
    Implement runtime.KeepAlive as an intrinsic.
    
    Update #15843
    
    Change-Id: Ib60225bd30d6770ece1c3c7d1339a06aa25b1cbc
    Reviewed-on: https://go-review.googlesource.com/28310
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit faf611a07a7630a075fba3a555db7831e002122a
Author: Brady Sullivan <brady@bsull.com>
Date:   Fri Sep 16 15:25:07 2016 -0700

    net/http: rename Post's parameter from bodyType to contentType
    
    Change-Id: Ie1b08215c02ce3ec72a4752f4b800f23345ff99d
    Reviewed-on: https://go-review.googlesource.com/29362
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 75ce89c20dab10857ab0b5102001b34767c45b6e
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 16 13:50:18 2016 -0700

    cmd/compile: cache CFG-dependent computations
    
    We compute a lot of stuff based off the CFG: postorder traversal,
    dominators, dominator tree, loop nest.  Multiple phases use this
    information and we end up recomputing some of it.  Add a cache
    for this information so if the CFG hasn't changed, we can reuse
    the previous computation.
    
    Change-Id: I9b5b58af06830bd120afbee9cfab395a0a2f74b2
    Reviewed-on: https://go-review.googlesource.com/29356
    Reviewed-by: David Chase <drchase@google.com>

commit 2679282da4e437ee086ec791ab73181c39ae3463
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 16 12:08:05 2016 -0700

    cmd/compile: fold ADDconsts together for PPC
    
    Change-Id: I571f03af6f791e78e7e18addcc310eb25747cdcf
    Reviewed-on: https://go-review.googlesource.com/29351
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6129f37367686edf7c2732fbb5300d5f28203743
Author: Keith Randall <khr@golang.org>
Date:   Sat Sep 17 15:04:36 2016 -0700

    cmd/compile: inline convT2{I,E} when result doesn't escape
    
    No point in calling a function when we can build the interface
    using a known type (or itab) and the address of a local.
    
    Get rid of third arg (preallocated stack space) to convT2{I,E}.
    
    Makes go binary smaller by 0.2%
    
    benchmark                   old ns/op     new ns/op     delta
    BenchmarkEfaceInteger-8     16.7          10.1          -39.52%
    
    Update #17118
    Update #15375
    
    Change-Id: I9724a1f802bfa1e3957bf1856b55558278e198a2
    Reviewed-on: https://go-review.googlesource.com/29373
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 892d146a7aae17e5fe22e04b16ba4da7e3d8c767
Author: Rob Pike <r@golang.org>
Date:   Mon Sep 19 11:22:52 2016 +1000

    cmd/vet: fix documentation for -structtags
    
    Was missing a title in the documentation, so it formatted wrong.
    
    Fixes #17124
    
    Change-Id: Ie8a9c36fbc54eed7d8a761f89a088e582b8c062d
    Reviewed-on: https://go-review.googlesource.com/29340
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit dcb954c3f7ab6d463882dbc3f88e075e8e096f74
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Sep 16 16:47:28 2016 +1200

    cmd/link: remove size and version from genasmsym's argument
    
    They are trivially available in the few places they are needed.
    
    Change-Id: I6544692e9027076ec9e6e9a295c66457039e55e1
    Reviewed-on: https://go-review.googlesource.com/29332
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 2266047556e7bc32e828dbfc4accdd1d4669f137
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Sep 16 16:22:08 2016 +1200

    cmd/link: give names and a type to the symbol types used by genasmsym
    
    Doing this revealed some dead code.
    
    Change-Id: I5202fcc3f73e3dfddfea3ec7b772e16da51195da
    Reviewed-on: https://go-review.googlesource.com/29331
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c1bee49cac710af467f6450e104f920eafe344f7
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Sep 18 15:28:41 2016 -0400

    cmd/dist: fix internal linking check for mips64le
    
    Fix mips64le build.
    
    Change-Id: Icf1b4901655463f582b49054a88edfb06ad6c676
    Reviewed-on: https://go-review.googlesource.com/29281
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 964639cc338db650ccadeafb7424bc8ebb2c0f6c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Sep 15 19:50:59 2016 -0400

    cmd/compile: intrinsify runtime/internal/atomic.Xaddint64
    
    This aliases to Xadd64.
    
    Change-Id: I95d49e1d03eecc242e9e6fd4b2742b1c1a1d5ade
    Reviewed-on: https://go-review.googlesource.com/29274
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 5c0fbf052b8ab98c0a1dd74365912d33572be4f3
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 16 15:08:17 2016 -0400

    plugin: cast dlerror return value for android
    
    Until a few weeks ago, bionic, the Andoid libc, incorrectly
    returned const char* (instead of char*) from dlerror(3).
    
    https://android.googlesource.com/platform/bionic/+/5e071a18ce88d93fcffaebb9e0f62524ae504908
    
    Change-Id: I30d33240c63a9f35b6c20ca7e3928ad33bc5e33f
    Reviewed-on: https://go-review.googlesource.com/29352
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 2e2db7a1704773082db547cbde70d8c0ce36a10c
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 16 16:56:29 2016 -0700

    cmd/compile: fix format verbs in ssa package
    
    %s is no longer valid.  Use %v instead.
    
    Change-Id: I5ec4fa6a9280082c1a0c75fd1cf94b4bb8096f5c
    Reviewed-on: https://go-review.googlesource.com/29365
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 246074d043f686c532cac88dccd68e01048a23bc
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Sep 16 19:19:42 2016 -0700

    cmd/internal/obj: remove ACHECKNIL
    
    Updates #16357.
    
    Change-Id: I35f938d675ca5c31f65c4419ee0732bbc593b5cb
    Reviewed-on: https://go-review.googlesource.com/29368
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e518962a27ce0333d42833456e09557ba37a95a0
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Sep 16 18:31:25 2016 -0700

    cmd/internal/obj: simplify Plists
    
    Keep Plists in a slice instead of a linked list.
    Eliminate unnecessary fields.
    Also, while here remove gc's unused breakpc and continpc vars.
    
    Change-Id: Ia04264036c0442843869965d247ccf68a5295115
    Reviewed-on: https://go-review.googlesource.com/29367
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit fc5df089da6c02397897f11a875a593353dc0590
Author: Trey Lawrence <lawrence.trey@gmail.com>
Date:   Tue Aug 23 16:43:43 2016 -0400

    cmd/compile: fix compiler bug for constant equality comparison
    
    The compiler incorrectly will error when comparing a nil pointer
    interface to a nil pointer of any other type. Example:
    (*int)(nil) == interface{}(nil)
    Will error with "gc: illegal constant expression: *int == interface {}"
    
    Fixes #16702
    
    Change-Id: I1a15d651df2cfca6762b1783a28b377b2e6ff8c6
    Reviewed-on: https://go-review.googlesource.com/27591
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6fe1febc867237fdf9ae40483044ed377144627f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Sep 16 15:31:04 2016 -0700

    cmd/internal/obj: replace AGLOBL with (*Link).Globl
    
    Replace the AGLOBL pseudo-op with a method to directly register an
    LSym as a global. Similar to how we previously already replaced the
    ADATA pseudo-op with directly writing out data bytes.
    
    Passes toolstash -cmp.
    
    Change-Id: I3631af0a2ab5798152d0c26b833dc309dbec5772
    Reviewed-on: https://go-review.googlesource.com/29366
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a1bf203b57ab854f2c1b7668ae96bc34e60e02fc
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Thu Sep 15 17:24:42 2016 -0300

    net: respect resolv.conf rotate option
    
    Instead of ranging over servers in the config, grab an offset
    from the config that is used to determine indices.
    
    When the rotate option is enabled, the offset increases which
    rotates queries through servers. Otherwise, it is always 0
    which uses servers in config order.
    
    Fixes #17126
    
    Change-Id: If233f6de7bfa42f88570055b9ab631be08a76b3e
    Reviewed-on: https://go-review.googlesource.com/29233
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f7e49f6644bde3f17b4a795218a35876347455a1
Author: Suyash <dextrous93@gmail.com>
Date:   Fri Sep 16 19:54:44 2016 +0530

    sort: fix search descending order example
    
    Change-Id: I27b82d8c63a06ddf7e148b15853aba24a22a6282
    Reviewed-on: https://go-review.googlesource.com/29336
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 569340ebafcd20c01227ed9cbeb5ccd29c52f2bc
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 16 18:53:46 2016 -0400

    cmd/dist: disable plugin test on arm64
    
    Mysterious error and no time or easy machine access to investigate
    now, so disabling the -buildmode=plugin test on arm64. (The arm
    version is working as expected.)
    
    Updates #17138
    
    Change-Id: I4cc56ddf47e7597213462e48d4934a765168bd07
    Reviewed-on: https://go-review.googlesource.com/29363
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 192548a5470d013f8782c7f1c8553338f7234fc7
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 16 18:25:52 2016 -0400

    cmd/dist: unify internal linking checks
    
    I missed one in CL 29360.
    
    Change-Id: I29fc6dcd920829a918c70734d646119133a0a9df
    Reviewed-on: https://go-review.googlesource.com/29361
    Reviewed-by: Keith Randall <khr@golang.org>

commit f4748f1e218c3584348f2f5c297e73f494182e54
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 16 17:15:12 2016 -0400

    cmd/dist: skip libgcc test on arm64 and mips64
    
    This test was always being run with external linking on these
    platforms because the linker was silently forcing external linking
    until CL 28971. Now it produces an error instead.
    
    Change-Id: I794e0812711e05b150daf805dc3451507bb4cae8
    Reviewed-on: https://go-review.googlesource.com/29360
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fb4e49059c28fc2de64f0c1b83b178f4bede4a75
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 16 17:10:44 2016 -0400

    cmd/dist: disable plugin test on ppc64le and s390x
    
    These are close to working, but untested and failing on
    build.golang.org. So disable for now.
    
    Change-Id: I330b8d1a91f0bf5139c894913868f01ec87e986d
    Reviewed-on: https://go-review.googlesource.com/29359
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 08feadbd6e724fd0377f37fbe88a3114dd7b45bb
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Sep 16 14:08:10 2016 -0700

    cmd/compile: remove Arch.REGCTXT
    
    Update #16357.
    
    Change-Id: I507676212d7137a62c76de7bfa0ba8dbd68e840f
    Reviewed-on: https://go-review.googlesource.com/29358
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit f323a7c6d28cb536e1924ccd2566617e4c0dec34
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 16 15:24:56 2016 -0400

    cmd/link: remove never-set Windows variable
    
    Change-Id: I3d64549b85b71bb63bcc815ce2276af6ca2eb215
    Reviewed-on: https://go-review.googlesource.com/29354
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d6194c1fd1bc7ba2c89f2e0eaa9c08228088f4fc
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Sep 16 13:56:39 2016 -0700

    cmd/compile: change goarch.Main into goarch.Init
    
    Similar idea to golang.org/cl/28970.
    
    Change-Id: I9d2feb1a669d71ffda1d612cf39ee0d3c08d22d2
    Reviewed-on: https://go-review.googlesource.com/29357
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 32db3f2756324616b7c856ac9501deccc2491239
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Sep 15 17:40:26 2016 -0700

    cmd/compile/internal/syntax: support for alias declarations
    
    Permits parsing of alias declarations with -newparser
    
            const/type/var/func T => p.T
    
    but the compiler will reject it with an error. For now this
    also accepts
    
            type T = p.T
    
    so we can experiment with a type-alias only scenario.
    
    - renamed _Arrow token to _Larrow (<-)
    - introduced _Rarrow token (=>)
    - introduced AliasDecl node
    - extended scanner to accept _Rarrow
    - extended parser and printer to handle alias declarations
    
    Change-Id: I0170d10a87df8255db9186d466b6fd405228c38e
    Reviewed-on: https://go-review.googlesource.com/29355
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 28ed2b0cd9ad6e2015f073931c05c08e8bf7b247
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 9 12:04:22 2016 -0400

    cmd/link: skip arch-specific main function
    
    Add some notes to main.go on what happens where.
    
    Change-Id: I4fb0b6c280e5f990ddc5d749267372b86870af6d
    Reviewed-on: https://go-review.googlesource.com/28970
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c199c76cb47f065f9f513bc6f74f58be0484765d
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 16 12:11:33 2016 -0700

    cmd/compile: turn live variable test off for ppc
    
    ppc64 has an extraneous variable live in some situations.
    We need a better tighten pass to get rid of this extra variable.
    I'm working on it, but fix the test in the meantime.
    
    Fixes build for ppc64.
    
    Change-Id: I1efb9ccb234a64f2a1c228abd2b3195f67fbeb41
    Reviewed-on: https://go-review.googlesource.com/29353
    Reviewed-by: David Chase <drchase@google.com>

commit 9f447c20efeaab1e60e1da03feee2fa0935cf95c
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Sep 16 11:26:41 2016 -0700

    cmd/compile: remove unused Label fields
    
    Updates #16357.
    
    Change-Id: I37f04d83134b5e1e7f6ba44eb9a566758ef594d3
    Reviewed-on: https://go-review.googlesource.com/29350
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 833ed7c4315d58f26875f512b5b5ef3b042ce7ed
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 16 09:36:00 2016 -0700

    cmd/compile: reorganize SSA register numbering
    
    Teach SSA about the cmd/internal/obj/$ARCH register numbering.
    It can then return that numbering when requested.  Each architecture
    now does not need to know anything about the internal SSA numbering
    of registers.
    
    Change-Id: I34472a2736227c15482e60994eebcdd2723fa52d
    Reviewed-on: https://go-review.googlesource.com/29249
    Reviewed-by: David Chase <drchase@google.com>

commit b87d7a5cf6da24b8dbe2f43cb69f61bd4dbc93f5
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 16 11:35:53 2016 -0400

    cmd/link: give RelocVariant a type, document Reloc
    
    Change-Id: Ib20263405a08674b5e160295fc965da4c8b54b34
    Reviewed-on: https://go-review.googlesource.com/29248
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6007c8c76beb6e9d8bccc966f0a1cc0f7518c539
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 9 17:34:07 2016 -0400

    cmd/link: attempt to rationalize linkmode init
    
    This CL gives Linkmode a type, switches it to the standard flag
    handling mechanism, and deduplicates some logic.
    
    There is a semantic change in this CL. Previously if a link was
    invoked explicitly with -linkmode=internal, any condition that forced
    external linking would silently override this and use external
    linking. Instead it now fails with a reason why. I believe this is an
    improvement, but will change it back if there's disagreement.
    
    Fixes #12848
    
    Change-Id: Ic80e341fff65ecfdd2b6fdd6079674cc7210fc5f
    Reviewed-on: https://go-review.googlesource.com/28971
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 1df438f79c440ddf9bdd342f089a55567254bc9a
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Aug 26 09:04:27 2016 -0400

    misc/cgo/testplugin: add test of -buildmode=plugin
    
    Change-Id: Ie9fea9814c850b084562ab2349b54d9ad9fa1f4a
    Reviewed-on: https://go-review.googlesource.com/27825
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c19382319abd444592a02e819db87fe77d9a888c
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Aug 26 08:51:58 2016 -0400

    cmd/go: support -buildmode=plugin on linux
    
    Change-Id: I0c8a04457db28c55c35c9a186b63c40f40730e39
    Reviewed-on: https://go-review.googlesource.com/27824
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 0cbb12f0bbaeb3893b3d011fdb1a270291747ab0
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Aug 26 08:50:50 2016 -0400

    plugin: new package for loading plugins
    
    Includes a linux implementation.
    
    Change-Id: Iacc2ed7da760ae9deebc928adf2b334b043b07ec
    Reviewed-on: https://go-review.googlesource.com/27823
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c87528d7761e3f133fe238937019a18bf7cf0d3c
Author: David Chase <drchase@google.com>
Date:   Thu Sep 15 15:14:32 2016 -0700

    cmd/compile: elide unnecessary sign/zeroExt, PPC64
    
    Bias {Eq,Neq}{8,16} to prefer the extension likely to match
    their operand's load (if loaded), and elide sign and zero
    extending MOV{B,W}, MOV{B,W}Z when their operands are already
    appropriately extended.
    
    Change-Id: Ic01b9cab55e170f68fc2369688b50ce78a818608
    Reviewed-on: https://go-review.googlesource.com/29236
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 46ba59025f527b2cfc5ef0d5ec47be45971ba672
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jul 20 10:09:40 2016 -0400

    cmd/compile: label LoadReg with line number of the use
    
    A tentative fix of #16380. It adds "line" everywhere...
    
    This also reduces binary size slightly (cmd/go on ARM as an example):
    
                            before          after
    total binary size       8068097         8018945 (-0.6%)
    .gopclntab              1195341         1179929 (-1.3%)
    .debug_line              689692          652017 (-5.5%)
    
    Change-Id: Ibda657c6999783c5bac180cbbba487006dbf0ed7
    Reviewed-on: https://go-review.googlesource.com/25082
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 882dd1c3baff19cfe8d59bc1f69d766b38540fa5
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Sep 16 09:27:06 2016 -0400

    cmd/compile: enable rewritings that are blocked by old backends
    
    Old backends did not implement them, but SSA do.
    
    Change-Id: I543b2281dcf4bab0da37c9b1f26a5ef55a0ea11b
    Reviewed-on: https://go-review.googlesource.com/29278
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 8607bed7445100993938ee96a028627461fce9d3
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Aug 25 22:04:04 2016 -0400

    runtime: avoid dependence on main symbol
    
    For -buildmode=plugin, this lets the linker drop the main.main symbol
    out of the binary while including most of the runtime.
    
    (In the future it should be possible to drop the entire runtime
    package from plugins.)
    
    Change-Id: I3e7a024ddf5cc945e3d8b84bf37a0b7cb2a00eb6
    Reviewed-on: https://go-review.googlesource.com/27821
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit eced6754c2f2ce98cb5bacbdbfcbbaa4a6a69d53
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Aug 25 21:58:45 2016 -0400

    cmd/link: -buildmode=plugin support for linux
    
    This CL contains several linker changes to support creating plugins.
    
    It collects the exported plugin symbols provided by the compiler and
    includes them in the moduledata.
    
    It treats a binary as being dynamically linked if it imports the plugin
    package. This lets the dynamic linker de-duplicate symbols.
    
    Change-Id: I099b6f38dda26306eba5c41dbe7862f5a5918d95
    Reviewed-on: https://go-review.googlesource.com/27820
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 79167bbd9c71017e18836a3e5e40261b57012358
Author: Suyash <dextrous93@gmail.com>
Date:   Wed Sep 14 22:36:58 2016 +0530

    sort: add examples for sort.Search
    
    This adds examples showing the different ways of using sort.Search.
    
    Change-Id: Iaa08b4501691f37908317fdcf2e618fbe9f99ade
    Reviewed-on: https://go-review.googlesource.com/29131
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9658de32c6141ca036a741006da5b83b4b2c6af5
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 19:48:53 2016 -0700

    cmd/compile: remove unused cases in Naddr
    
    Gins, and in turn Naddr, is only used with ONAME and OLITERAL Nodes,
    so we can drastically simplify Naddr.
    
    Passes toolstash/buildall.
    
    Change-Id: I2deb7eb771fd55e7c7f00040a9aee54588fcac11
    Reviewed-on: https://go-review.googlesource.com/29247
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c55c33af52c5ed97c93ec67bf7373d095bcb957d
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Wed Sep 14 02:11:02 2016 -0700

    os/exec: add examples for CombinedOutput, StdinPipe, StderrPipe
    
    Updates #16360.
    
    Adds examples for:
    + CombinedOutput
    + StdinPipe
    + StderrPipe
    
    Change-Id: I19293e64b34ed9268da00e0519173a73bfbc2c10
    Reviewed-on: https://go-review.googlesource.com/29150
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 1bd91d4ccc57d3dbb2e5452c16ff6281d53e9763
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 19:00:05 2016 -0700

    cmd/internal/obj: remove Addr.Etype and Addr.Width
    
    Since the legacy backends were removed, these fields are write-only.
    
    Change-Id: I4816c39267b7c10a4da2a6d22cd367dc475e564d
    Reviewed-on: https://go-review.googlesource.com/29246
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit 073d248bf50b1b1029b2a09c575111f38f6bf5eb
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Sep 16 11:00:54 2016 +1000

    cmd/compile/internal/gc: make Nod private
    
    Follow up to CL 29134. Generated with gofmt -r 'Nod -> nod', plus
    three manual adjustments to the comments in syntax/parser.go
    
    Change-Id: I02920f7ab10c70b6e850457b42d5fe35f1f3821a
    Reviewed-on: https://go-review.googlesource.com/29136
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit bb12894d2b51f8a50c0783db6043247758706466
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Sep 15 19:52:40 2016 -0400

    cmd/link: fix number-of-files entry in gopclntab
    
    According to golang.org/s/go12symtab, for N files, it should put N+1
    there.
    
    Fixes #17132.
    
    Change-Id: I0c84136855c6436be72b9d3c407bf10d4c81a099
    Reviewed-on: https://go-review.googlesource.com/29275
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 544010a05a90e45763ec9a8c149fc5137e1ec461
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 16:37:17 2016 -0700

    cmd/compile: remove Betypeinit
    
    Change-Id: I5c2fd0ff1b49f3826f2b9869b0b11329804b0e2a
    Reviewed-on: https://go-review.googlesource.com/29244
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit 1f2930cb5cf78cab648e39b9a0cee2f7a1e98d99
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 16:33:11 2016 -0700

    cmd/compile: remove traces of old register allocator
    
    Only added lines are moving amd64 and x86's ginsnop functions from
    gsubr.go to ggen.go to match other architectures, so all of the
    gsubr.go files can go away.
    
    Change-Id: Ib2292460c155ae6d9dcf5c9801f178031d8eea7a
    Reviewed-on: https://go-review.googlesource.com/29240
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit ae7e0ad7b8cdc16b6a7d5043375bb92a9cf6190a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 16:21:56 2016 -0700

    cmd/compile: remove more dead code
    
    Change-Id: I0131b0d7421ff1397f16a08eff758250abbdf8e2
    Reviewed-on: https://go-review.googlesource.com/29239
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit e888b81a11df98d352580358e1ba524d05904d69
Author: Vitor De Mario <vitordemario@gmail.com>
Date:   Thu Sep 15 20:46:20 2016 -0300

    doc: change variable name in Effective Go
    
    Effective Go has references to a function call f(c, req) made by ServeHTTP mixed with f(w,
    req). c is dropped in favor of w to maintain consistency
    
    Fixes #17128
    
    Change-Id: I6746fd115ed5a58971fd24e54024d29d18ead1fa
    Reviewed-on: https://go-review.googlesource.com/29311
    Reviewed-by: Rob Pike <r@golang.org>

commit e727e37090de987f2fafd48a02cd39455dff2ca5
Author: Alberto Bertogli <albertito@blitiri.com.ar>
Date:   Thu Sep 15 02:32:44 2016 +0100

    net: document dummy byte in ReadMsgUnix and WriteMsgUnix
    
    ReadMsgUnix and WriteMsgUnix both will read/write 1 byte from/to the
    socket if they were given no buffer to read/write, to avoid a common
    pitfall in out of band operations (they will usually block
    indefinitely if there's no actual data to read).
    
    This patch adds a note about this behaviour in their documentation, so
    users can be aware of it.
    
    Change-Id: I751f0e12bb4d80311e94ea8de023595c5d40ec3e
    Reviewed-on: https://go-review.googlesource.com/29180
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 897c0ebf61daaaf9c02be577f1afcb2c67ff3ed5
Author: David Chase <drchase@google.com>
Date:   Thu Sep 15 16:51:35 2016 -0400

    cmd/compile: adapt GOSSAHASH to set a DebugTest flag in ssa.Config
    
    Binary search remains our friend.
    Suppose you add an ought-to-be-benign pattern to PPC64.rules,
    and make.bash starts crashing.  You can guard the pattern(s)
    with config.DebugTest:
    
    (Eq8 x y) && config.DebugTest && isSigned(x.Type) &&
       isSigned(y.Type) ->
       (Equal (CMPW (SignExt8to32 x) (SignExt8to32 y)))
    
    and then
    
      gossahash -s ./make.bash
      ...
      (go drink beer while silicon minions toil)
      ...
      Trying ./make.bash args=[], env=[GOSSAHASH=100110010111110]
      ./make.bash failed (1 distinct triggers): exit status 1
      Trigger string is 'GOSSAHASH triggered (*importReader).readByte',
        repeated 1 times
      Review GSHS_LAST_FAIL.0.log for failing run
      Finished with GOSSAHASH=100110010111110
    
    Change-Id: I4eff46ebaf496baa2acedd32e217005cb3ac1c62
    Reviewed-on: https://go-review.googlesource.com/29273
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit b265d51789ea7ccf68787ebb30a2cdf03cc9d4fe
Author: Keith Randall <khr@golang.org>
Date:   Thu Sep 15 13:29:17 2016 -0700

    test,cmd/compile: remove _ssa file suffix
    
    Everything is SSA now.
    
    Update #16357
    
    Change-Id: I436dbe367b863ee81a3695a7d653ba4bfc5b0f6c
    Reviewed-on: https://go-review.googlesource.com/29232
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f13701bf2f32fe48f6dcf9149b35b9ce3effe022
Author: Keith Randall <khr@golang.org>
Date:   Thu Sep 15 13:04:50 2016 -0700

    test: make SSA tests unconditional
    
    Delete legacy backend tests, make SSA tests unconditional.
    
    Next CL will remove _ssa from the file names.
    
    Update #16357
    
    Change-Id: I2a7f5dcbc69455f63b5e6e6b2725df26ab86c8dd
    Reviewed-on: https://go-review.googlesource.com/29231
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 22d3bf1da49a6fa168cf16a619d3591100dba426
Author: Sam Whited <sam@samwhited.com>
Date:   Fri Sep 9 09:49:47 2016 -0500

    cmd/fix: add golang.org/x/net/context fix
    
    Fixes #17040
    
    Change-Id: I3682cc0367b919084c280d7dc64746495c1d4aaa
    Reviewed-on: https://go-review.googlesource.com/28872
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4466298df47fee61c31933e6fd5dde474dcfd2ad
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Wed Sep 7 09:36:58 2016 +0530

    cmd/go: add distribution-specific info for Linux to bug command
    
    Also remove the hard-coded path for getting glibc information.
    
    As an example, the following is the diff for `go bug` on Ubuntu before
    and after the change:
    
    >>>
    --- /tmp/01     2016-09-13 15:08:53.202758043 +0530
    +++ /tmp/02     2016-09-13 21:38:17.485039867 +0530
    @@ -1,7 +1,7 @@
     Please check whether the issue also reproduces on the latest release, go1.7.1.
    
     ```
    -go version devel +bdb3b79 Wed Sep 7 03:23:44 2016 +0000 linux/amd64
    +go version devel +cb13150 Wed Sep 7 09:46:58 2016 +0530 linux/amd64
     GOARCH="amd64"
     GOBIN=""
     GOEXE=""
    @@ -18,5 +18,23 @@
     CXX="g++"
     CGO_ENABLED="1"
     uname -sr: Linux 4.4.0-36-generic
    +Distributor ID:        Ubuntu
    +Description:   Ubuntu 16.04.1 LTS
    +Release:       16.04
    +Codename:      xenial
    +/lib/x86_64-linux-gnu/libc.so.6: GNU C Library (Ubuntu GLIBC 2.23-0ubuntu3) stable release version 2.23, by Roland McGrath et al.
     gdb --version: GNU gdb (Ubuntu 7.11.1-0ubuntu1~16.04) 7.11.1
     ```
    <<<
    
    Change-Id: I7e3730a797af0f94d6e43fe4743ab48bc0f11f1b
    Reviewed-on: https://go-review.googlesource.com/28581
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3eb16c059d301bc7242461ac8480d95e433e0374
Author: David Chase <drchase@google.com>
Date:   Thu Sep 15 12:43:29 2016 -0400

    cmd/compile: repair GOSSAFUNC functionality
    
    GOSSAFUNC=foo had previously only done printing for the
    single function foo, and didn't quite clean up after itself
    properly. Changes ensures that Config.HTML != nil iff
    GOSSAFUNC==name-of-current-function.
    
    Change-Id: I255e2902dfc64f715d93225f0d29d9525c06f764
    Reviewed-on: https://go-review.googlesource.com/29250
    Run-TryBot: David Chase <drchase@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ad7732a9acf6447e4b70845085263772ea788a6c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Sep 15 08:07:54 2016 -0700

    cmd/compile: remove gins
    
    The only remaining consumers of gins were
    ginsnop and arch-independent opcodes like GVARDEF.
    Rewrite ginsnop to create and populate progs directly.
    Move arch-independent opcodes to package gc
    and simplify.
    Delete some now unused code.
    There is more.
    Step one towards eliminating gc.Node.Reg.
    
    Change-Id: I7c34cd8a848f6fc3b030705ab8e293838e0b6c20
    Reviewed-on: https://go-review.googlesource.com/29220
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a06e931abec6082e8f2db65f29dd6c63be5c0de6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Sep 15 08:46:08 2016 -0700

    cmd/compile: nodintconst is the new Nodintconst
    
    Fixes the build.
    
    Change-Id: Ib9aca6cf86d595d20f22dbf730afa8dea4b44672
    Reviewed-on: https://go-review.googlesource.com/29221
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit c9fd997524ce7d531579500218f11b528bab4c88
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jun 25 13:40:43 2016 -0700

    cmd/compile: unroll comparisons to short constant strings
    
    Unroll s == "ab" to
    
    len(s) == 2 && s[0] == 'a' && s[1] == 'b'
    
    This generates faster and shorter code
    by avoiding a runtime call.
    Do something similar for !=.
    
    The cutoff length is 6. This was chosen empirically
    by examining binary sizes on arm, arm64, 386, and amd64
    using the SSA backend.
    
    For all architectures examined, 4, 5, and 6 were
    the ideal cutoff, with identical binary sizes.
    
    The distribution of constant string equality sizes
    during 'go build -a std' is:
    
     40.81%   622 len 0
     14.11%   215 len 4
      9.45%   144 len 1
      7.81%   119 len 3
      7.48%   114 len 5
      5.12%    78 len 7
      4.13%    63 len 2
      3.54%    54 len 8
      2.69%    41 len 6
      1.18%    18 len 10
      0.85%    13 len 9
      0.66%    10 len 14
      0.59%     9 len 17
      0.46%     7 len 11
      0.26%     4 len 12
      0.20%     3 len 19
      0.13%     2 len 13
      0.13%     2 len 15
      0.13%     2 len 16
      0.07%     1 len 20
      0.07%     1 len 23
      0.07%     1 len 33
      0.07%     1 len 36
    
    A cutoff of length 6 covers most of the cases.
    
    Benchmarks on amd64 comparing a string to a constant of length 3:
    
    Cmp/1same-8           4.78ns ± 6%  0.94ns ± 9%  -80.26%  (p=0.000 n=20+20)
    Cmp/1diffbytes-8      6.43ns ± 6%  0.96ns ±11%  -85.13%  (p=0.000 n=20+20)
    Cmp/3same-8           4.71ns ± 5%  1.28ns ± 5%  -72.90%  (p=0.000 n=20+20)
    Cmp/3difffirstbyte-8  6.33ns ± 7%  1.27ns ± 7%  -79.90%  (p=0.000 n=20+20)
    Cmp/3difflastbyte-8   6.34ns ± 8%  1.26ns ± 9%  -80.13%  (p=0.000 n=20+20)
    
    The change to the prove test preserves the
    existing intent of the test. When the string was
    short, there was a new "proved in bounds" report
    that referred to individual byte comparisons.
    
    Change-Id: I593ac303b0d11f275672090c5c786ea0c6b8da13
    Reviewed-on: https://go-review.googlesource.com/26758
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 584978f4b5a6abc6d3950114994e7281b525a232
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Sep 16 00:33:29 2016 +1000

    cmd/compile/internal/gc: unexport private variables
    
    Change-Id: I14a7c08105e6bdcee04a5cc21d7932e9ca753384
    Reviewed-on: https://go-review.googlesource.com/29138
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8e922759b381061cc4709c3e52d522d983521c5a
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Sep 16 00:17:33 2016 +1000

    cmd/compile/internal/gc: fix build
    
    Fix conflict between CL 29213 and 29134.
    
    Change-Id: Ie58bd7195893d7e634f1b257ee0bdd3250cd23c2
    Reviewed-on: https://go-review.googlesource.com/29137
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit df2b63f09be280544e58c8d17fe34c863521ed10
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Sep 14 20:07:13 2016 -0700

    cmd/compile: unwrap fewer CONVNOPs in staticassign
    
    staticassign unwraps all CONVNOPs.
    However, in the included test, we need the
    CONVNOP for everything to typecheck.
    Stop unwrapping unnecessarily.
    
    The code we generate for this example is
    suboptimal, but that's not new; see #17113.
    
    Fixes #17111.
    
    Change-Id: I29532787a074a6fe19a5cc53271eb9c84bf1b576
    Reviewed-on: https://go-review.googlesource.com/29213
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d7012ca282092e876c3f12b758e12d6eabc8cc3c
Author: Dave Cheney <dave@cheney.net>
Date:   Thu Sep 15 15:45:10 2016 +1000

    cmd/compile/internal/gc: unexport more helper functions
    
    After the removal of the old backend many types are no longer referenced
    outside internal/gc. Make these functions private so that tools like
    honnef.co/go/unused can spot when they become dead code. In doing so
    this CL identified several previously public helpers which are no longer
    used, so removes them.
    
    This should be the last of the public functions.
    
    Change-Id: I7e9c4e72f86f391b428b9dddb6f0d516529706c3
    Reviewed-on: https://go-review.googlesource.com/29134
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 896ac677b5e3e80278cc1ce179d8a077ac3a6d2f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Sep 14 14:51:41 2016 -0700

    cmd/go: make bug subcommand open the browser
    
    Instead of dumping information for the use
    to copy/paste into the issue tracker,
    open the issue tracker directly with a pre-filled
    template.
    
    Change-Id: I370d0063b609200497014ccda35244fa4314a662
    Reviewed-on: https://go-review.googlesource.com/29210
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 27eebbabc23366552508f4b13375f8729b90aacb
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 00:52:36 2016 -0700

    cmd/compile, runtime: remove throwreturn
    
    Change-Id: If8d27cf1cd8d650ed0ba332448d3174d80b6b0ca
    Reviewed-on: https://go-review.googlesource.com/29217
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 78c46581f4da82473f3efcbbe1f3b48625f88471
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 00:34:35 2016 -0700

    cmd/compile: deduplicate appendpp functions
    
    Change-Id: Ifa7a882b020f7b0c9602c28c3e5567e5d4c39e73
    Reviewed-on: https://go-review.googlesource.com/29216
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 8a0bd5dc01f21888525e58d42efdc9dd246748d6
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 15 00:21:28 2016 -0700

    cmd/compile: eliminate differences in appendpp functions
    
    Consistently use int16 for [ft]reg and int64 for [ft]offset.
    
    Change-Id: I7d279bb6e4fb735105429234a949074bf1cefb29
    Reviewed-on: https://go-review.googlesource.com/29215
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 838eaa738f2bc07c3706b96f9e702cb80877dfe1
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Sep 14 23:57:50 2016 -0700

    cmd/compile: remove more dead code
    
    Change-Id: Ib05a8e149db8accdb1474703cd7b7004243d91d4
    Reviewed-on: https://go-review.googlesource.com/29214
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit f03855f40ef131de6c0881ec12996a747be05a83
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Sep 14 16:49:44 2016 -0700

    cmd/compile: remove ginscon2 functions
    
    These are no longer reachable as gins dispatches to ginscon for all
    arch-specific instructions anyway.
    
    Change-Id: I7f34883c16058308d8afa0f960dcf554af31bfe4
    Reviewed-on: https://go-review.googlesource.com/29211
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 82703f84e4e2ab5b68d43825e0f046a5ad990612
Author: Dave Cheney <dave@cheney.net>
Date:   Thu Sep 15 14:34:20 2016 +1000

    cmd/compile/internal/gc: unexport helper functions
    
    After the removal of the old backend many types are no longer referenced
    outside internal/gc. Make these functions private so that tools like
    honnef.co/go/unused can spot when they become dead code. In doing so
    this CL identified several previously public helpers which are no longer
    used, so removes them.
    
    Change-Id: Idc2d485f493206de9d661bd3cb0ecb4684177b32
    Reviewed-on: https://go-review.googlesource.com/29133
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 24965bc96ce2fd1483b6b8d5759b0ea129688758
Author: Dave Cheney <dave@cheney.net>
Date:   Thu Sep 15 14:20:55 2016 +1000

    cmd/compile/internal/gc: remove dead code
    
    Remove unused functions spotted by honnef.co/go/unused.
    
    Change-Id: Iabf3b201215ce21e420a60f4ef2679b36231eda7
    Reviewed-on: https://go-review.googlesource.com/29132
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e7cc9c4f4dee2d6f8a77ce177172ef4ae4ff0c90
Author: Keith Randall <khr@golang.org>
Date:   Wed Sep 14 14:02:50 2016 -0700

    cmd/compile: delete lots of the legacy backend
    
    It's not everything, but it is a good start.
    
    I tried to make the CL delete only.  goimports forced
    a few exceptions to that rule.
    
    Update #16357
    
    Change-Id: I041925cb2fe68bb7ae1617af862b22c48da649c1
    Reviewed-on: https://go-review.googlesource.com/29168
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Martin Möhrmann <martisch@uos.de>

commit 3134ab3c2d58d2c9f93471c85405cb26ccdf0f4d
Author: Keith Randall <khr@golang.org>
Date:   Tue Sep 13 17:01:01 2016 -0700

    cmd/compile: redo nil checks
    
    Get rid of BlockCheck. Josh goaded me into it, and I went
    down a rabbithole making it happen.
    
    NilCheck now panics if the pointer is nil and returns void, as before.
    BlockCheck is gone, and NilCheck is no longer a Control value for
    any block. It just exists (and deadcode knows not to throw it away).
    
    I rewrote the nilcheckelim pass to handle this case.  In particular,
    there can now be multiple NilCheck ops per block.
    
    I moved all of the arch-dependent nil check elimination done as
    part of ssaGenValue into its own proper pass, so we don't have to
    duplicate that code for every architecture.
    
    Making the arch-dependent nil check its own pass means I needed
    to add a bunch of flags to the opcode table so I could write
    the code without arch-dependent ops everywhere.
    
    Change-Id: I419f891ac9b0de313033ff09115c374163416a9f
    Reviewed-on: https://go-review.googlesource.com/29120
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit f9e9412ce26ca208c1d25a6b854259c94aa54030
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Sep 10 18:53:30 2016 -0700

    cmd/dist: run vet/all on dedicated builders
    
    We will soon add dedicated builders for running vet/all.
    Their name will end with "-vetall".
    On those builders, run vet/all and nothing else.
    On all other builders, including local all.bash,
    don't run vet/all at all, because it is slow.
    
    This will probably be refined more over time.
    
    Change-Id: Ib1d0337adda84353931a325515c132068d4320cd
    Reviewed-on: https://go-review.googlesource.com/28962
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fb273fc3a3438f7a24b0901e2fcfd099eac4860d
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Aug 26 10:50:12 2016 -0700

    cmd/compile: fix comma-ok assignments for non-boolean ok
    
    Passes toolstash -cmp.
    
    Fixes #16870.
    
    Change-Id: I70dc3bbb3cd3031826e5a54b96ba1ea603c282d1
    Reviewed-on: https://go-review.googlesource.com/27910
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 150de948eeceac0ace02a0f93a9a7a1f7421d744
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sat Sep 10 22:44:00 2016 +0200

    cmd/compile: intrinsify slicebytetostringtmp when not instrumenting
    
    when not instrumenting:
    - Intrinsify uses of slicebytetostringtmp within the runtime package
      in the ssa backend.
    - Pass OARRAYBYTESTRTMP nodes to the compiler backends for lowering
      instead of generating calls to slicebytetostringtmp.
    
    name                    old time/op  new time/op  delta
    ConcatStringAndBytes-4  27.9ns ± 2%  24.7ns ± 2%  -11.52%  (p=0.000 n=43+43)
    
    Fixes #17044
    
    Change-Id: I51ce9c3b93284ce526edd0234f094e98580faf2d
    Reviewed-on: https://go-review.googlesource.com/29017
    Run-TryBot: Martin Möhrmann <martisch@uos.de>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 1c5ac0827d2d0d2f5fb3b7f2b34b37e170beff1d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Sep 14 14:10:14 2016 -0700

    cmd/vet/all: check platforms concurrently
    
    Change-Id: I63e7fd7f62aa80e1252b0c5b6c472439aa66da73
    Reviewed-on: https://go-review.googlesource.com/29169
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b78108d5dc198a1bb8ed294b8b974f0c0d51b55c
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sat Sep 3 18:16:52 2016 -0700

    time: document in UnixNano when the value is valid
    
    It is unlikely that the value of UnixNano overflow in most
    use cases. However, the max date of 2262 is also within the range
    where it may be of concern to some users. Rather than have each
    person recompute when this overflows to validate if its okay for
    their use case, we just document it as within the years 1678 and
    2262, for user's convenience.
    
    Fixes #16977
    
    Change-Id: I4988738c147f4a6d30f8b8735c3941b75113bb16
    Reviewed-on: https://go-review.googlesource.com/28478
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 4b8a1611b5d37fd41d2d11ef1bb3455c77b34e07
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Sep 14 13:19:20 2016 -0700

    cmd/compile: add Nodes.Prepend helper method
    
    Prepared with gofmt -r.
    
    Change-Id: Ib9f224cc20353acd9c5850dead1a2d32ca5427d3
    Reviewed-on: https://go-review.googlesource.com/29165
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6f135bfd922cafd000497467c73cc03cfa788fa6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Sep 14 18:44:59 2016 +0000

    math/big: cut 2 minutes off race tests
    
    No need to test so many sizes in race mode, especially for a package
    which doesn't use goroutines.
    
    Reduces test time from 2.5 minutes to 25 seconds.
    
    Updates #17104
    
    Change-Id: I7065b39273f82edece385c0d67b3f2d83d4934b8
    Reviewed-on: https://go-review.googlesource.com/29163
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 983e2fd4e6fecd78789384069c476eead4806d1b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Sep 14 18:14:34 2016 +0000

    sort: cut 140 seconds off race build tests
    
    No coverage is gained by running the 1e6 versions of the test over the
    1e4 versions. It just adds 140 seconds of race overhead time.
    
    Updates #17104
    
    Change-Id: I41408aedae34a8b1a148eebdda20269cdefffba3
    Reviewed-on: https://go-review.googlesource.com/29159
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 83676b93fb9591fb452612b55237ffdeb2a59119
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Sep 14 11:35:29 2016 -0700

    cmd/compile, cmd/link: fix printf verbs
    
    Found by vet.
    
    Change-Id: I9dbc6208ddbb5b407f4ddd20efbc166aac852cf7
    Reviewed-on: https://go-review.googlesource.com/29162
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d72c2846f1fac679d6def717b4e9c83fbafdf908
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Sep 14 11:35:11 2016 -0700

    cmd/vet/all: update whitelist
    
    CL 29110 brought the fix into the main tree.
    
    Change-Id: I7bf02670d40f22d35c63e05173419fdee9f93462
    Reviewed-on: https://go-review.googlesource.com/29161
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b92d39ef6924fd5174449f95505d782f3f75db16
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 7 16:50:46 2016 -0700

    cmd/compile/internal/obj/x86: eliminate some function prologues
    
    The standard sort swap method
    
    func (t T) Swap(i, j int) {
      t[i], t[j] = t[j], t[i]
    }
    
    uses no stack space on architectures for which
    FixedFrameSize == 0, currently 386 and amd64.
    
    Nevertheless, we insert a stack check prologue.
    This is because it contains a call to
    runtime.panicindex.
    
    However, for a few common runtime functions,
    we know at compile time that they require
    no arguments. Allow them to pass unnoticed.
    
    Triggers for 380 functions during make.bash.
    Cuts 4k off cmd/go.
    
    encoding/binary benchmarks:
    
    ReadSlice1000Int32s-8     9.49µs ± 3%    9.41µs ± 5%    ~     (p=0.075 n=29+27)
    ReadStruct-8              1.50µs ± 3%    1.48µs ± 2%  -1.49%  (p=0.000 n=30+28)
    ReadInts-8                 599ns ± 3%     600ns ± 3%    ~     (p=0.471 n=30+29)
    WriteInts-8                836ns ± 4%     841ns ± 3%    ~     (p=0.371 n=30+29)
    WriteSlice1000Int32s-8    8.84µs ± 3%    8.69µs ± 5%  -1.71%  (p=0.001 n=30+30)
    PutUvarint32-8            29.6ns ± 1%    28.1ns ± 3%  -5.21%  (p=0.000 n=28+28)
    PutUvarint64-8            82.6ns ± 5%    82.3ns ±10%  -0.43%  (p=0.014 n=27+30)
    
    Swap assembly before:
    
    "".T.Swap t=1 size=74 args=0x28 locals=0x0
            0x0000 00000 (swap.go:5)        TEXT    "".T.Swap(SB), $0-40
            0x0000 00000 (swap.go:5)        MOVQ    (TLS), CX
            0x0009 00009 (swap.go:5)        CMPQ    SP, 16(CX)
            0x000d 00013 (swap.go:5)        JLS     67
            0x000f 00015 (swap.go:5)        FUNCDATA        $0, gclocals·3cadd97b66f25a3a642be35e9362338f(SB)
            0x000f 00015 (swap.go:5)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
            0x000f 00015 (swap.go:5)        MOVQ    "".i+32(FP), AX
            0x0014 00020 (swap.go:5)        MOVQ    "".t+16(FP), CX
            0x0019 00025 (swap.go:5)        CMPQ    AX, CX
            0x001c 00028 (swap.go:5)        JCC     $0, 60
            0x001e 00030 (swap.go:5)        MOVQ    "".t+8(FP), DX
            0x0023 00035 (swap.go:5)        MOVBLZX (DX)(AX*1), BX
            0x0027 00039 (swap.go:5)        MOVQ    "".j+40(FP), SI
            0x002c 00044 (swap.go:5)        CMPQ    SI, CX
            0x002f 00047 (swap.go:5)        JCC     $0, 60
            0x0031 00049 (swap.go:5)        MOVBLZX (DX)(SI*1), CX
            0x0035 00053 (swap.go:5)        MOVB    CL, (DX)(AX*1)
            0x0038 00056 (swap.go:5)        MOVB    BL, (DX)(SI*1)
            0x003b 00059 (swap.go:5)        RET
            0x003c 00060 (swap.go:5)        PCDATA  $0, $1
            0x003c 00060 (swap.go:5)        CALL    runtime.panicindex(SB)
            0x0041 00065 (swap.go:5)        UNDEF
            0x0043 00067 (swap.go:5)        NOP
            0x0043 00067 (swap.go:5)        CALL    runtime.morestack_noctxt(SB)
            0x0048 00072 (swap.go:5)        JMP     0
    
    Swap assembly after:
    
    "".T.Swap t=1 size=52 args=0x28 locals=0x0
            0x0000 00000 (swap.go:5)        TEXT    "".T.Swap(SB), $0-40
            0x0000 00000 (swap.go:5)        FUNCDATA        $0, gclocals·3cadd97b66f25a3a642be35e9362338f(SB)
            0x0000 00000 (swap.go:5)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
            0x0000 00000 (swap.go:5)        MOVQ    "".i+32(FP), AX
            0x0005 00005 (swap.go:5)        MOVQ    "".t+16(FP), CX
            0x000a 00010 (swap.go:5)        CMPQ    AX, CX
            0x000d 00013 (swap.go:5)        JCC     $0, 45
            0x000f 00015 (swap.go:5)        MOVQ    "".t+8(FP), DX
            0x0014 00020 (swap.go:5)        MOVBLZX (DX)(AX*1), BX
            0x0018 00024 (swap.go:5)        MOVQ    "".j+40(FP), SI
            0x001d 00029 (swap.go:5)        CMPQ    SI, CX
            0x0020 00032 (swap.go:5)        JCC     $0, 45
            0x0022 00034 (swap.go:5)        MOVBLZX (DX)(SI*1), CX
            0x0026 00038 (swap.go:5)        MOVB    CL, (DX)(AX*1)
            0x0029 00041 (swap.go:5)        MOVB    BL, (DX)(SI*1)
            0x002c 00044 (swap.go:5)        RET
            0x002d 00045 (swap.go:5)        PCDATA  $0, $1
            0x002d 00045 (swap.go:5)        CALL    runtime.panicindex(SB)
            0x0032 00050 (swap.go:5)        UNDEF
    
    Change-Id: I57dad14af8aaa5e6112deac407cfadc2bfaf1f54
    Reviewed-on: https://go-review.googlesource.com/24814
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 33ed35647520f2162c2fed1b0e5f19cec2c65de3
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Sep 14 11:16:50 2016 -0700

    cmd: add internal/browser package
    
    cmd/cover, cmd/trace, and cmd/pprof all open browsers.
    'go bug' will soon also open a browser.
    It is time to unify the browser-handling code.
    
    Change-Id: Iee6b443e21e938aeaaac366a1aefb1afbc7d9b2c
    Reviewed-on: https://go-review.googlesource.com/29160
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9a7ce41d6c74ef30af361677d2077ad8dd0e92b7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Sep 14 18:03:26 2016 +0000

    bytes: cut 10 seconds off the race builder's benchmark test
    
    Don't benchmark so many sizes during the race builder's benchmark run.
    
    This package doesn't even use goroutines.
    
    Cuts off 10 seconds.
    
    Updates #17104
    
    Change-Id: Ibb2c7272c18b9014a775949c656a5b930f197cd4
    Reviewed-on: https://go-review.googlesource.com/29158
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit f09d0458d32a965602a346843354f7778b12a375
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Sep 14 17:48:33 2016 +0000

    regexp: don't run slow benchmarks on race builders
    
    Shave 6.5 minutes off the *-race build time.
    
    The *-race builders run:
    
        go test -short -race -run=^$ -benchtime=.1s -cpu=4 $PKG
    
    ... for each package with benchmarks.
    
    The point isn't to measure the speed of the packages, but rather to
    see if there are any races. (which is why a benchtime of 0.1 seconds
    is used)
    
    But running in race mode makes things slower and our benchmarks aren't
    all very fast to begin with.
    
    The regexp benchmarks in race were taking over 6.5 minutes. With this
    CL, it's now 8 seconds.
    
    Updates #17104
    
    Change-Id: I054528d09b1568d37aac9f9b515d6ed90a5cf5b0
    Reviewed-on: https://go-review.googlesource.com/29156
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 167e381f405d36f71ef152e45bb845b866592c80
Author: Keith Randall <khr@golang.org>
Date:   Wed Sep 14 10:01:05 2016 -0700

    cmd/compile: make ssa compilation unconditional
    
    Rip out the code that allows SSA to be used conditionally.
    
    No longer exists:
     ssa=0 flag
     GOSSAHASH
     GOSSAPKG
     SSATEST
    
    GOSSAFUNC now only controls the printing of the IR/html.
    
    Still need to rip out all of the old backend.  It should no longer be
    callable after this CL.
    
    Update #16357
    
    Change-Id: Ib30cc18fba6ca52232c41689ba610b0a94aa74f5
    Reviewed-on: https://go-review.googlesource.com/29155
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit eed061f89cad583c3dfa6b4e13c9a79f897920bf
Author: David Chase <drchase@google.com>
Date:   Tue Sep 13 16:51:42 2016 -0400

    cmd/compile: enable SSA for PowerPC 64 Big-endian
    
    It passed tests once, if anything's wrong, better to fail
    sooner than later.
    
    Change-Id: Ibb1c5db3f4c5535a4ff4681fd157db77082c5041
    Reviewed-on: https://go-review.googlesource.com/28982
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3ead49989ed037bda6009b96b0008ed83fb7d184
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Sep 14 16:34:27 2016 +0000

    cmd/dist: skip compiling 100 packages without benchmarks in race mode
    
    The go_test_bench:* tests run:
    
        go test -short -race -run=^$ -benchtime=.1s -cpu=4 $PKG
    
    ... on each discovered package with any tests. (The same set used for
    the "go_test:*" tests)
    
    That set was 168 packages:
    
    $ go tool dist test -list | grep go_test: | wc -l
    168
    
    But only 76 of those have a "func Benchmark", and running each
    "go_test_bench:" test and compiling it in race mode, just to do
    nothing took 1-2 seconds each.
    
    So stop doing that and filter out the useless packages earlier. Now:
    
    $ go tool dist test -list -race | grep go_test_bench:  | wc -l
    76
    
    Should save 90-180 seconds. (or maybe 45 seconds for trybots, since
    they're sharded)
    
    Updates #17104
    
    Change-Id: I08ccb072a0dc0454ea425540ee8e74b59f83b773
    Reviewed-on: https://go-review.googlesource.com/29153
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 059ada596ca5d510d1e1755b3b6dafa036195fb0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Sep 12 11:36:02 2016 -0700

    cmd/dist: skip broken cgo race tests on darwin
    
    CL 26668 exposed #17065.
    Skip the cgo race tests on darwin for now.
    
    Updates #17065
    
    Change-Id: I0ad0ce2ff1af6d515b8ce6184ddeabc49806950f
    Reviewed-on: https://go-review.googlesource.com/29077
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2e675142dd269348da992c1862a28e56e0a87eb6
Author: Jan Mercl <0xjnml@gmail.com>
Date:   Wed Sep 14 13:17:37 2016 +0200

    test/float_lit2: fix expressions in comment
    
    The change corrects the values of the largest float32 value (f1) and the
    value of the halfway point between f1 and the next, overflow value (f2).
    
    Fixes #17012
    
    Change-Id: Idaf9997b69d61fafbffdb980d751c9857732e14d
    Reviewed-on: https://go-review.googlesource.com/29171
    Reviewed-by: Robert Griesemer <gri@golang.org>
```
