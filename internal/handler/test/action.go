package test

import (
	cache "github.com/armnerd/go-skeleton/pkg/redis"
	auth "github.com/armnerd/go-skeleton/pkg/auth"
	curl "github.com/armnerd/go-skeleton/pkg/curl"
	syslog "github.com/armnerd/go-skeleton/pkg/log"
	response "github.com/armnerd/go-skeleton/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// SetCache redis插入
func SetCache(c *gin.Context) {
	// 参数验证
	key := c.DefaultPostForm("key", "")
	value := c.DefaultPostForm("value", "")
	if key == "" || value == "" {
		response.Fail(c, response.ParamsLost)
		return
	}
	res, _ := cache.Get().Do("SET", key, value)
	response.Succuss(c, res)
}

// GetCache redis插入
func GetCache(c *gin.Context) {
	// 参数验证
	key := c.DefaultPostForm("key", "")
	if key == "" {
		response.Fail(c, response.ParamsLost)
		return
	}
	res, _ := redis.String(cache.Get().Do("GET", key))
	response.Succuss(c, res)
}

// CurlGet 请求
func CurlGet(c *gin.Context) {
	var url = "http://127.0.0.1:9551"
	data := map[string]interface{}{}
	header := map[string]interface{}{}
	content, err := curl.Get(url, data, header)
	if err != nil {
		response.Fail(c, response.RequestFail)
	}
	res := content.Get("Welcome").Value()
	// info日志
	syslog.Info("curl get test", "")
	response.Succuss(c, res)
}

// CurlPost 请求
func CurlPost(c *gin.Context) {
	var url = "http://127.0.0.1:9551/api/article/info"
	data := map[string]interface{}{
		"id": "95",
	}
	header := map[string]interface{}{}
	content, err := curl.Post(url, data, header)
	if err != nil {
		response.Fail(c, response.RequestFail)
	}
	res := content.Get("data").Get("Author").Value()
	// debug日志
	syslog.Debug("curl post test", "")
	response.Succuss(c, res)
}

// Login 登录
func Login(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	if username == "" || password == "" {
		response.Fail(c, response.ParamsLost)
		return
	}
	token := auth.GetToken(1)
	response.Succuss(c, token)
}

// Auth 鉴权
func Auth(c *gin.Context) {
	user, _ := c.Get("user")
	response.Succuss(c, user)
}
