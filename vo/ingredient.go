package vo

type IngredientVO struct {
	Id          string
	Name        string
	Image       string
	Description string
}

type IngredientQuantityVO struct {
	Id       string
	Name     string
	Quantity int
	Unit     string
}
