#!/bin/bash
./stop.sh
# 启动mysql容器
docker run --name fabrictrace-mysql -p 3337:3306 -e MYSQL_ROOT_PASSWORD=fabrictrace -d mysql:8
# 要检查的镜像版本
image_versions=("2.5.6")

# 要检查的镜像列表
images=("hyperledger/fabric-tools" "hyperledger/fabric-peer" "hyperledger/fabric-orderer" "hyperledger/fabric-ccenv" "hyperledger/fabric-baseos")

# 遍历镜像列表
for image in "${images[@]}"
do
    for version in "${image_versions[@]}"
    do
        # 使用docker images命令查找特定镜像和版本
        if ! docker images -a | grep "$image" | grep "$version" &> /dev/null
        then
            echo "镜像 $image:$version 不存在，开始拉取..."
            docker pull "$image:$version"
        fi
    done
done


# 启动区块链网络、创建通道
./network.sh up createChannel
# 部署链码，使用trace链码
./network.sh deployCC -ccn trace -ccp ../chaincode -ccl go
cp -r organizations explorer/

# 创建文件存储目录
mkdir -p ../../application/backend/files/uploadfiles
mkdir -p ../../application/backend/files/downloadfiles
mkdir -p ../../application/backend/files/images

# 启动区块链浏览器
docker compose -f explorer/docker-compose.yaml up -d
