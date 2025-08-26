package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theabdullahishola/to-do/model"
	"github.com/theabdullahishola/to-do/util"
)

// Response structure from Google
type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func googleAuth(c *gin.Context) {

	var req struct {
		Code string `json:"code"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing code"})

		return
	}
	// Step 1: Exchange code for tokens

	token, err := util.ExchangeCode(req.Code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Google code"})
		return
	}

	// Step 2: Use access token to fetch user info
	client := util.GoogleOAuthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	fmt.Println("resp is ... ", resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	var gUser GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&gUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to decode user info"})
		return
	}

	// Step 3: Save/find user in DB (pseudo-code)
	// user, _ := model.FindOrCreateGoogleUser(gUser)

	user, err := model.CreateOrGetGoogleUser(gUser.Email, gUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to login with google"})
		return
	}
	ID := user.ID.Hex()
	email := user.Email

	// Step 4: Generate your own JWT
	accessToken, err := util.GenerateAccessToken(ID, email) // use your own function
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create token"})
		return
	}

	refreshToken, err := util.GenerateRefreshToken(ID, email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	c.SetCookie(
		"refresh_token", // name
		refreshToken,    // value
		24*60*60,        // maxAge (1 days in seconds)
		"/",             // path
		"",              // domain ("" = current domain)
		false,           // secure (true = only over HTTPS)
		true,            // httpOnly (not accessible by JS)
	)

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}
