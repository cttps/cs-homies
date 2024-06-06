package controllers

import (
	"encoding/hex"
	"net/http"

	"crypto/hmac"
	"crypto/sha256"
	"fmt"

	"github.com/gin-gonic/gin"
)

const curSeed string = "errPGoUfeA"

func hashCode(code string, seed string) string {
	key := []byte(seed)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(code))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckHash(code, seed, hashedCode string) bool {
	return hmac.Equal([]byte(hashCode(code, seed)), []byte(hashedCode))
}

func VerifyUser(c *gin.Context) {

	var body struct {
		LoginCode string `json:"logincode" form:"logincode"`
	}

	if err := c.Bind(&body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		panic(err)
	}

	hashedCode := hashCode(body.LoginCode, curSeed)

	// if err := bcrypt.CompareHashAndPassword([]byte(body.LoginCode), []byte(hashedCode)); err != nil {
	// 	c.String(http.StatusBadRequest, err.Error())
	// }

	if CheckHash(body.LoginCode, curSeed, "015ea09e27ebfe3c86edf38c772baa994652a6c41816eb5065bb24669273aa91") {
		c.String(http.StatusAccepted, "Loading...")
	} else {
		fmt.Println(hashedCode)
		c.String(http.StatusBadRequest, "User ID Mismatch Error") // htmx doesnt swap with status 400, 500 etc
	}

}
