package models

type Person struct {
	FirstName string
	ID        uint `mandatory`
	LastName  string
}

func (m PersonIDName) GetIDName() PersonIDName {
	return PersonIDName{
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
}

func (m PersonIDPrimary) GetIDPrimary() PersonIDPrimary {
	return PersonIDPrimary{
		ID: m.ID,
	}
}
