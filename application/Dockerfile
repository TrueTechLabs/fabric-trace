
# 第一阶段：构建阶段
FROM golang:1.20 AS backendbuilder
ENV GOPROXY=https://goproxy.cn,direct
# 设置工作目录
WORKDIR /app
# 复制 go.mod 和 go.sum 文件
COPY backend/go.mod backend/go.sum ./
# 下载依赖
RUN go mod download
# 复制后端代码
COPY backend/ ./
COPY backend/pkg/fabric_docker ./pkg/fabric.go
COPY backend/blockchain/network/organizations /app/network/organizations
# 编译 Go 程序为静态二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# 使用轻量级 Node.js 镜像作为前端构建环境
FROM node:16-alpine AS webbuilder
# 设置工作目录
WORKDIR /app/web
# 配置 npm 镜像源为国内镜像
RUN npm config set registry https://registry.npmmirror.com
# 复制前端代码到镜像中
COPY ./web .
# 安装前端依赖并构建
RUN npm install && npm run build:prod
# 使用轻量级 Alpine Linux 镜像作为运行环境
FROM alpine:latest

# 设置工作目录
WORKDIR /app
# 复制前端构建的产物到运行环境
COPY --from=backendbuilder /app/settings/config_docker.yaml ./settings/config.yaml
COPY --from=backendbuilder /app/server ./server
COPY --from=backendbuilder /app/network/organizations /app/network/organizations
COPY --from=webbuilder /app/web/dist ./dist
# 设置可执行权限（如果需要）
RUN chmod +x ./server
# 暴露后端服务端口
EXPOSE 9090
# 启动后端服务
CMD ["./server"]
