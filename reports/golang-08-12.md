# August 12, 2016 Report

Number of commits: 10

## Compilation time

* github.com/coreos/etcd/cmd: from 13.449608827s to 13.415211207s, -0.26%
* github.com/boltdb/bolt/cmd/bolt: from 723.857958ms to 615.353243ms, -14.99%
* github.com/gogits/gogs: from 13.979710365s to 13.798540772s, -1.30%
* github.com/spf13/hugo: from 9.32488364s to 7.659044634s, -17.86%
* github.com/nsqio/nsq/apps/nsqd: from 2.047527696s to 2.2435499s, +9.57%
* github.com/mholt/caddy: from 285.239281ms to 299.829287ms, +5.12%

## Binary size:

* github.com/coreos/etcd/cmd: from 26680798 to 26685015, ~
* github.com/boltdb/bolt/cmd/bolt: from 2679371 to 2679372, ~
* github.com/gogits/gogs: from 23918579 to 23930988, ~
* github.com/spf13/hugo: from 15877527 to 15885840, ~
* github.com/nsqio/nsq/apps/nsqd: from 10008067 to 10012284, ~
* github.com/mholt/caddy: from 13044558 to 13044558, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               187           189           +1.07%
BenchmarkMsgpUnmarshal-4             405           403           -0.49%
BenchmarkEasyJsonMarshal-4           1598          1604          +0.38%
BenchmarkEasyJsonUnmarshal-4         2245          2265          +0.89%
BenchmarkFlatBuffersMarshal-4        387           361           -6.72%
BenchmarkFlatBuffersUnmarshal-4      290           291           +0.34%
BenchmarkGogoprotobufMarshal-4       162           161           -0.62%
BenchmarkGogoprotobufUnmarshal-4     256           255           -0.39%

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
commit 31ad583ab24263f9dbcb5cbcce849eed64e74040
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Jun 25 15:11:25 2016 -0700

    testing: respect benchtime on very fast benchmarks
    
    When ns/op dropped below 1, the old code
    ignored benchtime and reverted to 1s.
    
    Change-Id: I59752cef88d8d73bfd5b085f5400ae657f78504e
    Reviewed-on: https://go-review.googlesource.com/26664
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Marcel van Lohuizen <mpvl@golang.org>

commit 392bf3a9cfee297ec106d5a67c37d8edb4c8c183
Author: Chris Broadfoot <cbro@golang.org>
Date:   Mon Aug 8 16:56:22 2016 -0700

    doc/go1.7.html: update compress/flate section
    
    Updates #15810.
    
    Change-Id: Ifa7d2fd7fbfe58dff8541b18a11f007a5ff5818a
    Reviewed-on: https://go-review.googlesource.com/25591
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7a622740655bb5fcbd160eb96887032314842e6e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Aug 8 17:14:01 2016 +0000

    net/http: make Transport use new connection if over HTTP/2 concurrency limit
    
    The Go HTTP/1 client will make as many new TCP connections as the user requests.
    
    The HTTP/2 client tried to have that behavior, but the policy of
    whether a connection is re-usable didn't take into account the extra 1
    stream counting against SETTINGS_MAX_CONCURRENT_STREAMS so in practice
    users were getting errors.
    
    For example, if the server's advertised max concurrent streams is 100
    and 200 concurrrent Go HTTP requests ask for a connection at once, all
    200 will think they can reuse that TCP connection, but then 100 will
    fail later when the number of concurrent streams exceeds 100.
    
    Instead, recognize the "no cached connections" error value in the
    shouldRetryRequest method, so those 100 will retry a new connection.
    
    This is the conservative fix for Go 1.7 so users don't get errors, and
    to match the HTTP/1 behavior. Issues #13957 and #13774 are the more
    involved bugs for Go 1.8.
    
    Updates #16582
    Updates #13957
    
    Change-Id: I1f15a7ce60c07a4baebca87675836d6fe03993e8
    Reviewed-on: https://go-review.googlesource.com/25580
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 219ca602ab3f9d7d857bc1640e2b9e01475cdc3d
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Aug 6 10:12:03 2016 -0700

    doc: fix required OS X version inconsistency for binary downloads
    
    Updates #16625
    
    Change-Id: Icac6705828bd9b29379596ba64b34d922b9002c3
    Reviewed-on: https://go-review.googlesource.com/25548
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 26015b95634de49bc5b4ac998e8a2d1fcb8eca79
Author: Shenghou Ma <minux@golang.org>
Date:   Fri Aug 5 19:16:07 2016 -0400

    runtime: make stack 16-byte aligned for external code in _rt0_amd64_linux_lib
    
    Fixes #16618.
    
    Change-Id: Iffada12e8672bbdbcf2e787782c497e2c45701b1
    Reviewed-on: https://go-review.googlesource.com/25550
    Run-TryBot: Minux Ma <minux@golang.org>
    Reviewed-by: Arjan Van De Ven <arjan.van.de.ven@intel.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9fde86b0124b8c75000eb5d05887eff922a24566
Author: Shenghou Ma <minux@golang.org>
Date:   Thu Aug 4 21:34:06 2016 -0400

    runtime, syscall: fix kernel gettimeofday ABI change on iOS 10
    
    Fixes #16570 on iOS.
    
    Thanks Daniel Burhans for reporting the bug and testing the fix.
    
    Change-Id: I43ae7b78c8f85a131ed3d93ea59da9f32a02cd8f
    Reviewed-on: https://go-review.googlesource.com/25481
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3a03e877cc03c1fd155055e60a3f1f9cb8bda8d0
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Aug 4 19:53:52 2016 -0700

    os: check for waitid returning ENOSYS
    
    Reportedly waitid is not available for Ubuntu on Windows.
    
    Fixes #16610.
    
    Change-Id: Ia724f45a85c6d3467b847da06d8c65d280781dcd
    Reviewed-on: https://go-review.googlesource.com/25507
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 10316757cec3c2744ea61088e0fc905cfeb28fb2
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Aug 5 16:42:31 2016 +0000

    net/http: update bundled http2 for flow control window adjustment fix
    
    Updates bundled http2 to x/net/http2 git rev 075e191 for:
    
       http2: adjust flow control on open streams when processing SETTINGS
       https://golang.org/cl/25508
    
    Fixes #16612
    
    Change-Id: Ib0513201bff44ab747a574ae6894479325c105d2
    Reviewed-on: https://go-review.googlesource.com/25543
    Run-TryBot: Chris Broadfoot <cbro@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Chris Broadfoot <cbro@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit da070bed19fb23a8dd1b9f974c7fdf1f247fdba0
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Aug 4 12:21:05 2016 -0700

    syscall: fix Gettimeofday on macOS Sierra
    
    Fixes #16606
    
    Change-Id: I57465061b90e901293cd8b6ef756d6aa84ebd4a3
    Reviewed-on: https://go-review.googlesource.com/25495
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f135c326402aaa757aa96aad283a91873d4ae124
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Aug 4 13:09:29 2016 -0400

    runtime: initialize hash algs before typemap
    
    When compiling with -buildmode=shared, a map[int32]*_type is created for
    each extra module mapping duplicate types back to a canonical object.
    This is done in the function typelinksinit, which is called before the
    init function that sets up the hash functions for the map
    implementation. The result is typemap becomes unusable after
    runtime initialization.
    
    The fix in this CL is to move algorithm init before typelinksinit in
    the runtime setup process. (For 1.8, we may want to turn typemap into
    a sorted slice of types and use binary search.)
    
    Manually tested on GOOS=linux with:
    
            GOHOSTARCH=386 GOARCH=386 ./make.bash && \
                    go install -buildmode=shared std && \
                    cd ../test && \
                    go run run.go -linkshared
    
    Fixes #16590
    
    Change-Id: Idc08c50cc70d20028276fbf564509d2cd5405210
    Reviewed-on: https://go-review.googlesource.com/25469
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
```
