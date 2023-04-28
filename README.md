# chatgpt-wework-robot
企业微信应用chatgpt

1. 申请域名，并配置解析 到自己服务器的IP

2. 安装nginx和配置ssl证书（安全） 创建  wxwork.conf 文件，并放到nginx的配置路径，比如 /etc/nginx/site-enable/

upstream work {
    server 127.0.0.1:6060;
    keepalive 300;
}

server {

        listen 443 ssl http2;
        listen [::]:443 ssl default_server;

        root /var/www/html;

 	  server_name xxx.xxx.com;

        index index.html index.htm index.php index.nginx-debian.html;


        ssl_certificate  xxx.pem;
        ssl_certificate_key xxx.key;     


        location / {
                  proxy_pass  http://work;
        }


}

3.  创建 wxwork.yaml配置文件,和主程序同一个目录，内容如下,修改成自己对应的值
	
Web:
    Port: ":6060"
    RootPath: "/var/www/html"
OpenAI:
    BaseURL: ""
    APIKEY: ""
    Engine: ""
WeiXin:
    CorpID: ""
    AgentID: 0
    Token: ""
    AESKey: ""
    Secret: "

参数说明：
    baseurl： azure 的入口地址
   engine： 模型部署名称
    apikey：  openAI的 可以，或者azure的key ,如果用的是openai的接口，上面的baseurl和engine留空就可以

4.启动服务器

5.配置企业微信，添加 企业app
   
