## 安装go

- ### linux 环境下安装

安装包可从[go官网](https://golang.google.cn/dl/) 或者 [go 中文网](https://studygolang.com/dl) ，找到安装包路径。如图所示，根据自己需求下载。

![image](https://github.com/sunct/learning-go/blob/master/images/install-1.png)

在操作环境下，执行下面的命令进行下载。（截止时间：2020年01月14日14:14:59）

#### 1、下载安装包
```linux
wget https://studygolang.com/dl/golang/go1.13.6.linux-amd64.tar.gz
```
根据你当前时间使用最新版本。

#### 2、提取压缩包

提取压缩包到合适的目录（例如: /usr/local ，该目录一般是默认的目录），（注：根据自己的安装包修改命令中的版本号。）

```linux
sudo tar -xzf go1.13.6.linux-amd64.tar.gz -C /usr/local
```
解压完，可删除压缩包：
```linux
rm -rf go1.13.6.linux-amd64.tar.gz
```


#### 3、查看安装版本
```$xslt
go version
```

### 4、配置环境变量
```$xslt
vim /etc/profile
```
进入编辑界面后 Shift+G 跳转至尾行，按 o 新插入一行，输入如下增加或修改：

```$xslt
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
```
保存，并执行
```$xslt
source /etc/profile
```
### 5、查看环境变量信息
```$xslt
go env
```
信息如下：
```$xslt
GO111MODULE=""
GOARCH="amd64"
GOBIN="/usr/local/go/bin"
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/root/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"

```

### 6、go版本升级
先删除旧版本，只需删除/usr/local/go 文件即可。运行命令：
```$xslt
sudo rm -rf /usr/local/go
```
再根据此文进行 下载新版本包并安装即可。

- ### mac 环境下安装

#### 1、安装
使用 brew 运行命令：
```$xslt
brew install go
```

#### 2、查看安装版本
```$xslt
go version
```
#### 3、配置环境变量
编辑
```$xslt
vim ~/.bash_profile
```
保存，并执行
```$xslt
source ~/.bash_profile
```


#### 4、升级版本
```$xslt
brew upgrade go
```
#### 5、切换版本
```$xslt
brew switch go 1.13.6
```

#### 6、卸载
```$xslt
brew uninstall go
```



