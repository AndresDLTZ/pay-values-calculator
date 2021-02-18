package pay_value

import "github.com/andresdltz/pay-values-calculator/src/pkg/line_name_schedule"

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