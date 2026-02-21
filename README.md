# 🔐 AuthService

AuthService is a centralized authentication and authorization microservice designed to work as an API Gateway component. It provides JWT-based authentication, role-based access control (RBAC), and permission management for distributed systems.

---

# 🚀 Features

* 🔐 JWT-based Authentication
* 📝 User Sign Up & Sign In
* 🔑 Role-Based Access Control (RBAC)
* 🧩 Permission Management
* 🔒 Password Hashing using bcrypt
* ✅ Request Validation via middlewares
* 🚦 Rate Limiting
* 🔁 Reverse Proxy / API Gateway support
* 🛡 Secure Middleware-based Authorization
* 💎 DB Migrations using goose

---

# 🗄 Database

**Database Name:** `airbnb_auth_dev`

## users

| Column    | Type                 |
| --------- | -------------------- |
| id        | INT                  |
| name      | String               |
| email     | String               |
| password  | String (bcrypt hash) |
| createdat | Timestamp            |
| updatedat | Timestamp            |

## roles

| Column    | Type      |
| --------- | --------- |
| id        | INT       |
| name      | String    |
| desc      | String    |
| createdat | Timestamp |
| updatedat | Timestamp |

## permissions

| Column    | Type      |
| --------- | --------- |
| id        | INT       |
| name      | String    |
| desc      | String    |
| resource  | String    |
| action    | String    |
| createdat | Timestamp |
| updatedat | Timestamp |

## users_roles

| Column    | Type      |
| --------- | --------- |
| id        | INT       |
| user_id   | INT       |
| role_id   | INT       |
| createdat | Timestamp |
| updatedat | Timestamp |

## roles_permissions

| Column        | Type      |
| ------------- | --------- |
| id            | INT       |
| role_id       | INT       |
| permission_id | INT       |
| createdat     | Timestamp |
| updatedat     | Timestamp |

---

# 🌐 Base URL

```
http://localhost:3004
```

---

# 👤 User Authentication & Management

---

## POST `/auth/signup`

**Description:** Register a new user
**Authorization:** Public

### Request Body

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret123"
}
```

**Required fields:** `name`, `email`, `password`

### Success Response — 200 OK

```json
{
  "message": "successfully created the user",
  "success": true,
  "data": "",
  "err": null
}
```

### Validation Error — 400

```json
{
  "message": "validation failed",
  "error": "validation error details",
  "data": null,
  "success": false
}
```

---

## POST `/auth/signin`

**Description:** User login (JWT issued)
**Authorization:** Public

### Request Body

```json
{
  "email": "john@example.com",
  "password": "secret123"
}
```

**Required fields:** `email`, `password`

### Success Response — 200 OK

```json
{
  "message": "successfully signed in",
  "success": true,
  "data": "<jwt-token>",
  "err": null
}
```

### Unauthorized — 401

```text
auth header required
```

---

## GET `/auth/user/{id}`

**Description:** Get user by ID
**Authorization:** User / Admin

### Success Response — 200 OK

```json
{
  "message": "user found",
  "success": true,
  "data": {
    "Id": 1,
    "Name": "John Doe",
    "Email": "john@example.com",
    "Password": "",
    "Created_at": "2026-01-01T00:00:00Z",
    "Updated_at": "2026-01-01T00:00:00Z"
  },
  "err": null
}
```

---

## GET `/auth/user/email/{email}`

**Description:** Get user by email
**Authorization:** User / Admin

### Success Response — 200 OK

```json
{
  "message": "user found",
  "success": true,
  "data": {
    "Id": 1,
    "Name": "John Doe",
    "Email": "john@example.com",
    "Password": "<hashed-password>",
    "Created_at": "2026-01-01T00:00:00Z",
    "Updated_at": "2026-01-01T00:00:00Z"
  },
  "err": null
}
```

---

## GET `/auth/users`

**Description:** Get all users
**Authorization:** User / Admin

### Success Response — 200 OK

```json
{
  "message": "users fetched successfully",
  "success": true,
  "data": [],
  "err": null
}
```

---

## DELETE `/auth/user/{id}`

**Description:** Delete user by ID
**Authorization:** Admin

### Success Response — 200 OK

```json
{
  "message": "user deleted successfully",
  "success": true,
  "data": "",
  "err": null
}
```

---

# 🧑‍💼 Role Management

---

## GET `/roles`

**Authorization:** Admin
**Description:** Get all roles

### Success Response — 200 OK

```json
{
  "message": "fetched all roles SUCCESSFULLY",
  "success": true,
  "data": [],
  "err": null
}
```

---

## POST `/roles`

**Authorization:** Admin
**Description:** Create a new role

### Request Body

```json
{
  "name": "manager",
  "description": "Manager role"
}
```

### Success Response — 202 Accepted

```json
{
  "message": "Successfully created the role",
  "success": true,
  "data": {},
  "err": null
}
```

---

## PATCH `/roles`

**Authorization:** Admin
**Description:** Update role

### Request Body

```json
{
  "id": "2",
  "name": "manager",
  "description": "Updated manager role"
}
```

### Success Response — 200 OK

```json
{
  "message": "Successfully updated the role",
  "success": true,
  "data": {},
  "err": null
}
```

---

## DELETE `/roles/id/{roleId}`

**Authorization:** Admin
**Description:** Delete role

### Success Response — 200 OK

```json
{
  "message": "Successfully deleted the role",
  "success": true,
  "data": null,
  "err": null
}
```

---

# 🔐 Permission Management

---

## POST `/permissions`

**Authorization:** Admin
**Description:** Create permission

### Request Body

```json
{
  "name": "booking:create",
  "desc": "Create booking",
  "resource": "booking",
  "action": "create"
}
```

### Success Response — 201 Created

```json
{
  "message": "permission created successfully",
  "success": true,
  "data": {},
  "err": null
}
```

---

## GET `/permissions`

**Authorization:** Admin
**Description:** Get all permissions

### Success Response — 200 OK

```json
{
  "message": "permissions fetched successfully",
  "success": true,
  "data": [],
  "err": null
}
```

---

## PUT `/permissions/{id}`

**Authorization:** Admin
**Description:** Update permission

### Request Body

```json
{
  "name": "booking:update",
  "desc": "Update booking",
  "resource": "booking",
  "action": "update"
}
```

### Success Response — 200 OK

```json
{
  "message": "permission updated successfully",
  "success": true,
  "data": {},
  "err": null
}
```

---

## DELETE `/permissions/{id}`

**Authorization:** Admin
**Description:** Delete permission

### Success Response — 200 OK

```json
{
  "message": "permission deleted successfully",
  "success": true,
  "data": null,
  "err": null
}
```

---

# 🔗 Role ↔ Permission Mapping

---

## POST `/roles-permissions/{roleId}/{permissionId}`

**Authorization:** Admin
**Description:** Assign permission to role

### Success Response — 200 OK

```json
{
  "message": "Successfully assigned permission to role",
  "success": true,
  "data": null,
  "err": null
}
```

---

## DELETE `/roles-permissions/{roleId}/{permissionId}`

**Authorization:** Admin
**Description:** Remove permission from role

### Success Response — 200 OK

```json
{
  "message": "Successfully removed permission to role",
  "success": true,
  "data": null,
  "err": null
}
```

---

## GET `/roles-permissions/{roleId}`

**Authorization:** Admin
**Description:** Get permissions of a role

### Success Response — 200 OK

```json
{
  "message": "fetched all permissions",
  "success": true,
  "data": [],
  "err": null
}
```

---

# 👥 User ↔ Role Assignment

---

## POST `/users-roles/assign/{userId}/{roleId}`

**Authorization:** Admin
**Description:** Assign role to user

### Success Response — 200 OK

```json
{
  "message": "Successfully assigned role to the user",
  "success": true,
  "data": null,
  "err": null
}
```

---

# 🛠 Setup Instructions

## 1️⃣ Clone the project

```bash
git clone https://github.com/Vivan1133/Airbnb-AuthInGo.git <ProjectName>
```

## 2️⃣ Move into project

```bash
cd <ProjectName>
```

## 3️⃣ Run the server

```bash
go run main.go
```

---

# 📌 Common Response Format

## Success

```json
{
  "message": "success message",
  "success": true,
  "data": {},
  "err": null
}
```

## Error

```json
{
  "message": "error message",
  "error": "error details",
  "data": null,
  "success": false
}
```
