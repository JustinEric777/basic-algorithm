package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type EchoDomain struct {
	Code int             `json:"code"`
	Data map[string]Data `json:"data"`
}

type Data struct {
	Domain string  `json:"domain"`
	Data   [][]int `json:"data"`
}

func main() {
	file, err := ioutil.ReadFile("/Users/iknow/Works/Go/interview/basic-algorithm/leetcode/test/test2.json")
	if err != nil {
		fmt.Println(err)
	}

	echoDomain := &EchoDomain{}
	err = json.Unmarshal(file, echoDomain)
	if err != nil {
		fmt.Println(err)
	}

	filePath := "/Users/iknow/Works/Go/interview/basic-algorithm/leetcode/test/test.txt"
	file1, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0775)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file1.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file1)

	//fmt.Println(echoDomain.Data)
	for _, val := range echoDomain.Data {
		if len(val.Data) > 0 {
			for _, v := range val.Data {
				line := fmt.Sprintf("%s,%d,%d\n", val.Domain, v[0], v[1])
				write.WriteString(line)
			}
		} else {
			line := fmt.Sprintf("%s,%d,%d\n", val.Domain, 0, 0)
			write.WriteString(line)
		}
	}
	write.Flush()
}
