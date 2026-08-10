# Spongebob revamp — design

## Goal

Turn the repo around from "library nested under a subfolder, binary at root" to
the standard Go layout used by the author's other projects (e.g. `governor`):
library at the module root, CLI under `cmd/`. Fix the `rand.Seed` deprecation,
tidy the README, and wire up automated releases that publish a Homebrew
formula to `calini/homebrew-tap`.

## 1. Layout change

- Move `spongebob/text.go` and `spongebob/text_test.go` to the repository
  root. Package stays `spongebob`; add a short package doc comment (mirrors
  `governor`'s `// Package governor is...` convention).
- Move `main.go` to `cmd/spongebob/main.go` (package `main`), importing the
  root package via `github.com/calini/spongebob`.
- Delete the now-empty `spongebob/` directory.
- This also fixes a latent inconsistency: the README already documented
  `import "github.com/calini/spongebob"` for library use, but the real
  package previously lived at the `/spongebob` subpath, so that import never
  actually worked.

## 2. `rand.Seed` deprecation

- Delete the `init()` function and its `rand.Seed(time.Now().UTC().UnixNano())`
  call in `text.go`, along with the now-unused `time` import.
- As of Go 1.20, the global `math/rand` source auto-seeds; there is nothing to
  replace the call with (per the deprecation message: "there is no reason to
  call Seed with a random value").
- Bump `go.mod` from `go 1.13` to `go 1.21` — the floor needed for the
  auto-seed guarantee to apply. CI's `setup-go` step reads `go-version-file:
  go.mod`, so this also pins the toolchain used for build/test.

## 3. README cleanup

- Keep the joke line: `It uses _CUTTING EDGE_ technology like *MARKOV
  CHAINS™* to generate _REALISTIC_ SPonGeBoBⓇ text.️`
- Remove the two FOSSA badges (stale, not part of this revamp).
- Add badges: GitHub Actions CI status, latest GitHub release, Go Report
  Card, License (MIT).
- Update code snippets for the new layout:
  - `go get -u github.com/calini/spongebob` stays correct now that the
    library really is at the root.
  - "Building it manually" snippet builds `./cmd/spongebob`.
  - Add a Homebrew install section: `brew install calini/tap/spongebob`.

## 4. CI workflow (basic)

New `.github/workflows/ci.yml`, triggered on push to `master` and on PRs:

- `go build ./...`
- `go vet ./...`
- `go test ./... -race -count=1`
- `gofmt -l .` check (non-Windows)

No golangci-lint job, no `.golangci.yml`, no coverage threshold — kept lean
for a small demo library, per explicit choice over `governor`'s fuller CI.

## 5. Release automation

- `.goreleaser.yaml` (v2 config), mirroring `governor`'s:
  - `project_name: spongebob`
  - `builds`: `main: ./cmd/spongebob`, binary `spongebob`, `CGO_ENABLED=0`,
    `ldflags: -s -w`, targets linux/darwin/windows × amd64/arm64
    (no `-X main.version=...`: unlike `governor`, `main.go` has no `version`
    var for it to target, and adding one isn't otherwise needed)
  - `archives`: tar.gz, zip on Windows
  - `checksum`, `changelog` (sort asc, exclude `test:`/`chore:` commits)
  - `release.github`: owner `calini`, name `spongebob`
  - `brews`: pushes to `calini/homebrew-tap` using
    `{{ .Env.HOMEBREW_TAP_GITHUB_TOKEN }}`, description "CLI for writing
    SPonGeBob TeXT", license MIT, commit author Calin Ilie
    <calin@ilie.io>, install block `bin.install "spongebob"`, test block
    invoking the built binary.
- `.github/workflows/release.yml`, triggered on `v*` tags, runs
  `goreleaser/goreleaser-action@v7` with `GITHUB_TOKEN` and
  `HOMEBREW_TAP_GITHUB_TOKEN` secrets (the latter already exists on
  `calini/spongebob`).
- The hand-written `spongebob.rb` in `calini/homebrew-tap` is deleted as part
  of this work — from the next tag onward, GoReleaser generates and commits
  it automatically.

## Out of scope

- No `--version` flag or build-time version injection.
- No golangci-lint / coverage gate (explicit choice above).
- No changes to `LICENSE.md` or the license itself (MIT, already correct).
