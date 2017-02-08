Number of commits: 147

Compilation time:
	* github.com/coreos/etcd: from 12.719228932s to 12.368341084s, -2.76%
	* github.com/boltdb/bolt/cmd/bolt: from 640.152912ms to 609.511947ms, -4.79%
	* github.com/gogits/gogs: from 15.455176657s to 15.179491219s, -1.78%
	* github.com/spf13/hugo: from 7.464503324s to 7.283847394s, -2.42%
	* github.com/influxdata/influxdb/cmd/influxd: from 7.138511194s to 6.903522753s, -3.29%
	* github.com/nsqio/nsq/apps/nsqd: from 2.571217446s to 2.459312286s, -4.35%
	* github.com/mholt/caddy: from 5.589300803s to 5.475329036s, -2.04%

Binary size:
	* github.com/coreos/etcd: from 24024520 to 21281448, -11.42%
	* github.com/boltdb/bolt/cmd/bolt: from 2937015 to 2618526, -10.84%
	* github.com/gogits/gogs: from 25603247 to 23021534, -10.08%
	* github.com/spf13/hugo: from 16603030 to 14973948, -9.81%
	* github.com/influxdata/influxdb/cmd/influxd: from 17502631 to 15567739, -11.05%
	* github.com/nsqio/nsq/apps/nsqd: from 10956217 to 9845738, -10.14%
	* github.com/mholt/caddy: from 14814810 to 13266771, -10.45%

Bechmarks:
BenchmarkMsgpMarshal-4                   	10000000	       211 ns/op	     128 B/op	       1 allocs/op
BenchmarkMsgpUnmarshal-4                 	 3000000	       402 ns/op	     112 B/op	       3 allocs/op
BenchmarkVmihailencoMsgpackMarshal-4     	 1000000	      1819 ns/op	     352 B/op	       5 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal-4   	 1000000	      2104 ns/op	     352 B/op	      13 allocs/op
BenchmarkJsonMarshal-4                   	  500000	      3593 ns/op	    1232 B/op	      10 allocs/op
BenchmarkJsonUnmarshal-4                 	  500000	      3471 ns/op	     416 B/op	       7 allocs/op
BenchmarkEasyJsonMarshal-4               	 1000000	      1572 ns/op	     784 B/op	       5 allocs/op
BenchmarkEasyJsonUnmarshal-4             	 1000000	      1515 ns/op	     160 B/op	       4 allocs/op
BenchmarkBsonMarshal-4                   	 1000000	      1571 ns/op	     392 B/op	      10 allocs/op
BenchmarkBsonUnmarshal-4                 	 1000000	      2117 ns/op	     248 B/op	      21 allocs/op
BenchmarkGobMarshal-4                    	 1000000	      1068 ns/op	      48 B/op	       2 allocs/op
BenchmarkGobUnmarshal-4                  	 1000000	      1101 ns/op	     112 B/op	       3 allocs/op
BenchmarkXdrMarshal-4                    	 1000000	      1943 ns/op	     455 B/op	      20 allocs/op
BenchmarkXdrUnmarshal-4                  	 1000000	      1572 ns/op	     240 B/op	      11 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal-4     	  500000	      3072 ns/op	    2753 B/op	       8 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal-4   	  500000	      3309 ns/op	    3008 B/op	       6 allocs/op
BenchmarkUgorjiCodecBincMarshal-4        	  500000	      3110 ns/op	    2785 B/op	       8 allocs/op
BenchmarkUgorjiCodecBincUnmarshal-4      	  500000	      3597 ns/op	    3168 B/op	       9 allocs/op
BenchmarkSerealMarshal-4                 	  500000	      3576 ns/op	     912 B/op	      21 allocs/op
BenchmarkSerealUnmarshal-4               	  500000	      3717 ns/op	    1008 B/op	      34 allocs/op
BenchmarkBinaryMarshal-4                 	 1000000	      1695 ns/op	     256 B/op	      16 allocs/op
BenchmarkBinaryUnmarshal-4               	 1000000	      1907 ns/op	     336 B/op	      22 allocs/op
BenchmarkFlatbuffersMarshal-4            	 5000000	       367 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlatBuffersUnmarshal-4          	 5000000	       305 ns/op	     112 B/op	       3 allocs/op
BenchmarkCapNProtoMarshal-4              	 3000000	       545 ns/op	      56 B/op	       2 allocs/op
BenchmarkCapNProtoUnmarshal-4            	 3000000	       533 ns/op	     216 B/op	       6 allocs/op
BenchmarkCapNProto2Marshal-4             	 1000000	      1627 ns/op	     436 B/op	       7 allocs/op
BenchmarkCapNProto2Unmarshal-4           	 1000000	      1876 ns/op	     608 B/op	      12 allocs/op
BenchmarkHproseMarshal-4                 	 1000000	      1088 ns/op	     473 B/op	       8 allocs/op
BenchmarkHproseUnmarshal-4               	 1000000	      1269 ns/op	     320 B/op	      10 allocs/op
BenchmarkProtobufMarshal-4               	 1000000	      1171 ns/op	     200 B/op	       7 allocs/op
BenchmarkProtobufUnmarshal-4             	 2000000	       879 ns/op	     192 B/op	      10 allocs/op
BenchmarkGoprotobufMarshal-4             	 2000000	       629 ns/op	     312 B/op	       4 allocs/op
BenchmarkGoprotobufUnmarshal-4           	 2000000	       919 ns/op	     432 B/op	       9 allocs/op
BenchmarkGogoprotobufMarshal-4           	10000000	       188 ns/op	      64 B/op	       1 allocs/op
BenchmarkGogoprotobufUnmarshal-4         	 5000000	       280 ns/op	      96 B/op	       3 allocs/op
BenchmarkGencodeMarshal-4                	10000000	       199 ns/op	      80 B/op	       2 allocs/op
BenchmarkGencodeUnmarshal-4              	10000000	       243 ns/op	     112 B/op	       3 allocs/op
PASS
ok  	github.com/alecthomas/go_serialization_benchmarks	70.072s

benchmark                                  old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4                     225           211           -6.22%
BenchmarkMsgpUnmarshal-4                   404           402           -0.50%
BenchmarkVmihailencoMsgpackMarshal-4       1838          1819          -1.03%
BenchmarkVmihailencoMsgpackUnmarshal-4     2087          2104          +0.81%
BenchmarkJsonMarshal-4                     3682          3593          -2.42%
BenchmarkJsonUnmarshal-4                   3463          3471          +0.23%
BenchmarkEasyJsonMarshal-4                 1612          1572          -2.48%
BenchmarkEasyJsonUnmarshal-4               1507          1515          +0.53%
BenchmarkBsonMarshal-4                     1579          1571          -0.51%
BenchmarkBsonUnmarshal-4                   2113          2117          +0.19%
BenchmarkGobMarshal-4                      1093          1068          -2.29%
BenchmarkGobUnmarshal-4                    1123          1101          -1.96%
BenchmarkXdrMarshal-4                      1974          1943          -1.57%
BenchmarkXdrUnmarshal-4                    1513          1572          +3.90%
BenchmarkUgorjiCodecMsgpackMarshal-4       3000          3072          +2.40%
BenchmarkUgorjiCodecMsgpackUnmarshal-4     3259          3309          +1.53%
BenchmarkUgorjiCodecBincMarshal-4          3056          3110          +1.77%
BenchmarkUgorjiCodecBincUnmarshal-4        3587          3597          +0.28%
BenchmarkSerealMarshal-4                   3620          3576          -1.22%
BenchmarkSerealUnmarshal-4                 3796          3717          -2.08%
BenchmarkBinaryMarshal-4                   1694          1695          +0.06%
BenchmarkBinaryUnmarshal-4                 2038          1907          -6.43%
BenchmarkFlatbuffersMarshal-4              371           367           -1.08%
BenchmarkFlatBuffersUnmarshal-4            307           305           -0.65%
BenchmarkCapNProtoMarshal-4                541           545           +0.74%
BenchmarkCapNProtoUnmarshal-4              546           533           -2.38%
BenchmarkCapNProto2Marshal-4               1643          1627          -0.97%
BenchmarkCapNProto2Unmarshal-4             1858          1876          +0.97%
BenchmarkHproseMarshal-4                   1141          1088          -4.65%
BenchmarkHproseUnmarshal-4                 1289          1269          -1.55%
BenchmarkProtobufMarshal-4                 1188          1171          -1.43%
BenchmarkProtobufUnmarshal-4               926           879           -5.08%
BenchmarkGoprotobufMarshal-4               627           629           +0.32%
BenchmarkGoprotobufUnmarshal-4             943           919           -2.55%
BenchmarkGogoprotobufMarshal-4             197           188           -4.57%
BenchmarkGogoprotobufUnmarshal-4           281           280           -0.36%
BenchmarkGencodeMarshal-4                  201           199           -1.00%
BenchmarkGencodeUnmarshal-4                245           243           -0.82%

benchmark                                  old allocs     new allocs     delta
BenchmarkMsgpMarshal-4                     1              1              +0.00%
BenchmarkMsgpUnmarshal-4                   3              3              +0.00%
BenchmarkVmihailencoMsgpackMarshal-4       5              5              +0.00%
BenchmarkVmihailencoMsgpackUnmarshal-4     13             13             +0.00%
BenchmarkJsonMarshal-4                     10             10             +0.00%
BenchmarkJsonUnmarshal-4                   7              7              +0.00%
BenchmarkEasyJsonMarshal-4                 5              5              +0.00%
BenchmarkEasyJsonUnmarshal-4               4              4              +0.00%
BenchmarkBsonMarshal-4                     10             10             +0.00%
BenchmarkBsonUnmarshal-4                   21             21             +0.00%
BenchmarkGobMarshal-4                      2              2              +0.00%
BenchmarkGobUnmarshal-4                    3              3              +0.00%
BenchmarkXdrMarshal-4                      20             20             +0.00%
BenchmarkXdrUnmarshal-4                    11             11             +0.00%
BenchmarkUgorjiCodecMsgpackMarshal-4       8              8              +0.00%
BenchmarkUgorjiCodecMsgpackUnmarshal-4     6              6              +0.00%
BenchmarkUgorjiCodecBincMarshal-4          8              8              +0.00%
BenchmarkUgorjiCodecBincUnmarshal-4        9              9              +0.00%
BenchmarkSerealMarshal-4                   21             21             +0.00%
BenchmarkSerealUnmarshal-4                 34             34             +0.00%
BenchmarkBinaryMarshal-4                   16             16             +0.00%
BenchmarkBinaryUnmarshal-4                 22             22             +0.00%
BenchmarkFlatbuffersMarshal-4              0              0              +0.00%
BenchmarkFlatBuffersUnmarshal-4            3              3              +0.00%
BenchmarkCapNProtoMarshal-4                2              2              +0.00%
BenchmarkCapNProtoUnmarshal-4              6              6              +0.00%
BenchmarkCapNProto2Marshal-4               7              7              +0.00%
BenchmarkCapNProto2Unmarshal-4             12             12             +0.00%
BenchmarkHproseMarshal-4                   8              8              +0.00%
BenchmarkHproseUnmarshal-4                 10             10             +0.00%
BenchmarkProtobufMarshal-4                 7              7              +0.00%
BenchmarkProtobufUnmarshal-4               10             10             +0.00%
BenchmarkGoprotobufMarshal-4               4              4              +0.00%
BenchmarkGoprotobufUnmarshal-4             9              9              +0.00%
BenchmarkGogoprotobufMarshal-4             1              1              +0.00%
BenchmarkGogoprotobufUnmarshal-4           3              3              +0.00%
BenchmarkGencodeMarshal-4                  2              2              +0.00%
BenchmarkGencodeUnmarshal-4                3              3              +0.00%

benchmark                                  old bytes     new bytes     delta
BenchmarkMsgpMarshal-4                     128           128           +0.00%
BenchmarkMsgpUnmarshal-4                   112           112           +0.00%
BenchmarkVmihailencoMsgpackMarshal-4       352           352           +0.00%
BenchmarkVmihailencoMsgpackUnmarshal-4     351           352           +0.28%
BenchmarkJsonMarshal-4                     1232          1232          +0.00%
BenchmarkJsonUnmarshal-4                   416           416           +0.00%
BenchmarkEasyJsonMarshal-4                 784           784           +0.00%
BenchmarkEasyJsonUnmarshal-4               160           160           +0.00%
BenchmarkBsonMarshal-4                     392           392           +0.00%
BenchmarkBsonUnmarshal-4                   248           248           +0.00%
BenchmarkGobMarshal-4                      48            48            +0.00%
BenchmarkGobUnmarshal-4                    112           112           +0.00%
BenchmarkXdrMarshal-4                      455           455           +0.00%
BenchmarkXdrUnmarshal-4                    239           240           +0.42%
BenchmarkUgorjiCodecMsgpackMarshal-4       2753          2753          +0.00%
BenchmarkUgorjiCodecMsgpackUnmarshal-4     3008          3008          +0.00%
BenchmarkUgorjiCodecBincMarshal-4          2785          2785          +0.00%
BenchmarkUgorjiCodecBincUnmarshal-4        3168          3168          +0.00%
BenchmarkSerealMarshal-4                   912           912           +0.00%
BenchmarkSerealUnmarshal-4                 1008          1008          +0.00%
BenchmarkBinaryMarshal-4                   256           256           +0.00%
BenchmarkBinaryUnmarshal-4                 335           336           +0.30%
BenchmarkFlatbuffersMarshal-4              0             0             +0.00%
BenchmarkFlatBuffersUnmarshal-4            112           112           +0.00%
BenchmarkCapNProtoMarshal-4                56            56            +0.00%
BenchmarkCapNProtoUnmarshal-4              216           216           +0.00%
BenchmarkCapNProto2Marshal-4               436           436           +0.00%
BenchmarkCapNProto2Unmarshal-4             608           608           +0.00%
BenchmarkHproseMarshal-4                   473           473           +0.00%
BenchmarkHproseUnmarshal-4                 320           320           +0.00%
BenchmarkProtobufMarshal-4                 200           200           +0.00%
BenchmarkProtobufUnmarshal-4               192           192           +0.00%
BenchmarkGoprotobufMarshal-4               312           312           +0.00%
BenchmarkGoprotobufUnmarshal-4             432           432           +0.00%
BenchmarkGogoprotobufMarshal-4             64            64            +0.00%
BenchmarkGogoprotobufUnmarshal-4           96            96            +0.00%
BenchmarkGencodeMarshal-4                  80            80            +0.00%
BenchmarkGencodeUnmarshal-4                112           112           +0.00%

Highlights: 
	https://github.com/golang/go/commit/f6db855d5595e170bfc70d90f1eaa26034d2e191 net: mirror Tom Sawyer in the net package for tests
	https://github.com/golang/go/commit/6c6089b3fdba9eb0cff863a03074dbac47c92f63 cmd/compile: bce when max and limit are consts
	https://github.com/golang/go/commit/1faa8869c6c72f055cdaa2b547964830909c96c6 net/http: set the Request context for incoming server requests
	https://github.com/golang/go/commit/ad7448fe982d83de15deec9c55c56d0cd9261c6c runtime: speed up makeslice by avoiding divisions
	https://github.com/golang/go/commit/f20b1809f213c662932106a68c76ea3545eab1ee compress/flate: eliminate most common bounds checks
	https://github.com/golang/go/commit/e4f1d9cf2e948eb0f0bb91d7c253ab61dfff3a59 runtime: make execution error panic values implement the Error interface
	https://github.com/golang/go/commit/e88f89028a55acf9c8b76b7f6ca284c3f9eb4cbd bytes, string: add Reset method to Reader
	https://github.com/golang/go/commit/187afdebef7953295189d4531e7dccdc0cb64500 math/big: re-use memory in Int.GCD


commit 2c9d773f7411de211389b9e1da441fae68f826d8
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Thu Apr 14 16:31:42 2016 +0200

    misc/trace: update trace viewer html
    
    The old trace-viewer is broken since Chrome 49:
    https://bugs.chromium.org/p/chromium/issues/detail?id=569417
    It was fixed in:
    https://github.com/catapult-project/catapult/commit/506457cbd726324f327b80ae11f46c1dfeb8710d
    
    This change updates trace-viewer to the latest version
    (now it is called catapult).
    
    This version has a bug in the lean config that we use, though:
    https://github.com/catapult-project/catapult/issues/2247
    So use full config for now (it works, but leads to larger html).
    When the bug is fixed we need to switch back to lean config (issue #15302).
    
    Change-Id: Ifb8d782ced66e3292d81c5604039fe18eaf267c5
    Reviewed-on: https://go-review.googlesource.com/22013
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0ec6d7c0bbfceb7b8e4857b775686ae5cf699e54
Author: Marcel van Lohuizen <mpvl@golang.org>
Date:   Thu Apr 14 15:44:48 2016 +0800

    testing: removed flakey test
    
    The synchronization in this test is a bit complicated and likely
    incorrect, judging from the sporadically hanging trybots.
    Most of what this is supposed to test is already tested in
    TestTestContext, so I'll just remove it.
    
    Fixes #15170
    
    Change-Id: If54db977503caa109cec4516974eda9191051888
    Reviewed-on: https://go-review.googlesource.com/22080
    Run-TryBot: Marcel van Lohuizen <mpvl@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 285e78609f4fd85948d056f581d3443d5f9b230a
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Apr 14 14:53:12 2016 +0900

    net: fix TestDialAddrError
    
    Fixes #15291.
    
    Change-Id: I563140c2acd37d4989a940488b217414cf73f6c2
    Reviewed-on: https://go-review.googlesource.com/22077
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit ed7cd2546eb997e976534f0542816e12448f34d5
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Apr 14 12:17:44 2016 +0900

    net: make use of internal/testenv package
    
    Change-Id: I6644081df495cb92b3d208f867066f9acb08946f
    Reviewed-on: https://go-review.googlesource.com/22074
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8f64336edc7c725abcbe564d21b3d2dc5ec250ec
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Apr 14 09:57:43 2016 +0900

    net: make newLocalPacketListener handle network argument correcly
    
    Change-Id: I41691134770d01805c19c0f84f8828b00b85de0c
    Reviewed-on: https://go-review.googlesource.com/22058
    Run-TryBot: Mikio Hara <mikioh.mikioh@gmail.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit babfb4ec3ba3e4e36b1003d6efbaeddf2e975240
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Apr 13 18:41:59 2016 -0700

    cmd/internal/obj: change Link.Flag_shared to bool
    
    Change-Id: I9bda2ce6f45fb8292503f86d8f9f161601f222b7
    Reviewed-on: https://go-review.googlesource.com/22053
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 980ab12ade53e70d037ab2ab475148b216d84a14
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Apr 13 18:37:18 2016 -0700

    cmd/compile/internal/gc: change flags to bool where possible
    
    Some of the Debug[x] flags are actually boolean too, but not all, so
    they need to be handled separately.
    
    While here, change some obj.Flagstr and obj.Flagint64 calls to
    directly use flag.StringVar and flag.Int64Var instead.
    
    Change-Id: Iccedf6fed4328240ee2257f57fe6d66688f237c4
    Reviewed-on: https://go-review.googlesource.com/22052
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit ae9804595879eb07efd23b9c98eab46693573447
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Apr 13 16:57:23 2016 -0700

    cmd/compile: use correct export function (fix debugFormat)
    
    Tested with debugFormat enabled and running
    (export GO_GCFLAGS=-newexport; sh all.bash).
    
    Change-Id: If7d43e1e594ea43c644232b89e670f7abb6b003e
    Reviewed-on: https://go-review.googlesource.com/22033
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 933d521a7aa5defc46d3336bcb71a2f3f2b8172d
Author: Rob Pike <r@golang.org>
Date:   Wed Apr 13 11:47:25 2016 -0700

    fmt: clarify that for %g precision determines number of significant digits
    
    Documentation change only.
    
    Fixes #15178.
    
    Change-Id: I3c7d80ce9e668ac7515f7ebb9da80f3bd8e534d6
    Reviewed-on: https://go-review.googlesource.com/22006
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 6e5027a37a851eb19dba7dad7ea5a8b43e27b842
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Apr 13 13:17:30 2016 -0700

    cmd/compile: don't export unneeded OAS, OASWB nodes
    
    Also:
    - "rewrite" node Op in exporter for some nodes instead of importer
    - more comments
    
    Change-Id: I809e6754d14987b28f1da9379951ffa2e690c2a7
    Reviewed-on: https://go-review.googlesource.com/22008
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 44f80f6d4925ae59b519ced3a626170099258904
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Wed Apr 6 11:07:12 2016 -0500

    syscall:  fix epoll_event struct for ppc64le/ppc64
    
    The existing epoll_event structure used by many of
    the epoll_* syscalls was defined incorrectly
    for use with ppc64le & ppc64 in the syscall
    directory.  This resulted in the caller getting
    incorrect information on return from these
    syscalls.  This caused failures in fsnotify as
    well as builds with upstream Docker.  The
    structure is defined correctly in gccgo.
    
    This adds a pad field that is expected for
    these syscalls on ppc64le, ppc64.
    Fixes #15135
    
    Change-Id: If7e8ea9eb1d1ca5182c8dc0f935b334127341ffd
    Reviewed-on: https://go-review.googlesource.com/21582
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>

commit f120936dfffa3ac935730699587e6957f2d5ea61
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Mar 31 10:02:10 2016 -0400

    cmd/compile, etc: use name for type pkgPath
    
    By replacing the *string used to represent pkgPath with a
    reflect.name everywhere, the embedded *string for package paths
    inside the reflect.name can be replaced by an offset, nameOff.
    This reduces the number of pointers in the type information.
    
    This also moves all reflect.name types into the same section, making
    it possible to use nameOff more widely in later CLs.
    
    No significant binary size change for normal binaries, but:
    
    linux/amd64 PIE:
    	cmd/go: -440KB (3.7%)
    	jujud:  -2.6MB (3.2%)
    
    For #6853.
    
    Change-Id: I3890b132a784a1090b1b72b32febfe0bea77eaee
    Reviewed-on: https://go-review.googlesource.com/21395
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 73e2ad20220050f88b1ea79bf5a2e4c4fbee0533
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Apr 13 11:33:42 2016 -0700

    runtime: rename os1_darwin.go to os_darwin.go
    
    Change-Id: If0e0bc5a85101db1e70faaab168fc2d12024eb93
    Reviewed-on: https://go-review.googlesource.com/22005
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d9712aa82af7192469d75802c6dc1734ea9858b2
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Apr 13 11:31:24 2016 -0700

    runtime: merge the darwin os*.go files together
    
    Merge them together into os1_darwin.go. A future CL will rename it.
    
    Change-Id: Ia4380d3296ebd5ce210908ce3582ff184566f692
    Reviewed-on: https://go-review.googlesource.com/22004
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 79048df2ccc2d4c2ccc4e15d481f7888d48cf440
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed Apr 6 07:11:24 2016 -0400

    cmd/link: handle long symbol names
    
    Fixes #15104.
    
    Change-Id: I9ddfbbf39ef0a873b703ee3e04fbb7d1192f5f39
    Reviewed-on: https://go-review.googlesource.com/21581
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3ea7cfabbb0549d62d524e4ad30cb464af250fde
Author: Todd Neal <todd@tneal.org>
Date:   Wed Apr 13 08:51:46 2016 -0400

    cmd/compile: sort partitions by dom to speed up cse
    
    We do two O(n) scans of all values in an eqclass when computing
    substitutions for CSE.
    
    In unfortunate cases, like those found in #15112, we can have a large
    eqclass composed of values found in blocks none of whom dominate the
    other.  This leads to O(n^2) behavior. The elements are removed one at a
    time, with O(n) scans each time.
    
    This CL removes the linear scan by sorting the eqclass so that dominant
    values will be sorted first.  As long as we also ensure we don't disturb
    the sort order, then we no longer need to scan for the maximally
    dominant value.
    
    For the code in issue #15112:
    
    Before:
    real    1m26.094s
    user    1m30.776s
    sys     0m1.125s
    
    Aefter:
    real    0m52.099s
    user    0m56.829s
    sys     0m1.092s
    
    Updates #15112
    
    Change-Id: Ic4f8680ed172e716232436d31963209c146ef850
    Reviewed-on: https://go-review.googlesource.com/21981
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Todd Neal <todd@tneal.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4721ea6abcde318a2f5d61ec249cde5e9c57ebea
Author: Austin Clements <austin@google.com>
Date:   Wed Apr 13 11:22:42 2016 -0400

    runtime/internal/atomic: rename Storep1 to StorepNoWB
    
    Make it clear that the point of this function stores a pointer
    *without* a write barrier.
    
    sed -i -e 's/Storep1/StorepNoWB/' $(git grep -l Storep1)
    
    Updates #15270.
    
    Change-Id: Ifad7e17815e51a738070655fe3b178afdadaecf6
    Reviewed-on: https://go-review.googlesource.com/21994
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit da6205b67e844503152b3be7bbb1a25c76cbbce2
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri Apr 8 17:56:36 2016 +0200

    cmd/pprof/internal/profile: always subtract 1 from PCs
    
    Go runtime never emits PCs that are not a return address
    (except for cpu profiler).
    
    Change-Id: I08d9dc5c7c71e23f34f2f0c16f8baeeb4f64fcd6
    Reviewed-on: https://go-review.googlesource.com/21735
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit eb79f21c48915454b372de7fee2c6b86d52ea0bc
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Apr 12 21:58:44 2016 -0700

    cmd/compile, go/importer: minor cleanups
    
    Change-Id: Ic7a1fb0dbbf108052c970a4a830269a5673df7df
    Reviewed-on: https://go-review.googlesource.com/21963
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 7d0d1222477ce50736ee24adb38c1f487d0801d9
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Apr 12 18:00:04 2016 -0700

    cmd/compile: move more compiler specifics into compiler specific export section
    
    Instead of indicating with each function signature if it has an inlineable
    body, collect all functions in order and export function bodies with function
    index in platform-specific section.
    
    Moves this compiler specific information out of the platform-independent
    export data section, and removes an int value for all functions w/o body.
    Also simplifies the code a bit.
    
    Change-Id: I8b2d7299dbe81f2706be49ecfb9d9f7da85fd854
    Reviewed-on: https://go-review.googlesource.com/21939
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit d8e8fc292ace5ae59a0da44dfca1dd5b1a71ecf1
Author: Austin Clements <austin@google.com>
Date:   Wed Apr 13 11:13:39 2016 -0400

    runtime/internal/atomic: remove write barrier from Storep1 on s390x
    
    atomic.Storep1 is not supposed to invoke a write barrier (that's what
    atomicstorep is for), but currently does on s390x. This causes a panic
    in runtime.mapzero when it tries to use atomic.Storep1 to store what's
    actually a scalar.
    
    Fix this by eliminating the write barrier from atomic.Storep1 on
    s390x. Also add some documentation to atomicstorep to explain the
    difference between these.
    
    Fixes #15270.
    
    Change-Id: I291846732d82f090a218df3ef6351180aff54e81
    Reviewed-on: https://go-review.googlesource.com/21993
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Austin Clements <austin@google.com>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit 6b85a45edc94786c7669823ee47a6ce1156d6a9a
Author: David Chase <drchase@google.com>
Date:   Mon Mar 21 11:32:04 2016 -0400

    cmd/compile: move spills to loop exits when easy.
    
    For call-free inner loops.
    
    Revised statistics:
      85 inner loop spills sunk
     341 inner loop spills remaining
    1162 inner loop spills that were candidates for sinking
         ended up completely register allocated
     119 inner loop spills could have been sunk were used in
         "shuffling" at the bottom of the loop.
       1 inner loop spill not sunk because the register assigned
         changed between def and exit,
    
     Understanding how to make an inner loop definition not be
     a candidate for from-memory shuffling (to force the shuffle
     code to choose some other value) should pick up some of the
     119 other spills disqualified for this reason.
    
     Modified the stats printing based on feedback from Austin.
    
    Change-Id: If3fb9b5d5a028f42ccc36c4e3d9e0da39db5ca60
    Reviewed-on: https://go-review.googlesource.com/21037
    Reviewed-by: Keith Randall <khr@golang.org>
    Run-TryBot: David Chase <drchase@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c4807d4cc759025854e354fee99ac20d125f0d79
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Wed Apr 13 08:58:10 2016 -0500

    runtime: improve memmove performance ppc64,ppc64le
    
    This change improves the performance of memmove
    on ppc64 & ppc64le mainly for moves >=32 bytes.
    In addition, the test to detect backward moves
     was enhanced to avoid backward moves if source
    and dest were in different types of storage, since
    backward moves might not always be efficient.
    
    Fixes #14507
    
    The following shows some of the improvements from the test
    in the runtime package:
    
    BenchmarkMemmove32                   4229.56      4717.13      1.12x
    BenchmarkMemmove64                   6156.03      7810.42      1.27x
    BenchmarkMemmove128                  7521.69      12468.54     1.66x
    BenchmarkMemmove256                  6729.90      18260.33     2.71x
    BenchmarkMemmove512                  8521.59      18033.81     2.12x
    BenchmarkMemmove1024                 9760.92      25762.61     2.64x
    BenchmarkMemmove2048                 10241.00     29584.94     2.89x
    BenchmarkMemmove4096                 10399.37     31882.31     3.07x
    
    BenchmarkMemmoveUnalignedDst16       1943.69      2258.33      1.16x
    BenchmarkMemmoveUnalignedDst32       3885.08      3965.81      1.02x
    BenchmarkMemmoveUnalignedDst64       5121.63      6965.54      1.36x
    BenchmarkMemmoveUnalignedDst128      7212.34      11372.68     1.58x
    BenchmarkMemmoveUnalignedDst256      6564.52      16913.59     2.58x
    BenchmarkMemmoveUnalignedDst512      8364.35      17782.57     2.13x
    BenchmarkMemmoveUnalignedDst1024     9539.87      24914.72     2.61x
    BenchmarkMemmoveUnalignedDst2048     9199.23      21235.11     2.31x
    BenchmarkMemmoveUnalignedDst4096     10077.39     25231.99     2.50x
    
    BenchmarkMemmoveUnalignedSrc32       3249.83      3742.52      1.15x
    BenchmarkMemmoveUnalignedSrc64       5562.35      6627.96      1.19x
    BenchmarkMemmoveUnalignedSrc128      6023.98      10200.84     1.69x
    BenchmarkMemmoveUnalignedSrc256      6921.83      15258.43     2.20x
    BenchmarkMemmoveUnalignedSrc512      8593.13      16541.97     1.93x
    BenchmarkMemmoveUnalignedSrc1024     9730.95      22927.84     2.36x
    BenchmarkMemmoveUnalignedSrc2048     9793.28      21537.73     2.20x
    BenchmarkMemmoveUnalignedSrc4096     10132.96     26295.06     2.60x
    
    Change-Id: I73af59970d4c97c728deabb9708b31ec7e01bdf2
    Reviewed-on: https://go-review.googlesource.com/21990
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 66afbf1010fa492fb9a266f9019f707bd09f066d
Author: David Crawshaw <crawshaw@golang.org>
Date:   Wed Apr 13 10:06:12 2016 -0400

    cmd/link: use a switch for name prefix switching
    
    Minor cleanup.
    
    Change-Id: I7574f58a7e55c2bb798ebe9c7c98d36b8c258fb8
    Reviewed-on: https://go-review.googlesource.com/21982
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7d469179e6e3dafe16700b7fc1cf8683ad9453fa
Author: David Crawshaw <crawshaw@golang.org>
Date:   Mon Mar 28 10:32:27 2016 -0400

    cmd/compile, etc: store method tables as offsets
    
    This CL introduces the typeOff type and a lookup method of the same
    name that can turn a typeOff offset into an *rtype.
    
    In a typical Go binary (built with buildmode=exe, pie, c-archive, or
    c-shared), there is one moduledata and all typeOff values are offsets
    relative to firstmoduledata.types. This makes computing the pointer
    cheap in typical programs.
    
    With buildmode=shared (and one day, buildmode=plugin) there are
    multiple modules whose relative offset is determined at runtime.
    We identify a type in the general case by the pair of the original
    *rtype that references it and its typeOff value. We determine
    the module from the original pointer, and then use the typeOff from
    there to compute the final *rtype.
    
    To ensure there is only one *rtype representing each type, the
    runtime initializes a typemap for each module, using any identical
    type from an earlier module when resolving that offset. This means
    that types computed from an offset match the type mapped by the
    pointer dynamic relocations.
    
    A series of followup CLs will replace other *rtype values with typeOff
    (and name/*string with nameOff).
    
    For types created at runtime by reflect, type offsets are treated as
    global IDs and reference into a reflect offset map kept by the runtime.
    
    darwin/amd64:
    	cmd/go:  -57KB (0.6%)
    	jujud:  -557KB (0.8%)
    
    linux/amd64 PIE:
    	cmd/go: -361KB (3.0%)
    	jujud:  -3.5MB (4.2%)
    
    For #6853.
    
    Change-Id: Icf096fd884a0a0cb9f280f46f7a26c70a9006c96
    Reviewed-on: https://go-review.googlesource.com/21285
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e0611b16645dba6768cab405f1ec1b3fce83334a
Author: Alexandru Moșoi <mosoi@google.com>
Date:   Wed Apr 13 10:58:38 2016 +0200

    cmd/compile: use shared dom tree for cse, too
    
    Missed this in the previous CL where the shared
    dom tree was introduced.
    
    Change-Id: If0bd85d4b4567d7e87814ed511603b1303ab3903
    Reviewed-on: https://go-review.googlesource.com/21970
    Run-TryBot: Alexandru Moșoi <alexandru@mosoi.ro>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 61b7a9c57bb6b9c259360239001b2d5be4876abd
Author: Shahar Kohanim <skohanim@gmail.com>
Date:   Tue Apr 12 23:18:47 2016 +0300

    cmd/link: rename Pcln to FuncInfo
    
    After non pcln fields were added to it in a previous commit.
    
    Change-Id: Icf92c0774d157c61399a6fc2a3c4d2cd47a634d2
    Reviewed-on: https://go-review.googlesource.com/21921
    Run-TryBot: Shahar Kohanim <skohanim@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 24967ec122710e73b35893925fd9a8390d7524ab
Author: Tal Shprecher <tshprecher@gmail.com>
Date:   Sun Apr 10 18:12:41 2016 -0700

    cmd/compile: make enqueued map keys fail validation on forward types
    
    Map keys are currently validated in multiple locations but share
    a common validation routine. The problem is that early validations
    should be lenient enough to allow for forward types while the final
    validations should not. The final validations should fail on forward
    types since they've already settled.
    
    This change also separates the key type checking from the creation
    of the map via typMap. Instead of the mapqueue being populated in
    copytype() by checking the map line number, it's populated in the
    same block that validates the key type. This isolates key validation
    logic while type checking.
    
    Fixes #14988
    
    Change-Id: Ia47cf6213585d6c63b3a35249104c0439feae658
    Reviewed-on: https://go-review.googlesource.com/21830
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0e01db4b8d6ac64e6661508bc6876fa41c799208
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Apr 12 17:46:41 2016 -0700

    cmd/compile: fix crash on bare package name in constant declarations
    
    Fixes #11361.
    
    Change-Id: I70b8808f97f0e07de680e7e6ede1322ea0fdbbc0
    Reviewed-on: https://go-review.googlesource.com/21936
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6531fab06fc4667b7d167c7e3ee936f28bac68e2
Author: Tal Shprecher <tshprecher@gmail.com>
Date:   Tue Apr 12 22:29:34 2016 -0700

    cmd/compile: remove unnecessary assignments while type checking.
    
    Change-Id: Ica0ec84714d7f01d800d62fa10cdb08321d43cf3
    Reviewed-on: https://go-review.googlesource.com/21967
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 6af4e996e2f0408f159a8553d11122b9fe052ffb
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Apr 12 15:51:24 2016 -0700

    runtime: simplify setPanicOnFault slightly
    
    No need to acquire the M just to change G's paniconfault flag, and the
    original C implementation of SetPanicOnFault did not. The M
    acquisition logic is an artifact of golang.org/cl/131010044, which was
    started before golang.org/cl/123640043 (which introduced the current
    "getg" function) was submitted.
    
    Change-Id: I6d1939008660210be46904395cf5f5bbc2c8f754
    Reviewed-on: https://go-review.googlesource.com/21935
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 381e5eee39edfb3a43e294023957aff05e9ff349
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Apr 13 04:35:37 2016 +0000

    all: use new io.SeekFoo constants instead of os.SEEK_FOO
    
    Automated change.
    
    Fixes #15269
    
    Change-Id: I8deb2ac0101d3f7c390467ceb0a1561b72edbb2f
    Reviewed-on: https://go-review.googlesource.com/21962
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 819e0b29bbbc07022e7b94c12b55860466a02e5b
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sat Mar 26 00:04:48 2016 +0100

    strings: improve explode and correct comment
    
    Merges explodetests into splittests which already contain
    some of the tests that cover explode.
    
    Adds a test to cover the utf8.RuneError branch in explode.
    
    name      old time/op  new time/op  delta
    Split1-2  14.9ms ± 0%  14.2ms ± 0%  -4.06%  (p=0.000 n=47+49)
    
    Change-Id: I00f796bd2edab70e926ea9e65439d820c6a28254
    Reviewed-on: https://go-review.googlesource.com/21609
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3f66d8c84b4b3d685db1031954d3343a7a8c9d0f
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Tue Apr 12 01:55:14 2016 -0700

    html/template: add examples of loading templates from files
    
    Adds examples showing loading templates from files and
    executing them.
    
    Shows examples:
    - Using ParseGlob.
    - Using ParseFiles.
    - Using helper functions to share and use templates
    in different contexts by adding them to an existing
    bundle of templates.
    - Using a group of driver templates with distinct sets
    of helper templates.
    
    Almost all of the code was directly copied from text/template.
    
    Fixes #8500
    
    Change-Id: Ic3d91d5232afc5a1cd2d8cd3d9a5f3b754c64225
    Reviewed-on: https://go-review.googlesource.com/21854
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit db5338f87982086a19310ad6e25c046280644b98
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Apr 12 17:12:26 2016 -0700

    cmd/compile: teach CSE that new objects are bespoke
    
    runtime.newobject never returns the same thing twice,
    so the resulting value will never be a common subexpression.
    
    This helps when compiling large static data structures
    that include pointers, such as maps and slices.
    No clear performance impact on other code. (See below.)
    
    For the code in issue #15112:
    
    Before:
      real	1m14.238s
      user	1m18.985s
      sys	0m0.787s
    
    After:
      real	0m47.172s
      user	0m52.248s
      sys	0m0.767s
    
    For the code in issue #15235, size 10k:
    
    Before:
      real	0m44.916s
      user	0m46.577s
      sys	0m0.304s
    
    After:
      real	0m7.703s
      user	0m9.041s
      sys	0m0.316s
    
    Still more work to be done, particularly for #15112.
    
    Updates #15112
    Updates #15235
    
    
    name       old time/op      new time/op      delta
    Template        330ms ±11%       333ms ±13%    ~           (p=0.749 n=20+19)
    Unicode         148ms ± 6%       152ms ± 8%    ~           (p=0.072 n=18+20)
    GoTypes         1.01s ± 7%       1.01s ± 3%    ~           (p=0.583 n=20+20)
    Compiler        5.04s ± 2%       5.06s ± 2%    ~           (p=0.314 n=20+20)
    
    name       old user-ns/op   new user-ns/op   delta
    Template   444user-ms ±11%  445user-ms ±10%    ~           (p=0.738 n=20+20)
    Unicode    215user-ms ± 5%  218user-ms ± 5%    ~           (p=0.239 n=18+18)
    GoTypes    1.45user-s ± 3%  1.45user-s ± 4%    ~           (p=0.620 n=20+20)
    Compiler   7.23user-s ± 2%  7.22user-s ± 2%    ~           (p=0.901 n=20+19)
    
    name       old alloc/op     new alloc/op     delta
    Template       55.0MB ± 0%      55.0MB ± 0%    ~           (p=0.547 n=20+20)
    Unicode        37.6MB ± 0%      37.6MB ± 0%    ~           (p=0.301 n=20+20)
    GoTypes         177MB ± 0%       177MB ± 0%    ~           (p=0.065 n=20+19)
    Compiler        798MB ± 0%       797MB ± 0%  -0.05%        (p=0.000 n=19+20)
    
    name       old allocs/op    new allocs/op    delta
    Template         492k ± 0%        493k ± 0%  +0.03%        (p=0.030 n=20+20)
    Unicode          377k ± 0%        377k ± 0%    ~           (p=0.423 n=20+19)
    GoTypes         1.40M ± 0%       1.40M ± 0%    ~           (p=0.102 n=20+20)
    Compiler        5.53M ± 0%       5.53M ± 0%    ~           (p=0.094 n=17+18)
    
    name       old text-bytes   new text-bytes   delta
    HelloSize        561k ± 0%        561k ± 0%    ~     (all samples are equal)
    CmdGoSize       6.13M ± 0%       6.13M ± 0%    ~     (all samples are equal)
    
    name       old data-bytes   new data-bytes   delta
    HelloSize        128k ± 0%        128k ± 0%    ~     (all samples are equal)
    CmdGoSize        306k ± 0%        306k ± 0%    ~     (all samples are equal)
    
    name       old exe-bytes    new exe-bytes    delta
    HelloSize        905k ± 0%        905k ± 0%    ~     (all samples are equal)
    CmdGoSize       9.64M ± 0%       9.64M ± 0%    ~     (all samples are equal)
    
    Change-Id: Id774e2901d7701a3ec7a1c1d1cf1d9327a4107fc
    Reviewed-on: https://go-review.googlesource.com/21937
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Todd Neal <todd@tneal.org>

commit 260b7daf0a3fa1548d976f2484325240d4bdb73a
Author: Keith Randall <khr@golang.org>
Date:   Tue Apr 12 16:25:48 2016 -0700

    cmd/compile: fix arg to getcallerpc
    
    getcallerpc's arg needs to point to the first argument slot.
    I believe this bug was introduced by Michel's itab changes
    (specifically https://go-review.googlesource.com/c/20902).
    
    Fixes #15145
    
    Change-Id: Ifb2e17f3658e2136c7950dfc789b4d5706320683
    Reviewed-on: https://go-review.googlesource.com/21931
    Reviewed-by: Michel Lespinasse <walken@google.com>

commit b0cbe158da10aac1876680e825a902d58a9d1bac
Author: Shahar Kohanim <skohanim@gmail.com>
Date:   Mon Apr 11 22:19:34 2016 +0300

    cmd/link: move function only lsym fields to pcln struct
    
    name       old secs    new secs    delta
    LinkCmdGo   0.53 ± 9%   0.53 ±10%  -1.30%  (p=0.022 n=100+99)
    
    name       old MaxRSS  new MaxRSS  delta
    LinkCmdGo   151k ± 4%   142k ± 6%  -5.92%  (p=0.000 n=98+100)
    
    Change-Id: Ic30e63a948f8e626b3396f458a0163f7234810c1
    Reviewed-on: https://go-review.googlesource.com/21920
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a0ab6cd6852cec430e280217a9516d6be3c1ef5f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Apr 12 00:22:39 2016 +0000

    net/http: add test that panic in a handler signals an error to the client
    
    Change-Id: Iba40edc9ddad62534b06c5af20bbc3dd3dc14d0a
    Reviewed-on: https://go-review.googlesource.com/21881
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 982274c96d6c9ad88a9deb07583d3b74ec2df357
Author: Ian Lance Taylor <iant@golang.org>
Date:   Tue Apr 12 15:47:17 2016 -0700

    reflect: test that Call results are not addressable
    
    Gccgo was erroneously marking Call results as addressable, which led to
    an obscure bug using text/template, as text/template calls CanAddr to
    check whether to take the address of a value when looking up methods.
    When a function returned a pointer, and CanAddr was true, the result was
    a pointer to a pointer that had no methods.
    
    Fixed in gccgo by https://golang.org/cl/21908.  Adding the test here so
    that it doesn't regress.
    
    Change-Id: I1d25b868e1b8e2348b21cbac6404a636376d1a4a
    Reviewed-on: https://go-review.googlesource.com/21930
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f0c5b8b9c9c7900033ddb11b584da6198d599454
Author: Martin Möhrmann <martisch@uos.de>
Date:   Tue Apr 12 21:16:27 2016 +0200

    image/color: optimize YCbCrToRGB
    
    Use one comparison to detect underflow and overflow simultaneously.
    Use a shift, bitwise complement and uint8 type conversion to handle
    clamping to upper and lower bound without additional branching.
    
    Overall the new code is faster for a mix of
    common case, underflow and overflow.
    
    name     old time/op  new time/op  delta
    YCbCr-2  1.12ms ± 0%  0.64ms ± 0%  -43.01%  (p=0.000 n=48+47)
    
    name              old time/op  new time/op  delta
    YCbCrToRGB/0-2    5.52ns ± 0%  5.77ns ± 0%  +4.48%  (p=0.000 n=50+49)
    YCbCrToRGB/128-2  6.05ns ± 0%  5.52ns ± 0%  -8.69%  (p=0.000 n=39+50)
    YCbCrToRGB/255-2  5.80ns ± 0%  5.77ns ± 0%  -0.58%  (p=0.000 n=50+49)
    
    Found in collaboration with Josh Bleecher Snyder and Ralph Corderoy.
    
    Change-Id: Ic5020320f704966f545fdc1ae6bc24ddb5d3d09a
    Reviewed-on: https://go-review.googlesource.com/21910
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1650ced98f4f6c5f0783f78cb9d0ffd3a6d1768f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Apr 12 22:27:33 2016 +0000

    net: skip failing or flaky TestInterfaces on freebsd-arm
    
    Updates #15262
    
    Change-Id: I3eb1f6f71d6285d039f11ba6a34b8a599a33bf49
    Reviewed-on: https://go-review.googlesource.com/21909
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit f6db855d5595e170bfc70d90f1eaa26034d2e191
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Apr 12 21:51:19 2016 +0000

    net: mirror Tom Sawyer in the net package for tests
    
    Fixes the darwin/arm builder, which has a special test runner which
    makes the assumption that tests never use testdata from another
    package.
    
    This looks large, but it's no more space in git.
    
    Change-Id: I81921b516443d12d21b77617d323ddebedbe40f8
    Reviewed-on: https://go-review.googlesource.com/21907
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e46b00a43b62bd67ec13ca75c51037db3b312043
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Apr 12 14:44:49 2016 -0700

    cmd/internal/obj: remove unused Pciter type
    
    Change-Id: Ie8323cfcd1193f390729d0d3dd67863aedf47d13
    Reviewed-on: https://go-review.googlesource.com/21906
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6fb8bf9d7940b2e2a90249e1894ad4e3d24fd3e7
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Apr 12 14:04:01 2016 -0700

    net: make two tests not parallel
    
    Running
    
    stress -p 1 go test -short std
    
    on a heavily loaded machine causes net timeouts
    every 15 or 20 runs.
    Making these tests not run in parallel helps.
    With this change, I haven’t seen a single failure
    in over 100 runs.
    
    Fixes #14986
    
    Change-Id: Ibaa14869ce8d95b00266aee94d62d195927ede68
    Reviewed-on: https://go-review.googlesource.com/21905
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 022548cfe82915e5bf14ce7cb28f3ec651550662
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Tue Apr 12 16:58:56 2016 -0300

    all: standardize RFC mention format
    
    Standardize on space between "RFC" and number. Additionally change
    the couple "a RFC" instances to "an RFC."
    
    Fixes #15258
    
    Change-Id: I2b17ecd06be07dfbb4207c690f52a59ea9b04808
    Reviewed-on: https://go-review.googlesource.com/21902
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 24fc3234428e138e693584185fab4146de6088db
Author: Shahar Kohanim <skohanim@gmail.com>
Date:   Mon Apr 11 17:35:55 2016 +0300

    cmd/internal/obj: split plist flushing from object writing
    
    Only splits into separate files, no other changes.
    
    Change-Id: Icc0da2c5f18e03e9ed7c0043bd7c790f741900f2
    Reviewed-on: https://go-review.googlesource.com/21804
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f028b9f9e2433662502283850d06e9e07e72a6bb
Author: David Crawshaw <crawshaw@golang.org>
Date:   Sun Mar 27 10:21:48 2016 -0400

    cmd/link, etc: store typelinks as offsets
    
    This is the first in a series of CLs to replace the use of pointers
    in binary read-only data with offsets.
    
    In standard Go binaries these CLs have a small effect, shrinking
    8-byte pointers to 4-bytes. In position-independent code, it also
    saves the dynamic relocation for the pointer. This has a significant
    effect on the binary size when building as PIE, c-archive, or
    c-shared.
    
    darwin/amd64:
    	cmd/go: -12KB (0.1%)
    	jujud:  -82KB (0.1%)
    
    linux/amd64 PIE:
    	cmd/go:  -86KB (0.7%)
    	jujud:  -569KB (0.7%)
    
    For #6853.
    
    Change-Id: Iad5625bbeba58dabfd4d334dbee3fcbfe04b2dcf
    Reviewed-on: https://go-review.googlesource.com/21284
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 23cbfa2545a735eca5ffc1ffd6c0e93c2eecac2a
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Apr 12 19:48:01 2016 +0000

    net: skip TestDialCancel on linux-arm64-buildlet
    
    These builders (on Linaro) have a different network configuration
    which is incompatible with this test. Or so it seems.
    
    Updates #15191
    
    Change-Id: Ibfeacddc98dac1da316e704b5c8491617a13e3bf
    Reviewed-on: https://go-review.googlesource.com/21901
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit da224a5c42e7fce7f1d190a86962b1c46be454ef
Author: Alberto Donizetti <alb.donizetti@gmail.com>
Date:   Sun Apr 10 20:14:27 2016 +0200

    cmd/pprof: pass the event to pprof_toggle_asm for the weblist command
    
    Fixes #15225
    
    Change-Id: I1f85590b2c3293463c6476beebcd3256adc1bf23
    Reviewed-on: https://go-review.googlesource.com/21802
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 37af06360039c96be707526596557a33885c3ad0
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Tue Apr 12 13:12:54 2016 -0300

    crypto/x509: remove broken link in ParsePKCS8PrivateKey documentation
    
    Fixes #14776
    
    Change-Id: I55423ac643f18542b9fd1386ed98dec47fb678aa
    Reviewed-on: https://go-review.googlesource.com/21890
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e07a4459a155789fb57bbf4e2c8eaca5b369fd17
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Tue Apr 12 12:16:20 2016 -0700

    cmd: replace x[i:][0] expressions with x[i]
    
    Passes toolstash -cmp.
    
    Change-Id: Id504e71ed1f23900e24a9aed25143c94f4d7d50c
    Reviewed-on: https://go-review.googlesource.com/21899
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 98080a6c64c2d9bc2a759b66a9b861af4ef7367b
Author: Shawn Walker-Salas <shawn.walker@oracle.com>
Date:   Fri Apr 8 15:59:04 2016 -0700

    net: broken sendfile on SmartOS/Solaris
    
    In the event of a partial write on Solaris and some BSDs, the offset
    pointer passed to sendfile() will be updated even though the function
    returns -1 if errno is set to EAGAIN/EINTR.  In that case, calculate the
    bytes written based on the difference between the updated offset and the
    original offset.  If no bytes were written, and errno is set to
    EAGAIN/EINTR, ignore the errno.
    
    Fixes #13892
    
    Change-Id: I6334b5ef2edcbebdaa7db36fa4f7785967313c2d
    Reviewed-on: https://go-review.googlesource.com/21769
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b1851a3c11a179d4eb55f9d0dd25ef81668a9f81
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Apr 12 11:31:16 2016 -0700

    cmd/compile: move compiler-specific flags into compiler-spec. export data section
    
    Also: Adjust go/importer accordingly.
    
    Change-Id: Ia6669563793e218946af45b9fba1cf986a21c031
    Reviewed-on: https://go-review.googlesource.com/21896
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit be7c786dd04db51076012618ea29ee528a654978
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Apr 12 13:38:24 2016 -0400

    cmd/objdump: skip TestDisasm* on s390x
    
    The disassembler is not yet implemented on s390x.
    
    Updates #15255.
    
    Change-Id: Ibab319c8c087b1a619baa1529398305c1e721877
    Reviewed-on: https://go-review.googlesource.com/21894
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b09c274bfabb3edef60b4df3375906852aab7da1
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Apr 12 12:46:54 2016 -0400

    net/http: fix TestLinuxSendfile on s390x
    
    s390x doesn't have sendfile64 so apply the same fix as MIPS
    (eebf7d27) and just use sendfile.
    
    Change-Id: If8fe2e974ed44a9883282430157c3545d5bd04bd
    Reviewed-on: https://go-review.googlesource.com/21892
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 613ba6cda845fef442995d705027a622984c6b3a
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Apr 12 12:26:17 2016 -0400

    cmd/compile/internal/gc: add s390x support
    
    Allows instructions with a From3 field to be used in regopt so
    long as From3 represents a constant. This is needed because the
    storage-to-storage instructions on s390x place the length of the
    data into From3.
    
    Change-Id: I12cd32d4f997baf2fe97937bb7d45bbf716dfcb5
    Reviewed-on: https://go-review.googlesource.com/20875
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>

commit 811ebb6ac961162b815f4fd50976df81ba4c47b0
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Apr 12 09:22:26 2016 -0700

    cmd/compile: temporarily disable inplace append special case
    
    Fixes #15246
    Re-opens #14969
    
    Change-Id: Ic0b41c5aa42bbb229a0d62b7f3e5888c6b29293d
    Reviewed-on: https://go-review.googlesource.com/21891
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 8edf4cb27d07a81ae340b0fda4e519c12f139618
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Apr 8 13:30:41 2016 -0400

    hash/crc32: invert build tags for go implementation
    
    It seems cleaner and more consistent with other files to list the
    architectures that have assembly implementations rather than to
    list those that do not.
    
    This means we don't have to add s390x and future platforms to this
    list.
    
    Change-Id: I2ad3f66b76eb1711333c910236ca7f5151b698e5
    Reviewed-on: https://go-review.googlesource.com/21770
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 78ecd61f6245197f701629f5f511be7f2bc1ff58
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Mar 18 19:20:34 2016 -0400

    runtime/cgo: add s390x support
    
    Change-Id: I64ada9fe34c3cfc4bd514ec5d8c8f4d4c99074fb
    Reviewed-on: https://go-review.googlesource.com/20950
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 921b2eba52906fc8b9bc4a8744dab63678f5ed3a
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Mar 21 13:30:50 2016 -0400

    debug/gosym: accept PC quantum of 2 (for s390x)
    
    Needed for the header check to accept the header generated for
    s390x as Go 1.2 style rather than Go 1.1 style.
    
    Change-Id: I7b3713d4cc7514cfc58f947a45702348f6d7b824
    Reviewed-on: https://go-review.googlesource.com/20966
    Reviewed-by: Minux Ma <minux@golang.org>

commit 7cbe7b1e867db9001db35ca41ee3e4a3b0de31c7
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Mar 18 19:13:59 2016 -0400

    runtime/internal/atomic: add s390x atomic operations
    
    Load and store instructions are atomic on the s390x.
    
    Change-Id: I0031ed2fba43f33863bca114d0fdec2e7d1ce807
    Reviewed-on: https://go-review.googlesource.com/20938
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7f5a063d157c777d8e78a567fc9538929bfd38f5
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Apr 12 10:27:16 2016 -0400

    cmd/compile/internal/gc: minor Cgen_checknil cleanup
    
    Most architectures can only generate nil checks when the
    the address to check is in a register. Currently only
    amd64 and 386 can generate checks for addresses that
    reside in memory. This is unlikely to change so the architecture
    check has been inverted.
    
    Change-Id: I73697488a183406c79a9039c62823712b510bb6a
    Reviewed-on: https://go-review.googlesource.com/21861
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cd85f711c0b6847cbfe4e05f4402df075ea936de
Author: Keith Randall <khr@golang.org>
Date:   Mon Apr 11 21:23:11 2016 -0700

    cmd/compile: add x.Uses==1 test to load combiners
    
    We need to make sure that when we combine loads, we only do
    so if there are no other uses of the load.  We can't split
    one load into two because that can then lead to inconsistent
    loaded values in the presence of races.
    
    Add some aggressive copy removal code so that phantom
    "dead copy" uses of values are cleaned up promptly.  This lets
    us use x.Uses==1 conditions reliably.
    
    Change-Id: I9037311db85665f3868dbeb3adb3de5c20728b38
    Reviewed-on: https://go-review.googlesource.com/21853
    Reviewed-by: Todd Neal <todd@tneal.org>

commit 204b6f48c5107d3132033324fd492ca0253568dc
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Tue Apr 12 09:41:11 2016 +0200

    cmd/pprof/internal: move to cmd/internal/pprof
    
    Make internal pprof packages available to cmd/trace.
    cmd/trace needs access to them to generate symbolized
    svg profiles (create and serialize Profile struct).
    And potentially generate svg programmatically instead
    of invoking go tool pprof.
    
    Change-Id: Iafd0c87ffdd4ddc081093be0b39761f19507907a
    Reviewed-on: https://go-review.googlesource.com/21870
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 9743e4b0311c37ebacc2c9063a1cd778510eae09
Author: Alexandru Moșoi <mosoi@google.com>
Date:   Mon Apr 11 21:51:29 2016 +0200

    cmd/compile: share dominator tree among many passes
    
    These passes do not modify the dominator tree too much.
    
    % benchstat old.txt new.txt
    name       old time/op     new time/op     delta
    Template       335ms ± 3%      325ms ± 8%    ~             (p=0.074 n=8+9)
    GoTypes        1.05s ± 1%      1.05s ± 3%    ~            (p=0.095 n=9+10)
    Compiler       5.37s ± 4%      5.29s ± 1%  -1.42%         (p=0.022 n=9+10)
    MakeBash       34.9s ± 3%      34.4s ± 2%    ~            (p=0.095 n=9+10)
    
    name       old alloc/op    new alloc/op    delta
    Template      55.4MB ± 0%     54.9MB ± 0%  -0.81%        (p=0.000 n=10+10)
    GoTypes        179MB ± 0%      178MB ± 0%  -0.89%        (p=0.000 n=10+10)
    Compiler       807MB ± 0%      798MB ± 0%  -1.10%        (p=0.000 n=10+10)
    
    name       old allocs/op   new allocs/op   delta
    Template        498k ± 0%       496k ± 0%  -0.29%          (p=0.000 n=9+9)
    GoTypes        1.42M ± 0%      1.41M ± 0%  -0.24%        (p=0.000 n=10+10)
    Compiler       5.61M ± 0%      5.60M ± 0%  -0.12%        (p=0.000 n=10+10)
    
    Change-Id: I4cd20cfba3f132ebf371e16046ab14d7e42799ec
    Reviewed-on: https://go-review.googlesource.com/21806
    Run-TryBot: Alexandru Moșoi <alexandru@mosoi.ro>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit a223ccae386449169774597b15a00f2d70addce7
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Apr 11 20:23:19 2016 -0400

    cmd/compile/internal/s390x: add s390x support
    
    s390x does not require duffzero/duffcopy since it has
    storage-to-storage instructions that can copy/clear up to 256
    bytes at a time.
    
    peep contains several new passes to optimize instruction
    sequences that match s390x instructions such as the
    compare-and-branch and load/store multiple instructions.
    
    copyprop and subprop have been extended to work with moves that
    require sign/zero extension. This work could be ported to other
    architectures that do not used sized math however it does add
    complexity and will probably be rendered unnecessary by ssa in
    the near future.
    
    Change-Id: I1b64b281b452ed82a85655a0df69cb224d2a6941
    Reviewed-on: https://go-review.googlesource.com/20873
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 944a0859b9a16a1951512b82870a31f371d1c417
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri Apr 8 11:53:23 2016 +0200

    internal/trace: fix int overflow in timestamps
    
    Fixes #15102
    
    Change-Id: I7fdb6464afd0b7af9b6652051416f0fddd34dc9a
    Reviewed-on: https://go-review.googlesource.com/21730
    Reviewed-by: Austin Clements <austin@google.com>

commit b6cd6d7d3211bd9030dec4115b6202d93fe570a3
Author: Dmitri Shuralyov <shurcooL@gmail.com>
Date:   Thu Mar 31 02:01:48 2016 -0700

    cmd/go: fix vcsFromDir returning bad root on Windows
    
    Apply golang/tools@5804fef4c0556d3e5e7d2440740a3d7aced7d58b.
    
    In the context of cmd/go build tool, import path is a '/'-separated path.
    This can be inferred from `go help importpath` and `go help packages`.
    vcsFromDir documentation says on return, root is the import path
    corresponding to the root of the repository. On Windows and other
    OSes where os.PathSeparator is not '/', that wasn't true since root
    would contain characters other than '/', and therefore it wasn't a
    valid import path corresponding to the root of the repository.
    Fix that by using filepath.ToSlash.
    
    Add test coverage for vcsFromDir, it was previously not tested.
    It's taken from golang.org/x/tools/go/vcs tests, and modified to
    improve style.
    
    Additionally, remove an unneccessary statement from the documentation
    "(thus root is a prefix of importPath)". There is no variable
    importPath that is being referred to (it's possible p.ImportPath
    was being referred to). Without it, the description of root value
    matches the documentation of repoRoot.root struct field:
    
    	// root is the import path corresponding to the root of the
    	// repository
    	root string
    
    Rename and change signature of vcsForDir(p *Package) to
    vcsFromDir(dir, srcRoot string). This is more in sync with the x/tools
    version. It's also simpler, since vcsFromDir only needs those two
    values from Package, and nothing more. Change "for" to "from" in name
    because it's more consistent and clear.
    
    Update usage of vcsFromDir to match the new signature, and respect
    that returned root is a '/'-separated path rather than a os.PathSeparator
    separated path.
    
    Fixes #15040.
    Updates #7723.
    Helps #11490.
    
    Change-Id: Idf51b9239f57248739daaa200aa1c6e633cb5f7f
    Reviewed-on: https://go-review.googlesource.com/21345
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>
    Run-TryBot: Alex Brainman <alex.brainman@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 07669d2737aa51107a4e54b61d6704f6ad8035b5
Author: Martin Möhrmann <martisch@uos.de>
Date:   Thu Apr 7 08:01:47 2016 +0200

    cmd/compile: cleanup pragcgo
    
    Removes dynimport, dynexport, dynlinker cases since they can not
    be reached due to prefix check for "go:cgo_" in getlinepragma.
    
    Replaces the if chains for verb distinction by a switch statement.
    Replaces fmt.Sprintf by fmt.Sprintln for string concatenation.
    
    Removes the more, getimpsym and getquoted functions by introducing a
    pragmaFields function that partitions a pragma into its components.
    
    Adds tests for cgo pragmas.
    
    Change-Id: I43c7b9550feb3ddccaff7fb02198a3f994444123
    Reviewed-on: https://go-review.googlesource.com/21607
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e569b10ebaed8fbf27d0b55886b6a81d635ddbc7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Apr 12 00:09:05 2016 +0000

    src: split nacltest.bash into naclmake.bash and keep nacltest.bash
    
    Needed by the build system to shard tests. nacl was the last unsharded
    builder.
    
    (I considered also adding a -make-only flag to nacltest.bash, but that
    wouldn't fail fast when the file didn't exist.)
    
    Updates #15242
    
    Change-Id: I6afc1c1fe4268ab98c0724b5764c67d3784caebe
    Reviewed-on: https://go-review.googlesource.com/21851
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 73a0185ad390761f21a3407858fcccc6a11c0858
Author: David Benjamin <davidben@google.com>
Date:   Mon Feb 15 11:51:54 2016 -0500

    crypto/tls: Enforce that version and cipher match on resume.
    
    Per RFC 5246, 7.4.1.3:
    
       cipher_suite
          The single cipher suite selected by the server from the list in
          ClientHello.cipher_suites.  For resumed sessions, this field is
          the value from the state of the session being resumed.
    
    The specifications are not very clearly written about resuming sessions
    at the wrong version (i.e. is the TLS 1.0 notion of "session" the same
    type as the TLS 1.1 notion of "session"?). But every other
    implementation enforces this check and not doing so has some odd
    semantics.
    
    Change-Id: I6234708bd02b636c25139d83b0d35381167e5cad
    Reviewed-on: https://go-review.googlesource.com/21153
    Reviewed-by: Adam Langley <agl@golang.org>

commit c9b66bb355ebbc6a26ee511e996cba4da3e1d644
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Mon Apr 11 21:47:37 2016 -0300

    io: document WriteString calls Write exactly once
    
    Fixes #13849
    
    Change-Id: Idd7f06b547a0179fe15571807a8c48b7c3b78d7c
    Reviewed-on: https://go-review.googlesource.com/21852
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 1f54410a61e2242285e366a5580943f78fbff741
Author: Mikio Hara <mikioh.mikioh@gmail.com>
Date:   Thu Apr 7 17:19:29 2016 +0900

    net: make IP.{String,MarshalText} return helpful information on address error
    
    This change makes String and MarshalText methods of IP return a
    hexadecial form of IP with no punctuation as part of error
    notification. It doesn't affect the existing behavior of ParseIP.
    
    Also fixes bad shadowing in ipToSockaddr and makes use of reserved
    IP address blocks for documnetation.
    
    Fixes #15052.
    Updates #15228.
    
    Change-Id: I9e9ecce308952ed5683066c3d1bb6a7b36458c65
    Reviewed-on: https://go-review.googlesource.com/21642
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a119f88f2c52e6eb969f51c5bf610d6f105348a3
Author: Alan Donovan <adonovan@google.com>
Date:   Mon Apr 11 18:39:46 2016 -0400

    go/importer: make For("gccgo",  nil) not panic
    
    Apparently we forgot to test this.
    
    Fixes #15092
    
    Change-Id: I33d4fef0f659dfbdfc1ebf8401e96610c8215592
    Reviewed-on: https://go-review.googlesource.com/21860
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit 501ddf7189cf97ef27eb870ad134a312f80ae585
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Apr 11 18:16:55 2016 +0000

    context: attempt to deflake timing tests
    
    Passes on OpenBSD now when running it with -count=500.
    
    Presumably this will also fix the same problems seen on FreeBSD and
    Windows.
    
    Fixes #15158
    
    Change-Id: I86451c901613dfa5ecff0c2ecc516527a3c011b3
    Reviewed-on: https://go-review.googlesource.com/21840
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit d1feddb7ae6389ee4175ee85b7168cf58a04b952
Author: Rob Pike <r@golang.org>
Date:   Mon Apr 4 13:22:34 2016 -0700

    cmd/vet: improve documentation for flags, slightly
    
    The way that -all works was unclear from the documentation and made
    worse by recent changes to the flag package. Improve matters by making
    the help message say "default true" for the tests that do default to true,
    and tweak some of the wording.
    
    Before:
    
    Usage of vet:
    	vet [flags] directory...
    	vet [flags] files... # Must be a single package
    For more information run
    	go doc cmd/vet
    
    Flags:
      -all
        	enable all non-experimental checks (default unset)
      -asmdecl
        	check assembly against Go declarations (default unset)
    ...
    
    After:
    
    Usage of vet:
    	vet [flags] directory...
    	vet [flags] files... # Must be a single package
    By default, -all is set and all non-experimental checks are run.
    For more information run
    	go doc cmd/vet
    
    Flags:
      -all
        	enable all non-experimental checks (default true)
      -asmdecl
        	check assembly against Go declarations (default true)
    ...
    
    Change-Id: Ie94b27381a9ad2382a10a7542a93bce1d59fa8f5
    Reviewed-on: https://go-review.googlesource.com/21495
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit a4650a2111b2bb826ca64a13bdad9c96e3095e47
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Apr 10 09:44:17 2016 -0700

    cmd/compile: avoid write barrier in append fast path
    
    When we are writing the result of an append back
    to the same slice, we don’t need a write barrier
    on the fast path.
    
    This re-implements an optimization that was present
    in the old backend.
    
    Updates #14921
    Fixes #14969
    
    Sample code:
    
    var x []byte
    
    func p() {
    	x = append(x, 1, 2, 3)
    }
    
    Before:
    
    "".p t=1 size=224 args=0x0 locals=0x48
    	0x0000 00000 (append.go:21)	TEXT	"".p(SB), $72-0
    	0x0000 00000 (append.go:21)	MOVQ	(TLS), CX
    	0x0009 00009 (append.go:21)	CMPQ	SP, 16(CX)
    	0x000d 00013 (append.go:21)	JLS	199
    	0x0013 00019 (append.go:21)	SUBQ	$72, SP
    	0x0017 00023 (append.go:21)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    	0x0017 00023 (append.go:21)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    	0x0017 00023 (append.go:19)	MOVQ	"".x+16(SB), CX
    	0x001e 00030 (append.go:19)	MOVQ	"".x(SB), DX
    	0x0025 00037 (append.go:19)	MOVQ	"".x+8(SB), BX
    	0x002c 00044 (append.go:19)	MOVQ	BX, "".autotmp_0+64(SP)
    	0x0031 00049 (append.go:22)	LEAQ	3(BX), BP
    	0x0035 00053 (append.go:22)	CMPQ	BP, CX
    	0x0038 00056 (append.go:22)	JGT	$0, 131
    	0x003a 00058 (append.go:22)	MOVB	$1, (DX)(BX*1)
    	0x003e 00062 (append.go:22)	MOVB	$2, 1(DX)(BX*1)
    	0x0043 00067 (append.go:22)	MOVB	$3, 2(DX)(BX*1)
    	0x0048 00072 (append.go:22)	MOVQ	BP, "".x+8(SB)
    	0x004f 00079 (append.go:22)	MOVQ	CX, "".x+16(SB)
    	0x0056 00086 (append.go:22)	MOVL	runtime.writeBarrier(SB), AX
    	0x005c 00092 (append.go:22)	TESTB	AL, AL
    	0x005e 00094 (append.go:22)	JNE	$0, 108
    	0x0060 00096 (append.go:22)	MOVQ	DX, "".x(SB)
    	0x0067 00103 (append.go:23)	ADDQ	$72, SP
    	0x006b 00107 (append.go:23)	RET
    	0x006c 00108 (append.go:22)	LEAQ	"".x(SB), CX
    	0x0073 00115 (append.go:22)	MOVQ	CX, (SP)
    	0x0077 00119 (append.go:22)	MOVQ	DX, 8(SP)
    	0x007c 00124 (append.go:22)	PCDATA	$0, $0
    	0x007c 00124 (append.go:22)	CALL	runtime.writebarrierptr(SB)
    	0x0081 00129 (append.go:23)	JMP	103
    	0x0083 00131 (append.go:22)	LEAQ	type.[]uint8(SB), AX
    	0x008a 00138 (append.go:22)	MOVQ	AX, (SP)
    	0x008e 00142 (append.go:22)	MOVQ	DX, 8(SP)
    	0x0093 00147 (append.go:22)	MOVQ	BX, 16(SP)
    	0x0098 00152 (append.go:22)	MOVQ	CX, 24(SP)
    	0x009d 00157 (append.go:22)	MOVQ	BP, 32(SP)
    	0x00a2 00162 (append.go:22)	PCDATA	$0, $0
    	0x00a2 00162 (append.go:22)	CALL	runtime.growslice(SB)
    	0x00a7 00167 (append.go:22)	MOVQ	40(SP), DX
    	0x00ac 00172 (append.go:22)	MOVQ	48(SP), AX
    	0x00b1 00177 (append.go:22)	MOVQ	56(SP), CX
    	0x00b6 00182 (append.go:22)	ADDQ	$3, AX
    	0x00ba 00186 (append.go:19)	MOVQ	"".autotmp_0+64(SP), BX
    	0x00bf 00191 (append.go:22)	MOVQ	AX, BP
    	0x00c2 00194 (append.go:22)	JMP	58
    	0x00c7 00199 (append.go:22)	NOP
    	0x00c7 00199 (append.go:21)	CALL	runtime.morestack_noctxt(SB)
    	0x00cc 00204 (append.go:21)	JMP	0
    
    After:
    
    "".p t=1 size=208 args=0x0 locals=0x48
    	0x0000 00000 (append.go:21)	TEXT	"".p(SB), $72-0
    	0x0000 00000 (append.go:21)	MOVQ	(TLS), CX
    	0x0009 00009 (append.go:21)	CMPQ	SP, 16(CX)
    	0x000d 00013 (append.go:21)	JLS	191
    	0x0013 00019 (append.go:21)	SUBQ	$72, SP
    	0x0017 00023 (append.go:21)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    	0x0017 00023 (append.go:21)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    	0x0017 00023 (append.go:19)	MOVQ	"".x+16(SB), CX
    	0x001e 00030 (append.go:19)	MOVQ	"".x+8(SB), DX
    	0x0025 00037 (append.go:19)	MOVQ	DX, "".autotmp_0+64(SP)
    	0x002a 00042 (append.go:19)	MOVQ	"".x(SB), BX
    	0x0031 00049 (append.go:22)	LEAQ	3(DX), BP
    	0x0035 00053 (append.go:22)	MOVQ	BP, "".x+8(SB)
    	0x003c 00060 (append.go:22)	CMPQ	BP, CX
    	0x003f 00063 (append.go:22)	JGT	$0, 84
    	0x0041 00065 (append.go:22)	MOVB	$1, (BX)(DX*1)
    	0x0045 00069 (append.go:22)	MOVB	$2, 1(BX)(DX*1)
    	0x004a 00074 (append.go:22)	MOVB	$3, 2(BX)(DX*1)
    	0x004f 00079 (append.go:23)	ADDQ	$72, SP
    	0x0053 00083 (append.go:23)	RET
    	0x0054 00084 (append.go:22)	LEAQ	type.[]uint8(SB), AX
    	0x005b 00091 (append.go:22)	MOVQ	AX, (SP)
    	0x005f 00095 (append.go:22)	MOVQ	BX, 8(SP)
    	0x0064 00100 (append.go:22)	MOVQ	DX, 16(SP)
    	0x0069 00105 (append.go:22)	MOVQ	CX, 24(SP)
    	0x006e 00110 (append.go:22)	MOVQ	BP, 32(SP)
    	0x0073 00115 (append.go:22)	PCDATA	$0, $0
    	0x0073 00115 (append.go:22)	CALL	runtime.growslice(SB)
    	0x0078 00120 (append.go:22)	MOVQ	40(SP), CX
    	0x007d 00125 (append.go:22)	MOVQ	56(SP), AX
    	0x0082 00130 (append.go:22)	MOVQ	AX, "".x+16(SB)
    	0x0089 00137 (append.go:22)	MOVL	runtime.writeBarrier(SB), AX
    	0x008f 00143 (append.go:22)	TESTB	AL, AL
    	0x0091 00145 (append.go:22)	JNE	$0, 168
    	0x0093 00147 (append.go:22)	MOVQ	CX, "".x(SB)
    	0x009a 00154 (append.go:22)	MOVQ	"".x(SB), BX
    	0x00a1 00161 (append.go:19)	MOVQ	"".autotmp_0+64(SP), DX
    	0x00a6 00166 (append.go:22)	JMP	65
    	0x00a8 00168 (append.go:22)	LEAQ	"".x(SB), DX
    	0x00af 00175 (append.go:22)	MOVQ	DX, (SP)
    	0x00b3 00179 (append.go:22)	MOVQ	CX, 8(SP)
    	0x00b8 00184 (append.go:22)	PCDATA	$0, $0
    	0x00b8 00184 (append.go:22)	CALL	runtime.writebarrierptr(SB)
    	0x00bd 00189 (append.go:22)	JMP	154
    	0x00bf 00191 (append.go:22)	NOP
    	0x00bf 00191 (append.go:21)	CALL	runtime.morestack_noctxt(SB)
    	0x00c4 00196 (append.go:21)	JMP	0
    
    Change-Id: I77a41ad3a22557a4bb4654de7d6d24a029efe34a
    Reviewed-on: https://go-review.googlesource.com/21813
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 32efa16c3d63dd630e2190a8c0f30c0a941f6fd7
Author: David Chase <drchase@google.com>
Date:   Fri Apr 1 14:51:29 2016 -0400

    cmd/compile: added stats printing to stackalloc
    
    This is controlled by the "regalloc" stats flag, since regalloc
    calls stackalloc.  The plan is for this to allow comparison
    of cheaper stack allocation algorithms with what we have now.
    
    Change-Id: Ibf64a780344c69babfcbb328fd6d053ea2e02cfc
    Reviewed-on: https://go-review.googlesource.com/21393
    Run-TryBot: David Chase <drchase@google.com>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 7e40627a0e595aa321efaf44f8507b678ee5eb1e
Author: Keith Randall <khr@golang.org>
Date:   Mon Apr 11 13:17:52 2016 -0700

    cmd/compile: zero all three argstorage slots
    
    These changes were missed when going from 2 to 3 argstorage slots.
    https://go-review.googlesource.com/20296/
    
    Change-Id: I930a307bb0b695bf1ae088030c9bbb6d14ca31d2
    Reviewed-on: https://go-review.googlesource.com/21841
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 7f53391f6b7f2387a5ed00398d34b046c321966f
Author: Keith Randall <khr@golang.org>
Date:   Mon Apr 11 12:22:26 2016 -0700

    cmd/compile: fix -N build
    
    The decomposer of builtin types is confused by having structs
    still around from the user-type decomposer.  They're all dead though,
    so just enabling a deadcode pass fixes things.
    
    Change-Id: I2df6bc7e829be03eabfd24c8dda1bff96f3d7091
    Reviewed-on: https://go-review.googlesource.com/21839
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 7166dfe0c11bd25b962d7f691ea1be857842dfbf
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Mon Apr 11 11:54:07 2016 -0700

    image/color: add YCbCrToRGB benchmark
    
    Change-Id: I9ba76d5b0861a901415fdceccaf2f5caa2facb7f
    Reviewed-on: https://go-review.googlesource.com/21837
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit cd6b2b7451c6feb277d38820f41f81ce4a036af2
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Apr 10 20:01:49 2016 -0400

    cmd/internal/obj/s390x: add MULHD instruction
    
    Emulate 64-bit signed high multiplication ((a*b)>>64). To do this
    we use the 64-bit unsigned high multiplication method and then
    fix the result as shown in Hacker's Delight 2nd ed., chapter 8-3.
    
    Required to enable some division optimizations.
    
    Change-Id: I9194f428e09d3d029cb1afb4715cd5424b5d922e
    Reviewed-on: https://go-review.googlesource.com/21774
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit cabb1402568ae7d05d9d5adf56953a4792085a81
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Apr 11 05:12:43 2016 +0000

    net/http: add ServerContextKey to let a handler access its Server
    
    Fixes #12438
    Updates #15229 (to decide context key variable naming convention)
    
    Change-Id: I3ba423e91b689e232143247d044495a12c97a7d2
    Reviewed-on: https://go-review.googlesource.com/21829
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e79fef8e55f8a893c65f41566bbec10339d45dec
Author: Shawn Walker-Salas <shawn.walker@oracle.com>
Date:   Thu Apr 7 15:26:57 2016 -0700

    cmd/link: external linking can fail on Solaris 11.2+
    
    Workaround external linking issues encountered on Solaris 11.2+ due to
    the go.o object file being created with a NULL STT_FILE symtab entry by
    using a placeholder name.
    
    Fixes #14957
    
    Change-Id: I89c501b4c548469f3c878151947d35588057982b
    Reviewed-on: https://go-review.googlesource.com/21636
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 3fafe2e8888dadb6877fa1e7569f5bd1f688dd3a
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Mon Apr 11 08:57:52 2016 +0200

    internal/trace: support parsing of 1.5 traces
    
    1. Parse out version from trace header.
    2. Restore handling of 1.5 traces.
    3. Restore optional symbolization of traces.
    4. Add some canned 1.5 traces for regression testing
       (http benchmark trace, runtime/trace stress traces,
        plus one with broken timestamps).
    
    Change-Id: Idb18a001d03ded8e13c2730eeeb37c5836e31256
    Reviewed-on: https://go-review.googlesource.com/21803
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Austin Clements <austin@google.com>

commit b04e145248d5d3721a41d4bb26704fdb43caaf38
Author: Keith Randall <khr@golang.org>
Date:   Thu Mar 31 21:24:10 2016 -0700

    cmd/compile: fix naming of decomposed structs
    
    When a struct is SSAable, we will name its component parts
    by their field names.  For example,
    type T struct {
         a, b, c int
    }
    If we ever need to spill a variable x of type T, we will
    spill its individual components to variables named x.a, x.b,
    and x.c.
    
    Change-Id: I857286ff1f2597f2c4bbd7b4c0b936386fb37131
    Reviewed-on: https://go-review.googlesource.com/21389
    Reviewed-by: David Chase <drchase@google.com>

commit 6c6089b3fdba9eb0cff863a03074dbac47c92f63
Author: Alexandru Moșoi <mosoi@google.com>
Date:   Fri Apr 1 15:09:19 2016 +0200

    cmd/compile: bce when max and limit are consts
    
    Removes 49 more bound checks in make.bash. For example:
    
    var a[100]int
    for i := 0; i < 50; i++ {
      use a[i+25]
    }
    
    Change-Id: I85e0130ee5d07f0ece9b17044bba1a2047414ce7
    Reviewed-on: https://go-review.googlesource.com/21379
    Reviewed-by: David Chase <drchase@google.com>
    Run-TryBot: Alexandru Moșoi <alexandru@mosoi.ro>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 00681eec6aec03b8b2822c9220fba27c18923c01
Author: Dan Peterson <dpiddy@gmail.com>
Date:   Mon Apr 11 11:15:00 2016 -0300

    net/http: document Error does not end the request
    
    Fixes #15205
    
    Change-Id: Ia650806756758ca8ed2272b1696e59b809b16c61
    Reviewed-on: https://go-review.googlesource.com/21836
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 683917a72154e3409e1ab5ef5b26030388312d0b
Author: Dominik Honnef <dominik@honnef.co>
Date:   Fri Apr 1 07:34:18 2016 +0200

    all: use bytes.Equal, bytes.Contains and strings.Contains, again
    
    The previous cleanup was done with a buggy tool, missing some potential
    rewrites.
    
    Change-Id: I333467036e355f999a6a493e8de87e084f374e26
    Reviewed-on: https://go-review.googlesource.com/21378
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4f12cc08132f3e5d2ba4b756c91d88c2e58a73b1
Author: Shahar Kohanim <skohanim@gmail.com>
Date:   Thu Apr 7 18:00:57 2016 +0300

    cmd/link: symbol generation optimizations
    
    After making dwarf generation backed by LSyms there was a performance regression
    of about 10%. These changes make on the fly symbol generation faster and
    are meant to help mitigate that.
    
    name       old secs    new secs    delta
    LinkCmdGo   0.55 ± 9%   0.53 ± 8%  -4.42%   (p=0.000 n=100+99)
    
    name       old MaxRSS  new MaxRSS  delta
    LinkCmdGo   152k ± 6%   149k ± 3%  -1.99%    (p=0.000 n=99+97)
    
    Change-Id: Iacca3ec924ce401aa83126bc0b10fe89bedf0ba6
    Reviewed-on: https://go-review.googlesource.com/21733
    Run-TryBot: Shahar Kohanim <skohanim@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit b0eeb8b0aaaf4997c25e3048bfc40e53d556a8eb
Author: Russ Cox <rsc@golang.org>
Date:   Sun Apr 3 12:52:12 2016 -0400

    net/http/pprof: accept fractional seconds in trace handler
    
    For heavily loaded servers, even 1 second of trace is too large
    to process with the trace viewer; using a float64 here allows
    fetching /debug/pprof/trace?seconds=0.1.
    
    Change-Id: I286c07abf04f9c1fe594b0e26799bf37f5c734db
    Reviewed-on: https://go-review.googlesource.com/21455
    Reviewed-by: Austin Clements <austin@google.com>

commit 720c4c016c75d37d14e0621696127819c8a73b0b
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Apr 8 17:50:40 2016 +1000

    runtime: merge lfstack_amd64.go into lfstack_64bit.go
    
    Merge the amd64 lfstack implementation into the general 64 bit
    implementation.
    
    Change-Id: Id9ed61b90d2e3bc3b0246294c03eb2c92803b6ca
    Reviewed-on: https://go-review.googlesource.com/21707
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Minux Ma <minux@golang.org>

commit 20375f64b18e9f904302d8f873e23702117bf4f5
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon Apr 11 11:04:15 2016 +1000

    cmd/go: document that -run=^$ skips all tests
    
    Change-Id: I7bbdd9600e0d9a647aeea16f1ae9e42a4e0cf44d
    Reviewed-on: https://go-review.googlesource.com/21823
    Reviewed-by: Rob Pike <r@golang.org>

commit 2a4158207edb499f8b210aaa7a9af103b93b5ac7
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Apr 10 21:58:37 2016 -0400

    cmd/compile/internal/gc: refactor cgen_div
    
    This commit adds two new functions to cgen.go: hasHMUL64 and
    hasRROTC64. These are used to determine whether or not an
    architecture supports the instructions needed to perform an
    optimization in cgen_div.
    
    This commit should not affect existing architectures (although it
    does add s390x to the new functions). However, since most
    architectures support HMUL the hasHMUL64 function could be
    modified to enable most of the optimizations in cgen_div on those
    platforms.
    
    Change-Id: I33bf329ddeb6cf2954bd17b7c161012de352fb62
    Reviewed-on: https://go-review.googlesource.com/21775
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a00ad5f47ea9c950dd041b33476c2bb4568496e1
Author: Andrew Gerrand <adg@golang.org>
Date:   Mon Apr 11 14:54:54 2016 +1000

    doc: document Go 1.6.1 and Go 1.5.4
    
    Change-Id: Icb9e947a43fb87fbfe0655b09e0d7e8f61825aeb
    Reviewed-on: https://go-review.googlesource.com/21825
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ba09d06e166a06b4405b2ffd92df6acf222d281f
Author: Jeremy Jackins <jeremyjackins@gmail.com>
Date:   Thu Apr 7 15:42:35 2016 +0900

    runtime: remove remaining references to TheChar
    
    After mdempsky's recent changes, these are the only references to
    "TheChar" left in the Go tree. Without the context, and without
    knowing the history, this is confusing.
    
    Also rename sys.TheGoos and sys.TheGoarch to sys.GOOS
    and sys.GOARCH.
    
    Also change the heap dump format to include sys.GOARCH
    rather than TheChar, which is no longer a concept.
    
    Updates #15169 (changes heapdump format)
    
    Change-Id: I3e99eeeae00ed55d7d01e6ed503d958c6e931dca
    Reviewed-on: https://go-review.googlesource.com/21647
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 1faa8869c6c72f055cdaa2b547964830909c96c6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Apr 6 12:31:55 2016 -0700

    net/http: set the Request context for incoming server requests
    
    Updates #13021
    Updates #15224
    
    Change-Id: Ia3cd608bb887fcfd8d81b035fa57bd5eb8edf09b
    Reviewed-on: https://go-review.googlesource.com/21810
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Emmanuel Odeke <emm.odeke@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bd7249766617fda12d112c3ad3ae2857ff97c71e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Apr 10 15:45:34 2016 +0000

    context: document that WithValue's key must be comparable
    
    Also, check it and explode earlier, rather than cryptic failures later.
    
    Change-Id: I319a425f60e2bc9d005a187fbdbd153faa96411c
    Reviewed-on: https://go-review.googlesource.com/21799
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Minux Ma <minux@golang.org>

commit e6a8daceb0b0df77f5d2ca34a73561375bb89c63
Author: David Symonds <dsymonds@golang.org>
Date:   Mon Apr 11 10:26:38 2016 +1000

    cmd/vet: refresh command for updating whitelist data.
    
    This excludes internal and testdata packages, as well as func types.
    
    No new whitelist entries were found.
    
    Change-Id: Ie7d42ce0a235394e4bcabf09e155726a35cd2d3d
    Reviewed-on: https://go-review.googlesource.com/21822
    Reviewed-by: Rob Pike <r@golang.org>

commit ad7448fe982d83de15deec9c55c56d0cd9261c6c
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sun Apr 10 17:32:35 2016 +0200

    runtime: speed up makeslice by avoiding divisions
    
    Only compute the number of maximum allowed elements per slice once.
    
    name         old time/op  new time/op  delta
    MakeSlice-2  55.5ns ± 1%  45.6ns ± 2%  -17.88%  (p=0.000 n=99+100)
    
    Change-Id: I951feffda5d11910a75e55d7e978d306d14da2c5
    Reviewed-on: https://go-review.googlesource.com/21801
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5b3e5766bcc5e1090d0512a5916886ffc24ab246
Author: Andrew Gerrand <adg@golang.org>
Date:   Fri Apr 8 14:43:35 2016 +1000

    cmd/go: remove special case that puts godoc in $GOROOT/bin
    
    Updates golang/go#15106
    
    Change-Id: I4214b841d63bb7e9c3c5ede2abe21a8a68f06c41
    Reviewed-on: https://go-review.googlesource.com/21701
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0fd270ab7abec08c050f29a3bbeb83d7740d0a47
Author: Andrew Gerrand <adg@golang.org>
Date:   Fri Apr 8 15:39:32 2016 +1000

    text/template: emit field error over nil pointer error where appropriate
    
    When evaluating "{{.MissingField}}" on a nil *T, Exec returns
    "can't evaluate field MissingField in type *T" instead of
    "nil pointer evaluating *T.MissingField".
    
    Fixes golang/go#15125
    
    Change-Id: I6e73f61b8a72c694179c1f8cdc808766c90b6f57
    Reviewed-on: https://go-review.googlesource.com/21705
    Reviewed-by: Rob Pike <r@golang.org>

commit 0004f34cefcdaad13a5131e3494fb2ff04877cd2
Author: Keith Randall <khr@golang.org>
Date:   Sun Apr 10 08:26:43 2016 -0700

    cmd/compile: regalloc enforces 2-address instructions
    
    Instead of being a hint, resultInArg0 is now enforced by regalloc.
    This allows us to delete all the code from amd64/ssa.go which
    deals with converting from a semantically three-address instruction
    into some copies plus a two-address instruction.
    
    Change-Id: Id4f39a80be4b678718bfd42a229f9094ab6ecd7c
    Reviewed-on: https://go-review.googlesource.com/21816
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 6b33b0e98e9be77d98b026ae2adf10dd71be5a1b
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Apr 10 09:08:00 2016 -0700

    cmd/compile: avoid a spill in append fast path
    
    Instead of spilling newlen, recalculate it.
    This removes a spill from the fast path,
    at the cost of a cheap recalculation
    on the (rare) growth path.
    This uses 8 bytes less of stack space.
    It generates two more bytes of code,
    but that is due to suboptimal register allocation;
    see far below.
    
    Runtime append microbenchmarks are all over the map,
    presumably due to incidental code movement.
    
    Sample code:
    
    func s(b []byte) []byte {
    	b = append(b, 1, 2, 3)
    	return b
    }
    
    Before:
    
    "".s t=1 size=160 args=0x30 locals=0x48
    	0x0000 00000 (append.go:8)	TEXT	"".s(SB), $72-48
    	0x0000 00000 (append.go:8)	MOVQ	(TLS), CX
    	0x0009 00009 (append.go:8)	CMPQ	SP, 16(CX)
    	0x000d 00013 (append.go:8)	JLS	149
    	0x0013 00019 (append.go:8)	SUBQ	$72, SP
    	0x0017 00023 (append.go:8)	FUNCDATA	$0, gclocals·6432f8c6a0d23fa7bee6c5d96f21a92a(SB)
    	0x0017 00023 (append.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    	0x0017 00023 (append.go:9)	MOVQ	"".b+88(FP), CX
    	0x001c 00028 (append.go:9)	LEAQ	3(CX), DX
    	0x0020 00032 (append.go:9)	MOVQ	DX, "".autotmp_0+64(SP)
    	0x0025 00037 (append.go:9)	MOVQ	"".b+96(FP), BX
    	0x002a 00042 (append.go:9)	CMPQ	DX, BX
    	0x002d 00045 (append.go:9)	JGT	$0, 86
    	0x002f 00047 (append.go:8)	MOVQ	"".b+80(FP), AX
    	0x0034 00052 (append.go:9)	MOVB	$1, (AX)(CX*1)
    	0x0038 00056 (append.go:9)	MOVB	$2, 1(AX)(CX*1)
    	0x003d 00061 (append.go:9)	MOVB	$3, 2(AX)(CX*1)
    	0x0042 00066 (append.go:10)	MOVQ	AX, "".~r1+104(FP)
    	0x0047 00071 (append.go:10)	MOVQ	DX, "".~r1+112(FP)
    	0x004c 00076 (append.go:10)	MOVQ	BX, "".~r1+120(FP)
    	0x0051 00081 (append.go:10)	ADDQ	$72, SP
    	0x0055 00085 (append.go:10)	RET
    	0x0056 00086 (append.go:9)	LEAQ	type.[]uint8(SB), AX
    	0x005d 00093 (append.go:9)	MOVQ	AX, (SP)
    	0x0061 00097 (append.go:9)	MOVQ	"".b+80(FP), BP
    	0x0066 00102 (append.go:9)	MOVQ	BP, 8(SP)
    	0x006b 00107 (append.go:9)	MOVQ	CX, 16(SP)
    	0x0070 00112 (append.go:9)	MOVQ	BX, 24(SP)
    	0x0075 00117 (append.go:9)	MOVQ	DX, 32(SP)
    	0x007a 00122 (append.go:9)	PCDATA	$0, $0
    	0x007a 00122 (append.go:9)	CALL	runtime.growslice(SB)
    	0x007f 00127 (append.go:9)	MOVQ	40(SP), AX
    	0x0084 00132 (append.go:9)	MOVQ	56(SP), BX
    	0x0089 00137 (append.go:8)	MOVQ	"".b+88(FP), CX
    	0x008e 00142 (append.go:9)	MOVQ	"".autotmp_0+64(SP), DX
    	0x0093 00147 (append.go:9)	JMP	52
    	0x0095 00149 (append.go:9)	NOP
    	0x0095 00149 (append.go:8)	CALL	runtime.morestack_noctxt(SB)
    	0x009a 00154 (append.go:8)	JMP	0
    
    After:
    
    "".s t=1 size=176 args=0x30 locals=0x40
    	0x0000 00000 (append.go:8)	TEXT	"".s(SB), $64-48
    	0x0000 00000 (append.go:8)	MOVQ	(TLS), CX
    	0x0009 00009 (append.go:8)	CMPQ	SP, 16(CX)
    	0x000d 00013 (append.go:8)	JLS	151
    	0x0013 00019 (append.go:8)	SUBQ	$64, SP
    	0x0017 00023 (append.go:8)	FUNCDATA	$0, gclocals·6432f8c6a0d23fa7bee6c5d96f21a92a(SB)
    	0x0017 00023 (append.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
    	0x0017 00023 (append.go:9)	MOVQ	"".b+80(FP), CX
    	0x001c 00028 (append.go:9)	LEAQ	3(CX), DX
    	0x0020 00032 (append.go:9)	MOVQ	"".b+88(FP), BX
    	0x0025 00037 (append.go:9)	CMPQ	DX, BX
    	0x0028 00040 (append.go:9)	JGT	$0, 81
    	0x002a 00042 (append.go:8)	MOVQ	"".b+72(FP), AX
    	0x002f 00047 (append.go:9)	MOVB	$1, (AX)(CX*1)
    	0x0033 00051 (append.go:9)	MOVB	$2, 1(AX)(CX*1)
    	0x0038 00056 (append.go:9)	MOVB	$3, 2(AX)(CX*1)
    	0x003d 00061 (append.go:10)	MOVQ	AX, "".~r1+96(FP)
    	0x0042 00066 (append.go:10)	MOVQ	DX, "".~r1+104(FP)
    	0x0047 00071 (append.go:10)	MOVQ	BX, "".~r1+112(FP)
    	0x004c 00076 (append.go:10)	ADDQ	$64, SP
    	0x0050 00080 (append.go:10)	RET
    	0x0051 00081 (append.go:9)	LEAQ	type.[]uint8(SB), AX
    	0x0058 00088 (append.go:9)	MOVQ	AX, (SP)
    	0x005c 00092 (append.go:9)	MOVQ	"".b+72(FP), BP
    	0x0061 00097 (append.go:9)	MOVQ	BP, 8(SP)
    	0x0066 00102 (append.go:9)	MOVQ	CX, 16(SP)
    	0x006b 00107 (append.go:9)	MOVQ	BX, 24(SP)
    	0x0070 00112 (append.go:9)	MOVQ	DX, 32(SP)
    	0x0075 00117 (append.go:9)	PCDATA	$0, $0
    	0x0075 00117 (append.go:9)	CALL	runtime.growslice(SB)
    	0x007a 00122 (append.go:9)	MOVQ	40(SP), AX
    	0x007f 00127 (append.go:9)	MOVQ	48(SP), CX
    	0x0084 00132 (append.go:9)	MOVQ	56(SP), BX
    	0x0089 00137 (append.go:9)	ADDQ	$3, CX
    	0x008d 00141 (append.go:9)	MOVQ	CX, DX
    	0x0090 00144 (append.go:8)	MOVQ	"".b+80(FP), CX
    	0x0095 00149 (append.go:9)	JMP	47
    	0x0097 00151 (append.go:9)	NOP
    	0x0097 00151 (append.go:8)	CALL	runtime.morestack_noctxt(SB)
    	0x009c 00156 (append.go:8)	JMP	0
    
    Observe that in the following sequence,
    we should use DX directly instead of using
    CX as a temporary register, which would make
    the new code a strict improvement on the old:
    
    	0x007f 00127 (append.go:9)	MOVQ	48(SP), CX
    	0x0084 00132 (append.go:9)	MOVQ	56(SP), BX
    	0x0089 00137 (append.go:9)	ADDQ	$3, CX
    	0x008d 00141 (append.go:9)	MOVQ	CX, DX
    	0x0090 00144 (append.go:8)	MOVQ	"".b+80(FP), CX
    
    Change-Id: I4ee50b18fa53865901d2d7f86c2cbb54c6fa6924
    Reviewed-on: https://go-review.googlesource.com/21812
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 974c201f74f730737964e5239da473fc548b408e
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Apr 10 10:43:04 2016 -0700

    runtime: avoid unnecessary map iteration write barrier
    
    Update #14921
    
    Change-Id: I5c5816d0193757bf7465b1e09c27ca06897df4bf
    Reviewed-on: https://go-review.googlesource.com/21814
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit a44c4256ae958b0aacecd5fd0b0e7f1156f8bcf4
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sun Apr 10 14:51:07 2016 +0000

    html: fix typo in UnescapeString string docs
    
    Fixes #15221
    
    Change-Id: I9e927a2f604213338b4572f1a32d0247c58bdc60
    Reviewed-on: https://go-review.googlesource.com/21798
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a56f5a032217e1e28c005886a98054caf7dc8201
Author: Klaus Post <klauspost@gmail.com>
Date:   Sun Apr 10 17:16:07 2016 +0200

    compress/flate: improve short writer error test
    
    This improves the short version of the writer test.
    
    First of all, it has a much quicker setup. Previously that
    could take up towards 0.5 second.
    
    Secondly, it will test all compression levels in short mode as well.
    
    Execution time is 1.7s/0.03s for normal/short mode.
    
    Change-Id: I275a21f712daff6f7125cc6a493415e86439cb19
    Reviewed-on: https://go-review.googlesource.com/21800
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f20b1809f213c662932106a68c76ea3545eab1ee
Author: Klaus Post <klauspost@gmail.com>
Date:   Sun Apr 10 13:43:24 2016 +0200

    compress/flate: eliminate most common bounds checks
    
    This uses the SSA compiler to eliminate various unneeded bounds
    checks in loops and various lookups.
    
    This fixes the low hanging fruit, without any major code changes.
    
    name                       old time/op    new time/op    delta
    EncodeDigitsHuffman1e4-8     49.9µs ± 1%    48.1µs ± 1%  -3.74%   (p=0.000 n=10+9)
    EncodeDigitsHuffman1e5-8      476µs ± 1%     458µs ± 1%  -3.58%  (p=0.000 n=10+10)
    EncodeDigitsHuffman1e6-8     4.80ms ± 2%    4.56ms ± 1%  -5.07%   (p=0.000 n=10+9)
    EncodeDigitsSpeed1e4-8        305µs ± 3%     290µs ± 2%  -5.03%   (p=0.000 n=10+9)
    EncodeDigitsSpeed1e5-8       3.67ms ± 2%    3.49ms ± 2%  -4.78%   (p=0.000 n=9+10)
    EncodeDigitsSpeed1e6-8       38.3ms ± 2%    35.8ms ± 1%  -6.58%   (p=0.000 n=9+10)
    EncodeDigitsDefault1e4-8      361µs ± 2%     346µs ± 3%  -4.12%   (p=0.000 n=10+9)
    EncodeDigitsDefault1e5-8     5.24ms ± 2%    4.96ms ± 3%  -5.38%  (p=0.000 n=10+10)
    EncodeDigitsDefault1e6-8     56.5ms ± 3%    52.2ms ± 2%  -7.68%  (p=0.000 n=10+10)
    EncodeDigitsCompress1e4-8     362µs ± 2%     343µs ± 1%  -5.20%   (p=0.000 n=10+9)
    EncodeDigitsCompress1e5-8    5.26ms ± 3%    4.98ms ± 2%  -5.48%  (p=0.000 n=10+10)
    EncodeDigitsCompress1e6-8    56.0ms ± 4%    52.1ms ± 1%  -7.01%  (p=0.000 n=10+10)
    EncodeTwainHuffman1e4-8      70.9µs ± 3%    64.7µs ± 1%  -8.68%   (p=0.000 n=10+9)
    EncodeTwainHuffman1e5-8       556µs ± 2%     524µs ± 2%  -5.84%  (p=0.000 n=10+10)
    EncodeTwainHuffman1e6-8      5.54ms ± 3%    5.22ms ± 2%  -5.70%  (p=0.000 n=10+10)
    EncodeTwainSpeed1e4-8         294µs ± 3%     284µs ± 1%  -3.71%  (p=0.000 n=10+10)
    EncodeTwainSpeed1e5-8        2.59ms ± 2%    2.48ms ± 1%  -4.14%   (p=0.000 n=10+9)
    EncodeTwainSpeed1e6-8        25.6ms ± 1%    24.3ms ± 1%  -5.28%   (p=0.000 n=9+10)
    EncodeTwainDefault1e4-8       419µs ± 2%     396µs ± 1%  -5.59%   (p=0.000 n=10+9)
    EncodeTwainDefault1e5-8      6.23ms ± 4%    5.75ms ± 1%  -7.83%   (p=0.000 n=10+9)
    EncodeTwainDefault1e6-8      66.2ms ± 2%    61.4ms ± 1%  -7.22%  (p=0.000 n=10+10)
    EncodeTwainCompress1e4-8      426µs ± 1%     405µs ± 1%  -4.97%   (p=0.000 n=9+10)
    EncodeTwainCompress1e5-8     6.80ms ± 1%    6.32ms ± 1%  -6.97%   (p=0.000 n=9+10)
    EncodeTwainCompress1e6-8     74.6ms ± 3%    68.7ms ± 1%  -7.90%   (p=0.000 n=10+9)
    
    name                       old speed      new speed      delta
    EncodeDigitsHuffman1e4-8    200MB/s ± 1%   208MB/s ± 1%  +3.88%   (p=0.000 n=10+9)
    EncodeDigitsHuffman1e5-8    210MB/s ± 1%   218MB/s ± 1%  +3.71%  (p=0.000 n=10+10)
    EncodeDigitsHuffman1e6-8    208MB/s ± 2%   219MB/s ± 1%  +5.32%   (p=0.000 n=10+9)
    EncodeDigitsSpeed1e4-8     32.8MB/s ± 3%  34.5MB/s ± 2%  +5.29%   (p=0.000 n=10+9)
    EncodeDigitsSpeed1e5-8     27.2MB/s ± 2%  28.6MB/s ± 2%  +5.29%  (p=0.000 n=10+10)
    EncodeDigitsSpeed1e6-8     26.1MB/s ± 2%  27.9MB/s ± 1%  +7.02%   (p=0.000 n=9+10)
    EncodeDigitsDefault1e4-8   27.7MB/s ± 2%  28.9MB/s ± 3%  +4.30%   (p=0.000 n=10+9)
    EncodeDigitsDefault1e5-8   19.1MB/s ± 2%  20.2MB/s ± 3%  +5.69%  (p=0.000 n=10+10)
    EncodeDigitsDefault1e6-8   17.7MB/s ± 3%  19.2MB/s ± 2%  +8.31%  (p=0.000 n=10+10)
    EncodeDigitsCompress1e4-8  27.6MB/s ± 2%  29.1MB/s ± 1%  +5.47%   (p=0.000 n=10+9)
    EncodeDigitsCompress1e5-8  19.0MB/s ± 3%  20.1MB/s ± 2%  +5.78%  (p=0.000 n=10+10)
    EncodeDigitsCompress1e6-8  17.9MB/s ± 4%  19.2MB/s ± 1%  +7.50%  (p=0.000 n=10+10)
    EncodeTwainHuffman1e4-8     141MB/s ± 3%   154MB/s ± 1%  +9.46%   (p=0.000 n=10+9)
    EncodeTwainHuffman1e5-8     180MB/s ± 2%   191MB/s ± 2%  +6.19%  (p=0.000 n=10+10)
    EncodeTwainHuffman1e6-8     181MB/s ± 3%   192MB/s ± 2%  +6.02%  (p=0.000 n=10+10)
    EncodeTwainSpeed1e4-8      34.0MB/s ± 3%  35.3MB/s ± 1%  +3.84%  (p=0.000 n=10+10)
    EncodeTwainSpeed1e5-8      38.7MB/s ± 2%  40.3MB/s ± 1%  +4.30%   (p=0.000 n=10+9)
    EncodeTwainSpeed1e6-8      39.1MB/s ± 1%  41.2MB/s ± 1%  +5.57%   (p=0.000 n=9+10)
    EncodeTwainDefault1e4-8    23.9MB/s ± 2%  25.3MB/s ± 1%  +5.91%   (p=0.000 n=10+9)
    EncodeTwainDefault1e5-8    16.0MB/s ± 4%  17.4MB/s ± 1%  +8.47%   (p=0.000 n=10+9)
    EncodeTwainDefault1e6-8    15.1MB/s ± 2%  16.3MB/s ± 1%  +7.76%  (p=0.000 n=10+10)
    EncodeTwainCompress1e4-8   23.5MB/s ± 1%  24.7MB/s ± 1%  +5.24%   (p=0.000 n=9+10)
    EncodeTwainCompress1e5-8   14.7MB/s ± 1%  15.8MB/s ± 1%  +7.50%   (p=0.000 n=9+10)
    EncodeTwainCompress1e6-8   13.4MB/s ± 3%  14.6MB/s ± 1%  +8.57%   (p=0.000 n=10+9)
    
    Change-Id: I5c7e84c2f9ea4d38a2115995705eebb93387e22f
    Reviewed-on: https://go-review.googlesource.com/21759
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 012557b3769f9286b9488fbfd4bddfeee66b6a55
Author: Martin Möhrmann <martisch@uos.de>
Date:   Sun Apr 10 08:48:55 2016 +0200

    all: replace magic 0x80 with named constant utf8.RuneSelf
    
    Change-Id: Id1c2e8e9d60588de866e8b6ca59cc83dd28f848f
    Reviewed-on: https://go-review.googlesource.com/21756
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 527ffebb2c9fe432a0ef0aa0c2449d83cd8a23cb
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Sun Apr 10 12:49:40 2016 +0200

    internal/trace: fix a typo in error message
    
    Change-Id: Id79eaa6d49dae80c334c7243b0a5bbcdcb9397d3
    Reviewed-on: https://go-review.googlesource.com/21758
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit de7ee57c7ead59899d5b412a839c995de0e813b5
Author: Marvin Stenger <marvin.stenger94@gmail.com>
Date:   Fri Apr 8 18:19:10 2016 +0200

    cmd: remove bio.Bread
    
    Replace calls to bio.Bread with calls to io.ReadFull.
    
    Change-Id: I2ee8739d01e04a4da9c20b6ce7d1d5b89914b8ad
    Reviewed-on: https://go-review.googlesource.com/21750
    Reviewed-by: Dave Cheney <dave@cheney.net>

commit e4f1d9cf2e948eb0f0bb91d7c253ab61dfff3a59
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sun Mar 27 17:29:53 2016 -0700

    runtime: make execution error panic values implement the Error interface
    
    Make execution panics implement Error as
    mandated by https://golang.org/ref/spec#Run_time_panics,
    instead of panics with strings.
    
    Fixes #14965
    
    Change-Id: I7827f898b9b9c08af541db922cc24fa0800ff18a
    Reviewed-on: https://go-review.googlesource.com/21214
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 824d8c10fe5e1026c15cbece41ee372b1fd333f3
Author: Wisdom Omuya <deafgoat@gmail.com>
Date:   Sat Apr 9 19:23:01 2016 -0400

    cmd/go: fix typo in findInternal documentation
    
    Fixes #15217
    
    Change-Id: Ib8f7af714197fd209e743f61f28a5b07c04a7f5c
    Reviewed-on: https://go-review.googlesource.com/21793
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3598d4b8387a9e1c9afd522e0d18201f855f613b
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Sat Apr 9 00:34:18 2016 -0700

    net/http/httputil: DumpRequest dumps Content-Length if set in header
    
    Fixes #7215
    
    Change-Id: I108171ef18cac66d0dc11ce3553c26fc49529807
    Reviewed-on: https://go-review.googlesource.com/21790
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit c31fdd4ee9fccd24a274cebd82dcc7123ad43d0e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Apr 9 22:54:26 2016 +0000

    cmd/go: fix typo in comment
    
    Thanks to deafgoat.
    
    Fixes #15215
    
    Change-Id: I9fababc7ecd201ce86020a438e4faee95e7623a8
    Reviewed-on: https://go-review.googlesource.com/21792
    Reviewed-by: Emmanuel Odeke <emm.odeke@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 0435e88a119fd057aa7209591ba3dff122c9f24c
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri Mar 18 10:56:23 2016 +0100

    runtime: revert "do not call timeBeginPeriod on windows"
    
    This reverts commit ab4c9298b8185a056ff1152f2c7bd9b38d3d06f3.
    
    Sysmon critically depends on system timer resolution for retaking
    of Ps blocked in system calls. See #14790 for an example
    of a program where execution time goes from 2ms to 30ms if
    timeBeginPeriod(1) is not used.
    
    We can remove timeBeginPeriod(1) when we support UMS (#7876).
    
    Update #14790
    
    Change-Id: I362b56154359b2c52d47f9f2468fe012b481cf6d
    Reviewed-on: https://go-review.googlesource.com/20834
    Reviewed-by: Austin Clements <austin@google.com>
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Alex Brainman <alex.brainman@gmail.com>

commit 9d4efdfd12f47f1ed8ce482ebeeb4d4e30a2dbc6
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Apr 8 20:44:56 2016 +1000

    cmd/internal/bio: move Bgetc to link/internal/ld
    
    Also, remove bio.Brdline.
    
    Change-Id: I3e0caed27a373fd71737cf6892de5e8fc208b776
    Reviewed-on: https://go-review.googlesource.com/21783
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Dave Cheney <dave@cheney.net>

commit bce9747ed00c53e7ddeea102e87aede1b3ec9bd3
Author: Dave Cheney <dave@cheney.net>
Date:   Sat Apr 9 15:04:45 2016 +1000

    cmd: remove unused code
    
    Generated with honnef.co/go/unused
    
    There is a large amount of unused code in cmd/internal/obj/s390x but
    that can wait til the s390x port is merged.
    
    There is some unused code in
    cmd/internal/unvendor/golang.org/x/arch/arm/armasm but that should be
    addressed upstream and a new revision imported.
    
    Change-Id: I252c0f9ea8c5bb1a0b530a374ef13a0a20ea56aa
    Reviewed-on: https://go-review.googlesource.com/21782
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Dave Cheney <dave@cheney.net>

commit 93368be61ebaf8069d0d70034097de580441c412
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Apr 8 20:37:54 2016 +1000

    cmd/internal/bio: embed bufio.{Reader,Writer} in bio.{Reader,Writer}
    
    Change-Id: Ie95b0b0d4f724f4769cf2d4f8063cb5019fa9bc9
    Reviewed-on: https://go-review.googlesource.com/21781
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ca397bb68e4b548843d2886e374f96ec3bb0f9c0
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Apr 8 19:30:41 2016 +1000

    cmd: remove bio.BufReader and bio.BufWriter
    
    bio.BufReader was never used.
    
    bio.BufWriter was used to wrap an existing io.Writer, but the
    bio.Writer returned would not be seekable, so replace all occurences
    with bufio.Reader instead.
    
    Change-Id: I9c6779e35c63178aa4e104c17bb5bb8b52de0359
    Reviewed-on: https://go-review.googlesource.com/21722
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6fee4aa5c76612d133d2b01441e608cfa696bae9
Author: Dave Cheney <dave@cheney.net>
Date:   Sat Apr 9 12:54:45 2016 +1000

    cmd/link/internal: make ld.Bso a *bio.Writer
    
    This is a pre requesite of CL 21722 and removes a lot of unidiomatic
    boilerplate in the linker.
    
    Change-Id: If7491b88212b2be7b0c8c588e9c196839131f8ad
    Reviewed-on: https://go-review.googlesource.com/21780
    Run-TryBot: Dave Cheney <dave@cheney.net>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 092ef8a2ca60e1a7573442757b02ec1efc456c2c
Author: Michael Hudson-Doyle <michael.hudson@canonical.com>
Date:   Fri Apr 8 14:27:35 2016 +1200

    cmd/cgo: fix cgo with gccgo
    
    Change-Id: I1780899255e22c16d7f8e9947609a1c284d7c42e
    Reviewed-on: https://go-review.googlesource.com/21690
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Michael Hudson-Doyle <michael.hudson@canonical.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c3b3e7b4ef9dff1fc0cc504f81465ded5663b4e4
Author: David Chase <drchase@google.com>
Date:   Fri Apr 8 13:33:43 2016 -0400

    cmd/compile: insert instrumentation more carefully in racewalk
    
    Be more careful about inserting instrumentation in racewalk.
    If the node being instrumented is an OAS, and it has a non-
    empty Ninit, then append instrumentation to the Ninit list
    rather than letting it be inserted before the OAS (and the
    compilation of its init list).  This deals with the case that
    the Ninit list defines a variable used in the RHS of the OAS.
    
    Fixes #15091.
    
    Change-Id: Iac91696d9104d07f0bf1bd3499bbf56b2e1ef073
    Reviewed-on: https://go-review.googlesource.com/21771
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Run-TryBot: David Chase <drchase@google.com>

commit 0fb7b4cccd02df10f239ed77d6d85566b6388b83
Author: Dmitry Vyukov <dvyukov@google.com>
Date:   Fri Apr 8 15:14:37 2016 +0200

    runtime: emit file:line info into traces
    
    This makes traces self-contained and simplifies trace workflow
    in modern cloud environments where it is simpler to reach
    a service via HTTP than to obtain the binary.
    
    Change-Id: I6ff3ca694dc698270f1e29da37d5efaf4e843a0d
    Reviewed-on: https://go-review.googlesource.com/21732
    Run-TryBot: Dmitry Vyukov <dvyukov@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Hyang-Ah Hana Kim <hyangah@gmail.com>

commit 9ada88aec271a2f08c998e9669331145803e7d5a
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Apr 8 13:02:49 2016 -0400

    cmd/cgo: increase s390x int type size to 8 bytes
    
    The size of the int type in Go on s390x is 8 bytes, not 4.
    
    Change-Id: I1a71ce8c9925f3499abb61c1aa4f6fa2d2ec0d7e
    Reviewed-on: https://go-review.googlesource.com/21760
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 59af53d681845a8b0be2a728ca1b59aee5ad9ea6
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Thu Apr 7 23:22:30 2016 -0700

    bytes: add ContainsRune
    
    Make package bytes consistent with strings
    by adding missing function ContainsRune.
    
    Fixes #15189
    
    Change-Id: Ie09080b389e55bbe070c57aa3bd134053a805423
    Reviewed-on: https://go-review.googlesource.com/21710
    Run-TryBot: Rob Pike <r@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 6c5352f181846b73d532c039df3017befe657d6a
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Apr 8 11:40:51 2016 -0400

    syscall: add assembly for Linux on s390x
    
    Change-Id: I42ade65a91f3effc03dd663ee449410baa9f8ca8
    Reviewed-on: https://go-review.googlesource.com/21740
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Minux Ma <minux@golang.org>
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 59d186832b94349d683431e01e084d6ce460f476
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Apr 8 17:50:03 2016 +0000

    context: disable more flaky tests on openbsd
    
    Updates #15158
    
    Change-Id: Icb3788152a7a5a9b0d56ea38da46d770ffdce413
    Reviewed-on: https://go-review.googlesource.com/21763
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit dae98d5c3b02b39f53168a9403d24e1ddd4a16d4
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Mar 20 19:35:59 2016 -0400

    test: skip nilptr3 test on s390x
    
    Fails for the same reason as ppc64 and mips64 (incomplete
    optimization).
    
    Change-Id: Ieb4d997fc27d4f2b756e63dd7f588abe10c0213a
    Reviewed-on: https://go-review.googlesource.com/20963
    Reviewed-by: Bill O'Farrell <billotosyr@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 78715cebcfcca3aaaaba4dd41ef6b82a46d7b93d
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Apr 8 11:55:05 2016 -0400

    cmd/link: add s390x to link tool main function
    
    Change-Id: I83bc2b4a00216b069f133113e4ae9ad76c98a708
    Reviewed-on: https://go-review.googlesource.com/21741
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 01360a64c5895a9ec8b8c34140415fe34c3de201
Author: Rob Pike <r@golang.org>
Date:   Wed Apr 6 20:19:47 2016 -0700

    io: change the name of ReadAtSizer to SizedReaderAt
    
    This is a proposal. The old name is pretty poor. The new one describes
    it better and may be easier to remember. It does not start with Read,
    though I think that inconsistency is worthwhile.
    
    Reworded the comment a bit for clarity.
    
    Change-Id: Icb4f9c663cc68958e0363d7ff78a0b29cc521f98
    Reviewed-on: https://go-review.googlesource.com/21629
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 8f2edf11998a30b497586ac0e9f75036a318280a
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Apr 8 19:14:03 2016 +1000

    cmd: replace bio.Buf with bio.Reader and bio.Writer
    
    Replace the bidirectional bio.Buf type with a pair of unidirectional
    buffered seekable Reader and Writers.
    
    Change-Id: I86664a06f93c94595dc67c2cbd21356feb6680ef
    Reviewed-on: https://go-review.googlesource.com/21720
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit d22357ce9dc650a69e78b37a6b25be1ee0b8b26c
Author: Michael Munday <munday@ca.ibm.com>
Date:   Thu Apr 7 15:31:49 2016 -0400

    cmd/compile: cleanup -dynlink/-shared support check
    
    Moves the list of architectures that support shared libraries into
    a function. Also adds s390x to that list.
    
    Change-Id: I99c8a9f6cd4816ce3d53abaabaf8d002e25e6b28
    Reviewed-on: https://go-review.googlesource.com/21661
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>
    Run-TryBot: Michael Munday <munday@ca.ibm.com>

commit 49e07f2b7e25a1f7a050f73fbb7807185e09e46b
Author: Dave Cheney <dave@cheney.net>
Date:   Fri Apr 8 20:09:10 2016 +1000

    cmd/compile/internal/gc: unexport Export
    
    Export does not need to be exported.
    
    Change-Id: I252f0c024732f1d056817cab13e8e3c589b54d13
    Reviewed-on: https://go-review.googlesource.com/21721
    Run-TryBot: Dave Cheney <dave@cheney.net>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 4dae828f77a37ed87401f7877998b241f0d2c33e
Author: Andrew Gerrand <adg@golang.org>
Date:   Fri Apr 8 15:09:45 2016 +1000

    cmd/go: fix failing tests since vet was moved from x/tools
    
    Change-Id: I3276a118ced78f3efd8f1bc5fb8b8fa2fde52496
    Reviewed-on: https://go-review.googlesource.com/21704
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 853f1a1a63b759686421196d419ddaa626a44bf5
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Thu Apr 7 19:59:59 2016 -0700

    net/http: fixed trivial go vet warnings
    
    Updates #15177
    
    Change-Id: I748f025461f313b5b426821ead695f90d3011a6b
    Reviewed-on: https://go-review.googlesource.com/21677
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 68ac1f774624faf99e7f6ec6592acb50f23b7874
Author: Keith Randall <khr@golang.org>
Date:   Thu Apr 7 10:21:35 2016 -0700

    cmd/compile: Fix constant-folding of unsigned shifts
    
    Make sure the results of unsigned constant-folded
    shifts are sign-extended into the AuxInt field.
    
    Fixes #15175
    
    Change-Id: I3490d1bc3d9b2e1578ed30964645508577894f58
    Reviewed-on: https://go-review.googlesource.com/21586
    Reviewed-by: Alexandru Moșoi <alexandru@mosoi.ro>
    Run-TryBot: Keith Randall <khr@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit dfded57819dd9111afffc25360cd3e147859d354
Author: Cheng-Lung Sung <clsung@gmail.com>
Date:   Wed Apr 6 23:05:20 2016 +0800

    cmd/go: revise importPath when ImportPath is 'command-line-arguments'
    
    Fixes #14613
    
    Change-Id: I40d9696db3879549e78373ef17f6c92bd4b3470b
    Reviewed-on: https://go-review.googlesource.com/21596
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e88f89028a55acf9c8b76b7f6ca284c3f9eb4cbd
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Thu Mar 31 16:05:23 2016 -0700

    bytes, string: add Reset method to Reader
    
    Currently, there is no easy allocation-free way to turn a
    []byte or string into an io.Reader. Thus, we add a Reset method
    to bytes.Reader and strings.Reader to allow the reuse of these
    Readers with another []byte or string.
    
    This is consistent with the fact that many standard library io.Readers
    already support a Reset method of some type:
    	bufio.Reader
    	flate.Reader
    	gzip.Reader
    	zlib.Reader
    	debug/dwarf.LineReader
    	bytes.Buffer
    	crypto/rc4.Cipher
    
    Fixes #15033
    
    Change-Id: I456fd1af77af6ef0b4ac6228b058ac1458ff3d19
    Reviewed-on: https://go-review.googlesource.com/21386
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e6f36f0cd5b45b9ce7809a34c45aeb66a5ca64a4
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Mar 18 19:09:39 2016 -0400

    runtime: add s390x support (new files and lfstack_64bit.go modifications)
    
    Change-Id: I51c0a332e3cbdab348564e5dcd27583e75e4b881
    Reviewed-on: https://go-review.googlesource.com/20946
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9658b7ef83ae9c34f4a52680e7102d958577d5bb
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Apr 7 14:27:15 2016 -0400

    cmd/link: hide go.dwarf symbols
    
    Fixes #15179
    
    Change-Id: I0f70b7ae1682eafaece7f22d8e76f0aa806f3ec9
    Reviewed-on: https://go-review.googlesource.com/21589
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e6181eb9e1dc4ab9e297a102ed192997582ac46c
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Apr 7 14:00:00 2016 -0400

    cmd/link: disable DWARF when not generating symtab
    
    Fixes #15166
    
    Change-Id: I30284e3c0fb2c80b26a2572e2fb249b8018e85f9
    Reviewed-on: https://go-review.googlesource.com/21587
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 39f1ecd157941420693239fa269643ce8a34c267
Author: Burcu Dogan <jbd@google.com>
Date:   Thu Apr 7 10:50:47 2016 -0700

    C: fix jbd's identity
    
    Change-Id: Ib4353710a742b1067723c7c6186e8639559668a3
    Reviewed-on: https://go-review.googlesource.com/21655
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit e2c09749af8c50fc2c0b515f2adc990cb0cb3cf6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Apr 7 09:00:32 2016 -0700

    context: mark more tests as flaky on OpenBSD
    
    Updates #15158
    
    Change-Id: I53e9e68d36efbf52736822e6caa047cfff501283
    Reviewed-on: https://go-review.googlesource.com/21653
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 187afdebef7953295189d4531e7dccdc0cb64500
Author: Aliaksandr Valialkin <valyala@gmail.com>
Date:   Mon Apr 4 19:28:15 2016 +0300

    math/big: re-use memory in Int.GCD
    
    This improves TLS handshake performance.
    
    benchmark                                 old ns/op     new ns/op     delta
    BenchmarkGCD10x10/WithoutXY-4             965           968           +0.31%
    BenchmarkGCD10x10/WithXY-4                1813          1391          -23.28%
    BenchmarkGCD10x100/WithoutXY-4            1093          1075          -1.65%
    BenchmarkGCD10x100/WithXY-4               2348          1676          -28.62%
    BenchmarkGCD10x1000/WithoutXY-4           1569          1565          -0.25%
    BenchmarkGCD10x1000/WithXY-4              4262          3242          -23.93%
    BenchmarkGCD10x10000/WithoutXY-4          6069          6066          -0.05%
    BenchmarkGCD10x10000/WithXY-4             12123         11331         -6.53%
    BenchmarkGCD10x100000/WithoutXY-4         52664         52610         -0.10%
    BenchmarkGCD10x100000/WithXY-4            97494         95649         -1.89%
    BenchmarkGCD100x100/WithoutXY-4           5244          5228          -0.31%
    BenchmarkGCD100x100/WithXY-4              22572         18630         -17.46%
    BenchmarkGCD100x1000/WithoutXY-4          6143          6233          +1.47%
    BenchmarkGCD100x1000/WithXY-4             24652         19357         -21.48%
    BenchmarkGCD100x10000/WithoutXY-4         15725         15804         +0.50%
    BenchmarkGCD100x10000/WithXY-4            60552         55973         -7.56%
    BenchmarkGCD100x100000/WithoutXY-4        107008        107853        +0.79%
    BenchmarkGCD100x100000/WithXY-4           349597        340994        -2.46%
    BenchmarkGCD1000x1000/WithoutXY-4         63785         64434         +1.02%
    BenchmarkGCD1000x1000/WithXY-4            373186        334035        -10.49%
    BenchmarkGCD1000x10000/WithoutXY-4        78038         78241         +0.26%
    BenchmarkGCD1000x10000/WithXY-4           543692        507034        -6.74%
    BenchmarkGCD1000x100000/WithoutXY-4       205607        207727        +1.03%
    BenchmarkGCD1000x100000/WithXY-4          2488113       2415323       -2.93%
    BenchmarkGCD10000x10000/WithoutXY-4       1731340       1714992       -0.94%
    BenchmarkGCD10000x10000/WithXY-4          10601046      7111329       -32.92%
    BenchmarkGCD10000x100000/WithoutXY-4      2239155       2212173       -1.21%
    BenchmarkGCD10000x100000/WithXY-4         30097040      26538887      -11.82%
    BenchmarkGCD100000x100000/WithoutXY-4     119845326     119863916     +0.02%
    BenchmarkGCD100000x100000/WithXY-4        768006543     426795966     -44.43%
    
    benchmark                                 old allocs     new allocs     delta
    BenchmarkGCD10x10/WithoutXY-4             5              5              +0.00%
    BenchmarkGCD10x10/WithXY-4                17             9              -47.06%
    BenchmarkGCD10x100/WithoutXY-4            6              6              +0.00%
    BenchmarkGCD10x100/WithXY-4               21             9              -57.14%
    BenchmarkGCD10x1000/WithoutXY-4           6              6              +0.00%
    BenchmarkGCD10x1000/WithXY-4              30             12             -60.00%
    BenchmarkGCD10x10000/WithoutXY-4          6              6              +0.00%
    BenchmarkGCD10x10000/WithXY-4             26             12             -53.85%
    BenchmarkGCD10x100000/WithoutXY-4         6              6              +0.00%
    BenchmarkGCD10x100000/WithXY-4            28             12             -57.14%
    BenchmarkGCD100x100/WithoutXY-4           5              5              +0.00%
    BenchmarkGCD100x100/WithXY-4              183            61             -66.67%
    BenchmarkGCD100x1000/WithoutXY-4          8              8              +0.00%
    BenchmarkGCD100x1000/WithXY-4             170            47             -72.35%
    BenchmarkGCD100x10000/WithoutXY-4         8              8              +0.00%
    BenchmarkGCD100x10000/WithXY-4            200            67             -66.50%
    BenchmarkGCD100x100000/WithoutXY-4        8              8              +0.00%
    BenchmarkGCD100x100000/WithXY-4           188            65             -65.43%
    BenchmarkGCD1000x1000/WithoutXY-4         5              5              +0.00%
    BenchmarkGCD1000x1000/WithXY-4            2435           1193           -51.01%
    BenchmarkGCD1000x10000/WithoutXY-4        8              8              +0.00%
    BenchmarkGCD1000x10000/WithXY-4           2211           1076           -51.33%
    BenchmarkGCD1000x100000/WithoutXY-4       8              8              +0.00%
    BenchmarkGCD1000x100000/WithXY-4          2271           1108           -51.21%
    BenchmarkGCD10000x10000/WithoutXY-4       5              5              +0.00%
    BenchmarkGCD10000x10000/WithXY-4          23183          11605          -49.94%
    BenchmarkGCD10000x100000/WithoutXY-4      8              8              +0.00%
    BenchmarkGCD10000x100000/WithXY-4         23421          11717          -49.97%
    BenchmarkGCD100000x100000/WithoutXY-4     5              5              +0.00%
    BenchmarkGCD100000x100000/WithXY-4        232976         116815         -49.86%
    
    benchmark                                 old bytes      new bytes     delta
    BenchmarkGCD10x10/WithoutXY-4             208            208           +0.00%
    BenchmarkGCD10x10/WithXY-4                736            432           -41.30%
    BenchmarkGCD10x100/WithoutXY-4            256            256           +0.00%
    BenchmarkGCD10x100/WithXY-4               896            432           -51.79%
    BenchmarkGCD10x1000/WithoutXY-4           368            368           +0.00%
    BenchmarkGCD10x1000/WithXY-4              1856           1152          -37.93%
    BenchmarkGCD10x10000/WithoutXY-4          1616           1616          +0.00%
    BenchmarkGCD10x10000/WithXY-4             7920           7376          -6.87%
    BenchmarkGCD10x100000/WithoutXY-4         13776          13776         +0.00%
    BenchmarkGCD10x100000/WithXY-4            68800          68176         -0.91%
    BenchmarkGCD100x100/WithoutXY-4           208            208           +0.00%
    BenchmarkGCD100x100/WithXY-4              6960           2112          -69.66%
    BenchmarkGCD100x1000/WithoutXY-4          544            560           +2.94%
    BenchmarkGCD100x1000/WithXY-4             7280           2400          -67.03%
    BenchmarkGCD100x10000/WithoutXY-4         2896           2912          +0.55%
    BenchmarkGCD100x10000/WithXY-4            15280          10002         -34.54%
    BenchmarkGCD100x100000/WithoutXY-4        27344          27365         +0.08%
    BenchmarkGCD100x100000/WithXY-4           88288          83427         -5.51%
    BenchmarkGCD1000x1000/WithoutXY-4         544            544           +0.00%
    BenchmarkGCD1000x1000/WithXY-4            178288         40043         -77.54%
    BenchmarkGCD1000x10000/WithoutXY-4        3344           3136          -6.22%
    BenchmarkGCD1000x10000/WithXY-4           188720         54432         -71.16%
    BenchmarkGCD1000x100000/WithoutXY-4       27792          27592         -0.72%
    BenchmarkGCD1000x100000/WithXY-4          373872         239447        -35.95%
    BenchmarkGCD10000x10000/WithoutXY-4       4288           4288          +0.00%
    BenchmarkGCD10000x10000/WithXY-4          11935584       481875        -95.96%
    BenchmarkGCD10000x100000/WithoutXY-4      31296          28834         -7.87%
    BenchmarkGCD10000x100000/WithXY-4         13237088       1662620       -87.44%
    BenchmarkGCD100000x100000/WithoutXY-4     40768          40768         +0.00%
    BenchmarkGCD100000x100000/WithXY-4        1165518864     14256010      -98.78%
    
    Change-Id: I652b3244bd074a03f3bc9a87c282330f9e5f1507
    Reviewed-on: https://go-review.googlesource.com/21506
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1e7c61d8c3fc0916b73cb7411afa45e99b50aac4
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Mar 20 21:49:52 2016 -0400

    math/big: add s390x function implementations
    
    Change-Id: I2aadc885d6330460e494c687757f07c5e006f3b0
    Reviewed-on: https://go-review.googlesource.com/20937
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 7da42d75975044df37aa3aa2499623e2084a12df
Author: Michael Munday <munday@ca.ibm.com>
Date:   Sun Mar 20 21:34:48 2016 -0400

    sync/atomic: add s390x implementations of atomic functions
    
    Load and store instructions are atomic on s390x.
    
    Change-Id: I33c641a75954f4fbd301b11a467cb57872038880
    Reviewed-on: https://go-review.googlesource.com/20947
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
