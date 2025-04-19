package product

import (
	"database/sql"
	"fmt"
	"log"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]Product, error) {
	query := "SELECT id, name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var productList []Product
	var productObj Product

	for rows.Next() {
		if err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price); err != nil {
			fmt.Println(err)
			return nil, err
		}

		productList = append(productList, productObj)
	}

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product Product) (int, error) {

	stmt, err := pr.connection.Prepare("INSERT INTO product(name, price) VALUES (?, ?);")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Price)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}

func (pr *ProductRepository) GetProductById(id int) (*Product, error) {
	stmt, err := pr.connection.Prepare("SELECT * FROM product WHERE id = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil
}
