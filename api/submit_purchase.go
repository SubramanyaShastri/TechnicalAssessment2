package api

import (
	"context"
	"errors"
	"strconv"

	ticketingapp "ticketing_app/proto-gen/ticket"
)

// SubmitPurchase processes a ticket purchase request.
func (s *TrainService) SubmitPurchase(ctx context.Context, req *ticketingapp.PurchaseRequest) (*ticketingapp.Receipt, error) {
	if err := validatePurchaseTicketRequest(req); err != nil {
		return nil, err
	}

	receiptID := strconv.Itoa(len(s.receiptMap) + 1)
	userID := strconv.Itoa(len(s.receiptMap) + 1)

	// Allocate seat using the AllocateSeat function
	allocateSeatReq := &ticketingapp.SeatAllocationRequest{
		UserId:  userID,
		SeatNo:  req.SeatNo,
		Section: req.Section,
	}
	allocateSeatResp, err := s.AllocateSeat(ctx, allocateSeatReq)
	if err != nil {
		return nil, err
	}

	// Apply discount if discount_code is provided
	priceDeduct := getDiscountPrice(req.DiscountCode)

	receipt := &ticketingapp.Receipt{
		ReceiptId: receiptID,
		From:      req.From,
		To:        req.To,
		User: &ticketingapp.User{
			UserId:    userID,
			FirstName: req.User.FirstName,
			LastName:  req.User.LastName,
			Email:     req.User.Email,
		},
		SeatNo:    allocateSeatResp.SeatNumber,
		Section:   allocateSeatResp.Section,
		PricePaid: req.PricePaid - priceDeduct,
	}

	s.receiptMap[userID] = receipt

	return receipt, nil
}

func validatePurchaseTicketRequest(req *ticketingapp.PurchaseRequest) error {
	if req.From == "" || req.To == "" || req.User.FirstName == "" || req.User.LastName == "" ||
		req.User.Email == "" || req.Section == "" || req.SeatNo == "" {
		return errors.New("all fields are required")
	}
	return nil
}
