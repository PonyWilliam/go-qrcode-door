package qrcode

import (
	"encoding/json"
	"fmt"
	serial "github.com/tarm/goserial"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"ponywilliam/go-qrcode-door/RFID"
	"ponywilliam/go-qrcode-door/door"
	"ponywilliam/go-qrcode-door/speak"
	"time"
)
type Res1 struct{
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
var (
	s  io.ReadWriteCloser
	dataurl string
)
func FatalErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}
func Command(command []byte){
	//对执行Command一个封装
	_,err := s.Write(command)
	FatalErr(err)
}
func CloseConfigQrcode(){
	//不允许扫描配置码
	//07 C6 04 08 00 EC 00 FE 3B
	command := []byte{0x07,0xc6,0x04,0x08,0x00,0xec,0x00,0xfe,0x3b}
	Command(command)
}
func OnlyData(){
	//仅接受
	//07 C6 04 08 00 EB 00 FE 3C
	command := []byte{0x07,0xc6,0x04,0x08,0x00,0xeb,0x00,0xfe,0x3c}
	Command(command)
}
func SettingTime(){
	//08 C6 04 08 00 F3 03 0F FE 21 2FFE21
	command := []byte{0x08,0xc6,0x04,0x08,0x00,0xf3,0x03,0x2f,0xfe,0x21}//大概5s
	Command(command)
}
func GetData(){
	for{
		bufTemp := make([]byte, 200)
		for{
			num, err := s.Read(bufTemp)
			if err != nil{
				log.Fatal(err)
			}
			door.Send([]byte("0"))
			if num > 0 {
				//strTemp := hex.EncodeToString(bufTemp)
				strTemp := string(bufTemp[:num])
				fmt.Println(strTemp)
				if (strTemp[0] < '0' || strTemp[0] > '9')  &&(strTemp[0]<'a' || strTemp[0] > 'z') &&(strTemp[0] < 'A' || strTemp[0] > 'Z'){
					continue
				}
				//将RFID信息发出
				val := url.Values{}
				val.Set("secret",strTemp)
				var str string

				for _,v := range RFID.RfidResult{
					str += v + ","
				}
				fmt.Println(str)
				val.Set("rfid",str)
				temp,_ := http.PostForm(fmt.Sprintf("%v/work/byqr",dataurl),val)
				response,err := ioutil.ReadAll(temp.Body)
				res := &Res1{}
				err = json.Unmarshal(response,&res)
				if err != nil{
					//报警，二维码错误
					fmt.Println("二维码出错")
					go speak.SayFail()
					time.Sleep(time.Second *2)//休眠2s
					continue
				}
				//扫描成功，发送token，如果返回信息成功，则可以放行
				fmt.Println(res)
				if res.Code != 200 && (res.Code == 2003 || res.Code == 2004){
					go speak.SayFail()
					time.Sleep(time.Second *2)//休眠2s
					continue
				}
				if len(RFID.RfidResult) < 1{
					go speak.SayNoRfid()
					//放行
					door.Send([]byte("2"))
					continue
				}
				if res.Code != 200{
					go speak.SayNot()
					time.Sleep( time.Second * 2)
					continue
				}
				//成功,开门
				go speak.SayOpen()
				time.Sleep(time.Second *2)//休眠2s
			}
		}
	}
}
func UniCode(){
	//08 C6 04 08 00 F2 06 02 FE 2C
	command := []byte{0x08,0xc6,0x04,0x08,0x00,0xf2,0x06,0x02,0xfe,0x2c}
	Command(command)
}
func init(){
	dataurl = "https://hyh.dadiqq.cn"
	var err error
	cfg := &serial.Config{Name: "COM20", Baud: 115200, ReadTimeout: 50 /*毫秒*/}
	s,err = serial.OpenPort(cfg)
	FatalErr(err)
	CloseConfigQrcode()
	OnlyData()
	SettingTime()
	UniCode()
}