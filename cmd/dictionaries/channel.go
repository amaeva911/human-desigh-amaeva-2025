package dictionaries

type Channel struct {
	id              uint8
	Name            string
	Description     string
	FirstGate       Gate
	SecondGate      Gate
	updatedDateTime string
}
