# gitops-echobot

Aplicación Golang utilizada en los ejercicios de la serie [GitOps Flux](https://github.com/Sngular/gitops-flux-series).

## Descripción

En función del valor de la variable `OUTPUT_TYPE`:

- `OUTPUT_TYPE=stdout`: imprime trazas en la consola con cierta frecuencia.


- `OUTPUT_TYPE=mongodb`: inserta trazas de log en una base de datos [MongoDB](https://www.mongodb.com/) con cierta frecuencia e imprime mensajes en la consola sobre con información de la traza.

El mensaje de la traza y la frecuencia de tiempo pueden ser modificados a través de las variables de entorno `MESSAGE` y `SLEEP_TIME` respectivamente.

La base de datos se configura mediante las variables de entorno `MONGODB_URI`, `MONGODB_DATABASE` y `MONGODB_COLLECTION`.

## Variables de entorno

El funcionamiento de la aplicación puede ser modificado a través de variables de entorno:

| Variable de entorno  | Descripción                                                  | Requerida | Valor por defecto     |
|----------------------|--------------------------------------------------------------|-----------|-----------------------|
| `OUTPUT_TYPE`        | Define el tipo de salida del servicio (`stdout`, `mongodb`). | No        | "stdout"              |
| `MESSAGE`            | Modifica el mensaje de impreso en la pantalla.               | No        | "Gitops Flux series!" |
| `SLEEP_TIME`         | Modifica el intervalo de tiempo entre mensajes.              | No        | 5s                    |
| `MONGODB_URI`        | MongoDB uri `mongodb://user:pass@host:port`.                 | No        |                       |
| `MONGODB_DATABASE`   | Nombre de la base de datos donde se insertarán los datos.    | No        | echobot               |
| `MONGODB_COLLECTION` | Nombre de la `collection` donde se insertarán los datos.     | No        | log                   |

## Funcionamiento

Para ver su funcionamiento utilice el siguiente comando:

```bash
docker container run --rm ghcr.io/sngular/gitops-echobot:v0.2.2

hostname: 86f43c549bd4 - Gitops Flux series!
hostname: 86f43c549bd4 - Gitops Flux series!
hostname: 86f43c549bd4 - Gitops Flux series!
```

Si se desean almacenar las trazas en una base de datos MongoDB:

```bash
docker container run --rm \
  --env OUTPUT_TYPE="mongodb" \
  --env SLEEP_TIME="5s" \
  --env MESSAGE="Sngular utiliza Gitops en sus entornos" \
  --env MONGODB_URI="mongodb://test:1234@localhost:27017" \
  --env MONGODB_DATABASE="echobot" \
  --env MONGODB_COLLECTION="log" \
  ghcr.io/sngular/gitops-echobot:v0.2.2

Traza insertada en base de datos (507f1f77bcf86cd799439011): 680c3892c04c - Sngular utiliza Gitops en sus entornos
Traza insertada en base de datos (757f191a810c19729de860ae): 680c3892c04c - Sngular utiliza Gitops en sus entornos
Traza insertada en base de datos (807f191a810c19729de860ae): 680c3892c04c - Sngular utiliza Gitops en sus entornos
```
