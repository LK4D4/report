# July 7, 2016 Report

Number of commits: 19

## Compilation time

* github.com/coreos/etcd/cmd: from 15.13358728s to 15.247372779s, +0.75%
* github.com/boltdb/bolt/cmd/bolt: from 555.067769ms to 558.217736ms, +0.57%
* github.com/gogits/gogs: from 12.984591967s to 12.932498538s, -0.40%
* github.com/spf13/hugo: from 7.215382531s to 9.196958072s, +27.46%
* github.com/nsqio/nsq/apps/nsqd: from 2.079530519s to 2.048365995s, -1.50%
* github.com/mholt/caddy: from 265.533664ms to 286.270739ms, +7.81%

## Binary size:

* github.com/coreos/etcd/cmd: from 26396353 to 26404545, ~
* github.com/boltdb/bolt/cmd/bolt: from 2666587 to 2666587, ~
* github.com/gogits/gogs: from 23642896 to 23642898, ~
* github.com/spf13/hugo: from 15198261 to 15202359, ~
* github.com/nsqio/nsq/apps/nsqd: from 10046829 to 10050925, ~
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               190           190           +0.00%
BenchmarkMsgpUnmarshal-4             406           403           -0.74%
BenchmarkEasyJsonMarshal-4           1590          1587          -0.19%
BenchmarkEasyJsonUnmarshal-4         1533          1538          +0.33%
BenchmarkFlatBuffersMarshal-4        367           359           -2.18%
BenchmarkFlatBuffersUnmarshal-4      294           292           -0.68%
BenchmarkGogoprotobufMarshal-4       163           163           +0.00%
BenchmarkGogoprotobufUnmarshal-4     260           258           -0.77%

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

	<-------------------HIGHLIGHTS HERE---------------------->


## GIT Log

```
commit d8722012afb789f1a2875a0d2ed50bfbae12bb9c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Jul 6 21:29:40 2016 +0000

    net/http: deflake TestClientRedirectContext
    
    The test was checking for 1 of 2 possible error values. But based on
    goroutine scheduling and the randomness of select statement receive
    cases, it was possible for a 3rd type of error to be returned.
    
    This modifies the code (not the test) to make that third type of error
    actually the second type of error, which is a nicer error message.
    
    The test is no longer flaky. The flake was very reproducible with a
    5ms sleep before the select at the end of Transport.getConn.
    
    Thanks to Github user @jaredborner for debugging.
    
    Fixes #16049
    
    Change-Id: I0d2a036c9555a8d2618b07bab01f28558d2b0b2c
    Reviewed-on: https://go-review.googlesource.com/24748
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit df7c159f06ab6d6c7ac6c953e491f8900f40d282
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jul 6 17:14:10 2016 -0700

    path/filepath: fix typo in comment
    
    Change-Id: I0c76e8deae49c1149647de421503c5175028b948
    Reviewed-on: https://go-review.googlesource.com/24781
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 94477121bd1a758a70393773c6ae40c58c54f005
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jul 6 13:07:53 2016 -0700

    path/filepath: document Clean behavior for each function
    
    Document explicitly which functions Clean the result rather than
    documenting it in the package comment.
    
    Updates #10122.
    Fixes #16111.
    
    Change-Id: Ia589c7ee3936c9a6a758725ac7f143054d53e41e
    Reviewed-on: https://go-review.googlesource.com/24747
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit bbe5da42600d5ab26cd58ffac3d6427994f08fb2
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 28 14:19:27 2016 -0700

    cmd/compile, syscall: add //go:uintptrescapes comment, and use it
    
    This new comment can be used to declare that the uintptr arguments to a
    function may be converted from pointers, and that those pointers should
    be considered to escape. This is used for the Call methods in
    dll_windows.go that take uintptr arguments, because they call Syscall.
    
    We can't treat these functions as we do syscall.Syscall, because unlike
    Syscall they may cause the stack to grow. For Syscall we can assume that
    stack arguments can remain on the stack, but for these functions we need
    them to escape.
    
    Fixes #16035.
    
    Change-Id: Ia0e5b4068c04f8d303d95ab9ea394939f1f57454
    Reviewed-on: https://go-review.googlesource.com/24551
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 820e30f5b0289d5df22ab604f2d831470f748dca
Author: Sam Whited <sam@samwhited.com>
Date:   Tue Jul 5 20:06:00 2016 -0500

    encoding/xml: update docs to follow convention
    
    Fixes #8833
    
    Change-Id: I4523a1de112ed02371504e27882659bce8028a45
    Reviewed-on: https://go-review.googlesource.com/24745
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5a9d5c37479231336efef0e0fa5b75645aa1c569
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Jul 3 14:57:35 2016 -0700

    encoding/gob: document Encode, EncodeValue nil pointer panics
    
    Fixes #16258.
    
    Docs for Encode and EncodeValue do not mention that
    nil pointers are not permitted hence we panic,
    because Gobs encode values yet nil pointers have no value
    to encode. It moves a comment that was internal to EncodeValue
    to the top level to make it clearer to users what to expect
    when they pass in nil pointers.
    Supplements test TestTopLevelNilPointer.
    
    Change-Id: Ie54f609fde4b791605960e088456047eb9aa8738
    Reviewed-on: https://go-review.googlesource.com/24740
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 003a68bc7fcb917b5a4d92a5c2244bb1adf8f690
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jul 1 13:52:26 2016 -0700

    cmd/vet: remove copylock warning about result types and calls
    
    Don't issue a copylock warning about a result type; the function may
    return a composite literal with a zero value, which is OK.
    
    Don't issue a copylock warning about a function call on the RHS, or an
    indirection of a function call; the function may return a composite
    literal with a zero value, which is OK.
    
    Updates #16227.
    
    Change-Id: I94f0e066bbfbca5d4f8ba96106210083e36694a2
    Reviewed-on: https://go-review.googlesource.com/24711
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 878e002bb9021822cc44a9e20cf92689a2c478e7
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Mon Jul 4 08:45:10 2016 +0900

    syscall: fix missing use of use function in Getfsstat
    
    Updates #13372.
    
    Change-Id: If383c14af14839a303425ba7b80b97e35ca9b698
    Reviewed-on: https://go-review.googlesource.com/24750
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit afccfb829f17f85a8d4aa74f1a11a27422437405
Author: Monty Taylor <mordred@inaugust.com>
Date:   Fri Jul 1 08:47:41 2016 -0500

    cmd/go: remove noVCSSuffix check for OpenStack
    
    The original intent of the code was to allow both with and without .git
    suffix for now to allow a transition period. The noVCSSuffix check was a
    copy pasta error.
    
    Fixes #15979.
    
    Change-Id: I3d39aba8d026b40fc445244d6d01d8bc1979d1e4
    Reviewed-on: https://go-review.googlesource.com/24645
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 519b469795287dc81ab3b994f8809f61a0c802da
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jul 1 15:44:12 2016 -0700

    cmd/compile: mark live heap-allocated pparamout vars as needzero
    
    If we don't mark them as needzero, we have a live pointer variable
    containing possible garbage, which will baffle the GC.
    
    Fixes #16249.
    
    Change-Id: I7c423ceaca199ddd46fc2c23e5965e7973f07584
    Reviewed-on: https://go-review.googlesource.com/24715
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 575a87166291e321745041944321002b3c0b72be
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Jul 1 10:51:59 2016 -0700

    cmd/compile: don't lose //go:nointerface pragma in export data
    
    Fixes #16243.
    
    Change-Id: I207d1e8aa48abe453a23c709ccf4f8e07368595b
    Reviewed-on: https://go-review.googlesource.com/24648
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 29f0984a3558ef6e3e58a621791473a71b510365
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Jun 30 06:36:31 2016 -0400

    cmd/compile: don't set line number to 0 when building SSA
    
    The frontend may emit node with line number missing. In this case,
    use the parent line number. Instead of changing every call site of
    pushLine, do it in pushLine itself.
    
    Fixes #16214.
    
    Change-Id: I80390550b56e4d690fc770b01ff725b892ffd6dc
    Reviewed-on: https://go-review.googlesource.com/24641
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b5aae1a2845f157a2565b856fb2d7773a0f7af25
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Jun 30 22:11:22 2016 +0000

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev b400c2e for https://golang.org/cl/24214,
    "http2: add additional blacklisted ciphersuites"
    
    Both TLS_RSA_WITH_AES_128_GCM_SHA256 & TLS_RSA_WITH_AES_256_GCM_SHA384
    are now blacklisted, per http://httpwg.org/specs/rfc7540.html#BadCipherSuites
    
    Change-Id: I8b9a7f4dc3c152d0675e196523ddd36111744984
    Reviewed-on: https://go-review.googlesource.com/24684
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 08086e624689e0fdf5b53030ecfb96ea709b6d86
Author: Alan Donovan <adonovan@google.com>
Date:   Thu Jun 30 14:32:03 2016 -0400

    cmd/vet: lostcancel: treat naked return as a use of named results
    
    + test.
    
    Fixes #16230
    
    Change-Id: Idac995437146a9df9e73f094d2a31abc25b1fa62
    Reviewed-on: https://go-review.googlesource.com/24681
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7ea62121a7de25559cb88389983086f45910aed6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Jun 30 19:24:06 2016 +0000

    all: be consistent about spelling of cancelation
    
    We had ~30 one way, and these four new occurrences the other way.
    
    Updates #11626
    
    Change-Id: Ic6403dc4905874916ae292ff739d33482ed8e5bf
    Reviewed-on: https://go-review.googlesource.com/24683
    Reviewed-by: Rob Pike <r@golang.org>

commit fc12bb263683e43c0b93eb00071f894f5cfcc772
Author: Alan Donovan <adonovan@google.com>
Date:   Thu Jun 30 14:55:01 2016 -0400

    context: cancel the context in ExampleWithTimeout, with explanation
    
    Fixes #16230
    
    Change-Id: Ibb10234a6c3ab8bd0cfd93c2ebe8cfa66f80f6b0
    Reviewed-on: https://go-review.googlesource.com/24682
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9c8809f82aa59e0725e93cffb03de863e61cbbae
Author: Austin Clements <austin@google.com>
Date:   Wed Jun 29 17:41:50 2016 -0400

    runtime/internal/sys: implement Ctz and Bswap in assembly for 386
    
    Ctz is a hot-spot in the Go 1.7 memory manager. In SSA it's
    implemented as an intrinsic that compiles to a few instructions, but
    on the old backend (all architectures other than amd64), it's
    implemented as a fairly complex Go function. As a result, switching to
    bitmap-based allocation was a significant hit to allocation-heavy
    workloads like BinaryTree17 on non-SSA platforms.
    
    For unknown reasons, this hit 386 particularly hard. We can regain a
    lot of the lost performance by implementing Ctz in assembly on the
    386. This isn't as good as an intrinsic, since it still generates a
    function call and prevents useful inlining, but it's much better than
    the pure Go implementation:
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              3.59s ± 1%     3.06s ± 1%  -14.74%  (p=0.000 n=19+20)
    Fannkuch11-12                3.72s ± 1%     3.64s ± 1%   -2.09%  (p=0.000 n=17+19)
    FmtFprintfEmpty-12          52.3ns ± 3%    52.3ns ± 3%     ~     (p=0.829 n=20+19)
    FmtFprintfString-12          156ns ± 1%     148ns ± 3%   -5.20%  (p=0.000 n=18+19)
    FmtFprintfInt-12             137ns ± 1%     136ns ± 1%   -0.56%  (p=0.000 n=19+13)
    FmtFprintfIntInt-12          227ns ± 2%     225ns ± 2%   -0.93%  (p=0.000 n=19+17)
    FmtFprintfPrefixedInt-12     210ns ± 1%     208ns ± 1%   -0.91%  (p=0.000 n=19+17)
    FmtFprintfFloat-12           375ns ± 1%     371ns ± 1%   -1.06%  (p=0.000 n=19+18)
    FmtManyArgs-12               995ns ± 2%     978ns ± 1%   -1.63%  (p=0.000 n=17+17)
    GobDecode-12                9.33ms ± 1%    9.19ms ± 0%   -1.59%  (p=0.000 n=20+17)
    GobEncode-12                7.73ms ± 1%    7.73ms ± 1%     ~     (p=0.771 n=19+20)
    Gzip-12                      375ms ± 1%     374ms ± 1%     ~     (p=0.141 n=20+18)
    Gunzip-12                   61.8ms ± 1%    61.8ms ± 1%     ~     (p=0.602 n=20+20)
    HTTPClientServer-12         87.7µs ± 2%    86.9µs ± 3%   -0.87%  (p=0.024 n=19+20)
    JSONEncode-12               20.2ms ± 1%    20.4ms ± 0%   +0.53%  (p=0.000 n=18+19)
    JSONDecode-12               65.3ms ± 0%    65.4ms ± 1%     ~     (p=0.385 n=16+19)
    Mandelbrot200-12            4.11ms ± 1%    4.12ms ± 0%   +0.29%  (p=0.020 n=19+19)
    GoParse-12                  3.75ms ± 1%    3.61ms ± 2%   -3.90%  (p=0.000 n=20+20)
    RegexpMatchEasy0_32-12       104ns ± 0%     103ns ± 0%   -0.96%  (p=0.000 n=13+16)
    RegexpMatchEasy0_1K-12       805ns ± 1%     803ns ± 1%     ~     (p=0.189 n=18+18)
    RegexpMatchEasy1_32-12       111ns ± 0%     111ns ± 3%     ~     (p=1.000 n=14+19)
    RegexpMatchEasy1_1K-12      1.00µs ± 1%    1.00µs ± 1%   +0.50%  (p=0.003 n=19+19)
    RegexpMatchMedium_32-12      133ns ± 2%     133ns ± 2%     ~     (p=0.218 n=20+20)
    RegexpMatchMedium_1K-12     41.2µs ± 1%    42.2µs ± 1%   +2.52%  (p=0.000 n=18+16)
    RegexpMatchHard_32-12       2.35µs ± 1%    2.38µs ± 1%   +1.53%  (p=0.000 n=18+18)
    RegexpMatchHard_1K-12       70.9µs ± 2%    72.0µs ± 1%   +1.42%  (p=0.000 n=19+17)
    Revcomp-12                   1.06s ± 0%     1.05s ± 0%   -1.36%  (p=0.000 n=20+18)
    Template-12                 86.2ms ± 1%    84.6ms ± 0%   -1.89%  (p=0.000 n=20+18)
    TimeParse-12                 425ns ± 2%     428ns ± 1%   +0.77%  (p=0.000 n=18+19)
    TimeFormat-12                517ns ± 1%     519ns ± 1%   +0.43%  (p=0.001 n=20+19)
    [Geo mean]                  74.3µs         73.5µs        -1.05%
    
    Prior to this commit, BinaryTree17-12 on 386 was 33% slower than at
    the go1.6 tag. With this commit, it's 13% slower.
    
    On arm and arm64, BinaryTree17-12 is only ~5% slower than it was at
    go1.6. It may be worth implementing Ctz for them as well.
    
    I consider this change low risk, since the functions it replaces are
    simple, very well specified, and well tested.
    
    For #16117.
    
    Change-Id: Ic39d851d5aca91330134596effd2dab9689ba066
    Reviewed-on: https://go-review.googlesource.com/24640
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 95483f262b619a53793baf86512aeabf44fc9d3a
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 29 19:20:51 2016 -0700

    os/exec: start checking for context cancelation in Start
    
    Previously we started checking for context cancelation in Wait, but
    that meant that when using StdoutPipe context cancelation never took
    effect.
    
    Fixes #16222.
    
    Change-Id: I89cd26d3499a6080bf1a07718ce38d825561899e
    Reviewed-on: https://go-review.googlesource.com/24650
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6c136493012b8a1f96f3edc9fa56aed70d34291a
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jun 30 08:22:27 2016 -0700

    syscall: accept more variants of id output when testing as root
    
    This should fix the report at #16224, and also fixes running the test as
    root on my Ubuntu Trusty system.
    
    Fixes #16224.
    
    Change-Id: I4e3b5527aa63366afb33a7e30efab088d34fb302
    Reviewed-on: https://go-review.googlesource.com/24670
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
```