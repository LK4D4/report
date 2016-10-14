# October 6, 2016 Report

Number of commits: 97

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 581.576461ms to 583.279194ms, +0.29%
* github.com/gogits/gogs: from 12.579064712s to 12.606124146s, +0.22%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 33.349483375s to 32.982180009s, -1.10%
* github.com/influxdata/influxdb/cmd/influxd: from 6.67968705s to 6.451083101s, -3.42%
* github.com/junegunn/fzf/src/fzf: from 1.045339652s to 1.157864619s, +10.76%
* github.com/mholt/caddy/caddy: from 6.247933711s to 7.459865729s, +19.40%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.439472696s to 1.595065363s, +10.81%
* github.com/nsqio/nsq/apps/nsqd: from 2.060026872s to 2.275248285s, +10.45%
* github.com/prometheus/prometheus/cmd/prometheus: from 11.569651657s to 11.994541527s, +3.67%
* github.com/spf13/hugo: from 8.061223392s to 7.621509542s, -5.45%
* golang.org/x/tools/cmd/guru: from 2.606247421s to 2.550537355s, -2.14%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2573942 to 2595413, +0.83%
* github.com/gogits/gogs: from 23532319 to 23544740, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 29132896 to 29174000, +0.14%
* github.com/influxdata/influxdb/cmd/influxd: from 16436219 to 16449344, ~
* github.com/junegunn/fzf/src/fzf: from 3153024 to 3158912, +0.19%
* github.com/mholt/caddy/caddy: from 14922835 to 14940165, +0.12%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4362327 to 4384116, +0.50%
* github.com/nsqio/nsq/apps/nsqd: from 9939114 to 9956471, +0.17%
* github.com/prometheus/prometheus/cmd/prometheus: from 25401056 to 25422024, ~
* github.com/spf13/hugo: from 16269250 to 16295664, +0.16%
* golang.org/x/tools/cmd/guru: from 7493167 to 7523237, +0.40%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               183           188           +2.73%
BenchmarkMsgpUnmarshal-4             372           373           +0.27%
BenchmarkEasyJsonMarshal-4           1452          1415          -2.55%
BenchmarkEasyJsonUnmarshal-4         1567          1993          +27.19%
BenchmarkFlatBuffersMarshal-4        350           349           -0.29%
BenchmarkFlatBuffersUnmarshal-4      272           272           +0.00%
BenchmarkGogoprotobufMarshal-4       156           155           -0.64%
BenchmarkGogoprotobufUnmarshal-4     239           242           +1.26%

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

* [sort: add Slice, SliceStable, and SliceIsSorted](https://github.com/golang/go/commit/22a2bdfedb95612984cec3141924953b88a607b7)
* [all: use sort.Slice where applicable](https://github.com/golang/go/commit/ad26bb5e3098cbfd7c0ad9a1dc9d38c92e50f06e)
* [runtime: make append only clear uncopied memory](https://github.com/golang/go/commit/c1e267cc734135a66af8a1a5015e572cbb598d44)
* [encoding/csv: avoid allocations when reading records](https://github.com/golang/go/commit/bd06d4827ae637cd08f85962f996760e76e28efc)
* [encoding/json: use standard ES6 formatting for numbers during marshal](https://github.com/golang/go/commit/92b3e3651dc44f54b458f171f641779f10fbaec0)
* [runtime: improve memmove for amd64](https://github.com/golang/go/commit/d7507e9d1109da424dd375365dc923257ebd0c23)

## GIT Log

```
commit 91706c04b93e9f14321a038943829e99dc333794
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Oct 5 23:08:25 2016 -0400

    cmd/asm, cmd/internal/obj/s390x: delete unused instructions
    
    Deletes the following s390x instructions:
    
     - ADDME
     - ADDZE
     - SUBME
     - SUBZE
    
    They appear to be emulated PPC instructions left over from the
    porting process and I don't think they will ever be useful.
    
    Change-Id: I9b1ba78019dbd1218d0c8f8ea2903878802d1990
    Reviewed-on: https://go-review.googlesource.com/30538
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d7507e9d1109da424dd375365dc923257ebd0c23
Author: Denis Nagorny <denis.nagorny@intel.com>
Date:   Thu Sep 22 04:05:39 2016 +0300

        runtime: improve memmove for amd64
    
        Use AVX if available on 4th generation of Intel(TM) Core(TM) processors.
    
        (collected on E5 2609v3 @1.9GHz)
        name                        old speed      new speed       delta
        Memmove/1-6                  158MB/s ± 0%    172MB/s ± 0%    +9.09% (p=0.000 n=16+16)
        Memmove/2-6                  316MB/s ± 0%    345MB/s ± 0%    +9.09% (p=0.000 n=18+16)
        Memmove/3-6                  517MB/s ± 0%    517MB/s ± 0%      ~ (p=0.445 n=16+16)
        Memmove/4-6                  687MB/s ± 1%    690MB/s ± 0%    +0.35% (p=0.000 n=20+17)
        Memmove/5-6                  729MB/s ± 0%    729MB/s ± 0%    +0.01% (p=0.000 n=16+18)
        Memmove/6-6                  875MB/s ± 0%    875MB/s ± 0%    +0.01% (p=0.000 n=18+18)
        Memmove/7-6                 1.02GB/s ± 0%   1.02GB/s ± 1%      ~ (p=0.139 n=19+20)
        Memmove/8-6                 1.26GB/s ± 0%   1.26GB/s ± 0%    +0.00% (p=0.000 n=18+18)
        Memmove/9-6                 1.42GB/s ± 0%   1.42GB/s ± 0%    +0.00% (p=0.000 n=17+18)
        Memmove/10-6                1.58GB/s ± 0%   1.58GB/s ± 0%    +0.00% (p=0.000 n=19+19)
        Memmove/11-6                1.74GB/s ± 0%   1.74GB/s ± 0%    +0.00% (p=0.001 n=18+17)
        Memmove/12-6                1.90GB/s ± 0%   1.90GB/s ± 0%    +0.00% (p=0.000 n=19+19)
        Memmove/13-6                2.05GB/s ± 0%   2.05GB/s ± 0%    +0.00% (p=0.000 n=18+19)
        Memmove/14-6                2.21GB/s ± 0%   2.21GB/s ± 0%    +0.00% (p=0.000 n=16+20)
        Memmove/15-6                2.37GB/s ± 0%   2.37GB/s ± 0%    +0.00% (p=0.004 n=19+20)
        Memmove/16-6                2.53GB/s ± 0%   2.53GB/s ± 0%    +0.00% (p=0.000 n=16+16)
        Memmove/32-6                4.67GB/s ± 0%   4.67GB/s ± 0%    +0.00% (p=0.000 n=17+17)
        Memmove/64-6                8.67GB/s ± 0%   8.64GB/s ± 0%    -0.33% (p=0.000 n=18+17)
        Memmove/128-6               12.6GB/s ± 0%   11.6GB/s ± 0%    -8.05% (p=0.000 n=16+19)
        Memmove/256-6               16.3GB/s ± 0%   16.6GB/s ± 0%    +1.66% (p=0.000 n=20+18)
        Memmove/512-6               21.5GB/s ± 0%   24.4GB/s ± 0%   +13.35% (p=0.000 n=18+17)
        Memmove/1024-6              24.7GB/s ± 0%   33.7GB/s ± 0%   +36.12% (p=0.000 n=18+18)
        Memmove/2048-6              27.3GB/s ± 0%   43.3GB/s ± 0%   +58.77% (p=0.000 n=19+17)
        Memmove/4096-6              37.5GB/s ± 0%   50.5GB/s ± 0%   +34.56% (p=0.000 n=19+19)
        MemmoveUnalignedDst/1-6      135MB/s ± 0%    146MB/s ± 0%    +7.69% (p=0.000 n=16+14)
        MemmoveUnalignedDst/2-6      271MB/s ± 0%    292MB/s ± 0%    +7.69% (p=0.000 n=18+18)
        MemmoveUnalignedDst/3-6      438MB/s ± 0%    438MB/s ± 0%      ~ (p=0.352 n=16+19)
        MemmoveUnalignedDst/4-6      584MB/s ± 0%    584MB/s ± 0%      ~ (p=0.876 n=17+17)
        MemmoveUnalignedDst/5-6      631MB/s ± 1%    632MB/s ± 0%    +0.25% (p=0.000 n=20+17)
        MemmoveUnalignedDst/6-6      759MB/s ± 0%    759MB/s ± 0%    +0.00% (p=0.000 n=19+16)
        MemmoveUnalignedDst/7-6      885MB/s ± 0%    883MB/s ± 1%      ~ (p=0.647 n=18+20)
        MemmoveUnalignedDst/8-6     1.08GB/s ± 0%   1.08GB/s ± 0%    +0.00% (p=0.035 n=19+18)
        MemmoveUnalignedDst/9-6     1.22GB/s ± 0%   1.22GB/s ± 0%      ~ (p=0.251 n=18+17)
        MemmoveUnalignedDst/10-6    1.35GB/s ± 0%   1.35GB/s ± 0%      ~ (p=0.327 n=17+18)
        MemmoveUnalignedDst/11-6    1.49GB/s ± 0%   1.49GB/s ± 0%      ~ (p=0.531 n=18+19)
        MemmoveUnalignedDst/12-6    1.63GB/s ± 0%   1.63GB/s ± 0%      ~ (p=0.886 n=19+18)
        MemmoveUnalignedDst/13-6    1.76GB/s ± 0%   1.76GB/s ± 1%    -0.24% (p=0.006 n=18+20)
        MemmoveUnalignedDst/14-6    1.90GB/s ± 0%   1.90GB/s ± 0%      ~ (p=0.818 n=20+19)
        MemmoveUnalignedDst/15-6    2.03GB/s ± 0%   2.03GB/s ± 0%      ~ (p=0.294 n=17+16)
        MemmoveUnalignedDst/16-6    2.17GB/s ± 0%   2.17GB/s ± 0%      ~ (p=0.602 n=16+18)
        MemmoveUnalignedDst/32-6    4.05GB/s ± 0%   4.05GB/s ± 0%    +0.00% (p=0.010 n=18+17)
        MemmoveUnalignedDst/64-6    7.59GB/s ± 0%   7.59GB/s ± 0%    +0.00% (p=0.022 n=18+16)
        MemmoveUnalignedDst/128-6   11.1GB/s ± 0%   11.4GB/s ± 0%    +2.79% (p=0.000 n=18+17)
        MemmoveUnalignedDst/256-6   16.4GB/s ± 0%   16.7GB/s ± 0%    +1.59% (p=0.000 n=20+17)
        MemmoveUnalignedDst/512-6   15.7GB/s ± 0%   21.3GB/s ± 0%   +35.87% (p=0.000 n=18+20)
        MemmoveUnalignedDst/1024-6  16.0GB/s ±20%   31.5GB/s ± 0%   +96.93% (p=0.000 n=20+14)
        MemmoveUnalignedDst/2048-6  19.6GB/s ± 0%   42.1GB/s ± 0%  +115.16% (p=0.000 n=17+18)
        MemmoveUnalignedDst/4096-6  6.41GB/s ± 0%  33.18GB/s ± 0%  +417.56% (p=0.000 n=17+18)
        MemmoveUnalignedSrc/1-6      171MB/s ± 0%    166MB/s ± 0%    -3.33% (p=0.000 n=19+16)
        MemmoveUnalignedSrc/2-6      343MB/s ± 0%    342MB/s ± 1%    -0.41% (p=0.000 n=17+20)
        MemmoveUnalignedSrc/3-6      508MB/s ± 0%    493MB/s ± 1%    -2.90% (p=0.000 n=17+17)
        MemmoveUnalignedSrc/4-6      677MB/s ± 0%    660MB/s ± 2%    -2.55% (p=0.000 n=17+20)
        MemmoveUnalignedSrc/5-6      790MB/s ± 0%    790MB/s ± 0%      ~ (p=0.139 n=17+17)
        MemmoveUnalignedSrc/6-6      948MB/s ± 0%    946MB/s ± 1%      ~ (p=0.330 n=17+19)
        MemmoveUnalignedSrc/7-6     1.11GB/s ± 0%   1.11GB/s ± 0%    -0.05% (p=0.026 n=17+17)
        MemmoveUnalignedSrc/8-6     1.38GB/s ± 0%   1.38GB/s ± 0%      ~ (p=0.091 n=18+16)
        MemmoveUnalignedSrc/9-6     1.42GB/s ± 0%   1.40GB/s ± 1%    -1.04% (p=0.000 n=19+20)
        MemmoveUnalignedSrc/10-6    1.58GB/s ± 0%   1.56GB/s ± 1%    -1.15% (p=0.000 n=18+19)
        MemmoveUnalignedSrc/11-6    1.73GB/s ± 0%   1.71GB/s ± 1%    -1.30% (p=0.000 n=20+20)
        MemmoveUnalignedSrc/12-6    1.89GB/s ± 0%   1.87GB/s ± 1%    -1.18% (p=0.000 n=17+20)
        MemmoveUnalignedSrc/13-6    2.05GB/s ± 0%   2.02GB/s ± 1%    -1.18% (p=0.000 n=17+20)
        MemmoveUnalignedSrc/14-6    2.21GB/s ± 0%   2.18GB/s ± 1%    -1.14% (p=0.000 n=17+20)
        MemmoveUnalignedSrc/15-6    2.36GB/s ± 0%   2.34GB/s ± 1%    -1.04% (p=0.000 n=17+20)
        MemmoveUnalignedSrc/16-6    2.52GB/s ± 0%   2.49GB/s ± 1%    -1.26% (p=0.000 n=19+20)
        MemmoveUnalignedSrc/32-6    4.82GB/s ± 0%   4.61GB/s ± 0%    -4.40% (p=0.000 n=19+20)
        MemmoveUnalignedSrc/64-6    5.03GB/s ± 4%   7.97GB/s ± 0%   +58.55% (p=0.000 n=20+16)
        MemmoveUnalignedSrc/128-6   11.1GB/s ± 0%   11.2GB/s ± 0%    +0.52% (p=0.000 n=17+18)
        MemmoveUnalignedSrc/256-6   16.5GB/s ± 0%   16.4GB/s ± 0%    -0.10% (p=0.000 n=20+18)
        MemmoveUnalignedSrc/512-6   21.0GB/s ± 0%   22.1GB/s ± 0%    +5.48% (p=0.000 n=14+17)
        MemmoveUnalignedSrc/1024-6  24.9GB/s ± 0%   31.9GB/s ± 0%   +28.20% (p=0.000 n=19+20)
        MemmoveUnalignedSrc/2048-6  23.3GB/s ± 0%   33.8GB/s ± 0%   +45.22% (p=0.000 n=17+19)
        MemmoveUnalignedSrc/4096-6  37.3GB/s ± 0%   42.7GB/s ± 0%   +14.30% (p=0.000 n=17+17)
    
    Change-Id: Id66aa3e499ccfb117cb99d623ef326b50d057b64
    Reviewed-on: https://go-review.googlesource.com/29590
    Run-TryBot: Denis Nagorny <denis.nagorny@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit dd1dcf949676a5f091d8f17ad9a64f6336aa371b
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Oct 4 15:35:28 2016 -0400

    cmd/{asm,compile}: add ANDW, ORW and XORW instructions to s390x
    
    Adds the following instructions and uses them in the SSA backend:
    
     - ANDW
     - ORW
     - XORW
    
    The instruction encodings for 32-bit operations are typically shorter,
    particularly when an immediate is used. For example, XORW $-1, R1
    only requires one instruction, whereas XOR requires two.
    
    Also removes some unused instructions (that were emulated):
    
     - ANDN
     - NAND
     - ORN
     - NOR
    
    Change-Id: Iff2a16f52004ba498720034e354be9771b10cac4
    Reviewed-on: https://go-review.googlesource.com/30291
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 1bddd2ee6aed261830131f824fe32e07de326066
Author: Keith Randall <khr@golang.org>
Date:   Wed Oct 5 14:35:47 2016 -0700

    cmd/compile: don't shuffle rematerializeable values around
    
    Better to just rematerialize them when needed instead of
    cross-register spilling or other techniques for keeping them in
    registers.
    
    This helps for amd64 code that does 1 << x. It is better to do
      loop:
        MOVQ $1, AX  // materialize arg to SLLQ
        SLLQ CX, AX
        ...
        goto loop
    than to do
      MOVQ $1, AX    // materialize outsize of loop
      loop:
        MOVQ AX, DX  // save value that's about to be clobbered
        SLLQ CX, AX
        MOVQ DX, AX  // move it back to the correct register
        goto loop
    
    Update #16092
    
    Change-Id: If7ac290208f513061ebb0736e8a79dcb0ba338c0
    Reviewed-on: https://go-review.googlesource.com/30471
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit e5421e21effb5b1db4e565babbddffeb4103d40e
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Oct 5 07:31:11 2016 -0700

    runtime: add threadprof tag for test that starts busy thread
    
    The CgoExternalThreadSIGPROF test starts a thread at constructor time
    that does a busy loop. That can throw off some other tests. So only
    build that code if testprogcgo is built with the tag threadprof, and
    adjust the tests that use that code to pass that build tag.
    
    This revealed that the CgoPprofThread test was not testing what it
    should have, as it never actually started the cpuHog thread. It was
    passing because of the busy loop thread. Fix it to start the thread as
    intended.
    
    Change-Id: I087a9e4fc734a86be16a287456441afac5676beb
    Reviewed-on: https://go-review.googlesource.com/30362
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5fd6bb4c14b395bc413f281987225b57ae5fe67c
Author: Larz Conwell <larzconwell@gmail.com>
Date:   Sat Mar 12 02:57:24 2016 -0500

    go/doc: hide methods on locally-declared predeclared types
    
    Currently if you declare a type overwriting a predeclared type
    and export methods on it they will be exposed in godoc, even
    though the type itself is not exported. This corrects that
    by making all methods on these types hidden, since that's
    the expected output.
    
    Fixes #9860
    
    Change-Id: I14037bdcef1b4bbefcf299a143bac8bf363718e0
    Reviewed-on: https://go-review.googlesource.com/20610
    Reviewed-by: Russ Cox <rsc@golang.org>

commit eee727d0855b9e78f9df87e08d57b1d7f264876c
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Tue Sep 13 13:21:23 2016 -0700

    cmd/go: note when some Go files were ignored on no-Go-files errors
    
    It is pretty confusing when there are Go files ignored for mismatching
    build tags and similar and we output "no buildable Go files" without
    giving any other information about some Go files have been ignored.
    
    Fixes #17008.
    
    Change-Id: I1766ee86a9a7a72f6694deae3f73b47bfc9d0be5
    Reviewed-on: https://go-review.googlesource.com/29113
    Run-TryBot: Jaana Burcu Dogan <jbd@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6abc4a7c3eb7b7ed35e7d14e9afb18900367b58a
Author: Dmitriy Dudkin <dudkin.dmitriy@gmail.com>
Date:   Thu Feb 25 04:00:01 2016 +0200

    cmd/go: fix go get -u wildcard corner case
    
    Suppose you have already downloaded "foo.bar/baz", where the repo
    is for all of foo.bar/, and you then "go get -u foo.bar/...".
    The command-line wildcard expands to foo.bar/baz,
    and go get updates the foo.bar/ repo.
    Suppose that the repo update brought in foo.bar/quux,
    though, which depends on other.site/bar.
    Download does not consider foo.bar/quux, since it's
    only looking at foo.bar/baz, so it didn't download other.site/bar.
    After the download, we call importPaths(args) to decide what to install.
    That call was reevaluating the original wildcard with the new repo
    and matching foo.bar/quux, which was missing its dependency
    other.site/bar, causing a build failure.
    
    The fix in this CL is to remember the pre-download expansion
    of the argument list and pass it to the installer. Then only the things
    we tried to download get installed.
    
    The case where foo.bar/ is not even checked out yet continues to work,
    because in that case we leave the wildcard in place, and download
    reevaluates it during the download.
    
    The fix in this CL may not be the right long-term fix, but it is at least a fix.
    It may be that download should be passed all the original wildcards
    so that it can reexpand them as new code is downloaded, ideally reaching
    a fixed point. That can be left for another day.
    
    In short:
    
    - The problem is that the "install" half of "go get" was trying to install
      more than the "download" half was properly downloading.
    - The fix in this CL is to install just what was downloaded (install less).
    - It may be that a future CL should instead download what will be installed (download more).
    
    Fixes #14450.
    
    Change-Id: Ia1984761d24439549b7cff322bc0dbc262c1a653
    Reviewed-on: https://go-review.googlesource.com/19892
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3107c91e2d390771888df6b47fd6f8fc7a364cd3
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Wed Oct 5 15:12:05 2016 -0500

    runtime: memclr perf improvements on ppc64x
    
    This updates runtime/memclr_ppc64x.s to improve performance,
    by unrolling loops for larger clears.
    
    Fixes #17348
    
    benchmark                    old MB/s     new MB/s     speedup
    BenchmarkMemclr/5-80         199.71       406.63       2.04x
    BenchmarkMemclr/16-80        693.66       1817.41      2.62x
    BenchmarkMemclr/64-80        2309.35      5793.34      2.51x
    BenchmarkMemclr/256-80       5428.18      14765.81     2.72x
    BenchmarkMemclr/4096-80      8611.65      27191.94     3.16x
    BenchmarkMemclr/65536-80     8736.69      28604.23     3.27x
    BenchmarkMemclr/1M-80        9304.94      27600.09     2.97x
    BenchmarkMemclr/4M-80        8705.66      27589.64     3.17x
    BenchmarkMemclr/8M-80        8575.74      23631.04     2.76x
    BenchmarkMemclr/16M-80       8443.10      19240.68     2.28x
    BenchmarkMemclr/64M-80       8390.40      9493.04      1.13x
    BenchmarkGoMemclr/5-80       263.05       630.37       2.40x
    BenchmarkGoMemclr/16-80      904.33       1148.49      1.27x
    BenchmarkGoMemclr/64-80      2830.20      8756.70      3.09x
    BenchmarkGoMemclr/256-80     6064.59      20299.46     3.35x
    
    Change-Id: Ic76c9183c8b4129ba3df512ca8b0fe6bd424e088
    Reviewed-on: https://go-review.googlesource.com/30373
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: David Chase <drchase@google.com>

commit ce645534e481368b4ec6141d48eb361422e423b8
Author: Quentin Smith <quentin@golang.org>
Date:   Wed Oct 5 18:04:54 2016 -0400

    crypto/x509: support RHEL 7 cert bundle
    
    RHEL 7 introduces a new tool, update-ca-trust(8), which places the
    certificate bundle in a new location. Add this path to the list of
    locations that are searched for the certificate bundle.
    
    Fixes #15749
    
    Change-Id: Idc97f885ee48ef085f1eb4dacbd1c2cf55f94ff5
    Reviewed-on: https://go-review.googlesource.com/30375
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 936ae27b9c89ffa7ffe6ebb33376bc55b82e5ccd
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Oct 5 13:53:39 2016 -0700

    cmd/compile: untyped arrays bounds representable as integers are valid
    
    Fixes #13485.
    
    Change-Id: I11dd15c7d14fc19d42a3b48427a4cc1208b18e6a
    Reviewed-on: https://go-review.googlesource.com/30456
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f6b4c88941f88380caf59a050c1e9805664aa2fa
Author: Joe Tsai <joetsai@google.com>
Date:   Wed Oct 5 21:23:56 2016 +0000

    Revert "net/http: improve performance for parsePostForm"
    
    This reverts commit 59320c396e6448132a52cb5a5d96491eee1e0ad8.
    
    Reasons:
    This CL was causing failures on a large regression test that we run
    within Google. The issues arises from two bugs in the CL:
    * The CL dropped support for ';' as a delimiter (see https://golang.org/issue/2210)
    * The handling of an empty string caused an empty record to be added when
    no record was added (see https://golang.org/cl/30454 for my attempted fix)
    
    The logic being added is essentially a variation of url.ParseQuery,
    but altered to accept an io.Reader instead of a string.
    Since it is duplicated (but modified) logic, there needs to be good
    tests to ensure that it's implementation doesn't drift in functionality
    from url.ParseQuery. Fixing the above issues and adding the associated
    regression tests leads to >100 lines of codes.
    For a 4% reduction in CPU time, I think this complexity and duplicated
    logic is not worth the effort.
    
    As such, I am abandoning my efforts to fix the existing issues and
    believe that reverting CL/20301 is the better course of action.
    
    Updates #14655
    
    Change-Id: Ibb5be0a5b48a16c46337e213b79467fcafee69df
    Reviewed-on: https://go-review.googlesource.com/30470
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a9b49537771c05b82923a256a47b73af98c3e87e
Author: Alexander Döring <email@alexd.ch>
Date:   Sun Oct 2 21:07:40 2016 +0200

    os/exec: add example for CommandContext
    
    Updates #16360
    
    Change-Id: I0e0afe7a89f2ebcb3e5bbc345f77a605d3afc398
    Reviewed-on: https://go-review.googlesource.com/30103
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b9fd510cd00b6aa26e2ea7001a07b90ebf97d2ed
Author: Jirka Daněk <dnk@mail.muni.cz>
Date:   Mon Jan 18 16:26:05 2016 +0100

    encoding/json: add struct and field name to UnmarshalTypeError message
    
    The UnmarshalTypeError has two new fields Struct and Field,
    used when constructing the error message.
    
    Fixes #6716.
    
    Change-Id: I67da171480a9491960b3ae81893770644180f848
    Reviewed-on: https://go-review.googlesource.com/18692
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cb986def671cd244e4682a476cff51c4cff2d8f8
Author: Quentin Smith <quentin@golang.org>
Date:   Wed Oct 5 14:37:25 2016 -0400

    syscall: relax TestUnshare
    
    Fixes #17224.
    
    Some systems have more than just "lo" in a fresh network namespace, due
    to IPv6. Instead of testing for exactly 3 lines of output (implying 1
    interface), just test to make sure that the unshare call resulted in
    fewer interfaces than before. This should still verify that unshare did
    something.
    
    Change-Id: Iaf84c2b0e673fc207059d62e2f4dd7583a753419
    Reviewed-on: https://go-review.googlesource.com/30372
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Jessica Frazelle <me@jessfraz.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 92b3e3651dc44f54b458f171f641779f10fbaec0
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 5 11:26:04 2016 -0400

    encoding/json: use standard ES6 formatting for numbers during marshal
    
    Change float32/float64 formatting to use non-exponential form
    for a slightly wider range, to more closely match ES6 JSON.stringify
    and other JSON generators.
    
    Most notably:
    
            1e20 now formats as 100000000000000000000 (previously 1e+20)
            1e-6 now formats as 0.000001 (previously 1e-06)
            1e-7 now formats as 1e-7 (previously 1e-07)
    
    This also brings the int64 and float64 formatting in line with each other,
    for all shared representable values. For example both int64(1234567)
    and float64(1234567) now format as "1234567", where before the
    float64 formatted as "1.234567e+06".
    
    The only variation now compared to ES6 JSON.stringify is that
    Go continues to encode negative zero as "-0", not "0", so that
    the value continues to be preserved during JSON round trips.
    
    Fixes #6384.
    Fixes #14135.
    
    Change-Id: Ib0e0e009cd9181d75edc0424a28fe776bcc5bbf8
    Reviewed-on: https://go-review.googlesource.com/30371
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b662e524e4d393f7c99fe281f8c95f7b2f7015a8
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Sep 28 11:20:58 2016 -0400

    cmd/compile: use CBZ/CBNZ instrunctions on ARM64
    
    These are conditional branches that takes a register instead of
    flags as control value.
    
    Reduce binary size by 0.7%, text size by 2.4% (cmd/go as an
    exmaple).
    
    Change-Id: I0020cfde745f9eab680b8b949ad28c87fe183afd
    Reviewed-on: https://go-review.googlesource.com/30030
    Reviewed-by: David Chase <drchase@google.com>

commit 4c9a372946347304094cbf5306cce6336d11e64b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Oct 2 17:10:13 2016 -0400

    runtime, cmd/internal/obj: get rid of rewindmorestack
    
    In the function prologue, we emit a jump to the beginning of
    the function immediately after calling morestack. And in the
    runtime stack growing code, it decodes and emulates that jump.
    This emulation was necessary before we had per-PC SP deltas,
    since the traceback code assumed that the frame size was fixed
    for the whole function, except on the first instruction where
    it was 0. Since we now have per-PC SP deltas and PCDATA, we
    can correctly record that the frame size is 0. This makes the
    emulation unnecessary.
    
    This may be helpful for registerized calling convention, where
    there may be unspills of arguments after calling morestack. It
    also simplifies the runtime.
    
    Change-Id: I7ebee31eaee81795445b33f521ab6a79624c4ceb
    Reviewed-on: https://go-review.googlesource.com/30138
    Reviewed-by: David Chase <drchase@google.com>

commit 56b746974cb8dcd44b09c3db384e8aeaae8a9d3e
Author: Yasuhiro Matsumoto <mattn.jp@gmail.com>
Date:   Thu Oct 6 02:19:01 2016 +0900

    cmd/asm: close file before remove
    
    Windows doesn't remove an open file.
    
    Fixes #17345
    
    Change-Id: I283930c7d6eb3bc09ad208191afefe989804ce32
    Reviewed-on: https://go-review.googlesource.com/30430
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6dd38ebae11dc0dfa607723de3c12355e556c6ed
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Tue Oct 4 14:14:48 2016 -0500

    cmd/compile: Improve const shifts in PPC64.rules
    
    This change updates PPC64.rules to recognize constant shift
    counts and generate more efficient code sequences in those cases.
    
    Fixes #17336
    
    Change-Id: I8a7b812408d7a68388df41e42bad045dd214be17
    Reviewed-on: https://go-review.googlesource.com/30310
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit f54c0db859867f415a0702c8ceaa30c1a8976b0b
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Sun Feb 21 02:28:37 2016 -0800

    cmd/compile, cmd/cgo: align complex{64,128} like GCC
    
    complex64 and complex128 are treated like [2]float32 and [2]float64,
    so it makes sense to align them the same way.
    
    Change-Id: Ic614bcdcc91b080aeb1ad1fed6fc15ba5a2971f8
    Reviewed-on: https://go-review.googlesource.com/19800
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 30088ac9a391d8505a3e016f36aaa23170109f6f
Author: Keith Randall <khr@golang.org>
Date:   Tue Oct 4 14:35:45 2016 -0700

    cmd/compile: make CSE faster
    
    To refine a set of possibly equivalent values, the old CSE algorithm
    picked one value, compared it against all the others, and made two sets
    out of the results (the values that match the picked value and the
    values that didn't).  Unfortunately, this leads to O(n^2) behavior. The
    picked value ends up being equal to no other values, we make size 1 and
    size n-1 sets, and then recurse on the size n-1 set.
    
    Instead, sort the set by the equivalence classes of its arguments.  Then
    we just look for spots in the sorted list where the equivalence classes
    of the arguments change.  This lets us do a multi-way split for O(n lg
    n) time.
    
    This change makes cmpDepth unnecessary.
    
    The refinement portion used to call the type comparator.  That is
    unnecessary as the type was already part of the initial partition.
    
    Lowers time of 16361 from 8 sec to 3 sec.
    Lowers time of 15112 from 282 sec to 20 sec. That's kind of unfair, as
    CL 30257 changed it from 21 sec to 282 sec. But that CL fixed other bad
    compile times (issue #17127) by large factors, so net still a big win.
    
    Fixes #15112
    Fixes #16361
    
    Change-Id: I351ce111bae446608968c6d48710eeb6a3d8e527
    Reviewed-on: https://go-review.googlesource.com/30354
    Reviewed-by: Todd Neal <todd@tneal.org>

commit bd06d4827ae637cd08f85962f996760e76e28efc
Author: Justin Nuß <nuss.justin@gmail.com>
Date:   Sun Jul 3 17:49:29 2016 +0200

    encoding/csv: avoid allocations when reading records
    
    This commit changes parseRecord to allocate a single string per record,
    instead of per field, by using indexes into the raw record.
    
    Benchstat (done with f69991c17)
    
    name                          old time/op    new time/op    delta
    Read-8                          3.17µs ± 0%    2.78µs ± 1%  -12.35%  (p=0.016 n=4+5)
    ReadWithFieldsPerRecord-8       3.18µs ± 1%    2.79µs ± 1%  -12.23%  (p=0.008 n=5+5)
    ReadWithoutFieldsPerRecord-8    4.59µs ± 0%    2.77µs ± 0%  -39.58%  (p=0.016 n=4+5)
    ReadLargeFields-8               57.0µs ± 0%    55.7µs ± 0%   -2.18%  (p=0.008 n=5+5)
    
    name                          old alloc/op   new alloc/op   delta
    Read-8                            660B ± 0%      664B ± 0%   +0.61%  (p=0.008 n=5+5)
    ReadWithFieldsPerRecord-8         660B ± 0%      664B ± 0%   +0.61%  (p=0.008 n=5+5)
    ReadWithoutFieldsPerRecord-8    1.14kB ± 0%    0.66kB ± 0%  -41.75%  (p=0.008 n=5+5)
    ReadLargeFields-8               3.86kB ± 0%    3.94kB ± 0%   +1.86%  (p=0.008 n=5+5)
    
    name                          old allocs/op  new allocs/op  delta
    Read-8                            30.0 ± 0%      18.0 ± 0%  -40.00%  (p=0.008 n=5+5)
    ReadWithFieldsPerRecord-8         30.0 ± 0%      18.0 ± 0%  -40.00%  (p=0.008 n=5+5)
    ReadWithoutFieldsPerRecord-8      50.0 ± 0%      18.0 ± 0%  -64.00%  (p=0.008 n=5+5)
    ReadLargeFields-8                 66.0 ± 0%      24.0 ± 0%  -63.64%  (p=0.008 n=5+5)
    
    For a simple application that I wrote, which reads in a CSV file (via
    ReadAll) and outputs the number of rows read (15857625 rows), this change
    reduces the total time on my notebook from ~58 seconds to ~48 seconds.
    
    This reduces time and allocations (bytes) each by ~6% for a real world
    CSV file at work (~230000 rows, 13 colums).
    
    Updates #16791
    
    Change-Id: Ia07177c94624e55cdd3064a7d2751fb69322d3e4
    Reviewed-on: https://go-review.googlesource.com/24723
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit dce0df29dd9052c0f00ce8217b9a51a84206e892
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 4 15:57:24 2016 -0700

    cmd/compile/internal/gc: change bvfoo functions into bvec methods
    
    plive.go (except for printeffects and livenessprintblock) and
    reflect.go changes were prepared mechanically with gofmt -r.
    
    Passes toolstash.
    
    name       old alloc/op    new alloc/op    delta
    Template      44.3MB ± 0%     44.3MB ± 0%    ~           (p=0.367 n=30+30)
    Unicode       37.4MB ± 0%     37.4MB ± 0%    ~           (p=0.665 n=30+30)
    GoTypes        125MB ± 0%      125MB ± 0%    ~           (p=0.067 n=30+30)
    Compiler       515MB ± 0%      515MB ± 0%    ~           (p=0.542 n=30+28)
    
    name       old allocs/op   new allocs/op   delta
    Template        434k ± 0%       434k ± 0%    ~           (p=0.076 n=30+29)
    Unicode         367k ± 0%       367k ± 0%    ~           (p=0.716 n=29+30)
    GoTypes        1.24M ± 0%      1.24M ± 0%    ~           (p=0.428 n=29+29)
    Compiler       4.47M ± 0%      4.47M ± 0%    ~           (p=0.225 n=28+30)
    
    Change-Id: Ibaf0668567b3f69fba06aa03b7997c8fb152113a
    Reviewed-on: https://go-review.googlesource.com/30356
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7478ea5dba7ed02ddffd91c1d17ec8141f7cf184
Author: Michael Fraenkel <michael.fraenkel@gmail.com>
Date:   Wed Oct 5 11:27:34 2016 -0400

    net/http: multipart ReadForm close file after copy
    
    Always close the file regardless of whether the copy succeeds or fails.
    Pass along the close error if the copy succeeds
    
    Fixes #16296
    
    Change-Id: Ib394655b91d25750f029f17b3846d985f673fb50
    Reviewed-on: https://go-review.googlesource.com/30410
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f69991c17d9dea88e927643e4b7fdc43ad789ac3
Author: Russ Cox <rsc@golang.org>
Date:   Tue Oct 4 23:58:42 2016 -0400

    context: make DeadlineExceeded implement net.Error
    
    It already implemented the Timeout method,
    but implementing the full net.Error is more convenient.
    
    Fixes #14238 (again).
    
    Change-Id: Ia87f897f0f35bcb49865e2355964049227951ca6
    Reviewed-on: https://go-review.googlesource.com/30370
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fb4f4f4e96058165c0e7be32aa9ce493515c22a3
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Sep 30 14:20:42 2016 -0400

    crypto/{aes,cipher}: add optimized implementation of AES-GCM for s390x
    
    Also adds two tests: one to exercise the counter incrementing code
    and one which checks the output of the optimized implementation
    against that of the generic implementation for large/unaligned data
    sizes.
    
    Uses the KIMD instruction for GHASH and the KMCTR instruction for AES
    in counter mode.
    
    AESGCMSeal1K  75.0MB/s ± 2%  1008.7MB/s ± 1%  +1245.71%  (p=0.000 n=10+10)
    AESGCMOpen1K  75.3MB/s ± 1%  1006.0MB/s ± 1%  +1235.59%   (p=0.000 n=10+9)
    AESGCMSeal8K  78.5MB/s ± 1%  1748.4MB/s ± 1%  +2127.34%   (p=0.000 n=9+10)
    AESGCMOpen8K  78.5MB/s ± 0%  1752.7MB/s ± 0%  +2134.07%   (p=0.000 n=10+9)
    
    Change-Id: I88dbcfcb5988104bfd290ae15a60a2721c1338be
    Reviewed-on: https://go-review.googlesource.com/30361
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f15f1ff46f94c70b55903963a32dfccdcbe1efe5
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Oct 5 11:24:07 2016 -0400

    runtime/testdata/testprogcgo: add explicit return value to signalThread
    
    Should fix the clang builder.
    
    Change-Id: I3ee34581b6a7ec902420de72a8a08a2426997782
    Reviewed-on: https://go-review.googlesource.com/30363
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit aad29eba296df2374e5f7d334d33649d01552c01
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Oct 5 04:35:59 2016 +0000

    sort: fix a slice benchmark not using the stable variant, add another
    
    Change-Id: I9783d8023d453a72c4605a308064bef98168bcb8
    Reviewed-on: https://go-review.googlesource.com/30360
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ee8ec42929541055a9e063b50f9ffd5ee9404517
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Sat Sep 10 17:04:41 2016 +0300

    cmd/vet: skip printf check for non-constant format string during failed import
    
    Fixes #17006
    
    Change-Id: I3c2060ca5384a4b9782a7d804305d2cf4388dd5a
    Reviewed-on: https://go-review.googlesource.com/29014
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 6c13a1db2ebe146fec7cc7261146ca0e8420f011
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Oct 3 16:58:34 2016 -0700

    runtime: don't call cgocallback from signal handler
    
    Calling cgocallback from a signal handler can fail when using the race
    detector. Calling cgocallback will lead to a call to newextram which
    will call oneNewExtraM which will call racegostart. The racegostart
    function will set up some race detector data structures, and doing that
    will sometimes call the C memory allocator. If we are running the signal
    handler from a signal that interrupted the C memory allocator, we will
    crash or hang.
    
    Instead, change the signal handler code to call needm and dropm. The
    needm function will grab allocated m and g structures and initialize the
    g to use the current stack--the signal stack. That is all we need to
    safely call code that allocates memory and checks whether it needs to
    split the stack. This may temporarily leave us with no m available to
    run a cgo callback, but that is OK in this case since the code we call
    will quickly either crash or call dropm to return the m.
    
    Implementing this required changing some of the setSignalstackSP
    functions to avoid a write barrier. These functions never need a write
    barrier but in some cases generated one anyhow because on some systems
    the ss_sp field is a pointer.
    
    Change-Id: I3893f47c3a66278f85eab7f94c1ab11d4f3be133
    Reviewed-on: https://go-review.googlesource.com/30218
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 7faf70239670c3c1f8b4b530aba8847a03860f2a
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Oct 4 21:15:42 2016 -0700

    runtime: avoid endless loop if printing the panic value panics
    
    Change-Id: I56de359a5ccdc0a10925cd372fa86534353c6ca0
    Reviewed-on: https://go-review.googlesource.com/30358
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit efaa36017e34b87a5731793594b42d483209d808
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Oct 4 23:34:01 2016 +0000

    encoding/csv: update and add CSV reading benchmarks
    
    Benchmarks broken off from https://golang.org/cl/24723 and modified to
    allocate less in the places we're not trying to measure.
    
    Updates #16791
    
    Change-Id: I508e4cfeac60322d56f1d71ff1912f6a6f183a63
    Reviewed-on: https://go-review.googlesource.com/30357
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d1d798dd15e7abbeab335804bd6dfaec6c016018
Author: Jeff R. Allen <jra@nella.org>
Date:   Sat Dec 5 21:06:05 2015 +0600

    image/gif: check handling of truncated GIF files
    
    All the prefixes of the testGIF produce errors today,
    but they differ wildly in which errors: some are io.EOF,
    others are io.ErrUnexpectedEOF, and others are gif-specific.
    Make them all gif-specific to explain context, and make
    any complaining about EOF be sure to mention the EOF
    is unexpected.
    
    Fixes #11390.
    
    Change-Id: I742c39c88591649276268327ea314e68d1de1845
    Reviewed-on: https://go-review.googlesource.com/17493
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a39920fdbbeca653a8b2ac52678378d260e2d396
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 3 22:13:57 2016 -0400

    math: fix Gamma(x) for x < -170.5 and other corner cases
    
    Fixes #11441.
    
    Test tables generated by
    
            package main
    
            import (
                    "bytes"
                    "fmt"
                    "log"
                    "os/exec"
                    "strconv"
                    "strings"
            )
    
            var inputs = []float64{
                    0.5,
                    1.5,
                    2.5,
                    3.5,
                    -0.5,
                    -1.5,
                    -2.5,
                    -3.5,
                    0.1,
                    0.01,
                    1e-8,
                    1e-16,
                    1e-3,
                    1e-16,
                    1e-308,
                    5.6e-309,
                    5.5e-309,
                    1e-309,
                    1e-323,
                    5e-324,
                    -0.1,
                    -0.01,
                    -1e-8,
                    -1e-16,
                    -1e-3,
                    -1e-16,
                    -1e-308,
                    -5.6e-309,
                    -5.5e-309,
                    -1e-300 / 1e9,
                    -1e-300 / 1e23,
                    -5e-300 / 1e24,
                    -0.9999999999999999,
                    -1.0000000000000002,
                    -1.9999999999999998,
                    -2.0000000000000004,
                    -100.00000000000001,
                    -99.999999999999986,
                    17,
                    171,
                    171.6,
                    171.624,
                    171.625,
                    172,
                    2000,
                    -100.5,
                    -160.5,
                    -170.5,
                    -171.5,
                    -176.5,
                    -177.5,
                    -178.5,
                    -179.5,
                    -201.0001,
                    -202.9999,
                    -1000.5,
                    -1000000000.3,
                    -4503599627370495.5,
                    -63.349078729022985,
                    -127.45117632943295,
            }
    
            func main() {
                    var buf bytes.Buffer
                    for _, v := range inputs {
                            fmt.Fprintf(&buf, "gamma(%.1000g)\n", v)
                    }
                    cmd := exec.Command("gp", "-q")
                    cmd.Stdin = &buf
                    out, err := cmd.CombinedOutput()
                    if err != nil {
                            log.Fatalf("gp: %v", err)
                    }
                    f := strings.Split(string(out), "\n")
                    if len(f) > 0 && f[len(f)-1] == "" {
                            f = f[:len(f)-1]
                    }
                    if len(f) != len(inputs) {
                            log.Fatalf("gp: wrong output count\n%s\n", out)
                    }
                    for i, g := range f {
                            gf, err := strconv.ParseFloat(strings.Replace(g, " E", "e", -1), 64)
                            if err != nil {
                                    if strings.Contains(err.Error(), "value out of range") {
                                            if strings.HasPrefix(g, "-") {
                                                    fmt.Printf("\t{%g, Inf(-1)},\n", inputs[i])
                                            } else {
                                                    fmt.Printf("\t{%g, Inf(1)},\n", inputs[i])
                                            }
                                            continue
                                    }
                                    log.Fatal(err)
                            }
                            if gf == 0 && strings.HasPrefix(g, "-") {
                                    fmt.Printf("\t{%g, Copysign(0, -1)},\n", inputs[i])
                                    continue
                            }
                            fmt.Printf("\t{%g, %g},\n", inputs[i], gf)
                    }
            }
    
    Change-Id: Ie98c7751d92b8ffb40e8313f5ea10df0890e2feb
    Reviewed-on: https://go-review.googlesource.com/30146
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit aab849e4297d60918602355a0335cc7b3ca4c5f2
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 3 22:33:49 2016 -0400

    math: use portable Exp instead of 387 instructions on 386
    
    The 387 implementation is less accurate and slower.
    
    name     old time/op  new time/op  delta
    Exp-8    29.7ns ± 2%  24.0ns ± 2%  -19.08%  (p=0.000 n=10+10)
    
    This makes Gamma more accurate too.
    
    Change-Id: Iad33b9cce0b087ccbce3e08ba7a6d285c4999d02
    Reviewed-on: https://go-review.googlesource.com/30230
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 84743c348b0a4a7ed1ea3d7122feb757ccc7ebae
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Wed Aug 3 00:31:17 2016 -0700

    cmd/doc: ensure summaries truly are only one line
    
    The documentation for doc says:
    > Doc prints the documentation comments associated with the item identified by its
    > arguments (a package, const, func, type, var, or method) followed by a one-line
    > summary of each of the first-level items "under" that item (package-level
    > declarations for a package, methods for a type, etc.).
    
    Certain variables (and constants, functions, and types) have value specifications
    that are multiple lines long. Prior to this change, doc would print out all of the
    lines necessary to display the value. This is inconsistent with the documented
    behavior, which guarantees a one-line summary for all first-level items.
    We fix this here by writing a general oneLineNode method that always returns
    a one-line summary (guaranteed!) of any input node.
    
    Packages like image/color/palette and unicode now become much
    more readable since large slices are now a single line.
    
    $ go doc image/color/palette
    <<<
    // Before:
    var Plan9 = []color.Color{
            color.RGBA{0x00, 0x00, 0x00, 0xff},
            color.RGBA{0x00, 0x00, 0x44, 0xff},
            color.RGBA{0x00, 0x00, 0x88, 0xff},
            ... // Hundreds of more lines!
    }
    var WebSafe = []color.Color{
            color.RGBA{0x00, 0x00, 0x00, 0xff},
            color.RGBA{0x00, 0x00, 0x33, 0xff},
            color.RGBA{0x00, 0x00, 0x66, 0xff},
            ... // Hundreds of more lines!
    }
    
    // After:
    var Plan9 = []color.Color{ ... }
    var WebSafe = []color.Color{ ... }
    >>>
    
    In order to test this, I ran `go doc` and `go doc -u` on all of the
    standard library packages and diff'd the output with and without the
    change to ensure that all differences were intended.
    
    Fixes #13072
    
    Change-Id: Ida10b7796b7e4e174a929b55c60813a9eb7158fe
    Reviewed-on: https://go-review.googlesource.com/25420
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 44009a24138b6c7b5fd5b4be113db44fdfd1678e
Author: Kevin Burke <kev@inburke.com>
Date:   Tue Oct 4 15:38:46 2016 -0700

    cmd/cover: fix typo
    
    Change-Id: I3f13488605ab62eba5d3c59d5e9df1bcf69dd571
    Reviewed-on: https://go-review.googlesource.com/30355
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 316f93f7164d015ce82f341bd58657cc84f2cc69
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Oct 3 13:32:11 2016 -0700

    go/types: minimal support for alias declarations: don't crash
    
    For #16339
    
    Change-Id: I8927f40e0fd166795f41c784ad92449743f73af5
    Reviewed-on: https://go-review.googlesource.com/30213
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 59c63c711c73f3872c3047c2e80debba5ff1b802
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Oct 3 13:31:07 2016 -0700

    go/printer: support for printing alias declarations
    
    For #16339.
    
    Change-Id: Ie2e3338b87e84f45cda0868213bbcd2dae9ab6e3
    Reviewed-on: https://go-review.googlesource.com/30212
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 57ae83307fc4cb90338b39dcc6fe3c61ee8ce0b7
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Oct 3 13:28:25 2016 -0700

    go/ast, go/parser: parse alias declarations
    
    For now, we also accept "type p = p.T" (using = instead of =>, for
    type aliases only), so we can experiment with an approach that only
    uses type aliases. This concession is implemened in the parser.
    
    For #16339
    
    Change-Id: I88b5522a8b6cfc2e97ca146ede8b32af340220f8
    Reviewed-on: https://go-review.googlesource.com/30211
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 776a90100f1f65fcf54dfd3d082d657341bdc323
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Oct 3 13:24:56 2016 -0700

    go/scanner, go/token: recognize => (ALIAS) token
    
    For #16339.
    
    Change-Id: I0f83e46f13b5c8801aacf48fc8b690049edbbbff
    Reviewed-on: https://go-review.googlesource.com/30210
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c1e267cc734135a66af8a1a5015e572cbb598d44
Author: Carl Mastrangelo <notcarl@google.com>
Date:   Mon Oct 3 12:45:12 2016 -0700

    runtime: make append only clear uncopied memory
    
    Also add a benchmark that shows off the new behavior.  The
    existing benchmarks reuse the same slice, and thus don't ever have
    to clear memory.  Running the Append|Grow benchmarks in runtime:
    
    name                              old time/op  new time/op  delta
    AppendSliceLarge/1024Bytes-12      265ns ± 1%   265ns ± 3%     ~     (p=0.524 n=17+20)
    AppendSliceLarge/4096Bytes-12      807ns ± 3%   772ns ± 1%   -4.38%  (p=0.000 n=20+20)
    AppendSliceLarge/16384Bytes-12    3.20µs ± 4%  2.82µs ± 4%  -11.93%  (p=0.000 n=19+20)
    AppendSliceLarge/65536Bytes-12    13.0µs ± 4%  11.0µs ± 3%  -15.22%  (p=0.000 n=20+20)
    AppendSliceLarge/262144Bytes-12   62.7µs ± 1%  51.6µs ± 1%  -17.67%  (p=0.000 n=19+20)
    AppendSliceLarge/1048576Bytes-12   337µs ± 3%   289µs ± 3%  -14.36%  (p=0.000 n=20+20)
    GrowSliceBytes-12                 31.2ns ± 4%  31.4ns ±11%     ~     (p=0.308 n=19+18)
    GrowSliceInts-12                  53.4ns ±14%  45.0ns ± 6%  -15.74%  (p=0.000 n=20+19)
    GrowSlicePtr-12                   87.0ns ± 3%  83.3ns ± 3%   -4.26%  (p=0.000 n=18+17)
    GrowSliceStruct24Bytes-12         88.9ns ± 5%  77.8ns ± 2%  -12.45%  (p=0.000 n=20+19)
    Append-12                         17.2ns ± 1%  17.3ns ± 2%     ~     (p=0.464 n=18+17)
    AppendGrowByte-12                 2.28ms ± 1%  1.92ms ± 2%  -15.65%  (p=0.000 n=20+18)
    AppendGrowString-12                255ms ± 3%   253ms ± 4%     ~     (p=0.065 n=19+19)
    AppendSlice/1Bytes-12             3.13ns ± 0%  3.11ns ± 1%   -0.65%  (p=0.000 n=17+18)
    AppendSlice/4Bytes-12             3.02ns ± 2%  3.11ns ± 1%   +3.27%  (p=0.000 n=18+17)
    AppendSlice/7Bytes-12             4.14ns ± 3%  4.13ns ± 2%     ~     (p=0.380 n=19+18)
    AppendSlice/8Bytes-12             3.74ns ± 3%  3.68ns ± 1%   -1.76%  (p=0.000 n=19+18)
    AppendSlice/15Bytes-12            4.03ns ± 2%  4.04ns ± 2%     ~     (p=0.261 n=19+20)
    AppendSlice/16Bytes-12            4.03ns ± 2%  4.03ns ± 0%     ~     (p=0.062 n=18+17)
    AppendSlice/32Bytes-12            3.23ns ± 4%  3.43ns ± 1%   +6.10%  (p=0.000 n=17+18)
    AppendStr/1Bytes-12               3.51ns ± 1%  3.52ns ± 1%     ~     (p=0.321 n=18+19)
    AppendStr/4Bytes-12               3.46ns ± 1%  3.46ns ± 1%     ~     (p=0.977 n=18+20)
    AppendStr/8Bytes-12               3.18ns ± 1%  3.19ns ± 1%     ~     (p=0.650 n=16+17)
    AppendStr/16Bytes-12              6.08ns ±27%  5.52ns ± 3%   -9.16%  (p=0.002 n=18+19)
    AppendStr/32Bytes-12              3.71ns ± 1%  3.53ns ± 1%   -4.73%  (p=0.000 n=20+19)
    AppendSpecialCase-12              17.7ns ± 1%  17.8ns ± 3%   +0.86%  (p=0.045 n=17+18)
    AppendInPlace/NoGrow/Byte-12       375ns ± 1%   376ns ± 1%   +0.35%  (p=0.021 n=20+18)
    AppendInPlace/NoGrow/1Ptr-12      1.01µs ± 1%  1.10µs ± 1%   +9.28%  (p=0.000 n=18+20)
    AppendInPlace/NoGrow/2Ptr-12      1.85µs ± 2%  1.71µs ± 1%   -7.51%  (p=0.000 n=19+18)
    AppendInPlace/NoGrow/3Ptr-12      2.57µs ± 2%  2.44µs ± 1%   -5.08%  (p=0.000 n=19+19)
    AppendInPlace/NoGrow/4Ptr-12      3.52µs ± 2%  3.35µs ± 2%   -4.70%  (p=0.000 n=20+19)
    AppendInPlace/Grow/Byte-12         212ns ± 1%   217ns ± 8%   +2.57%  (p=0.000 n=20+20)
    AppendInPlace/Grow/1Ptr-12         214ns ± 2%   217ns ± 3%   +1.23%  (p=0.001 n=18+19)
    AppendInPlace/Grow/2Ptr-12         298ns ± 2%   300ns ± 2%   +0.55%  (p=0.038 n=19+20)
    AppendInPlace/Grow/3Ptr-12         367ns ± 2%   366ns ± 2%     ~     (p=0.452 n=20+18)
    AppendInPlace/Grow/4Ptr-12         416ns ± 2%   411ns ± 2%   -1.18%  (p=0.000 n=20+19)
    StackGrowth-12                    43.4ns ± 1%  43.4ns ± 0%     ~     (p=1.000 n=16+16)
    StackGrowthDeep-12                11.4µs ± 4%  10.3µs ± 4%   -9.65%  (p=0.000 n=20+19)
    
    Change-Id: I69a8afbd942c787c591d95b9d9439bd6db4d1e49
    Reviewed-on: https://go-review.googlesource.com/30192
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3f952b75046548881709086cc20ab3c5fca2f52e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 4 14:40:38 2016 -0700

    cmd/compile/internal/ssa: update BlockKind documentation
    
    BlockCall was removed in golang.org/cl/28950.
    
    Change-Id: Ib8d9f3111bf3dc01956dd776afeb345ede8bc933
    Reviewed-on: https://go-review.googlesource.com/30353
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4caf93c9571ed637cf3876042d263b0d8938a2fe
Author: Oleg Vakheta <helginet@gmail.com>
Date:   Fri Nov 27 17:07:58 2015 +0200

    fmt: add tests for parsenum
    
    Change-Id: Ie7b869892816a171d8c71612998cc32a190aeff9
    Reviewed-on: https://go-review.googlesource.com/17227
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit c28f55c50219f0191e453ea02d57c1f20434b561
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 4 13:00:21 2016 -0700

    cmd/compile/internal/ssa: add Op.UsesScratch method
    
    Passes toolstash/buildall.
    
    Change-Id: I928a2ef39fb10091957f35bb3f1564498f6f1b83
    Reviewed-on: https://go-review.googlesource.com/30312
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 59320c396e6448132a52cb5a5d96491eee1e0ad8
Author: Quentin Renard <contact@asticode.com>
Date:   Sun Mar 6 17:27:50 2016 +0100

    net/http: improve performance for parsePostForm
    
    Remove the use of io.ReadAll in http.parsePostForm to avoid converting
    the whole input from []byte to string and not performing well
    space-allocated-wise.
    
    Instead a new function called parsePostFormURLEncoded is used and is
    fed directly an io.Reader that is parsed using a bufio.Reader.
    
    Benchmark:
    
    name         old time/op    new time/op    delta
    PostQuery-4    2.90µs ± 6%    2.82µs ± 4%     ~       (p=0.094 n=9+9)
    
    name         old alloc/op   new alloc/op   delta
    PostQuery-4    1.05kB ± 0%    0.90kB ± 0%  -14.49%  (p=0.000 n=10+10)
    
    name         old allocs/op  new allocs/op  delta
    PostQuery-4      6.00 ± 0%      7.00 ± 0%  +16.67%  (p=0.000 n=10+10)
    
    Fixes #14655
    
    Change-Id: I112c263d4221d959ed6153cfe88bc57a2aa8ea73
    Reviewed-on: https://go-review.googlesource.com/20301
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e6051de0351bb05d7409ee0d483f932e3530f816
Author: Keith Randall <khr@golang.org>
Date:   Tue Oct 4 10:01:48 2016 -0700

    cmd/compile: lower cse comparison depth
    
    Todd originally set cmpDepth to 4.  Quoting:
    
    I picked a depth of 4 by timing tests of `go tool compile arithConst_ssa.go` and `go test -c net/http`.
    
        3.89 / 3.92  CL w/cmpDepth = 1
        3.78 / 3.92  CL w/cmpDepth = 2
        3.44 / 3.96  CL w/cmpDepth = 3
        3.29 / 3.9   CL w/cmpDepth = 4
        3.3  / 3.93  CL w/cmpDepth = 5
        3.29 / 3.92  CL w/cmpDepth = 10
    
    I don't see the same behavior now, differences in those two benchmarks
    are in the noise (between 1 and 4).
    
    In issue 17127, CSE takes a really long time.  Lowering cmpDepth
    from 4 to 1 lowers compile time from 8 minutes to 1 minute.
    
    Fixes #17127
    
    Change-Id: I6dc544bbcf2a9dca73637d0182d3de1a5ae6c944
    Reviewed-on: https://go-review.googlesource.com/30257
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit fa49c3970a43ca77d46363397e7f4f3ae1f5957c
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Tue Oct 4 09:21:30 2016 -0700

    database/sql: fixup remaining driver call to use context
    
    Missed one in the prior CL.
    
    Change-Id: I6f6d84d52fe4d902a985971a402701fb3b1eed86
    Reviewed-on: https://go-review.googlesource.com/30255
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ddb77100a670d491d7553cc1beed630fc994a7f6
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Oct 3 12:13:22 2016 -0700

    reflect: ignore struct tags when converting structs
    
    Implementation of spec change https://golang.org/cl/24190/.
    
    For #16085.
    
    Change-Id: Ib7cb513354269282dfad663c7d2c6e624149f3cd
    Reviewed-on: https://go-review.googlesource.com/30191
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit f5b0012362f0ab801a657ff01d2d55f2391b1792
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Oct 3 12:12:20 2016 -0700

    go/types: ignore struct tags when converting structs
    
    Implementation of spec change https://golang.org/cl/24190/.
    
    For #16085.
    
    Change-Id: I17bbbce38d98a169bc64e84983a7ebfe7142f6e9
    Reviewed-on: https://go-review.googlesource.com/30190
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 39055700b1c69e791405518a914017b8c5551436
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Oct 3 11:40:43 2016 -0700

    cmd/compile: ignore struct tags when converting structs
    
    Implementation of spec change https://golang.org/cl/24190/.
    
    For #16085.
    
    Change-Id: Id71ef29af5031b073e8be163f578d1bb768ff97a
    Reviewed-on: https://go-review.googlesource.com/30169
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 9abaef93c75b8aef007624b66f99a671eb0cc5d6
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Oct 3 16:09:13 2016 -0700

    cmd/compile: cleanup artifacts from previous CL
    
    Does not pass toolstash, but only because it causes ATYPE instructions
    to be emitted in a different order, and it avoids emitting type
    metadata for unused variables.
    
    Change-Id: I3ec8f66a40b5af9213e0d6e852b267a8dd995838
    Reviewed-on: https://go-review.googlesource.com/30217
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5c7a005266f84ecea26859619630a862eccc0d48
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Jun 3 13:21:55 2016 -0700

    spec: ignore struct tags when converting structs
    
    This is a backwards-compatible language change.
    
    Per the proposal (#16085), the rules for conversions are relaxed
    such that struct tags in any of the structs involved in the conversion
    are ignored (recursively).
    
    Because this is loosening the existing rules, code that compiled so
    far will continue to compile.
    
    For #16085.
    Fixes #6858.
    
    Change-Id: I0feef651582db5f23046a2331fc3f179ae577c45
    Reviewed-on: https://go-review.googlesource.com/24190
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 8c24bff52b7d8e789382a8992af2e0adf0b491f2
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Oct 3 12:26:25 2016 -0700

    cmd/compile: layout stack frame during SSA
    
    Identify live stack variables during SSA and compute the stack frame
    layout earlier so that we can emit instructions with the correct
    offsets upfront.
    
    Passes toolstash/buildall.
    
    Change-Id: I191100dba274f1e364a15bdcfdc1d1466cdd1db5
    Reviewed-on: https://go-review.googlesource.com/30216
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit f239196b9e1c0d58ab2067630bd1db7210eb7f6a
Author: Keith Randall <khr@golang.org>
Date:   Tue Oct 4 09:49:33 2016 -0700

    cmd/compile: remove duplicate statement list function
    
    Probably a holdover from linked list vs. slice.
    
    Change-Id: Ib2540b08ef0ae48707d44a5d57bc23f8d65c760d
    Reviewed-on: https://go-review.googlesource.com/30256
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f0636bf6f97478aca313052a9661828d01357d75
Author: Dhananjay Nakrani <dhananjayn@google.com>
Date:   Sat Oct 1 18:01:30 2016 -0700

    cmd/cover: Fix compiler directives handling.
    
    Currently, it separates comments from rest of the AST. This causes problems when
    long counter increment statements are added before compiler directives.
    See Issue #17315.
    
    This change moves comments handling into AST Visitor so that when printer prints
    code from AST, position of compiler directives relative to the associated function
    is preserved.
    
    Tested with https://gist.github.com/dhananjay92/837df6bc1f171b1350f85d7a7d59ca1e
    and unit test.
    
    Fixes #17315
    
    Change-Id: I61a80332fc1923de6fc59ff63b953671598071fa
    Reviewed-on: https://go-review.googlesource.com/30161
    Reviewed-by: Rob Pike <r@golang.org>

commit 6300161d40e50902e16b3144c36aaf9279ab6208
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 30 09:03:38 2016 -0700

    cmd/compile: force folding of MOVDaddr into storezero
    
    Fold MOVDaddr ops into MOVXstorezero ops.
    Also fold ADDconst into MOVDaddr so we're sure there isn't
    (MOVDstorezero (ADDconst (MOVDaddr ..)))
    
    Without this CL, we get:
    
    v1 = MOVDaddr {s}
    v2 = VARDEF {s}
    v3 = MOVDstorezero v1 v2
    
    The liveness pass thinks the MOVDaddr is a read of s, so s is
    incorrectly thought to be live at the start of the function.
    
    Fixes #17194
    
    Change-Id: I2b4a2f13b12aa5b072941ee1c7b89f3793650cdc
    Reviewed-on: https://go-review.googlesource.com/30086
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit 7d0642d9d6a7a9c06d422904632ab61668e24a9e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Oct 4 05:21:58 2016 +0000

    regexp: remove dead code
    
    Wasn't convenient enough.
    
    Change-Id: I78270dc22cdb2e264641148e50029a9e4de953cd
    Reviewed-on: https://go-review.googlesource.com/30251
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6b795e77dfbaed05f0607ac14a3db3f62d905e70
Author: Nick Craig-Wood <nick@craig-wood.com>
Date:   Tue Oct 4 14:04:18 2016 +0100

    runtime: correct function name in throw message
    
    Change-Id: I8fd271066925734c3f7196f64db04f27c4ce27cb
    Reviewed-on: https://go-review.googlesource.com/30274
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f28cf8346c4ce7cb74bf97c7c69da21c43a78034
Author: Filippo Valsorda <hi@filippo.io>
Date:   Wed Dec 23 02:03:44 2015 +0000

    crypto/tls: implement countermeasures against CBC padding oracles
    
    The aim is to make the decrypt() timing profile constant, irrespective of
    the CBC padding length or correctness.  The old algorithm, on valid padding,
    would only MAC bytes up to the padding length threshold, making CBC
    ciphersuites vulnerable to plaintext recovery attacks as presented in the
    "Lucky Thirteen" paper.
    
    The new algorithm Write()s to the MAC all supposed payload, performs a
    constant time Sum()---which required implementing a constant time Sum() in
    crypto/sha1, see the "Lucky Microseconds" paper---and then Write()s the rest
    of the data. This is performed whether the padding is good or not.
    
    This should have no explicit secret-dependent timings, but it does NOT
    attempt to normalize memory accesses to prevent cache timing leaks.
    
    Updates #13385
    
    Change-Id: I15d91dc3cc6eefc1d44f317f72ff8feb0a9888f7
    Reviewed-on: https://go-review.googlesource.com/18130
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit ad26bb5e3098cbfd7c0ad9a1dc9d38c92e50f06e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Oct 4 03:01:09 2016 +0000

    all: use sort.Slice where applicable
    
    I avoided anywhere in the compiler or things which might be used by
    the compiler in the future, since they need to build with Go 1.4.
    
    I also avoided anywhere where there was no benefit to changing it.
    
    I probably missed some.
    
    Updates #16721
    
    Change-Id: Ib3c895ff475c6dec2d4322393faaf8cb6a6d4956
    Reviewed-on: https://go-review.googlesource.com/30250
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 2f184c65a5bdd422f88d841bb3a37fa60b3e1d52
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue Sep 27 19:54:05 2016 +0900

    net: implement network interface API for Solaris
    
    Fixes #7177.
    
    Change-Id: Iba6063905f4f9c6acef8aba76b55d996f186d835
    Reviewed-on: https://go-review.googlesource.com/29892
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>

commit cb6bb4062f6a36d0e76f6fe15f78e0bbcd4b71c0
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue Sep 27 19:51:47 2016 +0900

    vendor: import golang.org/x/net/lif
    
    golang.org/x/net/lif becomes vendor/golang_org/x/net/lif.
    
    At git rev 9f0e377 (golang.org/cl/29893)
    
    Updates #7177.
    
    Change-Id: Id838fcc234e71f735bb2609073f4c2214b48a970
    Reviewed-on: https://go-review.googlesource.com/29891
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 89c4cbd7acc36cf627746f379f0cc002b5b60383
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Oct 1 19:20:32 2016 +0900

    syscall: fix Send{msg,msgN}, Recvmsg and control message handling on solaris
    
    This change switches the use of socket implementation from the
    conventional SUS-based one to the latest POSIX-based one to make
    socket control message work correctly on Solaris.
    
    It looks like those two implementations, Socket over TLI/XTI and
    Socket, have different semantics in details but it wouldn't hurt
    the existing applications because the exposed syscall API doesn't
    support socket properties related to such a protocol independent
    application framework.
    
    Fixes #7402.
    
    Change-Id: I45a4e782d606bfbebe1404086c50a8c69af53461
    Reviewed-on: https://go-review.googlesource.com/30171
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 7e0218cdb204cec082601988324d54ef484a8f5f
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Sep 30 12:04:46 2016 -0400

    cmd/compile: remove unnecessary OpSB checks in S390X.rules
    
    Reversed, indexed and multi-register stores/loads cannot accept SB
    inputs. Therefore if one of these Ops is an input to a rule any
    pointer that is an argument to that Op cannot be OpSB.
    
    Change-Id: Ib8048362d1c6277122afec0d13a1c905290d69cb
    Reviewed-on: https://go-review.googlesource.com/30131
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 7aab88a31e4b689374832ebd1a3927bcef6e3f79
Author: Austin Clements <austin@google.com>
Date:   Mon Sep 26 13:41:39 2016 -0400

    runtime: fix missing space in error message
    
    Change-Id: I422708d50c3c727246e7991039877660ca034dc9
    Reviewed-on: https://go-review.googlesource.com/30144
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit bf776a988bf6fe82cbef5cfc954f33d127c5172c
Author: Austin Clements <austin@google.com>
Date:   Mon Sep 26 13:10:41 2016 -0400

    runtime: document bmap.tophash
    
    In particular, it wasn't obvious that some values are special (unless
    you also found those special values), so document that it isn't
    necessarily a hash value.
    
    Change-Id: Iff292822b44408239e26cd882dc07be6df2c1d38
    Reviewed-on: https://go-review.googlesource.com/30143
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 38f1df66ff231458a81cb07e7b147f30854b45d4
Author: Austin Clements <austin@google.com>
Date:   Fri Sep 9 10:22:10 2016 -0400

    runtime: make gcDumpObject useful on stack frames
    
    gcDumpObject is often used on a stack pointer (for example, when
    checkmark finds an unmarked object on the stack), but since stack
    spans don't have an elemsize, it doesn't print any of the memory from
    the frame. Make it at least slightly more useful by printing
    everything between obj and obj+off (inclusive). While we're here, also
    print out the span state.
    
    Change-Id: I51be064ea8791b4a365865bfdc7afa7b5aaecfbd
    Reviewed-on: https://go-review.googlesource.com/30142
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 6879dbde4e6dceceb272b83fc6682e97b9de2fa1
Author: Austin Clements <austin@google.com>
Date:   Fri Sep 9 10:31:27 2016 -0400

    runtime: introduce a type for span states
    
    Currently span states are untyped constants and the field is just a
    uint8. Make this more type-safe by introducing a type for the span
    state.
    
    Change-Id: I369bf59fe6e8234475f4921611424fceb7d0a6de
    Reviewed-on: https://go-review.googlesource.com/30141
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 5a6e511c614a158cb58150fb62bfbc207a33922d
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 30 10:12:32 2016 -0700

    cmd/compile: Use Sreedhar+Gao phi building algorithm
    
    Should be more asymptotically happy.
    
    We process each variable in turn to find all the
    locations where it needs a phi (the dominance frontier
    of all of its definitions).  Then we add all those phis.
    This takes O(n * #variables), although hopefully much less.
    
    Then we do a single tree walk to match all the
    FwdRefs with the nearest definition or phi.
    This takes O(n) time.
    
    The one remaining inefficiency is that we might end up
    introducing a bunch of dead phis in the first step.
    A TODO is to introduce phis only where they might be
    used by a read.
    
    The old algorithm is still faster on small functions,
    so there's a cutover size (currently 500 blocks).
    
    This algorithm supercedes the David's sparse phi
    placement algorithm for large functions.
    
    Lowers compile time of example from #14934 from
    ~10 sec to ~4 sec.
    Lowers compile time of example from #16361 from
    ~4.5 sec to ~3 sec.
    Lowers #16407 from ~20 min to ~30 sec.
    
    Update #14934
    Update #16361
    Fixes #16407
    
    Change-Id: I1cff6364e1623c143190b6a924d7599e309db58f
    Reviewed-on: https://go-review.googlesource.com/30163
    Reviewed-by: David Chase <drchase@google.com>

commit d0e92f61e5c5c59395d9b1a3b4f5c7b90dec5bc8
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Oct 3 15:21:16 2016 -0400

    cmd/compile: remove unnecessary write barriers for APPEND
    
    Updates #17330.
    
    Change-Id: I83fe80139a2213f3169db884b84a4c3bd15b58b6
    Reviewed-on: https://go-review.googlesource.com/30140
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit fe77c5b36994d907a6c0f4cd8ffb0a2ad6cfde5e
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Oct 3 11:25:25 2016 -0700

    doc: add PKG_CONFIG and GIT_ALLOW_PROTOCOL env vars to go1.8.txt
    
    Change-Id: I592b87f49fc636b89807d911132f69257d718afd
    Reviewed-on: https://go-review.googlesource.com/30168
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 55620a0e910fd9f2a9d7d631f01a6bddfe0f6a0a
Author: Billy Lynch <wlynch@google.com>
Date:   Fri Sep 30 16:24:31 2016 -0400

    cmd/go: add support for GIT_ALLOW_PROTOCOL
    
    Allows users to override the default secure protocol list by setting the
    GIT_ALLOW_PROTOCOL environment variable.
    
    Addresses #17299 for vcs.go.
    
    Change-Id: If575861d2b1b04b59029fed7e5d12b49690af50a
    Reviewed-on: https://go-review.googlesource.com/30135
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5dcb31b2d555be0c7d1c0c4b2001d078e7ac6078
Author: Boris Nagaev <nagaev@google.com>
Date:   Sat Jun 25 13:51:06 2016 +0200

    cmd/dist, cmd/go: add environment variable override for pkg-config
    
    Allow overriding default name of `pkg-config` tool via environment
    variable PKG_CONFIG (same as used by autoconf pkg.m4 macros). This
    facilitates easy cross-compilation of cgo code.
    
    Original patch against Go <= 1.4 was written by
    xnox_canonical <dimitri.ledkov@canonical.com> in 2014.
    Source: https://codereview.appspot.com/104960043/
    
    Fixes #16253
    
    Change-Id: I31c33ffc3ecbff65da31421e6188d092ab4fe7e4
    Reviewed-on: https://go-review.googlesource.com/29991
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6c5e377d23136bd371d205e1c2aae6ddccb4597e
Author: Than McIntosh <thanm@google.com>
Date:   Thu Jul 14 13:23:11 2016 -0400

    cmd/compile: relax liveness restrictions on ambiguously live
    
    Update gc liveness to remove special conservative treatment
    of ambiguously live vars, since there is no longer a need to
    protect against GCDEBUG=gcdead.
    
    Change-Id: Id6e2d03218f7d67911e8436d283005a124e6957f
    Reviewed-on: https://go-review.googlesource.com/24896
    Reviewed-by: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit 5f36e9a3062c1f133169d01d612da9458a7ea884
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Oct 1 20:28:09 2016 -0700

    net: clarify that Conn deadlines also affect currently-blocked I/O
    
    All implementations have always implemented this behavior, it's
    tested, and it's depended on by other packages. (notably, by net/http)
    
    The one exception is Plan 9 which doesn't support I/O deadlines at all
    (tracked in #11932). As a result, a bunch of tests fail on plan9
    (#7237). But once Plan 9 adds I/O deadline support, it'll also need
    this behavior.
    
    Change-Id: Idb71767f0c99279c66dce29f7bdc78ef467e47aa
    Reviewed-on: https://go-review.googlesource.com/30164
    Reviewed-by: Sam Whited <sam@samwhited.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 99339dd44537f82c9a1348b8937b68b6c61be005
Author: Austin Clements <austin@google.com>
Date:   Mon Oct 3 11:52:20 2016 -0400

    runtime: weaken claim about SetFinalizer panicking
    
    Currently the SetFinalizer documentation makes a strong claim that
    SetFinalizer will panic if the pointer is not to an object allocated
    by calling new, to a composite literal, or to a local variable. This
    is not true. For example, it doesn't panic when passed the address of
    a package-level variable. Nor can we practically make it true. For
    example, we can't distinguish between passing a pointer to a composite
    literal and passing a pointer to its first field.
    
    Hence, weaken the guarantee to say that it "may" panic.
    
    Updates #17311. (Might fix it, depending on what we want to do with
    package-level variables.)
    
    Change-Id: I1c68ea9d0a5bbd3dd1b7ce329d92b0f05e2e0877
    Reviewed-on: https://go-review.googlesource.com/30137
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 22a2bdfedb95612984cec3141924953b88a607b7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Aug 17 14:29:00 2016 +0000

    sort: add Slice, SliceStable, and SliceIsSorted
    
    Add helpers for sorting slices.
    
    Slice sorts slices:
    
        sort.Slice(s, func(i, j int) bool {
            if s[i].Foo != s[j].Foo {
                return s[i].Foo < s[j].Foo
            }
            return s[i].Bar < s[j].Bar
        })
    
    SliceStable is the same, but does a stable sort.
    
    SliceIsSorted reports whether a slice is already sorted.
    
    Fixes #16721
    
    Change-Id: I346530af1c5dee148ea9be85946fe08f23ae53e7
    Reviewed-on: https://go-review.googlesource.com/27321
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 003a598bf28ae8f77919c3bda2abb9d6d4f449bb
Author: Florian Uekermann <florian@uekermann.me>
Date:   Wed Aug 17 18:45:57 2016 +0200

    math/rand: add Rand.Uint64
    
    This adds Uint64 methods to Rand and rngSource.
    Rand.Uint64 uses Source.Uint64 directly if it is present.
    
    rngSource.Uint64 provides access to all 64 bits generated by the
    underlying ALFG. To ensure high seed quality a 64th bit has been added
    to all elements of the array of "cooked" random numbers that are used
    for seeding. gen_cooked.go generates both the 63 bit and 64 bit array.
    
    Fixes #4254
    
    Change-Id: I22855618ac69abae3d2799b3e7e59996d4c5a4b1
    Reviewed-on: https://go-review.googlesource.com/27253
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 998419575fb34dd5228cfdb353a73184e29db8da
Author: Adam Langley <agl@golang.org>
Date:   Fri Sep 30 14:48:11 2016 -0700

    crypto/ecdsa: correct code comment.
    
    The code comment mixed up max and min. In this case, min is correct
    because this entropy is only used to make the signature scheme
    probabilistic. (I.e. if it were fixed then the scheme would still be
    secure except that key.Sign(foo) would always give the same result for a
    fixed key and foo.)
    
    For this purpose, 256-bits is plenty.
    
    Fixes #16819.
    
    Change-Id: I309bb312b775cf0c4b7463c980ba4b19ad412c36
    Reviewed-on: https://go-review.googlesource.com/30153
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 49aa1d791be26de71ba7ed02d6c6cd1dd0092b71
Author: Adam Langley <agl@golang.org>
Date:   Fri Sep 30 14:16:12 2016 -0700

    crypto/x509: return better error when a certificate contains no names.
    
    Currently, if a certificate contains no names (that we parsed),
    verification will return the confusing error:
        x509: certificate is valid for , not example.com.
    
    This change improves the error for that situation.
    
    Fixes #16834.
    
    Change-Id: I2ed9ed08298d7d50df758e503bdb55277449bf55
    Reviewed-on: https://go-review.googlesource.com/30152
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e4dafa32620e80e4e39937d8e2033fb2ee6085f8
Author: Adam Langley <agl@golang.org>
Date:   Fri Sep 30 16:54:54 2016 -0700

    crypto/x509: fix name constraints handling.
    
    This change brings the behaviour of X.509 name constraints into line
    with NSS[1]. In this area, the behavior specified by the RFC and by NIST
    differs and this code follows the NIST behaviour.
    
    [1] https://github.com/servo/nss/blob/master/lib/certdb/genname.c
    
    Fixes #16347, fixes #14833.
    
    Change-Id: I5acd1970041291c2e3936f5b1fd36f2a0338e613
    Reviewed-on: https://go-review.googlesource.com/30155
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2d573eee8ae532a3720ef4efbff9c8f42b6e8217
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sat Jul 23 15:41:57 2016 -0700

    cmd/compile: improve error message for wrong number of arguments to return
    
    Fixes #4215.
    Fixes #6750.
    
    Improves the error message for wrong number of arguments by comparing
    the signature of the return call site arguments, versus the function's
    expected return arguments.
    
    In this CL, the signature representation of:
    + ideal numbers(TIDEAL) ie float*, complex*, rune, int is
    "number" instead of "untyped number".
    + idealstring is "string" instead of "untyped string".
    + idealbool is "bool" instead of "untyped bool".
    
    However, the representation of other types remains as the compiler
    would produce.
    
    * Example 1(in the error messages, if all lines were printed):
    $ cat main.go && go run main.go
    package main
    
    func foo() (int, int) {
      return 2.3
    }
    
    func foo2() {
      return int(2), 2
    }
    
    func foo3(v int) (a, b, c, d int) {
      if v >= 5 {
        return 1
      }
      return 2, 3
    }
    
    func foo4(name string) (string, int) {
      switch name {
      case "cow":
        return "moo"
      case "dog":
        return "dog", 10, true
      case "fish":
        return ""
      default:
        return "lizard", 10
      }
    }
    
    type S int
    type T string
    type U float64
    
    func foo5() (S, T, U) {
      if false {
        return ""
      } else {
        ptr := new(T)
        return ptr
      }
      return new(S), 12.34, 1 + 0i, 'r', true
    }
    
    func foo6() (T, string) {
      return "T"
    }
    
    ./issue4215.go:4: not enough arguments to return, got (number) want (int, int)
    ./issue4215.go:8: too many arguments to return, got (int, number) want ()
    ./issue4215.go:13: not enough arguments to return, got (number) want (int, int, int, int)
    ./issue4215.go:15: not enough arguments to return, got (number, number) want (int, int, int, int)
    ./issue4215.go:21: not enough arguments to return, got (string) want (string, int)
    ./issue4215.go:23: too many arguments to return, got (string, number, bool) want (string, int)
    ./issue4215.go:25: not enough arguments to return, got (string) want (string, int)
    ./issue4215.go:37: not enough arguments to return, got (string) want (S, T, U)
    ./issue4215.go:40: not enough arguments to return, got (*T) want (S, T, U)
    ./issue4215.go:42: too many arguments to return, got (*S, number, number, number, bool) want (S, T, U)
    ./issue4215.go:46: not enough arguments to return, got (string) want (T, string)
    ./issue4215.go:46: too many errors
    
    * Example 2:
    $ cat 6750.go && go run 6750.go
    package main
    
    import "fmt"
    
    func printmany(nums ...int) {
      for i, n := range nums {
        fmt.Printf("%d: %d\n", i, n)
      }
      fmt.Printf("\n")
    }
    
    func main() {
      printmany(1, 2, 3)
      printmany([]int{1, 2, 3}...)
      printmany(1, "abc", []int{2, 3}...)
    }
    ./issue6750.go:15: too many arguments in call to printmany, got (number, string, []int) want (...int)
    
    Change-Id: I6fdce78553ae81770840070e2c975d3e3c83d5d8
    Reviewed-on: https://go-review.googlesource.com/25156
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 2878cf14f3bb4c097771e50a481fec43962d7401
Author: Adam Langley <agl@golang.org>
Date:   Fri Sep 30 12:55:25 2016 -0700

    crypto/tls: simplify keylog tests.
    
    Since there's no aspect of key logging that OpenSSL can check for us,
    the tests for it might as well just connect to another goroutine as this
    is lower-maintainance.
    
    Change-Id: I746d1dbad1b4bbfc8ef6ccf136ee4824dbda021e
    Reviewed-on: https://go-review.googlesource.com/30089
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joonas Kuorilehto <joneskoo@derbian.fi>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7b40b0c3a332cbfaa1eb17bdafd2ddf12119ec45
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Wed Sep 28 01:54:38 2016 -0700

    strings, bytes: panic if Repeat overflows or if given a negative count
    
    Panic if Repeat is given a negative count or
    if the value of (len(*) * count) is detected
    to overflow.
    We panic because we cannot change the
    signature of Repeat to return an error.
    
    Fixes #16237
    
    Change-Id: I9f5ba031a5b8533db0582d7a672ffb715143f3fb
    Reviewed-on: https://go-review.googlesource.com/29954
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d166a369a89ef2d81efdc5d49fa782ee1c0186c4
Author: Matt Layher <mdlayher@gmail.com>
Date:   Sat Oct 1 10:13:52 2016 -0400

    bufio: remove redundant Writer.flush method
    
    Fixes #17232
    
    Change-Id: I34df86f79b643dce9f054c6df6782e6037c06910
    Reviewed-on: https://go-review.googlesource.com/30158
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 05f599594ac4f8db63d6a2d628b6ba5781273cad
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Oct 1 15:41:03 2016 +0900

    syscall: re-enable TestPassFD on dragonfly
    
    At least it works well on DragonFly BSD 4.6.
    
    Change-Id: I3b210745246c6d8d42e32ba65ee3b9a17d171ff7
    Reviewed-on: https://go-review.googlesource.com/30170
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 09fb7956fa277912d1af9dbebbbfba7502e3a051
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 30 20:31:26 2016 +0000

    net/http: don't sniff Request.Body on 100-continue requests in Transport
    
    Also, update bundled http2 to x/net git rev 0d8126f to include
    https://golang.org/cl/30150, the HTTP/2 version of this fix.
    
    Fixes #16002
    
    Change-Id: I8da1ca98250357aec012e3e85c8b13acfa2f3fec
    Reviewed-on: https://go-review.googlesource.com/30151
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 360f2e43b78a3ea119ea8dce9649f7c1227d793b
Author: Mike Appleby <mike@app.leby.org>
Date:   Wed Sep 28 16:01:27 2016 -0500

    runtime: sleep on CLOCK_MONOTONIC in futexsleep1 on freebsd
    
    In FreeBSD 10.0, the _umtx_op syscall was changed to allow sleeping on
    any supported clock, but the default clock was switched from a monotonic
    clock to CLOCK_REALTIME.
    
    Prior to 10.0, the __umtx_op_wait* functions ignored the fourth argument
    to _umtx_op (uaddr1), expected the fifth argument (uaddr2) to be a
    struct timespec pointer, and used a monotonic clock (nanouptime(9)) for
    timeout calculations.
    
    Since 10.0, if callers want a clock other than CLOCK_REALTIME, they must
    call _umtx_op with uaddr1 set to a value greater than sizeof(struct
    timespec), and with uaddr2 as pointer to a struct _umtx_time, rather
    than a timespec. Callers can set the _clockid field of the struct
    _umtx_time to request the clock they want.
    
    The relevant FreeBSD commit:
        https://svnweb.freebsd.org/base?view=revision&revision=232144
    
    Fixes #17168
    
    Change-Id: I3dd7b32b683622b8d7b4a6a8f9eb56401bed6bdf
    Reviewed-on: https://go-review.googlesource.com/30154
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d1f4e0413fd2b5cb78a7b645e802565f202d1926
Author: Mike Appleby <mike@app.leby.org>
Date:   Fri Sep 30 19:43:42 2016 -0500

    time: update comment to reflect correct file path.
    
    Update cross-reference in the comment for runtimeTimer to point to the
    new go file instead of the old .h file.
    
    Change-Id: Iddb3614c41e1989096d6caf77d6c0d5781005181
    Reviewed-on: https://go-review.googlesource.com/30157
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit eca4e446115be5653a3963c37459a263569390c5
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Aug 2 18:42:58 2016 -0700

    cmd/doc: perform type grouping for constants and variables
    
    In golang.org/cl/22354, we added functionality to group functions under the
    type that they construct to. In this CL, we extend the same concept to
    constants and variables. This makes the doc tool more consistent with what
    the godoc website does.
    
    $ go doc reflect | egrep "ChanDir|Kind|SelectDir"
    <<<
    // Before:
    const RecvDir ChanDir = 1 << iota ...
    const Invalid Kind = iota ...
    type ChanDir int
    type Kind uint
    type SelectDir int
        func ChanOf(dir ChanDir, t Type) Type
    
    // After:
    type ChanDir int
        const RecvDir ChanDir = 1 << iota ...
    type Kind uint
        const Invalid Kind = iota ...
    type SelectDir int
        const SelectSend SelectDir ...
        func ChanOf(dir ChanDir, t Type) Type
    >>>
    
    Furthermore, a fix was made to ensure that the type was printed in constant
    blocks when the iota was applied on an unexported field.
    
    $ go doc reflect SelectSend
    <<<
    // Before:
    const (
            SelectSend    // case Chan <- Send
            SelectRecv    // case <-Chan:
            SelectDefault // default
    )
    
    // After:
    const (
            SelectSend    SelectDir // case Chan <- Send
            SelectRecv              // case <-Chan:
            SelectDefault           // default
    )
    >>>
    
    Fixes #16569
    
    Change-Id: I26124c3d19e50caf9742bb936803a665e0fa6512
    Reviewed-on: https://go-review.googlesource.com/25419
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c5f064ee49e98c00c7959fdf095e2c61ff0747b8
Author: Matt Layher <mdlayher@gmail.com>
Date:   Fri Sep 30 17:23:24 2016 -0400

    context: discourage use of basic types as keys in WithValue
    
    Fixes #17302
    
    Change-Id: I375d5d4f2714ff415542f4fe56a548e53c5e8ba6
    Reviewed-on: https://go-review.googlesource.com/30134
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ca04091f5be4134291142cc4e7e577d0f627e788
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 30 18:12:37 2016 +0000

    reflect: add Swapper func
    
    Swapper returns a func that swaps two elements in a slice.
    
    Updates #16721
    
    Change-Id: I7f2287a675c10a05019e02b7d62fb870af31216f
    Reviewed-on: https://go-review.googlesource.com/30088
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 9491e7d65e09644eb7db4e2ed5ff0139571cedf3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 30 10:40:58 2016 -0700

    net/http: refactor testing of Request.Body on 0 ContentLength
    
    Code movement only, to look more like the equivalent http2 code, and
    to make an upcoming fix look more obvious.
    
    Updates #16002 (to be fixed once this code is in)
    
    Change-Id: Iaa4f965be14e98f9996e7c4624afe6e19bed1a80
    Reviewed-on: https://go-review.googlesource.com/30087
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit ab6ba99484b637bad0c5a5fa2c590834c14746c7
Author: David Benjamin <davidben@google.com>
Date:   Fri Jul 1 16:41:09 2016 -0400

    crypto/tls: Fix c.in.decrypt error handling.
    
    readRecord was not returning early if c.in.decrypt failed and ran
    through the rest of the function. It does set c.in.err, so the various
    checks in the callers do ultimately notice before acting on the result,
    but we should avoid running the rest of the function at all.
    
    Also rename 'err' to 'alertValue' since it isn't actually an error.
    
    Change-Id: I6660924716a85af704bd3fe81521b34766238695
    Reviewed-on: https://go-review.googlesource.com/24709
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit 01661612e433deadec99229c6075baa4174ac743
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Fri Sep 30 11:27:52 2016 +0200

    cmd/compile: update error messages in Mpint, Mpflt methods
    
    CL 20909 gave Mpint methods nice go-like names, but it
    didn't update the names in the error strings. Fix them.
    
    Same for a couple of Mpflt methods.
    
    Change-Id: I9c99653d4b922e32fd5ba18aba768a589a4c7869
    Reviewed-on: https://go-review.googlesource.com/30091
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
```
