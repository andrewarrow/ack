package dog

import (
	"encoding/json"
	"fmt"
	"infinity-dog/database"
	"infinity-dog/files"
	"io/ioutil"
	"os"
)

func Import() {
	sampleFiles, err := ioutil.ReadDir("samples")
	if err != nil {
		fmt.Println(err)
		return
	}

	//os.Remove("sqlite.db")
	database.CreateSchema()
	db := database.OpenTheDB()
	defer db.Close()

	metaMapBytes := map[string]int64{}
	metaMapExceptions := map[string]int64{}

	for i, file := range sampleFiles {
		jsonString := files.ReadFile("samples/" + file.Name())
		var logResponse LogResponse
		json.Unmarshal([]byte(jsonString), &logResponse)

		tx, _ := db.Begin()
		s := `insert into services (id,name,msg,message,exception,logged_at) values (?,?,?,?,?,?)`
		prep, _ := tx.Prepare(s)

		for _, d := range logResponse.Data {
			ts := d.Attributes.Timestamp
			prep.Exec(d.Id, d.Attributes.Service, d.Attributes.SubAttributes.Msg,
				d.Attributes.Message, d.Attributes.SubAttributes.Exception, ts)

			metaMapBytes[d.Attributes.Service] += int64(len(d.Attributes.SubAttributes.Msg)) +
				int64(len(d.Attributes.Message)) +
				int64(len(d.Attributes.SubAttributes.Exception))
			if len(d.Attributes.SubAttributes.Exception) > 0 {
				metaMapExceptions[d.Attributes.Service]++
			}
		}

		tx.Commit()
		fmt.Println("done", i)
		os.Remove("samples/" + file.Name())
	}

	for k, v := range metaMapBytes {
		tx, _ := db.Begin()
		s := `insert into service_meta (name,total_exceptions,total_bytes) values (?,?,?)`
		prep, _ := tx.Prepare(s)
		prep.Exec(k, metaMapExceptions[k], v)
		tx.Commit()
	}
	fmt.Println("done meta")
}
