# CSV to PocketBase CLI Tool

Este es un programa de línea de comandos (CLI) diseñado para cargar datos desde un archivo CSV a PocketBase, un servicio de base de datos en la nube. La herramienta utiliza la API de PocketBase para enviar datos estructurados desde un archivo CSV a una colección específica en tu instancia de PocketBase.

## Instalación

1. Asegúrate de tener Go instalado en tu sistema. Si no lo tienes instalado, puedes descargarlo desde [https://golang.org/dl/](https://golang.org/dl/).

2. Instala el programa utilizando el siguiente comando:
   
   ```bash
   go install asdrome.com/csv-to-pb/cmd/csv-to-pb
   ```

## Uso

Una vez instalado, puedes usar el programa de la siguiente manera:

```bash
csv-to-pb --collection NOMBRE_DE_LA_COLECCION --endpoint ENDPOINT_DE_POCKETBASE ARCHIVO_CSV
```

- `NOMBRE_DE_LA_COLECCION`: El nombre de la colección en PocketBase donde deseas cargar los datos.
- `ENDPOINT_DE_POCKETBASE`: El endpoint de PocketBase donde se encuentra tu instancia (por ejemplo, `https://asdrome-pos-pb.pockethost.io`).
- `ARCHIVO_CSV`: La ruta al archivo CSV que contiene los datos que deseas cargar.

Ejemplo:

```bash
csv-to-pb --collection perfumeria --endpoint https://asdrome-pos-pb.pockethost.io ./datos.csv
```

### Opciones Adicionales

- `--help`: Muestra información de ayuda sobre el uso del programa.

## Requisitos

- Tener una cuenta en PocketBase.
- Tener un archivo CSV con datos que deseas cargar.
- Conexión a Internet para acceder a la API de PocketBase.


## Licencia

Este programa está bajo la Licencia MIT - consulta el archivo `LICENSE` para más detalles.
