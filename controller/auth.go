package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"tani-hub-v2/constant"
	"tani-hub-v2/database"
	"tani-hub-v2/structs"
	"time"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Name     string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	user := structs.User{Email: body.Email, Name: body.Name, Role: constant.USER, Password: string(hash)}
	sql := "INSERT INTO users (email, name, password, role)" +
		" VALUES ($1, $2, $3, $4)"

	errs := database.DbConnection.QueryRow(sql, user.Email, user.Name, user.Password, user.Role)
	if errs.Err() != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Signup User",
	})
}

func SignupAdmin(c *gin.Context) {
	var body struct {
		Email    string
		Name     string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	user := structs.User{Email: body.Email, Name: body.Name, Role: constant.ADMIN, Password: string(hash)}
	sql := "INSERT INTO users (email, name, password, role)" +
		" VALUES ($1, $2, $3, $4)"

	errs := database.DbConnection.QueryRow(sql, user.Email, user.Name, user.Password, user.Role)
	if errs.Err() != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Signup Admin",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user structs.User

	sql := "SELECT * FROM users WHERE email = $1 LIMIT 1"

	err := database.DbConnection.QueryRow(sql, body.Email).Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		panic(err)
	}

	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	errs := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.Id,
		"role": user.Role,
		//"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	//c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Login User",
	})
}

//
//func Validate(c *gin.Context) {
//	user, _ := c.Get("user")
//
//	// user.(models.User).Email    -->   to access specific data
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": user,
//	})
//}
