package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)


func main() {
	// Получаем номера заказов из аргументов командной строки
	args := os.Args[1:] // Получаем аргументы командной строки
	if len(args) == 0 { // Если аргументов нет, то выходим с ошибкой
		log.Fatal("Необходимо указать номер заказа")
	}

	// Подключение к базе данных
	db, err := sql.Open("postgres", "user=postgres password=newPassword dbname=store sslmode=disable")
	if err != nil { // Если не удалось подключиться к БД, выходим с ошибкой
		log.Fatal("не удалось подключиться к БД")
	}

	// Формирование SQL-запроса для выборки данных из таблиц БД
	query := fmt.Sprintf(`
		SELECT 
			shelves.name, products.name, products.id, order_items.quantity, orders.order_number
		FROM
			order_items
		INNER JOIN 
			orders ON order_items.order_id = orders.id
		INNER JOIN 
			products ON order_items.product_id = products.id
		INNER JOIN 
			shelves ON order_items.shelf_id = shelves.id
		WHERE 
			orders.order_number IN ('%s')
	`, strings.Join(args, "','"))

	// Выполнение запроса и обработка результатов
	fmt.Println("=+=+=+=")
	fmt.Printf("Страница сборки заказов %s\n", strings.Join(args, ","))
	fmt.Println()

	result := make(map[string]map[string]int)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("не удалось выполнить запрос и обработать результат")
	}
	for rows.Next() {
		var shelfName, productName, productID, orderNum string
		var quantity int
		if err := rows.Scan(&shelfName, &productName, &productID, &quantity, &orderNum); err != nil {
			log.Fatal("не удалось сканировать строки")
		}

		if _, ok := result[shelfName]; !ok {
			result[shelfName] = make(map[string]int)
			fmt.Printf("===Стеллаж %s\n", shelfName)
		}

		fmt.Printf("%s (id=%s) \nзаказ %s, %d шт\n", productName, productID, orderNum, quantity)
		fmt.Println()

		result[shelfName][productName] += quantity
	}
}