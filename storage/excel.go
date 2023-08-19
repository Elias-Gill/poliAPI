package storage

import "github.com/elias-gill/poliapi/types"

// Lista los 5 archivos excel mas recientes en la base de datos
func GetAvailableExeclFiles() ([]*types.ExcelInfo, error) {
    return nil, nil
}

// lista las materias disponibles dentro de un archivo excel
func GetSubjectsFromExcel(fileId int) ([]parser.Materia, error) {
    return nil, nil
}

