package main

type Contact struct {
	ID int
	Name,
	Phone,
	Firstname,
	Lastname,
	Email,
	CreatedAt string
}

type ContactManager struct {
	contacts []Contact
}

func CreateContactManager() *ContactManager {
	p := new(ContactManager)
	p.contacts = make([]Contact, 0)
	return p
}

func (c *ContactManager) Create(cont *Contact) *Contact {
	for i := 0; i < len(c.contacts); i++ {
		if c.contacts[i].ID == cont.ID {
			return cont
		}
	}
	c.contacts = append(c.contacts, *cont)
	return &Contact{}
}

func (c *ContactManager) Update(cont *Contact) *Contact {
	for i := 0; i < len(c.contacts); i++ {
		if c.contacts[i].ID == cont.ID {
			c.contacts[i] = *cont
			return cont
		}
	}
	return &Contact{}
}

func (c *ContactManager) Delete(ID int) *Contact {
	for i := 0; i < len(c.contacts); i++ {
		if c.contacts[i].ID == ID {
			out := c.contacts[i]
			c.contacts = append(c.contacts[:i], c.contacts[i+1:]...)
			return &out
		}
	}
	return &Contact{}
}
func (c *ContactManager) Get(ID int) *Contact {
	for i := 0; i < len(c.contacts); i++ {
		if c.contacts[i].ID == ID {
			return &c.contacts[i]
		}
	}
	return &Contact{}
}

func (c *ContactManager) GetAll() *[]Contact {
	return &c.contacts
}
