package businessflow

import (
	"fmt"
	"go/tutorial/config"
	"go/tutorial/exception"
	"go/tutorial/models/entity"
	"go/tutorial/stage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	User              entity.User
	Accounting        entity.Accounting
	PaymentAccounting entity.Accounting
}

func CreateUserBusinessFlow(c *gin.Context) {
	request := createRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Can't bind request"})
		exception.PanicLog(err)
		return
	}
	user := entity.User{
		UserName:  request.Name,
		TypeUser:  request.TypeUser,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	tx := config.DBTransaction()
	if err := stage.CreateUserStage(tx, &user); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadGateway, gin.H{"errorMessage": "Can't create user: " + err.Error()})
		return
	}

	fmt.Printf("userID : %v\n", user.ID)

	_, err := CreateAccounting(tx, user)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadGateway, gin.H{"errorMessage": "Error while create accouting: " + err.Error()})
		return
	}

	accoutingStage, err := PaymentProcess(tx, user)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadGateway, gin.H{"errorMessage": "Error while payment process: " + err.Error()})
		return
	}
	if err := config.DBCommit(tx); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"errorMessage": "Can't commit into database: " + err.Error()})
		return
	}
	resposne := response{
		User:              user,
		Accounting:        accoutingStage.OldAccouting,
		PaymentAccounting: accoutingStage.NewAccounting,
	}
	c.JSON(http.StatusCreated, resposne)
}

type createRequest struct {
	TypeUser uint   `json:"typeUser"`
	Name     string `json:"name"`
}
