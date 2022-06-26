package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/loviatar101/order-consumer/database"
	"github.com/loviatar101/order-consumer/database/models"
)

//this is not the model(table)

//see this as a serialiser
type User struct{
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser(userModel models.User)User{
	return User{ID:userModel.ID,FirstName: userModel.FirstName,LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx)error{
	var user models.User
	if err:=c.BodyParser(&user);err !=nil{
		return c.Status(400).JSON(err.Error())

	}
	database.Database.Db.Create(&user)
	responseUser:=CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}


func GetUsers(c *fiber.Ctx)error{
	users:= []models.User{}
	database.Database.Db.Find(&users)
	responseUsers:=[]User{}
	for _,user:=range(users){
		responseUser:=CreateResponseUser(user)
		responseUsers=append(responseUsers,responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int,user *models.User)error{
	database.Database.Db.Find(&user,"id=?",id)
	if user.ID==0{
		return errors.New("user does not exist")

	}
	return nil
}

func GetUser(c *fiber.Ctx)error{
	id, err:=c.ParamsInt("id")
	var user models.User
	if err!=nil{
		return c.Status(400).JSON("please ensure that :id an integer")
		
	}

	if err:=findUser(id,&user);err!=nil{
		return c.Status(400).JSON(err.Error())
	}
	responseUser:=CreateResponseUser(user)


	return c.Status(200).JSON(responseUser)

}



func UpdateUser(c *fiber.Ctx)error{
	id, err:=c.ParamsInt("id")
	var user models.User
	if err!=nil{
		return c.Status(400).JSON("please ensure that :id an integer")
		
	}

	if err:=findUser(id,&user);err!=nil{
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct{
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}

	var updatedata UpdateUser
	if err:=c.BodyParser(&updatedata); err!=nil{
		return c.Status(500).JSON(err.Error())
	}
	user.FirstName=updatedata.FirstName
	user.LastName=updatedata.LastName

	database.Database.Db.Save(&user)


	responseuser:=CreateResponseUser(user)
	return c.Status(200).JSON(responseuser)
}



func DeleteUser(c *fiber.Ctx)error{
	id, err:=c.ParamsInt("id")
	var user models.User
	if err!=nil{
		return c.Status(400).JSON("please ensure that :id an integer")
		
	}

	if err:=findUser(id,&user);err!=nil{
		return c.Status(400).JSON(err.Error())
	}

	if err:=database.Database.Db.Delete(user).Error;err!=nil{
		return c.Status(404).JSON(err.Error())
	} 
	
	return c.Status(200).SendString("succefully deleted user")

	

}