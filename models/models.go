package models

import (
	"fmt"
	"kusnandartoni/starter/pkg/logging"
	"kusnandartoni/starter/pkg/setting"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // add database driver bridge
)

var db *gorm.DB

// Base :
type Base struct {
	ID         int64  `json:"id" gorm:"primary_key;"`
	CreatedOn  int64  `json:"-"`
	CreatedBy  string `json:"-"`
	ModifiedOn int64  `json:"-"`
	ModifiedBy string `json:"-"`
	DeletedOn  int64  `json:"-"`
	DeletedBy  string `json:"-"`
}

// Setup :
func Setup() {
	now := time.Now()
	var err error
	connectionParams := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port)

	db, err = gorm.Open(setting.DatabaseSetting.Type, connectionParams)

	if err != nil {
		logging.Fatal("0", "models.Setup err: ", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	go autoMigrate()

	timeSpent := time.Since(now)
	log.Printf("Config database is ready in %v", timeSpent)
}

func autoMigrate() {
	// Add auto migrate bellow this line
	log.Println("STARTING AUTO MIGRATE ")
	db.AutoMigrate(Members{}, Classes{}, AuditLog{})

	log.Println("FINISHING AUTO MIGRATE ")
}

// CloseDB :
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

// AuditLog :
type AuditLog struct {
	ID        int64  `json:"id"`
	CreatedOn int64  `json:"created_on"`
	Level     string `json:"level"`
	UUID      string `json:"uuid"`
	FuncName  string `json:"func_name"`
	FileName  string `json:"file_name"`
	Line      int64  `json:"line"`
	Time      string `json:"time"`
	Message   string `json:"message"`
}
