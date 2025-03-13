echo "============================================="
echo "         启动网络将清理旧的区块链网络        "
echo "============================================="
echo ""
echo "⚠️  清理网络将删除所有链上数据！"
echo "   若不希望清除链上数据，可执行以下命令启动网络："
echo "   docker start \$(docker ps -aq)"
echo ""
echo "============================================="
echo "            正在执行 stop.sh 脚本            "
echo "============================================="

# 关闭区块链浏览器
docker compose -f explorer/docker-compose.yaml down -v > /dev/null 2>&1

# 关闭区块链网络
./network.sh down  > /dev/null 2>&1

# 删除 organizations 目录
rm -rf explorer/organizations 

# 删除 MySQL 容器
docker rm -f fabrictrace-mysql > /dev/null 2>&1

echo ""
echo "✅ 区块链网络已关闭，相关内容已清空！"
