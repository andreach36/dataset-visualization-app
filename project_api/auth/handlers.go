package auth

import (
	"errors"
	"net/http"
	"os"
	"time"
	"project_api/database"
	"project_api/models"
	"project_api/shared"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var userInput models.UserInput

	// validar si el input ingresado es válido
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user := models.User{
		Username: userInput.Username,
		Password: userInput.Password,
	}
	// conecto con la base de datos
	if tx := database.CreateDbConnection().Create(&user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
			return
		}

	}

	// se inicializan los filtros predefinidos
	defaultFilter := models.Filter{
		UserID:         user.ID,
		MinAge:         "0", // Valores predeterminados
		MaxAge:         "100",
		Education:      "",
		MaritalStatus:  "",
		Occupation:     "",
		Income:         "",
		OrderBy:        "",
		OrderDirection: "",
	}
	// se crean los filtros predefinidos en la base de datos
	if err := database.CreateDbConnection().Create(&defaultFilter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create default filter"})
		return
	}

	// generar token session
	sessionToken := uuid.NewV5(uuid.UUID{}, "session").String()
	session := shared.Session{
		Uid:        user.ID,
		ExpiryTime: time.Now().Add(10 * time.Minute),
	}
	// almacenar la sesion
	shared.Sessions[sessionToken] = session

	// se crea el claim
	claims := shared.Payload{
		MapClaims: jwt.MapClaims{
			"iat": jwt.NewNumericDate(time.Now()),                       // issued at
			"eat": jwt.NewNumericDate(time.Now().Add(60 * time.Minute)), // expired at
		},
		Session: sessionToken,
	}

	// generar token con jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signinKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	tokenString, err := token.SignedString(signinKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"username": user.Username,
		"id":       user.ID,
		"session":  tokenString,
		"message":  "Succesful registration!",
	})
}

func Login(c *gin.Context) {
	var userInput models.UserInput
	var user models.User
	// validar si el input ingresado es válido
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	// conecto con la base de datos
	database.CreateDbConnection().Where("username=?", userInput.Username).Find(&user)
	// verificar la contraseña
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// generar token session
	sessionToken := uuid.NewV5(uuid.UUID{}, "session").String()
	session := shared.Session{
		Uid:        user.ID,
		ExpiryTime: time.Now().Add(10 * time.Minute),
	}
	// almacenar la sesion
	shared.Sessions[sessionToken] = session

	// se crea el claim
	claims := shared.Payload{
		MapClaims: jwt.MapClaims{
			"iat": jwt.NewNumericDate(time.Now()),                       // issued at
			"eat": jwt.NewNumericDate(time.Now().Add(60 * time.Minute)), // expired at
		},
		Session: sessionToken,
	}

	// generar token con jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signinKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	tokenString, err := token.SignedString(signinKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"session": tokenString,
		"message": "Login succesfully!",
	})

}

func Logout(c *gin.Context) {
	// obtener el usuario autorizado del contexto
	user, exists := c.Get("authorizedUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Obtener los claims de sesión del contexto
	sessionClaims, ok := c.Get("sessionClaims")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		return
	}

	// Convertir el valor de sessionClaims al tipo *shared.Payload
	claims, ok := sessionClaims.(*shared.Payload)
	if !ok || claims.Session == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session claims"})
		return
	}

	// remover la sesion de la memoria
	delete(shared.Sessions, claims.Session)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout succefully",
		"user":    user,
	})
}
