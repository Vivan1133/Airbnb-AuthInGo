# AuthService
AuthService is a centralized authentication and authorization microservice designed to work as an API Gateway component.

## 🚀 Features

🔐 JWT-based Authentication

📝 User Sign Up & Sign In

🔑 Role-Based Access Control (RBAC)

🧩 Permission Management

🔒 Password Hashing using bcrypt

✅ Request Validation using middlewares

🚦 Rate Limiting

🔁 Reverse Proxy / API Gateway

🛡 Secure Middleware-based Authorization

💎 DB Migrations using goose lib

## 🗄 Database Design
Database Name: ```airbnb_auth_dev```

### ```users```

| Column    | Type                 |
| --------- | -------------------- |
| id        | INT                  |
| name      | String               |
| email     | String               |
| password  | String (bcrypt hash) |
| createdat | Timestamp            |
| updatedat | Timestamp            |

### ```roles```

| Column    | Type      |
| --------- | --------- |
| id        | INT       |
| name      | String    |
| desc      | String    |
| createdat | Timestamp |
| updatedat | Timestamp |

### ```permissions```

| Column    | Type      |
| --------- | --------- |
| id        | INT       |
| name      | String    |
| desc      | String    |
| resource  | String    |
| action    | String    |
| createdat | Timestamp |
| updatedat | Timestamp |

### ```users_roles```

| Column    | Type      |
| --------- | --------- |
| id        | UUID      |
| user_id   | UUID      |
| role_id   | UUID      |
| createdat | Timestamp |
| updatedat | Timestamp |

### ```roles_permissions```

| Column        | Type      |
| ------------- | --------- |
| id            | UUID      |
| role_id       | UUID      |
| permission_id | UUID      |
| createdat     | Timestamp |
| updatedat     | Timestamp |

## Base URL
```
{server-url}
ex : http://localhost:3004
```

## ```👤 User Authentication & Management Routes```

| Method   | Endpoint                   | Authorization  | Description             |
| -------- | -------------------------- | -------------- | ----------------------- |
| `POST`   | `/auth/signup`             | ❌ Public       | Register a new user     |
| `POST`   | `/auth/signin`             | ❌ Public       | User login (JWT issued) |
| `GET`    | `/auth/user/{id}`          | ✅ User / Admin | Get user by ID          |
| `GET`    | `/auth/user/email/{email}` | ✅ User / Admin | Get user by email       |
| `GET`    | `/auth/users`              | ✅ User / Admin | Get all users           |
| `DELETE` | `/auth/user/{id}`          | ✅ Admin        | Delete user by ID       |

## ```🧑‍💼 Role Management Routes```

| Method   | Endpoint                 | Authorization | Description       |
| -------- | ------------------------ | ------------- | ----------------- |
| `GET`    | `/roles`                 | ✅ Admin       | Get all roles     |
| `GET`    | `/roles/id/{roleId}`     | ✅ Admin       | Get role by ID    |
| `GET`    | `/roles/name/{roleName}` | ✅ Admin       | Get role by name  |
| `POST`   | `/roles`                 | ✅ Admin       | Create a new role |
| `PATCH`  | `/roles`                 | ✅ Admin       | Update role       |
| `DELETE` | `/roles/id/{roleId}`     | ✅ Admin       | Delete role       |

## ```🔗 Role ↔ Permission Mapping```

| Method   | Endpoint                                     | Authorization | Description                 |
| -------- | -------------------------------------------- | ------------- | --------------------------- |
| `POST`   | `/roles-permissions/{roleId}/{permissionId}` | ✅ Admin       | Assign permission to role   |
| `DELETE` | `/roles-permissions/{roleId}/{permissionId}` | ✅ Admin       | Remove permission from role |
| `GET`    | `/roles-permissions/{roleId}`                | ✅ Admin       | Get permissions of a role   |

## ```🔐 Permission Management Routes```

| Method   | Endpoint            | Authorization | Description          |
| -------- | ------------------- | ------------- | --------------------  |
| `POST`   | `/permissions`      | ✅ Admin       | Create permission    |
| `GET`    | `/permissions`      | ✅ Admin       | Get all permissions  |
| `GET`    | `/permissions/{id}` | ✅ Admin       | Get permission by ID |
| `PUT`    | `/permissions/{id}` | ✅ Admin       | Update permission    |
| `DELETE` | `/permissions/{id}` | ✅ Admin       | Delete permission    |

## ```👥 User ↔ Role Assignment```

| Method | Endpoint                                | Authorization | Description         |
| ------ | --------------------------------------- | ------------- | ------------------- |
| `POST` | `/users-roles/assign/{userId}/{roleId}` | ✅ Admin       | Assign role to user |



