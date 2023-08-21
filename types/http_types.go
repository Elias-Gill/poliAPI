package types

type ExcelInfo struct {
	Id     int    `json:"id"`
	Fecha  string `json:"date"`
	Nombre string `json:"fileName"`
}

type JWTResponse struct {
	Token string `json:"token"`
}

type HttpError struct {
	Error string `json:"error"`
}
