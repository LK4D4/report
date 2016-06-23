# June 23, 2016 Report

Number of commits: 27

## Compilation time

* github.com/coreos/etcd/cmd: from 14.545252948s to 14.42240503s, -0.84%
* github.com/boltdb/bolt/cmd/bolt: from 553.946518ms to 567.132211ms, +2.38%
* github.com/gogits/gogs: from 12.97118458s to 12.946595833s, -0.19%
* github.com/spf13/hugo: from 6.683664056s to 6.743993915s, +0.90%
* github.com/nsqio/nsq/apps/nsqd: from 2.270817882s to 3.734972024s, +64.48%
* github.com/mholt/caddy: from 278.433244ms to 268.490585ms, -3.57%

## Binary size:

* github.com/coreos/etcd/cmd: from 25848182 to 25848182, ~
* github.com/boltdb/bolt/cmd/bolt: from 2665904 to 2665904, ~
* github.com/gogits/gogs: from 23603938 to 23603938, ~
* github.com/spf13/hugo: from 15171903 to 15171903, ~
* github.com/nsqio/nsq/apps/nsqd: from 10033736 to 10033736, ~
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               192           191           -0.52%
BenchmarkMsgpUnmarshal-4             405           406           +0.25%
BenchmarkEasyJsonMarshal-4           1609          1590          -1.18%
BenchmarkEasyJsonUnmarshal-4         1519          1534          +0.99%
BenchmarkFlatBuffersMarshal-4        362           363           +0.28%
BenchmarkFlatBuffersUnmarshal-4      292           294           +0.68%
BenchmarkGogoprotobufMarshal-4       163           164           +0.61%
BenchmarkGogoprotobufUnmarshal-4     260           259           -0.38%

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

* [cmd/vet: -lostcancel: check for discarded result of context.WithCancel](https://github.com/golang/go/commit/b65cb7f198836faf6605051b95bd60a169fa5e8b)

## GIT Log

```
commit aa6345d3c91167f1e81bff9e8655e7aaac7762bd
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 21 15:17:56 2016 -0700

    testing: document that logs are dumped to standard output
    
    Since at least 1.0.3, the testing package has said that logs are dumped
    to standard error, but has in fact dumped the logs to standard output.
    We could change to dump to standard error, but after doing it this way
    for so long I think it's better to change the docs.
    
    Fixes #16138.
    
    Change-Id: If39c7ce91f51c7113f33ebabfb8f84fd4611b9e1
    Reviewed-on: https://go-review.googlesource.com/24311
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit bc3bcfd4e76195ead984e9d2ae1a1783d1272dc4
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 22 09:47:42 2016 -0700

    html/template: update security model link
    
    Fixes #16148.
    
    Change-Id: Ifab773e986b768602476824d005eea9200761236
    Reviewed-on: https://go-review.googlesource.com/24327
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit b31ec5c564f02cf48d177853fd7bff9892be7ce6
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 22 06:27:31 2016 -0700

    cmd/yacc: error rather than panic when TEMPSIZE is too small
    
    I tried simply increasing the size of the slice but then I got an error
    because NSTATES was too small. Leaving a real fix for after 1.7.
    
    Update #16144.
    
    Change-Id: I8676772cb79845dd4ca1619977d4d54a2ce6de59
    Reviewed-on: https://go-review.googlesource.com/24321
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 0dae2fd149537a0385b48bbd1564b3cfefa1c85b
Author: Keith Randall <khr@golang.org>
Date:   Wed Jun 1 14:18:00 2016 -0700

    cmd/objdump: fix disassembly suffixes
    
    MOVB $1, (AX) was being disassembled as MOVL $1, (AX).
    
    Use the memory size to override the standard size.
    Fix the tests.
    
    Fixes #15922
    
    Change-Id: If92fe74c33a21e5427c8c5cc97dd15e087edb860
    Reviewed-on: https://go-review.googlesource.com/23608
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 395f6ebaf9fd2954f5db3da4c7ad2817b0aa7252
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Jun 17 16:09:15 2016 -0400

    CONTRIBUTORS: add people who contributed to s390x port (IBM CLA)
    
    Add Bill O'Farrell (corporate CLA for IBM)
    Add Karan Dhiman (corporate CLA for IBM)
    Add Sam Ding (corporate CLA for IBM)
    Add Tristan Amini (corporate CLA for IBM)
    Add Yu Heng Zhang (corporate CLA for IBM)
    Add Yu Xuan Zhang (corporate CLA for IBM)
    
    Change-Id: I9ab15e33954afc2c208fc2e420a72c5a4d865f9b
    Reviewed-on: https://go-review.googlesource.com/24350
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4764d6fd6e64d9d40e7131a3bf4ea0eb1507ef1f
Author: Alan Donovan <adonovan@google.com>
Date:   Wed Jun 22 10:41:30 2016 -0400

    cmd/vet/internal/cfg: don't crash on malformed goto statement
    
    Change-Id: Ib285c02e240f02e9d5511bd448163ec1d4e75516
    Reviewed-on: https://go-review.googlesource.com/24323
    Reviewed-by: Rob Pike <r@golang.org>

commit f2c13d713d85650e4a850813d64681d6be5d2e29
Author: Alan Donovan <adonovan@google.com>
Date:   Wed Jun 22 10:40:30 2016 -0400

    cmd/vet: fix a crash in lostcancel check
    
    Fixes issue 16143
    
    Change-Id: Id9d257aee54d31fbf0d478cb07339729cd9712c0
    Reviewed-on: https://go-review.googlesource.com/24325
    Reviewed-by: Rob Pike <r@golang.org>

commit 1f446432ddfd64f1507e7c85cd603d3c5ae60094
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Jun 21 14:27:40 2016 -0700

    cmd/compile: fix error msg mentioning different packages with same name
    
    This is a regression from 1.6. The respective code in importimport
    (export.go) was not exactly replicated with the new importer. Also
    copied over the missing cyclic import check.
    
    Added test cases.
    
    Fixes #16133.
    
    Change-Id: I1e0a39ff1275ca62a8054874294d400ed83fb26a
    Reviewed-on: https://go-review.googlesource.com/24312
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>

commit 845992eeed01643bfb2e88aa559413908b3cb508
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Jun 21 15:33:04 2016 -0700

    test: add -s flag to commands understood by run.go
    
    If -s is specified, each file is considered a separate
    package even if multiple files have the same package names.
    
    For instance, the action and flag "errorcheckdir -s"
    will compile all files in the respective directory as
    individual packages.
    
    Change-Id: Ic5c2f9e915a669433f66c2d3fe0ac068227a502f
    Reviewed-on: https://go-review.googlesource.com/24313
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit eaf4ad6f7493c222a6b27609fcb24e950eb540ab
Author: Alan Donovan <adonovan@google.com>
Date:   Tue Jun 21 11:11:50 2016 -0400

    doc: describe vet -lostcancel in go1.7 release notes
    
    Change-Id: Ie1c95fd0869307551bfcf76bf45c13372723fbba
    Reviewed-on: https://go-review.googlesource.com/24288
    Reviewed-by: Rob Pike <r@golang.org>

commit b65cb7f198836faf6605051b95bd60a169fa5e8b
Author: Alan Donovan <adonovan@google.com>
Date:   Tue Jun 14 23:50:39 2016 -0400

    cmd/vet: -lostcancel: check for discarded result of context.WithCancel
    
    The cfg subpackage builds a control flow graph of ast.Nodes.
    The lostcancel module checks this graph to find paths, from a call to
    WithCancel to a return statement, on which the cancel variable is
    not used.  (A trivial case is simply assigning the cancel result to
    the blank identifier.)
    
    In a sample of 50,000 source files, the new check found 2068 blank
    assignments and 118 return-before-cancel errors.  I manually inspected
    20 of the latter and didn't find a single false positive among them.
    
    Change-Id: I84cd49445f9f8d04908b04881eb1496a96611205
    Reviewed-on: https://go-review.googlesource.com/24150
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit d28242724872c6ab82d53a71fc775095d1171ee7
Author: qeed <qeed.quan@gmail.com>
Date:   Mon Jun 20 21:11:53 2016 -0400

    cmd/cgo: error, not panic, if not enough arguments to function
    
    Fixes #16116.
    
    Change-Id: Ic3cb0b95382bb683368743bda49b4eb5fdcc35c0
    Reviewed-on: https://go-review.googlesource.com/24286
    Reviewed-by: Emmanuel Odeke <emm.odeke@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 86b031018550ff1848ebe7c471c54c5a58fb1a3d
Author: Rob Pike <r@golang.org>
Date:   Mon Jun 20 15:39:03 2016 -0700

    text/template: clarify the default formatting used for values
    
    Fixes #16105.
    
    Change-Id: I94467f2adf861eb38f3119ad30d46a87456d5305
    Reviewed-on: https://go-review.googlesource.com/24281
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 252eda470a3684a1ead5956f7e703532f4213f11
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jun 20 14:00:58 2016 -0700

    cmd/pprof: don't use offset if we don't have a start address
    
    The test is in the runtime package because there are other tests of
    pprof there. At some point we should probably move them all into a pprof
    testsuite.
    
    Fixes #16128.
    
    Change-Id: Ieefa40c61cf3edde11fe0cf04da1debfd8b3d7c0
    Reviewed-on: https://go-review.googlesource.com/24274
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 09834d1c082a2437b12584bebaa7353377e66f1a
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jun 20 15:53:11 2016 -0700

    runtime: panic with the right error on iface conversion
    
    A straight conversion from a type T to an interface type I, where T does
    not implement I, should always panic with an interface conversion error
    that shows the missing method.  This was not happening if the conversion
    was done once using the comma-ok form (the result would not be OK) and
    then again in a straight conversion.  Due to an error in the runtime
    package the second conversion was failing with a nil pointer
    dereference.
    
    Fixes #16130.
    
    Change-Id: I8b9fca0f1bb635a6181b8b76de8c2385bb7ac2d2
    Reviewed-on: https://go-review.googlesource.com/24284
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Michel Lespinasse <walken@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 659b9a19aa509df35f984276e177c68ff7f6f632
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jun 20 09:16:17 2016 -0700

    runtime: set PPROF_TMPDIR before running pprof
    
    Fixes #16121.
    
    Change-Id: I7b838fb6fb9f098e6c348d67379fdc81fb0d69a4
    Reviewed-on: https://go-review.googlesource.com/24270
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 109823ec93e690f2a401c316210ee86bde53d6bf
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon Jun 20 13:30:04 2016 +1000

    cmd/go: for generate, use build context values for GOOS/GOARCH
    
    Fixes #16120
    
    Change-Id: Ia352558231e00baab5c698e93d7267564c07ec0c
    Reviewed-on: https://go-review.googlesource.com/24242
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e1a6e71e746d511570f269d43b9abf838505a8e5
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jun 20 15:43:21 2016 -0700

    test: add missing copyright notice
    
    Change-Id: I2a5353203ca2958fa37fc7a5ea3f22ad4fc62b0e
    Reviewed-on: https://go-review.googlesource.com/24282
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 349f0fb89a11c9aa579bc0facab375d6ffb7ad43
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon Jun 20 15:33:52 2016 +1000

    doc: update architectures on source install instructions
    
    Fixes #16099
    
    Change-Id: I334c1f04dfc98c4a07e33745819d890b5fcb1673
    Reviewed-on: https://go-review.googlesource.com/24243
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 20fd1bb8bdbd8ead3d707ad13b560a375a7b3967
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Mon Jun 20 10:40:07 2016 +0900

    doc/go1.7.html: don't mention obsolete RFC
    
    Change-Id: Ia978c10a97e4c24fd7cf1fa4a7c3bd886d20bbf8
    Reviewed-on: https://go-review.googlesource.com/24241
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 153d31da1629facdc855ad0e4e91369ec2124ac7
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Jun 19 17:30:29 2016 -0700

    doc/go1.7.html: net/http RFC 2616 conformation + timeoutHandler on empty body
    
    - Mention RFC 2616 conformation in which the server now only sends one
    "Transfer-Encoding" header when "chunked" is explicitly set.
    - Mention that a timeout handler now sends a 200 status code on
    encountering an empty response body instead of sending back 0.
    
    Change-Id: Id45e2867390f7e679ab40d7a66db1f7b9d92ce17
    Reviewed-on: https://go-review.googlesource.com/24250
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 691c5c156878799ec15683f10e24d9a924ea7996
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Fri Jun 17 14:05:35 2016 +1000

    debug/pe: handle files with no string table
    
    pecoff.doc (https://goo.gl/ayvckk) in section 5.6 says:
    
    Immediately following the COFF symbol table is the COFF string table.
    The position of this table is found by taking the symbol table address
    in the COFF header, and adding the number of symbols multiplied by
    the size of a symbol.
    
    So it is unclear what to do when symbol table address is 0.
    Lets assume executable does not have any string table.
    
    Added new test with executable with no symbol table. The
    
    gcc -s testdata\hello.c -o testdata\gcc-386-mingw-no-symbols-exec.
    
    command was used to generate the executable.
    
    Fixes #16084
    
    Change-Id: Ie74137ac64b15daadd28e1f0315f3b62d1bf2059
    Reviewed-on: https://go-review.googlesource.com/24200
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f3d54789f764f0f6dba1d2f12ad01986d66ea31c
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Jun 16 18:12:48 2016 -0400

    cmd/compile: use power5 instructions for uint64 to float casts
    
    Use the same technique as mips64 for these casts (CL 22835).
    
    We could use the FCFIDU instruction for ppc64le however it seems
    better to keep it the same as ppc64 for now.
    
    Updates #15539, updates #16004.
    
    Change-Id: I550680e485327568bf3238c4615a6cc8de6438d7
    Reviewed-on: https://go-review.googlesource.com/24191
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9e8fa1e99c2003cee53a6630aea9d8a3627492ab
Author: Austin Clements <austin@google.com>
Date:   Fri Jun 17 11:07:56 2016 -0400

    runtime: eliminate poisonStack checks
    
    We haven't used poisonStack since we switched to 1-bit stack maps
    (4d0f3a1), but the checks are still there. However, nothing prevents
    us from genuinely allocating an object at this address on 32-bit and
    causing the runtime to crash claiming that it's found a bad pointer.
    
    Since we're not using poisonStack anyway, just pull it out.
    
    Fixes #15831.
    
    Change-Id: Ia6ef604675b8433f75045e369f5acd4644a5bb38
    Reviewed-on: https://go-review.googlesource.com/24211
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit fca9fc52c831ab6af56e30f8c48062a99ded2580
Author: Austin Clements <austin@google.com>
Date:   Thu Jun 16 15:41:33 2016 -0400

    runtime: fix stale comment in lfstack
    
    Change-Id: I6ef08f6078190dc9df0b2df4f26a76456602f5e8
    Reviewed-on: https://go-review.googlesource.com/24176
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 79f2f008a31f769ae3db684eb48d8baeda731c00
Author: Austin Clements <austin@google.com>
Date:   Thu Jun 16 14:38:33 2016 -0400

    cmd/dist: make zosarch.go deterministic
    
    Currently zosarch.go is written out in non-deterministic map order.
    Sort the keys and write it out in sorted order to make the generated
    file contents deterministic.
    
    Change-Id: Id490f0e8665a2c619c5a7a00a30f4fc64f333258
    Reviewed-on: https://go-review.googlesource.com/24174
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c3818e56d0f60493a63b2bb03a09f261d3e0ada2
Author: Hana Kim <hyangah@gmail.com>
Date:   Wed Jun 15 12:53:05 2016 -0400

    internal/trace: err if binary is not supplied for old trace
    
    Change-Id: Id25c90993c4cbb7449d7031301b6d214a67d7633
    Reviewed-on: https://go-review.googlesource.com/24134
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
```
