package main

import (
	"fmt"
	"log"
	"strconv"
)

/**
 * interface to string
 */
func ToString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case nil:
		return ""
	}

	return fmt.Sprintf("%v", src)
}

/**
 * interface to int
 */
func ToInt(src interface{}) (dst int) {

	var ok bool

	if dst, ok = src.(int); !ok {
		switch src := src.(type) {
		case []byte:
			dst, _ = strconv.Atoi(string(src))
		case string:
			dst, _ = strconv.Atoi(src)
		case nil:
			dst = 0
		case bool:
			if src {
				dst = 1
			} else {
				dst = 0
			}
		default:
			str := fmt.Sprintf("%v", src)
			dst, _ = strconv.Atoi(str)
		}
	}

	return dst
}

/*
	用来检测通过命令行传递的参数是否有问题.
*/
func CheckArgs(args []string) (string, bool) {
	args_len := len(args)
	if args_len != 5 {
		return ARGS_NUM_ERR, false
	}
	return ARGS_OK, true
}
func CheckErr(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			return 0
		}
		log.Fatal(" error info ", err.Error())
		return -1
	}
	return 1
}
