#!/bin/bash

# 定义文件列表
FILES=(
    "web/.env.development"
    "web/.env.production"  
    "web/src/router/index.js"
)

# 定义预设的云服务器 IP
OLD_IP="119.45.247.29"

# 检查每个文件
for FILE in "${FILES[@]}"; do
    if [ -f "$FILE" ]; then
        echo "检查文件 $FILE..."
        # 检测文件中是否包含旧的云服务器 IP
        if grep -q "$OLD_IP" "$FILE"; then
            echo "文件 $FILE 包含旧的云服务器 IP ($OLD_IP)。需要更新。"
            echo "选择更新的 IP 地址类型："
            echo "1. 虚拟机 (使用 IP 127.0.0.1)"
            echo "2. 输入新的云服务器 IP"
            read -p "请选择 (输入 1 或 输入新的云服务器 IP): " USER_INPUT

            # 根据用户输入决定新的 IP
            if [ "$USER_INPUT" == "1" ]; then
                NEW_IP="127.0.0.1"
            else
                NEW_IP="$USER_INPUT"
            fi

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
