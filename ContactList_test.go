package main

import (
	"testing"
	"time"
)

var (
	Persons = []Contact{
		Contact{
			ID:        0,
			Name:      "Clown",
			Phone:     "+998937853311",
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "clown@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
		Contact{
			ID:        1,
			Name:      "Kale",
			Phone:     "+998937853311",
			Firstname: "Archi",
			Lastname:  "Doe",
			Email:     "clown@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
		Contact{
			ID:        2,
			Name:      "Charli",
			Phone:     "+009813239120",
			Firstname: "Charles",
			Lastname:  "Dickens",
			Email:     "IG@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
	}
	TestsForContactUpdate = []Contact{
		Contact{
			ID:        0,
			Name:      "Wick",
			Phone:     "+998937853311",
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "clown@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
		Contact{
			ID:        1,
			Name:      "Ocean",
			Phone:     "+998937853311",
			Firstname: "Archi",
			Lastname:  "Doe",
			Email:     "clown@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
		Contact{
			ID:        2,
			Name:      "Charli",
			Phone:     "+009813239120",
			Firstname: "Charles",
			Lastname:  "Dickens",
			Email:     "IG@gmail.com",
			CreatedAt: time.Now().Format(time.UnixDate),
		},
	}

	p = NewContactManager()
)

func TestCreate(t *testing.T) {
	for _, ppl := range Persons {
		p.Create(&ppl)
		Passed := false
		for _, val := range p.contacts {
			if val.ID == ppl.ID {
				Passed = true
			}
		}
		if !Passed {
			t.Error("Failed to Create object: ", ppl)
		}
	}
}

func TestUpdate(t *testing.T) {
	TestCreate(t)
	for _, ppl := range TestsForContactUpdate {
		if !(*p.Update(&ppl) == ppl) {
			t.Error("Failed to Update obeject: ", ppl)
		}
	}
}
func TestDelete(t *testing.T) {
	TestCreate(t)
	deletedContact := p.Delete(1)
	if deletedContact.ID != 1 {
		t.Error("Failed to Delete contact with id: 1")
	}
}

func TestGet(t *testing.T) {
	TestCreate(t)
	contact := p.Get(1)
	if contact.ID != 1 {
		t.Error("Failed to Get contact with id: 1")
	}
}
func TestGetAll(t *testing.T) {
	TestCreate(t)
	if p.GetAll() != &p.contacts {
		t.Error("failed to get the list of contacts")
	}
}
