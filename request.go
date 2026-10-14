package meteogo

import (
	"fmt"

	"github.com/zenviek/meteogo/current"
	"github.com/zenviek/meteogo/daily"
	"github.com/zenviek/meteogo/hourly"
)

type MeteoRequest struct {
	ForecastDays int
	PastDays     int
	Latitude     float64
	Longitude    float64
	Timezone     Timezone
	Daily        []daily.Daily
	Hourly       []hourly.Hourly
	Current      []current.Current
}

func NewMeteoRequestWithDays(latitude, longitude float64, timezone Timezone, forecast, past int) *MeteoRequest {
	return &MeteoRequest{
		ForecastDays: forecast,
		PastDays:     past,
		Latitude:     latitude,
		Longitude:    longitude,
		Timezone:     timezone,
	}
}

func NewMeteoRequest(latitude, longitude float64, timezone Timezone) *MeteoRequest {
	return NewMeteoRequestWithDays(latitude, longitude, timezone, 1, 0)
}

func (mr *MeteoRequest) SetDaily(dailies ...daily.Daily) *MeteoRequest {
	mr.Daily = dailies
	return mr
}

func (mr *MeteoRequest) SetHourly(hourlies ...hourly.Hourly) *MeteoRequest {
	mr.Hourly = hourlies
	return mr
}

func (mr *MeteoRequest) SetCurrent(currents ...current.Current) *MeteoRequest {
	mr.Current = currents
	return mr
}

func (mr *MeteoRequest) Body() string {
	currentString := ""
	if len(mr.Current) > 0 {
		currentString = "&current=" + FromUnits(mr.Current...)
	}
	dailyString := ""
	if len(mr.Daily) > 0 {
		dailyString = "&daily=" + FromUnits(mr.Daily...)
	}
	hourlyString := ""
	if len(mr.Hourly) > 0 {
		hourlyString = "&hourly=" + FromUnits(mr.Hourly...)
	}
	return fmt.Sprintf("latitude=%.4f&longitude=%.4f&timezone=%s&forecast_days=%d&past_days=%d%s%s%s",
		mr.Latitude, mr.Longitude, mr.Timezone, mr.ForecastDays, mr.PastDays, currentString, dailyString, hourlyString)
}
