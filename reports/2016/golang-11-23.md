# November 23, 2016 Report

Number of commits: 85

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 540.191512ms to 571.904842ms, +5.87%
* github.com/coreos/etcd/cmd/etcd: from 11.784349623s to 12.124302023s, +2.88%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 34.693633515s to 33.701957109s, -2.86%
* github.com/influxdata/influxdb/cmd/influxd: from 6.858310144s to 6.77067861s, -1.28%
* github.com/junegunn/fzf/src/fzf: from 1.054331168s to 999.16903ms, -5.23%
* github.com/mholt/caddy/caddy: from 5.377422595s to 5.611907361s, +4.36%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.228298513s to 1.415275784s, +15.22%
* github.com/nsqio/nsq/apps/nsqd: from 1.753382784s to 2.039412433s, +16.31%
* github.com/prometheus/prometheus/cmd/prometheus: from 39.864090897s to 40.298231132s, +1.09%
* github.com/spf13/hugo: from 5.844335307s to 6.169617139s, +5.57%
* golang.org/x/tools/cmd/guru: from 2.590902304s to 2.726630482s, +5.24%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 3069331 to 3196545, +4.14%
* github.com/coreos/etcd/cmd/etcd: from 25105000 to 26648088, +6.15%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 31188864 to 31867520, +2.18%
* github.com/influxdata/influxdb/cmd/influxd: from 16459661 to 17109623, +3.95%
* github.com/junegunn/fzf/src/fzf: from 3079424 to 3182008, +3.33%
* github.com/mholt/caddy/caddy: from 15041194 to 15507015, +3.10%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4247740 to 4350235, +2.41%
* github.com/nsqio/nsq/apps/nsqd: from 10103027 to 10410944, +3.05%
* github.com/prometheus/prometheus/cmd/prometheus: from 60854744 to 62387437, +2.52%
* github.com/spf13/hugo: from 16035035 to 16471342, +2.72%
* golang.org/x/tools/cmd/guru: from 7957974 to 8234248, +3.47%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               187           193           +3.21%
BenchmarkMsgpUnmarshal-4             369           385           +4.34%
BenchmarkEasyJsonMarshal-4           1410          1485          +5.32%
BenchmarkEasyJsonUnmarshal-4         1770          2442          +37.97%
BenchmarkFlatBuffersMarshal-4        358           510           +42.46%
BenchmarkFlatBuffersUnmarshal-4      293           303           +3.41%
BenchmarkGogoprotobufMarshal-4       155           163           +5.16%
BenchmarkGogoprotobufUnmarshal-4     247           256           +3.64%

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

* [doc: add FAQ: why no conversion from []T1 to []T2?](https://github.com/golang/go/commit/fbf92436b95d91151ce6717f40c46614ee68d487)
* [runtime: exit idle worker if there's higher-priority work](https://github.com/golang/go/commit/49ea9207b6512c2400de11bc097d974bb527ba63)
* [runtime: wake idle Ps when enqueuing GC work](https://github.com/golang/go/commit/0bae74e8c9b5fab3baf61bde0169f4aa5e287bdc)

## GIT Log

```
commit 3f69822a9a68c76a8562a43f226951c33ed5694a
Author: Russ Cox <rsc@golang.org>
Date:   Tue Nov 22 22:46:39 2016 -0500

    math/rand: export Source64, mainly for documentation value
    
    There is some code value too: types intending to implement
    Source64 can write a conversion confirming that.
    
    For #4254 and the Go 1.8 release notes.
    
    Change-Id: I7fc350a84f3a963e4dab317ad228fa340dda5c66
    Reviewed-on: https://go-review.googlesource.com/33456
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2f0a306d283c21b98089706878190c34bd591b9a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 23 03:46:58 2016 +0000

    doc: add net/http section to go1.8.html
    
    TBR=See https://golang.org/cl/33244 and review there.
    
    Updates #17929
    
    Change-Id: I752ec7a6d086f370feaf3cf282708620e891079b
    Reviewed-on: https://go-review.googlesource.com/33478
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 75c1381176e98357b01a67af7e9dbaf68de7fdff
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 23 01:43:33 2016 +0000

    cmd/gofmt: don't call Chmod on windows
    
    Fixes #18026
    
    Change-Id: Id510f427ceffb2441c3d6f5bb5c93244e46c6497
    Reviewed-on: https://go-review.googlesource.com/33477
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 8ace3461a46d09466f0e1a8ea3ff41ac9cbd46f4
Author: Elias Naur <elias.naur@gmail.com>
Date:   Wed Nov 23 00:50:46 2016 +0100

    doc: add SIGPIPE behaviour change to go1.8.txt
    
    CL 32796 changes the SIGPIPE behaviour for c-archive and c-shared
    programs. Add it to go1.8.txt.
    
    Change-Id: I31200187033349c642965a4bb077bcc77d5329a3
    Reviewed-on: https://go-review.googlesource.com/33397
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b1dbc9f8c0257aca065126d8dfc3ebfe0be2aada
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Nov 22 15:50:03 2016 -0800

    reflect: fix typo in comment
    
    Sigh, forgot to run `git mail`.
    
    Change-Id: Idc49be2bb20d6f0e392cb472a63267ffee2ca22c
    Reviewed-on: https://go-review.googlesource.com/33476
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 50c4dbced991a01d4d707c00dc40c1a6366e9458
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Nov 22 15:50:03 2016 -0800

    reflect: fix size of StructOf ending in zero-sized field
    
    Update #9401.
    Fixes #18016.
    
    Change-Id: Icc24dd10dab1ad8e5cf295e0727d437afa5025c0
    Reviewed-on: https://go-review.googlesource.com/33475
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit e12f6ee0ab1e7b4da818054d7ced4247c81fd164
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Mon Oct 31 07:58:41 2016 -0700

    database/sql: fix TestPendingConnsAfterErr
    
    TestPendingConnsAfterErr showed a failure on slower systems.
    Wait and check for the database to close all connections
    before pronouncing failure.
    
    A more careful method was attempted but the connection pool
    behavior is too dependent on the scheduler behavior to be
    predictable.
    
    Fixes #15684
    
    Change-Id: Iafdbc90ba51170c76a079db04c3d5452047433a4
    Reviewed-on: https://go-review.googlesource.com/33418
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 199d410df5f52926add3cb6e1ad972a1135b40b1
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Tue Nov 22 11:30:16 2016 -0800

    doc: fix typos in go1.8.html
    
    Change-Id: I51180e1c685e488f7ea4c51a63fd035148671b05
    Reviewed-on: https://go-review.googlesource.com/33470
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f756204f5bcbb6f62e187113eb44abd4e3ed4d7f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 22 11:57:49 2016 -0800

    doc: more go1.8.html content
    
    TBR=See https://golang.org/cl/33244 and review there.
    
    Updates #17929
    
    Change-Id: I7cb0b666469dba35426d1f0ae1b185e0bdfeac05
    Reviewed-on: https://go-review.googlesource.com/33474
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 46323795130e92f26cbbb6bc744298edc8c443e9
Author: David du Colombier <0intro@gmail.com>
Date:   Tue Nov 22 16:19:25 2016 +0100

    cmd/go: print CC environment variables on Plan 9
    
    This changes makes the output of `go env` the same
    as on other operating systems.
    
    Fixes #18013.
    
    Change-Id: I3079e14dcf7b30c75ec3fde6c78cb95721111320
    Reviewed-on: https://go-review.googlesource.com/33396
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5508561180950534fd623fb61a8a8f357ad50e4c
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Nov 22 15:39:51 2016 -0500

    runtime/pprof/internal/protopprof: fix test on s390x
    
    Applies the fix from CL 32920 to the new test TestSampledHeapAllocProfile
    introduced in CL 33422. The test should be skipped rather than fail if
    there is only one executable region of memory.
    
    Updates #17852.
    
    Change-Id: Id8c47b1f17ead14f02a58a024c9a04ebb8ec0429
    Reviewed-on: https://go-review.googlesource.com/33453
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f9feaffdf51967bc5cf3c9363db9ddee98c9b3a0
Author: Russ Cox <rsc@golang.org>
Date:   Sat Nov 12 23:01:37 2016 -0500

    runtime: do not print runtime panic frame at top of user stack
    
    The expected default behavior (no explicit GOTRACEBACK setting)
    is for the stack trace to start in user code, eliding unnecessary runtime
    frames that led up to the actual trace printing code. The idea was that
    the first line number printed was the one that crashed.
    
    For #5832 we added code to show 'panic' frames so that if code panics
    and then starts running defers and then we trace from there, the panic
    frame can help explain why the code seems to have made a call not
    present in the code. But that's only needed for panics between two different
    call frames, not the panic at the very top of the stack trace.
    Fix the fix to again elide the runtime code at the very top of the stack trace.
    
    Simple panic:
    
            package main
    
            func main() {
                    var x []int
                    println(x[1])
            }
    
    Before this CL:
    
            panic: runtime error: index out of range
    
            goroutine 1 [running]:
            panic(0x1056980, 0x1091bf0)
                    /Users/rsc/go/src/runtime/panic.go:531 +0x1cf
            main.main()
                    /tmp/x.go:5 +0x5
    
    After this CL:
    
            panic: runtime error: index out of range
    
            goroutine 1 [running]:
            main.main()
                    /tmp/x.go:5 +0x5
    
    Panic inside defer triggered by panic:
    
            package main
    
            func main() {
                    var x []int
                    defer func() {
                            println(x[1])
                    }()
                    println(x[2])
            }
    
    Before this CL:
    
            panic: runtime error: index out of range
                    panic: runtime error: index out of range
    
            goroutine 1 [running]:
            panic(0x1056aa0, 0x1091bf0)
                    /Users/rsc/go/src/runtime/panic.go:531 +0x1cf
            main.main.func1(0x0, 0x0, 0x0)
                    /tmp/y.go:6 +0x62
            panic(0x1056aa0, 0x1091bf0)
                    /Users/rsc/go/src/runtime/panic.go:489 +0x2cf
            main.main()
                    /tmp/y.go:8 +0x59
    
    The middle panic is important: it explains why main.main ended up calling main.main.func1 on a line that looks like a call to println. The top panic is noise.
    
    After this CL:
    
            panic: runtime error: index out of range
                    panic: runtime error: index out of range
    
            goroutine 1 [running]:
            main.main.func1(0x0, 0x0, 0x0)
                    /tmp/y.go:6 +0x62
            panic(0x1056ac0, 0x1091bf0)
                    /Users/rsc/go/src/runtime/panic.go:489 +0x2cf
            main.main()
                    /tmp/y.go:8 +0x59
    
    Fixes #17901.
    
    Change-Id: Id6d7c76373f7a658a537a39ca32b7dc23e1e76aa
    Reviewed-on: https://go-review.googlesource.com/33165
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ac1dbe6392f2d392f9554127f96597a9aaa721fd
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 22 08:24:13 2016 -0700

    doc: more go1.8.html content
    
    TBR=See https://golang.org/cl/33244 and review there.
    
    Updates #17929
    
    Change-Id: I37b49318a9203b16c0c788926039288b99a36ce5
    Reviewed-on: https://go-review.googlesource.com/33450
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 86ab09eed55e4b70e642ad97d243d3090550b624
Author: Michael Matloob <matloob@golang.org>
Date:   Mon Nov 21 12:10:07 2016 -0500

    runtime/pprof: generate heap profiles in compressed proto format
    
    When debug is 0, emit the compressed proto format.
    The debug>0 format stays the same.
    
    Updates #16093
    
    Change-Id: I45aa1874a22d34cf44dd4aa78bbff9302381cb34
    Reviewed-on: https://go-review.googlesource.com/33422
    Run-TryBot: Michael Matloob <matloob@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f88a33aeac427b57a9acb4ecf6e26d7191ab60cc
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 22 07:24:07 2016 -0500

    doc: go1.8.html updates from Joe Tsai
    
    Updates #17929
    
    Change-Id: Ibc711d39d9ff83458d213778117493796b678aa7
    Reviewed-on: https://go-review.googlesource.com/33437
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e47b7af6409e37bd395ff92dee9c98720d22fc26
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 15 08:13:46 2016 -0800

    doc: start of go1.8.html release notes
    
    Updates #17929
    
    Change-Id: Ie90736cfce3fc5f23cbe0a0f1971476705aac5f9
    Reviewed-on: https://go-review.googlesource.com/33436
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 323b5c9d37e3633c96a96303da71b5d45cc9bac6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 22 01:40:07 2016 +0000

    time: make Parse validate day's lower bound in addition to upper bound
    
    Day 0 is as invalid as day 32.
    
    Fixes #17874
    
    Change-Id: I52109d12bafd6d957d00c44d540cb88389fff0a7
    Reviewed-on: https://go-review.googlesource.com/33429
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 409a667f35f66b2b5fa219ee93f1a3ddc6e07b03
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 22 02:07:52 2016 +0000

    net/http: fix parallel tests using global DefaultTransport
    
    When I added t.Parallel to some tests earlier, I overlooked some using
    the global "Get" func, which uses DefaultTransport.
    
    The DefaultTransport can have its CloseIdleConnections called by other
    parallel tests. Use a private Transport instead.
    
    Fixes #18006
    
    Change-Id: Ia4faca5bac235cfa95dcf2703c25f3627112a5e9
    Reviewed-on: https://go-review.googlesource.com/33432
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 75055de84ab7ad0f36b4c93e5c851ea55b297c95
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Nov 15 21:15:40 2016 -0800

    runtime: sleep a bit to let a bad signal be delivered
    
    When we raise a signal that was delivered to C code, it's possible that
    the kernel will not deliver it immediately. This is especially possible
    on Darwin where we use send the signal to the entire process rather than
    just the current thread. Sleep for a millisecond after sending the
    signal to give it a chance to be delivered before we restore the Go
    signal handler. In most real cases the program is going to crash at this
    point, so sleeping is kind of irrelevant anyhow.
    
    Fixes #14809.
    
    Change-Id: Ib2c0d2c4e240977fb4535dc1dd2bdc50d430eb85
    Reviewed-on: https://go-review.googlesource.com/33300
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit e9ffda45c8b7d409a5b951d6a74b8241c026fad5
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Nov 15 14:48:54 2016 -0800

    cmd/go: don't clobber `go env GOGCCFLAGS`
    
    When CC is set in the environment, the mkEnv function sets its version
    of CC to the first word $CC and sets GOGCCFLAGS to the remainder. That
    worked since Go 1 but was broken accidentally by
    https://golang.org/cl/6409, which changed the code such that `go env`
    calls mkEnv twice. The second call to mkEnv would clobber GOGCCFLAGS
    based on the value of CC set by the first call. Go back to the old
    handling by only calling mkEnv once.
    
    Fixes #15457.
    
    Change-Id: I000a1ebcc48684667e48f2b9b24605867b9e06cd
    Reviewed-on: https://go-review.googlesource.com/33293
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 6f31abd23a6f768c21c8b308f355f3a1bae521d2
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Nov 21 16:58:55 2016 -0500

    cmd/compile, cmd/link: weak relocation for ptrTo
    
    Introduce R_WEAKADDROFF, a "weak" variation of the R_ADDROFF relocation
    that will only reference the type described if it is in some other way
    reachable.
    
    Use this for the ptrToThis field in reflect type information where it
    is safe to do so (that is, types that don't need to be included for
    interface satisfaction, and types that won't cause the compiler to
    recursively generate an endless series of ptr-to-ptr-to-ptr-to...
    types).
    
    Also fix a small bug in reflect, where StructOf was not clearing the
    ptrToThis field of new types.
    
    Fixes #17931
    
    Change-Id: I4d3b53cb9c916c97b3b16e367794eee142247281
    Reviewed-on: https://go-review.googlesource.com/33427
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit aeaa4c3c1da9e7a4afd4152913d6f2bfcf4fad2d
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 22 01:43:38 2016 +0000

    net/http: skip TestLinuxSendfile on mips64 for now
    
    See issues for details. We can expand this test during the Go 1.9
    cycle.
    
    Updates #18008
    
    Change-Id: I78b6b7e8dede414769be97898e29f969bc2a9651
    Reviewed-on: https://go-review.googlesource.com/33430
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 37d078ede386b5d0dff0bb1f3ea77e47122730d0
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 10 16:52:57 2016 -0400

    math/big: add Baillie-PSW test to (*Int).ProbablyPrime
    
    After x.ProbablyPrime(n) passes the n Miller-Rabin rounds,
    add a Baillie-PSW test before declaring x probably prime.
    
    Although the provable error bounds are unchanged, the empirical
    error bounds drop dramatically: there are no known inputs
    for which Baillie-PSW gives the wrong answer. For example,
    before this CL, big.NewInt(443*1327).ProbablyPrime(1) == true.
    Now it is (correctly) false.
    
    The new Baillie-PSW test is two pieces: an added Miller-Rabin
    round with base 2, and a so-called extra strong Lucas test.
    (See the references listed in prime.go for more details.)
    The Lucas test takes about 3.5x as long as the Miller-Rabin round,
    which is close to theoretical expectations.
    
    name                              time/op
    ProbablyPrime/Lucas             2.91ms ± 2%
    ProbablyPrime/MillerRabinBase2   850µs ± 1%
    ProbablyPrime/n=0               3.75ms ± 3%
    
    The speed of prime testing for a prime input does get slower:
    
    name                  old time/op  new time/op   delta
    ProbablyPrime/n=1    849µs ± 1%   4521µs ± 1%  +432.31%   (p=0.000 n=10+9)
    ProbablyPrime/n=5   4.31ms ± 3%   7.87ms ± 1%   +82.70%  (p=0.000 n=10+10)
    ProbablyPrime/n=10  8.52ms ± 3%  12.28ms ± 1%   +44.11%  (p=0.000 n=10+10)
    ProbablyPrime/n=20  16.9ms ± 2%   21.4ms ± 2%   +26.35%   (p=0.000 n=9+10)
    
    However, because the Baillie-PSW test is only added when the old
    ProbablyPrime(n) would return true, testing composites runs at
    the same speed as before, except in the case where the result
    would have been incorrect and is now correct.
    
    In particular, the most important use of this code is for
    generating random primes in crypto/rand. That use spends
    essentially all its time testing composites, so it is not
    slowed down by the new Baillie-PSW check:
    
    name                  old time/op  new time/op   delta
    Prime                104ms ±22%    111ms ±16%      ~     (p=0.165 n=10+10)
    
    Thanks to Serhat Şevki Dinçer for CL 20170, which this CL builds on.
    
    Fixes #13229.
    
    Change-Id: Id26dde9b012c7637c85f2e96355d029b6382812a
    Reviewed-on: https://go-review.googlesource.com/30770
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 526b2f85ce8e4b1b16f3122e0a5700c04b6de199
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Nov 21 18:23:12 2016 -0500

    runtime/internal/atomic: crash on unaligned 64-bit ops on 32-bit MIPS
    
    This check was originally implemented by Vladimir in
    https://go-review.googlesource.com/c/31489/1/src/runtime/internal/atomic/atomic_mipsx.go#30
    but removed due to my comment (Sorry!). This CL adds it back.
    
    Fixes #17786.
    
    Change-Id: I7ff4c2539fc9e2afd8199964b587a8ccf093b896
    Reviewed-on: https://go-review.googlesource.com/33431
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 67ce6af4567e6edb8b246494f7faa511485666ac
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 21 22:30:14 2016 +0000

    cmd/dist: skip plugin tests on noopt builder for now
    
    Updates #17937
    
    Change-Id: Ic822da1786a983b3b7bca21b68c3d5fc4bdfaee2
    Reviewed-on: https://go-review.googlesource.com/33428
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 21a3c53c541ee6366db371e333cdb5376f507c65
Author: Russ Cox <rsc@golang.org>
Date:   Fri Nov 11 15:45:09 2016 -0500

    build: fix cross-compile on Plan 9
    
    In Plan 9's shell,
    
            GOBIN= \
                    foo bar
    
    is the same as
    
            GOBIN=foo bar
    
    Write what was meant, which is
    
            GOBIN=() \
                    foo bar
    
    Fixes #17737.
    
    Change-Id: Ie5a1b51a7cec950b5e78bbbe99cbc3cfe102f980
    Reviewed-on: https://go-review.googlesource.com/33144
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Reviewed-by: David du Colombier <0intro@gmail.com>

commit 9073af247d602dff4633710adf90c8b3c1869c45
Author: Russ Cox <rsc@golang.org>
Date:   Fri Nov 11 11:06:32 2016 -0500

    encoding/json: document what happens to MarshalText's result
    
    Fixes #17743.
    
    Change-Id: Ib5afb6248bb060f2ad8dd3d5f78e95271af62a57
    Reviewed-on: https://go-review.googlesource.com/33135
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Reviewed-by: Caleb Spare <cespare@gmail.com>

commit 6e7e8b0f0d8644b3e1981229d597cf8ac2d898c7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 21 21:29:32 2016 +0000

    cmd/go: skip slow tests on mips when run under builders
    
    Change-Id: If754de6c44cf0ec4192101432e4065cc7a28e862
    Reviewed-on: https://go-review.googlesource.com/33425
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit ff191dd7268937902e9181495765a61497fcf5ca
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 21 21:51:01 2016 +0000

    net/http: maybe fix TestLinuxSendfile on mips64
    
    Updates #18008
    
    Change-Id: I8fde0d71d15b416db4d262f6db8ef32a209a192f
    Reviewed-on: https://go-review.googlesource.com/33426
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8d226da29d95d31b6016a669b6da8f1dd52027c2
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Nov 15 16:15:05 2016 -0800

    cmd/go: don't check standard packages when using gccgo
    
    The gccgo compiler does not have the standard packages available, so it
    can not check for violations of internal references.
    
    Also, the gccgo compiler can not read runtime/internal/sys/zversion.go;
    in fact, the file does not even exist for gccgo.
    
    Change-Id: Ibadf16b371621ad1b87b6e858c5eb233913e179d
    Reviewed-on: https://go-review.googlesource.com/33295
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1368de3db23ed51992a57cc828c5f14558b4b49f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 21 19:09:42 2016 +0000

    archive/zip: skip large concurrent tests in race mode
    
    We recently added these large zip64 tests. They're slow-ish already,
    but fast enough in non-race mode with t.Parallel. But in race mode,
    the concurrency makes them much slower than the normal
    non-race-to-race multiplier.
    
    They're taking so long now that it's causing test failures when it
    sometimes is over the test timeout threshold.
    
    Change-Id: I02f4ceaa9d6cab826708eb3860f47a57b05bdfee
    Reviewed-on: https://go-review.googlesource.com/33423
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 35231ec7c6a6d9277bf6ac53cb0142e4d37c2ece
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 21 14:58:23 2016 +0000

    net/http: deflake TestClientTimeout
    
    Should fix flakes like:
    
    https://build.golang.org/log/c8da331317064227f38d5ef57ed7dba563ba1b38
    
    --- FAIL: TestClientTimeout_h1 (0.35s)
        client_test.go:1263: timeout after 200ms waiting for timeout of 100ms
    FAIL
    
    Change-Id: I0a4dba607524e8d7a00f498e27d9598acde5d222
    Reviewed-on: https://go-review.googlesource.com/33420
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 01b4ddb37724b0cd0a1f0a62956f9e0e706bb10c
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Nov 14 10:17:23 2016 -0500

    runtime/internal/atomic: crash on unaligned 64-bit ops on 386 and ARM
    
    Updates #17786. Will fix mips(32) when the port is fully landed.
    
    Change-Id: I00d4ff666ec14a38cadbcd52569b347bb5bc8b75
    Reviewed-on: https://go-review.googlesource.com/33236
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit bbe96f5673fbb379ec3da581dba0d9cd603cc0d1
Author: Cherry Zhang <cherryyz@google.com>
Date:   Mon Nov 14 18:24:37 2016 -0500

    runtime: make work.bytesMarked 8-byte aligned
    
    Make atomic access on 32-bit architectures happy.
    
    Updates #17786.
    
    Change-Id: I42de63ff1381af42124dc51befc887160f71797d
    Reviewed-on: https://go-review.googlesource.com/33235
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Austin Clements <austin@google.com>

commit ccd69d058278ef97dfd9b122d3832ec027455e90
Author: Michael Matloob <matloob@golang.org>
Date:   Fri Nov 11 15:01:58 2016 -0500

    runtime/pprof: emit count profiles with debug=0 as proto profiles
    
    count profiles with debug=1 retain their previous format.
    Also add a test check for the proto profiles since all runtime/pprof
    tests only look at the debug=1 profiles.
    
    Change-Id: Ibe805585b597e5d3570807115940a1dc4535c03f
    Reviewed-on: https://go-review.googlesource.com/33148
    Run-TryBot: Michael Matloob <matloob@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 0bae74e8c9b5fab3baf61bde0169f4aa5e287bdc
Author: Austin Clements <austin@google.com>
Date:   Sun Oct 30 20:43:53 2016 -0400

    runtime: wake idle Ps when enqueuing GC work
    
    If the scheduler has no user work and there's no GC work visible, it
    puts the P to sleep (or blocks on the network). However, if we later
    enqueue more GC work, there's currently nothing that specifically
    wakes up the scheduler to let it start an idle GC worker. As a result,
    we can underutilize the CPU during GC if Ps have been put to sleep.
    
    Fix this by making GC wake idle Ps when work buffers are put on the
    full list. We already have a hook to do this, since we use this to
    preempt a random P if we need more dedicated workers. We expand this
    hook to instead wake an idle P if there is one. The logic we use for
    this is identical to the logic used to wake an idle P when we ready a
    goroutine.
    
    To make this really sound, we also fix the scheduler to re-check the
    idle GC worker condition after releasing its P. This closes a race
    where 1) the scheduler checks for idle work and finds none, 2) new
    work is enqueued but there are no idle Ps so none are woken, and 3)
    the scheduler releases its P.
    
    There is one subtlety here. Currently we call enlistWorker directly
    from putfull, but the gcWork is in an inconsistent state in the places
    that call putfull. This isn't a problem right now because nothing that
    enlistWorker does touches the gcWork, but with the added call to
    wakep, it's possible to get a recursive call into the gcWork
    (specifically, while write barriers are disallowed, this can do an
    allocation, which can dispose a gcWork, which can put a workbuf). To
    handle this, we lift the enlistWorker calls up a layer and delay them
    until the gcWork is in a consistent state.
    
    Fixes #14179.
    
    Change-Id: Ia2467a52e54c9688c3c1752e1fc00f5b37bbfeeb
    Reviewed-on: https://go-review.googlesource.com/32434
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 49ea9207b6512c2400de11bc097d974bb527ba63
Author: Austin Clements <austin@google.com>
Date:   Sun Oct 30 20:20:17 2016 -0400

    runtime: exit idle worker if there's higher-priority work
    
    Idle GC workers trigger whenever there's a GC running and the
    scheduler doesn't find any other work. However, they currently run for
    a full scheduler quantum (~10ms) once started.
    
    This is really bad for event-driven applications, where work may come
    in on the network hundreds of times during that window. In the
    go-gcbench rpc benchmark, this is bad enough to often cause effective
    STWs where all Ps are in the idle worker. When this happens, we don't
    even poll the network any more (except for the background 10ms poll in
    sysmon), so we don't even know there's more work to do.
    
    Fix this by making idle workers check with the scheduler roughly every
    100 µs to see if there's any higher-priority work the P should be
    doing. This check includes polling the network for incoming work.
    
    Fixes #16528.
    
    Change-Id: I6f62ebf6d36a92368da9891bafbbfd609b9bd003
    Reviewed-on: https://go-review.googlesource.com/32433
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 7dc97d9e328edc800e2ce41d5b211ef4e0ef41d0
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Nov 18 18:14:11 2016 -0800

    misc/cgo/testcshared: add explicit ./ to shared library argument
    
    Use an explicit ./ to make sure we link against the libgo.so we just
    built, not some other libgo.so that the compiler or linker may decide to
    seek out.
    
    Fixes #17986.
    
    Change-Id: Id23f6c95aa2b52f4f42c1b6dac45482c22b4290d
    Reviewed-on: https://go-review.googlesource.com/33413
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f42929ce9fcf8b32656900764881cfb84fdbe46b
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 18 14:47:51 2016 -0800

    go/internal/gccgoimporter: handle conversions in exported const values
    
    Also: handle version "v2" of export data format.
    
    Fixes #17981.
    
    Change-Id: I8042ce18c4a27c70cc1ede675daca019b047bcf3
    Reviewed-on: https://go-review.googlesource.com/33412
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f39050c8ebf894ccedc0b99de96f7412be97af89
Author: Keith Randall <khr@golang.org>
Date:   Mon Jun 6 16:58:27 2016 -0700

    cmd/cover: handle multiple samples from the same location
    
    So we can merge cover profiles from multiple runs.
    
    Change-Id: I1bf921e2b02063a2a62b35d21a6823062d10e5d0
    Reviewed-on: https://go-review.googlesource.com/23831
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit b01f612a693b2b39064c0b5bd75a5d0280e4179e
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 18 10:49:29 2016 -0800

    spec: add subtitles to section on "for" statements
    
    This matches what we already do for switch statements and makes
    this large section more visibly organized. No other changes besides
    introducing the titles.
    
    Fixes #4486.
    
    Change-Id: I73f274e4fdd27c6cfeaed79090b4553e57a9c479
    Reviewed-on: https://go-review.googlesource.com/33410
    Reviewed-by: Rob Pike <r@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d0b3c169acda68040d051c27627c08da4e3377bd
Author: Austin Clements <austin@google.com>
Date:   Fri Nov 4 11:13:27 2016 -0400

    cmd/trace: fix goroutine view
    
    Currently, trace processing interleaves state/statistics updates and
    emitting trace viewer objects. As a result, if events are being
    filtered, either by time or by goroutines, we'll miss those
    state/statistics updates. At best, this leads to bad statistics;
    however, since we're now strictly checking G state transitions, it
    usually leads to a failure to process the trace if there is any
    filtering.
    
    Fix this by separating state updates from emitting trace object. State
    updates are done before filtering, so we always have correct state
    information and statistics. Trace objects are only emitted if we pass
    the filter. To determine when we need to emit trace counters, rather
    than duplicating the knowledge of which events might modify
    statistics, we keep track of the previously emitted counters and emit
    a trace counter object whenever these have changed.
    
    Fixes #17719.
    
    Change-Id: Ic66e3ddaef60d1acaaf2ff4c62baa5352799cf99
    Reviewed-on: https://go-review.googlesource.com/32810
    Reviewed-by: Dmitry Vyukov <dvyukov@google.com>

commit 0eb26fa8ba531b21d183fd3a4d3fb8abf57db7aa
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 18 09:27:18 2016 -0800

    spec: remove => (alias) operator from Operators and Delimiters section
    
    (Revert of https://go-review.googlesource.com/#/c/32310/)
    
    For #16339.
    Fixes #17975.
    
    Change-Id: I36062703c423a81ea1c5b00f4429a4faf00b3782
    Reviewed-on: https://go-review.googlesource.com/33365
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 120cf676caff29296de2dd16a2997463eb6e1579
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Nov 17 16:39:11 2016 -0800

    spec: clarify type elision rules for composite literals
    
    - organize examples better
    - add an example illustrating behavior if element type is a named pointer type
    - both compilers and go/types (per https://go-review.googlesource.com/33358)
      follow this now
    
    See the issue for detailed discussion.
    
    Fixes #17954.
    
    Change-Id: I8d90507ff2347d9493813f75b73233819880d2b4
    Reviewed-on: https://go-review.googlesource.com/33361
    Reviewed-by: Rob Pike <r@golang.org>

commit a34fddf46c47a86e7a7cab32be858f7e8d0feb70
Author: Philip Hofer <phofer@umich.edu>
Date:   Mon Nov 14 17:05:46 2016 -0800

    cmd/compile: in cse, allow for new ssa values
    
    The table of rewrites in ssa/cse is not sized appropriately for
    ssa IDs that are created during copying of selects into new blocks.
    
    Fixes #17918
    
    Change-Id: I65fe86c6aab5efa679aa473aadc4ee6ea882cd41
    Reviewed-on: https://go-review.googlesource.com/33240
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 277bcbbdcd26f2d64493e596238e34b47782f98e
Author: Özgür Kesim <oec-go@kesim.org>
Date:   Fri Oct 21 13:14:57 2016 +0200

    text/template: handle option missingkey=error consistently
    
    The existing implementation of text/template handles the option
    "missingkey=error" in an inconsitent manner:  If the provided data is
    a nil-interface, no error is returned (despite the fact that no key
    can be found in it).
    
    This patch makes text/template return an error if "missingkey=error"
    is set and the provided data is a not a valid reflect.Value.
    
    Fixes #15356
    
    Change-Id: Ia0a83da48652ecfaf31f18bdbd78cb21dbca1164
    Reviewed-on: https://go-review.googlesource.com/31638
    Reviewed-by: Rob Pike <r@golang.org>
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 348275cda626edf60e1d11a83b2e78d83088ccef
Author: Cherry Zhang <cherryyz@google.com>
Date:   Fri Oct 28 23:11:04 2016 -0400

    cmd/compile: make a copy of Phi input if it is still live
    
    Register of Phi input is allocated to the Phi. So if the Phi
    input is still live after Phi, we may need to use a spill. In
    this case, copy the Phi input to a spare register to avoid a
    spill.
    
    Originally targeted the code in issue #16187, and this CL
    indeed removes the spill, but doesn't seem to help on benchmark
    result. It may help in general, though.
    
    On AMD64:
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.79s ± 1%     2.76s ± 0%  -1.33%  (p=0.000 n=10+10)
    Fannkuch11-12                3.02s ± 0%     3.14s ± 0%  +3.99%  (p=0.000 n=10+10)
    FmtFprintfEmpty-12          51.2ns ± 0%    51.4ns ± 3%    ~      (p=0.368 n=8+10)
    FmtFprintfString-12          145ns ± 0%     144ns ± 0%  -0.69%    (p=0.000 n=6+9)
    FmtFprintfInt-12             127ns ± 1%     124ns ± 1%  -2.79%   (p=0.000 n=10+9)
    FmtFprintfIntInt-12          186ns ± 0%     184ns ± 0%  -1.34%   (p=0.000 n=10+9)
    FmtFprintfPrefixedInt-12     196ns ± 0%     194ns ± 0%  -0.97%    (p=0.000 n=9+9)
    FmtFprintfFloat-12           293ns ± 2%     287ns ± 0%  -2.00%   (p=0.000 n=10+9)
    FmtManyArgs-12               847ns ± 1%     829ns ± 0%  -2.17%   (p=0.000 n=10+7)
    GobDecode-12                7.17ms ± 0%    7.18ms ± 0%    ~     (p=0.123 n=10+10)
    GobEncode-12                6.08ms ± 1%    6.08ms ± 0%    ~      (p=0.497 n=10+9)
    Gzip-12                      277ms ± 1%     275ms ± 1%  -0.47%   (p=0.028 n=10+9)
    Gunzip-12                   39.1ms ± 2%    38.2ms ± 1%  -2.20%   (p=0.000 n=10+9)
    HTTPClientServer-12         90.9µs ± 4%    87.7µs ± 2%  -3.51%   (p=0.001 n=9+10)
    JSONEncode-12               17.3ms ± 1%    16.5ms ± 0%  -5.02%    (p=0.000 n=9+9)
    JSONDecode-12               54.6ms ± 1%    54.1ms ± 0%  -0.99%    (p=0.000 n=9+9)
    Mandelbrot200-12            4.45ms ± 0%    4.45ms ± 0%  -0.02%    (p=0.006 n=8+9)
    GoParse-12                  3.44ms ± 0%    3.48ms ± 1%  +0.95%  (p=0.000 n=10+10)
    RegexpMatchEasy0_32-12      84.9ns ± 0%    85.0ns ± 0%    ~       (p=0.241 n=8+8)
    RegexpMatchEasy0_1K-12       867ns ± 3%     915ns ±11%  +5.55%  (p=0.037 n=10+10)
    RegexpMatchEasy1_32-12      82.7ns ± 5%    83.9ns ± 4%    ~      (p=0.161 n=9+10)
    RegexpMatchEasy1_1K-12       361ns ± 1%     363ns ± 0%    ~      (p=0.098 n=10+8)
    RegexpMatchMedium_32-12      126ns ± 0%     126ns ± 1%    ~      (p=0.549 n=8+10)
    RegexpMatchMedium_1K-12     38.8µs ± 0%    39.1µs ± 0%  +0.67%    (p=0.000 n=9+8)
    RegexpMatchHard_32-12       1.95µs ± 0%    1.96µs ± 0%  +0.43%    (p=0.000 n=9+9)
    RegexpMatchHard_1K-12       59.0µs ± 0%    59.1µs ± 0%  +0.27%   (p=0.000 n=10+9)
    Revcomp-12                   436ms ± 1%     431ms ± 1%  -1.19%  (p=0.005 n=10+10)
    Template-12                 56.7ms ± 1%    57.1ms ± 1%  +0.71%   (p=0.001 n=10+9)
    TimeParse-12                 312ns ± 0%     310ns ± 0%  -0.80%   (p=0.000 n=10+9)
    TimeFormat-12                336ns ± 0%     332ns ± 0%  -1.19%    (p=0.000 n=8+7)
    [Geo mean]                  59.2µs         58.9µs       -0.42%
    
    On PPC64:
    name                     old time/op    new time/op    delta
    BinaryTree17-2              4.67s ± 2%     4.71s ± 1%    ~     (p=0.421 n=5+5)
    Fannkuch11-2                3.92s ± 1%     3.94s ± 0%  +0.46%  (p=0.032 n=5+5)
    FmtFprintfEmpty-2           122ns ± 0%     120ns ± 2%  -1.80%  (p=0.016 n=4+5)
    FmtFprintfString-2          305ns ± 1%     299ns ± 1%  -1.84%  (p=0.008 n=5+5)
    FmtFprintfInt-2             243ns ± 0%     241ns ± 1%  -0.66%  (p=0.016 n=4+5)
    FmtFprintfIntInt-2          361ns ± 1%     356ns ± 1%  -1.49%  (p=0.016 n=5+5)
    FmtFprintfPrefixedInt-2     355ns ± 1%     357ns ± 1%    ~     (p=0.333 n=5+5)
    FmtFprintfFloat-2           502ns ± 2%     498ns ± 1%    ~     (p=0.151 n=5+5)
    FmtManyArgs-2              1.55µs ± 2%    1.59µs ± 1%  +2.52%  (p=0.008 n=5+5)
    GobDecode-2                13.0ms ± 1%    13.0ms ± 1%    ~     (p=0.841 n=5+5)
    GobEncode-2                11.8ms ± 1%    11.8ms ± 1%    ~     (p=0.690 n=5+5)
    Gzip-2                      499ms ± 1%     503ms ± 0%    ~     (p=0.421 n=5+5)
    Gunzip-2                   86.5ms ± 0%    86.4ms ± 1%    ~     (p=0.841 n=5+5)
    HTTPClientServer-2         68.2µs ± 2%    69.6µs ± 3%    ~     (p=0.151 n=5+5)
    JSONEncode-2               39.0ms ± 1%    37.2ms ± 1%  -4.65%  (p=0.008 n=5+5)
    JSONDecode-2                122ms ± 1%     126ms ± 1%  +2.63%  (p=0.008 n=5+5)
    Mandelbrot200-2            6.08ms ± 1%    5.89ms ± 1%  -3.06%  (p=0.008 n=5+5)
    GoParse-2                  5.95ms ± 2%    5.98ms ± 1%    ~     (p=0.421 n=5+5)
    RegexpMatchEasy0_32-2       331ns ± 1%     328ns ± 1%    ~     (p=0.056 n=5+5)
    RegexpMatchEasy0_1K-2      1.45µs ± 0%    1.47µs ± 0%  +1.13%  (p=0.008 n=5+5)
    RegexpMatchEasy1_32-2       359ns ± 0%     353ns ± 0%  -1.84%  (p=0.008 n=5+5)
    RegexpMatchEasy1_1K-2      1.79µs ± 0%    1.81µs ± 1%  +1.16%  (p=0.008 n=5+5)
    RegexpMatchMedium_32-2      420ns ± 2%     413ns ± 0%  -1.72%  (p=0.008 n=5+5)
    RegexpMatchMedium_1K-2     70.2µs ± 1%    69.5µs ± 1%  -1.09%  (p=0.032 n=5+5)
    RegexpMatchHard_32-2       3.87µs ± 1%    3.65µs ± 0%  -5.86%  (p=0.008 n=5+5)
    RegexpMatchHard_1K-2        111µs ± 0%     105µs ± 0%  -5.49%  (p=0.016 n=5+4)
    Revcomp-2                   1.00s ± 1%     1.01s ± 2%    ~     (p=0.151 n=5+5)
    Template-2                  113ms ± 1%     113ms ± 2%    ~     (p=0.841 n=5+5)
    TimeParse-2                 555ns ± 0%     550ns ± 1%  -0.87%  (p=0.032 n=5+5)
    TimeFormat-2                736ns ± 1%     704ns ± 1%  -4.35%  (p=0.008 n=5+5)
    [Geo mean]                  120µs          119µs       -0.77%
    
    Reduce "spilled value remains" by 0.6% in cmd/go on AMD64.
    
    Change-Id: If655df343b0f30d1a49ab1ab644f10c698b96f3e
    Reviewed-on: https://go-review.googlesource.com/32442
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit d24b57a6a1a3530e590b7c0a72dc78043e198630
Author: Elias Naur <elias.naur@gmail.com>
Date:   Sun Nov 6 21:40:57 2016 +0100

    runtime: handle SIGPIPE in c-archive and c-shared programs
    
    Before this CL, Go programs in c-archive or c-shared buildmodes
    would not handle SIGPIPE. That leads to surprising behaviour where
    writes on a closed pipe or socket would raise SIGPIPE and terminate
    the program. This CL changes the Go runtime to handle
    SIGPIPE regardless of buildmode. In addition, SIGPIPE from non-Go
    code is forwarded.
    
    Fixes #17393
    Updates #16760
    
    Change-Id: I155e82020a03a5cdc627a147c27da395662c3fe8
    Reviewed-on: https://go-review.googlesource.com/32796
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e54662dc859f65f90eefede76a6134f8d892cc77
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Nov 17 13:35:59 2016 -0800

    go/types: look at underlying type of element type of composite literals with elided types
    
    Match behavior of gc and gccgo.
    
    For #17954.
    
    Change-Id: I3f065e56d0a623bd7642c1438d0cab94d23fa2ae
    Reviewed-on: https://go-review.googlesource.com/33358
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit b21743c6d0484a6912d1a4ee20905c7b9b578ed8
Author: Adam Langley <agl@golang.org>
Date:   Thu Nov 17 12:15:19 2016 -0800

    crypto/tls: reject zero-length SCTs.
    
    The SignedCertificateTimestampList[1] specifies that both the list and
    each element must not be empty. Checking that the list is not empty was
    handled in [2] and this change checks that the SCTs themselves are not
    zero-length.
    
    [1] https://tools.ietf.org/html/rfc6962#section-3.3
    [2] https://golang.org/cl/33265
    
    Change-Id: Iabaae7a15f6d111eb079e5086e0bd2005fae9e48
    Reviewed-on: https://go-review.googlesource.com/33355
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c09945980a80a3b92362bd2e99a883051d2dd4d7
Author: woodsaj <awoods@raintank.io>
Date:   Thu Nov 17 20:14:32 2016 +0800

    crypto/tls: reject CT extension with no SCTs included
    
    When the CT extension is enabled but no SCTs are present, the existing
    code calls "continue" which causes resizing the data byte slice to be
    skipped. In fact, such extensions should be rejected.
    
    Fixes #17958
    
    Change-Id: Iad12da10d1ea72d04ae2e1012c28bb2636f06bcd
    Reviewed-on: https://go-review.googlesource.com/33265
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5cd6ab5b6d8232d5443f474c378f1307ce502613
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Thu Nov 17 20:06:27 2016 +0100

    runtime/pprof/internal/protopprof: fix TestTranslateCPUProfileWithSamples test for mips
    
    Change-Id: I01168a7530e18dd1098d467d0c8a330f727ba91f
    Reviewed-on: https://go-review.googlesource.com/33281
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 05dc6b26ca337b80f9424aa9b78f4b45f6e44c8d
Author: Austin Clements <austin@google.com>
Date:   Thu Nov 17 10:48:40 2016 -0500

    runtime: improve diagnostics for "scan missed a g"
    
    Currently there are no diagnostics for mark root check during marking.
    Fix this by printing out the same diagnostics we print during mark
    termination.
    
    Also, drop the allglock before throwing. Holding that across a throw
    causes a self-deadlock with tracebackothers.
    
    For #16083.
    
    Change-Id: Ib605f3ae0c17e70704b31d8378274cfaa2307dc2
    Reviewed-on: https://go-review.googlesource.com/33339
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 7061dc3f6e94dbc369db779ec6799c6f60d3466f
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Nov 16 09:55:24 2016 -0800

    cmd/cgo: ignore top-level qualifiers in function args/results
    
    The top-level qualifiers are unimportant for our purposes. If a C
    function is defined as `const int f(const int i)`, the `const`s are
    meaningless to C, and we want to avoid using them in the struct we
    create where the `const` has a completely different meaning.
    
    This unwinds https://golang.org/cl/33097 with regard to top-level
    qualifiers.
    
    Change-Id: I3d66b0eb43b6d9a586d9cdedfae5a2306b46d96c
    Reviewed-on: https://go-review.googlesource.com/33325
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit c1e9760d4ca102539b2d52a7c4021205c29070bf
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 15 11:33:10 2016 -0800

    archive/zip: avoid overflow in record count and byte offset fields
    
    This is Quentin's https://golang.org/cl/33012 with updated tests.
    
    Fixes #14186
    
    Change-Id: Ib51deaab0368c6bad32ce9d6345119ff44f3c2d6
    Reviewed-on: https://go-review.googlesource.com/33291
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 90b8a0ca2d0b565c7c7199ffcf77b15ea6b6db3a
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Thu Nov 17 09:33:31 2016 -0800

    database/sql: ensure all driver Stmt are closed once
    
    Previously  driver.Stmt could could be closed multiple times in
    edge cases that drivers may not test for initially. Make their
    job easier by ensuring the driver is only closed a single time.
    
    Fixes #16019
    
    Change-Id: I1e4777ef70697a849602e6ef9da73054a8feb4cd
    Reviewed-on: https://go-review.googlesource.com/33352
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e0942b76c735a69df18d24a23fa16da1e5db8c2e
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Nov 17 12:26:36 2016 -0500

    cmd/asm/internal/asm: fix copy/paste errors in comment
    
    Change-Id: I0249b60e340710bea7b6671c9b7405c278b037bd
    Reviewed-on: https://go-review.googlesource.com/33351
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit afb0ae67b7d37b6678179f236c98ef0d952c0403
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 17 16:28:45 2016 +0000

    runtime/pprof: fix typo in test
    
    Not sure what I was thinking.
    
    Change-Id: I143cdf7c5ef8e7b2394afeca6b30c46bb2c19a55
    Reviewed-on: https://go-review.googlesource.com/33340
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4ca3a8f7a807bba55e9db82b9aa8c43b1a186b8e
Author: Elias Naur <elias.naur@gmail.com>
Date:   Thu Nov 17 13:52:00 2016 +0100

    misc/cgo: decrease test failure timeouts
    
    CL 33239 changed the polling loops from using sched_yield to a sleep
    for 1/1000 of a second. The loop counters were not updated, so failing
    tests now take 100 seconds to complete. Lower the loop counts to 5
    seconds instead.
    
    Change-Id: I7c9a343dacc8188603ecf7e58bd00b535cfc87f5
    Reviewed-on: https://go-review.googlesource.com/33280
    Run-TryBot: Elias Naur <elias.naur@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit fd0f69c6802b92fed953659efe5f2e0e2e8aed14
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Nov 17 15:41:18 2016 +0900

    net: use testenv.SkipFlaky instead of testing.Skip
    
    Change-Id: Ic219fedbe6bbb846f31111fa21df6f2b8620e269
    Reviewed-on: https://go-review.googlesource.com/33263
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7534a72ea8b4ebb71cd8525029717fddd46c9dc6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 17 14:39:11 2016 +0000

    fmt: fix typo
    
    Fixes #17955
    
    Change-Id: Ia1a04796353c83358a38a6b63f2a0cd3c6926f09
    Reviewed-on: https://go-review.googlesource.com/33338
    Reviewed-by: Rob Pike <r@golang.org>

commit 03ca047dd334f6018f06f7fc9a7a4e2608b1f8d3
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sat Nov 12 17:57:35 2016 +1100

    debug/pe: do not create symbol table if FileHeader.PointerToSymbolTable is 0
    
    https://github.com/tpn/pdfs/raw/master/Microsoft Portable Executable and Common Object File Format Specification - 1999 (pecoff).doc
    says this about PointerToSymbolTable:
    
    File offset of the COFF symbol table or 0 if none is present.
    
    Do as it says.
    
    Fixes #17809.
    
    Change-Id: Ib1ad83532f36a3e56c7e058dc9b2acfbf60c4e72
    Reviewed-on: https://go-review.googlesource.com/33170
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit dadfd14babccc30757ddb3f3eb8fbb7cd3bf4b5a
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Thu Nov 10 13:00:35 2016 +1100

    os: add more tests in TestReadStdin
    
    TestReadStdin always fill up buffer provided by ReadFile caller full.
    But we do not know if real ReadFile does the same. Add tests where
    buffer is only filled with limited data.
    
    Change-Id: I0fc776325c2b1fe60511126c439f4b0560e9d653
    Reviewed-on: https://go-review.googlesource.com/33030
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit b8d56fdd9379bd5bb89454f3679f01acecd9df4d
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Nov 17 12:18:39 2016 +0900

    net: tweak comment on ExampleCIDRMask
    
    CIDRMask just returns a mask which corresponds to an address
    prefix in CIDR nonation. A subnet for an IPv6 mask sounds a bit
    confusing.
    
    Change-Id: Ic7859ce992bc2de4043d3b25caf9a1051d118b0e
    Reviewed-on: https://go-review.googlesource.com/33262
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b2d34fa51bd509f0aa780151a3d30c5ca77f1f4e
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Wed Nov 9 09:09:37 2016 -0600

    runtime: handle bad ftab index in symtab.go
    
    If a program has had its text section split into multiple
    sections then the ftab that is built is based on addresses
    prior to splitting.  That means all the function addresses
    are there and correct because of relocation but the
    but the computed idx won't always match up quite right and
    in some cases go beyond the end of the table, causing a panic.
    
    To resolve this, determine if the idx is too large and if it is,
    set it to the last index in ftab.  Then search backward to find the
    matching function address.
    
    Fixes #17854
    
    Change-Id: I6940e76a5238727b0a9ac23dc80000996db2579a
    Reviewed-on: https://go-review.googlesource.com/32972
    Reviewed-by: David Chase <drchase@google.com>

commit a1235f3179c4dbd6b16963d6b8f932586fa9bc1c
Author: Joonas Kuorilehto <joneskoo@derbian.fi>
Date:   Sun Sep 11 22:31:19 2016 +0300

    crypto/tls: add example for Config KeyLogWriter
    
    For #13057.
    
    Change-Id: Idbc50d5b08e055a23ab7cc9eb62dbc47b65b1815
    Reviewed-on: https://go-review.googlesource.com/29050
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 011cb6423187699e26553580d97518f40966d32b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Nov 16 16:13:22 2016 -0800

    cmd/compile, reflect: use field pkgPath if needed
    
    It's possible for the pkgPath of a field to be different than that of
    the struct type as a whole. In that case, store the field's pkgPath in
    the name field. Use the field's pkgPath when setting PkgPath and when
    checking for type identity.
    
    Fixes #17952.
    
    Change-Id: Iebaf92f0054b11427c8f6e4158c3bebcfff06f45
    Reviewed-on: https://go-review.googlesource.com/33333
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit fbf92436b95d91151ce6717f40c46614ee68d487
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Nov 16 11:23:53 2016 -0800

    doc: add FAQ: why no conversion from []T1 to []T2?
    
    Fixes #16934.
    
    Change-Id: I725704e4c4aae7023fd89edc42af7ba0d242fec8
    Reviewed-on: https://go-review.googlesource.com/33327
    Reviewed-by: Rob Pike <r@golang.org>

commit 48858a2386f7eef41cb3459e85d53ac9b1e8f70f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 16 23:41:15 2016 +0000

    net/http: deflake TestInterruptWithPanic_nil_h2, again
    
    Updates #17243
    
    Change-Id: Iaa737874e75fdac73452f1fc13a5749e8df78ebe
    Reviewed-on: https://go-review.googlesource.com/33332
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 14e9f4825bffd4339dbde43198ed1710a1e149b5
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Wed Nov 16 22:39:25 2016 +0000

    cmd/cover: don't ignore os.Create error
    
    Failing to create the output file would give confusing errors such as:
    
            cover: invalid argument
    
    Also do out.Close() even if Execute() errored.
    
    Fixes #17951.
    
    Change-Id: I897e1d31f7996871c54fde7cb09614cafbf6c3fc
    Reviewed-on: https://go-review.googlesource.com/33278
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit d7c0de98a96893e5608358f7578c85be7ba12b25
Author: Scott Bell <scott@sctsm.com>
Date:   Wed May 18 09:56:51 2016 -0700

    database/sql: additional underlying types in DefaultValueConverter
    
    The previous documentation purported to convert underlying strings to
    []byte, which it did not do. This adds support for underlying bool,
    string, and []byte, which convert directly to their underlying type.
    
    Fixes #15174.
    
    Change-Id: I7fc4e2520577f097a48f39c9ff6c8160fdfb7be4
    Reviewed-on: https://go-review.googlesource.com/27812
    Reviewed-by: Daniel Theophanes <kardianos@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit 0df762ed7b2ac9484a4f64450424529ecee8aa7f
Author: Kevin Burke <kev@inburke.com>
Date:   Thu Nov 3 15:34:05 2016 -0700

    net: add example for CIDRMask
    
    I had trouble translating the documentation language into a subnet
    - e.g. whether /31 was CIDRMask(1, 31) or CIDRMask(1, 32) or
    CIDRMask(31, 32) so I thought I'd add a short example showing how to
    create the right masks.
    
    Change-Id: Ia6a6de08c5c30b6d2249b3194cced2d3c383e317
    Reviewed-on: https://go-review.googlesource.com/32677
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fe057c1478c4309f715d689127125cabbb4efe08
Author: Ian Lance Taylor <iant@golang.org>
Date:   Wed Nov 16 13:51:32 2016 -0800

    runtime/cgo: fixes for calling sigaction in C
    
    Zero out the sigaction structs, in case the sa_restorer field is set.
    
    Clear the SA_RESTORER flag; it is part of the kernel interface, not the
    libc interface.
    
    Fixes #17947.
    
    Change-Id: I610348ce3c196d3761cf2170f06c24ecc3507cf7
    Reviewed-on: https://go-review.googlesource.com/33331
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Bryan Mills <bcmills@google.com>

commit 8dc47e3b3a939a39e5cc3ea59f4848f50fd0cb7b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 16 20:35:02 2016 +0000

    net: disable TestAcceptTimeout for now
    
    It's too flaky and doing more harm than good.
    
    Disable it until it can be made reliable.
    
    Updates #17948
    Updates #17927
    
    Change-Id: Iaab7f09a4060da377fcd3ca2262527fef50c558d
    Reviewed-on: https://go-review.googlesource.com/33330
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 49b77a87974c310b1d6a8437f3490f81811b2058
Author: Daniel Theophanes <kardianos@gmail.com>
Date:   Wed Nov 16 11:33:38 2016 -0800

    database/sql: guard against driver.Stmt.Close panics
    
    Do not retain a lock when driver.Stmt.Close panic as the rest
    of the sql package ensures.
    
    Updates #16019
    
    Change-Id: Idc7ea9258ae23f491e79cce3efc365684a708428
    Reviewed-on: https://go-review.googlesource.com/33328
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 81627f0e47f91a6e6e7bfd7c59b4e3ac596668ca
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Nov 17 06:06:34 2016 +0900

    net: deflake TestAcceptTimeout again
    
    This is a followup to CL 33257.
    
    It looks like active close operation at passive open side sometimes
    takes a bit long time on Darwin.
    
    Fixes #17948.
    
    Change-Id: Ida17639c4e66a43e1be1f74fd0ef3baddde25092
    Reviewed-on: https://go-review.googlesource.com/33258
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9f5673d9307572ff3b435a845470e3b7fd3c6a43
Author: David Chase <drchase@google.com>
Date:   Mon Nov 14 18:00:17 2016 -0500

    cmd/compile: ensure necessary types appear in .debug_info
    
    Autotmp filtering was too aggressive and excluded types
    necessary to make debuggers work properly.  Restore the
    "late filter" in dwarf.go based on names to exclude autotmps,
    and remove the "early filter" in pgen.go based on how the
    name was introduced.  However, the updated naming scheme
    with a dot prefix is retained to prevent accidental clashes
    with legal Go identifier names.
    
    Includes test (grouped with runtime gdb tests),
    verified to fail without the fix.
    
    Updates #17644.
    Fixes #17830.
    
    Change-Id: I7ec3f7230083889660236e5f6bc77ba5fe434e93
    Reviewed-on: https://go-review.googlesource.com/33233
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 1e3c57c2cc1500b12a35a859f3d6e8aa27aeebc5
Author: Cherry Zhang <cherryyz@google.com>
Date:   Wed Nov 16 10:43:54 2016 -0500

    cmd/internal/obj/arm64: fix branch too far for CBZ (and like)
    
    The assembler backend fixes too-far conditional branches, but only
    for BEQ and like. Add a case for CBZ and like.
    
    Fixes #17925.
    
    Change-Id: Ie516e6c5ca165b582367283a0110f7081e00c214
    Reviewed-on: https://go-review.googlesource.com/33304
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: David Chase <drchase@google.com>

commit cd66c38619cbf6e031d9af4cea8197cd6980ffa1
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 16 16:27:27 2016 +0000

    runtime/pprof: skip profiling tests on mips if highres timers not available
    
    Fixes #17936
    
    Change-Id: I20d09712b7d7303257994356904052ba64bc5bf2
    Reviewed-on: https://go-review.googlesource.com/33306
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit e279280d0db0a8ffb7b453e789a2e322722d2259
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Wed Nov 16 16:34:02 2016 +0900

    net: deflake TestAcceptTimeout
    
    This change makes use of synchronization primitive instead of
    context-based canceling not to depend on defer execution scheduling.
    
    Fixes #17927.
    
    Change-Id: I5ca9287a48bb5cdda6845a7f12757f95175c5db8
    Reviewed-on: https://go-review.googlesource.com/33257
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d8b14c524303f8d28bc5b496e918cfbb2758cbc5
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 16 19:08:25 2016 +0000

    math/rand: make floating point tests shorter on mips and mipsle
    
    Like GOARM=5 does.
    
    Fixes #17944
    
    Change-Id: Ica2a54a90fbd4a29471d1c6009ace2fcc5e82a73
    Reviewed-on: https://go-review.googlesource.com/33326
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 68fda1888e8026aef96590ac634bc35d4c71b6e0
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Wed Nov 16 15:26:41 2016 +0000

    all: call flag.Parse from TestMain only if used
    
    These don't use any flags in TestMain itself, so the call is redundant
    as M.Run will do it.
    
    Change-Id: I00f2ac7f846dc2c3ad3535eb8177616b2d900149
    Reviewed-on: https://go-review.googlesource.com/33275
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7ee793652307269c9fdee2c0cb222509371a6e36
Author: David Crawshaw <crawshaw@golang.org>
Date:   Sun Nov 13 21:20:58 2016 -0500

    cmd/link: handle R_GOTPCREL separately on darwin
    
    To generate the correct section offset the shared code path for
    R_CALL, R_PCREL, and R_GOTPCREL on darwin when externally linking
    walks up the symbol heirarchy adding the differences. This is fine,
    except in the case where we are generating a GOT lookup, because
    the topmost symbol is left in r.Xsym instead of the symbol we are
    looking up. So all funcsym GOT lookups were looking up the outer
    "go.func.*" symbol.
    
    Fix this by separating out the R_GOTPCREL code path.
    
    For #17828 (and may fix it).
    
    Change-Id: I2c9f4d135e77c17270aa064d8c876dc6d485d659
    Reviewed-on: https://go-review.googlesource.com/33211
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b75b9e1d65989753d0ee14ccc6007729e49e2e51
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Nov 15 17:14:13 2016 -0800

    database/sql: clarify when statement in transaction is closed
    
    Fixes #16346.
    
    Change-Id: Ie75a4ae7011036dd2c1f121a7a5e38d10177721e
    Reviewed-on: https://go-review.googlesource.com/33296
    Reviewed-by: Daniel Theophanes <kardianos@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 26069e1981ba8500baf35ff5aec79531e4882543
Author: Daniel Martí <mvdan@mvdan.cc>
Date:   Wed Nov 16 14:28:12 2016 +0000

    cmd/compile: remove some unused code
    
    The use of these has been removed in recent commits.
    
    Change-Id: Iff36a3ee4dcdfe39c40e93e2601f44d3c59f7024
    Reviewed-on: https://go-review.googlesource.com/33274
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
```
