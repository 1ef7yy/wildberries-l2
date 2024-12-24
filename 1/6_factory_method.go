package main

type Database interface {
	Connect() string
	Disconnect() string
}

type MySQLDatabase struct{}

func (m *MySQLDatabase) Connect() string {
	return "Подключение к MySQL..."
}

func (m *MySQLDatabase) Disconnect() string {
	return "Отключение от MySQL..."
}

type PostgreSQLDatabase struct{}

func (p *PostgreSQLDatabase) Connect() string {
	return "Подключение к PostgreSQL..."
}

func (p *PostgreSQLDatabase) Disconnect() string {
	return "Отключение от PostgreSQL..."
}

type MongoDBDatabase struct{}

func (m *MongoDBDatabase) Connect() string {
	return "Подключение к MongoDB..."
}

func (m *MongoDBDatabase) Disconnect() string {
	return "Отключение от MongoDB..."
}

// Фабрика баз данных
type DatabaseFactory struct{}

func (f *DatabaseFactory) CreateDatabase(databaseType string) Database {
	switch databaseType {
	case "mysql":
		return &MySQLDatabase{}
	case "postgres":
		return &PostgreSQLDatabase{}
	case "mongodb":
		return &MongoDBDatabase{}
	default:
		return nil
	}
}

// func main() {
// 	factory := &DatabaseFactory{}

// 	mysqlDB := factory.CreateDatabase("mysql")
// 	fmt.Println(mysqlDB.Connect())
// 	fmt.Println(mysqlDB.Disconnect())

// 	postgresDB := factory.CreateDatabase("postgres")
// 	fmt.Println(postgresDB.Connect())
// 	fmt.Println(postgresDB.Disconnect())

// 	mongoDB := factory.CreateDatabase("mongodb")
// 	fmt.Println(mongoDB.Connect())
// 	fmt.Println(mongoDB.Disconnect())
// }
