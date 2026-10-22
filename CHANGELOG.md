# Changelog

## v1.1.1 (2026-10-22)

- Fork established from CockroachDB `release-19.2` at e9e3265.
- License changed to AGPL-3.0-or-later (`LICENSE`). The BSL change
  date on the forked code (2022-10-01, see
  `licenses/COCKROACHDB-LICENSE`) has passed, so the forked files are
  Apache-2.0-effective; the combined work is distributed under AGPL.
  Original notices are preserved in `licenses/`.
- Proprietary enterprise code removed: `pkg/ccl`,
  `c-deps/libroach/ccl`, `pkg/ui/distccl`, `pkg/ui/ccl`. The CCL-gated
  `pkg/cmd/cockroach` entry point is gone with it; `pkg/cmd/cockroach-oss`
  is the binary. `pkg/cmd/cockroach-short` and `pkg/cmd/workload` no
  longer pull CCL init hooks; CDC roachtests (`pkg/cmd/roachtest/cdc.go`)
  and the enterprise syntax generator (`pkg/sql/tests/rsg_test.go`) are
  removed.
- Module path is now `github.com/podomy/tardigrade`. Historical
  issue/PR links to upstream are left pointing at CockroachDB.
- Vendor phone-home endpoints neutered (UI registration endpoint
  emptied); no telemetry leaves the cluster.
- `make buildoss` is the supported build. The default `make build`
  target still references the removed CCL tree and is not supported.
