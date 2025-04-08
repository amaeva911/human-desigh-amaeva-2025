package dictionaries

type Channel struct {
	Id              uint8
	Name            string
	Description     string
	FirstGate       Gate
	SecondGate      Gate
	UpdatedDateTime string
}
