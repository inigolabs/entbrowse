// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FilmsColumns holds the columns for the "films" table.
	FilmsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "director", Type: field.TypeString},
		{Name: "episode_id", Type: field.TypeInt},
		{Name: "opening_crawl", Type: field.TypeString},
		{Name: "producer", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
	}
	// FilmsTable holds the schema information for the "films" table.
	FilmsTable = &schema.Table{
		Name:       "films",
		Columns:    FilmsColumns,
		PrimaryKey: []*schema.Column{FilmsColumns[0]},
	}
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "birth_year", Type: field.TypeString},
		{Name: "eye_color", Type: field.TypeString},
		{Name: "gender", Type: field.TypeString},
		{Name: "hair_color", Type: field.TypeString},
		{Name: "height", Type: field.TypeInt},
		{Name: "mass", Type: field.TypeFloat64, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "skin_color", Type: field.TypeString},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
	}
	// PlanetsColumns holds the columns for the "planets" table.
	PlanetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "climate", Type: field.TypeString},
		{Name: "diameter", Type: field.TypeInt, Nullable: true},
		{Name: "gravity", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "orbital_period", Type: field.TypeString},
		{Name: "population", Type: field.TypeInt, Nullable: true},
		{Name: "rotation_period", Type: field.TypeString},
		{Name: "surface_water", Type: field.TypeString},
		{Name: "terrain", Type: field.TypeString},
	}
	// PlanetsTable holds the schema information for the "planets" table.
	PlanetsTable = &schema.Table{
		Name:       "planets",
		Columns:    PlanetsColumns,
		PrimaryKey: []*schema.Column{PlanetsColumns[0]},
	}
	// SpeciesColumns holds the columns for the "species" table.
	SpeciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "average_height", Type: field.TypeInt},
		{Name: "average_lifespan", Type: field.TypeString},
		{Name: "classification", Type: field.TypeString},
		{Name: "designation", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "skin_color", Type: field.TypeString},
		{Name: "eye_color", Type: field.TypeString},
		{Name: "hair_color", Type: field.TypeString},
		{Name: "language", Type: field.TypeString},
	}
	// SpeciesTable holds the schema information for the "species" table.
	SpeciesTable = &schema.Table{
		Name:       "species",
		Columns:    SpeciesColumns,
		PrimaryKey: []*schema.Column{SpeciesColumns[0]},
	}
	// StarshipsColumns holds the columns for the "starships" table.
	StarshipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "cargo_capacity", Type: field.TypeInt, Nullable: true},
		{Name: "class", Type: field.TypeString},
		{Name: "consumables", Type: field.TypeString},
		{Name: "cost_in_credits", Type: field.TypeInt},
		{Name: "crew", Type: field.TypeString},
		{Name: "hyperdrive_rating", Type: field.TypeString},
		{Name: "length", Type: field.TypeFloat64},
		{Name: "manufacturer", Type: field.TypeString},
		{Name: "max_atmosphering_speed", Type: field.TypeString},
		{Name: "maximum_megalights", Type: field.TypeString, Nullable: true},
		{Name: "model", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "passenger_capacity", Type: field.TypeInt, Nullable: true},
	}
	// StarshipsTable holds the schema information for the "starships" table.
	StarshipsTable = &schema.Table{
		Name:       "starships",
		Columns:    StarshipsColumns,
		PrimaryKey: []*schema.Column{StarshipsColumns[0]},
	}
	// VehiclesColumns holds the columns for the "vehicles" table.
	VehiclesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "cargo_capacity", Type: field.TypeInt, Nullable: true},
		{Name: "consumables", Type: field.TypeString},
		{Name: "cost_in_credits", Type: field.TypeInt},
		{Name: "crew", Type: field.TypeString},
		{Name: "length", Type: field.TypeFloat64, Nullable: true},
		{Name: "manufacturer", Type: field.TypeString},
		{Name: "max_atmosphering_speed", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "passenger_capacity", Type: field.TypeInt, Nullable: true},
	}
	// VehiclesTable holds the schema information for the "vehicles" table.
	VehiclesTable = &schema.Table{
		Name:       "vehicles",
		Columns:    VehiclesColumns,
		PrimaryKey: []*schema.Column{VehiclesColumns[0]},
	}
	// FilmHasPersonColumns holds the columns for the "film_has_person" table.
	FilmHasPersonColumns = []*schema.Column{
		{Name: "film_id", Type: field.TypeInt},
		{Name: "person_id", Type: field.TypeInt},
	}
	// FilmHasPersonTable holds the schema information for the "film_has_person" table.
	FilmHasPersonTable = &schema.Table{
		Name:       "film_has_person",
		Columns:    FilmHasPersonColumns,
		PrimaryKey: []*schema.Column{FilmHasPersonColumns[0], FilmHasPersonColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "film_has_person_film_id",
				Columns:    []*schema.Column{FilmHasPersonColumns[0]},
				RefColumns: []*schema.Column{FilmsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "film_has_person_person_id",
				Columns:    []*schema.Column{FilmHasPersonColumns[1]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FilmHasPlanetColumns holds the columns for the "film_has_planet" table.
	FilmHasPlanetColumns = []*schema.Column{
		{Name: "film_id", Type: field.TypeInt},
		{Name: "planet_id", Type: field.TypeInt},
	}
	// FilmHasPlanetTable holds the schema information for the "film_has_planet" table.
	FilmHasPlanetTable = &schema.Table{
		Name:       "film_has_planet",
		Columns:    FilmHasPlanetColumns,
		PrimaryKey: []*schema.Column{FilmHasPlanetColumns[0], FilmHasPlanetColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "film_has_planet_film_id",
				Columns:    []*schema.Column{FilmHasPlanetColumns[0]},
				RefColumns: []*schema.Column{FilmsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "film_has_planet_planet_id",
				Columns:    []*schema.Column{FilmHasPlanetColumns[1]},
				RefColumns: []*schema.Column{PlanetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FilmHasStarshipColumns holds the columns for the "film_has_starship" table.
	FilmHasStarshipColumns = []*schema.Column{
		{Name: "film_id", Type: field.TypeInt},
		{Name: "starship_id", Type: field.TypeInt},
	}
	// FilmHasStarshipTable holds the schema information for the "film_has_starship" table.
	FilmHasStarshipTable = &schema.Table{
		Name:       "film_has_starship",
		Columns:    FilmHasStarshipColumns,
		PrimaryKey: []*schema.Column{FilmHasStarshipColumns[0], FilmHasStarshipColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "film_has_starship_film_id",
				Columns:    []*schema.Column{FilmHasStarshipColumns[0]},
				RefColumns: []*schema.Column{FilmsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "film_has_starship_starship_id",
				Columns:    []*schema.Column{FilmHasStarshipColumns[1]},
				RefColumns: []*schema.Column{StarshipsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FilmHasVehicleColumns holds the columns for the "film_has_vehicle" table.
	FilmHasVehicleColumns = []*schema.Column{
		{Name: "film_id", Type: field.TypeInt},
		{Name: "vehicle_id", Type: field.TypeInt},
	}
	// FilmHasVehicleTable holds the schema information for the "film_has_vehicle" table.
	FilmHasVehicleTable = &schema.Table{
		Name:       "film_has_vehicle",
		Columns:    FilmHasVehicleColumns,
		PrimaryKey: []*schema.Column{FilmHasVehicleColumns[0], FilmHasVehicleColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "film_has_vehicle_film_id",
				Columns:    []*schema.Column{FilmHasVehicleColumns[0]},
				RefColumns: []*schema.Column{FilmsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "film_has_vehicle_vehicle_id",
				Columns:    []*schema.Column{FilmHasVehicleColumns[1]},
				RefColumns: []*schema.Column{VehiclesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FilmHasSpeciesColumns holds the columns for the "film_has_species" table.
	FilmHasSpeciesColumns = []*schema.Column{
		{Name: "film_id", Type: field.TypeInt},
		{Name: "species_id", Type: field.TypeInt},
	}
	// FilmHasSpeciesTable holds the schema information for the "film_has_species" table.
	FilmHasSpeciesTable = &schema.Table{
		Name:       "film_has_species",
		Columns:    FilmHasSpeciesColumns,
		PrimaryKey: []*schema.Column{FilmHasSpeciesColumns[0], FilmHasSpeciesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "film_has_species_film_id",
				Columns:    []*schema.Column{FilmHasSpeciesColumns[0]},
				RefColumns: []*schema.Column{FilmsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "film_has_species_species_id",
				Columns:    []*schema.Column{FilmHasSpeciesColumns[1]},
				RefColumns: []*schema.Column{SpeciesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PersonPilotedStarshipColumns holds the columns for the "person_piloted_starship" table.
	PersonPilotedStarshipColumns = []*schema.Column{
		{Name: "person_id", Type: field.TypeInt},
		{Name: "starship_id", Type: field.TypeInt},
	}
	// PersonPilotedStarshipTable holds the schema information for the "person_piloted_starship" table.
	PersonPilotedStarshipTable = &schema.Table{
		Name:       "person_piloted_starship",
		Columns:    PersonPilotedStarshipColumns,
		PrimaryKey: []*schema.Column{PersonPilotedStarshipColumns[0], PersonPilotedStarshipColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "person_piloted_starship_person_id",
				Columns:    []*schema.Column{PersonPilotedStarshipColumns[0]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "person_piloted_starship_starship_id",
				Columns:    []*schema.Column{PersonPilotedStarshipColumns[1]},
				RefColumns: []*schema.Column{StarshipsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PersonPilotedVehicleColumns holds the columns for the "person_piloted_vehicle" table.
	PersonPilotedVehicleColumns = []*schema.Column{
		{Name: "person_id", Type: field.TypeInt},
		{Name: "vehicle_id", Type: field.TypeInt},
	}
	// PersonPilotedVehicleTable holds the schema information for the "person_piloted_vehicle" table.
	PersonPilotedVehicleTable = &schema.Table{
		Name:       "person_piloted_vehicle",
		Columns:    PersonPilotedVehicleColumns,
		PrimaryKey: []*schema.Column{PersonPilotedVehicleColumns[0], PersonPilotedVehicleColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "person_piloted_vehicle_person_id",
				Columns:    []*schema.Column{PersonPilotedVehicleColumns[0]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "person_piloted_vehicle_vehicle_id",
				Columns:    []*schema.Column{PersonPilotedVehicleColumns[1]},
				RefColumns: []*schema.Column{VehiclesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PersonIsOfTypeColumns holds the columns for the "person_is_of_type" table.
	PersonIsOfTypeColumns = []*schema.Column{
		{Name: "person_id", Type: field.TypeInt},
		{Name: "species_id", Type: field.TypeInt},
	}
	// PersonIsOfTypeTable holds the schema information for the "person_is_of_type" table.
	PersonIsOfTypeTable = &schema.Table{
		Name:       "person_is_of_type",
		Columns:    PersonIsOfTypeColumns,
		PrimaryKey: []*schema.Column{PersonIsOfTypeColumns[0], PersonIsOfTypeColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "person_is_of_type_person_id",
				Columns:    []*schema.Column{PersonIsOfTypeColumns[0]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "person_is_of_type_species_id",
				Columns:    []*schema.Column{PersonIsOfTypeColumns[1]},
				RefColumns: []*schema.Column{SpeciesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PlanetHomeToColumns holds the columns for the "planet_home_to" table.
	PlanetHomeToColumns = []*schema.Column{
		{Name: "planet_id", Type: field.TypeInt},
		{Name: "person_id", Type: field.TypeInt},
	}
	// PlanetHomeToTable holds the schema information for the "planet_home_to" table.
	PlanetHomeToTable = &schema.Table{
		Name:       "planet_home_to",
		Columns:    PlanetHomeToColumns,
		PrimaryKey: []*schema.Column{PlanetHomeToColumns[0], PlanetHomeToColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "planet_home_to_planet_id",
				Columns:    []*schema.Column{PlanetHomeToColumns[0]},
				RefColumns: []*schema.Column{PlanetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "planet_home_to_person_id",
				Columns:    []*schema.Column{PlanetHomeToColumns[1]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SpeciesOriginatesFromColumns holds the columns for the "species_originates_from" table.
	SpeciesOriginatesFromColumns = []*schema.Column{
		{Name: "species_id", Type: field.TypeInt},
		{Name: "planet_id", Type: field.TypeInt},
	}
	// SpeciesOriginatesFromTable holds the schema information for the "species_originates_from" table.
	SpeciesOriginatesFromTable = &schema.Table{
		Name:       "species_originates_from",
		Columns:    SpeciesOriginatesFromColumns,
		PrimaryKey: []*schema.Column{SpeciesOriginatesFromColumns[0], SpeciesOriginatesFromColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "species_originates_from_species_id",
				Columns:    []*schema.Column{SpeciesOriginatesFromColumns[0]},
				RefColumns: []*schema.Column{SpeciesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "species_originates_from_planet_id",
				Columns:    []*schema.Column{SpeciesOriginatesFromColumns[1]},
				RefColumns: []*schema.Column{PlanetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FilmsTable,
		PersonsTable,
		PlanetsTable,
		SpeciesTable,
		StarshipsTable,
		VehiclesTable,
		FilmHasPersonTable,
		FilmHasPlanetTable,
		FilmHasStarshipTable,
		FilmHasVehicleTable,
		FilmHasSpeciesTable,
		PersonPilotedStarshipTable,
		PersonPilotedVehicleTable,
		PersonIsOfTypeTable,
		PlanetHomeToTable,
		SpeciesOriginatesFromTable,
	}
)

func init() {
	FilmHasPersonTable.ForeignKeys[0].RefTable = FilmsTable
	FilmHasPersonTable.ForeignKeys[1].RefTable = PersonsTable
	FilmHasPlanetTable.ForeignKeys[0].RefTable = FilmsTable
	FilmHasPlanetTable.ForeignKeys[1].RefTable = PlanetsTable
	FilmHasStarshipTable.ForeignKeys[0].RefTable = FilmsTable
	FilmHasStarshipTable.ForeignKeys[1].RefTable = StarshipsTable
	FilmHasVehicleTable.ForeignKeys[0].RefTable = FilmsTable
	FilmHasVehicleTable.ForeignKeys[1].RefTable = VehiclesTable
	FilmHasSpeciesTable.ForeignKeys[0].RefTable = FilmsTable
	FilmHasSpeciesTable.ForeignKeys[1].RefTable = SpeciesTable
	PersonPilotedStarshipTable.ForeignKeys[0].RefTable = PersonsTable
	PersonPilotedStarshipTable.ForeignKeys[1].RefTable = StarshipsTable
	PersonPilotedVehicleTable.ForeignKeys[0].RefTable = PersonsTable
	PersonPilotedVehicleTable.ForeignKeys[1].RefTable = VehiclesTable
	PersonIsOfTypeTable.ForeignKeys[0].RefTable = PersonsTable
	PersonIsOfTypeTable.ForeignKeys[1].RefTable = SpeciesTable
	PlanetHomeToTable.ForeignKeys[0].RefTable = PlanetsTable
	PlanetHomeToTable.ForeignKeys[1].RefTable = PersonsTable
	SpeciesOriginatesFromTable.ForeignKeys[0].RefTable = SpeciesTable
	SpeciesOriginatesFromTable.ForeignKeys[1].RefTable = PlanetsTable
}