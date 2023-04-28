# chatgpt-wework-robot
企业微信应用chatgpt

## 同时支持 openAI 服务器，openAI中转服务器和Azure服务器

#安装步骤


1. 申请域名，并配置A记录解析 到自己服务器的IP， 域名需要备案，否则腾讯无法验证
2. 安装nginx和配置ssl证书（安全） 将wxwork.conf 文件，并放到nginx的配置路径，比如 /etc/nginx/site-enable/3. 未域名申请SSL证书，并将证书放到路径，

3.  将wxwork.yaml配置文件放到主程序同一个目录：

```
Web:
    Port: ":6060"  //服务器端口，nginx 提供https服务器，反向代理到 本程序的端口，必须和nginx中代理目标端口一致
    RootPath: "/var/www/html"   //如果项目有静态网页，可以放在此目录
OpenAI:
    BaseURL: ""     //openAI的路径，为空使用默认路径， 也可以指定中转服务器的路径，比如： "http://chat.xxx.com/v1"  ，如果使用Azure服务，可以指定Azure的入口地址,比如： "htts://xxx.openai.azure.com"
    APIKEY: ""   // openAI的Key或者Azure的Key
    Engine: ""   // 使用Azure时的，模型部署名称， 为空代表不使用Azure服务，而使用openAI服务
WeiXin:
    CorpID: ""   //企业ID  包含以下参数，请从企业微信管理后台查找
    AgentID: 0   
    Token: ""    
    AESKey: ""
    Secret: "
```
	

4.启动服务器
```
#创建运行目录
    mkdir /wework
# 将编译好的执行文件放到  /wework
# 将 wework.services 放到 /lib/systemd/system 
# 打开开机启动，并启动服务：
    systemctl enable wework
    systemctl start wework

```

5.配置企业微信，添加 企业app
   
