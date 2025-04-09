package dictionaries

type Adress struct {
	id              uint64
	Country         string
	Region          string
	City            string
	Latitude        float64
	Longitude       float64
	Timezone        int8
	updatedDateTime string
}
