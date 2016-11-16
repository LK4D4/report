# November 16, 2016 Report

Number of commits: 64

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 510.077446ms to 534.499718ms, +4.79%
* github.com/coreos/etcd/cmd/etcd: from 11.965695276s to 11.095815219s, -7.27%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 33.194868885s to 33.887135463s, +2.09%
* github.com/influxdata/influxdb/cmd/influxd: from 6.725966255s to 6.4824467s, -3.62%
* github.com/junegunn/fzf/src/fzf: from 1.019189302s to 1.015741521s, -0.34%
* github.com/mholt/caddy/caddy: from 5.404512096s to 5.509901599s, +1.95%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.267080406s to 1.326506456s, +4.69%
* github.com/nsqio/nsq/apps/nsqd: from 1.85155661s to 1.942376145s, +4.91%
* github.com/prometheus/prometheus/cmd/prometheus: from 38.714590362s to 38.663045996s, -0.13%
* github.com/spf13/hugo: from 6.036109541s to 6.129917214s, +1.55%
* golang.org/x/tools/cmd/guru: from 2.668698026s to 2.572496497s, -3.60%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 3069005 to 3069331, ~
* github.com/coreos/etcd/cmd/etcd: from 25070840 to 25090560, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 29028248 to 29043752, ~
* github.com/influxdata/influxdb/cmd/influxd: from 16428223 to 16441265, ~
* github.com/junegunn/fzf/src/fzf: from 3042352 to 3051384, +0.30%
* github.com/mholt/caddy/caddy: from 14997114 to 15014389, +0.12%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4235000 to 4247740, +0.30%
* github.com/nsqio/nsq/apps/nsqd: from 10085889 to 10103027, +0.17%
* github.com/prometheus/prometheus/cmd/prometheus: from 57450696 to 57467834, ~
* github.com/spf13/hugo: from 15972514 to 15985548, ~
* golang.org/x/tools/cmd/guru: from 7941217 to 7957974, +0.21%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               189           187           -1.06%
BenchmarkMsgpUnmarshal-4             371           371           +0.00%
BenchmarkEasyJsonMarshal-4           1408          1453          +3.20%
BenchmarkEasyJsonUnmarshal-4         1961          1551          -20.91%
BenchmarkFlatBuffersMarshal-4        359           365           +1.67%
BenchmarkFlatBuffersUnmarshal-4      293           294           +0.34%
BenchmarkGogoprotobufMarshal-4       155           157           +1.29%
BenchmarkGogoprotobufUnmarshal-4     248           246           -0.81%

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

* [os/exec: add closeOnce.WriteString method](https://github.com/golang/go/commit/b906df653b58bc2ab9b93e18f62adccc8c1419b7)
* [all: don't call t.Fatal from a goroutine](https://github.com/golang/go/commit/a145890059e9c7aae870e1b9e74b204b6c8bc8d5)
* [text/template: efficient reporting of line numbers](https://github.com/golang/go/commit/24a088d20ad52c527f61b34217da72589e366833)

## GIT Log

```
commit d338f2e1470227afcadde977f2f2ab07d65088db
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Nov 16 13:35:59 2016 +0900

    net: don't run TestTCPBig unconditionally
    
    The test requires tons of memory and results various failures, mainly
    runtime errors and process termination by SIGKILL, caused by resource
    exhaustion when the node under test doesn't have much resources.
    
    This change makes use of -tcpbig flag to enable the test.
    
    Change-Id: Id53fa5d88543e2e60ca9bb4f55a1914ccca844e1
    Reviewed-on: https://go-review.googlesource.com/33254
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1f605175b0044d69ac2364c24f515344c9866fd6
Author: Bryan C. Mills <bcmills@google.com>
Date:   Wed Nov 9 15:28:24 2016 -0500

    runtime/cgo: use libc for sigaction syscalls when possible
    
    This ensures that runtime's signal handlers pass through the TSAN and
    MSAN libc interceptors and subsequent calls to the intercepted
    sigaction function from C will correctly see them.
    
    Fixes #17753.
    
    Change-Id: I9798bb50291a4b8fa20caa39c02a4465ec40bb8d
    Reviewed-on: https://go-review.googlesource.com/33142
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c69233be84eee26d1ea15869b86d76aad602d693
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Nov 16 13:51:45 2016 +0900

    net/http: fix a typo in test
    
    Change-Id: I897237667ffe9e9b2a5f92251a6f665d29479fd2
    Reviewed-on: https://go-review.googlesource.com/33255
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b906df653b58bc2ab9b93e18f62adccc8c1419b7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Nov 15 17:55:28 2016 -0800

    os/exec: add closeOnce.WriteString method
    
    Add an explicit WriteString method to closeOnce that acquires the
    writers lock.  This overrides the one promoted from the
    embedded *os.File field.  The promoted one naturally does not acquire
    the lock, and can therefore race with the Close method.
    
    Fixes #17647.
    
    Change-Id: I3460f2a0d503449481cfb2fd4628b4855ab0ecdf
    Reviewed-on: https://go-review.googlesource.com/33298
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1b66b38e25af567aa30a7b22581c05285be3564a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 15 20:16:15 2016 -0500

    api, doc: update go1.8.txt and next.txt
    
    Both automated updates with a few tweaks.
    
    Change-Id: I24579a8dcc32a84a4fff5c2212681ef30dda61d1
    Reviewed-on: https://go-review.googlesource.com/33297
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 272032d0b245635ca0c5ca4c22e64496174b2ae3
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:51:01 2016 +0200

    runtime: add support files for linux/mips{,le} port
    
    Only exe buildmode without cgo supported.
    
    Change-Id: Id104a79a99d3285c04db00fd98b8affa94ea3c37
    Reviewed-on: https://go-review.googlesource.com/31487
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 4d1fdd8b5e3783b8ca4e9f4fe1e524f5aa83383c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Nov 15 12:41:33 2016 -0800

    test: add test case that failed when built with gccgo
    
    Change-Id: Ie7512cc27436cde53b58686b32a0389849a365e4
    Reviewed-on: https://go-review.googlesource.com/33249
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 816aa99b9a905c5a1fb1d1853698d87b74099720
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Nov 15 16:33:22 2016 +0100

    syscall: define bind flags on Plan 9
    
    These bind flags were removed by mistake in CL 2167.
    
    Fixes #17921.
    
    Change-Id: I1e8089dade30a212b8db0b216c8299946d924d4b
    Reviewed-on: https://go-review.googlesource.com/33271
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 59dc9d7a89829127883dd5e2d8b042f1e5b40336
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Nov 14 15:51:30 2016 -0800

    cmd/cgo: add missing period in comment
    
    Change-Id: I05f31938f3736100bd8b20a150c9fe3a6ffcdeae
    Reviewed-on: https://go-review.googlesource.com/33245
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 27b68474ca8bffea06a4dd11424b293243ae846c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Nov 14 15:51:30 2016 -0800

    cmd/cgo: run cgo pointer checks for pointer to union
    
    If a C union type (or a C++ class type) can contain a pointer field,
    then run the cgo checks on pointers to that type. This will test the
    pointer as though it were an unsafe.Pointer, and will crash if it points
    to Go memory that contains a pointer.
    
    Fixes #15942.
    
    Change-Id: Ic2d07ed9648d4b27078ae7683e26196bcbc59fc9
    Reviewed-on: https://go-review.googlesource.com/33237
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fab3fcaf750445c0016a44376aa81c720133e02c
Author: David Crawshaw <crawshaw@golang.org>
Date:   Sat Nov 12 06:50:24 2016 -0500

    cmd/go: use build ID as plugin symbol prefix
    
    Updates #17821
    
    Change-Id: Iebd2e88b2d4f3d757ffad72456f4bfc0607d8110
    Reviewed-on: https://go-review.googlesource.com/33162
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 03da2690c9fefdbaff613f9ccc224b5f0abfbe16
Author: David Crawshaw <crawshaw@golang.org>
Date:   Sat Nov 12 06:24:36 2016 -0500

    cmd/link, runtime, plugin: versioning
    
    In plugins and every program that opens a plugin, include a hash of
    every imported package.
    
    There are two versions of each hash: one local and one exported.
    As the program starts and plugins are loaded, the first exported
    symbol for each package becomes the canonical version.
    
    Any subsequent plugin's local package hash symbol has to match the
    canonical version.
    
    Fixes #17832
    
    Change-Id: I4e62c8e1729d322e14b1673bada40fa7a74ea8bc
    Reviewed-on: https://go-review.googlesource.com/33161
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a145890059e9c7aae870e1b9e74b204b6c8bc8d5
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Nov 14 21:34:58 2016 -0800

    all: don't call t.Fatal from a goroutine
    
    Fixes #17900.
    
    Change-Id: I42cda6ac9cf48ed739d3a015a90b3cb15edf8ddf
    Reviewed-on: https://go-review.googlesource.com/33243
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9be14c4058287f88dc927ea847e3d6d57ff4047b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 15 03:31:47 2016 +0000

    net: add test that TCP Close unblocks blocked Reads
    
    I guess this was fixed at some point. Remove a skipped test in
    net/http and add an explicit test in net.
    
    Fixes #17695
    
    Change-Id: Idb9f3e236b726bb45098474b830c95c1fce57529
    Reviewed-on: https://go-review.googlesource.com/33242
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b687d6a7886fe7ffea0e36072c2a882c8c485838
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Nov 14 16:50:15 2016 -0800

    misc/cgo/testcarchive, misc/cgo/testcshared: sleep instead of sched_yield
    
    Apparently when GOMAXPROCS == 1 a simple sched_yield in a tight loop is
    not necessarily sufficient to permit a signal handler to run. Instead,
    sleep for 1/1000 of a second.
    
    Fixes #16649.
    
    Change-Id: I83910144228556e742b7a92a441732ef61aa49d9
    Reviewed-on: https://go-review.googlesource.com/33239
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1e917312511f6ce54bbdcea8cd0c25e66973d49e
Author: Caleb Spare <cespare@gmail.com>
Date:   Sun Nov 13 18:06:16 2016 -0800

    html/template: fix multiple Clones of redefined template
    
    This change redoes the fix for #16101 (CL 31092) in a different way by
    making t.Clone return the template associated with the t.Name() while
    allowing for the case that a template of the same name is define-d.
    
    Fixes #17735.
    
    Change-Id: I1e69672390a4c81aa611046a209008ae4a3bb723
    Reviewed-on: https://go-review.googlesource.com/33210
    Run-TryBot: Caleb Spare <cespare@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 91461002f354a5fc8b3898bbb0ade90d2921441e
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Mon Nov 14 09:17:15 2016 +0900

    os: gofmt -w -s
    
    Change-Id: I9a42cb55544185ade20b2a4a9de5d39a6cfc6fc6
    Reviewed-on: https://go-review.googlesource.com/33172
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 90d536f3ca2d6456ee0ab45408e8499ea815ddb3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 15 01:36:34 2016 +0000

    net/http: update bundled http2 for write scheduling order fix
    
    Updates x/net/http2 to x/net git rev 00ed5e9 for:
    
        http2: schedule RSTStream writes onto its stream's queue
        https://golang.org/cl/33238
    
    Fixes #17243
    
    Change-Id: I79cc5d15bf69ead28d549d4f798c12f4ee2a2201
    Reviewed-on: https://go-review.googlesource.com/33241
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 24a088d20ad52c527f61b34217da72589e366833
Author: Rob Pike <r@golang.org>
Date:   Tue Nov 8 15:43:20 2016 -0800

    text/template: efficient reporting of line numbers
    
    Instead of scanning the text to count newlines, which is n², keep track as we go
    and store the line number in the token.
    
    benchmark                 old ns/op      new ns/op     delta
    BenchmarkParseLarge-4     1589721293     38783310      -97.56%
    
    Fixes #17851
    
    Change-Id: I231225c61e667535e2ce55cd2facea6d279cc59d
    Reviewed-on: https://go-review.googlesource.com/33234
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit bb00a8d97faa70bf7a1cbdd4a43e95347a9c8709
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 14 22:11:05 2016 +0000

    net/http: update bundled http2, add TestServerKeepAlivesEnabled h1/h2 tests
    
    Updates x/net/http2 to x/net git rev 6dfeb344 for:
    
       http2: make Server respect http1 Server's SetKeepAlivesEnabled
       https://golang.org/cl/33153
    
    And adds a test in std.
    
    Fixes #17717
    
    Change-Id: I3ba000abb6f3f682261e105d8a4bb93bde6609fe
    Reviewed-on: https://go-review.googlesource.com/33231
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Tom Bergan <tombergan@google.com>

commit b83350a2e05d63eaae8da9ff4f957ab44e4cb9d9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 14 22:23:10 2016 +0000

    Revert "text/template: efficient reporting of line numbers"
    
    This reverts commit 794fb71d9c1018c4beae1657baca5229e6a02ad0.
    
    Reason for revert: submitted without TryBots and it broke all three race builders.
    
    Change-Id: I80a1e566616f0ee8fa3529d4eeee04268f8a713b
    Reviewed-on: https://go-review.googlesource.com/33232
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2442b49c47aa818bbc55e4c064e9ea0ca058735f
Author: Marcel Edmund Franke <marcel.edmund.franke@gmail.com>
Date:   Mon Nov 14 21:46:25 2016 +0100

    html/template: typo fix
    
    comment on unexported function starts with wrong functionname
    
    Change-Id: Ib16c2fe42b5a8d4606ed719f620923c17839d091
    Reviewed-on: https://go-review.googlesource.com/33203
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 794fb71d9c1018c4beae1657baca5229e6a02ad0
Author: Rob Pike <r@golang.org>
Date:   Tue Nov 8 15:43:20 2016 -0800

    text/template: efficient reporting of line numbers
    
    Instead of scanning the text to count newlines, which is n², keep track as we go
    and store the line number in the token.
    
    benchmark                 old ns/op      new ns/op     delta
    BenchmarkParseLarge-4     1589721293     38783310      -97.56%
    
    Fixes #17851
    
    Change-Id: Ieaf89a35e371b405ad92e38baa1e3fa98d18cfb4
    Reviewed-on: https://go-review.googlesource.com/32923
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 2f76c1985fa8bbb0fb09af6600445b5c7d4d5cb4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Nov 13 15:18:42 2016 -0800

    cmd/go/testdata/src: gofmt
    
    These are functionality tests, not formatter tests.
    
    I also tested manually that 'go test cmd/go'
    without -short still passes.
    
    
    Change-Id: Id146e1dc3b65e19ea531869725cd0b97f4801b8b
    Reviewed-on: https://go-review.googlesource.com/33169
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5f74ce394f02714a88dd375f54e8709ce58d1805
Author: Jesse Szwedko <jesse.szwedko@gmail.com>
Date:   Sun Nov 13 13:29:19 2016 -0800

    syscall: Clearenv now unsets env vars on Windows
    
    Previously, `os.Clearenv()` (by way of `syscall.Clearenv`) would simply
    set all environment variables' values to `""` rather than actually
    unsetting them causing subsequent `os.LookupEnv` calls to return that
    they were still set.
    
    Fixes #17902
    
    Change-Id: I54081b4b98665e9a39f55ea7582c8d40bb8a2a22
    Reviewed-on: https://go-review.googlesource.com/33168
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 4a381e3ee3ad6104fc1e1bc255f29d6bf47d7969
Author: David du Colombier <0intro@gmail.com>
Date:   Sun Nov 13 19:12:39 2016 +0100

    net/http: enable timeout tests on Plan 9
    
    Deadlines have been implemented on Plan 9 in CL 31521.
    
    Enable the following tests:
    
     - TestServerTimeouts
     - TestOnlyWriteTimeout
     - TestTLSHandshakeTimeout
     - TestIssue4191_InfiniteGetTimeout
     - TestIssue4191_InfiniteGetToPutTimeout
    
    Updates #7237.
    
    Change-Id: If5e75cfaa9133dcf9ce6aac9fc2badafc1612b64
    Reviewed-on: https://go-review.googlesource.com/33197
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8d3d23a124970475485c50385fc3c13780c9c306
Author: David du Colombier <0intro@gmail.com>
Date:   Sun Nov 13 22:00:38 2016 +0100

    net/http: fix error message in TestClientWriteShutdown
    
    Change-Id: I3c664201baef6d7dbed94dab63db0ac974bf6817
    Reviewed-on: https://go-review.googlesource.com/33198
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f8dc4f20f8f8cb2ac8ee14b15ed4bf5201b61e81
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Nov 13 12:59:07 2016 -0800

    Revert "cmd/vet: ignore printf failures in cmd"
    
    This reverts commit f15915af4effbbe6895ae69be02d22ac016927d5.
    
    CL 32851 fixed cmd/vet's handling of fmt.Formatter.
    
    Updates #17057.
    
    Change-Id: I3409100d16037645946fe7fe78fbb173e1648494
    Reviewed-on: https://go-review.googlesource.com/33166
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7c9f910607afbb5e33146d966e6c60ac5dbf3b31
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Nov 13 13:10:42 2016 -0800

    all: fix vet nits
    
    Fixes these vet complaints:
    
    net/error_test.go:254: unrecognized printf flag for verb 'T': '#'
    os/os_test.go:1067: arg mt for printf verb %d of wrong type: time.Time
    runtime/debug/garbage_test.go:83: arg dt for printf verb %d of wrong type: time.Time
    
    Change-Id: I0e986712a4b083b75fb111e687e424d06a85a47b
    Reviewed-on: https://go-review.googlesource.com/33167
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit afa68b36cc225075e87b53e2f9c2edc9dfb73b9e
Author: Thordur Bjornsson <thorduri@secnorth.net>
Date:   Sat Nov 12 17:05:17 2016 +0100

    encoding/hex: Document DecodedLen.
    
    Mention that it specifically returns x / 2, and do the same for
    EncodedLen.
    
    Change-Id: Ie334f5abecbc487caf4965abbcd14442591bef2a
    Change-Id: Idfa413faad487e534489428451bf736b009293d6
    Reviewed-on: https://go-review.googlesource.com/33191
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 582a421a8c245268f7c081fa1191374d50a601f7
Author: David du Colombier <0intro@gmail.com>
Date:   Sun Nov 13 15:04:54 2016 +0100

    net: enable timeout tests on Plan 9
    
    Deadlines have been implemented on Plan 9 in CL 31521.
    
    Enable the following tests:
    
     - TestReadTimeout
     - TestReadFromTimeout
     - TestWriteTimeout
     - TestWriteToTimeout
     - TestReadTimeoutFluctuation
     - TestVariousDeadlines
     - TestVariousDeadlines1Proc
     - TestVariousDeadlines4Proc
     - TestReadWriteDeadlineRace
    
    Change-Id: I221ed61d55f7f1e4345b37af6748c04e1e91e062
    Reviewed-on: https://go-review.googlesource.com/33196
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 662d25351542d624f397bb0a83ff1c0ca38b6428
Author: Dhananjay Nakrani <dhananjaynakrani@gmail.com>
Date:   Sun Nov 6 19:56:14 2016 -0800

    cmd/vet: ignore unrecognized verbs for fmt.Formatter
    
    Updates #17057.
    
    Change-Id: I54c838d3a44007d4023754e42971e91bfb5e8612
    Reviewed-on: https://go-review.googlesource.com/32851
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 524cd4855e11a26d07a7ca7ee45d3bee38c54425
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Sat Nov 12 20:08:02 2016 +0100

    time: simplify stringification of Month
    
    Simplifies https://golang.org/cl/33145
    which fixed #17720.
    
    Change-Id: Ib922d493cdc5920832dc95b55094796baca7243e
    Reviewed-on: https://go-review.googlesource.com/33194
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4bee9012b31a762799ea861cd5f55583f1f602f5
Author: David du Colombier <0intro@gmail.com>
Date:   Sat Nov 12 22:12:22 2016 +0100

    net/http/httptest: remove workaround on Plan 9
    
    This issue has been fixed in CL 31390.
    
    Change-Id: I0c2425fd33be878037d10d612a50116a7b693431
    Reviewed-on: https://go-review.googlesource.com/33195
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 091ba60bd87bf8b40a24196f31b5f3c18c2f2dc4
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Thu Sep 22 13:33:22 2016 -0700

    compress/flate: add examples
    
    Updates #16360
    
    Change-Id: I66ff23e0501363f58fe891d5e95806422071f93b
    Reviewed-on: https://go-review.googlesource.com/30162
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2f497263e4ff6121a1ba80e7a57e950061896626
Author: Patrick Lee <pattyshack101@gmail.com>
Date:   Fri Nov 11 19:24:07 2016 -0800

    cmd/pprof: add options to skip tls verification
    
    Don't verify tls host when profiling https+insecure://host/port/...,
    as per discussion in https://go-review.googlesource.com/#/c/20885/.
    
    Fixes: #11468
    
    Change-Id: Ibfc236e5442a00339334602a4014e017c62d9e7a
    Reviewed-on: https://go-review.googlesource.com/33157
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4966150af0139b5f26a943eb4c33fe5cb6758043
Author: David du Colombier <0intro@gmail.com>
Date:   Sat Nov 12 17:01:08 2016 +0100

    net: enable TestReadTimeoutUnblocksRead on Plan 9
    
    Deadlines have been implemented on Plan 9 in CL 31521.
    
    Fixes #17477.
    
    Change-Id: Icb742ac30933b6d2f9350fc4e6acbcd433c66c21
    Reviewed-on: https://go-review.googlesource.com/33190
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 37020dd510cf7da36e9eb1827a20890234e4ea79
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Nov 12 05:58:52 2016 +0000

    runtime/internal/atomic: add TestUnaligned64
    
    Add a variant of sync/atomic's TestUnaligned64 to
    runtime/internal/atomic.
    
    Skips the test on arm for now where it's currently failing.
    
    Updates #17786
    
    Change-Id: If63f9c1243e9db7b243a95205b2d27f7d1dc1e6e
    Reviewed-on: https://go-review.googlesource.com/33159
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c921d8f39d6da1afd1550787464d27f015412194
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 11 23:40:25 2016 +0000

    context: document appropriate WithValue key type more
    
    Fixes #17826
    Updates #17302
    
    Change-Id: I7c1ebd965e679e7169a97e62d27ae3ede2473aa1
    Reviewed-on: https://go-review.googlesource.com/33152
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit adb384ad2cfbd742fb106b7ec2a65d4ad844c35e
Author: David du Colombier <0intro@gmail.com>
Date:   Fri Nov 11 20:49:11 2016 +0100

    net: implement asynchonous cancelable I/O on Plan 9
    
    This change is an experimental implementation of asynchronous
    cancelable I/O operations on Plan 9, which are required to
    implement deadlines.
    
    There are no asynchronous syscalls on Plan 9. I/O operations
    are performed with blocking pread and pwrite syscalls.
    
    Implementing deadlines in Go requires a way to interrupt
    I/O operations.
    
    It is possible to interrupt reads and writes on a TCP connection
    by forcing the closure of the TCP connection. This approach
    has been used successfully in CL 31390.
    
    However, we can't implement deadlines with this method, since
    we require to be able to reuse the connection after the timeout.
    
    On Plan 9, I/O operations are interrupted when the process
    receives a note. We can rely on this behavior to implement
    a more generic approach.
    
    When doing an I/O operation (read or write), we start the I/O in
    its own process, then wait for the result asynchronously. The
    process is able to handle the "hangup" note. When receiving the
    "hangup" note, the currently running I/O operation is canceled
    and the process returns.
    
    This way, deadlines can be implemented by sending an "hangup"
    note to the process running the blocking I/O operation, after
    the expiration of a timer.
    
    Fixes #11932.
    Fixes #17498.
    
    Change-Id: I414f72c7a9a4f9b8f9c09ed3b6c269f899d9b430
    Reviewed-on: https://go-review.googlesource.com/31521
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 456f2f5cb8c2e06e7faba6ba298ffb65c7a19397
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Nov 10 17:10:45 2016 -0800

    time: use 1e9 rather than 1e-9 in Duration calculations
    
    1e-9 has a 1 in the last place, causing some Duration calculations to
    have unnecessary rounding errors.  1e9 does not, so use that instead.
    
    Change-Id: I96334a2c47e7a014b532eb4b8a3ef9550e7ed057
    Reviewed-on: https://go-review.googlesource.com/33116
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5b147122d6094b792a027b892884b994fe77a4d6
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Wed Oct 19 17:55:10 2016 +0200

    cmd/dist: add support for GOARCH=mips{,le}
    
    Change-Id: I6e24d22eada190e9aa2adc161be7a753c8e5054b
    Reviewed-on: https://go-review.googlesource.com/31514
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 241dccc4fdaf32830fa32a38a9c7953584b57a17
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Sat Sep 17 15:01:55 2016 +0100

    cmd/internal/browser: add chromium to the list of browsers
    
    Many linux distros distribute Chromium instead of Chrome.
    
    Change-Id: I5474d94da28a7c79bdd7181f77472d4ce73bb225
    Reviewed-on: https://go-review.googlesource.com/29293
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d8264de8683dac99ffbbbc1f46415e627b73c9ed
Author: Dmitri Shuralyov <shurcooL@gmail.com>
Date:   Wed Nov 9 14:49:12 2016 -0800

    all: spell "marshal" and "unmarshal" consistently
    
    The tree is inconsistent about single l vs double l in those
    words in documentation, test messages, and one error value text.
    
            $ git grep -E '[Mm]arshall(|s|er|ers|ed|ing)' | wc -l
                  42
            $ git grep -E '[Mm]arshal(|s|er|ers|ed|ing)' | wc -l
                1694
    
    Make it consistently a single l, per earlier decisions. This means
    contributors won't be confused by misleading precedence, and it helps
    consistency.
    
    Change the spelling in one error value text in newRawAttributes of
    crypto/x509 package to be consistent.
    
    This change was generated with:
    
            perl -i -npe 's,([Mm]arshal)l(|s|er|ers|ed|ing),$1$2,' $(git grep -l -E '[Mm]arshall' | grep -v AUTHORS | grep -v CONTRIBUTORS)
    
    Updates #12431.
    Follows https://golang.org/cl/14150.
    
    Change-Id: I85d28a2d7692862ccb02d6a09f5d18538b6049a2
    Reviewed-on: https://go-review.googlesource.com/33017
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9a78eadeb636689f79dbf6bd3c0a35d830678097
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 11 14:05:51 2016 -0800

    net: deflake TestTCPSupriousConnSetupCompletion [sic]
    
    And rename it.
    
    Fixes #17703
    
    Change-Id: I73c82a9b3f96180699c6d33c069a666018eb30f9
    Reviewed-on: https://go-review.googlesource.com/33149
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 02d79e95581e53edca21f0c105675b3f459ee736
Author: Quentin Smith <quentin@golang.org>
Date:   Thu Nov 10 11:45:56 2016 -0500

    cmd/go: skip TestCgoPkgConfig if pkg-config is too old
    
    pkg-config 0.24 adds support for quoting and escaping whitespace;
    distros like CentOS 6 are still shipping pkg-config 0.23. Skip the test
    there since there's no way to get whitespace into the pkg-config output.
    
    Fixes #17846.
    
    Change-Id: Ie4ea17e9b709372a20178b539498929754bcd51f
    Reviewed-on: https://go-review.googlesource.com/33027
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit a18b4b3fb9b1b98f6eefa038b723f99fd6d13efd
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 11 20:54:07 2016 +0000

    time: don't panic stringifying the zero Month
    
    Fixes #17720
    
    Change-Id: Ib95c230deef3934db729856c17908f8e5a1e2b7f
    Reviewed-on: https://go-review.googlesource.com/33145
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit e0aedfb496f414e1a869d27175b4dfcc6baef407
Author: Rhys Hiltner <rhys@justin.tv>
Date:   Tue Oct 25 11:56:29 2016 -0700

    runtime: include pre-panic/throw logs in core dumps
    
    When a Go program crashes with GOTRACEBACK=crash, the OS creates a
    core dump. Include the text-formatted output of some of the cause of
    that crash in the core dump.
    
    Output printed by the runtime before crashing is maintained in a
    circular buffer to allow access to messages that may be printed
    immediately before calling runtime.throw.
    
    The stack traces printed by the runtime as it crashes are not stored.
    The information required to recreate them should be included in the
    core file.
    
    Updates #16893
    
    There are no tests covering the generation of core dumps; this change
    has not added any.
    
    This adds (reentrant) locking to runtime.gwrite, which may have an
    undesired performance impact.
    
    Change-Id: Ia2463be3c12429354d290bdec5f3c8d565d1a2c3
    Reviewed-on: https://go-review.googlesource.com/32013
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 10d2efd0b0e3f4f92f9470435f63211cbeb82008
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 11 20:35:26 2016 +0000

    net/smtp: make Client.Auth trim final space if Auth.Start toServer is empty
    
    Users can implement the smtp.Auth interface and return zero bytes in
    the "toServer []byte" return value from the Auth.Start method. People
    apparently do this to implement the SMTP "LOGIN" method.
    
    But we were then sending "AUTH LOGIN \r\n" to the server, which some
    servers apparently choke on. So, trim it when the toServer value is
    empty.
    
    Fixes #17794
    
    Change-Id: I83662dba9e0f61b1c5000396c096cf7110f78361
    Reviewed-on: https://go-review.googlesource.com/33143
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e6da64b6c070eeb872ca141ab58435e7be5da388
Author: Russ Cox <rsc@golang.org>
Date:   Fri Nov 11 10:27:36 2016 -0500

    runtime: fix Windows profiling crash
    
    I don't have any way to test or reproduce this problem,
    but the current code is clearly wrong for Windows.
    Make it better.
    
    As I said on #17165:
    
    But the borrowing of M's and the profiling of M's by the CPU profiler
    seem not synchronized enough. This code implements the CPU profiler
    on Windows:
    
            func profileloop1(param uintptr) uint32 {
                    stdcall2(_SetThreadPriority, currentThread, _THREAD_PRIORITY_HIGHEST)
    
                    for {
                            stdcall2(_WaitForSingleObject, profiletimer, _INFINITE)
                            first := (*m)(atomic.Loadp(unsafe.Pointer(&allm)))
                            for mp := first; mp != nil; mp = mp.alllink {
                                    thread := atomic.Loaduintptr(&mp.thread)
                                    // Do not profile threads blocked on Notes,
                                    // this includes idle worker threads,
                                    // idle timer thread, idle heap scavenger, etc.
                                    if thread == 0 || mp.profilehz == 0 || mp.blocked {
                                            continue
                                    }
                                    stdcall1(_SuspendThread, thread)
                                    if mp.profilehz != 0 && !mp.blocked {
                                            profilem(mp)
                                    }
                                    stdcall1(_ResumeThread, thread)
                            }
                    }
            }
    
            func profilem(mp *m) {
                    var r *context
                    rbuf := make([]byte, unsafe.Sizeof(*r)+15)
    
                    tls := &mp.tls[0]
                    gp := *((**g)(unsafe.Pointer(tls)))
    
                    // align Context to 16 bytes
                    r = (*context)(unsafe.Pointer((uintptr(unsafe.Pointer(&rbuf[15]))) &^ 15))
                    r.contextflags = _CONTEXT_CONTROL
                    stdcall2(_GetThreadContext, mp.thread, uintptr(unsafe.Pointer(r)))
                    sigprof(r.ip(), r.sp(), 0, gp, mp)
            }
    
            func sigprof(pc, sp, lr uintptr, gp *g, mp *m) {
                    if prof.hz == 0 {
                            return
                    }
    
                    // Profiling runs concurrently with GC, so it must not allocate.
                    mp.mallocing++
    
                    ... lots of code ...
    
                    mp.mallocing--
            }
    
    A borrowed M may migrate between threads. Between the
    atomic.Loaduintptr(&mp.thread) and the SuspendThread, mp may have
    moved to a new thread, so that it's in active use. In particular
    it might be calling malloc, as in the crash stack trace. If so, the
    mp.mallocing++ in sigprof would provoke the crash.
    
    Those lines are trying to guard against allocation during sigprof.
    But on Windows, mp is the thread being traced, not the current
    thread. Those lines should really be using getg().m.mallocing, which
    is the same on Unix but not on Windows. With that change, it's
    possible the race on the actual thread is not a problem: the traceback
    would get confused and eventually return an error, but that's fine.
    The code expects that possibility.
    
    Fixes #17165.
    
    Change-Id: If6619731910d65ca4b1a6e7de761fa2518ef339e
    Reviewed-on: https://go-review.googlesource.com/33132
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b6a15683f0c4d177b3711b55724506aebb03f764
Author: Bill O'Farrell <billo@ca.ibm.com>
Date:   Sun Oct 30 00:11:37 2016 -0400

    math: use SIMD to accelerate some scalar math functions on s390x
    
    Note, most math functions are structured to use stubs, so that they can
    be accelerated with assembly on any platform.
    Sinh, cosh, and tanh were not structued with stubs, so this CL does
    that. This set of routines was chosen as likely to produce good speedups
    with assembly on any platform.
    
    Technique used was minimax polynomial approximation using tables of
    polynomial coefficients, with argument range reduction.
    A table of scaling factors was also used for cosh and log10.
    
                         before       after      speedup
    BenchmarkCos         22.1 ns/op   6.79 ns/op  3.25x
    BenchmarkCosh       125   ns/op  11.7  ns/op 10.68x
    BenchmarkLog10       48.4 ns/op  12.5  ns/op  3.87x
    BenchmarkSin         22.2 ns/op   6.55 ns/op  3.39x
    BenchmarkSinh       125   ns/op  14.2  ns/op  8.80x
    BenchmarkTanh        65.0 ns/op  15.1  ns/op  4.30x
    
    Accuracy was tested against a high precision
    reference function to determine maximum error.
    Approximately 4,000,000 points were tested for each function,
    producing the following result.
    Note: ulperr is error in "units in the last place"
    
           max
          ulperr
    sin    1.43 (returns NaN beyond +-2^50)
    cos    1.79 (returns NaN beyond +-2^50)
    cosh   1.05
    sinh   3.02
    tanh   3.69
    log10  1.75
    
    Also includes a set of tests to test non-vector functions even
    when SIMD is enabled
    
    Change-Id: Icb45f14d00864ee19ed973d209c3af21e4df4edc
    Reviewed-on: https://go-review.googlesource.com/32352
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit 9f9d83404f938a0dfb98d3f4a4d420261606069a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 11 18:43:39 2016 +0000

    net/http: make Server respect shutdown state after handler finishes
    
    If the Server's Shutdown (or SetKeepAlivesEnabled) method was called
    while a connection was in a Handler, but after the headers had been
    written, the connection was not later closed.
    
    Fixes #9478
    Updates #17754 (reverts that workaround)
    
    Change-Id: I65324ab8217373fbb38e12e2b8bffd0a91806072
    Reviewed-on: https://go-review.googlesource.com/33141
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 39e3cbfff623d7941b519f9d35883ef3b74cbcd6
Author: Russ Cox <rsc@golang.org>
Date:   Fri Nov 11 11:54:18 2016 -0500

    text/template: reintroduce implicit indirect of interface values in builtin funcs
    
    CL 31462 made it possible to operate directly on reflect.Values
    instead of always forcing a round trip to interface{} and back.
    The round trip was losing addressability, which hurt users.
    
    The round trip was also losing "interface-ness", which helped users.
    That is, using reflect.ValueOf(v.Interface()) instead of v was doing
    an implicit indirect any time v was itself an interface{} value: the result
    was the reflect.Value for the underlying concrete value contained in the
    interface, not the interface itself.
    
    CL 31462 eliminated some "unnecessary" reflect.Value round trips
    in order to preserve addressability, but in doing so it lost this implicit
    indirection. This CL adds the indirection back.
    
    It may help to compare the changes in this CL against funcs.go from CL 31462:
    https://go-review.googlesource.com/#/c/31462/4/src/text/template/funcs.go
    
    Everywhere CL 31462 changed 'v := reflect.ValueOf(x)' to 'v := x',
    this CL changes 'v := x' to 'v := indirectInterface(x)'.
    
    Fixes #17714.
    
    Change-Id: I67cec4eb41fed1d56e1c19f12b0abbd0e59d35a2
    Reviewed-on: https://go-review.googlesource.com/33139
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit fabb4115ed311ec4af19f87c2334e38497dbb8d0
Author: Russ Cox <rsc@golang.org>
Date:   Fri Nov 11 10:11:19 2016 -0500

    time: update Timer.Stop doc to account for AfterFunc
    
    Fixes #17600.
    
    Change-Id: I7aa0eb0dd959da031b6039b51f07db668d4fb468
    Reviewed-on: https://go-review.googlesource.com/33131
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Gudger <igudger@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 84ded8ba8a0233a7f38e3c777dc1c213f98d00a2
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Fri Aug 26 22:21:00 2016 +0100

    net/http: make Server log on bad requests from clients
    
    Fixes #12745
    
    Change-Id: Iebb7c97cb5b68dc080644d796a6ca1c120d41b26
    Reviewed-on: https://go-review.googlesource.com/27950
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 238247eb59cdddeedaa5c5db67734df2cd1049ab
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 11 17:14:02 2016 +0000

    net/http: deflake new TestInterruptWithPanic_h2
    
    TestInterruptWithPanic_h2 was added yesterday in
    https://golang.org/cl/33099 and https://golang.org/cl/33103
    
    Deflake it. The http2 server sends an error before logging.
    
    Rather than reorder the http2 code to log before writing the RSTStream
    frame, just loop for a bit waiting for the condition we're
    expecting.
    
    This goes from 2 in 500 flakes for me to unreproducible.
    
    Change-Id: I062866a5977f50c820965aaf83882ddd7bf98f91
    Reviewed-on: https://go-review.googlesource.com/33140
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 866e01457f69b58b31ccb95f223aac80e1285332
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 26 20:44:26 2016 -0400

    net: apply tcp4/tcp6 restrictions to literals in ResolveTCPAddr
    
    The restrictions were already being applied to the IP addresses
    received from the host resolver. Apply the same restrictions to
    literal IP addresses not passed to the host resolver.
    
    For example, ResolveTCPAddr("tcp4", "[2001:db8::1]:http") used
    to succeed and now does not (that's not an IPv4 address).
    
    Perhaps a bit surprisingly,
    ResolveTCPAddr("tcp4", "[::ffff:127.0.0.1]:http") succeeds,
    behaving identically to ResolveTCPAddr("tcp4", "127.0.0.1:http"), and
    ResolveTCPAddr("tcp6", "[::ffff:127.0.0.1]:http") fails,
    behaving identically to ResolveTCPAddr("tcp6", "127.0.0.1:http").
    Even so, it seems right to match (by reusing) the existing filtering
    as applied to addresses resolved by the host C library.
    If anyone can make a strong argument for changing the filtering
    of IPv4-inside-IPv6 addresses, the fix can be applied to all
    the code paths in a separate CL.
    
    Fixes #14037.
    
    Change-Id: I690dfdcbe93d730e11e00ea387fa7484cd524341
    Reviewed-on: https://go-review.googlesource.com/32100
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c4099c7593fea9c594ad0e8969452a741acba4c7
Author: Russ Cox <rsc@golang.org>
Date:   Fri Nov 11 11:25:15 2016 -0500

    runtime/pprof: delete new TestCPUProfileParse
    
    All the existing CPU profiler tests already parse the profile.
    That should be sufficient indication that profiles can be parsed.
    
    Fixes #17853.
    
    Change-Id: Ie8a190e2ae4eef125c8eb0d4e8b7adac420abbdb
    Reviewed-on: https://go-review.googlesource.com/33136
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit eafe48781a013554961b9bcad8a0262f9ea20acc
Author: Michael Matloob <matloob@golang.org>
Date:   Fri Nov 11 11:32:07 2016 -0500

    runtime/pprof/internal: delete package gzip0
    
    rsc's change golang.org/cl/32455 added a mechanism
    that allows pprof to depend on gzip without introducing
    an import cycle. This obsoletes the need for the gzip0
    package, which was created solely to remove the need
    for that dependency.
    
    Change-Id: Ifa3b98faac9b251f909b84b4da54742046c4e3ad
    Reviewed-on: https://go-review.googlesource.com/33137
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8eb88b0d8e1b76eabe37de67b7dd603da9fbaea3
Author: Kevin Burke <kev@inburke.com>
Date:   Fri Nov 11 08:39:33 2016 -0800

    cmd/gofmt, crypto/tls: fix typos
    
    Fix spelling of "original" and "occurred" in new gofmt docs. The same
    misspelling of "occurred" was also present in crypto/tls, I fixed it there as
    well.
    
    Change-Id: I67b4f1c09bd1a2eb1844207d5514f08a9f525ff9
    Reviewed-on: https://go-review.googlesource.com/33138
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8f215d8c1f90f2841df2ae319799e0bc22320751
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Nov 11 07:39:17 2016 -0800

    cmd/vet/all: add bitwidths for mips and mipsle
    
    cmd/vet/all still doesn't run for mips/mipsle,
    because the rest of the toolchain doesn't yet
    fully support it.
    
    Change-Id: I1a86b0edddbdcd5f43e752208508d99da7aabbb3
    Reviewed-on: https://go-review.googlesource.com/33134
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit eb8f2a832078d0748bcf1eb38357c8119fdbccb7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Fri Nov 11 07:37:32 2016 -0800

    all: fix vet nits
    
    Fixes these complaints from vet:
    
    cmd/compile/internal/gc/noder.go:32: cmd/compile/internal/syntax.Error composite literal uses unkeyed fields
    cmd/compile/internal/gc/noder.go:1035: cmd/compile/internal/syntax.Error composite literal uses unkeyed fields
    cmd/compile/internal/gc/noder.go:1051: cmd/compile/internal/syntax.Error composite literal uses unkeyed fields
    cmd/compile/internal/syntax/parser_test.go:182: possible formatting directive in Error call
    net/http/client_test.go:1334: possible formatting directive in Fatal call
    
    Change-Id: I5f90ec30f3c106c7e66c92e2b6f8d3b4874fec66
    Reviewed-on: https://go-review.googlesource.com/33133
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 50fed64dd90ca6a58cfe8fa7c1061aa8666cc76f
Author: Keegan Carruthers-Smith <keegan.csmith@gmail.com>
Date:   Fri Nov 4 09:31:35 2016 +0200

    go/doc: don't panic if method is missing recv type
    
    Fixes #17788
    
    Change-Id: I2f8a11321dc8f10bebbc8df90ba00ec65b9ee0fa
    Reviewed-on: https://go-review.googlesource.com/32790
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 9a5bddd7ed57596a259f3896dd31ea30e331027d
Author: Richard Gibson <richard.gibson@gmail.com>
Date:   Sat Oct 22 00:21:18 2016 -0400

    net: bring domain name length checks into RFC compliance
    
    The 255-octet limit applies to wire format, not presentation format.
    
    Fixes #17549
    
    Change-Id: I2b5181c53fba32fea60178e0d8df9114aa992b55
    Reviewed-on: https://go-review.googlesource.com/31722
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
```
