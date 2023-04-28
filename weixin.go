package main

import (
	"log"
	"sync"

	openai "github.com/sashabaranov/go-openai"
	"github.com/xen0n/go-workwx"
)

type dummyRxMessageHandler struct{}

var wxclient *workwx.WorkwxApp

var msgMap sync.Map

var _ workwx.RxMessageHandler = dummyRxMessageHandler{}

// OnIncomingMessage 一条消息到来时的回调。
func (dummyRxMessageHandler) OnIncomingMessage(msg *workwx.RxMessage) error {
	// You can do much more!
	//fmt.Printf("incoming message: %s\n", msg)

	if content, ok := msg.Text(); ok {
		//content.GetContent()
		go sendMessage(msg.FromUserID, content.GetContent())
	}

	return nil
}

func sendMessage(userID string, question string) {

	resp := &workwx.Recipient{UserIDs: []string{userID}}

	var mlist []openai.ChatCompletionMessage

	m1 := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	}

	if m, ok := msgMap.Load(userID); ok {

		mm := m.([]openai.ChatCompletionMessage)

		if len(mm) > 100 || question == "复位" || question == "结束会话" || question == "重新开始" || question == "reset" {
			msgMap.Delete(userID)
			wxclient.SendTextMessage(resp, "会话结束，请重新提问", false)
			return

		} else {

			mlist = append(mm, m1)
		}
		msgMap.Store(userID, mlist)

	} else {

		mlist = append(mlist, m1)
		msgMap.Store(userID, mlist)
	}

	m, err := chat(mlist)
	if err != nil {
		return
	}

	log.Printf("User %v :%v \n\t\t GPT:%v \n\n", userID, question, m)

	wxclient.SendTextMessage(resp, m, false)

}
