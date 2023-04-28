package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

var jsonextra = jsoniter.ConfigCompatibleWithStandardLibrary

type config struct {
	Web struct {
		sslcrt   string `yaml:"SSLCRT" json:"ssl_crt"`
		sslkey   string `yaml:"SSLKey" json:"ssl_key"`
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

	confpath := dir + "/wxwork.yaml"

	cc := flag.String("c", confpath, "config file path and name")
	oo := flag.String("o", "", "print config content to stdout and exit , yaml or json format")

	flag.Parse()

	if *cc != "" {
		confpath = *cc
	}

	yamlFile, err := os.ReadFile(confpath)

	if err != nil {
		log.Printf("wxwork.yaml open err #%v ", err)
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
