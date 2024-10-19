package structs

type Service struct {
	Name 	 string `json:"name"`
	Dockerfile string `json:"dockerfile"`
	DockerId string `json:"docker_id"`
	Codelang string `json:"codelang"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}