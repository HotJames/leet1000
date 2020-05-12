package main

import (
	"fmt"
	"reflect"
)

/* for test */
func main() {

	//testRes :=
	//is_palindrome.IsPalindrome(n)
	//is_happy.IsHappy(n)
	//fmt.Println(testRes)
	//fmt.Println(testRes)

	s := []string{"flower", "flow", "flight"}

	for i := 0; i ++ {

	}

	for  _, i := range s{
		for _, j := range []byte(i){
		fmt.Println(j)
		fmt.Println(reflect.TypeOf(j))
	}
	}

}
