package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"weather-monster/models"
)

func InitSqlDB(config *models.Config) (db *gorm.DB) {
	db, err := gorm.Open(config.Sql.Type, fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",config.Sql.Host,config.Sql.Port,config.Sql.User,config.Sql.DbName,config.Sql.Sslmode,config.Sql.Password))
	if err != nil {
		panic("failed to connect postgres database")
	}
	migration(db)
	db.LogMode(true)
	return db
}
func migration(db *gorm.DB){
	db.AutoMigrate(&models.City{})
	db.AutoMigrate(&models.Temperature{})
	db.AutoMigrate(&models.Webhook{})

	db.Model(&models.Temperature{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Webhook{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
}