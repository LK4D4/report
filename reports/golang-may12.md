# May 12, 2016 Report

Number of commits: 90

## Compilation time

* github.com/coreos/etcd: from 11.676090712s to 11.276484699s, -3.42%
* github.com/boltdb/bolt/cmd/bolt: from 554.564327ms to 552.10216ms, -0.44%
* github.com/gogits/gogs: from 12.904094088s to 12.244053491s, -5.11%
* github.com/spf13/hugo: from 6.683870631s to 6.458126868s, -3.38%
* github.com/influxdata/influxdb/cmd/influxd: from 6.425585765s to 6.280917177s, -2.25%

## Binary size:

* github.com/coreos/etcd: from 21794584 to 21822520, +0.13%
* github.com/boltdb/bolt/cmd/bolt: from 2565367 to 2565367, ~
* github.com/gogits/gogs: from 22713777 to 22752019, +0.17%
* github.com/spf13/hugo: from 14628785 to 14666866, +0.26%
* github.com/influxdata/influxdb/cmd/influxd: from 15640645 to 15685855, +0.29%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               294           289           -1.70%
BenchmarkMsgpUnmarshal-4             621           628           +1.13%
BenchmarkEasyJsonMarshal-4           2225          2229          +0.18%
BenchmarkEasyJsonUnmarshal-4         2394          2388          -0.25%
BenchmarkFlatBuffersMarshal-4        570           566           -0.70%
BenchmarkFlatBuffersUnmarshal-4      449           453           +0.89%
BenchmarkGogoprotobufMarshal-4       255           253           -0.78%
BenchmarkGogoprotobufUnmarshal-4     388           389           +0.26%

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

* [runtime: print signal name in panic, if name is known](1a7fc2357b1c26dcdf4fa57dee67a1172696801f)
* [crypto/sha256: Use AVX2 if possible](2210d88a889e0ea463bcdef2b658aaec1050cf01)
* [crypto/sha1: Add AVX2 version for AMD64](fafadc521ede90f8abed73e8d209e130c456e983)
* [time: don't depend on the io package](d88261fb6581106e4e7d8d6c63f0e33c2a24361e)
* [encoding/json: support maps with integer keys](f05c3aa24d815cd3869153750c9875e35fc48a6e)
* [crypto/sha1: disable crashing AVX2 optimizations for now](78ff74375930d5ae391beae562c91da40e5d92a4)


## GIT Log

```
commit 5bd37b8e78980beed2861bbdc7f8f28fc3f72671
Author: Joel Sing <joel@sing.id.au>
Date:   Mon May 9 02:13:03 2016 +1000

    runtime: stop using sigreturn on openbsd/386
    
    In future releases of OpenBSD, the sigreturn syscall will no longer
    exist. As such, stop using sigreturn on openbsd/386 and just return
    from the signal trampoline (as we already do for openbsd/amd64 and
    openbsd/arm).
    
    Change-Id: Ic4de1795bbfbfb062a685832aea0d597988c6985
    Reviewed-on: https://go-review.googlesource.com/23024
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 81b70f3751374ccd1eda2f536156dd91cd9f9c9b
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed May 11 16:16:37 2016 +1000

    syscall: make mksyscall_windows.go -systemdll flag true by default
    
    Updates #15167
    
    Change-Id: I826f67e75011ba79325a1294ac0d70d7c6a3e32f
    Reviewed-on: https://go-review.googlesource.com/23022
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 21d781070cea6001ee541933ed76dc6da96bde4c
Author: Robert Griesemer <gri@golang.org>
Date:   Wed May 11 13:39:36 2016 -0700

    cmd/compile: use one format for exporting calls of builtin functions
    
    Minor cleanup. Each of these cases appears both during export and
    import when running all.bash and thus is tested by all.bash.
    
    Change-Id: Iaa4a5a5b163cefe33e43d08d396e02a02e5c22a5
    Reviewed-on: https://go-review.googlesource.com/23060
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit eb9062b7bf77a5051b3db6f3e944739a486ca1e4
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed May 11 15:01:28 2016 -0700

    net/http: keep HTTP/1.0 keep-alive conns open if response can't have a body
    
    Fixes #15647
    
    Change-Id: I588bfa4eb336d1da1fcda8d06e32ed13c0b51c70
    Reviewed-on: https://go-review.googlesource.com/23061
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 114051aa1d5da5dec1b2707b1403261a3135b9b5
Author: Johan Sageryd <j@1616.se>
Date:   Sun May 8 18:06:03 2016 +0200

    text/template: fix typo in documentation
    
    Change-Id: I4ccfaa16e153aad001d670891b3848264e63cf6f
    Reviewed-on: https://go-review.googlesource.com/23031
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ef62f641c37431a870fa093c43b3ee51a06db0da
Author: Robert Griesemer <gri@golang.org>
Date:   Wed May 11 12:40:17 2016 -0700

    cmd/compile: use ONAME instead of OPACK in binary export format
    
    This is addressing feedback given on golang.org/cl/23052;
    we do it in a separate CL to separate the functional from
    the rename change.
    
    ONAME was not used in the export data, but it's the natural node op
    where we used OPACK instead. Renamed.
    
    Furthermore, OPACK and ONONAME nodes are replaced by the type checker
    with ONAME nodes, so OPACK nodes cannot occur when exporting type-checked
    code. Removed a special-case for OPACK nodes since they don't appear.
    
    Change-Id: I78b01a1badbf60e9283eaadeca2578a65d28cbd2
    Reviewed-on: https://go-review.googlesource.com/23053
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit aff4889089f970fb739acf5e3a5bddd3491a908b
Author: Robert Griesemer <gri@golang.org>
Date:   Wed May 11 11:28:36 2016 -0700

    cmd/compile: clean up encoding of method expressions and add test
    
    Fixes #15646.
    
    Change-Id: Ic13d1adc0a358149209195cdb03811eeee506fb8
    Reviewed-on: https://go-review.googlesource.com/23052
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e9407ae514df7d18e162ce03ebd530fe21aed16d
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Wed May 11 10:23:37 2016 -0700

    cmd/pprof: remove tempDir when no longer needed
    
    The pprof tools properly cleans up all files it creates, but forgets
    to clean up the temporary directory itself. This CL fixes that.
    
    Fixes #13863
    
    Change-Id: I1151c36cdad5ace7cc97e7e04001cf0149ef0f63
    Reviewed-on: https://go-review.googlesource.com/23019
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2ffb3e5d905b5622204d199128dec06cefd57790
Author: Marc-Antoine Ruel <maruel@chromium.org>
Date:   Thu Apr 7 14:24:24 2016 -0400

    os: fix Remove for file with read only attribute on Windows
    
    Include integration test. Confirmed that without the fix, the test case
    TestDeleteReadOnly fails.
    
    This permits to revert "cmd/go: reset read-only flag during TestIssue10952"
    This reverts commit 3b7841b3aff9204f054ffabbe4dd39d3e3dd3e91.
    
    Fixes #9606
    
    Change-Id: Ib55c151a8cf1a1da02ab18c34a9b58f615c34254
    Reviewed-on: https://go-review.googlesource.com/18235
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4d8031cf3c179d8682b62feae1d5a9109d14b382
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue May 10 15:09:23 2016 -0700

    net/http: make the MaxBytesReader.Read error sticky
    
    Fixes #14981
    
    Change-Id: I39b906d119ca96815801a0fbef2dbe524a3246ff
    Reviewed-on: https://go-review.googlesource.com/23009
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 20e362dae73b84e7b9dba9959444e5bc9d513ff1
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed May 11 09:37:46 2016 -0700

    cmd/cgo: remove //extern for check functions in gccgo Go prologue
    
    The //extern comments are incorrect and cause undefined symbol
    errorswhen building cgo code with -compiler=gccgo. The code is already
    designed to use weak references, and that support relies on the cgo
    check functions being treated as local functions.
    
    Change-Id: Ib38a640cc4ce6eba74cfbf41ba7147ec88769ec0
    Reviewed-on: https://go-review.googlesource.com/23014
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9a57fa31ff35024b9f628e7eae39bfd35bf90d77
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue May 10 16:09:16 2016 -0700

    net/http: document ResponseWriter read-vs-write concurrency rules
    
    Summary: Go's HTTP/1.x server closes the request body once writes are
    flushed. Go's HTTP/2 server supports concurrent read & write.
    
    Added a TODO to make the HTTP/1.x server also support concurrent
    read+write. But for now, document it.
    
    Updates #15527
    
    Change-Id: I81f7354923d37bfc1632629679c75c06a62bb584
    Reviewed-on: https://go-review.googlesource.com/23011
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 80423f1e64f1e939cddc455a29e5111527cd16f8
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Sat Dec 26 16:50:01 2015 +0900

    os/exec: cleanup and remove duplicated code
    
    Change-Id: Ia2f61427b1cc09064ac4c0563bccbd9b98767a0e
    Reviewed-on: https://go-review.googlesource.com/18118
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c1e88920606e78b06e936c9c249bd55f06dd8c51
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed May 11 13:04:22 2016 +0900

    net: fix nits found by vet
    
    Change-Id: I323231f31c4e1e7415661ebd943a90b2f1e9da1c
    Reviewed-on: https://go-review.googlesource.com/23020
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d1981ac313f6858cf1ec163dac94ea0d6904a731
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue May 10 04:29:32 2016 +0900

    net: reorganize interface tests to avoid vague flakiness
    
    This change reorganizes test cases for surveying network interfaces and
    address prefixes to make sure which part of the functionality is broken.
    
    Updates #7849.
    
    Change-Id: If6918075802eef69a7f1ee040010b3c46f4f4b97
    Reviewed-on: https://go-review.googlesource.com/22990
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b4538d7aaa1a600dc1d3724f9aecb5c8039e1324
Author: David du Colombier <0intro@gmail.com>
Date:   Tue May 10 07:43:17 2016 +0000

    Revert "os: enable TestGetppid on Plan 9"
    
    This reverts commit a677724edfc465193d2f79ee48d2c06defbc916b.
    
    Change-Id: I6a54ac26a6deca5b2a39ec9f899469a88b543d3d
    Reviewed-on: https://go-review.googlesource.com/22980
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0efc16284bc4fd5b8b31d3f6b6763f98700c5664
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Wed May 11 15:19:32 2016 +1000

    syscall: remove mksyscall_windows.go -xsys flag
    
    Also run "go generate" in
    internal/syscall/windows and internal/syscall/windows/registry
    
    Updates #15167
    
    Change-Id: I0109226962f81857fe11d308b869d561ea8ed9f9
    Reviewed-on: https://go-review.googlesource.com/23021
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d958dab095febfe542c6209b023d15f1d0de7128
Author: Robert Griesemer <gri@golang.org>
Date:   Tue May 10 17:06:19 2016 -0700

    go/importer: use correct path when checking if package was already imported
    
    The importer uses a global (shared) package map across multiple imports
    to determine if a package was imported before. That package map is usually
    indexed by package (import) path ('id' in this code). However, the binary
    importer was using the incoming (possibly unclean) path.
    
    Fixes #15517.
    
    Change-Id: I0c32a708dfccf345e0353fbda20ad882121e437c
    Reviewed-on: https://go-review.googlesource.com/23012
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b6712946c1b46eb629fb010e65e5b3735f94d171
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu Apr 7 15:12:32 2016 +1000

    runtime: make mksyscall_windows.go flags do what they say they do
    
    The -systemdll and -xsys flags generate broken code in some situations
    (see issue for details). Fix all that.
    
    This CL only fixes bugs in existing code, but I have more changes comming:
    
    golang.org/x/sys/windows is not the only package that uses mksyscall_windows.go.
    golang.org/x/exp/shiny and github.com/derekparker/delve do too. I also have
    few personal packages that use mksyscall_windows.go. None of those packages
    are aware of new -xsys flag. I would like to change mksyscall_windows.go, so
    external packages do not need to use -xsys flag. I would love to get rid of
    -xsys flag altogether, but I don't see how it is possible. So I will, probably,
    replace -xsys with a flag that means opposite to -xsys, and use new flag
    everywhere in standard libraries. Flag name suggestions are welcome.
    
    -systemdll flag makes users code more "secure". I would like to make -systemdll
    behaviour a default for all mksyscall_windows.go users. We use that already in
    standard library. If we think "secure" is important, we should encourage it in
    all users code. If mksyscall_windows.go user insist on using old code, provide
    -use_old_loaddll (need good name here) flag for that. So -systemdll flag will
    be replaced with -use_old_loaddll.
    
    Fixes #15167
    
    Change-Id: I516369507867358ba1b66aabe00a17a7b477016e
    Reviewed-on: https://go-review.googlesource.com/21645
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9628e6fd1d1afeedce7c4b45454e0bc5cbd0d5ff
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue May 10 15:32:00 2016 -0700

    runtime/testdata/testprogcgo: fix Windows C compiler warning
    
    Noticed and fix by Alex Brainman.
    
    Tested in https://golang.org/cl/23005 (which makes all compiler
    warnings fatal during development)
    
    Fixes #15623
    
    Change-Id: Ic19999fce8bb8640d963965cc328574efadd7855
    Reviewed-on: https://go-review.googlesource.com/23010
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 9780bf2a9587b6aa0c92526cc1d6d6d1ed4c7210
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue May 10 14:27:32 2016 -0700

    os/user: don't create C function mygetgrouplist
    
    Instead of exporting the C function mygetgrouplist as a global symbol to
    conflict with other symbols of the same name, use trivial Go code and a
    static C function.
    
    Change-Id: I98dd667814d0a0ed8f7b1d4cfc6483d5a6965b26
    Reviewed-on: https://go-review.googlesource.com/23008
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 81a89606ef7e1334a0a23dab2eaa295b381caebc
Author: Andrew Gerrand <adg@golang.org>
Date:   Tue May 10 11:48:48 2016 -0700

    doc: remove mention of %HOME% from installation instructions
    
    Fixes #15598
    
    Change-Id: I4cfb8799dab0e9e34cae2e61839911fd65e4cfa3
    Reviewed-on: https://go-review.googlesource.com/23004
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f77f22b2bf43f565ac0933c8e1068c387e4007c3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue May 10 09:46:39 2016 -0700

    net/http: update bundled x/net/http2
    
    Updates x/net/http2 to git rev 96dbb961 for golang.org/cl/23002
    
    Fixes #15366
    Updates #15134 (server part remains)
    
    Change-Id: I29336e624706f906b754da66381a620ae3293c6c
    Reviewed-on: https://go-review.googlesource.com/23003
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 78ff74375930d5ae391beae562c91da40e5d92a4
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue May 10 09:24:57 2016 -0700

    crypto/sha1: disable crashing AVX2 optimizations for now
    
    Updates #15617
    
    Change-Id: I2104776f8e789d987b4f2f7f08f2ebe979b747a1
    Reviewed-on: https://go-review.googlesource.com/23001
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 9e96ad851d9fb3feb7ee3f7b72213c5b7a9977aa
Author: Keith Randall <khr@golang.org>
Date:   Tue May 10 09:10:43 2016 -0700

    test: add test for unlowered ITab
    
    See #15604.  This was a bug in a CL that has since been
    rolled back.  Adding a test to challenge the next attempter.
    
    Change-Id: Ic43be254ea6eaab0071018cdc61d9b1c21f19cbf
    Reviewed-on: https://go-review.googlesource.com/23000
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 636670b8db9b72c80ec89da0a666be0d686269fd
Author: Scott Bell <scott@sctsm.com>
Date:   Sun May 8 18:17:59 2016 -0700

    net: use contexts for cgo-based DNS resolution
    
    Although calls to getaddrinfo can't be portably interrupted,
    we still benefit from more granular resource management by
    pushing the context downwards.
    
    Fixes #15321
    
    Change-Id: I5506195fc6493080410e3d46aaa3fe02018a24fe
    Reviewed-on: https://go-review.googlesource.com/22961
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9edb27e76f297c034e9383ad2d1bf48b23e1a25b
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Tue May 10 07:06:47 2016 -0700

    reflect: make Field panic when out of bounds, as documented
    
    Fixes #15046.
    
    Change-Id: Iba7216297735be8e1ec550ce5336d17dcd3fd6b7
    Reviewed-on: https://go-review.googlesource.com/22992
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 42b647bde669ffa1d6f056eac20a683b9d02a786
Author: David du Colombier <0intro@gmail.com>
Date:   Tue May 10 07:50:09 2016 +0200

    go/internal/gccgoimporter: remove workaround on Plan 9
    
    We fixed the implementation of the pread syscall in
    the Plan 9 kernel, so calling pread doesn't update the
    channel offset when reading a file.
    
    Fixes #11194.
    
    Change-Id: Ie4019e445542a73479728af861a50bb54caea3f6
    Reviewed-on: https://go-review.googlesource.com/22245
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a677724edfc465193d2f79ee48d2c06defbc916b
Author: David du Colombier <0intro@gmail.com>
Date:   Mon Apr 18 03:18:13 2016 +0200

    os: enable TestGetppid on Plan 9
    
    Fixes #8206.
    
    Change-Id: Iec1026ecc586495f5c9562cc84b3240c71d53da5
    Reviewed-on: https://go-review.googlesource.com/22164
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3d82432288885696b01357e921ec00116291a790
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Apr 19 20:59:03 2016 +0200

    os: add TestReadAtOffset
    
    In the Plan 9 kernel, there used to be a bug in the implementation of
    the pread syscall, where the channel offset was erroneously updated after
    calling pread on a file.
    
    This test verifies that ReadAt is behaving as expected.
    
    Fixes #14534.
    
    Change-Id: Ifc9fd40a1f94879ee7eb09b2ffc369aa2bec2926
    Reviewed-on: https://go-review.googlesource.com/22244
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f05c3aa24d815cd3869153750c9875e35fc48a6e
Author: Caleb Spare <cespare@gmail.com>
Date:   Wed Apr 13 16:51:25 2016 -0700

    encoding/json: support maps with integer keys
    
    This change makes encoding and decoding support integer types in map
    keys, converting to/from JSON string keys.
    
    JSON object keys are still sorted lexically, even though the keys may be
    integer strings.
    
    For backwards-compatibility, the existing Text(Un)Marshaler support for
    map keys (added in CL 20356) does not take precedence over the default
    encoding for string types. There is no such concern for integer types,
    so integer map key encoding is only used as a fallback if the map key
    type is not a Text(Un)Marshaler.
    
    Fixes #12529.
    
    Change-Id: I7e68c34f9cd19704b1d233a9862da15fabf0908a
    Reviewed-on: https://go-review.googlesource.com/22060
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9af83462c6f432b77a846a24b4d8efae9bdf0567
Author: Tilman Dilo <tilman.dilo@gmail.com>
Date:   Mon May 9 23:37:07 2016 +0200

    crypto/cipher: execute AES-GCM decryption example
    
    The decryption example for AES-GCM was not executed, hiding the fact
    that the provided ciphertext could not be authenticated.
    
    This commit adds the required output comment, replaces the ciphertext
    with a working example, and removes an unnecessary string conversion
    along the way.
    
    Change-Id: Ie6729ca76cf4a56c48b33fb3b39872105faa604b
    Reviewed-on: https://go-review.googlesource.com/22953
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d88261fb6581106e4e7d8d6c63f0e33c2a24361e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon May 9 17:21:11 2016 +0000

    time: don't depend on the io package
    
    The time package has never depended on the io package until
    a recent change during Go 1.7 to use the io.Seek* constants.
    
    The go/build dependency check didn't catch this because "time" was
    allowed to depend on meta package group "L0", which included "io".
    
    Adding the "io" package broke one of Dmitry's tools. The tool is
    fixable, but it's also not necessary for us to depend on "io" at all
    for some constants. Mirror the constants instead, and change
    deps_test.go to prevent an io dependency in the future.
    
    Change-Id: I74325228565279a74fa4a2f419643f5710e3e09f
    Reviewed-on: https://go-review.googlesource.com/22960
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 561c94884477f568bdb68aacebfeb4d0411a874b
Author: Hana Kim <hyangah@gmail.com>
Date:   Mon May 9 15:14:07 2016 -0400

    os: skip Lchown test on Android if symlink doesn't work
    
    After upgrading builder device (android/arm) to android 5.0.2,
    the test started failing. Running 'ln -s' from shell fails with
    permission error.
    
    Change-Id: I5b9e312806d58532b41ea3560ff079dabbc6424e
    Reviewed-on: https://go-review.googlesource.com/22962
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 256a9670cc9a0acd1fd70ad53ba7ab032d5b2933
Author: Austin Clements <austin@google.com>
Date:   Mon May 9 15:03:15 2016 -0400

    runtime: fix some out of date comments in bitmap code
    
    Change-Id: I4613aa6d62baba01686bbab10738a7de23daae30
    Reviewed-on: https://go-review.googlesource.com/22971
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 7ba54d45732219af86bde9a5b73c145db82b70c6
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Mon Mar 28 02:29:18 2016 -0700

    compress: update documentation regarding footer verification
    
    Address two documentation issues:
    1) Document that the GZIP and ZLIB footer is only verified when the
    reader has been fully consumed.
    2) The zlib reader is guaranteed to not read past the EOF if the
    input io.Reader is also a io.ByteReader. This functionality was
    documented in the flate and gzip packages but not on zlib.
    
    Fixes #14867
    
    Change-Id: I43d46b93e38f98a04901dc7d4f18ed2f9e09f6fb
    Reviewed-on: https://go-review.googlesource.com/21218
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5934523e75e2aa3725c4e709be56d9e84c472bfe
Author: Shenghou Ma <minux@golang.org>
Date:   Sun May 8 22:31:09 2016 -0400

    cmd/compile: document -l in godoc
    
    Fixes #15607.
    
    Change-Id: I3e68ad00ebe72027d064238d4e77f1ad6a52f533
    Reviewed-on: https://go-review.googlesource.com/22940
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3c090019172afd4517360606efc50750d3e278fa
Author: David Chase <drchase@google.com>
Date:   Fri May 6 17:05:02 2016 -0700

    cmd/compile: correct sparseSet probes in regalloc to avoid index error
    
    In regalloc, a sparse map is preallocated for later use by
    spill-in-loop sinking.  However, variables (spills) are added
    during register allocation before spill sinking, and a map
    query involving any of these new variables will index out of
    bounds in the map.
    
    To fix:
    1) fix the queries to use s.orig[v.ID].ID instead, to ensure
    proper indexing.  Note that s.orig will be nil for values
    that are not eligible for spilling (like memory and flags).
    
    2) add a test.
    
    Fixes #15585.
    
    Change-Id: I8f2caa93b132a0f2a9161d2178320d5550583075
    Reviewed-on: https://go-review.googlesource.com/22911
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 149ac34893ad1cc5023ae2fbabd0d553f21b0313
Author: Mikhail Gusarov <dottedmag@dottedmag.net>
Date:   Mon May 9 19:28:28 2016 +0200

    doc: update number of supported instruction sets
    
    Current  number was out-of-date since adding MIPS.
    
    Change-Id: I565342a92de3893b75cdfb76fa39f7fdf15672da
    Reviewed-on: https://go-review.googlesource.com/22952
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit feb6131b1a4c0da098821c516e06499add886182
Author: Russ Cox <rsc@golang.org>
Date:   Tue Apr 26 21:50:59 2016 -0400

    cmd/compile: add -linkobj flag to allow writing object file in two parts
    
    This flag is experimental and the semantics may change
    even after Go 1.7 is released. There are no changes to code
    not using the flag.
    
    The first part is for reading by future compiles.
    The second part is for reading by the final link step.
    Splitting the file this way allows distributed build systems
    to ship the compile-input part only to compile steps and
    the linker-input part only to linker steps.
    
    The first part is basically just the export data,
    and the second part is basically everything else.
    The overall files still have the same broad structure,
    so that existing tools will work with both halves.
    It's just that various pieces are empty in the two halves.
    
    This also copies the two bits of data the linker needed from
    export data into the object header proper, so that the linker
    doesn't need any export data at all. That eliminates a TODO
    that was left for switching to the binary export data.
    (Now the linker doesn't need to know about the switch.)
    
    The default is still to write out a combined output file.
    Nothing changes unless you pass -linkobj to the compiler.
    There is no support in the go command for -linkobj,
    since the go command doesn't copy objects around.
    The expectation is that other build systems (like bazel, say)
    might take advantage of this.
    
    The header adjustment and the option for the split output
    was intended as part of the zip archives, but the zip archives
    have been cut from Go 1.7. Doing this to the current archives
    both unblocks one step in the switch to binary export data
    and enables alternate build systems to experiment with the
    new flag using the Go 1.7 release.
    
    Change-Id: I8b6eab25b8a22b0a266ba0ac6d31e594f3d117f3
    Reviewed-on: https://go-review.googlesource.com/22500
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit aeecee8ce4cf1821dcb6b5e37e20f40696278498
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Mon May 9 15:11:24 2016 +0200

    runtime/race: deflake test
    
    The test sometimes fails on builders.
    The test uses sleeps to establish the necessary goroutine
    execution order. If sleeps undersleep/oversleep
    the race is still reported, but it can be reported when the
    main test goroutine returns. In such case test driver
    can't match the race with the test and reports failure.
    
    Wait for both test goroutines to ensure that the race
    is reported in the test scope.
    
    Fixes #15579
    
    Change-Id: I0b9bec0ebfb0c127d83eb5325a7fe19ef9545050
    Reviewed-on: https://go-review.googlesource.com/22951
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 87a2ae1fa25677dc9097a25292c54b7b9dac2c9d
Author: Robert Griesemer <gri@golang.org>
Date:   Sun May 8 20:59:53 2016 -0700

    cmd/compile: fix binary export of composite literals with implicit types
    
    Also:
    - replaced remaining panics with Fatal calls
    - more comments
    
    Fixes #15572.
    
    Change-Id: Ifb27e80b66700f5692a84078764a1e928d4b310d
    Reviewed-on: https://go-review.googlesource.com/22935
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 3696e469e5f8a5531c69ffcf091deaa692e81104
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun May 8 14:32:48 2016 -0700

    test: add test for issue 15602
    
    The problem was fixed by the rollback in CL 22930.
    This CL just adds a test to prevent regressions.
    
    Fixes #15602
    
    Change-Id: I37453f6e18ca43081266fe7f154c6d63fbaffd9b
    Reviewed-on: https://go-review.googlesource.com/22931
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 55546efeee9fb6104d3dfd76351e7765df0bdd71
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Sun May 8 12:54:31 2016 -0700

    Revert "cmd/compile: properly handle map assignments for OAS2DOTTYPE"
    
    This reverts commit 9d7c9b4384db01afd2acb27d3a4636b60e957f08.
    
    For #15602.
    
    Change-Id: I464184b05babe4cb8dedab6161efa730cea6ee2d
    Reviewed-on: https://go-review.googlesource.com/22930
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0b6e5e3d733c1da53244753b42940eddb7401c6c
Author: Joel Sing <joel@sing.id.au>
Date:   Sun May 8 01:27:45 2016 +1000

    cmd/link: specify correct size for dynamic symbols in 386 elf output
    
    Currently 386 ELF binaries are generated with dynamic symbols that have
    a size of zero bytes, even though the symbol in the symbol table has
    the correct size. Fix this by specifying the correct size when creating
    dynamic symbols.
    
    Issue found on OpenBSD -current, where ld.so is now producing link
    warnings due to mismatched symbol sizes.
    
    Fixes #15593.
    
    Change-Id: Ib1a12b23ff9159c61ac980bf48a983b86f3df256
    Reviewed-on: https://go-review.googlesource.com/22912
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 394ac818b037ab8a3714b8a23e06e17a1e05aace
Author: Robert Griesemer <gri@golang.org>
Date:   Thu May 5 18:03:59 2016 -0700

    cmd/compile: add and enable (internal) option to only track named types
    
    The new export format keeps track of all types that are exported.
    If a type is seen that was exported before, only a reference to
    that type is emitted. The importer maintains a list of all the
    seen types and uses that list to resolve type references.
    
    The existing compiler infrastructure's invariants assumes that
    only named types are referred to before they are fully set up.
    Referring to unnamed incomplete types causes problems. One of
    the issues was #15548.
    
    Added a new internal flag 'trackAllTypes' to enable/disable
    this type tracking. With this change only named types are
    tracked.
    
    Verified that this fix also addresses #15548, even w/o the
    prior fix for that issue (in fact that prior fix is turned
    off if trackAllTypes is disabled because it's not needed).
    
    The test for #15548 covers also this change.
    
    For #15548.
    
    Change-Id: Id0b3ff983629703d025a442823f99649fd728a56
    Reviewed-on: https://go-review.googlesource.com/22839
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit fa270ad98e77cd0625c97eb01ad01efe11a324e8
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sat May 7 07:24:39 2016 +0200

    cmd/go: add -shared to darwin/arm{,64} default build mode
    
    Buildmode c-archive now supports position independent code for
    darwin/arm (in addition to darwin/arm64). Make PIC (-shared) the
    default for both platforms in the default buildmode.
    
    Without this change, gomobile will go install the standard library
    into its separate package directory without PIC support.
    
    Also add -shared to darwin/arm64 in buildmode c-archive, for
    symmetry (darwin/arm64 always generates position independent code).
    
    Fixes #15519
    
    Change-Id: If27d2cbea8f40982e14df25da2703cbba572b5c6
    Reviewed-on: https://go-review.googlesource.com/22920
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9d7c9b4384db01afd2acb27d3a4636b60e957f08
Author: Tal Shprecher <tshprecher@gmail.com>
Date:   Thu May 5 15:14:08 2016 -0700

    cmd/compile: properly handle map assignments for OAS2DOTTYPE
    
    The boolean destination in an OAS2DOTTYPE expression craps out during
    compilation when trying to assign to a map entry because, unlike slice entries,
    map entries are not directly addressable in memory. The solution is to
    properly order the boolean destination node so that map entries are set
    via autotmp variables.
    
    Fixes #14678
    
    Change-Id: If344e8f232b5bdac1b53c0f0d21eeb43ab17d3de
    Reviewed-on: https://go-review.googlesource.com/22833
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e6ec82067a9068c93db6e7041017060a1c863452
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Mar 9 10:00:12 2016 +0100

    runtime: use entire address space on 32 bit
    
    In issue #13992, Russ mentioned that the heap bitmap footprint was
    halved but that the bitmap size calculation hadn't been updated. This
    presents the opportunity to either halve the bitmap size or double
    the addressable virtual space. This CL doubles the addressable virtual
    space. On 32 bit this can be tweaked further to allow the bitmap to
    cover the entire 4GB virtual address space, removing a failure mode
    if the kernel hands out memory with a too low address.
    
    First, fix the calculation and double _MaxArena32 to cover 4GB virtual
    memory space with the same bitmap size (256 MB).
    
    Then, allow the fallback mode for the initial memory reservation
    on 32 bit (or 64 bit with too little available virtual memory) to not
    include space for the arena. mheap.sysAlloc will automatically reserve
    additional space when the existing arena is full.
    
    Finally, set arena_start to 0 in 32 bit mode, so that any address is
    acceptable for subsequent (additional) reservations.
    
    Before, the bitmap was always located just before arena_start, so
    fix the two places relying on that assumption: Point the otherwise unused
    mheap.bitmap to one byte after the end of the bitmap, and use it for
    bitmap addressing instead of arena_start.
    
    With arena_start set to 0 on 32 bit, the cgoInRange check is no longer a
    sufficient check for Go pointers. Introduce and call inHeapOrStack to
    check whether a pointer is to the Go heap or stack.
    
    While we're here, remove sysReserveHigh which seems to be unused.
    
    Fixes #13992
    
    Change-Id: I592b513148a50b9d3967b5c5d94b86b3ec39acc2
    Reviewed-on: https://go-review.googlesource.com/20471
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 83676d694b64205e80c042ca7cf61f7ad4de6c62
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 6 18:33:39 2016 +0000

    net/url: remove RFC 3986 mention in package comment
    
    Change-Id: Ifd707a4bbfcb1721655b4fce2045f3b043e66818
    Reviewed-on: https://go-review.googlesource.com/22859
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 7c5c6645d2ac21073b146c3d1a83c9b8c6463c25
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 6 18:46:35 2016 +0000

    net: skip more flaky net tests on flaky net builders
    
    e.g. https://storage.googleapis.com/go-build-log/9b937dd8/linux-arm_df54a25a.log
    
    Change-Id: Ic5864c7bd840b4f0c6341f919fcbcd5c708b14e7
    Reviewed-on: https://go-review.googlesource.com/22881
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit a1813ae0a091e1b880c0d3472112d2c725c2fa18
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri May 6 11:32:18 2016 -0700

    misc/cgo/testcarchive: avoid possible pthread_create race
    
    The old code assumed that the thread ID set by pthread_create would be
    available in the newly created thread.  While that is clearly true
    eventually, it is not necessarily true immediately.  Rather than try to
    pass down the thread ID, just call pthread_self in the created thread.
    
    Fixes #15576 (I hope).
    
    Change-Id: Ic07086b00e4fd5676c04719a299c583320da64a1
    Reviewed-on: https://go-review.googlesource.com/22880
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f0e2d32fde77ad03616304ab42b8c7426cf3a350
Author: Russ Cox <rsc@golang.org>
Date:   Fri May 6 15:34:25 2016 +0000

    Revert "net/url: validate ports in IPv4 addresses"
    
    This reverts commit 9f1ccd647fcdb1b703c1042c90434e15aff75013.
    
    For #14860.
    
    Change-Id: I63522a4dda8915dc8b972ae2e12495553ed65f09
    Reviewed-on: https://go-review.googlesource.com/22861
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4eccc77f196edfa7646b0e92a11ef8d96ef85b57
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 6 18:21:22 2016 +0000

    net/http: wait longer for subprocess to startup in test
    
    Might deflake the occasional linux-amd64-race failures.
    
    Change-Id: I273b0e32bb92236168eb99887b166e079799c1f1
    Reviewed-on: https://go-review.googlesource.com/22858
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 31283dd4836542f94b063bfd2886fc32639358f7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 6 18:11:38 2016 +0000

    net/http: don't assume Response.Request is populated after redirect errors
    
    Fixes #15577
    
    Change-Id: I5f023790a393b17235db2e66c02c2483773ddc1a
    Reviewed-on: https://go-review.googlesource.com/22857
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 670a5cda2048af8d83958af0f4b2fda8f7b4ea72
Author: Russ Cox <rsc@golang.org>
Date:   Fri May 6 15:28:19 2016 +0000

    Revert "testing/quick: generate more map and slice states"
    
    This reverts commit 0ccabe2e0b42a2602e0f37ce28d5368aa811f530.
    
    Change-Id: Ib1c230fb6801c0ee26f4a352b0c1130fa240a76a
    Reviewed-on: https://go-review.googlesource.com/22860
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 131231b8db26b38c9c2fdc52fb788241f5c2de51
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 6 16:07:11 2016 +0000

    os: rename remaining four os1_*.go files to os_*.go
    
    Change-Id: Ice9c234960adc7857c8370b777a0b18e29d59281
    Reviewed-on: https://go-review.googlesource.com/22853
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1ff57143af65014c80e39cc0f19cd97a455f5b49
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 6 16:24:57 2016 +0000

    net: ignore network failures on some builders
    
    We run the external network tests on builders, but some of our
    builders have less-than-ideal DNS connectivity. This change continues
    to run the tests on all builders, but marks certain builders as flaky
    (network-wise), and only validates their DNS results if they got DNS
    results.
    
    Change-Id: I826dc2a6f6da55add89ae9c6db892b3b2f7b526b
    Reviewed-on: https://go-review.googlesource.com/22852
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 61602b0e9e1daa0490793ef9ada3a51f8f482265
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 6 16:06:02 2016 +0000

    runtime: delete empty files
    
    I meant to delete these in CL 22850, actually.
    
    Change-Id: I0c286efd2b9f1caf0221aa88e3bcc03649c89517
    Reviewed-on: https://go-review.googlesource.com/22851
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a5c5f6ea94dcd9caad0f0df8caaf68f8659900b2
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Fri May 6 18:16:52 2016 +0200

    all: fix copy-and-paste errors in tests
    
    Fixes #15570
    
    Change-Id: I95d1ac26e342c3bbf36ad1f0209711ea96eaf487
    Reviewed-on: https://go-review.googlesource.com/22870
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c81a3532fea42df33dea54497dfaa96873c2d976
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Fri Apr 15 00:33:28 2016 +0300

    cmd/vet: check sync.* types' copying
    
    Embed noLock struct into the following types, so `go vet -copylocks` catches
    their copying additionally to types containing sync.Mutex:
      - sync.Cond
      - sync.WaitGroup
      - sync.Pool
      - atomic.Value
    
    Fixes #14582
    
    Change-Id: Icb543ef5ad10524ad239a15eec8a9b334b0e0660
    Reviewed-on: https://go-review.googlesource.com/22015
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 258a4c3daf992958f5d7dc5bccf2c5b41e236959
Author: Richard Miller <miller.research@gmail.com>
Date:   Fri May 6 14:21:52 2016 +0100

    syscall,os,net: don't use ForkLock in plan9
    
    This is the follow-on to CL 22610: now that it's the child instead of
    the parent which lists unwanted fds to close in syscall.StartProcess,
    plan9 no longer needs the ForkLock to protect the list from changing.
    The readdupdevice function is also now unused and can be removed.
    
    Change-Id: I904c8bbf5dbaa7022b0f1a1de0862cd3064ca8c7
    Reviewed-on: https://go-review.googlesource.com/22842
    Reviewed-by: David du Colombier <0intro@gmail.com>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2dc680007e35f4cb87527582eb73a653392aa8c3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri May 6 08:26:37 2016 -0700

    runtime: merge the last four os-vs-os1 files together
    
    Change-Id: Ib0ba691c4657fe18a4659753e70d97c623cb9c1d
    Reviewed-on: https://go-review.googlesource.com/22850
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit bf151cc2aa4094b4633a7e5f07a34227d58231fe
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu May 5 16:51:54 2016 -0700

    cmd/compile/internal/mips64: fix large uint -> float conversion
    
    Re-enable TestFP in cmd/compile/internal/gc on mips64.
    
    Fixes #15552.
    
    Change-Id: I5c3a5564b94d28c723358f0862468fb6da371991
    Reviewed-on: https://go-review.googlesource.com/22835
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fafadc521ede90f8abed73e8d209e130c456e983
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Fri Apr 29 16:17:14 2016 +0300

    crypto/sha1: Add AVX2 version for AMD64
    
    name             old time/op    new time/op    delta
    Hash8Bytes-48       271ns ± 8%     273ns ± 5%     ~     (p=0.313 n=19+19)
    Hash320Bytes-48    1.04µs ± 7%    0.75µs ± 8%  -27.66%  (p=0.000 n=20+20)
    Hash1K-48          2.72µs ± 6%    1.75µs ± 6%  -35.79%  (p=0.000 n=19+20)
    Hash8K-48          19.9µs ± 7%    11.6µs ± 6%  -41.84%  (p=0.000 n=20+19)
    
    name             old speed      new speed      delta
    Hash8Bytes-48    29.5MB/s ± 8%  29.3MB/s ± 5%     ~     (p=0.314 n=19+19)
    Hash320Bytes-48   307MB/s ± 7%   424MB/s ± 8%  +38.29%  (p=0.000 n=20+20)
    Hash1K-48         377MB/s ± 6%   587MB/s ± 6%  +55.76%  (p=0.000 n=19+20)
    Hash8K-48         413MB/s ± 7%   709MB/s ± 6%  +71.85%  (p=0.000 n=20+19)
    
    Change-Id: I2963cf744eeb2e8191d4e4223fbf6f533a7fd405
    Reviewed-on: https://go-review.googlesource.com/22607
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2210d88a889e0ea463bcdef2b658aaec1050cf01
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Fri Apr 29 16:19:17 2016 +0300

    crypto/sha256: Use AVX2 if possible
    
    name          old time/op    new time/op    delta
    Hash8Bytes-4     376ns ± 0%     246ns ± 0%  -34.57%  (p=0.000 n=20+20)
    Hash1K-4        5.21µs ± 0%    2.82µs ± 0%  -45.83%  (p=0.000 n=20+20)
    Hash8K-4        38.6µs ± 0%    20.8µs ± 0%  -46.05%  (p=0.000 n=20+20)
    
    name          old speed      new speed      delta
    Hash8Bytes-4  21.2MB/s ± 0%  32.4MB/s ± 0%  +52.70%  (p=0.000 n=15+19)
    Hash1K-4       197MB/s ± 0%   363MB/s ± 0%  +84.60%  (p=0.000 n=20+20)
    Hash8K-4       212MB/s ± 0%   393MB/s ± 0%  +85.36%  (p=0.000 n=20+20)
    
    Change-Id: Ib50119c591074ff486d11d3566e24b691bcc6598
    Reviewed-on: https://go-review.googlesource.com/22608
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit e558fdbd9a3865dd383f50524bb8a18b9e6a0cb0
Author: Elias Naur <elias.naur@gmail.com>
Date:   Fri May 6 15:09:57 2016 +0200

    misc/cgo/testcarchive: don't force -no_pie on Darwin
    
    Now that darwin/arm supports position independent code, allow the
    binaries generated by the c-archive tests be position independent
    (PIE) as well.
    
    Change-Id: If0517f06e92349ada29a4e3e0a951f08b0fcc710
    Reviewed-on: https://go-review.googlesource.com/22841
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 009c002c925e391e5a7a406c9175aefafb6c9e3c
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Fri Apr 29 16:14:57 2016 +0300

    cmd/internal/obj/x86: add AVX2 instrutions needed for sha1/sha512/sha256 acceleration
    
    This means: VPSHUFB, VPSHUFD, VPERM2F128, VPALIGNR, VPADDQ, VPADDD, VPSRLDQ,
    VPSLLDQ, VPSRLQ, VPSLLQ, VPSRLD, VPSLLD, VPOR, VPBLENDD, VINSERTI128,
    VPERM2I128, RORXL, RORXQ.
    
    Change-Id: Ief27190ee6acfa86b109262af5d999bc101e923d
    Reviewed-on: https://go-review.googlesource.com/22606
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 2e32efc44ac86cce3bd0808e6049d8c9b0225ba8
Author: Shenghou Ma <minux@golang.org>
Date:   Fri May 6 00:53:42 2016 -0400

    runtime: get randomness from AT_RANDOM AUXV on linux/mips64x
    
    Fixes #15148.
    
    Change-Id: If3b628f30521adeec1625689dbc98aaf4a9ec858
    Reviewed-on: https://go-review.googlesource.com/22811
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ef92857e27556804d66e72e2360dc2c6b6554bd7
Author: Cherry Zhang <cherryyz@google.com>
Date:   Thu May 5 09:10:49 2016 -0700

    cmd/go, cmd/cgo: pass "-mabi=64" to gcc on mips64
    
    Change-Id: I9ac2ae57a00cee23d6255db02419b0a0f087d4f3
    Reviewed-on: https://go-review.googlesource.com/22801
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>

commit d68f800620b4295039912066970fb2be914f1d1e
Author: Robert Griesemer <gri@golang.org>
Date:   Thu May 5 17:46:58 2016 -0700

    test: update test for issue 15548
    
    Accidentally checked in the version of file c.go that doesn't
    exhibit the bug - hence the test was not testing the bug fix.
    Double-checked that this version exposes the bug w/o the fix.
    
    Change-Id: Ie4dc455229d1ac802a80164b5d549c2ad4d971f5
    Reviewed-on: https://go-review.googlesource.com/22837
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit b90cb3f4716d3fede57bf8e798d27406fba5c294
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu May 5 17:52:37 2016 -0700

    cmd/go: fail with nice error message on bad GOOS/GOARCH pair
    
    Fixes #12272
    
    Change-Id: I2115ec62ed4061084c482eb385a583a1c1909888
    Reviewed-on: https://go-review.googlesource.com/22838
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 30bfafc319288e8cfe54111664e3f2f259998a0a
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Thu Sep 17 00:22:56 2015 -0700

    archive/tar: centralize all information about tar header format
    
    The Reader and Writer have hard-coded constants regarding the
    offsets and lengths of certain fields in the tar format sprinkled
    all over. This makes it harder to verify that the offsets are
    correct since a reviewer would need to search for them throughout
    the code. Instead, all information about the layout of header
    fields should be centralized in one single file. This has the
    advantage of being both centralized, and also acting as a form
    of documentation about the header struct format.
    
    This method was chosen over using "encoding/binary" since that
    method would cause an allocation of a header struct every time
    binary.Read was called. This method causes zero allocations and
    its logic is no longer than if structs were declared.
    
    Updates #12594
    
    Change-Id: Ic7a0565d2a2cd95d955547ace3b6dea2b57fab34
    Reviewed-on: https://go-review.googlesource.com/14669
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 88d3db0a5b5196ed45c96014d5a2d32e4e41e34e
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 6 12:45:23 2016 -0500

    runtime: stop traceback at foreign function
    
    This can only happen when profiling and there is foreign code
    at the top of the g0 stack but we're not in cgo.
    That in turn only happens with the race detector.
    
    Fixes #13568.
    
    Change-Id: I23775132c9c1a3a3aaae191b318539f368adf25e
    Reviewed-on: https://go-review.googlesource.com/18322
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a633d766d1763f4f4648e423c4e8a8635b183d03
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Fri May 6 08:51:16 2016 +0900

    Revert "net: add support for Zone of IPNet"
    
    Updates #14518.
    
    This reverts commit 3e9264c9ae781a2cd28127deaed6ae26f84b4b15.
    
    Change-Id: I2531b04efc735b5b51ef675541172f2f5ae747d9
    Reviewed-on: https://go-review.googlesource.com/22836
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit acc757f678a42ba1ffbf8bb9886de4fe080302de
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Apr 5 11:22:53 2016 -0700

    all: use SeekStart, SeekCurrent, SeekEnd
    
    CL/19862 (f79b50b8d5bc159561c1dcf7c17e2a0db96a9a11) recently introduced the constants
    SeekStart, SeekCurrent, and SeekEnd to the io package. We should use these constants
    consistently throughout the code base.
    
    Updates #15269
    
    Change-Id: If7fcaca7676e4a51f588528f5ced28220d9639a2
    Reviewed-on: https://go-review.googlesource.com/22097
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Joe Tsai <joetsai@digital-static.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6db98a3c51549eb5e1e481e9bca6ede7e8e15f54
Author: David Chase <drchase@google.com>
Date:   Thu May 5 13:35:10 2016 -0700

    cmd/compile: repair MININT conversion bug in arm softfloat
    
    Negative-case conversion code was wrong for minimum int32,
    used negate-then-widen instead of widen-then-negate.
    
    Test already exists; this fixes the failure.
    
    Fixes #15563.
    
    Change-Id: I4b0b3ae8f2c9714bdcc405d4d0b1502ccfba2b40
    Reviewed-on: https://go-review.googlesource.com/22830
    Run-TryBot: David Chase <drchase@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5f83bf6053763801beb84a926cde7221874bc4f7
Author: Alan Donovan <adonovan@google.com>
Date:   Thu May 5 14:56:58 2016 -0700

    go/token: document postcondition of SetLines
    
    Change-Id: Ie163deade396b3e298a93845b9ca4d52333ea82a
    Reviewed-on: https://go-review.googlesource.com/22831
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 5bf9b39acc7c4e9658190d8606b5d719678db14a
Author: Richard Miller <miller.research@gmail.com>
Date:   Fri Apr 29 21:02:59 2016 +0100

    os/exec: re-enable TestExtraFiles for plan9
    
    This test should now succeed after CL 22610 which fixes issue #7118
    
    Change-Id: Ie785a84d77b27c832a1ddd81699bf25dab24b97d
    Reviewed-on: https://go-review.googlesource.com/22640
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: David du Colombier <0intro@gmail.com>
    Run-TryBot: David du Colombier <0intro@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 639a20da908c2646de46b93ca2933651363ec22a
Author: Richard Miller <miller.research@gmail.com>
Date:   Fri Apr 29 17:39:33 2016 +0100

    syscall: simplify closing of extra fds in plan9 StartProcess
    
    Reviving earlier work by @ality in https://golang.org/cl/57890043
    to make the closing of extra file descriptors in syscall.StartProcess
    less race-prone. Instead of making a list of open fds in the parent
    before forking, the child can read through the list of open fds and
    close the ones not explicitly requested.  Also eliminate the
    complication of keeping open any extra fds which were inherited by
    the parent when it started.
    
    This CL will be followed by one to eliminate the ForkLock in plan9,
    which is now redundant.
    
    Fixes #5605
    
    Change-Id: I6b4b942001baa54248b656c52dced3b62021c486
    Reviewed-on: https://go-review.googlesource.com/22610
    Run-TryBot: David du Colombier <0intro@gmail.com>
    Reviewed-by: David du Colombier <0intro@gmail.com>

commit 9b05ae612a4496df317e3c1c770b4b9c5648616d
Author: Andrew Gerrand <adg@golang.org>
Date:   Thu May 5 13:19:07 2016 -0700

    doc: update broken links in release notes
    
    Fixes #15559
    
    Change-Id: Ie58650f35e32c1f49669134b62876357abcdc583
    Reviewed-on: https://go-review.googlesource.com/22823
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8650c2303449a1cccf3a05609a63193e34d6bae8
Author: Robert Griesemer <gri@golang.org>
Date:   Thu May 5 09:39:50 2016 -0700

    cmd/compile: verify imported types after they are fully imported
    
    Fixes #15548.
    
    Change-Id: I1dfa9c8739a4b6d5e4c737c1a1e09e80e045b7aa
    Reviewed-on: https://go-review.googlesource.com/22803
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 1a7fc2357b1c26dcdf4fa57dee67a1172696801f
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Wed May 4 01:42:13 2016 -0600

    runtime: print signal name in panic, if name is known
    
    Adds a small function signame that infers a signal name
    from the signal table, otherwise will fallback to using
    hex(sig) as previously. No signal table is present for
    Windows hence it will always print the hex value.
    
    Sample code and new result:
    ```go
    package main
    
    import (
      "fmt"
      "time"
    )
    
    func main() {
      defer func() {
        if err := recover(); err != nil {
          fmt.Printf("err=%v\n", err)
        }
      }()
    
      ticker := time.Tick(1e9)
      for {
        <-ticker
      }
    }
    ```
    
    ```shell
    $ go run main.go &
    $ kill -11 <pid>
    fatal error: unexpected signal during runtime execution
    [signal SIGSEGV: segmentation violation code=0x1 addr=0xb01dfacedebac1e
    pc=0xc71db]
    ...
    ```
    
    Fixes #13969
    
    Change-Id: Ie6be312eb766661f1cea9afec352b73270f27f9d
    Reviewed-on: https://go-review.googlesource.com/22753
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fafd792de30f46cbd822fd6bb041c60c7b5fbe6d
Author: Russ Cox <rsc@golang.org>
Date:   Wed Jan 6 14:37:45 2016 -0500

    net: fix hostLookupOrder("")
    
    Fixes #13623.
    
    Change-Id: I1bd96aa7b6b715e4dbdcf0c37c2d29228df6565c
    Reviewed-on: https://go-review.googlesource.com/18329
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit eeca3ba92fdb07e44abf3e2bebfcede03e1eae12
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Thu Apr 28 07:16:08 2016 -0500

    sync/atomic, runtime/internal/atomic: improve ppc64x atomics
    
    The following performance improvements have been made to the
    low-level atomic functions for ppc64le & ppc64:
    
    - For those cases containing a lwarx and stwcx (or other sizes):
    sync, lwarx, maybe something, stwcx, loop to sync, sync, isync
    The sync is moved before (outside) the lwarx/stwcx loop, and the
     sync after is removed, so it becomes:
    sync, lwarx, maybe something, stwcx, loop to lwarx, isync
    
    - For the Or8 and And8, the shifting and manipulation of the
    address to the word aligned version were removed and the
    instructions were changed to use lbarx, stbcx instead of
    register shifting, xor, then lwarx, stwcx.
    
    - New instructions LWSYNC, LBAR, STBCC were tested and added.
    runtime/atomic_ppc64x.s was changed to use the LWSYNC opcode
    instead of the WORD encoding.
    
    Fixes #15469
    
    Ran some of the benchmarks in the runtime and sync directories.
    Some results varied from run to run but the trend was improvement
    based on best times for base and new:
    
    runtime.test:
    BenchmarkChanNonblocking-128         0.88          0.89          +1.14%
    BenchmarkChanUncontended-128         569           511           -10.19%
    BenchmarkChanContended-128           63110         53231         -15.65%
    BenchmarkChanSync-128                691           598           -13.46%
    BenchmarkChanSyncWork-128            11355         11649         +2.59%
    BenchmarkChanProdCons0-128           2402          2090          -12.99%
    BenchmarkChanProdCons10-128          1348          1363          +1.11%
    BenchmarkChanProdCons100-128         1002          746           -25.55%
    BenchmarkChanProdConsWork0-128       2554          2720          +6.50%
    BenchmarkChanProdConsWork10-128      1909          1804          -5.50%
    BenchmarkChanProdConsWork100-128     1624          1580          -2.71%
    BenchmarkChanCreation-128            237           212           -10.55%
    BenchmarkChanSem-128                 705           667           -5.39%
    BenchmarkChanPopular-128             5081190       4497566       -11.49%
    
    BenchmarkCreateGoroutines-128             532           473           -11.09%
    BenchmarkCreateGoroutinesParallel-128     35.0          34.7          -0.86%
    BenchmarkCreateGoroutinesCapture-128      4923          4200          -14.69%
    
    sync.test:
    BenchmarkUncontendedSemaphore-128      112           94.2          -15.89%
    BenchmarkContendedSemaphore-128        133           128           -3.76%
    BenchmarkMutexUncontended-128          1.90          1.67          -12.11%
    BenchmarkMutex-128                     353           310           -12.18%
    BenchmarkMutexSlack-128                304           283           -6.91%
    BenchmarkMutexWork-128                 554           541           -2.35%
    BenchmarkMutexWorkSlack-128            567           556           -1.94%
    BenchmarkMutexNoSpin-128               275           242           -12.00%
    BenchmarkMutexSpin-128                 1129          1030          -8.77%
    BenchmarkOnce-128                      1.08          0.96          -11.11%
    BenchmarkPool-128                      29.8          27.4          -8.05%
    BenchmarkPoolOverflow-128              40564         36583         -9.81%
    BenchmarkSemaUncontended-128           3.14          2.63          -16.24%
    BenchmarkSemaSyntNonblock-128          1087          1069          -1.66%
    BenchmarkSemaSyntBlock-128             897           893           -0.45%
    BenchmarkSemaWorkNonblock-128          1034          1028          -0.58%
    BenchmarkSemaWorkBlock-128             949           886           -6.64%
    
    Change-Id: I4403fb29d3cd5254b7b1ce87a216bd11b391079e
    Reviewed-on: https://go-review.googlesource.com/22549
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 0960c7c7eb30b3d651037c799aaa0d80722f063f
Author: Shenghou Ma <minux@golang.org>
Date:   Thu May 5 14:22:34 2016 -0400

    context: use https in docs
    
    Change-Id: I9354712768702e3b083c77f30165a34cb414d686
    Reviewed-on: https://go-review.googlesource.com/22810
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit bc1989f1154b5b8f235e7e4932d935e490d6e79e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed May 4 14:04:03 2016 -0700

    cmd/compile: optimize lookupVarOutgoing
    
    If b has exactly one predecessor, as happens
    frequently with static calls, we can make
    lookupVarOutgoing generate less garbage.
    
    Instead of generating a value that is just
    going to be an OpCopy and then get eliminated,
    loop. This can lead to lots of looping.
    However, this loop is way cheaper than generating
    lots of ssa.Values and then eliminating them.
    
    For a subset of the code in #15537:
    
    Before:
    
           28.31 real        36.17 user         1.68 sys
    2282450944  maximum resident set size
    
    After:
    
            9.63 real        11.66 user         0.51 sys
     638144512  maximum resident set size
    
    Updates #15537.
    
    Excitingly, it appears that this also helps
    regular code:
    
    name       old time/op      new time/op      delta
    Template        288ms ± 6%       276ms ± 7%   -4.13%        (p=0.000 n=21+24)
    Unicode         143ms ± 8%       141ms ±10%     ~           (p=0.287 n=24+25)
    GoTypes         932ms ± 4%       874ms ± 4%   -6.20%        (p=0.000 n=23+22)
    Compiler        4.89s ± 4%       4.58s ± 4%   -6.46%        (p=0.000 n=22+23)
    MakeBash        40.2s ±13%       39.8s ± 9%     ~           (p=0.648 n=23+23)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   388user-ms ±10%  373user-ms ± 5%   -3.80%        (p=0.000 n=24+25)
    Unicode    203user-ms ± 6%  202user-ms ± 7%     ~           (p=0.492 n=22+24)
    GoTypes    1.29user-s ± 4%  1.17user-s ± 4%   -9.67%        (p=0.000 n=25+23)
    Compiler   6.86user-s ± 5%  6.28user-s ± 4%   -8.49%        (p=0.000 n=25+25)
    
    name       old alloc/op     new alloc/op     delta
    Template       51.5MB ± 0%      47.6MB ± 0%   -7.47%        (p=0.000 n=22+25)
    Unicode        37.2MB ± 0%      37.1MB ± 0%   -0.21%        (p=0.000 n=25+25)
    GoTypes         166MB ± 0%       138MB ± 0%  -16.83%        (p=0.000 n=25+25)
    Compiler        756MB ± 0%       628MB ± 0%  -16.96%        (p=0.000 n=25+23)
    
    name       old allocs/op    new allocs/op    delta
    Template         450k ± 0%        445k ± 0%   -1.02%        (p=0.000 n=25+25)
    Unicode          356k ± 0%        356k ± 0%     ~           (p=0.374 n=24+25)
    GoTypes         1.31M ± 0%       1.25M ± 0%   -4.18%        (p=0.000 n=25+25)
    Compiler        5.29M ± 0%       5.02M ± 0%   -5.15%        (p=0.000 n=25+23)
    
    It also seems to help in other cases in which
    phi insertion is a pain point (#14774, #14934).
    
    Change-Id: Ibd05ed7b99d262117ece7bb250dfa8c3d1cc5dd2
    Reviewed-on: https://go-review.googlesource.com/22790
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>

commit 6ccab441de1063a16b263107426fe62d2cefa990
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed May 4 23:47:37 2016 -0700

    cmd/compile/internal/gc: skip TestFP on mips64x
    
    The legacy mips64 backend doesn't handle large uint->float conversion
    correctly. See #15552.
    
    Change-Id: I84ceeaa95cc4e85f09cc46dfb30ab5d151f6b205
    Reviewed-on: https://go-review.googlesource.com/22800
    Reviewed-by: Minux Ma <minux@golang.org>

commit bfa89c3cd42f4301c5dc8657e02372bdd449cfcb
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed May 4 23:42:10 2016 -0700

    cmd/compile/internal/gc: remove duplicated TestFP
    
    TestFp and TestFP are same, remove one.
    
    Change-Id: Iffdece634cd4572421974496298925e7c6ac13a9
    Reviewed-on: https://go-review.googlesource.com/22799
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4fa050024fbdfff9eb43c930c5b1f73b1a16fafa
Author: Keith Randall <khr@golang.org>
Date:   Thu Apr 28 16:52:47 2016 -0700

    cmd/compile: enable constant-time CFG editing
    
    Provide indexes along with block pointers for Preds
    and Succs arrays.  This allows us to splice edges in
    and out of those arrays in constant time.
    
    Fixes worst-case O(n^2) behavior in deadcode and fuse.
    
    benchmark                     old ns/op      new ns/op     delta
    BenchmarkFuse1-8              2065           2057          -0.39%
    BenchmarkFuse10-8             9408           9073          -3.56%
    BenchmarkFuse100-8            105238         76277         -27.52%
    BenchmarkFuse1000-8           3982562        1026750       -74.22%
    BenchmarkFuse10000-8          301220329      12824005      -95.74%
    BenchmarkDeadCode1-8          1588           1566          -1.39%
    BenchmarkDeadCode10-8         4333           4250          -1.92%
    BenchmarkDeadCode100-8        32031          32574         +1.70%
    BenchmarkDeadCode1000-8       590407         468275        -20.69%
    BenchmarkDeadCode10000-8      17822890       5000818       -71.94%
    BenchmarkDeadCode100000-8     1388706640     78021127      -94.38%
    BenchmarkDeadCode200000-8     5372518479     168598762     -96.86%
    
    Change-Id: Iccabdbb9343fd1c921ba07bbf673330a1c36ee17
    Reviewed-on: https://go-review.googlesource.com/22589
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
```
