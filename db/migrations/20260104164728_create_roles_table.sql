-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ROLES (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    NAME VARCHAR(100) NOT NULL UNIQUE,
    DESCRIPTION TEXT NOT NULL,
    CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    UPDATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO ROLES (
    NAME, DESCRIPTION
) VALUES 
('admin' , 'admin having full access'),
('user', 'user having limited access'),
('moderator', 'moderator having special privileges');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ROLES;
-- +goose StatementEnd
