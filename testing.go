package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "#01748d7057801c7a6231103.131.147.44linux(1279047e,1cfbc4c8,12e568e0)#01748d705781a73bb2aa29.165.251.39windows(5d7723fc,9c4d497e,a6b4d6b9,d196b465,fff61779)#01748d705782e427a588103.131.147.44linux(7c1d4136)#01748d705783addee65a214.141.37.118android(d97f6468,a11ecb48,a11ecb48,c90e8092)#01748d705784aacf54a7160.34.215.31mac(bcd8f4a8,10243dc3,1279047e,c90e8092,a48327da)#01748d7057856d6b84137.224.203.105linux(8b0f1a05,ca155786)#01748d705786c08ee17f165.243.180.246ios(e02c3c77,7c1d4136,8b0f1a05,a2df52d5,29a04e11)#01748d705787286121fe159.11.38.79linux(ed8d0b69,786d0253)#01748d70578844f0c12a188.192.16.31ios(2e3f28bb,467880b7,c122ae6f,e1cc7a8d,17543175)#01748d70578971eff03d126.151.131.211linux(86a98679,2f15ff16,de49ec23,1c1b06ae)#01748d70578a8c3b851a183.253.207.120mac(754ccb27)#01748d70578b60a1764b95.183.100.154android(73870fc9,80891a08,a48327da)#01748d70578c5776022b187.127.182.168windows(864b40a0,a4f8dc54,fff61779,1c1b06ae)#01748d70578d5e33a92a192.223.83.209mac(a4f8dc54)#01748d70578eb192e681103.47.104.88android(ca155786,ca155786,c3ea0a90)#01748d70578f83ff3ba4155.101.36.169windows(81d51ead,b35e919,e6ff51ec,1a47362c,e7543b90)#01748d705790fbdb1dfc7.241.248.181linux(fc5de8c5,5d7723fc,1279047e)"
	trimed := strings.TrimSpace(s)
	split1 := strings.Split(trimed, "#")
	split2 := split1[1:len(split1)]

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	// submatchall := re.FindAllStringIndex(split1[1], 1)

	for i := 0; i < len(split2); i++ {
		fmt.Println("id: ", split2[i][0:12])
		fmt.Println("buyer_id: ", split2[i][12:20])
		ipIndexRange := re.FindAllStringIndex(split2[i][20:len(split2[i])-1], 1)[0]
		//ipStart := ipIndexRange[0][0:1]
		fmt.Println(ipIndexRange)
		fmt.Println("IP: ", split2[i][20 : len(split2[i])-1][ipIndexRange[0]:ipIndexRange[1]])
		deviceString := split2[i][20+ipIndexRange[1] : strings.Index(split2[i], "(")]
		fmt.Println("device: ", deviceString)
		productIds := strings.Split(split2[i], "(")[1]
		fmt.Println("product_ids: ", strings.Split(productIds[0:len(productIds)-1], ","))
	}

	// fmt.Println(submatchall)
	// fmt.Println(split1[1])
	// fmt.Println(len(split2))
}
