package Components

import "time"

type History struct {
	Date          time.Time `json:"date"`
	Destination   string    `json:"destination"`
	OperationType string    `json:"operationType"`
	Sum           float64   `json:"sum"`
}

func createHistoryField(destination, operationType string, sum float64) History {
	return History{
		Date:          time.Now(),
		Destination:   destination,
		OperationType: operationType,
		Sum:           sum,
	}
}
