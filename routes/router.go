package routes

import (
	"net/http"
	"io/ioutil"

	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/TrAyZeN/gimme-repos/models"
)

func createRouter() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	})

	router.GET("reallyunknown", getReallyunknown)
	router.GET("/unknown", getUnknown)
	router.GET("/maybeknown", getMaybeknown)

	return router
}

func Listen(port string) {
	router := createRouter()
	router.Run(":" + port)
}

func request(url string) (models.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return models.Response{}, err
	}

	r, err := parseResponse(res)
	if err != nil {
		return models.Response{}, err
	}

	return r, nil
}

func parseResponse(res *http.Response) (models.Response, error) {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return models.Response{}, err
	}

	var t models.Response
	err = json.Unmarshal([]byte(bodyBytes), &t)
	if err != nil {
		return models.Response{}, err
	}

	return t, nil
}
