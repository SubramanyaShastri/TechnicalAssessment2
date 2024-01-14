package api

import (
	"context"
	"fmt"
	ticketingapp "ticketing_app/proto-gen/ticket"
)

// ModifySeat updates a user's seat assignment.
func (s *TrainService) ModifySeat(ctx context.Context, req *ticketingapp.ModifySeatRequest) (*ticketingapp.ModifySeatResponse, error) {
	userDetails, exists := s.receiptMap[req.UserId]
	if !exists {
		return nil, fmt.Errorf("user with ID '%s' not found", req.UserId)
	}

	if err := s.updateUserSeat(userDetails, req.NewSeatNumber); err != nil {
		return nil, err
	}

	return &ticketingapp.ModifySeatResponse{Success: true}, nil
}

func (s *TrainService) updateUserSeat(userDetails *ticketingapp.Receipt, newSeatNumber string) error {

	// Check if the new seat is available before assigning
	sectionMap := s.getSectionMap(userDetails.Section)
	if _, taken := sectionMap[newSeatNumber]; taken {
		return fmt.Errorf("seat %s in section %s is already taken", newSeatNumber, userDetails.Section)
	}

	// Remove user from the old seat
	s.removeFromSectionMap(userDetails.Section, userDetails.SeatNo, userDetails.User.UserId)

	// Add user to the new seat
	s.addToSectionMap(userDetails.Section, newSeatNumber, userDetails.User.UserId)

	userDetails.SeatNo = newSeatNumber

	return nil
}

func (s *TrainService) removeFromSectionMap(section string, seatNo string, userId string) {
	sectionMap := s.getSectionMap(section)
	if users, ok := sectionMap[seatNo]; ok {
		for i, id := range users {
			if id == userId {
				sectionMap[seatNo] = append(users[:i], users[i+1:]...)
				break
			}
		}
	}
}

func (s *TrainService) addToSectionMap(section string, seatNo string, userId string) {
	sectionMap := s.getSectionMap(section)
	sectionMap[seatNo] = append(sectionMap[seatNo], userId)
}
