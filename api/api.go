/*
 * Created on: 29/10/2018
 *     Author: guilhermehenrique
 */

package api

import (
	"encoding/json"
	"github.com/guiacarneiro/visualgolibs/logger"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

//PostHandler - tratamento do body de um post
func PostHandler(w http.ResponseWriter, r *http.Request, obj interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&obj)
	defer r.Body.Close()

	if err != nil {
		PostResponse(w, r, &obj, http.StatusBadRequest, err.Error())
		return false
	}

	if !validaParametrosObrigatorios(reflect.ValueOf(obj).Elem().Interface()) {
		PostResponse(w, r, &obj, http.StatusBadRequest, "Parâmetros incorretos")
		return false
	}

	return true
}

//PostHandlerBytes - tratamento do body de um post e retorna seus bytes
func PostHandlerBytes(w http.ResponseWriter, r *http.Request, obj interface{}) []byte {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		PostResponse(w, r, &obj, http.StatusBadRequest, err.Error())
		return nil
	}

	err = json.Unmarshal(body, &obj)
	if err != nil {
		PostResponse(w, r, &obj, http.StatusBadRequest, err.Error())
		return nil
	}

	if !validaParametrosObrigatorios(reflect.ValueOf(obj).Elem().Interface()) {
		PostResponse(w, r, &obj, http.StatusBadRequest, "Parâmetros incorretos")
		return nil
	}

	return body
}

func validaParametrosObrigatorios(obj interface{}) bool {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		typeField := t.Field(i)
		valueField := v.Field(i)
		if typeField.Tag.Get("obrigatorio") == "true" {
			if valueField.IsNil() ||
				((typeField.Tag.Get("nullable") == "false") && ((valueField.Elem().Interface() == "") || (valueField.Elem().Interface() == 0))) {
				return false
			}
		}
	}
	return true
}

//PostResponse - tratamento da resposta de um post
func PostResponse(w http.ResponseWriter, r *http.Request, parameters interface{}, httpStatus int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	var log = "endpoint: " + strconv.Itoa(httpStatus) + " " + r.Method + " " + r.RequestURI
	if parameters != nil {
		if reflect.ValueOf(parameters).Kind() == reflect.String {
			log += " - params: " + parameters.(string)
		} else {
			bytes, err := json.Marshal(parameters)
			if err == nil {
				log += " - params: " + string(bytes)
			}
		}
	}

	if httpStatus != 200 {
		log += " - erro: " + response.(string)
		logger.LogError.Println(log)

		type Erro struct {
			Mensagem string `json:"mensagem"`
		}
		var erro Erro
		erro.Mensagem = response.(string)
		bytes, err := json.Marshal(erro)
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			_, err = io.WriteString(w, string(bytes))
			if err != nil {
				logger.LogError.Println("Erro ao responder a requisicao.", err)
			}
		}
	} else {
		logger.LogInfo.Println(log)

		bytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			_, err = io.WriteString(w, string(bytes))
			if err != nil {
				logger.LogError.Println("Erro ao responder a requisicao.", err)
			}
		}
	}
}
