echo "===启动网络前先清理旧的区块链网络==="
echo "===========执行stop.sh脚本=========="
# 关闭区块链浏览器
docker compose -f explorer/docker-compose.yaml down -v > /dev/null 2>&1
# 关闭区块链网络
./network.sh down  > /dev/null 2>&1
# 删除organizations
rm -rf explorer/organizations 
# 删除mysql容器
docker rm -f fabrictrace-mysql > /dev/null 2>&1
echo "==区块链网络已关闭，相关内容已清空=="