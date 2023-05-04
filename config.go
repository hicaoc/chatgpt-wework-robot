package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

var jsonextra = jsoniter.ConfigCompatibleWithStandardLibrary

type config struct {
	Web struct {
		SSLcrt   string `yaml:"SSLCRT" json:"ssl_crt"`
		SSLkey   string `yaml:"SSLKey" json:"ssl_key"`
		Port     string `yaml:"Port" json:"port"`
		RootPath string `yaml:"RootPath" json:"root_path"`
	} `yaml:"Web" json:"web"`
	OpenAI struct {
		BaseURL string `yaml:"BaseURL" json:"base_url"`
		APIKEY  string `yaml:"APIKEY" json:"api_key"`
		Engine  string `yaml:"Engine" json:"engine"`
	} `yaml:"OpenAI" json:"openai"`
	WeiXin struct {
		CorpID  string `yaml:"CorpID" json:"corp_id" `
		AgentID int64  `yaml:"AgentID" json:"agent_id"`
		Token   string `yaml:"Token" json:"token"`
		AESKey  string `yaml:"AESKey" json:"aes_key"`
		Secret  string `yaml:"Secret" json:"secret"`
	} `yaml:"WeiXin" json:"weixin"`
}

var conf = &config{}

// LoadConfig load config from file
func (c *config) init() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("get filepath err #%v ", err)
		os.Exit(1)
	}

	confpath := dir + "/wework.yaml"

	cc := flag.String("config", confpath, "config file path and name")
	oo := flag.String("o", "", "print config content to stdout and exit , yaml or json format")
	help := flag.Bool("help", false, "Show help information")

	flag.Parse()

	// 如果用户指定了 help 参数，则输出帮助信息
	if *help {
		fmt.Println("Usage:")
		fmt.Println("  wework [OPTIONS]")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  --help             Show help information")
		fmt.Println("  --config  FILENAME  Config file")
		fmt.Println("  --o  Std Output config  format. Support yaml or json ")
		fmt.Println("Setting environment variables to override parameters in configuration files.:")
		fmt.Println("APPPORT BASEURL APIKEY ENGINE CORPID AGENTID TOKEN AESKEY SECRET")

		os.Exit(0)
	}

	if *cc != "" {
		confpath = *cc
	}

	yamlFile, err := os.ReadFile(confpath)

	if err != nil {
		log.Printf("wework.yaml open err #%v ", err)
		os.Exit(1)

	}
	err = yaml.Unmarshal(yamlFile, conf)

	if err != nil {
		log.Fatalf("Unmarshal: %v \n %s", err, yamlFile)
	}

	// c.Parm.iDCfilterIPMap = make(map[uint32]bool, 0)
	// for _, v := range c.Parm.IDCfilterIP {
	// 	c.Parm.iDCfilterIPMap[ipstrToUInt32(v)] = true
	// }
	readENV()

	if *oo != "" {
		if *oo == "json" {
			j, _ := jsonextra.MarshalIndent(conf, "", "    ")
			fmt.Println(string(j))
		} else if *oo == "yaml" {
			j, _ := yaml.Marshal(conf)
			fmt.Println(string(j))

		}
		os.Exit(0)
	}

}

func readENV() {

	port := os.Getenv("APPPORT")
	if port != "" {
		conf.Web.Port = ":" + port
	}
	rootpath := os.Getenv("ROOTPATH")
	if rootpath != "" {
		conf.Web.RootPath = rootpath
	}
	BaseURL := os.Getenv("BASEURL")
	if BaseURL != "" {
		conf.OpenAI.BaseURL = BaseURL
	}
	APIKEY := os.Getenv("APIKEY")
	if APIKEY != "" {
		conf.OpenAI.APIKEY = APIKEY
	}

	Engine := os.Getenv("ENGINE")
	if Engine != "" {
		conf.OpenAI.Engine = Engine
	}
	CorpID := os.Getenv("CORPID")
	if CorpID != "" {
		conf.WeiXin.CorpID = CorpID
	}
	AgentID := os.Getenv("AGENTID")
	if AgentID != "" {
		i, err := strconv.Atoi(AgentID)
		if err != nil {
			fmt.Println("Read AGENTID env parm error:", err)
		}
		conf.WeiXin.AgentID = int64(i)
	}
	Token := os.Getenv("TOKEN ")
	if Token != "" {
		conf.WeiXin.Token = Token
	}
	AESKey := os.Getenv("AESKEY")
	if AESKey != "" {
		conf.WeiXin.AESKey = AESKey
	}
	Secret := os.Getenv("SECRET")
	if Secret != "" {
		conf.WeiXin.Secret = Secret
	}

}
