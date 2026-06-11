package config

var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	DBType     string
)

func InitConfig() {
	DBType = "sqlite"
	DBHost = "localhost"
	DBPort = "3306"
	DBUser = "root"
	DBPassword = "password"
	DBName = "testdev_platform"
	JWTSecret = "testdev_learning_platform_jwt_secret_key"
}