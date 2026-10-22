#!/usr/bin/env bash
# Local/CI packaging only. Never installs or publishes.
set -euo pipefail
cd "$(dirname "$0")/.."
version="${1:?usage: package_personal.sh VERSION OUTPUT_DIR}"
output="${2:?output directory required}"
[[ "$version" =~ ^[0-9]+\.[0-9]+\.[0-9]+-x\.[0-9]+$ ]] || { echo 'Expected X.Y.Z-x.N' >&2; exit 2; }
mkdir -p "$output"
output="$(cd "$output" && pwd)"
target_os="${GOOS:-$(go env GOOS)}"
target_arch="${GOARCH:-$(go env GOARCH)}"
case "$target_os/$target_arch" in darwin/arm64|darwin/amd64|linux/amd64|linux/arm64) ;; *) echo 'Unsupported target' >&2; exit 2;; esac
stage="$(mktemp -d)"
trap 'rm -rf "$stage"' EXIT
# Use the pinned metadata submodule, not a live cloud API.
git submodule update --init --recursive
go generate ./bundledmeta
CGO_ENABLED=0 GOOS="$target_os" GOARCH="$target_arch" go build -trimpath -tags aliyun_cli_packed_meta   -ldflags "-X github.com/aliyun/aliyun-cli/v3/cli.Version=$version -X github.com/aliyun/aliyun-cli/v3/cli.CommandName=aliyunx -X github.com/aliyun/aliyun-cli/v3/cli.SourceCommit=$(git rev-parse --short=12 HEAD)" -o "$stage/aliyunx" ./main
cp LICENSE "$stage/"
git rev-parse HEAD > "$stage/SOURCE_COMMIT"
git submodule status > "$stage/SUBMODULES"
go version -m "$stage/aliyunx" > "$stage/BUILD_INFO"
archive="aliyunx-$version-$target_os-$target_arch.tar.gz"
COPYFILE_DISABLE=1 tar -czf "$output/$archive" -C "$stage" aliyunx LICENSE SOURCE_COMMIT SUBMODULES BUILD_INFO
(cd "$output" && shasum -a 256 "$archive" > "$archive.sha256")
