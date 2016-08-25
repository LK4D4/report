# August 25, 2016 Report

Number of commits: 330

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 597.698056ms to 581.66067ms, -2.68%
* github.com/coreos/etcd: from 13.541158832s to 13.772457773s, +1.71%
* github.com/gogits/gogs: from 14.09192473s to 14.373830374s, +2.00%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 37.797779594s to 37.105107337s, -1.83%
* github.com/influxdata/influxdb/cmd/influxd: from 6.546223843s to 7.259098139s, +10.89%
* github.com/junegunn/fzf/src/fzf: from 1.174456079s to 1.128131581s, -3.94%
* github.com/mholt/caddy/caddy: from 6.211016332s to 6.384370151s, +2.79%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.601172455s to 1.599379883s, -0.11%
* github.com/nsqio/nsq/apps/nsqd: from 2.205105765s to 2.256486762s, +2.33%
* github.com/prometheus/prometheus/cmd/prometheus: from 10.07719961s to 10.341088543s, +2.62%
* github.com/spf13/hugo: from 8.110848543s to 7.962506692s, -1.83%
* golang.org/x/tools/cmd/guru: from 2.875641975s to 3.182780277s, +10.68%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2679372 to 2597269, -3.06%
* github.com/coreos/etcd: from 24677744 to 23715968, -3.90%
* github.com/gogits/gogs: from 24178929 to 23240749, -3.88%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 28350800 to 27222888, -3.98%
* github.com/influxdata/influxdb/cmd/influxd: from 16741247 to 16103022, -3.81%
* github.com/junegunn/fzf/src/fzf: from 3246360 to 3159936, -2.66%
* github.com/mholt/caddy/caddy: from 15142750 to 14567639, -3.80%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4602599 to 4459055, -3.12%
* github.com/nsqio/nsq/apps/nsqd: from 10012284 to 9678149, -3.34%
* github.com/prometheus/prometheus/cmd/prometheus: from 23119508 to 21872363, -5.39%
* github.com/spf13/hugo: from 16139525 to 15553961, -3.63%
* golang.org/x/tools/cmd/guru: from 7851170 to 7616532, -2.99%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               190           192           +1.05%
BenchmarkMsgpUnmarshal-4             404           402           -0.50%
BenchmarkEasyJsonMarshal-4           1501          1491          -0.67%
BenchmarkEasyJsonUnmarshal-4         2246          2496          +11.13%
BenchmarkFlatBuffersMarshal-4        359           362           +0.84%
BenchmarkFlatBuffersUnmarshal-4      291           291           +0.00%
BenchmarkGogoprotobufMarshal-4       164           160           -2.44%
BenchmarkGogoprotobufUnmarshal-4     253           256           +1.19%

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

* [cmd/compile: add inline explainer](https://github.com/golang/go/commit/e4cae432d6185b75d45fcf0c3f9c6d49591c128a)
* [io: fix infinite loop bug in MultiReader](https://github.com/golang/go/commit/93372673ce51b9462d7ae0f87ac28ffe0c2ad37d)
* [encoding/hex: change lookup table from string to array](https://github.com/golang/go/commit/57370a87d80be0ab588eb8bb9a5e2a31f4613355)

## GIT Log

```
commit e71e1fe87e144ec10287a10b6a41a543762dabff
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Aug 22 12:25:23 2016 -0400

    cmd/compile: get MIPS64 SSA working
    
    - implement *, /, %, shifts, Zero, Move.
    - fix mistakes in comparison.
    - fix floating point rounding.
    - handle RetJmp in assembler (which was not handled, as a consequence
      Duff's device was disabled in the old backend.)
    
    all.bash now passes with SSA on.
    
    Updates #16359.
    
    Change-Id: Ia14eed0ed1176b5d800592080c8f53dded7fe73f
    Reviewed-on: https://go-review.googlesource.com/27592
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e90ae90b7a7de4079413b956472044a021bee7c6
Author: Dave Cheney <dave@cheney.net>
Date:   Thu Aug 25 16:40:50 2016 +1000

    cmd/{asm,compile/internal}: delete dead code
    
    Delete unused fields, methods, vars, and funcs. Spotted by
    honnef.co/go/unused.
    
    Change-Id: I0e65484bbd916e59369c4018be46f120b469d610
    Reviewed-on: https://go-review.googlesource.com/27731
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit f29ec7d74a17249b3dfa721fc6ee1dcf3c77b5d5
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Aug 24 16:06:10 2016 -0700

    runtime: remove unused type sigtabtt
    
    The type sigtabtt was introduced by an automated tool in
    https://golang.org/cl/167550043. It was the Go version of the C type
    SigTab. However, when the C code using SigTab was converted to Go in
    https://golang.org/cl/168500044 it was rewritten to use a different Go
    type, sigTabT, rather than sigtabtt (the difference being that sigTabT
    uses string where sigtabtt uses *int8 from the C type char*). So this is
    just a dreg from the conversion that was never actually used.
    
    Change-Id: I2ec6eb4b25613bf5e5ad1dbba1f4b5ff20f80f55
    Reviewed-on: https://go-review.googlesource.com/27691
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 62861889863d3d61f546d01aa7bd9824df1b33df
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu May 26 18:08:24 2016 -0700

    cmd/compile: optimize integer "in range" expressions
    
    Use unsigned comparisons to reduce from
    two comparisons to one for integer "in range"
    checks, such as a <= b && b < c.
    We already do this for bounds checks.
    Extend it to user code.
    
    This is much easier to do in the front end than SSA.
    A back end optimization would be more powerful,
    but this is a good start.
    
    This reduces the power of some of SSA prove
    inferences (#16653), but those regressions appear
    to be rare and not worth holding this CL for.
    
    Fixes #15844.
    Fixes #16697.
    
    strconv benchmarks:
    
    name                          old time/op  new time/op   delta
    Atof64Decimal-8               41.4ns ± 3%   38.9ns ± 2%   -5.89%  (p=0.000 n=24+25)
    Atof64Float-8                 48.5ns ± 0%   46.8ns ± 3%   -3.64%  (p=0.000 n=20+23)
    Atof64FloatExp-8              97.7ns ± 4%   93.5ns ± 1%   -4.25%  (p=0.000 n=25+20)
    Atof64Big-8                    187ns ± 8%    162ns ± 2%  -13.54%  (p=0.000 n=24+22)
    Atof64RandomBits-8             250ns ± 6%    233ns ± 5%   -6.76%  (p=0.000 n=25+25)
    Atof64RandomFloats-8           160ns ± 0%    152ns ± 0%   -5.00%  (p=0.000 n=21+22)
    Atof32Decimal-8               41.1ns ± 1%   38.7ns ± 2%   -5.86%  (p=0.000 n=24+24)
    Atof32Float-8                 46.1ns ± 1%   43.5ns ± 3%   -5.63%  (p=0.000 n=21+24)
    Atof32FloatExp-8               101ns ± 4%    100ns ± 2%   -1.59%  (p=0.000 n=24+23)
    Atof32Random-8                 136ns ± 3%    133ns ± 3%   -2.83%  (p=0.000 n=22+22)
    Atoi-8                        33.8ns ± 3%   30.6ns ± 3%   -9.51%  (p=0.000 n=24+25)
    AtoiNeg-8                     31.6ns ± 3%   29.1ns ± 2%   -8.05%  (p=0.000 n=23+24)
    Atoi64-8                      48.6ns ± 1%   43.8ns ± 1%   -9.81%  (p=0.000 n=20+23)
    Atoi64Neg-8                   47.1ns ± 4%   42.0ns ± 2%  -10.83%  (p=0.000 n=25+25)
    FormatFloatDecimal-8           177ns ± 9%    178ns ± 6%     ~     (p=0.460 n=25+25)
    FormatFloat-8                  282ns ± 6%    282ns ± 3%     ~     (p=0.954 n=25+22)
    FormatFloatExp-8               259ns ± 7%    255ns ± 6%     ~     (p=0.089 n=25+24)
    FormatFloatNegExp-8            253ns ± 6%    254ns ± 6%     ~     (p=0.941 n=25+24)
    FormatFloatBig-8               340ns ± 6%    341ns ± 8%     ~     (p=0.600 n=22+25)
    AppendFloatDecimal-8          79.4ns ± 0%   80.6ns ± 6%     ~     (p=0.861 n=20+25)
    AppendFloat-8                  175ns ± 3%    174ns ± 0%     ~     (p=0.722 n=25+20)
    AppendFloatExp-8               142ns ± 4%    142ns ± 2%     ~     (p=0.948 n=25+24)
    AppendFloatNegExp-8            137ns ± 2%    138ns ± 2%   +0.70%  (p=0.001 n=24+25)
    AppendFloatBig-8               218ns ± 3%    218ns ± 4%     ~     (p=0.596 n=25+25)
    AppendFloatBinaryExp-8        80.0ns ± 4%   78.0ns ± 1%   -2.43%  (p=0.000 n=24+21)
    AppendFloat32Integer-8        82.3ns ± 3%   79.3ns ± 4%   -3.69%  (p=0.000 n=24+25)
    AppendFloat32ExactFraction-8   143ns ± 2%    143ns ± 0%     ~     (p=0.177 n=23+19)
    AppendFloat32Point-8           175ns ± 3%    175ns ± 3%     ~     (p=0.062 n=24+25)
    AppendFloat32Exp-8             139ns ± 2%    137ns ± 4%   -1.05%  (p=0.001 n=24+24)
    AppendFloat32NegExp-8          134ns ± 0%    137ns ± 4%   +2.06%  (p=0.000 n=22+25)
    AppendFloat64Fixed1-8         97.8ns ± 0%   98.6ns ± 3%     ~     (p=0.711 n=20+25)
    AppendFloat64Fixed2-8          110ns ± 3%    110ns ± 5%   -0.45%  (p=0.037 n=24+24)
    AppendFloat64Fixed3-8          102ns ± 3%    102ns ± 3%     ~     (p=0.684 n=24+24)
    AppendFloat64Fixed4-8          112ns ± 3%    110ns ± 0%   -1.43%  (p=0.000 n=25+18)
    FormatInt-8                   3.18µs ± 4%   3.10µs ± 6%   -2.54%  (p=0.001 n=24+25)
    AppendInt-8                   1.81µs ± 5%   1.80µs ± 5%     ~     (p=0.648 n=25+25)
    FormatUint-8                   812ns ± 6%    816ns ± 6%     ~     (p=0.777 n=25+25)
    AppendUint-8                   536ns ± 4%    538ns ± 3%     ~     (p=0.798 n=20+22)
    Quote-8                        605ns ± 6%    602ns ± 9%     ~     (p=0.573 n=25+25)
    QuoteRune-8                   99.5ns ± 8%  100.2ns ± 7%     ~     (p=0.432 n=25+25)
    AppendQuote-8                  361ns ± 3%    363ns ± 4%     ~     (p=0.085 n=25+25)
    AppendQuoteRune-8             23.3ns ± 3%   22.4ns ± 2%   -3.79%  (p=0.000 n=25+24)
    UnquoteEasy-8                  146ns ± 4%    145ns ± 5%     ~     (p=0.112 n=24+24)
    UnquoteHard-8                  804ns ± 6%    771ns ± 6%   -4.10%  (p=0.000 n=25+24)
    
    Change-Id: Ibd384e46e90f1cfa40503c8c6352a54c65b72980
    Reviewed-on: https://go-review.googlesource.com/27652
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d61c07ffd81b804cbd746ed1dc1e0d9c8b7fca49
Author: Dave Cheney <dave@cheney.net>
Date:   Thu Aug 25 12:32:42 2016 +1000

    cmd/link/internal, cmd/internal/obj: introduce ctxt.Logf
    
    Replace the various calls to Fprintf(ctxt.Bso, ...) with a helper,
    ctxt.Logf. This also addresses the various inconsistent flushing of
    ctxt.Bso.
    
    Because we have two Link structures, add Link.Logf in both places.
    
    Change-Id: I23093f9b9b3bf33089a0ffd7f815f92dcd1a1fa1
    Reviewed-on: https://go-review.googlesource.com/27730
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 5f94ff4c8702e05bdee32b4a5dc6f8c736adab38
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Aug 24 20:27:38 2016 +1000

    cmd/link/internal/ld: move ld.Cpos to coutbuf.Offset
    
    This change moves the ld.Cpos function to a method on coutbuf. This is
    part of a larger change that makes ld.outbuf look more like a bio.Buf in
    an effort to eventually replace the former with the latter.
    
    Change-Id: I506f7131935a2aa903fa302a0fab0c5be50220fd
    Reviewed-on: https://go-review.googlesource.com/27578
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 61d5daea0aca4fcb14f0090685327132a8e1a7f2
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Aug 24 17:44:27 2016 -0400

    runtime: use clock_gettime for time.now() on s390x
    
    This should improve the precision of time.now() from microseconds
    to nanoseconds.
    
    Also, modify runtime.nanotime to keep it consistent with cleanup
    done to time.now.
    
    Updates #11222 for s390x.
    
    Change-Id: I27864115ea1fee7299360d9003cd3a8355f624d3
    Reviewed-on: https://go-review.googlesource.com/27710
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3167ff7ca91b56f3d06c5cda371b2a7c2741cd35
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Aug 24 20:44:09 2016 +1000

    cmd/internal/*: only call ctx.Bso.Flush when something has been written
    
    In many places where ctx.Bso.Flush is used as the target for some debug
    logging, ctx.Bso.Flush is called unconditionally. In the majority of
    cases where debug logging is not enabled, this means Flush is called
    many times when there is nothing to be flushed (it will be called anyway
    when ctx.Bso is eventually closed), sometimes in a loop.
    
    Avoid this by moving the ctx.Bso.Flush call into the same condition
    block as the debug print. This pattern was previously applied
    sporadically.
    
    Change-Id: I0444cb235cc8b9bac51a59b2e44e59872db2be06
    Reviewed-on: https://go-review.googlesource.com/27579
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 1472221a774c2a763dbe545b1686f4e8dc23613f
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Aug 17 10:16:55 2016 -0700

    cmd/go: refactor cgo logic
    
    Extract "cgo -dynimport" and "ld -r" logic into separate helper
    methods to make (*builder).cgo somewhat more manageable.
    
    Fixes #16650.
    
    Change-Id: I3e4d77ff3791528b1233467060d3ea83cb854acb
    Reviewed-on: https://go-review.googlesource.com/27270
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9741d83ceaf046e55077d68719cc2781cd96d5f1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed May 18 13:27:11 2016 -0700

    cmd/asm, go/build: invoke cmd/asm only once per package
    
    Prior to this CL, cmd/go invoked cmd/asm once
    for every assembly file.
    The exec and cmd/asm startup overhead dwarfed
    the actual time spent assembling.
    This CL adds support to cmd/asm to process
    multiple input files and uses it in cmd/go.
    
    This cuts 10% off the wall time for 'go build -a math'.
    
    Fixes #15680
    
    Change-Id: I12d2ee2c817207954961dc8f37b8f2b09f835550
    Reviewed-on: https://go-review.googlesource.com/27636
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 9d4623fe43f121bd98cdb78f207aa12bfc68dbc2
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 1 11:07:05 2016 -0700

    cmd/compile: handle OCONV[NOP] in samesafeexpr
    
    This increases the effectiveness of the
    "integer-in-range" CL that follows.
    
    Change-Id: I23b7b6809f0b2c25ed1d59dd2d5429c30f1db89c
    Reviewed-on: https://go-review.googlesource.com/27651
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 507051d6940c297bd93f5fac5cf786b0e3642d06
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 24 11:33:55 2016 -0700

    spec: complete list of special comma-ok forms
    
    The enumerations didn't include the syntactic form where the lhs is
    full variable declaration with type specification, as in:
    
    var x, ok T = ...
    
    Fixes #15782.
    
    Change-Id: I0f7bafc37dc9dcf62cdb0894a0d157074ccd4b3e
    Reviewed-on: https://go-review.googlesource.com/27670
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 78bc418f1d7f9346e3214ac40506624ab7f5f649
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 24 13:12:34 2016 -0700

    cmd/compile: remove unused importimport function
    
    Functionality is present in bimport.go in slightly modified form.
    
    Change-Id: I6be79d91588873e6ba70d6ab07ba2caa12346dfc
    Reviewed-on: https://go-review.googlesource.com/27672
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a99f812cba4c5a5207fed9be5488312a44a5df34
Author: Keith Randall <khr@golang.org>
Date:   Sat Jul 2 17:19:25 2016 -0700

    cmd/objdump: implement objdump of .o files
    
    Update goobj reader so it can provide all the information
    necessary to disassemble .o (and .a) files.
    
    Grab architecture of .o files from header.
    
    .o files have relocations in them.  This CL also contains a simple
    mechanism to disassemble relocations and add relocation info as an extra
    column in the output.
    
    Fixes #13862
    
    Change-Id: I608fd253ff1522ea47f18be650b38d528dae9054
    Reviewed-on: https://go-review.googlesource.com/24818
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 873dca4c17437d07ae97ef4f9e9a8e8c93d88bd7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Aug 23 17:47:11 2016 -0700

    net: use runtime.Keepalive for *netFD values
    
    The net package sets a finalizer on *netFD. I looked through all the
    uses of *netFD in the package, looking for each case where a *netFD
    was passed as an argument and the final reference to the argument was
    not a function or method call. I added a call to runtime.KeepAlive after
    each such final reference (there were only three).
    
    The code is safe today without the KeepAlive calls because the compiler
    keeps arguments alive for the duration of the function. However, that is
    not a language requirement, so adding the KeepAlive calls ensures that
    this code remains safe even if the compiler changes in the future.
    
    Change-Id: I4e2bd7c5a946035dc509ccefb4828f72335a9ee3
    Reviewed-on: https://go-review.googlesource.com/27650
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e69912e6f44b09d8bafde32f11642579272ab4af
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Aug 23 15:56:43 2016 -0700

    os: use runtime.Keepalive for *Process values
    
    The os package sets a finalizer on *Process. I looked through all the
    uses of *Process in the package, looking for each case where a *Process
    was passed as an argument and the final reference to the argument was
    not a function or method call. I added a call to runtime.KeepAlive after
    each such final reference (there were only three).
    
    The code is safe today without the KeepAlive calls because the compiler
    keeps arguments alive for the duration of the function. However, that is
    not a language requirement, so adding the KeepAlive calls ensures that
    this code remains safe even if the compiler changes in the future.
    
    I also removed an existing unnecessry call to runtime.KeepAlive. The
    syscall.Syscall function is handled specially by the compiler to keep
    its arguments alive.
    
    Change-Id: Ibd2ff20b31ed3de4f6a59dd1633c1b44001d91d9
    Reviewed-on: https://go-review.googlesource.com/27637
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9be2a279ee12d0e17646c0b0b12d8dff9b157a59
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Wed Aug 24 13:22:54 2016 +0900

    internal/testenv: make MustHaveSymlink message friendly
    
    Change-Id: If6e12ebc41152bc0534d3d383df80e960efe97f0
    Reviewed-on: https://go-review.googlesource.com/27577
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5bddca64a84edae33a8ffaa74abcd3a8c966ea0c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jun 20 08:34:39 2016 -0700

    cmd/compile: minor cleanup in mapinit
    
    Change-Id: I7d58d200f7e8b2c0a6e35371da0dafd9b44e9057
    Reviewed-on: https://go-review.googlesource.com/26757
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 41943d96390353695225da6ef44dacdca84d92df
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 16 14:28:44 2016 -0700

    cmd/compile: convert getdyn int arg to bool
    
    Passes toolstash -cmp.
    
    Change-Id: I5b893b8b82b358534fd85542f05e3aa7e666bcd3
    Reviewed-on: https://go-review.googlesource.com/26752
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 67bcee8d985697678ced806569c17cbbab067478
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 16 14:23:12 2016 -0700

    cmd/compile: convert Dodata to a bool, rename to IsStatic
    
    Passes toolstash -cmp.
    
    Change-Id: I2c204ec14b0a72b592fb336acdd4dff55650f7f6
    Reviewed-on: https://go-review.googlesource.com/26751
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a301b329e5c99693cf5a35ad93e9db5198f0b284
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 16 14:14:16 2016 -0700

    cmd/compile: simplify isglobal
    
    Passes toolstash -cmp.
    
    Change-Id: I16ec0c11096bf4c020cf41392effeb67436f32ba
    Reviewed-on: https://go-review.googlesource.com/26750
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit af9342ca10ea661ca34d6d6950edc553fdd27856
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Fri Aug 19 15:00:29 2016 +1000

    syscall, internal/syscall/windows, internal/syscall/windows/registry: make go generate work on every os
    
    Fixes #16368
    
    Change-Id: I2ef7a2deb5798e11cc1d3f8ca29a6e1655155422
    Reviewed-on: https://go-review.googlesource.com/27411
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e9267ca8259dd56ae01db3c6a6350f007c1b84f2
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Aug 24 00:44:03 2016 +0900

    vendor: update vendored route
    
    Updates golang_org/x/net/route to rev 4d38db7 for:
    - route: don't crash or hang up with corrupted messages
    
    Change-Id: I22492f56a5e211b5a0479f1e07ad8f42f7b9ea03
    Reviewed-on: https://go-review.googlesource.com/27574
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c6e0c49b4152ade7f8cc7368c82703a2d1c58f48
Author: Jess Frazelle <me@jessfraz.com>
Date:   Fri May 20 16:17:27 2016 -0700

    cmd/go: updates go get to return exit status 0 for test file only pkgs
    
    Updates the behavior of `go get` to return exit status 0 when a
    requested package only contains test files.
    
    Fixes #15093
    
    Change-Id: I76b80517d58748090f5e8c6f41178361e2d7ca54
    Reviewed-on: https://go-review.googlesource.com/23314
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3e270ab80bbbc259402f0ae22b5eb36f4daec426
Author: Keith Randall <khr@golang.org>
Date:   Tue Aug 23 10:43:47 2016 -0700

    cmd/compile: clean up ctz ops
    
    Now that we have ops that can return 2 results, have BSF return a result
    and flags.  We can then get rid of the redundant comparison and use CMOV
    instead of CMOVconst ops.
    
    Get rid of a bunch of the ops we don't use.  Ctz{8,16}, plus all the Clzs,
    and CMOVNEs.  I don't think we'll ever use them, and they would be easy
    to add back if needed.
    
    Change-Id: I8858a1d017903474ea7e4002fc76a6a86e7bd487
    Reviewed-on: https://go-review.googlesource.com/27630
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 6394eb378eb6b2c0e691c519b2a6664f930b427e
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 23 16:02:19 2016 -0700

    cmd/compile: export package for _ (blank) struct fields
    
    Blank struct fields are regular unexported fields. Two
    blank fields are different if they are from different
    packages. In order to correctly differentiate them, the
    compiler needs the package information. Add it to the
    export data.
    
    Fixes #15514.
    
    Change-Id: I421aaca22b542fcd0d66b2d2db777249cad78df6
    Reviewed-on: https://go-review.googlesource.com/27639
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a8eb6d51bbd668b9e6fefa9f0cf39fccde3a305e
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 23 15:18:09 2016 -0700

    cmd/compile: simplify field/method export (internal cleanup)
    
    Towards a fix for #15514.
    
    Change-Id: I62073e9fdcfe5ddda9b0a47fe8554b524191a77c
    Reviewed-on: https://go-review.googlesource.com/27638
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ec75230b55281a84cd1a9aa52d33d5e552c82dd4
Author: Ross Light <light@google.com>
Date:   Tue Aug 16 11:10:36 2016 -0700

    reflect: document equality guarantee for Type
    
    The only previous mention of this property was in the String() method.
    Since this is the only way to uniquely identify a type and we can't
    change this property without breaking the Go 1 guarantee, it seems
    better to document this property than hiding it on a method.
    
    Fixes #16348
    
    Change-Id: I3d25f7d6e6007e3c15c2e13010869888d0181fc2
    Reviewed-on: https://go-review.googlesource.com/27170
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8050782cba45c739bfb2cd735766879098c48d99
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Aug 23 13:54:01 2016 -0700

    flag: improve comment for calling String with zero value
    
    Update #16694.
    
    Change-Id: Id6be1535d8a146b3dac3bee429ce407a51272032
    Reviewed-on: https://go-review.googlesource.com/27634
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 78fac02cfd6bca33e3b8a13a875c593bbbc353cc
Author: Than McIntosh <thanm@google.com>
Date:   Mon Jul 11 14:26:18 2016 -0400

    test: add test for gccgo issue #15722
    
    Change-Id: I4faf9a55414e217f0c48528efb13ab8fdcd9bb16
    Reviewed-on: https://go-review.googlesource.com/24845
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit e14e67fff63b364f531667c1a2390feb4b3c1c64
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri May 6 09:24:16 2016 -0700

    cmd/compile: clean up one Node.Etype usage
    
    Whoever Marvin is, we're one step closer to realizing his dream.
    
    Change-Id: I8dece4417d0f9ec234be158d0ee7bc6735342d93
    Reviewed-on: https://go-review.googlesource.com/27465
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 874ea6a4c7fc402ff05a8e63dea5d1b41d55930e
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 23 11:20:34 2016 -0700

    cmd/compile: add comment
    
    Minor update on https://go-review.googlesource.com/27441 .
    
    Change-Id: I605a8bfbe67e259020aa53f1d2357808197d02b6
    Reviewed-on: https://go-review.googlesource.com/27631
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 9aea69d6dc20420c726782dc17d9760ab7b68890
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sun Aug 21 17:08:04 2016 +0200

    cmd/compile: fix binary import of unsafe.Pointer literals
    
    Add a type conversion to uintptr for untyped constants
    before the conversion to unsafe.Pointer.
    
    Fixes #16317
    
    Change-Id: Ib85feccad1019e687e7eb6135890b64b82fb87fb
    Reviewed-on: https://go-review.googlesource.com/27441
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ee333b4b7482873e24d39c5c0b3ea1f7ea4ebdcc
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Aug 22 15:56:27 2016 -0700

    cmd/compile: don't Fatal when printing -m debug information
    
    Some FmtSharp export formatting flag support was removed with
    commit b4e9f70. Don't panic if FmtSharp is set, just ignore it.
    
    Fixes #16820.
    
    Change-Id: Ie0c3d3774bd55002f6d2781b1212d070f083e6b2
    Reviewed-on: https://go-review.googlesource.com/27556
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8e24a98abe464161bc8937a84f78189684aa738d
Author: Kevin Burke <kev@inburke.com>
Date:   Sat Aug 20 22:05:47 2016 -0700

    cmd/compile: precompute constant square roots
    
    If a program wants to evaluate math.Sqrt for any constant value
    (for example, math.Sqrt(3)), we can replace that expression with
    its evaluation (1.7320508075688772) at compile time, instead of
    generating a SQRT assembly command or equivalent.
    
    Adds tests that math.Sqrt generates the correct values. I also
    compiled a short program and verified that the Sqrt expression was
    replaced by a constant value in the "after opt" step.
    
    Adds a short doc to the top of generic.rules explaining what the file
    does and how other files interact with it.
    
    Fixes #15543.
    
    Change-Id: I6b6e63ac61cec50763a09ba581024adeee03d4fa
    Reviewed-on: https://go-review.googlesource.com/27457
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d00890b5f3e7846a7b4ab084fe39d279f21aad04
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun Jul 10 18:44:09 2016 -0700

    runtime: add msan calls before calling traceback functions
    
    Tell msan that the arguments to the traceback functions are initialized,
    in case the traceback functions are compiled with -fsanitize=memory.
    
    Change-Id: I3ab0816604906c6cd7086245e6ae2e7fa62fe354
    Reviewed-on: https://go-review.googlesource.com/24856
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 24158644dddc414243fc3c033f85f0565b7ebd3e
Author: Michael Matloob <matloob@golang.org>
Date:   Mon Aug 22 10:33:13 2016 -0400

    cmd/link/internal/ld: camelCase a buch of snake_case names
    
    I've also unexported a few symbols that weren't used outside the
    package.
    
    Updates #16818
    
    Change-Id: I39d9d87b3eec30b88b4a17c1333cfbbfa6b3518f
    Reviewed-on: https://go-review.googlesource.com/27468
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 5e66ac9ac61efb045f8ae28ff0fc019a4a98b473
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun Aug 14 20:39:30 2016 -0700

    flag: document that Value.String must work on the zero value
    
    Otherwise flag.PrintDefaults will fail when it tries to determine
    whether the default is the zero value.
    
    Fixes #16694.
    
    Change-Id: I253fbf11ffc0a9069fd48c2c3cf3074df53e3a03
    Reviewed-on: https://go-review.googlesource.com/27003
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit eb29ac7a7cf97cd364c9d343377b104bc3e90927
Author: Gabriel Russell <gabriel.russell@gmail.com>
Date:   Tue Aug 23 08:31:26 2016 -0400

    time: fix optional fractional seconds range err
    
    The optional fractional seconds overrides any range error
    from the second parsing. Instead don't look for optional fractional
    seconds if a range error has occured.
    
    Change-Id: I27e0a2432740f6753668bd8833e48b9495bc4036
    Reviewed-on: https://go-review.googlesource.com/27590
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 03723c909ee7b75c0c88ddb3f547642b9f48b009
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Tue Aug 23 22:59:27 2016 +0900

    path/filepath: use testenv.MustHaveSymlink to simplify symlink tests
    
    Cleanup test code for symbolic links.
    
    Change-Id: I25f561cd34dc4d120a4143f933619d233a6cffc5
    Reviewed-on: https://go-review.googlesource.com/27573
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e893c72f2a5ec41bbf14f23b95ae10caf609260a
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Tue Aug 23 22:21:49 2016 +0900

    os: use testenv.MustHaveSymlink to simplify symlink tests
    
    Cleanup test code for symbolic links.
    
    Change-Id: I7a116e4d5c0e955578eca53c1af559e9092f60cd
    Reviewed-on: https://go-review.googlesource.com/27572
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 25d18954f68f0f90e337addfe1970b4f960d0418
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jul 21 15:31:56 2016 -0700

    test: add test case that gccgo miscompiled
    
    Change-Id: I384eac632a4a87c12977e56a7d7bad7614305c51
    Reviewed-on: https://go-review.googlesource.com/25143
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a7ed9ff754dd66a75070a376382ca0fa520deefd
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Mon Aug 22 05:52:15 2016 +0900

    net: document unimplemented methods and functions
    
    Fixes #16802.
    
    Change-Id: I41be7bb4e21e3beaa2136ee69771b0f455b2a7c6
    Reviewed-on: https://go-review.googlesource.com/27417
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4333d3823dd78f755f90e5f63b7de180d4a90025
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue Jul 26 12:22:45 2016 +0900

    net: fix typo in error message of TCPConn.ReadFrom
    
    On some error when using io.Copy with TCPConn, it displays an error
    correlation like the following:
    
    read tcp 192.0.2.1:1111->192.0.2.2:2222: read tcp [2001:db8::2]:2222->[2001:db8::3]:3333 read: connection reset by peer
    
    the correlation "some error on reading after reading operation" looks a
    bit confusing because the operation on the ReadFrom method of TCPConn is
    actually "writing after reading." To clarify and avoid confusion, this
    change sets "readfrom" to the Op field of outer-most OpError instead of
    "read."
    
    Change-Id: I6bf4e2e7247143fa54bbcf9cef7a8ae1ede1b35c
    Reviewed-on: https://go-review.googlesource.com/25220
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fe251d2581420811891741b49f38737c9fa4e1cc
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jul 7 21:01:21 2016 -0700

    runtime: remove unused function in test
    
    Change-Id: I43f14cdd9eb4a1d5471fc88c1b4759ceb2c674cf
    Reviewed-on: https://go-review.googlesource.com/24817
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2d85e87f08c325f8be869718c4ac0d7c069161c0
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Aug 2 11:42:53 2016 -0700

    runtime/cgo: add tsan acquire/release around setenv/unsetenv
    
    Change-Id: Iabb25e97714d070c31c657559a97a3bfc979da18
    Reviewed-on: https://go-review.googlesource.com/25403
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0889d2d04a0ab8c4a0d6c437c857f3f721818577
Author: Michael Matloob <matloob@golang.org>
Date:   Tue Aug 23 08:17:58 2016 -0400

    debug/macho: fix comment on Section64
    
    Change-Id: I7c809ec385b56ebb2ec784a1479d466df6ab4d1a
    Reviewed-on: https://go-review.googlesource.com/27565
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d98de0c3d027815bc05f049a528687b436ce984c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 16 13:56:15 2016 -0700

    cmd/compile: use two tables for table-driven map inserts
    
    This enables better packing when key and value
    types have different alignments.
    
    Cuts 57k off cmd/go.
    
    Change-Id: Ifdd125264caccd7852d622382c94e4689e757978
    Reviewed-on: https://go-review.googlesource.com/26669
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit c8941bb85ce434eccf5066f22dd3c161e6afc4e2
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Aug 23 06:14:31 2016 -0700

    doc: add note about CL 24180 to go1.8.txt
    
    Change-Id: Ie2bef1c181a49d7a02ed8068894d2bd81fc5bafa
    Reviewed-on: https://go-review.googlesource.com/27566
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d92a3606f57d3a400eea6e98ddc8db7a09625f44
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jun 16 12:59:09 2016 -0700

    cmd/go, cmd/link: build c-archive as position independent on ELF
    
    This permits people to use -buildmode=c-archive to produce an archive
    file that can be included in a PIE or shared library.
    
    Change-Id: Ie340ee2f08bcff4f6fd1415f7d96d51ee3a7c9a1
    Reviewed-on: https://go-review.googlesource.com/24180
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit dc9755c2a2b561af6c990399938980ab044406ba
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jul 6 15:30:33 2016 -0700

    runtime: add missing race and msan checks to reflect functions
    
    Add missing race and msan checks to reflect.typedmmemove and
    reflect.typedslicecopy. Missing these checks caused the race detector
    to miss races and caused msan to issue false positive errors.
    
    Fixes #16281.
    
    Change-Id: I500b5f92bd68dc99dd5d6f297827fd5d2609e88b
    Reviewed-on: https://go-review.googlesource.com/24760
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit d9504d4623556467f023da5c33236ec0cf4520cb
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Aug 22 22:29:24 2016 -0700

    cmd/link: don't record interpreter in flagInterpreter
    
    Keep flagInterpreter unchanged after flag parsing. This lets us replace
    flagInterpreterSet with flagInterpreter != "".
    
    Change-Id: Ifd2edbb2ce0011e97276ca18281b8ffbabde1c50
    Reviewed-on: https://go-review.googlesource.com/27563
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e3cecfdcaec55c4dd62c7d4fef3d03fed6d03e38
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Mon Aug 22 15:02:59 2016 +0900

    net: fix a typo
    
    Change-Id: I29fadde646095fa8507f239a339857bf53172c14
    Reviewed-on: https://go-review.googlesource.com/27418
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 80b31c05e6ae37c09162406590b9e3b99f0fff9b
Author: Elias Naur <elias.naur@gmail.com>
Date:   Mon Jun 27 21:38:04 2016 +0200

    time: load time zones from the system tzdata file on Android
    
    Android timezones are in a packed format, different from the separate
    files of a regular Unix system. This CL contain the necessary code to
    parse the packed tzdata file and extract time zones from it. It also
    adds a basic test to ensure the new parser works.
    
    Fixes #13581
    
    Change-Id: Idebe73726c3d4c2de89dd6ae1d7d19f975207500
    Reviewed-on: https://go-review.googlesource.com/24494
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0df5ab7e6569a8bf6bc87975e539706163bf664a
Author: Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Date:   Wed Aug 10 15:02:02 2016 -0300

    runtime: Use clock_gettime to get current time on ppc64x
    
    Fetch the current time in nanoseconds, not microseconds, by using
    clock_gettime rather than gettimeofday.
    
    Updates #11222
    
    Change-Id: I1c2c1b88f80ae82002518359436e19099061c6fb
    Reviewed-on: https://go-review.googlesource.com/26790
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 8a9b96ace4c0064d3c06cd483368bd655ad43d87
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Thu Aug 18 19:07:59 2016 +0900

    internal/testenv: add HasSymlink/MustHaveSymlink
    
    os package and path/filepath package have duplicated code for
    checking symlink supports in test code.
    This CL tries to simplify such test code.
    
    Change-Id: I0371488337f5e951eca699852daab9ccb16ddd62
    Reviewed-on: https://go-review.googlesource.com/27331
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0b5f2f0d1149bcff3c6b08458d7ffdd96970235c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jul 15 15:30:53 2016 -0700

    net/http: if context is canceled, return its error
    
    This permits the error message to distinguish between a context that was
    canceled and a context that timed out.
    
    Updates #16381.
    
    Change-Id: I3994b98e32952abcd7ddb5fee08fa1535999be6d
    Reviewed-on: https://go-review.googlesource.com/24978
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 643b9ec07c3a0e440a2b1669896d7a4b4d4dd64b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed May 25 14:10:01 2016 -0700

    cmd/compile: tidy up switch case expansion
    
    No functional changes.
    
    Change-Id: I0961227e8a7be2d7c611452896843b6955303fa6
    Reviewed-on: https://go-review.googlesource.com/26768
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit a9266eef9384090d748acf723351a282b71982a8
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jun 17 14:50:39 2016 -0700

    cmd/compile: simplify constant switch case sorting
    
    This sort is now only reachable for constant clauses
    for a non-interface switch expression value.
    
    Refactor a bit so that the few tests that remain
    are concise and easy to read.
    
    Add a test that string length takes priority
    over built-in string order.
    
    Change-Id: Iedaa11ff77049d5ad1bf14f54cbb8c3411d589a7
    Reviewed-on: https://go-review.googlesource.com/26767
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e26499153e938b21850a7b6b33f27bb3d98e01cd
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Jun 5 17:28:28 2016 -0700

    cmd/compile: use a map to track const switch cases
    
    This is simpler than the sorting technique.
    It also allows us to simplify or eliminate
    some of the sorting decisions.
    
    Most important, sorting will not work when case clauses
    represent ranges of integers: There is no correct
    sort order that allows overlap detection by comparing
    neighbors. Using a map allows of a cheap, simple
    approach to ranges, namely to insert every int
    in the map. The equivalent approach for sorting
    means juggling temporary Nodes for every int,
    which is a lot more expensive.
    
    Change-Id: I84df3cb805992a1b04d14e0e4b2334f943e0ce05
    Reviewed-on: https://go-review.googlesource.com/26766
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b7ac426ee31cfe4464345754745e38bc0af02a66
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jun 3 12:05:32 2016 -0700

    cmd/compile: split genCaseClauses by switch type
    
    The implementations are going to start diverging more.
    Instead of more if clauses and empty parameters,
    specialize.
    
    Change-Id: I44584450592e8c9f72a10d8ada859c07e9d9aa19
    Reviewed-on: https://go-review.googlesource.com/26764
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b046786528d6c802ae007f001a60341fd0d9c9d4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 1 15:21:56 2016 -0700

    cmd/compile: eliminate switch case kinds
    
    We used to have separate kinds for the default
    case and the nil type case. Now that those are
    gone, we can use a simple bool instead.
    
    Change-Id: I65488e945df68178e893cddd2e091ebb6e32ef4d
    Reviewed-on: https://go-review.googlesource.com/26763
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 4739dcf7fbe5f5bbe885c4349d36c3e23658c0c6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed May 25 14:08:13 2016 -0700

    cmd/compile: fix printing of OCASE nodes
    
    Switch lowering splits each case expression out
    into its own OCASE node.
    
    Change-Id: Ifcb72b99975ed36da8540f6e43343e9aa2058572
    Reviewed-on: https://go-review.googlesource.com/26769
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit bd2838be77caf367cd6aede50c48279ee6b3a6c4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 1 11:34:28 2016 -0700

    cmd/compile: use a map to detect duplicate type switch cases
    
    This is a bit simpler than playing sorting games,
    and it is clearer that it generates errors
    in the correct (source) order.
    
    It also allows us to simplify sorting.
    
    It also prevents quadratic error messages for
    (pathological) inputs with many duplicate type cases.
    
    While we’re here, refactoring deduping into separate functions.
    
    Negligible compilebench impact.
    
    Fixes #15912.
    
    Change-Id: I6cc19edd38875389a70ccbdbdf0d9b7d5ac5946f
    Reviewed-on: https://go-review.googlesource.com/26762
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 1a3006b03564936326e591f191a5af0630afca85
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jun 16 13:39:16 2016 -0700

    test: expand switch dead code test to include a range
    
    Change-Id: If443ffb50b140c466dcf4cc5340f44948bfa46a9
    Reviewed-on: https://go-review.googlesource.com/26765
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit cd5ab9798612a08fd39e29ccd6bc10112d1bef0a
Author: Dhaivat Pandit <dhaivatpandit@gmail.com>
Date:   Mon Aug 22 10:29:02 2016 -0700

    net/http/httptest: updated example to use Result()
    
    example for httptest.Recorder was inspecting Recoder directly.
    Using Result() to convert Recorder into a http.Response yields a much
    better user experience.
    
    Closes #16837
    
    Change-Id: Id0e636c12cd6adb1ba11f89953ff2b0f43758cf3
    Reviewed-on: https://go-review.googlesource.com/27495
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 93b753f525b62a2a860fc2ba2d4ea3f788c275f9
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Aug 22 21:49:32 2016 -0700

    cmd/link: fix handling of cgo_dynamic_interpreter
    
    CL 27473 accidentally changed `!Debug['I']` to `*flagInterpreter != ""`.
    Since the old `Debug['I']` was set when the new *flagInterpreter was
    set, this inverted the sense of the condition. The effect was to always
    ignore the cgo_dynamic_interpreter setting from runtime/cgo. This worked
    OK when the default interpreter was the correct one, but failed when it
    was not, as is currently the case on, at least, PPC64 and ARM.
    
    This CL restores the old behavior by using a separate variable to track
    whether the -I flag was used, just as we used to.
    
    Change-Id: Icf9b65fa41349ed2e4de477fec0a557ef1eb8189
    Reviewed-on: https://go-review.googlesource.com/27562
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 550caa1c87ea11aa54c6482ff66f98b4036c474f
Author: Michael Matloob <matloob@golang.org>
Date:   Mon Aug 22 18:53:01 2016 -0400

    cmd/link/internal/mips64: fix use of -s flags
    
    My flags change reversed the meaning of -s within mips64's
    linker code. This should fix that.
    
    Change-Id: Ia24002469e557fb29badfd830134e61c1dd7e16e
    Reviewed-on: https://go-review.googlesource.com/27555
    Reviewed-by: Minux Ma <minux@golang.org>

commit c043e90e559a18691be722f96272846808292fe0
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Aug 19 18:10:00 2016 -0700

    cmd/compile: clean up encoding of export version info
    
    Replace ad-hoc encoding of export version info with a
    more systematic approach.
    
    Continue to read (but not write) the Go1.7 format for backward-
    compatibility. This will avoid spurious errors with old installed
    packages.
    
    Fixes #16244.
    
    Change-Id: I945e79ffd5e22b883250f6f9fac218370c2505a2
    Reviewed-on: https://go-review.googlesource.com/27452
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e8ba80fbf681430d7c5872b7c228f31c6ec6e37e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Aug 22 22:22:44 2016 +0000

    io: fix comment in test
    
    Updates #16795
    
    Change-Id: I0bcc34bb5a92a2c480aebfb0eb6ba57bcc7f7cfd
    Reviewed-on: https://go-review.googlesource.com/27551
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a1dc9465a1db83912a54927d13367404bfbef363
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Thu Aug 18 13:49:10 2016 -0500

    cmd/compile: PPC64.rules for load/store address folding
    
    This adds some additional rules to improve loads and
    stores for ppc64x.
    
    Change-Id: I96b99c3a0019e6ac5393910c086f58330a04fc5a
    Reviewed-on: https://go-review.googlesource.com/27354
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 726190daf3cd66d020369c6ac7df586d53a58a1a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue May 31 13:11:15 2016 -0700

    cmd/compile: explicitly manage default and nil switch cases
    
    Rather than juggle default and nil cases as part
    of a slice, handle them explicitly.
    
    Change-Id: I97b200c9d3f23fe1a438acdbf3d13b0cf7e0851e
    Reviewed-on: https://go-review.googlesource.com/26761
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit cf20525bf414c8268d8fe111460767d52f03df44
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue May 31 15:02:40 2016 -0700

    cmd/compile: set correct line number for multiple defaults in switch error
    
    Fixes #15911.
    
    Change-Id: I500533484de61aa09abe4cecb010445e3176324e
    Reviewed-on: https://go-review.googlesource.com/26760
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 410974497365b2e6e23811dbc58475b00c69d0db
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 21 15:25:22 2016 -0700

    syscall: delete unreachable code
    
    Change-Id: Iacedb792c73591b7fd75e836aab8e0e117c8e738
    Reviewed-on: https://go-review.googlesource.com/27494
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e2103adb6cf4e7d7d0905e513852407355967638
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Aug 7 14:14:47 2016 -0700

    crypto/*, runtime: nacl asm fixes
    
    Found by vet.
    
    Updates #11041
    
    Change-Id: I5217b3e20c6af435d7500d6bb487b9895efe6605
    Reviewed-on: https://go-review.googlesource.com/27493
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5abfc97e84eef1fe5c5782eb8047cb911560ddb7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Aug 9 20:05:41 2016 -0700

    runtime: use correct MOV for plan9 brk_ ret value
    
    Updates #11041
    
    Change-Id: I78f8d48f00cfbb451e37c868cc472ef06ea0fd95
    Reviewed-on: https://go-review.googlesource.com/27491
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e80376ca6b1f3c4cbc11639d998c0c83787f9247
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Aug 8 16:46:48 2016 -0700

    runtime: ignore closeonexec ret val on openbsd/arm
    
    Fixes #16641
    Updates #11041
    
    Change-Id: I087208a486f535d74135591b2c9a73168cf80e1a
    Reviewed-on: https://go-review.googlesource.com/27490
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 355d7fa8a8c362ee148d3394ce834d0b742c6872
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Aug 22 18:59:20 2016 +0000

    net/http: make Transport.CancelRequest doc recommend Request.WithContext
    
    The old deprecation docs were referencing another deprecated field.
    
    Fixes #16752
    
    Change-Id: I44a690048e00ddc790a80214ecb7f5bb0a5b7b34
    Reviewed-on: https://go-review.googlesource.com/27510
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 4187e1f49f45b1cfa66e7cdf8c118ee926285019
Author: Billy Lynch <wlynch@google.com>
Date:   Fri Aug 19 01:40:08 2016 -0400

    net/http/httptrace: add simple example and fix copyright header
    
    Partially addresses #16360
    
    Change-Id: I67a328302d7d91231f348d934e4232fcb844830a
    Reviewed-on: https://go-review.googlesource.com/27398
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2c58cb36f971aed484e880769eb2b0a21654459a
Author: Eric Lagergren <ericscottlagergren@gmail.com>
Date:   Sat Aug 20 15:46:06 2016 -0700

    encoding/xml: do not ignore error return from copyValue
    
    The error return from copyValue was ignored causing some XML attribute
    parsing to swallow an error.
    
    Additionally, type MyMarshalerAttrTest had no UnmarshalXMLAttr method
    causing marshalTests not to be symmetrical and the test suite to fail
    for test case 101.
    
    Fixes #16158
    
    Change-Id: Icebc505295a2c656ca4b42ba37bb0957dd7260c6
    Reviewed-on: https://go-review.googlesource.com/27455
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 747a158ef314bb458b90da95f3e3d67aa4140622
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Mon Aug 8 16:02:09 2016 +0200

    runtime: speed up StartTrace with lots of blocked goroutines
    
    In StartTrace we emit EvGoCreate for all existing goroutines.
    This includes stack unwind to obtain current stack.
    Real Go programs can contain hundreds of thousands of blocked goroutines.
    For such programs StartTrace can take up to a second (few ms per goroutine).
    
    Obtain current stack ID once and use it for all EvGoCreate events.
    
    This speeds up StartTrace with 10K blocked goroutines from 20ms to 4 ms
    (win for StartTrace called from net/http/pprof hander will be bigger
    as stack is deeper).
    
    Change-Id: I9e5ff9468331a840f8fdcdd56c5018c2cfde61fc
    Reviewed-on: https://go-review.googlesource.com/25573
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>

commit 7c5f33b173d7bde6b3ae33bab940b76b4c991556
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jul 11 16:05:41 2016 -0700

    runtime: cull dead code
    
    They are unused, and vet wants them to have
    a function prototype.
    
    Updates #11041
    
    Change-Id: Idedc96ddd3c3cf1b1d2ab6d98796367eab29f032
    Reviewed-on: https://go-review.googlesource.com/27492
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 310a40b4f2d602786becb3a5ed28394311c41ffa
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Aug 19 16:35:36 2016 -0400

    cmd/compile: start MIPS64 port of SSA backend
    
    Fib with all int and float types run correctly.
    *, /, shifts, Zero, Move not implemented yet. No optimization yet.
    
    Updates #16359.
    
    Change-Id: I4b0412954d5fd4c13a5fcddd8689ed8ac701d345
    Reviewed-on: https://go-review.googlesource.com/27404
    Reviewed-by: David Chase <drchase@google.com>

commit e4cae432d6185b75d45fcf0c3f9c6d49591c128a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue May 3 17:21:32 2016 -0700

    cmd/compile: add inline explainer
    
    When compiling with -m -m, this adds output
    for every non-inlined function explaining why
    it was not inlined.
    
    Change-Id: Icb59ae912a835c996e6b3475b163ee5125113001
    Reviewed-on: https://go-review.googlesource.com/22782
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ef6fde26e5cac06c0dbd4ffe628a3ecd2ef7c450
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 21 11:40:01 2016 -0700

    cmd/vet: handle multiple arches in asm build directives
    
    If a build directive contains multiple arches,
    try to match the build context.
    
    Updates #11041
    
    Change-Id: I03b5d7bfb29d1ff6c7d36a9d7c7fabfcc1d871c1
    Reviewed-on: https://go-review.googlesource.com/27158
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 863ca99415915b19218a36e7c4bf836c135ca00c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 21 11:59:40 2016 -0700

    cmd/vet: fix mips64le arch name in asmdecl check
    
    Updates #11041
    
    Change-Id: Ic6df8ef25b7cf280db523950cd3640b060ad1a9b
    Reviewed-on: https://go-review.googlesource.com/27157
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit f410a668ee3f0a5c71e4d9d0bda83b826246ab5d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jul 11 15:20:43 2016 -0700

    cmd/vet: improve error message for cross-package assembly
    
    bytes.Compare has its go prototype in package bytes,
    but its implementation in package runtime.
    vet used to complain that the prototype was missing.
    Now instead:
    
    runtime/asm_amd64.s:1483: [amd64] cannot check cross-package assembly function: Compare is in package bytes
    
    Updates #11041
    
    Change-Id: Ied44fac10d0916d7a34e552c02d052e16fca0c8c
    Reviewed-on: https://go-review.googlesource.com/27153
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 0952a15cd17755c655910e4b2601d0f255d71c42
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jul 11 12:30:45 2016 -0700

    cmd/vet: clean up printing errors with no position
    
    Before:
    
    : runtime/asm_amd64.s:345: [amd64] morestack: use of 8(SP) points beyond argument frame
    
    After:
    
    runtime/asm_amd64.s:345: [amd64] morestack: use of 8(SP) points beyond argument frame
    
    Updates #11041
    
    Change-Id: Ic87a6d1a2a7b2a8bf737407bc981b159825c84f2
    Reviewed-on: https://go-review.googlesource.com/27152
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 4af1148079f00b461c9ae79df22aa647aa7ff5ef
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jul 9 17:10:57 2016 -0700

    cmd/vet: improve asmdecl parameter handling
    
    The asmdecl check had hand-rolled code that
    calculated the size and offset of parameters
    based only on the AST.
    It included a list of known named types.
    
    This CL changes asmdecl to use go/types instead.
    This allows us to easily handle named types.
    It also adds support for structs, arrays,
    and complex parameters.
    
    It improves the default names given to unnamed
    parameters. Previously, all anonymous arguments were
    called "unnamed", and the first anonymous return
    argument was called "ret".
    Anonymous arguments are now called arg, arg1, arg2,
    etc., depending on the index in the argument list.
    Return arguments are ret, ret1, ret2.
    
    This CL also fixes a bug in the printing of
    composite data type sizes.
    
    Updates #11041
    
    Change-Id: I1085116a26fe6199480b680eff659eb9ab31769b
    Reviewed-on: https://go-review.googlesource.com/27150
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 880c967ccd71013253a751452a83e6c6a0cf86df
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Aug 13 18:12:21 2016 -0700

    runtime: minor string/rune optimizations
    
    Eliminate a spill in concatstrings.
    Provide bounds elim hints in runetochar.
    No significant benchmark movement.
    
    Before:
    "".runetochar t=1 size=412 args=0x28 locals=0x0
    "".concatstrings t=1 size=736 args=0x30 locals=0x98
    
    After:
    "".runetochar t=1 size=337 args=0x28 locals=0x0
    "".concatstrings t=1 size=711 args=0x30 locals=0x90
    
    Change-Id: Icce646976cb20a223163b7e72a54761193ac17e3
    Reviewed-on: https://go-review.googlesource.com/27460
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Martin Möhrmann <martisch@uos.de>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fa8a28d55940adbf797e54a9ca6fa68f1e341a93
Author: Michael Matloob <matloob@golang.org>
Date:   Sun Aug 21 18:34:24 2016 -0400

    cmd/link: turn some globals into flag pointer variables
    
    This moves many of the flag globals into main and assigns them
    to their flag.String/Int64/... directly.
    
    Updates #16818
    
    Change-Id: Ibbff44a273bbc5cb7228e43f147900ee8848517f
    Reviewed-on: https://go-review.googlesource.com/27473
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 236901384dc9fc3d7810ae96b43c8404f0fea6c1
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Aug 19 23:13:29 2016 +0000

    net/http: fix unwanted HTTP/2 conn Transport crash after IdleConnTimeout
    
    Go 1.7 crashed after Transport.IdleConnTimeout if an HTTP/2 connection
    was established but but its caller no longer wanted it. (Assuming the
    connection cache was enabled, which it is by default)
    
    Fixes #16208
    
    Change-Id: I9628757f7669e344f416927c77f00ed3864839e3
    Reviewed-on: https://go-review.googlesource.com/27450
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 82c1e22e13aeb60866031a8dc23bd78c60c1e782
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Mon Aug 22 10:53:05 2016 +1200

    cmd/link: make listsort less generic
    
    It's always called with the same arguments now.
    
    Maybe the real fix is to make Symbol.Sub a slice but that requires a bit more
    brain.
    
    Change-Id: I1326d34a0a327554be6d54f9bd402ea328224766
    Reviewed-on: https://go-review.googlesource.com/27416
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit ec8d49c139d2052f200b6867122c678adbbfc142
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Aug 19 23:27:45 2016 +0000

    net/http: update bundled http2 for Transport double STREAM_ENDED error
    
    Updates bundled http2 to x/net/http2 git rev 7394c11 for:
    
    http2: fix protocol violation regression when writing certain request bodies
    https://golang.org/cl/27406
    
    Fixes #16788
    
    Change-Id: I0efcd36e2b4b34a1df79f763d35bf7a3a1858506
    Reviewed-on: https://go-review.googlesource.com/27451
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit a3765723c60b4cbbec3c8e62e3b2ee9e4080eb0b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun May 8 22:05:21 2016 -0700

    cmd/compile: remove inl escape analysis hack
    
    Relevant issues: #5056, #9537, and #11053.
    Their tests all pass.
    
    Change-Id: Ibbe05982ed5f332149ffd2cb6a232b8d677c4454
    Reviewed-on: https://go-review.googlesource.com/27464
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3c55ee642005db840f8d84bbdfaafa66f713559d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri May 6 09:06:05 2016 -0700

    cmd/compile: refactor out method-called-as-function check
    
    Change-Id: I417aae8622d7d363863704594680bd2502a09049
    Reviewed-on: https://go-review.googlesource.com/27463
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 376d9665a80e96f0b550c24be7ffedf0b467e40d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu May 5 11:45:27 2016 -0700

    cmd/compile: don’t consider recursive calls for inlining
    
    We will never inline recursive calls.
    Rather than simulate the recursion until we hit
    the complexity ceiling, just bail early.
    
    Also, remove a pointless n.Op check.
    visitBottomUp guarantees that n will be an
    ODCLFUNC, and caninl double-checks it.
    
    Change-Id: Ifa48331686b24289d34e68cf5bef385f464b6b92
    Reviewed-on: https://go-review.googlesource.com/27462
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d1faf3879ede9efc9f1907dfad04bac8ec08d598
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri May 27 15:33:11 2016 -0700

    cmd/compile: don’t generate pointless gotos during inlining
    
    Their only purpose in life was to suppress an error.
    Suppress that error explicitly instead by reusing
    an existing, aptly named Node field.
    
    This generates fewer blocks during ssa construction.
    
    name       old alloc/op     new alloc/op     delta
    Template       47.5MB ± 0%      47.2MB ± 0%  -0.72%        (p=0.000 n=15+15)
    Unicode        36.8MB ± 0%      36.8MB ± 0%    ~           (p=0.775 n=15+15)
    GoTypes         143MB ± 0%       142MB ± 0%  -1.03%        (p=0.000 n=15+14)
    Compiler        686MB ± 0%       674MB ± 0%  -1.75%        (p=0.000 n=15+15)
    
    name       old allocs/op    new allocs/op    delta
    Template         446k ± 0%        445k ± 0%  -0.20%        (p=0.000 n=15+15)
    Unicode          355k ± 0%        355k ± 0%    ~           (p=0.235 n=13+15)
    GoTypes         1.36M ± 0%       1.36M ± 0%  -0.41%        (p=0.000 n=13+15)
    Compiler        5.77M ± 0%       5.70M ± 0%  -1.16%        (p=0.000 n=15+15)
    
    
    Change-Id: I5f14afb833c9d355688d9a229eb820e95c7657bf
    Reviewed-on: https://go-review.googlesource.com/27461
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 16e3ce278591069c4f49c29f28d222768f026747
Author: Michael Matloob <matloob@golang.org>
Date:   Sun Aug 21 18:34:24 2016 -0400

    cmd/link/internal/ld: rename pobj.go to main.go
    
    The only thing pobj contains is the Ldmain function.
    
    Updates #16818
    
    Change-Id: Id114bdb264cb5ea2f372eb2166201f1f8eb99445
    Reviewed-on: https://go-review.googlesource.com/27472
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0a15d95091f4d7db396da018e49653fbb3b19a53
Author: Michael Matloob <matloob@golang.org>
Date:   Sun Aug 21 18:25:28 2016 -0400

    cmd/link: use standard library flag package where possible
    
    The obj library's flag functions are (mostly) light wrappers
    around the standard library flag package. Use the flag package
    directly where possible.
    
    Most uses of the 'count'-type flags (except for -v) only check
    against 0, so they can safely be replaced by bools. Only -v
    and the flagfns haven't been replaced.
    
    Debug has been turned into a slice of bools rather than ints.
    There was a copy of the -v verbosity in ctxt.Debugvlog, so don't use
    Debug['v'] and just use ctxt.Debugvlog.
    
    Updates #16818
    
    Change-Id: Icf6473a4823c9d35513bbd0c34ea02d5676d782a
    Reviewed-on: https://go-review.googlesource.com/27471
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 65c5d62420a539f2f0d06b3ea2ba837f0fbdd6cf
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jul 8 07:52:02 2016 -0700

    cmd/vet: re-quote struct tags when printing errors
    
    cmd/link/link_test.go contains several multi-line
    struct tags. Going through an unquote/quote cycle
    converts `a
    b
    c` to "a\nb\nc\n".
    
    This keeps all vet error messages for the standard
    library on a single line.
    
    Updates #11041
    
    Change-Id: Ifba1e87297a5174294d1fbf73463fd3db357464f
    Reviewed-on: https://go-review.googlesource.com/27129
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 6ad76718cfdd59977b0008c1e774150a7e39fbd8
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 7 17:40:37 2016 -0700

    cmd/vet: don't treat trailing % as possible formatting directive
    
    Eliminates the following false positive:
    
    cmd/go/go_test.go:1916: possible formatting directive in Error call
    
    The line in question:
    
    tg.t.Error("some coverage results are 0.0%")
    
    Updates #11041
    
    Change-Id: I3b7611fa3e0245714a19bd5388f21e39944f5296
    Reviewed-on: https://go-review.googlesource.com/27128
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 8e90b9026b64a47d68c80a079564b8c17611db0d
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Mon Aug 22 10:27:20 2016 +1200

    cmd/link: remove references to LSym (now Symbol)
    
    Mostly comments but some derived names too.
    
    Change-Id: I1e01dccca98de6688e1426c7a9309f6fd6a1e368
    Reviewed-on: https://go-review.googlesource.com/27415
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    Reviewed-by: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 85072d5f752903298fa52b8d92a6d052061ee5b8
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Apr 22 10:31:14 2016 +1200

    cmd/link: remove Symbol.Next
    
    Bye bye one more class of linked list manipulation!
    
    Change-Id: I2412b224c847dd640f9253125d30cd5f911ce00c
    Reviewed-on: https://go-review.googlesource.com/27414
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit 64f5023e6f4d1bcb7f293d29cdaf259a80864920
Author: Michael Matloob <matloob@golang.org>
Date:   Sun Aug 21 13:52:23 2016 -0400

    cmd/link: remove global Bso variable
    
    Bso is already a member on ld.Link. Use that instead of
    the global.
    
    Updates #16818
    
    Change-Id: Icfc0f6cb1ff551e8129253fb6b5e0d6a94479f51
    Reviewed-on: https://go-review.googlesource.com/27470
    Run-TryBot: Michael Matloob <matloob@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1756b665980613cf655d6ecde709a032568963b0
Author: LE Manh Cuong <cuong.manhle.vn@gmail.com>
Date:   Sun Jul 31 23:10:35 2016 +0700

    os: make ExpandEnv recognize '-' as a special shell parameter
    
    '-' is one of shell special parameters.
    
    The existing implementation of isShellSpecialVar missed '-'
    from the list, causing "$-" and "${-}" expand differently.
    
    Fixes #16554
    
    Change-Id: Iafc7984692cc83cff58f7c1e01267bf78b3a20a9
    Reviewed-on: https://go-review.googlesource.com/25352
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4338e5a8914be2216de09363288177c806ef845e
Author: Michael Matloob <matloob@golang.org>
Date:   Sun Aug 21 13:52:23 2016 -0400

    cmd/link/internal: remove global Ctxt variable
    
    This change threads the *ld.Link Ctxt variable through
    code in arch-specific packages. This removes all remaining
    uses of Ctxt, so remove the global variable too.
    
    This CL continues the work in golang.org/cl/27408
    
    Updates #16818
    
    Change-Id: I5f4536847a1825fd0b944824e8ae4e122ec0fb78
    Reviewed-on: https://go-review.googlesource.com/27459
    Run-TryBot: Michael Matloob <matloob@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e2b30e900064373bce2f4ba4d3917df3de99ac69
Author: Salman Aljammaz <s@0x65.net>
Date:   Sun Aug 21 16:59:56 2016 +0100

    net/http: prepend ./ to directory list hrefs in FileServer
    
    Certain browsers (Chrome 53, Safari 9.1.2, Firefox 46) won't correctly
    follow a directory listing's links if the file name begins with a run
    of characters then a colon, e.g. "foo:bar". Probably mistaking it for
    a URI. However, they are happy to follow "./foo:bar", so this change
    prepends "./" to all link hrefs in the directory listing of
    FileServer.
    
    Change-Id: I60ee8e1ebac73cbd3a3ac0f23e80fdf52e3dc352
    Reviewed-on: https://go-review.googlesource.com/27440
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6a393dc64fb06ba17f737432a45c63233aa4bd3d
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sun Aug 21 09:57:01 2016 +0200

    cmd/compile: fix compilation of math.Sqrt when used as a statement
    
    Fixes #16804
    
    Change-Id: I669c2c24d3135cd35e15a464894ac66945847d0c
    Reviewed-on: https://go-review.googlesource.com/27437
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2f783c3458744a95a01286eed5f4c0af0978cd46
Author: Michael Matloob <matloob@golang.org>
Date:   Fri Aug 19 22:40:38 2016 -0400

    cmd/link/internal: thread *ld.Link through calls
    
    Ctxt is a global defined in cmd/link/internal/ld of type *ld.Link.
    Start threading a *ld.Link through function calls instead of
    relying on the global variable.
    
    Ctxt is still used as a global by the architecture-specific packages,
    but I plan to fix that in a subsequent CL.
    
    Change-Id: I77a3a58bd396fafd959fa1d8b1c83008a9f5a7fb
    Reviewed-on: https://go-review.googlesource.com/27408
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 4b17b152a3a3d238669c93b31de34e87c2855f6e
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sat Aug 20 21:09:53 2016 -0400

    hash/crc32: fix optimized s390x implementation
    
    The code wasn't checking to see if the data was still >= 64 bytes
    long after aligning it.
    
    Aligning the data is an optimization and we don't actually need
    to do it. In fact for smaller sizes it slows things down due to
    the overhead of calling the generic function. Therefore for now
    I have simply removed the alignment stage. I have also added a
    check into the assembly to deliberately trigger a segmentation
    fault if the data is too short.
    
    Fixes #16779.
    
    Change-Id: Ic01636d775efc5ec97689f050991cee04ce8fe73
    Reviewed-on: https://go-review.googlesource.com/27409
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 94d9cc7741a13206e139efeab84fcc589dc390e5
Author: Konstantin Shaposhnikov <k.shaposhnikov@gmail.com>
Date:   Sat Aug 20 19:29:01 2016 +0800

    index/suffixarray: add Lookup example
    
    Updates #16360
    
    Change-Id: Idd8523b5a9a496ebd9c6e3b89c30df539842a139
    Reviewed-on: https://go-review.googlesource.com/27433
    Reviewed-by: C Cirello <uldericofilho@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7995cb86e54345cb2e3911bce4da00fbb68c2dce
Author: Ian Gudger <igudger@google.com>
Date:   Fri Jul 22 16:38:27 2016 -0700

    syscall: validate ParseUnixCredentials inputs
    
    Don't panic, crash, or return references to uninitialized memory when
    ParseUnixCredentials is passed invalid input.
    
    Fixes #16475
    
    Change-Id: I140d41612e8cd8caaa94be829a415159659c217b
    Reviewed-on: https://go-review.googlesource.com/25154
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fa897643a18d71a62bade50f80171f5e58449f5a
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sat Aug 20 18:36:27 2016 +0100

    runtime: remove unnecessary calls to memclr
    
    Go will have already cleared the structs (the original C wouldn't
    have).
    
    Change-Id: I4a5a0cfd73953181affc158d188aae2ce281bb33
    Reviewed-on: https://go-review.googlesource.com/27435
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8cd04da7627576197b4fb1ef28985ac9243d13a3
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Aug 16 16:03:00 2016 -0700

    compress/flate: make huffmanBitWriter errors persistent
    
    For persistent error handling, the methods of huffmanBitWriter have to be
    consistent about how they check errors. It must either consistently
    check error *before* every operation OR immediately *after* every
    operation. Since most of the current logic uses the previous approach,
    we apply the same style of error checking to writeBits and all calls
    to Write such that they only operate if w.err is already nil going
    into them.
    
    The error handling approach is brittle and easily broken by future commits to
    the code. In the near future, we should switch the logic to use panic at the
    lowest levels and a recover at the edge of the public API to ensure
    that errors are always persistent.
    
    Fixes #16749
    
    Change-Id: Ie1d83e4ed8842f6911a31e23311cd3cbf38abe8c
    Reviewed-on: https://go-review.googlesource.com/27200
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c10f8700e065338dcacda0f4659339e86c402358
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Fri Aug 19 10:46:49 2016 -0700

    net/http/httptrace: test the order of hooks when ctx has multi ClientTraces
    
    Change-Id: I95cae14bb5561947ada9577fb05053f93321a4a8
    Reviewed-on: https://go-review.googlesource.com/27400
    Run-TryBot: Jaana Burcu Dogan <jbd@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a072fc2e6714c20699ee627f88c5a5d5b948340f
Author: Michael Matloob <matloob@golang.org>
Date:   Fri Aug 19 11:35:54 2016 -0400

    cmd/link/internal: rename LSym to Symbol, and add a doc comment.
    
    I'd also like to document some of its fields, but I don't know
    what they are.
    
    Change-Id: I87d341e255f785d351a8a73e645be668e02b2689
    Reviewed-on: https://go-review.googlesource.com/27399
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0a2a64d85d52ad51de34d39bc5685c39c0e1e32a
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Jul 17 14:22:15 2016 -0700

    encoding/gob: error out instead of panicking on nil dereference
    
    Do not panic when we encounter nil interface values which are
    invalid values for gob. Previously this wasn't caught yet
    we were calling reflect.*.Type() on reflect.Invalid values
    thereby causing panic:
      `panic: reflect: call of reflect.Value.Type on zero Value.`
    which is a panic not enforced by encoding/gob itself.
    We can catch this and send back an error to the caller.
    
    Fixes #16204
    
    Change-Id: Ie646796db297759a74a02eee5267713adbe0c3a0
    Reviewed-on: https://go-review.googlesource.com/24989
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Rob Pike <r@golang.org>

commit 14e59511661303ab1406f7c21ee27e58bcd0750e
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Mon Jun 27 12:23:39 2016 +0200

    runtime: increase malloc size classes
    
    When we calculate class sizes, in some cases we discard considerable
    amounts of memory without an apparent reason. For example, we choose
    size 8448 with 6 objects in 7 pages. But we can well use object
    size 9472, which is also 6 objects in 7 pages but +1024 bytes (+12.12%).
    
    Increase class sizes to the max value that leads to the same
    page count/number of objects. Full list of affected size classes:
    
    class 36: pages: 2 size: 1664->1792 +128 (7.69%)
    class 39: pages: 1 size: 2560->2688 +128 (5.0%)
    class 40: pages: 3 size: 2816->3072 +256 (9.9%)
    class 41: pages: 2 size: 3072->3200 +128 (4.16%)
    class 42: pages: 3 size: 3328->3456 +128 (3.84%)
    class 44: pages: 3 size: 4608->4864 +256 (5.55%)
    class 47: pages: 4 size: 6400->6528 +128 (2.0%)
    class 48: pages: 5 size: 6656->6784 +128 (1.92%)
    class 51: pages: 7 size: 8448->9472 +1024 (12.12%)
    class 52: pages: 6 size: 8704->9728 +1024 (11.76%)
    class 53: pages: 5 size: 9472->10240 +768 (8.10%)
    class 54: pages: 4 size: 10496->10880 +384 (3.65%)
    class 57: pages: 7 size: 14080->14336 +256 (1.81%)
    class 59: pages: 9 size: 16640->18432 +1792 (10.76%)
    class 60: pages: 7 size: 17664->19072 +1408 (7.97%)
    class 62: pages: 8 size: 21248->21760 +512 (2.40%)
    class 64: pages: 10 size: 24832->27264 +2432 (9.79%)
    class 65: pages: 7 size: 28416->28672 +256 (0.90%)
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.59s ± 5%     2.52s ± 4%    ~     (p=0.132 n=6+6)
    Fannkuch11-12                2.13s ± 3%     2.17s ± 3%    ~     (p=0.180 n=6+6)
    FmtFprintfEmpty-12          47.0ns ± 3%    46.6ns ± 1%    ~     (p=0.355 n=6+5)
    FmtFprintfString-12          131ns ± 0%     131ns ± 1%    ~     (p=0.476 n=4+6)
    FmtFprintfInt-12             121ns ± 6%     122ns ± 2%    ~     (p=0.511 n=6+6)
    FmtFprintfIntInt-12          182ns ± 2%     186ns ± 1%  +2.20%  (p=0.015 n=6+6)
    FmtFprintfPrefixedInt-12     184ns ± 5%     181ns ± 2%    ~     (p=0.645 n=6+6)
    FmtFprintfFloat-12           272ns ± 7%     265ns ± 1%    ~     (p=1.000 n=6+5)
    FmtManyArgs-12               783ns ± 2%     802ns ± 2%  +2.38%  (p=0.017 n=6+6)
    GobDecode-12                7.04ms ± 4%    7.00ms ± 2%    ~     (p=1.000 n=6+6)
    GobEncode-12                6.36ms ± 6%    6.17ms ± 6%    ~     (p=0.240 n=6+6)
    Gzip-12                      242ms ±14%     233ms ± 7%    ~     (p=0.310 n=6+6)
    Gunzip-12                   36.6ms ±22%    36.0ms ± 9%    ~     (p=0.841 n=5+5)
    HTTPClientServer-12         93.1µs ±29%    88.0µs ±32%    ~     (p=0.240 n=6+6)
    JSONEncode-12               27.1ms ±39%    26.2ms ±35%    ~     (p=0.589 n=6+6)
    JSONDecode-12               71.7ms ±36%    71.5ms ±36%    ~     (p=0.937 n=6+6)
    Mandelbrot200-12            4.78ms ±10%    4.70ms ±16%    ~     (p=0.394 n=6+6)
    GoParse-12                  4.86ms ±34%    4.95ms ±36%    ~     (p=1.000 n=6+6)
    RegexpMatchEasy0_32-12       110ns ±37%     110ns ±36%    ~     (p=0.660 n=6+6)
    RegexpMatchEasy0_1K-12       240ns ±38%     234ns ±47%    ~     (p=0.554 n=6+6)
    RegexpMatchEasy1_32-12      77.2ns ± 2%    77.2ns ±10%    ~     (p=0.699 n=6+6)
    RegexpMatchEasy1_1K-12       337ns ± 5%     331ns ± 4%    ~     (p=0.552 n=6+6)
    RegexpMatchMedium_32-12      125ns ±13%     132ns ±26%    ~     (p=0.561 n=6+6)
    RegexpMatchMedium_1K-12     35.9µs ± 3%    36.1µs ± 5%    ~     (p=0.818 n=6+6)
    RegexpMatchHard_32-12       1.81µs ± 4%    1.82µs ± 5%    ~     (p=0.452 n=5+5)
    RegexpMatchHard_1K-12       52.4µs ± 2%    54.4µs ± 3%  +3.84%  (p=0.002 n=6+6)
    Revcomp-12                   401ms ± 2%     390ms ± 1%  -2.82%  (p=0.002 n=6+6)
    Template-12                 54.5ms ± 3%    54.6ms ± 1%    ~     (p=0.589 n=6+6)
    TimeParse-12                 294ns ± 1%     298ns ± 2%    ~     (p=0.160 n=6+6)
    TimeFormat-12                323ns ± 4%     318ns ± 5%    ~     (p=0.297 n=6+6)
    
    name                      old speed      new speed      delta
    GobDecode-12               109MB/s ± 4%   110MB/s ± 2%    ~     (p=1.000 n=6+6)
    GobEncode-12               121MB/s ± 6%   125MB/s ± 6%    ~     (p=0.240 n=6+6)
    Gzip-12                   80.4MB/s ±12%  83.3MB/s ± 7%    ~     (p=0.310 n=6+6)
    Gunzip-12                  495MB/s ±41%   541MB/s ± 9%    ~     (p=0.931 n=6+5)
    JSONEncode-12             80.7MB/s ±39%  82.8MB/s ±34%    ~     (p=0.589 n=6+6)
    JSONDecode-12             30.4MB/s ±40%  31.0MB/s ±37%    ~     (p=0.937 n=6+6)
    GoParse-12                13.2MB/s ±33%  13.2MB/s ±35%    ~     (p=1.000 n=6+6)
    RegexpMatchEasy0_32-12     321MB/s ±34%   326MB/s ±34%    ~     (p=0.699 n=6+6)
    RegexpMatchEasy0_1K-12    4.49GB/s ±31%  4.74GB/s ±37%    ~     (p=0.589 n=6+6)
    RegexpMatchEasy1_32-12     414MB/s ± 2%   415MB/s ± 9%    ~     (p=0.699 n=6+6)
    RegexpMatchEasy1_1K-12    3.03GB/s ± 5%  3.09GB/s ± 4%    ~     (p=0.699 n=6+6)
    RegexpMatchMedium_32-12   7.99MB/s ±12%  7.68MB/s ±22%    ~     (p=0.589 n=6+6)
    RegexpMatchMedium_1K-12   28.5MB/s ± 3%  28.4MB/s ± 5%    ~     (p=0.818 n=6+6)
    RegexpMatchHard_32-12     17.7MB/s ± 4%  17.0MB/s ±15%    ~     (p=0.351 n=5+6)
    RegexpMatchHard_1K-12     19.6MB/s ± 2%  18.8MB/s ± 3%  -3.67%  (p=0.002 n=6+6)
    Revcomp-12                 634MB/s ± 2%   653MB/s ± 1%  +2.89%  (p=0.002 n=6+6)
    Template-12               35.6MB/s ± 3%  35.5MB/s ± 1%    ~     (p=0.615 n=6+6)
    
    Change-Id: I465a47f74227f316e3abea231444f48c7a30ef85
    Reviewed-on: https://go-review.googlesource.com/24493
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit ab9137dd24e10a9f884475413437cc31e48dbdf7
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Fri Aug 19 11:13:11 2016 -0700

    context: test WithCancel with canceled parent
    
    Change-Id: I32079cc12cfffb8520f0073a8b5119705dc0cd1b
    Reviewed-on: https://go-review.googlesource.com/27401
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3de7dbb19145dcc6b7db8da7aef695961dbb5ece
Author: Austin Clements <austin@google.com>
Date:   Fri Aug 19 16:03:14 2016 -0400

    runtime: fix check for vacuous page boundary rounding again
    
    The previous fix for this, commit 336dad2a, had everything right in
    the commit message, but reversed the test in the code. Fix the test in
    the code.
    
    This reversal effectively disabled the scavenger on large page systems
    *except* in the rare cases where this code was originally wrong, which
    is why it didn't obviously show up in testing.
    
    Fixes #16644. Again. :(
    
    Change-Id: I27cce4aea13de217197db4b628f17860f27ce83e
    Reviewed-on: https://go-review.googlesource.com/27402
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a50f9859bd15784d047875624c9fa91ce69bd85b
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Mon Aug 8 15:41:18 2016 +0200

    internal/trace: fix analysis of EvGoWaiting/EvGoInSyscall events
    
    When tracing is started in the middle of program execution,
    we already have a number of runnable goroutines and a number
    of blocked/in syscall goroutines. In order to reflect these
    goroutines in the trace, we emit EvGoCreate for all existing
    goroutines. Then for blocked/in syscall goroutines we additionally
    emit EvGoWaiting/EvGoInSyscall events. These events don't reset g.ev
    during trace analysis. So next EvGoStart finds g.ev set to the
    previous EvGoCreate. As the result time between EvGoCreate and
    EvGoStart is accounted as scheduler latency. While in reality
    it is blocking/syscall time.
    
    Properly reset g.ev for EvGoWaiting/EvGoInSyscall events.
    
    Change-Id: I0615ba31ed7567600a0667ebb27458481da73adb
    Reviewed-on: https://go-review.googlesource.com/25572
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>

commit 93372673ce51b9462d7ae0f87ac28ffe0c2ad37d
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Aug 18 19:21:26 2016 -0700

    io: fix infinite loop bug in MultiReader
    
    If an io.Reader returned (non-zero, EOF), MultiReader would yield
    bytes forever.
    
    This bug has existed before Go 1 (!!), introduced in the original
    MultiReader implementation in https://golang.org/cl/1764043 and also
    survived basically the only update to this code since then
    (https://golang.org/cl/17873, git rev ccdca832c), which was added in
    Go 1.7.
    
    This just bit me when writing a test for some unrelated code.
    
    Fixes #16795
    
    Change-Id: I36e6a701269793935d19a47ac12f67b07179fbff
    Reviewed-on: https://go-review.googlesource.com/27397
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 244efebe7faa30c087a19a09eb2c72ef4c1824d9
Author: Austin Clements <austin@google.com>
Date:   Sun Jul 17 22:22:32 2016 -0400

    runtime: fix out of date comments
    
    The transition from mark 1 to mark 2 no longer enqueues new root
    marking jobs, but some of the comments still refer to this. Fix these
    comments.
    
    Change-Id: I3f98628dba32c5afe30495ab495da42b32291e9e
    Reviewed-on: https://go-review.googlesource.com/24965
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 8ad70a549364c216c386afcead4dccfdcd39748b
Author: Adam Langley <agl@golang.org>
Date:   Thu Aug 18 16:44:08 2016 -0700

    crypto/x509: allow a leaf certificate to be specified directly as root.
    
    In other systems, putting a leaf certificate in the root store works to
    express that exactly that certificate is acceptable. This change makes
    that work in Go too.
    
    Fixes #16763.
    
    Change-Id: I5c0a8dbc47aa631b23dd49061fb217ed8b0c719c
    Reviewed-on: https://go-review.googlesource.com/27393
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit bcd54f6ca514ae0e347c780b45d743f2c4145bf3
Author: Adam Langley <agl@golang.org>
Date:   Thu Aug 18 17:14:43 2016 -0700

    crypto/x509: recognise ISO OID for RSA+SHA1
    
    For some reason, ISO decided to duplicate the OID for RSA+SHA1. Most
    pertinantly, the makecert.exe utility on Windows is known to have used
    this OID.
    
    This change makes the ISO OID an alias for the normal one.
    
    Change-Id: I60b76265bf1721282bdb0d5c99c98d227c18a878
    Reviewed-on: https://go-review.googlesource.com/27394
    Run-TryBot: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0da545d1280af9895bacff587316b68f50799f16
Author: Adam Langley <agl@golang.org>
Date:   Thu Aug 18 16:20:44 2016 -0700

    encoding/pem: be stricter about the ending line.
    
    Previously the code didn't check the type and final five dashes of the
    ending line of a PEM block.
    
    Fixes #16335.
    
    Change-Id: Ia544e8739ea738d767cfe56c8d46204214ec0b5a
    Reviewed-on: https://go-review.googlesource.com/27391
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a25a7ad70323e8edea4b607aad7c9d2bb96fcc82
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Aug 18 18:24:21 2016 -0700

    cmd/internal/obj: update Bool2int to the form optimized by the compiler
    
    As of https://golang.org/cl/22711 the compiler optimizes this form.
    
    Updates #6011
    
    Change-Id: Ibc6c529dfa24d42f4aab78ebd6722e1d72cb6038
    Reviewed-on: https://go-review.googlesource.com/27395
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 70544c91ffac19f5ffa66c59e3097f3f1fe900f8
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Jun 6 17:59:05 2016 -0700

    cmd/compile/internal/syntax: match old parser errors and line numbers
    
    This makes a bunch of changes to package syntax to tweak line numbers
    for AST nodes. For example, short variable declaration statements are
    now associated with the location of the ":=" token, and function calls
    are associated with the location of the final ")" token. These help
    satisfy many unit tests that assume the old parser's behavior.
    
    Because many of these changes are questionable, they're guarded behind
    a new "gcCompat" const to make them easy to identify and revisit in
    the future.
    
    A handful of remaining tests are too difficult to make behave
    identically. These have been updated to execute with -newparser=0 and
    comments explaining why they need to be fixed.
    
    all.bash now passes with both the old and new parsers.
    
    Change-Id: Iab834b71ca8698d39269f261eb5c92a0d55a3bf4
    Reviewed-on: https://go-review.googlesource.com/27199
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 2ff463948c92e71651a31c621826281035a7071a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri May 6 23:17:29 2016 -0700

    cmd/compile/internal/gc: use new AST parser
    
    Introduce a new noder type to transform package syntax's AST into gc's
    Node tree. Hidden behind a new -newparser flag.
    
    Change-Id: Id0e862ef6196c41533876afc4ec289e21d422d18
    Reviewed-on: https://go-review.googlesource.com/27198
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 117793624b630b0ee63abb16dcb019301adc6472
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Aug 16 13:33:29 2016 -0700

    cmd/compile/internal/syntax: expose additional information for gc
    
    gc needs access to line offsets for Nodes. It also needs access to the
    end line offset for function bodies so it knows what line number to
    use for things like implicit returns and defer executions.
    
    Lastly, include an extra bool to distinguish between simple and full
    slice expressions. This is redundant in valid parse trees, but needed
    by gc for producing complete warnings in invalid inputs.
    
    Change-Id: I64baf334a35c72336d26fa6755c67eb9d6f4e93c
    Reviewed-on: https://go-review.googlesource.com/27196
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4e8c11379345f08ccf47239f7e0ea192917f602a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Aug 17 14:32:49 2016 +0000

    log/syslog: document that Dial uses net.Dial parameters, add example
    
    Fixes #16761
    
    Change-Id: I709daa87926a31e5f8fd46a4c5ef69718ae349b1
    Reviewed-on: https://go-review.googlesource.com/27209
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit 55ea153b785249cfa9b15550d78eeb3df7d825d8
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Aug 18 05:26:51 2016 +0000

    os: reference LookupEnv from the Getenv docs
    
    Fixes #16782
    
    Change-Id: If54917bf5ca1588d8a6d443c3aa6e1d4ada6b620
    Reviewed-on: https://go-review.googlesource.com/27322
    Reviewed-by: anatoly techtonik <techtonik@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 794442375d87e57d012bab2d7424575f6cdff018
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jul 18 17:05:54 2016 -0700

    cmd/vet: allow large shifts of constants
    
    Large shifts of constants are frequently
    used for fancy 32/64 bit detection.
    
    This removes 14 false positives from the
    standard library.
    
    Updates #11041
    
    Change-Id: Ib39346e5c161da04c38a6a3067932ef43bf74f2d
    Reviewed-on: https://go-review.googlesource.com/27155
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 4e79c152b1625857f9b8cdcde59fcf44dc11e95a
Author: Adam Langley <agl@golang.org>
Date:   Wed Aug 17 17:38:06 2016 -0700

    crypto/tls: don't generate random ticket keys if already set.
    
    If SetSessionTicketKeys was called on a fresh tls.Config, the configured
    keys would be overridden with a random key by serverInit.
    
    Fixes #15421.
    
    Change-Id: I5d6cc81fc3e5de4dfa15eb614d102fb886150d1b
    Reviewed-on: https://go-review.googlesource.com/27317
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 89d085de9fbc177ed53f09851b87f920c0322f67
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Aug 18 15:22:32 2016 -0700

    go/types: better doc string for Object.Parent and test
    
    Fixes #14647.
    
    Change-Id: Ib9012a9141e815f5b95f8ca2307e65ffc4587a5b
    Reviewed-on: https://go-review.googlesource.com/27370
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit da2a938910f0401ccd0396891c17b8f8974263de
Author: Adam Langley <agl@golang.org>
Date:   Thu Aug 18 14:49:01 2016 -0700

    crypto/tls: fix comment typo.
    
    This was pointed out in https://go-review.googlesource.com/#/c/27315/1
    but I changed and uploaded the wrong branch. This actually makes the
    fix.
    
    Change-Id: Ib615b06c9141b914648b6abbeeb688c5ffa0d2e3
    Reviewed-on: https://go-review.googlesource.com/27360
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 35bddbba278fe0528b1399ab317b5e8128f02201
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Aug 18 21:50:58 2016 +0000

    doc: upate go1.8.txt
    
    Change-Id: I42597785be6121d8180520b3f7d8e936464f0048
    Reviewed-on: https://go-review.googlesource.com/27361
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8741504888b8ca77736716926e472e90c0e1f765
Author: Adam Langley <agl@golang.org>
Date:   Wed Aug 17 16:45:47 2016 -0700

    crypto/tls: support AES-128-CBC cipher suites with SHA-256.
    
    These were new with TLS 1.2 and, reportedly, some servers require it.
    Since it's easy, this change adds suport for three flavours of
    AES-128-CBC with SHA-256 MACs.
    
    Other testdata/ files have to be updated because this changes the list
    of cipher suites offered by default by the client.
    
    Fixes #15487.
    
    Change-Id: I1b14330c31eeda20185409a37072343552c3464f
    Reviewed-on: https://go-review.googlesource.com/27315
    Run-TryBot: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Jonathan Rudenberg <jonathan@titanous.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ab0bd26da520db67add7e8cfd07b89ffe89df46d
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Aug 18 21:39:16 2016 +0000

    doc: update go1.8.txt
    
    Change-Id: I9a7654a6d623add8542a1c34ccc76ea136a9a7e3
    Reviewed-on: https://go-review.googlesource.com/27359
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 36f61ed7ed962a25f41ecd6ab57cbaa5ce938bf1
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri May 6 23:16:15 2016 -0700

    cmd/dist: build cmd/compile/internal/syntax
    
    Change-Id: Ie6dd2318e031be445c0b1ae65d4c78723d5a1167
    Reviewed-on: https://go-review.googlesource.com/27197
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit c8683ff7977c526fb48ae007971fed16ef32ff62
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Mar 4 17:09:08 2016 -0800

    cmd/compile/internal/syntax: fast Go syntax trees, initial commit.
    
    Syntax tree nodes, scanner, parser, basic printers.
    
    Builds syntax trees for entire Go std lib at a rate of ~1.8M lines/s
    in warmed up state (MacMini, 2.3 GHz Intel Core i7, 8GB RAM):
    
    $ go test -run StdLib -fast
    parsed 1074617 lines (2832 files) in 579.66364ms (1853863 lines/s)
    allocated 282.212Mb (486.854Mb/s)
    PASS
    
    Change-Id: Ie26d9a7bf4e5ff07457aedfcc9b89f0eba72ae3f
    Reviewed-on: https://go-review.googlesource.com/27195
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 3b967be4219c789ef9c47aa5e9607cab3005e1cd
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Aug 18 13:22:22 2016 -0700

    spec: undo spec date change introduced by prior commit
    
    I accidentally included a modified spec in
    https://go-review.googlesource.com/27290 .
    Remove that change.
    
    Change-Id: Icb62fe829072860e9eb74865d21e06f30efcfd26
    Reviewed-on: https://go-review.googlesource.com/27357
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit fd8028dec92166545d080de99021c9c51f05c670
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 17 15:35:15 2016 -0700

    go/types: fix scope extents for range and type switch variables
    
    The changes match the existing compilers, and assume an adjusted
    spec (per issue #16794).
    
    Fixes #15686.
    
    Change-Id: I72677ce75888c41a8f3c2963117a2f2d5501c42b
    Reviewed-on: https://go-review.googlesource.com/27290
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 604efe128195c1fbb1d14cbd36bf681fcff723a3
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Aug 17 16:22:24 2016 -0700

    runtime: disable TestCgoCallbackGC on FreeBSD
    
    The trybot flakes are a nuisance.
    
    Updates #16396
    
    Change-Id: I8202adb554391676ba82bca44d784c6a81bf2085
    Reviewed-on: https://go-review.googlesource.com/27313
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 021f2432131d9bdaa4e89592129a540324f9f22d
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 17 17:05:31 2016 -0700

    go/types: set Info.Types.Type for array composite literals of the form [...]T
    
    Fixes #14092.
    
    Change-Id: I00692f60a416348e38cab256b94fda07e334d258
    Reviewed-on: https://go-review.googlesource.com/27316
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 29df4c8f0004aa259093bf8dbf0bf966a392d44d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Aug 17 16:23:40 2016 -0700

    unsafe: document use of &^ to round/align pointers
    
    Follow-up to CL 27156
    
    Change-Id: I4f1cfced2dced9c9fc8a05bbc00ec4229e85c5c9
    Reviewed-on: https://go-review.googlesource.com/27314
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a7277e5494c696a4798b99e1e55d55acf61211de
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Aug 17 21:23:36 2016 -0400

    cmd/compile: compare size in dead store elimination
    
    Only remove stores that is shadowed by another store with same or
    larger size. Normally we don't need this check because we did check
    the types, but unsafe pointer casting can get around it.
    
    Fixes #16769.
    
    Change-Id: I3f7c6c57807b590a2f735007dec6c65a4fa01a34
    Reviewed-on: https://go-review.googlesource.com/27320
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 5b9ff11c3d50368c44ae7aa9cb4b58c67494e7bb
Author: David Chase <drchase@google.com>
Date:   Mon Aug 15 13:51:00 2016 -0700

    cmd/compile: ppc64le working, not optimized enough
    
    This time with the cherry-pick from the proper patch of
    the old CL.
    
    Stack size increased.
    Corrected NaN-comparison glitches.
    Marked g register as clobbered by calls.
    Fixed shared libraries.
    
    live_ssa.go still disabled because of differences.
    Presumably turning on more optimization will fix
    both the stack size and the live_ssa.go glitches.
    
    Enhanced debugging output for shared libs test.
    
    Rebased onto master.
    
    Updates #16010.
    
    Change-Id: I40864faf1ef32c118fb141b7ef8e854498e6b2c4
    Reviewed-on: https://go-review.googlesource.com/27159
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit dea6dab40c589b486ee6643db3e1204c3379aaee
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Aug 18 15:50:30 2016 +0000

    cmd/yacc: remove go tool yacc
    
    It is no longer used by Go.
    
    It's now moved to golang.org/x/tools/cmd/goyacc for anybody who needs it.
    
    Fixes #11229
    
    Change-Id: Ia431d5a380c7ff784a2050dee2f5bc8acee015da
    Reviewed-on: https://go-review.googlesource.com/27325
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 795ad07b3be3cb51e07d502409f815f7d1f97305
Author: Michael Matloob <matloob@golang.org>
Date:   Thu Jul 28 13:04:41 2016 -0400

    cmd: generate DWARF for functions in compile instead of link.
    
    This is a copy of golang.org/cl/22092 by Ryan Brown.
    
    Here's his original comment:
    On my machine this increases the average time for 'go build cmd/go' from
    2.25s to 2.36s. I tried to measure compile and link separately but saw
    no significant change.
    
    Change-Id: If0d2b756d52a0d30d4eda526929c82794d89dd7b
    Reviewed-on: https://go-review.googlesource.com/25311
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 11e93aa24a61e0c4b25600bf2a681b8779371fe8
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 7 16:09:08 2016 -0700

    cmd/vet: allow any printf verb with any interface
    
    fmt treats interfaces as being transparent.
    As a result, we cannot say with confidence
    that any particular verb is wrong.
    
    This fixes the following vet false positives
    in the standard library:
    
    database/sql/sql_test.go:210: arg dep for printf verb %p of wrong type: sql.finalCloser
    fmt/fmt_test.go:1663: arg nil for printf verb %s of wrong type: untyped nil
    go/ast/commentmap.go:328: arg node for printf verb %p of wrong type: ast.Node
    net/http/transport_test.go:120: arg c for printf verb %p of wrong type: net.Conn
    net/http/httptest/server.go:198: arg c for printf verb %p of wrong type: net.Conn
    net/http/httputil/dump_test.go:258: arg body for printf verb %p of wrong type: io.Reader
    reflect/set_test.go:81: arg x for printf verb %p of wrong type: io.Writer
    reflect/set_test.go:141: arg bb for printf verb %p of wrong type: io.Reader
    
    Updates #11041
    Updates #16314
    
    Change-Id: I76df01abb3c34a97b6960f551bed9c1c91377cfc
    Reviewed-on: https://go-review.googlesource.com/27127
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 0ece9c4b592502647699c2b15d9f0b7332b26de6
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Aug 17 14:57:07 2016 +1000

    debug/pe: revert CL 22720
    
    CL 22720 hid all recently added functionality for go1.7.
    Make everything exported again, so we could use it now.
    
    Updates #15345
    
    Change-Id: Id8ccba7199422b554407ec14c343d2c28fbb8f72
    Reviewed-on: https://go-review.googlesource.com/27212
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit e41b0e2bcb2667425b7eb223baa2b9945466651b
Author: Adam Langley <agl@golang.org>
Date:   Tue Jul 5 13:50:18 2016 -0700

    crypto/x509: support PSS signatures.
    
    Although the term “RSA” is almost synonymous with PKCS#1 v1.5, that
    standard is quite flawed, cryptographically speaking. Bellare and
    Rogaway fixed PKCS#1 v1.5 with OAEP (for encryption) and PSS (for
    signatures) but they only see a fraction of the use of v1.5.
    
    This change adds support for creating and verifying X.509 certificates
    that use PSS signatures. Sadly, every possible dimension of flexibility
    seems to have been reflected in the integration of X.509 and PSS
    resulting in a huge amount of excess complexity. This change only
    supports one “sane” configuration for each of SHA-{256, 384, 512}.
    Hopefully this is sufficient because it saves a lot of complexity in the
    code.
    
    Although X.509 certificates with PSS signatures are rare, I'm inclined
    to look favourably on them because they are sufficiently superior.
    
    Fixes #15958.
    
    Change-Id: I7282e0b68ad0177209f8b2add473b94aa5224c07
    Reviewed-on: https://go-review.googlesource.com/24743
    Run-TryBot: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 59aeac20c0412442848982a9287b4bab66c25682
Author: Adam Langley <agl@golang.org>
Date:   Wed Aug 17 15:55:15 2016 -0700

    crypto/x509: require a NULL parameters for RSA public keys.
    
    The RFC is clear that the Parameters in an AlgorithmIdentifer for an RSA
    public key must be NULL. BoringSSL enforces this so we have strong
    evidence that this is a widely compatible change.
    
    Embarrassingly enough, the major source of violations of this is us. Go
    used to get this correct in only one of two places. This was only fixed
    in 2013 (with 4874bc9b). That's why lots of test certificates are
    updated in this change.
    
    Fixes #16166.
    
    Change-Id: Ib9a4551349354c66e730d44eb8cee4ec402ea8ab
    Reviewed-on: https://go-review.googlesource.com/27312
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 57370a87d80be0ab588eb8bb9a5e2a31f4613355
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Wed Aug 17 20:34:09 2016 +0200

    encoding/hex: change lookup table from string to array
    
    name            old time/op  new time/op  delta
    Encode/256-4     431ns ± 2%   391ns ± 2%   -9.36%  (p=0.000 n=8+8)
    Encode/1024-4   1.68µs ± 0%  1.51µs ± 0%   -9.91%  (p=0.001 n=7+7)
    Encode/4096-4   6.68µs ± 0%  6.03µs ± 1%   -9.69%  (p=0.000 n=8+8)
    Encode/16384-4  27.0µs ± 1%  24.0µs ± 0%  -11.03%  (p=0.000 n=8+7)
    
    Change-Id: I6994e02f77797349c4e188377d84f97dffe98399
    Reviewed-on: https://go-review.googlesource.com/27254
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 17eee31020b982c10a2bf21f446743137968240b
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 17 15:56:29 2016 -0700

    go/types: enable disabled test for comma-ok expressions
    
    This was fixed long ago but the test was not enabled.
    
    For #8189.
    
    Change-Id: Ia44ef752b6bf076f3e243d2d0db326a392a20193
    Reviewed-on: https://go-review.googlesource.com/27310
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 69371671c722d3fcc4f2e1c57dd6a40fc4973ebc
Author: Adam Langley <agl@golang.org>
Date:   Wed Aug 17 13:18:43 2016 -0700

    crypto/hmac: don't test for length equality in Equal.
    
    subtle.ConstantTimeCompare now tests the length of the inputs (although
    it didn't when this code was written) so this test in crypto/hmac is now
    superfluous.
    
    Fixes #16336.
    
    Change-Id: Ic02d8537e776fa1dd5694d3af07a28c4d840d14b
    Reviewed-on: https://go-review.googlesource.com/27239
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b23b9a762c4f876a9f7575fa512074a386e3c6e1
Author: Adam Langley <agl@golang.org>
Date:   Wed Aug 17 13:15:28 2016 -0700

    crypto/x509: return error for missing SerialNumber.
    
    If the SerialNumber is nil in the template then the resulting panic is
    rather deep in encoding/asn1 and it's not obvious what went wrong.
    
    This change tests and returns a more helpful error in this case.
    
    Fixes #16603.
    
    Change-Id: Ib30d652555191eb78f705dff8d909e4b5808f9ca
    Reviewed-on: https://go-review.googlesource.com/27238
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c2322b7ea60f85b4fd3d566b17ab3f7dcb865c1a
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Wed Aug 17 16:08:10 2016 -0700

    runtime: fix the absolute URL to pprof tools
    
    Change-Id: I82eaf5c14a5b8b9ec088409f946adf7b5fd5dbe3
    Reviewed-on: https://go-review.googlesource.com/27311
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c3f05d95932c0d3d2b8461099ce82fa917c7b8a3
Author: Jess Frazelle <me@jessfraz.com>
Date:   Wed Aug 17 13:27:17 2016 -0700

    text/template: remove unused Tree.parse return value
    
    Fixes #13993
    
    Change-Id: Ic61b2bcd9f4f71457d3a8581574633d505d5750e
    Reviewed-on: https://go-review.googlesource.com/27240
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit d1272a8b5c429ed1f43f2935adcb6366abc80c05
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 17 14:05:47 2016 -0700

    go/types: better error message for invalid fallthrough case
    
    Now matches the gc compiler.
    
    Fixes #15594.
    
    Change-Id: I9f3942367bc0acf883c6216b8ca44820832f5fe3
    Reviewed-on: https://go-review.googlesource.com/27241
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 0c819b654f0e2b24f418aeac5c5627516905eda9
Author: Radu Berinde <radu@cockroachlabs.com>
Date:   Tue Aug 16 08:05:39 2016 -0400

    hash/crc32: improve the processing of the last bytes in the SSE4.2 code for AMD64
    
    This commit improves the processing of the final few bytes in
    castagnoliSSE42: instead of processing one byte at a time, we use all
    versions of the CRC32 instruction to process 4 bytes, then 2, then 1.
    The difference is only noticeable for small "odd" sized buffers.
    
    We do the similar improvement for processing the first few bytes in
    the case of unaligned buffer.
    
    Fixing the test which was not actually verifying the results for
    misaligned buffers (WriteString was creating an internal copy which
    was aligned).
    
    Adding benchmarks for length 15 (aligned and misaligned), results
    below.
    
    name                          old time/op    new time/op    delta
    CastagnoliCrc15B-4              25.1ns ± 0%    22.1ns ± 1%  -12.14%
    CastagnoliCrc15BMisaligned-4    25.2ns ± 0%    22.9ns ± 1%   -9.03%
    CastagnoliCrc40B-4              23.1ns ± 0%    23.4ns ± 0%   +1.08%
    CastagnoliCrc1KB-4               127ns ± 0%     128ns ± 0%   +1.18%
    CastagnoliCrc4KB-4               462ns ± 0%     464ns ± 0%     ~
    CastagnoliCrc32KB-4             3.58µs ± 0%    3.60µs ± 0%   +0.58%
    
    name                          old speed      new speed      delta
    CastagnoliCrc15B-4             597MB/s ± 0%   679MB/s ± 1%  +13.77%
    CastagnoliCrc15BMisaligned-4   596MB/s ± 0%   655MB/s ± 1%   +9.94%
    CastagnoliCrc40B-4            1.73GB/s ± 0%  1.71GB/s ± 0%   -1.14%
    CastagnoliCrc1KB-4            8.01GB/s ± 0%  7.93GB/s ± 1%   -1.06%
    CastagnoliCrc4KB-4            8.86GB/s ± 0%  8.83GB/s ± 0%     ~
    CastagnoliCrc32KB-4           9.14GB/s ± 0%  9.09GB/s ± 0%   -0.58%
    
    Change-Id: I499e37af2241d28e3e5d522bbab836c1a718430a
    Reviewed-on: https://go-review.googlesource.com/24470
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3d5cf72ca9beaedc5dcc8b094945de95fa35a670
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Aug 17 13:29:19 2016 -0400

    cmd/compile: CSE copied tuple selectors
    
    In CSE if a tuple generator is CSE'd to a different block, its
    selectors are copied to the same block. In this case, also CES
    the copied selectors.
    
    Test copied from Keith's CL 27202.
    
    Fixes #16741.
    
    Change-Id: I2fc8b9513d430f10d6104275cfff5fb75d3ef3d9
    Reviewed-on: https://go-review.googlesource.com/27236
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 2b8e143dc302d2f3817cb3df1c1cc0b2cde3bbc1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jul 11 12:19:51 2016 -0700

    cmd/vet: infer asm arch from build context
    
    If we cannot infer the asm arch from the filename
    or the build tags, assume that it is the
    current build arch. Assembly files with no
    restrictions ought to be usable on all arches.
    
    Updates #11041
    
    Change-Id: I0ae807dbbd5fb67ca21d0157fe180237a074113a
    Reviewed-on: https://go-review.googlesource.com/27151
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 752e16158751165c53538a564ca0a8bd7ba84f1f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jul 18 17:33:05 2016 -0700

    cmd/vet: allow ^& uintptr arithmetic
    
    The unsafe.Pointer check allows adding to
    and subtracting from uintptrs in order to do
    arithmetic.
    
    Some code needs to round uintptrs.
    Allow &^ for that purpose.
    
    Updates #11041
    
    Change-Id: Ib90dd2954bb6c78427058271e13f2ce4c4af38fb
    Reviewed-on: https://go-review.googlesource.com/27156
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit ebcd17979326b44f445cdfe7640dc53cf14c19ca
Author: Atin M <amalaviy@akamai.com>
Date:   Fri May 6 12:20:12 2016 -0400

    crypto/tls: set Conn.ConnectionState.ServerName unconditionally
    
    Moves the state.ServerName assignment to outside the if
    statement that checks for handshakeComplete.
    
    Fixes #15571
    
    Change-Id: I6c4131ddb16389aed1c410a975f9aa3b52816965
    Reviewed-on: https://go-review.googlesource.com/22862
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit 659dd4f1d7ed4e040d32346fa18c4ae3311ed81a
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Aug 16 14:17:33 2016 -0400

    cmd/compile: add more ARM64 optimizations
    
    - Use machine instructions for uint64<->float conversions
    - Do not enforce alignment on Zero/Move
            ARM64 supports unaligned load/stores, but only aligned offset
            or small offset can be encoded into instructions.
    - Do combined loads
    
    Change-Id: Iffca7dd0f13070b17b784861ce5a30af584680eb
    Reviewed-on: https://go-review.googlesource.com/27086
    Reviewed-by: David Chase <drchase@google.com>

commit cda633b39b1353d23965337fb3118a3fc532a0c1
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Wed May 11 20:55:53 2016 +0200

    math/big: avoid allocation in float.{Add, Sub} when there's no aliasing
    
    name               old time/op    new time/op    delta
    FloatAdd/10-4         116ns ± 1%      82ns ± 0%   -28.74%  (p=0.008 n=5+5)
    FloatAdd/100-4        124ns ± 0%      86ns ± 1%   -30.34%  (p=0.016 n=4+5)
    FloatAdd/1000-4       192ns ± 1%     123ns ± 0%   -35.94%  (p=0.008 n=5+5)
    FloatAdd/10000-4      826ns ± 0%     438ns ± 0%   -46.99%  (p=0.000 n=4+5)
    FloatAdd/100000-4    6.82µs ± 1%    3.36µs ± 0%   -50.74%  (p=0.008 n=5+5)
    FloatSub/10-4         108ns ± 1%      77ns ± 1%   -29.06%  (p=0.008 n=5+5)
    FloatSub/100-4        115ns ± 0%      79ns ± 0%   -31.48%  (p=0.029 n=4+4)
    FloatSub/1000-4       168ns ± 0%      99ns ± 0%   -41.09%  (p=0.029 n=4+4)
    FloatSub/10000-4      690ns ± 2%     288ns ± 1%   -58.24%  (p=0.008 n=5+5)
    FloatSub/100000-4    5.37µs ± 1%    2.10µs ± 1%   -60.89%  (p=0.008 n=5+5)
    
    name               old alloc/op   new alloc/op   delta
    FloatAdd/10-4         48.0B ± 0%     0.0B ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatAdd/100-4        64.0B ± 0%     0.0B ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatAdd/1000-4        176B ± 0%       0B ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatAdd/10000-4     1.41kB ± 0%   0.00kB ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatAdd/100000-4    13.6kB ± 0%    0.0kB ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatSub/10-4         48.0B ± 0%     0.0B ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatSub/100-4        64.0B ± 0%     0.0B ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatSub/1000-4        176B ± 0%       0B ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatSub/10000-4     1.41kB ± 0%   0.00kB ±NaN%  -100.00%  (p=0.008 n=5+5)
    FloatSub/100000-4    13.6kB ± 0%    0.0kB ±NaN%  -100.00%  (p=0.008 n=5+5)
    
    Fixes #14868
    
    Change-Id: Ia2b8b1a8ef0868288ecb25f812b17bd03ff40d1c
    Reviewed-on: https://go-review.googlesource.com/23568
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit f542576b9e576758b7e0a8ec7f8d07b5d0c1f29a
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Jun 24 15:03:04 2016 -0700

    cmd/compile: add compiler phase timing
    
    Timings is a simple data structure that collects times of labeled
    Start/Stop events describing timed phases, which later can be written
    to a file.
    
    Adjacent phases with common label prefix are automatically collected
    in a group together with the accumulated phase time.
    
    Timing data can be appended to a file in benchmark data format
    using the new -bench flag:
    
    $ go build -gcflags="-bench=/dev/stdout" -o /dev/null go/types
    commit: devel +8847c6b Mon Aug 15 17:51:53 2016 -0700
    goos: darwin
    goarch: amd64
    BenchmarkCompile:go/types:fe:init              1       663292 ns/op      0.07 %
    BenchmarkCompile:go/types:fe:loadsys           1      1337371 ns/op      0.14 %
    BenchmarkCompile:go/types:fe:parse             1     47008869 ns/op      4.91 %    10824 lines    230254 lines/s
    BenchmarkCompile:go/types:fe:typecheck:top1    1      2843343 ns/op      0.30 %
    BenchmarkCompile:go/types:fe:typecheck:top2    1       447457 ns/op      0.05 %
    BenchmarkCompile:go/types:fe:typecheck:func    1     15119595 ns/op      1.58 %      427 funcs     28241 funcs/s
    BenchmarkCompile:go/types:fe:capturevars       1        56314 ns/op      0.01 %
    BenchmarkCompile:go/types:fe:inlining          1      9805767 ns/op      1.02 %
    BenchmarkCompile:go/types:fe:escapes           1     53598646 ns/op      5.60 %
    BenchmarkCompile:go/types:fe:xclosures         1       199302 ns/op      0.02 %
    BenchmarkCompile:go/types:fe:subtotal          1    131079956 ns/op     13.70 %
    BenchmarkCompile:go/types:be:compilefuncs      1    692009428 ns/op     72.33 %      427 funcs       617 funcs/s
    BenchmarkCompile:go/types:be:externaldcls      1        54591 ns/op      0.01 %
    BenchmarkCompile:go/types:be:dumpobj           1    133478173 ns/op     13.95 %
    BenchmarkCompile:go/types:be:subtotal          1    825542192 ns/op     86.29 %
    BenchmarkCompile:go/types:unaccounted          1       106101 ns/op      0.01 %
    BenchmarkCompile:go/types:total                1    956728249 ns/op    100.00 %
    
    For #16169.
    
    Change-Id: I93265fe0cb08e47cd413608d0824c5dd35ba7899
    Reviewed-on: https://go-review.googlesource.com/24462
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 289df4e6e65466716aaf1cf65fcfdedada2a5599
Author: Tom Wilkie <tom@weave.works>
Date:   Wed Aug 17 10:13:03 2016 +0100

    net: don't avoid resolving .local addresses
    
    .local addresses are used by things like Kubernetes and Weave DNS; Go
    should not avoid resolving them.
    
    This is a partial revert of https://golang.org/cl/21328 which was too
    strict of an interpretation of RFC 6762.
    
    Fixes #16739
    
    Change-Id: I349415b4eab5d61240dd18217bd95dc7d2105cd5
    Reviewed-on: https://go-review.googlesource.com/27250
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 73fdba2601f1aeaa6565cfb03e4c5c8c98489b7a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Aug 17 08:57:15 2016 -0700

    cmd/compile/internal/s390x: cleanup betypeinit
    
    The Width{int,ptr,reg} assignments are no longer necessary since
    golang.org/cl/21623. The other arch's betypeinit functions were
    cleaned up, but apparently this one was missed.
    
    Change-Id: I1c7f074d7864a561659c1f98aef604f57f285fd0
    Reviewed-on: https://go-review.googlesource.com/27272
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e772c723665c098fb6a90b5e03a61d172e348703
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Aug 17 08:30:31 2016 -0700

    doc/go1.8: support "option ndots:0" in resolv.conf
    
    Updates #15419.
    
    Change-Id: If7c80adcb38b5731e337b2ae2d9d76fcf8513d8e
    Reviewed-on: https://go-review.googlesource.com/27271
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2f73fe7a0db269dcbe51e372809032fa52b8c68c
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Wed Jul 13 10:35:35 2016 -0600

    net: use libresolv rules for ndots range and validation
    
    BIND libresolv allows values from 0 to 15.
    
    For invalid values and negative numbers, 0 is used.
    For numbers greater than 15, 15 is used.
    
    Fixes #15419
    
    Change-Id: I1009bc119c3e87919bcb55a80a35532e9fc3ba52
    Reviewed-on: https://go-review.googlesource.com/24901
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 04e76f295f434bf1bd5ef3b01eed42b638a8b321
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Aug 17 10:32:39 2016 -0400

    test: add test for CL 26831
    
    Test nil check removal for access of PAUTOHEAP.
    
    Change-Id: Id739a9cda7cd3ff173bdcccfedcad93ee90711ef
    Reviewed-on: https://go-review.googlesource.com/27232
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 507144c011c64fd6f18cee6592e79461817e3791
Author: Florian Uekermann <florian@uekermann.me>
Date:   Tue Apr 19 16:06:37 2016 +0200

    math/rand: Document origin of cooked pseudo-random numbers
    
    The Source provided by math/rand relies on an array of cooked
    pseudo-random 63bit integers for seeding. The origin of these
    numbers is undocumented.
    
    Add a standalone program in math/rand folder that generates
    the 63bit integer array as well as a 64bit version supporting
    extension of the Source to 64bit pseudo-random number
    generation while maintaining the current sequence in the
    lower 63bit.
    
    The code is largely based on the initial implementation of the
    random number generator in the go repository by Ken Thompson
    (revision 399).
    
    Change-Id: Ib4192aea8127595027116a0f5a7be53f11dc110b
    Reviewed-on: https://go-review.googlesource.com/22230
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 336dad2a07b17e46e5e81eab8a3209dffbbb928d
Author: Austin Clements <austin@google.com>
Date:   Tue Aug 16 22:05:57 2016 -0400

    runtime: fix check for vacuous page boundary rounding
    
    sysUnused (e.g., madvise MADV_FREE) is only sensible to call on
    physical page boundaries, so scavengelist rounds in the bounds of the
    region being released to the nearest physical page boundaries.
    However, if the region is smaller than a physical page and neither the
    start nor end fall on a boundary, then rounding the start up to a page
    boundary and the end down to a page boundary will result in end < start.
    Currently, we only give up on the region if start == end, so if we
    encounter end < start, we'll call madvise with a negative length and
    the madvise will fail.
    
    Issue #16644 gives a concrete example of this:
    
        start = 0x1285ac000
        end   = 0x1285ae000 (1 8K page)
    
    This leads to the rounded values
    
        start = 0x1285b0000
        end   = 0x1285a0000
    
    which leads to len = -65536.
    
    Fix this by giving up on the region if end <= start, not just if
    end == start.
    
    Fixes #16644.
    
    Change-Id: I8300db492dbadc82ac1ad878318b36bcb7c39524
    Reviewed-on: https://go-review.googlesource.com/27230
    Reviewed-by: Keith Randall <khr@golang.org>

commit 5dc7525b3e04dc7384e11b986ed13ac130afa748
Author: Yasuhiro Matsumoto <mattn.jp@gmail.com>
Date:   Tue Jul 26 04:13:15 2016 +0900

    syscall: mksyscall_windows.go: put path separator in suffix for matching GOROOT
    
    fixes #16493
    
    Change-Id: I86bec2f9bd7965449c43e94733791f7cb18c5c4c
    Reviewed-on: https://go-review.googlesource.com/25165
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 00b779aeed9e0345e4bdbb38c85df63a6352b954
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Tue Aug 16 13:33:27 2016 -0300

    net: simplify internal dtoi and xtoi funcs
    
    Callers pass strings sliced as necessary instead of giving
    an offset.
    
    Fixes #16350
    
    Change-Id: I7ba896f6ff09e0fd0094ca6c5af5d9a81622f15e
    Reviewed-on: https://go-review.googlesource.com/27206
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9c13eb3729988af8be1fd00175bacdd20c7d4206
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Aug 16 15:06:01 2016 -0700

    go/build: introduce go1.8 build tag
    
    Change-Id: Ib8855f8125970fc7ffb271635c28d31d310fcb5b
    Reviewed-on: https://go-review.googlesource.com/27192
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1a7fc7b3a783031ec91844460a44ebffe34a6af1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jun 6 13:01:48 2016 -0700

    cmd/compile: handle e == T comparison more efficiently
    
    Instead of making a runtime call, compare types and values.
    
    Change-Id: Id302083d5a6a5f18e04f36f304f3d290c46976ad
    Reviewed-on: https://go-review.googlesource.com/26660
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 615a52b95b5eedb94297f8de6e7838b16445bd16
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jun 6 12:38:19 2016 -0700

    cmd/compile: inline x, ok := y.(T) where T is a scalar
    
    When T is a scalar, there are no runtime calls
    required, which makes this a clear win.
    
    encoding/binary:
    WriteInts-8                958ns ± 3%     864ns ± 2%   -9.80%  (p=0.000 n=15+15)
    
    This also considerably shrinks a core fmt
    routine:
    
    Before: "".(*pp).printArg t=1 size=3952 args=0x20 locals=0xf0
    After:  "".(*pp).printArg t=1 size=2624 args=0x20 locals=0x98
    
    Unfortunately, I find it very hard to get stable
    numbers out of the fmt benchmarks due to thermal scaling.
    
    Change-Id: I1278006b030253bf8e48dc7631d18985cdaa143d
    Reviewed-on: https://go-review.googlesource.com/26659
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 074d6a649c57a3731e273c8f9dcb36f1663e504a
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 16 17:32:54 2016 -0700

    cmd/compile: remove conditional code dealing with two export formats
    
    This removes some scaffolding introduced pre-1.7, introduced to
    fix an export format bug, and to minimize conflicts with older
    formats. The currently deployed and recognized format is "v1",
    so don't worry about other versions. This is a step towards a
    better scheme for internal export format versioning.
    
    For #16244.
    
    Change-Id: Ic7cf99dd2a24ad5484cc54aed44fa09332c2cf72
    Reviewed-on: https://go-review.googlesource.com/27205
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Robert Griesemer <gri@golang.org>

commit f50ced6d734038b7231bef3c674d541f95b2f23b
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 16 16:43:34 2016 -0700

    cmd/compile: remove encoding of safemode bit from export data
    
    Removes the encoding of this bit which was ignored but left behind
    for 1.7 to minimize pre-1.7 export format changes. See the issue
    for more details.
    
    Fixes #15772.
    
    Change-Id: I46cd7a66ad4c6003b78c64295cf3bda503ebf2dd
    Reviewed-on: https://go-review.googlesource.com/27201
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2b583a190eb14c69bffe5d488d2d6d3862fe76ea
Author: Jess Frazelle <me@jessfraz.com>
Date:   Wed Jun 22 21:57:52 2016 -0700

    text/template: fix Parse when called twice with empty text
    
    Fixes #16156
    
    Change-Id: I6989db4fd392583a2d490339cefc525b07c11b90
    Reviewed-on: https://go-review.googlesource.com/24380
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>

commit 9b88fac00cdbb4025a24fd20e87be121785da579
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jun 17 12:30:27 2016 -0700

    cmd/internal/obj: reduce per-architecture opcode space
    
    s390x took up the last available chunk of int16 opcodes.
    There are RISC-V and sparc64 ports in progress out of tree,
    and there will likely be other architectures.
    Reduce the opcode space to allow more architectures to
    fit without increasing to int32.
    
    This is the smallest power of two that accomodates all
    existing architectures. All else being equal, smaller is
    better--smaller numbers are easier to generate immediates
    for and easier on the eyes when debugging.
    
    Change-Id: I4d0824b28913892fbd0579d3f90bea34e44c8946
    Reviewed-on: https://go-review.googlesource.com/24223
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit b4e9f70412671c7ac06b4852c38a0ab82d94ddf5
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 16 12:55:17 2016 -0700

    cmd/compile: remove support for textual export format
    
    Fixes #15323.
    
    Change-Id: I50e996e6fde6b24327cb45dd84da31deef4dcc56
    Reviewed-on: https://go-review.googlesource.com/27171
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e492d9f01890cf61cb009b3b3617238a8947ebbe
Author: Keith Randall <khr@golang.org>
Date:   Wed Jul 6 15:02:49 2016 -0700

    runtime: fix map iterator concurrent map check
    
    We should check whether there is a concurrent writer at the
    start of every mapiternext, not just in mapaccessK (which is
    only called during certain map growth situations).
    
    Tests turned off by default because they are inherently flaky.
    
    Fixes #16278
    
    Change-Id: I8b72cab1b8c59d1923bec6fa3eabc932e4e91542
    Reviewed-on: https://go-review.googlesource.com/24749
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit a16a189fb96f824d1eaa53db9c0047c7ce334bd1
Author: Keith Randall <khr@golang.org>
Date:   Tue Aug 16 13:52:51 2016 -0700

    test: remove unused variable
    
    ssaMain is no longer needed.
    
    Change-Id: I0b77f0bcd482329d73018bd80a6e068e622e191b
    Reviewed-on: https://go-review.googlesource.com/27190
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5ae82307696458269b373d82c072347d87a2a554
Author: Keith Randall <khr@golang.org>
Date:   Thu Jun 2 12:41:42 2016 -0700

    cmd/compile: use shorter versions of zero-extend ops
    
    Only need to zero-extend to 32 bits and we get the top
    32 bits zeroed for free.
    
    Only the WQ change actually generates different code.
    The assembler did this optimization for us in the other two cases.
    But we might as well do it during SSA so -S output more closely
    matches the actual generated instructions.
    
    Change-Id: I3e4ac50dc4da124014d4e31c86e9fc539d94f7fd
    Reviewed-on: https://go-review.googlesource.com/23711
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 0cd8faf74494f400d278b9b0071908951bea6f40
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jun 17 12:28:31 2016 -0700

    cmd/internal/obj: add opcode space safety check
    
    This CL adds a safety mechanism
    for changing the number of opcodes
    available per architecture.
    
    A subsequent CL will actually make the change.
    
    Change-Id: I6332ed5514f2f153c54d11b7da0cc8a6be1c8066
    Reviewed-on: https://go-review.googlesource.com/24222
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 496174e32e8983f46c6178cab250ea08f4d89dd2
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Aug 16 07:36:13 2016 -0700

    doc/go1.7.html: fix name of tls.Config.NextProtos
    
    Updates #16737
    
    Change-Id: Ia51fc9b06df43b7c6f7136e90b40362263c20081
    Reviewed-on: https://go-review.googlesource.com/27126
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit f88a88402cbf93b37f701a8ca0c70b7d2558a057
Author: Tamir Duberstein <tamird@gmail.com>
Date:   Sun Jul 31 12:55:06 2016 -0400

    regexp: add some tests that were fixed in #12980
    
    Also includes a minor golint cleanup in the tests.
    
    Change-Id: I8c0fc81479e635e7cca18d5c48c28b654afa59d8
    Reviewed-on: https://go-review.googlesource.com/25380
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8c9a7978942296085628bc33eefcb9bab561f4ee
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Mon Aug 15 23:24:31 2016 -0700

    cmd/go: document -v flag for get
    
    Fixes #16719.
    
    Change-Id: I20550628814e3454f17d6f8ae8b66cce17f09859
    Reviewed-on: https://go-review.googlesource.com/27118
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 64214792e214bbacb8c00ffea92a7131e30fa59e
Author: Keith Randall <khr@golang.org>
Date:   Fri Jul 8 15:59:00 2016 -0700

    cmd/compile: allow unsafe.Pointer(nil) as static data
    
    Fixes #16306
    
    Change-Id: If8e2f411fe9a5a5c198f10765fee7261ba8feaf2
    Reviewed-on: https://go-review.googlesource.com/24836
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 9d2b988e4aeb59424411a314748aa8ffd1e71033
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 1 10:58:30 2016 -0700

    cmd/compile: accept literals in samesafeexpr
    
    This only triggers a few places in the stdlib,
    but it helps a lot when it does.
    
    Before:
    
    runtime.mapassign1 t=1 size=2400 args=0x20 locals=0xe0
    
    After:
    
    runtime.mapassign1 t=1 size=2352 args=0x20 locals=0xd8
    
    name           old time/op  new time/op  delta
    MapPop100-8    19.8µs ±11%  18.4µs ± 9%  -7.16%  (p=0.000 n=20+19)
    MapPop1000-8    367µs ±17%   335µs ±11%  -8.63%  (p=0.000 n=19+19)
    MapPop10000-8  7.29ms ±15%  6.86ms ±12%  -5.84%  (p=0.020 n=20+20)
    
    Change-Id: I9faf32f95a6ba6a6d5d0818eab32cc271e01d10a
    Reviewed-on: https://go-review.googlesource.com/26666
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e6f1a886bc49e920533b3e95e96f4965000b9821
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Aug 16 12:14:54 2016 -0400

    cmd/compile: fix uint<->float conversion on 386
    
    The frontend rewriting lowers them to runtime calls on 386. It
    matches explicitly uint32, but missed uint.
    
    Fixes #16738.
    
    Change-Id: Iece7a45edf74615baca052a53273c208f057636d
    Reviewed-on: https://go-review.googlesource.com/27085
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1faea596e4f94e0a4170eda3d1a217b4936d8aa6
Author: Keith Randall <khr@golang.org>
Date:   Sun May 29 11:16:13 2016 -0700

    cmd/compile: add size hint to map literal allocations
    
    Might as well tell the runtime how large the map is going to be.
    This avoids grow work and allocations while the map is being built.
    
    Will wait for 1.8.
    
    Fixes #15880
    Fixes #16279
    
    Change-Id: I377e3e5ec1e2e76ea2a50cc00810adda20ad0e79
    Reviewed-on: https://go-review.googlesource.com/23558
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 35e25ef62efc5917481e11ff6e7a5cc12468b0e2
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jul 20 10:18:23 2016 -0700

    cmd/internal/obj/x86: minor code cleanup
    
    Update #16415
    
    Change-Id: I83e0966931ada2f1ed02304685bb45effdd71268
    Reviewed-on: https://go-review.googlesource.com/26665
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e85265e8c2a41e4a0e703e5fb6fe762cc382d0af
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon May 2 16:16:46 2016 -0700

    cmd/compile: optimize bool to int conversion
    
    This CL teaches SSA to recognize code of the form
    
    // b is a boolean value, i is an int of some flavor
    if b {
            i = 1
    } else {
            i = 0
    }
    
    and use b's underlying 0/1 representation for i
    instead of generating jumps.
    
    Unfortunately, it does not work on the obvious code:
    
    func bool2int(b bool) int {
            if b {
                    return 1
            }
            return 0
    }
    
    This is left for future work.
    Note that the existing phiopt optimizations also don't work for:
    
    func neg(b bool) bool {
            if b {
                    return false
            }
            return true
    }
    
    In the meantime, runtime authors and the like can use:
    
    func bool2int(b bool) int {
            var i int
            if b {
                    i = 1
            } else {
                    i = 0
            }
            return i
    }
    
    This compiles to:
    
    "".bool2int t=1 size=16 args=0x10 locals=0x0
            0x0000 00000 (x.go:25)  TEXT    "".bool2int(SB), $0-16
            0x0000 00000 (x.go:25)  FUNCDATA        $0, gclocals·23e8278e2b69a3a75fa59b23c49ed6ad(SB)
            0x0000 00000 (x.go:25)  FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
            0x0000 00000 (x.go:32)  MOVBLZX "".b+8(FP), AX
            0x0005 00005 (x.go:32)  MOVBQZX AL, AX
            0x0008 00008 (x.go:32)  MOVQ    AX, "".~r1+16(FP)
            0x000d 00013 (x.go:32)  RET
    
    The extraneous MOVBQZX is #15300.
    
    This optimization also helps range and slice.
    The compiler must protect against pointers pointing
    to the end of a slice/string. It does this by increasing
    a pointer by either 0 or 1 * elemsize, based on a condition.
    This CL optimizes away a jump in that code.
    
    This CL triggers 382 times while compiling the standard library.
    
    Updating code to utilize this optimization is left for future CLs.
    
    Updates #6011
    
    Change-Id: Ia7c1185f8aa223c543f91a3cd6d4a2a09c691c70
    Reviewed-on: https://go-review.googlesource.com/22711
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit c7b9bd745642677c8d4b3b76803a39b4e50b4d81
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Jul 15 16:46:50 2016 -0700

    cmd/compile: don't crash when exporting self-recursive interfaces
    
    For #16369.
    
    Change-Id: I4c9f5a66b95558adcc1bcface164b9b2b4382d2f
    Reviewed-on: https://go-review.googlesource.com/24979
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 56752eb2b8cc52311c346f986a0f6e2a9577bfe4
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Aug 16 07:44:57 2016 -0400

    reflect: clear tflag on new types
    
    Fixes #16722
    
    Change-Id: I50a0e69d3e79d13bc1860cd983267c3db087a4b8
    Reviewed-on: https://go-review.googlesource.com/27119
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d251030fa6feac46e5686a78b9681625447a5871
Author: Keith Randall <khr@golang.org>
Date:   Tue Aug 16 09:19:05 2016 -0700

    cmd/compile: don't fold >32bit constants into a MULQ
    
    Don't fold constant factors into a multiply
    beyond the capacity of a MULQ instruction (32 bits).
    
    Fixes #16733
    
    Change-Id: Idc213c6cb06f7c94008a8cf9e60a9e77d085fd89
    Reviewed-on: https://go-review.googlesource.com/27160
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 562d06fc23261b21d12fbcbd407aadee9cb428cb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jun 6 08:29:52 2016 -0700

    cmd/compile: inline _, ok = i.(T)
    
    We already inlined
    
    _, ok = e.(T)
    _, ok = i.(E)
    _, ok = e.(E)
    
    The only ok-only variants not inlined are now
    
    _, ok = i.(I)
    _, ok = e.(I)
    
    These call getitab, so are non-trivial.
    
    Change-Id: Ie45fd8933ee179a679b92ce925079b94cff0ee12
    Reviewed-on: https://go-review.googlesource.com/26658
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit e6e26eeb299f9829ac70bd400d011bfdd266f1c1
Author: Michael Pratt <mpratt@google.com>
Date:   Mon Jul 18 21:59:14 2016 -0700

    cmd/internal/obj: convert Aconv to a stringer
    
    Now that assembler opcodes have their own type, they can have a true
    stringer, rather than explicit calls to Aconv, which makes for nicer
    format strings.
    
    Change-Id: Ic77f5f8ac38b4e519dcaa08c93e7b732226f7bfe
    Reviewed-on: https://go-review.googlesource.com/25045
    Run-TryBot: Michael Pratt <mpratt@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 5693bee0f1f245ffc58c58a41bfd9a431a416957
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 7 15:29:48 2016 -0700

    cmd/compile/internal/big: re-vendor
    
    Pick up a bunch of changes and fixes.
    
    Change-Id: If4101f7185d433a4c89096bc786ee5de8eeabac0
    Reviewed-on: https://go-review.googlesource.com/27123
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6f74c0774cf4fd906292bf0a733cb596f0849780
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jul 11 11:18:17 2016 -0700

    runtime: move printing of extra newline
    
    No functional changes, makes vet happy.
    
    Updates #11041
    
    Change-Id: I59f3aba46d19b86d605508978652d76a1fe7ac7b
    Reviewed-on: https://go-review.googlesource.com/27125
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 297d1d736e3e6db6fa390dd54b1e3de9ea8f1fba
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jul 9 17:09:35 2016 -0700

    net/http: use keyed composite literal
    
    Makes vet happy.
    
    Updates #11041
    
    Change-Id: I23ca413c03ff387359440af8114786cd7880a048
    Reviewed-on: https://go-review.googlesource.com/27124
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 302dd7b71eb565c5460966292e434fc903362cb6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jul 9 14:51:00 2016 -0700

    crypto/cipher, math/big: fix example names
    
    Fixes (legit) vet warnings.
    Fix some verb tenses while we're here.
    
    Updates #11041
    
    Change-Id: I27e995f55b38f4cf584e97a67b8545e8247e83d6
    Reviewed-on: https://go-review.googlesource.com/27122
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 6d2db0986faef807e8538db6f3d32adb0dfd78c5
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jul 9 13:59:36 2016 -0700

    crypto/tls: fix WriteTo method signature
    
    Give *recordingConn the correct WriteTo signature
    to be an io.WriterTo. This makes vet happy.
    It also means that it'll report errors,
    which were previously being ignored.
    
    Updates #11041
    
    Change-Id: I13f171407d63f4b62427679bff362eb74faddca5
    Reviewed-on: https://go-review.googlesource.com/27121
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 12292754d3c46a4172373971b1ba945ecc4c5cbc
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 7 17:41:25 2016 -0700

    net: change t.Error to t.Errorf
    
    Caught by vet.
    
    Updates #11041
    
    Change-Id: I4dbb2eeaf633eea5976074840064edc2349e01d8
    Reviewed-on: https://go-review.googlesource.com/27120
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 88858fa58f892ef19c6dbae2af15af41ecae4937
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jun 30 10:48:54 2016 -0700

    container/list: silence vet warnings
    
    container/list/list_test.go:274: self-assignment of e1 to e1
    container/list/list_test.go:274: self-assignment of e4 to e4
    container/list/list_test.go:282: self-assignment of e1 to e1
    container/list/list_test.go:286: self-assignment of e1 to e1
    container/list/list_test.go:286: self-assignment of e4 to e4
    
    Updates #11041
    
    Change-Id: Ibd90cf6a924e93497908f437b814c3fc82937f4a
    Reviewed-on: https://go-review.googlesource.com/27114
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4c4ca8312064cc47bc3b3cd5efc43ff5b89f4dff
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Aug 11 09:54:04 2016 -0400

    cmd/compile: remove nil check in accessing PAUTOHEAP variable
    
    CL 23393 introduces PAUTOHEAP, and access of PAUTOHEAP variable is
    rewritten to indirection of a PAUTO variable. Mark this variable
    non-nil, so this indirection does not introduce extra nil checks.
    
    Change-Id: I31853eed5e60238b6c5bc0546e2e9ab340dcddd9
    Reviewed-on: https://go-review.googlesource.com/26831
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 40cf4ad0ef2232d65a85d35897ea11aab95e9ef4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 7 15:32:49 2016 -0700

    all: fix "result not used" vet warnings
    
    For tests, assign to _.
    For benchmarks, assign to a sink.
    
    Updates #11041
    
    Change-Id: I87c5543245c7bc74dceb38902f4551768dd37948
    Reviewed-on: https://go-review.googlesource.com/27116
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c70bdd378811be6d56a3483fe6b9c5ed7b40cb2d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jun 30 10:53:47 2016 -0700

    cmd/compile: fix bad generated format strings in test
    
    We were generating format strings containing
    a lone %. Vet legitimately complains:
    
    cmd/compile/internal/gc/constFold_test.go:339: unrecognized printf verb ' '
    
    The fix doesn't make for very readable code,
    but it is simple and obviously correct.
    
    Updates #11041
    
    Change-Id: I90bd2d1d140887f5229752a279f7e46921472fbb
    Reviewed-on: https://go-review.googlesource.com/27115
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2cbe735366a16dcd2595c5f1b1ca38b5b2b5e5b2
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 21 09:02:13 2016 -0700

    syscall: unify unix/amd64 asm implementations
    
    Updates #11041
    
    Change-Id: I77e5ca0b61ffc530ee46848721a177867c81d548
    Reviewed-on: https://go-review.googlesource.com/25116
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a5464af55469e2810452fd62e189404f3b8406ea
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 21 08:57:50 2016 -0700

    syscall: split out unix Syscall9 asm support
    
    This is preliminary work to unifying them.
    Aside from Syscall9, all are identical.
    Syscall9 has a netbsd/openbsd variant
    and a dragonfly/freebsd variant.
    
    Updates #11041
    
    Change-Id: Ia5ce95d5e9115d4c0492d5e53aa7a4316deafd1f
    Reviewed-on: https://go-review.googlesource.com/25115
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 33f95ec4ecf96c795d69c643f80e4bd7f6a8f1bb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 21 08:49:07 2016 -0700

    syscall: superficial cleanup of amd64 unix assembly
    
    This is preliminary work to unifying them.
    
    Updates #11041
    
    Change-Id: Ibe83da3d626f1da9e8888e26cedd3af2152b42e6
    Reviewed-on: https://go-review.googlesource.com/25114
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 856342d844a66cb15e13d73b60b2d2de3ea56090
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 21 08:37:45 2016 -0700

    syscall: fix dragonfly/amd64 assembly argument sizes
    
    This is preliminary work to unifying the
    unix amd64 assembly implementations,
    which is preliminary work to making the
    assembly vet-friendly.
    
    Updates #11041
    
    Change-Id: Ic64985124f8fb86cc08898be2ec7fca972ced4ca
    Reviewed-on: https://go-review.googlesource.com/25113
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b173298c896ebec8e4f898bfae986fce7557c1ec
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 21 08:30:41 2016 -0700

    syscall: unify unix 386 implementations
    
    They were identical.
    
    This will allow us to do the TODO at the top
    of the file only once.
    
    Updates #11041
    
    Change-Id: I07aaca27ae46b66b65780082988bdc7546ed534b
    Reviewed-on: https://go-review.googlesource.com/25112
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1e94d79f9dbe032d9c8c0f461ca6fb96e20c0b85
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Aug 16 05:57:38 2016 -0400

    cmd/compile: disable Duff's device on darwin/arm64
    
    Darwin linker does not support BR26 reloc with non-zero addend.
    
    Fixes #16724.
    
    Change-Id: I1b5b4dc7159141bde3e273490f435c08c583afaf
    Reviewed-on: https://go-review.googlesource.com/27081
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit 6fd2d2cf161ec933f206ef57b8ca6062815545d3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Aug 16 05:05:00 2016 +0000

    net/http: make Transport retry non-idempotent requests if no bytes written
    
    If the server failed on us before we even tried to write any bytes,
    it's safe to retry the request on a new connection, regardless of the
    HTTP method/idempotence.
    
    Fixes #15723
    
    Change-Id: I25360f82aac530d12d2b3eef02c43ced86e62906
    Reviewed-on: https://go-review.googlesource.com/27117
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fe27291c0039d4de0748ebd512cb236ca3c24ff6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Aug 15 21:09:39 2016 -0700

    cmd/compile: reduce garbage from autolabel
    
    Follow-up to CL 26661
    
    Change-Id: I67c58d17313094675cf0f30ce50d486818ae0dcb
    Reviewed-on: https://go-review.googlesource.com/27113
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 77e68ea78abc08f827041384e0198576da44db9f
Author: Gyu-Ho Lee <gyuhox@gmail.com>
Date:   Sat Jun 4 23:08:19 2016 -0700

    archive/tar: preallocate slice from paxHeaders
    
    Preallocate keys slice with the length of paxHeaders map
    to prevent slice growth with append operations.
    
    Change-Id: Ic9a927c4eaa775690a4ef912d61dd06f38e11510
    Reviewed-on: https://go-review.googlesource.com/23782
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2cb471e40dfe3d3ae1b0c777e67f3737a823ae79
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue Jun 28 08:41:37 2016 +0900

    crypto/tls: gofmt -w -s
    
    Change-Id: Iedf9000e3bb1fa73b4c3669eae846e85f1f5fdfe
    Reviewed-on: https://go-review.googlesource.com/24489
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5a59516dd75e32cfa22441ffe313103ff72fe796
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Aug 15 17:20:21 2016 +0000

    net/http: deflake BenchmarkClient and its use of a fixed port for testing
    
    Let the kernel pick a port for testing, and have the server in the
    child process tell the parent (benchmarking) process the port that
    was selected.
    
    Fixes flakes like seen in https://golang.org/cl/27050 (and previously)
    
    Change-Id: Ia2b705dc4152f70e0a5725015bdae09984d09d53
    Reviewed-on: https://go-review.googlesource.com/27051
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 52fcff3ec147ea8ae48c023f3d5000a8bf42fe8c
Author: Jan Mercl <0xjnml@gmail.com>
Date:   Tue Aug 2 13:00:46 2016 +0200

    go/token: Fix race in FileSet.PositionFor.
    
    Methods of FileSet are documented to be safe for concurrent use by
    multiple goroutines, so FileSet is protected by a mutex and all its
    methods use it to prevent concurrent mutations. All methods of File that
    mutate the respective FileSet, including AddLine, do also lock its
    mutex, but that does not help when PositionFor is invoked concurrently
    and reads without synchronization what AddLine mutates.
    
    The change adds acquiring a RLock around the racy call of File.position
    and the respective test.
    
    Fixes #16548
    
    Change-Id: Iecaaa02630b2532cb29ab555376633ee862315dd
    Reviewed-on: https://go-review.googlesource.com/25345
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 14e446d909784fba7b349b76ec0a234c2a86a491
Author: Carlos C <uldericofilho@gmail.com>
Date:   Wed Jul 20 00:58:55 2016 +0200

    bytes: add examples
    
    `bytes` and `strings` are pretty similar to each other, this commit
    brings `strings` examples to its counter-part.
    
    Partially addresses #16360
    
    Change-Id: I551320eaa78be9df69012035f1c3333f500e04c9
    Reviewed-on: https://go-review.googlesource.com/25062
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 35f5517ce9f43abb4bbe5601bb2697b15a440783
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Aug 15 17:28:00 2016 -0700

    doc: add 1.7 to golang.org/project
    
    Change-Id: Ib17f6643efd49e2bca188c4faa505f79832d18b1
    Reviewed-on: https://go-review.googlesource.com/27110
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7c31043ccad67e193cd1d88cb1a62b97cd2ba294
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue Aug 16 09:38:26 2016 +0900

    os/exec: fix nit found by vet
    
    Change-Id: I8085ed43d63215237a4871cc1e44257132a7f5de
    Reviewed-on: https://go-review.googlesource.com/27130
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b98d8cd5ce0c279674472af247d86f8c0b73828a
Author: Sina Siadat <siadat@gmail.com>
Date:   Tue Jun 21 00:54:52 2016 +0430

    container/heap: remove one unnecessary comparison in Fix
    
    The heap.Fix function calls both down and up.  If the element is moved
    down, we don't need to call up and we could save a comparison.
    
    (per suggestion by Radu Berinde)
    
    Fixes #16098.
    
    Change-Id: I83a74710e66cf0d274d8c0743338c26f89f31afe
    Reviewed-on: https://go-review.googlesource.com/24273
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit b5e43e669a5e1591c9a6c7157b4dd0d2796d3037
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Aug 12 10:31:17 2016 +1200

    cmd/link: when dynlinking, do not mangle short symbol names
    
    When dynamically linking, a type symbol's name is replaced with a name based on
    the SHA1 of the name as type symbol's names can be very long.  However, this
    can make a type's symbol name longer in some cases. So skip it in that case.
    One of the symbols this changes the treatment of is 'type.string' and that fixes a
    bug where -X doesn't work when dynamically linking.
    
    Fixes #16671
    
    Change-Id: If5269038261b76fb0ec52e25a9c1d64129631e3c
    Reviewed-on: https://go-review.googlesource.com/26890
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 3ddc9ad9161c6d5ae07ce2304aa838d4b853cc78
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Jul 18 17:27:21 2016 +0000

    strings: add special cases for Join of 2 and 3 strings
    
    We already had special cases for 0 and 1. Add 2 and 3 for now too.
    To be removed if the compiler is improved later (#6714).
    
    This halves the number of allocations and total bytes allocated via
    common filepath.Join calls, improving filepath.Walk performance.
    
    Noticed as part of investigating filepath.Walk in #16399.
    
    Change-Id: If7b1bb85606d4720f3ebdf8de7b1e12ad165079d
    Reviewed-on: https://go-review.googlesource.com/25005
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit c88e868030f5f0bc68c11cb65af625e77acc43ba
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Jun 7 12:26:06 2016 -0700

    cmd/internal/obj: add generated String method for AddrType
    
    Generated with:
    
    stringer -type AddrType cmd/internal/obj
    
    Change-Id: I74509cffab774035c5ca2ac0634638d73dbd33f3
    Reviewed-on: https://go-review.googlesource.com/26657
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a9ed47735f2948d7391c8a663ededeb495f7753f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 1 10:15:02 2016 -0700

    cmd/compile: move auto label gen variables to local function
    
    This still depends on Curfn, but it's progress.
    
    Updates #15756
    
    Change-Id: Ic32fe56f44fcfbc023e7668d4dee07f8b47bf3a4
    Reviewed-on: https://go-review.googlesource.com/26661
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d94409d65186a5fcee2955e374e1d5c0f457eb2b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jul 7 17:20:00 2016 -0700

    go/types: fix bad printf verbs
    
    This fixes the following vet warnings:
    
    go/types/builtins.go:437: arg call for printf verb %s of wrong type: *go/ast.CallExpr
    go/types/builtins.go:598: arg call for printf verb %s of wrong type: *go/ast.CallExpr
    
    Updates #11041
    
    Change-Id: I746d054e8e49b330fbdf961912a98f55dd5f3ff9
    Reviewed-on: https://go-review.googlesource.com/26997
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit e0d8064ed4b53436f678960022c7092b2078094d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Aug 7 09:11:00 2016 -0700

    go/types: fix multiword data structure alignment on nacl
    
    Fixes #16464
    
    Change-Id: Ibf5625c1b5fa3abd18623023f18664e8f81fa45a
    Reviewed-on: https://go-review.googlesource.com/26996
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 7a974a4c602252ca7db5b35c1236b45ead4c1d54
Author: Carlos C <uldericofilho@gmail.com>
Date:   Wed Aug 10 16:24:11 2016 +0200

    encoding/json: add example for RawMessage marshalling
    
    Fixes #16648
    
    Change-Id: I3ab21ab33ca3f41219de9518ac6a39f49131e5e5
    Reviewed-on: https://go-review.googlesource.com/26692
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3357a02b74fea42c3348de83af81900c6623f584
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Jul 5 10:33:50 2016 -0700

    math/big: use array instead of slice for deBruijn lookups
    
    This allows the compiler to remove a bounds check.
    
    math/big/nat.go:681: index bounds check elided
    math/big/nat.go:683: index bounds check elided
    
    Change-Id: Ieecb89ec5e988761b06764bd671672015cd58e9d
    Reviewed-on: https://go-review.googlesource.com/26663
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4e24e1d9996b0b0155c8349e49244d9694c89708
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Mon Aug 1 17:43:24 2016 +0300

    cmd/internal/obj/x86: VPSHUFD takes an unsigned byte.
    
    VPSHUFD should take an unsigned argument to be consistent with
    PSHUFD. Also fix all usage.
    
    Fixes #16499
    
    Change-Id: Ie699c102afed0379445914a251710365b14d89b6
    Reviewed-on: https://go-review.googlesource.com/25383
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit df9eeb192227257e8399c629372cd8cc79513406
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Aug 7 09:25:15 2016 -0700

    go/types: remove struct Sizeof cache
    
    It was not responsive to the sizes param.
    Remove it, and unwind the extra layers.
    
    Fixes #16316
    
    Change-Id: I940a57184a1601f52348d4bff8638f3f7462f5cd
    Reviewed-on: https://go-review.googlesource.com/26995
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 133d231a89c2d9f8bbceeb42cf47d8cd68fde357
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Aug 10 10:36:42 2016 -0700

    cmd/compile/internal/gc: get rid of useless autopkg variable
    
    autopkg == localpkg, so it appears to be a remnant of earlier code.
    
    Change-Id: I65b6c074535e877317cbf9f1f35e94890f0ebf14
    Reviewed-on: https://go-review.googlesource.com/26662
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3dc082f8fea3ee2710f1d1929169fb49ddf2622a
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Jun 15 15:22:35 2016 -0700

    go/types: minor cleanups
    
    1) Removed mark field from declInfo struct. Instead use a visited map
       in ordering.go which was the only use place for the mark field.
    
    2) Introduced objSet type for the common map[Object]bool type.
    
    3) Improved comments.
    
    Change-Id: I7544e7458d844b0ca08193f11de6238d317eaf2d
    Reviewed-on: https://go-review.googlesource.com/24153
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 5c84441d888655ebcc57c2ba2db834f97fa6d102
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Jun 14 17:35:36 2016 -0700

    go/types: fix computation of initialization order
    
    The old algorithm operated on a dependency graph that included
    all objects (including functions) for simplicity: it was based
    directly on the dependencies collected for each object during
    type checking an object's initialization expression. It also
    used that graph to compute the objects involved in an erroneous
    initialization cycle.
    
    Cycles that consist only of (mutually recursive) functions are
    permitted in initialization code; so those cycles were silently
    ignored if encountered. However, such cycles still inflated the
    number of dependencies a variable might have (due to the cycle),
    which in some cases lead to the wrong variable being scheduled
    for initialization before the one with the inflated dependency
    count.
    
    Correcting for the cycle when it is found is too late since at
    that point another variable may have already been scheduled.
    
    The new algorithm computes the initialization dependency graph as
    before but adds an extra pass during which functions are eliminated
    from the graph (and their dependencies are "back-propagated").
    This eliminates the problem of cycles only involving functions
    (there are no functions).
    
    When a cycle is found, the new code computes the cycle path from
    the original object dependencies so it can still include functions
    on the path as before, for the same detailed error message.
    
    The new code also more clearly distinguishes between objects that
    can be in the dependency graph (constants, variables, functions),
    and objects that cannot, by introducing the dependency type, a new
    subtype of Object. As a consequence, the dependency graph is smaller.
    
    Fixes #10709.
    
    Change-Id: Ib58d6ea65cfb279041a0286a2c8e865f11d244eb
    Reviewed-on: https://go-review.googlesource.com/24131
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 66da885594b6dbf61d93b627f5f2d5cd34cf9023
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Aug 10 19:46:48 2016 +0000

    syscall: test Gettimeofday everywhere, not just on Darwin
    
    The Darwin-only restriction was because we were late in the Go 1.7
    cycle when the test was added.
    
    In the process, I noticed Gettimeofday wasn't in the "unimplemented
    midden heap" section of syscall_nacl.go, despite this line in the
    original go1.txt:
    
    pkg syscall, func Gettimeofday(*Timeval) error
    
    So, add it, returning ENOSYS like the others.
    
    Change-Id: Id7e02e857b753f8d079bee335c22368734e92254
    Reviewed-on: https://go-review.googlesource.com/26772
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit 664c4a1f87fb48d7af6880fd9e4b504049c37b9b
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Mon Aug 15 01:39:00 2016 +0900

    os: consolidate files
    
    Code movement only.
    
    If someone finds function 'foo' in "foo_linux.go",
    they will expect that the Window version of 'foo' exists in "foo_windows.go".
    
    Current code doesn't follow this manner.
    
    For example, 'sameFile' exists in "file_unix.go",
    "stat_plan9.go" and "types_windows.go".
    
    The CL address that problem by following rules:
    
    * readdir family => dir.go, dir_$GOOS.go
    * stat family => stat.go, stat_$GOOS.go
    * path-functions => path_$GOOS.go
    * sameFile => types.go, types_$GOOS.go
    * process-functions => exec.go, exec_$GOOS.go
    * hostname => sys.go, sys_$GOOS.go
    
    Change-Id: Ic3c64663ce0b2a364d7a414351cd3c772e70187b
    Reviewed-on: https://go-review.googlesource.com/27035
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c307e1628c1569fcd589d7e1a9b3adbe367b98e6
Merge: d47bcd1 88c8b7c
Author: Gerrit Code Review <noreply-gerritcodereview@google.com>
Date:   Tue Aug 16 00:10:35 2016 +0000

    Merge "Merge remote-tracking branch 'origin/dev.ssa' into merge"

commit 88c8b7c7f98f9df718fa9cdf45f4a46726753add
Merge: 31ad583 d08010f
Author: Keith Randall <khr@golang.org>
Date:   Mon Aug 15 09:31:59 2016 -0700

    Merge remote-tracking branch 'origin/dev.ssa' into merge
    
    Merging from dev.ssa back into master.
    
    Contains complete SSA backends for arm, arm64, 386, amd64p32.
    Work in progress for PPC64.
    
    Change-Id: Ifd7075e3ec6f88f776e29f8c7fd55830328897fd

commit d47bcd157cdf0e937dbca01939d0c287e7c49acd
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Aug 15 12:57:13 2016 -0700

    doc: document go1.7
    
    Change-Id: Ieae5831b35768a625bf735a38f3d938f23f0b77b
    Reviewed-on: https://go-review.googlesource.com/27057
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1003b3e1f06935070a0000a41011898325b100ce
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Aug 15 12:01:23 2016 -0700

    doc: update version tag in source install instructions
    
    Change-Id: Id83e0371b7232b01be83640ef1e47f9026cf2a23
    Reviewed-on: https://go-review.googlesource.com/27055
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c6f19fe1ea5292f75618e3dcb2e2c4a9d1684aa6
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Aug 8 17:44:33 2016 -0700

    doc/go1.7.html: remove DRAFT
    
    Fixes #15820.
    
    Change-Id: Ia5d5237754e77774a3a6049765eea163911f41c9
    Reviewed-on: https://go-review.googlesource.com/25592
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d08010f94e06e5d5b6d01c248ba2a976b75d95a2
Author: David Chase <drchase@google.com>
Date:   Wed Aug 10 11:44:57 2016 -0700

    [dev.ssa] cmd/compile: PPC64, FP to/from int conversions.
    
    Passes ssa_test.
    
    Requires a few new instructions and some scratchpad
    memory to move data between G and F registers.
    
    Also fixed comparisons to be correct in case of NaN.
    Added missing instructions for run.bash.
    Removed some FP registers that are apparently "reserved"
    (but that are also apparently also unused except for a
    gratuitous multiplication by two when y = x+x would work
    just as well).
    
    Currently failing stack splits.
    
    Updates #16010.
    
    Change-Id: I73b161bfff54445d72bd7b813b1479f89fc72602
    Reviewed-on: https://go-review.googlesource.com/26813
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit d99cee79b98dfb6c1cd8e64c96845ee29aa28b4c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Aug 10 13:24:03 2016 -0400

    [dev.ssa] cmd/compile, etc.: more ARM64 optimizations, and enable SSA by default
    
    Add more ARM64 optimizations:
    - use hardware zero register when it is possible.
    - use shifted ops.
      The assembler supports shifted ops but not documented, nor knows
      how to print it. This CL adds them.
    - enable fast division.
      This was disabled because it makes the old backend generate slower
      code. But with SSA it generates faster code.
    
    Turn on SSA by default, also adjust tests.
    
    Change-Id: I7794479954c83bb65008dcb457bc1e21d7496da6
    Reviewed-on: https://go-review.googlesource.com/26950
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 94c8e59ae11d374cd8dd46afec4710ad10500ad9
Author: Keith Randall <khr@golang.org>
Date:   Thu Aug 11 12:56:36 2016 -0700

    [dev.ssa] cmd/compile: simplify 386+PIC+globals a bit
    
    We shouldn't issue instructions like MOVL foo(SB), AX directly from the
    SSA backend.  Instead we should do LEAL foo(SB), AX; MOVL (AX), AX.
    
    This simplifies obj logic because now only LEAL needs to be treated
    specially.  The register allocator uses the LEAL to in effect allocate
    the temporary register required for the shared library thunk calls.
    
    Also, the LEALs can now be CSEd.  So code like
        var g int
        func f() { g += 5 }
    Requires only one thunk call instead of 2.
    
    Change-Id: Ib87d465f617f73af437445871d0ea91a630b2355
    Reviewed-on: https://go-review.googlesource.com/26814
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 8f955d3664813c831b35cb02c6e7b48dd0341ece
Author: Keith Randall <khr@golang.org>
Date:   Thu Aug 11 09:49:48 2016 -0700

    [dev.ssa] cmd/compile: fix fp constant loads for 386+PIC
    
    In position-independent 386 code, loading floating-point constants from
    the constant pool requires two steps: materializing the address of
    the constant pool entry (requires calling a thunk) and then loading
    from that address.
    
    Before this CL, the materializing happened implicitly in CX, which
    clobbered that register.
    
    Change-Id: Id094e0fb2d3be211089f299e8f7c89c315de0a87
    Reviewed-on: https://go-review.googlesource.com/26811
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ed1ad8f56cc51cc55a8c12514e1c2b3098c1218b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Aug 3 09:56:36 2016 -0400

    [dev.ssa] cmd/compile: add some ARM64 optimizations
    
    Mostly mirrors ARM, includes:
    - constant folding
    - simplification of load, store, extension, and arithmetics
    - nilcheck removal
    
    Change-Id: Iffaa5fcdce100fe327429ecab316cb395e543469
    Reviewed-on: https://go-review.googlesource.com/26710
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 748aa84424418fb71c2528e7340df0ad6075b265
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Aug 3 11:47:40 2016 -0400

    [dev.ssa] cmd/internal/obj/arm64: fix encoding constant into some instructions
    
    When a constant can be encoded in a logical instruction (BITCON), do
    it this way instead of using the constant pool. The BITCON testing
    code runs faster than table lookup (using map):
    
    (on AMD64 machine, with pseudo random input)
    BenchmarkIsBitcon-4     300000000                4.04 ns/op
    BenchmarkTable-4        50000000                27.3 ns/op
    
    The equivalent C code of BITCON testing is formally verified with
    model checker CBMC against linear search of the lookup table.
    
    Also handle cases when a constant can be encoded in a MOV instruction.
    In this case, materializa the constant into REGTMP without using the
    constant pool.
    
    When constants need to be added to the constant pool, make sure to
    check whether it fits in 32-bit. If not, store 64-bit.
    
    Both legacy and SSA compiler backends are happy with this.
    
    Fixes #16226.
    
    Change-Id: I883e3069dee093a1cdc40853c42221a198a152b0
    Reviewed-on: https://go-review.googlesource.com/26631
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 31ad583ab24263f9dbcb5cbcce849eed64e74040
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jun 25 15:11:25 2016 -0700

    testing: respect benchtime on very fast benchmarks
    
    When ns/op dropped below 1, the old code
    ignored benchtime and reverted to 1s.
    
    Change-Id: I59752cef88d8d73bfd5b085f5400ae657f78504e
    Reviewed-on: https://go-review.googlesource.com/26664
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>

commit c069bc49963bd5c2c152fe600fdcb9a2b7b58f76
Author: Keith Randall <khr@golang.org>
Date:   Tue Jul 26 11:51:33 2016 -0700

    [dev.ssa] cmd/compile: implement GO386=387
    
    Last part of the 386 SSA port.
    
    Modify the x86 backend to simulate SSE registers and
    instructions with 387 registers and instructions.
    The simulation isn't terribly performant, but it works,
    and the old implementation wasn't very performant either.
    Leaving to people who care about 387 to optimize if they want.
    
    Turn on SSA backend for 386 by default.
    
    Fixes #16358
    
    Change-Id: I678fb59132620b2c47e993c1c10c4c21135f70c0
    Reviewed-on: https://go-review.googlesource.com/25271
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 77ef597f38e11e03522d1ccac34cfd39a1ca8d8e
Author: Keith Randall <khr@golang.org>
Date:   Tue Aug 9 13:58:06 2016 -0700

    [dev.ssa] cmd/compile: more fixes for 386 shared libraries
    
    Use the destination register for materializing the pc
    for GOT references also. See https://go-review.googlesource.com/c/25442/
    The SSA backend assumes CX does not get clobbered for these instructions.
    
    Mark duffzero as clobbering CX. The linker needs to clobber CX
    to materialize the address to call. (This affects the non-shared-library
    duffzero also, but hopefully forbidding one register across duffzero
    won't be a big deal.)
    
    Hopefully this is all the cases where the linker is clobbering CX
    under the hood and SSA assumes it isn't.
    
    Change-Id: I080c938170193df57cd5ce1f2a956b68a34cc886
    Reviewed-on: https://go-review.googlesource.com/26611
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit ff37d0e681a8b35c0ebff34bc1e73105f8845928
Author: David Chase <drchase@google.com>
Date:   Mon Aug 8 14:19:56 2016 -0700

    [dev.ssa] cmd/compile: PPC: FP load/store/const/cmp/neg; div/mod
    
    FP<->int conversions remain.
    
    Updates #16010.
    
    Change-Id: I38d7a4923e34d0a489935fffc4c96c020cafdba2
    Reviewed-on: https://go-review.googlesource.com/25589
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 2cbdd55d640314e37e43d7dd8e60c457846a2876
Author: Keith Randall <khr@golang.org>
Date:   Wed Aug 3 13:00:49 2016 -0700

    [dev.ssa] cmd/compile: fix PIC for SSA-generated code
    
    Access to globals requires a 2-instruction sequence on PIC 386.
    
        MOVL foo(SB), AX
    
    is translated by the obj package into:
    
        CALL getPCofNextInstructionInTempRegister(SB)
        MOVL (&foo-&thisInstruction)(tmpReg), AX
    
    The call returns the PC of the next instruction in a register.
    The next instruction then offsets from that register to get the
    address required.  The tricky part is the allocation of the
    temp register.  The legacy compiler always used CX, and forbid
    the register allocator from allocating CX when in PIC mode.
    We can't easily do that in SSA because CX is actually a required
    register for shift instructions. (I think the old backend got away
    with this because the register allocator never uses CX, only
    codegen knows that shifts must use CX.)
    
    Instead, we allow the temp register to be anything.  When the
    destination of the MOV (or LEA) is an integer register, we can
    use that register.  Otherwise, we make sure to compile the
    operation using an LEA to reference the global.  So
    
        MOVL AX, foo(SB)
    
    is never generated directly.  Instead, SSA generates:
    
        LEAL foo(SB), DX
        MOVL AX, (DX)
    
    which is then rewritten by the obj package to:
    
        CALL getPcInDX(SB)
        LEAL (&foo-&thisInstruction)(DX), AX
        MOVL AX, (DX)
    
    So this CL modifies the obj package to use different thunks
    to materialize the pc into different registers.  We use the
    registers that regalloc chose so that SSA can still allocate
    the full set of registers.
    
    Change-Id: Ie095644f7164a026c62e95baf9d18a8bcaed0bba
    Reviewed-on: https://go-review.googlesource.com/25442
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 69a755b6020e20b7c424628e9c1ba4e28d311373
Author: Keith Randall <khr@golang.org>
Date:   Mon Aug 8 11:26:25 2016 -0700

    [dev.ssa] cmd/compile: port SSA backend to amd64p32
    
    It's not a new backend, just a PtrSize==4 modification
    of the existing AMD64 backend.
    
    Change-Id: Icc63521a5cf4ebb379f7430ef3f070894c09afda
    Reviewed-on: https://go-review.googlesource.com/25586
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 392bf3a9cfee297ec106d5a67c37d8edb4c8c183
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Aug 8 16:56:22 2016 -0700

    doc/go1.7.html: update compress/flate section
    
    Updates #15810.
    
    Change-Id: Ifa7d2fd7fbfe58dff8541b18a11f007a5ff5818a
    Reviewed-on: https://go-review.googlesource.com/25591
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f3b4e785164d6047947931724f54139a42b5293e
Merge: 0484052 01dbfb8
Author: Gerrit Code Review <noreply-gerritcodereview@google.com>
Date:   Mon Aug 8 18:21:58 2016 +0000

    Merge "[dev.ssa] Merge commit 'f135c326402aaa757aa96aad283a91873d4ae124' into mergebranch" into dev.ssa

commit 0484052358ffdfb64cd533b3ea55f7c2b9d0b7bd
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Aug 4 06:57:34 2016 -0400

    [dev.ssa] cmd/compile: remove flags from regMask
    
    Reg allocator skips flag-typed values. Flag allocator uses the type
    and whether the op has "clobberFlags" set.
    
    Tested on AMD64, ARM, ARM64, 386. Passed 'toolstash -cmp' on AMD64.
    PPC64 is coded blindly.
    
    Change-Id: Ib1cc27efecef6a1bb27f7d7ed035a582660d244f
    Reviewed-on: https://go-review.googlesource.com/25480
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 01ae4b1da4fd38c0cc5f370c958aa67735fdcd40
Author: David Chase <drchase@google.com>
Date:   Thu Aug 4 11:34:43 2016 -0700

    [dev.ssa] cmd/compile: PPC64, load/store by type, shifts, divisions, bools
    
    Updates #16010.
    
    Change-Id: Ie520d64fd1c4f881f45623303ed0b7cbdf0e4764
    Reviewed-on: https://go-review.googlesource.com/25493
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit dd1d9b36c65ac71610cf69fbc425519990af7ada
Author: David Chase <drchase@google.com>
Date:   Tue Aug 2 13:17:09 2016 -0700

    [dev.ssa] cmd/compile: PPC64, add cmp->bool, some shifts, hmul
    
    Includes hmul (all widths)
    compare for boolean result and simplifications
    shift operations plus changes/additions for implementation
    (ORN, ADDME, ADDC)
    
    Also fixed a backwards-operand CMP.
    
    Change-Id: Id723c4e25125c38e0d9ab9ec9448176b75f4cdb4
    Reviewed-on: https://go-review.googlesource.com/25410
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 01dbfb81a0d444cdaa867c30c6306b756c91dea6
Merge: d2286ea f135c32
Author: Keith Randall <khr@golang.org>
Date:   Thu Aug 4 10:51:10 2016 -0700

    [dev.ssa] Merge commit 'f135c326402aaa757aa96aad283a91873d4ae124' into mergebranch
    
    Pick up shared library fix in dev.ssa.
    
    Change-Id: I5bdd0e9e0f1d6f7c14b518343ee323ed9a894b9c

commit d2286ea2843569c7d277587f3d3ef06ae4092b78
Merge: 6a1153a 50edddb
Author: Keith Randall <khr@golang.org>
Date:   Thu Aug 4 10:08:20 2016 -0700

    [dev.ssa] Merge remote-tracking branch 'origin/master' into mergebranch
    
    Semi-regular merge from tip into dev.ssa.
    
    Change-Id: Iadb60e594ef65a99c0e1404b14205fa67c32a9e9

commit 6a1153acb4aff2d481be1611862be0f9f5203e62
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Jul 1 11:05:29 2016 -0700

    [dev.ssa] cmd/compile: refactor out rulegen value parsing
    
    Previously, genMatch0 and genResult0 contained
    lots of duplication: locating the op, parsing
    the value, validation, etc.
    Parsing and validation was mixed in with code gen.
    
    Extract a helper, parseValue. It is responsible
    for parsing the value, locating the op, and doing
    shared validation.
    
    As a bonus (and possibly as my original motivation),
    make op selection pay attention to the number
    of args present.
    This allows arch-specific ops to share a name
    with generic ops as long as there is no ambiguity.
    It also detects and reports unresolved ambiguity,
    unlike before, where it would simply always
    pick the generic op, with no warning.
    
    Also use parseValue when generating the top-level
    op dispatch, to ensure its opinion about ops
    matches genMatch0 and genResult0.
    
    The order of statements in the generated code used
    to depend on the exact rule. It is now somewhat
    independent of the rule. That is the source
    of some of the generated code changes in this CL.
    See rewritedec64 and rewritegeneric for examples.
    It is a one-time change.
    
    The op dispatch switch and functions used to be
    sorted by opname without architecture. The sort
    now includes the architecture, leading to further
    generated code changes.
    See rewriteARM and rewriteAMD64 for examples.
    Again, it is a one-time change.
    
    There are no functional changes.
    
    Change-Id: I22c989183ad5651741ebdc0566349c5fd6c6b23c
    Reviewed-on: https://go-review.googlesource.com/24649
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit dede2061f3c7d593df471eb8b22b349dd310c71f
Author: David Chase <drchase@google.com>
Date:   Wed Jul 27 13:54:07 2016 -0700

    [dev.ssa] cmd/compile: PPC64, add more zeroing and moves
    
    Passes light testing.
    Modified to avoid possible exposure of "exterior" pointers
    to GC.
    
    Updates #16010.
    
    Change-Id: I41fced4fa83cefb9542dff8c8dee1a0c48056b3c
    Reviewed-on: https://go-review.googlesource.com/25310
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 00692402162ecc3df33af2b3ce48142b0ff9429c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jul 27 16:46:51 2016 -0400

    [dev.ssa] cmd/compile: fix build for old backend on ARM64
    
    Apparently the old backend needs NEG instruction having RegRead set,
    even this instruction does not take a Reg field... I don't think SSA
    uses this flag, so just leave it as it was. SSA is still happy.
    
    Fix ARM64 build on https://build.golang.org/?branch=dev.ssa
    
    Change-Id: Ia7e7f2ca217ddae9af314d346af5406bbafb68e8
    Reviewed-on: https://go-review.googlesource.com/25302
    Reviewed-by: David Chase <drchase@google.com>

commit 114c05962cd5a9924cd23f1263d08f0fd757bdb7
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jul 27 12:33:08 2016 -0400

    [dev.ssa] cmd/compile: fix possible invalid pointer spill in large Zero/Move on ARM
    
    Instead of comparing the address of the end of the memory to zero/copy,
    comparing the address of the last element, which is a valid pointer.
    Also unify large and unaligned Zero/Move, by passing alignment as AuxInt.
    
    Fixes #16515 for ARM.
    
    Change-Id: I19a62b31c5acf5c55c16a89bea1039c926dc91e5
    Reviewed-on: https://go-review.googlesource.com/25300
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 83208504fe2bbe91dae99111593de54cca1cdca0
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Jul 22 06:41:14 2016 -0400

    [dev.ssa] cmd/compile: add more on ARM64 SSA
    
    Support the following:
    - Shifts. ARM64 machine instructions only use lowest 6 bits of the
      shift (i.e. mod 64). Use conditional selection instruction to
      ensure Go semantics.
    - Zero/Move. Alignment is ensured.
    - Hmul, Avg64u, Sqrt.
    - reserve R18 (platform register in ARM64 ABI) and R29 (frame pointer
      in ARM64 ABI).
    
    Everything compiles, all.bash passed (with non-SSA test disabled).
    
    Change-Id: Ia8ed58dae5cbc001946f0b889357b258655078b1
    Reviewed-on: https://go-review.googlesource.com/25290
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 2d16e43158bb3f85eeaa983c7f12c7c81c6bc187
Author: David Chase <drchase@google.com>
Date:   Tue Jul 26 09:24:18 2016 -0700

    [dev.ssa] cmd/compile: PPC64, basic support for all calls and "miscellaneous"
    
    Added support for ClosureCall, DeferCall, InterCall
    (GoCall not yet tested).
    
    Added support for GetClosurePtr, IsNonNil, IsInBounds, IsSliceInBounds, NilCheck
    (Convert and GetG not yet tested)
    
    Still need to implement NilCheck optimizations.
    Fixed move boolean constant, order of operands to subtract.
    
    Updates #16010.
    
    Change-Id: Ibe0f6a6e688df4396cd77de0e9095997e4ca8ed2
    Reviewed-on: https://go-review.googlesource.com/25241
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 806cacc7c664ad46f3bec10a3a4e0274678eaeef
Author: David Chase <drchase@google.com>
Date:   Fri Jul 22 10:47:16 2016 -0700

    [dev.ssa] cmd/compile: replace storeconst w/ storezero, fold addressing
    
    Because PPC lacks store-immediate, remove the instruction
    that implies that it exists.  Replace it with storezero for
    the special case of storing zero, because R0 is reserved zero
    for Go (though the assembler knows this, do it in SSA).
    
    Also added address folding for storezero.
    (Now corrected to use right-sized stores in bulk-zero code.)
    
    Hello.go now compiles to
    genssa main
        00000 (...hello.go:7) TEXT "".main(SB), $0
        00001 (...hello.go:7) FUNCDATA $0, "".gcargs·0(SB)
        00002 (...hello.go:7) FUNCDATA $1, "".gclocals·1(SB)
    v23 00003 (...hello.go:8) MOVD $go.string."Hello, World!\n"(SB), R3
    v11 00004 (...hello.go:8) MOVD R3, 32(R1)
    v22 00005 (...hello.go:8) MOVD $14, R3
    v6  00006 (...hello.go:8) MOVD R3, 40(R1)
    v20 00007 (...hello.go:8) MOVD R0, 48(R1)
    v18 00008 (...hello.go:8) MOVD R0, 56(R1)
    v9  00009 (...hello.go:8) MOVD R0, 64(R1)
    v10 00010 (...hello.go:8) CALL fmt.Printf(SB)
    b2  00011 (...hello.go:9) RET
        00012 (<unknown line number>) END
    
    Updates #16010
    
    Change-Id: I33cfd98c21a1617502260ac753fa8cad68c8d85a
    Reviewed-on: https://go-review.googlesource.com/25151
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ae9570a5b95b0b321f91f504661e9c36dc1caa0e
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Jul 21 12:42:49 2016 -0400

    [dev.ssa] cmd/compile: initial ARM64 SSA port
    
    Mostly copied from ARM port, with instruction names and Prog fields
    adjusted, and 64-bit int ops added. Not complete.
    
    Fib compiles and runs correctly.
    
    Change-Id: Id3ecb0d4b571200a035344b3e8e4408769f76221
    Reviewed-on: https://go-review.googlesource.com/25130
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7bca2c599d593544d449bafc2fe1978c93b94a2f
Author: David Chase <drchase@google.com>
Date:   Wed Jul 6 13:32:52 2016 -0700

    [dev.ssa] cmd/compile: some improvements to PPC codegen
    
    Runs fibonacci for all integer types.
    Fold addressing arithmetic into stores.
    
    Updates #16010.
    
    Change-Id: I257982c82c00c80b00679757c3da345045968022
    Reviewed-on: https://go-review.googlesource.com/25103
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: David Chase <drchase@google.com>

commit df2f813bd230105e8afd6f58a3e12c3109481a3e
Author: Keith Randall <khr@golang.org>
Date:   Thu Jul 21 10:37:59 2016 -0700

    [dev.ssa] cmd/compile: 386 port now works
    
    GOARCH=386 SSATEST=1 ./all.bash passes
    
    Caveat: still needs changes to test/ files to use *_ssa.go versions.  I
    won't check those changes in with this CL because the builders will
    complain as they don't have SSATEST=1.
    
    Mostly minor fixes.
    
    Implement float <-> uint32 in assembly.  It seems the simplest option
    for now.
    
    GO386=387 does not work.  That's why I can't make SSA the default for
    386 yet.
    
    Change-Id: Ic4d4402104d32bcfb1fd612f5bb6539f9acb8ae0
    Reviewed-on: https://go-review.googlesource.com/25119
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit d8181d5d75821ad5b78ea7f4163dd86ac29f740a
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Jul 15 14:07:15 2016 -0400

    [dev.ssa] cmd/compile: simplify MOVWreg on ARM
    
    For register-register move, if there is only one use, allocate it in
    the same register so we don't need to emit an instruction.
    
    Updates #15365.
    
    Change-Id: Iad41843854a506c521d577ad93fcbe73e8de8065
    Reviewed-on: https://go-review.googlesource.com/25059
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 7b9873b9b9daba332be582cf8a9249b7430311f8
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jun 29 15:20:48 2016 -0400

    [dev.ssa] cmd/internal/obj, etc.: add and use NEGF, NEGD instructions on ARM
    
    Updates #15365.
    
    Change-Id: I372a5617c2c7d91de545cac0464809b96711b63a
    Reviewed-on: https://go-review.googlesource.com/24646
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: David Chase <drchase@google.com>

commit 4a33af6bb63eaa69a4a2cc0d4f222d37d7531b9c
Author: Keith Randall <khr@golang.org>
Date:   Mon Jul 18 15:52:59 2016 -0700

    [dev.ssa] cmd/compile: more 386 port changes
    
    Fix up zero/move code, including duff calls and rep movs.
    
    Handle the new ops generated by dec64.rules.
    
    Fix constant shifts.
    
    Change-Id: I7d89194b29b04311bfafa0fd93b9f5644af04df9
    Reviewed-on: https://go-review.googlesource.com/25033
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 1b0404c4cab18bae9c9e11d0699a1aeb32f08908
Author: Keith Randall <khr@golang.org>
Date:   Mon Jul 18 14:00:25 2016 -0700

    [dev.ssa] cmd/compile: fix verbose typing of DIV
    
    Use Cherry's awesome pair type constructor.
    
    Change-Id: I282156a570ee4dd3548bd82fbf15b8d8eb5bedf6
    Reviewed-on: https://go-review.googlesource.com/25009
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit aee8d8b9dd633f43754cdc310e5264026dc3bc42
Author: Keith Randall <khr@golang.org>
Date:   Mon Jul 18 11:51:48 2016 -0700

    [dev.ssa] cmd/compile: implement more 64-bit ops on 386
    
    add/sub/mul, plus constant input variants.
    
    Change-Id: I1c8006727c4fdf73558da0e646e7d1fa130ed773
    Reviewed-on: https://go-review.googlesource.com/25006
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit cf92e3845f2a7355d8270d419b7bcab4daf220fc
Author: Keith Randall <khr@golang.org>
Date:   Mon Jul 18 10:18:12 2016 -0700

    [dev.ssa] cmd/compile: use 2-result divide op
    
    We now allow Values to have 2 outputs.  Use that ability for amd64.
    This allows x,y := a/b,a%b to use just a single divide instruction.
    
    Update #6815
    
    Change-Id: Id70bcd20188a2dd8445e631a11d11f60991921e4
    Reviewed-on: https://go-review.googlesource.com/25004
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Chase <drchase@google.com>

commit 25e0a367da6254d89b497f392ea9d1679455d000
Author: Keith Randall <khr@golang.org>
Date:   Wed Jul 13 16:15:54 2016 -0700

    [dev.ssa] cmd/compile: clean up tuple types and selects
    
    Make tuple types and their SelectX ops fully generic.
    These ops no longer need to be lowered.
    Regalloc understands them and their tuple-generating arguments.
    We can now have opcodes returning arbitrary pairs of results.
    (And it would be easy to move to >2 results if needed.)
    
    Update arm implementation to the new standard.
    Implement just enough in 386 port to do 64-bit add.
    
    Change-Id: I370ed5aacce219c82e1954c61d1f63af76c16f79
    Reviewed-on: https://go-review.googlesource.com/24976
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6b6de15d327142a19c978c8b9811310b174fd60b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Jul 7 10:49:43 2016 -0400

    [dev.ssa] cmd/compile: support NaCl in SSA for ARM
    
    NaCl code runs in sandbox and there are restrictions for its
    instruction uses
    (https://developer.chrome.com/native-client/reference/sandbox_internals/arm-32-bit-sandbox).
    
    Like the legacy backend, on NaCl,
    - don't use R9, which is used as NaCl's "thread pointer".
    - don't use Duff's device.
    - don't use indexed load/stores.
    - the assembler rewrites DIV/MOD to runtime calls, which on NaCl
      clobbers R12, so R12 is marked as clobbered for DIV/MOD.
    - other restrictions are satisfied by the assembler.
    
    Enable SSA specific tests on nacl/arm, and disable non-SSA ones.
    
    Updates #15365.
    
    Change-Id: I9262693ec6756b89ca29d3ae4e52a96fe5403b02
    Reviewed-on: https://go-review.googlesource.com/24859
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 7d70f84f547a1b60279985fa91c407ddfde9bd64
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jul 6 10:04:45 2016 -0400

    [dev.ssa] cmd/compile: add floating point optimizations in SSA for ARM
    
    Add some simplification rules for floating point ops.
    
    cmd/internal/obj/arm supports instructions that compare FP register
    to 0, but runtime softfloat simulator does not. This CL adds these
    instructions to softfloat simulator as well.
    
    Updates #15365.
    
    Change-Id: I29405b2bfcb4c8cf106cb7a1a811409fec91b170
    Reviewed-on: https://go-review.googlesource.com/24790
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 6adb97bde72b97310f9a75a4e286cd2ef236b271
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jul 13 09:22:48 2016 -0600

    [dev.ssa] cmd/compile: fix argument size of runtime call in SSA for ARM
    
    The argument size for runtime call was incorrectly includes the size
    of LR (FixedFrameSize in general). This makes the stack frame
    sometimes unnecessarily 4 bytes larger on ARM.
    For example,
            func f(b []byte) byte { return b[0] }
    compiles to
            0x0000 00000 (h.go:6)   TEXT    "".f(SB), $4-16 // <-- framesize = 4
            0x0000 00000 (h.go:6)   MOVW    8(g), R1
            0x0004 00004 (h.go:6)   CMP     R1, R13
            0x0008 00008 (h.go:6)   BLS     52
            0x000c 00012 (h.go:6)   MOVW.W  R14, -8(R13)
            0x0010 00016 (h.go:6)   FUNCDATA        $0, gclocals·8355ad952265fec823c17fcf739bd009(SB)
            0x0010 00016 (h.go:6)   FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
            0x0010 00016 (h.go:6)   MOVW    "".b+4(FP), R0
            0x0014 00020 (h.go:6)   CMP     $0, R0
            0x0018 00024 (h.go:6)   BLS     44
            0x001c 00028 (h.go:6)   MOVW    "".b(FP), R0
            0x0020 00032 (h.go:6)   MOVBU   (R0), R0
            0x0024 00036 (h.go:6)   MOVB    R0, "".~r1+12(FP)
            0x0028 00040 (h.go:6)   MOVW.P  8(R13), R15
            0x002c 00044 (h.go:6)   PCDATA  $0, $1
            0x002c 00044 (h.go:6)   CALL    runtime.panicindex(SB)
            0x0030 00048 (h.go:6)   UNDEF
            0x0034 00052 (h.go:6)   NOP
            0x0034 00052 (h.go:6)   MOVW    R14, R3
            0x0038 00056 (h.go:6)   CALL    runtime.morestack_noctxt(SB)
            0x003c 00060 (h.go:6)   JMP     0
    
    Note that the frame size is 4, but there is actually no local. It
    incorrectly thinks call to runtime.panicindex needs 4 bytes space
    for argument.
    
    This CL fixes it.
    
    Updates #15365.
    
    Change-Id: Ic65d55283a6aa8a7861d7a3fbc7b63c35785eeec
    Reviewed-on: https://go-review.googlesource.com/24909
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 7bd88a651d0d9f8b314989475b337a0edf225bd1
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Jul 10 23:34:46 2016 -0600

    [dev.ssa] cmd/compile: don't sink spills that satisfy merge edges in SSA
    
    If a spill is used to satisfy a merge edge (in shuffle), don't sink
    it out of loop.
    
    This is found in the following code (on ARM) where there is a stack
    Phi (v268) inside a loop (b36 -> ... -> b47 -> b38 -> b36).
    
    (before shuffle)
      b36: <- b34 b38
        ...
        v268 = Phi <int> v410 v360 : autotmp_198[int]
        ...
        ... -> b47
      b47: <- b44
        ...
        v360 = ... : R6
        v230 = StoreReg <int> v360 : autotmp_198[int]
        v261 = CMPconst <flags> [0] v360
        EQ v261 -> b49 b38 (unlikely)
      b38: <- b47
        ...
        Plain -> b36
    
    During shuffle, v230 (as spill of v360) is found to satisfy v268, but
    it didn't record its use in shuffle, and v230 is sunk out of the loop
    (to b49), which leads to bad value in v268.
    
    This seems never happened on AMD64 (in make.bash), until 4 registers
    are removed.
    
    Change-Id: I01dfc28ae461e853b36977c58bcfc0669e556660
    Reviewed-on: https://go-review.googlesource.com/24858
    Reviewed-by: David Chase <drchase@google.com>

commit 8cc3f4a17e2f4d63e090fd7bd39bee697521fddf
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Jun 17 10:34:06 2016 -0400

    [dev.ssa] cmd/compile: use shifted and indexed ops in SSA for ARM
    
    This CL implements the following optimizations for ARM:
    - use shifted ops (e.g. ADD R1<<2, R2) and indexed load/stores
    - break up shift ops. Shifts used to be one SSA op that generates
      multiple instructions. We break them up to multiple ops, which
      allows constant folding and CSE for comparisons. Conditional moves
      are introduced for this.
    - simplify zero/sign-extension ops.
    
    Updates #15365.
    
    Change-Id: I55e262a776a7ef2a1505d75e04d1208913c35d39
    Reviewed-on: https://go-review.googlesource.com/24512
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 14cf6e20832dd64d79c345e0fd59169c5bd0eb35
Author: Keith Randall <khr@golang.org>
Date:   Wed Jul 13 13:43:08 2016 -0700

    [dev.ssa] cmd/compile: initial 386 SSA port
    
    Basically just copied all the amd64 files, removed all the *Q ops,
    and rebuilt.
    
    Compiles fib successfully.
    
    Still need to do:
     - all the 64->32 bit op translations.
     - audit for instructions that aren't available on 386.
     - GO386=387?
    
    Update #16358
    
    Change-Id: Ib8c684586416a554a527a5eefa0cff71424e36f5
    Reviewed-on: https://go-review.googlesource.com/24912
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit efefd11725885c850be617389faa4d9f56137641
Merge: f0bab31 76da649
Author: Keith Randall <khr@golang.org>
Date:   Wed Jul 13 11:03:12 2016 -0700

    [dev.ssa] Merge remote-tracking branch 'origin/master' into mergebranch
    
    Semi-regular merge of tip into dev.ssa.
    
    Change-Id: I855817c4746237792a2dab6eaf471087a3646be4

commit f0bab3166012ec914571cb9b7a2c780b107f1225
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Jul 5 10:03:02 2016 -0700

    [dev.ssa] cmd/compile: add some constant folding optimizations
    
    These were helpful for some autogenerated code
    I'm working with.
    
    Change-Id: I7b89c69552ca99bf560a14bfbcd6bd238595ddf6
    Reviewed-on: https://go-review.googlesource.com/24742
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8599fdd9b66a384ac1e82f301a9ff4adfe448b08
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Jun 13 16:49:09 2016 -0400

    [dev.ssa] cmd/compile: add some ARM optimization rewriting rules
    
    Mostly constant folding rules, analogous to AMD64 ones. Along with
    some simplifications.
    
    Updates #15365.
    
    Change-Id: If83bc1188bb05acb982ef3a1c21704c187e3eb24
    Reviewed-on: https://go-review.googlesource.com/24210
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 42181ad852700b790906b4c7ab24c4c23dd874e2
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Jun 7 14:18:04 2016 -0400

    [dev.ssa] cmd/compile: enable SSA on ARM by default
    
    As Josh mentioned in CL 24716, there has been requests for using SSA
    for ARM. SSA can still be disabled by setting -ssa=0 for cmd/compile,
    or partially enabled with GOSSAFUNC, GOSSAPKG, and GOSSAHASH.
    
    Not enable SSA by default on NaCl, which is not supported yet.
    
    Enable SSA-specific tests on ARM: live_ssa.go and nilptr3_ssa.go;
    disable non-SSA tests: live.go, nilptr3.go, and slicepot.go.
    
    Updates #15365.
    
    Change-Id: Ic2ca8d166aeca8517b9d262a55e92f2130683a16
    Reviewed-on: https://go-review.googlesource.com/23953
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Chase <drchase@google.com>

commit 41a7dca2722b7defafb05b0919fb8dde38819efb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Jul 3 13:40:03 2016 -0700

    [dev.ssa] cmd/compile: unify and check LoweredGetClosurePtr
    
    The comments were mostly duplicated; unify them.
    Add a check that the required invariant holds.
    
    Change-Id: I42fe09dcd1fac76d3c4e191f7a58c591c5ce429b
    Reviewed-on: https://go-review.googlesource.com/24719
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Chase <drchase@google.com>

commit ad8b8f644ea6cf99c81afc62ce4a5cc301df0ecc
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Jul 3 12:56:29 2016 -0700

    [dev.ssa] cmd/compile: remove dead amd64 ITab lowering rule
    
    ITab is handled by decomposition.
    The rule is vestigial. Remove it.
    
    Change-Id: I6fdf3d14d466761c7665c7ea14f34ca0e1e3e646
    Reviewed-on: https://go-review.googlesource.com/24718
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit f55317828b65dc35bec511e9225d9a5761a12cac
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Jun 27 16:54:57 2016 -0400

    [dev.ssa] cmd/compile: ensure alignment for Zero and Move in SSA for ARM
    
    Encode the size and the alignment into AuxInt of Zero and Move ops.
    On AMD64, we simply don't look at the alignment. On ARM and PPC64, we
    only generate aligned stores.
    
    Updates #15365.
    
    Change-Id: Ifdcc205c364f67c4516b9adebfe7d50d223b6863
    Reviewed-on: https://go-review.googlesource.com/24511
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 95427d2549a613fd1fb925ed118194488d6ec6bb
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Jun 30 11:13:24 2016 -0700

    [dev.ssa] cmd/compile: improve stability of generated code
    
    If the files in cmd/compile/internal/ssa/gen
    are passed to go run in a different order,
    e.g. due to shell differences or manual entry,
    then the order of constants in opGen churns.
    
    Sort archs by name to enforce stability.
    The movement of the PPC constants is a one time cost.
    
    Change-Id: Iebcfdb9e612d7dd8cde575f920f1292891f2f24a
    Reviewed-on: https://go-review.googlesource.com/24680
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 03d152f36fd3b64e0c0c86e93c50f5bafeb8e9b3
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Fri Jun 24 14:37:17 2016 -0500

    [dev.ssa] cmd/compile: Add initial SSA configuration for PPC64
    
    This adds the initial SSA implementation for PPC64.
    Builds golang and all.bash runs correctly.  Simple hello.go
    builds but does not run.
    
    Change-Id: I7cec211b934cd7a2dd75a6cdfaf9f71867063466
    Reviewed-on: https://go-review.googlesource.com/24453
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 68dc102ed1e2263c0c7469d6e48046eb35954a55
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jun 25 16:07:56 2016 -0700

    [dev.ssa] cmd/compile: provide default types for all extension ops
    
    Change-Id: I655327818297cc6792c81912f2cebdc321381561
    Reviewed-on: https://go-review.googlesource.com/24465
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit df43cf033f41d784a9f263b932eb8c859132eec4
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Jun 24 15:06:17 2016 -0400

    [dev.ssa] cmd/compile: optimize NilCheck in SSA for ARM
    
    Like AMD64, don't issue NilCheck instruction if the subsequent block
    has a load or store at the same address.
    
    Pass test/nilptr3_ssa.go.
    
    Updates #15365.
    
    Change-Id: Ic88780dab8c4893c57d1c95f663760cc185fe51e
    Reviewed-on: https://go-review.googlesource.com/24451
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: David Chase <drchase@google.com>

commit 8eadb89266a5a785e568f936b176d746a6d7de7c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Jun 14 11:18:39 2016 -0400

    [dev.ssa] cmd/compile: move tuple selectors to generator's block in CSE
    
    CSE may substitute a tuple generator with another one in a different
    block. In this case, since we want tuple selectors to stay together
    with the tuple generator, copy the selector to the new generator's
    block and rewrite its use.
    
    Op.isTupleGenerator and Op.isTupleSelector are introduced to assert
    tuple ops. Use it in tighten as well.
    
    Updates #15365.
    
    Change-Id: Ia9e8c734b9cc3bc9fca4a2750041eef9cdfac5a5
    Reviewed-on: https://go-review.googlesource.com/24137
    Reviewed-by: David Chase <drchase@google.com>

commit 8086ce44c4ddaba03fe0edb62aed8ca723cf0cfe
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 15 15:56:52 2016 -0700

    [dev.ssa] cmd/compile: unify OpARMMOVWaddr cases
    
    Minor code cleanup. Done as part of understanding
    OpARMMOVWaddr, since other architectures will
    need to do something similar.
    
    Change-Id: Iea2ecf3defb4f884e63902c369cd55e4647bce7a
    Reviewed-on: https://go-review.googlesource.com/24157
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 22d1318e7b6e9eba747bd90939703fff7660add1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 15 15:26:47 2016 -0700

    [dev.ssa] cmd/compile: refactor out CheckLoweredPhi
    
    This will be used verbatim in other architectures.
    
    Change-Id: I307891ae597d797fd45f296b6a38ffe9fac6b975
    Reviewed-on: https://go-review.googlesource.com/24155
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a2beee000b8e175489284abda4b82453d08af758
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 15 10:07:16 2016 -0700

    [dev.ssa] cmd/compile: improve special register error checking
    
    Provide better diagnostic messages.
    
    Use an int for numRegs comparisons,
    to avoid asking whether a uint8 is > 255.
    
    Change-Id: I33ae193ce292b24b369865abda3902c3207d7d3f
    Reviewed-on: https://go-review.googlesource.com/24135
    Reviewed-by: Keith Randall <khr@golang.org>

commit d0fa6c2f9ed213db9010d9b188fabdd14f342d6a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Jun 15 15:17:45 2016 -0700

    [dev.ssa] cmd/compile: add and use SSAReg
    
    This will be needed by other architectures as well.
    Put a cleaner encapsulation around it.
    
    Change-Id: I0ac25d600378042b2233301678e9d037e20701d8
    Reviewed-on: https://go-review.googlesource.com/24154
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 93b8aab5c992faa447893872df4f69c81444f37d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue May 31 14:01:34 2016 -0400

    [dev.ssa] cmd/compile: handle GetG on ARM
    
    Use hardware g register (R10) for GetG, allow g to appear at LHS of
    some ops.
    
    Progress on SSA backend for ARM. Now everything compiles and runs.
    
    Updates #15365.
    
    Change-Id: Icdf93585579faa86cc29b1e17ab7c90f0119fc4e
    Reviewed-on: https://go-review.googlesource.com/23952
    Reviewed-by: David Chase <drchase@google.com>

commit 0393ed8201751d58ac71288e6ef902ec4e03efde
Merge: c40dcff 53242e4
Author: Keith Randall <khr@golang.org>
Date:   Tue Jun 14 07:33:48 2016 -0700

    [dev.ssa] Merge remote-tracking branch 'origin/master' into mergebranch
    
    Change-Id: Idd150294aaeced0176b53d6b95852f5d21ff4fdc

commit c40dcff2f23c75e15785687eee7d56b524c54634
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Jun 6 22:36:45 2016 -0400

    [dev.ssa] cmd/compile: use MOVWaddr for address on ARM
    
    Introduce an op MOVWaddr for addresses on ARM, instead of overuse
    ADDconst.
    
    Mark MOVWaddr as rematerializable. This fixes a liveness problem: if
    it were not rematerializable, the address of a variable may be spilled
    and later use of the address may just load the spilled value without
    mentioning the variable, and the liveness code may think it is dead
    prematurely.
    
    Update #15365.
    
    Change-Id: Ib0b0fa826bdb75c9e6bb362b95c6cf132cc6b1c0
    Reviewed-on: https://go-review.googlesource.com/23942
    Reviewed-by: David Chase <drchase@google.com>

commit e3a6d00876488c70b6ecc641954ed5cecb918cac
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jun 8 10:12:30 2016 -0400

    [dev.ssa] cmd/compile: ensure OffPtr has pointer type
    
    SSA treats SP as constant throughout a function, so as OffPtr [off] SP.
    When the stack moves, spilled OffPtr values become invalid, if they are
    not pointer-typed.
    
    (Currently it is fine because of the optimization rules that folds OffPtr
    into Load/Store. But it'd better be "optimization", not requirement.)
    
    Updates #15365.
    
    Change-Id: I76cf4008dfdc169e1cb5a55a2605b6678efc915d
    Reviewed-on: https://go-review.googlesource.com/23941
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit fa54bf16e0080296487407c8cc883a1e039c31c8
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Jun 3 18:03:29 2016 -0400

    [dev.ssa] cmd/compile: fix a few bugs for SSA for ARM
    
    - 64x signed right shift was wrong for shift larger than 0x80000000.
    - for Lsh-followed-by-Rsh, the intermediate value should be full int
      width, so when it is spilled MOVW should be used.
    - use RET for RetJmp, so the assembler can take case of restoring LR
      for non-leaf case.
    - reserve R9 in dynlink mode. R9 is used for GOT by the assembler.
    
    Progress on SSA backend for ARM. Still not complete.
    
    Updates #15365.
    
    Change-Id: I3caca256b92ff7cf96469da2feaf4868a592efc5
    Reviewed-on: https://go-review.googlesource.com/23793
    Reviewed-by: David Chase <drchase@google.com>

commit 225ef76c250fc9ab9794fd723952209e2ff440aa
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jun 1 06:41:08 2016 -0400

    [dev.ssa] cmd/compile: fix scheduling of tuple ops
    
    We want tuple-reading ops immediately follow tuple-generating op, so
    that tuple values will not be spilled/copied.
    
    The mechanism introduced in the previous CL cannot really avoid tuples
    interleaving. In this CL we always emit tuple and their selectors together.
    Maybe remove the tuple scores if it does not help on performance (todo).
    
    Also let tighten not move tuple-reading ops across blocks.
    
    In the previous CL a special case of regenerating flags with tuple-reading
    pseudo-op is added, but it did not cover end-of-block case. This is fixed
    in this CL and the condition is generalized.
    
    Progress on SSA backend for ARM. Still not complete.
    
    Updates #15365.
    
    Change-Id: I8980b34e7a64eb98153540e9e19a3782e20406ff
    Reviewed-on: https://go-review.googlesource.com/23792
    Reviewed-by: David Chase <drchase@google.com>

commit 59e11d782717407fcdf288664a48beb52336d42a
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue May 31 11:27:16 2016 -0400

    [dev.ssa] cmd/compile: handle floating point on ARM
    
    Machine supports (or the runtime simulates in soft float mode)
    (u)int32<->float conversions. The frontend rewrites int64<->float
    conversions to call to runtime function.
    
    For int64->float32 conversion, the frontend generates
    
    .   .   AS u(100) l(10) tc(1)
    .   .   .   NAME-main.~r1 u(1) a(true) g(1) l(9) x(8+0) class(PPARAMOUT) f(1) float32
    .   .   .   CALLFUNC u(100) l(10) tc(1) float32
    .   .   .   .   NAME-runtime.int64tofloat64 u(1) a(true) x(0+0) class(PFUNC) tc(1) used(true) FUNC-func(int64) float64
    
    The CALLFUNC node has type float32, whereas runtime.int64tofloat64
    returns float64. The legacy backend implicitly makes a float64->float32
    conversion. The SSA backend does not do implicit conversion, so we
    insert an explicit CONV here.
    
    All cmd/compile/internal/gc/testdata/*_ssa.go tests passed.
    
    Progress on SSA for ARM. Still not complete.
    
    Update #15365.
    
    Change-Id: I30937c8ff977271246b068f48224693776804339
    Reviewed-on: https://go-review.googlesource.com/23652
    Reviewed-by: Keith Randall <khr@golang.org>

commit e78d90beebbb6fde602ceb3999535ac4b49da295
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed May 25 09:49:28 2016 -0400

    [dev.ssa] cmd/compile: handle Div, Convert, GetClosurePtr etc. on ARM
    
    This CL adds support of Div, Mod, Convert, GetClosurePtr and 64-bit indexing
    support to SSA backend for ARM.
    
    Add tests for 64-bit indexing to cmd/compile/internal/gc/testdata/string_ssa.go.
    
    Tests cmd/compile/internal/gc/testdata/*_ssa.go passed, except compound_ssa.go
    and fp_ssa.go.
    
    Progress on SSA for ARM. Still not complete. Essentially the only unsupported
    part is floating point.
    
    Updates #15365.
    
    Change-Id: I269e88b67f641c25e7a813d910c96d356d236bff
    Reviewed-on: https://go-review.googlesource.com/23542
    Reviewed-by: David Chase <drchase@google.com>

commit 4636d02244e683a0d9c078a49b4c614bed401d6b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed May 25 23:17:42 2016 -0400

    [dev.ssa] cmd/compile: handle 64-bit shifts on ARM
    
    Also fix a mistake in previous CL about x8 and x16 shifts:
    the shift needs ZeroExt.
    
    Progress on SSA for ARM. Still not complete.
    
    Updates #15365.
    
    Change-Id: Ibc352760023d38bc6b9c5251e929fe26e016637a
    Reviewed-on: https://go-review.googlesource.com/23486
    Reviewed-by: David Chase <drchase@google.com>

commit 90883091ff5f7170c83f847d6748cf36713e8c9b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu May 19 12:33:30 2016 -0400

    [dev.ssa] cmd/compile: clean up hardcoded regmasks in ssa/regalloc.go
    
    Auto-generate register masks and load them through Config.
    
    Passed toolstash -cmp on AMD64.
    
    Tests phi_ssa.go and regalloc_ssa.go in cmd/compile/internal/gc/testdata
    passed on ARM.
    
    Updates #15365.
    
    Change-Id: I393924d68067f2dbb13dab82e569fb452c986593
    Reviewed-on: https://go-review.googlesource.com/23292
    Reviewed-by: David Chase <drchase@google.com>

commit 8756d9253f56f28167543fbd41c15e5695e654b2
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed May 18 18:14:36 2016 -0400

    [dev.ssa] cmd/compile: decompose 64-bit integer on ARM
    
    Introduce dec64 rules to (generically) decompose 64-bit integer on
    32-bit architectures. 64-bit integer is composed/decomposed with
    Int64Make/Hi/Lo ops, as for complex types.
    
    The idea of dealing with Add64 is the following:
    
    (Add64 (Int64Make xh xl) (Int64Make yh yl))
    ->
    (Int64Make
            (Add32withcarry xh yh (Select0 (Add32carry xl yl)))
            (Select1 (Add32carry xl yl)))
    
    where Add32carry returns a tuple (flags,uint32). Select0 and Select1
    read the first and the second component of the tuple, respectively.
    The two Add32carry will be CSE'd.
    
    Similarly for multiplication, Mul32uhilo returns a tuple (hi, lo).
    
    Also add support of KeepAlive, to fix build after merge.
    
    Tests addressed_ssa.go, array_ssa.go, break_ssa.go, chan_ssa.go,
    cmp_ssa.go, ctl_ssa.go, map_ssa.go, and string_ssa.go in
    cmd/compile/internal/gc/testdata passed.
    
    Progress on SSA for ARM. Still not complete.
    
    Updates #15365.
    
    Change-Id: I7867c76785a456312de5d8398a6b3f7ca5a4f7ec
    Reviewed-on: https://go-review.googlesource.com/23213
    Reviewed-by: Keith Randall <khr@golang.org>

commit 31e13c83c26c5addad6c9a15a8f06a11edc7c519
Merge: d108bc0 496cf21
Author: David Chase <drchase@google.com>
Date:   Fri May 27 15:18:49 2016 -0400

    [dev.ssa] Merge branch 'master' into dev.ssa
    
    Change-Id: Iabc80b6e0734efbd234d998271e110d2eaad41dd

commit d108bc0e73046826b27a3a3f6cfc1b33b40ae11d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun May 15 00:12:56 2016 -0400

    [dev.ssa] cmd/compile: implement Defer, RetJmp on SSA for ARM
    
    Also fix argument offset for runtime calls.
    
    Also fix LoadReg/StoreReg by generating instructions by type.
    
    Progress on SSA backend for ARM. Still not complete.
    Tests append_ssa.go, assert_ssa.go, loadstore_ssa.go, short_ssa.go, and
    deferNoReturn.go in cmd/compile/internal/gc/testdata passed.
    
    Updates #15365.
    
    Change-Id: I0f0a2398cab8bbb461772a55241a16a7da2ecedf
    Reviewed-on: https://go-review.googlesource.com/23212
    Reviewed-by: David Chase <drchase@google.com>

commit 8357ec37ae6a7580e928dbabbb99dd6cf1958017
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri May 13 15:31:14 2016 -0400

    [dev.ssa] cmd/compile: implement Zero, Move, Copy for SSA on ARM
    
    Generate load/stores for small zeroing/move, DUFFZERO/DUFFCOPY for
    medium zeroing/move, and loops for large zeroing/move.
    
    cmd/compile/internal/gc/testdata/{copy_ssa.go,zero_ssa.go} tests
    passed.
    
    Progress on SSA backend for ARM. Still not complete. A few packages
    in the standard library compile and tests passed, including
    container/list, hash/crc32, unicode/utf8, etc.
    
    Updates #15365.
    
    Change-Id: Ieb4b68b44ee7de66bf7b68f5f33a605349fcc6fa
    Reviewed-on: https://go-review.googlesource.com/23097
    Reviewed-by: Keith Randall <khr@golang.org>

commit 8f726907113e6be6dd886d6a790619b2535330fa
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri May 13 15:22:56 2016 -0400

    [dev.ssa] cmd/compile: implement shifts & multiplications for SSA on ARM
    
    Implement shifts and multiplications for up to 32-bit values.
    
    Also handle Exit block.
    
    Progress on SSA backend for ARM. Still not complete.
    container/heap, crypto/subtle, hash/adler32 packages compile and
    tests passed.
    
    Updates #15365.
    
    Change-Id: I6bee4d5b0051e51d5de97e8a1938c4b87a36cbf8
    Reviewed-on: https://go-review.googlesource.com/23096
    Reviewed-by: Keith Randall <khr@golang.org>

commit ccaed50c7bf6381275d49adcf54974441752fd11
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri May 13 11:25:07 2016 -0400

    [dev.ssa] cmd/compile: handle boolean values for SSA on ARM
    
    Fix hardcoded flag register mask in ssa/flagalloc.go by auto-generating
    the mask.
    
    Also fix a mistake (in previous CL) about conditional branches.
    
    Progress on SSA backend for ARM. Still not complete. Now "container/ring"
    package compiles and tests passed.
    
    Updates #15365.
    
    Change-Id: Id7c8805c30dbb8107baedb485ed0f71f59ed6ea8
    Reviewed-on: https://go-review.googlesource.com/23093
    Reviewed-by: Keith Randall <khr@golang.org>

commit e2848de9efc599e1af54079ab1f8e79e0e26764c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri May 6 10:13:31 2016 -0700

    [dev.ssa] cmd/compile: implement the following for SSA on ARM
    
    - generic Ops: Phi, CALL variants, NilCheck
    - generic Blocks: Plain, Check
    - 32-bit arithmetics
    - CMP and conditional branches
    - load/store
    - zero/sign-extensions (8 to 16, 8 to 32, 16 to 32)
    
    Progress on SSA backend for ARM. Still not complete. Now "errors"
    package compiles and tests passed.
    
    Updates #15365.
    
    Change-Id: If126fd17f8695cbf55d64085bb3f1a4a53205701
    Reviewed-on: https://go-review.googlesource.com/22856
    Reviewed-by: Keith Randall <khr@golang.org>

commit fdc4a964d24f0c975e2db2cf16a53327ad36b24d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri May 6 10:12:57 2016 -0700

    [dev.ssa] cmd/compile/internal/gc, runtime: use 32-bit load for writeBarrier check
    
    Use 32-bit load for writeBarrier check on all architectures.
    Padding added to runtime structure.
    
    Updates #15365, #15492.
    
    Change-Id: I5d3dadf8609923fe0fe4fcb384a418b7b9624998
    Reviewed-on: https://go-review.googlesource.com/22855
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e1a2ea88d02280e0661d66dc7483c95b39e18646
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri May 6 10:11:41 2016 -0700

    [dev.ssa] cmd/compile: handle symbolic constant for SSA on ARM
    
    Progress on SSA backend for ARM. Still not complete. Now "helloworld"
    function compiles and runs.
    
    Updates #15365.
    
    Change-Id: I02f66983cefdf07a6aed262fb4af8add464d8e9a
    Reviewed-on: https://go-review.googlesource.com/22854
    Reviewed-by: Keith Randall <khr@golang.org>

commit 802966f7b3827d0e462742293c259a2c1408595d
Merge: ab150e1 9b05ae6
Author: Keith Randall <khr@golang.org>
Date:   Thu May 5 14:24:16 2016 -0700

    [dev.ssa] Merge remote-tracking branch 'origin/master' into mergebranch
    
    Merge from tip into ssa.
    
    Change-Id: Icbc1c46d9f4721e4a0f99a24dd708044407ee9f7

commit ab150e1ac93b720b4591d618e4fe35988044ee54
Author: Keith Randall <khr@golang.org>
Date:   Thu May 5 13:05:16 2016 -0700

    [dev.ssa] all: merge from tip to get dev.ssa current
    
    So we can start working on other architectures here.
    
    Change is a dummy to keep git happy.
    
    Change-Id: I1caa62a242790601810a1ff72af7ea9773d4da76
    Reviewed-on: https://go-review.googlesource.com/22822
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
```
