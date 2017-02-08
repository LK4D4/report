# September 30, 2016 Report

Number of commits: 102

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 584.047345ms to 584.699694ms, +0.11%
* github.com/gogits/gogs: from 13.133907231s to 12.506842438s, -4.77%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 36.891499644s to 35.13498484s, -4.76%
* github.com/influxdata/influxdb/cmd/influxd: from 7.03874552s to 6.512827287s, -7.47%
* github.com/junegunn/fzf/src/fzf: from 1.147404325s to 1.021519297s, -10.97%
* github.com/mholt/caddy/caddy: from 5.993338291s to 5.842411688s, -2.52%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.312899578s to 1.337131317s, +1.85%
* github.com/nsqio/nsq/apps/nsqd: from 2.039219266s to 2.153102178s, +5.58%
* github.com/prometheus/prometheus/cmd/prometheus: from 11.208619024s to 11.733360999s, +4.68%
* github.com/spf13/hugo: from 8.054097951s to 7.66085059s, -4.88%
* golang.org/x/tools/cmd/guru: from 2.742956805s to 2.514716628s, -8.32%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2573545 to 2573942, ~
* github.com/gogits/gogs: from 23156037 to 23527751, +1.61%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 28730928 to 29106016, +1.31%
* github.com/influxdata/influxdb/cmd/influxd: from 16121288 to 16412924, +1.81%
* github.com/junegunn/fzf/src/fzf: from 3144936 to 3151120, +0.20%
* github.com/mholt/caddy/caddy: from 14622830 to 14914426, +1.99%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4366026 to 4362327, ~
* github.com/nsqio/nsq/apps/nsqd: from 9643381 to 9939114, +3.07%
* github.com/prometheus/prometheus/cmd/prometheus: from 25096437 to 25396266, +1.19%
* github.com/spf13/hugo: from 16015606 to 16298965, +1.77%
* golang.org/x/tools/cmd/guru: from 7496866 to 7493167, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               182           183           +0.55%
BenchmarkMsgpUnmarshal-4             371           372           +0.27%
BenchmarkEasyJsonMarshal-4           1431          1430          -0.07%
BenchmarkEasyJsonUnmarshal-4         1492          1683          +12.80%
BenchmarkFlatBuffersMarshal-4        525           379           -27.81%
BenchmarkFlatBuffersUnmarshal-4      280           271           -3.21%
BenchmarkGogoprotobufMarshal-4       155           155           +0.00%
BenchmarkGogoprotobufUnmarshal-4     247           237           -4.05%

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

* [expvar: export http.Handler](https://github.com/golang/go/commit/6ba5b32922c438a608a11561100a8a80abf0fd3a)
* [runtime: optimize defer code](https://github.com/golang/go/commit/f8b2314c563be4366f645536e8031a132cfdf3dd)
* [testing: add Name method to *T and *B](https://github.com/golang/go/commit/594cddd62598dcfc1fe6ee1c3e5978063f498dc1)
* [net: add Buffers type, do writev on unix](https://github.com/golang/go/commit/8e69d43b32be578cd36eebe491b6e1205dbf32a4)
* [runtime: remove defer from standard cgo call](https://github.com/golang/go/commit/441502154fa5f78e93c9c7985fbea78a02c21f4f)

## GIT Log

```
commit 962dc4b44d5260e72dad8d636cb09bcf7f4aa6bb
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Sep 27 20:30:01 2016 -0400

    cmd/compile: improve load/store merging on s390x
    
    This commit makes the process of load/store merging more incremental
    for both big and little endian operations. It also adds support for
    32-bit shifts (needed to merge 16- and 32-bit loads/stores).
    
    In addition, the merging of little endian stores is now supported.
    Little endian stores are now up to 30 times faster.
    
    Change-Id: Iefdd81eda4a65b335f23c3ff222146540083ad9c
    Reviewed-on: https://go-review.googlesource.com/29956
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 41fa42b4475fda8b0c4205827ce0e2388608cb72
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Fri Sep 30 10:35:03 2016 +0200

    cmd/compile: delete unused Convconst function
    
    Convconst is not used in the new backend, and all its callers
    were deleted in CL 29168 (cmd/compile: delete lots of the legacy
    backend). iconv was an helper function for Convconst.
    
    Updates #16357
    
    Change-Id: I65c7345586d7af81cdc2fb09c68f744ffb161a17
    Reviewed-on: https://go-review.googlesource.com/30090
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 893d6866213ca539195076f0e3338da99f321c9c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Sep 29 18:05:06 2016 -0700

    net/http: update bundled http2, add h2 Transport.IdleConnTimeout tests
    
    Updates bundled http2 to x/net git rev a333c53 for:
    
       http2: add Transport support for IdleConnTimeout
       https://golang.org/cl/30075
    
    And add tests.
    
    The bundled http2 also includes a change adding a Ping method to
    http2.ClientConn, but that type isn't exposed in the standard
    library. Nevertheless, the code gets compiled and adds a dependency on
    "crypto/rand", requiring an update to go/build's dependency
    test. Because net/http already depends on crypto/tls, which uses
    crypto/rand, it's not really a new dependency.
    
    Fixes #16808
    
    Change-Id: I1ec8666ea74762f27c70a6f30a366a6647f923f7
    Reviewed-on: https://go-review.googlesource.com/30078
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 441502154fa5f78e93c9c7985fbea78a02c21f4f
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Sep 29 20:21:52 2016 -0400

    runtime: remove defer from standard cgo call
    
    The endcgo function call is currently deferred in case a cgo
    callback into Go panics and unwinds through cgocall. Typical cgo
    calls do not have callbacks into Go, and even fewer panic, so we
    pay the cost of this defer for no typical benefit.
    
    Amazingly, there is another defer on the cgocallback path also used
    to cleanup in case the Go code called by cgo panics. This CL folds
    the first defer into the second, to reduce the cost of typical cgo
    calls.
    
    This reduces the overhead for a no-op cgo call significantly:
    
            name       old time/op  new time/op  delta
            CgoNoop-8  93.5ns ± 0%  51.1ns ± 1%  -45.34%  (p=0.016 n=4+5)
    
    The total effect between Go 1.7 and 1.8 is even greater, as CL 29656
    reduced the cost of defer recently. Hopefully a future Go release
    will drop the cost of defer to nothing, making this optimization
    unnecessary. But until then, this is nice.
    
    Change-Id: Id1a5648f687a87001d95bec6842e4054bd20ee4f
    Reviewed-on: https://go-review.googlesource.com/30080
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9828b7c468029d54e90846e5e2fc23fd6d39782a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 29 21:34:44 2016 -0700

    runtime, syscall: use FP instead of SP for parameters
    
    Consistently access function parameters using the FP pseudo-register
    instead of SP (e.g., x+0(FP) instead of x+4(SP) or x+8(SP), depending
    on register size). Two reasons: 1) doc/asm says the SP pseudo-register
    should use negative offsets in the range [-framesize, 0), and 2)
    cmd/vet only validates parameter offsets when indexed from the FP
    pseudo-register.
    
    No binary changes to the compiled object files for any of the affected
    package/OS/arch combinations.
    
    Change-Id: I0efc6079bc7519fcea588c114ec6a39b245d68b0
    Reviewed-on: https://go-review.googlesource.com/30085
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 79db1625b929b8dad46c1537175b9412fd020851
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 29 19:09:36 2016 -0700

    cmd/compile: eliminate stkdelta
    
    At this point in the compiler we haven't assigned Xoffset values for
    PAUTO variables anyway, so just immediately store the stack offsets
    into Xoffset rather than into a global map.
    
    Change-Id: I61eb471c857c8b145fd0895cbd98fd4e8d3c3365
    Reviewed-on: https://go-review.googlesource.com/30081
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit bb2bbfa08630c65b8751159515f3a22ec5f933ee
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 29 16:22:43 2016 -0700

    cmd/compile: split Addrconst out of Naddr
    
    There are only three Prog types that we were creating with an OLITERAL
    Node: ATEXT, ATYPE, and AFUNCDATA. ATEXT's value we later overwrite in
    defframe, and ATYPE's we don't even need. AFUNCDATA only needs integer
    constants, so get rid of all the non-int constant logic and skip
    creating a Node representation for the constant.
    
    While here, there are a few other Naddr code paths that are no longer
    needed, so turn those into Fatalfs.
    
    Passes toolstash/buildall.
    
    Change-Id: I4cc9b92c3011890afd4f31ebeba8b1b42b753cab
    Reviewed-on: https://go-review.googlesource.com/30074
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 214bf6809791170b88b41249850f2f01534725c4
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Sep 29 19:35:43 2016 -0700

    net/http: remove TODO about the Server's base context
    
    I decided not to expand the API for this per discusion on #16220.
    
    Fixes #16220
    
    Change-Id: I65cb2eacd4ec28c79438a8f7c30024524a484ce6
    Reviewed-on: https://go-review.googlesource.com/30082
    Reviewed-by: Daniel Theophanes <kardianos@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit dd748cf3141d45eeeedcb9b88aa11e47b16a0008
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 29 15:43:10 2016 -0700

    cmd/compile: make Afunclit the default/only behavior for Naddr
    
    Naddr used to translate PFUNC Nodes into references to the function
    literal wrapper, and then Afunclit could be used to rewrite it to
    reference the function text itself. But now everywhere we use Naddr on
    PFUNC Nodes, we immediately call Afunclit anyway. So just merge
    Afunclit's behavior into Naddr.
    
    Passes toolstash/buildall.
    
    Change-Id: If2ca6d7f314c1a0711df9b8209aace16ba4b8bc0
    Reviewed-on: https://go-review.googlesource.com/30073
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d2df8498f366669acbae24f38e3683b3acdab102
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Wed Sep 28 12:51:39 2016 -0700

    database/sql: close Rows when context is cancelled
    
    To prevent leaking connections, close any open Rows when the
    context is cancelled. Also enforce context cancel while reading
    rows off of the wire.
    
    Change-Id: I62237ecdb7d250d6734f6ce3d2b0bcb16dc6fda7
    Reviewed-on: https://go-review.googlesource.com/29957
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6dc356a76a405ff12c884ab0a4acb2296d1618b7
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 29 15:08:37 2016 -0700

    cmd/compile/internal/ssa: erase register copies deterministically
    
    Fixes #17288.
    
    Change-Id: I2ddd01d14667d5c6a2e19bd70489da8d9869d308
    Reviewed-on: https://go-review.googlesource.com/30072
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c1e06dcb611822ba3c881b170f278d18237372c9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Sep 29 08:51:29 2016 -0700

    net/http: use atomic.Value for Transport's alternate protocol map
    
    Fix an old TODO and use atomic.Value for holding the Transport's
    alternate protocol map. It is very frequently accessed and almost
    never set or updated.
    
    Change-Id: Ic5a71c504bdac76678114c6390d1fc0673e07aa9
    Reviewed-on: https://go-review.googlesource.com/29967
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8e69d43b32be578cd36eebe491b6e1205dbf32a4
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 27 20:50:57 2016 +0000

    net: add Buffers type, do writev on unix
    
    No fast path currently for solaris, windows, nacl, plan9.
    
    Fixes #13451
    
    Change-Id: I24b3233a2e3a57fc6445e276a5c0d7b097884007
    Reviewed-on: https://go-review.googlesource.com/29951
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ffd1c781b77aab542713b66ef387fa9307e4060b
Author: Nodir Turakulov <nodir@google.com>
Date:   Sat Sep 5 06:38:13 2015 -0700

    html/template: check "type" attribute in <script>
    
    Currently any script tag is treated as a javascript container, although
    <script type="text/template"> must not be. Check "type" attribute of
    "script" tag. If it is present and it is not a JS MIME type, do not
    transition to elementScript state.
    
    Fixes #12149, where // inside text template was treated as regexp.
    Fixes #6701
    
    Change-Id: I8fc9e504f7280bdd800f40383c061853665ac8a2
    Reviewed-on: https://go-review.googlesource.com/14336
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit f5516559e65175887f2fadb73cd8e5fdfc44bcd6
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 29 11:38:47 2016 -0700

    cmd/compile/internal/x86: fix zero-range merging logic
    
    All other architectures merge stack-zeroing ranges if there are at
    most two pointers/registers of memory between them, but x86 is
    erroneously coded to require *exactly* two.
    
    Shaves a tiny amount of text size off cmd/go when building for
    GOARCH=386 and eliminates an unnecessary inconsistency between x86's
    defframe and the other GOARCHes'.
    
       text    data     bss     dec     hex filename
    5241015  191051   93336 5525402  544f9a go.before
    5240224  191051   93336 5524611  544c83 go.after
    
    Change-Id: Ib15ec8c07bca11e824640f0ab32abfc4bb160496
    Reviewed-on: https://go-review.googlesource.com/30050
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2b49d129c48b0036ee595fbf7b91a0e9292f87ee
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Thu Sep 29 19:49:36 2016 +0200

    cmd/compile: delete unused (*Node) SetInt, SetBigInt, Bool
    
    Introduced in CL 9263 (prepare to unexport gc.Mp*) and CL 9267
    (prepare Node.Val to be unexported), their only callers were in
    the old backend and all got deleted in CL 29168 (cmd/compile:
    delete lots of the legacy backend).
    
    Update #16357
    
    Change-Id: I0a5d76b98b418e8ec0984c033c3bc0ac3fc5f38a
    Reviewed-on: https://go-review.googlesource.com/29997
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6fea452e38012e167e8a8f08f571e0240b248c97
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Fri Sep 2 20:15:12 2016 -0700

    archive/tar: move parse/format functionality into strconv.go
    
    Move all parse/format related functionality into strconv.go
    and thoroughly test them. This also reduces the amount of noise
    inside reader.go and writer.go.
    
    There was zero functionality change other than moving code around.
    
    Change-Id: I3bc288d10c20ebb3814b30b75d8acd7be62b85d7
    Reviewed-on: https://go-review.googlesource.com/28470
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a09e1de0ea7fdc30f3761d12fe52248946c08205
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Fri Sep 16 16:46:19 2016 -0700

    net/http: document how Request.Cookie deals with duplicate cookies
    
    RFC 6265, section 4.2.2 says:
    <<<
    Although cookies are serialized linearly in the Cookie header,
    servers SHOULD NOT rely upon the serialization order.  In particular,
    if the Cookie header contains two cookies with the same name (e.g.,
    that were set with different Path or Domain attributes), servers
    SHOULD NOT rely upon the order in which these cookies appear in the
    header.
    >>>
    
    This statement seems to indicate that cookies should conceptually
    be thought of as a map of keys to sets of values (map[key][]value).
    However, in practice, everyone pretty much treats cookies as a
    map[key]value and the API for Request.Cookie seems to indicate that.
    
    We should update the documentation for Request.Cookie to warn the
    user what happens when there is are multiple cookies with the same
    key. I deliberately did not want to say *which* cookie is returned.
    
    Change-Id: Id3e0e24b2b14ca2d9ea8b13f82ba739edaa71cf0
    Reviewed-on: https://go-review.googlesource.com/29364
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 518cc7f3079bbd5ad141efb1e16c5a6eae52b831
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Sep 29 11:00:47 2016 -0700

    cmd/internal/obj/arm: cleanup some unnecessary temps and conversions
    
    Change-Id: I573278c9aee80e62463b2542774dabeec7c3b098
    Reviewed-on: https://go-review.googlesource.com/29969
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 985d3d307c3669094f77b52caffef60157b7d648
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Sep 29 09:09:14 2016 -0700

    net: make proto and port lookups fall back to baked-in maps on Windows
    
    In https://golang.org/cl/28951 I cleaned up the lookupProtocol and
    lookupPort paths to be consistently case-insensitive across operating
    systems and to share the same baked-in maps of port & proto values
    that can be relied on to exist on any platform.
    
    I missed the fallback to the baked-in maps on Windows, though, which
    broke Windows XP. This should fix it.
    
    Fixes #17175
    
    Change-Id: Iecd434fb684304137ee27f5521cfaa8c351a1bde
    Reviewed-on: https://go-review.googlesource.com/29968
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d58219e50ba1fd9bf577be7332bdcabe0ef8b7d5
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Thu Sep 29 18:46:24 2016 +0200

    cmd/compile: delete unused IntLiteral function
    
    IntLiteral was only called by the gins functions in
    cmd/compile/internal/{arm64,mips64,ppc64}/gsubr.go
    but CL 29220 (cmd/compile: remove gins) deleted them,
    so IntLiteral is now unused.
    
    Change-Id: I2652b6d2ace6fdadc1982f65e749f3982513371e
    Reviewed-on: https://go-review.googlesource.com/29996
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c5434f2973a87acff76bac359236e690d632ce95
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Thu Sep 29 13:59:10 2016 +0200

    time: update test for tzdata-2016g
    
    Fixes #17276
    
    Change-Id: I0188cf9bc5fdb48c71ad929cc54206d03e0b96e4
    Reviewed-on: https://go-review.googlesource.com/29995
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 731b3ed18dcb854912cb06b3486bb633917e4cb7
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Wed Sep 14 15:23:41 2016 +0300

    math: make sqrt smaller on AMD64
    
    This makes function fit in 16 bytes, saving 16 bytes.
    
    Change-Id: Iac5d2add42f6dae985b2a5cbe19ad4bd4bcc92ec
    Reviewed-on: https://go-review.googlesource.com/29151
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit fedb0b30188952dc082672cbd45b39a49136d29c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Sep 29 09:22:01 2016 -0400

    cmd/internal/obj/arm: optimize MOVW $-off(R), R
    
    When offset < 0 and -offset fits in instruction, generate SUB
    instruction, instead of ADD with constant from the pool.
    
    Fixes #13280.
    
    Change-Id: I57d97fe9300fe1f6554365e2262393ef50acbdd3
    Reviewed-on: https://go-review.googlesource.com/30014
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 26531b3846694b311aec87fc4f5de65b4cc48a44
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Wed Sep 14 16:56:24 2016 +0300

    cmd/internal/obj/x86: cleanup
    
    Remove duplicate vars, commented out code  and duplicate lines.
    When choosing between 2 aliases, on with more uses was chosen.
    
    Change-Id: I7bc15f1693de3f6d378cef9c09469970a659db40
    Reviewed-on: https://go-review.googlesource.com/29152
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 03a1dc3522f99bf5045fc41730e6682c1cc7402a
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Sep 28 15:04:00 2016 -0700

    cmd/compile: don't crash on (unsafe.Sizeof)(0)
    
    Fixes #17270.
    
    Change-Id: I4affa57e10baf1a758bc0977265d160f220b2945
    Reviewed-on: https://go-review.googlesource.com/29960
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4d07d3e29c467484801b84dfeb762d2ee00979a9
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Sep 28 10:20:24 2016 -0400

    cmd/compile: re-enable nilcheck removal for newobject
    
    Also add compiler debug ouput and add a test.
    
    Fixes #15390.
    
    Change-Id: Iceba1414c29bcc213b87837387bf8ded1f3157f1
    Reviewed-on: https://go-review.googlesource.com/30011
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 456a01ac47c1473d6b241c56eb8db0cb832d9be2
Author: Blixt <me@blixt.nyc>
Date:   Wed Aug 24 12:59:01 2016 -0400

    encoding/binary: add bool support
    
    This change adds support for decoding and encoding the bool type. The
    encoding is a single byte, with a zero value for false and a non-zero
    value for true.
    
    Closes #16856.
    
    Change-Id: I1d1114b320263691473bb100cad0f380e0204186
    Reviewed-on: https://go-review.googlesource.com/28514
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 5e7bae5082dcc2cd241b40c2bf50db6b28943827
Author: Andrew Gerrand <adg@golang.org>
Date:   Wed Sep 28 20:15:57 2016 +1000

    doc: add testing Name method to go1.8.txt
    
    Change-Id: I6d413f747e6a6c30c5e0e9afdffd5ec18dce7e08
    Reviewed-on: https://go-review.googlesource.com/29974
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a0d330f1ddeaa55cbd35191814546fd98e09a117
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Sep 28 17:26:12 2016 +0900

    runtime, runtime/cgo: revert CL 18814; don't drop signal stack in new thread on dragonfly
    
    This change reverts CL 18814 which is a workaroud for older DragonFly
    BSD kernels, and fixes #13945 and #13947 in a more general way the
    same as other platforms except NetBSD.
    
    This is a followup to CL 29491.
    
    Updates #16329.
    
    Change-Id: I771670bc672c827f2b3dbc7fd7417c49897cb991
    Reviewed-on: https://go-review.googlesource.com/29971
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit eb268cb321edf6e2bbaa832acb2e61db6b081f98
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Sep 27 22:24:51 2016 -0700

    runtime: minor simplifications to signal code
    
    Change setsig, setsigstack, getsig, raise, raiseproc to take uint32 for
    signal number parameter, as that is the type mostly used for signal
    numbers.  Same for dieFromSignal, sigInstallGoHandler, raisebadsignal.
    
    Remove setsig restart parameter, as it is always either true or
    irrelevant.
    
    Don't check the handler in setsigstack, as the only caller does that
    anyhow.
    
    Don't bother to convert the handler from sigtramp to sighandler in
    getsig, as it will never be called when the handler is sigtramp or
    sighandler.
    
    Don't check the return value from rt_sigaction in the GNU/Linux version
    of setsigstack; no other setsigstack checks it, and it never fails.
    
    Change-Id: I6bbd677e048a77eddf974dd3d017bc3c560fbd48
    Reviewed-on: https://go-review.googlesource.com/29953
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 594cddd62598dcfc1fe6ee1c3e5978063f498dc1
Author: Andrew Gerrand <adg@golang.org>
Date:   Wed Sep 28 13:31:33 2016 +1000

    testing: add Name method to *T and *B
    
    Fixes #17231
    
    Change-Id: I0d6007ab504f2277cb6affc9e2050157a6ad4d5e
    Reviewed-on: https://go-review.googlesource.com/29970
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit db82cf4e506938a36a57a64dbe1f79eb0365ea89
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Mon Sep 26 16:44:35 2016 +1000

    runtime: use RtlGenRandom instead of CryptGenRandom
    
    This change replaces the use of CryptGenRandom with RtlGenRandom in
    Windows to generate cryptographically random numbers during process
    startup. RtlGenRandom uses the same RNG as CryptGenRandom, but it has many
    fewer DLL dependencies and so does not affect process startup time as
    much.
    
    This makes running simple Go program on my computers faster.
    
    Windows XP:
    benchmark                      old ns/op     new ns/op     delta
    BenchmarkRunningGoProgram-2     47408573      10784148      -77.25%
    
    Windows 7 (VM):
    benchmark                    old ns/op     new ns/op     delta
    BenchmarkRunningGoProgram     16260390      12792150      -21.33%
    
    Windows 7:
    benchmark                      old ns/op     new ns/op     delta
    BenchmarkRunningGoProgram-2     13600778      10050574      -26.10%
    
    Fixes #15589
    
    Change-Id: I2816239a2056e3d4a6dcd86a6fa2bb619c6008fe
    Reviewed-on: https://go-review.googlesource.com/29700
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 98938189a16a764007bce7fcd4bfeb2386043208
Author: Keith Randall <khr@golang.org>
Date:   Tue Sep 27 14:39:27 2016 -0700

    cmd/compile: remove duplicate nilchecks
    
    Mark nil check operations as faulting if their arg is zero.
    This lets the late nilcheck pass remove duplicates.
    
    Fixes #17242.
    
    Change-Id: I4c9938d8a5a1e43edd85b4a66f0b34004860bcd9
    Reviewed-on: https://go-review.googlesource.com/29952
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit ba94dd34385af3352660fb4bfa2a2d97fb937088
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Sep 26 13:21:21 2016 -0400

    math: add some assembly implementations on ARM64
    
    Also add GP<->FP move addressing mode to FMOVS, FMOVD
    instructions.
    
    Ceil-8                 37.1ns ± 0%   7.9ns ± 0%  -78.64%          (p=0.000 n=4+5)
    Dim-8                  20.9ns ± 1%  11.3ns ± 0%  -45.93%          (p=0.008 n=5+5)
    Floor-8                22.9ns ± 0%   7.9ns ± 0%  -65.41%          (p=0.029 n=4+4)
    Gamma-8                 117ns ± 0%    94ns ± 1%  -19.50%          (p=0.016 n=4+5)
    PowInt-8                121ns ± 0%   108ns ± 1%  -11.07%          (p=0.008 n=5+5)
    PowFrac-8               331ns ± 0%   318ns ± 0%   -3.93%          (p=0.000 n=5+4)
    Trunc-8                18.8ns ± 0%   7.9ns ± 0%  -57.83%          (p=0.016 n=4+5)
    
    Change-Id: I709b7f1a914b28adc27414522db551e2630cfb92
    Reviewed-on: https://go-review.googlesource.com/29734
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit a73020f847b3cf8575250569ebefb02573d19224
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 27 18:27:02 2016 +0000

    net/http: add more IDNA2008 tests and fix some omissions
    
    It wasn't lowercasing the string, folding widths, and putting strings
    into NFC form. Do those.
    
    Fixes #13835
    
    Change-Id: Ia3de6159417cacec203b48e206e51d79f945df58
    Reviewed-on: https://go-review.googlesource.com/29860
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>

commit 36c164ec9c8597ec3507bb4aa9390603e9d29f7a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 27 18:28:29 2016 +0000

    vendor: add golang.org/x/text/unicode/norm + x/test/width for IDNA support
    
    Add golang.org/x/text/unicode/norm from x/text git rev a7c02369.
    
    Needed by net/http for IDNA normalization.
    
    Updates #13835
    
    Change-Id: I8b024e179d573f2b093c209a4b9e4f71f7d4a1f2
    Reviewed-on: https://go-review.googlesource.com/29859
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>

commit 15b4d187b8dceb3384dd3bf8ea9f368c517f082d
Author: mike andrews <mra@xoba.com>
Date:   Tue Sep 27 16:22:29 2016 -0400

        encoding/json: fix a bug in the documentation
    
        Documentation made reference to an unknown entity "DisableHTMLEscaping,"
        but I think it actually meant the method "Encoder.SetEscapeHTML."
    
        Fixes #17255
    
    Change-Id: I18fda76f8066110caef85fd33698de83d632e646
    Reviewed-on: https://go-review.googlesource.com/29931
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit fdc167164ecde259bd356cc8e7ae5ccb0903469c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Sep 27 13:42:28 2016 -0700

    runtime: remove sigmask type, use sigset instead
    
    The OS-independent sigmask type was not pulling its weight. Replace it
    with the OS-dependent sigset type. This requires adding an OS-specific
    sigaddset function, but permits removing the OS-specific sigmaskToSigset
    function.
    
    Change-Id: I43307b512b0264ec291baadaea902f05ce212305
    Reviewed-on: https://go-review.googlesource.com/29950
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e13df02e5fbf2c0cd8811b826a8c8567efa882dd
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Mon Sep 19 11:19:32 2016 -0700

    database/sql: add context methods
    
    Add context methods to sql and sql/driver methods. If
    the driver doesn't implement context methods the connection
    pool will still handle timeouts when a query fails to return
    in time or when a connection is not available from the pool
    in time.
    
    There will be a follow-up CL that will add support for
    context values that specify transaction levels and modes
    that a driver can use.
    
    Fixes #15123
    
    Change-Id: Ia99f3957aa3f177b23044dd99d4ec217491a30a7
    Reviewed-on: https://go-review.googlesource.com/29381
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 54a72d90f62030034f03cacbac1d1ec02c0444c6
Author: Luigi Riefolo <luigi.riefolo@gmail.com>
Date:   Tue Sep 27 03:02:40 2016 +0200

    go/doc: add IsPredeclared function
    
    IsPredeclared allows simplifying src/golang.org/x/tools/godoc/linkify.go
    
    Change-Id: I56b3223896f844630bc2e940255572d1682f0d06
    Reviewed-on: https://go-review.googlesource.com/29870
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 097a581dc0d97efac1dfbe5d79819bbf6bf681a7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Sep 27 07:20:10 2016 -0700

    runtime: simplify signalstack by dropping nil as argument
    
    Change the two calls to signalstack(nil) to inline the code
    instead (it's two lines).
    
    Change-Id: Ie92a05494f924f279e40ac159f1b677fda18f281
    Reviewed-on: https://go-review.googlesource.com/29854
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 60482d8a8b11a3dfdf9b582b9f666694d84bb9c4
Author: Elias Naur <elias.naur@gmail.com>
Date:   Thu Sep 22 09:37:28 2016 +0200

    runtime: relax SetFinalizer documentation to allow &local
    
    The SetFinalizer documentation states that
    
    "The argument obj must be a pointer to an object allocated by calling
    new or by taking the address of a composite literal."
    
    which precludes pointers to local variables. According to a comment
    on #6591, this case is expected to work. This CL updates the documentation
    for SetFinalizer accordingly.
    
    Fixes #6591
    
    Change-Id: Id861b3436bc1c9521361ea2d51c1ce74a121c1af
    Reviewed-on: https://go-review.googlesource.com/29592
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 413f1ef4de98f2e98853c3753f785871bfc32e5d
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Sep 27 10:29:43 2016 -0400

    cmd/asm: fix parsing of the s390x VLE{G,F,H,B} instructions
    
    This commit makes the assembler frontend reorder the operands so that
    they are in the order the backend expects. The index should be first
    for consistency with the other vector instructions.
    
    Before this commit no operand order would have been accepted so this
    isn't a breaking change.
    
    Change-Id: I188d57eeb338d27fa1fa6845de0d6d1521b7a6c3
    Reviewed-on: https://go-review.googlesource.com/29855
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a30b5a3d1916deb6e366aa4557d6c21eb835d737
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Sep 27 10:45:05 2016 -0400

    cmd/asm: add s390x instructions BLTU and BLEU
    
    These instructions are the same as BLT and BLE except that they
    also branch if the 'unordered' bit is set in the condition code.
    
    They are already used by the SSA backend. This change allows them
    to be used in hand-written assembly code.
    
    Change-Id: Ie9b5985a5e87ea22e8043567a286e09dce16a2db
    Reviewed-on: https://go-review.googlesource.com/29930
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 17a8ec2c4f702039652a4bc9630d233b454cfae8
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Sep 26 00:45:01 2016 -0400

    cmd/asm, cmd/internal/obj/s390x: improve add/multiply-immediate codegen
    
    Use the A{,G}HI instructions where possible (4 bytes instead of 6 bytes
    for A{,G}FI). Also, use 32-bit operations where appropriate for
    multiplication.
    
    Change-Id: I4041781cda26be52b54e4804a9e71552310762d0
    Reviewed-on: https://go-review.googlesource.com/29733
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3436f0776f4f373b5ba1aacf9f66689c833168b0
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Aug 30 15:25:28 2016 -0400

    reflect, runtime: optimize Value.Call on s390x and add benchmark
    
    Use an MVC loop to copy arguments in runtime.call* rather than copying
    bytes individually.
    
    I've added the benchmark CallArgCopy to test the speed of Value.Call
    for various argument sizes.
    
    name                    old speed      new speed       delta
    CallArgCopy/size=128     439MB/s ± 1%    582MB/s ± 1%   +32.41%  (p=0.000 n=10+10)
    CallArgCopy/size=256     695MB/s ± 1%   1172MB/s ± 1%   +68.67%  (p=0.000 n=10+10)
    CallArgCopy/size=1024    573MB/s ± 8%   4175MB/s ± 2%  +628.11%  (p=0.000 n=10+10)
    CallArgCopy/size=4096   1.46GB/s ± 2%  10.19GB/s ± 1%  +600.52%  (p=0.000 n=10+10)
    CallArgCopy/size=65536  1.51GB/s ± 0%  12.30GB/s ± 1%  +716.30%   (p=0.000 n=9+10)
    
    Change-Id: I87dae4809330e7964f6cb4a9e40e5b3254dd519d
    Reviewed-on: https://go-review.googlesource.com/28096
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6c43c0c2fdc866bf6d85a4689ad42c42b9c527dd
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Tue Sep 27 17:23:08 2016 +0200

    cmd/compile: remove commented-out old c code
    
    Change-Id: I9b2e6c45f7e83543a06d0aafd08a911f7b6485fd
    Reviewed-on: https://go-review.googlesource.com/29874
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f876fb9baeef1798b2f6c30fde8f695b127fdad2
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Sep 23 09:15:51 2016 -0400

    cmd/compile: move value around before kick it out of register
    
    When allocating registers, before kicking out the existing value,
    copy it to a spare register if there is one. So later use of this
    value can be found in register instead of reload from spill. This
    is very helpful for instructions of which the input and/or output
    can only be in specific registers, e.g. DIV on x86, MUL/DIV on
    MIPS. May also be helpful in general.
    
    For "go build -a cmd/go" on AMD64, reduce "spilled value remains"
    by 1% (not including args, which almost certainly remain).
    
    For the code in issue #16061 on AMD64:
    MaxRem-12   111µs ± 1%    94µs ± 0%  -15.38%  (p=0.008 n=5+5)
    
    Go1 benchmark on AMD64:
    BinaryTree17-12              2.32s ± 2%     2.30s ± 1%    ~     (p=0.421 n=5+5)
    Fannkuch11-12                2.52s ± 0%     2.44s ± 0%  -3.44%  (p=0.008 n=5+5)
    FmtFprintfEmpty-12          39.9ns ± 3%    39.8ns ± 0%    ~     (p=0.635 n=5+4)
    FmtFprintfString-12          114ns ± 1%     113ns ± 1%    ~     (p=0.905 n=5+5)
    FmtFprintfInt-12             102ns ± 6%      98ns ± 1%    ~     (p=0.087 n=5+5)
    FmtFprintfIntInt-12          146ns ± 5%     147ns ± 1%    ~     (p=0.238 n=5+5)
    FmtFprintfPrefixedInt-12     155ns ± 2%     151ns ± 1%  -2.58%  (p=0.008 n=5+5)
    FmtFprintfFloat-12           231ns ± 1%     232ns ± 1%    ~     (p=0.286 n=5+5)
    FmtManyArgs-12               657ns ± 1%     649ns ± 0%  -1.31%  (p=0.008 n=5+5)
    GobDecode-12                6.35ms ± 0%    6.29ms ± 1%    ~     (p=0.056 n=5+5)
    GobEncode-12                5.38ms ± 1%    5.45ms ± 1%    ~     (p=0.056 n=5+5)
    Gzip-12                      209ms ± 0%     209ms ± 1%    ~     (p=0.690 n=5+5)
    Gunzip-12                   31.2ms ± 1%    31.1ms ± 1%    ~     (p=0.548 n=5+5)
    HTTPClientServer-12          123µs ± 4%     130µs ± 8%    ~     (p=0.151 n=5+5)
    JSONEncode-12               14.0ms ± 1%    14.0ms ± 1%    ~     (p=0.421 n=5+5)
    JSONDecode-12               41.2ms ± 1%    41.1ms ± 2%    ~     (p=0.421 n=5+5)
    Mandelbrot200-12            3.96ms ± 1%    3.98ms ± 0%    ~     (p=0.421 n=5+5)
    GoParse-12                  2.88ms ± 1%    2.88ms ± 1%    ~     (p=0.841 n=5+5)
    RegexpMatchEasy0_32-12      68.0ns ± 3%    66.6ns ± 1%  -2.00%  (p=0.024 n=5+5)
    RegexpMatchEasy0_1K-12       728ns ± 8%     682ns ± 1%  -6.26%  (p=0.008 n=5+5)
    RegexpMatchEasy1_32-12      66.8ns ± 2%    66.0ns ± 1%    ~     (p=0.302 n=5+5)
    RegexpMatchEasy1_1K-12       291ns ± 2%     288ns ± 1%    ~     (p=0.111 n=5+5)
    RegexpMatchMedium_32-12      103ns ± 2%     100ns ± 0%  -2.53%  (p=0.016 n=5+4)
    RegexpMatchMedium_1K-12     31.9µs ± 1%    31.3µs ± 0%  -1.75%  (p=0.008 n=5+5)
    RegexpMatchHard_32-12       1.59µs ± 2%    1.59µs ± 1%    ~     (p=0.548 n=5+5)
    RegexpMatchHard_1K-12       48.3µs ± 2%    47.7µs ± 1%    ~     (p=0.222 n=5+5)
    Revcomp-12                   340ms ± 1%     338ms ± 1%    ~     (p=0.421 n=5+5)
    Template-12                 46.3ms ± 1%    46.5ms ± 1%    ~     (p=0.690 n=5+5)
    TimeParse-12                 252ns ± 1%     247ns ± 0%  -1.91%  (p=0.000 n=5+4)
    TimeFormat-12                277ns ± 1%     267ns ± 0%  -3.82%  (p=0.008 n=5+5)
    [Geo mean]                  48.8µs         48.3µs       -0.93%
    
    It has very little effect on binary size and compiler speed.
    compilebench:
    Template       230ms ±10%      231ms ± 8%    ~             (p=0.546 n=9+9)
    Unicode        123ms ± 6%      124ms ± 9%    ~           (p=0.481 n=10+10)
    GoTypes        742ms ± 6%      755ms ± 3%    ~           (p=0.123 n=10+10)
    Compiler       3.10s ± 3%      3.08s ± 1%    ~           (p=0.631 n=10+10)
    
    Fixes #16061.
    
    Change-Id: Id99cdc7a182ee10a704fa0f04e8e0d0809b2ac56
    Reviewed-on: https://go-review.googlesource.com/29732
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 9d4b40f55d2298fcb69e049b031e2e3ce8a1de8c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Sep 27 08:57:02 2016 -0400

    runtime, cmd/compile: implement and use DUFFCOPY on ARM64
    
    Change-Id: I8984eac30e5df78d4b94f19412135d3cc36969f8
    Reviewed-on: https://go-review.googlesource.com/29910
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 7de7d20e9f14c3a8fb9a3bcf6a36a299c74e9ddd
Author: Oliver Tonnhofer <olt@bogosoft.com>
Date:   Tue Sep 27 15:24:00 2016 +0200

    image/png: improve compression by skipping filter for paletted images
    
    Compression of paletted images is more efficient if they are not filtered.
    This patch skips filtering for cbP8 images.
    The improvements are demonstrated at https://github.com/olt/compressbench
    
    Fixes #16196
    
    Change-Id: Ie973aad287cacf9057e394bb01cf0e4448a77618
    Reviewed-on: https://go-review.googlesource.com/29872
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 590f3f0c9dffdadeeec4d9e79af6e0974a574a1b
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sat Sep 24 21:38:58 2016 +0200

    cmd/compile: fix misaligned comments
    
    Realign multi-line comments that got misaligned by the c->go
    conversion.
    
    Change-Id: I584b902e95cf588aa14febf1e0b6dfa499c303c2
    Reviewed-on: https://go-review.googlesource.com/29871
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8f9e2ab55786d37158a3aaf27d054944e0742717
Author: Sam Whited <sam@samwhited.com>
Date:   Mon Sep 26 20:23:36 2016 -0500

    database/sql: add doc comment for ErrTxDone
    
    Change-Id: Idffb82cdcba4985954d061bdb021217f47ff4985
    Reviewed-on: https://go-review.googlesource.com/29850
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9ed0715bb66dbbd0f597f93d0bc70a3d769b1b10
Author: Michal Bohuslávek <mbohuslavek@gmail.com>
Date:   Tue Sep 20 22:56:57 2016 +0100

    math/big: support negative numbers in ModInverse
    
    Fixes #16984
    
    Change-Id: I3a330e82941a068ca6097985af4ab221275fd336
    Reviewed-on: https://go-review.googlesource.com/29299
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit f8b2314c563be4366f645536e8031a132cfdf3dd
Author: Austin Clements <austin@google.com>
Date:   Thu Sep 22 17:08:37 2016 -0400

    runtime: optimize defer code
    
    This optimizes deferproc and deferreturn in various ways.
    
    The most important optimization is that it more carefully arranges to
    prevent preemption or stack growth. Currently we do this by switching
    to the system stack on every deferproc and every deferreturn. While we
    need to be on the system stack for the slow path of allocating and
    freeing defers, in the common case we can fit in the nosplit stack.
    Hence, this change pushes the system stack switch down into the slow
    paths and makes everything now exposed to the user stack nosplit. This
    also eliminates the need for various acquirem/releasem pairs, since we
    are now preventing preemption by preventing stack split checks.
    
    As another smaller optimization, we special case the common cases of
    zero-sized and pointer-sized defer frames to respectively skip the
    copy and perform the copy in line instead of calling memmove.
    
    This speeds up the runtime defer benchmark by 42%:
    
    name           old time/op  new time/op  delta
    Defer-4        75.1ns ± 1%  43.3ns ± 1%  -42.31%   (p=0.000 n=8+10)
    
    In reality, this speeds up defer by about 2.2X. The two benchmarks
    below compare a Lock/defer Unlock pair (DeferLock) with a Lock/Unlock
    pair (NoDeferLock). NoDeferLock establishes a baseline cost, so these
    two benchmarks together show that this change reduces the overhead of
    defer from 61.4ns to 27.9ns.
    
    name           old time/op  new time/op  delta
    DeferLock-4    77.4ns ± 1%  43.9ns ± 1%  -43.31%  (p=0.000 n=10+10)
    NoDeferLock-4  16.0ns ± 0%  15.9ns ± 0%   -0.39%    (p=0.000 n=9+8)
    
    This also shaves 34ns off cgo calls:
    
    name       old time/op  new time/op  delta
    CgoNoop-4   122ns ± 1%  88.3ns ± 1%  -27.72%  (p=0.000 n=8+9)
    
    Updates #14939, #16051.
    
    Change-Id: I2baa0dea378b7e4efebbee8fca919a97d5e15f38
    Reviewed-on: https://go-review.googlesource.com/29656
    Reviewed-by: Keith Randall <khr@golang.org>

commit d211c2d3774d78173e004f0ffb1e2eae9ae19706
Author: Austin Clements <austin@google.com>
Date:   Thu Sep 22 17:02:22 2016 -0400

    runtime: implement getcallersp in Go
    
    This makes it possible to inline getcallersp. getcallersp is on the
    hot path of defers, so this slightly speeds up defer:
    
    name           old time/op  new time/op  delta
    Defer-4        78.3ns ± 2%  75.1ns ± 1%  -4.00%   (p=0.000 n=9+8)
    
    Updates #14939.
    
    Change-Id: Icc1cc4cd2f0a81fc4c8344432d0b2e783accacdd
    Reviewed-on: https://go-review.googlesource.com/29655
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit aaf4099a5cabfee52b1c481f2a30ee0dd02ef247
Author: Austin Clements <austin@google.com>
Date:   Fri Sep 23 11:47:24 2016 -0400

    runtime: update malloc.go documentation
    
    The big documentation comment at the top of malloc.go has gotten
    woefully out of date. Update it.
    
    Change-Id: Ibdb1bdcfdd707a6dc9db79d0633a36a28882301b
    Reviewed-on: https://go-review.googlesource.com/29731
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit f67c9de65621a46905a2f923b35634f74fa8635b
Author: Austin Clements <austin@google.com>
Date:   Fri Sep 9 21:43:50 2016 -0400

    runtime: document MemStats
    
    This documents all fields in MemStats and more clearly documents where
    mstats differs from MemStats.
    
    Fixes #15849.
    
    Change-Id: Ie09374bcdb3a5fdd2d25fe4bba836aaae92cb1dd
    Reviewed-on: https://go-review.googlesource.com/28972
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>

commit 2098e5d39afc3e6f687fcabd7704e8013f9c0dbd
Author: Austin Clements <austin@google.com>
Date:   Thu Sep 15 14:30:31 2016 -0400

    runtime: eliminate memstats.heap_reachable
    
    We used to compute an estimate of the reachable heap size that was
    different from the marked heap size. This ultimately caused more
    problems than it solved, so we pulled it out, but memstats still has
    both heap_reachable and heap_marked, and there are some leftover TODOs
    about the problems with this estimate.
    
    Clean this up by eliminating heap_reachable in favor of heap_marked
    and deleting the stale TODOs.
    
    Change-Id: I713bc20a7c90683d2b43ff63c0b21a440269cc4d
    Reviewed-on: https://go-review.googlesource.com/29271
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit ec9c84c8840124b4cf883ab08f6c2dc6baf20732
Author: Austin Clements <austin@google.com>
Date:   Thu Sep 15 14:08:04 2016 -0400

    runtime: disentangle next_gc from GC trigger
    
    Back in Go 1.4, memstats.next_gc was both the heap size at which GC
    would trigger, and the size GC kept the heap under. When we switched
    to concurrent GC in Go 1.5, we got somewhat confused and made this
    variable the trigger heap size, while gcController.heapGoal became the
    goal heap size.
    
    memstats.next_gc is exposed to the user via MemStats.NextGC, while
    gcController.heapGoal is not. This is unfortunate because 1) the heap
    goal is far more useful for diagnostics, and 2) the trigger heap size
    is just part of the GC trigger heuristic, which means it wouldn't be
    useful to an application even if it tried to use it.
    
    We never noticed this mess because MemStats.NextGC is practically
    undocumented. Now that we're trying to document MemStats, it became
    clear that this field had diverged from its original usefulness.
    
    Clean up this mess by shuffling things back around so that next_gc is
    the goal heap size and the new (unexposed) memstats.gc_trigger field
    is the trigger heap size. This eliminates gcController.heapGoal.
    
    Updates #15849.
    
    Change-Id: I2cbbd43b1d78bdf613cb43f53488bd63913189b7
    Reviewed-on: https://go-review.googlesource.com/29270
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 196df6f0c9ddd7043984cdf5d7becfb647ec1a31
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Sep 26 13:22:08 2016 -0700

    go/ast: better documentation for CallExpr node
    
    Fixes #17222.
    
    Change-Id: Iffffc8cbb8627d06afa9066246b68fa2da4600e3
    Reviewed-on: https://go-review.googlesource.com/29810
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 0a185610321ddbc68142ed8ed2ca3dc2b08578b4
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Sep 26 12:43:21 2016 -0700

    cmd/printer: document that Fprint doesn't match gofmt output
    
    Fixes #16963.
    
    Change-Id: Iaadf0da4ee9cc97146c5e6ac2d93de9ae6893880
    Reviewed-on: https://go-review.googlesource.com/29790
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e2e11f02a4627e4090083d433e6c66602b514ab7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Sep 26 11:35:55 2016 -0700

    runtime: unify Unix implementations of unminit
    
    Change-Id: I2cbb13eb85876ad05a52cbd498a9b86e7a28899c
    Reviewed-on: https://go-review.googlesource.com/29772
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6bcd258095da307dc60ffe1b45b014c0be849dc9
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sat Sep 24 15:53:09 2016 +0200

    math/big: better SetFloat64 example in doc
    
    Fixes #17221
    
    Change-Id: Idaa2af6b8646651ea72195671d1a4b5c370a5a22
    Reviewed-on: https://go-review.googlesource.com/29711
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit ac24388e5e5bdc129451c074a349a982e1d55ffa
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Sep 26 11:14:41 2016 -0700

    runtime: merge setting new signal mask in minit
    
    All the variants that sets the new signal mask in minit do the same
    thing, so merge them. This requires an OS-specific sigdelset function;
    the function already exists for linux, and is now added for other OS's.
    
    Change-Id: Ie96f6f02e2cf09c43005085985a078bd9581f670
    Reviewed-on: https://go-review.googlesource.com/29771
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3390294308462d5ddab95feaf42acc8fde767c98
Author: David Chase <drchase@google.com>
Date:   Mon Sep 26 10:06:10 2016 -0700

    cmd/compile: PPC64, find compare-with-immediate
    
    Added rules for compare double and word immediate,
    including those that use invertflags to cope with
    flipped operands.
    
    Change-Id: I594430a210e076e52299a2cc6ab074dbb04a02bd
    Reviewed-on: https://go-review.googlesource.com/29763
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>

commit 7f1bc53379f548845a23275c2f83d75071cf8e13
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sat Sep 10 13:24:35 2016 +0200

    cmd/compile: only allow integer expressions as keys in array literals
    
    Fixes #16439
    Updates #16679
    
    Change-Id: Idff4b313f29351866b1a649786501adee85fd580
    Reviewed-on: https://go-review.googlesource.com/29011
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit c2735039f3ea4e44a3c1df6ef6715e83bc5257b1
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun Sep 25 21:33:27 2016 -0700

    runtime: unify sigtrampgo
    
    Combine the various versions of sigtrampgo into a single function in
    signal_unix.go. This requires defining a fixsigcode method on sigctxt
    for all operating systems; it only does something on Darwin. This also
    requires changing the darwin/amd64 signal handler to call sigreturn
    itself, rather than relying on sigtrampgo to call sigreturn for it. We
    can then drop the Darwin sigreturn function, as it is no longer used.
    
    Change-Id: I5a0b9d2d2c141957e151b41e694efeb20e4b4b9a
    Reviewed-on: https://go-review.googlesource.com/29761
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1906d93bfda1c7b4af51457bc0a137b8418500b7
Author: David Chase <drchase@google.com>
Date:   Mon Sep 26 08:32:49 2016 -0700

    cmd/compile: On PPC, nilcheck load should be MOVBZ
    
    There's no load-signed-byte on PPC, so MOVB
    causes the assembler to macro-expand in a
    useless sign extension.
    
    Fixes #17211.
    
    Change-Id: Ibcd73aea4c94ba6df0a998b0091e45508113be2a
    Reviewed-on: https://go-review.googlesource.com/29762
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 375092bdcbc5b095e1591558952ce537b9fa5fa3
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sat May 7 21:34:04 2016 +0200

    cmd/compile: fix bogus "fallthrough statement out of place"
    
    When processing a fallthrough, the casebody function in swt.go
    checks that the last statement has indeed Op == OXFALL (not-processed
    fallthrough) before setting it to OFALL (processed fallthrough).
    
    Unfortunately, sometimes the fallthrough statement won't be in the
    last node. For example, in
    
    case 0:
             return func() int {return 1}()
             fallthrough
    
    the compiler generates
    
    autotmp_0 = (func literal)(); return autotmp_0; fallthrough; <node VARKILL>
    
    with an OVARKILL node in the last position. casebody will find that
    last.Op != OXFALL, won't mark the fallthrough as processed, and the
    fallthrough line will cause a "fallthrough statement out of place" error.
    
    To fix this, we change casebody so that it searches for the fallthrough
    statement backwards in the statements list, without assuming that it'll
    be in the last position.
    
    Fixes #13262
    
    Change-Id: I366c6caa7fd7442d365bd7a08cc66a552212d9b2
    Reviewed-on: https://go-review.googlesource.com/22921
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit d15295c6790b70eba0e4a3aa7ddead251aa440da
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun Sep 25 13:38:54 2016 -0700

    runtime: unify handling of alternate signal stack
    
    Change all Unix systems to use stackt for the alternate signal
    stack (some were using sigaltstackt). Add OS-specific setSignalstackSP
    function to handle different types for ss_sp field, and unify all
    OS-specific signalstack functions into one. Unify handling of alternate
    signal stack in OS-specific minit and sigtrampgo functions via new
    functions minitSignalstack and setGsignalStack.
    
    Change-Id: Idc316dc69b1dd725717acdf61a1cd8b9f33ed174
    Reviewed-on: https://go-review.googlesource.com/29757
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e6143e17d3e0c3ab8a7bd8357001217eb01dc6c6
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Sep 25 18:31:05 2016 -0700

    net/http: add Client tests for various 3xx redirect codes
    
    Updates #13994
    Updates #16840
    
    Change-Id: Ia3cad5c211e0c688a945ed6b6277c2552592774c
    Reviewed-on: https://go-review.googlesource.com/29760
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5bbb854cee2df329e031e50181ccc022c9d93a85
Author: Gyu-Ho Lee <gyuhox@gmail.com>
Date:   Sat Jun 4 23:26:12 2016 -0700

    net/http/httputil: preallocate trailerKeys slice
    
    To prevent slice growths with append operations.
    
    Change-Id: Icdb745b23cc44dfaf3e16746b94c06997f814e15
    Reviewed-on: https://go-review.googlesource.com/23784
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4383e4387b30ffbd8f85f053e399d53d7eef9330
Author: Antonio Murdaca <runcom@redhat.com>
Date:   Sun Jun 26 12:14:41 2016 +0200

    net/url: avoid if statement
    
    Change-Id: I894a8f49d29dbb6f9265e4b3df5767318b225460
    Signed-off-by: Antonio Murdaca <runcom@redhat.com>
    Reviewed-on: https://go-review.googlesource.com/24492
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5df7f5220f42cb283147ab271cd965720e233759
Author: Sam Whited <sam@samwhited.com>
Date:   Thu Sep 1 09:19:32 2016 -0500

    net/rpc: Add documentation for client.Close
    
    Fixes #16678
    
    Change-Id: I48c2825d4fef55a75d2f99640a7079c56fce39db
    Reviewed-on: https://go-review.googlesource.com/28370
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6ba5b32922c438a608a11561100a8a80abf0fd3a
Author: Anders Pearson <anders@columbia.edu>
Date:   Sun Jul 3 14:22:06 2016 +0100

    expvar: export http.Handler
    
    Add a method to expose the handler to allow it to be installed at a
    non-standard location or used with a different ServeMux.
    
    fixes #15030
    
    Change-Id: If778ad6fcc200f124a05c0a493511e364fca6078
    Reviewed-on: https://go-review.googlesource.com/24722
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f05cd4cde5f7023b1db23b7499cf9b1a6f3ffdec
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun Sep 25 12:37:00 2016 -0700

    runtime: simplify conditions testing g.paniconfault
    
    Implement a comment by Ralph Corderoy on CL 29754.
    
    Change-Id: I22bbede211ddcb8a057f16b4f47d335a156cc8d2
    Reviewed-on: https://go-review.googlesource.com/29756
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 343bec53c7317e1c780b3faf04aa124f19849b61
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sat Sep 24 21:17:34 2016 -0700

    runtime: merge sigpanic_unix.go into signal_unix.go
    
    Change-Id: Iba541045b4878405834c637095627631b6559a35
    Reviewed-on: https://go-review.googlesource.com/29754
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 38765eba739461e5c5dc463860c62daee2eef4ee
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Sat Sep 24 17:07:35 2016 +0200

    runtime/race: don't crash on invalid PCs
    
    Currently raceSymbolizeCode uses funcline, which is internal runtime
    function which crashes on incorrect PCs. Use FileLine instead,
    it is public and does not crash on invalid data.
    
    Note: FileLine returns "?" file on failure. That string is not NUL-terminated,
    so we need to additionally check what FileLine returns.
    
    Fixes #17190
    
    Change-Id: Ic6fbd4f0e68ddd52e9b2dd25e625b50adcb69a98
    Reviewed-on: https://go-review.googlesource.com/29714
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9f1c78781b320b6d7cf83378b857c1168cb7fd0f
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Sat Sep 24 17:03:30 2016 +0200

    cmd/cgo: fix line info in _cgo_gotypes.go
    
    Don't write line info for types, we don't have it.
    Otherwise types look like:
    
    type _Ctype_struct_cb struct {
    //line :1
          on_test *[0]byte
    //line :1
    }
    
    Which is not useful. Moreover we never override source info,
    so subsequent source code uses the same source info.
    Moreover, empty file name makes compile emit no source debug info at all.
    
    Update #17190
    
    Change-Id: I7ae6fa4964520d7665743d340419b787df0b51e8
    Reviewed-on: https://go-review.googlesource.com/29713
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c14050646f621db5f54c1a300b80c65f99fbd03b
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Sat Sep 24 16:59:23 2016 +0200

    runtime: fix newextram PC passed to race detector
    
    PC passed to racegostart is expected to be a return PC
    of the go statement. Race runtime will subtract 1 from the PC
    before symbolization. Passing start PC of a function is wrong.
    Add sys.PCQuantum to the function start PC.
    
    Update #17190
    
    Change-Id: Ia504c49e79af84ed4ea360c2aea472b370ea8bf5
    Reviewed-on: https://go-review.googlesource.com/29712
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 893edc9b3fceead9c31151ce46fb871ea8c225b1
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Sat Sep 24 22:06:39 2016 -0700

    net/http/httptrace: remove the mention of http client
    
    This sentence is partially guilty why httptrace is considered as an
    http.Client tracing package. Removing the mention.
    
    Updates #17152.
    
    Change-Id: I69f78a6e10817db933f44e464a949ae896e44ec6
    Reviewed-on: https://go-review.googlesource.com/29755
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 159a90b93a962cb942688f099b42d00d164e436f
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Sep 23 22:05:51 2016 -0700

    runtime: merge Unix sighandler functions
    
    Replace all the Unix sighandler functions with a single instance.
    Push the relatively small amount of processor-specific code into five
    methods on sigctxt: sigpc, sigsp, siglr, fault, preparePanic.
    (Some processors already had a fault method.)
    
    Change-Id: Ib459412ff8f7e0f5ad06bfd43eb827c8b196fc32
    Reviewed-on: https://go-review.googlesource.com/29752
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 890c09efb72f11b3e6fd95d859260bfee7de7f31
Author: Kale Blankenship <kale@lemnisys.com>
Date:   Sat Sep 24 16:39:36 2016 -0700

    os: make IsExist report true on ERROR_DIR_NOT_EMPTY on Windows
    
    Fixes #17164
    
    Change-Id: I3e626d92293c1379e2922276f033fdee6f48dda3
    Reviewed-on: https://go-review.googlesource.com/29753
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 60074b0fd35e438b06ba8e42c0e90b1a8a2945b7
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 23 19:57:16 2016 -0700

    runtime: remove TestCollisions from -short
    
    Takes a bit too long to run it all the time.
    
    Fixes #17217
    Update #17104
    
    Change-Id: I4802190ea16ee0f436a7f95b093ea0f995f5b11d
    Reviewed-on: https://go-review.googlesource.com/29751
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ab552aa3b69deb208b38677880e86aa41c3a9e47
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Sep 23 17:54:51 2016 -0700

    runtime: unify some signal handling functions
    
    Unify the OS-specific versions of msigsave, msigrestore, sigblock,
    updatesigmask, and unblocksig into single versions in signal_unix.go.
    To do this, make sigprocmask work the same way on all systems, which
    required adding a definition of sigprocmask for linux and openbsd.
    Also add a single OS-specific function sigmaskToSigset.
    
    Change-Id: I7cbf75131dddb57eeefe648ef845b0791404f785
    Reviewed-on: https://go-review.googlesource.com/29689
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit fd296282e0a5058351954f1a7ea2dac5ef87f052
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 19 14:14:58 2016 -0400

    cmd/dist: enable plugin test on darwin/amd64
    
    Change-Id: I6071881a5f7b9638bca0bfef76d6f4f45c9202a6
    Reviewed-on: https://go-review.googlesource.com/29396
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 96acfaaefae3378dcde009a184d3f4e5795ee5cc
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 19 14:14:45 2016 -0400

    cmd/go: enable -buildmode=plugin on darwin/amd64
    
    Change-Id: I8e594e059448879a9f451801064729186ac7c11b
    Reviewed-on: https://go-review.googlesource.com/29395
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 73e7a569b4b2917e76f9e1ad95f638127b7638ad
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 19 14:13:07 2016 -0400

    cmd/link: plugin support on darwin/amd64
    
    This CL turns some special section marker symbols into real symbols
    laid out in the sections they mark. This is to deal with the fact
    that dyld on OS X resolves the section marker symbols in any dlopen-ed
    Go program to the original section marker symbols in the host program.
    
    More details in a comment in cmd/link/internal/ld/data.go.
    
    Change-Id: Ie9451cfbf06d0bdcccb9959219c791b829f3f771
    Reviewed-on: https://go-review.googlesource.com/29394
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 26a6131bacb5dbc491f77329557580df0a310858
Author: Keith Randall <khr@golang.org>
Date:   Mon Sep 19 14:35:41 2016 -0700

    cmd/compile: fix 4-byte unaligned load rules
    
    The 2-byte rule was firing before the 4-byte rule, preventing
    the 4-byte rule from firing.  Update the 4-byte rule to use
    the results of the 2-byte rule instead.
    
    Add some tests to make sure we don't regress again.
    
    Fixes #17147
    
    Change-Id: Icfeccd9f2b96450981086a52edd76afb3191410a
    Reviewed-on: https://go-review.googlesource.com/29382
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 120e9ff34f577376f8d1c25a8966c88e0e92fee5
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Sep 23 10:38:50 2016 -0400

    cmd/compile: recognize OpS390XLoweredNilCheck as a nil check in the scheduler
    
    Before this change a nil check on s390x could be scheduled after the
    target pointer has been dereferenced.
    
    Change-Id: I7ea40a4b52f975739f6db183a2794be4981c4e3d
    Reviewed-on: https://go-review.googlesource.com/29730
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3311275ce8eef87a64b78589e0da0bf115f9be07
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Thu Sep 22 14:54:45 2016 -0500

    math, cmd/internal/obj/ppc64: improve floor, ceil, trunc with asm
    
    This adds the instructions frim, frip, and friz to the ppc64x
    assembler for use in implementing the math.Floor, math.Ceil, and
    math.Trunc functions to improve performance.
    
    Fixes #17185
    
    BenchmarkCeil-128                    21.4          6.99          -67.34%
    BenchmarkFloor-128                   13.9          6.37          -54.17%
    BenchmarkTrunc-128                   12.7          6.33          -50.16%
    
    Change-Id: I96131bd4e8c9c8dbafb25bfeb544cf9d2dbb4282
    Reviewed-on: https://go-review.googlesource.com/29654
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit ed915ad421e61c3158ccae4bfbbcfbc796514ee2
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 23 00:02:38 2016 -0400

    runtime: use sched_yield instead of pthread_yield
    
    Attempt to fix the linux-amd64-clang builder, which broke
    with CL 29472.
    
    Turns out pthread_yield is a non-portable Linux function, and
    should have #define _GNU_SOURCE before #include <pthread.h>.
    GCC doesn't complain about this, but Clang does:
    
            ./raceprof.go:44:3: warning: implicit declaration of function 'pthread_yield' is invalid in C99 [-Wimplicit-function-declaration]
    
    (Though the error, while explicable, certainly could be clearer.)
    
    There is a portable POSIX equivalent, sched_yield, so this
    CL uses it instead.
    
    Change-Id: I58ca7a3f73a2b3697712fdb02e72a8027c391169
    Reviewed-on: https://go-review.googlesource.com/29675
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b444d438c061fa934130fed17d34c23e77174851
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 19 14:09:07 2016 -0400

    plugin: darwin support
    
    Change-Id: I76981d1d83da401178226634d076371a04f5ccb7
    Reviewed-on: https://go-review.googlesource.com/29392
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b4c9829c2299772eb0aaffcd437085f9f43e7166
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 19 14:08:21 2016 -0400

    runtime: check plugin-loaded moduledata addresses
    
    Inspired by difficulties with plugin support on darwin.
    
    Change-Id: I2cef8410837946454e75d00e94e46791f03f2267
    Reviewed-on: https://go-review.googlesource.com/29391
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9d8522fdc72ecc8eaa2d318a2cc04abde9beeb42
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Sep 21 09:44:40 2016 -0700

    cmd/compile: don't instrument copy and append in runtime
    
    Instrumenting copy and append for the race detector changes them to call
    different functions. In the runtime package the alternate functions are
    not marked as nosplit. This caused a crash in the SIGPROF handler when
    invoked on a non-Go thread in a program built with the race detector. In
    some cases the handler can call copy, the race detector changed that to
    a call to a non-nosplit function, the function tried to check the stack
    guard, and crashed because it was running on a non-Go thread. The
    SIGPROF handler is written carefully to avoid such problems, but hidden
    function calls are difficult to avoid.
    
    Fix this by changing the compiler to not instrument copy and append when
    compiling the runtime package. Change the runtime package to add
    explicit race checks for the only code I could find where copy is used
    to write to user data (append is never used).
    
    Change-Id: I11078a66c0aaa459a7d2b827b49f4147922050af
    Reviewed-on: https://go-review.googlesource.com/29472
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 88d2f9112a0a8afd3a29ac1479d4f17847f16803
Author: Keith Randall <khr@golang.org>
Date:   Tue Sep 20 16:34:30 2016 -0700

    cmd/compile: fix type of static closure pointer
    
      var x *X = ...
      defer x.foo()
    
    As part of the defer, we need to calculate &(*X).foo·f.  This expression
    is the address of the static closure that will call (*X).foo when a
    pointer to that closure is used in a call/defer/go.  This pointer is not
    currently properly typed in SSA.  It is a pointer type, but the base
    type is nil, not a proper type.
    
    This turns out not to be a problem currently because we never use the
    type of these SSA values.  But I'm trying to change that (to be able to
    spill them) in CL 28391.  To fix, use uint8 as the fake type of the
    closure.
    
    Change-Id: Ieee388089c9af398ed772ee8c815122c347cb633
    Reviewed-on: https://go-review.googlesource.com/29444
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d586aae1f44ebdf0e8f92137856b4b62c41cac6a
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Sep 22 13:50:16 2016 -0400

    test: errorcheck auto-generated functions
    
    Add an "errorcheckwithauto" action which performs error check
    including lines with auto-generated functions (excluded by
    default). Comment "// ERRORAUTO" matches these lines.
    
    Add testcase for CL 29570 (as an example).
    
    Updates #16016, #17186.
    
    Change-Id: Iaba3727336cd602f3dda6b9e5f97dafe0848e632
    Reviewed-on: https://go-review.googlesource.com/29652
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 3dfb92f254ed4f94e3c98a789c171a1cd9c2563d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Sep 21 18:53:31 2016 -0400

    cmd/compile: ensure args are live in tail calls for LR machines
    
    On link-register machines we uses RET (sym), instead of JMP (sym),
    for tail call (so the assembler knows and may rewrite it to
    restore link register if necessary). Add RET to the analysis.
    
    Fixes #17186.
    Fixes #16016 on link-register machines.
    
    Change-Id: I8690ac57dd9d49beeea76a5f291988e9a1d3afe5
    Reviewed-on: https://go-review.googlesource.com/29570
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 254169d7bb6f96409ba4b405c80d28deefdb000b
Author: Adam Langley <agl@golang.org>
Date:   Wed Sep 14 11:50:36 2016 -0700

    crypto/tls: fix deadlock when racing to complete handshake.
    
    After renegotiation support was added (af125a5193c) it's possible for a
    Write to block on a Read when racing to complete the handshake:
       1. The Write determines that a handshake is needed and tries to
          take the neccesary locks in the correct order.
       2. The Read also determines that a handshake is needed and wins
          the race to take the locks.
       3. The Read goroutine completes the handshake and wins a race
          to unlock and relock c.in, which it'll hold when waiting for
          more network data.
    
    If the application-level protocol requires the Write to complete before
    data can be read then the system as a whole will deadlock.
    
    Unfortunately it doesn't appear possible to reverse the locking order of
    c.in and handshakeMutex because we might read a renegotiation request at
    any point and need to be able to do a handshake without unlocking.
    
    So this change adds a sync.Cond that indicates that a goroutine has
    committed to doing a handshake. Other interested goroutines can wait on
    that Cond when needed.
    
    The test for this isn't great. I was able to reproduce the deadlock with
    it only when building with -race. (Because -race happened to alter the
    timing just enough.)
    
    Fixes #17101.
    
    Change-Id: I4e8757f7b82a84e46c9963a977d089f0fb675495
    Reviewed-on: https://go-review.googlesource.com/29164
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ad5d91c17a3c0bc4acf9e4036b050517972432f0
Author: Kale Blankenship <kale@lemnisys.com>
Date:   Wed Sep 21 19:03:06 2016 -0700

    net/url: prefix relative paths containing ":" in the first segment with "./"
    
    This change modifies URL.String to prepend "./" to a relative URL which
    contains a colon in the first path segment.
    
    Per RFC 3986 §4.2:
    
    > A path segment that contains a colon character (e.g., "this:that")
    > cannot be used as the first segment of a relative-path reference, as
    > it would be mistaken for a scheme name.  Such a segment must be
    > preceded by a dot-segment (e.g., "./this:that") to make a relative-
    > path reference.
    
    https://go-review.googlesource.com/27440 corrects the behavior for http.FileServer,
    but URL.String will still return an invalid URL. This CL reverts the changes to
    http.FileServer as they are unnecessary with this fix.
    
    Fixes #17184
    
    Change-Id: I9211ae20f82c91b785d1b079b2cd766487d94225
    Reviewed-on: https://go-review.googlesource.com/29610
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cddddbc6231177c87b72f95209ab51abb74bcbc5
Author: David Chase <drchase@google.com>
Date:   Fri Sep 16 15:02:47 2016 -0700

    cmd/compile: use ISEL, cleanup use of zero & extensions
    
    Abandoned earlier efforts to expose zero register,
    but left it in numbering to decrease squirrelyness of
    register allocator.
    
    ISELrelOp used in code generation of bool := x relOp y.
    Some patterns added to better elide zero case and
    some sign extension.
    
    Updates: #17109
    
    Change-Id: Ida7839f0023ca8f0ffddc0545f0ac269e65b05d9
    Reviewed-on: https://go-review.googlesource.com/29380
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
```
