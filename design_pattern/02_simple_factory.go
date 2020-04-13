package design_pattern

import "fmt"

type API interface {
	Say()
}

func NewAPI(lang string) API {
	switch lang {
	case "cn":
		return &Chinese{}
	case "en":
		return &English{}
	default:
		return nil
	}
}

type Chinese struct {
}

func (c *Chinese) Say() {
	fmt.Println("你好")
}

type English struct {
}

func (e *English) Say() {
	fmt.Println("Hello")
}
