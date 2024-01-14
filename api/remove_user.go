package api

import (
	"context"
	"errors"
	"ticketing_app/proto-gen/ticket"
)

// RemoveUser removes a user from the system.
func (s *TrainService) RemoveUser(ctx context.Context, req *ticket.RemoveUserRequest) (*ticket.RemoveUserResponse, error) {
	userDetails, exists := s.receiptMap[req.UserId]
	if !exists {
		return nil, errors.New("user not found")
	}

	if err := s.removeFromSection(userDetails); err != nil {
		return nil, err
	}

	delete(s.receiptMap, req.UserId)

	return &ticket.RemoveUserResponse{Success: true}, nil
}

func (s *TrainService) removeFromSection(userDetails *ticket.Receipt) error {
	sectionMap := s.getSectionMap(userDetails.Section)
	if sectionMap == nil {
		return errors.New("invalid section")
	}

	if _, ok := sectionMap[userDetails.SeatNo]; ok {
		delete(sectionMap, userDetails.SeatNo)
		return nil
	}

	return errors.New("user seat not found in section")
}
