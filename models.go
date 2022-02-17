package snipeit

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"items,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Coffee   Coffee `json:"coffee"`
	Quantity int    `json:"quantity"`
}

// Coffee -
type Coffee struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Teaser      string             `json:"teaser"`
	Collection  string             `json:"collection"`
	Origin      string             `json:"origin"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Image       string             `json:"image"`
	Ingredient  []CoffeeIngredient `json:"ingredients"`
}

// Ingredient -
type CoffeeIngredient struct {
	ID       int    `json:"ingredient_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// Ingredient -
type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type Company struct {
	Total int          `json:"total"`
	Rows  []CompanyRow `json:"rows"`
}

type CompanyRow struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Image            string `json:"image"`
	AssetsCount      int    `json:"assets_count"`
	LicensesCount    int    `json:"licenses_count"`
	AccessoriesCount int    `json:"accessories_count"`
	ConsumablesCount int    `json:"consumables_count"`
	ComponentsCount  int    `json:"components_count"`
	UsersCount       int    `json:"users_count"`
}

type Consumable struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Image    string   `json:"image"`
	Category Category `json:"category"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
