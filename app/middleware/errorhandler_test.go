package middleware

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/axiaoxin/gin-skeleton/app/utils/response"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode("release")
}

func TestErrorHandler404(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.JSON(404, 1)
	ErrorHandler()(c)
	j := w.Body.Bytes()[1:] // 这里的测试方式有点问题，body中会保存c.JSON在前面返回的1，手动去掉
	r := response.Response{}
	err := json.Unmarshal(j, &r)
	if err != nil {
		t.Error(err)
	}
	if r.Code != 4 {
		t.Error("json code error")
	}
}

func TestErrorHandler500(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.JSON(500, 1)
	ErrorHandler()(c)
	j := w.Body.Bytes()[1:] // 这里的测试方式有点问题，body中会保存c.JSON在前面返回的1，手动去掉
	r := response.Response{}
	err := json.Unmarshal(j, &r)
	if err != nil {
		t.Error(err)
	}
	if r.Code != 5 {
		t.Error("json code error")
	}
}