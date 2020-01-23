package routes

import (
	// "fmt"
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

	router.GET("/unknown", get_unknown)
	router.GET("/maybeknown", get_maybeknown)

	return router
}

func Listen(port int) {
	router := createRouter()
	router.Run(":" + strconv.Itoa(port))
}

func get_unknown(c *gin.Context) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	stars_l := rng.Intn(20)
	stars_u := 20 + rng.Intn(200)

	request_query := c.Request.URL.Query()

	lang_param := ""
	if request_query["language"] != nil {
		lang_param = "language:" + request_query["language"][0] + "+"
	}

	q := utils.Query{
		"q": lang_param + "stars:" + strconv.Itoa(stars_l) + ".." + strconv.Itoa(stars_u) + "+is:public",
		"sort": "updated",
		"order": "desc",
	}

	r := parseResponse(request(utils.BuildRequestString(q), c), c)
	if r != nil {
		c.JSON(http.StatusOK, r)
	}
}

func get_maybeknown(c *gin.Context) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	stars := 100 + rng.Intn(150)

	request_query := c.Request.URL.Query()

	lang_param := ""
	if request_query["language"] != nil {
		lang_param = "language:" + request_query["language"][0] + "+"
	}

	q := utils.Query{
		"q": lang_param + "stars:>" + strconv.Itoa(stars) + "+is:public",
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
	json.Unmarshal([]byte(bodyBytes), &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return nil;
	}

	return t.ToMap();
}
