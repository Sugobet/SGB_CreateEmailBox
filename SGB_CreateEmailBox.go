package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"

	"./config"
	"./httpclient"
	supermanyuncv "./yuncv"
)

/*	version: 3.0.0

	邮箱官方登录地址：https://mail.ggo.net/
	支持收发邮件(发QQ邮箱可能进垃圾箱)
	请不要做调包侠;
		author: Sugobet
		github: https://github.com/Sugobet/SGB_GetMail
*/

var createNum uint
var isYundama bool
var isCreateConfig bool
var asImageName string = "asImage.png"
var apiImageURL string = "https://api.ggo.net/api.php?op=checkcode&code_len=4&font_size=18"
var (
	saveEmailFileName         string
	proxyAddr                 = make([]string, 1, 10)
	yUsername, yPassword, yID string
)

func randMailName(lens int) (mailname string) {
	stringList := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

	for i := 0; i <= lens; i++ {
		mailname += string(stringList[rand.Intn(52)])
	}
	return mailname
}

func randPassWord(lens int) (mailpw string) {
	stringList := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"

	for i := 0; i <= lens; i++ {
		mailpw += string(stringList[rand.Intn(62)])
	}
	return mailpw
}

func opFile(path string, mode int) *os.File {
	f, err := os.OpenFile(path, mode|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func testRegex(text, re string) bool {
	match, _ := regexp.MatchString(re, text)

	return match
}

// client 必须是自定义的httpclient.Client类型
// 的NewProxyClient方法封装后返回的http.Client对象
func getImageAndCookie(client *http.Client) (cookie string, err error) {
	request, _ := http.NewRequest("GET", apiImageURL, nil)

	res, err := client.Do(request)
	if res != nil {
		defer res.Body.Close()
	} else {
		err = errors.New("response is nil")
		return "", err
	}
	if err != nil {
		err = errors.New("踏马的, 是不是垃圾代理地址gg了")
		return "", err
	}

	cookie = res.Header["Set-Cookie"][0]
	r, _ := regexp.Compile("(PHPSESSID=.*);")
	cookie = r.FindString(cookie)

	body, _ := ioutil.ReadAll(res.Body)
	file := opFile(asImageName, os.O_WRONLY|os.O_TRUNC)
	defer file.Close()
	file.Write(body)

	return cookie, nil
}

func yundama() (code string) {
YDM:
	y := supermanyuncv.Yundama{Username: yUsername, Password: yPassword, ID: yID}
	info, code := y.Send(asImageName)
	infos := int(info)
	if infos == 0 {
		goto YDM
	} else if infos == -1 {
		goto YDM
	} else if infos == -2 {
		panic("你的超人云打码账号余额不足,请及时充值")
	} else if infos == -3 {
		panic("帐号未绑定软件ID	登录平台后台,查看是否绑定软件ID")
	} else if infos == -5 || infos == -9 {
		panic("用户校验失败	检查用户及密码是否正确")
	}

	return code
}

func newClient() *httpclient.Client {
	client := &httpclient.Client{Client: &http.Client{}}
	return client
}

func newConfig() *config.Config {
	return &config.Config{}
}

func saveMailAccountNumber(MailName, PassWord string) {
	Mn := MailName + "@ggo.la"
	file := opFile(saveEmailFileName, os.O_APPEND)
	defer file.Close()

	_str := "账号:" + Mn + "\t" + "密码:" + PassWord + "\n"
	file.WriteString(_str)
}

// client 必须是自定义的httpclient.Client类型
// 的NewProxyClient方法封装后返回的http.Client对象
func regEmail(client *http.Client, cookie, code, emailname, emailpw string) error {
	regMailURL := "https://api.ggo.net/box.php?op=email&callback=jQuery19106083270282093554_1579312976604&action=newreg&username=" + emailname + "&domain=%40ggo.la&password=" + emailpw + "&code=" + code
	request, _ := http.NewRequest("GET", regMailURL, nil)
	request.Header.Add("cookie", cookie)
	request.Header.Add("Host", "api.ggo.net")
	request.Header.Add("Referer", "https://mail.ggo.net/reg.html")

	res, err := client.Do(request)
	if res != nil {
		defer res.Body.Close()
	} else {
		return errors.New("response is nil")
	}
	if err != nil {
		return err
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if testRegex(string(body), "200") {
		saveMailAccountNumber(emailname, emailpw)
		fmt.Println("账号：", emailname, "已保存")
		return nil
	} else if testRegex(string(body), `\\u9a8c\\u8bc1\\u7801\\u9519\\u8bef`) {
		fmt.Println("由于输入的验证码是错误的，重新输入验证码")
		return errors.New("验证码错误")
	}
	return errors.New("未知错误")
}

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.UintVar(&createNum, "num", 0, "注册邮箱账号的个数")
	flag.BoolVar(&isYundama, "isYundama", false, "是否使用云打码 默认false, true为使用云打码，false为不用云打码（使用手动）")
	flag.BoolVar(&isCreateConfig, "isCreateConfig", false, "是否重建配置文件 默认false  如果是true则创建或重置配置文件")
	flag.Parse()

	if isCreateConfig {
		_, err := (newConfig()).CreateDefultConfigFile()
		if err != nil {
			panic(err)
		}
		fmt.Println("配置文件重建成功, 请自行根据个人需求去修改配置文件")
		os.Exit(0)
	}

	if createNum == 0 {
		fmt.Println("Where is your brain?")
		os.Exit(-1)
	}
}

func main() {
	config := newConfig()
	c, err := config.UnConfig()
	if err != nil {
		panic(err)
	}
	{
		saveEmailFileName = c.SaveEmailBoxNumberFileName
		proxyAddr = c.ProxyAddr
		yUsername = c.YundamaUserInfo.Username
		yPassword = c.YundamaUserInfo.Password
		yID = c.YundamaUserInfo.ID
	}

	randint := func() (int, bool) {
		if proxyAddr[0] == "" {
			return 0, false
		}
		return rand.Intn(len(proxyAddr)), true
	}

	{
		randProxyAddr, isProxy := randint()
		client := newClient()
		var c *http.Client
		if isProxy == false {
			c = &http.Client{}
		} else {
			c, err = client.NewProxyClient("http://" + proxyAddr[randProxyAddr])
			if err != nil {
				panic(err)
			}
		}

		for i := uint(0); i < createNum; i++ {
		E:
			emailName := randMailName(10)
			emailPW := randPassWord(8)
			cookie, err := getImageAndCookie(c)
			if err != nil {
				randProxyAddr, isProxy = randint()
				if isProxy {
					proxyAddr = append(proxyAddr[:randProxyAddr], proxyAddr[randProxyAddr+1:]...)
					c, err = client.NewProxyClient(proxyAddr[randProxyAddr])
					if err != nil {
						panic(err)
					}
					goto E
				} else {
					c = &http.Client{}
					goto E
				}
			}

			code := ""
			if isYundama {
				code = yundama()
			} else {
				print("请输入验证码(验证码图片在程序相对路径下 asImage.png) :")
				fmt.Scanln(&code)
			}
			err = regEmail(c, cookie, code, emailName, emailPW)
			if err != nil {
				goto E
			}

		}
	}
}
