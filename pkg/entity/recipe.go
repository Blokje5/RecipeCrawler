package entity

// TODO add reviews as well
type RecipeMetaData struct {
	name         string
	originalUrl  string
	prepTimeRaw  string
	cookTimeRaw  string
	totalTimeRaw string
	imageUrls    []string
	categories   []string
	description  string
	ingredients  []string
	instructions string
	yield        string

	author struct {
		typeAuthor string
		name       string
	}

	nutrition struct {
		typeNutrition       string
		calories            string
		fatContent          string
		cholesterolContent  string
		sodiumContent       string
		carbohydrateContent string
		fiberContent        string
		proteinContent      string
	}
}
