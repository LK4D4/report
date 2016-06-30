# June 30, 2016 Report

Number of commits: 41

## Compilation time

* github.com/coreos/etcd/cmd: from 15.154866207s to 15.272458147s, +0.78%
* github.com/boltdb/bolt/cmd/bolt: from 555.130635ms to 565.195687ms, +1.81%
* github.com/gogits/gogs: from 13.042114125s to 12.921796999s, -0.92%
* github.com/spf13/hugo: from 8.913891062s to 7.896182241s, -11.42%
* github.com/nsqio/nsq/apps/nsqd: from 2.009413476s to 2.007659682s, ~
* github.com/mholt/caddy: from 283.861781ms to 281.226338ms, -0.93%

## Binary size:

* github.com/coreos/etcd/cmd: from 26370221 to 26383277, ~
* github.com/boltdb/bolt/cmd/bolt: from 2665904 to 2666587, ~
* github.com/gogits/gogs: from 23612374 to 23625550, ~
* github.com/spf13/hugo: from 15172129 to 15189680, +0.12%
* github.com/nsqio/nsq/apps/nsqd: from 10033736 to 10046829, +0.13%
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               190           191           +0.53%
BenchmarkMsgpUnmarshal-4             410           411           +0.24%
BenchmarkEasyJsonMarshal-4           1589          1588          -0.06%
BenchmarkEasyJsonUnmarshal-4         1543          1544          +0.06%
BenchmarkFlatBuffersMarshal-4        369           414           +12.20%
BenchmarkFlatBuffersUnmarshal-4      292           292           +0.00%
BenchmarkGogoprotobufMarshal-4       165           162           -1.82%
BenchmarkGogoprotobufUnmarshal-4     259           255           -1.54%

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

* [cmd/go: restore support for git submodules and update docs](https://github.com/golang/go/commit/069289180816e2f8b40ad6f9e167dc5071cefcdf)


## GIT Log

```
commit e0c8af090ea1ccc32d06ae75b653446d2a9d6f87
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Jun 29 17:01:10 2016 -0700

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 8e573f40 for https://golang.org/cl/24600,
    "http2: merge multiple GOAWAY frames' contents into error message"
    
    Fixes #14627 (more)
    
    Change-Id: I5231607c2c9e0d854ad6199ded43c59e59f62f52
    Reviewed-on: https://go-review.googlesource.com/24612
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 51b08d511e8b42eace59588a7eea73c4d21d222d
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Jun 29 16:56:43 2016 -0700

    net/http: be consistent about spelling of HTTP/1.x
    
    There was only one use of "HTTP/1.n" compared to "HTTP/1.x":
    
    h2_bundle.go://   "Just as in HTTP/1.x, header field names are strings of ASCII
    httputil/dump.go:// DumpRequest returns the given request in its HTTP/1.x wire
    httputil/dump.go:// intact. HTTP/2 requests are dumped in HTTP/1.x form, not in their
    response.go:// Write writes r to w in the HTTP/1.x server response format,
    server.go:      // Request.Body. For HTTP/1.x requests, handlers should read any
    server.go:// The default HTTP/1.x and HTTP/2 ResponseWriter implementations
    server.go:// The default ResponseWriter for HTTP/1.x connections supports
    server.go:// http1ServerSupportsRequest reports whether Go's HTTP/1.x server
    server.go:      // about HTTP/1.x Handlers concurrently reading and writing, like
    server.go:      // HTTP/1.x from here on.
    transport.go:   return fmt.Errorf("net/http: HTTP/1.x transport connection broken: %v", err)
    
    Be consistent.
    
    Change-Id: I93c4c873e500f51af2b4762055e22f5487a625ac
    Reviewed-on: https://go-review.googlesource.com/24610
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit cc6f5f6ce1daca51f475764a3e4fa9420cafcefa
Author: Nick Harper <nharper@google.com>
Date:   Wed Jun 29 11:18:01 2016 -0700

    crypto/ecdsa: Update documentation for Sign
    
    Change-Id: I2b7a81cb809d109f10d5f0db957c614f466d6bfd
    Reviewed-on: https://go-review.googlesource.com/24582
    Reviewed-by: Adam Langley <agl@golang.org>

commit ad82f2cf4b7e8e5f5398b5546b7d834432347355
Author: Tom Bergan <tombergan@google.com>
Date:   Wed Jun 29 07:45:23 2016 -0700

    crypto/tls: Use the same buffer size in the client and server in the TLS throughput benchmark
    
    I believe it's necessary to use a buffer size smaller than 64KB because
    (at least some versions of) Window using a TCP receive window less than
    64KB. Currently the client and server use buffer sizes of 16KB and 32KB,
    respectively (the server uses io.Copy, which defaults to 32KB internally).
    Since the server has been using 32KB, it should be safe for the client to
    do so as well.
    
    Fixes #15899
    
    Change-Id: I36d44b29f2a5022c03fc086213d3c1adf153e983
    Reviewed-on: https://go-review.googlesource.com/24581
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bb337372fb6e171a6e8a7665ce91eda734f8cdd2
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Wed Jun 29 11:09:36 2016 +0200

    runtime: fix race atomic operations on external memory
    
    The assembly is broken: it does `MOVQ g(R12), R14` expecting that
    R12 contains tls address, but it does not do get_tls(R12) before.
    This magically works on linux: `MOVQ g(R12), R14` is compiled to
    `mov %fs:0xfffffffffffffff8,%r14` which does not use R12.
    But it crashes on windows.
    
    Add explicit `get_tls(R12)`.
    
    Fixes #16206
    
    Change-Id: Ic1f21a6fef2473bcf9147de6646929781c9c1e98
    Reviewed-on: https://go-review.googlesource.com/24590
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 25a609556aff7700c864f2dc69be01652fd801ab
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 28 17:23:01 2016 -0700

    runtime: correct printing of blocked field in scheduler trace
    
    When the blocked field was first introduced back in
    https://golang.org/cl/61250043 the scheduler trace code incorrectly used
    m->blocked instead of mp->blocked.  That has carried through the
    conversion to Go.  This CL fixes it.
    
    Change-Id: Id81907b625221895aa5c85b9853f7c185efd8f4b
    Reviewed-on: https://go-review.googlesource.com/24571
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c7ae41e5770b2258074eee68a6a3c4d0d71a251f
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 28 17:06:59 2016 -0700

    runtime: better error message for newosproc failure
    
    If creating a new thread fails with EAGAIN, point the user at ulimit.
    
    Fixes #15476.
    
    Change-Id: Ib36519614b5c72776ea7f218a0c62df1dd91a8ea
    Reviewed-on: https://go-review.googlesource.com/24570
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8641e6fe2131ac342647fa34398a727f96d15fb5
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Jun 28 16:25:45 2016 -0700

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev ef2e00e88 for https://golang.org/cl/24560,
    "http2: make Transport return server's GOAWAY error back to the user"
    
    Fixes #14627
    
    Change-Id: I2bb123a3041e168db7c9446beef4ee47638f17ee
    Reviewed-on: https://go-review.googlesource.com/24561
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 85a4f44745859ecdd71c034171d40263651f8594
Author: Konstantin Shaposhnikov <k.shaposhnikov@gmail.com>
Date:   Mon Jun 27 17:13:15 2016 +0800

    cmd/vet: make checking example names in _test packages more robust
    
    Prior to this change package "foo" had to be installed in order to check
    example names in "foo_test" package.
    
    However by the time "foo_test" package is checked a parsed "foo" package
    has been already constructed. Use it to check example names.
    
    Also change TestDivergentPackagesExamples test to pass directory of the
    package to the vet tool as it is the most common way to invoke it. This
    requires changes to errchk to add support for grabbing source files from
    a directory.
    
    Fixes #16189
    
    Change-Id: Ief103d07b024822282b86c24250835cc591793e8
    Reviewed-on: https://go-review.googlesource.com/24488
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 733aefd06e5cf708637308a4ad7a048aa97db5cd
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Jun 28 12:06:08 2016 -0700

    database/sql: deflake TestPendingConnsAfterErr and fix races, panics
    
    TestPendingConnsAfterErr only cared that things didn't deadlock, so 5
    seconds is a sufficient timer. We don't need 100 milliseconds.
    
    I was able to reproduce with a tiny (5 nanosecond) timeout value,
    instead of 100 milliseconds. In the process of testing with -race and
    a high -count= value, I noticed several data races and panics
    (sendings on a closed channel) which are also fixed in this change.
    
    Fixes #15684
    
    Change-Id: Ib4605fcc0f296e658cb948352ed642b801cb578c
    Reviewed-on: https://go-review.googlesource.com/24550
    Reviewed-by: Marko Tiikkaja <marko@joh.to>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 069289180816e2f8b40ad6f9e167dc5071cefcdf
Author: Andrew Gerrand <adg@golang.org>
Date:   Tue Jun 28 18:09:56 2016 +1000

    cmd/go: restore support for git submodules and update docs
    
    Fixes #16165
    
    Change-Id: Ic90e5873e0c8ee044f09543177192dcae1dcdbed
    Reviewed-on: https://go-review.googlesource.com/24531
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b0838ca292f0c62ac9d45a92b520160ed052cb26
Author: Justyn Temme <justyntemme@gmail.com>
Date:   Sun Jun 19 20:39:58 2016 +0000

    strconv: clarify doc for Atoi return type
    
    Change-Id: I47bd98509663d75b0d4dedbdb778e803d90053cf
    Reviewed-on: https://go-review.googlesource.com/24216
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b5f0aff49503e31002b33198e06708e263c445a7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Jun 27 16:39:40 2016 -0700

    net/http: conditionally configure HTTP/2 in Server.Serve(Listener)
    
    Don't configure HTTP/2 in http.Server.Serve(net.Listener) if the
    Server's TLSConfig is set and doesn't include the "h2" NextProto
    value. This avoids mutating a *tls.Config already in use if
    previously passed to tls.NewListener.
    
    Also document this. (it's come up a few times now)
    
    Fixes #15908
    
    Change-Id: I283eed82fdb29a791f80d801aadd9f75db244de0
    Reviewed-on: https://go-review.googlesource.com/24508
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 996ed3be9a10ace6cd7a8a6a8080c0c8db7ab1fe
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Tue Jun 28 17:47:27 2016 +0200

    doc: update 1.7 release notes on Unicode upgrade
    
    Fixes #16201
    
    Change-Id: I38c17859db78c2868905da24217e0ad47739c320
    Reviewed-on: https://go-review.googlesource.com/24541
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3001334e7752af67036a6d9c30c919cbfea81a7c
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Mon Jun 27 23:27:54 2016 -0700

    doc/go1.7.html: mention recent changes to Rand.Read
    
    Updates #16124
    
    Change-Id: Ib58f2bb37fd1559efc512a2e3cba976f09b939a0
    Reviewed-on: https://go-review.googlesource.com/24520
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit a2a4db7375204c22898559d3b96a748c6df96294
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Tue Jun 28 12:31:02 2016 +0200

    unicode: upgrade to version 9.0.0
    
    Changes beyond generated tables:
    - Now supports aliases to handle deprecated
      property classes.
    - Some Mongolian letters are now modifiers.
    
    Other changes:
    - strconv: newly generated table to be in sync
    - regexp/syntax: updated maxFold
    
    Fixes #16191
    
    Change-Id: I56bdf21ee2f775f2a82d0465b3772faf5c24cb61
    Reviewed-on: https://go-review.googlesource.com/24496
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ed9362f769626b1cdaf2eb1da63d5f25cadc979b
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Jun 27 21:37:19 2016 -0400

    reflect, runtime: optimize Name method
    
    Several minor changes that remove a good chunk of the overhead added
    to the reflect Name method over the 1.7 cycle, as seen from the
    non-SSA architectures.
    
    In particular, there are ~20 fewer instructions in reflect.name.name
    on 386, and the method now qualifies for inlining.
    
    The simple JSON decoding benchmark on darwin/386:
    
            name           old time/op    new time/op    delta
            CodeDecoder-8    49.2ms ± 0%    48.9ms ± 1%  -0.77%  (p=0.000 n=10+9)
    
            name           old speed      new speed      delta
            CodeDecoder-8  39.4MB/s ± 0%  39.7MB/s ± 1%  +0.77%  (p=0.000 n=10+9)
    
    On darwin/amd64 the effect is less pronounced:
    
            name           old time/op    new time/op    delta
            CodeDecoder-8    38.9ms ± 0%    38.7ms ± 1%  -0.38%  (p=0.005 n=10+10)
    
            name           old speed      new speed      delta
            CodeDecoder-8  49.9MB/s ± 0%  50.1MB/s ± 1%  +0.38%  (p=0.006 n=10+10)
    
    Counterintuitively, I get much more useful benchmark data out of my
    MacBook Pro than a linux workstation with more expensive Intel chips.
    While the laptop has fewer cores and an active GUI, the single-threaded
    performance is significantly better (nearly 1.5x decoding throughput)
    so the differences are more pronounced.
    
    For #16117.
    
    Change-Id: I4e0cc1cc2d271d47d5127b1ee1ca926faf34cabf
    Reviewed-on: https://go-review.googlesource.com/24510
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b75b0630fe6d8f66f937f78f15f540b5b6dab24f
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Thu Jun 16 14:57:45 2016 -0500

    runtime/internal/atomic: Use power5 compatible instructions for ppc64
    
    This modifies a recent performance improvement to the
    And8 and Or8 atomic functions which required both ppc64le
    and ppc64 to use power8 instructions. Since then it was
    decided that ppc64 (BE) should work for power5 and later.
    This change uses instructions compatible with power5 for
    ppc64 and uses power8 for ppc64le.
    
    Fixes #16004
    
    Change-Id: I623c75e8e6fd1fa063a53d250d86cdc9d0890dc7
    Reviewed-on: https://go-review.googlesource.com/24181
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 05ecf534566c0cd05b4afdaa9619522e4204328e
Author: Raul Silvera <rsilvera@google.com>
Date:   Mon Jun 27 19:06:17 2016 -0700

    net/http/pprof: remove comments pointing to gperftools
    
    The version of pprof in gperftools has been deprecated.
    No need to have a pointer to that version since go tool pprof
    is included with the Go distro.
    
    Change-Id: I6d769a68f64280f5db89ff6fbc67bfea9c8f1526
    Reviewed-on: https://go-review.googlesource.com/24509
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 73516c5f481250db4ccbfdddb8f68ef261897fcb
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Jun 27 11:07:08 2016 -0400

    encoding/gob: avoid allocating string for map key
    
    On linux/386 compared to tip:
    
            name                     old time/op  new time/op  delta
            DecodeInterfaceSlice-40  1.23ms ± 1%  1.17ms ± 1%  -4.93%  (p=0.000 n=9+10)
    
    Recovers about half the performance regression from Go 1.6 on 386.
    
    For #16117.
    
    Change-Id: Ie8676d92a4da3e27ff21b91a98b3e13d16730ba1
    Reviewed-on: https://go-review.googlesource.com/24468
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8d966bad6e4e9e28295bccbd739bf5280f729a7b
Author: Dmitri Popov <operator@cv.dp-net.com>
Date:   Sun Jun 19 20:58:40 2016 -0700

    math/rand: fix io.Reader implementation
    
    Do not throw away the rest of Int63 value used for
    generation random bytes. Save it in Rand struct and
    re-use during the next Read call.
    
    Fixes #16124
    
    Change-Id: Ic70bd80c3c3a6590e60ac615e8b3c2324589bea3
    Reviewed-on: https://go-review.googlesource.com/24251
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0ce100dc9646eb9015e31502e184e45605df1393
Author: Vladimir Mihailenco <vladimir.webdev@gmail.com>
Date:   Thu Jun 23 07:42:22 2016 +0000

    compress/flate: don't ignore dict in Reader.Reset
    
    Fixes #16162.
    
    Change-Id: I6f4ae906630079ef5fc29ee5f70e2e3d1c962170
    Reviewed-on: https://go-review.googlesource.com/24390
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit db5802104797cadcb4f44c5198a0fc39e13f9bc3
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 21 07:00:41 2016 -0700

    crypto/tls: don't copy Mutex or Once values
    
    This fixes some 40 warnings from go vet.
    
    Fixes #16134.
    
    Change-Id: Ib9fcba275fe692f027a2a07b581c8cf503b11087
    Reviewed-on: https://go-review.googlesource.com/24287
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit b43fe463ff93e5ab0849d9d31bf53f87b42a9b28
Author: Konstantin Shaposhnikov <k.shaposhnikov@gmail.com>
Date:   Sat Jun 25 20:32:40 2016 +0800

    net/http/httptest: show usage of httptest.NewRequest in example
    
    Change ExampleResponseRecorder to use httptest.NewRequest instead of
    http.NewRequest. This makes the example shorter and shows how to use
    one more function from the httptest package.
    
    Change-Id: I3d35869bd0a4daf1c7551b649428bb2f2a45eba2
    Reviewed-on: https://go-review.googlesource.com/24480
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3e9d6e064d554edeed9c55325832844403ae5d3f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Jun 27 10:28:08 2016 -0700

    net/http: reject faux HTTP/0.9 and HTTP/2+ requests
    
    Fixes #16197
    
    Change-Id: Icaabacbb22bc18c52b9e04b47385ac5325fcccd1
    Reviewed-on: https://go-review.googlesource.com/24505
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e0f986bf26b72749553e45cad34f14d7d4166acb
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jun 27 08:21:55 2016 -0700

    cmd/compile: avoid function literal name collision with "glob"
    
    The compiler was treating all global function literals as occurring in a
    function named "glob", which caused a symbol name collision when there
    was an actual function named "glob".  Fixed by adding a period.
    
    Fixes #16193.
    
    Change-Id: I67792901a8ca04635ba41d172bfaee99944f594d
    Reviewed-on: https://go-review.googlesource.com/24500
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit c0e5d445066a4c6417aa9e54fde4a3cb9d9c2c05
Author: Raul Silvera <rsilvera@google.com>
Date:   Mon Jun 27 09:43:14 2016 -0700

    runtime/pprof: update comments to point to new pprof
    
    In the comments for this file there is a reference to gperftools
    for more info on pprof. pprof now live on its own repo on github,
    and the version in gperftools is deprecated.
    
    Change-Id: I8a188f129534f73edd132ef4e5a2d566e69df7e9
    Reviewed-on: https://go-review.googlesource.com/24502
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6effdd28deec546798c00fd83c2a7b5b328ac5ab
Author: Keith Randall <khr@golang.org>
Date:   Sat Jun 18 19:40:57 2016 -0700

    cmd/compile: keep heap pointer for escaping output parameters live
    
    Make sure the pointer to the heap copy of an output parameter is kept
    live throughout the function.  The function could panic at any point,
    and then a defer could recover.  Thus, we need the pointer to the heap
    copy always available so the post-deferreturn code can copy the return
    value back to the stack.
    
    Before this CL, the pointer to the heap copy could be considered dead in
    certain situations, like code which is reverse dominated by a panic call.
    
    Fixes #16095.
    
    Change-Id: Ic3800423e563670e5b567b473bf4c84cddb49a4c
    Reviewed-on: https://go-review.googlesource.com/24213
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit e96b1ef99bdd18a6f777892008f614a4401d6655
Author: Konstantin Shaposhnikov <k.shaposhnikov@gmail.com>
Date:   Mon Jun 27 13:37:01 2016 +0800

    cmd/vet: fix name check for examples in _test package
    
    This fixes the obvious bug and makes go vet look for identifiers in foo
    package when checking example names in foo_test package.
    
    Note that for this check to work the foo package have to be
    installed (using go install).
    
    This commit however doesn't fix TestDivergentPackagesExamples test that
    is not implemented correctly and passes only by chance.
    
    Updates #16189
    
    Change-Id: I5c2f675cd07e5b66cf0432b2b3e422ab45c3dedd
    Reviewed-on: https://go-review.googlesource.com/24487
    Reviewed-by: Dmitri Shuralyov <shurcool@gmail.com>
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 5f209aba6d903688fd5d801bad3fbb5572c85a02
Author: David Crawshaw <crawshaw@golang.org>
Date:   Sat Jun 25 10:23:30 2016 -0400

    encoding/json: copy-on-write cacheTypeFields
    
    Swtich from a sync.RWMutex to atomic.Value for cacheTypeFields.
    
    On GOARCH=386, this recovers most of the remaining performance
    difference from the 1.6 release. Compared with tip on linux/386:
    
            name            old time/op    new time/op    delta
            CodeDecoder-40    92.8ms ± 1%    87.7ms ± 1%  -5.50%  (p=0.000 n=10+10)
    
            name            old speed      new speed      delta
            CodeDecoder-40  20.9MB/s ± 1%  22.1MB/s ± 1%  +5.83%  (p=0.000 n=10+10)
    
    With more time and care, I believe more of the JSON decoder's work
    could be shifted so it is done before decoding, and independent of
    the number of bytes processed. Maybe someone could explore that for
    Go 1.8.
    
    For #16117.
    
    Change-Id: I049655b2e5b76384a0d5f4b90e3ec7cc8d8c4340
    Reviewed-on: https://go-review.googlesource.com/24472
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 33fa855e6cfa906e12c624663b1010862bf4df6f
Author: Konstantin Shaposhnikov <k.shaposhnikov@gmail.com>
Date:   Thu May 26 09:45:57 2016 +0800

    math/rand: fix comment about bits of seed used by the default Source
    
    Fixes #15788
    
    Change-Id: I5a1fd1e5992f1c16cf8d8437d742bf02e1653b9c
    Reviewed-on: https://go-review.googlesource.com/23461
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4fcb4eb27986e5fa7d26aeb85b04909e03d11ff9
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jun 24 16:59:39 2016 -0700

    cmd/pprof: don't use local symbolization for remote source
    
    If we are using a remote source (a URL), and the user did not specify
    the executable file to use, then don't try to use a local source.
    This was misbehaving because the local symbolizer will not fail
    if there is any memory map available, but the presence of a memory map
    does not ensure that the files and symbols are actually available.
    
    We still need a pprof testsuite.
    
    Fixes #16159.
    
    Change-Id: I0250082a4d5181c7babc7eeec6bc95b2f3bcaec9
    Reviewed-on: https://go-review.googlesource.com/24464
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 83e839f86fcc0762ea513f8d5dd9a50a8338e9c2
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jun 24 13:19:46 2016 -0700

    cmd/pprof: ignore symbols with address 0 and size 0
    
    Handling a symbol with address 0 and size 0, such as an ELF STT_FILE
    symbols, was causing us to disassemble the entire program.  We started
    adding STT_FILE symbols to help fix issue #13247.
    
    Fixes #16154.
    
    Change-Id: I174b9614e66ddc3d65801f7c1af7650f291ac2af
    Reviewed-on: https://go-review.googlesource.com/24460
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 5e43dc943a9265ec65690242eb8076727c18a958
Author: Nathan VanBenschoten <nvanbenschoten@gmail.com>
Date:   Thu Jun 23 19:46:13 2016 -0400

    math/big: special-case a 0 mantissa during Rat parsing
    
    Previously, a 0 mantissa was special-cased during big.Float
    parsing, but not during big.Rat parsing. This meant that a value
    like 0e9999999999 would parse successfully in big.Float.SetString,
    but would hang in big.Rat.SetString. This discrepancy became an
    issue in https://golang.org/src/go/constant/value.go?#L250,
    where the big.Float would report an exponent of 0, so
    big.Rat.SetString would be used and would subsequently hang.
    
    A Go Playground example of this is https://play.golang.org/p/3fy28eUJuF
    
    The solution is to special-case a zero mantissa during big.Rat
    parsing as well, so that neither big.Rat nor big.Float will hang when
    parsing a value with 0 mantissa but a large exponent.
    
    This was discovered using go-fuzz on CockroachDB:
    https://github.com/cockroachdb/go-fuzz/blob/master/examples/parser/main.go
    
    Fixes #16176
    
    Change-Id: I775558a8682adbeba1cc9d20ba10f8ed26259c56
    Reviewed-on: https://go-review.googlesource.com/24430
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 797dc584577c66ee1e181a3f423133ee83647247
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Jun 24 15:28:58 2016 -0400

    cmd/compile, etc: use tflag to optimize Name()==""
    
    Improves JSON decoding benchmark:
    
            name                  old time/op    new time/op    delta
            CodeDecoder-8           41.3ms ± 6%    39.8ms ± 1%  -3.61%  (p=0.000 n=10+10)
    
            name                  old speed      new speed      delta
            CodeDecoder-8         47.0MB/s ± 6%  48.7MB/s ± 1%  +3.66%  (p=0.000 n=10+10)
    
    Change-Id: I524ee05c432fad5252e79b29222ec635c1dee4b4
    Reviewed-on: https://go-review.googlesource.com/24452
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2834526fd9de26079bebc726d3ce3ccaaf38a0aa
Author: Rob Pike <r@golang.org>
Date:   Thu Jun 23 15:28:36 2016 -0700

    time: update documentation for Duration.String regarding the zero value
    
    It was out of date; in 1.7 the format changes to 0s.
    
    Change-Id: I2013a1b0951afc5607828f313641b51c74433257
    Reviewed-on: https://go-review.googlesource.com/24421
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e4dc7f1ed2f90b2f72b21f5f8e00287e3c409876
Author: Sameer Ajmani <sameer@golang.org>
Date:   Fri Jun 24 10:48:06 2016 -0400

    context: update documentation on cancelation and go vet check.
    
    Also replace double spaces after periods with single spaces.
    
    Change-Id: Iedaea47595c5ce64e7e8aa3a368f36d49061c555
    Reviewed-on: https://go-review.googlesource.com/24431
    Reviewed-by: Alan Donovan <adonovan@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3c6ed76da2feb45f8fba9177e9d6e0f19671353b
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Jun 24 11:09:48 2016 -0400

    reflect: avoid lock for some NumMethod()==0 cases
    
    The encoding/json package uses NumMethod()==0 as a fast check for
    interface satisfaction. In the case when a type has no methods at
    all, we don't need to grab the RWMutex.
    
    Improves JSON decoding benchmark on linux/amd64:
    
            name           old time/op    new time/op    delta
            CodeDecoder-8    44.2ms ± 2%    40.6ms ± 1%  -8.11%  (p=0.000 n=10+10)
    
            name           old speed      new speed      delta
            CodeDecoder-8  43.9MB/s ± 2%  47.8MB/s ± 1%  +8.82%  (p=0.000 n=10+10)
    
    For #16117
    
    Change-Id: Id717e7fcd2f41b7d51d50c26ac167af45bae3747
    Reviewed-on: https://go-review.googlesource.com/24433
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f9d6b909b10e80d862932a935907e36c8cc24eeb
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Jun 23 20:51:39 2016 +0000

    A+C: automated updates
    
    Add Aaron Zinman (corporate CLA for Empirical Interfaces Inc.)
    Add Ayanamist Yang (individual CLA)
    Add Christian Couder (individual CLA)
    Add Eric Engestrom (individual CLA)
    Add Filippo Valsorda (individual CLA)
    Add Gyu-Ho Lee (individual CLA)
    Add H. İbrahim Güngör (individual CLA)
    Add Jacob Hoffman-Andrews (individual CLA)
    Add Jason Barnett (individual CLA)
    Add Joe Farrell (individual CLA)
    Add Julian Kornberger (individual CLA)
    Add Kris Rousey (corporate CLA for Google Inc.)
    Add Miguel Mendez (individual CLA)
    Add Nic Day (individual CLA)
    Add Paulo Casaretto (individual CLA)
    Add Philip Børgesen (individual CLA)
    Add Quan Tran (individual CLA)
    Add Sai Cheemalapati (corporate CLA for Google Inc.)
    Add Sasha Sobol (individual CLA)
    Add Seth Vargo (individual CLA)
    Add Simon Thulbourn (individual CLA)
    Add Wisdom Omuya (individual CLA)
    
    Updates #12042
    
    Change-Id: Ie8ab5e3500ee62000c0b176d4d71340446e72ab7
    Reviewed-on: https://go-review.googlesource.com/24420
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e75c899a10c9321d1b4935c34401f66b2abad83a
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Jun 23 13:32:50 2016 -0400

    reflect: optimize (reflect.Type).Name
    
    Improves JSON decoding on linux/amd64.
    
    name                   old time/op    new time/op    delta
    CodeUnmarshal-40         89.3ms ± 2%    86.3ms ± 2%  -3.31%  (p=0.000 n=22+22)
    
    name                   old speed      new speed      delta
    CodeUnmarshal-40       21.7MB/s ± 2%  22.5MB/s ± 2%  +3.44%  (p=0.000 n=22+22)
    
    Updates #16117
    
    Change-Id: I52acf31d7729400cfe6693e46292d41e1addba3d
    Reviewed-on: https://go-review.googlesource.com/24410
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e369490fb7db5f2d42bb0e8ee19b48378dee0ebf
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Jun 23 10:59:38 2016 -0400

    cmd/compile, etc: bring back ptrToThis
    
    This was removed in CL 19695 but it slows down reflect.New, which ends
    up on the hot path of things like JSON decoding.
    
    There is no immediate cost in binary size, but it will make it harder to
    further shrink run time type information in Go 1.8.
    
    Before
    
            BenchmarkNew-40         30000000                36.3 ns/op
    
    After
    
            BenchmarkNew-40         50000000                29.5 ns/op
    
    Fixes #16161
    Updates #16117
    
    Change-Id: If7cb7f3e745d44678f3f5cf3a5338c59847529d2
    Reviewed-on: https://go-review.googlesource.com/24400
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
```
