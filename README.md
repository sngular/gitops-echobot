# gitops-echobot

Aplicación Golang utilizada en los ejercicios de la serie [GitOps Flux](https://github.com/Sngular/gitops-flux-series).

## Descripción

La aplicación imprime mensajes en la consola con cierta frecuencia. El mensaje a imprimir y la frecuencia de tiempo pueden ser modificados a través de las variables de entorno `CHARACTER` y `SLEEP_TIME` respectivamente.

## Funcionamiento

Para ver su funcionamiento utilice el siguiente comando:

```bash
docker container run --rm ghcr.io/sngular/gitops-echobot:v0.1.0

hostname: 86f43c549bd4 - gitops flux series
hostname: 86f43c549bd4 - gitops flux series
hostname: 86f43c549bd4 - gitops flux series
```

## Variables de entorno

El funcionamiento de la aplicación puede ser modificado a través de variables de entorno:

| Variable de entorno | Descripción                                     | Valor por defecto     |
|---------------------|-------------------------------------------------|-----------------------|
| `CHARACTER`         | Modifica el mensaje de impreso en la pantalla.  | "gitops flux series" |
| `SLEEP_TIME`        | Modifica el intervalo de tiempo entre mensajes. | 1s                    |

```bash
docker container run --rm \
  --env SLEEP_TIME="3s" \
  --env CHARACTER="Sngular utiliza gitops en sus entornos" \
  ghcr.io/sngular/gitops-echobot:v0.1.0

hostname: 680c3892c04c - Sngular utiliza gitops en sus entornos
hostname: 680c3892c04c - Sngular utiliza gitops en sus entornos
hostname: 680c3892c04c - Sngular utiliza gitops en sus entornos
```
