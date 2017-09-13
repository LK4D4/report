# September 13, 2017 Report

Number of commits: 288

## Compilation time

* github.com/coreos/etcd/cmd/etcd: from 6.278164722s to 6.342214117s, +1.02%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 20.991429963s to 20.87406285s, -0.56%
* github.com/prometheus/prometheus/cmd/prometheus: from 15.176728466s to 15.153904527s, -0.15%
* code.gitea.io/gitea: from 7.604317098s to 7.564748685s, -0.52%

## Binary size:

* github.com/coreos/etcd/cmd/etcd: from 29753120 to 29488552, -0.89%
* github.com/grafana/grafana/pkg/cmd/grafana-server: from 38416752 to 38103232, -0.82%
* github.com/prometheus/prometheus/cmd/prometheus: from 61003712 to 60493171, -0.84%
* code.gitea.io/gitea: from 31733639 to 31431738, -0.95%
## Highlights: 

* [testing: parallelize tests over count](https://github.com/golang/go/commit/f04d5836181dec3ec1b7e427607f02fa7a204a2d)
* [strconv: optimize Atoi for common case](https://github.com/golang/go/commit/46aa9f5437b000fad3696b0cd9fd97995da16411)
* [fmt: document verbs %b %d %o %x %X for printing pointers](https://github.com/golang/go/commit/a4140b745ce22c56821750001f30fca4020b4650)
* [cmd/fix: rewrite x/net/context by default](https://github.com/golang/go/commit/bf90da97c1aaec78d2f8ad8b74a506d3b6f0ee75)
* [math/rand: add Shuffle](https://github.com/golang/go/commit/a2dfe5d278eae0864397a046a8206342a426d2bd)
* [runtime: improve timers scalability on multi-CPU systems](https://github.com/golang/go/commit/76f4fd8a5251b4f63ea14a3c1e2fe2e78eb74f81)
* [runtime: optimize siftupTimer and siftdownTimer a bit](https://github.com/golang/go/commit/6a7c63a08ab353c7e41cb24ae66e73fb3cb7cb56)









## GIT Log

```
git log 8f1e2a2610765528068107e33ab0d1d2ff224ce3..de2231888821add783305e7674bbb43d4d8453dc
```
