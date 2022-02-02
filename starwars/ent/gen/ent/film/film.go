// Code generated by entc, DO NOT EDIT.

package film

const (
	// Label holds the string label denoting the film type in the database.
	Label = "film"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDirector holds the string denoting the director field in the database.
	FieldDirector = "director"
	// FieldEpisodeID holds the string denoting the episode_id field in the database.
	FieldEpisodeID = "episode_id"
	// FieldOpeningCrawl holds the string denoting the opening_crawl field in the database.
	FieldOpeningCrawl = "opening_crawl"
	// FieldProducer holds the string denoting the producer field in the database.
	FieldProducer = "producer"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// EdgeHasPerson holds the string denoting the has_person edge name in mutations.
	EdgeHasPerson = "has_person"
	// EdgeHasPlanet holds the string denoting the has_planet edge name in mutations.
	EdgeHasPlanet = "has_planet"
	// EdgeHasStarship holds the string denoting the has_starship edge name in mutations.
	EdgeHasStarship = "has_starship"
	// EdgeHasVehicle holds the string denoting the has_vehicle edge name in mutations.
	EdgeHasVehicle = "has_vehicle"
	// EdgeHasSpecies holds the string denoting the has_species edge name in mutations.
	EdgeHasSpecies = "has_species"
	// Table holds the table name of the film in the database.
	Table = "films"
	// HasPersonTable is the table that holds the has_person relation/edge. The primary key declared below.
	HasPersonTable = "film_has_person"
	// HasPersonInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	HasPersonInverseTable = "persons"
	// HasPlanetTable is the table that holds the has_planet relation/edge. The primary key declared below.
	HasPlanetTable = "film_has_planet"
	// HasPlanetInverseTable is the table name for the Planet entity.
	// It exists in this package in order to avoid circular dependency with the "planet" package.
	HasPlanetInverseTable = "planets"
	// HasStarshipTable is the table that holds the has_starship relation/edge. The primary key declared below.
	HasStarshipTable = "film_has_starship"
	// HasStarshipInverseTable is the table name for the Starship entity.
	// It exists in this package in order to avoid circular dependency with the "starship" package.
	HasStarshipInverseTable = "starships"
	// HasVehicleTable is the table that holds the has_vehicle relation/edge. The primary key declared below.
	HasVehicleTable = "film_has_vehicle"
	// HasVehicleInverseTable is the table name for the Vehicle entity.
	// It exists in this package in order to avoid circular dependency with the "vehicle" package.
	HasVehicleInverseTable = "vehicles"
	// HasSpeciesTable is the table that holds the has_species relation/edge. The primary key declared below.
	HasSpeciesTable = "film_has_species"
	// HasSpeciesInverseTable is the table name for the Species entity.
	// It exists in this package in order to avoid circular dependency with the "species" package.
	HasSpeciesInverseTable = "species"
)

// Columns holds all SQL columns for film fields.
var Columns = []string{
	FieldID,
	FieldDirector,
	FieldEpisodeID,
	FieldOpeningCrawl,
	FieldProducer,
	FieldTitle,
}

var (
	// HasPersonPrimaryKey and HasPersonColumn2 are the table columns denoting the
	// primary key for the has_person relation (M2M).
	HasPersonPrimaryKey = []string{"film_id", "person_id"}
	// HasPlanetPrimaryKey and HasPlanetColumn2 are the table columns denoting the
	// primary key for the has_planet relation (M2M).
	HasPlanetPrimaryKey = []string{"film_id", "planet_id"}
	// HasStarshipPrimaryKey and HasStarshipColumn2 are the table columns denoting the
	// primary key for the has_starship relation (M2M).
	HasStarshipPrimaryKey = []string{"film_id", "starship_id"}
	// HasVehiclePrimaryKey and HasVehicleColumn2 are the table columns denoting the
	// primary key for the has_vehicle relation (M2M).
	HasVehiclePrimaryKey = []string{"film_id", "vehicle_id"}
	// HasSpeciesPrimaryKey and HasSpeciesColumn2 are the table columns denoting the
	// primary key for the has_species relation (M2M).
	HasSpeciesPrimaryKey = []string{"film_id", "species_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
