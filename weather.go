package main
import (
	"encoding/xml"
	"io"
	"log"
	"net/http"

)

type Weatherdata struct {
	XMLName xml.Name `xml:"current"`
	Text    string   `xml:",chardata"`
	City    struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Name  string `xml:"name,attr"`
		Coord struct {
			Text string `xml:",chardata"`
			Lon  string `xml:"lon,attr"`
			Lat  string `xml:"lat,attr"`
		} `xml:"coord"`
		Country  string `xml:"country"`
		Timezone string `xml:"timezone"`
		Sun      struct {
			Text string `xml:",chardata"`
			Rise string `xml:"rise,attr"`
			Set  string `xml:"set,attr"`
		} `xml:"sun"`
	} `xml:"city"`
	Temperature struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Min   string `xml:"min,attr"`
		Max   string `xml:"max,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"temperature"`
	FeelsLike struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"feels_like"`
	Humidity struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"humidity"`
	Pressure struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"pressure"`
	Wind struct {
		Text  string `xml:",chardata"`
		Speed struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
			Unit  string `xml:"unit,attr"`
			Name  string `xml:"name,attr"`
		} `xml:"speed"`
		Gusts string `xml:"gusts"`
	} `xml:"wind"`
	Clouds struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Name  string `xml:"name,attr"`
	} `xml:"clouds"`
	Visibility struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"visibility"`
	Precipitation struct {
		Text string `xml:",chardata"`
		Mode string `xml:"mode,attr"`
	} `xml:"precipitation"`
	Weather struct {
		Text   string `xml:",chardata"`
		Number string `xml:"number,attr"`
		Value  string `xml:"value,attr"`
		Icon   string `xml:"icon,attr"`
	} `xml:"weather"`
	Lastupdate struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"lastupdate"`
}

func main() {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=moscow&units=metric&mode=xml&appid=ac528d9a31cb6c0dffe8feccd1497739")
if err != nil {
	log.Fatal(err)
}
byteValue, err := io.ReadAll(resp.Body) //передача по значению

var weather Weatherdata
err = xml.Unmarshal(byteValue, &weather) // & - передача по ссылке
	if err != nil {
		log.Fatal(err)
	}
	println(weather.City.Name)
    println("FeelsLike", weather.FeelsLike.Value)
	println("Clouds:", weather.Clouds.Value)
}

