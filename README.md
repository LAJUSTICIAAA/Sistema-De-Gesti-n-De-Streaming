
#  Sistema de Gestión de Streaming en Go

**Autor:** Justin Salazar  
**Carrera:** Ingeniería en Software  
**Correo:** jusalazarro@uide.edu.ec  
**Fecha:** Junio 2025  

---

##  Descripción General

Este sistema de streaming fue desarrollado en el lenguaje de programación **Go (Golang)** como parte del proyecto final de la materia **Programación Orientada a Objetos**. Permite realizar acciones como:

- Registro e inicio de sesión de usuarios.
- Subida y listado de videos simulados.
- Subida real de archivos de video.
- Listado y búsqueda por título.
- Eliminación de videos.

Todo está organizado con una estructura modular, servicios REST y buenas prácticas de programación.

---

##  Estructura del Proyecto

```
 Sistema-De-Gesti-n-De-Streaming/
├── handlers/             # Rutas y lógica HTTP
│   ├── routes.go
│   ├── user_handlers.go
│   └── video_handlers.go
│
├── models/               # Estructuras de datos
│   ├── user.go
│   └── video.go
│
├── services/             # Servicios de negocio
│   └── video_service.go
│
├── utils/                # Funciones de utilidad
│   └── hash.go
│
├── main.go               # Punto de entrada
├── go.mod                # Configuración del módulo Go
├── go.sum
└── README.md             # Este archivo
```

---

##  Tecnologías Usadas

- **Go 1.24.3**
- **Gin Gonic** - Framework HTTP para APIs
- **bcrypt** - Hashing de contraseñas
- **Postman** - Para pruebas
- **HTML5** - Para visualización simulada de video

---

##  Funcionalidades REST

| Método | Endpoint            | Descripción                             |
|--------|---------------------|-----------------------------------------|
| POST   | `/register`         | Registrar un nuevo usuario              |
| POST   | `/login`            | Iniciar sesión                          |
| GET    | `/videos`           | Listar todos los videos simulados       |
| POST   | `/videos`           | Subir video simulado (JSON)             |
| POST   | `/upload`           | Subir video real (archivo `form-data`)  |
| GET    | `/videos/:title`    | Obtener video por título                |
| DELETE | `/videos/:title`    | Eliminar video por título               |
| GET    | `/users`            | Listar todos los usuarios registrados   |

---

##  Cómo Probar en Postman

### 1. Registrar usuario
**POST** `http://localhost:8080/register`  
```json
{
  "nombre": "Justin",
  "email": "justin@correo.com",
  "password": "123456"
}
```

### 2. Iniciar sesión  
**POST** `http://localhost:8080/login`  
```json
{
  "username": "Justin",
  "password": "123456"
}
```

### 3. Subir video simulado  
**POST** `http://localhost:8080/videos`  
```json
{
  "title": "Mi primer video",
  "url": "http://localhost:8080/static/demo1.mp4"
}
```

### 4. Subir video real  
**POST** `http://localhost:8080/upload`  
- Tipo: `form-data`
- Clave: `video`  
- Valor: Selecciona un archivo `.mp4` desde tu PC

### 5. Ver lista de videos  
**GET** `http://localhost:8080/videos`

### 6. Buscar video por título  
**GET** `http://localhost:8080/videos/Mi primer video`

### 7. Eliminar video por título  
**DELETE** `http://localhost:8080/videos/Mi primer video`

### 8. Ver usuarios registrados  
**GET** `http://localhost:8080/users`

---

##  Cómo Ejecutar el Proyecto

### 1. Clona el repositorio
```bash
git clone https://github.com/LAJUSTICIAAA/Sistema-De-Gesti-n-De-Streaming.git
cd Sistema-De-Gesti-n-De-Streaming
```

### 2. Instala las dependencias
```bash
go mod tidy
```

### 3. Ejecuta el servidor
```bash
go run main.go
```

### 4. Abre Postman y prueba todos los endpoints listados arriba.

---

##  Contenidos Aplicados

- ✔️ Encapsulamiento y estructuras (`struct`)
- ✔️ Interfaces y modularización
- ✔️ JSON y servicios REST
- ✔️ Buen manejo de errores
- ✔️ Hashing de contraseñas con bcrypt

---

##  Reflexión Final

Este sistema me permitió aplicar de manera práctica todos los conceptos aprendidos en la materia de POO, estructurar un backend profesional en Go, trabajar con rutas RESTful y asegurar la seguridad básica del sistema. Me ayudó a mejorar mi lógica de programación y el uso de herramientas como Postman y GitHub.

---

##  Licencia

Desarrollado con fines académicos en la Universidad Internacional del Ecuador (UIDE).
