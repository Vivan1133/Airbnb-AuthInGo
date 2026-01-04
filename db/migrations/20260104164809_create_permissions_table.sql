-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS PERMISSIONS (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    NAME VARCHAR(100) NOT NULL UNIQUE,
    DESCRIPTION TEXT NOT NULL,
    RESOURCE TEXT NOT NULL,
    ACTION TEXT NOT NULL,
    CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO PERMISSIONS (NAME, DESCRIPTION, RESOURCE, ACTION) VALUES
('user:create', 'Create a new user', 'user', 'create'),
('user:read', 'View user details', 'user', 'read'),
('user:update', 'Update user profile or data', 'user', 'update'),
('user:delete', 'Delete a user', 'user', 'delete'),
('user:list', 'View list of users', 'user', 'list'),
('permission:create', 'Create a new permission', 'permission', 'create'),
('permission:read', 'View permission details', 'permission', 'read'),
('permission:assign', 'Assign permissions to roles', 'permission', 'assign'),
('permission:revoke', 'Revoke permissions from roles', 'permission', 'revoke'),
('permission:delete', 'Delete a permission', 'permission', 'delete'),
('permission:list', 'View all permissions', 'permission', 'list'),
('role:create', 'Create a new role', 'role', 'create'),
('role:read', 'View role details', 'role', 'read'),
('role:update', 'Update role details', 'role', 'update'),
('role:delete', 'Delete a role', 'role', 'delete'),
('role:list', 'View all roles', 'role', 'list'),
('role:assign', 'Assign roles to users', 'role', 'assign'),
('role:revoke', 'Revoke roles from users', 'role', 'revoke');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS PERMISSIONS;
-- +goose StatementEnd
