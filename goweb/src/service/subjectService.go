package service

import (
	"encoding/json"
	"io"
	"net/http"
	"wserver/src/entity"
	"wserver/src/util"
)

func GetSubject(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	grade := params["grade"][0]
	subject := params["subject"][0]
	data := util.Get(grade + "_" + subject)
	var dataJson entity.SubjectAllData
	if err := json.Unmarshal([]byte(data), &dataJson); err == nil {
	}
	if subject != "6010" {
		for i := 0; i < len(dataJson.Result.SubjectList); i++ {
			dataJson.Result.SubjectList[i].SumPrice = dataJson.Result.SubjectList[i].MinPrice * dataJson.Result.SubjectList[i].SoldCount
		}
		result, _ := json.Marshal(dataJson.Result.SubjectList)
		_, _ = io.WriteString(w, string(result))
	} else {
		result, _ := json.Marshal(dataJson.Result)
		_, _ = io.WriteString(w, string(result))
	}

}

func GetGrades(w http.ResponseWriter, r *http.Request) {
	data := util.Get("grades")
	_, _ = io.WriteString(w, data)
}
