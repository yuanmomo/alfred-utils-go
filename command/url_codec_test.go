package command

// Package is called aw
import (
	"fmt"
	"log"
	"testing"
)

var testStrMap = map[string]string{
	"0;tEG2}VsWeS4&PIei&)8%3E*i(t+34=N:Hs 你好": "0%3BtEG2%7DVsWeS4%26PIei%26%298%253E%2Ai%28t%2B34%3DN%3AHs%20%E4%BD%A0%E5%A5%BD",
	";,/?:@&=+$":                   "%3B%2C%2F%3F%3A%40%26%3D%2B%24",
	"-_.!~*'()":                    "-_.%21~%2A%27%28%29",
	"#":                            "%23",
	"ABC abc 123":                  "ABC%20abc%20123",
	"22.546065|113.946405|深圳市|南山区": "22.546065%7C113.946405%7C%E6%B7%B1%E5%9C%B3%E5%B8%82%7C%E5%8D%97%E5%B1%B1%E5%8C%BA",
	`{"maxMember":2,"totalGameNums":10,"gameType":"doudizhu","autoPlayRoom":"2","keChai":"2","baseScore":2000,"maxLimit":1000000000,"arenaid":3,"limitProp":"2","force3":"1","multiple":5,"serverid":"2","showLeft":"2","force1":"1","wanfaType":"1","minLimit":100000,"roomChargeType":"1","roomLevel":3,"force2":"1"}`: `%7B%22maxMember%22%3A2%2C%22totalGameNums%22%3A10%2C%22gameType%22%3A%22doudizhu%22%2C%22autoPlayRoom%22%3A%222%22%2C%22keChai%22%3A%222%22%2C%22baseScore%22%3A2000%2C%22maxLimit%22%3A1000000000%2C%22arenaid%22%3A3%2C%22limitProp%22%3A%222%22%2C%22force3%22%3A%221%22%2C%22multiple%22%3A5%2C%22serverid%22%3A%222%22%2C%22showLeft%22%3A%222%22%2C%22force1%22%3A%221%22%2C%22wanfaType%22%3A%221%22%2C%22minLimit%22%3A100000%2C%22roomChargeType%22%3A%221%22%2C%22roomLevel%22%3A3%2C%22force2%22%3A%221%22%7D`,
}

var encodeContainsAndMap = map[string]string{
	`sessionid=ec06f00000850fa04901910293000706&version=1.01.0071&position=22.546065%7C113.946405%7C%E6%B7%B1%E5%9C%B3%E5%B8%82%7C%E5%8D%97%E5%B1%B1%E5%8C%BA&gameType=doudizhu&roomType=4&action=createroom&userid=DT100285&ruleInfo=%7B%22maxMember%22%3A2%2C%22totalGameNums%22%3A10%2C%22gameType%22%3A%22doudizhu%22%2C%22autoPlayRoom%22%3A%222%22%2C%22keChai%22%3A%222%22%2C%22baseScore%22%3A2000%2C%22maxLimit%22%3A1000000000%2C%22arenaid%22%3A3%2C%22limitProp%22%3A%222%22%2C%22force3%22%3A%221%22%2C%22multiple%22%3A5%2C%22serverid%22%3A%222%22%2C%22showLeft%22%3A%222%22%2C%22force1%22%3A%221%22%2C%22wanfaType%22%3A%221%22%2C%22minLimit%22%3A100000%2C%22roomChargeType%22%3A%221%22%2C%22roomLevel%22%3A3%2C%22force2%22%3A%221%22%7D`:`sessionid=ec06f00000850fa04901910293000706&version=1.01.0071&position=22.546065|113.946405|深圳市|南山区&gameType=doudizhu&roomType=4&action=createroom&userid=DT100285&ruleInfo={"maxMember":2,"totalGameNums":10,"gameType":"doudizhu","autoPlayRoom":"2","keChai":"2","baseScore":2000,"maxLimit":1000000000,"arenaid":3,"limitProp":"2","force3":"1","multiple":5,"serverid":"2","showLeft":"2","force1":"1","wanfaType":"1","minLimit":100000,"roomChargeType":"1","roomLevel":3,"force2":"1"}`,
}


func Test_EncodeURL(t *testing.T) {
	log.Println("Encode test ########################",)
	for origin, encode := range testStrMap {
		encodeResult := EncodeURL(origin)

		if encodeResult != encode {
			error := fmt.Sprintf("Encode ERROR:\n origin:[%v], \n result:[%v],\n encode:[%v] \n", origin, encodeResult, encode)
			log.Println(error)
			t.Error(error) // 如果不是如预期的那么就报错
		}else{
			log.Printf("\n origin:[%v], \n result:[%v],\n encode:[%v] \n", origin, encodeResult, encode)
		}
	}
	t.Log("Encode test finish success!!!") //记录一些你期望记录的信息
}

func Test_DecodeURLContainsAnd(t *testing.T) {
	fmt.Printf("Deocde test ######################## \n",)
	for origin, encode := range testStrMap {
		decodeResult := DecodeURLContainsAnd(encode)
		if decodeResult != origin {
			error := fmt.Sprintf("Decode ERROR:\n decode:[%v], \n result:[%v],\n origin:[%v] \n", encode, decodeResult, origin)
			log.Println(error)
			t.Error(error) // 如果不是如预期的那么就报错
		}else{
			log.Printf("\n decode:[%v], \n result:[%v],\n origin:[%v] \n", encode, decodeResult, origin)
		}
	}

	for encoded, origin := range encodeContainsAndMap {
		decodeResult := DecodeURLContainsAnd(encoded)
		if decodeResult != origin {
			error := fmt.Sprintf("Decode ERROR:\n decode:[%v], \n result:[%v],\n origin:[%v] \n", encoded, decodeResult, origin)
			log.Println(error)
			t.Error(error) // 如果不是如预期的那么就报错
		}else{
			log.Printf("\n decode:[%v], \n result:[%v],\n origin:[%v] \n", encoded, decodeResult, origin)
		}
	}
	t.Log("Decode test finish success!!!") //记录一些你期望记录的信息
}
