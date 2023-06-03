package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type url struct {
	Params params
}

// params構造体を引数にして、URLを生成するメソッド
func (u *url) generate() string {
	return "https://api.openweathermap.org/data/2.5/weather?q=" + u.Params.City + "&appid=" + u.Params.APIKey + "&units=" + u.Params.Units
}

type params struct {
	City   string
	APIKey string
	Units  string
}

type weather struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	// params構造体を初期化
	p := params{
		City:   "Sapporo",
		APIKey: "cbefcca05496b015533c61d9740fcc28",
		Units:  "metric",
	}
	// url構造体を初期化
	u := url{
		Params: p,
	}
	url := u.generate()

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
