

**项目演示站点:** http://realcool.top:9090/

与B站官方合作，本项目代码讲解与二次开发课程已发布，支持试看: [B站：Fabric V2.5通用溯源项目讲解与二次开发课程](https://www.bilibili.com/cheese/play/ss15923?bsource=link_copy)

Fabric V2.5通用溯源课程活动：[腾讯文档：提交源码仓库地址/PPT/视频获得返现50-300元](https://docs.qq.com/form/page/DQ3ZJUERPVXJCdUtX)

#### 如果项目有帮助请给项目点上Star，我们将十分感谢！ 欢迎加入TrueTechLabs Fabric学习交流QQ群：776873343
群内可免费阅读此项目的区块链部分、应用后端部分的代码说明文档。项目二次开发以及添加IPFS上传文件功能等文档收录在专栏：[《Fabric项目学习笔记》](https://blog.csdn.net/qq_41575489/article/details/128637560)。


![TrueTechLabs Fabric学习交流QQ群：776873343](https://cdn.jsdelivr.net/gh/trueTechLabs/picture@pic/img202403062054644.jpg)
#### 一、项目介绍
基于区块链Hyperledger FabricV2.5的农产品溯源/商品/通用溯源应用模板，部署简单，附压测工具tape、区块链浏览器，文档详细。可以快速使用本系统搭建自己的溯源系统，帮助想法快速落地。采用的技术栈：Fabric V2.5、Gin、Vue.js、Mysql。
[10b244362bd18f0688ae6fd3cb62397](https://github.com/TrueTechLabs/fabric-trace/assets/49872902/60337eb8-1799-435f-b0a5-8ac61761aa28)




![多类型用户注册](https://cdn.jsdelivr.net/gh/trueTechLabs/picture@pic/img202403062029060.png)

![农产品上链](https://cdn.jsdelivr.net/gh/trueTechLabs/picture@pic/img202403062031538.png)
![农产品溯源](https://cdn.jsdelivr.net/gh/trueTechLabs/picture@pic/img202403062032189.png)
![区块链浏览器可视化](https://cdn.jsdelivr.net/gh/trueTechLabs/picture@pic/img202403062033224.png)

##### 项目地址：
**Github**:
[https://github.com/TrueTechLabs/fabric-trace](https://github.com/TrueTechLabs/fabric-trace)

**Gitee**:
[https://gitee.com/real__cool/fabric-trace](https://gitee.com/real__cool/fabric-trace)

##### 搭建视频地址：
[https://www.bilibili.com/video/BV1Ar421H7TK](https://www.bilibili.com/video/BV1Ar421H7TK)

##### 项目文档地址(部分内容需要订阅专栏）：
[https://blog.csdn.net/qq_41575489/category_12075943.html](https://blog.csdn.net/qq_41575489/category_12075943.html)

#### 二、版权声明
本项目基于Apache License 2.0开源协议，在个人的科研、学习范围内可以自由使用，请附上项目链接。如有商业需求或合作需求，需要联系作者购买授权。

#### 三、项目特点
本项目采用Hyperledger Fabric V2.5，属于目前最新的Fabric版本，具有更好的性能和稳定性，调用链码使用Fabric-gateway模式，是当前版本的推荐方式。内置了tape压测工具，可以方便的对区块链网络进行压测；内置了区块链浏览器，可以方便地查询交易信息。
项目结构清晰，代码注释详细，方便二次开发。结合了mysql实现账户注册登录功能，更符合真实业务场景。

#### 四、项目背景
区块链技术的出现，为溯源系统的建设提供了新的思路。区块链技术的不可篡改性、去中心化、可追溯等特点，使得区块链技术成为溯源系统的理想选择。本项目基于Hyperledger Fabric V2.5，实现了一个农产品溯源系统。 在本区块链系统中，有5个内置的角色：种植户、工厂、驾驶员、商店、消费者。其中种植户、工厂、驾驶员、商店可以将信息上链，消费者有信息溯源权限。上述可以上链信息的角色各可以输入5个农产品的属性，方便二次开发。本项目的目标是作为Fabric V2.5下的一个通用溯源模板。

#### 五、搭建步骤
> 如果部分内容与视频不一致请以本文档为准

 强烈推荐：使用云服务器搭建本系统，虚拟机问题较多。点击此链接购买腾讯云服务器：[https://curl.qcloud.com/Sjy0zKjy](https://curl.qcloud.com/Sjy0zKjy) 点击首单特惠，购买2核4G或以上的服务器，218/年（价格经常会调整），如果后续准备做程序开发可以用新用户优惠买三年的，安装Ubuntu20.04系统。

严格按照以下步骤操作：

1. 安装docker 

	```bash
	#下载docker 
	curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun 
	#添加当前用户到docker用户组 
	sudo usermod -aG docker $USER 
	newgrp docker 
	#配置docker镜像加速
	sudo mkdir -p /etc/docker

	sudo tee /etc/docker/daemon.json <<-'EOF'
	{
	  "registry-mirrors": ["https://punulfd2.mirror.aliyuncs.com"]
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
	# 仅在首次使用执行：下载Fabric Docker镜像
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

5. 启动后端 在fabric-trace/application/backend目录下： 执行： `go run main.go`

6. 修改后端IP，将以下文件中的IP：`119.45.247.29`，换成自己云服务的IP。
	```bash
	fabric-trace/application/web/.env.development
	fabric-trace/application/web/src/router/index.js
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
8. 在腾讯云轻量应用服务器防火墙页面，放行TCP端口`8080,9090,9528`
![防火墙配置](https://cdn.jsdelivr.net/gh/trueTechLabs/picture@pic/img202403062033585.png)
9. 在浏览器中打开：http://云服务器IP:9528 即可看到前端页面。
10. 关闭项目步骤
前后端使用ctrl+c指令即可关闭，区块链部分在network目录下使用./stop.sh脚本关闭，如果不正确关闭可能会影响下一次启动，处理方案是使用命令关闭相关进程。
11. 使用tape对项目进行压力测试
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


#### 六、本项目相关的后续计划：

1. 本项目将持续维护，欢迎给项目点亮Star与B站三连，非常感谢！本项目的区块链部分、前后端的说明文档以及简易的二次开发流程已在[《Fabric项目学习笔记 》](https://blog.csdn.net/qq_41575489/category_12075943.html)专栏发布,欢迎订阅支持！
2. （付费）本系统的讲解课程（约90分钟）已发布：[B站：Fabric V2.5通用溯源项目讲解与二次开发](https://www.bilibili.com/cheese/play/ss15923?bsource=link_copy)
3. 支持Docker方式部署，简化部署步骤
#### 七、目前已知存在的问题：
1. 由于网络原因，区块链浏览器镜像拉取较慢，因为不是Docker官方的仓库，需要耐心等待。
2. 区块链浏览器有时候会出现无法访问的情况，可以尝试重启浏览器容器。

#### 如果不能按照步骤运行项目
上述部署步骤已经多人验证并顺利完成，如果您通过上述步骤未能运行项目，请检查环境是否与本项目要求的一致，任何修改或遗漏步骤都可能引起项目不能正常运行，请严格按照视频与文章步骤再次尝试。若还是有问题请在文章评论区留言，附上第一次遇到的报错，如果问题不够明确，我们也很难帮助到您。

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=TrueTechLabs/fabric-trace&type=Date)](https://star-history.com/#TrueTechLabs/fabric-trace&Date)
