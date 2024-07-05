package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"tani-hub-v2/constant"
	"tani-hub-v2/database"
	"tani-hub-v2/structs"
	"time"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user structs.User

		sql := "SELECT * FROM users WHERE id = $1 LIMIT 1"

		err := database.DbConnection.QueryRow(sql, claims["sub"]).Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			panic(err)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func RequireAuthUser(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user structs.User

		sql := "SELECT * FROM users WHERE id = $1 LIMIT 1"

		err := database.DbConnection.QueryRow(sql, claims["sub"]).Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			panic(err)
		}

		if claims["role"] != constant.USER && user.Role != constant.USER {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func RequireAuthAdmin(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user structs.User

		sql := "SELECT * FROM users WHERE id = $1 LIMIT 1"

		err := database.DbConnection.QueryRow(sql, claims["sub"]).Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			panic(err)
		}

		if claims["role"] != constant.ADMIN && user.Role != constant.ADMIN {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
