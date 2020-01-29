package supermanyuncv

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"os"
)


// Yundama 类型
type Yundama struct {
	Username string `json:"账号"`
	Password string `json:"密码"`
	ID string `json:"软件id"`

}


// Send 发送请求 获取验证码
func (y *Yundama) Send(imgName string) (info float64, code string) {
	f, _ := os.OpenFile(imgName, os.O_RDONLY, 0666)
	defer f.Close()
	d := make([]byte, 6000)
	f.Read(d)
	response, _ := http.PostForm("http://api2.sz789.net:88/RecvByte.ashx", url.Values{"username": {y.Username}, "password": {y.Password}, "softid": {y.ID}, "imgdata": {fmt.Sprintf("%x", d)}})
	if response != nil{
		defer response.Body.Close()
	}

	body, _ := ioutil.ReadAll(response.Body)

	var j map[string]interface{}
	json.Unmarshal([]byte(body), &j)
	
	info = j["info"].(float64)
	code = j["result"].(string)
	
	return info, code
}
