# X distribution

The fork build channel publishes `aliyunx` version `3.5.1-x.1`. Commands are
`aliyunx` and `maxcx`; neither installs an `aliyun` or `maxc` alias. Homebrew
formulae have the same names and can coexist with the official installation.
The repository and maintenance branch names are unchanged. Upstream PR branches
remain separate; branding changes belong only to `build/personal-packages`.

Install with `brew install xuhuanxxx/tap/aliyunx xuhuanxxx/tap/maxcx`.
Use `brew upgrade aliyunx maxcx` to upgrade. AliyunX disables the official
self-updater so it cannot replace itself with an upstream binary.

MaxCX bundles Python and requires no project virtual environment. Its bearer
provider defaults to `aliyunx`; an explicitly configured executable is respected.
Existing auth/configuration and job formats remain compatible and are not migrated
by installation. The official Go `maxc` plugin wrapper is not redirected to MaxCX;
invoke `maxcx` directly.

Four CI platforms: macOS Intel/Apple Silicon and Linux amd64/arm64. Linux requires
glibc 2.35 or later. macOS binaries are not Developer ID signed/notarized. The
release records source revision and dependency/build metadata; this is not a
fully locked reproducible build. Version output includes the X version and source
revision. GitHub Release assets are immutable inputs to SHA256-pinned formulae.

Bearer use remains opt-in and experimental; successful cross-region tests are not
an official support commitment. Python 0.9.4 upstream has no explicit Python
license declaration in the inspected metadata; do not infer its license from the
Java tree. Personal release authorization does not resolve that upstream question.
