# July 28, 2016 Report

Number of commits: 15

## Compilation time

* github.com/coreos/etcd/cmd: from 15.023236464s to 13.174489415s, -12.31%
* github.com/boltdb/bolt/cmd/bolt: from 563.644886ms to 567.33251ms, +0.65%
* github.com/gogits/gogs: from 13.006048329s to 13.223952731s, +1.68%
* github.com/spf13/hugo: from 7.89039783s to 9.666340224s, +22.51%
* github.com/nsqio/nsq/apps/nsqd: from 2.151817449s to 1.988566218s, -7.59%
* github.com/mholt/caddy: from 274.303597ms to 286.510548ms, +4.45%

## Binary size:

* github.com/coreos/etcd/cmd: from 26607017 to 26623999, ~
* github.com/boltdb/bolt/cmd/bolt: from 2679371 to 2679371, ~
* github.com/gogits/gogs: from 23744607 to 23761589, ~
* github.com/spf13/hugo: from 15843701 to 15864779, +0.13%
* github.com/nsqio/nsq/apps/nsqd: from 10051324 to 10064210, +0.13%
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               191           188           -1.57%
BenchmarkMsgpUnmarshal-4             404           405           +0.25%
BenchmarkEasyJsonMarshal-4           1593          1598          +0.31%
BenchmarkEasyJsonUnmarshal-4         1554          1525          -1.87%
BenchmarkFlatBuffersMarshal-4        369           370           +0.27%
BenchmarkFlatBuffersUnmarshal-4      290           288           -0.69%
BenchmarkGogoprotobufMarshal-4       163           161           -1.23%
BenchmarkGogoprotobufUnmarshal-4     252           257           +1.98%

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
commit be915159073ed93fa511ceef7256bc8ee396d1c7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Jul 27 08:51:47 2016 +0200

    doc/go1.7.html: add known issues section for FreeBSD crashes
    
    Updates #16396
    
    Change-Id: I7b4f85610e66f2c77c17cf8898cc41d81b2efc8c
    Reviewed-on: https://go-review.googlesource.com/25283
    Reviewed-by: Chris Broadfoot <cbro@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit ccca9c9cc0fa5b6ea6e5c8276a96eee8c27ebd87
Author: Rhys Hiltner <rhys@justin.tv>
Date:   Fri Jul 22 16:36:30 2016 -0700

    runtime: reduce GC assist extra credit
    
    Mutator goroutines that allocate memory during the concurrent mark
    phase are required to spend some time assisting the garbage
    collector. The magnitude of this mandatory assistance is proportional
    to the goroutine's allocation debt and subject to the assistance
    ratio as calculated by the pacer.
    
    When assisting the garbage collector, a mutator goroutine will go
    beyond paying off its allocation debt. It will build up extra credit
    to amortize the overhead of the assist.
    
    In fast-allocating applications with high assist ratios, building up
    this credit can take the affected goroutine's entire time slice.
    Reduce the penalty on each goroutine being selected to assist the GC
    in two ways, to spread the responsibility more evenly.
    
    First, do a consistent amount of extra scan work without regard for
    the pacer's assistance ratio. Second, reduce the magnitude of the
    extra scan work so it can be completed within a few hundred
    microseconds.
    
    Commentary on gcOverAssistWork is by Austin Clements, originally in
    https://golang.org/cl/24704
    
    Updates #14812
    Fixes #16432
    
    Change-Id: I436f899e778c20daa314f3e9f0e2a1bbd53b43e1
    Reviewed-on: https://go-review.googlesource.com/25155
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit c80e0d374ba3caf8ee32c6fe4a5474fa33928086
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Jul 26 23:44:00 2016 +0200

    net/http: fix data race with concurrent use of Server.Serve
    
    Fixes #16505
    
    Change-Id: I0afabcc8b1be3a5dbee59946b0c44d4c00a28d71
    Reviewed-on: https://go-review.googlesource.com/25280
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit 4a15508c663429652d32f5363c0964152b28dd74
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Jul 26 23:58:44 2016 +0200

    crypto/x509: detect OS X version for FetchPEMRoots at run time
    
    https://golang.org/cl/25233 was detecting the OS X release at compile
    time, not run time. Detect it at run time instead.
    
    Fixes #16473 (again)
    
    Change-Id: I6bec4996e57aa50c52599c165aa6f1fae7423fa7
    Reviewed-on: https://go-review.googlesource.com/25281
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit 66b47431cba75ce23630e17c1a3aa000e7b33d06
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Jul 27 00:17:38 2016 +0200

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 6a513af for:
    
      http2: return flow control for closed streams
      https://golang.org/cl/25231
    
      http2: make Transport prefer HTTP response header recv before body write error
      https://golang.org/cl/24984
    
      http2: make Transport treat "Connection: close" the same as Request.Close
      https://golang.org/cl/24982
    
    Fixes golang/go#16481
    
    Change-Id: Iaddb166387ca2df1cfbbf09a166f8605578bec49
    Reviewed-on: https://go-review.googlesource.com/25282
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit b11fff3886705bd98aec4923f43216ed8e13cb1a
Author: Austin Clements <austin@google.com>
Date:   Mon Jul 25 13:41:07 2016 -0400

    runtime/pprof: document use of pprof package
    
    Currently the pprof package gives almost no guidance for how to use it
    and, despite the standard boilerplate used to create CPU and memory
    profiles, this boilerplate appears nowhere in the pprof documentation.
    
    Update the pprof package documentation to give the standard
    boilerplate in a form people can copy, paste, and tweak. This
    boilerplate is based on rsc's 2011 blog post on profiling Go programs
    at https://blog.golang.org/profiling-go-programs, which is where I
    always go when I need to copy-paste the boilerplate.
    
    Change-Id: I74021e494ea4dcc6b56d6fb5e59829ad4bb7b0be
    Reviewed-on: https://go-review.googlesource.com/25182
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit ff60da6962f71871fac3dd6a5406686ea92de8dc
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Jul 26 16:55:40 2016 +0200

    crypto/x509: use Go 1.6 implementation for FetchPEMRoots for OS X 10.8
    
    Conservative fix for the OS X 10.8 crash. We can unify them back together
    during the Go 1.8 dev cycle.
    
    Fixes #16473
    
    Change-Id: If07228deb2be36093dd324b3b3bcb31c23a95035
    Reviewed-on: https://go-review.googlesource.com/25233
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 887606114902bd58c3838767ac2b66dadba27e5e
Author: Jack Lindamood <jlindamo@justin.tv>
Date:   Fri Jul 15 13:28:27 2016 -0700

    context: add test for WithDeadline in the past
    
    Adds a test case for calling context.WithDeadline() where the deadline
    exists in the past.  This change increases the code coverage of the
    context package.
    
    Change-Id: Ib486bf6157e779fafd9dab2b7364cdb5a06be36e
    Reviewed-on: https://go-review.googlesource.com/25007
    Reviewed-by: Sameer Ajmani <sameer@golang.org>
    Run-TryBot: Sameer Ajmani <sameer@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ea2376fcea0be75c856ebd199c0ad0f98192d406
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Jul 22 22:51:05 2016 +0000

    net/http: make Transport.RoundTrip return raw Conn.Read error on peek failure
    
    From at least Go 1.4 to Go 1.6, Transport.RoundTrip would return the
    error value from net.Conn.Read directly when the initial Read (1 byte
    Peek) failed while reading the HTTP response, if a request was
    outstanding. While never a documented or tested promise, Go 1.7 changed the
    behavior (starting at https://golang.org/cl/23160).
    
    This restores the old behavior and adds a test (but no documentation
    promises yet) while keeping the fix for spammy logging reported in #15446.
    
    This looks larger than it is: it just changes errServerClosedConn from
    a variable to a type, where the type preserves the underlying
    net.Conn.Read error, for unwrapping later in Transport.RoundTrip.
    
    Fixes #16465
    
    Change-Id: I6fa018991221e93c0cfe3e4129cb168fbd98bd27
    Reviewed-on: https://go-review.googlesource.com/25153
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 67f799c42cbe5fa667dbad0139a98728624cbf4b
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sat Jul 23 23:27:25 2016 -0400

    doc: add s390x information to asm.html
    
    Fixes #16362
    
    Change-Id: I676718a1149ed2f3ff80cb031e25de7043805399
    Reviewed-on: https://go-review.googlesource.com/25157
    Reviewed-by: Rob Pike <r@golang.org>

commit d0256118de0b397494a3f4ca6d2e1e889b8c114e
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Mon Jul 25 15:49:35 2016 -0700

    compress/flate: document HuffmanOnly
    
    Fixes #16489
    
    Change-Id: I13e2ed6de59102f977566de637d8d09b4e541980
    Reviewed-on: https://go-review.googlesource.com/25200
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 10538a8f9e2e718a47633ac5a6e90415a2c3f5f1
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Jul 22 21:58:18 2016 +0000

    net/http: fix potential for-select spin with closed Context.Done channel
    
    Noticed when investigating a separate issue.
    
    No external bug report or repro yet.
    
    Change-Id: I8a1641a43163f22b09accd3beb25dd9e2a68a238
    Reviewed-on: https://go-review.googlesource.com/25152
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 243d51f05e5dc263e185f9b1f7f1fe96a2098644
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Jul 21 16:58:18 2016 +0000

    misc/trace: disable trace resolution warning
    
    It was removed in upstream Chrome https://codereview.chromium.org/2016863004
    
    Rather than update to the latest version, make the minimal change for Go 1.7 and
    change the "showToUser" boolean from true to false.
    
    Tested by hand that it goes away after this change.
    
    Updates #16247
    
    Change-Id: I051f49da878e554b1a34a88e9abc70ab50e18780
    Reviewed-on: https://go-review.googlesource.com/25117
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 846bc6c5ab396490717f8753cc3c271f9c1391e4
Author: David Chase <drchase@google.com>
Date:   Wed Jul 20 10:44:49 2016 -0400

    cmd/compile: change phi location to be optimistic at backedges
    
    This is:
    
    (1) a simple trick that cuts the number of phi-nodes
    (temporarily) inserted into the ssa representation by a factor
    of 10, and can cut the user time to compile tricky inputs like
    gogo/protobuf tests from 13 user minutes to 9.5, and memory
    allocation from 3.4GB to 2.4GB.
    
    (2) a fix to sparse lookup, that does not rely on
    an assumption proven false by at least one pathological
    input "etldlen".
    
    These two changes fix unrelated compiler performance bugs,
    both necessary to obtain good performance compiling etldlen.
    Without them it takes 20 minutes or longer, with them it
    completes in 2 minutes, without a gigantic memory footprint.
    
    Updates #16407
    
    Change-Id: Iaa8aaa8c706858b3d49de1c4865a7fd79e6f4ff7
    Reviewed-on: https://go-review.googlesource.com/23136
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 305a0ac123cf99d469f5519f8974f4911e690c48
Author: Keith Randall <khr@golang.org>
Date:   Mon Jul 18 23:06:04 2016 -0700

    cmd/compile: move phi args which are constants closer to the phi
    
    entry:
       x = MOVQconst [7]
       ...
    b1:
       goto b2
    b2:
       v = Phi(x, y, z)
    
    Transform that program to:
    
    entry:
       ...
    b1:
       x = MOVQconst [7]
       goto b2
    b2:
       v = Phi(x, y, z)
    
    This CL moves constant-generating instructions used by a phi to the
    appropriate immediate predecessor of the phi's block.
    
    We used to put all constants in the entry block.  Unfortunately, in
    large functions we have lots of constants at the start of the
    function, all of which are used by lots of phis throughout the
    function.  This leads to the constants being live through most of the
    function (especially if there is an outer loop).  That's an O(n^2)
    problem.
    
    Note that most of the non-phi uses of constants have already been
    folded into instructions (ADDQconst, MOVQstoreconst, etc.).
    
    This CL may be generally useful for other instances of compiler
    slowness, I'll have to check.  It may cause some programs to run
    slower, but probably not by much, as rematerializeable values like
    these constants are allocated late (not at their originally scheduled
    location) anyway.
    
    This CL is definitely a minimal change that can be considered for 1.7.
    We probably want to do a better job in the tighten pass generally, not
    just for phi args.  Leaving that for 1.8.
    
    Update #16407
    
    Change-Id: If112a8883b4ef172b2f37dea13e44bda9346c342
    Reviewed-on: https://go-review.googlesource.com/25046
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
```
