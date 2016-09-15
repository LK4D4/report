# June 16, 2016 Report

Number of commits: 32

## Compilation time

* github.com/coreos/etcd/cmd: from 14.458248604s to 14.518299282s, +0.42%
* github.com/boltdb/bolt/cmd/bolt: from 547.141317ms to 553.29184ms, +1.12%
* github.com/gogits/gogs: from 12.988071383s to 12.911730568s, -0.59%
* github.com/spf13/hugo: from 6.724718371s to 8.411540397s, +25.08%
* github.com/influxdata/influxdb/cmd/influxd: from 6.624849543s to 6.09212415s, -8.04%
* github.com/nsqio/nsq/apps/nsqd: from 2.179579787s to 2.216044118s, +1.67%
* github.com/mholt/caddy: from 270.394164ms to 267.701648ms, -1.00%

## Binary size:

* github.com/coreos/etcd/cmd: from 25685467 to 25727105, +0.16%
* github.com/boltdb/bolt/cmd/bolt: from 2652945 to 2665904, +0.49%
* github.com/gogits/gogs: from 23557580 to 23595278, +0.16%
* github.com/spf13/hugo: from 15134034 to 15167579, +0.22%
* github.com/influxdata/influxdb/cmd/influxd: from 16260059 to 16293562, +0.21%
* github.com/nsqio/nsq/apps/nsqd: from 10008482 to 10033736, +0.25%
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               192           192           +0.00%
BenchmarkMsgpUnmarshal-4             404           402           -0.50%
BenchmarkEasyJsonMarshal-4           1581          1592          +0.70%
BenchmarkEasyJsonUnmarshal-4         1515          1525          +0.66%
BenchmarkFlatBuffersMarshal-4        365           359           -1.64%
BenchmarkFlatBuffersUnmarshal-4      290           291           +0.34%
BenchmarkGogoprotobufMarshal-4       164           164           +0.00%
BenchmarkGogoprotobufUnmarshal-4     250           260           +4.00%

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
BenchmarkEasyJsonUnmarshal-4         159           160           +0.63%
BenchmarkFlatBuffersMarshal-4        0             0             +0.00%
BenchmarkFlatBuffersUnmarshal-4      112           112           +0.00%
BenchmarkGogoprotobufMarshal-4       64            64            +0.00%
BenchmarkGogoprotobufUnmarshal-4     96            96            +0.00%
```

## Highlights: 


## GIT Log

```
commit ea2ac3fe5fb2011b077809e60bc018e0c6caa66c
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 15 10:36:48 2016 -0700

    runtime: remove useless loop from CgoCCodeSIGPROF test program
    
    I verified that the test fails if I undo the change that it tests for.
    
    Updates #14732.
    
    Change-Id: Ib30352580236adefae946450ddd6cd65a62b7cdf
    Reviewed-on: https://go-review.googlesource.com/24151
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit d78d0de4d15019f87b04b49ba640a455f7c42512
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Jun 15 12:33:47 2016 -0700

    go/ast: fix comments misinterpreted as documentation
    
    The comments describing blocks of Pos/End implementations for various
    nodes types are being misinterpreted as documentation for BadDecl,
    BadExpr, BadStmt, and ImportSpec's Pos methods.
    
    Change-Id: I935b0bc38dbc13e9305f3efeb437dd3a6575d9a1
    Reviewed-on: https://go-review.googlesource.com/24152
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 26d6dc6bf8a7e5487844a63aa26a4de3afdd688e
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 14 14:38:22 2016 -0700

    runtime: if the test program hangs, try to get a stack trace
    
    This is an attempt to get more information for #14809, which seems to
    occur rarely.
    
    Updates #14809.
    
    Change-Id: Idbeb136ceb57993644e03266622eb699d2685d02
    Reviewed-on: https://go-review.googlesource.com/24110
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Austin Clements <austin@google.com>

commit 48cc3c4b587f9549f7426776d032da99b3ba471b
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Jun 14 15:33:15 2016 -0400

    syscall: skip TestUnshare if kernel does not support net namespace
    
    Fixes #16056.
    
    Change-Id: Ic3343914742713851b8ae969b101521f25e85e7b
    Reviewed-on: https://go-review.googlesource.com/24132
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0ec62565f911575beaf7d047dfe1eae2ae02bf67
Author: Andrew Gerrand <adg@golang.org>
Date:   Wed Jun 15 10:52:42 2016 +1000

    net/http: pass through server side Transfer-Encoding headers
    
    Fixes #16063
    
    Change-Id: I2e8695beb657b0aef067e83f086828d8857787ed
    Reviewed-on: https://go-review.googlesource.com/24130
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c4692da9231c244a1275d42055e703b3f1dac25b
Author: Sameer Ajmani <sameer@golang.org>
Date:   Tue Jun 14 16:48:42 2016 -0400

    context: document how to release resources associated with Contexts.
    
    Some users don't realize that creating a Context with a CancelFunc
    attaches a subtree to the parent, and that that subtree is not released
    until the CancelFunc is called or the parent is canceled.  Make this
    clear early in the package docs, so that people learning about this
    package have the right conceptual model.
    
    Change-Id: I7c77a546c19c3751dd1f3a5bc827ad106dd1afbf
    Reviewed-on: https://go-review.googlesource.com/24090
    Reviewed-by: Alan Donovan <adonovan@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 68697a3e82e19fef04c6af4a02340a1aa6e3bcf2
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 14 15:32:39 2016 -0700

    net: don't run TestLookupDotsWithLocalSource in short mode
    
    Do run it on the builders, though.
    
    Fixes #15881.
    
    Change-Id: Ie42204d553cb18547ffd6441afc261717bbd9205
    Reviewed-on: https://go-review.googlesource.com/24111
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9208ed322459809cf26f65485d0e6d248dadb830
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Jun 15 01:09:51 2016 +0900

    os: fix blockUntilWaitable on freebsd/{386,arm}
    
    The previous fix was wrong because it had two misunderstandings on
    freebsd32 calling convention like the following:
    - 32-bit id1 implies that it is the upper half of 64-bit id, indeed it
      depends on machine endianness.
    - 32-bit ARM calling convension doesn't conform to freebsd32_args,
      indeed it does.
    
    This change fixes the bugs and makes blockUntilWaitable work correctly
    on freebsd/{386,arm}.
    
    Fixes #16064.
    
    Change-Id: I820c6d01d59a43ac4f2ab381f757c03b14bca75e
    Reviewed-on: https://go-review.googlesource.com/24064
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit af0fc83985860776551d15be3a8fefde35514bcb
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Jun 14 10:20:11 2016 -0400

    cmd/compile, etc: handle many struct fields
    
    This adds 8 bytes of binary size to every type that has methods. It is
    the smallest change I could come up with for 1.7.
    
    Fixes #16037
    
    Change-Id: Ibe15c3165854a21768596967757864b880dbfeed
    Reviewed-on: https://go-review.googlesource.com/24070
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 53242e49b127ede6d7b258c7e90c39a5afa70c25
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 7 09:42:48 2016 -0700

    crypto/x509: don't ignore asn1.Marshal error
    
    I don't see how the call could fail, so, no test. Just a code cleanup in
    case it can fail in the future.
    
    Fixes #15987.
    
    Change-Id: If4af0d5e7d19cc8b13fb6a4f7661c37fb0015e83
    Reviewed-on: https://go-review.googlesource.com/23860
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit 0deb49f9c09d15bf0e4c5ec843bd374f9a377e97
Author: Ian Lance Taylor <iant@golang.org>
Date:   Mon Jun 13 11:55:30 2016 -0700

    cmd/go: include .syso files even if CGO_ENABLED=0
    
    A .syso file may include information that should go into the object file
    that is not object code, and should be included even if not using cgo.
    The example in the issue is a Windows manifest file.
    
    Fixes #16050.
    
    Change-Id: I1f4f3f80bb007e84d153ca2d26e5919213ea4f8d
    Reviewed-on: https://go-review.googlesource.com/24032
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 9273e25eccbe82edff839b125b49bfb5578f24eb
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Jun 10 10:41:36 2016 -0700

    cmd/go: remove obsolete comment referring to deleted parameter
    
    The dir parameter was removed in https://golang.org/cl/5732045.
    
    Fixes #15503.
    
    Change-Id: I02a6d8317233bea08633715a095ea2514822032b
    Reviewed-on: https://go-review.googlesource.com/24011
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitri Shuralyov <shurcool@gmail.com>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit cab87a60dec2c771c12ba08b82bb7645228192e6
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Tue Jun 14 10:14:28 2016 +0900

    os: fix build on freebsd/arm
    
    Change-Id: I21fad94ff94e342ada18e0e41ca90296d030115f
    Reviewed-on: https://go-review.googlesource.com/24061
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5d876e3e2eb3a30a8c66888912cf41785fa65a96
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Jun 11 19:55:34 2016 +0900

    os: use wait6 to avoid wait/kill race on freebsd
    
    This change is a followup to https://go-review.googlesource.com/23967
    for FreeBSD.
    
    Updates #13987.
    Updates #16028.
    
    Change-Id: I0f0737372fce6df89d090fe9847305749b79eb4c
    Reviewed-on: https://go-review.googlesource.com/24021
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ccd9a55609fcc8814146a7c61898b47e3a7aea7d
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Sat Jun 11 19:36:17 2016 +0900

    os: use waitid to avoid wait/kill race on darwin
    
    This change is a followup to https://go-review.googlesource.com/23967
    for Darwin.
    
    Updates #13987.
    Updates #16028.
    
    Change-Id: Ib1fb9f957fafd0f91da6fceea56620e29ad82b00
    Reviewed-on: https://go-review.googlesource.com/24020
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 84d8aff94cf48439047c7edc68ae2ea0aac6ddf5
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Jun 7 21:46:25 2016 -0700

    runtime: collect stack trace if SIGPROF arrives on non-Go thread
    
    Fixes #15994.
    
    Change-Id: I5aca91ab53985ac7dcb07ce094ec15eb8ec341f8
    Reviewed-on: https://go-review.googlesource.com/23891
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5701174c52a2d42621ec3c5c59dca3bde9a14bc6
Author: Keith Randall <khr@golang.org>
Date:   Sat Jun 11 17:12:28 2016 -0700

    cmd/link: put padding between functions, not at the end of a function
    
    Functions should be declared to end after the last real instruction, not
    after the last padding byte. We achieve this by adding the padding while
    assembling the text section in the linker instead of adding the padding
    to the function symbol in the compiler. This change makes dtrace happy.
    
    TODO: check that this works with external linking
    
    Fixes #15969
    
    Change-Id: I973e478d0cd34b61be1ddc55410552cbd645ad62
    Reviewed-on: https://go-review.googlesource.com/24040
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 595426c0d903a3686bdfe6d0e8ef268a60c19896
Author: David Chase <drchase@google.com>
Date:   Fri Jun 10 11:51:46 2016 -0400

    cmd/compile: fix OASWB rewriting in racewalk
    
    Special case for rewriting OAS inits omitted OASWB, added
    that and OAS2FUNC.  The special case cannot be default case,
    that causes racewalk to fail in horrible ways.
    
    Fixes #16008.
    
    Change-Id: Ie0d2f5735fe9d8255a109597b36d196d4f86703a
    Reviewed-on: https://go-review.googlesource.com/23954
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2ba3d5fc9661846cc3ef66ffc2b8bf2f91909d73
Author: Dmitri Shuralyov <shurcooL@gmail.com>
Date:   Fri Jun 10 01:38:43 2016 -0700

    cmd/go: remove invalid space in import comment docs
    
    Generate package comment in alldocs.go using line comments rather than
    general comments. This scales better, general comments cannot contain the
    "*/" character sequence. Line comments do not have any restrictions on
    the comment text that can be contained.
    
    Remove the dependency on sed, which is not cross-platform, not go-gettable
    external command.
    
    Remove trailing whitespace from usage string in test.go. It's unnecessary.
    
    Fixes #16030.
    
    Change-Id: I3c0bc9955e7c7603c3d1fb4878218b0719d02e04
    Reviewed-on: https://go-review.googlesource.com/23968
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c83e6f50d983d81166d21736ff9ab0ad2182f0fa
Author: Keith Randall <khr@golang.org>
Date:   Thu May 26 08:56:49 2016 -0700

    runtime: aeshash, xor seed in earlier
    
    Instead of doing:
    
    x = input
    one round of aes on x
    x ^= seed
    two rounds of aes on x
    
    Do:
    
    x = input
    x ^= seed
    three rounds of aes on x
    
    This change provides some additional seed-dependent scrambling
    which should help prevent collisions.
    
    Change-Id: I02c774d09c2eb6917cf861513816a1024a9b65d7
    Reviewed-on: https://go-review.googlesource.com/23577
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cea29c4a358004d84d8711a07628c2f856b381e8
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jun 9 22:24:40 2016 -0700

    os: on GNU/Linux use waitid to avoid wait/kill race
    
    On systems that support the POSIX.1-2008 waitid function, we can use it
    to block until a wait will succeed. This avoids a possible race
    condition: if a program calls p.Kill/p.Signal and p.Wait from two
    different goroutines, then it is possible for the wait to complete just
    before the signal is sent. In that case, it is possible that the system
    will start a new process using the same PID between the wait and the
    signal, causing the signal to be sent to the wrong process. The
    Process.isdone field attempts to avoid that race, but there is a small
    gap of time between when wait returns and isdone is set when the race
    can occur.
    
    This CL avoids that race by using waitid to wait until the process has
    exited without actually collecting the PID. Then it sets isdone, then
    waits for any active signals to complete, and only then collects the PID.
    
    No test because any plausible test would require starting enough
    processes to recycle all the process IDs.
    
    Update #13987.
    Update #16028.
    
    Change-Id: Id2939431991d3b355dfb22f08793585fc0568ce8
    Reviewed-on: https://go-review.googlesource.com/23967
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e980a3d8856ec3b4f11daa7e5ec417ad4f5c5256
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Jun 10 10:27:37 2016 -0700

    go/parser: document that parse functions need valid token.FileSet
    
    + panic with explicit error if no file set it provided
    
    (Not providing a file set is invalid use of the API; panic
    is the appropriate action rather than returning an error.)
    
    Fixes #16018.
    
    Change-Id: I207f5b2a2e318d65826bdd9522fce46d614c24ee
    Reviewed-on: https://go-review.googlesource.com/24010
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fee02d270bf850d5b390000d8545c3609718e9a5
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jun 9 06:42:42 2016 -0700

    cmd/go: clarify go get documentation
    
    Make the documentation for `go get` match the documentation for `go
    install`, since `go get` essentially invokes `go install`.
    
    Update #15825.
    
    Change-Id: I374d80efd301814b6d98b86b7a4a68dd09704c92
    Reviewed-on: https://go-review.googlesource.com/23925
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 5f3eb432884ae1b3c61977d93a4bdf0263fcdce6
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Jun 9 12:01:43 2016 -0400

    syscall: add a padding field to EpollEvent on s390x
    
    Fixes #16021.
    
    Change-Id: I55df38bbccd2641abcb54704115002a9aa04325d
    Reviewed-on: https://go-review.googlesource.com/23962
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8042bfe347d5b2c4f6af372b09b23c8945fb196e
Author: Jess Frazelle <me@jessfraz.com>
Date:   Fri May 20 14:35:28 2016 -0700

    encoding/csv: update doc about comments whitespace
    
    This patch updates the doc about comments whitespace for the
    encoding/csv package to reflect that leading whitespace before
    the hash will treat the line as not a comment.
    
    Fixes #13775.
    
    Change-Id: Ia468c75b242a487b4b2b4cd3d342bfb8e07720ba
    Reviewed-on: https://go-review.googlesource.com/23302
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit cbc26869b7835e45359dad7dfb70e85c02c820cd
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Jun 8 22:22:35 2016 -0400

    runtime: set $sp before $pc in gdb python script
    
    When setting $pc, gdb does a backtrace using the current value of $sp,
    and it may complain if $sp does not match that $pc (although the
    assignment went through successfully).
    
    This happens with ARM SSA backend: when setting $pc it prints
    > Cannot access memory at address 0x0
    
    As well as occasionally on MIPS64:
    > warning: GDB can't find the start of the function at 0xc82003fe07.
    > ...
    
    Setting $sp before setting $pc makes it happy.
    
    Change-Id: Idd96dbef3e9b698829da553c6d71d5b4c6d492db
    Reviewed-on: https://go-review.googlesource.com/23940
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit e3f1c66f313a59888620c415163b93c12153574e
Author: Keith Randall <khr@golang.org>
Date:   Thu Jun 9 11:07:36 2016 -0700

    cmd/compile: for tail calls in stubs, ensure args are alive
    
    The generated code for interface stubs sometimes just messes
    with a few of the args and then tail-calls to the target routine.
    The args that aren't explicitly modified appear to not be used.
    But they are used, by the thing we're tail calling.
    
    Fixes #16016
    
    Change-Id: Ib9b3a8311bb714a201daee002885fcb59e0463fa
    Reviewed-on: https://go-review.googlesource.com/23960
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 1bdf1c3024d75a3c4913d031d55257b311f0133f
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Jun 9 11:19:26 2016 -0700

    cmd/cgo: fix use of unsafe argument in new deferred function
    
    The combination of https://golang.org/cl/23650 and
    https://golang.org/cl/23675 did not work--they were tested separately
    but not together.
    
    The problem was that 23650 introduced deferred argument checking, and
    the deferred function loses the type that 23675 started requiring. The
    fix is to go back to using an empty interface type in a deferred
    argument check.
    
    No new test required--fixes broken build.
    
    Change-Id: I5ea023c5aed71d70e57b11c4551242d3ef25986d
    Reviewed-on: https://go-review.googlesource.com/23961
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 837984f37291f4fc48f9c99b65b0ab3f050bf4b9
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 1 21:32:47 2016 -0700

    cmd/cgo: use function arg type for _cgoCheckPointerN function
    
    When cgo writes a _cgoCheckPointerN function to handle unsafe.Pointer,
    use the function's argument type rather than interface{}. This permits
    type errors to be detected at build time rather than run time.
    
    Fixes #13830.
    
    Change-Id: Ic7090905e16b977e2379670e0f83640dc192b565
    Reviewed-on: https://go-review.googlesource.com/23675
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 894803c11e4eab128869be759463510580a68602
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Jun 8 21:15:26 2016 -0700

    time: document that RFC822/1123 don't parse all RFC formats
    
    Fixes #14505.
    
    Change-Id: I46196b26c9339609e6e3ef9159de38c5b50c2a1b
    Reviewed-on: https://go-review.googlesource.com/23922
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Emmanuel Odeke <emm.odeke@gmail.com>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit e2a30b8ffba065ea45bb1e18d1bb584898e3bf11
Author: Kenny Grant <kennygrant@gmail.com>
Date:   Sat May 21 17:19:22 2016 +0100

    time: genzabbrs.go skips Feb when checking months
    
    getAbbrs looks like it is checking each month looking for a change
    in the time zone abbreviation, but starts in Dec of the previous year
    and skips the month of February because of the overflow rules for
    AddDate. Changing the day to 1 starts at Jan 1 and tries all months
    in the current year. This isn't very important or likely to change
    output as zones usually span several months. Discovered when
    looking into time.AddDate behavior when adding months.
    
    Change-Id: I685254c8d21c402ba82cc4176e9a86b64ce8f7f7
    Reviewed-on: https://go-review.googlesource.com/23322
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6662897b2a3260393fff9dcf64faf3abfc773181
Author: Jason Barnett <jason.w.barnett@gmail.com>
Date:   Tue May 24 15:50:02 2016 -0400

    crypto/subtle: expand abbreviation to eliminate confusion
    
    Change-Id: I68d66fccf9cc8f7137c92b94820ce7d6f478a185
    Reviewed-on: https://go-review.googlesource.com/23310
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
```
