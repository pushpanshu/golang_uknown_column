package controller

import (
	"GIN/jones-gin/db"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func rowsToStrings(rows *sql.Rows) []map[string]string {
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	result := []map[string]string{}
	columns := make([]interface{}, len(cols))
	for rows.Next() {
		for i := range columns {
			columns[i] = new(interface{})
		}
		if err := rows.Scan(columns...); err != nil {
			panic(err)
		}

		data := map[string]string{}
		for i := range columns {
			val := *columns[i].(*interface{})
			var str string
			if val == nil {
				str = "NULL"
			} else {
				switch v := val.(type) {
				case []byte:
					str = string(v)
				default:
					str = fmt.Sprintf("%v", v)
				}
			}

			data[cols[i]] = str
		}
		result = append(result, data)
	}
	return result
}
func GetAll(c *gin.Context) {
	rows, _ := db.DB.Query("select * from hiring")
	newresult := rowsToStrings(rows)
	c.JSON(http.StatusOK, gin.H{"success": true, "result": newresult})
}
