package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"cafebean.xyz/cafeshort/v2/urlrepo"
	"github.com/gin-gonic/gin"
)

const SCHEME string = "http"
const HOST string = "localhost:8080"

type ResponseStructure struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func main() {
	r := gin.Default()
	var store urlrepo.URLRepository = urlrepo.NewInMemoryURLRepository()
	r.GET("/:id", func(ctx *gin.Context) {
		shortId, ok := ctx.Params.Get("id")
		if !ok || len(shortId) == 0 {
			ctx.JSON(http.StatusBadRequest, ResponseStructure{StatusCode: http.StatusBadRequest, Message: "No ID specified in URI"})
			return
		}
		longUrl, err := store.GetUrlById(shortId)
		if err != nil {
			switch {
			case errors.Is(err, urlrepo.ErrUrlNotExists):
				ctx.String(http.StatusNotFound, "No URL found for id")
			default:
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
			log.Println(err)
			return
		}
		ctx.Redirect(http.StatusSeeOther, longUrl)
	})
	r.POST("/create", func(ctx *gin.Context) {
		// get url from form data
		longUrl := ctx.Request.FormValue("url")
		urlRegex, err := regexp.Compile(`http[s]*://([a-z]*|[A-Z]*|\\.*)(/(.)*)*`)
		if err != nil {
			fmt.Println("Error compiling regex: ", err)
			ctx.JSON(http.StatusInternalServerError, ResponseStructure{
				StatusCode: http.StatusInternalServerError,
				Message:    "Something went wrong inside",
			})
			return
		}
		if !urlRegex.Match([]byte(longUrl)) {
			fmt.Println("URL did not match regex")
			ctx.JSON(http.StatusBadRequest, ResponseStructure{
				StatusCode: http.StatusBadRequest,
				Message:    "The given URL is not a valid HTTP(S) URL.",
			})
			return
		}
		// store the url
		shortId, err := store.AddUrl(longUrl)
		if err != nil {
			fmt.Println("Error storing URL: ", err)
			ctx.JSON(http.StatusInternalServerError, ResponseStructure{StatusCode: http.StatusInternalServerError, Message: "Something went wrong inside"})
			return
		}
		// return the short url in plaintext
		shortUrl := fmt.Sprintf("%s://%s/%s", SCHEME, HOST, shortId)
		ctx.String(http.StatusCreated, shortUrl)
	})
	r.Run()
}
