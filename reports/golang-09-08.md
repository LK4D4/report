# September 8, 2016 Report

Number of commits: 76

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 673.493348ms to 585.52799ms, -13.06%
* github.com/coreos/etcd: from 13.305897414s to 13.580010649s, +2.06%
* github.com/gogits/gogs: from 14.126338599s to 13.449209482s, -4.79%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 37.556980687s to 36.177503934s, -3.67%
* github.com/influxdata/influxdb/cmd/influxd: from 7.067195297s to 7.044287736s, -0.32%
* github.com/junegunn/fzf/src/fzf: from 1.040538943s to 1.101978384s, +5.90%
* github.com/mholt/caddy/caddy: from 6.085946741s to 6.162519742s, +1.26%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.398461523s to 1.540469242s, +10.15%
* github.com/nsqio/nsq/apps/nsqd: from 2.125554383s to 2.388553667s, +12.37%
* github.com/prometheus/prometheus/cmd/prometheus: from 11.459816911s to 11.151191333s, -2.69%
* github.com/spf13/hugo: from 8.080073993s to 8.09087006s, +0.13%
* golang.org/x/tools/cmd/guru: from 2.988210892s to 2.761086272s, -7.60%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2589363 to 2598068, +0.34%
* github.com/coreos/etcd: from 23589536 to 23590976, ~
* github.com/gogits/gogs: from 23125252 to 23121702, ~
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 27202104 to 27196320, ~
* github.com/influxdata/influxdb/cmd/influxd: from 16047028 to 16047590, ~
* github.com/junegunn/fzf/src/fzf: from 3129560 to 3131480, ~
* github.com/mholt/caddy/caddy: from 14591663 to 14591507, ~
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4425925 to 4426407, ~
* github.com/nsqio/nsq/apps/nsqd: from 9617376 to 9622037, ~
* github.com/prometheus/prometheus/cmd/prometheus: from 23464686 to 23465252, ~
* github.com/spf13/hugo: from 16011403 to 16016060, ~
* golang.org/x/tools/cmd/guru: from 7570389 to 7570871, ~

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               185           187           +1.08%
BenchmarkMsgpUnmarshal-4             392           398           +1.53%
BenchmarkEasyJsonMarshal-4           1459          1460          +0.07%
BenchmarkEasyJsonUnmarshal-4         2048          1866          -8.89%
BenchmarkFlatBuffersMarshal-4        402           359           -10.70%
BenchmarkFlatBuffersUnmarshal-4      285           287           +0.70%
BenchmarkGogoprotobufMarshal-4       158           158           +0.00%
BenchmarkGogoprotobufUnmarshal-4     250           244           -2.40%

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

* [math/rand: document that NewSource sources race](https://github.com/golang/go/commit/82bc0d4e80870f25805029ef0e1e844ace7bf068)
* [io: make MultiReader nil exhausted Readers for earlier GC](https://github.com/golang/go/commit/269ff8e6030cacd3a8ef5804f39c50566ce6f57e)
* [math: fix sqrt regression on AMD64](https://github.com/golang/go/commit/6e703ae7093b8921ce8e64a08e600d94ea1f9f28)
* [regexp: reduce mallocs in Regexp.Find\* and Regexp.ReplaceAll\*.](https://github.com/golang/go/commit/bea39e63ecf0e29323a93b3353a40eacbd815dc9)
* [runtime: bound scanobject to ~100 µs](https://github.com/golang/go/commit/cf4f1d07a189125a8774a923a3259126599e942b)
* [syscall: make Getpagesize return page size from runtime](https://github.com/golang/go/commit/1b9499b06989d2831e5b156161d6c07642926ee1)
* [bytes: make IndexRune faster](https://github.com/golang/go/commit/e10286aeda6b1412f8f64734412bff74836637f9)
* [strings: use AVX2 for Index if available](https://github.com/golang/go/commit/0cff219c1279cb76f042004bffcefba0a169cb67)


## GIT Log

```
commit 3a59b5626da498de0e74a5c02298f04a330f2911
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Sep 6 14:26:56 2016 -0700

    cmd/compile: remove unnecessary FuncType cloning
    
    Since FuncTypes are represented as structs rather than linking the
    parameter lists together, we no longer need to worry about duplicating
    the parameter lists.
    
    Change-Id: I3767aa3cd1cbeddfb80a6eef6b42290dc2ac14ae
    Reviewed-on: https://go-review.googlesource.com/28574
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 24d8f3fa4b02784af2419eec8a28aee303aae0c5
Author: Sina Siadat <siadat@gmail.com>
Date:   Sun Sep 4 12:20:14 2016 +0430

    net/http/httputil: copy header map if necessary in ReverseProxy
    
    We were already making a copy of the map before removing
    hop-by-hop headers. This commit does the same for proxied
    headers mentioned in the "Connection" header.
    
    A test is added to ensure request headers are not modified.
    
    Updates #16875
    
    Change-Id: I85329d212787958d5ad818915eb0538580a4653a
    Reviewed-on: https://go-review.googlesource.com/28493
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit b6f44923c0f88eb36816d90fb8fff2fd78422df5
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Wed Aug 10 19:01:06 2016 +0200

    go/format: add format.Node example
    
    Updates #16360
    
    Change-Id: I5927cffa961cd85539a3ba9606b116c5996d1096
    Reviewed-on: https://go-review.googlesource.com/26696
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 614dfe9b02d69f96f4b222d818ec5ff47f26cb31
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Sep 5 21:13:24 2016 +0000

    io: add test that MultiReader zeros exhausted Readers
    
    Updates #16983
    Updates #16996
    
    Change-Id: I76390766385b2668632c95e172b2d243d7f66651
    Reviewed-on: https://go-review.googlesource.com/28771
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 9e040979bd07acfcd93d59667b5f51c6aa183047
Author: Chris Broadfoot <cbro@golang.org>
Date:   Wed Sep 7 11:59:58 2016 -0700

    doc: document go1.7.1
    
    Change-Id: I4dc1ff7bfc67351a046f199dee8b7a9eadb1e524
    Reviewed-on: https://go-review.googlesource.com/28693
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3ef0e8f8235fe938ad5d1c99859cc63470877ec7
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Fri Sep 2 01:14:57 2016 -0700

    net: remove parsing of negative decimals in IPv4 literal
    
    https://golang.org/cl/27206 fixed the dtoi function such that
    it now properly parses negative number. Ironically, this causes
    several other functions that depended on dtoi to now (incorrectly)
    parse negative numbers.
    
    For example, ParseCIDR("-1.0.0.0/32") used to be rejected prior to the
    above CL, but is now accepted even though it is an invalid CIDR notation.
    This CL fixes that regression.
    
    We fix this by removing the signed parsing logic entirely from dtoi.
    It was introduced relatively recently in https://golang.org/cl/12447
    to fix a bug where an invalid port was improperly being parsed as OK.
    It seems to me that the fix in that CL to the port handling logic was
    sufficient such that a change to dtoi was unnecessary.
    
    Updates #16350
    
    Change-Id: I414bb1aa27d0a226ebd4b05a09cb40d784691b43
    Reviewed-on: https://go-review.googlesource.com/28414
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>

commit 42433e27b0467c862b97f3515185f8053807c648
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Sep 2 15:12:49 2016 -0700

    cmd/compile: add format verification test
    
    TestFormats finds potential (Printf, etc.) format strings.
    If they are used in a call, the format verbs are verified
    based on the matching argument type against a precomputed
    table of valid formats (formatMapping, below). The table
    can be used to automatically rewrite format strings with
    the -u flag.
    
    Run as: go test -run Formats [-u]
    
    A formatMapping based on the existing formats is printed
    when the test is run in verbose mode (-v flag). The table
    needs to be updated whenever a new (type, format) combination
    is found and the format verb is not 'v' (as in "%v").
    
    Known bugs:
    - indexed format strings ("%[2]s", etc.) are not suported
      (the test will fail)
    - format strings that are not simple string literals cannot
      be updated automatically
      (the test will fail with respective warnings)
    
    Change-Id: I1ca5bb6421d57ac78a00f1a80b9547a72837adc9
    Reviewed-on: https://go-review.googlesource.com/28419
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 0cff219c1279cb76f042004bffcefba0a169cb67
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Thu Apr 28 17:39:55 2016 +0300

    strings: use AVX2 for Index if available
    
    IndexHard4-4      1.50ms ± 2%  0.71ms ± 0%  -52.36%  (p=0.000 n=20+19)
    
    This also fixes a bug, that caused a string of length 16 to use
    two 8-byte comparisons instead of one 16-byte. And adds a test for
    cases when partial_match fails.
    
    Change-Id: I1ee8fc4e068bb36c95c45de78f067c822c0d9df0
    Reviewed-on: https://go-review.googlesource.com/22551
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 83c73a85db84a04c8e60e52cfa348fc6b675fbf7
Author: Keith Randall <khr@golang.org>
Date:   Tue Sep 6 21:08:21 2016 -0700

    cmd/compile: ignore contentEscapes for marking nodes as escaping
    
    Redo of CL 28575 with fixed test.
    We're in a pre-KeepAlive world for a bit yet, the old tests
    were in a client which was in a post-KeepAlive world.
    
    Change-Id: I114fd630339d761ab3306d1d99718d3cb973678d
    Reviewed-on: https://go-review.googlesource.com/28582
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ca264cdc6247141a1e042f38c83fff4783f03324
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Sep 4 19:07:22 2016 +0000

    syscall: avoid convT2I allocs for common Windows error values
    
    This is was already done for Unix in https://golang.org/cl/6701 +
    https://golang.org/cl/8192. Do it for Windows also.
    
    Fixes #16988
    
    Change-Id: Ia7832b0d0d48566b0cd205652b85130df529592e
    Reviewed-on: https://go-review.googlesource.com/28484
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 238274df4eee1752f51b288a11eddaf5365123bf
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Sep 6 17:26:15 2016 -0700

    go/constant: document that Value.String and ExactString return quoted strings
    
    This has always been the case but it was not obvious from the documentation.
    The reason for the quoting is that String() may return an abbreviated string,
    starting with double-quote (") but ending in ... (w/o a quote). The missing
    quote indicates the abbreviation (in contrast to a string ending in ...").
    
    constant.StringVal can be used to obtain the unquoted string of a String Value.
    
    Change-Id: Id0ba45b6ff62b3e024386ba8d907d6b3a4fcb6d7
    Reviewed-on: https://go-review.googlesource.com/28576
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit bdb3b790c66444c388529fa1d9b3f4d6aaa4c13f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Sep 7 03:23:20 2016 +0000

    Revert of cmd/compile: ignore contentEscapes for marking nodes as escaping
    
    Reason for revert: broke the build due to cherrypick;
    relies on an unsubmitted parent CL.
    
    Original issue's description:
    > cmd/compile: ignore contentEscapes for marking nodes as escaping
    >
    > We can still stack allocate and VarKill nodes which don't
    > escape but their content does.
    >
    > Fixes #16996
    >
    > Change-Id: If8aa0fcf2c327b4cb880a3d5af8d213289e6f6bf
    > Reviewed-on: https://go-review.googlesource.com/28575
    > Run-TryBot: Keith Randall <khr@golang.org>
    > TryBot-Result: Gobot Gobot <gobot@golang.org>
    > Reviewed-by: David Chase <drchase@google.com>
    >
    
    Change-Id: Ie1a325209de14d70af6acb2d78269b7a0450da7a
    Reviewed-on: https://go-review.googlesource.com/28578
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 923a74ce7710c1e3b24b4cc3220e2ba38d7673af
Author: Keith Randall <khr@golang.org>
Date:   Tue Sep 6 14:48:47 2016 -0700

    cmd/compile: ignore contentEscapes for marking nodes as escaping
    
    We can still stack allocate and VarKill nodes which don't
    escape but their content does.
    
    Fixes #16996
    
    Change-Id: If8aa0fcf2c327b4cb880a3d5af8d213289e6f6bf
    Reviewed-on: https://go-review.googlesource.com/28575
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit a6edffb28696b739136172995ed3ec000e8e6fdb
Author: David Chase <drchase@google.com>
Date:   Tue Sep 6 13:39:22 2016 -0700

    cmd/compile: add BVC/BVS to branch ops in ppc64/prog.go
    
    Includes test case shown to fail with unpatched compiler.
    
    Fixes #17005.
    
    Change-Id: I49b7b1a3f02736d85846a2588018b73f68d50320
    Reviewed-on: https://go-review.googlesource.com/28573
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 8737dac1f27db2596f1d24aab8e5c942734c3bb4
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Tue Sep 6 20:23:40 2016 +0900

    strings: make IndexRune faster
    
    re-implement IndexRune by Index which is well optimized to get
    performance gain.
    
    name                   old time/op  new time/op  delta
    IndexRune-4            30.2ns ± 1%  28.3ns ± 1%   -6.22%  (p=0.000 n=20+19)
    IndexRuneLongString-4   156ns ± 1%    49ns ± 1%  -68.72%  (p=0.000 n=19+19)
    IndexRuneFastPath-4    10.6ns ± 2%  10.0ns ± 1%   -6.30%  (p=0.000 n=18+18)
    
    Change-Id: Ie663b8f7860ca51892dd4be182fca3caa5f8ae61
    Reviewed-on: https://go-review.googlesource.com/28546
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a4bdd645550608fdb39bfea0bb83eb39b95c6c0b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 6 17:02:19 2016 +0000

    syscall: use MNT_NOWAIT in TestGetfsstat
    
    Fixes test failure when VMWare's shared folder filesystem is present.
    
    MNT_NOWAIT is what the mount(8) command does.
    
    Fixes #16937
    
    Change-Id: Id436185f544b7069db46c8716d6a0bf580b31da0
    Reviewed-on: https://go-review.googlesource.com/28550
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 9a243303b86a0c5aaf25875a7d07dabd629b0662
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Sep 4 16:59:46 2016 -0700

    cmd/compile: omit some temp panicdivide calls
    
    When the divisor is known to be a constant
    non-zero, don't insert panicdivide calls
    that will just be eliminated later.
    
    The main benefit here is readability of the SSA
    form for compiler developers.
    
    Change-Id: Icb7d07fc996941fbaff84524ac3e4b53d8e75fda
    Reviewed-on: https://go-review.googlesource.com/28530
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit e10286aeda6b1412f8f64734412bff74836637f9
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Tue Sep 6 08:09:27 2016 +0900

    bytes: make IndexRune faster
    
    re-implement IndexRune by IndexByte and Index which are well optimized
    to get performance gain.
    
    name                  old time/op   new time/op     delta
    IndexRune/10-4         53.2ns ± 1%     29.1ns ± 1%    -45.32%  (p=0.008 n=5+5)
    IndexRune/32-4          191ns ± 1%       27ns ± 1%    -85.75%  (p=0.008 n=5+5)
    IndexRune/4K-4         23.5µs ± 1%      1.0µs ± 1%    -95.77%  (p=0.008 n=5+5)
    IndexRune/4M-4         23.8ms ± 0%      1.0ms ± 2%    -95.90%  (p=0.008 n=5+5)
    IndexRune/64M-4         384ms ± 1%       15ms ± 1%    -95.98%  (p=0.008 n=5+5)
    IndexRuneASCII/10-4    61.5ns ± 0%     10.3ns ± 4%    -83.17%  (p=0.008 n=5+5)
    IndexRuneASCII/32-4     203ns ± 0%       11ns ± 5%    -94.68%  (p=0.008 n=5+5)
    IndexRuneASCII/4K-4    23.4µs ± 0%      0.3µs ± 2%    -98.60%  (p=0.008 n=5+5)
    IndexRuneASCII/4M-4    24.0ms ± 1%      0.3ms ± 1%    -98.60%  (p=0.008 n=5+5)
    IndexRuneASCII/64M-4    386ms ± 2%        6ms ± 1%    -98.57%  (p=0.008 n=5+5)
    
    name                  old speed     new speed       delta
    IndexRune/10-4        188MB/s ± 1%    344MB/s ± 1%    +82.91%  (p=0.008 n=5+5)
    IndexRune/32-4        167MB/s ± 0%   1175MB/s ± 1%   +603.52%  (p=0.008 n=5+5)
    IndexRune/4K-4        174MB/s ± 1%   4117MB/s ± 1%  +2262.71%  (p=0.008 n=5+5)
    IndexRune/4M-4        176MB/s ± 0%   4299MB/s ± 2%  +2340.46%  (p=0.008 n=5+5)
    IndexRune/64M-4       175MB/s ± 1%   4354MB/s ± 1%  +2388.57%  (p=0.008 n=5+5)
    IndexRuneASCII/10-4   163MB/s ± 0%    968MB/s ± 4%   +494.66%  (p=0.008 n=5+5)
    IndexRuneASCII/32-4   157MB/s ± 0%   2974MB/s ± 4%  +1788.59%  (p=0.008 n=5+5)
    IndexRuneASCII/4K-4   175MB/s ± 0%  12481MB/s ± 2%  +7027.71%  (p=0.008 n=5+5)
    IndexRuneASCII/4M-4   175MB/s ± 1%  12510MB/s ± 1%  +7061.15%  (p=0.008 n=5+5)
    IndexRuneASCII/64M-4  174MB/s ± 2%  12143MB/s ± 1%  +6881.70%  (p=0.008 n=5+5)
    
    Change-Id: I0632eadb83937c2a9daa7f0ce79df1dee64f992e
    Reviewed-on: https://go-review.googlesource.com/28537
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8259cf3c72e149ccdec2070d7f885119d92d94c6
Author: Austin Clements <austin@google.com>
Date:   Fri Aug 19 15:51:00 2016 -0400

    runtime/debug: enable TestFreeOSMemory on all arches
    
    TestFreeOSMemory was disabled on many arches because of issue #9993.
    Since that's been fixed, enable the test everywhere.
    
    Change-Id: I298c38c3e04128d9c8a1f558980939d5699bea03
    Reviewed-on: https://go-review.googlesource.com/27403
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 1b9499b06989d2831e5b156161d6c07642926ee1
Author: Austin Clements <austin@google.com>
Date:   Mon Jul 18 22:00:45 2016 -0400

    syscall: make Getpagesize return page size from runtime
    
    syscall.Getpagesize currently returns hard-coded page sizes on all
    architectures (some of which are probably always wrong, and some of
    which are definitely not always right). The runtime now has this
    information, queried from the OS during runtime init, so make
    syscall.Getpagesize return the page size that the runtime knows.
    
    Updates #10180.
    
    Change-Id: I4daa6fbc61a2193eb8fa9e7878960971205ac346
    Reviewed-on: https://go-review.googlesource.com/25051
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6dda7b2f5fb675a2747fea5ae020248245b8903f
Author: Austin Clements <austin@google.com>
Date:   Mon Jul 18 12:24:02 2016 -0400

    runtime: don't hard-code physical page size
    
    Now that the runtime fetches the true physical page size from the OS,
    make the physical page size used by heap growth a variable instead of
    a constant. This isn't used in any performance-critical paths, so it
    shouldn't be an issue.
    
    sys.PhysPageSize is also renamed to sys.DefaultPhysPageSize to make it
    clear that it's not necessarily the true page size. There are no uses
    of this constant any more, but we'll keep it around for now.
    
    Updates #12480 and #10180.
    
    Change-Id: I6c23b9df860db309c38c8287a703c53817754f03
    Reviewed-on: https://go-review.googlesource.com/25022
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 276a52de55fb48c4e56a778f1f7cac9292d8fad7
Author: Austin Clements <austin@google.com>
Date:   Mon Jul 18 21:40:02 2016 -0400

    runtime: fetch physical page size from the OS
    
    Currently the physical page size assumed by the runtime is hard-coded.
    On Linux the runtime at least fetches the OS page size during init and
    sanity checks against the hard-coded value, but they may still differ.
    On other OSes we wouldn't even notice.
    
    Add support on all OSes to fetch the actual OS physical page size
    during runtime init and lift the sanity check of PhysPageSize from the
    Linux init code to general malloc init. Currently this is the only use
    of the retrieved page size, but we'll add more shortly.
    
    Updates #12480 and #10180.
    
    Change-Id: I065f2834bc97c71d3208edc17fd990ec9058b6da
    Reviewed-on: https://go-review.googlesource.com/25050
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit d7de8b6d231289b7a6b205508c6f02a5a475cc84
Author: Austin Clements <austin@google.com>
Date:   Mon Jul 18 12:32:29 2016 -0400

    runtime: assume 64kB physical pages on ARM
    
    Currently we assume the physical page size on ARM is 4kB. While this
    is usually true, the architecture also supports 16kB and 64kB physical
    pages, and Linux (and possibly other OSes) can be configured to use
    these larger page sizes.
    
    With Go 1.6, such a configuration could potentially run, but generally
    resulted in memory corruption or random panics. With current master,
    this configuration will cause the runtime to panic during init on
    Linux when it checks the true physical page size (and will still cause
    corruption or panics on other OSes).
    
    However, the assumed physical page size only has to be a multiple of
    the true physical page size, the scavenger can now deal with large
    physical page sizes, and the rest of the runtime can deal with a
    larger assumed physical page size than the true size. Hence, there's
    little disadvantage to conservatively setting the assumed physical
    page size to 64kB on ARM.
    
    This may result in some extra memory use, since we can only return
    memory at multiples of the assumed physical page size. However, it is
    a simple change that should make Go run on systems configured for
    larger page sizes. The following commits will make the runtime query
    the actual physical page size from the OS, but this is a simple step
    there.
    
    Updates #12480.
    
    Change-Id: I851829595bc9e0c76235c847a7b5f62ad82b5302
    Reviewed-on: https://go-review.googlesource.com/25021
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit cf4f1d07a189125a8774a923a3259126599e942b
Author: Austin Clements <austin@google.com>
Date:   Fri May 27 21:04:40 2016 -0400

    runtime: bound scanobject to ~100 µs
    
    Currently the time spent in scanobject is proportional to the size of
    the object being scanned. Since scanobject is non-preemptible, large
    objects can cause significant goroutine (and even whole application)
    delays through several means:
    
    1. If a GC assist picks up a large object, the allocating goroutine is
       blocked for the whole scan, even if that scan well exceeds that
       goroutine's debt.
    
    2. Since the scheduler does not run on the P performing a large object
       scan, goroutines in that P's run queue do not run unless they are
       stolen by another P (which can take some time). If there are a few
       large objects, all of the Ps may get tied up so the scheduler
       doesn't run anywhere.
    
    3. Even if a large object is scanned by a background worker and other
       Ps are still running the scheduler, the large object scan doesn't
       flush background credit until the whole scan is done. This can
       easily cause all allocations to block in assists, waiting for
       credit, causing an effective STW.
    
    Fix this by splitting large objects into 128 KB "oblets" and scanning
    at most one oblet at a time. Since we can scan 1–2 MB/ms, this equates
    to bounding scanobject at roughly 100 µs. This improves assist
    behavior both because assists can no longer get "unlucky" and be stuck
    scanning a large object, and because it causes the background worker
    to flush credit and unblock assists more frequently when scanning
    large objects. This also improves GC parallelism if the heap consists
    primarily of a small number of very large objects by letting multiple
    workers scan a large objects in parallel.
    
    Fixes #10345. Fixes #16293.
    
    This substantially improves goroutine latency in the benchmark from
    issue #16293, which exercises several forms of very large objects:
    
    name                 old max-latency    new max-latency    delta
    SliceNoPointer-12           154µs ± 1%        155µs ±  2%     ~     (p=0.087 n=13+12)
    SlicePointer-12             314ms ± 1%       5.94ms ±138%  -98.11%  (p=0.000 n=19+20)
    SliceLivePointer-12        1148ms ± 0%       4.72ms ±167%  -99.59%  (p=0.000 n=19+20)
    MapNoPointer-12           72509µs ± 1%        408µs ±325%  -99.44%  (p=0.000 n=19+18)
    ChanPointer-12              313ms ± 0%       4.74ms ±140%  -98.49%  (p=0.000 n=18+20)
    ChanLivePointer-12         1147ms ± 0%       3.30ms ±149%  -99.71%  (p=0.000 n=19+20)
    
    name                 old P99.9-latency  new P99.9-latency  delta
    SliceNoPointer-12           113µs ±25%         107µs ±12%     ~     (p=0.153 n=20+18)
    SlicePointer-12          309450µs ± 0%         133µs ±23%  -99.96%  (p=0.000 n=20+20)
    SliceLivePointer-12         961ms ± 0%        1.35ms ±27%  -99.86%  (p=0.000 n=20+20)
    MapNoPointer-12            448µs ±288%         119µs ±18%  -73.34%  (p=0.000 n=18+20)
    ChanPointer-12           309450µs ± 0%         134µs ±23%  -99.96%  (p=0.000 n=20+19)
    ChanLivePointer-12          961ms ± 0%        1.35ms ±27%  -99.86%  (p=0.000 n=20+20)
    
    This has negligible effect on all metrics from the garbage, JSON, and
    HTTP x/benchmarks.
    
    It shows slight improvement on some of the go1 benchmarks,
    particularly Revcomp, which uses some multi-megabyte buffers:
    
    name                      old time/op    new time/op    delta
    BinaryTree17-12              2.46s ± 1%     2.47s ± 1%  +0.32%  (p=0.012 n=20+20)
    Fannkuch11-12                2.82s ± 0%     2.81s ± 0%  -0.61%  (p=0.000 n=17+20)
    FmtFprintfEmpty-12          50.8ns ± 5%    50.5ns ± 2%    ~     (p=0.197 n=17+19)
    FmtFprintfString-12          131ns ± 1%     132ns ± 0%  +0.57%  (p=0.000 n=20+16)
    FmtFprintfInt-12             117ns ± 0%     116ns ± 0%  -0.47%  (p=0.000 n=15+20)
    FmtFprintfIntInt-12          180ns ± 0%     179ns ± 1%  -0.78%  (p=0.000 n=16+20)
    FmtFprintfPrefixedInt-12     186ns ± 1%     185ns ± 1%  -0.55%  (p=0.000 n=19+20)
    FmtFprintfFloat-12           263ns ± 1%     271ns ± 0%  +2.84%  (p=0.000 n=18+20)
    FmtManyArgs-12               741ns ± 1%     742ns ± 1%    ~     (p=0.190 n=19+19)
    GobDecode-12                7.44ms ± 0%    7.35ms ± 1%  -1.21%  (p=0.000 n=20+20)
    GobEncode-12                6.22ms ± 1%    6.21ms ± 1%    ~     (p=0.336 n=20+19)
    Gzip-12                      220ms ± 1%     219ms ± 1%    ~     (p=0.130 n=19+19)
    Gunzip-12                   37.9ms ± 0%    37.9ms ± 1%    ~     (p=1.000 n=20+19)
    HTTPClientServer-12         82.5µs ± 3%    82.6µs ± 3%    ~     (p=0.776 n=20+19)
    JSONEncode-12               16.4ms ± 1%    16.5ms ± 2%  +0.49%  (p=0.003 n=18+19)
    JSONDecode-12               53.7ms ± 1%    54.1ms ± 1%  +0.71%  (p=0.000 n=19+18)
    Mandelbrot200-12            4.19ms ± 1%    4.20ms ± 1%    ~     (p=0.452 n=19+19)
    GoParse-12                  3.38ms ± 1%    3.37ms ± 1%    ~     (p=0.123 n=19+19)
    RegexpMatchEasy0_32-12      72.1ns ± 1%    71.8ns ± 1%    ~     (p=0.397 n=19+17)
    RegexpMatchEasy0_1K-12       242ns ± 0%     242ns ± 0%    ~     (p=0.168 n=17+20)
    RegexpMatchEasy1_32-12      72.1ns ± 1%    72.1ns ± 1%    ~     (p=0.538 n=18+19)
    RegexpMatchEasy1_1K-12       385ns ± 1%     384ns ± 1%    ~     (p=0.388 n=20+20)
    RegexpMatchMedium_32-12      112ns ± 1%     112ns ± 3%    ~     (p=0.539 n=20+20)
    RegexpMatchMedium_1K-12     34.4µs ± 2%    34.4µs ± 2%    ~     (p=0.628 n=18+18)
    RegexpMatchHard_32-12       1.80µs ± 1%    1.80µs ± 1%    ~     (p=0.522 n=18+19)
    RegexpMatchHard_1K-12       54.0µs ± 1%    54.1µs ± 1%    ~     (p=0.647 n=20+19)
    Revcomp-12                   387ms ± 1%     369ms ± 5%  -4.89%  (p=0.000 n=17+19)
    Template-12                 62.3ms ± 1%    62.0ms ± 0%  -0.48%  (p=0.002 n=20+17)
    TimeParse-12                 314ns ± 1%     314ns ± 0%    ~     (p=1.011 n=20+13)
    TimeFormat-12                358ns ± 0%     354ns ± 0%  -1.12%  (p=0.000 n=17+20)
    [Geo mean]                  53.5µs         53.3µs       -0.23%
    
    Change-Id: I2a0a179d1d6bf7875dd054b7693dd12d2a340132
    Reviewed-on: https://go-review.googlesource.com/23540
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit b275e55d86c78b8fdcfc70ea835ab3e00c7d6eeb
Author: Austin Clements <austin@google.com>
Date:   Fri Sep 2 13:56:52 2016 -0400

    runtime: clean up more traces of the old mark bit
    
    Commit 59877bf renamed bitMarked to bitScan, since the bitmap is no
    longer used for marking. However, there were several other references
    to this strewn about comments and in some other constant names. Fix
    these up, too.
    
    Change-Id: I4183d28c6b01977f1d75a99ad55b150f2211772d
    Reviewed-on: https://go-review.googlesource.com/28450
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rick Hudson <rlh@golang.org>

commit 4d5bb76279a7b0043bef97156a39ae1bef923e6d
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Sep 6 13:56:48 2016 -0400

    cmd/compile: remove nil check if followed by storezero on ARM64, MIPS64
    
    Change-Id: Ib90c92056fa70b27feb734837794ef53e842c41a
    Reviewed-on: https://go-review.googlesource.com/28513
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 0e0ab20334358ab009cbcd9d570ef6e884750733
Author: David Chase <drchase@google.com>
Date:   Tue Sep 6 09:05:02 2016 -0700

    cmd/compile: remove ld/st-followed nil checks for PPC64
    
    Enabled checks (except for DUFF-ops which aren't implemented yet).
    Added ppc64le to relevant test.
    
    Also updated register list to reflect no-longer-reserved-
    for-constants status (file was missed in that change).
    
    Updates #16010.
    
    Change-Id: I31b1aac19e14994f760f2ecd02edbeb1f78362e7
    Reviewed-on: https://go-review.googlesource.com/28548
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit b926bf83b0d58d7cb177dae46c011847442498c2
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Sep 6 12:33:36 2016 -0400

    cmd/link: remove outdated cast and comment
    
    This program is written in Go now.
    
    Change-Id: Ieec21a1bcac7c7a59e88cd1e1359977659de1757
    Reviewed-on: https://go-review.googlesource.com/28549
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bea39e63ecf0e29323a93b3353a40eacbd815dc9
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Wed May 11 14:57:24 2016 +0300

    regexp: reduce mallocs in Regexp.Find* and Regexp.ReplaceAll*.
    
    This improves Regexp.Find* and Regexp.ReplaceAll* speed:
    
    name                  old time/op    new time/op    delta
    Find-4                   345ns ± 1%     314ns ± 1%    -8.94%    (p=0.000 n=9+8)
    FindString-4             341ns ± 1%     308ns ± 0%    -9.85%   (p=0.000 n=10+9)
    FindSubmatch-4           440ns ± 1%     404ns ± 0%    -8.27%   (p=0.000 n=10+8)
    FindStringSubmatch-4     426ns ± 0%     387ns ± 0%    -9.07%   (p=0.000 n=10+9)
    ReplaceAll-4            1.75µs ± 1%    1.67µs ± 0%    -4.45%   (p=0.000 n=9+10)
    
    name                  old alloc/op   new alloc/op   delta
    Find-4                   16.0B ± 0%     0.0B ±NaN%  -100.00%  (p=0.000 n=10+10)
    FindString-4             16.0B ± 0%     0.0B ±NaN%  -100.00%  (p=0.000 n=10+10)
    FindSubmatch-4           80.0B ± 0%     48.0B ± 0%   -40.00%  (p=0.000 n=10+10)
    FindStringSubmatch-4     64.0B ± 0%     32.0B ± 0%   -50.00%  (p=0.000 n=10+10)
    ReplaceAll-4              152B ± 0%      104B ± 0%   -31.58%  (p=0.000 n=10+10)
    
    name                  old allocs/op  new allocs/op  delta
    Find-4                    1.00 ± 0%     0.00 ±NaN%  -100.00%  (p=0.000 n=10+10)
    FindString-4              1.00 ± 0%     0.00 ±NaN%  -100.00%  (p=0.000 n=10+10)
    FindSubmatch-4            2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    FindStringSubmatch-4      2.00 ± 0%      1.00 ± 0%   -50.00%  (p=0.000 n=10+10)
    ReplaceAll-4              8.00 ± 0%      5.00 ± 0%   -37.50%  (p=0.000 n=10+10)
    
    Fixes #15643
    
    Change-Id: I594fe51172373e2adb98d1d25c76ca2cde54ff48
    Reviewed-on: https://go-review.googlesource.com/23030
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5923df1af9f058a4f1bf095dfb1d2722cd4120a1
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Aug 25 20:29:15 2016 -0400

    cmd/compile: generate table of main symbol types
    
    For each exported symbol in package main, add its name and type to
    go.plugin.tabs symbol. This is used by the runtime when loading a
    plugin to return a typed interface{} value.
    
    Change-Id: I23c39583e57180acb8f7a74d218dae4368614f46
    Reviewed-on: https://go-review.googlesource.com/27818
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6e703ae7093b8921ce8e64a08e600d94ea1f9f28
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Thu Sep 1 21:04:08 2016 +0300

    math: fix sqrt regression on AMD64
    
    1.7 introduced a significant regression compared to 1.6:
    
    SqrtIndirect-4  2.32ns ± 0%  7.86ns ± 0%  +238.79%        (p=0.000 n=20+18)
    
    This is caused by sqrtsd preserving upper part of destination register.
    Which introduces dependency on previous  value of X0.
    In 1.6 benchmark loop didn't use X0 immediately after call:
    
    callq  *%rbx
    movsd  0x8(%rsp),%xmm2
    movsd  0x20(%rsp),%xmm1
    addsd  %xmm2,%xmm1
    mov    0x18(%rsp),%rax
    inc    %rax
    jmp    loop
    
    In 1.7 however xmm0 is used just after call:
    
    callq  *%rbx
    mov    0x10(%rsp),%rcx
    lea    0x1(%rcx),%rax
    movsd  0x8(%rsp),%xmm0
    movsd  0x18(%rsp),%xmm1
    
    I've  verified that this is caused by dependency, by inserting
    XORPS X0,X0 in the beginning of math.Sqrt, which puts performance back on 1.6 level.
    
    Splitting SQRTSD mem,reg into:
    MOVSD mem,reg
    SQRTSD reg,reg
    
    Removes dependency, because MOVSD (load version)
    doesn't need to preserve upper part of a register.
    And reg,reg operation is solved by renamer in CPU.
    
    As a result of this change regression is gone:
    SqrtIndirect-4  7.86ns ± 0%  2.33ns ± 0%  -70.36%  (p=0.000 n=18+17)
    
    This also removes old Sqrt benchmarks, in favor of benchmarks measuring latency.
    Only SqrtIndirect is kept, to show impact of this patch.
    
    Change-Id: Ic7eebe8866445adff5bc38192fa8d64c9a6b8872
    Reviewed-on: https://go-review.googlesource.com/28392
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 6bcca5e91836cae42fc9a88c4ac9a6a5ffed67f6
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Sep 6 07:43:15 2016 -0700

    cmd/go: run mkalldocs.sh
    
    This should have happened as part of CL 28485.
    
    Change-Id: I63cd31303e542ceaec3f4002c5573f186a1e9a52
    Reviewed-on: https://go-review.googlesource.com/28547
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 644c16c76cc11034b124763a039a7164cb3e1abc
Author: Cherry Zhang <cherryyz@google.com>
Date:   Tue Sep 6 08:48:14 2016 -0400

    cmd/compile: fix intrinsifying sync/atomic.Swap* on AMD64
    
    It should alias to Xchg instead of Swap. Found when testing #16985.
    
    Change-Id: If9fd734a1f89b8b2656f421eb31b9d1b0d95a49f
    Reviewed-on: https://go-review.googlesource.com/28512
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit f1ef5a06d296fd0ba604a62ebb30531cef7ae74a
Author: Cherry Zhang <cherryyz@google.com>
Date:   Sun Sep 4 23:51:22 2016 -0400

    cmd/compile: mark some AMD64 atomic ops as clobberFlags
    
    Fixes #16985.
    
    Change-Id: I5954db28f7b70dd3ac7768e471d5df871a5b20f9
    Reviewed-on: https://go-review.googlesource.com/28510
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 6db13e071b8b35b9efc8aeae6434217733ee8e94
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Sep 6 04:11:59 2016 +0000

    syscall: add yet more TestGetfsstat debugging
    
    Updates #16937
    
    Change-Id: I98aa203176f8f2ca2fcca6e334a65bc60d6f824d
    Reviewed-on: https://go-review.googlesource.com/28535
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 66121ce8a95f587be8641f79c66c8536660a16d5
Author: Erik Staab <estaab@google.com>
Date:   Sat Sep 3 15:57:48 2016 -0700

    runtime: remove redundant expression from SetFinalizer
    
    The previous if condition already checks the same expression and doesn't
    have side effects.
    
    Change-Id: Ieaf30a786572b608d0a883052b45fd3f04bc6147
    Reviewed-on: https://go-review.googlesource.com/28475
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5514332ed9e770bfb37011f6500a27be53f905c4
Author: Shenghou Ma <minux@golang.org>
Date:   Mon Sep 5 05:19:30 2016 -0400

    os: deduplicate File definition
    
    Fixes #16993.
    
    Change-Id: Ibe406f97d2a49acae94531d969c56dbac8ce53b2
    Reviewed-on: https://go-review.googlesource.com/28511
    Run-TryBot: Minux Ma <minux@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 269ff8e6030cacd3a8ef5804f39c50566ce6f57e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Sep 5 20:16:16 2016 +0000

    io: make MultiReader nil exhausted Readers for earlier GC
    
    No test because the language spec makes no promises in this area.
    
    Fixes #16983
    
    Change-Id: I1a6aa7ff87dd14aa27e8400040a6f6fc908aa1fd
    Reviewed-on: https://go-review.googlesource.com/28533
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 29272b1e611bf20b706e26757e5d0d872e61adff
Author: Caio Marcelo de Oliveira Filho <caio.oliveira@intel.com>
Date:   Mon Sep 5 15:14:55 2016 -0300

    cmd/go: use httpGET helper in bug command
    
    Use existing helper function instead of importing "net/http". This
    allows the go_bootstrap build to not depend on "net/http" package.
    See cmd/go/http.go for details.
    
    Fixes build bootstrap build with all.bash.
    
    Change-Id: I2fd0fb01af7774f1690a353af22137680ec78170
    Reviewed-on: https://go-review.googlesource.com/28531
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2939f395e0ab9cd70fb1af51f65af711d2f5f222
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Aug 28 10:38:41 2016 -0700

    cmd/go: add bug command
    
    This is a slightly rough, skeletal implementation.
    We will polish and add to it through use.
    
    .github/ISSUE_TEMPLATE will be updated in a
    separate CL.
    
    Fixes #16635
    
    Change-Id: Icf284170d87e61b5b643366c85cffc48f149f730
    Reviewed-on: https://go-review.googlesource.com/28485
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d58886409398abfa1a3aed1864e9ee14d516c089
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Sep 5 01:20:50 2016 +0000

    syscall: add more TestGetfsstat debugging
    
    Updates #16937
    
    Change-Id: I6d1b210c741269b58040bd68bf3b51644f891737
    Reviewed-on: https://go-review.googlesource.com/28487
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 460b76aef98490b3cf8d374c589db631eab85957
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Sep 4 09:34:03 2016 -0700

    cmd/compile: clean up ctxt params in sinit
    
    The ctxt parameter is always set to 0 on entry into anylit so make this
    parameter a literal constant, and where possibly remove ctxt as a parameter
    where it is known to be a constant zero.
    
    Passes toolstash -cmp.
    
    This is a re-creation of CL 28221 by Dave Cheney.
    That CL was graciously reverted in CL 28480
    to make merging other CLs easier.
    
    Change-Id: If7a57bf0e27774d9890adbc30af9fabb4aff1058
    Reviewed-on: https://go-review.googlesource.com/28483
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit 199b17cca8daf5bac6210e5719b1e85a51bd311d
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Fri Aug 26 13:50:34 2016 +0900

    path/filepath: handle "C:." correctly in EvalSymlinks on Windows
    
    Fixes #16886
    
    Change-Id: Idfacb0cf44d9994559c8e09032b4595887e76433
    Reviewed-on: https://go-review.googlesource.com/28214
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1a04b4abe78a152da5ccd801bf13a1df7ddfa8aa
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Sep 4 14:52:59 2016 -0700

    net/http/httputil: t.Error -> t.Errorf
    
    Found by vet.
    
    Change-Id: I09b79d68c7a5fc97e0edda4700a82bfbb00a4f45
    Reviewed-on: https://go-review.googlesource.com/28486
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit db9796dac50c65d300d2b262f2f107d86bf1703f
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Jun 20 08:18:22 2016 -0700

    cmd/compile: simplify staticname
    
    Add docs.
    Give it a more natural signature.
    
    Passes toolstash -cmp.
    
    Change-Id: Iab368dd10e8f16e41b725c2980020bbead2cdefb
    Reviewed-on: https://go-review.googlesource.com/26756
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7b5df0c195330d9c776139ffc8caaab2c66bbf2a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Jun 19 14:12:59 2016 -0700

    cmd/compile: document sinit ctxt and pass/kind arguments
    
    No functional changes. Passes toolstash -cmp.
    
    Change-Id: I1ad467e574fd2ea80ab1459c0c943d9ff66c23ec
    Reviewed-on: https://go-review.googlesource.com/26755
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit cd2d5ee0485fd0fa5a102c9de372c70b7a31d154
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Jun 19 12:11:47 2016 -0700

    cmd/compile: unify arraylit and structlit
    
    They were almost identical.
    Merge them and some of their calling code.
    
    Passes toolstash -cmp.
    
    Change-Id: I9e92a864a6c09c9e18ed52dc247a678467e344ba
    Reviewed-on: https://go-review.googlesource.com/26754
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 49cce1a62e452c1f2e2acbed37038837edef937d
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Jun 19 07:20:28 2016 -0700

    cmd/compile: add OSLICELIT
    
    Does not pass toolstash -cmp due to changed export data,
    but the cmd/go binary (which doesn't contain export data)
    is bit-for-bit identical.
    
    Change-Id: I6b12f9de18cf7da528e9207dccbf8f08c969f142
    Reviewed-on: https://go-review.googlesource.com/26753
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 41e1c420280b04431ae3a757ba9ec9e51963b64d
Author: Michael Pratt <mpratt@google.com>
Date:   Sat Sep 3 18:48:30 2016 -0700

    cmd/compile: refactor out KeepAlive
    
    Reduce the duplication in every arch by moving the code into package gc.
    
    Change-Id: Ia111add8316492571825431ecd4f0154c8792ae1
    Reviewed-on: https://go-review.googlesource.com/28481
    Run-TryBot: Michael Pratt <mpratt@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 960016eca27f2e727886c51ed98dd5ae47c150dc
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sat Sep 3 18:29:43 2016 -0700

    compress/flate: clarify the behavior of Writer.Flush
    
    Fixes #16068
    
    Change-Id: I04e80a181c0b7356996f7a1158ea4895ff9e1e39
    Reviewed-on: https://go-review.googlesource.com/28477
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f1ebe8a03c397fd5a14dde71994f170c303bbf26
Author: Dave Cheney <dave@cheney.net>
Date:   Sun Sep 4 01:39:16 2016 +0000

    Revert "cmd/compile/internal/gc: clean up sinit.go"
    
    Revert to make josharians branch land cleanly
    
    This reverts commit 38abd43b6a4d215375901d137a3eac9d0d3393a5.
    
    Change-Id: Idde1df953baf6e5742c87c4edd4bee0b6b418aca
    Reviewed-on: https://go-review.googlesource.com/28480
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4c8baa0ad0f709b857710d34e662f168c16c533f
Author: Michael Pratt <mpratt@google.com>
Date:   Sat Sep 3 18:32:51 2016 -0700

    cmd/compile: use CheckLoweredPhi on PPC64
    
    This custom version is identical to CheckLoweredPhi. The addition of
    CheckLoweredPhi likely raced with adding PPC64.
    
    Change-Id: I294dcb758d312e93fb8842f4d1e12bf0f63a1e06
    Reviewed-on: https://go-review.googlesource.com/28479
    Run-TryBot: Michael Pratt <mpratt@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 3a67d595f50ff7d589d4a097e162707342f1bbc4
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Sat Sep 3 12:01:27 2016 -0700

    strconv: fix function name in errors for Atoi
    
    Fixes #16980
    
    Change-Id: I902a02b157c2c7d1772f5122b850dc48b1d7a224
    Reviewed-on: https://go-review.googlesource.com/28474
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit be8a1c61392c97328db0026c7c80cd7d69c0aa36
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Sep 3 16:31:56 2016 -0700

    test: add test for issue 15895
    
    It was fixed earlier in the Go 1.8 cycle.
    Add a test.
    
    Fixes #15895
    
    Change-Id: I5834831235d99b9fcf21b435932cdd7ac6dc2c6e
    Reviewed-on: https://go-review.googlesource.com/28476
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fd41951c2b243a80e7b2b5d28a90839c928cfd2e
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sat Sep 3 18:39:25 2016 +0200

    unicode/utf8: reduce bounds checks in EncodeRune
    
    Provide bounds elim hints in EncodeRune.
    
    name                  old time/op  new time/op  delta
    EncodeASCIIRune-4     2.69ns ± 2%  2.69ns ± 2%    ~     (p=0.193 n=47+46)
    EncodeJapaneseRune-4  5.97ns ± 2%  5.38ns ± 2%  -9.93%  (p=0.000 n=49+50)
    
    Change-Id: I1a6dcffff3bdd64ab93c2130021e3b00981de4c8
    Reviewed-on: https://go-review.googlesource.com/28492
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 38abd43b6a4d215375901d137a3eac9d0d3393a5
Author: Dave Cheney <dave@cheney.net>
Date:   Wed Aug 31 14:39:41 2016 +1000

    cmd/compile/internal/gc: clean up sinit.go
    
    The ctxt parameter is always set to 0 on entry into anylit so make this
    parameter a literal constant, and where possibly remove ctxt as a parameter
    where it is known to be a constant zero.
    
    Change-Id: I3e76e06456d7b1a1ea875ffeb2efefa4a1ff5a7e
    Reviewed-on: https://go-review.googlesource.com/28221
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7cbc1058ea9240d9df92a339932d5c6dce694e7a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Sep 2 05:00:05 2016 +0000

    net/http/httputil: make ReverseProxy send nil Body requests when possible
    
    The http.Transport's retry can't retry requests with non-nil
    bodies. When cloning an incoming server request into an outgoing
    client request, nil out the Body field if the ContentLength is 0. (For
    server requests, Body is always non-nil, even for GET, HEAD, etc.)
    
    Also, don't use the deprecated CancelRequest and use Context instead.
    
    And don't set Proto, ProtoMajor, ProtoMinor. Those are ignored in
    client requests, which was probably a later documentation
    clarification.
    
    Fixes #16036
    Updates #16696 (remove useless Proto lines)
    
    Change-Id: I70a869e9bd4bf240c5838e82fb5aa695a539b343
    Reviewed-on: https://go-review.googlesource.com/28412
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 0b84a64da173d811d01a8a59545c22a7e1fd986a
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Fri Sep 2 14:37:35 2016 -0700

    archive/tar: reapply Header.Size to regFileReader after merging
    
    The use of PAX headers can modify the overall file size, thus the
    formerly created regFileReader may be stale.
    
    The relevant PAX specification for this behavior is:
    <<<
    Any fields in the preceding optional extended header shall override
    the associated fields in this header block for this file.
    >>>
    Where "optional extended header" refers to the preceding PAX header.
    Where "this header block" refers to the subsequent USTAR header.
    
    Fixes #15573
    Fixes #15564
    
    Change-Id: I83b1c3f05a9ca2d3be38647425ad21a9fe450ee2
    Reviewed-on: https://go-review.googlesource.com/28418
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 549ca046ffeb1f76833c15059d3a5da301cf1eb3
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Sep 2 14:21:57 2016 -0700

    cmd/compile: fix argument for given format verb
    
    Follow-up to https://go-review.googlesource.com/28394.
    
    Change-Id: Ic4147e9ae786a4de0a3454131fac03e940ae2e76
    Reviewed-on: https://go-review.googlesource.com/28417
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 0318d80e51ff148f3a94008b59b36d601846152e
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Thu Sep 1 23:28:26 2016 -0700

    path/filepath: use new style deprecation message
    
    Change-Id: I242a8960583e333f372929aad4adb8efbe441cd4
    Reviewed-on: https://go-review.googlesource.com/28413
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Jaana Burcu Dogan <jbd@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0e7e43688d2ad8b6c78bb865591eec96b6ce0e60
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun May 1 18:02:41 2016 -0700

    runtime: remove a load and shift from scanobject
    
    hbits.morePointers and hbits.isPointer both
    do a load and a shift. Do it only once.
    
    Benchmarks using compilebench (because it is
    the benchmark I have the most tooling around),
    on a quiet machine.
    
    name       old time/op      new time/op      delta
    Template        291ms ±14%       290ms ±15%    ~          (p=0.702 n=100+99)
    Unicode         143ms ± 9%       142ms ± 9%    ~           (p=0.126 n=99+98)
    GoTypes         934ms ± 4%       933ms ± 4%    ~         (p=0.937 n=100+100)
    Compiler        4.92s ± 2%       4.90s ± 1%  -0.28%        (p=0.003 n=98+98)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   360user-ms ± 5%  355user-ms ± 4%  -1.37%        (p=0.000 n=97+96)
    Unicode    178user-ms ± 6%  176user-ms ± 6%  -1.24%        (p=0.001 n=96+99)
    GoTypes    1.22user-s ± 5%  1.21user-s ± 5%  -0.94%      (p=0.000 n=100+100)
    Compiler   6.50user-s ± 2%  6.44user-s ± 3%  -0.94%        (p=0.000 n=96+98)
    
    On amd64, before:
    
    "".scanobject t=1 size=581 args=0x10 locals=0x78
    
    After:
    
    "".scanobject t=1 size=540 args=0x10 locals=0x78
    
    
    Change-Id: I420ac3704549d484a5d85e19fea82c85da389514
    Reviewed-on: https://go-review.googlesource.com/22712
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cd285f1c6fc613fa0f097443ae1d21d6c4491386
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Mon Aug 8 16:17:14 2016 +0200

    runtime: fix global buffer reset in StopTrace
    
    We reset global buffer only if its pos != 0.
    We ought to do it always, but queue it only if pos != 0.
    This is a latent bug. Currently it does not fire because
    whenever we create a global buffer, we increment pos.
    
    Change-Id: I01e28ae88ce9a5412497c524391b8b7cb443ffd9
    Reviewed-on: https://go-review.googlesource.com/25574
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 59877bfaaf0778d4cc2cbaf06cf5030144271349
Author: Gleb Stepanov <glebstepanov1992@gmail.com>
Date:   Mon Jul 25 16:25:44 2016 +0300

    runtime: rename variable
    
    Rename variable to bitScan according to
    TODO comment.
    
    Change-Id: I81dd8cc1ca28c0dc9308a654ad65cdf5b2fd2ce3
    Reviewed-on: https://go-review.googlesource.com/25175
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit 3df926d52addd638e23e5b29e692bb670e484966
Author: Austin Clements <austin@google.com>
Date:   Fri Sep 2 12:14:38 2016 -0400

    runtime: improve message when a bad pointer is found on the stack
    
    Currently this message says "invalid stack pointer", which could be
    interpreted as the value of SP being invalid. Change it to "invalid
    pointer found on stack" to emphasize that it's a pointer on the stack
    that's invalid.
    
    Updates #16948.
    
    Change-Id: I753624f8cc7e08cf13d3ea5d9c790cc4af9fa372
    Reviewed-on: https://go-review.googlesource.com/28430
    Run-TryBot: Austin Clements <austin@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 210ac4d5e0fea2bfd4287b0865104bdaaeaffe05
Author: Adam Langley <agl@golang.org>
Date:   Thu Sep 1 16:00:25 2016 -0700

    crypto/cipher: enforce message size limits for GCM.
    
    The maximum input plaintext for GCM is 64GiB - 64. Since the GCM
    interface is one-shot, it's very hard to hit this in Go (one would need
    a 64GiB buffer in memory), but we should still enforce this limit.
    
    Thanks to Quan Nguyen for pointing it out.
    
    Change-Id: Icced47bf8d4d5dfbefa165cf13e893205c9577b8
    Reviewed-on: https://go-review.googlesource.com/28410
    Run-TryBot: Adam Langley <agl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 03cff2e115277951108d3e00298ae1cb0b0b7fb6
Author: Sina Siadat <siadat@gmail.com>
Date:   Sat Aug 27 20:46:25 2016 +0430

    net/http/httputil: remove proxied headers mentioned in connection-tokens
    
    RFC 2616, section 14.10 says:
    
    >>>
    HTTP/1.1 proxies MUST parse the Connection header field before a message
    is forwarded and, for each connection-token in this field, remove any
    header field(s) from the message with the same name as the
    connection-token. Connection options are signaled by the presence of a
    connection-token in the Connection header field, not by any
    corresponding additional header field(s), since the additional header
    field may not be sent if there are no parameters associated with that
    connection option.
    <<<
    
    The same requirement was included in RFC 7230, section 6.1.
    
    Fixes #16875
    
    Change-Id: I57ad4a4a17775537c8810d0edd7de1604317b5fa
    Reviewed-on: https://go-review.googlesource.com/27970
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 82bc0d4e80870f25805029ef0e1e844ace7bf068
Author: David Glasser <glasser@meteor.com>
Date:   Thu Aug 4 16:49:05 2016 -0700

    math/rand: document that NewSource sources race
    
    While it was previously explicitly documented that "the default Source"
    is safe for concurrent use, a careless reader can interpret that as
    meaning "the implementation of the Source interface created by functions
    in this package" rather than "the default shared Source used by
    top-level functions". Be explicit that the Source returned by NewSource
    is not safe for use by multiple goroutines.
    
    Fixes #3611.
    
    Change-Id: Iae4bc04c3887ad6e2491e36e38feda40324022c5
    Reviewed-on: https://go-review.googlesource.com/25501
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 98def53f5610247b3d635ff85bf47fd19e848c01
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Aug 30 18:11:10 2016 -0700

    cmd/dist: make test names consistent
    
    Current banners are:
    
    ##### Building Go bootstrap tool.
    ##### Building Go toolchain using /Users/josh/go/1.4.
    ##### Building go_bootstrap for host, darwin/amd64.
    ##### Building packages and commands for darwin/amd64.
    ##### Testing packages.
    ##### GOMAXPROCS=2 runtime -cpu=1,2,4
    ##### Testing without libgcc.
    ##### sync -cpu=10
    ##### ../misc/cgo/stdio
    ##### ../misc/cgo/life
    ##### ../misc/cgo/fortran
    ##### ../misc/cgo/test
    ##### Testing race detector
    ##### ../misc/cgo/testso
    ##### ../misc/cgo/testsovar
    ##### misc/cgo/testcarchive
    ##### ../misc/cgo/testcshared
    ##### ../misc/cgo/errors
    ##### ../test/bench/go1
    ##### ../test
    ##### API check
    
    One of these things is not like the others.
    Fix that.
    
    Change-Id: If0bd8ea9293d73b5d1b70d6bf676bd9192991505
    Reviewed-on: https://go-review.googlesource.com/26759
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7eddaf5f0b275cd2ff37b260246d276748ac6cd4
Author: Matt Layher <mdlayher@gmail.com>
Date:   Mon Aug 1 12:50:11 2016 -0400

    go/doc: allow ToHTML to properly handle URLs containing semicolons
    
    Fixes #16565
    
    Change-Id: I3edfd2576a7ca5270644a4e7f126854f821f2c9a
    Reviewed-on: https://go-review.googlesource.com/25385
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 88ccb3c945daeba7c08dfe9b39be18ec78941a45
Author: Matt Layher <mdlayher@gmail.com>
Date:   Thu Sep 1 16:30:15 2016 -0400

    net/http: omit Content-Length in Response.Write for 1xx or 204 status
    
    Per RFC 7230, Section 3.3.2: "A server MUST NOT send a Content-Length
    header field in any response with a status code of 1xx (Informational)
    or 204 (No Content).".
    
    Fixes #16942
    
    Change-Id: I8006c76c126304e13618966e6eafb08a3885d3cd
    Reviewed-on: https://go-review.googlesource.com/28351
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8895a99c9ff522cf41f3a1bee365bd0c7e0c7900
Author: Martin Möhrmann <martisch@uos.de>
Date:   Thu Sep 1 08:31:37 2016 +0200

    cmd/compile: disallow typed non-integer constant len and cap make arguments
    
    make(T, n, m) returns a slice of type T with length n and capacity m
    where "The size arguments n and m must be of integer type or untyped."
    https://tip.golang.org/ref/spec#Making_slices_maps_and_channels
    
    The failure to reject typed non-integer size arguments in make
    during compile time was uncovered after https://golang.org/cl/27851
    changed the generation of makeslice calls.
    
    Fixes   #16940
    Updates #16949
    
    Change-Id: Ib1e3576f0e6ad199c9b16b7a50c2db81290c63b4
    Reviewed-on: https://go-review.googlesource.com/28301
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 809bb3a71c2cdceb93c349888b186787589d5ec7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Sep 1 12:23:43 2016 -0700

    cmd/compile: fix missing format verb
    
    Found by vet.
    
    Change-Id: I50420771678b1a3695348ce1a81f410479ed09a1
    Reviewed-on: https://go-review.googlesource.com/28394
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit fe5619b479859f199a244929770a11cc4cbd1911
Author: Keith Randall <khr@golang.org>
Date:   Thu Sep 1 10:13:11 2016 -0700

    cmd/compile: be more aggressive in tighten pass for booleans
    
    Fixes #15509
    
    Change-Id: I44073533f02d38795f9ba9b255db4d1ee426d70e
    Reviewed-on: https://go-review.googlesource.com/28390
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit c53879512f38982b1e773fd2aeecee4ebf382006
Author: Keith Randall <khr@golang.org>
Date:   Wed Aug 31 13:36:10 2016 -0700

    cmd/compile: missing float indexed loads/stores on amd64
    
    Update #16141
    
    Change-Id: I7d32c5cdc197d86491a67ea579fa16cb3d675b51
    Reviewed-on: https://go-review.googlesource.com/28273
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 0f1a8d3c2ddce2d96d2b2bb81094ac3401aff03d
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Thu Sep 1 02:49:27 2016 -0700

    doc: fix stale gofrontend/gccgo contribution link
    
    Change-Id: I63af8f0a19ec91f4a2001aa7a2eadcd2232a47df
    Reviewed-on: https://go-review.googlesource.com/28348
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 44f1854c9dc82d8dba415ef102e93896d57c2c0d
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Thu Apr 28 17:34:24 2016 +0300

    bytes: Use the same algorithm as strings for Index
    
    name                     old time/op    new time/op      delta
    IndexByte32-48             9.05ns ± 7%      9.59ns ±11%     +5.93%  (p=0.001 n=19+20)
    IndexByte4K-48              118ns ± 4%       122ns ± 8%     +3.52%  (p=0.002 n=19+19)
    IndexByte4M-48              172µs ±13%       188µs ±12%     +9.49%  (p=0.000 n=20+20)
    IndexByte64M-48            8.00ms ±14%      8.05ms ±23%       ~     (p=0.799 n=20+20)
    IndexBytePortable32-48     41.7ns ±15%      42.5ns ±12%       ~     (p=0.372 n=20+20)
    IndexBytePortable4K-48     3.08µs ±16%      3.26µs ±10%     +5.77%  (p=0.018 n=20+20)
    IndexBytePortable4M-48     3.12ms ±17%      3.20ms ±10%       ~     (p=0.157 n=20+20)
    IndexBytePortable64M-48    54.0ms ±14%      55.3ms ±14%       ~     (p=0.640 n=20+20)
    Index32-48                  230ns ±12%        46ns ± 6%    -79.87%  (p=0.000 n=20+19)
    Index4K-48                 43.2µs ± 9%       3.2µs ±12%    -92.58%  (p=0.000 n=19+20)
    Index4M-48                 44.4ms ± 7%       3.3ms ±13%    -92.59%  (p=0.000 n=19+20)
    Index64M-48                 714ms ±10%        56ms ± 8%    -92.22%  (p=0.000 n=19+19)
    IndexEasy32-48             52.7ns ±10%      31.0ns ±11%    -41.21%  (p=0.000 n=20+20)
    IndexEasy4K-48              139ns ± 5%      1598ns ± 6%  +1046.37%  (p=0.000 n=19+19)
    IndexEasy4M-48              179µs ± 8%      1674µs ±10%   +834.31%  (p=0.000 n=19+20)
    IndexEasy64M-48            8.56ms ±10%     27.82ms ±16%   +225.14%  (p=0.000 n=19+20)
    
    name                     old speed      new speed        delta
    IndexByte32-48           3.52GB/s ± 7%    3.35GB/s ±11%     -4.99%  (p=0.001 n=20+20)
    IndexByte4K-48           34.5GB/s ± 7%    33.2GB/s ±10%     -3.67%  (p=0.002 n=20+20)
    IndexByte4M-48           24.6GB/s ±14%    22.4GB/s ±14%     -8.73%  (p=0.000 n=20+20)
    IndexByte64M-48          8.42GB/s ±16%    8.42GB/s ±19%       ~     (p=0.799 n=20+20)
    IndexBytePortable32-48    770MB/s ±13%     756MB/s ±11%       ~     (p=0.383 n=20+20)
    IndexBytePortable4K-48   1.34GB/s ±14%    1.26GB/s ±10%     -5.76%  (p=0.018 n=20+20)
    IndexBytePortable4M-48   1.35GB/s ±15%    1.31GB/s ±11%       ~     (p=0.157 n=20+20)
    IndexBytePortable64M-48  1.25GB/s ±16%    1.22GB/s ±13%       ~     (p=0.640 n=20+20)
    Index32-48                138MB/s ± 8%     687MB/s ± 8%   +398.57%  (p=0.000 n=19+20)
    Index4K-48               94.9MB/s ± 9%  1280.5MB/s ±11%  +1249.11%  (p=0.000 n=19+20)
    Index4M-48               94.6MB/s ± 7%  1278.5MB/s ±12%  +1250.99%  (p=0.000 n=19+20)
    Index64M-48              94.2MB/s ±10%  1210.9MB/s ± 8%  +1185.04%  (p=0.000 n=19+19)
    IndexEasy32-48            608MB/s ±10%    1035MB/s ±10%    +70.15%  (p=0.000 n=20+20)
    IndexEasy4K-48           29.3GB/s ± 6%     2.6GB/s ± 6%    -91.24%  (p=0.000 n=19+19)
    IndexEasy4M-48           23.3GB/s ±10%     2.5GB/s ± 9%    -89.23%  (p=0.000 n=20+20)
    IndexEasy64M-48          7.86GB/s ±11%    2.42GB/s ±14%    -69.18%  (p=0.000 n=19+20)
    
    Change-Id: Ia191f0a6ca80e113397d9ed98d25f195768b65bc
    Reviewed-on: https://go-review.googlesource.com/22550
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>
```
