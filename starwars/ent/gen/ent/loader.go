// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"io"

	yaml "github.com/goccy/go-yaml"
)

func (c *Client) LoadYML(ctx context.Context, reader io.Reader) error {
	return LoadYML(ctx, reader, c)
}

func (c *Client) LoadJson(ctx context.Context, reader io.Reader) error {
	return LoadJson(ctx, reader, c)
}

func (c *Client) WriteInput(ctx context.Context, writer io.Writer) error {
	return WriteInput(ctx, writer, c)
}

func (c *Client) DumpInput(ctx context.Context) interface{} {
	return DumpInput(ctx, c)
}

func (c *Client) WriteOutput(ctx context.Context, writer io.Writer) error {
	return WriteOutput(ctx, writer, c)
}

func (c *Client) DumpOutput(ctx context.Context) interface{} {
	return DumpOutput(ctx, c)
}

func (c *Client) Clear(ctx context.Context) {
	Clear(ctx, c)
}

func LoadYML(ctx context.Context, reader io.Reader, client *Client) error {
	var input Input
	decoder := yaml.NewDecoder(reader)
	err := decoder.Decode(&input)
	if err != nil {
		return err
	}
	loadInput(ctx, &input, client)
	return nil
}

func LoadJson(ctx context.Context, reader io.Reader, client *Client) error {
	var input Input
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&input)
	if err != nil {
		return err
	}
	loadInput(ctx, &input, client)
	return nil
}

func LoadInput(ctx context.Context, input *Input, client *Client) {
	loadInput(ctx, input, client)
}

func WriteInput(ctx context.Context, writer io.Writer, client *Client) error {
	input := DumpInput(ctx, client)
	encoder := yaml.NewEncoder(writer)
	err := encoder.Encode(input)
	if err != nil {
		return err
	}
	return nil
}

func DumpInput(ctx context.Context, client *Client) *Input {
	return dumpInput(ctx, client)
}

func WriteOutput(ctx context.Context, writer io.Writer, client *Client) error {
	input := DumpOutput(ctx, client)
	encoder := yaml.NewEncoder(writer)
	err := encoder.Encode(input)
	if err != nil {
		return err
	}
	return nil
}

func DumpOutput(ctx context.Context, client *Client) *Output {
	return dumpOutput(ctx, client)
}

func Clear(ctx context.Context, client *Client) {
	client.Film.Delete().Where().ExecX(ctx)
	client.Person.Delete().Where().ExecX(ctx)
	client.Planet.Delete().Where().ExecX(ctx)
	client.Species.Delete().Where().ExecX(ctx)
	client.Starship.Delete().Where().ExecX(ctx)
	client.Vehicle.Delete().Where().ExecX(ctx)
}

type InputID = int
type OutputID = int

// Input holds the top level data for all nodes in the schema
type Input struct {
	Film     []*FilmInput     `yaml:"film"`
	Person   []*PersonInput   `yaml:"person"`
	Planet   []*PlanetInput   `yaml:"planet"`
	Species  []*SpeciesInput  `yaml:"species"`
	Starship []*StarshipInput `yaml:"starship"`
	Vehicle  []*VehicleInput  `yaml:"vehicle"`
}

// Output holds the top level data for all nodes in the schema
type Output struct {
	Film     []*FilmOutput     `yaml:"film"`
	Person   []*PersonOutput   `yaml:"person"`
	Planet   []*PlanetOutput   `yaml:"planet"`
	Species  []*SpeciesOutput  `yaml:"species"`
	Starship []*StarshipOutput `yaml:"starship"`
	Vehicle  []*VehicleOutput  `yaml:"vehicle"`
}

// FilmInput is the load model entity for the Film schema.
type FilmInput struct {
	ID           *int   `json:"id,omitempty"`
	Director     string `json:"director,omitempty"`
	EpisodeID    int    `json:"episode_id,omitempty"`
	OpeningCrawl string `json:"opening_crawl,omitempty"`
	Producer     string `json:"producer,omitempty"`
	Title        string `json:"title,omitempty"`

	Edges FilmEdgeInput `json:"edges"`
}

// FilmOutput is the dump model entity for the Film schema.
type FilmOutput struct {
	ID           int    `json:"id,omitempty"`
	Director     string `json:"director,omitempty"`
	EpisodeID    int    `json:"episode_id,omitempty"`
	OpeningCrawl string `json:"opening_crawl,omitempty"`
	Producer     string `json:"producer,omitempty"`
	Title        string `json:"title,omitempty"`

	Edges FilmEdgeInput `json:"edges"`
}

// FilmEdgeInput holds the relations/edges for other nodes in the graph.
type FilmEdgeInput struct {
	HasPerson   []InputID `json:"has_person,omitempty"`
	HasPlanet   []InputID `json:"has_planet,omitempty"`
	HasStarship []InputID `json:"has_starship,omitempty"`
	HasVehicle  []InputID `json:"has_vehicle,omitempty"`
	HasSpecies  []InputID `json:"has_species,omitempty"`
}

func (pc *FilmCreate) WithFilmInput(p *FilmInput) *FilmCreate {

	pc.mutation.SetDirector(p.Director)
	pc.mutation.SetEpisodeID(p.EpisodeID)
	pc.mutation.SetOpeningCrawl(p.OpeningCrawl)
	pc.mutation.SetProducer(p.Producer)
	pc.mutation.SetTitle(p.Title)
	return pc
}

// PersonInput is the load model entity for the Person schema.
type PersonInput struct {
	ID        *int     `json:"id,omitempty"`
	BirthYear string   `json:"birth_year,omitempty"`
	EyeColor  string   `json:"eye_color,omitempty"`
	Gender    string   `json:"gender,omitempty"`
	HairColor string   `json:"hair_color,omitempty"`
	Height    int      `json:"height,omitempty"`
	Mass      *float64 `json:"mass,omitempty"`
	Name      string   `json:"name,omitempty"`
	SkinColor string   `json:"skin_color,omitempty"`

	Edges PersonEdgeInput `json:"edges"`
}

// PersonOutput is the dump model entity for the Person schema.
type PersonOutput struct {
	ID        int      `json:"id,omitempty"`
	BirthYear string   `json:"birth_year,omitempty"`
	EyeColor  string   `json:"eye_color,omitempty"`
	Gender    string   `json:"gender,omitempty"`
	HairColor string   `json:"hair_color,omitempty"`
	Height    int      `json:"height,omitempty"`
	Mass      *float64 `json:"mass,omitempty"`
	Name      string   `json:"name,omitempty"`
	SkinColor string   `json:"skin_color,omitempty"`

	Edges PersonEdgeInput `json:"edges"`
}

// PersonEdgeInput holds the relations/edges for other nodes in the graph.
type PersonEdgeInput struct {
	PilotedStarship []InputID `json:"piloted_starship,omitempty"`
	PilotedVehicle  []InputID `json:"piloted_vehicle,omitempty"`
	IsOfType        []InputID `json:"is_of_type,omitempty"`
}

func (pc *PersonCreate) WithPersonInput(p *PersonInput) *PersonCreate {

	pc.mutation.SetBirthYear(p.BirthYear)
	pc.mutation.SetEyeColor(p.EyeColor)
	pc.mutation.SetGender(p.Gender)
	pc.mutation.SetHairColor(p.HairColor)
	pc.mutation.SetHeight(p.Height)

	if p.Mass != nil {
		pc.mutation.SetMass(*p.Mass)
	}
	pc.mutation.SetName(p.Name)
	pc.mutation.SetSkinColor(p.SkinColor)
	return pc
}

// PlanetInput is the load model entity for the Planet schema.
type PlanetInput struct {
	ID             *int   `json:"id,omitempty"`
	Climate        string `json:"climate,omitempty"`
	Diameter       int    `json:"diameter,omitempty"`
	Gravity        string `json:"gravity,omitempty"`
	Name           string `json:"name,omitempty"`
	OrbitalPeriod  string `json:"orbital_period,omitempty"`
	Population     int    `json:"population,omitempty"`
	RotationPeriod string `json:"rotation_period,omitempty"`
	SurfaceWater   string `json:"surface_water,omitempty"`
	Terrain        string `json:"terrain,omitempty"`

	Edges PlanetEdgeInput `json:"edges"`
}

// PlanetOutput is the dump model entity for the Planet schema.
type PlanetOutput struct {
	ID             int    `json:"id,omitempty"`
	Climate        string `json:"climate,omitempty"`
	Diameter       int    `json:"diameter,omitempty"`
	Gravity        string `json:"gravity,omitempty"`
	Name           string `json:"name,omitempty"`
	OrbitalPeriod  string `json:"orbital_period,omitempty"`
	Population     int    `json:"population,omitempty"`
	RotationPeriod string `json:"rotation_period,omitempty"`
	SurfaceWater   string `json:"surface_water,omitempty"`
	Terrain        string `json:"terrain,omitempty"`

	Edges PlanetEdgeInput `json:"edges"`
}

// PlanetEdgeInput holds the relations/edges for other nodes in the graph.
type PlanetEdgeInput struct {
	HomeTo []InputID `json:"home_to,omitempty"`
}

func (pc *PlanetCreate) WithPlanetInput(p *PlanetInput) *PlanetCreate {

	pc.mutation.SetClimate(p.Climate)
	pc.mutation.SetDiameter(p.Diameter)
	pc.mutation.SetGravity(p.Gravity)
	pc.mutation.SetName(p.Name)
	pc.mutation.SetOrbitalPeriod(p.OrbitalPeriod)
	pc.mutation.SetPopulation(p.Population)
	pc.mutation.SetRotationPeriod(p.RotationPeriod)
	pc.mutation.SetSurfaceWater(p.SurfaceWater)
	pc.mutation.SetTerrain(p.Terrain)
	return pc
}

// SpeciesInput is the load model entity for the Species schema.
type SpeciesInput struct {
	ID              *int   `json:"id,omitempty"`
	AverageHeight   int    `json:"average_height,omitempty"`
	AverageLifespan string `json:"average_lifespan,omitempty"`
	Classification  string `json:"classification,omitempty"`
	Designation     string `json:"designation,omitempty"`
	Name            string `json:"name,omitempty"`
	SkinColor       string `json:"skin_color,omitempty"`
	EyeColor        string `json:"eye_color,omitempty"`
	HairColor       string `json:"hair_color,omitempty"`
	Language        string `json:"language,omitempty"`

	Edges SpeciesEdgeInput `json:"edges"`
}

// SpeciesOutput is the dump model entity for the Species schema.
type SpeciesOutput struct {
	ID              int    `json:"id,omitempty"`
	AverageHeight   int    `json:"average_height,omitempty"`
	AverageLifespan string `json:"average_lifespan,omitempty"`
	Classification  string `json:"classification,omitempty"`
	Designation     string `json:"designation,omitempty"`
	Name            string `json:"name,omitempty"`
	SkinColor       string `json:"skin_color,omitempty"`
	EyeColor        string `json:"eye_color,omitempty"`
	HairColor       string `json:"hair_color,omitempty"`
	Language        string `json:"language,omitempty"`

	Edges SpeciesEdgeInput `json:"edges"`
}

// SpeciesEdgeInput holds the relations/edges for other nodes in the graph.
type SpeciesEdgeInput struct {
	OriginatesFrom []InputID `json:"originates_from,omitempty"`
}

func (pc *SpeciesCreate) WithSpeciesInput(p *SpeciesInput) *SpeciesCreate {

	pc.mutation.SetAverageHeight(p.AverageHeight)
	pc.mutation.SetAverageLifespan(p.AverageLifespan)
	pc.mutation.SetClassification(p.Classification)
	pc.mutation.SetDesignation(p.Designation)
	pc.mutation.SetName(p.Name)
	pc.mutation.SetSkinColor(p.SkinColor)
	pc.mutation.SetEyeColor(p.EyeColor)
	pc.mutation.SetHairColor(p.HairColor)
	pc.mutation.SetLanguage(p.Language)
	return pc
}

// StarshipInput is the load model entity for the Starship schema.
type StarshipInput struct {
	ID                   *int    `json:"id,omitempty"`
	CargoCapacity        int     `json:"cargo_capacity,omitempty"`
	Class                string  `json:"class,omitempty"`
	Consumables          string  `json:"consumables,omitempty"`
	CostInCredits        int     `json:"cost_in_credits,omitempty"`
	Crew                 string  `json:"crew,omitempty"`
	HyperdriveRating     string  `json:"hyperdrive_rating,omitempty"`
	Length               float64 `json:"length,omitempty"`
	Manufacturer         string  `json:"manufacturer,omitempty"`
	MaxAtmospheringSpeed string  `json:"max_atmosphering_speed,omitempty"`
	MaximumMegalights    string  `json:"maximum_megalights,omitempty"`
	Model                string  `json:"model,omitempty"`
	Name                 string  `json:"name,omitempty"`
	PassengerCapacity    int     `json:"passenger_capacity,omitempty"`

	Edges StarshipEdgeInput `json:"edges"`
}

// StarshipOutput is the dump model entity for the Starship schema.
type StarshipOutput struct {
	ID                   int     `json:"id,omitempty"`
	CargoCapacity        int     `json:"cargo_capacity,omitempty"`
	Class                string  `json:"class,omitempty"`
	Consumables          string  `json:"consumables,omitempty"`
	CostInCredits        int     `json:"cost_in_credits,omitempty"`
	Crew                 string  `json:"crew,omitempty"`
	HyperdriveRating     string  `json:"hyperdrive_rating,omitempty"`
	Length               float64 `json:"length,omitempty"`
	Manufacturer         string  `json:"manufacturer,omitempty"`
	MaxAtmospheringSpeed string  `json:"max_atmosphering_speed,omitempty"`
	MaximumMegalights    string  `json:"maximum_megalights,omitempty"`
	Model                string  `json:"model,omitempty"`
	Name                 string  `json:"name,omitempty"`
	PassengerCapacity    int     `json:"passenger_capacity,omitempty"`

	Edges StarshipEdgeInput `json:"edges"`
}

// StarshipEdgeInput holds the relations/edges for other nodes in the graph.
type StarshipEdgeInput struct {
}

func (pc *StarshipCreate) WithStarshipInput(p *StarshipInput) *StarshipCreate {

	pc.mutation.SetCargoCapacity(p.CargoCapacity)
	pc.mutation.SetClass(p.Class)
	pc.mutation.SetConsumables(p.Consumables)
	pc.mutation.SetCostInCredits(p.CostInCredits)
	pc.mutation.SetCrew(p.Crew)
	pc.mutation.SetHyperdriveRating(p.HyperdriveRating)
	pc.mutation.SetLength(p.Length)
	pc.mutation.SetManufacturer(p.Manufacturer)
	pc.mutation.SetMaxAtmospheringSpeed(p.MaxAtmospheringSpeed)
	pc.mutation.SetMaximumMegalights(p.MaximumMegalights)
	pc.mutation.SetModel(p.Model)
	pc.mutation.SetName(p.Name)
	pc.mutation.SetPassengerCapacity(p.PassengerCapacity)
	return pc
}

// VehicleInput is the load model entity for the Vehicle schema.
type VehicleInput struct {
	ID                   *int     `json:"id,omitempty"`
	CargoCapacity        int      `json:"cargo_capacity,omitempty"`
	Consumables          string   `json:"consumables,omitempty"`
	CostInCredits        int      `json:"cost_in_credits,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Length               *float64 `json:"length,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Model                string   `json:"model,omitempty"`
	Name                 string   `json:"name,omitempty"`
	PassengerCapacity    int      `json:"passenger_capacity,omitempty"`

	Edges VehicleEdgeInput `json:"edges"`
}

// VehicleOutput is the dump model entity for the Vehicle schema.
type VehicleOutput struct {
	ID                   int      `json:"id,omitempty"`
	CargoCapacity        int      `json:"cargo_capacity,omitempty"`
	Consumables          string   `json:"consumables,omitempty"`
	CostInCredits        int      `json:"cost_in_credits,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Length               *float64 `json:"length,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Model                string   `json:"model,omitempty"`
	Name                 string   `json:"name,omitempty"`
	PassengerCapacity    int      `json:"passenger_capacity,omitempty"`

	Edges VehicleEdgeInput `json:"edges"`
}

// VehicleEdgeInput holds the relations/edges for other nodes in the graph.
type VehicleEdgeInput struct {
}

func (pc *VehicleCreate) WithVehicleInput(p *VehicleInput) *VehicleCreate {

	pc.mutation.SetCargoCapacity(p.CargoCapacity)
	pc.mutation.SetConsumables(p.Consumables)
	pc.mutation.SetCostInCredits(p.CostInCredits)
	pc.mutation.SetCrew(p.Crew)

	if p.Length != nil {
		pc.mutation.SetLength(*p.Length)
	}
	pc.mutation.SetManufacturer(p.Manufacturer)
	pc.mutation.SetMaxAtmospheringSpeed(p.MaxAtmospheringSpeed)
	pc.mutation.SetModel(p.Model)
	pc.mutation.SetName(p.Name)
	pc.mutation.SetPassengerCapacity(p.PassengerCapacity)
	return pc
}

func loadInput(ctx context.Context, input *Input, client *Client) {
	filmMap := make(map[InputID]*Film)
	personMap := make(map[InputID]*Person)
	planetMap := make(map[InputID]*Planet)
	speciesMap := make(map[InputID]*Species)
	starshipMap := make(map[InputID]*Starship)
	vehicleMap := make(map[InputID]*Vehicle)

	// load required objects

	// load objects

	for id, obj := range input.Film {
		filmMap[id] = client.Film.Create().
			WithFilmInput(obj).
			SaveX(ctx)
	}

	for id, obj := range input.Person {
		personMap[id] = client.Person.Create().
			WithPersonInput(obj).
			SaveX(ctx)
	}

	for id, obj := range input.Planet {
		planetMap[id] = client.Planet.Create().
			WithPlanetInput(obj).
			SaveX(ctx)
	}

	for id, obj := range input.Species {
		speciesMap[id] = client.Species.Create().
			WithSpeciesInput(obj).
			SaveX(ctx)
	}

	for id, obj := range input.Starship {
		starshipMap[id] = client.Starship.Create().
			WithStarshipInput(obj).
			SaveX(ctx)
	}

	for id, obj := range input.Vehicle {
		vehicleMap[id] = client.Vehicle.Create().
			WithVehicleInput(obj).
			SaveX(ctx)
	}

	// load edges

	for id, obj := range input.Film {
		u := client.Film.UpdateOne(filmMap[id])

		for _, iid := range obj.Edges.HasPerson {
			u.AddHasPerson(personMap[iid])
		}

		for _, iid := range obj.Edges.HasPlanet {
			u.AddHasPlanet(planetMap[iid])
		}

		for _, iid := range obj.Edges.HasStarship {
			u.AddHasStarship(starshipMap[iid])
		}

		for _, iid := range obj.Edges.HasVehicle {
			u.AddHasVehicle(vehicleMap[iid])
		}

		for _, iid := range obj.Edges.HasSpecies {
			u.AddHasSpecies(speciesMap[iid])
		}

		_ = obj
		u.SaveX(ctx)
	}
	for id, obj := range input.Person {
		u := client.Person.UpdateOne(personMap[id])

		for _, iid := range obj.Edges.PilotedStarship {
			u.AddPilotedStarship(starshipMap[iid])
		}

		for _, iid := range obj.Edges.PilotedVehicle {
			u.AddPilotedVehicle(vehicleMap[iid])
		}

		for _, iid := range obj.Edges.IsOfType {
			u.AddIsOfType(speciesMap[iid])
		}

		_ = obj
		u.SaveX(ctx)
	}
	for id, obj := range input.Planet {
		u := client.Planet.UpdateOne(planetMap[id])

		for _, iid := range obj.Edges.HomeTo {
			u.AddHomeTo(personMap[iid])
		}

		_ = obj
		u.SaveX(ctx)
	}
	for id, obj := range input.Species {
		u := client.Species.UpdateOne(speciesMap[id])

		for _, iid := range obj.Edges.OriginatesFrom {
			u.AddOriginatesFrom(planetMap[iid])
		}

		_ = obj
		u.SaveX(ctx)
	}
	for id, obj := range input.Starship {
		u := client.Starship.UpdateOne(starshipMap[id])

		_ = obj
		u.SaveX(ctx)
	}
	for id, obj := range input.Vehicle {
		u := client.Vehicle.UpdateOne(vehicleMap[id])

		_ = obj
		u.SaveX(ctx)
	}
}

func dumpInput(ctx context.Context, client *Client) *Input {
	var input Input
	filmMap := make(map[int]InputID)
	personMap := make(map[int]InputID)
	planetMap := make(map[int]InputID)
	speciesMap := make(map[int]InputID)
	starshipMap := make(map[int]InputID)
	vehicleMap := make(map[int]InputID)

	filmObjects := client.Film.Query().
		WithHasPerson().
		WithHasPlanet().
		WithHasStarship().
		WithHasVehicle().
		WithHasSpecies().
		AllX(ctx)
	input.Film = make([]*FilmInput, len(filmObjects))
	for i := 0; i < len(filmObjects); i++ {
		obj := filmObjects[i]
		inputObj := &FilmInput{
			Director:     obj.Director,
			EpisodeID:    obj.EpisodeID,
			OpeningCrawl: obj.OpeningCrawl,
			Producer:     obj.Producer,
			Title:        obj.Title,
		}
		input.Film[i] = inputObj
		filmMap[obj.ID] = i
	}

	personObjects := client.Person.Query().
		WithPilotedStarship().
		WithPilotedVehicle().
		WithIsOfType().
		WithAppearedIn().
		WithFromPlanet().
		AllX(ctx)
	input.Person = make([]*PersonInput, len(personObjects))
	for i := 0; i < len(personObjects); i++ {
		obj := personObjects[i]
		inputObj := &PersonInput{
			BirthYear: obj.BirthYear,
			EyeColor:  obj.EyeColor,
			Gender:    obj.Gender,
			HairColor: obj.HairColor,
			Height:    obj.Height,
			Mass:      obj.Mass,
			Name:      obj.Name,
			SkinColor: obj.SkinColor,
		}
		input.Person[i] = inputObj
		personMap[obj.ID] = i
	}

	planetObjects := client.Planet.Query().
		WithHomeTo().
		WithAppearedIn().
		WithOriginOf().
		AllX(ctx)
	input.Planet = make([]*PlanetInput, len(planetObjects))
	for i := 0; i < len(planetObjects); i++ {
		obj := planetObjects[i]
		inputObj := &PlanetInput{
			Climate:        obj.Climate,
			Diameter:       obj.Diameter,
			Gravity:        obj.Gravity,
			Name:           obj.Name,
			OrbitalPeriod:  obj.OrbitalPeriod,
			Population:     obj.Population,
			RotationPeriod: obj.RotationPeriod,
			SurfaceWater:   obj.SurfaceWater,
			Terrain:        obj.Terrain,
		}
		input.Planet[i] = inputObj
		planetMap[obj.ID] = i
	}

	speciesObjects := client.Species.Query().
		WithOriginatesFrom().
		WithAppearedIn().
		WithIncludesPerson().
		AllX(ctx)
	input.Species = make([]*SpeciesInput, len(speciesObjects))
	for i := 0; i < len(speciesObjects); i++ {
		obj := speciesObjects[i]
		inputObj := &SpeciesInput{
			AverageHeight:   obj.AverageHeight,
			AverageLifespan: obj.AverageLifespan,
			Classification:  obj.Classification,
			Designation:     obj.Designation,
			Name:            obj.Name,
			SkinColor:       obj.SkinColor,
			EyeColor:        obj.EyeColor,
			HairColor:       obj.HairColor,
			Language:        obj.Language,
		}
		input.Species[i] = inputObj
		speciesMap[obj.ID] = i
	}

	starshipObjects := client.Starship.Query().
		WithAppearedIn().
		WithPilotedBy().
		AllX(ctx)
	input.Starship = make([]*StarshipInput, len(starshipObjects))
	for i := 0; i < len(starshipObjects); i++ {
		obj := starshipObjects[i]
		inputObj := &StarshipInput{
			CargoCapacity:        obj.CargoCapacity,
			Class:                obj.Class,
			Consumables:          obj.Consumables,
			CostInCredits:        obj.CostInCredits,
			Crew:                 obj.Crew,
			HyperdriveRating:     obj.HyperdriveRating,
			Length:               obj.Length,
			Manufacturer:         obj.Manufacturer,
			MaxAtmospheringSpeed: obj.MaxAtmospheringSpeed,
			MaximumMegalights:    obj.MaximumMegalights,
			Model:                obj.Model,
			Name:                 obj.Name,
			PassengerCapacity:    obj.PassengerCapacity,
		}
		input.Starship[i] = inputObj
		starshipMap[obj.ID] = i
	}

	vehicleObjects := client.Vehicle.Query().
		WithAppearedIn().
		WithPilotedBy().
		AllX(ctx)
	input.Vehicle = make([]*VehicleInput, len(vehicleObjects))
	for i := 0; i < len(vehicleObjects); i++ {
		obj := vehicleObjects[i]
		inputObj := &VehicleInput{
			CargoCapacity:        obj.CargoCapacity,
			Consumables:          obj.Consumables,
			CostInCredits:        obj.CostInCredits,
			Crew:                 obj.Crew,
			Length:               obj.Length,
			Manufacturer:         obj.Manufacturer,
			MaxAtmospheringSpeed: obj.MaxAtmospheringSpeed,
			Model:                obj.Model,
			Name:                 obj.Name,
			PassengerCapacity:    obj.PassengerCapacity,
		}
		input.Vehicle[i] = inputObj
		vehicleMap[obj.ID] = i
	}

	for i := 0; i < len(filmObjects); i++ {
		obj := filmObjects[i]
		inputObj := input.Film[i]

		for _, edgeObj := range obj.Edges.HasPerson {
			inputObj.Edges.HasPerson = append(inputObj.Edges.HasPerson, personMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.HasPlanet {
			inputObj.Edges.HasPlanet = append(inputObj.Edges.HasPlanet, planetMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.HasStarship {
			inputObj.Edges.HasStarship = append(inputObj.Edges.HasStarship, starshipMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.HasVehicle {
			inputObj.Edges.HasVehicle = append(inputObj.Edges.HasVehicle, vehicleMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.HasSpecies {
			inputObj.Edges.HasSpecies = append(inputObj.Edges.HasSpecies, speciesMap[edgeObj.ID])
		}
		_ = obj
		_ = inputObj
	}

	for i := 0; i < len(personObjects); i++ {
		obj := personObjects[i]
		inputObj := input.Person[i]

		for _, edgeObj := range obj.Edges.PilotedStarship {
			inputObj.Edges.PilotedStarship = append(inputObj.Edges.PilotedStarship, starshipMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.PilotedVehicle {
			inputObj.Edges.PilotedVehicle = append(inputObj.Edges.PilotedVehicle, vehicleMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.IsOfType {
			inputObj.Edges.IsOfType = append(inputObj.Edges.IsOfType, speciesMap[edgeObj.ID])
		}
		_ = obj
		_ = inputObj
	}

	for i := 0; i < len(planetObjects); i++ {
		obj := planetObjects[i]
		inputObj := input.Planet[i]

		for _, edgeObj := range obj.Edges.HomeTo {
			inputObj.Edges.HomeTo = append(inputObj.Edges.HomeTo, personMap[edgeObj.ID])
		}
		_ = obj
		_ = inputObj
	}

	for i := 0; i < len(speciesObjects); i++ {
		obj := speciesObjects[i]
		inputObj := input.Species[i]

		for _, edgeObj := range obj.Edges.OriginatesFrom {
			inputObj.Edges.OriginatesFrom = append(inputObj.Edges.OriginatesFrom, planetMap[edgeObj.ID])
		}
		_ = obj
		_ = inputObj
	}

	for i := 0; i < len(starshipObjects); i++ {
		obj := starshipObjects[i]
		inputObj := input.Starship[i]

		_ = obj
		_ = inputObj
	}

	for i := 0; i < len(vehicleObjects); i++ {
		obj := vehicleObjects[i]
		inputObj := input.Vehicle[i]

		_ = obj
		_ = inputObj
	}

	return &input
}

func dumpOutput(ctx context.Context, client *Client) *Output {
	var output Output
	filmMap := make(map[int]OutputID)
	personMap := make(map[int]OutputID)
	planetMap := make(map[int]OutputID)
	speciesMap := make(map[int]OutputID)
	starshipMap := make(map[int]OutputID)
	vehicleMap := make(map[int]OutputID)

	filmObjects := client.Film.Query().
		WithHasPerson().
		WithHasPlanet().
		WithHasStarship().
		WithHasVehicle().
		WithHasSpecies().
		AllX(ctx)
	output.Film = make([]*FilmOutput, len(filmObjects))
	for i := 0; i < len(filmObjects); i++ {
		obj := filmObjects[i]
		outputObj := &FilmOutput{
			ID:           obj.ID,
			Director:     obj.Director,
			EpisodeID:    obj.EpisodeID,
			OpeningCrawl: obj.OpeningCrawl,
			Producer:     obj.Producer,
			Title:        obj.Title,
		}
		output.Film[i] = outputObj
		filmMap[obj.ID] = i
	}

	personObjects := client.Person.Query().
		WithPilotedStarship().
		WithPilotedVehicle().
		WithIsOfType().
		WithAppearedIn().
		WithFromPlanet().
		AllX(ctx)
	output.Person = make([]*PersonOutput, len(personObjects))
	for i := 0; i < len(personObjects); i++ {
		obj := personObjects[i]
		outputObj := &PersonOutput{
			ID:        obj.ID,
			BirthYear: obj.BirthYear,
			EyeColor:  obj.EyeColor,
			Gender:    obj.Gender,
			HairColor: obj.HairColor,
			Height:    obj.Height,
			Mass:      obj.Mass,
			Name:      obj.Name,
			SkinColor: obj.SkinColor,
		}
		output.Person[i] = outputObj
		personMap[obj.ID] = i
	}

	planetObjects := client.Planet.Query().
		WithHomeTo().
		WithAppearedIn().
		WithOriginOf().
		AllX(ctx)
	output.Planet = make([]*PlanetOutput, len(planetObjects))
	for i := 0; i < len(planetObjects); i++ {
		obj := planetObjects[i]
		outputObj := &PlanetOutput{
			ID:             obj.ID,
			Climate:        obj.Climate,
			Diameter:       obj.Diameter,
			Gravity:        obj.Gravity,
			Name:           obj.Name,
			OrbitalPeriod:  obj.OrbitalPeriod,
			Population:     obj.Population,
			RotationPeriod: obj.RotationPeriod,
			SurfaceWater:   obj.SurfaceWater,
			Terrain:        obj.Terrain,
		}
		output.Planet[i] = outputObj
		planetMap[obj.ID] = i
	}

	speciesObjects := client.Species.Query().
		WithOriginatesFrom().
		WithAppearedIn().
		WithIncludesPerson().
		AllX(ctx)
	output.Species = make([]*SpeciesOutput, len(speciesObjects))
	for i := 0; i < len(speciesObjects); i++ {
		obj := speciesObjects[i]
		outputObj := &SpeciesOutput{
			ID:              obj.ID,
			AverageHeight:   obj.AverageHeight,
			AverageLifespan: obj.AverageLifespan,
			Classification:  obj.Classification,
			Designation:     obj.Designation,
			Name:            obj.Name,
			SkinColor:       obj.SkinColor,
			EyeColor:        obj.EyeColor,
			HairColor:       obj.HairColor,
			Language:        obj.Language,
		}
		output.Species[i] = outputObj
		speciesMap[obj.ID] = i
	}

	starshipObjects := client.Starship.Query().
		WithAppearedIn().
		WithPilotedBy().
		AllX(ctx)
	output.Starship = make([]*StarshipOutput, len(starshipObjects))
	for i := 0; i < len(starshipObjects); i++ {
		obj := starshipObjects[i]
		outputObj := &StarshipOutput{
			ID:                   obj.ID,
			CargoCapacity:        obj.CargoCapacity,
			Class:                obj.Class,
			Consumables:          obj.Consumables,
			CostInCredits:        obj.CostInCredits,
			Crew:                 obj.Crew,
			HyperdriveRating:     obj.HyperdriveRating,
			Length:               obj.Length,
			Manufacturer:         obj.Manufacturer,
			MaxAtmospheringSpeed: obj.MaxAtmospheringSpeed,
			MaximumMegalights:    obj.MaximumMegalights,
			Model:                obj.Model,
			Name:                 obj.Name,
			PassengerCapacity:    obj.PassengerCapacity,
		}
		output.Starship[i] = outputObj
		starshipMap[obj.ID] = i
	}

	vehicleObjects := client.Vehicle.Query().
		WithAppearedIn().
		WithPilotedBy().
		AllX(ctx)
	output.Vehicle = make([]*VehicleOutput, len(vehicleObjects))
	for i := 0; i < len(vehicleObjects); i++ {
		obj := vehicleObjects[i]
		outputObj := &VehicleOutput{
			ID:                   obj.ID,
			CargoCapacity:        obj.CargoCapacity,
			Consumables:          obj.Consumables,
			CostInCredits:        obj.CostInCredits,
			Crew:                 obj.Crew,
			Length:               obj.Length,
			Manufacturer:         obj.Manufacturer,
			MaxAtmospheringSpeed: obj.MaxAtmospheringSpeed,
			Model:                obj.Model,
			Name:                 obj.Name,
			PassengerCapacity:    obj.PassengerCapacity,
		}
		output.Vehicle[i] = outputObj
		vehicleMap[obj.ID] = i
	}

	for i := 0; i < len(filmObjects); i++ {
		obj := filmObjects[i]
		outputObj := output.Film[i]

		for _, edgeObj := range obj.Edges.HasPerson {
			outputObj.Edges.HasPerson = append(outputObj.Edges.HasPerson, personMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.HasPlanet {
			outputObj.Edges.HasPlanet = append(outputObj.Edges.HasPlanet, planetMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.HasStarship {
			outputObj.Edges.HasStarship = append(outputObj.Edges.HasStarship, starshipMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.HasVehicle {
			outputObj.Edges.HasVehicle = append(outputObj.Edges.HasVehicle, vehicleMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.HasSpecies {
			outputObj.Edges.HasSpecies = append(outputObj.Edges.HasSpecies, speciesMap[edgeObj.ID])
		}
		_ = obj
		_ = outputObj
	}

	for i := 0; i < len(personObjects); i++ {
		obj := personObjects[i]
		outputObj := output.Person[i]

		for _, edgeObj := range obj.Edges.PilotedStarship {
			outputObj.Edges.PilotedStarship = append(outputObj.Edges.PilotedStarship, starshipMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.PilotedVehicle {
			outputObj.Edges.PilotedVehicle = append(outputObj.Edges.PilotedVehicle, vehicleMap[edgeObj.ID])
		}
		for _, edgeObj := range obj.Edges.IsOfType {
			outputObj.Edges.IsOfType = append(outputObj.Edges.IsOfType, speciesMap[edgeObj.ID])
		}
		_ = obj
		_ = outputObj
	}

	for i := 0; i < len(planetObjects); i++ {
		obj := planetObjects[i]
		outputObj := output.Planet[i]

		for _, edgeObj := range obj.Edges.HomeTo {
			outputObj.Edges.HomeTo = append(outputObj.Edges.HomeTo, personMap[edgeObj.ID])
		}
		_ = obj
		_ = outputObj
	}

	for i := 0; i < len(speciesObjects); i++ {
		obj := speciesObjects[i]
		outputObj := output.Species[i]

		for _, edgeObj := range obj.Edges.OriginatesFrom {
			outputObj.Edges.OriginatesFrom = append(outputObj.Edges.OriginatesFrom, planetMap[edgeObj.ID])
		}
		_ = obj
		_ = outputObj
	}

	for i := 0; i < len(starshipObjects); i++ {
		obj := starshipObjects[i]
		outputObj := output.Starship[i]

		_ = obj
		_ = outputObj
	}

	for i := 0; i < len(vehicleObjects); i++ {
		obj := vehicleObjects[i]
		outputObj := output.Vehicle[i]

		_ = obj
		_ = outputObj
	}

	return &output
}
