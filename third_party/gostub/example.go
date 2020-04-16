package gostub

import (
	"os"
)

// 全局变量
var globalVar = 100

// 函数 -> 先转为函数变量
func GetGlobalVar() int {
	return globalVar
}

var funcGetVar = GetGlobalVar

// 过程 -> 一般用于回收资源
func DoMain(v int) {
	globalVar = v
}

var funcDoMain = DoMain

// 无法重构的函数
var osHostName = os.Hostname
