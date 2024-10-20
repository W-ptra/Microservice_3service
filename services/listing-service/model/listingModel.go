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

	err = db.AutoMigrate(&Listing{})
	if err != nil{
		log.Fatalln("Error while migrating database:\n",err)
	}
}

func GetAllListing(pageNum, pageSize, userId int) ([]Listing, error) {
    db, err := getConnection()

    if err != nil {
        return nil, err
    }

    var listings []Listing
    offset := (pageNum - 1) * pageSize

    query := db.Order("created_at DESC").
        Limit(pageSize).
        Offset(offset)

    if userId != 0 {
        query = query.Where("user_id = ?", userId)
    }
    result := query.Find(&listings)

    return listings, result.Error
}

func CreateListing(listing Listing)(Listing,error){
	db,err := getConnection()

	if err != nil{
		return Listing{},err
	}

	result := db.Create(&listing)
	return listing,result.Error
}