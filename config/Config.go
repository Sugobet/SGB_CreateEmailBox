package config

import (
	"os"
	"io/ioutil"
	"encoding/json"

	"../yuncv"
)


// Config 类型
type Config struct{
	// CreateEmailBoxNumbern uint `json:"创建邮箱账号个数"`
	SaveEmailBoxNumberFileName string `json:"保存邮箱账号文件名"`
	ProxyAddr []string `json:"代理ip"`
	YundamaUserInfo supermanyuncv.Yundama `json:"超人云打码用户信息"`
}

// CreateDefultConfigFile 创建默认配置文件
// 创建成功则返回 状态为true, error is nil
// 失败则返回false, error not is nil
func (c *Config) CreateDefultConfigFile() (bool, error) {
	file, err := os.OpenFile("config.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil{
		return false, err
	}
	defer file.Close()

	data, err := json.Marshal(Config{})
	if err != nil{
		return false, err
	}

	file.Write(data)

	return true, nil
}

// UnConfig 将配置文件(json)反序列化为结构体
func (c *Config) UnConfig() (*Config, error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil{
		return nil, err
	}
	con := &Config{}
	err = json.Unmarshal(data, con)
	if err != nil{
		return nil, err
	}

	return con, nil
}
