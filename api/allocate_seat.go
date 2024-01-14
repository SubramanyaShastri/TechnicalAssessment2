package api

import (
	"context"
	"errors"
	ticketingapp "ticketing_app/proto-gen/ticket"
)

// AllocateSeat is a function to allocate a seat.
func (s *TrainService) AllocateSeat(ctx context.Context, req *ticketingapp.SeatAllocationRequest) (*ticketingapp.SeatAllocationResponse, error) {

	err := s.assignSeat(req.UserId, req.Section, req.SeatNo)
	if err != nil {
		return nil, err
	}

	return &ticketingapp.SeatAllocationResponse{
		SeatNumber: req.SeatNo,
		Section:    req.Section,
	}, nil
}

// assignSeat is a function to assign a seat to a user in a specific section.
func (s *TrainService) assignSeat(userID string, section string, seatNo string) error {
	sectionMap := s.getSectionMap(section)
	if sectionMap == nil {
		return errors.New("invalid section")
	}

	if seats, exists := sectionMap[seatNo]; exists && len(seats) > 0 {
		return errors.New("seat already taken")
	}

	sectionMap[seatNo] = []string{userID}
	return nil
}
