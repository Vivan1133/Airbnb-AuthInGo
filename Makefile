MIGRATION_FOLDER=db/migrations
DB_URL=root:adminadmin@tcp(127.0.0.1:3306)/airbnb_auth_dev

migrate-create:	# make migrate-create name="whatever-migration-name"
	goose -dir $(MIGRATION_FOLDER) create $(name) sql

migrate-up:	
	goose -dir $(MIGRATION_FOLDER) mysql "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATION_FOLDER) mysql "$(DB_URL)" down

# Rollback all migrations and reset database # make migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" reset

# Show current migration status # make migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" status

# Redo last migration (Down then Up) # make migrate-redo
migrate-redo:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" redo

# Run specific migration version # make migrate-version version=20200101120000
migrate-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version # make migrate-down-to version=20200101120000
migrate-down-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version # make migrate-force version=20200101120000
migrate-force:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" force $(version)

# Print Goose help # make migrate-help
migrate-help:
	goose -h