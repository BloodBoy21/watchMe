package shared

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb"

	"watch-me/structs"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	dbPath := "/tmp/watch-me/data.ddb"
	db, err := sql.Open("duckdb", dbPath)
	if err != nil {
		panic(err)
	}
	createServicesTable(db)
	return db
}

func createServicesTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS services (docker_id TEXT,name TEXT,dockerfile TEXT,codelang TEXT,created_at DATE,updated_at DATE)")
	if err != nil {
		panic(err)
	}
}

func SaveService(service *structs.Service) {
	db := GetDB()

	_, err := db.Exec(`INSERT INTO services (name, dockerfile, codelang, created_at, updated_at,docker_id) VALUES (?,?,?,?,?,?)`, service.Name, service.Dockerfile, service.Codelang, service.CreatedAt, service.UpdatedAt, service.DockerId)
	if err != nil {
		panic(err)
	}
}

func GetAllServices() []*structs.Service {
	db := GetDB()
	// Ensure the column order in the SELECT statement matches the Scan order
	rows, err := db.Query(`SELECT docker_id, name, dockerfile, codelang, created_at, updated_at FROM services`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var services []*structs.Service
	for rows.Next() {
		service := &structs.Service{}
		// Make sure the Scan order matches the SELECT column order
		err := rows.Scan(&service.DockerId, &service.Name, &service.Dockerfile, &service.Codelang, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			panic(err)
		}
		services = append(services, service)
	}


	return services
}
