package integration_test

import (
	"bytes"
	"encoding/json"

	"github.com/Aspiand/lego/database"
	"github.com/Aspiand/lego/routers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTest() (*gorm.DB, *gin.Engine) {
	db := database.Initialize(":memory:")
	router := gin.Default()
	routers.SetupRouters(db, router)

	return db, router
}

func toJSONReader(data any) *bytes.Reader {
	b, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return bytes.NewReader(b)
}
