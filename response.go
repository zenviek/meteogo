package meteogo

import (
	"github.com/zenviek/meteogo/current"
	"github.com/zenviek/meteogo/daily"
	"github.com/zenviek/meteogo/hourly"
)

type MeteoResponse struct {
	Latitude             float32        `json:"latitude"`
	Longitude            float32        `json:"longitude"`
	GenerationMS         float32        `json:"generation_ms"`
	UTCOffsetSeconds     int            `json:"utc_offset_seconds"`
	Timezone             Timezone       `json:"timezone"`
	TimezoneAbbreviation string         `json:"timezone_abbreviation"`
	Elevation            float32        `json:"elevation"`
	CurrentUnits         *current.Units `json:"current_units"`
	Current              *current.Data  `json:"current"`
	HourlyUnits          *hourly.Units  `json:"hourly_units"`
	Hourly               *hourly.Data   `json:"hourly"`
	DailyUnits           *daily.Units   `json:"daily_units"`
	Daily                *daily.Data    `json:"daily"`
}
