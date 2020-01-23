package routes

import (
	"strconv"

	"math/rand"
	"time"
	"net/http"
	"io/ioutil"

	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/TrAyZeN/gimme-repos/utils"
	"github.com/TrAyZeN/gimme-repos/models"
)

func createRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/unknown", getUnknown)
	router.GET("/maybeknown", getMaybeknown)

	return router
}

func Listen(port int) {
	router := createRouter()
	router.Run(":" + strconv.Itoa(port))
}

func getUnknown(c *gin.Context) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	starsLowerBound := rng.Intn(20)
	starsUpperBound := 20 + rng.Intn(200)

	requestQuery := c.Request.URL.Query()

	langParam := ""
	if requestQuery["language"] != nil {
		langParam = "language:" + requestQuery["language"][0] + "+"
	}

	q := utils.Query{
		"q": langParam + "stars:" + strconv.Itoa(starsLowerBound) + ".." + strconv.Itoa(starsUpperBound) + "+is:public",
		"sort": "updated",
		"order": "desc",
	}

	r := parseResponse(request(utils.BuildRequestString(q), c), c)
	if r != nil {
		c.JSON(http.StatusOK, r)
	}
}

func getMaybeknown(c *gin.Context) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	stars := 100 + rng.Intn(150)

	requestQuery := c.Request.URL.Query()

	langParam := ""
	if requestQuery["language"] != nil {
		langParam = "language:" + requestQuery["language"][0] + "+"
	}

	q := utils.Query{
		"q": langParam + "stars:>" + strconv.Itoa(stars) + "+is:public",
		"sort": "updated",
		"order": "desc",
	}

	r := parseResponse(request(utils.BuildRequestString(q), c), c)
	if r != nil {
		c.JSON(http.StatusOK, r)
	}
}

func request(url string, c *gin.Context) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return nil;
	}

	return res;
}

func parseResponse(res *http.Response, c *gin.Context) map[string]interface{} {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return nil;
	}

	var t models.Response
	err = json.Unmarshal([]byte(bodyBytes), &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return nil;
	}

	return t.ToMap();
}
