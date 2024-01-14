package api

// getSectionMap is a function to retrieve the section map based on the given section.
func (s *TrainService) getSectionMap(section string) map[string][]string {
	switch section {
	case "A":
		return s.sectionA
	case "B":
		return s.sectionB
	default:
		return nil
	}
}
