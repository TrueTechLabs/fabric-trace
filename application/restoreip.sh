#!/bin/bash

# 要处理的文件列表
files=(
    "web/.env.development"
    "web/.env.production"
    "web/src/router/index.js"
    "backend/settings/config.yaml"
)

# 遍历文件列表并替换 IP 地址
for file in "${files[@]}"; do
    if [ -f "$file" ]; then
        echo "正在处理文件: $file"
        # 使用 sed 替换文件中的 IP 地址
        sed -i 's/127.0.0.1/119.45.247.29/g' "$file"
        echo "$file 中的 IP 地址已替换。"
    else
        echo "文件 $file 不存在！"
    fi
done

echo "所有文件的 IP 地址已替换完成！"
