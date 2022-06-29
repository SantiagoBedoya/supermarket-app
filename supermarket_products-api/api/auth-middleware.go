package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func IsAuth(c *gin.Context) {
	bearerToken, err := c.Cookie("jwt")
	if err != nil {
		log.Printf("%+v\n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	if bearerToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	rawURL := fmt.Sprintf("%s/verify-token", os.Getenv("ACCOUNT_API_URL"))
	urlObj, err := url.Parse(rawURL)
	if err != nil {
		log.Printf("%+v\n", err)
		restErr := ParseError(err)
		c.AbortWithStatusJSON(restErr.StatusCode, restErr)
		return
	}
	request, _ := http.NewRequest(http.MethodGet, urlObj.String(), nil)
	request.AddCookie(&http.Cookie{
		Name:  "jwt",
		Value: bearerToken,
	})
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		log.Printf("%+v\n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Printf("%+v\n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}
	byteData, _ := ioutil.ReadAll(response.Body)
	data := make(map[string]string)
	if err := json.Unmarshal(byteData, &data); err != nil {
		log.Printf("%+v\n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	c.Set("userFirstName", data["first_name"])
	c.Set("userLastName", data["last_name"])
	c.Set("userEmail", data["email"])
	c.Set("userRole", data["role"])
	c.Next()
}
