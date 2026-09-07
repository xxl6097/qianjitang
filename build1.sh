#!/bin/bash
# 使用 UPX 压缩多平台交叉编译
#   cmd/attendance → dist-u/attendance-<goos>-<goarch>[.exe]
#   cmd/userbook   → dist-u/userbook-<goos>-<goarch>[.exe]
# 依赖：upx（macOS: brew install upx，或 https://github.com/upx/upx/releases）
# 可选：USE_GARBLE=1 时使用 garble 符号混淆后再压缩（依赖：go install mvdan.cc/garble@latest）
# 注意：UPX 4.x 起已移除 macOS Mach-O 支持，darwin 产物只编译不压缩。

set -e
cd "$(dirname "$0")"

if [ -f dist ]; then
    rm -rf dist
fi

mkdir -p dist

# 用法：build <GOOS> <GOARCH> <包路径> <输出文件>
build() {
  local goos=$1 goarch=$2 pkg=$3 out=$4
  echo "=== ${pkg} ${goos}/${goarch} → ${out} ==="
  # 先删除旧产物：go build 可能跳过重写未变化的输出，残留 UPX 压缩包会触发 AlreadyPacked
  rm -f "$out"
  if [ "${USE_GARBLE:-0}" = "1" ]; then
    CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch \
      garble -literals -seed=random \
        build -trimpath -ldflags "-s -w -buildid=" -buildvcs=false \
        -o "$out" "$pkg"
  else
    CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch \
      go build -trimpath -ldflags "-s -w -buildid=" -buildvcs=false \
        -o "$out" "$pkg"
  fi
}


# 平台列表：3 系统 × 2 架构
platforms=(
  "linux   amd64"
)

for plat in "${platforms[@]}"; do
  read -r goos goarch <<< "$plat"
  ext=""
  [ "$goos" = "windows" ] && ext=".exe"
  for pkg in ./cmd/agent ; do #./cmd/userbook
    out="dist/agent-${goos}-${goarch}${ext}"
    build "$goos" "$goarch" "$pkg" "$out"
  done
done

echo ""
echo "=== 产物 ==="
ls -lh dist/
