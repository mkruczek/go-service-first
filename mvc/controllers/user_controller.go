package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkruczek/go-service-first/mvc/services"
	"github.com/mkruczek/go-service-first/mvc/utils"
)

func GetUsers(ctx *gin.Context) {
	users := services.UserService.GetUsers()
	utils.RespondOK(ctx, users)
}

func GetUser(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Msg:    fmt.Sprintf("Wrong user ID param : %s.", idParam),
			Status: http.StatusBadRequest,
			Code:   "bad_reques",
		}
		utils.RespondError(ctx, appErr)
		return
	}

	log.Printf("Procesing for  user ID : %d.\n", id)

	u, appErr := services.UserService.GetUser(id)
	if appErr != nil {
		utils.RespondError(ctx, appErr)
		return
	}

	utils.RespondOK(ctx, u)
}

// without gin

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	users := services.UserService.GetUsers()
// 	jsonValue, _ := json.Marshal(users)
// 	w.Write(jsonValue)
// }

// func GetUser(resp http.ResponseWriter, req *http.Request) {

// 	idParam := req.URL.Query().Get("id")
// 	id, err := strconv.ParseUint(idParam, 10, 64)
// 	if err != nil {
// 		appErr := &utils.ApplicationError{
// 			Msg:    fmt.Sprintf("Wrong user ID param : %s.", idParam),
// 			Status: http.StatusBadRequest,
// 			Code:   "bad_reques",
// 		}
// 		resp.WriteHeader(appErr.Status)
// 		jsonValue, _ := json.Marshal(appErr)
// 		resp.Write(jsonValue)
// 		return
// 	}

// 	log.Printf("Procesing for  user ID : %d.\n", id)

// 	u, appErr := services.UserService.GetUser(id)
// 	if appErr != nil {
// 		resp.WriteHeader(appErr.Status)
// 		jsonvalue, _ := json.Marshal(appErr)
// 		resp.Write(jsonvalue)
// 		return
// 	}

// 	jsonvalue, _ := json.Marshal(u)
// 	resp.Write(jsonvalue)
// }
