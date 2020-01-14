## 配置编辑工具

go常用的编辑器包括但不限于：[GoLand](https://www.jetbrains.com/go/) ,[Visual Studio Code](https://code.visualstudio.com/)

目前我使用的是  GoLand

### 1、goland配置

打开goland设置
![image](https://github.com/sunct/learning-go/blob/master/images/tool-1.png)

- #### 配置go modules

![image](https://github.com/sunct/learning-go/blob/master/images/tool-2.jpg)

勾选 [ ] Enable Go Modules (vgo) integration
选择并配置。

- #### 自动格式化代码

可以通过添加一个File Watcher来在文件发生变化的时候调用gofmt进行代码格式化，具体方法是，点击Preferences -> Tools -> File Watchers，点加号添加一个go fmt模版，Goland中预置的go fmt模版使用的是go fmt命令，将其替换为gofmt，然后在参数中增加-l -w -s参数，启用代码简化功能。添加配置后，保存源码时，goland就会执行代码格式化了。

如图所示：

![image](https://github.com/sunct/learning-go/blob/master/images/tool-3.jpg)

