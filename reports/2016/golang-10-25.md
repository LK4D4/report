# October 25, 2016 Report

Number of commits: 119

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 566.765601ms to 525.024808ms, -7.36%
* github.com/gogits/gogs: from 13.308492029s to 11.941720129s, -10.27%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 34.175010322s to 33.512432856s, -1.94%
* github.com/influxdata/influxdb/cmd/influxd: from 6.813920783s to 6.279078837s, -7.85%
* github.com/junegunn/fzf/src/fzf: from 1.042464591s to 960.074818ms, -7.90%
* github.com/mholt/caddy/caddy: from 6.00496195s to 5.956745536s, -0.80%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.427543585s to 1.356437815s, -4.98%
* github.com/prometheus/prometheus/cmd/prometheus: from 40.204993089s to 38.363351078s, -4.58%
* github.com/spf13/hugo: from 7.729125826s to 7.170379034s, -7.23%
* golang.org/x/tools/cmd/guru: from 2.798379465s to 2.4127721s, -13.78%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2538814 to 2538596, ~
* github.com/gogits/gogs: from 23390074 to 23234693, -0.66%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 29462592 to 29236896, -0.77%
* github.com/influxdata/influxdb/cmd/influxd: from 16392035 to 16347162, -0.27%
* github.com/junegunn/fzf/src/fzf: from 3087488 to 3088024, ~
* github.com/mholt/caddy/caddy: from 14810042 to 14778274, -0.21%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4297905 to 4286763, -0.26%
* github.com/prometheus/prometheus/cmd/prometheus: from 57021771 to 56844649, -0.31%
* github.com/spf13/hugo: from 16331085 to 16209558, -0.74%
* golang.org/x/tools/cmd/guru: from 7341299 to 7336332, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               190           188           -1.05%
BenchmarkMsgpUnmarshal-4             372           380           +2.15%
BenchmarkEasyJsonMarshal-4           1444          1440          -0.28%
BenchmarkEasyJsonUnmarshal-4         1595          1676          +5.08%
BenchmarkFlatBuffersMarshal-4        349           488           +39.83%
BenchmarkFlatBuffersUnmarshal-4      272           270           -0.74%
BenchmarkGogoprotobufMarshal-4       155           155           +0.00%
BenchmarkGogoprotobufUnmarshal-4     247           245           -0.81%

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

* [os: add ErrClosed, return for use of closed File](https://github.com/golang/go/commit/212d2f82e05018f1ebb5e40e2c328865201da356)
* [runtime: make sweep time proportional to in-use spans](https://github.com/golang/go/commit/f9497a6747abe8738728eeb08f80849c88404d18)
* [runtime: make markrootSpans time proportional to in-use spans](https://github.com/golang/go/commit/c95a8e458fdf9f3cb0c176ac92a513e5dc9b32c1)

## GIT Log

```
commit 2481481ff798636907376bfdf7e8c7558b8b930e
Author: Alexander Morozov <lk4d4math@gmail.com>
Date:   Tue Oct 25 14:37:54 2016 -0700

    runtime: fix comments in time.go
    
    Change-Id: I5c501f598f41241e6d7b21d98a126827a3c3ad9a
    Reviewed-on: https://go-review.googlesource.com/32018
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6e02750dd6747735e401e13b168bd6e03ab7dbe6
Author: Francesc Campoy <campoy@golang.org>
Date:   Tue Oct 25 19:09:36 2016 -0700

    cmd/dist: ignore stderr when listing packages to test
    
    Currently any warning will make dist fail because the
    text will be considered as part of the package list.
    
    Change-Id: I09a14089cd0448c3779e2f767e9356fe3325d8d9
    Reviewed-on: https://go-review.googlesource.com/32111
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit d89b70d43357e7f016331356014253647af02971
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 25 16:25:57 2016 -0700

    cmd/compile: slightly regularize interface method types
    
    Use a single *struct{} type instance rather than reconstructing one
    for every declared/imported interface method. Minor allocations win:
    
    name       old alloc/op    new alloc/op    delta
    Template      41.8MB ± 0%     41.7MB ± 0%  -0.10%         (p=0.000 n=9+10)
    Unicode       34.2MB ± 0%     34.2MB ± 0%    ~           (p=0.971 n=10+10)
    GoTypes        123MB ± 0%      122MB ± 0%  -0.03%         (p=0.000 n=9+10)
    Compiler       495MB ± 0%      495MB ± 0%  -0.01%        (p=0.000 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        409k ± 0%       408k ± 0%  -0.13%        (p=0.000 n=10+10)
    Unicode         354k ± 0%       354k ± 0%    ~           (p=0.516 n=10+10)
    GoTypes        1.22M ± 0%      1.22M ± 0%  -0.03%        (p=0.009 n=10+10)
    Compiler       4.43M ± 0%      4.43M ± 0%  -0.02%        (p=0.000 n=10+10)
    
    Change-Id: Id3a4ca3dd09112bb96ccc982b06c9e79f661d31f
    Reviewed-on: https://go-review.googlesource.com/32051
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit c78d072c8e4a8b0283da21423732f93c503711a4
Author: Keith Randall <khr@golang.org>
Date:   Tue Oct 25 23:34:35 2016 +0000

    Revert "Revert "cmd/compile: inline convI2E""
    
    This reverts commit 7dd9c385f6896c7dcb5d76353e52e36c81af2838.
    
    Reason for revert: Reverting the revert, which will re-enable the convI2E optimization.  We originally reverted the convI2E optimization because it was making the builder fail, but the underlying cause was later determined to be unrelated.
    
    Original CL: https://go-review.googlesource.com/31260
    Revert CL: https://go-review.googlesource.com/31310
    Real bug: https://go-review.googlesource.com/c/25159
    Real fix: https://go-review.googlesource.com/c/31316
    
    Change-Id: I17237bb577a23a7675a5caab970ccda71a4124f2
    Reviewed-on: https://go-review.googlesource.com/32023
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c60d9a33bfd4af38399b4caf76be0ced4c64c839
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Oct 18 14:56:19 2016 -0700

    net/http: fix redirect logic to handle mutations of cookies
    
    In the situation where the Client.Jar is set and the Request.Header
    has cookies manually inserted, the redirect logic needs to be
    able to apply changes to cookies from "Set-Cookie" headers to both
    the Jar and the manually inserted Header cookies.
    
    Since Header cookies lack information about the original domain
    and path, the logic in this CL simply removes cookies from the
    initial Header if any subsequent "Set-Cookie" matches. Thus,
    in the event of cookie conflicts, the logic preserves the behavior
    prior to change made in golang.org/cl/28930.
    
    Fixes #17494
    Updates #4800
    
    Change-Id: I645194d9f97ff4d95bd07ca36de1d6cdf2f32429
    Reviewed-on: https://go-review.googlesource.com/31435
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 70d685dc7244d46b3c22c4ac9588e51d76087ded
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 25 15:43:05 2016 -0700

    cmd/compile: don't wrap numeric or type literals in OPAREN
    
    It's only necessary to wrap named OTYPE or OLITERAL nodes, because
    their line numbers reflect the line number of the declaration, rather
    than use.
    
    Saves a lot of wrapper nodes in composite-literal-heavy packages like
    Unicode.
    
    name       old alloc/op    new alloc/op    delta
    Template      41.8MB ± 0%     41.8MB ± 0%  -0.07%        (p=0.000 n=10+10)
    Unicode       36.6MB ± 0%     34.2MB ± 0%  -6.55%        (p=0.000 n=10+10)
    GoTypes        123MB ± 0%      123MB ± 0%  -0.02%        (p=0.004 n=10+10)
    Compiler       495MB ± 0%      495MB ± 0%  -0.03%        (p=0.000 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        409k ± 0%       409k ± 0%  -0.05%        (p=0.029 n=10+10)
    Unicode         371k ± 0%       354k ± 0%  -4.48%         (p=0.000 n=10+9)
    GoTypes        1.22M ± 0%      1.22M ± 0%    ~           (p=0.075 n=10+10)
    Compiler       4.44M ± 0%      4.44M ± 0%  -0.02%        (p=0.000 n=10+10)
    
    Change-Id: Id1183170835125c778fb41b7e76d06d5ecd4f7a1
    Reviewed-on: https://go-review.googlesource.com/32021
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 57df2f802f0417f08100ff8002f3b062e695e148
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 25 11:50:13 2016 -0700

    cmd/compile: remove old lexer and parser
    
    Change-Id: I7306d28930dc4538a3bee31ff5d22f3f40681ec5
    Reviewed-on: https://go-review.googlesource.com/32020
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 575b1dda4ec845aea6b5c475e9d748dc76d1bc06
Author: Austin Clements <austin@google.com>
Date:   Wed Oct 5 21:22:33 2016 -0400

    runtime: eliminate allspans snapshot
    
    Now that sweeping and span marking use the sweep list, there's no need
    for the work.spans snapshot of the allspans list. This change
    eliminates the few remaining uses of it, which are either dead code or
    can use allspans directly, and removes work.spans and its support
    functions.
    
    Change-Id: Id5388b42b1e68e8baee853d8eafb8bb4ff95bb43
    Reviewed-on: https://go-review.googlesource.com/30537
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit c95a8e458fdf9f3cb0c176ac92a513e5dc9b32c1
Author: Austin Clements <austin@google.com>
Date:   Wed Oct 5 18:32:21 2016 -0400

    runtime: make markrootSpans time proportional to in-use spans
    
    Currently markrootSpans iterates over all spans ever allocated to find
    the in-use spans. Since we now have a list of in-use spans, change it
    to iterate over that instead.
    
    This, combined with the previous change, fixes #9265. Before these two
    changes, blowing up the heap to 8GB and then shrinking it to a 0MB
    live set caused the small-heap portion of the test to run 60x slower
    than without the initial blowup. With these two changes, the time is
    indistinguishable.
    
    No significant effect on other benchmarks.
    
    Change-Id: I4a27e533efecfb5d18cba3a87c0181a81d0ddc1e
    Reviewed-on: https://go-review.googlesource.com/30536
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit f9497a6747abe8738728eeb08f80849c88404d18
Author: Austin Clements <austin@google.com>
Date:   Wed Oct 5 17:50:39 2016 -0400

    runtime: make sweep time proportional to in-use spans
    
    Currently sweeping walks the list of all spans, which means the work
    in sweeping is proportional to the maximum number of spans ever used.
    If the heap was once large but is now small, this causes an
    amortization failure: on a small heap, GCs happen frequently, but a
    full sweep still has to happen in each GC cycle, which means we spent
    a lot of time in sweeping.
    
    Fix this by creating a separate list consisting of just the in-use
    spans to be swept, so sweeping is proportional to the number of in-use
    spans (which is proportional to the live heap). Specifically, we
    create two lists: a list of unswept in-use spans and a list of swept
    in-use spans. At the start of the sweep cycle, the swept list becomes
    the unswept list and the new swept list is empty. Allocating a new
    in-use span adds it to the swept list. Sweeping moves spans from the
    unswept list to the swept list.
    
    This fixes the amortization problem because a shrinking heap moves
    spans off the unswept list without adding them to the swept list,
    reducing the time required by the next sweep cycle.
    
    Updates #9265. This fix eliminates almost all of the time spent in
    sweepone; however, markrootSpans has essentially the same bug, so now
    the test program from this issue spends all of its time in
    markrootSpans.
    
    No significant effect on other benchmarks.
    
    Change-Id: Ib382e82790aad907da1c127e62b3ab45d7a4ac1e
    Reviewed-on: https://go-review.googlesource.com/30535
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 45baff61e37cb8ac497ca395d8da3f3e87601bb2
Author: Austin Clements <austin@google.com>
Date:   Wed Oct 5 12:31:00 2016 -0400

    runtime: expand comment on work.spans
    
    Change-Id: I4b8a6f5d9bc5aba16026d17f99f3512dacde8d2d
    Reviewed-on: https://go-review.googlesource.com/30534
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 5915ce66742415d96d5a082c41d31f965b719f26
Author: Austin Clements <austin@google.com>
Date:   Tue Oct 4 16:22:41 2016 -0400

    runtime: use len(h.spans) to indicate mapped region
    
    Currently we set the len and cap of h.spans to the full reserved
    region of the address space and track the actual mapped region
    separately in h.spans_mapped. Since we have both the len and cap at
    our disposal, change things so len(h.spans) tracks how much of the
    spans array is mapped and eliminate h.spans_mapped. This simplifies
    mheap and means we'll get nice "index out of bounds" exceptions if we
    do try to go off the end of the spans rather than a SIGSEGV.
    
    Change-Id: I8ed9a1a9a844d90e9fd2e269add4704623dbdfe6
    Reviewed-on: https://go-review.googlesource.com/30533
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 6b0f668044b7a8dd5d2360426cbcf6bd6785374d
Author: Austin Clements <austin@google.com>
Date:   Tue Oct 4 16:03:00 2016 -0400

    runtime: consolidate h_spans and mheap_.spans
    
    Like h_allspans and mheap_.allspans, these were two ways of referring
    to the spans array from when the runtime was split between C and Go.
    Clean this up by making mheap_.spans a slice and eliminating h_spans.
    
    Change-Id: I3aa7038d53c3a4252050aa33e468c48dfed0b70e
    Reviewed-on: https://go-review.googlesource.com/30532
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 66e849b168eef36a4159a4b038fe89eecd2f22e3
Author: Austin Clements <austin@google.com>
Date:   Tue Oct 4 15:56:19 2016 -0400

    runtime: eliminate mheap.nspan and use range loops
    
    This was necessary in the C days when allspans was an mspan**, but now
    that allspans is a Go slice, this is redundant with len(allspans) and
    we can use range loops over allspans.
    
    Change-Id: Ie1dc39611e574e29a896e01690582933f4c5be7e
    Reviewed-on: https://go-review.googlesource.com/30531
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 4d6207790b2f08daa00d2a7a67854a159ab2f601
Author: Austin Clements <austin@google.com>
Date:   Tue Oct 4 15:51:31 2016 -0400

    runtime: consolidate h_allspans and mheap_.allspans
    
    These are two ways to refer to the allspans array that hark back to
    when the runtime was split between C and Go. Clean this up by making
    mheap_.allspans a slice and eliminating h_allspans.
    
    Change-Id: Ic9360d040cf3eb590b5dfbab0b82e8ace8525610
    Reviewed-on: https://go-review.googlesource.com/30530
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit adda7ad29551d0880df1805ae22401551b1fbfa8
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 16 16:33:05 2016 -0700

    cmd/compile/internal/gc: enable new parser by default
    
    Change-Id: I3c784986755cfbbe1b8eb8da4d64227bd109a3b0
    Reviewed-on: https://go-review.googlesource.com/27203
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 263a825b0534a55916d472ba17feff7d1ed47569
Author: David du Colombier <0intro@gmail.com>
Date:   Wed Oct 26 00:07:27 2016 +0200

    syscall: use name+(NN)FP on plan9/amd64
    
    Generated from go vet.
    
    Change-Id: Ie775c29b505166e0bd511826ef20eeb153a0424c
    Reviewed-on: https://go-review.googlesource.com/32071
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 71f72e9d4f1cd3fd7e27515d1a5c5159fca99f0e
Author: David du Colombier <0intro@gmail.com>
Date:   Wed Oct 26 00:06:45 2016 +0200

    syscall: use name+(NN)FP on plan9/386
    
    Generated from go vet.
    
    Change-Id: I2620e5544be46485a876c7dce26b0592bf5a4101
    Reviewed-on: https://go-review.googlesource.com/32070
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f6aec889e1c880316b1989bdc6ce3b926cbe5fe4
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Oct 13 06:57:00 2016 -0400

    cmd/compile: add a writebarrier phase in SSA
    
    When the compiler insert write barriers, the frontend makes
    conservative decisions at an early stage. This may have false
    positives which result in write barriers for stack writes.
    
    A new phase, writebarrier, is added to the SSA backend, to delay
    the decision and eliminate false positives. The frontend still
    makes conservative decisions. When building SSA, instead of
    emitting runtime calls directly, it emits WB ops (StoreWB,
    MoveWB, etc.), which will be expanded to branches and runtime
    calls in writebarrier phase. Writes to static locations on stack
    are detected and write barriers are removed.
    
    All write barriers of stack writes found by the script from
    issue #17330 are eliminated (except two false positives).
    
    Fixes #17330.
    
    Change-Id: I9bd66333da9d0ceb64dcaa3c6f33502798d1a0f8
    Reviewed-on: https://go-review.googlesource.com/31131
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 698bfa17a842890043098b972446e9b8dbc20841
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Oct 23 22:32:16 2016 -0400

    cmd/internal/obj: save link register in leaf function with non-empty frame on PPC64, ARM64, S390X
    
    The runtime traceback code assumes non-empty frame has link
    link register saved on LR architectures. Make sure it is so in
    the assember.
    
    Also make sure that LR is stored before update SP, so the traceback
    code will not see a half-updated stack frame if a signal comes
    during the execution of function prologue.
    
    Fixes #17381.
    
    Change-Id: I668b04501999b7f9b080275a2d1f8a57029cbbb3
    Reviewed-on: https://go-review.googlesource.com/31760
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit cf73bbfa259afe29962a5ca5e207441f63c9bcf2
Author: Tom Bergan <tombergan@google.com>
Date:   Tue Oct 25 12:18:59 2016 -0700

    net/http: add an interface for server push
    
    This interface will be implemented by golang.org/x/net/http2 in
    https://go-review.googlesource.com/c/29439/.
    
    Updates golang/go#13443
    
    Change-Id: Ib6bdd403b0878cfe36fa9875c07c2c7239232556
    Reviewed-on: https://go-review.googlesource.com/32012
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 45b43f6198436b18a0b5f64a3176d632e4f1c855
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Tue Oct 25 15:13:07 2016 -0500

    cmd/objdump: updates from golang.org/x/arch/ppc64/ppc64asm
    
    Update the ppc64x disassembly code for use by objdump
    from golang.org/x/arch/ppc64/ppc64asm commit fcea5ea.
    Enable the objdump testcase for external linking on ppc64le
    make a minor fix to the expected output.
    
    Fixes #17447
    
    Change-Id: I769cc7f8bfade594690a476dfe77ab33677ac03b
    Reviewed-on: https://go-review.googlesource.com/32015
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a850dbdef2f1875d81ad09024480f648ce3eac32
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 25 15:49:00 2016 -0400

    cmd/vet: accept space-separated tag lists for compatibility with cmd/go
    
    Fixes #17148.
    
    Change-Id: I4c66aa0733c249ee6019d1c4e802a7e30457d4b6
    Reviewed-on: https://go-review.googlesource.com/32030
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit e24ccfc6fc6289073cdbe6f47f8a915e798578e9
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Oct 25 12:59:36 2016 -0700

    misc/cgo/errors: fix malloc test for dragonfly
    
    The Dragonfly libc returns a non-zero value for malloc(-1).
    
    Fixes #17585.
    
    Change-Id: Icfe68011ccbc75c676273ee3c3efdf24a520a004
    Reviewed-on: https://go-review.googlesource.com/32050
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d391dc260a71e9f67530d7c2ca20d595514d3514
Author: shaharko <skohanim@gmail.com>
Date:   Mon Oct 24 23:15:41 2016 +0300

    cmd/internal/obj: Use bitfield for LSym attributes
    
    Reduces the size of LSym struct.
    
    On 32bit: before 84  after 76
    On 64bit: before 136 after 128
    
    name       old time/op     new time/op     delta
    Template       182ms ± 3%      182ms ± 3%    ~           (p=0.607 n=19+20)
    Unicode       93.5ms ± 4%     94.2ms ± 3%    ~           (p=0.141 n=20+19)
    GoTypes        608ms ± 1%      605ms ± 2%    ~           (p=0.056 n=20+20)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        249M ± 7%       249M ± 4%    ~           (p=0.605 n=18+19)
    Unicode         149M ±14%       151M ± 5%    ~           (p=0.724 n=20+17)
    GoTypes         855M ± 4%       853M ± 3%    ~           (p=0.537 n=19+19)
    
    name       old alloc/op    new alloc/op    delta
    Template      40.3MB ± 0%     40.3MB ± 0%  -0.11%        (p=0.000 n=19+20)
    Unicode       33.8MB ± 0%     33.8MB ± 0%  -0.08%        (p=0.000 n=20+20)
    GoTypes        119MB ± 0%      119MB ± 0%  -0.10%        (p=0.000 n=19+20)
    
    name       old allocs/op   new allocs/op   delta
    Template        383k ± 0%       383k ± 0%    ~           (p=0.703 n=20+20)
    Unicode         317k ± 0%       317k ± 0%    ~           (p=0.982 n=19+18)
    GoTypes        1.14M ± 0%      1.14M ± 0%    ~           (p=0.086 n=20+20)
    
    Change-Id: Id6ba0db3ecc4503a4e9af3ed0d5884d4366e8bf9
    Reviewed-on: https://go-review.googlesource.com/31870
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Shahar Kohanim <skohanim@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2ee82edfc2c904c952daef6f442c223b1568cb66
Author: Rob Pike <r@golang.org>
Date:   Mon Oct 24 12:38:06 2016 -0700

    cmd/doc: show documentation for interface methods when requested explicitly
    
    For historical reasons, the go/doc package does not include
    the methods within an interface as part of the documented
    methods for that type. Thus,
    
            go doc ast.Node.Pos
    
    gives an incorrect and confusing error message:
    
            doc: no method Node.Pos in package go/ast
    
    This CL does some dirty work to dig down to the methods
    so interface methods now present their documentation:
    
    % go doc ast.node.pos
    func Pos() token.Pos  // position of first character belonging to the node
    %
    
    It must largely sidestep the doc package to do this, which
    is a shame. Perhaps things will improve there one day.
    
    The change does not handle embeddings, and in principle the
    same approach could be done for struct fields, but that is also
    not here yet. But this CL fixes the thing that was bugging me.
    
    Change-Id: Ic10a91936da96f54ee0b2f4a4fe4a8c9b93a5b4a
    Reviewed-on: https://go-review.googlesource.com/31852
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 050f378085da91ce65d2c0157b9bebcced5f883f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Oct 25 11:45:26 2016 -0700

    cmd/go: add more env variables to "go bug"
    
    CL 31330 added more envvars to "go env".
    This CL brings them to "go bug" as well.
    
    Change-Id: Iae122072c8178007eda8b765aaa3f38c3c6e39a0
    Reviewed-on: https://go-review.googlesource.com/32011
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d8d445280a6a7e5f3535e8b49e9ae248f2627759
Author: shaharko <skohanim@gmail.com>
Date:   Wed Oct 19 07:33:16 2016 +0300

    cmd/compile, cmd/link: more efficient typelink generation
    
    Instead of generating typelink symbols in the compiler
    mark types that should have typelinks with a flag.
    The linker detects this flag and adds the marked types
    to the typelink table.
    
    name            old s/op    new s/op    delta
    LinkCmdCompile   0.27 ± 6%   0.25 ± 6%  -6.93%    (p=0.000 n=97+98)
    LinkCmdGo        0.30 ± 5%   0.29 ±10%  -4.22%    (p=0.000 n=97+99)
    
    name            old MaxRSS  new MaxRSS  delta
    LinkCmdCompile   112k ± 3%   106k ± 2%  -4.85%  (p=0.000 n=100+100)
    LinkCmdGo        107k ± 3%   103k ± 3%  -3.00%  (p=0.000 n=100+100)
    
    Change-Id: Ic95dd4b0101e90c1fa262c9c6c03a2028d6b3623
    Reviewed-on: https://go-review.googlesource.com/31772
    Run-TryBot: Shahar Kohanim <skohanim@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 5a9549260df1f5ffcbdd5938861fea9f74478661
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Tue Oct 25 19:33:50 2016 +0530

    math/cmplx: prevent infinite loop in tanSeries
    
    The condition to determine if any further iterations are needed is
    evaluated to false in case it encounters a NaN. Instead, flip the
    condition to keep looping until the factor is greater than the machine
    roundoff error.
    
    Updates #17577
    
    Change-Id: I058abe73fcd49d3ae4e2f7b33020437cc8f290c3
    Reviewed-on: https://go-review.googlesource.com/31952
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit a2f77e9ef8ba2e956453ad0dda1ebdf4ae7c4fdb
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Oct 19 13:13:31 2016 -0700

    cmd/compile: cleanup gdata slightly
    
    In sinit.go, gdata can already handle strings and complex, so no
    reason to handle them separately.
    
    In obj.go, inline gdatastring and gdatacomplex into gdata, since it's
    the only caller. Allows extracting out the common Linksym calls.
    
    Passes toolstash -cmp.
    
    Change-Id: I3cb18d9b4206a8a269c36e0d30a345d8e6caba1f
    Reviewed-on: https://go-review.googlesource.com/31498
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 213ee3d20ed35bf1bce8bb3e93fd7b0fca562536
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 25 08:39:19 2016 -0700

    go/types: match cmd/compile's alignment for complex64
    
    Fixes #17584.
    
    Change-Id: I3af31cc1f2e9c906f3b73e77f3b092624ba78fbe
    Reviewed-on: https://go-review.googlesource.com/31939
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 1986a450ddd2e5ad213b9aff92cf7fe495234bbf
Author: David Chase <drchase@google.com>
Date:   Tue Oct 25 11:11:40 2016 -0400

    cmd/compile: added test to ensure that accidental fix remains
    
    Bug 15141 was apparently fixed by some other change to the
    compiler (this is plausible, it was a weird bug dependent
    on a particular way of returning a large named array result),
    add the test to ensure that it stays fixed.
    
    Updates #15141.
    
    Change-Id: I3d6937556413fab1af31c5a1940e6931563ce2f3
    Reviewed-on: https://go-review.googlesource.com/31972
    Run-TryBot: David Chase <drchase@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d8f7e4fadb785d90444245159f4c5a32653263d6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Oct 25 06:44:58 2016 -0700

    runtime, syscall: appease vet
    
    No functional changes.
    
    Change-Id: I0842b2560f4296abfc453410fdd79514132cab83
    Reviewed-on: https://go-review.googlesource.com/31935
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit eb6ae3ca7b0bb9f4bbeac03cc3cbd2362fc05342
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Oct 25 05:56:03 2016 -0700

    cmd/vet/all: remove cmd/compile/internal/big special case
    
    It no longer exists as of CL 31010.
    
    Change-Id: Idd61f392544cad8b3f3f8d984dc5c953b473e2e5
    Reviewed-on: https://go-review.googlesource.com/31934
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b10b2f8d407198d59525197e0b94a16fcc4a7cf0
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Mon Oct 24 14:40:42 2016 -0500

    cmd/internal: add shift opcodes with shift operands on ppc64x
    
    Some original shift opcodes for ppc64x expected an operand to be
    a mask instead of a shift count, preventing some valid shift counts
    from being written.
    
    This adds new opcodes for shifts where needed, using mnemonics that
    match the ppc64 asm and allowing the assembler to accept the full set
    of valid shift counts.
    
    Fixes #15016
    
    Change-Id: Id573489f852038d06def279c13fd0523736878a7
    Reviewed-on: https://go-review.googlesource.com/31853
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
    Reviewed-by: David Chase <drchase@google.com>

commit 5db7c6d32cef3d9dd917adb81c4a38978853a17b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Oct 25 06:46:26 2016 -0700

    cmd/vet/all: update whitelists
    
    Change-Id: Ie505b5d8cdfe4ffda71f909d6f81603b6d752eed
    Reviewed-on: https://go-review.googlesource.com/31937
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit cf09920c0f9dcb8ef9c85cf02e51e876e52845e2
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Oct 25 06:46:51 2016 -0700

    cmd/compile: place OIDATA next to OITAB
    
    Change-Id: Ia499125714e272af87562de5e5d23e68a112df58
    Reviewed-on: https://go-review.googlesource.com/31938
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 79b5d329011823c2dff403a98e5a6a4f5946cde2
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Oct 25 06:45:35 2016 -0700

    crypto/tls: fix vet issues again
    
    While we're here, use test[%d] in place of #%d.
    
    Change-Id: Ie30afcab9673e78d3ea7ca80f5e662fbea897488
    Reviewed-on: https://go-review.googlesource.com/31936
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1abefc1ff0ae6a3f3df7affa704cd2c71ab48a05
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Oct 25 05:45:52 2016 -0700

    cmd/compile: clean up rule logging helpers
    
    Introduced in CLs 29380 and 30011.
    
    Change-Id: I3d3641e8748ce0adb57b087a1fcd62f295ade665
    Reviewed-on: https://go-review.googlesource.com/31933
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit a3faa80033ae2ffa3ee56439759f0ea0200a3a3e
Author: Russ Cox <rsc@golang.org>
Date:   Fri Oct 21 09:21:06 2016 -0400

    cmd/go: bypass install to os.DevNull entirely, test mayberemovefile(os.DevNull)
    
    Fixes #16811.
    
    Change-Id: I7d018015f691838482ccf845d621209b96935ba4
    Reviewed-on: https://go-review.googlesource.com/31657
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit 3ef07c412f068144554648c0d209bef444a2ee27
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Oct 23 16:11:13 2016 -0400

    cmd, runtime: remove s390x 3 operand immediate logical ops
    
    These are emulated by the assembler and we don't need them.
    
    Change-Id: I2b07c5315a5b642fdb5e50b468453260ae121164
    Reviewed-on: https://go-review.googlesource.com/31758
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 517a44d57e6a0834e98b763045fab9be8f4673d0
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Oct 19 16:41:01 2016 -0400

    cmd/compile: intrinsify atomic operations on s390x
    
    Implements the following intrinsics on s390x:
     - AtomicAdd{32,64}
     - AtomicCompareAndSwap{32,64}
     - AtomicExchange{32,64}
     - AtomicLoad{32,64,Ptr}
     - AtomicStore{32,64,PtrNoWB}
    
    I haven't added rules for And8 or Or8 yet.
    
    Change-Id: I647af023a8e513718e90e98a60191e7af6167314
    Reviewed-on: https://go-review.googlesource.com/31614
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2113c9ad0d82b3d1a734c0b5fc0efc9c44a920d5
Author: Martin Möhrmann <martisch@uos.de>
Date:   Tue Oct 18 06:18:39 2016 +0200

    image/color: improve speed of RGBA methods
    
    Apply the optimizations added to color conversion functions in
    https://go-review.googlesource.com/#/c/21910/ to the RGBA methods.
    
    YCbCrToRGBA/0-4      6.32ns ± 3%  6.58ns ± 2%   +4.15%  (p=0.000 n=20+19)
    YCbCrToRGBA/128-4    8.02ns ± 2%  5.89ns ± 2%  -26.57%  (p=0.000 n=20+19)
    YCbCrToRGBA/255-4    8.06ns ± 2%  6.59ns ± 3%  -18.18%  (p=0.000 n=20+20)
    NYCbCrAToRGBA/0-4    8.71ns ± 2%  8.78ns ± 2%   +0.86%  (p=0.036 n=19+20)
    NYCbCrAToRGBA/128-4  10.3ns ± 4%   7.9ns ± 2%  -23.44%  (p=0.000 n=20+20)
    NYCbCrAToRGBA/255-4  9.64ns ± 2%  8.79ns ± 3%   -8.80%  (p=0.000 n=20+20)
    
    Fixes: #15260
    
    Change-Id: I225efdf74603e8d2b4f063054f7baee7a5029de6
    Reviewed-on: https://go-review.googlesource.com/31773
    Run-TryBot: Martin Möhrmann <martisch@uos.de>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Nigel Tao <nigeltao@golang.org>

commit 80a034642e9eac3fde60dbff920a70c4aae7fbc9
Author: shaharko <skohanim@gmail.com>
Date:   Thu Oct 13 22:31:46 2016 +0300

    cmd/compile, cmd/link: stop generating unused go.string.hdr symbols.
    
    name       old s/op    new s/op    delta
    LinkCmdGo   0.29 ± 5%   0.29 ± 8%  -2.60%   (p=0.000 n=97+98)
    
    name       old MaxRSS  new MaxRSS  delta
    LinkCmdGo   106k ± 4%   105k ± 3%  -1.00%  (p=0.000 n=100+99)
    
    Change-Id: I75a1c3b24ea711a15a5d2eae026b70b97ee7bad4
    Reviewed-on: https://go-review.googlesource.com/31030
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 001743813971e0870144443399f2fb396b3781b7
Author: shaharko <skohanim@gmail.com>
Date:   Thu Oct 13 22:56:38 2016 +0300

    cmd/compile: only generate ·f symbols when necessary
    
    Before go supported buildmode=shared ·f symbols used to be defined
    only when they were used. In order to solve #11480 the strategy
    was changed to have these symbols defined on declaration which is
    less efficient and generates many unneeded symbols.
    With this change the best strategy is chosen for each situation,
    improving static linking time:
    
    name            old s/op    new s/op    delta
    LinkCmdCompile   0.27 ± 5%   0.25 ± 6%  -8.22%   (p=0.000 n=98+96)
    LinkCmdGo        0.30 ± 6%   0.29 ± 8%  -5.03%   (p=0.000 n=95+99)
    
    name            old MaxRSS  new MaxRSS  delta
    LinkCmdCompile   107k ± 2%    98k ± 3%  -8.32%  (p=0.000 n=99+100)
    LinkCmdGo        106k ± 3%   104k ± 3%  -1.94%  (p=0.000 n=99+100)
    
    Change-Id: I965eeee30541e724fd363804adcd6fda10f965a4
    Reviewed-on: https://go-review.googlesource.com/31031
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 5d8324e6822e34c900e36d67adf640cee6693d25
Author: Paul Marks <pmarks@google.com>
Date:   Mon Oct 24 17:49:22 2016 -0700

    net: add hostname warnings to all first(isIPv4) functions.
    
    In general, these functions cannot behave correctly when given a
    hostname, because a hostname may represent multiple IP addresses, and
    first(isIPv4) chooses at most one.
    
    Updates #9334
    
    Change-Id: Icfb629f84af4d976476385a3071270253c0000b1
    Reviewed-on: https://go-review.googlesource.com/31931
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 71cf409dbdcba20f220d65ab83a6494c1f79b2a0
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 24 20:47:31 2016 -0400

    runtime: accept timeout from non-timeout semaphore wait on OS X
    
    Looking at the kernel sources, I don't see how this is possible.
    But obviously it is. Just try again.
    
    Fixes #17161.
    
    Change-Id: Iea7d53f7cf75944792d2f75a0d07129831c7bcdb
    Reviewed-on: https://go-review.googlesource.com/31823
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b4ce38ec5769a270f0545dce43b2e926230609c3
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Oct 24 10:19:04 2016 -0700

    cmd/cgo: throw if C.malloc returns nil
    
    Change-Id: If7740ac7b6c4190db5a1ab4100d12cf16dc79c84
    Reviewed-on: https://go-review.googlesource.com/31768
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 643c6b3c74409871d7f96cbad145dabdefbbc8c1
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Mon Oct 24 08:52:04 2016 +0900

    path/filepath: make TestToNorm robust
    
    The old code leaves garbages in a temporary directory because it
    cannot remove the current working directory on windows.
    The new code changes the directory before calling os.Remove.
    
    Furthermore, the old code assumes that ioutil.TempDir (os.TempDir)
    doesn't return a relative path nor an UNC path.
    If it isn't the case, the new code calls t.Fatal earlier for preventing
    ambiguous errors.
    
    Finally, the old code reassigns the variable which is used by the defer
    function. It could cause unexpected results, so avoid that.
    
    Change-Id: I5fc3902059ecaf18dc1341ecc4979d1206034cd7
    Reviewed-on: https://go-review.googlesource.com/31790
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 02c1d8a1589b47f12deaaa63d4a6a084640f4389
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Mon Oct 24 15:06:20 2016 +1100

    cmd/link/internal/ld: remove goto from ldpe.go
    
    Updates #15345
    
    Change-Id: I447d133512e99a9900928a910e161a85db6e8b75
    Reviewed-on: https://go-review.googlesource.com/31792
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c0e2318f7c9a83b237a08b208eb145d520e3a233
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Oct 21 17:23:01 2016 -0700

    cmd/compile: simplify parsing of type aliases
    
    Change-Id: Ia86841cf84bc17ff6ecc6e5ac4cec86384a0da00
    Reviewed-on: https://go-review.googlesource.com/31719
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 1b0cf430dd130ad53e9f43bc04d2ed91bcd87b26
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Oct 19 15:32:14 2016 -0700

    cmd/compile: implement package-level aliases (no export yet)
    
    Requires -newparser=1.
    
    For #17487.
    For #16339.
    
    Change-Id: I156fb0c0f8a97e8c72dbbfbd7fe821efee12b957
    Reviewed-on: https://go-review.googlesource.com/31597
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 7e9f420ddfb1ce8882bb715158cdb8b977b93955
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Oct 24 15:56:48 2016 -0700

    test: delete bugs directory
    
    It appears to be a vestigial holding ground for bugs.
    But we have an issue tracker, and #1909 is there and open.
    
    Change-Id: I912ff222a24c51fab483be0c67dad534f5a84488
    Reviewed-on: https://go-review.googlesource.com/31859
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c63db157bd669845a23255ab268596669a5ea5df
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Oct 25 00:31:45 2016 +0200

    net: handle "dns failure" as errNoSuchHost on Plan 9
    
    CL 31468 added TestLookupNonLDH, which was failing on Plan 9,
    because LookupHost was expecting to return errNoSuchHost
    on DNS resolution failure, while Plan 9 returned the
    "dns failure" string.
    
    In the Plan 9 implementation of lookupHost, we now return
    errNoSuchHost instead of the "dns failure" string, so
    the behavior is more consistant with other operating systems.
    
    Fixes #17568.
    
    Change-Id: If64f580dc0626a4a4f19e5511ba2ca5daff5f789
    Reviewed-on: https://go-review.googlesource.com/31873
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6e78f7697405e8c24af796322e4c3c325ed97d95
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Oct 24 23:24:11 2016 +0000

    testing/quick, text/tabwriter: freeze packages
    
    Fixes #15557
    
    Change-Id: I02ad98068894e75d4e08e271fdd16cb420519460
    Reviewed-on: https://go-review.googlesource.com/31910
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 4c9c023346415908390b68cfab0677ef53cf38ac
Author: Alexander Döring <email@alexd.ch>
Date:   Mon Oct 24 22:40:31 2016 +0200

    math,math/cmplx: fix linter issues
    
    Change-Id: If061f1f120573cb109d97fa40806e160603cd593
    Reviewed-on: https://go-review.googlesource.com/31871
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 426c287eb64abfe54e51fba3014c8eaeff7fc270
Author: Joshua Boelter <joshua.boelter@intel.com>
Date:   Wed Jul 13 16:22:28 2016 -0600

    crypto/tls: add VerifyPeerCertificate to tls.Config
    
    VerifyPeerCertificate returns an error if the peer should not be
    trusted. It will be called after the initial handshake and before
    any other verification checks on the cert or chain are performed.
    This provides the callee an opportunity to augment the certificate
    verification.
    
    If VerifyPeerCertificate is not nil and returns an error,
    then the handshake will fail.
    
    Fixes #16363
    
    Change-Id: I6a22f199f0e81b6f5d5f37c54d85ab878216bb22
    Reviewed-on: https://go-review.googlesource.com/26654
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 93bca632d9d4662c0e04f6ae24122579130a4bc2
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Oct 24 15:36:26 2016 -0700

    cmd/compile: preserve type information in inrange
    
    Fixes #17551.
    
    Change-Id: I84b7d82654cda3559c119aa56b07f30d0d224865
    Reviewed-on: https://go-review.googlesource.com/31857
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b55cee1893283cc55e99fb041fc0067f56924926
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Oct 21 18:22:13 2016 -0400

    cmd/internal/obj/mips: store LR before update SP in function prologue
    
    This prevents the traceback code from seeing a half-updated
    stack frame when a profiling signal comes during the execution
    of function prologue. Also fixes mips64x part of #17381.
    
    Change-Id: Iec9683427e546e3648b2e8b1dde956d13f6eb938
    Reviewed-on: https://go-review.googlesource.com/31721
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 0c02280fe0b4a178f78633721201f899e5213414
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Oct 25 00:12:57 2016 +0200

    net: fix TestCloseError on Plan 9
    
    Since CL 30614, TestCloseError is failing on Plan 9,
    because File.Write now checks f.fd == badFd before
    calling syscall.Write.
    
    The f.fd == badFd check returns os.ErrClosed, while
    syscall.Write returned a syscall.ErrorString error.
    
    TestCloseError was failing because it expected a
    syscall.ErrorString error.
    
    We add a case in parseCloseError to handle the
    os.ErrClosed case.
    
    Fixes #17569.
    
    Change-Id: I6b4d956d18ed6d3c2ac5211ffd50a4888f7521e1
    Reviewed-on: https://go-review.googlesource.com/31872
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1fcad29341ba4daa8762f2ead932fd001a9d2476
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Oct 24 14:33:22 2016 -0700

    cmd/compile: remove OREGISTER, Node.Reg
    
    OREGISTER is unused.
    
    All remaining uses of Node.Reg use REGSP.
    
    Change-Id: I51cf06826867e576baabd568e04f96d2634f5cad
    Reviewed-on: https://go-review.googlesource.com/31856
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 28269796c8ef6c59cca746b090080d99043cdeef
Author: Russ Cox <rsc@golang.org>
Date:   Fri Oct 21 00:02:33 2016 -0400

    cmd/internal/objfile: remove debugging print
    
    Crept into CL 9682, committed last week.
    
    Change-Id: I5b8e9119dbfeb0bc3005623ab74dbd29311d17ae
    Reviewed-on: https://go-review.googlesource.com/31814
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit aeb8b9591c254560c1ecc4eafbc38b31d632da64
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 14:51:58 2016 -0400

    encoding/json: fix bad formatting introduced in CL 20356
    
    Change-Id: I39a8b543e472e5ec5d4807a9b7f61657465c5ce5
    Reviewed-on: https://go-review.googlesource.com/31816
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a7cad4110ac229ecf5b0fcf0a035c5a03c699415
Author: Michael Fraenkel <michael.fraenkel@gmail.com>
Date:   Sat Oct 8 06:57:49 2016 -0400

    net/http/httputil: log err encountered during reverseproxy body copying
    
    Fixes #16659
    
    Change-Id: I13dd797e93e0b572eaf8726f1be594870d40183b
    Reviewed-on: https://go-review.googlesource.com/30692
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 95f3e47456e42899b64d3740eab1dd7ee1db5bf9
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Mon Oct 24 14:19:18 2016 -0500

    cmd/compile: add rule to use ANDN for a&^b on ppc64x
    
    Adds a rule to generate ANDN for AND x ^y.
    
    Fixes #17567
    
    Change-Id: I3b978058d5663f32c42b1af19bb207eac5622615
    Reviewed-on: https://go-review.googlesource.com/31769
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 7124056f7e6f44faba822e4d96c18fde002b4566
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Oct 24 11:46:06 2016 -0700

    cmd/internal/obj: drop Addr's Gotype field
    
    The Gotype field is only used for ATYPE instructions. Instead of
    specially storing the Go type symbol in From.Gotype, just store it in
    To.Sym like any other 2-argument instruction would.
    
    Modest reduction in allocations:
    
    name       old alloc/op    new alloc/op    delta
    Template      42.0MB ± 0%     41.8MB ± 0%  -0.40%         (p=0.000 n=9+10)
    Unicode       34.3MB ± 0%     34.1MB ± 0%  -0.48%         (p=0.000 n=9+10)
    GoTypes        122MB ± 0%      122MB ± 0%  -0.14%         (p=0.000 n=9+10)
    Compiler       518MB ± 0%      518MB ± 0%  -0.04%         (p=0.000 n=9+10)
    
    Passes toolstash -cmp.
    
    Change-Id: I0e603266b5d7d4e405106a26369e22773a0d3a91
    Reviewed-on: https://go-review.googlesource.com/31850
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 07a22dbd34d439182bea6a966b80baa2df7c72f4
Author: Alan Donovan <adonovan@google.com>
Date:   Mon Oct 24 11:00:19 2016 -0400

    cmd/vet: cgo: emit no error for calls to C.CBytes
    
    Fixes issue golang/go#17563
    
    Change-Id: Ibb41ea9419907193526cc601f6afd07d8689b1fe
    Reviewed-on: https://go-review.googlesource.com/31810
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2a3db5c0175e26ba99adca90e5bab6cb4e85cc2d
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 22:22:25 2016 -0400

    cmd/cgo: document C.malloc behavior
    
    Fixes #16309.
    
    Change-Id: Ifcd28b0746e1af30e2519a7b118485aecfb12396
    Reviewed-on: https://go-review.googlesource.com/31811
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a6141ebd3f33c1f4330a0ce3507f686b7fd64489
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Mon Oct 24 22:39:04 2016 +0530

    math/big: fix alignment in Float.Parse docs
    
    Leading spaces in a couple of lines instead of tabs cause those to be
    misaligned (as seen on <https://golang.org/pkg/math/big/#Float.Parse>):
    
    <<<
            number   = [ sign ] [ prefix ] mantissa [ exponent ] | infinity .
            sign     = "+" | "-" .
         prefix   = "0" ( "x" | "X" | "b" | "B" ) .
            mantissa = digits | digits "." [ digits ] | "." digits .
            exponent = ( "E" | "e" | "p" ) [ sign ] digits .
            digits   = digit { digit } .
            digit    = "0" ... "9" | "a" ... "z" | "A" ... "Z" .
         infinity = [ sign ] ( "inf" | "Inf" ) .
    >>>
    
    Replace the leading spaces with tabs so that those align well.
    
    Change-Id: Ibba6cd53f340001bbd929067dc587feb071dc3bd
    Reviewed-on: https://go-review.googlesource.com/31830
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0fff67d191514d4eab9e5fd1f078a8c66c8550fd
Author: Alexander Döring <email@alexd.ch>
Date:   Sun Oct 23 23:13:59 2016 +0200

    database/sql: fix possible context leak in test
    
    Fixes #17560
    
    Change-Id: I96fcdec87220391ef5432571b5c090b5be27491a
    Reviewed-on: https://go-review.googlesource.com/31771
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit accf5cc3862830bb8bb54af31aa1861e6a2b8481
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Oct 23 14:10:11 2016 -0700

    all: minor vet fixes
    
    Change-Id: I22f0f3e792052762499f632571155768b4052bc9
    Reviewed-on: https://go-review.googlesource.com/31759
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 212d2f82e05018f1ebb5e40e2c328865201da356
Author: Dan Caddigan <goldcaddy77@gmail.com>
Date:   Fri Oct 7 00:46:56 2016 -0400

    os: add ErrClosed, return for use of closed File
    
    This is clearer than syscall.EBADF.
    
    Fixes #17320.
    
    Change-Id: I14c6a362f9a6044c9b07cd7965499f4a83d2a860
    Reviewed-on: https://go-review.googlesource.com/30614
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 452bbfc179d6739a404aacc819ec66acc71fc55c
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 19 00:01:31 2016 -0400

    path/filepath: fix match of \\?\c:\* on Windows
    
    \\?\c:\ is a "root directory" that is not subject to further matching,
    but the ? makes it look like a pattern, which was causing an
    infinite recursion. Make sure the code understands the ? is not a pattern.
    
    Fixes #15879.
    
    Change-Id: I3a4310bbc398bcae764b9f8859c875317345e757
    Reviewed-on: https://go-review.googlesource.com/31460
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit c5ccbdd22bdbdc43d541b7e7d4ed66ceb559030e
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 16:53:49 2016 -0400

    net/url: reject colon in first segment of relative path in Parse
    
    RFC 3986 §3.3 disallows relative URL paths in which the first segment
    contains a colon, presumably to avoid confusion with scheme:foo syntax,
    which is exactly what happened in #16822.
    
    Fixes #16822.
    
    Change-Id: Ie4449e1dd21c5e56e3b126e086c3a0b05da7ff24
    Reviewed-on: https://go-review.googlesource.com/31582
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit afe675c2fa9deacd61f5684be54f1ebbdc94fc0c
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 14:37:31 2016 -0400

    doc/faq: mention that copying discussion is semantics, not implementation
    
    Fixes #17181.
    
    Change-Id: If7cc4865e391acf76512f7ec7167d5a31377b598
    Reviewed-on: https://go-review.googlesource.com/31574
    Reviewed-by: Rob Pike <r@golang.org>

commit 39690beb5885d378d98117c3d57b494e97f16eea
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 09:11:20 2016 -0400

    runtime: fix invariant comment in chan.go
    
    Change-Id: Ic6317f186d0ee68ab1f2d15be9a966a152f61bfb
    Reviewed-on: https://go-review.googlesource.com/31610
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 2693fa15ee12acd67e45d8fa57626675903ab605
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 19 13:34:15 2016 -0400

    html/template: add test case for unbounded template expansion
    
    Fixed by CL 31092 already, but that change is a few steps away
    from the problem observed here, so add an explicit test.
    
    Fixes #17019.
    
    Change-Id: If4ece1418e6596b1976961347889ce12c5969637
    Reviewed-on: https://go-review.googlesource.com/31466
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit 604146ce8961d32f410949015fc8ee31f9052209
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 19 10:27:05 2016 -0400

    html/template, text/template: docs and fixes for template redefinition
    
    All prior versions of Go have allowed redefining empty templates
    to become non-empty. Unfortunately, that has never consistently
    taken effect in html/template after the first execution:
    
            // define and execute
            t := template.New("root")
            t.Parse(`{{define "T"}}{{end}}<a href="{{template "T"}}">`)
            t.Execute(w, nil) // <a href="">
    
            // redefine
            t.Parse(`{{define "T"}}my.url{{end}}`) // succeeds, but ignored
            t.Execute(w, nil) // <a href="">
    
    When Go 1.6 added {{block...}} to text/template, that loosened the
    redefinition rules to allow redefinition at any time. The loosening was
    undone a bit in html/template, although inconsistently:
    
            // define and execute
            t := template.New("root")
            t.Parse(`{{define "T"}}body{{end}}`)
            t.Lookup("T").Execute(ioutil.Discard, nil)
    
            // attempt to redefine
            t.Parse(`{{define "T"}}body{{end}}`) // rejected in all Go versions
            t.Lookup("T").Parse("body") // OK as of Go 1.6, likely unintentionally
    
    Like in the empty->non-empty case, whether future execution takes
    notice of a redefinition basically can't be explained without going into
    the details of the template escape analysis.
    
    Address both the original inconsistencies in whether a redefinition
    would have any effect and the new inconsistencies about whether a
    redefinition is allowed by adopting a new rule: no parsing or modifying
    any templates after the first execution of any template in the same set.
    Template analysis begins at first execution, and once template analysis
    has begun, we simply don't have the right logic to update the analysis
    for incremental modifications (and never have).
    
    If this new rule breaks existing uses of templates that we decide need
    to be supported, we can try to invalidate all escape analysis for the
    entire set after any modifications. But let's wait on that until we know
    we need to and why.
    
    Also fix documentation of text/template redefinition policy
    (redefinition is always OK).
    
    Fixes #15761.
    
    Change-Id: I7d58d7c08a7d9df2440ee0d651a5b2ecaff3006c
    Reviewed-on: https://go-review.googlesource.com/31464
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit fa90f9b909286dc815fde1f83f77b80bd686127d
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 19 16:48:21 2016 -0400

    net: there are no invalid domain names anymore
    
    The Go resolver reports invalid domain name for '!!!.local',
    but that is allowed by multicast DNS. In general we can't predict
    what future relaxations might come along, and libc resolvers
    do not distinguish 'no such host' from 'invalid name', so stop
    making that distinction here too. Always use 'no such host'.
    
    Fixes #12421.
    
    Change-Id: I8f22604767ec9e270434e483da52b337833bad71
    Reviewed-on: https://go-review.googlesource.com/31468
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 157ce90abe24e7ff8ac51b05660eb7190101138c
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 10:06:54 2016 -0400

    go/build: allow % in ${SRCDIR} expansion for Jenkins
    
    Fixes #16959.
    
    Change-Id: Ibbb28fdf26c53788a0edb3e3ea54ec030fa2a8cf
    Reviewed-on: https://go-review.googlesource.com/31611
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 86324f29c6892e56ea756e7a5d81a26c86ce2e12
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 15:45:51 2016 -0400

    go/build: do not record go:binary-only-package if build tags not satisfied
    
    This is the documented (and now implemented) behavior.
    
    Fixes #16841.
    
    Change-Id: Ic75adc5ba18303ed9578e04284f32933f905d6a3
    Reviewed-on: https://go-review.googlesource.com/31577
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit ee4b58df61e25bec2d526947dd76cd2ab5d3051d
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 15:51:26 2016 -0400

    log: document that log messages end in newlines
    
    Fixes #16564.
    
    Change-Id: Idd7b3c8f1d8415acd952d1efb6dc35ba4191805d
    Reviewed-on: https://go-review.googlesource.com/31578
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit fc88a0f4ce66704d04f2b6d7730e722ef2b9a5de
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 16:43:09 2016 -0400

    net/mail: expose ParseDate, for use parsing Resent-Date headers
    
    Fixes #16657.
    
    Change-Id: I9425af91a48016b1d7465b9f43cafa792bc00bb3
    Reviewed-on: https://go-review.googlesource.com/31581
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit 699fb0fc5b654b714d7e8be5628646c0b063d523
Author: Quentin Smith <quentin@golang.org>
Date:   Fri Oct 21 16:48:50 2016 -0400

    cmd/doc: continue searching after error reading directory
    
    If a directory in GOPATH is unreadable, we should keep looking for other
    packages. Otherwise we can give the misleading error "no buildable Go
    source files".
    
    Fixes #16240
    
    Change-Id: I38e1037f56ec463d3c141f0508fb74211cb90f13
    Reviewed-on: https://go-review.googlesource.com/31713
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 8419c85eaa014a6b8f3485a4e27520a3acd31601
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 20:37:28 2016 -0400

    runtime, cmd/link: fix netbsd/arm EABI support
    
    Fixes reported by oshimaya (see #13806).
    
    Fixes #13806.
    
    Change-Id: I9b659ab918a34bc5f7c58f3d7f59058115b7f776
    Reviewed-on: https://go-review.googlesource.com/31651
    Reviewed-by: Minux Ma <minux@golang.org>

commit 19adf8aeaae7c898e8efeab18e3162c3807a8756
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 20 20:05:41 2016 -0400

    reflect: fix DeepEqual for some cyclic corner cases
    
    Fixes #15610.
    
    Change-Id: Idbc8a9b328b92034d53b8009471678a166d5cf3f
    Reviewed-on: https://go-review.googlesource.com/31588
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit 17ad60b8a4b92ee36f5b14c609ad1d0b5805b886
Author: Russ Cox <rsc@golang.org>
Date:   Fri Oct 21 10:45:47 2016 -0400

    cmd/go: fix test for moved package in go get -u
    
    What matters during go get -u is not whether there is an import comment
    but whether we resolved the path by an HTML <meta> tag.
    
    Fixes #16471.
    
    Change-Id: I6b194a3f73a7962a0170b4d5cf51cfed74e02c00
    Reviewed-on: https://go-review.googlesource.com/31658
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit f3b4abd806c392ff0bf8566009e750ebc01ae355
Author: Russ Cox <rsc@golang.org>
Date:   Fri Oct 21 11:37:54 2016 -0400

    cmd/go: allow 'go generate' even if imports do not resolve
    
    Maybe the go generate is generating the imports,
    or maybe there's some other good reason the code
    is incomplete.
    
    The help text already says:
    
            Note that go generate does not parse the file, so lines that look
            like directives in comments or multiline strings will be treated
            as directives.
    
    We'll still reject Go source files that don't begin with a package statement
    or have a syntax error in the import block, but those are I think more
    defensible rejections.
    
    Fixes #16307.
    
    Change-Id: I4f8496c02fdff993f038adfed2df4db7f067dc06
    Reviewed-on: https://go-review.googlesource.com/31659
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 9575882461b79c833d206ace186db178e5aff9fa
Author: Russ Cox <rsc@golang.org>
Date:   Fri Oct 21 12:19:52 2016 -0400

    cmd/go: document that cmd/foo is only in Go repo, never GOPATH
    
    It's always been like this, so document it.
    
    Fixes #14351.
    
    Change-Id: Ic6a7c44881bac0209fa6863a487fabec5ec0214e
    Reviewed-on: https://go-review.googlesource.com/31663
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit a1813fcb839e21e46208a5294df87097a44bd05a
Author: Russ Cox <rsc@golang.org>
Date:   Fri Oct 21 15:01:03 2016 -0400

    cmd/go: referee another vendor vs symlink fight
    
    Avoid crash in the specific case reported in #15201 but also
    print more useful error message, avoiding slice panic.
    
    Fixes #15201.
    Fixes #16167.
    Fixes #16566.
    
    Change-Id: I66499621e9678a05bc9b12b0da77906cd7027bdd
    Reviewed-on: https://go-review.googlesource.com/31665
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit 0eaa8fe03f46d9cc35d81e872b4b23442609ed80
Author: Quentin Smith <quentin@golang.org>
Date:   Fri Oct 21 17:30:15 2016 -0400

    bufio: remove unnecessary "continue"
    
    After resizing the scan buffer, we can immediately read into the
    newly-resized buffer since we know there is now space.
    
    Fixes #15712.
    
    Change-Id: I56fcfaeb67045ee753a012c37883aa7c81b6e877
    Reviewed-on: https://go-review.googlesource.com/31715
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 7ff0c8267f25a628dc4c8b5fa356b26cbb72530d
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sat Oct 22 15:43:23 2016 +0200

    cmd/compile: replace ANDL with MOV?ZX
    
    According to "Intel 64 and IA-32 Architectures Optimization Reference
    Manual" Section: "3.5.1.13 Zero-Latency MOV Instructions"
    MOV?ZX instructions have zero latency on newer processors.
    
    during make.bash:
    (ANDLconst [0xFF] x) -> (MOVBQZX x)
    applies 422 times
    (ANDLconst [0xFFFF] x) -> (MOVWQZX x)
    applies 114 times
    
    Updates #15105
    
    Change-Id: I10933af599de3c26449c52f4b5cd859331028f39
    Reviewed-on: https://go-review.googlesource.com/31639
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: David Chase <drchase@google.com>

commit 9ac60181e2e54b6404e67852d6e1e65a8cbd3616
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Mon Oct 24 12:04:12 2016 +1100

    runtime/cgo: do not link math lib by default on windows
    
    Makes windows same as others.
    
    Change-Id: Ib4651e06d0bd37473ac345d36c91f39aa8f5e662
    Reviewed-on: https://go-review.googlesource.com/31791
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3cbfcaa4ba898faba516fbb9f3debf4ceac1a557
Author: Austin Clements <austin@google.com>
Date:   Sun Oct 2 18:46:02 2016 -0400

    runtime: make mspan.isFree do what's on the tin
    
    Currently mspan.isFree technically returns whether the object was not
    allocated *during this cycle*. Fix it so it actually returns whether
    or not the object is allocated so the method is more generally useful
    (especially for debugging).
    
    It has one caller, which is carefully written to be insensitive to
    this distinction, but this lets us simplify this caller.
    
    Change-Id: I9d79cf784a56015e434961733093c1d8d03fc091
    Reviewed-on: https://go-review.googlesource.com/30145
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit bf9c71cb434a730679f54a3a87c2e9e36ec400d0
Author: Austin Clements <austin@google.com>
Date:   Wed Oct 19 18:27:39 2016 -0400

    runtime: make morestack less subtle
    
    morestack writes the context pointer to gobuf.ctxt, but since
    morestack is written in assembly (and has to be very careful with
    state), it does *not* invoke the requisite write barrier for this
    write. Instead, we patch this up later, in newstack, where we invoke
    an explicit write barrier for ctxt.
    
    This already requires some subtle reasoning, and it's going to get a
    lot hairier with the hybrid barrier.
    
    Fix this by simplifying the whole mechanism. Instead of writing
    gobuf.ctxt in morestack, just pass the value of the context register
    to newstack and let it write it to gobuf.ctxt. This is a normal Go
    pointer write, so it gets the normal Go write barrier. No subtle
    reasoning required.
    
    Updates #17503.
    
    Change-Id: Ia6bf8459bfefc6828f53682ade32c02412e4db63
    Reviewed-on: https://go-review.googlesource.com/31550
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit cdccd6a79c5391c21c8e7316e13f8b8d1697ea63
Author: Alexander Döring <email@alexd.ch>
Date:   Sun Oct 23 16:03:38 2016 +0200

    doc: update size of "hello, world" binary in FAQ
    
    Fixes #17159
    
    Change-Id: I44d7081ef7a973dcd1cc2eb7124e3454c94bc6e3
    Reviewed-on: https://go-review.googlesource.com/31770
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0e7f9700f69944a993230d706e41a86ac47da415
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Fri Oct 21 18:46:44 2016 +0900

    path/filepath: pass TestToNorm even if VolumeName(tmpdir) != VolumeName(pwd) on windows
    
    Fixes #17504
    
    Change-Id: Ic83578cf2019e5d8778e4b324f04931eb802f603
    Reviewed-on: https://go-review.googlesource.com/31544
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit b7477f386926e65bb99db4eb90820576f6533614
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Fri Sep 23 16:58:31 2016 +1000

    syscall: use ERROR_IO_PENDING value in errnoErr
    
    So errnoErr can be used in other packages.
    This is something I missed when I sent CL 28990.
    
    Fixes #17539
    
    Change-Id: I8ee3b79c4d70ca1e5b29e5b40024f7ae9a86061e
    Reviewed-on: https://go-review.googlesource.com/29690
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b992c391d4aae64e147fc64c77ad41d61be8e2e7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Oct 22 09:47:05 2016 -0700

    net/http: add NoBody, don't return nil from NewRequest on zero bodies
    
    This is an alternate solution to https://golang.org/cl/31445
    
    Instead of making NewRequest return a request with Request.Body == nil
    to signal a zero byte body, add a well-known variable that means
    explicitly zero.
    
    Too many tests inside Google (and presumably the outside world)
    broke.
    
    Change-Id: I78f6ecca8e8aa1e12179c234ccfb6bcf0ee29ba8
    Reviewed-on: https://go-review.googlesource.com/31726
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 448e1db103df7a9b29aa360f42fdcdc9b89fa399
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Oct 22 08:56:07 2016 -0700

    runtime: skip TestLldbPython
    
    The test is broken on macOS Sierra.
    
    Updates #17463.
    
    Change-Id: Ifbb2379c640b9353a01bc55a5cb26dfaad9b4bdc
    Reviewed-on: https://go-review.googlesource.com/31725
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 659570915481c87559f0197c9980e8cbac8e2c33
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 19 14:06:54 2016 -0400

    net/url: make URL implement encoding.BinaryMarshaler, BinaryUnmarshaler
    
    This makes it possible to use URLs with gob.
    
    Ideally we'd also implement TextMarshaler and TextUnmarshaler,
    but that would change the JSON encoding of a URL from something like:
    
            {"Scheme":"https","Opaque":"","User":null,"Host":"www.google.com","Path":"/x","RawPath":"","ForceQuery":false,"RawQuery":"y=z","Fragment":""}
    
    to something like:
    
            "https://www.google.com/x?y=z"
    
    That'd be nice, but it would break code expecting the old form.
    
    Fixes #10964.
    
    Change-Id: I83f06bc2bedd2ba8a5d8eef03ea0056d045c258f
    Reviewed-on: https://go-review.googlesource.com/31467
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6eae03e136f031649683d359a0879f6d6ae5e023
Author: Adam Langley <agl@golang.org>
Date:   Thu Oct 20 09:48:24 2016 -0700

    net/http: drop custom tls.Config cloning code.
    
    Now that we have the Clone method on tls.Config, net/http doesn't need
    any custom functions to do that any more.
    
    Change-Id: Ib60707d37f1a7f9a7d7723045f83e59eceffd026
    Reviewed-on: https://go-review.googlesource.com/31595
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3d4ea227c6b8062c436fc9417034f2d01cf8c82c
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Oct 18 17:22:25 2016 -0700

    archive/tar: validate sparse headers in parsePAX
    
    According to the GNU manual, the format is:
    <<<
    GNU.sparse.size=size
    GNU.sparse.numblocks=numblocks
    repeat numblocks times
      GNU.sparse.offset=offset
      GNU.sparse.numbytes=numbytes
    end repeat
    >>>
    
    The logic in parsePAX converts the repeating sequence of
    (offset, numbytes) pairs (which is not PAX compliant) into a single
    comma-delimited list of numbers (which is now PAX compliant).
    
    Thus, we validate the following:
    * The (offset, numbytes) headers must come in the correct order.
    * The ',' delimiter cannot appear in the value.
    We do not validate that the value is a parsible decimal since that
    will be determined later.
    
    Change-Id: I8d6681021734eb997898227ae8603efb1e17c0c8
    Reviewed-on: https://go-review.googlesource.com/31439
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ece4e23d9aaed3e11f7a0b9a3f15c592c96b065d
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Oct 22 06:39:12 2016 -0700

    net/http: document Transport.ExpectContinueTimeout a bit more
    
    Fixes #16003
    
    Change-Id: I76a8da24b9944647ec40ef2ca4fc93c175ff5a25
    Reviewed-on: https://go-review.googlesource.com/31723
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit eb15cf16aee1ec4387b7397b102bd2cc3ff33a64
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Oct 21 18:03:49 2016 -0700

    net: clarify LookupAddr docs on libc's behavior, and alternatives
    
    Text from rsc.
    
    Fixes #17093
    
    Change-Id: I13c3018b1584f152b53f8576dd16ebef98aa5182
    Reviewed-on: https://go-review.googlesource.com/31720
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit aa1e063efd7376e268ee592ebe078c6d05b0bdf8
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Oct 21 12:03:41 2016 +0100

    net/http: add Request.GetBody func for 307/308 redirects
    
    Updates #10767
    
    Change-Id: I197535f71bc2dc45e783f38d8031aa717d50fd80
    Reviewed-on: https://go-review.googlesource.com/31733
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Emmanuel Odeke <emm.odeke@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ca4431a3846de4b1c5cf2388ca22d915f510f7fd
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Oct 21 16:07:47 2016 -0700

    cmd/compile: avoid one symbol lookup for qualified identifiers
    
    For -newparser only.
    
    Change-Id: I0eaa05035df11734e2bda7ad456b9b30485d9465
    Reviewed-on: https://go-review.googlesource.com/31718
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6eede325ab9e444c19c8dc4fd5cc6c3603abb93c
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Oct 21 14:43:59 2016 -0700

    cmd/compile: fix detection of duplicate cases for integer ranges
    
    Previously, the check to make sure we only considered constant cases
    for duplicates was skipping past integer ranges, because those use
    n.List instead of n.Left. Thanks to Emmanuel Odeke for investigating
    and helping to identify the root cause.
    
    Fixes #17517.
    
    Change-Id: I46fcda8ed9c346ff3a9647d50b83f1555587b740
    Reviewed-on: https://go-review.googlesource.com/31716
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Emmanuel Odeke <emm.odeke@gmail.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 28279238004da8e83e3d652b4bcd14d6795c6148
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Oct 21 13:44:29 2016 -0700

    cmd/compile: prevent ICE from misuse of [...]T arrays
    
    Fixes #16428.
    
    Change-Id: I78d37472e228402bb3c06d7ebd441952386fa38a
    Reviewed-on: https://go-review.googlesource.com/31731
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit d553c29dc1f0057d1223af79d6054ce51afe2b7d
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Oct 20 17:33:45 2016 -0700

    cmd/compile: directly construct Fields instead of ODCLFIELD nodes
    
    Avoids some garbage allocations while loading import data. Seems to
    especially benefit html/template for some reason, but significant
    allocation improvements for other packages too.
    
    name       old time/op     new time/op     delta
    Template       345ms ± 6%      332ms ± 6%   -3.76%        (p=0.000 n=49+47)
    Unicode        185ms ±10%      184ms ±12%     ~           (p=0.401 n=50+49)
    GoTypes        1.04s ± 3%      1.04s ± 3%   -0.72%        (p=0.012 n=48+47)
    Compiler       4.52s ± 7%      4.49s ± 9%     ~           (p=0.465 n=48+47)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        532M ±17%       471M ±23%  -11.48%        (p=0.000 n=50+50)
    Unicode         298M ±29%       311M ±28%     ~           (p=0.065 n=50+50)
    GoTypes        1.52G ± 7%      1.54G ± 9%     ~           (p=0.062 n=49+50)
    Compiler       6.37G ± 7%      6.42G ± 8%     ~           (p=0.157 n=49+48)
    
    name       old alloc/op    new alloc/op    delta
    Template      43.9MB ± 0%     42.3MB ± 0%   -3.51%        (p=0.000 n=48+48)
    Unicode       34.3MB ± 0%     34.3MB ± 0%     ~           (p=0.945 n=50+50)
    GoTypes        123MB ± 0%      122MB ± 0%   -0.82%        (p=0.000 n=50+50)
    Compiler       522MB ± 0%      519MB ± 0%   -0.51%        (p=0.000 n=50+50)
    
    name       old allocs/op   new allocs/op   delta
    Template        414k ± 0%       397k ± 0%   -4.14%        (p=0.000 n=50+49)
    Unicode         320k ± 0%       320k ± 0%     ~           (p=0.988 n=48+49)
    GoTypes        1.18M ± 0%      1.17M ± 0%   -0.97%        (p=0.000 n=50+50)
    Compiler       4.44M ± 0%      4.41M ± 0%   -0.66%        (p=0.000 n=50+50)
    
    Passes toolstash.
    
    Change-Id: I0f54c0fa420d4f4ed3584c47cec0dde100c70c03
    Reviewed-on: https://go-review.googlesource.com/31670
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 6ca662ca0e1a41077bd93501a5d04668474bca69
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Oct 21 11:43:12 2016 -0700

    net/http: make Redirect escape non-ASCII in Location header
    
    Only ASCII is permitted there.
    
    Fixes #4385
    
    Change-Id: I63708b04a041cdada0fdfc1f2308fcb66889a27b
    Reviewed-on: https://go-review.googlesource.com/31732
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e64e858c51bc7739a308fa0749ded1123da2e89c
Author: Adam Langley <agl@golang.org>
Date:   Thu Oct 20 09:35:19 2016 -0700

    vendor/golang_org/x/crypto/curve25519: update to f62085100e1abe3d5c9b3b8c9a38d50b71323f64
    
    This change updates the vendored copy of x/crypto/curve25519,
    specifically to include the following changes:
      f620851 curve25519: eliminate unnecessary "callee save" prologues
      722a7b7 curve25519: fix confusing SP adjustments
    
    Change-Id: I1b16aba12c744ac32f925654a136a8f52cd40fc2
    Reviewed-on: https://go-review.googlesource.com/31666
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a50fbcd331c57b885c13a6c0c2502202417ce312
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Oct 21 13:55:25 2016 -0700

    net/http: update bundled http2
    
    Updates http2 to x/net/http2 git rev 40a0a18 for:
    
        http2: fix Server race with concurrent Read/Close
        http2: make Server reuse 64k request body buffer between requests
        http2: never Read from Request.Body in Transport to determine ContentLength
    
    Fixes #17480
    Updates #17071
    
    Change-Id: If142925764a2e148f95957f559637cfc1785ad21
    Reviewed-on: https://go-review.googlesource.com/31737
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 40d4be59ccd83c29a72fb6a9a8c2eee4c91f628b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Oct 21 10:27:57 2016 -0700

    net: make all Resolver methods respect Resolver.PreferGo
    
    Fixes #17532
    
    Change-Id: Id62671d505c77ea924b3570a504cdc3b157e5a0d
    Reviewed-on: https://go-review.googlesource.com/31734
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 23173fc025f769aaa9e19f10aa0f69c851ca2f3b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Oct 21 12:14:36 2016 +0100

    net/http/httptrace: clarify ClientTrace docs
    
    The old wording over-promised.
    
    Fixes #16957
    
    Change-Id: Iaac04de0d24eb17a0db66beeeab9de70d0f6d391
    Reviewed-on: https://go-review.googlesource.com/31735
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Tom Bergan <tombergan@google.com>

commit f0e347b1a8161c41be9dbcfe5db5e4e4dae57cfd
Author: Joe Tsai <joetsai@google.com>
Date:   Fri Oct 21 19:43:33 2016 +0000

    Revert "cmd/compile: cleanup toolstash hacks from previous CL"
    
    This partially reverts commit 01bf5cc21912ff8642171d8255a7fff87f1da00b.
    
    For unknown reasons, this CL was causing an internal test to allocate
    1.2GB when it used to allocate less than 300MB.
    
    Change-Id: I41d767781e0ae9e43bf670e2a186ee074821eca4
    Reviewed-on: https://go-review.googlesource.com/31674
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b7c7949817e5bee801dfb333f483a080fa1f29e7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Oct 21 07:54:46 2016 -0700

    cmd/cgo: preserve original call arguments when pointer checking
    
    With the old code rewriting refs would rewrite the inner arguments
    rather than the outer ones, leaving a reference to C.val in the outer
    arguments.
    
    Change-Id: I9b91cb4179eccd08500d14c6591bb15acf8673eb
    Reviewed-on: https://go-review.googlesource.com/31672
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a73d68e75eece80f8514bb1b368420843c1f58ad
Author: Austin Clements <austin@google.com>
Date:   Thu Oct 20 22:27:39 2016 -0400

    runtime: fix call* signatures and deferArgs with siz=0
    
    This commit fixes two bizarrely related bugs:
    
    1. The signatures for the call* functions were wrong, indicating that
    they had only two pointer arguments instead of three. We didn't notice
    because the call* functions are defined by a macro expansion, which go
    vet doesn't see.
    
    2. deferArgs on a defer object with a zero-sized frame returned a
    pointer just past the end of the allocated object, which is illegal in
    Go (and can cause the "sweep increased allocation count" crashes).
    
    In a fascinating twist, these two bugs canceled each other out, which
    is why I'm fixing them together. The pointer returned by deferArgs is
    used in only two ways: as an argument to memmove and as an argument to
    reflectcall. memmove is NOSPLIT, so the argument was unobservable.
    reflectcall immediately tail calls one of the call* functions, which
    are not NOSPLIT, but the deferArgs pointer just happened to be the
    third argument that was accidentally marked as a scalar. Hence, when
    the garbage collector scanned the stack, it didn't see the bad
    pointer as a pointer.
    
    I believe this was all ultimately benign. In principle, stack growth
    during the reflectcall could fail to update the args pointer, but it
    never points to the stack, so it never needs to be updated. Also in
    principle, the garbage collector could fail to mark the args object
    because of the incorrect call* signatures, but in all calls to
    reflectcall (including the ones spelled "call" in the reflect package)
    the args object is kept live by the calling stack.
    
    Change-Id: Ic932c79d5f4382be23118fdd9dba9688e9169e28
    Reviewed-on: https://go-review.googlesource.com/31654
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit c2425178669c0c6a26f58f90b59086d7e4313c64
Author: Austin Clements <austin@google.com>
Date:   Wed Oct 19 16:01:02 2016 -0400

    runtime: replace *g with guintptr in trace
    
    trace's reader *g is going to cause write barriers in unfortunate
    places, so replace it with a guintptr.
    
    Change-Id: Ie8fb13bb89a78238f9d2a77ec77da703e96df8af
    Reviewed-on: https://go-review.googlesource.com/31469
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>
```
