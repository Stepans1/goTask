package DTO

type PaternDescription struct {
	Key         string `json:"key"`
	Description string `json:"description"`
}

type GenerationOptions struct {
	Length  int      `json:"length"`
	Options []string `json:"options"`
}

type CharSet struct {
	Key string
	Set string
}