package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB // Variável global para armazenar a conexão com o banco de dados

// Função para conectar ao banco de dados PostgreSQL
func ConnectarBancoDados() {
	// Tenta carregar .env em locais possíveis (padrão, ../.env, ../../.env)
	if err := godotenv.Load(); err != nil {
		_ = godotenv.Load("../.env")
		_ = godotenv.Load("../../.env")
	}

	// Obtendo as variáveis de ambiente e validando
	host := os.Getenv("HOST")
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port == 0 {
		fmt.Printf("PORT inválida ou não definida: %q\n", portStr)
		panic("porta do banco inválida ou não definida")
	}

	user := os.Getenv("users")
	if user == "" {
		user = os.Getenv("USER")
	}
	if user == "" {
		fmt.Println("Usuário do banco não definido nas variáveis de ambiente (users ou USER)")
		panic("usuário do banco não definido")
	}

	dbname := os.Getenv("NAME")
	pass := os.Getenv("PASSWORD")
	if host == "" || dbname == "" || pass == "" {
		fmt.Println("HOST, NAME ou PASSWORD não definidos nas variáveis de ambiente")
		panic("variáveis de conexão ausentes")
	}

	// Criando a string de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, errSql := sql.Open("postgres", psqlInfo)
	if errSql != nil {
		fmt.Println("Erro ao abrir conexão com o banco de dados:", errSql)
		panic(errSql)
	}

	if errPing := db.Ping(); errPing != nil {
		fmt.Println("Erro ao conectar ao banco de dados (ping):", errPing)
		panic(errPing)
	}

	DB = db // Atribuindo a conexão à variável global
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}
