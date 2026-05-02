# 🔗 Fabric V2.5 通用区块链溯源系统  
> 农产品 / 商品 / 通用溯源应用模板（Hyperledger Fabric）

基于 **Hyperledger Fabric V2.5** 的企业级区块链溯源系统模板，集成区块链网络、业务系统、区块链浏览器与性能压测工具，**部署简单、文档完善、支持快速二次开发**，可用于农产品溯源、商品溯源及通用可信存证场景。

📊 **项目累计部署 / 下载次数已超过 3000 次**

---

## 🌐 项目演示

**在线演示地址：**  
👉 https://t.realcool.top/

---

## 🎓 官方课程（B 站合作）

本项目已与 **B 站官方** 合作推出完整的源码讲解与二次开发课程：

🎥 **Fabric V2.5 通用溯源项目讲解与二次开发课程**  
👉 https://www.bilibili.com/cheese/play/ss15923

- 💰 限时特价：219 元（支持试看）
- 📅 发布时间：2024 年 4 月
- 📈 截至 2026-01-31：
  - 学员数：150+
  - 播放量：约 2 万
  - 完播率：60%

### 📦 付费课程包含

1. 项目源码讲解 + 二次开发视频
2. 一键部署脚本（减少约 90% 手动步骤）
3. 配置管理功能（5 分钟生成指定溯源应用）
4. 项目交流群（含助教答疑）
5. 项目全套资料  
   - IPFS 改进方案  
   - Caliper 压测代码  
   - 简易二次开发文档等  

📢 **课程活动**  
提交源码仓库 / PPT / 视频，可返现 **50–219 元**  
👉 https://docs.qq.com/form/page/DQ3ZJUERPVXJCdUtX

---

## ⭐ 项目影响力与推荐情况

- 📥 项目总下载量：约 3000 次（含Github）
- 📊 截至 2026-02-01（Gitee历史统计）：
  - Gitee 拉取量：1600+
  - Gitee ZIP 下载：281
- 🏅 2025 年被 **Gitee 官方评为推荐项目**
![Gitee评为推荐项目](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/202602060009685.png)
---

## 🤝 对外合作
- 🎓 与多个企业、高校合作完成真实区块链溯源项目
- 🏢 多家农产品企业、科研机构试用与关注
- ✅ 非常适合国内中小企业进行产品溯源落地

欢迎交流与合作 B站@TrueTechLabs

---

## 💬 学习与交流

⭐ 如果本项目对你有帮助，欢迎 **Star 支持**，这是对我们最大的鼓励！

📌 **TrueTechLabs Fabric 学习交流群（QQ）：776873343**  
群内可免费阅读本项目 **区块链部分与后端部分代码说明文档**

---

## 一、项目介绍

本项目是一个基于 **Hyperledger Fabric V2.5** 的通用区块链溯源系统模板，面向真实业务场景设计，可快速搭建具备完整业务流程的区块链应用。
![项目系统架构图](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/picgo319183738-60337eb8-1799-435f-b0a5-8ac61761aa28.png)
![多类型用户注册](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/picgo202404151236356.jpg)
![农产品上链](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/202602060035854.png)
![农产品溯源](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/202602060036273.png)
![区块链浏览器可视化](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/picgo202404151239574.png)


![配置管理功能](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/202602060034592.png)

### 技术栈

- 区块链：Hyperledger Fabric V2.5
- 后端：Go + Gin / Java Spring Boot
- 前端：Vue.js
- 数据库：MySQL
- 压测工具：Tape（课程包含 Caliper）


### 系统角色

系统内置 5 类角色：

- 种植户  
- 工厂  
- 驾驶员  
- 商店  
- 消费者（仅溯源查询）

支持多角色信息多字段上链，便于二次开发。

---

## 二、项目地址

- **GitHub**  
  https://github.com/TrueTechLabs/fabric-trace

- **Gitee（与 GitHub 同步）**  
  https://gitee.com/real__cool/fabric-trace

---

## 三、文档与视频资源

- 🎬 项目搭建视频  
  https://www.bilibili.com/video/BV1Ar421H7TK

- 🎥 完整未剪辑搭建视频  
  https://www.bilibili.com/video/BV1mF4m1P7Go

- 📖 项目文档（部分需订阅）  
  https://blog.csdn.net/qq_41575489/category_12075943.html

---

## 四、版权声明

本项目基于 **Apache License 2.0** 开源协议：

- ✅ 允许个人科研、学习使用（需注明项目来源）
- ❌ 未经授权禁止用于任何商业收费项目（含教学案例）

📄 商业授权 / 商务合作申请：  
👉 https://docs.qq.com/form/page/DQ1hIck5OQkNGQXF2  

若不认可本声明内容，请勿使用本项目。

---

## 五、项目特点

- 使用 Fabric 最新稳定版本 V2.5
- 采用官方推荐的 fabric-gateway 调用模式
- 内置区块链浏览器，交易可视化
- 集成 Tape / Caliper 压测方案
- 真实业务流程（注册、登录、上链、查询）
- 代码结构清晰、注释完善，适合二次开发

---

## 六、项目背景

区块链的不可篡改、可追溯特性，使其成为溯源系统的理想技术方案。

本项目的目标是提供一个 **可直接落地的 Fabric V2.5 通用溯源模板**，降低学习与部署成本，帮助想法快速转化为可运行系统。

---

## 七、搭建步骤


> ⚠️ 由于相关资源的失效，如果部分内容与视频不一致请以本文档为准

> ⚠️强烈推荐按照步骤使用全新的腾讯云轻量云服务器搭建本系统，虚拟机的问题较多。点击此链接直达新人活动页面：[https://curl.qcloud.com/Sjy0zKjy](https://curl.qcloud.com/Sjy0zKjy) ，**购买轻量2核4G或以上的服务器**，188/年（价格经常会调整），如果后续准备做程序开发可以用新用户优惠买三年的，安装**Ubuntu20.04**系统。系统版本不一致极有可能部署不成功。

> 注：不建议使用阿里云等其他云服务器，与演示步骤会不一致，且会有额外问题。



> 付费课程用户可使用一键部署脚本，减少 90% 操作步骤  
> 🚀 https://www.bilibili.com/video/BV119ffB4EoA

**严格按照以下步骤操作，以下步骤已经经过上千人次的验证，如果遇到报错请仔细检查是否遗漏某个步骤。已做好换源处理，不要使用任何代理工具。
如果对Linux命令不熟悉，请一定先学习下：[快速入门Linux及使用VSCode远程连接Linux服务器](https://blog.csdn.net/qq_41575489/article/details/139434933)，遇到报错后一定不能跳过，解决完再往下一步！**

1. 安装docker 

	```bash
	#下载docker 
	# 官方脚本当前已无法下载，使用gitee备份的脚本:
	# 本项目使用Docker版本：27.2.1 ，Ubuntu 20.04直接用以下指令即可
	# 近期源不稳定，如遇超时可使用：--mirror Aliyun
	curl -fsSL https://gitee.com/real__cool/fabric_install/raw/main/docker_install.sh | bash -s docker --mirror AzureChinaCloud
	#添加当前用户到docker用户组 
	sudo usermod -aG docker $USER 
	newgrp docker 
	sudo mkdir -p /etc/docker
	#配置docker镜像加速，近期非常不稳定，如果以下源不好用可以再找下其他源
	#下边的源2026.2.6日测试可用
	sudo tee /etc/docker/daemon.json <<-'EOF'
	{
	    "registry-mirrors": [
			"https://docker.1ms.run",
	        "https://docker.1panel.live",
			"https://docker.m.ixdev.cn",
			"https://docker.xuanyuan.me",
	        "https://hub.rat.dev"
	    ]
	}
	EOF

	#重启docker 
	sudo systemctl daemon-reload
	sudo systemctl restart docker
	```

2. 安装开发使用的go、node、jq

	```bash
	#下载二进制包
	wget https://golang.google.cn/dl/go1.19.linux-amd64.tar.gz
	#将下载的二进制包解压至 /usr/local目录
	sudo tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
	mkdir $HOME/go
	#将以下内容添加至环境变量 ~/.bashrc
	export GOPATH=$HOME/go
	export GOROOT=/usr/local/go
	export PATH=$GOROOT/bin:$PATH
	export PATH=$GOPATH/bin:$PATH
	#更新环境变量
	source  ~/.bashrc 
	#设置代理
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	
	#下载nvm安装脚本
	wget https://gitee.com/real__cool/fabric_install/raw/main/nvminstall.sh
	#安装nvm；屏幕输出内容添加环境变量
	chmod +x nvminstall.sh
	./nvminstall.sh
	# 将环境变量写入.bashrc
	export NVM_DIR="$HOME/.nvm"
	[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
	[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
	export NVM_NODEJS_ORG_MIRROR=http://npmmirror.com/mirrors/node/ #更换阿里云nvm node源
 
	# 更新环境变量
	source  ~/.bashrc
	# 安装node16
	nvm install 16
	#换源
	npm config set registry https://registry.npmmirror.com
	
	#安装jq 
	sudo apt install jq
	```



3. 克隆本项目 

	```bash
	git clone https://gitee.com/real__cool/fabric-trace
	```

4. 启动区块链部分。在fabric-trace/blockchain/network目录下:

	```bash
	# 仅在首次使用执行：下载Fabric Docker镜像。如果拉取速度过慢或失败请检查是否完成docker换源并执行了重启docker命令。
	./install-fabric.sh -f 2.5.6 d 
	```
	```bash
	# 启动区块链网络
	./start.sh
	```	
 	 **如果在启动区块链网络时遇到报错可以尝试:**
	```bash
	# 执行清理所有的容器指令：
	docker rm -f $(docker ps -aq)
	```
	**然后再重新启动区块链网络**

5. 启动后端（Go 或 Java 二选一）

	**Go 后端**（默认）：在 `fabric-trace/application/backend` 目录下：
	```bash
	go run main.go
	```

	**Java 后端**（Spring Boot）：在 `fabric-trace/application/backend-java` 目录下：
	```bash
	# 首次需编译打包
	mvn clean package -DskipTests
	# 启动服务
	java -jar target/trace-backend-1.0.0.jar
	```
	> Java 后端需要 JDK 17+ 和 Maven 环境，详细部署步骤见 [backend-java/README.md](application/backend-java/README.md)

6. 修改后端IP，将以下文件中的IP：`119.45.247.29`，换成自己云服务器的IP。
	```bash
	fabric-trace/application/web/.env.production
	fabric-trace/application/web/.env.development
	fabric-trace/application/web/src/router/index.js
	```
	或使用application/replaceip.sh脚本根据指引修改IP，在fabric-trace/application目录下
	```bash
	./replaceip.sh
	```

7. 新开一个窗口，启动前端 在fabric-trace/application/web目录下： 执行： 

	```bash
	# 仅在首次运行执行：安装依赖
	npm install 
	```

	```bash
	# 启动前端
	npm run dev
	```
8. 在腾讯云轻量应用服务器防火墙页面，放行TCP端口`8080,9090,9528`。如果是云服务器，则修改[安全组](https://cloud.tencent.com.cn/document/product/215/39790)。
![防火墙配置](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/picgo202404151240899.png)
9. 在浏览器中打开：http://云服务器IP:9528 即可看到前端页面。如果出现network error等网络报错，请按步骤6更换IP后重启项目。
10. 使用tape对项目进行压力测试（仅做测试demo，信息上链时有bug，建议使用课程里的caliper）
根据blockchain/chaincode/chaincode/smartcontract.go中的合约函数的签名，编写压测的参数，需要修改的内容是tape目录下的yaml文件中的args。args第一个参数是函数名，后面的参数是函数的参数。例如：
	```yaml
	args:
	# 函数名
	- RegisterUser
	# userID string
	- 1
	#userType string
	- randomString8
	# realInfoHash string
	- randomString8
	```
执行`./tape --config config_register.yaml -n 1`即可完成用户1的注册，然后可以对农产品上链操作与获取用户信息函数进行压测。更多的压测案例可以根据合约函数的签名进行修改。
附农产品上链操作与获取用户信息函数进行压测操作指令：

```	bash
./tape --config config_invoke.yaml -n 100
./tape --config config_query.yaml -n 100
```

#### 八、关闭项目与重新运行步骤
##### 关闭项目：
1. 前端（`npm run dev`界面）与后端（`go run main.go`界面：

	使用键盘组合键：`ctrl+c`

2. 区块链部分：

	在`fabric-trace/blockchain/network`目录`./stop.sh`，此步骤会清理所有的区块链、Mysql中的数据。

##### 开发模式启动项目：
1. 在`fabric-trace/blockchain/network`目录
`./start.sh` 如果遇到报错可以执行以下命令后再试：
执行清理所有的容器指令：
`docker rm -f $(docker ps -aq)`
2. 在`fabric-trace/application/backend`目录下： 执行： `go run main.go`
3. 在`fabric-trace/application/web`目录下： 执行：
`npm run dev`
4. 在http://服务器IP:9528打开

##### 生产环境部署项目(后台运行，访问速度更快)
1. 在`fabric-trace/blockchain/network`目录
`./start.sh` 如果遇到报错可以执行以下命令后再试：
执行清理所有的容器指令：
`docker rm -f $(docker ps -aq)`
2. 在`fabric-trace/application`目录下： 执行： `./start_prod.sh`
3. 在http://服务器IP:9090打开

	注意：此方式部署项目会在后台运行，如果后续遇到端口号占用可以尝试关闭占用9090端口号的进程，可以参考：
	[解决端口占用 bind:address already in use](https://blog.csdn.net/qq_41575489/article/details/137434008?spm=1001.2014.3001.5501)

#### 九、本项目相关的后续计划：

 本项目将持续维护，欢迎给项目点亮Star与B站三连，非常感谢！与B站官方合作，本项目代码讲解与二次开发课程已发布，限时特价219元，支持试看: [B站：Fabric V2.5通用溯源项目讲解与二次开发课程](https://www.bilibili.com/cheese/play/ss15923?bsource=link_copy),购买课程将赠送[《Fabric项目学习笔记》](https://blog.csdn.net/qq_41575489/article/details/128637560)中所有与本项目相关的资料。购买课程与订阅专栏均配套社群，方便交流与答疑。


#### 十、特别提示：
1. 使用虚拟机时区块链浏览器有时候会出现无法访问的情况，可以尝试重启浏览器容器。
2. 为了减少用户运行本项目时的难度，区块链目录的start.sh脚本在启动区块链时同时会清理掉所有的历史数据！如果重启机器后不希望清理原来的数据启动区块链，可以使用指令：`docker start $(docker ps -aq)`启动所有节点

#### 常见问题总结（检查第一个报错的位置）
上述部署步骤已经上百人次验证并顺利完成，如果您通过上述步骤未能运行项目，请检查环境是否与本项目要求的一致，任何修改或遗漏步骤都可能引起项目不能正常运行，请严格按照视频与文章步骤再次尝试或查看以下常见问题列表。若还是有问题请在[B站项目搭建视频](https://www.bilibili.com/video/BV1Ar421H7TK)评论区查看其他人的留言是否有相同的问题，如果还是没有解决请在视频下评论问题并附上与【安装步骤不一致的地方】或【第一个遇到的报错】，如果问题不够明确，我们也很难帮助到您。购买[B站：Fabric V2.5通用溯源项目讲解与二次开发课程](https://www.bilibili.com/cheese/play/ss15923?bsource=link_copy)可以加入配套社群，方便本项目交流与答疑。如果需要远程搭建服务或商业合作请填写收集表，对于政府等机构公益项目可以提供免费技术支持：[【腾讯文档】本项目搭建服务或商务合作意向收集](https://docs.qq.com/form/page/DQ1hIck5OQkNGQXF2)

1. 需要给机器安装mysql吗？
按照项目搭建过程即可部署好mysql，mysql容器与区块链节点一起启动，因此不需要单独安装mysql。
2. docker镜像拉取过慢或提示timeout
按照步骤docker换源并重启再试。
3. 安装链码时报错：exec: "go": executable file not found in $PATH
可能是go环境未安装好；可能使用了sudo，不要与安装步骤不一致！
4. jq:未找到命令
漏掉安装步骤中的安装jq，使用sudo apt install jq即可解决
5. docker权限报错：docker: permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock:
重新执行安装步骤中的：

	```bash
	#添加当前用户到docker用户组 
	sudo usermod -aG docker $USER 
	newgrp docker 
	```

6. 提示mysql或3337端口错误
mysql未启动，重启区块链网络部分可以一并重启mysql。
7. 前端提示network error、timeout
修改好IP后重启整个项目；检查防火墙是否放行相关端口。
8. 前端登录页面提示404
检查是否修改好IP并重启前端服务器，除了IP不要修改任何字符。
9. npm run dev 不能启动前端
检查npm install是否完整把所有包装上了。
10. Chaincode installation on peerg.orgl has failed
Deploying chaincode failed
此问题并非第一个报错，仔细检查之前的步骤中第一个报错，然后与上述问题对照。
11. eslint相关报错 此报错为代码规范性报错，不影响运行，可以根据提示自行修复（可不修复）
12. 其他问题
使用全新Ubuntu20.04系统，重新严格按照步骤再试，出现第一个报错后在视频下留言。



## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=TrueTechLabs/fabric-trace&type=Date)](https://star-history.com/#TrueTechLabs/fabric-trace&Date)
