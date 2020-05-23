package chapter3

import (
	"testing"
)

func Test_newShelter(t *testing.T) {
	sh := newShelter()

	sh.enqueue(&animal{animalType: animalTypeDog, name: "doge1"})
	sh.enqueue(&animal{animalType: animalTypeDog, name: "doge2"})
	sh.enqueue(&animal{animalType: animalTypeCat, name: "cat1"})
	sh.enqueue(&animal{animalType: animalTypeDog, name: "doge3"})

	if got := sh.dequeueAny(); got.name != "doge1" {
		t.Errorf("dequeueAny() want dog with name %s got %s", "doge1", got.name)
	}

	if got := sh.dequeueCat(); got.name != "cat1" {
		t.Errorf("dequeueCat() want cat with name %s got %s", "cat1", got.name)
	}

	if got := sh.dequeueCat(); got != nil {
		t.Errorf("dequeueCat() want nil, got cat with name %s", got.name)
	}

	if got := sh.dequeueDog(); got.name != "doge2" {
		t.Errorf("dequeueDog() want dog with name %s got %s", "doge2", got.name)
	}

	if got := sh.dequeueDog(); got.name != "doge3" {
		t.Errorf("dequeueDog() want dog with name %s got %s", "doge3", got.name)
	}

	if got := sh.dequeueDog(); got != nil {
		t.Errorf("dequeueDog() want nil, got dog with name %s", got.name)
	}

	if got := sh.dequeueAny(); got != nil {
		t.Errorf("dequeueAny() want nil, got animal with name %s", got.name)
	}

	sh.enqueue(&animal{name: "Garfield", animalType: animalTypeCat})
	if got := sh.dequeueAny(); got.name != "Garfield" {
		t.Errorf("dequeueAny() want animal with name %s got %s", "Garfield", got.name)
	}

	sh.enqueue(&animal{name: "Snoopy", animalType: animalTypeDog})
	if got := sh.dequeueAny(); got.name != "Snoopy" {
		t.Errorf("dequeueAny() want animal with name %s got %s", "Garfield", got.name)
	}

	sh.enqueue(&animal{name: "Felix", animalType: animalTypeCat})
	sh.enqueue(&animal{name: "Goofy", animalType: animalTypeDog})
	if got := sh.dequeueAny(); got.name != "Felix" {
		t.Errorf("dequeueAny() want animal with name %s got %s", "Garfield", got.name)
	}

}
