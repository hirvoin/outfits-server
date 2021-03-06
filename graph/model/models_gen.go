// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Garment struct {
	ID         string `json:"id"`
	User       *User  `json:"user"`
	Title      string `json:"title"`
	Category   string `json:"category"`
	Color      string `json:"color"`
	WearCount  int    `json:"wearCount"`
	IsFavorite bool   `json:"isFavorite"`
	ImageURI   string `json:"imageUri"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewGarment struct {
	UserID   string `json:"userId"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Color    string `json:"color"`
	ImageURI string `json:"imageUri"`
}

type NewOutfit struct {
	UserID   string   `json:"userId"`
	Garments []string `json:"garments"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Outfit struct {
	ID       string     `json:"id"`
	User     *User      `json:"user"`
	Date     string     `json:"date"`
	Garments []*Garment `json:"garments"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type UpdatedGarment struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Color    string `json:"color"`
	ImageURI string `json:"imageUri"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
