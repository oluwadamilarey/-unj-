package controllers

import "github.com/gin-gonic/gin"

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {
	return "gdhfflffh"
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	return false, "ghdjd"
}
