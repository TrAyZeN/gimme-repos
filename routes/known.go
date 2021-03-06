package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TrAyZeN/gimme-repos/utils"
)

func getKnown(c *gin.Context) {
	requestQuery := c.Request.URL.Query()

	langParam := ""
	if requestQuery["language"] != nil {
		langParam = "language:" + requestQuery["language"][0] + "+"
	}

	q := utils.Query{
		"q": langParam + "stars:>1000+is:public",
		"sort": "updated",
		"order": "desc",
	}

	res, err := request(utils.BuildRequestString(q))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
