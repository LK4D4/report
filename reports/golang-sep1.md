# September 1, 2016 Report

Number of commits: 117

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 609.175035ms to 605.841615ms, -0.55%
* github.com/gogits/gogs: from 13.302608381s to 13.624899752s, +2.42%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 32.558075879s to 32.622095471s, +0.20%
* github.com/influxdata/influxdb/cmd/influxd: from 6.980130038s to 7.074529379s, +1.35%
* github.com/junegunn/fzf/src/fzf: from 1.102806841s to 1.189099401s, +7.82%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.499478436s to 1.469706458s, -1.99%
* github.com/nsqio/nsq/apps/nsqd: from 2.313674102s to 2.312804545s, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 11.389928432s to 11.565382461s, +1.54%
* github.com/spf13/hugo: from 7.572876912s to 10.014020204s, +32.24%
* golang.org/x/tools/cmd/guru: from 3.089123108s to 3.094063816s, +0.16%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2597269 to 2572052, -0.97%
* github.com/gogits/gogs: from 23327625 to 23112184, -0.92%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 27252536 to 27151560, -0.37%
* github.com/influxdata/influxdb/cmd/influxd: from 16084573 to 16047173, -0.23%
* github.com/junegunn/fzf/src/fzf: from 3159936 to 3129560, -0.96%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4459642 to 4425925, -0.76%
* github.com/nsqio/nsq/apps/nsqd: from 9678149 to 9617376, -0.63%
* github.com/prometheus/prometheus/cmd/prometheus: from 23472559 to 23460240, ~
* github.com/spf13/hugo: from 15553961 to 15450646, -0.66%
* golang.org/x/tools/cmd/guru: from 7616532 to 7570389, -0.61%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               191           188           -1.57%
BenchmarkMsgpUnmarshal-4             404           393           -2.72%
BenchmarkEasyJsonMarshal-4           1502          1472          -2.00%
BenchmarkEasyJsonUnmarshal-4         2492          2137          -14.25%
BenchmarkFlatBuffersMarshal-4        549           500           -8.93%
BenchmarkFlatBuffersUnmarshal-4      291           285           -2.06%
BenchmarkGogoprotobufMarshal-4       161           158           -1.86%
BenchmarkGogoprotobufUnmarshal-4     258           248           -3.88%

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

* [hash/crc32: improve the AMD64 implementation using SSE4.2](https://github.com/golang/go/commit/90c3cf4b52cc9373a96009da0d013019c1f5bcd8)
* [time: Add Until helper function](https://github.com/golang/go/commit/67ea710792eabdae1182e2bf4845f512136cccce)
* [cmd/compile: intrinsify sync/atomic for amd64](https://github.com/golang/go/commit/d6098e4277bab633c2df752ed90e1e826918ca67)
* [cmd/compile: improve string iteration performance](https://github.com/golang/go/commit/0dae9dfb08a30983cce1114742c974077bdf5e18)
* [runtime: improve memmove for amd64](https://github.com/golang/go/commit/3607c5f4f18ad4d423e40996ebf7f46b2f79ce02)
* [Revert "runtime: improve memmove for amd64"](https://github.com/golang/go/commit/6fb4b15f98bba7ef3966c5edc6b8fe2cc99c6beb)
* [doc/faq: explain the meaning of "runtime"](https://github.com/golang/go/commit/d1a19235212d62843c17dc4f7c61d46bb1bf56ff)
* [crypto/tls: add Config.Clone](https://github.com/golang/go/commit/d24f446a90ea94b87591bf16228d7d871fec3d92)


## GIT Log

```
commit 1c53a1b1975adf69c594fbbd5b1ca13d783f9817
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Sep 1 07:22:23 2016 -0400

    cmd/compile: fix scheduling of memory-producing tuple ops
    
    Intrinsified atomic op produces <value,memory>. Make sure this
    memory is considered in the store chain calculation.
    
    Fixes #16948.
    
    Change-Id: I029f164b123a7e830214297f8373f06ea0bf1e26
    Reviewed-on: https://go-review.googlesource.com/28350
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit f8555ea6fdbbfc32e26f351ac16138fad31a2d62
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Aug 18 13:14:30 2016 -0700

    spec: update language on type switches to match implementations
    
    See the issue below for details.
    
    Fixes #16794.
    
    Change-Id: I7e338089fd80ddcb634fa80bfc658dee2772361c
    Reviewed-on: https://go-review.googlesource.com/27356
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d24f446a90ea94b87591bf16228d7d871fec3d92
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Aug 30 03:19:01 2016 +0000

    crypto/tls: add Config.Clone
    
    In Go 1.0, the Config struct consisted only of exported fields.
    
    In Go 1.1, it started to grow private, uncopyable fields (sync.Once,
    sync.Mutex, etc).
    
    Ever since, people have been writing their own private Config.Clone
    methods, or risking it and doing a language-level shallow copy and
    copying the unexported sync variables.
    
    Clean this up and export the Config.clone method as Config.Clone.
    This matches the convention of Template.Clone from text/template and
    html/template at least.
    
    Fixes #15771
    Updates #16228 (needs update in x/net/http2 before fixed)
    Updates #16492 (not sure whether @agl wants to do more)
    
    Change-Id: I48c2825d4fef55a75d2f99640a7079c56fce39ca
    Reviewed-on: https://go-review.googlesource.com/28075
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit cd0ba4c169b591cc22f51cb61463eb45af7b930d
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Mon Aug 29 16:10:32 2016 -0700

    archive/tar: make Reader error handling consistent
    
    The tar.Reader guarantees stickiness of errors. Ensuring this property means
    that the methods of Reader need to be consistent about whose responsibility it
    is to actually ensure that errors are sticky.
    
    In this CL, we make it only the responsibility of the exported methods
    (Next and Read) to store tr.err. All other methods just return the error as is.
    
    As part of this change, we also check the error value of mergePAX (and test
    that it properly detects invalid PAX files). Since the value of mergePAX was
    never used before, we change it such that it always returns ErrHeader instead
    of strconv.SyntaxError. This keeps it consistent with other usages of strconv
    in the same tar package.
    
    Change-Id: Ia1c31da71f1de4c175da89a385dec665d3edd167
    Reviewed-on: https://go-review.googlesource.com/28215
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d1a19235212d62843c17dc4f7c61d46bb1bf56ff
Author: Rob Pike <r@golang.org>
Date:   Sat Aug 27 12:09:38 2016 +1000

    doc/faq: explain the meaning of "runtime"
    
    This truly is a common point of confusion that deserves
    explanation in the FAQ.
    
    Change-Id: Ie624e31a2042ca99626fe7570d9c8c075aae6a84
    Reviewed-on: https://go-review.googlesource.com/28275
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6fb4b15f98bba7ef3966c5edc6b8fe2cc99c6beb
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Wed Aug 31 20:44:42 2016 +0000

    Revert "runtime: improve memmove for amd64"
    
    This reverts commit 3607c5f4f18ad4d423e40996ebf7f46b2f79ce02.
    
    This was causing failures on amd64 machines without AVX.
    
    Fixes #16939
    
    Change-Id: I70080fbb4e7ae791857334f2bffd847d08cb25fa
    Reviewed-on: https://go-review.googlesource.com/28274
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit cc0248aea53b252ec5c0e1c57e32edb102bc36fe
Author: Keith Randall <khr@golang.org>
Date:   Wed Aug 31 12:35:32 2016 -0700

    cmd/compile: don't reserve X15 for float sub/div any more
    
    We used to reserve X15 to implement the 3-operand floating-point
    sub/div ops with the 2-operand sub/div that 386/amd64 gives us.
    
    Now that resultInArg0 is implemented, we no longer need to
    reserve X15 (X7 on 386).
    
    Fixes #15584
    
    Change-Id: I978e6c0a35236e89641bfc027538cede66004e82
    Reviewed-on: https://go-review.googlesource.com/28272
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 33bb597d859e6d98b4abb27592cb925753764136
Author: Keith Randall <khr@golang.org>
Date:   Wed Aug 31 12:30:46 2016 -0700

    cmd/compile: print SizeAndAlign AuxInt values correctly
    
    Makes the AuxInt arg to Move/Zero print in a readable format.
    
    Change-Id: I12295959b00ff7c1638d35836cc6d64d112c11ca
    Reviewed-on: https://go-review.googlesource.com/28271
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 00459f05e099d92b9ae4d50980b11767aca5e102
Author: Martin Möhrmann <martisch@uos.de>
Date:   Wed Aug 31 12:37:19 2016 +0200

    cmd/compile: fold negation into comparison operators
    
    This allows for example AMD64 ssa to generate
    (SETNE x) instead of (XORLconst [1] SETE).
    
    make.bash trigger count on AMD64:
    691 generic.rules:225
      1 generic.rules:226
      4 generic.rules:228
      1 generic.rules:229
      8 generic.rules:231
      6 generic.rules:238
      2 generic.rules:257
    
    Change-Id: I5b9827b2df63c8532675079e5a6026aa47bfd8dc
    Reviewed-on: https://go-review.googlesource.com/28232
    Run-TryBot: Martin Möhrmann <martisch@uos.de>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit ee161e859166b8b8b077998c0101f56c18b18329
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 30 16:31:53 2016 -0700

    cmd/compile: handle pragmas immediately with -newparser=1
    
    Instead of saving all pragmas and processing them after parsing is
    finished, process them immediately during scanning like the current
    lexer does.
    
    This is a bit unfortunate because it means we can't use
    syntax.ParseFile to concurrently parse files yet, but it fixes how we
    report syntax errors in the presence of //line pragmas.
    
    While here, add a bunch more gcCompat entries to syntax/parser.go to
    get "go build -toolexec='toolstash -cmp' std cmd" passing. There are
    still a few remaining cases only triggered building unit tests, but
    this seems like a nice checkpoint.
    
    Change-Id: Iaf3bbcf2849857a460496f31eea228e0c585ce13
    Reviewed-on: https://go-review.googlesource.com/28226
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 69e7e8a696825fa818b70587563ac68e52f8b1a1
Author: Edward Muller <edwardam@interlix.com>
Date:   Tue Aug 30 19:14:46 2016 -0700

    doc: update go tour installation instructions
    
    Fixes #16933
    
    Change-Id: I2054abd28bc555b018309934774fc4ecc44826b3
    Reviewed-on: https://go-review.googlesource.com/28217
    Reviewed-by: Emmanuel Odeke <emm.odeke@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ffa2bd27a47ef16e4d6a404dd15781ed5ba21e5d
Author: Kevin Burke <kev@inburke.com>
Date:   Wed Aug 31 09:09:01 2016 -0700

    runtime: fix typo
    
    Change-Id: I47e3cfa8b49e3d0b55c91387df31488b37038a8f
    Reviewed-on: https://go-review.googlesource.com/28225
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3607c5f4f18ad4d423e40996ebf7f46b2f79ce02
Author: Denis Nagorny <denis.nagorny@intel.com>
Date:   Thu Apr 28 12:25:46 2016 +0300

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
    
    Change-Id: Iab488d93a293cdf573ab5cd89b95a818bbb5d531
    Reviewed-on: https://go-review.googlesource.com/22515
    Run-TryBot: Denis Nagorny <denis.nagorny@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 04ade8e428fdd302ab8666d3fc5d8953caa4abcb
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Thu Aug 25 16:39:07 2016 +0300

    cmd/internal/obj/x86: Make VPSHUFD accept negative constant
    
    This partially reverts commit 4e24e1d9996b0b0155c8349e49244d9694c89708.
    Since in release 1.7 VPSHUFD support negative constant as an argument,
    removing it as part of 4e24e1d9996b0b0155c8349e49244d9694c89708 was wrong.
    Add it back.
    
    Change-Id: Id1a3e062fe8fb4cf538edb3f9970f0664f3f545f
    Reviewed-on: https://go-review.googlesource.com/27712
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bdde10137b3e67383f38329b02b329a906b78d5d
Author: Radu Berinde <radu@cockroachlabs.com>
Date:   Sun Aug 28 14:36:06 2016 -0400

    hash/crc32: cleanup code and improve tests
    
    Major reorganization of the crc32 code:
    
     - The arch-specific files now implement a well-defined interface
       (documented in crc32.go). They no longer have the responsibility of
       initializing and falling back to a non-accelerated implementation;
       instead, that happens in the higher level code.
    
     - The non-accelerated algorithms are moved to a separate file with no
       dependencies on other code.
    
     - The "cutoff" optimization for slicing-by-8 is moved inside the
       algorithm itself (as opposed to every callsite).
    
    Tests are significantly improved:
     - direct tests for the non-accelerated algorithms.
     - "cross-check" tests for arch-specific implementations (all archs).
     - tests for misaligned buffers for both IEEE and Castagnoli.
    
    Fixes #16909.
    
    Change-Id: I9b6dd83b7a57cd615eae901c0a6d61c6b8091c74
    Reviewed-on: https://go-review.googlesource.com/27935
    Reviewed-by: Keith Randall <khr@golang.org>

commit 2a2cab2911083f1ed2ea5c6bf24a12e2c5f6bcfc
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Tue Aug 30 18:07:39 2016 +0300

    math: speed up bessel functions on AMD64
    
    J0-4            71.9ns ± 1%  54.6ns ± 0%  -24.08%  (p=0.000 n=20+18)
    J1-4            71.6ns ± 0%  55.4ns ± 0%  -22.60%  (p=0.000 n=19+20)
    Jn-4             153ns ± 0%   118ns ± 1%  -22.71%  (p=0.000 n=20+20)
    Y0-4            70.8ns ± 0%  53.9ns ± 0%  -23.87%  (p=0.000 n=19+19)
    Y1-4            70.8ns ± 0%  54.1ns ± 0%  -23.54%  (p=0.000 n=20+20)
    Yn-4             149ns ± 0%   116ns ± 0%  -22.15%  (p=0.000 n=19+20)
    
    Fixes #16889
    
    Change-Id: Ie88496407b42f6acb918ffae1226b1b4c0500cb9
    Reviewed-on: https://go-review.googlesource.com/28086
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit dfbbe06a205e7048a8541c4c97b250c24c40db96
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Aug 10 11:06:46 2016 +1000

    cmd/link, cmd/go: delay linking of mingwex and mingw32 until very end
    
    cmd/go links mingwex and mingw32 libraries to every package it builds.
    This breaks when 2 different packages call same gcc standard library
    function pow. gcc linker appends pow implementation to the compiled
    package, and names that function "pow". But when these 2 compiled
    packages are linked together into the final executable, linker
    complains, because it finds two "pow" functions with the same name.
    
    This CL stops linking of mingwex and mingw32 during package build -
    that leaves pow function reference unresolved. pow reference gets
    resolved as final executable is built, by having both internal and
    external linker use mingwex and mingw32 libraries.
    
    Fixes #8756
    
    Change-Id: I50ddc79529ea5463c67118d668488345ecf069bc
    Reviewed-on: https://go-review.googlesource.com/26670
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b040bc9c062e4c5593792b6754d001509989a9df
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Aug 31 04:30:03 2016 +0000

    syscall: add some debugging to TestGetfsstat
    
    TestGetfsstat is failing on OS X 10.8.
    
    Not sure why. Add more debug info.
    
    Change-Id: I7dabb70dd7aeffda7e8959103db9e4886b84741e
    Reviewed-on: https://go-review.googlesource.com/28220
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 448d3952af2e60eab7fddd0d8a7b8e0ea6905073
Author: Kevin Burke <kev@inburke.com>
Date:   Tue Jul 12 16:54:36 2016 -0700

    crypto/sha256: add examples for New, Sum256
    
    The goal for these examples is to show how to mirror the
    functionality of the sha256sum Unix utility, a common checksumming
    tool, using the Go standard library.
    
    Add a newline at the end of the input, so users will get the same
    output if they type `echo 'hello world' | sha256sum`, since the
    builtin shell echo appends a newline by default. Also use hex output
    (instead of the shorter base64) since this is the default output
    encoding for shasum/sha256sum.
    
    Change-Id: I0036874b3cc5ba85432bfcb86f81b51c4e0238fd
    Reviewed-on: https://go-review.googlesource.com/24868
    Reviewed-by: Emmanuel Odeke <emm.odeke@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3968ac2caf4b7bbabcad6504ab5e3589fc525b1e
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Aug 31 09:54:00 2016 +1000

    cmd/compile/internal/gc: clean up closure.go
    
    Change-Id: I01bfab595c50582c5adf958dcecbd58524dbc28f
    Reviewed-on: https://go-review.googlesource.com/28212
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d3092464624f1f0ad29fa0ac4c4069fdd0697155
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Aug 30 16:08:06 2016 -0700

    compress/flate: always return uncompressed data in the event of error
    
    In the event of an unexpected error, we should always flush available
    decompressed data to the user.
    
    Fixes #16924
    
    Change-Id: I0bc0824c3201f3149e84e6a26e3dbcba72a1aae5
    Reviewed-on: https://go-review.googlesource.com/28216
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6ebacf18a5769d071f467fae455a142a964a43da
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 30 17:09:30 2016 -0700

    doc: more tweaks to the FAQ
    
    Change-Id: I0a3726f841122643bd1680ef6bd450c2039f362b
    Reviewed-on: https://go-review.googlesource.com/28213
    Reviewed-by: Rob Pike <r@golang.org>

commit 859cab099c5a9a9b4939960b630b78e468c8c39e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Aug 30 03:40:51 2016 +0000

    net/http: make DefaultTransport's Dialer enable DualStack ("Happy Eyeballs")
    
    As @pmarks-net said in the bug, this is something of a prisoner's
    dilemma, but it does help people who occasionally report problems.
    
    This is temporary. IPv6 is happening regardless of our decision here,
    so we'll do this for now.
    
    Fixes #15324
    
    Change-Id: I8cc29c6efa56222970996c71182fc9ee89d78539
    Reviewed-on: https://go-review.googlesource.com/28077
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 8e6948fe88f8300c8a939f062b9c58d688bbab62
Author: Bryan Alexander <Kozical@msn.com>
Date:   Sun Aug 28 14:11:31 2016 -0500

    crypto/x509: Fix bug in UnknownAuthorityError.Error
    
    Fix bug in UnknownAuthorityError.Error that would never allow Org
    Name to be inserted into error message if the Common Name was empty.
    Create tests for all three paths in UnknownAuthorityError.Error
    
    Change-Id: Id8afc444e897ef549df682d93a8563fd9de22a2b
    Reviewed-on: https://go-review.googlesource.com/27992
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2b74de3ed91c495d63868acef0471b0286e7b432
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Jun 28 09:22:46 2016 -0700

    runtime: rename fastrand1 to fastrand
    
    Change-Id: I37706ff0a3486827c5b072c95ad890ea87ede847
    Reviewed-on: https://go-review.googlesource.com/28210
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f9dafc742d7c0e892b6e4ff17cb9ec7165887e44
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Aug 30 14:46:25 2016 -0400

    cmd/compile, runtime, etc: get rid of constant FP registers
    
    On ARM64, MIPS64, and PPC64, some floating point registers were
    reserved for constants 0, 1, 2, 0.5, etc. This CL removes them.
    
    On ARM64, they are never used. On MIPS64 and PPC64, the only use
    case is a multiplication-by-2 in the old backend of the compiler,
    which is replaced with an addition.
    
    Change-Id: I737cbf43283756e3408964fc88c567a938c57036
    Reviewed-on: https://go-review.googlesource.com/28095
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b2e0e9688a512970ea8d270238c8ff3bbf85cbe1
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Aug 26 15:41:51 2016 -0400

    cmd/compile: remove Zero and NilCheck for newobject
    
    Recognize runtime.newobject and don't Zero or NilCheck it.
    
    Fixes #15914 (?)
    Updates #15390.
    
    TBD: add test
    
    Change-Id: Ia3bfa5c2ddbe2c27c92d9f68534a713b5ce95934
    Reviewed-on: https://go-review.googlesource.com/27930
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 842b05832fb5088a212e30962f58c95a38296d32
Author: Keith Randall <khr@golang.org>
Date:   Tue Aug 30 11:08:47 2016 -0700

    all: use testing.GoToolPath instead of "go"
    
    This change makes sure that tests are run with the correct
    version of the go tool.  The correct version is the one that
    we invoked with "go test", not the one that is first in our path.
    
    Fixes #16577
    
    Change-Id: If22c8f8c3ec9e7c35d094362873819f2fbb8559b
    Reviewed-on: https://go-review.googlesource.com/28089
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit be23e98e06b1e1c65de19d460537c4df21ebf555
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Aug 29 13:29:46 2016 -0700

    reflect: cleanup wording for type identity/equality
    
    Use terms like "equal" and "identical types" to match the Go spec,
    rather than inventing a new explanation. See also discussion on
    golang.org/cl/27170.
    
    Updates #16348.
    
    Change-Id: I0fe0bd01c0d1da3c8937a579c2ba44cf1eb16b71
    Reviewed-on: https://go-review.googlesource.com/28054
    Reviewed-by: Rob Pike <r@golang.org>

commit 55875977eb7e4f5b926127ec76217f37c7fd3713
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 30 14:48:01 2016 -0700

    cmd/compile: dedup Pragma switch
    
    Change-Id: I2d01f692ae30a166079976b86bf0b7a439f05d5c
    Reviewed-on: https://go-review.googlesource.com/28178
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d1383b5b8f94f7d94cf703b27f58e329979568d3
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Aug 30 14:18:39 2016 -0700

    misc/cgo/testsigfwd: add missing return statement
    
    Fixes C compiler warning:
    
    ./main.go:54:1: warning: control reaches end of non-void function [-Wreturn-type]
    
    Should help fix the linux builders
    that broke due to CL 23005.
    
    Change-Id: Ib0630798125e35a12f99d666b7ffe7b3196f0ecc
    Reviewed-on: https://go-review.googlesource.com/28176
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d57a4a656a8579b0ea3570c24329252ed536f98c
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 30 14:21:06 2016 -0700

    cmd/compile: eliminate addmethod tpkg parameter
    
    It's only needed for a check that can be pushed up into bimport.go,
    where it makes more sense anyway.
    
    Change-Id: I6ef381ff4f29627b0f390ce27fef08902932bea6
    Reviewed-on: https://go-review.googlesource.com/28177
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4710e16d34354a555cc6ebe54c6855bf5502fcbf
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 30 13:50:03 2016 -0700

    cmd/compile: cleanup addmethod
    
    Change-Id: Icb1671187d70edd962e2bda2cc45771b17a8e770
    Reviewed-on: https://go-review.googlesource.com/28175
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4eb2fa17659fb6a2337c72841b8695317813ad8f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 30 13:43:37 2016 -0700

    cmd/compile: eliminate methtype's mustname parameter
    
    Change-Id: Idd3e677dec00eb36a2cf7baa34e772335e1f2bc8
    Reviewed-on: https://go-review.googlesource.com/28173
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8c85e23087d90e831a70ccd199cac49a38d91027
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jun 17 16:27:23 2016 -0700

    cmd/compile: recognize integer ranges in switch statements
    
    Consider a switch statement like:
    
    switch x {
    case 1:
      // ...
    case 2, 3, 4, 5, 6:
      // ...
    case 5:
      // ...
    }
    
    Prior to this CL, the generated code treated
    2, 3, 4, 5, and 6 independently in a binary search.
    With this CL, the generated code checks whether
    2 <= x && x <= 6.
    walkinrange then optimizes that range check
    into a single unsigned comparison.
    
    Experiments suggest that the best min range size
    is 2, using binary size as a proxy for optimization.
    
    Binary sizes before/after this CL:
    
    cmd/compile: 14209728 / 14165360
    cmd/go:       9543100 /  9539004
    
    Change-Id: If2f7fb97ca80468fa70351ef540866200c4c996c
    Reviewed-on: https://go-review.googlesource.com/26770
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 5c3edc46a6014f9ff74a5e46a69a8891cad3190d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Aug 30 09:57:48 2016 -0700

    cmd/compile: add Trunc-of-Ext simplifications
    
    This is a follow-up to the discussion
    in CL 27853.
    
    During make.bash, trigger count:
    
    24 rewrite generic.rules:57
    22 rewrite generic.rules:69
    10 rewrite generic.rules:54
    10 rewrite generic.rules:58
    10 rewrite generic.rules:67
     7 rewrite generic.rules:66
     4 rewrite generic.rules:59
     3 rewrite generic.rules:50
     3 rewrite generic.rules:51
     3 rewrite generic.rules:52
     1 rewrite generic.rules:64
    
    Change-Id: Id96cb6a707a4a564831f763c2d4d0e180c94bbef
    Reviewed-on: https://go-review.googlesource.com/28088
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Martin Möhrmann <martisch@uos.de>

commit 74a00b249b049459b6f8b3c3999969f2beef3d31
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 30 11:15:38 2016 -0700

    cmd/compile: get rid of ugly {Recvs,Params,Results}P methods
    
    These were a hack abstraction for before FuncType existed.
    
    The result value from calling FuncType() could be saved, but this
    maintains the current idiom of consistently using t.FuncType().foo
    everywhere in case we choose to evolve the API further.
    
    Change-Id: I81f19aaeab6fb7caa2d4da8bf0bbbc358ab970d0
    Reviewed-on: https://go-review.googlesource.com/28150
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 7722d0f90383750784377bb395a8c799868bbab8
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Fri Aug 19 09:37:19 2016 +0900

    path/filepath: handle ".." in normalizing a path on Windows
    
    Current code assumes there are not ".." in the Clean(path).
    That's not true. Clean doesn't handle leading "..", so we need to stop
    normalization if we see "..".
    
    Fixes #16793
    
    Change-Id: I0a7901bedac17f1210b134d593ebd9f5e8483775
    Reviewed-on: https://go-review.googlesource.com/27410
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1319a0ffc79e6a3f278ce39bee90bf6823c647be
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 30 12:46:36 2016 -0700

    cmd/compile: remove unused FmtWidth flag
    
    Change-Id: I6c48683b620b0f119d7f0ae4a88502773202756b
    Reviewed-on: https://go-review.googlesource.com/28170
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6a982c390138fb2af49f85c4aeea2bce3222eb20
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 30 11:33:18 2016 -0700

    cmd/compile: remove unused Type.Printed field
    
    Change-Id: Iff2b1507dce08ef7c27085c8e0f45d0e3e88c476
    Reviewed-on: https://go-review.googlesource.com/28152
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2340f4713dff4ecd4ed082f80b1d38150e0b5348
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 30 11:19:07 2016 -0700

    cmd/compile: remove unused FmtBody flag and associated code
    
    For #15323.
    
    Change-Id: I23192a05ce57012aa2f96909d90d6a33b913766b
    Reviewed-on: https://go-review.googlesource.com/28151
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f42f20ad391c510ef394bc66cf3cf5bedef48e1e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Aug 30 04:08:10 2016 +0000

    net/http: fix ordering & data race in TestTransportEventTrace_h2
    
    Ordering fix: this CL swaps the order of the log write and the channel close
    in WroteRequest. I could reproduce the bug by putting a sleep between the two
    when the channel close was first. It needs to happen after the log.
    
    Data race: use the log buffer's mutex when reading too. Not really
    important once the ordering fix above is fixed (since nobody is
    concurrently writing anymore), but for consistency.
    
    Fixes #16414
    
    Change-Id: If6657884e67be90b4455c8f5a6f7bc6981999ee4
    Reviewed-on: https://go-review.googlesource.com/28078
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 0dae9dfb08a30983cce1114742c974077bdf5e18
Author: Martin Möhrmann <martisch@uos.de>
Date:   Fri Aug 26 15:00:46 2016 +0200

    cmd/compile: improve string iteration performance
    
    Generate a for loop for ranging over strings that only needs to call
    the runtime function charntorune for non ASCII characters.
    
    This provides faster iteration over ASCII characters and slightly
    faster iteration for other characters.
    
    The runtime function charntorune is changed to take an index from where
    to start decoding and returns the index after the last byte belonging
    to the decoded rune.
    
    All call sites of charntorune in the runtime are replaced by a for loop
    that will be transformed by the compiler instead of calling the charntorune
    function directly.
    
    go binary size decreases by 80 bytes.
    godoc binary size increases by around 4 kilobytes.
    
    runtime:
    
    name                           old time/op  new time/op  delta
    RuneIterate/range/ASCII-4      43.7ns ± 3%  10.3ns ± 4%  -76.33%  (p=0.000 n=44+45)
    RuneIterate/range/Japanese-4   72.5ns ± 2%  62.8ns ± 2%  -13.41%  (p=0.000 n=49+50)
    RuneIterate/range1/ASCII-4     43.5ns ± 2%  10.4ns ± 3%  -76.18%  (p=0.000 n=50+50)
    RuneIterate/range1/Japanese-4  72.5ns ± 2%  62.9ns ± 2%  -13.26%  (p=0.000 n=50+49)
    RuneIterate/range2/ASCII-4     43.5ns ± 3%  10.3ns ± 2%  -76.22%  (p=0.000 n=48+47)
    RuneIterate/range2/Japanese-4  72.4ns ± 2%  62.7ns ± 2%  -13.47%  (p=0.000 n=50+50)
    
    strings:
    
    name                 old time/op    new time/op    delta
    IndexRune-4            64.7ns ± 5%    22.4ns ± 3%  -65.43%  (p=0.000 n=25+21)
    MapNoChanges-4          269ns ± 2%     157ns ± 2%  -41.46%  (p=0.000 n=23+24)
    Fields-4               23.0ms ± 2%    19.7ms ± 2%  -14.35%  (p=0.000 n=25+25)
    FieldsFunc-4           23.1ms ± 2%    19.6ms ± 2%  -14.94%  (p=0.000 n=25+24)
    
    name                 old speed      new speed      delta
    Fields-4             45.6MB/s ± 2%  53.2MB/s ± 2%  +16.87%  (p=0.000 n=24+25)
    FieldsFunc-4         45.5MB/s ± 2%  53.5MB/s ± 2%  +17.57%  (p=0.000 n=25+24)
    
    Updates #13162
    
    Change-Id: I79ffaf828d82bf9887592f08e5cad883e9f39701
    Reviewed-on: https://go-review.googlesource.com/27853
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Martin Möhrmann <martisch@uos.de>

commit 0d7a2241cb684236f2728bb08514e7f256ac4a43
Author: Keith Randall <khr@golang.org>
Date:   Tue Aug 30 09:29:16 2016 -0700

    runtime: update a few comments
    
    noescape is now 0 instructions with the SSA backend.
    fast atomics are no longer a TODO (at least for amd64).
    
    Change-Id: Ib6e06f7471bef282a47ba236d8ce95404bb60a42
    Reviewed-on: https://go-review.googlesource.com/28087
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f7ac5da4956fb2db129848be331345ece8e7faa6
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Aug 29 17:56:15 2016 -0700

    cmd/compile: make internal objects directly print to printer
    
    Internal objects that satisfy the Printable interface can print
    directly to a printer w/o going through the conversion to a string
    first.
    
    Made printer.f understand and special-case %v so that Printable
    objects use the printer directly.
    
    This is work in progress and we may end up doing something else
    eventually (perhaps using fmt.Formatter) - or even undo these
    changes if this exploration doesn't get us to a significantly
    better place.
    
    Allocations numbers relative to commit c85b77c (still up, but
    reduced from most recent change):
    
    name       old time/op     new time/op     delta
    Template       307ms ± 4%      315ms ± 4%   +2.55%        (p=0.000 n=29+29)
    Unicode        164ms ± 4%      165ms ± 4%     ~           (p=0.057 n=30+30)
    GoTypes        1.01s ± 3%      1.03s ± 3%   +1.72%        (p=0.000 n=30+30)
    Compiler       5.49s ± 1%      5.62s ± 2%   +2.31%        (p=0.000 n=30+28)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        397M ± 3%       406M ± 6%   +2.21%        (p=0.000 n=28+30)
    Unicode         225M ± 4%       226M ± 3%     ~           (p=0.230 n=29+30)
    GoTypes        1.31G ± 3%      1.34G ± 5%   +2.79%        (p=0.000 n=30+30)
    Compiler       7.39G ± 2%      7.50G ± 2%   +1.43%        (p=0.000 n=30+29)
    
    name       old alloc/op    new alloc/op    delta
    Template      46.8MB ± 0%     47.5MB ± 0%   +1.48%        (p=0.000 n=29+28)
    Unicode       37.8MB ± 0%     38.1MB ± 0%   +0.64%        (p=0.000 n=30+28)
    GoTypes        143MB ± 0%      145MB ± 0%   +1.72%        (p=0.000 n=30+30)
    Compiler       683MB ± 0%      706MB ± 0%   +3.31%        (p=0.000 n=30+29)
    
    name       old allocs/op   new allocs/op   delta
    Template        444k ± 0%       481k ± 0%   +8.38%        (p=0.000 n=30+30)
    Unicode         369k ± 0%       379k ± 0%   +2.74%        (p=0.000 n=30+30)
    GoTypes        1.35M ± 0%      1.50M ± 0%  +10.78%        (p=0.000 n=30+30)
    Compiler       5.66M ± 0%      6.25M ± 0%  +10.31%        (p=0.000 n=29+29)
    
    For #16897.
    
    Change-Id: I37f95ab60508018ee6d29a98d238482b60e3e4b5
    Reviewed-on: https://go-review.googlesource.com/28072
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit fcb45e7ceff09548eb0308909256296086eedf9c
Author: Terrel Shumway <gopher@shumway.us>
Date:   Tue Aug 30 07:58:52 2016 -0600

    doc: clarify FAQ wording for float sizes
    
    I was confused by the current wording. This wording
    answers the question more clearly.
    
    Thanks to Robert Griesemer for suggestions.
    
    Fixes #16916
    
    Change-Id: I50187c8df2db661b9581f4b3c5d5c279d2f9af41
    Reviewed-on: https://go-review.googlesource.com/28052
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit d6098e4277bab633c2df752ed90e1e826918ca67
Author: Keith Randall <khr@golang.org>
Date:   Mon Aug 29 20:28:20 2016 -0700

    cmd/compile: intrinsify sync/atomic for amd64
    
    Uses the same implementation as runtime/internal/atomic.
    
    Reorganize the intrinsic detector to make it more table-driven.
    
    Also works on amd64p32.
    
    Change-Id: I7a5238951d6018d7d5d1bc01f339f6ee9282b2d0
    Reviewed-on: https://go-review.googlesource.com/28076
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit adb1e67f02fa58b13d1baf60c84556f375f6ceeb
Author: Dave Cheney <dave@cheney.net>
Date:   Mon Aug 29 09:44:50 2016 +1000

    reflect: avoid zeroing memory that will be overwritten
    
    Avoid new'ing memory that will be overwritten by assignment.
    
    name              old time/op    new time/op    delta
    Call-4               160ns ± 4%     155ns ± 2%  -3.19%        (p=0.003 n=10+10)
    FieldByName1-4      94.5ns ± 2%    95.2ns ± 1%  +0.65%          (p=0.026 n=9+9)
    FieldByName2-4      3.09µs ± 4%    3.13µs ± 2%    ~           (p=0.165 n=10+10)
    FieldByName3-4      19.8µs ± 1%    19.9µs ± 1%    ~            (p=0.395 n=10+8)
    InterfaceBig-4      11.6ns ± 0%    11.7ns ± 0%  +0.86%          (p=0.000 n=8+9)
    InterfaceSmall-4    11.7ns ± 0%    11.7ns ± 0%    ~     (all samples are equal)
    New-4               26.6ns ± 0%    26.4ns ± 0%  -0.64%         (p=0.000 n=10+9)
    
    name              old alloc/op   new alloc/op   delta
    Call-4              0.00B ±NaN%    0.00B ±NaN%    ~     (all samples are equal)
    
    name              old allocs/op  new allocs/op  delta
    Call-4               0.00 ±NaN%     0.00 ±NaN%    ~     (all samples are equal)
    
    Change-Id: I12c85d4e65245598669dd6f66beb0744ec9b9d6d
    Reviewed-on: https://go-review.googlesource.com/28011
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b6e3a98cf3a7acb3d5c1431eb04e9c3edad6c6ed
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue May 10 13:03:42 2016 -0700

    cmd/go: make C compiler warnings fatal on builders
    
    Fixes #14698
    
    Change-Id: I82fa781bf136c30e900d8e910ff576ba8b218acb
    Reviewed-on: https://go-review.googlesource.com/23005
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 2af00eb63cde716b59c0d64f4c3855b83a9d5a63
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Aug 30 01:09:32 2016 +0000

    net/http: stop timeout timer if user cancels a request
    
    Change-Id: I84faeae69f294b9a70e545faac6a070feba67770
    Reviewed-on: https://go-review.googlesource.com/28074
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 6f43a989b0701ef2f87827dc72ee93378e09bb9b
Author: Jonathan Boulle <jonathanboulle@gmail.com>
Date:   Tue Aug 30 02:09:19 2016 +0200

    os: fix typo in comment (Readir -> Readdir)
    
    Change-Id: I8434925661dc11396380af65c192c9f0dc191287
    Reviewed-on: https://go-review.googlesource.com/27973
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 298791a94af8b787c38fb95c51cb2dbc94668dad
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Aug 30 01:05:18 2016 +0000

    all: use time.Until where applicable
    
    Updates #14595
    
    Change-Id: Idf60b3004c7a0ebb59dd48389ab62c854069e09f
    Reviewed-on: https://go-review.googlesource.com/28073
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6c6ad08eb920d02947410919889229bbfa8c9915
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Aug 29 16:19:23 2016 -0700

    go/types: fix bad variable shadowing in offsetsof
    
    Introduced in CL 26995.
    
    Fixes #16902
    
    Change-Id: I8e749f598167e1f8b82cd5e735a7eb5291362e5e
    Reviewed-on: https://go-review.googlesource.com/28070
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 67ea710792eabdae1182e2bf4845f512136cccce
Author: Sam Whited <sam@samwhited.com>
Date:   Wed Mar 2 00:00:23 2016 -0600

    time: Add Until helper function
    
    Adds an Until() function that returns the duration until the given time.
    This compliments the existing Since() function and makes writing
    expressions that have expiration times more readable; for example:
    
        <-After(time.Until(connExpires)):
    
    Fixes #14595
    
    Change-Id: I87998a924b11d4dad5512e010b29d2da6b123456
    Reviewed-on: https://go-review.googlesource.com/20118
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 9f8335b7e72bdb2b95055c69f2d4b453b73646df
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Aug 29 21:27:40 2016 +0000

    os: don't let File.Readdir return an empty slice and nil error
    
    In the case of a file being deleted while Readdir was running, it was
    possible for File.Readdir to return an empty slice and a nil error,
    counter to its documentation.
    
    Fixes #16919
    
    Change-Id: If0e42882eea52fbf5530317a1895f3829ea8e67b
    Reviewed-on: https://go-review.googlesource.com/28056
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 11e3955e10ccd9105b78e07f4785402a5c8ceaf7
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Aug 29 13:53:32 2016 -0700

    net: restore per-query timeout logic
    
    The handling of "options timeout:n" is supposed to be per individual
    DNS server exchange, not per Lookup call.
    
    Fixes #16865.
    
    Change-Id: I2304579b9169c1515292f142cb372af9d37ff7c1
    Reviewed-on: https://go-review.googlesource.com/28057
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit aaa6b5352420aaaf9d46848facb5854927b3130e
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Fri Jul 29 14:02:26 2016 -0300

    runtime: insufficient padding in the `p` structure
    
    The current padding in the 'p' struct is hardcoded at 64 bytes. It should be the
    cache line size. On ppc64x, the current value is only okay because sys.CacheLineSize
    is wrong at 64 bytes. This change fixes that by making the padding equal to the
    cache line size. It also fixes the cache line size for ppc64/ppc64le to 128 bytes.
    
    Fixes #16477
    
    Change-Id: Ib7ec5195685116eb11ba312a064f41920373d4a3
    Reviewed-on: https://go-review.googlesource.com/25370
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit db1fef7b2afd12625e3edd90f879e5d05511d1d6
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun Jul 10 17:51:27 2016 -0700

    cmd/go: for -msan build runtime/cgo with -fsanitize=memory
    
    The go tool used to avoid passing -fsanitize=memory when building
    runtime/cgo. That was originally to avoid an msan error, but that error
    was fixed anyhow for issue #13815. And building runtime/cgo with
    -fsanitize=memory corrects the handling of the context traceback
    function when the traceback function itself is built with
    -fsanitize=memory.
    
    Change-Id: I4bf5c3d21de6b2eb540600435ae47f5820d17464
    Reviewed-on: https://go-review.googlesource.com/24855
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bd9aa9811d63c121a84ef92bc9b6a4101af46235
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Wed Jun 22 20:16:41 2016 +0300

    cmd/vet: check for copying of array of locks
    
    Updates #14664
    
    Change-Id: I1f7b1116cfe91466816c760f136ce566da3e80a9
    Reviewed-on: https://go-review.googlesource.com/24340
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit c9fbe0f29321602ce791834f600dcc453580c22b
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Thu Jun 23 14:59:26 2016 +0300

    cmd/vet: properly handle indexed arguments in printf
    
    Fixes #15884
    
    Change-Id: I33d98db861d74e3c37a546efaf83ce6f2f76d335
    Reviewed-on: https://go-review.googlesource.com/24391
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 4955147291812fd78049d47ef985095e6442264a
Author: Ethan Miller <eamiller@us.ibm.com>
Date:   Fri Aug 12 13:45:50 2016 -0500

    math/big: add assembly implementation of arith for ppc64{le}
    
    The existing implementation used a pure go implementation, leading to slow
    cryptographic performance.
    
    Implemented mulWW, subVV, mulAddVWW, addMulVVW, and bitLen for
    ppc64{le}.
    Implemented divWW for ppc64le only, as the DIVDEU instruction is only
    available on Power8 or newer.
    
    benchcmp output:
    
    benchmark                         old ns/op     new ns/op     delta
    BenchmarkSignP384                 28934360      10877330      -62.41%
    BenchmarkRSA2048Decrypt           41261033      5139930       -87.54%
    BenchmarkRSA2048Sign              45231300      7610985       -83.17%
    Benchmark3PrimeRSA2048Decrypt     20487300      2481408       -87.89%
    
    Fixes #16621
    
    Change-Id: If8b68963bb49909bde832f2bda08a3791c4f5b7a
    Reviewed-on: https://go-review.googlesource.com/26951
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit 0a7c73b5db9c78c4ecf8a5e8274ddf352b41562c
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Aug 29 13:53:43 2016 -0700

    cmd/compile: use printer in typefmt, Tconv
    
    Change-Id: Ib3ac0177761af1edea6b7951ffbbea042fb836d2
    Reviewed-on: https://go-review.googlesource.com/28055
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d3134b6450d08bba13e904f572a34d411614533e
Author: Kevin Burke <kev@inburke.com>
Date:   Thu Aug 18 09:14:22 2016 -0700

    cmd/compile: document more Node fields
    
    Change-Id: Ic8d37e5612b68bc73c4b50b59db54d8966b69838
    Reviewed-on: https://go-review.googlesource.com/27326
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e6f9f39ce52e880b54e4cb08bf0cde73cf6c7dc2
Author: Martin Möhrmann <martisch@uos.de>
Date:   Thu Aug 25 14:17:52 2016 +0200

    cmd/compile: generate makeslice calls with int arguments
    
    Where possible generate calls to runtime makeslice with int arguments
    during compile time instead of makeslice with int64 arguments.
    
    This eliminates converting arguments for calls to makeslice with
    int64 arguments for platforms where int64 values do not fit into
    arguments of type int.
    
    godoc 386 binary shrinks by approximately 12 kilobyte.
    
    amd64:
    name         old time/op  new time/op  delta
    MakeSlice-2  29.8ns ± 1%  29.8ns ± 1%   ~     (p=1.000 n=24+24)
    
    386:
    name         old time/op  new time/op  delta
    MakeSlice-2  52.3ns ± 0%  45.9ns ± 0%  -12.17%  (p=0.000 n=25+22)
    
    Fixes  #15357
    
    Change-Id: Icb8701bb63c5a83877d26c8a4b78e782ba76de7c
    Reviewed-on: https://go-review.googlesource.com/27851
    Run-TryBot: Martin Möhrmann <martisch@uos.de>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 595cebb055d327e52bd447985b53dcca869cea1d
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Aug 29 10:37:13 2016 -0700

    cmd/compile: remove ignored bool from exported ODCL nodes
    
    This shortens the export format by 1 byte for each exported ODCL
    node in inlined function bodies.
    
    Maintain backward compatibility by updating format version and
    continue to accept older format.
    
    Change-Id: I549bb3ade90bc0f146decf8016d5c9c3f14eb293
    Reviewed-on: https://go-review.googlesource.com/27999
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 428d79bd38db30405fe2b5d264d856e52030c338
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Sun Aug 28 23:11:00 2016 -0700

    os: add example for OpenFile
    
    New beginners are not familiar with open(2)-style masking of the
    flags. Add an example demonstrates the flag or'ing.
    
    Change-Id: Ifa8009c55173ba0dc6642c1d3b3124c766b1ebbb
    Reviewed-on: https://go-review.googlesource.com/27996
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 650c2c173dc8d929d1955045df4fe77c85d4037b
Author: Carlos C <uldericofilho@gmail.com>
Date:   Tue Aug 9 01:24:22 2016 +0200

    mime/quotedprintable: add examples
    
    Partially addresses #16360
    
    Change-Id: Ic098d2c465742fb50aee325a3fd0e1d50b7b3c99
    Reviewed-on: https://go-review.googlesource.com/25575
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 967fa427fde2f7a22adc04399d0de3090d34f883
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue May 31 06:58:33 2016 -0700

    database/sql: don't hang if the driver Exec method panics
    
    Fixes #13677.
    Fixes #15901.
    
    Change-Id: Idffb82cdcba4985954d061bdb021217f47ff4984
    Reviewed-on: https://go-review.googlesource.com/23576
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e6a96a6277da312d71c76ecb6f4a9a99ba88c389
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Aug 28 09:05:33 2016 -0700

    doc: add cmd/go pkgdir changes to go1.8.txt
    
    Change-Id: I451ca386781b50ab47b313e07a610867fa14aeaf
    Reviewed-on: https://go-review.googlesource.com/27990
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 26124030a051b6b05a29349fdc4b92880e18c8a8
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Aug 28 09:03:08 2016 -0700

    androidtest.bash: use go list to get pkg dir
    
    This will be more robust in the faces of
    future changes to the pkg dir layout.
    
    Change-Id: Iaf078093f02ef3a10884a19c25e2068cbbf5f36a
    Reviewed-on: https://go-review.googlesource.com/27929
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 6dd0b2d70292fc57e5664e2806f1d46cd621f531
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Aug 28 08:56:06 2016 -0700

    androidtest.bash: fix pkg dir
    
    CL 24930 altered the default InstallSuffix
    for mobile platforms.
    Update androidtest.bash to reflect this.
    This reverts CL 16151.
    
    A subsequent CL will make this more robust,
    but it will take more discussion and review.
    In the meantime, this fixes the build.
    
    Change-Id: Ia19ca2c9bab7b79c9cf24beeca64ecddaa60289c
    Reviewed-on: https://go-review.googlesource.com/27927
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Elias Naur <elias.naur@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 8dd069b52a87d482e219e16e92390b1d415db486
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Aug 28 08:53:20 2016 -0700

    androidtest.bash: require GOARCH
    
    When finding the pkg dir, androidtest.bash assumes
    that GOARCH is set. Require it up front.
    
    Change-Id: I143f7b59ad9d98b9c3cfb53c1d65c2d33a6acc12
    Reviewed-on: https://go-review.googlesource.com/27926
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Elias Naur <elias.naur@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 7c04633e0c4babc8da4c8a582aee4f1754c9db12
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Aug 28 17:04:46 2016 -0700

    all: fix obsolete inferno-os links
    
    Fixes #16911.
    
    Fix obsolete inferno-os links, since code.google.com shutdown.
    This CL points to the right files by replacing
    http://code.google.com/p/inferno-os/source/browse
    with
    https://bitbucket.org/inferno-os/inferno-os/src/default
    
    To implement the change I wrote and ran this script in the root:
    $ grep -Rn 'http://code.google.com/p/inferno-os/source/browse' * \
    | cut -d":" -f1 | while read F;do perl -pi -e \
    's/http:\/\/code.google.com\/p\/inferno-os\/source\/browse/https:\/\/bitbucket.org\/inferno-os\/inferno-os\/src\/default/g'
    $F;done
    
    I excluded any cmd/vendor changes from the commit.
    
    Change-Id: Iaaf828ac8f6fc949019fd01832989d00b29b6749
    Reviewed-on: https://go-review.googlesource.com/27994
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9ac67f56f776ccf1708d25ce3f4f02ac771551de
Author: Giovanni Bajo <rasky@develer.com>
Date:   Wed Aug 17 17:31:12 2016 +0200

    doc: improve issue template
    
    The previous template used an ordered list, but the formatting always
    breaks when users paste quoted snippets of code or command outputs.
    It is also harder to visually parse because items in ordered lists
    are only indented but not highlighted in any way.
    
    Change-Id: I73c89e9f0465aef41093f5c54d11bb0d12ff8c8d
    Reviewed-on: https://go-review.googlesource.com/27252
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2f679d74e638ac7514a6b6b32f5d28a9980c22c3
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sat Aug 27 19:08:14 2016 -0400

    cmd/compile: fix load int32 to FP register on big-endian MIPS64
    
    Fixes #16903.
    
    Change-Id: I1f6fcd57e14b2b62e208b7bb3adccd5fd7f8bdbc
    Reviewed-on: https://go-review.googlesource.com/27933
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 8c15a1725147692e1106f2e32fae657e1b7f27aa
Author: Radu Berinde <radu@cockroachlabs.com>
Date:   Sun Aug 28 11:47:14 2016 -0400

    hash/crc32: fix nil Castagnoli table problem
    
    When SSE is available, we don't need the Table. However, it is
    returned as a handle by MakeTable. Fix this to always generate
    the table.
    
    Further cleanup is discussed in #16909.
    
    Change-Id: Ic05400d68c6b5d25073ebd962000451746137afc
    Reviewed-on: https://go-review.googlesource.com/27934
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0c6c3d1de7bfe40d3589109bf2adb1726d6caca8
Author: Keith Randall <khr@golang.org>
Date:   Sun Aug 28 11:17:37 2016 -0700

    cmd/compile: fix noopt build
    
    Atomic add rules were depending on CSE to combine duplicate atomic ops.
    With -N, CSE doesn't run.
    
    Redo the rules for atomic add so there's only one atomic op.
    Introduce an add-to-first-part-of-tuple pseudo-ops to make the atomic add result correct.
    
    Change-Id: Ib132247051abe5f80fefad6c197db8df8ee06427
    Reviewed-on: https://go-review.googlesource.com/27991
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 84aac622a40d4707513b02cf056977d852b38e18
Author: Keith Randall <khr@golang.org>
Date:   Thu Aug 25 16:02:57 2016 -0700

    cmd/compile: intrinsify the rest of runtime/internal/atomic for amd64
    
    Atomic swap, add/and/or, compare and swap.
    
    Also works on amd64p32.
    
    Change-Id: Idf2d8f3e1255f71deba759e6e75e293afe4ab2ba
    Reviewed-on: https://go-review.googlesource.com/27813
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit e2e2d10b9afe0cf4845aede06c6dd083a1d281c9
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Fri Aug 26 16:57:00 2016 +1000

    time: always use $GOROOT/lib/time/zoneinfo.zip with genzabbrs.go
    
    genzabbrs.go uses whatever zoneinfo database available on the system.
    This makes genzabbrs.go output change from system to system. Adjust
    go:generate line to always use $GOROOT/lib/time/zoneinfo.zip, so it
    does not matter who runs the command.
    
    Also move go:generate line into zoneinfo.go, so it can be run
    on Unix (see #16368 for details).
    
    Fixes #15802.
    
    Change-Id: I8ae4818aaf40795364e180d7bb4326ad7c07c370
    Reviewed-on: https://go-review.googlesource.com/27832
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 90c3cf4b52cc9373a96009da0d013019c1f5bcd8
Author: Radu Berinde <radu@cockroachlabs.com>
Date:   Sat Aug 27 13:17:30 2016 -0400

    hash/crc32: improve the AMD64 implementation using SSE4.2
    
    The algorithm is explained in the comments. The improvement in
    throughput is about 1.4x for buffers between 500b-4Kb and 2.5x-2.6x
    for larger buffers.
    
    Additionally, we no longer initialize the software tables if SSE4.2 is
    available.
    
    Adding a test for the SSE implementation (restricted to amd64 and
    amd64p32).
    
    Benchmarks on a Haswell i5-4670 @ 3.4 GHz:
    
    name                           old time/op    new time/op     delta
    CastagnoliCrc15B-4               21.9ns ± 1%     22.9ns ± 0%    +4.45%
    CastagnoliCrc15BMisaligned-4     22.6ns ± 0%     23.4ns ± 0%    +3.43%
    CastagnoliCrc40B-4               23.3ns ± 0%     23.9ns ± 0%    +2.58%
    CastagnoliCrc40BMisaligned-4     25.4ns ± 0%     26.1ns ± 0%    +2.86%
    CastagnoliCrc512-4               72.6ns ± 0%     52.8ns ± 0%   -27.33%
    CastagnoliCrc512Misaligned-4     76.3ns ± 1%     56.3ns ± 0%   -26.18%
    CastagnoliCrc1KB-4                128ns ± 1%       89ns ± 0%   -30.04%
    CastagnoliCrc1KBMisaligned-4      130ns ± 0%       88ns ± 0%   -32.65%
    CastagnoliCrc4KB-4                461ns ± 0%      187ns ± 0%   -59.40%
    CastagnoliCrc4KBMisaligned-4      463ns ± 0%      191ns ± 0%   -58.77%
    CastagnoliCrc32KB-4              3.58µs ± 0%     1.35µs ± 0%   -62.22%
    CastagnoliCrc32KBMisaligned-4    3.58µs ± 0%     1.36µs ± 0%   -61.84%
    
    name                           old speed      new speed       delta
    CastagnoliCrc15B-4              684MB/s ± 1%    655MB/s ± 0%    -4.32%
    CastagnoliCrc15BMisaligned-4    663MB/s ± 0%    641MB/s ± 0%    -3.32%
    CastagnoliCrc40B-4             1.72GB/s ± 0%   1.67GB/s ± 0%    -2.69%
    CastagnoliCrc40BMisaligned-4   1.58GB/s ± 0%   1.53GB/s ± 0%    -2.82%
    CastagnoliCrc512-4             7.05GB/s ± 0%   9.70GB/s ± 0%   +37.59%
    CastagnoliCrc512Misaligned-4   6.71GB/s ± 1%   9.09GB/s ± 0%   +35.43%
    CastagnoliCrc1KB-4             7.98GB/s ± 1%  11.46GB/s ± 0%   +43.55%
    CastagnoliCrc1KBMisaligned-4   7.86GB/s ± 0%  11.70GB/s ± 0%   +48.75%
    CastagnoliCrc4KB-4             8.87GB/s ± 0%  21.80GB/s ± 0%  +145.69%
    CastagnoliCrc4KBMisaligned-4   8.83GB/s ± 0%  21.39GB/s ± 0%  +142.25%
    CastagnoliCrc32KB-4            9.15GB/s ± 0%  24.22GB/s ± 0%  +164.62%
    CastagnoliCrc32KBMisaligned-4  9.16GB/s ± 0%  24.00GB/s ± 0%  +161.94%
    
    Fixes #16107.
    
    Change-Id: Ibe50ea76574674ce0571ef31c31015e0ed66b907
    Reviewed-on: https://go-review.googlesource.com/27931
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 320bd562cbb24a01beb02706c42d06a290160645
Author: Joonas Kuorilehto <joneskoo@derbian.fi>
Date:   Sat Aug 20 14:41:42 2016 +0300

    crypto/tls: add KeyLogWriter for debugging
    
    Add support for writing TLS client random and master secret
    in NSS key log format.
    
    https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format
    
    Normally this is enabled by a developer debugging TLS based
    applications, especially HTTP/2, by setting the KeyLogWriter
    to an open file. The keys negotiated in handshake are then
    logged and can be used to decrypt TLS sessions e.g. in Wireshark.
    
    Applications may choose to add support similar to NSS where this
    is enabled by environment variable, but no such mechanism is
    built in to Go. Instead each application must explicitly enable.
    
    Fixes #13057.
    
    Change-Id: If6edd2d58999903e8390b1674ba4257ecc747ae1
    Reviewed-on: https://go-review.googlesource.com/27434
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3427f16642a1c207db4a4c3cce912dfdce2ac9f5
Author: Keith Randall <khr@golang.org>
Date:   Sat Aug 27 16:48:22 2016 +0000

    Revert "hash/crc32: improve the AMD64 implementation using SSE4.2"
    
    This reverts commit 54d7de7dd62bab764125c48fd159bb938da078e1.
    
    It was breaking non-amd64 builds.
    
    Change-Id: I22650e922498eeeba3d4fa08bb4ea40a210c8f97
    Reviewed-on: https://go-review.googlesource.com/27925
    Reviewed-by: Keith Randall <khr@golang.org>

commit 54d7de7dd62bab764125c48fd159bb938da078e1
Author: Radu Berinde <radu@cockroachlabs.com>
Date:   Tue Aug 23 17:33:21 2016 -0400

    hash/crc32: improve the AMD64 implementation using SSE4.2
    
    The algorithm is explained in the comments. The improvement in
    throughput is about 1.4x for buffers between 500b-4Kb and 2.5x-2.6x
    for larger buffers.
    
    Additionally, we no longer initialize the software tables if SSE4.2 is
    available.
    
    Benchmarks on a Haswell i5-4670 @ 3.4 GHz:
    
    name                           old time/op    new time/op     delta
    CastagnoliCrc15B-4               21.9ns ± 1%     22.9ns ± 0%    +4.45%
    CastagnoliCrc15BMisaligned-4     22.6ns ± 0%     23.4ns ± 0%    +3.43%
    CastagnoliCrc40B-4               23.3ns ± 0%     23.9ns ± 0%    +2.58%
    CastagnoliCrc40BMisaligned-4     25.4ns ± 0%     26.1ns ± 0%    +2.86%
    CastagnoliCrc512-4               72.6ns ± 0%     52.8ns ± 0%   -27.33%
    CastagnoliCrc512Misaligned-4     76.3ns ± 1%     56.3ns ± 0%   -26.18%
    CastagnoliCrc1KB-4                128ns ± 1%       89ns ± 0%   -30.04%
    CastagnoliCrc1KBMisaligned-4      130ns ± 0%       88ns ± 0%   -32.65%
    CastagnoliCrc4KB-4                461ns ± 0%      187ns ± 0%   -59.40%
    CastagnoliCrc4KBMisaligned-4      463ns ± 0%      191ns ± 0%   -58.77%
    CastagnoliCrc32KB-4              3.58µs ± 0%     1.35µs ± 0%   -62.22%
    CastagnoliCrc32KBMisaligned-4    3.58µs ± 0%     1.36µs ± 0%   -61.84%
    
    name                           old speed      new speed       delta
    CastagnoliCrc15B-4              684MB/s ± 1%    655MB/s ± 0%    -4.32%
    CastagnoliCrc15BMisaligned-4    663MB/s ± 0%    641MB/s ± 0%    -3.32%
    CastagnoliCrc40B-4             1.72GB/s ± 0%   1.67GB/s ± 0%    -2.69%
    CastagnoliCrc40BMisaligned-4   1.58GB/s ± 0%   1.53GB/s ± 0%    -2.82%
    CastagnoliCrc512-4             7.05GB/s ± 0%   9.70GB/s ± 0%   +37.59%
    CastagnoliCrc512Misaligned-4   6.71GB/s ± 1%   9.09GB/s ± 0%   +35.43%
    CastagnoliCrc1KB-4             7.98GB/s ± 1%  11.46GB/s ± 0%   +43.55%
    CastagnoliCrc1KBMisaligned-4   7.86GB/s ± 0%  11.70GB/s ± 0%   +48.75%
    CastagnoliCrc4KB-4             8.87GB/s ± 0%  21.80GB/s ± 0%  +145.69%
    CastagnoliCrc4KBMisaligned-4   8.83GB/s ± 0%  21.39GB/s ± 0%  +142.25%
    CastagnoliCrc32KB-4            9.15GB/s ± 0%  24.22GB/s ± 0%  +164.62%
    CastagnoliCrc32KBMisaligned-4  9.16GB/s ± 0%  24.00GB/s ± 0%  +161.94%
    
    Fixes #16107.
    
    Change-Id: I8fa827ec03f708ba27ee71c833f7544ad9dc5bc3
    Reviewed-on: https://go-review.googlesource.com/24471
    Reviewed-by: Keith Randall <khr@golang.org>

commit 0d23c28526223f5239581e845e0682f704724525
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Aug 26 16:38:06 2016 -0700

    cmd/compile: make dumpdepth a global again
    
    Fixes indenting in debug output like -W.
    
    Change-Id: Ia16b0bad47428cee71fe036c297731e841ec9ca0
    Reviewed-on: https://go-review.googlesource.com/27924
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 57331b79da3920ba579a07a20827ee4499cf1c7e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Aug 26 23:39:50 2016 +0000

    Revert "cmd/compile: use printer in typefmt, Tconv"
    
    This reverts commit 8fb0893307b0d4ab4a115a6151ee8d344d3c1d74.
    
    Broke go/ast tests.
    
    Change-Id: I5c314cb29731d4bc3a0873af8ebfe376f5faba8a
    Reviewed-on: https://go-review.googlesource.com/27923
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8fb0893307b0d4ab4a115a6151ee8d344d3c1d74
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 24 23:53:16 2016 -0700

    cmd/compile: use printer in typefmt, Tconv
    
    Change-Id: I9e99289070d63a2509aec1e91b9dd7437a08af5e
    Reviewed-on: https://go-review.googlesource.com/27921
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 71d2b42bf6c7ccd6a92ad10e8d0ac9272e312d1a
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 24 23:44:55 2016 -0700

    cmd/compile: use printer in exprfmt
    
    Change-Id: I7376c3bb55529a575e74985c2d7f0cf07c8996e7
    Reviewed-on: https://go-review.googlesource.com/27920
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a46ea907053d6ad116e49f76f58b69633c6039d3
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 14 14:07:18 2016 -0700

    go/build: don't alter InstallSuffix for default compile options
    
    Fixes #16378.
    
    Change-Id: I99a064f1afec78fb63cb3719061d20be0f21d45d
    Reviewed-on: https://go-review.googlesource.com/24930
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 87061054e709f5408db4bc1b9ac8e5685883d9fa
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 24 23:26:07 2016 -0700

    cmd/compile: use printer in stmtfmt, hconv
    
    Change-Id: If11d328101a82de5ead04159d3085e3d59869283
    Reviewed-on: https://go-review.googlesource.com/27919
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 62a296cd5482312cc656a4ead421acef00538cc1
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 24 23:02:08 2016 -0700

    cmd/compile: use printer in sconv, symfmt
    
    Change-Id: Iec33775ff5a786f6c52024d592f634231acf91c0
    Reviewed-on: https://go-review.googlesource.com/27918
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 73be5d82d443def1f721474bff9819155eb5bdf0
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 24 22:15:07 2016 -0700

    cmd/compile: use printer in vconv
    
    Change-Id: Ib30ed686448c4c0a5777cdf1d505ea06eb8b2a47
    Reviewed-on: https://go-review.googlesource.com/27917
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit dc2a0d59a2b9b5dd87065e6227ca112839cd389b
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 24 21:47:58 2016 -0700

    cmd/compile: introduce printer for internal formatting; use in jconv
    
    Starting point for uniform use of printer in fmt.go.
    It provides a hook to store additional state (and
    remove global variables) and should also be more
    efficient and cleaner than the mix of string concatenation
    and bytes.Buffer use we have now.
    
    Change-Id: I72de14b01850cca32d407a1cb16c894179ea8848
    Reviewed-on: https://go-review.googlesource.com/27916
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c85b77c22b78c801d18a456a8f242c007a520217
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Aug 26 13:06:59 2016 -0700

    cmd/compile: reintroduce support for version 0 export format
    
    The Go1.7 export format didn't encode the field package for
    blank struct fields (#15514). Re-introduce support for that
    format so we can read it w/o error.
    
    For #16881.
    
    Change-Id: Ib131d41aac56dbf970aab15ae7e75ef3944b412d
    Reviewed-on: https://go-review.googlesource.com/27912
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit fba8f4deba81b8c5d903ec2f52dcb151f13a147b
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Aug 25 17:17:50 2016 -0700

    go/internal/gcimporter: fail gracefully on export format skew
    
    Port of changes made to compiler in
    https://go-review.googlesource.com/27814.
    
    Correctly handle export format version 0 (we only do this
    in x/tools/gcimporter15 at the moment - this is a backport
    of that code for struct fields).
    
    Added tests for version handling and detection of corrupted
    export data.
    
    Fixes #16881.
    
    Change-Id: I246553c689c89ef5c7fedd1e43717504c2838804
    Reviewed-on: https://go-review.googlesource.com/27816
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit bee42067649390557c772c561e02d46b3f066fe3
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Aug 25 22:19:16 2016 -0400

    runtime: have typelinksinit work forwards
    
    For reasons I have forgotten typelinksinit processed modules backwards.
    (I suspect this was an attempt to process types in the executing
    binary first.)
    
    It does not appear to be necessary, and it is not the order we want
    when a module can be loaded at an arbitrary point during a program's
    execution as a plugin. So reverse the order.
    
    While here, make it safe to call typelinksinit multiple times.
    
    Change-Id: Ie10587c55c8e5efa0542981efb6eb3c12dd59e8c
    Reviewed-on: https://go-review.googlesource.com/27822
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8f3c8a33fad917abb45ef98b3a1cd34fe9715370
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Aug 25 21:06:10 2016 -0400

    cmd/link: make DynlinkingGo a method
    
    This will allow it to depend on whether plugin.Open is a symbol to be
    linked in.
    
    Change-Id: Ie9aa4216f2510fe8b10bc4665c8b19622b7122ea
    Reviewed-on: https://go-review.googlesource.com/27819
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7f27f1dfdd81c978d4868917d7622e09b288ecb0
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Aug 25 15:52:45 2016 -0400

    cmd/compile: add MIPS64 optimizations, SSA on by default
    
    Add the following optimizations:
    - fold constants
    - fold address into load/store
    - simplify extensions and conditional branches
    - remove nil checks
    
    Turn on SSA on MIPS64 by default, and toggle the tests.
    
    Fixes #16359.
    
    Change-Id: I7f1e38c2509e22e42cd024e712990ebbe47176bd
    Reviewed-on: https://go-review.googlesource.com/27870
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9f7ea616742f8d58ecc140afcb0fa53bcb84fe9b
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Aug 26 11:39:47 2016 -0400

    math: optimize Ceil, Floor and Trunc on s390x
    
    Use the FIDBR instruction to round floating-point numbers to integers.
    
    name   old time/op  new time/op  delta
    Ceil   14.1ns ± 0%   3.0ns ± 0%  -78.89%  (p=0.000 n=10+10)
    Floor  6.42ns ± 0%  3.03ns ± 0%  -52.80%  (p=0.000 n=10+10)
    Trunc  6.67ns ± 0%  3.03ns ± 0%  -54.57%   (p=0.000 n=10+9)
    
    Change-Id: I3b416f6d0bccaaa9b547de86356471365862399c
    Reviewed-on: https://go-review.googlesource.com/27827
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d2dd0dfda813deb27415962656e7a6496170cef3
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Aug 26 11:28:41 2016 -0400

    cmd/internal/obj/s390x: add FIDBR and FIEBR instructions
    
    FIDBR and FIEBR can be used for floating-point to integer rounding.
    The relevant functions (Ceil, Floor and Trunc) will be updated
    in a future CL.
    
    Change-Id: I5952d67ab29d5ef8923ff1143e17a8d30169d692
    Reviewed-on: https://go-review.googlesource.com/27826
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 266b349b2d28bf69f778320adb7e8ecc6bf848cd
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Aug 26 10:33:34 2016 -0400

    cmd/internal/obj/s390x: add atomic operation instructions
    
    Adds the following s390x instructions from the interlocked access
    facility:
    
     * LAA(G)  - load and add
     * LAAL(G) - load and add logical
     * LAN(G)  - load and and
     * LAX(G)  - load and exclusive or
     * LAO(G)  - load and or
    
    These instructions can be used for atomic arithmetic/logical
    operations. The atomic packages will be updated in future CLs.
    
    Change-Id: Idc850ac6749b3e778fda3da66bcd864f6b1df375
    Reviewed-on: https://go-review.googlesource.com/27871
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 14efaa0dc3d3ff5a3919c27297297ef0cd5bb625
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Aug 25 14:31:50 2016 -0400

    cmd/compile: qualify unexported fields of unnamed types
    
    The compiler was canonicalizing unnamed types of the form
    
            struct { i int }
    
    across packages, even though an unexported field i should not be
    accessible from other packages.
    
    The fix requires both qualifying the field name in the string used by
    the compiler to distinguish the type, and ensuring the struct's pkgpath
    is set in the rtype version of the data when the type being written is
    not part of the localpkg.
    
    Fixes #16616
    
    Change-Id: Ibab160b8b5936dfa47b17dbfd48964a65586785b
    Reviewed-on: https://go-review.googlesource.com/27791
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 3e59b20d41c6dc4ed1e528279da3017555df2ceb
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Fri Aug 26 17:44:30 2016 +0900

    os: fix build error on plan9
    
    https://go-review.googlesource.com/#/c/27580 added the test.
    However the test use syscall.ELOOP which is not defined on plan9.
    Move test code from "os_test.go" to "os_windows_test.go" to prevent
    build error.
    
    Change-Id: Ie7f05bfb9ab229e06a8e82a4b3b8a7ca82d4663b
    Reviewed-on: https://go-review.googlesource.com/27833
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David du Colombier <0intro@gmail.com>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit a656390b6986c3d559873b64683aa3454a151115
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Aug 25 21:01:17 2016 -0700

    spec: clarify text on init functions
    
    For #16874.
    
    Change-Id: I2e13f582297606e506d805755a6cfc1f3d4306a2
    Reviewed-on: https://go-review.googlesource.com/27817
    Reviewed-by: Rob Pike <r@golang.org>

commit 2eb46e8c57c4dab0197ca82d9899fa1356500fc0
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Thu Aug 18 16:10:28 2016 +0900

    os: prevent infinite symlink loop of Stat on Windows
    
    The Windows version of Stat calls Readlink iteratively until
    reaching a non-symlink file.
    If the given file is a circular symlink, It never stops.
    This CL defines the maximum number of symlink loop count.
    If the loop count will exceed that number, Stat will return error.
    
    Fixes #16538
    
    Change-Id: Ia9f3f2259a8d32801461c5041cc24a34f9f81009
    Reviewed-on: https://go-review.googlesource.com/27580
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 5a6f973565ccb54c77e03cbb4844fd0ea392d3fe
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Aug 25 16:53:10 2016 -0700

    cmd/compile: fail gracefully on export format skew
    
    Import errors due to unexpected format are virtually
    always due to version skew. Don't panic but report a
    good error message (incl. hint that the imported package
    needs to be reinstalled) if not in debugFormat mode.
    
    Recognize export data format version and store it so
    it can be used to automatically handle minor version
    differences. We did this before, but not very well.
    
    No export data format changes.
    
    Manually tested with corrupted export data.
    
    For #16881.
    
    Change-Id: I53ba98ef747b1c81033a914bb61ee52991f35a90
    Reviewed-on: https://go-review.googlesource.com/27814
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 7c3fc4b8e48f061e242e8d04803c7bb249ba6995
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Aug 26 12:15:00 2016 +1200

    cmd/link: hide funcsym symbols
    
    As far as I can tell, this check has been
    non-functional since it was introduced.
    
    This cuts 57k off cmd/go and 70k off cmd/compile.
    
    Based on golang.org/cl/24710 by Josh Bleecher Snyder.
    
    Change-Id: I1162a066971df1a067b50afa1cfa0819a6913574
    Reviewed-on: https://go-review.googlesource.com/27830
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit e022dcd35f3047134c661ef27cafdf1d845df447
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Aug 25 17:19:28 2016 -0700

    syscall: fix plan9/386 RawSyscall6
    
    Fixes the build.
    
    Change-Id: I34bcae08cfb43257aeb9086336966ef85f15fe1d
    Reviewed-on: https://go-review.googlesource.com/27815
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit f9acd3918e5eb8819f8f5e9697af55f395a1074d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Aug 25 13:42:49 2016 -0700

    net/http, cmd/compile: minor vet fixes
    
    Updates #11041
    
    Change-Id: Ia0151723e3bc0d163cc687a02bfc5e0285d95ffa
    Reviewed-on: https://go-review.googlesource.com/27810
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 320ddcf8344beb1c322f3a7f0a251eea5e442a10
Author: Keith Randall <khr@golang.org>
Date:   Tue Aug 23 16:49:28 2016 -0700

    cmd/compile: inline atomics from runtime/internal/atomic on amd64
    
    Inline atomic reads and writes on amd64.  There's no reason
    to pay the overhead of a call for these.
    
    To keep atomic loads from being reordered, we make them
    return a <value,memory> tuple.
    
    Change the meaning of resultInArg0 for tuple-generating ops
    to mean the first part of the result tuple, not the second.
    This means we can always put the store part of the tuple last,
    matching how arguments are laid out.  This requires reordering
    the outputs of add32carry and sub32carry and their descendents
    in various architectures.
    
    benchmark                    old ns/op     new ns/op     delta
    BenchmarkAtomicLoad64-8      2.09          0.26          -87.56%
    BenchmarkAtomicStore64-8     7.54          5.72          -24.14%
    
    TBD (in a different CL): Cas, Or8, ...
    
    Change-Id: I713ea88e7da3026c44ea5bdb56ed094b20bc5207
    Reviewed-on: https://go-review.googlesource.com/27641
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 71ab9fa312f8266379dbb358b9ee9303cde7bd6b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jul 11 16:05:57 2016 -0700

    all: fix assembly vet issues
    
    Add missing function prototypes.
    Fix function prototypes.
    Use FP references instead of SP references.
    Fix variable names.
    Update comments.
    Clean up whitespace. (Not for vet.)
    
    All fairly minor fixes to make vet happy.
    
    Updates #11041
    
    Change-Id: Ifab2cdf235ff61cdc226ab1d84b8467b5ac9446c
    Reviewed-on: https://go-review.googlesource.com/27713
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6af7639ae147689cbabd06287bf4ff15a4dfd896
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sat Aug 20 01:46:32 2016 -0700

    archive/tar: isolate regular and sparse file handling as methods
    
    Factor out the regular file handling logic into handleRegularFile
    from nextHeader. We will need to reuse this logic when fixing #15573
    in a future CL.
    
    Factor out the sparse file handling logic into handleSparseFile.
    Currently this logic is split between nextHeader (for GNU sparse
    files) and Next (for PAX sparse files). Instead, we move this
    related code into a single method.
    
    There is no overall logic change. Thus, no unit tests.
    
    Updates #15573 #15564
    
    Change-Id: I3b8270d8b4e080e77d6c0df6a123d677c82cc466
    Reviewed-on: https://go-review.googlesource.com/27454
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit aa9b3d70142afb75a510c2c92b8c387fce10b2c9
Author: Sina Siadat <siadat@gmail.com>
Date:   Fri Jun 17 21:02:59 2016 +0430

    net/http: send Content-Range if no byte range overlaps
    
    RFC 7233, section 4.4 says:
    >>>
    For byte ranges, failing to overlap the current extent means that the
    first-byte-pos of all of the byte-range-spec values were greater than the
    current length of the selected representation.  When this status code is
    generated in response to a byte-range request, the sender SHOULD generate a
    Content-Range header field specifying the current length of the selected
    representation
    <<<
    
    Thus, we should send the Content-Range only if none of the ranges
    overlap.
    
    Fixes #15798.
    
    Change-Id: Ic9a3e1b3a8730398b4bdff877a8f2fd2e30149e3
    Reviewed-on: https://go-review.googlesource.com/24212
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0bc94a8864f0cb8392c094f58dd84d28a06f35d5
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Apr 15 15:49:00 2016 -0700

    cmd/compile: when inlining ==, don’t take the address of the values
    
    This CL reworks walkcompare for clarity and concision.
    It also makes one significant functional change.
    (The functional change is hard to separate cleanly
    from the cleanup, so I just did them together.)
    When inlining and unrolling an equality comparison
    for a small struct or array, compare the elements like:
    
    a[0] == b[0] && a[1] == b[1]
    
    rather than
    
    pa := &a
    pb := &b
    pa[0] == pb[0] && pa[1] == pb[1]
    
    The result is the same, but taking the address
    and working through the indirect
    forces the backends to generate less efficient code.
    
    This is only an improvement with the SSA backend.
    However, every port but s390x now has a working
    SSA backend, and switching to the SSA backend
    by default everywhere is a priority for Go 1.8.
    It thus seems reasonable to start to prioritize
    SSA performance over the old backend.
    
    Updates #15303
    
    
    Sample code:
    
    type T struct {
            a, b int8
    }
    
    func g(a T) bool {
            return a == T{1, 2}
    }
    
    
    SSA before:
    
    "".g t=1 size=80 args=0x10 locals=0x8
            0x0000 00000 (badeq.go:7)       TEXT    "".g(SB), $8-16
            0x0000 00000 (badeq.go:7)       SUBQ    $8, SP
            0x0004 00004 (badeq.go:7)       FUNCDATA        $0, gclocals·23e8278e2b69a3a75fa59b23c49ed6ad(SB)
            0x0004 00004 (badeq.go:7)       FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
            0x0004 00004 (badeq.go:8)       MOVBLZX "".a+16(FP), AX
            0x0009 00009 (badeq.go:8)       MOVB    AL, "".autotmp_0+6(SP)
            0x000d 00013 (badeq.go:8)       MOVBLZX "".a+17(FP), AX
            0x0012 00018 (badeq.go:8)       MOVB    AL, "".autotmp_0+7(SP)
            0x0016 00022 (badeq.go:8)       MOVB    $0, "".autotmp_1+4(SP)
            0x001b 00027 (badeq.go:8)       MOVB    $1, "".autotmp_1+4(SP)
            0x0020 00032 (badeq.go:8)       MOVB    $2, "".autotmp_1+5(SP)
            0x0025 00037 (badeq.go:8)       MOVBLZX "".autotmp_0+6(SP), AX
            0x002a 00042 (badeq.go:8)       MOVBLZX "".autotmp_1+4(SP), CX
            0x002f 00047 (badeq.go:8)       CMPB    AL, CL
            0x0031 00049 (badeq.go:8)       JNE     70
            0x0033 00051 (badeq.go:8)       MOVBLZX "".autotmp_0+7(SP), AX
            0x0038 00056 (badeq.go:8)       CMPB    AL, $2
            0x003a 00058 (badeq.go:8)       SETEQ   AL
            0x003d 00061 (badeq.go:8)       MOVB    AL, "".~r1+24(FP)
            0x0041 00065 (badeq.go:8)       ADDQ    $8, SP
            0x0045 00069 (badeq.go:8)       RET
            0x0046 00070 (badeq.go:8)       MOVB    $0, AL
            0x0048 00072 (badeq.go:8)       JMP     61
    
    SSA after:
    
    "".g t=1 size=32 args=0x10 locals=0x0
            0x0000 00000 (badeq.go:7)       TEXT    "".g(SB), $0-16
            0x0000 00000 (badeq.go:7)       NOP
            0x0000 00000 (badeq.go:7)       NOP
            0x0000 00000 (badeq.go:7)       FUNCDATA        $0, gclocals·23e8278e2b69a3a75fa59b23c49ed6ad(SB)
            0x0000 00000 (badeq.go:7)       FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
            0x0000 00000 (badeq.go:8)       MOVBLZX "".a+8(FP), AX
            0x0005 00005 (badeq.go:8)       CMPB    AL, $1
            0x0007 00007 (badeq.go:8)       JNE     25
            0x0009 00009 (badeq.go:8)       MOVBLZX "".a+9(FP), CX
            0x000e 00014 (badeq.go:8)       CMPB    CL, $2
            0x0011 00017 (badeq.go:8)       SETEQ   AL
            0x0014 00020 (badeq.go:8)       MOVB    AL, "".~r1+16(FP)
            0x0018 00024 (badeq.go:8)       RET
            0x0019 00025 (badeq.go:8)       MOVB    $0, AL
            0x001b 00027 (badeq.go:8)       JMP     20
    
    
    Change-Id: I120185d58012b7bbcdb1ec01225b5b08d0855d86
    Reviewed-on: https://go-review.googlesource.com/22277
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 157fc454eccb850a0a74029a49f8d947ff1a3762
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jul 6 17:11:29 2016 -0700

    path/filepath: don't return SkipDir at top
    
    If the walker function called on a top-level file returns SkipDir,
    then (before this change) Walk would return SkipDir, which the
    documentation implies will not happen.
    
    Fixes #16280.
    
    Change-Id: I37d63bdcef7af4b56e342b624cf0d4b42e65c297
    Reviewed-on: https://go-review.googlesource.com/24780
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 307de6540a5e6e5c2353c410240bb0f98bab1624
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 7 16:21:53 2016 -0700

    cmd/compile/internal/obj/x86: clean up "is leaf?" check
    
    Minor code cleanup. No functional changes.
    
    Change-Id: I2e631b43b122174302a182a1a286c0f873851ce6
    Reviewed-on: https://go-review.googlesource.com/24813
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 64e152910e5b8ef3cad5aa4d02070fdda645c378
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jun 16 10:49:32 2016 -0700

    cmd/internal/obj/x86: remove pointless NOPs
    
    They are no longer needed by stkcheck.
    
    Fixes #16057
    
    Change-Id: I57cb55de5b7a7a1d31a3da200a3a2d51576b68f5
    Reviewed-on: https://go-review.googlesource.com/26667
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ef694a01104168ef4e48579ebdd0d840445d7fd7
Author: Egon Elbre <egonelbre@gmail.com>
Date:   Thu Aug 11 09:05:54 2016 +0300

    website: recreate 16px and 32px favicon
    
    Recreated original favicon with svg. Note, the rasterizations are hand
    tweaked for crispness and straight export will not give the same results.
    
    Fixes #6938
    
    Change-Id: I9bf7b59028711361c29365b145932d90af419b69
    Reviewed-on: https://go-review.googlesource.com/26850
    Reviewed-by: Chris Broadfoot <cbro@golang.org>
```
