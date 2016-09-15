# May 26, 2016 Report

Number of commits: 77

## Compilation time

* github.com/coreos/etcd: from 14.046480989s to 13.919707458s, -0.90%
* github.com/boltdb/bolt/cmd/bolt: from 544.857861ms to 554.538169ms, +1.78%
* github.com/gogits/gogs: from 12.507480694s to 12.613879282s, +0.85%
* github.com/spf13/hugo: from 6.476413014s to 6.499049693s, +0.35%
* github.com/influxdata/influxdb/cmd/influxd: from 6.416976896s to 6.390121149s, -0.42%
* github.com/nsqio/nsq/apps/nsqd: from 2.123233099s to 2.145951074s, +1.07%
* github.com/mholt/caddy: from 4.56006107s to 4.579368543s, +0.42%

## Binary size:

* github.com/coreos/etcd: from 22043872 to 22073312, +0.13%
* github.com/boltdb/bolt/cmd/bolt: from 2581853 to 2581982, ~
* github.com/gogits/gogs: from 22852110 to 22883091, +0.14%
* github.com/spf13/hugo: from 14763758 to 14790476, +0.18%
* github.com/influxdata/influxdb/cmd/influxd: from 15859285 to 15890381, +0.20%
* github.com/nsqio/nsq/apps/nsqd: from 9784401 to 9811237, +0.27%
* github.com/mholt/caddy: from 12747269 to 12769891, +0.18%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               185           186           +0.54%
BenchmarkMsgpUnmarshal-4             390           391           +0.26%
BenchmarkEasyJsonMarshal-4           1438          1409          -2.02%
BenchmarkEasyJsonUnmarshal-4         1476          1488          +0.81%
BenchmarkFlatBuffersMarshal-4        347           343           -1.15%
BenchmarkFlatBuffersUnmarshal-4      287           286           -0.35%
BenchmarkGogoprotobufMarshal-4       164           163           -0.61%
BenchmarkGogoprotobufUnmarshal-4     244           243           -0.41% 
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

* [context: make DeadlineExceeded have a Timeout method](https://github.com/golang/go/commit/dc4427f3727804ded270bc6a7a8066ccb3c151d0)
* [syscall: add Unshare flags to SysProcAttr on Linux](https://github.com/golang/go/commit/8527b8ef9b00c72b1a8e30e5917c7bdd3c0e79ef)
* [os/exec: remove Cmd.RunContext and Cmd.WaitContext, add CommandContext](https://github.com/golang/go/commit/4cad610401edc11fe921205438a7b3ab4faa3982)
* [io: remove SizedReaderAt](524956f8b976be2b7be829a2d0d87c2951932ac6)
* [net/http: add missing HTTP status codes](https://github.com/golang/go/commit/b9ec0024fbc18dd94eff7240afd82fac6b4d8fdc)

## GIT Log

```
commit 56e5e0b69c92c9157c7db39112e27a4b5c026b48
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed May 25 13:19:11 2016 -0400

    runtime: tell race detector about reflectOffs.lock
    
    Fixes #15832
    
    Change-Id: I6f3f45e3c21edd0e093ecb1d8a067907863478f5
    Reviewed-on: https://go-review.googlesource.com/23441
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 6247ca2dbbdc13d6c80666119d182e119a2e7a5b
Author: Russ Cox <rsc@golang.org>
Date:   Thu May 26 09:59:29 2016 -0400

    cmd/dist: drop testcarchive on ppc64le
    
    It is timing out on the dashboard.
    (We enabled it as an experiment to see if it was still broken. Looks that way.)
    
    Change-Id: I425b7e54a2ab95b623ab7a15554b4173078f75e2
    Reviewed-on: https://go-review.googlesource.com/23480
    Reviewed-by: Russ Cox <rsc@golang.org>

commit b92f4238790c590168e7dae03165d75deb89fe41
Author: Austin Clements <austin@google.com>
Date:   Wed May 25 20:56:56 2016 -0400

    runtime: unwind BP in jmpdefer to match SP unwind
    
    The irregular calling convention for defers currently incorrectly
    manages the BP if frame pointers are enabled. Specifically, jmpdefer
    manipulates the SP as if its own caller, deferreturn, had returned.
    However, it does not manipulate the BP to match. As a result, when a
    BP-based traceback happens during a deferred function call, it unwinds
    to the function that performed the defer and then thinks that function
    called itself in an infinite regress.
    
    Fix this by making jmpdefer manipulate the BP as if deferreturn had
    actually returned.
    
    Fixes #12968.
    
    Updates #15840.
    
    Change-Id: Ic9cc7c863baeaf977883ed0c25a7e80e592cf066
    Reviewed-on: https://go-review.googlesource.com/23457
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit adff422779709c83db91700fdcc0a0bd5dee6a21
Author: Austin Clements <austin@google.com>
Date:   Wed May 25 22:59:19 2016 -0400

    cmd/link/internal/ld: fix DWARF offsets with GOEXPERIMENT=framepointer
    
    The offsets computed by the DWARF expressions for local variables
    currently don't account for the extra stack slot used by the frame
    pointer when GOEXPERIMENT=framepointer is enabled.
    
    Fix this by adding the extra stack slot to the offset.
    
    This fixes TestGdbPython with GOEXPERIMENT=framepointer.
    
    Updates #15840.
    
    Change-Id: I1b2ebb2750cd22266f4a89ec8d9e8bfa05fabd19
    Reviewed-on: https://go-review.googlesource.com/23458
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d9557523c2febc29ad3ecab8b1a4358abd709b30
Author: Russ Cox <rsc@golang.org>
Date:   Wed May 25 20:01:25 2016 -0400

    runtime: make framepointer mode safe for Windows
    
    A few other architectures have already defined a NOFRAME flag.
    Use it to disable frame pointer code on a few very low-level functions
    that must behave like Windows code.
    
    Makes the failing os/signal test pass on a Windows gomote.
    
    Change-Id: I982365f2c59a0aa302b4428c970846c61027cf3e
    Reviewed-on: https://go-review.googlesource.com/23456
    Reviewed-by: Austin Clements <austin@google.com>

commit 805eaeef33a52778ba6ee624389c2cbfe6896f6f
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Tue May 17 20:55:55 2016 +0300

    crypto/sha1: fix AVX2 variant on AMD64
    
    AVX2 variant reads next blocks while calculating current block.
    Avoid reading past the end of data, by switching back to original,
    for last blocks.
    
    Fixes #15617.
    
    Change-Id: I04fa2d83f1b47995117c77b4a3d403a7dff594d4
    Reviewed-on: https://go-review.googlesource.com/23138
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8a1dc3244725c2afd170025fc616df840b464a99
Author: Russ Cox <rsc@golang.org>
Date:   Wed May 25 14:54:21 2016 -0400

    runtime: add library startup support for ppc64le
    
    I have been running this patch inside Google against Go 1.6 for the last month.
    
    The new tests will probably break the builders but let's see
    exactly how they break.
    
    Change-Id: Ia65cf7d3faecffeeb4b06e9b80875c0e57d86d9e
    Reviewed-on: https://go-review.googlesource.com/23452
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 30282b091d6c80f4aa37d7c457fa288c3a181573
Author: Robert Griesemer <gri@golang.org>
Date:   Wed May 25 16:38:02 2016 -0700

    cmd/compile: correctly import labels, gotos, and fallthroughs
    
    The importer had several bugs with respect to labels and gotos:
    - it didn't create a new ONAME node for label names (label dcl,
      goto, continue, and break)
    - it overwrote the symbol for gotos with the dclstack
    - it didn't set the dclstack for labels
    
    In the process changed export format slightly to always assume
    a label name for labels and gotos, and never assume a label for
    fallthroughs.
    
    For fallthroughs and switch cases, now also set Xoffset like in
    the parser. (Not setting it, i.e., using 0 was ok since this is
    only used for verifying correct use of fallthroughs, which was
    checked already. But it's an extra level of verification of the
    import.)
    
    Fixes #15838.
    
    Change-Id: I3637f6314b8651c918df0c8cd70cd858c92bd483
    Reviewed-on: https://go-review.googlesource.com/23445
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b9ec0024fbc18dd94eff7240afd82fac6b4d8fdc
Author: Seth Vargo <sethvargo@gmail.com>
Date:   Thu May 12 16:26:27 2016 -0400

    net/http: add missing HTTP status codes
    
    This commit adds missing status codes:
    
    * 102 - Processing
    * 207 - Multi-Status
    * 208 - Already Reported
    * 226 - IM Used
    * 308 - Permanent Redirect
    * 422 - Unprocessable Entity
    * 423 - Locked
    * 424 - Failed Dependency
    * 426 - Upgrade Required
    * 506 - Variant Also Negotiates
    * 507 - Insufficient Storage
    * 508 - Loop Detected
    * 510 - Not Extended
    * 511 - Network Authentication Required
    
    Change-Id: Ife0e5b064f4b1e3542d2fd41abc9e7b1e410b644
    Reviewed-on: https://go-review.googlesource.com/23090
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a5d1a72a40b59db0d2f3f5d3fbb2ed60aafb7fdf
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu May 19 16:27:23 2016 -0700

    cmd/cgo, runtime, runtime/cgo: TSAN support for malloc
    
    Acquire and release the TSAN synchronization point when calling malloc,
    just as we do when calling any other C function. If we don't do this,
    TSAN will report false positive errors about races calling malloc and
    free.
    
    We used to have a special code path for malloc and free, going through
    the runtime functions cmalloc and cfree. The special code path for cfree
    was no longer used even before this CL. This CL stops using the special
    code path for malloc, because there is no place along that path where we
    could conditionally insert the TSAN synchronization. This CL removes
    the support for the special code path for both functions.
    
    Instead, cgo now automatically generates the malloc function as though
    it were referenced as C.malloc.  We need to automatically generate it
    even if C.malloc is not called, even if malloc and size_t are not
    declared, to support cgo-provided functions like C.CString.
    
    Change-Id: I829854ec0787a80f33fa0a8a0dc2ee1d617830e2
    Reviewed-on: https://go-review.googlesource.com/23260
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 10c8b2374f413ef6225555893cad5d2a530f77d5
Author: Russ Cox <rsc@golang.org>
Date:   Wed May 25 17:24:11 2016 -0400

    runtime: align C library startup calls on amd64
    
    This makes GOEXPERIMENT=framepointer, GOOS=darwin, and buildmode=carchive coexist.
    
    Change-Id: I9f6fb2f0f06f27df683e5b51f2fa55cd21872453
    Reviewed-on: https://go-review.googlesource.com/23454
    Reviewed-by: Austin Clements <austin@google.com>

commit 3be48b4dc8c50a6568afcac0113e61a8f6e5224a
Author: Austin Clements <austin@google.com>
Date:   Mon May 23 22:14:53 2016 -0400

    runtime: pass gcWork to scanstack
    
    Currently scanstack obtains its own gcWork from the P for the duration
    of the stack scan and then, if called during mark termination,
    disposes the gcWork.
    
    However, this means that the number of workbufs allocated will be at
    least the number of stacks scanned during mark termination, which may
    be very high (especially during a STW GC). This happens because, in
    steady state, each scanstack will obtain a fresh workbuf (either from
    the empty list or by allocating it), fill it with the scan results,
    and then dispose it to the full list. Nothing is consuming from the
    full list during this (and hence nothing is recycling them to the
    empty list), so the length of the full list by the time mark
    termination starts draining it is at least the number of stacks
    scanned.
    
    Fix this by pushing the gcWork acquisition up the stack to either the
    gcDrain that calls markroot that calls scanstack (which batches across
    many stack scans and is the path taken during STW GC) or to newstack
    (which is still a single scanstack call, but this is roughly bounded
    by the number of Ps).
    
    This fix reduces the workbuf allocation for the test program from
    issue #15319 from 213 MB (roughly 2KB * 1e5 goroutines) to 10 MB.
    
    Fixes #15319.
    
    Note that there's potentially a similar issue in write barriers during
    mark 2. Fixing that will be more difficult since there's no broader
    non-preemptible context, but it should also be less of a problem since
    the full list is being drained during mark 2.
    
    Some overall improvements in the go1 benchmarks, plus the usual noise.
    No significant change in the garbage benchmark (time/op or GC memory).
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.54s ± 1%     2.51s ± 1%  -1.09%  (p=0.000 n=20+19)
    Fannkuch11-12                2.12s ± 0%     2.17s ± 0%  +2.18%  (p=0.000 n=19+18)
    FmtFprintfEmpty-12          45.1ns ± 1%    45.2ns ± 0%    ~     (p=0.078 n=19+18)
    FmtFprintfString-12          127ns ± 0%     128ns ± 0%  +1.08%  (p=0.000 n=19+16)
    FmtFprintfInt-12             125ns ± 0%     122ns ± 1%  -2.71%  (p=0.000 n=14+18)
    FmtFprintfIntInt-12          196ns ± 0%     190ns ± 1%  -2.91%  (p=0.000 n=12+20)
    FmtFprintfPrefixedInt-12     196ns ± 0%     194ns ± 1%  -0.94%  (p=0.000 n=13+18)
    FmtFprintfFloat-12           253ns ± 1%     251ns ± 1%  -0.86%  (p=0.000 n=19+20)
    FmtManyArgs-12               807ns ± 1%     784ns ± 1%  -2.85%  (p=0.000 n=20+20)
    GobDecode-12                7.13ms ± 1%    7.12ms ± 1%    ~     (p=0.351 n=19+20)
    GobEncode-12                5.89ms ± 0%    5.95ms ± 0%  +0.94%  (p=0.000 n=19+19)
    Gzip-12                      219ms ± 1%     221ms ± 1%  +1.35%  (p=0.000 n=18+20)
    Gunzip-12                   37.5ms ± 1%    37.4ms ± 0%    ~     (p=0.057 n=20+19)
    HTTPClientServer-12         81.4µs ± 4%    81.9µs ± 3%    ~     (p=0.118 n=17+18)
    JSONEncode-12               15.7ms ± 1%    15.8ms ± 1%  +0.73%  (p=0.000 n=17+18)
    JSONDecode-12               57.9ms ± 1%    57.2ms ± 1%  -1.34%  (p=0.000 n=19+19)
    Mandelbrot200-12            4.12ms ± 1%    4.10ms ± 0%  -0.33%  (p=0.000 n=19+17)
    GoParse-12                  3.22ms ± 2%    3.25ms ± 1%  +0.72%  (p=0.000 n=18+20)
    RegexpMatchEasy0_32-12      70.6ns ± 1%    71.1ns ± 2%  +0.63%  (p=0.005 n=19+20)
    RegexpMatchEasy0_1K-12       240ns ± 0%     239ns ± 1%  -0.59%  (p=0.000 n=19+20)
    RegexpMatchEasy1_32-12      71.3ns ± 1%    71.3ns ± 1%    ~     (p=0.844 n=17+17)
    RegexpMatchEasy1_1K-12       384ns ± 2%     371ns ± 1%  -3.45%  (p=0.000 n=19+20)
    RegexpMatchMedium_32-12      109ns ± 1%     108ns ± 2%  -0.48%  (p=0.029 n=19+19)
    RegexpMatchMedium_1K-12     34.3µs ± 1%    34.5µs ± 2%    ~     (p=0.160 n=18+20)
    RegexpMatchHard_32-12       1.79µs ± 9%    1.72µs ± 2%  -3.83%  (p=0.000 n=19+19)
    RegexpMatchHard_1K-12       53.3µs ± 4%    51.8µs ± 1%  -2.82%  (p=0.000 n=19+20)
    Revcomp-12                   386ms ± 0%     388ms ± 0%  +0.72%  (p=0.000 n=17+20)
    Template-12                 62.9ms ± 1%    62.5ms ± 1%  -0.57%  (p=0.010 n=18+19)
    TimeParse-12                 325ns ± 0%     331ns ± 0%  +1.84%  (p=0.000 n=18+19)
    TimeFormat-12                338ns ± 0%     343ns ± 0%  +1.34%  (p=0.000 n=18+20)
    [Geo mean]                  52.7µs         52.5µs       -0.42%
    
    Change-Id: Ib2d34736c4ae2ec329605b0fbc44636038d8d018
    Reviewed-on: https://go-review.googlesource.com/23391
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit a1f7db88f8c70a7520ce870a71706032f173739f
Author: Austin Clements <austin@google.com>
Date:   Mon May 23 22:05:51 2016 -0400

    runtime: document scanstack
    
    Also mark it go:systemstack and explain why.
    
    Change-Id: I88baf22741c04012ba2588d8e03dd3801d19b5c0
    Reviewed-on: https://go-review.googlesource.com/23390
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit a689f6b8af6a1c4a388d6aabb64253e0ce9ab84f
Author: Robert Griesemer <gri@golang.org>
Date:   Wed May 25 11:23:56 2016 -0700

    cmd/compile: document how to update builtin.go
    
    No code changes.
    
    Fixes #15835.
    
    Change-Id: Ibae3f20882f976babc4093df5e9fea0b2cf0e9d9
    Reviewed-on: https://go-review.googlesource.com/23443
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit b7d96b8e059ce446f25e615ab9ef277eae2ef1c9
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed May 25 08:03:46 2016 -0400

    doc: reflect {Num,}Method skips unexported methods
    
    For #15673
    
    Change-Id: I3ce8d4016854d41860c5a9f05a54cda3de49f337
    Reviewed-on: https://go-review.googlesource.com/23430
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 07f0c19a30a5c83592bcd2f19e52eb8a40b32790
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Wed May 25 14:57:49 2016 +0200

    math/big: use run for benchmarks
    
    shortens code and gives an example of the use of Run.
    
    Change-Id: I75ffaf762218a589274b4b62e19022e31e805d1b
    Reviewed-on: https://go-review.googlesource.com/23424
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 23cb8864b52e5f2f60618a551ca564574e0575b0
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Wed May 25 19:34:01 2016 +0200

    runtime: use Run for more benchmarks
    
    Names for Append?Bytes are slightly changed in addition to adding a slash.
    
    Change-Id: I0291aa29c693f9040fd01368eaad9766259677df
    Reviewed-on: https://go-review.googlesource.com/23426
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d2aa5f95cce5055d71a77166a60b574cd3f8ecd5
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Wed May 25 11:19:17 2016 +0200

    compress/flate: simplify using subtests and sub-benchmarks
    
    This causes the large files to be loaded only once per benchmark.
    
    This CL also serves as an example use case of sub(tests|-benchmarks).
    
    This CL ensures that names are identical to the original
    except for an added slashes. Things could be
    simplified further if this restriction were dropped.
    
    Change-Id: I45e303e158e3152e33d0d751adfef784713bf997
    Reviewed-on: https://go-review.googlesource.com/23420
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 89283781c69314f09e999fe131b12c57059c1ba1
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Tue May 24 20:03:31 2016 +0200

    testing: added package doc for sub(tests/benchmarks)
    
    Change-Id: I6991cd7a41140da784a1ff8d69c5ea2032d05850
    Reviewed-on: https://go-review.googlesource.com/23354
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5dd922c935d28ded082f76a81e0c963938d7c3c3
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Wed May 25 14:54:41 2016 +0200

    compress/lzw: use Run for benchmarks
    
    load file only once per group.
    
    Change-Id: I965661507055e6e100506bf14d37133ecdd2cc5e
    Reviewed-on: https://go-review.googlesource.com/23423
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 095fbdcc91e41abf52a690dd6c64d701682ca96b
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Wed May 25 15:44:39 2016 +0200

    runtime: use of Run for some benchmarks
    
    Names of sub-benchmarks are preserved, short of the additional slash.
    
    Change-Id: I9b3f82964f9a44b0d28724413320afd091ed3106
    Reviewed-on: https://go-review.googlesource.com/23425
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 824e1f2e083c9c1df8455554744e49471becbaa2
Author: Robert Griesemer <gri@golang.org>
Date:   Tue May 24 12:54:14 2016 -0700

    text/scanner: better error message if no error handler is installed
    
    This is reverting golang.org/cl/19622 and introducing "<input>"
    as filename if no filename is specified.
    
    Fixes #15813.
    
    Change-Id: Iafc74b789fa33f48ee639c42d4aebc6f06435f95
    Reviewed-on: https://go-review.googlesource.com/23402
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 786e51d7baf6535290ecb0f34c9e73d7dcc21b02
Author: Keith Randall <khr@golang.org>
Date:   Tue May 24 16:53:48 2016 -0700

    cmd/compile: add generated tests for constant folding
    
    Covers a bunch of constant-folding rules in generic.rules that aren't
    being covered currently.
    
    Increases coverage in generic.rules from 65% to 72%.
    
    Change-Id: I7bf58809faf22e97070183b42e6dd7d3f35bf5f9
    Reviewed-on: https://go-review.googlesource.com/23407
    Run-TryBot: Todd Neal <todd@tneal.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Todd Neal <todd@tneal.org>

commit 9f38796270a017de8d9e1f102b28576826e6a188
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue May 24 19:04:51 2016 -0400

    reflect: remove type info for unexported methods
    
    Also remove some of the now unnecessary corner case handling and
    tests I've been adding recently for unexported method data.
    
    For #15673
    
    Change-Id: Ie0c7b03f2370bbe8508cdc5be765028f08000bd7
    Reviewed-on: https://go-review.googlesource.com/23410
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f2f3b6cd8fbe9f823fd6946f055bb70c3ef6f9db
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed May 25 13:24:36 2016 +0200

    cmd/link: fix ARM gold linker check
    
    CL 23400 introduced a check to make sure the gold linker is used
    on ARM host links. The check itself works, but the error checking
    logic was reversed; fix it.
    
    I manually verified that the check now correctly rejects host links
    on my RPi2 running an ancient rasbian without the gold linker
    installed.
    
    Updates #15696
    
    Change-Id: I927832620f0a60e91a71fdedf8cbd2550247b666
    Reviewed-on: https://go-review.googlesource.com/23421
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 72eb46c5a086051e3677579a0810922724eb6a6d
Author: Elias Naur <elias.naur@gmail.com>
Date:   Mon May 16 15:51:07 2016 +0200

    runtime,runtime/cgo: save callee-saved FP register on arm
    
    Other GOARCHs already handle their callee-saved FP registers, but
    arm was missing. Without this change, code using Cgo and floating
    point code might fail in mysterious and hard to debug ways.
    
    There are no floating point registers when GOARM=5, so skip the
    registers when runtime.goarm < 6.
    
    darwin/arm doesn't support GOARM=5, so the check is left out of
    rt0_darwin_arm.s.
    
    Fixes #14876
    
    Change-Id: I6bcb90a76df3664d8ba1f33123a74b1eb2c9f8b2
    Reviewed-on: https://go-review.googlesource.com/23140
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit fa3f484800415662cc741bbb8968ebb72896e20a
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue May 24 11:39:48 2016 -0700

    encoding/csv: clarify that this package supports RFC 4180
    
    The intent of this comment is to reduce the number of issues opened
    against the package to add support for new kinds of CSV formats, such as
    issues #3150, #8458, #12372, #12755.
    
    Change-Id: I452c0b748e4ca9ebde3e6cea188bf7774372148e
    Reviewed-on: https://go-review.googlesource.com/23401
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 93e8e704996ce100fe46b2249324442947e47a9d
Author: Robert Griesemer <gri@golang.org>
Date:   Tue May 24 14:12:35 2016 -0700

    all: fixed a handful of typos
    
    Change-Id: Ib0683f27b44e2f107cca7a8dcc01d230cbcd5700
    Reviewed-on: https://go-review.googlesource.com/23404
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit a640d95172ec516a3b727a8ca796c64f5aea89ec
Author: Austin Clements <austin@google.com>
Date:   Fri May 20 14:57:55 2016 -0400

    runtime: update SP when jumping stacks in traceback
    
    When gentraceback starts on a system stack in sigprof, it is
    configured to jump to the user stack when it reaches the end of the
    system stack. Currently this updates the current frame's FP, but not
    its SP. This is okay on non-LR machines (x86) because frame.sp is only
    used to find defers, which the bottom-most frame of the user stack
    will never have.
    
    However, on LR machines, we use frame.sp to find the saved LR. We then
    use to resolve the function of the next frame, which is used to
    resolved the size of the next frame. Since we're not updating frame.sp
    on a stack jump, we read the saved LR from the system stack instead of
    the user stack and wind up resolving the wrong function and hence the
    wrong frame size for the next frame.
    
    This has had remarkably few ill effects (though the resulting profiles
    must be wrong). We noticed it because of a bad interaction with stack
    barriers. Specifically, once we get the next frame size wrong, we also
    get the location of its LR wrong. If we happen to get a stack slot
    that contains a stale stack barrier LR (for a stack barrier we already
    hit) and hasn't been overwritten with something else as we re-grew the
    stack, gentraceback will fail with a "found next stack barrier at ..."
    error, pointing at the slot that it thinks is an LR, but isn't.
    
    Fixes #15138.
    
    Updates #15313 (might fix it).
    
    Change-Id: I13cfa322b44c0c2f23ac2b3d03e12631e4a6406b
    Reviewed-on: https://go-review.googlesource.com/23291
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 7b6b5e340403c32756a92294729dc52f308b841a
Author: Jeff R. Allen <jra@nella.org>
Date:   Wed May 25 00:12:22 2016 +0600

    doc: add notes on good commit messages
    
    Explain Brad's algorithm for generating commit headlines.
    
    Fixes #15700
    
    Change-Id: Ic602f17629b3dd7675e2bb1ed119062c03353ee9
    Reviewed-on: https://go-review.googlesource.com/23355
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3474610fbc81f7e9f3f2cb23dc1554b3f5cec657
Author: Jeff R. Allen <jra@nella.org>
Date:   Tue May 24 23:00:06 2016 +0600

    math/rand: Doc fix for how many bits Seed uses
    
    Document the fact that the default Source uses only
    the bottom 31 bits of the given seed.
    
    Fixes #15788
    
    Change-Id: If20d1ec44a55c793a4a0a388f84b9392c2102bd1
    Reviewed-on: https://go-review.googlesource.com/23352
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2f9eeea212681927385a4713004438cd864966e0
Author: Jeff R. Allen <jra@nella.org>
Date:   Tue May 24 23:46:11 2016 +0600

    cmd/go: document testdata directory in "go help test"
    
    Document the correct use of the testdata directory
    where test writers might be expecting to find it.
    
    It seems that alldocs.go was out of date, so it
    has picked up some other changes with this commit.
    
    Fixes #14715.
    
    Change-Id: I0a22676bb7a64b2a61b56495f7ea38db889d8b37
    Reviewed-on: https://go-review.googlesource.com/23353
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a0abecf1020c33e82c464f7891b317e83f0c6a78
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue May 24 13:40:02 2016 -0400

    cmd/link: ensure -fuse-ld=gold uses gold
    
    Fixes #15696
    
    Change-Id: I134e918dc56f79a72a04aa54f415371884113d2a
    Reviewed-on: https://go-review.googlesource.com/23400
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 7b9d3ff4cbd93c0b9e69c78cbb2cb891839c7fb3
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Sat May 21 14:37:29 2016 +0200

    testing: don't be silent if a test's goroutine fails a test after test exits
    
    Fixes #15654
    
    Change-Id: I9bdaa9b76d480d75f24d95f0235efd4a79e3593e
    Reviewed-on: https://go-review.googlesource.com/23320
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit dcc42c7d11ad06bebc9d13d1e812629f930f14a7
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Tue May 24 13:53:44 2016 +0300

    cmd/vet: do not check print-like functions with unknown type
    
    Fixes #15787
    
    Change-Id: I559ba886527b474dbdd44fe884c78973b3012377
    Reviewed-on: https://go-review.googlesource.com/23351
    Run-TryBot: Rob Pike <r@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 524956f8b976be2b7be829a2d0d87c2951932ac6
Author: Russ Cox <rsc@golang.org>
Date:   Tue May 24 09:42:33 2016 -0400

    io: remove SizedReaderAt
    
    It's not clear we want to enshrine an io interface in which Size cannot
    return an error. Because this requires more thought before committing
    to the API, remove from Go 1.7.
    
    Fixes #15818.
    
    Change-Id: Ic4138ffb0e033030145a12d33f78078350a8381f
    Reviewed-on: https://go-review.googlesource.com/23392
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit ba867a86fa28f9edca64b682bc2df66e73967f56
Author: Russ Cox <rsc@golang.org>
Date:   Tue May 24 02:50:17 2016 -0400

    api: update next.txt
    
    Change-Id: I7b38309d927409a92f68f5d26f491b0166eba838
    Reviewed-on: https://go-review.googlesource.com/23378
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit f117a5359af607d8e62fa9cd66eb327d18a7d6d7
Author: Russ Cox <rsc@golang.org>
Date:   Tue May 24 03:12:41 2016 -0400

    doc: first draft of Go 1.7 release notes
    
    Mostly complete but a few TODOs remain for future CLs.
    
    For #15810.
    
    Change-Id: I81ee19d1088d192cf709a5f7e6b7bcc44ad892ac
    Reviewed-on: https://go-review.googlesource.com/23379
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit 85e3c9e6b863792135c8cd49bebfd1028e87cee5
Author: Russ Cox <rsc@golang.org>
Date:   Tue May 24 02:52:31 2016 -0400

    cmd/compile, go/types: omit needless word in error message
    
    CL 21462 and CL 21463 made this message say explicitly that the problem
    was a struct field in a map, but the word "directly" is unnecessary,
    sounds wrong, and makes the error long.
    
    Change-Id: I2fb68cdaeb8bd94776b8022cf3eae751919ccf6f
    Reviewed-on: https://go-review.googlesource.com/23373
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 34b17d4dc5726eebde437f2c1b680d039cc3e7c0
Author: Russ Cox <rsc@golang.org>
Date:   Mon May 23 12:28:56 2016 -0400

    encoding/json: rename Indent method to SetIndent
    
    CL 21057 added this method during the Go 1.7 cycle
    (so it is not yet released and still possible to revise).
    
    This makes it clearer that the method is not doing something
    (like func Indent does), but just changing a setting about doing
    something later.
    
    Also document that this is in some sense irreversible.
    I think that's probably a mistake but the original CL discussion
    claimed it as a feature, so I'll leave it alone.
    
    For #6492.
    
    Change-Id: If4415c869a9196501056c143811a308822d5a420
    Reviewed-on: https://go-review.googlesource.com/23295
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit 4aea7a12b6a6621a67267050df0688f28adfe6b4
Author: Russ Cox <rsc@golang.org>
Date:   Mon May 23 11:41:00 2016 -0400

    encoding/json: change DisableHTMLEscaping to SetEscapeHTML
    
    DisableHTMLEscaping is now SetEscapeHTML, allowing the escaping
    to be toggled, not just disabled. This API is new for Go 1.7,
    so there are no compatibility concerns (quite the opposite,
    the point is to fix the API before we commit to it in Go 1.7).
    
    Change-Id: I96b9f8f169a9c44995b8a157a626eb62d0b6dea7
    Reviewed-on: https://go-review.googlesource.com/23293
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 658814fa386a08761032b659f64f4eecb79ca084
Author: Russ Cox <rsc@golang.org>
Date:   Tue May 24 02:25:56 2016 -0400

    A+C: add Andre Nathan (individual CLA)
    
    Andre wrote https://golang.org/cl/13454043 (never submitted),
    which served as the basis for Ross Light's https://golang.org/cl/19235.
    
    Individual CLA verified by hand.
    
    Change-Id: Ic09e8efd84b7ded3ae472c204133e40cb85d97f7
    Reviewed-on: https://go-review.googlesource.com/23377
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7f9255c2120f784c334431661145ee89e75f64fe
Author: Russ Cox <rsc@golang.org>
Date:   Mon May 23 20:40:52 2016 -0400

    net: revise IP.String result for malformed IP address to add ? back
    
    In earlier versions of Go the result was simply "?".
    A change in this cycle made the result echo back the hex bytes
    of the address, which is certainly useful, but now the result is
    not clearly indicating an error. Put the "?" back, at the beginning
    of the hex string, to make the invalidity of the string clearer.
    
    Change-Id: I3e0f0b6a005601cd98d982a62288551959185b40
    Reviewed-on: https://go-review.googlesource.com/23376
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8e724e7bbad07d530f547b6214c71ddfe26ba92a
Author: Russ Cox <rsc@golang.org>
Date:   Mon May 23 20:34:39 2016 -0400

    cmd/go: fix //go:binary-only-package check
    
    The use of a prefix check was too liberal.
    Noted in review after submit.
    
    Change-Id: I4fe1df660997efd225609e818040b8392fab79f0
    Reviewed-on: https://go-review.googlesource.com/23375
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 12610236375e2e981e370508548f3d51920df7e2
Author: Russ Cox <rsc@golang.org>
Date:   Mon May 23 12:21:57 2016 -0400

    encoding/json: additional tests and fixes for []typedByte encoding/decoding
    
    CL 19725 changed the encoding of []typedByte to look for
    typedByte.MarshalJSON and typedByte.MarshalText.
    Previously it was handled like []byte, producing a base64 encoding of the underlying byte data.
    
    CL 19725 forgot to look for (*typedByte).MarshalJSON and (*typedByte).MarshalText,
    as the marshaling of other slices would. Add test and fix for those.
    
    This CL also adds tests that the decoder can handle both the old and new encodings.
    (This was true even in Go 1.6, which is the only reason we can consider this
    not an incompatible change.)
    
    For #13783.
    
    Change-Id: I7cab8b6c0154a7f2d09335b7fa23173bcf856c37
    Reviewed-on: https://go-review.googlesource.com/23294
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit ab4414773e27624abf4361e48a0ca0979e804970
Author: Russ Cox <rsc@golang.org>
Date:   Tue May 24 02:45:11 2016 -0400

    math/big: write t*10 to multiply t by 10
    
    The compiler has caught up.
    In fact the compiler is ahead; it knows about a magic multiply-by-5 instruction:
    
    	// compute '0' + byte(r - t*10) in AX
    	MOVQ	t, AX
    	LEAQ	(AX)(AX*4), AX
    	SHLQ	$1, AX
    	MOVQ	r, CX
    	SUBQ	AX, CX
    	LEAL	48(CX), AX
    
    For comparison, the shifty version compiles to:
    
    	// compute '0' + byte(r - t*10) in AX
    	MOVQ	t, AX
    	MOVQ	AX, CX
    	SHLQ	$3, AX
    	MOVQ	r, DX
    	SUBQ	AX, DX
    	SUBQ	CX, DX
    	SUBQ	CX, DX
    	LEAL	48(DX), AX
    
    Fixes #2671.
    
    Change-Id: Ifbf23dbfeb19c0bb020fa44eb2f025943969fb6b
    Reviewed-on: https://go-review.googlesource.com/23372
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7a9f6c2b56bd87ff7f9296344c9e63cc46194428
Author: Keith Randall <khr@golang.org>
Date:   Mon May 23 13:09:12 2016 -0700

    cmd/compile: benchmark needs dominator tree
    
    Now that CSE uses dom tree to order partitions, we need the
    dom tree computed before benchmarking CSE.
    
    Fixes #15801
    
    Change-Id: Ifa4702c7b75250f34de185e69a880b3f3cc46a12
    Reviewed-on: https://go-review.googlesource.com/23361
    Reviewed-by: David Chase <drchase@google.com>

commit 30fc940c70ee5ee27f0a455248735b8b57f34fb7
Author: Robert Griesemer <gri@golang.org>
Date:   Fri May 20 17:26:24 2016 -0700

    go/types: don't drop type in n:1 var decl if one is given
    
    In n:1 variable declarations (multiple lhs variables with single
    multi-valued initialization expression) where also a variable
    type is provided, make sure that that type is assigned to all
    variables on the lhs before the init expression assignment is
    checked. Otherwise, (some) variables are assumed to take the type
    of the corresponding value of the multi-valued init expression.
    
    Fixes #15755.
    
    Change-Id: I969cb5a95c85e28dbb38abd7fa7df16ff5554c03
    Reviewed-on: https://go-review.googlesource.com/23313
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 65adb6ab9a91c2f4336e1e48f9e7c325dafa0213
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Sat May 21 16:57:37 2016 +0100

    time: run genzabbrs.go with new source data
    
    The source xml data has changed, so running genzabbrs.go
    regenerates a new time zone file in zoneinfo_abbrs_windows.go
    which adds some zones and adjusts others.
    
    Now set export ZONEINFO=$GOROOT/lib/time/zoneinfo.zip to use zoneinfo.zip in go tip.
    
    Change-Id: I19f72359cc808094e5dcb420e480a00c6b2205d7
    Reviewed-on: https://go-review.googlesource.com/23321
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4b6e560b69f3fe6dfadcf6a43eda8a78b086448c
Author: djherbis <djherbis@gmail.com>
Date:   Sat May 21 13:25:47 2016 -0700

    AUTHORS: correcting my last name Herbis -> Herbison
    
    Change-Id: I91608b15e00c8eaf732db3a99a890d4ceeb41955
    Reviewed-on: https://go-review.googlesource.com/23317
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 1f8d2768987ea1d7f1bc4d6bbfe59b2a8e98d9b9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 20 23:28:56 2016 +0000

    A+C: automated update (subrepos)
    
    Add Abe Haskins (individual CLA)
    Add Ahmy Yulrizka (individual CLA)
    Add Akihiro Suda (individual CLA)
    Add Alex Vaghin (corporate CLA for Google Inc.)
    Add Arlo Breault (individual CLA)
    Add Audrey Lim (individual CLA)
    Add Benjamin Wester (corporate CLA for Square, Inc.)
    Add Bryan Chan (corporate CLA for IBM)
    Add Christy Perez (corporate CLA for IBM)
    Add Colin Edwards (individual CLA)
    Add David Brophy (individual CLA)
    Add David Sansome (individual CLA)
    Add Diwaker Gupta (individual CLA)
    Add Doug Anderson (corporate CLA for Google Inc.)
    Add Dustin Carlino (corporate CLA for Google Inc.)
    Add Ernest Chiang (individual CLA)
    Add Ethan Burns (corporate CLA for Google Inc.)
    Add Gary Elliott (corporate CLA for Google Inc.)
    Add Hallgrimur Gunnarsson (corporate CLA for Google Inc.)
    Add Hironao OTSUBO (individual CLA)
    Add Holden Huang (individual CLA)
    Add Idora Shinatose (individual CLA)
    Add Irieda Noboru (individual CLA)
    Add Jeff Craig (corporate CLA for Google Inc.)
    Add Joe Henke (individual CLA)
    Add John Schnake (individual CLA)
    Add Jonathan Amsterdam (corporate CLA for Google Inc.)
    Add Kenji Kaneda (individual CLA)
    Add Kenneth Shaw (individual CLA)
    Add Mark Severson (individual CLA)
    Add Martin Garton (individual CLA)
    Add Mathias Leppich (individual CLA)
    Add Maxwell Krohn (individual CLA)
    Add Niall Sheridan (individual CLA)
    Add Nick Patavalis (individual CLA)
    Add Nick Petroni (individual CLA)
    Add Omar Jarjur (corporate CLA for Google Inc.)
    Add Özgür Kesim (individual CLA)
    Add Peter Gonda (corporate CLA for Google Inc.)
    Add Pierre Durand (individual CLA)
    Add Quentin Smith (corporate CLA for Google Inc.)
    Add Ricardo Padilha (individual CLA)
    Add Riku Voipio (corporate CLA for Linaro Limited)
    Add Roland Shoemaker (individual CLA)
    Add Sam Hug (individual CLA)
    Add Sam Whited (individual CLA)
    Add Sami Commerot (corporate CLA for Google Inc.)
    Add Scott Mansfield (corporate CLA for Netflix, Inc.)
    Add Sean Harger (corporate CLA for Google Inc.)
    Add Simon Jefford (individual CLA)
    Add Sridhar Venkatakrishnan (individual CLA)
    Add Tim Swast (corporate CLA for Google Inc.)
    Add Timothy Studd (individual CLA)
    Add Tipp Moseley (corporate CLA for Google Inc.)
    Add Toby Burress (corporate CLA for Google Inc.)
    Add Tzu-Jung Lee (corporate CLA for Currant)
    Add Vadim Grek (individual CLA)
    Add Xudong Zhang (individual CLA)
    
    Updates #12042
    
    Change-Id: I4119a8829119a2b8a9abbea9f52ceebb04878764
    Reviewed-on: https://go-review.googlesource.com/23306
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit def50f8e488dcfb12362d4b84feb38de3af5cadc
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat May 21 00:25:48 2016 +0000

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 0c607074 for https://golang.org/cl/23311,
    "http2: prevent Server from sending status 100 header after anything else"
    
    New test is in the x/net/http2 package (not bundled to std).
    
    Fixes #14030
    
    Change-Id: Ifc6afa4a5fe35977135428f6d0e9f7c164767720
    Reviewed-on: https://go-review.googlesource.com/23312
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 82ec4cd79f117191d12fc14060c4b4b786feca5b
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri May 13 12:13:00 2016 +0900

    net: don't crash DNS flood test on darwin
    
    Also renames the test function to TestDNSFlood.
    
    Updates #15659.
    
    Change-Id: Ia562004c43bcc19c2fee9440321c27b591f85da5
    Reviewed-on: https://go-review.googlesource.com/23077
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit da5ac69bd4dce05443220c73c9eadb606b9777f8
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 20 21:52:59 2016 +0000

    A+C: automated updates
    
    Add Aiden Scandella (individual CLA)
    Add Alessandro Arzilli (individual CLA)
    Add Augusto Roman (individual CLA)
    Add Brady Catherman (individual CLA)
    Add Brady Sullivan (individual CLA)
    Add Caio Marcelo de Oliveira Filho (corporate CLA for Intel Corporation)
    Add Catalin Nicutar (corporate CLA for Google Inc.)
    Add Cherry Zhang (corporate CLA for Google Inc.)
    Add Chris Zou (corporate CLA for IBM)
    Add Christopher Nelson (individual CLA)
    Add Conrad Irwin (individual CLA)
    Add Cuihtlauac ALVARADO (corporate CLA for Orange)
    Add Daniel Speichert (individual CLA)
    Add Datong Sun (individual CLA)
    Add Denys Honsiorovskyi (individual CLA)
    Add Derek Shockey (individual CLA)
    Add Dmitriy Dudkin (individual CLA)
    Add Dustin Herbis (individual CLA)
    Add Frits van Bommel (individual CLA)
    Add Harshavardhana (individual CLA)
    Add Hitoshi Mitake (individual CLA)
    Add James Bardin (individual CLA)
    Add James Chacon (corporate CLA for Google Inc.)
    Add Jamil Djadala (individual CLA)
    Add Jess Frazelle (individual CLA)
    Add Joe Sylve (individual CLA)
    Add Johan Sageryd (individual CLA)
    Add John Jeffery (individual CLA)
    Add Julia Hansbrough (corporate CLA for Google Inc.)
    Add Jure Ham (corporate CLA for Zemanta d.o.o.)
    Add Kamal Aboul-Hosn (corporate CLA for Google Inc.)
    Add Kevin Burke (individual CLA)
    Add Kevin Kirsche (individual CLA)
    Add Kevin Vu (individual CLA)
    Add Lee Hinman (individual CLA)
    Add Luan Santos (individual CLA)
    Add Marc-Antoine Ruel (corporate CLA for Google Inc.)
    Add Matt Robenolt (individual CLA)
    Add Michael McConville (individual CLA)
    Add Michael Munday (corporate CLA for IBM)
    Add Michael Pratt (corporate CLA for Google Inc.)
    Add Michel Lespinasse (corporate CLA for Google Inc.)
    Add Mike Danese (corporate CLA for Google Inc.)
    Add Mikhail Gusarov (individual CLA)
    Add Monty Taylor (individual CLA)
    Add Morten Siebuhr (individual CLA)
    Add Muhammed Uluyol (individual CLA)
    Add Niels Widger (individual CLA)
    Add Niko Dziemba (individual CLA)
    Add Olivier Poitrey (individual CLA)
    Add Paul Wankadia (corporate CLA for Google Inc.)
    Add Philip Hofer (individual CLA)
    Add Prashant Varanasi (individual CLA)
    Add Rhys Hiltner (corporate CLA for Amazon.com, Inc)
    Add Richard Miller (individual CLA)
    Add Scott Bell (individual CLA)
    Add Shahar Kohanim (individual CLA)
    Add Shinji Tanaka (individual CLA)
    Add Suharsh Sivakumar (corporate CLA for Google Inc.)
    Add Tal Shprecher (individual CLA)
    Add Tilman Dilo (individual CLA)
    Add Tim Ebringer (individual CLA)
    Add Tom Bergan (corporate CLA for Google Inc.)
    Add Vishvananda Ishaya (individual CLA)
    Add Wedson Almeida Filho (corporate CLA for Google Inc.)
    Add Zhongwei Yao (corporate CLA for ARM Ltd.)
    
    Updates #12042
    
    Change-Id: Ia118adc2eb38e5ffc8448de2d9dd3ca792ee7227
    Reviewed-on: https://go-review.googlesource.com/23303
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 7ab698b0221fd4a5b2842fb50a34ba8a5f49c6d5
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 20 22:43:14 2016 +0000

    time: document that After uses memory until duration times out
    
    Fixes #15698
    
    Change-Id: I616fc06dcf04092bafdaf56fb1afba2a998a6d83
    Reviewed-on: https://go-review.googlesource.com/23304
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 054a721dcac9b30610af0898b3ed8bf3ffa9f8b1
Author: Robert Griesemer <gri@golang.org>
Date:   Fri May 20 11:19:19 2016 -0700

    cmd/compile: read safemode bit from package header
    
    Ignore respective bit in export data, but leave the info to
    minimize format changes for 1.7. Scheduled to remove by 1.8.
    
    For #15772.
    
    Change-Id: Ifb3beea655367308a4e2d5dc8cb625915f904287
    Reviewed-on: https://go-review.googlesource.com/23285
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit cc0d8c86e3437c1eec697809bdc9b2bcc8e0ed92
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 20 18:13:49 2016 +0000

    net/http: deflake TestTransportEventTrace_h2
    
    Fixes #15765
    
    Change-Id: Id0a89d90ef9d3fffa9af0affca8c10a26fe6b7bc
    Reviewed-on: https://go-review.googlesource.com/23284
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit db5af0d7115781be12f510322ee01556fb1e6d16
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 20 20:58:44 2016 +0000

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 4d07e8a49 for CL 23287:
    
       http2: let handlers close Request.Body without killing streams
       https://golang.org/cl/23287
    
    Fixes #15425
    
    Change-Id: I20b6e37cd09aa1d5a040c122ca0daf14b8916559
    Reviewed-on: https://go-review.googlesource.com/23301
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4cad610401edc11fe921205438a7b3ab4faa3982
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 20 20:42:21 2016 +0000

    os/exec: remove Cmd.RunContext and Cmd.WaitContext, add CommandContext
    
    Fixes #15775
    
    Change-Id: I0a6c2ca09d3850c3538494711f7a9801b9500411
    Reviewed-on: https://go-review.googlesource.com/23300
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 85e39f838722a1521e09288cddfe378843d662fb
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 20 16:30:49 2016 +0000

    net/http: also clone DynamicRecordSizingDisabled in cloneTLSConfig
    
    Updates #15771
    
    Change-Id: I5dad96bdca19d680dd00cbd17b72a03e43eb557e
    Reviewed-on: https://go-review.googlesource.com/23283
    Reviewed-by: Tom Bergan <tombergan@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit be1b93065356e71362ca8469fc53c9ab102c4be5
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu May 19 13:31:58 2016 -0400

    reflect: hide unexported methods that do not satisfy interfaces
    
    Fixes #15673
    
    Change-Id: Ib36d8db3299a93d92665dbde012d52c2c5332ac0
    Reviewed-on: https://go-review.googlesource.com/23253
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b3bf2e7803174b77838a357ee05b1351ece6a29d
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 20 03:01:27 2016 +0000

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 8a52c78 for golang.org/cl/23258
    (http2: fix Transport.CloseIdleConnections when http1+http2 are wired together)
    
    Fixes #14607
    
    Change-Id: I038badc69e230715b8ce4e398eb5e6ede73af918
    Reviewed-on: https://go-review.googlesource.com/23280
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1ded9fdcff8722ae961fb9da015faac874b7690e
Author: Jess Frazelle <me@jessfraz.com>
Date:   Thu May 19 22:26:01 2016 -0700

    syscall: fix unshare test on mips
    
    Change-Id: Iedce3770a92112802f3a45c7b95ee145ab5b187e
    Reviewed-on: https://go-review.googlesource.com/23282
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit 16f846a9cbe747b13498761f1dd1a298478ec43e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu May 19 17:35:23 2016 +0000

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 202ff482 for https://golang.org/cl/23235 (Expect:
    100-continue support for HTTP/2)
    
    Fixes a flaky test too, and changes the automatic HTTP/2 behavior to
    no longer special-case the DefaultTransport, because
    ExpectContinueTimeout is no longer unsupported by the HTTP/2
    transport.
    
    Fixes #13851
    Fixes #15744
    
    Change-Id: I3522aace14179a1ca070fd7063368a831167a0f7
    Reviewed-on: https://go-review.googlesource.com/23254
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 9cd2c700deccc6dfcc8f264857e406c53bf07859
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu May 19 06:15:18 2016 +0900

    net: deflake TestDialTimeoutMaxDuration
    
    Fixes #15745.
    
    Change-Id: I6f9a1dcf0b1d97cb443900c7d8da09ead83d4b6a
    Reviewed-on: https://go-review.googlesource.com/23243
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8527b8ef9b00c72b1a8e30e5917c7bdd3c0e79ef
Author: Jess Frazelle <me@jessfraz.com>
Date:   Wed May 18 18:47:24 2016 -0700

    syscall: add Unshare flags to SysProcAttr on Linux
    
    This patch adds Unshare flags to SysProcAttr for Linux systems.
    
    Fixes #1954
    
    Change-Id: Id819c3f92b1474e5a06dd8d55f89d74a43eb770c
    Reviewed-on: https://go-review.googlesource.com/23233
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 448246adff7feb868d66cfde82b36fcfd0e66b75
Author: Robert Griesemer <gri@golang.org>
Date:   Wed May 18 17:43:15 2016 -0700

    cmd/compile: don't exit early because of hidden error messages
    
    Non-syntax errors are always counted to determine if to exit
    early, but then deduplication eliminates them. This can lead
    to situations which report "too many errors" and only one
    error is shown.
    
    De-duplicate non-syntax errors early, at least the ones that
    appear consecutively, and only count the ones actually being
    shown. This doesn't work perfectly as they may not appear in
    sequence, but it's cheap and good enough.
    
    Fixes #14136.
    
    Change-Id: I7b11ebb2e1e082f0d604b88e544fe5ba967af1d7
    Reviewed-on: https://go-review.googlesource.com/23259
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit dc4427f3727804ded270bc6a7a8066ccb3c151d0
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu May 19 18:08:43 2016 +0000

    context: make DeadlineExceeded have a Timeout method
    
    Fixes #14238
    
    Change-Id: I1538bfb5cfa63e36a89df1f6eb9f5a0dcafb6ce5
    Reviewed-on: https://go-review.googlesource.com/23256
    Reviewed-by: Dave Cheney <dave@cheney.net>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0b80659832ec72532ee1210cdb51422ee6012c66
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu May 19 18:05:10 2016 +0000

    net/http/httptest: restore historic ResponseRecorder.HeaderMap behavior
    
    In Go versions 1 up to and including Go 1.6,
    ResponseRecorder.HeaderMap was both the map that handlers got access
    to, and was the map tests checked their results against. That did not
    mimic the behavior of the real HTTP server (Issue #8857), so HeaderMap
    was changed to be a snapshot at the first write in
    https://golang.org/cl/20047. But that broke cases where the Handler
    never did a write (#15560), so revert the behavior.
    
    Instead, introduce the ResponseWriter.Result method, returning an
    *http.Response. It subsumes ResponseWriter.Trailers which was added
    for Go 1.7 in CL 20047. Result().Header now contains the correct
    answer, and HeaderMap is unchanged in behavior from previous Go
    releases, so we don't break people's tests. People wanting the correct
    behavior can use ResponseWriter.Result.
    
    Fixes #15560
    Updates #8857
    
    Change-Id: I7ea9b56a6b843103784553d67f67847b5315b3d2
    Reviewed-on: https://go-review.googlesource.com/23257
    Reviewed-by: Damien Neil <dneil@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3b50adbc4f1a9d775f0434166ad71220e8a4b8ce
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu May 19 17:43:04 2016 +0000

    build: unset GOBIN during build
    
    Fixes #14340
    
    Change-Id: I43e1624fafc972fb868708c3857fc8acf1bfbbd7
    Reviewed-on: https://go-review.googlesource.com/23255
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 44497ebacb6336e4cc9ce2934840bdd68e8c46c0
Author: Austin Clements <austin@google.com>
Date:   Tue May 17 18:46:03 2016 -0400

    runtime: fix goroutine priority elevation
    
    Currently it's possible for user code to exploit the high scheduler
    priority of the GC worker in conjunction with the runnext optimization
    to elevate a user goroutine to high priority so it will always run
    even if there are other runnable goroutines.
    
    For example, if a goroutine is in a tight allocation loop, the
    following can happen:
    
    1. Goroutine 1 allocates, triggering a GC.
    2. G 1 attempts an assist, but fails and blocks.
    3. The scheduler runs the GC worker, since it is high priority.
       Note that this also starts a new scheduler quantum.
    4. The GC worker does enough work to satisfy the assist.
    5. The GC worker readies G 1, putting it in runnext.
    6. GC finishes and the scheduler runs G 1 from runnext, giving it
       the rest of the GC worker's quantum.
    7. Go to 1.
    
    Even if there are other goroutines on the run queue, they never get a
    chance to run in the above sequence. This requires a confluence of
    circumstances that make it unlikely, though not impossible, that it
    would happen in "real" code. In the test added by this commit, we
    force this confluence by setting GOMAXPROCS to 1 and GOGC to 1 so it's
    easy for the test to repeated trigger GC and wake from a blocked
    assist.
    
    We fix this by making GC always put user goroutines at the end of the
    run queue, instead of in runnext. This makes it so user code can't
    piggy-back on the GC's high priority to make a user goroutine act like
    it has high priority. The only other situation where GC wakes user
    goroutines is waking all blocked assists at the end, but this uses the
    global run queue and hence doesn't have this problem.
    
    Fixes #15706.
    
    Change-Id: I1589dee4b7b7d0c9c8575ed3472226084dfce8bc
    Reviewed-on: https://go-review.googlesource.com/23172
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 91740582c3ec1e57621cbb0ec0f9163431f1b688
Author: Austin Clements <austin@google.com>
Date:   Tue May 17 18:21:54 2016 -0400

    runtime: add 'next' flag to ready
    
    Currently ready always puts the readied goroutine in runnext. We're
    going to have to change this for some uses, so add a flag for whether
    or not to use runnext.
    
    For now we always pass true so this is a no-op change.
    
    For #15706.
    
    Change-Id: Iaa66d8355ccfe4bbe347570cc1b1878c70fa25df
    Reviewed-on: https://go-review.googlesource.com/23171
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 79ba1e44c7c2d7ff186f9ac142a85869f352f0f6
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu May 19 10:07:41 2016 -0700

    cmd/cgo: mark stub functions as no_sanitize_thread
    
    When the generated stub functions write back the results to the stack,
    they can in some cases be writing to the same memory on the g0 stack.
    There is no race here (assuming there is no race in the Go code), but
    the thread sanitizer does not know that.  Turn off the thread sanitizer
    for the stub functions to prevent false positive warnings.
    
    Current clang suggests the no_sanitize("thread") attribute, but that
    does not work with clang 3.6 or GCC.  clang 3.6, GCC, and current clang
    all support the no_sanitize_thread attribute, so use that
    unconditionally.
    
    The test case and first version of the patch are from Dmitriy Vyukov.
    
    Change-Id: I80ce92824c6c8cf88ea0fe44f21cf50cf62474c9
    Reviewed-on: https://go-review.googlesource.com/23252
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0dcd330bc8ab19db15b4517b80e940cf154071bc
Author: Joel Sing <joel@sing.id.au>
Date:   Wed May 18 23:39:23 2016 +1000

    runtime/cgo: make cgo work with openbsd ABI changes
    
    OpenBSD 6.0 (due out November 2016) will support PT_TLS, which will
    allow for the OpenBSD cgo pthread_create() workaround to be removed.
    
    However, in order for Go to continue working on supported OpenBSD
    releases (the current release and the previous release - 5.9 and 6.0,
    once 6.0 is released), we cannot enable PT_TLS immediately. Instead,
    adjust the existing code so that it works with the previous TCB
    allocation and the new TIB allocation. This allows the same Go
    runtime to work on 5.8, 5.9 and later 6.0.
    
    Once OpenBSD 5.9 is no longer supported (May 2017, when 6.1 is
    released), PT_TLS can be enabled and the additional cgo runtime
    code removed.
    
    Change-Id: I3eed5ec593d80eea78c6656cb12557004b2c0c9a
    Reviewed-on: https://go-review.googlesource.com/23197
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Joel Sing <joel@sing.id.au>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d603c27c6b462f044a7079ce5113d90bb3ca4814
Author: Keith Randall <khr@golang.org>
Date:   Wed May 18 13:04:00 2016 -0700

    cmd/compile: large ptr literals must escape
    
    They get rewritten to NEWs, and they must be marked as escaping
    so walk doesn't try to allocate them back onto the stack.
    
    Fixes #15733
    
    Change-Id: I433033e737c3de51a9e83a5a273168dbc9110b74
    Reviewed-on: https://go-review.googlesource.com/23223
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 1ab9428eec6cd1595de571aac4c093645a6629d0
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu May 19 06:15:18 2016 +0900

    net: deflake TestDialerDualStack
    
    Fixes #15316.
    Fixes #15574.
    
    Change-Id: I3ec8bffd35b9e5123de4be983a53fc0b8c2a0895
    Reviewed-on: https://go-review.googlesource.com/23242
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 086d7b0e9e34555f32248c9242b641273a32bc7e
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu May 19 17:36:47 2016 +0900

    net: deflake TestDialerDualStackFDLeak
    
    Fixes #14717.
    Updates #15157.
    
    Change-Id: I7238b4fe39f3670c2dfe09b3a3df51a982f261ed
    Reviewed-on: https://go-review.googlesource.com/23244
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
```
