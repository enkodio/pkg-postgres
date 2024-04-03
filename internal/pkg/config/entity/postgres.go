package entity

type Config interface {
	GetDSN(string) string
	GetMaxOpenConns() int
	GetMaxAttempts() int
	GetMaxDelay() int
}

type PostgresConfig struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	DBName       string `json:"dbName"`
	Password     string `json:"password"`
	MaxOpenConns int    `json:"maxOpenConns"`
	MaxAttempts  int    `json:"maxAttempts"`
	MaxDelay     int    `json:"maxDelay"`
}

func (p PostgresConfig) GetDSN(serviceName string) string {
	return "host=" + p.Host +
		" port=" + p.Port +
		" user=" + p.User +
		" dbname=" + p.DBName +
		" sslmode=disable password=" + p.Password +
		" application_name=" + serviceName
}

func (p PostgresConfig) GetMaxOpenConns() int {
	return p.MaxOpenConns
}

func (p PostgresConfig) GetMaxAttempts() int {
	return p.MaxAttempts
}

func (p PostgresConfig) GetMaxDelay() int {
	return p.MaxDelay
}
