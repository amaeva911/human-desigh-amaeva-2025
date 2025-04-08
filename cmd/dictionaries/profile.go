package dictionaries

type Profile struct {
	Id              uint8
	PersonalityLine Line
	DesignLine      Line
	Name            string
	Description     string
	UpdatedDateTime string
}
