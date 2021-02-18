package pay_value

import (
	"fmt"
	"github.com/andresdltz/pay-values-calculator/src/pkg/line_name_schedule"
	"github.com/andresdltz/pay-values-calculator/src/pkg/payment_rate"
	"github.com/andresdltz/pay-values-calculator/src/pkg/tools"
)

var acmeDailyPaymentRates *[]payment_rate.DaysScheduleRate

func init() {
	acmeDailyPaymentRates = payment_rate.GetAcmeDailyPaymentRates()
}

// CalculatePayValue will retrieve a lineCh channel, line, doneCh channel to calculate pay value
func CalculatePayValue(lineCh chan *line_name_schedule.Line, doneCh chan *bool) {

	// loop for create a goroutine for each line
	for {
		_, more := <- lineCh
		if more {
			// todo: goroutine to calculate pay value
		}
		close(doneCh)
		return
	}
}

// calculateTotalPayValue calculates pay value of all schedules worked
func calculateTotalPayValue(line *line_name_schedule.Line) float64 {
	payValue := float64(0)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s %s \nAdvertencia: %s\n\n", line.Name, line.Schedule, r)
		}
	}()

	// loop for iterates schedules worked
	for _, dayAndRangeHour := range line.Schedule {
		// loop for iterates ACME daily payment rate
		for _, acmeDailyScheduleRate := range *acmeDailyPaymentRates {
			if tools.IsElementInArray(acmeDailyScheduleRate.Days, dayAndRangeHour[:2]) {
				payValue += calculatePartialPayValue(
					acmeDailyScheduleRate.ScheduleRates,
					dayAndRangeHour[2:],
					len(acmeDailyScheduleRate.ScheduleRates) - 1)
				break
			}
		}
	}
	fmt.Printf("%s %s\nEl importe a pagar a %s es: %.2f USD\n\n", line.Name, line.Schedule, line.Name, payValue)
	return payValue
}

// calculatePartialPayValue calculates pay value of one schedule worked
func calculatePartialPayValue(acmeSR []payment_rate.ScheduleRate, hrRange string, counter int) float64 {
	if counter == -1 {
		panic(fmt.Sprintf("No se puede calcular el horario a pagar, horario %s fuera de referencia", hrRange))
	}
	duration := tools.GetDurationOfValidHourRange(hrRange, &acmeSR[0])
	if duration == -1 { return calculatePartialPayValue(acmeSR[1:], hrRange, counter - 1) }
	return duration * acmeSR[0].Rate
}