package api

import (
	"context"
	"fmt"
	"ticketing_app/proto-gen/ticket"
)

// GetReceiptDetails retrieves receipt details for a given user.
func (s *TrainService) GetReceiptDetails(ctx context.Context, req *ticket.ReceiptDetailsRequest) (*ticket.Receipt, error) {

	userDetails, found := s.receiptMap[req.ReceiptId]
	if !found {
		return nil, fmt.Errorf("receipt with ID '%s' not found", req.ReceiptId)
	}

	receipt := &ticket.Receipt{
		From:      userDetails.From,
		To:        userDetails.To,
		User:      userDetails.User,
		ReceiptId: userDetails.ReceiptId,
		SeatNo:    userDetails.SeatNo,
		Section:   userDetails.Section,
		PricePaid: userDetails.PricePaid,
	}
	return receipt, nil
}
