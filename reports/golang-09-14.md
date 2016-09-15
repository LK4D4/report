# September 14, 2016 Report

Number of commits: 119

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 630.086443ms to 587.081681ms, -6.83%
* github.com/coreos/etcd: from 13.878589831s to 13.602607687s, -1.99%
* github.com/gogits/gogs: from 13.452833954s to 14.155295676s, +5.22%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 33.700518888s to 39.268723209s, +16.52%
* github.com/influxdata/influxdb/cmd/influxd: from 7.845155891s to 7.170265082s, -8.60%
* github.com/junegunn/fzf/src/fzf: from 1.14285616s to 1.06083151s, -7.18%
* github.com/mholt/caddy/caddy: from 7.050351884s to 6.861908564s, -2.67%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.518100098s to 1.501055337s, -1.12%
* github.com/nsqio/nsq/apps/nsqd: from 2.350204869s to 2.2358869s, -4.86%
* github.com/prometheus/prometheus/cmd/prometheus: from 11.60342751s to 11.067167949s, -4.62%
* github.com/spf13/hugo: from 8.572691205s to 8.359096979s, -2.49%
* golang.org/x/tools/cmd/guru: from 2.930856654s to 3.110333713s, +6.12%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2598068 to 2590003, -0.31%
* github.com/coreos/etcd: from 23727032 to 23765040, +0.16%
* github.com/gogits/gogs: from 23121900 to 23104459, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 27208776 to 27246656, +0.14%
* github.com/influxdata/influxdb/cmd/influxd: from 16068296 to 16095202, +0.17%
* github.com/junegunn/fzf/src/fzf: from 3130128 to 3126704, -0.11%
* github.com/mholt/caddy/caddy: from 14627013 to 14662466, +0.24%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4426645 to 4394640, -0.72%
* github.com/nsqio/nsq/apps/nsqd: from 9622037 to 9657057, +0.36%
* github.com/prometheus/prometheus/cmd/prometheus: from 23469583 to 23496910, +0.12%
* github.com/spf13/hugo: from 16056117 to 16068033, ~
* golang.org/x/tools/cmd/guru: from 7570871 to 7570536, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               192           193           +0.52%
BenchmarkMsgpUnmarshal-4             416           414           -0.48%
BenchmarkEasyJsonMarshal-4           1499          1526          +1.80%
BenchmarkEasyJsonUnmarshal-4         2676          2129          -20.44%
BenchmarkFlatBuffersMarshal-4        683           371           -45.68%
BenchmarkFlatBuffersUnmarshal-4      301           296           -1.66%
BenchmarkGogoprotobufMarshal-4       166           161           -3.01%
BenchmarkGogoprotobufUnmarshal-4     257           262           +1.95%

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

* [bytes: improve WriteRune performance](https://github.com/golang/go/commit/2321895fe2a2def7b511453329f4cd8662230256)
* [encoding/json: Use a lookup table for safe characters](https://github.com/golang/go/commit/ed8f207940c8787d344664a43071b235e2ce5c68)
* [context: reduce memory usage of context tree](https://github.com/golang/go/commit/39382793d3a9e7a0720e6dbf8be4b19e8519af19)
* [cmd/compile: add SSA backend for s390x and enable by default](https://github.com/golang/go/commit/6ec993adc3373d31392b301ebe0c376b02d68051)
* [encoding/asn1: reduce allocations in Marshal](https://github.com/golang/go/commit/ae4aac00bba5d1d616408a1c07bd4ef5691e3a00)

## GIT Log

```
commit 7f583a4d721c94967fe3ce098b3eff902cba043b
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 13 19:19:36 2016 -0400

    cmd/dist: re-enable internal PIE test
    
    For #17068
    
    Change-Id: I4e3ab166f08100292b779b651a9acfbfb44a55cd
    Reviewed-on: https://go-review.googlesource.com/29119
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5940a007c14af11fe35f48d8f7daa3cd9036aaa3
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 13 08:24:07 2016 -0400

    cmd/link: R_ADDR dynamic relocs for internal PIE
    
    This gets -buildmode=pie -ldflags=-linkmode=internal working on
    Ubuntu 16.04.
    
    Fixes #17068
    
    Change-Id: Ice5036199005fb528cc58279a7f057170dc6b73d
    Reviewed-on: https://go-review.googlesource.com/29118
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6fd8c006adf792e523232a026823a5444893f28f
Author: David Stainton <dstainton415@gmail.com>
Date:   Fri Aug 12 22:15:21 2016 +0000

    syscall: add bounds checking and error returns to ParseNetlinkMessage
    
    Fixes #16681
    
    Change-Id: I6ff7ec81fe48ab06be3aae5b7ff92e9dc70960c3
    Reviewed-on: https://go-review.googlesource.com/26990
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit 33e63ebc20e38d20077c1f184f05a4a0656e189a
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Sep 11 05:06:29 2016 -0700

    os: add more examples
    
    Updates #16360.
    
    Adds examples for:
    + Chmod
    + Chtimes
    + FileMode
    
    Change-Id: I1b61ee0392fa3774593a7f36aaf0fa1e484c778b
    Reviewed-on: https://go-review.googlesource.com/28963
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit a562351e512682046623ecbd106f4973f3525be9
Author: Aaron Jacobs <jacobsa@google.com>
Date:   Wed Sep 14 12:57:15 2016 +1000

    net/http: clarify Request.ContentLength behavior on the client.
    
    While you could argue the previous wording technically said that -1 is
    an acceptable way to indicate "unknown" on the client, it could be read
    as ambiguous. Now it's clear that both 0 and -1 mean unknown.
    
    Change-Id: I3bc5a3fd5afd1999e487296ec121eb548415e6b0
    Reviewed-on: https://go-review.googlesource.com/29130
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6488201b9b980277b41de540a88f4289aee64e1c
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 13 19:04:25 2016 -0400

    cmd/internal/obj: regenerate stringer values
    
    Created by running 'go generate'.
    
    That made debugging fun today.
    
    Change-Id: I9ffe00877851f2b198275220ad6058b9005daa72
    Reviewed-on: https://go-review.googlesource.com/29117
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 567aefd99ddc66d596fffd442b6c014d83ac5eb9
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 6 07:46:59 2016 -0400

    cmd/link: address comments from CL 28540
    
    Change-Id: I11899096c71ee0e24e902c87914601fcd7ffd7a9
    Reviewed-on: https://go-review.googlesource.com/28967
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 429eb3c6960a0b919a150aaa570e6a3f6f8758e9
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Sep 13 15:33:55 2016 -0700

    cmd/compile: remove two unnecessary Pkg fields
    
    Exported is no longer used since removing the text-format exporter,
    and Safe is only used within importfile so it can be made into a local
    variable.
    
    Change-Id: I92986f704d7952759c79d9243620a22c24602333
    Reviewed-on: https://go-review.googlesource.com/29115
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2b5c18c99e1be7c779706e23f99717a84d7d882b
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Sep 13 14:05:14 2016 -0700

    cmd/compile/internal/gc: eliminate bstdout
    
    Just use Ctxt.Bso instead.
    
    Change-Id: I68f1639f0b4c238ae5499ef49e78a5d734417979
    Reviewed-on: https://go-review.googlesource.com/29114
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ae4aac00bba5d1d616408a1c07bd4ef5691e3a00
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Tue Aug 2 14:41:53 2016 +0900

    encoding/asn1: reduce allocations in Marshal
    
    Current code uses trees of bytes.Buffer as data representation.
    Each bytes.Buffer takes 4k bytes at least, so it's waste of memory.
    The change introduces trees of lazy-encoder as
    alternative one which reduce allocations.
    
    name       old time/op    new time/op    delta
    Marshal-4    64.7µs ± 2%    42.0µs ± 1%  -35.07%   (p=0.000 n=9+10)
    
    name       old alloc/op   new alloc/op   delta
    Marshal-4    35.1kB ± 0%     7.6kB ± 0%  -78.27%  (p=0.000 n=10+10)
    
    name       old allocs/op  new allocs/op  delta
    Marshal-4       503 ± 0%       293 ± 0%  -41.75%  (p=0.000 n=10+10)
    
    Change-Id: I32b96c20b8df00414b282d69743d71a598a11336
    Reviewed-on: https://go-review.googlesource.com/27030
    Reviewed-by: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ee3f3a60070ee9edeb3f10fa2e4b90404068cb3a
Author: Adam Langley <agl@golang.org>
Date:   Sun Sep 11 17:14:51 2016 -0700

    crypto/rsa: ensure that generating toy RSA keys doesn't loop.
    
    If there are too few primes of the given length then it can be
    impossible to generate an RSA key with n distinct primes.
    
    This change approximates the expected number of candidate primes and
    causes key generation to return an error if it's unlikely to succeed.
    
    Fixes #16596.
    
    Change-Id: I53b60d0cb90e2d0e6f0662befa64d13f24af51a7
    Reviewed-on: https://go-review.googlesource.com/28969
    Reviewed-by: Minux Ma <minux@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7e2b5a102e1c7fcc314b5e58151043530ea1ffe9
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Sep 13 15:57:52 2016 -0400

    test: re-enable phi optimization test
    
    CL 28978 (6ec993a) accidentally disabled the test (it would only
    run if amd64 AND s390x, whereas it should be amd64 OR s390x).
    
    Change-Id: I23c1ad71724ff55f5808d5896b19b62c8ec5af76
    Reviewed-on: https://go-review.googlesource.com/28981
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6ec993adc3373d31392b301ebe0c376b02d68051
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Sep 12 14:50:10 2016 -0400

    cmd/compile: add SSA backend for s390x and enable by default
    
    The new SSA backend modifies the ABI slightly: R0 is now a usable
    general purpose register.
    
    Fixes #16677.
    
    Change-Id: I367435ce921e0c7e79e021c80cf8ef5d1d1466cf
    Reviewed-on: https://go-review.googlesource.com/28978
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit b7e53038b8abb4d82cf25cb844395af602150a29
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Sep 11 02:00:38 2016 +0000

    net/http: make Transport support international domain names
    
    This CL makes code like this work:
    
         res, err := http.Get("https://фу.бар/баз")
    
    So far, IDNA support is limited to the http1 and http2 Transports.
    The http package is currently responsible for converting domain names
    into Punycode before calling the net layer. The http package also has
    to Punycode-ify the hostname for the Host & :authority headers for
    HTTP/1 and HTTP/2, respectively.
    
    No automatic translation from Punycode back to Unicode is performed,
    per Go's historical behavior. Docs are updated where relevant.  No
    changes needed to the Server package. Things are already in ASCII
    at that point.
    
    No changes to the net package, at least yet.
    
    Updates x/net/http2 to git rev 57c7820 for https://golang.org/cl/29071
    
    Updates #13835
    
    Change-Id: I1e9a74c60d00a197ea951a9505da5c3c3187099b
    Reviewed-on: https://go-review.googlesource.com/29072
    Reviewed-by: Chris Broadfoot <cbro@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 802cb5927f1e163749331c9f6cfb414cb0c753b9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 13 18:51:16 2016 +0000

    net/http: update bundled x/net/http2
    
    Updates x/net/http2 (and x/net/lex/httplex) to git rev 749a502 for:
    
       http2: don't sniff first Request.Body byte in Transport until we have a conn
       https://golang.org/cl/29074
       Fixes #17071
    
       http2: add Transport support for unicode domain names
       https://golang.org/cl/29071
       Updates #13835
    
       http2: don't send bogus :path pseudo headers if Request.URL.Opaque is set
       https://golang.org/cl/27632
         +
       http2: fix bug where '*' as a valid :path value in Transport
       https://golang.org/cl/29070
       Updates #16847
    
       http2: fix all vet warnings
       https://golang.org/cl/28344
       Updates #16228
       Updates #11041
    
    Also uses the new -underscore flag to x/tools/cmd/bundle from
    https://golang.org/cl/29086
    
    Change-Id: Ica0f6bf6e33266237e37527a166a783d78c059c4
    Reviewed-on: https://go-review.googlesource.com/29110
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit 9980b70cb460f27907a003674ab1b9bea24a847c
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Jul 19 11:47:53 2016 -0700

    runtime: limit the number of map overflow buckets
    
    Consider repeatedly adding many items to a map
    and then deleting them all, as in #16070. The map
    itself doesn't need to grow above the high water
    mark of number of items. However, due to random
    collisions, the map can accumulate overflow
    buckets.
    
    Prior to this CL, those overflow buckets were
    never removed, which led to a slow memory leak.
    
    The problem with removing overflow buckets is
    iterators. The obvious approach is to repack
    keys and values and eliminate unused overflow
    buckets. However, keys, values, and overflow
    buckets cannot be manipulated without disrupting
    iterators.
    
    This CL takes a different approach, which is to
    reuse the existing map growth mechanism,
    which is well established, well tested, and
    safe in the presence of iterators.
    When a map has accumulated enough overflow buckets
    we trigger map growth, but grow into a map of the
    same size as before. The old overflow buckets will
    be left behind for garbage collection.
    
    For the code in #16070, instead of climbing
    (very slowly) forever, memory usage now cycles
    between 264mb and 483mb every 15 minutes or so.
    
    To avoid increasing the size of maps,
    the overflow bucket counter is only 16 bits.
    For large maps, the counter is incremented
    stochastically.
    
    Fixes #16070
    
    Change-Id: If551d77613ec6836907efca58bda3deee304297e
    Reviewed-on: https://go-review.googlesource.com/25049
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 0cd3ecb016e0c3f0656877a20ca37eabe4fd0f8f
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Sep 12 17:30:35 2016 -0700

    cmd/compile: reduce allocs some more
    
    Also: update fmt_test.go.
    
    Together with the previous commits, we are now at or below c85b77c
    levels in terms of allocation for the benchmark described in #16897
    (old = c85b77c, new = this commit):
    
    name       old time/op     new time/op     delta
    Template       297ms ± 5%      284ms ± 3%  -4.53%  (p=0.000 n=27+29)
    Unicode        159ms ± 5%      151ms ± 5%  -4.91%  (p=0.000 n=28+30)
    GoTypes        985ms ± 5%      935ms ± 2%  -5.13%  (p=0.000 n=28+29)
    
    name       old alloc/op    new alloc/op    delta
    Template      46.8MB ± 0%     45.7MB ± 0%  -2.37%  (p=0.000 n=30+30)
    Unicode       37.8MB ± 0%     37.9MB ± 0%  +0.29%  (p=0.000 n=29+30)
    GoTypes        143MB ± 0%      138MB ± 0%  -3.64%  (p=0.000 n=29+30)
    
    name       old allocs/op   new allocs/op   delta
    Template        444k ± 0%       440k ± 0%  -0.94%  (p=0.000 n=30+30)
    Unicode         369k ± 0%       369k ± 0%  +0.19%  (p=0.000 n=29+30)
    GoTypes        1.35M ± 0%      1.34M ± 0%  -1.24%  (p=0.000 n=30+30)
    
    For #16897.
    
    Change-Id: Iedbeb408e2f1e68dd4a3201bf8813c8066ebf7ed
    Reviewed-on: https://go-review.googlesource.com/29089
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b6946fb120ed0c176162e4f632fb8cd062144af3
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Wed Sep 7 14:42:43 2016 -0500

    cmd/asm: ppc64le support for ISEL for use by SSA
    
    This adds the support for the ppc64le isel instruction so
    it can be used by SSA.
    
    Fixed #16771
    
    Change-Id: Ia2517f0834ff5e7ad927e218b84493e0106ab4a7
    Reviewed-on: https://go-review.googlesource.com/28611
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 07b8011393dc3d3a78b3cd0857a31da339985994
Author: Paul Borman <borman@google.com>
Date:   Tue Jul 12 08:54:09 2016 -0700

    text/template: improve lexer performance in finding left delimiters.
    
    The existing implementation calls l.next for each run up to the next
    instance of the left delimiter ({{).  For ascii text, this is multiple
    function calls per byte.  Change to use strings.Index to find the left
    delimiter.  The performace improvement ranges from 1:1 (no text outside
    of {{}}'s) to multiple times faster (9:1 was seen on 8K of text with no
    {{ }}'s).
    
    Change-Id: I2f82bea63b78b6714f09a725f7b2bbb00a3448a3
    Reviewed-on: https://go-review.googlesource.com/24863
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Rob Pike <r@golang.org>

commit 8f9671d11a219d8fc9a6176ddf9939e743982ffc
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Tue Sep 13 13:29:13 2016 +1200

    cmd/link: fix -buildmode=pie / -linkshared combination
    
    main.main and main.init were not being marked as reachable.
    
    Fixes #17076
    
    Change-Id: Ib3e29bd35ba6252962e6ba89173ca321ed6849b9
    Reviewed-on: https://go-review.googlesource.com/28996
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 38d35e714a55f2e4bb67caadac7e61f8c1967d88
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Sep 12 15:24:11 2016 -0400

    cmd/compile, runtime/internal/atomic: intrinsify And8, Or8 on ARM64
    
    Also add assembly implementation, in case intrinsics is disabled.
    
    Change-Id: Iff0a8a8ce326651bd29f6c403f5ec08dd3629993
    Reviewed-on: https://go-review.googlesource.com/28979
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 09686a58734382ace059f1dbd882dadbb39b2268
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Sep 12 17:08:21 2016 -0700

    cmd/compile: remove another bytes.Buffer use in fmt.go
    
    Missed in prior commit.
    
    Change-Id: Ib3a41fb4e4d41feeb28c316fe70a329c73e72379
    Reviewed-on: https://go-review.googlesource.com/29088
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 37d452c3e9c1fe354375ad41ae0b952b563dfbe4
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Sep 12 13:44:43 2016 -0700

    cmd/compile: reduce allocs to c85b77c (pre-fmt.go change) levels
    
    Linker and reflect info generation (reflect.go) relies on formatting
    of types (tconv). The fmt.Format based approach introduces extra
    allocations, which matter in those cases. Resurrected sconv and tconv
    code from commit c85b77c (fmt.go only); and adjusted it slightly.
    The formatter-based approach is still used throughout the rest of the
    compiler, but reflect.go now uses the tconv method that simply returns
    the desired string.
    
    (The timing data below may not be accurate; I've included it only for
    comparison with the numbers in issue #16897).
    
    name       old time/op     new time/op     delta
    Template       297ms ± 2%      288ms ± 3%  -3.12%        (p=0.000 n=27+29)
    Unicode        155ms ± 5%      150ms ± 5%  -3.26%        (p=0.000 n=30+30)
    GoTypes        1.00s ± 3%      0.95s ± 3%  -4.51%        (p=0.000 n=28+29)
    
    name       old alloc/op    new alloc/op    delta
    Template      46.8MB ± 0%     46.5MB ± 0%  -0.65%        (p=0.000 n=28+30)
    Unicode       37.9MB ± 0%     37.8MB ± 0%  -0.24%        (p=0.000 n=29+30)
    GoTypes        144MB ± 0%      143MB ± 0%  -0.68%        (p=0.000 n=30+30)
    
    name       old allocs/op   new allocs/op   delta
    Template        469k ± 0%       446k ± 0%  -5.01%        (p=0.000 n=29+30)
    Unicode         375k ± 0%       369k ± 0%  -1.62%        (p=0.000 n=30+28)
    GoTypes        1.47M ± 0%      1.37M ± 0%  -6.29%        (p=0.000 n=30+30)
    
    The code for sconv/tconv in fmt.go now closely match the code from c85b77c
    again; except that the functions are now methods. Removing the use of
    the bytes.Buffer in tconv and special-caseing interface{} has helped a
    small amount as well:
    
    name       old time/op     new time/op     delta
    Template       299ms ± 3%      288ms ± 3%  -3.83%        (p=0.000 n=29+29)
    Unicode        156ms ± 5%      150ms ± 5%  -3.56%        (p=0.000 n=30+30)
    GoTypes        960ms ± 2%      954ms ± 3%  -0.58%        (p=0.037 n=26+29)
    
    name       old alloc/op    new alloc/op    delta
    Template      46.6MB ± 0%     46.5MB ± 0%  -0.22%        (p=0.000 n=30+30)
    Unicode       37.8MB ± 0%     37.8MB ± 0%    ~           (p=0.075 n=30+30)
    GoTypes        143MB ± 0%      143MB ± 0%  -0.31%        (p=0.000 n=30+30)
    
    name       old allocs/op   new allocs/op   delta
    Template        447k ± 0%       446k ± 0%  -0.28%        (p=0.000 n=30+30)
    Unicode         369k ± 0%       369k ± 0%  -0.03%        (p=0.032 n=30+28)
    GoTypes        1.38M ± 0%      1.37M ± 0%  -0.35%        (p=0.000 n=29+30)
    
    Comparison between c85b77c and now (see issue #16897):
    
    name       old time/op     new time/op     delta
    Template       307ms ± 4%      288ms ± 3%  -6.24%  (p=0.000 n=29+29)
    Unicode        164ms ± 4%      150ms ± 5%  -8.20%  (p=0.000 n=30+30)
    GoTypes        1.01s ± 3%      0.95s ± 3%  -5.72%  (p=0.000 n=30+29)
    
    name       old alloc/op    new alloc/op    delta
    Template      46.8MB ± 0%     46.5MB ± 0%  -0.66%  (p=0.000 n=29+30)
    Unicode       37.8MB ± 0%     37.8MB ± 0%  -0.13%  (p=0.000 n=30+30)
    GoTypes        143MB ± 0%      143MB ± 0%  -0.11%  (p=0.000 n=30+30)
    
    name       old allocs/op   new allocs/op   delta
    Template        444k ± 0%       446k ± 0%  +0.48%  (p=0.000 n=30+30)
    Unicode         369k ± 0%       369k ± 0%  +0.09%  (p=0.000 n=30+28)
    GoTypes        1.35M ± 0%      1.37M ± 0%  +1.47%  (p=0.000 n=30+30)
    
    There's still a small increase (< 1.5%) for GoTypes but pending a complete
    rewrite of fmt.go, this seems ok again.
    
    Fixes #16897.
    
    Change-Id: I7e0e56cd1b9f981252eded917f5752259d402354
    Reviewed-on: https://go-review.googlesource.com/29087
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c345a3913f54ab88720654b3dc9b2c34779fca0c
Author: Keith Randall <khr@golang.org>
Date:   Fri Sep 9 13:11:07 2016 -0700

    cmd/compile: get rid of BlockCall
    
    No need for it, we can treat calls as (mostly) normal values
    that take a memory and return a memory.
    
    Lowers the number of basic blocks needed to represent a function.
    "go test -c net/http" uses 27% fewer basic blocks.
    Probably doesn't affect generated code much, but should help
    various passes whose running time and/or space depends on
    the number of basic blocks.
    
    Fixes #15631
    
    Change-Id: I0bf21e123f835e2cfa382753955a4f8bce03dfa6
    Reviewed-on: https://go-review.googlesource.com/28950
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit d00a3cead8423c8af6c5781aa2a3efe0a9a442ae
Author: Keith Randall <khr@golang.org>
Date:   Mon Sep 12 15:19:36 2016 -0700

    runtime: make gdb test resilient to line numbering
    
    Don't break on line number, instead break on the actual call.
    This makes the test more robust to line numbering changes in the backend.
    
    A CL (28950) changed the generated code line numbering slightly.  A MOVW
    $0, R0 instruction at the start of the function changed to line
    10 (because several constant zero instructions got CSEd, and one gets
    picked arbitrarily).  That's too fragile for a test.
    
    Change-Id: I5d6a8ef0603de7d727585004142780a527e70496
    Reviewed-on: https://go-review.googlesource.com/29085
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6ca7c2055bd114c08426d3dbdee444b280b569b6
Author: Keith Randall <khr@golang.org>
Date:   Mon Sep 12 13:47:55 2016 -0700

    cmd/compile: fix tuple-generating flag ops as clobbering flags
    
    If an op generates a tuple, and part of that tuple is of flags type,
    then treat the op as clobbering flags.
    
    Normally this doesn't matter because we do:
    
    v1 = ADDS        <int32, flags>
    v2 = Select0 v1  <int32>
    v3 = Select1 v1  <flags>
    
    And v3 will do the right clobbering of flags.  But in the rare
    cases where we issue a tuple-with-flag op and the flag portion
    is dead, then we never issue a Select1.  But v1 still clobbers flags,
    so we need to respect that.
    
    Fixes builder failure in CL 28950.
    
    Change-Id: I589089fd81aaeaaa9750bb8d85e7b10199aaa002
    Reviewed-on: https://go-review.googlesource.com/29083
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit e42ae65a85079a9f39d0fb4837a78172ad898d84
Author: Quentin Smith <quentin@golang.org>
Date:   Wed Sep 7 18:07:45 2016 -0400

    time: improve Truncate and Round documentation
    
    Truncate and Round operate on absolute time, which means that
    Truncate(Hour) may return a time with non-zero Minute(). Document that
    more clearly, and remove the misleading example which suggests it is
    safe.
    
    Updates #16647
    
    Change-Id: I930584ca030dd12849195d45e49ed2fb74e0c9ac
    Reviewed-on: https://go-review.googlesource.com/28730
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 1ee544641450236e8c78d8d408e6cb8ab69cee60
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Sep 12 13:39:54 2016 -0700

    cmd/compile: remove incannedimport
    
    This used to be used to give special semantics to the builtin
    definitions of package runtime and unsafe, but none of those are
    relevant anymore:
    
    - The builtin runtime and unsafe packages do not risk triggering false
      import cycles since they no longer contain `import "runtime"`.
    
    - bimport.go never creates ODCLTYPE, so no need to special case them.
    
    - "incannedimport != 0" is only true when "importpkg != nil" anyway,
      so "incannedimport == 0 && importpkg == nil" is equivalent to just
      "importpkg == nil".
    
    Change-Id: I076f15dd705d4962e7a4c33972e304ef67e7effb
    Reviewed-on: https://go-review.googlesource.com/29084
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4c83d2914467e670d680b967ec689f8434136188
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 12 15:10:04 2016 -0400

    cmd/dist: disable test of internal PIE linking
    
    Updates #17068
    
    Change-Id: I61b75ec07ca8705a678677d262e11b16848cddf3
    Reviewed-on: https://go-review.googlesource.com/29079
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 6758eedf898c48d6ca4abd42f44622abf7c005c2
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Sep 12 13:27:36 2016 -0700

    cmd/compile: remove Pointer from builtin/unsafe.go
    
    We already explicitly construct the "unsafe.Pointer" type in typeinit
    because we need it for Types[TUNSAFEPTR]. No point in also having it
    in builtin/unsafe.go if it just means (*importer).importtype needs to
    fix it.
    
    Change-Id: Ife8a5a73cbbe2bfcabe8b25ee4f7e0f5fd0570b4
    Reviewed-on: https://go-review.googlesource.com/29082
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b11c79fd07c80e6902fa26045ac566ddc2f1250d
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Mon Sep 12 11:43:55 2016 -0700

    cmd/compile: deduplicate importtype and (*importer).importtype
    
    Change-Id: I7bfb0e5e71fc26448b0d5d3801cd6e50c8b48f5d
    Reviewed-on: https://go-review.googlesource.com/29078
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 36b32911bdfa3c5323f9afa63047844cb5b0d0be
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Sep 9 22:24:44 2016 -0700

    cmd/compile: update fmt.go internal documentation
    
    No code changes.
    
    Change-Id: I7a22b3fbd6d727b276c7559f064cb0fdf385c02b
    Reviewed-on: https://go-review.googlesource.com/28955
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6537e18f02448bba92d5ea9c2f224b8c2c7ef888
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Sep 9 21:08:46 2016 -0700

    cmd/compile: rewrite %1v and %2v formats to %S and %L (short and long)
    
    - also consistently use %v instead of %s when we have a (gc) Formatter
    - rewrite done automatically using Formats test in -u (update) mode
    - manual update of format strings that were not single string constants
    - updated fmt.go, fmt_test.go accordingly
    - fmt_test: permit "%T" always
    
    Change-Id: I8f0704286aba5704600ad0c4a4484005b79b905d
    Reviewed-on: https://go-review.googlesource.com/28954
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit af8ca3387a62f83f9ee740d62c0007273e5fe1af
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Sep 9 19:54:09 2016 -0700

    cmd/compile: improved format string handling, faster operation
    
    - only accept a-z, A-Z as format verbs
    - blacklist vendored math package (no need to include it)
    
    Change-Id: Ica0fcbfe712369f79dd1d3472dfd4759b8bc3752
    Reviewed-on: https://go-review.googlesource.com/28953
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8458a387e3a0aa29449767f633514747a1ee7e72
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Sep 8 21:39:33 2016 -0700

    cmd/compile: make fmt_test work on entire compiler
    
    - process all directories recursively
    
    Change-Id: I27a737013d17fd3c2cc8ae9de4722dcbe989e6e5
    Reviewed-on: https://go-review.googlesource.com/28789
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8ff4260777aabe4ec7a92cba8c7dcce24f7fbf2b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Aug 30 09:12:22 2016 -0400

    cmd/compile: intrinsify Ctz, Bswap on ARM
    
    Atomic ops on ARM are implemented with kernel calls, so they are
    not intrinsified.
    
    Change-Id: I0e7cc2e5526ae1a3d24b4b89be1bd13db071f8ef
    Reviewed-on: https://go-review.googlesource.com/28977
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit cfea26026bc49be1710ed742465514e84bd31ab5
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 9 05:26:20 2016 +0000

    doc: document minimum OS X version as 10.8
    
    The download page says "OS X 10.8 or later", but other pages said 10.7.
    Say 10.8 everywhere.
    
    Turns out Go doesn't even compile on OS X 10.7 (details in bug) and we
    only run builders for OS X 10.8+, which is likely why 10.7
    regressed. Until recently we only had OS X 10.10 builders, even.
    
    We could run 10.7 builders, but there's basically no reason to do so,
    especially with 10.12 coming out imminently.
    
    Fixes #16625
    
    Change-Id: Ida6e20fb6c54aea0a3757235b708ac1c053b8c04
    Reviewed-on: https://go-review.googlesource.com/28870
    Reviewed-by: Chris Broadfoot <cbro@golang.org>

commit e4691d92dac9c46ea9ae380d5110dbceff45fad3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Sep 8 23:58:17 2016 +0000

    net/http: skip test needing good DNS in short mode, except on builders
    
    Fixes #16732
    
    Change-Id: If0a7f9425cf75b9e31b3091c43cb23d6e039f568
    Reviewed-on: https://go-review.googlesource.com/28782
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit d185cc3b643abbd88efcc1d9b92c87422a25d955
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 12 13:27:31 2016 -0400

    cmd/link: disable internal PIE for now
    
    There's more work to do.
    
    Updates #17068
    
    Change-Id: I4e16c0e8e9ac739e1fe266224c3769f6c5b2e070
    Reviewed-on: https://go-review.googlesource.com/29076
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f1515a01fd5d77b964194d3830d36ae006823ea3
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Sep 12 13:33:00 2016 -0400

    runtime, math/big: allow R0 on s390x to contain values other than 0
    
    The new SSA backend for s390x can use R0 as a general purpose register.
    This change modifies assembly code to either avoid using R0 entirely
    or explicitly set R0 to 0.
    
    R0 can still be safely used as 0 in address calculations.
    
    Change-Id: I3efa723e9ef322a91a408bd8c31768d7858526c8
    Reviewed-on: https://go-review.googlesource.com/28976
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 43bdfa9337c136f4e19122914c082f34045d9509
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Sep 12 12:13:34 2016 -0400

    cmd/asm, cmd/internal/obj/s390x: add new s390x instructions for SSA
    
    This commit adds the following instructions to support the new SSA
    backend for s390x:
    
    32-bit operations:
    ADDW
    SUBW
    NEGW
    FNEGS
    
    Conditional moves:
    MOVDEQ
    MOVDGE
    MOVDGT
    MOVDLE
    MOVDLT
    MOVDNE
    
    Unordered branches (for floating point comparisons):
    BLEU
    BLTU
    
    Modulo operations:
    MODW
    MODWU
    MODD
    MODDU
    
    The modulo operations might be removed in a future commit because
    I'd like to change DIV to produce a tuple once the old backend is
    removed.
    
    This commit also removes uses of REGZERO from the assembler. They
    aren't necessary and R0 will be used as a GPR by SSA.
    
    Change-Id: I05756c1cbb74bf4a35fc492f8f0cd34b50763dc9
    Reviewed-on: https://go-review.googlesource.com/29075
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit dfc56a4cd313c9c5de37f4fadb14912286edc42f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu May 12 17:22:47 2016 -0700

    cmd/compile: statically initialize some interface values
    
    When possible, emit static data rather than
    init functions for interface values.
    
    This:
    
    * cuts 32k off cmd/go
    * removes several error values from runtime init
    * cuts the size of the image/color/palette compiled package from 103k to 34k
    * reduces the time to build the package in #15520 from 8s to 1.5s
    
    Fixes #6289
    Fixes #15528
    
    Change-Id: I317112da17aadb180c958ea328ab380f83e640b4
    Reviewed-on: https://go-review.googlesource.com/26668
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b8eb5b53eadd10ed8c84b94a46301a3fc3715282
Author: David du Colombier <0intro@gmail.com>
Date:   Sun Sep 11 23:29:25 2016 +0200

    net: make lookupPort case-insensitive on Plan 9
    
    The CL 28951 added TestLookupPort_Minimal, which was failing
    on Plan 9, because lookupPort was not case-insensitive.
    
    Change-Id: Ic80dd29dad4ffd1c84c2590e3d5d0e588ab2e6c2
    Reviewed-on: https://go-review.googlesource.com/29051
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 62ba15a492a5a26fce70aab5f44cac4a841b956e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Sep 11 08:29:04 2016 -0700

    cmd/compile: add more non-returning runtime calls
    
    This list now matches the one in popt.go.
    
    Change-Id: Ib24de531cc35252f0ef276e5c6d247654b021533
    Reviewed-on: https://go-review.googlesource.com/28965
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 8086e7c6ab85646cd53d5cb6d6183750a76e6214
Author: Rob Pike <r@golang.org>
Date:   Mon Sep 12 13:17:02 2016 +1000

    testing: improve the documentation for the -run flag
    
    It's not intuitive, especially in the presence of subtests, so improve the
    explanation and extend and explain the examples.
    
    Change-Id: I6c4d3f8944b60b12311d0c0f0a8e952e7c35a9ed
    Reviewed-on: https://go-review.googlesource.com/28995
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 9e876861017991fe1478421b85a83964ea7abc6d
Author: Dave Day <djd@golang.org>
Date:   Fri Sep 9 17:59:43 2016 +1000

    net/url: modernise parse and unit tests
    
    Remove the naked returns and goto statements from parse.
    
    Make tests more consistent in the got/want ordering, and clean up some
    unnecessary helper functions.
    
    Change-Id: Iaa244cb8c00dd6b42836d95448bf02caa72bfabd
    Reviewed-on: https://go-review.googlesource.com/28890
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4ba2a4913c4e29754ade9d7329ff324a8f5de59a
Author: Michal Bohuslávek <mbohuslavek@gmail.com>
Date:   Tue Sep 6 10:49:26 2016 +0100

    crypto/rsa: remove unused variable y
    
    Change-Id: I70beb844cd6928dbfbfd8de365e0cb708e54f71e
    Reviewed-on: https://go-review.googlesource.com/28496
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit 5a59b66f23bc2eb11cc8445aea1dcf7a71bf2954
Author: Filippo Valsorda <hi@filippo.io>
Date:   Fri Sep 9 14:07:30 2016 +0100

    crypto/tls: flush the buffer on handshake errors
    
    Since 2a8c81ff handshake messages are not written directly to wire but
    buffered.  If an error happens at the wrong time the alert will be
    written to the buffer but never flushed, causing an EOF on the client
    instead of a more descriptive alert.
    
    Thanks to Brendan McMillion for reporting this.
    
    Fixes #17037
    
    Change-Id: Ie093648aa3f754f4bc61c2e98c79962005dd6aa2
    Reviewed-on: https://go-review.googlesource.com/28818
    Reviewed-by: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2e4dc86bfb19b1eb2a69ac36c87ada22bacc98c9
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Sep 11 14:43:37 2016 -0700

    cmd/compile: add Node.IsMethod helper
    
    Changes generated with eg:
    
    func before(n *gc.Node) bool { return n.Type.Recv() != nil }
    func after(n *gc.Node) bool  { return n.IsMethod() }
    
    func before(n *gc.Node) bool { return n.Type.Recv() == nil }
    func after(n *gc.Node) bool  { return !n.IsMethod() }
    
    Change-Id: I28e544490d17bbdc06ab11ed32464af5802ab206
    Reviewed-on: https://go-review.googlesource.com/28968
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3bf141955be1d1f791f643700fbde60c258546d3
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 6 09:16:48 2016 -0400

    cmd/dist: test PIE internal linking on linux/amd64
    
    Change-Id: I88dd0811db3a9864106def47b89848f5c8de94d4
    Reviewed-on: https://go-review.googlesource.com/28545
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 675ba53c7671d5506887eec2a65fabb4cbda0d59
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 6 09:14:26 2016 -0400

    cmd/go: internal PIE does not need runtime/cgo
    
    Part of adding PIE internal linking on linux/amd64.
    
    Change-Id: I57f0596cb254cbe6569e4d4e39fe4f48437733f2
    Reviewed-on: https://go-review.googlesource.com/28544
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 1a42d8fbd6fdce1ef9122c2c4c60833eee9cefa0
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 6 08:08:47 2016 -0400

    cmd/link: allow internal PIE linking
    
    Part of adding PIE internal linking on linux/amd64.
    
    Change-Id: I5ce01d1974e5d4b1a8cbcc8b08157477631d8d24
    Reviewed-on: https://go-review.googlesource.com/28543
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 04b4dbe1f0ec58efe8a1bf8e05a1042b17176c3b
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 6 08:05:19 2016 -0400

    cmd/link: mark PIE binaries as ET_DYN
    
    Part of adding PIE internal linking on linux/amd64.
    
    Change-Id: I586e7c2afba349281168df5e20d2fdcb697f6e37
    Reviewed-on: https://go-review.googlesource.com/28542
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d9d1399eeac037f4e2057d32831e574c0f1e6fc8
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 6 08:02:30 2016 -0400

    cmd/link: mark rel.ro segment as PT_GNU_RELRO
    
    Details: http://www.airs.com/blog/archives/189
    
    Part of adding PIE internal linking on linux/amd64.
    
    Change-Id: I8843a97f22f6f120346cccd694c7fff32f09f60b
    Reviewed-on: https://go-review.googlesource.com/28541
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 44ee2b00db9ae21c4304dbfbd19b4bc70f9ff4c9
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 6 07:46:59 2016 -0400

    cmd/link: optimize TLS IE to LE in build mode PIE
    
    When cmd/compile generates position-independent code on linux
    (the -shared flag), it refers to runtime.tlsg as a TLS IE variable.
    
    When cmd/link is linking a PIE executable internally, all TLS IE
    relocations are generated by cmd/compile, and the variable they
    refer to, runtime.tlsg, is local to the binary. This means we can
    optimize this particular IE case to LE, and thus implement IE
    support when internally linking.
    
    To do this optimization in the linker, we need to rewrite the
    PC-relative MOVD to a constant load. This may seem like an
    unconscionable act born of enthusiasm, but it turns out this is
    standard operating procedure for linkers. GNU gold does exactly
    the same optimization. I spent some time reading it and documented
    at least one missing feature from this version.
    
    Part of adding PIE internal linking on linux/amd64.
    
    Change-Id: I1eb068d0ec774724944c6b308aa5084582ecde0b
    Reviewed-on: https://go-review.googlesource.com/28540
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1fe4c81282f22b5ac9ba25e7972109255e173b04
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Sep 11 12:16:11 2016 -0700

    cmd/compile: don't crash on complex(0())
    
    Fixes #17038.
    
    Change-Id: Iaf6294361050040830af1d60cd48f263223d9356
    Reviewed-on: https://go-review.googlesource.com/28966
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit aada57f39b60c2aef88dbafc7b406df9c2680f12
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Sep 11 12:34:58 2016 -0700

    misc/cgo/test: add skipped test for issue 17065
    
    Updates #17065
    
    Change-Id: I113caced6de666a9b032ab2684ece79482aa7357
    Reviewed-on: https://go-review.googlesource.com/28964
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 81ee93679dad4138d1c3431aa66237a31a99f1b2
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 5 23:49:53 2016 -0400

    cmd/link: generate dynamic relocs for internal PIE
    
    This reuses the machinery built for dynamic loading of shared
    libraries. The significant difference with PIE is we generate
    dynamic relocations for known internal symbols, not just
    dynamic external symbols.
    
    Part of adding PIE internal linking on linux/amd64.
    
    Change-Id: I4afa24070bfb61f94f8d3648f2433d5343bac3fe
    Reviewed-on: https://go-review.googlesource.com/28539
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 276803d6111b46c66956c99d982d70f23820ba5d
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Sep 5 23:29:16 2016 -0400

    cmd/link: introduce a rel.ro segment
    
    When internally linking with using rel.ro sections, this segment covers
    the sections. To do this, move to other read-only sections, SELFROSECT
    and SMACHOPLT, out of the way.
    
    Part of adding PIE internal linking on linux/amd64.
    
    Change-Id: I4fb3d180e92f7e801789ab89864010faf5a2cb6d
    Reviewed-on: https://go-review.googlesource.com/28538
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 13dc4d378bb40261fb2e7f9e1f0eb2f840250040
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Sep 11 01:32:22 2016 +0000

    vendor: add golang.org/x/net/idna to the vendor directory for Punycode
    
    Adds golang.org/x/net/idna to the Go repo from the
    golang.org/x/net repo's git rev 7db922ba (Dec 2012).
    
    Punycode is needed for http.Get("привет.рф") etc., which will
    come in separate commits.
    
    Updates #13835
    
    Change-Id: I313ed82d292737579a3ec5dcf8a9e766f2f75138
    Reviewed-on: https://go-review.googlesource.com/28961
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 74190735be5e722a0a977d056204307c2468cf21
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 9 22:51:11 2016 +0000

    net: make LookupPort and lookupProtocol work on nacl
    
    Also, flesh out the baked-in /etc/services table for LookupPort a bit.
    
    This services map moves from a unix-specific file to a portable file
    where nacl can use it.
    
    Also, remove the duplicated entries in the protocol map in different
    cases, and just canonicalize the input before looking in the map. Now
    it handles any case, including MiXeD cAse.
    
    In the process, add a test that service names for LookupPort are case
    insensitive. They were on Windows, but not cgo. Now there's a test and
    they're case insensitive in all 3+ paths. Maybe it breaks plan9. We'll
    see.
    
    Fixes #17045
    
    Change-Id: Idce7d68703f371727c7505cda03a32bd842298cd
    Reviewed-on: https://go-review.googlesource.com/28951
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit f7fce1e2088ed888c5ced3f2649af84bc0a672b9
Author: Shenghou Ma <minux@golang.org>
Date:   Mon Aug 22 17:20:57 2016 -0400

    mime/quotedprintable: accept trailing soft line-break at the end of message
    
    Fixes #15486.
    
    Change-Id: Id879dc9acef9232003df9a0f6f54312191374a60
    Reviewed-on: https://go-review.googlesource.com/27530
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fca3dd3718080563f4bc6c4c8b6fbe681a1602fa
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sat Sep 10 14:04:46 2016 +1000

    syscall: avoid convT2I allocs for ERROR_IO_PENDING instead of WSAEINPROGRESS
    
    CL 28484 mistakenly assumed that WSARecv returns WSAEINPROGRESS
    when there is nothing to read. But the error is ERROR_IO_PENDING.
    Fix that mistake.
    
    I was about to write a test for it. But I have found
    TestTCPReadWriteAllocs in net package that does nearly what I need,
    but was conveniently disabled. So enable and extend the test.
    
    Fixes #16988
    
    Change-Id: I55e5cf8998a9cf29e92b398d702280bdf7d6fc85
    Reviewed-on: https://go-review.googlesource.com/28990
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f15915af4effbbe6895ae69be02d22ac016927d5
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Sep 10 15:21:12 2016 -0700

    cmd/vet: ignore printf failures in cmd
    
    This is a temporary measure to work around #17057.
    It will be reverted when #17057 is fixed.
    
    Change-Id: I21c02f63f3530774c91065cfed5d9c566839ed9f
    Reviewed-on: https://go-review.googlesource.com/28959
    Reviewed-by: Rob Pike <r@golang.org>

commit 8e84a5dd47d5c27e83f40439a4240a7504134039
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Sep 10 15:24:28 2016 -0700

    cmd/vet/all/whitelist: add readme.txt
    
    This was written for CL 27811,
    but it got lost in the sea of new files.
    
    Change-Id: I5c52cb23dda499b21a6bb32ed5c586779ccbc2f1
    Reviewed-on: https://go-review.googlesource.com/28960
    Reviewed-by: Rob Pike <r@golang.org>

commit 5de2e5b73e79aea02463d661ec87cee4e944990b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jul 9 15:29:32 2016 -0700

    cmd/vet: add vet runner script for core
    
    This CL adds a script to run vet on std and cmd.
    
    There are a considerable number of false positives,
    mostly from legitimate but unusual assembly in
    the runtime and reflect packages.
    
    There are also a few false positives that need fixes.
    They are noted as such in the whitelists;
    they're not worth holding this CL for.
    
    The unsafe pointer check is disabled.
    The false positive rate is just too high to be worth it.
    
    There is no cmd/dist/test integration yet.
    The tentative plan is that we'll check the local platform
    during all.bash, and that there'll be a fast builder that
    checks all platforms (to cover platforms that can't exec).
    
    Fixes #11041
    
    Change-Id: Iee4ed32b05447888368ed86088e3ed3771f84442
    Reviewed-on: https://go-review.googlesource.com/27811
    Reviewed-by: Rob Pike <r@golang.org>

commit 4cf95fda64d76ca044319fd0a292ad3a77c1da0b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Sep 10 14:05:51 2016 -0700

    encoding/hex: fix example function name
    
    Found by vet.
    
    Change-Id: I556d87f853a734002f779b04ba5a3588a3117106
    Reviewed-on: https://go-review.googlesource.com/28958
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0e435347b1ed888f51706d5467fabb829292a5ff
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Sep 10 14:05:31 2016 -0700

    cmd: fix format strings used with obj.Headtype
    
    Found by vet. Introduced by CL 28853.
    
    Change-Id: I3199e0cbdb1c512ba29eb7e4d5c1c98963f5a954
    Reviewed-on: https://go-review.googlesource.com/28957
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f30598dd71b5bb7109eda56b02ea10d2b6cc1362
Author: Joonas Kuorilehto <joneskoo@derbian.fi>
Date:   Sat Sep 10 22:07:33 2016 +0300

    crypto/tls: Add mutex to protect KeyLogWriter
    
    Concurrent use of tls.Config is allowed, and may lead to
    KeyLogWriter being written to concurrently. Without a mutex
    to protect it, corrupted output may occur. A mutex is added
    for correctness.
    
    The mutex is made global to save size of the config struct as
    KeyLogWriter is rarely enabled.
    
    Related to #13057.
    
    Change-Id: I5ee55b6d8b43a191ec21f06e2aaae5002a71daef
    Reviewed-on: https://go-review.googlesource.com/29016
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c564aebce99fb92b8dc26b203f4f32e4977c0aed
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Sep 8 18:50:59 2016 -0700

    internal/testenv: add GoTool
    
    GoToolPath requires a *testing.T to handle errors.
    GoTool provides a variant that returns errors
    for clients without a *testing.T,
    such as that found in CL 27811.
    
    Change-Id: I7ac8b7ec9d472894c37223c5f7b121ec823e7f61
    Reviewed-on: https://go-review.googlesource.com/28787
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3d562de8e38317a06fbf860e60f71c5e4f0ceab6
Author: bogem <albertnigma@gmail.com>
Date:   Sat Sep 10 18:34:07 2016 +0500

    os: delete code duplications in error files
    
    Change-Id: I1ec2fcf81a7a9e45a2fae8c02c8adabc7841b4fa
    Reviewed-on: https://go-review.googlesource.com/29013
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ca993d679729dbf312530d15f0489a45d1b61eaa
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Sep 8 09:12:50 2016 -0700

    builtin: clarify that make(map[K]V, n) allocates space for n elements
    
    Change-Id: Id6265b6093edaa4be2c59e4799351082f7228b5d
    Reviewed-on: https://go-review.googlesource.com/28815
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit e273a2794639692a216c46a0043e5e6929e0a959
Author: bogem <albertnigma@gmail.com>
Date:   Fri Sep 9 23:23:32 2016 +0500

    flag: use strconv instead of fmt in values' String funcs
    
    The existing implementation of flag values with fmt package uses
    more memory and works slower than the implementation with strconv
    package.
    
    Change-Id: I9e749179f66d5c50cafe98186641bcdbc546d2db
    Reviewed-on: https://go-review.googlesource.com/28914
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7b269195d30195f29080e17114aeec7821851870
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Thu Sep 8 11:21:20 2016 -0700

    cmd/go: organize test profiling flags into their own sub-section
    
    Fixes #17020.
    
    Change-Id: Ice21bd8eb4dbc208f244b275c3be604bc8e3efe7
    Reviewed-on: https://go-review.googlesource.com/28783
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Jaana Burcu Dogan <jbd@google.com>

commit 81dfcba331f43bd14c8933eca83c433e53cb7b55
Author: Edward Muller <edwardam@interlix.com>
Date:   Wed Sep 7 11:39:31 2016 -0700

    go/build: add help info for unset $GOPATH
    
    We relay this info in a few places, in a few different ways, but not
    consistently everywhere.  This led one of our users to start googling
    and not find https://golang.org/doc/code.html#Workspaces, of which `go
    help gopath` is the most equivalent.
    
    Change-Id: I28a94375739f3aa4f200e145293ca2a5f65101e1
    Reviewed-on: https://go-review.googlesource.com/28690
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 1ff19201fd898c3e1a0ed5d3458c81c1f062570b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 9 19:29:25 2016 +0000

    net/url: add URL.Hostname and URL.Port accessors
    
    Fixes #16142
    
    Change-Id: I7609faaf00c69646b0bd44a60a63a22d9265feb0
    Reviewed-on: https://go-review.googlesource.com/28933
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Francesc Campoy Flores <campoy@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6e87082d41f0267b39e6a1854d655b1d1c2f7541
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 9 18:06:56 2016 +0000

    net/http: make Client copy headers on redirect
    
    Copy all of the original request's headers on redirect, unless they're
    sensitive. Only send sensitive ones to the same origin, or subdomains
    thereof.
    
    Fixes #4800
    
    Change-Id: Ie9fa75265c9d5e4c1012c028d31fd1fd74465712
    Reviewed-on: https://go-review.googlesource.com/28930
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Francesc Campoy Flores <campoy@golang.org>
    Reviewed-by: Ross Light <light@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1161a19c1ef536f8db2ca7eddf0e424e138e37db
Author: Carlos C <uldericofilho@gmail.com>
Date:   Fri Aug 12 12:45:14 2016 +0200

    context: add examples
    
    Add function level examples to the package.
    
    Partially addresses #16360
    
    Change-Id: I7162aed4e4a969743c19b79c9ffaf9217d2c1c08
    Reviewed-on: https://go-review.googlesource.com/26930
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d817c4ec6d458777537aae3a04201b7182a5f78a
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Sep 9 11:29:33 2016 -0700

    cmd/compile: use regular rather than indexed format string
    
    This enables the format test to process this file (the format
    test doesn't handle indexed formats, and this is the only place
    in the compiler where they occur).
    
    Change-Id: I99743f20c463f181a589b210365f70162227d4e0
    Reviewed-on: https://go-review.googlesource.com/28932
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit cfc0d6d8847eb959da2b6bd1c0fe1c0c7a19873e
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Fri Sep 9 11:28:07 2016 -0700

    cmd/compile/internal/syntax: remove strbyteseql
    
    cmd/compile already optimizes "string(b) == s" to avoid allocating a
    temporary string.
    
    Change-Id: I4244fbeae8d350261494135c357f9a6e2ab7ace3
    Reviewed-on: https://go-review.googlesource.com/28931
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3877f820a61886b74bb596bdb128a8d705a44628
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed Sep 7 14:45:27 2016 -0400

    cmd/link, etc: introduce SymKind type
    
    Moves the grouping of symbol kinds (sections) into cmd/internal/obj
    to keep it near the definition. Groundwork for CL 28538.
    
    Change-Id: I99112981e69b028f366e1333f31cd7defd4ff82c
    Reviewed-on: https://go-review.googlesource.com/28691
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 791f71d1921a5d4d775167486eab20c1e8f97248
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 9 08:13:16 2016 -0400

    cmd: use obj.GOOS, obj.GOARCH, etc
    
    As cmd/internal/obj is coordinating the definition of GOOS, GOARCH,
    etc across the compiler and linker, turn its functions into globals
    and use them everywhere.
    
    Change-Id: I5db5addda3c6b6435c37fd5581c7c3d9a561f492
    Reviewed-on: https://go-review.googlesource.com/28854
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f2f374d125312cd6f9c91581e85a80ee1f143cc1
Author: Tormod Erevik Lea <tormodlea@gmail.com>
Date:   Fri Sep 9 16:41:30 2016 +0200

    reflect: update location of vet tool in comment
    
    Change-Id: Ic5160edbbca4a8ffc7c7e6246e34fae1978470fd
    Reviewed-on: https://go-review.googlesource.com/28912
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 428c349552759dc1e11965d1d483a4749818ec11
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Sep 9 06:20:44 2016 -0400

    cmd/link, cmd/internal/obj: give Headtype a type
    
    Separate out windows/windowsgui properly so we are not encoding some
    of the Headtype value in a separate headstring global.
    
    Remove one of the two copies of the variable from cmd/link.
    
    Remove duplicate string to headtype list.
    
    Change-Id: Ifa20fb9652a1dc95161e154aac11f15ad0f709d0
    Reviewed-on: https://go-review.googlesource.com/28853
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 361d2738d5d750bd249661e725d3d9070fc3f2f5
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Sep 8 22:25:50 2016 -0400

    cmd/link: remove the -shared flag
    
    The -shared flag has been superseded by the -buildmode flag.
    
    Change-Id: I3682cc0367b919084c280d7dc64746485c1d4ddd
    Reviewed-on: https://go-review.googlesource.com/28852
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 57d4e5763559634733452f62aba562a5df345192
Author: Dhaivat Pandit <dhaivatpandit@gmail.com>
Date:   Fri Aug 26 20:13:44 2016 -0700

    net/http/cookiejar: added simple example test
    
    Fixes #16884
    Updates #16360
    
    Change-Id: I01563031a1c105e54499134eed4789f6219f41ec
    Reviewed-on: https://go-review.googlesource.com/27993
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 39382793d3a9e7a0720e6dbf8be4b19e8519af19
Author: Jack Lindamood <jlindamo@justin.tv>
Date:   Tue Jul 26 14:20:36 2016 -0700

    context: reduce memory usage of context tree
    
    Modifies context package to use map[]struct{} rather than map[]bool,
    since the map is intended as a set object.  Also adds Benchmarks to
    the context package switching between different types of root nodes
    and a tree with different depths.
    
    Included below are bytes deltas between the old and new code, using
    these benchmarks.
    
    benchmark                                                       old bytes     new bytes     delta
    BenchmarkContextCancelTree/depth=1/Root=Background-8            176           176           +0.00%
    BenchmarkContextCancelTree/depth=1/Root=OpenCanceler-8          560           544           -2.86%
    BenchmarkContextCancelTree/depth=1/Root=ClosedCanceler-8        352           352           +0.00%
    BenchmarkContextCancelTree/depth=10/Root=Background-8           3632          3488          -3.96%
    BenchmarkContextCancelTree/depth=10/Root=OpenCanceler-8         4016          3856          -3.98%
    BenchmarkContextCancelTree/depth=10/Root=ClosedCanceler-8       1936          1936          +0.00%
    BenchmarkContextCancelTree/depth=100/Root=Background-8          38192         36608         -4.15%
    BenchmarkContextCancelTree/depth=100/Root=OpenCanceler-8        38576         36976         -4.15%
    BenchmarkContextCancelTree/depth=100/Root=ClosedCanceler-8      17776         17776         +0.00%
    BenchmarkContextCancelTree/depth=1000/Root=Background-8         383792        367808        -4.16%
    BenchmarkContextCancelTree/depth=1000/Root=OpenCanceler-8       384176        368176        -4.16%
    BenchmarkContextCancelTree/depth=1000/Root=ClosedCanceler-8     176176        176176        +0.00%
    
    Change-Id: I699ad704d9f7b461214e1651d24941927315b525
    Reviewed-on: https://go-review.googlesource.com/25270
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 0ab6bb42e197c4c766769bfc4a8807f93cc630b9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 9 01:30:49 2016 +0000

    doc: clarify that any C compiler works for cgo
    
    Currently the footnote says "gcc is required only if you plan to use cgo",
    but the footnote was referenced from the text:
    
       "use the clang or gcc† that comes with Xcode‡ for cgo support"
    
    That seems to imply that clang doesn't get you cgo support on OS X,
    which isn't true. The update text matches what the install-source.html
    page says.
    
    Change-Id: Ib88464a0d138227d357033123f6675a77d5d777f
    Reviewed-on: https://go-review.googlesource.com/28786
    Reviewed-by: Minux Ma <minux@golang.org>

commit f782a7e075d049979c67cdc72b6fb38fac36fdad
Author: Nigel Tao <nigeltao@golang.org>
Date:   Thu Jul 14 10:51:30 2016 +1000

    image/draw: optimize drawFillOver as drawFillSrc for opaque fills.
    
    Benchmarks are much better for opaque fills and slightly worse on non
    opaque fills. I think that on balance, this is still a win.
    
    When the source is uniform(color.RGBA{0x11, 0x22, 0x33, 0xff}):
    name        old time/op  new time/op  delta
    FillOver-8   966µs ± 1%    32µs ± 1%  -96.67%  (p=0.000 n=10+10)
    FillSrc-8   32.4µs ± 1%  32.2µs ± 1%     ~      (p=0.053 n=9+10)
    
    When the source is uniform(color.RGBA{0x11, 0x22, 0x33, 0x44}):
    name        old time/op  new time/op  delta
    FillOver-8   962µs ± 0%  1018µs ± 0%  +5.85%   (p=0.000 n=9+10)
    FillSrc-8   32.2µs ± 1%  32.1µs ± 0%    ~     (p=0.148 n=10+10)
    
    Change-Id: I52ec6d5fcd0fbc6710cef0e973a21ee7827c0dd9
    Reviewed-on: https://go-review.googlesource.com/28790
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 06eeea210b0fd764cda8d86b555343fcfac6e194
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Sep 8 16:51:26 2016 -0700

    cmd/compile: permit use of %S (short) and %L (long) instead of %1v and %2v
    
    First step towards cleaning up format use. Not yet enabled.
    
    Change-Id: Ia8d76bf02fe05882fffb9d17c9a30dc38d28bf81
    Reviewed-on: https://go-review.googlesource.com/28784
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d38d59ffb5d56c838b1ed7cc346b8d63398b5452
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Sep 8 19:27:24 2016 -0400

    runtime: fix SIGILL in checkvectorfacility on s390x
    
    STFLE does not necessarily write to all the double-words that are
    requested. It is therefore necessary to clear the target memory
    before calling STFLE in order to ensure that the facility list does
    not contain false positives.
    
    Fixes #17032.
    
    Change-Id: I7bec9ade7103e747b72f08562fe57e6f091bd89f
    Reviewed-on: https://go-review.googlesource.com/28850
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 14ff7cc94c4e167dbd80c96996b43f96bb9c17fb
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu Sep 8 08:33:19 2016 -0400

    runtime/cgo: fix callback on big-endian MIPS64
    
    Use MOVW, instead of MOVV, to pass an int32 arg. Also no need to
    restore arg registers.
    
    Fix big-endian MIPS64 build.
    
    Change-Id: Ib43c71075c988153e5e5c5c6e7297b3fee28652a
    Reviewed-on: https://go-review.googlesource.com/28830
    Reviewed-by: Minux Ma <minux@golang.org>

commit 2a36f78e8795c939630e494319c5790bc519ba72
Author: Jim Kingdon <jim@bolt.me>
Date:   Thu Sep 8 11:27:04 2016 -0700

    doc: avoid mentioning non-existence of u flag to fmt.Printf.
    
    It is better to document what golang does, rather than how it differs
    from languages which readers may or may not know.
    
    That the output format is based on the type is basically self-evident
    if you consider this just in go terms.
    
    Change-Id: I0223e9b4cb67cc83a9ebe4d424e6c151d7ed600f
    Reviewed-on: https://go-review.googlesource.com/28393
    Reviewed-by: Rob Pike <r@golang.org>

commit f27c1bda5165f94115458908b5222d992010cbee
Author: Dave Day <djd@golang.org>
Date:   Thu Sep 1 13:13:37 2016 +1000

    net/url: handle escaped paths in ResolveReference
    
    Currently, path resolution is done using the un-escaped version of
    paths. This means that path elements like one%2ftwo%2fthree are
    handled incorrectly, and optional encodings (%2d vs. -) are dropped.
    
    This function makes escaped handling consistent with Parse: provided
    escapings are honoured, and RawPath is only set if necessary.
    
    A helper method setPath is introduced to handle the correct setting of
    Path and RawPath given the encoded path.
    
    Fixes #16947
    
    Change-Id: I40b1215e9066e88ec868b41635066eee220fde37
    Reviewed-on: https://go-review.googlesource.com/28343
    Run-TryBot: Dave Day <djd@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 29f18d7983d9ae3752a36cf22f4c83601f6da9ea
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Sep 8 14:50:24 2016 -0700

    cmd/compile: fix bug in oconv
    
    Introduced by https://go-review.googlesource.com/#/c/28331/ .
    
    Change-Id: Id75aed6410f06b302d5347f6ca6a2e19c61f6fb6
    Reviewed-on: https://go-review.googlesource.com/28779
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 9d6677ca7bd5de65b94046dbe90b4ef8418502b0
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Sep 8 14:39:38 2016 -0700

    cmd/compile: update and reenable Formats test
    
    Change-Id: I9c0da13d21551dbf766156472224370ab9badfe9
    Reviewed-on: https://go-review.googlesource.com/28778
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 58de26810ccfd20d3c430a20483b171938337aed
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 17:49:49 2016 -0700

    cmd/compile: remove superfluous returns in fmt.go
    
    Change-Id: Ie73fb460a3838c6d1b9348965a8b69c1bfa6a882
    Reviewed-on: https://go-review.googlesource.com/28341
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit 0ca5f269d4cb04cccd0012cc6fc0cb52b0002728
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 16:28:18 2016 -0700

    cmd/compile: remove fmt.go printer again, now that it's unused
    
    Change-Id: I9a6e5b9cbcfc264c61fd39ed65330ca737707e1f
    Reviewed-on: https://go-review.googlesource.com/28340
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 266c6223df73c5f90d2da6e271c8a415392e066c
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 16:19:50 2016 -0700

    cmd/compile: implement fmt.Formatter for Nodes formats %s, %v
    
    Change-Id: Iac3a72cb6c5394f3c1a49f39125b0256d570e006
    Reviewed-on: https://go-review.googlesource.com/28339
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit daf61797330a340c1341bd5428df8f623135b81b
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 16:07:49 2016 -0700

    cmd/compile: use fmt.State in nodedump
    
    Change-Id: Icd83e88fc879b30b34f8697d540619efeb25c25b
    Reviewed-on: https://go-review.googlesource.com/28338
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 0510f0dfe72ff9488c9c9aa23dbbc700837acfba
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 16:02:36 2016 -0700

    cmd/compile: use fmt.State in exprfmt
    
    Change-Id: If6c2d469c66a7aa8471bf54310354efdac3e0235
    Reviewed-on: https://go-review.googlesource.com/28337
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ebdc8faf0411d04fa4e2e47d1c7861682141f866
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 15:32:35 2016 -0700

    cmd/compile: use fmt.State in nodefmt and stmtfmt
    
    Change-Id: Iac87007af4ee268b45f11ec05bf4757f2e7eedd8
    Reviewed-on: https://go-review.googlesource.com/28336
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ff046d2e28791a5db4802306a3d1324a9e58e383
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 15:22:36 2016 -0700

    cmd/compile: implement fmt.Formatter for *Node formats %s, %v
    
    Change-Id: I80ed668cdeab0c4342b734d34b429927e0213e5a
    Reviewed-on: https://go-review.googlesource.com/28335
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a0d2010208199d2cfca0581348efde4e04d0ab06
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 14:12:35 2016 -0700

    cmd/compile: implement fmt.Formatter for *Sym formats %s, %v
    
    Change-Id: I0c362edba66c763e84990e3c5508013021f3e6fe
    Reviewed-on: https://go-review.googlesource.com/28334
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 8d0bbe2b48fae8c41b990c5605d614e8f4b5e6d4
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Aug 31 10:32:40 2016 -0700

    cmd/compile: implement fmt.Formatter for *Type formats %s, %v
    
    Change-Id: I878ac549430abc7859c30d176d52d52ce02c5827
    Reviewed-on: https://go-review.googlesource.com/28333
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 97ba3c824265bd6917c13c1ebc59be3085324b4c
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 30 15:01:48 2016 -0700

    cmd/compile: implement fmt.Formatter for Val formats %s, %v
    
    Change-Id: Id56e886793161b48b445439e9a12109142064d3f
    Reviewed-on: https://go-review.googlesource.com/28332
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 9e1f6a326e4fe9081975d73da03d115c01173552
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 30 14:21:33 2016 -0700

    cmd/compile: implement fmt.Formatter for Op formats %s, %v
    
    Change-Id: I59e18fab37fd688fc1e578e2192e32e29fdf37f0
    Reviewed-on: https://go-review.googlesource.com/28331
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit adcd34c732d01d60ca336627fc0fd7647d46a0c1
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Aug 30 14:13:41 2016 -0700

    cmd/compile: implement fmt.Formatter for *Node formats %s, %v, %j
    
    Change-Id: I44ee5843bb9dfd65b9a18091f365355e84888f21
    Reviewed-on: https://go-review.googlesource.com/28330
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b6948ce7c2ee7ab7bbb2377a8b76d1473075753c
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Sep 8 14:12:28 2016 -0700

    cmd/compile: temporarily disable Formats test
    
    So we can submit a sequence of older changes that don't yet
    update the formats in this file. We will then re-enable the
    test with the updated formats.
    
    Change-Id: I6ed559b83adc891bbf4b3d855a7dc1e428366f7f
    Reviewed-on: https://go-review.googlesource.com/28776
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 141f1a0e24836abbc85933787025edef260fe1f0
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Sep 4 03:49:02 2016 -0700

    encoding/hex: implement examples using all exported functions
    
    Fixes #11254.
    Updates #16360.
    
    Implements examples using all exported functions.
    
    This CL also updates Decode documentation to
    state that only hexadecimal characters are accepted
    in the source slice src, but also that the length
    of src must be even.
    
    Change-Id: Id016a4ba814f940cd300f26581fb4b9d2aded306
    Reviewed-on: https://go-review.googlesource.com/28482
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 4354ffd38b7ebdf7b4ee9ff614939ed77f872acd
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Aug 29 16:26:57 2016 -0400

    cmd/compile: intrinsify Ctz, Bswap, and some atomics on ARM64
    
    Change-Id: Ia5bf72b70e6f6522d6fb8cd050e78f862d37b5ae
    Reviewed-on: https://go-review.googlesource.com/27936
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit daa7c607d269e4779b74174032639b552174868f
Author: Sina Siadat <siadat@gmail.com>
Date:   Thu Sep 8 11:39:12 2016 +0430

    net/http/httputil: remove custom hop-by-hop headers from response in ReverseProxy
    
    Hop-by-hop headers (explicitly mentioned in RFC 2616) were already
    removed from the response. This removes the custom hop-by-hop
    headers listed in the "Connection" header of the response.
    
    Updates #16875
    
    Change-Id: I6b8f261d38b8d72040722f3ded29755ef0303427
    Reviewed-on: https://go-review.googlesource.com/28810
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ed8f207940c8787d344664a43071b235e2ce5c68
Author: Kevin Burke <kev@inburke.com>
Date:   Sun Jun 26 09:47:43 2016 -0700

    encoding/json: Use a lookup table for safe characters
    
    The previous check for characters inside of a JSON string that needed
    to be escaped performed seven different boolean comparisons before
    determining that a ASCII character did not need to be escaped. Most
    characters do not need to be escaped, so this check can be done in a
    more performant way.
    
    Use the same strategy as the unicode package for precomputing a range
    of characters that need to be escaped, then do a single lookup into a
    character array to determine whether the character needs escaping.
    
    On an AWS c4.large node:
    
    $ benchstat benchmarks/master-bench benchmarks/json-table-bench
    name                   old time/op    new time/op     delta
    CodeEncoder-2            19.0ms ± 0%     15.5ms ± 1%  -18.16%        (p=0.000 n=19+20)
    CodeMarshal-2            20.1ms ± 1%     16.8ms ± 2%  -16.35%        (p=0.000 n=20+21)
    CodeDecoder-2            49.3ms ± 1%     49.5ms ± 2%     ~           (p=0.498 n=16+20)
    DecoderStream-2           416ns ± 0%      416ns ± 1%     ~           (p=0.978 n=19+19)
    CodeUnmarshal-2          51.0ms ± 1%     50.9ms ± 1%     ~           (p=0.490 n=19+17)
    CodeUnmarshalReuse-2     48.5ms ± 2%     48.5ms ± 2%     ~           (p=0.989 n=20+19)
    UnmarshalString-2         541ns ± 1%      532ns ± 1%   -1.75%        (p=0.000 n=20+21)
    UnmarshalFloat64-2        485ns ± 1%      481ns ± 1%   -0.92%        (p=0.000 n=20+21)
    UnmarshalInt64-2          429ns ± 1%      427ns ± 1%   -0.49%        (p=0.000 n=19+20)
    Issue10335-2              631ns ± 1%      619ns ± 1%   -1.84%        (p=0.000 n=20+20)
    NumberIsValid-2          19.1ns ± 0%     19.1ns ± 0%     ~     (all samples are equal)
    NumberIsValidRegexp-2     689ns ± 1%      690ns ± 0%     ~           (p=0.150 n=20+20)
    SkipValue-2              14.0ms ± 0%     14.0ms ± 0%   -0.05%        (p=0.000 n=18+18)
    EncoderEncode-2           525ns ± 2%      512ns ± 1%   -2.33%        (p=0.000 n=20+18)
    
    name                   old speed      new speed       delta
    CodeEncoder-2           102MB/s ± 0%    125MB/s ± 1%  +22.20%        (p=0.000 n=19+20)
    CodeMarshal-2          96.6MB/s ± 1%  115.6MB/s ± 2%  +19.56%        (p=0.000 n=20+21)
    CodeDecoder-2          39.3MB/s ± 1%   39.2MB/s ± 2%     ~           (p=0.464 n=16+20)
    CodeUnmarshal-2        38.1MB/s ± 1%   38.1MB/s ± 1%     ~           (p=0.525 n=19+17)
    SkipValue-2             143MB/s ± 0%    143MB/s ± 0%   +0.05%        (p=0.000 n=18+18)
    
    I also took the data set reported in #5683 (browser
    telemetry data from Mozilla), added named structs for
    the data set, and turned it into a proper benchmark:
    https://github.com/kevinburke/jsonbench/blob/master/go/bench_test.go
    
    The results from that test are similarly encouraging. On a 64-bit
    Mac:
    
    $ benchstat benchmarks/master-benchmark benchmarks/json-table-benchmark
    name              old time/op    new time/op    delta
    CodeMarshal-4       1.19ms ± 2%    1.08ms ± 2%   -9.33%  (p=0.000 n=21+17)
    Unmarshal-4         3.09ms ± 3%    3.06ms ± 1%   -0.83%  (p=0.027 n=22+17)
    UnmarshalReuse-4    3.04ms ± 1%    3.04ms ± 1%     ~     (p=0.169 n=20+15)
    
    name              old speed      new speed      delta
    CodeMarshal-4     80.3MB/s ± 1%  88.5MB/s ± 1%  +10.29%  (p=0.000 n=21+17)
    Unmarshal-4       31.0MB/s ± 2%  31.2MB/s ± 1%   +0.83%  (p=0.025 n=22+17)
    
    On the c4.large:
    
    $ benchstat benchmarks/master-bench benchmarks/json-table-bench
    name              old time/op    new time/op    delta
    CodeMarshal-2       1.10ms ± 1%    0.98ms ± 1%  -10.12%  (p=0.000 n=20+54)
    Unmarshal-2         2.82ms ± 1%    2.79ms ± 0%   -1.09%  (p=0.000 n=20+51)
    UnmarshalReuse-2    2.80ms ± 0%    2.77ms ± 0%   -1.03%  (p=0.000 n=20+52)
    
    name              old speed      new speed      delta
    CodeMarshal-2     87.3MB/s ± 1%  97.1MB/s ± 1%  +11.27%  (p=0.000 n=20+54)
    Unmarshal-2       33.9MB/s ± 1%  34.2MB/s ± 0%   +1.10%  (p=0.000 n=20+51)
    
    For what it's worth, I tried other heuristics - short circuiting the
    conditional for common ASCII characters, for example:
    
    if (b >= 63 && b != 92) || (b >= 39 && b <= 59) || (rest of the conditional)
    
    This offered a speedup around 7-9%, not as large as the submitted
    change.
    
    Change-Id: Idcf88f7b93bfcd1164cdd6a585160b7e407a0d9b
    Reviewed-on: https://go-review.googlesource.com/24466
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2321895fe2a2def7b511453329f4cd8662230256
Author: Martin Möhrmann <martisch@uos.de>
Date:   Wed Sep 7 08:59:00 2016 +0200

    bytes: improve WriteRune performance
    
    Remove the runeBytes buffer and write the utf8 encoding directly
    to the internal buf byte slice.
    
    name         old time/op   new time/op   delta
    WriteRune-4   80.5µs ± 2%   57.1µs ± 2%  -29.06%  (p=0.000 n=20+20)
    
    name         old speed     new speed     delta
    WriteRune-4  153MB/s ± 2%  215MB/s ± 2%  +40.96%  (p=0.000 n=20+20)
    
    Change-Id: Ic15f6e2d6e56a3d15c74f56159e2eae020ba73ba
    Reviewed-on: https://go-review.googlesource.com/28816
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 07bcc165475f3c34433ebf48b05f704fd40e5639
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Sep 8 09:49:31 2016 -0700

    runtime: simplify getargp
    
    Change-Id: I9ed62e8a6d8b9204c18748efd7845adabf3460b9
    Reviewed-on: https://go-review.googlesource.com/28775
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 70fd814f53ba57e6523363d865a6ba49063bfa15
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Sep 7 12:06:43 2016 -0400

    cmd/compile/internal/ssa/gen: fix error message for wrong arg length
    
    When arg length is wrong, op is not set, so it always prints
    "should have 0 args".
    
    Change-Id: If7bcb41d993919d0038d2a09e16188c79dfbd858
    Reviewed-on: https://go-review.googlesource.com/28831
    Reviewed-by: Keith Randall <khr@golang.org>

commit bec84c728a3ef14a15740f856d0fbcf6c7a70b01
Author: Michal Bohuslávek <mbohuslavek@gmail.com>
Date:   Thu Sep 8 09:52:36 2016 +0100

    doc: fix typo in the release notes
    
    Change-Id: I003795d8dc2176532ee133740bf35e23a3aa3878
    Reviewed-on: https://go-review.googlesource.com/28811
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 252093f1203da13e1face4f71141ac75482ccf11
Author: Martin Möhrmann <martisch@uos.de>
Date:   Tue Sep 6 10:38:16 2016 +0200

    runtime: remove maxstring
    
    Before this CL the runtime prevented printing of overlong strings with the print
    function when the length of the string was determined to be corrupted.
    Corruption was checked by comparing the string size against the limit
    which was stored in maxstring.
    
    However maxstring was not updated everywhere were go strings were created
    e.g. for string constants during compile time. Thereby the check for maximum
    string length prevented the printing of some valid strings.
    
    The protection maxstring provided did not warrant the bookkeeping
    and global synchronization needed to keep maxstring updated to the
    correct limit everywhere.
    
    Fixes #16999
    
    Change-Id: I62cc2f4362f333f75b77f199ce1a71aac0ff7aeb
    Reviewed-on: https://go-review.googlesource.com/28813
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fd975c6aa535f2aa066653235be992731d691cfb
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Sep 8 04:51:21 2016 +0000

    io/ioutil: return better error when TempDir called with non-extant dir
    
    Fixes #14196
    
    Change-Id: Ife7950289ac6adbcfc4d0f2fce31f20bc2657858
    Reviewed-on: https://go-review.googlesource.com/28772
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
```
