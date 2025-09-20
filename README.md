# Go + React Inventory

Un **mini sistema de inventario** full‑stack para practicar CRUD end‑to‑end con **Go (Gin)**, **MySQL** y **React + Vite**. Incluye una API REST sencilla y un frontend que consume los productos y los muestra en una tabla.

---

## ✨ Características

- **API REST** con Gin: `GET /products`, `POST /products`, `PUT /products/:id`, `DELETE /products/:id`  
- **Persistencia** en **MySQL** (tabla `products`)  
- **Frontend React + Vite** que consume la API y lista productos  
- **CORS** configurado para desarrollo (`http://localhost:5173`)  
- **ESLint** ya integrado en el frontend para mantener calidad de código

---

## 🧱 Arquitectura

```
┌───────────────┐   HTTP/JSON (fetch)   ┌─────────────────┐      SQL/TCP      ┌───────────┐
│ React (Vite)  │  ───────────────────▶ │ Go API (Gin)    │  ───────────────▶ │ MySQL DB  │
│ http://5173   │   ◀────────────────── │ :8080 /products │   ◀────────────── │ localhost │
└───────────────┘         CORS          └─────────────────┘                   └───────────┘
```

---

## 📁 Estructura del repo

```
timosboy-go-react-inventory/
├── LICENSE
├── backend/
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── frontend/
    ├── README.md
    ├── eslint.config.js
    ├── index.html
    ├── package.json
    ├── vite.config.js
    └── src/
        ├── App.css
        ├── App.jsx
        ├── index.css
        └── main.jsx
```

---

## 🔧 Requisitos

- **Go** 1.25+  
- **Node.js** 20+ y **npm**  
- **MySQL** 8.x

---

## 🚀 Puesta en marcha

### 1) Base de datos

Crea la BD y la tabla `products`:

```sql
CREATE DATABASE IF NOT EXISTS inventory_db;
USE inventory_db;

CREATE TABLE IF NOT EXISTS products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  quantity INT NOT NULL DEFAULT 0
);

-- Datos de ejemplo
INSERT INTO products (name, price, quantity) VALUES
('Mouse', 29.99, 50),
('Teclado', 49.99, 20);
```

### 2) Backend (Go + Gin)

> Por defecto el backend escucha en `localhost:8080` y se conecta a MySQL en `127.0.0.1:3306` con la BD `inventory_db`.

- Entrar a la carpeta y ejecutar:

```bash
cd backend
go run main.go
```

> 💡 **Sugerencia** (producción): mover el DSN a una variable de entorno como `DB_DSN` y leerla en el código (en vez de tenerla hardcodeada). También parametriza `CLIENT_ORIGIN` y el puerto del servidor.

### 3) Frontend (React + Vite)

- Instalar dependencias y ejecutar en modo desarrollo:

```bash
cd frontend
npm install
npm run dev
```

- Vite abrirá el sitio en `http://localhost:5173` y hará **fetch** a `http://localhost:8080/products` (coincide con el CORS del backend).

---

## 📚 API Reference

**Base URL:** `http://localhost:8080`

### Obtener productos
```http
GET /products
```

### Crear producto
```http
POST /products
Content-Type: application/json

{
  "name": "Laptop",
  "price": 999.99,
  "quantity": 5
}
```

### Actualizar producto
```http
PUT /products/:id
Content-Type: application/json

{
  "name": "Laptop",
  "price": 899.00,
  "quantity": 4
}
```

### Eliminar producto
```http
DELETE /products/:id
```

## 🧩 Frontend (UI)

- `src/App.jsx` hace `fetch` a `/products` y pinta una tabla con `id`, `name`, `price`, `quantity`.  
- Estilos base con `App.css` e `index.css`.  
- Scripts disponibles:
  - `npm run dev` — desarrollo
  - `npm run build` — build producción
  - `npm run preview` — previsualización
  - `npm run lint` — lint con ESLint

---

## 📄 Licencia

Este proyecto está bajo la licencia **MIT**. Consulta el archivo `LICENSE` para más detalles.

---

## 👤 Autor

- **Timothy Kuno** — 2025
