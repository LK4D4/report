# May 5, 2016 Report

Number of commits: 137

## Compilation time

* github.com/coreos/etcd: from 11.145033182s to 11.052837388s, -0.83%
* github.com/boltdb/bolt/cmd/bolt: from 549.946287ms to 549.075444ms, -0.16%
* github.com/gogits/gogs: from 12.713240493s to 12.556416868s, -1.23%
* github.com/spf13/hugo: from 6.618978324s to 6.558986297s, -0.91%
* github.com/influxdata/influxdb/cmd/influxd: from 6.302705879s to 6.273564048s, -0.46%
* github.com/nsqio/nsq/apps/nsqd: from 2.128987593s to 2.117584667s, -0.54%
* github.com/mholt/caddy: from 4.70856066s to 4.678892572s, -0.63%

## Binary size:

* github.com/coreos/etcd: from 21114336 to 21161984, +0.23%
* github.com/boltdb/bolt/cmd/bolt: from 2552515 to 2565367, +0.50%
* github.com/gogits/gogs: from 22521922 to 22572472, +0.22%
* github.com/spf13/hugo: from 14634329 to 14667718, +0.23%
* github.com/influxdata/influxdb/cmd/influxd: from 15517860 to 15585756, +0.44%
* github.com/nsqio/nsq/apps/nsqd: from 9619978 to 9683310, +0.66%
* github.com/mholt/caddy: from 12977503 to 13041133, +0.49%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               189           185           -2.12%
BenchmarkMsgpUnmarshal-4             403           389           -3.47%
BenchmarkEasyJsonMarshal-4           1463          1425          -2.60%
BenchmarkEasyJsonUnmarshal-4         1524          1477          -3.08%
BenchmarkFlatBuffersMarshal-4        366           344           -6.01%
BenchmarkFlatBuffersUnmarshal-4      291           284           -2.41%
BenchmarkGogoprotobufMarshal-4       160           161           +0.63%
BenchmarkGogoprotobufUnmarshal-4     244           245           +0.41%

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

* [os/exec: add Cmd.RunContext and Cmd.WaitContext](https://github.com/golang/go/commit/2cc27a7de9e7d14cb6702153688d02746c6a49ea)
* [os/exec: fix variable shadow, don't leak goroutine](https://github.com/golang/go/commit/1b591dfb1f071d978448966e979e40b1f265c1a5)
* [net/http, net/http/httptrace: new package for tracing HTTP client requests](https://github.com/golang/go/commit/1518d431321100cd9f0e18d740da7c835ba438dd) 
* [compress/flate: replace "Best Speed" with specialized version](https://github.com/golang/go/commit/d8b7bd6a1f89df1fbcf43fcaee72a94e291dcdb1)
* [net/http: add Transport.IdleConnTimeout](https://github.com/golang/go/commit/abc1472d78c70888473634497b49b1c2e1bb6569)
* [net/http: provide access to the listener address an HTTP request arrived on](https://github.com/golang/go/commit/a9cf0b1e1e2a66db547fcabb7188465e4ac54700)

## GIT Log

```
commit bcd4b84bc56889b5a9a8a5d457f35fc0188e8315
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed May 4 22:07:50 2016 -0700

    runtime: skip TestCgoCallbackGC on linux/mips64x
    
    Builder is too slow. This test passed on builder machines but took
    15+ min.
    
    Change-Id: Ief9d67ea47671a57e954e402751043bc1ce09451
    Reviewed-on: https://go-review.googlesource.com/22798
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3c4ebd20235e965237f5856d30d6ce6513b9f6c9
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed May 4 21:38:45 2016 -0700

    cmd/link: fix external linker argument for mips64
    
    I overlooked it when rebasing CL 19803.
    
    Change-Id: Ife9d6bcc6a772715d137af903c64bafac0cdb216
    Reviewed-on: https://go-review.googlesource.com/22797
    Reviewed-by: Minux Ma <minux@golang.org>

commit 34f97d28d2ff435b8ac85ad6645aaf79a5d061bd
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed May 4 13:27:27 2016 -0700

    runtime: put tracebackctxt C functions in .c file
    
    Since tracebackctxt.go uses //export functions, the C functions can't be
    externally visible in the C comment. The code was using attributes to
    work around that, but that failed on Windows.
    
    Change-Id: If4449fd8209a8998b4f6855ea89e5db1471b2981
    Reviewed-on: https://go-review.googlesource.com/22786
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 57be1607d975ebec2f5faecea068f2db31abc041
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Tue May 3 16:37:33 2016 +1000

    debug/pe: unexport newly introduced identifiers
    
    CLs 22181, 22332 and 22336 intorduced new functionality to be used
    in cmd/link (see issue #15345 for details). But we didn't have chance
    to use new functionality yet. Unexport newly introduced identifiers,
    so we don't have to commit to the API until we actually tried it.
    
    Rename File.COFFSymbols into File._COFFSymbols,
    COFFSymbol.FullName into COFFSymbol._FullName,
    Section.Relocs into Section._Relocs,
    Reloc into _Relocs,
    File.StringTable into File._StringTable and
    StringTable into _StringTable.
    
    Updates #15345
    
    Change-Id: I770eeb61f855de85e0c175225d5d1c006869b9ec
    Reviewed-on: https://go-review.googlesource.com/22720
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 27d0d849fe7802794c74d049f5e8e1b0be018f9a
Author: Nigel Tao <nigeltao@golang.org>
Date:   Thu May 5 06:43:13 2016 +1000

    compress/flate: distinguish between base and min match length.
    
    Change-Id: I93db5cd86e3fb568e4444cad95268ba4a02ce8a0
    Reviewed-on: https://go-review.googlesource.com/22787
    Reviewed-by: Nigel Tao <nigeltao@golang.org>

commit 1f4f2d0d39ceac483f3d42eb25ec992cdaf257f3
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed May 4 14:26:46 2016 -0400

    cmd/compile: fix uint64 to float casts on ppc64
    
    Adds the FCFIDU instruction and uses it instead of the FCFID
    instruction for unsigned integer to float casts. This change means
    that unsigned integers do not have to be cast to signed integers
    before being cast to a floating point value. Therefore it is no
    longer necessary to insert instructions to detect and fix
    values that overflow int64.
    
    The previous code generating the uint64 to int64 cast handled
    overflow by truncating the uint64 value. This truncation can
    change the result of the rounding performed by the integer to
    float cast.
    
    The FCFIDU instruction was added in Power ISA 2.06B.
    
    Fixes #15539.
    
    Change-Id: Ia37a9631293eff91032d4cd9a9bec759d2142437
    Reviewed-on: https://go-review.googlesource.com/22772
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3747c00923b140847f5d8f7978e672fa79e3ec37
Author: Robert Griesemer <gri@golang.org>
Date:   Wed May 4 15:13:36 2016 -0700

    cmd/compile: fix debugFormat for new export format
    
    Change-Id: Ic3415f3ee643636eab4ff7d2351b8ad0dae62895
    Reviewed-on: https://go-review.googlesource.com/22792
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 4d6788ecae127017ab4ab4d5d5907b8912a08bd6
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Thu May 5 00:04:54 2016 +0530

    runtime: clean up profiling data files produced by TestCgoPprof
    
    Fixes #15541
    
    Change-Id: I9b6835157db0eb86de13591e785f971ffe754baa
    Reviewed-on: https://go-review.googlesource.com/22783
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bfcb5b64062b3bdf367e0888cd42a4a46d37a913
Author: Shenghou Ma <minux@golang.org>
Date:   Wed May 4 17:06:50 2016 -0400

    net/http: correct RFC for MethodPatch
    
    Fixes #15546.
    
    Change-Id: I39c29ea6999812dd5f1c45f67bddad28f20b6c6b
    Reviewed-on: https://go-review.googlesource.com/22773
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2e2481ed345347197d0a5685ef200ac761def2f2
Author: Shenghou Ma <minux@golang.org>
Date:   Thu Apr 7 16:03:53 2016 -0400

    cmd/dist: remove the use of debug/elf package
    
    debug/elf is only needed to determine the endianness of the host
    machine, which is easy to do without debug/elf.
    
    Fixes #15180.
    
    Change-Id: I21035ed3884871270765a1ca3b812a5d4890a7ee
    Reviewed-on: https://go-review.googlesource.com/21662
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 24744f6561f3ff7bc58046ba62abbc1c07e9fd4e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed May 4 11:43:42 2016 -0700

    cmd/compile: check that SSA memory args are in the right place
    
    Fixes #15510
    
    Change-Id: I2e0568778ef90cf29712753b8c42109ef84a0256
    Reviewed-on: https://go-review.googlesource.com/22784
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 0f84afe24bd7954ed2408008c17b1930881d89d5
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 4 09:08:27 2016 -0700

    cmd/dist: add flag to build tests only, but not run them
    
    To speed up the ssacheck check builder and make it on by default as a
    trybot.
    
    Change-Id: I91a3347491507c84f4878dff744ca426ba3e2e9f
    Reviewed-on: https://go-review.googlesource.com/22755
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 43d2a10e2647b2d6ee10a2a7629ee4055dcce740
Author: Shenghou Ma <minux@golang.org>
Date:   Wed Apr 27 13:47:41 2016 -0400

    runtime/internal/atomic: fix vet warnings
    
    Change-Id: Ib29cf7abbbdaed81e918e5e41bca4e9b8da24621
    Reviewed-on: https://go-review.googlesource.com/22503
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7a964e1df4c94941d7ca53491de4ff7f12bc0941
Author: Shenghou Ma <minux@golang.org>
Date:   Wed Apr 27 13:47:19 2016 -0400

    cmd/vet: add mips64x assembly check support
    
    Change-Id: I0a6a92604dbfa4b0f9c5ae483b574331f246dcad
    Reviewed-on: https://go-review.googlesource.com/22502
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit c5a95e47a02fa32a5a560d1e1cccc30f0a502dda
Author: Shenghou Ma <minux@golang.org>
Date:   Wed May 4 12:48:20 2016 -0400

    doc: update go1.7.txt
    
    Change-Id: I78efbfe0d7e9825126109a568c99f548a099b0fb
    Reviewed-on: https://go-review.googlesource.com/22771
    Reviewed-by: Minux Ma <minux@golang.org>

commit 04a30025dbc8ff7757a763562d7d4a2b5ff68583
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:58 2016 -0400

    test: enable fixedbugs/issue10607.go test on linux/mips64x
    
    external linking is now supported.
    
    Change-Id: I13e90c39dad86e60781adecdbe8e6bc9e522f740
    Reviewed-on: https://go-review.googlesource.com/19811
    Reviewed-by: Minux Ma <minux@golang.org>

commit 3b4180c67425ead0631a8c86eb164f9eaa6c922d
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:48 2016 -0400

    cmd/go: enable TestNodeReading on linux/mips64x
    
    external linking is now supported.
    
    Change-Id: I3f552f5f09391205fced509fe8a5a38297ea8153
    Reviewed-on: https://go-review.googlesource.com/19810
    Reviewed-by: Minux Ma <minux@golang.org>

commit 1cec0fec8caff1eacaba99ff71b2c57574887aab
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:44 2016 -0400

    cmd/dist: enable cgo and external linking on linux/mips64x
    
    Fixes #14126
    
    Change-Id: I21c8e06c01d3ef02ee09dc185d4443e2da8fd52b
    Reviewed-on: https://go-review.googlesource.com/19809
    Reviewed-by: Minux Ma <minux@golang.org>

commit 094e5a92886948ed2ba42031e29919ecb6087c8b
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:38 2016 -0400

    misc/cgo/test: add mips64x test case for issue9400
    
    Change-Id: If2b4abb6ff322c20e35de025298c8e5ab53edd42
    Reviewed-on: https://go-review.googlesource.com/19808
    Reviewed-by: Minux Ma <minux@golang.org>

commit dcd613862bc51db8455ccd22c03336111afe6883
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:33 2016 -0400

    cmd/link/internal/ld: force external linking for mips64x with cgo
    
    cgo internal linking is not supported yet (issue #14449).
    
    Change-Id: Ic968916383d77b7f449db8f230c928a1e81939e0
    Reviewed-on: https://go-review.googlesource.com/19807
    Reviewed-by: Minux Ma <minux@golang.org>

commit b6687c8933e2e123cd336d6ee96aa43df40bfc24
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:24 2016 -0400

    runtime: add linux/mips64x cgo support
    
    Change-Id: Id40dd05b7b264f3b779fdf9ccc2421ba4bc70589
    Reviewed-on: https://go-review.googlesource.com/19806
    Reviewed-by: Minux Ma <minux@golang.org>

commit 6e9043234221678e53d3b46ea98fb2ed56c260b7
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Tue May 3 13:49:54 2016 -0700

    runtime/cgo: add context argument to crosscall2 on mips64
    
    Change-Id: Id018516075842afd8af12fbf207763a851d5a851
    Reviewed-on: https://go-review.googlesource.com/22754
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 082881daf84a0ce3aae847c64b483c04d471e088
Author: Martin Möhrmann <martisch@uos.de>
Date:   Tue May 3 11:10:26 2016 +0200

    misc/cgo/fortran: fix gfortran compile test
    
    Fixes #14544
    
    Change-Id: I58b0b164ebbfeafe4ab32039a063df53e3018a6d
    Reviewed-on: https://go-review.googlesource.com/22730
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Sean Lake <odysseus9672@gmail.com>

commit e155bf0ce8543ebffb3b5a699861402f7abe08d7
Author: Niko Dziemba <niko@dziemba.com>
Date:   Wed Feb 10 15:57:16 2016 +0100

    archive/zip: pool flate readers
    
    Similar to the flate Writer pools already used,
    this adds pooling for flate Readers.
    
    compress/flate allows re-using of Readers, see
    https://codereview.appspot.com/97140043/
    
    In a real-world scenario when reading ~ 500 small files from a ZIP
    archive this gives a speedup of 1.5x-2x.
    
    Fixes #14289
    
    Change-Id: I2d98ad983e95ab7d97e06fd0145f619b4f47caa4
    Reviewed-on: https://go-review.googlesource.com/19416
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2f41edf120923000c92ed65ab501590fb1c8c548
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Wed May 4 11:23:24 2016 +1200

    cmd/link: always read type data for dynimport symbols
    
    Consider three shared libraries:
    
     libBase.so -- defines a type T
     lib2.so    -- references type T
     lib3.so    -- also references type T, and something from lib2
    
    lib2.so will contain a type symbol for T in its symbol table, but no
    definition. If, when linking lib3.so the linker reads the symbols from lib2.so
    before libBase.so, the linker didn't read the type data and later crashed.
    
    The fix is trivial but the test change is a bit messy because the order the
    linker reads the shared libraries in ends up depending on the order of the
    import statements in the file so I had to rename one of the test packages so
    that gofmt doesn't fix the test by accident...
    
    Fixes #15516
    
    Change-Id: I124b058f782c900a3a54c15ed66a0d91d0cde5ce
    Reviewed-on: https://go-review.googlesource.com/22744
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ba6765c237ed4dece0056b774d81e160b3839db1
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed May 4 12:20:14 2016 +1000

    cmd/api: remove debug/pe replated changes from next.txt
    
    See CL 22720 for details.
    
    Updates #15345
    
    Change-Id: If93ddbb8137d57da9846b671160b4cebe1992570
    Reviewed-on: https://go-review.googlesource.com/22752
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 9f8ecd75fcbee462dc9f1f2d1ce1025a14e1cf5b
Author: Keith Randall <khr@golang.org>
Date:   Tue May 3 17:21:36 2016 -0700

    cmd/compile: use SSA tests on legacy compiler
    
    Why not?  Because the 386 backend can't handle one of them.
    But other than that, it should work.
    
    Change-Id: Iaeb9735f8c3c281136a0734376dec5ddba21be3b
    Reviewed-on: https://go-review.googlesource.com/22748
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 9dee7771f561cf6aee081c0af6658cc81fac3918
Author: Vishvananda Ishaya <vishvananda@gmail.com>
Date:   Tue Feb 16 17:58:11 2016 -0800

    net: allow netgo to use lookup from nsswitch.conf
    
    Change https://golang.org/cl/8945 allowed Go to use its own DNS resolver
    instead of libc in a number of cases. The code parses nsswitch.conf and
    attempts to resolve things in the same order. Unfortunately, builds with
    netgo completely ignore this parsing and always search via
    hostLookupFilesDNS.
    
    This commit modifies the logic to allow binaries built with netgo to
    parse nsswitch.conf and attempt to resolve using the order specified
    there. If the parsing results in hostLookupCGo, it falls back to the
    original hostLookupFilesDNS. Tests are also added to ensure that both
    the parsing and the fallback work properly.
    
    Fixes #14354
    
    Change-Id: Ib079ad03d7036a4ec57f18352a15ba55d933f261
    Reviewed-on: https://go-review.googlesource.com/19523
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 84e808043ff40992ea1e25beb58365fd8e4f2591
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Apr 29 15:20:27 2016 -0700

    runtime: use cgo traceback for SIGPROF
    
    If we collected a cgo traceback when entering the SIGPROF signal
    handler, record it as part of the profiling stack trace.
    
    This serves as the promised test for https://golang.org/cl/21055 .
    
    Change-Id: I5f60cd6cea1d9b7c3932211483a6bfab60ed21d2
    Reviewed-on: https://go-review.googlesource.com/22650
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 4d5ac10f690bbc742a3cbf186ad8f3169a45ee26
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue May 3 23:33:32 2016 +0000

    cmd/vet: fix test's dependence on perl
    
    Change-Id: I774dbd4f90ef271a0969c3c8e65d145669312e3e
    Reviewed-on: https://go-review.googlesource.com/22745
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ross Light <light@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b64c7fc6832902acb8eebc67c887d2ef9114f644
Author: Keith Randall <khr@golang.org>
Date:   Tue May 3 13:58:28 2016 -0700

    cmd/compile: never CSE two memories
    
    It never makes sense to CSE two ops that generate memory.
    We might as well start those ops off in their own partition.
    
    Fixes #15520
    
    Change-Id: I0091ed51640f2c10cd0117f290b034dde7a86721
    Reviewed-on: https://go-review.googlesource.com/22741
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 01182425f847e4c98a53c60d0994175e21fd06dd
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue May 3 19:54:49 2016 +0000

    strings, bytes: fix Reader 0 byte read at EOF
    
    0 byte reads at EOF weren't returning EOF.
    
    Change-Id: I19b5fd5a72e83d49566a230ce4067be03f00d14b
    Reviewed-on: https://go-review.googlesource.com/22740
    Reviewed-by: Bryan Mills <bcmills@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 15f7a66f3686d24fd5ad233c6c6b1ff22daa42ae
Author: Robert Griesemer <gri@golang.org>
Date:   Mon May 2 17:03:36 2016 -0700

    cmd/compile: use correct packages when exporting/importing _ (blank) names
    
    1) Blank parameters cannot be accessed so the package doesn't matter.
       Do not export it, and consistently use localpkg when importing a
       blank parameter.
    
    2) More accurately replicate fmt.go and parser.go logic when importing
       a blank struct field. Blank struct fields get exported without
       package qualification.
       (This is actually incorrect, even with the old textual export format,
       but we will fix that in a separate change. See also issue 15514.)
    
    Fixes #15491.
    
    Change-Id: I7978e8de163eb9965964942aee27f13bf94a7c3c
    Reviewed-on: https://go-review.googlesource.com/22714
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit babdbfb8260bbe8c6305c9d3023d83cc0b3645bf
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Sun Apr 24 13:33:33 2016 +0200

    cmd/trace: make binary argument optional
    
    1.7 traces embed symbol info and we now generate symbolized pprof profiles,
    so we don't need the binary. Make binary argument optional as 1.5 traces
    still need it.
    
    Change-Id: I65eb13e3d20ec765acf85c42d42a8d7aae09854c
    Reviewed-on: https://go-review.googlesource.com/22410
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>
    Reviewed-by: Austin Clements <austin@google.com>

commit caa21475328999c1cd108b71ceb6efb7f4cf8fc4
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri Feb 26 21:57:16 2016 +0100

    runtime: per-P contexts for race detector
    
    Race runtime also needs local malloc caches and currently uses
    a mix of per-OS-thread and per-goroutine caches. This leads to
    increased memory consumption. But more importantly cache of
    synchronization objects is per-goroutine and we don't always
    have goroutine context when feeing memory in GC. As the result
    synchronization object descriptors leak (more precisely, they
    can be reused if another synchronization object is recreated
    at the same address, but it does not always help). For example,
    the added BenchmarkSyncLeak has effectively runaway memory
    consumption (based on a real long running server).
    
    This change updates race runtime with support for per-P contexts.
    BenchmarkSyncLeak now stabilizes at ~1GB memory consumption.
    
    Long term, this will allow us to remove race runtime dependency
    on glibc (as malloc is the main cornerstone).
    
    I've also implemented a different scheme to pass P context to
    race runtime: scheduler notified race runtime about association
    between G and P by calling procwire(g, p)/procunwire(g, p).
    But it turned out to be very messy as we have lots of places
    where the association changes (e.g. syscalls). So I dropped it
    in favor of the current scheme: race runtime asks scheduler
    about the current P.
    
    Fixes #14533
    
    Change-Id: Iad10d2f816a44affae1b9fed446b3580eafd8c69
    Reviewed-on: https://go-review.googlesource.com/19970
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fcd7c02c70a110c6f6dbac30ad4ac3eb435ac3fd
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri Mar 18 16:34:11 2016 +0100

    runtime: fix CPU underutilization
    
    Runqempty is a critical predicate for scheduler. If runqempty spuriously
    returns true, then scheduler can fail to schedule arbitrary number of
    runnable goroutines on idle Ps for arbitrary long time. With the addition
    of runnext runqempty predicate become broken (can spuriously return true).
    Consider that runnext is not nil and the main array is empty. Runqempty
    observes that the array is empty, then it is descheduled for some time.
    Then queue owner pushes another element to the queue evicting runnext
    into the array. Then queue owner pops runnext. Then runqempty resumes
    and observes runnext is nil and returns true. But there were no point
    in time when the queue was empty.
    
    Fix runqempty predicate to not return true spuriously.
    
    Change-Id: Ifb7d75a699101f3ff753c4ce7c983cf08befd31e
    Reviewed-on: https://go-review.googlesource.com/20858
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 499cd3371997bdb6e33377266754d20782ef134d
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Mon May 2 14:46:40 2016 +1200

    cmd/cgo: an approach to tsan that works with gcc
    
    GCC, unlike clang, does not provide any way for code being compiled to tell if
    -fsanitize-thread was passed. But cgo can look to see if that flag is being
    passed and generate different code in that case.
    
    Fixes #14602
    
    Change-Id: I86cb5318c2e35501ae399618c05af461d1252d2d
    Reviewed-on: https://go-review.googlesource.com/22688
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f459660cb85abd504845f93fdb65b1932bd6ac37
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon May 2 22:22:25 2016 +0000

    net/http: keep idle conns sorted by usage
    
    Addressing feedback from Alan Su in https://golang.org/cl/22655
    
    Change-Id: Ie0724efea2b4da67503c074e265ec7f8d7de7791
    Reviewed-on: https://go-review.googlesource.com/22709
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6b019e216b4521c2375572a4edf2e08c1d9bc754
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon May 2 17:32:14 2016 -0400

    cmd/link: bump object file version number
    
    The format has been tweaked several times in the latest cycle, so
    replace go13ld with go17ld.
    
    Change-Id: I343c49b02b7516fd781bc96ad46640579da68c59
    Reviewed-on: https://go-review.googlesource.com/22708
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 097e2c0a8a759819dde3c9b169058b2fb55b0de3
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun May 1 17:18:18 2016 -0600

    net/http: fix typo in comment in main_test
    
    Change-Id: I22d4b5a0d5c146a65d4ef77a32e23f7780ba1d95
    Reviewed-on: https://go-review.googlesource.com/22684
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 35d342b4fa749f7a3d45527580d3aa0c2a42e1f4
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon May 2 16:07:18 2016 +0000

    net/http: remove some TODOs
    
    Change-Id: Iaf200ba9a308bc8f511eec4a70dbeb014bf5fdc3
    Reviewed-on: https://go-review.googlesource.com/22690
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 53fd522c0db58f3bd75d85295f46bb06e8ab1a9b
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Apr 10 14:32:26 2016 -0700

    all: make copyright headers consistent with one space after period
    
    Follows suit with https://go-review.googlesource.com/#/c/20111.
    
    Generated by running
    $ grep -R 'Go Authors.  All' * | cut -d":" -f1 | while read F;do perl -pi -e 's/Go
    Authors.  All/Go Authors. All/g' $F;done
    
    The code in cmd/internal/unvendor wasn't changed.
    
    Fixes #15213
    
    Change-Id: I4f235cee0a62ec435f9e8540a1ec08ae03b1a75f
    Reviewed-on: https://go-review.googlesource.com/21819
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e50346d26a935cd43023856d0df65a158d867c00
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun May 1 17:03:46 2016 -0700

    cmd/cgo, misc/cgo/test: make -Wdeclaration-after-statement clean
    
    I got a complaint that cgo output triggers warnings with
    -Wdeclaration-after-statement.  I don't think it's worth testing for
    this--C has permitted declarations after statements since C99--but it is
    easy enough to fix.  It may break again; so it goes.
    
    This CL also fixes errno handling to avoid getting confused if the tsan
    functions happen to change the global errno variable.
    
    Change-Id: I0ec7c63a6be5653ef44799d134c8d27cb5efa441
    Reviewed-on: https://go-review.googlesource.com/22686
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 45f39fb46747d0c46bc25e6ef605c00e96e2dc07
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Mon May 2 13:58:48 2016 +1200

    cmd/link: pass -Wl,-z,relro to host linker in -buildmode=PIE
    
    Fixes #15485
    
    Change-Id: I8e9314be91db89873130b232b589a284822e6643
    Reviewed-on: https://go-review.googlesource.com/22687
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5325fbc7dbab155a88efab0bcd128236e97b3349
Author: Keith Randall <khr@golang.org>
Date:   Fri Apr 29 12:09:32 2016 -0700

    cmd/compile: don't SSA any variables when -N
    
    Turn SSAing of variables off when compiling with optimizations off.
    This helps keep variable names around that would otherwise be
    optimized away.
    
    Fixes #14744
    
    Change-Id: I31db8cf269c068c7c5851808f13e5955a09810ca
    Reviewed-on: https://go-review.googlesource.com/22681
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 5776c20164e5852ec13828db10277019db86180b
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Mon May 2 04:44:46 2016 +0900

    net/http: gofmt -w -s
    
    Change-Id: I7e07888e90c7449f119e74b97995efcd7feef76e
    Reviewed-on: https://go-review.googlesource.com/22682
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ade0eb2f0689b4d88d425451387c516013fd4b20
Author: Keith Randall <khr@golang.org>
Date:   Sat Apr 30 22:28:37 2016 -0700

    cmd/compile: fix reslice
    
    := is the wrong thing here.  The new variable masks the old
    variable so we allocate the slice afresh each time around the loop.
    
    Change-Id: I759c30e1bfa88f40decca6dd7d1e051e14ca0844
    Reviewed-on: https://go-review.googlesource.com/22679
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 24c05e7e695767a46ce6e48d2492c29ba7adffc4
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun May 1 15:18:13 2016 +0000

    net/http: fix typo in comment
    
    Change-Id: I753e62879a56582a9511e3f34fdeac929202efbf
    Reviewed-on: https://go-review.googlesource.com/22680
    Reviewed-by: Ralph Corderoy <ralph@inputplus.co.uk>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ca8b6270724026fb7697e9f9510d1e6865ed7045
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 16 09:35:32 2016 -0700

    net/http: add Response.Uncompressed bool
    
    The Transport's automatic gzip uncompression lost information in the
    process (the compressed Content-Length, if known). Normally that's
    okay, but it's not okay for reverse proxies which have to be able to
    generate a valid HTTP response from the Transport's provided
    *Response.
    
    Reverse proxies should normally be disabling compression anyway and
    just piping the compressed pipes though and not wasting CPU cycles
    decompressing them. So also document that on the new Uncompressed
    field.
    
    Then, using the new field, fix Response.Write to not inject a bogus
    "Connection: close" header when it doesn't see a transfer encoding or
    content-length.
    
    Updates #15366 (the http2 side remains, once this is submitted)
    
    Change-Id: I476f40aa14cfa7aa7b3bf99021bebba4639f9640
    Reviewed-on: https://go-review.googlesource.com/22671
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a9cf0b1e1e2a66db547fcabb7188465e4ac54700
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 30 21:11:26 2016 -0700

    net/http: provide access to the listener address an HTTP request arrived on
    
    This adds a context key named LocalAddrContextKey (for now, see #15229) to
    let users access the net.Addr of the net.Listener that accepted the connection
    that sent an HTTP request. This is similar to ServerContextKey which provides
    access to the *Server. (A Server may have multiple Listeners)
    
    Fixes #6732
    
    Change-Id: I74296307b68aaaab8df7ad4a143e11b5227b5e62
    Reviewed-on: https://go-review.googlesource.com/22672
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit abc1472d78c70888473634497b49b1c2e1bb6569
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 30 21:11:42 2016 -0500

    net/http: add Transport.IdleConnTimeout
    
    Don't keep idle HTTP client connections open forever. Add a new knob,
    Transport.IdleConnTimeout, and make the default be 90 seconds. I
    figure 90 seconds is more than a minute, and less than infinite, and I
    figure enough code has things waking up once a minute polling APIs.
    
    This also removes the Transport's idleCount field which was unused and
    redundant with the size of the idleLRU map (which was actually used).
    
    Change-Id: Ibb698a9a9a26f28e00a20fe7ed23f4afb20c2322
    Reviewed-on: https://go-review.googlesource.com/22670
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0ab78df9ea602d6bc9cf45dbd610c3d6f534cb58
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 30 21:27:04 2016 -0500

    net/http: fix a few crashes with a ClientTrace with nil funcs
    
    And add a test.
    
    Updates #12580
    
    Change-Id: Ia7eaba09b8e7fd0eddbcaefb948d01ab10af876e
    Reviewed-on: https://go-review.googlesource.com/22659
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b3a130e81a0c3c2508f483af15e57d181c5cdc1e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 30 20:11:34 2016 -0700

    net/http: document some errors more, mark ErrWriteAfterFlush as unused
    
    Fixes #15150
    
    Change-Id: I1a892d5b0516a37dac050d3bb448e0a2571db16e
    Reviewed-on: https://go-review.googlesource.com/22658
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit d713e8e8069ee052ef4d9eac49a0c74a3b5c2c92
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Apr 28 15:01:41 2016 -0700

    archive/zip: improve BenchmarkCompressedZipGarbage
    
    Before this CL:
    
    $ go test -bench=CompressedZipGarbage -count=5 -run=NONE archive/zip
    BenchmarkCompressedZipGarbage-8        50  20677087 ns/op   42973 B/op      47 allocs/op
    BenchmarkCompressedZipGarbage-8       100  20584764 ns/op   24294 B/op      47 allocs/op
    BenchmarkCompressedZipGarbage-8        50  20859221 ns/op   42973 B/op      47 allocs/op
    BenchmarkCompressedZipGarbage-8       100  20901176 ns/op   24294 B/op      47 allocs/op
    BenchmarkCompressedZipGarbage-8        50  21282409 ns/op   42973 B/op      47 allocs/op
    
    The B/op number is effectively meaningless. There
    is a surprisingly large one-time cost that gets
    divided by the number of iterations that your
    machine can get through in a second.
    
    This CL discards the first run, which helps.
    It is not a panacea. Running with -benchtime=10s
    will allow the sync.Pool to be emptied,
    which brings the problem back.
    However, since there are more iterations to divide
    the cost through, it’s not quite as bad,
    and running with a high benchtime is rare.
    
    This CL changes the meaning of the B/op number,
    which is unfortunate, since it won’t have the
    same order of magnitude as previous Go versions.
    But it wasn’t really comparable before anyway,
    since it didn’t have any reliable meaning at all.
    
    After this CL:
    
    $ go test -bench=CompressedZipGarbage -count=5 -run=NONE archive/zip
    BenchmarkCompressedZipGarbage-8   	     100	  20881890 ns/op	    5616 B/op	      47 allocs/op
    BenchmarkCompressedZipGarbage-8   	      50	  20622757 ns/op	    5616 B/op	      47 allocs/op
    BenchmarkCompressedZipGarbage-8   	      50	  20628193 ns/op	    5616 B/op	      47 allocs/op
    BenchmarkCompressedZipGarbage-8   	     100	  20756612 ns/op	    5616 B/op	      47 allocs/op
    BenchmarkCompressedZipGarbage-8   	     100	  20639774 ns/op	    5616 B/op	      47 allocs/op
    
    Change-Id: Iedee04f39328974c7fa272a6113d423e7ffce50f
    Reviewed-on: https://go-review.googlesource.com/22585
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3836354f113367a6a405ac17a65f406514ea9313
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Apr 28 13:51:12 2016 -0500

    doc: update go1.7.txt
    
    Change-Id: I53dd5affc3a1e1f741fe44c7ce691bb2cd432764
    Reviewed-on: https://go-review.googlesource.com/22657
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3b0b30727ca6caeddcc6b49bc97960826b9f15d1
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:14 2016 -0400

    cmd/internal/obj/mips, cmd/link: add support TLS relocation for mips64x
    
    a new relocation R_ADDRMIPSTLS is added, which resolves to 16-bit offset
    of a TLS address on mips64x.
    
    Change-Id: Ic60d0e1ba49ff1c433cead242f5884677ab227a5
    Reviewed-on: https://go-review.googlesource.com/19804
    Reviewed-by: Minux Ma <minux@golang.org>

commit 77c7f12438aa83d9ba02c7fd48e4a0a288ed9123
Author: Austin Clements <austin@google.com>
Date:   Sat Apr 30 21:47:30 2016 -0400

    runtime: update some comments
    
    This updates some comments that became out of date when we moved the
    mark bit out of the heap bitmap and started using the high bit for the
    first word as a scan/dead bit.
    
    Change-Id: I4a572d16db6114cadff006825466c1f18359f2db
    Reviewed-on: https://go-review.googlesource.com/22662
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 5d002dbc2179c99beb47490d316c53f687c9565a
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:20 2016 -0400

    runtime/cgo: add linux/mips64x cgo support
    
    MIPS N64 ABI passes arguments in registers R4-R11, return value in R2.
    R16-R23, R28, R30 and F24-F31 are callee-save. gcc PIC code expects
    to be called with indirect call through R25.
    
    Change-Id: I24f582b4b58e1891ba9fd606509990f95cca8051
    Reviewed-on: https://go-review.googlesource.com/19805
    Reviewed-by: Minux Ma <minux@golang.org>

commit 073d292c453f85f6c50f063ba0a4d0d1b328dadc
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:09 2016 -0400

    cmd/link, runtime: add external linking support for linux/mips64x
    
    Fixes #12560
    
    Change-Id: Ic2004fc7b09f2dbbf83c41f8c6307757c0e1676d
    Reviewed-on: https://go-review.googlesource.com/19803
    Reviewed-by: Minux Ma <minux@golang.org>

commit b13b249f43d4d38b145cd01135026286052bbc88
Author: Frits van Bommel <fvbommel@gmail.com>
Date:   Sat Apr 30 11:13:29 2016 +0200

    cmd/compile: Improve readability of HTML produced by GOSSAFUNC
    
    Factor out the Aux/AuxInt handling in (*Value).LongString() and
    use it in (*Value).LongHTML() as well.
    This especially improves readability of auxFloat32, auxFloat64,
    and auxSymValAndOff values which would otherwise be printed as
    opaque integers.
    This change also makes LongString() slightly less verbose by
    eliding offsets that are zero (as is very often the case).
    
    Additionally, ensure the HTML is interpreted as UTF-8 so that
    non-ASCII characters (especially the "middle dots" in some symbols)
    show up correctly.
    
    Change-Id: Ie26221df876faa056d322b3e423af63f33cd109d
    Reviewed-on: https://go-review.googlesource.com/22641
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Frits van Bommel <fvbommel@gmail.com>

commit 981395103e9addfc494245979063dc59a327e888
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:18:02 2016 -0400

    cmd/internal/obj/mips et al.: introduce SB register on mips64x
    
    SB register (R28) is introduced for access external addresses with shorter
    instruction sequences. It is loaded at entry points. External data within
    2G of SB can be accessed this way.
    
    cmd/internal/obj: relocaltion R_ADDRMIPS is split into two relocations
    R_ADDRMIPS and R_ADDRMIPSU, handling the low 16 bits and the "upper" 16
    bits of external addresses, respectively, since the instructios may not
    be adjacent. It might be better if relocation Variant could be used.
    
    cmd/link/internal/mips64: support new relocations.
    
    cmd/compile/internal/mips64: reserve SB register.
    
    runtime: initialize SB register at entry points.
    
    Change-Id: I5f34868f88c5a9698c042a8a1f12f76806c187b9
    Reviewed-on: https://go-review.googlesource.com/19802
    Reviewed-by: Minux Ma <minux@golang.org>

commit 8dc0444a04a8a43887b4ca3753ee63b430cf2602
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:17:44 2016 -0400

    cmd/asm, cmd/internal/obj/mips: add an alias of RSB on mips64x
    
    Change-Id: I724ce0a48c1aeed14267c049fa415a6fa2fffbcf
    Reviewed-on: https://go-review.googlesource.com/19864
    Reviewed-by: Minux Ma <minux@golang.org>

commit a409fb80b0ebc1353336a9c03f8db408fbc6d1d5
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:17:36 2016 -0400

    cmd/internal/obj/mips, runtime: change REGTMP to R23
    
    Leave R28 to SB register, which will be introduced in CL 19802.
    
    Change-Id: I1cf7a789695c5de664267ec8086bfb0b043ebc14
    Reviewed-on: https://go-review.googlesource.com/19863
    Reviewed-by: Minux Ma <minux@golang.org>

commit 9bc1e2065c3a4ac257bebc5c7c07cd5c844b11ba
Author: Cherry Zhang <lunaria21@gmail.com>
Date:   Wed Apr 27 22:16:51 2016 -0400

    cmd/asm/internal/asm/testdata: remove WORD $foo(SB) from mips64.s
    
    on mips64, address is 64 bit, not a WORD. also it is never used anywhere.
    
    Change-Id: Ic6bf6d6a21c8d2f1eb7bfe9efc5a29186ec2a8ef
    Reviewed-on: https://go-review.googlesource.com/19801
    Reviewed-by: Minux Ma <minux@golang.org>

commit 81b2ea4d34a42bee14e1ed17d5166546be957849
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 30 17:17:26 2016 -0500

    net/http: add Transport.MaxIdleConns limit
    
    The HTTP client had a limit for the maximum number of idle connections
    per-host, but not a global limit.
    
    This CLs adds a global idle connection limit too,
    Transport.MaxIdleConns.
    
    All idle conns are now also stored in a doubly-linked list. When there
    are too many, the oldest one is closed.
    
    Fixes #15461
    
    Change-Id: I72abbc28d140c73cf50f278fa70088b45ae0deef
    Reviewed-on: https://go-review.googlesource.com/22655
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 38cfaa5f0ac1e8b9c7528649f01e4b0edcd4a788
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 30 17:29:28 2016 -0500

    net/http: expand documentation of Server.MaxHeaderBytes
    
    Clarify that it includes the RFC 7230 "request-line".
    
    Fixes #15494
    
    Change-Id: I9cc5dd5f2d85ebf903229539208cec4da5c38d04
    Reviewed-on: https://go-review.googlesource.com/22656
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 4e0cd1eeef419b221fda3dd3966be71095f0b4ce
Author: Kevin Burke <kev@inburke.com>
Date:   Sat Apr 23 11:00:05 2016 -0700

    database/sql: clone data for named []byte types
    
    Previously named byte types like json.RawMessage could get dirty
    database memory from a call to Scan. These types would activate a
    code path that didn't clone the byte data coming from the database
    before assigning it. Another thread could then overwrite the byte
    array in src, which has unexpected consequences.
    
    Originally reported by Jason Moiron; the patch and test are his
    suggestions. Fixes #13905.
    
    Change-Id: Iacfef61cbc9dd51c8fccef9b2b9d9544c77dd0e0
    Reviewed-on: https://go-review.googlesource.com/22393
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a20fd1f6ba668ec0bd8c432d26def2b65cc6609a
Author: Austin Clements <austin@google.com>
Date:   Fri Apr 29 14:51:48 2016 -0400

    runtime: reclaim scan/dead bit in first word
    
    With the switch to separate mark bitmaps, the scan/dead bit for the
    first word of each object is now unused. Reclaim this bit and use it
    as a scan/dead bit, just like words three and on. The second word is
    still used for checkmark.
    
    This dramatically simplifies heapBitsSetTypeNoScan and hasPointers,
    since they no longer need different cases for 1, 2, and 3+ word
    objects. They can instead just manipulate the heap bitmap for the
    first word and be done with it.
    
    In order to enable this, we change heapBitsSetType and runGCProg to
    always set the scan/dead bit to scan for the first word on every code
    path. Since these functions only apply to types that have pointers,
    there's no need to do this conditionally: it's *always* necessary to
    set the scan bit in the first word.
    
    We also change every place that scans an object and checks if there
    are more pointers. Rather than only checking morePointers if the word
    is >= 2, we now check morePointers if word != 1 (since that's the
    checkmark word).
    
    Looking forward, we should probably reclaim the checkmark bit, too,
    but that's going to be quite a bit more work.
    
    Tested by setting doubleCheck in heapBitsSetType and running all.bash
    on both linux/amd64 and linux/386, and by running GOGC=10 all.bash.
    
    This particularly improves the FmtFprintf* go1 benchmarks, since they
    do a large amount of noscan allocation.
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.34s ± 1%     2.38s ± 1%  +1.70%  (p=0.000 n=17+19)
    Fannkuch11-12                2.09s ± 0%     2.09s ± 1%    ~     (p=0.276 n=17+16)
    FmtFprintfEmpty-12          44.9ns ± 2%    44.8ns ± 2%    ~     (p=0.340 n=19+18)
    FmtFprintfString-12          127ns ± 0%     125ns ± 0%  -1.57%  (p=0.000 n=16+15)
    FmtFprintfInt-12             128ns ± 0%     122ns ± 1%  -4.45%  (p=0.000 n=15+20)
    FmtFprintfIntInt-12          207ns ± 1%     193ns ± 0%  -6.55%  (p=0.000 n=19+14)
    FmtFprintfPrefixedInt-12     197ns ± 1%     191ns ± 0%  -2.93%  (p=0.000 n=17+18)
    FmtFprintfFloat-12           263ns ± 0%     248ns ± 1%  -5.88%  (p=0.000 n=15+19)
    FmtManyArgs-12               794ns ± 0%     779ns ± 1%  -1.90%  (p=0.000 n=18+18)
    GobDecode-12                7.14ms ± 2%    7.11ms ± 1%    ~     (p=0.072 n=20+20)
    GobEncode-12                5.85ms ± 1%    5.82ms ± 1%  -0.49%  (p=0.000 n=20+20)
    Gzip-12                      218ms ± 1%     215ms ± 1%  -1.22%  (p=0.000 n=19+19)
    Gunzip-12                   36.8ms ± 0%    36.7ms ± 0%  -0.18%  (p=0.006 n=18+20)
    HTTPClientServer-12         77.1µs ± 4%    77.1µs ± 3%    ~     (p=0.945 n=19+20)
    JSONEncode-12               15.6ms ± 1%    15.9ms ± 1%  +1.68%  (p=0.000 n=18+20)
    JSONDecode-12               55.2ms ± 1%    53.6ms ± 1%  -2.93%  (p=0.000 n=17+19)
    Mandelbrot200-12            4.05ms ± 1%    4.05ms ± 0%    ~     (p=0.306 n=17+17)
    GoParse-12                  3.14ms ± 1%    3.10ms ± 1%  -1.31%  (p=0.000 n=19+18)
    RegexpMatchEasy0_32-12      69.3ns ± 1%    70.0ns ± 0%  +0.89%  (p=0.000 n=19+17)
    RegexpMatchEasy0_1K-12       237ns ± 1%     236ns ± 0%  -0.62%  (p=0.000 n=19+16)
    RegexpMatchEasy1_32-12      69.5ns ± 1%    70.3ns ± 1%  +1.14%  (p=0.000 n=18+17)
    RegexpMatchEasy1_1K-12       377ns ± 1%     366ns ± 1%  -3.03%  (p=0.000 n=15+19)
    RegexpMatchMedium_32-12      107ns ± 1%     107ns ± 2%    ~     (p=0.318 n=20+19)
    RegexpMatchMedium_1K-12     33.8µs ± 3%    33.5µs ± 1%  -1.04%  (p=0.001 n=20+19)
    RegexpMatchHard_32-12       1.68µs ± 1%    1.73µs ± 0%  +2.50%  (p=0.000 n=20+18)
    RegexpMatchHard_1K-12       50.8µs ± 1%    52.0µs ± 1%  +2.50%  (p=0.000 n=19+18)
    Revcomp-12                   381ms ± 1%     385ms ± 1%  +1.00%  (p=0.000 n=17+18)
    Template-12                 64.9ms ± 3%    62.6ms ± 1%  -3.55%  (p=0.000 n=19+18)
    TimeParse-12                 324ns ± 0%     328ns ± 1%  +1.25%  (p=0.000 n=18+18)
    TimeFormat-12                345ns ± 0%     334ns ± 0%  -3.31%  (p=0.000 n=15+17)
    [Geo mean]                  52.1µs         51.5µs       -1.00%
    
    Change-Id: I13e74da3193a7f80794c654f944d1f0d60817049
    Reviewed-on: https://go-review.googlesource.com/22632
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d5e3d08b3ad0048c540727b5512f29ecc70ef51a
Author: Austin Clements <austin@google.com>
Date:   Fri Apr 29 15:19:11 2016 -0400

    runtime: use morePointers and isPointer in more places
    
    This makes this code better self-documenting and makes it easier to
    find these places in the future.
    
    Change-Id: I31dc5598ae67f937fb9ef26df92fd41d01e983c3
    Reviewed-on: https://go-review.googlesource.com/22631
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a5d3f7ece9033393166d0c74cd1121c15032ba4d
Author: Austin Clements <austin@google.com>
Date:   Fri Apr 29 15:17:34 2016 -0400

    runtime: avoid conditional execution in morePointers and isPointer
    
    heapBits.bits is carefully written to produce good machine code. Use
    it in heapBits.morePointers and heapBits.isPointer to get good machine
    code there, too.
    
    Change-Id: I208c7d0d38697e7a22cad67f692162589b75f1e2
    Reviewed-on: https://go-review.googlesource.com/22630
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7a60a962b9284763f8ff239aae54d1d858dd4543
Author: Keith Randall <khr@golang.org>
Date:   Fri Apr 29 17:10:02 2016 -0700

    cmd/compile: ecx is reserved for PIC, don't let peep work on it
    
    Fixes #15496
    
    Change-Id: Ieb5be1caa4b1c23e23b20d56c1a0a619032a9f5d
    Reviewed-on: https://go-review.googlesource.com/22652
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 58f52cbb79d6ed369bd3fed44fe615a23c721189
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Apr 29 20:17:06 2016 -0400

    runtime: fix cgocallback_gofunc on ppc64x
    
    Fix issues introduced in 5f9a870.
    
    Change-Id: Ia75945ef563956613bf88bbe57800a96455c265d
    Reviewed-on: https://go-review.googlesource.com/22661
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9fe572e509470d53c220e7f43a825d70a7f0d8b3
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Apr 29 15:44:36 2016 -0700

    runtime: fix cgocallback_gofunc argument passing on arm64
    
    Change-Id: I4b34bcd5cde71ecfbb352b39c4231de6168cc7f3
    Reviewed-on: https://go-review.googlesource.com/22651
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit 36b6c03827f51a8464a135d27212100d893f582c
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Apr 29 15:56:57 2016 -0700

    root: remove dev.garbage file
    
    Change-Id: I99b2ca52824341d986090f5c78ab4f396594bcdf
    Reviewed-on: https://go-review.googlesource.com/22660
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5f9a870bf1bf461ca3609502608b12cc4aab189a
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Apr 27 14:18:29 2016 -0700

    cmd/cgo, runtime, runtime/cgo: use cgo context function
    
    Add support for the context function set by runtime.SetCgoTraceback.
    The context function was added in CL 17761, without support.
    This CL is the support.
    
    This CL has not been tested for real C code, as a working context
    function for C code requires unwind support that does not seem to exist.
    I wanted to get the CL out before the freeze.
    
    I apologize for the length of this CL.  It's mostly plumbing, but
    unfortunately the plumbing is processor-specific.
    
    Change-Id: I8ce11a0de9b3dafcc29efd2649d776e93bff0e90
    Reviewed-on: https://go-review.googlesource.com/22508
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c717675c35cb436bdab62091a6288843aa1c863c
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Apr 17 21:26:23 2016 -0400

    crypto/cipher, crypto/aes: add s390x implementation of AES-CTR
    
    This commit adds the new 'ctrAble' interface to the crypto/cipher
    package. The role of ctrAble is the same as gcmAble but for CTR
    instead of GCM. It allows block ciphers to provide optimized CTR
    implementations.
    
    The primary benefit of adding CTR support to the s390x AES
    implementation is that it allows us to encrypt the counter values
    in bulk, giving the cipher message instruction a larger chunk of
    data to work on per invocation.
    
    The xorBytes assembly is necessary because xorBytes becomes a
    bottleneck when CTR is done in this way. Hopefully it will be
    possible to remove this once s390x has migrated to the ssa
    backend.
    
    name      old speed     new speed     delta
    AESCTR1K  160MB/s ± 6%  867MB/s ± 0%  +442.42%  (p=0.000 n=9+10)
    
    Change-Id: I1ae16b0ce0e2641d2bdc7d7eabc94dd35f6e9318
    Reviewed-on: https://go-review.googlesource.com/22195
    Reviewed-by: Adam Langley <agl@golang.org>

commit 2f8475648a5500830561ea03960a1425e1ff0993
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Apr 25 21:46:02 2016 -0400

    crypto/cipher, crypto/aes: add s390x implementation of AES-CBC
    
    This commit adds the cbcEncAble and cbcDecAble interfaces that
    can be implemented by block ciphers that support an optimized
    implementation of CBC. This is similar to what is done for GCM
    with the gcmAble interface.
    
    The cbcEncAble, cbcDecAble and gcmAble interfaces all now have
    tests to ensure they are detected correctly in the cipher
    package.
    
    name             old speed     new speed      delta
    AESCBCEncrypt1K  152MB/s ± 1%  1362MB/s ± 0%  +795.59%   (p=0.000 n=10+9)
    AESCBCDecrypt1K  143MB/s ± 1%  1362MB/s ± 0%  +853.00%   (p=0.000 n=10+9)
    
    Change-Id: I715f686ab3686b189a3dac02f86001178fa60580
    Reviewed-on: https://go-review.googlesource.com/22523
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit cd956576ae618a1de829e42f9ec557e6cee338d1
Author: Keith Randall <khr@golang.org>
Date:   Fri Apr 29 09:02:27 2016 -0700

    cmd/compile: make vet happy with ssa code
    
    Fixes #15488
    
    Change-Id: I054eb1e1c859de315e3cdbdef5428682bce693fd
    Reviewed-on: https://go-review.googlesource.com/22609
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 56b54912628934707977a2a0a3824288c0286830
Merge: d8d3351 e9eaa18
Author: Rick Hudson <rlh@golang.org>
Date:   Fri Apr 29 13:49:18 2016 -0400

    Merge remote-tracking branch 'origin/dev.garbage'
    
    This commit moves the GC from free list allocation to
    bit mark allocation. Instead of using the bitmaps
    generated during the mark phases to generate free
    list and then using the free lists for allocation we
    allocate directly from the bitmaps.
    
    The change in the garbage benchmark
    
    name              old time/op  new time/op  delta
    XBenchGarbage-12  2.22ms ± 1%  2.13ms ± 1%  -3.90%  (p=0.000 n=18+18)
    
    Change-Id: I17f57233336f0ca5ef5404c3be4ecb443ab622aa

commit e9eaa181fcadc2162baa62ccd8bfeb610acfdd55
Author: Rick Hudson <rlh@golang.org>
Date:   Fri Apr 29 12:09:36 2016 -0400

    [dev.garbage] runtime: simplify nextFreeFast so it is inlined
    
    nextFreeFast is currently not inlined by the compiler due
    to its size and complexity. This CL simplifies
    nextFreeFast by letting the slow path handle (nextFree)
    handle a corner cases.
    
    Change-Id: Ia9c5d1a7912bcb4bec072f5fd240f0e0bafb20e4
    Reviewed-on: https://go-review.googlesource.com/22598
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Austin Clements <austin@google.com>

commit d8d33514f9e8c80d504ab4a61ef96621afc3647d
Author: David Chase <drchase@google.com>
Date:   Fri Apr 29 11:15:16 2016 -0400

    cmd/compile: Move divconst_test out of test/bench/go1
    
    This is necessary to avoid disrupting the go1 suite and gives
    us a place to put other tests of basic compiler function and
    correctness.
    
    Change-Id: I36933819ff2bfe6a2121fff2be9a98efd2123d9a
    Reviewed-on: https://go-review.googlesource.com/22597
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit fa9435cdff443d12c526b0436435925dd52e8503
Author: Keith Randall <khr@golang.org>
Date:   Tue Apr 26 12:08:31 2016 -0700

    cmd/compile: clean up rewrite rules
    
    Break really long lines.
    Add spacing to line up columns.
    
    In AMD64, put all the optimization rules after all the
    lowering rules.
    
    Change-Id: I45cc7368bf278416e67f89e74358db1bd4326a93
    Reviewed-on: https://go-review.googlesource.com/22470
    Reviewed-by: David Chase <drchase@google.com>

commit b3579c095e00f89d8c92c2aa4fb4af222a96f429
Author: Austin Clements <austin@google.com>
Date:   Fri Apr 29 10:57:06 2016 -0400

    [dev.garbage] runtime: revive sweep fast path
    
    sweep used to skip mcental.freeSpan (and its locking) if it didn't
    find any new free objects. We lost that optimization when the
    freed-object counting changed in dad83f7 to count total free objects
    instead of newly freed objects.
    
    The previous commit brings back counting of newly freed objects, so we
    can easily revive this optimization by checking that count (like we
    used to) instead of the total free objects count.
    
    Change-Id: I43658707a1c61674d0366124d5976b00d98741a9
    Reviewed-on: https://go-review.googlesource.com/22596
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit d97625ae9e7195a68d1c9f2b2ff54eb85545982e
Author: Austin Clements <austin@google.com>
Date:   Fri Apr 29 09:44:53 2016 -0400

    [dev.garbage] runtime: fix nfree accounting
    
    Commit 8dda1c4 changed the meaning of "nfree" in sweep from the number
    of newly freed objects to the total number of free objects in the
    span, but didn't update where sweep added nfree to c.local_nsmallfree.
    Hence, we're over-accounting the number of frees. This is causing
    TestArrayHash to fail with "too many allocs NNN - hash not balanced".
    
    Fix this by computing the number of newly freed objects and adding
    that to c.local_nsmallfree, so it behaves like it used to. Computing
    this requires a small tweak to mallocgc: apparently we've never set
    s.allocCount when allocating a large object; fix this by setting it to
    1 so sweep doesn't get confused.
    
    Change-Id: I31902ffd310110da4ffd807c5c06f1117b872dc8
    Reviewed-on: https://go-review.googlesource.com/22595
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>

commit 6d11490539e3aa459066b94c6587f5429dfe7a31
Author: Austin Clements <austin@google.com>
Date:   Thu Apr 28 15:49:39 2016 -0400

    [dev.garbage] runtime: fix allocfreetrace
    
    We broke tracing of freed objects in GODEBUG=allocfreetrace=1 mode
    when we removed the sweep over the mark bitmap. Fix it by
    re-introducing the sweep over the bitmap specifically if we're in
    allocfreetrace mode. This doesn't have to be even remotely efficient,
    since the overhead of allocfreetrace is huge anyway, so we can keep
    the code for this down to just a few lines.
    
    Change-Id: I9e176b3b04c73608a0ea3068d5d0cd30760ebd40
    Reviewed-on: https://go-review.googlesource.com/22592
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 38f674687a5dbce63af60a0a52892f666d7c626c
Author: Austin Clements <austin@google.com>
Date:   Thu Apr 28 15:32:01 2016 -0400

    [dev.garbage] runtime: reintroduce no-zeroing optimization
    
    Currently we always zero objects when we allocate them. We used to
    have an optimization that would not zero objects that had not been
    allocated since the whole span was last zeroed (either by getting it
    from the system or by getting it from the heap, which does a bulk
    zero), but this depended on the sweeper clobbering the first two words
    of each object. Hence, we lost this optimization when the bitmap
    sweeper went away.
    
    Re-introduce this optimization using a different mechanism. Each span
    already keeps a flag indicating that it just came from the OS or was
    just bulk zeroed by the mheap. We can simply use this flag to know
    when we don't need to zero an object. This is slightly less efficient
    than the old optimization: if a span gets allocated and partially
    used, then GC happens and the span gets returned to the mcentral, then
    the span gets re-acquired, the old optimization knew that it only had
    to re-zero the objects that had been reclaimed, whereas this
    optimization will re-zero everything. However, in this case, you're
    already paying for the garbage collection, and you've only wasted one
    zeroing of the span, so in practice there seems to be little
    difference. (If we did want to revive the full optimization, each span
    could keep track of a frontier beyond which all free slots are zeroed.
    I prototyped this and it didn't obvious do any better than the much
    simpler approach in this commit.)
    
    This significantly improves BinaryTree17, which is allocation-heavy
    (and runs first, so most pages are already zeroed), and slightly
    improves everything else.
    
    name              old time/op  new time/op  delta
    XBenchGarbage-12  2.15ms ± 1%  2.14ms ± 1%  -0.80%  (p=0.000 n=17+17)
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.71s ± 1%     2.56s ± 1%  -5.73%        (p=0.000 n=18+19)
    DivconstI64-12              1.70ns ± 1%    1.70ns ± 1%    ~           (p=0.562 n=18+18)
    DivconstU64-12              1.74ns ± 2%    1.74ns ± 1%    ~           (p=0.394 n=20+20)
    DivconstI32-12              1.74ns ± 0%    1.74ns ± 0%    ~     (all samples are equal)
    DivconstU32-12              1.66ns ± 1%    1.66ns ± 0%    ~           (p=0.516 n=15+16)
    DivconstI16-12              1.84ns ± 0%    1.84ns ± 0%    ~     (all samples are equal)
    DivconstU16-12              1.82ns ± 0%    1.82ns ± 0%    ~     (all samples are equal)
    DivconstI8-12               1.79ns ± 0%    1.79ns ± 0%    ~     (all samples are equal)
    DivconstU8-12               1.60ns ± 0%    1.60ns ± 1%    ~           (p=0.603 n=17+19)
    Fannkuch11-12                2.11s ± 1%     2.11s ± 0%    ~           (p=0.333 n=16+19)
    FmtFprintfEmpty-12          45.1ns ± 4%    45.4ns ± 5%    ~           (p=0.111 n=20+20)
    FmtFprintfString-12          134ns ± 0%     129ns ± 0%  -3.45%        (p=0.000 n=18+16)
    FmtFprintfInt-12             131ns ± 1%     129ns ± 1%  -1.54%        (p=0.000 n=16+18)
    FmtFprintfIntInt-12          205ns ± 2%     203ns ± 0%  -0.56%        (p=0.014 n=20+18)
    FmtFprintfPrefixedInt-12     200ns ± 2%     197ns ± 1%  -1.48%        (p=0.000 n=20+18)
    FmtFprintfFloat-12           256ns ± 1%     256ns ± 0%  -0.21%        (p=0.008 n=18+20)
    FmtManyArgs-12               805ns ± 0%     804ns ± 0%  -0.19%        (p=0.001 n=18+18)
    GobDecode-12                7.21ms ± 1%    7.14ms ± 1%  -0.92%        (p=0.000 n=19+20)
    GobEncode-12                5.88ms ± 1%    5.88ms ± 1%    ~           (p=0.641 n=18+19)
    Gzip-12                      218ms ± 1%     218ms ± 1%    ~           (p=0.271 n=19+18)
    Gunzip-12                   37.1ms ± 0%    36.9ms ± 0%  -0.29%        (p=0.000 n=18+17)
    HTTPClientServer-12         78.1µs ± 2%    77.4µs ± 2%    ~           (p=0.070 n=19+19)
    JSONEncode-12               15.5ms ± 1%    15.5ms ± 0%    ~           (p=0.063 n=20+18)
    JSONDecode-12               56.1ms ± 0%    55.4ms ± 1%  -1.18%        (p=0.000 n=19+18)
    Mandelbrot200-12            4.05ms ± 0%    4.06ms ± 0%  +0.29%        (p=0.001 n=18+18)
    GoParse-12                  3.28ms ± 1%    3.21ms ± 1%  -2.30%        (p=0.000 n=20+20)
    RegexpMatchEasy0_32-12      69.4ns ± 2%    69.3ns ± 1%    ~           (p=0.205 n=18+16)
    RegexpMatchEasy0_1K-12       239ns ± 0%     239ns ± 0%    ~     (all samples are equal)
    RegexpMatchEasy1_32-12      69.4ns ± 1%    69.4ns ± 1%    ~           (p=0.620 n=15+18)
    RegexpMatchEasy1_1K-12       370ns ± 1%     369ns ± 2%    ~           (p=0.088 n=20+20)
    RegexpMatchMedium_32-12      108ns ± 0%     108ns ± 0%    ~     (all samples are equal)
    RegexpMatchMedium_1K-12     33.6µs ± 3%    33.5µs ± 3%    ~           (p=0.718 n=20+20)
    RegexpMatchHard_32-12       1.68µs ± 1%    1.67µs ± 2%    ~           (p=0.316 n=20+20)
    RegexpMatchHard_1K-12       50.5µs ± 3%    50.4µs ± 3%    ~           (p=0.659 n=20+20)
    Revcomp-12                   381ms ± 1%     381ms ± 1%    ~           (p=0.916 n=19+18)
    Template-12                 66.5ms ± 1%    65.8ms ± 2%  -1.08%        (p=0.000 n=20+20)
    TimeParse-12                 317ns ± 0%     319ns ± 0%  +0.48%        (p=0.000 n=19+12)
    TimeFormat-12                338ns ± 0%     338ns ± 0%    ~           (p=0.124 n=19+18)
    [Geo mean]                  5.99µs         5.96µs       -0.54%
    
    Change-Id: I638ffd9d9f178835bbfa499bac20bd7224f1a907
    Reviewed-on: https://go-review.googlesource.com/22591
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 1fb4e4de26034cb2822bd8a9eadeb8e2b215d796
Author: Nigel Tao <nigeltao@golang.org>
Date:   Fri Apr 29 17:17:44 2016 +1000

    compress/flate: use a constant hash table size for Best Speed.
    
    This makes compress/flate's version of Snappy diverge from the upstream
    golang/snappy version, but the latter has a goal of matching C++ snappy
    output byte-for-byte. Both C++ and the asm version of golang/snappy can
    use a smaller N for the O(N) zero-initialization of the hash table when
    the input is small, even if the pure Go golang/snappy algorithm cannot:
    "var table [tableSize]uint16" zeroes all tableSize elements.
    
    For this package, we don't have the match-C++-snappy goal, so we can use
    a different (constant) hash table size.
    
    This is a small win, in terms of throughput and output size, but it also
    enables us to re-use the (constant size) hash table between
    encodeBestSpeed calls, avoiding the cost of zero-initializing the hash
    table altogether. This will be implemented in follow-up commits.
    
    This package's benchmarks:
    name                    old speed      new speed      delta
    EncodeDigitsSpeed1e4-8  72.8MB/s ± 1%  73.5MB/s ± 1%  +0.86%  (p=0.000 n=10+10)
    EncodeDigitsSpeed1e5-8  77.5MB/s ± 1%  78.0MB/s ± 0%  +0.69%  (p=0.000 n=10+10)
    EncodeDigitsSpeed1e6-8  82.0MB/s ± 1%  82.7MB/s ± 1%  +0.85%   (p=0.000 n=10+9)
    EncodeTwainSpeed1e4-8   65.1MB/s ± 1%  65.6MB/s ± 0%  +0.78%   (p=0.000 n=10+9)
    EncodeTwainSpeed1e5-8   80.0MB/s ± 0%  80.6MB/s ± 1%  +0.66%   (p=0.000 n=9+10)
    EncodeTwainSpeed1e6-8   81.6MB/s ± 1%  82.1MB/s ± 1%  +0.55%  (p=0.017 n=10+10)
    
    Input size in bytes, output size (and time taken) before and after on
    some larger files:
    1073741824   57269781 (  3183ms)   57269781 (  3177ms) adresser.001
    1000000000  391052000 ( 11071ms)  391051996 ( 11067ms) enwik9
    1911399616  378679516 ( 13450ms)  378679514 ( 13079ms) gob-stream
    8558382592 3972329193 ( 99962ms) 3972329193 ( 91290ms) rawstudio-mint14.tar
     200000000  200015265 (   776ms)  200015265 (   774ms) sharnd.out
    
    Thanks to Klaus Post for the original suggestion on cl/21021.
    
    Change-Id: Ia4c63a8d1b92c67e1765ec5c3c8c69d289d9a6ce
    Reviewed-on: https://go-review.googlesource.com/22604
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 5edcff01343f36618dd6330438cf8b456bd914ef
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Apr 29 14:17:04 2016 +1000

    cmd/compile/internal/gc: bv.go cleanup
    
    Drive by gardening of bv.go.
    
    - Unexport the Bvec type, it is not used outside internal/gc.
      (machine translated with gofmt -r)
    - Removed unused constants and functions.
      (driven by cmd/unused)
    
    Change-Id: I3433758ad4e62439f802f4b0ed306e67336d9aba
    Reviewed-on: https://go-review.googlesource.com/22602
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 94e523cb522959f10b345bf27a32bf087c094108
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Apr 27 18:03:49 2016 -0400

    misc/cgo/testcarchive: fix C include path for darwin/arm
    
    After CL 22461, c-archive build on darwin/arm is by default compiled
    with -shared and installed in pkg/darwin_arm_shared.
    
    Fix build (2nd time...)
    
    Change-Id: Ia2bb09bb6e1ebc9bc74f7570dd80c81d05eaf744
    Reviewed-on: https://go-review.googlesource.com/22534
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d8b7bd6a1f89df1fbcf43fcaee72a94e291dcdb1
Author: Nigel Tao <nigeltao@golang.org>
Date:   Tue Apr 26 19:16:30 2016 +1000

    compress/flate: replace "Best Speed" with specialized version
    
    This encoding algorithm, which prioritizes speed over output size, is
    based on Snappy's LZ77-style encoder: github.com/golang/snappy
    
    This commit keeps the diff between this package's encodeBestSpeed
    function and and Snappy's encodeBlock function as small as possible (see
    the diff below). Follow-up commits will improve this package's
    performance and output size.
    
    This package's speed benchmarks:
    
    name                    old speed      new speed      delta
    EncodeDigitsSpeed1e4-8  40.7MB/s ± 0%  73.0MB/s ± 0%   +79.18%  (p=0.008 n=5+5)
    EncodeDigitsSpeed1e5-8  33.0MB/s ± 0%  77.3MB/s ± 1%  +134.04%  (p=0.008 n=5+5)
    EncodeDigitsSpeed1e6-8  32.1MB/s ± 0%  82.1MB/s ± 0%  +156.18%  (p=0.008 n=5+5)
    EncodeTwainSpeed1e4-8   42.1MB/s ± 0%  65.0MB/s ± 0%   +54.61%  (p=0.008 n=5+5)
    EncodeTwainSpeed1e5-8   46.3MB/s ± 0%  80.0MB/s ± 0%   +72.81%  (p=0.008 n=5+5)
    EncodeTwainSpeed1e6-8   47.3MB/s ± 0%  81.7MB/s ± 0%   +72.86%  (p=0.008 n=5+5)
    
    Here's the milliseconds taken, before and after this commit, to compress
    a number of test files:
    
    Go's src/compress/testdata files:
    
         4          1 e.txt
         8          4 Mark.Twain-Tom.Sawyer.txt
    
    github.com/golang/snappy's benchmark files:
    
         3          1 alice29.txt
        12          3 asyoulik.txt
         6          1 fireworks.jpeg
         1          1 geo.protodata
         1          0 html
         2          2 html_x_4
         6          3 kppkn.gtb
        11          4 lcet10.txt
         5          1 paper-100k.pdf
        14          6 plrabn12.txt
        17          6 urls.10K
    
    Larger files linked to from
    https://docs.google.com/spreadsheets/d/1VLxi-ac0BAtf735HyH3c1xRulbkYYUkFecKdLPH7NIQ/edit#gid=166102500
    
      2409       3182 adresser.001
     16757      11027 enwik9
     13764      12946 gob-stream
    153978      74317 rawstudio-mint14.tar
      4371        770 sharnd.out
    
    Output size is larger. In the table below, the first column is the input
    size, the second column is the output size prior to this commit, the
    third column is the output size after this commit.
    
        100003      47707      50006 e.txt
        387851     172707     182930 Mark.Twain-Tom.Sawyer.txt
        152089      62457      66705 alice29.txt
        125179      54503      57274 asyoulik.txt
        123093     122827     123108 fireworks.jpeg
        118588      18574      20558 geo.protodata
        102400      16601      17305 html
        409600      65506      70313 html_x_4
        184320      49007      50944 kppkn.gtb
        426754     166957     179355 lcet10.txt
        102400      82126      84937 paper-100k.pdf
        481861     218617     231988 plrabn12.txt
        702087     241774     258020 urls.10K
    1073741824   43074110   57269781 adresser.001
    1000000000  365772256  391052000 enwik9
    1911399616  340364558  378679516 gob-stream
    8558382592 3807229562 3972329193 rawstudio-mint14.tar
     200000000  200061040  200015265 sharnd.out
    
    The diff between github.com/golang/snappy's encodeBlock function and
    this commit's encodeBestSpeed function:
    
    1c1,7
    < func encodeBlock(dst, src []byte) (d int) {
    ---
    > func encodeBestSpeed(dst []token, src []byte) []token {
    > 	// This check isn't in the Snappy implementation, but there, the caller
    > 	// instead of the callee handles this case.
    > 	if len(src) < minNonLiteralBlockSize {
    > 		return emitLiteral(dst, src)
    > 	}
    >
    4c10
    < 	// and len(src) <= maxBlockSize and maxBlockSize == 65536.
    ---
    > 	// and len(src) <= maxStoreBlockSize and maxStoreBlockSize == 65535.
    65c71
    < 			if load32(src, s) == load32(src, candidate) {
    ---
    > 			if s-candidate < maxOffset && load32(src, s) == load32(src, candidate) {
    73c79
    < 		d += emitLiteral(dst[d:], src[nextEmit:s])
    ---
    > 		dst = emitLiteral(dst, src[nextEmit:s])
    90c96
    < 			// This is an inlined version of:
    ---
    > 			// This is an inlined version of Snappy's:
    93c99,103
    < 			for i := candidate + 4; s < len(src) && src[i] == src[s]; i, s = i+1, s+1 {
    ---
    > 			s1 := base + maxMatchLength
    > 			if s1 > len(src) {
    > 				s1 = len(src)
    > 			}
    > 			for i := candidate + 4; s < s1 && src[i] == src[s]; i, s = i+1, s+1 {
    96c106,107
    < 			d += emitCopy(dst[d:], base-candidate, s-base)
    ---
    > 			// matchToken is flate's equivalent of Snappy's emitCopy.
    > 			dst = append(dst, matchToken(uint32(s-base-3), uint32(base-candidate-minOffsetSize)))
    114c125
    < 			if uint32(x>>8) != load32(src, candidate) {
    ---
    > 			if s-candidate >= maxOffset || uint32(x>>8) != load32(src, candidate) {
    124c135
    < 		d += emitLiteral(dst[d:], src[nextEmit:])
    ---
    > 		dst = emitLiteral(dst, src[nextEmit:])
    126c137
    < 	return d
    ---
    > 	return dst
    
    This change is based on https://go-review.googlesource.com/#/c/21021/ by
    Klaus Post, but it is a separate changelist as cl/21021 seems to have
    stalled in code review, and the Go 1.7 feature freeze approaches.
    
    Golang-dev discussion:
    https://groups.google.com/d/topic/golang-dev/XYgHX9p8IOk/discussion and
    of course cl/21021.
    
    Change-Id: Ib662439417b3bd0b61c2977c12c658db3e44d164
    Reviewed-on: https://go-review.googlesource.com/22370
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 3e2462387f39db99a9a2b551c444c22fae460949
Author: Austin Clements <austin@google.com>
Date:   Thu Apr 28 11:21:01 2016 -0400

    [dev.garbage] runtime: eliminate mspan.start
    
    This converts all remaining uses of mspan.start to instead use
    mspan.base(). In many cases, this actually reduces the complexity of
    the code.
    
    Change-Id: If113840e00d3345a6cf979637f6a152e6344aee7
    Reviewed-on: https://go-review.googlesource.com/22590
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>

commit b7adc41fbacac446c1daf0cb282cb2a921d4a15b
Author: Austin Clements <austin@google.com>
Date:   Thu Apr 28 10:59:00 2016 -0400

    [dev.garbage] runtime: use s.base() everywhere it makes sense
    
    Currently we have lots of (s.start << _PageShift) and variants. We now
    have an s.base() function that returns this. It's faster and more
    readable, so use it.
    
    Change-Id: I888060a9dae15ea75ca8cc1c2b31c905e71b452b
    Reviewed-on: https://go-review.googlesource.com/22559
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>

commit 2e8b74b69574e969b5565e69cb54d39064b2dba1
Author: Austin Clements <austin@google.com>
Date:   Thu Apr 28 11:19:53 2016 -0400

    [dev.garbage] runtime: document sysAlloc
    
    In particular, it always returns an aligned pointer.
    
    Change-Id: I763789a539a4bfd8b0efb36a39a80be1a479d3e2
    Reviewed-on: https://go-review.googlesource.com/22558
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 15744c92de5e6a2295bfbae2126b19c124bbb46a
Author: Austin Clements <austin@google.com>
Date:   Thu Apr 28 10:53:25 2016 -0400

    [dev.garbage] runtime: remove unused head/end arguments from freeSpan
    
    These used to be used for the list of newly freed objects, but that's
    no longer a thing.
    
    Change-Id: I5a4503137b74ec0eae5372ca271b1aa0b32df074
    Reviewed-on: https://go-review.googlesource.com/22557
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c884f6594a594d2d18c4d21106592bd3cdfcbe9b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Apr 28 22:04:30 2016 -0500

    context: produce a nicer panic message for a nil WithValue key
    
    Change-Id: I2e8ae403622ba7131cadaba506100d79613183f1
    Reviewed-on: https://go-review.googlesource.com/22601
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 694846a548f23cebb9b913999fa4fa6756e2c545
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu Apr 28 15:19:11 2016 +1000

    debug/pe: .bss section must contain only zeros
    
    .bss section has no data stored in PE file. But when .bss section data
    is used by the linker it is assumed that its every byte is set to zero.
    (*Section).Data returns garbage at this moment. Change (*Section).Data
    so it returns slice filled with 0s.
    
    Updates #15345
    
    Change-Id: I1fa5138244a9447e1d59dec24178b1dd0fd4c5d7
    Reviewed-on: https://go-review.googlesource.com/22544
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d954f9c4d19bbe459ad4a6de95af47349da1d40b
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Apr 28 14:24:19 2016 -0700

    test: added test case for (fixed) issue 15470
    
    Follow-up to https://golang.org/cl/22543.
    
    Change-Id: I873b4fa6616ac2aea8faada2fccd126233bbc07f
    Reviewed-on: https://go-review.googlesource.com/22583
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit af6aa0fd745d48c2db70712ebfe6833d30a9a85d
Author: Russ Cox <rsc@golang.org>
Date:   Mon Apr 25 11:54:24 2016 -0400

    cmd/go, go/build: add support for binary-only packages
    
    See https://golang.org/design/2775-binary-only-packages for design.
    
    Fixes #2775.
    
    Change-Id: I33e74eebffadc14d3340bba96083af0dec5172d5
    Reviewed-on: https://go-review.googlesource.com/22433
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4618dd8704b868a095a98ee8beaf465578aaec30
Author: Nigel Tao <nigeltao@golang.org>
Date:   Thu Apr 28 17:19:31 2016 +1000

    image/gif: accept an out-of-bounds transparent color index.
    
    This is an error according to the spec, but Firefox and Google Chrome
    seem OK with this.
    
    Fixes #15059.
    
    Change-Id: I841cf44e96655e91a2481555f38fbd7055a32202
    Reviewed-on: https://go-review.googlesource.com/22546
    Reviewed-by: Rob Pike <r@golang.org>

commit 2fb75ea6c65d03c3fda89c8e954712a2fa97b052
Author: Rick Hudson <rlh@golang.org>
Date:   Thu Mar 31 10:45:36 2016 -0400

    [dev.garbage] runtime: use sys.Ctz64 intrinsic
    
    Our compilers now provides instrinsics including
    sys.Ctz64 that support CTZ (count trailing zero)
    instructions. This CL replaces the Go versions
    of CTZ with the compiler intrinsic.
    
    Count trailing zeros CTZ finds the least
    significant 1 in a word and returns the number
    of less significant 0s in the word.
    
    Allocation uses the bitmap created by the garbage
    collector to locate an unmarked object. The logic
    takes a word of the bitmap, complements, and then
    caches it. It then uses CTZ to locate an available
    unmarked object. It then shifts marked bits out of
    the bitmap word preparing it for the next search.
    Once all the unmarked objects are used in the
    cached work the bitmap gets another word and
    repeats the process.
    
    Change-Id: Id2fc42d1d4b9893efaa2e1bd01896985b7e42f82
    Reviewed-on: https://go-review.googlesource.com/21366
    Reviewed-by: Austin Clements <austin@google.com>

commit 2063d5d903718962de58a86a692626fe89919a4d
Author: Rick Hudson <rlh@golang.org>
Date:   Mon Mar 14 12:17:48 2016 -0400

    [dev.garbage] runtime: restructure alloc and mark bits
    
    Two changes are included here that are dependent on the other.
    The first is that allocBits and gcamrkBits are changed to
    a *uint8 which points to the first byte of that span's
    mark and alloc bits. Several places were altered to
    perform pointer arithmetic to locate the byte corresponding
    to an object in the span. The actual bit corresponding
    to an object is indexed in the byte by using the lower three
    bits of the objects index.
    
    The second change avoids the redundant calculation of an
    object's index. The index is returned from heapBitsForObject
    and then used by the functions indexing allocBits
    and gcmarkBits.
    
    Finally we no longer allocate the gc bits in the span
    structures. Instead we use an arena based allocation scheme
    that allows for a more compact bit map as well as recycling
    and bulk clearing of the mark bits.
    
    Change-Id: If4d04b2021c092ec39a4caef5937a8182c64dfef
    Reviewed-on: https://go-review.googlesource.com/20705
    Reviewed-by: Austin Clements <austin@google.com>

commit ac0ee77d630c4a692b02cad630a61e974b0c52ce
Author: Nigel Tao <nigeltao@golang.org>
Date:   Thu Apr 28 19:01:48 2016 +1000

    image/gif: be stricter on parsing graphic control extensions.
    
    See Section 23. Graphic Control Extension of the spec:
    https://www.w3.org/Graphics/GIF/spec-gif89a.txt
    
    Change-Id: Ie78b4ff4aa97e1b332ade67ae4fa25f7c0770610
    Reviewed-on: https://go-review.googlesource.com/22547
    Reviewed-by: Rob Pike <r@golang.org>

commit cb97fd7741fc8bfa257bb020dab756a14c420daf
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Apr 29 10:37:37 2016 +1200

    cmd/link: fix -no-pie / -race check
    
    golang.org/cl/22453 was supposed to pass -no-pie to the linker when linking a
    race-enabled binary if the host toolchain supports it. But I bungled the
    supported check as I forgot to pass -c to the host compiler so it tried to
    compile a 0 byte .c file into an executable, which will never work. Fix it to
    pass -c as it should have all along.
    
    Change-Id: I4801345c7a29cb18d5f22cec5337ce535f92135d
    Reviewed-on: https://go-review.googlesource.com/22587
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 6ed79fbd1abf018b018088c6a1699cd34ee6d56c
Author: Keith Randall <khr@golang.org>
Date:   Thu Apr 28 15:04:10 2016 -0700

    cmd/compile: remove BlockDead state
    
    It is unused, remove the clutter.
    
    Change-Id: I51a44326b125ef79241459c463441f76a289cc08
    Reviewed-on: https://go-review.googlesource.com/22586
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3cb090f93c4b0d9ebdf111efb0c5383d8ca97bd2
Author: David Symonds <dsymonds@golang.org>
Date:   Fri Apr 29 08:43:53 2016 +1000

    lib/time: update to IANA release 2016d (Apr 2016).
    
    Change-Id: I46d9ea31cf5836d054a9ce22af4dd1742a418a07
    Reviewed-on: https://go-review.googlesource.com/22588
    Run-TryBot: David Symonds <dsymonds@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit be730b49ca41e8a7e98a21cf61bb6c9c7fc7857e
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Apr 28 09:35:19 2016 +0900

    runtime: drop _SigUnblock for SIGSYS on Linux
    
    The _SigUnblock flag was appended to SIGSYS slot of runtime signal table
    for Linux in https://go-review.googlesource.com/22202, but there is
    still no concrete opinion on whether SIGSYS must be an unblocked signal
    for runtime.
    
    This change removes _SigUnblock flag from SIGSYS on Linux for
    consistency in runtime signal handling and adds a reference to #15204 to
    runtime signal table for FreeBSD.
    
    Updates #15204.
    
    Change-Id: I42992b1d852c2ab5dd37d6dbb481dba46929f665
    Reviewed-on: https://go-review.googlesource.com/22537
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit bec0863b53777f313396a10f0bc1349139d1009e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Apr 28 13:44:15 2016 -0700

    net: remove unneeded tags from dnsRR structs
    
    DNS packing and unpacking uses hand-coded struct walking functions
    rather than reflection, so these tags are unneeded and just contribute
    to their runtime reflect metadata size.
    
    Change-Id: I2db09d5159912bcbc3b482cbf23a50fa8fa807fa
    Reviewed-on: https://go-review.googlesource.com/22594
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c231dd21e1fe992e196d4a36f9e990523b9cc45a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Apr 28 13:26:14 2016 -0700

    net: remove internal support for obsolete DNS record types
    
    There are no real world use cases for HINFO, MINFO, MB, MG, or MR
    records, and package net's exposed APIs don't provide any way to
    access them even if there were. If a use ever does show up, we can
    revive them. In the mean time, this is just effectively-dead code that
    sticks around because of rr_mk.
    
    Change-Id: I6c188b5ee32f3b3a04588b79a0ee9c2e3e725ccc
    Reviewed-on: https://go-review.googlesource.com/22593
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1518d431321100cd9f0e18d740da7c835ba438dd
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 16 11:57:06 2016 -0700

    net/http, net/http/httptrace: new package for tracing HTTP client requests
    
    Updates #12580
    
    Change-Id: I9f9578148ef2b48dffede1007317032d39f6af55
    Reviewed-on: https://go-review.googlesource.com/22191
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Tom Bergan <tombergan@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1b591dfb1f071d978448966e979e40b1f265c1a5
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Apr 28 15:16:46 2016 -0500

    os/exec: fix variable shadow, don't leak goroutine
    
    Goroutine leak checking is still too tedious, so untested.
    
    See #6705 which is my fault for forgetting to mail out.
    
    Change-Id: I899fb311c9d4229ff1dbd3f54fe307805e17efee
    Reviewed-on: https://go-review.googlesource.com/22581
    Reviewed-by: Ahmed W. <oneofone@gmail.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6c11e2710e96171e7c202940bf2b14aa859f5ca2
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Apr 28 12:43:12 2016 -0700

    cmd/compile: use delta encoding for filenames in export data position info
    
    This reduces the export data size significantly (15%-25%) for some packages,
    especially where the paths are very long or if there are many files involved.
    Slight (2%) reduction on average, with virtually no increases in export data
    size.
    
    Selected export data sizes for packages with |delta %| > 3%:
    
                         package   before    after   delta     %
    
            cmd/asm/internal/arch   11647    11088    -559   -4%
       cmd/compile/internal/amd64     838      600    -238  -27%
         cmd/compile/internal/arm    7323     6793    -530   -6%
       cmd/compile/internal/arm64   19948    18971    -977   -4%
         cmd/compile/internal/big    9043     8548    -495   -4%
      cmd/compile/internal/mips64     645      482    -163  -24%
       cmd/compile/internal/ppc64     695      497    -198  -27%
       cmd/compile/internal/s390x     553      433    -120  -21%
         cmd/compile/internal/x86     744      555    -189  -24%
                         cmd/dist     145      121     -24  -16%
             cmd/internal/objfile   17359    16474    -885   -4%
       cmd/internal/pprof/symbolz    8346     7941    -405   -4%
          cmd/link/internal/amd64   11178    10604    -574   -4%
            cmd/link/internal/arm     204      171     -33  -15%
          cmd/link/internal/arm64     210      175     -35  -16%
         cmd/link/internal/mips64     213      177     -36  -16%
          cmd/link/internal/ppc64     211      176     -35  -16%
          cmd/link/internal/s390x     210      175     -35  -16%
            cmd/link/internal/x86     203      170     -33  -15%
                        cmd/trace     782      744     -38   -4%
                     compress/lzw     402      383     -19   -4%
                       crypto/aes     311      262     -49  -15%
                    crypto/cipher    1138      959    -179  -15%
                       crypto/des     315      288     -27   -8%
                  crypto/elliptic    6063     5746    -317   -4%
                       crypto/rc4     317      295     -22   -6%
                    crypto/sha256     348      312     -36   -9%
                    crypto/sha512     487      451     -36   -6%
                           go/doc    3871     3649    -222   -5%
        go/internal/gccgoimporter    2063     1949    -114   -5%
           go/internal/gcimporter    3253     3096    -157   -4%
                             math    4343     3572    -771  -17%
                       math/cmplx    1580     1274    -306  -18%
                        math/rand     982      926     -56   -5%
            net/internal/socktest    2159     2049    -110   -4%
                          os/exec    7928     7492    -436   -4%
                        os/signal     237      208     -29  -11%
                          os/user     717      682     -35   -4%
          runtime/internal/atomic     728      693     -35   -4%
             runtime/internal/sys    2287     2107    -180   -7%
                             sync    1306     1214     -92   -6%
    
                     all packages 1509255  1465507  -43748   -2%
    
    Change-Id: I98a11521b552166b7f47f2039a29f106748bf5d4
    Reviewed-on: https://go-review.googlesource.com/22580
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit f04eb356732c205f9e8c2f38a4f5fdb9def991b2
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Apr 28 10:58:55 2016 -0700

    cmd/compile: remove unused Bputname function
    
    Change-Id: Icecbf9bae8c39670d1ceef62dd94b36e90b27b04
    Reviewed-on: https://go-review.googlesource.com/22570
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3c8ef0e0c9c26f15926a396688b0fe8acd4e3dcf
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Apr 24 14:37:14 2016 -0400

    cmd/compile: allow 64-bit multiplication with immediates on s390x
    
    MGHI (16-bit signed immediate) is now used where possible for both
    MULLW and MULLD. MGHI is 2-bytes shorter than MSGFI.
    
    Change-Id: I5d0648934f28b3403b1126913fd703d8f62b9e9f
    Reviewed-on: https://go-review.googlesource.com/22398
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5fe1b35ed214a8ece13449f5788dd9f5c927379f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Apr 28 11:31:59 2016 -0700

    net: ensure dnsConfig search list is rooted
    
    Avoids some extra work and string concatenation at query time.
    
    benchmark                                      old allocs     new allocs     delta
    BenchmarkGoLookupIP-32                         154            150            -2.60%
    BenchmarkGoLookupIPNoSuchHost-32               446            442            -0.90%
    BenchmarkGoLookupIPWithBrokenNameServer-32     564            568            +0.71%
    
    benchmark                                      old bytes     new bytes     delta
    BenchmarkGoLookupIP-32                         10824         10704         -1.11%
    BenchmarkGoLookupIPNoSuchHost-32               43140         42992         -0.34%
    BenchmarkGoLookupIPWithBrokenNameServer-32     46616         46680         +0.14%
    
    BenchmarkGoLookupIPWithBrokenNameServer's regression appears to be
    because it's actually only performing 1 LookupIP call, so the extra
    work done parsing the DNS config file doesn't amortize as well as for
    BenchmarkGoLookupIP or BenchmarkGoLOokupIPNoSuchHost, which perform
    2000+ LookupIP calls per run.
    
    Update #15473.
    
    Change-Id: I98c8072f2f39e2f2ccd6c55e9e9bd309f5ad68f8
    Reviewed-on: https://go-review.googlesource.com/22571
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4d9bda51ff91f79944a12106b77315c9414b851a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Apr 28 11:15:44 2016 -0700

    net: append ":53" to DNS servers when reading resolv.conf
    
    Avoids generating some redundant garbage from re-concatenating the
    same string for every DNS query.
    
    benchmark                                      old allocs     new allocs     delta
    BenchmarkGoLookupIP-32                         156            154            -1.28%
    BenchmarkGoLookupIPNoSuchHost-32               456            446            -2.19%
    BenchmarkGoLookupIPWithBrokenNameServer-32     577            564            -2.25%
    
    benchmark                                      old bytes     new bytes     delta
    BenchmarkGoLookupIP-32                         10873         10824         -0.45%
    BenchmarkGoLookupIPNoSuchHost-32               43303         43140         -0.38%
    BenchmarkGoLookupIPWithBrokenNameServer-32     46824         46616         -0.44%
    
    Update #15473.
    
    Change-Id: I3b0173dfedf31bd08eaea1069968b416850864a1
    Reviewed-on: https://go-review.googlesource.com/22556
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2cc27a7de9e7d14cb6702153688d02746c6a49ea
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Apr 28 11:53:58 2016 -0500

    os/exec: add Cmd.RunContext and Cmd.WaitContext
    
    Updates #14660
    
    Change-Id: Ifa5c97ba327ad7ceea0a9a252e3dbd9d079dae54
    Reviewed-on: https://go-review.googlesource.com/22529
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit af125a5193c75dd59307fcf1b26d885010ce8bfd
Author: Adam Langley <agl@golang.org>
Date:   Tue Apr 26 10:45:35 2016 -0700

    crypto/tls: allow renegotiation to be handled by a client.
    
    This change adds Config.Renegotiation which controls whether a TLS
    client will accept renegotiation requests from a server. This is used,
    for example, by some web servers that wish to “add” a client certificate
    to an HTTPS connection.
    
    This is disabled by default because it significantly complicates the
    state machine.
    
    Originally, handshakeMutex was taken before locking either Conn.in or
    Conn.out. However, if renegotiation is permitted then a handshake may
    be triggered during a Read() call. If Conn.in were unlocked before
    taking handshakeMutex then a concurrent Read() call could see an
    intermediate state and trigger an error. Thus handshakeMutex is now
    locked after Conn.in and the handshake functions assume that Conn.in is
    locked for the duration of the handshake.
    
    Additionally, handshakeMutex used to protect Conn.out also. With the
    possibility of renegotiation that's no longer viable and so
    writeRecordLocked has been split off.
    
    Fixes #5742.
    
    Change-Id: I935914db1f185d507ff39bba8274c148d756a1c8
    Reviewed-on: https://go-review.googlesource.com/22475
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit d610d304f86021cc5f388b8f02d99bc73fca0d9b
Author: Keith Randall <khr@golang.org>
Date:   Wed Apr 27 16:58:50 2016 -0700

    cmd/compile: reorg copyelim to avoid O(n^2) problem
    
    Make sure we don't do O(n^2) work to eliminate a chain
    of n copies.
    
    benchmark                     old ns/op       new ns/op     delta
    BenchmarkCopyElim1-8          1418            1406          -0.85%
    BenchmarkCopyElim10-8         5289            5162          -2.40%
    BenchmarkCopyElim100-8        52618           41684         -20.78%
    BenchmarkCopyElim1000-8       2473878         424339        -82.85%
    BenchmarkCopyElim10000-8      269373954       6367971       -97.64%
    BenchmarkCopyElim100000-8     31272781165     104357244     -99.67%
    
    Change-Id: I680f906f70f2ee1a8615cb1046bc510c77d59284
    Reviewed-on: https://go-review.googlesource.com/22535
    Reviewed-by: Alexandru Moșoi <alexandru@mosoi.ro>

commit 5ec87ba554c2a83cdc188724f815e53fede91b66
Author: David Chase <drchase@google.com>
Date:   Thu Apr 28 10:46:08 2016 -0400

    cmd/compile: fix.gc.Type.cmp for map.notBucket cmp map.Bucket
    
    Comparison of certain map types could fail to be antisymmetric.
    This corrects that.
    
    Change-Id: I88c6256053ce29950ced4ba4d538e241ee8591fe
    Reviewed-on: https://go-review.googlesource.com/22552
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: jcd . <jcd@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 23aeb34df172b17b7bfaa85fb59ca64bef9073bb
Merge: 1354b32 d3c79d3
Author: Rick Hudson <rlh@golang.org>
Date:   Wed Apr 27 18:19:16 2016 -0400

    [dev.garbage] Merge remote-tracking branch 'origin/master' into HEAD
    
    Change-Id: I282fd9ce9db435dfd35e882a9502ab1abc185297

commit 1354b32cd70f2702381764fd595dd2faa996840c
Author: Rick Hudson <rlh@golang.org>
Date:   Mon Mar 14 12:17:48 2016 -0400

    [dev.garbage] runtime: add gc work buffer tryGet and put fast paths
    
    The complexity of the GC work buffers put and tryGet
    prevented them from being inlined. This CL simplifies
    the fast path thus enabling inlining. If the fast
    path does not succeed the previous put and tryGet
    functions are called.
    
    Change-Id: I6da6495d0dadf42bd0377c110b502274cc01acf5
    Reviewed-on: https://go-review.googlesource.com/20704
    Reviewed-by: Austin Clements <austin@google.com>

commit f8d0d4fd59b6cb6f875eac7753f036b10a28f995
Author: Rick Hudson <rlh@golang.org>
Date:   Mon Mar 14 12:02:02 2016 -0400

    [dev.garbage] runtime: cleanup and optimize span.base()
    
    Prior to this CL the base of a span was calculated in various
    places using shifts or calls to base(). This CL now
    always calls base() which has been optimized to calculate the
    base of the span when the span is initialized and store that
    value in the span structure.
    
    Change-Id: I661f2bfa21e3748a249cdf049ef9062db6e78100
    Reviewed-on: https://go-review.googlesource.com/20703
    Reviewed-by: Austin Clements <austin@google.com>

commit 8dda1c4c08adf8b2107dec1c0d70d24443269ccd
Author: Rick Hudson <rlh@golang.org>
Date:   Wed Mar 2 12:15:02 2016 -0500

    [dev.garbage] runtime: remove heapBitsSweepSpan
    
    Prior to this CL the sweep phase was responsible for locating
    all objects that were about to be freed and calling a function
    to process the object. This was done by the function
    heapBitsSweepSpan. Part of processing included calls to
    tracefree and msanfree as well as counting how many objects
    were freed.
    
    The calls to tracefree and msanfree have been moved into the
    gcmalloc routine and called when the object is about to be
    reallocated. The counting of free objects has been optimized
    using an array based popcnt algorithm and if all the objects
    in a span are free then span is freed.
    
    Similarly the code to locate the next free object has been
    optimized to use an array based ctz (count trailing zero).
    Various hot paths in the allocation logic have been optimized.
    
    At this point the garbage benchmark is within 3% of the 1.6
    release.
    
    Change-Id: I00643c442e2ada1685c010c3447e4ea8537d2dfa
    Reviewed-on: https://go-review.googlesource.com/20201
    Reviewed-by: Austin Clements <austin@google.com>

commit 4093481523b1e064e998d5d586276db45f4d11a7
Author: Rick Hudson <rlh@golang.org>
Date:   Wed Feb 24 14:36:30 2016 -0500

    [dev.garbage] runtime: add bit and cache ctz64 (count trailing zero)
    
    Add to each span a 64 bit cache (allocCache) of the allocBits
    at freeindex. allocCache is shifted such that the lowest bit
    corresponds to the bit freeindex. allocBits uses a 0 to
    indicate an object is free, on the other hand allocCache
    uses a 1 to indicate an object is free. This facilitates
    ctz64 (count trailing zero) which counts the number of 0s
    trailing the least significant 1. This is also the index of
    the least significant 1.
    
    Each span maintains a freeindex indicating the boundary
    between allocated objects and unallocated objects. allocCache
    is shifted as freeindex is incremented such that the low bit
    in allocCache corresponds to the bit a freeindex in the
    allocBits array.
    
    Currently ctz64 is written in Go using a for loop so it is
    not very efficient. Use of the hardware instruction will
    follow. With this in mind comparisons of the garbage
    benchmark are as follows.
    
    1.6 release        2.8 seconds
    dev:garbage branch 3.1 seconds.
    
    Profiling shows the go implementation of ctz64 takes up
    1% of the total time.
    
    Change-Id: If084ed9c3b1eda9f3c6ab2e794625cb870b8167f
    Reviewed-on: https://go-review.googlesource.com/20200
    Reviewed-by: Austin Clements <austin@google.com>

commit 44fe90d0b393c961e3fb1b4c37e93ce268da46bc
Author: Rick Hudson <rlh@golang.org>
Date:   Wed Feb 17 11:27:52 2016 -0500

    [dev.garbage] runtime: logic that uses count trailing zero (ctz)
    
    Most (all?) processors that Go supports supply a hardware
    instruction that takes a byte and returns the number
    of zeros trailing the first 1 encountered, or 8
    if no ones are found. This is the index within the
    byte of the first 1 encountered. CTZ should improve the
    performance of the nextFreeIndex function.
    
    Since nextFreeIndex wants the next unmarked (0) bit
    a bit-wise complement is needed before calling ctz.
    Furthermore unmarked bits associated with previously
    allocated objects need to be ignored. Instead of writing
    a 1 as we allocate the code masks all bits less than the
    freeindex after loading the byte.
    
    While this CL does not actual execute a CTZ instruction
    it supplies a ctz function with the appropiate signature
    along with the logic to execute it.
    
    Change-Id: I5c55ce0ed48ca22c21c4dd9f969b0819b4eadaa7
    Reviewed-on: https://go-review.googlesource.com/20169
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit e4ac2d4acc8cb44df2107e3fa1067755feaaa005
Author: Rick Hudson <rlh@golang.org>
Date:   Tue Feb 16 17:16:43 2016 -0500

    [dev.garbage] runtime: replace ref with allocCount
    
    This is a renaming of the field ref to the
    more appropriate allocCount. The field
    holds the number of objects in the span
    that are currently allocated. Some throws
    strings were adjusted to more accurately
    convey the meaning of allocCount.
    
    Change-Id: I10daf44e3e9cc24a10912638c7de3c1984ef8efe
    Reviewed-on: https://go-review.googlesource.com/19518
    Reviewed-by: Austin Clements <austin@google.com>

commit 3479b065d43f2990ac12e7b00ddff6f63a876ca9
Author: Rick Hudson <rlh@golang.org>
Date:   Thu Feb 11 13:57:58 2016 -0500

    [dev.garbage] runtime: allocate directly from GC mark bits
    
    Instead of building a freelist from the mark bits generated
    by the GC this CL allocates directly from the mark bits.
    
    The approach moves the mark bits from the pointer/no pointer
    heap structures into their own per span data structures. The
    mark/allocation vectors consist of a single mark bit per
    object. Two vectors are maintained, one for allocation and
    one for the GC's mark phase. During the GC cycle's sweep
    phase the interpretation of the vectors is swapped. The
    mark vector becomes the allocation vector and the old
    allocation vector is cleared and becomes the mark vector that
    the next GC cycle will use.
    
    Marked entries in the allocation vector indicate that the
    object is not free. Each allocation vector maintains a boundary
    between areas of the span already allocated from and areas
    not yet allocated from. As objects are allocated this boundary
    is moved until it reaches the end of the span. At this point
    further allocations will be done from another span.
    
    Since we no longer sweep a span inspecting each freed object
    the responsibility for maintaining pointer/scalar bits in
    the heapBitMap containing is now the responsibility of the
    the routines doing the actual allocation.
    
    This CL is functionally complete and ready for performance
    tuning.
    
    Change-Id: I336e0fc21eef1066e0b68c7067cc71b9f3d50e04
    Reviewed-on: https://go-review.googlesource.com/19470
    Reviewed-by: Austin Clements <austin@google.com>

commit dc65a82eff0a3af5a26f6c6d31c53bdac9b31168
Author: Rick Hudson <rlh@golang.org>
Date:   Tue Feb 9 09:38:44 2016 -0500

    [dev.garbage] runtime: mark/allocation helper functions
    
    The gcmarkBits is a bit vector used by the GC to mark
    reachable objects. Once a GC cycle is complete the gcmarkBits
    swap places with the allocBits. allocBits is then used directly
    by malloc to locate free objects, thus avoiding the
    construction of a linked free list. This CL introduces a set
    of helper functions for manipulating gcmarkBits and allocBits
    that will be used by later CLs to realize the actual
    algorithm. Minimal attempts have been made to optimize these
    helper routines.
    
    Change-Id: I55ad6240ca32cd456e8ed4973c6970b3b882dd34
    Reviewed-on: https://go-review.googlesource.com/19420
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Rick Hudson <rlh@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e1c4e9a754833e169a41ea98a49c3712513879ab
Author: Rick Hudson <rlh@golang.org>
Date:   Mon Feb 8 12:36:23 2016 -0500

    [dev.garbage] runtime: refactor next free object
    
    In preparation for changing how the next free object is chosen
    refactor and consolidate code into a single function.
    
    Change-Id: I6836cd88ed7cbf0b2df87abd7c1c3b9fabc1cbd8
    Reviewed-on: https://go-review.googlesource.com/19317
    Reviewed-by: Austin Clements <austin@google.com>

commit aed861038f876643a67c2297b384b6be140c46c1
Author: Rick Hudson <rlh@golang.org>
Date:   Mon Feb 8 09:53:14 2016 -0500

    [dev.garbage] runtime: add stackfreelist
    
    The freelist for normal objects and the freelist
    for stacks share the same mspan field for holding
    the list head but are operated on by different code
    sequences. This overloading complicates the use of bit
    vectors for allocation of normal objects. This change
    refactors the use of the stackfreelist out from the
    use of freelist.
    
    Change-Id: I5b155b5b8a1fcd8e24c12ee1eb0800ad9b6b4fa0
    Reviewed-on: https://go-review.googlesource.com/19315
    Reviewed-by: Austin Clements <austin@google.com>

commit 2ac8bdc52ae1ea0418df465de3f1ef36f49e2274
Author: Rick Hudson <rlh@golang.org>
Date:   Thu Feb 4 11:41:48 2016 -0500

    [dev.garbage] runtime: bitmap allocation data structs
    
    The bitmap allocation data structure prototypes. Before
    this is released these underlying data structures need
    to be more performant but the signatures of helper
    functions utilizing these structures will remain stable.
    
    Change-Id: I5ace12f2fb512a7038a52bbde2bfb7e98783bcbe
    Reviewed-on: https://go-review.googlesource.com/19221
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6d6e16001bb04d4bf60c6d14d1f64684a043ef6c
Merge: 22972d2 5e1b7bd
Author: Austin Clements <austin@google.com>
Date:   Tue Apr 5 16:37:29 2016 -0400

    [dev.garbage] Merge branch 'master' into dev.garbage
    
    Change-Id: I47ac4112befc07d3674d7a88827227199edd93b4

commit 22972d2207424edc481b7c127788f573a726dfe7
Merge: c67f09a 0f86d1e
Author: Austin Clements <austin@google.com>
Date:   Mon Mar 28 13:33:44 2016 -0400

    [dev.garbage] Merge remote-tracking branch 'origin/master' into dev.garbage
    
    Change-Id: Icb6811a9eb08fbde297d256db9f135a4e85e7cd4

commit c67f09ac5cf3c5a6af67d0e302a1985ade44be1f
Merge: 132d7ae bbd3ffb
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 9 13:16:50 2016 -0500

    [dev.garbage] Merge branch 'master' into dev.garbage
    
    Change-Id: Iede021da8fdb7ac87f1e0c495b7401e50a5b0a83

commit 132d7ae99e2996349515a41832194a0d5c15725f
Merge: 29257f5 0d1a98e
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 2 11:05:13 2016 -0500

    [dev.garbage] Merge branch 'master' into dev.garbage
    
    Change-Id: I8327de2ac5eeb56d7f0371776a0d9131e3204f12

commit 29257f5effa30a8a0cc2d79df8ff25593a861b2b
Merge: 830ce3f 91911e3
Author: Rick Hudson <rlh@golang.org>
Date:   Wed Feb 3 18:14:46 2016 -0500

    Merge remote-tracking branch 'origin/master' into toc
    
    Change-Id: Iba5d3ce5955b5e95c3c40b4408727e28ade67c61

commit 830ce3f1ed2771a94ed768816304ce2dd1017da0
Author: Russ Cox <rsc@golang.org>
Date:   Mon Jan 11 21:10:46 2016 -0500

    [dev.garbage] dev.garbage: create new dev.garbage branch
    
    This is for a GC experiment that may or may not go anywhere.
    
    Change-Id: I46a4535cc768ce8bbe33c72961f1fa87658493f7
    Reviewed-on: https://go-review.googlesource.com/18534
    Reviewed-by: Rick Hudson <rlh@golang.org>
```
