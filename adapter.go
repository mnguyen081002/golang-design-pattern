package main

import "fmt"

type ChangerMobile interface {
	Changing10V()
}

type IOSChanger struct {
}

type AndroidChanger struct {
}

func (ios *IOSChanger) Changing15V() {
	fmt.Println("IOS Changing to 15V")
}

type IOSAdapter struct {
	ios *IOSChanger
}

func (iosAdapter *IOSAdapter) Changing10V() {
	iosAdapter.ios.Changing15V()
}

func (s *AndroidChanger) Changing10V() {
	fmt.Println("Android Changing to 10V")
}

func main() {
	ios := &IOSAdapter{&IOSChanger{}}
	android := &AndroidChanger{}

	var changer []ChangerMobile
	changer = append(changer, ios)
	changer = append(changer, android)

	for _, v := range changer {
		v.Changing10V()
	}
}
