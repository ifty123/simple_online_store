package migration

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ifty123/simple_online_store/database"
	"github.com/ifty123/simple_online_store/internal/model"
)

var tables = []interface{}{
	&model.User{},
	&model.Category{},
	&model.Product{},
	&model.Cart{},
	&model.Transaction{},
	&model.TransactionDetail{},
}

func Migrate() {

	conn := database.GetConnection()

	conn.AutoMigrate(tables...)
}

func Rollback() {

	conn := database.GetConnection()

	for i := len(tables) - 1; i >= 0; i-- {
		conn.Migrator().DropTable(tables[i])
	}
}

func Status() {
	var (
		conn        = database.GetConnection()
		colorReset  = "\033[0m"
		colorGreen  = "\033[32m"
		colorYellow = "\033[33m"
	)

	fmt.Printf("In database %s:\n", conn.Migrator().CurrentDatabase())
	for _, table := range tables {
		var name string

		t := reflect.TypeOf(table)
		if t.Kind() == reflect.Ptr {
			name = strings.ToLower(t.Elem().Name()) + "s"
		} else {
			name = strings.ToLower(t.Name()) + "s"
		}

		if conn.Migrator().HasTable(table) {
			fmt.Println("\t", name, "===>", string(colorGreen), "migrated", string(colorReset))
		} else {
			fmt.Println("\t", name, "===>", string(colorYellow), "not migrated", string(colorReset))
		}
	}
}
