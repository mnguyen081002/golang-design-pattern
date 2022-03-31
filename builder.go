package main

import "fmt"

type Show interface {
	showTotal()
}

type HomeBuilder interface {
	buildRoof(string) HomeBuilder
	buildBedRoom(string) HomeBuilder
	buildKitchen(string) HomeBuilder
	buildLivingRoom(string) HomeBuilder
	buildGarage(string) HomeBuilder
	buildHome() Home
}

type Home struct {
	total      string
	roof       string
	bedRoom    string
	kitchen    string
	livingRoom string
	garage     string
}

func (h Home) buildRoof(s string) HomeBuilder {
	h.roof = s
	return h
}

func (h Home) buildBedRoom(s string) HomeBuilder {
	h.bedRoom = s
	return h
}

func (h Home) buildKitchen(s string) HomeBuilder {
	h.kitchen = s
	return h
}

func (h Home) buildLivingRoom(s string) HomeBuilder {
	h.livingRoom = s
	return h
}

func (h Home) buildGarage(s string) HomeBuilder {
	h.garage = s
	return h
}

func (h Home) buildHome() Home {
	h.total = h.roof + " " + h.bedRoom + " " + h.kitchen + " " + h.livingRoom + " " + h.garage
	return h
}

func (h Home) showTotal() {
	fmt.Println(h.total)
}

func main() {
	home := new(Home).
		buildRoof("roof").
		buildBedRoom("bedRoom").
		buildGarage("garage").
		buildHome()
	home.showTotal()
}
