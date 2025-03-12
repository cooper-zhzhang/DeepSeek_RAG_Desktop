package main

import (
	"dp_client/storage/model"

	"gorm.io/gen"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../storage/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	//DeepSeekDB, _ := gorm.Open(mysql.Open("we:@(127.0.0.1:3306)/deepseek?charset=utf8mb4&parseTime=True&loc=Local"))
	//g.UseDB(DeepSeekDB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.Agent{}, model.Conversation{}, model.Message{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, model.Agent{}, model.Conversation{}, model.Message{})

	// Generate the code
	g.Execute()
}
