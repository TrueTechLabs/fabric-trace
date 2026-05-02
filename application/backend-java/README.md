# Fabric-Trace 后端服务（Spring Boot）

基于 Spring Boot 2.7 的农产品溯源系统后端，通过 Hyperledger Fabric Gateway 与区块链网络交互，提供用户管理和溯源数据上链/查询接口。

## 技术栈

| 组件 | 版本 | 说明 |
|------|------|------|
| Spring Boot | 2.7.18 | Web 框架 |
| MyBatis-Plus | 3.5.5 | ORM，简化 MySQL 操作 |
| Hutool | 5.8.25 | 工具库（雪花ID、MD5、SHA256） |
| jjwt | 0.9.1 | JWT 令牌生成与解析 |
| Fabric Gateway | 1.7.0 | Hyperledger Fabric Java SDK |
| gRPC | 1.67.1 | 区块链网络通信 |
| MySQL | 5.7+ | 用户数据存储 |
| JDK | 17 | Java 运行环境 |
| Maven | 3.6+ | 构建工具 |

## 项目结构

```
src/main/java/com/fabric/trace/
├── BackendApplication.java          # 启动类
├── controller/
│   ├── UserController.java          # 用户注册/登录/登出/信息查询
│   └── TraceController.java         # 上链/溯源查询/图片获取
├── service/
│   ├── UserService.java             # 用户服务接口
│   ├── TraceService.java            # 溯源服务接口
│   ├── FabricService.java           # 区块链服务接口
│   └── impl/
│       ├── UserServiceImpl.java     # 用户业务实现
│       ├── TraceServiceImpl.java    # 溯源业务实现
│       └── FabricServiceImpl.java   # 区块链调用实现
├── mapper/
│   └── UserMapper.java              # MyBatis-Plus Mapper
├── entity/
│   └── User.java                    # 用户实体
├── config/
│   ├── CorsConfig.java              # 跨域配置
│   ├── WebMvcConfig.java            # 静态资源 + 拦截器注册
│   └── FabricConfig.java            # Fabric Gateway 连接配置
├── common/
│   └── Result.java                  # 统一响应封装
├── interceptor/
│   └── JwtAuthInterceptor.java      # JWT 认证拦截器
└── util/
    ├── JwtUtil.java                 # JWT 工具类
    └── SnowflakeUtil.java           # 雪花ID/MD5/SHA256 工具类
```

## API 接口

### 公开接口（无需认证）

| 方法 | 路径 | 说明 | 参数 |
|------|------|------|------|
| GET | `/ping` | 健康检查 | - |
| POST | `/register` | 用户注册 | username, password, userType |
| POST | `/login` | 用户登录 | username, password |
| POST | `/getFruitInfo` | 查询溯源信息 | traceability_code |
| GET | `/getImg/{filename}` | 获取图片 | 路径参数 filename |

### 认证接口（需 JWT Token）

请求头：`Authorization: <token>`

| 方法 | 路径 | 说明 | 参数 |
|------|------|------|------|
| POST | `/logout` | 登出 | - |
| POST | `/getInfo` | 获取用户信息 | - |
| POST | `/uplink` | 农产品上链 | arg1-arg5, traceability_code(非种植户), file(可选) |
| POST | `/getFruitList` | 获取用户产品列表 | - |
| POST | `/getAllFruitInfo` | 获取所有产品信息 | - |
| POST | `/getFruitHistory` | 获取溯源历史 | traceability_code |

### 响应格式

```json
// 成功
{"code": 200, "message": "success", "data": "..."}

// 登录成功
{"code": 200, "message": "login success", "jwt": "eyJ..."}

// 注册成功
{"code": 200, "message": "register success", "txid": "abc..."}

// 上链成功
{"code": 200, "message": "uplink success", "txid": "abc...", "traceability_code": "123..."}

// 错误
{"code": 500, "message": "错误信息"}

// 未认证
{"code": 401, "msg": "请求未携带token，无权限访问"}
```

## 环境要求

- **JDK 17+**
- **Maven 3.6+**
- **MySQL 5.7+**（需提前启动，默认端口 3337）
- **Hyperledger Fabric 网络**（需提前启动，默认端口 7051）

## 配置说明

配置文件位于 `src/main/resources/`：

| 文件 | 说明 |
|------|------|
| `application.yml` | 主配置（默认开发环境） |
| `application-dev.yml` | 开发环境覆盖配置 |
| `application-docker.yml` | Docker 环境覆盖配置 |

### 关键配置项

```yaml
# 服务端口
server:
  port: 9090

# MySQL
spring:
  datasource:
    url: jdbc:mysql://127.0.0.1:3337/fabrictrace?createDatabaseIfNotExist=true
    username: root
    password: fabrictrace

# JWT
jwt:
  secret: testsecret          # 生产环境请更换
  expiration: 172800000       # 48小时（毫秒）

# Fabric 区块链
fabric:
  msp-id: Org1MSP
  peer-endpoint: 127.0.0.1:7051
  tls-cert-path: ../../blockchain/network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
  cert-path: ../../blockchain/network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/User1@org1.example.com-cert.pem
  key-path: ../../blockchain/network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/
  channel-name: mychannel
  chaincode-name: trace

# 文件存储
file:
  upload-path: files/uploadfiles/
  image-path: files/images/
```

## 部署步骤

### 1. 安装 JDK 和 Maven

```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install -y openjdk-17-jdk maven

# 验证
java -version
mvn -version
```

### 2. 配置 Maven 阿里云镜像（推荐）

```bash
mkdir -p ~/.m2
cat > ~/.m2/settings.xml << 'EOF'
<?xml version="1.0" encoding="UTF-8"?>
<settings xmlns="http://maven.apache.org/SETTINGS/1.2.0"
          xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
          xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.2.0 https://maven.apache.org/xsd/settings-1.2.0.xsd">
    <mirrors>
        <mirror>
            <id>aliyun</id>
            <mirrorOf>central</mirrorOf>
            <name>Aliyun Maven Mirror</name>
            <url>https://maven.aliyun.com/repository/public</url>
        </mirror>
    </mirrors>
</settings>
EOF
```

### 3. 启动前置服务

确保以下服务已启动：

```bash
# MySQL（端口 3337）
# Hyperledger Fabric 网络（端口 7051）
# 具体启动方式参见项目根目录的 blockchain/ 和 docker-compose 配置
```

### 4. 编译打包

```bash
cd application/backend-java
mvn clean package -DskipTests
```

打包产物在 `target/trace-backend-1.0.0.jar`。

### 5. 运行

```bash
# 开发环境（默认配置）
java -jar target/trace-backend-1.0.0.jar

# 指定环境
java -jar target/trace-backend-1.0.0.jar --spring.profiles.active=dev

# 自定义端口
java -jar target/trace-backend-1.0.0.jar --server.port=8080
```

### 6. 验证服务

```bash
# 健康检查
curl http://127.0.0.1:9090/ping

# 注册用户
curl -X POST http://127.0.0.1:9090/register \
  -F "username=testuser" \
  -F "password=123456" \
  -F "userType=种植户"

# 登录
curl -X POST http://127.0.0.1:9090/login \
  -F "username=testuser" \
  -F "password=123456"
```


## 开发说明

- 数据库表会在首次启动时自动创建（`schema.sql`），不会重复创建
- Fabric Gateway 在应用启动时建立连接，作为 Spring 单例 Bean 管理
- JWT Token 有效期 48 小时，支持 `Bearer token` 和直接传 token 两种格式
- 上链时种植户自动生成溯源码，其他角色需提供已有的溯源码
