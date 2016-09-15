# May 19, 2016 Report

Number of commits: 81

## Compilation time

* github.com/coreos/etcd: from 11.199752117s to 13.748744927s, +22.76%
* github.com/boltdb/bolt/cmd/bolt: from 537.330708ms to 541.080265ms, +0.70%
* github.com/gogits/gogs: from 12.160700419s to 12.507641894s, +2.85%
* github.com/spf13/hugo: from 6.368027571s to 6.466065948s, +1.54%
* github.com/influxdata/influxdb/cmd/influxd: from 6.212942347s to 6.374237614s, +2.60%
* github.com/nsqio/nsq/apps/nsqd: from 2.080219878s to 2.146539308s, +3.19%
* github.com/mholt/caddy: from 4.503140483s to 4.583235646s, +1.78%

## Binary size:

* github.com/coreos/etcd: from 21870584 to 21988376, +0.54%
* github.com/boltdb/bolt/cmd/bolt: from 2565367 to 2581853, +0.64%
* github.com/gogits/gogs: from 22752367 to 22852182, +0.44%
* github.com/spf13/hugo: from 14675939 to 14759370, +0.57%
* github.com/influxdata/influxdb/cmd/influxd: from 15778904 to 15862335, +0.53%
* github.com/nsqio/nsq/apps/nsqd: from 9713258 to 9784401, +0.73%
* github.com/mholt/caddy: from 12667934 to 12747269, +0.63%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               185           188           +1.62%
BenchmarkMsgpUnmarshal-4             392           393           +0.26%
BenchmarkEasyJsonMarshal-4           1431          1428          -0.21%
BenchmarkEasyJsonUnmarshal-4         1462          1486          +1.64%
BenchmarkFlatBuffersMarshal-4        352           345           -1.99%
BenchmarkFlatBuffersUnmarshal-4      283           287           +1.41%
BenchmarkGogoprotobufMarshal-4       162           163           +0.62%
BenchmarkGogoprotobufUnmarshal-4     245           245           +0.00%

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

* [hash/crc64: Use slicing by 8.](https://github.com/golang/go/commit/9d73e146dade6553f2365de2ada0156dcb6026d9)


## GIT Log

```
commit 1f7a0d4b5ec7ef94b96755e9b95168abf86e9d71
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon May 16 15:39:43 2016 -0700

    runtime: don't do a plain throw when throwsplit == true
    
    The test case in #15639 somehow causes an invalid syscall frame. The
    failure is obscured because the throw occurs when throwsplit == true,
    which causes a "stack split at bad time" error when trying to print the
    throw message.
    
    This CL fixes the "stack split at bad time" by using systemstack. No
    test because there shouldn't be any way to trigger this error anyhow.
    
    Update #15639.
    
    Change-Id: I4240f3fd01bdc3c112f3ffd1316b68504222d9e1
    Reviewed-on: https://go-review.googlesource.com/23153
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 2a12035f8ec18f0a577853fda78faf2826397131
Author: Scott Bell <scott@sctsm.com>
Date:   Wed May 18 18:44:46 2016 -0700

    expvar: slightly expand documentation for Var's String method
    
    Fixes #15088.
    
    Change-Id: I7727829a4062e15c0e5e3beff4d0bfc1fa327b0f
    Reviewed-on: https://go-review.googlesource.com/23232
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 255e206b2bae9e7632043e08cf8cc0c7ce445c31
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu May 19 02:13:36 2016 +0000

    net/http: update bundled http2
    
    Updates x/net/http2 to git rev 5916dcb1 for:
    
    * http2, lex/httplex: make Transport reject bogus headers before sending
      https://golang.org/cl/23229
    
    * http2: reject more trailer values
      https://golang.org/cl/23230
    
    Fixes #14048
    Fixes #14188
    
    Change-Id: Iaa8beca6e005267a3e849a10013eb424a882f2bb
    Reviewed-on: https://go-review.googlesource.com/23234
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8d428ed218d2b65dbb4abbd9be870c95439a2b14
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu May 19 12:04:10 2016 +0900

    net: don't return io.EOF from zero byte reads on Plan 9
    
    Updates #15735.
    Fixes #15741.
    
    Change-Id: Ic4ad7e948e8c3ab5feffef89d7a37417f82722a1
    Reviewed-on: https://go-review.googlesource.com/23199
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5bcdd639331cd7f8d844fd38a674c4751423f938
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 18 21:54:12 2016 +0000

    net: don't return io.EOF from zero byte reads
    
    Updates #15735
    
    Change-Id: I42ab2345443bbaeaf935d683460fc2c941b7679c
    Reviewed-on: https://go-review.googlesource.com/23227
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c08436d1c897996055892882d23ce6778f3492f7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed May 18 15:20:56 2016 -0700

    runtime: print PC, not the counter, for a cgo traceback
    
    Change-Id: I54ed7a26a753afb2d6a72080e1f50ce9fba7c183
    Reviewed-on: https://go-review.googlesource.com/23228
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 538537a28dc956f069b83e3f1966683901205331
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed May 18 13:19:24 2016 -0700

    runtime: check only up to ptrdata bytes for pointers
    
    Fixes #14508.
    
    Change-Id: I237d0c5a79a73e6c97bdb2077d8ede613128b978
    Reviewed-on: https://go-review.googlesource.com/23224
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit ebbe4f8db76b947663cc535602054c84b01b080d
Author: David Benjamin <davidben@google.com>
Date:   Mon Feb 15 11:41:40 2016 -0500

    crypto/tls: Never resume sessions across different versions.
    
    Instead, decline the session and do a full handshake. The semantics of
    cross-version resume are unclear, and all major client implementations
    treat this as a fatal error. (This doesn't come up very much, mostly if
    the client does the browser version fallback without sharding the
    session cache.)
    
    See BoringSSL's bdf5e72f50e25f0e45e825c156168766d8442dde and OpenSSL's
    9e189b9dc10786c755919e6792e923c584c918a1.
    
    Change-Id: I51ca95ac1691870dd0c148fd967739e2d4f58824
    Reviewed-on: https://go-review.googlesource.com/21152
    Reviewed-by: Adam Langley <agl@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d8bd7b24fcc72fb4117f7fc249ceaa79f69d4e00
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 18 17:59:12 2016 +0000

    net/http: update bundled x/net/http2 for Server context changes
    
    Updates x/net/http2 to golang.org/cl/23220
    (http2: with Go 1.7 set Request.Context in ServeHTTP handlers)
    
    Fixes #15134
    
    Change-Id: I73bac2601118614528f051e85dab51dc48e74f41
    Reviewed-on: https://go-review.googlesource.com/23221
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 1efec481d0997e260f4524d45d11cc35bed63f73
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 18 20:07:16 2016 +0000

    net/http: further restrict when Transport's automatic HTTP/2 happens
    
    Make the temporary, conservative restrictions from rev 79d9f48c in Go
    1.6 permanent, and also don't do automatic TLS if the user configured
    a Dial or DialTLS hook. (Go 1.7 has Transport.Dialer instead, for
    tweaking dialing parameters)
    
    Fixes #14275
    
    Change-Id: I5550d5c1e3a293e103eb4251a3685dc204a23941
    Reviewed-on: https://go-review.googlesource.com/23222
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1249197936aef58cb2296a3cd57b519ba3243042
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed May 18 13:43:50 2016 -0700

    doc/go1.7: add runtime.KeepAlive
    
    Update #13347.
    
    Change-Id: I04bf317ed409478a859355f833d4a5e30db2b9c9
    Reviewed-on: https://go-review.googlesource.com/23226
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 075880a8e8f6363a554c100ad09a85d108953eea
Author: Keith Randall <khr@golang.org>
Date:   Wed May 18 13:28:48 2016 -0700

    cmd/compile: fix build
    
    Run live vars test only on ssa builds.
    We can't just drop KeepAlive ops during regalloc.  We need
    to replace them with copies.
    
    Change-Id: Ib4b3b1381415db88fdc2165fc0a9541b73ad9759
    Reviewed-on: https://go-review.googlesource.com/23225
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 6ab45c09f6fc1bde56e3a72e50505b9a5021aaaf
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri May 13 09:02:40 2016 -0700

    runtime: add KeepAlive function
    
    Fixes #13347.
    
    Change-Id: I591a80a1566ce70efb5f68e3ad69e7e3ab98cd9b
    Reviewed-on: https://go-review.googlesource.com/23102
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3572c6418b5032fbd7e888e14fd9ad5afac85dfc
Author: Keith Randall <khr@golang.org>
Date:   Thu Apr 21 19:28:28 2016 -0700

    cmd/compile: keep pointer input arguments live throughout function
    
    Introduce a KeepAlive op which makes sure that its argument is kept
    live until the KeepAlive.  Use KeepAlive to mark pointer input
    arguments as live after each function call and at each return.
    
    We do this change only for pointer arguments.  Those are the
    critical ones to handle because they might have finalizers.
    Doing compound arguments (slices, structs, ...) is more complicated
    because we would need to track field liveness individually (we do
    that for auto variables now, but inputs requires extra trickery).
    
    Turn off the automatic marking of args as live.  That way, when args
    are explicitly nulled, plive will know that the original argument is
    dead.
    
    The KeepAlive op will be the eventual implementation of
    runtime.KeepAlive.
    
    Fixes #15277
    
    Change-Id: I5f223e65d99c9f8342c03fbb1512c4d363e903e5
    Reviewed-on: https://go-review.googlesource.com/22365
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit d52022676da939cb183083da4ee0b614f86ac3b0
Author: Andrew Gerrand <adg@golang.org>
Date:   Thu May 5 13:25:19 2016 -0700

    html/template: mention risks of the CSS, HTML, JS, etc. types
    
    Fixes #15399
    
    Change-Id: I5b9645cb9ddede6981ce0a005e0c6fdd8a751c6f
    Reviewed-on: https://go-review.googlesource.com/22824
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 8f13080267d0ddbb50da9029339796841224116a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 18 15:42:54 2016 +0000

    net/http: allow Client.CheckRedirect to use most recent response
    
    Fixes #10069
    
    Change-Id: I3819ff597d5a0c8e785403bf9d65a054f50655a6
    Reviewed-on: https://go-review.googlesource.com/23207
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 5d92aefc18317578226a3873fb8fc37411cd2184
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 18 16:18:32 2016 +0000

    vendor, net/http: update vendored hpack
    
    Updates x/net/http2/hpack to rev 6050c111 for:
    
       http2/hpack: forbid excess and invalid padding in hpack decoder
       https://golang.org/cl/23067
    
    Updates #15614
    
    Change-Id: I3fbf9b265bfa5e49e6aa97d8c34e08214cfcc49a
    Reviewed-on: https://go-review.googlesource.com/23208
    Reviewed-by: Carl Mastrangelo <notcarl@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1119af89767dc4086cba336e732afcea084c8c34
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 18 04:42:37 2016 +0000

    net/http: update bundled x/net/http2 for httptrace changes
    
    Updates x/net/http2 to 3b99394 for golang.org/cl/23205
    
    And associated tests.
    
    Fixes #12580
    
    Change-Id: I1f4b59267b453d241f2afaa315b7fe10d477e52d
    Reviewed-on: https://go-review.googlesource.com/23206
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4d2ac544a437aaf7bbd78d1a46baa5108945f06e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 30 19:57:28 2016 -0700

    net/http: fix spurious logging in Transport when server closes idle conn
    
    In https://golang.org/3210, Transport errors occurring before
    receiving response headers were wrapped in another error type to
    indicate to the retry logic elsewhere that the request might be
    re-tryable. But a check for err == io.EOF was missed, which then became
    false once io.EOF was wrapped in the beforeRespHeaderError type.
    
    The beforeRespHeaderError was too fragile. Remove it. I tried to fix
    it in an earlier version of this CL and just broke different things
    instead.
    
    Also remove the "markBroken" method. It's redundant and confusing.
    
    Also, rename the checkTransportResend method to shouldRetryRequest and
    make it return a bool instead of an error. This also helps readability.
    
    Now the code recognizes the two main reasons we'd want to retry a
    request: because we never wrote the request in the first place (so:
    count the number of bytes we've written), or because the server hung
    up on us before we received response headers for an idempotent request.
    
    As an added bonus, this could make POST requests safely re-tryable
    since we know we haven't written anything yet. But it's too late in Go
    1.7 to enable that, so we'll do that later (filed #15723).
    
    This also adds a new internal (package http) test, since testing this
    blackbox at higher levels in transport_test wasn't possible.
    
    Fixes #15446
    
    Change-Id: I2c1dc03b1f1ebdf3f04eba81792bd5c4fb6b6b66
    Reviewed-on: https://go-review.googlesource.com/23160
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9d73e146dade6553f2365de2ada0156dcb6026d9
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Tue Apr 19 19:05:53 2016 +0300

    hash/crc64: Use slicing by 8.
    
    Similar to crc32 slicing by 8.
    This also fixes a Crc64KB benchmark actually using 1024 bytes.
    
    Crc64/ISO64KB-4       147µs ± 0%      37µs ± 0%   -75.05%  (p=0.000 n=18+18)
    Crc64/ISO4KB-4       9.19µs ± 0%    2.33µs ± 0%   -74.70%  (p=0.000 n=19+20)
    Crc64/ISO1KB-4       2.31µs ± 0%    0.60µs ± 0%   -73.81%  (p=0.000 n=19+15)
    Crc64/ECMA64KB-4      147µs ± 0%      37µs ± 0%   -75.05%  (p=0.000 n=20+20)
    Crc64/Random64KB-4    147µs ± 0%      41µs ± 0%   -72.17%  (p=0.000 n=20+18)
    Crc64/Random16KB-4   36.7µs ± 0%    36.5µs ± 0%    -0.54%  (p=0.000 n=18+19)
    
    name                old speed     new speed      delta
    Crc64/ISO64KB-4     446MB/s ± 0%  1788MB/s ± 0%  +300.72%  (p=0.000 n=18+18)
    Crc64/ISO4KB-4      446MB/s ± 0%  1761MB/s ± 0%  +295.20%  (p=0.000 n=18+20)
    Crc64/ISO1KB-4      444MB/s ± 0%  1694MB/s ± 0%  +281.46%  (p=0.000 n=19+20)
    Crc64/ECMA64KB-4    446MB/s ± 0%  1788MB/s ± 0%  +300.77%  (p=0.000 n=20+20)
    Crc64/Random64KB-4  446MB/s ± 0%  1603MB/s ± 0%  +259.32%  (p=0.000 n=20+18)
    Crc64/Random16KB-4  446MB/s ± 0%   448MB/s ± 0%    +0.54%  (p=0.000 n=18+20)
    
    Change-Id: I1c7621d836c486d6bfc41dbe1ec2ff9ab11aedfc
    Reviewed-on: https://go-review.googlesource.com/22222
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 6cd698d71da92aeb4540c378213ac4a1c6687097
Author: Lee Hinman <hinman@gmail.com>
Date:   Mon Mar 7 22:31:31 2016 -0600

    crypto/x509: add Admin & User Keychains to FetchPEMRoots on Darwin
    
    in root_cgo_darwin.go only certificates from the System Domain
    were being used in FetchPEMRoots.  This patch adds support for
    getting certificates from all three domains (System, Admin,
    User).  Also it will only read trusted certificates from those
    Keychains.  Because it is possible to trust a non Root certificate,
    this patch also adds a checks to see if the Subject and Issuer
    name are the same.
    
    Fixes #14514
    
    Change-Id: Ia03936d7a61d1e24e99f31c92f9927ae48b2b494
    Reviewed-on: https://go-review.googlesource.com/20351
    Reviewed-by: Russ Cox <rsc@golang.org>

commit b30fcbc9f59ca4bf1723eb6743b47fa89b3847a3
Author: Adam Langley <agl@golang.org>
Date:   Thu Apr 14 13:52:56 2016 -0700

    crypto/ecdsa: reject negative inputs.
    
    The fact that crypto/ecdsa.Verify didn't reject negative inputs was a
    mistake on my part: I had unsigned numbers on the brain. However, it
    doesn't generally cause problems. (ModInverse results in zero, which
    results in x being zero, which is rejected.)
    
    The amd64 P-256 code will crash when given a large, negative input.
    
    This fixes both crypto/ecdsa to reject these values and also the P-256
    code to ignore the sign of inputs.
    
    Change-Id: I6370ed7ca8125e53225866f55b616a4022b818f8
    Reviewed-on: https://go-review.googlesource.com/22093
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2ba8fc5b086942dbb23282702f61c813298867f3
Author: Monty Taylor <mordred@inaugust.com>
Date:   Tue May 17 08:24:18 2016 -0500

    vcs: Add support for git.openstack.org
    
    Go is being proposed as an officially supported language for elements of
    OpenStack:
    
      https://review.openstack.org/#/c/312267/
    
    As such, repos that exist in OpenStack's git infrastructure
    are likely to become places from which people might want to go get
    things. Allow optional .git suffixes to allow writing code that depends
    on git.openstack.org repos that will work with older go versions while
    we wait for this support to roll out.
    
    Change-Id: Ia64bdb1dafea33b1c3770803230d30ec1059df22
    Reviewed-on: https://go-review.googlesource.com/23135
    Reviewed-by: Russ Cox <rsc@golang.org>

commit d35a4158ab66aef99d9204c65cc2e2fa74b57a73
Author: David Chase <drchase@google.com>
Date:   Mon May 9 14:59:25 2016 -0400

    cmd/compile: reduce element size of arrays in sparse{map,set}
    
    sparseSet and sparseMap only need 32 bit integers in their
    arrays, since a sparseEntry key is also limited to 32 bits.
    This appears to reduce the space allocated for at least
    one pathological compilation by 1%, perhaps more.
    
    Not necessarily for 1.7, but it saves a little and is very
    low-risk.
    
    Change-Id: Icf1185859e9f5fe1261a206b441e02c34f7d02fd
    Reviewed-on: https://go-review.googlesource.com/22972
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 2380a039c0457141e28f8f927139e1f9c38f8205
Author: Cuihtlauac ALVARADO <cuihtlauac.alvarado@orange.com>
Date:   Tue May 17 09:27:00 2016 +0200

    runtime: in tests, make sure gdb does not start with a shell
    
    On some systems, gdb is set to: "startup-with-shell on". This
    breaks runtime_test. This just make sure gdb does not start by
    spawning a shell.
    
    Fixes #15354
    
    Change-Id: Ia040931c61dea22f4fdd79665ab9f84835ecaa70
    Reviewed-on: https://go-review.googlesource.com/23142
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c65647d6204531e93c19ea2dba01ff13d1b8ef31
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Apr 19 15:02:06 2016 -0700

    cmd/compile: handle unsafe.Pointer(f()) correctly
    
    Previously statements like
    
        f(unsafe.Pointer(g()), int(h()))
    
    would be reordered into a sequence of statements like
    
        autotmp_g := g()
        autotmp_h := h()
        f(unsafe.Pointer(autotmp_g), int(autotmp_h))
    
    which can leave g's temporary value on the stack as a uintptr, rather
    than an unsafe.Pointer. Instead, recognize uintptr-to-unsafe.Pointer
    conversions when reordering function calls to instead produce:
    
        autotmp_g := unsafe.Pointer(g())
        autotmp_h := h()
        f(autotmp_g, int(autotmp_h))
    
    Fixes #15329.
    
    Change-Id: I2cdbd89d233d0d5c94791513a9fd5fd958d11ed5
    Reviewed-on: https://go-review.googlesource.com/22273
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit c6a5b3602a87b2d1321ad11aa64b7f588bbb683b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue May 17 18:07:07 2016 -0700

    doc/effective_go: clarify backward function reference
    
    Fixes #14656.
    
    Change-Id: I37a9aa51705ae18bd034f2cc6dbf06a55f969197
    Reviewed-on: https://go-review.googlesource.com/23202
    Reviewed-by: Rob Pike <r@golang.org>

commit d4ed8da9969dd04a3b10683971185359e5ec5302
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed May 18 13:02:39 2016 +0900

    net: don't increase test table rows when using -test.count flag
    
    Change-Id: I7881e3353dc5cd9755a79ea0eab146c6a0a08306
    Reviewed-on: https://go-review.googlesource.com/23195
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6de34e6e25732757b7b40e4053c6ac7fb6d00df4
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed May 18 10:54:20 2016 +0900

    net: deflake TestLookupPort on Android
    
    Looks like some version of Android still fails with "servname not
    supported for ai_socktype". It probably doesn't support
    ai_socktype=SOCK_STREAM.
    
    Updates #14576.
    
    Change-Id: I77ecff147d5b759e3281b3798c60f150a4aab811
    Reviewed-on: https://go-review.googlesource.com/23194
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit cdcb8271a411bac78aa886a5998ac2c10b23f058
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 18 01:00:32 2016 +0000

    regexp/syntax: clarify that \Z means Perl's \Z
    
    Fixes #14793
    
    Change-Id: I408056d096cd6a999fa5e349704b5ea8e26d2e4e
    Reviewed-on: https://go-review.googlesource.com/23201
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 23a59ba17cbfeb5380845f309f88165b2e38930b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue May 17 18:16:47 2016 -0700

    runtime: deflake TestSignalExitStatus
    
    The signal might get delivered to a different thread, and that thread
    might not run again before the currently running thread returns and
    exits. Sleep to give the other thread time to pick up the signal and
    crash.
    
    Not tested for all cases, but, optimistically:
    Fixes #14063.
    
    Change-Id: Iff58669ac6185ad91cce85e0e86f17497a3659fd
    Reviewed-on: https://go-review.googlesource.com/23203
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit ac66bb343181bef154342638d45dcc2c695ded00
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Tue May 17 20:22:48 2016 -0600

    crypto/x509: fix typo in docs for CreateCertificateRequest
    
    Update the doc for CreateCertificateRequest
    to state that it creates a
      `new certificate request`
    instead of just a
      `new certificate`
    
    Fixes #14649.
    
    Change-Id: Ibbbcf91d74168998990990e78e5272a6cf294d51
    Reviewed-on: https://go-review.googlesource.com/23204
    Reviewed-by: Russ Cox <rsc@golang.org>

commit f962e6e0e21b9e73981e6cf2407ea01fce04b989
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 18 00:35:43 2016 +0000

    net/http: add test confirming a connection reuse case
    
    Verify that for a server doing chunked encoding, with the final data
    and EOF arriving together, the client will reuse the connection even
    if it closes the body without seeing an EOF. The server sends at least
    one non-zero chunk and one zero chunk. This verifies that the client's
    bufio reading reads ahead and notes the EOF, so even if the JSON
    decoder doesn't read the EOF itself, as long as somebody sees it, a
    close won't forcible tear down the connection. This was true at least
    of https://golang.org/cl/21291
    
    No code change. Test already passed (even with lots of runs, including
    in race mode with randomized goroutine scheduling).
    
    Updates #15703
    
    Change-Id: I2140b3eec6b099b6b6e54f153fe271becac5d949
    Reviewed-on: https://go-review.googlesource.com/23200
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 733162fd6c0df8bd700859974957b25045fe9ee4
Author: James Chacon <jchacon@google.com>
Date:   Fri May 6 11:51:56 2016 -0700

    runtime: prevent racefini from being invoked more than once
    
    racefini calls __tsan_fini which is C code and at the end of it
    invoked the standard C library exit(3) call. This has undefined
    behavior if invoked more than once. Specifically in C++ programs
    it caused static destructors to run twice. At least on glibc
    impls it also means the at_exit handlers list (where those are
    stored) also free's a list entry when it completes these. So invoking
    twice results in a double free at exit which trips debug memory
    allocation tracking.
    
    Fix all of this by using an atomic as a boolean barrier around
    calls to racefini being invoked > 1 time.
    
    Fixes #15578
    
    Change-Id: I49222aa9b8ded77160931f46434c61a8379570fc
    Reviewed-on: https://go-review.googlesource.com/22882
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f744717d1924340b8f5e5a385e99078693ad9097
Author: Alessandro Arzilli <alessandro.arzilli@gmail.com>
Date:   Sat May 14 07:54:02 2016 +0200

    debug/gosym: parse remote packages correctly
    
    Fixes #15675
    
    Change-Id: I8bad220988e5d690f20804db970b2db037c81187
    Reviewed-on: https://go-review.googlesource.com/23086
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7b597f4d92a844e30694095485c335baa93a1ad1
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri May 13 05:02:00 2016 +0900

    net: deflake TestLookupPort for embedded, security-hardened platforms
    
    Fixes #14576.
    
    Change-Id: I760907c67c97cb827cf48ba7eb0bb2f4f8d4d790
    Reviewed-on: https://go-review.googlesource.com/23111
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9d36cac99c8248292be1cb6a196bbe0715d0c057
Author: Colin Cross <ccross@android.com>
Date:   Tue May 17 13:09:11 2016 -0700

    reflect: remove out of date UTF-8 StructOf restriction
    
    The initial implementation of reflect.StructOf in
    https://golang.org/cl/9251 had a limitation that field names had to be
    ASCII, which was later lifted by https://golang.org/cl/21777.  Remove
    the out-of-date documentation disallowing UTF-8 field names.
    
    Updates: #5748
    Updates: #15064
    
    Change-Id: I2c5bfea46bfd682449c6e847fc972a1a131f51b7
    Reviewed-on: https://go-review.googlesource.com/23170
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 1a3e4f05a067f37e8ee8f7a0d5eec39a7e0cad84
Author: Konstantin Shaposhnikov <k.shaposhnikov@gmail.com>
Date:   Tue May 17 20:16:41 2016 +0800

    os/signal: fix wrong constant name in the documentation
    
    os.SIGINT is not defined, os.Interrupt or syscall.SIGINT should be used.
    
    Change-Id: I39867726d28e179d1160a4fd353b7bea676c9dbb
    Reviewed-on: https://go-review.googlesource.com/23127
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 495e3c60aa61615dd603050ac47f86468f8222b6
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue May 17 12:20:16 2016 +0900

    net: use IPv4/IPv6 reserved address blocks for documentation
    
    Also replaces google.com with golang.org in package documentation.
    
    Updates #15228.
    
    Change-Id: I554fa960878fa44557a522635ed412d8d7548d3f
    Reviewed-on: https://go-review.googlesource.com/23126
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ca831135b34d13fe5b774a6b23867dd1a277786a
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue May 17 07:07:58 2016 +0900

    net: simplify interfaceTable for BSD variants
    
    This change drops parseInterfaceTable which becomes unnecessary by the
    golang.org/x/net/route plumbing.
    
    Change-Id: I05f96e347de950bb1e9292bb3eeff01bb40e292f
    Reviewed-on: https://go-review.googlesource.com/23125
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5ccd571f3e2798e4afe8affa354351b5055cb20d
Author: Scott Bell <scott@sctsm.com>
Date:   Mon May 16 12:51:52 2016 -0700

    crypto/tls: document certificate chains in LoadX509KeyPair
    
    Fixes #15348
    
    Change-Id: I9e0e1e3a26fa4cd697d2c613e6b4952188b7c7e1
    Reviewed-on: https://go-review.googlesource.com/23150
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ccdca832c569727f7985966a3324421a69739f57
Author: andrew werner <andrew@upthere.com>
Date:   Tue Dec 15 14:42:28 2015 -0800

    io: make chained multiReader Read more efficient
    
    before this change, when io.MultiReader was called many times but contain few
    underlying readers, calls to Read were unnecessarily expensive.
    
    Fixes #13558
    
    Change-Id: I3ec4e88c7b50c075b148331fb1b7348a5840adbe
    Reviewed-on: https://go-review.googlesource.com/17873
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 28c201860e0838b10972e805582007f9eb61e7ac
Author: Simon Thulbourn <simon+github@thulbourn.com>
Date:   Mon Dec 7 16:36:11 2015 +0000

    mime/multipart: sort header keys to ensure reproducible output
    
    Adds a transparent sort to the mime/multipart package, which is
    only used in the CreatePart func. This will ensure the ordering
    of the MIMEHeader.
    
    The point of this change was to ensure the output would be consistent
    and something that could be depended on.
    
    Fixes #13522
    
    Change-Id: I9584ef9dbe98ce97d536d897326914653f8d9ddf
    Reviewed-on: https://go-review.googlesource.com/17497
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 99ef42fe7b310749d83f9b76d814e78fe8139b42
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon May 16 15:45:18 2016 -0700

    A+C: add Andrew Werner (corporate CLA for Upthere, Inc)
    
    Change-Id: I7627e480d5d2366cba223fd81635c4115649f752
    Reviewed-on: https://go-review.googlesource.com/23154
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c1b32acefb3b1438981ba9dc4f5259999e9fc2ab
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon May 16 14:11:01 2016 -0700

    runtime: yield after raising signal that should kill process
    
    Issue #15613 points out that the darwin builders have been getting
    regular failures in which a process that should exit with a SIGPIPE
    signal is instead exiting with exit status 2. The code calls
    runtime.raise. On most systems runtime.raise is the equivalent of
    pthread_kill(gettid(), sig); that is, it kills the thread with the
    signal, which should ensure that the program does not keep going. On
    darwin, however, runtime.raise is actually kill(getpid(), sig); that is,
    it sends a signal to the entire process. If the process decides to
    deliver the signal to a different thread, then it is possible that in
    some cases the thread that calls raise is able to execute the next
    system call before the signal is actually delivered. That would cause
    the observed error.
    
    I have not been able to recreate the problem myself, so I don't know
    whether this actually fixes it. But, optimistically:
    
    Fixed #15613.
    
    Change-Id: I60c0a9912aae2f46143ca1388fd85e9c3fa9df1f
    Reviewed-on: https://go-review.googlesource.com/23152
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1b86862d0d6eeb818bc622fee5f140951bd31063
Author: Scott Bell <scott@sctsm.com>
Date:   Mon May 16 13:13:25 2016 -0700

    doc: fix broken link to the vet command documentation
    
    Fixes #15188
    
    Change-Id: I0ab7791f7db499cef6bc52292d3d93ff4da7caff
    Reviewed-on: https://go-review.googlesource.com/23151
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6b99fb5bea56b9674161b6c5415c1d05f19dfdc6
Author: David Chase <drchase@google.com>
Date:   Thu Apr 21 13:24:58 2016 -0400

    cmd/compile: use sparse algorithm for phis in large program
    
    This adds a sparse method for locating nearest ancestors
    in a dominator tree, and checks blocks with more than one
    predecessor for differences and inserts phi functions where
    there are.
    
    Uses reversed post order to cut number of passes, running
    it from first def to last use ("last use" for paramout and
    mem is end-of-program; last use for a phi input from a
    backedge is the source of the back edge)
    
    Includes a cutover from old algorithm to new to avoid paying
    large constant factor for small programs.  This keeps normal
    builds running at about the same time, while not running
    over-long on large machine-generated inputs.
    
    Add "phase" flags for ssa/build -- ssa/build/stats prints
    number of blocks, values (before and after linking references
    and inserting phis, so expansion can be measured), and their
    product; the product governs the cutover, where a good value
    seems to be somewhere between 1 and 5 million.
    
    Among the files compiled by make.bash, this is the shape of
    the tail of the distribution for #blocks, #vars, and their
    product:
    
    	 #blocks	#vars	    product
     max	6171	28180	173,898,780
    99.9%	1641	 6548	 10,401,878
      99%	 463	 1909	    873,721
      95%	 152	  639	     95,235
      90%	  84	  359	     30,021
    
    The old algorithm is indeed usually fastest, for 99%ile
    values of usually.
    
    The fix to LookupVarOutgoing
    ( https://go-review.googlesource.com/#/c/22790/ )
    deals with some of the same problems addressed by this CL,
    but on at least one bug ( #15537 ) this change is still
    a significant help.
    
    With this CL:
    /tmp/gopath$ rm -rf pkg bin
    /tmp/gopath$ time go get -v -gcflags -memprofile=y.mprof \
       github.com/gogo/protobuf/test/theproto3/combos/...
    ...
    real	4m35.200s
    user	13m16.644s
    sys	0m36.712s
    and pprof reports 3.4GB allocated in one of the larger profiles
    
    With tip:
    /tmp/gopath$ rm -rf pkg bin
    /tmp/gopath$ time go get -v -gcflags -memprofile=y.mprof \
       github.com/gogo/protobuf/test/theproto3/combos/...
    ...
    real	10m36.569s
    user	25m52.286s
    sys	4m3.696s
    and pprof reports 8.3GB allocated in the same larger profile
    
    With this CL, most of the compilation time on the benchmarked
    input is spent in register/stack allocation (cumulative 53%)
    and in the sparse lookup algorithm itself (cumulative 20%).
    
    Fixes #15537.
    
    Change-Id: Ia0299dda6a291534d8b08e5f9883216ded677a00
    Reviewed-on: https://go-review.googlesource.com/22342
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 466cae6ca9f28c971a2d716a5a49dc76bbd1d5bb
Author: Austin Clements <austin@google.com>
Date:   Mon May 16 15:27:48 2016 -0400

    runtime: use GOTRACEBACK=system for TestStackBarrierProfiling
    
    This should help with debugging failures.
    
    For #15138 and #15477.
    
    Change-Id: I77db2b6375d8b4403d3edf5527899d076291e02c
    Reviewed-on: https://go-review.googlesource.com/23134
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 181000896e381f07e8f105eef2667d566729f6eb
Author: Scott Bell <scott@sctsm.com>
Date:   Mon May 16 12:36:02 2016 -0700

    encoding/json: document that object keys are sorted
    
    Fixes #15424
    
    Change-Id: Ib9e97509f5ac239ee54fe6fe37152a7f5fc75087
    Reviewed-on: https://go-review.googlesource.com/23109
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 64770f642fe4b2f3af6f30d1e058e934e73d0b9b
Author: Austin Clements <austin@google.com>
Date:   Mon May 9 11:29:34 2016 -0400

    runtime: use conventional shift style for gcBitsChunkBytes
    
    The convention for writing something like "64 kB" is 64<<10, since
    this is easier to read than 1<<16. Update gcBitsChunkBytes to follow
    this convention.
    
    Change-Id: I5b5a3f726dcf482051ba5b1814db247ff3b8bb2f
    Reviewed-on: https://go-review.googlesource.com/23132
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 30ded16596246c719ede90acf45ecb31d8f428f6
Author: Austin Clements <austin@google.com>
Date:   Wed May 11 16:29:07 2016 -0400

    runtime: remove obsolete comment from scanobject
    
    Change-Id: I5ebf93b60213c0138754fc20888ae5ce60237b8c
    Reviewed-on: https://go-review.googlesource.com/23131
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit b66b97e0a120880e37b03eba00c0c7679f0a70c1
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Mon May 16 10:11:59 2016 -0300

    net/http: mention ALPN in http.Server.TLSNextProto documentation
    
    Make clear negotiation can happen via NPN or ALPN, similar to
    http.Transport.TLSNextProto and x/net/http2.NextProtoTLS.
    
    Change-Id: Ied00b842bc04e11159d6d2107beda921cefbc6ca
    Reviewed-on: https://go-review.googlesource.com/23108
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6a6c1d9841a1957a2fd292df776ea920ae38ea00
Author: Artyom Pervukhin <artyom.pervukhin@gmail.com>
Date:   Mon May 16 15:30:28 2016 +0300

    net/http/httputil: don't add User-Agent header by proxy made with NewSingleHostReverseProxy
    
    If client does not provided User-Agent header, do not set default one
    used by net/http package when doing request to backend.
    
    Fixes #15524
    
    Change-Id: I9a46bb3b7ec106bc7c3071e235b872d279994d67
    Reviewed-on: https://go-review.googlesource.com/23089
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a101b85e00f302706d8b1de1d2173a154d5f54cc
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sun May 15 17:24:51 2016 +0900

    syscall: fix missing use of use function in sysctl
    
    Updates #13372.
    
    Change-Id: Id2402a781474e9d0bb0901c5844adbd899f76cbd
    Reviewed-on: https://go-review.googlesource.com/23123
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5fae488633ff247b9b7964dc45b8fe4b491f5a16
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sun May 15 17:16:04 2016 +0900

    syscall: deprecate BPF/LSF
    
    Updates #14982.
    
    Change-Id: Id12b1e61456832d2b2ffbdbe8cf0a1db4444b1e4
    Reviewed-on: https://go-review.googlesource.com/23122
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6d66819587b9de3d7602721830884fd92a0f7090
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Apr 23 22:36:41 2016 +0900

    syscall: deprecate routing message APIs for BSD variants
    
    Also removes unnecessary test cases for avoiding unexpected failures on
    newer operating systems.
    
    Updates #14724.
    
    Change-Id: I2291585d951fb70383da68293a6ac1ff3524c7f7
    Reviewed-on: https://go-review.googlesource.com/22452
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6cc6ef82ea1ea5c904bb44c14bb6f4eb33937bb1
Author: David du Colombier <0intro@gmail.com>
Date:   Sun May 15 20:12:34 2016 +0200

    mime: fix mime type file name on Plan 9
    
    There was a typo introduced in the initial
    implementation of the Plan 9 support of
    the mime package.
    
    On Plan 9, the mime type file name should be
    /sys/lib/mimetype instead of /sys/lib/mimetypes.
    
    Change-Id: If0f0a9b6f3fbfa8dde551f790e83bdd05e8f0acb
    Reviewed-on: https://go-review.googlesource.com/23087
    Run-TryBot: Minux Ma <minux@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b4bf0663fa3334d053981f222eed5015a0a1b8df
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Apr 23 22:36:41 2016 +0900

    net: golang.org/x/net/route plumbing
    
    This change makes use of new routing message APIs for BSD variants to
    support FreeBSD 11 and newer versions of other BSDs.
    
    Fixes #7849.
    Fixes #14724.
    
    Change-Id: I56c7886d6622cdeddd7cc29c8a8062dcc06216d5
    Reviewed-on: https://go-review.googlesource.com/22451
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0bc14f57ec7e5518af711a64103ca2ac72f19a6e
Author: Keith Randall <khr@golang.org>
Date:   Sat May 14 17:33:23 2016 -0700

    strings: fix Contains on amd64
    
    The 17-31 byte code is broken.  Disabled it.
    
    Added a bunch of tests to at least cover the cases
    in indexShortStr.  I'll channel Brad and wonder why
    this CL ever got in without any tests.
    
    Fixes #15679
    
    Change-Id: I84a7b283a74107db865b9586c955dcf5f2d60161
    Reviewed-on: https://go-review.googlesource.com/23106
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit aba7b3e91cea1c95a1803357bcf219b0591d3c12
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Apr 23 22:36:41 2016 +0900

    vendor: import golang.org/x/net/route
    
    golang.org/x/net/route becomes vendor/golang.org/x/net/route.
    
    At git rev 30be488 (golang.org/cl/22446)
    
    Updates #14724.
    
    Change-Id: I41cfb5443aeecac4c71e843c09eb8c1d4b7413ea
    Reviewed-on: https://go-review.googlesource.com/22450
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2699da180967e5d5dab2cc64deeca4680bf2b2fb
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Fri May 13 14:26:30 2016 +1000

    time: set Local.name on windows
    
    Local.String() returns "Local" on every OS, but windows.
    Change windows code to do like others.
    
    Updates #15568
    
    Change-Id: I7a4d2713d940e2a01cff9d7f5cefc89def07546a
    Reviewed-on: https://go-review.googlesource.com/23078
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5f833121cb8e7722667d17dcf07bb32e4e524f23
Author: Kevin Burke <kev@inburke.com>
Date:   Sat May 14 10:23:09 2016 -0700

    archive/zip: use HTTPS for documentation link
    
    The resource is available over (and redirects to) HTTPS, it seems like a good
    idea to save a redirect and ensure an encrypted connection.
    
    Change-Id: I262c7616ae289cdd756b6f67573ba6bd7e3e0ca6
    Reviewed-on: https://go-review.googlesource.com/23104
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 094afe0cf1aa4182d62a3aea6970933b7ae4a27d
Author: Konstantin Shaposhnikov <k.shaposhnikov@gmail.com>
Date:   Sat May 14 10:27:31 2016 +0800

    cmd/vendor: move cmd/internal/unvendor packages to cmd/vendor
    
    Updates #14047
    
    Change-Id: I4b150533393bfb90e840497095ac32bcca4f04c2
    Reviewed-on: https://go-review.googlesource.com/23114
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6181db53dbfec513300a236debbdc01735f00c07
Author: Austin Clements <austin@google.com>
Date:   Thu May 12 18:10:03 2016 -0400

    runtime: improve heapBitsSetType documentation
    
    Currently the heapBitsSetType documentation says that there are no
    races on the heap bitmap, but that isn't exactly true. There are no
    *write-write* races, but there are read-write races. Expand the
    documentation to explain this and why it's okay.
    
    Change-Id: Ibd92b69bcd6524a40a9dd4ec82422b50831071ed
    Reviewed-on: https://go-review.googlesource.com/23092
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit d8b08c3aa49d9aaac6ff34dbb8516040cc88a13a
Author: Austin Clements <austin@google.com>
Date:   Wed May 11 14:57:33 2016 -0400

    runtime: perform publication barrier even for noscan objects
    
    Currently we only execute a publication barrier for scan objects (and
    skip it for noscan objects). This used to be okay because GC would
    never consult the object itself (so it wouldn't observe uninitialized
    memory even if it found a pointer to a noscan object), and the heap
    bitmap was pre-initialized to noscan.
    
    However, now we explicitly initialize the heap bitmap for noscan
    objects when we allocate them. While the GC will still never consult
    the contents of a noscan object, it does need to see the initialized
    heap bitmap. Hence, we need to execute a publication barrier to make
    the bitmap visible before user code can expose a pointer to the newly
    allocated object even for noscan objects.
    
    Change-Id: Ie4133c638db0d9055b4f7a8061a634d970627153
    Reviewed-on: https://go-review.googlesource.com/23043
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 041cc148faae23714c38ec9e4388715d99aef518
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri May 13 05:02:00 2016 +0900

    net: deflake TestPointToPointInterface and TestInterfaceArrivalAndDeparture
    
    Fixes #6879.
    
    Change-Id: I9ed2460cf14cb9322d9521e7af910efa48abdaf0
    Reviewed-on: https://go-review.googlesource.com/23112
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 19619c21c36d1695000f5b798241971dfdb2cc2d
Author: Tom Bergan <tombergan@google.com>
Date:   Thu May 12 22:03:46 2016 -0700

    net, net/http: don't trace DNS dials
    
    This fixes change https://go-review.googlesource.com/#/c/23069/, which
    assumes all DNS requests are UDP. This is not true -- DNS requests can
    be TCP in some cases. See:
    https://tip.golang.org/src/net/dnsclient_unix.go#L154
    https://en.wikipedia.org/wiki/Domain_Name_System#Protocol_transport
    
    Also, the test code added by the above change doesn't actually test
    anything because the test uses a faked DNS resolver that doesn't
    actually make any DNS queries. I fixed that by adding another test
    that uses the system DNS resolver.
    
    Updates #12580
    
    Change-Id: I6c24c03ebab84d437d3ac610fd6eb5353753c490
    Reviewed-on: https://go-review.googlesource.com/23101
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0cc710dca63b79ed2dd6ce9375502e76e5fc0484
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri May 13 12:33:27 2016 -0400

    reflect: fix method type string
    
    By picking up a spurious tFlagExtraStar, the method type was printing
    as unc instead of func.
    
    Updates #15673
    
    Change-Id: I0c2c189b99bdd4caeb393693be7520b8e3f342bf
    Reviewed-on: https://go-review.googlesource.com/23103
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2cc0f2209653f9f6931e0c3a1fb63e581a0fe87f
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri May 13 11:10:29 2016 +0900

    Revert "cmd/build: reject non-existant directories in ImportDir"
    
    This reverts commit 7af2ce3f159760033c903b3730bfb5995b4edd40.
    
    The commit had a wrong prefix in the description line, probably
    misreconginized something. As a result it broke golang.org/x/tools/godoc
    and golang.org/x/tools/cmd/godoc like the following:
    
    --- FAIL: TestCLI (10.90s)
    --- FAIL: TestWeb (13.74s)
    FAIL
    FAIL        golang.org/x/tools/cmd/godoc    36.428s
    --- FAIL: TestCommandLine (0.00s)
    FAIL
    FAIL        golang.org/x/tools/godoc        0.068s
    
    Change-Id: I362a862a4ded8592dec7488a28e7a256adee148f
    Reviewed-on: https://go-review.googlesource.com/23076
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit be5782c330f2c743f81942f5bc1b9c1e04296d44
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri May 13 09:25:25 2016 +0200

    doc/go1.7.txt: add cmd/trace changes
    
    Change-Id: Iaf455d1a2863ff752e0c398e1c364373b4d36614
    Reviewed-on: https://go-review.googlesource.com/23084
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 15f2d0e45227f68024f3415d9466055877b70426
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu May 12 22:12:11 2016 +0000

    net, net/http: don't trace UDP dials
    
    The httptrace.ConnectStart and ConnectDone hooks are just about the
    post-DNS connection to the host. We were accidentally also firing on
    the UDP dials to DNS. Exclude those for now. We can add them back
    later as separate hooks if desired. (but they'd only work for pure Go
    DNS)
    
    This wasn't noticed earlier because I was developing on a Mac at the
    time, which always uses cgo for DNS. When running other tests on
    Linux, I started seeing UDP dials.
    
    Updates #12580
    
    Change-Id: I2b2403f2483e227308fe008019f1100f6300250b
    Reviewed-on: https://go-review.googlesource.com/23069
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4cffe44e361deb39e3274774a7984ab78a5b3931
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu May 12 15:04:05 2016 +1000

    syscall: separate stdlib imports from others in mksyscall_windows.go
    
    Change-Id: I6610b872578d161e535565258039d9f064f01456
    Reviewed-on: https://go-review.googlesource.com/23070
    Reviewed-by: Nigel Tao <nigeltao@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit eb69476c66339ca494f98e65a78d315da99a9c79
Author: Andrew Gerrand <adg@golang.org>
Date:   Thu May 12 13:55:46 2016 -0700

    text/template: detect pathologically recursive template invocations
    
    Return an error message instead of eating memory and eventually
    triggering a stack overflow.
    
    Fixes #15618
    
    Change-Id: I3dcf1d669104690a17847a20fbfeb6d7e39e8751
    Reviewed-on: https://go-review.googlesource.com/23091
    Reviewed-by: Rob Pike <r@golang.org>

commit 8f48efb31c7cdddeec7d4221174254466b0891dd
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Fri May 13 02:05:48 2016 +0530

    fmt: remove extra space in too few arguments example
    
    Change-Id: Iae4855c52c4da9755277251d22121226507ea26a
    Reviewed-on: https://go-review.googlesource.com/23074
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7af2ce3f159760033c903b3730bfb5995b4edd40
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Tue May 10 21:28:30 2016 -0700

    cmd/build: reject non-existant directories in ImportDir
    
    Re-apply @adg's CL https://golang.org/cl/7129048 that was
    previously disabled in https://golang.org/cl/7235052 because
    it broke `godoc net/http` for go1.1.
    
    Currently `godoc net/http` seems to work fine with this CL.
    
    Fixes #3428.
    
    Change-Id: I7df06df02fd62dededac6ec60bea62561be59cf1
    Reviewed-on: https://go-review.googlesource.com/23013
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 7ae273923cdd5d00b72c293b57ade8a1e290a4a3
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Tue May 3 16:44:25 2016 +0200

    cmd/trace: split large traces into parts
    
    Trace viewer cannot handle traces larger than 256MB (limit on js string size):
    https://github.com/catapult-project/catapult/issues/627
    And even that is problematic (chrome hangs and crashes).
    Split large traces into 100MB parts. Somewhat clumsy, but I don't see any other
    solution (other than rewriting trace viewer). At least it works reliably now.
    
    Fixes #15482
    
    Change-Id: I993b5f43d22072c6f5bd041ab5888ce176f272b2
    Reviewed-on: https://go-review.googlesource.com/22731
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>

commit ccf2c019921999f49ba2ab8cbfe70ebecc986f46
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Thu May 12 15:27:30 2016 +0300

    go/types: fix certain vet warnings
    
    Updates #11041
    
    Change-Id: I4e1c670d2b7fc04927d77c6f933cee39b7d48b6e
    Reviewed-on: https://go-review.googlesource.com/23083
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a8a2b38fb9e3886b6942621bf4b24ae062f0460b
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu May 12 12:09:18 2016 -0400

    cmd/compile/internal/gc: minor cleanup of init.go comments
    
    Step 5 was deleted in f3575a9 however the numbering of the other
    steps wasn't adjusted accordingly.
    
    While we're here: clean up the whitespace, add curly braces where
    appropriate and delete semicolons.
    
    Change-Id: I4e77b2d3ee8460abe4bfb993674f83e35be8ff17
    Reviewed-on: https://go-review.googlesource.com/23066
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e54dfc2ec4a057aa1bf06f9bef5cdb2e769a669d
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Thu May 12 15:03:22 2016 +0300

    testing: fix vet warning
    
    Updates #11041
    
    Change-Id: I32a381854e6a4fd791db380150efab57e6dfc38c
    Reviewed-on: https://go-review.googlesource.com/23081
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 376e6415402b4e62f96fb7f8f7a99d352aa9c1b3
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Thu May 12 15:13:22 2016 +0300

    cmd: fixed certain vet warnings
    
    Updates #11041
    
    Change-Id: I7f2583d08f344d6622027c5e8a5de1f5d2f2881c
    Reviewed-on: https://go-review.googlesource.com/23082
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a71584975dedd4f4975d65047ec7660191a49613
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Thu May 12 15:00:10 2016 +0300

    reflect: fix vet warnings
    
    Updated #11041
    
    Change-Id: I4a110ba8fefb367a1049b4a65dd20c39eb890ea2
    Reviewed-on: https://go-review.googlesource.com/23080
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
```
