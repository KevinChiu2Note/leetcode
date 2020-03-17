package design_pattern

import (
	"sync"
)

// 懒汉方式 - 非线程安全。当正在创建时，有线程同时访问ins = nil就会再创建，单例类就会有多个实例了。
type singleton struct {
	Name string
}

var ins01 *singleton

func getSingleton01() *singleton {
	if ins01 == nil {
		ins01 = &singleton{}
	}
	return ins01
}

// 饿汉方式 - 如果singleton创建初始化比较复杂耗时时，加载时间会延长。
var ins02 = &singleton{}

func getSingleton02() *singleton {
	return ins02
}

// 懒汉加锁方式 - 虽然解决并发的问题，但每次加锁是要付出代价的
var (
	ins03  *singleton
	lock03 sync.Mutex
)

func getSingleton03() *singleton {
	lock03.Lock()
	defer lock03.Unlock()
	if ins03 == nil {
		ins03 = &singleton{}
	}
	return ins03
}

// 懒汉初始加锁方式 -
var (
	ins04  *singleton
	lock04 sync.Mutex
)

func getSingleton04() *singleton {
	if ins04 == nil {
		lock04.Lock()
		defer lock04.Unlock()
		if ins04 == nil {
			ins04 = &singleton{}
		}
	}
	return ins04
}

// sync.Once - 核心思想是使用原子计数记录被执行的次数。使用Mutex Lock Unlock锁定被执行函数，防止被重复执行。
var (
	ins05 *singleton
	once  sync.Once
)

func getSingleton05() *singleton {
	once.Do(func() {
		ins05 = &singleton{}
	})
	return ins05
}
