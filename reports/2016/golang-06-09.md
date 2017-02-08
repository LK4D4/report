# June 9, 2016 Report

Number of commits: 38

## Compilation time

* github.com/coreos/etcd: from 14.127932902s to 14.323570285s, +1.38%
* github.com/boltdb/bolt/cmd/bolt: from 558.411301ms to 548.796344ms, -1.72%
* github.com/gogits/gogs: from 12.924443958s to 12.736883022s, -1.45%
* github.com/spf13/hugo: from 7.222385433s to 6.557703885s, -9.20%
* github.com/influxdata/influxdb/cmd/influxd: from 6.374784733s to 6.590248407s, +3.38%
* github.com/nsqio/nsq/apps/nsqd: from 2.183671731s to 2.196358709s, +0.58%
* github.com/mholt/caddy: from 265.682637ms to 272.739581ms, +2.66%

## Binary size:

* github.com/coreos/etcd: from 22693608 to 22696192, ~
* github.com/boltdb/bolt/cmd/bolt: from 2652945 to 2652945, ~
* github.com/gogits/gogs: from 23347590 to 23348082, ~
* github.com/spf13/hugo: from 15111102 to 15111540, ~
* github.com/influxdata/influxdb/cmd/influxd: from 16224307 to 16228895, ~
* github.com/nsqio/nsq/apps/nsqd: from 10008044 to 10008482, ~
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               189           190           +0.53%
BenchmarkMsgpUnmarshal-4             401           403           +0.50%
BenchmarkEasyJsonMarshal-4           1456          1466          +0.69%
BenchmarkEasyJsonUnmarshal-4         1546          1519          -1.75%
BenchmarkFlatBuffersMarshal-4        360           359           -0.28%
BenchmarkFlatBuffersUnmarshal-4      291           293           +0.69%
BenchmarkGogoprotobufMarshal-4       164           164           +0.00%
BenchmarkGogoprotobufUnmarshal-4     253           253           +0.00%

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

## GIT Log

```
commit a8c6c4837c17ab4ec3ba78e40b9a72dc70d9cf5a
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 8 21:09:09 2016 -0700

    os: document that the runtime can write to standard error
    
    Fixes #15970.
    
    Change-Id: I3f7d8316069a69d0e3859aaa96bc1414487fead0
    Reviewed-on: https://go-review.googlesource.com/23921
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 763883632e4d7fea145b6f3a7ee501b5ad9096f2
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 8 17:45:55 2016 -0700

    cmd/go: only run TestGoGetHTTPS404 where it works
    
    The test TestGoGetHTTPS404 downloads a package that does not build on
    every OS, so change it to only run where the package builds. It's not
    great for the test to depend on an external package, but this is an
    improvement on the current situation.
    
    Fixes #15644.
    
    Change-Id: I1679cee5ab1e61a5b26f4ad39dc8a397fbc0da69
    Reviewed-on: https://go-review.googlesource.com/23920
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit f3689d138256c36dbe0004459f79d570f5345f74
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Jun 6 16:00:33 2016 -0400

    cmd/compile: nilcheck interface value in go/defer interface call for SSA
    
    This matches the behavior of the legacy backend.
    
    Fixes #15975 (if this is the intended behavior)
    
    Change-Id: Id277959069b8b8bf9958fa8f2cbc762c752a1a19
    Reviewed-on: https://go-review.googlesource.com/23820
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0324a3f828d3a2f48751df2bbd54ad20499e598a
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Jun 8 16:11:44 2016 +0000

    runtime/cgo: restore the g pointer correctly in crosscall_s390x
    
    R13 needs to be set to g because C code may have clobbered R13.
    
    Fixes #16006.
    
    Change-Id: I66311fe28440e85e589a1695fa1c42416583b4c6
    Reviewed-on: https://go-review.googlesource.com/23910
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 09eedc32e13484d9171519e9f07f3210ba5a7afd
Author: Elias Naur <elias.naur@gmail.com>
Date:   Fri Jun 3 11:41:26 2016 +0200

    misc/android: make the exec wrapper exit code parsing more robust
    
    Before, the Android exec wrapper expected the trailing exit code
    output on its own line, like this:
    
    PASS
    exitcode=0
    
    However, some tests can sometimes squeeze in some output after
    the test harness outputs "PASS" and the newline. The
    TestWriteHeapDumpFinalizers test is particularly prone to this,
    since its finalizers println to standard out. When it happens, the
    output looks like this:
    
    PASS
    finalizedexitcode=0
    
    Two recent failures caused by this race:
    
    https://build.golang.org/log/185605e1b936142c22350eef22d20e982be53c29
    https://build.golang.org/log/e61cf6a050551d10360bd90be3c5f58c3eb07605
    
    Since the "exitcode=" string is always echoed after the test output,
    the fix is simple: instead of looking for the last newline in the
    output, look for the last exitcode string instead.
    
    Change-Id: Icd6e53855eeba60b982ad3108289d92549328b86
    Reviewed-on: https://go-review.googlesource.com/23750
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d1b5d08f341cbc702e1d2a3cd86ab1ad93a90c41
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jun 6 21:44:24 2016 -0700

    misc/cgo/testsanitizers: don't run some TSAN tests on GCC < 7
    
    Before GCC 7 defined __SANITIZE_THREAD__ when using TSAN,
    runtime/cgo/libcgo.h could not determine reliably whether TSAN was in
    use when using GCC.
    
    Fixes #15983.
    
    Change-Id: I5581c9f88e1cde1974c280008b2230fe5e971f44
    Reviewed-on: https://go-review.googlesource.com/23833
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit f605c77bbcc7946531e0914f13a0a14aae5f2991
Author: Andrew Gerrand <adg@golang.org>
Date:   Wed Jun 8 12:57:00 2016 +1000

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 313cf39 for CLs 23812 and 23880:
    
    	http2: GotFirstResponseByte hook should only fire once
    	http2: fix data race on pipe
    
    Fixes #16000
    
    Change-Id: I9c3f1b2528bbd99968aa5a0529ae9c5295979d1d
    Reviewed-on: https://go-review.googlesource.com/23881
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit afad74ec30c208f7cab08b7b80081adc7591dcb3
Author: Keith Randall <khr@golang.org>
Date:   Tue Jun 7 15:43:48 2016 -0700

    cmd/compile: cgen_append can handle complex targets
    
    Post-liveness fix, the slices on both sides can now be
    indirects of & variables.  The cgen code handles those
    cases just fine.
    
    Fixes #15988
    
    Change-Id: I378ad1d5121587e6107a9879c167291a70bbb9e4
    Reviewed-on: https://go-review.googlesource.com/23863
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 41dd1696ab13755bf7a129e0c73523ffb9fcbe66
Author: Keith Randall <khr@golang.org>
Date:   Tue Jun 7 09:54:09 2016 -0700

    cmd/compile: fix heap dump test on android
    
    go_android_exec is looking for "exitcode=" to decide the result
    of running a test.  The heap dump test nondeterministically prints
    "finalized" right at the end of the test.  When the timing is just
    right, we print "finalizedexitcode=0" and confuse go_android_exec.
    
    This failure happens occasionally on the android builders.
    
    Change-Id: I4f73a4db05d8f40047ecd3ef3a881a4ae3741e26
    Reviewed-on: https://go-review.googlesource.com/23861
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 2f088884aeab7f34ca69c0b8ab21c9694c628e19
Author: Keith Randall <khr@golang.org>
Date:   Mon Jun 6 10:56:42 2016 -0700

    cmd/compile: use fake package for allocating autos
    
    Make sure auto names don't conflict with function names. Before this CL,
    we confused name a.len (the len field of the slice a) with a.len (the function
    len declared on a).
    
    Fixes #15961
    
    Change-Id: I14913de697b521fb35db9a1b10ba201f25d552bb
    Reviewed-on: https://go-review.googlesource.com/23789
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 27c5dcffef590ac9dbc31f3d513f1b7d058907f6
Author: Andrew Gerrand <adg@golang.org>
Date:   Tue Jun 7 08:33:00 2016 +1000

    cmd/dist: use "set" instead of "export" in diagnostic message
    
    On Windows, "export" doesn't mean anything, but Windows users are the
    most likely to see this message.
    
    Fixes #15977
    
    Change-Id: Ib09ca08a7580713cacb65d0cdc43730765c377a8
    Reviewed-on: https://go-review.googlesource.com/23840
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3ba31558d1bca8ae6d2f03209b4cae55381175b3
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon Jun 6 10:23:49 2016 +1000

    net/http: send StatusOK on empty body with TimeoutHandler
    
    Fixes #15948
    
    Change-Id: Idd79859b3e98d61cd4e3ef9caa5d3b2524fd026a
    Reviewed-on: https://go-review.googlesource.com/23810
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a71af25401d68645ca23b2303ac6ae426739aa8b
Author: Andrew Gerrand <adg@golang.org>
Date:   Tue May 31 15:30:52 2016 +1000

    time: warn about correct use of a Timer's Stop/Reset methods
    
    Updates #14038
    Fixes #14383
    
    Change-Id: Icf6acb7c5d13ff1d3145084544c030a778482a38
    Reviewed-on: https://go-review.googlesource.com/23575
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f9b4556de01710a964ffd0513eb7574a2d1fd62c
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon Jun 6 10:41:47 2016 +1000

    net/http: send one Transfer-Encoding header when "chunked" set manually
    
    Fixes #15960
    
    Change-Id: I7503f6ede33e6a1a93cee811d40f7b297edf47bc
    Reviewed-on: https://go-review.googlesource.com/23811
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a871464e5aca9b81a6dc54cde8c31629387cb785
Author: Keith Randall <khr@golang.org>
Date:   Sun Jun 5 09:24:09 2016 -0700

    runtime: fix typo
    
    Fixes #15962
    
    Change-Id: I1949e0787f6c2b1e19b9f9d3af2f712606a6d4cf
    Reviewed-on: https://go-review.googlesource.com/23786
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0326e28f17ca760f76105fbcba9c5f55bb6ef1ce
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Jun 4 07:33:28 2016 +0900

    Revert "cmd/go: re-enable TestCgoConsistentResults on solaris"
    
    This reverts commit b89bcc1daeed9980c5ba8a255b37877493952874.
    
    Change-Id: Ief2f317ffc175f7e6002d0c39694876f46788c69
    Reviewed-on: https://go-review.googlesource.com/23744
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit b89bcc1daeed9980c5ba8a255b37877493952874
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri Jun 3 16:45:21 2016 +0900

    cmd/go: re-enable TestCgoConsistentResults on solaris
    
    Updates #13247.
    
    Change-Id: If5e4c9f4db05f58608b0eeed1a2312a04015b207
    Reviewed-on: https://go-review.googlesource.com/23741
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 7b48020cfeb64d1f841a7523aa841dbe53b3b465
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 1 15:24:14 2016 -0700

    cmd/cgo: check pointers for deferred C calls at the right time
    
    We used to check time at the point of the defer statement. This change
    fixes cgo to check them when the deferred function is executed.
    
    Fixes #15921.
    
    Change-Id: I72a10e26373cad6ad092773e9ebec4add29b9561
    Reviewed-on: https://go-review.googlesource.com/23650
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 55559f159e4ba7645a864c89caba0e29498425f9
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jun 3 12:39:54 2016 -0700

    doc/go1.7.html: html tidy
    
    Change-Id: I0e07610bae641cd63769b520089f5d854d796648
    Reviewed-on: https://go-review.googlesource.com/23770
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit adf261b7d6ed17bdefadc29ff71ddaba667b1a99
Author: Stephen McQuay (smcquay) <stephen@mcquay.me>
Date:   Fri Jun 3 02:12:17 2016 -0700

    cmd/go: match go-import package prefixes by slash
    
    The existing implementation for path collision resolution would
    incorrectly determine that:
    
        example.org/aa
    
    collides with:
    
        example.org/a
    
    This change splits by slash rather than comparing on a byte-by-byte
    basis.
    
    Fixes: #15947
    
    Change-Id: I18b3aaafbc787c81253203cf1328bb3c4420a0c4
    Reviewed-on: https://go-review.googlesource.com/23732
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit 4b64c53c034b3a99a8a5dc3e8081342e77048561
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Jun 3 11:27:33 2016 -0400

    reflect: clear tflag for StructOf type
    
    Fixes #15923
    
    Change-Id: I3e56564365086ceb0bfc15db61db6fb446ab7448
    Reviewed-on: https://go-review.googlesource.com/23760
    Reviewed-by: Sebastien Binet <seb.binet@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cf862478c89fd94c4fe8d9ce1cb481d71e5136bf
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jun 3 10:49:24 2016 -0700

    runtime/cgo: add TSAN locks around mmap call
    
    Change-Id: I806cc5523b7b5e3278d01074bc89900d78700e0c
    Reviewed-on: https://go-review.googlesource.com/23736
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit b4c7f6280ed00316e410261adbc804f6ddd209cc
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jun 3 10:09:08 2016 -0700

    doc/go1.7.html: add missing <code> and </a>
    
    Change-Id: I5f4bf89345dc139063dcf34da653e914386bcde6
    Reviewed-on: https://go-review.googlesource.com/23735
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6901b084824244122ea108eb7305295e44136be8
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jun 3 07:11:52 2016 -0700

    cmd/link: avoid name collision with DWARF .def suffix
    
    Adding a .def suffix for DWARF info collided with the DWARF info,
    without the suffix, for a method named def. Change the suffix to ..def
    instead.
    
    Fixes #15926.
    
    Change-Id: If1bf1bcb5dff1d7f7b79f78e3f7a3bbfcd2201bb
    Reviewed-on: https://go-review.googlesource.com/23733
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit c3bd93aa264383c0c7928516ca102a225c83ea23
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri Jun 3 18:06:54 2016 +0900

    net: don't leak test helper goroutine in TestAcceptTimeout
    
    Fixes #15109.
    
    Change-Id: Ibfdedd6807322ebec84bacfeb492fb53fe066960
    Reviewed-on: https://go-review.googlesource.com/23742
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>

commit 9e112a3fe4c001530184c2edc918a854d0b6d7e4
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Thu May 26 10:54:25 2016 +0200

    bytes: use Run method for benchmarks
    
    Change-Id: I34ab1003099570f0ba511340e697a648de31d08a
    Reviewed-on: https://go-review.googlesource.com/23427
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 26849746c9c7ca290d6cbb7ca5f3cf71c971e980
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Thu Jun 2 11:07:55 2016 +1200

    cmd/internal/obj, runtime: fixes for defer in 386 shared libraries
    
    Any defer in a shared object crashed when GOARCH=386. This turns out to be two
    bugs:
    
     1) Calls to morestack were not processed to be PIC safe (must have been
        possible to trigger this another way too)
     2) jmpdefer needs to rewind the return address of the deferred function past
        the instructions that load the GOT pointer into BX, not just past the call
    
    Bug 2) requires re-introducing the a way for .s files to know when they are
    being compiled for dynamic linking but I've tried to do that in as minimal
    a way as possible.
    
    Fixes #15916
    
    Change-Id: Ia0d09b69ec272a176934176b8eaef5f3bfcacf04
    Reviewed-on: https://go-review.googlesource.com/23623
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5799973c3e25545ac0e7d20f32a1453531c69399
Author: Andrew Gerrand <adg@golang.org>
Date:   Fri Jun 3 11:33:37 2016 +1000

    cmd/go: fix staleness test for releases, also deflake it
    
    Fixes #15933
    
    Change-Id: I2cd6365e6d0ca1cafdc812fbfaaa55aa64b2b289
    Reviewed-on: https://go-review.googlesource.com/23731
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7825ca6a63502e3b8decb0b569513dd6b5954aa7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jun 2 12:33:34 2016 -0700

    doc/go1.7.html: net/mail.ParseAddress is stricter
    
    Fixes #15940.
    
    Change-Id: Ie6da6fef235c6a251caa96d45f606c05d118a0ac
    Reviewed-on: https://go-review.googlesource.com/23710
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Travis Beatty <travisby@gmail.com>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 92cd6e3af9f423ab4d8ac78f24e7fd81c31a8ce6
Author: David Glasser <glasser@meteor.com>
Date:   Tue May 31 12:28:57 2016 -0700

    encoding/json: fix docs on valid key names
    
    This has been inaccurate since https://golang.org/cl/6048047.
    
    Fixes #15317.
    
    Change-Id: If93d2161f51ccb91912cb94a35318cf33f4d526a
    Reviewed-on: https://go-review.googlesource.com/23691
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 49c680f948310cfc7ab3062ca9a96a4adb6ae8cd
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Jun 2 17:17:02 2016 +0900

    syscall: deflake TestUnshare
    
    Change-Id: I21a08c2ff5ebb74e158723cca323574432870ba8
    Reviewed-on: https://go-review.googlesource.com/23662
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0f5697a81deab54f1673a48bd0ce613ebf1ddae6
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Thu May 26 12:10:28 2016 +0200

    strconv: use Run for some benchmarks
    
    This serves as an example of table-driven benchmarks which are analoguous to the common pattern for table-driven tests.
    
    Change-Id: I47f94c121a7117dd1e4ba03b3f2f8bcb5da38063
    Reviewed-on: https://go-review.googlesource.com/23470
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 03abde49713b46366fa47a037040697c8fdad3bb
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jun 2 12:01:03 2016 -0700

    runtime: only permit SetCgoTraceback to be called once
    
    Accept a duplicate call, but nothing else.
    
    Change-Id: Iec24bf5ddc3b0f0c559ad2158339aca698601743
    Reviewed-on: https://go-review.googlesource.com/23692
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 88e0ec2979bb39bd8811ec50a69fcb5007a24623
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue May 31 16:04:00 2016 -0700

    runtime/cgo: avoid races on cgo_context_function
    
    Change-Id: Ie9e6fda675e560234e90b9022526fd689d770818
    Reviewed-on: https://go-review.googlesource.com/23610
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 853cd1f4a61396cccb91522ed59af52d61aa8371
Author: Alexander Morozov <lk4d4math@gmail.com>
Date:   Fri May 27 15:02:31 2016 -0700

    syscall: call setgroups for no groups on GNU/Linux
    
    Skip setgroups only for one particular case: GidMappings != nil and
    GidMappingsEnableSetgroup == false and list of supplementary groups is
    empty.
    This patch returns pre-1.5 behavior for simple exec and still allows to
    use GidMappings with non-empty Credential.
    
    Change-Id: Ia91c77e76ec5efab7a7f78134ffb529910108fc1
    Reviewed-on: https://go-review.googlesource.com/23524
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e90a49a0f5380c6f68502b1febfb73f696c2f610
Author: Steve Phillips <steve@tryingtobeawesome.com>
Date:   Thu Jun 2 02:40:37 2016 -0700

    doc/go1.7.html: typo fix; replace "," at end of sentence with "."
    
    Signed-off-by: Steven Phillips <steve@tryingtobeawesome.com>
    
    Change-Id: Ie7c3253a5e1cd43be8fa12bad340204cc6c5ca76
    Reviewed-on: https://go-review.googlesource.com/23677
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 04888c9770560e99de63fafdfc2ce39b47844bfd
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Thu Jun 2 14:34:37 2016 +0200

    doc/go1.7: fix typo in nsswitch.conf name
    
    Fixes #15939
    
    Change-Id: I120cbeac73a052fb3f328774e6d5e1534f11bf6b
    Reviewed-on: https://go-review.googlesource.com/23682
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 905ced0e6ae33dca7ae6aa984d50cb02287952b1
Author: Sebastien Binet <seb.binet@gmail.com>
Date:   Thu Jun 2 09:25:30 2016 +0200

    reflect: document StructOf embedded fields limitation
    
    This CL documents that StructOf currently does not generate wrapper
    methods for embedded fields.
    
    Updates #15924
    
    Change-Id: I932011b1491d68767709559f515f699c04ce70d4
    Reviewed-on: https://go-review.googlesource.com/23681
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
```
