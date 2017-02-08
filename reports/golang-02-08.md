# February 8, 2017 Report

Number of commits: 203

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 570.774264ms to 575.575064ms, +0.84%
* github.com/coreos/etcd/cmd/etcd: from 12.325686748s to 12.272951822s, -0.43%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 39.172222644s to 37.36356983s, -4.62%
* github.com/influxdata/influxdb/cmd/influxd: from 7.845077183s to 6.681726416s, -14.83%
* github.com/junegunn/fzf/src/fzf: from 985.641184ms to 956.05655ms, -3.00%
* github.com/mholt/caddy/caddy: from 6.930668491s to 6.14109582s, -11.39%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.437564787s to 1.447549369s, +0.69%
* github.com/nsqio/nsq/apps/nsqd: from 2.003728493s to 1.979226248s, -1.22%
* github.com/prometheus/prometheus/cmd/prometheus: from 42.190078386s to 40.489451161s, -4.03%
* github.com/spf13/hugo: from 6.58692492s to 6.455602513s, -1.99%
* golang.org/x/tools/cmd/guru: from 2.606907902s to 2.521335475s, -3.28%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 3188113 to 3197173, +0.28%
* github.com/coreos/etcd/cmd/etcd: from 27000232 to 27079960, +0.30%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 34380888 to 34485168, +0.30%
* github.com/influxdata/influxdb/cmd/influxd: from 17412925 to 17436983, +0.14%
* github.com/junegunn/fzf/src/fzf: from 3207488 to 3211413, +0.12%
* github.com/mholt/caddy/caddy: from 15828229 to 15819478, ~
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4362729 to 4367217, +0.10%
* github.com/nsqio/nsq/apps/nsqd: from 10340775 to 10373154, +0.31%
* github.com/prometheus/prometheus/cmd/prometheus: from 62704582 to 62281481, -0.67%
* github.com/spf13/hugo: from 16758494 to 16801550, +0.26%
* golang.org/x/tools/cmd/guru: from 8234173 to 8226235, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               203           203           +0.00%
BenchmarkMsgpUnmarshal-4             390           398           +2.05%
BenchmarkEasyJsonMarshal-4           1567          1552          -0.96%
BenchmarkEasyJsonUnmarshal-4         1578          1671          +5.89%
BenchmarkFlatBuffersMarshal-4        408           391           -4.17%
BenchmarkFlatBuffersUnmarshal-4      312           336           +7.69%
BenchmarkGogoprotobufMarshal-4       167           167           +0.00%
BenchmarkGogoprotobufUnmarshal-4     257           272           +5.84%

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

* [testing: synchronize writes to the root's Writer](https://github.com/golang/go/commit/7f31971f594edbacbdba5407aaee042850fbd220)
* [context: lazily initialize cancelCtx done channel](https://github.com/golang/go/commit/986768de7fcf4def65cecd7eb0c34e2cdf92e78c)
* [cmd/compile: disable memory profiling when not in use](https://github.com/golang/go/commit/302474c61c15095406325773172bfb0a819ce3af)
* [cmd/compile: use len(s)<=cap(s) to remove more bounds checks](https://github.com/golang/go/commit/73f92f9b0405e98427bbb445f24cffb5d3c4d01b)
* [cmd/compile: convert constants to interfaces without allocating](https://github.com/golang/go/commit/c682d3239e5aa05a77ad21f2267efc4e2e60c05f)
* [time: record monotonic clock reading in time.Now, for more accurate comparisons](https://github.com/golang/go/commit/0e3355903d2ebcf5ee9e76096f51ac9a116a9dbb)
* [sort: optimize average calculation in binary search](https://github.com/golang/go/commit/fd37b8ccf2262bb3f0a608f7545f78a72e8d661f)
* [spec: introduce alias declarations and type definitions](https://github.com/golang/go/commit/56c9b51b937cca7d3db517add96bd9517bbffb80)
* [encoding/json: add Valid for checking validity of input bytes](https://github.com/golang/go/commit/3f7a35d91c7079269dec5cefef7599148f0279e0)

## GIT Log

```
commit e410d2a81ef26d7dcef0c712b584d2345b15148e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Feb 7 15:29:18 2017 -0800

    cmd/gofmt: clear pattern match map at the correct time
    
    We need to clear the pattern match map after the recursive rewrite
    applications, otherwise there might be lingering entries that cause
    match to fail.
    
    Fixes #18987.
    
    Change-Id: I7913951c455c98932bda790861db6a860ebad032
    Reviewed-on: https://go-review.googlesource.com/36546
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 87ad863f359de3760578acb7f7a4d7e333c9cdc8
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Feb 3 21:13:30 2017 -0800

    runtime: use atomic ops for fwdSig, make sigtable immutable
    
    The fwdSig array is accessed by the signal handler, which may run in
    parallel with other threads manipulating it via the os/signal package.
    Use atomic accesses to ensure that there are no problems.
    
    Move the _SigHandling flag out of the sigtable array. This makes sigtable
    immutable and safe to read from the signal handler.
    
    Change-Id: Icfa407518c4ebe1da38580920ced764898dfc9ad
    Reviewed-on: https://go-review.googlesource.com/36321
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 14c2849c3ebe498971413ee5e8c9780fabc8578e
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Feb 7 17:35:51 2017 -0500

    runtime: update android time_now call
    
    This was broken in https://golang.org/cl/36255
    
    Change-Id: Ib23323a745a650ac51b0ead161076f97efe6d7b7
    Reviewed-on: https://go-review.googlesource.com/36543
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 48d71990722f3cc81305c794d3d3b6d9007770c8
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sat Jan 28 11:29:53 2017 +0100

    cmd/go: clarify that tag lists are space-separated
    
    Apparently the current documentation is confusing users that
    quickly skim the flags list at the top. Make very clear that
    build tags are space-separated.
    
    Updates #18800
    
    Change-Id: I473552c5a2b70ca03d8bbbd2c76805f7f82b49a2
    Reviewed-on: https://go-review.googlesource.com/35951
    Reviewed-by: Daniel Martí <mvdan@mvdan.cc>
    Reviewed-by: Minux Ma <minux@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3b84a3c9acaaba04a232f7e73a40c36bccd5e988
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Dec 7 10:49:45 2016 +1100

    os: make Stdin.Stat() return ModeCharDevice if Stdin is console
    
    CL 20845 changed Stdin.Stat() so it returns ModeNamedPipe.
    But introduced TestStatStdin does not test what Stdin.Stat()
    returns when Stdin is console.
    
    This CL adjusts both TestStatStdin and Stdin.Stat
    implementations to handle console. Return ModeCharDevice
    from Stdin.Stat() when Stdin is console on windows,
    just like it does on unix.
    
    Fixes #14853.
    
    Change-Id: I54d73caee2aea45a99618d11600d8e82fe20d0c0
    Reviewed-on: https://go-review.googlesource.com/34090
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3f7a35d91c7079269dec5cefef7599148f0279e0
Author: Matt Layher <mdlayher@gmail.com>
Date:   Tue Dec 13 17:57:06 2016 -0500

    encoding/json: add Valid for checking validity of input bytes
    
    Fixes #18086
    
    Change-Id: Idc501dd37893e04a01c6ed9920147d24c0c1fa18
    Reviewed-on: https://go-review.googlesource.com/34202
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1f93ba66d6ccb200e8b0a037a01265f6563fdf58
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 7 11:34:20 2017 -0800

    math/big: add IsInt64/IsUint64 predicates
    
    Change-Id: Ia5ed3919cb492009ac8f66d175b47a69f83ee4f1
    Reviewed-on: https://go-review.googlesource.com/36487
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 7bad00366b2e3e8440e8c870d8c53efaa8fe3811
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Feb 6 17:06:02 2017 -0800

    cmd/internal/obj: remove ATYPE
    
    In cmd/compile, we can directly construct obj.Auto to represent local
    variables and attach them to the function's obj.LSym.
    
    In preparation for being able to emit more precise DWARF info based on
    other compiler available information (e.g., lexical scoping).
    
    Change-Id: I9c4225ec59306bec42552838493022e0e9d70228
    Reviewed-on: https://go-review.googlesource.com/36420
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Heschi Kreinick <heschi@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 38cb9d28a9a3970f5bfeacdbeaa2f95aab3ebe73
Author: Sameer Ajmani <sameer@golang.org>
Date:   Tue Feb 7 13:12:25 2017 -0500

    runtime/pprof: document that profile names should not contain spaces.
    
    Change-Id: I967d897e812bee63b32bc2a7dcf453861b89b7e3
    Reviewed-on: https://go-review.googlesource.com/36533
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a833485828f02b098c67127ed8be8e78493aaf4b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Jan 16 12:37:11 2017 -0500

    cmd/compile: do not use statictmp for zeroing
    
    Also fixes #18687.
    
    Change-Id: I7c6d47c71e632adf4c16937a29074621f771844c
    Reviewed-on: https://go-review.googlesource.com/35261
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8cf1766930269ea09a4234cc42570abcf425355f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Feb 6 18:18:49 2017 -0800

    cmd/compile/internal/ssa: use *obj.LSym in ExternSymbol
    
    Change-Id: I713120f90fd1d2df6698c40622ccac6eae907919
    Reviewed-on: https://go-review.googlesource.com/36423
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 1a7582f5e980ecc8f23631336e8010db4b754c83
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Feb 6 15:52:36 2017 -0800

    cmd/internal/dwarf: use []*Var instead of linked lists
    
    Passes toolstash -cmp.
    
    Change-Id: I202b29495ca1aaf3c52879fa99fdc0a4b86703af
    Reviewed-on: https://go-review.googlesource.com/36419
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6cf7918e7335041bd7d40be4c6e976a044cc000c
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Tue Feb 7 11:01:51 2017 -0800

    runtime/pprof: clarify CPU profile's captured during the lifetime of the prog
    
    Fixes #18504.
    
    Change-Id: I3716fc58fc98472eea15ce3617aee3890670c276
    Reviewed-on: https://go-review.googlesource.com/36430
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 67c3d4dab06530b4a444c3a076fceaa5573cf2b6
Author: Sameer Ajmani <sameer@golang.org>
Date:   Tue Feb 7 12:27:29 2017 -0500

    time: delete incorrect docs about day-of-month checks.
    
    Documentation was introduced by CL https://golang.org/cl/14123
    but that behavior was changed later by CL https://golang.org/cl/17710.
    This CL deletes the stale paragraph.
    
    Fixes #18980
    
    Change-Id: Ib434f1eac6fc814fde1be112a8f52afe6e3e0fcc
    Reviewed-on: https://go-review.googlesource.com/36532
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 57d06fff3e7e020510fe9460ccfa247370c472ba
Author: Russ Cox <rsc@golang.org>
Date:   Tue Feb 7 10:29:32 2017 -0500

    cmd/go, go/build: better defenses against GOPATH=GOROOT
    
    Fixes #18863.
    
    Change-Id: I0723563cd23728b0d43ebcc25979bf8d21e2a72c
    Reviewed-on: https://go-review.googlesource.com/36427
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4af6b81d41c648f31ed8113151ed8f7fee6180c8
Author: Austin Clements <austin@google.com>
Date:   Fri Jan 13 13:23:41 2017 -0500

    runtime: fix confusion between _MaxMem and _MaxArena32
    
    Currently both _MaxMem and _MaxArena32 represent the maximum arena
    size on 32-bit hosts (except on MIPS32 where _MaxMem is confusingly
    smaller than _MaxArena32).
    
    Clean up sysAlloc so that it always uses _MaxMem, which is the maximum
    arena size on both 32- and 64-bit architectures and is the arena size
    we allocate auxiliary structures for. This lets us simplify and unify
    some code paths and eliminate _MaxArena32.
    
    Fixes #18651. mheap.sysAlloc currently assumes that if the arena is
    small, we must be on a 32-bit machine and can therefore grow the arena
    to _MaxArena32. This breaks down on darwin/arm64, where _MaxMem is
    only 2 GB. As a result, on darwin/arm64, we only reserve spans and
    bitmap space for a 2 GB heap, and if the application tries to allocate
    beyond that, sysAlloc takes the 32-bit path, tries to grow the arena
    beyond 2 GB, and panics when it tries to grow the spans array
    allocation past its reserved size. This has probably been a problem
    for several releases now, but was only noticed recently because
    mapSpans didn't check the bounds on the span reservation until
    recently. Most likely it corrupted the bitmap before. By using _MaxMem
    consistently, we avoid thinking that we can grow the arena larger than
    we have auxiliary structures for.
    
    Change-Id: Ifef28cb746a3ead4b31c1d7348495c2242fef520
    Reviewed-on: https://go-review.googlesource.com/35253
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Elias Naur <elias.naur@gmail.com>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1cc24690b80b5f78e07b6e14b8614317462da1ea
Author: Austin Clements <austin@google.com>
Date:   Fri Jan 13 14:19:52 2017 -0500

    runtime: simplify and cleanup mallocinit
    
    mallocinit has evolved organically. Make a pass to clean it up in
    various ways:
    
    1. Merge the computation of spansSize and bitmapSize. These were
       computed on every loop iteration of two different loops, but always
       have the same value, which can be derived directly from _MaxMem.
       This also avoids over-reserving these on MIPS, were _MaxArena32 is
       larger than _MaxMem.
    
    2. Remove the ulimit -v logic. It's been disabled for many releases
       and the dead code paths to support it are even more wrong now than
       they were when it was first disabled, since now we *must* reserve
       spans and bitmaps for the full address space.
    
    3. Make it clear that we're using a simple linear allocation to lay
       out the spans, bitmap, and arena spaces. Previously there were a
       lot of redundant pointer computations. Now we just bump p1 up as we
       reserve the spaces.
    
    In preparation for #18651.
    
    Updates #5049 (respect ulimit).
    
    Change-Id: Icbe66570d3a7a17bea227dc54fb3c4978b52a3af
    Reviewed-on: https://go-review.googlesource.com/35252
    Reviewed-by: Russ Cox <rsc@golang.org>

commit efb5eae3cf1c5f9be8cc5d4c85a7314204513b4c
Author: Austin Clements <austin@google.com>
Date:   Fri Jan 13 15:32:53 2017 -0500

    runtime: make _MaxMem an untyped constant
    
    Currently _MaxMem is a uintptr, which is going to complicate some
    further changes. Make it untyped so we'll be able to do untyped math
    on it before truncating it to a uintptr.
    
    The runtime assembly is identical before and after this change on
    {linux,windows}/{amd64,386}.
    
    Updates #18651.
    
    Change-Id: I0f64511faa9e0aa25179a556ab9f185ebf8c9cf8
    Reviewed-on: https://go-review.googlesource.com/35251
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 46085c4b3620fb3be29ea6ecc6206ffdb963f8bf
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Feb 7 09:31:58 2017 -0800

    cmd/compile: cmd/internal/obj: cull dead code
    
    This code is dead as a result of
    
    * removing the Follow pass
    * moving rotation detection from walk to ssa
    
    Change-Id: I14599c85bedb4e3148347b547e724187920182c4
    Reviewed-on: https://go-review.googlesource.com/36484
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 160914e33ca6521d74297291d801062cc44d794d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Feb 6 14:24:16 2017 -0500

    cmd/compile: do not use "oaslit" for global
    
    The compiler did not emit write barrier for assigning global with
    struct literal, like global = T{} where T contains pointer.
    
    The relevant code path is:
    walkexpr OAS var_ OSTRUCTLIT
        oaslit
            anylit OSTRUCTLIT
                walkexpr OAS var_ nil
                return without adding write barrier
        return true
    break (without adding write barrier)
    
    This CL makes oaslit not apply to globals. See also CL
    https://go-review.googlesource.com/c/36355/ for an alternative
    fix.
    
    The downside of this is that it generates static data for zeroing
    struct now. Also this only covers global. If there is any lurking
    bug with implicit zeroing other than globals, this doesn't fix.
    
    Fixes #18956.
    
    Change-Id: Ibcd27e4fae3aa38390ffa94a32a9dd7a802e4b37
    Reviewed-on: https://go-review.googlesource.com/36410
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1ead0bd1dc8958939b16b8fc3ab2cc8242f5e831
Author: Russ Cox <rsc@golang.org>
Date:   Tue Feb 7 11:59:38 2017 -0500

    crypto/x509: check for new tls-ca-bundle.pem last
    
    We added CentOS 7's /etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem
    to the list in response to #17549 - not being able to find any certs otherwise.
    
    Now we have #18813, where CentOS 6 apparently has both that file
    and /etc/pki/tls/certs/ca-bundle.crt, and the latter is complete while
    the former is not.
    
    Moving the new CentOS 7 file to the bottom of the list should fix both
    problems: the CentOS 7 system that didn't have any of the other files
    in the list will still find the new one, and existing systems will still
    keep using what they were using instead of preferring the new path
    that may or may not be complete on some systems.
    
    Fixes #18813.
    
    Change-Id: I5275ab67424b95e7210e14938d3e986c8caee0ba
    Reviewed-on: https://go-review.googlesource.com/36429
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit 99df7c9caa19d99747c4766be171c9487c9645cf
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Mon Feb 6 11:03:58 2017 +0000

    cmd/link, crypto/tls: don't use append loops
    
    Change-Id: Ib47e295e8646b769c30fd81e5c7f20f964df163e
    Reviewed-on: https://go-review.googlesource.com/36335
    Reviewed-by: Filippo Valsorda <hi@filippo.io>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e62aab1274c25364fe8c4609c17cb0d8e57b78d0
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 6 23:33:21 2017 -0800

    spec: clarify alignment of arrays
    
    Fixes #18950.
    
    Change-Id: I9f94748f36a896bcadc96f0642eb1f3bff387950
    Reviewed-on: https://go-review.googlesource.com/36481
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3e366ec6a7066567b4f747984389ab10cb6f1a46
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Fri Feb 3 10:18:04 2017 +0000

    testing: clarify T.Parallel() godoc wording
    
    Fixes #18914.
    
    Change-Id: Iec90d6aaa62595983db28b17794429f3c9a3dc36
    Reviewed-on: https://go-review.googlesource.com/36272
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 14347ee480968c712ea885a4ea62779fd8a0dc44
Author: Russ Cox <rsc@golang.org>
Date:   Tue Feb 7 15:21:44 2017 +0000

    Revert "image: fix the overlap check in Rectangle.Intersect."
    
    This reverts commit a855da29dbd7a80c4d87a421c1f88a8603c020fa.
    
    Change-Id: I23c0351b0708877e0b3d1b44a2bc2799cee52cd1
    Reviewed-on: https://go-review.googlesource.com/36426
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 50c7783f599d3af0e65f1c301c7ff05f6876def7
Author: Seth Vargo <sethvargo@gmail.com>
Date:   Thu Jan 19 13:19:22 2017 -0500

    text/template: remove duplicate logic in conditional
    
    It looks like this conditional may have been refactored at some point,
    but the logic was still very confusing. The outer conditional checks if
    the function is variadic, so there's no need to verify that in the
    result. Additionally, since the function isn't variadic, there is no
    reason to permit the function call if the number of input arguments is
    less than the function signature requires.
    
    Change-Id: Ia957cf83d1c900c08dd66384efcb74f0c368422e
    Reviewed-on: https://go-review.googlesource.com/35491
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit bed8129ee69f5b2d32cd84e4cc0e3e3be50366dc
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Feb 2 09:44:26 2017 -0500

    cmd/internal/obj: remove Follow pass
    
    The Follow pass in the assembler backend reorders and copies
    instructions. This even applies to hand-written assembly code,
    which in many cases don't want to be reordered. Now that the
    SSA compiler does a good job for laying out instructions, the
    benefit of this pass is very little:
    
    AMD64: (old = with Follow, new = without Follow)
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.78s ± 1%     2.79s ± 1%  +0.44%  (p=0.000 n=20+19)
    Fannkuch11-12                3.11s ± 0%     3.31s ± 1%  +6.16%  (p=0.000 n=19+19)
    FmtFprintfEmpty-12          50.9ns ± 1%    51.6ns ± 3%  +1.40%  (p=0.000 n=17+20)
    FmtFprintfString-12          127ns ± 0%     128ns ± 1%  +0.88%  (p=0.000 n=17+17)
    FmtFprintfInt-12             122ns ± 0%     123ns ± 1%  +0.76%  (p=0.000 n=20+19)
    FmtFprintfIntInt-12          185ns ± 1%     186ns ± 1%  +0.65%  (p=0.000 n=20+19)
    FmtFprintfPrefixedInt-12     192ns ± 1%     202ns ± 1%  +4.99%  (p=0.000 n=20+19)
    FmtFprintfFloat-12           284ns ± 0%     288ns ± 0%  +1.33%  (p=0.000 n=15+19)
    FmtManyArgs-12               807ns ± 0%     804ns ± 0%  -0.44%  (p=0.000 n=16+18)
    GobDecode-12                7.23ms ± 1%    7.21ms ± 1%    ~     (p=0.052 n=20+20)
    GobEncode-12                6.09ms ± 1%    6.12ms ± 1%  +0.41%  (p=0.002 n=19+19)
    Gzip-12                      253ms ± 1%     255ms ± 1%  +0.95%  (p=0.000 n=18+20)
    Gunzip-12                   38.4ms ± 0%    38.5ms ± 0%  +0.34%  (p=0.000 n=17+17)
    HTTPClientServer-12         95.4µs ± 2%    96.1µs ± 1%  +0.78%  (p=0.002 n=19+19)
    JSONEncode-12               16.5ms ± 1%    16.6ms ± 1%  +1.17%  (p=0.000 n=19+19)
    JSONDecode-12               54.6ms ± 1%    55.3ms ± 1%  +1.23%  (p=0.000 n=18+18)
    Mandelbrot200-12            4.47ms ± 0%    4.47ms ± 0%  +0.06%  (p=0.000 n=18+18)
    GoParse-12                  3.47ms ± 1%    3.47ms ± 1%    ~     (p=0.583 n=20+20)
    RegexpMatchEasy0_32-12      84.8ns ± 1%    85.2ns ± 2%  +0.51%  (p=0.022 n=20+20)
    RegexpMatchEasy0_1K-12       206ns ± 1%     206ns ± 1%    ~     (p=0.770 n=20+20)
    RegexpMatchEasy1_32-12      82.8ns ± 1%    83.4ns ± 1%  +0.64%  (p=0.000 n=20+19)
    RegexpMatchEasy1_1K-12       363ns ± 1%     361ns ± 1%  -0.48%  (p=0.007 n=20+20)
    RegexpMatchMedium_32-12      126ns ± 1%     126ns ± 0%  +0.72%  (p=0.000 n=20+20)
    RegexpMatchMedium_1K-12     39.1µs ± 1%    39.8µs ± 0%  +1.73%  (p=0.000 n=19+19)
    RegexpMatchHard_32-12       1.97µs ± 0%    1.98µs ± 1%  +0.29%  (p=0.005 n=18+20)
    RegexpMatchHard_1K-12       59.5µs ± 1%    59.8µs ± 1%  +0.36%  (p=0.000 n=18+20)
    Revcomp-12                   442ms ± 1%     445ms ± 2%  +0.67%  (p=0.000 n=19+20)
    Template-12                 58.0ms ± 1%    57.5ms ± 1%  -0.85%  (p=0.000 n=19+19)
    TimeParse-12                 311ns ± 0%     314ns ± 0%  +0.94%  (p=0.000 n=20+18)
    TimeFormat-12                350ns ± 3%     346ns ± 0%    ~     (p=0.076 n=20+19)
    [Geo mean]                  55.9µs         56.4µs       +0.80%
    
    ARM32:
    name                     old time/op    new time/op    delta
    BinaryTree17-4              30.4s ± 0%     30.1s ± 0%  -1.14%  (p=0.000 n=10+8)
    Fannkuch11-4                13.7s ± 0%     13.6s ± 0%  -0.75%  (p=0.000 n=10+10)
    FmtFprintfEmpty-4           664ns ± 1%     651ns ± 1%  -1.96%  (p=0.000 n=7+8)
    FmtFprintfString-4         1.83µs ± 2%    1.77µs ± 2%  -3.21%  (p=0.000 n=10+10)
    FmtFprintfInt-4            1.57µs ± 2%    1.54µs ± 2%  -2.25%  (p=0.007 n=10+10)
    FmtFprintfIntInt-4         2.37µs ± 2%    2.31µs ± 1%  -2.68%  (p=0.000 n=10+10)
    FmtFprintfPrefixedInt-4    2.14µs ± 2%    2.10µs ± 1%  -1.83%  (p=0.006 n=10+10)
    FmtFprintfFloat-4          3.69µs ± 2%    3.74µs ± 1%  +1.60%  (p=0.000 n=10+10)
    FmtManyArgs-4              9.43µs ± 1%    9.17µs ± 1%  -2.70%  (p=0.000 n=10+10)
    GobDecode-4                76.3ms ± 1%    75.5ms ± 1%  -1.14%  (p=0.003 n=10+10)
    GobEncode-4                70.7ms ± 2%    69.0ms ± 1%  -2.36%  (p=0.000 n=10+10)
    Gzip-4                      2.64s ± 1%     2.65s ± 0%  +0.59%  (p=0.002 n=10+10)
    Gunzip-4                    402ms ± 0%     398ms ± 0%  -1.11%  (p=0.000 n=10+9)
    HTTPClientServer-4          458µs ± 0%     457µs ± 0%    ~     (p=0.247 n=10+10)
    JSONEncode-4                171ms ± 0%     172ms ± 0%  +0.56%  (p=0.000 n=10+10)
    JSONDecode-4                672ms ± 1%     668ms ± 1%    ~     (p=0.105 n=10+10)
    Mandelbrot200-4            33.5ms ± 0%    33.5ms ± 0%    ~     (p=0.156 n=9+10)
    GoParse-4                  33.9ms ± 0%    34.0ms ± 0%  +0.36%  (p=0.031 n=9+9)
    RegexpMatchEasy0_32-4       823ns ± 1%     835ns ± 1%  +1.49%  (p=0.000 n=8+8)
    RegexpMatchEasy0_1K-4      3.99µs ± 0%    4.02µs ± 1%  +0.92%  (p=0.000 n=8+10)
    RegexpMatchEasy1_32-4       877ns ± 3%     904ns ± 2%  +3.07%  (p=0.012 n=10+10)
    RegexpMatchEasy1_1K-4      5.99µs ± 0%    5.97µs ± 1%  -0.38%  (p=0.023 n=8+8)
    RegexpMatchMedium_32-4     1.40µs ± 2%    1.40µs ± 2%    ~     (p=0.590 n=10+9)
    RegexpMatchMedium_1K-4      357µs ± 0%     355µs ± 1%  -0.72%  (p=0.000 n=7+8)
    RegexpMatchHard_32-4       22.3µs ± 0%    22.1µs ± 0%  -0.49%  (p=0.000 n=8+7)
    RegexpMatchHard_1K-4        661µs ± 0%     658µs ± 0%  -0.42%  (p=0.000 n=8+7)
    Revcomp-4                  46.3ms ± 0%    46.3ms ± 0%    ~     (p=0.393 n=10+10)
    Template-4                  753ms ± 1%     750ms ± 0%    ~     (p=0.211 n=10+9)
    TimeParse-4                4.28µs ± 1%    4.22µs ± 1%  -1.34%  (p=0.000 n=8+10)
    TimeFormat-4               9.00µs ± 0%    9.05µs ± 0%  +0.59%  (p=0.000 n=10+10)
    [Geo mean]                  538µs          535µs       -0.55%
    
    ARM64:
    name                     old time/op    new time/op    delta
    BinaryTree17-8              8.39s ± 0%     8.39s ± 0%    ~     (p=0.684 n=10+10)
    Fannkuch11-8                5.95s ± 0%     5.99s ± 0%  +0.63%  (p=0.000 n=10+10)
    FmtFprintfEmpty-8           116ns ± 0%     116ns ± 0%    ~     (all equal)
    FmtFprintfString-8          361ns ± 0%     360ns ± 0%  -0.31%  (p=0.003 n=8+6)
    FmtFprintfInt-8             290ns ± 0%     290ns ± 0%    ~     (p=0.620 n=9+9)
    FmtFprintfIntInt-8          476ns ± 1%     469ns ± 0%  -1.47%  (p=0.000 n=10+6)
    FmtFprintfPrefixedInt-8     412ns ± 2%     417ns ± 2%  +1.39%  (p=0.006 n=9+10)
    FmtFprintfFloat-8           652ns ± 1%     652ns ± 0%    ~     (p=0.161 n=10+8)
    FmtManyArgs-8              1.94µs ± 0%    1.94µs ± 2%    ~     (p=0.781 n=10+10)
    GobDecode-8                17.7ms ± 1%    17.7ms ± 0%    ~     (p=0.962 n=10+7)
    GobEncode-8                15.6ms ± 0%    15.6ms ± 1%    ~     (p=0.063 n=10+10)
    Gzip-8                      786ms ± 0%     787ms ± 0%    ~     (p=0.356 n=10+9)
    Gunzip-8                    127ms ± 0%     127ms ± 0%  +0.08%  (p=0.028 n=10+9)
    HTTPClientServer-8          198µs ± 6%     198µs ± 7%    ~     (p=0.796 n=10+10)
    JSONEncode-8               42.5ms ± 0%    42.2ms ± 0%  -0.73%  (p=0.000 n=9+8)
    JSONDecode-8                158ms ± 1%     162ms ± 0%  +2.28%  (p=0.000 n=10+9)
    Mandelbrot200-8            10.1ms ± 0%    10.1ms ± 0%  -0.01%  (p=0.000 n=10+9)
    GoParse-8                  8.54ms ± 1%    8.63ms ± 1%  +1.06%  (p=0.000 n=10+9)
    RegexpMatchEasy0_32-8       231ns ± 1%     225ns ± 0%  -2.52%  (p=0.000 n=9+10)
    RegexpMatchEasy0_1K-8      1.63µs ± 0%    1.63µs ± 0%    ~     (p=0.170 n=10+10)
    RegexpMatchEasy1_32-8       253ns ± 0%     249ns ± 0%  -1.41%  (p=0.000 n=9+10)
    RegexpMatchEasy1_1K-8      2.08µs ± 0%    2.08µs ± 0%  -0.32%  (p=0.000 n=9+10)
    RegexpMatchMedium_32-8      355ns ± 1%     351ns ± 0%  -1.04%  (p=0.007 n=10+7)
    RegexpMatchMedium_1K-8      104µs ± 0%     104µs ± 0%    ~     (p=0.148 n=10+10)
    RegexpMatchHard_32-8       5.79µs ± 0%    5.79µs ± 0%    ~     (p=0.578 n=10+10)
    RegexpMatchHard_1K-8        176µs ± 0%     176µs ± 0%    ~     (p=0.137 n=10+10)
    Revcomp-8                   1.37s ± 1%     1.36s ± 1%  -0.26%  (p=0.023 n=10+10)
    Template-8                  151ms ± 1%     154ms ± 1%  +2.14%  (p=0.000 n=9+10)
    TimeParse-8                 723ns ± 2%     721ns ± 1%    ~     (p=0.592 n=10+10)
    TimeFormat-8                804ns ± 2%     798ns ± 3%    ~     (p=0.344 n=10+10)
    [Geo mean]                  154µs          154µs       -0.02%
    
    Therefore remove this pass. Also reduce text size by 0.5~2%.
    
    Comment out some dead code in runtime/sys_nacl_amd64p32.s
    which contains undefined symbols.
    
    Change-Id: I1473986fe5b18b3d2554ce96cdc6f0999b8d955d
    Reviewed-on: https://go-review.googlesource.com/36205
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 76d427449106d0eb2842d31934e2cea9f049a562
Author: Mura Li <mura_li@castech.com.tw>
Date:   Mon Aug 29 10:22:20 2016 +0800

    crypto/des: improve the throughput of DES and 3DES
    
    For detailed explanation of the adopted (Eric Young's) algorithm,
    see http://ftp.nluug.nl/security/coast/libs/libdes/ALGORITHM
    
    benchmark                   old ns/op     new ns/op     delta
    BenchmarkEncrypt-16         649           164           -74.73%
    BenchmarkDecrypt-16         546           156           -71.43%
    BenchmarkTDESEncrypt-16     1651          385           -76.68%
    BenchmarkTDESDecrypt-16     1645          378           -77.02%
    
    benchmark                   old MB/s     new MB/s     speedup
    BenchmarkEncrypt-16         12.31        48.76        3.96x
    BenchmarkDecrypt-16         14.64        51.03        3.49x
    BenchmarkTDESEncrypt-16     4.84         20.74        4.29x
    BenchmarkTDESDecrypt-16     4.86         21.16        4.35x
    
    Change-Id: Ic3e1fe3340419ec5a0e6379434911eb41e0246f6
    Reviewed-on: https://go-review.googlesource.com/36490
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 08bb7ccb75cd6ca5c52c5a09386a6479168607d6
Author: Alan Donovan <adonovan@google.com>
Date:   Fri Dec 9 09:55:50 2016 -0500

    go/types: permit f(nil...) for variadic arguments
    
    This code may be pointless, but it is legal.
    
    Fixes golang/go#18268
    
    Change-Id: Ibacae583606e1a6fdf0c0f01abe2e22e9e608393
    Reviewed-on: https://go-review.googlesource.com/34194
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit a855da29dbd7a80c4d87a421c1f88a8603c020fa
Author: Nigel Tao <nigeltao@golang.org>
Date:   Thu Jan 5 17:37:54 2017 +1100

    image: fix the overlap check in Rectangle.Intersect.
    
    The doc comment for Rectangle.Intersect clearly states, "If the two
    rectangles do not overlap then the zero rectangle will be returned."
    Prior to this fix, calling Intersect on adjacent but non-overlapping
    rectangles would return an empty but non-zero rectangle.
    
    The fix essentially changes
    if r.Min.X > r.Max.X || r.Min.Y > r.Max.Y { etc }
    to
    if r.Min.X >= r.Max.X || r.Min.Y >= r.Max.Y { etc }
    (note that the > signs have become >= signs), but changing that line to:
    if r.Empty() { etc }
    seems clearer (and equivalent).
    
    Change-Id: Ia654e4b9dc805978db3e94d7a9718b6366005360
    Reviewed-on: https://go-review.googlesource.com/34853
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit cbef450df797c473c9ca01f8d0c81ea26d106c24
Author: Michael Matloob <matloob@golang.org>
Date:   Thu Dec 8 16:54:07 2016 -0500

    runtime/pprof: symbolize proto profiles
    
    When generating pprof profiles in proto format, symbolize the profiles.
    
    Change-Id: I2471ed7f919483e5828868306418a63e41aff5c5
    Reviewed-on: https://go-review.googlesource.com/34192
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 936749efb07f87f99753f47d921e7659414fad2d
Author: Shintaro Kaneko <kaneshin0120@gmail.com>
Date:   Sat Jan 28 10:08:10 2017 +0000

    test: improve output format of issue10607a.go test
    
    Change-Id: Iad5ff820a95f5082b75aa5260e40c33c7b0ecf22
    Reviewed-on: https://go-review.googlesource.com/35990
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 53c6ac54190ae21dd1a7dacf7f066785834407b2
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 6 22:01:07 2017 -0800

    cmd/compile/internal/syntax: avoid follow-up error for incorrect if statement
    
    This is a follow-up on https://go-review.googlesource.com/36470
    and leads to a more stable fix. The above CL relied on filtering
    of multiple errors on the same line to avoid more than one error
    for an `if` statement of the form `if a := 10 {}`. This CL avoids
    the secondary error ("missing condition in if statement") in the
    first place.
    
    For #18915.
    
    Change-Id: I8517f485cc2305965276c17d8f8797d61ef9e999
    Reviewed-on: https://go-review.googlesource.com/36479
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6b742b2f84fc4ddea27076aa1e581197d17bd863
Author: Quentin Smith <quentin@golang.org>
Date:   Mon Feb 6 11:59:01 2017 -0500

    testing: print extra labels on benchmarks
    
    When running benchmarks, print "goos", "goarch", and "pkg"
    labels. This makes it easier to refer to benchmark logs and understand
    how they were generated. "pkg" is printed only for benchmarks located
    in GOPATH.
    
    Change-Id: I397cbdd57b9fe8cbabbb354ec7bfba59f5625c42
    Reviewed-on: https://go-review.googlesource.com/36356
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit c0bd4f33ccc9a9454d50245a1dba1fa46e62a1ad
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 6 15:57:00 2017 -0800

    spec: pick up a few corrections missed in prior commit
    
    This CL picks up a couple of minor fixes that were present
    in https://go-review.googlesource.com/#/c/36213/6..5 but
    accidentally got dropped in https://go-review.googlesource.com/#/c/36213/
    because I submitted from the wrong client.
    
    Change-Id: I3ad0d20457152ea9a116cbb65a23eb0dc3a8525e
    Reviewed-on: https://go-review.googlesource.com/36471
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 56c9b51b937cca7d3db517add96bd9517bbffb80
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 2 15:43:56 2017 -0800

    spec: introduce alias declarations and type definitions
    
    To avoid confusion caused by the term "named type" (which now just
    means a type with a name, but formerly meant a type declared with
    a non-alias type declaration), a type declaration now comes in two
    forms: alias declarations and type definitions. Both declare a type
    name, but type definitions also define new types.
    
    Replace the use of "named type" with "defined type" elsewhere in
    the spec.
    
    For #18130.
    
    Change-Id: I49f5ddacefce90354eb65ee5fbf10ba737221995
    Reviewed-on: https://go-review.googlesource.com/36213
    Reviewed-by: Rob Pike <r@golang.org>

commit 3b68a647696ebfb61d199155f5f1faa5740e5c55
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 6 15:08:28 2017 -0800

    cmd/compile/internal/syntax: make a parser error "1.7 compliant"
    
    For code such as
    
            if a := 10 { ...
    
    the 1.7 compiler reported
    
            a := 10 used as value
    
    while the 1.8 compiler reported
    
            invalid condition, tag, or type switch guard
    
    Changed the error message to match the 1.7 compiler.
    
    Fixes #18915.
    
    Change-Id: I01308862e461922e717f9f8295a9db53d5a914eb
    Reviewed-on: https://go-review.googlesource.com/36470
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6a29440dcc5b71ded72d35e00a26d96c401f49d4
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Feb 6 14:46:48 2017 -0800

    cmd/compile/internal/gc: remove more backend Sym uses
    
    Removes all external uses of Linksym and Pkglookup, which are the only
    two exported functions that return Syms.
    
    Also add Duffcopy and Duffzero since they're used often enough across
    SSA backends.
    
    Passes toolstash -cmp.
    
    Change-Id: I8d3fd048ad5cd676fc46378f09a917569ffc9b2c
    Reviewed-on: https://go-review.googlesource.com/36418
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit ab067cde34b515172cf51b4c562b2b4ef3cbe587
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Feb 6 17:52:26 2017 -0500

    cmd/link: use external linking for PIE by default
    
    Now `go test -buildmode=pie std -short` passes on linux/amd64.
    
    Updates #18968
    
    Change-Id: Ide21877713e00edc64c1700c950016d6bff8de0e
    Reviewed-on: https://go-review.googlesource.com/36417
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5f374ea8fb3754703a01ddf94e729f926317bf67
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Feb 6 13:40:19 2017 -0800

    cmd/compile/internal/gc: stop exporting *gc.Sym-typed globals
    
    The arch-specific SSA backends now no longer use gc.Sym either.
    
    Passes toolstash -cmp.
    
    Change-Id: Ic13b934b92a1b89b4b79c6c4796ab0a137608163
    Reviewed-on: https://go-review.googlesource.com/36416
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 87c475c2276977a0bf1208f95884261b0426fddc
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Feb 6 13:30:40 2017 -0800

    cmd/compile/internal/ssa: use obj.LSym instead of gc.Sym
    
    Gc's Sym type represents a package-qualified identifier, which is a
    frontend concept and doesn't belong in SSA. Bonus: we can replace some
    interface{} types with *obj.LSym.
    
    Passes toolstash -cmp.
    
    Change-Id: I456eb9957207d80f99f6eb9b8eab4a1f3263e9ed
    Reviewed-on: https://go-review.googlesource.com/36415
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit c2bc727f944fe41e1567b8ac1cf306809801bec9
Author: Andrew Gerrand <adg@golang.org>
Date:   Tue Feb 7 08:05:01 2017 +1100

    doc: remove inactive members of the CoC working group
    
    Dave and Jason have moved on to other things.
    
    Change-Id: I702d11bedfab1f47a33679a48c2309f49021229e
    Reviewed-on: https://go-review.googlesource.com/36450
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit 62956897c1743ee7e79895496180f84432f21d0a
Author: Michael Matloob <matloob@golang.org>
Date:   Fri Dec 9 16:00:02 2016 -0500

    runtime: add definitions for SetGoroutineLabels and Do
    
    This change defines runtime/pprof.SetGoroutineLabels and runtime/pprof.Do, which
    are used to set profiler labels on goroutines. The change defines functions
    in the runtime for setting and getting profile labels, and sets and unsets
    profile labels when goroutines are created and deleted. The change also adds
    the package runtime/internal/proflabel, which defines the structure the runtime
    uses to store profile labels.
    
    Change-Id: I747a4400141f89b6e8160dab6aa94ca9f0d4c94d
    Reviewed-on: https://go-review.googlesource.com/34198
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-on: https://go-review.googlesource.com/35010

commit bc548d71b9634341c514dcef10ce8141f330d95c
Author: Russ Cox <rsc@golang.org>
Date:   Mon Feb 6 14:41:12 2017 -0500

    vendor/golang.org/x/crypto/curve25519: avoid loss of R15 in -dynlink mode
    
    Original code fixed in https://go-review.googlesource.com/#/c/36359/.
    
    Fixes #18820.
    
    Change-Id: I060e6c9d0e312b4fd5d0674aff131055bf5cf61d
    Reviewed-on: https://go-review.googlesource.com/36412
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit 807c80fce346eb2e895dc3de7be8fe0ad70c3894
Author: Keith Randall <khr@golang.org>
Date:   Wed Jan 11 16:40:24 2017 -0800

    cmd/compile: using CONV instead of CONVNOP for interface conversions
    
    We shouldn't use CONVNOP for conversions between two different
    nonempty interface types, because we want to update the itab
    in those situations.
    
    Fixes #18595
    
    After this CL, we are guaranteed that itabs are unique, that is
    there is only one itab per compile-time-type/concrete type pair.
    See also the tests in CL 35115 and 35116 which make sure this
    invariant holds even for shared libraries and plugins.
    
    Unique itabs are required for CL 34810 (faster type switch code).
    
    R=go1.9
    
    Change-Id: Id27d2e01ded706680965e4cb69d7c7a24ac2161b
    Reviewed-on: https://go-review.googlesource.com/35119
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit fdbae7d77ed1d365f1fd6735d234f7c277102f12
Author: Sameer Ajmani <sameer@golang.org>
Date:   Mon Feb 6 13:43:11 2017 -0500

    net/http/httputil: don't log read error when it's context.Canceled
    
    Fixes #18838
    
    Change-Id: I44976cadb0dc3c23eacb8cdd58429a572cd8d28a
    Reviewed-on: https://go-review.googlesource.com/36358
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2ca5d105b00353d3f3aa4da023e405b0550117d0
Author: Sameer Ajmani <sameer@golang.org>
Date:   Mon Feb 6 13:21:57 2017 -0500

    os/user: cache the result of user.Current
    
    This has a notable impact on systems with very large passwd files.
    
    Before:
    BenchmarkCurrent-12        30000             42546 ns/op
    
    After:
    BenchmarkCurrent-12     20000000                77.5 ns/op
    
    Saved in perf dashboard:
    https://perf.golang.org/search?q=upload:20170206.1
    
    Fixes #11625
    
    Change-Id: Iebc9bf122cc64a4cab24ac06843c7b2bc450ded9
    Reviewed-on: https://go-review.googlesource.com/36391
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit fd37b8ccf2262bb3f0a608f7545f78a72e8d661f
Author: David R. Jenni <david.r.jenni@gmail.com>
Date:   Sun Feb 5 11:02:03 2017 +0100

    sort: optimize average calculation in binary search
    
    Use fewer instructions to calculate the average of i and j without
    overflowing at the addition.
    
    Even if both i and j are math.MaxInt{32,64}, the sum fits into a
    uint{32,64}. Because the sum of i and j is always ≥ 0, the right
    shift by one does the same as a division by two. The result of the
    shift operation is at most math.MaxInt{32,64} and fits again into
    an int{32,64}.
    
    name              old time/op  new time/op  delta
    SearchWrappers-4   153ns ± 3%   143ns ± 6%  -6.33%  (p=0.000 n=90+100)
    
    This calculation is documented in:
    https://research.googleblog.com/2006/06/extra-extra-read-all-about-it-nearly.html
    
    Change-Id: I2be7922afc03b3617fce32e59364606c37a83678
    Reviewed-on: https://go-review.googlesource.com/36332
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 91ad2a219445d6df3ddb796d0f44190da24f429d
Author: Michael Matloob <matloob@golang.org>
Date:   Fri Dec 9 16:00:02 2016 -0500

    runtime/pprof: add definitions of profile label types
    
    This change defines WithLabels, Labels, Label, and ForLabels.
    This is the first step of the profile labels implemention for go 1.9.
    
    Updates #17280
    
    Change-Id: I2dfc9aae90f7a4aa1ff7080d5747f0a1f0728e75
    Reviewed-on: https://go-review.googlesource.com/34198
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 47d2a4dafa6e84f834f677790449c3c5998a5b98
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Feb 4 21:20:23 2017 -0800

    cmd/compile: remove walkmul
    
    Replace with generic rewrite rules.
    
    Change-Id: I3ee32076cfd9db5801f1a7bdbb73a994255884a9
    Reviewed-on: https://go-review.googlesource.com/36323
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 6aee6b895c219300c9c089b79f81c80f0e468dd7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sat Feb 4 14:05:20 2017 -0800

    runtime: remove markBits.clearMarkedNonAtomic
    
    It's not used, it's never been used, and it doesn't do what its doc
    comment says it does.
    
    Fixes #18941.
    
    Change-Id: Ia89d97fb87525f5b861d7701f919e0d6b7cbd376
    Reviewed-on: https://go-review.googlesource.com/36322
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 172311ce10c823c470d23060317f79c5059f041b
Author: Alexey Palazhchenko <alexey.palazhchenko@gmail.com>
Date:   Sun Feb 5 23:06:34 2017 +0300

    time: Fix typo in Time.String() description.
    
    Refs #12914.
    
    Change-Id: Iadac4cbef70db6a95b47f86eaffcfc63bfdb8e90
    Reviewed-on: https://go-review.googlesource.com/36334
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b53f0f8c96a46b3cce0f1073787b74dd23f57a80
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Feb 3 20:54:52 2017 -0500

    cmd/compile: do not fold large offset on ARM64
    
    Fixes #18933.
    
    Change-Id: I8bb98e95bb4486a086d93bcf99e3a37488e77b03
    Reviewed-on: https://go-review.googlesource.com/36318
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 34b455da4484dee20e8c355b50d24680224b58f2
Author: Yasuhiro Matsumoto <mattn.jp@gmail.com>
Date:   Wed Jan 25 16:48:17 2017 +0900

    path/filepath: ignore dot for Dir(`\\server\share`)
    
    Dir(`\\server\share`) returns `\\server\share.`. Change Dir so it
    returns `\\server\share` instead.
    
    Fixes #18783
    
    Change-Id: I9e0dd71ea6aea85e6c6114aaa4bb3bea3270d818
    Reviewed-on: https://go-review.googlesource.com/35690
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 769be04feb724e03c1f2b757fc19326a1486896c
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sat Feb 4 17:33:14 2017 +1100

    cmd/nm: skip TestInternalLinkerCgoFile if no internal linking is supported
    
    Fixes build.
    
    Change-Id: I2fee624c8a4b228bb9f2889e241ea016a317bb11
    Reviewed-on: https://go-review.googlesource.com/36373
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit c7a7c5a9b425259e17976b978b60651b636b8979
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Fri Jan 13 18:02:07 2017 +1100

    cmd/link: do not prefix external symbols with underscore on windows/386/cgo
    
    CL 18057 added underscore to most external pe symbols
    on windows/386/cgo. The CL changed runtime.epclntab and
    runtime.pclntab pe symbols into _runtime.pclntab and
    _runtime.epclntab, and now cmd/nm cannot find them.
    Revert correspondent CL 18057 changes, because most pe
    symbols do not need underscore prefix.
    
    This CL also removes code that added obj.SHOSTOBJ symbols
    explicitly, because each of those was also added via
    genasmsym call. These created duplicate pe symbols (like
    _GetProcAddress@8 and __GetProcAddress@8), and external
    linker would complain.
    
    This CL adds new test in cmd/nm to verify go programs
    built with cgo.
    
    Fixes #18416
    
    Change-Id: I68b1be8fb631d95ec69bd485c77c79604fb23f26
    Reviewed-on: https://go-review.googlesource.com/35076
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit afa0247c5d28eb9558311729c8edf3f0c898644f
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Nov 23 08:45:15 2016 -0800

    buildall.bash: clarify target selection
    
    Avoid confusing use of $(( in non-arithmetic context.
    
    Permit added targets linux-386-387 linux-arm-arm5 to be correctly
    matched against pattern argument.
    
    Change-Id: Ib004c926457acb760c7e270fdd2f4095b1787a6d
    Reviewed-on: https://go-review.googlesource.com/33492
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6ee8c6a7ce30a3c0613617dcc182874d36b2ea94
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Feb 3 14:45:26 2017 -0800

    cmd/compile/internal/gc: simplify generating static data
    
    Passes toolstash -cmp.
    
    Change-Id: I4a72e3e130c38868ee8ecef32cad58748aa5be52
    Reviewed-on: https://go-review.googlesource.com/36353
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ad784caa746515c28033435318c8b82c187fa583
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Oct 28 11:42:40 2016 -0700

    cmd/compile: move Heapaddr field from Name to Param
    
    No performance impact, just cleanup.
    
    Passes toolstash -cmp.
    
    Change-Id: Ic7957d2686de53a9680c2bdefe926cccccd73a5c
    Reviewed-on: https://go-review.googlesource.com/36316
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5c90e1cf8ab39625b1f73c499cf47c06a60e9c08
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Feb 3 14:28:32 2017 -0800

    cmd/compile/internal/ssa: remove Func.StaticData field
    
    Rather than collecting static data nodes to be written out later, just
    write them out immediately.
    
    Change-Id: I51708b690e94bc3e288b4d6ba3307bf738a80f64
    Reviewed-on: https://go-review.googlesource.com/36352
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 592c97fc8f55d65dc668b1acb3b5bdf46e851f9e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Feb 3 14:38:47 2017 -0800

    cmd/dist: ignore .#foo.go files created by Emacs
    
    go/build already ignores them, but they cause make.bash to fail.
    
    Fixes #18931.
    
    Change-Id: Idd5c8c2a6f2309ecd5f0d669660704d6f5612710
    Reviewed-on: https://go-review.googlesource.com/36351
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7751d56e3955aaf5285f16b1643e6d8153e32b80
Author: Peter Nguyen <peter@mictis.com>
Date:   Fri Feb 3 21:45:50 2017 +0100

    net/rpc/jsonrpc: Update package doc with info about JSON-RPC 2.0
    
    Currently the net/rpc/jsonrpc package only implements JSON-RPC version
    1.0. This change updates the package's documentation with link to find
    packages for JSON-RPC 2.0.
    
    Fixes #10929
    
    Change-Id: I3b6f1d17738a1759d7b62ab7b3ecef5b248d30ca
    Reviewed-on: https://go-review.googlesource.com/36330
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ddf807fce821c26da32f653a2483bb7d96b20e26
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Feb 3 04:55:34 2017 -0500

    cmd/compile: fix type propagation through s390x SSA rules
    
    This CL fixes two issues:
    
    1. Load ops were initially always lowered to unsigned loads, even
       for signed types. This was fine by itself however LoadReg ops
       (used to re-load spilled values) were lowered to signed loads
       for signed types. This meant that spills could invalidate
       optimizations that assumed the original unsigned load.
    
    2. Types were not always being maintained correctly through rules
       designed to eliminate unnecessary zero and sign extensions.
    
    Fixes #18906.
    
    Change-Id: I95785dcadba03f7e3e94524677e7d8d3d3b9b737
    Reviewed-on: https://go-review.googlesource.com/36256
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 178307c3a72a9da3d731fecf354630761d6b246c
Author: Russ Cox <rsc@golang.org>
Date:   Fri Feb 3 13:29:06 2017 -0500

    cmd/go: address review comments
    
    Address review comments from earlier CLs.
    These are changes I was too scared to try to push
    down into the original CLs (thanks, Git).
    
    Change-Id: I0e428fad73d71bd2a7d08178cf2e856de3cef19f
    Reviewed-on: https://go-review.googlesource.com/36257
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 707cadd7fa1c8653a7b3409be4dc79823c45306c
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 13:49:50 2017 -0500

    cmd/go: split out cmd/go/internal/clean,doc,fix,generate,list,run,tool,version,vet
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: Ib22fc435827d4a05a77a5200ac437ce00e2a4da3
    Reviewed-on: https://go-review.googlesource.com/36204
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 76db88ab4db6c95731d32dc4eefafa674db1203a
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 13:40:24 2017 -0500

    cmd/go: split out cmd/go/internal/bug
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I05629567cc33fef41bc74eba4f7ff66e4851343c
    Reviewed-on: https://go-review.googlesource.com/36203
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 6dad55aa25debdeb43e6b7c8348cf923fb8e5cff
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 13:39:12 2017 -0500

    cmd/go: split out cmd/go/internal/get
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: Iec17bf2243de129942ae5fba126ec5f217be7303
    Reviewed-on: https://go-review.googlesource.com/36202
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit cc03ba3289eb1a3bd4b6454fff24646912f5bf12
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 13:36:46 2017 -0500

    cmd/go: split out cmd/go/internal/web
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I2f349150659b6ddf6be4c675abba38dfe57ff652
    Reviewed-on: https://go-review.googlesource.com/36201
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 6bc9844b63efaa5e29e3c93c1d27477b1663db09
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 13:24:01 2017 -0500

    cmd/go: split out cmd/go/internal/env
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I28b20d53d20dff06eede574eb5c20359db0d3991
    Reviewed-on: https://go-review.googlesource.com/36200
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 36ce197c858ea9e187ee59fc18576c73a0a513e7
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 13:18:56 2017 -0500

    cmd/go: split out cmd/go/internal/fmt
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I8e325d75f553b5d0b6224b56a705d2e2cb895de4
    Reviewed-on: https://go-review.googlesource.com/36199
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 4efe9250e5e182c8269a2b98497a9bdea4875c8f
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 12:56:37 2017 -0500

    cmd/go: split out cmd/go/internal/test
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I2d0ccdb84814537ab8b8842aa1b5f5bc0a88a0fc
    Reviewed-on: https://go-review.googlesource.com/36198
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 3c667ef421becd1acdbd8ac57503a1cbed47fc5c
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 00:06:08 2017 -0500

    cmd/go: split out cmd/go/internal/work
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: Icdd181098f9f0e81f68bf201e6867cdd8f820300
    Reviewed-on: https://go-review.googlesource.com/36197
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit eb93b20c2ebb9cba89657437615957d5f45b5f6c
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 13 16:24:19 2017 -0500

    cmd/go: split out cmd/go/internal/load
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: Ic802483e50598def638f1e2e706d5fdf7822d32d
    Reviewed-on: https://go-review.googlesource.com/36196
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 461c3e52633638ee6e385a6fed3d78dc0a02f210
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 13 14:59:37 2017 -0500

    cmd/go: split out cmd/go/internal/buildid
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I20dbc352c3df3c83a75811dd8e78c580a46b2202
    Reviewed-on: https://go-review.googlesource.com/36195
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 2cab41d5cb813956a367c464afc0b16b2c418331
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 13 14:51:57 2017 -0500

    cmd/go: split out cmd/go/internal/help
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I4cf05b076d81b780c87a31378523929b5da8964b
    Reviewed-on: https://go-review.googlesource.com/36194
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 6dc9e31f5e86112737bb60e54b6f903929ca11d8
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 13 14:41:42 2017 -0500

    cmd/go: split out cmd/go/internal/base
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I7c5dde6e7fe4f390e6607303b4d42535c674eac3
    Reviewed-on: https://go-review.googlesource.com/36193
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit b60e61ab1bf7cca20a03364143e76acf2925ccb6
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 18 12:58:25 2017 -0500

    cmd/dist: move cmd/go z files to cmd/go/internal/cfg
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I6ee5b053683034ea9462a9a0a4ea4f5ad24fa5a1
    Reviewed-on: https://go-review.googlesource.com/36192
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit a918864c52355d2f73657825b4e0905c9d7fa279
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 13 14:10:06 2017 -0500

    cmd/go: split out cmd/go/internal/cfg
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: Icb3f168ade91e7da5fcab89ac75b768daefff359
    Reviewed-on: https://go-review.googlesource.com/36191
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 762eb40892061b43a2dabbfd99486e9940a149b9
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 13 13:52:44 2017 -0500

    cmd/go: split out cmd/go/internal/str
    
    This is one CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I63f578f5ac99c707b599ac5659293c46b275567d
    Reviewed-on: https://go-review.googlesource.com/36190
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit d9e6835b1d0ac90d82c62e33a2aa0daf602940b0
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 13 11:49:16 2017 -0500

    cmd/go: break a few dependencies
    
    This CL makes a few naming changes to break dependencies
    between different parts of the go command, to make it easier
    to split into different packages.
    
    This is the first CL in a long sequence of changes to break up the
    go command from one package into a plausible group of packages.
    
    This sequence is concerned only with moving code, not changing
    or cleaning up code. There will still be more cleanup after this sequence.
    
    The entire sequence will be submitted together: it is not a goal
    for the tree to build at every step.
    
    For #18653.
    
    Change-Id: I69a98b9ea48e61b1e1cda95273d29860b525415f
    Reviewed-on: https://go-review.googlesource.com/36129
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 78074f6850c34a955d69f578e363d1d3f3e00e25
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sun Jan 29 15:34:50 2017 +0100

    runtime: handle SIGPIPE in c-archive and c-shared programs
    
    Before this CL, Go programs in c-archive or c-shared buildmodes
    would not handle SIGPIPE. That leads to surprising behaviour where
    writes on a closed pipe or socket would raise SIGPIPE and terminate
    the program. This CL changes the Go runtime to handle
    SIGPIPE regardless of buildmode. In addition, SIGPIPE from non-Go
    code is forwarded.
    
    This is a refinement of CL 32796 that fixes the case where a non-default
    handler for SIGPIPE is installed by the host C program.
    
    Fixes #17393
    
    Change-Id: Ia41186e52c1ac209d0a594bae9904166ae7df7de
    Reviewed-on: https://go-review.googlesource.com/35960
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b612ab3acbf3a11ea6dbaac8f244b4bdfed308cd
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jan 14 23:43:26 2017 -0800

    cmd/compile: make liveness more efficient
    
    When the number of variables in a function is very large,
    liveness analysis gets less efficient, since every bit vector
    is O(number of variables).
    
    Improve the situation by returning a sparse representation
    from progeffects. In all scenarios, progeffects either
    returns a slice that is shared function-wide,
    and which is usually small, or a slice that is guaranteed
    to have at most three values.
    
    Reduces compilation time for the code in #8225 Comment 1 by ~10%.
    Minor effects on regular packages (below).
    
    Passes toolstash -cmp.
    
    Updates #8225
    
    name       old time/op      new time/op      delta
    Template        215ms ± 2%       212ms ± 4%  -1.31%  (p=0.001 n=30+30)
    Unicode        98.3ms ± 3%      98.4ms ± 5%    ~     (p=0.971 n=30+30)
    GoTypes         657ms ± 3%       651ms ± 2%  -0.98%  (p=0.001 n=30+27)
    Compiler        2.78s ± 2%       2.77s ± 2%  -0.60%  (p=0.006 n=30+30)
    Flate           130ms ± 4%       130ms ± 4%    ~     (p=0.712 n=29+30)
    GoParser        159ms ± 5%       158ms ± 3%    ~     (p=0.331 n=29+30)
    Reflect         406ms ± 3%       404ms ± 3%  -0.69%  (p=0.041 n=29+30)
    Tar             117ms ± 4%       117ms ± 3%    ~     (p=0.886 n=30+29)
    XML             219ms ± 2%       217ms ± 2%    ~     (p=0.091 n=29+24)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   272user-ms ± 3%  270user-ms ± 3%  -1.03%  (p=0.004 n=30+30)
    Unicode    138user-ms ± 2%  138user-ms ± 3%    ~     (p=0.902 n=29+29)
    GoTypes    891user-ms ± 2%  883user-ms ± 2%  -0.95%  (p=0.000 n=29+29)
    Compiler   3.85user-s ± 2%  3.84user-s ± 2%    ~     (p=0.236 n=30+30)
    Flate      167user-ms ± 2%  166user-ms ± 4%    ~     (p=0.511 n=28+30)
    GoParser   211user-ms ± 4%  210user-ms ± 3%    ~     (p=0.287 n=29+30)
    Reflect    539user-ms ± 3%  536user-ms ± 2%  -0.59%  (p=0.034 n=29+30)
    Tar        154user-ms ± 3%  155user-ms ± 4%    ~     (p=0.786 n=30+30)
    XML        289user-ms ± 3%  288user-ms ± 4%    ~     (p=0.249 n=30+26)
    
    name       old alloc/op     new alloc/op     delta
    Template       40.7MB ± 0%      40.8MB ± 0%  +0.09%  (p=0.001 n=30+30)
    Unicode        30.8MB ± 0%      30.8MB ± 0%    ~     (p=0.112 n=30+30)
    GoTypes         123MB ± 0%       124MB ± 0%  +0.09%  (p=0.000 n=30+30)
    Compiler        473MB ± 0%       473MB ± 0%  +0.05%  (p=0.000 n=30+30)
    Flate          26.5MB ± 0%      26.5MB ± 0%    ~     (p=0.186 n=29+30)
    GoParser       32.3MB ± 0%      32.4MB ± 0%  +0.07%  (p=0.021 n=28+30)
    Reflect        84.4MB ± 0%      84.6MB ± 0%  +0.21%  (p=0.000 n=30+30)
    Tar            27.3MB ± 0%      27.3MB ± 0%  +0.09%  (p=0.010 n=30+28)
    XML            44.7MB ± 0%      44.7MB ± 0%  +0.07%  (p=0.002 n=30+30)
    
    name       old allocs/op    new allocs/op    delta
    Template         401k ± 1%        400k ± 1%    ~     (p=0.321 n=30+30)
    Unicode          331k ± 1%        331k ± 1%    ~     (p=0.357 n=30+28)
    GoTypes         1.24M ± 0%       1.24M ± 1%  -0.19%  (p=0.001 n=30+30)
    Compiler        4.27M ± 0%       4.27M ± 0%  -0.13%  (p=0.000 n=30+30)
    Flate            252k ± 1%        251k ± 1%  -0.30%  (p=0.005 n=30+30)
    GoParser         325k ± 1%        325k ± 1%    ~     (p=0.224 n=28+30)
    Reflect         1.06M ± 0%       1.05M ± 0%  -0.34%  (p=0.000 n=30+30)
    Tar              266k ± 1%        266k ± 1%    ~     (p=0.333 n=30+30)
    XML              416k ± 1%        415k ± 1%    ~     (p=0.144 n=30+29)
    
    
    Change-Id: I6ba67a9203516373062a2618122306da73333d98
    Reviewed-on: https://go-review.googlesource.com/36211
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ef2593905de3bcdd760bab020667a8aa972a951e
Author: Chris Broadfoot <cbro@golang.org>
Date:   Thu Feb 2 16:43:24 2017 -0800

    readme: add attribution for the Gopher image
    
    Change-Id: I3b1317f0ab46e03d8c5a0af74c83183710a75055
    Reviewed-on: https://go-review.googlesource.com/36214
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 0e3355903d2ebcf5ee9e76096f51ac9a116a9dbb
Author: Russ Cox <rsc@golang.org>
Date:   Thu Feb 2 16:20:58 2017 -0500

    time: record monotonic clock reading in time.Now, for more accurate comparisons
    
    See https://golang.org/design/12914-monotonic for details.
    
    Fixes #12914.
    
    Change-Id: I80edc2e6c012b4ace7161c84cf067d444381a009
    Reviewed-on: https://go-review.googlesource.com/36255
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Caleb Spare <cespare@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8179b9b462eb2946de8488a26dca91a89b3d22e6
Author: Keith Randall <khr@golang.org>
Date:   Mon Jan 30 14:55:12 2017 -0800

    cmd/compile: make sure output params are live if there is a defer
    
    If there is a defer, and that defer recovers, then the caller
    can see all of the output parameters.  That means that we must
    mark all the output parameters live at any point which might panic.
    
    If there is no defer then this is not necessary.  This is implemented.
    
    We could also detect whether there is a recover in any of the defers.
    If not, we would need to mark only output params that the defer
    actually references (and the closure mechanism already does that).
    This is not implemented.
    
    Fixes #18860.
    
    Change-Id: If984fe6686eddce9408bf25e725dd17fc16b8578
    Reviewed-on: https://go-review.googlesource.com/36030
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 12c58bbf81c0feca25292a2291a59e16b5ed00f6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Feb 2 22:38:04 2017 -0800

    cmd/compile: optimize (ZeroExt (Const [c]))
    
    These rules trigger 116 times while running make.bash.
    And at least for the sample code at
    https://github.com/golang/go/issues/18906#issuecomment-277174241
    they are providing optimizations not already present
    in amd64.
    
    Updates #18906
    
    Change-Id: I410a480f566f5ab176fc573fb5ac74f9cffec225
    Reviewed-on: https://go-review.googlesource.com/36217
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 61fb2f6d634aeaf46d3e546267639ad832058d81
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jan 19 14:22:26 2017 -0800

    cmd/compile: speed up hot phi insertion code
    
    This speeds up compilation of the code in #8225 by 25%-30%.
    The complexity of the algorithm is unchanged,
    but this shrinks the constant factor so much that it doesn't matter,
    even the size of the giant type switch gets scaled up dramatically.
    
    name       old time/op      new time/op      delta
    Template        218ms ± 5%       217ms ±10%    ~           (p=0.163 n=27+30)
    Unicode        98.2ms ± 6%      97.7ms ±10%    ~           (p=0.150 n=27+29)
    GoTypes         654ms ± 5%       650ms ± 5%    ~           (p=0.350 n=30+30)
    Compiler        2.70s ± 4%       2.68s ± 3%    ~           (p=0.128 n=30+29)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   276user-ms ± 6%  271user-ms ± 7%  -1.83%        (p=0.003 n=29+28)
    Unicode    138user-ms ± 5%  137user-ms ± 4%    ~           (p=0.071 n=27+27)
    GoTypes    881user-ms ± 4%  877user-ms ± 4%    ~           (p=0.423 n=30+30)
    Compiler   3.76user-s ± 4%  3.72user-s ± 2%  -0.84%        (p=0.028 n=30+29)
    
    name       old alloc/op     new alloc/op     delta
    Template       40.7MB ± 0%      40.7MB ± 0%    ~           (p=0.936 n=30+30)
    Unicode        30.8MB ± 0%      30.8MB ± 0%    ~           (p=0.859 n=28+30)
    GoTypes         123MB ± 0%       123MB ± 0%    ~           (p=0.273 n=30+30)
    Compiler        472MB ± 0%       472MB ± 0%    ~           (p=0.432 n=30+30)
    
    name       old allocs/op    new allocs/op    delta
    Template         401k ± 1%        401k ± 1%    ~           (p=0.859 n=30+30)
    Unicode          331k ± 0%        331k ± 1%    ~           (p=0.823 n=28+30)
    GoTypes         1.24M ± 0%       1.24M ± 0%    ~           (p=0.286 n=30+30)
    Compiler        4.26M ± 0%       4.26M ± 0%    ~           (p=0.359 n=30+30)
    
    Change-Id: Ia850065a9a84c07a5b0b4e23c1758b5679498da7
    Reviewed-on: https://go-review.googlesource.com/36112
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 69e1634985f8d839f74f342925bd7546faea0a19
Author: Keith Randall <khr@golang.org>
Date:   Thu Feb 2 18:50:45 2017 -0800

    runtime: darwin/amd64, don't depend on outarg slots being unmodified
    
    sigtramp was calling sigtrampgo and depending on the fact that
    the 3rd argument slot will not be modified on return.  Our calling
    convention doesn't guarantee that.  Avoid that assumption.
    
    There's no actual bug here, as sigtrampgo does not in fact modify its
    argument slots.  But I found this while working on the dead stack slot
    clobbering tool.  https://go-review.googlesource.com/c/23924/
    
    Change-Id: Ia7e791a2b4c1c74fff24cba8169e7840b4b06ffc
    Reviewed-on: https://go-review.googlesource.com/36216
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 64c5280844aa7d6cbec18c47750f9d3e65f5f72a
Author: Russ Cox <rsc@golang.org>
Date:   Thu Feb 2 20:37:29 2017 -0500

    net/http: fix dns hijacking test
    
    The name lookups are unrooted; the test should be unrooted too.
    Correctly skips the tests if the DNS config specifies a domain
    suffix that has a wildcard entry causing all unrooted names to resolve.
    
    Change-Id: I80470326a5d332f3b8d64663f765fd304c5e0811
    Reviewed-on: https://go-review.googlesource.com/36253
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 57546d67ec8bf66f62bdac58542533c18fe42402
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jan 27 12:01:05 2017 -0800

    cmd/compile: add reusable []Location to ssa.Config
    
    name       old time/op      new time/op      delta
    Template        218ms ± 3%       214ms ± 3%  -1.70%  (p=0.000 n=30+30)
    Unicode         100ms ± 3%       100ms ± 4%    ~     (p=0.614 n=29+30)
    GoTypes         657ms ± 1%       660ms ± 3%  +0.46%  (p=0.046 n=29+30)
    Compiler        2.80s ± 2%       2.80s ± 1%    ~     (p=0.451 n=28+29)
    Flate           131ms ± 2%       132ms ± 4%    ~     (p=1.000 n=29+29)
    GoParser        159ms ± 3%       160ms ± 5%    ~     (p=0.341 n=28+30)
    Reflect         406ms ± 3%       408ms ± 4%    ~     (p=0.511 n=28+30)
    Tar             118ms ± 4%       118ms ± 4%    ~     (p=0.827 n=29+30)
    XML             222ms ± 6%       222ms ± 3%    ~     (p=0.532 n=30+30)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   274user-ms ± 3%  272user-ms ± 3%  -0.87%  (p=0.015 n=29+30)
    Unicode    140user-ms ± 4%  140user-ms ± 3%    ~     (p=0.735 n=29+30)
    GoTypes    890user-ms ± 1%  897user-ms ± 2%  +0.88%  (p=0.002 n=29+30)
    Compiler   3.88user-s ± 2%  3.89user-s ± 1%    ~     (p=0.132 n=30+29)
    Flate      168user-ms ± 2%  157user-ms ± 4%  -6.21%  (p=0.000 n=25+28)
    GoParser   211user-ms ± 2%  213user-ms ± 5%    ~     (p=0.086 n=28+30)
    Reflect    539user-ms ± 2%  541user-ms ± 3%    ~     (p=0.267 n=27+29)
    Tar        156user-ms ± 7%  155user-ms ± 5%    ~     (p=0.708 n=30+30)
    XML        291user-ms ± 5%  294user-ms ± 3%  +0.83%  (p=0.029 n=29+30)
    
    name       old alloc/op     new alloc/op     delta
    Template       40.7MB ± 0%      39.4MB ± 0%  -3.26%  (p=0.000 n=29+26)
    Unicode        30.8MB ± 0%      30.7MB ± 0%  -0.40%  (p=0.000 n=28+30)
    GoTypes         123MB ± 0%       119MB ± 0%  -3.47%  (p=0.000 n=30+29)
    Compiler        472MB ± 0%       455MB ± 0%  -3.60%  (p=0.000 n=30+30)
    Flate          26.5MB ± 0%      25.6MB ± 0%  -3.21%  (p=0.000 n=28+30)
    GoParser       32.3MB ± 0%      31.4MB ± 0%  -2.98%  (p=0.000 n=29+30)
    Reflect        84.4MB ± 0%      82.1MB ± 0%  -2.83%  (p=0.000 n=30+30)
    Tar            27.3MB ± 0%      26.5MB ± 0%  -2.70%  (p=0.000 n=29+29)
    XML            44.6MB ± 0%      43.1MB ± 0%  -3.49%  (p=0.000 n=30+30)
    
    name       old allocs/op    new allocs/op    delta
    Template         401k ± 1%        399k ± 0%  -0.35%  (p=0.000 n=30+28)
    Unicode          331k ± 0%        331k ± 1%    ~     (p=0.907 n=28+30)
    GoTypes         1.24M ± 0%       1.23M ± 0%  -0.43%  (p=0.000 n=30+30)
    Compiler        4.26M ± 0%       4.25M ± 0%  -0.34%  (p=0.000 n=29+30)
    Flate            252k ± 1%        251k ± 1%  -0.41%  (p=0.000 n=30+30)
    GoParser         325k ± 1%        324k ± 1%  -0.31%  (p=0.000 n=27+30)
    Reflect         1.06M ± 0%       1.05M ± 0%  -0.69%  (p=0.000 n=30+30)
    Tar              266k ± 1%        265k ± 1%  -0.51%  (p=0.000 n=29+30)
    XML              416k ± 1%        415k ± 1%  -0.36%  (p=0.002 n=30+30)
    
    Change-Id: I8f784001324df83b2764c44f0e83a540e5beab34
    Reviewed-on: https://go-review.googlesource.com/36212
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f69a6defd1da6509b8f0f54f9ae60e4bf740891d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Feb 2 16:00:26 2017 -0500

    runtime: skip flaky TestGdbPythonCgo on MIPS
    
    It seems the problem is on gdb and the dynamic linker. Skip the
    test for now until we figure out what's going on with the system.
    
    Updates #18784.
    
    Change-Id: Ic9320ffd463f6c231b2c4192652263b1cf7f4231
    Reviewed-on: https://go-review.googlesource.com/36250
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ecbf90840471a8cd5500c90baea7a60a65b14b10
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Thu Feb 2 11:39:02 2017 -0800

    doc: consistently link to the SettingGOPATH page
    
    Change-Id: I4fdd81aa7c9b180cb72ec4af3e7d9d803c99ecac
    Reviewed-on: https://go-review.googlesource.com/36033
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6ad2d6aa922368224730a347d6a82387d4770c40
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Jan 20 12:14:22 2017 -0500

    cmd/compile: simplify IsNonNil ConstNil
    
    Change-Id: I9ed5a2065cef06708e319b16c801da2eff42004e
    Reviewed-on: https://go-review.googlesource.com/35497
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit fddc004537b90e46c894e9b22c291759fcd1209e
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Nov 21 11:31:39 2016 -0500

    cmd/compile: remove nil check for Zero/Move on 386, AMD64, S390X
    
    Fixes #18003.
    
    Change-Id: Iadcc5c424c64badecfb5fdbd4dbd9197df56182c
    Reviewed-on: https://go-review.googlesource.com/33421
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit c1363b2d91d9aa152ef17a68d7a1778426b33727
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Feb 1 15:17:32 2017 -0800

    cmd/compile: provide line number for cgo directive error (fix a TODO)
    
    Also: Remove double "go:" prefix in related error message.
    
    Fixes #18882.
    
    Change-Id: Ifbbd8e2f7529b43f035d3dbf7ca4a91f212bc6b6
    Reviewed-on: https://go-review.googlesource.com/36121
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e546b295b8a3b6f8996a02aeecfbdfedd07c21aa
Author: Lars Wiegman <lars@namsral.com>
Date:   Tue Jan 17 11:38:18 2017 +0100

    runtime: use mach_absolute_time for runtime.nanotime
    
    The existing darwin/amd64 implementation of runtime.nanotime returns the
    wallclock time, which results in timers not functioning properly when
    system time runs backwards. By implementing the algorithm used by the
    darwin syscall mach_absolute_time, timers will function as expected.
    
    The algorithm is described at
    https://opensource.apple.com/source/xnu/xnu-3248.60.10/libsyscall/wrappers/mach_absolute_time.s
    
    Fixes #17610
    
    Change-Id: I9c8d35240d48249a6837dca1111b1406e2686f67
    Reviewed-on: https://go-review.googlesource.com/35292
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8958d8ce37fdae15693787019869ab21e7b8347a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Jan 22 10:17:12 2017 -0800

    cmd/compile: skip convT2E for empty structs
    
    Fixes #18402
    
    Change-Id: I5af800857fb2e365ce4224eece9171277106ec7d
    Reviewed-on: https://go-review.googlesource.com/35562
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3d5eb4a6bef97c16308442a0e7b87fcdf5fd0f02
Author: Keith Randall <khr@golang.org>
Date:   Tue Jan 24 12:48:10 2017 -0800

    cmd/compile: better implementation of Slicemask
    
    Use (-x)>>63 instead of ((x-1)>>63)^-1 to get a mask that
    is 0 when x is 0 and all ones when x is positive.
    
    Saves one instruction when slicing.
    
    Change-Id: Ib46d53d3aac6530ac481fa2f265a6eadf3df0567
    Reviewed-on: https://go-review.googlesource.com/35641
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 03583675765933d4a5fb394cfa89fb41b274aaa7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jan 21 19:52:09 2017 -0800

    cmd/compile, runtime: convert byte-sized values to interfaces without allocation
    
    Based in part on khr's CL 2500.
    
    Updates #17725
    Updates #18121
    
    Change-Id: I744e1f92fc2104e6c5bd883a898c30b2eea8cc31
    Reviewed-on: https://go-review.googlesource.com/35555
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit c682d3239e5aa05a77ad21f2267efc4e2e60c05f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jan 21 13:41:06 2017 -0800

    cmd/compile: convert constants to interfaces without allocating
    
    The order pass is responsible for ensuring that
    values passed to runtime functions, including
    convT2E/convT2I, are addressable.
    
    Prior to this CL, this was always accomplished
    by creating a temp, which frequently escaped to
    the heap, causing allocations, perhaps most
    notably in code like:
    
    fmt.Println(1, 2, 3) // allocates three times
    
    None of the runtime routines modify the contents
    of the pointers they receive, so in the case of
    constants, instead of creating a temp value,
    we can create a static value.
    
    (Marking the static value as read-only provides
    protection against accidental attempts by the runtime
    to modify the constant data.)
    
    This improves code generation for code like:
    
    panic("abc")
    c <- 2 // c is a chan int
    
    which can now simply refer to "abc" and 2,
    rather than going by way of a temporary.
    
    It also allows us to optimize convT2E/convT2I,
    by recognizing static readonly values
    and directly constructing the interface.
    
    This CL adds ~0.5% to binary size, despite
    decreasing the size of many functions,
    because it also adds many static symbols.
    
    This binary size regression could be recovered in
    future (but currently unplanned) work.
    
    There is a lot of content-duplication in these
    symbols; this statement generates six new symbols,
    three containing an int 1 and three containing
    a pointer to the string "a":
    
    fmt.Println(1, 1, 1, "a", "a", "a")
    
    These symbols could be made content-addressable.
    
    Furthermore, these symbols are small, so the
    alignment and naming overhead is large.
    As with the go.strings section, these symbols
    could be hidden and have their alignment reduced.
    
    The changes to test/live.go make it impossible
    (at least with current optimization techniques)
    to place the values being passed to the runtime
    in static symbols, preserving autotmp creation.
    
    Fixes #18704
    
    Benchmarks from fmt and go-kit's logging package:
    
    github.com/go-kit/kit/log
    
    name                      old time/op    new time/op    delta
    JSONLoggerSimple-8          1.91µs ± 2%    2.11µs ±22%     ~     (p=1.000 n=9+10)
    JSONLoggerContextual-8      2.60µs ± 6%    2.43µs ± 2%   -6.29%  (p=0.000 n=9+10)
    Discard-8                    101ns ± 2%      34ns ±14%  -66.33%  (p=0.000 n=10+9)
    OneWith-8                    161ns ± 1%     102ns ±16%  -36.78%  (p=0.000 n=10+10)
    TwoWith-8                    175ns ± 3%     106ns ± 7%  -39.36%  (p=0.000 n=10+9)
    TenWith-8                    293ns ± 3%     227ns ±15%  -22.44%  (p=0.000 n=9+10)
    LogfmtLoggerSimple-8         704ns ± 2%     608ns ± 2%  -13.65%  (p=0.000 n=10+9)
    LogfmtLoggerContextual-8     962ns ± 1%     860ns ±17%  -10.57%  (p=0.003 n=9+10)
    NopLoggerSimple-8            188ns ± 1%     120ns ± 1%  -36.39%  (p=0.000 n=9+10)
    NopLoggerContextual-8        379ns ± 1%     243ns ± 0%  -35.77%  (p=0.000 n=9+10)
    ValueBindingTimestamp-8      577ns ± 1%     499ns ± 1%  -13.51%  (p=0.000 n=10+10)
    ValueBindingCaller-8         898ns ± 2%     844ns ± 2%   -6.00%  (p=0.000 n=10+10)
    
    name                      old alloc/op   new alloc/op   delta
    JSONLoggerSimple-8            904B ± 0%      872B ± 0%   -3.54%  (p=0.000 n=10+10)
    JSONLoggerContextual-8      1.20kB ± 0%    1.14kB ± 0%   -5.33%  (p=0.000 n=10+10)
    Discard-8                    64.0B ± 0%     32.0B ± 0%  -50.00%  (p=0.000 n=10+10)
    OneWith-8                    96.0B ± 0%     64.0B ± 0%  -33.33%  (p=0.000 n=10+10)
    TwoWith-8                     160B ± 0%      128B ± 0%  -20.00%  (p=0.000 n=10+10)
    TenWith-8                     672B ± 0%      640B ± 0%   -4.76%  (p=0.000 n=10+10)
    LogfmtLoggerSimple-8          128B ± 0%       96B ± 0%  -25.00%  (p=0.000 n=10+10)
    LogfmtLoggerContextual-8      304B ± 0%      240B ± 0%  -21.05%  (p=0.000 n=10+10)
    NopLoggerSimple-8             128B ± 0%       96B ± 0%  -25.00%  (p=0.000 n=10+10)
    NopLoggerContextual-8         304B ± 0%      240B ± 0%  -21.05%  (p=0.000 n=10+10)
    ValueBindingTimestamp-8       159B ± 0%      127B ± 0%  -20.13%  (p=0.000 n=10+10)
    ValueBindingCaller-8          112B ± 0%       80B ± 0%  -28.57%  (p=0.000 n=10+10)
    
    name                      old allocs/op  new allocs/op  delta
    JSONLoggerSimple-8            19.0 ± 0%      17.0 ± 0%  -10.53%  (p=0.000 n=10+10)
    JSONLoggerContextual-8        25.0 ± 0%      21.0 ± 0%  -16.00%  (p=0.000 n=10+10)
    Discard-8                     3.00 ± 0%      1.00 ± 0%  -66.67%  (p=0.000 n=10+10)
    OneWith-8                     3.00 ± 0%      1.00 ± 0%  -66.67%  (p=0.000 n=10+10)
    TwoWith-8                     3.00 ± 0%      1.00 ± 0%  -66.67%  (p=0.000 n=10+10)
    TenWith-8                     3.00 ± 0%      1.00 ± 0%  -66.67%  (p=0.000 n=10+10)
    LogfmtLoggerSimple-8          4.00 ± 0%      2.00 ± 0%  -50.00%  (p=0.000 n=10+10)
    LogfmtLoggerContextual-8      7.00 ± 0%      3.00 ± 0%  -57.14%  (p=0.000 n=10+10)
    NopLoggerSimple-8             4.00 ± 0%      2.00 ± 0%  -50.00%  (p=0.000 n=10+10)
    NopLoggerContextual-8         7.00 ± 0%      3.00 ± 0%  -57.14%  (p=0.000 n=10+10)
    ValueBindingTimestamp-8       5.00 ± 0%      3.00 ± 0%  -40.00%  (p=0.000 n=10+10)
    ValueBindingCaller-8          4.00 ± 0%      2.00 ± 0%  -50.00%  (p=0.000 n=10+10)
    
    fmt
    
    name                             old time/op    new time/op    delta
    SprintfPadding-8                   88.9ns ± 3%    79.1ns ± 1%   -11.09%  (p=0.000 n=10+7)
    SprintfEmpty-8                     12.6ns ± 3%    12.8ns ± 3%      ~     (p=0.136 n=10+10)
    SprintfString-8                    38.7ns ± 5%    26.9ns ± 6%   -30.65%  (p=0.000 n=10+10)
    SprintfTruncateString-8            56.7ns ± 2%    47.0ns ± 3%   -17.05%  (p=0.000 n=10+10)
    SprintfQuoteString-8                164ns ± 2%     153ns ± 2%    -7.01%  (p=0.000 n=10+10)
    SprintfInt-8                       38.9ns ±15%    26.5ns ± 2%   -31.93%  (p=0.000 n=10+9)
    SprintfIntInt-8                    60.3ns ± 9%    38.2ns ± 1%   -36.67%  (p=0.000 n=10+8)
    SprintfPrefixedInt-8               58.6ns ±13%    51.2ns ±11%   -12.66%  (p=0.001 n=10+10)
    SprintfFloat-8                     71.4ns ± 3%    64.2ns ± 3%   -10.08%  (p=0.000 n=8+10)
    SprintfComplex-8                    175ns ± 3%     159ns ± 2%    -9.03%  (p=0.000 n=10+10)
    SprintfBoolean-8                   33.5ns ± 4%    25.7ns ± 5%   -23.28%  (p=0.000 n=10+10)
    SprintfHexString-8                 65.3ns ± 3%    51.7ns ± 5%   -20.86%  (p=0.000 n=10+9)
    SprintfHexBytes-8                  67.2ns ± 5%    67.9ns ± 4%      ~     (p=0.383 n=10+10)
    SprintfBytes-8                      129ns ± 7%     124ns ± 7%      ~     (p=0.074 n=9+10)
    SprintfStringer-8                   127ns ± 4%     126ns ± 8%      ~     (p=0.506 n=9+10)
    SprintfStructure-8                  357ns ± 3%     359ns ± 3%      ~     (p=0.469 n=10+10)
    ManyArgs-8                          203ns ± 6%     126ns ± 3%   -37.94%  (p=0.000 n=10+10)
    FprintInt-8                         119ns ±10%      74ns ± 3%   -37.54%  (p=0.000 n=10+10)
    FprintfBytes-8                      122ns ± 4%     120ns ± 3%      ~     (p=0.124 n=10+10)
    FprintIntNoAlloc-8                 78.2ns ± 5%    74.1ns ± 3%    -5.28%  (p=0.000 n=10+10)
    ScanInts-8                          349µs ± 1%     349µs ± 0%      ~     (p=0.606 n=9+8)
    ScanRecursiveInt-8                 43.8ms ± 7%    40.1ms ± 2%    -8.42%  (p=0.000 n=10+10)
    ScanRecursiveIntReaderWrapper-8    43.5ms ± 4%    40.4ms ± 2%    -7.16%  (p=0.000 n=10+9)
    
    name                             old alloc/op   new alloc/op   delta
    SprintfPadding-8                    24.0B ± 0%     16.0B ± 0%   -33.33%  (p=0.000 n=10+10)
    SprintfEmpty-8                      0.00B          0.00B           ~     (all equal)
    SprintfString-8                     21.0B ± 0%      5.0B ± 0%   -76.19%  (p=0.000 n=10+10)
    SprintfTruncateString-8             32.0B ± 0%     16.0B ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfQuoteString-8                48.0B ± 0%     32.0B ± 0%   -33.33%  (p=0.000 n=10+10)
    SprintfInt-8                        16.0B ± 0%      1.0B ± 0%   -93.75%  (p=0.000 n=10+10)
    SprintfIntInt-8                     24.0B ± 0%      3.0B ± 0%   -87.50%  (p=0.000 n=10+10)
    SprintfPrefixedInt-8                72.0B ± 0%     64.0B ± 0%   -11.11%  (p=0.000 n=10+10)
    SprintfFloat-8                      16.0B ± 0%      8.0B ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfComplex-8                    48.0B ± 0%     32.0B ± 0%   -33.33%  (p=0.000 n=10+10)
    SprintfBoolean-8                    8.00B ± 0%     4.00B ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfHexString-8                  96.0B ± 0%     80.0B ± 0%   -16.67%  (p=0.000 n=10+10)
    SprintfHexBytes-8                    112B ± 0%      112B ± 0%      ~     (all equal)
    SprintfBytes-8                      96.0B ± 0%     96.0B ± 0%      ~     (all equal)
    SprintfStringer-8                   32.0B ± 0%     32.0B ± 0%      ~     (all equal)
    SprintfStructure-8                   256B ± 0%      256B ± 0%      ~     (all equal)
    ManyArgs-8                          80.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
    FprintInt-8                         8.00B ± 0%     0.00B       -100.00%  (p=0.000 n=10+10)
    FprintfBytes-8                      32.0B ± 0%     32.0B ± 0%      ~     (all equal)
    FprintIntNoAlloc-8                  0.00B          0.00B           ~     (all equal)
    ScanInts-8                         15.2kB ± 0%    15.2kB ± 0%      ~     (p=0.248 n=9+10)
    ScanRecursiveInt-8                 21.6kB ± 0%    21.6kB ± 0%      ~     (all equal)
    ScanRecursiveIntReaderWrapper-8    21.7kB ± 0%    21.7kB ± 0%      ~     (all equal)
    
    name                             old allocs/op  new allocs/op  delta
    SprintfPadding-8                     2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfEmpty-8                       0.00           0.00           ~     (all equal)
    SprintfString-8                      2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfTruncateString-8              2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfQuoteString-8                 2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfInt-8                         2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfIntInt-8                      3.00 ± 0%      1.00 ± 0%   -66.67%  (p=0.000 n=10+10)
    SprintfPrefixedInt-8                 2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfFloat-8                       2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfComplex-8                     2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfBoolean-8                     2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfHexString-8                   2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    SprintfHexBytes-8                    2.00 ± 0%      2.00 ± 0%      ~     (all equal)
    SprintfBytes-8                       2.00 ± 0%      2.00 ± 0%      ~     (all equal)
    SprintfStringer-8                    4.00 ± 0%      4.00 ± 0%      ~     (all equal)
    SprintfStructure-8                   7.00 ± 0%      7.00 ± 0%      ~     (all equal)
    ManyArgs-8                           8.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
    FprintInt-8                          1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
    FprintfBytes-8                       1.00 ± 0%      1.00 ± 0%      ~     (all equal)
    FprintIntNoAlloc-8                   0.00           0.00           ~     (all equal)
    ScanInts-8                          1.60k ± 0%     1.60k ± 0%      ~     (all equal)
    ScanRecursiveInt-8                  1.71k ± 0%     1.71k ± 0%      ~     (all equal)
    ScanRecursiveIntReaderWrapper-8     1.71k ± 0%     1.71k ± 0%      ~     (all equal)
    
    Change-Id: I7ba72a25fea4140a0ba40a9f443103ed87cc69b5
    Reviewed-on: https://go-review.googlesource.com/35554
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit f395e878887df1a28966353f32f5ee70973c8f3f
Author: David Lazar <lazard@golang.org>
Date:   Fri Jan 27 15:24:48 2017 -0500

    io: fix test when MultiReader is inlined with -l=3
    
    This ensures there isn't a live reference to buf1 on our stack
    when MultiReader is inlined.
    
    Fixes #18819.
    
    Change-Id: I96a8cdc1ffad8f8a10c0ddcbf0299005f3176b61
    Reviewed-on: https://go-review.googlesource.com/35931
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 16e430e1ef8b9c9259bb4f07a0787de4310b3041
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jan 11 13:58:20 2017 -0800

    cmd/compile: reduce slice growth in fuseBlockPlain
    
    Instead of always appending to c.Values,
    choose whichever slice is larger;
    b.Values will be set to nil anyway.
    
    Appending once instead of in a loop also
    limits slice growth to once per function call
    and is more efficient.
    
    Reduces max rss for the program in #18602 by 6.5%,
    and eliminates fuseBlockPlain from the alloc_space
    pprof output. fuseBlockPlain previously accounted
    for 16.74% of allocated memory.
    
    Updates #18602.
    
    Change-Id: I417b03722d011a59a679157da43dc91f4425210e
    Reviewed-on: https://go-review.googlesource.com/35114
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 01c8719f8b6cc689c6b446ed01ce570bce7d6287
Author: Keith Randall <khr@golang.org>
Date:   Thu Dec 8 16:17:20 2016 -0800

    cmd/compile: move rotate instruction generation to SSA
    
    Remove rotate generation from walk.  Remove OLROT and ssa.Lrot* opcodes.
    Generate rotates during SSA lowering for architectures that have them.
    
    This CL will allow rotates to be generated in more situations,
    like when the shift values are determined to be constant
    only after some analysis.
    
    Fixes #18254
    
    Change-Id: I8d6d684ff5ce2511aceaddfda98b908007851079
    Reviewed-on: https://go-review.googlesource.com/34232
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 8a9dc05312ece25aae057e618f3715208d201f8e
Author: Keith Randall <khr@golang.org>
Date:   Fri Dec 9 16:59:38 2016 -0800

    cmd/compile: allow inlining of functions with intrinsics in them
    
    Intrinsics are ok to inline as they don't rewrite to actual calls.
    
    Change-Id: Ieb19c834c61579823c62c6d1a1b425d6c4d4de23
    Reviewed-on: https://go-review.googlesource.com/34272
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 73f92f9b0405e98427bbb445f24cffb5d3c4d01b
Author: Keith Randall <khr@golang.org>
Date:   Sun Nov 27 11:43:08 2016 -0800

    cmd/compile: use len(s)<=cap(s) to remove more bounds checks
    
    When we discover a relation x <= len(s), also discover the relation
    x <= cap(s).  That way, in situations like:
    
    a := s[x:]  // tests 0 <= x <= len(s)
    b := s[:x]  // tests 0 <= x <= cap(s)
    
    the second check can be eliminated.
    
    Fixes #16813
    
    Change-Id: Ifc037920b6955e43bac1a1eaf6bac63a89cfbd44
    Reviewed-on: https://go-review.googlesource.com/33633
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alexandru Moșoi <alexandru@mosoi.ro>
    Reviewed-by: David Chase <drchase@google.com>

commit 6317f92f6e51f679712deec6094c6b5fc2948a5b
Author: Keith Randall <khr@golang.org>
Date:   Sun Nov 27 10:41:37 2016 -0800

    cmd/compile: fix CSE with commutative ops
    
    CSE opportunities were being missed for commutative ops. We used to
    order the args of commutative ops (by arg ID) once at the start of CSE.
    But that may not be enough.
    
    i1 = (Load ptr mem)
    i2 = (Load ptr mem)
    x1 = (Add i1 j)
    x2 = (Add i2 j)
    
    Equivalent commutative ops x1 and x2 may not get their args ordered in
    the same way because because at the start of CSE, we don't know that
    the i values will be CSEd. If x1 and x2 get opposite orders we won't
    CSE them.
    
    Instead, (re)order the args of commutative operations by their
    equivalence class IDs each time we partition an equivalence class.
    
    Change-Id: Ic609fa83b85299782a5e85bf93dc6023fccf4b0c
    Reviewed-on: https://go-review.googlesource.com/33632
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Todd Neal <todd@tneal.org>

commit 34b563f447280f9d386f208646ac4f94cafc4ab6
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sat Jan 14 05:23:23 2017 -0700

    cmd/compile: improve error for wrong type in switch
    
    Fixes #10561.
    
    Provides a better diagnostic message for failed type switch
    satisfaction in the case that a value receiver is being used
    in place of the pointer receiver that implements and satisfies
    the interface.
    
    Change-Id: If8c13ba13f2a8d81bf44bac7c3a66c12921ba921
    Reviewed-on: https://go-review.googlesource.com/35235
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ba1a65fc518c367bd4a3e18324036d457e6a07c3
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Thu Feb 2 21:53:52 2017 +0900

    cmd/cgo: don't track same node twice in guessKinds
    
    Change-Id: Ib2c1490a42e3485913a05a0b2fecdcc425d42871
    Reviewed-on: https://go-review.googlesource.com/36083
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 329fff0db01dd1252444cf106dffe72ba20ddbc8
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sat Nov 26 15:27:29 2016 +0000

    misc/cgo/testshared: remove unused flag.Parse()
    
    TestMain doesn't make use of any flags.
    
    Change-Id: I98ec582fb004045a5067618f605ccfeb1f9f4bbb
    Reviewed-on: https://go-review.googlesource.com/33613
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 312ea5bf8ff76d8ea4ea0b92df29de5204446b46
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Feb 1 18:41:27 2017 +0100

    misc/ios: allow exit code 0 to mean test success
    
    Tests that use TestMain might never call m.Run(), and simply return
    from TestMain. In that case, the iOS test harness never sees the
    PASS from the testing framework and assumes the test failed.
    
    Allow an exit with exit code 0 to also mean test success, thereby
    fixing the objdump test on iOS.
    
    Change-Id: I1fe9077b05931aa0905e41b88945cd153c5b35b6
    Reviewed-on: https://go-review.googlesource.com/36065
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cb6e0639fb090ea0e129b1ddb956a7e645cff285
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Jan 1 03:08:48 2017 -0700

    cmd/compile: improve error message if init is directly invoked
    
    Fixes #8481.
    
    Inform the user that init functions cannot be directly invoked
    in user code, as mandated by the spec at:
    http://golang.org/ref/spec#Program_initialization_and_execution.
    
    Change-Id: Ib12c0c08718ffd48b76b6f9b13c76bb6612d2e7b
    Reviewed-on: https://go-review.googlesource.com/34790
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit dbd51ce99c140766808c17b334b8795b8040c0b3
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Wed Nov 2 17:20:22 2016 +0100

    cmd/compile: intrinsify math.Sqrt by using only the ssa backend
    
    Change-Id: If3cb64f52fe0fd7331b6f1acf3d15dd705dfd633
    Reviewed-on: https://go-review.googlesource.com/32591
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit dba0d38298f0af466274795ce35ad5f310b391db
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sat Jul 30 01:45:27 2016 -0700

    cmd/compile: fix conversion error message for printed slices
    
    Fixes #15055.
    
    Updates exprfmt printing using fmt verb "%v" to check that n.Left
    is non-nil before attempting to print it, otherwise we'll print
    the nodes in the list using verb "%.v".
    
    Credit to @mdempsky for this approach and for finding
    the root cause of the issue.
    
    Change-Id: I20a6464e916dc70d5565e145164bb9553e5d3865
    Reviewed-on: https://go-review.googlesource.com/25361
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b443babad49f90e9507d91819736c97a7495e308
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jan 21 16:58:01 2017 -0800

    test: add extra test for issue18661
    
    Make sure that the lack of an lvalue doesn't
    cause extra side-effects.
    
    Updates #18661
    Updates #18739
    
    Change-Id: I52eb4b4a5c6f8ff5cddd2115455f853c18112c19
    Reviewed-on: https://go-review.googlesource.com/36126
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit edad59cfae70d2bfb9cdf66e2492f9a1c1318ddc
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Feb 1 16:40:46 2017 -0800

    cmd/compile: skip reexporting types in reexportdep
    
    The binary export format embeds type definitions inline as necessary,
    so there's no need to add them to exportlist. Also, constants are
    embedded directly by value, so they can be omitted too.
    
    Change-Id: Id1879eb97c298a5a52f615cf9883c346c7f7bd69
    Reviewed-on: https://go-review.googlesource.com/36170
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 16dd0624c204ced87ea950b129c5c26d82e2aad4
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Feb 1 15:13:48 2017 -0800

    cmd/compile/internal/gc: add comment and test for #15550
    
    When switching to the new parser, I changed cmd/compile to handle iota
    per an intuitive interpretation of how nested constant declarations
    should work (which also matches go/types).
    
    Note: if we end up deciding that the current spec wording is
    intentional (i.e., confirming gccgo's current behavior), the test will
    need to be updated to expect 4 instead of 1.
    
    Updates #15550.
    
    Change-Id: I441f5f13209f172b73ef75031f2a9daa5e985277
    Reviewed-on: https://go-review.googlesource.com/36122
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 99e1e4f657c24769a2b42a4aa26c226b6e1db915
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Tue Jan 17 15:06:12 2017 +1100

    cmd/link: assume that runtime.epclntab lives in .text section
    
    Sometimes STEXT symbols point to the first byte of .data
    section, instead of the end of .text section. But, while writing
    pe symbol table, we should treat them as if they belong to the
    .text section. Change pe symbol table records for these symbols.
    
    Fixes #14710
    
    Change-Id: I1356e61aa8fa37d590d7b1677b2bac214ad0ba4e
    Reviewed-on: https://go-review.googlesource.com/35272
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 08a3a7c08a04a0041db3ee6923d9dccb8aaf764d
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Feb 1 15:12:51 2017 -0800

    cmd/compile: update maxPtrmaskBytes comment for larger value
    
    The comment for maxPtrmaskBytes implied that the value was still 16,
    but that changed in CL 10815.
    
    Change-Id: I86e304bc7d9d1a0a6b22b600fefcc1325e4372d9
    Reviewed-on: https://go-review.googlesource.com/36120
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 4aa7b142681dba2a1d76e59ecb5dc6923a192eb0
Author: Kale Blankenship <kale@lemnisys.com>
Date:   Sat Dec 17 20:40:29 2016 -0800

    net/http: detect Comcast et al DNS and auto-skip DNS tests
    
    Adds helper function to auto-skip tests when DNS returns
    a successful response for a domain known not to exist.
    
    The error from `net.LookupHost` is intentionally ignored
    because the DNS tests will fail anyway if there are issues
    unrelated to NXDOMAIN responses.
    
    Fixes #17884
    
    Change-Id: I729391bd702218507561818668f791331295299e
    Reviewed-on: https://go-review.googlesource.com/34516
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a27b78141b85d3b9733647de3f3863977d2f9f81
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Feb 1 12:35:53 2017 -0800

    cmd/compile/internal/gc: inline typedcl0 and typedcl1
    
    It's easier to understand what's happening after inlining these into
    noder.typeDecl.
    
    Change-Id: I7beed5a1e18047bf09f2d4ddf64b9646c324d8d6
    Reviewed-on: https://go-review.googlesource.com/36111
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit bb41b4d599f5758e25091666e123c41b401ac890
Author: Michael Fraenkel <michael.fraenkel@gmail.com>
Date:   Thu Dec 15 09:58:30 2016 -0500

    net/http: make Server validate HTTP method
    
    Fixes #18319
    
    Change-Id: If88e60a86828f60d8d93fc291932c19bab19e8dc
    Reviewed-on: https://go-review.googlesource.com/34470
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 90de5e817c257ffea8dbba12a9f012e22b8c0c7c
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Feb 1 12:28:47 2017 +0100

    misc/ios: use the default go test timeout
    
    If -test.timeout is not specified to go test, it will time out after
    a default 10 minutes.
    
    The iOS exec wrapper also contains a fail safe timeout mechanism for
    a stuck device. However, if no explicit -test.timeout is specified,
    it will use a timeout of 0, plus some constant amount.
    
    Use the same default timeout in the exec wrapper as for go test,
    10 minutes.
    
    Change-Id: I6465ccd9f7b9ce08fa302e6697f7938a0ea9af34
    Reviewed-on: https://go-review.googlesource.com/36062
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7ef746961cec772635b83553fec6836956387954
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Feb 1 12:11:54 2017 +0100

    syscall: regenerate zsyscall_darwin_arm64.go with mksyscall.pl
    
    Notably, this change fixes the TestTCPReadWriteAllocs test because
    the errnoErr wrapper is now used, elimitating the allocation for
    common errnos.
    
    The change to Dup is caused by a CL 8095 that changed the Dup* calls
    to use Syscall instead of RawSyscall.
    
    Found while working on the new iOS builders.
    
    Change-Id: I44ab9dcad27db190e175aa149865b33944f48674
    Reviewed-on: https://go-review.googlesource.com/36061
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7d889af26d40a6d81e668c9780086e8c8ef14ceb
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Feb 1 09:37:47 2017 +0100

    misc/ios: include the bundle id in the GOIOS_APP_ID env variable
    
    The iOS exec wrapper use the constant bundle id "golang.gotest" for
    running Go programs on iOS. However, that only happens to work on
    the old iOS builders where their provisioning profile covers
    that bundle id.
    
    Expand the detection script to list all available provisioning
    profiles for the attached device and include the bundle id in the
    GOIOS_APP_ID environment variable.
    
    To allow the old builders to continue, the "golang.gotest" bundle
    id is used as a fallback if only the app id prefix is specified in
    GOIOS_APP_ID.
    
    For the new builders.
    
    Change-Id: I8baa1d4d57f845de851c3fad3f178e05e9a01b17
    Reviewed-on: https://go-review.googlesource.com/36060
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ade6bcf1d5d0fdddb7ce779e7aa3f1479e1e77a3
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Feb 1 08:57:46 2017 +0100

    misc/ios: ignore stderr from iOS tools
    
    On (at least) macOS 10.12, the `security cms` subcommand used by the
    iOS detection script will output an error to stderr. The command
    otherwise succeeds, but the extra line confuses a later parsing step.
    
    To fix it, use only stdout and ignore stderr from every command run
    by detect.go.
    
    For the new iOS builders.
    
    Change-Id: Iee426da7926d7f987ba1be061fa92ebb853ef53d
    Reviewed-on: https://go-review.googlesource.com/36059
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3f45916433c7e868c75b7a23c9288f8c67447acc
Author: Adam Langley <agl@golang.org>
Date:   Mon Dec 5 10:24:30 2016 -0800

    crypto/tls: reject SNI values with a trailing dot.
    
    SNI values may not include a trailing dot according to
    https://tools.ietf.org/html/rfc6066#section-3. Although crypto/tls
    handled this correctly as a client, it didn't reject this as a server.
    
    This change makes sending an SNI value with a trailing dot a fatal
    error.
    
    Updates #18114.
    
    Change-Id: Ib7897ab40e98d4a7a4646ff8469a55233621f631
    Reviewed-on: https://go-review.googlesource.com/33904
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e56be943753d454b3eeb938c30de812266a8549e
Author: Adam Langley <agl@golang.org>
Date:   Wed Dec 14 14:10:26 2016 -0800

    crypto/x509: add test for v1 intermediates.
    
    X.509v1 certificates are ancient and should be dead. (They are even
    prohibited by the Baseline requirements, section 7.1.1.)
    
    However, there are a number of v1 roots from the 1990's that are still
    in operation. Thus crypto/x509.Certificate.CheckSignatureFrom allows
    X.509v1 certificates to sign other certificates.
    
    The chain building code, however, only allows v1 certificates to sign
    others if they're a root. This change adds a test to check that.
    
    Change-Id: Ib8d81e522f30d41932b89bdf3b19ef3782d8ec12
    Reviewed-on: https://go-review.googlesource.com/34383
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c67f0d949941279681b53b585eb967326811a93b
Author: Anmol Sethi <anmol@aubble.com>
Date:   Sun Jan 29 03:18:17 2017 -0500

    crypto/tls: document ConnectionState.NegotiatedProtocol more clearly
    
    ConnectionState.NegotiatedProtocol's documentation implies that it will
    always be from Config.NextProtos. This commit clarifies that there is no
    guarantee.
    
    This commit also adds a note to
    ConnectionState.NegotiatedProtocolIsMutual, making it clear that it is
    client side only.
    
    Fixes #18841
    
    Change-Id: Icd028af8042f31e45575f1080c5e9bd3012e03d7
    Reviewed-on: https://go-review.googlesource.com/35917
    Reviewed-by: Filippo Valsorda <hi@filippo.io>
    Reviewed-by: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b48c419c4109a1c0e1e7a60a7a28659f6a92e827
Author: Blake Mizerany <blake.mizerany@gmail.com>
Date:   Mon Jan 23 11:03:03 2017 -0800

    net/http/httputil: eliminate duplicate alloc/copy in ReverseProxy
    
    This commit elimates the request allocation and shallow copy duplication
    already done by req.WithContext.
    
    name         old time/op    new time/op    delta
    ServeHTTP-4     216µs ±36%     212µs ±15%     ~     (p=0.853 n=10+10)
    
    name         old alloc/op   new alloc/op   delta
    ServeHTTP-4     917kB ±36%    1137kB ± 0%     ~     (p=0.352 n=10+10)
    
    name         old allocs/op  new allocs/op  delta
    ServeHTTP-4      5.00 ± 0%      4.00 ± 0%  -20.00%  (p=0.000 n=10+10)
    
    Change-Id: I514a59c30b037c7a65c355b06fd82c2d6ff17bb0
    Reviewed-on: https://go-review.googlesource.com/35569
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 435450bf3c6efcc65111e96a42fc1c8acd3081e3
Author: Thomas Bonfort <thomas.bonfort@gmail.com>
Date:   Sun Jan 1 17:02:44 2017 +0100

    image/jpeg: improve performance when encoding *image.YCbCr
    
    The existing implementation falls back to using image.At()
    for each pixel when encoding an *image.YCbCr which is
    inefficient and causes many memory allocations.
    
    This change makes the jpeg encoder directly read Y, Cb, and Cr
    pixel values.
    
    benchmark                  old ns/op     new ns/op     delta
    BenchmarkEncodeYCbCr-4     43990846      24201148      -44.99%
    
    benchmark                  old MB/s     new MB/s     speedup
    BenchmarkEncodeYCbCr-4     20.95        38.08        1.82x
    
    Fixes #18487
    
    Change-Id: Iaf2ebc646997e3e1fffa5335f1b0d642e15bd453
    Reviewed-on: https://go-review.googlesource.com/34773
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Nigel Tao <nigeltao@golang.org>

commit de479267ef9db0911dac68e94d75186313bae11d
Author: Filippo Valsorda <hi@filippo.io>
Date:   Wed Feb 1 19:57:11 2017 +0000

    doc: mention SHA-256 CBC suites are off by default
    
    Change-Id: I82c41bd1d82adda457ddb5dd08caf0647905da22
    Reviewed-on: https://go-review.googlesource.com/36091
    Reviewed-by: Matt Layher <mdlayher@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 99fb0dc1b17b5f203334b92d8af4761aff51b5c1
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Feb 1 21:23:03 2017 +0000

    Revert "testing: delete unused stopAlarm()"
    
    This reverts commit ed8c62b7fb47b717dc2adc9f6d0c90c924c67712.
    
    Turns out it was needed in later commits.
    
    Change-Id: I07a7bc2429976d8a5a89f915a11625c118b85500
    Reviewed-on: https://go-review.googlesource.com/36113
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 871300308a6c14ccaaed16e201752ee53b3b4037
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Wed Dec 21 11:07:11 2016 +0100

    cmd/compile: never report "truncated to real" for toint calls
    
    Whoever called toint() is expecting the {Mpint, Mpflt, Mpcplx} arg to
    be converted to an integer expression, so it never makes sense to
    report an error as "constant X truncated to real".
    
    Fixes #11580
    
    Change-Id: Iadcb105f0802358a7f77188c2b1e63fe80c5580c
    Reviewed-on: https://go-review.googlesource.com/34638
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>

commit 8b11e0b42d6466726639ac702c073047d9fadedb
Author: Anmol Sethi <anmol@aubble.com>
Date:   Sun Jan 8 14:46:54 2017 -0500

    net/http: remove check for null bytes in Dir.Open()
    
    The syscall package checks for null bytes now.
    This was added in https://codereview.appspot.com/6458050
    
    Change-Id: I59a2fed3757a25b85e2668905ff5cf2ec8c3a0d3
    Reviewed-on: https://go-review.googlesource.com/34978
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6d6360648a6820306a5ead6e2a0bf75ad0bb78d1
Author: Jeff Johnson <jrjohnson@google.com>
Date:   Tue Jan 24 11:45:20 2017 -0800

    time: defer loading ZONEINFO until first time.LoadLocation call
    
    A user application can now use os.Setenv("ZONEINFO", ..) becase the
    value is no longer read on init of the time package.
    
    Fixes #18619
    
    Change-Id: Id8e303d67e6fb9c5d6ea9f969d8c94f6fff1bee3
    Reviewed-on: https://go-review.googlesource.com/35639
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ed8c62b7fb47b717dc2adc9f6d0c90c924c67712
Author: Kyrylo Silin <silin@kyrylo.org>
Date:   Sun Jan 22 17:56:03 2017 +0200

    testing: delete unused stopAlarm()
    
    The function call was removed in:
    ead08e91f6468ab1c35c250ec487935103c580f6
    
    Change-Id: I78fe563c9ea4554811c74130533d2186a65d3033
    Reviewed-on: https://go-review.googlesource.com/35532
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0ce3ce010cf03698494944518e8801b17a00aaa1
Author: Ibrahim AshShohail <ibra.sho@gmail.com>
Date:   Sun Jan 22 02:46:25 2017 +0300

    archive/zip: update the ZIP spec link
    
    Update the link to PKWARE "Application Notes on the .ZIP file format" document.
    Now uses the permanent link according to 1.5 in version 6.3.3 (https://pkware.cachefly.net/webdocs/APPNOTE/APPNOTE-6.3.3.TXT):
    http://www.pkware.com/appnote
    
    Fixes #18738
    
    Change-Id: If252a5fca1dd666e70c2591a83d8714672d02932
    Reviewed-on: https://go-review.googlesource.com/35500
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5b1c6efb314f1426c22a7ba9c5983f41b28e76a7
Author: Ibrahim AshShohail <ibra.sho@gmail.com>
Date:   Tue Jan 24 22:20:04 2017 +0300

    cmd/pprof: remove redundant URLs from error messages in fetch.FetchURL
    
    Errors from http.Client already includes the URL in the message.
    
    Fixes #18754
    
    Change-Id: I65fc25a8f3aa6a2d4627aac3fb47eed8d3c4151a
    Reviewed-on: https://go-review.googlesource.com/35650
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1ef3a77e182b322eaaaf767eb176a093f5f68282
Author: Patrick Pelletier <pp.pelletier@gmail.com>
Date:   Wed Jan 25 22:09:26 2017 -0800

    mime/multipart: allow boundary len <= 70
    
    As per RFC 2046, the boundary for multipart MIME is allowed up to 70
    characters. The old SetBoundary implementation only allowed up to 69 so
    this bumps it to the correct value of 70.
    
    The relevant RFC is at https://www.ietf.org/rfc/rfc2046.txt and section
    5.1.1 defines the boundary specification.
    
    Fixes #18793
    
    Change-Id: I91d2ed4549c3d27d6049cb473bac680a750fb520
    Reviewed-on: https://go-review.googlesource.com/35830
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ec4062f8eab9ce771d0ecece5b5388f960498606
Author: Kevin Burke <kev@inburke.com>
Date:   Thu Jan 26 13:41:56 2017 -0800

    database/sql: fix spelling mistake
    
    Change-Id: I67db3b342929a7bd11f01bf3b9afb49f4da69a0a
    Reviewed-on: https://go-review.googlesource.com/35841
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a246f61637b80e2f3426fae03ede072c8a28474e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Nov 8 22:18:38 2016 -0800

    cmd/compile: report more non-inlineable functions
    
    Many non-inlineable functions were not being
    reported in '-m -m' mode.
    
    Updates #17858.
    
    Change-Id: I7d96361b39dd317f5550e57334a8a6dd1a836598
    Reviewed-on: https://go-review.googlesource.com/32971
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 457ac38e7a4066d27a64845c2569fad2b4d7bc8a
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Feb 1 14:40:58 2017 -0500

    cmd/compile: fix generic.rules
    
    generic.rules wasn't updated when rewritegeneric.go was. This commit
    updates it so that the rewritegeneric.go file can be regenerated.
    
    Fixes #18885.
    
    Change-Id: Ie7dab653ca0a9ea1c255fd12e311a0d9e66afdd2
    Reviewed-on: https://go-review.googlesource.com/36032
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 166b1219b8a5b246c83986c7ecef3d15c85c8150
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Jan 31 14:32:11 2017 -0800

    cmd/compile: allocate Nodes together with Name/Param/Func
    
    After allocating a Node that needs a Name, Param, and/or Func field,
    we never clear that field, so we can reduce GC overhead slightly by
    allocating them together with the owner Node.
    
    name       old time/op     new time/op     delta
    Template       325ms ± 7%      325ms ± 7%    ~           (p=0.910 n=29+30)
    Unicode        177ms ±12%      173ms ±11%    ~           (p=0.110 n=29+30)
    GoTypes        1.06s ± 7%      1.05s ± 5%  -1.22%        (p=0.027 n=30+30)
    Compiler       4.48s ± 3%      4.47s ± 3%    ~           (p=0.423 n=30+30)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        476M ±22%       467M ±14%    ~           (p=0.310 n=29+30)
    Unicode         298M ±22%       294M ±25%    ~           (p=0.335 n=30+30)
    GoTypes        1.54G ± 9%      1.48G ± 9%  -4.06%        (p=0.000 n=30+30)
    Compiler       6.26G ± 6%      6.14G ± 6%  -1.90%        (p=0.004 n=30+30)
    
    name       old alloc/op    new alloc/op    delta
    Template      40.9MB ± 0%     41.1MB ± 0%  +0.53%        (p=0.000 n=30+30)
    Unicode       30.9MB ± 0%     31.0MB ± 0%  +0.16%        (p=0.000 n=30+30)
    GoTypes        122MB ± 0%      123MB ± 0%  +0.37%        (p=0.000 n=30+30)
    Compiler       477MB ± 0%      479MB ± 0%  +0.37%        (p=0.000 n=30+29)
    
    name       old allocs/op   new allocs/op   delta
    Template        400k ± 1%       376k ± 1%  -5.96%        (p=0.000 n=30+30)
    Unicode         330k ± 1%       325k ± 1%  -1.48%        (p=0.000 n=30+30)
    GoTypes        1.22M ± 0%      1.16M ± 0%  -4.38%        (p=0.000 n=30+30)
    Compiler       4.35M ± 0%      4.13M ± 0%  -5.08%        (p=0.000 n=30+29)
    
    Change-Id: I9bdc7d9223bb32f785df71810564e82d9a76d109
    Reviewed-on: https://go-review.googlesource.com/36022
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b761b07bf909d9ff7d1fdc11083104fd4e28f252
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Jan 31 16:31:11 2017 -0800

    cmd/compile: simplify noding const declarations
    
    By grouping all the logic into constDecl, we're able to get rid of the
    lastconst and lasttype globals, and simplify the logic slightly. Still
    clunky, but much easier to reason about.
    
    Change-Id: I446696c31084b3bfc1fd5d3651655a81ddd159ab
    Reviewed-on: https://go-review.googlesource.com/36023
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1cbc5aa5290437e81859911cf8e022be2448cc09
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Nov 15 22:33:10 2016 -0800

    cmd/compile: insertVarPhis micro-optimization
    
    Algorithmic improvements here are hard.
    Lifting a lookup out of the loop helps a little, though.
    
    To compile the code in #17926:
    
    name  old s/op   new s/op   delta
    Real   146 ± 3%   140 ± 4%  -3.87%  (p=0.002 n=10+10)
    User   143 ± 3%   139 ± 4%  -3.08%  (p=0.005 n=10+10)
    Sys   8.28 ±35%  8.08 ±28%    ~     (p=0.684 n=10+10)
    
    Updates #17926.
    
    Change-Id: Ic255ac8b7b409c1a53791058818b7e2cf574abe3
    Reviewed-on: https://go-review.googlesource.com/33305
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 302474c61c15095406325773172bfb0a819ce3af
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jan 27 14:04:23 2017 -0800

    cmd/compile: disable memory profiling when not in use
    
    The default value of runtime.MemProfileRate
    is non-zero, which means that a small portion
    of allocations go through the (slow) profiled
    allocation path.
    
    This is never useful in the compiler
    unless the -memprofile flag has been passed.
    I noticed this when samples from mprof.go
    showed up in a compiler cpu pprof listing.
    
    name       old time/op      new time/op      delta
    Template        207ms ± 4%       205ms ± 4%  -0.86%  (p=0.001 n=97+90)
    Unicode        91.8ms ± 4%      91.4ms ± 4%  -0.44%  (p=0.030 n=93+93)
    GoTypes         628ms ± 4%       624ms ± 3%  -0.73%  (p=0.001 n=95+92)
    Compiler        2.70s ± 3%       2.69s ± 3%  -0.39%  (p=0.000 n=97+95)
    Flate           131ms ± 5%       130ms ± 4%  -0.82%  (p=0.000 n=93+90)
    GoParser        154ms ± 5%       153ms ± 4%  -0.57%  (p=0.019 n=98+96)
    Reflect         394ms ± 5%       392ms ± 5%  -0.62%  (p=0.026 n=94+97)
    Tar             112ms ± 6%       112ms ± 5%    ~     (p=0.455 n=97+98)
    XML             214ms ± 3%       213ms ± 4%  -0.68%  (p=0.000 n=91+93)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   246user-ms ± 3%  244user-ms ± 4%  -0.48%  (p=0.016 n=92+91)
    Unicode    114user-ms ± 5%  113user-ms ± 4%  -0.78%  (p=0.002 n=98+94)
    GoTypes    817user-ms ± 3%  813user-ms ± 2%  -0.50%  (p=0.006 n=96+94)
    Compiler   3.58user-s ± 2%  3.57user-s ± 2%  -0.38%  (p=0.003 n=97+95)
    Flate      158user-ms ± 5%  157user-ms ± 4%  -0.80%  (p=0.000 n=94+90)
    GoParser   191user-ms ± 4%  191user-ms ± 4%    ~     (p=0.122 n=98+98)
    Reflect    500user-ms ± 4%  498user-ms ± 4%    ~     (p=0.057 n=95+99)
    Tar        134user-ms ± 3%  134user-ms ± 4%    ~     (p=0.529 n=98+98)
    XML        265user-ms ± 3%  265user-ms ± 3%  -0.30%  (p=0.033 n=92+96)
    
    
    Change-Id: Ied5384e337800d567895ff8d47f15d631edf4f0b
    Reviewed-on: https://go-review.googlesource.com/35916
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9f26b9b93e2622ce9d5d5e7824eedf8a4a8957e7
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Oct 31 16:20:42 2016 -0700

    cmd/compile: eliminate iota_
    
    Change-Id: Iad9c1961aedcc754ad2f6010a49f94c5a0a4bfee
    Reviewed-on: https://go-review.googlesource.com/32487
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 986768de7fcf4def65cecd7eb0c34e2cdf92e78c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Jan 8 13:22:24 2017 -0800

    context: lazily initialize cancelCtx done channel
    
    This CL reduces allocations when a context
    created with WithCancel either
    
    (1) never has its Done channel used or
    (2) gets cancelled before its Done channel is used
    
    This is not uncommon. Many contexts are created
    for tasks that end up not using them.
    
    name                                                old time/op    new time/op    delta
    ContextCancelTree/depth=1/Root=Background-8            112ns ± 2%      74ns ± 1%  -34.03%  (p=0.000 n=17+18)
    ContextCancelTree/depth=1/Root=OpenCanceler-8          601ns ± 3%     544ns ± 1%   -9.56%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1/Root=ClosedCanceler-8        367ns ± 4%     257ns ± 1%  -30.01%  (p=0.000 n=20+20)
    ContextCancelTree/depth=10/Root=Background-8          2.91µs ± 2%    2.87µs ± 0%   -1.38%  (p=0.000 n=20+18)
    ContextCancelTree/depth=10/Root=OpenCanceler-8        4.36µs ± 2%    4.26µs ± 1%   -2.34%  (p=0.000 n=20+18)
    ContextCancelTree/depth=10/Root=ClosedCanceler-8      2.02µs ± 2%    1.51µs ± 1%  -25.18%  (p=0.000 n=19+19)
    ContextCancelTree/depth=100/Root=Background-8         30.5µs ± 6%    30.5µs ± 1%     ~     (p=0.941 n=20+20)
    ContextCancelTree/depth=100/Root=OpenCanceler-8       39.8µs ± 1%    41.1µs ± 1%   +3.15%  (p=0.000 n=18+19)
    ContextCancelTree/depth=100/Root=ClosedCanceler-8     17.8µs ± 1%    13.9µs ± 1%  -21.61%  (p=0.000 n=18+20)
    ContextCancelTree/depth=1000/Root=Background-8         302µs ± 1%     313µs ± 0%   +3.62%  (p=0.000 n=20+18)
    ContextCancelTree/depth=1000/Root=OpenCanceler-8       412µs ± 2%     427µs ± 1%   +3.55%  (p=0.000 n=18+19)
    ContextCancelTree/depth=1000/Root=ClosedCanceler-8     178µs ± 1%     139µs ± 1%  -21.80%  (p=0.000 n=19+17)
    
    name                                                old alloc/op   new alloc/op   delta
    ContextCancelTree/depth=1/Root=Background-8             176B ± 0%       80B ± 0%  -54.55%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1/Root=OpenCanceler-8           544B ± 0%      448B ± 0%  -17.65%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1/Root=ClosedCanceler-8         352B ± 0%      160B ± 0%  -54.55%  (p=0.000 n=20+20)
    ContextCancelTree/depth=10/Root=Background-8          3.49kB ± 0%    3.39kB ± 0%   -2.75%  (p=0.000 n=20+20)
    ContextCancelTree/depth=10/Root=OpenCanceler-8        3.86kB ± 0%    3.76kB ± 0%   -2.49%  (p=0.000 n=20+20)
    ContextCancelTree/depth=10/Root=ClosedCanceler-8      1.94kB ± 0%    0.88kB ± 0%  -54.55%  (p=0.000 n=20+20)
    ContextCancelTree/depth=100/Root=Background-8         36.6kB ± 0%    36.5kB ± 0%   -0.26%  (p=0.000 n=20+20)
    ContextCancelTree/depth=100/Root=OpenCanceler-8       37.0kB ± 0%    36.9kB ± 0%   -0.26%  (p=0.000 n=20+20)
    ContextCancelTree/depth=100/Root=ClosedCanceler-8     17.8kB ± 0%     8.1kB ± 0%  -54.55%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1000/Root=Background-8         368kB ± 0%     368kB ± 0%   -0.03%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1000/Root=OpenCanceler-8       368kB ± 0%     368kB ± 0%   -0.03%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1000/Root=ClosedCanceler-8     176kB ± 0%      80kB ± 0%  -54.55%  (p=0.000 n=20+20)
    
    name                                                old allocs/op  new allocs/op  delta
    ContextCancelTree/depth=1/Root=Background-8             3.00 ± 0%      2.00 ± 0%  -33.33%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1/Root=OpenCanceler-8           8.00 ± 0%      7.00 ± 0%  -12.50%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1/Root=ClosedCanceler-8         6.00 ± 0%      4.00 ± 0%  -33.33%  (p=0.000 n=20+20)
    ContextCancelTree/depth=10/Root=Background-8            48.0 ± 0%      47.0 ± 0%   -2.08%  (p=0.000 n=20+20)
    ContextCancelTree/depth=10/Root=OpenCanceler-8          53.0 ± 0%      52.0 ± 0%   -1.89%  (p=0.000 n=20+20)
    ContextCancelTree/depth=10/Root=ClosedCanceler-8        33.0 ± 0%      22.0 ± 0%  -33.33%  (p=0.000 n=20+20)
    ContextCancelTree/depth=100/Root=Background-8            498 ± 0%       497 ± 0%   -0.20%  (p=0.000 n=20+20)
    ContextCancelTree/depth=100/Root=OpenCanceler-8          503 ± 0%       502 ± 0%   -0.20%  (p=0.000 n=20+20)
    ContextCancelTree/depth=100/Root=ClosedCanceler-8        303 ± 0%       202 ± 0%  -33.33%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1000/Root=Background-8         5.00k ± 0%     5.00k ± 0%   -0.02%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1000/Root=OpenCanceler-8       5.00k ± 0%     5.00k ± 0%   -0.02%  (p=0.000 n=20+20)
    ContextCancelTree/depth=1000/Root=ClosedCanceler-8     3.00k ± 0%     2.00k ± 0%  -33.33%  (p=0.000 n=20+20)
    
    Change-Id: Ibd7a0c3d5c847861cf1497f8fead34329413d26d
    Reviewed-on: https://go-review.googlesource.com/34979
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Sameer Ajmani <sameer@golang.org>

commit fd118b69fabda2f5def2f46606a0ec23f093cbb9
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Jan 22 12:12:22 2017 -0500

    cmd/asm, cmd/internal/obj/s390x: fix encoding of VREPI{H,F,G}
    
    Also adds tests for all missing VRI-a instructions (which may be
    affected by this change).
    
    Fixes #18749.
    
    Change-Id: I48249dda626f32555da9ab58659e2e140de6504a
    Reviewed-on: https://go-review.googlesource.com/35561
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 7f31971f594edbacbdba5407aaee042850fbd220
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sat Jan 21 20:07:26 2017 -0800

    testing: synchronize writes to the root's Writer
    
    Prior to this change it was possible to see interleaved messages:
    <<<
    === RUN   Test/LongLongLongLongName48
    === RUN   Test/LongLon=== RUN   Test/LongLongLongLongName50
    gLongLongName49
    === RUN   Test/LongLongLongLongName51
    >>>
    
    This change fixes it such that you see:
    <<<
    === RUN   Test/LongLongLongLongName48
    === RUN   Test/LongLongLongLongName49
    === RUN   Test/LongLongLongLongName50
    === RUN   Test/LongLongLongLongName51
    >>>
    
    Fixes #18741
    
    Change-Id: I2529d724065dc65b3e9eb3d7cbeeda82a2d0cfd4
    Reviewed-on: https://go-review.googlesource.com/35556
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>

commit 048b8cecc6e74b50205e803ca387ffaa7e9f37fe
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Jan 31 12:54:45 2017 -0800

    reflect: adjust documentation on Value
    
    Make the documentation more explicit that it is not safe to directly
    compare Value. Get straight to the point on how to do it correctly.
    
    Updates #18871
    
    Change-Id: I2aa3253f779636b2f72a1aae8c9bb45d3c32c902
    Reviewed-on: https://go-review.googlesource.com/36018
    Reviewed-by: Keith Randall <khr@golang.org>

commit 47ce87877b1e2d4f34bb93fe6c7d88785b318cd5
Merge: c47df7ae17 ec63158d71
Author: Russ Cox <rsc@golang.org>
Date:   Wed Feb 1 09:35:27 2017 -0500

    all: merge dev.inline into master
    
    Change-Id: I7715581a04e513dcda9918e853fa6b1ddc703770

commit c47df7ae171b1470f8304c6759caf68f3f37ea90
Merge: 7d8bfdde45 f8b4123613
Author: Russ Cox <rsc@golang.org>
Date:   Tue Jan 31 12:57:12 2017 -0500

    all: merge dev.typealias into master
    
    For #18130.
    
    f8b4123613 [dev.typealias] spec: use term 'embedded field' rather than 'anonymous field'
    9ecc3ee252 [dev.typealias] cmd/compile: avoid false positive cycles from type aliases
    49b7af8a30 [dev.typealias] reflect: add test for type aliases
    9bbb07ddec [dev.typealias] cmd/compile, reflect: fix struct field names for embedded byte, rune
    43c7094386 [dev.typealias] reflect: fix StructOf use of StructField to match StructField docs
    9657e0b077 [dev.typealias] cmd/doc: update for type alias
    de2e5459ae [dev.typealias] cmd/compile: declare methods after resolving receiver type
    9259f3073a [dev.typealias] test: match gccgo error messages on alias2.go
    5d92916770 [dev.typealias] cmd/compile: change Func.Shortname to *Sym
    a7c884efc1 [dev.typealias] go/internal/gccgoimporter: support for type aliases
    5802cfd900 [dev.typealias] cmd/compile: export/import test cases for type aliases
    d7cabd40dd [dev.typealias] go/types: clarified doc string
    cc2dcce3d7 [dev.typealias] cmd/compile: a few better comments related to alias types
    5c160b28ba [dev.typealias] cmd/compile: improved error message for cyles involving type aliases
    b2386dffa1 [dev.typealias] cmd/compile: type-check type alias declarations
    ac8421f9a5 [dev.typealias] cmd/compile: various minor cleanups
    f011e0c6c3 [dev.typealias] cmd/compile, go/types, go/importer: various alias related fixes
    49de5f0351 [dev.typealias] cmd/compile, go/importer: define export format and implement importing of type aliases
    5ceec42dc0 [dev.typealias] go/types: export TypeName.IsAlias so clients can use it
    aa1f0681bc [dev.typealias] go/types: improved Object printing
    c80748e389 [dev.typealias] go/types: remove some more vestiges of prior alias implementation
    80d8b69e95 [dev.typealias] go/types: implement type aliases
    a917097b5e [dev.typealias] go/build: add go1.9 build tag
    3e11940437 [dev.typealias] cmd/compile: recognize type aliases but complain for now (not yet supported)
    e0a05c274a [dev.typealias] cmd/gofmt: added test cases for alias type declarations
    2e5116bd99 [dev.typealias] go/ast, go/parser, go/printer, go/types: initial type alias support
    
    Change-Id: Ia65f2e011fd7195f18e1dce67d4d49b80a261203

commit 7d8bfdde453445affb50fcaeacc050938ec98467
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sun Jan 29 20:53:35 2017 +0100

    testing: stop timeout-timer after running tests
    
    Fixes #18845
    
    Change-Id: Icdc3e2067807781e42f2ffc94d1824aed94d3713
    Reviewed-on: https://go-review.googlesource.com/35956
    Run-TryBot: Alberto Donizetti <alb.donizetti@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f8b4123613a2cb0c453726033a03a1968205ccae
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Jan 10 16:19:14 2017 -0800

    [dev.typealias] spec: use term 'embedded field' rather than 'anonymous field'
    
    First steps towards defining type aliases in the spec.
    This is a nomenclature clarification, not a language change.
    
    The spec used all three terms 'embedded type', 'anonymous field',
    and 'embedded field'. Users where using the terms inconsistently.
    
    The notion of an 'anonymous' field was always misleading since they
    always had a de-facto name. With type aliases that name becomes even
    more important because we may have different names for the same type.
    
    Use the term 'embedded field' consistently and remove competing
    terminology.
    
    For #18130.
    
    Change-Id: I2083bbc85788cab0b2e2cb1ff58b2f979491f001
    Reviewed-on: https://go-review.googlesource.com/35108
    Reviewed-by: Alan Donovan <adonovan@google.com>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 9ecc3ee2523f2db87b5b2d79efdd04abda93fb6e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Jan 26 09:00:56 2017 -0800

    [dev.typealias] cmd/compile: avoid false positive cycles from type aliases
    
    For #18130.
    Fixes #18640.
    
    Change-Id: I26cf1d1b78cca6ef207cc4333f30a9011ef347c9
    Reviewed-on: https://go-review.googlesource.com/35831
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 49b7af8a308c24b1e3f6e83ded9e97513316c8d5
Author: Russ Cox <rsc@golang.org>
Date:   Tue Jan 24 14:59:22 2017 -0500

    [dev.typealias] reflect: add test for type aliases
    
    For #18130.
    
    Change-Id: Idd77cb391178c185227cfd779c70fec16351f825
    Reviewed-on: https://go-review.googlesource.com/35733
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9bbb07ddec63e5e747f1cd9dbf82b7504b29dd09
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 25 10:19:33 2017 -0500

    [dev.typealias] cmd/compile, reflect: fix struct field names for embedded byte, rune
    
    Will also fix type aliases.
    
    Fixes #17766.
    For #18130.
    
    Change-Id: I9e1584d47128782152e06abd0a30ef423d5c30d2
    Reviewed-on: https://go-review.googlesource.com/35732
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 43c70943861ce39b44e5bd577a8c3c2ef18538db
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 25 09:50:36 2017 -0500

    [dev.typealias] reflect: fix StructOf use of StructField to match StructField docs
    
    The runtime internal structField interprets name=="" as meaning anonymous,
    but the exported reflect.StructField has always set Name, even for anonymous
    fields, and also set Anonymous=true.
    
    The initial implementation of StructOf confused the internal and public
    meanings of the StructField, expecting the runtime representation of
    anonymous fields instead of the exported reflect API representation.
    It also did not document this fact, so that users had no way to know how
    to create an anonymous field.
    
    This CL changes StructOf to use the previously documented interpretation
    of reflect.StructField instead of an undocumented one.
    
    The implementation of StructOf also, in some cases, allowed creating
    structs with unexported fields (if you knew how to ask) but set the
    PkgPath incorrectly on those fields. Rather than try to fix that, this CL
    changes StructOf to reject attempts to create unexported fields.
    (I think that may be the right design choice, not just a temporary limitation.
    In any event, it's not the topic for today's work.)
    
    For #17766.
    Fixes #18780.
    
    Change-Id: I585a4e324dc5a90551f49d21ae04d2de9ea04b6c
    Reviewed-on: https://go-review.googlesource.com/35731
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 9657e0b0777f3af3b48908cc39e5ab6d06022422
Author: Russ Cox <rsc@golang.org>
Date:   Tue Jan 24 14:53:31 2017 -0500

    [dev.typealias] cmd/doc: update for type alias
    
    For #18130.
    
    Change-Id: I06b05a2b45a2aa6764053fc51e05883063572dad
    Reviewed-on: https://go-review.googlesource.com/35670
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit de2e5459aecb531a67dad274b789ffeb61dca020
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Jan 23 14:24:24 2017 -0800

    [dev.typealias] cmd/compile: declare methods after resolving receiver type
    
    For #18130.
    Fixes #18655.
    
    Change-Id: I58e2f076b9d8273f128cc033bba9edcd06c81567
    Reviewed-on: https://go-review.googlesource.com/35575
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 9259f3073afe0830ab1484bfee46bfa1f322e7e7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jan 24 12:43:52 2017 -0800

    [dev.typealias] test: match gccgo error messages on alias2.go
    
    For #18130.
    
    Change-Id: I9561ee2b8a9f7b11f0851f281a899f78b9e9703e
    Reviewed-on: https://go-review.googlesource.com/35640
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 5d92916770ef57aeb2ae2cb556285d5e093c3aa0
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Jan 23 13:40:25 2017 -0800

    [dev.typealias] cmd/compile: change Func.Shortname to *Sym
    
    A Func's Shortname is just an identifier. No need for an entire ONAME
    Node.
    
    Change-Id: Ie4d397e8d694c907fdf924ce57bd96bdb4aaabca
    Reviewed-on: https://go-review.googlesource.com/35574
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit a7c884efc14368750e30067367b6eab57ed06c0e
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Jan 17 15:06:04 2017 -0800

    [dev.typealias] go/internal/gccgoimporter: support for type aliases
    
    For #18130.
    
    Change-Id: Iac182a6c5bc62633eb02191d9da6166d3b254c4c
    Reviewed-on: https://go-review.googlesource.com/35268
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 5802cfd900c238baeb835bff62bad61c4f4c9852
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Jan 13 17:23:01 2017 -0800

    [dev.typealias] cmd/compile: export/import test cases for type aliases
    
    Plus a few minor changes.
    
    For #18130.
    
    Change-Id: Ica6503fe9c888cc05c15b46178423f620c087491
    Reviewed-on: https://go-review.googlesource.com/35233
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit d7cabd40dd92e5364969273229373d515bee14e8
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Jan 17 11:40:04 2017 -0800

    [dev.typealias] go/types: clarified doc string
    
    Also: removed internal TODO and added better comment
    
    Fixes #18644.
    
    Change-Id: I3e3763d3afdad6937173cdd32fc661618fb60820
    Reviewed-on: https://go-review.googlesource.com/35245
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit ec63158d7104ab6eb3765f7d4ea48744f97d9ff9
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Jan 11 15:48:30 2017 -0800

    [dev.inline] cmd/compile: parse source files concurrently
    
    Conversion to Nodes still happens sequentially at the moment.
    
    Change-Id: I3407ba0711b8b92e22ece0a06fefaff863c3ccc9
    Reviewed-on: https://go-review.googlesource.com/35126
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b90aed020dc9bd430c9a451386550d26c2355ea5
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Jan 11 15:31:48 2017 -0800

    [dev.inline] cmd/compile: reorganize file parsing logic
    
    Preparation for concurrent parsing. No behavior change.
    
    Change-Id: Ic1ec45fc3cb316778c29065cf055c82e92ffa874
    Reviewed-on: https://go-review.googlesource.com/35125
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e48919bcde90d283e01e903eae8f92da4a3c1103
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Jan 11 15:20:38 2017 -0800

    [dev.inline] cmd/compile: split mkpackage into separate functions
    
    Previously, mkpackage jumbled together three unrelated tasks: handling
    package declarations, clearing imports from processing previous source
    files, and assigning a default value to outfile.
    
    Change-Id: I1e124335768aeabfd1a6d9cc2499fbb980d951cf
    Reviewed-on: https://go-review.googlesource.com/35124
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cc2dcce3d748b5585a1da5a9aa2e7ab8b8be00cd
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Jan 12 15:21:21 2017 -0800

    [dev.typealias] cmd/compile: a few better comments related to alias types
    
    For #18130.
    
    Change-Id: I50bded3af0db673fc92b20c41a86b9cae614acd9
    Reviewed-on: https://go-review.googlesource.com/35191
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 5c160b28baceb263fbd95ea0c95f5083e191c114
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Jan 12 14:35:20 2017 -0800

    [dev.typealias] cmd/compile: improved error message for cyles involving type aliases
    
    Known issue: #18640 (requires a bit more work, I believe).
    
    For #18130.
    
    Change-Id: I53dc26012070e0c79f63b7c76266732190a83d47
    Reviewed-on: https://go-review.googlesource.com/35129
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b2386dffa1f646f06c230f9b317cb3640fef11d4
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Jan 11 11:24:35 2017 -0800

    [dev.typealias] cmd/compile: type-check type alias declarations
    
    Known issues:
    - needs many more tests
    - duplicate method declarations via type alias names are not detected
    - type alias cycle error messages need to be improved
    - need to review setup of byte/rune type aliases
    
    For #18130.
    
    Change-Id: Icc2fefad6214e5e56539a9dcb3fe537bf58029f8
    Reviewed-on: https://go-review.googlesource.com/35121
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ac8421f9a58c2c4df9072d1702783baa62eb99f3
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 29 15:49:01 2016 -0800

    [dev.typealias] cmd/compile: various minor cleanups
    
    Also: Don't allow type pragmas with type alias declarations.
    
    For #18130.
    
    Change-Id: Ie54ea5fefcd677ad87ced03466bbfd783771e974
    Reviewed-on: https://go-review.googlesource.com/35102
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f011e0c6c378427f32bbf09f24ba211f7bd96b9c
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Dec 27 16:53:33 2016 -0800

    [dev.typealias] cmd/compile, go/types, go/importer: various alias related fixes
    
    cmd/compile:
    - remove crud from prior alias implementation
    - better comments in places
    
    go/types:
    - fix TypeName.IsAlias predicate
    - more tests
    
    go/importer (go/internal/gcimporter15):
    - handle "@" format for anonymous fields using aliases
      (currently tested indirectly via x/tools/gcimporter15 tests)
    
    For #18130.
    
    Change-Id: I23a6d4e3a4c2a5c1ae589513da73fde7cad5f386
    Reviewed-on: https://go-review.googlesource.com/35101
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 49de5f035169526675b9d5897753d257bf2c7965
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 22 13:04:04 2016 -0800

    [dev.typealias] cmd/compile, go/importer: define export format and implement importing of type aliases
    
    This defines the (tentative) export/import format for type aliases.
    
    The compiler doesn't support type aliases yet, so while the code is present
    it is guarded with a flag.
    
    The export format for embedded (anonymous) fields now has three modes (mode 3 is new):
    1) The original type name and the anonymous field name are the same, and the name is exported:
       we don't need the field name and write "" instead
    2) The original type name and the anonymous field name are the same, and the name is not exported:
       we don't need the field name and write "?" instead, indicating that there is package info
    3) The original type name and the anonymous field name are different:
       we do need the field name and write "@" followed by the field name (and possible package info)
    
    For #18130.
    
    Change-Id: I790dad826757233fa71396a210f966c6256b75d3
    Reviewed-on: https://go-review.googlesource.com/35100
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 5ceec42dc0db9342bc4f37503844b46cf2689c65
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 22 14:42:05 2016 -0800

    [dev.typealias] go/types: export TypeName.IsAlias so clients can use it
    
    For #18130.
    
    Change-Id: I634eaaeaa11e92fc31219d70419fdb4a7aa6e0b4
    Reviewed-on: https://go-review.googlesource.com/35099
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit aa1f0681bc34da2088fec08773eacebc3aee7391
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 22 11:27:48 2016 -0800

    [dev.typealias] go/types: improved Object printing
    
    - added internal isAlias predicated and test
    - use it for improved Object printing
    - when printing a basic type object, don't repeat type name
      (i.e., print "type int" rather than "type int int")
    - added another test to testdata/decls4.src
    
    For #18130.
    
    Change-Id: Ice9517c0065a2cc465c6d12f87cd27c01ef801e6
    Reviewed-on: https://go-review.googlesource.com/35093
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit c80748e3894b5623681fc5f1059ffdbd2cff6b7c
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 22 09:48:43 2016 -0800

    [dev.typealias] go/types: remove some more vestiges of prior alias implementation
    
    For #18130.
    
    Change-Id: Ibec8efd158d32746978242910dc71e5ed23e9d91
    Reviewed-on: https://go-review.googlesource.com/35092
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 80d8b69e95a4514f6567d3b314aa3434ec924363
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Dec 19 17:55:47 2016 -0800

    [dev.typealias] go/types: implement type aliases
    
    Now a TypeName is just that: a name for a type (not just Named and Basic types
    as before). If it happens to be an alias, its type won't be a Named or Basic type,
    or it won't have the same name. We can determine this externally.
    
    It may be useful to provide a helper predicate to make that test easily accessible,
    but we can get to that if there's an actual need.
    
    The field/method lookup code has become more general an simpler, which is a good sign.
    The changes in methodset.go are symmetric to the changes in lookup.go.
    
    Known issue: Cycles created via alias types are not properly detected at the moment.
    
    For #18130.
    
    Change-Id: I90a3206be13116f89c221b5ab4d0f577eec6c78a
    Reviewed-on: https://go-review.googlesource.com/35091
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit a917097b5e5fd42bb4e6f4884a58544330d34984
Author: Russ Cox <rsc@golang.org>
Date:   Mon Jan 9 19:41:19 2017 -0500

    [dev.typealias] go/build: add go1.9 build tag
    
    It's earlier than usual but this will help us put the type alias-aware
    code into x/tools without breaking clients on go1.6, go1.7,
    or (eventually) go1.8.
    
    Change-Id: I43e7ea804922de07d153c7e356cf95e2a11fc592
    Reviewed-on: https://go-review.googlesource.com/35050
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3e119404372fd0d47de1458802b68522f593bf36
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 16 16:28:30 2016 -0800

    [dev.typealias] cmd/compile: recognize type aliases but complain for now (not yet supported)
    
    Added test file.
    
    For #18130.
    
    Change-Id: Ifcfd7cd1acf9dd6a2f4f3d85979d232bb6b8c6b1
    Reviewed-on: https://go-review.googlesource.com/34988
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e0a05c274aa5a3917c5e53f72537e38bb05c10d6
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 16 15:56:05 2016 -0800

    [dev.typealias] cmd/gofmt: added test cases for alias type declarations
    
    For #18130.
    
    Change-Id: I95e84130df40db5241e0cc25c36873c3281199ff
    Reviewed-on: https://go-review.googlesource.com/34987
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2e5116bd999be18691d860e47cb87f1446cf70fe
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 16 15:10:07 2016 -0800

    [dev.typealias] go/ast, go/parser, go/printer, go/types: initial type alias support
    
    Parsing and printing support for type aliases complete.
    go/types recognizes them an issues an "unimplemented" error for now.
    
    For #18130.
    
    Change-Id: I9f2f7b1971b527276b698d9347bcd094ef0012ee
    Reviewed-on: https://go-review.googlesource.com/34986
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 33c036867f11e1e6b874af68a606be7605cd6daf
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Dec 28 17:58:51 2016 -0800

    [dev.inline] cmd/internal/obj: remove vestiges of LineHist - not used anymore
    
    Change-Id: I9d3fcdd5b002953fa9d2f001bf7a834073443794
    Reviewed-on: https://go-review.googlesource.com/34722
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 472c792e0a09bd3d6483ff31863bb0492f27fe33
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 15 17:17:01 2016 -0800

    [dev.inline] cmd/internal/src: introduce compact source position representation
    
    XPos is a compact (8 instead of 16 bytes on a 64bit machine) source
    position representation. There is a 1:1 correspondence between each
    XPos and each regular Pos, translated via a global table.
    
    In some sense this brings back the LineHist, though positions can
    track line and column information; there is a O(1) translation
    between the representations (no binary search), and the translation
    is factored out.
    
    The size increase with the prior change is brought down again and
    the compiler speed is in line with the master repo (measured on
    the same "quiet" machine as for prior change):
    
    name       old time/op     new time/op     delta
    Template       256ms ± 1%      262ms ± 2%    ~             (p=0.063 n=5+4)
    Unicode        132ms ± 1%      135ms ± 2%    ~             (p=0.063 n=5+4)
    GoTypes        891ms ± 1%      871ms ± 1%  -2.28%          (p=0.016 n=5+4)
    Compiler       3.84s ± 2%      3.89s ± 2%    ~             (p=0.413 n=5+4)
    MakeBash       47.1s ± 1%      46.2s ± 2%    ~             (p=0.095 n=5+5)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        309M ± 1%       314M ± 2%    ~             (p=0.111 n=5+4)
    Unicode         165M ± 1%       172M ± 9%    ~             (p=0.151 n=5+5)
    GoTypes        1.14G ± 2%      1.12G ± 1%    ~             (p=0.063 n=5+4)
    Compiler       5.00G ± 1%      4.96G ± 1%    ~             (p=0.286 n=5+4)
    
    Change-Id: Icc570cc60ab014d8d9af6976f1f961ab8828cc47
    Reviewed-on: https://go-review.googlesource.com/34506
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4808fc444307fa683bf3df6d55f9ad1828891a36
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 9 17:15:05 2016 -0800

    [dev.inline] cmd/internal/src: replace src.Pos with syntax.Pos
    
    This replaces the src.Pos LineHist-based position tracking with
    the syntax.Pos implementation and updates all uses.
    
    The LineHist table is not used anymore - the respective code is still
    there but should be removed eventually. CL forthcoming.
    
    Passes toolstash -cmp when comparing to the master repo (with the
    exception of a couple of swapped assembly instructions, likely due
    to different instruction scheduling because the line-based sorting
    has changed; though this is won't affect correctness).
    
    The sizes of various important compiler data structures have increased
    significantly (see the various sizes_test.go files); this is probably
    the reason for an increase of compilation times (to be addressed). Here
    are the results of compilebench -count 5, run on a "quiet" machine (no
    apps running besides a terminal):
    
    name       old time/op     new time/op     delta
    Template       256ms ± 1%      280ms ±15%  +9.54%          (p=0.008 n=5+5)
    Unicode        132ms ± 1%      132ms ± 1%    ~             (p=0.690 n=5+5)
    GoTypes        891ms ± 1%      917ms ± 2%  +2.88%          (p=0.008 n=5+5)
    Compiler       3.84s ± 2%      3.99s ± 2%  +3.95%          (p=0.016 n=5+5)
    MakeBash       47.1s ± 1%      47.2s ± 2%    ~             (p=0.841 n=5+5)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        309M ± 1%       326M ± 2%  +5.18%          (p=0.008 n=5+5)
    Unicode         165M ± 1%       168M ± 4%    ~             (p=0.421 n=5+5)
    GoTypes        1.14G ± 2%      1.18G ± 1%  +3.47%          (p=0.008 n=5+5)
    Compiler       5.00G ± 1%      5.16G ± 1%  +3.12%          (p=0.008 n=5+5)
    
    Change-Id: I241c4246cdff627d7ecb95cac23060b38f9775ec
    Reviewed-on: https://go-review.googlesource.com/34273
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2d429f01bd917c42e66e1991eab9c2e33d813d16
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 9 15:38:23 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: add predicates to syntax.Pos
    
    This moves syntax.Pos closer to cmd/internal/src.Pos so that
    we can more easily replace src.Pos with syntax.Pos going forward.
    
    Change-Id: I9f93a65fecb4c22591edca4b9d6cda39cf0e872e
    Reviewed-on: https://go-review.googlesource.com/34270
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a0c5405c18568900a9b7365297adc1cd846bbdac
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 9 14:28:49 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: add tests for //line directives
    
    Change-Id: I77dc73bfe79e43bbadf85d7eb3c5f8990ec72023
    Reviewed-on: https://go-review.googlesource.com/34248
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f3b56de4d2a9ad5a3ed538455158b8e003b2e25e
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 9 11:14:26 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: report byte offset rather then rune count for column value
    
    This will only become user-visible if error messages show column information.
    Per the discussion in #10324.
    
    For #10324.
    
    Change-Id: I5959c1655aba74bb1a22fdc261cd728ffcfa6912
    Reviewed-on: https://go-review.googlesource.com/34244
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 48d029fe431f2c19e0ccc62a33de059c7725ee93
Author: David Lazar <lazard@golang.org>
Date:   Fri Dec 9 14:30:40 2016 -0500

    [dev.inline] cmd/internal/obj: rename Prog.Lineno to Prog.Pos
    
    Change-Id: I7585d85907869f5a286b36936dfd035f1e8e9906
    Reviewed-on: https://go-review.googlesource.com/34197
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit ad4efedc6ce317aa9eb1208950daee4c31b023cc
Author: David Lazar <lazard@golang.org>
Date:   Fri Dec 9 12:34:01 2016 -0500

    [dev.inline] cmd/internal/obj: use src.Pos in obj.Prog
    
    This will let us use the src.Pos struct to thread inlining
    information through to obj.
    
    Change-Id: I96a16d3531167396988df66ae70f0b729049cc82
    Reviewed-on: https://go-review.googlesource.com/34195
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 4b8895e2ddb8b9aa324417a0d01e6c09c9822e75
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 2 16:22:45 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: remove gcCompat uses in scanner
    
    - make the scanner unconditionally gc compatible
    - consistently use "invalid" instead "illegal" in errors
    
    Reviewed in and cherry-picked from https://go-review.googlesource.com/#/c/33896/.
    
    Change-Id: I4c4253e7392f3311b0d838bbe503576c9469b203
    Reviewed-on: https://go-review.googlesource.com/34237
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 3d5df64b3fe16757c9f271c2421715ba6d79b02d
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Dec 2 10:44:34 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: use syntax.Pos for all external positions
    
    - use syntax.Pos in syntax.Error (rather than line, col)
    - use syntax.Pos in syntax.PragmaHandler (rather than just line)
    - update uses
    - better documentation in various places
    
    Also:
    - make Pos methods use Pos receiver (rather than *Pos)
    
    Reviewed in and cherry-picked from https://go-review.googlesource.com/#/c/33891/.
    With minor adjustments to noder.go to make merge compile.
    
    Change-Id: I5507cea6c2be46a7677087c1aeb69382d31033eb
    Reviewed-on: https://go-review.googlesource.com/34236
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 54ef0447fed1a59b95111b86a037c3443daf0b9b
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 1 22:04:49 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: clean up error and pragma handling
    
    Reviewed in and cherry-picked from https://go-review.googlesource.com/#/c/33873/.
    
    - simplify error handling in source.go
      (move handling of first error into parser, where it belongs)
    
    - clean up error handling in scanner.go
    
    - move pragma and position base handling from scanner
      to parser where it belongs
    
    - have separate error methods in parser to avoid confusion
      with handlers from scanner.go and source.go
    
    - (source.go) and (scanner.go, source.go, tokens.go)
      may be stand-alone packages if so desired, which means
      these files are now less entangled and easier to maintain
    
    Change-Id: I81510fc7ef943b78eaa49092c0eab2075a05878c
    Reviewed-on: https://go-review.googlesource.com/34235
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e97c8a592f20d390a97db1d516782c56badf258d
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 1 15:25:07 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: simplified position code
    
    Reviewed in and cherry-picked from https://go-review.googlesource.com/#/c/33805/.
    
    Change-Id: I859d9bd5f2256ca78f7b24b330290f7ae600854d
    Reviewed-on: https://go-review.googlesource.com/34234
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 32bf2829a17a90bdbd472335707639ba35776da6
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Nov 30 23:28:40 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: process //line pragmas in scanner
    
    Reviewed in and cherry-picked from https://go-review.googlesource.com/#/c/33764/.
    
    Minor adjustment in noder.go to make merge compile again.
    
    Change-Id: Ib5029b52b59944f207b0f2438c8a5aa576eb25b8
    Reviewed-on: https://go-review.googlesource.com/34233
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8d20b25779d4ce32e8eaeb52374fba1e74f7df57
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Nov 29 16:13:09 2016 -0800

    [dev.inline] cmd/compile/internal/syntax: introduce general position info for nodes
    
    Reviewed in and cherry-picked from https://go-review.googlesource.com/#/c/33758/.
    Minor adjustments in noder.go to fix merge.
    
    Change-Id: Ibe429e327c7f8554f8ac205c61ce3738013aed98
    Reviewed-on: https://go-review.googlesource.com/34231
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit eaca0e0529b780f4c862a97aa47008aa1b403adf
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 8 15:19:47 2016 -0800

    [dev.inline] cmd/internal/src: introduce NoPos and use it instead Pos{}
    
    Using a variable instead of a composite literal makes
    the code independent of implementation changes of Pos.
    
    Per David Lazar's suggestion.
    
    Change-Id: I336967ac12a027c51a728a58ac6207cb5119af4a
    Reviewed-on: https://go-review.googlesource.com/34148
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c10499b539b964d647a9153cbf44e9c39661c397
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Dec 8 13:49:51 2016 -0800

    [dev.inline] cmd/compile/internal/ssa: another round of renames from line -> pos (cleanup)
    
    Mostly mechanical renames. Make variable names consistent with use.
    
    Change-Id: Iaa89d31deab11eca6e784595b58e779ad525c8a3
    Reviewed-on: https://go-review.googlesource.com/34146
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit cfd17f51c87765fbd2b9c32e54722a32975bedf2
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Dec 7 18:14:35 2016 -0800

    [dev.inline] cmd/compile/internal/ssa: rename various fields from Line to Pos
    
    This is a mostly mechanical rename followed by manual fixes where necessary.
    
    Change-Id: Ie5c670b133db978f15dc03e50dc2da0c80fc8842
    Reviewed-on: https://go-review.googlesource.com/34137
    Reviewed-by: David Lazar <lazard@golang.org>

commit eab3707d6d5a746eb60011c40831ea9083ae533c
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Dec 7 17:40:46 2016 -0800

    [dev.inline] cmd/compile: rename various fields from Lineno to Pos
    
    Various minor adjustments.
    
    Change-Id: Iedfb97989f7bedaa3e9e8993b167e05f162434a7
    Reviewed-on: https://go-review.googlesource.com/34136
    Reviewed-by: David Lazar <lazard@golang.org>

commit 82d0caea2c5041a0d0260ff5ec7f7b61ee2bb0af
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Dec 7 16:02:42 2016 -0800

    [dev.inline] cmd/internal/src: make Pos implementation abstract
    
    Adjust cmd/compile accordingly.
    
    This will make it easier to replace the underlying implementation.
    
    Change-Id: I33645850bb18c839b24785b6222a9e028617addb
    Reviewed-on: https://go-review.googlesource.com/34133
    Reviewed-by: David Lazar <lazard@golang.org>

commit 24597c080bdba1de8f7e5d46aa250e5f25d24311
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Dec 6 17:08:06 2016 -0800

    [dev.inline] cmd/compile: introduce cmd/internal/src.Pos type for line numbers
    
    This is a step toward chosing a different position representation.
    By introducing an explicit type, it will be easier to make the
    transition step-wise while ensuring everything keeps running.
    
    This has been reviewed via https://go-review.googlesource.com/#/c/34025/.
    
    Change-Id: Ibceddcd62d8f346321ac3250e3940e9c436ed684
    Reviewed-on: https://go-review.googlesource.com/34132
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Lazar <lazard@golang.org>
```
