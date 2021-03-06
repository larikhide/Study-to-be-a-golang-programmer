package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//MMWEATHER преобразованный xml с помощью сайта https://www.onlinetool.io/xmltogo/
type MMWEATHER struct {
	XMLName xml.Name `xml:"MMWEATHER"`
	Text    string   `xml:",chardata"`
	REPORT  struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		TOWN struct {
			Text      string `xml:",chardata"`
			Index     string `xml:"index,attr"`
			Sname     string `xml:"sname,attr"`
			Latitude  string `xml:"latitude,attr"`
			Longitude string `xml:"longitude,attr"`
			FORECAST  []struct {
				Text      string `xml:",chardata"`
				Day       string `xml:"day,attr"`
				Month     string `xml:"month,attr"`
				Year      string `xml:"year,attr"`
				Hour      string `xml:"hour,attr"`
				Tod       string `xml:"tod,attr"`
				Predict   string `xml:"predict,attr"`
				Weekday   string `xml:"weekday,attr"`
				PHENOMENA struct {
					Text          string `xml:",chardata"`
					Cloudiness    string `xml:"cloudiness,attr"`
					Precipitation string `xml:"precipitation,attr"`
					Rpower        string `xml:"rpower,attr"`
					Spower        string `xml:"spower,attr"`
				} `xml:"PHENOMENA"`
				PRESSURE struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"PRESSURE"`
				TEMPERATURE struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"TEMPERATURE"`
				WIND struct {
					Text      string `xml:",chardata"`
					Min       string `xml:"min,attr"`
					Max       string `xml:"max,attr"`
					Direction string `xml:"direction,attr"`
				} `xml:"WIND"`
				RELWET struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"RELWET"`
				HEAT struct {
					Text string `xml:",chardata"`
					Min  string `xml:"min,attr"`
					Max  string `xml:"max,attr"`
				} `xml:"HEAT"`
			} `xml:"FORECAST"`
		} `xml:"TOWN"`
	} `xml:"REPORT"`
}

func main() {
	//запрос на xml в сеть
	responce, err := http.Get("https://xml.meteoservice.ru/export/gismeteo/point/32277.xml")
	//обработка ошибки
	if err != nil {
		log.Fatal(err)
	}
	//закрыть поток как только процедура закончится
	defer responce.Body.Close()

	//чтение из потока в массив байт byteValue
	byteValue, err := ioutil.ReadAll(responce.Body)

	if err != nil {
		log.Fatal(err)
	}

	var weather MMWEATHER
	//распаковываем полученный массив байт, передаем weather по ссылке, чтобы unmarshall поменял значение переменной weather
	err = xml.Unmarshal(byteValue, &weather)

	if err != nil {
		log.Fatal(err)
	}

	//раскодировать имя города
	//TODO: url.PathUnescape возвращает строку и ошибку. Узнать как не выводить ошибку
	fmt.Println(url.PathUnescape(weather.REPORT.TOWN.Sname))

	// short name of area
	forecast := weather.REPORT.TOWN.FORECAST
	for a := range forecast {
		fmt.Printf("%s-%s-%s at %s o'clock max temperature is %s celcium degree\n", forecast[a].Day, forecast[a].Month, forecast[a].Year, forecast[a].Hour, forecast[a].TEMPERATURE.Max)
	}
}
