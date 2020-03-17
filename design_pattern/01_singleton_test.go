package design_pattern

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestSingleton01(t *testing.T) {
	singleton := getSingleton01()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			tmp := getSingleton01()
			tmp.Name = strconv.Itoa(i)
			p1 := fmt.Sprintf("%p", singleton)
			p2 := fmt.Sprintf("%p", tmp)
			if p1 != p2 {
				t.Error("不是单例")
			}
		}()
	}
	wg.Wait()
}

func TestSingleton02(t *testing.T) {
	singleton := getSingleton02()
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			tmp := getSingleton02()
			tmp.Name = strconv.Itoa(i)
			p1 := fmt.Sprintf("%p", singleton)
			p2 := fmt.Sprintf("%p", tmp)
			if p1 != p2 {
				t.Error("不是单例")
			}
		}()
	}
	wg.Wait()
}

func TestSingleton03(t *testing.T) {
	singleton := getSingleton03()
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			tmp := getSingleton03()
			tmp.Name = strconv.Itoa(i)
			p1 := fmt.Sprintf("%p", singleton)
			p2 := fmt.Sprintf("%p", tmp)
			if p1 != p2 {
				t.Error("不是单例")
			}
		}()
	}
	wg.Wait()
}
func TestSingleton04(t *testing.T) {
	singleton := getSingleton04()
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			tmp := getSingleton04()
			tmp.Name = strconv.Itoa(i)
			p1 := fmt.Sprintf("%p", singleton)
			p2 := fmt.Sprintf("%p", tmp)
			if p1 != p2 {
				t.Error("不是单例")
			}
		}()
	}
	wg.Wait()
}
func TestSingleton05(t *testing.T) {
	singleton := getSingleton05()
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			tmp := getSingleton05()
			tmp.Name = strconv.Itoa(i)
			p1 := fmt.Sprintf("%p", singleton)
			p2 := fmt.Sprintf("%p", tmp)
			if p1 != p2 {
				t.Error("不是单例")
			}
		}()
	}
	wg.Wait()
}
