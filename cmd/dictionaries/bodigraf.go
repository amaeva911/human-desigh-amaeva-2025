package dictionaries

type Bodigraf struct {
	Id               uint64
	BirthDateTimeUtc string
	Agress           Adress
	UserName         string
	Type             Type
	Profile          Profile
	Authority        Authority
	Certainty        Certainty
	SvgBodigraf      string
	createdDateTime  string
}
