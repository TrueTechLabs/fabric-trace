## Docker版部署fabric-trace步骤
> 相较于主要的部署方式，不需要安装nvm，且使用docker部署项目更加方便管理

强烈推荐按照步骤使用云服务器搭建本系统，虚拟机的问题较多。点击此链接直达新人活动页面：[https://curl.qcloud.com/Sjy0zKjy](https://curl.qcloud.com/Sjy0zKjy) 点击首单特惠，**购买2核4G或以上的服务器**，199/年（价格经常会调整），如果后续准备做程序开发可以用新用户优惠买三年的，安装Ubuntu20.04系统。

**严格按照以下步骤操作，以下步骤已经经过上百人次的验证，如果遇到报错请仔细检查是否遗漏某个步骤。
如果对Linux命令不熟悉，请一定先学习下：[快速入门Linux及使用VSCode远程连接Linux服务器](https://blog.csdn.net/qq_41575489/article/details/139434933)，遇到报错后一定不能跳过，解决完再往下一步！**

1. 安装docker 

    ```bash
    #下载docker 
    # 官方脚本当前已无法下载，使用gitee备份的脚本:
    curl -fsSL https://gitee.com/real__cool/fabric_install/raw/main/docker_install.sh | bash -s docker --mirror Aliyun
    #添加当前用户到docker用户组 
    sudo usermod -aG docker $USER 
    newgrp docker 
    sudo mkdir -p /etc/docker
    #配置docker镜像加速，近期非常不稳定，如果以下源不好用可以再找下其他源
    #下边的源2024.8.29日测试可用
    sudo tee /etc/docker/daemon.json <<-'EOF'
    {
        "registry-mirrors": [
            "https://docker.m.daocloud.io",
            "https://docker.1panel.live",
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
    # 更新环境变量
    source  ~/.bashrc
    #安装jq 
    sudo apt install jq
    ```



3. 克隆本项目 

    ```bash
    git clone https://gitee.com/real__cool/fabric-trace
    ```

4. 启动区块链部分。在fabric-trace/blockchain/network目录下:

    ```bash
    # 仅在首次使用执行：下载Fabric Docker镜像。如果拉取速度过慢或失败请检查是否完成docker换源，或者更换一个其他的镜像源再试。
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

5. 修改后端IP，将以下文件中的IP：`119.45.247.29`，换成自己云服务器的IP。
    ```bash
    fabric-trace/application/web/.env.production
    fabric-trace/application/web/.env.development
    fabric-trace/application/web/src/router/index.js
    ```
    或使用application/replaceip.sh脚本根据指引修改IP，在fabric-trace/application目录下
    ```bash
    ./replaceip.sh
    ```

6. 启动app，在fabric-trace/application目录执行： 

    ```bash
    ./stop_docker.sh
    ```
8. 在腾讯云轻量应用服务器防火墙页面，放行TCP端口`8080,9090,9528`
![防火墙配置](https://truetechlabs-1259203851.cos.ap-shanghai.myqcloud.com/picgo202404151240899.png)
9. 在浏览器中打开：http://云服务器IP:9090 即可看到前端页面。
10. 使用tape对项目进行压力测试
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

```bash
./tape --config config_invoke.yaml -n 100
./tape --config config_query.yaml -n 100
```

11. 关闭项目


    ```bash
    cd application
    #关闭前后端
    ./stop_docker.sh
    #关闭区块链
    cd blockchain/network
    ./stop.sh
    ```


