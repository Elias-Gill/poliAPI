#!/bin/bash
# Este script se encarga de buscar los archivos excel disponibles en la web de la POLI,
# los lista y guarda los links de descarga disponibles para descarga en un archivo que se llama
# "Disponibles". 
# INFO: solo lista los links, queda a total responsabilidad de la base de datos y del backend de
# asegurarse de que estos archivos no estan repetidos para evitar dobles parseos.

# WARN: crear la variable de entorno $CACHED_SHEETS
# Revisar que la variable exista

if [[ -z "${CACHED_SHEETS}" ]]; then
    exit 1;
fi

# asegurarse que exista la carpeta
if [ ! -d "$CACHED_SHEETS" ]; then
    mkdir "$CACHED_SHEETS"
fi

old_list=$(cat "$CACHED_SHEETS/disponible")

# Listar los archivos disponibles en la web
new_list=$(curl -s --basic --insecure https://www.pol.una.py/academico/horarios-de-clases-y-examenes/ | grep -o '<a[^>]*href="[^"]*"' $1 | sed 's/<a[^>]*href="//' | awk -F '"' '!/^[[:space:]]/ && !/^#/ {print $1}' | grep -i "xls" | grep -i "clases")

# "Limpiar" la chache de archivos antiguos
for VAR in $old_list
do
    if ! grep -q "$VAR" <<< "$new_list"; then
        name=$(basename $VAR)
        name="${name%.*}" # remove extension
        rm "$CACHED_SHEETS/json/$name.json"
    fi
done

# actualizar el archivo de excels disponibles
rm "$CACHED_SHEETS/disponible"
$new_list >> "$CACHED_SHEETS/disponible"
