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

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func rowsToStrings(rows *sql.Rows) []map[string]string {
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	pretty := [][]string{cols}
	//res := map[string]string{}

	result := []map[string]string{}
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for rows.Next() {
		for i := range columns {
			columns[i] = new(interface{})
			columnPointers[i] = &columns[i]
		}
		//fmt.Println(columns[:]...)
		if err := rows.Scan(columns[:]...); err != nil {
			panic(err)
		}
		cur := make([]string, len(cols))
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
			cur[i] = str
			fmt.Println(cols[i])
			data[cols[i]] = str
		}
		result = append(result, data)
		pretty = append(pretty, cur)
	}
	return result
}
func GetAll(c *gin.Context) {
	rows, _ := db.DB.Query("select * from hiring")
	newresult := rowsToStrings(rows)
	cols, _ := rows.Columns()
	data := make(map[string]string)
	result := []map[string]string{}
	defer rows.Close()
	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)

		for i, colName := range cols {

			var str string
			if columnPointers[i] == nil {
				str = "NULL"
			} else {
				switch v := columnPointers[i].(type) {
				case []byte:
					str = string(v)
				default:
					str = fmt.Sprintf("%v", v)
				}
			}
			//data[colName] = columns[i]
			data[colName] = str
		}
		result = append(result, data)

	}

	//ff := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	c.JSON(http.StatusOK, gin.H{"success": true, "result": newresult})
}
