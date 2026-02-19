-- Tabela users compatível com in/models/user.go e in/repositories/user_repository.go
-- Executar no PostgreSQL: psql -U postgres -d SEU_DATABASE -f migrations/001_create_users.sql

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Índice para busca por email (login)
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
