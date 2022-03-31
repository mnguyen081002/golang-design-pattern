package main

import "fmt"

type IOSChanger struct {
}

type Socket struct {
}

func (ios *IOSChanger) Changing15V() {
	fmt.Println("IOS Changing to 15V")
}

type IOSAdapter struct {
	ios *IOSChanger
}

func (iosAdapter *IOSAdapter) Changer10V() {
	iosAdapter.ios.Changing15V()
}

func (s *Socket) Changer10V() {
}

func main() {
	ios := &IOSAdapter{&IOSChanger{}}
	ios.Changer10V()
}
