package main


type Profile struct {
	Name string `json:"name"`
}

type User struct {
	Id uint64 `json:"id"`
	Profile Profile `json:"Profile"`
}