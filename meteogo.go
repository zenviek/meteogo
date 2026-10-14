package meteogo

import (
	"errors"

	"github.com/zenviek/reqtango/v2"
)

var url = "https://api.open-meteo.com/v1/forecast?"

type Meteo struct {
	rb *reqtango.RequestBuilder
}

func NewMeteo() *Meteo {
	return &Meteo{
		rb: reqtango.NewRequestBuilderSimple(),
	}
}

func (m *Meteo) SendRequest(request *MeteoRequest) (*MeteoResponse, error) {
	if request == nil {
		return nil, errors.New("request is nil")
	}
	var response MeteoResponse
	if _, err := m.rb.GetToStruct(url+request.Body(), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
