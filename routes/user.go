package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theabdullahishola/to-do/model"
	"github.com/theabdullahishola/to-do/util"
)

func signUp(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cant parse data"})
		return
	}
	err = user.NewUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cant sign up"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
func login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cant parse data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}
	accessToken, err := util.GenerateAccessToken(user.ID.Hex(), user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	refreshToken, err := util.GenerateRefreshToken(user.ID.Hex(), user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		MaxAge:   24 * 60 * 60,       // 1 day
		Path:     "/",
		Domain:   "",                 // "" = current domain
		Secure:   true,               // true in prod (HTTPS), false in local dev
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // 🔑 allow cross-site (frontend <> backend)
	})
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken})

}
func refreshToken(c *gin.Context) {

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing refresh token"})
		return
	}

	userID, err := util.VerifyRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid refresh token"})
		return
	}
	user, err := model.GetUserbyID(userID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid refresh token"})
		return
	}
	newAccessToken, err := util.GenerateAccessToken(userID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate new access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": newAccessToken})
}
func logout(c *gin.Context) {
	// Overwrite the refresh token cookie with empty value + expired time
	c.SetCookie(
		"refresh_token",
		"",
		-1, // expire immediately
		"/",
		"",
		false,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
