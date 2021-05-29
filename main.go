package main

import (
	"encoding/hex"
	"ponywilliam/go-qrcode-door/RFID"
	"ponywilliam/go-qrcode-door/qrcode"
	"time"
)
func RemoveRepByLoop(slc []string) []string {
	var result []string // 存放结果
	for i := range slc{
		flag := true
		for j := range result{
			if slc[i] == result[j] {
				flag = false  // 存在重复元素，标识为false
				break
			}
		}
		if flag {  // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}
func AddRfid(){
	for{
		rfids := RFID.GetNearRfid()
		for _,v := range rfids{
			RFID.RfidResult = append(RFID.RfidResult,hex.EncodeToString(v))
			RFID.RfidResult = RemoveRepByLoop(RFID.RfidResult)//根据算法原则，在这一层数据元素较少，个人感觉时间复杂度会更好。
		}
	}
}
func ResetRfid(){
	tiker := time.NewTicker(time.Second * 5)
	for{
		<-tiker.C
		RFID.RfidResult = RFID.RfidResult[0:0]
	}
}
func main(){
	go AddRfid()
	go qrcode.GetData()//时刻扫描，允许不租借物品出门
	go ResetRfid()
	select{}//防止主线程退出
}
