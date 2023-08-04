package controller

import (
	"context"
	"fmt"
	"net"
	"net/http"

	//"net/url"
	"time"
	"urlshortnerService/constant"
	"urlshortnerService/database"
	"urlshortnerService/helper"
	"urlshortnerService/types"

	"github.com/gin-gonic/gin"
	//"google.golang.org/genproto/googleapis/cloud/location"
	"gopkg.in/mgo.v2/bson"
)

func ShortTheUrl(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var shortUrlBody types.ShortUrlBody
	err := c.BindJSON(&shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": constant.BindError})
		return
	}
	code := helper.GenRandomString(6)

	//record,_:=  database.Mgr.GetUrlFromCode(code, constant.UrlCollection)
	




	//if record.UrlCode != "" {
		//c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "this code is already in use"})
		//return
	//}

	var url types.UrlDb

	url.CreatedAt = time.Now().Unix()
	url.ExpiredAt = time.Now().Unix()
	url.UrlCode = code
	url.LongUrl = shortUrlBody.LongUrl
	url.ShortUrl = constant.BaseUrl + code
	url.Locations = net.ParseIP(c.ClientIP()).String()
	//url.Clicks = database.Mgr.UpdateClicks(code, constant.UrlCollection).Clicks
	
	 result, err := database.GetCollection(database.ConnectToDB(), "url").InsertOne(ctx, url)
	 if err != nil {
	 	c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	 	return
	 }
	 fmt.Println(result)

	c.JSON(http.StatusOK, gin.H{"error": false, "data": result, "short_url": url.ShortUrl, "long_url": url.LongUrl})

	// resp, err := database.Mgr.Insert(url, constant.UrlCollection)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"error": false, "data": resp, "short_url": url.ShortUrl})
}

func RedirectURL(c *gin.Context) {
	
	code := c.Param("code")

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "there is no url found"})
		return
	}
	fmt.Println(record.LongUrl)
        click := database.Mgr.UpdateClicks(code, constant.UrlCollection)
	fmt.Println(click)
	location := c.Copy().ClientIP()
	
fmt.Println(location)


	c.Redirect(http.StatusPermanentRedirect, record.LongUrl)
}

func RedirectToURL(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	code := c.Param("code") 



	result := database.GetCollection(database.ConnectToDB(), "new").FindOne(ctx, bson.M{"url_code": code}).Decode(&types.UrlDb{})
	if result == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "there is no url found"})
		return
	}

	fmt.Println(result)
	url := types.UrlDb{}
	c.Redirect(http.StatusPermanentRedirect, url.LongUrl)




	
	

	
	
	

	
	//fmt.Println(result)
	//record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)
	

	//c.Redirect(http.StatusPermanentRedirect, result.LongUrl )

}
