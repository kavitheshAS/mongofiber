package responses

import "github.com/gofiber/fiber/v2"

type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

//this creates a userresponse struct with sstatus.message,and data property to represnt the api responses type
//the json:"status" .. etc are called sstruct tags, these are used to attack meta data to the struct
