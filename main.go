package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
)

/*
	输入用户名密码来登录系统保存cookie
	连接chat,进行认证并进行conn的管理
	随时进行chat的收发
*/

var (
	MOBILE = "18201065931"
	PASSWD = "111"
	COOKIE = ""
	//conn *net.Conn
)

func main() {
	print("start go chat\n")

	res, err := http.PostForm("http://localhost:3000/api/pub/login", url.Values{"mobile": {MOBILE}, "passwd": {PASSWD}})
	if err != nil {
		log.Fatalf("login err, mobile : %s, passwd : %s, err : %v", MOBILE, PASSWD, err)
	}

	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("read res err : %v", err)
	}
	COOKIE = res.Cookies()[0].String()
	log.Printf("cookie : %s", res.Cookies()[0].String())

	//tcpAddr, err := net.ResolveTCPAddr("tcp", "123.59.64.205:7243")
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":3001")
	if err != nil {
		log.Fatalln("解析地址错误", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalln("连接chat接口错误", err)
	}


	go func(){
		//异步处理信息的接收与发送
	}()


	//认证规则:
	n := uint32(len([]byte(COOKIE)))
	nnw := bytes.NewBuffer(nil)
	err = binary.Write(nnw, binary.BigEndian, n)
	if err != nil {
		log.Printf("binary write err : %v", err)
	}

	//log.Printf("这是写入binary后为q%", nnw.Bytes())
	//log.Printf("这是写入binary后为s%", string(nnw.Bytes()))

	//var nb uint32
	//binary.Read(nnw, binary.BigEndian, &nb)
	//log.Println("反解析回来的:", nb)

	nbb := nnw.Bytes()
	log.Printf("第一次发送的auth信息之前:", nbb)
	nbb = append(nbb, []byte(COOKIE)...)
	log.Printf("cookie:", []byte(COOKIE))
	log.Printf("第一次发送的auth信息:", nbb)
	conn.Write(nbb)









	for {
	}
}
