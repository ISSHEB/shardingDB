<h2>ShardinngDB</h2>

В этом примере мы создаем две базы данных (db1 и db2) с помощью пакета github.com/lib/pq. 
Затем мы определяем, какая из них имеет наименьший размер, используя функцию getDatabaseWithLessMemory.

Функция getDatabaseSize получает размер базы данных, используя запрос SELECT pg_database_size($1). 
Она возвращает размер в байтах и ошибку, если возникла.

Функция getDatabaseWithLessMemory сравнивает размеры двух баз данных и возвращает ссылку на базу данных с наименьшим размером. 
Если возникает ошибка, она также возвращается.

Функция SendToDB вставляет данные в таблицу базы данных. 
Она принимает указатель на объект sql.DB и возвращает ошибку, если возникает.

В функции main мы создаем две базы данных и определяем, какая из них имеет наименьший размер. 
Затем мы отправляем данные в эту базу данных с помощью функции SendToDB.