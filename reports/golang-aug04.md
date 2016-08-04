# August 4, 2016 Report

Number of commits: 17

## Compilation time

* github.com/coreos/etcd/cmd: from 13.38525805s to 13.195611041s, -1.42%
* github.com/boltdb/bolt/cmd/bolt: from 647.687973ms to 629.838236ms, -2.76%
* github.com/gogits/gogs: from 13.366564493s to 13.412733112s, +0.35%
* github.com/spf13/hugo: from 7.702782581s to 9.068950077s, +17.74%
* github.com/nsqio/nsq/apps/nsqd: from 2.1388922s to 2.082663096s, -2.63%
* github.com/mholt/caddy: from 277.424657ms to 292.226838ms, +5.34%

## Binary size:

* github.com/coreos/etcd/cmd: from 26646244 to 26654613, ~
* github.com/boltdb/bolt/cmd/bolt: from 2679371 to 2679371, ~
* github.com/gogits/gogs: from 23810085 to 23818454, ~
* github.com/spf13/hugo: from 15860751 to 15869120, ~
* github.com/nsqio/nsq/apps/nsqd: from 10064210 to 10072579, ~
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               188           187           -0.53%
BenchmarkMsgpUnmarshal-4             407           407           +0.00%
BenchmarkEasyJsonMarshal-4           1604          1613          +0.56%
BenchmarkEasyJsonUnmarshal-4         2142          2321          +8.36%
BenchmarkFlatBuffersMarshal-4        385           366           -4.94%
BenchmarkFlatBuffersUnmarshal-4      291           289           -0.69%
BenchmarkGogoprotobufMarshal-4       163           160           -1.84%
BenchmarkGogoprotobufUnmarshal-4     251           254           +1.20%

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
commit 50edddb7383a0e4aef990a085777f958c5d8a2b8
Author: Chris Broadfoot <cbro@golang.org>
Date:   Tue Aug 2 14:30:06 2016 -0700

    VERSION: remove erroneously committed VERSION file
    
    Change-Id: I1134a4758b7e1a7da243c56f12ad9d2200c8ba41
    Reviewed-on: https://go-review.googlesource.com/25414
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ae68090d000c0a4a6cda0ef44a59809f914e2358
Merge: c628d83 2da5633
Author: Chris Broadfoot <cbro@golang.org>
Date:   Tue Aug 2 14:05:48 2016 -0700

    all: merge master into release-branch.go1.7
    
    Change-Id: I177856ea2bc9943cbde28ca9afa145b6ea5b0942

commit 2da5633eb9091608047881953f75b489a3134cdc
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Aug 1 21:54:40 2016 -0700

    runtime: fix nanotime for macOS Sierra, again.
    
    macOS Sierra beta4 changed the kernel interface for getting time.
    DX now optionally points to an address for additional info.
    Set it to zero to avoid corrupting memory.
    
    Fixes #16570
    
    Change-Id: I9f537e552682045325cdbb68b7d0b4ddafade14a
    Reviewed-on: https://go-review.googlesource.com/25400
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6317c213c953d0879fe88593b4372f03d25f369b
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Mon Aug 1 20:04:25 2016 -0700

    cmd/doc: ensure functions with unexported return values are shown
    
    The commit in golang.org/cl/22354 groups constructors functions under
    the type that they construct to. However, this caused a minor regression
    where functions that had unexported return values were not being printed
    at all. Thus, we forgo the grouping logic if the type the constructor falls
    under is not going to be printed.
    
    Fixes #16568
    
    Change-Id: Idc14f5d03770282a519dc22187646bda676af612
    Reviewed-on: https://go-review.googlesource.com/25369
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c628d83ec5309cd679e16c734456fed1b9a85806
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Aug 1 18:52:52 2016 -0700

    go1.7rc4
    
    Change-Id: Icf861dd28bfe29a2e4b90529e53644b43b6f7969
    Reviewed-on: https://go-review.googlesource.com/25368
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f5758739a8f011c1d146a7736ab8f0d2834e1783
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Mon Aug 1 14:33:19 2016 -0700

    cmd/doc: handle embedded interfaces properly
    
    Changes made:
    * Disallow star expression on interfaces as this is not possible.
    * Show an embedded "error" in an interface as public similar to
    how godoc does it.
    * Properly handle selector expressions in both structs and interfaces.
    This is possible since a type may refer to something defined in
    another package (e.g. io.Reader).
    
    Before:
    <<<
    $ go doc runtime.Error
    type Error interface {
    
        // RuntimeError is a no-op function but
        // serves to distinguish types that are run time
        // errors from ordinary errors: a type is a
        // run time error if it has a RuntimeError method.
        RuntimeError()
        // Has unexported methods.
    }
    
    $ go doc compress/flate Reader
    doc: invalid program: unexpected type for embedded field
    doc: invalid program: unexpected type for embedded field
    type Reader interface {
        io.Reader
        io.ByteReader
    }
    >>>
    
    After:
    <<<
    $ go doc runtime.Error
    type Error interface {
        error
    
        // RuntimeError is a no-op function but
        // serves to distinguish types that are run time
        // errors from ordinary errors: a type is a
        // run time error if it has a RuntimeError method.
        RuntimeError()
    }
    
    $ go doc compress/flate Reader
    type Reader interface {
        io.Reader
        io.ByteReader
    }
    >>>
    
    Fixes #16567
    
    Change-Id: I272dede971eee9f43173966233eb8810e4a8c907
    Reviewed-on: https://go-review.googlesource.com/25365
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8ea89ba858823212114de952b8a1375ceb82587f
Merge: 8707f31 28ee179
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Aug 1 18:26:06 2016 -0700

    all: merge master into release-branch.go1.7
    
    Change-Id: Ifb9647fa9817ed57aa4835a35a05020aba00a24e

commit 28ee17965703c4ef81cc97e5088539fe3e8e541f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Jul 28 13:42:11 2016 +0200

    net: prevent cancelation goroutine from adjusting fd timeout after connect
    
    This was previously fixed in https://golang.org/cl/21497 but not enough.
    
    Fixes #16523
    
    Change-Id: I678543a656304c82d654e25e12fb094cd6cc87e8
    Reviewed-on: https://go-review.googlesource.com/25330
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2629446df0cb906986f377d45cde307ffdae9675
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Aug 2 00:41:12 2016 +0000

    doc/go1.7.html: mention Server.Serve HTTP/2 behavior change
    
    Fixes #16550
    Updates #15908
    
    Change-Id: Ic951080dbc88f96e4c00cdb3ffe24a5c03079efd
    Reviewed-on: https://go-review.googlesource.com/25389
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit c558a539b5efaeda4b6f8e61f51c21f64d1b94f6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Aug 1 23:44:22 2016 +0000

    net/http: update bundled http2
    
    Updates bundled http2 to x/net/http2 rev 28d1bd4f for:
    
        http2: make Transport work around mod_h2 bug
        https://golang.org/cl/25362
    
        http2: don't ignore DATA padding in flow control
        https://golang.org/cl/25382
    
    Updates #16519
    Updates #16556
    Updates #16481
    
    Change-Id: I51f5696e977c91bdb2d80d2d56b8a78e3222da3f
    Reviewed-on: https://go-review.googlesource.com/25388
    Reviewed-by: Chris Broadfoot <cbro@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 111d590f86e2c9a55ec08d95fc4e9adea9232f0c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Jul 28 12:22:49 2016 -0400

    cmd/compile: fix possible spill of invalid pointer with DUFFZERO on AMD64
    
    SSA compiler on AMD64 may spill Duff-adjusted address as scalar. If
    the object is on stack and the stack moves, the spilled address become
    invalid.
    
    Making the spill pointer-typed does not work. The Duff-adjusted address
    points to the memory before the area to be zeroed and may be invalid.
    This may cause stack scanning code panic.
    
    Fix it by doing Duff-adjustment in genValue, so the intermediate value
    is not seen by the reg allocator, and will not be spilled.
    
    Add a test to cover both cases. As it depends on allocation, it may
    be not always triggered.
    
    Fixes #16515.
    
    Change-Id: Ia81d60204782de7405b7046165ad063384ede0db
    Reviewed-on: https://go-review.googlesource.com/25309
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 8707f31c0abc6b607014e843b7cc188b3019daa9
Author: Chris Broadfoot <cbro@golang.org>
Date:   Thu Jul 21 14:06:27 2016 -0700

    go1.7rc3
    
    Change-Id: Iaef13003979c68926c260c415d6074a50ae137b2
    Reviewed-on: https://go-review.googlesource.com/25142
    Run-TryBot: Chris Broadfoot <cbro@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 16a2af03f17e5b2bcf468442e66ef7a99ae55c70
Merge: 0ebf6ce 243d51f
Author: Chris Broadfoot <cbro@golang.org>
Date:   Thu Jul 21 12:38:13 2016 -0700

    all: merge master into release-branch.go1.7
    
    Change-Id: I2511c3f7583887b641c9b3694aae54789fbc5342

commit 0ebf6ce087388cdd501a02ff92f2f8cafc3e1378
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Jul 18 08:19:17 2016 -0700

    [release-branch.go1.7] go1.7rc2
    
    Change-Id: I5473071f672f8352fbd3352e158d5be12823b58a
    Reviewed-on: https://go-review.googlesource.com/25017
    Run-TryBot: Chris Broadfoot <cbro@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit cad4e97af8f2e0b9f09b97f67fb3a89ced2e9021
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Jul 18 06:05:24 2016 +0000

    [release-branch.go1.7] net/http, net/http/cgi: fix for CGI + HTTP_PROXY security issue
    
    Because,
    
    * The CGI spec defines that incoming request header "Foo: Bar" maps to
      environment variable HTTP_FOO == "Bar". (see RFC 3875 4.1.18)
    
    * The HTTP_PROXY environment variable is conventionally used to configure
      the HTTP proxy for HTTP clients (and is respected by default for
      Go's net/http.Client and Transport)
    
    That means Go programs running in a CGI environment (as a child
    process under a CGI host) are vulnerable to an incoming request
    containing "Proxy: attacker.com:1234", setting HTTP_PROXY, and
    changing where Go by default proxies all outbound HTTP requests.
    
    This is CVE-2016-5386, aka https://httpoxy.org/
    
    Fixes #16405
    
    Change-Id: I6f68ade85421b4807785799f6d98a8b077e871f0
    Reviewed-on: https://go-review.googlesource.com/25010
    Run-TryBot: Chris Broadfoot <cbro@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>
    Reviewed-on: https://go-review.googlesource.com/25013

commit 53da5fd4d431881bb3583c9790db7735a6530a1b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jul 7 16:41:29 2016 -0700

    [release-branch.go1.7] runtime: fix nanotime for macOS Sierra
    
    In the beta version of the macOS Sierra (10.12) release, the
    gettimeofday system call changed on x86. Previously it always returned
    the time in the AX/DX registers. Now, if AX is returned as 0, it means
    that the system call has stored the values into the memory pointed to by
    the first argument, just as the libc gettimeofday function does. The
    libc function handles both cases, and we need to do so as well.
    
    Fixes #16272.
    
    Change-Id: Ibe5ad50a2c5b125e92b5a4e787db4b5179f6b723
    Reviewed-on: https://go-review.googlesource.com/24812
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-on: https://go-review.googlesource.com/24755
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit a91416e7abf2236909f99aea85accfe98a9ba1fd
Author: Chris Broadfoot <cbro@google.com>
Date:   Thu Jul 7 19:32:35 2016 -0700

    [release-branch.go1.7] go1.7rc1
    
    Change-Id: Ifbf1c13ce740428add68d68133c7f10876bad404
    Reviewed-on: https://go-review.googlesource.com/24816
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
```
