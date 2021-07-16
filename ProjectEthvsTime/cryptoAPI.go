package main

import (
	"encoding/json"
	"github.com/wcharczuk/go-chart"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)
type EthPricePerDay struct {
	Data []struct {
		PriceUsd string    `json:"priceUsd"`
		Time     int64     `json:"time"`
		Date     time.Time `json:"date"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

var (
	url = "https://api.coincap.io/v2/assets/ethereum/history?interval=d1"
	method = "GET"
	x = []time.Time{}
	y = []float64{}
)
func main(){
	x,y = dataSource()
	//fmt.Println(x)
	//fmt.Println(y)
	http.HandleFunc("/",chartServer)
	http.ListenAndServe(":8080",nil)
}

func chartServer(w http.ResponseWriter,r *http.Request){
	chartObj := chart.Chart{
		Title: "Eth($) vs Time",
		XAxis: chart.XAxis{
			Name: "Time",
		},
		YAxis: chart.YAxis{
			Name: "Eth($)",
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: x,
				YValues: y,
			},
			},
		}
	err := chartObj.Render(chart.PNG,w)
	if err != nil {
		log.Fatalln("unable to render graph")
	}
}


func dataSource()(x []time.Time,y []float64){
	res,err := http.Get(url)
	if err != nil {
		log.Fatalln("unable to make api request ", err)
	}
	res_data,_:=ioutil.ReadAll(res.Body)
	var eth EthPricePerDay
	json.Unmarshal(res_data,&eth)
	for i:=0;i<len(eth.Data);i++{
		date := eth.Data[i].Date
		price,_ := strconv.ParseFloat(string([]rune(eth.Data[i].PriceUsd)[0:6]),64)
		x,y  = append(x,date),append(y,price)
	}
	return
}
