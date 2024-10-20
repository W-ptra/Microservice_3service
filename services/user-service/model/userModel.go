package model

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"fmt"
	"log"
)

var dbConnection *gorm.DB

func getConnection()(*gorm.DB,error){
	if dbConnection == nil{
		err := godotenv.Load()
		if err != nil {
			log.Fatalln("Error loading .env file:", err)
		}
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Shanghai", 
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_SSLMODE"),
		)
		connection,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
		if err!=nil{
			return nil,err
		}
		dbConnection = connection
	}
	return dbConnection,nil
}

func Migration(){
	db,err := getConnection()
	if err != nil{
		log.Println("Error to connect to database:\n",err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil{
		log.Fatalln("Error while migrating database:\n",err)
	}
}

func GetAllUsers(pageNum int, pageSize int) ([]User, error) {
	db, err := getConnection()

	if err != nil {
		return nil, err
	}

	var users []User
	offset := (pageNum - 1) * pageSize

	result := db.Order("created_at DESC").
		Limit(pageSize).    
		Offset(offset).     
		Find(&users)

	return users, result.Error
}


func GetUserById(id int)(User,error){
	db,err := getConnection()

	if err != nil{
		return User{},err
	}

	var user User
	result := db.Find(&user,"id = ?",id)

	return user,result.Error
}

func CreateUser(user User)(User,error){
	db,err := getConnection()

	if err != nil{
		return User{},err
	}

	result := db.Create(&user)

	return user,result.Error
}