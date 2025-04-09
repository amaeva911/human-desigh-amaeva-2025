package dictionaries

type GateLine struct {
	id                    uint8
	Name                  string
	Description           string
	Gates                 Gate
	ExaltationPlanet      CelestialBody
	FallPlanet            CelestialBody
	ExaltationDescription string
	FallDescription       string
	updatedDateTime       string
}
