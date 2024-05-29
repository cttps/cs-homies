package middleware

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"websitetest1/controllers"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func RequireAuth(c *gin.Context) {
// 	fmt.Println("mw")

// 	tokenStr, err := c.Cookie("Authorization")
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		panic(err)
// 	}

// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return controllers.RandSecretString, nil
// 	})
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		panic(err)
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		if float64(time.Now().Unix()) > claims["exp"].(float64) { // expired token/cookie
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 		var user models.User
// 		controllers.DB.First(&user, claims["sub"])
// 		if float64(user.ID) != claims["sub"] {
// 			c.String(http.StatusUnauthorized, "User ID Mismatch unauth")
// 		}

// 		c.Set("user", user)

// 	} else {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		panic(err)
// 	}

// 	c.Next()
// }
