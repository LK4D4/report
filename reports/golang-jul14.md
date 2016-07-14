# July 14, 2016 Report

Number of commits: 20

## Compilation time

* github.com/coreos/etcd/cmd: from 15.320427747s to 15.558997012s, +1.56%
* github.com/boltdb/bolt/cmd/bolt: from 557.769126ms to 552.974298ms, -0.86%
* github.com/gogits/gogs: from 13.026002745s to 12.859951085s, -1.27%
* github.com/spf13/hugo: from 6.894096167s to 6.788778391s, -1.53%
* github.com/nsqio/nsq/apps/nsqd: from 2.33865218s to 2.11327643s, -9.64%
* github.com/mholt/caddy: from 278.858609ms to 276.067353ms, -1.00%

## Binary size:

* github.com/coreos/etcd/cmd: from 26485214 to 26493811, ~
* github.com/boltdb/bolt/cmd/bolt: from 2666587 to 2675184, +0.32%
* github.com/gogits/gogs: from 23704257 to 23716950, ~
* github.com/spf13/hugo: from 15202443 to 15206944, ~
* github.com/nsqio/nsq/apps/nsqd: from 10042636 to 10051233, ~
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               193           190           -1.55%
BenchmarkMsgpUnmarshal-4             407           408           +0.25%
BenchmarkEasyJsonMarshal-4           1608          1612          +0.25%
BenchmarkEasyJsonUnmarshal-4         1531          1521          -0.65%
BenchmarkFlatBuffersMarshal-4        359           368           +2.51%
BenchmarkFlatBuffersUnmarshal-4      291           294           +1.03%
BenchmarkGogoprotobufMarshal-4       163           164           +0.61%
BenchmarkGogoprotobufUnmarshal-4     256           257           +0.39%

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
commit 29ed5da5f2804cab0f6f1c97309673ac5d22a99d
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jul 13 13:22:47 2016 -0700

    runtime/pprof: don't print extraneous 0 after goexit
    
    This fixes erroneous handling of the more result parameter of
    runtime.Frames.Next.
    
    Fixes #16349.
    
    Change-Id: I4f1c0263dafbb883294b31dbb8922b9d3e650200
    Reviewed-on: https://go-review.googlesource.com/24911
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4d00937cecdea85b6f1eb894a6d28a53f5f2ff8a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Jul 13 10:49:48 2016 -0600

    all: rename vendored golang.org/x/net packages to golang_org
    
    Regression from Go 1.6 to Go 1.7rc1: we had broken the ability for
    users to vendor "golang.org/x/net/http2" or "golang.org/x/net/route"
    because we were vendoring them ourselves and cmd/go and cmd/compile do
    not understand multiple vendor directories across multiple GOPATH
    workspaces (e.g. user's $GOPATH and default $GOROOT).
    
    As a short-term fix, since fixing cmd/go and cmd/compile is too
    invasive at this point in the cycle, just rename "golang.org" to
    "golang_org" for the standard library's vendored copy.
    
    Fixes #16333
    
    Change-Id: I9bfaed91e9f7d4ca6bab07befe80d71d437a21af
    Reviewed-on: https://go-review.googlesource.com/24902
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 1cb3f7169ccff3ae2197784676404e8d0d3f5e32
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jul 13 09:38:41 2016 -0700

    doc/go1.7.html: earlier Go versions don't work on macOS Sierra
    
    Updates #16272.
    
    Change-Id: If5444b8de8678eeb9be10b62a929e2e101d1dd91
    Reviewed-on: https://go-review.googlesource.com/24900
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 76da6491e802410bf84e122b8694bf01a6cf57cd
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Mon Jul 4 23:57:05 2016 -0700

    doc/go1.7.html: document that http.Server now enforces request versions
    
    Document that the http.Server is now stricter about rejecting
    requests with invalid HTTP versions, and also that it rejects plaintext
    HTTP/2 requests, except for `PRI * HTTP/2.0` upgrade requests.
    The relevant CL is https://golang.org/cl/24505.
    
    Updates #15810.
    
    Change-Id: Ibbace23e001b5e2eee053bd341de50f9b6d3fde8
    Reviewed-on: https://go-review.googlesource.com/24731
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2fcb25e07b2549f607aa174ceab974f8732ea0f4
Author: Bryan C. Mills <bcmills@google.com>
Date:   Tue Jul 12 18:56:07 2016 -0400

    doc/effective_go: clarify advice on returning interfaces
    
    New Gophers sometimes misconstrue the advice in the "Generality" section
    as "export interfaces instead of implementations" and add needless
    interfaces to their code as a result.  Down the road, they end up
    needing to add methods and either break existing callers or have to
    resort to unpleasant hacks (e.g. using "magic method" type-switches).
    
    Weaken the first paragraph of this section to only advise leaving types
    unexported when they will never need additional methods.
    
    Change-Id: I32a1ae44012b5896faf167c02e192398a4dfc0b8
    Reviewed-on: https://go-review.googlesource.com/24892
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit a1110c39301b21471c27dad0e50cdbe499587fc8
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Jul 12 11:19:16 2016 -0600

    cmd/go: don't fail on invalid GOOS/GOARCH pair when using gccgo
    
    Fixes #12272
    
    Change-Id: I0306ce0ef4a87df2158df3b7d4d8d93a1cb6dabc
    Reviewed-on: https://go-review.googlesource.com/24864
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit b30814bbd6840e1574a27c87c37515af22caa5d9
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jul 11 16:15:03 2016 -0700

    runtime: add ctxt parameter to cgocallback called from Go
    
    The cgocallback function picked up a ctxt parameter in CL 22508.
    That CL updated the assembler implementation, but there are a few
    mentions in Go code that were not updated. This CL fixes that.
    
    Fixes #16326
    
    Change-Id: I5f68e23565c6a0b11057aff476d13990bff54a66
    Reviewed-on: https://go-review.googlesource.com/24848
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 1f4e68d92b33a668f2afa2ab5f8114c1a4bee682
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jul 11 22:34:30 2016 -0700

    reflect: an unnamed type has no PkgPath
    
    The reflect package was returning a non-empty PkgPath for an unnamed
    type with methods, such as a type whose methods have a pointer
    receiver.
    
    Fixes #16328.
    
    Change-Id: I733e93981ebb5c5c108ef9b03bf5494930b93cf3
    Reviewed-on: https://go-review.googlesource.com/24862
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit a84b18ac865257c50d8812e39d244b57809fc8c8
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jul 12 03:43:05 2016 +0000

    Revert "regexp: add the Fanout benchmark
    
    This is a copy of the "FANOUT" benchmark recently added to RE2 with the
    following comment:
    
        // This has quite a high degree of fanout.
        // NFA execution will be particularly slow.
    
    Most of the benchmarks on the regexp package have very little fanout and
    are designed for comparing the regexp package's NFA with backtracking
    engines found in other regular expression libraries. This benchmark
    exercises the performance of the NFA on expressions with high fanout.Reviewed-on: https://go-review.googlesource.com/24846
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    "
    
    This reverts commit fc803874d3a509ddd99a897da1c6a62dc4ce631e.
    
    Reason for revert: Breaks the -race build because the benchmark takes too long to run.
    
    Change-Id: I6ed4b466f74a4108d8bcd5b019b9abe971eb483e
    Reviewed-on: https://go-review.googlesource.com/24861
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit fc803874d3a509ddd99a897da1c6a62dc4ce631e
Author: Michael Matloob <matloob@golang.org>
Date:   Mon Jul 11 14:59:03 2016 -0600

    regexp: add the Fanout benchmark
    
    This is a copy of the "FANOUT" benchmark recently added to RE2 with the
    following comment:
    
        // This has quite a high degree of fanout.
        // NFA execution will be particularly slow.
    
    Most of the benchmarks on the regexp package have very little fanout and
    are designed for comparing the regexp package's NFA with backtracking
    engines found in other regular expression libraries. This benchmark
    exercises the performance of the NFA on expressions with high fanout.
    
    Change-Id: Ie9c8e3bbeffeb1fe9fb90474ddd19e53f2f57a52
    Reviewed-on: https://go-review.googlesource.com/24846
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 296b618dc8c8f59d7327b4d322f7ceef4032d94b
Author: Francesc Campoy <campoy@golang.org>
Date:   Mon Jul 11 12:31:52 2016 -0600

    gofmt: remove unneeded call to os.Exit
    
    PrintDefaults already calls os.Exit(2).
    
    Change-Id: I0d783a6476f42b6157853cdb34ba69618e3f3fcb
    Reviewed-on: https://go-review.googlesource.com/24844
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 38de5b71f274ff93581a302b2f3ec4b9937afa51
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun Jul 10 21:47:56 2016 -0700

    doc/go1.7.html: no concurrent calls of math/rand methods
    
    A follow-on to https://golang.org/cl/24852 that mentions the
    documentation clarifications.
    
    Updates #16308.
    
    Change-Id: Ic2a6e1d4938d74352f93a6649021fb610efbfcd0
    Reviewed-on: https://go-review.googlesource.com/24857
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit fb3cf5c686b09bab8c1bef5f7589aaef0e6d9712
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sat Jul 9 19:38:04 2016 -0700

    math/rand: fix raciness in Rand.Read
    
    There are no synchronization points protecting the readVal and readPos
    variables. This leads to a race when Read is called concurrently.
    Fix this by adding methods to lockedSource, which is the case where
    a race matters.
    
    Fixes #16308.
    
    Change-Id: Ic028909955700906b2d71e5c37c02da21b0f4ad9
    Reviewed-on: https://go-review.googlesource.com/24852
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 54ffdf364f77c62ffeb205debe26347ca5961373
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Jul 9 17:24:45 2016 -0700

    net/http: fix vet warning of leaked context in error paths
    
    Updates #16230
    
    Change-Id: Ie38f85419c41c00108f8843960280428a39789b5
    Reviewed-on: https://go-review.googlesource.com/24850
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 54b499e3f1d3ef37765c209919d30f0abf55a2e1
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jul 8 11:42:19 2016 -0700

    syscall: add another output for TestGroupCleanupUserNamespace
    
    Fixes #16303.
    
    Change-Id: I2832477ce0117a66da53ca1f198ebb6121953d56
    Reviewed-on: https://go-review.googlesource.com/24833
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 12f2b4ff0ea694fc31e5b25d61d36cf058a88f35
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jul 8 07:56:52 2016 -0700

    runtime: fix case in KeepAlive comment
    
    Fixes #16299.
    
    Change-Id: I76f541c7f11edb625df566f2f1035147b8bcd9dd
    Reviewed-on: https://go-review.googlesource.com/24830
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 915398f14fff28f7ba8592f134e22079de044745
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jul 8 07:50:07 2016 -0700

    doc/go1.7.html: fix name of IsExist
    
    For better or for worse, it's IsExist, not IsExists.
    
    Change-Id: I4503f961486edd459c0c81cf3f32047dff7703a4
    Reviewed-on: https://go-review.googlesource.com/24819
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fad2bbdc6a686a20174d2e73cf78f1659722bb39
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jul 7 16:41:29 2016 -0700

    runtime: fix nanotime for macOS Sierra
    
    In the beta version of the macOS Sierra (10.12) release, the
    gettimeofday system call changed on x86. Previously it always returned
    the time in the AX/DX registers. Now, if AX is returned as 0, it means
    that the system call has stored the values into the memory pointed to by
    the first argument, just as the libc gettimeofday function does. The
    libc function handles both cases, and we need to do so as well.
    
    Fixes #16272.
    
    Change-Id: Ibe5ad50a2c5b125e92b5a4e787db4b5179f6b723
    Reviewed-on: https://go-review.googlesource.com/24812
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 84bb9e62f06dbb62279241fa0bd7a6c8846271ac
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jul 7 17:43:08 2016 -0700

    runtime: handle selects with duplicate channels in shrinkstack
    
    The shrinkstack code locks all the channels a goroutine is waiting for,
    but didn't handle the case of the same channel appearing in the list
    multiple times. This led to a deadlock. The channels are sorted so it's
    easy to avoid locking the same channel twice.
    
    Fixes #16286.
    
    Change-Id: Ie514805d0532f61c942e85af5b7b8ac405e2ff65
    Reviewed-on: https://go-review.googlesource.com/24815
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit e5ff529679b3adbed06d509b0fc21a76b62e89e9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Jul 6 00:16:05 2016 +0000

    lib/time: update to IANA release 2016f (July 2016)
    
    Fixes #16273
    
    Change-Id: I443e1f254fd683c4ff61beadae89c1c45ff5d972
    Reviewed-on: https://go-review.googlesource.com/24744
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
```
