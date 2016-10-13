# October 13, 2016 Report

Number of commits: 113

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 544.829483ms to 553.442567ms, +1.58%
* github.com/gogits/gogs: from 12.8953699s to 12.421203394s, -3.68%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 37.618453151s to 33.008875159s, -12.25%
* github.com/influxdata/influxdb/cmd/influxd: from 6.518353717s to 6.935917576s, +6.41%
* github.com/junegunn/fzf/src/fzf: from 1.150302967s to 999.566549ms, -13.10%
* github.com/mholt/caddy/caddy: from 5.845098924s to 6.090833097s, +4.20%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.225447717s to 1.399797789s, +14.23%
* github.com/nsqio/nsq/apps/nsqd: from 1.932091122s to 2.068474627s, +7.06%
* github.com/prometheus/prometheus/cmd/prometheus: from 11.59941787s to 11.647932067s, +0.42%
* github.com/spf13/hugo: from 7.94061088s to 7.645704151s, -3.71%
* golang.org/x/tools/cmd/guru: from 2.620725592s to 2.543048275s, -2.96%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2595413 to 2542636, -2.03%
* github.com/gogits/gogs: from 23557599 to 23252185, -1.30%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 29292816 to 29130136, -0.56%
* github.com/influxdata/influxdb/cmd/influxd: from 16508307 to 16260591, -1.50%
* github.com/junegunn/fzf/src/fzf: from 3158912 to 3077808, -2.57%
* github.com/mholt/caddy/caddy: from 14944471 to 14683761, -1.74%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4384116 to 4297359, -1.98%
* github.com/nsqio/nsq/apps/nsqd: from 9956471 to 9797434, -1.60%
* github.com/prometheus/prometheus/cmd/prometheus: from 25430957 to 25172578, -1.02%
* github.com/spf13/hugo: from 16361534 to 16116853, -1.50%
* golang.org/x/tools/cmd/guru: from 7523237 to 7340867, -2.42%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               188           188           +0.00%
BenchmarkMsgpUnmarshal-4             375           373           -0.53%
BenchmarkEasyJsonMarshal-4           1446          1436          -0.69%
BenchmarkEasyJsonUnmarshal-4         1555          1785          +14.79%
BenchmarkFlatBuffersMarshal-4        346           411           +18.79%
BenchmarkFlatBuffersUnmarshal-4      274           270           -1.46%
BenchmarkGogoprotobufMarshal-4       156           157           +0.64%
BenchmarkGogoprotobufUnmarshal-4     242           239           -1.24%

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

* [cmd/compile: fix choice of phi building algorithm](https://github.com/golang/go/commit/433be563b6246eb132aed6e9e58f46a0d46f7010)
* [cmd/compile,runtime: redo how map assignments work](https://github.com/golang/go/commit/442de98c14d49bf306ab880e9f9c898ca0ae7c19)

## GIT Log

```
commit dc46b882d583a32c7de37a7cfa8b423eb58d0296
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Oct 13 02:20:51 2016 -0700

    cmd/compile: stop allocating Name and Param for ODCLFIELD
    
    name       old time/op     new time/op     delta
    Template       349ms ± 5%      339ms ± 7%  -2.89%        (p=0.000 n=27+29)
    Unicode        187ms ±11%      182ms ±11%  -2.77%        (p=0.039 n=29+29)
    GoTypes        1.05s ± 3%      1.04s ± 4%    ~           (p=0.103 n=29+29)
    Compiler       4.57s ± 3%      4.55s ± 3%    ~           (p=0.202 n=30+29)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        510M ±21%       521M ±18%    ~           (p=0.281 n=30+29)
    Unicode         303M ±34%       300M ±28%    ~           (p=0.592 n=30+30)
    GoTypes        1.52G ± 9%      1.50G ± 9%    ~           (p=0.314 n=30+30)
    Compiler       6.50G ± 5%      6.44G ± 5%    ~           (p=0.362 n=29+30)
    
    name       old alloc/op    new alloc/op    delta
    Template      44.7MB ± 0%     44.0MB ± 0%  -1.63%        (p=0.000 n=28+28)
    Unicode       34.6MB ± 0%     34.5MB ± 0%  -0.18%        (p=0.000 n=30+29)
    GoTypes        125MB ± 0%      123MB ± 0%  -1.14%        (p=0.000 n=30+30)
    Compiler       515MB ± 0%      513MB ± 0%  -0.52%        (p=0.000 n=30+30)
    
    name       old allocs/op   new allocs/op   delta
    Template        427k ± 0%       416k ± 0%  -2.66%        (p=0.000 n=30+30)
    Unicode         323k ± 0%       322k ± 0%  -0.28%        (p=0.000 n=30+30)
    GoTypes        1.21M ± 0%      1.18M ± 0%  -1.84%        (p=0.000 n=29+30)
    Compiler       4.40M ± 0%      4.36M ± 0%  -0.95%        (p=0.000 n=30+30)
    
    Passes toolstash -cmp.
    
    Change-Id: Ifee7d012b1cddadda01450e027eef8d4ecf5581f
    Reviewed-on: https://go-review.googlesource.com/30980
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 01bf5cc21912ff8642171d8255a7fff87f1da00b
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Oct 12 17:39:28 2016 -0700

    cmd/compile: cleanup toolstash hacks from previous CL
    
    Change-Id: I36cf3523e00b80e2d3a690f251edd5d6f665d156
    Reviewed-on: https://go-review.googlesource.com/30975
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit add3ff549a97adbd7715b5609deb4dcbf5e7624c
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Oct 12 15:48:18 2016 -0700

    cmd/compile: add OSTRUCTKEY for keyed struct literals
    
    Previously, we used OKEY nodes to represent keyed struct literal
    elements. The field names were represented by an ONAME node, but this
    is clumsy because it's the only remaining case where ONAME was used to
    represent a bare identifier and not a variable.
    
    This CL introduces a new OSTRUCTKEY node op for use in struct
    literals. These ops instead store the field name in the node's own Sym
    field. This is similar in spirit to golang.org/cl/20890.
    
    Significant reduction in allocations for struct literal heavy code
    like package unicode:
    
    name       old time/op     new time/op     delta
    Template       345ms ± 6%      341ms ± 6%     ~           (p=0.141 n=29+28)
    Unicode        200ms ± 9%      184ms ± 7%   -7.77%        (p=0.000 n=29+30)
    GoTypes        1.04s ± 3%      1.05s ± 3%     ~           (p=0.096 n=30+30)
    Compiler       4.47s ± 9%      4.49s ± 6%     ~           (p=0.890 n=29+29)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        523M ±13%       516M ±17%     ~           (p=0.400 n=29+30)
    Unicode         334M ±27%       314M ±30%     ~           (p=0.093 n=30+30)
    GoTypes        1.53G ±10%      1.52G ±10%     ~           (p=0.572 n=30+30)
    Compiler       6.28G ± 7%      6.34G ±11%     ~           (p=0.300 n=30+30)
    
    name       old alloc/op    new alloc/op    delta
    Template      44.5MB ± 0%     44.4MB ± 0%   -0.35%        (p=0.000 n=27+30)
    Unicode       39.2MB ± 0%     34.5MB ± 0%  -11.79%        (p=0.000 n=26+30)
    GoTypes        125MB ± 0%      125MB ± 0%   -0.12%        (p=0.000 n=29+30)
    Compiler       515MB ± 0%      515MB ± 0%   -0.10%        (p=0.000 n=29+30)
    
    name       old allocs/op   new allocs/op   delta
    Template        426k ± 0%       424k ± 0%   -0.39%        (p=0.000 n=29+30)
    Unicode         374k ± 0%       323k ± 0%  -13.67%        (p=0.000 n=29+30)
    GoTypes        1.21M ± 0%      1.21M ± 0%   -0.14%        (p=0.000 n=29+29)
    Compiler       4.40M ± 0%      4.39M ± 0%   -0.13%        (p=0.000 n=29+30)
    
    Passes toolstash/buildall.
    
    Change-Id: Iba4ee765dd1748f67e52fcade1cd75c9f6e13fa9
    Reviewed-on: https://go-review.googlesource.com/30974
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 032e2bd1eb215cdeec605b33e42878ec7186cb53
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Oct 12 23:28:36 2016 -0700

    cmd/compile: replace aindex with typArray
    
    aindex is overkill when it's only ever used with known integer
    constants, so just use typArray directly instead.
    
    Change-Id: I43fc14e604172df859b3ad9d848d219bbe48e434
    Reviewed-on: https://go-review.googlesource.com/30979
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1af769da8260b277bd5aa92b5074b3400b1f8d9d
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed Sep 21 11:19:36 2016 +1000

    os: make readConsole handle its input and output correctly
    
    This CL introduces first test for readConsole. And new test
    discovered couple of problems with readConsole.
    
    Console characters consist of multiple bytes each, but byte blocks
    returned by syscall.ReadFile have no character boundaries. Some
    multi-byte characters might start at the end of one block, and end
    at the start of next block. readConsole feeds these blocks to
    syscall.MultiByteToWideChar to convert them into utf16, but if some
    multi-byte characters have no ending or starting bytes, the
    syscall.MultiByteToWideChar might get confused. Current version of
    syscall.MultiByteToWideChar call will make
    syscall.MultiByteToWideChar ignore all these not complete
    multi-byte characters.
    
    The CL solves this issue by changing processing from "randomly
    sized block of bytes at a time" to "one multi-byte character at a
    time". New readConsole code calls syscall.ReadFile to get 1 byte
    first. Then it feeds this byte to syscall.MultiByteToWideChar.
    The new syscall.MultiByteToWideChar call uses MB_ERR_INVALID_CHARS
    flag to make syscall.MultiByteToWideChar return error if input is
    not complete character. If syscall.MultiByteToWideChar returns
    correspondent error, we read another byte and pass 2 byte buffer
    into syscall.MultiByteToWideChar, and so on until success.
    
    Old readConsole code would also sometimes return no data if user
    buffer was smaller then uint16 size, which would confuse callers
    that supply 1 byte buffer. This CL fixes that problem too.
    
    Fixes #17097
    
    Change-Id: I88136cdf6a7bf3aed5fbb9ad2c759b6c0304ce30
    Reviewed-on: https://go-review.googlesource.com/29493
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 0a0f4bc181835882717d2f4be75fced4061d4572
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Wed Oct 12 11:34:47 2016 +0900

    cmd/compile/internal/gc: cleanup esc.go
    
    * convert important functions to methods
    * rename EscXXX to XXX in NodeEscState
    * rename local variables more friendly
    * simplify redundant code
    * update comments
    
    Change-Id: I8442bf4f8dde84523d9a2ad3d04b1cd326bd4719
    Reviewed-on: https://go-review.googlesource.com/30893
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 4940a8379096a48af682c266a0e104f249a80816
Author: Alex Browne <stephenalexbrowne@gmail.com>
Date:   Sat Nov 7 23:54:41 2015 -0500

    cmd/vet: check for duplicate json, xml struct field tags
    
    It is easy to make the mistake of duplicating json struct field
    tags especially when copy/pasting. This commit causes go vet to
    report the mistake. Only field tags in the same struct type are
    considered, because that is the only case which is undoubtedly an
    error.
    
    Fixes #12791.
    
    Change-Id: I4130e4c04b177694cc0daf8f1acaf0751d4f062b
    Reviewed-on: https://go-review.googlesource.com/16704
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit ab019da7272285fb4f634e7fca00c3fa973c76c4
Author: Michael Pratt <mpratt@google.com>
Date:   Wed Feb 17 23:06:56 2016 -0800

    cmd/internal/obj: document Prog
    
    Change-Id: Iafc392ba06452419542ec85e91d44991839eb6f8
    Reviewed-on: https://go-review.googlesource.com/19593
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 0a55a16c91c5bed397a0716cb9d67f8b0a71de9b
Author: Shenghou Ma <minux@golang.org>
Date:   Mon May 4 01:23:56 2015 -0400

    cmd/objdump: enable tests on ppc64/ppc64le
    
    Fixes #9039.
    
    Change-Id: I7d213b4f8e4cda73ea7687fb97dbd22e58163949
    Reviewed-on: https://go-review.googlesource.com/9683
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 94cf54e8619e03671b33113d39306646235feb60
Author: Shenghou Ma <minux@golang.org>
Date:   Mon May 4 01:21:27 2015 -0400

    cmd/internal/objfile: add ppc64/ppc64le disassembler support
    
    Change-Id: I7d213b4f8e4cda73ea7687fb97dbd22e58163948
    Reviewed-on: https://go-review.googlesource.com/9682
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit ed0a9567469e63f7d47e399f7b300b731d5c3489
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 12 12:41:58 2016 -0400

    cmd: add golang.org/x/arch/ppc64/ppc64asm for disassembly
    
    For #9039.
    
    Change-Id: I2b1bcd76857ff332411ca21a0cc5def3097a8eaf
    Reviewed-on: https://go-review.googlesource.com/30936
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit daa121167b6ce630aba00195f1c3872cda39a50c
Author: Allan Simon <allan.simon@supinfo.com>
Date:   Sun Oct 11 04:16:58 2015 +0800

    encoding/xml: prevent omitempty from omitting non-nil pointers to empty values
    
    There was an inconsistency between the (json encoding + documentation)
    and the xml encoding implementation. Pointer to an empty value was
    not being serialized (i.e simply ignored). Which had the effect of making
    impossible to have a struct with a string field for which we wanted to
    serialize the value ""
    
    Fixes #5452
    
    Change-Id: Id858701801158409be01e962d2cda843424bd22a
    Reviewed-on: https://go-review.googlesource.com/15684
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 347259cbae30f876be11fe5e71710969afa24374
Author: Xia Bin <snyh@snyh.org>
Date:   Sun Jan 24 14:22:54 2016 +0800

    misc/cgo/test: add test that gccgo fails
    
    Gccgo isn't locking the OS thread properly during calls.
    
    Change-Id: Idb2475291405e390cbb83abb27a402fd0381d0c4
    Reviewed-on: https://go-review.googlesource.com/18882
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 90a750857c6146ce5d8d24cda5a1e70aa3d4cc58
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 12 15:10:15 2016 -0400

    syscall: update darwin/amd64 for timespec change
    
    Change-Id: I74f47f519dfee10cd079ad9a4e09e36e8d74c6dc
    Reviewed-on: https://go-review.googlesource.com/30937
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 14e545b60a8c78aa6609d807b4f50e54e9bfe1eb
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Wed Oct 12 11:36:18 2016 -0700

    archive/tar: reduce allocations in formatOctal
    
    Change-Id: I9ddb7d2a97d28aba7a107b65f278993daf7807fa
    Reviewed-on: https://go-review.googlesource.com/30960
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6da8bdd2cc7a10f037a2025ffed57627d97a990c
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Wed Oct 12 07:50:30 2016 -0500

    cmd/asm: recognize CR1-CR7 on ppc64x branch instructions
    
    Some of the branch instructions (BEQ, BNE, BLT, etc.) accept
    all the valid CR values as operands, but the CR register value is
    not parsed and not put into the instruction, so that CR0 is always
    used regardless of what was specified on the instruction.  For example
    BEQ CR2,label becomes beq cr0,label.
    
    This adds the change to the PPC64 assembler to recognize the CR value
    and set the approppriate field in the instruction so the correct
    CR is used.  This also adds some general comments on the branch
    instruction BC and its operand values.
    
    Fixes #17408
    
    Change-Id: I8e956372a42846a4c09a7259e9172eaa29118e71
    Reviewed-on: https://go-review.googlesource.com/30930
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 442de98c14d49bf306ab880e9f9c898ca0ae7c19
Author: Keith Randall <khr@golang.org>
Date:   Tue Oct 11 08:36:38 2016 -0700

    cmd/compile,runtime: redo how map assignments work
    
    To compile:
      m[k] = v
    instead of:
      mapassign(maptype, m, &k, &v), do
    do:
      *mapassign(maptype, m, &k) = v
    
    mapassign returns a pointer to the value slot in the map.  It is just
    like mapaccess except that it will allocate a new slot if k is not
    already present in the map.
    
    This makes map accesses faster but potentially larger (codewise).
    
    It is faster because the write into the map is done when the compiler
    knows the concrete type, so it can be done with a few store
    instructions instead of calling typedmemmove.  We also potentially
    avoid stack temporaries to hold v.
    
    The code can be larger when the map has pointers in its value type,
    since there is a write barrier call in addition to the mapassign call.
    That makes the code at the callsite a bit bigger (go binary is 0.3%
    bigger).
    
    This CL is in preparation for doing operations like m[k] += v with
    only a single runtime call.  That will roughly double the speed of
    such operations.
    
    Update #17133
    Update #5147
    
    Change-Id: Ia435f032090a2ed905dac9234e693972fe8c2dc5
    Reviewed-on: https://go-review.googlesource.com/30815
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 55ef67f2f85c51d415a030ae144a0b3301a097bd
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 12 13:11:02 2016 -0400

    unicode: change SimpleFold to handle invalid runes
    
    Functions like ToLower and ToUpper return the invalid rune back,
    so we might as well do the same here.
    
    I changed my mind about panicking when I tried to document the behavior.
    
    Fixes #16690 (again).
    
    Change-Id: If1c68bfcd66daea160fd19948e7672b0e1add106
    Reviewed-on: https://go-review.googlesource.com/30935
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 413afcafbfd45b6b58023a49484d8480143960e0
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed Oct 12 12:45:16 2016 -0400

    cmd/link: force external linking for plugins
    
    Fixes #17415
    
    Change-Id: I6f896d549092e5e0dba72351e5385992b4cbe90f
    Reviewed-on: https://go-review.googlesource.com/30933
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9e98e7e6686a500f87352e3721038d194b1db33c
Author: Adam Langley <agl@golang.org>
Date:   Tue Oct 11 15:08:35 2016 -0700

    crypto/tls: enable X25519 by default.
    
    Since this changes the offered curves in the ClientHello, all the test
    data needs to be updated too.
    
    Change-Id: I227934711104349c0f0eab11d854e5a2adcbc363
    Reviewed-on: https://go-review.googlesource.com/30825
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8a11cb318f5c4c15b67ffe9ff9b2fa772bf1dd85
Author: Adam Langley <agl@golang.org>
Date:   Mon Oct 10 18:23:37 2016 -0700

    crypto/tls: support X25519.
    
    X25519 (RFC 7748) is now commonly used for key agreement in TLS
    connections, as specified in
    https://tools.ietf.org/html/draft-ietf-tls-curve25519-01.
    
    This change adds support for that in crypto/tls, but does not enabled it
    by default so that there's less test noise. A future change will enable
    it by default and will update all the test data at the same time.
    
    Change-Id: I91802ecd776d73aae5c65bcb653d12e23c413ed4
    Reviewed-on: https://go-review.googlesource.com/30824
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8c74139cada0b7da5a0914ee0afd75837b6f682d
Author: Adam Langley <agl@golang.org>
Date:   Tue Oct 11 14:45:29 2016 -0700

    crypto/tls: fix printing of OpenSSL output when updating a test fails.
    
    When updating the test data against OpenSSL, the handshake can fail and
    the stdout/stderr output of OpenSSL is very useful in finding out why.
    
    However, printing that output has been broken for some time because its
    no longer sent to a byte.Buffer. This change fixes that.
    
    Change-Id: I6f846c7dc80f1ccee9fa1be36f0b579b3754e05f
    Reviewed-on: https://go-review.googlesource.com/30823
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e875fe42eee942c35cdecc7b4b5d4e762f47bade
Author: Adam Langley <agl@golang.org>
Date:   Tue Oct 11 14:59:55 2016 -0700

    vendor/golang_org/x/crypto/curve25519: new package
    
    This change imports the curve25519 package from x/crypto at revision
    594708b89f21ece706681be23d04a6513a22de6e.
    
    Change-Id: I379eaa71492959e404259fc1273d0057573bc243
    Reviewed-on: https://go-review.googlesource.com/30822
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9d88292c83c17f065485d0c8fa0439f2d9e5cca2
Author: Adam Langley <agl@golang.org>
Date:   Tue Oct 11 10:08:57 2016 -0700

    crypto/tls: switch to OpenSSL 1.1.0 for test data.
    
    We will need OpenSSL 1.1.0 in order to test some of the features
    expected for Go 1.8. However, 1.1.0 also disables (by default) some
    things that we still want to test, such as RC4, 3DES and SSLv3. Thus
    developers wanting to update the crypto/tls test data will need to build
    OpenSSL from source.
    
    This change updates the test data with transcripts generated by 1.1.0
    (in order to reduce future diffs) and also causes a banner to be printed
    if 1.1.0 is not used when updating.
    
    (The test for an ALPN mismatch is removed because OpenSSL now terminates
    the connection with a fatal alert if no known ALPN protocols are
    offered. There's no point testing against this because it's an OpenSSL
    behaviour.)
    
    Change-Id: I957516975e0b8c7def84184f65c81d0b68f1c551
    Reviewed-on: https://go-review.googlesource.com/30821
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0e36456bf9aa7b3041987634e478449215cdbd82
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Fri Sep 2 16:17:37 2016 -0700

    archive/tar: fix and cleanup readOldGNUSparseMap
    
    * Assert that the format is GNU.
    Both GNU and STAR have some form of sparse file support with
    incompatible header structures. Worse yet, both formats use the
    'S' type flag to indicate the presence of a sparse file.
    As such, we should check the format (based on magic numbers)
    and fail early.
    
    * Move realsize parsing logic into readOldGNUSparseMap.
    This is related to the sparse parsing logic and belongs here.
    
    * Fix the termination condition for parsing sparse fields.
    The termination condition for reading the sparse fields
    is to simply check if the first byte of the offset field is NULL.
    This does not seem to be documented in the GNU manual, but this is
    the check done by the both the GNU and BSD implementations:
            http://git.savannah.gnu.org/cgit/tar.git/tree/src/sparse.c?id=9a33077a7b7ad7d32815a21dee54eba63b38a81c#n731
            https://github.com/libarchive/libarchive/blob/1fa9c7bf90f0862036a99896b0501c381584451a/libarchive/archive_read_support_format_tar.c#L2207
    
    * Fix the parsing of sparse fields to use parseNumeric.
    This is what GNU and BSD do. The previous two links show that
    GNU and BSD both handle base-256 and base-8.
    
    * Detect truncated streams.
    The call to io.ReadFull does not check if the error is io.EOF.
    Getting io.EOF in this method is never okay and should always be
    converted to io.ErrUnexpectedEOF.
    
    * Simplify the function.
    The logic is essentially a do-while loop so we can remove
    some redundant code.
    
    Change-Id: Ib2f601b1a283eaec1e41b1d3396d649c80749c4e
    Reviewed-on: https://go-review.googlesource.com/28471
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit 14204662c8a73ea6d5b6489b0c5a6b0345b99a0d
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Fri Sep 2 21:03:57 2016 -0700

    archive/tar: handle integer overflow on 32bit machines
    
    Most calls to strconv.ParseInt(x, 10, 0) should really be
    calls to strconv.ParseInt(x, 10, 64) in order to ensure that they
    do not overflow on 32b architectures.
    
    Furthermore, we should document a bug where Uid and Gid may
    overflow on 32b machines since the type is declared as int.
    
    Change-Id: I99c0670b3c2922e4a9806822d9ad37e1a364b2b8
    Reviewed-on: https://go-review.googlesource.com/28472
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 94f49fd40dd907bcf9e16ab9a798409b8fcc13fd
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Wed Oct 12 23:05:35 2016 +0900

    unicode: panic if given rune is negative in SimpleFold
    
    Fixes #16690
    
    Change-Id: I6db588c4b0f23c5ec6bc9b85a488b60fab3f2f13
    Reviewed-on: https://go-review.googlesource.com/30892
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit be302e6d43790c3398e5b03c955f257868855a80
Author: Momchil Velikov <momchil.velikov@gmail.com>
Date:   Wed Sep 7 15:10:07 2016 +0300

    cmd/compile: trim more blocks
    
     - trim blocks with multiple predecessors
     - trim blocks, which contain only phi-functions
     - trim blocks, which can be merged into the successor block
    
    As an example, compiling the following source:
    
    ---8<------
    package p
    
    type Node struct {
            Key         int
            Left, Right *Node
    }
    
    func Search(r *Node, k int) *Node {
            for r != nil {
                    switch {
                    case k == r.Key:
                            return r
                    case k < r.Key:
                            r = r.Left
                    default:
                            r = r.Right
                    }
            }
            return nil
    }
    ---8<------
    
    with `GOSSAFUNC=Search" go tool compile t.go`, results in the following
    code:
    
    ---8<------
    genssa
    
          00000 (t.go:8)    TEXT    "".Search(SB), $0
          00001 (t.go:8)    FUNCDATA        $0, "".gcargs·0(SB)
          00002 (t.go:8)    FUNCDATA        $1, "".gclocals·1(SB)
          00003 (t.go:8)    TYPE    "".r(FP)type.*"".Node, $8
          00004 (t.go:8)    TYPE    "".k+8(FP)type.int, $8
          00005 (t.go:8)    TYPE    "".~r2+16(FP)type.*"".Node, $8
    v40   00006 (t.go:9)    MOVQ    "".k+8(FP), AX
    v34   00007 (t.go:9)    MOVQ    "".r(FP), CX
    v33   00008 (t.go:9)    TESTQ   CX, CX
    b2    00009 (t.go:9)    JEQ     $0, 22
    v16   00010 (t.go:11)   MOVQ    (CX), DX
    v21   00011 (t.go:11)   CMPQ    DX, AX
    b9    00012 (t.go:11)   JEQ     $0, 19
    v64   00013 (t.go:13)   CMPQ    AX, DX
    b13   00014 (t.go:13)   JGE     17
    v36   00015 (t.go:14)   MOVQ    8(CX), CX
    b4    00016 (t.go:9)    JMP     8                  <---+
    v42   00017 (t.go:16)   MOVQ    16(CX), CX             |
    b21   00018 (t.go:10)   JMP     16                 ----+
    v28   00019 (t.go:12)   VARDEF  "".~r2+16(FP)
    v29   00020 (t.go:12)   MOVQ    CX, "".~r2+16(FP)
    b10   00021 (t.go:12)   RET
    v44   00022 (t.go:19)   VARDEF  "".~r2+16(FP)
    v45   00023 (t.go:19)   MOVQ    $0, "".~r2+16(FP)
    b5    00024 (t.go:19)   RET
    00025 (<unknown line number>)   END
    ---8<------
    
    Note the jump at 18 jumps to another jump at 16.
    
    Looking at the function after trimming:
    
    --8<------
    after trim [199 ns]
    
    b1:
    v1 = InitMem <mem>
    v2 = SP <uintptr> : SP
    v67 = Arg <*Node> {r} : r[*Node]
    v59 = Arg <int> {k} : k[int]
    v40 = LoadReg <int> v59 : AX
    v34 = LoadReg <*Node> v67 : CX
    Plain → b2
    
    b2: ← b1 b4
    v8 = Phi <*Node> v34 v68 : CX
    v33 = TESTQ <flags> v8 v8
    NE v33 → b9 b5 (likely)
    
    b9: ← b2
    v16 = MOVQload <int> v8 v1 : DX
    v21 = CMPQ <flags> v16 v40
    EQ v21 → b10 b13 (unlikely)
    
    b13: ← b9
    v64 = CMPQ <flags> v40 v16
    LT v64 → b19 b21
    
    b19: ← b13
    v36 = MOVQload <*Node> [8] v8 v1 : CX
    Plain → b4
    
    b4: ← b21 b19                       <
    v68 = Phi <*Node> v42 v36 : CX      <- no actual code
    Plain → b2                          <
    
    b21: ← b13
    v42 = MOVQload <*Node> [16] v8 v1 : CX
    Plain → b4
    
    b10: ← b9
    v28 = VarDef <mem> {~r2} v1
    v29 = MOVQstore <mem> {~r2} v2 v8 v28
    v30 = Copy <mem> v29
    Ret v30
    
    b5: ← b2
    v44 = VarDef <mem> {~r2} v1
    v45 = MOVQstoreconst <mem> {~r2} [val=0,off=0] v2 v44
    v47 = Copy <mem> v45
    Ret v47
    
    --8<------
    
    The jump at 16 corresponds to the edge b21 -> b4. The block b4 contains
    only phi-ops, i.e. no actual code besides the jump to b2. However b4 is
    not trimmed, because it a) has more than one predecessor, and b) it is
    not empty.
    
    This change enhances trim.go to remove more blocks, subject to the
    following criteria:
    
     - block has predecessors (i.e. not the start block)
    
     - block is BlockPlain
    
     - block does not loop back to itself
    
     - block is the single predecessor of its successor; the instructions of
       the block are merged into the successor
    
     - block does no emit actual code, besides a possible unconditional
       jump.
         Currently only OpPhi are considered to not be actual code,
       perhaps OpKeepAlive/others should be considered too?
    
    As an example, after the change, the block b4 is trimmed and the jump at
    18 jumps directly to 8.
    
    Revision 1: Adjust phi-ops arguments after merge
    
    Ensure the number of phi-ops arguments matches the new number of
    predecessors in the merged block.
    When moving values, make them refer to the merged block.
    
    Revision 2:
     - Make clear the intent that we do not want to trim the entry block
     - Double check that we are merging a phi operation
     - Minor code style fix
     - Fix a potentially dangerous situation when a blocks refers to the
       inline value space in another block
    
    Change-Id: I0ab91779f931f404d11008f5c45606d985d7fbaa
    Reviewed-on: https://go-review.googlesource.com/28812
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 8d1d9292ff024f6c7586d27edd2c84c1ca8d9bf5
Author: Kevin Burke <kev@inburke.com>
Date:   Tue Sep 13 23:31:07 2016 -0700

    syscall: document that Exec wraps execve(2)
    
    Change-Id: I611511434f37c75f77c22f61f469108243bc0101
    Reviewed-on: https://go-review.googlesource.com/29121
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 68998433556a5c9ee0f7d8d544ffb006a724adaf
Author: Cyrill Schumacher <cyrill@schumacher.fm>
Date:   Sat Aug 20 11:32:32 2016 +0200

    net/http: optimize internal cookie functions
    
    - precalculate *Cookie slice in read cookie functions
    - readSetCookies: pre-allocs depending on the count of Set-Cookies
    - rename success variable to ok; avoid else
    - refactor Cookie.String to use less allocations
    - remove fmt package and replace with writes to a bytes.Buffer
    - add BenchmarkReadSetCookies and BenchmarkReadCookies
    
    name              old time/op    new time/op    delta
    CookieString-8      1.42µs ± 2%    0.78µs ± 1%  -45.36%        (p=0.000 n=10+10)
    ReadSetCookies-8    3.46µs ± 1%    3.42µs ± 2%   -1.39%        (p=0.001 n=10+10)
    ReadCookies-8       5.12µs ± 1%    5.15µs ± 2%     ~           (p=0.393 n=10+10)
    
    name              old alloc/op   new alloc/op   delta
    CookieString-8        520B ± 0%      384B ± 0%  -26.15%        (p=0.000 n=10+10)
    ReadSetCookies-8      968B ± 0%      960B ± 0%   -0.83%        (p=0.000 n=10+10)
    ReadCookies-8       2.01kB ± 0%    2.01kB ± 0%     ~     (all samples are equal)
    
    name              old allocs/op  new allocs/op  delta
    CookieString-8        10.0 ± 0%       3.0 ± 0%  -70.00%        (p=0.000 n=10+10)
    ReadSetCookies-8      18.0 ± 0%      17.0 ± 0%   -5.56%        (p=0.000 n=10+10)
    ReadCookies-8         16.0 ± 0%      16.0 ± 0%     ~     (all samples are equal)
    
    Change-Id: I870670987f10f3e52f9c657cfb8e6eaaa97a6162
    Reviewed-on: https://go-review.googlesource.com/27850
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 9c850958cea35d019142f8341beacb4151e1511b
Author: Alan Donovan <adonovan@google.com>
Date:   Fri Oct 7 18:11:06 2016 -0400

    go/types: expose Default function, which converts untyped T to T
    
    Change-Id: Ibcf5e0ba694b280744a00c2c6fda300f0a653455
    Reviewed-on: https://go-review.googlesource.com/30715
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 41a005d458558b41b4bc7a4c837953b35609c9a2
Author: Alan Donovan <adonovan@google.com>
Date:   Fri Sep 9 11:14:41 2016 -0400

    test: add test for issue 17039
    
    Change-Id: Ieb3d605f03a7185a707621bef7160090c9bdb51f
    Reviewed-on: https://go-review.googlesource.com/28873
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 898ca6ba0af0ea1180fea1f226ff6ef731018ac2
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Sep 11 19:05:29 2016 -0700

    runtime: update mkduff legacy comments
    
    Update comments for duffzero and duffcopy
    which referred to legacy locations:
    + cmd/?g/cgen.go
    + cmd/?g/ggen.go
    
    Remnants of the old days when we had 5g, 6g etc.
    
    Those locations have since moved to:
    + cmd/compile/internal/<arch>/cgen.go
    + cmd/compile/internal/<arch>/ggen.go
    
    Change-Id: Ie2ea668559d52d42b747260ea69a6d5b3d70e859
    Reviewed-on: https://go-review.googlesource.com/29073
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 6d702d8ed2a980cd0c96499379eee5eb218f8339
Author: Jean-Nicolas Moal <jn.moal@gmail.com>
Date:   Mon Aug 22 19:02:33 2016 +0200

    os: add examples of environment functions
    
    For #16360.
    
    Change-Id: Iaa3548704786018eacec530f7a907b976fa532fe
    Reviewed-on: https://go-review.googlesource.com/27443
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b65cdc28882bfd7c4be46e811a6a7841d9fb7d53
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Thu Aug 18 18:40:02 2016 +0900

    path/filepath: add a test case for EvalSymlinks error
    
    EvalSymlinks returns error if given path or its target path don't exist.
    Add a test for future improvement.
    
    Change-Id: Ic9a4aa5eaee0fe7ac523d54d8eb3132a11b380b3
    Reviewed-on: https://go-review.googlesource.com/27330
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit 6e759ad2e2ae53f08db4470aa60e972eb908f2ef
Author: Dave Cheney <dave@cheney.net>
Date:   Mon Aug 1 20:34:12 2016 +1000

    cmd/compile/internal/gc: add runtime/trace support
    
    This change adds runtime/trace support to go tool compile.
    
    Change-Id: I6c496b9b063796123f75eba6af511c53a57c0196
    Reviewed-on: https://go-review.googlesource.com/25354
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 1a3b739b26e2742fd2febc7c0f87aa8115f390ab
Author: Joshua Boelter <joshua.boelter@intel.com>
Date:   Tue Aug 9 22:37:26 2016 -0700

    runtime: check for errors returned by windows sema calls
    
    Add checks for failure of CreateEvent, SetEvent or
    WaitForSingleObject. Any failures are considered fatal and
    will throw() after printing an informative message.
    
    Updates #16646
    
    Change-Id: I3bacf9001d2abfa8667cc3aff163ff2de1c99915
    Reviewed-on: https://go-review.googlesource.com/26655
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9d36ab222d838f0e493653e2f31d77adb15ab9b5
Author: David du Colombier <0intro@gmail.com>
Date:   Wed Oct 12 14:57:19 2016 +0200

    cmd/link: use HEADR to define FlagTextAddr (cosmetic change)
    
    This cosmetic change defines ld.FlagTextAddr using ld.HEADR in
    the Plan 9 cases, like it is done for other operating systems.
    
    Change-Id: Ic929c1c437f25661058682cf3e159f0b16cdc538
    Reviewed-on: https://go-review.googlesource.com/30912
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8fc9c504968be0bbea4c8317998813782d077cf4
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Oct 12 19:41:29 2016 +0900

    net: update doc for unimplemented feature on Plan 9
    
    Also removes unnecessary allocation.
    
    Change-Id: I3406cf75a7b64d93b2b99c7f1f5c78f580452b60
    Reviewed-on: https://go-review.googlesource.com/30891
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 29c600bba1279d63a40bc99a7e6212d3f0bffa06
Author: David du Colombier <0intro@gmail.com>
Date:   Wed Oct 12 14:24:39 2016 +0200

    cmd/link: fix build on plan9/amd64
    
    Support for multiple text sections was added in CL 27790.
    However, this change broke the build on plan9/amd64.
    
    In relocsym, the R_ADDROFF relocation was changed to
    use offsets relative to the start of the first text
    section. However, Segtext.Vaddr is the address of
    the text segment, while we expect to start from
    the first section (text.runtime) of the text segment.
    
    Fixes #17411.
    
    Change-Id: I86bbcbda81cea735b0ecf156eab2e6e5d63acce3
    Reviewed-on: https://go-review.googlesource.com/30911
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6c517df4daa0acaa25a16baeef5ea037c9a0194c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Oct 11 21:04:16 2016 -0700

    syscall: unify NsecToTime{spec,val}, fix for times < 1970
    
    All the implementations of NsecToTimespec and NsecToTimeval were the
    same other than types. Write a single version that uses
    GOARCH/GOOS-specific setTimespec and setTimeval functions to handle the
    types.
    
    The logic in NsecToTimespec and NsecToTimeval caused times before 1970
    to have a negative usec/nsec. The Linux kernel requires that usec
    contain a positive number; for consistency, we do this for both
    NsecToTimespec and NsecToTimeval.
    
    Change-Id: I525eaba2e7cdb00cb57fa00182dabf19fec298ae
    Reviewed-on: https://go-review.googlesource.com/30826
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 6ca48f710f6ae163aa87e883a0a4cf8a91dad0a4
Author: Chris McGee <sirnewton_01@yahoo.ca>
Date:   Wed Sep 28 21:29:08 2016 -0400

    net: implement network interface API for Plan 9
    
    The new implementation parses the plan9 interface files
    and builds Interface representations for the net package.
    
    Updates #17218
    
    Change-Id: I3199d18d9e96a17e922186c3abff1d7cd9cbec2e
    Reviewed-on: https://go-review.googlesource.com/29963
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2287296dbe43c5dfa9afebf0e03af42318a6a8f0
Author: Sina Siadat <siadat@gmail.com>
Date:   Mon Aug 8 02:04:52 2016 +0430

    os: add example for IsNotExist
    
    Show usage of os.IsNotExist in an example.
    
    Change-Id: I5306ea06c370099de5b02668dfa02b87b0c2beac
    Reviewed-on: https://go-review.googlesource.com/25571
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f373bf1eb9a21f843dc0094c3702c2855a2e0d3f
Author: Shawn Walker-Salas <shawn.walker@oracle.com>
Date:   Wed Jun 15 13:44:03 2016 -0700

    cmd/link: non-executable stack support for Solaris
    
    Support the tagging of binaries created with the internal linker
    on Solaris as having a non-executable stack by writing a PT_SUNWSTACK
    program header.
    
    Fixes #16074
    
    Change-Id: I3888f2153083385d04a52f341570f93e5738b276
    Reviewed-on: https://go-review.googlesource.com/24142
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 456b7f5a974f229e0ef7ad5a1925fa72fc4182e2
Author: Gyu-Ho Lee <gyuhox@gmail.com>
Date:   Sat Jun 4 23:20:45 2016 -0700

    runtime/pprof: preallocate slice in pprof.go
    
    To prevent slice growth when appending.
    
    Change-Id: I2cdb9b09bc33f63188b19573c8b9a77601e63801
    Reviewed-on: https://go-review.googlesource.com/23783
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 81b9af7cccce8234319551330cb6406469f32bab
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu Jul 28 10:33:26 2016 +1000

    os: add new tests for symbolic links and directory junctions
    
    Updates #15978
    Updates #16145
    
    Change-Id: I161f5bc97d41c08bf5e1405ccafa86232d70886d
    Reviewed-on: https://go-review.googlesource.com/25320
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit d2ef288c8be22ccc39f33251bb33530105b17cfa
Author: Tristan Ooohry <ooohry@gmail.com>
Date:   Fri Jul 22 04:11:58 2016 +0000

    cmd/go: added verbose error from matchGoImport
    
    The error coming out of matchGoImport does not differentiate between
    having no imports, and having some invalid imports.
    
    This some extra context to the error message to help debug these issues.
    
    Fixes #16467
    
    Change-Id: I3e9a119ed73da1bed5e07365be0221ea6b7f19db
    Reviewed-on: https://go-review.googlesource.com/25121
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 87b1aaa37cefec8deacdf9c3c30d26015bdfb00b
Author: Xuyang Kang <xuyangkang@gmail.com>
Date:   Sun Jul 17 00:23:56 2016 -0700

    encoding/base64: This change modifies Go to take strict option when decoding base64
    
    If strict option is enabled, when decoding, instead of skip the padding
    bits, it will do strict check to enforce they are set to zero.
    
    Fixes #15656
    
    Change-Id: I869fb725a39cc9dde44dbc4ff0046446e7abc642
    Reviewed-on: https://go-review.googlesource.com/24964
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 56b5546b9138f80bb0b20aebcc1fa551096e87df
Author: Gustav Westling <gustav@westling.xyz>
Date:   Fri Aug 5 16:13:23 2016 +0200

    cmd/gofmt: simplify map key literals
    
    Simplify map key literals in "gofmt -s"
    
    Fixes #16461.
    
    Change-Id: Ia61739b34a30ac27f6696f94a98809109a8a7b61
    Reviewed-on: https://go-review.googlesource.com/25530
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit d07345a6334def1a8273107f8bffab33f7a701bc
Author: Keith Randall <khr@golang.org>
Date:   Tue Oct 11 13:50:44 2016 -0700

    cmd/compile: update ssa html help text
    
    Update the description of the conditions under which highlighting might
    be misleading.
    
    Fixes #16754
    
    Change-Id: I3078a09e0b9a76d12078352e15a3f26ba3f1bbee
    Reviewed-on: https://go-review.googlesource.com/30818
    Reviewed-by: David Chase <drchase@google.com>

commit 6597bcbe53cad07b0ac4070b4f428f5db3331383
Author: Dmitri Shuralyov <shurcooL@gmail.com>
Date:   Tue Jul 26 13:01:18 2016 -0400

    path/filepath: remove unneeded doc statement for SplitList
    
    This is a followup to CL 24747, where the package doc phrase
    "Functions in this package replace occurrences of slash unless otherwise specified."
    was removed. The phrase was originally added in CL 7310 together
    with this explicit opt out statement for SplitList.
    
    Remove it since it's no longer neccessary. This helps consistency.
    
    Updates #16111.
    Updates #10122.
    
    Change-Id: Iba86de57c24100adecac9cb5892ce180126c0ea6
    Reviewed-on: https://go-review.googlesource.com/25250
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 460d112f6c3716a45e99651ed769e7d8d74aca41
Author: Gleb Stepanov <glebstepanov1992@gmail.com>
Date:   Mon Jul 25 15:53:15 2016 +0300

    runtime: fix typo in comments
    
    Fix typo in word synchronization in comments.
    
    Change-Id: I453b4e799301e758799c93df1e32f5244ca2fb84
    Reviewed-on: https://go-review.googlesource.com/25174
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 8728df645c70e6420eb59e5886bc839022998322
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Apr 24 17:36:41 2016 -0400

    runtime: remove canBackTrace variable from TestGdbPython
    
    The canBackTrace variable is true for all of the architectures
    Go supports and this is likely to remain the case as new
    architectures are added.
    
    Change-Id: I73900c018eb4b2e5c02fccd8d3e89853b2ba9d90
    Reviewed-on: https://go-review.googlesource.com/22423
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0aadaf2ad5fa8cc1c00256840e37f422dda9689f
Author: Michael Pratt <mpratt@google.com>
Date:   Wed Jul 20 22:01:33 2016 -0700

    cmd/pprof: instruction-level granularity in callgrind output
    
    When generating callgrind format output, produce cost lines at
    instruction granularity. This allows visualizers supporting the
    callgrind format to display instruction-level profiling information.
    
    We also need to provide the object file (ob=) in order for tools to find
    the object file to disassemble when displaying assembly.
    
    We opportunistically group cost lines corressponding to the same
    function together, reducing the number of superfluous description lines.
    Subposition compression (relative position numbering) is also used to
    reduce the output size.
    
    Change-Id: Id8e960b81dc7a47ec1dfbae877521f76972431c4
    Reviewed-on: https://go-review.googlesource.com/23781
    Run-TryBot: Michael Pratt <mpratt@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Raul Silvera <rsilvera@google.com>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit f64c67018103c200c9e150ac5f1199a3175f54a6
Author: Filippo Valsorda <hi@filippo.io>
Date:   Tue May 24 12:50:38 2016 +0100

    cmd/trace: add option to output pprof files
    
    The trace tool can generate some interesting profiles, but it was only
    exposing them as svg through the web UI.  This adds command line options
    to generate the raw pprof file.
    
    Change-Id: I52e4f909fdca6f65c3616add444e3892783640f4
    Reviewed-on: https://go-review.googlesource.com/23324
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 7f6eadb64fb16e73b15630dd3089edc568d047cb
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Oct 11 14:20:50 2016 -0700

    syscall: unify TimespecToNsec and TimevalToNsec
    
    All implementations of these functions are identical.
    
    Change-Id: I7cbea53c02bb0cee75e30beed19d29ba0a7ef657
    Reviewed-on: https://go-review.googlesource.com/30819
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e1fc292500aa70c265937aebad00ac010031cbaf
Author: James Clarke <jrtc27@jrtc27.com>
Date:   Tue Oct 11 19:08:45 2016 +0100

    debug/elf: add sparc64 relocations
    
    Change-Id: I1a2504ad9ca8607588d2d366598115fe360435b5
    Reviewed-on: https://go-review.googlesource.com/30870
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 35220534d5edbbbcd0eed59133bcfae54d140287
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Mon Oct 10 18:23:56 2016 -0700

    archive/zip: only use Extended Timestamp on non-zero MS-DOS timestamps
    
    We should preserve the fact that a roundtrip read on fields with the zero
    value should remain the zero for those that are reasonable to stay that way.
    If the zero value for a MS-DOS timestamp was used, then it is sensible for
    that zero value to also be read back later.
    
    Fixes #17403
    
    Change-Id: I32c3915eab180e91ddd2499007374f7b85f0bd76
    Reviewed-on: https://go-review.googlesource.com/30811
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 303b69feb7b26b583b53a3a82d824088064bbf2b
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 11 11:34:20 2016 -0700

    cmd/compile, runtime: stop padding stackmaps to 4 bytes
    
    Shrinks cmd/go's text segment by 0.9%.
    
       text    data     bss     dec     hex filename
    6447148  231643  146328 6825119  68249f go.before
    6387404  231643  146328 6765375  673b3f go.after
    
    Change-Id: I431e8482dbb11a7c1c77f2196cada43d5dad2981
    Reviewed-on: https://go-review.googlesource.com/30817
    Reviewed-by: Austin Clements <austin@google.com>

commit 943f5afe22c1a07da8954756a3701ac765bbf3c1
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Oct 11 10:23:14 2016 -0700

    cmd/compile: refactor stackmap dumping code
    
    Also, fix a byte-ordering problem with stack maps for assembly
    function signatures on big-endian targets.
    
    Change-Id: I6e8698f5fbb04b31771a65f4a8f3f9c045ff3c98
    Reviewed-on: https://go-review.googlesource.com/30816
    Reviewed-by: Austin Clements <austin@google.com>

commit 15817e409bce6eeaa7ea2ae158db6ce4618f27f2
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Oct 6 15:06:45 2016 -0400

    cmd/compile: make link register allocatable in non-leaf functions
    
    We save and restore the link register in non-leaf functions because
    it is clobbered by CALLs. It is therefore available for general
    purpose use.
    
    Only enabled on s390x currently. The RC4 benchmarks in particular
    benefit from the extra register:
    
    name     old speed     new speed     delta
    RC4_128  243MB/s ± 2%  341MB/s ± 2%  +40.46%  (p=0.008 n=5+5)
    RC4_1K   267MB/s ± 0%  359MB/s ± 1%  +34.32%  (p=0.008 n=5+5)
    RC4_8K   271MB/s ± 0%  362MB/s ± 0%  +33.61%  (p=0.008 n=5+5)
    
    Change-Id: Id23bff95e771da9425353da2f32668b8e34ba09f
    Reviewed-on: https://go-review.googlesource.com/30597
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 809a1de1ac1ccc76f7a4faf630017626f2f68231
Author: Adam Langley <agl@golang.org>
Date:   Mon Oct 10 16:26:51 2016 -0700

    crypto/x509: parse all names in an RDN.
    
    The Subject and Issuer names in a certificate look like they should be a
    list of key-value pairs. However, they're actually a list of lists of
    key-value pairs. Previously we only looked at the first element of each
    sublist and the vast majority of certificates only have one element per
    sublist.
    
    However, it's possible to have multiple elements and some 360
    certificates from the “Pilot” log are so constructed.
    
    This change causes all elements of the sublists to be processed.
    
    Fixes #16836.
    
    Change-Id: Ie0a5159135b08226ec517fcf251aa17aada37857
    Reviewed-on: https://go-review.googlesource.com/30810
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c5368123fedba827967628943bf45ed3d1d369ac
Author: Hajime Hoshi <hajimehoshi@gmail.com>
Date:   Sun Oct 9 01:09:52 2016 +0900

    cmd/compile: remove redundant function idom
    
    Change-Id: Ib14b5421bb5e407bbd4d3cbfc68c92d3dd257cb1
    Reviewed-on: https://go-review.googlesource.com/30732
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f1eed92fd04cc93db44655a5a1870d361c9c5137
Author: David Chase <drchase@google.com>
Date:   Sat Oct 8 16:45:58 2016 -0400

    cmd/compile: escape analysis needs to run "flood" to fixed point
    
    In some cases the members of the root set from which flood
    runs themselves escape, without their referents being also
    tagged as escaping.  Fix this by reflooding from those roots
    whose escape increases, and also enhance the "leak" test to
    include reachability from a heap-escaped root.
    
    Fixes #17318.
    
    Change-Id: Ied1e75cee17ede8ca72a8b9302ce8201641ec593
    Reviewed-on: https://go-review.googlesource.com/30693
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 88562dc83ecc3e0c5ce85f22356cb7114e4756df
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 10 16:03:13 2016 -0400

    math/big: move ProbablyPrime into its own source file
    
    A later CL will be adding more code here.
    It will help to keep it separate from the other code.
    
    Change-Id: I971ba53de819cd10991b51fdec665984939a5f9b
    Reviewed-on: https://go-review.googlesource.com/30709
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9927f25d711ec7aa0876e33e3bd174e09cc032bd
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 10 16:45:30 2016 -0400

    math/big: test and optimize Exp(2, y, n) for large y, odd n
    
    The Montgomery multiply code is applicable to this case
    but was being bypassed. Don't do that.
    
    The old test len(x) > 1 was really just a bad approximation to x > 1.
    
    name    old time/op  new time/op  delta
    Exp-8   5.56ms ± 4%  5.73ms ± 3%     ~     (p=0.095 n=5+5)
    Exp2-8  7.59ms ± 1%  5.66ms ± 1%  -25.40%  (p=0.008 n=5+5)
    
    This comes up especially when doing Fermat (Miller-Rabin)
    primality tests with base 2.
    
    Change-Id: I4cc02978db6dfa93f7f3c8f32718e25eedb4f5ed
    Reviewed-on: https://go-review.googlesource.com/30708
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9a8832f1422f7fa72e4855757e4a951957cc62ae
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 10 16:44:53 2016 -0400

    math/big: move exhaustive tests behind -long flag
    
    This way you can still run 'go test' or 'go bench -run Foo'
    without wondering why it is taking so very long.
    
    Change-Id: Icfa097a6deb1d6682acb7be9f34729215c29eabb
    Reviewed-on: https://go-review.googlesource.com/30707
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 2756d56c894e5b044907da74b6d5f3c684eab00d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Oct 6 15:43:47 2016 -0400

    cmd/compile: intrinsify math/big.mulWW, divWW on AMD64
    
    Change-Id: I59f7afa7a5803d19f8b21fe70fc85ef997bb3a85
    Reviewed-on: https://go-review.googlesource.com/30542
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 7c431cb7f9780fcaf58b9ef07028d5129e1e5fe7
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Sep 14 14:47:12 2016 -0400

    cmd/link: insert trampolines for too-far jumps on ARM
    
    ARM direct CALL/JMP instruction has 24 bit offset, which can only
    encodes jumps within +/-32M. When the target is too far, the top
    bits get truncated and the program jumps wild.
    
    This CL detects too-far jumps and automatically insert trampolines,
    currently only internal linking on ARM.
    
    It is necessary to make the following changes to the linker:
    - Resolve direct jump relocs when assigning addresses to functions.
      this allows trampoline insertion without moving all code that
      already laid down.
    - Lay down packages in dependency order, so that when resolving a
      inter-package direct jump reloc, the target address is already
      known. Intra-package jumps are assumed never too far.
    - a linker flag -debugtramp is added for debugging trampolines:
        "-debugtramp=1 -v" prints trampoline debug message
        "-debugtramp=2"    forces all inter-package jump to use
                           trampolines (currently ARM only)
        "-debugtramp=2 -v" does both
    - Some data structures are changed for bookkeeping.
    
    On ARM, pseudo DIV/DIVU/MOD/MODU instructions now clobber R8
    (unfortunate). In the standard library there is no ARM assembly
    code that uses these instructions, and the compiler no longer emits
    them (CL 29390).
    
    all.bash passes with -debugtramp=2, except a disassembly test (this
    is unavoidable as we changed the instruction).
    
    TBD: debug info of trampolines?
    
    Fixes #17028.
    
    Change-Id: Idcce347ea7e0af77c4079041a160b2f6e114b474
    Reviewed-on: https://go-review.googlesource.com/29397
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d03e8b226cd11692ca9505a815af559ce7989700
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Oct 4 07:11:55 2016 -0700

    runtime: record current PC for SIGPROF on non-Go thread
    
    If we get a SIGPROF on a non-Go thread, and the program has not called
    runtime.SetCgoTraceback so we have no way to collect a stack trace, then
    record a profile that is just the PC where the signal occurred. That
    will at least point the user to the right area.
    
    Retrieving the PC from the sigctxt in a signal handler on a non-G thread
    required marking a number of trivial sigctxt methods as nosplit, and,
    for extra safety, nowritebarrierrec.
    
    The test shows that the existing test CgoPprofThread test does not test
    the stack trace, just the profile signal. Leaving that for later.
    
    Change-Id: I8f8f3ff09ac099fc9d9df94b5a9d210ffc20c4ab
    Reviewed-on: https://go-review.googlesource.com/30252
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit c24cc40075d77b96bbf1f217dcdcff651229e89b
Author: Hyang-Ah (Hana) Kim <hyangah@gmail.com>
Date:   Mon Aug 8 17:24:07 2016 -0400

    cmd/trace: fix a runnable goroutine count bug
    
    When starting tracing, EvGoCreate events are added for existing
    goroutines that may have been blocking in syscall. EvGoCreate
    increments the runnable goroutine count. This change makes the
    following EvGoInSyscall event decrement the runnable goroutine count
    because we now know that goroutine is in syscall, and not runnable.
    
    Made generateTrace return an error, at any given time, the number
    of runnable/running/insyscall goroutines becomes non-negative.
    
    Added a basic test that checks the number of runnable/running
    goroutines don't include the goroutines in syscall - the test failed
    before this change.
    
    Change-Id: Ib732c382e7bd17158a437576f9d589ab89097ce6
    Reviewed-on: https://go-review.googlesource.com/25552
    Run-TryBot: Hyang-Ah Hana Kim <hyangah@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit dd307da10c3582c195928d9bf073d1b0b01f2135
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Tue Oct 11 14:24:59 2016 +1100

    runtime/cgo: do not explicitly link msvcrt.dll
    
    CL 14472 solved issue #12030 by explicitly linking msvcrt.dll
    to every cgo executable we build. This CL achieves the same
    by manually loading ntdll.dll during startup.
    
    Updates #12030
    
    Change-Id: I5d9cd925ef65cc34c5d4031c750f0f97794529b2
    Reviewed-on: https://go-review.googlesource.com/30737
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 8f65379f947f77f5caadf283f219b2cc6e9bdb84
Author: John Dethridge <jcd@golang.org>
Date:   Fri Oct 7 15:06:09 2016 +1100

    cmd/link: more efficient encoding of DWARF line number information
    
    The (pc, line) deltas in the line number information are currently encoded
    either with a special opcode, or with a triplet of DW_LNS_advance_pc,
    DW_LNS_advance_line, and DW_LNS_copy instructions.  Instead of DW_LNS_copy,
    this change always uses a special opcode, which can make DW_LNS_advance_pc or
    DW_LNS_advance_line unnecessary, or make their operands take fewer bytes.  It
    chooses the special opcode so that the encoding of the remaining deltas is as
    small as possible.
    
    Use DW_LNS_const_add_pc or DW_LNS_fixed_advance_pc instead of DW_LNS_advance_pc
    for deltas where they save a byte.
    
    Update LINE_BASE and LINE_RANGE constants to optimal values for this strategy.
    
    This reduces line number information by about 35% and total size by about 2%
    for a typical binary.
    
    Change-Id: Ia61d6bf19c95c1d34ba63c67ed32b376beda225f
    Reviewed-on: https://go-review.googlesource.com/30577
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3a9072829e7d315ecda030281b622632c1bbe1b6
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 6 22:42:20 2016 -0400

    math/big: make division faster
    
    - Add new BenchmarkQuoRem.
    - Eliminate allocation in divLarge nat pool
    - Unroll mulAddVWW body 4x
    - Remove some redundant slice loads in divLarge
    
    name      old time/op  new time/op  delta
    QuoRem-8  2.18µs ± 1%  1.93µs ± 1%  -11.38%  (p=0.000 n=19+18)
    
    The starting point in the comparison here is Cherry's
    pending CL to turn mulWW and divWW into intrinsics.
    The optimizations in divLarge work best because all
    the function calls are gone. The effect of this CL is not
    as large if you don't assume Cherry's CL.
    
    Change-Id: Ia6138907489c5b9168497912e43705634e163b35
    Reviewed-on: https://go-review.googlesource.com/30613
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 68331750dac5a38c5158f57ab19e3e99d11a59e3
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Oct 3 23:01:26 2016 -0400

    cmd/compile: remove some write barriers for stack writes
    
    This, along with CL 30140, removes ~50% of stack write barriers
    mentioned in issue #17330. The remaining are most due to Phi and
    FwdRef, which is not resolved when building SSA. We might be
    able to do it at a later stage where Phi and Copy propagations
    are done, but matching an if-(store-store-call)+ sequence seems
    not very pleasant.
    
    Updates #17330.
    
    Change-Id: Iaa36c7b1f4c4fc3dc10a27018a3b0e261094cb21
    Reviewed-on: https://go-review.googlesource.com/30290
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 2f0b8f88df33ca48494a90eea705ad092322a1ca
Author: David Chase <drchase@google.com>
Date:   Wed Oct 5 13:21:09 2016 -0700

    cmd/compile: PPC64, elide unnecessary sign extension
    
    Inputs to store[BHW] and cmpW(U) need not be correct
    in more bits than are used by the instruction.
    
    Added a pattern tailored to what appears to be cgo boilerplate.
    Added a pattern (also seen in cgo boilerplate and hashing)
    to replace {EQ,NE}-CMP-ANDconst with {EQ-NE}-ANDCCconst.
    Added a pattern to clean up ANDconst shift distance inputs
    (this was seen in hashing).
    
    Simplify repeated and,or,xor.
    
    Fixes #17109.
    
    Change-Id: I68eac83e3e614d69ffe473a08953048c8b066d88
    Reviewed-on: https://go-review.googlesource.com/30455
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 672e57944458fc3c5f5ee1dd11d4f32d1aeaebe1
Author: Tal Shprecher <tshprecher@gmail.com>
Date:   Wed Jul 13 12:29:39 2016 -0600

    cmd/compile: avoid leak of dottype expression on double assignment form
    
    This is a followup to issue #13805. That change avoid leaks for types that
    don't have any pointers for the single assignment form of a dottype expression.
    This does the same for the double assignment form.
    
    Fixes #15796
    
    Change-Id: I27474cade0ff1f3025cb6392f47b87b33542bc0f
    Reviewed-on: https://go-review.googlesource.com/24906
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 795289b114e4dd4f69eaa7f74cb9eb7b6f963fff
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Oct 9 21:03:00 2016 -0700

    net/http: fix typo in server commont
    
    Change-Id: I5b04ba7e12eff933fc67eb7a1cbdfde536e4db88
    Reviewed-on: https://go-review.googlesource.com/30722
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 334cbe3bc81e2f3a6c968877b1ae170547989f1e
Author: Anmol Sethi <anmol@aubble.com>
Date:   Sat Oct 8 16:07:40 2016 -0400

    io: simplified a small part of copyBuffer
    
    Change-Id: I0b7052103174f0864ee9714f76f8f78f2a988777
    Reviewed-on: https://go-review.googlesource.com/30719
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5f1a013ea7e755e963518cc6c17a17a270041864
Author: Quentin Renard <contact@asticode.com>
Date:   Thu Oct 6 22:47:53 2016 +0200

    net/http: Add missing tests for parsePostForm
    
    Renamed TestPOSTQuery to TestParseFormQuery and added testing
    for the ";" delimiter, an empty key, an empty value and an
    empty key + value.
    
    Also added TestParseFormQueryMethods to make sure forms sent in
    PATCH and PUT (and no others) request  are parsed correctly in
    ParseForm.
    
    Fixes #17368
    
    Change-Id: I445aad324ffc7b38d179ea41953bffbac0cddffe
    Reviewed-on: https://go-review.googlesource.com/30555
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f13372c9f7df12bfa4fbba5d9802f17fd186a0bc
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Oct 9 11:50:21 2016 -0400

    cmd/internal/obj/s390x: remove support for stores of global addresses
    
    This CL removes support for MOVD instructions that store the address
    of a global variable. For example:
    
      MOVD $main·a(SB), (R1)
      MOVD $main·b(SB), main·c(SB)
    
    These instructions are emulated and the new backend doesn't need them
    (the stores now always go through an intermediate register).
    
    Change-Id: I3a1bcb3f19c5096ad0426afd76d35a4d7975733b
    Reviewed-on: https://go-review.googlesource.com/30720
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0ff40a76ad81e2b02c24e83eee8aa93352498f6f
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Mon May 30 00:41:03 2016 -0700

    crypto/x509: check that the issuer name matches the issuer's subject name.
    
    Fixes #14955.
    
    Change-Id: I157432584bb51088bec565f6bb9e64348345cff9
    Reviewed-on: https://go-review.googlesource.com/23571
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit 13c829e5f6c541359c7bb213774ef9bbba9ddb77
Author: Wedson Almeida Filho <wedsonaf@google.com>
Date:   Tue May 3 00:28:23 2016 +0100

    cmd/internal/obj/x86: On amd64, relocation type for and indirect call is pc-relative.
    
    With this change, the code in bug #15609 compiles and runs properly:
    
    0000000000401070 <main.jump>:
      401070:       ff 15 aa 7e 06 00       callq  *0x67eaa(%rip)        # 468f20 <main.pointer>
      401076:       c3                      retq
    
    0000000000468f20 g     O .rodata        0000000000000008 main.pointer
    
    Fixes #15609
    
    Change-Id: Iebb4d5a9f9fff335b693f4efcc97882fe04eefd7
    Reviewed-on: https://go-review.googlesource.com/22950
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit d5a8b9f57157e64b894ef9a1f48418f45ff7984c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Oct 9 10:20:39 2016 -0700

    doc: update go1.8.txt using new tool
    
    With help of a new interactive commit classifier tool (tool location
    TBD, likely x/build/cmd/writenotes), classify all commits from go1.7
    up to 56d35d4.
    
    We can selectively cull this list later. When in doubt, I erred on the
    side of inclusion for now.
    
    Change-Id: I458945004e1b1a148fb2f294b454a390ef4f92c2
    Reviewed-on: https://go-review.googlesource.com/30696
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 56d35d42e188b8767665a368384e084da50ef634
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Fri Oct 7 21:27:32 2016 -0700

    compress/gzip: document Reader.Read
    
    Fixes #17374.
    
    Change-Id: Ic89c35aaa31f35a8a4e3ffa09f49b68f08127625
    Reviewed-on: https://go-review.googlesource.com/30718
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 9b6ced931eb5cbde677877adcf59b0a4b8422300
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Oct 8 07:24:45 2016 +0900

    vendor: update vendored lif
    
    Updates golang_org/x/net/lif to rev 084869a for:
    - lif: rename internal types and constants generated by cgo
    
    Change-Id: Icf478d60f5ef35800966c62dcf046f7fe50204ff
    Reviewed-on: https://go-review.googlesource.com/30731
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b108009d00e55c34486541202dc11ab5cb06b63f
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Oct 8 07:22:47 2016 +0900

    vendor: update vendored route
    
    Updates golang_org/x/net/route to rev f09c466 for:
    - route: fix typo
    - route: test helper code cleanup
    
    Change-Id: If39f0e947dc56f3b0f38190035d2f47c8d847c74
    Reviewed-on: https://go-review.googlesource.com/30730
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 45b26a93f31071deee38b6579da34c2ebe98b978
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Oct 7 14:29:55 2016 -0400

    cmd/{asm,compile}: replace TESTB op with CMPWconst on s390x
    
    TESTB was implemented as AND $0xff, Rx, REGTMP. Unfortunately there
    is no 3-operand AND-with-immediate instruction and so it was emulated
    by the assembler using two instructions.
    
    This CL uses CMPW instead of AND and also optimizes CMPW to use
    the chi instruction where possible.
    
    Overall this CL reduces the size of the .text section of the
    bin/go binary by ~2%.
    
    Change-Id: Ic335c29fc1129378fcbb1265bfb10f5b744a0f3f
    Reviewed-on: https://go-review.googlesource.com/30690
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f4e37c8ec5f54964221fc950c2f0260140f438d8
Author: Keith Randall <khr@golang.org>
Date:   Fri Oct 7 10:08:10 2016 -0700

    cmd/compile: use standard dom tree in nilcheckelim
    
    No need to build a bespoke dom tree here when we might
    have one cached already.  The allocations for the dom tree
    were also more expensive than they needed to be.
    
    Fixes #12021
    
    Change-Id: I6a967880aee03660ad6fc293f8fc783779cae11d
    Reviewed-on: https://go-review.googlesource.com/30671
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 433be563b6246eb132aed6e9e58f46a0d46f7010
Author: Andrew Pogrebnoy <absourd.noise@gmail.com>
Date:   Tue Oct 4 02:39:33 2016 +0300

    cmd/compile: fix choice of phi building algorithm
    
    The algorithm for placing a phi nodes in small functions now
    unreachable. This patch fix that.
    
    Change-Id: I253d745b414fa12ee0719459c28e78a69c6861ae
    Reviewed-on: https://go-review.googlesource.com/30106
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 94589054d3ed88c049fbe468d58b995e0e154f1d
Author: Austin Clements <austin@google.com>
Date:   Thu Sep 29 11:59:56 2016 -0400

    cmd/trace: label mark termination spans as such
    
    Currently these are labeled "MARK", which was accurate in the STW
    collector, but these really indicate mark termination now, since
    marking happens for the full duration of the concurrent GC. Re-label
    them as "MARK TERMINATION" to clarify this.
    
    Change-Id: Ie98bd961195acde49598b4fa3f9e7d90d757c0a6
    Reviewed-on: https://go-review.googlesource.com/30018
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit eed309f5fb223a879630e9dced8c44c2b049cf71
Author: Austin Clements <austin@google.com>
Date:   Thu Sep 29 11:53:42 2016 -0400

    cmd/trace: move process-wide GC events to their own row
    
    Currently, the process-wide GC state is attributed to the P that
    happened to perform the allocation that exceeded the GC trigger. This
    is pretty arbitrary and makes it hard to see when GC is running since
    the GC spans are intermingled with a lot of other trace noise.
    
    The current display is particularly confusing because it creates three
    sub-rows in the P row that can overlap each other. Usually a P has
    just two sub-rows: one showing the current G and another showing that
    G's activity. However, because GC is attributed to a proc, it winds up
    as a third row that neither subsumes nor is subsumed by any other row.
    This in turn screws up the trace's layout and results in overlapping
    events.
    
    Fix these problems by creating a new dedicated row like the existing
    "Network" and "Timer" rows and displaying process-wide GC events in
    this row. Mark termination and sweep events still appear in their
    respective P rows because these are meaningfully attributed.
    
    Change-Id: Ie1a1c6cf8c446e4b043f10f3968f91ff1b546d15
    Reviewed-on: https://go-review.googlesource.com/30017
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit fa9b57bb1d84bf81b49346341ad14297d0195187
Author: Austin Clements <austin@google.com>
Date:   Thu Sep 29 11:46:53 2016 -0400

    runtime: make next_gc ^0 when GC is disabled
    
    When GC is disabled, we set gcpercent to -1. However, we still use
    gcpercent to compute several values, such as next_gc and gc_trigger.
    These calculations are meaningless when gcpercent is -1 and result in
    meaningless values. This is okay in a sense because we also never use
    these values if gcpercent is -1, but they're confusing when exposed to
    the user, for example via MemStats or the execution trace. It's
    particularly unfortunate in the execution trace because it attempts to
    plot the underflowed value of next_gc, which scales all useful
    information in the heap row into oblivion.
    
    Fix this by making next_gc ^0 when gcpercent < 0. This has the
    advantage of being true in a way: next_gc is effectively infinite when
    gcpercent < 0. We can also detect this special value when updating the
    execution trace and report next_gc as 0 so it doesn't blow up the
    display of the heap line.
    
    Change-Id: I4f366e4451f8892a4908da7b2b6086bdc67ca9a9
    Reviewed-on: https://go-review.googlesource.com/30016
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit f8a0c15e902cc6555610a32a6dbf36f4ff2c2682
Author: Keith Randall <khr@golang.org>
Date:   Fri Oct 7 10:10:55 2016 -0700

    test: re-enable live2 test on amd64
    
    Not sure why it was ever disabled (early SSA work?) but it passes now.
    
    Change-Id: I76439cacdbd286ce077f7e08c4d0663396a0cd8f
    Reviewed-on: https://go-review.googlesource.com/30672
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 93d5f43a29dea00a5dffd1fa8faed911a31b55fb
Author: Keith Randall <khr@golang.org>
Date:   Fri Oct 7 09:35:04 2016 -0700

    cmd/compile: do regalloc check only when checkEnabled
    
    No point doing this check all the time.
    
    Fixes #15621
    
    Change-Id: I1966c061986fe98fe9ebe146d6b9738c13cef724
    Reviewed-on: https://go-review.googlesource.com/30670
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0db9518ab30f35c7c185aed337037e8305a98b34
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 6 17:08:44 2016 -0400

    fmt: document and adjust Scanf space handling to eliminate a few paradoxes
    
    This CL makes minor adjustment to the handling of spaces and newlines
    and then documents the result.
    
    The semantic adjustment mainly concerns the handling of a run of
    spaces following a newline in the format, like in "\n ".
    Previously, that run of spaces was ignored entirely, leading to paradoxes
    like the format "1 \n 2" not matching itself as input.
    Now, spaces following a newline in the format match zero or more
    spaces following the corresponding newline in the input.
    
    The changes to the test suite show how minor the semantic adjustments are
    and how they make the behavior more regular than previously.
    
    This CL also updates the documentation to explain the handling of
    spaces more precisely, incorporating the draft from CL 17723 but
    describing the newly introduced behavior.
    
    Fixes #13565.
    
    Change-Id: I129666e9ba42de3c28b67f75cb47488e9a4c1867
    Reviewed-on: https://go-review.googlesource.com/30611
    Reviewed-by: Rob Pike <r@golang.org>

commit 23606c6fc4ed21d563f8e9a49b2ed6e18489e222
Author: Russ Cox <rsc@golang.org>
Date:   Thu Dec 10 21:43:51 2015 -0500

    fmt: add tests showing current Scanf space handling
    
    There are no semantic changes here, just tests to establish
    the status quo. A followup CL will make some semantic changes,
    the (limited) scope of which should be clear from the number of
    tests that change.
    
    For #13565.
    
    Change-Id: I960749cf59d4dfe39c324875bcc575096654f883
    Reviewed-on: https://go-review.googlesource.com/30610
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 78a267e3796e5f0a9939b1152576a072de1a2669
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Oct 6 15:50:47 2016 -0700

    cmd/compile: cache pointer and slice types
    
    Anonymous pointer and slice types are very common and identical
    anyway, so just reuse them rather than allocating new ones everywhere
    they appear.
    
    Turns out to be a small code/stack size win because SSA relies on
    gc.Type identity for reusing temporary stack slots:
    
       text    data     bss     dec     hex filename
    6453005  231643  146328 6830976  683b80 go.old
    6446660  231643  146328 6824631  6822b7 go.new
    
    Saves on memory usage during compile time too, and maybe a small CPU
    time win, but the benchmarks are pretty noisy:
    
    name       old time/op     new time/op     delta
    Template       342ms ± 8%      339ms ± 9%    ~           (p=0.332 n=99+99)
    Unicode        183ms ± 9%      181ms ±11%    ~           (p=0.274 n=95+98)
    GoTypes        1.05s ± 4%      1.04s ± 3%  -1.22%        (p=0.000 n=97+95)
    Compiler       4.49s ± 7%      4.46s ± 6%    ~           (p=0.058 n=96+91)
    
    name       old user-ns/op  new user-ns/op  delta
    Template        520M ±17%       522M ±20%    ~          (p=0.544 n=98+100)
    Unicode         331M ±27%       327M ±30%    ~           (p=0.615 n=98+98)
    GoTypes        1.54G ±10%      1.53G ±12%    ~          (p=0.173 n=99+100)
    Compiler       6.33G ±10%      6.33G ±10%    ~           (p=0.682 n=98+98)
    
    name       old alloc/op    new alloc/op    delta
    Template      44.5MB ± 0%     44.1MB ± 0%  -0.80%        (p=0.000 n=97+99)
    Unicode       37.5MB ± 0%     37.3MB ± 0%  -0.44%       (p=0.000 n=98+100)
    GoTypes        126MB ± 0%      124MB ± 0%  -1.41%        (p=0.000 n=98+99)
    Compiler       518MB ± 0%      508MB ± 0%  -1.90%       (p=0.000 n=98+100)
    
    name       old allocs/op   new allocs/op   delta
    Template        441k ± 0%       434k ± 0%  -1.76%       (p=0.000 n=100+97)
    Unicode         368k ± 0%       365k ± 0%  -0.69%        (p=0.000 n=99+99)
    GoTypes        1.26M ± 0%      1.23M ± 0%  -2.27%       (p=0.000 n=100+99)
    Compiler       4.60M ± 0%      4.46M ± 0%  -2.96%       (p=0.000 n=100+99)
    
    Change-Id: I94abce5c57aed0f9c48f567b3ac24c627d4c7c91
    Reviewed-on: https://go-review.googlesource.com/30632
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 15937ccb8915ef941e08feb2500f5acf61bd5427
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Oct 6 15:14:58 2016 -0700

    runtime: fix sigset type for ppc64 big-endian GNU/Linux
    
    On 64-bit big-endian GNU/Linux machines we need to treat sigset as a
    single uint64, not as a pair of uint32 values. This fix was already made
    for s390x, but not for ppc64 (which is big-endian--the little endian
    version is known as ppc64le). So copy os_linux_390.x to
    os_linux_be64.go, and use build constraints as needed.
    
    Fixes #17361
    
    Change-Id: Ia0eb18221a8f5056bf17675fcfeb010407a13fb0
    Reviewed-on: https://go-review.googlesource.com/30602
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a016ecfdcbc266f45f33350238777fba9a391b8d
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Oct 5 15:59:09 2016 -0700

    spec: clarify acceptable indices in array/slice composite literals
    
    This simply documents the status quo accepted by cmd/compile, gccgo,
    and go/types. The new language matches the language used for indices
    of index expressions for arrays and slices.
    
    Fixes #16679.
    
    Change-Id: I65447889fbda9d222f2a9e6c10334d1b38c555f0
    Reviewed-on: https://go-review.googlesource.com/30474
    Reviewed-by: Rob Pike <r@golang.org>

commit 95a6572b2b84e8cbfd821c4e5f774f20d37c007e
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Oct 5 15:01:15 2016 -0700

    math/big: Rat.SetString to report error if input is not consumed entirely
    
    Also, document behavior explicitly for all SetString implementations.
    
    Fixes #17001.
    
    Change-Id: Iccc882b4bc7f8b61b6092f330e405c146a80dc98
    Reviewed-on: https://go-review.googlesource.com/30472
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 4103fedf199dbc80c11455ffe75e3ecf89c77da5
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Oct 6 19:06:49 2016 +0000

    runtime: skip gdb tests on linux/ppc64 for now
    
    Updates #17366
    
    Change-Id: Ia4bd3c74c48b85f186586184a7c2b66d3b80fc9c
    Reviewed-on: https://go-review.googlesource.com/30596
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 7b4a2246670b4ec5a19214f8290fce21173f6b14
Author: Alexander Döring <email@alexd.ch>
Date:   Thu Oct 6 20:46:50 2016 +0200

    math/cmplx: add examples for Abs, Exp, Polar
    
    Updates #16360
    
    Change-Id: I941519981ff5bda3a113e14fa6be718eb4d2bf83
    Reviewed-on: https://go-review.googlesource.com/30554
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1c09b4dde645b63b34673d0b2c93f4371f7f19df
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Oct 6 18:22:19 2016 +0000

    time: document that calendrical calculations are with no leap seconds
    
    Fixes #15247
    
    Change-Id: I942fb2eacd1b54bab66cc147a6b047a3ffce0b84
    Reviewed-on: https://go-review.googlesource.com/30595
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 4c79ed5f4483f1de065ba2b409de58ae2b6907d2
Author: Yasuhiro Matsumoto <mattn.jp@gmail.com>
Date:   Wed Jan 6 21:36:31 2016 +0900

    archive/zip: handle mtime in NTFS/UNIX/ExtendedTS extra fields
    
    Handle NTFS timestamp, UNIX timestamp, Extended extra timestamp.
    Writer supports only Extended extra timestamp field, matching most
    zip creators.
    
    Fixes #10242.
    
    Change-Id: Id665db274e63def98659231391fb77392267ac1e
    Reviewed-on: https://go-review.googlesource.com/18274
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 122abe6b122a885c52a53044eedbd3b6905b6124
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Oct 6 13:49:57 2016 -0400

    cmd/compile: fold extensions into constants on s390x
    
    We insert extensions when lowering comparisons and they were
    blocking constant folding.
    
    Change-Id: I804bbf91c7606612ffe921a90853844a57e55955
    Reviewed-on: https://go-review.googlesource.com/30541
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b4f3c9339e77e70cbd5b81ef7f2a4dec557d40de
Author: Simon Rawet <simon@rawet.se>
Date:   Sat May 28 16:46:45 2016 +0200

    time: fix AddDate with nil location
    
    AddDate now retrieves location from t.Location() to ensure that
    it never calls Date with a nil location.
    
    Added test for this bug on all Time's methods
    
    Fixes #15852
    
    Change-Id: Id2a222af56993f741ad0b802a2c3b89e8e463926
    Reviewed-on: https://go-review.googlesource.com/23561
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit c79ba22ece2cbb86b00325df7712ff0803cfd6e9
Author: David Chase <drchase@google.com>
Date:   Thu Oct 6 11:08:11 2016 -0700

    test: delete sliceopt.go
    
    It tests the behavior of the old deleted compiler.
    
    Fixes #17362.
    
    Change-Id: Ia2fdec734c5cbe724a9de562ed71598f67244ab3
    Reviewed-on: https://go-review.googlesource.com/30593
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fc47ee23322d668b3b023cd8c1d3541b1ad9703e
Author: Hajime Hoshi <hajimehoshi@gmail.com>
Date:   Fri Oct 7 02:06:33 2016 +0900

    cmd/compile/internal/gc: unexport global variable Pc
    
    Change-Id: Id2a9fc1e9e70eaf5f25ddc7476061e06abcf60e4
    Reviewed-on: https://go-review.googlesource.com/30573
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 33faa4ebb9aedbe821b57b9c645d3be72100ebe3
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Oct 5 15:25:28 2016 -0700

    go/importer: better error messages when export data is not found
    
    Fixes #17281.
    
    Change-Id: I4e639998dbe3baa98879f1becc37d7c4d19351e7
    Reviewed-on: https://go-review.googlesource.com/30473
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 8aadcc551e1610d4185c36624c8105f4303fe7ec
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Oct 6 10:16:17 2016 -0400

    cmd/compile: intrinsify math.Sqrt when compiling "math" itself
    
    Fixes #17354.
    
    Change-Id: I0e018c8c3e791fc6cc1925dbbc18c2151ba9a111
    Reviewed-on: https://go-review.googlesource.com/30539
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 4f3a641e6e44fab414f7e384ac55e5b9e8d6fc7e
Author: Russ Cox <rsc@golang.org>
Date:   Thu Oct 6 10:39:50 2016 -0400

    math: fix Gamma(-171.5) on all platforms
    
    Using 387 mode was computing it without underflow to zero,
    apparently due to an 80-bit intermediate. Avoid underflow even
    with 64-bit floats.
    
    This eliminates the TODOs in the test suite.
    
    Fixes linux-386-387 build and fixes #11441.
    
    Change-Id: I8abaa63bfdf040438a95625d1cb61042f0302473
    Reviewed-on: https://go-review.googlesource.com/30540
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 20c48c9557e8d33e19a7e770d4f045ed548f26a2
Author: Richard Gibson <richard.gibson@gmail.com>
Date:   Mon Mar 28 17:15:01 2016 -0400

    encoding/json: explicitly document and test "-" key tag
    
    Struct fields can be suppressed in JSON serialization by "-" tags, but
    that doesn't preclude generation of "-" object keys.
    Document and verify the mechanism for doing so.
    
    Change-Id: I7f60e1759cfee15cb7b2447cd35fab91c5b004e6
    Reviewed-on: https://go-review.googlesource.com/21204
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
```
