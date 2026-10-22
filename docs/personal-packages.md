# Personal build channel

This branch adds build-only automation on top of the existing candidate branch.
It does not publish to OSS, PyPI, GitHub Releases or a Homebrew tap, install a CLI,
or access live cloud credentials. Push this branch to the personal fork to run;
workflow_dispatch is also available after the workflow reaches the default branch.
The four upstream PR branches remain separate and unchanged.

The first matrix targets macOS Intel/Apple Silicon and Linux amd64/arm64.
Linux MaxC bundles require the runner baseline (Ubuntu 22.04 / glibc 2.35);
they do not claim the official manylinux2014 compatibility. macOS packages are
not Developer ID signed or notarized. Windows packaging is not in this matrix.
Artifacts are CI candidates, not stable releases, and expire under Actions retention.
Source commit, checksums and dependency/build information accompany each package.
Dependency versions are recorded but are not yet a fully locked reproducible build.

The Python candidate includes the experimental bearer provider; the Go candidate
includes its credential interface and OAuth fixes. Packaging is not official
endorsement of cross-region bearer use. The RFC support boundaries still apply.
Go native mc and the old maxc wrapper redesign are not included in this branch.

## Homebrew follow-up

Use a personal homebrew-tap repository and separately named, keg-only formulae
so the official CLI is not overwritten. Publish immutable, tested versioned release
assets first; the formula then pins their URLs and SHA256. A Python bundle contains
its own runtime and must install the entire PyInstaller onedir, not just maxc.
Wire the intended Go credential executable explicitly; do not assume PATH selects
the patched CLI. Do not redirect the official plugin updater to personal packages
or claim Go's maxc wrapper already selects this bundle. That integration needs its
own isolated acceptance test before recommending an installation command.

## Local validation (2026-10-22)

- actionlint and bash syntax validation passed.
- Invalid version rejected before creating output.
- macOS arm64 archive built, checksum verified, extracted version/help passed in a temporary HOME.
- Race tests passed for config, cli and main.
- Other matrix platforms require GitHub Actions execution; they are not claimed tested locally.
