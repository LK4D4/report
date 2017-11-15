# November 15, 2017 Report

Number of commit/s: 621

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 9.987581754s to 9.935628112s, -0.52%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 23.742275068s to 23.945551277s, +0.86%
* github.com/prometheus/prometheus/cmd/prometheus: from 19.272130797s to 19.314289069s, +0.22%
* code.gitea.io/gitea: from 11.297784156s to 11.288622624s, ~

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 31244560 to 31311744, +0.22%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 40880080 to 40348536, -1.30%
* github.com/prometheus/prometheus/cmd/prometheus: from 64412576 to 63768112, -1.00%
* code.gitea.io/gitea: from 32604190 to 32541722, -0.19%
## Highlights: 

* [runtime: make LockOSThread/UnlockOSThread nested](https://github.com/golang/go/commit/c85b12b5796c7efd4d8311253208b47449161361)
* [cmd/compile: optimize signed non-negative div/mod by a power of 2](https://github.com/golang/go/commit/0011cfbe2b57b385bac25a3daf9de581ee263661)
* [runtime: don't start new threads from locked threads](https://github.com/golang/go/commit/2595fe7fb6f272f9204ca3ef0b0c55e66fb8d90f)
* [runtime: make it possible to exit Go-created threads](https://github.com/golang/go/commit/eff2b2620db005cb58c266c0f25309d6f466cb25)
* [math/big: implement Lehmer's GCD algorithm](https://github.com/golang/go/commit/1643d4f33a0ed45cef0f6d33aff207ad530f9c94)
* [cmd/compile: compiler support for buffered write barrier](https://github.com/golang/go/commit/7e343134d334f7317b342db19c3e90d1f3f200cc)
* [cmd/vet: tighten printf format error messages](https://github.com/golang/go/commit/fc768da8b8030e6f344be6bbc86ae08c30f02849)
* [encoding/json: disallow unknown fields in Decoder](https://github.com/golang/go/commit/2596a0c075aeddec571cd658f748ac7a712a2b69)
* [cmd/compile: specialize map creation for small hint sizes](https://github.com/golang/go/commit/fbfc2031a673c95700e46ddf56404a0f648fc8a9)
* [cmd/go: cache built packages](https://github.com/golang/go/commit/de4b6ebf5d0a12f57ace43948b8b1b90f200fae9)
* [cmd/go: cache successful test results](https://github.com/golang/go/commit/bd95f889cdd241202fac01b29a3f3d7c03131a20)
* [cmd/go: run vet automatically during go test](https://github.com/golang/go/commit/0d188752524282496ebd0ab4b382bb4ff8750c90)
* [cmd/go: allow -coverprofile with multiple packages being tested](https://github.com/golang/go/commit/283558e42b88a6afa39da6ad4ae87558dc053776)













## GIT Log

```
git log 273b657b4e970f510afb258aa73dc2e264a701e3..f71cbc8a96056248d6789581b214ac44a2e6f91e
```
