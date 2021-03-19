package aggregator

import (
	"encoding/json"
	"net/http"
	"reflect"
	"time"
)

func ByteArraytoAgg(b []byte) (Agg, error) {
	body := Agg{}
	var err error

	err = json.Unmarshal(b, &body)
	return body, err
}

func (s Agg) IsEmpty() bool {
	return reflect.DeepEqual(s, Agg{})
}

type Agg struct {
	WindowStartTime   time.Time `json:"windowStartTime"`
	WindowEndTime     time.Time `json:"windowEndTime"`
	ReferenceTemp     float64   `json:"referenceTemp"`
	ReferenceHumidity float64   `json:"referenceHumidity"`

	Thermometers []Thermometer `json:"Thermometers"`
	Hsensors     []Hsensor     `json:"Hsensors"`
}

func (s Agg) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Thermometer struct {
	Name      string `json:"name"`
	Precision string `json:"Precision"`
}

func (a *Thermometer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Hsensor struct {
	Name    string `json:"name"`
	Discard bool   `json:"Discard"`
}

func (a *Hsensor) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
