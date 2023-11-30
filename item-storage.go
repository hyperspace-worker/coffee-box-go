package main

type ItemStorage struct {
	cups int
}

func (s *ItemStorage) areAnyGlasses() bool {
	return s.cups > 0
}

func (s *ItemStorage) getCup() {
	s.cups--
}

func (s *ItemStorage) tryFillUpGlasses(glasses int) bool {
	if s.cups+glasses <= GLASSES_CAPACITY {
		s.cups += glasses
		return true
	}
	return false
}
