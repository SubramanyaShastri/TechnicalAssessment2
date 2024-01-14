package api

import (
	"context"
	"fmt"
	ticketingapp "ticketing_app/proto-gen/ticket"
)

// GetUsersBySection retrieves user seat assignments for a specified section.
func (s *TrainService) GetUsersBySection(ctx context.Context, req *ticketingapp.SectionRequest) (*ticketingapp.UsersBySectionResponse, error) {
	var sectionMap map[string][]string
	switch req.Section {
	case "A":
		sectionMap = s.sectionA
	case "B":
		sectionMap = s.sectionB
	default:
		return nil, fmt.Errorf("invalid section: %v", req.Section)
	}

	userSeatMappings := make([]*ticketingapp.UserSeatMapping, 0)
	for seatNo, users := range sectionMap {
		for _, userId := range users {
			if _, ok := s.receiptMap[userId]; ok {
				userSeatMappings = append(userSeatMappings, &ticketingapp.UserSeatMapping{
					UserId:     userId,
					SeatNumber: seatNo,
				})
			}
		}
	}

	return &ticketingapp.UsersBySectionResponse{UserSeatMapping: userSeatMappings}, nil
}
