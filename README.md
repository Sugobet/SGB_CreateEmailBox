# SGB_CreateEmailBox

<br>

### 程序介绍:  

　　程序由 go 语言开发，可以大量注册可用的 [ggo.la](https://mail.ggo.net/) 随机邮箱账号<br>　　
已接入[超人云打码平台](http://www.chaorendama.com/)，只需在配置文件中添加账号密码软件 id,并且确保你的账户余额充足, 即可在运行程序时添加命令行参数:isYundama=true(默认 false, false 则为手动打码)  
+ demo:SGB_CreateEmailBox.exe -num 3 -isYundama=true  
<br>
　　若不使用云打码，想自己上 可不设置参数isYundama或将参数isYundama设置为false, 默认为false
<br><br>
  　　注册到的邮箱账号以及密码会根据配置文件的"保存邮箱账号文件名"参数保存到相应的文件中。
    <br>
　　(上面已有现成的exe可执行文件，可直接下载exe 无需下载源码即可使用)

### 命令行参数:  
　　Usage of SGB_CreateEmailBox.exe:  
　　　　-isCreateConfig  
　　　　是否重建配置文件 默认false  如果是true则创建或重置配置文件  
　　　　-isYundama  
　　　　是否使用云打码 默认false, true为使用云打码，false为不用云打码（使用手动）  
　　　　-num uint  
　　　　注册邮箱账号的个数  
    <br>
　　若只想程序正常运行只需添加 num参数和isYundama参数，若不需要使用云打码可不添加isYundama  
    　　当参数isCreateConfig为true时，程序只会为你重建配置文件，之后退出程序
<br>


### 配置文件(config.json):  
　　demo配置文件:  <br>　　{  
　　"保存邮箱账号文件名": "test1.txt",  
　　"代理ip": [],　　// 代理ip可为多个，若不使用代理ip请设置为[""]若不使用代理ip请设置为[""]若不使用代理ip请设置为[""],若需要使用一个或多个代理则["1.1.1.1:1234", "2.2.2.2:4321"...]  
　　"超人云打码用户信息": {　　// 若需使用云打码请填写以下信息  
　　　　"账号": "",  
　　　　"密码": "",  
　　　　"软件id": ""  
　　　　}  
　　}  
  配置文件必须跟随程序所在的路径
  <br>

### 补充说明:  
　　asImage.png(验证码图片)会自动生成在程序所在路径

<br>
<br>

# 免费开源 请别调包 更别倒卖
