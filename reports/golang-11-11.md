# November 11, 2016 Report

Number of commits: 177

## Compilation time

* github.com/boltdb/bolt/cmd/bolt: from 510.85271ms to 575.410616ms, +12.64%
* github.com/junegunn/fzf/src/fzf: from 947.456904ms to 974.278145ms, +2.83%
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 1.185812348s to 1.172555943s, -1.12%
* golang.org/x/tools/cmd/guru: from 2.520040588s to 2.391774108s, -5.09%

## Binary size:

* github.com/boltdb/bolt/cmd/bolt: from 2492280 to 3069005, +23.14%
* github.com/junegunn/fzf/src/fzf: from 3041992 to 3042256, ~
* github.com/monochromegane/the_platinum_searcher/cmd/pt: from 4239163 to 4235000, ~
* golang.org/x/tools/cmd/guru: from 7588998 to 7941217, +4.64%

## Bechmarks:

```
benchmark                            old ns/op     new ns/op     delta
BenchmarkMsgpMarshal-4               191           185           -3.14%
BenchmarkMsgpUnmarshal-4             370           369           -0.27%
BenchmarkEasyJsonMarshal-4           1410          1401          -0.64%
BenchmarkEasyJsonUnmarshal-4         1786          1553          -13.05%
BenchmarkFlatBuffersMarshal-4        439           363           -17.31%
BenchmarkFlatBuffersUnmarshal-4      294           295           +0.34%
BenchmarkGogoprotobufMarshal-4       157           153           -2.55%
BenchmarkGogoprotobufUnmarshal-4     249           249           +0.00%

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

* [cmd/compile/internal/ssa: generate bswap on AMD64](https://github.com/golang/go/commit/10f757486e94f60a5e0af180dcd61c9eef7534c6)
* [testing: mark tests and benchmarks failed if a race occurs during execution](https://github.com/golang/go/commit/43f954e09858449cc5f3650720e81b7e879ab349)
* [testing: add T.Context method](https://github.com/golang/go/commit/26827bc2fe4c80dc68b3793631d24975425c9467)
* port to mips,le by [Imagination](https://imgtec.com/)
* [Revert "spec: add new language for alias declarations"](https://github.com/golang/go/commit/87f4e36ce7d7dffbf1f2a869f3014321f6cfff3c)
* [os: add Executable() (string, error)](https://github.com/golang/go/commit/2fc67e71af142bfa1e7662a4fde361f43509d2d7)
* [go/build: implement default GOPATH](https://github.com/golang/go/commit/dc4a815d100b82643656ec88fd9fa8e7c705ebba)

## GIT Log

```
commit add721ef91ed533cf578ff7a604124e377329ae4
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Mon Oct 24 21:20:04 2016 -0700

    encoding/json: encode nil Marshaler as "null"
    
    Fixes #16042.
    
    Change-Id: I0a28aa004246b7b0ffaaab457e077ad9035363c2
    Reviewed-on: https://go-review.googlesource.com/31932
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit c439a5d8b77a3b87d94bc49714faeeb2a04112f4
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Fri Sep 16 16:25:48 2016 +0900

    cmd/pprof: don't print binary outputs in interactive mode
    
    Some commands generate binary outputs which are not human readable.
    In interactive mode, there are no use-cases for such outputs.
    Instead, the new code writes it to the temporary file on the $CWD and
    shows the file name. So the user can use any program to display the
    file outside interactive shell.
    
    Fixes #17465
    
    Change-Id: I5c479db26017607f7a28eafbff2385533e5c584e
    Reviewed-on: https://go-review.googlesource.com/31123
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 428df5e39c0a696b71724237879a22a718a854a7
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Nov 10 22:06:08 2016 -0800

    cmd/go: don't set default GOPATH to GOROOT
    
    It will just cause confusion later as the go tool will say
    "warning: GOPATH set to GOROOT (%s) has no effect".
    Better to just leave GOPATH unset and get that warning instead.
    
    Change-Id: I78ff9e87fdf4bb0460f4f6d6ee76e1becaa3e7b0
    Reviewed-on: https://go-review.googlesource.com/33105
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 66477ec8307e18f751996d92ac8741596a23615a
Author: David Crawshaw <crawshaw@golang.org>
Date:   Thu Nov 10 19:02:07 2016 -0500

    reflect: rename, document TestUnaddressableField
    
    Change-Id: I94e0f3e4bccd44a67934ddb4d5fc7da57bb8ac9f
    Reviewed-on: https://go-review.googlesource.com/33112
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 53aec79ce05cd5eff1c8f5576b553d3c429227c3
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Nov 10 22:08:51 2016 -0800

    cmd/link: for -buildmode=exe pass -no-pie to external linker
    
    On some systems the external linker defaults to PIE. On some systems
    DT_TEXTREL does not work correctly. When both are true we have a bad
    situation: any Go program built with the default buildmode (exe) that
    uses external linking will fail to run. Fix this by passing -no-pie to
    the external linker, if the option is supported.
    
    Fixes #17847.
    
    Change-Id: I9b5ff97825d8b7f494f96d29c4c04f72b53dbf4e
    Reviewed-on: https://go-review.googlesource.com/33106
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 0631f292d370cd838e1fb57e193c6c09e74fe9e4
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 11 00:02:45 2016 +0000

    net/http: document relation and interaction with golang.org/x/net/http2
    
    Fixes #16412
    
    Change-Id: Idc65d2a62414a9b1573e6bd9f8601b52985b5dea
    Reviewed-on: https://go-review.googlesource.com/33110
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit fb8c896aff9549e868df58f9d40fd06b67ae7d07
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Nov 10 14:34:32 2016 -0800

    cmd/cgo: don't ignore qualifiers, don't cast to void*
    
    The cgo tool used to simply ignore C type qualifiers. To avoid problems
    when a C function expected a qualifier that was not present, cgo emitted
    a cast to void* around all pointer arguments. Unfortunately, that broke
    code that contains both a function declaration and a macro, when the
    macro required the argument to have the right type. To fix this problem,
    don't ignore qualifiers. They are easy enough to handle for the limited
    set of cases that matter for cgo, in which we don't care about array or
    function types.
    
    Fixes #17537.
    
    Change-Id: Ie2988d21db6ee016a3e99b07f53cfb0f1243a020
    Reviewed-on: https://go-review.googlesource.com/33097
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit dc4a815d100b82643656ec88fd9fa8e7c705ebba
Author: Francesc Campoy <campoy@golang.org>
Date:   Mon Oct 31 21:36:38 2016 -0700

    go/build: implement default GOPATH
    
    Whenever GOPATH is not defined in the environment, use $HOME/go
    as its default value. For Windows systems use %USERPROFILE%/go
    and $home/go for plan9.
    
    The choice of these environment variables is based on what Docker
    currently does. The os/user package is not used to avoid having
    a cgo dependency.
    
    Updates #17262. Documentation changes forthcoming.
    
    Change-Id: I6368fbfbc5afda99d6e64c35c1980076fcf45344
    Reviewed-on: https://go-review.googlesource.com/32019
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit ebc0b625a07ccce6ade7a0082f4ab49c2817e965
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Nov 10 16:38:22 2016 -0800

    doc/go1.8.txt: mention that struct conversions ignore tags
    
    Also:
    - update performance improvements for CL 31275.
    
    Change-Id: I2f2ec0a42b248643e76df8654e11bf0b01a5d030
    Reviewed-on: https://go-review.googlesource.com/33114
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit eb4e17b73b4ad486f9e9b0ea0fe2a6050ceb54fc
Author: David Crawshaw <crawshaw@golang.org>
Date:   Tue Nov 8 13:59:25 2016 -0500

    cmd/link: use plugin path in visibility analysis
    
    CL 32355 switched from using the output file as a
    plugin prefix to the full package path. The linker dead code analysis
    was not updated.
    
    Updates #17821
    
    Change-Id: I13fc45e0264b425d28524ec54c829e2c3e895b0b
    Reviewed-on: https://go-review.googlesource.com/32916
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 7bdb77af5fea5b94cf3d5d7840ca9162e76b3e9b
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Nov 10 15:19:11 2016 -0800

    cmd/cgo: don't depend on runtime/cgo if !CgoEnabled
    
    Fixes the build when CGO_ENABLED=0.
    
    Change-Id: I7f3c67d61e156e69536558fda0a0a4b429b82bbd
    Reviewed-on: https://go-review.googlesource.com/33104
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 35ea53dcc8d8350898250e87a0b5ffa03e14173e
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Nov 10 14:40:32 2016 -0800

    cmd/gofmt: don't overwrite read-only files
    
    This reverts the changes from https://golang.org/cl/33018: Instead
    of writing the result of gofmt to a tmp file and then rename that
    to the original (which doesn't preserve the original file's perm
    bits, uid, gid, and possibly other properties because it is hard
    to do in a platform-independent way - see #17869), use the original
    code that simply overwrites the processed file if gofmt was able to
    create a backup first. Upon success, the backup is removed, otherwise
    it remains.
    
    Fixes #17873.
    For #8984.
    
    Change-Id: Ifcf2bf1f84f730e6060f3517d63b45eb16215ae1
    Reviewed-on: https://go-review.googlesource.com/33098
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 0457957c991dde4bbdeefb73bc9fb01827298bd9
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 23:16:14 2016 +0000

    net/http: update bundled http2 for ErrAbortHandler support, document it more
    
    Updates http2 to x/net/http2 git rev 0e2717d for:
    
       http2: conditionally log stacks from panics in Server Handlers like net/http
       https://golang.org/cl/33102
    
    Fixes #17790
    
    Change-Id: Idd3f0c65540398d41b412a33f1d80de3f7f31409
    Reviewed-on: https://go-review.googlesource.com/33103
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit a501fef3455a8e0ff0424bb29a9403d7539c6164
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 23:03:06 2016 +0000

    net/http: deflake TestClientTimeout
    
    This test was only enabled by default today so it hasn't been hardened
    by build.golang.org. Welcome to the ring, TestClientTimeout.
    
    Change-Id: I1967f6c825699f13f6c659dc14d3c3c22b965272
    Reviewed-on: https://go-review.googlesource.com/33101
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit caa434d28063b4532bc362d50285230597d6d1f6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 21:16:31 2016 +0000

    net/http: update Transport doc example to not disable http2
    
    The old Transport example ended up disabling HTTP/2.
    
    Use a better example.
    
    Fixes #17051
    Fixes #17296
    
    Change-Id: I6feca168744131916e8bf56c829b4d4b50e304ee
    Reviewed-on: https://go-review.googlesource.com/33094
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b20c055230a20663c75e3099f672c15c39d46b9e
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 22:53:16 2016 +0000

    net/http: update bundled http2
    
    Updates http2 to x/net/http2 git rev 9ef22118 for:
    
       http2: fix CloseNotify data race
       https://golang.org/cl/33013
    
       http2: don't overflow stream IDs in server push
       https://golang.org/cl/32488
    
       http2: disable server push on receiving a GOAWAY
       https://golang.org/cl/32887
    
       http2: fix state tracking for pushed streams
       https://golang.org/cl/32755
    
    Change-Id: Ie7d675857423c102c9ec164d3c943093c749c7cf
    Reviewed-on: https://go-review.googlesource.com/33100
    Reviewed-by: Tom Bergan <tombergan@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9997545a8626bf1a73002f44a7b7538988da4e76
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 22:49:16 2016 +0000

    net/http: add ErrAbortHandler, make Server quiet if used as panic value
    
    Add an explicit way for Handlers to abort their response to the client
    and also not spam their error log with stack traces.
    
    panic(nil) also worked in the past (for http1 at least), so continue
    to make that work (and test it). But ErrAbortHandler is more explicit.
    
    Updates #17790 (needs http2 updates also)
    
    Change-Id: Ib1456905b27e2ae8cf04c0983dc73e314a4a751e
    Reviewed-on: https://go-review.googlesource.com/33099
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 1c54119315d9f3bd9212c01db2fd4653314959e0
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 21:35:01 2016 +0000

    net/http: document that Server.Close and Shutdown don't track hijacked conns
    
    Fixes #17721
    
    Change-Id: I19fd81c9909a22b01a4dc9c75f3f0e069c8608ca
    Reviewed-on: https://go-review.googlesource.com/33095
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c9ed065fbb7514a5da52b92575576cf359aead73
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Nov 10 12:49:06 2016 -0800

    debug/elf: SPARC64 relocation type is only 8 bits
    
    https://docs.oracle.com/cd/E53394_01/html/E54813/chapter6-54839.html#OSLLGchapter6-24:
    
    "For 64–bit SPARC Elf64_Rela structures, the r_info field is further
    broken down into an 8–bit type identifier and a 24–bit type dependent
    data field. For the existing relocation types, the data field is
    zero. New relocation types, however, might make use of the data bits.
    
     #define ELF64_R_TYPE_ID(info)         (((Elf64_Xword)(info)<<56)>>56)
    "
    
    No test for this because the only test would be an invalid object file.
    
    Change-Id: I5052ca3bfaf0759e920f9a24a16fd97543b24486
    Reviewed-on: https://go-review.googlesource.com/33091
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit 63224cab54055eaf1b3af62f3acaf64ff304316c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 22:12:50 2016 +0000

    net/http: document and deprecate type and errors of type ProtocolError
    
    Clean up & document the ProtocolError gunk.
    
    Fixes #17558
    
    Change-Id: I5e54c25257907c9cac7433f7a5bdfb176e8c3eee
    Reviewed-on: https://go-review.googlesource.com/33096
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 41027cc460b8db8fe2c23fa6fb97eb1ddab44799
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Nov 10 12:58:16 2016 -0800

    cmd/go: remove "x" in TestImportMain
    
    Interestingly, this only became a problem when CL 32850 marked
    TestImportMain as parallel.  Before that, "x" was overwritten and remove
    in a later test, TestGoBuildOutput.  The latter test is not marked as
    parallel, so now it is run first.  It is rather fragile for two tests to
    compete over the same filename, but this change is correct regardless.
    
    Change-Id: I1db7929c0bc20a2fd0cc6a02999bef2dca9e0cc0
    Reviewed-on: https://go-review.googlesource.com/33092
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit ab0ae44e91efec4440828406ef1929488a8b1b06
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 20:28:14 2016 +0000

    net/http: fix documentation on Server.TLSNextProto controlling http2
    
    Server.TLSNextProto being nil is necessary but not sufficient but
    http2 being automatically enabled.
    
    Fixes #16588
    
    Change-Id: I5b18690582f9b12ef05b58235e1eaa52483be285
    Reviewed-on: https://go-review.googlesource.com/33090
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 76f12cdaa2be0b96d314762ff5b2e403d1359cd8
Author: Michael Matloob <matloob@golang.org>
Date:   Thu Nov 10 13:31:41 2016 -0500

    runtime/pprof: output CPU profiles in pprof protobuf format
    
    This change buffers the entire profile and converts in one shot
    in the profile writer, and could use more memory than necessary
    to output protocol buffer formatted profiles. It should be
    possible to convert each chunk in a stream (maybe maintaining
    some minimal state to output in the end) which could save on
    memory usage.
    
    Fixes #16093
    
    Change-Id: I946c6a2b044ae644c72c8bb2d3bd82c415b1a847
    Reviewed-on: https://go-review.googlesource.com/33071
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 7448eb4172bfc8f704b9ea39d77d0113a042b9dc
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Tue Nov 1 00:44:48 2016 -0700

    net/http: don't wrap request cancellation errors in timeouts
    
    Based on Filippo Valsorda's https://golang.org/cl/24230
    
    Fixes #16094
    
    Change-Id: Ie39b0834e220f0a0f4fbfb3bbb271e70837718c3
    Reviewed-on: https://go-review.googlesource.com/32478
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 47bdae942242eca4be94989bab485bb1335f354d
Author: Francesc Campoy <campoy@golang.org>
Date:   Mon Nov 7 23:37:21 2016 -0800

    cmd/vet: detect defer resp.Body.Close() before error check
    
    This check detects the code
    
            resp, err := http.Get("http://foo.com")
            defer resp.Body.Close()
            if err != nil {
                    ...
            }
    
    For every call to a function on the net/http package or any method
    on http.Client that returns (*http.Response, error), it checks
    whether the next line is a defer statement that calls on the response.
    
    Fixes #17780.
    
    Change-Id: I9d70edcbfa2bad205bf7f45281597d074c795977
    Reviewed-on: https://go-review.googlesource.com/32911
    Reviewed-by: Rob Pike <r@golang.org>

commit 91135f27415ae582081e59d1b70f77d4d4d58112
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 17:49:19 2016 +0000

    A+C: update for Go 1.8
    
    Add Albert Nigmatzianov (individual CLA)
    Add Alex Browne (individual CLA)
    Add Alex Carol (individual CLA)
    Add Alexander Döring (individual CLA)
    Add Allan Simon (individual CLA)
    Add Alok Menghrajani (individual CLA)
    Add Andreas Auernhammer (individual CLA)
    Add Andreas Litt (individual CLA)
    Add Andrew Pogrebnoy (individual CLA)
    Add Antonio Murdaca (corporate CLA for Red Hat, Inc.)
    Add Atin Malaviya (individual CLA)
    Add Billy Lynch (corporate CLA for Google Inc.)
    Add Blixt (individual CLA)
    Add Boris Nagaev (corporate CLA for Google Inc.)
    Add Braden Bassingthwaite (corporate CLA for Vendasta)
    Add Brian Kennedy (individual CLA)
    Add Bryan Alexander (individual CLA)
    Add Carl Johnson (individual CLA)
    Add Cixtor (individual CLA)
    Add Cyrill Schumacher (individual CLA)
    Add Daniel Martí (individual CLA)
    Add Daria Kolistratova (corporate CLA for Intel Corporation)
    Add David Hubbard (corporate CLA for Google Inc.)
    Add David Stainton (individual CLA)
    Add Deepak Jois (individual CLA)
    Add Denis Nagorny (corporate CLA for Intel Corporation)
    Add Dhaivat Pandit (individual CLA)
    Add Dhananjay Nakrani (corporate CLA for Google Inc.)
    Add Dmitri Popov (individual CLA)
    Add Erik Staab (corporate CLA for Google Inc.)
    Add Ethan Miller (corporate CLA for IBM)
    Add Faiyaz Ahmed (individual CLA)
    Add Fedor Indutny (individual CLA)
    Add Gabriel Russell (individual CLA)
    Add Gareth Paul Jones (individual CLA)
    Add Geoffroy Lorieux (individual CLA)
    Add Gleb Stepanov (individual CLA)
    Add Henrik Hodne (individual CLA)
    Add Ivan Babrou (individual CLA)
    Add Jack Lindamood (corporate CLA for Amazon.com, Inc)
    Add James Clarke (individual CLA)
    Add Jamie Beverly (individual CLA)
    Add Jason Smale (individual CLA)
    Add Jean-Nicolas Moal (individual CLA)
    Add Jeroen Bobbeldijk (individual CLA)
    Add Jim Kingdon (corporate CLA for Bolt)
    Add Jirka Daněk (individual CLA)
    Add Jon Chen (corporate CLA for Amazon.com, Inc)
    Add Joonas Kuorilehto (individual CLA)
    Add Josh Chorlton (individual CLA)
    Add Joshua Boelter (corporate CLA for Intel Corporation)
    Add Justyn Temme (individual CLA)
    Add Kale Blankenship (individual CLA)
    Add LE Manh Cuong (individual CLA)
    Add Luigi Riefolo (individual CLA)
    Add Manfred Touron (individual CLA)
    Add Martin Bertschler (individual CLA)
    Add Martin Hamrle (individual CLA)
    Add Matthew Denton (individual CLA)
    Add Matthieu Hauglustaine (individual CLA)
    Add Michael Darakananda (corporate CLA for Google Inc.)
    Add Mike Appleby (individual CLA)
    Add Mike Houston (individual CLA)
    Add Mike Strosaker (corporate CLA for IBM)
    Add Miroslav Genov (individual CLA)
    Add Momchil Velikov (individual CLA)
    Add Nick Harper (corporate CLA for Google Inc.)
    Add Oleg Vakheta (individual CLA)
    Add Parker Moore (individual CLA)
    Add Prasanna Swaminathan (corporate CLA for MediaMath, Inc)
    Add Radu Berinde (individual CLA)
    Add Ramesh Dharan (corporate CLA for Google Inc.)
    Add Richard Gibson (individual CLA)
    Add Samuel Tan (corporate CLA for Google Inc.)
    Add Samuele Pedroni (individual CLA)
    Add Sarah Adams (corporate CLA for Google Inc.)
    Add Sean Rees (individual CLA)
    Add Simon Rawet (individual CLA)
    Add Sina Siadat (individual CLA)
    Add Song Gao (individual CLA)
    Add Suyash (individual CLA)
    Add Sven Blumenstein (corporate CLA for Google Inc.)
    Add Syohei YOSHIDA (individual CLA)
    Add Terrel Shumway (individual CLA)
    Add Than McIntosh (corporate CLA for Google Inc.)
    Add Thomas de Zeeuw (individual CLA)
    Add Tim Henderson (individual CLA)
    Add Tom Wilkie (corporate CLA for Weaveworks)
    Add Trey Lawrence (individual CLA)
    Add Tristan Ooohry (individual CLA)
    Add Tuo Shan (corporate CLA for Google Inc.)
    Add Victor Chudnovsky (corporate CLA for Google Inc.)
    Add Vitor De Mario (individual CLA)
    Add Vladimir Mihailenco (individual CLA)
    Add Vladimir Stefanovic (individual CLA)
    Add Walter Poupore (corporate CLA for Google Inc.)
    Add Xuyang Kang (individual CLA)
    Add Zev Goldstein (individual CLA)
    
    Updates #12042
    
    Change-Id: I28d63babe225683b88f3f1501e529aed636c9ead
    Reviewed-on: https://go-review.googlesource.com/33028
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d54463f4fc9e6137b31f660afa5a785ba65e1879
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 18:00:10 2016 +0000

    lib/time: update tzdata to 2016i
    
    Fixes #17678
    
    Change-Id: I01d12a827e6106efed1ec024f736c640b86906b4
    Reviewed-on: https://go-review.googlesource.com/33029
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b77bff97c4817656dd3bf49f14d1834c411be084
Author: Ian Lance Taylor <iant@golang.org>
Date:   Thu Nov 10 09:59:42 2016 -0800

    cmd/go: -ldflags=-linkmode=external requires runtime/cgo
    
    We add runtime/cgo to the list of import paths for various cases that
    imply external linking mode, but before this change we did not add for
    an explicit request of external linking mode. This fixes the case where
    you are using a non-default buildmode that implies a different
    compilation option (for example, -buildmode=pie implies -shared) and the
    runtime/cgo package for that option is stale.
    
    No test, as I'm not sure how to write one. It would require forcing a
    stale runtime/cgo.
    
    Change-Id: Id0409c7274ce67fe15d910baf587d3220cb53d83
    Reviewed-on: https://go-review.googlesource.com/33070
    Reviewed-by: Michael Hudson-Doyle <michael.hudson@canonical.com>

commit a0d2e9699f1cc83a854c52843ff15d07f83bce47
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Nov 9 14:15:59 2016 -0800

    go/printer: don't drop required semi/linebreak after /*-comment
    
    For details, see the issues.
    
    Fixes #11274.
    Fixes #15137.
    
    Change-Id: Ia11e71a054b3195e3007f490418a9c53a7e9cdf1
    Reviewed-on: https://go-review.googlesource.com/33016
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 8cd55615d4d00f48c30dff85d5d5e3f2adce70ce
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 10 16:43:15 2016 +0000

    net/http: fix Server.Close double Lock
    
    Fixes #17878
    
    Change-Id: I062ac514239068c58175c9ee7964b3590f956a82
    Reviewed-on: https://go-review.googlesource.com/33026
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8d0c105407d235c85a163c0cda5bda86e5219c36
Author: David Crawshaw <crawshaw@golang.org>
Date:   Fri Nov 4 18:22:06 2016 -0400

    reflect: unexported fields are tied to a package
    
    An unexported field of a struct is not visible outside of the package
    that defines it, so the package path is implicitly part of the
    definition of any struct with an unexported field.
    
    Change-Id: I17c6aac822bd0c24188ab8ba1cc406d6b5d82771
    Reviewed-on: https://go-review.googlesource.com/32820
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9e2c3f4c7e78b01d635e16287789913d02807569
Author: Kevin Burke <kev@inburke.com>
Date:   Wed Jun 22 10:00:31 2016 -0700

    sync: add example for Pool
    
    It was a little tricky to figure out how to go from the documentation
    to figuring out the best way to implement a Pool, so I thought I'd
    try to provide a simple example. The implementation is mostly taken
    from the fmt package.
    
    I'm not happy with the verbosity of the calls to WriteString() etc,
    but I wanted to provide a non-trivial example.
    
    Change-Id: Id33a8b6cbf8eb278f71e1f78e20205b436578606
    Reviewed-on: https://go-review.googlesource.com/24371
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 73497c7656fa55ac33bac960ecee806b9b07ae5e
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Nov 9 17:02:33 2016 -0800

    cmd/gofmt: don't leave tmp file if -w failed
    
    Follow-up on https://golang.org/cl/33018.
    
    For #8984.
    
    Change-Id: I6655a5537a60d4ea3ee13029a56a75b150f8c8f8
    Reviewed-on: https://go-review.googlesource.com/33020
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b188b4cc110261a004674df5a4e209cc4894d314
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Nov 9 15:29:41 2016 -0800

    cmd/gofmt: don't eat source if -w fails
    
    Write output to a temp file first and only upon success
    rename that file to source file name.
    
    Fixes #8984.
    
    Change-Id: Ie40e49d2a4eb3c9462fe769ccbf055b4366eceb0
    Reviewed-on: https://go-review.googlesource.com/33018
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit add8028eb9925db0ec9d3d0f95dec65c8cab1d96
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Nov 9 16:11:50 2016 -0800

    go/types: remove unused alias-related testdata files
    
    They interfere with gofmt -w across this directory.
    
    Follow-up on https://go-review.googlesource.com/32819.
    
    For #16339 (comment).
    
    Change-Id: I4298b6117d89517d4fe6addce3942d950d821817
    Reviewed-on: https://go-review.googlesource.com/33019
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2f2b57853eeccf854e8e47e5f77ff043d34b2a34
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 9 21:35:37 2016 +0000

    net/http: deflake TestIdleConnH2Crash
    
    Fixes #17838
    
    Change-Id: Ifafb4542a0ed6f2e29c9a83e30842e2fc18d6546
    Reviewed-on: https://go-review.googlesource.com/33015
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Tom Bergan <tombergan@google.com>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit 22c70f268b30c703856143df848556515a824071
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 9 20:35:46 2016 +0000

    syscall: use 32-bit setuid/setgid syscalls on linux/{386,arm}
    
    Fixes #17092
    
    Change-Id: Ib14e4db13116ebbe4d72c414fb979d27a06d6174
    Reviewed-on: https://go-review.googlesource.com/33011
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 48c6048e554ff4f428aefd41b9345ed5ec634783
Author: Quentin Smith <quentin@golang.org>
Date:   Tue Nov 8 16:47:04 2016 -0500

    encoding/xml: check type when unmarshaling innerxml field
    
    We only support unmarshaling into a string or a []byte, but we
    previously would try (and panic while) setting a slice of a different
    type. The docs say ",innerxml" is ignored if the type is not string or
    []byte, so do that for other slices as well.
    
    Fixes #15600.
    
    Change-Id: Ia64815945a14c3d04a0a45ccf413e38b58a69416
    Reviewed-on: https://go-review.googlesource.com/32919
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 9c2037fbcf1a732f55e29062f3d30ddd21ca36d3
Author: Michael Munday <munday@ca.ibm.com>
Date:   Tue Nov 8 17:11:18 2016 -0500

    runtime/pprof/internal/protopprof: skip TestTranslateCPUProfileWithSamples if < 2 mappings
    
    A Go binary may only have 1 executable memory region if it has been
    linked using internal linking. This change means that the test will
    be skipped if this is the case, rather than fail.
    
    Fixes #17852.
    
    Change-Id: I59459a0f90ae8963aeb9908e5cb9fb64d7d0e0f4
    Reviewed-on: https://go-review.googlesource.com/32920
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Matloob <matloob@golang.org>

commit 60a9bf9f957d48856839873c6dcb699afe7da359
Author: Robert Griesemer <gri@golang.org>
Date:   Tue Nov 8 16:01:56 2016 -0800

    cmd/compile/internal/syntax: fix error handling for Read/Parse calls
    
    - define syntax.Error for cleaner error reporting
    - abort parsing after first error if no error handler is installed
    - make sure to always report the first error, if any
    - document behavior of API calls
    - while at it: rename ReadXXX -> ParseXXX (clearer)
    - adjust cmd/compile noder.go accordingly
    
    Fixes #17774.
    
    Change-Id: I7893eedea454a64acd753e32f7a8bf811ddbb03c
    Reviewed-on: https://go-review.googlesource.com/32950
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit ad020477f4dfe731450b6dd3dd15ea43aab0d0f1
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sun Nov 6 11:58:49 2016 -0800

    cmd/cgo: delete unused variable in log statement
    
    visit is just a func, and there's no formatting
    verb for it, and it's on an internal-error path.
    It has been thus many years, unchanged and unexecuted.
    
    Change-Id: I4c2e2673ee9996218c24143bcc3be3eb4abdff25
    Reviewed-on: https://go-review.googlesource.com/32970
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit ef462153de8f756a6f214ec24ccc541fe99f58ad
Author: Martin Möhrmann <moehrmann@google.com>
Date:   Wed Nov 9 09:24:50 2016 +0100

    C: add Martin Möhrmann's google.com email (Google CLA)
    
    Change-Id: Ia439c4a3c873ef24f60f8ee54a74f767fdaafd29
    Reviewed-on: https://go-review.googlesource.com/32799
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5af7553f9d8dbba8d798cadf62c10e504f5c3d36
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Tue Nov 8 16:10:26 2016 -0800

    cmd/compile: ensure that knownFormats is up to date
    
    Change-Id: I4febdddfe5be569a8bba0a4cddf52dec7f1be1bf
    Reviewed-on: https://go-review.googlesource.com/32930
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 2925427a47f41622f28f84889ad7aade27581144
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 8 01:06:06 2016 +0000

    os: on Windows, don't fix long paths that aren't long
    
    Notably, don't allocate.
    
    Follow-up to https://golang.org/cl/32451 which added long path
    cleaning.
    
    Updates #3358
    
    Change-Id: I89c59cbd660d0a030f31b6acd070fa9f3250683b
    Reviewed-on: https://go-review.googlesource.com/32886
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 76d8e60451c7b51eb7ea362e7ce3f66050864493
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:47 2016 +0200

    cmd/link: add support for GOARCH=mips{,le}
    
    Only internal linking without cgo is supported for now.
    
    Change-Id: I772d2ba496a613c78bee7e93f29e9538e6407bdc
    Reviewed-on: https://go-review.googlesource.com/31481
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit f4c997578ad962074b8e3aaf28dcfaa2f48a8593
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:44 2016 +0200

    cmd/compile: add support for GOARCH=mips{,le}
    
    Change-Id: Ib489dc847787aaab7ba1be96792f885469e346ae
    Reviewed-on: https://go-review.googlesource.com/31479
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 247fc4a98ee7e41bc317ba2245516f279927ae65
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:42 2016 +0200

    cmd/compile/internal/ssa: add support for GOARCH=mips{,le}
    
    Change-Id: I632d4aef7295778ba5018d98bcb06a68bcf07ce1
    Reviewed-on: https://go-review.googlesource.com/31478
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit f72a629dbd6d4ed65cb295b896de214b0f7790d5
Author: Michael Matloob <matloob@golang.org>
Date:   Fri Nov 4 11:27:54 2016 -0400

    runtime/pprof/internal: add package protopprof
    
    This change adds code, originally written by Russ Cox <rsc@golang.org>
    and open-sourced by Google, that converts from the "legacy"
    binary pprof profile format to a struct representation of the
    new protocol buffer pprof profile format.
    
    This code reads the entire binary format for conversion to the
    protobuf format. In a future change, we will update the code
    to incrementally read and convert segments of the binary format,
    so that the entire profile does not need to be stored in memory.
    
    This change also contains contributions by Daria Kolistratova
    <daria.kolistratova@intel.com> from the rolled-back change
    golang.org/cl/30556 adapting the code to be used by the package
    runtime/pprof.
    
    This code also appeared in the change golang.org/cl/32257, which was based
    on Daria Kolistratova's change, but was also rolled back.
    
    Updates #16093
    
    Change-Id: I5c768b1134bc15408d80a3ccc7ed867db9a1c63d
    Reviewed-on: https://go-review.googlesource.com/32811
    Run-TryBot: Michael Matloob <matloob@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 7465bfb1ea1a7ddfec9b267587ee9e6200514f3f
Author: Jaana Burcu Dogan <jbd@google.com>
Date:   Mon Oct 31 10:23:50 2016 -0700

    path: document that filepath is recommended to manipulate filename paths
    
    Fixes #17690.
    
    Change-Id: Ifd300980aa4c11498ed7c083d08bcdd23f5b307a
    Reviewed-on: https://go-review.googlesource.com/32423
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 59d5835f144fae0c9e8726ffc7e42663eadef229
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 8 18:23:43 2016 +0000

    doc: add a CL to go1.8.txt mentioned by Alberto Donizetti
    
    Change-Id: I43617e6dfd5b8227a8ef907dc22c00188de87b94
    Reviewed-on: https://go-review.googlesource.com/32915
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b990558162fa038f3651dc0f1821f64b282dda6f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Oct 28 18:03:21 2016 +0000

    doc: reference go1.4-bootstrap-20161024.tar.gz
    
    Updates #16352
    
    Change-Id: I214c87579ef21ced8d0ba94aa170dd7780afec4b
    Reviewed-on: https://go-review.googlesource.com/32312
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f58ce7fe7954bd788072beaf1517303ebb5316eb
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:38 2016 +0200

    cmd/asm: add support for GOARCH=mips{,le}
    
    Change-Id: I6a5256a42f895bb93ac56764e91ade1861c00e04
    Reviewed-on: https://go-review.googlesource.com/31476
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 27a3d30dd09cdd869b1b67f0154fa698bdf8ead2
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 8 17:44:53 2016 +0000

    net/http: deflake TestClientRedirects
    
    Fix another case of a parallel test relying on a global variable
    (DefaultTransport) implicitly.
    
    Use the private Transport already in scope instead. It's closed at the
    end, instead of randomly via another test.
    
    Change-Id: I95e51926177ad19a766cabbb306782ded1bbb59b
    Reviewed-on: https://go-review.googlesource.com/32913
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 41eb9bb993b22979eebffe4eaeeca53db8e7b388
Author: Volodymyr Paprotski <vpaprots@ca.ibm.com>
Date:   Fri Oct 14 16:19:25 2016 -0400

    crypto/elliptic: add s390x assembly implementation of NIST P-256 Curve
    
    A paranoid go at constant time implementation of P256 curve.
    
    This code relies on z13 SIMD instruction set. For zEC12 and below,
    the fallback is the existing P256 implementation. To facilitate this
    fallback mode, I've refactored the code so that implementations can
    be picked at run-time.
    
    Its 'slightly' difficult to grok, but there is ASCII art..
    
    name            old time/op  new time/op  delta
    BaseMultP256     419µs ± 3%    27µs ± 1%  -93.65% (p=0.000 n=10+8)
    ScalarMultP256  1.05ms ±10%  0.09ms ± 1%  -90.94% (p=0.000 n=10+8)
    
    Change-Id: Ic1ded898a2ceab055b1c69570c03179c4b85b177
    Reviewed-on: https://go-review.googlesource.com/31231
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 5d28bc58b6524b2043e2864b8de99fb05e7160d5
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:37 2016 +0200

    cmd/internal/obj/mips: add support for GOARCH=mips{,le}
    
    Implements subset of MIPS32(r1) instruction set.
    
    Change-Id: Iba017350f6c2763de05d4d1bc2f123e8eb76d0ff
    Reviewed-on: https://go-review.googlesource.com/31475
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit c41137d2428587b7e6df09952cddabd6441bc4b7
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Tue Nov 8 02:25:23 2016 +0000

    syscall: fix name of prlimit parameters
    
    Fixes #17606
    
    Change-Id: I040c7621cef265d44b58f16556e6d58660a2245d
    Reviewed-on: https://go-review.googlesource.com/32889
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3a3f672eda6f720f68950492a85254528ac53dad
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Tue Nov 8 17:46:10 2016 +0530

    os: cleanup directories created by TestLongPath
    
    Add tmpdir as a parameter to the closure otherwise the subsequent
    modifications to tmpdir causes only the last subdirectory to be
    removed.
    
    Additionally, add the missing argument for the t.Fatalf call.
    
    Change-Id: I3df53f9051f7ea40cf3f846d47d9cefe445e9b9d
    Reviewed-on: https://go-review.googlesource.com/32892
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 8a2a999311c22079c3b9f2e6fac2bbd38435a7ab
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Nov 7 15:56:41 2016 -0800

    go/types: document that selectors are not recorded in Info.Types
    
    Fixes #11944.
    
    Change-Id: I424ba93725f22fd599e052eb182f9ba2fca8e8bd
    Reviewed-on: https://go-review.googlesource.com/32881
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 248a59447165ebac2779cb54ee2a10c021009d20
Author: Shenghou Ma <minux@golang.org>
Date:   Mon Nov 7 19:08:51 2016 -0500

    doc/devel/release.html: document go1.6.3 doesn't actually support macOS Sierra
    
    Updates #17824.
    
    Change-Id: I73cf89c21b418158c7014c3271cd1103a17a5c86
    Reviewed-on: https://go-review.googlesource.com/32882
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit de847be6a4ca3e5d22c056e09aadc30a0a777a08
Author: Shenghou Ma <minux@golang.org>
Date:   Mon Nov 7 19:16:17 2016 -0500

    doc/go1.8.txt: mention os.Executable addition
    
    Change-Id: Id3d571666b9275e3fa5cb20762afbd391dbcdeba
    Reviewed-on: https://go-review.googlesource.com/32883
    Reviewed-by: Minux Ma <minux@golang.org>

commit 2fc67e71af142bfa1e7662a4fde361f43509d2d7
Author: Shenghou Ma <minux@golang.org>
Date:   Sun Nov 1 04:18:58 2015 -0500

    os: add Executable() (string, error)
    
    // Executable returns the path name for the executable that started
    // the current process. There is no guarantee that the path is still
    // pointing to the correct executable. If a symlink was used to start
    // the process, depending on the operating system, the result might
    // be the symlink or the path it pointed to. If a stable result is
    // needed, path/filepath.EvalSymlinks might help.
    //
    // Executable returns an absolute path unless an error occurred.
    //
    // The main use case is finding resources located relative to an
    // executable.
    //
    // Executable is not supported on nacl or OpenBSD (unless procfs is
    // mounted.)
    func Executable() (string, error) {
            return executable()
    }
    
    Fixes #12773.
    
    Change-Id: I469738d905b12f0b633ea4d88954f8859227a88c
    Reviewed-on: https://go-review.googlesource.com/16551
    Run-TryBot: Minux Ma <minux@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 119c30eaf24c3d4f46ba1411f5cddbabb86bd840
Author: Shenghou Ma <minux@golang.org>
Date:   Sun Nov 1 04:17:34 2015 -0500

    internal/syscall/windows: add GetModuleFileName
    
    For os.Executable. Updates #12773.
    
    Change-Id: Iff6593514b7453b6c5e1f20079e35cb4992cc74a
    Reviewed-on: https://go-review.googlesource.com/32877
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit a9a1d020ec4a4deb417160a091c0ed41123063bc
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Nov 7 15:57:04 2016 -0500

    cmd/internal/sys, runtime/internal/sys: gofmt
    
    Change-Id: Ice8f3b42194852f7ee8f00f004e80014d1ea119b
    Reviewed-on: https://go-review.googlesource.com/32875
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit c2917af6280ea1bcc5ebf85224055ae1a1882af3
Author: Quentin Smith <quentin@golang.org>
Date:   Thu Nov 3 18:45:01 2016 -0400

    cmd/go: handle escapes in pkg-config output
    
    This commit also adds a test for pkg-config usage in cgo.
    
    Fixes #16455.
    
    Change-Id: I95fb6a288a4d19093c4613c93878017d95cbe4a2
    Reviewed-on: https://go-review.googlesource.com/32735
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 47d1c42aff6bb84a654cc69b1dbc42f855b03415
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Nov 7 15:40:48 2016 -0500

    crypto/tls: use default cipher suites in BenchmarkThroughput
    
    CL 32871 updated the default cipher suites to use AES-GCM in
    preference to ChaCha20-Poly1305 on platforms which have hardware
    implementations of AES-GCM. This change makes BenchmarkThroughput
    use the default cipher suites instead of the test cipher suites to
    ensure that the recommended (fastest) algorithms are used.
    
    Updates #17779.
    
    Change-Id: Ib551223e4a00b5ea197d4d73748e1fdd8a47c32d
    Reviewed-on: https://go-review.googlesource.com/32838
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Adam Langley <agl@golang.org>

commit 1a279b34f674369c05694bd2d8e493ec2d3bba97
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Nov 7 11:20:48 2016 -0800

    go/constant: follow-up for https://go-review.googlesource.com/32870
    
    For #17812.
    
    Change-Id: I58411aaa0e8b2250a16ddb20c951e39da3d601e8
    Reviewed-on: https://go-review.googlesource.com/32872
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 231aa9d6d7e942ac43130356ba3554972251f015
Author: Quentin Smith <quentin@golang.org>
Date:   Fri Oct 28 13:01:51 2016 -0400

    os: use extended-length paths on Windows when possible
    
    Windows has a limit of 260 characters on normal paths, but it's possible
    to use longer paths by using "extended-length paths" that begin with
    `\\?\`. This commit attempts to transparently convert an absolute path
    to an extended-length path, following the subtly different rules those
    paths require. It does not attempt to handle relative paths, which
    continue to be passed to the operating system unmodified.
    
    This adds a new test, TestLongPath, to the os package. This test makes
    sure that it is possible to write a path at least 400 characters long
    and runs on every platform. It also tests symlinks and hardlinks, though
    symlinks are not testable with our builder configuration.
    
    HasLink is moved to internal/testenv so it can be used by multiple tests.
    
    https://msdn.microsoft.com/en-us/library/windows/desktop/aa365247(v=vs.85).aspx
    has Microsoft's documentation on extended-length paths.
    
    Fixes #3358.
    Fixes #10577.
    Fixes #17500.
    
    Change-Id: I4ff6bb2ef9c9a4468d383d98379f65cf9c448218
    Reviewed-on: https://go-review.googlesource.com/32451
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 2058511e4e5966a7d482beb6033c68e324aa09ac
Author: Shenghou Ma <minux@golang.org>
Date:   Sun Nov 1 04:16:52 2015 -0500

    runtime: os.Executable runtime support for Darwin
    
    Change-Id: Ie21df37016c90cd0479c23ec4845f8195dd90fda
    Reviewed-on: https://go-review.googlesource.com/16518
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit f335fcf0fc297b9d1318fb21fa608a7ebc751b3e
Author: Shenghou Ma <minux@golang.org>
Date:   Sun Nov 1 04:18:26 2015 -0500

    syscall: add Getexecname on Solaris for os.Executable
    
    Change-Id: Icd77ccbfe6a31117a11effb949b5826950df75a9
    Reviewed-on: https://go-review.googlesource.com/16550
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit a9ce0f96e1f2ab69ce3319c5a97c1d01beb9472c
Author: Adam Langley <agl@golang.org>
Date:   Mon Nov 7 10:25:57 2016 -0800

    crypto/{cipher,tls,internal/cryptohw}: prioritise AES-GCM when hardware support is present.
    
    Support for ChaCha20-Poly1305 ciphers was recently added to crypto/tls.
    These ciphers are preferable in software, but they cannot beat hardware
    support for AES-GCM, if present.
    
    This change moves detection for hardware AES-GCM support into
    cipher/internal/cipherhw so that it can be used from crypto/tls. Then,
    when AES-GCM hardware is present, the AES-GCM cipher suites are
    prioritised by default in crypto/tls. (Some servers, such as Google,
    respect the client's preference between AES-GCM and ChaCha20-Poly1305.)
    
    Fixes #17779.
    
    Change-Id: I50de2be486f0b0b8052c4628d3e3205a1d54a646
    Reviewed-on: https://go-review.googlesource.com/32871
    Run-TryBot: Adam Langley <agl@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 9e4a70e8fd3e6fed93fcd6159228b3a8ecae1d80
Author: Quentin Smith <quentin@golang.org>
Date:   Tue Oct 25 15:41:14 2016 -0400

    mime/multipart: test for overreading on a stream
    
    Some multipart data arrives in a stream, where subsequent parts may not
    be ready yet. Read should return a complete part as soon as
    possible.
    
    Fixes #15431
    
    Change-Id: Ie8c041b853f3e07f0f2a66fbf4bcab5fe9132a7c
    Reviewed-on: https://go-review.googlesource.com/32032
    Run-TryBot: Quentin Smith <quentin@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 44febb28b626695e2daa693b1aa9ad7e516518e3
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 7 01:21:15 2016 +0000

    cmd/go: parallelize some tests
    
    Cuts tests from 35 to 25 seconds.
    
    Many of these could be parallel if the test runner were modified to
    give each test its own workdir cloned from the tempdir files they
    use. But later. This helps for now.
    
    Updates #17751
    
    Change-Id: Icc2ff87cca60a33ec5fd8abb1eb0a9ca3e85bf95
    Reviewed-on: https://go-review.googlesource.com/32850
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4eb9832724e94d04cd737ac2554cf02a79d87f23
Author: Robert Griesemer <gri@golang.org>
Date:   Mon Nov 7 10:43:25 2016 -0800

    go/constant: improved fatal error messages
    
    Fixes #17812.
    
    Change-Id: I08202165dd3f72ae04420e7b6129b8b689e74f5c
    Reviewed-on: https://go-review.googlesource.com/32870
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 9d139ac3fa1f4c0c468042bdc6248a30044fb1bb
Author: Ian Lance Taylor <iant@golang.org>
Date:   Sun Nov 6 10:35:58 2016 -0800

    unsafe: remove incorrect type conversion in docs
    
    Fixes #17818.
    
    Change-Id: Id7242b0bdd5e1db254b44ae29900fc4f3362c743
    Reviewed-on: https://go-review.googlesource.com/32828
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f815499fab5f650f3eba12d48c53ec01d1362265
Author: Joe Farrell <joe2farrell@gmail.com>
Date:   Mon Nov 7 18:43:13 2016 +0000

    doc: fix broken links in 1.7 release docs
    
    Change-Id: Ibf73ee7be4591393f4e08d464edfa325c3ec2c11
    Reviewed-on: https://go-review.googlesource.com/32798
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 4cc83d49d2a69bb9d348cb32d71ca34651aec3b5
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Nov 7 11:50:48 2016 -0500

    cmd/dist: enable more cgo tests on ppc64le
    
    The tests all pass (for me at least) so I don't think there is any
    reason not to enable them.
    
    Change-Id: I96e71383e573273f442a849914cf6458ada14f82
    Reviewed-on: https://go-review.googlesource.com/32855
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 6c2a35ae0c29306b53a4a5925627ee7235a8c450
Author: Michael Munday <munday@ca.ibm.com>
Date:   Mon Nov 7 11:08:36 2016 -0500

    test/fixedbugs: enable issue 10607 test on ppc64le
    
    ppc64le supports both internal and external linking so I don't
    think there is any reason for it to skip this test.
    
    Change-Id: I05c80cc25909c0364f0a1fb7d20766b011ea1ebb
    Reviewed-on: https://go-review.googlesource.com/32854
    Reviewed-by: Lynn Boger <laboger@linux.vnet.ibm.com>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit ec77b8e09cace7338ebc39faadd599e5cb676295
Author: Lynn Boger <laboger@linux.vnet.ibm.com>
Date:   Fri Nov 4 14:53:59 2016 -0500

    cmd/link: don't use trampolines in ppc64le ext linking
    
    On ppc64x, trampolines are used to resolve too-far
    branches for internal linking.  The external linking,
    solution on ppc64x is to split text sections when they
    get too large, allowing the external linker to handle
    the long branches.
    
    On arm trampolines are generanted for too-far branches
    for internal and external linking.  When the change
    was made recently to enable trampolines for external linking
    on arm, that broke the ppc64x fix for too-far branches
    with external linking.
    
    The fix adds a check to use trampolines only for internal
    linking with ppc64x.
    
    Fixes #17795
    
    Change-Id: Icce968fb96545f10a913e07654514643bce96261
    Reviewed-on: https://go-review.googlesource.com/32853
    Run-TryBot: Lynn Boger <laboger@linux.vnet.ibm.com>
    Run-TryBot: Cherry Zhang <cherryyz@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: David Chase <drchase@google.com>

commit 3df059ece5d4c575abdf61b4b955f0ba292e5168
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Mon Nov 7 04:05:18 2016 +0000

    net/http/fcgi: fix link to protocol docs
    
    Fixes #17815
    
    Change-Id: I766082d28a14c77f5dfb6cd1974b86cb0a8fe31a
    Reviewed-on: https://go-review.googlesource.com/32852
    Reviewed-by: Minux Ma <minux@golang.org>

commit f8187ceacf886195dc9f72a6578bd266efc52681
Author: Alex Brainman <alex.brainman@gmail.com>
Date:   Sun Nov 6 18:43:49 2016 +1100

    runtime/race: allow TestFail to run longer than 0.00s
    
    Fixes #17811
    
    Change-Id: I7bf9cbc5245417047ad28a14d9b9ad6592607d3d
    Reviewed-on: https://go-review.googlesource.com/32774
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e0178025978470e9f7b5fa9365891d20db809a7c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 4 19:12:09 2016 +0000

    cmd/vet: parallelize tests
    
    Was 2.3 seconds. Now 1.4 seconds.
    
    Next win would be not running a child process and refactoring main so
    it could be called from tests easily. But that would also require
    rewriting the errchk written in Perl. This appears to be the last user
    of errchk in the tree.
    
    Updates #17751
    
    Change-Id: Id7c3cec76f438590789b994e756f55b5397be07f
    Reviewed-on: https://go-review.googlesource.com/32754
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 2b445c76453ee79d02ea50f50f619d8f7250fa87
Author: Robert Griesemer <gri@golang.org>
Date:   Sat Nov 5 13:14:23 2016 -0700

    go/constant: document that BinaryOp doesn't handle comparisons or shifts
    
    Fixes #17797.
    
    Change-Id: I544df81c4bcf3cbd36a793be40050f14f9a9974f
    Reviewed-on: https://go-review.googlesource.com/32761
    Reviewed-by: Dominik Honnef <dominik@honnef.co>

commit d0cf0421717a93b705efcbce0770a24361582445
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Sat Nov 5 10:05:27 2016 -0700

    net: fix vet nit
    
    net/fd_windows.go:121: syscall.WSABuf composite literal uses unkeyed fields
    
    Change-Id: I91cbe38199d5b6828379a854d08f6ceaf687dd82
    Reviewed-on: https://go-review.googlesource.com/32760
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit f9d406ebeaa0af42d2a88bab498a44a91c6c03ee
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Nov 5 00:25:38 2016 +0000

    net/http: deflake TestLinuxSendfile
    
    Fixes #17805
    
    Change-Id: I30d3e63a82b3690a76f2bb33d59ae34c62a7fa59
    Reviewed-on: https://go-review.googlesource.com/32759
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 276c29ff6da13ef7089075d8cd2172fe3d3956fb
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Nov 5 00:21:59 2016 +0000

    net/http: deflake TestClientRedirectTypes and maybe some similar ones
    
    A few tests were using the global DefaultTransport implicitly.
    Be explicit instead. And then make some parallel while I'm there.
    
    Change-Id: I3c617e75429ecc8f6d23567d1470f5e5d0cb7cfd
    Reviewed-on: https://go-review.googlesource.com/32758
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 65269e7066c0e97532bf0fd082e95cf8444ee378
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 16:37:24 2016 -0700

    cmd/go/internal/syntax: reintroduce reverted comments
    
    These comments were originally introduced together with the changes
    for alias declarations, and then reverted when we backed out alias
    support.
    
    Reintroduce them.
    
    Change-Id: I3ef2c4f4672d6af8a900f5d73df273edf28d1a14
    Reviewed-on: https://go-review.googlesource.com/32826
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 429edcff1049899cef4e3c5cb4b473e13a31d85c
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 16:27:31 2016 -0700

    Revert "cmd/compile/internal/syntax: support for alias declarations"
    
    This reverts commit 32db3f2756324616b7c856ac9501deccc2491239.
    
    Reason: Decision to back out current alias implementation.
    
    For #16339.
    
    Change-Id: Ib05e3d96041d8347e49cae292f66bec791a1fdc8
    Reviewed-on: https://go-review.googlesource.com/32825
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit a1a688fa0012f7ce3a37e9ac0070461fe8e3f28e
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 16:11:48 2016 -0700

    Revert "go/scanner, go/token: recognize => (ALIAS) token"
    
    This reverts commit 776a90100f1f65fcf54dfd3d082d657341bdc323.
    
    Reason: Decision to back out current alias implementation.
    
    For #16339.
    
    Change-Id: Icb451a122c661ded05d9293356b466fa72b965f3
    Reviewed-on: https://go-review.googlesource.com/32824
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2808f1f41583ef9943b1d7426d7affa6cfc998ba
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 16:04:11 2016 -0700

    Revert "go/ast, go/parser: parse alias declarations"
    
    This reverts commit 57ae83307fc4cb90338b39dcc6fe3c61ee8ce0b7.
    
    Reason: Decision to back out current alias implementation.
    
    For #16339.
    
    Change-Id: I7bcc04ac87ea3590999e58ff65a7f2e1e6c6bc77
    Reviewed-on: https://go-review.googlesource.com/32823
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit c350c5cfdc5a6ed1a135c06517ef29cdd716a935
Author: Leon Klingele <git@leonklingele.de>
Date:   Sat Nov 5 00:57:45 2016 +0100

    plugin: fix doc example fmt usage
    
    Change-Id: I0520a37a48a56d231a8ac2dc58b2bf1762282760
    Reviewed-on: https://go-review.googlesource.com/32795
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 7179c1acd967daea44fb806865e57308ea0c3679
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 15:59:48 2016 -0700

    Revert "go/printer: support for printing alias declarations"
    
    This reverts commit 59c63c711c73f3872c3047c2e80debba5ff1b802.
    
    Reason: Decision to back out current alias implementation.
    
    For #16339.
    
    Change-Id: Idd135fe84b7ce4814cb3632f717736fc6985634c
    Reviewed-on: https://go-review.googlesource.com/32822
    Reviewed-by: Chris Manghane <cmang@golang.org>

commit 26e43779f119683e8571ec109a7bf502ebe95d9a
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 17:01:34 2016 -0700

    Revert "cmd/vet: teach vet about ast.AliasSpec"
    
    This reverts commit aa8c8e770e6db895405b66d38867c2368d94024a.
    
    Reason: Decision to back out current alias implementation.
    
    For #16339.
    
    Change-Id: I4db9a8d6b3625c794be9d2f1ff0e9c047f383d28
    Reviewed-on: https://go-review.googlesource.com/32827
    Run-TryBot: Robert Griesemer <gri@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Chris Manghane <cmang@golang.org>

commit 8e970536dfe0b8ce74bfd0e83ae608c4a012d3c6
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 15:48:39 2016 -0700

    cmd/compile: revert user-visible changes related to aliases
    
    Reason: Decision to back out current alias implementation.
    
    Leaving import/export related code in place for now.
    
    For #16339.
    
    TBR=mdempsky
    
    Change-Id: Ib0897cab2c1c3dc8a541f2efb9893271b0b0efe2
    Reviewed-on: https://go-review.googlesource.com/32757
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit dd1e7b3be0f64438e58f956bfb989608c7fa61bc
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 2 14:45:27 2016 -0700

    crypto/x509: update __MAC_OS_X_VERSION_MAX_ALLOWED on Mac
    
    Reportedly, -mmacosx-version-min=10.6 -D__MAC_OS_X_VERSION_MAX_ALLOWED=1060
    is problematic.
    
    It means min 10.6 and max 10.6, thus exactly 10.6. But we only support
    10.8+.
    
    It never caused us problems, because we build on Macs, but apparently
    if you cross-compile from Linux with some Mac compiler SDK thing, then
    things break?
    
    This was added in https://golang.org/cl/5700083 for #3131, and the
    intent at the time was to pin to exactly 10.6. So it wasn't a mistake,
    but it is definitely outdated.
    
    Given that we now support 10.8 as the min, update it to 1080.
    
    Fixes #17732
    
    Change-Id: I6cc8ab6ac62b8638a5025952b830f23e8822b2a6
    Reviewed-on: https://go-review.googlesource.com/32580
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 039e60ce4e0763f5c67e11227858a4d508df1299
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 15:01:09 2016 -0700

    go/types: revert user-visible changes related to aliases
    
    Reason: Decision to back out current alias implementation.
    For #16339 (comment).
    
    Change-Id: Ie04f24e529db2d29c5dd2e36413f5f37f628df39
    Reviewed-on: https://go-review.googlesource.com/32819
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 2c6949ec89817caecbb441422fe1d6729ee16462
Author: Alan Donovan <adonovan@google.com>
Date:   Fri Nov 4 17:11:51 2016 -0400

    go/types: avoid redundant call to recordUse for anonymous fields
    
    Anonymous fields are type expressions, and Checker.typexpr already
    correctly records uses within them.  There's no need for a second
    call, and the second call caused a bug when we implemented aliases.
    
    Change-Id: I1bf2429cd4948d68b085e75dfb1bdc03ad8caffd
    Reviewed-on: https://go-review.googlesource.com/32837
    Reviewed-by: Robert Griesemer <gri@golang.org>

commit cfd89164bb6af2d1a660b75ded8c0801372924e2
Author: Michael Munday <munday@ca.ibm.com>
Date:   Fri Nov 4 16:30:12 2016 -0400

    all: make copyright headers consistent with one space after period
    
    Continuation of CL 20111.
    
    Change-Id: Ie2f62237e6ec316989c021de9b267cc9d6ee6676
    Reviewed-on: https://go-review.googlesource.com/32830
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 7f033933ce8488ff3c7510ec35680510f9ccc83b
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:40 2016 +0200

    cmd/compile/internal/gc: add support for GOARCH=mips{,le}
    
    Change-Id: Ida4cd647525abce3441bfcb9fdee059344fe717f
    Reviewed-on: https://go-review.googlesource.com/31477
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 87f4e36ce7d7dffbf1f2a869f3014321f6cfff3c
Author: Robert Griesemer <gri@golang.org>
Date:   Fri Nov 4 12:38:53 2016 -0700

    Revert "spec: add new language for alias declarations"
    
    This reverts commit aff37662d1f70f2bf9e47b4f962e85521e7c18d1.
    
    Reason: Decision to back out current alias implementation.
    https://github.com/golang/go/issues/16339#issuecomment-258527920
    
    Fixes #16339.
    Fixes #17746.
    Fixes #17784.
    
    Change-Id: I5737b830d7f6fb79cf36f26403b4ad8533ba1dfe
    Reviewed-on: https://go-review.googlesource.com/32813
    Reviewed-by: Russ Cox <rsc@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 585a0b03b243e1847b3028bc0f2e77bb3688adba
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:45 2016 +0200

    cmd/link/internal/ld: add support for GOARCH=mips{,le}
    
    Change-Id: Ida214ccc5858969ea60abb0787f4d98bab4336d6
    Reviewed-on: https://go-review.googlesource.com/31480
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 5bfafff03f42df38c3088c266ca23d57c9fc3342
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:51:11 2016 +0200

    cmd/vet: add support for GOARCH=mips{,le}
    
    Change-Id: Ie7b40cc67e9901e252a4a48225bbd745a66d2673
    Reviewed-on: https://go-review.googlesource.com/31511
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 23416315060bf7601e5779c3a6a2529d4d604584
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 4 05:28:01 2016 +0000

    all: sprinkle t.Parallel on some slow tests
    
    I used the slowtests.go tool as described in
    https://golang.org/cl/32684 on packages that stood out.
    
    go test -short std drops from ~56 to ~52 seconds.
    
    This isn't a huge win, but it was mostly an exercise.
    
    Updates #17751
    
    Change-Id: I9f3402e36a038d71e662d06ce2c1d52f6c4b674d
    Reviewed-on: https://go-review.googlesource.com/32751
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3f69909851813006216f49083932badee396c4f1
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:51:12 2016 +0200

    test: disable unsupported test for GOARCH=mips{,le}
    
    External linking on mips/mipsle is not supported yet (issue #17792).
    
    Change-Id: Ic25f4f8fe9e0ec35c72ca9f85c053b398df4952c
    Reviewed-on: https://go-review.googlesource.com/31512
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 04d71354eb54b62e9fc520a2fbff8dd76a5782d0
Author: Ian Lance Taylor <iant@golang.org>
Date:   Fri Nov 4 06:43:48 2016 -0700

    runtime: remove useless assignment in test code
    
    Change-Id: I5fecdf52e9e3035ea8feb5768985ed5200dbd6af
    Reviewed-on: https://go-review.googlesource.com/32752
    Run-TryBot: Ian Lance Taylor <iant@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Martin Möhrmann <martisch@uos.de>

commit d62b31386338364b3f319337ca0a1e511c6f23f7
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Thu Nov 3 21:17:45 2016 -0700

    net/http: move extra redirect logic for 307/308 into redirectBehavior
    
    Follow up of CL https://go-review.googlesource.com/32595.
    
    Change-Id: I2b3ff7e6b2c764bb6bc5e9aa692d0aed79eb5626
    Reviewed-on: https://go-review.googlesource.com/32750
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit cd670a61c193e6717400c6cc130484d6ad97c96b
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 4 03:23:37 2016 +0000

    net/http: speed up tests, use t.Parallel when it's safe
    
    Before: 8.9 seconds for go test -short
     After: 2.8 seconds
    
    There are still 250 tests without t.Parallel, but I got the important
    onces using a script:
    
        $ go test -short -v 2>&1 | go run ~/slowtests.go
    
    Where slowtests.go is https://play.golang.org/p/9mh5Wg1nLN
    
    The remaining 250 (the output lines from slowtests.go) all have a
    reported duration of 0ms, except one 50ms test which has to be serial.
    
    Where tests can't be parallel, I left a comment at the top saying why,
    so people don't add t.Parallel later and get surprised at failures.
    
    Updates #17751
    
    Change-Id: Icbe32cbe2b996e23c89f1af6339287fa22af5115
    Reviewed-on: https://go-review.googlesource.com/32684
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 9f5859759eccebb5d0e9d3ef4de1db41d66b41ac
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Thu Nov 3 19:20:27 2016 -0700

    cmd/go: fix minor typo in 'go bug'
    
    Change-Id: I6bb594576e174cb0df8d25d11b84f5a4752ebfd6
    Reviewed-on: https://go-review.googlesource.com/32683
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b5c0470c8d819e0f8a3accbb5a614d47b8ce0c7c
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Fri Nov 4 03:33:44 2016 +0000

    net/http: fix type name in comment
    
    Change-Id: Ia03f993287d2929f35b4c92d00fe25c7243bd8b3
    Reviewed-on: https://go-review.googlesource.com/32685
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>

commit 3440c7bc4c238e1d75d728536ca8f5efe883dbe6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 3 20:53:05 2016 +0000

    net/http: tweak the new Client 307/308 redirect behavior a bit
    
    This CL tweaks the new (unreleased) 307/308 support added in
    https://golang.org/cl/29852 for #10767.
    
    Change 1: if a 307/308 response doesn't have a Location header in its
    response (as observed in the wild in #17773), just do what we used to
    do in Go 1.7 and earlier, and don't try to follow that redirect.
    
    Change 2: don't follow a 307/308 if we sent a body on the first
    request and the caller's Request.GetBody func is nil so we can't
    "rewind" the body to send it again.
    
    Updates #17773 (will be fixed more elsewhere)
    
    Change-Id: I183570f7346917828a4b6f7f1773094122a30406
    Reviewed-on: https://go-review.googlesource.com/32595
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 6e269256264599a49739ad1145bac3845e94ead4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Nov 2 22:56:10 2016 -0700

    cmd/objdump: speed up tests
    
    Rebuild cmd/objdump once instead of twice.
    Speeds up standalone 'go test cmd/objdump' on my
    machine from ~1.4s to ~1s.
    
    Updates #17751
    
    Change-Id: I15fd79cf18c310f892bc28a9e9ca47ee010c989a
    Reviewed-on: https://go-review.googlesource.com/32673
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 3797446150ecc7adbc4e6f9a6315214264ac11f9
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Thu Nov 3 15:58:55 2016 -0700

    cmd/compile: prevent Noalg from breaking user types
    
    Use a separate symbol for reflect metadata for types with Noalg set.
    
    Fixes #17752.
    
    Change-Id: Icb6cade7e3004fc4108f67df61105dc4085cd7e2
    Reviewed-on: https://go-review.googlesource.com/32679
    Reviewed-by: David Crawshaw <crawshaw@golang.org>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit fa770016484165d65200243811f177f1332a9e94
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:25 2016 +0200

    runtime/internal/sys: add arch defs for GOARCH=mips{,le}
    
    Change-Id: I6288f1fca1ae4c64b3907af700811ee842053020
    Reviewed-on: https://go-review.googlesource.com/31472
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 66fdf433b09c72e29b7e5bb5238ec9b3cf745479
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:29 2016 +0200

    cmd/internal/sys: add support for GOARCH=mips{,le}
    
    Change-Id: I8c6b8839c68818430510702719dca15b8d748fb8
    Reviewed-on: https://go-review.googlesource.com/31473
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit c408266e2861513eac79d999f7553dc93992a965
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:59 2016 +0200

    internal/syscall/unix: add randomTrap const for GOARCH=mips{,le}
    
    Change-Id: I76c62a7b79ea526f59f281e933e4fd431539d2da
    Reviewed-on: https://go-review.googlesource.com/31486
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit d3a3b74aa1eec3417a754d798a4cad487949fa77
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Wed Aug 31 16:58:56 2016 +0200

    runtime: 8-byte align the heap_live field for atomic access
    
    mheap_.heap_live is an atomically accessed uint64. It is currently not 8-byte
    aligned on 32-bit platforms, which has been okay because it's only accessed via
    Xadd64, which doesn't require alignment on 386 or ARM32. However, Xadd64 on
    MIPS32 does require 8-byte alignment.
    
    Add a padding field to force 8-byte alignment of heap_live and prevent an
    alignment check crash on MIPS32.
    
    Change-Id: I7eddf7883aec7a0a7e0525af5d58ed4338a401d0
    Reviewed-on: https://go-review.googlesource.com/31635
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 58e3bf11fb8af887a97a7329edfffb1126ba5fee
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:51:03 2016 +0200

    syscall: add support for GOARCH=mips{,le}
    
    Change-Id: I39a46b2a9412981db1780b688a86fec791f68b6f
    Reviewed-on: https://go-review.googlesource.com/31488
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit ad366fdbe40eea7f4fbdc75e8a27304c696793f1
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 3 22:45:06 2016 +0000

    cmd/dist: add mips and mipsle as GOARCH values
    
    Change-Id: I7a51d5d96a7cb87c40ade5be276136c465010bb9
    Reviewed-on: https://go-review.googlesource.com/32596
    Reviewed-by: Cherry Zhang <cherryyz@google.com>
    Reviewed-by: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>

commit 1a07257777ba49231fbf8b81bc6efe34b8422fec
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:57 2016 +0200

    reflect: add support for GOARCH=mips{,le}
    
    Change-Id: I8b0c4bfe1e4c401d5c36a51b937671e6362c73a4
    Reviewed-on: https://go-review.googlesource.com/31485
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 3f2b85a8d14cb0435a84501ac9f25d80d6129348
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:54 2016 +0200

    cmd/cgo: add support for GOARCH=mips{,le}
    
    Change-Id: I47c6867fc653c8388ad32e210a8027baa592eda3
    Reviewed-on: https://go-review.googlesource.com/31483
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit d1e9104fb25a71aa459ee329452545a0ceebdcf9
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:55 2016 +0200

    math, math/big: add support for GOARCH=mips{,le}
    
    Change-Id: I54e100cced5b49674937fb87d1e0f585f962aeb7
    Reviewed-on: https://go-review.googlesource.com/31484
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 9788e3d4d72792da22a12e54f8a86e25a03b58ab
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:51:06 2016 +0200

    sync/atomic: add support for GOARCH=mips{,le}
    
    Change-Id: I10f36710dd95b9bd31b3b82a3c32edcadb90ffa9
    Reviewed-on: https://go-review.googlesource.com/31510
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit b241a06479864b85c2377e52b775b176f686e192
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:51:04 2016 +0200

    runtime/internal/atomic: add GOARCH=mips{,le} support
    
    Change-Id: I99a48f719fd1a8178fc59787084a074e91c89ac6
    Reviewed-on: https://go-review.googlesource.com/31489
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 2c39e50995bb02325b2c17f253b10f5ada0e337f
Author: Vladimir Stefanovic <vladimir.stefanovic@imgtec.com>
Date:   Tue Oct 18 23:50:48 2016 +0200

    debug/elf: add support for GOARCH=mips{,le}
    
    Change-Id: Ia6f8ae7e56a49ad66b60a24c4afb606f3cfe5efd
    Reviewed-on: https://go-review.googlesource.com/31482
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 26827bc2fe4c80dc68b3793631d24975425c9467
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Sat Oct 22 07:25:21 2016 -0700

    testing: add T.Context method
    
    From the doc comment:
    
    Context returns the context for the current test or benchmark.
    The context is cancelled when the test or benchmark finishes.
    A goroutine started during a test or benchmark can wait for the
    context's Done channel to become readable as a signal that the
    test or benchmark is over, so that the goroutine can exit.
    
    Fixes #16221.
    Fixes #17552.
    
    Change-Id: I657df946be2c90048cc74615436c77c7d9d1226c
    Reviewed-on: https://go-review.googlesource.com/31724
    Reviewed-by: Rob Pike <r@golang.org>

commit 606f81eef37e5a232f43a208f6eeaddd82dadf34
Author: Russ Cox <rsc@golang.org>
Date:   Thu Nov 3 15:01:30 2016 -0400

    context: adjust tests to avoid importing "testing" in package context
    
    So that testing can use context in its public API.
    
    For #16221.
    
    Change-Id: I6263fa7266c336c9490f20164ce79336df44a57e
    Reviewed-on: https://go-review.googlesource.com/32648
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2b59b15f6b7e6fa6ac725367acff5aef7c666e36
Author: Emmanuel Odeke <emm.odeke@gmail.com>
Date:   Wed Sep 21 23:11:28 2016 -0700

    encoding/json: example on MarshalJSON, UnmarshalJSON
    
    Updates #16360.
    
    Change-Id: I5bf13d3367e68c5d8435f6ef2161d5a74cc747a7
    Reviewed-on: https://go-review.googlesource.com/29611
    Reviewed-by: Andrew Gerrand <adg@golang.org>
    Run-TryBot: Andrew Gerrand <adg@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 43f954e09858449cc5f3650720e81b7e879ab349
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 21:40:47 2016 -0400

    testing: mark tests and benchmarks failed if a race occurs during execution
    
    Before:
    
    $ go test -race -v -run TestRace
    === RUN   TestRace
    ==================
    WARNING: DATA RACE
    Write at 0x00c420076420 by goroutine 7:
      _/Users/rsc/go/src/cmd/go/testdata/src/testrace.TestRace.func1()
          /Users/rsc/go/src/cmd/go/testdata/src/testrace/race_test.go:10 +0x3b
    
    Previous write at 0x00c420076420 by goroutine 6:
      _/Users/rsc/go/src/cmd/go/testdata/src/testrace.TestRace()
          /Users/rsc/go/src/cmd/go/testdata/src/testrace/race_test.go:13 +0xcc
      testing.tRunner()
          /Users/rsc/go/src/testing/testing.go:656 +0x104
    
    Goroutine 7 (running) created at:
      _/Users/rsc/go/src/cmd/go/testdata/src/testrace.TestRace()
          /Users/rsc/go/src/cmd/go/testdata/src/testrace/race_test.go:12 +0xbb
      testing.tRunner()
          /Users/rsc/go/src/testing/testing.go:656 +0x104
    
    Goroutine 6 (running) created at:
      testing.(*T).Run()
          /Users/rsc/go/src/testing/testing.go:693 +0x536
      testing.runTests.func1()
          /Users/rsc/go/src/testing/testing.go:877 +0xaa
      testing.tRunner()
          /Users/rsc/go/src/testing/testing.go:656 +0x104
      testing.runTests()
          /Users/rsc/go/src/testing/testing.go:883 +0x4ac
      testing.(*M).Run()
          /Users/rsc/go/src/testing/testing.go:818 +0x1c3
      main.main()
          _/Users/rsc/go/src/cmd/go/testdata/src/testrace/_test/_testmain.go:42 +0x20f
    ==================
    --- PASS: TestRace (0.00s)
    PASS
    Found 1 data race(s)
    FAIL    _/Users/rsc/go/src/cmd/go/testdata/src/testrace 1.026s
    $
    
    After:
    
    $ go test -race -v -run TestRace
    === RUN   TestRace
    ==================
    WARNING: DATA RACE
    Write at 0x00c420076420 by goroutine 7:
      _/Users/rsc/go/src/cmd/go/testdata/src/testrace.TestRace.func1()
          /Users/rsc/go/src/cmd/go/testdata/src/testrace/race_test.go:10 +0x3b
    
    Previous write at 0x00c420076420 by goroutine 6:
      _/Users/rsc/go/src/cmd/go/testdata/src/testrace.TestRace()
          /Users/rsc/go/src/cmd/go/testdata/src/testrace/race_test.go:13 +0xcc
      testing.tRunner()
          /Users/rsc/go/src/testing/testing.go:656 +0x104
    
    Goroutine 7 (running) created at:
      _/Users/rsc/go/src/cmd/go/testdata/src/testrace.TestRace()
          /Users/rsc/go/src/cmd/go/testdata/src/testrace/race_test.go:12 +0xbb
      testing.tRunner()
          /Users/rsc/go/src/testing/testing.go:656 +0x104
    
    Goroutine 6 (running) created at:
      testing.(*T).Run()
          /Users/rsc/go/src/testing/testing.go:693 +0x536
      testing.runTests.func1()
          /Users/rsc/go/src/testing/testing.go:877 +0xaa
      testing.tRunner()
          /Users/rsc/go/src/testing/testing.go:656 +0x104
      testing.runTests()
          /Users/rsc/go/src/testing/testing.go:883 +0x4ac
      testing.(*M).Run()
          /Users/rsc/go/src/testing/testing.go:818 +0x1c3
      main.main()
          _/Users/rsc/go/src/cmd/go/testdata/src/testrace/_test/_testmain.go:42 +0x20f
    ==================
    --- FAIL: TestRace (0.00s)
            testing.go:609: race detected during execution of test
    FAIL
    FAIL    _/Users/rsc/go/src/cmd/go/testdata/src/testrace 0.022s
    $
    
    Fixes #15972.
    
    Change-Id: Idb15b8ab81d65637bb535c7e275595ca4a6e450e
    Reviewed-on: https://go-review.googlesource.com/32615
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit f3862742b67a84edf939f41276360ada4e7197a6
Author: Dan Harrington <harringtond@google.com>
Date:   Tue Oct 25 12:51:39 2016 -0700

    net/http: support If-Match in ServeContent
    
    - Added support for If-Match and If-Unmodified-Since.
    - Precondition checks now more strictly follow RFC 7232 section 6, which
    affects precedence when multiple condition headers are present.
    - When serving a 304, Last-Modified header is now removed when no ETag is
    present (as suggested by RFC 7232 section 4.1).
    - If-None-Match supports multiple ETags.
    - ETag comparison now correctly handles weak ETags.
    
    Fixes #17572
    
    Change-Id: I35039dea6811480ccf2889f8ed9c6a39ce34bfff
    Reviewed-on: https://go-review.googlesource.com/32014
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 18f0e88103aaa429e92564312b4ee966dcb77102
Author: Robert Griesemer <gri@golang.org>
Date:   Thu Nov 3 10:47:44 2016 -0700

    go/internal/gcimporter: backport changes from x/tools/go/gcimporter15
    
    See https://go-review.googlesource.com/32581.
    
    This makes x/tools/go/gcimporter15/bimport.go a close copy again
    and the importer more robust.
    
    Change-Id: If96ad6acd611878b7dfa6a13d005d847ece82ab6
    Reviewed-on: https://go-review.googlesource.com/32647
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Robert Griesemer <gri@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 322436b78c2b0a8dd346f460df6fc00f49adba8e
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Thu Nov 3 22:25:42 2016 +0530

    cmd/compile: don't panic if syntax.ReadFile returns an error
    
    Fixes #17772
    
    Change-Id: I0f2094400c454828aa57a8d172dadeac4ddb6d35
    Reviewed-on: https://go-review.googlesource.com/32691
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit 4141054d19f6d5bcdc448be00a6ce4a4c864d488
Author: David Chase <drchase@google.com>
Date:   Thu Nov 3 11:50:14 2016 -0400

    runtime/cgo: correct type declaration for Windows
    
    Newer versions of gcc notice a type mismatch and complain.
    Fix code to match documented signature in MSDN.
    Trybots say this still compiles with the older (5.1) version
    of gcc.
    
    Fixes #17771.
    
    Change-Id: Ib3fe6f71b40751e1146249e31232da5ac69b9e00
    Reviewed-on: https://go-review.googlesource.com/32646
    Run-TryBot: David Chase <drchase@google.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit b89135777be7c2f123dbf2eea3c92a5402e4c63c
Author: Russ Cox <rsc@golang.org>
Date:   Thu Nov 3 09:39:32 2016 -0400

    crypto/x509: expose UnknownAuthorityError.Cert
    
    This matches exposing CertificateInvalidError.Cert.
    and (exposing but not the spelling of) HostnameError.Certificate.
    
    Fixes #13519.
    
    Change-Id: Ifae9a09e063d642c09de3cdee8a728ff06d3a5df
    Reviewed-on: https://go-review.googlesource.com/32644
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit b2290229c23fed8e4a63241568e436b1daa1196f
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 19:41:01 2016 -0400

    cmd/cgo: add #line directives to avoid printing bogus references to Go source files
    
    A bit contrived to come up with an example, but it happened in #15836, somehow.
    
            $ cat /tmp/x.go
            package main
    
            /*
            #include <stddef.h>
    
            int foo(void);
    
            int foo(void) {
                    return 2;
            }
    
            #define int asdf
            */
            import "C"
    
            func main() {
                    println(C.foo())
            }
    
            $ go run /tmp/x.go
            # command-line-arguments
            cgo-builtin-prolog:9:31: error: unknown type name 'asdf'   <<<<<
            _GoString_ GoStringN(char *p, int l);
                                          ^
            /tmp/x.go:12:13: note: expanded from macro 'int'
            #define int asdf
                        ^
            cgo-builtin-prolog:10:28: error: unknown type name 'asdf'  <<<<<
            _GoBytes_ GoBytes(void *p, int n);
                                       ^
            /tmp/x.go:12:13: note: expanded from macro 'int'
            #define int asdf
                        ^
            2 errors generated.
    
    The two marked lines used to refer incorrectly to /tmp/x.go.
    
    Fixes #15836.
    
    Change-Id: I08ef60a53cfd148112fceb651eaf7b75d94a7a8d
    Reviewed-on: https://go-review.googlesource.com/32613
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit a4a4d43028ba20ecec50ef761012d17800553cde
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 20:18:47 2016 -0400

    cmd/cover: change covered block for switch/select case to exclude expression
    
    Consider a switch like
    
            switch x {
            case foo:
                    f()
                    g()
            }
    
    Before, the coverage annotation for the block calling f and g included
    in its position span the text for 'case foo:'. This looks nice in the coverage
    report, but it breaks the invariant that coverage blocks are disjoint if
    you have a more complex expression like:
    
            switch x {
            case func() int { return foo }():
                    f()
                    g()
            }
    
    Then the coverage analysis wants to annotate the func literal body,
    which overlaps with the case body, because the case body is considered
    to begin at the case token.
    
    Change the annotation for a case body to start just after the colon of
    the case clause, avoiding any potential conflict with complex case
    expressions. Could have started at the colon instead, but it seemed
    less weird to start just after it.
    
    Fixes #16540.
    
    Change-Id: I1fec4bc2a53c7092e649dc0d4be1680a697cb79b
    Reviewed-on: https://go-review.googlesource.com/32612
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit 95e809f0e525fca000ec634e1e8cc10d6cebdf30
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 23:12:10 2016 -0400

    cmd/go: clear GIT_ALLOW_PROTOCOL during tests
    
    Clear it before any tests begin.
    Clear it again after TestIsSecureGitAllowProtocol sets it.
    
    Fixes #17700.
    
    Change-Id: I6ea50d37f8222d8c7c9fee0b1eac3bbdfb5d133e
    Reviewed-on: https://go-review.googlesource.com/32640
    Reviewed-by: Quentin Smith <quentin@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>

commit a35decf5a7a722acb32642b6a096271d10545376
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 23:18:54 2016 -0400

    cmd/go: fix TestIssue11457
    
    The goal of the test is to provoke a custom import path check error,
    which will contain the current repo path, to see that it says ssh:// in it.
    
    But the fix to #16471 made the test no longer provoke that error.
    Provoke the error by checking out from rsc.io instead of github.com/rsc.
    
    Fixes #17701.
    
    Change-Id: I750ffda2ff59c2be8e111d26160997214a73fd9a
    Reviewed-on: https://go-review.googlesource.com/32641
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5f9fb1ce0293c3d5bc2b0fd81ad7e0aaafd414f0
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 23:22:18 2016 -0400

    cmd/go: remove .o and .a files in builder.collect
    
    This matches the removal of .so files and makes the processing
    of '-L/path/to -lfoo' and plain '/path/to/foo.a' match.
    
    Fixes #16463.
    
    Change-Id: I1464c5390d7eb6a3a33b4b2c951f87ef392ec94a
    Reviewed-on: https://go-review.googlesource.com/32642
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 23e6e1124522267655b05a5f47f62fc99f0e56c3
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 26 12:18:00 2016 -0400

    encoding/asn1: document that default:x tag only has meaning with optional tag
    
    Fixes #16712.
    
    Change-Id: Ib216059c6c0c952162c19e080dcf3799f0652a8d
    Reviewed-on: https://go-review.googlesource.com/32171
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Quentin Smith <quentin@golang.org>

commit ced137fad48ad18518cdd4ed3a77e75d5e58c78c
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 16:33:33 2016 -0400

    misc/cgo/testsanitizers: skip tests when vm.overcommit_memory=2
    
    Fixes #17689.
    
    Change-Id: I45a14e6bf4b2647431105f3e0b63b7076b6655d2
    Reviewed-on: https://go-review.googlesource.com/32635
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3345802e8ac39e4ea6b2772ace15983f2c3be66b
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 17:19:42 2016 -0400

    time: document that only Jan and January, Mon and Monday are recognized
    
    Not "J", not "JAN", not "jan", etc.
    
    Fixes #17523.
    
    Change-Id: I16b5da97e73d88c6680c36401d30f8a195061527
    Reviewed-on: https://go-review.googlesource.com/32636
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: Rob Pike <r@golang.org>

commit 1c08c728dceeb63f6df915831bd0584be1517e00
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 21:15:41 2016 -0400

    cmd/go: fix coverage in xtest of cgo package
    
    Cover-annotated cgo-rebuilt package for xtest was
    not linked into package graph and so not being rebuilt.
    Link into package graph.
    
    Fixes #13625.
    
    Change-Id: I685f7276f92bbc85fbc4b389111c83d9fe517637
    Reviewed-on: https://go-review.googlesource.com/32614
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 3ce46e3e9c25d6234f49ab483300e538c65089d0
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 21:52:27 2016 -0400

    cmd/go: fix TestGoGetDashTIssue8181
    
    The test case was importing golang.org/x/build/cmd/cl,
    which is a package main and cannot be imported.
    The test case (stored in a separate repo) has been changed
    to import golang.org/x/build/gerrit. Update the test accordingly.
    
    Fixes #17702.
    
    Change-Id: I80e150092111b5a04bb00c992b32edb271edb086
    Reviewed-on: https://go-review.googlesource.com/32616
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 9efcdd4e694789646e7de79dfbd8c7636803785a
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 23:01:20 2016 -0400

    cmd/go: use new HTTP-only server insecure.go-get-issue-15410.appspot.com instead of wh3rd.net
    
    Fixes #15410.
    
    Change-Id: I9964d0162a3cae690afeb889b1822cf79c80b89a
    Reviewed-on: https://go-review.googlesource.com/32639
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Andrew Gerrand <adg@golang.org>

commit f4f6b647500438eaff21e36f48ea4399727c41bd
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Nov 2 20:34:29 2016 -0700

    cmd/compile: update/remove outdated comments
    
    Change-Id: I5a74be1593dca8ba5e0829f0bae35dc9ce702671
    Reviewed-on: https://go-review.googlesource.com/32672
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 21c114e930fc1b118c2b8b8716d7b1b3b38027e0
Author: Quentin Smith <quentin@golang.org>
Date:   Wed Nov 2 16:18:22 2016 -0400

    runtime/testdata/testprog: increase GCFairness2 timeout to 1s
    
    OpenBSD's scheduler causes preemption to take 20+ms, so 30ms is not
    enough time for 3 goroutines to run. This change continues to sleep for
    30ms, but if it finds that the 3 goroutines have not run, it sleeps for
    an additional 1s before declaring failure.
    
    Updates #17712
    
    Change-Id: I3e886e40d05192b7cb71b4f242af195836ef62a8
    Reviewed-on: https://go-review.googlesource.com/32634
    Reviewed-by: Rick Hudson <rlh@golang.org>
    Run-TryBot: Quentin Smith <quentin@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit e83911d6fc05c0ccb11bd95cf60694f8aec1698a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Nov 2 18:16:51 2016 -0700

    cmd/vet/all: add s390x support
    
    Some of these whitelist entries could be
    eliminated, e.g. by the addition of Go
    declarations, but this is a start.
    
    Change-Id: I2fb3234cf05ebc6b161aacac2d4c15d810d50527
    Reviewed-on: https://go-review.googlesource.com/32671
    Reviewed-by: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 62dafbb4c6cd51332b2e0519eb4267c7a03da446
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Nov 2 19:02:43 2016 -0700

    runtime: fix s390x asm frame sizes
    
    Found by vet.
    
    Change-Id: I1d78454facdd3522509ecfe7c73b21c4602ced8a
    Reviewed-on: https://go-review.googlesource.com/32670
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit 7f4c3e87106a367babb6662b2d1e3b485e0d07af
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Nov 2 19:02:15 2016 -0700

    all: update s390x assembly to use vet-friendly names
    
    No functional changes.
    
    Change-Id: Ibf592c04be506a76577d48574e84ab20c3238b49
    Reviewed-on: https://go-review.googlesource.com/32589
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Michael Munday <munday@ca.ibm.com>

commit a3aaa189a484e18cad30391b7b5a70d04bc3cbe3
Author: Mohit Agarwal <mohit@sdf.org>
Date:   Thu Nov 3 19:53:27 2016 +0530

    cmd/go: run mkalldocs.sh
    
    Follow-up to CL 32114
    
    Change-Id: I75247ed9c1c0a0e8a278eb75a60d4c5bee355409
    Reviewed-on: https://go-review.googlesource.com/32690
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit fc2e282c04a3c36761254eb662bb20bad4eb4a35
Author: Russ Cox <rsc@golang.org>
Date:   Wed Nov 2 23:31:08 2016 -0400

    cmd/go: add version of GOROOT to go bug details
    
    Fixes #15877.
    
    Change-Id: Ia1e327c0cea3be43e5f8ba637c97c223cee4bb5a
    Reviewed-on: https://go-review.googlesource.com/32643
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 89ccd5795c5e41ef478eab3d4a292a836a7fe3ad
Author: Russ Cox <rsc@golang.org>
Date:   Thu Nov 3 09:13:16 2016 -0400

    cmd/pprof: use correct default handler for Go programs
    
    The original go tool pprof (written in Perl) got this right.
    The Go rewrite never has, but should.
    
    Change-Id: Ie1fc571214c61b1b5654a0bc90e15eb889adf059
    Reviewed-on: https://go-review.googlesource.com/32617
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 8eb9fdaa0123fc98cb70f58801eb74c8a6f92817
Author: David Crawshaw <crawshaw@golang.org>
Date:   Sat Sep 24 08:39:31 2016 +1000

    cmd/compile: write type symbols referenced in ptabs
    
    The exported symbol for a plugin can be the only reference to a
    type in a program. In particular, "var F func()" will have
    the type *func(), which is uncommon.
    
    Fixes #17140
    
    Change-Id: Ide2104edbf087565f5377374057ae54e0c00c57e
    Reviewed-on: https://go-review.googlesource.com/29692
    Run-TryBot: David Crawshaw <crawshaw@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 5ac3e7d6a7717bf2f722803e1852ba991af4e724
Author: David du Colombier <0intro@gmail.com>
Date:   Thu Nov 3 11:50:45 2016 +0100

    net: disallow dialing and listening on port > 65535 on Plan 9
    
    Since Dial(":80") was implemented on Plan 9 (CL 32593),
    TestProtocolDialError is failing because dialing a port
    superior to 65535 is supported on Plan 9.
    
    This change disallows dialing and listening on ports
    superior to 65535.
    
    Fixes #17761.
    
    Change-Id: I95e8a163eeacf1ccd8ece7b650f16a0531c59709
    Reviewed-on: https://go-review.googlesource.com/32594
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 97b49f660cbb37ff6789dbb92b1a6ded394d3d7d
Author: David du Colombier <0intro@gmail.com>
Date:   Thu Nov 3 14:57:05 2016 +0100

    net: fix Dial(":80") on Plan 9
    
    CL 32101 fixed Dial(":80") on Windows and added TestDialLocal,
    which was failing on Plan 9, because it wasn't implemented
    on Plan 9.
    
    This change implements Dial(":80") by connecting to 127.0.0.1
    or ::1 (depending on network), so it works as documented.
    
    Fixes #17760.
    
    Change-Id: If0ff769299e09bebce11fc3708639c1d8c96c280
    Reviewed-on: https://go-review.googlesource.com/32593
    Reviewed-by: Russ Cox <rsc@golang.org>

commit aa8c8e770e6db895405b66d38867c2368d94024a
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Nov 2 18:25:57 2016 -0700

    cmd/vet: teach vet about ast.AliasSpec
    
    Fixes #17755
    
    Change-Id: I1ad1edc382b1312d992963054eb82648cb5112d2
    Reviewed-on: https://go-review.googlesource.com/32588
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit 10f757486e94f60a5e0af180dcd61c9eef7534c6
Author: Ilya Tocar <ilya.tocar@intel.com>
Date:   Thu Oct 27 16:58:45 2016 +0300

    cmd/compile/internal/ssa: generate bswap on AMD64
    
    Generate bswap+load/store for reading/writing big endian data.
    Helps encoding/binary.
    
    name                    old time/op    new time/op    delta
    ReadSlice1000Int32s-8     5.06µs ± 8%    4.58µs ± 8%   -9.50%        (p=0.000 n=10+10)
    ReadStruct-8              1.07µs ± 0%    1.05µs ± 0%   -1.51%         (p=0.000 n=9+10)
    ReadInts-8                 367ns ± 0%     363ns ± 0%   -1.15%          (p=0.000 n=8+9)
    WriteInts-8                475ns ± 1%     469ns ± 0%   -1.45%        (p=0.000 n=10+10)
    WriteSlice1000Int32s-8    5.03µs ± 3%    4.50µs ± 3%  -10.45%          (p=0.000 n=9+9)
    PutUvarint32-8            17.2ns ± 0%    17.2ns ± 0%     ~     (all samples are equal)
    PutUvarint64-8            46.7ns ± 0%    46.7ns ± 0%     ~           (p=0.509 n=10+10)
    
    name                    old speed      new speed      delta
    ReadSlice1000Int32s-8    791MB/s ± 8%   875MB/s ± 8%  +10.53%        (p=0.000 n=10+10)
    ReadStruct-8            70.0MB/s ± 0%  71.1MB/s ± 0%   +1.54%         (p=0.000 n=9+10)
    ReadInts-8              81.6MB/s ± 0%  82.6MB/s ± 0%   +1.21%          (p=0.000 n=9+9)
    WriteInts-8             63.0MB/s ± 1%  63.9MB/s ± 0%   +1.45%        (p=0.000 n=10+10)
    WriteSlice1000Int32s-8   796MB/s ± 4%   888MB/s ± 3%  +11.65%          (p=0.000 n=9+9)
    PutUvarint32-8           233MB/s ± 0%   233MB/s ± 0%     ~           (p=0.089 n=10+10)
    PutUvarint64-8           171MB/s ± 0%   171MB/s ± 0%     ~            (p=0.137 n=10+9)
    
    Change-Id: Ia2dbdef92198eaa7e2af5443a8ed586d4b401ffb
    Reviewed-on: https://go-review.googlesource.com/32222
    Run-TryBot: Ilya Tocar <ilya.tocar@intel.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit ca5cea9dca0bc635b9a23cfe65226f3e1a423342
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Nov 2 17:55:53 2016 -0700

    cmd/compile: add OMOD to list of ops that might panic
    
    Follow-up to CL 32551.
    
    Change-Id: If68f9581a7f13e04796aaff2007c09f8ea2c3611
    Reviewed-on: https://go-review.googlesource.com/32586
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Keith Randall <khr@golang.org>

commit 154d013155978ba9c714d931f0e91c9c964dd82c
Author: Hiroshi Ioka <hirochachacha@gmail.com>
Date:   Fri Oct 21 09:00:07 2016 +0900

    encoding/asn1: return error for unexported fields in Marshal, Unmarshal
    
    The old code cannot handle unexported fields, it panics.
    The new code returns error instead.
    
    Fixes #17462
    
    Change-Id: I927fc46b21d60e86cb52e84c65f2122f9159b21d
    Reviewed-on: https://go-review.googlesource.com/31540
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Russ Cox <rsc@golang.org>

commit be7c50a7101a6e3cbd39648814c9f1030980c0f6
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 3 01:11:49 2016 +0000

    net/http: deflake TestServerSetKeepAlivesEnabledClosesConns
    
    Fixes #17754
    Updates #9478 (details in here)
    
    Change-Id: Iae2c1ca05a18ed266b53b2594c22fc57fab33c5e
    Reviewed-on: https://go-review.googlesource.com/32587
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 66504485ebd8d7a75b68c025cc15835b016a1ae4
Author: Josh Bleecher Snyder <josharian@gmail.com>
Date:   Wed Nov 2 17:55:29 2016 -0700

    cmd/compile/internal/gc: make tests run faster
    
    TestAssembly takes 20s on my machine,
    which is too slow for normal operation.
    Marking as -short has its dangers (#17472),
    but hopefully we'll soon have a builder for that.
    
    All the SSA tests are hermetic and not time sensitive
    and can thus be run in parallel.
    Reduces the cmd/compile/internal/gc test time during
    all.bash on my laptop from 42s to 7s.
    
    Updates #17751
    
    Change-Id: Idd876421db23b9fa3475e8a9b3355a5dc92a5a29
    Reviewed-on: https://go-review.googlesource.com/32585
    Run-TryBot: Josh Bleecher Snyder <josharian@gmail.com>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 827f2accc1a25b77003303d3c35db5fc054bf8f1
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Thu Nov 3 00:44:16 2016 +0000

    net/http: update bundled http2 to fix test flake
    
    Updates http2 to x/net/http2 git rev 569280fa for:
    
       http2: fix over-aggressive ignoring of frames while in "go away" mode
       https://golang.org/cl/32583
    
    Fixes #17733
    
    Change-Id: I4008d2e14ce89782ce7e18b441b1181f98623b9d
    Reviewed-on: https://go-review.googlesource.com/32584
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    Reviewed-by: David Crawshaw <crawshaw@golang.org>

commit 627f4d85ba4fa71e5af11ee047ba42196cea1f2c
Author: Robert Griesemer <gri@golang.org>
Date:   Wed Nov 2 14:54:46 2016 -0700

    go/types: set up correct type with NewAlias
    
    Change-Id: I4b035b3539c98e5b1442d1009d457cbc199b42ee
    Reviewed-on: https://go-review.googlesource.com/32637
    Reviewed-by: Matthew Dempsky <mdempsky@google.com>

commit 5513f855f6e627606f0b516c4a16ae32d82d99fd
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Nov 2 13:43:57 2016 -0400

    cmd/vet: add test case for ppc64{,le}
    
    Adapted from the mips64 test case.
    
    Fixes #17745.
    
    Change-Id: I46f0900028adb936dcab2cdc701ea11d0a3cb95e
    Reviewed-on: https://go-review.googlesource.com/32611
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: David Chase <drchase@google.com>
    Reviewed-by: Rob Pike <r@golang.org>

commit cf28e5cc9d51928ae05df0b193edc7e39a28c413
Author: Keith Randall <khr@golang.org>
Date:   Tue Nov 1 15:28:10 2016 -0700

    cmd/compile: compute faulting args before writing args to stack
    
    when compiling f(a, b, c), we do something like:
      *(SP+0) = eval(a)
      *(SP+8) = eval(b)
      *(SP+16) = eval(c)
      call f
    
    If one of those evaluations is later determined to unconditionally panic
    (say eval(b) in this example), then the call is deadcode eliminated. But
    any previous argument write (*(SP+0)=... here) is still around. Becuase
    we only compute the size of the outarg area for calls which are still
    around at the end of optimization, the space needed for *(SP+0)=v is not
    accounted for and thus the outarg area may be too small.
    
    The fix is to make sure that we evaluate any potentially panicing
    operation before we write any of the args to the stack. It turns out
    that fix is pretty easy, as we already have such a mechanism available
    for function args. We just need to extend it to possibly panicing args
    as well.
    
    The resulting code (if b and c can panic, but a can't) is:
      tmpb = eval(b)
      *(SP+16) = eval(c)
      *(SP+0) = eval(a)
      *(SP+8) = tmpb
      call f
    
    This change tickled a bug in how we find the arguments for intrinsic
    calls, so that latent bug is fixed up as well.
    
    Update #16760.
    
    Change-Id: I0bf5edf370220f82bc036cf2085ecc24f356d166
    Reviewed-on: https://go-review.googlesource.com/32551
    Reviewed-by: Cherry Zhang <cherryyz@google.com>

commit 688995d1e91fcc36ac33bf34d6303e935d1b0a70
Author: Keith Randall <khr@golang.org>
Date:   Fri Oct 28 11:37:45 2016 -0700

    cmd/compile: do more type conversion inline
    
    The code to do the conversion is smaller than the
    call to the runtime.
    The 1-result asserts need to call panic if they fail, but that
    code is out of line.
    
    The only conversions left in the runtime are those which
    might allocate and those which might need to generate an itab.
    
    Given the following types:
      type E interface{}
      type I interface { foo() }
      type I2 iterface { foo(); bar() }
      type Big [10]int
      func (b Big) foo() { ... }
    
    This CL inlines the following conversions:
    
    was assertE2T
      var e E = ...
      b := i.(Big)
    was assertE2T2
      var e E = ...
      b, ok := i.(Big)
    was assertI2T
      var i I = ...
      b := i.(Big)
    was assertI2T2
      var i I = ...
      b, ok := i.(Big)
    was assertI2E
      var i I = ...
      e := i.(E)
    was assertI2E2
      var i I = ...
      e, ok := i.(E)
    
    These are the remaining runtime calls:
    
    convT2E:
      var b Big = ...
      var e E = b
    convT2I:
      var b Big = ...
      var i I = b
    convI2I:
      var i2 I2 = ...
      var i I = i2
    assertE2I:
      var e E = ...
      i := e.(I)
    assertE2I2:
      var e E = ...
      i, ok := e.(I)
    assertI2I:
      var i I = ...
      i2 := i.(I2)
    assertI2I2:
      var i I = ...
      i2, ok := i.(I2)
    
    Fixes #17405
    Fixes #8422
    
    Change-Id: Ida2367bf8ce3cd2c6bb599a1814f1d275afabe21
    Reviewed-on: https://go-review.googlesource.com/32313
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: David Chase <drchase@google.com>

commit 761443edd56832cc1b62f9193f157ca822dfa09e
Author: Keith Randall <khr@golang.org>
Date:   Mon Oct 31 21:18:00 2016 -0700

    cmd/compile: On a runtime.KeepAlive call, keep whole variable alive
    
    We generate an OpKeepAlive for the idata portion of the interface
    for a runtime.KeepAlive call.  But given such an op, we need to keep
    the entire containing variable alive, not just the range that was
    passed to the OpKeepAlive operation.
    
    Fixes #17710
    
    Change-Id: I90de66ec8065e22fb09bcf9722999ddda289ae6e
    Reviewed-on: https://go-review.googlesource.com/32477
    Run-TryBot: Keith Randall <khr@golang.org>
    Reviewed-by: Joe Tsai <thebrokentoaster@gmail.com>

commit 09bb6434f9a8681b81a667e1ff186c61bbe0a50f
Author: Brad Fitzpatrick <bradfitz@golang.org>
Date:   Wed Nov 2 19:48:35 2016 +0000

    net/http: update bundled http2
    
    Update bundled x/net/http2 to x/net git rev 6c4ac8bd for:
    
       http2: fix Transport race sending RST_STREAM while reading DATA on cancels
       https://golang.org/cl/32571
    
       http2: remove h2-14 ALPN proto
       https://golang.org/cl/32576
    
    Fixes #16974
    
    Change-Id: I6ff8493a13d2641499fedf33e8005004735352ff
    Reviewed-on: https://go-review.googlesource.com/32578
    Run-TryBot: Brad Fitzpatrick <bradfitz@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 2d3cd51dbe5b5266b9baf68fb42ea19fedf4d23e
Author: Joe Tsai <joetsai@digital-static.net>
Date:   Thu Oct 27 16:23:53 2016 -0700

    archive/tar: disable prefix field in Writer
    
    The proper fix for the Writer is too involved to be done in time
    for Go 1.8. Instead, we do a localized fix that simply disables the
    prefix encoding logic. While this will prevent some legitimate uses
    of prefix, it will ensure that we don't keep outputting invalid
    GNU format files that have the prefix field populated.
    
    For headers with long filenames that could have used the prefix field,
    they will be promoted to use the PAX format, which ensures that we
    will still be able to encode all headers that we were able to do before.
    
    Updates #12594
    Fixes #17630
    Fixes #9683
    
    Change-Id: Ia97b524ac69865390e2ae8bb0dfb664d40a05add
    Reviewed-on: https://go-review.googlesource.com/32234
    Reviewed-by: Russ Cox <rsc@golang.org>
    Run-TryBot: Joe Tsai <thebrokentoaster@gmail.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit eed2bb71d235e2c8b6d6eae57af83fbff23d420e
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Nov 2 13:11:56 2016 -0400

    cmd/vet: fix go vet on s390x assembly
    
    Test adapted from the mips64 test.
    
    Fixes #15454.
    
    Change-Id: If890c2d18a4a03a08faaa2e674edd7223af60290
    Reviewed-on: https://go-review.googlesource.com/22472
    Run-TryBot: Michael Munday <munday@ca.ibm.com>
    Reviewed-by: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>

commit bcc0247331a77015053d3fa28f458d0c639d6730
Author: Matthew Dempsky <mdempsky@google.com>
Date:   Wed Nov 2 09:47:43 2016 -0700

    cmd/compile: avoid double export of aliased objects
    
    Instead of writing out the original object for each alias, ensure we
    export the original object before any aliases. This allows the aliases
    to simply refer back to the original object by qualified name.
    
    Fixes #17636.
    
    Change-Id: If80fa8c66b8fee8344a00b55d25a8aef22abd859
    Reviewed-on: https://go-review.googlesource.com/32575
    Run-TryBot: Matthew Dempsky <mdempsky@google.com>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Robert Griesemer <gri@golang.org>
    Reviewed-by: Alan Donovan <adonovan@google.com>

commit 1a0b1cca4c26d41fe7508ffdb355de78b4ea2a19
Author: Russ Cox <rsc@golang.org>
Date:   Wed Oct 26 21:07:52 2016 -0400

    net: fix Dial(":80") on Windows
    
    Windows sockets allow bind to 0.0.0.0:80 but not connect to it.
    To make Listen(":80") / Dial(":80") work as documented on Windows,
    connect to 127.0.0.1 or ::1 (depending on network) in place of 0.0.0.0.
    
    Fixes #6290.
    
    Change-Id: Ia27537067276871648546678fbe0f1b8478329fe
    Reviewed-on: https://go-review.googlesource.com/32101
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Mikio Hara <mikioh.mikioh@gmail.com>

commit c56cc9b3b5727647c2afb3d57f5793151558a0a7
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 31 19:27:26 2016 -0400

    testing: introduce testing/internal/testdeps for holding testmain dependencies
    
    Currently, we don't have package testing to import package regexp directly,
    because then regexp can't have internal tests (or at least they become more
    difficult to write), for fear of an import cycle. The solution we've been using
    is for the generated test main package (pseudo-import path "testmain", package main)
    to import regexp and pass in a matchString function for use by testing when
    implementing the -run flags. This lets testing use regexp but without depending
    on regexp and creating unnecessary cycles.
    
    We want to add a few dependencies to runtime/pprof, notably regexp
    but also compress/gzip, without causing those packages to have to work
    hard to write internal tests.
    
    Restructure the (dare I say it) dependency injection of regexp.MatchString
    to be more general, and use it for the runtime/pprof functionality in addition
    to the regexp functionality. The new package testing/internal/testdeps is
    the root for the testing dependencies handled this way.
    
    Code using testing.MainStart will have to change from passing in a matchString
    implementation to passing in testdeps.TestDeps{}. Users of 'go test' don't do this,
    but other build systems that have recreated 'go test' (for example, Blaze/Bazel)
    may need to be updated. The new testdeps setup should make future updates
    unnecessary, but even so we keep the comment about MainStart not being
    subject to Go 1 compatibility.
    
    Change-Id: Iec821d2afde10c79f95f3b23de5e71b219f47b92
    Reviewed-on: https://go-review.googlesource.com/32455
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 5ad3bd99b5f997c55fa7c3d01eeeddc835631ba0
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 31 14:16:53 2016 -0400

    api: sort except.txt
    
    Make it easier to find lines and update the file.
    
    Change-Id: I9db78ffd7316fbc17c5488e178e23777756d8f47
    Reviewed-on: https://go-review.googlesource.com/32454
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit d5b97f614eb02399f7b4ed6615fae094362d151d
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 31 12:09:03 2016 -0400

    cmd/pprof: move cmd/internal/pprof back to cmd/pprof/internal
    
    CL 21870 moved the entire cmd/pprof/internal directory to cmd/internal/pprof
    for use by cmd/trace, but really cmd/trace only needed cmd/pprof/internal/profile,
    which became cmd/internal/pprof/profile, and then internal/pprof/profile.
    
    Move the rest back under cmd/pprof so that it is clear that no other code
    is reaching into the guts of cmd/pprof. Just like functions should not be
    exported unless necessary, internals should not be made visible to more
    code than necessary.
    
    Raúl Silvera noted after the commit of CL 21870 that only the profile package
    should have moved, but there was no followup fix (until now).
    
    Change-Id: I603f4dcb0616df1e5d5eb7372e6fccda57e05079
    Reviewed-on: https://go-review.googlesource.com/32453
    Run-TryBot: Russ Cox <rsc@golang.org>
    TryBot-Result: Gobot Gobot <gobot@golang.org>
    Reviewed-by: Ian Lance Taylor <iant@golang.org>

commit 682ffae6db749ba63df4b8bc1739974346bb14d7
Author: Russ Cox <rsc@golang.org>
Date:   Mon Oct 31 12:07:13 2016 -0400

    internal/pprof/profile: new package, moved from cmd/internal/pprof/profile
    
    This allows both the runtime and the cmd/pprof code to use the package,
    just like we do for internal/trace.
    
    Change-Id: I7606977284e1def36c9647354c58e7c1e93dba6b
    Reviewed-on: https://go-review.googlesource.com/32452
    Run-TryBot: Russ Cox <rsc@golang.org>
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 53cc69170ae1a96163d3b6c14467e85dc8aa7266
Author: Michael Munday <munday@ca.ibm.com>
Date:   Wed Nov 2 11:41:40 2016 -0400

    bytes, strings: update s390x code to match amd64 changes
    
    Updates the s390x-specific files in these packages with the changes
    to the amd64-specific files made during the review of CL 31690. I'd
    like to keep these files in sync unless there is a reason to
    diverge.
    
    Change-Id: Id83e5ce11a45f877bdcc991d02b14416d1a2d8d2
    Reviewed-on: https://go-review.googlesource.com/32574
    Reviewed-by: Brad Fitzpatrick <bradfitz@golang.org>

commit 235f2c72e95d486b530b45e1f2a48fb629260504
Author: Jan Mercl <0xjnml@gmail.com>
Date:   Wed Nov 2 11:23:05 2016 +0100

    go1.8.txt: Add CL 25345.
    
    Change-Id: I436528a4f81634448a60b1183d1b65a3bf4f48c1
    Reviewed-on: https://go-review.googlesource.com/32590
    Reviewed-by: Josh Bleecher Snyder <josharian@gmail.com>
```
