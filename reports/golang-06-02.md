# June 2, 2016 Report

Number of commits: 88

## Compilation time

* github.com/coreos/etcd: from 14.706641003s to 14.430088826s, -1.88%
* github.com/boltdb/bolt/cmd/bolt: from 546.389751ms to 583.045047ms, +6.71%
* github.com/gogits/gogs: from 12.551555069s to 12.717827007s, +1.32%
* github.com/spf13/hugo: from 6.432768583s to 6.643672829s, +3.28%
* github.com/influxdata/influxdb/cmd/influxd: from 6.396162823s to 6.558451873s, +2.54%
* github.com/nsqio/nsq/apps/nsqd: from 2.158815757s to 2.16409228s, +0.24%
* github.com/mholt/caddy: from 4.636343412s to 4.813305357s, +3.82%

## Binary size:

* github.com/coreos/etcd: from 22155240 to 22583488, +1.93%
* github.com/boltdb/bolt/cmd/bolt: from 2586078 to 2652945, +2.59%
* github.com/gogits/gogs: from 22929471 to 23339270, +1.79%
* github.com/spf13/hugo: from 14798695 to 15093806, +1.99%
* github.com/influxdata/influxdb/cmd/influxd: from 15903871 to 16207174, +1.91%
* github.com/nsqio/nsq/apps/nsqd: from 9811237 to 10008044, +2.01%
* github.com/mholt/caddy: from 12778119 to 13044558, +2.09%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               188           192           +2.13%
BenchmarkMsgpUnmarshal-4             394           406           +3.05%
BenchmarkEasyJsonMarshal-4           1413          1464          +3.61%
BenchmarkEasyJsonUnmarshal-4         1532          1520          -0.78%
BenchmarkFlatBuffersMarshal-4        346           367           +6.07%
BenchmarkFlatBuffersUnmarshal-4      285           293           +2.81%
BenchmarkGogoprotobufMarshal-4       162           164           +1.23%
BenchmarkGogoprotobufUnmarshal-4     245           253           +3.27%

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
* [build: enable framepointer mode by default](https://github.com/golang/go/commit/7fdec6216c0a25c6dbcc8159b755da6682dd9080)


## GIT Log

```
commit 14968bc1e52842b098408516472ebd3fb97e4714
Author: Elias Naur <elias.naur@gmail.com>
Date:   Thu Jun 2 15:42:14 2016 +0200

    cmd/dist: skip an unsupported test on darwin/arm
    
    Fixes the darwin/arm builder (I hope)
    
    Change-Id: I8a3502a1cdd468d4bf9a1c895754ada420b305ce
    Reviewed-on: https://go-review.googlesource.com/23684
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 6c4f8cd0d17c8147319effcffcb608fc42eaf307
Author: Elias Naur <elias.naur@gmail.com>
Date:   Thu Jun 2 15:00:34 2016 +0200

    misc/cgo/test: fix issue9400 test on android/386
    
    The test for #9400 relies on an assembler function that manipulates
    the stack pointer. Meanwile, it uses a global variable for
    synchronization. However, position independent code on 386 use a
    function call to fetch the base address for global variables.
    That function call in turn overwrites the Go stack.
    
    Fix that by fetching the global variable address once before the
    stack register manipulation.
    
    Fixes the android/386 builder.
    
    Change-Id: Ib77bd80affaa12f09d582d09d8b84a73bd021b60
    Reviewed-on: https://go-review.googlesource.com/23683
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 42c51debe824bd9b1fd93b3d50ff7187530754d3
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Jun 1 22:51:30 2016 +0200

    misc/cgo/test,cmd/dist: enable (more) Cgo tests on iOS
    
    For #15919
    
    Change-Id: I9fc38d9c8a9cc9406b551315e1599750fe212d0d
    Reviewed-on: https://go-review.googlesource.com/23635
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ba22172832a971f0884106a5a8ff26a98a65623c
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Thu Jun 2 07:43:21 2016 +0200

    runtime: fix typo in comment
    
    Change-Id: I82e35770b45ccd1433dfae0af423073c312c0859
    Reviewed-on: https://go-review.googlesource.com/23680
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 15db3654b8e9fda6d41f4389879c8cd370f71a7e
Author: Anmol Sethi <anmol@aubble.com>
Date:   Wed Jun 1 22:35:09 2016 -0400

    net/http: http.Request.Context doc fix
    
    The comment on http.Request.Context says that the context
    is canceled when the client's connection closes even though
    this has not been implemented. See #15927
    
    Change-Id: I50b68638303dafd70f77f8f778e6caff102d3350
    Reviewed-on: https://go-review.googlesource.com/23672
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 6a0fd18016794a681580c8ca971c7d2d26f287bf
Author: Andrew Gerrand <adg@golang.org>
Date:   Thu Jun 2 14:31:16 2016 +1000

    doc: mention net/http/httptrace package in release notes
    
    Updates #15810
    
    Change-Id: I689e18409a88c9e8941aa2e98f472c331efd455e
    Reviewed-on: https://go-review.googlesource.com/23674
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 17f7461ed66bbc66fef02ef7ca6901d116b6ff3d
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 1 21:09:58 2016 -0700

    doc/go1.7.html: fix spelling of cancelation
    
    We say "cancelation," not "cancellation."
    
    Fixes #15928.
    
    Change-Id: I66d545404665948a27281133cb9050eebf1debbb
    Reviewed-on: https://go-review.googlesource.com/23673
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d25c3eadea9bc5c8b6451c3502d6063dd618a3af
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri May 27 15:41:55 2016 +1200

    cmd/compile: do not generate tail calls when dynamic linking on ppc64le
    
    When a wrapper method calls the real implementation, it's not possible to use a
    tail call when dynamic linking on ppc64le. The bad scenario is when a local
    call is made to the wrapper: the wrapper will call the implementation, which
    might be in a different module and so set the TOC to the appropriate value for
    that module. But if it returns directly to the wrapper's caller, nothing will
    reset it to the correct value for that function.
    
    Change-Id: Icebf24c9a2a0a9a7c2bce6bd6f1358657284fb10
    Reviewed-on: https://go-review.googlesource.com/23468
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 068c745e1e44875c411de5d5aea3f96574fbee12
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Jun 2 08:53:11 2016 +0900

    vendor: update vendored route
    
    Updates golang.org/x/net/route to rev fac978c for:
    - route: fix typos in test
    
    Change-Id: I35de1d3f8e887c6bb5fe50e7299f2fc12e4426de
    Reviewed-on: https://go-review.googlesource.com/23660
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3c6b6684ce21c1092ba208a0f1744ad7c930248a
Author: Andrew Gerrand <adg@golang.org>
Date:   Thu Jun 2 10:00:23 2016 +1000

    api: promote next.txt to go1.7.txt and update api tool
    
    Change-Id: Idb348be00f949da553aa6aab62836f59dfee298d
    Reviewed-on: https://go-review.googlesource.com/23671
    Reviewed-by: Chris Broadfoot <cbro@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>

commit 36358b16062fb419a4c78b3f03b24106cb057222
Author: Andrew Gerrand <adg@golang.org>
Date:   Thu Jun 2 09:59:06 2016 +1000

    api: remove os.File.Size and http.Transport.Dialer
    
    This method and field were added and then later removed during the 1.7
    development cycle.
    
    Change-Id: I0482a6356b91d2be67880b44ef5d8a1daab49ec8
    Reviewed-on: https://go-review.googlesource.com/23670
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit d7ae8b3c11b027721f0878caac0620ccb7f81048
Author: Andrew Gerrand <adg@golang.org>
Date:   Thu Jun 2 09:24:43 2016 +1000

    api: update next.txt
    
    Change-Id: I04da6a56382d3bd96e3c849a022618553039b2db
    Reviewed-on: https://go-review.googlesource.com/23651
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit 2a8c81ffaadc69add6ff85b241691adb7f9f24ff
Author: Adam Langley <agl@golang.org>
Date:   Wed Jun 1 14:41:09 2016 -0700

    crypto/tls: buffer handshake messages.
    
    This change causes TLS handshake messages to be buffered and written in
    a single Write to the underlying net.Conn.
    
    There are two reasons to want to do this:
    
    Firstly, it's slightly preferable to do this in order to save sending
    several, small packets over the network where a single one will do.
    
    Secondly, since 37c28759ca46cf381a466e32168a793165d9c9e9 errors from
    Write have been returned from a handshake. This means that, if a peer
    closes the connection during a handshake, a “broken pipe” error may
    result from tls.Conn.Handshake(). This can mask any, more detailed,
    fatal alerts that the peer may have sent because a read will never
    happen.
    
    Buffering handshake messages means that the peer will not receive, and
    possibly reject, any of a flow while it's still being written.
    
    Fixes #15709
    
    Change-Id: I38dcff1abecc06e52b2de647ea98713ce0fb9a21
    Reviewed-on: https://go-review.googlesource.com/23609
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f6c0241999bffe0fe52e8b7f5bbcc8f9e02edbdf
Author: Tom Bergan <tombergan@google.com>
Date:   Fri May 27 16:53:13 2016 -0700

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 6bdd4be4 for CL 23526:
    
      http2: GotFirstResponseByte hook should only fire once
    
    Also updated the trace hooks test to verify that all trace hooks are called
    exactly once except ConnectStart/End, which may be called multiple times (due
    to happy-eyeballs).
    
    Fixes #15777
    
    Change-Id: Iea5c64eb322b58be27f9ff863b3a6f90e996fa9b
    Reviewed-on: https://go-review.googlesource.com/23527
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit d0c1888739ac0d5d0c9f82a4b86945c0351caef6
Author: Andrew Gerrand <adg@golang.org>
Date:   Thu Jun 2 08:11:01 2016 +1000

    doc: revert copyright date to 2009
    
    Somehow this date was changed in error (by me) to 2012.
    It should have always been 2009.
    
    Change-Id: I87029079458d4c4eeeff2f2fc0574f10afa9af09
    Reviewed-on: https://go-review.googlesource.com/23622
    Reviewed-by: Rob Pike <r@golang.org>

commit 6de014b9e2a655e093c2e3b5617a90b97d66c152
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Jun 1 20:58:02 2016 +0200

    misc/cgo/test,cmd/dist: enable (most) Cgo tests on Android
    
    Some tests cannot build for Android; use build tags and stubs to
    skip them.
    
    For #15919
    
    Change-Id: Ieedcb73d4cabe23c3775cfb1d44c1276982dccd9
    Reviewed-on: https://go-review.googlesource.com/23634
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit bbd1dcdf7da68a3759a2d86f851391c1ec974f77
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Jun 1 13:46:49 2016 -0700

    cmd/compile: correctly export underlying type of predecl. error type
    
    Fixes #15920.
    
    Change-Id: I78cd79b91a58d0f7218b80f9445417f4ee071a6e
    Reviewed-on: https://go-review.googlesource.com/23606
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5db44c17a2391bbdfbc3c04e83e66025ca5dea3d
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Wed Jun 1 19:16:56 2016 +0200

    math/big: avoid panic in float.Text with negative prec
    
    Fixes #15918
    
    Change-Id: I4b434aed262960a2e6c659d4c2296fbf662c3a52
    Reviewed-on: https://go-review.googlesource.com/23633
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 77026ef902d3fa21597400d230701979bc1f0efc
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun May 22 02:20:11 2016 -0700

    runtime: document heap scavenger memory summary
    
    Fixes #15212.
    
    Change-Id: I2628ec8333330721cddc5145af1ffda6f3e0c63f
    Reviewed-on: https://go-review.googlesource.com/23319
    Reviewed-by: Austin Clements <austin@google.com>

commit bc4fdfdbfe6b971fcceaf4d75514a882917df10d
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 1 09:31:31 2016 -0700

    os/signal: deflake TestReset/TestIgnore
    
    Fixes #15661.
    
    Change-Id: Ic3a8296fc7107f491880900ef52563e52caca1a3
    Reviewed-on: https://go-review.googlesource.com/23615
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 52fe47247217f5126dacc5a8c8e80b85d2fb25c6
Author: Keith Randall <khr@golang.org>
Date:   Tue May 31 14:55:12 2016 -0700

    cmd/compile: for arm, zero unaligned memory 1 byte at a time
    
    If memory might be unaligned, zero it one byte at a time
    instead of 4 bytes at a time.
    
    Fixes #15902
    
    Change-Id: I4eff0840e042e2f137c1a4028f08793eb7dfd703
    Reviewed-on: https://go-review.googlesource.com/23587
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit e29e0ba19af26c30c95b59aeda482e60ae594113
Author: David du Colombier <0intro@gmail.com>
Date:   Wed Jun 1 15:13:55 2016 +0200

    cmd/compile: fix TestAssembly on Plan 9
    
    Since CL 23620, TestAssembly is failing on Plan 9.
    
    In CL 23620, the process environment is passed to 'go tool compile'
    after setting GOARCH. On Plan 9, if GOARCH is already set in the
    process environment, it would take precedence. On Unix, it works
    as expected because the first GOARCH found takes precedence.
    
    This change uses the mergeEnvLists function from cmd/go/main.go
    to merge the two environment lists such that variables with the
    same name in "in" replace those in "out".
    
    Change-Id: Idee22058343932ee18666dda331c562c89c33507
    Reviewed-on: https://go-review.googlesource.com/23593
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bd2dc2d819b85beb8887466a165242e2d540e4b9
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Wed Jun 1 09:44:38 2016 -0300

    doc: rename Unshare to Unshareflags in go1.7 release notes
    
    Implementation changed in https://golang.org/cl/23612.
    
    Updates #15810
    
    Change-Id: I8fff9e3aa3e54162546bb9ec1cc2ebba2b6d9fed
    Reviewed-on: https://go-review.googlesource.com/23614
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ad074e205e4aa4c7762e223df65695d5157b0c4e
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Thu May 26 11:46:19 2016 +0200

    regexp: use Run for benchmark
    
    Change-Id: I4d19e3221d3789d4c460b421b2d1484253778068
    Reviewed-on: https://go-review.googlesource.com/23429
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>

commit 88ae6495d086ed5b0acb94d5adc49434ec47a675
Author: Alexander Morozov <lk4d4math@gmail.com>
Date:   Tue May 31 19:44:48 2016 -0700

    syscall: rename SysProcAttr.Unshare to Unshareflags
    
    For symmetry with Cloneflags and it looks slightly weird because there
    is syscall.Unshare method.
    
    Change-Id: I3d710177ca8f27c05b344407f212cbbe3435094b
    Reviewed-on: https://go-review.googlesource.com/23612
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 2885e07c259ffda336d6965fcca03b4df617d812
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Wed Jun 1 13:32:53 2016 +1200

    cmd/compile: pass process env to 'go tool compile' in compileToAsm
    
    In particular, this stops the test failing when GOROOT and GOROOT_FINAL are
    different.
    
    Change-Id: Ibf6cc0a173f1d965ee8aa31eee2698b223f1ceec
    Reviewed-on: https://go-review.googlesource.com/23620
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 04acd625d7a1044c8ca78464f6727276577ffb3d
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Tue May 31 22:30:37 2016 +0100

    context: fix typo in comments
    
    Change-Id: I41310ec88c889fda79d80eaf4a742a1000284f60
    Reviewed-on: https://go-review.googlesource.com/23591
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fe62a9ee872d4f61a47cc4e8c7bc0fb67cc4ebb6
Author: Robert Griesemer <gri@golang.org>
Date:   Fri May 27 19:47:55 2016 -0700

    crypto/tls: remove unused variable in benchmark code
    
    This fixes `go test go/types`.
    
    https://golang.org/cl/23487/ introduced this code which contains
    two unused variables (declared and assigned to, but never read).
    cmd/compile doesn't report the error due open issue #8560 (the
    variables are assigned to in a closure), but go/types does. The
    build bot only runs go/types tests in -short mode (which doesn't
    typecheck the std lib), hence this doesn't show up on the dashboard
    either.
    
    We cannot call b.Fatal and friends in the goroutine. Communicating
    the error to the invoking function requires a channel or a mutex.
    Unless the channel/sycnhronized variable is tested in each iteration
    that follows, the iteration blocks if there's a failure. Testing in
    each iteration may affect benchmark times.
    
    One could use a time-out but that time depends on the underlying system.
    Panicking seems good enough in this unlikely case; better than hanging
    or affecting benchmark times.
    
    Change-Id: Idce1172da8058e580fa3b3e398825b0eb4316325
    Reviewed-on: https://go-review.googlesource.com/23528
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e106122200032cdf4f0a993cdd89c7531aaf8d75
Author: Keith Randall <khr@golang.org>
Date:   Tue May 31 13:48:29 2016 -0700

    cmd/compile: test non-constant shifts
    
    Test all the weird shifts, like int8 shifted right by uint16.
    Increases coverage for shift lowerings in AMD64.rules.
    
    Change-Id: I066fe6ad6bfc05253a8d6a2ee17ff244d3a7652e
    Reviewed-on: https://go-review.googlesource.com/23585
    Run-TryBot: Todd Neal <todd@tneal.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Todd Neal <todd@tneal.org>

commit 3d81d4adc9419e2fcba5888ab074d3f17eb5eb03
Author: Robert Griesemer <gri@golang.org>
Date:   Tue May 31 13:32:34 2016 -0700

    spec: document that duplicate types are invalid in type switches
    
    Both compilers and also go/types don't permit duplicate types in
    type switches; i.e., this spec change is documenting a status quo
    that has existed for some time.
    
    Furthermore, duplicate nils are not accepted by gccgo or go/types;
    and more recently started causing a compiler error in gc. Permitting
    them is inconsistent with the existing status quo.
    
    Rather than making it an implementation restriction (as we have for
    expression switches), this is a hard requirement since it was enforced
    from the beginning (except for duplicate nils); it is also a well
    specified requirement that does not pose a significant burden for
    an implementation.
    
    Fixes #15896.
    
    Change-Id: If12db5bafa87598b323ea84418cb05421e657dd8
    Reviewed-on: https://go-review.googlesource.com/23584
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3659645cb1f32d7b1eeefdb65f1339fe54f0f6eb
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue May 31 11:23:50 2016 -0700

    flag: recognize "0s" as the zero value for a flag.Duration
    
    Implemented by using a reflect-based approach to recognize the zero
    value of any non-interface type that implements flag.Value.  Interface
    types will fall back to the old code.
    
    Fixes #15904.
    
    Change-Id: I594c3bfb30e9ab1aca3e008ef7f70be20aa41a0b
    Reviewed-on: https://go-review.googlesource.com/23581
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 8003e791549f011edcc2a9d1eacbd5674826d38c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue May 31 15:01:05 2016 -0700

    test: add more switch error handling tests
    
    Some of these errors are reported in the wrong places.
    That’s issue #15911 and #15912.
    
    Change-Id: Ia09d7f89be4d15f05217a542a61b6ac08090dd87
    Reviewed-on: https://go-review.googlesource.com/23588
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1846c632ee37681eb92fb27f7071a58bdf6d7a3c
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue May 31 20:48:42 2016 +1200

    cmd/go: combine gccgo's ld and ldShared methods
    
    This fixes handling of cgo flags and makes sure packages that are only
    implicitly included in the shared library are passed to the link.
    
    Fixes #15885
    
    Change-Id: I1e8a72b5314261973ca903c78834700fb113dde9
    Reviewed-on: https://go-review.googlesource.com/23537
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 690de51ffac1473820212c88a11685b40f7bde3b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue May 31 14:07:38 2016 -0700

    runtime: fix restoring PC in ARM version of cgocallback_gofunc
    
    Fixes #15856.
    
    Change-Id: Ia8def161642087e4bd92a87298c77a0f9f83dc86
    Reviewed-on: https://go-review.googlesource.com/23586
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 3d037cfaf8c70b8af87cb5d57553a7e3e9dc2117
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri May 27 10:05:52 2016 -0700

    runtime: pass signal context to cgo traceback function
    
    When doing a backtrace from a signal that occurs in C code compiled
    without using -fasynchronous-unwind-tables, we have to rely on frame
    pointers. In order to do that, the traceback function needs the signal
    context to reliably pick up the frame pointer.
    
    Change-Id: I7b45930fced01685c337d108e0f146057928f876
    Reviewed-on: https://go-review.googlesource.com/23494
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c52dff0727c58cb7a6e768d91d15e3eaafcb420a
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue May 31 12:42:34 2016 -0700

    doc/go1.7.html: make RFC an actual link
    
    Change-Id: I5e8dad0c2534b5c3654cf0a0b51a38186d627a3c
    Reviewed-on: https://go-review.googlesource.com/23582
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 2256e38978a38a954df72ab50423c1883f1063d7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri May 27 14:33:23 2016 -0700

    runtime: update pprof binary header URL
    
    The code has moved from code.google.com to github.com.
    
    Change-Id: I0cc9eb69b3fedc9e916417bc7695759632f2391f
    Reviewed-on: https://go-review.googlesource.com/23523
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 2c1791b13b27dfc69adf2d19ecf9a180d089cd22
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue May 31 19:24:30 2016 +1200

    cmd/link: suppress PIE whenever externally linking with a sanitizer
    
    golang.org/issue/15443 complained that a race-enabled PIE binary crashed at
    startup, but other ways of linking in tsan (or other sanitizers) such as
    
     #cgo CFLAGS: -fsanitize=thread
     #cgo LDFLAGS: -fsanitize=thread
    
    have the same problem. Pass -no-pie to the host linker (if supported) if any
    -fsanitizer=foo cgo LDFLAG is seen when linking.
    
    Fixes #15887
    
    Change-Id: Id799770f8d045f6f40fa8c463563937a5748d1a8
    Reviewed-on: https://go-review.googlesource.com/23535
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 2a6544b604b61898d666ffbe456ccde720c04577
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue May 31 20:05:39 2016 +1200

    cmd/go, cmd/link: set LC_ALL=C when checking if host compiler supports -no-pie
    
    Fixes #15900
    
    Change-Id: Ieada5f4e3b3b2ae358414e013f3090b4b820569b
    Reviewed-on: https://go-review.googlesource.com/23536
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 66736880ca2e50fc7c5428a171fbbe6d344a853b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu May 26 17:47:03 2016 -0700

    runtime/cgo: add TSAN acquire/release calls
    
    Add TSAN acquire/release calls to runtime/cgo to match the ones
    generated by cgo.  This avoids a false positive race around the malloc
    memory used in runtime/cgo when other goroutines are simultaneously
    calling malloc and free from cgo.
    
    These new calls will only be used when building with CGO_CFLAGS and
    CGO_LDFLAGS set to -fsanitize=thread, which becomes a requirement to
    avoid all false positives when using TSAN.  These are needed not just
    for runtime/cgo, but also for any runtime package that uses cgo (such as
    net and os/user).
    
    Add an unused attribute to the _cgo_tsan_acquire and _cgo_tsan_release
    functions, in case there are no actual cgo function calls.
    
    Add a test that checks that setting CGO_CFLAGS/CGO_LDFLAGS avoids a
    false positive report when using os/user.
    
    Change-Id: I0905c644ff7f003b6718aac782393fa219514c48
    Reviewed-on: https://go-review.googlesource.com/23492
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 0e13dbc1a91cbe00e3c83a55f56db69380fe8f68
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 30 16:42:38 2016 -0700

    cmd/compile: disallow multiple nil cases in a type switch
    
    Fixes #15898.
    
    Change-Id: I66e2ad21f283563c7142aa820f0354711d964768
    Reviewed-on: https://go-review.googlesource.com/23573
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 7cd6cae6a63f09caa88bbcc394053b40cdeeccd1
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sun May 29 00:47:45 2016 -0700

    compress/flate: use seperate const block for exported constants
    
    As rendered on https://tip.golang.org/pkg/compress/flate/, there is an
    extra new-line because of the unexported constants in the same block.
    
    <<<
    const (
        NoCompression      = 0
        BestSpeed          = 1
        BestCompression    = 9
        DefaultCompression = -1
        HuffmanOnly        = -2 // Disables match search and only does Huffman entropy reduction.
    
    )
    >>>
    
    Instead, seperate the exported compression level constants into its own
    const block. This is both more readable and also fixes the issue.
    
    Change-Id: I60b7966c83fb53356c02e4640d05f55a3bee35b7
    Reviewed-on: https://go-review.googlesource.com/23557
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4223294eab3dee0f6c03fd57fc24be3dc3e2d53a
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri May 27 16:03:44 2016 -0700

    runtime/pprof, cmd/pprof: fix profiling for PIE
    
    In order to support pprof for position independent executables, pprof
    needs to adjust the PC addresses stored in the profile by the address at
    which the program is loaded. The legacy profiling support which we use
    already supports recording the GNU/Linux /proc/self/maps data
    immediately after the CPU samples, so do that. Also change the pprof
    symbolizer to use the information, if available, when looking up
    addresses in the Go pcline data.
    
    Fixes #15714.
    
    Change-Id: I4bf679210ef7c51d85cf873c968ce82db8898e3e
    Reviewed-on: https://go-review.googlesource.com/23525
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 87ee12cece96ec5837fe89c37899d725e7e852d9
Author: Andrew Gerrand <adg@golang.org>
Date:   Tue May 31 13:21:35 2016 +1000

    crypto/tls: reduce size of buffer in throughput benchmarks
    
    The Windows builders run the throughput benchmarks really slowly with a
    64kb buffer. Lowering it to 16kb brings the performance back into line
    with the other builders.
    
    This is a work-around to get the build green until we can figure out why
    the Windows builders are slow with the larger buffer size.
    
    Update #15899
    
    Change-Id: I215ebf115e8295295c87f3b3e22a4ef1f9e77f81
    Reviewed-on: https://go-review.googlesource.com/23574
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8e6e9e83818596048cfca1e121ad39b9b007ba28
Author: Richard Miller <miller.research@gmail.com>
Date:   Sat May 28 10:06:37 2016 +0100

    syscall: plan9 - mark gbit16 as go:nosplit
    
    This is a correction to CL 22610.  The gbit16 function is called in
    StartProcess between fork and exec, and therefore must not split the
    stack.  Normally it's inlined so this is not an issue, but on one
    occasion I've observed it to be compiled without inlining, and the
    result was a panic.  Mark it go:nosplit to be safe.
    
    Change-Id: I0381754397b766431bf406d9767c73598d23b901
    Reviewed-on: https://go-review.googlesource.com/23560
    Reviewed-by: David du Colombier <0intro@gmail.com>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b3f98d7a42d04f6f173cd61ce3fe2106e4877496
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon May 30 15:17:14 2016 +1000

    sync: document that RWMutex read locks may not be held recursively
    
    Fixes #15418
    
    Change-Id: Ibc51d602eb28819d0e44e5ca13a5c61573e4111c
    Reviewed-on: https://go-review.googlesource.com/23570
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit d2c92f8453cab8d042b794c8ce398f6ff8e6f650
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 30 12:57:20 2016 -0700

    path/filepath: prevent infinite recursion on Windows on UNC input
    
    This is a minimal fix to prevent this and
    other possible future infinite recursion.
    We can put in a proper fix for UNC in Go 1.8.
    
    Updates #15879
    
    Change-Id: I3653cf5891bab8511adf66fa3c1a1d8912d1a293
    Reviewed-on: https://go-review.googlesource.com/23572
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 81a8f1a7946c28eaa9c187b7aaa349020b7a9ba4
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon May 30 15:17:14 2016 +1000

    doc: remove remnant mention of io.SizedReaderAt from Go 1.7 docs
    
    Updates #15810
    
    Change-Id: I37f14a0ed1f5ac24ea2169a7e65c0469bfddd928
    Reviewed-on: https://go-review.googlesource.com/23559
    Reviewed-by: Michael McGreevy <mcgreevy@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit dc5b5239e8020ca0b366ba02f99fe87728fa290c
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat May 28 03:06:33 2016 +0900

    net: don't call forceCloseSockets in non-TestMain functions
    
    forceCloseSockets is just designed as a kingston valve for TestMain
    function and is not suitable to keep track of inflight sockets.
    
    Fixes #15525.
    
    Change-Id: Id967fe5b8da99bb08b699cc45e07bbc3dfc3ae3d
    Reviewed-on: https://go-review.googlesource.com/23505
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4e01c132d03ee1f862ae8ba9db465d6047f950f2
Author: Augusto Roman <aroman@gmail.com>
Date:   Sat May 28 16:59:28 2016 -0700

    doc: correct release notes for non-string map keys in encoding/json
    
    The original draft mentioned support for json.Marshaler, but that's
    not the case.  JSON supports only string keys (not arbitrary JSON)
    so only encoding.TextMarshaller is supported.
    
    Change-Id: I7788fc23ac357da88e92aa0ca17b513260840cee
    Reviewed-on: https://go-review.googlesource.com/23529
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 17396575135ba71472ab8a5b82f31af50d8bf312
Author: Keith Randall <khr@golang.org>
Date:   Sat May 28 21:15:24 2016 -0700

    cmd/compile: shift tests, fix triple-shift rules
    
    Add a bunch of tests for shifts.
    
    Fix triple-shift rules to always take constant shifts as 64 bits.
    (Earlier rules always promote shift amounts to 64 bits.)
    Add overflow checks.
    
    Increases generic rule coverage to 91%
    
    Change-Id: I6b42d368d19d36ac482dbb8e0d4f67e30ad7145d
    Reviewed-on: https://go-review.googlesource.com/23555
    Reviewed-by: Todd Neal <todd@tneal.org>

commit 79f7ccf2c3931745aeb97c5c985b6ac7b44befb4
Author: Keith Randall <khr@golang.org>
Date:   Sat May 28 21:59:17 2016 -0700

    cmd/compile: add constant fold comparison tests
    
    Increases generic.rules coverage from 91% to 95%.
    
    Change-Id: I981eb94f3cd10d2f87c836576a43786787a25d83
    Reviewed-on: https://go-review.googlesource.com/23556
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Todd Neal <todd@tneal.org>

commit b859a78e0a71d769274dac8cf0108bdf41ec55a5
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sat May 28 02:14:25 2016 -0700

    io: use SeekStart, SeekCurrent, and SeekEnd in io.Seeker documentation
    
    The documentation previously used C style enumerations: 0, 1, 2.
    While this is pretty much universally correct, it does not help a user
    become aware of the existence of the SeekStart, SeekCurrent, and SeekEnd
    constants. Thus, we should use them in the documentation to direct people's
    attention to them.
    
    Updates #6885
    
    Change-Id: I44b5e78d41601c68a0a1c96428c853df53981d52
    Reviewed-on: https://go-review.googlesource.com/23551
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 4969b46a316888950bd1910d7ef123883ab6c9f3
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sat May 28 02:40:32 2016 -0700

    doc/go1.7.html: update documentation regarding compress/flate
    
    Document the following:
    * That the algorithmic changes are still compliant with RFC 1951. I remember
    people having questions regarding this issue, and it would be good to re-assure
    them that it is still standards compliant.
    * io.EOF can now be returned early (c27efce66bce7534dbb357ac1779bbc08395b267)
    * Use the term "decompress" when referred to as an action. The term "uncompressed"
    or "decompressed" are both valid as ways to represent the current state of the data.
    
    Change-Id: Ie29ebce709357359e7c36d3e7f3d53b260eaadfa
    Reviewed-on: https://go-review.googlesource.com/23552
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 24996832c68b9d0aa4cd0e51189d148aae7a2772
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sat May 28 02:33:23 2016 -0600

    net/http/httputil: fix typos in deprecation comments
    
    Fixes #15868
    
    Change-Id: I4e4471e77091309c4ea1d546b2c4f20dfbb4314e
    Reviewed-on: https://go-review.googlesource.com/23550
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 795809b5c7d7e281e392399b9a366cbe92aa9e98
Author: Robert Griesemer <gri@golang.org>
Date:   Mon May 23 17:43:09 2016 -0700

    go/types: better debugging output for init order computation
    
    Also: Added some test cases for issue #10709.
    No impact when debugging output is disabled (default).
    
    For #10709.
    
    Change-Id: I0751befb222c86d46225377a674f6bad2990349e
    Reviewed-on: https://go-review.googlesource.com/23442
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 429bbf331247ef598802a94a23670bfe1cf61d6f
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Wed May 25 16:33:19 2016 +0300

    strings: fix and reenable amd64 Index for 17-31 byte strings
    
    Fixes #15689
    
    Change-Id: I56d0103738cc35cd5bc5e77a0e0341c0dd55530e
    Reviewed-on: https://go-review.googlesource.com/23440
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Nigel Tao <nigeltao@golang.org>

commit 42da35c699853002b7695052a8eeb3f10019cfd5
Author: Keith Randall <khr@golang.org>
Date:   Fri May 27 14:07:37 2016 -0700

    cmd/compile: SSA, don't let write barrier clobber return values
    
    When we do *p = f(), we might need to copy the return value from
    f to p with a write barrier.  The write barrier itself is a call,
    so we need to copy the return value of f to a temporary location
    before we call the write barrier function.  Otherwise, the call
    itself (specifically, marshalling the args to typedmemmove) will
    clobber the value we're trying to write.
    
    Fixes #15854
    
    Change-Id: I5703da87634d91a9884e3ec098d7b3af713462e7
    Reviewed-on: https://go-review.googlesource.com/23522
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3a6a41868eb620912235f2dd3f9738c76035e731
Author: Russ Cox <rsc@golang.org>
Date:   Fri May 27 16:30:03 2016 -0400

    doc: mention frame pointers in Go 1.7 release notes
    
    For #15840.
    
    Change-Id: I2ecf5c7b00afc2034cf3d7a1fd78636a908beb67
    Reviewed-on: https://go-review.googlesource.com/23517
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 496cf215cf7f36c2b1b14c00aadccf3e16f67eae
Author: Austin Clements <austin@google.com>
Date:   Fri May 27 14:25:16 2016 -0400

    crypto/tls: gofmt
    
    Commit fa3543e introduced formatting errors.
    
    Change-Id: I4b921f391a9b463cefca4318ad63b70ae6ce6865
    Reviewed-on: https://go-review.googlesource.com/23514
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: David Chase <drchase@google.com>

commit e149624ebb00a2fcc59bc02b9f122e3c4bae6e9c
Author: Austin Clements <austin@google.com>
Date:   Fri May 27 14:24:26 2016 -0400

    cmd/compile/internal/gc: gofmt
    
    Commit 36a80c5 introduced formatting errors.
    
    Change-Id: I6d5b231200cd7abcd5b94c1a3f4e99f10ee11c4f
    Reviewed-on: https://go-review.googlesource.com/23513
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: David Chase <drchase@google.com>

commit 53af0d3476e3c5f5b71f0c5fcf2141c24cc102b2
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat May 28 03:20:11 2016 +0900

    crypto/tls: fix race in Benchmark{Throughput,Latency}
    
    Fixes #15864.
    
    Change-Id: Ic12aa3654bf0b7e4a26df20ea92d07d7efe7339c
    Reviewed-on: https://go-review.googlesource.com/23504
    Reviewed-by: David Chase <drchase@google.com>

commit b0b2f7d6dda2b01a06a1dd99b87c97c81934c184
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri May 27 17:35:45 2016 +0900

    net/http/httptrace: fix nit in test
    
    Change-Id: I6dc3666398b4cd7a7195bb9c0e321fa8b733fa15
    Reviewed-on: https://go-review.googlesource.com/23502
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c340f4867b61d0c9dab167df88f56efc4ed7f17b
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri May 27 17:34:22 2016 +0900

    runtime: skip TestGdbBacktrace on netbsd
    
    Also adds missing copyright notice.
    
    Updates #15603.
    
    Change-Id: Icf4bb45ba5edec891491fe5f0039a8a25125d168
    Reviewed-on: https://go-review.googlesource.com/23501
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6a86dbe75f6d5a135eefbae807d98e856136514f
Author: Austin Clements <austin@google.com>
Date:   Fri May 27 12:21:14 2016 -0400

    runtime: always call stackfree on the system stack
    
    Currently when the garbage collector frees stacks of dead goroutines
    in markrootFreeGStacks, it calls stackfree on a regular user stack.
    This is a problem, since stackfree manipulates the stack cache in the
    per-P mcache, so if it grows the stack or gets preempted in the middle
    of manipulating the stack cache (which are both possible since it's on
    a user stack), it can easily corrupt the stack cache.
    
    Fix this by calling markrootFreeGStacks on the system stack, so that
    all calls to stackfree happen on the system stack. To prevent this bug
    in the future, mark stack functions that manipulate the mcache as
    go:systemstack.
    
    Fixes #15853.
    
    Change-Id: Ic0d1c181efb342f134285a152560c3a074f14a3d
    Reviewed-on: https://go-review.googlesource.com/23511
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 966baedfea6b09fa203b0cf0e6388830cc9f9fa7
Author: Austin Clements <austin@google.com>
Date:   Thu May 26 11:05:01 2016 -0400

    runtime: record Python stack on TestGdbPython failure
    
    For #15599.
    
    Change-Id: Icc2e58a3f314b7a098d78fe164ba36f5b2897de6
    Reviewed-on: https://go-review.googlesource.com/23481
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit fa3543e33782fd90e0a8f36366d9889d39a7575e
Author: Russ Cox <rsc@golang.org>
Date:   Fri May 27 09:50:06 2016 -0400

    crypto/tls: adjust dynamic record sizes to grow arithmetically
    
    The current code, introduced after Go 1.6 to improve latency on
    low-bandwidth connections, sends 1 kB packets until 1 MB has been sent,
    and then sends 16 kB packets (the maximum record size).
    
    Unfortunately this decreases throughput for 1-16 MB responses by 20% or so.
    
    Following discussion on #15713, change cutoff to 128 kB sent
    and also grow the size allowed for successive packets:
    1 kB, 2 kB, 3 kB, ..., 15 kB, 16 kB.
    This fixes the throughput problems: the overhead is now closer to 2%.
    
    I hope this still helps with latency but I don't have a great way to test it.
    At the least, it's not worse than Go 1.6.
    
    Comparing MaxPacket vs DynamicPacket benchmarks:
    
    name              maxpkt time/op  dyn. time/op delta
    Throughput/1MB-8    5.07ms ± 7%   5.21ms ± 7%  +2.73%  (p=0.023 n=16+16)
    Throughput/2MB-8   15.7ms ±201%    8.4ms ± 5%    ~     (p=0.604 n=20+16)
    Throughput/4MB-8    14.3ms ± 1%   14.5ms ± 1%  +1.53%  (p=0.000 n=16+16)
    Throughput/8MB-8    26.6ms ± 1%   26.8ms ± 1%  +0.47%  (p=0.003 n=19+18)
    Throughput/16MB-8   51.0ms ± 1%   51.3ms ± 1%  +0.47%  (p=0.000 n=20+20)
    Throughput/32MB-8    100ms ± 1%    100ms ± 1%  +0.24%  (p=0.033 n=20+20)
    Throughput/64MB-8    197ms ± 0%    198ms ± 0%  +0.56%   (p=0.000 n=18+7)
    
    The small MB runs are bimodal in both cases, probably GC pauses.
    But there's clearly no general slowdown anymore.
    
    Fixes #15713.
    
    Change-Id: I5fc44680ba71812d24baac142bceee0e23f2e382
    Reviewed-on: https://go-review.googlesource.com/23487
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 52c2db7e6d68b68938d904864ee484e7b5dd5d52
Author: Russ Cox <rsc@golang.org>
Date:   Fri May 27 12:15:04 2016 -0400

    doc/go1.7.html: fix broken sentence
    
    Change-Id: Ia540c890767dcb001d3b3b55d98d9517b13b21da
    Reviewed-on: https://go-review.googlesource.com/23510
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 605e751b53a8bec7d1f51b2ccc0093e063358dc6
Author: Russ Cox <rsc@golang.org>
Date:   Fri May 27 11:05:14 2016 -0400

    net/http: change Transport.Dialer to Transport.DialContext
    
    New in Go 1.7 so still possible to change.
    This allows implementations not tied to *net.Dialer.
    
    Fixes #15748.
    
    Change-Id: I5fabbf13c7f1951c06587a4ccd120def488267ce
    Reviewed-on: https://go-review.googlesource.com/23489
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 36a80c5941ec36d9c44d6f3c068d13201e023b5f
Author: Russ Cox <rsc@golang.org>
Date:   Fri May 27 00:56:19 2016 -0400

    cmd/compile: clean up, document Node closure fields
    
    Requested during CL 23431.
    
    Change-Id: I513ae42166b3a9fcfe51231ff55c163ab672e7d2
    Reviewed-on: https://go-review.googlesource.com/23485
    Reviewed-by: David Chase <drchase@google.com>

commit 93369001c76e01b2fe8b0d8a5074d62d0b8fdc81
Author: Russ Cox <rsc@golang.org>
Date:   Thu May 26 23:43:19 2016 -0400

    cmd/compile: delete Func.Outer
    
    This was just storage for a linked list.
    
    Change-Id: I850e8db1e1f5e72410f5c904be9409179b56a94a
    Reviewed-on: https://go-review.googlesource.com/23484
    Reviewed-by: David Chase <drchase@google.com>

commit cedc7c8f2091129514276c3c2f5f046f523a4684
Author: Russ Cox <rsc@golang.org>
Date:   Fri May 27 10:58:00 2016 -0400

    doc/go1.7.html: incorporate Rob's comments from CL 23379
    
    For #15810.
    
    Change-Id: Ib529808f664392feb9b36770f3d3d875fcb54528
    Reviewed-on: https://go-review.googlesource.com/23488
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 65dd08197ece4f64b990aa0023286c8f6abc25fa
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Fri May 27 02:31:59 2016 -0600

    doc/go1.7: document signal name printing during panics
    
    Document new behavior about signal name printing
    in panics as per CL golang.org/cl/22753.
    
    For #15810
    
    Change-Id: I9c677d5dd779b41e82afa25e3c797d8e739600d3
    Reviewed-on: https://go-review.googlesource.com/23493
    Reviewed-by: Russ Cox <rsc@golang.org>

commit dec1bae916fc75a6718fb7fa667e419cc902097a
Author: Russ Cox <rsc@golang.org>
Date:   Wed May 25 10:01:58 2016 -0400

    cmd/compile: additional paranoia and checking in plive.go
    
    The main check here is that liveness now crashes if it finds an instruction
    using a variable that should be tracked but is not.
    
    Comments and adjustments in nodarg to explain what's going on and
    to remove the "-1" argument added a few months ago, plus a sketch
    of a future simplification.
    
    The need for n.Orig in the earlier CL seems to have been an intermediate
    problem rather than fundamental: the new explanations in nodarg make
    clear that nodarg is not causing the problem I thought, and in fact now
    using n instead of n.Orig works fine in plive.go.
    
    Change-Id: I3f5cf9f6e4438a6d27abac7d490e7521545cd552
    Reviewed-on: https://go-review.googlesource.com/23450
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit e9228dd9490c3d4827170abeb8c82e68c175ecf0
Author: Dmitri Shuralyov <shurcooL@gmail.com>
Date:   Wed May 25 23:22:11 2016 -0700

    cmd/go: fixup for parsing SCP-like addresses
    
    This is a fixup change for commit 5cd294480364eb166751838a3df8f58649c214e1
    that added parsing of SCP-like addresses. To get the expected output
    from (*url.URL).String(), Path needs to be set, not RawPath.
    
    Add a test for this, since it has already regressed multiple times.
    
    Updates #11457.
    
    Change-Id: I806f5abbd3cf65e5bdcef01aab872caa8a5b8891
    Reviewed-on: https://go-review.googlesource.com/23447
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 20803b845f26111eb9281f8ece98383d48ea1b3f
Author: Russ Cox <rsc@golang.org>
Date:   Wed May 25 10:29:50 2016 -0400

    cmd/compile: eliminate PPARAMREF
    
    As in the elimination of PHEAP|PPARAM in CL 23393,
    this is something the front end can trivially take care of
    and then not bother the back ends with.
    It also eliminates some suspect (and only lightly exercised)
    code paths in the back ends.
    
    I don't have a smoking gun for this one but it seems
    more clearly correct.
    
    Change-Id: I3b3f5e669b3b81d091ff1e2fb13226a6f14c69d5
    Reviewed-on: https://go-review.googlesource.com/23431
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit b6dc3e6f668e4e0a2d2f710a7604d163d8ca45e1
Author: Russ Cox <rsc@golang.org>
Date:   Wed May 25 01:33:24 2016 -0400

    cmd/compile: fix liveness computation for heap-escaped parameters
    
    The liveness computation of parameters generally was never
    correct, but forcing all parameters to be live throughout the
    function covered up that problem. The new SSA back end is
    too clever: even though it currently keeps the parameter values live
    throughout the function, it may find optimizations that mean
    the current values are not written back to the original parameter
    stack slots immediately or ever (for example if a parameter is set
    to nil, SSA constant propagation may replace all later uses of the
    parameter with a constant nil, eliminating the need to write the nil
    value back to the stack slot), so the liveness code must now
    track the actual operations on the stack slots, exposing these
    problems.
    
    One small problem in the handling of arguments is that nodarg
    can return ONAME PPARAM nodes with adjusted offsets, so that
    there are actually multiple *Node pointers for the same parameter
    in the instruction stream. This might be possible to correct, but
    not in this CL. For now, we fix this by using n.Orig instead of n
    when considering PPARAM and PPARAMOUT nodes.
    
    The major problem in the handling of arguments is general
    confusion in the liveness code about the meaning of PPARAM|PHEAP
    and PPARAMOUT|PHEAP nodes, especially as contrasted with PAUTO|PHEAP.
    The difference between these two is that when a local variable "moves"
    to the heap, it's really just allocated there to start with; in contrast,
    when an argument moves to the heap, the actual data has to be copied
    there from the stack at the beginning of the function, and when a
    result "moves" to the heap the value in the heap has to be copied
    back to the stack when the function returns
    This general confusion is also present in the SSA back end.
    
    The PHEAP bit worked decently when I first introduced it 7 years ago (!)
    in 391425ae. The back end did nothing sophisticated, and in particular
    there was no analysis at all: no escape analysis, no liveness analysis,
    and certainly no SSA back end. But the complications caused in the
    various downstream consumers suggest that this should be a detail
    kept mainly in the front end.
    
    This CL therefore eliminates both the PHEAP bit and even the idea of
    "heap variables" from the back ends.
    
    First, it replaces the PPARAM|PHEAP, PPARAMOUT|PHEAP, and PAUTO|PHEAP
    variable classes with the single PAUTOHEAP, a pseudo-class indicating
    a variable maintained on the heap and available by indirecting a
    local variable kept on the stack (a plain PAUTO).
    
    Second, walkexpr replaces all references to PAUTOHEAP variables
    with indirections of the corresponding PAUTO variable.
    The back ends and the liveness code now just see plain indirected
    variables. This may actually produce better code, but the real goal
    here is to eliminate these little-used and somewhat suspect code
    paths in the back end analyses.
    
    The OPARAM node type goes away too.
    
    A followup CL will do the same to PPARAMREF. I'm not sure that
    the back ends (SSA in particular) are handling those right either,
    and with the framework established in this CL that change is trivial
    and the result clearly more correct.
    
    Fixes #15747.
    
    Change-Id: I2770b1ce3cbc93981bfc7166be66a9da12013d74
    Reviewed-on: https://go-review.googlesource.com/23393
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 99d29d5a43da0efde2ed9a137627d0d310e3baad
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu May 19 15:58:40 2016 +1000

    path/filepath: fix globbing of c:\*dir\... pattern
    
    The problem was introduced by the recent filepath.Join change.
    
    Fixes #14949
    
    Change-Id: I7ee52f210e12bbb1369e308e584ddb2c7766e095
    Reviewed-on: https://go-review.googlesource.com/23240
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit b5d18b50ac591d41cb4aab522fa9044c61b2c1e7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed May 18 18:42:25 2016 -0700

    cmd/cgo: remove -O options when generating compiler errors
    
    The cgo tool generates compiler errors to find out what kind of name it
    is using.  Turning on optimization can confuse that process by producing
    new unexpected messages.
    
    Fixes #14669.
    
    Change-Id: Idc8e35fd259711ecc9638566b691c11d17140325
    Reviewed-on: https://go-review.googlesource.com/23231
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 9369f22b8444d4e4afd791273ae13121d2ec7a6d
Author: Keith Randall <khr@golang.org>
Date:   Tue May 24 14:09:02 2016 -0700

    cmd/compile: testing harness for checking generated assembly
    
    Add a test which compiles a function and checks the
    generated assembly to make sure certain patterns are present.
    This test allows us to do white box tests of the compiler
    to make sure optimizations don't regress.
    
    Added a few simple tests for now.  More to come.
    
    Change-Id: I4ab5ce5d95b9e04e7d0d9328ffae47b8d1f95e74
    Reviewed-on: https://go-review.googlesource.com/23403
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b1894bb5cc781f1c59af727cea07ba4e84181830
Author: Quentin Smith <quentin@golang.org>
Date:   Thu May 26 17:53:21 2016 -0400

    encoding/json: improve Decode example
    
    Decoding a JSON message does not touch unspecified or null fields;
    always use a new underlying struct to prevent old field values from
    sticking around.
    
    Fixes: #14640
    
    Change-Id: Ica78c208ce104e2cdee1d4e92bf58596ea5587c8
    Reviewed-on: https://go-review.googlesource.com/23483
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 02bf6b5ef93cdb981949e8164e4364a9d8424dc9
Author: Keith Randall <khr@golang.org>
Date:   Thu May 26 14:31:35 2016 -0700

    cmd/compile: add tests for logical simplification rewrite rules
    
    Increases coverage of generic.rules from 72% to 84%.
    
    Change-Id: I1b139aeeb6410d025d49cbe4e4601f6f935ce1e5
    Reviewed-on: https://go-review.googlesource.com/23490
    Reviewed-by: Todd Neal <todd@tneal.org>

commit 0e930015c16b73dc9f98776b6624f02ea41d8268
Author: Keith Randall <khr@golang.org>
Date:   Tue May 24 15:43:25 2016 -0700

    cmd/compile: log rules to a file for rule coverage tool
    
    When rules are generated with -log, log rule application to a file.
    
    The file is opened in append mode so multiple calls to the compiler
    union their logs.
    
    Change-Id: Ib35c7c85bf58e5909ea9231043f8cbaa6bf278b7
    Reviewed-on: https://go-review.googlesource.com/23406
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 13a5b1faee06b59df456930d04edd2b5e083b019
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu May 26 12:16:53 2016 -0700

    cmd/compile: improve domorder documentation
    
    domorder has some non-obvious useful properties
    that we’re relying on in cse.
    Document them and provide an argument that they hold.
    While we’re here, do some minor renaming.
    
    The argument is a re-working of a private email
    exchange with Todd Neal and David Chase.
    
    Change-Id: Ie154e0521bde642f5f11e67fc542c5eb938258be
    Reviewed-on: https://go-review.googlesource.com/23449
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 2deb9209dec81792156c8e865a409a4ee5c331f6
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Thu May 26 11:29:53 2016 +0200

    math/big: using Run for some more benchmarks
    
    Change-Id: I3ede8098f405de5d88e51c8370d3b68446d40744
    Reviewed-on: https://go-review.googlesource.com/23428
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit bfccf4071eb9401cb906b203659010f9422c524f
Author: Russ Cox <rsc@golang.org>
Date:   Thu May 26 15:02:55 2016 -0400

    cmd/dist: drop ppc64le from testcshared
    
    I'm glad my CL fixed the library use case inside Google.
    It fixes neither of the two tests here.
    
    Change-Id: Ica91722dced8955a0a8ba3aad3d288816b46564e
    Reviewed-on: https://go-review.googlesource.com/23482
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 7fdec6216c0a25c6dbcc8159b755da6682dd9080
Author: Russ Cox <rsc@golang.org>
Date:   Wed May 25 14:37:43 2016 -0400

    build: enable framepointer mode by default
    
    This has a minor performance cost, but far less than is being gained by SSA.
    As an experiment, enable it during the Go 1.7 beta.
    Having frame pointers on by default makes Linux's perf, Intel VTune,
    and other profilers much more useful, because it lets them gather a
    stack trace efficiently on profiling events.
    (It doesn't help us that much, since when we walk the stack we usually
    need to look up PC-specific information as well.)
    
    Fixes #15840.
    
    Change-Id: I4efd38412a0de4a9c87b1b6e5d11c301e63f1a2a
    Reviewed-on: https://go-review.googlesource.com/23451
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2168f2a68bb438996d14869ff7dd10a47cc0552c
Author: Robert Griesemer <gri@golang.org>
Date:   Wed May 25 17:29:56 2016 -0700

    math/big: simplify benchmarking code some more
    
    Follow-up cleanup to https://golang.org/cl/23424/ .
    
    Change-Id: Ifb05c1ff5327df6bc5f4cbc554e18363293f7960
    Reviewed-on: https://go-review.googlesource.com/23446
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>
```
