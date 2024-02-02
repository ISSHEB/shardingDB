<h2>ShardinngDB</h2>

В этом примере мы создаем две базы данных (db1 и db2) с помощью пакета github.com/lib/pq. 
Затем мы определяем, какая из них имеет наименьший размер, используя функцию getDatabaseWithLessMemory.

    func getDatabaseSize(db *sql.DB, dbName string) (int64, error) {
    	var size int64
    	err := db.QueryRow("SELECT pg_database_size($1)", dbName).Scan(&size)
    	if err != nil {
    		return 0, err
    	}
    	return size, nil
    }

Функция getDatabaseSize получает размер базы данных, используя запрос SELECT pg_database_size($1). 
Она возвращает размер в байтах и ошибку, если возникла.

    func getDatabaseSize(db *sql.DB, dbName string) (int64, error) {
    var size int64
    err := db.QueryRow("SELECT pg_database_size($1)",
    dbName).Scan(&size) 
    if err != nil { 
    return 0, err 
    }
    return size, nil 
    }

Функция getDatabaseWithLessMemory сравнивает размеры двух баз данных и возвращает ссылку на базу данных с наименьшим размером. 
Если возникает ошибка, она также возвращается.

    getDatabaseWithLessMemory(db1 *Database, db2 *Database) (*Database, error) {
    size1, err := getDatabaseSize(db1.DB, "db1") 
    if err != nil {
    return nil, err 
    } 
    size2, err := getDatabaseSize(db2.DB, "db2") 
    if err != nil {
    return nil, err 
    }
    if size1 < size2 {
    	return db1, nil
    }
    return db2, nil
    }

Функция SendToDB вставляет данные в таблицу базы данных. 
Она принимает указатель на объект sql.DB и возвращает ошибку, если возникает.

    func SendToDB(db *sql.DB) error { 
    query := "INSERT INTO my_table (column1, column2) VALUES ($1, $2)" 
    args := []interface{}{"value1", "value2"}
    _, err := db.Exec(query, args...)
    if err != nil {
    	return err
    }
    return nil
    }

В функции main мы создаем две базы данных и определяем, какая из них имеет наименьший размер. 
Затем мы отправляем данные в эту базу данных с помощью функции SendToDB.

    func main() {
    _ = pq.Driver{}
    connStr1 := "host=postgres_db1 user=postgres password=postgres dbname=db1 port=5432 sslmode=disable"
    db1, err := NewDatabase("db1", connStr1)
    if err != nil {
    	log.Fatal(err)
    }
    
    connStr2 := "host=postgres_db2 user=postgres password=postgres dbname=db2 port=5432 sslmode=disable"
    db2, err := NewDatabase("db2", connStr2)
    if err != nil {
    	log.Fatal(err)
    }
    
    leastUsedDB, err := getDatabaseWithLessMemory(db1, db2)
    if err != nil {
    	log.Fatal(err)
    }
    
    s := SendToDB(leastUsedDB.DB)
    
    fmt.Printf("Data %s sent successfully.", s)
    }
