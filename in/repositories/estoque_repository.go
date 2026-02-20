package repositories

import (
	"github.com/ManelFer/APIgo/in/config/database"
	"github.com/ManelFer/APIgo/in/models"
)

func CreateEstoque(e models.Estoque) error {
	query := `INSERT INTO estoque (equipamento, marca, modelo, patrimonio, quantidade)
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := database.DB.Exec(query, e.Equipamento, e.Marca, e.Modelo, e.Patrimonio, e.Quantidade)
	return err
}

func GetAllEstoque() ([]models.Estoque, error) {
	rows, err := database.DB.Query("SELECT id, equipamento, marca, modelo, patrimonio, quantidade FROM estoque")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Estoque
	for rows.Next() {
		var e models.Estoque
		if err := rows.Scan(&e.ID, &e.Equipamento, &e.Marca, &e.Modelo, &e.Patrimonio, &e.Quantidade); err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}
