package main

import (
	"github.com/json-iterator/go/extra"
	openai "github.com/sashabaranov/go-openai"
	"github.com/xen0n/go-workwx"
)

func main() {

	conf.init()

	var wx = workwx.New(conf.WeiXin.CorpID)

	wxclient = wx.WithApp(conf.WeiXin.Secret, conf.WeiXin.AgentID)

	if conf.OpenAI.Engine == "" {
		openaiconfig = openai.DefaultConfig(conf.OpenAI.APIKEY)

		if conf.OpenAI.BaseURL != "" {
			openaiconfig.BaseURL = conf.OpenAI.BaseURL
		}

		openaiclient = openai.NewClientWithConfig(openaiconfig)
	} else {

		openaiconfig = openai.DefaultAzureConfig(conf.OpenAI.APIKEY, conf.OpenAI.BaseURL, conf.OpenAI.Engine)
		openaiclient = openai.NewClientWithConfig(openaiconfig)

	}

	wxclient.SpawnAccessTokenRefresher()

	extra.RegisterFuzzyDecoders()

	jsonhttp.init()

}
