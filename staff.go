package main 

// Вывод результата
fmt.Printf("Страница сборки заказа %d\n", orderID)
fmt.Println("=+=+=+=")

for shelfName, products := range result {
    fmt.Printf("===Стеллаж %s\n", shelfName)
    for productName, quantity := range products {
        fmt.Printf("%s (количество: %d)\n", productName, quantity)
    }
}

// --------------------------
fmt.Println("Order details:")
fmt.Println("--------------")
for rows.Next() {
	err := rows.Scan(&orderNumber, &productName, &quantity, &shelfName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Order #%d: %d x %s on shelf %s\n", orderNumber, quantity, productName, shelfName)
}
if rows.Err() != nil {
	panic(rows.Err())
}
}

func main_() {
	// Получаем номера заказов из аргументов командной строки
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Необходимо указать номер заказа")
	}

	// Подключение к базе данных
	db, err := sql.Open("postgres", "user=postgres password=newPassword dbname=store sslmode=disable")
	if err != nil {
		log.Fatal("не удалось подключиться к БД")
	}

	// Формирование SQL-запроса для выборки данных из таблиц БД
	queries := make(map[string]string)
	for _, orderNum := range args {
		query := fmt.Sprintf(`
			SELECT 
				shelves.name, products.name, products.id, order_items.quantity
			FROM
				order_items
			INNER JOIN 
				orders ON order_items.order_id = orders.id
			INNER JOIN 
				products ON order_items.product_id = products.id
			INNER JOIN 
				shelves ON order_items.shelf_id = shelves.id
			WHERE 
				orders.order_number = '%s'
		`, orderNum)
		queries[orderNum] = query
	}

	// Выполнение запросов и обработка результатов
	fmt.Println("=+=+=+=")
	fmt.Printf("Страница сборки заказов %s\n", strings.Trim(fmt.Sprint(args), "[]"))
	fmt.Println()

	result := make(map[string]map[string]int)
	for _, query := range queries {
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal("не удалось выполнить запрос и обработать результат")
		}

		for rows.Next() {
			var shelfName, productName, productID string
			var quantity int
			if err := rows.Scan(&shelfName, &productName, &productID, &quantity, ); err != nil {
				log.Fatal("не удалось сканировать строки")
			}

			if _, ok := result[shelfName]; !ok {
				result[shelfName] = make(map[string]int)
				fmt.Printf("===Стеллаж %s\n", shelfName)
			}

			fmt.Printf("%s (id=%s) \nзаказ %s, %d шт\n", productName, productID, args[0], quantity)
			fmt.Println()
			if _, ok := result[shelfName]; !ok {
				result[shelfName] = make(map[string]int)
			}
			result[shelfName][productName] += quantity
		}
	}
	// // fmt.Println("=+=+=+=")
	// // fmt.Println("Итоговая статистика:")
	// for shelfName, products := range result {
	// 	fmt.Printf("===Стеллаж %s\n", shelfName)
	// 	for productName, quantity := range products {
	// 		fmt.Printf("%s: %d шт\n", productName, quantity)
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("=+=+=+=")
	// fmt.Println("Информация о товарах на каждом стеллаже:")
	// fmt.Println()
	// for shelfName, products := range result {
	// 	fmt.Printf("Стеллаж %s\n", shelfName)
	// 	for productName, quantity := range products {
	// 		fmt.Printf("%s: %d шт\n", productName, quantity)
	// 	}
	// 	fmt.Println()
	// }

}

func main() {
	// Получаем номера заказов из аргументов командной строки
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Необходимо указать номер заказа")
	}

	// Подключение к базе данных
	db, err := sql.Open("postgres", "user=postgres password=newPassword dbname=store sslmode=disable")
	if err != nil {
		log.Fatal("не удалось подключиться к БД")
	}

	// Формирование SQL-запроса для выборки данных из таблиц БД
	queries := make(map[string]string)
	for _, orderNum := range args {
		query := fmt.Sprintf(`
			SELECT 
				shelves.name, products.name, products.id, order_items.quantity
			FROM
				order_items
			INNER JOIN 
				orders ON order_items.order_id = orders.id
			INNER JOIN 
				products ON order_items.product_id = products.id
			INNER JOIN 
				shelves ON order_items.shelf_id = shelves.id
			WHERE 
				orders.order_number = '%s'
		`, orderNum)
		queries[orderNum] = query
	}

	// Выполнение запросов и обработка результатов
	fmt.Println("=+=+=+=")
	fmt.Printf("Страница сборки заказов %s\n", strings.Join(args, ","))
	fmt.Println()

	result := make(map[string]map[string]int)
	for orderNum, query := range queries {
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal("не удалось выполнить запрос и обработать результат")
		}

		for rows.Next() {
			var shelfName, productName, productID string
			var quantity int
			if err := rows.Scan(&shelfName, &productName, &productID, &quantity); err != nil {
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

	// // Поиск дополнительных стеллажей
	// for shelfName, products := range result {
	// 	fmt.Printf("доп стеллаж %s: ", shelfName)
	// 	var additionalShelves []string
	// 	for productName, quantity := range products {
	// 		if quantity > 5 {
	// 			additionalShelves = append(additionalShelves, productName)
	// 		}
	// 	}
	// 	if len(additionalShelves) == 0 {
	// 		fmt.Println("нет дополнительных стеллажей")
	// 	} else {
	// 		fmt.Println(strings.Join(additionalShelves, ","))
	// 	}
	// }
}

func main() {
	// Получаем номера заказов из аргументов командной строки
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Необходимо указать номер заказа")
	}

	// Подключение к базе данных
	db, err := sql.Open("postgres", "user=postgres password=newPassword dbname=store sslmode=disable")
	if err != nil {
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