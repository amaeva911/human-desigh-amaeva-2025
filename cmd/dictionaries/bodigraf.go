package dictionaries

type Bodigraf struct {
	id               uint64
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
