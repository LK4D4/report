# April 4, 2017 Report

Number of commits: 851

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 1.454714055s to 1.281112328s, -11.93%
* github.com/coreos/etcd/cmd/etcd: from 21.230020206s to 21.111207837s, -0.56%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 1m46.379191537s to 1m24.783226565s, -20.30%
* github.com/junegunn/fzf/src/fzf: from 2.431548854s to 2.443536946s, +0.49%
* github.com/mholt/caddy/caddy: from 11.879836126s to 12.392737118s, +4.32%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 2.409353758s to 2.451338094s, +1.74%
* github.com/nsqio/nsq/apps/nsqd: from 4.289575855s to 3.887553981s, -9.37%
* github.com/prometheus/prometheus/cmd/prometheus: from 1m16.015199849s to 1m10.527950959s, -7.22%
* github.com/spf13/hugo: from 13.947993964s to 13.25205243s, -4.99%
* golang.org/x/tools/cmd/guru: from 5.630548621s to 5.60178682s, -0.51%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 3201269 to 2983095, -6.82%
* github.com/coreos/etcd/cmd/etcd: from 27770632 to 28364400, +2.14%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 34608824 to 35411216, +2.32%
* github.com/junegunn/fzf/src/fzf: from 3211624 to 3425864, +6.67%
* github.com/mholt/caddy/caddy: from 16223462 to 16908916, +4.23%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4375596 to 4714140, +7.74%
* github.com/nsqio/nsq/apps/nsqd: from 10398320 to 10842483, +4.27%
* github.com/prometheus/prometheus/cmd/prometheus: from 62363413 to 63843797, +2.37%
* github.com/spf13/hugo: from 17418348 to 18142858, +4.16%
* golang.org/x/tools/cmd/guru: from 8276836 to 8570919, +3.55%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               338           360           +6.51%
BenchmarkMsgpUnmarshal-4             1142          879           -23.03%
BenchmarkEasyJsonMarshal-4           3774          3445          -8.72%
BenchmarkEasyJsonUnmarshal-4         6979          5812          -16.72%
BenchmarkFlatBuffersMarshal-4        1761          1404          -20.27%
BenchmarkFlatBuffersUnmarshal-4      769           657           -14.56%
BenchmarkGogoprotobufMarshal-4       396           353           -10.86%
BenchmarkGogoprotobufUnmarshal-4     627           680           +8.45%

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

* [bytes: use Index in Count](https://github.com/golang/go/commit/438818d9f1b5b4ffae3ca63d5ce2a2f5cef97552)
* [bytes, strings: optimize Split*](https://github.com/golang/go/commit/894650277670eed065566f803c642a8f80437456)
* [database/sql: do not exhaust connection pool on conn request timeout](https://github.com/golang/go/commit/4f6d4bb3f4461e7e25eff24254115b689495e834)
* [time: optimize Now on darwin, windows](https://github.com/golang/go/commit/e4371fb179ad69cbd057f2430120843948e09f2f)
* [runtime: implement fastrand in go](https://github.com/golang/go/commit/d03c1248604679e1e6a01253144065bc57da48b8)
* [runtime: use two-level list for semaphore address search in semaRoot](https://github.com/golang/go/commit/45c6f59e1fd94ccb11fde61ca8d5b33b3d06dd9f)
* [runtime: speed up fastrand() % n](https://github.com/golang/go/commit/46a75870ad5b9b9711e69ffce3738a3ab2057789)
* [os: use poller for file i/o](https://github.com/golang/go/commit/c05b06a12d005f50e4776095a60d6bd9c2c91fac)
* [runtime: do not call wakep from enlistWorker, to avoid possible deadlock](https://github.com/golang/go/commit/1f77db94f8a453ae96e490fe729c8c6b0ba9479f)
* [runtime: use balanced tree for addr lookup in semaphore implementation](https://github.com/golang/go/commit/990124da2a6ca5a54b38733b51018e2f8758cfae)
* [sync: make Mutex more fair](https://github.com/golang/go/commit/0556e26273f704db73df9e7c4c3d2e8434dec7be)
* [runtime: do not allocate on every time.Sleep](https://github.com/golang/go/commit/a1261b8b0a38814df453defb2fc2cae3ba0c956a)
* [os/user: add Go implementation of LookupGroup, LookupGroupId](https://github.com/golang/go/commit/949f95e7a40715ad05015dc4cb039e78a5260ef8)
* [bytes: make bytes.Buffer cache-friendly](https://github.com/golang/go/commit/55310403ddf051634fa398b4c9e6013d530753f5)
* [runtime: don't rescan finalizers queue during mark termination](https://github.com/golang/go/commit/f1ba75f8c577e1471f646ef3715fc2f41dd423ef)
* [runtime: make ReadMemStats STW for < 25µs](https://github.com/golang/go/commit/4a7cf960c38d72e9f0c6f00e46e013be2a35d56e)
* [runtime: optimize slicebytestostring](https://github.com/golang/go/commit/23be728950be7973a4c4852449e1987c64dc2739)
* [strconv: optimize formatting for small decimal ints](https://github.com/golang/go/commit/bc8b9b23ca3f33e738d85c4734bd35dfe63e9ac4)
* [bytes: add optimized countByte for amd64](https://github.com/golang/go/commit/01cd22c68792b659ca0912c104b14c86044110cb)
* [syscall: use CLONE_VFORK and CLONE_VM](https://github.com/golang/go/commit/9e6b79a5dfb2f6fe4301ced956419a0da83bd025)
* [cmd/go: exclude vendored packages from ... matches](https://github.com/golang/go/commit/fa1d54c2edad607866445577fe4949fbe55166e1)
* [strings: speed up Fields](https://github.com/golang/go/commit/bebfd4ba415cbfee578f64177fe1c59dab5a1df8)


## GIT Log

```
commit 4c1622082e493dea24a936930be8b324aae54505
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Apr 4 15:29:31 2017 -0700

    cmd/compile: don't catch panics during rewrite
    
    This is a holdover from the days when we did not
    have full SSA coverage and compiled things optimistically,
    and catching the panic obscures useful information.
    
    Change-Id: I196790cb6b97419d92b318a2dfa7f1e1097cefb7
    Reviewed-on: https://go-review.googlesource.com/39534
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit fc327a14c490db9d2e14c4c44ba8791dc54be02a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Apr 4 15:26:04 2017 -0700

    cmd/compile: remove order canonicalization rules from mips
    
    CL 38801 introduced automatic commutative rule generation.
    Manual order canonicalization rules thus lead to infinite loops.
    
    Fixes #19842
    
    Change-Id: I877c476152f4d207fdc67bc6f3018265aa9bc5ac
    Reviewed-on: https://go-review.googlesource.com/39533
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 6c5a819a5edd602236200627694215a1017aded2
Author: Filip Gruszczyński <gruszczy@gmail.com>
Date:   Fri Mar 17 20:10:38 2017 -0700

    reflect: add MakeMapWithSize for creating maps with size hint
    
    Providing size hint when creating a map allows avoiding re-allocating
    underlying data structure if we know how many elements are going to
    be inserted. This can be used for example during decoding maps in
    gob.
    
    Fixes #19599
    
    Change-Id: I108035fec29391215d2261a73eaed1310b46bab1
    Reviewed-on: https://go-review.googlesource.com/38335
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit be5a201e2a59dbe45157b8939a507830206d86fb
Author: Rob Pike <r@golang.org>
Date:   Tue Mar 21 10:00:30 2017 -0700

    text/template: fix handling of empty blocks
    
    This was a subtle bug introduced in the previous release's fix for
    issue 16156.
    
    The definition of empty template was broken, causing the answer
    to depend on the order of templates in the map.
    
    Fixes #16156 (for real).
    Fixes #19294.
    Fixes #19204.
    
    Change-Id: I1cd915c94534cad3116d83bd158cbc28700510b9
    Reviewed-on: https://go-review.googlesource.com/38420
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bebfd4ba415cbfee578f64177fe1c59dab5a1df8
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Mon Mar 6 09:34:39 2017 +0100

    strings: speed up Fields
    
    - use a string lookup to detect if a single byte is a space character
    - determine the exact number of fields for ASCII and
      a possibly underestimated number of fields for non ASCII strings
      by doing a separate byte for byte scan of the input string
      before collecting the fields in an extra pass
    - provide a fast path for ASCII only strings when collecting the fields
    - avoid utf8.DecodeRuneInString and unicode.IsSpace for ASCII characters
    
    Used golang.org/cl/33108 from Joe Tsai as starting point.
    
    name                      old time/op    new time/op     delta
    Fields/ASCII/16              284ns ± 1%      116ns ± 2%   -59.30%  (p=0.000 n=9+10)
    Fields/ASCII/256            3.81µs ± 1%     0.80µs ± 1%   -79.10%  (p=0.000 n=10+10)
    Fields/ASCII/4096           61.4µs ± 1%     12.3µs ± 1%   -79.96%  (p=0.000 n=10+9)
    Fields/ASCII/65536           982µs ± 1%      235µs ± 0%   -76.04%  (p=0.000 n=10+9)
    Fields/ASCII/1048576        16.7ms ± 2%      5.4ms ± 1%   -67.52%  (p=0.000 n=10+10)
    Fields/Mixed/16              314ns ± 1%      168ns ± 1%   -46.33%  (p=0.000 n=9+10)
    Fields/Mixed/256            3.92µs ± 1%     1.17µs ± 1%   -70.19%  (p=0.000 n=10+10)
    Fields/Mixed/4096           69.1µs ± 1%     19.0µs ± 1%   -72.53%  (p=0.000 n=10+10)
    Fields/Mixed/65536          1.12ms ± 1%     0.39ms ± 0%   -65.37%  (p=0.000 n=10+9)
    Fields/Mixed/1048576        19.0ms ± 2%      7.3ms ± 4%   -61.75%  (p=0.000 n=10+9)
    
    name                      old speed      new speed       delta
    Fields/ASCII/16           56.3MB/s ± 1%  138.1MB/s ± 2%  +145.31%  (p=0.000 n=9+10)
    Fields/ASCII/256          67.1MB/s ± 1%  321.0MB/s ± 1%  +378.26%  (p=0.000 n=10+10)
    Fields/ASCII/4096         66.7MB/s ± 1%  333.0MB/s ± 1%  +398.97%  (p=0.000 n=10+9)
    Fields/ASCII/65536        66.7MB/s ± 1%  278.4MB/s ± 0%  +317.39%  (p=0.000 n=10+9)
    Fields/ASCII/1048576      62.7MB/s ± 2%  192.9MB/s ± 1%  +207.82%  (p=0.000 n=10+10)
    Fields/Mixed/16           51.0MB/s ± 2%   94.9MB/s ± 1%   +85.87%  (p=0.000 n=10+10)
    Fields/Mixed/256          65.4MB/s ± 1%  219.2MB/s ± 1%  +235.33%  (p=0.000 n=10+10)
    Fields/Mixed/4096         59.3MB/s ± 1%  215.7MB/s ± 1%  +263.98%  (p=0.000 n=10+10)
    Fields/Mixed/65536        58.6MB/s ± 1%  169.1MB/s ± 0%  +188.73%  (p=0.000 n=10+9)
    Fields/Mixed/1048576      55.1MB/s ± 2%  144.0MB/s ± 4%  +161.44%  (p=0.000 n=10+9)
    
    Updates #19789
    Updates #17856
    
    Change-Id: If2ce1479542702e9cd65a82a462ba55ac8eb3876
    Reviewed-on: https://go-review.googlesource.com/37959
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 5cadc91b3ced9614b1055c448f7784a15907fff5
Author: Keith Randall <keithr@alum.mit.edu>
Date:   Thu Mar 16 21:33:03 2017 -0700

    cmd/compile: intrinsics for math/bits.OnesCount
    
    Popcount instructions on amd64 are not guaranteed to be
    present, so we must guard their call.  Rewrite rules can't
    generate control flow at the moment, so the intrinsifier
    needs to generate that code.
    
    name           old time/op  new time/op  delta
    OnesCount-8    2.47ns ± 5%  1.04ns ± 2%  -57.70%  (p=0.000 n=10+10)
    OnesCount16-8  1.05ns ± 1%  0.78ns ± 0%  -25.56%    (p=0.000 n=9+8)
    OnesCount32-8  1.63ns ± 5%  1.04ns ± 2%  -35.96%  (p=0.000 n=10+10)
    OnesCount64-8  2.45ns ± 0%  1.04ns ± 1%  -57.55%   (p=0.000 n=6+10)
    
    Update #18616
    
    Change-Id: I4aff2cc9aa93787898d7b22055fe272a7cf95673
    Reviewed-on: https://go-review.googlesource.com/38320
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 59f6549d1c7e0e074472c46f55716267225f4fd6
Author: Eric Lagergren <ericscottlagergren@gmail.com>
Date:   Mon Apr 3 16:08:13 2017 -0700

    bytes, strings: declare variables inside loop they're used in
    
    The recently updated Count functions declare variables before
    special-cased returns.
    
    Change-Id: I8f726118336b7b0ff72117d12adc48b6e37e60ea
    Reviewed-on: https://go-review.googlesource.com/39357
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 42426ed41167d6a99cfc9e5a91a4aff1b95093ca
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Apr 3 16:14:26 2017 -0700

    cmd/compile: Fatal instead of panic in large bvbulkalloc
    
    This provides better diagnostics when it occurs.
    
    Updates #19751
    
    Change-Id: I87db54c22e1345891b418c1741dc76ac5fb8ed00
    Reviewed-on: https://go-review.googlesource.com/39358
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 094498c9a13cd711ed45a65b153393eb8ae1566b
Author: Eric Lagergren <ericscottlagergren@gmail.com>
Date:   Mon Apr 3 15:54:20 2017 -0700

    all: fix minor misspellings
    
    Change-Id: I1f1cfb161640eb8756fb1a283892d06b30b7a8fa
    Reviewed-on: https://go-review.googlesource.com/39356
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 50688fcb6ef0f74d6ff91dce95f8823b1f56bdf7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Apr 3 11:27:06 2017 -0700

    cmd/compile: unroll small static maps
    
    When a map is small, it's not worth putting
    the contents in an array and then looping over the array.
    Just generate code instead.
    
    This makes smaller binaries.
    It might also be better for cache lines.
    
    It also can avoids adding control flow in the middle
    of the init function, which can be very large.
    Eliminating this source of extra blocks
    makes phi insertion easier for temp-heavy init functions.
    This reduces the time required for compiler to
    panic while compiling the code in #19751
    from 15 minutes to 45 seconds.
    
    The cutoff of 25 was chosen fairly unscientifically
    by looking at the size of cmd/go.
    
    Cutoff of   0: 10689604
    Cutoff of   5: 10683572
    Cutoff of  15: 10682324
    Cutoff of  25: 10681700
    Cutoff of  50: 10685476
    Cutoff of 100: 10689412
    
    There are probably more sophisticated mechanisms available.
    For example, the smaller the key/value sizes, the better
    generated code will be vs a table.
    Nevertheless this is simple and seems like a good start.
    
    Updates #19751
    
    name       old time/op     new time/op     delta
    Template       204ms ± 6%      202ms ± 5%  -0.78%  (p=0.027 n=47+45)
    Unicode       84.8ms ± 6%     85.2ms ± 7%    ~     (p=0.146 n=46+45)
    GoTypes        551ms ± 2%      556ms ± 3%  +0.76%  (p=0.004 n=43+45)
    SSA            3.93s ± 3%      3.95s ± 4%    ~     (p=0.179 n=50+49)
    Flate          123ms ± 4%      123ms ± 5%    ~     (p=0.201 n=47+49)
    GoParser       145ms ± 3%      145ms ± 4%    ~     (p=0.937 n=50+50)
    Reflect        356ms ± 3%      354ms ± 5%  -0.44%  (p=0.048 n=46+50)
    Tar            107ms ± 6%      106ms ± 6%    ~     (p=0.188 n=50+49)
    XML            201ms ± 4%      200ms ± 4%    ~     (p=0.085 n=50+49)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        252M ± 9%       250M ± 7%    ~     (p=0.206 n=49+47)
    Unicode         106M ± 7%       106M ± 9%    ~     (p=0.331 n=47+46)
    GoTypes         724M ± 5%       729M ± 5%    ~     (p=0.160 n=47+49)
    SSA            5.64G ± 2%      5.62G ± 4%    ~     (p=0.148 n=47+50)
    Flate           147M ± 6%       147M ± 5%    ~     (p=0.466 n=50+49)
    GoParser        179M ± 5%       179M ± 6%    ~     (p=0.584 n=50+49)
    Reflect         448M ± 6%       441M ± 8%  -1.39%  (p=0.027 n=50+49)
    Tar             124M ± 6%       123M ± 5%    ~     (p=0.221 n=50+47)
    XML             244M ± 5%       243M ± 4%    ~     (p=0.275 n=49+49)
    
    name       old alloc/op    new alloc/op    delta
    Template      39.9MB ± 0%     39.4MB ± 0%  -1.28%  (p=0.008 n=5+5)
    Unicode       29.8MB ± 0%     29.8MB ± 0%    ~     (p=0.310 n=5+5)
    GoTypes        113MB ± 0%      113MB ± 0%    ~     (p=0.421 n=5+5)
    SSA            854MB ± 0%      854MB ± 0%    ~     (p=0.151 n=5+5)
    Flate         25.3MB ± 0%     25.3MB ± 0%    ~     (p=1.000 n=5+5)
    GoParser      31.8MB ± 0%     31.8MB ± 0%    ~     (p=0.222 n=5+5)
    Reflect       78.2MB ± 0%     78.2MB ± 0%    ~     (p=1.000 n=5+5)
    Tar           26.7MB ± 0%     26.7MB ± 0%    ~     (p=0.841 n=5+5)
    XML           42.3MB ± 0%     42.3MB ± 0%  -0.15%  (p=0.008 n=5+5)
    
    name       old allocs/op   new allocs/op   delta
    Template        390k ± 1%       386k ± 1%  -1.05%  (p=0.016 n=5+5)
    Unicode         319k ± 0%       320k ± 0%    ~     (p=0.310 n=5+5)
    GoTypes        1.14M ± 0%      1.14M ± 0%    ~     (p=0.421 n=5+5)
    SSA            7.60M ± 0%      7.59M ± 0%    ~     (p=0.310 n=5+5)
    Flate           234k ± 0%       235k ± 1%    ~     (p=1.000 n=5+5)
    GoParser        315k ± 1%       317k ± 0%    ~     (p=0.151 n=5+5)
    Reflect         978k ± 0%       978k ± 0%    ~     (p=0.841 n=5+5)
    Tar             251k ± 1%       251k ± 1%    ~     (p=0.690 n=5+5)
    XML             394k ± 0%       392k ± 0%    ~     (p=0.056 n=5+5)
    
    
    Change-Id: Ic53a18627082abe075a1cbc33330ce015e50850a
    Reviewed-on: https://go-review.googlesource.com/39354
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 53f8a6aeb089f981c9d4fc7871ef712d1669d0a4
Author: Keith Randall <khr@golang.org>
Date:   Thu Mar 30 03:30:22 2017 +0000

    cmd/compile: automatically handle commuting ops in rewrite rules
    
    Note that this is a redo of an undo of the original buggy CL 38666.
    
    We have lots of rewrite rules that vary only in the fact that
    we have 2 versions for the 2 different orderings of various
    commuting ops. For example:
    
    (ADDL x (MOVLconst [c])) -> (ADDLconst [c] x)
    (ADDL (MOVLconst [c]) x) -> (ADDLconst [c] x)
    
    It can get unwieldly quickly, especially when there is more than
    one commuting op in a rule.
    
    Our existing "fix" for this problem is to have rules that
    canonicalize the operations first. For example:
    
    (Eq64 x (Const64 <t> [c])) && x.Op != OpConst64 -> (Eq64 (Const64 <t> [c]) x)
    
    Subsequent rules can then assume if there is a constant arg to Eq64,
    it will be the first one. This fix kinda works, but it is fragile and
    only works when we remember to include the required extra rules.
    
    The fundamental problem is that the rule matcher doesn't
    know anything about commuting ops. This CL fixes that fact.
    
    We already have information about which ops commute. (The register
    allocator takes advantage of commutivity.)  The rule generator now
    automatically generates multiple rules for a single source rule when
    there are commutative ops in the rule. We can now drop all of our
    almost-duplicate source-level rules and the canonicalization rules.
    
    I have some CLs in progress that will be a lot less verbose when
    the rule generator handles commutivity for me.
    
    I had to reorganize the load-combining rules a bit. The 8-way OR rules
    generated 128 different reorderings, which was causing the generator
    to put too much code in the rewrite*.go files (the big ones were going
    from 25K lines to 132K lines). Instead I reorganized the rules to
    combine pairs of loads at a time. The generated rule files are now
    actually a bit (5%) smaller.
    
    Make.bash times are ~unchanged.
    
    Compiler benchmarks are not observably different. Probably because
    we don't spend much compiler time in rule matching anyway.
    
    I've also done a pass over all of our ops adding commutative markings
    for ops which hadn't had them previously.
    
    Fixes #18292
    
    Change-Id: Ic1c0e43fbf579539f459971625f69690c9ab8805
    Reviewed-on: https://go-review.googlesource.com/38801
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 63a72fd447abb8a07bee9166e87bfe27780492c3
Author: Keith Randall <khr@golang.org>
Date:   Mon Apr 3 10:17:48 2017 -0700

    cmd/compile: strength-reduce floating point
    
    x*2 -> x+x
    x/c, c power of 2 -> x*(1/c)
    
    Fixes #19827
    
    Change-Id: I74c9f0b5b49b2ed26c0990314c7d1d5f9631b6f1
    Reviewed-on: https://go-review.googlesource.com/39295
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 095a62c3494e5ca6290357d1b4d307f502257578
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 29 09:08:39 2017 -0700

    cmd/compile: refactor maplit
    
    Instead of walking the list of nodes twice,
    once to find static entries to add to an array
    and once to find dynamic entries to generate code for,
    do the split once up front, into two slices.
    Then process each slice individually.
    This makes the code easier to read
    and more importantly, easier to modify.
    
    While we're here, add a TODO to avoid
    using temporaries for mapassign_fast calls.
    It's not an important TODO;
    the generated code would be basically identical.
    It would just avoid a minor amount of
    pointless SSA optimization work.
    
    Passes toolstash-check.
    No measureable compiler performance impact.
    
    Updates #19751
    
    Change-Id: I84a8f2c22f9025c718ef34639059d7bd02a3c406
    Reviewed-on: https://go-review.googlesource.com/39351
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1d6a499cc02012275b3c172dc98a143e49fb8ddc
Author: Russ Cox <rsc@golang.org>
Date:   Mon Apr 3 14:34:51 2017 -0400

    encoding/pem: yet another fuzz fake failure
    
    Fixes #19829.
    
    Change-Id: I8500fd73c37b504d6ea25f5aff7017fbc0718570
    Reviewed-on: https://go-review.googlesource.com/39314
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 69fe9ea43e66552c11f2b87cdf63a39448b46287
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Apr 3 17:50:34 2017 +0000

    cmd/compile/internal/ssa: use recently agreed upon generated code header
    
    Updates #13560
    
    Change-Id: I9bc08ca5cf0627e653d55f748ebb83be8b69ea3b
    Reviewed-on: https://go-review.googlesource.com/39296
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7439ba32ffe5461372a5b51cc315d8c0e8dae03e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Apr 3 10:00:02 2017 -0700

    cmd/compile: respect Node.Bounded when inserting OpArraySelect
    
    This triggers 119 times during make.bash.
    
    This CL reduces the time it takes for the
    compiler to panic while compiling the code in #19751
    from 22 minutes to 15 minutes. Yay, I guess.
    
    Updates #19751
    
    Change-Id: I8ca7f1ae75f89d1eb2a361d67b3055a975221734
    Reviewed-on: https://go-review.googlesource.com/39294
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 9d5987d79f42b018ce8c57ac2d90ce0d81f0d646
Author: David Chase <drchase@google.com>
Date:   Mon Apr 3 11:50:54 2017 -0400

    cmd/compile: rewrite upper-bit-clear idiom to use shift-rotate
    
    Old buggy hardware incorrectly executes the shift-left-K
    then shift-right-K idiom for clearing K leftmost bits.
    Use a right rotate instead of shift to avoid triggering the
    bug.
    
    Fixes #19809.
    
    Change-Id: I6dc646b183c29e9d01aef944729f34388dcc687d
    Reviewed-on: https://go-review.googlesource.com/39310
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit d6b999436a5ad7f303c20c018867e1e118572fa0
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Mon Apr 3 15:18:30 2017 +0100

    go/parser: fix example to run on the playground
    
    The example shouldn't rely on the existance of example_test.go. That
    breaks in the playground, which is what the "run" button in
    https://golang.org/pkg/go/parser/#example_ParseFile does.
    
    Make the example self-sufficient by using a small piece of source via a
    string literal instead.
    
    Fixes #19823.
    
    Change-Id: Ie8a3c6c5d00724e38ff727862b62e6a3621adc88
    Reviewed-on: https://go-review.googlesource.com/39236
    Run-TryBot: Daniel Martí <mvdan@mvdan.cc>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 26308fb4813377def1391ad4ea383f9178c2d16a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Apr 3 07:50:56 2017 -0700

    cmd/internal/obj: use string instead of LSym in Pcln
    
    In a concurrent backend, Ctxt.Lookup will need some
    form of concurrency protection, which will make it
    more expensive.
    
    This CL changes the pcln table builder to track
    filenames as strings rather than LSyms.
    Those strings are then converted into LSyms
    at the last moment, for writing the object file.
    
    This CL removes over 85% of the calls to Ctxt.Lookup
    in a run of make.bash.
    
    Passes toolstash-check.
    
    Updates #15756
    
    Change-Id: I3c53deff6f16f2643169f3bdfcc7aca2ca58b0a4
    Reviewed-on: https://go-review.googlesource.com/39291
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 719c7b03ba5d8bdea937a6b21564fa82539d3661
Author: Russ Cox <rsc@golang.org>
Date:   Fri Mar 31 12:46:35 2017 -0400

    testing/quick: generate all possible int64, uint64 values
    
    When generating a random int8, uint8, int16, uint16, int32, uint32,
    quick.Value chooses among all possible values.
    
    But when generating a random int64 or uint64, it only chooses
    values in the range [-2⁶², 2⁶²) (even for uint64).
    It should, like for all the other integers, use the full range.
    
    If it had, this would have caught #19807 earlier.
    Instead it let us discover the presence of #19809.
    
    While we are here, also make the default source of
    randomness not completely deterministic.
    
    Fixes #19808.
    
    Change-Id: I070f852531c92b3670bd76523326c9132bfc9416
    Reviewed-on: https://go-review.googlesource.com/39152
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 65c17a05e98866d1e29a5d53fc21b0221760698d
Author: Russ Cox <rsc@golang.org>
Date:   Fri Mar 31 16:15:51 2017 -0400

    encoding/pem: do not try to round trip value with leading/trailing space
    
    The header is literally
    
            Key: Value
    
    If the value or the key has leading or trailing spaces, those will
    be lost by the round trip.
    
    Found because testing/quick returns different values now.
    
    Change-Id: I0f574bdbb5990689509c24309854d8f814b5efa0
    Reviewed-on: https://go-review.googlesource.com/39211
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 64f00fb150fff62a144e188e00f59f044ffb7d23
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Apr 2 19:12:19 2017 -0700

    cmd/compile: len(n.List.Slice()) -> n.List.Len()
    
    Minor cleanup.
    
    This is the only such instance in the compiler.
    
    Change-Id: I4e8ecde57d71867c7e1ac4d17e2154a91dd262b0
    Reviewed-on: https://go-review.googlesource.com/39209
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 96af8174975dcc18b6d13dad46c35bd1d7264d37
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Apr 2 18:37:04 2017 -0700

    cmd/compile: add block profiling support
    
    Updates #15756
    
    Change-Id: Ic635812b324af926333122c02908cebfb24d7bce
    Reviewed-on: https://go-review.googlesource.com/39208
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 877574725ba251be743c47ecbe49958b6cf8b814
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 31 18:06:19 2017 -0700

    cmd/compile: enforce that all nodes are used when generating DWARF
    
    No particular need for this,
    but it's nice to enforce invariants
    when they are available.
    
    Change-Id: Ia6fa88dc4116f65dac2879509746e123e2c1862a
    Reviewed-on: https://go-review.googlesource.com/39201
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 37515135626ba0bb81c3f2befc3393f862290826
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 31 16:52:50 2017 -0700

    cmd/compile: don't modify nodfp in AllocFrame
    
    nodfp is a global, so modifying it is unsafe in a concurrent backend.
    It is also not necessary, since the Used marks
    are only relevant for nodes in fn.Dcl.
    For good measure, mark nodfp as always used.
    
    Passes toolstash-check.
    
    Updates #15756
    
    Change-Id: I5320459f5eced2898615a17b395a10c1064bcaf5
    Reviewed-on: https://go-review.googlesource.com/39200
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a1cedf08428bdb91916bb5317c8413212308048c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Mar 31 07:14:16 2017 -0400

    cmd/link: canonicalize the "package" of dupok text symbols
    
    Dupok symbols may be defined in multiple packages. Its associated
    package is chosen sort of arbitrarily (the first containing package
    that the linker loads). Canonicalize its package to the package
    with which it will be laid down in text, which is the first package
    in dependency order that defines the symbol. So later passes (for
    example, trampoline insertion pass) know that the dupok symbol
    is laid down along with the package.
    
    Fixes #19764.
    
    Change-Id: I7cbc7474ff3016d5069c8b7be04af934abab8bc3
    Reviewed-on: https://go-review.googlesource.com/39150
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>
    Reviewed-by: David Chase <drchase@google.com>

commit 78f6622b818cb0198a9c7c45eca58ba6f9f235c5
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Mar 29 14:21:50 2017 +1100

    cmd/internal/obj/*: rename Rconv to rconv
    
    Each architecture's Rconv function is only used inside its
    respective package, so it does not need to be exported.
    
    Change-Id: Ifbd629964d7a9edd66501d7cdf4750621d66d646
    Reviewed-on: https://go-review.googlesource.com/39110
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d96f9cbb87bfdbb809d1f4f18e0552b94475b86e
Author: Jamie Stackhouse <contin673@gmail.com>
Date:   Sat Apr 1 01:07:59 2017 -0300

    mime/multipart: add Size to FileHeader
    
    This change makes it possible to retrieve the size of a file part
    without having to Seek to determine file-size.
    
    Resolves #19501
    
    Change-Id: I7b9994c4cf41c9b06a046eb7046f8952ae1f15e9
    Reviewed-on: https://go-review.googlesource.com/39223
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8a4cee67afeda7c89e4a6e97cd95820f1904095d
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Thu Mar 23 21:22:32 2017 -0600

    cmd/compile: use yyerrorl in typecheckswitch
    
    Replace yyerror usages with yyerrorl in function
    typecheckswitch.
    
    Updates #19683.
    
    Change-Id: I7188cdecddd2ce4e06b8cee45b57f3765a979405
    Reviewed-on: https://go-review.googlesource.com/38597
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e4e1d089bea9afa1c97a9613d1f820a2e896954e
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 31 10:09:27 2017 -0700

    go/types: use std "DO NOT EDIT" comment for generated hilbert test
    
    For #13560.
    
    Change-Id: I884e63f89d0756ca87b8c2092b4fd8360f791a2f
    Reviewed-on: https://go-review.googlesource.com/39171
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a33903b02c4e13f881676bd2619986b058a87897
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 28 13:52:33 2017 -0700

    cmd/compile: evaluate itabname during walk instead of SSA
    
    For backend concurrency safety. Follow-up to CL 38721.
    
    This does introduce a Nodes where there wasn't one before,
    but these are so rare that the performance impact is negligible.
    
    Does not pass toolstash-check, but the only change is line numbers,
    and the new line numbers appear preferable.
    
    Updates #15756
    
    name       old alloc/op    new alloc/op    delta
    Template      39.9MB ± 0%     39.9MB ± 0%    ~     (p=0.841 n=5+5)
    Unicode       29.8MB ± 0%     29.8MB ± 0%    ~     (p=0.690 n=5+5)
    GoTypes        113MB ± 0%      113MB ± 0%  +0.09%  (p=0.008 n=5+5)
    SSA            854MB ± 0%      855MB ± 0%    ~     (p=0.222 n=5+5)
    Flate         25.3MB ± 0%     25.3MB ± 0%    ~     (p=0.690 n=5+5)
    GoParser      31.8MB ± 0%     31.9MB ± 0%    ~     (p=0.421 n=5+5)
    Reflect       78.2MB ± 0%     78.3MB ± 0%    ~     (p=0.548 n=5+5)
    Tar           26.7MB ± 0%     26.7MB ± 0%    ~     (p=0.690 n=5+5)
    XML           42.3MB ± 0%     42.3MB ± 0%    ~     (p=0.222 n=5+5)
    
    name       old allocs/op   new allocs/op   delta
    Template        391k ± 1%       391k ± 0%    ~     (p=0.841 n=5+5)
    Unicode         320k ± 0%       320k ± 0%    ~     (p=0.841 n=5+5)
    GoTypes        1.14M ± 0%      1.14M ± 0%  +0.26%  (p=0.008 n=5+5)
    SSA            7.60M ± 0%      7.60M ± 0%    ~     (p=0.548 n=5+5)
    Flate           234k ± 0%       234k ± 1%    ~     (p=1.000 n=5+5)
    GoParser        316k ± 1%       317k ± 0%    ~     (p=0.841 n=5+5)
    Reflect         979k ± 0%       980k ± 0%    ~     (p=0.690 n=5+5)
    Tar             251k ± 1%       251k ± 0%    ~     (p=0.595 n=5+5)
    XML             394k ± 0%       393k ± 0%    ~     (p=0.222 n=5+5)
    
    
    Change-Id: I237ae5502db4560f78ce021dc62f6d289797afd6
    Reviewed-on: https://go-review.googlesource.com/39197
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit bfeda6ccc7098c177f351670e3c102974a847377
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 31 16:04:52 2017 -0700

    cmd/compile: add comment to statictmp name generation
    
    Follow-up to review comments on CL 39193.
    
    Change-Id: I7649af9d70ad73e039061a7a66fea416a7476192
    Reviewed-on: https://go-review.googlesource.com/39199
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 8e36575ebe36aba9c42be4f965fa30ec0f2b41dc
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 28 07:12:57 2017 -0700

    cmd/compile: don't mutate shared nodes in orderinit
    
    A few gc.Node ops may be shared across functions.
    The compiler is (mostly) already careful to avoid mutating them.
    However, from a concurrency perspective, replacing (say)
    an empty list with an empty list still counts as a mutation.
    One place this occurs is orderinit. Avoid it.
    
    This requires fixing one spot where shared nodes were mutated.
    It doesn't result in any functional or performance changes.
    
    Passes toolstash-check.
    
    Updates #15756
    
    Change-Id: I63c93b31baeeac62d7574804acb6b7f2bc9d14a9
    Reviewed-on: https://go-review.googlesource.com/39196
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit a8b2e4a630a5991e91095d85c604dc1fa23c1e56
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Thu Mar 30 11:07:36 2017 -0400

    cmd/compile: improve LoweredMove performance on ppc64x
    
    This change improves the performance for LoweredMove on ppc64le
    and ppc64.
    
    benchmark                   old ns/op     new ns/op     delta
    BenchmarkCopyFat8-16        0.93          0.69          -25.81%
    BenchmarkCopyFat12-16       2.61          1.85          -29.12%
    BenchmarkCopyFat16-16       9.68          1.89          -80.48%
    BenchmarkCopyFat24-16       4.48          1.85          -58.71%
    BenchmarkCopyFat32-16       6.12          1.82          -70.26%
    BenchmarkCopyFat64-16       21.2          2.70          -87.26%
    BenchmarkCopyFat128-16      29.6          3.97          -86.59%
    BenchmarkCopyFat256-16      52.6          13.4          -74.52%
    BenchmarkCopyFat512-16      97.1          18.7          -80.74%
    BenchmarkCopyFat1024-16     186           35.3          -81.02%
    
    BenchmarkAssertE2TLarge-16      14.2          5.06          -64.37%
    
    Fixes #19785
    
    Change-Id: I7d5e0052712b75811c02c7d86c5112e5649ad782
    Reviewed-on: https://go-review.googlesource.com/38950
    Reviewed-by: Keith Randall <khr@golang.org>

commit 105cc2bd6396f47bc613721fb6c1db66050e15ab
Author: Russ Cox <rsc@golang.org>
Date:   Fri Mar 31 12:34:25 2017 -0400

    time: test and fix Time.Round, Duration.Round for d > 2⁶²
    
    Round uses r+r < d to decide whether the remainder is
    above or below half of d (to decide whether to round up or down).
    This is wrong when r+r wraps negative, because it looks < d
    but is really > d.
    
    No one will ever care about rounding to a multiple of
    d > 2⁶² (about 146 years), but might as well get it right.
    
    Fixes #19807.
    
    Change-Id: I1b55a742dc36e02a7465bc778bf5dd74fe71f7c0
    Reviewed-on: https://go-review.googlesource.com/39151
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8ab71304d4f02e4280eb5c04422cdec5feb27c11
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 31 12:16:29 2017 -0700

    cmd/compile: use newnamel in typenamesym
    
    The node in typenamesym requires neither
    a position nor a curfn.
    
    Passes toolstash-check.
    
    Updates #15756
    
    Change-Id: I6d39a8961e5578fe5924aaceb29045b6de2699df
    Reviewed-on: https://go-review.googlesource.com/39194
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8caf21da47d09124ba4163b66a99eb08ea72c7c0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 28 15:51:05 2017 -0700

    cmd/compile: use newnamel in ssa.go
    
    For concurrency safety.
    
    Passes toolstash-check.
    
    Updates #15756.
    
    Change-Id: I1caca231a962781ff8f4f589b2e0454d2820ffb6
    Reviewed-on: https://go-review.googlesource.com/39192
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3d90378df5bb97ecadf4a4436fbbf2ca6746a99f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 31 11:10:01 2017 -0700

    cmd/compile: add newnamel, use in tempAt
    
    newnamel is newname but with no dependency on lineno or Curfn.
    This makes it suitable for use in a concurrent back end.
    Use it now to make tempAt global-free.
    
    The decision to push the assignment to n.Name.Curfn
    to the caller of newnamel is based on mdempsky's
    comments in #19683 that he'd like to do that
    for callers of newname as well.
    
    Passes toolstash-check. No compiler performance impact.
    
    Updates #19683
    Updates #15756
    
    Change-Id: Idc461a1716916d268c9ff323129830d9a6e4a4d9
    Reviewed-on: https://go-review.googlesource.com/39191
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 4927b9a9ffeb5e33f6586b0f9000387d8ea20730
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 16:23:01 2017 -0700

    cmd/compile: remove makefuncdatasym_nsym global
    
    This causes a minor reduction in allocations,
    because the old funcdatasym names were
    being interned unnecessarily.
    
    Updates #15756
    
    name       old alloc/op    new alloc/op    delta
    Template      39.9MB ± 0%     39.9MB ± 0%    ~     (p=0.280 n=10+10)
    Unicode       29.9MB ± 0%     29.8MB ± 0%  -0.26%  (p=0.000 n=10+10)
    GoTypes        113MB ± 0%      113MB ± 0%  -0.12%  (p=0.000 n=10+10)
    SSA            855MB ± 0%      855MB ± 0%  -0.03%  (p=0.001 n=10+10)
    Flate         25.4MB ± 0%     25.3MB ± 0%  -0.30%  (p=0.000 n=10+10)
    GoParser      31.9MB ± 0%     31.8MB ± 0%    ~     (p=0.065 n=10+9)
    Reflect       78.4MB ± 0%     78.2MB ± 0%  -0.15%  (p=0.000 n=9+10)
    Tar           26.7MB ± 0%     26.7MB ± 0%  -0.17%  (p=0.000 n=9+10)
    XML           42.3MB ± 0%     42.4MB ± 0%  +0.07%  (p=0.011 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        390k ± 0%       390k ± 0%    ~     (p=0.905 n=9+10)
    Unicode         319k ± 1%       319k ± 1%    ~     (p=0.724 n=10+10)
    GoTypes        1.14M ± 0%      1.14M ± 0%    ~     (p=0.393 n=10+10)
    SSA            7.60M ± 0%      7.60M ± 0%    ~     (p=0.604 n=9+10)
    Flate           235k ± 1%       234k ± 1%    ~     (p=0.105 n=10+10)
    GoParser        317k ± 0%       316k ± 1%    ~     (p=0.280 n=10+10)
    Reflect         979k ± 0%       979k ± 0%    ~     (p=0.315 n=10+10)
    Tar             251k ± 0%       251k ± 1%    ~     (p=0.762 n=8+10)
    XML             393k ± 0%       394k ± 1%    ~     (p=0.095 n=9+10)
    
    name       old text-bytes  new text-bytes  delta
    HelloSize       684k ± 0%       684k ± 0%    ~     (all equal)
    
    name       old data-bytes  new data-bytes  delta
    HelloSize       138k ± 0%       138k ± 0%    ~     (all equal)
    
    name       old exe-bytes   new exe-bytes   delta
    HelloSize      1.03M ± 0%      1.03M ± 0%    ~     (all equal)
    
    Change-Id: Idba33da4e89c325984ac46e4852cf12e4a7fd1a9
    Reviewed-on: https://go-review.googlesource.com/39032
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 2f072a427a9ac20bd6deda18d95e4abb25ea6297
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 31 09:57:38 2017 -0700

    cmd/compile: clean up methodsym
    
    Convert yyerrors into Fatals.
    Remove the goto.
    Move variable declaration closer to use.
    Unify printing strings a bit.
    Convert an int param into a bool.
    
    Passes toolstash-check. No compiler performance impact.
    
    Change-Id: I9017681417b785cf8693d18b124ac4f1ff37f2b5
    Reviewed-on: https://go-review.googlesource.com/39170
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3237af2da88c82d4938b392d2644d65637a11526
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 31 10:28:54 2017 -0700

    cmd/compile: don't use lookupN for statictmps
    
    The names never occur more than once,
    so interning the results is counterproductive.
    
    The impact is not very big, but neither is the fix.
    
    name     old time/op     new time/op     delta
    Unicode     90.2ms ± 3%     88.3ms ± 5%  -2.10%  (p=0.000 n=94+98)
    
    
    Change-Id: I1e3a24433db4ae0c9a6e98166568941824ff0779
    Reviewed-on: https://go-review.googlesource.com/39193
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9e28ea0c08d00fda167f47669711cefd83aa942a
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 31 10:26:30 2017 -0700

    cmd/compile: use std "DO NOT EDIT" comment for generated files
    
    Also: Fix (testdata/gen/) copyGen.go, zeroGen.go, and arithConstGen.go
    to actually match (testdata/) copy.go, zero.go, and arithConst.go, all
    of which were manually edited in https://go-review.googlesource.com/20823
    and https://go-review.googlesource.com/22748 despite the 'do not edit'
    (or perhaps because it was missing in the case of arithConst.go).
    
    For #13560.
    
    Change-Id: I366e1b521e51885e0d318ae848760e5e14ccd488
    Reviewed-on: https://go-review.googlesource.com/39172
    Reviewed-by: Rob Pike <r@golang.org>

commit 654c977b26282e28ee884db579a1c48a9d07af20
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 19:12:36 2017 -0700

    runtime/race: print output when TestRace parsing fails
    
    Change-Id: I986f0c106e059455874692f5bfe2b5af25cf470e
    Reviewed-on: https://go-review.googlesource.com/39090
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0323895cc015aa155013a945caf05f08615ccdcc
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 11:40:06 2017 -0700

    cmd/compile: catch and report nowritebarrier violations later
    
    Prior to this CL, the SSA backend reported violations
    of the //go:nowritebarrier annotation immediately.
    This necessitated emitting errors during SSA compilation,
    which is not compatible with a concurrent backend.
    
    Instead, check for such violations later.
    We already save the data required to do a late check
    for violations of the //go:nowritebarrierrec annotation.
    Use the same data, and check //go:nowritebarrier at the same time.
    
    One downside to doing this is that now only a single
    violation will be reported per function.
    Given that this is for the runtime only,
    and violations are rare, this seems an acceptable cost.
    
    While we are here, remove several 'nerrors != 0' checks
    that are rendered pointless.
    
    Updates #15756
    Fixes #19250 (as much as it ever can be)
    
    Change-Id: Ia44c4ad5b6fd6f804d9f88d9571cec8d23665cb3
    Reviewed-on: https://go-review.googlesource.com/38973
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit ca33e1097155019d0636a0da805054b4182e9b7d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 11:46:45 2017 -0700

    cmd/compile: rework reporting of oversized stack frames
    
    We don't support stack frames over 2GB.
    Rather than detect this during backend compilation,
    check for it at the end of compilation.
    This is arguably a more accurate check anyway,
    since it takes into account the full frame,
    including local stack, arguments, and arch-specific
    rounding, although it's unlikely anyone would ever notice.
    
    Also, rather than reporting the error right away,
    take note of it and report it later, at the top level.
    This is not relevant now, but it will help with making
    the backend concurrent, as the append to the list of
    oversized functions can be cheaply protected by a plain mutex.
    
    Updates #15756
    Updates #19250
    
    Change-Id: Id3fa21906616d62e9dc66e27a17fd5f83304e96e
    Reviewed-on: https://go-review.googlesource.com/38972
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 8577f81a10f4e5b5736119cdd960c44a5ad600f5
Author: Ben Shi <powerman1st@163.com>
Date:   Tue Jan 24 09:48:58 2017 +0000

    cmd/compile/internal: Optimization with RBIT and REV
    
    By checking GOARM in ssa/gen/ARM.rules, each intermediate operator
    can be implemented via different instruction serials.
    
    It is up to the user to choose between compitability and efficiency.
    
    The Bswap32(x) is optimized to REV(x) when GOARM >= 6.
    The CTZ(x) is optimized to CLZ(RBIT x) when GOARM == 7.
    
    Change-Id: Ie9ee645fa39333fa79ad84ed4d1cefac30422814
    Reviewed-on: https://go-review.googlesource.com/35610
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 4a1140472b3093edf9cd653666242393d7f4c2bd
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Thu Mar 30 18:30:07 2017 -0300

    math/big: Unify divWW implementation for ppc64 and ppc64le.
    
    Starting in go1.9, the minimum processor requirement for ppc64 is POWER8. So it
    may now use the same divWW implementation as ppc64le.
    
    Updates #19074
    
    Change-Id: If1a85f175cda89eee06a1024ccd468da6124c844
    Reviewed-on: https://go-review.googlesource.com/39010
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit 5a45a157f2e94cb3fec38a3be8afa3bffd800067
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Thu Mar 30 16:03:03 2017 -0700

    database/sql: support scanning into user defined string types
    
    User defined numeric types such as "type Int int64" have
    been able to be scanned into without a custom scanner by
    using the reflect scan code path used to convert between
    various numeric types. Add in a path for string types
    for symmetry and least surprise.
    
    Fixes #18101
    
    Change-Id: I00553bcf021ffe6d95047eca0067ee94b54ff501
    Reviewed-on: https://go-review.googlesource.com/39031
    Run-TryBot: Daniel Theophanes <kardianos@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit bfd8093c969d2b7b7e1e60866031508ea6e462d6
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Mar 29 14:35:27 2017 +1100

    cmd/asm/internal/arch: use generic obj.Rconv function everywhere
    
    Rather than using arm64.Rconv directly in the archArm64 constructor
    use the generic obj.Rconv helper. This removes the only use of
    arm64.Rconv outside the arm64 package itself.
    
    Change-Id: I99e9e7156b52cd26dc134f610f764ec794264e2c
    Reviewed-on: https://go-review.googlesource.com/38756
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 3e4afe2307c3db0208bc33b06829ec00f8dd6f37
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Mar 30 18:00:10 2017 -0700

    syscall: skip test on TestUnshareMountNameSpace permission error
    
    TestUnshareMountNameSpace fails on arm64 due to permission problems.
    
    Skip that test for now when permission problems are encountered, so we
    don't regress elsewhere in the meantime.
    
    Updates #19698
    
    Change-Id: I9058928afa474b813652c9489f343b8957160a6c
    Reviewed-on: https://go-review.googlesource.com/39052
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 9ffbdabdb02c1fd95eabee82457aaa0dd79d46ac
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 23 21:50:19 2017 -0500

    runtime: make runtime.GC() trigger a concurrent GC
    
    Currently runtime.GC() triggers a STW GC. For common uses in tests and
    benchmarks, it doesn't matter whether it's STW or concurrent, but for
    uses in servers for things like collecting heap profiles and
    controlling memory footprint, this pause can be a bit problem for
    latency.
    
    This changes runtime.GC() to trigger a concurrent GC. In order to
    remain as close as possible to its current meaning, we define it to
    always perform a full mark/sweep GC cycle before returning (even if
    that means it has to finish up a cycle we're in the middle of first)
    and to publish the heap profile as of the triggered mark termination.
    While it must perform a full cycle, simultaneous runtime.GC() calls
    can be consolidated into a single full cycle.
    
    Fixes #18216.
    
    Change-Id: I9088cc5deef4ab6bcf0245ed1982a852a01c44b5
    Reviewed-on: https://go-review.googlesource.com/37520
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 44ed88a5a768a54c1880c58f3438bf2fa8639c4c
Author: Austin Clements <austin@google.com>
Date:   Thu Mar 2 16:28:35 2017 -0500

    runtime: track the number of active sweepone calls
    
    sweepone returns ^uintptr(0) when there are no more spans to *start*
    sweeping, but there may be spans being swept concurrently at the time
    and there's currently no efficient way to tell when the sweeper is
    done sweeping all the spans.
    
    We'll need this for concurrent runtime.GC(), so add a count of the
    number of active sweepone calls to make it possible to block until
    sweeping is truly done.
    
    This is also useful for more accurately printing the gcpacertrace,
    since that should be printed after all of the sweeping stats are in
    (currently we can print it slightly too early).
    
    For #18216.
    
    Change-Id: I06e6240c9e7b40aca6fd7b788bb6962107c10a0f
    Reviewed-on: https://go-review.googlesource.com/37716
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 2919132e1b7ade3faa394339d3199f243103dd72
Author: Austin Clements <austin@google.com>
Date:   Thu Mar 30 14:59:53 2017 -0400

    runtime: don't adjust GC trigger on forced GC
    
    Forced GCs don't provide good information about how to adjust the GC
    trigger. Currently we avoid adjusting the trigger on forced GC because
    forced GC is also STW and we don't adjust the trigger on STW GC.
    However, this will become a problem when forced GC is concurrent.
    
    Fix this by skipping trigger adjustment if the GC was user-forced.
    
    For #18216.
    
    Change-Id: I03dfdad12ecd3cfeca4573140a0768abb29aac5e
    Reviewed-on: https://go-review.googlesource.com/38951
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 29fdbcfea37f3bf519f678b1426277b70406c029
Author: Austin Clements <austin@google.com>
Date:   Mon Feb 27 10:46:12 2017 -0500

    runtime: track forced GCs independent of gcMode
    
    Currently gcMode != gcBackgroundMode implies this was a user-forced GC
    cycle. This is no longer going to be true when we make runtime.GC()
    trigger a concurrent GC, so replace this with an explicit
    work.userForced bit.
    
    For #18216.
    
    Change-Id: If7d71bbca78b5f0b35641b070f9d457f5c9a52bd
    Reviewed-on: https://go-review.googlesource.com/37519
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 786eb5b754780d649f5d85481096d268728d383f
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 23 21:55:37 2017 -0500

    runtime: make debug.FreeOSMemory call runtime.GC()
    
    Currently freeOSMemory calls gcStart directly, but we really just want
    it to behave like runtime.GC() and then perform a scavenge, so make it
    call runtime.GC() rather than gcStart.
    
    For #18216.
    
    Change-Id: I548ec007afc788e87d383532a443a10d92105937
    Reviewed-on: https://go-review.googlesource.com/37518
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 3d58498fdb047c442abafff4277159d6c3842e05
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 23 11:54:43 2017 -0500

    runtime: simplify forced GC triggering
    
    Now that the gcMode is no longer involved in the GC trigger condition,
    we can simplify the triggering of forced GCs. By making the trigger
    condition for forced GCs true even if gcphase is not _GCoff, we don't
    need any special case path in gcStart to ensure that forced GCs don't
    get consolidated.
    
    Change-Id: I6067a13d76e40ff2eef8fade6fc14adb0cb58ee5
    Reviewed-on: https://go-review.googlesource.com/37517
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 29be3f1999d1ccf01110f4e1f2d628f54f6c65da
Author: Austin Clements <austin@google.com>
Date:   Mon Jan 9 11:35:42 2017 -0500

    runtime: generalize GC trigger
    
    Currently the GC triggering condition is an awkward combination of the
    gcMode (whether or not it's gcBackgroundMode) and a boolean
    "forceTrigger" flag.
    
    Replace this with a new gcTrigger type that represents the range of
    transition predicates we need. This has several advantages:
    
    1. We can remove the awkward logic that affects the trigger behavior
       based on the gcMode. Now gcMode purely controls whether to run a
       STW GC or not and the gcTrigger controls whether this is a forced
       GC that cannot be consolidated with other GC cycles.
    
    2. We can lift the time-based triggering logic in sysmon to just
       another type of GC trigger and move the logic to the trigger test.
    
    3. This sets us up to have a cycle count-based trigger, which we'll
       use to make runtime.GC trigger concurrent GC with the desired
       consolidation properties.
    
    For #18216.
    
    Change-Id: If9cd49349579a548800f5022ae47b8128004bbfc
    Reviewed-on: https://go-review.googlesource.com/37516
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 640cd3b3220f7a06820365e53f8fc6cb0acd1b20
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 23 11:04:37 2017 -0500

    runtime: check transition condition before triggering periodic GC
    
    Currently sysmon triggers periodic GC if GC is not currently running
    and it's been long enough since the last GC. This misses some
    important conditions; for example, whether GC is enabled at all by
    GOGC. As a result, if GOGC is off, once we pass the timeout for
    periodic GC, sysmon will attempt to trigger a GC every 10ms. This GC
    will be a no-op because gcStart will check all of the appropriate
    conditions and do nothing, but it still goes through the motions of
    waking the forcegc goroutine and printing a gctrace line.
    
    Fix this by making sysmon call gcShouldStart to check *all* of the
    appropriate transition conditions before attempting to trigger a
    periodic GC.
    
    Fixes #19247.
    
    Change-Id: Icee5521ce175e8419f934723849853d53773af31
    Reviewed-on: https://go-review.googlesource.com/37515
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 1be3e76e7628cae8500c0c1f3aa620638aec351d
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 1 21:03:20 2017 -0500

    runtime: simplify heap profile flushing
    
    Currently the heap profile is flushed by *either* gcSweep in STW mode
    or by gcMarkTermination in concurrent mode. Simplify this by making
    gcMarkTermination always flush the heap profile and by making gcSweep
    do one extra flush (instead of two) in STW mode.
    
    Change-Id: I62147afb2a128e1f3d92ef4bb8144c8a345f53c4
    Reviewed-on: https://go-review.googlesource.com/37715
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit eee85fc5a1071c8c7d8d804e898da6394ffbdec3
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 1 13:58:22 2017 -0500

    runtime: snapshot heap profile during mark termination
    
    Currently we snapshot the heap profile just *after* mark termination
    starts the world because it's a relatively expensive operation.
    However, this means any alloc or free events that happen between
    starting the world and snapshotting the heap profile can be accounted
    to the wrong cycle. In the worst case, a free can be accounted to the
    cycle before the alloc; if the heap is small, this can result
    temporarily in a negative "in use" count in the profile.
    
    Fix this without making STW more expensive by using a global heap
    profile cycle counter. This lets us split up the operation into a two
    parts: 1) a super-cheap snapshot operation that simply increments the
    global cycle counter during STW, and 2) a more expensive cleanup
    operation we can do after starting the world that frees up a slot in
    all buckets for use by the next heap profile cycle.
    
    Fixes #19311.
    
    Change-Id: I6bdafabf111c48b3d26fe2d91267f7bef0bd4270
    Reviewed-on: https://go-review.googlesource.com/37714
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 3ebe7d7d110ba2573cfd597bb786d14749181cc3
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 1 11:50:38 2017 -0500

    runtime: pull heap profile cycle into a type
    
    Currently memRecord has the same set of four fields repeated three
    times. Pull these into a type and use this type three times. This
    cleans up and simplifies the code a bit and will make it easier to
    switch to a globally tracked heap profile cycle for #19311.
    
    Change-Id: I414d15673feaa406a8366b48784437c642997cf2
    Reviewed-on: https://go-review.googlesource.com/37713
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 42aa608f8acd50f963b0f1807a6c3ecbf651b6e2
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 30 16:40:52 2017 -0700

    cmd/compile: remove confusing comment, fix comment for symExport
    
    The symExport flag tells whether a symbol is in the export list
    already or not (and it's also used to avoid being added to that
    list). Exporting is based on that export list - no need to check
    again.
    
    Change-Id: I6056f97aa5c24a19376957da29199135c8da35f9
    Reviewed-on: https://go-review.googlesource.com/39033
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit 673a8fdfe60929f61657dbfbdf5534eabe8cd6f5
Author: Austin Clements <austin@google.com>
Date:   Mon Feb 27 11:36:37 2017 -0500

    runtime: diagram flow of stats through heap profile
    
    Every time I modify heap profiling, I find myself redrawing this
    diagram, so add it to the comments. This shows how allocations and
    frees are accounted, how we arrive at consistent profile snapshots,
    and when those snapshots are published to the user.
    
    Change-Id: I106aba1200af3c773b46e24e5f50205e808e2c69
    Reviewed-on: https://go-review.googlesource.com/37514
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit ef1829d1debc0bcd32052d9686adb75704e75984
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 23 21:48:34 2017 -0500

    runtime: improve TestMemStats checks
    
    Now that we have a nice predicate system, improve the tests performed
    by TestMemStats. We add some more non-zero checks (now that we force a
    GC, things like NumGC must be non-zero), checks for trivial boolean
    fields, and a few more range checks.
    
    Change-Id: I6da46d33fa0ce5738407ee57d587825479413171
    Reviewed-on: https://go-review.googlesource.com/37513
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit bda74b0e4a8eb1acae8c202efd62d298ec3268f0
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 23 21:40:55 2017 -0500

    runtime: make TestMemStats failure messages useful
    
    Currently most TestMemStats failures dump the whole MemStats object if
    anything is amiss without telling you what is amiss, or even which
    field is wrong. This makes it hard to figure out what the actual
    problem is.
    
    Replace this with a reflection walk over MemStats and a map of
    predicates to check. If one fails, we can construct a detailed and
    descriptive error message. The predicates are a direct translation of
    the current tests.
    
    Change-Id: I5a7cafb8e6a1eeab653d2e18bb74e2245eaa5444
    Reviewed-on: https://go-review.googlesource.com/37512
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit bc972e8ef870471c3c4ba95b92c5194b37ba2871
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 30 15:44:12 2017 -0700

    cmd/compile: remove lookupf
    
    Change-Id: I4de5173fa50fbf90802d1d2428824702f2118dde
    Reviewed-on: https://go-review.googlesource.com/39030
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 361af94d5d4df310a90d924b4ce3dd4e3c96ee38
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Mar 29 15:14:55 2017 +1100

    cmd/internal/obj, cmd/link: remove Hwindowsgui everywhere
    
    Hwindowsgui has the same meaning as Hwindows - build PE
    executable. So use Hwindows everywhere.
    
    Change-Id: I2cae5777f17c7bc3a043dfcd014c1620cc35fc20
    Reviewed-on: https://go-review.googlesource.com/38761
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4e3a1e409ae1bf74deb3c93745ad49b71311f4de
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Mon Mar 27 15:58:14 2017 +1100

    cmd/link/internal/ld: introduce and use windowsgui variable
    
    cmd/link -H flag is stored in variable of type
    cmd/internal/obj.HeadType. The HeadType type from cmd/internal/obj
    accepts Hwindows and Hwindowsgui values, but these values have
    same meaning - build PE executable, except for 2 places in
    cmd/link/internal/ld package.
    
    This CL introduces code to store cmd/link "windowsgui" -H flag
    in cmd/link/internal/ld, so cmd/internal/obj.Hwindowsgui can be
    removed in the next CL.
    
    This CL also includes 2 changes to code where distinction
    between Hwindows and Hwindowsgui is important.
    
    Change-Id: Ie5ee1f374e50c834652a037f2770118d56c21a2a
    Reviewed-on: https://go-review.googlesource.com/38760
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 94c62efe9cfb9b3d9f0934dcb6d40a0ae522cdff
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 15:07:05 2017 -0700

    cmd/link: skip TestDWARF when cgo is disabled
    
    While we're here, fix a Skip/Skipf error I noticed.
    
    Fixes #19796.
    
    Change-Id: I59b1f5b5ea727fc314acfee8445b3de0b5af1e46
    Reviewed-on: https://go-review.googlesource.com/38992
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ac99ade5a09a685f20df3b0d62c98cd3de7a575e
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 30 13:19:18 2017 -0700

    cmd/compile: remove Pkglookup in favor of Lookup
    
    Remove one of the many lookup variants.
    
    Change-Id: I4095aa030da4227540badd6724bbf50b728fbe93
    Reviewed-on: https://go-review.googlesource.com/38990
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit f7027b4b2dc9c822efd94f1d84189a60291ae152
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 30 14:04:54 2017 -0700

    cmd/compile: remove lookupBytes
    
    Change-Id: I08c264f5f3744d835e407534c492ef8c43e1a700
    Reviewed-on: https://go-review.googlesource.com/38991
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7efb0779be523f8629919abc90067382f370f4c0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 14:34:16 2017 -0700

    cmd/compile: remove scratchFpMem global
    
    Instead, add a scratchFpMem field to ssafn,
    so that it may be passed on to genssa.
    
    Updates #15756
    
    Change-Id: Icdeae290d3098d14d31659fa07a9863964bb76ed
    Reviewed-on: https://go-review.googlesource.com/38728
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 399086f2db0257f9f90a87dfd8a251944b52120a
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Mar 30 15:58:01 2017 -0400

    cmd/compile/internal/ssa/gen: add comment on SB-addressing on s390x
    
    During the review of CL 38801 it was noted that it would be nice
    to have a bit more clarity on how-and-why SB addressing is handled
    strangely on s390x. This additional comment should hopefully help.
    
    In general SB is handled differently because not all instructions
    have variants that use relative addressing.
    
    Change-Id: I3379012ae3f167478c191c435939c3b876c645ed
    Reviewed-on: https://go-review.googlesource.com/38952
    Reviewed-by: Keith Randall <khr@golang.org>

commit ebcb9cdf67e53d2ad50da1910d6c5f944c32b163
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 11:16:23 2017 -0700

    cmd/compile: cull unused ssa construction Error function
    
    The preceding passes have caught any errors
    that could occur during SSA construction.
    
    Updates #19250
    
    Change-Id: I736edb2017da3f111fb9f74be12d437b5a24d2b4
    Reviewed-on: https://go-review.googlesource.com/38971
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit beb833f12456f8ae49e003df79f1e0813b610032
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 11:13:01 2017 -0700

    cmd/compile: initialize SSA runtime functions in initssaconfig
    
    This is a better home for it.
    
    Change-Id: I7ce96c16378d841613edaa53c07347b0ac99ea6e
    Reviewed-on: https://go-review.googlesource.com/38970
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ed6f2978a21d39abd1efbc2b0c06431e17519a63
Author: David du Colombier <0intro@gmail.com>
Date:   Thu Mar 30 21:33:06 2017 +0200

    cmd/link: skip TestDWARF on Plan 9
    
    TestDWARF has been added in CL 38855. This test is
    failing on Plan 9 because executables don't have
    a DWARF symbol table.
    
    Fixes #19793.
    
    Change-Id: I7fc547a7c877b58cc4ff6b4eb5b14852e8b4668b
    Reviewed-on: https://go-review.googlesource.com/38931
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e82c925f5e3af26ded71e85dec1eea3464e4fa19
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 29 21:04:00 2017 -0700

    cmd/compile: remove Type haspointers caches
    
    Even very large Types are not very big.
    The haspointer cache looks like premature optimization.
    Removing them has no detectable compiler performance impact,
    and it removes mutable shared state used by the backend.
    
    Updates #15756
    
    Change-Id: I2d2cf03f470f5eef5bcd50ff693ef6a01d481700
    Reviewed-on: https://go-review.googlesource.com/38912
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e86168430f0aab8f971763e4b00c2aae7bec55f0
Author: Egon Elbre <egonelbre@gmail.com>
Date:   Fri Feb 10 11:30:22 2017 +0200

    cmd/fix,cmd/gofmt: flush to disk before diffing
    
    Flush file content to disk before diffing files,
    may cause unpredictable results on Windows.
    
    Convert from \r\n to \n when comparing diff result.
    
    Change-Id: Ibcd6154a2382dba1338ee5674333611aea16bb65
    Reviewed-on: https://go-review.googlesource.com/36750
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>

commit c5ddc558baa9884050ddf26dd93c91e9297509b8
Author: Ben Shi <powerman1st@163.com>
Date:   Mon Jan 23 09:14:35 2017 +0000

    cmd/internal/obj/arm: support more ARMv5/ARMv6/ARMv7 instructions
    
    REV/REV16/REVSH were introduced in ARMv6, they offered more efficient
    byte reverse operatons.
    
    MMUL/MMULA/MMULS were introduced in ARMv6, they simplified
    a serial of mul->shift->add/sub operations into a single instruction.
    
    RBIT was introduced in ARMv7, it inversed a 32-bit word's bit order.
    
    MULS was introduced in ARMv7, it corresponded to MULA.
    
    MULBB/MULABB were introduced in ARMv5TE, they performed 16-bit
    multiplication (and accumulation).
    
    Change-Id: I6365b17b3c4eaf382a657c210bb0094b423b11b8
    Reviewed-on: https://go-review.googlesource.com/35565
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 944d56d7635f1f5e599805d621cd3a171534e333
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 29 21:03:09 2017 -0700

    cmd/compile: move haspointers to type.go
    
    100% code movement.
    
    Change-Id: Idb51c61b7363229258a3b48045e901bea68c7a85
    Reviewed-on: https://go-review.googlesource.com/38911
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 371aa23d104237fe72d84618210c5d61da8019aa
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 29 21:00:55 2017 -0700

    cmd/compile: convert Type.Trecur to a boolean flag
    
    Change-Id: I162e86e5f92c8b827a74ee860d16abadf83bc43e
    Reviewed-on: https://go-review.googlesource.com/38910
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit efc47819c080c800a161f099b7bdbacb53ea311e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 29 16:28:51 2017 -0700

    cmd/compile: eliminate use of Trecur in formatting routines
    
    CL 38147 eliminated package gc globals in formatting routines.
    However, tconv still used the Type field Trecur
    to avoid infinite recursion when formatting recursive
    interfaces types such as (test/fixedbugs398.go):
    
    type i1 interface {
            F() interface {
                    i1
            }
    }
    
    type i2 interface {
            F() interface {
                    i2
            }
    }
    
    This CL changes the recursion prevention to use a parameter,
    and threads it through the formatting routines.
    Because this fundamentally limits the embedding depth
    of all types, it sets the depth limit to be much higher.
    In practice, it is unlikely to impact any code at all,
    one way or the other.
    
    The remaining uses of Type.Trecur are boolean in nature.
    A future CL will change Type.Trecur to be a boolean flag.
    
    The removal of a couple of mode.Sprintf calls
    makes this a very minor net performance improvement:
    
    name       old alloc/op    new alloc/op    delta
    Template      40.0MB ± 0%     40.0MB ± 0%  -0.13%  (p=0.032 n=5+5)
    Unicode       30.0MB ± 0%     29.9MB ± 0%    ~     (p=0.310 n=5+5)
    GoTypes        114MB ± 0%      113MB ± 0%  -0.25%  (p=0.008 n=5+5)
    SSA            856MB ± 0%      855MB ± 0%  -0.04%  (p=0.008 n=5+5)
    Flate         25.5MB ± 0%     25.4MB ± 0%  -0.27%  (p=0.008 n=5+5)
    GoParser      31.9MB ± 0%     31.9MB ± 0%    ~     (p=0.222 n=5+5)
    Reflect       79.0MB ± 0%     78.6MB ± 0%  -0.45%  (p=0.008 n=5+5)
    Tar           26.8MB ± 0%     26.7MB ± 0%  -0.25%  (p=0.032 n=5+5)
    XML           42.4MB ± 0%     42.4MB ± 0%    ~     (p=0.151 n=5+5)
    
    name       old allocs/op   new allocs/op   delta
    Template        395k ± 0%       391k ± 0%  -1.00%  (p=0.008 n=5+5)
    Unicode         321k ± 1%       319k ± 0%  -0.56%  (p=0.008 n=5+5)
    GoTypes        1.16M ± 0%      1.14M ± 0%  -1.61%  (p=0.008 n=5+5)
    SSA            7.63M ± 0%      7.60M ± 0%  -0.30%  (p=0.008 n=5+5)
    Flate           239k ± 0%       234k ± 0%  -1.94%  (p=0.008 n=5+5)
    GoParser        320k ± 0%       317k ± 1%  -0.86%  (p=0.008 n=5+5)
    Reflect        1.00M ± 0%      0.98M ± 0%  -2.17%  (p=0.016 n=4+5)
    Tar             255k ± 1%       251k ± 0%  -1.35%  (p=0.008 n=5+5)
    XML             398k ± 0%       395k ± 0%  -0.89%  (p=0.008 n=5+5)
    
    Updates #15756
    
    Change-Id: Id23e647d347aa841f9a69d51f7d2d7d27b259239
    Reviewed-on: https://go-review.googlesource.com/38797
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 18e77673d85fe8d230b8c2dfffe7cd5341b72c4f
Author: Russ Cox <rsc@golang.org>
Date:   Wed Mar 29 20:50:34 2017 -0400

    cmd/link: emit a mach-o dwarf segment that dsymutil will accept
    
    Right now, at least with Xcode 8.3, we invoke dsymutil and dutifully
    copy what it produces back into the binary, but it has actually dropped
    all the DWARF information that we wanted, because it didn't like
    the look of go.o.
    
    Make it like the look of go.o.
    
    DWARF is tested in other ways, but typically indirectly and not for cgo programs.
    Add a direct test, and one that exercises cgo.
    This detects missing dwarf information in cgo-using binaries on macOS,
    at least with Xcode 8.3, and possibly earlier versions as well.
    
    Fixes #19772.
    
    Change-Id: I0082e52c0bc8fc4e289770ec3dc02f39fd61e743
    Reviewed-on: https://go-review.googlesource.com/38855
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit eca90561c39d4d8cc587fbcc70baa9d90f3f0707
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 07:38:46 2017 -0700

    cmd/compile: minor init handling cleanup
    
    Place comments correctly.
    Simplify control flow.
    Reduce variable scope.
    
    Passes toolstash-check.
    
    Change-Id: Iea47ed3502c15491c2ca6db8149fe0949b8849aa
    Reviewed-on: https://go-review.googlesource.com/38914
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5272a2cdc551c041a9f744ede72506be5f622196
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 30 06:45:36 2017 -0700

    cmd/compile: avoid infinite loops in dead blocks during phi insertion
    
    Now that we no longer generate dead code,
    it is possible to follow block predecessors
    into infinite loops with no variable definitions,
    causing an infinite loop during phi insertion.
    
    To fix that, check explicitly whether the predecessor
    is dead in lookupVarOutgoing, and if so, bail.
    
    The loop in lookupVarOutgoing is very hot code,
    so I am wary of adding anything to it.
    However, a long, CPU-only benchmarking run shows no
    performance impact at all.
    
    Fixes #19783
    
    Change-Id: I8ef8d267e0b20a29b5cb0fecd7084f76c6f98e47
    Reviewed-on: https://go-review.googlesource.com/38913
    Reviewed-by: David Chase <drchase@google.com>

commit 3431d9113c7e01a0ddb458a075ca571d3873e061
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 28 13:52:14 2017 -0700

    cmd/compile: add global autogeneratedPos
    
    We use an "autogenerated" position in several places.
    Rather than recreate it each time, make one early on and reuse it.
    This removes the creation of new positions during the backend,
    which was not concurrency-safe.
    
    Updates #15756
    
    Change-Id: Ic116b2e60f0e99de1a2ea87fe763831b50b645f8
    Reviewed-on: https://go-review.googlesource.com/38915
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b28f2f73991c849657e5898dfa2f8cae54e885a0
Author: Russ Cox <rsc@golang.org>
Date:   Wed Mar 29 20:53:32 2017 -0400

    cmd/link: make mach-o dwarf segment properly aligned
    
    Without this, the load fails during kernel exec, which results in the
    mysterious and completely uninformative "Killed: 9" error.
    
    It appears that the stars (or at least the inputs) were properly aligned
    with earlier versions of Xcode so that this happened accidentally.
    Make it happen on purpose.
    
    Gregory Man bisected the breakage to this change in LLVM,
    which fits the theory nicely:
    https://github.com/llvm-mirror/llvm/commit/9a41e59c
    
    Fixes #19734.
    
    Change-Id: Ice67a09af2de29d3c0d5e3fcde6a769580897c95
    Reviewed-on: https://go-review.googlesource.com/38854
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 23dc3633ca38b8167a50250eeec948e4cd619056
Author: Russ Cox <rsc@golang.org>
Date:   Wed Mar 29 20:46:20 2017 -0400

    cmd/link: disable mach-o dwarf munging with -w (in addition to -s)
    
    Might as well provide a way around the mach-o munging
    that doesn't require stripping all symbols.
    After all, -w does mean no DWARF.
    
    For #11887, #19734, and anyone else that needs to disable
    this code path without losing the symbol table.
    
    Change-Id: I254b7539f97fb9211fa90f446264b383e7f3980f
    Reviewed-on: https://go-review.googlesource.com/38853
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 371c83b594bf7af5a5b56f3ecfe1347e0b666dc2
Author: Russ Cox <rsc@golang.org>
Date:   Wed Mar 29 20:39:32 2017 -0400

    cmd/link: do not pass -s through to host linker on macOS
    
    This keeps the host linker from printing
    ld: warning: option -s is obsolete and being ignored
    
    Fixes #19775.
    
    Change-Id: I18dd4e4b3f59cbf35dad770fd65e6baea5a7347f
    Reviewed-on: https://go-review.googlesource.com/38851
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6e7d5d0326a50b06fd47553cb2d96ede28b1680c
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Mon Mar 27 15:55:15 2017 +1100

    debug/pe: add TestBuildingWindowsGUI
    
    Change-Id: I6b6a6dc57e48e02ff0d452755b8dcf5543b3caed
    Reviewed-on: https://go-review.googlesource.com/38759
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 23f56c186d5a1dc198bdbb597b834ce208f09485
Author: Elias Naur <elias.naur@gmail.com>
Date:   Thu Mar 30 01:23:20 2017 +0200

    misc/cgo/testcshared: use the gold linker on android/arm64
    
    The gold linker is used by default in the Android NDK, except on
    arm64:
    
    https://github.com/android-ndk/ndk/issues/148
    
    The Go linker already forces the use of the gold linker on arm and
    arm64 (CL 22141) for other reasons. However, the test.bash script in
    testcshared doesn't, resulting in linker errors on android/arm64:
    
    warning: liblog.so, needed by ./libgo.so, not found (try using -rpath or
    -rpath-link)
    
    Add -fuse-ld=gold when running testcshared on Android. Fixes the
    android/arm64 builder.
    
    Change-Id: I35ca96f01f136bae72bec56d71b7ca3f344df1ed
    Reviewed-on: https://go-review.googlesource.com/38832
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 592037f38188a7357f2103c8ddf1bd62206c778f
Author: Caleb Spare <cespare@gmail.com>
Date:   Wed Mar 29 15:18:54 2017 -0700

    runtime: fix for implementation notes appearing in godoc
    
    Change-Id: I31cfae1e98313b68e3bc8f49079491d2725a662b
    Reviewed-on: https://go-review.googlesource.com/38850
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4f579cc65bfedca93ad00f9c5eb138c4a785a4ad
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Tue Mar 28 14:09:28 2017 -0500

    math: speed up Log on amd64
    
    After https://golang.org/cl/31490 we break false
    output dependency for CVTS.. in compiler generated code.
    I've looked through asm code, which uses CVTS..
    and added XOR to the only case where it affected performance.
    
    Log-6                  21.6ns ± 0%  19.9ns ± 0%  -7.87%  (p=0.000 n=10+10)
    
    Change-Id: I25d9b405e3041a3839b40f9f9a52e708034bb347
    Reviewed-on: https://go-review.googlesource.com/38771
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 41fd8d6401b6702b801678a493001ec9e08d6f2d
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Mar 28 15:10:20 2017 -0400

    cmd/internal/obj: make morestack cutoff the same on all architectures
    
    There is always 128 bytes available below the stackguard. Allow functions
    with medium-sized stack frames to use this, potentially allowing them to
    avoid growing the stack.
    
    This change makes all architectures use the same calculation as x86.
    
    Change-Id: I2afb1a7c686ae5a933e50903b31ea4106e4cd0a0
    Reviewed-on: https://go-review.googlesource.com/38734
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit fbe6723903cc7ec06d0158d4909c6cb15c1ff977
Author: haya14busa <haya14busa@gmail.com>
Date:   Tue Mar 28 13:24:17 2017 +0900

    regexp: reduce allocations at makeOnePass
    
    It reduces needless allocations on compiling onepass regex.
    
    Following CL 38750
    
    name                                      old time/op    new time/op    delta
    CompileOnepass/^(?:(?:(?:.(?:$))?))...-4    5.75µs ± 1%    5.51µs ± 2%   -4.25%  (p=0.008 n=5+5)
    CompileOnepass/^abcd$-4                     4.76µs ± 0%    4.52µs ± 1%   -5.06%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a{0,})*?)$-4          5.56µs ± 0%    5.56µs ± 3%     ~     (p=0.524 n=5+5)
    CompileOnepass/^(?:(?:a+)*)$-4              5.09µs ± 0%    5.15µs ± 5%     ~     (p=0.690 n=5+5)
    CompileOnepass/^(?:(?:a|(?:aa)))$-4         6.53µs ± 0%    6.43µs ± 5%     ~     (p=0.151 n=5+5)
    CompileOnepass/^(?:[^\s\S])$-4              4.05µs ± 1%    4.00µs ± 2%     ~     (p=0.095 n=5+5)
    CompileOnepass/^(?:(?:(?:a*)+))$-4          5.47µs ± 0%    5.36µs ± 1%   -1.91%  (p=0.008 n=5+5)
    CompileOnepass/^[a-c]+$-4                   4.13µs ± 1%    4.05µs ± 0%   -2.07%  (p=0.008 n=5+5)
    CompileOnepass/^[a-c]*$-4                   4.59µs ± 2%    4.93µs ± 7%   +7.30%  (p=0.016 n=5+5)
    CompileOnepass/^(?:a*)$-4                   4.67µs ± 1%    4.82µs ± 8%     ~     (p=0.730 n=4+5)
    CompileOnepass/^(?:(?:aa)|a)$-4             6.43µs ± 1%    6.18µs ± 1%   -3.91%  (p=0.008 n=5+5)
    CompileOnepass/^...$-4                      4.71µs ± 0%    4.31µs ± 1%   -8.51%  (p=0.008 n=5+5)
    CompileOnepass/^(?:a|(?:aa))$-4             6.37µs ± 0%    6.17µs ± 0%   -3.23%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))c$-4                  6.85µs ± 1%    6.50µs ± 1%   -5.15%  (p=0.008 n=5+5)
    CompileOnepass/^a.[l-nA-Cg-j]?e$-4          6.99µs ± 1%    6.66µs ± 1%   -4.81%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))$-4                   6.15µs ± 1%    5.87µs ± 0%   -4.57%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:(b)|(c))c$-4            8.62µs ± 1%    8.21µs ± 1%   -4.77%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b|c)$-4                 5.76µs ±42%    4.42µs ± 1%  -23.35%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c)$-4                7.17µs ± 6%    6.86µs ± 0%   -4.39%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c+)$-4               8.08µs ± 2%    7.67µs ± 2%   -4.97%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:bc)+$-4                 5.53µs ± 3%    5.35µs ± 1%   -3.34%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:[bcd])+$-4              5.08µs ± 1%    4.98µs ± 0%   -2.02%  (p=0.008 n=5+5)
    CompileOnepass/^a((?:[bcd])+)$-4            6.49µs ± 1%    6.29µs ± 1%   -3.03%  (p=0.008 n=5+5)
    CompileOnepass/^a(:?b|c)*d$-4               11.8µs ± 1%    11.4µs ± 3%   -3.98%  (p=0.008 n=5+5)
    CompileOnepass/^.bc(d|e)*$-4                8.02µs ± 1%    7.54µs ± 1%   -6.00%  (p=0.008 n=5+5)
    CompileOnepass/^loooooooooooooooooo...-4     228µs ±18%     196µs ± 0%  -14.02%  (p=0.016 n=5+4)
    
    name                                      old alloc/op   new alloc/op   delta
    CompileOnepass/^(?:(?:(?:.(?:$))?))...-4    3.41kB ± 0%    3.38kB ± 0%   -0.94%  (p=0.008 n=5+5)
    CompileOnepass/^abcd$-4                     2.75kB ± 0%    2.74kB ± 0%   -0.29%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a{0,})*?)$-4          3.34kB ± 0%    3.34kB ± 0%     ~     (all equal)
    CompileOnepass/^(?:(?:a+)*)$-4              2.95kB ± 0%    2.95kB ± 0%     ~     (all equal)
    CompileOnepass/^(?:(?:a|(?:aa)))$-4         3.75kB ± 0%    3.74kB ± 0%   -0.43%  (p=0.008 n=5+5)
    CompileOnepass/^(?:[^\s\S])$-4              2.46kB ± 0%    2.45kB ± 0%   -0.49%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:(?:a*)+))$-4          3.13kB ± 0%    3.13kB ± 0%     ~     (all equal)
    CompileOnepass/^[a-c]+$-4                   2.48kB ± 0%    2.48kB ± 0%     ~     (all equal)
    CompileOnepass/^[a-c]*$-4                   2.52kB ± 0%    2.52kB ± 0%     ~     (all equal)
    CompileOnepass/^(?:a*)$-4                   2.63kB ± 0%    2.63kB ± 0%     ~     (all equal)
    CompileOnepass/^(?:(?:aa)|a)$-4             3.64kB ± 0%    3.62kB ± 0%   -0.44%  (p=0.008 n=5+5)
    CompileOnepass/^...$-4                      2.91kB ± 0%    2.87kB ± 0%   -1.37%  (p=0.008 n=5+5)
    CompileOnepass/^(?:a|(?:aa))$-4             3.64kB ± 0%    3.62kB ± 0%   -0.44%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))c$-4                  4.39kB ± 0%    4.38kB ± 0%   -0.18%  (p=0.008 n=5+5)
    CompileOnepass/^a.[l-nA-Cg-j]?e$-4          4.32kB ± 0%    4.30kB ± 0%   -0.56%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))$-4                   4.06kB ± 0%    4.05kB ± 0%   -0.39%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:(b)|(c))c$-4            5.31kB ± 0%    5.30kB ± 0%   -0.15%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b|c)$-4                 2.88kB ± 0%    2.87kB ± 0%   -0.28%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c)$-4                4.36kB ± 0%    4.35kB ± 0%   -0.18%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c+)$-4               4.59kB ± 0%    4.58kB ± 0%   -0.17%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:bc)+$-4                 3.15kB ± 0%    3.15kB ± 0%     ~     (all equal)
    CompileOnepass/^a(?:[bcd])+$-4              2.94kB ± 0%    2.94kB ± 0%     ~     (all equal)
    CompileOnepass/^a((?:[bcd])+)$-4            4.09kB ± 0%    4.08kB ± 0%   -0.20%  (p=0.008 n=5+5)
    CompileOnepass/^a(:?b|c)*d$-4               6.15kB ± 0%    6.10kB ± 0%   -0.78%  (p=0.008 n=5+5)
    CompileOnepass/^.bc(d|e)*$-4                4.47kB ± 0%    4.46kB ± 0%   -0.36%  (p=0.008 n=5+5)
    CompileOnepass/^loooooooooooooooooo...-4     135kB ± 0%     135kB ± 0%     ~     (p=0.810 n=5+5)
    
    name                                      old allocs/op  new allocs/op  delta
    CompileOnepass/^(?:(?:(?:.(?:$))?))...-4      49.0 ± 0%      47.0 ± 0%   -4.08%  (p=0.008 n=5+5)
    CompileOnepass/^abcd$-4                       41.0 ± 0%      41.0 ± 0%     ~     (all equal)
    CompileOnepass/^(?:(?:a{0,})*?)$-4            49.0 ± 0%      49.0 ± 0%     ~     (all equal)
    CompileOnepass/^(?:(?:a+)*)$-4                44.0 ± 0%      44.0 ± 0%     ~     (all equal)
    CompileOnepass/^(?:(?:a|(?:aa)))$-4           54.0 ± 0%      54.0 ± 0%     ~     (all equal)
    CompileOnepass/^(?:[^\s\S])$-4                33.0 ± 0%      33.0 ± 0%     ~     (all equal)
    CompileOnepass/^(?:(?:(?:a*)+))$-4            46.0 ± 0%      46.0 ± 0%     ~     (all equal)
    CompileOnepass/^[a-c]+$-4                     36.0 ± 0%      36.0 ± 0%     ~     (all equal)
    CompileOnepass/^[a-c]*$-4                     41.0 ± 0%      41.0 ± 0%     ~     (all equal)
    CompileOnepass/^(?:a*)$-4                     42.0 ± 0%      42.0 ± 0%     ~     (all equal)
    CompileOnepass/^(?:(?:aa)|a)$-4               53.0 ± 0%      53.0 ± 0%     ~     (all equal)
    CompileOnepass/^...$-4                        43.0 ± 0%      39.0 ± 0%   -9.30%  (p=0.008 n=5+5)
    CompileOnepass/^(?:a|(?:aa))$-4               53.0 ± 0%      53.0 ± 0%     ~     (all equal)
    CompileOnepass/^a((b))c$-4                    53.0 ± 0%      53.0 ± 0%     ~     (all equal)
    CompileOnepass/^a.[l-nA-Cg-j]?e$-4            58.0 ± 0%      56.0 ± 0%   -3.45%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))$-4                     47.0 ± 0%      47.0 ± 0%     ~     (all equal)
    CompileOnepass/^a(?:(b)|(c))c$-4              65.0 ± 0%      65.0 ± 0%     ~     (all equal)
    CompileOnepass/^a(?:b|c)$-4                   40.0 ± 0%      40.0 ± 0%     ~     (all equal)
    CompileOnepass/^a(?:b?|c)$-4                  57.0 ± 0%      57.0 ± 0%     ~     (all equal)
    CompileOnepass/^a(?:b?|c+)$-4                 63.0 ± 0%      63.0 ± 0%     ~     (all equal)
    CompileOnepass/^a(?:bc)+$-4                   46.0 ± 0%      46.0 ± 0%     ~     (all equal)
    CompileOnepass/^a(?:[bcd])+$-4                43.0 ± 0%      43.0 ± 0%     ~     (all equal)
    CompileOnepass/^a((?:[bcd])+)$-4              49.0 ± 0%      49.0 ± 0%     ~     (all equal)
    CompileOnepass/^a(:?b|c)*d$-4                  105 ± 0%       101 ± 0%   -3.81%  (p=0.008 n=5+5)
    CompileOnepass/^.bc(d|e)*$-4                  62.0 ± 0%      60.0 ± 0%   -3.23%  (p=0.008 n=5+5)
    CompileOnepass/^loooooooooooooooooo...-4     1.09k ± 0%     1.09k ± 0%     ~     (all equal)
    
    Fixes #19735
    
    Change-Id: Ib90e18e1b06166407b26b2a68b88afbb1f486024
    Reviewed-on: https://go-review.googlesource.com/38751
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8a5175df35d20aa97afcbea63e86ba14ecafdc88
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Mar 22 21:34:12 2017 -0400

    cmd/compile: improve startRegs calculation
    
    In register allocation, we calculate what values are used in
    and after the current block. If a value is used only after a
    function call, since registers are clobbered in call, we don't
    need to mark the value live at the entrance of the block.
    Before this CL it is considered live, and unnecessary copy or
    load may be generated when resolving merge edge.
    
    Fixes #14761.
    
    On AMD64:
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.84s ± 1%     2.81s ± 1%   -1.06%  (p=0.000 n=10+9)
    Fannkuch11-12                3.61s ± 0%     3.55s ± 1%   -1.77%  (p=0.000 n=10+9)
    FmtFprintfEmpty-12          50.4ns ± 4%    50.0ns ± 1%     ~     (p=0.785 n=9+8)
    FmtFprintfString-12         80.0ns ± 3%    78.2ns ± 3%   -2.35%  (p=0.004 n=10+9)
    FmtFprintfInt-12            81.3ns ± 4%    81.8ns ± 2%     ~     (p=0.159 n=10+10)
    FmtFprintfIntInt-12          120ns ± 4%     118ns ± 2%     ~     (p=0.218 n=10+10)
    FmtFprintfPrefixedInt-12     152ns ± 3%     155ns ± 2%   +2.11%  (p=0.026 n=10+10)
    FmtFprintfFloat-12           240ns ± 1%     238ns ± 1%   -0.79%  (p=0.005 n=9+9)
    FmtManyArgs-12               504ns ± 1%     510ns ± 1%   +1.14%  (p=0.000 n=8+9)
    GobDecode-12                7.00ms ± 1%    6.99ms ± 0%     ~     (p=0.497 n=9+10)
    GobEncode-12                5.47ms ± 1%    5.48ms ± 1%     ~     (p=0.218 n=10+10)
    Gzip-12                      258ms ± 2%     256ms ± 1%   -0.96%  (p=0.043 n=10+9)
    Gunzip-12                   38.6ms ± 0%    38.3ms ± 0%   -0.64%  (p=0.000 n=9+8)
    HTTPClientServer-12         90.4µs ± 3%    87.2µs ±11%     ~     (p=0.053 n=9+10)
    JSONEncode-12               15.6ms ± 0%    15.6ms ± 1%     ~     (p=0.077 n=9+9)
    JSONDecode-12               55.1ms ± 1%    54.6ms ± 1%   -0.85%  (p=0.010 n=10+9)
    Mandelbrot200-12            4.49ms ± 0%    4.47ms ± 0%   -0.25%  (p=0.000 n=10+8)
    GoParse-12                  3.38ms ± 0%    3.37ms ± 1%     ~     (p=0.315 n=8+10)
    RegexpMatchEasy0_32-12      82.5ns ± 4%    82.0ns ± 0%     ~     (p=0.164 n=10+8)
    RegexpMatchEasy0_1K-12       203ns ± 1%     202ns ± 1%   -0.85%  (p=0.000 n=9+10)
    RegexpMatchEasy1_32-12      82.3ns ± 1%    81.1ns ± 0%   -1.39%  (p=0.000 n=10+8)
    RegexpMatchEasy1_1K-12       357ns ± 1%     357ns ± 1%     ~     (p=0.697 n=8+9)
    RegexpMatchMedium_32-12      125ns ± 2%     126ns ± 2%     ~     (p=0.197 n=10+10)
    RegexpMatchMedium_1K-12     39.6µs ± 3%    39.6µs ± 1%     ~     (p=0.971 n=10+10)
    RegexpMatchHard_32-12       1.99µs ± 2%    1.99µs ± 4%     ~     (p=0.891 n=10+9)
    RegexpMatchHard_1K-12       60.1µs ± 3%    60.4µs ± 3%     ~     (p=0.684 n=10+10)
    Revcomp-12                   531ms ± 6%     441ms ± 0%  -16.94%  (p=0.000 n=10+9)
    Template-12                 58.9ms ± 1%    58.7ms ± 1%     ~     (p=0.315 n=10+10)
    TimeParse-12                 319ns ± 1%     320ns ± 4%     ~     (p=0.215 n=9+9)
    TimeFormat-12                345ns ± 0%     333ns ± 1%   -3.36%  (p=0.000 n=9+10)
    [Geo mean]                  52.2µs         51.6µs        -1.13%
    
    On ARM64:
    name                     old time/op    new time/op    delta
    BinaryTree17-8              8.53s ± 0%     8.36s ± 0%   -1.89%  (p=0.000 n=10+10)
    Fannkuch11-8                6.15s ± 0%     6.10s ± 0%   -0.67%  (p=0.000 n=10+10)
    FmtFprintfEmpty-8           117ns ± 0%     117ns ± 0%     ~     (all equal)
    FmtFprintfString-8          192ns ± 0%     192ns ± 0%     ~     (all equal)
    FmtFprintfInt-8             198ns ± 0%     198ns ± 0%     ~     (p=0.211 n=10+10)
    FmtFprintfIntInt-8          289ns ± 0%     291ns ± 0%   +0.59%  (p=0.000 n=7+10)
    FmtFprintfPrefixedInt-8     320ns ± 2%     317ns ± 0%     ~     (p=0.431 n=10+8)
    FmtFprintfFloat-8           538ns ± 0%     538ns ± 0%     ~     (all equal)
    FmtManyArgs-8              1.17µs ± 1%    1.18µs ± 1%     ~     (p=0.063 n=10+10)
    GobDecode-8                17.0ms ± 1%    17.2ms ± 1%   +0.83%  (p=0.000 n=10+10)
    GobEncode-8                14.2ms ± 0%    14.1ms ± 1%   -0.78%  (p=0.001 n=9+10)
    Gzip-8                      806ms ± 0%     797ms ± 0%   -1.12%  (p=0.000 n=6+9)
    Gunzip-8                    131ms ± 0%     130ms ± 0%   -0.51%  (p=0.000 n=10+9)
    HTTPClientServer-8          206µs ± 9%     212µs ± 2%     ~     (p=0.829 n=10+8)
    JSONEncode-8               40.1ms ± 0%    40.1ms ± 0%     ~     (p=0.136 n=9+9)
    JSONDecode-8                157ms ± 0%     151ms ± 0%   -3.32%  (p=0.000 n=9+9)
    Mandelbrot200-8            10.1ms ± 0%    10.1ms ± 0%   -0.05%  (p=0.000 n=9+8)
    GoParse-8                  8.43ms ± 0%    8.43ms ± 0%     ~     (p=0.912 n=10+10)
    RegexpMatchEasy0_32-8       228ns ± 1%     227ns ± 0%   -0.26%  (p=0.026 n=10+9)
    RegexpMatchEasy0_1K-8      1.92µs ± 0%    1.63µs ± 0%  -15.18%  (p=0.001 n=7+7)
    RegexpMatchEasy1_32-8       258ns ± 1%     250ns ± 0%   -2.83%  (p=0.000 n=10+10)
    RegexpMatchEasy1_1K-8      2.39µs ± 0%    2.13µs ± 0%  -10.94%  (p=0.000 n=9+9)
    RegexpMatchMedium_32-8      352ns ± 0%     351ns ± 0%   -0.29%  (p=0.004 n=9+10)
    RegexpMatchMedium_1K-8      104µs ± 0%     105µs ± 0%   +0.58%  (p=0.000 n=8+9)
    RegexpMatchHard_32-8       5.84µs ± 0%    5.82µs ± 0%   -0.27%  (p=0.000 n=9+10)
    RegexpMatchHard_1K-8        177µs ± 0%     177µs ± 0%   -0.07%  (p=0.000 n=9+9)
    Revcomp-8                   1.57s ± 1%     1.50s ± 1%   -4.60%  (p=0.000 n=9+10)
    Template-8                  157ms ± 1%     153ms ± 1%   -2.28%  (p=0.000 n=10+9)
    TimeParse-8                 779ns ± 1%     770ns ± 1%   -1.18%  (p=0.013 n=10+10)
    TimeFormat-8                823ns ± 2%     826ns ± 1%     ~     (p=0.324 n=10+9)
    [Geo mean]                  144µs          142µs        -1.45%
    
    Reduce cmd/go text size by 0.5%.
    
    Change-Id: I9288ff983c4a7cf03fc0cb35b9b1750828013117
    Reviewed-on: https://go-review.googlesource.com/38457
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit fa1d54c2edad607866445577fe4949fbe55166e1
Author: Russ Cox <rsc@golang.org>
Date:   Tue Mar 28 14:54:10 2017 -0400

    cmd/go: exclude vendored packages from ... matches
    
    By overwhelming popular demand, exclude vendored packages from ... matches,
    by making ... never match the "vendor" element above a vendored package.
    
    go help packages now reads:
    
        An import path is a pattern if it includes one or more "..." wildcards,
        each of which can match any string, including the empty string and
        strings containing slashes.  Such a pattern expands to all package
        directories found in the GOPATH trees with names matching the
        patterns.
    
        To make common patterns more convenient, there are two special cases.
        First, /... at the end of the pattern can match an empty string,
        so that net/... matches both net and packages in its subdirectories, like net/http.
        Second, any slash-separted pattern element containing a wildcard never
        participates in a match of the "vendor" element in the path of a vendored
        package, so that ./... does not match packages in subdirectories of
        ./vendor or ./mycode/vendor, but ./vendor/... and ./mycode/vendor/... do.
        Note, however, that a directory named vendor that itself contains code
        is not a vendored package: cmd/vendor would be a command named vendor,
        and the pattern cmd/... matches it.
    
    Fixes #19090.
    
    Change-Id: I985bf9571100da316c19fbfd19bb1e534a3c9e5f
    Reviewed-on: https://go-review.googlesource.com/38745
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 68da265c8e6b32753fb5788716953bac16b374c0
Author: Keith Randall <khr@golang.org>
Date:   Wed Mar 29 18:06:04 2017 +0000

    Revert "cmd/compile: automatically handle commuting ops in rewrite rules"
    
    This reverts commit 041ecb697f0e867a2bb0bf219cc2fd5f77057c2e.
    
    Reason for revert: Not working on S390x and some 386 archs.
    I have a guess why the S390x is failing.  No clue on the 386 yet.
    Revert until I can figure it out.
    
    Change-Id: I64f1ce78fa6d1037ebe7ee2a8a8107cb4c1db70c
    Reviewed-on: https://go-review.googlesource.com/38790
    Reviewed-by: Keith Randall <khr@golang.org>

commit 8295dbda03ef2bfaccb6e2c139f1981d0c69964d
Author: Russ Cox <rsc@golang.org>
Date:   Tue Mar 28 14:59:13 2017 -0400

    cmd/go: make pattern matching tests less repetitive
    
    Change-Id: I25db1d637dd461cec67ba70659d523b46895c113
    Reviewed-on: https://go-review.googlesource.com/38744
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit f0a3018b00bb535259e7d46f0f1503048130f51c
Author: Russ Cox <rsc@golang.org>
Date:   Wed Mar 29 12:12:02 2017 -0400

    cmd/go: fix bug in test of go get ./path needing to download path
    
    rsc.io/toolstash is gone; use rsc.io/pprof_mac_fix.
    
    This fixes a bug in the test. It turns out the code being tested here
    is also broken, so the test still doesn't pass after this CL (filed #19769).
    
    Change-Id: Ieb725c321d7fab600708e133ae28f531e55521ad
    Reviewed-on: https://go-review.googlesource.com/38743
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 24e94766c082338d83e5f3e1f8eccf4f070d8956
Author: David Chase <drchase@google.com>
Date:   Tue Mar 28 17:55:26 2017 -0400

    cmd/compile: added special case for reflect header fields to esc
    
    The uintptr-typed Data field in reflect.SliceHeader and
    reflect.StringHeader needs special treatment because it is
    really a pointer.  Add the special treatment in walk for
    bug #19168 to escape analysis.
    
    Includes extra debugging that was helpful.
    
    Fixes #19743.
    
    Change-Id: I6dab5002f0d436c3b2a7cdc0156e4fc48a43d6fe
    Reviewed-on: https://go-review.googlesource.com/38738
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 83843b16100d889bc84580c9427e02c7c9cee769
Author: David Lazar <lazard@golang.org>
Date:   Wed Mar 8 17:08:31 2017 -0500

    cmd/compile: fix names of inlined methods from other packages
    
    Previously, an inlined call to wg.Done() in package main would have the
    following incorrect symbol name:
    
        main.(*sync.WaitGroup).Done
    
    This change modifies methodname to return the correct symbol name:
    
        sync.(*WaitGroup).Done
    
    This fix was suggested by @mdempsky.
    
    Fixes #19467.
    
    Change-Id: I0117838679ac5353789299c618ff8c326712d94d
    Reviewed-on: https://go-review.googlesource.com/37866
    Reviewed-by: Austin Clements <austin@google.com>

commit 7bf0adc6add500054ccbc37a868ab0fab120fa24
Author: David Lazar <lazard@golang.org>
Date:   Tue Mar 7 21:14:12 2017 -0500

    runtime: include inlined calls in result of CallersFrames
    
    Change-Id: If1a3396175f2afa607d56efd1444181334a9ae3e
    Reviewed-on: https://go-review.googlesource.com/37862
    Reviewed-by: Austin Clements <austin@google.com>

commit ee97216a1787a979911d43c0c5c582b5492a2205
Author: David Lazar <lazard@golang.org>
Date:   Mon Mar 6 14:48:36 2017 -0500

    runtime: handle inlined calls in runtime.Callers
    
    The `skip` argument passed to runtime.Caller and runtime.Callers should
    be interpreted as the number of logical calls to skip (rather than the
    number of physical stack frames to skip). This changes runtime.Callers
    to skip inlined calls in addition to physical stack frames.
    
    The result value of runtime.Callers is a slice of program counters
    ([]uintptr) representing physical stack frames. If the `skip` parameter
    to runtime.Callers skips part-way into a physical frame, there is no
    convenient way to encode that in the resulting slice. To avoid changing
    the API in an incompatible way, our solution is to store the number of
    skipped logical calls of the first frame in the _second_ uintptr
    returned by runtime.Callers. Since this number is a small integer, we
    encode it as a valid PC value into a small symbol called:
    
        runtime.skipPleaseUseCallersFrames
    
    For example, if f() calls g(), g() calls `runtime.Callers(2, pcs)`, and
    g() is inlined into f, then the frame for f will be partially skipped,
    resulting in the following slice:
    
        pcs = []uintptr{pc_in_f, runtime.skipPleaseUseCallersFrames+1, ...}
    
    We store the skip PC in pcs[1] instead of pcs[0] so that `pcs[i:]` will
    truncate the captured stack trace rather than grow it for all i.
    
    Updates #19348.
    
    Change-Id: I1c56f89ac48c29e6f52a5d085567c6d77d499cf1
    Reviewed-on: https://go-review.googlesource.com/37854
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit f3f5b10e06f6bb29731e9213dd8745a9b1857568
Author: David Lazar <lazard@golang.org>
Date:   Tue Mar 14 00:47:03 2017 -0400

    test: allow flags in run action
    
    Previously, we could not run tests with -l=4 on NaCl since the buildrun
    action is not supported on NaCl. This lets us run tests with build flags
    on NaCl.
    
    Change-Id: I103370c7b823b4ff46f47df97e802da0dc2bc7c3
    Reviewed-on: https://go-review.googlesource.com/38170
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 94c95d3e52ed55ef57e60ace559b684e4d5ebcbf
Author: Russ Cox <rsc@golang.org>
Date:   Wed Mar 29 10:16:50 2017 -0400

    cmd/go: build test binaries with -s in addition to -w
    
    Fixes #19753.
    
    Change-Id: Ib20a69b1d0bcc42aa9e924918bcb578d6a560a31
    Reviewed-on: https://go-review.googlesource.com/38742
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6983b9a57955fa12ecd81ab8394ee09e64ef21b9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Mar 29 16:55:58 2017 +0000

    net, net/http: adjust time-in-past constant even earlier
    
    The aLongTimeAgo time value in net and net/http is used to cancel
    in-flight read and writes. It was set to time.Unix(233431200, 0)
    which seemed like far enough in the past.
    
    But Raspberry Pis, lacking a real time clock, had to spoil the fun and
    boot in 1970 at the Unix epoch time, breaking assumptions in net and
    net/http.
    
    So change aLongTimeAgo to time.Unix(1, 0), which seems like the
    earliest safe value. I don't trust subsecond values on all operating
    systems, and I don't trust the Unix zero time. The Raspberry Pis do
    advance their clock at least. And the reported problem was that Hijack
    on a ResponseWriter hung forever, waiting for the connection read
    operation to finish. So now, even if kernel + userspace boots in under
    a second (unlikely), the Hijack will just have to wait for up to a
    second.
    
    Fixes #19747
    
    Change-Id: Id59430de2e7b5b5117d4903a788863e9d344e53a
    Reviewed-on: https://go-review.googlesource.com/38785
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 041ecb697f0e867a2bb0bf219cc2fd5f77057c2e
Author: Keith Randall <khr@golang.org>
Date:   Sat Mar 25 15:05:42 2017 -0700

    cmd/compile: automatically handle commuting ops in rewrite rules
    
    We have lots of rewrite rules that vary only in the fact that
    we have 2 versions for the 2 different orderings of various
    commuting ops. For example:
    
    (ADDL x (MOVLconst [c])) -> (ADDLconst [c] x)
    (ADDL (MOVLconst [c]) x) -> (ADDLconst [c] x)
    
    It can get unwieldly quickly, especially when there is more than
    one commuting op in a rule.
    
    Our existing "fix" for this problem is to have rules that
    canonicalize the operations first. For example:
    
    (Eq64 x (Const64 <t> [c])) && x.Op != OpConst64 -> (Eq64 (Const64 <t> [c]) x)
    
    Subsequent rules can then assume if there is a constant arg to Eq64,
    it will be the first one. This fix kinda works, but it is fragile and
    only works when we remember to include the required extra rules.
    
    The fundamental problem is that the rule matcher doesn't
    know anything about commuting ops. This CL fixes that fact.
    
    We already have information about which ops commute. (The register
    allocator takes advantage of commutivity.)  The rule generator now
    automatically generates multiple rules for a single source rule when
    there are commutative ops in the rule. We can now drop all of our
    almost-duplicate source-level rules and the canonicalization rules.
    
    I have some CLs in progress that will be a lot less verbose when
    the rule generator handles commutivity for me.
    
    I had to reorganize the load-combining rules a bit. The 8-way OR rules
    generated 128 different reorderings, which was causing the generator
    to put too much code in the rewrite*.go files (the big ones were going
    from 25K lines to 132K lines). Instead I reorganized the rules to
    combine pairs of loads at a time. The generated rule files are now
    actually a bit (5%) smaller.
    [Note to reviewers: check these carefully. Most of the other rule
    changes are trivial.]
    
    Make.bash times are ~unchanged.
    
    Compiler benchmarks are not observably different. Probably because
    we don't spend much compiler time in rule matching anyway.
    
    I've also done a pass over all of our ops adding commutative markings
    for ops which hadn't had them previously.
    
    Fixes #18292
    
    Change-Id: I999b1307272e91965b66754576019dedcbe7527a
    Reviewed-on: https://go-review.googlesource.com/38666
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 627798db4e42e4fd3f80a4149cbfafff80b304f3
Author: Austin Clements <austin@google.com>
Date:   Tue Mar 28 16:03:24 2017 -0400

    reflect: fix out-of-bounds pointers calling no-result method
    
    reflect.callReflect heap-allocates a stack frame and then constructs
    pointers to the arguments and result areas of that frame. However, if
    there are no results, the results pointer will point past the end of
    the frame allocation. If there are also no arguments, the arguments
    pointer will also point past the end of the frame allocation. If the
    GC observes either these pointers, it may panic.
    
    Fix this by not constructing these pointers if these areas of the
    frame are empty.
    
    This adds a test of calling no-argument/no-result methods via reflect,
    since nothing in std did this before. However, it's quite difficult to
    demonstrate the actual failure because it depends on both exact
    allocation patterns and on GC scanning the goroutine's stack while
    inside one of the typedmemmovepartial calls.
    
    I also audited other uses of typedmemmovepartial and
    memclrNoHeapPointers in reflect, since these are the most susceptible
    to this. These appear to be the only two cases that can construct
    out-of-bounds arguments to these functions.
    
    Fixes #19724.
    
    Change-Id: I4b83c596b5625dc4ad0567b1e281bad4faef972b
    Reviewed-on: https://go-review.googlesource.com/38736
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6e9ec14186cad6058625415abba2744e2bd83ec7
Author: Rick Hudson <rlh@golang.org>
Date:   Mon Mar 27 14:20:35 2017 -0400

    runtime: redo insert/remove of large spans
    
    Currently for spans with up to 1 MBytes (128 pages) we
    maintain an array indexed by the number of pages in the
    span. This is efficient both in terms of space as well
    as time to insert or remove a span of a particular size.
    
    Unfortunately for spans larger than 1 MByte we currently
    place them on a separate linked list. This results in
    O(n) behavior. Now that we are seeing heaps approaching
    100 GBytes n is large enough to be noticed in real programs.
    
    This change replaces the linked list now used with a balanced
    binary tree structure called a treap. A treap is a
    probabilistically balanced tree offering O(logN) behavior for
    inserting and removing spans.
    
    To verify that this approach will work we start with noting
    that only spans with sizes > 1MByte will be put into the treap.
    This means that to support 1 TByte a treap will need at most
    1 million nodes and can ideally be held in a treap with a
    depth of 20. Experiments with adding and removing randomly
    sized spans from the treap seem to result in treaps with
    depths of about twice the ideal or 40. A petabyte would
    require a tree of only twice again that depth again so this
    algorithm should last well into the future.
    
    Fixes #19393
    
    Go1 benchmarks indicate this is basically an overall wash.
    Tue Mar 28 21:29:21 EDT 2017
    name                     old time/op    new time/op    delta
    BinaryTree17-4              2.42s ± 1%     2.42s ± 1%    ~     (p=0.980 n=21+21)
    Fannkuch11-4                3.00s ± 1%     3.18s ± 4%  +6.10%  (p=0.000 n=22+24)
    FmtFprintfEmpty-4          40.5ns ± 1%    40.3ns ± 3%    ~     (p=0.692 n=22+25)
    FmtFprintfString-4         65.9ns ± 3%    64.6ns ± 1%  -1.98%  (p=0.000 n=24+23)
    FmtFprintfInt-4            69.6ns ± 1%    68.0ns ± 7%  -2.30%  (p=0.001 n=21+22)
    FmtFprintfIntInt-4          102ns ± 2%      99ns ± 1%  -3.07%  (p=0.000 n=23+23)
    FmtFprintfPrefixedInt-4     126ns ± 0%     125ns ± 0%  -0.79%  (p=0.000 n=19+17)
    FmtFprintfFloat-4           206ns ± 2%     205ns ± 1%    ~     (p=0.671 n=23+21)
    FmtManyArgs-4               441ns ± 1%     445ns ± 1%  +0.88%  (p=0.000 n=22+23)
    GobDecode-4                5.73ms ± 1%    5.86ms ± 1%  +2.37%  (p=0.000 n=23+22)
    GobEncode-4                4.51ms ± 1%    4.89ms ± 1%  +8.32%  (p=0.000 n=22+22)
    Gzip-4                      197ms ± 0%     202ms ± 1%  +2.75%  (p=0.000 n=23+24)
    Gunzip-4                   32.9ms ± 8%    32.7ms ± 2%    ~     (p=0.466 n=23+24)
    HTTPClientServer-4         57.3µs ± 1%    56.7µs ± 1%  -0.94%  (p=0.000 n=21+22)
    JSONEncode-4               13.8ms ± 1%    13.9ms ± 2%  +1.14%  (p=0.000 n=22+23)
    JSONDecode-4               47.4ms ± 1%    48.1ms ± 1%  +1.49%  (p=0.000 n=23+23)
    Mandelbrot200-4            3.92ms ± 0%    3.92ms ± 1%  +0.21%  (p=0.000 n=22+22)
    GoParse-4                  2.89ms ± 1%    2.87ms ± 1%  -0.68%  (p=0.000 n=21+22)
    RegexpMatchEasy0_32-4      73.6ns ± 1%    72.0ns ± 2%  -2.15%  (p=0.000 n=21+22)
    RegexpMatchEasy0_1K-4       173ns ± 1%     173ns ± 1%    ~     (p=0.847 n=22+24)
    RegexpMatchEasy1_32-4      71.9ns ± 1%    69.8ns ± 1%  -2.99%  (p=0.000 n=23+20)
    RegexpMatchEasy1_1K-4       314ns ± 1%     308ns ± 1%  -1.91%  (p=0.000 n=22+23)
    RegexpMatchMedium_32-4      106ns ± 0%     105ns ± 1%  -0.58%  (p=0.000 n=19+21)
    RegexpMatchMedium_1K-4     34.3µs ± 1%    34.3µs ± 1%    ~     (p=0.871 n=23+22)
    RegexpMatchHard_32-4       1.67µs ± 1%    1.67µs ± 7%    ~     (p=0.224 n=22+23)
    RegexpMatchHard_1K-4       51.5µs ± 1%    50.4µs ± 1%  -1.99%  (p=0.000 n=22+23)
    Revcomp-4                   383ms ± 1%     415ms ± 0%  +8.51%  (p=0.000 n=22+22)
    Template-4                 51.5ms ± 1%    51.5ms ± 1%    ~     (p=0.555 n=20+23)
    TimeParse-4                 279ns ± 2%     277ns ± 1%  -0.95%  (p=0.000 n=24+22)
    TimeFormat-4                294ns ± 1%     296ns ± 1%  +0.58%  (p=0.003 n=24+23)
    [Geo mean]                 43.7µs         43.8µs       +0.32%
    
    name                     old speed      new speed      delta
    GobDecode-4               134MB/s ± 1%   131MB/s ± 1%  -2.32%  (p=0.000 n=23+22)
    GobEncode-4               170MB/s ± 1%   157MB/s ± 1%  -7.68%  (p=0.000 n=22+22)
    Gzip-4                   98.7MB/s ± 0%  96.1MB/s ± 1%  -2.68%  (p=0.000 n=23+24)
    Gunzip-4                  590MB/s ± 7%   593MB/s ± 2%    ~     (p=0.466 n=23+24)
    JSONEncode-4              141MB/s ± 1%   139MB/s ± 2%  -1.13%  (p=0.000 n=22+23)
    JSONDecode-4             40.9MB/s ± 1%  40.3MB/s ± 0%  -1.47%  (p=0.000 n=23+23)
    GoParse-4                20.1MB/s ± 1%  20.2MB/s ± 1%  +0.69%  (p=0.000 n=21+22)
    RegexpMatchEasy0_32-4     435MB/s ± 1%   444MB/s ± 2%  +2.21%  (p=0.000 n=21+22)
    RegexpMatchEasy0_1K-4    5.89GB/s ± 1%  5.89GB/s ± 1%    ~     (p=0.439 n=22+24)
    RegexpMatchEasy1_32-4     445MB/s ± 1%   459MB/s ± 1%  +3.06%  (p=0.000 n=23+20)
    RegexpMatchEasy1_1K-4    3.26GB/s ± 1%  3.32GB/s ± 1%  +1.97%  (p=0.000 n=22+23)
    RegexpMatchMedium_32-4   9.40MB/s ± 1%  9.44MB/s ± 1%  +0.43%  (p=0.000 n=23+21)
    RegexpMatchMedium_1K-4   29.8MB/s ± 1%  29.8MB/s ± 1%    ~     (p=0.826 n=23+22)
    RegexpMatchHard_32-4     19.1MB/s ± 1%  19.1MB/s ± 7%    ~     (p=0.233 n=22+23)
    RegexpMatchHard_1K-4     19.9MB/s ± 1%  20.3MB/s ± 1%  +2.03%  (p=0.000 n=22+23)
    Revcomp-4                 664MB/s ± 1%   612MB/s ± 0%  -7.85%  (p=0.000 n=22+22)
    Template-4               37.6MB/s ± 1%  37.7MB/s ± 1%    ~     (p=0.558 n=20+23)
    [Geo mean]                134MB/s        133MB/s       -0.76%
    Tue Mar 28 22:16:54 EDT 2017
    
    Change-Id: I4a4f5c2b53d3fb85ef76c98522d3ed5cf8ae5b7e
    Reviewed-on: https://go-review.googlesource.com/38732
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 846f925464e20fe9c09b289aefa4db0cdcd886b0
Author: haya14busa <haya14busa@gmail.com>
Date:   Wed Mar 29 13:51:00 2017 +0900

    cmd/go: add -json flag to go env
    
    "go env" prints Go environment information as a shell script format by
    default but it's difficult for some tools (e.g. editor packages) to
    interpret it.
    
    The -json flag prints the environment in JSON format which
    can be easily interpreted by a lot of tools.
    
    $ go env -json
    {
            "CC": "gcc",
            "CGO_CFLAGS": "-g -O2",
            "CGO_CPPFLAGS": "",
            "CGO_CXXFLAGS": "-g -O2",
            "CGO_ENABLED": "1",
            "CGO_FFLAGS": "-g -O2",
            "CGO_LDFLAGS": "-g -O2",
            "CXX": "g++",
            "GCCGO": "gccgo",
            "GOARCH": "amd64",
            "GOBIN": "/home/haya14busa/go/bin",
            "GOEXE": "",
            "GOGCCFLAGS": "-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build498013955=/tmp/go-build -gno-record-gcc-switches",
            "GOHOSTARCH": "amd64",
            "GOHOSTOS": "linux",
            "GOOS": "linux",
            "GOPATH": "/home/haya14busa",
            "GORACE": "",
            "GOROOT": "/home/haya14busa/src/go.googlesource.com/go",
            "GOTOOLDIR": "/home/haya14busa/src/go.googlesource.com/go/pkg/tool/linux_amd64",
            "PKG_CONFIG": "pkg-config"
    }
    
    Also, it supports arguments with -json flag.
    
    $ go env -json GOROOT GOPATH GOBIN
    {
            "GOBIN": "/home/haya14busa/go/bin",
            "GOPATH": "/home/haya14busa",
            "GOROOT": "/home/haya14busa/src/go.googlesource.com/go"
    }
    
    Fixes #12567
    
    Change-Id: I75db3780f14a8ab8c7fa58cc3c9cc488ef7b66a1
    Reviewed-on: https://go-review.googlesource.com/38757
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9f232c1786f57051d34848db48ece398ad0329db
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Mar 29 10:55:46 2017 +1100

    cmd/compile/internal/gc: remove unused state.placeholder field
    
    gc.state.placeholder was added in 5a6e511c61 but never used.
    
    Change-Id: I5a621507279d5bb1f3991b7a412d9a63039c464e
    Reviewed-on: https://go-review.googlesource.com/38755
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 54af18708b4459beec6b54fcd9bb51848d3fbe6c
Author: Dave Cheney <dave@cheney.net>
Date:   Mon Mar 27 15:27:44 2017 +1100

    cmd/internal/obj/x86: clean up byteswapreg
    
    Make byteswapreg more Go like.
    
    Change-Id: Ibdf3603cae9cad2b3465b4c224a28a4c4c745c2e
    Reviewed-on: https://go-review.googlesource.com/38615
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b72636cbde8e241ae216748ad915a9f9cf620988
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 28 15:16:30 2017 -0700

    cmd/compile/internal/gc: cleanup selecttype
    
    Use namedfield consistently.
    
    Passes toolstash-check.
    
    Change-Id: Ic5a3acb4bfaa1f60dd2eac94612160509e8d7f94
    Reviewed-on: https://go-review.googlesource.com/38741
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3c0072e0567c55d7172f137389cbfdfe0c428888
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 28 15:13:19 2017 -0700

    cmd/compile/internal/gc: use anonfield and namedfield
    
    Automated refactoring using gofmt.
    
    Passes toolstash-check.
    
    Change-Id: I8624e1c231dc736e1bb4cc800acaf629a0af91d7
    Reviewed-on: https://go-review.googlesource.com/38740
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3f4f119b5e5a48c558c2b318b32f036e1062b108
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 28 15:01:18 2017 -0700

    cmd/compile/internal/gc: npos(p, nod(o, l, r)) -> nodl(p, o, l, r)
    
    Prepared with gofmt -r.
    
    Passes toolstash-check.
    
    Change-Id: I8a4f7a8c365dfe464dfc5868a18c67d56af37749
    Reviewed-on: https://go-review.googlesource.com/38739
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 835b17c85f62c0e952646ace959510aca6c28568
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Mar 28 12:41:40 2017 -0700

    test: add test for gccgo compiler crash
    
    Gccgo crashed compiling a function that returned multiple zero-sized values.
    
    Change-Id: I499112cc310e4a4f649962f4d2bc9fee95dee1b6
    Reviewed-on: https://go-review.googlesource.com/38772
    Reviewed-by: Than McIntosh <thanm@google.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4b50c81356b7e36ae1de1651eb3f5fb9df3211dd
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Tue Feb 21 16:57:37 2017 -0600

    test/fixedbugs: add a test for 19201
    
    This was cherry-picked to 1.8 as CL 38587, but on master issue was fixed
    by CL 37661. Add still relevant part (test) and close issue, since test passes.
    
    Fixes #19201
    
    Change-Id: I6415792e2c465dc6d9bd6583ba1e54b107bcf5cc
    Reviewed-on: https://go-review.googlesource.com/37376
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 95bfd927f5d05fd2fb82dba918ff8d36bbc6d400
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 28 11:17:52 2017 -0700

    cmd/compile: fix two instances of { lineno = ...; yyerror }
    
    Updates #19683
    
    Change-Id: Ic00d5a9807200791cf37553f4f802dbf27beea19
    Reviewed-on: https://go-review.googlesource.com/38770
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b87fcc6e0641b351fb5a0daa5537baf4d0c7316e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 28 10:36:18 2017 -0700

    cmd/compile: number autotmps per-func, not per-package
    
    Prior to this CL, autotmps were global to a package.
    They also shared numbering with static variables.
    Switch autotmp numbering to be per-function instead,
    and do implicit numbering based on len(Func.Dcl).
    This eliminates a dependency on a global variable
    from the backend without adding to the Func struct.
    While we're here, move statuniqgen closer to its
    sole remaining user.
    
    This actually improves compiler performance,
    because the autotmp_* names can now be
    reused across functions.
    
    name       old alloc/op    new alloc/op    delta
    Template      40.6MB ± 0%     40.1MB ± 0%  -1.38%  (p=0.000 n=10+10)
    Unicode       29.9MB ± 0%     29.9MB ± 0%    ~     (p=0.912 n=10+10)
    GoTypes        116MB ± 0%      114MB ± 0%  -1.53%  (p=0.000 n=10+10)
    SSA            865MB ± 0%      856MB ± 0%  -1.04%  (p=0.000 n=10+10)
    Flate         25.8MB ± 0%     25.4MB ± 0%  -1.36%  (p=0.000 n=10+10)
    GoParser      32.2MB ± 0%     32.0MB ± 0%  -0.72%  (p=0.000 n=10+10)
    Reflect       80.3MB ± 0%     79.0MB ± 0%  -1.65%  (p=0.000 n=9+10)
    Tar           27.0MB ± 0%     26.7MB ± 0%  -0.86%  (p=0.000 n=10+9)
    XML           42.8MB ± 0%     42.4MB ± 0%  -0.95%  (p=0.000 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        398k ± 1%       396k ± 1%  -0.59%  (p=0.002 n=10+10)
    Unicode         321k ± 1%       321k ± 0%    ~     (p=0.912 n=10+10)
    GoTypes        1.17M ± 0%      1.16M ± 0%  -0.77%  (p=0.000 n=10+10)
    SSA            7.65M ± 0%      7.62M ± 0%  -0.40%  (p=0.000 n=10+10)
    Flate           240k ± 1%       238k ± 1%  -0.56%  (p=0.001 n=10+10)
    GoParser        323k ± 1%       320k ± 1%  -0.65%  (p=0.002 n=10+10)
    Reflect        1.01M ± 0%      1.00M ± 0%  -0.37%  (p=0.001 n=9+10)
    Tar             256k ± 1%       255k ± 0%    ~     (p=0.101 n=10+8)
    XML             400k ± 1%       398k ± 1%    ~     (p=0.063 n=10+10)
    
    
    Change-Id: I3c23ca98129137d373106990b1a3e1507bbe0cc3
    Reviewed-on: https://go-review.googlesource.com/38729
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8ee2d5bc00c5cc2313d87e001961294907cee133
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 23 11:40:47 2017 -0700

    cmd/compile: strengthen SetFields/Width safety guarantee
    
    It is currently possible in the compiler to create a struct type,
    calculate the widths of types that depend on it,
    and then alter the struct type.
    
    transformclosure has local protection against this.
    Protect against it at a deeper level.
    
    This is preparation to call dowidth automatically,
    rather than explicitly.
    
    This is a re-roll of CL 38469.
    
    Change-Id: Ic5b4baa250618504611fc57cbf51ab01d1eddf80
    Reviewed-on: https://go-review.googlesource.com/38534
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8dafdb1be147f0cd5a2811c69ce6e27769f5e5f1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 23 11:29:37 2017 -0700

    cmd/compile: add Type.WidthCalculated
    
    Prior to this CL, Type.Width != 0 was the mark
    of a Type whose Width had been calculated.
    As a result, dowidth always recalculated
    the width of struct{}.
    This, combined with the prohibition on calculating
    the width of a FuncArgsStruct and the use of
    struct{} as a function argument,
    meant that there were circumstances in which
    it was forbidden to call dowidth on a type.
    This inhibits refactoring to call dowidth automatically,
    rather than explicitly.
    Instead add a helper method, Type.WidthCalculated,
    and implement as Type.Align > 0.
    Type.Width is not a good candidate for tracking
    whether the width has been calculated;
    0 is a value type width, and Width is subject to
    too much magic value game-playing.
    
    For good measure, add a test for #11354.
    
    Change-Id: Ie9a9fb5d924e7a2010c1904ae5e38ed4a38eaeb2
    Reviewed-on: https://go-review.googlesource.com/38468
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f5c1926e93d4f383ee36d4c0351f2348440cbbe8
Author: haya14busa <haya14busa@gmail.com>
Date:   Tue Mar 28 12:35:06 2017 +0900

    regexp: reduce allocations at onePassCopy
    
    It reduces needless allocations on compiling onepass regex.
    
    name                                      old time/op    new time/op    delta
    CompileOnepass/^(?:(?:(?:.(?:$))?))...-4    6.31µs ± 3%    6.11µs ± 3%     ~     (p=0.056 n=5+5)
    CompileOnepass/^abcd$-4                     5.69µs ±12%    4.93µs ± 4%  -13.42%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a{0,})*?)$-4          7.10µs ±12%    5.82µs ± 5%  -17.95%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a+)*)$-4              5.99µs ±10%    6.07µs ±11%     ~     (p=1.000 n=5+5)
    CompileOnepass/^(?:(?:a|(?:aa)))$-4         7.36µs ± 4%    7.81µs ±19%     ~     (p=0.310 n=5+5)
    CompileOnepass/^(?:[^\s\S])$-4              4.71µs ± 3%    4.71µs ± 5%     ~     (p=1.000 n=5+5)
    CompileOnepass/^(?:(?:(?:a*)+))$-4          6.06µs ± 2%    6.23µs ± 9%     ~     (p=0.310 n=5+5)
    CompileOnepass/^[a-c]+$-4                   4.74µs ± 4%    4.64µs ± 6%     ~     (p=0.421 n=5+5)
    CompileOnepass/^[a-c]*$-4                   5.17µs ± 2%    4.68µs ± 0%   -9.57%  (p=0.016 n=5+4)
    CompileOnepass/^(?:a*)$-4                   5.34µs ± 3%    5.08µs ±12%     ~     (p=0.151 n=5+5)
    CompileOnepass/^(?:(?:aa)|a)$-4             7.24µs ± 5%    7.33µs ±12%     ~     (p=0.841 n=5+5)
    CompileOnepass/^...$-4                      5.28µs ± 3%    4.99µs ± 9%     ~     (p=0.095 n=5+5)
    CompileOnepass/^(?:a|(?:aa))$-4             7.20µs ± 4%    7.24µs ±10%     ~     (p=0.841 n=5+5)
    CompileOnepass/^a((b))c$-4                  7.99µs ± 3%    7.76µs ± 8%     ~     (p=0.151 n=5+5)
    CompileOnepass/^a.[l-nA-Cg-j]?e$-4          8.30µs ± 5%    7.29µs ± 4%  -12.08%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))$-4                   7.34µs ± 4%    7.24µs ±19%     ~     (p=0.690 n=5+5)
    CompileOnepass/^a(?:(b)|(c))c$-4            9.80µs ± 6%    9.49µs ±18%     ~     (p=0.151 n=5+5)
    CompileOnepass/^a(?:b|c)$-4                 5.23µs ± 3%    4.80µs ±10%     ~     (p=0.056 n=5+5)
    CompileOnepass/^a(?:b?|c)$-4                8.26µs ± 3%    7.30µs ± 3%  -11.62%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c+)$-4               9.18µs ± 2%    8.16µs ± 2%  -11.06%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:bc)+$-4                 6.16µs ± 3%    6.41µs ±13%     ~     (p=0.548 n=5+5)
    CompileOnepass/^a(?:[bcd])+$-4              5.75µs ± 5%    5.50µs ±12%     ~     (p=0.151 n=5+5)
    CompileOnepass/^a((?:[bcd])+)$-4            7.65µs ± 5%    6.93µs ± 9%     ~     (p=0.056 n=5+5)
    CompileOnepass/^a(:?b|c)*d$-4               13.0µs ± 1%    12.1µs ± 2%   -6.91%  (p=0.008 n=5+5)
    CompileOnepass/^.bc(d|e)*$-4                9.20µs ± 4%    8.25µs ± 3%  -10.38%  (p=0.008 n=5+5)
    CompileOnepass/^loooooooooooooooooo...-4     254µs ± 2%     220µs ± 6%  -13.47%  (p=0.008 n=5+5)
    
    name                                      old alloc/op   new alloc/op   delta
    CompileOnepass/^(?:(?:(?:.(?:$))?))...-4    3.92kB ± 0%    3.41kB ± 0%  -13.06%  (p=0.008 n=5+5)
    CompileOnepass/^abcd$-4                     3.20kB ± 0%    2.75kB ± 0%  -14.00%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a{0,})*?)$-4          3.85kB ± 0%    3.34kB ± 0%  -13.31%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a+)*)$-4              3.46kB ± 0%    2.95kB ± 0%  -14.78%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a|(?:aa)))$-4         4.20kB ± 0%    3.75kB ± 0%  -10.67%  (p=0.008 n=5+5)
    CompileOnepass/^(?:[^\s\S])$-4              3.10kB ± 0%    2.46kB ± 0%  -20.62%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:(?:a*)+))$-4          3.64kB ± 0%    3.13kB ± 0%  -14.07%  (p=0.008 n=5+5)
    CompileOnepass/^[a-c]+$-4                   3.06kB ± 0%    2.48kB ± 0%  -18.85%  (p=0.008 n=5+5)
    CompileOnepass/^[a-c]*$-4                   3.10kB ± 0%    2.52kB ± 0%  -18.60%  (p=0.008 n=5+5)
    CompileOnepass/^(?:a*)$-4                   3.21kB ± 0%    2.63kB ± 0%  -17.96%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:aa)|a)$-4             4.09kB ± 0%    3.64kB ± 0%  -10.96%  (p=0.008 n=5+5)
    CompileOnepass/^...$-4                      3.42kB ± 0%    2.91kB ± 0%  -14.95%  (p=0.008 n=5+5)
    CompileOnepass/^(?:a|(?:aa))$-4             4.09kB ± 0%    3.64kB ± 0%  -10.96%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))c$-4                  5.67kB ± 0%    4.39kB ± 0%  -22.59%  (p=0.008 n=5+5)
    CompileOnepass/^a.[l-nA-Cg-j]?e$-4          5.73kB ± 0%    4.32kB ± 0%  -24.58%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))$-4                   5.41kB ± 0%    4.06kB ± 0%  -24.85%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:(b)|(c))c$-4            6.40kB ± 0%    5.31kB ± 0%  -17.00%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b|c)$-4                 3.46kB ± 0%    2.88kB ± 0%  -16.67%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c)$-4                5.77kB ± 0%    4.36kB ± 0%  -24.41%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c+)$-4               5.94kB ± 0%    4.59kB ± 0%  -22.64%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:bc)+$-4                 3.60kB ± 0%    3.15kB ± 0%  -12.44%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:[bcd])+$-4              3.46kB ± 0%    2.94kB ± 0%  -14.81%  (p=0.008 n=5+5)
    CompileOnepass/^a((?:[bcd])+)$-4            5.50kB ± 0%    4.09kB ± 0%  -25.62%  (p=0.008 n=5+5)
    CompileOnepass/^a(:?b|c)*d$-4               7.24kB ± 0%    6.15kB ± 0%  -15.03%  (p=0.008 n=5+5)
    CompileOnepass/^.bc(d|e)*$-4                5.75kB ± 0%    4.47kB ± 0%  -22.25%  (p=0.008 n=5+5)
    CompileOnepass/^loooooooooooooooooo...-4     225kB ± 0%     135kB ± 0%  -39.99%  (p=0.008 n=5+5)
    
    name                                      old allocs/op  new allocs/op  delta
    CompileOnepass/^(?:(?:(?:.(?:$))?))...-4      52.0 ± 0%      49.0 ± 0%   -5.77%  (p=0.008 n=5+5)
    CompileOnepass/^abcd$-4                       44.0 ± 0%      41.0 ± 0%   -6.82%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a{0,})*?)$-4            52.0 ± 0%      49.0 ± 0%   -5.77%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a+)*)$-4                47.0 ± 0%      44.0 ± 0%   -6.38%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:a|(?:aa)))$-4           57.0 ± 0%      54.0 ± 0%   -5.26%  (p=0.008 n=5+5)
    CompileOnepass/^(?:[^\s\S])$-4                36.0 ± 0%      33.0 ± 0%   -8.33%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:(?:a*)+))$-4            49.0 ± 0%      46.0 ± 0%   -6.12%  (p=0.008 n=5+5)
    CompileOnepass/^[a-c]+$-4                     39.0 ± 0%      36.0 ± 0%   -7.69%  (p=0.008 n=5+5)
    CompileOnepass/^[a-c]*$-4                     44.0 ± 0%      41.0 ± 0%   -6.82%  (p=0.008 n=5+5)
    CompileOnepass/^(?:a*)$-4                     45.0 ± 0%      42.0 ± 0%   -6.67%  (p=0.008 n=5+5)
    CompileOnepass/^(?:(?:aa)|a)$-4               56.0 ± 0%      53.0 ± 0%   -5.36%  (p=0.008 n=5+5)
    CompileOnepass/^...$-4                        46.0 ± 0%      43.0 ± 0%   -6.52%  (p=0.008 n=5+5)
    CompileOnepass/^(?:a|(?:aa))$-4               56.0 ± 0%      53.0 ± 0%   -5.36%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))c$-4                    57.0 ± 0%      53.0 ± 0%   -7.02%  (p=0.008 n=5+5)
    CompileOnepass/^a.[l-nA-Cg-j]?e$-4            62.0 ± 0%      58.0 ± 0%   -6.45%  (p=0.008 n=5+5)
    CompileOnepass/^a((b))$-4                     51.0 ± 0%      47.0 ± 0%   -7.84%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:(b)|(c))c$-4              69.0 ± 0%      65.0 ± 0%   -5.80%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b|c)$-4                   43.0 ± 0%      40.0 ± 0%   -6.98%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c)$-4                  61.0 ± 0%      57.0 ± 0%   -6.56%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:b?|c+)$-4                 67.0 ± 0%      63.0 ± 0%   -5.97%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:bc)+$-4                   49.0 ± 0%      46.0 ± 0%   -6.12%  (p=0.008 n=5+5)
    CompileOnepass/^a(?:[bcd])+$-4                46.0 ± 0%      43.0 ± 0%   -6.52%  (p=0.008 n=5+5)
    CompileOnepass/^a((?:[bcd])+)$-4              53.0 ± 0%      49.0 ± 0%   -7.55%  (p=0.008 n=5+5)
    CompileOnepass/^a(:?b|c)*d$-4                  109 ± 0%       105 ± 0%   -3.67%  (p=0.008 n=5+5)
    CompileOnepass/^.bc(d|e)*$-4                  66.0 ± 0%      62.0 ± 0%   -6.06%  (p=0.008 n=5+5)
    CompileOnepass/^loooooooooooooooooo...-4     1.10k ± 0%     1.09k ± 0%   -0.91%  (p=0.008 n=5+5)
    
    Fixes #19735
    
    Change-Id: Ic68503aaa08e42fafcf7e11cf1f584d674f5ea7b
    Reviewed-on: https://go-review.googlesource.com/38750
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit aca58647b7b71a0efa2634e4ab67f8f0ba24d2e3
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 27 11:50:43 2017 -0700

    cmd/internal/obj: eliminate stray ctxt.Cursym write
    
    It is explicitly assigned in each of the
    assemblers as needed.
    I plan to remove Cursym entirely eventually,
    but this is a helpful intermediate step.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: Id7ddefae2def439af44d03053886ca8cc935731f
    Reviewed-on: https://go-review.googlesource.com/38727
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 73912a1b914251d29eb60c8e85262d931f895b90
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 27 15:25:44 2017 -0700

    cmd/vet: remove Peek from list of canonical methods
    
    It is insufficiently canonical;
    see the discussion at issue 19719.
    
    Fixes #19719
    
    Change-Id: I0559ff3b1b39d7bc4b446d104f36fdf8ce3da50e
    Reviewed-on: https://go-review.googlesource.com/38722
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit a0d6d3855f103db492afb43b243d2ed52959f575
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 27 14:48:24 2017 -0700

    cmd/compile: construct typename in walk instead of SSA conversion
    
    This eliminates references to lineno and
    other globals from ssa conversion.
    
    Passes toolstash-check.
    
    Updates #15756
    
    Change-Id: I9792074fab0036b42f454b79139d0b27db913fb5
    Reviewed-on: https://go-review.googlesource.com/38721
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 579297e1e14423a7d48733f5c1b475873bd3bfef
Author: wei xiao <wei.xiao@arm.com>
Date:   Mon Nov 28 10:38:48 2016 +0800

    cmd/asm: add support to shift operands on arm64
    
    Fixes: #18070
    Also added a test in: cmd/asm/internal/asm/testdata/arm64.s
    
    Change-Id: Icc43ff7383cc06b8eaccabd9ff0aefa61c4ecb88
    Reviewed-on: https://go-review.googlesource.com/33595
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 7e817859b349864c67cfcf4dfe3cb2f752463521
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 27 11:22:57 2017 -0700

    cmd/internal/obj: eliminate Curp
    
    Remove the global obj.Link.Curp.
    
    In asmz.go, replace the only use by passing it as an argument.
    In asm0.go and asm9.go, it was written but never read.
    In asm5.go and asm7.go, thread it through as an argument.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: I1a0faa89e768820f35d73a8b37ec8088d78d15f7
    Reviewed-on: https://go-review.googlesource.com/38715
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1acba7d4fafef57d44bbd757abce58d632ee8475
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 27 11:02:13 2017 -0700

    cmd/internal/obj: remove prasm
    
    Fold the printing of the offending instruction
    into the neighboring Diag call, if it is not
    already present.
    
    Change-Id: I310f1479e16a4d2a24ff3c2f7e2c60e5e2015c1b
    Reviewed-on: https://go-review.googlesource.com/38714
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2b4274d66767039bab5dee4639a7558b101f46a0
Author: Elias Naur <elias.naur@gmail.com>
Date:   Mon Mar 27 11:14:53 2017 +0200

    runtime/cgo: CFRelease result from CFBundleCopyResourceURL
    
    The result from CFBundleCopyResourceURL is owned by the caller. This
    CL adds the necessary CFRelease to release it after use.
    
    Fixes #19722
    
    Change-Id: I7afe22ef241d21922a7f5cef6498017e6269a5c3
    Reviewed-on: https://go-review.googlesource.com/38639
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>

commit f97830114497d88a980859440bea51b252d0601e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Mar 26 15:16:11 2017 -0700

    cmd/internal/dwarf: remove global encbuf
    
    The global encbuf helped avoid allocations.
    It is incompatible with a concurrent backend.
    To avoid a performance regression while removing it,
    introduce two optimizations.
    First, re-use a buffer in dwarf.PutFunc.
    Second, avoid a buffer entirely when the int
    being encoded fits in seven bits, which is about 75%
    of the time.
    
    Passes toolstash-check.
    
    Updates #15756
    
    
    name       old alloc/op    new alloc/op    delta
    Template      40.6MB ± 0%     40.6MB ± 0%  -0.08%  (p=0.001 n=8+9)
    Unicode       29.9MB ± 0%     29.9MB ± 0%    ~     (p=0.068 n=8+10)
    GoTypes        116MB ± 0%      116MB ± 0%  +0.05%  (p=0.043 n=10+9)
    SSA            864MB ± 0%      864MB ± 0%  +0.01%  (p=0.010 n=10+9)
    Flate         25.8MB ± 0%     25.8MB ± 0%    ~     (p=0.353 n=10+10)
    GoParser      32.2MB ± 0%     32.2MB ± 0%    ~     (p=0.353 n=10+10)
    Reflect       80.2MB ± 0%     80.2MB ± 0%    ~     (p=0.165 n=10+10)
    Tar           27.0MB ± 0%     26.9MB ± 0%    ~     (p=0.143 n=10+10)
    XML           42.8MB ± 0%     42.8MB ± 0%    ~     (p=0.400 n=10+9)
    
    name       old allocs/op   new allocs/op   delta
    Template        398k ± 0%       397k ± 0%  -0.20%  (p=0.002 n=8+9)
    Unicode         320k ± 0%       321k ± 1%    ~     (p=0.122 n=8+10)
    GoTypes        1.16M ± 0%      1.17M ± 0%    ~     (p=0.053 n=10+9)
    SSA            7.65M ± 0%      7.65M ± 0%    ~     (p=0.122 n=10+8)
    Flate           240k ± 1%       240k ± 1%    ~     (p=0.243 n=10+9)
    GoParser        322k ± 1%       322k ± 1%    ~     (p=0.481 n=10+10)
    Reflect        1.00M ± 0%      1.00M ± 0%    ~     (p=0.211 n=9+10)
    Tar             256k ± 0%       255k ± 1%    ~     (p=0.052 n=10+10)
    XML             400k ± 1%       400k ± 0%    ~     (p=0.631 n=10+10)
    
    
    Change-Id: Ia39d9de09232fdbfc9c9cec14587bbf6939c9492
    Reviewed-on: https://go-review.googlesource.com/38713
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 89ebe5bbca0638567771cb2b0376b059122a77ba
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Mar 27 16:32:04 2017 +0000

    net/http/httptest: don't panic on Close of user-constructed Server value
    
    If the user created an httptest.Server directly without using a
    constructor it won't have the new unexported 'client' field. So don't
    assume it's non-nil.
    
    Fixes #19729
    
    Change-Id: Ie92e5da66cf4e7fb8d95f3ad0f4e3987d3ae8b77
    Reviewed-on: https://go-review.googlesource.com/38710
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Kevin Burke <kev@inburke.com>

commit 4234d1decd853b2373c17340eb4c0033c0fe9566
Author: Austin Clements <austin@google.com>
Date:   Thu Mar 23 11:20:17 2017 -0400

    runtime: improve systemstack-on-Go stack message
    
    We reused the old C stack check mechanism for the implementation of
    //go:systemstack, so when we execute a //go:systemstack function on a
    user stack, the system fails by calling morestackc. However,
    morestackc's message still talks about "executing C code".
    
    Fix morestackc's message to reflect its modern usage.
    
    Change-Id: I7e70e7980eab761c0520f675d3ce89486496030f
    Reviewed-on: https://go-review.googlesource.com/38572
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 0476c7a7b50614f98f37a51ee2100af922288b68
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sat Mar 25 12:55:40 2017 +0100

    runtime/cgo: raise the thread-local storage slot search limit on Android
    
    On Android, the thread local offset is found by looping through memory
    starting at the TLS base address. The search is limited to
    PTHREAD_KEYS_MAX, but issue 19472 made it clear that in some cases, the
    slot is located further from the TLS base.
    
    The limit is merely a sanity check in case our assumptions about the
    thread-local storage layout are wrong, so this CL raises it to 384, which
    is enough for the test case in issue 19472.
    
    Fixes #19472
    
    Change-Id: I89d1db3e9739d3a7fff5548ae487a7483c0a278a
    Reviewed-on: https://go-review.googlesource.com/38636
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit aa4c2ca316b1553ae65a3e8afb4cf862e11b78d0
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sat Mar 25 14:54:28 2017 +0100

    runtime/pprof: fix proto tests on NetBSD
    
    The proto_test tests are failing on NetBSD:
    
    https://build.golang.org/log/a3a577144ac48c6ef8e384ce6a700ad30549fb78
    
    the failures seem similar to previous failures on Android:
    
    https://build.golang.org/log/b5786e0cd6d5941dc37b6a50be5172f6b99e22f0
    
    The Android failures where fixed by CL 37896. This CL is an attempt
    to fix the NetBSD failures with a similar fix.
    
    Change-Id: I3834afa5b32303ca226e6a31f0f321f66fef9a3f
    Reviewed-on: https://go-review.googlesource.com/38637
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4909ecc46266b5afa121ef5f29382172188dbc36
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Mar 26 12:47:20 2017 -0700

    cmd/internal/obj/x86: change AsmBuf.Lock to bool
    
    Follow-up to CL 38668.
    
    Passes toolstash-check -all.
    
    Change-Id: I78a62509c610b5184b5e7ef2c4aa146fc8038840
    Reviewed-on: https://go-review.googlesource.com/38670
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit cf2b32e71924140919562e6c221d6def24c2fd91
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Mar 26 08:05:40 2017 -0700

    cmd/internal/obj: eliminate Prog.Mode
    
    Follow-up to CL 38446.
    
    Passes toolstash-check -all.
    
    Change-Id: I04cadc058cbaa5f396136502c574e5a395a33276
    Reviewed-on: https://go-review.googlesource.com/38669
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4f122e82fecb135e6b56dbaf907b132401cf1f9f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Mar 26 07:53:06 2017 -0700

    cmd/internal/obj: move fields from obj.Link to x86.AsmBuf
    
    These fields are used to encode a single instruction.
    Add them to AsmBuf, which is also per-instruction,
    and which is not global.
    
    Updates #15756
    
    Change-Id: I0e5ea22ffa641b07291e27de6e2ff23b6dc534bd
    Reviewed-on: https://go-review.googlesource.com/38668
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 166160b44664e75ffc6e51cd795956fd586196e4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Mar 26 07:46:30 2017 -0700

    cmd/internal/obj/x86: make ctxt.Cursym local
    
    Thread it through as an argument instead of using a global.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: Ia8c6ce09b43dbb2e6c7d889ded8dbaeb5366048d
    Reviewed-on: https://go-review.googlesource.com/38667
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c06679f7b5101781d6b8e36cf884568243b539ce
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 25 21:09:16 2017 -0700

    cmd/internal/obj/x86: remove ctxt.Curp references
    
    Empirically, p == ctxt.Curp here.
    A scan of (the thousands of lines of) asm6.go
    shows no clear opportunity for them to diverge.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: I9f5ee9585a850fbe24be3b851d8fdc2c966c65ce
    Reviewed-on: https://go-review.googlesource.com/38665
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 214be5b302d6ae978064d7895c7dcd9b6c403421
Author: Keith Randall <khr@golang.org>
Date:   Sat Mar 25 16:55:42 2017 -0700

    cmd/compile: remove likely bits from generated assembly
    
    We don't need them any more since #15837 was fixed.
    
    Fixes #19718
    
    Change-Id: I13e46c62b321b2c9265f44c977b63bfb23163ca2
    Reviewed-on: https://go-review.googlesource.com/38664
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 2ebe1bdd894a4bbd71408ae5c49d906f52320d2e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 25 10:44:21 2017 -0700

    cmd/internal/obj: make x86's asmbuf a local variable
    
    The x86 assembler requires a buffer to build
    variable-length instructions.
    It used to be an obj.Link field.
    That doesn't play nicely with concurrent assembly.
    Move the AsmBuf type to the x86 package,
    where it belongs anyway,
    and make it a local variable.
    
    Passes toolstash-check -all.
    No compiler performance impact.
    
    Updates #15756
    
    Change-Id: I8014e52145380bfd378ee374a0c971ee5bada917
    Reviewed-on: https://go-review.googlesource.com/38663
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3b251e603d09739d466ff2f44c5a362e1ba5c0b1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 25 07:22:53 2017 -0700

    cmd/internal/obj: eagerly initialize x86 assembler
    
    Prior to this CL, instinit was called as needed.
    This does not work well in a concurrent backend.
    Initialization is very cheap; do it on startup instead.
    
    Passes toolstash-check -all.
    No compiler performance impact.
    
    Updates #15756
    
    Change-Id: Ifa5e82e8abf4504435e1b28766f5703a0555f42d
    Reviewed-on: https://go-review.googlesource.com/38662
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e76d6a456b734bd02f275303571c38bec295a6b5
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sat Mar 25 16:17:59 2017 -0600

    cmd/compile: add test for non interface type switch
    
    Ensure that we have a test for when the compiler
    encounters a type switch on a non-interface value.
    
    Change-Id: Icb222f986894d0190e1241ca65396b4950e7d14f
    Reviewed-on: https://go-review.googlesource.com/38661
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 783f166e659558414d846d16bbe65e6fe9c7099b
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 24 18:03:17 2017 -0700

    cmd/compile/internal/syntax: remove need for missing_statement (fixed TODO)
    
    Now that we have consistent use of xOrNil parse methods, we don't
    need a special missing_statement singleton to distinguish between
    missing actually statements and other errors (which have returned
    nil before).
    
    For #19663.
    
    Change-Id: I8364f1441bdf8dd966bcd6d8219b2a42d6b88abd
    Reviewed-on: https://go-review.googlesource.com/38656
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d1f5e5f48249c120c9eed301ed07d546c5c65698
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 24 16:23:21 2017 -0700

    cmd/compile/internal/syntax: always construct a correct syntax tree
    
    - parser creates sensible nodes in case of syntax errors instead of nil
    - a new BadExpr node is used in places where we can't do better
    - fixed error message for incorrect type switch guard
    - minor cleanups
    
    Fixes #19663.
    
    Change-Id: I028394c6db9cba7371f0e417ebf93f594659786a
    Reviewed-on: https://go-review.googlesource.com/38653
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ecc6a81617477ddfa961f44e309707a4f864104a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 24 16:15:10 2017 -0700

    cmd/compile: prevent modification of ONAME/OLITERAL/OTYPES nodes in walkexpr
    
    ONAME, OLITERAL, and OTYPE nodes can be shared between functions.
    In a concurrent backend, such nodes might be walked concurrently
    with being read in other functions.
    Arrange for them to be unmodified by walk.
    
    This is a follow-up to CL 38609.
    
    Passes toolstash-check.
    
    Updates #15756
    
    Change-Id: I03ff1d2c0ad81dafac3fd55caa218939cf7c0565
    Reviewed-on: https://go-review.googlesource.com/38655
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ef1ab0f0f4f56e9be490aaca43c799d2eeeed289
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 23 16:39:11 2017 -0700

    cmd/compile: enforce no uses of Curfn in backend
    
    Updates #15756
    
    Change-Id: Id8d65ca9a3f1a7f9ea43e26cdd5e7d3befef8ba0
    Reviewed-on: https://go-review.googlesource.com/38593
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 34975095d0fbf61e1c22c0adf71ca1568106862f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 22 20:27:54 2017 -0700

    cmd/compile: provide pos and curfn to temp
    
    Concurrent compilation requires providing an
    explicit position and curfn to temp.
    This implementation of tempAt temporarily
    continues to use the globals lineno and Curfn,
    so as not to collide with mdempsky's
    work for #19683 eliminating the Curfn dependency
    from func nod.
    
    Updates #15756
    Updates #19683
    
    Change-Id: Ib3149ca4b0740e9f6eea44babc6f34cdd63028a9
    Reviewed-on: https://go-review.googlesource.com/38592
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 34396adac1949438e8f059a199d813ac4619f158
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Fri Mar 17 23:03:34 2017 +0000

    net/http: Fix TestLinuxSendfile without strace permissions
    
    If go doesn't have permission to run strace, this test hangs while
    waiting for strace to run. Instead try invoking strace with
    Run() first - on fail skip and report error, otherwise run
    the test normally using strace.
    
    Also fix link to open mips64 issue in same test.
    
    Fixes #9711
    
    Change-Id: Ibbc5fbb143ea6d0f8b6cfdca4b385ef4c8960b3d
    Reviewed-on: https://go-review.googlesource.com/38633
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0b647ffb98fa40077b521ea0b4aa1dfa5dcb7912
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 24 16:22:08 2017 -0700

    cmd/compile: combine walkexpr cases
    
    The type switch in walkexpr is giant.
    Shrink it a little by coalescing identical cases
    and removing some vertical whitespace.
    
    No functional changes.
    
    Passes toolstash-check.
    
    Change-Id: I7f7efb4faae1f8657dfafac04585172f99d8b37d
    Reviewed-on: https://go-review.googlesource.com/38652
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2c50bffee1c49d71ec1ea2c6cad10570ad780a4b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 24 11:33:29 2017 -0700

    cmd/compile: simplify funcsyms
    
    Sym.Fsym is used only to avoid adding duplicate
    entries to funcsyms, but that is easily
    accomplished by detecting the first lookup
    vs subsequent lookups of the func sym name.
    
    This avoids creating an unnecessary ONAME node
    during funcsym, which eliminates a dependency
    in the backend on Curfn and lineno.
    
    It also makes the code a lot simpler and clearer.
    
    Updates #15756
    
    Passes toolstash-check -all.
    No compiler performance changes.
    funcsymname does generate garbage via string
    concatenation, but it is not called very much,
    and this CL also eliminates allocation of several
    Nodes and Names.
    
    Change-Id: I7116c78fa39d975b7bd2c65a1d228749cf0dd46b
    Reviewed-on: https://go-review.googlesource.com/38605
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 3a89065c6c72162ac9c5f7b268c3e46ebee3cc7f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 24 15:57:12 2017 -0700

    cmd/compile: replace nod(ONAME) with newname
    
    Passes toolstash-check -all.
    
    Change-Id: Ib9f969e5ecc1537b7eab186dc4fd504a50f800f2
    Reviewed-on: https://go-review.googlesource.com/38586
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e67d881bc3708d38fbe485d2264f38a699ce11fd
Author: Keith Randall <khr@golang.org>
Date:   Fri Mar 24 14:03:15 2017 -0700

    cmd/compile: simplify efaceeq and ifaceeq
    
    Clean up code that does interface equality. Avoid doing checks
    in efaceeq/ifaceeq that we already did before calling those routines.
    
    No noticeable performance changes for existing benchmarks.
    
    name            old time/op  new time/op  delta
    EfaceCmpDiff-8   604ns ± 1%   553ns ± 1%  -8.41%  (p=0.000 n=9+10)
    
    Fixes #18618
    
    Change-Id: I3bd46db82b96494873045bc3300c56400bc582eb
    Reviewed-on: https://go-review.googlesource.com/38606
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Chase <drchase@google.com>

commit c026c37f33c6037dcf71e16a1e79f78f3b5165c4
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 24 13:10:24 2017 -0700

    cmd/compile/internal/gc: remove unused parameter to importfile
    
    Change-Id: Icf69862554d0121ec24e3c162d5c48630a03b99a
    Reviewed-on: https://go-review.googlesource.com/38583
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b3a8beb9d1eb7bd28b22e3a26f65c6025017d7a3
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 17:09:07 2017 -0700

    cmd/compile: minor cleanup in debug code
    
    Change-Id: I9885606801b9c8fcb62c16d0856025c4e83e658b
    Reviewed-on: https://go-review.googlesource.com/38650
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 26be4b9113e9c4f14399388652e961137658b6cb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 24 13:30:19 2017 -0700

    cmd/compile: avoid an assignment of n.Type in walk
    
    In the future, walk will probably run concurrently
    with SSA construction. It is possible for walk
    to be walking a function node that is referred
    to by another function undergoing SSA construction.
    In that case, this particular assignment to n.Type
    is race-y.
    
    This assignment is also not necessary;
    evconst does not change the type of n.
    Both arguments to evconst must have the same type,
    and at the end of evconst, n is replaced with n.Left.
    
    Remove the assignment, and add a check to ensure
    that its removal remains correct.
    
    Updates #15756
    
    Change-Id: Id95faaff42d5abd76be56445d1d3e285775de8bf
    Reviewed-on: https://go-review.googlesource.com/38609
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5e954047bcb77b219629676db3e5d057bed48360
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 24 11:43:08 2017 -0700

    cmd/compile: be slightly more tolerant in case of certain syntax errors
    
    Avoid construction of incorrect syntax trees in presence of errors.
    
    For #19663.
    
    Change-Id: I43025a3cf0fe02cae9a57e8bb9489b5f628c3fd7
    Reviewed-on: https://go-review.googlesource.com/38604
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 0b9607d1d648ae77d2db86a991db4a1fe921dbd8
Author: Filip Gruszczyński <gruszczy@gmail.com>
Date:   Wed Mar 15 20:11:30 2017 -0700

    encoding/gob: Speedup map decoding by reducing the allocations.
    
    The improvementis achieved in encoding/gob/decode.go decodeMap by
    allocate keyInstr and elemInstr only once and pass it to
    decodeIntoValue, instead of allocating a new instance on every loop
    cycle.
    
    name                     old time/op  new time/op  delta
    DecodeComplex128Slice-8  64.2µs ±10%  62.2µs ± 8%     ~     (p=0.686 n=4+4)
    DecodeFloat64Slice-8     37.1µs ± 3%  36.5µs ± 5%     ~     (p=0.343 n=4+4)
    DecodeInt32Slice-8       33.7µs ± 3%  32.7µs ± 4%     ~     (p=0.200 n=4+4)
    DecodeStringSlice-8      59.7µs ± 5%  57.3µs ± 1%     ~     (p=0.114 n=4+4)
    DecodeInterfaceSlice-8    543µs ± 7%   497µs ± 3%     ~     (p=0.057 n=4+4)
    DecodeMap-8              3.78ms ± 8%  2.66ms ± 2%  -29.69%  (p=0.029 n=4+4)
    
    Updates #19525
    
    Change-Id: Iec5fa4530de76f0a70da5de8a129a567b4aa096e
    Reviewed-on: https://go-review.googlesource.com/38317
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f1e880386b668a26dd6b7afdea43c9041a917fa5
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Wed Mar 15 21:55:04 2017 +0000

    net/http: strip port from host in mux Handler
    
    This change strips the port in mux.Handler before attempting to
    match handlers and adds a test for a request with port.
    
    CONNECT requests continue to use the original path and port.
    
    Fixes #10463
    
    Change-Id: Iff3a2ca2b7f1d884eca05a7262ad6b7dffbcc30f
    Reviewed-on: https://go-review.googlesource.com/38194
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 369d1083a74b6a965a33510489ab381d937812ae
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 24 09:19:46 2017 -0700

    spec: for non-constant map keys, add reference to evaluation order section
    
    The section on map literals mentions constant map keys but doesn't say
    what happens for equal non-constant map keys - that is covered in the
    section on evaluation order. Added respective link for clarity.
    
    Fixes #19689.
    
    Change-Id: If9a5368ba02e8250d4bb0a1d60d0de26a1f37bbb
    Reviewed-on: https://go-review.googlesource.com/38598
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e00e57d67cea33d4faef8506ced6f3c2416cfd15
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 24 10:36:13 2017 -0700

    cmd/compile: ignore all unreachable values during simple phi insertion
    
    Simple phi insertion already had a heuristic to check
    for dead blocks, namely having no predecessors.
    When we stopped generating code for dead blocks,
    we eliminated some values contained in more subtle
    dead blocks, which confused phi insertion.
    Compensate by beefing up the reachability check.
    
    Fixes #19678
    
    Change-Id: I0081e4a46f7ce2f69b131a34a0553874a0cb373e
    Reviewed-on: https://go-review.googlesource.com/38602
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit ad8c17b70328b8771ed5bbfe9161cb98f1995b84
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 24 09:52:39 2017 -0700

    cmd/compile: don't export dead code in inlineable fuctions
    
    CL 37499 allows inlining more functions by ignoring dead code.
    However, that dead code can contain non-exportable constructs.
    Teach the exporter not to export dead code.
    
    Fixes #19679
    
    Change-Id: Idb1d3794053514544b6f1035d29262aa6683e1e7
    Reviewed-on: https://go-review.googlesource.com/38601
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit a69754e30c9582e830ba244578724449955d4160
Author: Keith Randall <keithr@alum.mit.edu>
Date:   Fri Mar 24 09:00:17 2017 -0700

    cmd/compile: unnamed parameters do not escape
    
    Fixes #19687
    
    Change-Id: I2e4769b4ec5812506df4ac5dc6bc6a7c5774ecb0
    Reviewed-on: https://go-review.googlesource.com/38600
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 7202341de92927484a3eed101d3b77653b8b8bd1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 24 09:36:31 2017 -0700

    cmd/compile: only SSA [0]T when T is SSA-able
    
    Almost never happens in practice.
    The compiler will generate reasonable code anyway,
    since assignments involving [0]T never do any work.
    
    Fixes #19696
    Fixes #19671
    
    Change-Id: I350d2e0c5bb326c4789c74a046ab0486b2cee49c
    Reviewed-on: https://go-review.googlesource.com/38599
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6652572b75bae4d358cf193a95f688d9b67c5722
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 23 16:39:04 2017 -0700

    cmd/compile: thread Curfn through to debuginfo
    
    Updates #15756
    
    Change-Id: I860dd45cae9d851c7844654621bbc99efe7c7f03
    Reviewed-on: https://go-review.googlesource.com/38591
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 3a1ce1085ad08296557e8a87573fae4634ce7d8e
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Mar 23 22:47:56 2017 -0400

    runtime: access _cgo_yield indirectly
    
    The darwin linker for ARM does not allow PC-relative relocation
    of external symbol in text section. Work around it by accessing
    it indirectly: putting its address in a global variable (which is
    not external), and accessing through that variable.
    
    Fixes #19684.
    
    Change-Id: I41361bbb281b5dbdda0d100ae49d32c69ed85a81
    Reviewed-on: https://go-review.googlesource.com/38596
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Elias Naur <elias.naur@gmail.com>

commit 48de5a85fbc452bfc3af7422cd8aba0fab132d3d
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Thu Feb 16 16:14:16 2017 +0100

    net/http: import updated idna package and adjust request.go
    
    Custom logic from request.go has been removed.
    
    Created by running: “go run gen.go -core” from x/text
    at fc7fa097411d30e6708badff276c4c164425590c.
    
    Fixes golang/go#17268
    
    Change-Id: Ie440d6ae30288352283d303e5126e5837f11bece
    Reviewed-on: https://go-review.googlesource.com/37111
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3e63cdf8507709ebfb0906a3dbdc14c402cc0cd6
Author: Alexandru Moșoi <brtzsnr@gmail.com>
Date:   Thu Mar 23 22:29:59 2017 +0100

    cmd/compile: optimize shift when counter has different type.
    
    We already handle n << (uint64(c)&63).
    This change also handles n << (uint8(c)&63)
    where the SSA compiler promotes the counter to 32 bits.
    
    Fixes #19681
    
    Change-Id: I9327d64a994286aa0dbf76eb995578880be6923a
    Reviewed-on: https://go-review.googlesource.com/38550
    Run-TryBot: Alexandru Moșoi <alexandru@mosoi.ro>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d039d01fe9786a35ba4f6beea79ff2e964990c97
Author: wei xiao <wei.xiao@arm.com>
Date:   Mon Nov 28 10:35:12 2016 +0800

    cmd/asm: fix TBZ/TBNZ instructions on arm64
    
    Fixes #18069
    Also added a test in: cmd/asm/internal/asm/testdata/arm64.s
    
    Change-Id: Iee400bda4f30503ea3c1dc5bb8301568f19c92d1
    Signed-off-by: Wei Xiao <wei.xiao@arm.com>
    Reviewed-on: https://go-review.googlesource.com/33594
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 1911087dee1a6544067e94cbf430f1fd6e20cf23
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 22 20:28:12 2017 -0700

    cmd/compile: eliminate all references to Curfn in liveness
    
    Updates #15756
    
    Change-Id: I5ad87ef44b8ee48e1294820e0b1ab0ec07c480eb
    Reviewed-on: https://go-review.googlesource.com/38590
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f498929cdbafea0ba346186841b38e8a5d282ca1
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 23 17:39:28 2017 -0700

    cmd/compile: remove global var importpkg in favor of simple bool
    
    Pass around the imported package explicitly instead of relying
    on a global variable.
    
    Unfortunately we still need a global variable to communicate to
    the typechecker that we're in an import, but the semantic load
    is significantly reduced as it's just a bool, set/reset in a
    couple of places only.
    
    Change-Id: I4ebeae4064eb76ca0c4e2a15e4ca53813f005c29
    Reviewed-on: https://go-review.googlesource.com/38595
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 39e22f04231511f2448897e3c392e98cdbf7abea
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 23 16:55:31 2017 -0700

    cmd/compile: pass in importpkg to importer rather than rely on global
    
    First step towards removing global var importpkg.
    
    Change-Id: Ifdda7c295e5720a7ff2da9baea17f03f190d48fa
    Reviewed-on: https://go-review.googlesource.com/38594
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 39c8d2b7faed06b0e91a1ad7906231f53aab45d1
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sat Apr 23 20:37:12 2016 -0700

    os: parse command line without shell32.dll
    
    Go uses CommandLineToArgV from shell32.dll to parse command
    line parameters. But shell32.dll is slow to load. Implement
    Windows command line parsing in Go. This should make starting
    Go programs faster.
    
    I can see these speed ups for runtime.BenchmarkRunningGoProgram
    
    on my Windows 7 amd64:
    name                old time/op  new time/op  delta
    RunningGoProgram-2  11.2ms ± 1%  10.4ms ± 2%  -6.63%  (p=0.000 n=9+10)
    
    on my Windows XP 386:
    name                old time/op  new time/op  delta
    RunningGoProgram-2  19.0ms ± 3%  12.1ms ± 1%  -36.20%  (p=0.000 n=10+10)
    
    on @egonelbre Windows 10 amd64:
    name                old time/op  new time/op  delta
    RunningGoProgram-8  17.0ms ± 1%  15.3ms ± 2%  -9.71%  (p=0.000 n=10+10)
    
    This CL is based on CL 22932 by John Starks.
    
    Fixes #15588.
    
    Change-Id: Ib14be0206544d0d4492ca1f0d91fac968be52241
    Reviewed-on: https://go-review.googlesource.com/37915
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cc48b01883271920f7c111b0815790492b2a95c7
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Mar 23 13:38:15 2017 -0700

    cmd/compile/internal/gc: cleanup FuncDecl noding
    
    Collapse funcHeader into funcDecl.
    Initialize pragmas earlier.
    Move empty / non-empty body errors closer to fun.Body handling.
    Switch some yyerror to yyerrorl.
    
    Change-Id: I71fb7a3c0b77d656af560e4d88da894ba6183826
    Reviewed-on: https://go-review.googlesource.com/38475
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 25a1d5d0f472f0da82aa8b4a1f3322463df5bbe6
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Mar 23 15:26:10 2017 -0700

    cmd/compile/internal/gc: remove a Curfn dependency from nod
    
    Change-Id: I5daeb8f00044c86bb10510afbc6886898e61ba15
    Reviewed-on: https://go-review.googlesource.com/38570
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 9330ef869ce9cf281ec5d78e73824a8b3a6c7c2b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 23 11:21:42 2017 -0700

    cmd/compile: use Widthptr instead of Types[Tptr].Width
    
    Change-Id: I21e3abcfd1859f933f55fe875476dec07e43b038
    Reviewed-on: https://go-review.googlesource.com/38466
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 0ca8701ea2a9d5972285fd7833b67f9f15a403c6
Author: Michael Fraenkel <michael.fraenkel@gmail.com>
Date:   Thu Mar 23 15:08:43 2017 -0400

    net/rpc: Create empty maps and slices as return type
    
    When a map or slice is used as a return type create an empty value
    rather than a nil value.
    
    Fixes #19588
    
    Change-Id: I577fd74956172329745d614ac37d4db8f737efb8
    Reviewed-on: https://go-review.googlesource.com/38474
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 70ea0ec30fe37326d24249d9c9330be1ad655a90
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 23 10:54:08 2017 -0700

    math/big: replace local versions of bitLen, nlz with math/bits versions
    
    Verified that BenchmarkBitLen time went down from 2.25 ns/op to 0.65 ns/op
    an a 2.3 GHz Intel Core i7, before removing that benchmark (now covered by
    math/bits benchmarks).
    
    Change-Id: I3890bb7d1889e95b9a94bd68f0bdf06f1885adeb
    Reviewed-on: https://go-review.googlesource.com/38464
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 536a2257fba6dd18c74506988bdf3d6a15e52831
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 23 19:40:16 2017 +0000

    Revert "cmd/compile: strengthen SetFields/Width safety guarantee"
    
    This reverts commit b1b4f67169c5ceb3c81ba900c5022722d28755ab.
    
    Reason for revert: Broke the build.
    
    Change-Id: I5c99779896e39137c93c77d016ce683c872a69d7
    Reviewed-on: https://go-review.googlesource.com/38532
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit b1b4f67169c5ceb3c81ba900c5022722d28755ab
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 23 11:40:47 2017 -0700

    cmd/compile: strengthen SetFields/Width safety guarantee
    
    It is currently possible in the compiler to create a struct type,
    calculate the widths of types that depend on it,
    and then alter the struct type.
    
    transformclosure has local protection against this.
    Protect against it at a deeper level.
    
    This is preparation to call dowidth automatically,
    rather than explicitly.
    
    Change-Id: Ic1578ca014610197cfe54a9f4d044d122a7217e8
    Reviewed-on: https://go-review.googlesource.com/38469
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 7ac2e413ebc42cb2e8cac73f72a679b60f9c906b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 23 10:59:41 2017 -0700

    cmd/compile: minor cleanup in widstruct
    
    Change-Id: I9e52a2c52b754568412d719b415f91a998d247fe
    Reviewed-on: https://go-review.googlesource.com/38467
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 899532487f88a68ed6b7861fb701a1d4fab9456b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 22 20:33:23 2017 -0700

    cmd/compile: don't crash when calling String on a TFUNCARGS Type
    
    Change-Id: If5eabd622700a6b82dc4961ae9174c9d907eedb7
    Reviewed-on: https://go-review.googlesource.com/38465
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e95989c1c1251d479e92a84180dd50384afdec8b
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Mar 22 11:21:35 2017 -0700

    cmd/compile/internal/gc: remove unnecessary bitvector in plive
    
    In livenessepilogue, if we save liveness information for instructions
    before updating liveout, we can avoid an extra bitvector temporary and
    some extra copying around.
    
    Passes toolstash-check -all.
    
    Change-Id: I10d5803167ef3eba2e9e95094adc7e3d33929cc7
    Reviewed-on: https://go-review.googlesource.com/38408
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit d8ed449d8eae5b39ffe227ef7f56785e978dd5e2
Author: Ronald G. Minnich <rminnich@gmail.com>
Date:   Wed Mar 22 14:40:55 2017 -0700

    os/exec: handle Unshareflags with CLONE_NEWNS
    
    In some newer Linux distros, systemd forces
    all mount namespaces to be shared, starting
    at /. This disables the CLONE_NEWNS
    flag in unshare(2) and clone(2).
    While this problem is most commonly seen
    on systems with systemd, it can happen anywhere,
    due to how Linux namespaces now work.
    
    Hence, to create a private mount namespace,
    it is not sufficient to just set
    CLONE_NEWS; you have to call mount(2) to change
    the behavior of namespaces, i.e.
    mount("none", "/", NULL, MS_REC|MS_PRIVATE, NULL)
    
    This is tested and working and we can now correctly
    start child process with private namespaces on Linux
    distros that use systemd.
    
    The new test works correctly on Ubuntu 16.04.2 LTS.
    It fails if I comment out the new Mount, and
    succeeds otherwise. In each case it correctly
    cleans up after itself.
    
    Fixes #19661
    
    Change-Id: I52240b59628e3772b529d9bbef7166606b0c157d
    Reviewed-on: https://go-review.googlesource.com/38471
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9ecfd177cfe9783919175780fe8f29a0e4a99f4e
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 23 10:07:25 2017 -0700

    math/big: fix TestFloatSetFloat64String
    
    A -0 constant is the same as 0. Use explicit negative zero
    for float64 -0.0. Also, fix two test cases that were wrong.
    
    Fixes #19673.
    
    Change-Id: Ic09775f29d9bc2ee7814172e59c4a693441ea730
    Reviewed-on: https://go-review.googlesource.com/38463
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ff80d8ba4b4c374da80dd2393a5cf0cba336cf3e
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Mar 23 12:21:43 2017 -0400

    cmd/compile: remove redundant checks in s390x SSA rules
    
    CL 38337 modified canMergeLoad to reject loads with multiple uses
    so it is no longer necessary to check this in the SSA rules.
    
    Change-Id: I03498390e778da1be8cb59ae0948e99289008315
    Reviewed-on: https://go-review.googlesource.com/38473
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c3a50ad3c79ed15a413c24d4ae3497d4d629ff04
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 22 16:43:42 2017 -0700

    cmd/compile: eliminate Prog-related globals
    
    Introduce a new type, gc.Progs, to manage
    generation of Progs for a function.
    Use it to replace globals pc and pcloc.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: I2206998d7c58fe2a76b620904909f2e1cec8a57d
    Reviewed-on: https://go-review.googlesource.com/38418
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 27bc723b5130d11e4d3bae9566a3043a2dff2b8a
Author: Keith Randall <khr@golang.org>
Date:   Tue Mar 21 17:12:33 2017 -0700

    cmd/compile: initialize loop depths
    
    Regalloc uses loop depths - make sure they are initialized!
    
    Test to make sure we aren't pushing spills into loops.
    
    This fixes a generated-code performance bug introduced with
    the better spill placement change:
    https://go-review.googlesource.com/c/34822/
    
    Update #19595
    
    Change-Id: Ib9f0da6fb588503518847d7aab51e569fd3fa61e
    Reviewed-on: https://go-review.googlesource.com/38434
    Reviewed-by: David Chase <drchase@google.com>

commit 86dc86b4f948e16001903879162e9cf8da8f0537
Author: Keith Randall <khr@golang.org>
Date:   Sat Mar 18 11:16:30 2017 -0700

    cmd/compile: don't merge load+op if other op arg is still live
    
    We want to merge a load and op into a single instruction
    
        l = LOAD ptr mem
        y = OP x l
    
    into
    
        y = OPload x ptr mem
    
    However, all of our OPload instructions require that y uses
    the same register as x. If x is needed past this instruction, then
    we must copy x somewhere else, losing the whole benefit of merging
    the instructions in the first place.
    
    Disable this optimization if x is live past the OP.
    
    Also disable this optimization if the OP is in a deeper loop than the load.
    
    Update #19595
    
    Change-Id: I87f596aad7e91c9127bfb4705cbae47106e1e77a
    Reviewed-on: https://go-review.googlesource.com/38337
    Reviewed-by: Ilya Tocar <ilya.tocar@intel.com>

commit d0ff9ece2b36bd4470b4388602a757d50c4c3607
Author: Ben Shi <powerman1st@163.com>
Date:   Thu Mar 2 05:24:53 2017 +0000

    cmd/internal/obj/arm: Fix wrong assembly in the arm assembler
    
    As #19357 reported,
    
    TST has 3 different sub forms
    TST $imme, Rx
    TST Ry << imme, Rx
    TST Ry << Rz, Rx
    
    just like CMP/CMN/TEQ has. But current arm assembler assembles all TST
    instructions wrongly. This patch fixes it and adds more tests.
    
    Fixes #19357
    
    Change-Id: Iafedccfaab2cbb2631e7acf259837a782e2e8e2f
    Reviewed-on: https://go-review.googlesource.com/37662
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c46226a6222eb8bc8b30b40e8ce09ebe00f5dade
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Mar 23 03:56:23 2017 +0000

    net: mark TestDialerDualStack as flaky
    
    It was already marked flaky for everything but the dashboard.
    Remove that restriction. It's just flaky overall.
    
    It's doing more harm than good.
    
    Updates #13324
    
    Change-Id: I36feff32a1b8681e77700f74b9c70cb4073268eb
    Reviewed-on: https://go-review.googlesource.com/38459
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit b5e1ae46adb2d3552d5df78e6a410f479ef1efb2
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Mar 22 20:50:32 2017 -0700

    cmd/compile: don't crash when reporting some syntax errors
    
    Fixes #19667.
    
    Change-Id: Iaa71e2020af123c1bd3ac25e0b760956688e8bdf
    Reviewed-on: https://go-review.googlesource.com/38458
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2ae79d0711c7ff039e440d1bf1d6a4e45193f533
Author: Dave Cheney <dave@cheney.net>
Date:   Thu Mar 23 15:05:31 2017 +1100

    cmd/compile/internal/gc: remove unused exporter.nesting field
    
    exporter.nesting was added in c7b9bd74 to mitigate #16369 which was
    closed in ee272bbf. Remove the exporter.nesting field as it is now unused.
    
    Change-Id: I07873d1a07d6a08b11994b817a1483ffc2f5e45f
    Reviewed-on: https://go-review.googlesource.com/38490
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6f4a4585f2080c74e7e6a409d275cbcd5e0136a5
Author: Dave Cheney <dave@cheney.net>
Date:   Thu Mar 23 10:56:58 2017 +1100

    cmd/internal/obj/mips: standardize on sys.MIPS and sys.MIPS64 constants
    
    CL 38446 introduced the use of the sys.ArchFamily type into the
    cmd/internal/obj/mips package and redefined the mips.Mips32 and
    mips.Mips64 constants in terms of their sys.ArchFamily counterparts.
    This CL removes these local declarations and consolidates on sys.MIPS
    and sys.MIPS64 respectively.
    
    Change-Id: Id7aab6c7fd0de42ff43dde605df6bd4c85a3d895
    Reviewed-on: https://go-review.googlesource.com/38287
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e74c6cd3c05fda74fc8cac7a24b22b8b55a2239d
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sat Mar 4 07:18:26 2017 +0100

    regexp: add ASCII fast path for context methods
    
    The step method implementations check directly if the next rune
    only needs one byte to be decoded and avoid calling utf8.DecodeRune
    for such ASCII characters.
    
    Introduce the same fast path optimization for rune decoding
    for the context methods.
    
    Results for regexp benchmarks that use the context methods:
    
    name                            old time/op  new time/op  delta
    AnchoredLiteralShortNonMatch-4  97.5ns ± 1%  94.8ns ± 2%  -2.80%  (p=0.000 n=45+43)
    AnchoredShortMatch-4             163ns ± 1%   160ns ± 1%  -1.84%  (p=0.000 n=46+47)
    NotOnePassShortA-4               742ns ± 2%   742ns ± 2%    ~     (p=0.440 n=49+50)
    NotOnePassShortB-4               535ns ± 1%   533ns ± 2%  -0.37%  (p=0.005 n=46+48)
    OnePassLongPrefix-4              169ns ± 2%   166ns ± 2%  -2.06%  (p=0.000 n=50+49)
    
    Change-Id: Ib302d9e8c63333f02695369fcf9963974362e335
    Reviewed-on: https://go-review.googlesource.com/38256
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8a16d7d40a371e61f6d30604224039cf9a46d106
Author: haya14busa <haya14busa@gmail.com>
Date:   Fri Mar 17 00:41:58 2017 +0900

    regexp: reduce allocs in regexp.Match for onepass regex
    
    There were no allocations in regexp.Match for *non* onepass regex
    because m.matchcap length is reset to zero (ncap=0 for regexp.Match).
    
    But, as for onepass regex, m.matchcap length remains as it is even when
    ncap=0 and it leads needless allocations.
    
    benchmark                                    old ns/op      new ns/op      delta
    BenchmarkMatch_onepass_regex/32-4      6465           4628           -28.41%
    BenchmarkMatch_onepass_regex/1K-4      208324         151558         -27.25%
    BenchmarkMatch_onepass_regex/32K-4     7230259        5834492        -19.30%
    BenchmarkMatch_onepass_regex/1M-4      234379810      166310682      -29.04%
    BenchmarkMatch_onepass_regex/32M-4     7903529363     4981119950     -36.98%
    
    benchmark                                    old MB/s     new MB/s     speedup
    BenchmarkMatch_onepass_regex/32-4      4.95         6.91         1.40x
    BenchmarkMatch_onepass_regex/1K-4      4.92         6.76         1.37x
    BenchmarkMatch_onepass_regex/32K-4     4.53         5.62         1.24x
    BenchmarkMatch_onepass_regex/1M-4      4.47         6.30         1.41x
    BenchmarkMatch_onepass_regex/32M-4     4.25         6.74         1.59x
    
    benchmark                                    old allocs     new allocs     delta
    BenchmarkMatch_onepass_regex/32-4      32             0              -100.00%
    BenchmarkMatch_onepass_regex/1K-4      1024           0              -100.00%
    BenchmarkMatch_onepass_regex/32K-4     32768          0              -100.00%
    BenchmarkMatch_onepass_regex/1M-4      1048576        0              -100.00%
    BenchmarkMatch_onepass_regex/32M-4     104559255      0              -100.00%
    
    benchmark                                    old bytes      new bytes     delta
    BenchmarkMatch_onepass_regex/32-4      512            0             -100.00%
    BenchmarkMatch_onepass_regex/1K-4      16384          0             -100.00%
    BenchmarkMatch_onepass_regex/32K-4     524288         0             -100.00%
    BenchmarkMatch_onepass_regex/1M-4      16777216       0             -100.00%
    BenchmarkMatch_onepass_regex/32M-4     2019458128     0             -100.00%
    
    Fixes #19573
    
    Change-Id: I033982d0003ebb0360bb40b92eb3941c781ec74d
    Reviewed-on: https://go-review.googlesource.com/38270
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9e6b79a5dfb2f6fe4301ced956419a0da83bd025
Author: Richard Musiol <mail@richard-musiol.de>
Date:   Sat Feb 25 18:37:17 2017 +0100

    syscall: use CLONE_VFORK and CLONE_VM
    
    This greatly improves the latency of starting a child process when
    the Go process is using a lot of memory. Even though the kernel uses
    copy-on-write, preparation for that can take up to several 100ms under
    certain conditions. All other goroutines are suspended while starting
    a subprocess so this latency directly affects total throughput.
    
    With CLONE_VM the child process shares the same memory with the parent
    process. On its own this would lead to conflicting use of the same
    memory, so CLONE_VFORK is used to suspend the parent process until the
    child releases the memory when switching to to the new program binary
    via the exec syscall. When the parent process continues to run, one
    has to consider the changes to memory that the child process did,
    namely the return address of the syscall function needs to be restored
    from a register.
    
    A simple benchmark has shown a difference in latency of 16ms vs. 0.5ms
    at 10GB memory usage. However, much higher latencies of several 100ms
    have been observed in real world scenarios. For more information see
    comments on #5838.
    
    Fixes #5838
    
    Change-Id: I6377d7bd8dcd00c85ca0c52b6683e70ce2174ba6
    Reviewed-on: https://go-review.googlesource.com/37439
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0a0186fb7832928c9b5b1966854a8abc31678ea8
Author: Sarah Adams <shadams@google.com>
Date:   Mon Mar 20 16:11:46 2017 -0700

    encoding/xml: unmarshal allow empty, non-string values
    
    When unmarshaling, if an element is empty, eg. '<tag></tag>', and
    destination type is int, uint, float or bool, do not attempt to parse
    value (""). Set to its zero value instead.
    
    Fixes #13417
    
    Change-Id: I2d79f6d8f39192bb277b1a9129727d5abbb2dd1f
    Reviewed-on: https://go-review.googlesource.com/38386
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1295b745d13fe1402d5b645c9c20cc3adf85d563
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Sun Mar 12 10:02:47 2017 +0000

    net/http: improve speed of default mux
    
    The DefaultServeMux included in net/http uses a map to store routes,
    but iterates all keys for every request to allow longer paths.
    
    This change checks the map for an exact match first.
    
    To check performance was better, BenchmarkServeMux has been added -
    this adds >100 routes and checks the matches.
    
    Exact matches are faster and more predictable on this benchmark
    and on most existing package benchmarks.
    
    https://perf.golang.org/search?q=upload:20170312.1
    
    ServeMux-4  2.02ms ± 2% 0.04ms ± 2%  −98.08%  (p=0.004 n=5+6)
    
    https://perf.golang.org/search?q=upload:20170312.2
    
    ReadRequestChrome-4     184MB/s  ± 8%   186MB/s  ± 1%   ~
    ReadRequestCurl-4       45.0MB/s ± 1%   46.2MB/s ± 1%   +2.71%
    Read...Apachebench-4    45.8MB/s ±13%   48.7MB/s ± 1%   ~
    ReadRequestSiege-4      63.6MB/s ± 5%   69.2MB/s ± 1%   +8.75%
    ReadRequestWrk-4        30.9MB/s ± 9%   34.4MB/s ± 2%   +11.25%
    
    Change-Id: I8afafcb956f07197419d545a9f1c03ecaa307384
    Reviewed-on: https://go-review.googlesource.com/38057
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b5f81eae17b68c9a34d23dcf4669e3d879781b35
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Mar 21 22:23:15 2017 -0700

    cmd/compile/internal/syntax: replace inlined statement lists with syntax.BlockStmt
    
    This simplifies the code and removes a premature optimization.
    It increases the amount of allocated syntax.Node space by ~0.4%
    for parsing all of std lib, which is negligible.
    
    Before the change (best of 5 runs):
    
      $ go test -run StdLib -fast
      parsed 1517022 lines (3394 files) in 793.487886ms (1911840 lines/s)
      allocated 387.086Mb (267B/line, 487.828Mb/s)
    
    After the change (best of 5 runs):
    
      $ go test -run StdLib -fast
      parsed 1516911 lines (3392 files) in 805.028655ms (1884294 lines/s)
      allocated 388.466Mb (268B/line, 482.549Mb/s)
    
    Change-Id: Id19d6210fdc62393862ba3b04913352d95c599be
    Reviewed-on: https://go-review.googlesource.com/38439
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e0329248d5cda9a6a6c1492a2fdeeacd982afc9c
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Mar 21 15:22:13 2017 -0700

    cmd/compile/internal/syntax: add position info for { and } braces
    
    This change adds position information for { and } braces in the
    source. There's a 1.9% increase in memory use for syntax.Nodes,
    which is negligible relative to overall compiler memory consumption.
    
    Parsing the std library (using syntax package only) and memory
    consumption before this change (fastest of 5 runs):
    
      $ go test -run StdLib -fast
      parsed 1516827 lines (3392 files) in 780.612335ms (1943124 lines/s)
      allocated 379.903Mb (486.673Mb/s)
    
    After this change (fastest of 5 runs):
    
      $ go test -run StdLib -fast
      parsed 1517022 lines (3394 files) in 793.487886ms (1911840 lines/s)
      allocated 387.086Mb (267B/line, 487.828Mb/s)
    
    While not an exact apples-to-apples comparison (the syntax package
    has changed and is also parsed), the overall impact is small.
    
    Also: Small improvements to nodes_test.go.
    
    Change-Id: Ib8a7f90bbe79de33d83684e33b1bf8dbc32e644a
    Reviewed-on: https://go-review.googlesource.com/38435
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ec512340148f80aa0be3da90f86043ff535c4081
Author: Sam Whited <sam@samwhited.com>
Date:   Wed Mar 8 16:12:58 2017 -0600

    encoding/xml: format test output using subtests
    
    Change-Id: I2d155c838935cd8427abd142a462ff4c56829715
    Reviewed-on: https://go-review.googlesource.com/37948
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0d3cd51c9c0eeec13aa0c2eb139659ebc7d09008
Author: Josselin Costanzi <josselin@costanzi.fr>
Date:   Wed Mar 22 20:26:33 2017 +0100

    bytes: fix typo in comment
    
    Change-Id: Ia739337dc9961422982912cc6a669022559fb991
    Reviewed-on: https://go-review.googlesource.com/38365
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 352e19c92c780e0c5592f1ddaa0b81b1ea1a66af
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 22 11:31:56 2017 -0700

    cmd/compile: eliminate Gins and Naddr
    
    Preparation for eliminating Prog-related globals.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: Ia199fcb282cc3a84903a6e92a3ce342c5faba79c
    Reviewed-on: https://go-review.googlesource.com/38409
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c644a76e1f215d4ae2b210e7c9134b2291077ad3
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Wed Mar 22 12:25:44 2017 -0300

    cmd/compile/internal/ppc64, cmd/compile/internal/ssa: Remove OldArch checks
    
    Starting in go1.9, the minimum processor requirement for ppc64 is POWER8.
    Therefore, the checks for OldArch and the code enabled by it are not necessary
    anymore.
    
    Updates #19074
    
    Change-Id: I33d6a78b2462c80d57c5dbcba2e13424630afab4
    Reviewed-on: https://go-review.googlesource.com/38404
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 189053aee2705a16ed660b7e036e9b6d825c0e9b
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Wed Mar 22 14:25:30 2017 -0300

    runtime/internal/atomic: Remove unnecessary checks for GOARCH_ppc64
    
    Starting in go1.9, the minimum processor requirement for ppc64 is POWER8. This
    means the checks for GOARCH_ppc64 in asm_ppc64x.s can be removed, since we can
    assume LBAR and STBCCC instructions (both from ISA 2.06) will always be
    available.
    
    Updates #19074
    
    Change-Id: Ib4418169cd9fc6f871a5ab126b28ee58a2f349e2
    Reviewed-on: https://go-review.googlesource.com/38406
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 22ea7fc1a9d3af3b09c823b60f7b7cc81cea4f48
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Mar 22 10:27:30 2017 -0700

    cmd/compile/internal/gc: make SSAGenFPJump a method of SSAGenState
    
    Change-Id: Ie22a08c93dfcfd4b336e7b158415448dd55b2c11
    Reviewed-on: https://go-review.googlesource.com/38407
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cfb3c8df62deb607726fbd8a7a90f4e67f990a27
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 22 10:17:46 2017 -0700

    cmd/internal/obj: eliminate Link.Asmode
    
    Asmode is always set to p.Mode,
    which is always set based on the arch family.
    Instead, use the arch family directly.
    
    Passes toolstash-check -all.
    
    Change-Id: Id982472dcc8eeb6dd22cac5ad2f116b54a44caee
    Reviewed-on: https://go-review.googlesource.com/38451
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit a470e5d4b823c7a3ada993d2e76f191d4c51555a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 21 22:23:45 2017 -0700

    cmd/internal/obj: eliminate Ctxt.Mode
    
    Replace Ctxt.Mode with a method, Ctxt.RegWidth,
    which is calculated directly off the arch info.
    
    I believe that Prog.Mode can also be removed; future CL.
    
    This is a step towards obj.Link immutability.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: Ifd7f8f6ed0a2fdc032d1dd306fcd695a14aa5bc5
    Reviewed-on: https://go-review.googlesource.com/38446
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 0a94daa3789a52bea9856f9f8b3fa32477eab28a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 20 08:01:28 2017 -0700

    cmd/compile: funnel SSA Prog creation through SSAGenState
    
    Step one in eliminating Prog-related globals.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: I3b777fb5a7716f2d9da3067fbd94c28ca894a465
    Reviewed-on: https://go-review.googlesource.com/38450
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 3b39f523e1181499827321cedd8b7370b14ee762
Author: Alan Donovan <adonovan@google.com>
Date:   Wed Mar 22 12:38:09 2017 -0400

    cmd/vet: -lostcancel: fix crash in ill-typed code
    
    Fixes golang/go#19656
    
    Change-Id: Ied20d3f25b6e147cc693a1dd1aeb9480adc6687e
    Reviewed-on: https://go-review.googlesource.com/38405
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>

commit b029e943448c2ddf6faa9007a8ba4ab76238f18b
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri Feb 10 06:22:39 2017 +0900

    net/http: fix possible nil pointer dereference in TestOnlyWriteTimeout
    
    TestOnlyWriteTimeout assumes wrongly that:
    - the Accept method of trackLastConnListener is called only once
    - the shared variable conn never becomes nil
    and crashes on some circumstances.
    
    Updates #19032.
    
    Change-Id: I61de22618cd90b84a2b6401afdb6e5d9b3336b12
    Reviewed-on: https://go-review.googlesource.com/36735
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0b4e8d00fe6267550e7e11149f6cd992e6edd04d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 21 20:58:00 2017 -0700

    cmd/compile: pass an explicit position to namedAuto
    
    To enable this, inline the call to nod and simplify.
    Eliminates a reference to lineno from the backend.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: I9c4bd77d10d727aa8f5e6c6bb16b0e05de165631
    Reviewed-on: https://go-review.googlesource.com/38441
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2f2cd557a642b96a697b09cc84c0a5e342d41d3b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 21 22:06:39 2017 -0700

    cmd/internal/obj: clean up brloop
    
    Add docs.
    Reduce indentation.
    
    Passes toolstash-check -all.
    
    Change-Id: I968d1af25989886ae9945052e05e211a107dde9c
    Reviewed-on: https://go-review.googlesource.com/38443
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 4c8023bfed57276498427ce17235c1bfe210d611
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Thu Mar 16 19:19:29 2017 +0200

    strconv: optimize decimal ints formatting with smallsString
    
    Benchmark results for GOARCH=amd64:
    
    name                                     old time/op  new time/op  delta
    FormatInt-4                              2.51µs ± 2%  2.40µs ± 2%   -4.51%  (p=0.000 n=9+10)
    AppendInt-4                              1.67µs ± 2%  1.61µs ± 3%   -3.74%  (p=0.000 n=9+9)
    FormatUint-4                              698ns ± 2%   643ns ± 3%   -7.95%  (p=0.000 n=10+8)
    AppendUint-4                              478ns ± 1%   418ns ± 2%  -12.61%  (p=0.000 n=8+10)
    AppendUintVarlen/1-4                     9.30ns ± 6%  9.15ns ± 1%     ~     (p=0.199 n=9+10)
    AppendUintVarlen/12-4                    9.12ns ± 0%  9.16ns ± 2%     ~     (p=0.307 n=9+9)
    AppendUintVarlen/123-4                   18.6ns ± 2%  18.7ns ± 0%     ~     (p=0.091 n=10+6)
    AppendUintVarlen/1234-4                  19.1ns ± 4%  17.7ns ± 1%   -7.35%  (p=0.000 n=10+9)
    AppendUintVarlen/12345-4                 21.5ns ± 3%  20.7ns ± 3%   -3.78%  (p=0.002 n=9+10)
    AppendUintVarlen/123456-4                23.5ns ± 3%  20.9ns ± 1%  -11.14%  (p=0.000 n=10+9)
    AppendUintVarlen/1234567-4               25.0ns ± 2%  23.6ns ± 7%   -5.48%  (p=0.004 n=9+10)
    AppendUintVarlen/12345678-4              26.8ns ± 2%  23.4ns ± 2%  -12.79%  (p=0.000 n=9+10)
    AppendUintVarlen/123456789-4             29.8ns ± 3%  26.5ns ± 5%  -11.03%  (p=0.000 n=10+10)
    AppendUintVarlen/1234567890-4            31.6ns ± 3%  26.9ns ± 3%  -14.95%  (p=0.000 n=10+9)
    AppendUintVarlen/12345678901-4           33.8ns ± 3%  29.3ns ± 5%  -13.21%  (p=0.000 n=10+10)
    AppendUintVarlen/123456789012-4          35.5ns ± 4%  29.2ns ± 4%  -17.82%  (p=0.000 n=10+10)
    AppendUintVarlen/1234567890123-4         37.6ns ± 4%  31.4ns ± 3%  -16.48%  (p=0.000 n=10+10)
    AppendUintVarlen/12345678901234-4        39.8ns ± 6%  32.0ns ± 7%  -19.60%  (p=0.000 n=10+10)
    AppendUintVarlen/123456789012345-4       40.7ns ± 0%  34.4ns ± 4%  -15.55%  (p=0.000 n=6+10)
    AppendUintVarlen/1234567890123456-4      45.4ns ± 6%  35.1ns ± 4%  -22.66%  (p=0.000 n=10+10)
    AppendUintVarlen/12345678901234567-4     45.1ns ± 1%  36.7ns ± 4%  -18.77%  (p=0.000 n=9+10)
    AppendUintVarlen/123456789012345678-4    46.9ns ± 0%  36.4ns ± 3%  -22.49%  (p=0.000 n=9+10)
    AppendUintVarlen/1234567890123456789-4   50.6ns ± 6%  38.8ns ± 3%  -23.28%  (p=0.000 n=10+10)
    AppendUintVarlen/12345678901234567890-4  51.3ns ± 2%  38.4ns ± 0%  -25.00%  (p=0.000 n=9+8)
    
    Benchmark results for GOARCH=386:
    
    name                                     old time/op  new time/op  delta
    FormatInt-4                              6.21µs ± 0%  6.14µs ± 0%  -1.11%  (p=0.008 n=5+5)
    AppendInt-4                              4.95µs ± 0%  4.85µs ± 0%  -1.99%  (p=0.016 n=5+4)
    FormatUint-4                             1.89µs ± 1%  1.83µs ± 1%  -2.94%  (p=0.008 n=5+5)
    AppendUint-4                             1.59µs ± 0%  1.57µs ± 2%  -1.72%  (p=0.040 n=5+5)
    FormatIntSmall-4                         8.48ns ± 0%  8.48ns ± 0%    ~     (p=0.905 n=5+5)
    AppendIntSmall-4                         12.2ns ± 0%  12.2ns ± 0%    ~     (all equal)
    AppendUintVarlen/1-4                     10.6ns ± 1%  10.7ns ± 0%    ~     (p=0.238 n=5+4)
    AppendUintVarlen/12-4                    10.7ns ± 0%  10.7ns ± 1%    ~     (p=0.333 n=4+5)
    AppendUintVarlen/123-4                   29.9ns ± 1%  30.2ns ± 0%  +1.07%  (p=0.016 n=5+4)
    AppendUintVarlen/1234-4                  32.4ns ± 1%  30.4ns ± 0%  -6.30%  (p=0.008 n=5+5)
    AppendUintVarlen/12345-4                 35.1ns ± 2%  34.9ns ± 0%    ~     (p=0.238 n=5+5)
    AppendUintVarlen/123456-4                36.6ns ± 0%  35.3ns ± 0%  -3.55%  (p=0.029 n=4+4)
    AppendUintVarlen/1234567-4               38.9ns ± 0%  39.6ns ± 0%  +1.80%  (p=0.029 n=4+4)
    AppendUintVarlen/12345678-4              41.3ns ± 0%  40.1ns ± 0%  -2.91%  (p=0.000 n=5+4)
    AppendUintVarlen/123456789-4             44.9ns ± 1%  44.8ns ± 0%    ~     (p=0.667 n=5+5)
    AppendUintVarlen/1234567890-4            65.6ns ± 0%  66.2ns ± 1%  +0.88%  (p=0.016 n=4+5)
    AppendUintVarlen/12345678901-4           77.9ns ± 0%  76.3ns ± 0%  -2.00%  (p=0.000 n=4+5)
    AppendUintVarlen/123456789012-4          80.7ns ± 0%  79.1ns ± 1%  -2.01%  (p=0.008 n=5+5)
    AppendUintVarlen/1234567890123-4         83.6ns ± 0%  80.2ns ± 1%  -4.07%  (p=0.008 n=5+5)
    AppendUintVarlen/12345678901234-4        86.2ns ± 1%  83.3ns ± 0%  -3.39%  (p=0.008 n=5+5)
    AppendUintVarlen/123456789012345-4       88.5ns ± 0%  83.7ns ± 0%  -5.42%  (p=0.008 n=5+5)
    AppendUintVarlen/1234567890123456-4      90.6ns ± 0%  88.3ns ± 0%  -2.54%  (p=0.008 n=5+5)
    AppendUintVarlen/12345678901234567-4     92.7ns ± 0%  89.0ns ± 1%  -4.01%  (p=0.008 n=5+5)
    AppendUintVarlen/123456789012345678-4    95.6ns ± 1%  92.6ns ± 0%  -3.18%  (p=0.016 n=5+4)
    AppendUintVarlen/1234567890123456789-4    118ns ± 0%   114ns ± 0%    ~     (p=0.079 n=4+5)
    AppendUintVarlen/12345678901234567890-4   138ns ± 0%   136ns ± 0%  -1.45%  (p=0.008 n=5+5)
    
    Updates #19445
    
    Change-Id: Iafbe5c074898187c150dc3854e5b9fc19c10be05
    Reviewed-on: https://go-review.googlesource.com/38255
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 33266df861b6d231ed48b34ebb088cab92454de9
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 21 21:51:10 2017 -0700

    cmd/internal/obj: clean up checkaddr
    
    Coalesce identical cases.
    Give it a proper doc comment.
    Fix comment locations.
    Update/delete old comments.
    
    Passes toolstash-check -all.
    
    Change-Id: I88d9cf20e6e04b0c1c6583e92cd96335831f183f
    Reviewed-on: https://go-review.googlesource.com/38442
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 3394633db0e3b507f89f5c187fdfd50561fca8bc
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 21 22:21:33 2017 -0700

    cmd/internal/obj: eliminate AMODE
    
    AMODE appears to have been intended to allow
    a Prog to switch between 16 (!), 32, or 64 bit x86.
    It is unused anywhere in the tree.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: Ic57b257cfe580f29dad81d97e4193bf3c330c598
    Reviewed-on: https://go-review.googlesource.com/38445
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 88d4ab82d5f89f1421a39d9342bbd1c1e2ff08b0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 21 22:07:23 2017 -0700

    cmd/internal/obj: eliminate unnecessary ctxt.Cursym assignment
    
    None of the following code uses it.
    
    Passes toolstash-check -all.
    
    Updates #15756
    
    Change-Id: Ieeaaca8ba31e5c345c0c8a758d520b24be88e173
    Reviewed-on: https://go-review.googlesource.com/38444
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit 2c47c3e22e0d4f238a844da6904ce4e98bb4efd1
Author: Rob Pike <r@golang.org>
Date:   Tue Mar 21 20:27:47 2017 -0700

    cmd/doc: implement "go doc struct.field"
    
    By analogy with the handling of methods on types, show the documentation
    for a single field of a struct.
    
            % go doc ast.structtype.fields
            struct StructType {
                Fields *FieldList  // list of field declarations
            }
            %
    
    Fixes #19169.
    
    Change-Id: I002f992e4aa64bee667e2e4bccc7082486149842
    Reviewed-on: https://go-review.googlesource.com/38438
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 19040ac871bb6873b7fa4747f0ffe479a2ce0ef9
Author: Jason Travis <infomaniac7@gmail.com>
Date:   Tue Mar 21 20:12:47 2017 -0700

    test/bench/go1: fix typo in parserdata_test.go comment
    
    Change-Id: Iaca02660bdc8262db2b003a94aca661b5cec5576
    Reviewed-on: https://go-review.googlesource.com/38437
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f9fb4579e1766246d615566dec202bcd54e096a7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 18 08:11:49 2017 -0700

    cmd/compile: disable typPtr caching in the backend
    
    The only new Types that the backend introduces
    are pointers to Types generated by the frontend.
    Usually, when we generate a *T,
    we cache the resulting Type in T,
    to avoid recreating it later.
    However, that caching is not concurrency safe.
    Rather than add mutexes, this CL disables that
    caching before starting the backend.
    The backend generates few enough new *Ts that the
    performance impact of this is small, particularly
    if we pre-create some commonly used *Ts.
    
    Updates #15756
    
    name       old alloc/op    new alloc/op    delta
    Template      40.3MB ± 0%     40.4MB ± 0%  +0.18%  (p=0.001 n=10+10)
    Unicode       29.8MB ± 0%     29.8MB ± 0%  +0.11%  (p=0.043 n=10+9)
    GoTypes        114MB ± 0%      115MB ± 0%  +0.33%  (p=0.000 n=9+10)
    SSA            855MB ± 0%      859MB ± 0%  +0.40%  (p=0.000 n=10+10)
    Flate         25.7MB ± 0%     25.8MB ± 0%  +0.35%  (p=0.000 n=10+10)
    GoParser      31.9MB ± 0%     32.1MB ± 0%  +0.58%  (p=0.000 n=10+10)
    Reflect       79.6MB ± 0%     79.9MB ± 0%  +0.31%  (p=0.000 n=10+10)
    Tar           26.9MB ± 0%     26.9MB ± 0%  +0.21%  (p=0.000 n=10+10)
    XML           42.5MB ± 0%     42.7MB ± 0%  +0.52%  (p=0.000 n=10+9)
    
    name       old allocs/op   new allocs/op   delta
    Template        394k ± 1%       393k ± 0%    ~     (p=0.529 n=10+10)
    Unicode         319k ± 1%       319k ± 0%    ~     (p=0.720 n=10+9)
    GoTypes        1.15M ± 0%      1.15M ± 0%  +0.14%  (p=0.035 n=10+10)
    SSA            7.53M ± 0%      7.56M ± 0%  +0.45%  (p=0.000 n=9+10)
    Flate           238k ± 0%       238k ± 1%    ~     (p=0.579 n=10+10)
    GoParser        318k ± 1%       320k ± 1%  +0.64%  (p=0.001 n=10+10)
    Reflect        1.00M ± 0%      1.00M ± 0%    ~     (p=0.393 n=10+10)
    Tar             254k ± 0%       254k ± 1%    ~     (p=0.075 n=10+10)
    XML             395k ± 0%       397k ± 0%  +0.44%  (p=0.001 n=10+9)
    
    Change-Id: I6c031ed4f39108f26969c5712b73aa2fc08cd10a
    Reviewed-on: https://go-review.googlesource.com/38417
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2730c17a863e0eeb7afa3589608eece8cc50e6f3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Mar 21 22:16:16 2017 +0000

    syscall, os: fix FreeBSD 9 build
    
    I broke FreeBSD 9 in https://golang.org/cl/38426 by using Pipe2.
    
    We still want to support FreeBSD 9 for one last release (Go 1.9 will
    be the last), and FreeBSD 9 doesn't have Pipe2.
    
    So this still uses Pipe2, but falls back to Pipe on error.
    
    Updates #18854
    Updates #19072
    
    Change-Id: I1de90fb83606c93fb84b4b86fba31e207a702835
    Reviewed-on: https://go-review.googlesource.com/38430
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ee80afe326bab0a4829bd39186c4f343ac680a40
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 21 13:14:35 2017 -0700

    cmd/compile/internal/gc: remove unneeded effects cache fields
    
    Must have been lost when rebasing the SSA liveness CLs.
    
    Change-Id: Iaac33158cc7c92ea44a023c242eb914a7d6979c6
    Reviewed-on: https://go-review.googlesource.com/38427
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 01cd22c68792b659ca0912c104b14c86044110cb
Author: Josselin Costanzi <josselin@costanzi.fr>
Date:   Sun Mar 19 12:18:08 2017 +0100

    bytes: add optimized countByte for amd64
    
    Use SSE/AVX2 when counting a single byte.
    Inspired from runtime indexbyte implementation.
    
    Benchmark against previous implementation, where
    1 byte in every 8 is the one we are looking for:
    
    * On a machine without AVX2
    name               old time/op   new time/op     delta
    CountSingle/10-4    61.8ns ±10%     15.6ns ±11%    -74.83%  (p=0.000 n=10+10)
    CountSingle/32-4     100ns ± 4%       17ns ±10%    -82.54%  (p=0.000 n=10+9)
    CountSingle/4K-4    9.66µs ± 3%     0.37µs ± 6%    -96.21%  (p=0.000 n=10+10)
    CountSingle/4M-4    11.0ms ± 6%      0.4ms ± 4%    -96.04%  (p=0.000 n=10+10)
    CountSingle/64M-4    194ms ± 8%        8ms ± 2%    -95.64%  (p=0.000 n=10+10)
    
    name               old speed     new speed       delta
    CountSingle/10-4   162MB/s ±10%    645MB/s ±10%   +297.00%  (p=0.000 n=10+10)
    CountSingle/32-4   321MB/s ± 5%   1844MB/s ± 9%   +474.79%  (p=0.000 n=10+9)
    CountSingle/4K-4   424MB/s ± 3%  11169MB/s ± 6%  +2533.10%  (p=0.000 n=10+10)
    CountSingle/4M-4   381MB/s ± 7%   9609MB/s ± 4%  +2421.88%  (p=0.000 n=10+10)
    CountSingle/64M-4  346MB/s ± 7%   7924MB/s ± 2%  +2188.78%  (p=0.000 n=10+10)
    
    * On a machine with AVX2
    name               old time/op   new time/op     delta
    CountSingle/10-8    37.1ns ± 3%      8.2ns ± 1%    -77.80%  (p=0.000 n=10+10)
    CountSingle/32-8    66.1ns ± 3%      9.8ns ± 2%    -85.23%  (p=0.000 n=10+10)
    CountSingle/4K-8    7.36µs ± 3%     0.11µs ± 1%    -98.54%  (p=0.000 n=10+10)
    CountSingle/4M-8    7.46ms ± 2%     0.15ms ± 2%    -97.95%  (p=0.000 n=10+9)
    CountSingle/64M-8    124ms ± 2%        6ms ± 4%    -95.09%  (p=0.000 n=10+10)
    
    name               old speed     new speed       delta
    CountSingle/10-8   269MB/s ± 3%   1213MB/s ± 1%   +350.32%  (p=0.000 n=10+10)
    CountSingle/32-8   484MB/s ± 4%   3277MB/s ± 2%   +576.66%  (p=0.000 n=10+10)
    CountSingle/4K-8   556MB/s ± 3%  37933MB/s ± 1%  +6718.36%  (p=0.000 n=10+10)
    CountSingle/4M-8   562MB/s ± 2%  27444MB/s ± 3%  +4783.43%  (p=0.000 n=10+9)
    CountSingle/64M-8  543MB/s ± 2%  11054MB/s ± 3%  +1935.81%  (p=0.000 n=10+10)
    
    Fixes #19411
    
    Change-Id: Ieaf20b1fabccabe767c55c66e242e86f3617f883
    Reviewed-on: https://go-review.googlesource.com/38258
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 0ebaca6ba27534add5930a95acffa9acff182e2b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Mar 21 10:23:08 2017 -0700

    syscall, os: use pipe2 syscall on FreeBSD instead of pipe
    
    The pipe2 syscall exists in all officially supported FreeBSD
    versions: 10, 11 and future 12.
    The pipe syscall no longer exists in 11 and 12. To build and
    run Go on these versions, kernel needs COMPAT_FREEBSD10 option.
    
    Based on Gleb Smirnoff's https://golang.org/cl/38422
    
    Fixes #18854
    
    Change-Id: I8e201ee1b15dca10427c3093b966025d160aaf61
    Reviewed-on: https://go-review.googlesource.com/38426
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 051cbf3f3720086ec6d3fd159a234bae3ffd12ef
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 21 12:38:23 2017 -0700

    cmd/compile: add regress test for issue 19632
    
    Updates #19632.
    
    Change-Id: I1411dd997c8c6a789d17d0dcc0bfbd2281447b16
    Reviewed-on: https://go-review.googlesource.com/38401
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Daniel Martí <mvdan@mvdan.cc>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 10fdf769f07fdbc51cb50965a5de1674892057d6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 21 09:57:47 2017 -0700

    cmd/compile: remove function-name-based debuglive tweaks
    
    It's easier to grep output than recompile the compiler anyway.
    
    For concurrent compilation.
    
    Updates #15756
    
    Change-Id: I151cb5dc77056469cd9019d516f86454e931a197
    Reviewed-on: https://go-review.googlesource.com/38424
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 89840e4ac9525ddff2b04a8f17fe5b85e96f9bdc
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 21 12:13:12 2017 -0700

    cmd/compile: eliminate a Curfn reference in plive
    
    I think this got lost in a rebase somewhere.
    
    Updates #15756
    
    Change-Id: Ia3e7c60d1b9254f2877217073732b46c91059ade
    Reviewed-on: https://go-review.googlesource.com/38425
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit baa0fdd0934cb9dca88ea0effb46cf42089c9ccd
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 21 10:15:14 2017 -0700

    cmd/compile/internal/gc: fix liveness regression
    
    During AllocFrame, we drop unused variables from Curfn.Func.Dcl, but
    there might still be OpVarFoo instructions that reference those
    variables. This wasn't a problem though because gvardefx used to emit
    ANOP for unused variables instead of AVARFOO.
    
    As an easy fix, if we see OpVarFoo (or OpKeepAlive) referencing an
    unused variable, we can ignore it.
    
    Fixes #19632.
    
    Change-Id: I4e9ffabdb4058f7cdcc4663b540f5a5a692daf8b
    Reviewed-on: https://go-review.googlesource.com/38400
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 2e29eb57dbd2bc7b022fadc33943b0a5ee69324d
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sat Mar 18 15:55:41 2017 +0000

    runtime: remove unused *chantype parameters
    
    The chanrecv funcs don't use it at all. The chansend ones do, but the
    element type is now part of the hchan struct, which is already a
    parameter.
    
    hchan can be nil in chansend when sending to a nil channel, so when
    instrumenting we must copy to the stack to be able to read the channel
    type.
    
    name             old time/op  new time/op  delta
    ChanUncontended  6.42µs ± 1%  6.22µs ± 0%  -3.06%  (p=0.000 n=19+18)
    
    Initially found by github.com/mvdan/unparam.
    
    Fixes #19591.
    
    Change-Id: I3a5e8a0082e8445cc3f0074695e3593fd9c88412
    Reviewed-on: https://go-review.googlesource.com/38351
    Run-TryBot: Daniel Martí <mvdan@mvdan.cc>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c65ceff125ded084c6f3b47f830050339e7cc74e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Mar 21 06:34:35 2017 +0000

    flag: validate Int and Uint values to be in range
    
    Fixes #19230
    
    Change-Id: I38df9732b88f0328506e74f1a46f52adf47db1e5
    Reviewed-on: https://go-review.googlesource.com/38419
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 24dc8c6cb52aaa2680d0817298d109cb12098cda
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Mar 21 14:57:58 2017 +0100

    cmd/compile,runtime: fix atomic And8 for mipsle
    
    Removing stray xori that came from big endian copy/paste.
    Adding atomicand8 check to runtime.check() that would have revealed
    this error.
    Might fix #19396.
    
    Change-Id: If8d6f25d3e205496163541eb112548aa66df9c2a
    Reviewed-on: https://go-review.googlesource.com/38257
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 23bd9191361cccab5ca03aa2d65989efdf9d839c
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Mon Mar 13 10:16:30 2017 -0400

    cmd/compile: improve LoweredZero performance for ppc64x
    
    This change improves the performance of the LoweredZero rule
    on ppc64x.
    
    The improvement can be seen in the runtime ClearFat
    benchmarks:
    
    BenchmarkClearFat12-16       2.40          0.69          -71.25%
    BenchmarkClearFat16-16       9.98          0.93          -90.68%
    BenchmarkClearFat24-16       4.75          0.93          -80.42%
    BenchmarkClearFat32-16       6.02          0.93          -84.55%
    BenchmarkClearFat40-16       7.19          1.16          -83.87%
    BenchmarkClearFat48-16       15.0          1.39          -90.73%
    BenchmarkClearFat56-16       9.95          1.62          -83.72%
    BenchmarkClearFat64-16       18.0          1.86          -89.67%
    BenchmarkClearFat128-16      30.0          8.08          -73.07%
    BenchmarkClearFat256-16      52.5          11.3          -78.48%
    BenchmarkClearFat512-16      97.0          19.0          -80.41%
    BenchmarkClearFat1024-16     244           34.2          -85.98%
    
    Fixes: #19532
    
    Change-Id: If493e28bc1d8e61bc79978498be9f5336a36cd3f
    Reviewed-on: https://go-review.googlesource.com/38096
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit d972dc2de9d6b2b85656654b4d3a01dd02e446ec
Author: Volker Dobler <dr.volker.dobler@gmail.com>
Date:   Mon Mar 6 10:07:25 2017 +0100

    net/http/cookiejar: fix out-of-bounds errors on malformed domains
    
    The old implementation of Jar made the assumption that the host names
    in the URLs given to SetCookies() and Cookies() methods are well-formed.
    This is not an unreasonable assumption as malformed host names do not
    trigger calls to SetCookies or Cookies (at least not from net/http)
    as the HTTP request themselves are not executed. But there can be other
    invocations of these methods and at least on Linux it was possible to
    make DNS lookup to domain names with two trailing dots (see issue #7122).
    
    This is an old bug and this CL revives an old change (see
    https://codereview.appspot.com/52100043) to fix the issue. The discussion
    around 52100043 focused on the interplay between the jar and the
    public suffix list and who is responsible for which type if domain name
    canonicalization. The new bug report in issue #19384 used a nil public
    suffix list which demonstrates that the package cookiejar alone exhibits
    this problem and any solution cannot be fully delegated to the
    implementation of the used PublicSuffixList: Package cookiejar itself
    needs to protect against host names of the form ".." which triggered an
    out-of-bounds error.
    
    This CL does not address the issue of host name canonicalization and
    the question who is responsible for it. This CL just prevents the
    out-of-bounds error: It is a very conservative change, i.e. one might
    still set and retrieve cookies for host names like "weird.stuf...".
    Several more test cases document how the current code works.
    
    Fixes #19384.
    
    Change-Id: I14be080e8a2a0b266ced779f2aeb18841b730610
    Reviewed-on: https://go-review.googlesource.com/37843
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 5d6b7fcaa1444f6c17d519c9ce7bc0771bfd96ec
Author: Hugues Bruant <hugues.bruant@gmail.com>
Date:   Tue Mar 14 11:11:28 2017 -0700

    runtime: add mapdelete_fast*
    
    Add benchmarks for map delete with int32/int64/string key
    
    Benchmark results on darwin/amd64
    
    name                 old time/op  new time/op  delta
    MapDelete/Int32/1-8   151ns ± 8%    99ns ± 3%  -34.39%  (p=0.008 n=5+5)
    MapDelete/Int32/2-8   128ns ± 2%   111ns ±15%  -13.40%  (p=0.040 n=5+5)
    MapDelete/Int32/4-8   128ns ± 5%   114ns ± 2%  -10.82%  (p=0.008 n=5+5)
    MapDelete/Int64/1-8   144ns ± 0%   104ns ± 3%  -27.53%  (p=0.016 n=4+5)
    MapDelete/Int64/2-8   153ns ± 1%   126ns ± 3%  -17.17%  (p=0.008 n=5+5)
    MapDelete/Int64/4-8   178ns ± 3%   136ns ± 2%  -23.60%  (p=0.008 n=5+5)
    MapDelete/Str/1-8     187ns ± 3%   171ns ± 3%   -8.54%  (p=0.008 n=5+5)
    MapDelete/Str/2-8     221ns ± 3%   206ns ± 4%   -7.18%  (p=0.016 n=5+4)
    MapDelete/Str/4-8     256ns ± 5%   232ns ± 2%   -9.36%  (p=0.016 n=4+5)
    
    name                     old time/op    new time/op    delta
    BinaryTree17-8              2.78s ± 7%     2.70s ± 1%    ~     (p=0.151 n=5+5)
    Fannkuch11-8                3.21s ± 2%     3.19s ± 1%    ~     (p=0.310 n=5+5)
    FmtFprintfEmpty-8          49.1ns ± 3%    50.2ns ± 2%    ~     (p=0.095 n=5+5)
    FmtFprintfString-8         78.6ns ± 4%    80.2ns ± 5%    ~     (p=0.460 n=5+5)
    FmtFprintfInt-8            79.7ns ± 1%    81.0ns ± 3%    ~     (p=0.103 n=5+5)
    FmtFprintfIntInt-8          117ns ± 2%     119ns ± 0%    ~     (p=0.079 n=5+4)
    FmtFprintfPrefixedInt-8     153ns ± 1%     146ns ± 3%  -4.19%  (p=0.024 n=5+5)
    FmtFprintfFloat-8           239ns ± 1%     237ns ± 1%    ~     (p=0.246 n=5+5)
    FmtManyArgs-8               506ns ± 2%     509ns ± 2%    ~     (p=0.238 n=5+5)
    GobDecode-8                7.06ms ± 4%    6.86ms ± 1%    ~     (p=0.222 n=5+5)
    GobEncode-8                6.01ms ± 5%    5.87ms ± 2%    ~     (p=0.222 n=5+5)
    Gzip-8                      246ms ± 4%     236ms ± 1%  -4.12%  (p=0.008 n=5+5)
    Gunzip-8                   37.7ms ± 4%    37.3ms ± 1%    ~     (p=0.841 n=5+5)
    HTTPClientServer-8         64.9µs ± 1%    64.4µs ± 0%  -0.80%  (p=0.032 n=5+4)
    JSONEncode-8               16.0ms ± 2%    16.2ms ±11%    ~     (p=0.548 n=5+5)
    JSONDecode-8               53.2ms ± 2%    53.1ms ± 4%    ~     (p=1.000 n=5+5)
    Mandelbrot200-8            4.33ms ± 2%    4.32ms ± 2%    ~     (p=0.841 n=5+5)
    GoParse-8                  3.24ms ± 2%    3.27ms ± 4%    ~     (p=0.690 n=5+5)
    RegexpMatchEasy0_32-8      86.2ns ± 1%    85.2ns ± 3%    ~     (p=0.286 n=5+5)
    RegexpMatchEasy0_1K-8       198ns ± 2%     199ns ± 1%    ~     (p=0.310 n=5+5)
    RegexpMatchEasy1_32-8      82.6ns ± 2%    81.8ns ± 1%    ~     (p=0.294 n=5+5)
    RegexpMatchEasy1_1K-8       359ns ± 2%     354ns ± 1%  -1.39%  (p=0.048 n=5+5)
    RegexpMatchMedium_32-8      123ns ± 2%     123ns ± 1%    ~     (p=0.905 n=5+5)
    RegexpMatchMedium_1K-8     38.2µs ± 2%    38.6µs ± 8%    ~     (p=0.690 n=5+5)
    RegexpMatchHard_32-8       1.92µs ± 2%    1.91µs ± 5%    ~     (p=0.460 n=5+5)
    RegexpMatchHard_1K-8       57.6µs ± 1%    57.0µs ± 2%    ~     (p=0.310 n=5+5)
    Revcomp-8                   483ms ± 7%     441ms ± 1%  -8.79%  (p=0.016 n=5+4)
    Template-8                 58.0ms ± 1%    58.2ms ± 7%    ~     (p=0.310 n=5+5)
    TimeParse-8                 324ns ± 6%     312ns ± 2%    ~     (p=0.087 n=5+5)
    TimeFormat-8                330ns ± 1%     329ns ± 1%    ~     (p=0.968 n=5+5)
    
    name                     old speed      new speed      delta
    GobDecode-8               109MB/s ± 4%   112MB/s ± 1%    ~     (p=0.222 n=5+5)
    GobEncode-8               128MB/s ± 5%   131MB/s ± 2%    ~     (p=0.222 n=5+5)
    Gzip-8                   78.9MB/s ± 4%  82.3MB/s ± 1%  +4.25%  (p=0.008 n=5+5)
    Gunzip-8                  514MB/s ± 4%   521MB/s ± 1%    ~     (p=0.841 n=5+5)
    JSONEncode-8              121MB/s ± 2%   120MB/s ±10%    ~     (p=0.548 n=5+5)
    JSONDecode-8             36.5MB/s ± 2%  36.6MB/s ± 4%    ~     (p=1.000 n=5+5)
    GoParse-8                17.9MB/s ± 2%  17.7MB/s ± 4%    ~     (p=0.730 n=5+5)
    RegexpMatchEasy0_32-8     371MB/s ± 1%   375MB/s ± 3%    ~     (p=0.310 n=5+5)
    RegexpMatchEasy0_1K-8    5.15GB/s ± 1%  5.13GB/s ± 1%    ~     (p=0.548 n=5+5)
    RegexpMatchEasy1_32-8     387MB/s ± 2%   391MB/s ± 1%    ~     (p=0.310 n=5+5)
    RegexpMatchEasy1_1K-8    2.85GB/s ± 2%  2.89GB/s ± 1%    ~     (p=0.056 n=5+5)
    RegexpMatchMedium_32-8   8.07MB/s ± 2%  8.06MB/s ± 1%    ~     (p=0.730 n=5+5)
    RegexpMatchMedium_1K-8   26.8MB/s ± 2%  26.6MB/s ± 7%    ~     (p=0.690 n=5+5)
    RegexpMatchHard_32-8     16.7MB/s ± 2%  16.7MB/s ± 5%    ~     (p=0.421 n=5+5)
    RegexpMatchHard_1K-8     17.8MB/s ± 1%  18.0MB/s ± 2%    ~     (p=0.310 n=5+5)
    Revcomp-8                 527MB/s ± 6%   577MB/s ± 1%  +9.44%  (p=0.016 n=5+4)
    Template-8               33.5MB/s ± 1%  33.4MB/s ± 7%    ~     (p=0.310 n=5+5)
    
    Updates #19495
    
    Change-Id: Ib9ece1690813d9b4788455db43d30891e2138df5
    Reviewed-on: https://go-review.googlesource.com/38172
    Reviewed-by: Hugues Bruant <hugues.bruant@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f2b79cadfde436a824a12b40e096b4fe6c8249d7
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Mar 8 14:58:41 2017 +1100

    runtime: import os package in BenchmarkRunningGoProgram
    
    I would like to use BenchmarkRunningGoProgram to measure
    changes for issue #15588. So the program in the benchmark
    should import "os" package.
    
    It is also reasonable that basic Go program includes
    "os" package.
    
    For #15588.
    
    Change-Id: Ida6712eab22c2e79fbe91b6fdd492eaf31756852
    Reviewed-on: https://go-review.googlesource.com/37914
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 165a96e28131b703a1d318da245285bd98546bf6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 20 19:00:14 2017 -0700

    cmd/compile: fix pos of typenames created during SSA construction
    
    Prior to this CL, the function's position was used.
    The dottype Node's position is clearly better.
    
    I'm not thrilled about introducing a reference to
    lineno in the middle of SSA construction;
    I will have to remove it later.
    My immediate goal is stability and correctness of positions,
    though, since that aids refactoring, so this is an improvement.
    
    An example from package io:
    
    func (t *multiWriter) WriteString(s string) (n int, err error) {
            var p []byte // lazily initialized if/when needed
            for _, w := range t.writers {
                    if sw, ok := w.(stringWriter); ok {
                            n, err = sw.WriteString(s)
    
    The w.(stringWriter) type assertion includes loading
    the address of static type data for stringWriter:
    
    LEAQ    type."".stringWriter(SB), R10
    
    Prior to this CL, this instruction was given the line number
    of the function declaration.
    After this CL, this instruction is given the line number
    of the type assertion itself.
    
    Change-Id: Ifcca274b581a5a57d7e3102c4d7b7786bf307210
    Reviewed-on: https://go-review.googlesource.com/38389
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 67a46cc1636362bea54ab68b521f77a778968ea2
Author: Rob Pike <r@golang.org>
Date:   Mon Mar 20 20:24:26 2017 -0700

    encoding/gob: document the extra byte after a singleton
    
    This paragraph has been added, as the notion was missing from the
    documentation.
    
    If a value is passed to Encode and the type is not a struct (or pointer to struct,
    etc.), for simplicity of processing it is represented as a struct of one field.
    The only visible effect of this is to encode a zero byte after the value, just as
    after the last field of an encoded struct, so that the decode algorithm knows when
    the top-level value is complete.
    
    Fixes #16978
    
    Change-Id: I5f008e792d1b6fe80d2e026a7ff716608889db32
    Reviewed-on: https://go-review.googlesource.com/38414
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5c5a10690e56bf127832b98d73c83720e0093eef
Author: Rob Pike <r@golang.org>
Date:   Mon Mar 20 20:18:02 2017 -0700

    text/template,html/template: state that Funcs must happen before parsing
    
    Any method that affects the parse must happen before parsing.
    This obvious point is clear, but it's not clear to some that the
    set of defined functions affect the parse.
    
    Fixes #18971
    
    Change-Id: I8b7f8c8cf85b028c18e5ca3b9797de92ea910669
    Reviewed-on: https://go-review.googlesource.com/38413
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 0dafb7d9624c4c2bc876f03f2a6fa8d4a2ce6963
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 20 13:16:07 2017 -0700

    cmd/compile: check for missing function body earlier
    
    Tested by fixedbugs/issue3705.go.
    
    This removes a dependency on lineno
    from near the backend.
    
    Change-Id: I228bd0ad7295cf881b9bdeb0df9d18483fb96821
    Reviewed-on: https://go-review.googlesource.com/38382
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 49a533e212973d4403de640005fdcf88a23f91c4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 20 15:32:00 2017 -0700

    cmd/compile: use autogenerated position for init functions
    
    This eliminates an old TODO,
    and stabilizes the position information
    for init functions.
    
    Change-Id: Idf2d9a16a60e097ee08f42541b87e170da2f9d3a
    Reviewed-on: https://go-review.googlesource.com/38388
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 346d5883eb53d3e3a5bac7ee37beaa3db5c01fd8
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Fri Mar 17 16:05:31 2017 +0000

    go/build: remove unused returnImports parameter
    
    The code uses the filename suffix instead of the bool parameter to
    determine what to do.
    
    Fixes #19474.
    
    Change-Id: Ic552a54e50194592a4b4ae7f74d3109af54e6d36
    Reviewed-on: https://go-review.googlesource.com/38265
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 80c4b53e1e5159cc440e52c906583edc1eb79abc
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 20 12:27:25 2017 -0700

    cmd/compile/internal/gc: remove unneeded Type.Pos field
    
    Change-Id: I9ab650d9d2d0a99186009362454e1eabc9f6bad6
    Reviewed-on: https://go-review.googlesource.com/38393
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ee272bbf36afa97b51669e1e8d1aed4fcb7013ab
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 20 12:14:16 2017 -0700

    cmd/compile/internal/gc: export interface embedding information
    
    Fixes #16369.
    
    Change-Id: I23f8c36370d0da37ac5b5126d012d22f78782782
    Reviewed-on: https://go-review.googlesource.com/38392
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 07de3465be8efafd66c96552de38c2cbb5851f28
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 20 11:56:15 2017 -0700

    cmd/compile/internal/gc: handle recursive interfaces better
    
    Previously, we handled recursive interfaces by deferring typechecking
    of interface methods, while eagerly expanding interface embeddings.
    
    This CL switches to eagerly evaluating interface methods, and
    deferring expanding interface embeddings to dowidth. This allows us to
    detect recursive interface embeddings with the same mechanism used for
    detecting recursive struct embeddings.
    
    Updates #16369.
    
    Change-Id: If4c0320058047f8a2d9b52b9a79de47eb9887f95
    Reviewed-on: https://go-review.googlesource.com/38391
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 4e35e5fcabb504a14b9533692c9ae1a8c38b1139
Author: Jakob Borg <jakob@nym.se>
Date:   Tue Mar 14 08:21:51 2017 +0900

    net/http: fix ProxyFromEnvironment panic on invalid $NO_PROXY value
    
    Given an entry in $no_proxy like ":1" we would interpret it as an empty
    host name and a port number, then check the first character of the host
    name for dots. This would then cause an index out of range panic. This
    change simply skips these entries, as the following checks would anyway
    have returned false.
    
    Fixes #19536
    
    Change-Id: Iafe9c7a77ad4a6278c8ccb00a1575b56e4bdcd79
    Reviewed-on: https://go-review.googlesource.com/38067
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit df68afd07ce67727bcc2ad8e4afaa42dbcbf58e7
Author: Pascal S. de Kloe <pascal@quies.net>
Date:   Wed Nov 16 16:35:42 2016 +0100

    encoding/json: reduce unmarshal mallocs for unmapped fields
    
    JSON decoding performs poorly for unmapped and ignored fields. We noticed better
    performance when unmarshalling unused fields. The loss comes mostly from calls
    to scanner.error as described at #17914.
    
    benchmark                 old ns/op     new ns/op     delta
    BenchmarkIssue10335-8     431           408           -5.34%
    BenchmarkUnmapped-8       1744          1314          -24.66%
    
    benchmark                 old allocs     new allocs     delta
    BenchmarkIssue10335-8     4              3              -25.00%
    BenchmarkUnmapped-8       18             4              -77.78%
    
    benchmark                 old bytes     new bytes     delta
    BenchmarkIssue10335-8     320           312           -2.50%
    BenchmarkUnmapped-8       568           344           -39.44%
    
    Fixes #17914, improves #10335
    
    Change-Id: I7d4258a94eb287c0fe49e7334795209b90434cd0
    Reviewed-on: https://go-review.googlesource.com/33276
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7bb5b2d33a522d80e35f93f4c7bd6471d1d73552
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Sat Mar 11 17:17:15 2017 -0800

    cmd/internal/obj: remove unneeded Addr.Node and Prog.Opt fields
    
    Change-Id: I218b241c32a5948b66ad0d95ecc368648cf4ddf5
    Reviewed-on: https://go-review.googlesource.com/38130
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit cce4c319d6aa2fbcf16fbab6a3dc74baf482ce51
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 10 15:52:26 2017 -0800

    cmd/internal/obj: remove unneeded AVARFOO ops
    
    Change-Id: I10e36046ebce8a8741ef019cfe266b9ac9fa322d
    Reviewed-on: https://go-review.googlesource.com/38088
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 8ada52228cac7ef265915cf95d8a0e515069fe9d
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 10 15:03:25 2017 -0800

    cmd/compile: remove ProgInfo tables
    
    Change-Id: Id807c702ad71edddd23f2eb6f5e69e9a62e60bcd
    Reviewed-on: https://go-review.googlesource.com/38089
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 5dc14af6824ed31eab5a8a16e8e08082c5ddcb14
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Mar 17 13:48:56 2017 -0700

    runtime: clear signal stack on main thread
    
    This is a workaround for a FreeBSD kernel bug. It can be removed when
    we are confident that all people are using the fixed kernel. See #15658.
    
    Updates #15658.
    
    Change-Id: I0ecdccb77ddd0c270bdeac4d3a5c8abaf0449075
    Reviewed-on: https://go-review.googlesource.com/38325
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 325904fe6a6b3fc4324c517a62fa570fd6efb163
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Mar 9 18:32:17 2017 -0800

    cmd/compile: port liveness analysis to SSA
    
    Passes toolstash-check -all.
    
    Change-Id: I92c3c25d6c053f971f346f4fa3bbc76419b58183
    Reviewed-on: https://go-review.googlesource.com/38087
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit ea8c7dae4fe2fa9cc5ef258582086941aec751ae
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 10 16:03:56 2017 -0800

    cmd/compile: sort CFG blocks in PC order during liveness
    
    This CL changes the order that liveness analysis visits CFG blocks to
    PC order, rather than RPO. This doesn't meaningfully change anything
    except that the PCDATA_StackMapIndex values will be assigned in PC
    order too.
    
    However, this does have the benefit that the subsequent CL to port
    liveness analysis to the SSA CFG (which has blocks in PC order) will
    now pass toolstash-check.
    
    Change-Id: I1de5a2eecb8027723a6e422d46186d0c63d48c8d
    Reviewed-on: https://go-review.googlesource.com/38086
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 42a915c933e44784e70c9e61e4fe77c133580895
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 20 15:01:20 2017 -0700

    cmd/internal/obj: convert Debug* Link fields into bools
    
    Change-Id: I9ac274dbfe887675a7820d2f8f87b5887b1c9b0e
    Reviewed-on: https://go-review.googlesource.com/38383
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 236ef852bea6a154c5b4eab708ff535c6762b9ee
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 20 02:32:42 2017 -0700

    cmd/compile/internal/gc: split SetInterface from SetFields
    
    Change-Id: I4e568414faf64d3d47b1795382f0615f6caf53bc
    Reviewed-on: https://go-review.googlesource.com/38390
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit df47b821747184361f9b3d6038fc66eda4916619
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Mar 20 17:31:35 2017 -0400

    cmd/internal/obj/s390x: cleanup objz.go
    
    This CL deletes some unnecessary code in objz.go that existed to
    support instruction scheduling. It's likely instruction scheduling
    will never be done in this part of the backend so this code can
    just be deleted.
    
    This file can probably be cleaned up a bit more, but I think this
    is a good start.
    
    Passes: go build -toolexec 'toolstash -cmp' -a std.
    
    Change-Id: I1645632ac551a90a4f4be418045c046b488e9469
    Reviewed-on: https://go-review.googlesource.com/38394
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d68bb16b1e2f2fb55b347bdb6d1374103b5cb6a0
Author: philhofer <phofer@umich.edu>
Date:   Wed Mar 15 15:34:52 2017 -0700

    cmd/compile/internal/ssa: recognize constant pointer comparison
    
    Teach the backend to recognize that the address of a symbol
    is equal with itself, and that the addresses of two different
    symbols are different.
    
    Some examples of where this rule hits in the standard library:
    
     - inlined uses of (*time.Time).setLoc (e.g. time.UTC)
     - inlined uses of bufio.NewReader (via type assertion)
    
    Change-Id: I23dcb068c2ec333655c1292917bec13bbd908c24
    Reviewed-on: https://go-review.googlesource.com/38338
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f8b0231639859de7b8f1bfe7df1be0132aec9ad6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Mar 20 09:56:50 2017 -0700

    cmd/go, cmd/compile: always optimize when building runtime
    
    When optimizations are disabled, the compiler
    cannot eliminate enough write barriers to satisfy
    the runtime's nowritebarrier and nowritebarrierrec
    annotations.
    
    Enforce that requirement, and for convenience,
    have cmd/go elide -N when compiling the runtime.
    
    This came up in practice for me when running
    toolstash -cmp. When toolstash -cmp detected
    mismatches, it recompiled with -N, which caused
    runtime compilation failures.
    
    Change-Id: Ifcdef22c725baf2c59a09470f00124361508a8f3
    Reviewed-on: https://go-review.googlesource.com/38380
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a955ece6cdfb7490e184271398f0b51a0aa2ae8f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Mar 19 22:31:02 2017 -0700

    cmd/internal/obj: reduce variable scope
    
    Minor cleanup, to make it clearer
    that the two p's are unrelated.
    
    Change-Id: Icb6386c626681f60e5e631b33aa3a0fc84f40e4a
    Reviewed-on: https://go-review.googlesource.com/38381
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 07af21308c88c754fa2cd69a6d34d2b40b40c191
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 20 02:21:50 2017 -0700

    cmd/compile/internal/gc: eliminate two uses of Type.Pos
    
    Instead we can use t.nod.Pos.
    
    Change-Id: I643ee3226e402e38d4c77e8f328cbe83e55eac5c
    Reviewed-on: https://go-review.googlesource.com/38309
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit e22ba7f0fbd7a92418834dafbc1b539de1a85219
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Aug 30 19:11:19 2016 -0700

    cmd/compile: enable CSE of constant strings
    
    CL 27254 changed a constant string to a byte array
    in encoding/hex and got significant performance
    improvements.
    
    hex.Encode used the string twice in a single function.
    The rewrite rules lower constant strings into components.
    The pointer component requires an aux symbol.
    The existing implementation created a new aux symbol every time.
    As a result, constant string pointers were never CSE'd.
    Tighten then moved the pointer calculation next to the uses, i.e.
    into the loop.
    
    The re-use of aux syms enabled by this CL
    occurs 3691 times during make.bash.
    
    This CL should not go in without CL 38338
    or something like it.
    
    Change-Id: Ibbf5b17283c0e31821d04c7e08d995c654de5663
    Reviewed-on: https://go-review.googlesource.com/28219
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 25b51810014cefb8dba31321fcf40eb1a008fc3e
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Mar 20 13:08:26 2017 -0700

    cmd/gofmt: clarify doc string even more
    
    Since "columns of alignment" are terminated whenever indentation
    changes from one line to the next, alignment with spaces will work
    independent of the actually chosen tab width. Don't mention tab width
    anymore.
    
    Follow-up on https://golang.org/cl/38374/.
    
    For #19618.
    
    Change-Id: I58e47dfde57834f56a98d9119670757a12fb9c41
    Reviewed-on: https://go-review.googlesource.com/38379
    Reviewed-by: Rob Pike <r@golang.org>

commit 422c7fea70424a1443ad2841aabd262dc01bd9fe
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Mar 20 11:36:15 2017 -0700

    cmd/compile: don't permit declarations in post statement of for loop
    
    Report syntax error that was missed when moving to new parser.
    
    Fixes #19610.
    
    Change-Id: Ie5625f907a84089dc56fcccfd4f24df546042783
    Reviewed-on: https://go-review.googlesource.com/38375
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 17570a9afb5dc2d7d11eb3e132917e8d153a1ec9
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Mar 13 14:39:17 2017 -0400

    cmd/compile: emit fused multiply-{add,subtract} on ppc64x
    
    A follow on to CL 36963 adding support for ppc64x.
    
    Performance changes (as posted on the issue):
    
    poly1305:
    benchmark               old ns/op new ns/op delta
    Benchmark64-16          172       151       -12.21%
    Benchmark1K-16          1828      1523      -16.68%
    Benchmark64Unaligned-16 172       151       -12.21%
    Benchmark1KUnaligned-16 1827      1523      -16.64%
    
    math:
    BenchmarkAcos-16        43.9      39.9      -9.11%
    BenchmarkAcosh-16       57.0      45.8      -19.65%
    BenchmarkAsin-16        35.8      33.0      -7.82%
    BenchmarkAsinh-16       68.6      60.8      -11.37%
    BenchmarkAtan-16        19.8      16.2      -18.18%
    BenchmarkAtanh-16       65.5      57.5      -12.21%
    BenchmarkAtan2-16       45.4      34.2      -24.67%
    BenchmarkGamma-16       37.6      26.0      -30.85%
    BenchmarkLgamma-16      40.0      28.2      -29.50%
    BenchmarkLog1p-16       35.1      29.1      -17.09%
    BenchmarkSin-16         22.7      18.4      -18.94%
    BenchmarkSincos-16      31.7      23.7      -25.24%
    BenchmarkSinh-16        146       131       -10.27%
    BenchmarkY0-16          130       107       -17.69%
    BenchmarkY1-16          127       107       -15.75%
    BenchmarkYn-16          278       235       -15.47%
    
    Updates #17895.
    
    Change-Id: I1c16199715d20c9c4bd97c4a950bcfa69eb688c1
    Reviewed-on: https://go-review.googlesource.com/38095
    Reviewed-by: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit 01ac5b8dcfe5342af3770b0834220b87ea328fad
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Mar 20 10:58:40 2017 -0700

    cmd/gofmt: clarify documentation re: tab width
    
    Fixes #19618.
    
    Change-Id: I0ac450ff717ec1f16eb12758c6bf5e98b5de20e8
    Reviewed-on: https://go-review.googlesource.com/38374
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit eb6c1dd7ebd7035978eda533d1c0470261306bff
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Mar 20 16:44:34 2017 +0000

    net/http: deflake TestServerAllowsBlockingRemoteAddr more
    
    As noted in https://github.com/golang/go/issues/19161#issuecomment-287554171,
    CL 37771 (adding use of the new httptest.Server.Client to all net/http
    tests) accidentally reverted DisableKeepAlives for this test. For
    many tests, DisableKeepAlives was just present to prevent goroutines
    from staying active after the test exited.  In this case it might
    actually be important. (We'll see)
    
    Updates #19161
    
    Change-Id: I11f889f86c932b51b11846560b68dbe5993cdfc3
    Reviewed-on: https://go-review.googlesource.com/38373
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit d80166ebbef1e57457dc383959af83ea2286726e
Author: Martin Lindhe <martin.j.lindhe@gmail.com>
Date:   Mon Mar 20 12:14:47 2017 +0100

    crypto/*: fix spelling of 'below'
    
    Change-Id: Ic9d65206ec27f6d54bb71395802929e9c769e80a
    Reviewed-on: https://go-review.googlesource.com/38355
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit df6025bc0d7746fdf40a39398e5d8799ccf78a55
Author: Austin Clements <austin@google.com>
Date:   Thu Mar 16 17:02:24 2017 -0400

    runtime: disallow malloc or panic in scavenge
    
    Mallocs and panics in the scavenge path are particularly nasty because
    they're likely to silently self-deadlock on the mheap.lock. Avoid
    sinking lots of time into debugging these issues in the future by
    turning these into immediate throws.
    
    Change-Id: Ib36fdda33bc90b21c32432b03561630c1f3c69bc
    Reviewed-on: https://go-review.googlesource.com/38293
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 13ae271d5d007dcd630d9f43d6a43016b9af6e5c
Author: Austin Clements <austin@google.com>
Date:   Tue Mar 7 16:38:29 2017 -0500

    runtime: introduce a type for lfstacks
    
    The lfstack API is still a C-style API: lfstacks all have unhelpful
    type uint64 and the APIs are package-level functions. Make the code
    more readable and Go-style by creating an lfstack type with methods
    for push, pop, and empty.
    
    Change-Id: I64685fa3be0e82ae2d1a782a452a50974440a827
    Reviewed-on: https://go-review.googlesource.com/38290
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 2805d206890344f685579ac5b72ba2d9e5da485d
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sun Mar 19 09:51:22 2017 +0100

    cmd/compile: replace all uses of ptrto by typPtr
    
    This makes the overall naming and use of the functions
    to create a Type more consistent.
    
    Passes toolstash -cmp.
    
    Change-Id: Ie0d40b42cc32b5ecf5f20502675a225038ea40e4
    Reviewed-on: https://go-review.googlesource.com/38354
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 00f4cacb4994ba9004b79a1dd3329dd827785667
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sat Feb 25 00:34:21 2017 +0100

    cmd/compile: reduce allocs when appending to Node slices
    
    Rewrite Append function such that the *Node slice argument does not escape.
    
    Passes toolstash -cmp.
    
    name      old alloc/op    new alloc/op    delta
    Template     40.8MB ± 0%     40.8MB ± 0%  -0.17%  (p=0.000 n=20+19)
    Unicode      30.3MB ± 0%     30.2MB ± 0%  -0.11%  (p=0.000 n=19+20)
    GoTypes       115MB ± 0%      115MB ± 0%  -0.20%  (p=0.000 n=20+20)
    Compiler      492MB ± 0%      491MB ± 0%  -0.25%  (p=0.000 n=20+20)
    SSA           858MB ± 0%      858MB ± 0%  -0.08%  (p=0.000 n=20+20)
    Flate        26.2MB ± 0%     26.2MB ± 0%  -0.13%  (p=0.000 n=20+19)
    GoParser     32.5MB ± 0%     32.4MB ± 0%  -0.14%  (p=0.000 n=20+20)
    Reflect      80.6MB ± 0%     80.4MB ± 0%  -0.27%  (p=0.000 n=20+20)
    Tar          27.3MB ± 0%     27.3MB ± 0%  -0.12%  (p=0.000 n=20+19)
    XML          43.1MB ± 0%     43.0MB ± 0%  -0.14%  (p=0.000 n=20+20)
    
    name      old allocs/op   new allocs/op   delta
    Template       400k ± 1%       397k ± 0%  -0.81%  (p=0.000 n=20+18)
    Unicode        321k ± 1%       320k ± 0%  -0.43%  (p=0.000 n=20+20)
    GoTypes       1.17M ± 0%      1.16M ± 0%  -0.89%  (p=0.000 n=20+20)
    Compiler      4.59M ± 0%      4.54M ± 0%  -1.26%  (p=0.000 n=20+19)
    SSA           7.68M ± 0%      7.65M ± 0%  -0.37%  (p=0.000 n=18+18)
    Flate          242k ± 1%       240k ± 1%  -0.70%  (p=0.000 n=20+20)
    GoParser       323k ± 1%       321k ± 1%  -0.64%  (p=0.000 n=20+20)
    Reflect       1.01M ± 0%      1.00M ± 0%  -0.92%  (p=0.000 n=20+19)
    Tar            258k ± 1%       256k ± 1%  -0.60%  (p=0.000 n=20+19)
    XML            403k ± 1%       400k ± 0%  -0.78%  (p=0.000 n=20+20)
    
    Change-Id: Ie1eb603dc46f729574f6a76c08085b2619249be4
    Reviewed-on: https://go-review.googlesource.com/37437
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 81bcc470416e97a63570536fd6f1f191347a9adb
Author: Michel Lespinasse <walken@google.com>
Date:   Sat Mar 18 19:02:20 2017 -0700

    cmd/pprof: use proxy from environment
    
    See #18736
    
    Change-Id: I9c16357c05c16db677125d3077ee466b71559c7a
    Reviewed-on: https://go-review.googlesource.com/38343
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a68e5d94fad421e64d471dae1c5ee0b95b933242
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 18 22:00:28 2017 -0700

    cmd/compile: clean up SSA test API
    
    I noted in CL 38327 that the SSA test API felt a bit
    clunky after the ssa.Func/ssa.Cache/ssa.Config refactoring,
    and promised to clean it up once the dust settled.
    The dust has settled.
    
    Along the way, this CL fixes a potential latent bug,
    in which the amd64 test context was used for all dummy Syslook calls.
    The lone SSA test using the s390x context did not depend on the
    Syslook context being correct, so the bug did not arise in practice.
    
    Change-Id: If964251d1807976073ad7f47da0b1f1f77c58413
    Reviewed-on: https://go-review.googlesource.com/38346
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 62947bedd28a884b46f5df71070a9e86dad17081
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 18 14:39:48 2017 -0700

    cmd/compile: canonicalize empty interface types
    
    Mapping all empty interfaces onto the same Type
    allows better reuse of the ptrTo and sliceOf
    Type caches for *interface{} and []interface{}.
    
    This has little compiler performance impact now,
    but it will be helpful in the future,
    when we will eagerly populate some of those caches.
    
    Passes toolstash-check.
    
    Change-Id: I17daee599a129b0b2f5f3025c1be43d569d6782c
    Reviewed-on: https://go-review.googlesource.com/38344
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 872db7998937b310635a99055e066904425559bb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 18 10:16:03 2017 -0700

    cmd/compile: add more types to ssa.Types
    
    This reduces the number of calls back into the
    gc Type routines, which will help performance
    in a concurrent backend.
    It also reduces the number of callsites
    that must be considered in making the transition.
    
    Passes toolstash-check -all. No compiler performance changes.
    
    Updates #15756
    
    Change-Id: Ic7a8f1daac7e01a21658ae61ac118b2a70804117
    Reviewed-on: https://go-review.googlesource.com/38340
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit aea3aff66911e31cba9eddd93c02eb591ae483bf
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 16:04:46 2017 -0700

    cmd/compile: separate ssa.Frontend and ssa.TypeSource
    
    Prior to this CL, the ssa.Frontend field was responsible
    for providing types to the backend during compilation.
    However, the types needed by the backend are few and static.
    It makes more sense to use a struct for them
    and to hang that struct off the ssa.Config,
    which is the correct home for readonly data.
    Now that Types is a struct, we can clean up the names a bit as well.
    
    This has the added benefit of allowing early construction
    of all types needed by the backend.
    This will be useful for concurrent backend compilation.
    
    Passes toolstash-check -all. No compiler performance change.
    
    Updates #15756
    
    Change-Id: I021658c8cf2836d6a22bbc20cc828ac38c7da08a
    Reviewed-on: https://go-review.googlesource.com/38336
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2c397c7a753963494ce5dd5d7eda471354074698
Author: Damien Lespiau <damien.lespiau@gmail.com>
Date:   Sat Mar 18 19:44:37 2017 +0000

    cmd/gofmt: unindent the second line of a BUG note
    
    Currently, this second line is treated a pre-formatted text as it's
    indented relatively to the BUG() line.
    
    The current state can be seen at:
    
      https://golang.org/cmd/gofmt/#pkg-note-BUG
    
    Unindenting makes the rest of the sentence part of the same paragraph.
    
    Change-Id: I6dee55c9c321b1a03b41c7124c6a1ea15772c878
    Reviewed-on: https://go-review.googlesource.com/38353
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 32cb0ce65b39fc91923ac12a0a94f34b5dfd04be
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sat Mar 18 18:43:20 2017 +0100

    encoding/gob: speedup floats encoding and decoding
    
    By replacing bytes-reversing routines with bits.ReverseBytes64 calls.
    
    name                     old time/op  new time/op  delta
    EncodeComplex128Slice-4  35.1µs ± 1%  23.2µs ± 2%  -33.94%  (p=0.000 n=20+20)
    EncodeFloat64Slice-4     17.9µs ± 1%  11.0µs ± 1%  -38.36%  (p=0.000 n=17+18)
    
    name                     old time/op  new time/op  delta
    DecodeComplex128Slice-4  79.7µs ± 0%  69.9µs ± 1%  -12.31%  (p=0.000 n=20+20)
    DecodeFloat64Slice-4     47.3µs ± 1%  42.2µs ± 1%  -10.65%  (p=0.000 n=17+17)
    
    Change-Id: I91a6401c6009b5712fca6258dd1e57c6fe68ea64
    Reviewed-on: https://go-review.googlesource.com/38352
    Run-TryBot: Alberto Donizetti <alb.donizetti@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b6074a417dd707af3a9b39cc54769d7f8185961c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 16:59:32 2017 -0700

    cmd/compile: use testConfig consistently in SSA tests
    
    Change-Id: Iae41e14ee55eb4068fcb2189a77b345a7c5468b4
    Reviewed-on: https://go-review.googlesource.com/38333
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3928e847bdb9eb04b1fe6f914fff79c5c167b83a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 17 13:34:56 2017 -0700

    net: fix tests for /etc/hosts with entries named "test"
    
    Fixes #19592.
    
    Change-Id: I8946b33fd36ae1f39bdcc4bf0bd4b5b99618efe8
    Reviewed-on: https://go-review.googlesource.com/38300
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit dc4434a0c0aff7cb709b519dafdb8176a862e3ea
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 09:19:56 2017 -0700

    cmd/compile: make stkptrsize local
    
    While we're here, also eliminate a few more Curfn uses.
    
    Passes toolstash -cmp. No compiler performance impact.
    
    Updates #15756
    
    Change-Id: Ib8db9e23467bbaf16cc44bf62d604910f733d6b8
    Reviewed-on: https://go-review.googlesource.com/38331
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 758b5b3284b1e04131ad55ecd7da284bac463e38
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 09:10:57 2017 -0700

    cmd/compile: make Stksize local
    
    Passes toolstash -cmp. No compiler performance impact.
    
    Updates #15756
    
    Change-Id: I85b45244453ae28d4da76be4313badddcbf3f5dc
    Reviewed-on: https://go-review.googlesource.com/38330
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e0a5e69be2fcebe02c4ff94da31d3445c0e3cd60
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 08:58:36 2017 -0700

    cmd/compile: make Maxarg local
    
    Passes toolstash -cmp. No compiler performance impact.
    
    Updates #15756
    
    Change-Id: I1294058716d83dd1be495d399ed7ab2277754dc6
    Reviewed-on: https://go-review.googlesource.com/38329
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 174b858f7855077839335a7e54ff7b18d7a59a8c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 08:55:40 2017 -0700

    cmd/compile: pass frame size to defframe
    
    Preparation for de-globalizing Stksize and MaxArg.
    
    Passes toolstash -cmp. No compiler performance impact.
    
    Updates #15756
    
    Change-Id: I312f0bbd15587a6aebf472cd66c8e62b89e55c8a
    Reviewed-on: https://go-review.googlesource.com/38328
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d83af90a894cc6edaed97cea0edac62dfe8ba69a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Mar 17 23:53:19 2017 +0000

    Revert "go/types: enforce Check path restrictions via panics"
    
    This reverts commit b744a11a966ad3999c190fea9909ec8df0570b87.
    
    Reason for revert: Broke trybots. (misc-vetall builder is busted)
    
    Change-Id: I651d1c18db2fb3cb6ec12c2ae62024627baf8d77
    Reviewed-on: https://go-review.googlesource.com/38332
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit da8e939ba9707108d5a4100824546b8f26e9c6c3
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 07:49:22 2017 -0700

    cmd/compile: thread Curfn through SSA
    
    This is a first step towards eliminating the
    Curfn global in the backend.
    There's more to do.
    
    Passes toolstash -cmp. No compiler performance impact.
    
    Updates #15756
    
    Change-Id: Ib09f550a001e279a5aeeed0f85698290f890939c
    Reviewed-on: https://go-review.googlesource.com/38232
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ce584e516e79c2b4a3fc9570db695d2d4629485a
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Mar 17 11:53:24 2017 -0400

    cmd/compile: using a single Store op for non-pointer non-skip store
    
    This makes fewer Values around until decomposing, reducing
    allocation in compiler.
    
    name       old alloc/op    new alloc/op    delta
    Template      41.4MB ± 0%     40.8MB ± 0%  -1.29%  (p=0.000 n=10+10)
    Unicode       30.3MB ± 0%     30.2MB ± 0%  -0.24%  (p=0.000 n=10+10)
    GoTypes        118MB ± 0%      115MB ± 0%  -2.23%  (p=0.000 n=10+10)
    Compiler       505MB ± 0%      493MB ± 0%  -2.47%  (p=0.000 n=10+10)
    SSA            881MB ± 0%      872MB ± 0%  -1.03%  (p=0.000 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        401k ± 1%       400k ± 1%    ~     (p=0.631 n=10+10)
    Unicode         321k ± 0%       321k ± 1%    ~     (p=0.684 n=10+10)
    GoTypes        1.18M ± 0%      1.17M ± 0%  -0.34%  (p=0.000 n=10+10)
    Compiler       4.63M ± 0%      4.61M ± 0%  -0.43%  (p=0.000 n=10+10)
    SSA            7.83M ± 0%      7.82M ± 0%  -0.13%  (p=0.000 n=10+10)
    
    Change-Id: I8f736396294444248a439bd4c90be1357024ce88
    Reviewed-on: https://go-review.googlesource.com/38294
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 2cdb7f118ab86adb6fef5485d96831df3446b747
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 16 22:42:10 2017 -0700

    cmd/compile: move Frontend field from ssa.Config to ssa.Func
    
    Suggested by mdempsky in CL 38232.
    This allows us to use the Frontend field
    to associate frontend state and information
    with a function.
    See the following CL in the series for examples.
    
    This is a giant CL, but it is almost entirely routine refactoring.
    
    The ssa test API is starting to feel a bit unwieldy.
    I will clean it up separately, once the dust has settled.
    
    Passes toolstash -cmp.
    
    Updates #15756
    
    Change-Id: I71c573bd96ff7251935fce1391b06b1f133c3caf
    Reviewed-on: https://go-review.googlesource.com/38327
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 193510f2f6a0d01bb03595ba12dd2b05109980e3
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 17 10:50:20 2017 -0700

    cmd/compile: evaluate config as needed in rewrite rules
    
    Prior to this CL, config was an explicit argument
    to the SSA rewrite rules, and rules that needed
    a Frontend got at it via config.
    An upcoming CL moves Frontend from Config to Func,
    so rules can no longer reach Frontend via Config.
    Passing a Frontend as an argument to the rewrite rules
    causes a 2-3% regression in compile times.
    This CL takes a different approach:
    It treats the variable names "config" and "fe"
    as special and calculates them as needed.
    The "as needed part" is also important to performance:
    If they are calculated eagerly, the nilchecks themselves
    cause a regression.
    
    This introduces a little bit of magic into the rewrite
    generator. However, from the perspective of the rules,
    the config variable was already more or less magic.
    And it makes the upcoming changes much clearer.
    
    Passes toolstash -cmp.
    
    Change-Id: I173f2bcc124cba43d53138bfa3775e21316a9107
    Reviewed-on: https://go-review.googlesource.com/38326
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 09272ae981905dcdc76ab7ffbda996c49d86595c
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 17 13:35:36 2017 -0700

    cmd/compile/internal/gc: rename Thearch to thearch
    
    Prepared using gorename.
    
    Change-Id: Id55dac9ae5446a8bfeac06e7995b35f4c249eeca
    Reviewed-on: https://go-review.googlesource.com/38302
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 3e2f980e2721c05eb4a324b1e26080e082568f88
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 17 13:35:31 2017 -0700

    cmd/compile: eliminate direct uses of gc.Thearch in backends
    
    This CL changes the GOARCH.Init functions to take gc.Thearch as a
    parameter, which gc.Main supplies.
    
    Additionally, the x86 backend is refactored to decide within Init
    whether to use the 387 or SSE2 instruction generators, rather than for
    each individual SSA Value/Block.
    
    Passes toolstash-check -all.
    
    Change-Id: Ie6305a6cd6f6ab4e89ecbb3cbbaf5ffd57057a24
    Reviewed-on: https://go-review.googlesource.com/38301
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit aea44109cf23946332319ba51a4a373a9de432e6
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 16 21:13:29 2017 -0700

    strconv: replace small int string table with constant string
    
    This reduces memory use yet still provides the significant
    performance gain seen when using a fast path for small integers.
    
    Improvement of this CL comparing to code without fast path:
    
    name              old time/op  new time/op  delta
    FormatIntSmall-8  35.6ns ± 1%   4.5ns ± 1%  -87.30%  (p=0.008 n=5+5)
    AppendIntSmall-8  17.4ns ± 1%   9.4ns ± 3%  -45.70%  (p=0.008 n=5+5)
    
    For comparison, here's the improvement before this CL to code without
    fast path (1% better for FormatIntSmall):
    
    name              old time/op  new time/op  delta
    FormatIntSmall-8  35.6ns ± 1%   4.0ns ± 3%  -88.64%  (p=0.008 n=5+5)
    AppendIntSmall-8  17.4ns ± 1%   8.2ns ± 1%  -52.80%  (p=0.008 n=5+5)
    
    Thus, the code in this CL performs slower for small integers using fast
    path then the prior version, but this is relative to an already very fast
    version:
    
    name              old time/op  new time/op  delta
    FormatIntSmall-8  4.05ns ± 3%  4.52ns ± 1%  +11.81%  (p=0.008 n=5+5)
    AppendIntSmall-8  8.21ns ± 1%  9.45ns ± 3%  +15.05%  (p=0.008 n=5+5)
    
    Measured on 2.3 GHz Intel Core i7 running macOS Sierra 10.12.3.
    
    Overall, it's still ~88% faster than without fast path for small integers,
    so probably worth it as it removes 100 global string slices in favor of
    a single string.
    
    Credits: This is based on the original (but cleaned up) version of the
    code by Aliaksandr Valialkin (https://go-review.googlesource.com/c/37963/).
    
    Change-Id: Icda78679c8c14666d46257894e9fa3d7f35e58b8
    Reviewed-on: https://go-review.googlesource.com/38319
    Reviewed-by: Martin Möhrmann <moehrmann@google.com>

commit b744a11a966ad3999c190fea9909ec8df0570b87
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Fri Mar 17 15:05:53 2017 +0000

    go/types: enforce Check path restrictions via panics
    
    Its godoc says that path must not be empty or dot, while the existing
    implementation happily accepts both.
    
    Change-Id: I64766271c35152dc7adb21ff60eb05c52237e6b6
    Reviewed-on: https://go-review.googlesource.com/38262
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>

commit ed00cd94f2cd01f49ee8da8b1dc0c06b66d34b2f
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Fri Mar 17 18:42:58 2017 +0100

    encoding/gob: make integers encoding faster
    
    name                old time/op  new time/op  delta
    EncodeInt32Slice-4  14.6µs ± 2%  12.2µs ± 1%  -16.65%  (p=0.000 n=19+18)
    
    Change-Id: I078a171f1633ff81d7e3f981dc9a398309ecb2c0
    Reviewed-on: https://go-review.googlesource.com/38269
    Reviewed-by: Rob Pike <r@golang.org>

commit 42e97468a1fd4b9f08bccd076edb1598435c72fb
Author: Keith Randall <keithr@alum.mit.edu>
Date:   Thu Mar 16 22:34:38 2017 -0700

    cmd/compile: intrinsic for math/bits.Reverse on ARM64
    
    I don't know that it exists for any other architectures.
    
    Update #18616
    
    Change-Id: Idfe5dee251764d32787915889ec0be4bebc5be24
    Reviewed-on: https://go-review.googlesource.com/38323
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit be04da8f0c5cb57e2736cfac8c76971d7d8cfc6f
Author: Alexander Menzhinsky <amenzhinsky@gmail.com>
Date:   Fri Mar 3 15:57:19 2017 +0300

    cmd/go: fix race libraries rebuilding by `go test -i`
    
    `go test -i -race` adds the "sync/atomic" package to every package dependency tree
    that makes buildIDs different from packages installed with `go install -race`
    and causes cache rebuilding.
    
    Fixes #19133
    Fixes #19151
    
    Change-Id: I0536c6fa41b0d20fe361b5d35b3c0937b146d07d
    Reviewed-on: https://go-review.googlesource.com/37598
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b9f6b22a0100716db1f196a395f9ded8456337f8
Author: Alexey Neganov <neganovalexey@gmail.com>
Date:   Tue Mar 14 21:11:42 2017 +0300

    mime: handling invalid mime media parameters
    
    Sometimes it's necessary to deal with emails that do not follow the specification; in particular, it's possible to download such email via gmail.
    When the existing implementation handle invalid mime media parameters, it returns nils and error, although there is a valid media type, which may be returned.
    If this behavior changes, it may not affect any existing programs, but it will help to parse some emails.
    
    Fixes #19498
    
    Change-Id: Ieb2fdbddfd93857faee941d2aa49d59e286d57fd
    Reviewed-on: https://go-review.googlesource.com/38190
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b6cd22c277f49df493e3033430c32e56ecb949e3
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Mon Mar 13 16:25:15 2017 -0400

    hash/crc32: improve performance for ppc64le
    
    This change improves the performance of crc32 for ppc64le by using
    vpmsum and other vector instructions in the algorithm.
    
    The testcase was updated to test more sizes.
    
    Fixes #19570
    
    BenchmarkCRC32/poly=IEEE/size=15/align=0-8             90.5          81.8          -9.61%
    BenchmarkCRC32/poly=IEEE/size=15/align=1-8             89.7          81.7          -8.92%
    BenchmarkCRC32/poly=IEEE/size=40/align=0-8             93.2          61.1          -34.44%
    BenchmarkCRC32/poly=IEEE/size=40/align=1-8             92.8          60.9          -34.38%
    BenchmarkCRC32/poly=IEEE/size=512/align=0-8            501           55.8          -88.86%
    BenchmarkCRC32/poly=IEEE/size=512/align=1-8            502           132           -73.71%
    BenchmarkCRC32/poly=IEEE/size=1kB/align=0-8            947           69.9          -92.62%
    BenchmarkCRC32/poly=IEEE/size=1kB/align=1-8            946           144           -84.78%
    BenchmarkCRC32/poly=IEEE/size=4kB/align=0-8            3602          186           -94.84%
    BenchmarkCRC32/poly=IEEE/size=4kB/align=1-8            3603          263           -92.70%
    BenchmarkCRC32/poly=IEEE/size=32kB/align=0-8           28404         1338          -95.29%
    BenchmarkCRC32/poly=IEEE/size=32kB/align=1-8           28856         1405          -95.13%
    BenchmarkCRC32/poly=Castagnoli/size=15/align=0-8       89.7          81.8          -8.81%
    BenchmarkCRC32/poly=Castagnoli/size=15/align=1-8       89.8          81.9          -8.80%
    BenchmarkCRC32/poly=Castagnoli/size=40/align=0-8       93.8          61.4          -34.54%
    BenchmarkCRC32/poly=Castagnoli/size=40/align=1-8       94.3          61.3          -34.99%
    BenchmarkCRC32/poly=Castagnoli/size=512/align=0-8      503           56.4          -88.79%
    BenchmarkCRC32/poly=Castagnoli/size=512/align=1-8      502           132           -73.71%
    BenchmarkCRC32/poly=Castagnoli/size=1kB/align=0-8      941           70.2          -92.54%
    BenchmarkCRC32/poly=Castagnoli/size=1kB/align=1-8      943           145           -84.62%
    BenchmarkCRC32/poly=Castagnoli/size=4kB/align=0-8      3588          186           -94.82%
    BenchmarkCRC32/poly=Castagnoli/size=4kB/align=1-8      3595          264           -92.66%
    BenchmarkCRC32/poly=Castagnoli/size=32kB/align=0-8     28266         1323          -95.32%
    BenchmarkCRC32/poly=Castagnoli/size=32kB/align=1-8     28344         1404          -95.05%
    
    Change-Id: Ic4d8274c66e0e87bfba5f609f508a3877aee6bb5
    Reviewed-on: https://go-review.googlesource.com/38184
    Reviewed-by: David Chase <drchase@google.com>

commit 16663a85ba0b2814c47f54ddfdcb45782e10dc42
Author: Nigel Tao <nigeltao@golang.org>
Date:   Fri Mar 17 15:34:39 2017 +1100

    image/png: decode Gray8 transparent images.
    
    Fixes #19553.
    
    Change-Id: I414cb3b1c2dab20f41a7f4e7aba49c534ff19942
    Reviewed-on: https://go-review.googlesource.com/38271
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 88e47187c165e477a189049094e5e988b7651a6a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 15 22:43:40 2017 -0700

    cmd/compile: relocate code from config.go to func.go
    
    This is a follow-up to CL 38167.
    Pure code movement.
    
    Change-Id: I13e58f7eac6718c77076d89e13fc721a5205ec57
    Reviewed-on: https://go-review.googlesource.com/38322
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a5e3cac89587d2d6235e9a7217185dee9be6852a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 15 11:15:13 2017 -0700

    cmd/compile: rearrange fields between ssa.Func, ssa.Cache, and ssa.Config
    
    This makes ssa.Func, ssa.Cache, and ssa.Config fulfill
    the roles laid out for them in CL 38160.
    
    The only non-trivial change in this CL is how cached
    values and blocks get IDs. Prior to this CL, their IDs were
    assigned as part of resetting the cache, and only modified
    IDs were reset. This required knowing how many values and
    blocks were modified, which required a tight coupling between
    ssa.Func and ssa.Config. To eliminate that coupling,
    we now zero values and blocks during reset,
    and assign their IDs when they are used.
    Since unused values and blocks have ID == 0,
    we can efficiently find the last used value/block,
    to avoid zeroing everything.
    Bulk zeroing is efficient, but not efficient enough
    to obviate the need to avoid zeroing everything every time.
    As a happy side-effect, ssa.Func.Free is no longer necessary.
    
    DebugHashMatch and friends now belong in func.go.
    They have been left in place for clarity and review.
    I will move them in a subsequent CL.
    
    Passes toolstash -cmp. No compiler performance impact.
    No change in 'go test cmd/compile/internal/ssa' execution time.
    
    Change-Id: I2eb7af58da067ef6a36e815a6f386cfe8634d098
    Reviewed-on: https://go-review.googlesource.com/38167
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit ccaa8e3c6fbba08f238c5ccea437b7cac513685f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 15 23:01:31 2017 -0700

    cmd/compile: avoid calling unnecessary Sym format routine
    
    Minor cleanup only.
    
    No reason to go through String() when it is
    just as easy to do a direct string comparison.
    
    Eliminates a surprising number of allocations.
    
    name       old alloc/op    new alloc/op    delta
    Template      40.9MB ± 0%     40.9MB ± 0%    ~     (p=0.190 n=10+10)
    Unicode       30.3MB ± 0%     30.3MB ± 0%    ~     (p=0.218 n=10+10)
    GoTypes        116MB ± 0%      116MB ± 0%  -0.09%  (p=0.000 n=10+10)
    SSA            871MB ± 0%      869MB ± 0%  -0.14%  (p=0.000 n=10+9)
    Flate         26.2MB ± 0%     26.2MB ± 0%  -0.15%  (p=0.002 n=10+10)
    GoParser      32.5MB ± 0%     32.5MB ± 0%    ~     (p=0.165 n=10+10)
    Reflect       80.5MB ± 0%     80.4MB ± 0%  -0.12%  (p=0.003 n=9+10)
    Tar           27.3MB ± 0%     27.3MB ± 0%  -0.13%  (p=0.008 n=10+9)
    XML           43.1MB ± 0%     43.1MB ± 0%    ~     (p=0.218 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        402k ± 1%       400k ± 1%  -0.64%  (p=0.002 n=10+10)
    Unicode         322k ± 1%       321k ± 1%    ~     (p=0.075 n=10+10)
    GoTypes        1.19M ± 0%      1.18M ± 0%  -0.90%  (p=0.000 n=10+10)
    SSA            7.94M ± 0%      7.81M ± 0%  -1.66%  (p=0.000 n=10+9)
    Flate           246k ± 0%       242k ± 1%  -1.42%  (p=0.000 n=10+10)
    GoParser        325k ± 1%       323k ± 1%  -0.84%  (p=0.000 n=10+10)
    Reflect        1.02M ± 0%      1.01M ± 0%  -0.99%  (p=0.000 n=10+10)
    Tar             259k ± 0%       257k ± 1%  -0.72%  (p=0.009 n=10+10)
    XML             406k ± 1%       403k ± 1%  -0.69%  (p=0.001 n=10+10)
    
    Change-Id: Ia129a4cd272027d627e1f3b27e9f07f93e3aa27e
    Reviewed-on: https://go-review.googlesource.com/38230
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0cfb23135c67314dbb9fc2e78fd0f364b6882f25
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 15 22:55:21 2017 -0700

    cmd/compile: move hasdefer to Func
    
    Passes toolstash -cmp.
    
    Updates #15756
    
    Change-Id: Ia071dbbd7f2ee0f8433d8c37af4f7b588016244e
    Reviewed-on: https://go-review.googlesource.com/38231
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 604e4841d6bd0a2c207444481928053d1ad3b822
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 16 00:24:51 2017 -0700

    cmd/internal/obj/ppc64: remove stackbarrier function check
    
    Stack barriers were removed in CL 36620.
    
    Change-Id: If124d65a73a7b344a42be2a4b386a14d7a0a428b
    Reviewed-on: https://go-review.googlesource.com/38169
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>
    Reviewed-by: David Chase <drchase@google.com>

commit faeda66c60dbc080720b30d42acbf67c4541e053
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 16 17:29:14 2017 -0700

    go/types: better error for assignment count mismatches
    
    This matches the error message of cmd/compile (for assignments).
    
    Change-Id: I42a428f5d72f034e7b7e97b090a929e317e812af
    Reviewed-on: https://go-review.googlesource.com/38315
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 3c7a812485334eb57c0856e6b152aa3d50f9f0a0
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 16 17:07:26 2017 -0700

    cmd/compile: eliminate "assignment count mismatch" - not needed anymore
    
    See https://go-review.googlesource.com/#/c/38313/ for background.
    It turns out that only a few tests checked for this.
    
    The new error message is shorter and very clear.
    
    Change-Id: I8ab4ad59fb023c8b54806339adc23aefd7dc7b07
    Reviewed-on: https://go-review.googlesource.com/38314
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 73a44f0456389c60db8459981868b4c64874ec0b
Author: Jeremy Jackins <jeremyjackins@gmail.com>
Date:   Thu Mar 16 15:49:06 2017 -0700

    cmd/compile: further clarify assignment count mismatch error message
    
    This is an evolution of https://go-review.googlesource.com/33616, as discussed
    via email with Robert (gri):
    
    $ cat foobar.go
    package main
    
    func main() {
            a := "foo", "bar"
    }
    
    before:
    ./foobar.go:4:4: assignment count mismatch: want 1 values, got 2
    
    after:
    ./foobar.go:4:4: assignment count mismatch: cannot assign 2 values to 1 variables
    
    We could likely also eliminate the "assignment count mismatch" prefix now
    without losing any information, but that string is matched by a number of
    tests.
    
    Change-Id: Ie6fc8a7bbd0ebe841d53e66e5c2f49868decf761
    Reviewed-on: https://go-review.googlesource.com/38313
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 495b167919b1473e02843135f57d310cdd4f5789
Author: Keith Randall <keithr@alum.mit.edu>
Date:   Thu Mar 16 14:08:31 2017 -0700

    cmd/compile: intrinsics for math/bits.{Len,LeadingZeros}
    
    name              old time/op  new time/op  delta
    LeadingZeros-4    2.00ns ± 0%  1.34ns ± 1%  -33.02%  (p=0.000 n=8+10)
    LeadingZeros16-4  1.62ns ± 0%  1.57ns ± 0%   -3.09%  (p=0.001 n=8+9)
    LeadingZeros32-4  2.14ns ± 0%  1.48ns ± 0%  -30.84%  (p=0.002 n=8+10)
    LeadingZeros64-4  2.06ns ± 1%  1.33ns ± 0%  -35.08%  (p=0.000 n=8+8)
    
    8-bit args is a special case - the Go code is really fast because
    it is just a single table lookup.  So I've disabled that for now.
    Intrinsics were actually slower:
    LeadingZeros8-4   1.22ns ± 3%  1.58ns ± 1%  +29.56%  (p=0.000 n=10+10)
    
    Update #18616
    
    Change-Id: Ia9c289b9ba59c583ea64060470315fd637e814cf
    Reviewed-on: https://go-review.googlesource.com/38311
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 5f3e7aa76ddc190b199e207f2347ffffbb06b05a
Author: Steve Francia <spf@golang.org>
Date:   Thu Feb 9 15:33:13 2017 -0500

    doc: reorganize the contribution guidelines into a guide
    
    Updates #17802
    
    Change-Id: I65ea0f4cde973604c04051e7eb25d12e4facecd3
    Reviewed-on: https://go-review.googlesource.com/36626
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit bc8b9b23ca3f33e738d85c4734bd35dfe63e9ac4
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Thu Mar 9 12:29:45 2017 +0200

    strconv: optimize formatting for small decimal ints
    
    Avoid memory allocations by returning pre-calculated strings
    for decimal ints in the range 0..99.
    
    Benchmark results:
    
    name              old time/op    new time/op    delta
    FormatInt-4         2.45µs ± 1%    2.40µs ± 1%    -1.86%  (p=0.000 n=8+9)
    AppendInt-4         1.67µs ± 1%    1.65µs ± 0%    -0.92%  (p=0.000 n=10+10)
    FormatUint-4         676ns ± 3%     669ns ± 1%      ~     (p=0.146 n=10+10)
    AppendUint-4         467ns ± 2%     474ns ± 0%    +1.58%  (p=0.000 n=10+10)
    FormatIntSmall-4    29.6ns ± 2%     3.3ns ± 0%   -88.98%  (p=0.000 n=10+9)
    AppendIntSmall-4    16.0ns ± 1%     8.5ns ± 0%   -46.98%  (p=0.000 n=10+9)
    
    name              old alloc/op   new alloc/op   delta
    FormatInt-4           576B ± 0%      576B ± 0%      ~     (all equal)
    AppendInt-4          0.00B          0.00B           ~     (all equal)
    FormatUint-4          224B ± 0%      224B ± 0%      ~     (all equal)
    AppendUint-4         0.00B          0.00B           ~     (all equal)
    FormatIntSmall-4     2.00B ± 0%     0.00B       -100.00%  (p=0.000 n=10+10)
    AppendIntSmall-4     0.00B          0.00B           ~     (all equal)
    
    name              old allocs/op  new allocs/op  delta
    FormatInt-4           37.0 ± 0%      35.0 ± 0%    -5.41%  (p=0.000 n=10+10)
    AppendInt-4           0.00           0.00           ~     (all equal)
    FormatUint-4          6.00 ± 0%      6.00 ± 0%      ~     (all equal)
    AppendUint-4          0.00           0.00           ~     (all equal)
    FormatIntSmall-4      1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
    AppendIntSmall-4      0.00           0.00           ~     (all equal)
    
    Fixes #19445
    
    Change-Id: Ib1f8922f2e0b13743c847ee9e703d1dab77f705c
    Reviewed-on: https://go-review.googlesource.com/37963
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit dd9892e31b24a496865cea4db6cdca6d73142895
Author: Keith Randall <keithr@alum.mit.edu>
Date:   Wed Mar 15 21:28:29 2017 -0700

    cmd/compile: intrinsify math/bits.ReverseBytes
    
    Update #18616
    
    Change-Id: I0c2d643cbbeb131b4c9b12194697afa4af48e1d2
    Reviewed-on: https://go-review.googlesource.com/38166
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 793e4ec3dd440ddb68ec4d154f5711ef76d32b60
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Mar 16 15:26:13 2017 -0400

    cmd/compile: fix MIPS Zero lower rule
    
    A copy-paste error in CL 38150. Fix build.
    
    Change-Id: Ib2afc83564ebe7dab934d45522803e1a191dea18
    Reviewed-on: https://go-review.googlesource.com/38292
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit f37ee0f33ba461dbaa58645c3ebdf148a43911a5
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Mar 16 10:43:54 2017 -0700

    cmd/compile/internal/syntax: track column position at function end
    
    Fixes #19576.
    
    Change-Id: I11034fb08e989f6eb7d54bde873b92804223598d
    Reviewed-on: https://go-review.googlesource.com/38291
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c8f38b339875e020f50bedbdb8cfd8a7a1ef12b1
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Mar 13 21:51:08 2017 -0400

    cmd/compile: use type information in Aux for Store size
    
    Remove size AuxInt in Store, and alignment in Move/Zero. We still
    pass size AuxInt to Move/Zero, as it is used for partial Move/Zero
    lowering (e.g. cmd/compile/internal/ssa/gen/386.rules:288).
    SizeAndAlign is gone.
    
    Passes "toolstash -cmp" on std.
    
    Change-Id: I1ca34652b65dd30de886940e789fcf41d521475d
    Reviewed-on: https://go-review.googlesource.com/38150
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d75925d6bad2576bd80c442ff3b23ba22ffb2c68
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Feb 13 21:27:53 2017 -0500

    cmd/compile: add a test for writebarrier pass with single-block loop
    
    The old writebarrier implementation fails to handle single-block
    loop where a memory Phi value depends on the write barrier store
    in the same block. The new implementation (CL 36834) doesn't have
    this problem. Add a test to ensure it.
    
    Fix #19067.
    
    Change-Id: Iab13c6817edc12be8a048d18699b4450fa7ed712
    Reviewed-on: https://go-review.googlesource.com/36940
    Reviewed-by: David Chase <drchase@google.com>

commit 1b85300602f8530e505ad5b8b033a15f5521d1a7
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Feb 10 10:15:10 2017 -0500

    cmd/compile: clean up SSA-building code
    
    Now that the write barrier insertion is moved to SSA, the SSA
    building code can be simplified.
    
    Updates #17583.
    
    Change-Id: I5cacc034b11aa90b0abe6f8dd97e4e3994e2bc25
    Reviewed-on: https://go-review.googlesource.com/36840
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 9ebf3d5100a52b2c0ebcbf9754c02d1edf7a035f
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Feb 5 23:43:31 2017 -0500

    cmd/compile: move write barrier insertion to SSA
    
    When the compiler insert write barriers, the frontend makes
    conservative decisions at an early stage. This sometimes have
    false positives because of the lack of information, for example,
    writes on stack. SSA's writebarrier pass identifies writes on
    stack and eliminates write barriers for them.
    
    This CL moves write barrier insertion into SSA. The frontend no
    longer makes decisions about write barriers, and simply does
    normal assignments and emits normal Store ops when building SSA.
    SSA writebarrier pass inserts write barrier for Stores when needed.
    There, it has better information about the store because Phi and
    Copy propagation are done at that time.
    
    This CL only changes StoreWB to Store in gc/ssa.go. A followup CL
    simplifies SSA building code.
    
    Updates #17583.
    
    Change-Id: I4592d9bc0067503befc169c50b4e6f4765673bec
    Reviewed-on: https://go-review.googlesource.com/36839
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 211c8c9f1a1d11a6f5c42e85ec78cd06a69fbe0c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Feb 9 09:46:44 2017 -0500

    cmd/compile: pass types on SSA Store/Move/Zero ops
    
    For SSA Store/Move/Zero ops, attach the type of the value being
    stored to the op as the Aux field. This type will be used for
    write barrier insertion (in a followup CL). Since SSA passes
    do not accurately propagate types of values (because of type
    casting), we can't simply use type of the store's arguments
    for write barrier insertion.
    
    Passes "toolstash -cmp" on std.
    
    Updates #17583.
    
    Change-Id: I051d5e5c482931640d1d7d879b2a6bb91f2e0056
    Reviewed-on: https://go-review.googlesource.com/36838
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 77b09b8b8dc9d27f0c583f4eb94a563e68fe0af6
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Tue Mar 7 16:10:59 2017 +0000

    runtime: remove unused g parameter
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: I20145440ff1bcd27fcf15a740354c52f313e536c
    Reviewed-on: https://go-review.googlesource.com/37894
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit d60166d5eea5084e0957e9028237cc87ecadbf7d
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Mon Feb 27 21:32:29 2017 -0300

    runtime: improve IndexByte for ppc64x
    
    This change adds a better implementation of IndexByte for ppc64x.
    
    Improvement for bytes·IndexByte:
    
    benchmark                             old ns/op     new ns/op     delta
    BenchmarkIndexByte/10-16              12.5          8.48          -32.16%
    BenchmarkIndexByte/32-16              34.4          9.85          -71.37%
    BenchmarkIndexByte/4K-16              3089          217           -92.98%
    BenchmarkIndexByte/4M-16              3154810       207051        -93.44%
    BenchmarkIndexByte/64M-16             50564811      5579093       -88.97%
    
    benchmark                             old MB/s     new MB/s     speedup
    BenchmarkIndexByte/10-16              800.41       1179.64      1.47x
    BenchmarkIndexByte/32-16              930.60       3249.10      3.49x
    BenchmarkIndexByte/4K-16              1325.71      18832.53     14.21x
    BenchmarkIndexByte/4M-16              1329.49      20257.29     15.24x
    BenchmarkIndexByte/64M-16             1327.19      12028.63     9.06x
    
    Improvement for strings·IndexByte:
    
    benchmark                             old ns/op     new ns/op     delta
    BenchmarkIndexByte-16                 25.9          7.69          -70.31%
    
    Fixes #19030
    
    Change-Id: Ifb82bbb3d643ec44b98eaa2d08a07f47e5c2fd11
    Reviewed-on: https://go-review.googlesource.com/37670
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit d5dc4905191c3f7cfeb52e93331d10ebd33301f5
Author: Keith Randall <khr@golang.org>
Date:   Tue Mar 14 13:25:12 2017 -0700

    cmd/compile: intrinsics for math/bits.TrailingZerosX
    
    Implement math/bits.TrailingZerosX using intrinsics.
    
    Generally reorganize the intrinsic spec a bit.
    The instrinsics data structure is now built at init time.
    This will make doing the other functions in math/bits easier.
    
    Update sys.CtzX to return int instead of uint{64,32} so it
    matches math/bits.TrailingZerosX.
    
    Improve the intrinsics a bit for amd64.  We don't need the CMOV
    for <64 bit versions.
    
    Update #18616
    
    Change-Id: Ic1c5339c943f961d830ae56f12674d7b29d4ff39
    Reviewed-on: https://go-review.googlesource.com/38155
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 16200c73331a679b43efc4699b5806c64a556f09
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sat Feb 25 23:50:56 2017 +0100

    runtime: make complex division c99 compatible
    
    - changes tests to check that the real and imaginary part of the go complex
      division result is equal to the result gcc produces for c99
    - changes complex division code to satisfy new complex division test
    - adds float functions isNan, isFinite, isInf, abs and copysign
      in the runtime package
    
    Fixes #14644.
    
    name                   old time/op  new time/op  delta
    Complex128DivNormal-4  21.8ns ± 6%  13.9ns ± 6%  -36.37%  (p=0.000 n=20+20)
    Complex128DivNisNaN-4  14.1ns ± 1%  15.0ns ± 1%   +5.86%  (p=0.000 n=20+19)
    Complex128DivDisNaN-4  12.5ns ± 1%  16.7ns ± 1%  +33.79%  (p=0.000 n=19+20)
    Complex128DivNisInf-4  10.1ns ± 1%  13.0ns ± 1%  +28.25%  (p=0.000 n=20+19)
    Complex128DivDisInf-4  11.0ns ± 1%  20.9ns ± 1%  +90.69%  (p=0.000 n=16+19)
    ComplexAlgMap-4        86.7ns ± 1%  86.8ns ± 2%     ~     (p=0.804 n=20+20)
    
    Change-Id: I261f3b4a81f6cc858bc7ff48f6fd1b39c300abf0
    Reviewed-on: https://go-review.googlesource.com/37441
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 4b8f41daa63154949104d27d70acc3857a0b4b0b
Author: Austin Clements <austin@google.com>
Date:   Fri Mar 10 10:59:39 2017 -0500

    runtime: print user stack on other threads during GOTRACBEACK=crash
    
    Currently, when printing tracebacks of other threads during
    GOTRACEBACK=crash, if the thread is on the system stack we print only
    the header for the user goroutine and fail to print its stack. This
    happens because we passed the g0 to traceback instead of curg. The g0
    never has anything set in its gobuf, so traceback doesn't print
    anything.
    
    Fix this by passing _g_.m.curg to traceback instead of the g0.
    
    Fixes #19494.
    
    Change-Id: Idfabf94d6a725e9cdf94a3923dead6455ef3b217
    Reviewed-on: https://go-review.googlesource.com/38012
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f2e87158f06902daa44d34b4846611f8520e4491
Author: Austin Clements <austin@google.com>
Date:   Wed Mar 15 14:48:23 2017 -0400

    runtime: make GOTRACEBACK=crash crash promptly in cgo binaries
    
    GOTRACEBACK=crash works by bouncing a SIGQUIT around the process
    sched.mcount times. However, sched.mcount includes the extra Ms
    allocated by oneNewExtraM for cgo callbacks. Hence, if there are any
    extra Ms that don't have real OS threads, we'll try to send SIGQUIT
    more times than there are threads to catch it. Since nothing will
    catch these extra signals, we'll fall back to blocking for five
    seconds before aborting the process.
    
    Avoid this five second delay by subtracting out the number of extra Ms
    when sending SIGQUITs.
    
    Of course, in a cgo binary, it's still possible for the SIGQUIT to go
    to a cgo thread and cause some other failure mode. This does not fix
    that.
    
    Change-Id: I4fbf3c52dd721812796c4c1dcb2ab4cb7026d965
    Reviewed-on: https://go-review.googlesource.com/38182
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c03e75e53915af9905bc261c66b5276de042ea1c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 14 11:05:03 2017 -0700

    cmd/compile: check labels and gotos before building SSA
    
    This CL introduces yet another compiler pass,
    which checks for correct control flow constructs
    prior to converting from AST to SSA form.
    
    It cannot be integrated with walk, since walk rewrites
    switch and select statements on the fly.
    
    To reduce code duplication, this CL also does some
    minor refactoring.
    
    With this pass in place, the AST to SSA converter
    can now stop generating SSA for any known-dead code.
    This minor savings pays for the minor cost of the new pass.
    
    Performance is almost a wash:
    
    name       old time/op     new time/op     delta
    Template       206ms ± 4%      205ms ± 4%   ~     (p=0.108 n=43+43)
    Unicode       84.0ms ± 4%     84.0ms ± 4%   ~     (p=0.979 n=43+43)
    GoTypes        550ms ± 3%      553ms ± 3%   ~     (p=0.065 n=40+41)
    Compiler       2.57s ± 4%      2.58s ± 2%   ~     (p=0.103 n=44+41)
    SSA            3.94s ± 3%      3.93s ± 2%   ~     (p=0.833 n=44+42)
    Flate          126ms ± 6%      125ms ± 4%   ~     (p=0.941 n=43+39)
    GoParser       147ms ± 4%      148ms ± 3%   ~     (p=0.164 n=42+39)
    Reflect        359ms ± 3%      357ms ± 5%   ~     (p=0.241 n=43+44)
    Tar            106ms ± 5%      106ms ± 7%   ~     (p=0.853 n=40+43)
    XML            202ms ± 3%      203ms ± 3%   ~     (p=0.488 n=42+41)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        240M ± 4%       239M ± 4%   ~     (p=0.844 n=42+43)
    Unicode         107M ± 5%       107M ± 4%   ~     (p=0.332 n=40+43)
    GoTypes         735M ± 3%       731M ± 4%   ~     (p=0.141 n=43+44)
    Compiler       3.51G ± 3%      3.52G ± 3%   ~     (p=0.208 n=42+43)
    SSA            5.72G ± 4%      5.72G ± 3%   ~     (p=0.928 n=44+42)
    Flate           151M ± 7%       150M ± 8%   ~     (p=0.662 n=44+43)
    GoParser        181M ± 5%       181M ± 4%   ~     (p=0.379 n=41+44)
    Reflect         447M ± 4%       445M ± 4%   ~     (p=0.344 n=43+43)
    Tar             125M ± 7%       124M ± 6%   ~     (p=0.353 n=43+43)
    XML             248M ± 4%       250M ± 6%   ~     (p=0.158 n=44+44)
    
    name       old alloc/op    new alloc/op    delta
    Template      40.3MB ± 0%     40.2MB ± 0%  -0.27%  (p=0.000 n=10+10)
    Unicode       30.3MB ± 0%     30.2MB ± 0%  -0.10%  (p=0.015 n=10+10)
    GoTypes        114MB ± 0%      114MB ± 0%  -0.06%  (p=0.000 n=7+9)
    Compiler       480MB ± 0%      481MB ± 0%  +0.07%  (p=0.000 n=10+10)
    SSA            864MB ± 0%      862MB ± 0%  -0.25%  (p=0.000 n=9+10)
    Flate         25.9MB ± 0%     25.9MB ± 0%    ~     (p=0.123 n=10+10)
    GoParser      32.1MB ± 0%     32.1MB ± 0%    ~     (p=0.631 n=10+10)
    Reflect       79.9MB ± 0%     79.6MB ± 0%  -0.39%  (p=0.000 n=10+9)
    Tar           27.1MB ± 0%     27.0MB ± 0%  -0.18%  (p=0.003 n=10+10)
    XML           42.6MB ± 0%     42.6MB ± 0%    ~     (p=0.143 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        401k ± 0%       401k ± 1%    ~     (p=0.353 n=10+10)
    Unicode         322k ± 0%       322k ± 0%    ~     (p=0.739 n=10+10)
    GoTypes        1.18M ± 0%      1.18M ± 0%  +0.25%  (p=0.001 n=7+8)
    Compiler       4.51M ± 0%      4.53M ± 0%  +0.37%  (p=0.000 n=10+10)
    SSA            7.91M ± 0%      7.93M ± 0%  +0.20%  (p=0.000 n=9+10)
    Flate           244k ± 0%       245k ± 0%    ~     (p=0.123 n=10+10)
    GoParser        323k ± 1%       324k ± 1%  +0.40%  (p=0.035 n=10+10)
    Reflect        1.01M ± 0%      1.02M ± 0%  +0.37%  (p=0.000 n=10+9)
    Tar             258k ± 1%       258k ± 1%    ~     (p=0.661 n=10+9)
    XML             403k ± 0%       405k ± 0%  +0.47%  (p=0.004 n=10+10)
    
    Updates #15756
    Updates #19250
    
    Change-Id: I647bfbb745c35630447eb79dfcaa994b490ce942
    Reviewed-on: https://go-review.googlesource.com/38159
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 604455a46c3dac17422d9ca941848dbf7ae116b9
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 15 11:27:26 2017 -0700

    cmd/compile: ensure TESTQconst AuxInt is in range
    
    Fixes #19555
    
    Change-Id: I7aa0551a90f6bb630c0ba721f3525a8a9cf793fd
    Reviewed-on: https://go-review.googlesource.com/38164
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d0a045daaf022515ddf518cdded04f672210f8c4
Author: Bryan C. Mills <bcmills@google.com>
Date:   Fri Feb 10 12:46:14 2017 -0500

    archive/zip: parallelize benchmarks
    
    Add subbenchmarks for BenchmarkZip64Test with different sizes to tease
    apart construction costs vs. steady-state throughput.
    
    Results remain comparable with the non-parallel version with -cpu=1:
    
    benchmark                           old ns/op     new ns/op     delta
    BenchmarkCompressedZipGarbage       26832835      27506953      +2.51%
    BenchmarkCompressedZipGarbage-6     27172377      4321534       -84.10%
    BenchmarkZip64Test                  196758732     197765510     +0.51%
    BenchmarkZip64Test-6                193850605     192625458     -0.63%
    
    benchmark                           old allocs     new allocs     delta
    BenchmarkCompressedZipGarbage       44             44             +0.00%
    BenchmarkCompressedZipGarbage-6     44             44             +0.00%
    
    benchmark                           old bytes     new bytes     delta
    BenchmarkCompressedZipGarbage       5592          5664          +1.29%
    BenchmarkCompressedZipGarbage-6     5592          21946         +292.45%
    
    updates #18177
    
    Change-Id: Icfa359d9b1a8df5e085dacc07d2b9221b284764c
    Reviewed-on: https://go-review.googlesource.com/36719
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 15b37655bc0b786d3cd6e41553263963fbebd642
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Mar 13 08:24:46 2017 -0400

    cmd/link: on PPC64, put plt stubs at beginning of Textp
    
    Put call stubs at the beginning (instead of the end). So the
    trampoline pass knows the addresses of the stubs, and it can
    insert trampolines when necessary.
    
    Fixes #19425.
    
    Change-Id: I1e06529ef837a6130df58917315610d45a6819ca
    Reviewed-on: https://go-review.googlesource.com/38131
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit 43afcb5c969112332e46b4cb07bf2fd6497c3919
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 14 16:44:48 2017 -0700

    cmd/compile: define roles for ssa.Func, ssa.Config, and ssa.Cache
    
    The line between ssa.Func and ssa.Config has blurred.
    Concurrent compilation in the backend will require more precision.
    This CL lays out an (aspirational) organization.
    The implementation will come in follow-up CLs,
    once the organization is settled.
    
    ssa.Config holds basic compiler configuration,
    mostly arch-specific information.
    It is configured once, early on, and is readonly,
    so it is safe for concurrent use.
    
    ssa.Func is a single-shot object used for
    compiling a single Func. It is not concurrency-safe
    and not re-usable.
    
    ssa.Cache is a multi-use object used to avoid
    expensive allocations during compilation.
    Each ssa.Func is given an ssa.Cache to use.
    ssa.Cache is not concurrency-safe.
    
    Change-Id: Id02809b6f3541541cac6c27bbb598834888ce1cc
    Reviewed-on: https://go-review.googlesource.com/38160
    Reviewed-by: Keith Randall <khr@golang.org>

commit 886e9e6065588d6c40567f3323883ded7ad3d946
Author: David Chase <drchase@google.com>
Date:   Tue Mar 7 14:45:46 2017 -0500

    cmd/compile: put spills in better places
    
    Previously we always issued a spill right after the op
    that was being spilled.  This CL pushes spills father away
    from the generator, hopefully pushing them into unlikely branches.
    For example:
    
      x = ...
      if unlikely {
        call ...
      }
      ... use x ...
    
    Used to compile to
    
      x = ...
      spill x
      if unlikely {
        call ...
        restore x
      }
    
    It now compiles to
    
      x = ...
      if unlikely {
        spill x
        call ...
        restore x
      }
    
    This is particularly useful for code which appends, as the only
    call is an unlikely call to growslice.  It also helps for the
    spills needed around write barrier calls.
    
    The basic algorithm is walk down the dominator tree following a
    path where the block still dominates all of the restores.  We're
    looking for a block that:
     1) dominates all restores
     2) has the value being spilled in a register
     3) has a loop depth no deeper than the value being spilled
    
    The walking-down code is iterative.  I was forced to limit it to
    searching 100 blocks so it doesn't become O(n^2).  Maybe one day
    we'll find a better way.
    
    I had to delete most of David's code which pushed spills out of loops.
    I suspect this CL subsumes most of the cases that his code handled.
    
    Generally positive performance improvements, but hard to tell for sure
    with all the noise.  (compilebench times are unchanged.)
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.91s ±15%     2.80s ±12%    ~     (p=0.063 n=10+10)
    Fannkuch11-12                3.47s ± 0%     3.30s ± 4%  -4.91%   (p=0.000 n=9+10)
    FmtFprintfEmpty-12          48.0ns ± 1%    47.4ns ± 1%  -1.32%    (p=0.002 n=9+9)
    FmtFprintfString-12         85.6ns ±11%    79.4ns ± 3%  -7.27%  (p=0.005 n=10+10)
    FmtFprintfInt-12            91.8ns ±10%    85.9ns ± 4%    ~      (p=0.203 n=10+9)
    FmtFprintfIntInt-12          135ns ±13%     127ns ± 1%  -5.72%   (p=0.025 n=10+9)
    FmtFprintfPrefixedInt-12     167ns ± 1%     168ns ± 2%    ~      (p=0.580 n=9+10)
    FmtFprintfFloat-12           249ns ±11%     230ns ± 1%  -7.32%  (p=0.000 n=10+10)
    FmtManyArgs-12               504ns ± 7%     506ns ± 1%    ~       (p=0.198 n=9+9)
    GobDecode-12                6.95ms ± 1%    7.04ms ± 1%  +1.37%  (p=0.001 n=10+10)
    GobEncode-12                6.32ms ±13%    6.04ms ± 1%    ~     (p=0.063 n=10+10)
    Gzip-12                      233ms ± 1%     235ms ± 0%  +1.01%   (p=0.000 n=10+9)
    Gunzip-12                   40.1ms ± 1%    39.6ms ± 0%  -1.12%   (p=0.000 n=10+8)
    HTTPClientServer-12          227µs ± 9%     221µs ± 5%    ~       (p=0.114 n=9+8)
    JSONEncode-12               16.1ms ± 2%    15.8ms ± 1%  -2.09%    (p=0.002 n=9+8)
    JSONDecode-12               61.8ms ±11%    57.9ms ± 1%  -6.30%   (p=0.000 n=10+9)
    Mandelbrot200-12            4.30ms ± 3%    4.28ms ± 1%    ~      (p=0.203 n=10+8)
    GoParse-12                  3.18ms ± 2%    3.18ms ± 2%    ~     (p=0.579 n=10+10)
    RegexpMatchEasy0_32-12      76.7ns ± 1%    77.5ns ± 1%  +0.92%    (p=0.002 n=9+8)
    RegexpMatchEasy0_1K-12       239ns ± 3%     239ns ± 1%    ~     (p=0.204 n=10+10)
    RegexpMatchEasy1_32-12      71.4ns ± 1%    70.6ns ± 0%  -1.15%   (p=0.000 n=10+9)
    RegexpMatchEasy1_1K-12       383ns ± 2%     390ns ±10%    ~       (p=0.181 n=8+9)
    RegexpMatchMedium_32-12      114ns ± 0%     113ns ± 1%  -0.88%    (p=0.000 n=9+8)
    RegexpMatchMedium_1K-12     36.3µs ± 1%    36.8µs ± 1%  +1.59%   (p=0.000 n=10+8)
    RegexpMatchHard_32-12       1.90µs ± 1%    1.90µs ± 1%    ~     (p=0.341 n=10+10)
    RegexpMatchHard_1K-12       59.4µs ±11%    57.8µs ± 1%    ~      (p=0.968 n=10+9)
    Revcomp-12                   461ms ± 1%     462ms ± 1%    ~       (p=1.000 n=9+9)
    Template-12                 67.5ms ± 1%    66.3ms ± 1%  -1.77%   (p=0.000 n=10+8)
    TimeParse-12                 314ns ± 3%     309ns ± 0%  -1.56%    (p=0.000 n=9+8)
    TimeFormat-12                340ns ± 2%     331ns ± 1%  -2.79%  (p=0.000 n=10+10)
    
    The go binary is 0.2% larger.  Not really sure why the size
    would change.
    
    Change-Id: Ia5116e53a3aeb025ef350ffc51c14ae5cc17871c
    Reviewed-on: https://go-review.googlesource.com/34822
    Reviewed-by: David Chase <drchase@google.com>

commit 710f4d3e7e0901f8fa2f04c31c0d28c603903ff2
Author: Philip Hofer <phofer@umich.edu>
Date:   Tue Mar 14 14:00:38 2017 -0700

    cmd/compile/internal/gc: mark generated wrappers as DUPOK
    
    Interface wrapper functions now get compiled eagerly in some cases.
    Consequently, they may be present in multiple translation units.
    Mark them as DUPOK, just like closures.
    
    Fixes #19548
    Fixes #19550
    
    Change-Id: Ibe74adb5a62dbf6447db37fde22dcbb3479969ef
    Reviewed-on: https://go-review.googlesource.com/38156
    Reviewed-by: David Chase <drchase@google.com>

commit 8a44c8efaefbbda6dd7ab4ee9a5e449fefbf5e1a
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Mar 14 18:21:23 2017 -0400

    cmd/compile: don't spill rematerializeable value when resolving merge edges
    
    Fixes #19515.
    
    Change-Id: I4bcce152cef52d00fbb5ab4daf72a6e742bae27c
    Reviewed-on: https://go-review.googlesource.com/38158
    Reviewed-by: Keith Randall <khr@golang.org>

commit 2d78538c12ed63cc931b552c9c98b7c93b91b21d
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Mar 9 12:15:41 2017 -0800

    cmd/compile: refactor liveness analysis for moving to SSA
    
    In the SSA CFG, TEXT, RET, and JMP instructions correspond to Blocks,
    not Values. Rework liveness analysis so that progeffects only cares
    about Progs that result from Values, and handle Blocks separately.
    
    Passes toolstash-check -all.
    
    Change-Id: Ic23719c75b0421fdb51382a08dac18c3ba042b32
    Reviewed-on: https://go-review.googlesource.com/38085
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit a9824cd47c1927ffc4eca040d60aecaa26130329
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Fri Mar 10 17:15:46 2017 +0000

    *.bash: always use the same string equality operator
    
    POSIX Shell only supports = to compare variables inside '[' tests. But
    this is Bash, where == is an alias for =. In practice they're the same,
    but the current form is inconsisnent and breaks POSIX for no good
    reason.
    
    Change-Id: I38fa7a5a90658dc51acc2acd143049e510424ed8
    Reviewed-on: https://go-review.googlesource.com/38031
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 723ba180b39969463f3644bac6422f0f81da26ea
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Nov 7 16:14:32 2016 -0800

    cmd/compile: eliminate fmtmode and fmtpkgpfx globals
    
    The fmtmode and fmtpkgpfx globals stand in the
    way of making the compiler more concurrent (#15756).
    This CL removes them.
    
    The natural way to eliminate a global is to explicitly
    thread it as a parameter through all function calls.
    However, most of the functions in gc/fmt.go
    get called indirectly, by way of fmt format strings,
    so there's nowhere natural to add a parameter.
    
    Since there are only a few fmtmode modes,
    use named types to distinguish between modes.
    For example, fmtNodeErr, fmtNodeDbg, and fmtNodeTypeId
    are all gc.Node, but they print in different modes.
    Varying the type allows us to thread mode through fmt.
    Handle fmtpkgpfx by converting it to a printing mode,
    FTypeIdName, and using the same type-based approach.
    
    To avoid a loss of readability and danger of bugs
    from introducing conversions at all call sites,
    instead add a helper that systematically modifies the args.
    
    The only remaining gc/fmt.go global is dumpdepth.
    Since that is used for debugging only,
    it that can be handled with a global mutex,
    or some similarly basic, if inefficient, protection.
    
    Passes toolstash -cmp. No compiler performance impact.
    
    For future reference, other options for threading state
    that were considered and rejected:
    
    * Wrapping values in structs, such as:
    
      type fmtNode struct {
            n *Node
            mode fmtMode
      }
    
      This reduces the proliferation of types, and supports
      easily adding extra local parameters.
      However, putting such a struct into an interface{} allocates.
      This is unacceptable in this particular area of code.
    
    * Passing state via precision, such as:
    
      fmt.Fprintf("%*v", mode, n)
    
      where mode is the state encoded as an integer.
      This avoids extra allocations, but it is out of keeping
      with the intended semantics of precision, and is less readable.
    
    * Modify the fmt package to support setting/getting context
      via fmt.State. Unavailable due to Go 1 compatibility,
      and probably the wrong solution anyway.
    
    * Give up on package fmt. This would be a huge readability
      regression and cause high code churn.
    
    * Attempt a de-novo rewrite that circumvents these problems.
      Too high a risk of bugs, with insufficient reward for the effort,
      particularly since long term plans call for elimination
      of gc.Node.
    
    
    Change-Id: Iea2440d5a34a938e64273707de27e3a897cb41d1
    Reviewed-on: https://go-review.googlesource.com/38147
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 88cf932e98db8fe014399e05320128d76183f3fa
Author: Michael Stapelberg <stapelberg@google.com>
Date:   Tue Nov 29 02:45:11 2016 -0800

    cmd/compile: improve assignment count mismatch error message
    
    Given the following test cases:
    
        $ cat left_too_many.go
        package main
    
        func main() {
            a, err := make([]int, 1)
        }
    
        $ cat right_too_many.go
        package main
    
        func main() {
            a := "foo", "bar"
        }
    
    Before this change, the error messages are:
    
        ./left_too_many.go:4: assignment count mismatch: 2 = 1
    
        ./right_too_many.go:4: assignment count mismatch: 1 = 2
    
    After this change, the error messages are:
    
        ./left_too_many.go:4: assignment count mismatch: want 2 values, got 1
    
        ./right_too_many.go:4: assignment count mismatch: want 1 values, got 2
    
    Change-Id: I9ad346f122406bc9a785bf690ed7b3de76a422da
    Reviewed-on: https://go-review.googlesource.com/33616
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 91d08e3bca1c12f45d105ada3f4f46a73375dac9
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 14 12:18:10 2017 -0700

    cmd/compile/internal/ssa: remove unused OpFunc
    
    Change-Id: I0f7eec2e0c15a355422d5ae7289508a5bd33b971
    Reviewed-on: https://go-review.googlesource.com/38171
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 295307ae78f8dd463a2ab8d85a1592ca76619d36
Author: philhofer <phofer@umich.edu>
Date:   Mon Mar 13 15:03:17 2017 -0700

    cmd/compile: de-virtualize interface calls
    
    With this change, code like
    
        h := sha1.New()
        h.Write(buf)
        sum := h.Sum()
    
    gets compiled into static calls rather than
    interface calls, because the compiler is able
    to prove that 'h' is really a *sha1.digest.
    
    The InterCall re-write rule hits a few dozen times
    during make.bash, and hundreds of times during all.bash.
    
    The most common pattern identified by the compiler
    is a constructor like
    
        func New() Interface { return &impl{...} }
    
    where the constructor gets inlined into the caller,
    and the result is used immediately. Examples include
    {sha1,md5,crc32,crc64,...}.New, base64.NewEncoder,
    base64.NewDecoder, errors.New, net.Pipe, and so on.
    
    Some existing benchmarks that change on darwin/amd64:
    
    Crc64/ISO4KB-8        2.67µs ± 1%    2.66µs ± 0%  -0.36%  (p=0.015 n=10+10)
    Crc64/ISO1KB-8         694ns ± 0%     690ns ± 1%  -0.59%  (p=0.001 n=10+10)
    Adler32KB-8            473ns ± 1%     471ns ± 0%  -0.39%  (p=0.010 n=10+9)
    
    On architectures like amd64, the reduction in code size
    appears to contribute more to benchmark improvements than just
    removing the indirect call, since that branch gets predicted
    accurately when called in a loop.
    
    Updates #19361
    
    Change-Id: I57d4dc21ef40a05ec0fbd55a9bb0eb74cdc67a3d
    Reviewed-on: https://go-review.googlesource.com/38139
    Run-TryBot: Philip Hofer <phofer@umich.edu>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 691755304ce49887f65f8f84e18cda3814b46e5c
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Mar 9 14:46:43 2017 -0800

    cmd/compile/internal/ssa: populate SymEffects for SSA Ops
    
    Changes to ${GOARCH}Ops.go files were mechanically produced using
    github.com/mdempsky/ssa-symops, a one-off tool that inserts
    "SymEffect: X" elements by pattern matching against the Op names.
    
    Change-Id: Ibf3e481ffd588647f2a31662d72114b740ccbfcf
    Reviewed-on: https://go-review.googlesource.com/38084
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 1cdf4bf33f57ca7910ab4ee1121ea7f05a6adcd1
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Mar 9 14:45:37 2017 -0800

    cmd/compile/internal/ssa: add SymEffect attribute to SSA Ops
    
    To replace the progeffects tables for liveness analysis.
    
    Change-Id: Idc4b990665cb0a9aa300d62cdf8ad12e51c5b991
    Reviewed-on: https://go-review.googlesource.com/38083
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 5f5d882f5cbd7f39d8bbbd5481a0884c0a042032
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 14 09:29:23 2017 -0700

    cmd/compile: catch bad pragma combination earlier
    
    Bad pragmas should never make it to the backend.
    I've confirmed manually that the error position is unchanged.
    
    Updates #15756
    Updates #19250
    
    Change-Id: If14f7ce868334f809e337edc270a49680b26f48e
    Reviewed-on: https://go-review.googlesource.com/38152
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 0d87f6a62e497fabd473a9a9847c355a42d64f16
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Mar 14 17:43:50 2017 +0000

    net/http: deflake TestServerTimeouts
    
    Retry the test several times with increasingly long timeouts.
    
    Fixes #19538 (hopefully)
    
    Change-Id: Ia3bf2b63b4298a6ee1e4082e14d9bfd5922c293a
    Reviewed-on: https://go-review.googlesource.com/38154
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 50520f1543194b591dd517f6c664aa31c29e3b3f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 14 09:46:45 2017 -0700

    cmd/compile: use Fatalf for more internal errors
    
    There were a surprising number of places
    in the tree that used yyerror for failed internal
    consistency checks. Switch them to Fatalf.
    
    Updates #15756
    Updates #19250
    
    Change-Id: Ie4278148185795a28ff3c27dacffc211cda5bbdd
    Reviewed-on: https://go-review.googlesource.com/38153
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 57107f300a53c0e52bc0a1ff03d52faa628c11e5
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Mar 14 09:12:55 2017 -0700

    cmd/compile: use local fn variable in compile
    
    fn == Curfn in this context. Prefer the local variable.
    
    Passes toolstash -cmp. Updates #15756.
    
    Change-Id: I75b589c682d0c1b524cac2bbf2bba368a6027b06
    Reviewed-on: https://go-review.googlesource.com/38151
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ccae744f8017971b59303f62702ecd5c269891ec
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 14 08:18:11 2017 -0700

    cmd/compile/internal/gc: better loop var names in esc.go
    
    Used gorename.
    
    Change-Id: Ib33305dc95876ec18e2473ad2999788a32eb21c0
    Reviewed-on: https://go-review.googlesource.com/38146
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 7a9aa06902540803773f22cd868bc5809dddad02
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Nov 7 15:50:16 2016 -0800

    cmd/compile: remove FmtFlag save and restore
    
    It is unnecessary.
    
    Passes toolstash -cmp.
    
    Change-Id: I7c03523b6110c3d9bd5ba2b37d9a1e17a7ae570e
    Reviewed-on: https://go-review.googlesource.com/38145
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8ee37c6a66fd71fa0204204719b0489085f85ca0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Nov 7 11:52:21 2016 -0800

    cmd/compile: eliminate format string FmtUnsigned support
    
    Passes toolstash -cmp.
    
    Change-Id: I678fc40c0f2a6e9a434bcdd4ea17bb7f319a6063
    Reviewed-on: https://go-review.googlesource.com/38144
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 84b9329bd9085a47325818005b6162af265ddfa4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Nov 6 11:08:08 2016 -0800

    cmd/compile: add Type.LongString and Type.ShortString
    
    Reduces duplication and centralizes documentation.
    Moves all uses of FmtUnsigned and tconv inside fmt.go.
    
    Passes toolstash -cmp.
    
    Change-Id: If6d906e7e839de05f36423523a3a1d596e29807d
    Reviewed-on: https://go-review.googlesource.com/38141
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 6088fab92015d4901e515a6991db874889508eba
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Nov 6 11:09:09 2016 -0800

    cmd/compile: rename Nconv to nconv
    
    For consistency.
    
    Change-Id: Ic687fea95f7a4a3be576945af3e9c97086309b07
    Reviewed-on: https://go-review.googlesource.com/38142
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0fd6df352205e2781ab34349527cf677b6c1283f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Nov 7 10:28:46 2016 -0800

    cmd/compile: simplify printing of constant bools
    
    Change-Id: I9339e83e39075826bf5819e55804a94208fe84ae
    Reviewed-on: https://go-review.googlesource.com/38140
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ec091b6af2ad6f33f3b36c39171d7ef65b12668b
Author: Hugues Bruant <hugues.bruant@gmail.com>
Date:   Sun Mar 12 14:47:59 2017 -0700

    runtime: add mapassign_fast*
    
    Add benchmarks for map assignment with int32/int64/string key
    
    Benchmark results on darwin/amd64
    
    name                  old time/op  new time/op  delta
    MapAssignInt32_255-8  24.7ns ± 3%  17.4ns ± 2%  -29.75%  (p=0.000 n=10+10)
    MapAssignInt32_64k-8  45.5ns ± 4%  37.6ns ± 4%  -17.18%  (p=0.000 n=10+10)
    MapAssignInt64_255-8  26.0ns ± 3%  17.9ns ± 4%  -31.03%  (p=0.000 n=10+10)
    MapAssignInt64_64k-8  46.9ns ± 5%  38.7ns ± 2%  -17.53%  (p=0.000 n=9+10)
    MapAssignStr_255-8    47.8ns ± 3%  24.8ns ± 4%  -48.01%  (p=0.000 n=10+10)
    MapAssignStr_64k-8    83.0ns ± 3%  51.9ns ± 3%  -37.45%  (p=0.000 n=10+9)
    
    name                     old time/op    new time/op    delta
    BinaryTree17-8              3.11s ±19%     2.78s ± 3%    ~     (p=0.095 n=5+5)
    Fannkuch11-8                3.26s ± 1%     3.21s ± 2%    ~     (p=0.056 n=5+5)
    FmtFprintfEmpty-8          50.3ns ± 1%    50.8ns ± 2%    ~     (p=0.246 n=5+5)
    FmtFprintfString-8         82.7ns ± 4%    80.1ns ± 5%    ~     (p=0.238 n=5+5)
    FmtFprintfInt-8            82.6ns ± 2%    81.9ns ± 3%    ~     (p=0.508 n=5+5)
    FmtFprintfIntInt-8          124ns ± 4%     121ns ± 3%    ~     (p=0.111 n=5+5)
    FmtFprintfPrefixedInt-8     158ns ± 6%     160ns ± 2%    ~     (p=0.341 n=5+5)
    FmtFprintfFloat-8           249ns ± 2%     245ns ± 2%    ~     (p=0.095 n=5+5)
    FmtManyArgs-8               513ns ± 2%     519ns ± 3%    ~     (p=0.151 n=5+5)
    GobDecode-8                7.48ms ±12%    7.11ms ± 2%    ~     (p=0.222 n=5+5)
    GobEncode-8                6.25ms ± 1%    6.03ms ± 2%  -3.56%  (p=0.008 n=5+5)
    Gzip-8                      252ms ± 4%     252ms ± 4%    ~     (p=1.000 n=5+5)
    Gunzip-8                   38.4ms ± 3%    38.6ms ± 2%    ~     (p=0.690 n=5+5)
    HTTPClientServer-8         76.9µs ±41%    66.4µs ± 6%    ~     (p=0.310 n=5+5)
    JSONEncode-8               16.5ms ± 3%    16.7ms ± 3%    ~     (p=0.421 n=5+5)
    JSONDecode-8               54.6ms ± 1%    54.3ms ± 2%    ~     (p=0.548 n=5+5)
    Mandelbrot200-8            4.45ms ± 3%    4.47ms ± 1%    ~     (p=0.841 n=5+5)
    GoParse-8                  3.43ms ± 1%    3.32ms ± 2%  -3.28%  (p=0.008 n=5+5)
    RegexpMatchEasy0_32-8      88.2ns ± 3%    89.4ns ± 2%    ~     (p=0.333 n=5+5)
    RegexpMatchEasy0_1K-8       205ns ± 1%     206ns ± 1%    ~     (p=0.905 n=5+5)
    RegexpMatchEasy1_32-8      85.1ns ± 1%    85.5ns ± 5%    ~     (p=0.690 n=5+5)
    RegexpMatchEasy1_1K-8       365ns ± 1%     371ns ± 9%    ~     (p=1.000 n=5+5)
    RegexpMatchMedium_32-8      129ns ± 2%     128ns ± 3%    ~     (p=0.730 n=5+5)
    RegexpMatchMedium_1K-8     39.8µs ± 0%    39.7µs ± 4%    ~     (p=0.730 n=4+5)
    RegexpMatchHard_32-8       1.99µs ± 3%    2.05µs ±16%    ~     (p=0.794 n=5+5)
    RegexpMatchHard_1K-8       59.3µs ± 1%    60.3µs ± 7%    ~     (p=1.000 n=5+5)
    Revcomp-8                   1.36s ±63%     0.52s ± 5%    ~     (p=0.095 n=5+5)
    Template-8                 62.6ms ±14%    60.5ms ± 5%    ~     (p=0.690 n=5+5)
    TimeParse-8                 330ns ± 2%     324ns ± 2%    ~     (p=0.087 n=5+5)
    TimeFormat-8                350ns ± 3%     340ns ± 1%  -2.86%  (p=0.008 n=5+5)
    
    name                     old speed      new speed      delta
    GobDecode-8               103MB/s ±11%   108MB/s ± 2%    ~     (p=0.222 n=5+5)
    GobEncode-8               123MB/s ± 1%   127MB/s ± 2%  +3.71%  (p=0.008 n=5+5)
    Gzip-8                   77.1MB/s ± 4%  76.9MB/s ± 3%    ~     (p=1.000 n=5+5)
    Gunzip-8                  505MB/s ± 3%   503MB/s ± 2%    ~     (p=0.690 n=5+5)
    JSONEncode-8              118MB/s ± 3%   116MB/s ± 3%    ~     (p=0.421 n=5+5)
    JSONDecode-8             35.5MB/s ± 1%  35.8MB/s ± 2%    ~     (p=0.397 n=5+5)
    GoParse-8                16.9MB/s ± 1%  17.4MB/s ± 2%  +3.45%  (p=0.008 n=5+5)
    RegexpMatchEasy0_32-8     363MB/s ± 3%   358MB/s ± 2%    ~     (p=0.421 n=5+5)
    RegexpMatchEasy0_1K-8    4.98GB/s ± 1%  4.97GB/s ± 1%    ~     (p=0.548 n=5+5)
    RegexpMatchEasy1_32-8     376MB/s ± 1%   375MB/s ± 5%    ~     (p=0.690 n=5+5)
    RegexpMatchEasy1_1K-8    2.80GB/s ± 1%  2.76GB/s ± 9%    ~     (p=0.841 n=5+5)
    RegexpMatchMedium_32-8   7.73MB/s ± 1%  7.76MB/s ± 3%    ~     (p=0.730 n=5+5)
    RegexpMatchMedium_1K-8   25.8MB/s ± 0%  25.8MB/s ± 4%    ~     (p=0.651 n=4+5)
    RegexpMatchHard_32-8     16.1MB/s ± 3%  15.7MB/s ±14%    ~     (p=0.794 n=5+5)
    RegexpMatchHard_1K-8     17.3MB/s ± 1%  17.0MB/s ± 7%    ~     (p=0.984 n=5+5)
    Revcomp-8                 273MB/s ±83%   488MB/s ± 5%    ~     (p=0.095 n=5+5)
    Template-8               31.1MB/s ±13%  32.1MB/s ± 5%    ~     (p=0.690 n=5+5)
    
    Updates #19495
    
    Change-Id: I116e9a2a4594769318b22d736464de8a98499909
    Reviewed-on: https://go-review.googlesource.com/38091
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 53937aad99f1d1e177fca3a3d0ec5c83825e015b
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Mon Mar 13 18:29:46 2017 +0200

    cmd/vet: check shift calculations with "unsafe" package
    
    vet should properly handle shift calculations via "unsafe"
    package after the CL 37950.
    
    Change-Id: I7737f2e656a5166337a17b92db46a0997f2a4e0e
    Reviewed-on: https://go-review.googlesource.com/38064
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cc71aa9ac4f2f1dc62f395a6f13ada8709d58213
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 10 22:09:43 2017 -0800

    cmd/compile/internal/ssa: make ARM's udiv like other calls
    
    Passes toolstash-check -all.
    
    Change-Id: Id389f8158cf33a3c0fcef373615b5351e7c74b5b
    Reviewed-on: https://go-review.googlesource.com/38082
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit b59a405656bbd79aefe3620553bee771628f9209
Author: David Chase <drchase@google.com>
Date:   Mon Mar 13 21:09:27 2017 +0000

    Revert "cmd/compile: de-virtualize interface calls"
    
    This reverts commit 4e0c7c3f61475116c4ae8d11ef796819d9c404f0.
    
    Reason for revert: The presence-of-optimization test program is fragile, breaks under noopt, and might break if the Go libraries are tweaked.  It needs to be (re)written without reference to other packages.
    
    Change-Id: I3aaf1ab006a1a255f961a978e9c984341740e3c7
    Reviewed-on: https://go-review.googlesource.com/38097
    Reviewed-by: Keith Randall <khr@golang.org>

commit 118b3fe7bbf855196db727daefbb403b84a4f67d
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 10 18:34:41 2017 -0800

    cmd/compile/internal/gc: refactor ACALL Prog creation
    
    This abstracts creation of ACALL Progs into package gc. The main
    benefit of this today is we can refactor away a lot of common
    boilerplate code.
    
    Later, once liveness analysis happens on the SSA graph, this will also
    provide an easy insertion point for emitting the PCDATA Progs
    immediately before call instructions.
    
    Passes toolstash-check -all.
    
    Change-Id: Ia15108ace97201cd84314f1ca916dfeb4f09d61c
    Reviewed-on: https://go-review.googlesource.com/38081
    Reviewed-by: Keith Randall <khr@golang.org>

commit 2e7c3b3f555853202fe0bdf2ea5ce37d7a56a7f7
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Mon Mar 13 20:35:07 2017 +0100

    encoding/gob: add Encode-Decode Int slices tests
    
    Tinkering with the gob package shows that is currently possible to
    *completely destroy* Int slices encoding without triggering a single
    test failure.
    
    The various encInt{8,16,32,64}Slice methods are only called during the
    execution of the GobMapInterfaceEncode test, which only encodes a few
    slices of length exactly 1 and then just checks that the error
    returned by Encode is nil (without trying to Decode back the data).
    
    This patch adds a few tests for signed integer slices encoding.
    
    Change-Id: Ifaaee2f32132873118b241f79aa8203e4ad31416
    Reviewed-on: https://go-review.googlesource.com/38066
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 08d8d5c986cd11951a6c2eb76587d4ec3ea9ecc7
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 10 17:42:02 2017 -0800

    cmd/compile/internal/ssa: replace {Defer,Go}Call with StaticCall
    
    Passes toolstash-check -all.
    
    Change-Id: Icf8b75364e4761a5e56567f503b2c1cb17382ed2
    Reviewed-on: https://go-review.googlesource.com/38080
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit a51e4cc9cea152a203ab508197dc0965c00e3a76
Author: khr <khr@khr-glaptop.roam.corp.google.com>
Date:   Thu Mar 9 10:38:45 2017 -0800

    cmd/compile: zero return parameters earlier
    
    Move the zeroing of results earlier.  In particular, they need to
    come before any move-to-heap operations, as those require allocation.
    Those allocations are points at which the GC can see the uninitialized
    result slots.
    
    For the function:
    
    func f() (x, y, z *int) {
      defer(){}()
      escape(&y)
      return
    }
    
    We used to generate code like this:
    
    x = nil
    y = nil
    &y = new(int)
    z = nil
    
    Now we will generate:
    
    x = nil
    y = nil
    z = nil
    &y = new(int)
    
    Since the fix for #18860, the return slots are always live if there
    is a defer, so the former ordering allowed the GC to see junk
    in the z slot.
    
    Fixes #19078
    
    Change-Id: I71554ae437549725bb79e13b2c100b2911d47ed4
    Reviewed-on: https://go-review.googlesource.com/38133
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 27492a2a549e4e03a6aed93811cdd458ce529e32
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Mon Mar 13 13:56:41 2017 -0500

    cmd/internal/obj/x86: remove unused const
    
    Since https://go-review.googlesource.com/24040 we no longer pad functions
    in asm6, so funcAlign is unused. Delete it.
    
    Change-Id: Id710e545a76b1797398f2171fe7e0928811fcb31
    Reviewed-on: https://go-review.googlesource.com/38134
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c694f6f3a80bcb42960e021b279eb6d23baf2c50
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 9 13:30:20 2017 -0800

    cmd/compile: eliminate more nil checks of phis
    
    The existing implementation started by eliminating
    nil checks for OpAddr, OpAddPtr, and OpPhis with
    all non-nil args.
    
    However, some OpPhis had all non-nil args,
    but their args had not been processed yet.
    
    Pull the OpPhi checks into their own loop,
    and repeat until stabilization.
    
    Eliminates a dozen additional nilchecks during make.bash.
    
    Negligible compiler performance impact.
    
    Change-Id: If7b803c3ad7582af7d9867d05ca13e03e109d864
    Reviewed-on: https://go-review.googlesource.com/37999
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 4e0c7c3f61475116c4ae8d11ef796819d9c404f0
Author: Philip Hofer <phofer@umich.edu>
Date:   Fri Mar 3 10:57:53 2017 -0800

    cmd/compile: de-virtualize interface calls
    
    With this change, code like
    
        h := sha1.New()
        h.Write(buf)
        sum := h.Sum()
    
    gets compiled into static calls rather than
    interface calls, because the compiler is able
    to prove that 'h' is really a *sha1.digest.
    
    The InterCall re-write rule hits a few dozen times
    during make.bash, and hundreds of times during all.bash.
    
    The most common pattern identified by the compiler
    is a constructor like
    
        func New() Interface { return &impl{...} }
    
    where the constructor gets inlined into the caller,
    and the result is used immediately. Examples include
    {sha1,md5,crc32,crc64,...}.New, base64.NewEncoder,
    base64.NewDecoder, errors.New, net.Pipe, and so on.
    
    Some existing benchmarks that change on darwin/amd64:
    
    Crc64/ISO4KB-8        2.67µs ± 1%    2.66µs ± 0%  -0.36%  (p=0.015 n=10+10)
    Crc64/ISO1KB-8         694ns ± 0%     690ns ± 1%  -0.59%  (p=0.001 n=10+10)
    Adler32KB-8            473ns ± 1%     471ns ± 0%  -0.39%  (p=0.010 n=10+9)
    
    On architectures like amd64, the reduction in code size
    appears to contribute more to benchmark improvements than just
    removing the indirect call, since that branch gets predicted
    accurately when called in a loop.
    
    Updates #19361
    
    Change-Id: Ia9d30afdd5f6b4d38d38b14b88f308acae8ce7ed
    Reviewed-on: https://go-review.googlesource.com/37751
    Run-TryBot: Philip Hofer <phofer@umich.edu>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 26e726c3092264584053a4f81714dcc8c91d2153
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 10 17:17:23 2017 -0800

    spec: remove superfluous terms "delimiter" and "special tokens"
    
    The (original) section on "Operators and Delimiters" introduced
    superfluous terminology ("delimiter", "special token") which
    didn't matter and was used inconsistently.
    
    Removed any mention of "delimiter" or "special token" and now
    simply group the special character tokens into "operators"
    (clearly defined via links), and "punctuation" (everything else).
    
    Fixes #19450.
    
    Change-Id: Ife31b24b95167ace096f93ed180b7eae41c66808
    Reviewed-on: https://go-review.googlesource.com/38073
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Rob Pike <r@golang.org>

commit eb715fbbbdc6b155512aed7c577e39b81bd35840
Author: Chew Choon Keat <choonkeat@gmail.com>
Date:   Mon Mar 13 03:21:21 2017 +0000

    net/http: unset proxy environment after test
    
    Fix last proxy in TestProxyFromEnvironment bleeds into other tests
    Change ResetProxyEnv to use the newer os.Unsetenv, instead of hard
    coding as ""
    
    Change-Id: I67cf833dbcf4bec2e10ea73c354334160cf05f84
    Reviewed-on: https://go-review.googlesource.com/38115
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit dd0e1acfeb50f33f79738b2ef7e21a61ecec9d22
Author: Dave Cheney <dave@cheney.net>
Date:   Mon Mar 13 14:30:44 2017 +1100

    cmd/compile/internal/gc: remove unused exportsize variable
    
    In Go 1.7 and earlier, gc.exportsize tracked the number of bytes
    written through exportf. With the removal of the old exporter in Go 1.8
    exportf is only used for printing the build id, and the header and
    trailer of the binary export format. The size of the export data is
    now returned directly from the exporter and exportsize is never
    referenced. Remove it.
    
    Change-Id: Id301144b3c26c9004c722d0c55c45b0e0801a88c
    Reviewed-on: https://go-review.googlesource.com/38116
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d5a9fbd556c441fa537c730681d6f8e7673b6e47
Author: Dave Cheney <dave@cheney.net>
Date:   Mon Mar 13 11:19:47 2017 +1100

    cmd/go/internal/get: remove unused tag selection code
    
    selectTag has been hard coded to only understand the tag `go1` since
    CL 6112060 which landed in 2012. The commit message asserted;
    
      Right now (before go1.0.1) there is only one possible tag,
      "go1", and I'd like to keep it that way.
    
    Remove goTag and the unused matching code in selectTag.
    
    Change-Id: I85f7c10f95704e22f8e8681266afd72bbcbe8fbd
    Reviewed-on: https://go-review.googlesource.com/38112
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7ee43faf78f3b0c97c315c28f13dd802047af0c8
Author: Dave Cheney <dave@cheney.net>
Date:   Mon Mar 13 11:35:39 2017 +1100

    runtime: remove sizeToClass
    
    CL 32219 added precomputed sizeclass tables.
    
    Remove the unused sizeToClass method which was previously only
    called from initSizes.
    
    Change-Id: I907bf9ed78430ecfaabbec7fca77ef2375010081
    Reviewed-on: https://go-review.googlesource.com/38113
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e831bd1fcae591a4dd985c4f75b455355033c75b
Author: Marcel Edmund Franke <marcel.edmund.franke@gmail.com>
Date:   Sun Mar 12 12:19:32 2017 +0100

    net/http: fix body close statement is missing
    
    Call body close after ioutil.ReadAll is done
    
    Change-Id: Ieceb1965a6a8f2dbc024e983acdfe22df17d07d1
    Reviewed-on: https://go-review.googlesource.com/38059
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b71ed4edc6a5663caac434f9c2bea47dbc37db15
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sat Mar 11 08:48:56 2017 +0100

    strconv: fix performance regression in integer formatting on 32bit platforms
    
    Some of the changes in CL golang.org/cl/38071/ assumed that / and %
    could always be combined to use only one DIV instruction. However,
    this is not the case for 64bit operands on a 32bit platform which use
    seperate runtime functions to calculate division and modulo.
    
    This CL restores the original optimizations that help on 32bit platforms
    with negligible impact on 64bit platforms.
    
    386:
    name          old time/op  new time/op  delta
    FormatInt-2   6.06µs ± 0%  6.02µs ± 0%  -0.70%  (p=0.000 n=20+20)
    AppendInt-2   4.98µs ± 0%  4.98µs ± 0%    ~     (p=0.747 n=18+18)
    FormatUint-2  1.93µs ± 0%  1.85µs ± 0%  -4.19%  (p=0.000 n=20+20)
    AppendUint-2  1.71µs ± 0%  1.64µs ± 0%  -3.68%  (p=0.000 n=20+20)
    
    amd64:
    name          old time/op  new time/op  delta
    FormatInt-2   2.41µs ± 0%  2.41µs ± 0%  -0.09%  (p=0.010 n=18+18)
    AppendInt-2   1.77µs ± 0%  1.77µs ± 0%  +0.08%  (p=0.000 n=18+18)
    FormatUint-2   653ns ± 1%   653ns ± 0%    ~     (p=0.178 n=20+20)
    AppendUint-2   514ns ± 0%   513ns ± 0%  -0.13%  (p=0.000 n=20+17)
    
    Change-Id: I574a18e54fb41b25fbe51ce696e7a8765abc79a6
    Reviewed-on: https://go-review.googlesource.com/38051
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit d343478295f8635e60014fa9481ccea04b6c53a8
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 10 22:22:11 2017 -0800

    cmd/link: eliminate markextra
    
    This appears to be leftover from when instruction selection happened
    in the linker. Many of the morestackX functions listed don't even
    exist anymore.
    
    Now that we select instructions within the compiler and assembler,
    normal deadcode elimination mechanisms should suffice for these
    symbols.
    
    Change-Id: I2cb1e435101392e7c983957c4acfbbcc87a5ca7d
    Reviewed-on: https://go-review.googlesource.com/38077
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6bc593805f2c2b0ef8b194cc186b804237744cb4
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Wed Mar 8 11:36:30 2017 +0200

    cmd/vet: eliminate "might be too small for shift" warnings
    
    Determine int, uint and uintptr bit sizes from GOARCH environment
    variable if it is set. Otherwise use host-specific sizes.
    
    Fixes #19321
    
    Change-Id: I494b8e4b49b59d32794f50ff2ce06ba040cb8460
    Reviewed-on: https://go-review.googlesource.com/37950
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit fe3458550bfee0c842393e44e096fd64c666d909
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Mar 10 17:03:52 2017 -0800

    cmd/go: if we get a C compiler dwarf2 warning, try without -g
    
    This avoids a problem that occurs on FreeBSD 11, in which the clang
    3.8 assembler issues a pointless warning when invoked with -g on a
    file that contains an empty .note.GNU-stack section.
    
    No test because there is no reasonable way to write one, but should
    fix the build on FreeBSD 11.
    
    Fixes #14705.
    
    Change-Id: I8c49bbf79a2c715c0e75495da19045fc92280e81
    Reviewed-on: https://go-review.googlesource.com/38072
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2de773d45f202d38c981a433880e867a7b5d0745
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 10 15:09:05 2017 -0800

    math/big: make nat.setUint64 vet-friendly
    
    nat.setUint64 is nicely generic.
    By assuming 32- or 64-bit words, however,
    we can write simpler code,
    and eliminate some shifts
    in dead code that vet complains about.
    
    Generated code for 64 bit systems is unaltered.
    Generated code for 32 bit systems is much better.
    For 386, the routine length drops from 325
    bytes of code to 271 bytes of code, with fewer loops.
    
    Change-Id: I1bc14c06272dee37a7fcb48d33dd1e621eba945d
    Reviewed-on: https://go-review.googlesource.com/38070
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 95c5227c15da3c0c61eeea70f0a8288088301b98
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 10 15:14:54 2017 -0800

    strconv: use % instead of computing the remainder from the quotient
    
    The compiler recognizes that in a sequence q = x/y; r = x%y only
    one division is required. Remove prior work-arounds and write
    more readable straight-line code (this also results in fewer
    instructions, though it doesn't appear to affect the benchmarks
    significantly).
    
    name          old time/op  new time/op  delta
    FormatInt-8   2.95µs ± 1%  2.92µs ± 5%   ~     (p=0.952 n=5+5)
    AppendInt-8   1.91µs ± 1%  1.89µs ± 2%   ~     (p=0.421 n=5+5)
    FormatUint-8   795ns ± 2%   782ns ± 4%   ~     (p=0.444 n=5+5)
    AppendUint-8   557ns ± 1%   557ns ± 2%   ~     (p=0.548 n=5+5)
    
    https://perf.golang.org/search?q=upload:20170310.1
    
    Also:
    - use uint instead of uintptr where we want to guarantee single-
      register operations
    - remove some unnecessary conversions (before indexing)
    - add more comments and fix some comments
    
    Change-Id: I04858dc2d798a6495879d9c7cfec2fdc2957b704
    Reviewed-on: https://go-review.googlesource.com/38071
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e19f184b8f61529980c24973d5522dc67e3d8525
Author: David NewHamlet <david@newhamlet.com>
Date:   Sat Mar 11 09:13:20 2017 +1300

    runtime: use cpuset_getaffinity for runtime.NumCPU() on FreeBSD
    
    In FreeBSD when run Go proc under a given sub-list of
    processors(e.g. 'cpuset -l 0 ./a.out' in multi-core system),
    runtime.NumCPU() still return all physical CPUs from sysctl
    hw.ncpu instead of account from sub-list.
    
    Fix by use syscall cpuset_getaffinity to account the number of sub-list.
    
    Fixes #15206
    
    Change-Id: If87c4b620e870486efa100685db5debbf1210a5b
    Reviewed-on: https://go-review.googlesource.com/29341
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit 135ce43c8731506d541329a1dfea2c737c6dd0b1
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Mar 10 10:04:42 2017 -0800

    cmd/go: when expanding "cmd", skip vendored main packages
    
    We are vendoring pprof from github.com/google/pprof, which comes with
    a main package. If we don't explicitly skip that main package, then
    `go install cmd` will install the compiled program in $GOROOT/bin.
    
    Fixes #19441.
    
    Change-Id: Ib268ffd16d4be65f7d80e4f8d9dc6e71523a94de
    Reviewed-on: https://go-review.googlesource.com/38007
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Raul Silvera <rsilvera@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit da0d23e5cdb305681a55c5475ff2db3e9a254cd4
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Fri Mar 10 17:07:42 2017 +0000

    runtime: remove unused ratep parameter
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: Iabcdfec2ae42c735aa23210b7183080d750682ca
    Reviewed-on: https://go-review.googlesource.com/38030
    Reviewed-by: Peter Weinberger <pjw@google.com>
    Run-TryBot: Peter Weinberger <pjw@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7e036521d505708a3e0d0b3d9bbe1e4853111211
Author: Bryan C. Mills <bcmills@google.com>
Date:   Thu Mar 9 16:45:14 2017 -0500

    expvar: add benchmark for (*Map).Set with per-goroutine keys
    
    Change-Id: I0fa68ca9812fe5e82ffb9d0b9598e95b47183eb8
    Reviewed-on: https://go-review.googlesource.com/38011
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d9fe2332ba63d1dd9416438a53b58bd6a91626b6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 9 22:02:13 2017 -0800

    net/http: change TestServerAllowsBlockingRemoteAddr to non-parallel
    
    It appears that this test is particularly
    sensitive to resource starvation.
    Returning it to non-parallel should reduce flakiness,
    by giving it the full system resources to run.
    
    Fixes #19161
    
    Change-Id: I6e8906516629badaa0cffeb5712af649dc197f39
    Reviewed-on: https://go-review.googlesource.com/38005
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4b3a55ec0243a291f4794c7057e120e000538886
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 9 14:28:32 2017 -0800

    cmd/compile: add 32 bit (AddPtr (Const)) rule
    
    This triggers about 50k times during 32 bit make.bash.
    
    Change-Id: Ia0c2b1a8246b92173b4b0d94a4037626f76b6e73
    Reviewed-on: https://go-review.googlesource.com/37998
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 3dcfce8d19e6fc98131106a6b1b7ce5445efc959
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 8 12:50:00 2017 -0800

    cmd/compile: add OpOffPtr [c] SP to constant cache
    
    They accounted for almost 30% of all CSE'd values.
    
    By never creating the duplicates in the first place,
    we reduce the high water mark of Value IDs,
    which in turn makes all SSA phases cheaper,
    particularly regalloc.
    
    name       old time/op     new time/op     delta
    Template       200ms ± 3%      198ms ± 4%  -0.87%  (p=0.016 n=50+49)
    Unicode       86.9ms ± 2%     85.5ms ± 3%  -1.56%  (p=0.000 n=49+50)
    GoTypes        553ms ± 4%      551ms ± 4%    ~     (p=0.183 n=50+49)
    SSA            3.97s ± 3%      3.93s ± 2%  -1.06%  (p=0.000 n=48+48)
    Flate          124ms ± 4%      124ms ± 3%    ~     (p=0.545 n=48+50)
    GoParser       146ms ± 4%      146ms ± 4%    ~     (p=0.810 n=49+49)
    Reflect        357ms ± 3%      355ms ± 3%  -0.59%  (p=0.049 n=50+48)
    Tar            106ms ± 4%      107ms ± 5%    ~     (p=0.454 n=49+50)
    XML            203ms ± 4%      203ms ± 4%    ~     (p=0.726 n=48+50)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        237M ± 3%       235M ± 4%    ~     (p=0.208 n=47+48)
    Unicode         111M ± 4%       108M ± 9%  -2.50%  (p=0.000 n=47+50)
    GoTypes         736M ± 5%       729M ± 4%  -0.95%  (p=0.017 n=50+46)
    SSA            5.73G ± 4%      5.74G ± 4%    ~     (p=0.765 n=50+50)
    Flate           150M ± 5%       148M ± 6%  -0.89%  (p=0.045 n=48+47)
    GoParser        180M ± 5%       178M ± 7%  -1.34%  (p=0.012 n=50+50)
    Reflect         450M ± 4%       444M ± 4%  -1.40%  (p=0.000 n=50+49)
    Tar             124M ± 7%       123M ± 7%    ~     (p=0.092 n=50+50)
    XML             248M ± 6%       245M ± 5%    ~     (p=0.057 n=50+50)
    
    name       old alloc/op    new alloc/op    delta
    Template      39.4MB ± 0%     39.3MB ± 0%  -0.37%  (p=0.000 n=50+50)
    Unicode       30.9MB ± 0%     30.9MB ± 0%  -0.27%  (p=0.000 n=48+50)
    GoTypes        114MB ± 0%      113MB ± 0%  -1.03%  (p=0.000 n=50+49)
    SSA            882MB ± 0%      865MB ± 0%  -1.95%  (p=0.000 n=49+49)
    Flate         25.8MB ± 0%     25.7MB ± 0%  -0.21%  (p=0.000 n=50+50)
    GoParser      31.7MB ± 0%     31.6MB ± 0%  -0.33%  (p=0.000 n=50+50)
    Reflect       79.7MB ± 0%     79.3MB ± 0%  -0.49%  (p=0.000 n=44+49)
    Tar           27.2MB ± 0%     27.1MB ± 0%  -0.31%  (p=0.000 n=50+50)
    XML           42.7MB ± 0%     42.3MB ± 0%  -1.05%  (p=0.000 n=48+49)
    
    name       old allocs/op   new allocs/op   delta
    Template        379k ± 1%       380k ± 1%  +0.26%  (p=0.000 n=50+50)
    Unicode         324k ± 1%       324k ± 1%    ~     (p=0.964 n=49+50)
    GoTypes        1.14M ± 0%      1.15M ± 0%  +0.14%  (p=0.000 n=50+49)
    SSA            7.89M ± 0%      7.89M ± 0%  -0.05%  (p=0.000 n=49+49)
    Flate           240k ± 1%       241k ± 1%  +0.27%  (p=0.001 n=50+50)
    GoParser        310k ± 1%       311k ± 1%  +0.48%  (p=0.000 n=50+49)
    Reflect        1.00M ± 0%      1.00M ± 0%  +0.17%  (p=0.000 n=48+50)
    Tar             254k ± 1%       255k ± 1%  +0.23%  (p=0.005 n=50+50)
    XML             395k ± 1%       395k ± 1%  +0.19%  (p=0.002 n=49+47)
    
    Change-Id: Iaa8f5f37e23bd81983409f7359f9dcd4dfe2961f
    Reviewed-on: https://go-review.googlesource.com/38003
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit d11a2184fb29d0f8a447b2e70299dc410c5642ed
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 8 14:48:43 2017 -0800

    cmd/compile: allow earlier GC of freed constant value
    
    Minor fix, because it's the right thing to do.
    No significant impact.
    
    Change-Id: I2138285d397494daa9a88c414149c2a7860edd7e
    Reviewed-on: https://go-review.googlesource.com/38001
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 678f35b676de075375066ade2935296dfb8050ec
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 8 12:48:04 2017 -0800

    cmd/compile: fix SSA type for first runtime call arg/result
    
    CLs 37254 and 37869 contained similar fixes.
    
    Change-Id: I0cbf01c691b54d82acef398489df6e9c89ebb83f
    Reviewed-on: https://go-review.googlesource.com/38000
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit c63ad970f6dc497b2cd529357201dd46bc3ee9b0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 8 12:45:12 2017 -0800

    cmd/compile: rename Func.constVal arg for clarity
    
    Values have an Aux and an AuxInt.
    We're setting AuxInt, not Aux.
    Say so.
    
    Change-Id: I41aa783273bb7e1ba47c941aa4233f818e37dadd
    Reviewed-on: https://go-review.googlesource.com/37997
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2a5cf48f91c0a59eeb01b97ca6afaca311324206
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Mar 8 14:26:23 2017 -0800

    cmd/compile: print columns (not just lines) in error messages
    
    Compiler errors now show the exact line and line byte offset (sometimes
    called "column") of where an error occured. For `go tool compile x.go`:
    
            package p
            const c int = false
            //line foo.go:123
            type t intg
    
    reports
    
            x.go:2:7: cannot convert false to type int
            foo.go:123[x.go:4:8]: undefined: intg
    
    (Some errors use the "wrong" position for the error message; arguably
    the byte offset for the first error should be 15, the position of 'false',
    rathen than 7, the position of 'c'. But that is an indepedent issue.)
    
    The byte offset (column) values are measured in bytes; they start at 1,
    matching the convention used by editors and IDEs.
    
    Positions modified by //line directives show the line offset only for the
    actual source location (in square brackets), not for the "virtual" file and
    line number because that code is likely generated and the //line directive
    only provides line information.
    
    Because the new format might break existing tools or scripts, printing
    of line offsets can be disabled with the new compiler flag -C. We plan
    to remove this flag eventually.
    
    Fixes #10324.
    
    Change-Id: I493f5ee6e78457cf7b00025aba6b6e28e50bb740
    Reviewed-on: https://go-review.googlesource.com/37970
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 1f9f0ea32b2dcee027b107f2c3d0bc723274a810
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 9 13:38:10 2017 -0800

    cmd/compile/internal/syntax: start line offset (column) numbers at 1
    
    We could leave it alone and fix line offset (column) numbers when
    reporting errors, but that is likely to cause confusion (internal
    numbers don't match reported numbers). Instead, switch to default
    numbering starting at 1.
    
    For package syntax-internal use only, introduced constants defining
    the line and column bases, and use them throughout the code and its
    tests. It is possible to change these constants and package syntax
    will continue to work. But changing them is going to break any client
    that makes explicit assumptions about line and column numbers (which
    is "all of them").
    
    Change-Id: Ia3d136a8ec8d9372ed9c05ca47d3dff222cf030e
    Reviewed-on: https://go-review.googlesource.com/37996
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit bfc164c64d33edfaf774b5c29b9bf5648a6447fb
Author: Paul Marks <pmarks@google.com>
Date:   Tue Nov 1 21:01:08 2016 -0700

    net: add Resolver.StrictErrors
    
    When LookupIP is performing multiple subqueries, this option causes a
    timeout/servfail affecting a single query to abort the whole operation,
    instead of returning a partial (IPv4/IPv6-only) result.
    
    Similarly, operations that walk the DNS search list will also abort when
    encountering one of these errors.
    
    Fixes #17448
    
    Change-Id: Ice22e4aceb555c5a80d19bd1fde8b8fe87ac9517
    Reviewed-on: https://go-review.googlesource.com/32572
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b0e91d836a0abd46899cf78fdd93303afcf6c637
Author: Philip Hofer <phofer@umich.edu>
Date:   Fri Mar 3 13:44:18 2017 -0800

    cmd/compile: clean up ssa.Value memory arg usage
    
    This change adds a method to replace expressions
    of the form
    
       v.Args[len(v.Args)-1]
    
    so that the code's intention to walk memory arguments
    is explicit.
    
    Passes toolstash-check.
    
    Change-Id: I0c80d73bc00989dd3cdf72b4f2c8e1075a2515e0
    Reviewed-on: https://go-review.googlesource.com/37757
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit e471ad9189d1eba54c8cb5414c47e413cea78df2
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Mar 8 06:56:40 2017 -0500

    cmd/compile: remove duplicated zeroing of outgoing args
    
    Outgoing arg zeroing code is inserted at walk.go:paramstoheap.
    Don't do it twice.
    
    Change-Id: I70afac6af9e39b3efce0a6a79d6193428d922708
    Reviewed-on: https://go-review.googlesource.com/37863
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 027500ce385709ecaa8fe11320a23a12f8e3b3de
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Thu Mar 9 20:50:24 2017 +0000

    src/*.bash: use tabs consistently
    
    make.bash used mostly tabs and buildall.bash used mostly spaces, but
    they were both mixing them. Be consistent and use tabs, as that's what's
    more common and what the Go code uses.
    
    Change-Id: Ia6affbfccfe64fda800c1ac400965df364d2c545
    Reviewed-on: https://go-review.googlesource.com/37967
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 945180fe2aa3238bbc23f336a00eba934daa9ccc
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Mar 9 13:03:07 2017 -0500

    cmd/compile: fix OffPtr type in 2-field struct Store rule
    
    The type of the OffPtr for the first field was incorrect. It should
    have been a pointer to the field type, rather than the field
    type itself.
    
    Fixes #19475.
    
    Change-Id: I3960b404da0f4bee759331126cce6140d2ce1df7
    Reviewed-on: https://go-review.googlesource.com/37869
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 4210930a2861c8938baabaa5b8097e8b28d92934
Author: Bryan C. Mills <bcmills@google.com>
Date:   Tue Mar 7 18:50:52 2017 -0500

    runtime/cgo: return correct sa_flags
    
    A typo in the previous revision ("act" instead of "oldact") caused us
    to return the sa_flags from the new (or zeroed) sigaction rather than
    the old one.
    
    In the presence of a signal handler registered before
    runtime.libpreinit, this caused setsigstack to erroneously zero out
    important sa_flags (such as SA_SIGINFO) in its attempt to re-register
    the existing handler with SA_ONSTACK.
    
    Change-Id: I3cd5152a38ec0d44ae611f183bc1651d65b8a115
    Reviewed-on: https://go-review.googlesource.com/37852
    Run-TryBot: Bryan Mills <bcmills@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e57350f4c091c80333d37b2ba5af50de193147fa
Author: Bryan C. Mills <bcmills@google.com>
Date:   Wed Mar 8 15:40:28 2017 -0500

    runtime: fix _cgo_yield usage with sysmon and on BSD
    
    There are a few problems from change 35494, discovered during testing
    of change 37852.
    
    1. I was confused about the usage of n.key in the sema variant, so we
       were looping on the wrong condition. The error was not caught by
       the TryBots (presumably due to missing TSAN coverage in the BSD and
       darwin builders?).
    
    2. The sysmon goroutine sometimes skips notetsleep entirely, using
       direct usleep syscalls instead. In that case, we were not calling
       _cgo_yield, leading to missed signals under TSAN.
    
    3. Some notetsleep calls have long finite timeouts. They should be
       broken up into smaller chunks with a yield at the end of each
       chunk.
    
    updates #18717
    
    Change-Id: I91175af5dea3857deebc686f51a8a40f9d690bcc
    Reviewed-on: https://go-review.googlesource.com/37867
    Run-TryBot: Bryan Mills <bcmills@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 91563ced5897faf729a34be7081568efcfedda31
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Thu Mar 9 13:20:54 2017 +0100

    time: make the ParseInLocation test more robust
    
    The tzdata 2017a update (2017-02-28) changed the abbreviation of the
    Asia/Baghdad time zone (used in TestParseInLocation) from 'AST' to the
    numeric '+03'.
    
    Update the test so that it skips the checks if we're using a recent
    tzdata release.
    
    Fixes #19457
    
    Change-Id: I45d705a5520743a611bdd194dc8f8d618679980c
    Reviewed-on: https://go-review.googlesource.com/37964
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3b689225751dd5a09569d94b787c01fb166c3e1d
Author: Elias Naur <elias.naur@gmail.com>
Date:   Thu Mar 9 10:23:10 2017 +0100

    misc/cgo/testcarchive: add missing header
    
    write(2) is defined in unistd.h.
    
    For the iOS builder.
    
    Change-Id: I411ffe81988d8fbafffde89e4732a20af1a63325
    Reviewed-on: https://go-review.googlesource.com/37962
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b2c391b70c43edea2a6c536afba6ebf925844e1b
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Mon Mar 6 20:39:24 2017 +0200

    cmd/compile/internal/gc: shrink Sym by 8 bytes on amd64
    
    Move 8-bit flags field after 32-bit Block field
    
    Change-Id: I8e5e9a2285477aac2402a839a105e710d5340224
    Reviewed-on: https://go-review.googlesource.com/37848
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 67d69f899ce0fcd3d7aee0a07554fb84770020c0
Author: Keith Randall <khr@golang.org>
Date:   Wed Mar 8 16:31:02 2017 -0800

    cmd/compile: set base register of spill/restore to SP
    
    Previously the base register was unset, which lead to the disassembler
    using "FP" instead of "SP" as the base register.  That lead to some
    confusion as to what the difference betweeen the two was.
    Be consistent and always use SP.
    
    Fixes #19458
    
    Change-Id: Ie8f8ee54653bd202c0cf6fbf1d350e3c8c8b67a0
    Reviewed-on: https://go-review.googlesource.com/37971
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 3a1271dad7617a9766d21987b2c3cc0e683d2010
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Mar 8 22:54:52 2017 +0100

    go/internal/srcimporter: skip tests on iOS
    
    The iOS test harness only includes the current test directory in its
    app bundles, but the tests need access to all source code.
    
    Change-Id: I8a902b183bc2745b4fbfffef867002d573abb1f5
    Reviewed-on: https://go-review.googlesource.com/37961
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 23be728950be7973a4c4852449e1987c64dc2739
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 4 16:54:50 2017 -0800

    runtime: optimize slicebytestostring
    
    Inline rawstringtmp and simplify.
    Use memmove instead of copy.
    
    name                     old time/op  new time/op  delta
    SliceByteToString/1-8    19.4ns ± 2%  14.1ns ± 1%  -27.04%  (p=0.000 n=20+17)
    SliceByteToString/2-8    20.8ns ± 2%  15.5ns ± 2%  -25.46%  (p=0.000 n=20+20)
    SliceByteToString/4-8    20.7ns ± 1%  14.9ns ± 1%  -28.30%  (p=0.000 n=20+20)
    SliceByteToString/8-8    23.2ns ± 1%  17.1ns ± 1%  -26.22%  (p=0.000 n=19+19)
    SliceByteToString/16-8   29.4ns ± 1%  23.6ns ± 1%  -19.76%  (p=0.000 n=17+20)
    SliceByteToString/32-8   31.4ns ± 1%  26.0ns ± 1%  -17.11%  (p=0.000 n=16+19)
    SliceByteToString/64-8   36.1ns ± 0%  30.0ns ± 0%  -16.96%  (p=0.000 n=16+16)
    SliceByteToString/128-8  46.9ns ± 0%  38.9ns ± 0%  -17.15%  (p=0.000 n=17+19)
    
    Change-Id: I422e688830e4a9bd21897d1f74964625b735f436
    Reviewed-on: https://go-review.googlesource.com/37791
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Marvin Stenger <marvin.stenger94@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 228438e0973b688829fdd601e31352920a5914f4
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Mar 8 21:56:57 2017 +0100

    os/user: fake Current on Android
    
    On Android devices where the stub fallback for Current fails to
    extract a User from the environment, return a dummy fallback instead
    of failing.
    
    While we're here, use / instead of /home/nacl for the NaCL fallback.
    
    Hopefully fixes the Android builder.
    
    Change-Id: Ia29304fbc224ee5f9c0f4e706d1756f765a7eae5
    Reviewed-on: https://go-review.googlesource.com/37960
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 29edf0f9feb0e7412788a20e7d8d473270cb9342
Author: Bryan C. Mills <bcmills@google.com>
Date:   Thu Jan 19 16:09:10 2017 -0500

    runtime: poll libc to deliver signals under TSAN
    
    fixes #18717
    
    Change-Id: I7244463d2e7489e0b0fe3b74c4b782e71210beb2
    Reviewed-on: https://go-review.googlesource.com/35494
    Run-TryBot: Bryan Mills <bcmills@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d71f36b5aa1eadc6cd86ada2c0d5dd621bd9fd82
Author: David Chase <drchase@google.com>
Date:   Thu Feb 2 11:53:41 2017 -0500

    cmd/compile: check loop rescheduling with stack bound, not counter
    
    After benchmarking with a compiler modified to have better
    spill location, it became clear that this method of checking
    was actually faster on (at least) two different architectures
    (ppc64 and amd64) and it also provides more timely interruption
    of loops.
    
    This change adds a modified FOR loop node "FORUNTIL" that
    checks after executing the loop body instead of before (i.e.,
    always at least once).  This ensures that a pointer past the
    end of a slice or array is not made visible to the garbage
    collector.
    
    Without the rescheduling checks inserted, the restructured
    loop from this  change apparently provides a 1% geomean
    improvement on PPC64 running the go1 benchmarks; the
    improvement on AMD64 is only 0.12%.
    
    Inserting the rescheduling check exposed some peculiar bug
    with the ssa test code for s390x; this was updated based on
    initial code actually generated for GOARCH=s390x to use
    appropriate OpArg, OpAddr, and OpVarDef.
    
    NaCl is disabled in testing.
    
    Change-Id: Ieafaa9a61d2a583ad00968110ef3e7a441abca50
    Reviewed-on: https://go-review.googlesource.com/36206
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 6fbedc1afead2de1fd554e72f6df47a4b7948b5a
Author: Kevin Burke <kev@inburke.com>
Date:   Wed Mar 8 10:22:00 2017 -0800

    database/sql: fix spelling mistake in tests
    
    Change-Id: I04e150d4e4123aad2f277e5c6e9f2abd15628a28
    Reviewed-on: https://go-review.googlesource.com/37941
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2123a6c64455f6e1cceeefa97e4a033c873e2631
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Mar 7 15:58:28 2017 -0800

    cmd/compile: fix recorded export data position info
    
    The position information recorded now consists of the line-
    directive relative filename and line number. It would be
    relatively easy to also encode absolute position information
    as necessary (by serializing src.PosBase data).
    
    For example, given $GOROOT/src/tmp/x.go:
    
            package p
    
            const C0 = 0
    
            //line c.go:10
            const C1 = 1
    
            //line t.go:20
            type T int
    
            //line v.go:30
            var V T
    
            //line f.go:40
            func F() {}
    
    The recorded positions for the exported entities are:
    
            C0 $GOROOT/src/tmp/x.go 3
            C1 c.go 10
            T t.go 20
            V v.go 30
            F f.go 40
    
    Fix verified by manual inspection. There's currently no easy way
    to test this, but it will eventually be tested when we fix #7311.
    
    Fixes #19391.
    
    Change-Id: I6269067ea58358250fe6dd1f73bdf9e5d2adfe3d
    Reviewed-on: https://go-review.googlesource.com/37936
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 82e1732f14d6b50ba29dcc3f0eb71c80fa52c2d3
Author: Sarah Adams <shadams@google.com>
Date:   Thu Jan 19 14:17:10 2017 -0800

    database/sql: proper prepared statement support in transactions
    
    This change was originally written by Marko Tiikkaja <marko@joh.to>.
    https://go-review.googlesource.com/#/c/2035/
    
    Previously *Tx.Stmt always prepared a new statement, even if an
    existing one was available on the connection the transaction was on.
    Now we first see if the statement is already available on the
    connection and only prepare if it isn't. Additionally, when we do
    need to prepare one, we store it in the parent *Stmt to allow it to be
    later reused by other calls to *Tx.Stmt on that statement or just
    straight up by *Stmt.Exec et al.
    
    To make sure that the statement doesn't disappear unexpectedly, we
    record a dependency from the statement returned by *Tx.Stmt to the
    *Stmt it came from and set a new field, parentStmt, to point to the
    originating *Stmt. When the transaction's *Stmt is closed, we remove
    the dependency. This way the "parent" *Stmt can be closed by the user
    without her having to know whether any transactions are still using it
    or not.
    
    Fixes #15606
    
    Change-Id: I41b5056847e117ac61130328b0239d1e000a4a08
    Reviewed-on: https://go-review.googlesource.com/35476
    Run-TryBot: Daniel Theophanes <kardianos@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Daniel Theophanes <kardianos@gmail.com>

commit 3b988eb643e6f2a1cae8981590891fd95e8112e5
Author: Johan Brandhorst <johan.brandhorst@gmail.com>
Date:   Sat Mar 4 18:24:44 2017 +0000

    net/http: use httptest.Server Client in tests
    
    After merging https://go-review.googlesource.com/c/34639/,
    it was pointed out to me that a lot of tests under net/http
    could use the new functionality to simplify and unify testing.
    
    Using the httptest.Server provided Client removes the need to
    call CloseIdleConnections() on all Transports created, as it
    is automatically called on the Transport associated with the
    client when Server.Close() is called.
    
    Change the transport used by the non-TLS
    httptest.Server to a new *http.Transport rather than using
    http.DefaultTransport implicitly. The TLS version already
    used its own *http.Transport. This change is to prevent
    concurrency problems with using DefaultTransport implicitly
    across several httptest.Server's.
    
    Add tests to ensure the httptest.Server.Client().Transport
    RoundTripper interface is implemented by a *http.Transport,
    as is now assumed across large parts of net/http tests.
    
    Change-Id: I9f9d15f59d72893deead5678d314388718c91821
    Reviewed-on: https://go-review.googlesource.com/37771
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2bd6360e3b0bda74b637ad37536b2f95f5b8574f
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Wed Mar 8 07:47:31 2017 +0900

    net/mail: fix wrong error message in consumePhrase
    
    Fixes #19415
    
    Change-Id: I6414f82e42bd09f1793156befce326aeac919ea2
    Reviewed-on: https://go-review.googlesource.com/37911
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c797256a8f7dcf24fdc798b1e0845d58aeeef7a2
Author: Russ Cox <rsc@golang.org>
Date:   Mon Mar 6 20:39:57 2017 -0500

    runtime/pprof: add GNU build IDs to Mappings recorded from /proc/self/maps
    
    This helps systems that maintain an external database mapping
    build ID to symbol information for the given binary, especially
    in the case where /proc/self/maps lists many different files
    (for example, many shared libraries).
    
    Avoid importing debug/elf to avoid dragging in that whole
    package (and its dependencies like debug/dwarf) into the
    build of every program that generates a profile.
    
    Fixes #19431.
    
    Change-Id: I6d4362a79fe23e4f1726dffb0661d20bb57f766f
    Reviewed-on: https://go-review.googlesource.com/37855
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 752d7bad4fc6cb4c70edbe0b735ca89d7da16732
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Mar 7 14:53:56 2017 -0800

    cmd/internal/src: fix Pos.String() for positions after line directives
    
    The old code simply printed the position of the line directive in
    square brackets for a position modified by a line directive. Now
    we print the corresponding actual source file position instead.
    
    Fixes #19392.
    
    Change-Id: I933f3e435d03a6ee8269df36ae35f9202b7b2e76
    Reviewed-on: https://go-review.googlesource.com/37932
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 38409f5f35b00979cfe491a4fec6c93a6f58e037
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Mon Mar 6 18:39:02 2017 +0900

    internal/poll: code cleanup
    
    This change adds missing docs, collapses single-line import paths,
    removes unsed method placeholders and renames str.go to strconv.go.
    
    Change-Id: I2d155c838935cd8427abd142a462ff4c56829703
    Reviewed-on: https://go-review.googlesource.com/37814
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d50f892abca46b794a23f20777c0b2425467d407
Author: Austin Clements <austin@google.com>
Date:   Tue Mar 7 15:36:49 2017 -0500

    runtime: join selectgo and selectgoImpl
    
    Currently selectgo is just a wrapper around selectgoImpl. This keeps
    the hard-coded frame skip counts for tracing the same between the
    channel implementation and the select implementation.
    
    However, this is fragile and confusing, so pass a skip parameter to
    send and recv, join selectgo and selectgoImpl into one function, and
    use decrease all of the skips in selectgo by one.
    
    Change-Id: I11b8cbb7d805b55f5dc6ab4875ac7dde79412ff2
    Reviewed-on: https://go-review.googlesource.com/37860
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 5e4a958351222233fbc6f82ab621a7d15299eea5
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 7 12:31:59 2017 -0800

    cmd/compile: refactor portable SSA Op handling
    
    Several SSA ops will always behave identically regardless of target
    architecture, so handle those within gc/ssa.go instead.
    
    Passes toolstash-check -all.
    
    Change-Id: I54d514e80ab86723e44332a5a38e3054cbca8c5d
    Reviewed-on: https://go-review.googlesource.com/37931
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b992c2649e6c30d25baeb3c26aba0e90b800a1f4
Author: Austin Clements <austin@google.com>
Date:   Tue Mar 7 15:20:40 2017 -0500

    runtime: print SP/FP on bad pointer crashes
    
    If the bad pointer is on a stack, this makes it possible to find the
    frame containing the bad pointer.
    
    Change-Id: Ieda44e054aa9ebf22d15d184457c7610b056dded
    Reviewed-on: https://go-review.googlesource.com/37858
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit caa7dacfd27beb493c262b18a9aae514863ddec7
Author: Austin Clements <austin@google.com>
Date:   Tue Mar 7 15:24:02 2017 -0500

    runtime: honor GOTRACEBACK=crash even if _g_.m.traceback != 0
    
    Change-Id: I6de1ef8f67bde044b8706c01e98400e266e1f8f0
    Reviewed-on: https://go-review.googlesource.com/37857
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit c310c688ffa46e2f91e40284c16d71f3921feed9
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Mar 1 15:50:57 2017 -0800

    cmd/compile, runtime: simplify multiway select implementation
    
    This commit reworks multiway select statements to use normal control
    flow primitives instead of the previous setjmp/longjmp-like behavior.
    This simplifies liveness analysis and should prevent issues around
    "returns twice" function calls within SSA passes.
    
    test/live.go is updated because liveness analysis's CFG is more
    representative of actual control flow. The case bodies are the only
    real successors of the selectgo call, but previously the selectsend,
    selectrecv, etc. calls were included in the successors list too.
    
    Updates #19331.
    
    Change-Id: I7f879b103a4b85e62fc36a270d812f54c0aa3e83
    Reviewed-on: https://go-review.googlesource.com/37661
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 5ed952368e3777845afd934e38219c5567b09cc4
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Tue Mar 7 16:05:40 2017 +0000

    runtime/pprof: actually use tag parameter
    
    It's only ever called with the value it was using, but the code was
    counterintuitive. Use the parameter instead, like the other funcs near
    it.
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: I45855e11d749380b9b2a28e6dd1d5dedf119a19b
    Reviewed-on: https://go-review.googlesource.com/37893
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 143fd8ef2ab4450d7298b59a45d15232518f982d
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sat Mar 4 04:55:56 2017 +0100

    os/user: use the stubs fallback for Android
    
    Using the stubs, user.Current will no longer fail on Android, fixing
    the os/exec.TestCredentialNoSetGroups test.
    
    Change-Id: I8b9842aa6704c0cde383c549a614bab0a0ed7695
    Reviewed-on: https://go-review.googlesource.com/37765
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e99dafc4a8d631992903250378a8007daf794f2c
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Wed Jan 18 11:48:14 2017 +0100

    cmd/compile: fix misleading "truncated to int" messages
    
    When defining an int const, the compiler tries to cast the RHS
    expression to int. The cast may fail for three reasons:
    
      1. expr is an integer constant that overflows int
      2. expr is a floating point constant
      3. expr is a complex constant, or not a number
    
    In the second case, in order to print a sensible error message, we
    must distinguish between a floating point constant that should be
    included in the error message and a floating point constant that
    cannot be reasonably formatted for inclusion in an error message.
    
    For example, in:
    
      const a int = 1.1
      const b int = 1 + 1e-100
    
    a is in the former group, while b is in the latter, since the floating
    point value resulting from the evaluation of the rhs of the assignment
    (1.00...01) is too long to be fully printed in an error message, and
    cannot be shortened without making the error message misleading
    (rounding or truncating it would result in a "1", which looks like an
    integer constant, and it makes little sense in an error message about
    an invalid floating point expression).
    
    To fix this problem, we try to format the float value using fconv
    (which is used by the error reporting mechanism to format float
    arguments), and then parse the resulting string back to a
    big.Float. If the result is an integer, we assume that expr is a float
    value that cannot be reasonably be formatted as a string, and we emit
    an error message that does not include its string representation.
    
    Also, change the error message for overflows to a more conservative
    "integer too large", which does not mention overflows that are only
    caused by an internal implementation restriction.
    
    Also, change (*Mpint) SetFloat so that it returns a bool (instead of
    0/-1 for success/failure).
    
    Fixes #11371
    
    Change-Id: Ibbc73e2ed2eaf41f07827b0649d0eb637150ecaa
    Reviewed-on: https://go-review.googlesource.com/35411
    Run-TryBot: Alberto Donizetti <alb.donizetti@gmail.com>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 527f3518dac8554161fd0535ab2c35df507062ff
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 7 10:57:55 2017 -0800

    cmd/compile/internal/gc: skip autotmp vars in gc again
    
    Instead of skipping them based on string matching much later in the
    compilation process, skip them up front using the proper API.
    
    Passes toolstash-check.
    
    Change-Id: Ibd4c0448a0701ba0de3235d4689ef300235fa1d9
    Reviewed-on: https://go-review.googlesource.com/37930
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b1a4424a52687e5abca29cd795c5701b5639a52f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Mar 7 10:44:12 2017 -0800

    cmd/internal/obj: change started to bool
    
    Change-Id: I90143e3c6e95a1495f300ffeb10de554aa41f56a
    Reviewed-on: https://go-review.googlesource.com/37889
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b91b694b37b12724ae3c3925eb9a215ac9b36271
Author: Elias Naur <elias.naur@gmail.com>
Date:   Tue Mar 7 20:04:23 2017 +0100

    runtime/pprof: fix the protobuf tests on Android
    
    Change-Id: I5f85a7980b9a18d3641c4ee8b0992671a8421bb0
    Reviewed-on: https://go-review.googlesource.com/37896
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c5cdda401e483a3a4dae7cc61eb78521ab953b04
Author: Kevin Burke <kev@inburke.com>
Date:   Tue Mar 7 09:03:40 2017 -0800

    encoding/base64, html/template: fix grammar mistakes
    
    Replace 'does not contains' with 'does not contain' where it appears
    in the source code.
    
    Change-Id: Ie7266347c429512c8a41a7e19142afca7ead3922
    Reviewed-on: https://go-review.googlesource.com/37887
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f639353330797b819f29fa2b9b4b73d09b4c7584
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Mon Mar 6 09:59:32 2017 +0900

    mime: fix panic parsing 'encoded-word', be stricter
    
    Fixes #19416
    
    Change-Id: I23c69ff637abaa202909f1cba6ed41b3cfe3d117
    Reviewed-on: https://go-review.googlesource.com/37812
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 68177d9ec034b5d2f0ef35fa1a225616c2af3827
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 6 07:32:37 2017 -0800

    cmd/internal/obj: move dwarf.Var generation into compiler
    
    Passes toolstash -cmp.
    
    Change-Id: I4bd60f7ebba5457e7b3ece688fee2351bfeeb59a
    Reviewed-on: https://go-review.googlesource.com/37874
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alessandro Arzilli <alessandro.arzilli@gmail.com>
    Reviewed-by: Than McIntosh <thanm@google.com>

commit e577a55b78bb2d36841504c00ff1d984c167308e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 6 23:49:49 2017 -0800

    cmd/compile: change signatlist to []*Type
    
    No need to keep as Nodes when they're all Types anyway.
    
    Change-Id: I8157914ba5b09cadf2263247844680a60233a0f2
    Reviewed-on: https://go-review.googlesource.com/37886
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2a26f5809e4e80e7d8d4e20b9965efb2eefe71c5
Author: Andrew Benton <andrewmbenton@gmail.com>
Date:   Mon Mar 6 23:17:58 2017 -0800

    crypto/x509: rename and reposition rsaPublicKey struct declaration
    
    For consistency with the other named types in this package, this
    change renames the unexported rsaPublicKey struct to pkcs1PublicKey
    and positions the declaration up with the other similarly-named
    types in pkcs1.go.
    
    See the final comment of #19355 for discussion.
    
    Change-Id: I1fa0366a8efa01602b81bc69287ef747abce84f5
    Reviewed-on: https://go-review.googlesource.com/37885
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit aeac77dce6335395309ad236adb1627b389a1d86
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Tue Mar 7 13:09:05 2017 +0000

    net: remove unused Interface parameter
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: I4795dd0221784d10cf7c9f7b84ea00787d5789f2
    Reviewed-on: https://go-review.googlesource.com/37892
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d3f5e3691712a2098fc8171574748389cec931a1
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sun Mar 5 19:57:26 2017 +1100

    cmd/link: use IMAGE_SYM_CLASS_STATIC for local symbols
    
    Sometimes asm code in 2 different packages name its global
    symbols with the same name. When these symbols are passed
    to gcc, it refuses to link them thinking they are duplicate.
    Mark these symbols with IMAGE_SYM_CLASS_STATIC.
    
    Fixes #19198.
    
    Change-Id: Ia5f59ede47354a2b48ce60b7d406c9f097ff2000
    Reviewed-on: https://go-review.googlesource.com/37810
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6491496d10c8e7e62f875c7781a2887564976b89
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Sat Mar 4 19:47:19 2017 +0900

    net/mail: ignore whitespace between adjacent 'encoded-word's
    
    rfc2047 says:
      White space between adjacent 'encoded-word's is not displayed.
    
    Although, mime package already have that feature,
    we cannot simply reuse that code,
    because there is a subtle difference in quoted-string handling.
    
    Fixes #19363
    
    Change-Id: I754201aa3c6b701074ad78fe46818af5b96cbd00
    Reviewed-on: https://go-review.googlesource.com/37811
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9207a7437e382e4a05436fc26a2e24740a0492ec
Author: Josselin Costanzi <josselin@costanzi.fr>
Date:   Sun Mar 5 18:55:18 2017 +0100

    encoding/base64: add alphabet and padding restrictions
    
    Document and check that the alphabet cannot contain '\n' or '\r'.
    Document that the alphabet cannot contain the padding character.
    Document that the padding character must be equal or bellow '\xff'.
    Document that the padding character must not be '\n' or '\r'.
    
    Fixes #19343
    Fixes #19318
    
    Change-Id: I6de0034d347ffdf317d7ea55d6fe38b01c2c4c03
    Reviewed-on: https://go-review.googlesource.com/37838
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8d3f29577d95aa06b2653d20e331aa47f759db06
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Mar 6 19:43:21 2017 -0500

    cmd/compile: regenerate knownFormats
    
    Should fix the build dashboard.
    
    Change-Id: Id4c8a996d9f689e1fa865a9cff9a7f52c700c691
    Reviewed-on: https://go-review.googlesource.com/37877
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 29be58039577a69bfe239a8aaf2cf26d8debf566
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Thu Feb 2 11:48:22 2017 -0800

    doc: add link to the setting GOPATH guide
    
    Change-Id: I4718c82540ef214728393824b89c8c7f6656823b
    Reviewed-on: https://go-review.googlesource.com/36210
    Reviewed-by: Russ Cox <rsc@golang.org>

commit cf710949a98aaac48a3d97660f7f6bb7d14f1ad7
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Mar 6 15:41:56 2017 -0800

    Revert "cmd/compile: improve error message if init is directly invoked"
    
    This reverts commit cb6e0639fb090ea0e129b1ddb956a7e645cff285.
    
    The fix is incorrect as it's perfectly fine to refer to an
    identifier 'init' inside a function, and 'init' may even be
    a variable of function value. Misspelling 'init' in that
    context would lead to an incorrect error message.
    
    Reopened #8481.
    
    Change-Id: I49787fdf7738213370ae6f0cab54013e9e3394a8
    Reviewed-on: https://go-review.googlesource.com/37876
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a6bd42f2630aa8907009a1ec0c35048c20a7afae
Author: philhofer <phofer@umich.edu>
Date:   Sat Mar 4 16:17:12 2017 -0800

    cmd/compile: emit OffPtr for first field in SSA'd structs
    
    Given
    
      (Store [c] (OffPtr <T1> [0] (Addr <T> _)) _
        (Store [c] (Addr <T> _) _ _))
    
    dead store elimination doesn't eliminate the inner
    Store, because it addresses a type of a different width
    than the first store.
    
    When decomposing StructMake operations, always generate
    an OffPtr to address struct fields so that dead stores to
    the first field of the struct can be optimized away.
    
    benchmarks affected on darwin/amd64:
    HTTPClientServer-8        73.2µs ± 1%    72.7µs ± 1%  -0.69%  (p=0.022 n=9+10)
    TimeParse-8                304ns ± 1%     300ns ± 0%  -1.61%  (p=0.000 n=9+9)
    RegexpMatchEasy1_32-8     80.1ns ± 0%    79.5ns ± 1%  -0.84%  (p=0.000 n=8+9)
    GobDecode-8               6.78ms ± 0%    6.81ms ± 1%  +0.46%  (p=0.000 n=9+10)
    Gunzip-8                  36.1ms ± 1%    36.2ms ± 0%  +0.37%  (p=0.019 n=10+10)
    JSONEncode-8              15.6ms ± 0%    15.7ms ± 0%  +0.69%  (p=0.000 n=9+10)
    
    Change-Id: Ia80d73fd047f9400c616ca64fdee4f438a0e7f21
    Reviewed-on: https://go-review.googlesource.com/37769
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit cd6f18779fd0c8be723d6eb1f1891796bfe98aa3
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Mon Mar 6 10:54:27 2017 +0100

    net/http: remove unused ResponseWriter params
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: I66f5a191cf9c9a11a7c3c4d7ee0a02e2c185f019
    Reviewed-on: https://go-review.googlesource.com/37841
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit d10b50dc3447ec69563320b0538b7a1b1f4cfc81
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Mar 6 15:30:32 2017 -0800

    cmd/compile/internal/syntax: print position info for names in tree dump
    
    Debugging support.
    
    Change-Id: Ia518aaed36eaba76e6233306f718ad8aff3ce744
    Reviewed-on: https://go-review.googlesource.com/37875
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 1874d4a883805056727e7c2fec01dbb7bf30fc6e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 3 16:55:05 2017 -0800

    cmd/internal/obj, cmd/compile: rip off some toolstash bandaids
    
    Change-Id: I402383e893223facae451adbd640113126d5edd9
    Reviewed-on: https://go-review.googlesource.com/37873
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9cb2ee0ff2efb3ec281718c6e301db695f9a0870
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 3 16:45:21 2017 -0800

    cmd/internal/obj: move STEXT-only LSym fields into new FuncInfo struct
    
    Shrinks LSym somewhat for non-STEXT LSyms, which are much more common.
    
    While here, switch to tracking Automs in a slice instead of a linked
    list. (Previously, this would have made LSyms larger.)
    
    Passes toolstash-check.
    
    Change-Id: I082e50e1d1f1b544c9e06b6e412a186be6a4a2b5
    Reviewed-on: https://go-review.googlesource.com/37872
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 7a98bdf1c258dfa6aa539035f422389dc4c994f0
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 3 14:27:24 2017 -0800

    cmd/internal/obj: remove AUSEFIELD pseudo-op
    
    Instead, cmd/compile can directly emit R_USEFIELD relocations.
    
    Manually verified rsc.io/tmp/fieldtrack still passes.
    
    Change-Id: Ib1fb5ab902ff0ad17ef6a862a9a5692caf7f87d1
    Reviewed-on: https://go-review.googlesource.com/37871
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 5d0c20efc7bb373107535543f75741465fe93d3f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 6 12:57:44 2017 -0800

    cmd/compile: preserve Type.nod in copytype
    
    By clearing out t.nod in copytype, we effectively lose the reference
    from a Type back to its declaring OTYPE Node. This means later in
    typenamesym when we add typenod(t) to signatlist, we end up creating a
    new OTYPE Node. Moreover, this Node's position information will depend
    on whatever context it happens be needed, and will be used for the
    Type's position in the export data.
    
    Updates #19391.
    
    Change-Id: Ied93126449f75d7c5e3275cbdcc6fa657a8aa21d
    Reviewed-on: https://go-review.googlesource.com/37870
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 4b261a141006667b687978a61aedb7222d6be0e4
Author: Quentin Smith <quentin@golang.org>
Date:   Mon Mar 6 11:45:19 2017 -0500

    test/fixedbugs: add test for #19403
    
    Change-Id: Ie52dac8eb4daed95e049ad74d5ae101e8a5cb854
    Reviewed-on: https://go-review.googlesource.com/37725
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit e4f73769bcb06fca7b22d4ac3ec17c7247e84f39
Author: Austin Clements <austin@google.com>
Date:   Mon Mar 6 14:13:24 2017 -0500

    runtime: strongly encourage CallersFrames with the result of Callers
    
    For historical reasons, it's still commonplace to iterate over the
    slice returned by runtime.Callers and call FuncForPC on each PC. This
    is broken in gccgo and somewhat broken in gc and will become more
    broken in gc with mid-stack inlining.
    
    In Go 1.7, we introduced runtime.CallersFrames to deal with these
    problems, but didn't strongly direct people toward using it. Reword
    the documentation on runtime.Callers to more strongly encourage people
    to use CallersFrames and explicitly discourage them from iterating
    over the PCs or using FuncForPC on the results.
    
    Fixes #19426.
    
    Change-Id: Id0d14cb51a0e9521c8fdde9612610f2c2b9383c4
    Reviewed-on: https://go-review.googlesource.com/37726
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 1fa063cbb6efd16a13d9c0e44374aac8791fbcb7
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sun Mar 5 19:10:47 2017 +0100

    go/internal/gccimporter: actually use pkg parameter
    
    We're passed a pkg, so it makes little sense to not use it. This was
    probably a typo and not the intended behaviour.
    
    Fixes #19407.
    
    Change-Id: Ia1c9130c0e474daf47753cf51914a2d7db272c96
    Reviewed-on: https://go-review.googlesource.com/37839
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 61d9cd73fb3f3748be84320a4e78841c445a38f3
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Mar 3 17:15:11 2017 -0800

    internal/poll: only start Windows goroutines when we need them
    
    We don't need to start the goroutines if the program isn't going to do
    any I/O.
    
    Fixes #19390.
    
    Change-Id: I47eef992d3ad05ed5f3150f4d6e5b3e0cb16a551
    Reviewed-on: https://go-review.googlesource.com/37762
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 5ce06cf71d62e6fc1740d97b4ff4dda7e039c606
Author: Josselin Costanzi <josselin@costanzi.fr>
Date:   Sun Mar 5 18:04:30 2017 +0100

    encoding/base64: fix decode reports incorrect index
    
    Fix Decode to return the correct illegal data index from a corrupted
    input that contains whitespaces.
    
    Fixes #19406
    
    Change-Id: Ib2b2b6ed7e41f024d0da2bd035caec4317c2869c
    Reviewed-on: https://go-review.googlesource.com/37837
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0efc8b21881ab35fdb45547088b1935fc8ebf263
Author: Austin Clements <austin@google.com>
Date:   Mon Feb 20 22:37:07 2017 -0500

    runtime: avoid repeated findmoduledatap calls
    
    Currently almost every function that deals with a *_func has to first
    look up the *moduledata for the module containing the function's entry
    point. This means we almost always do at least two identical module
    lookups whenever we deal with a *_func (one to get the *_func and
    another to get something from its module data) and sometimes several
    more.
    
    Fix this by making findfunc return a new funcInfo type that embeds
    *_func, but also includes the *moduledata, and making all of the
    functions that currently take a *_func instead take a funcInfo and use
    the already-found *moduledata.
    
    This transformation is trivial for the most part, since the *_func
    type is usually inferred. The annoying part is that we can no longer
    use nil to indicate failure, so this introduces a funcInfo.valid()
    method and replaces nil checks with calls to valid.
    
    Change-Id: I9b8075ef1c31185c1943596d96dec45c7ab5100f
    Reviewed-on: https://go-review.googlesource.com/37331
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 6533cc1ce899fa3c7fac1a85ad724e333fb9710f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 6 10:14:15 2017 -0800

    cmd/internal/goobj: update to support go19ld
    
    Updates the disassembler to support the same object file version used
    by the assembler and linker.
    
    Related #14782.
    
    Change-Id: I4cd7560c4e4e1350cfb27ca9cbe0fde25fe693cc
    Reviewed-on: https://go-review.googlesource.com/37797
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 7c84dc79fdd1e26fc117e170ada81444694e6bed
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Mar 6 10:51:37 2017 -0800

    cmd/internal/obj, cmd/link: bump magic string to go19ld
    
    golang.org/cl/37231 changed the object file format, but forgot to bump
    the version string.
    
    Change-Id: I8351ec8ed55e65479006e7c0df20254d0e31015f
    Reviewed-on: https://go-review.googlesource.com/37798
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ad74e450ca94b3d9e37efd47c234ca1eeee2889d
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Mon Mar 6 09:59:36 2017 +0100

    regexp/syntax: remove unused flags parameter
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: I186d2afd067e97eb05d65c4599119b347f82867d
    Reviewed-on: https://go-review.googlesource.com/37840
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d8485ee2e7c86def7318ce72b6a3714941452195
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 3 13:42:41 2017 -0800

    go/internal/gcimporter: return (possibly incomplete) package in case of error
    
    For #16088.
    
    Change-Id: Ib38bda06a5c5d110ca86510043775c5cf229e6a8
    Reviewed-on: https://go-review.googlesource.com/37756
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 63f8cca95bbfb48a8b0d96f00a4a25dfa4ec6bb8
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 3 13:07:54 2017 -0800

    go/internal/srcimporter: return (possibly incomplete) package in case of error
    
    For #16088.
    
    Change-Id: I0ff480e95ef5af375be2ccc655f8b233a7bcd39d
    Reviewed-on: https://go-review.googlesource.com/37755
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 2ad7453bf42654a6a1615fa1cc867123570c4595
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 2 14:06:35 2017 -0800

    go/types: continue type-checking with fake packages if imports failed
    
    This will make type-checking more robust in the presence of import errors.
    
    Also:
    - import is now relative to directory containing teh file containing the import
      (matters for relative imports)
    - factored out package import code from main resolver loop
    - fixed a couple of minor bugs
    
    Fixes #16088.
    
    Change-Id: I1ace45c13cd0fa675d1762877cec0a30afd9ecdc
    Reviewed-on: https://go-review.googlesource.com/37697
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2ef88f7fcf5cebccdf9a20a5af6a374864ccf149
Author: Austin Clements <austin@google.com>
Date:   Sat Dec 17 22:07:27 2016 -0500

    runtime: lock-free fast path for mark bits allocation
    
    Currently we acquire a global lock for every newMarkBits call. This is
    unfortunate since every span sweep operation calls newMarkBits.
    
    However, most allocations are simply linear allocations from the
    current arena. Take advantage of this to add a lock-free fast path for
    allocating from the current arena. With this change, the global lock
    only protects the lists of arenas, not the free offset in the current
    arena.
    
    Change-Id: I6cf6182af8492c8bfc21276114c77275fe3d7826
    Reviewed-on: https://go-review.googlesource.com/34595
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 6c4a8d195b627bf216c6cd22e8237c6faf99cad5
Author: Austin Clements <austin@google.com>
Date:   Fri Dec 16 15:56:13 2016 -0500

    runtime: don't hold global gcBitsArenas lock over allocation
    
    Currently, newArena holds the gcBitsArenas lock across allocating
    memory from the OS for a new gcBits arena. This is a global lock and
    allocating physical memory can be expensive, so this has the potential
    to cause high lock contention, especially since every single span
    sweep operation calls newArena (via newMarkBits).
    
    Improve the situation by temporarily dropping the lock across
    allocation. This means the caller now has to revalidate its
    assumptions after the lock is dropped, so this also factors out that
    code path and reinvokes it after the lock is acquired.
    
    Change-Id: I1113200a954ab4aad16b5071512583cfac744bdc
    Reviewed-on: https://go-review.googlesource.com/34594
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 4e428907c5d34b31e5d21c17917f70b0d1f0e4f6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 3 16:15:35 2017 -0800

    cmd/compile: avoid generating some dead blocks
    
    We generate a lot of pointless dead blocks
    during the AST to SSA conversion.
    There are a few commonly occurring kinds
    of statements that contain neither variables
    nor code and that switch to a new block themselves.
    Stop making dead blocks for them.
    
    For the code in #19379, this reduces compilation
    wall time by 36% and max rss by 28%.
    
    This also helps a little for regular code,
    particularly code heavy on switch statements.
    
    name       old time/op      new time/op      delta
    Template        231ms ± 3%       230ms ± 5%    ~     (p=0.402 n=17+16)
    Unicode         101ms ± 4%       103ms ± 5%    ~     (p=0.221 n=19+18)
    GoTypes         635ms ± 5%       625ms ± 4%    ~     (p=0.063 n=20+18)
    Compiler        2.93s ± 2%       2.89s ± 2%  -1.22%  (p=0.003 n=20+19)
    SSA             4.53s ± 3%       4.52s ± 3%    ~     (p=0.380 n=20+19)
    Flate           132ms ± 4%       133ms ± 5%    ~     (p=0.647 n=20+19)
    GoParser        161ms ± 3%       161ms ± 4%    ~     (p=0.749 n=20+19)
    Reflect         403ms ± 4%       397ms ± 3%  -1.53%  (p=0.030 n=20+19)
    Tar             121ms ± 2%       121ms ± 8%    ~     (p=0.544 n=19+19)
    XML             225ms ± 3%       224ms ± 4%    ~     (p=0.396 n=20+19)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   302user-ms ± 1%  297user-ms ± 7%  -1.49%  (p=0.048 n=15+18)
    Unicode    142user-ms ± 3%  143user-ms ± 5%    ~     (p=0.363 n=19+17)
    GoTypes    852user-ms ± 5%  851user-ms ± 3%    ~     (p=0.851 n=20+18)
    Compiler   4.11user-s ± 6%  3.98user-s ± 3%  -3.08%  (p=0.000 n=20+19)
    SSA        6.91user-s ± 5%  6.82user-s ± 7%    ~     (p=0.113 n=20+19)
    Flate      164user-ms ± 4%  168user-ms ± 4%  +2.42%  (p=0.001 n=18+19)
    GoParser   207user-ms ± 4%  206user-ms ± 4%    ~     (p=0.176 n=20+18)
    Reflect    509user-ms ± 4%  505user-ms ± 4%    ~     (p=0.113 n=20+19)
    Tar        153user-ms ± 7%  151user-ms ± 9%    ~     (p=0.283 n=20+19)
    XML        284user-ms ± 4%  282user-ms ± 4%    ~     (p=0.270 n=20+19)
    
    name       old alloc/op     new alloc/op     delta
    Template       42.6MB ± 0%      41.9MB ± 0%  -1.55%  (p=0.000 n=19+19)
    Unicode        31.7MB ± 0%      31.7MB ± 0%    ~     (p=0.828 n=20+18)
    GoTypes         124MB ± 0%       121MB ± 0%  -2.11%  (p=0.000 n=20+17)
    Compiler        534MB ± 0%       523MB ± 0%  -2.06%  (p=0.000 n=20+19)
    SSA             989MB ± 0%       977MB ± 0%  -1.28%  (p=0.000 n=20+19)
    Flate          27.8MB ± 0%      27.5MB ± 0%  -0.98%  (p=0.000 n=20+19)
    GoParser       34.3MB ± 0%      34.0MB ± 0%  -0.81%  (p=0.000 n=20+19)
    Reflect        84.6MB ± 0%      82.9MB ± 0%  -2.00%  (p=0.000 n=17+18)
    Tar            28.8MB ± 0%      28.3MB ± 0%  -1.52%  (p=0.000 n=16+18)
    XML            47.2MB ± 0%      45.8MB ± 0%  -2.99%  (p=0.000 n=20+19)
    
    name       old allocs/op    new allocs/op    delta
    Template         421k ± 1%        419k ± 1%  -0.41%  (p=0.001 n=20+19)
    Unicode          338k ± 1%        338k ± 1%    ~     (p=0.478 n=20+19)
    GoTypes         1.28M ± 0%       1.28M ± 0%  -0.36%  (p=0.000 n=20+18)
    Compiler        5.06M ± 0%       5.03M ± 0%  -0.63%  (p=0.000 n=20+19)
    SSA             9.14M ± 0%       9.11M ± 0%  -0.34%  (p=0.000 n=20+19)
    Flate            267k ± 1%        266k ± 1%    ~     (p=0.149 n=20+19)
    GoParser         347k ± 0%        347k ± 1%    ~     (p=0.103 n=19+19)
    Reflect         1.07M ± 0%       1.07M ± 0%  -0.42%  (p=0.000 n=16+18)
    Tar              274k ± 0%        273k ± 1%    ~     (p=0.116 n=19+19)
    XML              449k ± 0%        446k ± 1%  -0.60%  (p=0.000 n=20+19)
    
    Updates #19379
    
    Change-Id: Ie798c347a0c081f5e349e1529880bebaae290967
    Reviewed-on: https://go-review.googlesource.com/37760
    Reviewed-by: David Chase <drchase@google.com>

commit a5a1fd4bc9c45f817c15733f2576048546325b28
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Mon Mar 6 20:00:54 2017 +0200

    cmd/compile/internal/gc: convert Sym.Flags to bitset8
    
    This makes Sym flags consistent with the rest of the code after
    the CL 37445.
    
    No functional changes.
    
    Passes toolstash -cmp.
    
    Change-Id: Ica919f2ab98581371c717fff9a70aeb11058ca17
    Reviewed-on: https://go-review.googlesource.com/37847
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0332b6cf589a51f421a88176f756f0315c9f1491
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sun Mar 5 14:04:56 2017 +0100

    encoding/gob: remove unused ut and atyp parameters
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: Ic97f05a2ecb5b17caa36aafe403e2266abea3e0e
    Reviewed-on: https://go-review.googlesource.com/37836
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 06a6b3a4134f1c7b93308a4f45eed6aeaa9a5f00
Author: Russ Cox <rsc@golang.org>
Date:   Thu Mar 2 09:24:42 2017 -0500

    test/locklinear: deflake again
    
    On overloaded machines once we get to big N, the machine slowness dominates.
    But we only retry once we get to a big N.
    Instead, retry for small N too, and die on the first big N that fails.
    
    Change-Id: I3ab9cfb88832ad86e2ba1389a926045091268aeb
    Reviewed-on: https://go-review.googlesource.com/37543
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2ec77d3457aaa9e07ac5c765a0323fc9c3ef889f
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sun Mar 5 13:54:16 2017 +0100

    go/doc: remove unused tok parameter
    
    Found via github.com/mvdan/unparam.
    
    Change-Id: I12cb0c35b14c880425c347fb3eb146712a86f310
    Reviewed-on: https://go-review.googlesource.com/37834
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9600c32cc516b66f5337bebc61e249af4638c106
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sun Mar 5 13:45:10 2017 +0100

    go/printer: remove unused comment parameter
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: I5b0c7cfdc1ab4fe0d79ef4c5a31612bbcf2ff3ad
    Reviewed-on: https://go-review.googlesource.com/37833
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5f3281139d2173f4d742aaf3ed9ba57a93b46737
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sun Mar 5 13:42:55 2017 +0100

    go/types: remove unused field parameter
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: Ie26e963176eb7afb35d16fed5cbca6530f7731c3
    Reviewed-on: https://go-review.googlesource.com/37832
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 789c5255a4ee2bca8906baa15fc3e400054ff44d
Author: Eitan Adler <lists@eitanadler.com>
Date:   Sun Mar 5 09:14:38 2017 -0800

    all: remove the the duplicate words
    
    Change-Id: I6343c162e27e2e492547c96f1fc504909b1c03c0
    Reviewed-on: https://go-review.googlesource.com/37793
    Reviewed-by: Daniel Martí <mvdan@mvdan.cc>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 694f9e36aa508116b8fd0bf2c42e680f2937dd56
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sun Mar 5 13:56:30 2017 +0100

    encoding/xml: remove unused start parameter
    
    Found by github.com/mvdan/unparam.
    
    Change-Id: I5a6664cceeba1cf1c2f3236ddf4db5ce7a64b02a
    Reviewed-on: https://go-review.googlesource.com/37835
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit de3669901a6c9551067071ab410775e57e3b26c2
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sun Mar 5 13:37:33 2017 +0100

    strconv: remove unused append rune width param
    
    Found by github.com/mvdan/unparam. Small performance win when the
    utf8.RuneLen call is removed.
    
    name               old time/op    new time/op    delta
    AppendQuoteRune-4    21.7ns ± 0%    21.4ns ± 0%  -1.38%  (p=0.008 n=5+5)
    
    Change-Id: Ieb3b3e1148db7a3d854c81555a491edeff549f43
    Reviewed-on: https://go-review.googlesource.com/37831
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6fd5e2549ad530dc6f8504de61f1e49fa4f336ea
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Mar 5 11:37:09 2017 -0500

    cmd/compile: mark MOVWF/MOVFW clobbering F15 on ARM
    
    The assembler back end uses F15 as a temporary register in these
    instructions.
    
    Checked the assembler back end and made sure that this is the
    only case clobbering F15.
    
    Fixes #19403.
    
    Change-Id: I02b9e00fdd9229db899f501c8e9b306e02912d83
    Reviewed-on: https://go-review.googlesource.com/37792
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d4451362c0fa47b25fc2e69129e68cbbee4a6bdf
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 4 16:28:59 2017 -0800

    runtime: add slicebytetostring benchmark
    
    Change-Id: I666d2c6ea8d0b54a71260809d1a2573b122865b2
    Reviewed-on: https://go-review.googlesource.com/37790
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 57e038615d945b610f4b62c40ddeb1fd40130649
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Mar 4 07:09:33 2017 -0800

    cmd/internal/src: cache prefixed filenames
    
    CL 37234 introduced string concatenation into some hot code.
    This CL does that work earlier and caches the result.
    
    Updates #19386
    
    Performance impact vs master:
    
    name       old time/op      new time/op      delta
    Template        223ms ± 5%       216ms ± 5%   -2.98%  (p=0.001 n=20+20)
    Unicode        98.7ms ± 4%      99.0ms ± 4%     ~     (p=0.749 n=20+19)
    GoTypes         631ms ± 4%       626ms ± 4%     ~     (p=0.253 n=20+20)
    Compiler        2.91s ± 1%       2.87s ± 3%   -1.11%  (p=0.005 n=18+20)
    SSA             4.48s ± 2%       4.36s ± 2%   -2.77%  (p=0.000 n=20+20)
    Flate           130ms ± 2%       129ms ± 6%     ~     (p=0.428 n=19+20)
    GoParser        160ms ± 4%       157ms ± 3%   -1.62%  (p=0.005 n=20+18)
    Reflect         395ms ± 2%       394ms ± 4%     ~     (p=0.445 n=20+20)
    Tar             120ms ± 5%       118ms ± 6%     ~     (p=0.101 n=19+20)
    XML             224ms ± 3%       223ms ± 3%     ~     (p=0.544 n=19+19)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   291user-ms ± 5%  265user-ms ± 5%   -9.02%  (p=0.000 n=20+19)
    Unicode    140user-ms ± 3%  139user-ms ± 8%     ~     (p=0.904 n=20+20)
    GoTypes    844user-ms ± 3%  849user-ms ± 3%     ~     (p=0.251 n=20+18)
    Compiler   4.06user-s ± 5%  3.98user-s ± 2%     ~     (p=0.056 n=20+20)
    SSA        6.89user-s ± 5%  6.50user-s ± 3%   -5.61%  (p=0.000 n=20+20)
    Flate      164user-ms ± 5%  163user-ms ± 4%     ~     (p=0.365 n=20+19)
    GoParser   206user-ms ± 6%  204user-ms ± 4%     ~     (p=0.534 n=20+18)
    Reflect    501user-ms ± 4%  505user-ms ± 5%     ~     (p=0.383 n=20+20)
    Tar        151user-ms ± 3%  152user-ms ± 7%     ~     (p=0.798 n=17+20)
    XML        283user-ms ± 7%  280user-ms ± 5%     ~     (p=0.301 n=20+20)
    
    name       old alloc/op     new alloc/op     delta
    Template       42.5MB ± 0%      40.2MB ± 0%   -5.59%  (p=0.000 n=20+20)
    Unicode        31.7MB ± 0%      31.0MB ± 0%   -2.19%  (p=0.000 n=20+18)
    GoTypes         124MB ± 0%       117MB ± 0%   -5.90%  (p=0.000 n=20+20)
    Compiler        533MB ± 0%       490MB ± 0%   -8.07%  (p=0.000 n=20+20)
    SSA             989MB ± 0%       893MB ± 0%   -9.74%  (p=0.000 n=20+20)
    Flate          27.8MB ± 0%      26.1MB ± 0%   -5.92%  (p=0.000 n=20+20)
    GoParser       34.3MB ± 0%      32.1MB ± 0%   -6.43%  (p=0.000 n=19+20)
    Reflect        84.6MB ± 0%      81.4MB ± 0%   -3.84%  (p=0.000 n=20+20)
    Tar            28.8MB ± 0%      27.7MB ± 0%   -3.89%  (p=0.000 n=20+20)
    XML            47.2MB ± 0%      44.2MB ± 0%   -6.45%  (p=0.000 n=20+19)
    
    name       old allocs/op    new allocs/op    delta
    Template         420k ± 1%        381k ± 1%   -9.35%  (p=0.000 n=20+20)
    Unicode          338k ± 1%        324k ± 1%   -4.29%  (p=0.000 n=20+19)
    GoTypes         1.28M ± 0%       1.15M ± 0%  -10.30%  (p=0.000 n=20+20)
    Compiler        5.06M ± 0%       4.41M ± 0%  -12.92%  (p=0.000 n=20+20)
    SSA             9.14M ± 0%       7.91M ± 0%  -13.46%  (p=0.000 n=19+20)
    Flate            267k ± 0%        241k ± 1%   -9.53%  (p=0.000 n=20+20)
    GoParser         347k ± 1%        312k ± 0%  -10.15%  (p=0.000 n=19+20)
    Reflect         1.07M ± 0%       1.00M ± 0%   -6.86%  (p=0.000 n=20+20)
    Tar              274k ± 1%        256k ± 1%   -6.73%  (p=0.000 n=20+20)
    XML              448k ± 0%        398k ± 0%  -11.17%  (p=0.000 n=20+18)
    
    
    Performance impact when applied together with CL 37234
    atop CL 37234's parent commit (i.e. as if it were
    a part of CL 37234), to show that this commit
    makes CL 37234 completely performance-neutral:
    
    name       old time/op      new time/op      delta
    Template        222ms ±14%       222ms ±14%    ~     (p=1.000 n=14+15)
    Unicode         104ms ±18%       106ms ±18%    ~     (p=0.650 n=13+14)
    GoTypes         653ms ± 7%       638ms ± 5%    ~     (p=0.145 n=14+12)
    Compiler        3.10s ± 1%       3.13s ±10%    ~     (p=1.000 n=2+2)
    SSA             4.73s ±11%       4.68s ±11%    ~     (p=0.567 n=15+15)
    Flate           136ms ± 4%       133ms ± 7%    ~     (p=0.231 n=12+14)
    GoParser        163ms ±11%       169ms ±10%    ~     (p=0.352 n=14+14)
    Reflect         415ms ±15%       423ms ±20%    ~     (p=0.715 n=15+14)
    Tar             133ms ±17%       130ms ±23%    ~     (p=0.252 n=14+15)
    XML             236ms ±16%       235ms ±14%    ~     (p=0.874 n=14+14)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   271user-ms ±10%  271user-ms ±10%    ~     (p=0.780 n=14+15)
    Unicode    143user-ms ± 5%  146user-ms ±11%    ~     (p=0.432 n=12+14)
    GoTypes    864user-ms ± 5%  866user-ms ± 9%    ~     (p=0.905 n=14+13)
    Compiler   4.17user-s ± 1%  4.26user-s ± 7%    ~     (p=1.000 n=2+2)
    SSA        6.79user-s ± 8%  6.79user-s ± 6%    ~     (p=0.902 n=15+15)
    Flate      169user-ms ± 8%  164user-ms ± 5%  -3.13%  (p=0.014 n=14+14)
    GoParser   212user-ms ± 7%  217user-ms ±22%    ~     (p=1.000 n=13+15)
    Reflect    521user-ms ± 7%  533user-ms ±15%    ~     (p=0.511 n=14+14)
    Tar        165user-ms ±17%  161user-ms ±15%    ~     (p=0.345 n=15+15)
    XML        294user-ms ±11%  292user-ms ±10%    ~     (p=0.839 n=14+14)
    
    name       old alloc/op     new alloc/op     delta
    Template       39.9MB ± 0%      39.9MB ± 0%    ~     (p=0.621 n=15+14)
    Unicode        31.0MB ± 0%      31.0MB ± 0%    ~     (p=0.098 n=13+15)
    GoTypes         117MB ± 0%       117MB ± 0%    ~     (p=0.775 n=15+15)
    Compiler        488MB ± 0%       488MB ± 0%    ~     (p=0.333 n=2+2)
    SSA             892MB ± 0%       892MB ± 0%  +0.03%  (p=0.000 n=15+15)
    Flate          26.1MB ± 0%      26.1MB ± 0%    ~     (p=0.098 n=15+15)
    GoParser       31.8MB ± 0%      31.8MB ± 0%    ~     (p=0.525 n=15+13)
    Reflect        81.2MB ± 0%      81.2MB ± 0%  +0.06%  (p=0.001 n=12+14)
    Tar            27.5MB ± 0%      27.5MB ± 0%    ~     (p=0.595 n=15+15)
    XML            44.1MB ± 0%      44.1MB ± 0%    ~     (p=0.486 n=15+15)
    
    name       old allocs/op    new allocs/op    delta
    Template         378k ± 1%        378k ± 0%    ~     (p=0.949 n=15+14)
    Unicode          324k ± 0%        324k ± 1%    ~     (p=0.057 n=14+15)
    GoTypes         1.15M ± 0%       1.15M ± 0%    ~     (p=0.461 n=15+15)
    Compiler        4.39M ± 0%       4.39M ± 0%    ~     (p=0.333 n=2+2)
    SSA             7.90M ± 0%       7.90M ± 0%  +0.06%  (p=0.008 n=15+15)
    Flate            240k ± 1%        241k ± 0%    ~     (p=0.233 n=15+15)
    GoParser         309k ± 1%        309k ± 0%    ~     (p=0.867 n=15+12)
    Reflect         1.00M ± 0%       1.00M ± 0%    ~     (p=0.139 n=12+15)
    Tar              254k ± 1%        253k ± 1%    ~     (p=0.345 n=15+15)
    XML              398k ± 0%        397k ± 1%    ~     (p=0.267 n=15+15)
    
    
    Change-Id: Ic999a0f456a371c99eebba0f9747263a13836e33
    Reviewed-on: https://go-review.googlesource.com/37766
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c2eb06193f70475650a85a0e279f865181c1cece
Author: Kevin Burke <kev@inburke.com>
Date:   Wed Mar 1 10:31:57 2017 -0800

    os/user: add non-cgo versions of Lookup, LookupId
    
    If you cross compile for a Unix target and call user.Lookup("root")
    or user.LookupId("0"), we'll try to read the answer out of
    /etc/passwd instead of returning an "unimplemented" error.
    
    The equivalent cgo function calls getpwuid_r in glibc, which
    may reach out to the NSS database or allow callers to register
    extensions. The pure Go implementation only reads from /etc/passwd.
    
    Change-Id: I56a302d634b15ba5097f9f0d6a758c68e486ba6d
    Reviewed-on: https://go-review.googlesource.com/37664
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4fc45ae8794eaa84505629e82997b90503c89aa2
Author: Giovanni Bajo <rasky@develer.com>
Date:   Sun Feb 26 02:54:32 2017 +0100

    cmd/compile: improve generic rules for BCE based on AND operations.
    
    Match more patterns generated by the compiler where the index for
    a bound check is bounded through a AND operation, with different
    register sizes.
    
    These rules trigger a dozen of times in a bootstrap.
    
    Change-Id: Ic9fff16f21d08580f19a366c3ee1a372e58357d1
    Reviewed-on: https://go-review.googlesource.com/37442
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 4a7cf960c38d72e9f0c6f00e46e013be2a35d56e
Author: Austin Clements <austin@google.com>
Date:   Tue Jan 3 10:15:55 2017 -0700

    runtime: make ReadMemStats STW for < 25µs
    
    Currently ReadMemStats stops the world for ~1.7 ms/GB of heap because
    it collects statistics from every single span. For large heaps, this
    can be quite costly. This is particularly unfortunate because many
    production infrastructures call this function regularly to collect and
    report statistics.
    
    Fix this by tracking the necessary cumulative statistics in the
    mcaches. ReadMemStats still has to stop the world to stabilize these
    statistics, but there are only O(GOMAXPROCS) mcaches to collect
    statistics from, so this pause is only 25µs even at GOMAXPROCS=100.
    
    Fixes #13613.
    
    Change-Id: I3c0a4e14833f4760dab675efc1916e73b4c0032a
    Reviewed-on: https://go-review.googlesource.com/34937
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 3399fd254dcdf4e8a9be8c327076de5f9efe1b3a
Author: Austin Clements <austin@google.com>
Date:   Thu Dec 22 17:45:55 2016 -0700

    runtime: remove unused gcstats
    
    The gcstats structure is no longer consumed by anything and no longer
    tracks statistics that are particularly relevant to the concurrent
    garbage collector. Remove it. (Having statistics is probably a good
    idea, but these aren't the stats we need these days and we don't have
    a way to get them out of the runtime.)
    
    In preparation for #13613.
    
    Change-Id: Ib63e2f9067850668f9dcbfd4ed89aab4a6622c3f
    Reviewed-on: https://go-review.googlesource.com/34936
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 7523baed098789b8b4e2a5b295b14db91ab5e5cf
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Feb 1 16:04:07 2017 +0100

    misc/ios,cmd/go, runtime/cgo: fix iOS test harness (again)
    
    The iOS test harness was recently changed in response to lldb bugs
    to replace breakpoints with the SIGUSR2 signal (CL 34926), and to
    pass the current directory in the test binary arguments (CL 35152).
    Both the signal sending and working directory setup is done from
    the go test driver.
    
    However, the new method doesn't work with tests where a C program is
    the test driver instead of go test: the current working directory
    will not be changed and SIGUSR2 is not raised.
    
    Instead of copying that logic into any C test program, rework the
    test harness (again) to move the setup logic to the early runtime
    cgo setup code. That way, the harness will run even in the library
    build modes.
    
    Then, use the app Info.plist file to pass the working
    directory, removing the need to alter the arguments after running.
    
    Finally, use the SIGINT signal instead of SIGUSR2 to avoid
    manipulating the signal masks or handlers.
    
    Fixes the testcarchive tests on iOS.
    
    With this CL, both darwin/arm and darwin/arm64 passes all.bash.
    
    This CL replaces CL 34926, CL 35152 as well as the fixup CL
    35123 and CL 35255. They are reverted in CLs earlier in the
    relation chain.
    
    Change-Id: I8485c7db1404fbd8daa261efd1ea89e905121a3e
    Reviewed-on: https://go-review.googlesource.com/36090
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit fec40bd106f30d95939ec7ca6066dd3cd90c0af1
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sat Mar 4 01:44:45 2017 +0100

    Revert "cmd/go, misc: switch from breakpoint to SIGUSR2"
    
    This reverts commit 333f764df3d78930a5a3097fc34ac1374b7c3187.
    
    Replaced by a improved strategy later in the CL relation chain.
    
    Change-Id: I70a1d2f0aa5aa0d3d0ec85b5a956c6fb60d88908
    Reviewed-on: https://go-review.googlesource.com/36069
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 03e2a4d1f1d2ee74799eb103d0051c129523e4d6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Mar 3 16:33:06 2017 -0800

    cmd/compile: cull UINF
    
    It was used with Node.Ullman, which is now gone.
    
    Change-Id: I83b167645659ae7ef70043b7915d642e42ca524f
    Reviewed-on: https://go-review.googlesource.com/37761
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e646d07329e7edc56c8cf3284c729937979dae4a
Author: Dmitri Shuralyov <shurcooL@gmail.com>
Date:   Fri Nov 11 20:49:41 2016 -0800

    go/build: fix lack of error for Import of nonexistent local import path
    
    When calling build.Import, normally, an error is returned if the
    directory doesn't exist. However, that didn't happen for local
    import paths when build.FindOnly ImportMode was used.
    
    This change fixes that, and adds tests. It also makes the error
    value more consistent in all scenarios where it occurs.
    
    When calling build.Import with a local import path, the package
    can only exist in a single deterministic directory. That makes
    it possible verify that directory exists earlier in the path,
    and return a "cannot find package" error if it doesn't.
    Previously, this occurred only when build.FindOnly ImportMode
    was not set. It occurred quite late, after getting past Found
    label, to line that calls ctxt.readDir. Doing so would return
    an error like "no such file or directory" when the directory
    does not exist.
    
    Fixes #17863.
    Updates #17888 (relevant issue I ran into while working on this CL).
    
    Change-Id: If6a6996ac6176ac203a88bd31419748f88d89d7c
    Reviewed-on: https://go-review.googlesource.com/33158
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2dc714e1cf956edc6137be2674b5c8e44601873a
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sat Mar 4 01:28:53 2017 +0100

    Revert "cmd/go: add comment about SIGUSR2 on iOS"
    
    This reverts commit 4f0aac52d926270255fa2b682aca15e8ff404c59.
    
    Replaced by a improved strategy later in the CL relation chain.
    
    Change-Id: Iff0333f172443bb5b01a42ad06b01edeb6aa15bc
    Reviewed-on: https://go-review.googlesource.com/36068
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 3ce5371aaf6abc63ade70447198e9a88cb186910
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sat Mar 4 01:15:14 2017 +0100

    Revert "cmd/go, misc: rework cwd handling for iOS tests"
    
    This reverts commit 593ea3b3606a16da39e38406e22e373eeb944287.
    
    Replaced by a improved strategy later in the CL relation chain.
    
    Change-Id: I6963e4d1bf38e7028cf545a953e28054d83548
    Change-Id: I6963e4d1bf38e7028cf545a953e28054d8354870
    Reviewed-on: https://go-review.googlesource.com/36067
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 2b780af08e1a8767792b5990b06a102cfd96d7c4
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Feb 1 19:54:03 2017 +0100

    Revert "all: test adjustments for the iOS builder"
    
    This reverts commit 467109bf56fb560d1fd8a27c6184dbfe4f64ffef.
    
    Replaced by a improved strategy later in the CL relation chain.
    
    Change-Id: Ib90813b5a6c4716b563c8496013d2d57f9c022b8
    Reviewed-on: https://go-review.googlesource.com/36066
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 870d079c76ffd3766fb336a4071ae273867761d1
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 3 13:38:49 2017 -0800

    cmd/compile/internal/gc: replace Node.Ullman with Node.HasCall
    
    Since switching to SSA, the only remaining use for the Ullman field
    was in tracking whether or not an expression contained a function
    call. Give it a new name and encode it in our fancy new bitset field.
    
    Passes toolstash-check.
    
    Change-Id: I95b7f9cb053856320c0d66efe14996667e6011c2
    Reviewed-on: https://go-review.googlesource.com/37721
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 9fd359a29a8cc55ed665542d2a3fe9fef8baaa7d
Author: David Lazar <lazard@golang.org>
Date:   Mon Feb 20 09:55:54 2017 -0500

    cmd/compile: include position info in exported function bodies
    
    This gives accurate line numbers to inlined functions from another
    package. Previously AST nodes from another package would get the line
    number of the import statement for that package.
    
    The following benchmark results show how the size of package export data
    is impacted by this change. These benchmarks were created by compiling
    the go1 benchmark and running `go tool pack x` to extract the export
    data from the resulting .a files.
    
    name                                          old bytes   new bytes    delta
    bufio                                         3.59k ± 0%   4.17k ± 0%  +16.25%
    bytes                                         5.51k ± 0%   6.40k ± 0%  +16.21%
    compress/bzip2                                2.69k ± 0%   3.21k ± 0%  +19.74%
    compress/flate                                5.14k ± 0%   5.57k ± 0%   +8.43%
    compress/gzip                                 8.91k ± 0%  10.46k ± 0%  +17.32%
    container/list                                1.76k ± 0%   2.13k ± 0%  +21.51%
    context                                       4.51k ± 0%   5.47k ± 0%  +21.43%
    crypto                                        1.11k ± 0%   1.13k ± 0%   +1.90%
    crypto/aes                                      475 ± 0%     475 ± 0%   +0.00%
    crypto/cipher                                 1.18k ± 0%   1.18k ± 0%   +0.00%
    crypto/des                                      502 ± 0%     502 ± 0%   +0.00%
    crypto/dsa                                    5.96k ± 0%   6.54k ± 0%   +9.82%
    crypto/ecdsa                                  6.93k ± 0%   7.69k ± 0%  +10.91%
    crypto/elliptic                               6.53k ± 0%   7.17k ± 0%   +9.72%
    crypto/hmac                                     464 ± 0%     464 ± 0%   +0.00%
    crypto/internal/cipherhw                        313 ± 0%     313 ± 0%   +0.00%
    crypto/md5                                      695 ± 0%     711 ± 0%   +2.30%
    crypto/rand                                   5.62k ± 0%   6.21k ± 0%  +10.44%
    crypto/rc4                                      512 ± 0%     512 ± 0%   +0.00%
    crypto/rsa                                    7.31k ± 0%   8.10k ± 0%  +10.86%
    crypto/sha1                                     760 ± 0%     777 ± 0%   +2.24%
    crypto/sha256                                   523 ± 0%     523 ± 0%   +0.00%
    crypto/sha512                                   663 ± 0%     663 ± 0%   +0.00%
    crypto/subtle                                   873 ± 0%    1007 ± 0%  +15.35%
    crypto/tls                                    29.6k ± 0%   33.8k ± 0%  +14.03%
    crypto/x509                                   18.7k ± 0%   21.0k ± 0%  +12.56%
    crypto/x509/pkix                              10.6k ± 0%   12.2k ± 0%  +15.22%
    encoding                                        473 ± 0%     473 ± 0%   +0.00%
    encoding/asn1                                 1.42k ± 0%   1.50k ± 0%   +5.99%
    encoding/base64                               1.69k ± 0%   1.80k ± 0%   +6.88%
    encoding/binary                               2.76k ± 0%   3.51k ± 0%  +27.09%
    encoding/gob                                  13.5k ± 0%   15.2k ± 0%  +12.98%
    encoding/hex                                    857 ± 0%     881 ± 0%   +2.80%
    encoding/json                                 12.5k ± 0%   14.9k ± 0%  +19.37%
    encoding/pem                                    484 ± 0%     484 ± 0%   +0.00%
    errors                                          361 ± 0%     370 ± 0%   +2.49%
    flag                                          10.5k ± 0%   12.1k ± 0%  +14.92%
    fmt                                           1.42k ± 0%   1.42k ± 0%   +0.00%
    go/ast                                        15.8k ± 0%   17.5k ± 0%  +10.31%
    go/parser                                     8.13k ± 0%   9.86k ± 0%  +21.28%
    go/scanner                                    3.94k ± 0%   4.53k ± 0%  +14.73%
    go/token                                      3.53k ± 0%   3.75k ± 0%   +6.11%
    hash                                            507 ± 0%     507 ± 0%   +0.00%
    hash/crc32                                      685 ± 0%     685 ± 0%   +0.00%
    internal/nettrace                               474 ± 0%     474 ± 0%   +0.00%
    internal/poll                                 7.23k ± 0%   8.38k ± 0%  +15.90%
    internal/race                                   511 ± 0%     515 ± 0%   +0.78%
    internal/singleflight                           969 ± 0%    1075 ± 0%  +10.94%
    internal/syscall/unix                           427 ± 0%     427 ± 0%   +0.00%
    io                                            3.52k ± 0%   3.69k ± 0%   +4.82%
    io/ioutil                                     8.48k ± 0%   9.90k ± 0%  +16.72%
    log                                           5.06k ± 0%   5.98k ± 0%  +18.15%
    math                                          4.02k ± 0%   4.35k ± 0%   +8.26%
    math/big                                      9.28k ± 0%   9.94k ± 0%   +7.13%
    math/bits                                     3.47k ± 0%   4.33k ± 0%  +24.83%
    math/rand                                     1.30k ± 0%   1.32k ± 0%   +2.00%
    mime                                          2.98k ± 0%   3.70k ± 0%  +24.21%
    mime/multipart                                3.68k ± 0%   4.22k ± 0%  +14.65%
    mime/quotedprintable                          2.26k ± 0%   2.65k ± 0%  +17.60%
    net                                           23.0k ± 0%   25.7k ± 0%  +11.74%
    net/http                                      59.1k ± 0%   66.7k ± 0%  +13.00%
    net/http/httptest                             35.3k ± 0%   40.9k ± 0%  +15.80%
    net/http/httptrace                            15.3k ± 0%   17.6k ± 0%  +15.26%
    net/http/internal                             2.77k ± 0%   3.27k ± 0%  +17.89%
    net/textproto                                 4.60k ± 0%   5.25k ± 0%  +14.22%
    net/url                                       1.73k ± 0%   1.84k ± 0%   +6.59%
    os                                            14.3k ± 0%   16.4k ± 0%  +14.86%
    path                                            589 ± 0%     606 ± 0%   +2.89%
    path/filepath                                 5.07k ± 0%   6.17k ± 0%  +21.79%
    reflect                                       6.43k ± 0%   6.81k ± 0%   +5.90%
    regexp                                        5.88k ± 0%   6.46k ± 0%   +9.77%
    regexp/syntax                                 3.24k ± 0%   3.29k ± 0%   +1.73%
    runtime                                       13.1k ± 0%   14.9k ± 0%  +13.73%
    runtime/cgo                                     229 ± 0%     229 ± 0%   +0.00%
    runtime/debug                                 4.23k ± 0%   5.15k ± 0%  +21.79%
    runtime/internal/atomic                         905 ± 0%     905 ± 0%   +0.00%
    runtime/internal/sys                          2.04k ± 0%   2.20k ± 0%   +7.64%
    runtime/pprof                                 4.73k ± 0%   5.65k ± 0%  +19.41%
    runtime/trace                                   354 ± 0%     354 ± 0%   +0.00%
    sort                                          1.68k ± 0%   1.85k ± 0%  +10.17%
    strconv                                       1.85k ± 0%   1.95k ± 0%   +5.51%
    strings                                       3.98k ± 0%   4.53k ± 0%  +13.91%
    sync                                          1.52k ± 0%   1.58k ± 0%   +4.28%
    sync/atomic                                   1.60k ± 0%   1.74k ± 0%   +8.50%
    syscall                                       53.3k ± 0%   54.3k ± 0%   +1.84%
    testing                                       8.77k ± 0%  10.09k ± 0%  +14.96%
    testing/internal/testdeps                       598 ± 0%     600 ± 0%   +0.33%
    text/tabwriter                                3.63k ± 0%   4.41k ± 0%  +21.64%
    text/template                                 15.7k ± 0%   18.1k ± 0%  +15.67%
    text/template/parse                           9.12k ± 0%  10.35k ± 0%  +13.48%
    time                                          6.38k ± 0%   7.14k ± 0%  +11.81%
    unicode                                       4.62k ± 0%   4.66k ± 0%   +0.98%
    unicode/utf16                                   707 ± 0%     791 ± 0%  +11.88%
    unicode/utf8                                  1.06k ± 0%   1.20k ± 0%  +12.63%
    vendor/golang_org/x/crypto/chacha20poly1305   1.26k ± 0%   1.43k ± 0%  +13.54%
    vendor/golang_org/x/crypto/curve25519           392 ± 0%     392 ± 0%   +0.00%
    vendor/golang_org/x/crypto/poly1305             426 ± 0%     426 ± 0%   +0.00%
    vendor/golang_org/x/net/http2/hpack           4.75k ± 0%   5.77k ± 0%  +21.42%
    vendor/golang_org/x/net/idna                    355 ± 0%     355 ± 0%   +0.00%
    vendor/golang_org/x/net/lex/httplex             616 ± 0%     644 ± 0%   +4.55%
    vendor/golang_org/x/net/proxy                 7.76k ± 0%   9.58k ± 0%  +23.37%
    vendor/golang_org/x/text/transform            1.31k ± 0%   1.32k ± 0%   +0.46%
    vendor/golang_org/x/text/unicode/norm         5.89k ± 0%   6.84k ± 0%  +16.06%
    vendor/golang_org/x/text/width                1.24k ± 0%   1.27k ± 0%   +2.66%
    [Geo mean]                                    2.51k        2.74k        +9.14%
    
    Change-Id: I9ded911bb0ff63c530795fc85253d76b56d8abbc
    Reviewed-on: https://go-review.googlesource.com/37239
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 0824ae6dc1e867c327bdcae54be51a8a179c0f7d
Author: David Lazar <lazard@golang.org>
Date:   Fri Feb 17 16:55:40 2017 -0500

    cmd/compile: add flag for debugging PC-value tables
    
    For example, `-d pctab=pctoinline` prints the PC-inline table and
    inlining tree for every function.
    
    Change-Id: Ia6b9ce4d83eed0b494318d40ffe06481ec5d58ab
    Reviewed-on: https://go-review.googlesource.com/37235
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 301149b9e4f9e1256e80e76b0d8f6beace103731
Author: David Lazar <lazard@golang.org>
Date:   Fri Feb 17 16:20:52 2017 -0500

    cmd/internal/obj: avoid duplicate file name symbols
    
    The meaning of Version=1 was overloaded: it was reserved for file name
    symbols (to avoid conflicts with non-file name symbols), but was also
    used to mean "give me a fresh version number for this symbol."
    
    With the new inlining tree, the same file name symbol can appear in
    multiple entries, but each one would become a distinct symbol with its
    own version number.
    
    Now, we avoid duplicating symbols by using Version=0 for file name
    symbols and we avoid conflicts with other symbols by prefixing the
    symbol name with "gofile..".
    
    Change-Id: I8d0374053b8cdb6a9ca7fb71871b69b4dd369a9c
    Reviewed-on: https://go-review.googlesource.com/37234
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 781fd3998e3c1031379cf5709043a9c0e2987287
Author: David Lazar <lazard@golang.org>
Date:   Fri Feb 17 16:08:36 2017 -0500

    runtime: use inlining tables to generate accurate tracebacks
    
    The code in https://play.golang.org/p/aYQPrTtzoK now produces the
    following stack trace:
    
    goroutine 1 [running]:
    main.(*point).negate(...)
            /tmp/go/main.go:8
    main.main()
            /tmp/go/main.go:14 +0x23
    
    Previously the stack trace missed the inlined call:
    
    goroutine 1 [running]:
    main.main()
            /tmp/go/main.go:14 +0x23
    
    Fixes #10152.
    Updates #19348.
    
    Change-Id: Ib43c67012f53da0ef1a1e69bcafb65b57d9cecb2
    Reviewed-on: https://go-review.googlesource.com/37233
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 1c6ef9aeedf951a7a1fa7f510aa42150d3051567
Author: David Lazar <lazard@golang.org>
Date:   Fri Feb 17 16:07:47 2017 -0500

    cmd/compile: copy literals when inlining
    
    Without this, literals keep their original source positions through
    inlining, which results in strange jumps in line numbers of inlined
    function bodies. By copying literals, inlining can update their source
    position like other nodes.
    
    Fixes #15453.
    
    Change-Id: Iad5d9bbfe183883794213266dc30e31bab89ee69
    Reviewed-on: https://go-review.googlesource.com/37232
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 699175a11adfe57e859f4b995f4f5dfdaa5a5911
Author: David Lazar <lazard@golang.org>
Date:   Fri Feb 17 12:28:05 2017 -0500

    cmd/compile,link: generate PC-value tables with inlining information
    
    In order to generate accurate tracebacks, the runtime needs to know the
    inlined call stack for a given PC. This creates two tables per function
    for this purpose. The first table is the inlining tree (stored in the
    function's funcdata), which has a node containing the file, line, and
    function name for every inlined call. The second table is a PC-value
    table that maps each PC to a node in the inlining tree (or -1 if the PC
    is not the result of inlining).
    
    To give the appearance that inlining hasn't happened, the runtime also
    needs the original source position information of inlined AST nodes.
    Previously the compiler plastered over the line numbers of inlined AST
    nodes with the line number of the call. This meant that the PC-line
    table mapped each PC to line number of the outermost call in its inlined
    call stack, with no way to access the innermost line number.
    
    Now the compiler retains line numbers of inlined AST nodes and writes
    the innermost source position information to the PC-line and PC-file
    tables. Some tools and tests expect to see outermost line numbers, so we
    provide the OutermostLine function for displaying line info.
    
    To keep track of the inlined call stack for an AST node, we extend the
    src.PosBase type with an index into a global inlining tree. Every time
    the compiler inlines a call, it creates a node in the global inlining
    tree for the call, and writes its index to the PosBase of every inlined
    AST node. The parent of this node is the inlining tree index of the
    call. -1 signifies no parent.
    
    For each function, the compiler creates a local inlining tree and a
    PC-value table mapping each PC to an index in the local tree.  These are
    written to an object file, which is read by the linker.  The linker
    re-encodes these tables compactly by deduplicating function names and
    file names.
    
    This change increases the size of binaries by 4-5%. For example, this is
    how the go1 benchmark binary is impacted by this change:
    
    section             old bytes   new bytes   delta
    .text               3.49M ± 0%  3.49M ± 0%   +0.06%
    .rodata             1.12M ± 0%  1.21M ± 0%   +8.21%
    .gopclntab          1.50M ± 0%  1.68M ± 0%  +11.89%
    .debug_line          338k ± 0%   435k ± 0%  +28.78%
    Total               9.21M ± 0%  9.58M ± 0%   +4.01%
    
    Updates #19348.
    
    Change-Id: Ic4f180c3b516018138236b0c35e0218270d957d3
    Reviewed-on: https://go-review.googlesource.com/37231
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit ed70f37e73a656551077332449074b6b19686ab3
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Mon Feb 27 19:56:38 2017 +0200

    cmd/compile: pack bool fields in Node, Name, Func and Type structs to bitsets
    
    This reduces compiler memory usage by up to 4% - see compilebench
    results below.
    
    name       old time/op     new time/op     delta
    Template       245ms ± 4%      241ms ± 2%  -1.88%  (p=0.029 n=10+10)
    Unicode        126ms ± 3%      124ms ± 3%    ~     (p=0.105 n=10+10)
    GoTypes        805ms ± 2%      813ms ± 3%    ~     (p=0.515 n=8+10)
    Compiler       3.95s ± 2%      3.83s ± 1%  -2.96%  (p=0.000 n=9+10)
    MakeBash       47.4s ± 4%      46.6s ± 1%  -1.59%  (p=0.028 n=9+10)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        324M ± 5%       326M ± 3%    ~     (p=0.935 n=10+10)
    Unicode         186M ± 5%       178M ±10%    ~     (p=0.067 n=9+10)
    GoTypes        1.08G ± 7%      1.09G ± 4%    ~     (p=0.956 n=10+10)
    Compiler       5.34G ± 4%      5.31G ± 1%    ~     (p=0.501 n=10+8)
    
    name       old alloc/op    new alloc/op    delta
    Template      41.0MB ± 0%     39.8MB ± 0%  -3.03%  (p=0.000 n=10+10)
    Unicode       32.3MB ± 0%     31.0MB ± 0%  -4.13%  (p=0.000 n=10+10)
    GoTypes        119MB ± 0%      116MB ± 0%  -2.39%  (p=0.000 n=10+10)
    Compiler       499MB ± 0%      487MB ± 0%  -2.48%  (p=0.000 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        380k ± 1%       379k ± 1%    ~     (p=0.436 n=10+10)
    Unicode         324k ± 1%       324k ± 0%    ~     (p=0.853 n=10+10)
    GoTypes        1.15M ± 0%      1.15M ± 0%    ~     (p=0.481 n=10+10)
    Compiler       4.41M ± 0%      4.41M ± 0%  -0.12%  (p=0.007 n=10+10)
    
    name       old text-bytes  new text-bytes  delta
    HelloSize       623k ± 0%       623k ± 0%    ~     (all equal)
    CmdGoSize      6.64M ± 0%      6.64M ± 0%    ~     (all equal)
    
    name       old data-bytes  new data-bytes  delta
    HelloSize      5.81k ± 0%      5.81k ± 0%    ~     (all equal)
    CmdGoSize       238k ± 0%       238k ± 0%    ~     (all equal)
    
    name       old bss-bytes   new bss-bytes   delta
    HelloSize       134k ± 0%       134k ± 0%    ~     (all equal)
    CmdGoSize       152k ± 0%       152k ± 0%    ~     (all equal)
    
    name       old exe-bytes   new exe-bytes   delta
    HelloSize       967k ± 0%       967k ± 0%    ~     (all equal)
    CmdGoSize      10.2M ± 0%      10.2M ± 0%    ~     (all equal)
    
    Change-Id: I1f40af738254892bd6c8ba2eb43390b175753d52
    Reviewed-on: https://go-review.googlesource.com/37445
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fbf4dd91b9762eeb038c141b21712534310795f1
Author: Johan Brandhorst <johan.brandhorst@gmail.com>
Date:   Wed Dec 21 13:49:04 2016 +0000

    net/http/httptest: add Client and Certificate methods to Server
    
    Adds a function for easily accessing the x509.Certificate
    of a Server, if there is one. Also adds a helper function
    for getting a http.Client suitable for use with the server.
    
    This makes the steps required to test a httptest
    TLS server simpler.
    
    Fixes #18411
    
    Change-Id: I2e78fe1e54e31bed9c641be2d9a099f698c7bbde
    Reviewed-on: https://go-review.googlesource.com/34639
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 02e36f8c87901fde841cb5e82e0409b7914572b8
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 3 11:35:44 2017 -0800

    cmd/compile/internal/ssa: remove Hmul{8,16}{,u} ops
    
    Change-Id: I90865921584ae4bdfb6c220d439b14593d72b6f9
    Reviewed-on: https://go-review.googlesource.com/37752
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit c8eaeb8cba52a1eb688245e0f6935d560cf1569d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Feb 3 16:18:01 2017 -0500

    cmd/compile: remove zeroing after newobject
    
    The Zero op right after newobject has been removed. But this rule
    does not cover Store of constant zero (for SSA-able types). Add
    rules to cover Store op as well.
    
    Updates #19027.
    
    Change-Id: I5d2b62eeca0aa9ce8dc7205b264b779de01c660b
    Reviewed-on: https://go-review.googlesource.com/36836
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 9b480521d84654c90f6030675e1e10655b180274
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Mar 3 13:53:13 2017 -0500

    cmd/compile: fix optimization of Zero newobject on amd64p32
    
    On amd64p32, PtrSize and RegSize don't agree, and function return
    value is aligned with RegSize. Fix this rule. Other architectures
    are not affected, where PtrSize and RegSize are the same.
    
    Change-Id: If187d3dfde3dc3b931b8e97db5eeff49a781551b
    Reviewed-on: https://go-review.googlesource.com/37720
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d8a0f748013dae2a731a9a37f94ab0e37d096ee2
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Mar 3 02:43:44 2017 -0800

    cmd/compile/internal/gc: remove OHMUL Op
    
    Previously the compiler rewrote constant division into OHMUL
    operations, but that rewriting was moved to SSA in CL 37015. Now OHMUL
    is unused, so we can get rid of it.
    
    Change-Id: Ib6fc7c2b6435510bafb5735b3b4f42cfd8ed8cdb
    Reviewed-on: https://go-review.googlesource.com/37750
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 77f64c50dbaca5fcec8198a575f6c345cb80ad69
Author: Austin Clements <austin@google.com>
Date:   Sun Feb 5 19:34:16 2017 -0500

    runtime: clarify access to mheap_.busy
    
    There are two accesses to mheap_.busy that are guarded by checks
    against len(mheap_.free). This works because both lists are (and must
    be) the same length, but it makes the code less clear. Change these to
    use len(mheap_.busy) so the access more clearly parallels the check.
    
    Fixes #18944.
    
    Change-Id: I9bacbd3663988df351ed4396ae9018bc71018311
    Reviewed-on: https://go-review.googlesource.com/36354
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit b50b72858714f29f7147c0d58444398345ac8188
Author: Austin Clements <austin@google.com>
Date:   Thu Dec 22 18:00:18 2016 -0700

    runtime: simplify sweep allocation counting
    
    Currently sweep counts the number of allocated objects, computes the
    number of free objects from that, then re-computes the number of
    allocated objects from that. Simplify and clean this up by skipping
    these intermediate steps.
    
    Change-Id: I3ed98e371eb54bbcab7c8530466c4ab5fde35f0a
    Reviewed-on: https://go-review.googlesource.com/34935
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Marvin Stenger <marvin.stenger94@gmail.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit f1ba75f8c577e1471f646ef3715fc2f41dd423ef
Author: Austin Clements <austin@google.com>
Date:   Tue Jan 31 11:46:36 2017 -0500

    runtime: don't rescan finalizers queue during mark termination
    
    Currently we scan the finalizers queue both during concurrent mark and
    during mark termination. This costs roughly 20ns per queued finalizer
    and about 1ns per unused finalizer queue slot (allocated queue length
    never decreases), which can drive up STW time if there are many
    finalizers.
    
    However, we only add finalizers to this queue during sweeping, which
    means that the second scan will never find anything new. Hence, we can
    fix this by simply not scanning the finalizers queue during mark
    termination. This brings the STW time under the 100µs goal even with
    1,000,000 queued finalizers.
    
    Fixes #18869.
    
    Change-Id: I4ce5620c66fb7f13ebeb39ca313ce57047d1d0fb
    Reviewed-on: https://go-review.googlesource.com/36013
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 98da2d1f91a8f4e6bdaecd8a98fc77cbac211c80
Author: Austin Clements <austin@google.com>
Date:   Sun Jan 29 22:47:27 2017 -0500

    runtime: remove wbufptr
    
    Since workbuf is now marked go:notinheap, the write barrier-preventing
    wrapper type wbufptr is no longer necessary. Remove it.
    
    Change-Id: I3e5b5803a1547d65de1c1a9c22458a38e08549b7
    Reviewed-on: https://go-review.googlesource.com/35971
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 8eb14e9de5018afbcf7eefd7b4bce7a200d0ce3f
Author: Austin Clements <austin@google.com>
Date:   Wed Feb 22 16:13:06 2017 -0500

    cmd/compile: accept string debug flags
    
    The compiler's -d flag accepts string-valued flags, but currently only
    for SSA debug flags. Extend it to support string values for other
    flags. This also makes the syntax somewhat more sane so flag=value and
    flag:value now both accept integers and strings.
    
    Change-Id: Idd144d8479a430970cc1688f824bffe0a56ed2df
    Reviewed-on: https://go-review.googlesource.com/37345
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 5bfd1ef036f2cd549f78a0acd3e2666b42bcc07d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Feb 2 19:47:59 2017 -0500

    cmd/compile: get rid of "volatile" in SSA
    
    A value is "volatile" if it is a pointer to the argument region
    on stack which will be clobbered by function call. This is used
    to make sure the value is safe when inserting write barrier calls.
    The writebarrier pass can tell whether a value is such a pointer.
    Therefore no need to mark it when building SSA and thread this
    information through.
    
    Passes "toolstash -cmp" on std.
    
    Updates #17583.
    
    Change-Id: Idc5fc0d710152b94b3c504ce8db55ea9ff5b5195
    Reviewed-on: https://go-review.googlesource.com/36835
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 4775b7feb10014751da2669da973fb852f09aebe
Author: Will Storey <will@summercat.com>
Date:   Sun Feb 19 21:24:17 2017 -0800

    image/gif: handle an extra data sub-block byte.
    
    This changes the decoder's behaviour when there is stray/extra data
    found after an image is decompressed (e.g., data sub-blocks after an LZW
    End of Information Code). Instead of raising an error, we silently skip
    over such data until we find the end of the image data marked by a Block
    Terminator. We skip at most one byte as sample problem GIFs exhibit this
    property.
    
    GIFs should not have and do not need such stray data (though the
    specification is arguably ambiguous). However GIFs with such properties
    have been seen in the wild.
    
    Fixes #16146
    
    Change-Id: Ie7e69052bab5256b4834992304e6ca58e93c1879
    Reviewed-on: https://go-review.googlesource.com/37258
    Reviewed-by: Nigel Tao <nigeltao@golang.org>
    Run-TryBot: Nigel Tao <nigeltao@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9b15c13dc567d92c6ac628c762d42b4ae2c9469f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 15:10:55 2017 -0800

    runtime/pprof: fix data race between Profile.Add and Profile.WriteTo
    
    p.m is accessed in WriteTo without holding p.mu.
    Move the access inside the critical section.
    
    The race detector catches this bug using this program:
    
    
    package main
    
    import (
            "os"
            "runtime/pprof"
            "time"
    )
    
    func main() {
            p := pprof.NewProfile("ABC")
            go func() {
                    p.WriteTo(os.Stdout, 1)
                    time.Sleep(time.Second)
            }()
            p.Add("abc", 0)
            time.Sleep(time.Second)
    }
    
    
    $ go run -race x.go
    ==================
    WARNING: DATA RACE
    Write at 0x00c42007c240 by main goroutine:
      runtime.mapassign()
          /Users/josh/go/tip/src/runtime/hashmap.go:485 +0x0
      runtime/pprof.(*Profile).Add()
          /Users/josh/go/tip/src/runtime/pprof/pprof.go:281 +0x255
      main.main()
          /Users/josh/go/tip/src/p.go:15 +0x9d
    
    Previous read at 0x00c42007c240 by goroutine 6:
      runtime/pprof.(*Profile).WriteTo()
          /Users/josh/go/tip/src/runtime/pprof/pprof.go:314 +0xc5
      main.main.func1()
          /Users/josh/go/tip/src/x.go:12 +0x69
    
    Goroutine 6 (running) created at:
      main.main()
          /Users/josh/go/tip/src/x.go:11 +0x6e
    ==================
    ABC profile: total 1
    1 @ 0x110ccb4 0x111aeee 0x1055053 0x107f031
    
    Found 1 data race(s)
    exit status 66
    
    
    (Exit status 66?)
    
    Change-Id: I49d884dc3af9cce2209057a3448fe6bf50653523
    Reviewed-on: https://go-review.googlesource.com/37730
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 72359cf840801adc81058cc0430effa1a34da0e2
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Mar 2 11:13:06 2017 -0800

    go/types: don't exclude package unsafe from a Package's Imports list
    
    There's no good reason to exclude it and it only makes the code more
    complicated and less consistent. Having it in the list provides an
    easy way to detect if a package uses operations from package unsafe.
    
    Change-Id: I2f9b0485db0a680bd82f3b93a350b048db3f7701
    Reviewed-on: https://go-review.googlesource.com/37694
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 359ca5ccc8b801d6ca8f8e417135436ccde00212
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Mar 1 17:15:43 2017 -0800

    go/types: support type checking of external tests with gotype
    
    - renamed -a flag to -t
    - added -x flag to specify external test files
    - improved documentation and usage string
    
    Change-Id: I7c850bd28a10ceaa55d599c22db07774147aa3f7
    Reviewed-on: https://go-review.googlesource.com/37656
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 9eac1c87a60de7b2b9a4fba01e31a852c01aaf97
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Mar 1 15:35:24 2017 -0800

    go/types: gotype to always report the same first error now
    
    The old code may have reported different errors given an
    erroneous package depending on the order in which files
    were parsed concurrently. The new code always reports
    errors in "file order", independent of processing order.
    
    Also:
    - simplified parsing code and internal concurrency control
    - removed -seq flag which didn't really add useful functionality
    
    Change-Id: I18e24e630f458f2bc107a7b83926ae761d63c334
    Reviewed-on: https://go-review.googlesource.com/37655
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 3a90bfb2531c373be05bfe8ff27ba475f9c75cb2
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 14:12:09 2017 -0800

    cmd/dist, cmd/compile: eliminate mergeEnvLists copies
    
    This is now handled by os/exec.
    
    Updates #12868
    
    Change-Id: Ic21a6ff76a9b9517437ff1acf3a9195f9604bb45
    Reviewed-on: https://go-review.googlesource.com/37698
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9d29be468eb9092f9dea3e10d32e1f7848a55458
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Feb 27 05:41:50 2017 +0000

    net/http: clean up Transport.RoundTrip error handling
    
    If I put a 10 millisecond sleep at testHookWaitResLoop, before the big
    select in (*persistConn).roundTrip, two flakes immediately started
    happening, TestTransportBodyReadError (#19231) and
    TestTransportPersistConnReadLoopEOF.
    
    The problem was that there are many ways for a RoundTrip call to fail
    (errors reading from Request.Body while writing the response, errors
    writing the response, errors reading the response due to server
    closes, errors due to servers sending malformed responses,
    cancelations, timeouts, etc.), and many of those failures then tear
    down the TCP connection, causing more failures, since there are always
    at least three goroutines involved (reading, writing, RoundTripping).
    
    Because the errors were communicated over buffered channels to a giant
    select, the error returned to the caller was a function of which
    random select case was called, which was why a 10ms delay before the
    select brought out so many bugs. (several fixed in my previous CLs the past
    few days).
    
    Instead, track the error explicitly in the transportRequest, guarded
    by a mutex.
    
    In addition, this CL now:
    
    * differentiates between the two ways writing a request can fail: the
      io.Copy reading from the Request.Body or the io.Copy writing to the
      network. A new io.Reader type notes read errors from the
      Request.Body. The read-from-body vs write-to-network errors are now
      prioritized differently.
    
    * unifies the two mapRoundTripErrorFromXXX methods into one
      mapRoundTripError method since their logic is now the same.
    
    * adds a (*Request).WithT(*testing.T) method in export_test.go, usable
      by tests, to call t.Logf at points during RoundTrip. This is disabled
      behind a constant except when debugging.
    
    * documents and deflakes TestClientRedirectContext
    
    I've tested this CL with high -count values, with/without -race,
    with/without delays before the select, etc. So far it seems robust.
    
    Fixes #19231 (TestTransportBodyReadError flake)
    Updates #14203 (source of errors unclear; they're now tracked more)
    Updates #15935 (document Transport errors more; at least understood more now)
    
    Change-Id: I3cccc3607f369724b5344763e35ad2b7ea415738
    Reviewed-on: https://go-review.googlesource.com/37495
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 87649d32ad16a9a0b7bd5dbd1c124b2032a270f1
Author: Mike Danese <mikedanese@google.com>
Date:   Wed Mar 1 10:43:57 2017 -0800

    crypto/tls: make Config.Clone also clone the GetClientCertificate field
    
    Using GetClientCertificate with the http client is currently completely
    broken because inside the transport we clone the tls.Config and pass it
    off to the tls.Client. Since tls.Config.Clone() does not pass forward
    the GetClientCertificate field, GetClientCertificate is ignored in this
    context.
    
    Fixes #19264
    
    Change-Id: Ie214f9f0039ac7c3a2dab8ffd14d30668bdb4c71
    Signed-off-by: Mike Danese <mikedanese@google.com>
    Reviewed-on: https://go-review.googlesource.com/37541
    Reviewed-by: Filippo Valsorda <hi@filippo.io>
    Reviewed-by: Adam Langley <agl@golang.org>
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2f5aea7c1305158a6d87ea356f0ed2a8186bc73d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 18:46:00 2017 +0000

    Revert "Revert "cmd/vet/all: remove pprof from the whitelist""
    
    This reverts commit 9bd1cc3fa1145182e9ce041d0e96bd2051cd7fcf.
    
    Reason for revert: New fixes in from upstream. Try this again.
    
    Change-Id: Iea46f32857e8467f8d5a49b31e20a52fda8bce60
    Reviewed-on: https://go-review.googlesource.com/37693
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 85bae0a9df6ea04256c7787d14ae789fa697618a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 09:11:03 2017 -0800

    cmd/vet/all: remove printf hacks
    
    Now that vet loads from source,
    fmt can always be correctly resolved,
    so the fmt.Formatter type is always available,
    so we can reinstate the check.
    
    Change-Id: I17f0c7fccf6960c9415de8774b15123135d57be8
    Reviewed-on: https://go-review.googlesource.com/37692
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8a93546d686223e7f03658c69f08be397b62cb25
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 09:07:14 2017 -0800

    cmd/vet/all: vet using only source
    
    This simplifies the code and speeds it up.
    It also allows us to eliminate some other TODOs;
    those will come in a follow-up CL.
    
    Running for the host platform, before:
    
    real    0m9.907s
    user    0m14.566s
    sys     0m1.058s
    
    After:
    
    real    0m7.841s
    user    0m12.339s
    sys     0m0.572s
    
    Running for a single non-host platform, before:
    
    real    0m8.784s
    user    0m15.451s
    sys     0m3.445s
    
    After:
    
    real    0m7.681s
    user    0m12.122s
    sys     0m0.577s
    
    Running for all platforms, before:
    
    real    7m4.480s
    user    8m43.398s
    sys     1m15.683s
    
    After:
    
    real    4m37.596s
    user    7m30.729s
    sys     0m18.533s
    
    It also makes my laptop considerably more
    responsive while running for all platforms.
    
    Change-Id: I748689fea0d2d4ef61aca2ce5524d03d8fafa5ca
    Reviewed-on: https://go-review.googlesource.com/37691
    Reviewed-by: Rob Pike <r@golang.org>

commit ddbee9abd45bbb955fec973fc0e395276127d431
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 09:04:03 2017 -0800

    cmd/vet: support importing from source
    
    Add a -source flag to cmd/vet that instructs
    it to typecheck purely from source code.
    
    Updates #16086
    Fixes #19332
    
    Change-Id: Ic83d0f14d5bb837a329d539b2873aeccdf7bf669
    Reviewed-on: https://go-review.googlesource.com/37690
    Reviewed-by: Rob Pike <r@golang.org>

commit 7e74d432915a8f22b07f8d29aa8e02245f8d8cd1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 07:56:53 2017 -0800

    cmd/vet: refactor to support alternative importers
    
    Instead of constructing the importer in init, do it lazily as needed.
    This lets us select the importer using a command line flag.
    The addition of the command line flag will come in a follow-up CL.
    
    Change-Id: Ieb3a5f01a34fb5bd220a95086daf5d6b37e83bb5
    Reviewed-on: https://go-review.googlesource.com/37669
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4de2efe927e3c71798dc55f7b029a8476b756a45
Author: Heschi Kreinick <heschi@google.com>
Date:   Wed Feb 22 17:58:31 2017 -0500

    cmd/trace: traces may end with pending mark assists
    
    There's no guarantee that all in-progress mark assists will finish
    before the trace does. Don't crash if that happens.
    
    I haven't added a test because there's quite a bit of ceremony involved
    and the bug is fairly straightforward.
    
    Change-Id: Ia1369a8e2260fc6a328ad204a1eab1063d2e2c90
    Reviewed-on: https://go-review.googlesource.com/37540
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 04fc88776187a3c9c35a575eef65a7e9b9276e6d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 06:30:26 2017 -0800

    runtime: delay marking maps as writing until after first alg call
    
    Fixes #19359
    
    Change-Id: I196b47cf0471915b6dc63785e8542aa1876ff695
    Reviewed-on: https://go-review.googlesource.com/37665
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 0ee9c46cb19f04e713a5db30c8ae0e719c6d228b
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Mar 1 21:46:08 2017 -0800

    cmd/compile: add missing WBs for reflect.{Slice,String}Header.Data
    
    Fixes #19168.
    
    Change-Id: I3f4fcc0b189c53819ac29ef8de86fdad76a17488
    Reviewed-on: https://go-review.googlesource.com/37663
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3d77bc081d16b02bb937fdd648f4e83ebc0b9b95
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 07:23:43 2017 -0800

    cmd/vet/all: use SizesFor to calculate archbits
    
    Change-Id: I99706807782f11e8d24baf953424a9e292a2cbac
    Reviewed-on: https://go-review.googlesource.com/37668
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 99fbccbd930fbcf86e41fa7e969894f69bc8b95b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 1 21:11:25 2017 -0800

    cmd/vet: use types.SizesFor
    
    This eliminates a duplicate copy of
    the SizesFor map.
    
    Change-Id: I51e44ea8ee860901086616e3f4dfa32aaa9b4d2d
    Reviewed-on: https://go-review.googlesource.com/37667
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 32a1736d2416363cea43a297632acc7414e77032
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Mar 2 07:02:26 2017 -0800

    go/types: add a compiler param to SizesFor
    
    The current StdSizes most closely matches
    the gc compiler, and the uses I know of that care
    which compiler the sizes are for are all for
    the gc compiler, so call the existing
    implementation "gc".
    
    Updates #17586
    Fixes #19351
    
    Change-Id: I2bdd694518fbe233473896321a1f9758b46ed79b
    Reviewed-on: https://go-review.googlesource.com/37666
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 542a60fbdea55ca022fd87146a31b09058ee6850
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Feb 28 13:44:33 2017 -0800

    cmd/compile: don't crash when slicing non-slice
    
    Fixes #19323
    
    Change-Id: I92d1bdefb15de6178a577a4fa0f0dc004f791904
    Reviewed-on: https://go-review.googlesource.com/37584
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6c85fb08c27c0346125567573e0f0a2b8124bc3b
Author: Russ Cox <rsc@golang.org>
Date:   Wed Mar 1 21:09:52 2017 -0500

    time: strip monotonic time in t.Round, t.Truncate
    
    The original analysis of the Go corpus assumed that these
    stripped monotonic time. During the design discussion we
    decided to try not stripping monotonic time here, but existing
    code works better if we do.
    
    See the discussion on golang.org/issue/18991 for more details.
    
    For #18991.
    
    Change-Id: I04d355ffe56ca0317acdd2ca76cb3033c277f6d1
    Reviewed-on: https://go-review.googlesource.com/37542
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f072283bcecb66f8c22046cad4d8ddcc458d32e7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Mar 2 02:45:39 2017 +0000

    net/http: add more debugging to TestServerAllowsBlockingRemoteAddr
    
    It fails on Solaris often, but nowhere else.
    
    Not sure why. Add some debugging.
    
    Change-Id: I79fc710bd339ae972d624c73a46bd8d215729c10
    Reviewed-on: https://go-review.googlesource.com/37659
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a143f5d646e5e6a835f70235ee5e604deda5f3be
Author: Philip Hofer <phofer@umich.edu>
Date:   Tue Feb 28 14:35:37 2017 -0800

    cmd/internal/obj/arm: improve static branch prediction for wrapper prologue
    
    This is a follow-up to CL 36893.
    
    Move the unlikely branch in the wrapper prologue to the end
    of the function, where it has minimal impact on the instruction
    cache. Static branch prediction is also less likely to choose
    a forward branch.
    
    Updates #19042
    
    sort benchmarks:
    name                  old time/op  new time/op  delta
    SearchWrappers-4      1.44µs ± 0%  1.45µs ± 0%  +1.15%  (p=0.000 n=9+10)
    SortString1K-4        1.02ms ± 0%  1.04ms ± 0%  +2.39%  (p=0.000 n=10+10)
    SortString1K_Slice-4   960µs ± 0%   989µs ± 0%  +2.95%  (p=0.000 n=9+10)
    StableString1K-4       218µs ± 0%   213µs ± 0%  -2.13%  (p=0.000 n=10+10)
    SortInt1K-4            541µs ± 0%   543µs ± 0%  +0.30%  (p=0.003 n=9+10)
    StableInt1K-4          760µs ± 1%   763µs ± 1%  +0.38%  (p=0.011 n=10+10)
    StableInt1K_Slice-4    840µs ± 1%   779µs ± 0%  -7.31%  (p=0.000 n=9+10)
    SortInt64K-4          55.2ms ± 0%  55.4ms ± 1%  +0.34%  (p=0.012 n=10+8)
    SortInt64K_Slice-4    56.2ms ± 0%  55.6ms ± 1%  -1.16%  (p=0.000 n=10+10)
    StableInt64K-4        70.9ms ± 1%  71.0ms ± 0%    ~     (p=0.315 n=10+7)
    Sort1e2-4              250µs ± 0%   249µs ± 1%    ~     (p=0.315 n=9+10)
    Stable1e2-4            600µs ± 0%   594µs ± 0%  -1.09%  (p=0.000 n=9+10)
    Sort1e4-4             51.2ms ± 0%  51.4ms ± 1%  +0.40%  (p=0.001 n=9+10)
    Stable1e4-4            204ms ± 1%   199ms ± 1%  -2.27%  (p=0.000 n=10+10)
    Sort1e6-4              8.42s ± 0%   8.44s ± 0%  +0.28%  (p=0.000 n=8+9)
    Stable1e6-4            43.3s ± 0%   42.5s ± 1%  -1.89%  (p=0.000 n=9+9)
    
    Change-Id: I827559aa557fdba211a38ce3f77137b471c5c67e
    Reviewed-on: https://go-review.googlesource.com/37611
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit a2cc8b20fd6576b1db729159c096aadfb91fbb9e
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Mar 1 17:31:06 2017 -0800

    cmd/go: fix TestFFLAGS for Fortran compilers that accept unknown options
    
    The test assumed that passing an unknown option to the Fortran
    compiler would cause the compiler to fail. Unfortunately it appears
    that some succeed. It's irrelevant to the actual test, which is
    verifying that the flag was indeed passed.
    
    Fixes #19080.
    
    Change-Id: Ib9e89447a2104e4742f4b98938373fc2522772aa
    Reviewed-on: https://go-review.googlesource.com/37658
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 86abfbb931312db022ed8437a89219791ced6435
Author: Russ Cox <rsc@golang.org>
Date:   Mon Feb 27 20:22:48 2017 -0500

    doc/devel: update release.html for new support policy
    
    Fixes #19069.
    
    Change-Id: I211a304ec57d6b94366af4c0db413c8055b9634d
    Reviewed-on: https://go-review.googlesource.com/37531
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Steve Francia <spf@golang.org>

commit f6698cf3402f2d3049216b599ea1bc06d9de15c2
Author: Michel Lespinasse <walken@google.com>
Date:   Sat Jan 21 00:13:36 2017 -0800

    vendor: import golang.org/x/net/proxy
    
    Add golang.org/x/net/proxy from x/net git rev a689eb3bc4b5
    
    Change-Id: I4ceb2cf5686042c545fe69868537a66e083139de
    Reviewed-on: https://go-review.googlesource.com/35552
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 36f55a8b6125c9ae951487a0ad074b5c991f7b92
Author: Michel Lespinasse <walken@google.com>
Date:   Fri Jan 20 16:52:23 2017 -0800

    net/http: add support for socks5 proxy
    
    See #18508
    
    This commit adds http Client support for socks5 proxies.
    
    Change-Id: Ib015f3819801da13781d5acdd780149ae1f5857b
    Reviewed-on: https://go-review.googlesource.com/35488
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4be4da6331b4acfc379113bd5603079a4f36741a
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Mar 1 14:20:27 2017 -0800

    go/types: change local gotype command to use source importer
    
    Also: Remove -gccgo flag (not supported after 1.5), minor
    cleanups.
    
    Change-Id: I625241b07b277ac50ff836e2230b7b285887d35e
    Reviewed-on: https://go-review.googlesource.com/37654
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 13c35a1b204f6e580b220e0df409a2c186e648a4
Author: Keith Randall <khr@golang.org>
Date:   Wed Mar 1 13:08:22 2017 -0800

    cmd/compile: ppc64x no longer needs a scratch stack location
    
    After https://go-review.googlesource.com/c/36725/, ppc64x no longer
    needs a temp stack location for int reg <-> fp reg moves.
    
    Update #18922
    
    Change-Id: Ib4319784f7a855f593dfa5231604ca2c24e4c882
    Reviewed-on: https://go-review.googlesource.com/37651
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit 466a8915e3274c8b33d82c68c187bd6c5653d80e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 1 13:40:24 2017 -0800

    net/smtp: skip flaky TestTLSClient on freebsd/amd64
    
    Updates #19229
    
    Change-Id: Ibe1ea090ac064c7eb5abd225214ab43744bafbc4
    Reviewed-on: https://go-review.googlesource.com/37653
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a047f72dcfee7504a4dadc5e4c392dda761a3738
Author: Raul Silvera <rsilvera@google.com>
Date:   Wed Mar 1 13:21:29 2017 -0800

    cmd/vendor/github.com/google/pprof: refresh from upstream
    
    Updating to commit dec22b42d9eee442222c36c8da24ddc9905e7ee6
    from github.com/google/pprof
    
    Fixes #19322.
    
    Change-Id: I1bc3fcd381f22d52557f61c6fb694f54fc64470c
    Reviewed-on: https://go-review.googlesource.com/37652
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 1eed80f09adc7641227ac80a9031fd3781e54d7d
Author: Keith Randall <khr@golang.org>
Date:   Sun Feb 19 09:37:22 2017 -0800

    cmd/compile: fix disassembly of invalid instructions
    
    Make sure that if we encode an explicit base register, we print it.
    That will ensure that if we make an Addr with an auto variable but
    a base that isn't SP, then it will be obvious from the disassembly.
    
    Update #19184
    
    Change-Id: If5556a5183f344d719ec7197aa935a0166061e6f
    Reviewed-on: https://go-review.googlesource.com/37255
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit ffe923f6f427f5ab47f4e1f4584369212add07e2
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Mar 1 17:44:11 2017 +0000

    net/http: deflake TestClientRedirect308NoGetBody
    
    In an unrelated CL I found a way to increase the likelihood of latent
    flaky tests and found this one.
    
    This is just like yesterday's https://golang.org/cl/37624 and dozens
    before it (all remnants from the great net/http test parallelization
    of Nov 2016 in https://golang.org/cl/32684).
    
    Change-Id: I3fe61d1645062e5109206ff27d74f573ef6ebb2e
    Reviewed-on: https://go-review.googlesource.com/37627
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 95c9583a182ff2433fde9fae98f5dde9edd337e1
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Fri Feb 10 17:45:12 2017 -0500

    cmd/compile: intrinsify atomics on ppc64x
    
    This adds the necessary changes so that atomics are treated as
    intrinsics on ppc64x.
    
    The implementations of And8 and Or8 require power8 for
    both ppc64 and ppc64le.  This is a new requirement
    for ppc64.
    
    Fixes #8739
    
    Change-Id: Icb85e2755a49166ee3652668279f6ed5ebbca901
    Reviewed-on: https://go-review.googlesource.com/36832
    Reviewed-by: Keith Randall <khr@golang.org>

commit a6a0b1903d7029abe14f00c6cf2138cde9d81160
Author: Andreas Auernhammer <aead@mail.de>
Date:   Sun Feb 12 21:41:29 2017 +0100

    crypto: add BLAKE2b and BLAKE2s hash constants
    
    Fixes golang/go#19060
    Change-Id: I1844edc3dcccc8d83a11d1145b60b2b92f2658ca
    Reviewed-on: https://go-review.googlesource.com/36876
    Reviewed-by: Adam Langley <agl@golang.org>

commit d271576a0f7578288d663afee9d308e67e4a9d48
Author: Joe Shaw <joe@joeshaw.org>
Date:   Fri Feb 17 11:55:42 2017 -0500

    encoding/pem: refuse extra data on ending line
    
    Previously the code didn't check for extra data after the final five
    dashes of the ending line of a PEM block.
    
    Fixes #19147
    Fixes #7042
    
    Change-Id: Idaab2390914a2bed8c2c12b14dfb6d68233fdfec
    Reviewed-on: https://go-review.googlesource.com/37147
    Reviewed-by: Adam Langley <agl@golang.org>

commit b2a2a6054a015eddad4043f55fa280aed0334607
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 28 17:40:33 2017 -0800

    go/internal/srcimporter: report reimport of incomplete packages
    
    See the issue below for details.
    
    For #19337.
    
    Change-Id: I7637dcd4408f1bc4a9b3050a107aadb4de6f950b
    Reviewed-on: https://go-review.googlesource.com/37620
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit e54bc92a2ca9a3726dcf5a72cd7e993766e16dd8
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Wed Feb 22 13:32:54 2017 -0500

    runtime, cmd/go: roll back stale message, test detail
    
    Some debugging code was recently added to:
    1) provide more detail for the stale reason when it is
    determined that a package is stale
    2) provide file and package time and date information when
    it is determined that runtime.a is stale
    
    This backs out those those debugging messages.
    
    Fixes #19116
    
    Change-Id: I8dd0cbe29324820275b481d8bbb78ff2c5fbc362
    Reviewed-on: https://go-review.googlesource.com/37382
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 32bb13abbeeeeb91dd4deea746bcb85cc2634682
Author: David du Colombier <0intro@gmail.com>
Date:   Wed Mar 1 18:39:01 2017 +0100

    cmd/vendor/github.com/google/pprof: refresh from upstream
    
    Updating to commit b1c91b9f8fa7647e4c43c96c50f245df551f7013
    from github.com/google/pprof
    
    Fixes #19342.
    
    Change-Id: Iaaa283cdce3f9bf3e5fe713be7d23c477b579092
    Reviewed-on: https://go-review.googlesource.com/37634
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 44053de36509f4634befbf7ad442b5debd0f0cdf
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Fri Feb 10 15:40:44 2017 -0500

    cmd/compile: use reg moves for int <-> float conversions on ppc64x
    
    This makes a change in the SSA code generated for OpPPC64Xf2i64
    and OpPPC64Xi2f64 to use register based instructions to convert
    between float and integer.  This will require at least power8.
    Currently the conversion is done by storing to and loading
    from memory, which is more expensive.
    
    This improves some of the math functions:
    
    BenchmarkExp-128                     74.1          66.8          -9.85%
    BenchmarkExpGo-128                   87.4          66.3          -24.14%
    BenchmarkExp2-128                    72.2          64.3          -10.94%
    BenchmarkExp2Go-128                  74.3          65.9          -11.31%
    
    BenchmarkLgamma-128                  51.0          39.7          -22.16%
    BenchmarkLog-128                     42.9          40.6          -5.36%
    BenchmarkLogb-128                    11.5          9.16          -20.35%
    BenchmarkLog1p-128                   38.9          36.2          -6.94%
    
    BenchmarkSin-128                     29.5          23.7          -19.66%
    BenchmarkTan-128                     32.8          27.4          -16.46%
    
    Fixes #18922
    
    Change-Id: I8e1cf14d3880d7cd720dc5188dd174cba1f7fef7
    Reviewed-on: https://go-review.googlesource.com/36725
    Reviewed-by: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3123df3464ff8df1d15452b51360e1b7f05dbcd3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Mar 1 05:42:32 2017 +0000

    net/http: fix flaky TestClientRedirect308NoLocation
    
    This was a t.Parallel test but it was using the global DefaultTransport
    via the global Get func.
    
    Use a private Transport that won't have its CloseIdleConnections etc
    methods called by other tests.
    
    (I hit this flake myself while testing a different change.)
    
    Change-Id: If0665e3e8580ee198f8e5f3079bfaea55f036eca
    Reviewed-on: https://go-review.googlesource.com/37624
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 6d32b1a3431dba378d22fa8b187ccfeab259cbe2
Author: Kevin Burke <kev@inburke.com>
Date:   Tue Feb 28 16:28:34 2017 -0800

    os: add OpenFile example for appending data
    
    Fixes #19329.
    
    Change-Id: I6d8bb112a56d751a6d3ea9bd6021803cb9f59234
    Reviewed-on: https://go-review.googlesource.com/37619
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9bd1cc3fa1145182e9ce041d0e96bd2051cd7fcf
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Mar 1 13:29:17 2017 +0000

    Revert "cmd/vet/all: remove pprof from the whitelist"
    
    This reverts commit 12b6c18139233abd7b1af1fc0a07279d56df3642.
    
    Reason for revert: Broke vet builder. #19322 was not fully fixed.
    
    Change-Id: Id85131d4d0b8915480d65e3532da62b769463d70
    Reviewed-on: https://go-review.googlesource.com/37625
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 5e90bbcc6de54ecbad1d72bc8e71e167829069b5
Author: Heschi Kreinick <heschi@google.com>
Date:   Mon Feb 27 18:26:33 2017 -0500

    testing: fix Benchmark() to start at 1 iteration, not 100
    
    The run1 call removed in golang.org/cl/36990 was necessary to
    initialize the duration of the benchmark. With it gone, the math in
    launch() starts from 100. This doesn't work out well for second-long
    benchmark methods. Put it back.
    
    Updates #18815
    
    Change-Id: I461f3466c805d0c61124a2974662f7ad45335794
    Reviewed-on: https://go-review.googlesource.com/37530
    Run-TryBot: Heschi Kreinick <heschi@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>

commit 29f061960d5008170541a886feab721bf754f0fd
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Feb 8 12:47:43 2017 +1100

    cmd/link: write dwarf sections
    
    Also stop skipping TestExternalLinkerDWARF and
    TestDefaultLinkerDWARF.
    
    Fixes #10776.
    
    Change-Id: Ia596a684132e3cdee59ce5539293eedc1752fe5a
    Reviewed-on: https://go-review.googlesource.com/36983
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit aada49038c683d048fd0a146366d7ce52dc17e97
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Feb 8 12:30:30 2017 +1100

    cmd/link: write dwarf relocations
    
    For #10776.
    
    Change-Id: I11dd441d8e5d6316889ffa8418df8b58c57c677d
    Reviewed-on: https://go-review.googlesource.com/36982
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 15442178c801476f873b0678a99b27f06c8e38d6
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Feb 28 15:01:38 2017 -0800

    os: don't use waitid on Darwin
    
    According to issue #19314 waitid on Darwin returns if the process is
    stopped, even though we specify WEXITED.
    
    Fixes #19314.
    
    Change-Id: I95faf196c11e43b7741efff79351bab45c811bc2
    Reviewed-on: https://go-review.googlesource.com/37610
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d945b286758d034b3bb07cbf3a4055b90684c38b
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Mar 1 08:41:34 2017 +1100

    cmd/compile/internal/ssa: remove unused PrintFunc variable
    
    Change-Id: I8c581eec77beacaddc0aac29e7d380a4d5ca8acc
    Reviewed-on: https://go-review.googlesource.com/37551
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c861a4c78695038832bb4cf92a47c5e09566721b
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 28 13:38:04 2017 -0800

    go/internal/srcimporter: parse files concurrently (fixes TODO)
    
    Passes go test -race.
    
    Change-Id: I14b5b1b1a8ad1e43d60013823d71d78a6519581f
    Reviewed-on: https://go-review.googlesource.com/37588
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b6c600fc9a75fd6f4b6f4478058b95902ae6be94
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Feb 28 15:51:29 2017 -0800

    cmd/compile/internal/gc: separate builtin and real runtime packages
    
    The builtin runtime package definitions intentionally diverge from the
    actual runtime package's, but this only works as long as they never
    overlap.
    
    To make it easier to expand the builtin runtime package, this CL now
    loads their definitions into a logically separate "go.runtime"
    package.  By resetting the package's Prefix field to "runtime", any
    references to builtin definitions will still resolve against the real
    package runtime.
    
    Fixes #14482.
    
    Passes toolstash -cmp.
    
    Change-Id: I539c0994deaed4506a331f38c5b4d6bc8c95433f
    Reviewed-on: https://go-review.googlesource.com/37538
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 12b6c18139233abd7b1af1fc0a07279d56df3642
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Mar 1 00:26:16 2017 +0000

    cmd/vet/all: remove pprof from the whitelist
    
    Updates #19322
    
    Change-Id: I610f40d874f499e52db3356a3da54538dac55242
    Reviewed-on: https://go-review.googlesource.com/37618
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 218313555473ebd2c47bf2b1bb9aee70a7f0164a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 6 10:55:39 2017 -0800

    cmd/compile: recognize bit test patterns on amd64
    
    Updates #18943
    
    Change-Id: If3080d6133bb6d2710b57294da24c90251ab4e08
    Reviewed-on: https://go-review.googlesource.com/36329
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit ac7761e1a4c04154b557fa2294e0778bbf9373f9
Author: Heschi Kreinick <heschi@google.com>
Date:   Fri Feb 17 16:52:16 2017 -0500

    cmd/compile, cmd/asm: remove Link.Plists
    
    Link.Plists never contained more than one Plist, and sometimes none.
    Passing around the Plist being worked on is straightforward and makes
    the data flow easier to follow.
    
    Change-Id: I79cb30cb2bd3d319fdbb1dfa5d35b27fcb748e5c
    Reviewed-on: https://go-review.googlesource.com/37169
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ac4a86523c2521555b9ea104157fcc8cf5ce79f5
Author: Raul Silvera <rsilvera@google.com>
Date:   Tue Feb 28 16:07:36 2017 -0800

    cmd/vendor/github.com/google/pprof: refresh from upstream
    
    Updating to commit e41fb7133e7ebb84ba6af2f6443032c728db26d3
    from github.com/google/pprof
    
    This fixes #19322
    
    Change-Id: Ia1c008a09f46ed19ef176046e38868eacb715682
    Reviewed-on: https://go-review.googlesource.com/37617
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bca0320641000d842341637f22f140c262adb360
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 24 18:13:29 2017 -0800

    compress/flate: use math/bits.Reverse8/16 instead of local implementation
    
    No measurable impact on performance (specifically, no degradation).
    Reverse is used in Huffman en/de-coding. For completeness, here are
    all the speed-related benchmark results:
    
    name                             old time/op    new time/op    delta
    Decode/Digits/Huffman/1e4-8         181µs ± 0%     178µs ± 1%   ~             (p=0.100 n=3+3)
    Decode/Digits/Huffman/1e5-8        1.60ms ± 3%    1.56ms ± 3%   ~             (p=0.400 n=3+3)
    Decode/Digits/Huffman/1e6-8        15.7ms ± 1%    15.3ms ± 3%   ~             (p=0.700 n=3+3)
    Decode/Digits/Speed/1e4-8           179µs ± 0%     180µs ± 0%   ~             (p=0.200 n=3+3)
    Decode/Digits/Speed/1e5-8          1.68ms ± 0%    1.66ms ± 3%   ~             (p=0.700 n=3+3)
    Decode/Digits/Speed/1e6-8          16.6ms ± 2%    16.6ms ± 5%   ~             (p=0.700 n=3+3)
    Decode/Digits/Default/1e4-8         179µs ± 1%     178µs ± 1%   ~             (p=0.700 n=3+3)
    Decode/Digits/Default/1e5-8        1.62ms ± 3%    1.62ms ± 4%   ~             (p=1.000 n=3+3)
    Decode/Digits/Default/1e6-8        16.0ms ± 2%    16.0ms ± 3%   ~             (p=1.000 n=3+3)
    Decode/Digits/Compression/1e4-8     179µs ± 1%     179µs ± 0%   ~             (p=0.200 n=3+3)
    Decode/Digits/Compression/1e5-8    1.62ms ± 2%    1.62ms ± 3%   ~             (p=1.000 n=3+3)
    Decode/Digits/Compression/1e6-8    16.1ms ± 3%    16.0ms ± 3%   ~             (p=1.000 n=3+3)
    Decode/Twain/Huffman/1e4-8          205µs ± 2%     207µs ± 1%   ~             (p=1.000 n=3+3)
    Decode/Twain/Huffman/1e5-8         1.77ms ± 2%    1.77ms ± 4%   ~             (p=0.700 n=3+3)
    Decode/Twain/Huffman/1e6-8         17.4ms ± 2%    17.4ms ± 3%   ~             (p=1.000 n=3+3)
    Decode/Twain/Speed/1e4-8            186µs ± 1%     186µs ± 1%   ~             (p=0.400 n=3+3)
    Decode/Twain/Speed/1e5-8           1.53ms ± 2%    1.52ms ± 0%   ~             (p=0.700 n=3+3)
    Decode/Twain/Speed/1e6-8           14.9ms ± 1%    14.8ms ± 1%   ~             (p=1.000 n=3+3)
    Decode/Twain/Default/1e4-8          176µs ± 1%     174µs ± 0%   ~             (p=0.200 n=3+3)
    Decode/Twain/Default/1e5-8         1.30ms ± 2%    1.31ms ± 1%   ~             (p=0.700 n=3+3)
    Decode/Twain/Default/1e6-8         12.6ms ± 3%    12.5ms ± 0%   ~             (p=0.700 n=3+3)
    Decode/Twain/Compression/1e4-8      177µs ± 0%     174µs ± 1%   ~             (p=0.100 n=3+3)
    Decode/Twain/Compression/1e5-8     1.30ms ± 1%    1.31ms ± 0%   ~             (p=0.700 n=3+3)
    Decode/Twain/Compression/1e6-8     12.5ms ± 1%    12.5ms ± 1%   ~             (p=1.000 n=3+3)
    Encode/Digits/Huffman/1e4-8        47.4µs ± 1%    46.5µs ± 0%   ~             (p=0.100 n=3+3)
    Encode/Digits/Huffman/1e5-8         453µs ± 2%     446µs ± 1%   ~             (p=0.700 n=3+3)
    Encode/Digits/Huffman/1e6-8        4.44ms ± 3%    4.39ms ± 0%   ~             (p=1.000 n=3+3)
    Encode/Digits/Speed/1e4-8           190µs ± 4%     185µs ± 0%   ~             (p=0.100 n=3+3)
    Encode/Digits/Speed/1e5-8          1.78ms ± 5%    1.75ms ± 1%   ~             (p=1.000 n=3+3)
    Encode/Digits/Speed/1e6-8          17.9ms ± 7%    17.3ms ± 1%   ~             (p=0.400 n=3+3)
    Encode/Digits/Default/1e4-8         366µs ± 1%     361µs ± 0%   ~             (p=0.200 n=3+3)
    Encode/Digits/Default/1e5-8        5.58ms ± 5%    5.44ms ± 1%   ~             (p=0.400 n=3+3)
    Encode/Digits/Default/1e6-8        59.0ms ± 3%    58.2ms ± 1%   ~             (p=0.700 n=3+3)
    Encode/Digits/Compression/1e4-8     369µs ± 3%     362µs ± 0%   ~             (p=0.100 n=3+3)
    Encode/Digits/Compression/1e5-8    5.50ms ± 2%    5.47ms ± 1%   ~             (p=1.000 n=3+3)
    Encode/Digits/Compression/1e6-8    59.4ms ± 2%    58.5ms ± 1%   ~             (p=0.400 n=3+3)
    Encode/Twain/Huffman/1e4-8         64.4µs ± 3%    64.7µs ± 1%   ~             (p=0.700 n=3+3)
    Encode/Twain/Huffman/1e5-8          526µs ± 1%     526µs ± 2%   ~             (p=1.000 n=3+3)
    Encode/Twain/Huffman/1e6-8         5.18ms ± 2%    5.17ms ± 1%   ~             (p=0.700 n=3+3)
    Encode/Twain/Speed/1e4-8            206µs ± 1%     204µs ± 0%   ~             (p=0.100 n=3+3)
    Encode/Twain/Speed/1e5-8           1.73ms ± 2%    1.70ms ± 0%   ~             (p=0.100 n=3+3)
    Encode/Twain/Speed/1e6-8           16.7ms ± 0%    16.7ms ± 1%   ~             (p=1.000 n=3+3)
    Encode/Twain/Default/1e4-8          423µs ± 3%     418µs ± 1%   ~             (p=1.000 n=3+3)
    Encode/Twain/Default/1e5-8         6.34ms ± 4%    6.23ms ± 0%   ~             (p=1.000 n=3+3)
    Encode/Twain/Default/1e6-8         68.0ms ± 3%    67.5ms ± 0%   ~             (p=0.700 n=3+3)
    Encode/Twain/Compression/1e4-8      435µs ± 3%     424µs ± 0%   ~             (p=0.700 n=3+3)
    Encode/Twain/Compression/1e5-8     7.01ms ± 1%    6.92ms ± 0%   ~             (p=0.100 n=3+3)
    Encode/Twain/Compression/1e6-8     77.1ms ± 4%    75.5ms ± 1%   ~             (p=0.400 n=3+3)
    
    name                             old speed      new speed      delta
    Decode/Digits/Huffman/1e4-8      55.2MB/s ± 0%  56.2MB/s ± 1%   ~             (p=0.100 n=3+3)
    Decode/Digits/Huffman/1e5-8      62.4MB/s ± 3%  64.1MB/s ± 3%   ~             (p=0.400 n=3+3)
    Decode/Digits/Huffman/1e6-8      63.8MB/s ± 1%  65.3MB/s ± 3%   ~             (p=0.700 n=3+3)
    Decode/Digits/Speed/1e4-8        55.8MB/s ± 0%  55.4MB/s ± 0%   ~             (p=0.200 n=3+3)
    Decode/Digits/Speed/1e5-8        59.6MB/s ± 0%  60.3MB/s ± 3%   ~             (p=0.700 n=3+3)
    Decode/Digits/Speed/1e6-8        60.1MB/s ± 2%  60.3MB/s ± 4%   ~             (p=0.700 n=3+3)
    Decode/Digits/Default/1e4-8      55.8MB/s ± 1%  56.1MB/s ± 1%   ~             (p=0.700 n=3+3)
    Decode/Digits/Default/1e5-8      61.8MB/s ± 3%  61.7MB/s ± 4%   ~             (p=1.000 n=3+3)
    Decode/Digits/Default/1e6-8      62.4MB/s ± 2%  62.4MB/s ± 3%   ~             (p=1.000 n=3+3)
    Decode/Digits/Compression/1e4-8  55.7MB/s ± 1%  56.0MB/s ± 0%   ~             (p=0.300 n=3+3)
    Decode/Digits/Compression/1e5-8  61.7MB/s ± 2%  61.9MB/s ± 3%   ~             (p=1.000 n=3+3)
    Decode/Digits/Compression/1e6-8  62.2MB/s ± 3%  62.6MB/s ± 3%   ~             (p=1.000 n=3+3)
    Decode/Twain/Huffman/1e4-8       48.8MB/s ± 2%  48.4MB/s ± 1%   ~             (p=1.000 n=3+3)
    Decode/Twain/Huffman/1e5-8       56.4MB/s ± 2%  56.6MB/s ± 4%   ~             (p=0.700 n=3+3)
    Decode/Twain/Huffman/1e6-8       57.6MB/s ± 2%  57.5MB/s ± 3%   ~             (p=1.000 n=3+3)
    Decode/Twain/Speed/1e4-8         53.7MB/s ± 1%  53.9MB/s ± 1%   ~             (p=0.400 n=3+3)
    Decode/Twain/Speed/1e5-8         65.5MB/s ± 2%  65.6MB/s ± 0%   ~             (p=0.700 n=3+3)
    Decode/Twain/Speed/1e6-8         66.9MB/s ± 1%  67.4MB/s ± 1%   ~             (p=1.000 n=3+3)
    Decode/Twain/Default/1e4-8       56.9MB/s ± 1%  57.3MB/s ± 0%   ~             (p=0.200 n=3+3)
    Decode/Twain/Default/1e5-8       77.2MB/s ± 2%  76.6MB/s ± 1%   ~             (p=0.700 n=3+3)
    Decode/Twain/Default/1e6-8       79.3MB/s ± 3%  80.0MB/s ± 0%   ~             (p=0.700 n=3+3)
    Decode/Twain/Compression/1e4-8   56.4MB/s ± 0%  57.5MB/s ± 1%   ~             (p=0.100 n=3+3)
    Decode/Twain/Compression/1e5-8   76.8MB/s ± 1%  76.5MB/s ± 0%   ~             (p=0.700 n=3+3)
    Decode/Twain/Compression/1e6-8   80.1MB/s ± 1%  79.8MB/s ± 1%   ~             (p=1.000 n=3+3)
    Encode/Digits/Huffman/1e4-8       211MB/s ± 1%   215MB/s ± 0%   ~             (p=0.100 n=3+3)
    Encode/Digits/Huffman/1e5-8       221MB/s ± 2%   224MB/s ± 1%   ~             (p=0.700 n=3+3)
    Encode/Digits/Huffman/1e6-8       225MB/s ± 3%   228MB/s ± 0%   ~             (p=1.000 n=3+3)
    Encode/Digits/Speed/1e4-8        52.8MB/s ± 4%  54.1MB/s ± 0%   ~             (p=0.100 n=3+3)
    Encode/Digits/Speed/1e5-8        56.2MB/s ± 5%  57.0MB/s ± 1%   ~             (p=1.000 n=3+3)
    Encode/Digits/Speed/1e6-8        56.0MB/s ± 6%  57.7MB/s ± 1%   ~             (p=0.400 n=3+3)
    Encode/Digits/Default/1e4-8      27.3MB/s ± 1%  27.7MB/s ± 0%   ~             (p=0.200 n=3+3)
    Encode/Digits/Default/1e5-8      17.9MB/s ± 4%  18.4MB/s ± 1%   ~             (p=0.400 n=3+3)
    Encode/Digits/Default/1e6-8      17.0MB/s ± 3%  17.2MB/s ± 1%   ~             (p=0.500 n=3+3)
    Encode/Digits/Compression/1e4-8  27.1MB/s ± 3%  27.6MB/s ± 0%   ~             (p=0.100 n=3+3)
    Encode/Digits/Compression/1e5-8  18.2MB/s ± 2%  18.3MB/s ± 1%   ~             (p=1.000 n=3+3)
    Encode/Digits/Compression/1e6-8  16.9MB/s ± 2%  17.1MB/s ± 1%   ~             (p=0.400 n=3+3)
    Encode/Twain/Huffman/1e4-8        155MB/s ± 3%   155MB/s ± 1%   ~             (p=0.700 n=3+3)
    Encode/Twain/Huffman/1e5-8        190MB/s ± 1%   190MB/s ± 2%   ~             (p=1.000 n=3+3)
    Encode/Twain/Huffman/1e6-8        193MB/s ± 2%   193MB/s ± 1%   ~             (p=0.700 n=3+3)
    Encode/Twain/Speed/1e4-8         48.5MB/s ± 1%  49.1MB/s ± 0%   ~             (p=0.100 n=3+3)
    Encode/Twain/Speed/1e5-8         57.7MB/s ± 2%  59.0MB/s ± 0%   ~             (p=0.100 n=3+3)
    Encode/Twain/Speed/1e6-8         59.7MB/s ± 0%  59.7MB/s ± 1%   ~             (p=1.000 n=3+3)
    Encode/Twain/Default/1e4-8       23.6MB/s ± 3%  23.9MB/s ± 1%   ~             (p=1.000 n=3+3)
    Encode/Twain/Default/1e5-8       15.8MB/s ± 4%  16.1MB/s ± 0%   ~             (p=1.000 n=3+3)
    Encode/Twain/Default/1e6-8       14.7MB/s ± 3%  14.8MB/s ± 0%   ~             (p=0.700 n=3+3)
    Encode/Twain/Compression/1e4-8   23.0MB/s ± 3%  23.6MB/s ± 0%   ~             (p=0.700 n=3+3)
    Encode/Twain/Compression/1e5-8   14.3MB/s ± 1%  14.5MB/s ± 0%   ~             (p=0.100 n=3+3)
    Encode/Twain/Compression/1e6-8   13.0MB/s ± 4%  13.2MB/s ± 1%   ~             (p=0.400 n=3+3)
    
    Measured on a "quiet" (no browser running) 2.3 GHz Intel Core i7, running macOS 10.12.3.
    
    See also #19279.
    
    Change-Id: Ice759eb34eb37442b543957447c264e0aadc1fa9
    Reviewed-on: https://go-review.googlesource.com/37460
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 32b41c8dc75a731e4053b59b19c542a79eb56c1f
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 28 14:55:00 2017 -0800

    math/bits: move left-over functionality from bits_impl.go to bits.go
    
    Removes an extra function call for TrailingZeroes and thus may
    increase chances for inlining.
    
    Change-Id: Iefd8d4402dc89b64baf4e5c865eb3dadade623af
    Reviewed-on: https://go-review.googlesource.com/37613
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 09294ab75461c088e3676e1575acb44c38371b08
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Feb 28 15:24:18 2017 -0800

    cmd/vet/all: disable cgo when running 'go install'
    
    Change-Id: Iab1e84624c0288ebdd33fbe83bd60948b5d91fc4
    Reviewed-on: https://go-review.googlesource.com/37612
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e73f4894949c4ced611881329ff8f37805152585
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Feb 28 22:25:06 2017 +0000

    os/exec: remove duplicate environment variables in Cmd.Start
    
    Nobody intends to have duplicates anyway because it's so undefined
    and everything handles it so poorly.
    
    Removing duplicates automatically simplifies code and makes existing
    code do what people already expect.
    
    Fixes #12868
    
    Change-Id: I95eeba8c59ff94d0f018012a6f4e031aaabfd5d9
    Reviewed-on: https://go-review.googlesource.com/37586
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3c023f75a62f903273c688432f95e77fc945b6fb
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Tue Feb 28 21:21:45 2017 +0100

    strings: fix handling of invalid UTF-8 sequences in Map
    
    The new Map implementation introduced in golang.org/cl/33201
    did not differentiate if an invalid UTF-8 sequence was decoded
    or the RuneError rune. It would therefore always advance by
    3 bytes (which is the length of the RuneError rune) instead
    of 1 for an invalid sequences. This cl adds a check to correctly
    determine the length of bytes needed to advance to the next rune.
    
    Fixes #19330.
    
    Change-Id: I1e7f9333f3ef6068ffc64015bb0a9f32b0b7111d
    Reviewed-on: https://go-review.googlesource.com/37597
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0fe58bf6500a615e934e1ab2a7921dfbc2803cea
Author: Keith Randall <khr@golang.org>
Date:   Tue Feb 28 14:01:59 2017 -0800

    cmd/compile: simplify load+op rules
    
    There's no need to use @block rules, as canMergeLoad makes sure that
    the load and op are already in the same block.
    With no @block needed, we also don't need to set the type explicitly.
    It can just be inherited from the op being rewritten.
    
    Noticed while working on #19284.
    
    Change-Id: Ied8bcc8058260118ff7e166093112e29107bcb7e
    Reviewed-on: https://go-review.googlesource.com/37585
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 45055f21ab4b6005446e8d680f315ece410e75b5
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 28 13:18:55 2017 -0800

    go/types: implement SizesFor convenience function
    
    SizesFor returns a Sizes implementation for a supported architecture.
    Use functionality in srcimporter.
    
    Change-Id: I197e641b419c678030dfaab5c5b8c569fd0410f3
    Reviewed-on: https://go-review.googlesource.com/37583
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 83bc4a2feed1c7dc37026278364772483fe73618
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 28 11:52:43 2017 -0800

    math/bits: faster LeadingZeros and Len functions
    
    benchmark                     old ns/op     new ns/op     delta
    BenchmarkLeadingZeros-8       8.43          3.10          -63.23%
    BenchmarkLeadingZeros8-8      8.13          1.33          -83.64%
    BenchmarkLeadingZeros16-8     7.34          2.07          -71.80%
    BenchmarkLeadingZeros32-8     7.99          2.87          -64.08%
    BenchmarkLeadingZeros64-8     8.13          2.96          -63.59%
    
    Measured on 2.3 GHz Intel Core i7 running macOS 10.12.3.
    
    Change-Id: Id343531b408d42ac45f10c76f60e85bdb977f91e
    Reviewed-on: https://go-review.googlesource.com/37582
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9515cb511a1210e013c26354ea09e786acd61365
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 28 10:20:32 2017 -0800

    math/bits: faster TrailingZeroes8
    
    For sizes > 8, the existing code is faster.
    
    benchmark                     old ns/op     new ns/op     delta
    BenchmarkTrailingZeros8-8     1.95          1.29          -33.85%
    
    Measured on 2.3 GHz Intel Core i7 running macOS 10.12.3.
    
    Change-Id: I6f3a33ec633a2c544ec29693c141f2f99335c745
    Reviewed-on: https://go-review.googlesource.com/37581
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d7a659b11b9ab9132fd4302ffe9250b30bbe431e
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 28 10:09:44 2017 -0800

    math/bits: faster OnesCount using table lookups for sizes 8,16,32
    
    For uint64, the existing algorithm is faster.
    
    benchmark                  old ns/op     new ns/op     delta
    BenchmarkOnesCount8-8      1.95          0.97          -50.26%
    BenchmarkOnesCount16-8     2.54          1.39          -45.28%
    BenchmarkOnesCount32-8     2.61          1.96          -24.90%
    
    Measured on 2.3 GHz Intel Core i7 running macOS 10.12.3.
    
    Change-Id: I6cc42882fef3d24694720464039161e339a9ae99
    Reviewed-on: https://go-review.googlesource.com/37580
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 064e44f218f62247e894733d861208257102b0eb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Feb 14 11:00:34 2017 -0800

    runtime: evacuate old map buckets more consistently
    
    During map growth, buckets are evacuated in two ways.
    When a value is altered, its containing bucket is evacuated.
    Also, an evacuation mark is maintained and advanced every time.
    Prior to this CL, the evacuation mark was always incremented,
    even if the next bucket to be evacuated had already been evacuated.
    This CL changes evacuation mark advancement to skip previously
    evacuated buckets. This has the effect of making map evacuation both
    more aggressive and more consistent.
    
    Aggressive map evacuation is good. While the map is growing,
    map accesses must check two buckets, which may be far apart in memory.
    Map growth also delays garbage collection.
    And if map evacuation is not aggressive enough, there is a risk that
    a populate-once read-many map may be stuck permanently in map growth.
    This CL does not eliminate that possibility, but it shrinks the window.
    
    There is minimal impact on map benchmarks:
    
    name                         old time/op    new time/op    delta
    MapPop100-8                    12.4µs ±11%    12.4µs ± 7%    ~     (p=0.798 n=15+15)
    MapPop1000-8                    240µs ± 8%     235µs ± 8%    ~     (p=0.217 n=15+14)
    MapPop10000-8                  4.49ms ±10%    4.51ms ±15%    ~     (p=1.000 n=15+13)
    MegMap-8                       11.9ns ± 2%    11.8ns ± 0%  -1.01%  (p=0.000 n=15+11)
    MegOneMap-8                    9.30ns ± 1%    9.29ns ± 1%    ~     (p=0.955 n=14+14)
    MegEqMap-8                     31.9µs ± 5%    31.9µs ± 3%    ~     (p=0.935 n=15+15)
    MegEmptyMap-8                  2.41ns ± 2%    2.41ns ± 0%    ~     (p=0.594 n=12+14)
    SmallStrMap-8                  12.8ns ± 1%    12.7ns ± 1%    ~     (p=0.569 n=14+13)
    MapStringKeysEight_16-8        13.6ns ± 1%    13.7ns ± 2%    ~     (p=0.100 n=13+15)
    MapStringKeysEight_32-8        12.1ns ± 1%    12.1ns ± 2%    ~     (p=0.340 n=15+15)
    MapStringKeysEight_64-8        12.1ns ± 1%    12.1ns ± 2%    ~     (p=0.582 n=15+14)
    MapStringKeysEight_1M-8        12.0ns ± 1%    12.1ns ± 1%    ~     (p=0.267 n=15+14)
    IntMap-8                       7.96ns ± 1%    7.97ns ± 2%    ~     (p=0.991 n=15+13)
    RepeatedLookupStrMapKey32-8    15.8ns ± 2%    15.8ns ± 1%    ~     (p=0.393 n=15+14)
    RepeatedLookupStrMapKey1M-8    35.3µs ± 2%    35.3µs ± 1%    ~     (p=0.815 n=15+15)
    NewEmptyMap-8                  36.0ns ± 4%    36.4ns ± 7%    ~     (p=0.270 n=15+15)
    NewSmallMap-8                  85.5ns ± 1%    85.6ns ± 1%    ~     (p=0.674 n=14+15)
    MapIter-8                      89.9ns ± 6%    90.8ns ± 6%    ~     (p=0.467 n=15+15)
    MapIterEmpty-8                 10.0ns ±22%    10.0ns ±25%    ~     (p=0.846 n=15+15)
    SameLengthMap-8                4.18ns ± 1%    4.17ns ± 1%    ~     (p=0.653 n=15+14)
    BigKeyMap-8                    20.2ns ± 1%    20.1ns ± 1%  -0.82%  (p=0.002 n=15+15)
    BigValMap-8                    22.5ns ± 8%    22.3ns ± 6%    ~     (p=0.615 n=15+15)
    SmallKeyMap-8                  15.3ns ± 1%    15.3ns ± 1%    ~     (p=0.754 n=15+14)
    ComplexAlgMap-8                58.4ns ± 1%    58.7ns ± 1%  +0.52%  (p=0.000 n=14+15)
    
    There is a tiny but detectable difference in the compiler:
    
    name       old time/op      new time/op      delta
    Template        218ms ± 5%       219ms ± 4%    ~     (p=0.094 n=98+98)
    Unicode        93.6ms ± 5%      93.6ms ± 4%    ~     (p=0.910 n=94+95)
    GoTypes         596ms ± 5%       598ms ± 6%    ~     (p=0.533 n=98+100)
    Compiler        2.72s ± 3%       2.72s ± 4%    ~     (p=0.238 n=100+99)
    SSA             4.11s ± 3%       4.11s ± 3%    ~     (p=0.864 n=99+98)
    Flate           129ms ± 6%       129ms ± 4%    ~     (p=0.522 n=98+96)
    GoParser        151ms ± 4%       151ms ± 4%  -0.48%  (p=0.017 n=96+96)
    Reflect         379ms ± 3%       376ms ± 4%  -0.57%  (p=0.011 n=99+99)
    Tar             112ms ± 5%       112ms ± 6%    ~     (p=0.688 n=93+95)
    XML             214ms ± 4%       214ms ± 5%    ~     (p=0.968 n=100+99)
    StdCmd          16.2s ± 2%       16.2s ± 2%  -0.26%  (p=0.048 n=99+99)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   252user-ms ± 4%  250user-ms ± 4%  -0.63%  (p=0.020 n=98+97)
    Unicode    113user-ms ± 7%  114user-ms ± 5%    ~     (p=0.057 n=97+94)
    GoTypes    776user-ms ± 5%  777user-ms ± 5%    ~     (p=0.375 n=97+96)
    Compiler   3.61user-s ± 3%  3.60user-s ± 3%    ~     (p=0.445 n=98+93)
    SSA        5.84user-s ± 6%  5.85user-s ± 5%    ~     (p=0.542 n=100+95)
    Flate      154user-ms ± 5%  154user-ms ± 5%    ~     (p=0.699 n=99+99)
    GoParser   184user-ms ± 6%  183user-ms ± 4%    ~     (p=0.557 n=98+95)
    Reflect    461user-ms ± 5%  462user-ms ± 4%    ~     (p=0.853 n=97+99)
    Tar        130user-ms ± 5%  129user-ms ± 6%    ~     (p=0.567 n=93+100)
    XML        257user-ms ± 6%  258user-ms ± 6%    ~     (p=0.205 n=99+100)
    
    Change-Id: Id92dd54a152904069aac415e6aaaab5c67f5f476
    Reviewed-on: https://go-review.googlesource.com/37011
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9bc67bb4f4caba59443c504e09758812ac63a046
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Feb 28 11:34:56 2017 -0800

    cmd/internal/obj: remove unused Getcallerpc function
    
    Change-Id: I0c7b677657326f318e906e109cbda0cfa78c4973
    Reviewed-on: https://go-review.googlesource.com/37537
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 379567aad19e97eb2145d7bbe45a7a2febccd3e4
Author: philhofer <phofer@umich.edu>
Date:   Mon Feb 20 08:43:54 2017 -0800

    cmd/compile/ssa: more aggressive constant folding
    
    Add rewrite rules that canonicalize the location
    of constants in expressions, and fold conststants
    that appear in operations that can be trivially
    reassociated.
    
    After this change, the compiler constant-folds
    expressions like "4 + x - 1" and "4 & x & 1"
    
    Benchmarks affected on darwin/amd64:
    
    name                     old time/op    new time/op    delta
    FmtFprintfInt-8            82.1ns ± 1%    81.7ns ± 1%  -0.46%  (p=0.023 n=8+9)
    FmtFprintfIntInt-8          122ns ± 2%     120ns ± 2%  -1.48%  (p=0.047 n=10+10)
    FmtManyArgs-8               493ns ± 0%     486ns ± 1%  -1.37%  (p=0.000 n=8+10)
    Gzip-8                      230ms ± 0%     229ms ± 1%  -0.46%  (p=0.001 n=10+9)
    HTTPClientServer-8         74.5µs ± 1%    73.7µs ± 1%  -1.11%  (p=0.000 n=10+10)
    JSONDecode-8               51.7ms ± 0%    51.9ms ± 1%  +0.42%  (p=0.017 n=10+9)
    RegexpMatchEasy0_32-8      82.6ns ± 1%    81.7ns ± 0%  -1.02%  (p=0.000 n=9+8)
    RegexpMatchMedium_32-8      121ns ± 1%     120ns ± 1%  -1.48%  (p=0.001 n=10+10)
    Revcomp-8                   426ms ± 1%     400ms ± 1%  -6.16%  (p=0.000 n=10+10)
    TimeFormat-8                330ns ± 1%     327ns ± 0%  -0.82%  (p=0.000 n=10+10)
    
    name                     old speed      new speed      delta
    Gzip-8                   84.4MB/s ± 0%  84.8MB/s ± 1%  +0.47%  (p=0.001 n=10+9)
    JSONDecode-8             37.6MB/s ± 0%  37.4MB/s ± 1%  -0.42%  (p=0.016 n=10+9)
    RegexpMatchEasy0_32-8     387MB/s ± 1%   392MB/s ± 0%  +1.06%  (p=0.000 n=9+8)
    RegexpMatchMedium_32-8   8.21MB/s ± 1%  8.34MB/s ± 1%  +1.58%  (p=0.000 n=10+9)
    Revcomp-8                 597MB/s ± 1%   636MB/s ± 1%  +6.57%  (p=0.000 n=10+10)
    
    Change-Id: Ie37ff91605b76a984a8400dfd1e34f50bf61c864
    Reviewed-on: https://go-review.googlesource.com/37290
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 504bc3ed24765294cf3d665a68d57a6e4fc7d23a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Feb 3 22:40:56 2017 -0800

    cmd/compile, runtime: specialize convT2x, don't alloc for zero vals
    
    Prior to this CL, all runtime conversions
    from a concrete value to an interface went
    through one of two runtime calls: convT2E or convT2I.
    However, in practice, basic types are very common.
    Specializing convT2x for those basic types allows
    for a more efficient implementation for those types.
    For basic scalars and strings, allocation and copying
    can use the same methods as normal code.
    For pointer-free types, allocation can occur without
    zeroing, and copying can take place without GC calls.
    For slices, copying is cheaper and simpler.
    
    This CL adds twelve runtime routines:
    
    convT2E16, convT2I16
    convT2E32, convT2I32
    convT2E64, convT2I64
    convT2Estring, convT2Istring
    convT2Eslice, convT2Islice
    convT2Enoptr, convT2Inoptr
    
    While compiling make.bash, 93% of all convT2x calls
    are now to one of these specialized convT2x call.
    
    Within specialized convT2x routines, it is cheap to check
    for a zero value, in a way that it is not in general.
    When we detect a zero value there, we return a pointer
    to zeroVal, rather than allocating.
    
    name                         old time/op  new time/op  delta
    ConvT2Ezero/zero/16-8        17.9ns ± 2%   3.0ns ± 3%  -83.20%  (p=0.000 n=56+56)
    ConvT2Ezero/zero/32-8        17.8ns ± 2%   3.0ns ± 3%  -83.15%  (p=0.000 n=59+60)
    ConvT2Ezero/zero/64-8        20.1ns ± 1%   3.0ns ± 2%  -84.98%  (p=0.000 n=57+57)
    ConvT2Ezero/zero/str-8       32.6ns ± 1%   3.0ns ± 4%  -90.70%  (p=0.000 n=59+60)
    ConvT2Ezero/zero/slice-8     36.7ns ± 2%   3.0ns ± 2%  -91.78%  (p=0.000 n=59+59)
    ConvT2Ezero/zero/big-8       91.9ns ± 2%  85.9ns ± 2%   -6.52%  (p=0.000 n=57+57)
    ConvT2Ezero/nonzero/16-8     17.7ns ± 2%  12.7ns ± 3%  -28.38%  (p=0.000 n=55+60)
    ConvT2Ezero/nonzero/32-8     17.8ns ± 1%  12.7ns ± 1%  -28.44%  (p=0.000 n=54+57)
    ConvT2Ezero/nonzero/64-8     20.0ns ± 1%  15.0ns ± 1%  -24.90%  (p=0.000 n=56+58)
    ConvT2Ezero/nonzero/str-8    32.6ns ± 1%  25.7ns ± 1%  -21.17%  (p=0.000 n=58+55)
    ConvT2Ezero/nonzero/slice-8  36.8ns ± 2%  30.4ns ± 1%  -17.32%  (p=0.000 n=60+52)
    ConvT2Ezero/nonzero/big-8    92.1ns ± 2%  85.9ns ± 2%   -6.70%  (p=0.000 n=57+59)
    
    Benchmarks on a real program (the compiler):
    
    name       old time/op      new time/op      delta
    Template        227ms ± 5%       221ms ± 2%  -2.48%  (p=0.000 n=30+26)
    Unicode         102ms ± 5%       100ms ± 3%  -1.30%  (p=0.009 n=30+26)
    GoTypes         656ms ± 5%       659ms ± 4%    ~     (p=0.208 n=30+30)
    Compiler        2.82s ± 2%       2.82s ± 1%    ~     (p=0.614 n=29+27)
    Flate           128ms ± 2%       128ms ± 5%    ~     (p=0.783 n=27+28)
    GoParser        158ms ± 3%       158ms ± 3%    ~     (p=0.261 n=28+30)
    Reflect         408ms ± 7%       401ms ± 3%    ~     (p=0.075 n=30+30)
    Tar             123ms ± 6%       121ms ± 8%    ~     (p=0.287 n=29+30)
    XML             220ms ± 2%       220ms ± 4%    ~     (p=0.805 n=29+29)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   281user-ms ± 4%  279user-ms ± 3%  -0.87%  (p=0.044 n=28+28)
    Unicode    142user-ms ± 4%  141user-ms ± 3%  -1.04%  (p=0.015 n=30+27)
    GoTypes    884user-ms ± 3%  886user-ms ± 2%    ~     (p=0.532 n=30+30)
    Compiler   3.94user-s ± 3%  3.92user-s ± 1%    ~     (p=0.185 n=30+28)
    Flate      165user-ms ± 2%  165user-ms ± 4%    ~     (p=0.780 n=27+29)
    GoParser   209user-ms ± 2%  208user-ms ± 3%    ~     (p=0.453 n=28+30)
    Reflect    533user-ms ± 6%  526user-ms ± 3%    ~     (p=0.057 n=30+30)
    Tar        156user-ms ± 6%  154user-ms ± 6%    ~     (p=0.133 n=29+30)
    XML        288user-ms ± 4%  288user-ms ± 4%    ~     (p=0.633 n=30+30)
    
    name       old alloc/op     new alloc/op     delta
    Template       41.0MB ± 0%      40.9MB ± 0%  -0.11%  (p=0.000 n=29+29)
    Unicode        32.6MB ± 0%      32.6MB ± 0%    ~     (p=0.572 n=29+30)
    GoTypes         122MB ± 0%       122MB ± 0%  -0.10%  (p=0.000 n=30+30)
    Compiler        482MB ± 0%       481MB ± 0%  -0.07%  (p=0.000 n=30+29)
    Flate          26.6MB ± 0%      26.6MB ± 0%    ~     (p=0.096 n=30+30)
    GoParser       32.7MB ± 0%      32.6MB ± 0%  -0.06%  (p=0.011 n=28+28)
    Reflect        84.2MB ± 0%      84.1MB ± 0%  -0.17%  (p=0.000 n=29+30)
    Tar            27.7MB ± 0%      27.7MB ± 0%  -0.05%  (p=0.032 n=27+28)
    XML            44.7MB ± 0%      44.7MB ± 0%    ~     (p=0.131 n=28+30)
    
    name       old allocs/op    new allocs/op    delta
    Template         373k ± 1%        370k ± 1%  -0.76%  (p=0.000 n=30+30)
    Unicode          325k ± 1%        325k ± 1%    ~     (p=0.383 n=29+30)
    GoTypes         1.16M ± 0%       1.15M ± 0%  -0.75%  (p=0.000 n=29+30)
    Compiler        4.15M ± 0%       4.13M ± 0%  -0.59%  (p=0.000 n=30+29)
    Flate            238k ± 1%        237k ± 1%  -0.62%  (p=0.000 n=30+30)
    GoParser         304k ± 1%        302k ± 1%  -0.64%  (p=0.000 n=30+28)
    Reflect         1.00M ± 0%       0.99M ± 0%  -1.10%  (p=0.000 n=29+30)
    Tar              245k ± 1%        244k ± 1%  -0.59%  (p=0.000 n=27+29)
    XML              391k ± 1%        389k ± 1%  -0.59%  (p=0.000 n=29+30)
    
    Change-Id: Id7f456d690567c2b0a96b0d6d64de8784b6e305f
    Reviewed-on: https://go-review.googlesource.com/36476
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit f6fc0dd620a838be3c98acc798ba58d6cbb0bac2
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Feb 19 23:40:24 2017 -0500

    cmd/compile: update signature of runtime.memclr*
    
    runtime.memclr* functions have signatures
    
    func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
    func memclrHasPointers(ptr unsafe.Pointer, n uintptr)
    
    Update compiler's copy. Also teach gc/mkbuiltin.go to handle
    unsafe.Pointer. The import statement and its support is not
    really necessary, but just to make it look like real Go code.
    
    Fixes #19185.
    
    Change-Id: I251d02571fde2716d4727e31e04d56ec04b6f22a
    Reviewed-on: https://go-review.googlesource.com/37257
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit d3d2a67c137939003eada835f28a62b26ab1f89f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Feb 28 11:10:34 2017 -0800

    cmd/vet/all: temporarily ignore vendored pprof
    
    Change-Id: I3d96b9803dbbd7184f96240bd7944af919ca1376
    Reviewed-on: https://go-review.googlesource.com/37579
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d99d5f7caa10b679f8509c22aafb35a51ab716ae
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Feb 28 10:29:58 2017 -0800

    cmd/vet: allow shifts by amounts calculated using unsafe
    
    The real world code that inspired this fix,
    from runtime/pprof/map.go:
    
            // Compute hash of (stk, tag).
            h := uintptr(0)
            for _, x := range stk {
                    h = h<<8 | (h >> (8 * (unsafe.Sizeof(h) - 1)))
                    h += uintptr(x) * 41
            }
            h = h<<8 | (h >> (8 * (unsafe.Sizeof(h) - 1)))
            h += uintptr(tag) * 41
    
    Change-Id: I99a95b97cba73811faedb0b9a1b9b54e9a1784a3
    Reviewed-on: https://go-review.googlesource.com/37574
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 016569f204ed1c1060778b03ecacb33bc882d69a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Feb 28 11:02:08 2017 -0800

    cmd/vet/all: move suspicious shift whitelists to 64 bit
    
    This is an inconsequential consequence of updating
    math/big to use math/bits.
    
    Better would be to teach the vet shift test
    to size int/uint/uintptr to the platform in use,
    eliminating the whole category of "might be too small".
    Filed #19321 for that.
    
    Change-Id: I7e0b837bd329132d7a564468c18502dd2e724fc6
    Reviewed-on: https://go-review.googlesource.com/37576
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 31f9769c910e1470125fc80ff134e872980a3951
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Feb 28 18:12:32 2017 +0000

    cmd/dist: make the vetall builder have test shards per os/arch
    
    This makes the vetall builder friendly to auto-sharding by the build
    coordinator.
    
    Change-Id: I0893f5051ec90e7a6adcb89904ba08cd2d590549
    Reviewed-on: https://go-review.googlesource.com/37572
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 8defd9f708126a1c6968490ac970c279c1fdb0c5
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Feb 28 10:08:10 2017 -0800

    cmd/vet/all: exit with non-zero error code on failure
    
    Change-Id: I68e60b155c583fa47aa5ca13d591851009a4e571
    Reviewed-on: https://go-review.googlesource.com/37571
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit bd8a39b67a12ec3d271305105dea3b8521aa70bf
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Feb 12 22:12:12 2017 -0500

    cmd/compile: emit fused multiply-{add,subtract} instructions on s390x
    
    Explcitly block fused multiply-add pattern matching when a cast is used
    after the multiplication, for example:
    
        - (a * b) + c        // can emit fused multiply-add
        - float64(a * b) + c // cannot emit fused multiply-add
    
    float{32,64} and complex{64,128} casts of matching types are now kept
    as OCONV operations rather than being replaced with OCONVNOP operations
    because they now imply a rounding operation (and therefore aren't a
    no-op anymore).
    
    Operations (for example, multiplication) on complex types may utilize
    fused multiply-add and -subtract instructions internally. There is no
    way to disable this behavior at the moment.
    
    Improves the performance of the floating point implementation of
    poly1305:
    
    name         old speed     new speed     delta
    64           246MB/s ± 0%  275MB/s ± 0%  +11.48%   (p=0.000 n=10+8)
    1K           312MB/s ± 0%  357MB/s ± 0%  +14.41%  (p=0.000 n=10+10)
    64Unaligned  246MB/s ± 0%  274MB/s ± 0%  +11.43%  (p=0.000 n=10+10)
    1KUnaligned  312MB/s ± 0%  357MB/s ± 0%  +14.39%   (p=0.000 n=10+8)
    
    Updates #17895.
    
    Change-Id: Ia771d275bb9150d1a598f8cc773444663de5ce16
    Reviewed-on: https://go-review.googlesource.com/36963
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit a38a2d02d5c26adb34ec18b495a1d698517c1735
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Feb 28 14:20:44 2017 +0100

    crypto/sha512: fix checkAVX2
    
    The checkAVX2 test doesn't appear to be correct,
    because it always returns the value of support_bmi2,
    even if the value of support_avx2 is false.
    
    Consequently, checkAVX2 always returns true, as long
    as BMI2 is supported, even if AVX2 is not supported.
    
    We change checkAVX2 to return false when support_avx2
    is false.
    
    Fixes #19316.
    
    Change-Id: I2ec9dfaa09f4b54c4a03d60efef891b955d60578
    Reviewed-on: https://go-review.googlesource.com/37590
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit a8f07310e3a08910dde2b7e9550848ec400753ad
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Mon Dec 19 00:06:46 2016 +0100

    cmd/compile: fix assignment order in string range loop
    
    Fixes #18376.
    
    Change-Id: I4fe24f479311cd4cd1bdad9a966b681e50e3d500
    Reviewed-on: https://go-review.googlesource.com/35955
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 55310403ddf051634fa398b4c9e6013d530753f5
Author: Carlo Alberto Ferraris <cafxx@strayorange.com>
Date:   Sat Feb 25 08:48:00 2017 +0900

    bytes: make bytes.Buffer cache-friendly
    
    During benchmark of an internal tool we found out that (*Buffer).Reset() was
    surprisingly showing up in CPU profiles.
    
    This CL contains two related changes aimed at speeding up Reset():
    1. Create a fast path for Truncate(0) by moving the logic to Reset()
       (this makes Reset() a simple leaf func that gets inlined since it
       gets compiled to 3 MOVx instructions). Accordingly change calls in
       the rest of the Buffer methods to call Reset() instead of Truncate(0).
    2. Reorder the fields in the Buffer struct so that frequently accessed
       fields are packed together (buf, off, lastRead). This also make them
       likely to be in the same cacheline.
    
    Ideally it would be advisable to have Buffer{} cacheline-aligned, but I
    couldn't find a way to do this without changing the size of the bootstrap
    array (but this will cause some regressions, because it will make duffcopy
    show up in CPU profiles where it wasn't showing up before).
    
    go1 benchmarks are not really affected, but some other benchmarks that
    exercise Buffer more show improvements:
    
    name                     old time/op    new time/op    delta
    BinaryTree17-4              2.46s ± 9%     2.43s ± 3%    ~     (p=0.982 n=14+14)
    Fannkuch11-4                2.98s ± 1%     2.90s ± 1%  -2.58%  (p=0.000 n=15+14)
    FmtFprintfEmpty-4          45.2ns ± 1%    45.2ns ± 1%    ~     (p=0.494 n=14+15)
    FmtFprintfString-4         76.8ns ± 1%    83.1ns ± 2%  +8.23%  (p=0.000 n=10+15)
    FmtFprintfInt-4            78.0ns ± 2%    74.6ns ± 1%  -4.46%  (p=0.000 n=15+15)
    FmtFprintfIntInt-4          113ns ± 1%     109ns ± 2%  -2.91%  (p=0.000 n=13+15)
    FmtFprintfPrefixedInt-4     152ns ± 2%     143ns ± 2%  -6.04%  (p=0.000 n=15+14)
    FmtFprintfFloat-4           224ns ± 1%     222ns ± 2%  -1.08%  (p=0.001 n=15+14)
    FmtManyArgs-4               464ns ± 2%     463ns ± 2%    ~     (p=0.303 n=14+15)
    GobDecode-4                6.25ms ± 2%    6.32ms ± 3%  +1.20%  (p=0.002 n=14+14)
    GobEncode-4                5.41ms ± 2%    5.41ms ± 2%    ~     (p=0.967 n=15+15)
    Gzip-4                      215ms ± 2%     218ms ± 2%  +1.35%  (p=0.002 n=15+15)
    Gunzip-4                   34.3ms ± 2%    34.2ms ± 2%    ~     (p=0.539 n=15+15)
    HTTPClientServer-4         76.4µs ± 2%    75.4µs ± 1%  -1.31%  (p=0.000 n=15+15)
    JSONEncode-4               14.7ms ± 2%    14.6ms ± 3%    ~     (p=0.094 n=14+14)
    JSONDecode-4               48.0ms ± 1%    48.5ms ± 1%  +0.92%  (p=0.001 n=14+12)
    Mandelbrot200-4            4.04ms ± 2%    4.06ms ± 1%    ~     (p=0.108 n=15+13)
    GoParse-4                  2.99ms ± 2%    3.00ms ± 1%    ~     (p=0.130 n=15+13)
    RegexpMatchEasy0_32-4      78.3ns ± 1%    79.5ns ± 1%  +1.51%  (p=0.000 n=15+14)
    RegexpMatchEasy0_1K-4       185ns ± 1%     186ns ± 1%  +0.76%  (p=0.005 n=15+15)
    RegexpMatchEasy1_32-4      79.0ns ± 2%    76.7ns ± 1%  -2.87%  (p=0.000 n=14+15)
    
    name                     old speed      new speed      delta
    GobDecode-4               123MB/s ± 2%   121MB/s ± 3%  -1.18%  (p=0.002 n=14+14)
    GobEncode-4               142MB/s ± 2%   142MB/s ± 1%    ~     (p=0.959 n=15+15)
    Gzip-4                   90.3MB/s ± 2%  89.1MB/s ± 2%  -1.34%  (p=0.002 n=15+15)
    Gunzip-4                  565MB/s ± 2%   567MB/s ± 2%    ~     (p=0.539 n=15+15)
    JSONEncode-4              132MB/s ± 2%   133MB/s ± 3%    ~     (p=0.091 n=14+14)
    JSONDecode-4             40.4MB/s ± 1%  40.0MB/s ± 1%  -0.92%  (p=0.001 n=14+12)
    GoParse-4                19.4MB/s ± 2%  19.3MB/s ± 1%    ~     (p=0.121 n=15+13)
    RegexpMatchEasy0_32-4     409MB/s ± 1%   403MB/s ± 1%  -1.47%  (p=0.000 n=15+14)
    RegexpMatchEasy0_1K-4    5.53GB/s ± 1%  5.49GB/s ± 1%  -0.86%  (p=0.002 n=15+15)
    RegexpMatchEasy1_32-4     405MB/s ± 2%   417MB/s ± 1%  +2.94%  (p=0.000 n=14+15)
    
    name                old time/op  new time/op  delta
    PoolsSingle1K-4     34.9ns ± 2%  30.4ns ± 4%  -12.80%  (p=0.000 n=15+15)
    PoolsSingle64K-4    36.9ns ± 1%  34.4ns ± 4%   -6.72%  (p=0.000 n=14+15)
    PoolsRandomSmall-4  34.8ns ± 3%  29.5ns ± 1%  -15.19%  (p=0.000 n=15+14)
    PoolsRandomLarge-4  38.6ns ± 1%  34.3ns ± 3%  -11.17%  (p=0.000 n=14+15)
    PoolSingle1K-4      26.1ns ± 1%  21.2ns ± 2%  -18.59%  (p=0.000 n=15+14)
    PoolSingle64K-4     26.7ns ± 2%  21.5ns ± 2%  -19.72%  (p=0.000 n=15+15)
    MakeSingle1K-4      24.2ns ± 2%  24.3ns ± 3%     ~     (p=0.132 n=13+15)
    MakeSingle64K-4     6.76µs ± 1%  6.96µs ± 5%   +2.94%  (p=0.002 n=13+13)
    MakeRandomSmall-4    531ns ± 4%   538ns ± 5%     ~     (p=0.066 n=14+15)
    MakeRandomLarge-4    152µs ± 0%   152µs ± 1%   -0.31%  (p=0.001 n=14+13)
    
    Change-Id: I86d7d9d2cac65335baf62214fbb35ba0fd8f9528
    Reviewed-on: https://go-review.googlesource.com/37416
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 417f49a363e7dbd01e91c79fc4f13c03681c2cc8
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Feb 18 13:58:59 2017 -0800

    cmd/compile: fold (NegNN (ConstNN ...))
    
    Fix up and enable a few rules.
    They trigger a handful of times in std,
    despite the frontend handling.
    
    Change-Id: I83378c057cbbc95a4f2b58cd8c36aec0e9dc547f
    Reviewed-on: https://go-review.googlesource.com/37227
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 4dbcb53d0b5842d4db5735f9a67935405aa84eab
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Feb 21 15:20:38 2017 -0500

    cmd/compile: fix merging of s390x conditional moves into branch conditions
    
    A type conversion inserted between MOVD{LT,LE,GT,GE,EQ,NE} and CMPWconst
    by CL 36256 broke the rewrite rule designed to merge the two.
    This results in simple for loops (e.g. for i := 0; i < N; i++ {})
    emitting two comparisons instead of one, plus a conditional move.
    
    This CL explicitly types the input to CMPWconst so that the type conversion
    can be omitted. It also adds a test to check that conditional moves aren't
    emitted for loops with 'less than' conditions (i.e. i < N) on s390x.
    
    Fixes #19227.
    
    Change-Id: Ia39e806ed723791c3c755951aef23f957828ea3e
    Reviewed-on: https://go-review.googlesource.com/37334
    Reviewed-by: Keith Randall <khr@golang.org>

commit 1b31c9ff679c98deccd06477ec48fc190bd5ca53
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Thu Feb 2 00:14:25 2017 -0800

    net/url: document the package better
    
    Changes made:
    * Adjust the documented form for a URL to make it more obvious what
    happens when the scheme is missing.
    * Remove references to Go1.5. We are sufficiently far along enough
    that this distinction no longer matters.
    * Remove the "Opaque" example which provides a hacky and misleading
    use of the Opaque field. This workaround is no longer necessary
    since RawPath was added in Go1.5 and the obvious approach just works:
            // The raw string "/%2f/" will be sent as expected.
            req, _ := http.NewRequest("GET", "https://example.com/%2f/")
    
    Fixes #18824
    
    Change-Id: Ie33d27222e06025ce8025f8a0f04b601aaee1513
    Reviewed-on: https://go-review.googlesource.com/36127
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f81466ce9c663ddc1d136d2c435c60d988cf316a
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 23 20:25:09 2017 -0800

    go/importer: support importing directly from source
    
    For #11415.
    
    Change-Id: I5da39dad059113cfc4276152390aa4925bd18862
    Reviewed-on: https://go-review.googlesource.com/37405
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit bab191042b92b125f5fac83dc8f0201160b79ecd
Author: Austin Clements <austin@google.com>
Date:   Mon Feb 20 17:49:56 2017 -0500

    cmd/internal/obj, runtime: update funcdata comments
    
    The comments in cmd/internal/obj/funcdata.go are identical to the
    comments in runtime/funcdata.h, but the majority of the definitions
    they refer to don't apply to Go sources and have been stripped out of
    funcdata.go.
    
    Remove these stale comments from funcdata.go and clean up the
    references to other copies of the PCDATA and FUNCDATA indexes.
    
    Change-Id: I5d6e49a6e586cc9aecd7c3ce1567679f2a605884
    Reviewed-on: https://go-review.googlesource.com/37330
    Reviewed-by: Keith Randall <khr@golang.org>

commit 949f95e7a40715ad05015dc4cb039e78a5260ef8
Author: Kevin Burke <kev@inburke.com>
Date:   Wed Nov 30 10:03:09 2016 -0800

    os/user: add Go implementation of LookupGroup, LookupGroupId
    
    If cgo is not available, parse /etc/group in Go to find the name/gid
    we need. This does not consult the Network Information System (NIS),
    /etc/nsswitch.conf or any other libc extensions to /etc/group.
    
    Fixes #18102.
    
    Change-Id: I6ae4fe0e2c899396c45cdf243d5483113932657c
    Reviewed-on: https://go-review.googlesource.com/33713
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 459d061c99b8bcd0ab688e2536f5429c9f125a4b
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Feb 22 15:24:30 2017 -0800

    go/internal/srcimporter: implemented srcimporter
    
    For #11415.
    
    Change-Id: I87a8f534ab9dfd5022422457ea637b342c057d77
    Reviewed-on: https://go-review.googlesource.com/37393
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 1e29cd8c2b560825494e6ae079a17d9f3201b73b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 27 11:37:54 2017 -0800

    cmd/compile: ignore some dead code during escape analysis
    
    This is the escape analysis analog of CL 37499.
    
    Fixes #12397
    Fixes #16871
    
    The only "moved to heap" decisions eliminated by this
    CL in std+cmd are:
    
    cmd/compile/internal/gc/const.go:1514: moved to heap: ac
    cmd/compile/internal/gc/const.go:1515: moved to heap: bd
    cmd/compile/internal/gc/const.go:1516: moved to heap: bc
    cmd/compile/internal/gc/const.go:1517: moved to heap: ad
    cmd/compile/internal/gc/const.go:1546: moved to heap: ac
    cmd/compile/internal/gc/const.go:1547: moved to heap: bd
    cmd/compile/internal/gc/const.go:1548: moved to heap: bc
    cmd/compile/internal/gc/const.go:1549: moved to heap: ad
    cmd/compile/internal/gc/const.go:1550: moved to heap: cc_plus
    cmd/compile/internal/gc/export.go:162: moved to heap: copy
    cmd/compile/internal/gc/mpfloat.go:66: moved to heap: b
    cmd/compile/internal/gc/mpfloat.go:97: moved to heap: b
    
    Change-Id: I0d420b69c84a41ba9968c394e8957910bab5edea
    Reviewed-on: https://go-review.googlesource.com/37508
    Reviewed-by: David Chase <drchase@google.com>

commit 7b8f51188bd24746b7d0a624b2e9979a425745eb
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Feb 24 16:02:31 2017 -0800

    cmd/compile/internal/gc: refactor liveness bitmap generation
    
    Keep liveness bit vectors as simple live-variable vectors during
    liveness analysis. We can defer expanding them into runtime heap
    bitmaps until we're actually writing out the symbol data, and then we
    only need temporary memory to expand one bitmap at a time.
    
    This is logically cleaner (e.g., we no longer depend on stack frame
    layout during analysis) and saves a little bit on allocations.
    
    name       old alloc/op    new alloc/op    delta
    Template      41.4MB ± 0%     41.3MB ± 0%  -0.28%        (p=0.000 n=60+60)
    Unicode       32.6MB ± 0%     32.6MB ± 0%  -0.11%        (p=0.000 n=59+60)
    GoTypes        119MB ± 0%      119MB ± 0%  -0.35%        (p=0.000 n=60+59)
    Compiler       483MB ± 0%      481MB ± 0%  -0.47%        (p=0.000 n=59+60)
    
    name       old allocs/op   new allocs/op   delta
    Template        381k ± 1%       380k ± 1%  -0.32%        (p=0.000 n=60+60)
    Unicode         325k ± 1%       325k ± 1%    ~           (p=0.867 n=60+60)
    GoTypes        1.16M ± 0%      1.15M ± 0%  -0.40%        (p=0.000 n=60+59)
    Compiler       4.22M ± 0%      4.19M ± 0%  -0.61%        (p=0.000 n=59+60)
    
    Passes toolstash -cmp.
    
    Change-Id: I8175efe55201ffb5017f79ae6cb90df03f1b7e99
    Reviewed-on: https://go-review.googlesource.com/37458
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 88f423edaca505d612f9d5ebb6ec35fa0efb61dc
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Feb 10 15:26:47 2017 -0800

    cmd/internal/obj/x86: improve static branch prediction for wrapper prologue
    
    Static branch prediction assumes that forward branches are not taken.
    The existing wrapper prologue almost always takes the first forward
    branch.
    Move the rare case to the end of the function.
    
    This CL is amd64 only. Other architectures will be done in separate CLs.
    
    Updates #19042.
    
    Package sort benchmarks:
    
    SearchWrappers-8       104ns ± 2%   104ns ± 0%  -0.41%  (p=0.006 n=30+41)
    SortString1K-8         128µs ± 1%   128µs ± 1%  -0.25%  (p=0.045 n=30+56)
    SortString1K_Slice-8   117µs ± 1%   117µs ± 1%    ~     (p=0.855 n=30+59)
    StableString1K-8      18.6µs ± 1%  18.6µs ± 1%    ~     (p=0.599 n=29+60)
    SortInt1K-8           61.0µs ± 1%  56.5µs ± 1%  -7.36%  (p=0.000 n=29+58)
    StableInt1K-8         74.6µs ± 1%  70.4µs ± 3%  -5.54%  (p=0.000 n=28+60)
    StableInt1K_Slice-8   59.9µs ± 1%  58.3µs ± 4%  -2.64%  (p=0.000 n=29+60)
    SortInt64K-8          6.02ms ± 2%  5.98ms ± 2%  -0.60%  (p=0.000 n=29+59)
    SortInt64K_Slice-8    5.07ms ± 2%  5.05ms ± 2%  -0.38%  (p=0.006 n=30+58)
    StableInt64K-8        6.41ms ± 1%  6.22ms ± 1%  -3.00%  (p=0.000 n=27+58)
    Sort1e2-8             37.4µs ± 1%  37.1µs ± 1%  -0.91%  (p=0.000 n=30+57)
    Stable1e2-8           74.8µs ± 1%  75.2µs ± 1%  +0.52%  (p=0.000 n=30+57)
    Sort1e4-8             8.11ms ± 1%  8.01ms ± 1%  -1.20%  (p=0.000 n=30+59)
    Stable1e4-8           24.3ms ± 1%  24.3ms ± 1%    ~     (p=0.157 n=30+60)
    Sort1e6-8              1.25s ± 1%   1.23s ± 1%  -1.43%  (p=0.000 n=29+58)
    Stable1e6-8            4.93s ± 1%   4.90s ± 1%  -0.56%  (p=0.000 n=29+59)
    [Geo mean]             720µs        709µs       -1.52%
    
    Assembly for sort.(*intPairs).Swap:
    
    Before:
    
    "".(*intPairs).Swap t=1 size=147 args=0x18 locals=0x8
            0x0000 00000 (<autogenerated>:1)        TEXT    "".(*intPairs).Swap(SB), $8-24
            0x0000 00000 (<autogenerated>:1)        MOVQ    (TLS), CX
            0x0009 00009 (<autogenerated>:1)        SUBQ    $8, SP
            0x000d 00013 (<autogenerated>:1)        MOVQ    BP, (SP)
            0x0011 00017 (<autogenerated>:1)        LEAQ    (SP), BP
            0x0015 00021 (<autogenerated>:1)        MOVQ    32(CX), BX
            0x0019 00025 (<autogenerated>:1)        TESTQ   BX, BX
            0x001c 00028 (<autogenerated>:1)        JEQ     43
            0x001e 00030 (<autogenerated>:1)        LEAQ    16(SP), DI
            0x0023 00035 (<autogenerated>:1)        CMPQ    (BX), DI
            0x0026 00038 (<autogenerated>:1)        JNE     43
            0x0028 00040 (<autogenerated>:1)        MOVQ    SP, (BX)
            0x002b 00043 (<autogenerated>:1)        NOP
            0x002b 00043 (<autogenerated>:1)        FUNCDATA        $0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
            0x002b 00043 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
            0x002b 00043 (<autogenerated>:1)        MOVQ    ""..this+16(FP), AX
            0x0030 00048 (<autogenerated>:1)        TESTQ   AX, AX
            0x0033 00051 (<autogenerated>:1)        JEQ     $0, 140
            0x0035 00053 (<autogenerated>:1)        MOVQ    (AX), CX
            0x0038 00056 (<autogenerated>:1)        MOVQ    8(AX), AX
            0x003c 00060 (<autogenerated>:1)        MOVQ    "".i+24(FP), DX
            0x0041 00065 (<autogenerated>:1)        CMPQ    DX, AX
            0x0044 00068 (<autogenerated>:1)        JCC     $0, 133
            0x0046 00070 (<autogenerated>:1)        SHLQ    $4, DX
            0x004a 00074 (<autogenerated>:1)        MOVQ    8(CX)(DX*1), BX
            0x004f 00079 (<autogenerated>:1)        MOVQ    (CX)(DX*1), SI
            0x0053 00083 (<autogenerated>:1)        MOVQ    "".j+32(FP), DI
            0x0058 00088 (<autogenerated>:1)        CMPQ    DI, AX
            0x005b 00091 (<autogenerated>:1)        JCC     $0, 133
            0x005d 00093 (<autogenerated>:1)        SHLQ    $4, DI
            0x0061 00097 (<autogenerated>:1)        MOVQ    8(CX)(DI*1), AX
            0x0066 00102 (<autogenerated>:1)        MOVQ    (CX)(DI*1), R8
            0x006a 00106 (<autogenerated>:1)        MOVQ    R8, (CX)(DX*1)
            0x006e 00110 (<autogenerated>:1)        MOVQ    AX, 8(CX)(DX*1)
            0x0073 00115 (<autogenerated>:1)        MOVQ    SI, (CX)(DI*1)
            0x0077 00119 (<autogenerated>:1)        MOVQ    BX, 8(CX)(DI*1)
            0x007c 00124 (<autogenerated>:1)        MOVQ    (SP), BP
            0x0080 00128 (<autogenerated>:1)        ADDQ    $8, SP
            0x0084 00132 (<autogenerated>:1)        RET
            0x0085 00133 (<autogenerated>:1)        PCDATA  $0, $1
            0x0085 00133 (<autogenerated>:1)        CALL    runtime.panicindex(SB)
            0x008a 00138 (<autogenerated>:1)        UNDEF
            0x008c 00140 (<autogenerated>:1)        PCDATA  $0, $1
            0x008c 00140 (<autogenerated>:1)        CALL    runtime.panicwrap(SB)
            0x0091 00145 (<autogenerated>:1)        UNDEF
    
    After:
    
    "".(*intPairs).Swap t=1 size=149 args=0x18 locals=0x8
            0x0000 00000 (<autogenerated>:1)        TEXT    "".(*intPairs).Swap(SB), $8-24
            0x0000 00000 (<autogenerated>:1)        MOVQ    (TLS), CX
            0x0009 00009 (<autogenerated>:1)        SUBQ    $8, SP
            0x000d 00013 (<autogenerated>:1)        MOVQ    BP, (SP)
            0x0011 00017 (<autogenerated>:1)        LEAQ    (SP), BP
            0x0015 00021 (<autogenerated>:1)        MOVQ    32(CX), BX
            0x0019 00025 (<autogenerated>:1)        TESTQ   BX, BX
            0x001c 00028 (<autogenerated>:1)        JNE     134
            0x001e 00030 (<autogenerated>:1)        NOP
            0x001e 00030 (<autogenerated>:1)        FUNCDATA        $0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
            0x001e 00030 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
            0x001e 00030 (<autogenerated>:1)        MOVQ    ""..this+16(FP), AX
            0x0023 00035 (<autogenerated>:1)        TESTQ   AX, AX
            0x0026 00038 (<autogenerated>:1)        JEQ     $0, 127
            0x0028 00040 (<autogenerated>:1)        MOVQ    (AX), CX
            0x002b 00043 (<autogenerated>:1)        MOVQ    8(AX), AX
            0x002f 00047 (<autogenerated>:1)        MOVQ    "".i+24(FP), DX
            0x0034 00052 (<autogenerated>:1)        CMPQ    DX, AX
            0x0037 00055 (<autogenerated>:1)        JCC     $0, 120
            0x0039 00057 (<autogenerated>:1)        SHLQ    $4, DX
            0x003d 00061 (<autogenerated>:1)        MOVQ    8(CX)(DX*1), BX
            0x0042 00066 (<autogenerated>:1)        MOVQ    (CX)(DX*1), SI
            0x0046 00070 (<autogenerated>:1)        MOVQ    "".j+32(FP), DI
            0x004b 00075 (<autogenerated>:1)        CMPQ    DI, AX
            0x004e 00078 (<autogenerated>:1)        JCC     $0, 120
            0x0050 00080 (<autogenerated>:1)        SHLQ    $4, DI
            0x0054 00084 (<autogenerated>:1)        MOVQ    8(CX)(DI*1), AX
            0x0059 00089 (<autogenerated>:1)        MOVQ    (CX)(DI*1), R8
            0x005d 00093 (<autogenerated>:1)        MOVQ    R8, (CX)(DX*1)
            0x0061 00097 (<autogenerated>:1)        MOVQ    AX, 8(CX)(DX*1)
            0x0066 00102 (<autogenerated>:1)        MOVQ    SI, (CX)(DI*1)
            0x006a 00106 (<autogenerated>:1)        MOVQ    BX, 8(CX)(DI*1)
            0x006f 00111 (<autogenerated>:1)        MOVQ    (SP), BP
            0x0073 00115 (<autogenerated>:1)        ADDQ    $8, SP
            0x0077 00119 (<autogenerated>:1)        RET
            0x0078 00120 (<autogenerated>:1)        PCDATA  $0, $1
            0x0078 00120 (<autogenerated>:1)        CALL    runtime.panicindex(SB)
            0x007d 00125 (<autogenerated>:1)        UNDEF
            0x007f 00127 (<autogenerated>:1)        PCDATA  $0, $1
            0x007f 00127 (<autogenerated>:1)        CALL    runtime.panicwrap(SB)
            0x0084 00132 (<autogenerated>:1)        UNDEF
            0x0086 00134 (<autogenerated>:1)        LEAQ    16(SP), DI
            0x008b 00139 (<autogenerated>:1)        CMPQ    (BX), DI
            0x008e 00142 (<autogenerated>:1)        JNE     30
            0x0090 00144 (<autogenerated>:1)        MOVQ    SP, (BX)
            0x0093 00147 (<autogenerated>:1)        JMP     30
    
    Change-Id: Ie8c37f384bba10fbacaa754bb0a6b0a7e520ef01
    Reviewed-on: https://go-review.googlesource.com/36893
    Reviewed-by: Keith Randall <khr@golang.org>

commit f7f3514bd874844f2091a123e55b19fdaf4773b5
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Feb 22 14:43:21 2017 -0800

    cmd/compile/internal/gc: simplify ascompatte
    
    Passes toolstash -cmp.
    
    Change-Id: Ibb51ccaf29ee97c3463543175c9ac7b85ea10a7f
    Reviewed-on: https://go-review.googlesource.com/37339
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit ba6e5776fd6796282533ba49e80a5afd8b1ced23
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Mon Feb 27 20:39:49 2017 +0100

    runtime: remove unused RaceSemacquire declaration
    
    These functions are not defined and are not used.
    
    Fixes #19290
    
    Change-Id: I2978147220af83cf319f7439f076c131870fb9ee
    Reviewed-on: https://go-review.googlesource.com/37448
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 31e633464481a0748fb95eda2a2d3d604755d286
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 27 11:14:45 2017 -0800

    go/build: move math/bits into L1 set of dependencies
    
    Per suggestion from rsc.
    
    Change-Id: I4b61ec6f35ffaaa792b75e011fbba1bdfbabc1f6
    Reviewed-on: https://go-review.googlesource.com/37501
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5ae7cbfff6f98d2a87dca7362e34e48d7bba0d69
Author: Tom Bergan <tombergan@google.com>
Date:   Mon Feb 27 11:09:42 2017 -0800

    net/http: update bundles http2
    
    Updates http2 to x/net/http2 git rev 906cda9 for:
    
    http2: add configurable knobs for the server's receive window
    https://golang.org/cl/37226
    
    http2/hpack: speedup Encoder.searchTable
    https://golang.org/cl/37406
    
    http2: Add opt-in option to Framer to allow DataFrame struct reuse
    https://golang.org/cl/34812
    
    http2: replace fixedBuffer with dataBuffer
    https://golang.org/cl/37400
    
    http2/hpack: remove hpack's constant time string comparison
    https://golang.org/cl/37394
    
    Updates golang/go#16512
    Updates golang/go#18404
    
    Change-Id: I1ad7c95c404ead4ced7f85af061cf811b299a288
    Reviewed-on: https://go-review.googlesource.com/37500
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0df81e88876e0b2b40b13e49d6be12c26334070b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 27 10:45:26 2017 -0800

    cmd/compile: simplify and clean up inlnode
    
    Change-Id: I0d14d68b57e8605cdae8a45d6fa97255a42297d8
    Reviewed-on: https://go-review.googlesource.com/37521
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 566e72d0cedc593054dd36f9d3e91b588e849074
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Feb 26 15:25:57 2017 -0800

    cmd/compile: ignore some dead code when deciding whether to inline
    
    Constant evaluation provides some rudimentary
    knowledge of dead code at inlining decision time.
    Use it.
    
    This CL addresses only dead code inside if statements.
    For statements are never inlined anyway,
    and dead code inside for statements is rare.
    Analyzing switch statements is worth doing,
    but it is more complicated, since we would have
    to evaluate each case; leave it for later.
    
    Fixes #9274
    
    After this CL, the following functions in std+cmd
    can be newly inlined:
    
    cmd/internal/obj/x86/asm6.go:3122: can inline subreg
    cmd/vendor/golang.org/x/arch/x86/x86asm/decode.go:172: can inline instPrefix
    cmd/vendor/golang.org/x/arch/x86/x86asm/decode.go:202: can inline truncated
    go/constant/value.go:234: can inline makeFloat
    go/types/labels.go:52: can inline (*block).insert
    math/big/float.go:231: can inline (*Float).Sign
    math/bits/bits.go:57: can inline OnesCount
    net/http/server.go:597: can inline (*Server).newConn
    runtime/hashmap.go:1165: can inline reflect_maplen
    runtime/proc.go:207: can inline os_beforeExit
    runtime/signal_unix.go:55: can inline init.5
    runtime/stack.go:1081: can inline gostartcallfn
    
    Change-Id: I4c92fb96aa0c3d33df7b3f2da548612e79b56b5b
    Reviewed-on: https://go-review.googlesource.com/37499
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c7894924c7702c88e2b9836323bd4f40cd0257e7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Feb 11 14:44:16 2017 -0800

    runtime/pprof: handle empty stack traces in Profile.Add
    
    If the caller passes a large number to Profile.Add,
    the list of pcs is empty, which results in junk
    (a nil pc) being recorded. Check for that explicitly,
    and replace such stack traces with a lostProfileEvent.
    
    Fixes #18836.
    
    Change-Id: I99c96aa67dd5525cd239ea96452e6e8fcb25ce02
    Reviewed-on: https://go-review.googlesource.com/36891
    Reviewed-by: Russ Cox <rsc@golang.org>

commit eae657e9ee2897ffac8c8918738b0a4bab5864d6
Author: Kevin Burke <kev@inburke.com>
Date:   Mon Feb 27 02:28:24 2017 -0800

    os/user: rename group cgo file
    
    In another CL, I'll add a pure Go implementation of lookupGroup and
    lookupGroupId in lookup_unix.go, but attempting that in one CL makes
    the diff too difficult to read.
    
    Updates #18102.
    
    Change-Id: If8e26cee5efd30385763430f34304c70165aef32
    Reviewed-on: https://go-review.googlesource.com/37497
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4b2248527ffec97ab0706470e5654bd1aacaa6ab
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Feb 26 22:27:13 2017 +0000

    os: skip atime-going-backwards test on NetBSD for now
    
    That failing test is preventing other tests from running.
    Let's see what else is broken.
    
    Updates #19293
    
    Change-Id: I4c5784be94103ef882f29dec9db08d76a48aff28
    Reviewed-on: https://go-review.googlesource.com/37492
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matt Layher <mdlayher@gmail.com>

commit e18adbf88d35bc200d2b1c07ccb8f55f551942a0
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 24 17:50:53 2017 -0800

    math/bits: faster Reverse8/16 functions using table lookups
    
    Measured on 2.3 GHz Intel Core i7, running macOS 10.12.3:
    
    benchmark                old ns/op     new ns/op     delta
    BenchmarkReverse8-8      1.70          0.99          -41.76%
    BenchmarkReverse16-8     2.24          1.32          -41.07%
    
    Fixes #19279.
    
    Change-Id: I398cf8a3513b7fa63c130efc7846a7c5353999d4
    Reviewed-on: https://go-review.googlesource.com/37459
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit bf584b15d6104661c272e2ab900c3d83d015b7d0
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 24 21:52:37 2017 -0800

    cmd/dist: ran mkdeps.bash
    
    Change-Id: Iae9fe2db69c02cd442cba01a78820dc7c0fdda51
    Reviewed-on: https://go-review.googlesource.com/37462
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e458264aca63432b225a83267690baced3eb8f26
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Feb 24 21:40:57 2017 -0800

    cmd/compile: fix dolinkobj flag in TestAssembly
    
    Follow-up to CL 37270.
    
    This considerably reduces the time to run the test.
    
    Before:
    
    real    0m7.638s
    user    0m14.341s
    sys     0m2.244s
    
    After:
    
    real    0m4.867s
    user    0m7.107s
    sys     0m1.842s
    
    Change-Id: I8837a5da0979a1c365e1ce5874d81708249a4129
    Reviewed-on: https://go-review.googlesource.com/37461
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit a6b480bc07e400cee109eda90caddb20c2cc7e82
Author: Bill O'Farrell <billo@ca.ibm.com>
Date:   Fri Feb 24 17:03:06 2017 -0500

    cmd/go: implement -buildmode=plugin for s390x
    
    Simple change to allow plugins for linux on s390x
    
    Change-Id: I5c262ab81aac10d1dcb03381a48e5b9694b7a87a
    Reviewed-on: https://go-review.googlesource.com/37451
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit febafe60d469ce129d4662d7f3550218fa548616
Author: David Chase <drchase@google.com>
Date:   Fri Feb 24 17:21:54 2017 -0500

    cmd/compile: added cheapexpr call to simplify operand of CONVIFACE
    
    New special case for booleans and byte-sized integer types
    converted to interfaces needs to ensure that the operand is
    not too complex, if it were to appear in a parameter list
    for example.
    
    Added test, also increased the recursive node dump depth to
    a level that was actually useful for an actual bug.
    
    Fixes #19275.
    
    Change-Id: If36ac3115edf439e886703f32d149ee0a46eb2a5
    Reviewed-on: https://go-review.googlesource.com/37470
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit ac91a514ff521999b142901ad9714ca3f47f01a0
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 24 16:50:48 2017 -0800

    math/bits: fix incorrect doc strings for TrailingZeros functions
    
    Change-Id: I3e40018ab1903d3b9ada7ad7812ba71ea2a428e7
    Reviewed-on: https://go-review.googlesource.com/37456
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4b3e6fe123d95f461d8f9febfe782a138ba2387c
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Mon Nov 14 08:05:45 2016 +0100

    strings: speed up Map
    
    name                  old time/op  new time/op  delta
    ByteByteMap-4         2.03µs ± 2%  1.03µs ± 2%  -49.24%  (p=0.000 n=10+10)
    Map/identity/ASCII-4   246ns ± 0%   158ns ± 0%  -35.90%    (p=0.000 n=9+9)
    Map/identity/Greek-4   367ns ± 1%   273ns ± 1%  -25.63%  (p=0.000 n=10+10)
    Map/change/ASCII-4     582ns ± 1%   324ns ± 1%  -44.34%  (p=0.000 n=10+10)
    Map/change/Greek-4     709ns ± 2%   623ns ± 2%  -12.16%  (p=0.000 n=10+10)
    MapNoChanges-4         171ns ± 1%   111ns ± 1%  -35.36%   (p=0.000 n=8+10)
    
    Updates #17859
    
    Change-Id: I55d7d261fdc1ce2dcd0ebe23b0fa20b9889bf54c
    Reviewed-on: https://go-review.googlesource.com/33201
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 221bc23af6feeab821c49ad11f4661270d310ecd
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Feb 23 16:09:05 2017 -0500

    os/exec: deflake TestPipeLookPathLeak
    
    The number of open file descriptors reported by lsof is unreliable
    because it depends on whether the parent process (the test) closed
    the file descriptors it passed into the child process (lsof) before
    lsof runs.
    
    Reading /proc/self/fd directly on Linux appears to be much more
    reliable and still detects any file descriptor leaks originating
    from attempting to run an executable that cannot be found (issue
    #5071). If /proc/self/fd is not available (e.g. on Darwin) then we
    fall back to lsof and tolerate small differences in open file
    descriptor counts.
    
    Fixes #19243.
    
    Change-Id: I052b0c129e609010f1083e43a9911cba154117bf
    Reviewed-on: https://go-review.googlesource.com/37343
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit fdef951116ea5e201866b7d4a53c8c90056770f4
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sun Feb 19 15:57:58 2017 +0100

    cmd/compile: make setting and accessing of node slice elements more uniform
    
    Add Set3 function to complement existing Set1 and Set2 functions.
    Consistently use Set1, Set2 and Set3 for []*Node instead of Set where applicable.
    
    Add SetFirst and SetSecond for setting elements of []*Node to mirror
    First and Second for accessing elements in []*Node.
    
    Replace uses of Index by First and Second and
    SetIndex with SetFirst and SetSecond where applicable.
    
    Passes toolstash -cmp.
    
    Change-Id: I8255aae768cf245c8f93eec2e9efa05b8112b4e5
    Reviewed-on: https://go-review.googlesource.com/37430
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit fb1f47a77c642f36fc7a80d468ad1cda8bb66361
Author: Lorenzo Masini <rugginoso@develer.com>
Date:   Mon Feb 20 17:17:28 2017 +0100

    cmd/compile: speed up TestAssembly
    
    TestAssembly was very slow, leading to it being skipped by default.
    This is not surprising, it separately invoked the compiler and
    parsed the result many times.
    
    Now the test assembles one source file for arch/os combination,
    containing the relevant functions.
    
    Tests for each arch/os run in parallel.
    
    Now the test runs approximately 10x faster on my Intel(R) Core(TM)
    i5-6600 CPU @ 3.30GHz.
    
    Fixes #18966
    
    Change-Id: I45ab97630b627a32e17900c109f790eb4c0e90d9
    Reviewed-on: https://go-review.googlesource.com/37270
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 349b7820eb8bd30d42afad945b22e4e9fe74eff1
Author: Russ Cox <rsc@golang.org>
Date:   Fri Feb 24 16:01:50 2017 -0500

    test: deflake locklinear a little
    
    This should help on the openbsd systems where the test mostly passes.
    
    I don't expect it to help on s390x where the test reliably fails.
    But it should give more information when it does fail.
    
    For #19276.
    
    Change-Id: I496c291f2b4b0c747b8dd4315477d87d03010059
    Reviewed-on: https://go-review.googlesource.com/37348
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 67fcd9c5d9f84e4f66df1e357ca9f76523ecff4e
Author: Kevin Burke <kev@inburke.com>
Date:   Fri Feb 24 12:07:25 2017 -0800

    cmd/internal/browser: fix typo
    
    Change-Id: I3c31f10c1082c7bc57aac18856014c55f79e0fed
    Reviewed-on: https://go-review.googlesource.com/37409
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8c24e52247365725c97b1ed5bea63a84642fd0f7
Author: Russ Cox <rsc@golang.org>
Date:   Fri Feb 17 15:27:12 2017 -0500

    runtime: check that pprof accepts but doesn't need executable
    
    The profiles are self-contained now.
    Check that they work by themselves in the tests that invoke pprof,
    but also keep checking that the old command lines work.
    
    Change-Id: I24c74b5456f0b50473883c3640625c6612f72309
    Reviewed-on: https://go-review.googlesource.com/37166
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit 0b8c983ece56b63c433a65fd3de6a411cb2aac87
Author: Russ Cox <rsc@golang.org>
Date:   Fri Feb 17 00:17:26 2017 -0500

    runtime/pprof/internal/profile: move internal/pprof/profile here
    
    Nothing needs internal/pprof anymore except the runtime/pprof tests.
    Move the package here to prevent new dependencies.
    
    Change-Id: Ia119af91cc2b980e0fa03a15f46f69d7f71d2926
    Reviewed-on: https://go-review.googlesource.com/37165
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit cbab65fdfa1cf74f65a480de0447388286169f26
Author: Russ Cox <rsc@golang.org>
Date:   Fri Feb 17 00:10:39 2017 -0500

    runtime/pprof: add streaming protobuf encoder
    
    The existing code builds a full profile in memory.
    Then it translates that profile into a data structure (in memory).
    Then it marshals that data structure into a protocol buffer (in memory).
    Then it gzips that marshaled form into the underlying writer.
    So there are three copies of the full profile data in memory
    at the same time before we're done. This is obviously dumb.
    
    This CL implements a fully streaming conversion from
    the original in-memory profile to the underlying writer.
    There is now only one copy of the profile in memory.
    
    For the non-CPU profiles, this is optimal, since we have to
    have a full copy in memory to start with.
    
    For the CPU profiles, we could still try to bound the profile
    size stored in memory and stream fragments out during
    the actual profiling, as Go 1.7 did (with a simpler format),
    but so far that hasn't been necessary.
    
    Change-Id: Ic36141021857791bf0cd1fce84178fb5e744b989
    Reviewed-on: https://go-review.googlesource.com/37164
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit 322fff8ac855390dc2a2876e9051a8dd526d2c6a
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 16 15:04:33 2017 -0800

    math/big: use math/bits where appropriate
    
    This change adds math/bits as a new dependency of math/big.
    
    - use bits.LeadingZeroes instead of local implementation
      (they are identical, so there's no performance loss here)
    
    - leave other functionality local (ntz, bitLen) since there's
      faster implementations in math/big at the moment
    
    Change-Id: I1218aa8a1df0cc9783583b090a4bb5a8a145c4a2
    Reviewed-on: https://go-review.googlesource.com/37141
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7844ef427a61bd68db81912a254a3f99633a9354
Author: Raul Silvera <rsilvera@google.com>
Date:   Fri Feb 10 14:52:02 2017 -0800

    cmd/pprof: vendor pprof from github.com/google/pprof
    
    Import the github.com/google/pprof and github.com/ianlancetaylor/demangle
    packages, without modification.
    
    Build the golang version of pprof from cmd/pprof/pprof.go
    by importing the packages from src/cmd/vendot/github.com/google/pprof
    
    The versions upstreamed are:
    
    github.com/ianlancetaylor/demangle 4883227f66371e02c4948937d3e2be1664d9be38
    github.com/google/pprof            7eb5ba977f28f2ad8dd5f6bb82cc9b454e123cdc
    
    Update misc/nacl/testzip.proto for new tests.
    
    Change-Id: I076584856491353607a3b98b67d0ca6838be50d6
    Reviewed-on: https://go-review.googlesource.com/36798
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 2818cb5c9e183aed539d6a539a821e229671fe56
Author: Chris Broadfoot <cbro@golang.org>
Date:   Wed Feb 22 15:50:24 2017 -0800

    cmd/internal/browser: wait 3 seconds for non-zero exit codes
    
    Wait a short period between trying commands. Many commands
    will return a non-zero exit code if the browser couldn't be launched.
    
    For example, google-chrome returns quickly with a non-zero
    exit code in a headless environment.
    
    Updates #19131.
    
    Change-Id: I0ae5356dd4447969d9e216615449cead7a8fd5c9
    Reviewed-on: https://go-review.googlesource.com/37391
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit d9270ecb3ad35079df62ad85b3a5e52e46e4a1c0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Feb 22 21:19:29 2017 -0800

    cmd/compile: evaluate zero-sized values converted to interfaces
    
    CL 35562 substituted zerobase for the pointer for
    interfaces containing zero-sized values.
    However, it failed to evaluate the zero-sized value
    expression for side-effects. Fix that.
    
    The other similar interface value optimizations
    are not affected, because they all actually use the
    value one way or another.
    
    Fixes #19246
    
    Change-Id: I1168a99561477c63c29751d5cd04cf81b5ea509d
    Reviewed-on: https://go-review.googlesource.com/37395
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit f8ae30c4a201dbdb6652cbb72cd51762863c7447
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Feb 22 13:43:23 2017 -0800

    cmd/compile/internal/parser: improved a couple of error messages
    
    The new syntax tree introduced with 1.8 represents send statements
    (ch <- x) as statements; the old syntax tree represented them as
    expressions (and parsed them as such) but complained if they were
    used in expression context. As a consequence, some of the errors
    that in the past were of the form "ch <- x used as value" now look
    like "unexpected <- ..." because a "<-" is not valid according to
    Go syntax in those situations. Accept the new error message.
    
    Also: Fine-tune handling of misformed for loop headers.
    
    Also: Minor cleanups/better comments.
    
    Fixes #17590.
    
    Change-Id: Ia541dea1f2f015c1b21f5b3ae44aacdec60a8aba
    Reviewed-on: https://go-review.googlesource.com/37386
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2fa09a20e56eb27f7cec635be42fc3137c091085
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Feb 24 10:35:43 2017 -0800

    vendor/golang_org/x/crypto/chacha20poly1305: revendor
    
    Brings in chacha20poly1305 directory from golang.org/x/crypto revision
    453249f01cfeb54c3d549ddb75ff152ca243f9d8, adding:
    
    CL 35874: crypto/chacha20poly1305/internal/chacha20: add missing copyright header
    CL 35875: crypto/chacha20poly1305: rename test vectors file
    
    Fixes #19155.
    
    Change-Id: I25cf83d060113f6b2a197f243a25614440008f7e
    Reviewed-on: https://go-review.googlesource.com/37408
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2a0d7e24aad70e01a5a301b7e5248fc9adfb0e26
Author: Sean Christopherson <sean.j.christopherson@intel.com>
Date:   Wed Feb 22 06:56:26 2017 -0800

    run.bash: set GOPATH to $GOROOT/nil before running tests
    
    Set $GOPATH to a semantically valid, non-empty string that cannot
    conflict with $GOROOT to avoid false test failures that occur when
    $GOROOT resides under $GOPATH.  Unsetting GOPATH is no longer viable
    as Go now defines a default $GOPATH that may conflict with $GOROOT.
    
    Fixes #19237
    
    Change-Id: I376a2ad3b18e9c4098211b988dde7e76bc4725d2
    Reviewed-on: https://go-review.googlesource.com/37396
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1564817d8c941344caa14d14fac55f7e7b46055e
Author: Russ Cox <rsc@golang.org>
Date:   Thu Feb 16 21:13:15 2017 -0500

    runtime/pprof: use more efficient hash table for staging profile
    
    The old hash table was a place holder that allocates memory
    during every lookup for key generation, even for keys that hit
    in the the table.
    
    Change-Id: I4f601bbfd349f0be76d6259a8989c9c17ccfac21
    Reviewed-on: https://go-review.googlesource.com/37163
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit 1a680a902ada111a0d8122e12b42f5b437ba2566
Author: Russ Cox <rsc@golang.org>
Date:   Thu Feb 9 15:57:57 2017 -0500

    runtime/pprof: use new profile buffers for CPU profiling
    
    This doesn't change the functionality of the current code,
    but it sets us up for exporting the profiling labels into the profile.
    
    The old code had a hash table of profile samples maintained
    during the signal handler, with evictions going into a log.
    The new code just logs every sample directly, leaving the
    hash-based deduplication to an ordinary goroutine.
    
    The new code also avoids storing the entire profile in two
    forms in memory, an unfortunate regression introduced
    when binary profile support was added. After this CL the
    entire profile is only stored once in memory. We'd still like
    to get back down to storing it zero times (streaming it to
    the underlying io.Writer).
    
    Change-Id: I0893a1788267c564aa1af17970d47377b2a43457
    Reviewed-on: https://go-review.googlesource.com/36712
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit a1261b8b0a38814df453defb2fc2cae3ba0c956a
Author: Russ Cox <rsc@golang.org>
Date:   Fri Feb 17 10:17:42 2017 -0500

    runtime: do not allocate on every time.Sleep
    
    It's common for some goroutines to loop calling time.Sleep.
    Allocate once per goroutine, not every time.
    This comes up in runtime/pprof's background reader.
    
    Change-Id: I89d17dc7379dca266d2c9cd3aefc2382f5bdbade
    Reviewed-on: https://go-review.googlesource.com/37162
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit a1ea91219faa7c35098ffbb958582897fcd33123
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Wed Jan 11 16:53:49 2017 -0800

    cmd/doc: truncate long lists of arguments
    
    Some field-lists (especially in generated code) can be excessively long.
    In the one-line printout, it does not make sense to print all elements
    of the list if line-wrapping causes the "one-line" to become multi-line.
    
    // Before:
    var LongLine = newLongLine("someArgument1", "someArgument2", "someArgument3", "someArgument4", "someArgument5", "someArgument6", "someArgument7", "someArgument8")
    
    // After:
    var LongLine = newLongLine("someArgument1", "someArgument2", "someArgument3", "someArgument4", ...)
    
    Change-Id: I4bbbe2dbd1d7be9f02d63431d213088c3dee332c
    Reviewed-on: https://go-review.googlesource.com/36031
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit ea5529de155cfd3f2c31698344b1ca001e0f8819
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Thu Feb 23 17:41:31 2017 -0800

    crypto/tls: use io.ReadFull in conn_test.go
    
    An io.Reader does not guarantee that it will read in the entire buffer.
    To ensure that property, io.ReadFull should be used instead.
    
    Change-Id: I0b863135ab9abc40e813f9dac07bfb2a76199950
    Reviewed-on: https://go-review.googlesource.com/37403
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit abdb2c35b6a0684b3eabe9892ff9e3518d2f4e78
Author: David Chase <drchase@google.com>
Date:   Thu Feb 23 13:49:25 2017 -0500

    cmd/compile: repaired loop-finder to handle trickier nesting
    
    The loop-A-encloses-loop-C code did not properly handle the
    case where really C was already known to be enclosed by B,
    and A was nearest-outer to B, not C.
    
    Fixes #19217.
    
    Change-Id: I755dd768e823cb707abdc5302fed39c11cdb34d4
    Reviewed-on: https://go-review.googlesource.com/37340
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 8ca68c3fec18bec7739ceac0f55681f915baa7f9
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 23 12:52:33 2017 -0800

    go/types: fix doc string for Named.Obj
    
    Fixes #19249.
    
    Change-Id: I6327192eca11fa24f1078c016c9669e4ba4bdb4e
    Reviewed-on: https://go-review.googlesource.com/37399
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d55f52882656122869b6b409be783bfb5a2fd2fb
Author: David R. Jenni <david.r.jenni@gmail.com>
Date:   Fri Feb 10 19:25:58 2017 +0100

    cmd/compile: silence superfluous assignment error message
    
    Avoid printing a second error message when a field of an undefined
    variable is accessed.
    
    Fixes #8440.
    
    Change-Id: I3fe0b11fa3423cec3871cb01b5951efa8ea7451a
    Reviewed-on: https://go-review.googlesource.com/36751
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b788fd80e68addd181983ca589d28bc1075f4a36
Author: Russ Cox <rsc@golang.org>
Date:   Thu Feb 9 13:58:48 2017 -0500

    runtime: new profile buffer implementation supporting label pointers
    
    The existing CPU profiling buffer is a slice of uintptr, but we want to
    start including profiling label data in the profiles, and those labels need
    to be pointers in order to let them describe rich information.
    
    This CL implements a new profBuf type that holds both a slice of uint64
    for data and a slice of unsafe.Pointer for profiling labels (aka tags).
    Making the runtime use these buffers will happen in followup CLs.
    
    Change-Id: I9ff16b532d8edaf4ce0cbba1098229a561834efc
    Reviewed-on: https://go-review.googlesource.com/36713
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 9230ee20f3893814cbfb5e0bdc1a117b092eae35
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Thu Feb 23 19:36:34 2017 +0100

    website: mention go1.8 in project page
    
    Fixes #19253
    
    Change-Id: Ia473f51bfe4cf42cf64938993a81d9b1dbc2594d
    Reviewed-on: https://go-review.googlesource.com/37433
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit d580972d5933c77d9319ec99ff1b2f111f05c727
Author: Chris Broadfoot <cbro@golang.org>
Date:   Wed Feb 22 15:48:38 2017 -0800

    cmd/internal/browser: use xdg-open only from a desktop session
    
    xdg-open's man page says:
    > xdg-open is for use inside a desktop session only.
    
    Use the DISPLAY environment variable to detect this.
    
    Updates #19131.
    
    Change-Id: I3926b3e1042393939b2ec6aacd9b63ac8192df3b
    Reviewed-on: https://go-review.googlesource.com/37390
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 731fd009f0acef70d939f3cb62f81a83e3e9e2bb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Feb 22 12:47:52 2017 -0800

    cmd/vet/all: use -dolinkobj=false to speed up runs
    
    When running on the host platform,
    the standard library has almost certainly already been built.
    However, all other platforms will probably need building.
    Use the new -dolinkobj=false flag to cmd/compile
    to only build the export data instead of doing a full compile.
    
    Having partial object files could be confusing for people
    doing subsequent cross-compiles, depending on what happens with #18369.
    However, cmd/vet/all will mainly be run by builders
    and core developers, who are probably fairly well-placed
    to handle any such confusion.
    
    This reduces the time on my machine for a cold run of
    'go run main.go -all' by almost half:
    
    benchmark           old ns/op        new ns/op        delta
    BenchmarkVetAll     240670814551     130784517074     -45.66%
    
    Change-Id: Ieb866ffb2cb714b361b0a6104077652f8eacd166
    Reviewed-on: https://go-review.googlesource.com/37385
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 005c77dde89d6a062c021a3ed0e180a6848d82b4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Feb 22 00:05:18 2017 -0800

    cmd/compile: add -dolinkobj flag
    
    When set to false, the -dolinkobj flag instructs the compiler
    not to generate or emit linker information.
    
    This is handy when you need the compiler's export data,
    e.g. for use with go/importer,
    but you want to avoid the cost of full compilation.
    
    This must be used with care, since the resulting
    files are unusable for linking.
    
    This CL interacts with #18369,
    where adding gcflags and ldflags to buildid has been mooted.
    On the one hand, adding gcflags would make safe use of this
    flag easier, since if the full object files were needed,
    a simple 'go install' would fix it.
    On the other hand, this would mean that
    'go install -gcflags=-dolinkobj=false' would rebuild the object files,
    although any existing object files would probably suffice.
    
    Change-Id: I8dc75ab5a40095c785c1a4d2260aeb63c4d10f73
    Reviewed-on: https://go-review.googlesource.com/37384
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2b2870fff8ebf712c05ed46ae63cc6174ed72ed1
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Feb 22 22:23:35 2017 +0000

    doc: fix broken link in go1.8.html
    
    Fixes #19244
    
    Change-Id: Ia6332941b229c83d6fd082af49f31003a66b90db
    Reviewed-on: https://go-review.googlesource.com/37388
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 72a071c1da98458e9f7ccf1812b401903acf5b1d
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Feb 20 15:54:43 2017 -0500

    cmd/compile: rewrite pairs of shifts to extensions
    
    Replaces pairs of shifts with sign/zero extension where possible.
    
    For example:
    (uint64(x) << 32) >> 32 -> uint64(uint32(x))
    
    Reduces the execution time of the following code by ~4.5% on s390x:
    
    for i := 0; i < N; i++ {
            x += (uint64(i)<<32)>>32
    }
    
    Change-Id: Idb2d56f27e80a2e1366bc995922ad3fd958c51a7
    Reviewed-on: https://go-review.googlesource.com/37292
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 8321be63396363cd18e9d23b4b05bcb3e5791fa7
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Sat Feb 18 16:56:45 2017 +0000

    sort: new example: Sorting slices with sort.SliceStable
    
    ExampleSliceStable echoes the sort.Slice example, to demonstrate sorting
    on two fields together preserving order between sorts.
    
    Change-Id: I8afc20c0203991bfd57260431eda73913c165355
    Reviewed-on: https://go-review.googlesource.com/37196
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c12cd31a3326e6b2119525cd07cebc6d6e1b52ee
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Thu Feb 16 14:58:47 2017 -0200

    cmd/internal/obj/ppc64: Fix RLDIMI
    
    Fix the encoding of the SH field for rldimi.
    
    The SH field of rldimi is 6-bit wide and it is not contiguous in the instruction.
    Bits 0-4 are placed in bit fields 16-20 in the instruction, while bit 5 is
    placed in bit field 30. The current implementation does not consider this and,
    therefore, any SH field between 32 and 63 are encoded wrongly in the instruciton.
    
    Change-Id: I4d25a0a70f4219569be0e18160dea5505bd7fff0
    Reviewed-on: https://go-review.googlesource.com/37350
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit c804fd8927dabe1484b52e31f752136b7c0b32b7
Author: Yuval Pavel Zholkover <paulzhol@gmail.com>
Date:   Wed Feb 22 21:24:44 2017 +0200

    net: update IP.MarshalText documentation regarding len(ip) == 0
    
    Describe the difference from String encoding when len(ip) is zero.
    
    Change-Id: Ia9b36b405d4fec3fee9a77498a839b6d90c2ec0d
    Reviewed-on: https://go-review.googlesource.com/37379
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8c6643846ef5572cb138c8f7c9ac2b1b3cb1d06c
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sat Dec 10 08:04:40 2016 +0100

    math: speed up and improve accuracy of Pow10
    
    Removes init function from the math package.
    
    Allows stripping of arrays with pre-computed values
    used for Pow10 from binaries if Pow10 is not used.
    cmd/go shrinks by 128 bytes.
    
    Fixed small values like 10**-323 being 0 instead of 1e-323.
    
    Overall precision is increased but still not as good as
    predefined constants for some inputs.
    
    Samples:
    
    Pow10(208)
    before: 1.0000000000000006662e+208
    after:  1.0000000000000000959e+208
    
    Pow10(202)
    before 1.0000000000000009895e+202
    after  1.0000000000000001193e+202
    
    Pow10(60)
    before 1.0000000000000001278e+60
    after  0.9999999999999999494e+60
    
    Pow10(-100)
    before 0.99999999999999938551e-100
    after  0.99999999999999989309e-100
    
    Pow10(-200)
    before 0.9999999999999988218e-200
    after  1.0000000000000001271e-200
    
    name        old time/op  new time/op  delta
    Pow10Pos-4  44.6ns ± 2%   1.2ns ± 1%  -97.39%  (p=0.000 n=19+17)
    Pow10Neg-4  50.8ns ± 1%   4.1ns ± 2%  -92.02%  (p=0.000 n=17+19)
    
    Change-Id: If094034286b8ac64be3a95fd9e8ffa3d4ad39b31
    Reviewed-on: https://go-review.googlesource.com/36331
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6694a01016e2fb842ee6d8ae053b9c87f56185ff
Author: Michael Pratt <mpratt@google.com>
Date:   Tue Feb 21 17:00:10 2017 -0800

    cmd/dist: fix negative test filtering
    
    std and race bench tests fail to check against t.runRxWant, so what
    should be negative filters act as positive filters.
    
    Fixes #19239
    
    Change-Id: Icf02b2192bcd806a162fca9fb0af68a27ccfc936
    Reviewed-on: https://go-review.googlesource.com/37336
    Run-TryBot: Michael Pratt <mpratt@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 19d2061d502d26086d6db75fa818dde668a888bf
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Fri Feb 10 23:41:53 2017 -0700

    cmd/compile: suppress callsite signatures if any type is unknown
    
    Fixes #19012.
    
    Fallback to return signatures without detailed types.
    These error message will be of the form of issue:
    * https://golang.org/issues/4215
    * https://golang.org/issues/6750
    
    So:
    func f(x int, y uint) {
        return x > y
    }
    
    f(10, "a" < 3)
    
    will give errors:
    too many errors to return
    too many arguments in call to f
    
    instead of:
    
    too many errors to return
      have (<T>)
      want ()
    too many arguments in call to f
      have (number, <T>)
      want (number, number)
    
    Change-Id: I680abc7cdd8444400e234caddf3ff49c2d69f53d
    Reviewed-on: https://go-review.googlesource.com/36806
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit ffb3b3698c695b2ada052b803f11a2ffa12e5537
Author: Alexander Döring <email@alexd.ch>
Date:   Mon Feb 20 20:34:36 2017 +0100

    math: add more tests for special cases of Bessel functions Y0, Y1, Yn
    
    Test finite negative x with Y0(-1), Y1(-1), Yn(2,-1), Yn(-3,-1).
    
    Also test the special case Yn(0,0).
    
    Fixes #19130.
    
    Change-Id: I95f05a72e1c455ed8ddf202c56f4266f03f370fd
    Reviewed-on: https://go-review.googlesource.com/37310
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit dc6af19ff8b44e56abc1217af27fe098c78c932b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Feb 21 12:48:45 2017 -0800

    context: document that Err is unspecified before Done
    
    It could have been defined the other way, but since the behavior has
    been unspecified, this is the conservative approach for people writing
    different implementations of the Context interface.
    
    Change-Id: I7334a4c674bc2330cca6874f7cac1eb0eaea3cff
    Reviewed-on: https://go-review.googlesource.com/37375
    Reviewed-by: Matt Layher <mdlayher@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Sameer Ajmani <sameer@golang.org>
    Run-TryBot: Sameer Ajmani <sameer@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ea48c9d2325dfe3ccd64a2bfeea9516cb5a1d2e3
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Tue Feb 21 11:28:49 2017 -0500

    runtime: more detail for crash_test.go
    
    This updates the testcase to display the timestamps for the
    runtime.a, it dependent packages atomic.a and sys.a, and
    source files.
    
    Change-Id: Id2901b4e8aa8eb9775c4f404ac01cc07b394ba91
    Reviewed-on: https://go-review.googlesource.com/37332
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 094992e22aaa1d7091659b63b5cadfe947372277
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Feb 16 19:53:08 2017 -0500

    cmd/compile: zero extend when replacing load-hit-store on s390x
    
    Keith pointed out that these rules should zero extend during the review
    of CL 36845. In practice the generic rules are responsible for eliminating
    most load-hit-stores and they do not have this problem. When the s390x
    rules are triggered any cast following the elided load-hit-store is
    kept because of the sequence the rules are applied in (i.e. the load is
    removed before the zero extension gets a chance to be merged into the load).
    It is therefore not clear that this issue results in any functional bugs.
    
    This CL includes a test, but it only tests the generic rules currently.
    
    Change-Id: Idbc43c782097a3fb159be293ec3138c5b36858ad
    Reviewed-on: https://go-review.googlesource.com/37154
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 11b283092a29a9d402ce05706fd3a85683576218
Author: David Chase <drchase@google.com>
Date:   Tue Feb 21 15:22:52 2017 -0500

    cmd/compile: add opcode flag hasSideEffects for do-not-remove
    
    Added a flag to generic and various architectures' atomic
    operations that are judged to have observable side effects
    and thus cannot be dead-code-eliminated.
    
    Test requires GOMAXPROCS > 1 without preemption in loop.
    
    Fixes #19182.
    
    Change-Id: Id2230031abd2cca0bbb32fd68fc8a58fb912070f
    Reviewed-on: https://go-review.googlesource.com/37333
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit b5e51943063a1d78045f0c3ce6c87b424795e643
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Feb 16 14:20:24 2017 -0800

    reflect: fix bucketOf to only look at ptrdata entries in gcdata
    
    The gcdata field only records ptrdata entries, not size entries.
    
    Also fix an obsolete comment: the enforced limit on pointer maps is
    now 2048 bytes, not 16 bytes.
    
    I wasn't able to contruct a test case for this. It would require
    building a type whose size is greater than 64 bytes but less than 128
    bytes, with at least one pointer in first 64 bytes but no pointers
    after the first 64 bytes, such that the linker arranges for the one
    byte gcbits value to be immediately followed by a non-zero byte.
    
    Change-Id: I9118d3e4ec6f07fd18b72f621c1e5f4fdfe5f80b
    Reviewed-on: https://go-review.googlesource.com/37142
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit db6e27c38d20cdd6af205bbf99c1b1d3327e6c6f
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Feb 21 15:57:06 2017 -0800

    cmd/compile: update builtin writeBarrier to match runtime
    
    The definition of writeBarrier in the runtime was changed in CL 22855
    to include padding. Update the definition built in to the compiler to match.
    This doesn't affect the generated code, as the compiler sets the type
    to use anyhow, but having them be different seems clearly wrong.
    
    Change-Id: I8eac05bf70a424a0b2338ba5e9e41af231316de0
    Reviewed-on: https://go-review.googlesource.com/37377
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 7603aa7907fa92336fd9f7702b709e559ee0191b
Author: Kevin Burke <kev@inburke.com>
Date:   Tue Feb 21 11:23:04 2017 -0800

    doc: use appropriate type to describe return value
    
    Fixes #19223.
    
    Change-Id: I4cc8e81559a1313e1477ee36902e1b653155a888
    Reviewed-on: https://go-review.googlesource.com/37374
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 6464e5dc4b84a4348e6698f97c5bfc462a0e3a5e
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sat Feb 18 21:03:15 2017 -0500

    cmd/compile: do not fold offset into load/store for args on ARM64
    
    Args may be not at 8-byte aligned offset to SP. When the stack
    frame is large, folding the offset of args may cause large
    unaligned offsets that does not fit in a machine instruction on
    ARM64. Therefore disable folding offsets for args.
    
    This has small performance impact (see below). A better fix would
    be letting the assembler backend fix up the offset by loading it
    into a register if it doesn't fit into an instruction. And the
    compiler can simply generate large load/stores with offset. Since
    in most of the cases the offset is aligned or the stack frame is
    small, it can fit in an instruction and no fixup is needed. But
    this is too complicated for Go 1.8.
    
    name                     old time/op    new time/op    delta
    BinaryTree17-8              8.30s ± 0%     8.31s ± 0%    ~     (p=0.579 n=10+10)
    Fannkuch11-8                6.14s ± 0%     6.18s ± 0%  +0.53%  (p=0.000 n=9+10)
    FmtFprintfEmpty-8           117ns ± 0%     117ns ± 0%    ~     (all equal)
    FmtFprintfString-8          196ns ± 0%     197ns ± 0%  +0.72%  (p=0.000 n=10+10)
    FmtFprintfInt-8             204ns ± 0%     205ns ± 0%  +0.49%  (p=0.000 n=9+10)
    FmtFprintfIntInt-8          302ns ± 0%     307ns ± 1%  +1.46%  (p=0.000 n=10+10)
    FmtFprintfPrefixedInt-8     329ns ± 2%     326ns ± 0%    ~     (p=0.083 n=10+10)
    FmtFprintfFloat-8           540ns ± 0%     542ns ± 0%  +0.46%  (p=0.000 n=8+7)
    FmtManyArgs-8              1.20µs ± 1%    1.19µs ± 1%  -1.02%  (p=0.000 n=10+10)
    GobDecode-8                17.3ms ± 1%    17.8ms ± 0%  +2.75%  (p=0.000 n=10+7)
    GobEncode-8                15.3ms ± 1%    15.4ms ± 0%  +0.57%  (p=0.004 n=9+10)
    Gzip-8                      789ms ± 0%     803ms ± 0%  +1.78%  (p=0.000 n=9+10)
    Gunzip-8                    128ms ± 0%     130ms ± 0%  +1.73%  (p=0.000 n=10+9)
    HTTPClientServer-8          202µs ± 6%     201µs ±10%    ~     (p=0.739 n=10+10)
    JSONEncode-8               42.0ms ± 0%    42.1ms ± 0%  +0.19%  (p=0.028 n=10+9)
    JSONDecode-8                159ms ± 0%     161ms ± 0%  +1.05%  (p=0.000 n=9+10)
    Mandelbrot200-8            10.1ms ± 0%    10.1ms ± 0%  -0.07%  (p=0.000 n=10+9)
    GoParse-8                  8.46ms ± 1%    8.61ms ± 1%  +1.77%  (p=0.000 n=10+10)
    RegexpMatchEasy0_32-8       227ns ± 1%     226ns ± 0%  -0.35%  (p=0.001 n=10+9)
    RegexpMatchEasy0_1K-8      1.63µs ± 0%    1.63µs ± 0%  -0.13%  (p=0.000 n=10+9)
    RegexpMatchEasy1_32-8       250ns ± 0%     249ns ± 0%  -0.40%  (p=0.001 n=8+9)
    RegexpMatchEasy1_1K-8      2.07µs ± 0%    2.08µs ± 0%  +0.05%  (p=0.027 n=9+9)
    RegexpMatchMedium_32-8      350ns ± 0%     350ns ± 0%    ~     (p=0.412 n=9+8)
    RegexpMatchMedium_1K-8      104µs ± 0%     104µs ± 0%  +0.31%  (p=0.000 n=10+7)
    RegexpMatchHard_32-8       5.82µs ± 0%    5.82µs ± 0%    ~     (p=0.937 n=9+9)
    RegexpMatchHard_1K-8        176µs ± 0%     176µs ± 0%  +0.03%  (p=0.000 n=9+8)
    Revcomp-8                   1.36s ± 1%     1.37s ± 1%    ~     (p=0.218 n=10+10)
    Template-8                  151ms ± 1%     156ms ± 1%  +3.21%  (p=0.000 n=10+10)
    TimeParse-8                 737ns ± 0%     758ns ± 2%  +2.74%  (p=0.000 n=10+10)
    TimeFormat-8                801ns ± 2%     789ns ± 1%  -1.51%  (p=0.000 n=10+10)
    [Geo mean]                  142µs          143µs       +0.50%
    
    Fixes #19137.
    
    Change-Id: Ib8a21ea98c0ffb2d282a586535b213cc163e1b67
    Reviewed-on: https://go-review.googlesource.com/37251
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 174058038c72aa0e2cc254ef91d4dbf2956a8d1e
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 21 10:22:05 2017 -0800

    math/big: define Word as uint instead of uintptr
    
    For compatibility with math/bits uint operations.
    
    When math/big was written originally, the Go compiler used 32bit
    int/uint values even on a 64bit machine. uintptr was the type that
    represented the machine register size. Now, the int/uint types are
    sized to the native machine register size, so they are the natural
    machine Word type.
    
    On most machines, the size of int/uint correspond to the size of
    uintptr. On platforms where uint and uintptr have different sizes,
    this change may lead to performance differences (e.g., amd64p32).
    
    Change-Id: Ief249c160b707b6441848f20041e32e9e9d8d8ca
    Reviewed-on: https://go-review.googlesource.com/37372
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 10d718b9839dc7b1c55761d6c9f3001fac498cd0
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Feb 19 12:10:18 2017 -0500

    cmd/compile: fix type of OffPtr generated by ODOTPTR
    
    The type of the OffPtr should be consistent with the type of the
    following load. Before this CL it was typed as a pointer to the
    struct.
    
    Fixes #19164.
    
    Change-Id: Ibcdec4411c6f719702f76f8dba3cce8691bfbe0c
    Reviewed-on: https://go-review.googlesource.com/37254
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit ea020ff3de9482726ce7019ac43c1d301ce5e3de
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Tue Feb 21 14:30:17 2017 +0100

    fmt: add short note about %g precision
    
    Fixes #18772
    
    Change-Id: Ib5d9ffa0abd35b9d3ca83bac139aece0f3c9702d
    Reviewed-on: https://go-review.googlesource.com/37313
    Reviewed-by: Rob Pike <r@golang.org>

commit 689fa9cc2842141ad663845af2b6188cc1c12b5d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 20 22:58:36 2017 -0800

    syscall: fix linux/mipsx ret value FP offsets for Syscall9
    
    Found by vet.
    
    Change-Id: Idf910405566816ddce6781c8e99f90b59f33d63c
    Reviewed-on: https://go-review.googlesource.com/37308
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 00e2524d8aa7d4ee045530696a76efc9038962c0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 20 22:58:11 2017 -0800

    sync/atomic: fix mipsx frame sizes
    
    Found by vet.
    
    Change-Id: Ied3089a2cc8757ae5377fb5fa05bbb385d26ad9c
    Reviewed-on: https://go-review.googlesource.com/37307
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4208fcdcd40f0359e117d850dc180bda7fea3f92
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 20 22:57:45 2017 -0800

    runtime: use standard linux/mipsx clone variable names
    
    Change-Id: I62118e197190af1d11a89921d5769101ce6d2257
    Reviewed-on: https://go-review.googlesource.com/37306
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e51737aea14375ac72fb8c50241969d8972bb8fb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 20 22:57:10 2017 -0800

    crypto/aes: minor ppc64 assembly naming improvements
    
    doEncryptKeyAsm is tail-called from other assembly routines.
    Give it a proper prototype so that vet can check it.
    Adjust one assembly FP reference accordingly.
    
    Change-Id: I263fcb0191529214b16e6bd67330fadee492eef4
    Reviewed-on: https://go-review.googlesource.com/37305
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 2774085bdc9f0e64a3dd00633d5d6ae860988c78
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 20 22:55:51 2017 -0800

    cmd/vet/all: update windows whitelist
    
    A last-minute rollback of a change left some
    unreachable code that we don't want to remove.
    
    Change-Id: Ida0af5b18ed1a2e13ef66c303694afcc49d7bff4
    Reviewed-on: https://go-review.googlesource.com/37304
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b6e0d4647f9f9534c5dc215570ee80b6a4717928
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 20 22:54:39 2017 -0800

    runtime: update assembly var names after monotonic time changes
    
    Change-Id: I721045120a4df41462c02252e2e5e8529ae2d694
    Reviewed-on: https://go-review.googlesource.com/37303
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ea52f4b374613134e12367288813019589f85f33
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 20 22:54:07 2017 -0800

    cmd/vet/all: update whitelists for monotonic time changes
    
    Change-Id: Ib942cb9e0cb20821aea4274bc3ddc83a215afbcb
    Reviewed-on: https://go-review.googlesource.com/37302
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 84a855e547b8ef9a1a7aaaaedc3b1058ef7d1c09
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 13 16:19:51 2017 -0800

    cmd/vet/all: add mips and mipsle
    
    Change-Id: I689b2e8e214561350f88fa4e20c8f34cf69dc6a7
    Reviewed-on: https://go-review.googlesource.com/37301
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8ccd007f24673adca90feb5fd23f2131b42b99d0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 13 16:11:13 2017 -0800

    cmd/vet/all: work around vet printf checker deficiencies
    
    cmd/vet has a known deficiency in its handling of fmt.Formatters.
    This causes a spurious printf error only for non-host platforms.
    Since cmd/vet/all may get run on any given platform,
    whitelists cannot help here.
    
    Work around the issue by skipping printf tests entirely
    for non-host platforms.
    
    Work around the one known acceptable false positive from vet
    by whitelisting the file that contains it.
    
    Change-Id: Id74b3d4db0519cf9a670a065683715f856266e45
    Reviewed-on: https://go-review.googlesource.com/36936
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 35ffca31b180e6f9da6035326132f048980dc58c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Feb 21 07:14:51 2017 -0800

    os/exec: deflake TestStdinCloseRace
    
    Stop reporting errors from cmd.Process.Kill; they don't matter for
    purposes of this test, and they can occur if the process exits quickly.
    
    Fixes #19211.
    Fixes #19213.
    
    Change-Id: I1a0bb9170220ca69199abb8e8811b1dde43e1897
    Reviewed-on: https://go-review.googlesource.com/37309
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Daniel Martí <mvdan@mvdan.cc>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a355639c6046edae6b3a5190bc78260c2fe3b063
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sat Feb 18 18:37:38 2017 -0500

    cmd/compile: fix storeOrder
    
    storeOrder visits values in DFS order. It should "break" after
    pushing one argument to stack, instead of "continue".
    
    Fixes #19179.
    
    Change-Id: I561afb44213df40ebf8bf7d28e0fd00f22a81ac0
    Reviewed-on: https://go-review.googlesource.com/37250
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: David Chase <drchase@google.com>

commit b32170abdfe6979af249fccba513609fc52cfb15
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Feb 8 12:04:11 2017 +1100

    cmd/link: simplify peemitreloc
    
    No functional changes.
    
    For #10776.
    
    Change-Id: If9a5ef832af116c5802b06a38e0c050d7363f2d5
    Reviewed-on: https://go-review.googlesource.com/36981
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d0a978f5b503028bb9c34944edb59e52c4070f6f
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Feb 8 11:44:09 2017 +1100

    cmd/link: reorder pe sections
    
    dwarf writing code assumes that dwarf sections follow
    .data and .bss, not .ctors. Make pe section writing code
    match that assumption.
    
    For #10776.
    
    Change-Id: I128c3ad125f7d0db19e922f165704a054b2af7ba
    Reviewed-on: https://go-review.googlesource.com/36980
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6db4d92e4c01c854e15391d18100c8d99cbbd7bc
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Feb 8 11:42:24 2017 +1100

    cmd/link: do not add __image_base__ and _image_base__ if external linker
    
    The symbols get in a way when using external linker. They are
    not associated with a section. And linker fails when
    generating relocations for them.
    
    __image_base__ and _image_base__ have been added long time ago.
    I do not think they are needed anymore. If I delete them, all
    tests still PASS. I tried going back to the commit that added
    them to see if I can reproduce original error, but I cannot
    build it. I don't have hg version of go repo, and my gcc is
    complaining about cc source code.
    
    I wasted too much time with this, so I decided to leave them only
    for internal linker. That is what they were originally added for.
    
    For #10776.
    
    Change-Id: Ibb72b04f3864947c782f964a7badc69f4b074e25
    Reviewed-on: https://go-review.googlesource.com/36979
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b660a4b04d0a88e86d15c1235a4d3bdf1efcd12c
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Feb 8 13:59:19 2017 +1100

    cmd/link: add all pe section names to pe symbol table
    
    dwarf relocations refer to dwarf section symbols, so dwarf
    section symbols must be present in pe symbol table before we
    write dwarf relocations.
    
    .ctors pe section already refer to .text symbol.
    
    Write all pe section name symbols into symbol table, so we
    can use them whenever we need them.
    
    This CL also simplified some code.
    
    For #10776.
    
    Change-Id: I9b8c680ea75904af90c797a06bbb1f4df19e34b6
    Reviewed-on: https://go-review.googlesource.com/36978
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e9abf1a7166186cdcf84ccfd028c2c66f1ec4788
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Feb 8 13:58:21 2017 +1100

    cmd/link: introduce shNames
    
    Introduce a slice that keeps long pe section names as we add them.
    It will be used later to output pe symbol table and dwarf relocations.
    
    For #10776.
    
    Change-Id: I02f808a456393659db2354031baf1d4f9e0b2d61
    Reviewed-on: https://go-review.googlesource.com/36977
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a7e25562555a508571b63aedadb088eab8fc5a48
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Mon Feb 6 15:18:26 2017 +1100

    cmd/link: set VirtualAddress to 0 for external linker
    
    This is what gcc does when it generates object files.
    And pecoff.doc says: "for simplicity, compilers should
     set this to zero". It is easier to count everything,
    when it starts from 0. Make go linker do the same.
    
    For #10776.
    
    Change-Id: Iffa4b3ad86160624ed34adf1c6ba13feba34c658
    Reviewed-on: https://go-review.googlesource.com/36976
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a37f9d8a17cffc6fb79120fad667b0684fd03bc7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Feb 20 19:48:28 2017 +0000

    runtime/pprof: mark TestMutexProfile as flaky for now
    
    Flaky tests hurt productivity. Disable for now.
    
    Updates #19139
    
    Change-Id: I2e3040bdf0e53597a1c4f925b788e3268ea284c1
    Reviewed-on: https://go-review.googlesource.com/37291
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Peter Weinberger <pjw@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b9574f46f9079a4237a9dd42d30a6ec91723ee5a
Author: David Lazar <lazard@golang.org>
Date:   Sun Feb 19 14:01:42 2017 -0500

    cmd/objdump: make test independent of inlining
    
    Fixes #19189.
    
    Change-Id: Ice69216c7fc2eaeb3dbbdcd08a8284204c7f52ef
    Reviewed-on: https://go-review.googlesource.com/37237
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3892d507964dce100c55c3fa14f414bd14c34e98
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sun Feb 19 11:34:29 2017 +0100

    cmd/compile: remove unused constant divide strength reduction code
    
    Change list https://golang.org/cl/37015/ moved the optimization
    of division by constants to the generic ssa backend.
    This removes the old now unused code that was used
    for this optimization outside of the ssa backend.
    
    Change-Id: I86223e56742e48dbb372ba8d779681e66448c513
    Reviewed-on: https://go-review.googlesource.com/37198
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Daniel Martí <mvdan@mvdan.cc>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 177dfba1120d2d5976bb5fb5a68bf20bb6ca9ada
Author: Robert Griesemer <gri@golang.org>
Date:   Sat Feb 18 11:14:35 2017 -0800

    math/bits: faster OnesCount
    
    Using some additional suggestions per "Hacker's Delight".
    Added documentation and extra tests.
    
    Measured on 1.7 GHz Intel Core i7, running macOS 10.12.3.
    
    benchmark                  old ns/op     new ns/op     delta
    BenchmarkOnesCount-4       7.34          5.38          -26.70%
    BenchmarkOnesCount8-4      2.03          1.98          -2.46%
    BenchmarkOnesCount16-4     2.56          2.50          -2.34%
    BenchmarkOnesCount32-4     2.98          2.39          -19.80%
    BenchmarkOnesCount64-4     4.22          2.96          -29.86%
    
    Change-Id: I566b0ef766e55cf5776b1662b6016024ebe5d878
    Reviewed-on: https://go-review.googlesource.com/37223
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d9a19f86fb5297aee62242ad14b6a69d2c990a79
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sun Feb 19 10:42:34 2017 +0100

    fmt: remove unused global variable byteType
    
    Change list https://golang.org/cl/20686/ removed the last use
    of the variable byteType.
    
    Change-Id: I4ea79095136a49a9d22767b37f48f3404da05056
    Reviewed-on: https://go-review.googlesource.com/37197
    Run-TryBot: Martin Möhrmann <moehrmann@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit cfb0d349921a8f1ce3af8eb56188fc020d8828e5
Author: Keith Randall <khr@golang.org>
Date:   Sun Feb 19 00:16:27 2017 -0800

    cmd/compile: amd64, allow XCHG on stack pointers
    
    XCHG needs to allow the stack pointer as an argument because we have a
    rewrite that incorporates the address of a local variable into the
    instruction.
    
    Fixes #19184
    
    Change-Id: Ic438e6e1946332cdce3864d15abecd41b911b2a9
    Reviewed-on: https://go-review.googlesource.com/37253
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit f37428d8b7e06edc4be472695db677f82c9564c5
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Sat Feb 18 15:07:13 2017 -0800

    cmd/go/internal/envcmd: report PKG_CONFIG after the CGO group
    
    Before the change, `go env` reports PKG_CONFIG in between the
    CGO env group:
    
        GOARCH="amd64"
        GOBIN=""
        GOEXE=""
        GOHOSTARCH="amd64"
        GOHOSTOS="darwin"
        GOOS="darwin"
        GOPATH="/Users/jbd"
        GORACE=""
        GOROOT="/Users/jbd/go"
        GOTOOLDIR="/Users/jbd/go/pkg/tool/darwin_amd64"
        GCCGO="gccgo"
        CC="clang"
        GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/lq/qcn67khn4_1b41_g48x3zchh005d21/T/go-build184491598=/tmp/go-build -gno-record-gcc-switches -fno-common"
        CXX="clang++"
        CGO_ENABLED="1"
        PKG_CONFIG="pkg-config"
        CGO_CFLAGS="-g -O2"
        CGO_CPPFLAGS=""
        CGO_CXXFLAGS="-g -O2"
        CGO_FFLAGS="-g -O2"
        CGO_LDFLAGS="-g -O2"
    
    The change makes PKG_CONFIG to be reported as the final item,
    and not breaking the CGO_* group apart.
    
    Change-Id: I1e7ed6bdec83009ff118f85c9f0f7b78a67fdd76
    Reviewed-on: https://go-review.googlesource.com/37228
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e97f407ec164174187d6fbb2f7332d6fb3176e9e
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Wed Feb 15 12:41:02 2017 +0100

    fmt: support sharp flag for float and complex value printing
    
    Added an alternate form of printing floats and complex values
    by specifying the sharp flag.
    
    Output formatted using the the verbs v, e, E, f, F, g and G in
    combination with the sharp flag will always include a decimal point.
    
    The alternate form specified by the sharp flag for %g and %G verbs
    will not truncate trailing zeros and assume a default precision of 6.
    
    Fixes #18857.
    
    Change-Id: I4d776239e06d7a6a90f2d8556240a359888cb7c3
    Reviewed-on: https://go-review.googlesource.com/37051
    Reviewed-by: Rob Pike <r@golang.org>

commit 1e69aefb7e7ed34f8e425287b126b0f3edbf144e
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Sat Feb 18 07:28:49 2017 +0000

    net/url: document that Query returns only valid values
    
    Fixes #19110
    
    Change-Id: I291fa4ec3c61145162acd019e3f0e5dd3d7c97e9
    Reviewed-on: https://go-review.googlesource.com/37194
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6cfc3b25e95d9f41215e0df8eadfdbeb472594fe
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sat Feb 18 09:32:46 2017 +0100

    math: protect benchmarked functions from being optimized away
    
    Add exported global variables and store the results of benchmarked
    functions in them. This prevents the current compiler optimizations
    from removing the instructions that are needed to compute the return
    values of the benchmarked functions.
    
    Change-Id: If8b08424e85f3796bb6dd73e761c653abbabcc5e
    Reviewed-on: https://go-review.googlesource.com/37195
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6ef92b6e3bce369feeb114dd3267a3f18038fc8c
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sat Feb 18 07:46:41 2017 +0100

    os: remove incorrect detection of O_CLOEXEC flag on darwin
    
    The below range loop will not stop when encountering
    the first '.' character in a Darwin version string like "15.6.0".
    
    for i = range osver {
       if osver[i] != '.' {
             continue
          }
       }
    }
    
    Therefore, the condition i > 2 was always satisfied and
    supportsCloseOnExec was always set to true.
    
    Since the minimum supported version of OSX for go is currently 10.8
    and O_CLOEXEC is implemented from OSX 10.7 on the detection code
    can be removed and support for O_CLOEXEC is always assumed to exist.
    
    Change-Id: Idd10094d8385dd4adebc8d7a6d9e9a8f29455867
    Reviewed-on: https://go-review.googlesource.com/37193
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 497b608faba2a9c450d03e2bca2ef005ea32b933
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Sun Dec 18 07:24:58 2016 +0000

    go/doc: allow : in godoc links
    
    The emphasize function used a complex regexp to find URLs, which
    truncated some types of URL and did not match others.
    This has been simplified and adjusted to allow valid punctuation
    like :: or ! in the path part and :[] in the host part.
    Comments were added to clarify what this regexp allows.
    The path part matches query and fragment also so document this.
    Removed news, telnet, wais, and prospero protocols.
    
    Tests were added for:
     IPV6 URLs
     URLs surrounded by brackets
     URLs containing ::
     URLs containing :;!- in the path
    
    In order to allow punctuation and yet preserve current behaviour,
    URLs are not permitted to end in .,:;?! to allow the use of
    normal punctuation surrounding URLs in comments.
    
    Fixes #18139
    
    Change-Id: I38b2d7a85fe0d171e4bf4aac420f8c2d3ced8a2f
    Reviewed-on: https://go-review.googlesource.com/37192
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a4a3d63dbeb57174ada4b2e5f0fa54c9ec83803b
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 17 15:21:36 2017 -0800

    math/bits: added benchmarks for Leading/TrailingZeros
    
    BenchmarkLeadingZeros-8         200000000                8.80 ns/op
    BenchmarkLeadingZeros8-8        200000000                8.21 ns/op
    BenchmarkLeadingZeros16-8       200000000                7.49 ns/op
    BenchmarkLeadingZeros32-8       200000000                7.80 ns/op
    BenchmarkLeadingZeros64-8       200000000                8.67 ns/op
    
    BenchmarkTrailingZeros-8        1000000000               2.05 ns/op
    BenchmarkTrailingZeros8-8       2000000000               1.94 ns/op
    BenchmarkTrailingZeros16-8      2000000000               1.94 ns/op
    BenchmarkTrailingZeros32-8      2000000000               1.92 ns/op
    BenchmarkTrailingZeros64-8      2000000000               2.03 ns/op
    
    Change-Id: I45497bf2d6369ba6cfc88ded05aa735908af8908
    Reviewed-on: https://go-review.googlesource.com/37220
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 19028bdd18483689a3743639fa89d272cbb96c7b
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 17 15:02:49 2017 -0800

    math/bits: faster Rotate functions, added respective benchmarks
    
    Measured on 2.3 GHz Intel Core i7, running maxOS 10.12.3.
    
    benchmark                    old ns/op     new ns/op     delta
    BenchmarkRotateLeft-8        7.87          7.00          -11.05%
    BenchmarkRotateLeft8-8       8.41          4.52          -46.25%
    BenchmarkRotateLeft16-8      8.07          4.55          -43.62%
    BenchmarkRotateLeft32-8      8.36          4.73          -43.42%
    BenchmarkRotateLeft64-8      7.93          4.78          -39.72%
    
    BenchmarkRotateRight-8       8.23          6.72          -18.35%
    BenchmarkRotateRight8-8      8.76          4.39          -49.89%
    BenchmarkRotateRight16-8     9.07          4.44          -51.05%
    BenchmarkRotateRight32-8     8.85          4.46          -49.60%
    BenchmarkRotateRight64-8     8.11          4.43          -45.38%
    
    Change-Id: I79ea1e9e6fc65f95794a91f860a911efed3aa8a1
    Reviewed-on: https://go-review.googlesource.com/37219
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a12edb8db6f3fa93a1ccd96a0f84b647d08429ef
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 17 14:20:11 2017 -0800

    math/bits: faster OnesCount, added respective benchmarks
    
    Also: Changed Reverse/ReverseBytes implementations to use
    the same (smaller) masks as OnesCount.
    
    BenchmarkOnesCount-8          37.0          6.26          -83.08%
    BenchmarkOnesCount8-8         7.24          1.99          -72.51%
    BenchmarkOnesCount16-8        11.3          2.47          -78.14%
    BenchmarkOnesCount32-8        18.4          3.02          -83.59%
    BenchmarkOnesCount64-8        40.0          3.78          -90.55%
    BenchmarkReverse-8            6.69          6.22          -7.03%
    BenchmarkReverse8-8           1.64          1.64          +0.00%
    BenchmarkReverse16-8          2.26          2.18          -3.54%
    BenchmarkReverse32-8          2.88          2.87          -0.35%
    BenchmarkReverse64-8          5.64          4.34          -23.05%
    BenchmarkReverseBytes-8       2.48          2.17          -12.50%
    BenchmarkReverseBytes16-8     0.63          0.95          +50.79%
    BenchmarkReverseBytes32-8     1.13          1.24          +9.73%
    BenchmarkReverseBytes64-8     2.50          2.16          -13.60%
    
    OnesCount-8       37.0ns ± 0%   6.3ns ± 0%   ~             (p=1.000 n=1+1)
    OnesCount8-8      7.24ns ± 0%  1.99ns ± 0%   ~             (p=1.000 n=1+1)
    OnesCount16-8     11.3ns ± 0%   2.5ns ± 0%   ~             (p=1.000 n=1+1)
    OnesCount32-8     18.4ns ± 0%   3.0ns ± 0%   ~             (p=1.000 n=1+1)
    OnesCount64-8     40.0ns ± 0%   3.8ns ± 0%   ~             (p=1.000 n=1+1)
    Reverse-8         6.69ns ± 0%  6.22ns ± 0%   ~             (p=1.000 n=1+1)
    Reverse8-8        1.64ns ± 0%  1.64ns ± 0%   ~     (all samples are equal)
    Reverse16-8       2.26ns ± 0%  2.18ns ± 0%   ~             (p=1.000 n=1+1)
    Reverse32-8       2.88ns ± 0%  2.87ns ± 0%   ~             (p=1.000 n=1+1)
    Reverse64-8       5.64ns ± 0%  4.34ns ± 0%   ~             (p=1.000 n=1+1)
    ReverseBytes-8    2.48ns ± 0%  2.17ns ± 0%   ~             (p=1.000 n=1+1)
    ReverseBytes16-8  0.63ns ± 0%  0.95ns ± 0%   ~             (p=1.000 n=1+1)
    ReverseBytes32-8  1.13ns ± 0%  1.24ns ± 0%   ~             (p=1.000 n=1+1)
    ReverseBytes64-8  2.50ns ± 0%  2.16ns ± 0%   ~             (p=1.000 n=1+1)
    
    Change-Id: I591b0ffc83fc3a42828256b6e5030f32c64f9497
    Reviewed-on: https://go-review.googlesource.com/37218
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 21c71d77887733c02a7ae31b65b1b12041485ee5
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Fri Feb 10 13:17:20 2017 -0600

    cmd/compile/internal/ssa: combine load + op on AMD64
    
    On AMD64 Most operation can have one operand in memory.
    Combine load and dependand operation into one new operation,
    where possible. I've seen no significant performance changes on go1,
    but this allows to remove ~1.8kb code from go tool. And in math package
    I see e. g.:
    
    Remainder-6            70.0ns ± 0%   64.6ns ± 0%   -7.76%  (p=0.000 n=9+1
    Change-Id: I88b8602b1d55da8ba548a34eb7da4b25d59a297e
    Reviewed-on: https://go-review.googlesource.com/36793
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit a9292b833bdc8d95f738db20a1fc7bcffc0b33d2
Author: Keith Randall <khr@golang.org>
Date:   Fri Feb 17 11:36:08 2017 -0800

    cmd/compile: fix 32-bit unsigned division on 64-bit machines
    
    The type of an intermediate multiply was wrong.  When that
    intermediate multiply was spilled, the top 32 bits were lost.
    
    Fixes #19153
    
    Change-Id: Ib29350a4351efa405935b7f7ee3c112668e64108
    Reviewed-on: https://go-review.googlesource.com/37212
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 4498b6839042fd0eb0e2854bc93fbf26c3f78046
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 17 13:27:09 2017 -0800

    math/bits: faster Reverse, ReverseBytes
    
    - moved from: x&m>>k | x&^m<<k to: x&m>>k | x<<k&m
      This permits use of the same constant m twice (*) which may be
      better for machines that can't use large immediate constants
      directly with an AND instruction and have to load them explicitly.
      *) CPUs don't usually have a &^ instruction, so x&^m becomes x&(^m)
    
    - simplified returns
      This improves the generated code because the compiler recognizes
      x>>k | x<<k as ROT when k is the bitsize of x.
    
    The 8-bit versions of these instructions can be significantly faster
    still if they are replaced with table lookups, as long as the table
    is in cache. If the table is not in cache, table-lookup is probably
    slower, hence the choice of an explicit register-only implementation
    for now.
    
    BenchmarkReverse-8            8.50          6.86          -19.29%
    BenchmarkReverse8-8           2.17          1.74          -19.82%
    BenchmarkReverse16-8          2.89          2.34          -19.03%
    BenchmarkReverse32-8          3.55          2.95          -16.90%
    BenchmarkReverse64-8          6.81          5.57          -18.21%
    BenchmarkReverseBytes-8       3.49          2.48          -28.94%
    BenchmarkReverseBytes16-8     0.93          0.62          -33.33%
    BenchmarkReverseBytes32-8     1.55          1.13          -27.10%
    BenchmarkReverseBytes64-8     2.47          2.47          +0.00%
    
    Reverse-8         8.50ns ± 0%  6.86ns ± 0%   ~             (p=1.000 n=1+1)
    Reverse8-8        2.17ns ± 0%  1.74ns ± 0%   ~             (p=1.000 n=1+1)
    Reverse16-8       2.89ns ± 0%  2.34ns ± 0%   ~             (p=1.000 n=1+1)
    Reverse32-8       3.55ns ± 0%  2.95ns ± 0%   ~             (p=1.000 n=1+1)
    Reverse64-8       6.81ns ± 0%  5.57ns ± 0%   ~             (p=1.000 n=1+1)
    ReverseBytes-8    3.49ns ± 0%  2.48ns ± 0%   ~             (p=1.000 n=1+1)
    ReverseBytes16-8  0.93ns ± 0%  0.62ns ± 0%   ~             (p=1.000 n=1+1)
    ReverseBytes32-8  1.55ns ± 0%  1.13ns ± 0%   ~             (p=1.000 n=1+1)
    ReverseBytes64-8  2.47ns ± 0%  2.47ns ± 0%   ~     (all samples are equal)
    
    Change-Id: I0064de8c7e0e568ca7885d6f7064344bef91a06d
    Reviewed-on: https://go-review.googlesource.com/37215
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c61cf5e6b7920be423ba02bc13f716969265756d
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Feb 15 21:16:49 2017 -0800

    cmd/compile/internal/gc: remove Node.IsStatic field
    
    We can immediately emit static assignment data rather than queueing
    them up to be processed during SSA building.
    
    Passes toolstash -cmp.
    
    Change-Id: I8bcea4b72eafb0cc0b849cd93e9cde9d84f30d5e
    Reviewed-on: https://go-review.googlesource.com/37024
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 3557d546090c7fedd69562c88d20767397de835d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Feb 17 10:27:43 2017 -0500

    cmd/compile: check both syms when folding address into load/store on ARM64
    
    The rules for folding addresses into load/stores checks sym1 is
    not on stack (because the stack offset is not known at that point).
    But sym1 could be nil, which invalidates the check. Check merged
    sym instead.
    
    Fixes #19137.
    
    Change-Id: I8574da22ced1216bb5850403d8f08ec60a8d1005
    Reviewed-on: https://go-review.googlesource.com/37145
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 3a239a6ae44163d43cde40d9b83dc1c5b7359cb2
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 17 12:37:38 2017 -0800

    math/bits: fix benchmarks (make sure calls don't get optimized away)
    
    Sum up function results and store them in an exported (global)
    variable. This prevents the compiler from optimizing away the
    otherwise side-effect free function calls.
    
    We now have more realistic set of benchmark numbers...
    
    Measured on 2.3 GHz Intel Core i7, running maxOS 10.12.3.
    
    Note: These measurements are based on the same "old"
    implementation as the prior measurements (commit 7d5c003).
    
    benchmark                     old ns/op     new ns/op     delta
    BenchmarkReverse-8            72.9          8.50          -88.34%
    BenchmarkReverse8-8           13.2          2.17          -83.56%
    BenchmarkReverse16-8          21.2          2.89          -86.37%
    BenchmarkReverse32-8          36.3          3.55          -90.22%
    BenchmarkReverse64-8          71.3          6.81          -90.45%
    BenchmarkReverseBytes-8       11.2          3.49          -68.84%
    BenchmarkReverseBytes16-8     6.24          0.93          -85.10%
    BenchmarkReverseBytes32-8     7.40          1.55          -79.05%
    BenchmarkReverseBytes64-8     10.5          2.47          -76.48%
    
    Reverse-8         72.9ns ± 0%   8.5ns ± 0%   ~     (p=1.000 n=1+1)
    Reverse8-8        13.2ns ± 0%   2.2ns ± 0%   ~     (p=1.000 n=1+1)
    Reverse16-8       21.2ns ± 0%   2.9ns ± 0%   ~     (p=1.000 n=1+1)
    Reverse32-8       36.3ns ± 0%   3.5ns ± 0%   ~     (p=1.000 n=1+1)
    Reverse64-8       71.3ns ± 0%   6.8ns ± 0%   ~     (p=1.000 n=1+1)
    ReverseBytes-8    11.2ns ± 0%   3.5ns ± 0%   ~     (p=1.000 n=1+1)
    ReverseBytes16-8  6.24ns ± 0%  0.93ns ± 0%   ~     (p=1.000 n=1+1)
    ReverseBytes32-8  7.40ns ± 0%  1.55ns ± 0%   ~     (p=1.000 n=1+1)
    ReverseBytes64-8  10.5ns ± 0%   2.5ns ± 0%   ~     (p=1.000 n=1+1)
    
    Change-Id: I8aef1334b84f6cafd25edccad7e6868b37969efb
    Reviewed-on: https://go-review.googlesource.com/37213
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ddb15cea4a02c403160c2d9772f85c122cbc8248
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 17 11:32:08 2017 -0800

    math/bits: much faster ReverseBytes, added respective benchmarks
    
    Measured on 2.3 GHz Intel Core i7, running maxOS 10.12.3.
    
    benchmark                     old ns/op     new ns/op     delta
    BenchmarkReverseBytes-8       11.4          3.51          -69.21%
    BenchmarkReverseBytes16-8     6.87          0.64          -90.68%
    BenchmarkReverseBytes32-8     7.79          0.65          -91.66%
    BenchmarkReverseBytes64-8     11.6          0.64          -94.48%
    
    name              old time/op  new time/op  delta
    ReverseBytes-8    11.4ns ± 0%   3.5ns ± 0%   ~     (p=1.000 n=1+1)
    ReverseBytes16-8  6.87ns ± 0%  0.64ns ± 0%   ~     (p=1.000 n=1+1)
    ReverseBytes32-8  7.79ns ± 0%  0.65ns ± 0%   ~     (p=1.000 n=1+1)
    ReverseBytes64-8  11.6ns ± 0%   0.6ns ± 0%   ~     (p=1.000 n=1+1)
    
    Change-Id: I67b529652b3b613c61687e9e185e8d4ee40c51a2
    Reviewed-on: https://go-review.googlesource.com/37211
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 7d5c003a3a630dc82e10d72a86ae6103c4d3809a
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Feb 17 11:12:49 2017 -0800

    math/bits: much faster Reverse, added respective benchmarks
    
    Measured on 2.3 GHz Intel Core i7, running maxOS 10.12.3.
    
    name         old time/op  new time/op  delta
    Reverse-8    76.6ns ± 0%   8.1ns ± 0%   ~     (p=1.000 n=1+1)
    Reverse8-8   12.6ns ± 0%   0.6ns ± 0%   ~     (p=1.000 n=1+1)
    Reverse16-8  20.8ns ± 0%   0.6ns ± 0%   ~     (p=1.000 n=1+1)
    Reverse32-8  36.5ns ± 0%   0.6ns ± 0%   ~     (p=1.000 n=1+1)
    Reverse64-8  74.0ns ± 0%   6.4ns ± 0%   ~     (p=1.000 n=1+1)
    
    benchmark                old ns/op     new ns/op     delta
    BenchmarkReverse-8       76.6          8.07          -89.46%
    BenchmarkReverse8-8      12.6          0.64          -94.92%
    BenchmarkReverse16-8     20.8          0.64          -96.92%
    BenchmarkReverse32-8     36.5          0.64          -98.25%
    BenchmarkReverse64-8     74.0          6.38          -91.38%
    
    Change-Id: I6b99b10cee2f2babfe79342b50ee36a45a34da30
    Reviewed-on: https://go-review.googlesource.com/37149
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c4b8dadb4060a8456801ad64c9c5642a737dba19
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Feb 5 23:41:41 2017 -0500

    cmd/compile: fix some types in SSA
    
    These seem not to really matter, but good to be correct.
    
    Change-Id: I02edb9797c3d6739725cfbe4723c75f151acd05e
    Reviewed-on: https://go-review.googlesource.com/36837
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit c4ef597c47a00c3f78916425153aefa171a3b12f
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Feb 1 14:27:40 2017 -0500

    cmd/compile: redo writebarrier pass
    
    SSA's writebarrier pass requires WB store ops are always at the
    end of a block. If we move write barrier insertion into SSA and
    emits normal Store ops when building SSA, this requirement becomes
    impractical -- it will create too many blocks for all the Store
    ops.
    
    Redo SSA's writebarrier pass, explicitly order values in store
    order, so it no longer needs this requirement.
    
    Updates #17583.
    Fixes #19067.
    
    Change-Id: I66e817e526affb7e13517d4245905300a90b7170
    Reviewed-on: https://go-review.googlesource.com/36834
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 98061fa5f3a2410c97625cf5eb5a2cd8816bb558
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Jan 20 10:38:05 2017 -0500

    cmd/compile: re-enable nilcheck removal in same block
    
    Nil check removal in the same block is disabled due to issue 18725:
    because the values are not ordered, a nilcheck may influence a
    value that is logically before it. This CL re-enables same-block
    nilcheck removal by ordering values in store order first.
    
    Updates #18725.
    
    Change-Id: I287a38525230c14c5412cbcdbc422547dabd54f6
    Reviewed-on: https://go-review.googlesource.com/35496
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 81acd308a4577af3f0c3e191b16e125c5d10bbf4
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 16 14:15:23 2017 -0800

    math/bits: expand doc strings for all functions
    
    Follow-up on https://go-review.googlesource.com/36315.
    No functionality change.
    
    For #18616.
    
    Change-Id: Id4df34dd7d0381be06eea483a11bf92f4a01f604
    Reviewed-on: https://go-review.googlesource.com/37140
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 045ad5bab812657a85707e480c29de9144881be1
Author: Koki Ide <niconegoto@yahoo.co.jp>
Date:   Fri Feb 17 00:50:41 2017 +0900

    all: fix a few typos in comments
    
    Change-Id: I0455ffaa51c661803d8013c7961910f920d3c3cc
    Reviewed-on: https://go-review.googlesource.com/37043
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0556e26273f704db73df9e7c4c3d2e8434dec7be
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Tue Dec 13 16:45:55 2016 +0100

    sync: make Mutex more fair
    
    Add new starvation mode for Mutex.
    In starvation mode ownership is directly handed off from
    unlocking goroutine to the next waiter. New arriving goroutines
    don't compete for ownership.
    Unfair wait time is now limited to 1ms.
    Also fix a long standing bug that goroutines were requeued
    at the tail of the wait queue. That lead to even more unfair
    acquisition times with multiple waiters.
    
    Performance of normal mode is not considerably affected.
    
    Fixes #13086
    
    On the provided in the issue lockskew program:
    
    done in 1.207853ms
    done in 1.177451ms
    done in 1.184168ms
    done in 1.198633ms
    done in 1.185797ms
    done in 1.182502ms
    done in 1.316485ms
    done in 1.211611ms
    done in 1.182418ms
    
    name                    old time/op  new time/op   delta
    MutexUncontended-48     0.65ns ± 0%   0.65ns ± 1%     ~           (p=0.087 n=10+10)
    Mutex-48                 112ns ± 1%    114ns ± 1%   +1.69%        (p=0.000 n=10+10)
    MutexSlack-48            113ns ± 0%     87ns ± 1%  -22.65%         (p=0.000 n=8+10)
    MutexWork-48             149ns ± 0%    145ns ± 0%   -2.48%         (p=0.000 n=9+10)
    MutexWorkSlack-48        149ns ± 0%    122ns ± 3%  -18.26%         (p=0.000 n=6+10)
    MutexNoSpin-48           103ns ± 4%    105ns ± 3%     ~           (p=0.089 n=10+10)
    MutexSpin-48             490ns ± 4%    515ns ± 6%   +5.08%        (p=0.006 n=10+10)
    Cond32-48               13.4µs ± 6%   13.1µs ± 5%   -2.75%        (p=0.023 n=10+10)
    RWMutexWrite100-48      53.2ns ± 3%   41.2ns ± 3%  -22.57%        (p=0.000 n=10+10)
    RWMutexWrite10-48       45.9ns ± 2%   43.9ns ± 2%   -4.38%        (p=0.000 n=10+10)
    RWMutexWorkWrite100-48   122ns ± 2%    134ns ± 1%   +9.92%        (p=0.000 n=10+10)
    RWMutexWorkWrite10-48    206ns ± 1%    188ns ± 1%   -8.52%         (p=0.000 n=8+10)
    Cond32-24               12.1µs ± 3%   12.4µs ± 3%   +1.98%         (p=0.043 n=10+9)
    MutexUncontended-24     0.74ns ± 1%   0.75ns ± 1%     ~           (p=0.650 n=10+10)
    Mutex-24                 122ns ± 2%    124ns ± 1%   +1.31%        (p=0.007 n=10+10)
    MutexSlack-24           96.9ns ± 2%  102.8ns ± 2%   +6.11%        (p=0.000 n=10+10)
    MutexWork-24             146ns ± 1%    135ns ± 2%   -7.70%         (p=0.000 n=10+9)
    MutexWorkSlack-24        135ns ± 1%    128ns ± 2%   -5.01%         (p=0.000 n=10+9)
    MutexNoSpin-24           114ns ± 3%    110ns ± 4%   -3.84%        (p=0.000 n=10+10)
    MutexSpin-24             482ns ± 4%    475ns ± 8%     ~           (p=0.286 n=10+10)
    RWMutexWrite100-24      43.0ns ± 3%   43.1ns ± 2%     ~           (p=0.956 n=10+10)
    RWMutexWrite10-24       43.4ns ± 1%   43.2ns ± 1%     ~            (p=0.085 n=10+9)
    RWMutexWorkWrite100-24   130ns ± 3%    131ns ± 3%     ~           (p=0.747 n=10+10)
    RWMutexWorkWrite10-24    191ns ± 1%    192ns ± 1%     ~           (p=0.210 n=10+10)
    Cond32-12               11.5µs ± 2%   11.7µs ± 2%   +1.98%        (p=0.002 n=10+10)
    MutexUncontended-12     1.48ns ± 0%   1.50ns ± 1%   +1.08%        (p=0.004 n=10+10)
    Mutex-12                 141ns ± 1%    143ns ± 1%   +1.63%        (p=0.000 n=10+10)
    MutexSlack-12            121ns ± 0%    119ns ± 0%   -1.65%          (p=0.001 n=8+9)
    MutexWork-12             141ns ± 2%    150ns ± 3%   +6.36%         (p=0.000 n=9+10)
    MutexWorkSlack-12        131ns ± 0%    138ns ± 0%   +5.73%         (p=0.000 n=9+10)
    MutexNoSpin-12          87.0ns ± 1%   83.7ns ± 1%   -3.80%        (p=0.000 n=10+10)
    MutexSpin-12             364ns ± 1%    377ns ± 1%   +3.77%        (p=0.000 n=10+10)
    RWMutexWrite100-12      42.8ns ± 1%   43.9ns ± 1%   +2.41%         (p=0.000 n=8+10)
    RWMutexWrite10-12       39.8ns ± 4%   39.3ns ± 1%     ~            (p=0.433 n=10+9)
    RWMutexWorkWrite100-12   131ns ± 1%    131ns ± 0%     ~            (p=0.591 n=10+9)
    RWMutexWorkWrite10-12    173ns ± 1%    174ns ± 0%     ~            (p=0.059 n=10+8)
    Cond32-6                10.9µs ± 2%   10.9µs ± 2%     ~           (p=0.739 n=10+10)
    MutexUncontended-6      2.97ns ± 0%   2.97ns ± 0%     ~     (all samples are equal)
    Mutex-6                  122ns ± 6%    122ns ± 2%     ~           (p=0.668 n=10+10)
    MutexSlack-6             149ns ± 3%    142ns ± 3%   -4.63%        (p=0.000 n=10+10)
    MutexWork-6              136ns ± 3%    140ns ± 5%     ~           (p=0.077 n=10+10)
    MutexWorkSlack-6         152ns ± 0%    138ns ± 2%   -9.21%         (p=0.000 n=6+10)
    MutexNoSpin-6            150ns ± 1%    152ns ± 0%   +1.50%         (p=0.000 n=8+10)
    MutexSpin-6              726ns ± 0%    730ns ± 1%     ~           (p=0.069 n=10+10)
    RWMutexWrite100-6       40.6ns ± 1%   40.9ns ± 1%   +0.91%         (p=0.001 n=8+10)
    RWMutexWrite10-6        37.1ns ± 0%   37.0ns ± 1%     ~            (p=0.386 n=9+10)
    RWMutexWorkWrite100-6    133ns ± 1%    134ns ± 1%   +1.01%         (p=0.005 n=9+10)
    RWMutexWorkWrite10-6     152ns ± 0%    152ns ± 0%     ~     (all samples are equal)
    Cond32-2                7.86µs ± 2%   7.95µs ± 2%   +1.10%        (p=0.023 n=10+10)
    MutexUncontended-2      8.10ns ± 0%   9.11ns ± 4%  +12.44%         (p=0.000 n=9+10)
    Mutex-2                 32.9ns ± 9%   38.4ns ± 6%  +16.58%        (p=0.000 n=10+10)
    MutexSlack-2            93.4ns ± 1%   98.5ns ± 2%   +5.39%         (p=0.000 n=10+9)
    MutexWork-2             40.8ns ± 3%   43.8ns ± 7%   +7.38%         (p=0.000 n=10+9)
    MutexWorkSlack-2        98.6ns ± 5%  108.2ns ± 2%   +9.80%         (p=0.000 n=10+8)
    MutexNoSpin-2            399ns ± 1%    398ns ± 2%     ~             (p=0.463 n=8+9)
    MutexSpin-2             1.99µs ± 3%   1.97µs ± 1%   -0.81%          (p=0.003 n=9+8)
    RWMutexWrite100-2       37.6ns ± 5%   46.0ns ± 4%  +22.17%         (p=0.000 n=10+8)
    RWMutexWrite10-2        50.1ns ± 6%   36.8ns ±12%  -26.46%         (p=0.000 n=9+10)
    RWMutexWorkWrite100-2    136ns ± 0%    134ns ± 2%   -1.80%          (p=0.001 n=7+9)
    RWMutexWorkWrite10-2     140ns ± 1%    138ns ± 1%   -1.50%        (p=0.000 n=10+10)
    Cond32                  5.93µs ± 1%   5.91µs ± 0%     ~            (p=0.411 n=9+10)
    MutexUncontended        15.9ns ± 0%   15.8ns ± 0%   -0.63%          (p=0.000 n=8+8)
    Mutex                   15.9ns ± 0%   15.8ns ± 0%   -0.44%        (p=0.003 n=10+10)
    MutexSlack              26.9ns ± 3%   26.7ns ± 2%     ~           (p=0.084 n=10+10)
    MutexWork               47.8ns ± 0%   47.9ns ± 0%   +0.21%          (p=0.014 n=9+8)
    MutexWorkSlack          54.9ns ± 3%   54.5ns ± 3%     ~           (p=0.254 n=10+10)
    MutexNoSpin              786ns ± 2%    765ns ± 1%   -2.66%        (p=0.000 n=10+10)
    MutexSpin               3.87µs ± 1%   3.83µs ± 0%   -0.85%          (p=0.005 n=9+8)
    RWMutexWrite100         21.2ns ± 2%   21.0ns ± 1%   -0.88%         (p=0.018 n=10+9)
    RWMutexWrite10          22.6ns ± 1%   22.6ns ± 0%     ~             (p=0.471 n=9+9)
    RWMutexWorkWrite100      132ns ± 0%    132ns ± 0%     ~     (all samples are equal)
    RWMutexWorkWrite10       124ns ± 0%    123ns ± 0%     ~           (p=0.656 n=10+10)
    
    Change-Id: I66412a3a0980df1233ad7a5a0cd9723b4274528b
    Reviewed-on: https://go-review.googlesource.com/34310
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 79f6a5c7bd684f2e6007ee505b522440beb86bf0
Author: Wander Lairson Costa <wcosta@mozilla.com>
Date:   Fri Feb 10 04:10:48 2017 -0200

    syscall: only call setgroups if we need to
    
    If the caller set ups a Credential in os/exec.Command,
    os/exec.Command.Start will end up calling setgroups(2), even if no
    supplementary groups were given.
    
    Only root can call setgroups(2) on BSD kernels, which causes Start to
    fail for non-root users when they try to set uid and gid for the new
    process.
    
    We fix by introducing a new field to syscall.Credential named
    NoSetGroups, and setgroups(2) is only called if it is false.
    We make this field with inverted logic to preserve backward
    compatibility.
    
    RELNOTES=yes
    
    Change-Id: I3cff1f21c117a1430834f640ef21fd4e87e06804
    Reviewed-on: https://go-review.googlesource.com/36697
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 708ba22a0c7b6c2e8f46fccb35998c21c60629b9
Author: Keith Randall <khr@golang.org>
Date:   Mon Feb 13 16:00:09 2017 -0800

    cmd/compile: move constant divide strength reduction to SSA rules
    
    Currently the conversion from constant divides to multiplies is mostly
    done during the walk pass.  This is suboptimal because SSA can
    determine that the value being divided by is constant more often
    (e.g. after inlining).
    
    Change-Id: If1a9b993edd71be37396b9167f77da271966f85f
    Reviewed-on: https://go-review.googlesource.com/37015
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 794f1ebff7aeb4085ce7059011330a5efd946156
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Feb 15 18:43:34 2017 -0800

    cmd/compile: simplify needwritebarrier
    
    Currently, whether we need a write barrier is simply a property of the
    pointer slot being written to.
    
    The only optimization we currently apply using the value being written
    is that pointers to stack variables can omit write barriers because
    they're only written to stack slots... but we already omit write
    barriers for all writes to the stack anyway.
    
    Passes toolstash -cmp.
    
    Change-Id: I7f16b71ff473899ed96706232d371d5b2b7ae789
    Reviewed-on: https://go-review.googlesource.com/37109
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 211102c85f7c7c626853813567188379b5fd7292
Author: Shenghou Ma <minux@golang.org>
Date:   Sat Jan 28 02:01:50 2017 -0500

    math: fix typos in Bessel function docs
    
    While we're at it, also document Yn(0, 0) = -Inf for completeness.
    
    Fixes #18823.
    
    Change-Id: Ib6db68f76d29cc2373c12ebdf3fab129cac8c167
    Reviewed-on: https://go-review.googlesource.com/35970
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 661e2179e54710a83ca1779b9d6ab18c1e2d3679
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 2 16:59:34 2017 -0800

    math/bits: added package for bit-level counting and manipulation
    
    Initial platform-independent implementation.
    
    For #18616.
    
    Change-Id: I4585c55b963101af9059c06c1b8a866cb384754c
    Reviewed-on: https://go-review.googlesource.com/36315
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 1693e7b6f2ad1bd2a800161e92b5ac8d3d882663
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 16 12:52:01 2017 -0800

    cmd/compile/internal/syntax: better errors and recovery for invalid character literals
    
    Fixes #15611.
    
    Change-Id: I352b145026466cafef8cf87addafbd30716bda24
    Reviewed-on: https://go-review.googlesource.com/37138
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 990124da2a6ca5a54b38733b51018e2f8758cfae
Author: Russ Cox <rsc@golang.org>
Date:   Sun Feb 12 13:19:02 2017 -0500

    runtime: use balanced tree for addr lookup in semaphore implementation
    
    CL 36792 fixed #17953, a linear scan caused by n goroutines piling into
    two different locks that hashed to the same bucket in the semaphore table.
    In that CL, n goroutines contending for 2 unfortunately chosen locks
    went from O(n²) to O(n).
    
    This CL fixes a different linear scan, when n goroutines are contending for
    n/2 different locks that all hash to the same bucket in the semaphore table.
    In this CL, n goroutines contending for n/2 unfortunately chosen locks
    goes from O(n²) to O(n log n). This case is much less likely, but any linear
    scan eventually hurts, so we might as well fix it while the problem is fresh
    in our minds.
    
    The new test in this CL checks for both linear scans.
    
    The effect of this CL on the sync benchmarks is negligible
    (but it fixes the new test).
    
    name                      old time/op    new time/op    delta
    Cond1-48                     576ns ±10%     575ns ±13%     ~     (p=0.679 n=71+71)
    Cond2-48                    1.59µs ± 8%    1.61µs ± 9%     ~     (p=0.107 n=73+69)
    Cond4-48                    4.56µs ± 7%    4.55µs ± 7%     ~     (p=0.670 n=74+72)
    Cond8-48                    9.87µs ± 9%    9.90µs ± 7%     ~     (p=0.507 n=69+73)
    Cond16-48                   20.4µs ± 7%    20.4µs ±10%     ~     (p=0.588 n=69+71)
    Cond32-48                   45.4µs ±10%    45.4µs ±14%     ~     (p=0.944 n=73+73)
    UncontendedSemaphore-48     19.7ns ±12%    19.7ns ± 8%     ~     (p=0.589 n=65+63)
    ContendedSemaphore-48       55.4ns ±26%    54.9ns ±32%     ~     (p=0.441 n=75+75)
    MutexUncontended-48         0.63ns ± 0%    0.63ns ± 0%     ~     (all equal)
    Mutex-48                     210ns ± 6%     213ns ±10%   +1.30%  (p=0.035 n=70+74)
    MutexSlack-48                210ns ± 7%     211ns ± 9%     ~     (p=0.184 n=71+72)
    MutexWork-48                 299ns ± 5%     300ns ± 5%     ~     (p=0.678 n=73+75)
    MutexWorkSlack-48            302ns ± 6%     300ns ± 5%     ~     (p=0.149 n=74+72)
    MutexNoSpin-48               135ns ± 6%     135ns ±10%     ~     (p=0.788 n=67+75)
    MutexSpin-48                 693ns ± 5%     689ns ± 6%     ~     (p=0.092 n=65+74)
    Once-48                     0.22ns ±25%    0.22ns ±24%     ~     (p=0.882 n=74+73)
    Pool-48                     5.88ns ±36%    5.79ns ±24%     ~     (p=0.655 n=69+69)
    PoolOverflow-48             4.79µs ±18%    4.87µs ±20%     ~     (p=0.233 n=75+75)
    SemaUncontended-48          0.80ns ± 1%    0.82ns ± 8%   +2.46%  (p=0.000 n=60+74)
    SemaSyntNonblock-48          103ns ± 4%     102ns ± 5%   -1.11%  (p=0.003 n=75+75)
    SemaSyntBlock-48             104ns ± 4%     104ns ± 5%     ~     (p=0.231 n=71+75)
    SemaWorkNonblock-48          128ns ± 4%     129ns ± 6%   +1.51%  (p=0.000 n=63+75)
    SemaWorkBlock-48             129ns ± 8%     130ns ± 7%     ~     (p=0.072 n=75+74)
    RWMutexUncontended-48       2.35ns ± 1%    2.35ns ± 0%     ~     (p=0.144 n=70+55)
    RWMutexWrite100-48           139ns ±18%     141ns ±21%     ~     (p=0.071 n=75+73)
    RWMutexWrite10-48            145ns ± 9%     145ns ± 8%     ~     (p=0.553 n=75+75)
    RWMutexWorkWrite100-48       297ns ±13%     297ns ±15%     ~     (p=0.519 n=75+74)
    RWMutexWorkWrite10-48        588ns ± 7%     585ns ± 5%     ~     (p=0.173 n=73+70)
    WaitGroupUncontended-48     0.87ns ± 0%    0.87ns ± 0%     ~     (all equal)
    WaitGroupAddDone-48         63.2ns ± 4%    62.7ns ± 4%   -0.82%  (p=0.027 n=72+75)
    WaitGroupAddDoneWork-48      109ns ± 5%     109ns ± 4%     ~     (p=0.233 n=75+75)
    WaitGroupWait-48            0.17ns ± 0%    0.16ns ±16%   -8.55%  (p=0.000 n=56+75)
    WaitGroupWaitWork-48        1.78ns ± 1%    2.08ns ± 5%  +16.92%  (p=0.000 n=74+70)
    WaitGroupActuallyWait-48    52.0ns ± 3%    50.6ns ± 5%   -2.70%  (p=0.000 n=71+69)
    
    https://perf.golang.org/search?q=upload:20170215.1
    
    Change-Id: Ia29a8bd006c089e401ec4297c3038cca656bcd0a
    Reviewed-on: https://go-review.googlesource.com/37103
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fc456c7f7b3eed348329483b4ad4014e05d58820
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Feb 15 19:51:05 2017 -0800

    cmd/compile/internal/gc: drop unused src.XPos params in SSA builder
    
    Passes toolstash -cmp.
    
    Change-Id: I037278404ebf762482557e2b6867cbc595074a83
    Reviewed-on: https://go-review.googlesource.com/37023
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 58d762176a6e2216760bc9b56b40665172b9f6fa
Author: Russ Cox <rsc@golang.org>
Date:   Mon Feb 13 09:24:05 2017 -0500

    runtime: run mutexevent profiling without holding semaRoot lock
    
    Suggested by Dmitry in CL 36792 review.
    Clearly safe since there are many different semaRoots
    that could all have profiled sudogs calling mutexevent.
    
    Change-Id: I45eed47a5be3e513b2dad63b60afcd94800e16d1
    Reviewed-on: https://go-review.googlesource.com/37104
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 83f95b85deb97b1f4643362bcd43bee62fd9cc76
Author: Russ Cox <rsc@golang.org>
Date:   Mon Feb 13 10:17:52 2017 -0500

    sync: deflake TestWaitGroupMisuse2
    
    Also runs 100X faster on average, because it takes so many
    fewer attempts to trigger the failure.
    
    Fixes #11443.
    
    Change-Id: I8c39ee48bb3ff6c36fa63083e04076771b65a80d
    Reviewed-on: https://go-review.googlesource.com/36841
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 863035efceaf33e4b7ebfb162930119534e9e5eb
Author: Chris Broadfoot <cbro@golang.org>
Date:   Thu Feb 16 08:29:46 2017 -0800

    doc: document go1.8
    
    Change-Id: Ie2144d001c6b4b2293d07b2acf62d7e3cd0b46a7
    Reviewed-on: https://go-review.googlesource.com/37130
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 0ad247c6f0335f44a27217d0411b8c92e367ebd8
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Tue Feb 14 12:01:01 2017 +1100

    cmd/link: delay calculating pe file parameters after Linkmode is set
    
    For #10776.
    
    Change-Id: Id64a7e35c7cdcd9be16cbe3358402fa379090e36
    Reviewed-on: https://go-review.googlesource.com/36975
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e31144f128e2a491845dc4fcc57d45e22fc1b963
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sun Feb 5 14:19:19 2017 +1100

    cmd/link: set pe section and file alignment to 0 during external linking
    
    This is what gcc does when it generates object files.
    And it is easier to count everything, when it starts from 0.
    Make go linker do the same.
    
    gcc also does not output IMAGE_OPTIONAL_HEADER or
    PE64_IMAGE_OPTIONAL_HEADER for object files.
    Perhaps we should do the same, but not in this CL.
    
    For #10776.
    
    Change-Id: I9789c337648623b6cfaa7d18d1ac9cef32e180dc
    Reviewed-on: https://go-review.googlesource.com/36974
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 64c02460d703f718a44326646d19c417eeb1999a
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu Feb 2 18:37:59 2017 +1100

    debug/pe: add test to check dwarf info
    
    For #10776.
    
    Change-Id: I7931558257c1f6b895e4d44b46d320a54de0d677
    Reviewed-on: https://go-review.googlesource.com/36973
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a6b33312366ad321ea6ac6957e6a53bdd41e892e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Feb 15 13:12:53 2017 -0800

    cmd/compile/internal/gc: skip useless loads for non-SSA params
    
    Change-Id: I78ca43a0f0a6a162a2ade1352e2facb29432d4ac
    Reviewed-on: https://go-review.googlesource.com/37102
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 862fde81fc015720741fcb4ba9593bcc511f9aaf
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Feb 14 10:28:40 2017 -0800

    cmd/compile/internal/gc: document (*state).checkgoto
    
    No behavior change.
    
    Change-Id: I595c15ee976adf21bdbabdf24edf203c9e446185
    Reviewed-on: https://go-review.googlesource.com/36958
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 45a5f79c24677517270083eb56a931192c7e1e7e
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Feb 15 14:26:42 2017 -0800

    internal/poll: define PollDescriptor on plan9
    
    Fixes #19114.
    
    Change-Id: I352add53d6ee8bf78792564225099f8537ac6b46
    Reviewed-on: https://go-review.googlesource.com/37106
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: David du Colombier <0intro@gmail.com>

commit 025dfb130ae131a8005959afa7ee57a2c3778962
Author: Sarah Adams <shadams@google.com>
Date:   Tue Feb 14 15:34:46 2017 -0800

    doc: update Code of Conduct wording and scope
    
    This change removes the punitive language and anonymous reporting mechanism
    from the Code of Conduct document. Read on for the rationale.
    
    More than a year has passed since the Go Code of Conduct was introduced.
    In that time, there have been a small number (<30) of reports to the Working Group.
    Some reports we handled well, with positive outcomes for all involved.
    A few reports we handled badly, resulting in hurt feelings and a bad
    experience for all involved.
    
    On reflection, the reports that had positive outcomes were ones where the
    Working Group took the role of advisor/facilitator, listening to complaints and
    providing suggestions and advice to the parties involved.
    The reports that had negative outcomes were ones where the subject of the
    report felt threatened by the Working Group and Code of Conduct.
    
    After some discussion among the Working Group, we saw that we are most
    effective as facilitators, rather than disciplinarians. The various Go spaces
    already have moderators; this change to the CoC acknowledges their authority
    and places the group in a purely advisory role. If an incident is
    reported to the group we may provide information to or make a
    suggestion the moderators, but the Working Group need not (and should not) have
    any authority to take disciplinary action.
    
    In short, we want it to be clear that the Working Group are here to help
    resolve conflict, period.
    
    The second change made here is the removal of the anonymous reporting mechanism.
    To date, the quality of anonymous reports has been low, and with no way to
    reach out to the reporter for more information there is often very little we
    can do in response. Removing this one-way reporting mechanism strengthens the
    message that the Working Group are here to facilitate a constructive dialogue.
    
    Change-Id: Iee52aff5446accd0dae0c937bb3aa89709ad5fb4
    Reviewed-on: https://go-review.googlesource.com/37014
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit ae1d05981fd97a07e4dc26c37e887a8bfa5ebc89
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Feb 15 12:36:24 2017 -0800

    os: skip TestPipeThreads on Solaris
    
    I don't know why it is not working.  Filed issue 19111 for this.
    
    Fixes build.
    
    Update #19111.
    
    Change-Id: I76f8d6aafba5951da2f3ad7d10960419cca7dd1f
    Reviewed-on: https://go-review.googlesource.com/37092
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0fe62e7575a342decef0d5a00f6740fde15d5d7b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Feb 15 12:32:14 2017 -0800

    os: skip TestPipeThreads on Plan 9
    
    It can't work since Plan 9 does not support the runtime poller.
    
    Fixes build.
    
    Change-Id: I9ec33eb66019d9364c6ff6519b61b32e59498559
    Reviewed-on: https://go-review.googlesource.com/37091
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1f77db94f8a453ae96e490fe729c8c6b0ba9479f
Author: Russ Cox <rsc@golang.org>
Date:   Wed Feb 15 15:41:50 2017 -0500

    runtime: do not call wakep from enlistWorker, to avoid possible deadlock
    
    We have seen one instance of a production job suddenly spinning to
    100% CPU and becoming unresponsive. In that one instance, a SIGQUIT
    was sent after 328 minutes of spinning, and the stacks showed a single
    goroutine in "IO wait (scan)" state.
    
    Looking for things that might get stuck if a goroutine got stuck in
    scanning a stack, we found that injectglist does:
    
            lock(&sched.lock)
            var n int
            for n = 0; glist != nil; n++ {
                    gp := glist
                    glist = gp.schedlink.ptr()
                    casgstatus(gp, _Gwaiting, _Grunnable)
                    globrunqput(gp)
            }
            unlock(&sched.lock)
    
    and that casgstatus spins on gp.atomicstatus until the _Gscan bit goes
    away. Essentially, this code locks sched.lock and then while holding
    sched.lock, waits to lock gp.atomicstatus.
    
    The code that is doing the scan is:
    
            if castogscanstatus(gp, s, s|_Gscan) {
                    if !gp.gcscandone {
                            scanstack(gp, gcw)
                            gp.gcscandone = true
                    }
                    restartg(gp)
                    break loop
            }
    
    More analysis showed that scanstack can, in a rare case, end up
    calling back into code that acquires sched.lock. For example:
    
            runtime.scanstack at proc.go:866
            calls runtime.gentraceback at mgcmark.go:842
            calls runtime.scanstack$1 at traceback.go:378
            calls runtime.scanframeworker at mgcmark.go:819
            calls runtime.scanblock at mgcmark.go:904
            calls runtime.greyobject at mgcmark.go:1221
            calls (*runtime.gcWork).put at mgcmark.go:1412
            calls (*runtime.gcControllerState).enlistWorker at mgcwork.go:127
            calls runtime.wakep at mgc.go:632
            calls runtime.startm at proc.go:1779
            acquires runtime.sched.lock at proc.go:1675
    
    This path was found with an automated deadlock-detecting tool.
    There are many such paths but they all go through enlistWorker -> wakep.
    
    The evidence strongly suggests that one of these paths is what caused
    the deadlock we observed. We're running those jobs with
    GOTRACEBACK=crash now to try to get more information if it happens
    again.
    
    Further refinement and analysis shows that if we drop the wakep call
    from enlistWorker, the remaining few deadlock cycles found by the tool
    are all false positives caused by not understanding the effect of calls
    to func variables.
    
    The enlistWorker -> wakep call was intended only as a performance
    optimization, it rarely executes, and if it does execute at just the
    wrong time it can (and plausibly did) cause the deadlock we saw.
    
    Comment it out, to avoid the potential deadlock.
    
    Fixes #19112.
    Unfixes #14179.
    
    Change-Id: I6f7e10b890b991c11e79fab7aeefaf70b5d5a07b
    Reviewed-on: https://go-review.googlesource.com/37093
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 8833af3f4b891b13d747b1af375cc57e8db29909
Author: Hana Kim <hyangah@gmail.com>
Date:   Tue Feb 14 16:11:35 2017 -0500

    runtime/pprof: print newly added fields of runtime.MemStats
    
    in heap profile with debug mode
    
    Change-Id: I3a80d03a4aa556614626067a8fd698b3b00f4290
    Reviewed-on: https://go-review.googlesource.com/36962
    Reviewed-by: Austin Clements <austin@google.com>

commit 35a95df5710e485f0621cb123dc2e528d0a6146c
Author: Heschi Kreinick <heschi@google.com>
Date:   Tue Feb 7 15:49:43 2017 -0500

    cmd/compile/internal/ssa: display NamedValues in SSA html output.
    
    Change-Id: If268b42b32e6bcd6e7913bffa6e493dc78af40aa
    Reviewed-on: https://go-review.googlesource.com/36539
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Heschi Kreinick <heschi@google.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2ac32b6360b47e3e5bb87ad8cbc51c3d91467c85
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Tue Feb 14 12:30:53 2017 -0500

    cmd/go: improve stale reason for packages
    
    This adds more information to the pkg stale reason for debugging
    purposes.
    
    Change-Id: I7b626db4520baa1127195ae859f4da9b49304636
    Reviewed-on: https://go-review.googlesource.com/36944
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c05b06a12d005f50e4776095a60d6bd9c2c91fac
author: ian lance taylor <iant@golang.org>
date:   fri feb 10 15:17:38 2017 -0800

    os: use poller for file i/o
    
    This changes the os package to use the runtime poller for file I/O
    where possible. When a system call blocks on a pollable descriptor,
    the goroutine will be blocked on the poller but the thread will be
    released to run other goroutines. When using a non-pollable
    descriptor, the os package will continue to use thread-blocking system
    calls as before.
    
    For example, on GNU/Linux, the runtime poller uses epoll. epoll does
    not support ordinary disk files, so they will continue to use blocking
    I/O as before. The poller will be used for pipes.
    
    Since this means that the poller is used for many more programs, this
    modifies the runtime to only block waiting for the poller if there is
    some goroutine that is waiting on the poller. Otherwise, there is no
    point, as the poller will never make any goroutine ready. This
    preserves the runtime's current simple deadlock detection.
    
    This seems to crash FreeBSD systems, so it is disabled on FreeBSD.
    This is issue 19093.
    
    Using the poller on Windows requires opening the file with
    FILE_FLAG_OVERLAPPED. We should only do that if we can remove that
    flag if the program calls the Fd method. This is issue 19098.
    
    Update #6817.
    Update #7903.
    Update #15021.
    Update #18507.
    Update #19093.
    Update #19098.
    
    Change-Id: Ia5197dcefa7c6fbcca97d19a6f8621b2abcbb1fe
    Reviewed-on: https://go-review.googlesource.com/36800
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 81ec3f6a6ccd65abc85fc1f5d16af0a4b426029b
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Feb 15 19:22:45 2017 +1100

    internal/poll: remove unused poll.pollDesc methods
    
    Change-Id: Ic2b20c8238ff0ca5513d32e54ef2945fa4d0c3d2
    Reviewed-on: https://go-review.googlesource.com/37033
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 79fab70a63e03897db384fbf1b5fdeeb6f9f4b92
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Tue Feb 14 13:01:18 2017 +0100

    testing: fix stats bug for sub benchmarks
    
    Fixes golang/go#18815.
    
    Change-Id: Ic9d5cb640a555c58baedd597ed4ca5dd9f275c97
    Reviewed-on: https://go-review.googlesource.com/36990
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d390283ff42c44230ac25800efca231b952fd3ed
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 14 21:34:19 2017 -0800

    cmd/compile/internal/syntax: compiler directives must start at beginning of line
    
    - ignore them, if they don't.
    - added tests
    
    Fixes #18393.
    
    Change-Id: I13f87b81ac6b9138ab5031bb3dd6bebc4c548156
    Reviewed-on: https://go-review.googlesource.com/37020
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a8dc43edd1d04cfd9acabaf1e65fffe1e5bdeb32
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Feb 15 12:47:51 2017 +1100

    internal/testenv: do not delete target file
    
    We did not create it. We should not delete it.
    
    Change-Id: If98454ab233ce25367e11a7c68d31b49074537dd
    Reviewed-on: https://go-review.googlesource.com/37030
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2770c507a54770b38c8654357cf0f47e3f0f3052
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Feb 14 17:13:13 2017 -0800

    cmd/compile: fix position for "missing type in composite literal" error
    
    Fixes #18231.
    
    Change-Id: If1615da4db0e6f0516369a1dc37340d80c78f237
    Reviewed-on: https://go-review.googlesource.com/37018
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 5267ac2732edd1ba4a13773987dff08e8b4a2dde
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 13 16:00:53 2017 -0800

    cmd/compile/internal/syntax: establish principled position information
    
    Until now, the parser set the position for each Node to the position of
    the first token belonging to that node. For compatibility with the now
    defunct gc parser, in many places that position information was modified
    when the gcCompat flag was set (which it was, by default). Furthermore,
    in some places, position information was not set at all.
    
    This change removes the gcCompat flag and all associated code, and sets
    position information for all nodes in a more principled way, as proposed
    by mdempsky (see #16943 for details). Specifically, the position of a
    node may not be at the very beginning of the respective production. For
    instance for an Operation `a + b`, the position associated with the node
    is the position of the `+`. Thus, for `a + b + c` we now get different
    positions for the two additions.
    
    This change does not pass toolstash -cmp because position information
    recorded in export data and pcline tables is different. There are no
    other functional changes.
    
    Added test suite testing the position of all nodes.
    
    Fixes #16943.
    
    Change-Id: I3fc02bf096bc3b3d7d2fa655dfd4714a1a0eb90c
    Reviewed-on: https://go-review.googlesource.com/37017
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6910756f9b8c7a97b1435ec40b8ebff9655611d7
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Tue Feb 14 15:52:44 2017 +0000

    math/big: simplify bool expression
    
    Change-Id: I280c53be455f2fe0474ad577c0f7b7908a4eccb2
    Reviewed-on: https://go-review.googlesource.com/36993
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 72aa757dddad7e915f4faad87aacf8010d91561b
Author: Russ Cox <rsc@golang.org>
Date:   Tue Feb 14 00:17:50 2017 -0500

    encoding/xml: fix incorrect indirect code in chardata, comment, innerxml fields
    
    The new tests in this CL have been checked against Go 1.7 as well
    and all pass in Go 1.7, with the one exception noted in a comment
    (an intentional change to omitempty already present before this CL).
    
    CL 15684 made the intentional change to omitempty.
    This CL fixes bugs introduced along the way.
    
    Most of these are corner cases that are arguably not that important,
    but they've always worked all the way back to Go 1, and someone
    cared enough to file #19063. The most significant problem found
    while adding tests is that in the case of a nil *string field with
    `xml:",chardata"`, the existing code silently stops processing not just
    that field but the entire remainder of the struct.
    Even if #19063 were not worth fixing, this chardata bug would be.
    
    Fixes #19063.
    
    Change-Id: I318cf8f9945e1a4615982d9904e109fde577ebf9
    Reviewed-on: https://go-review.googlesource.com/36954
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Daniel Martí <mvdan@mvdan.cc>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit eebd8f51e8f358575bb5fe2867c96a8fe4605ca7
Author: Bryan C. Mills <bcmills@google.com>
Date:   Tue Feb 14 17:06:57 2017 -0500

    mime: add benchmarks for TypeByExtension and ExtensionsByType
    
    These are possible use-cases for sync.Map.
    
    Updates golang/go#18177
    
    Change-Id: I5e2a3d1249967c37d3f89a41122bf4a90522db11
    Reviewed-on: https://go-review.googlesource.com/36964
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4477fd097fcc16c6d2703ec6228f47c9af030655
Author: Kirill Smelkov <kirr@nexedi.com>
Date:   Thu Dec 1 23:43:21 2016 +0300

    cmd/compile/internal/ssa: combine 2 byte loads + shifts into word load + rolw 8 on AMD64
    
    ... and same for stores. This does for binary.BigEndian.Uint16() what
    was already done for Uint32 and Uint64 with BSWAP in 10f75748 (CL 32222).
    
    Here is how generated code changes e.g. for the following function
    (omitting saying the same prologue/epilogue):
    
            func get16(b [2]byte) uint16 {
                    return binary.BigEndian.Uint16(b[:])
            }
    
    "".get16 t=1 size=21 args=0x10 locals=0x0
    
            // before
            0x0000 00000 (x.go:15)  MOVBLZX "".b+9(FP), AX
            0x0005 00005 (x.go:15)  MOVBLZX "".b+8(FP), CX
            0x000a 00010 (x.go:15)  SHLL    $8, CX
            0x000d 00013 (x.go:15)  ORL     CX, AX
    
            // after
            0x0000 00000 (x.go:15)  MOVWLZX "".b+8(FP), AX
            0x0005 00005 (x.go:15)  ROLW    $8, AX
    
    encoding/binary is speedup overall a bit:
    
    name                    old time/op    new time/op    delta
    ReadSlice1000Int32s-4     4.83µs ± 0%    4.83µs ± 0%     ~     (p=0.206 n=4+5)
    ReadStruct-4              1.29µs ± 2%    1.28µs ± 1%   -1.27%  (p=0.032 n=4+5)
    ReadInts-4                 384ns ± 1%     385ns ± 1%     ~     (p=0.968 n=4+5)
    WriteInts-4                534ns ± 3%     526ns ± 0%   -1.54%  (p=0.048 n=4+5)
    WriteSlice1000Int32s-4    5.02µs ± 0%    5.11µs ± 3%     ~     (p=0.175 n=4+5)
    PutUint16-4               0.59ns ± 0%    0.49ns ± 2%  -16.95%  (p=0.016 n=4+5)
    PutUint32-4               0.52ns ± 0%    0.52ns ± 0%     ~     (all equal)
    PutUint64-4               0.53ns ± 0%    0.53ns ± 0%     ~     (all equal)
    PutUvarint32-4            19.9ns ± 0%    19.9ns ± 1%     ~     (p=0.556 n=4+5)
    PutUvarint64-4            54.5ns ± 1%    54.2ns ± 0%     ~     (p=0.333 n=4+5)
    
    name                    old speed      new speed      delta
    ReadSlice1000Int32s-4    829MB/s ± 0%   828MB/s ± 0%     ~     (p=0.190 n=4+5)
    ReadStruct-4            58.0MB/s ± 2%  58.7MB/s ± 1%   +1.30%  (p=0.032 n=4+5)
    ReadInts-4              78.0MB/s ± 1%  77.8MB/s ± 1%     ~     (p=0.968 n=4+5)
    WriteInts-4             56.1MB/s ± 3%  57.0MB/s ± 0%     ~     (p=0.063 n=4+5)
    WriteSlice1000Int32s-4   797MB/s ± 0%   783MB/s ± 3%     ~     (p=0.190 n=4+5)
    PutUint16-4             3.37GB/s ± 0%  4.07GB/s ± 2%  +20.83%  (p=0.016 n=4+5)
    PutUint32-4             7.73GB/s ± 0%  7.72GB/s ± 0%     ~     (p=0.556 n=4+5)
    PutUint64-4             15.1GB/s ± 0%  15.1GB/s ± 0%     ~     (p=0.905 n=4+5)
    PutUvarint32-4           201MB/s ± 0%   201MB/s ± 0%     ~     (p=0.905 n=4+5)
    PutUvarint64-4           147MB/s ± 1%   147MB/s ± 0%     ~     (p=0.286 n=4+5)
    
    ( "a bit" only because most of the time is spent in reflection-like things
      there, not actual bytes decoding. Even for direct PutUint16 benchmark the
      looping adds overhead and lowers visible benefit. For code-generated encoders /
      decoders actual effect is more than 20% )
    
    Adding Uint32 and Uint64 raw benchmarks too for completeness.
    
    NOTE I had to adjust load-combining rule for bswap case to match first 2 bytes
    loads as result of "2-bytes load+shift" -> "loadw + rorw 8" rewrite. Reason is:
    for loads+shift, even e.g. into uint16 var
    
            var b []byte
            var v uin16
            v = uint16(b[1]) | uint16(b[0])<<8
    
    the compiler eventually generates L(ong) shift - SHLLconst [8], probably
    because it is more straightforward / other reasons to work on the whole
    register. This way 2 bytes rewriting rule is using SHLLconst (not SHLWconst) in
    its pattern, and then it always gets matched first, even if 2-byte rule comes
    syntactically after 4-byte rule in AMD64.rules because 4-bytes rule seemingly
    needs more applyRewrite() cycles to trigger. If 2-bytes rule gets matched for
    inner half of
    
            var b []byte
            var v uin32
            v = uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
    
    and we keep 4-byte load rule unchanged, the result will be MOVW + RORW $8 and
    then series of byte loads and shifts - not one MOVL + BSWAPL.
    
    There is no such problem for stores: there compiler, since it probably knows
    store destination is 2 bytes wide, uses SHRWconst 8 (not SHRLconst 8) and thus
    2-byte store rule is not a subset of rule for 4-byte stores.
    
    Fixes #17151  (int16 was last missing piece there)
    
    Change-Id: Idc03ba965bfce2b94fef456b02ff6742194748f6
    Reviewed-on: https://go-review.googlesource.com/34636
    Reviewed-by: Ilya Tocar <ilya.tocar@intel.com>
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7ffdb757758c086556e5eba277202d9d8940c2bd
Author: Bryan C. Mills <bcmills@google.com>
Date:   Tue Feb 14 01:00:49 2017 -0500

    expvar: add benchmarks for steady-state Map Add calls
    
    Add a benchmark for setting a String value, which we may
    want to treat differently from Int or Float due to the need to support
    Add methods for the latter.
    
    Update tests to use only the exported API instead of making (fragile)
    assumptions about unexported fields.
    
    The existing Map benchmarks construct a new Map for each iteration, which
    focuses the benchmark results on the initial allocation costs for the
    Map and its entries. This change adds variants of the benchmarks which
    use a long-lived map in order to measure steady-state performance for
    Map updates on existing keys.
    
    Updates #18177
    
    Change-Id: I62c920991d17d5898c592446af382cd5c04c528a
    Reviewed-on: https://go-review.googlesource.com/36959
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d2fea0447fbfb26fa675baa8c628da48a23e52b4
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Feb 14 12:46:37 2017 -0500

    math/big: fix s390x test build tags
    
    The tests failed to compile when using the math_big_pure_go tag on
    s390x.
    
    Change-Id: I2a09f53ff6562ab9bc9b886cffc0f6205bbfcfbb
    Reviewed-on: https://go-review.googlesource.com/36956
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 78200799a290da7d53ebbd50c04e432a4ab14eec
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Feb 14 11:01:04 2017 -0500

    cmd/compile: undo special handling of zero-valued STRUCTLIT
    
    CL 35261 introduces special handling of zero-valued STRUCTLIT for
    efficient struct zeroing. But it didn't cover all use cases, for
    example, CONVNOP STRUCTLIT is not handled.
    
    On the other hand, CL 34566 handles zeroing earlier, so we don't
    need the change in CL 35261 for efficient zeroing. Other uses of
    zero-valued struct literals are very rare. So undo the change in
    walk.go in CL 35261.
    
    Add a test for efficient zeroing.
    
    Fixes #19084.
    
    Change-Id: I0807f7423fb44d47bf325b3c1ce9611a14953853
    Reviewed-on: https://go-review.googlesource.com/36955
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit bd91e3569a55e6ceed502422fca9db15f7198c35
Author: Kirill Smelkov <kirr@nexedi.com>
Date:   Thu Dec 1 22:13:16 2016 +0300

    cmd/compile/internal/ssa: generate bswap/store for indexed bigendian byte stores too on AMD64
    
    Commit 10f75748 (CL 32222) added rewrite rules to combine byte loads/stores +
    shifts into larger loads/stores + bswap. For loads both MOVBload and
    MOVBloadidx1 were handled but for store only MOVBstore was there without
    MOVBstoreidx added to rewrite pattern. Fix it.
    
    Here is how generated code changes for the following 2 functions
    (ommitting staying the same prologue/epilogue):
    
        func put32(b []byte, i int, v uint32) {
                binary.BigEndian.PutUint32(b[i:], v)
        }
    
        func put64(b []byte, i int, v uint64) {
                binary.BigEndian.PutUint64(b[i:], v)
        }
    
    "".put32 t=1 size=100 args=0x28 locals=0x0
    
            // before
            0x0032 00050 (x.go:5)   MOVL    CX, DX
            0x0034 00052 (x.go:5)   SHRL    $24, CX
            0x0037 00055 (x.go:5)   MOVQ    "".b+8(FP), BX
            0x003c 00060 (x.go:5)   MOVB    CL, (BX)(AX*1)
            0x003f 00063 (x.go:5)   MOVL    DX, CX
            0x0041 00065 (x.go:5)   SHRL    $16, DX
            0x0044 00068 (x.go:5)   MOVB    DL, 1(BX)(AX*1)
            0x0048 00072 (x.go:5)   MOVL    CX, DX
            0x004a 00074 (x.go:5)   SHRL    $8, CX
            0x004d 00077 (x.go:5)   MOVB    CL, 2(BX)(AX*1)
            0x0051 00081 (x.go:5)   MOVB    DL, 3(BX)(AX*1)
    
            // after
            0x0032 00050 (x.go:5)   BSWAPL  CX
            0x0034 00052 (x.go:5)   MOVQ    "".b+8(FP), DX
            0x0039 00057 (x.go:5)   MOVL    CX, (DX)(AX*1)
    
    "".put64 t=1 size=155 args=0x28 locals=0x0
    
            // before
            0x0037 00055 (x.go:9)   MOVQ    CX, DX
            0x003a 00058 (x.go:9)   SHRQ    $56, CX
            0x003e 00062 (x.go:9)   MOVQ    "".b+8(FP), BX
            0x0043 00067 (x.go:9)   MOVB    CL, (BX)(AX*1)
            0x0046 00070 (x.go:9)   MOVQ    DX, CX
            0x0049 00073 (x.go:9)   SHRQ    $48, DX
            0x004d 00077 (x.go:9)   MOVB    DL, 1(BX)(AX*1)
            0x0051 00081 (x.go:9)   MOVQ    CX, DX
            0x0054 00084 (x.go:9)   SHRQ    $40, CX
            0x0058 00088 (x.go:9)   MOVB    CL, 2(BX)(AX*1)
            0x005c 00092 (x.go:9)   MOVQ    DX, CX
            0x005f 00095 (x.go:9)   SHRQ    $32, DX
            0x0063 00099 (x.go:9)   MOVB    DL, 3(BX)(AX*1)
            0x0067 00103 (x.go:9)   MOVQ    CX, DX
            0x006a 00106 (x.go:9)   SHRQ    $24, CX
            0x006e 00110 (x.go:9)   MOVB    CL, 4(BX)(AX*1)
            0x0072 00114 (x.go:9)   MOVQ    DX, CX
            0x0075 00117 (x.go:9)   SHRQ    $16, DX
            0x0079 00121 (x.go:9)   MOVB    DL, 5(BX)(AX*1)
            0x007d 00125 (x.go:9)   MOVQ    CX, DX
            0x0080 00128 (x.go:9)   SHRQ    $8, CX
            0x0084 00132 (x.go:9)   MOVB    CL, 6(BX)(AX*1)
            0x0088 00136 (x.go:9)   MOVB    DL, 7(BX)(AX*1)
    
            // after
            0x0033 00051 (x.go:9)   BSWAPQ  CX
            0x0036 00054 (x.go:9)   MOVQ    "".b+8(FP), DX
            0x003b 00059 (x.go:9)   MOVQ    CX, (DX)(AX*1)
    
    Updates #17151
    
    Change-Id: I3f4a7f28f210e62e153e60da5abd1d39508cc6c4
    Reviewed-on: https://go-review.googlesource.com/34635
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ilya Tocar <ilya.tocar@intel.com>

commit a0645fcaf9a1a6b76483806f66eedb09ed7c71b5
Author: Kale Blankenship <kale@lemnisys.com>
Date:   Tue Feb 14 08:11:07 2017 -0800

    net/http: document ErrServerClosed
    
    Fixes #19085
    
    Change-Id: Ib11b9a22ea8092aca9e1c9c36b1fb015dd555c4b
    Reviewed-on: https://go-review.googlesource.com/36943
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0993b2fd06a43a2e51b68dd7d8b0643e50c54b9d
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 9 14:11:13 2017 -0500

    runtime: remove g.stackAlloc
    
    Since we're no longer stealing space for the stack barrier array from
    the stack allocation, the stack allocation is simply
    g.stack.hi-g.stack.lo.
    
    Updates #17503.
    
    Change-Id: Id9b450ae12c3df9ec59cfc4365481a0a16b7c601
    Reviewed-on: https://go-review.googlesource.com/36621
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit d089a6c7187f1ff85277515405ec6c641588a7ff
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 9 14:03:49 2017 -0500

    runtime: remove stack barriers
    
    Now that we don't rescan stacks, stack barriers are unnecessary. This
    removes all of the code and structures supporting them as well as
    tests that were specifically for stack barriers.
    
    Updates #17503.
    
    Change-Id: Ia29221730e0f2bbe7beab4fa757f31a032d9690c
    Reviewed-on: https://go-review.googlesource.com/36620
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c5ebcd2c8ac6c5bdf85ec0a346974efd4b0cbe49
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 9 11:50:26 2017 -0500

    runtime: remove rescan list
    
    With the hybrid barrier, rescanning stacks is no longer necessary so
    the rescan list is no longer necessary. Remove it.
    
    This leaves the gcrescanstacks GODEBUG variable, since it's useful for
    debugging, but changes it to simply walk all of the Gs to rescan
    stacks rather than using the rescan list.
    
    We could also remove g.gcscanvalid, which is effectively a distributed
    rescan list. However, it's still useful for gcrescanstacks mode and it
    adds little complexity, so we'll leave it in.
    
    Fixes #17099.
    Updates #17503.
    
    Change-Id: I776d43f0729567335ef1bfd145b75c74de2cc7a9
    Reviewed-on: https://go-review.googlesource.com/36619
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 7aeb915d6b1f90657e67e4722d13624b7109b8d5
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 9 11:36:25 2017 -0500

    runtime: remove unused debug.wbshadow
    
    The wbshadow implementation was removed a year and a half ago in
    1635ab7dfe, but the GODEBUG setting remained. Remove the GODEBUG
    setting since it doesn't do anything.
    
    Change-Id: I19cde324a79472aff60acb5cc9f7d4aa86c0c0ed
    Reviewed-on: https://go-review.googlesource.com/36618
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit a610957f2e48c5b917126b1d0b5f930f05ca8841
Author: Nathan Caza <mastercactapus@gmail.com>
Date:   Fri Feb 10 21:09:21 2017 -0600

    net/http: handle absolute paths in mapDirOpenError
    
    The current implementation does not account for Dir being
    initialized with an absolute path on systems that start
    paths with filepath.Separator. In this scenario, the
    original error is returned, and not checked for file
    segments.
    
    This change adds a test for this case, and corrects the
    behavior by ignoring blank path segments in the loop.
    
    Refs #18984
    
    Change-Id: I9b79fd0a73a46976c8e2feda0283ef0bb2b62ea1
    Reviewed-on: https://go-review.googlesource.com/36804
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ef30a1c8aa68f5226de1ed7397751f15f2956d62
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 13 09:06:08 2017 -0800

    runtime: fix some assembly offset names
    
    For vet. There are more. This is a start.
    
    Change-Id: Ibbbb2b20b5db60ee3fac4a1b5913d18fab01f6b9
    Reviewed-on: https://go-review.googlesource.com/36939
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 785cb7e098b689e9d8c2d4be856f3ffa1825042e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Feb 12 22:56:50 2017 -0800

    all: fix some printf format strings
    
    Appease vet.
    
    Change-Id: Ie88de08b91041990c0eaf2e15628cdb98d40c660
    Reviewed-on: https://go-review.googlesource.com/36938
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cc2a52adef473aa94cbbcc148eef4dfd79259ae7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Feb 12 22:56:28 2017 -0800

    all: use keyed composite literals
    
    Makes vet happy.
    
    Change-Id: I7250f283c96e82b9796c5672a0a143ba7568fa63
    Reviewed-on: https://go-review.googlesource.com/36937
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c0165a38fd3cf00b3ab8ec8c283e44c0c2383d74
Author: Dave Cheney <dave@cheney.net>
Date:   Tue Feb 14 12:51:50 2017 +1100

    internal/poll: only build str.go on plan9
    
    Alternatively the contents of str.go could be moved into fd_io_plan9.go
    
    Change-Id: I9d7ec85bbb376f4244eeca732f25c0b77cadc6a6
    Reviewed-on: https://go-review.googlesource.com/36971
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 84cf1f050dea573e4ef8706ace275d8d13ebe192
Author: Dave Cheney <dave@cheney.net>
Date:   Tue Feb 14 09:18:12 2017 +1100

    internal/poll: remove named return values and naked returns
    
    Change-Id: I283f4453e5cf8b22995b3abffccae182cfbb6945
    Reviewed-on: https://go-review.googlesource.com/36970
    Reviewed-by: Dave Cheney <dave@cheney.net>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 45356c1a082ae0e1f4233a8bc5644d60470e9e52
Author: Caleb Spare <cespare@gmail.com>
Date:   Wed Feb 8 19:47:23 2017 -0800

    time: add Duration.Truncate and Duration.Round
    
    Fixes #18996
    
    Change-Id: I0b0f7270960b368ce97ad4456f60bcc1fc2a8313
    Reviewed-on: https://go-review.googlesource.com/36615
    Run-TryBot: Caleb Spare <cespare@gmail.com>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 46a75870ad5b9b9711e69ffce3738a3ab2057789
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Feb 13 12:46:17 2017 -0800

    runtime: speed up fastrand() % n
    
    This occurs a fair amount in the runtime for non-power-of-two n.
    Use an alternative, faster formulation.
    
    name           old time/op  new time/op  delta
    Fastrandn/2-8  4.45ns ± 2%  2.09ns ± 3%  -53.12%  (p=0.000 n=14+14)
    Fastrandn/3-8  4.78ns ±11%  2.06ns ± 2%  -56.94%  (p=0.000 n=15+15)
    Fastrandn/4-8  4.76ns ± 9%  1.99ns ± 3%  -58.28%  (p=0.000 n=15+13)
    Fastrandn/5-8  4.96ns ±13%  2.03ns ± 6%  -59.14%  (p=0.000 n=15+15)
    
    name                    old time/op  new time/op  delta
    SelectUncontended-8     33.7ns ± 2%  33.9ns ± 2%  +0.70%  (p=0.000 n=49+50)
    SelectSyncContended-8   1.68µs ± 4%  1.65µs ± 4%  -1.54%  (p=0.000 n=50+45)
    SelectAsyncContended-8   282ns ± 1%   277ns ± 1%  -1.50%  (p=0.000 n=48+43)
    SelectNonblock-8        5.31ns ± 1%  5.32ns ± 1%    ~     (p=0.275 n=45+44)
    SelectProdCons-8         585ns ± 3%   577ns ± 2%  -1.35%  (p=0.000 n=50+50)
    GoroutineSelect-8       1.59ms ± 2%  1.59ms ± 1%    ~     (p=0.084 n=49+48)
    
    Updates #16213
    
    Change-Id: Ib555a4d7da2042a25c3976f76a436b536487d5b7
    Reviewed-on: https://go-review.googlesource.com/36932
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 83c58ac710517e0b3fb2654cdba187e45900cca2
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Feb 13 15:33:03 2017 -0800

    internal/poll: return error if WriteConsole fails
    
    Fixes #19068.
    
    Change-Id: Id76037826376b5fe8b588fe3dc02182dfaff8c21
    Reviewed-on: https://go-review.googlesource.com/36935
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 62237c2c8ef26767d61c9fb35adbf6a82de2b312
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Feb 13 15:30:42 2017 -0800

    runtime: if runtime is stale while testing, show StaleReason
    
    Update #19062.
    
    Change-Id: I7397b573389145b56e73d2150ce0fc9aa75b3caa
    Reviewed-on: https://go-review.googlesource.com/36934
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit efb3cab960e7cec3262f41705ec5b69431815411
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 13 15:08:04 2017 -0800

    cmd/compile/internal/syntax: generalize error about var decls in init clauses
    
    Change-Id: I62f9748b97bec245338ebf9686fbf6ad6dc6a9c2
    Reviewed-on: https://go-review.googlesource.com/36931
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f823d305141be6b18f82c875652019eccd0c6679
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 13 12:48:39 2017 -0800

    cmd/compile/internal/syntax: better error for malformed 'if' statements
    
    Use distinction between explicit and automatically inserted semicolons
    to provide a better error message if the condition in an 'if' statement
    is missing.
    
    For #18747.
    
    Change-Id: Iac167ae4e5ad53d2dc73f746b4dee9912434bb59
    Reviewed-on: https://go-review.googlesource.com/36930
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 02de5ed748e419f43d12d2bfaad35fdc3af5143b
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Feb 13 13:34:30 2017 -0800

    cmd/internal/obj: add AddrName type and cleanup AddrType values
    
    Passes toolstash -cmp.
    
    Change-Id: Ida3eda9bd9d79a34c1c3f18cb41aea9392698076
    Reviewed-on: https://go-review.googlesource.com/36950
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit e2948f7efeacf34bab99e33fbbf44c9bd502844a
Author: Kirill Smelkov <kirr@nexedi.com>
Date:   Mon Feb 13 22:28:26 2017 +0300

    cmd/compile: Show arch/os when something in TestAssembly fails
    
    It is not always obvious from the first glance when looking at
    TestAssembly failure in which context the code was generated. For
    example x86 and x86-64 are similar, and those of us who do not work with
    assembly every day can even take s390x version as something similar to x86.
    
    So when something fails lets print the whole test context - this
    includes os and arch which were previously missing. An example failure:
    
    before:
    
    --- FAIL: TestAssembly (40.48s)
            asm_test.go:46: expected:       MOVWZ   \(.*\),
                    go:
                    import "encoding/binary"
                    func f(b []byte) uint32 {
                            return binary.LittleEndian.Uint32(b)
                    }
    
                    asm:"".f t=1 size=160 args=0x20 locals=0x0
                    ...
    
    after:
    
    --- FAIL: TestAssembly (40.43s)
            asm_test.go:46: linux/s390x: expected:  MOVWZ   \(.*\),
                    go:
                    import "encoding/binary"
                    func f(b []byte) uint32 {
                            return binary.LittleEndian.Uint32(b)
                    }
    
                    asm:"".f t=1 size=160 args=0x20 locals=0x0
    
    Motivated-by: #18946#issuecomment-279491071
    
    Change-Id: I61089ceec05da7a165718a7d69dec4227dd0e993
    Reviewed-on: https://go-review.googlesource.com/36881
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8d7722ede284af5da0f4757141e261cdc465db47
Author: Sameer Ajmani <sameer@golang.org>
Date:   Mon Feb 13 14:59:45 2017 -0500

    cmd/go: add "syscall" to the set of packages that run extFiles++
    
    This eliminates the need for syscall/asm.s, which is now empty.
    
    Change-Id: Ied060195e03e9653251f54ea8ef6572444b37fdf
    Reviewed-on: https://go-review.googlesource.com/36844
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 663226d8e130470c1a627c64cf489261ebb6da08
Author: Sokolov Yura <funny.falcon@gmail.com>
Date:   Sun Feb 12 13:18:22 2017 +0300

    runtime: make fastrand to generate 32bit values
    
    Extend period of fastrand from (1<<31)-1 to (1<<32)-1 by
    choosing other polynom and reacting on high bit before shift.
    
    Polynomial is taken at https://users.ece.cmu.edu/~koopman/lfsr/index.html
    from 32.dat.gz . It is referred as F7711115 cause this list of
    polynomials is for LFSR with shift to right (and fastrand uses shift to
    left). (old polynomial is referred in 31.dat.gz as 7BB88888).
    
    There were couple of places with conversation of fastrand to int, which
    leads to negative values on 32bit platforms. They are fixed.
    
    Change-Id: Ibee518a3f9103e0aea220ada494b3aec77babb72
    Reviewed-on: https://go-review.googlesource.com/36875
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 15c62e8535125f096c2425330fe9b561c38e7ee4
Author: Sameer Ajmani <sameer@golang.org>
Date:   Mon Feb 13 14:40:48 2017 -0500

    net/http: document Response.Header values that are subordinate to other fields
    
    I noticed that Content-Length may appear in http.Response.Header, but the docs
    say it should be omitted.  Per discussion with bradfitz@, updating the docs to
    indicate that the struct fields are authoritative.
    
    Change-Id: Id1807ff9d4ba5de425d8b147205f29b18351230f
    Reviewed-on: https://go-review.googlesource.com/36842
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 074b73b1b2db1a1d1bb17d25bb335802e7b59f69
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Feb 13 14:39:58 2017 -0500

    cmd/compile: fix s390x load-combining rules
    
    MOVD{reg,nop} operations (added in CL 36256) inserted to preserve
    type information were blocking the load-combining rules. Fix this
    by merging type changes into loads wherever possible.
    
    Fixes #19059.
    
    Change-Id: I8a1df06eb0f231b40ae43107d4a3bd0b9c441b59
    Reviewed-on: https://go-review.googlesource.com/36843
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 40c27ed5bc9297718809e00e3158967414aa9ee5
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Feb 13 11:46:48 2017 -0800

    runtime: if runtime is stale while testing, show cmd/go output
    
    Update #19062.
    
    Change-Id: If6a4c4f8d12e148b162256f13a8ee423f6e30637
    Reviewed-on: https://go-review.googlesource.com/36918
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5a303aa1e96b669051a18bdc44465396a45037ea
Author: Sameer Ajmani <sameer@golang.org>
Date:   Fri Feb 10 09:22:35 2017 -0500

    syscall: delete the "use" function and calls in non-generated files.
    
    Delete use stub from asm.s, leaving only a dummy file.
    Deleting the file causes Windows build to fail.
    
    Fixes #16607
    
    Change-Id: Ic5a55e042e588f1e1bc6605a3d309d1eabdeb288
    Reviewed-on: https://go-review.googlesource.com/36716
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e9bb9e597e2683839e17f27349cf80ef395a06ce
Author: Chris Manghane <cmang@golang.org>
Date:   Fri Feb 10 14:36:59 2017 -0800

    cmd/go: respect group sticky bit on install.
    
    When installing a package to a different directory using `go build`,
    `mv` cannot be used if the destination directory has the group sticky
    bit set.  Instead, `cp` should be used to make sure the destination
    file has the correct permissions.
    
    Fixes golang/go#18878.
    
    Change-Id: I5423f559e7f84df080ed47816e19a22c6d00ab6d
    Reviewed-on: https://go-review.googlesource.com/36797
    Run-TryBot: Chris Manghane <cmang@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4e0f63940cccfabea084a7608e8ba9c55b8ed952
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Feb 13 11:08:21 2017 -0800

    net: use internal/poll for DragonFly setKeepAlivePeriod
    
    Fixes DragonFly build.
    
    Change-Id: Id6b439cd4023ea8e3ed7cd9b70eec553c9eee4be
    Reviewed-on: https://go-review.googlesource.com/36916
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fc13da1648b513fb6d9f1ec1521a0065727a19b4
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Feb 13 10:57:07 2017 -0800

    internal/poll: only export FD.eofError for testing on posix systems
    
    Fixes build on plan9.
    
    Change-Id: Idbb1e6887c24a873de77c92095198847ed953278
    Reviewed-on: https://go-review.googlesource.com/36915
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 50ab37d008fa830a6a1a7e27856ab248d5de816b
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Sun Feb 12 15:12:52 2017 -0800

    database/sql: convert test timeouts to explicit waits with checks
    
    When testing context cancelation behavior do not rely on context
    timeouts. Use explicit checks in all such tests. In closeDB
    convert the simple check for zero open conns with a wait loop
    for zero open conns.
    
    Fixes #19024
    Fixes #19041
    
    Change-Id: Iecfcc4467e91249fceb21ffd1f7c62c58140d8e9
    Reviewed-on: https://go-review.googlesource.com/36902
    Run-TryBot: Daniel Theophanes <kardianos@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 3792db518327c685da17ca6c6faa4e1d2da4c33c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Feb 10 14:59:38 2017 -0800

    net: refactor poller into new internal/poll package
    
    This will make it possible to use the poller with the os package.
    
    This is a lot of code movement but the behavior is intended to be
    unchanged.
    
    Update #6817.
    Update #7903.
    Update #15021.
    Update #18507.
    
    Change-Id: I1413685928017c32df5654ded73a2643820977ae
    Reviewed-on: https://go-review.googlesource.com/36799
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit b548eee3d96fc0b6e962a243b28121e1f37ad792
Author: Keith Randall <khr@golang.org>
Date:   Mon Feb 13 09:37:06 2017 -0800

    cmd/compile: fix load-combining rules
    
    CL 33632 reorders args of commutative ops in order to make
    CSE for commutative ops more robust.  Unfortunately, that
    broke the load-combining rules which depend on a certain ordering
    of OR ops' arguments.
    
    Introduce some additional rules that order OR ops' arguments
    consistently so that the load-combining rules fire.
    
    Note: there's also something else wrong with the s390x rules.
    I've filed #19059 for that.
    
    Fixes #18946
    
    Change-Id: I0a5447196bd88a55ccee683c69a57b943a9972e1
    Reviewed-on: https://go-review.googlesource.com/36911
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 76b4b8c72dc319454ff3ecb83bf49831e4e528c3
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Fri Feb 10 17:23:55 2017 -0800

    cmd/trace: document the final step to use pprof-like profiles
    
    The tutorial ends without mentioning how to use the generated
    pprof-like profile with the pprof tool. This may be very trivial
    for users who are already very familiar with the Go tools, but
    for the newcomers, it saves a lot of time to finalize the tutorial
    with an example of `go tool pprof` invocation.
    
    Change-Id: Idf034eb4bfb9672ef10190e66fcbf873e8f08f6a
    Reviewed-on: https://go-review.googlesource.com/36803
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>

commit c5fed5bb246b794993cc32bfc22a1a202e385bc5
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Feb 12 21:11:16 2017 -0800

    cmd/compile: cull some dead arch-specific Ops
    
    Change-Id: Iee7daa5b91b7896ce857321e307f2ee47b7f095f
    Reviewed-on: https://go-review.googlesource.com/36906
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 5a75d6a08ebb727c0e2bf5c8fbcbc104d260d302
Author: Keith Randall <khr@golang.org>
Date:   Tue Jan 3 16:15:38 2017 -0800

    cmd/compile: optimize non-empty-interface type conversions
    
    When doing i.(T) for non-empty-interface i and concrete type T,
    there's no need to read the type out of the itab. Just compare the
    itab to the itab we expect for that interface/type pair.
    
    Also optimize type switches by putting the type hash of the
    concrete type in the itab. That way we don't need to load the
    type pointer out of the itab.
    
    Update #18492
    
    Change-Id: I49e280a21e5687e771db5b8a56b685291ac168ce
    Reviewed-on: https://go-review.googlesource.com/34810
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Chase <drchase@google.com>

commit ee2f5fafd88b5ce1404fa40e3645a409e9630897
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Feb 13 09:22:46 2017 -0800

    cmd/compile/internal/parser: don't crash after unexpected token
    
    Added missing nil-check. We will get rid of the gcCompat corrections
    shortly but it's still worthwhile having the new test case added.
    
    Fixes #19056.
    
    Change-Id: I35bd938a4d789058da15724e34c05e5e631ecad0
    Reviewed-on: https://go-review.googlesource.com/36908
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8da91a6297d5960b69ca22d764ef73906f6d61e9
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Feb 11 14:56:16 2017 -0800

    runtime: add Frames example
    
    Based on sample code from iant.
    
    Fixes #18788.
    
    Change-Id: I6bb33ed05af2538fbde42ddcac629280ef7c00a6
    Reviewed-on: https://go-review.googlesource.com/36892
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 39fcf8bf0e3a3b447780983c1f88df5bb9f5bc98
Author: Erik Dubbelboer <erik@dubbelboer.com>
Date:   Thu Feb 9 12:54:25 2017 +0800

    net: use bytes.Equal instead of bytesEqual
    
    bytes.Equal is written in assembly and is slightly faster than the
    current Go bytesEqual from the net package.
    
    benchcmp:
    benchmark                 old ns/op     new ns/op     delta
    BenchmarkIPCompare4-8     7.74          7.01          -9.43%
    BenchmarkIPCompare6-8     8.47          6.86          -19.01%
    
    Change-Id: I2a7ad35867489b46f0943aef5776a2fe1b46e2df
    Reviewed-on: https://go-review.googlesource.com/36850
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 61bf0d1c4033ef2cc6905c2ca6e03046cf54d2bc
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sun Feb 12 18:23:34 2017 +1100

    path/filepath: add test for directory junction walk
    
    For #10424.
    
    Change-Id: Ie4e87503b0ed04f65d2444652bd1db647d3529f4
    Reviewed-on: https://go-review.googlesource.com/36851
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 45c6f59e1fd94ccb11fde61ca8d5b33b3d06dd9f
Author: Russ Cox <rsc@golang.org>
Date:   Fri Feb 10 14:45:41 2017 -0500

    runtime: use two-level list for semaphore address search in semaRoot
    
    If there are many goroutines contending for two different locks
    and both locks hash to the same semaRoot, the scans to find the
    goroutines for a particular lock can end up being O(n), making
    n lock acquisitions quadratic.
    
    As long as only one actively-used lock hashes to each semaRoot
    there's no problem, since the list operations in that case are O(1).
    But when the second actively-used lock hits the same semaRoot,
    then scans for entries with for a given lock have to scan over the
    entries for the other lock.
    
    Fix this problem by changing the semaRoot to hold only one sudog
    per unique address. In the running example, this drops the length of
    that list from O(n) to 2. Then attach other goroutines waiting on the
    same address to a separate list headed by the sudog in the semaRoot list.
    Those "same address list" operations are still O(1), so now the
    example from above works much better.
    
    There is still an assumption here that in real programs you don't have
    many many goroutines queueing up on many many distinct addresses.
    If we end up with that problem, we can replace the top-level list with
    a treap.
    
    Fixes #17953.
    
    Change-Id: I78c5b1a5053845275ab31686038aa4f6db5720b2
    Reviewed-on: https://go-review.googlesource.com/36792
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 93a18acf1e32c37c73e450319b78b5f9f4e11fe7
Author: Cezar Sa Espinola <cezarsa@gmail.com>
Date:   Wed Dec 7 22:45:06 2016 -0200

    image/png: reduce memory allocs encoding images by reusing buffers
    
    This change allows greatly reducing memory allocations with a slightly
    performance improvement as well.
    
    Instances of (*png).Encoder can have a optional BufferPool attached to
    them. This allows reusing temporary buffers used when encoding a new
    image. This buffers include instances to zlib.Writer and bufio.Writer.
    
    Also, buffers for current and previous rows are saved in the encoder
    instance and reused as long as their cap() is enough to fit the current
    image row.
    
    A new benchmark was added to demonstrate the performance improvement
    when setting a BufferPool to an Encoder instance:
    
    $ go test -bench BenchmarkEncodeGray -benchmem
    BenchmarkEncodeGray-4                       1000           2349584 ns/op         130.75 MB/s      852230 B/op         32 allocs/op
    BenchmarkEncodeGrayWithBufferPool-4         1000           2241650 ns/op         137.04 MB/s         900 B/op          3 allocs/op
    
    Change-Id: I4488201ae53cb2ad010c68c1e0118ee12beae14e
    Reviewed-on: https://go-review.googlesource.com/34150
    Reviewed-by: Nigel Tao <nigeltao@golang.org>
    Run-TryBot: Nigel Tao <nigeltao@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5030bfdf81ba3bf4d66cf6e9ddfd80ef194f07b6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Feb 10 14:56:05 2017 -0800

    cmd/internal/obj/x86: add comments to wrapper prologue insertion
    
    Make the comments a bit clearer and more accurate,
    in anticipation of updating the code.
    
    Change-Id: I1111e6c3405a8688fcd29b809a48a762ff41edaa
    Reviewed-on: https://go-review.googlesource.com/36833
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2c91bb4c8af6112074dca24b97b096c888e3b583
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Feb 10 21:49:07 2017 -0800

    cmd/compile: make panicwrap argument-free
    
    When code defines a method on T,
    the compiler generates a corresponding wrapper method on *T.
    The first thing the wrapper does is check whether
    the pointer is nil and if so, call panicwrap.
    This is done to provide a useful error message.
    
    The existing implementation gets its information
    from arguments set up by the compiler.
    However, with some trouble, this information can
    be extracted from the name of the wrapper method itself.
    
    Removing the arguments to panicwrap simplifies and
    shrinks the wrapper method.
    It also means that the call to panicwrap does not
    require any stack space.
    This enables a further optimization on amd64/x86,
    which is to skip the function prologue if nothing
    else in the method requires stack space.
    This is frequently the case in simple, hot methods,
    such as Less and Swap in sort.Interface implementations.
    
    Fixes #19040.
    
    Benchmarks for package sort on amd64:
    
    name                  old time/op  new time/op  delta
    SearchWrappers-8       104ns ± 1%   104ns ± 1%    ~     (p=0.286 n=27+27)
    SortString1K-8         128µs ± 1%   128µs ± 1%  -0.44%  (p=0.004 n=30+30)
    SortString1K_Slice-8   118µs ± 2%   117µs ± 1%    ~     (p=0.106 n=30+30)
    StableString1K-8      18.6µs ± 1%  18.6µs ± 1%    ~     (p=0.446 n=28+26)
    SortInt1K-8           65.9µs ± 1%  60.7µs ± 1%  -7.96%  (p=0.000 n=28+30)
    StableInt1K-8         75.3µs ± 2%  72.8µs ± 1%  -3.41%  (p=0.000 n=30+30)
    StableInt1K_Slice-8   57.7µs ± 1%  57.7µs ± 1%    ~     (p=0.515 n=30+30)
    SortInt64K-8          6.28ms ± 1%  6.01ms ± 1%  -4.19%  (p=0.000 n=28+28)
    SortInt64K_Slice-8    5.04ms ± 1%  5.04ms ± 1%    ~     (p=0.927 n=28+27)
    StableInt64K-8        6.65ms ± 1%  6.38ms ± 1%  -3.97%  (p=0.000 n=26+30)
    Sort1e2-8             37.9µs ± 1%  37.2µs ± 1%  -1.89%  (p=0.000 n=29+27)
    Stable1e2-8           77.0µs ± 1%  74.7µs ± 1%  -3.06%  (p=0.000 n=27+30)
    Sort1e4-8             8.21ms ± 2%  7.98ms ± 1%  -2.77%  (p=0.000 n=29+30)
    Stable1e4-8           24.8ms ± 1%  24.3ms ± 1%  -2.31%  (p=0.000 n=28+30)
    Sort1e6-8              1.27s ± 4%   1.22s ± 1%  -3.42%  (p=0.000 n=30+29)
    Stable1e6-8            5.06s ± 1%   4.92s ± 1%  -2.77%  (p=0.000 n=25+29)
    [Geo mean]             731µs        714µs       -2.29%
    
    Before/after assembly for sort.(*intPairs).Less follows.
    It can be optimized further, but that's for a follow-up CL.
    
    Before:
    
    "".(*intPairs).Less t=1 size=214 args=0x20 locals=0x38
            0x0000 00000 (<autogenerated>:1)        TEXT    "".(*intPairs).Less(SB), $56-32
            0x0000 00000 (<autogenerated>:1)        MOVQ    (TLS), CX
            0x0009 00009 (<autogenerated>:1)        CMPQ    SP, 16(CX)
            0x000d 00013 (<autogenerated>:1)        JLS     204
            0x0013 00019 (<autogenerated>:1)        SUBQ    $56, SP
            0x0017 00023 (<autogenerated>:1)        MOVQ    BP, 48(SP)
            0x001c 00028 (<autogenerated>:1)        LEAQ    48(SP), BP
            0x0021 00033 (<autogenerated>:1)        MOVQ    32(CX), BX
            0x0025 00037 (<autogenerated>:1)        TESTQ   BX, BX
            0x0028 00040 (<autogenerated>:1)        JEQ     55
            0x002a 00042 (<autogenerated>:1)        LEAQ    64(SP), DI
            0x002f 00047 (<autogenerated>:1)        CMPQ    (BX), DI
            0x0032 00050 (<autogenerated>:1)        JNE     55
            0x0034 00052 (<autogenerated>:1)        MOVQ    SP, (BX)
            0x0037 00055 (<autogenerated>:1)        NOP
            0x0037 00055 (<autogenerated>:1)        FUNCDATA        $0, gclocals·4032f753396f2012ad1784f398b170f4(SB)
            0x0037 00055 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
            0x0037 00055 (<autogenerated>:1)        MOVQ    ""..this+64(FP), AX
            0x003c 00060 (<autogenerated>:1)        TESTQ   AX, AX
            0x003f 00063 (<autogenerated>:1)        JEQ     $0, 135
            0x0041 00065 (<autogenerated>:1)        MOVQ    (AX), CX
            0x0044 00068 (<autogenerated>:1)        MOVQ    8(AX), AX
            0x0048 00072 (<autogenerated>:1)        MOVQ    "".i+72(FP), DX
            0x004d 00077 (<autogenerated>:1)        CMPQ    DX, AX
            0x0050 00080 (<autogenerated>:1)        JCC     $0, 128
            0x0052 00082 (<autogenerated>:1)        SHLQ    $4, DX
            0x0056 00086 (<autogenerated>:1)        MOVQ    (CX)(DX*1), DX
            0x005a 00090 (<autogenerated>:1)        MOVQ    "".j+80(FP), BX
            0x005f 00095 (<autogenerated>:1)        CMPQ    BX, AX
            0x0062 00098 (<autogenerated>:1)        JCC     $0, 128
            0x0064 00100 (<autogenerated>:1)        SHLQ    $4, BX
            0x0068 00104 (<autogenerated>:1)        MOVQ    (CX)(BX*1), AX
            0x006c 00108 (<autogenerated>:1)        CMPQ    DX, AX
            0x006f 00111 (<autogenerated>:1)        SETLT   AL
            0x0072 00114 (<autogenerated>:1)        MOVB    AL, "".~r2+88(FP)
            0x0076 00118 (<autogenerated>:1)        MOVQ    48(SP), BP
            0x007b 00123 (<autogenerated>:1)        ADDQ    $56, SP
            0x007f 00127 (<autogenerated>:1)        RET
            0x0080 00128 (<autogenerated>:1)        PCDATA  $0, $1
            0x0080 00128 (<autogenerated>:1)        CALL    runtime.panicindex(SB)
            0x0085 00133 (<autogenerated>:1)        UNDEF
            0x0087 00135 (<autogenerated>:1)        LEAQ    go.string."sort_test"(SB), AX
            0x008e 00142 (<autogenerated>:1)        MOVQ    AX, (SP)
            0x0092 00146 (<autogenerated>:1)        MOVQ    $9, 8(SP)
            0x009b 00155 (<autogenerated>:1)        LEAQ    go.string."intPairs"(SB), AX
            0x00a2 00162 (<autogenerated>:1)        MOVQ    AX, 16(SP)
            0x00a7 00167 (<autogenerated>:1)        MOVQ    $8, 24(SP)
            0x00b0 00176 (<autogenerated>:1)        LEAQ    go.string."Less"(SB), AX
            0x00b7 00183 (<autogenerated>:1)        MOVQ    AX, 32(SP)
            0x00bc 00188 (<autogenerated>:1)        MOVQ    $4, 40(SP)
            0x00c5 00197 (<autogenerated>:1)        PCDATA  $0, $1
            0x00c5 00197 (<autogenerated>:1)        CALL    runtime.panicwrap(SB)
            0x00ca 00202 (<autogenerated>:1)        UNDEF
            0x00cc 00204 (<autogenerated>:1)        NOP
            0x00cc 00204 (<autogenerated>:1)        PCDATA  $0, $-1
            0x00cc 00204 (<autogenerated>:1)        CALL    runtime.morestack_noctxt(SB)
            0x00d1 00209 (<autogenerated>:1)        JMP     0
    
    After:
    
    "".(*intPairs).Swap t=1 size=147 args=0x18 locals=0x8
            0x0000 00000 (<autogenerated>:1)        TEXT    "".(*intPairs).Swap(SB), $8-24
            0x0000 00000 (<autogenerated>:1)        MOVQ    (TLS), CX
            0x0009 00009 (<autogenerated>:1)        SUBQ    $8, SP
            0x000d 00013 (<autogenerated>:1)        MOVQ    BP, (SP)
            0x0011 00017 (<autogenerated>:1)        LEAQ    (SP), BP
            0x0015 00021 (<autogenerated>:1)        MOVQ    32(CX), BX
            0x0019 00025 (<autogenerated>:1)        TESTQ   BX, BX
            0x001c 00028 (<autogenerated>:1)        JEQ     43
            0x001e 00030 (<autogenerated>:1)        LEAQ    16(SP), DI
            0x0023 00035 (<autogenerated>:1)        CMPQ    (BX), DI
            0x0026 00038 (<autogenerated>:1)        JNE     43
            0x0028 00040 (<autogenerated>:1)        MOVQ    SP, (BX)
            0x002b 00043 (<autogenerated>:1)        NOP
            0x002b 00043 (<autogenerated>:1)        FUNCDATA        $0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
            0x002b 00043 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
            0x002b 00043 (<autogenerated>:1)        MOVQ    ""..this+16(FP), AX
            0x0030 00048 (<autogenerated>:1)        TESTQ   AX, AX
            0x0033 00051 (<autogenerated>:1)        JEQ     $0, 140
            0x0035 00053 (<autogenerated>:1)        MOVQ    (AX), CX
            0x0038 00056 (<autogenerated>:1)        MOVQ    8(AX), AX
            0x003c 00060 (<autogenerated>:1)        MOVQ    "".i+24(FP), DX
            0x0041 00065 (<autogenerated>:1)        CMPQ    DX, AX
            0x0044 00068 (<autogenerated>:1)        JCC     $0, 133
            0x0046 00070 (<autogenerated>:1)        SHLQ    $4, DX
            0x004a 00074 (<autogenerated>:1)        MOVQ    8(CX)(DX*1), BX
            0x004f 00079 (<autogenerated>:1)        MOVQ    (CX)(DX*1), SI
            0x0053 00083 (<autogenerated>:1)        MOVQ    "".j+32(FP), DI
            0x0058 00088 (<autogenerated>:1)        CMPQ    DI, AX
            0x005b 00091 (<autogenerated>:1)        JCC     $0, 133
            0x005d 00093 (<autogenerated>:1)        SHLQ    $4, DI
            0x0061 00097 (<autogenerated>:1)        MOVQ    8(CX)(DI*1), AX
            0x0066 00102 (<autogenerated>:1)        MOVQ    (CX)(DI*1), R8
            0x006a 00106 (<autogenerated>:1)        MOVQ    R8, (CX)(DX*1)
            0x006e 00110 (<autogenerated>:1)        MOVQ    AX, 8(CX)(DX*1)
            0x0073 00115 (<autogenerated>:1)        MOVQ    SI, (CX)(DI*1)
            0x0077 00119 (<autogenerated>:1)        MOVQ    BX, 8(CX)(DI*1)
            0x007c 00124 (<autogenerated>:1)        MOVQ    (SP), BP
            0x0080 00128 (<autogenerated>:1)        ADDQ    $8, SP
            0x0084 00132 (<autogenerated>:1)        RET
            0x0085 00133 (<autogenerated>:1)        PCDATA  $0, $1
            0x0085 00133 (<autogenerated>:1)        CALL    runtime.panicindex(SB)
            0x008a 00138 (<autogenerated>:1)        UNDEF
            0x008c 00140 (<autogenerated>:1)        PCDATA  $0, $1
            0x008c 00140 (<autogenerated>:1)        CALL    runtime.panicwrap(SB)
            0x0091 00145 (<autogenerated>:1)        UNDEF
    
    Change-Id: I15bb8435f0690badb868799f313ed8817335efd3
    Reviewed-on: https://go-review.googlesource.com/36809
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 61e963e9c649714a88433d50afa679ec91e05e33
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Feb 11 14:22:33 2017 -0800

    testing: fix copy/paste in docs
    
    Follow-up to CL 36791.
    
    Change-Id: I1c4831e5dfe90c205782e970ada7faff8a009daa
    Reviewed-on: https://go-review.googlesource.com/36890
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1cde87b312c5687cd0d70457c48586055e8c53ca
Author: Dhananjay Nakrani <dhananjaynakrani@gmail.com>
Date:   Fri Dec 23 22:28:45 2016 -0800

    cmd/compile: Ensure left-to-right assignment
    
    Add temporaries to reorder the assignment for OAS2XXX nodes.
    This makes orderstmt(), rewrite
      a, b, c = ...
    as
      tmp1, tmp2, tmp3 = ...
      a, b, c = tmp1, tmp2, tmp3
    and
      a, ok = ...
    as
      t1, t2 = ...
      a  = t1
      ok = t2
    
    Fixes #13433.
    
    Change-Id: Id0f5956e3a254d0a6f4b89b5f7b0e055b1f0e21f
    Reviewed-on: https://go-review.googlesource.com/34713
    Run-TryBot: Dhananjay Nakrani <dhananjayn@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit e0d50a5830cbd17810cd488bf70d86fd0c2757ec
Author: Paul Jolly <paul@myitcv.org.uk>
Date:   Fri Dec 9 11:15:23 2016 +0000

    doc: improve issue template and contribution guidelines
    
    Encourage people towards the various help forums as a first port of
    call. Better sign-posting will reduce the incidence or questions being
    asked in the issue tracker that should otherwise be handled elsewhere,
    thereby keeping the issue tracker email traffic more focussed.
    
    Change-Id: I13b2e498d88be010fca421067ae6fb579a46d6b7
    Reviewed-on: https://go-review.googlesource.com/34250
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 272ec231b722c5aa5eacf9c86ba30206fde3dd5e
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sat Feb 11 20:06:54 2017 +0100

    strings: make parameters names less confusing
    
    Using 'sep' as parameter name for strings functions that take a
    separator argument is fine, but for functions like Index or Count that
    look for a substring it's better to use 'substr' (like Contains
    already does).
    
    Fixes #19039
    
    Change-Id: Idd557409c8fea64ce830ab0e3fec37d3d56a79f0
    Reviewed-on: https://go-review.googlesource.com/36874
    Run-TryBot: Alberto Donizetti <alb.donizetti@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2ceeb7b090464a8da421f72e19267869b86b0901
Author: Remi Gillig <remigillig@gmail.com>
Date:   Sat Feb 11 17:34:48 2017 +0000

    path/filepath: fix TestWinSplitListTestsAreValid on some systems
    
    The environment variables used in those tests override the default
    OS ones. However, one of them (SystemRoot) seems to be required on
    some Windows systems for invoking cmd.exe properly.
    
    This fixes #4930 and #6568.
    
    Change-Id: I23dfb67c1de86020711a3b59513f6adcbba12561
    Reviewed-on: https://go-review.googlesource.com/36873
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 60d7d247a1a5d96152061aa9cfeb5466b90a6787
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Fri Feb 10 14:34:37 2017 -0800

    cmd/nm: extend help text to document the flags
    
    Change-Id: Ia2852666ef44e7ef0bba2360e92caccc83fd0e5c
    Reviewed-on: https://go-review.googlesource.com/36796
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 948b21a3d7419ba42c574bde89c199b522221dc6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Feb 10 11:44:19 2017 -0800

    testing: only call ReadMemStats if necessary when benchmarking
    
    When running benchmarks with -cpuprofile,
    the entire process gets profiled,
    and ReadMemStats is surprisingly expensive.
    Running the sort benchmarks right now with
    -cpuprofile shows almost half of all execution
    time in ReadMemStats.
    
    Since ReadMemStats is not required if the benchmark
    does not need allocation stats, simply skip it.
    This will make cpu profiles nicer to read
    and significantly speed up the process of running benchmarks.
    It might also make sense to toggle cpu profiling
    on/off as we begin/end individual benchmarks,
    but that wouldn't get us the time savings of
    skipping ReadMemStats, so this CL is useful in itself.
    
    Change-Id: I425197b1ee11be4bc91d22b929e2caf648ebd7c5
    Reviewed-on: https://go-review.googlesource.com/36791
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5157039fbd55e2fb7dfac6c461ca743ddced44bd
Author: Nigel Tao <nigeltao@golang.org>
Date:   Fri Feb 10 14:04:59 2017 +1100

    image/color: tweak the YCbCr to RGBA conversion formula again.
    
    The 0x10101 magic constant is a little more principled than 0x10100, as
    the rounding adjustment now spans the complete range [0, 0xffff] instead
    of [0, 0xff00].
    
    Consider this round-tripping code:
    
    y, cb, cr := color.RGBToYCbCr(r0, g0, b0)
    r1, g1, b1 := color.YCbCrToRGB(y, cb, cr)
    
    Due to rounding errors both ways, we often but not always get a perfect
    round trip (where r0 == r1 && g0 == g1 && b0 == b1). This is true both
    before and after this commit. In some cases we got luckier, in others we
    got unluckier.
    
    For example, before this commit, (180, 135, 164) doesn't round trip
    perfectly (it's off by 1) but (180, 135, 165) does. After this commit,
    both cases are reversed: the former does and the latter doesn't (again
    off by 1). Over all possible (r, g, b) triples, there doesn't seem to be
    a big change for better or worse.
    
    There is some history in these CLs:
    
    image/color: tweak the YCbCr to RGBA conversion formula.
    https://go-review.googlesource.com/#/c/12220/2/src/image/color/ycbcr.go
    
    image/color: have YCbCr.RGBA work in 16-bit color, per the Color
    interface.
    https://go-review.googlesource.com/#/c/8073/2/src/image/color/ycbcr.go
    
    Change-Id: Ib25ba7039f49feab2a9d1a4141b86db17db7b3e1
    Reviewed-on: https://go-review.googlesource.com/36732
    Run-TryBot: Nigel Tao <nigeltao@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit bdb9b945b954d01b490a468d97abf9592c98dce9
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Dec 19 10:30:44 2016 -0800

    cmd/compile: eliminate OASWB
    
    Instead we can just call needwritebarrier when constructing the SSA
    representation.
    
    Change-Id: I6fefaad49daada9cdb3050f112889e49dca0047b
    Reviewed-on: https://go-review.googlesource.com/34566
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit e7ec06e000627ad699a274a5a672bc3f63d6a709
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Feb 7 07:28:49 2017 -0800

    cmd/go: copy FFLAGS from build.Package
    
    Fixes #18975.
    
    Change-Id: I60dfb299233ecfed4b2da93750ea84e7921f1fbb
    Reviewed-on: https://go-review.googlesource.com/36482
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Daniel Martí <mvdan@mvdan.cc>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit aa1538799601bdbc881f2ad3133184377b7e7aca
Author: Bryan C. Mills <bcmills@google.com>
Date:   Fri Feb 10 14:46:22 2017 -0500

    expvar: make BenchmarkAdd{Same,Different} comparable to 1.8
    
    bradfitz noted in change 36717 that the new behavior was no longer
    comparable with the old.  This change restores comparable behavior
    for -cpu=1.
    
    BenchmarkMapAddSame                 909           909           +0.00%
    BenchmarkMapAddSame-6               1309          262           -79.98%
    BenchmarkMapAddDifferent            2856          3030          +6.09%
    BenchmarkMapAddDifferent-6          3803          581           -84.72%
    
    updates #18177
    
    Change-Id: Ifaff5a1f48be92002d86c296220313b7efdc81d6
    Reviewed-on: https://go-review.googlesource.com/36723
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 275e1fdb20a3751f5a30f9ec2f0c3fb0b1c8d1d1
Author: Mark Adams <mark@markadams.me>
Date:   Fri Feb 3 08:46:25 2017 -0600

    cmd/go: use Bitbucket v2 REST API when determining VCS
    
    The existing implementation uses v1.0 of Bitbucket's REST API. The newer
    version 2.0 of Bitbucket's REST API provides the same information but
    with support for partial responses allowing the client to request only
    the response fields that are relevant to their usage of the API
    resulting in a much smaller payload size.
    
    The partial response functionality in the Bitbucket API is documented here:
    https://developer.atlassian.com/bitbucket/api/2/reference/meta/partial-response
    
    The v2.0 of the Bitbucket repositories API is documented here:
    https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D#get
    
    Fixes #18919
    
    Change-Id: I319947d5c51adc241cfe3a2228a667cc43fb1f56
    Reviewed-on: https://go-review.googlesource.com/36219
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit a896869a49ade49a029d1f6bf39e5d2f2f4f4254
Author: Alan Donovan <adonovan@google.com>
Date:   Thu Feb 9 14:51:29 2017 -0500

    go/types: unsafe.Pointer is not an alias
    
    Change-Id: Ieb0808caa24c9a5e599084183ba5ee8a6536f7d8
    Reviewed-on: https://go-review.googlesource.com/36622
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit d03c1248604679e1e6a01253144065bc57da48b8
Author: Sokolov Yura <funny.falcon@gmail.com>
Date:   Thu Jan 5 09:36:27 2017 +0300

    runtime: implement fastrand in go
    
    So it could be inlined.
    
    Using bit-tricks it could be implemented without condition
    (improved trick version by Minux Ma).
    
    Simple benchmark shows it is faster on i386 and x86_64, though
    I don't know will it be faster on other architectures?
    
    benchmark                       old ns/op     new ns/op     delta
    BenchmarkFastrand-3             2.79          1.48          -46.95%
    BenchmarkFastrandHashiter-3     25.9          24.9          -3.86%
    
    Change-Id: Ie2eb6d0f598c0bb5fac7f6ad0f8b5e3eddaa361b
    Reviewed-on: https://go-review.googlesource.com/34782
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9f75ecd5e12f2b9988086954933d610cd5647918
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Dec 6 19:55:10 2016 +0000

    runtime/debug: don't run a GC when setting SetGCPercent negative
    
    If the user is calling SetGCPercent(-1), they intend to disable GC.
    They probably don't intend to run one. If they do, they can call
    runtime.GC themselves.
    
    Change-Id: I40ef40dfc7e15193df9ff26159cd30e56b666f73
    Reviewed-on: https://go-review.googlesource.com/34013
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 2a74b9e81405e34c67880866552b5d7bcab74de1
Author: Heschi Kreinick <heschi@google.com>
Date:   Tue Jan 31 14:09:14 2017 -0500

    cmd/trace: Record mark assists in execution traces
    
    During the mark phase of garbage collection, goroutines that allocate
    may be recruited to assist. This change creates trace events for mark
    assists and displays them similarly to sweep assists in the trace
    viewer.
    
    Mark assists are different than sweeps in that they can be preempted, so
    displaying them in the trace viewer is a little tricky -- we may need to
    synthesize multiple slices for one mark assist. This could have been
    done in the parser instead, but I thought it might be preferable to keep
    the parser as true to the event stream as possible.
    
    Change-Id: I381dcb1027a187a354b1858537851fa68a620ea7
    Reviewed-on: https://go-review.googlesource.com/36015
    Run-TryBot: Heschi Kreinick <heschi@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>

commit 249aca5deebcc22cc5169814207416a40f14cd86
Author: Hajime Hoshi <hajimehoshi@gmail.com>
Date:   Sat Feb 11 02:09:57 2017 +0900

    cmd/compile/internal/gc: unexport or remove global functions
    
    Change-Id: Ib2109ab773fbf2a35188300cf91a54735f75fc7c
    Reviewed-on: https://go-review.googlesource.com/36736
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 39651bb50bc0e1038e77e63baa37729bc4676e8a
Author: Bryan C. Mills <bcmills@google.com>
Date:   Thu Feb 9 18:39:16 2017 -0500

    expvar: parallelize BenchmarkMapAdd{Same,Different}
    
    The other expvar tests are already parallelized, and this will help to
    measure the impact of potential implementations for #18177.
    
    updates #18177
    
    Change-Id: I0f4f1a16a0285556cbcc8339855b6459af412675
    Reviewed-on: https://go-review.googlesource.com/36717
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 066ac428cdcb0b220a8a58f31c884d054cecd118
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Feb 10 13:08:13 2017 +1300

    reflect: clear ptrToThis in Ptr when allocating result on heap
    
    Otherwise, calling PtrTo on the result will fail.
    
    Fixes #19003
    
    Change-Id: I8d7d1981a5d0417d5aee52740469d71e90734963
    Reviewed-on: https://go-review.googlesource.com/36731
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 78e6abd244eb4f75180fdb3bc72d69a2f471feca
Author: David R. Jenni <david.r.jenni@gmail.com>
Date:   Wed Feb 8 11:32:58 2017 +0100

    sort: optimize average calculation in symMerge and doPivot.
    
    Change code of the form `i + (j-i)/2` to `int(uint(i+j) >> 1)`.
    
    The optimized average calculation uses fewer instructions to calculate
    the average without overflowing at the addition.
    
    Analogous to https://golang.org/cl/36332.
    
    name                 old time/op  new time/op  delta
    StableString1K-4     49.6µs ± 3%  49.1µs ± 8%    ~     (p=0.659 n=16+19)
    StableInt1K-4         160µs ±10%   148µs ± 5%  -7.29%  (p=0.000 n=20+16)
    StableInt1K_Slice-4   139µs ± 4%   136µs ± 3%  -2.52%  (p=0.000 n=20+16)
    StableInt64K-4       8.84ms ± 6%  8.57ms ± 5%  -3.07%  (p=0.001 n=20+19)
    Stable1e2-4           162µs ±19%   147µs ±16%  -8.79%  (p=0.002 n=20+20)
    Stable1e4-4          31.0ms ± 5%  30.6ms ± 5%    ~     (p=0.221 n=20+20)
    Stable1e6-4           6.37s ± 3%   6.27s ± 2%  -1.67%  (p=0.000 n=19+20)
    
    Change-Id: I1cea0bcb9ace8ef7e48b8fab772e41b4b2170da9
    Reviewed-on: https://go-review.googlesource.com/36570
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 6ac8ccf4b3b7ffe946b99e5031b88edc611e32ec
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Wed Feb 8 15:31:54 2017 -0600

    crypto/sha512: Add AVX2 version for AMD64
    
    name          old time/op    new time/op     delta
    Hash8Bytes-6     913ns ± 0%      667ns ± 0%  -26.91%  (p=0.000 n=10+10)
    Hash1K-6        6.58µs ± 0%     4.23µs ± 0%  -35.69%  (p=0.000 n=10+9)
    Hash8K-6        45.9µs ± 0%     28.1µs ± 0%  -38.93%  (p=0.000 n=10+10)
    
    name          old speed      new speed       delta
    Hash8Bytes-6  8.76MB/s ± 0%  11.99MB/s ± 0%  +36.87%  (p=0.000 n=10+8)
    Hash1K-6       156MB/s ± 0%    242MB/s ± 0%  +55.49%  (p=0.000 n=10+9)
    Hash8K-6       178MB/s ± 0%    292MB/s ± 0%  +63.74%  (p=0.000 n=10+10)
    
    Change-Id: Ic9211d68b02935b2195995f264ec57d6bc36f713
    Reviewed-on: https://go-review.googlesource.com/36630
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 630e93ed2d8a13226903451a0e85e62efd78cdcd
Author: Adam Langley <agl@golang.org>
Date:   Thu Feb 9 15:57:53 2017 -0800

    crypto/x509: ignore CN if SAN extension present.
    
    The code previously tested only whether DNS-name SANs were present in a
    certificate which is only approximately correct. In fact, /any/ SAN
    extension, including one with no DNS names, should cause the CN to be
    ignored.
    
    Change-Id: I3d9824918975be6d4817e7cbb48ed1b0c5a2fc8b
    Reviewed-on: https://go-review.googlesource.com/36696
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a52461686013767d9f3e43d1de6eebf6f92fb62c
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Nov 17 11:53:21 2016 -0500

    cmd/{asm,internal/obj/s390x}, math: remove emulated float instructions
    
    The s390x port was based on the ppc64 port and, because of the way the
    port was done, inherited some instructions from it. ppc64 supports
    3-operand (4-operand for FMADD etc.) floating point instructions
    but s390x doesn't (the destination register is always an input) and
    so these were emulated.
    
    There is a bug in the emulation of FMADD whereby if the destination
    register is also a source for the multiplication it will be
    clobbered. This doesn't break any assembly code in the std lib but
    could affect future work.
    
    To fix this I have gone through the floating point instructions and
    removed all unnecessary 3-/4-operand emulation. The compiler doesn't
    need it and assembly writers don't need it, it's just a source of
    bugs.
    
    I've also deleted the FNMADD family of emulated instructions. They
    aren't used anywhere.
    
    Change-Id: Ic07cedcf141a6a3b43a0c84895460f6cfbf56c04
    Reviewed-on: https://go-review.googlesource.com/33350
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit f44e58703115af61e7b03416273031d788c076f1
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Fri Feb 3 10:36:47 2017 +0100

    math: check overflow in amd64 Exp implementation
    
    Unlike the pure go implementation used by every other architecture,
    the amd64 asm implementation of Exp does not fail early if the
    argument is known to overflow. Make it fail early.
    
    Cost of the check is < 1ns (on an old Sandy Bridge machine):
    
    name   old time/op  new time/op  delta
    Exp-4  18.3ns ± 1%  18.7ns ± 1%  +2.08%  (p=0.000 n=18+20)
    
    Fixes #14932
    Fixes #18912
    
    Change-Id: I04b3f9b4ee853822cbdc97feade726fbe2907289
    Reviewed-on: https://go-review.googlesource.com/36271
    Run-TryBot: Alberto Donizetti <alb.donizetti@gmail.com>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4fafc843a27940c76f43d0d897d918dab2efd529
Author: Tuo Shan <shantuo@google.com>
Date:   Mon Feb 6 18:06:40 2017 -0800

    encoding/json: clarify documention for Unmarshal into a pointer.
    
    Fixes #18730.
    
    Change-Id: If3ef28e62f7e449d4c8dc1dfd78f7d6f5a87ed26
    Reviewed-on: https://go-review.googlesource.com/36478
    Reviewed-by: Russ Cox <rsc@golang.org>

commit a335c344fac8501ecdf49a0654d1701fb48efe61
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Wed Feb 8 21:38:51 2017 -0800

    database/sql: ensure driverConns are closed if not returned to pool
    
    Previously if a connection was requested but timed out during the
    request and when acquiring the db.Lock the connection request
    is fulfilled and the request is unable to be returned to the
    connection pool, then then driver connection would not be closed.
    
    No tests were added or modified because I was unable to determine
    how to trigger this situation without something invasive.
    
    Change-Id: I9d4dc680e3fdcf63d79d212174a5b8b313f363f1
    Reviewed-on: https://go-review.googlesource.com/36641
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 9a7544395ae5075cca2f3f5485f3503f891b5e97
Author: Russ Cox <rsc@golang.org>
Date:   Thu Feb 9 16:01:11 2017 -0500

    runtime/pprof: merge internal/protopprof into pprof package
    
    These are very tightly coupled, and internal/protopprof is small.
    There's no point to having a separate package.
    
    Change-Id: I2c8aa49c9e18a7128657bf2b05323860151b5606
    Reviewed-on: https://go-review.googlesource.com/36711
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3a20928157356f98db74db859b36d744400fc462
Author: Nigel Tao <nigeltao@golang.org>
Date:   Fri Feb 10 14:40:38 2017 +1100

    image: fix the overlap check in Rectangle.Intersect.
    
    This is a re-roll of a previous commit,
    a855da29dbd7a80c4d87a421c1f88a8603c020fa, which was rolled back in
    14347ee480968c712ea885a4ea62779fd8a0dc44.
    
    It was rolled back because it broke a unit test in image/gif. The
    image/gif code was fixed by 9ef65dbe0683634a2e8a557d12267d0309ae1570
    "image/gif: fix frame-inside-image bounds checking".
    
    The original commit message:
    
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
    
    Change-Id: I2e3af1f1686064a573b2e513b39246fe60c03631
    Reviewed-on: https://go-review.googlesource.com/36734
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Nigel Tao <nigeltao@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5faba3057dacdf365572b89b4c9ec9e27f3a6133
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Feb 9 14:00:23 2017 -0800

    cmd/compile: use constants directly for fast map access calls
    
    CL 35554 taught order.go to use static variables
    for constants that needed to be addressable for runtime routines.
    However, there is one class of runtime routines that
    do not actually need an addressable value: fast map access routines.
    This CL teaches order.go to avoid using static variables
    for addressability in those cases.
    Instead, it avoids introducing a temp at all,
    which the backend would just have to optimize away.
    
    Fixes #19015.
    
    Change-Id: I5ef780c604fac3fb48dabb23a344435e283cb832
    Reviewed-on: https://go-review.googlesource.com/36693
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 9ef65dbe0683634a2e8a557d12267d0309ae1570
Author: Nigel Tao <nigeltao@golang.org>
Date:   Thu Feb 9 15:18:47 2017 +1100

    image/gif: fix frame-inside-image bounds checking.
    
    The semantics of the Go image.Rectangle type is that the In and
    Intersects methods treat empty rectangles specially. There are multiple
    valid representations of an empty image.Rectangle. One of them is the
    zero image.Rectangle but there are others. They're obviously not all
    equal in the == sense, so we shouldn't use != to check GIF's semantics.
    
    This change will allow us to re-roll
    a855da29dbd7a80c4d87a421c1f88a8603c020fa "image: fix the overlap check
    in Rectangle.Intersect" which was rolled back in
    14347ee480968c712ea885a4ea62779fd8a0dc44.
    
    Change-Id: Ie1a0d092510a7bb6170e61adbf334b21361ff9e6
    Reviewed-on: https://go-review.googlesource.com/36639
    Run-TryBot: Nigel Tao <nigeltao@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit ee60d39a21f459b07cbc1550606db7ed231799e2
Author: Nathan Caza <mastercactapus@gmail.com>
Date:   Wed Feb 8 18:42:52 2017 -0600

    net/http: improve handling of errors in Dir.Open
    
    The current implementation fails to produce an "IsNotExist" error on some
    platforms (unix) for certain situations where it would be expected. This causes
    downstream consumers, like FileServer, to emit 500 errors instead of a 404 for
    some non-existant paths on certain platforms but not others.
    
    As an example, os.Open("/index.html/foo") on a unix-type system will return
    syscall.ENOTDIR, which os.IsNotExist cannot return true for (because the
    error code is ambiguous without context). On windows, this same example
    would result in os.IsNotExist returning true -- since the returned error is
    specific.
    
    This change alters Dir.Open to look up the tree for an "IsPermission" or
    "IsNotExist" error to return, or a non-directory, returning os.ErrNotExist in
    the last case. For all other error scenarios, the original error is returned.
    This ensures that downstream code, like FileServer, receive errors that behave
    the same across all platforms.
    
    Fixes #18984
    
    Change-Id: Id7d16591c24cd96afddb6d8ae135ac78da42ed37
    Reviewed-on: https://go-review.googlesource.com/36635
    Reviewed-by: Yasuhiro MATSUMOTO <mattn.jp@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 866f63e84eb7096e64b7a39b993c2ca3e943e425
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 9 13:22:37 2017 -0800

    spec: refer to "not defined type" rather than "unnamed type" in conversions
    
    We missed this in https://golang.org/cl/36213.
    Thanks to Chris Hines for pointing it out.
    
    For #18130.
    
    Change-Id: I6279ab19966c4391c4b4458b21fd2527d3f949dd
    Reviewed-on: https://go-review.googlesource.com/36691
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3fd3171c2c2b3f55cb9692b45f2ebb842e9b0b42
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Feb 9 16:00:23 2017 -0800

    cmd/compile/internal/syntax: removed gcCompat code needed to pass orig. tests
    
    The gcCompat mode was introduced to match the new parser's node position
    setup exactly with the positions used by the original parser. Some of the
    gcCompat adjustments were required to satisfy syntax error test cases,
    and the rest were required to make toolstash cmp pass.
    
    This change removes the former gcCompat adjustments and instead adjusts
    the respective test cases as necessary. In some cases this makes the error
    lines consistent with the ones reported by gccgo.
    
    Where it has changed, the position associated with a given syntactic construct
    is the position (line/col number) of the left-most token belonging to the
    construct.
    
    Change-Id: I5b60c00c5999a895c4d6d6e9b383c6405ccf725c
    Reviewed-on: https://go-review.googlesource.com/36695
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 09762ccff79ebf08839f8954ab99c17d41f071e2
Author: Adam Langley <agl@golang.org>
Date:   Wed Nov 30 08:33:17 2016 -0800

    crypto/dsa: also use fromHex in TestSignAndVerify.
    
    This change contains a very minor tidy-up to a test.
    
    Change-Id: I3a8c0168bcdcbf90cacbbac2566c8423c92129f8
    Reviewed-on: https://go-review.googlesource.com/33726
    Reviewed-by: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 450472989b318b7242d2c3d0db4c09d84727c065
Author: Austin Clements <austin@google.com>
Date:   Thu Feb 9 17:08:27 2017 -0500

    cmd/compile: disallow combining nosplit and systemstack
    
    go:systemstack works by tweaking the stack check prologue to check
    against a different bound, while go:nosplit removes the stack check
    prologue entirely. Hence, they can't be used together. Make the build
    fail if they are.
    
    Change-Id: I2d180c4b1d31ff49ec193291ecdd42921d253359
    Reviewed-on: https://go-review.googlesource.com/36710
    Run-TryBot: Austin Clements <austin@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 23df52747d5bd8a0b74f7da7c0c2002b7141f32b
Author: Paulo Flabiano Smorigo <pfsmorigo@linux.vnet.ibm.com>
Date:   Thu Feb 9 18:28:08 2017 -0200

    crypto/aes: fix build failure by changing VORL to VOR
    
    Recently, a commit (85ecc51c) changed the instruction from VORL to VOR.
    
    Fixes #19014
    
    Change-Id: I9a7e0b5771842b1abb5afc73dc41d5e7960cf390
    Reviewed-on: https://go-review.googlesource.com/36625
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4e9874f86e24381fa4305f939d78ed857fe416ca
Author: Bryan C. Mills <bcmills@google.com>
Date:   Thu Feb 9 14:34:38 2017 -0500

    net/rpc: fix aliasing in TestAcceptExitAfterListenerClose
    
    TestRPC writes to newServer and newServerAddr guarded with a
    sync.Once.
    TestAcceptExitAfterListenerClose was overwriting those variables,
    which caused the second invocation of TestRPC within a single process
    to fail.
    
    A second invocation can occur as a result of running the test with
    multiple values for the -cpu flag.
    
    fixes #19001.
    
    Change-Id: I291bacf44aefb49c2264ca0290a28248c026f80e
    Reviewed-on: https://go-review.googlesource.com/36624
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit f791b288d1c65420ae6051cdc180e82716952737
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Feb 9 10:45:35 2017 -0800

    cmd/compile: remove some allocs from CSE
    
    Pick up a few pennies:
    
    * CSE gets run twice for each function,
    but the set of Aux values doesn't change.
    Avoid populating it twice.
    
    * Don't bother populating auxmap for values
    that can't be CSE'd anyway.
    
    name       old alloc/op     new alloc/op     delta
    Template       41.0MB ± 0%      40.7MB ± 0%  -0.61%  (p=0.008 n=5+5)
    Unicode        32.3MB ± 0%      32.3MB ± 0%  -0.22%  (p=0.008 n=5+5)
    GoTypes         122MB ± 0%       121MB ± 0%  -0.55%  (p=0.008 n=5+5)
    Compiler        482MB ± 0%       479MB ± 0%  -0.58%  (p=0.008 n=5+5)
    SSA             865MB ± 0%       862MB ± 0%  -0.35%  (p=0.008 n=5+5)
    Flate          26.5MB ± 0%      26.5MB ± 0%    ~     (p=0.056 n=5+5)
    GoParser       32.6MB ± 0%      32.4MB ± 0%  -0.58%  (p=0.008 n=5+5)
    Reflect        84.2MB ± 0%      83.8MB ± 0%  -0.57%  (p=0.008 n=5+5)
    Tar            27.7MB ± 0%      27.6MB ± 0%  -0.37%  (p=0.008 n=5+5)
    XML            44.7MB ± 0%      44.5MB ± 0%  -0.53%  (p=0.008 n=5+5)
    
    name       old allocs/op    new allocs/op    delta
    Template         373k ± 0%        373k ± 1%    ~     (p=1.000 n=5+5)
    Unicode          326k ± 0%        325k ± 0%    ~     (p=0.548 n=5+5)
    GoTypes         1.16M ± 0%       1.16M ± 0%    ~     (p=0.841 n=5+5)
    Compiler        4.16M ± 0%       4.15M ± 0%    ~     (p=0.222 n=5+5)
    SSA             7.57M ± 0%       7.56M ± 0%  -0.22%  (p=0.008 n=5+5)
    Flate            238k ± 1%        239k ± 1%    ~     (p=0.690 n=5+5)
    GoParser         304k ± 0%        304k ± 0%    ~     (p=1.000 n=5+5)
    Reflect         1.01M ± 0%       1.00M ± 0%  -0.31%  (p=0.016 n=4+5)
    Tar              245k ± 0%        245k ± 1%    ~     (p=0.548 n=5+5)
    XML              393k ± 0%        391k ± 1%    ~     (p=0.095 n=5+5)
    
    Change-Id: I78f1ffe129bd8fd590b7511717dd2bf9f5ecbd6d
    Reviewed-on: https://go-review.googlesource.com/36690
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Daniel Martí <mvdan@mvdan.cc>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 06e5a558207f1d2bc6f61365b87925120c8549b3
Author: Paulo Flabiano Smorigo <pfsmorigo@linux.vnet.ibm.com>
Date:   Thu Dec 8 12:06:34 2016 -0200

    crypto/aes: improve performance for aes on ppc64le
    
    Add asm implementation for AES in order to make use of VMX cryptographic
    acceleration instructions for POWER8. There is a speed boost of over 10
    times using those instructions:
    
    Fixes #18076
    
                            old ns/op  new ns/op  delta
    BenchmarkEncrypt-20     337        30.3       -91.00%
    BenchmarkDecrypt-20     347        30.5a      -91.21%
    BenchmarkExpand-20      1180       130        -88.98%
    
                            old MB/s   new MB/s   speedup
    BenchmarkEncrypt-20     47.38      527.68     11.13x
    BenchmarkDecrypt-20     46.05      524.45     11.38x
    
    Change-Id: Ifa4d1b508f4803cc72dcaad97acc8495d651b019
    Reviewed-on: https://go-review.googlesource.com/33587
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit f02dda50e8b4e268987d269e22b6d7410a52587b
Author: Adam Langley <agl@golang.org>
Date:   Wed Feb 8 10:06:34 2017 -0800

    crypto/tls: don't hold lock when closing underlying net.Conn.
    
    There's no need to hold the handshake lock across this call and it can
    lead to deadlocks if the net.Conn calls back into the tls.Conn.
    
    Fixes #18426.
    
    Change-Id: Ib1b2813cce385949d970f8ad2e52cfbd1390e624
    Reviewed-on: https://go-review.googlesource.com/36561
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 695f12c21a217e0116a80c2c1a518d382cfea22e
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Fri Feb 3 08:40:18 2017 -0500

    cmd/compile: rules change to use ANDN more effectively on ppc64x
    
    Currently there are cases where an XOR with -1 followed by an AND
    is generanted when it could be done with just an ANDN instruction.
    
    Changes to PPC64.rules and required files allows this change
    in generated code.  Examples of this occur in sha3 among others.
    
    Fixes: #18918
    
    Change-Id: I647cb9b4a4aaeebb27db85f8bf75487d78f720c9
    Reviewed-on: https://go-review.googlesource.com/36218
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>

commit e24228af2566593a18932b563e548d288ea3cbb7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Dec 6 20:54:41 2016 -0800

    runtime: enable/disable SIGPROF if needed when profiling
    
    This ensures that SIGPROF is handled correctly when using
    runtime/pprof in a c-archive or c-shared library.
    
    Separate profiler handling into pre-process changes and per-thread
    changes. Simplify the Windows code slightly accordingly.
    
    Fixes #18220.
    
    Change-Id: I5060f7084c91ef0bbe797848978bdc527c312777
    Reviewed-on: https://go-review.googlesource.com/34018
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Austin Clements <austin@google.com>

commit 6a29806e018acc3bf0c17c9d6b946b9ba8edeb4d
Author: Adam Langley <agl@golang.org>
Date:   Wed Feb 8 09:41:39 2017 -0800

    crypto/x509: sort the list of fields used by CreateCertificateRequest.
    
    Change-Id: I67589cb9e728e6c7df5ef6e981189193154338d3
    Reviewed-on: https://go-review.googlesource.com/36559
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 7853b090dd8e9f02147581c4e7bdbb45021e787c
Author: Adam Langley <agl@golang.org>
Date:   Wed Feb 8 09:40:14 2017 -0800

    crypto/x509: CreateCertificateRequest reads ExtraExtensions, not Extensions.
    
    Fixes #18899.
    
    Change-Id: I6a4bf0aad9cf1dbe6691ba4e4c478fcb33c44528
    Reviewed-on: https://go-review.googlesource.com/36558
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 95011d4e018c9499c8a4739bc728e38ddc12b9e1
Author: Adam Langley <agl@golang.org>
Date:   Wed Feb 8 09:32:24 2017 -0800

    crypto/x509: sort the list of fields used by CreateCertificate.
    
    Change-Id: I20f4419ca377ee9428075e42db0bad46a75d983f
    Reviewed-on: https://go-review.googlesource.com/36557
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bfe7c81906212c69df754d5f4e3baa6413f092b3
Author: Adam Langley <agl@golang.org>
Date:   Wed Feb 8 09:28:53 2017 -0800

    crypto/x509: document AuthorityKeyId and don't mutate it.
    
    The AuthorityKeyId value from the template was used by
    CreateCertificate, but that wasn't documented. Also, CreateCertificate
    would stash a value in the template if it needed to override it, which
    was wrong: template should be read-only.
    
    Fixes #18962.
    
    Change-Id: Ida15c54c341e5bbf553756e8aa65021d8085f453
    Reviewed-on: https://go-review.googlesource.com/36556
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 38b3661b45ff62de28abc20057503422337330f7
Author: Francesc Campoy <campoy@golang.org>
Date:   Wed Feb 8 19:12:19 2017 -0800

    plugin: remove unnecessary import "C" from example
    
    It seems that it is not needed to import the pseudo package "C"
    for the plugin to be built correctly.
    Removing it to avoid confusion.
    
    Change-Id: I62838a953ad2889881bfbfd1a36141661565f033
    Reviewed-on: https://go-review.googlesource.com/36638
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit bd2f7c7c411bdbd75a194e745bec222390c1c99e
Author: Sameer Ajmani <sameer@golang.org>
Date:   Thu Feb 9 10:44:26 2017 -0500

    syscall: remove "use" function and calls from generated code.
    
    Update syscall code generators to set build tags.
    
    Regenerate zsyscall files, which makes the following changes:
    - remove calls to "use"
    - update build tags, adding missing ones in some cases
    - "stat" renamed to "st" in some cases
    - "libc_Utimes" renamed "libc_utimes" in one case
    
    I'll mirror this change to x/sys/unix once committed.
    
    Change-Id: Ic07e0ae1433dd133eb57e8dd2a3b86a62aab4eda
    Reviewed-on: https://go-review.googlesource.com/36616
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 85ecc51c4852c49b5e5bc4a52067210c58e266e0
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Thu Feb 2 17:59:18 2017 -0200

    cmd/asm, cmd/internal/obj/ppc64: Add ISA 2.05, 2.06 and 2.07 instructions.
    
    This change adds instructions from ISA 2.05, 2.06 and 2.07 that are frequently
    used in assembly optimizations for ppc64.
    
    It also fixes two problems:
    
      * the implementation of RLDICR[CC]/RLDICL[CC] did not consider all possible
      cases for the bit mask.
      * removed two non-existing instructions that were added by mistake in the VMX
      implementation (VORL/VANDL).
    
    Change-Id: Iaef4e5c6a5240c2156c6c0f28ad3bcd8780e9830
    Reviewed-on: https://go-review.googlesource.com/36230
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit 06637fb3146cd7fd19efc7463cd9d2c0006bff67
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 27 14:14:05 2017 -0500

    text/template: fix method lookup on addressable nil pointer
    
    Fixes #18816.
    
    Change-Id: I4f8f1cac2680dbde492c56d3a5a038577605e7c1
    Reviewed-on: https://go-review.googlesource.com/36542
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e4371fb179ad69cbd057f2430120843948e09f2f
Author: Russ Cox <rsc@golang.org>
Date:   Fri Feb 3 19:26:13 2017 -0500

    time: optimize Now on darwin, windows
    
    Fetch both monotonic and wall time together when possible.
    Avoids skew and is cheaper.
    
    Also shave a few ns off in conversion in package time.
    
    Compared to current implementation (after monotonic changes):
    
    name   old time/op  new time/op  delta
    Now    19.6ns ± 1%   9.7ns ± 1%  -50.63%  (p=0.000 n=41+49) darwin/amd64
    Now    23.5ns ± 4%  10.6ns ± 5%  -54.61%  (p=0.000 n=30+28) windows/amd64
    Now    54.5ns ± 5%  29.8ns ± 9%  -45.40%  (p=0.000 n=27+29) windows/386
    
    More importantly, compared to Go 1.8:
    
    name   old time/op  new time/op  delta
    Now     9.5ns ± 1%   9.7ns ± 1%   +1.94%  (p=0.000 n=41+49) darwin/amd64
    Now    12.9ns ± 5%  10.6ns ± 5%  -17.73%  (p=0.000 n=30+28) windows/amd64
    Now    15.3ns ± 5%  29.8ns ± 9%  +94.36%  (p=0.000 n=30+29) windows/386
    
    This brings time.Now back in line with Go 1.8 on darwin/amd64 and windows/amd64.
    
    It's not obvious why windows/386 is still noticeably worse than Go 1.8,
    but it's better than before this CL. The windows/386 speed is not too
    important; the changes just keep the two architectures similar.
    
    Change-Id: If69b94970c8a1a57910a371ee91e0d4e82e46c5d
    Reviewed-on: https://go-review.googlesource.com/36428
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3a6842a0ecf66cf06ce4f0a5fcb9c09fbfdbecc1
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Feb 9 12:32:33 2017 +0900

    database/sql: replace the expr of timeunit * N with N * timeunit in test
    
    Change-Id: I97981b30a9629916f896cb989cc2a42a8bdbef47
    Reviewed-on: https://go-review.googlesource.com/36672
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5630d39f0c726037e28b16b6d80bba64b848067a
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Feb 9 12:28:51 2017 +0900

    database/sql: fix nits in test
    
    Change-Id: I451b33d8da8d97917f2b257e6a25392c6e6582db
    Reviewed-on: https://go-review.googlesource.com/36671
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3c22e5ca271ca4ad17b34d965e6a9baf7b5cf94d
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Feb 8 17:56:41 2017 -0800

    cmd/compile/internal/parser: improved syntax error for incorrect if/for/switch header
    
    Starting the error message with "expecting" rather than "missing"
    causes the syntax error mechanism to add additional helpful info
    (it recognizes "expecting" but not "missing").
    
    Fixes #17328.
    
    Change-Id: I8482ca5e5a6a6b22e0ed0d831b7328e264156334
    Reviewed-on: https://go-review.googlesource.com/36637
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7ad512e7ffe576c4894ea84b02e954846fbda643
Author: Caleb Spare <cespare@gmail.com>
Date:   Wed Feb 8 13:05:25 2017 -0800

    time: format negative monotonic times correctly in Time.String
    
    Fixes #18993
    
    Change-Id: Ia1fa20b6d82384b07e9ba5512b909439e0bec2a5
    Reviewed-on: https://go-review.googlesource.com/36611
    Run-TryBot: Caleb Spare <cespare@gmail.com>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9799622f09ba2ece6fa8eb7607d0d471d75d9915
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Feb 8 17:30:45 2017 -0800

    cmd/compile/internal/syntax: differentiate between ';' and '\n' in syntax errors
    
    Towards better syntax error messages: With this change, the parser knows whether
    a semicolon was an actual ';' in the source, or whether it was an automatically
    inserted semicolon as result of a '\n' or EOF. Using this information in error
    messages makes them more understandable.
    
    For #17328.
    
    Change-Id: I8cd9accee8681b62569d0ecef922d38682b401eb
    Reviewed-on: https://go-review.googlesource.com/36636
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 4f6d4bb3f4461e7e25eff24254115b689495e834
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Wed Feb 8 10:32:22 2017 -0800

    database/sql: do not exhaust connection pool on conn request timeout
    
    Previously if a context was canceled while it was waiting for a
    connection request, that connection request would leak.
    
    To prevent this remove the pending connection request if the
    context is canceled and ensure no connection has been sent on the channel.
    This requires a change to how the connection requests are represented in the DB.
    
    Fixes #18995
    
    Change-Id: I9a274b48b8f4f7ca46cdee166faa38f56d030852
    Reviewed-on: https://go-review.googlesource.com/36563
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c57d91e34cf9a9d6d39b75e2f401bdf6a27447aa
Author: Kevin Burke <kev@inburke.com>
Date:   Wed Feb 8 15:31:33 2017 -0800

    database/sql: fix typo
    
    Change-Id: I09fdcebb939417f18af09ed57f24460724cab64f
    Reviewed-on: https://go-review.googlesource.com/36632
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 27520cc4c5deaa14e46cb97a382150e0f313b099
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Feb 8 15:03:56 2017 -0800

    net: merge FreeBSD and DragonFly sendfile support
    
    The two files were identical except for comments.
    
    Change-Id: Ifc300026c8e4584afa50a7b669099eaff146ea5d
    Reviewed-on: https://go-review.googlesource.com/36631
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 12991a75e0d93443642137d63719d63c286729be
Author: David du Colombier <0intro@gmail.com>
Date:   Wed Feb 8 23:11:15 2017 +0100

    cmd/gofmt: fix diff on Plan 9
    
    On Plan 9, GNU diff is called ape/diff.
    
    Fixes #18999.
    
    Change-Id: I7cf6c23c97bcc47172bbf838fd9dd72aefa4c18b
    Reviewed-on: https://go-review.googlesource.com/36650
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit af59742d0f6abe96eb5d68151b4d2ccb45d5ed89
Author: Dmitri Shuralyov <shurcooL@gmail.com>
Date:   Tue Feb 7 10:57:33 2017 -0500

    net/http: don't modify Request in StripPrefix
    
    As of https://golang.org/cl/21530, rules are updated to state
    that Handlers shouldn't modify the provided Request. This change
    updates StripPrefix to follow that rule.
    
    Resolves #18952.
    
    Change-Id: I29bbb580722e871131fa75a97e6e038ec64fdfcd
    Reviewed-on: https://go-review.googlesource.com/36483
    Reviewed-by: Matt Layher <mdlayher@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Matt Layher <mdlayher@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a146dd3a2fdf87bab90ee1f636c38cd3444e55fa
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Feb 8 15:31:24 2017 -0500

    cmd/compile: handle DOT STRUCTLIT for zero-valued struct in SSA
    
    CL 35261 makes SSA handle zero-valued STRUCTLIT, but DOT operation
    was not handled.
    
    Fixes #18994.
    
    Change-Id: Ic7976036acca1523b0b14afac4d170797e8aee20
    Reviewed-on: https://go-review.googlesource.com/36565
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e3efdffacdd27786ecf0647272e54c664daf4c94
Author: David Lazar <lazard@golang.org>
Date:   Fri Dec 2 13:55:25 2016 -0500

    cmd/compile: include linknames in export data
    
    This lets the compiler inline functions that contain a linknamed symbol.
    Previously, the net/http tests would fail to build with -l=4 because
    the compiler inlined functions that call net.byteIndex (which is
    linknamed to strings.IndexByte).
    
    This changes only the compiler-specific export data, so we don't need to
    bump the export format version number.
    
    The following benchmark results show how the size of package export data
    is impacted by this change. These benchmarks were created by compiling
    the go1 benchmark and running `go tool pack x` to extract the export
    data from the resulting .a files.
    
    name                                          old bytes   new bytes   delta
    bufio                                        3.48k ± 0%  3.58k ± 0%  +2.90%
    bytes                                        5.05k ± 0%  5.16k ± 0%  +2.16%
    compress/bzip2                               2.61k ± 0%  2.68k ± 0%  +2.68%
    compress/flate                               5.07k ± 0%  5.14k ± 0%  +1.40%
    compress/gzip                                8.26k ± 0%  8.40k ± 0%  +1.70%
    container/list                               1.69k ± 0%  1.76k ± 0%  +4.07%
    context                                      3.93k ± 0%  4.01k ± 0%  +1.86%
    crypto                                       1.03k ± 0%  1.03k ± 0%  +0.39%
    crypto/aes                                     475 ± 0%    475 ± 0%  +0.00%
    crypto/cipher                                1.18k ± 0%  1.18k ± 0%  +0.00%
    crypto/des                                     502 ± 0%    502 ± 0%  +0.00%
    crypto/dsa                                   5.71k ± 0%  5.77k ± 0%  +1.16%
    crypto/ecdsa                                 6.67k ± 0%  6.75k ± 0%  +1.08%
    crypto/elliptic                              6.28k ± 0%  6.35k ± 0%  +1.07%
    crypto/hmac                                    464 ± 0%    464 ± 0%  +0.00%
    crypto/internal/cipherhw                       313 ± 0%    313 ± 0%  +0.00%
    crypto/md5                                     691 ± 0%    695 ± 0%  +0.58%
    crypto/rand                                  5.37k ± 0%  5.43k ± 0%  +1.23%
    crypto/rc4                                     512 ± 0%    512 ± 0%  +0.00%
    crypto/rsa                                   7.05k ± 0%  7.12k ± 0%  +1.05%
    crypto/sha1                                    756 ± 0%    760 ± 0%  +0.53%
    crypto/sha256                                  523 ± 0%    523 ± 0%  +0.00%
    crypto/sha512                                  662 ± 0%    662 ± 0%  +0.00%
    crypto/subtle                                  835 ± 0%    873 ± 0%  +4.55%
    crypto/tls                                   28.1k ± 0%  28.5k ± 0%  +1.30%
    crypto/x509                                  17.7k ± 0%  17.9k ± 0%  +1.04%
    crypto/x509/pkix                             9.75k ± 0%  9.90k ± 0%  +1.50%
    encoding                                       473 ± 0%    473 ± 0%  +0.00%
    encoding/asn1                                1.41k ± 0%  1.42k ± 0%  +1.00%
    encoding/base64                              1.67k ± 0%  1.69k ± 0%  +0.90%
    encoding/binary                              2.65k ± 0%  2.76k ± 0%  +4.07%
    encoding/gob                                 13.3k ± 0%  13.5k ± 0%  +1.65%
    encoding/hex                                   854 ± 0%    857 ± 0%  +0.35%
    encoding/json                                11.9k ± 0%  12.1k ± 0%  +1.71%
    encoding/pem                                   484 ± 0%    484 ± 0%  +0.00%
    errors                                         360 ± 0%    361 ± 0%  +0.28%
    flag                                         7.32k ± 0%  7.42k ± 0%  +1.48%
    fmt                                          1.42k ± 0%  1.42k ± 0%  +0.00%
    go/ast                                       15.7k ± 0%  15.8k ± 0%  +1.07%
    go/parser                                    7.48k ± 0%  7.59k ± 0%  +1.55%
    go/scanner                                   3.88k ± 0%  3.94k ± 0%  +1.39%
    go/token                                     3.51k ± 0%  3.53k ± 0%  +0.60%
    hash                                           507 ± 0%    507 ± 0%  +0.00%
    hash/crc32                                     685 ± 0%    685 ± 0%  +0.00%
    internal/nettrace                              474 ± 0%    474 ± 0%  +0.00%
    internal/pprof/profile                       8.29k ± 0%  8.36k ± 0%  +0.89%
    internal/race                                  511 ± 0%    511 ± 0%  +0.00%
    internal/singleflight                          966 ± 0%    969 ± 0%  +0.31%
    internal/syscall/unix                          427 ± 0%    427 ± 0%  +0.00%
    io                                           3.48k ± 0%  3.52k ± 0%  +1.15%
    io/ioutil                                    5.30k ± 0%  5.38k ± 0%  +1.53%
    log                                          4.46k ± 0%  4.53k ± 0%  +1.59%
    math                                         3.72k ± 0%  3.75k ± 0%  +0.75%
    math/big                                     8.91k ± 0%  9.01k ± 0%  +1.15%
    math/rand                                    1.29k ± 0%  1.30k ± 0%  +0.46%
    mime                                         2.59k ± 0%  2.63k ± 0%  +1.55%
    mime/multipart                               3.61k ± 0%  3.68k ± 0%  +1.80%
    mime/quotedprintable                         2.20k ± 0%  2.25k ± 0%  +2.50%
    net                                          21.1k ± 0%  21.3k ± 0%  +1.10%
    net/http                                     56.6k ± 0%  57.3k ± 0%  +1.28%
    net/http/httptest                            33.6k ± 0%  34.1k ± 0%  +1.38%
    net/http/httptrace                           14.4k ± 0%  14.5k ± 0%  +1.29%
    net/http/internal                            2.70k ± 0%  2.77k ± 0%  +2.59%
    net/textproto                                4.51k ± 0%  4.60k ± 0%  +1.82%
    net/url                                      1.71k ± 0%  1.73k ± 0%  +1.41%
    os                                           11.3k ± 0%  11.4k ± 0%  +1.36%
    path                                           587 ± 0%    589 ± 0%  +0.34%
    path/filepath                                4.46k ± 0%  4.55k ± 0%  +1.88%
    reflect                                      6.39k ± 0%  6.43k ± 0%  +0.72%
    regexp                                       5.82k ± 0%  5.88k ± 0%  +1.12%
    regexp/syntax                                3.22k ± 0%  3.24k ± 0%  +0.62%
    runtime                                      12.9k ± 0%  13.2k ± 0%  +1.94%
    runtime/cgo                                    229 ± 0%    229 ± 0%  +0.00%
    runtime/debug                                3.66k ± 0%  3.72k ± 0%  +1.86%
    runtime/internal/atomic                        905 ± 0%    905 ± 0%  +0.00%
    runtime/internal/sys                         2.00k ± 0%  2.05k ± 0%  +2.55%
    runtime/pprof                                4.16k ± 0%  4.23k ± 0%  +1.66%
    runtime/pprof/internal/protopprof            11.5k ± 0%  11.7k ± 0%  +1.27%
    runtime/trace                                  354 ± 0%    354 ± 0%  +0.00%
    sort                                         1.63k ± 0%  1.68k ± 0%  +2.94%
    strconv                                      1.84k ± 0%  1.85k ± 0%  +0.54%
    strings                                      3.87k ± 0%  3.97k ± 0%  +2.48%
    sync                                         1.51k ± 0%  1.52k ± 0%  +0.33%
    sync/atomic                                  1.58k ± 0%  1.60k ± 0%  +1.27%
    syscall                                      53.2k ± 0%  53.3k ± 0%  +0.20%
    testing                                      8.14k ± 0%  8.26k ± 0%  +1.49%
    testing/internal/testdeps                      597 ± 0%    598 ± 0%  +0.17%
    text/tabwriter                               3.09k ± 0%  3.14k ± 0%  +1.85%
    text/template                                15.4k ± 0%  15.7k ± 0%  +1.89%
    text/template/parse                          8.90k ± 0%  9.12k ± 0%  +2.46%
    time                                         5.75k ± 0%  5.86k ± 0%  +1.86%
    unicode                                      4.62k ± 0%  4.62k ± 0%  +0.07%
    unicode/utf16                                  693 ± 0%    706 ± 0%  +1.88%
    unicode/utf8                                 1.05k ± 0%  1.07k ± 0%  +1.14%
    vendor/golang_org/x/crypto/chacha20poly1305  1.25k ± 0%  1.26k ± 0%  +0.64%
    vendor/golang_org/x/crypto/curve25519          392 ± 0%    392 ± 0%  +0.00%
    vendor/golang_org/x/crypto/poly1305            426 ± 0%    426 ± 0%  +0.00%
    vendor/golang_org/x/net/http2/hpack          4.19k ± 0%  4.26k ± 0%  +1.69%
    vendor/golang_org/x/net/idna                   355 ± 0%    355 ± 0%  +0.00%
    vendor/golang_org/x/net/lex/httplex            609 ± 0%    615 ± 0%  +0.99%
    vendor/golang_org/x/text/transform           1.31k ± 0%  1.31k ± 0%  +0.08%
    vendor/golang_org/x/text/unicode/norm        5.78k ± 0%  5.90k ± 0%  +2.06%
    vendor/golang_org/x/text/width               1.24k ± 0%  1.24k ± 0%  +0.16%
    [Geo mean]                                    2.49k       2.52k       +1.10%
    
    Fixes #18167.
    
    Change-Id: Ia5b7e70adc9652c7ee9954ca2efc1c59fa79be2b
    Reviewed-on: https://go-review.googlesource.com/33911
    Run-TryBot: David Lazar <lazard@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 39366326cc2600792a9d9f8c170aae5bc1613d25
Author: Kale Blankenship <kale@lemnisys.com>
Date:   Sun Jan 22 17:16:59 2017 -0800

    net/http/pprof: return error when requested profile duration exceeds WriteTimeout
    
    Updates Profile and Trace handlers to reject requests for durations >=
    WriteTimeout.
    
    Modifies go tool pprof to print the body of the http response when
    status != 200.
    
    Fixes #18755
    
    Change-Id: I6faed21685693caf39f315f003039538114937b0
    Reviewed-on: https://go-review.googlesource.com/35564
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7bd968fbfdfd943d8bfc3f6f48b47c5fe990f9ba
Author: Максим Федосеев <max.faceless.frei@gmail.com>
Date:   Mon Jan 30 17:11:01 2017 +0500

    crypto/tls: fix link to more info about channel bindings
    
    Link in the description of TLSUnique field of ConnectionState struct
    leads to an article that is no longer available, so this commit
    replaces it with link to a copy of the very same article on another
    site.
    
    Fixes #18842.
    
    Change-Id: I8f8d298c4774dc0fbbad5042db0684bb3220aee8
    Reviewed-on: https://go-review.googlesource.com/36052
    Reviewed-by: Filippo Valsorda <hi@filippo.io>
    Reviewed-by: Adam Langley <agl@golang.org>

commit e2390ec18352f25066074331449cbcb74957068e
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Wed Feb 8 11:17:47 2017 -0800

    doc: remove the confusing use of CL
    
    CL (change list) pops out of nowhere and confuses the
    reader. Use "change" instead to be consistent with the
    rest of the document.
    
    Fixes #18989.
    
    Change-Id: I525a63a195dc6bb992c8ad0f10c2f2e1b2b952df
    Reviewed-on: https://go-review.googlesource.com/36564
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a16e631e7434b620cf74835625265ab07e973630
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Feb 7 18:32:27 2017 -0500

    cmd/compile: remove unnecessary type conversions on s390x
    
    Some rules insert MOVDreg ops to ensure that type changes are kept.
    If there is no type change (or the input is constant) then the MOVDreg
    can be omitted, allowing further optimization.
    
    Reduces the size of the .text section in the asm tool by ~33KB.
    
    Change-Id: I386883bb35b843c7b99a269cd6840dca77cf4371
    Reviewed-on: https://go-review.googlesource.com/36547
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 92cdde016ab64416188113c72e5d6b5ade87e89f
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Feb 8 10:56:30 2017 -0800

    go/constant: use new math/big.IsInt and isUint predicates
    
    Slightly cleaner and more readable code.
    
    Change-Id: I35263dbf338861b0a1bd62d59417b6a2c6a4e670
    Reviewed-on: https://go-review.googlesource.com/36562
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ee7fdc26477a0eaf5382542a7822b45966a9f844
Author: haya14busa <hayabusa1419@gmail.com>
Date:   Sat Feb 4 22:22:14 2017 +0900

    cmd/gofmt: use actual filename in gofmt -d output
    
    By using actual filename, diff output of "gofmt -d" can be used with
    other commands like "diffstat" and "patch".
    
    Example:
      $ gofmt -d path/to/file.go | diffstat
      $ gofmt -d path/to/file.go > gofmt.patch
      $ patch -u -p0 < gofmt.patch
    
    Fixes #18932
    
    Change-Id: I21ce15eb77870d72f2c14bfd5e7c21e2c77dc9ab
    Reviewed-on: https://go-review.googlesource.com/36374
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 894650277670eed065566f803c642a8f80437456
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Tue Feb 7 13:38:52 2017 +0200

    bytes, strings: optimize Split*
    
    The relevant benchmark results on linux/amd64:
    
    bytes:
    
    SplitSingleByteSeparator-4   25.7ms ± 5%   9.1ms ± 4%  -64.40%  (p=0.000 n=10+10)
    SplitMultiByteSeparator-4    13.8ms ±20%   4.3ms ± 8%  -69.23%  (p=0.000 n=10+10)
    SplitNSingleByteSeparator-4  1.88µs ± 9%  0.88µs ± 4%  -53.25%  (p=0.000 n=10+10)
    SplitNMultiByteSeparator-4   4.83µs ±10%  1.32µs ± 9%  -72.65%  (p=0.000 n=10+10)
    
    strings:
    
    name                         old time/op  new time/op  delta
    SplitSingleByteSeparator-4   21.4ms ± 8%   8.5ms ± 5%  -60.19%  (p=0.000 n=10+10)
    SplitMultiByteSeparator-4    13.2ms ± 9%   3.9ms ± 4%  -70.29%  (p=0.000 n=10+10)
    SplitNSingleByteSeparator-4  1.54µs ± 5%  0.75µs ± 7%  -51.21%  (p=0.000 n=10+10)
    SplitNMultiByteSeparator-4   3.57µs ± 8%  1.01µs ±11%  -71.76%  (p=0.000 n=10+10)
    
    Fixes #18973
    
    Change-Id: Ie4bc010c6cc389001e72eab530497c81e5b26f34
    Reviewed-on: https://go-review.googlesource.com/36510
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c026845bd2984002168e7496cb9d0150f79164d0
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Tue Feb 7 10:19:02 2017 -0800

    database/sql: record the context error in Rows if canceled
    
    Previously it was intended that Rows.Scan would return
    an error and Rows.Err would return nil. This was problematic
    because drivers could not differentiate between a normal
    Rows.Close or a context cancel close.
    
    The alternative is to require drivers to return a Scan to return
    an error if the driver is closed while there are still rows to be read.
    This is currently not how several drivers currently work and may be
    difficult to detect when there are additional rows.
    
    At the same time guard the the Rows.lasterr and prevent a close
    while a Rows operation is active.
    
    For the drivers that do not have Context methods, do not check for
    context cancelation after the operation, but before for any operation
    that may modify the database state.
    
    Fixes #18961
    
    Change-Id: I49a25318ecd9f97a35d5b50540ecd850c01cfa5e
    Reviewed-on: https://go-review.googlesource.com/36485
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0c9325e13d5d9e2ab7459522e3556f6f44cbcb10
Author: Adam Langley <agl@golang.org>
Date:   Wed Feb 8 09:47:34 2017 -0800

    crypto/tls: document that only tickets are supported.
    
    This change clarifies that only ticket-based resumption is supported by
    crypto/tls. It's not clear where to document this for a server,
    although perhaps it's obvious there because there's nowhere to plug in
    the storage that would be needed by SessionID-based resumption.
    
    Fixes #18607
    
    Change-Id: Iaaed53e8d8f2f45c2f24c0683052df4be6340922
    Reviewed-on: https://go-review.googlesource.com/36560
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 438818d9f1b5b4ffae3ca63d5ce2a2f5cef97552
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Tue Feb 7 16:00:39 2017 -0600

    bytes: use Index in Count
    
    Similar to https://go-review.googlesource.com/28586,
    but for package bytes instead of strings.
    This provides simpler code and some performance gain.
    Also update strings.Count to use the same code.
    
    On AMD64 with heavily optimized Index I see:
    
    name             old time/op    new time/op     delta
    Count/10-6         47.3ns ± 0%     36.8ns ± 0%    -22.35%  (p=0.000 n=10+10)
    Count/32-6          286ns ± 0%       38ns ± 0%    -86.71%  (p=0.000 n=10+10)
    Count/4K-6         50.1µs ± 0%      4.4µs ± 0%    -91.18%  (p=0.000 n=10+10)
    Count/4M-6         48.1ms ± 1%      4.5ms ± 0%    -90.56%  (p=0.000 n=10+9)
    Count/64M-6         784ms ± 0%       73ms ± 0%    -90.73%  (p=0.000 n=10+10)
    CountEasy/10-6     28.4ns ± 0%     31.0ns ± 0%     +9.23%  (p=0.000 n=10+10)
    CountEasy/32-6     30.6ns ± 0%     37.0ns ± 0%    +20.92%  (p=0.000 n=10+10)
    CountEasy/4K-6      186ns ± 0%      198ns ± 0%     +6.45%  (p=0.000 n=9+10)
    CountEasy/4M-6      233µs ± 2%      234µs ± 2%       ~     (p=0.912 n=10+10)
    CountEasy/64M-6    6.70ms ± 0%     6.68ms ± 1%       ~     (p=0.762 n=8+10)
    
    name             old speed      new speed       delta
    Count/10-6        211MB/s ± 0%    272MB/s ± 0%    +28.77%  (p=0.000 n=10+9)
    Count/32-6        112MB/s ± 0%    842MB/s ± 0%   +652.84%  (p=0.000 n=10+10)
    Count/4K-6       81.8MB/s ± 0%  927.6MB/s ± 0%  +1033.63%  (p=0.000 n=10+9)
    Count/4M-6       87.2MB/s ± 1%  924.0MB/s ± 0%   +959.25%  (p=0.000 n=10+9)
    Count/64M-6      85.6MB/s ± 0%  922.9MB/s ± 0%   +978.31%  (p=0.000 n=10+10)
    CountEasy/10-6    352MB/s ± 0%    322MB/s ± 0%     -8.41%  (p=0.000 n=10+10)
    CountEasy/32-6   1.05GB/s ± 0%   0.87GB/s ± 0%    -17.35%  (p=0.000 n=9+10)
    CountEasy/4K-6   22.0GB/s ± 0%   20.6GB/s ± 0%     -6.33%  (p=0.000 n=10+10)
    CountEasy/4M-6   18.0GB/s ± 2%   18.0GB/s ± 2%       ~     (p=0.912 n=10+10)
    CountEasy/64M-6  10.0GB/s ± 0%   10.0GB/s ± 1%       ~     (p=0.762 n=8+10)
    
    On 386, without asm version of Index:
    
    Count/10-6         57.0ns ± 0%     56.9ns ± 0%   -0.11%  (p=0.006 n=10+9)
    Count/32-6          340ns ± 0%      274ns ± 0%  -19.48%  (p=0.000 n=10+9)
    Count/4K-6         49.5µs ± 0%     37.1µs ± 0%  -24.96%  (p=0.000 n=10+10)
    Count/4M-6         51.1ms ± 0%     38.2ms ± 0%  -25.21%  (p=0.000 n=10+10)
    Count/64M-6         818ms ± 0%      613ms ± 0%  -25.07%  (p=0.000 n=8+10)
    CountEasy/10-6     60.0ns ± 0%     70.4ns ± 0%  +17.34%  (p=0.000 n=10+10)
    CountEasy/32-6     81.1ns ± 0%     94.0ns ± 0%  +15.97%  (p=0.000 n=9+10)
    CountEasy/4K-6     4.37µs ± 0%     4.39µs ± 0%   +0.30%  (p=0.000 n=10+9)
    CountEasy/4M-6     4.43ms ± 0%     4.43ms ± 0%     ~     (p=0.579 n=10+10)
    CountEasy/64M-6    70.9ms ± 0%     70.9ms ± 0%     ~     (p=0.912 n=10+10)
    
    name             old speed      new speed       delta
    Count/10-6        176MB/s ± 0%    176MB/s ± 0%   +0.10%  (p=0.000 n=10+9)
    Count/32-6       93.9MB/s ± 0%  116.5MB/s ± 0%  +24.06%  (p=0.000 n=10+9)
    Count/4K-6       82.7MB/s ± 0%  110.3MB/s ± 0%  +33.26%  (p=0.000 n=10+10)
    Count/4M-6       82.1MB/s ± 0%  109.7MB/s ± 0%  +33.70%  (p=0.000 n=10+10)
    Count/64M-6      82.0MB/s ± 0%  109.5MB/s ± 0%  +33.46%  (p=0.000 n=8+10)
    CountEasy/10-6    167MB/s ± 0%    142MB/s ± 0%  -14.75%  (p=0.000 n=9+10)
    CountEasy/32-6    395MB/s ± 0%    340MB/s ± 0%  -13.77%  (p=0.000 n=10+10)
    CountEasy/4K-6    936MB/s ± 0%    934MB/s ± 0%   -0.29%  (p=0.000 n=10+9)
    CountEasy/4M-6    947MB/s ± 0%    946MB/s ± 0%     ~     (p=0.591 n=10+10)
    CountEasy/64M-6   947MB/s ± 0%    947MB/s ± 0%     ~     (p=0.867 n=10+10)
    
    Change-Id: Ia76b247372b6f5b5d23a9f10253a86536a5153b3
    Reviewed-on: https://go-review.googlesource.com/36489
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 04e0a7622c6f0f55c6a8cc4f16f7877b743ac5bb
Author: Russ Cox <rsc@golang.org>
Date:   Fri Jan 27 14:14:50 2017 -0500

    hash/crc32: use sub-benchmarks
    
    Change-Id: Iae68a097a6897f1616f94fdc3548837ef200e66f
    Reviewed-on: https://go-review.googlesource.com/36541
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit bd5616991b9310b95262c2dfea9a6590187c05ee
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Feb 7 23:15:24 2017 +0000

    time: bound file reads and validate LoadLocation argument
    
    Fixes #18985
    
    Change-Id: I956117f47d1d2b453b4786c7b78c1c944defeca0
    Reviewed-on: https://go-review.googlesource.com/36551
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
```
