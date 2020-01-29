package global

import "github.com/jinzhu/gorm"

const DBName = "swapid.db"
const DBDialect = "sqlite3"
const ListeningAddress = ":8089"

var DB *gorm.DB
