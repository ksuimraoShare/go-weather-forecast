package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	const apiKey = "cbefcca05496b015533c61d9740fcc28"
	const url = "https://api.openweathermap.org/data/2.5/weather?q=Sapporo&appid=" + apiKey

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var w weather
	err = json.Unmarshal(body, &w)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Temperture in %s : %f\n", w.Name, w.Main.Temp)
}
