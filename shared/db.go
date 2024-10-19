package shared

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb"

	"watch-me/structs"
)


func GetDB() *sql.DB {
	//dbPath := "~/.config/watch-me/data.duckdb"
	db, err := sql.Open("duckdb", "")
	if err != nil {
		panic(err)
	}
	createServicesTable(db)
	return db
}

func createServicesTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS services (service_id INTEGER PRIMARY KEY,docker_id TEXT,name TEXT,dockerfile TEXT,codelang TEXT,created_at DATE,updated_at DATE)")
	if err != nil {
		panic(err)
	}
}

func SaveService(service *structs.Service) {
	db := GetDB()

	_, err := db.Exec(`INSERT INTO services (name, dockerfile, codelang, created_at, updated_at) VALUES (?,?,?,?,?)`, service.Name, service.Dockerfile, service.Codelang, service.CreatedAt, service.UpdatedAt)
	if err != nil {
		panic(err)
	}
}

func GetAllServices() []*structs.Service {
	db := GetDB()
	rows, err := db.Query(`SELECT * FROM services`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var services []*structs.Service
	for rows.Next() {
		service := &structs.Service{}
		err := rows.Scan(&service.Name, &service.Dockerfile, &service.Codelang, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			panic(err)
		}
		services = append(services, service)
	}
	return services
}
