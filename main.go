package main

import (
	"net/http"

	//"encoding/json"

	"strconv"

	gf "github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

const RECORDS = 100

type Record struct {
	Id      int
	Name    string
	Email   string
	Phone   string
	Company string
}

func main() {

	records := []Record{}
	for i := 0; i < 100; i++ {
		records = append(records, Record{Id: i, Name: gf.Name(), Email: gf.Email(), Phone: gf.Phone(), Company: gf.Company()})
	}

	r := gin.Default()
	r.GET("/export", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"data": records,
		// })
		//j, _ := json.Marshal(records)
		//c.String(http.StatusOK, string(j))

		// Create a new sheet.
		f := excelize.NewFile()

		f.SetColWidth("Sheet1", "A", "A", 10)
		f.SetColWidth("Sheet1", "B", "B", 30)
		f.SetColWidth("Sheet1", "C", "C", 40)
		f.SetColWidth("Sheet1", "D", "D", 20)
		f.SetColWidth("Sheet1", "E", "E", 50)

		for n, row := range records {
			ns := strconv.Itoa(n)
			f.SetCellValue("Sheet1", "A"+ns, row.Id)
			f.SetCellValue("Sheet1", "B"+ns, row.Name)
			f.SetCellValue("Sheet1", "C"+ns, row.Email)
			f.SetCellValue("Sheet1", "D"+ns, row.Phone)
			f.SetCellValue("Sheet1", "E"+ns, row.Company)
		}

		buff, err := f.WriteToBuffer()
		if err != nil {
			return
		}
		// Save spreadsheet by the given path.
		//return buf.WriteTo(w)
		// if err := f.SaveAs("Book1.xlsx"); err != nil {
		// 	fmt.Println(err)
		// }
		c.String(http.StatusOK, buff.String())
	})
	r.Run()
}
