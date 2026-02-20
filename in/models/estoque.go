package models

type Estoque struct {
	ID          int    `json:"id"`
	Equipamento string `json:"equipamento"`
	Marca       string `json:"marca"`
	Modelo      string `json:"modelo"`
	Patrimonio  string `json:"patrimonio"`
	Quantidade  int    `json:"quantidade"`
}
