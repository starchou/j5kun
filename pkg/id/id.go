package id

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type IdCard struct {
	Province   string
	Area       string
	City       string
	CardNumber string `json`
	Age        time.Time
	Sex        int //性别，0为女，1为男
}

func NewIDCard(number string) *IdCard {
	return &IdCard{CardNumber: number}
}

func (c *IdCard) Analysis() *IdCard {
	//取省份，地区，区县
	c.Province = list[substr(c.CardNumber, 0, 2)+"0000"]
	c.Area = list[substr(c.CardNumber, 0, 4)+"00"]
	c.City = list[substr(c.CardNumber, 0, 6)]

	//取年龄
	ageCode := substr(c.CardNumber, 6, 8)
	c.Age, _ = time.Parse("20060102", ageCode)
	//取性别
	sexCode := substr(c.CardNumber, 14, 3)
	sexNum, _ := strconv.Atoi(sexCode)
	c.Sex = sexNum % 2
	return c
}

func (c *IdCard) Json() ([]byte, error) {
	return json.Marshal(c)
}

func (c *IdCard) String() string {
	hours := time.Now().Sub(c.Age).Hours()
	sex := "男"
	if c.Sex == 0 {
		sex = "女"
	}
	return fmt.Sprintf("%s,%s,%s 年龄:%d 性别:%s", c.Province, c.Area, c.City, int(hours/(24*365)), sex)
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	if pos > l {
		return ""
	}
	return string(runes[pos:l])
}
