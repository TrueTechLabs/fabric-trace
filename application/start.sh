# 构建前端静态目录
cd web
npm run build:prod
# 构建后端
cd ../backend
rm -rf dist
cp -r ../web/dist dist
go build -o fabrictrace
# 启动后端
nohup ./fabrictrace &!