package main

import (  "time"
		"fmt"
	    "reflect"
		"encoding/json"
		"log")



func main(){

	currentTime:= time.Now()

	fmt.Println(currentTime)

	fmt.Println(reflect.TypeOf(currentTime))

	god := currentTime.Format("2006-01-02")

	fmt.Println(god)

	fmt.Println(reflect.TypeOf(god))

	type FruitBasket struct {
		Created string
		NextDay string
	}
	
	basket := FruitBasket{
		Created: time.Now().Format("2006-01-02"),
		NextDay: time.Now().AddDate(0,0,1).Format("2006-01-02"),
	}
	
	var jsonData []byte
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(jsonData)
	fmt.Println(string(jsonData))
	

	s := string([]byte{65, 66, 67, 226, 130, 172})
	fmt.Println(s) 

	now := time.Now()
    fmt.Println("\nToday:", now)
     
    after := now.AddDate(0, 0, 1)
    fmt.Println("\nAdd 1 day:", after)

	fmt.Println(after.Format("2006-01-02"))
}