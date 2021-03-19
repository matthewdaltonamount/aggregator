package aggregator

import (
	"bufio"
	"context"
	"github.com/matthewdaltonamount/collector/pkg/utils"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseLogService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	agg, _ := ParseLogController(ctx)
	utils.RenderOr500(w, r, agg)
}

func ParseLogController(ctx context.Context) (Agg, error) {

	f, err := os.Open(ctx.Value("log").(string))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	agg := new(Agg)
	var currentProductName = ""
	var currentProductType = ""
	var currentProductData = new([]data)

	for scanner.Scan() {

		split := strings.Split(scanner.Text(), " ")

		if split[0] == "reference" {
			agg.ReferenceTemp, err = strconv.ParseFloat(split[1], 64)
			if err != nil {
				return Agg{}, err
			}

			agg.ReferenceHumidity, err = strconv.ParseFloat(split[2], 64)
			if err != nil {
				return Agg{}, err
			}
		} else if checkTimeStamp(split[0]) {
			p, err := strconv.ParseFloat(split[1], 64)
			if err != nil {
				return Agg{}, err
			}
			t, err := parseShortenedTimestamp(split[0])
			if err != nil {
				return Agg{}, err
			}
			d := &data{
				Time: t,
				Data: p,
			}
			if len(*currentProductData) == 0 {
				agg.WindowStartTime = t
			}
			agg.WindowEndTime = t
			*currentProductData = append(*currentProductData, *d)
		} else {
			if currentProductType != "" {
				if currentProductType == "thermometer" {
					therm := &Thermometer{
						Name:      currentProductName,
						Precision: calcTherm(agg.ReferenceTemp, *currentProductData),
					}

					agg.Thermometers = append(agg.Thermometers, *therm)
				}

				if currentProductType == "humidity" {
					h := &Hsensor{
						Name:    currentProductName,
						Discard: calcHsensor(agg.ReferenceHumidity, *currentProductData),
					}

					agg.Hsensors = append(agg.Hsensors, *h)
				}
			}

			currentProductData = new([]data)
			currentProductName = split[1]
			currentProductType = split[0]
		}

	}

	if currentProductType == "thermometer" {
		therm := &Thermometer{
			Name:      currentProductName,
			Precision: calcTherm(agg.ReferenceTemp, *currentProductData),
		}

		agg.Thermometers = append(agg.Thermometers, *therm)
	}

	if currentProductType == "humidity" {
		h := &Hsensor{
			Name:    currentProductName,
			Discard: calcHsensor(agg.ReferenceHumidity, *currentProductData),
		}

		agg.Hsensors = append(agg.Hsensors, *h)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return *agg, nil
}

func calcTherm(ref float64, productData []data) string {

	var sum, mean, sd float64

	for i := range productData {

		sum += productData[i].Data
	}

	mean = sum / float64(len(productData))

	for j := range productData {
		sd += math.Pow(productData[j].Data-mean, 2)
	}
	sd = math.Sqrt(sd / float64(len(productData)))

	meanCheck := math.Abs(mean-ref) < .5

	if meanCheck && sd < 3 {
		return "ultra precise"
	} else if meanCheck && sd < 5 {
		return "very precise"
	} else {
		return "precise"
	}
}

func calcHsensor(ref float64, productData []data) bool {
	for i := range productData {
		if math.Abs(productData[i].Data-ref) > 1 {
			return false
		}
	}
	return true
}

func checkTimeStamp(possibleTimestamp string) bool {
	_, err := parseShortenedTimestamp(possibleTimestamp)
	if err != nil {
		return false
	}
	return true
}

func parseShortenedTimestamp(timestring string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, timestring+":00Z")
	if err != nil {
		return time.Time{}, err
	}
	return t, nil

}
