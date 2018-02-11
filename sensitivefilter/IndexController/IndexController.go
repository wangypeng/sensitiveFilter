package indexController

import (
    "fmt"
    "net/http"
    "gopkg.in/mgo.v2"

    "encoding/json"

    "github.com/julienschmidt/httprouter"

    "goSensitive/sensitivefilter/dao"
    "goSensitive/sensitivefilter/model"
    "goSensitive/sensitivefilter/wordFilter"
)

type Controller struct{
    session *mgo.Session
}

type ResultData struct {
    Success bool        `json:"success"`
    Code string         `json:"code"`      
    Message string      `json:"message"`
    Word string         `json:"word"`
}
/**
 * add sensitive controller function
 */
func (clt *Controller) addSensitive (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    fmt.Println("hello world ....")
    var result ResultData

    dao := dao.Dao{
        S:clt.session,
    }
    word := &model.WordStruct{
        Word:params.ByName("word"),
    }

    queryResult  := dao.FindOne(params.ByName("word"))
    if  queryResult {
        result.Code = "501"
        result.Success = false
        result.Message = "this sensitive word is having"
    }else{
        // orm
        err := dao.Insert(word)

        var wordlist []string
        wordlist = append(wordlist,params.ByName("word"))
        wordFilter.LoadSensitiveWord(wordlist)
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
    }
    jsondata , _ := json.Marshal(result)

    fmt.Fprint(w, string(jsondata))
}

/**
 * add sensitive controller function
 */
func (clt *Controller) delSensitive (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    fmt.Println("hello world ....")
    var result ResultData

    dao := dao.Dao{
        S:clt.session,
    }

    err := dao.Delete(params.ByName("word"))

    wordFilter.DelSensitiveWord(params.ByName("word"))

    if err == nil {
        fmt.Println("del sensitive true")
        result.Code = "200"
        result.Success = true
        result.Message = "success"
    }else{
        fmt.Println("del sensitive false")
        result.Code = "500"
        result.Success = false
        result.Message = "false"
    }
    jsondata , _ := json.Marshal(result)

    fmt.Fprint(w, string(jsondata))
}

/**
 *  replace sensitive word to '*' 
 */
func (clt *Controller) check (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    var result ResultData

    text := ps.ByName("word")

    resultText := wordFilter.ReplaceSensitiveWord(text)

    fmt.Println("add sensitive success")
    result.Code = "200"
    result.Success = true
    result.Message = "success"
    result.Word = resultText

    data ,jsonerr := json.Marshal(result)
    if jsonerr != nil {
        fmt.Println("check sensitive false ...")
    }
    fmt.Fprint(w, string(data))
}


