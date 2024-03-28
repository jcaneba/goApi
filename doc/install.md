# install
mkdir /var/www/html/goApi
cd /var/wwww/html/goApi
go mod init goApi
go get github.com/gin-gonic/gin
go get gorm.io/driver/mysql gorm.io/gorm
go get github.com/joho/godotenv
mkdir doc controllers models
touch .env main.go
echo -e "PORT=8081\nDB_USER='root'\nDB_PASS='1234'\nDB_HOST='172.19.0.2'\nDB_TABLE='test'" > .env
# deploy
go mod tidy
go run main.go