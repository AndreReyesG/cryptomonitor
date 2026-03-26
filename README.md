# CryptoMonitor
**CryptoMonitor** es una aplicación web desarrollada en
**Golang** que permite monitorear el precio de criptomonedas
en tiempo real.
El sistema obtiene los precios desde APIs públicas y
los muestra en un dashboard web simple, claro y fácil de usar.

Actualmente el proyecto permite visualizar los precios de:
- Bitcoin (BTC)
- Ethereum (ETH)

## Descripción del proyecto
CryptoMonitor fue desarrollado como parte de un proyecto académico
enfocado en el desarrollo de software utilizando arquitectura modular y 
buenas prácticas en Go.

El sistema está dividido en dos partes principales:

- Un **API Server** encargado de obtener los precios desde una API externa
- Un **Web Server** encargado de mostrar la información en un dashboard web

Esto permite que el sistema sea más escalable, mantenible y fácil de mejorar en el futuro

## Tecnologías utilizadas
Las principales tecnologías utilizadas en el proyecto son:
- Golang
- HTML
- CSS
- Templates del lado del servidor con Go
- API REST
- Consumo de APIs externas (CoinGecko)

## Features
El proyecto incluye las siguientes funcionalidades:
- Obtener precios de criptomonedas desde una API externa
- Exponer los precios mediante una API REST
- Mostrar los precios en un dashboard web
- Actualización manual de precios mediante un botón
- Renderizado HTML desde el servidor

## El proceso de desarrollo
El proyecto fue desarrollado utilizando una metodología ágil basada en historias de usuario y sprints.
Durante el desarrollo se implementaron las siguientes funcionalidades principales:
- Obtener precios desde CoinGecko
- Crear un endpoint REST para exponer los precios
- Construir un dashboard web para visualizar los datos
- Implementar la actualización manual de precios

## Cómo ejecutar el proyecto
Para ejecutar el proyecto se utiliza un Makefile, el cual permite iniciar el
servidor web, la API REST y ejecutar diferentes pruebas de forma rápida.

### 1. Clonar repositorio
```bash
git clone https://github.com/AndreReyesG/cryptomonitor.git
cd cryptomonitor
```
### 2. Ejecutar el proyecto usando Makefile
#### Ejecutar primero la API REST
```bash
make run/api
```
#### Después ejecutar el servidor web
```bash
make run/web
```
### 3. Abrir el proyecto en el navegador
Después de ejecutar el servidor web, abre tu navegador y entra a:
```
http://localhost:9000
```

## Pruebas y comandos del Makefile
El proyecto incluye un Makefile que permite ejecutar el sistema y
realizar pruebas de forma rápida y organizada.
### Ejecutar todas las pruebas
```bash
make test
```
### Ejecutar pruebas por módulo
```bash
make test/api
make test/web
make test/exchanges
make test/ui
```
### Otros comandos útiles
```bash
make tidy
```
Este comando:
- Limpia dependencias del proyecto
- Formatea todos los archivos .go
- Mantiene el código organizado

## Estructura del proyecto
```
cryptomonitor
│
├── cmd
│   ├── api
│   └── webserver
│
├── internal
│   ├── api
│   ├── domain
│   ├── exchanges
│   ├── platform
│   └── web
│
└── ui
```
