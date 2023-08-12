package main

import (
	"fmt"
	"strings"
)


func testPrint() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	
	go func(){
		i := 1
		for {
			select {
				case <-ch1:
					fmt.Print(i)
					i++
					fmt.Print(i)
					i++
					ch2 <- true
			}
		}
	}()
	
	go func(){
		j := 'A'
		for {
			select {
				case <-ch2:
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					if j > 'Z' {
						ch3 <- true
						return
					}
					ch1 <- true
			}
		}
	}()
	
	ch1<-true
	<-ch3
}

func isRepeat(s string) bool {

	if len(s) > 3000 {
		return false
	}
	
	i := 0
	for key, val := range s {
		if val > 256 {
			return false
		}
		
		for kk, vv := range s {
			if vv == val && kk != key {
				fmt.Println(string(vv))
				i++
			}
		}
	}		
	
	if i > 0 {
		return true
	}
	
	return false
}

func reverse(s string) string {
	str := []int32(s)
	
	length := len(str)
	
	if length > 5000 {
		return ""
	}
	
	for k,_ := range str{
		if k > length/2 {
			break
		}
		
		str[k], str[length - k - 1] = str[length - k - 1], str[k]
	}
	
	return string(str)
	
}

func check(s1, s2 string) bool{
	ss1 := []rune(s1)
	ss2 := []rune(s2)
	
	if len(ss1) != len(ss2) {
		return false
	}
	
	if len(ss1) > 5000 || len(ss2) > 5000 {
		return false
	}
	
	for _, val := range ss1 {
		if strings.Count(s1, string(val)) != strings.Count(s2, string(val)) {
			return false
		}
	}
	
	return true
}

//冒泡排序
func sortStr(s []rune) ([]rune, bool) {
	length := len(s) 
	isEnd := true
	
	for k,_ := range s {
		if length -1 > k && s[k] > s[k + 1] {
			s[k], s[k + 1] = s[k + 1], s[k]
			isEnd = false
		}
	}
	
	if isEnd {
		return s, isEnd
	} else {
		return sortStr(s)
	}
}	

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func bigIntAdd(str1, str2 string) {
	len1 := len(str1)
	len2 := len(str2)
	length := 0
	if len1 > len2 {
		length = len1
	} else {
		length = len2
	}
	
	for  i:=0; i<length; i++{
		m := ""
		n := ""
		if len1 > 0 {
			m = str1[len1-1:len1]
			len1--
			fmt.Println("str1:"+m)
		}
		
		if len2 > 0 {
			n = str2[len2-1:len2]
			len2--
			fmt.Println("str2:"+n)
		}
		
	}
		
}

func main() {
	//s := "123456asdgeABFTGDaB"
	//if isRepeat(s) {
	//	fmt.Println("重复")
	//} else {
	//	fmt.Println("不重复")
	//}
	
	//str := "我爱你中国123456hello"
	//s := reverse(str)
	//fmt.Println(s)
	
	/*
	str1 := "我爱你中国123456hello11"
	str2 := "22123456hello我爱你中国"
	if check(str1, str2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	
	
	str1 := "我爱你中国123456来来来11"
	str2 := "1我爱你中国123456hello1"
	
	s1 := []rune(str1)
	s2 := []rune(str2)
	
	s1, _ = sortStr(s1)
	s2, _ = sortStr(s2)
	
	fmt.Println(string(s1))
	fmt.Println(string(s2))
	
	if string(s1) == string(s2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	*/
	
	//fmt.Println("return:", b())
	//fmt.Println("return111:", *(c()))
	
	a := "9876543111111111111111111111111111111"
	b := "2345678922222222222222222222222222222"
	bigIntAdd(a, b)
	
}


