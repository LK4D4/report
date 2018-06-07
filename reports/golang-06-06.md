# June 6, 2018 Report

Number of commits: 1553

## Compilation time

* github.com/coreos/etcd: from 10.71919656s to 10.690175004s, -0.27%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 27.006956158s to 27.711735784s, +2.61%
* github.com/prometheus/prometheus/cmd/prometheus: from 20.4609371s to 20.085572287s, -1.83%
* code.gitea.io/gitea: from 12.024799611s to 11.574938383s, -3.74%

## Binary size:

* github.com/coreos/etcd: from 33753992 to 45108176, +33.64%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 51005872 to 63255696, +24.02%
* github.com/prometheus/prometheus/cmd/prometheus: from 71371645 to 93033895, +30.35%
* code.gitea.io/gitea: from 36537556 to 49270637, +34.85%

## Highlights: 

* [os: add UserCacheDir](https://github.com/golang/go/commit/816154b06553a4cf8ee7ad089f5e444b37bed43d)
* [regexp: don't allocate when All methods find no matches](https://github.com/golang/go/commit/df5997b99b9a89e1198596366230fa6c4dd50b70)
* [sync: enable profiling of RWMutex](https://github.com/golang/go/commit/88ba64582703cea0d66a098730215554537572de)
* [cmd/compile: make math.Ceil/Floor/Round/Trunc intrinsics on arm64](https://github.com/golang/go/commit/07f0f0956355266dafc36aadb66928e7450210ea)
* [regexp: Regexp shouldn't keep references to inputs](https://github.com/golang/go/commit/7263540146c75de8037501b3d6fb64f59a0d1956)
* [cmd/compile: avoid mapaccess at ..](https://github.com/golang/go/commit/c12b185a6ed143e7b397bd58489866505756be0e)
* [cmd/trace: beautify goroutine page](https://github.com/golang/go/commit/ea1f4832401afb6bd89bf145db3791e7de6cadc4)
* [cmd/pprof: add readline support similar to upstream](https://github.com/golang/go/commit/3f89214940d1f922bc4fde923de658a2ec1e4ac3)

## GIT Log

```
git log go1.10..6decd3d984dd0bb213837b64ab6870568b33f197
```
