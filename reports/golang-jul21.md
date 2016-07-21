# July 21, 2016 Report

Number of commits: 15

## Compilation time

* github.com/coreos/etcd/cmd: from 14.920296003s to 15.143449284s, +1.50%
* github.com/boltdb/bolt/cmd/bolt: from 547.781711ms to 555.377808ms, +1.39%
* github.com/gogits/gogs: from 13.280493742s to 13.014074849s, -2.01%
* github.com/spf13/hugo: from 6.883819829s to 6.680270726s, -2.96%
* github.com/nsqio/nsq/apps/nsqd: from 2.111771695s to 2.177042188s, +3.09%
* github.com/mholt/caddy: from 297.642053ms to 293.165304ms, -1.50%

## Binary size:

* github.com/coreos/etcd/cmd: from 26572558 to 26572649, ~
* github.com/boltdb/bolt/cmd/bolt: from 2675184 to 2679371, +0.16%
* github.com/gogits/gogs: from 23697107 to 23701294, ~
* github.com/spf13/hugo: from 15211322 to 15211413, ~
* github.com/nsqio/nsq/apps/nsqd: from 10051233 to 10051324, ~
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               189           192           +1.59%
BenchmarkMsgpUnmarshal-4             408           407           -0.25%
BenchmarkEasyJsonMarshal-4           1604          1614          +0.62%
BenchmarkEasyJsonUnmarshal-4         1525          1521          -0.26%
BenchmarkFlatBuffersMarshal-4        370           365           -1.35%
BenchmarkFlatBuffersUnmarshal-4      293           292           -0.34%
BenchmarkGogoprotobufMarshal-4       163           164           +0.61%
BenchmarkGogoprotobufUnmarshal-4     252           250           -0.79%

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
commit ff227b8a56b66e72de744a39f5b68d6e6ce7f3fe
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jul 20 15:40:10 2016 -0700

    runtime: add explicit `INT $3` at end of Darwin amd64 sigtramp
    
    The omission of this instruction could confuse the traceback code if a
    SIGPROF occurred during a signal handler.  The traceback code would
    trace up to sigtramp, but would then get confused because it would see a
    PC address that did not appear to be in the function.
    
    Fixes #16453.
    
    Change-Id: I2b3d53e0b272fb01d9c2cb8add22bad879d3eebc
    Reviewed-on: https://go-review.googlesource.com/25104
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit f407ca9288c8556c466e316f390ee7e7e99647ae
Author: Austin Clements <austin@google.com>
Date:   Mon Jul 18 16:01:22 2016 -0400

    runtime: support smaller physical pages than PhysPageSize
    
    Most operations need an upper bound on the physical page size, which
    is what sys.PhysPageSize is for (this is checked at runtime init on
    Linux). However, a few operations need a *lower* bound on the physical
    page size. Introduce a "minPhysPageSize" constant to act as this lower
    bound and use it where it makes sense:
    
    1) In addrspace_free, we have to query each page in the given range.
       Currently we increment by the upper bound on the physical page
       size, which means we may skip over pages if the true size is
       smaller. Worse, we currently pass a result buffer that only has
       enough room for one page. If there are actually multiple pages in
       the range passed to mincore, the kernel will overflow this buffer.
       Fix these problems by incrementing by the lower-bound on the
       physical page size and by passing "1" for the length, which the
       kernel will round up to the true physical page size.
    
    2) In the write barrier, the bad pointer check tests for pointers to
       the first physical page, which are presumably small integers
       masquerading as pointers. However, if physical pages are smaller
       than we think, we may have legitimate pointers below
       sys.PhysPageSize. Hence, use minPhysPageSize for this test since
       pointers should never fall below that.
    
    In particular, this applies to ARM64 and MIPS. The runtime is
    configured to use 64kB pages on ARM64, but by default Linux uses 4kB
    pages. Similarly, the runtime assumes 16kB pages on MIPS, but both 4kB
    and 16kB kernel configurations are common. This also applies to ARM on
    systems where the runtime is recompiled to deal with a larger page
    size. It is also a step toward making the runtime use only a
    dynamically-queried page size.
    
    Change-Id: I1fdfd18f6e7cbca170cc100354b9faa22fde8a69
    Reviewed-on: https://go-review.googlesource.com/25020
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Austin Clements <austin@google.com>

commit d73ca5f4d8f6aef0c2e738cd1614d4dbf87735fb
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Wed Jul 20 13:52:03 2016 +0200

    runtime/race: fix memory leak
    
    The leak was reported internally on a sever canary that runs for days.
    After a day server consumes 5.6GB, after 6 days -- 12.2GB.
    The leak is exposed by the added benchmark.
    The leak is fixed upstream in :
    http://llvm.org/viewvc/llvm-project/compiler-rt/trunk/lib/tsan/rtl/tsan_rtl_thread.cc?view=diff&r1=276102&r2=276103&pathrev=276103
    
    Fixes #16441
    
    Change-Id: I9d4f0adef48ca6cf2cd781b9a6990ad4661ba49b
    Reviewed-on: https://go-review.googlesource.com/25091
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>

commit 50048a4e8ee11016227c283be2d073e14e1c006b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jul 18 23:00:43 2016 -0700

    runtime: add as many extra M's as needed
    
    When a non-Go thread calls into Go, the runtime needs an M to run the Go
    code. The runtime keeps a list of extra M's available. When the last
    extra M is allocated, the needextram field is set to tell it to allocate
    a new extra M as soon as it is running in Go. This ensures that an extra
    M will always be available for the next thread.
    
    However, if many threads need an extra M at the same time, this
    serializes them all. One thread will get an extra M with the needextram
    field set. All the other threads will see that there is no M available
    and will go to sleep. The one thread that succeeded will create a new
    extra M. One lucky thread will get it. All the other threads will see
    that there is no M available and will go to sleep. The effect is
    thundering herd, as all the threads looking for an extra M go through
    the process one by one. This seems to have a particularly bad effect on
    the FreeBSD scheduler for some reason.
    
    With this change, we track the number of threads waiting for an M, and
    create all of them as soon as one thread gets through. This still means
    that all the threads will fight for the lock to pick up the next M. But
    at least each thread that gets the lock will succeed, instead of going
    to sleep only to fight again.
    
    This smooths out the performance greatly on FreeBSD, reducing the
    average wall time of `testprogcgo CgoCallbackGC` by 74%.  On GNU/Linux
    the average wall time goes down by 9%.
    
    Fixes #13926
    Fixes #16396
    
    Change-Id: I6dc42a4156085a7ed4e5334c60b39db8f8ef8fea
    Reviewed-on: https://go-review.googlesource.com/25047
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 883e128f4571a59842e1156b5ebe25d8420162d9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Jul 19 20:27:55 2016 -0700

    net/smtp: document that the smtp package is frozen
    
    This copies the frozen wording from the log/syslog package.
    
    Fixes #16436
    
    Change-Id: If5d478023328925299399f228d8aaf7fb117c1b4
    Reviewed-on: https://go-review.googlesource.com/25080
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 1d2ca9e30c22bc9e8cd0b21dff58367443696c91
Author: Austin Clements <austin@google.com>
Date:   Mon Jul 18 11:34:11 2016 -0400

    doc/go1.7.html: start sentence on a new line
    
    Change-Id: Ia1c2ebcd2ccf7b98d89b378633bf4fc435d2364d
    Reviewed-on: https://go-review.googlesource.com/25019
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3ad586155bb8cd41fa0c0650a6b5feca871dfeed
Author: Austin Clements <austin@google.com>
Date:   Mon Jul 18 11:33:43 2016 -0400

    doc/go1.7.html: avoid term of art
    
    Rather than saying "stop-the-world", say "garbage collection pauses".
    
    Change-Id: Ifb2931781ab3094e04bea93f01f18f1acb889bdc
    Reviewed-on: https://go-review.googlesource.com/25018
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit d66cbec37afb7936b1ea0f7f2433cc070f667112
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jul 18 08:14:10 2016 -0700

    doc/go1.7.html: the 1.6.3 release supports Sierra
    
    Updates #16354
    Updates #16272
    
    Change-Id: I73e8df40621a0a17a1990f3b10ea996f4fa738aa
    Reviewed-on: https://go-review.googlesource.com/25014
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b3b0b7a1825c9249f2b323ffd23cbb128044fb6a
Author: Chris Broadfoot <cbro@golang.org>
Date:   Sun Jul 17 23:30:19 2016 -0700

    doc: document go1.6.3
    
    Change-Id: Ib33d7fb529aafcaf8ca7d43b2c9480f30d5c28cc
    Reviewed-on: https://go-review.googlesource.com/25011
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b97df54c31d6c4cc2a28a3c83725366d52329223
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Jul 18 06:05:24 2016 +0000

    net/http, net/http/cgi: fix for CGI + HTTP_PROXY security issue
    
    Because,
    
    * The CGI spec defines that incoming request header "Foo: Bar" maps to
      environment variable HTTP_FOO == "Bar". (see RFC 3875 4.1.18)
    
    * The HTTP_PROXY environment variable is conventionally used to configure
      the HTTP proxy for HTTP clients (and is respected by default for
      Go's net/http.Client and Transport)
    
    That means Go programs running in a CGI environment (as a child
    process under a CGI host) are vulnerable to an incoming request
    containing "Proxy: attacker.com:1234", setting HTTP_PROXY, and
    changing where Go by default proxies all outbound HTTP requests.
    
    This is CVE-2016-5386, aka https://httpoxy.org/
    
    Fixes #16405
    
    Change-Id: I6f68ade85421b4807785799f6d98a8b077e871f0
    Reviewed-on: https://go-review.googlesource.com/25010
    Run-TryBot: Chris Broadfoot <cbro@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit 2837c395526476e31fb15dbb948ed77389cdc75b
Author: Austin Clements <austin@google.com>
Date:   Sun Jul 17 23:12:41 2016 -0400

    doc/go1.7.html: mention specific runtime improvements
    
    Most of the runtime improvements are hard to quantify or summarize,
    but it's worth mentioning some of the substantial improvements in STW
    time, and that the scavenger now actually works on ARM64, PPC64, and
    MIPS.
    
    Change-Id: I0e951038516378cc3f95b364716ef1c183f3445a
    Reviewed-on: https://go-review.googlesource.com/24966
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a6dbfc12c640c90e8dc552443d3ece04cbae4a9c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Jul 16 23:56:45 2016 +0000

    net: demote TestDialerDualStack to a flaky test
    
    Only run TestDialerDualStack on the builders, as to not annoy or
    otherwise distract users when it's not their fault.
    
    Even though the intention is to only run this on the builders, very
    few of the builders have IPv6 support. Oh well. We'll get some
    coverage.
    
    Updates #13324
    
    Change-Id: I13e7e3bca77ac990d290cabec88984cc3d24fb67
    Reviewed-on: https://go-review.googlesource.com/24985
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit 510fb6397dfc93067dc90d42c58dfc5f8b995285
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sat Jul 16 02:42:52 2016 -0700

    fmt: properly handle early io.EOF Reads in readRune.readByte
    
    Change https://golang.org/cl/19895 caused a regression
    where the last character in a string would be dropped if it was
    accompanied by an io.EOF.
    
    This change fixes the logic so that the last byte is still returned
    without a problem.
    
    Fixes #16393
    
    Change-Id: I7a4d0abf761c2c15454136a79e065fe002d736ea
    Reviewed-on: https://go-review.googlesource.com/24981
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2b6eb276517ecba08985d59b6b1928e29743d3e0
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jul 14 08:50:44 2016 -0700

    doc/go1.7.html: remove erroneous note about ppc64 and power8
    
    We decided that ppc64 should maintain power5 compatibility.
    ppc64le requires power8.
    
    Fixes #16372.
    
    Change-Id: If5b309a0563f55a3c1fe9c853d29a463f5b71101
    Reviewed-on: https://go-review.googlesource.com/24915
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4054769a31f66039f5fd74ca3164e9233f724da8
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 14 07:25:05 2016 -0700

    runtime/internal/atomic: fix assembly arg sizes
    
    Change-Id: I80ccf40cd3930aff908ee64f6dcbe5f5255198d3
    Reviewed-on: https://go-review.googlesource.com/24914
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
```
