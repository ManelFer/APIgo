CREATE TABLE estoque (
    id SERIAL PRIMARY KEY,
    equipamento VARCHAR(255) NOT NULL,
    marca VARCHAR(255) NOT NULL,
    modelo VARCHAR(255) NOT NULL,
    patrimonio VARCHAR(255) NOT NULL,
    quantidade INTEGER NOT NULL
);