#!/bin/bash
# 关闭旧app容器
./stop_docker.sh > /dev/null 2>&1
# 复制区块链密钥证书材料
rm -rf backend/blockchain/network/organizations && mkdir -p backend/blockchain/network
cp -r ../blockchain/network/organizations backend/blockchain/network/organizations
# 构建 Docker 镜像
docker build -t fabric-trace-app .
echo "容器ID："
# 运行 Docker 容器
docker run -d -p 9090:9090 --name fabric-trace-app fabric-trace-app
# 检查容器是否启动成功
if docker ps | grep -q fabric-trace-app; then
    echo "系统已成功运行！ 在浏览器上打开服务器IP：9090 即可查看"
else
    echo "容器启动失败！请检查日志。"
fi
