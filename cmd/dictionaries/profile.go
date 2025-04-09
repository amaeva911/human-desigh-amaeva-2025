package dictionaries

type Profile struct {
	id              uint8
	PersonalityLine Line
	DesignLine      Line
	Name            string
	Description     string
	updatedDateTime string
}
