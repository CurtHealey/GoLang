package main

import ("fmt"
"encoding/json"
"log"
"time"
"github.com/jasonlvhit/gocron")

func task(){


	type FruitBasket struct {
		Created string
		NextDay string
	}
	
	basket := FruitBasket{
		Created: time.Now().Format("2006-01-02 15:04:05"),
		NextDay: time.Now().AddDate(0,0,1).Format("2006-01-02 15:04:05"),
	}
	
	var jsonData []byte
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}

func main() {
    s := gocron.NewScheduler()
    s.Every(2).Seconds().Do(task)
    <- s.Start()
}