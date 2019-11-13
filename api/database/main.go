// GORM has wrapped some drivers,

import "github.com/jinzhu/gorm/dialects/mysql"

// AutoMigration 
// AutoMigrate will only create tables, missing columns and missing indexes

db.AutoMigrate(&User{})
db.AutoMigrate(&User{}, &Product{}, &Order{})

db.HasTable(&User{})

db.HasTable("users")
// Check wheater table with model User is there or not

db.CreateTable(&User{})
db.DropTable(&User{})
db.DropTable("users")

// Models - gorm.Model struct
// gorm.Model - base model definition includes ID, CreatedAt, UpdatedAt, DeletedAt you could embed it in your model or only write those fields you want/

type User struct {
	gorm.Model
	Name string
	
}