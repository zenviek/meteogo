# meteogo

A simple library for working with the Open Meteo API (not all functionality is implemented)

# Import
```sh
go get -u github.com/zenviek/meteogo
```

# Usage example
```go
package main

import (
	"fmt"

	"github.com/zenviek/meteogo"
	"github.com/zenviek/meteogo/current"
	"github.com/zenviek/meteogo/hourly"
	"github.com/zenviek/meteogo/wmc"
)

func main() {
	meteo := meteogo.NewMeteo()
	resp, err := meteo.SendRequest(
		meteogo.NewMeteoRequest(
			0.0,
			0.0,
			meteogo.Auto,
		).SetCurrent(
			current.AllCurrents()...,
		).SetHourly(
			hourly.WeatherCode,
			hourly.Temperature2M,
			hourly.RelativeHumidity2M,
			hourly.ApparentTemperature,
		),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
	enDesc, ok := wmc.GetWMCDescription(resp.Current.WeatherCode, wmc.English)
	fmt.Println(enDesc, ok)
}
```