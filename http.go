package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/xen0n/go-workwx"
)

type jsonAPI struct{}

var jsonhttp = &jsonAPI{}

func (j *jsonAPI) init() {

	j.msghttp()
	select {}

}

func (j *jsonAPI) msghttp() {

	hh, err := workwx.NewHTTPHandler(conf.WeiXin.Token, conf.WeiXin.AESKey, dummyRxMessageHandler{})
	if err != nil {
		fmt.Println("parm error:", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/", hh)

	log.Println("Server started !")

	err = http.ListenAndServe(conf.Web.Port, mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (j *jsonAPI) msghttp2() {

	router := mux.NewRouter()

	router.HandleFunc("/message", j.httpMessage)

	router.Handle("/", http.FileServer(http.Dir(conf.Web.RootPath)))

	if conf.Web.SSLcrt == "" || conf.Web.SSLkey == "" {
		err := http.ListenAndServe(":"+conf.Web.Port, router)
		if err != nil {
			fmt.Println("start http err:", err)
		}

	} else {

		//http.ListenAndServe(":9999", nil)
		err := http.ListenAndServeTLS(":"+conf.Web.Port, conf.Web.SSLcrt, conf.Web.SSLkey, router)
		if err != nil {
			fmt.Println("start http ssl err:", err)
		}
	}

}

func (j *jsonAPI) httpMessage(w http.ResponseWriter, req *http.Request) {

	// hh, err := workwx.NewHTTPHandler(conf.WeiXin.Token, conf.WeiXin.AESKey, dummyRxMessageHandler{})
	// if err != nil {
	// 	panic(err)
	// }

}
