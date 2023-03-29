package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct{
	sisi float64
}

func (l lingkaran) jarijari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}




func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("\n========= persegi ========")
	fmt.Println("|| luas : ", bangunDatar.luas() , "\t\t||")
	fmt.Println("|| keliling : ",bangunDatar.keliling(), "\t||")
	fmt.Println("==========================")
	
	bangunDatar = lingkaran{14.0}
	fmt.Println("\n=============== lingkaran ================")
	fmt.Println("|| luas : ",bangunDatar.luas(), "\t\t||")
	fmt.Println("|| keliling : ",bangunDatar.keliling(), "\t||")
	fmt.Println("|| jari-jari :", bangunDatar.(lingkaran).jarijari(), "\t\t\t||")
	fmt.Println("==========================================")


}