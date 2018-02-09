package IndexController

import (
    "fmt"
    "net/http"
    "gopkg.in/mgo.v2"

    "encoding/json"

    "github.com/julienschmidt/httprouter"

    "goSensitive/sensitivefilter/Dao"
    "goSensitive/sensitivefilter/model"
    "goSensitive/sensitivefilter/WordFilter"
)

type Controller struct{
    session *mgo.Session
}

type ResultData struct {
    Success bool        `json:"success"`
    Code string         `json:"code"`      
    Message string      `json:"message"`
}

func (clt *Controller) addSensitive (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    fmt.Println("hello world ....")
    var result ResultData

    dao := Dao.Dao{
        S:clt.session,
    }
    fmt.Println(params.ByName("word"))
    word := &model.WordStruct{
        Word:params.ByName("word"),
    }
    err := dao.Insert(word)
    if err == nil {
        fmt.Println("add sensitive success")
        result.Code = "200"
        result.Success = true
        result.Message = "success"
    }else{
        fmt.Println("add sensitive false")
        result.Code = "500"
        result.Success = false
        result.Message = "false"
    }

    fmt.Println(result)
    jsondata ,jsonerr := json.Marshal(result)
    if jsonerr != nil {
        fmt.Println("add sensitive ")
    }
    fmt.Fprint(w, string(jsondata))
}

func (clt *Controller) check (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    var result ResultData

    text := ps.ByName("word")

    resultText := WordFilter.ReplaceSensitiveWord(text)

    fmt.Println("add sensitive success")
    result.Code = "200"
    result.Success = true
    result.Message = "success"

    data ,jsonerr := json.Marshal(resultText)
    if jsonerr != nil {
        fmt.Println("check sensitive false ...")
    }
    fmt.Fprint(w, data)
}


