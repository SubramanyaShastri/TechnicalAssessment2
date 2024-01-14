package api

import (
	"ticketing_app/proto-gen/ticket"
)

// TrainService represents the train service.
type TrainService struct {
	ticket.UnimplementedTrainTicketServiceServer
	receiptMap map[string]*ticket.Receipt
	sectionA   map[string][]string
	sectionB   map[string][]string
}

// NewTrainService creates a new instance of TrainService.
func NewTrainService() *TrainService {
	return &TrainService{
		receiptMap: make(map[string]*ticket.Receipt),
		sectionA:   make(map[string][]string),
		sectionB:   make(map[string][]string),
	}
}
