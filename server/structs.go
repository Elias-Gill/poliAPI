package server

type ExcelInfo struct {
	Id     int    `json:"id"`
	Fecha  string `json:"date"`
	Nombre string `json:"fileName"`
}

type User struct {
	Name  *string `json:"userName" db:"nombre"`
	Pasw  *string `json:"password" db:"password"`
	Email *string `json:"email" db:"email"`
}
