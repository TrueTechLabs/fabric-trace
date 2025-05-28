#!/bin/bash

# 定义文件列表
FILES=(
    "web/.env.development"
    "web/.env.production"  
    "web/src/router/index.js"
    "backend/settings/config.yaml"
)

# 定义预设的云服务器 IP
# 119.45.247.29
# 127.0.0.1
OLD_IP="119.45.247.29"


echo "选择更新的 IP 地址类型："
echo "1. 仅在虚拟机内运行 (包括使用内置浏览器打开网页，使用 IP 127.0.0.1)"
echo "2. 输入新的云服务器 IP"

while true; do
    read -p "请选择 (输入 1 或 输入新的云服务器 IP): " USER_INPUT
    # 根据用户输入决定新的 IP
    if [ "$USER_INPUT" == "1" ]; then
        NEW_IP="127.0.0.1"
        break
    elif [[ "$USER_INPUT" =~ ^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        IFS='.' read -ra ADDR <<< "$USER_INPUT"
        valid=true
        for i in "${ADDR[@]}"; do
            # 检查每个数字是否在0到255之间
            if [[ "$i" -lt 0 || "$i" -gt 255 ]]; then
                valid=false
                break
            fi
        done
        if [ "$valid" = true ]; then
            NEW_IP="$USER_INPUT"
            break
        fi
    fi
    echo "无效的输入，请输入有效的IPv4地址或选择1使用虚拟机IP。"
done

# 检查每个文件
for FILE in "${FILES[@]}"; do
    if [ -f "$FILE" ]; then
        echo "检查文件 $FILE..."
        # 检测文件中是否包含旧的云服务器 IP
        if grep -q "$OLD_IP" "$FILE"; then
            # 替换文件中的所有旧 IP 地址
            sed -i "s/$OLD_IP/$NEW_IP/g" "$FILE"
            echo "已将 $FILE 中的 IP 更新为 $NEW_IP。"
        else
            echo "文件 $FILE 中没有找到需要更新的 IP ($OLD_IP)。"
        fi
    else
        echo "警告: 文件 $FILE 不存在。"
    fi
done
