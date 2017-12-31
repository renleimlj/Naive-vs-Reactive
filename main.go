package main

import (
	"fmt"
	"time"
	"Naive-vs-Reactive/agency"
)

func synchronous() [10]agency.Infomation {
	start := time.Now()
	customer := agency.GetCustomers()
	destinations := agency.GetDestinations(customer)
	var infos [10]agency.Infomation
	for index, dest := range destinations {
		q := agency.GetQuoting(dest)
		w := agency.GetWeather(dest)
		infos[index] = agency.Infomation{W : w,D : dest,Q : q}
	}
	fmt.Println(time.Since(start))
	return infos
}

func asynchronous() [10]agency.Infomation {
	start := time.Now()
	customer := agency.GetCustomers()
	destinations := agency.GetDestinations(customer)
	var infos [10]agency.Infomation
	quotes := [10]chan agency.Quoting{}
	weathers := [10]chan agency.Weather{}

	for i, _ := range quotes {
		quotes[i] = make(chan agency.Quoting)
	}

	for i, _ := range weathers {
		weathers[i] = make(chan agency.Weather)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		go func() {
			quotes[i] <- agency.GetQuoting(d)
		}()

		go func() {
			weathers[i] <- agency.GetWeather(d)
		}()
	}

	for index, dest := range destinations {
		infos[index] = agency.Infomation{W:<-weathers[index],D:dest, Q:<-quotes[index]}
	}

	fmt.Println(time.Since(start))
	return infos
}

func main() {
	fmt.Println("Using synchronous way")
	synchronous()
	fmt.Println("Using asynchronous way")
	asynchronous()
}
