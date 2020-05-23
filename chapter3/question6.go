package chapter3

import (
	"container/list"
	"time"
)

type animalType int

const (
	animalTypeDog animalType = iota
	animalTypeCat
)

type animal struct {
	name        string
	animalType  animalType
	shelteredAt time.Time
}

// isOlderThan checks if the given animal is older.
func (a *animal) isOlderThan(b *animal) bool {
	return a.shelteredAt.Before(b.shelteredAt)
}

type shelter struct {
	dogs, cats *list.List
}

// newShelter creates a new shelter for cats and dogs.
func newShelter() *shelter {
	return &shelter{
		dogs: list.New(),
		cats: list.New(),
	}
}

// enqueue shelters an animal.
func (s *shelter) enqueue(a *animal) {
	a.shelteredAt = time.Now()
	switch a.animalType {
	case animalTypeCat:
		s.cats.PushBack(a)
	case animalTypeDog:
		s.dogs.PushBack(a)
	}
}

// dequeueAny returns the oldest animal regardless of type and removes it from the shelter.
// Returns nil if there are no more animals in the shelter.
func (s *shelter) dequeueAny() *animal {
	if s.dogs.Len() == 0 {
		return s.dequeueCat()
	}
	if s.cats.Len() == 0 {
		return s.dequeueDog()
	}
	oldestDog := s.dogs.Front().Value.(*animal)
	oldestCat := s.cats.Front().Value.(*animal)
	if oldestDog.isOlderThan(oldestCat) {
		return s.dequeueDog()
	}
	return s.dequeueCat()
}

// dequeueDog returns the oldest dog and removes it from the shelter.
// Returns nil if there are no more dogs left
func (s *shelter) dequeueDog() *animal {
	front := s.dogs.Front()
	if front == nil {
		return nil
	}
	s.dogs.Remove(front)
	return front.Value.(*animal)
}

// dequeueCat returns the oldest cat and removes it from the shelter.
// Returns nil if there are no more cats left
func (s *shelter) dequeueCat() *animal {
	front := s.cats.Front()
	if front == nil {
		return nil
	}
	s.cats.Remove(front)
	return front.Value.(*animal)
}
