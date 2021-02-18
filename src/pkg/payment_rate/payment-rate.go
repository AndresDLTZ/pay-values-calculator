package payment_rate

import (
	"encoding/json"
	"fmt"
	"github.com/andresdltz/pay-values-calculator/src/pkg/tools"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// ScheduleRate type provide a structure that takes care of getting StartTime, EndTime and Rate
type ScheduleRate struct {
	StartTime string
	EndTime string
	Rate float64
}

// DaysScheduleRate type provide a structure that takes care of ACME daily payment rate
type DaysScheduleRate struct {
	Days []string
	ScheduleRates []ScheduleRate
}

var (
	acmeDailyPaymentRate *[]DaysScheduleRate
	once          		 sync.Once
)

// GetAcmeDailyPaymentRates singleton function get acme daily payment rate data
func GetAcmeDailyPaymentRates() *[]DaysScheduleRate {
	// once.Do ensures that this function is executed only once
	once.Do(func() {
		acmeDailyPaymentRate = getJsonToDailyPaymentRate()
	})
	return acmeDailyPaymentRate
}

// getJsonToDailyPaymentRate parse acme daily payment rate json to DaysScheduleRate type
func getJsonToDailyPaymentRate() *[]DaysScheduleRate {
	scheduleRates := make([]DaysScheduleRate, 2)
	acmeDailyPaymentRateFilePath := "files/acme-daily-payment-rate.json"
	data, err := ioutil.ReadFile(acmeDailyPaymentRateFilePath)
	if err != nil {
		fmt.Printf("Error: %s no existe tal archivo o directorio\n", acmeDailyPaymentRateFilePath)
		fmt.Println("\t1. cree un directorio files\n\t" +
			"2. dentro del directorio cree un archivo con el nombre de acme-daily-paymen-rate.json\n\t" +
			"3. ingrese las tarifas de pago")
		os.Exit(0)
	}
	_ = json.Unmarshal(data, &scheduleRates)
	return &scheduleRates
}

// GetDurationOfValidHourRange will retrieve a hour range and payment rate,
// check what schedule it is and return to the hour difference
func GetDurationOfValidHourRange(hrRange string, sr *ScheduleRate) float64 {
	hrArr := strings.Split(hrRange, "-")
	hr1 := tools.ParseStringToTime(&hrArr[0])
	hr2 := tools.ParseStringToTime(&hrArr[1])

	if !(hr1.Before(tools.ParseStringToTime(&sr.StartTime)) || hr2.After(tools.ParseStringToTime(&sr.EndTime))) {
		return tools.DiffHourTime(&hr1, &hr2)
	}
	return -1
}
