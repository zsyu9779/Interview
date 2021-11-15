/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	//pattern.TestDecorator()

	url := "http://localhost:8300/api/v1/snap/up/1421000814075387905?w=1080&h=2028&model=MI%25208&vendor=Xiaomi&dpi=440&pkg=com.hunt.daily.baitao&v=1&vn=2&vs=&lc=default&lang=zh&os=android&op=default_op&locale=CN&ntt=WIFI&ts=1628048862104&telecom=&adid=&brand=Xiaomi&imei=&imeimd5=&mac=&tk=9le%252FkW4%252FvZsCkt2dX4pikg%253D%253D&vc=46757e44ff15c15d40f52f5b43b71131"
	method := "POST"

	payload := strings.NewReader(`{
    "payOrderId":1441325526890323969,
    "skuId":"1441341704027344897-1441341704228671489_1441341704262225922",
    "type":2
}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("accessToken", "MGTSK1441398422509731842")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}