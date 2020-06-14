package core

import (
	"database/sql/driver"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

type BaseModel struct {
	ID int64 `gorm:"primary_key" json:"id"`

	CreatedAt MyTime `json:"created_at"`
	UpdatedAt MyTime `json:"updated_at"`

	Isdeleted bool `gorm:"column:is_deleted" json:"is_deleted,omitempty"`
}

func Setup() {
	var (
		err error
	)

	c, err := Conf.Map("mysql")

	Db, err = gorm.Open("mysql", c["Url"])

	if err != nil {
		log.Println(err)
	}

	Db.LogMode(true)
	//Db.SingularTable(true)

	// https://github.com/go-sql-driver/mysql/issues/461
	Db.DB().SetMaxIdleConns(5)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(50 * time.Second)
}

// 用于时间的格式化输出
type MyTime struct {
	time.Time
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	// https://segmentfault.com/q/1010000010976398
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
