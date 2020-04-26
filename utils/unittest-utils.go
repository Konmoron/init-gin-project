package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// References:
// 	1. https://studygolang.com/articles/11836

// UnittestParseToStr 将map中的键值对输出成querystring形式
func UnittestParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

// UnittestGet 根据特定请求uri，发起get请求，返回响应
func UnittestGet(uri string, router *gin.Engine) (*http.Response, []byte) {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)

	return result, body
}

// UnittestDelete 根据特定请求uri，发起 delete 请求，返回响应
func UnittestDelete(uri string, router *gin.Engine) (*http.Response, []byte) {
	// 构造get请求
	req := httptest.NewRequest("DELETE", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)

	return result, body
}

// UnittestPostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func UnittestPostForm(uri string, param map[string]string, router *gin.Engine) (*http.Response, []byte) {
	// 构造post请求，表单数据以querystring的形式加在uri之后
	req := httptest.NewRequest("POST", uri+UnittestParseToStr(param), nil)

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return result, body
}

// UnittestPostJSON 根据特定请求uri和参数param，以Json形式传递参数，发起post请求返回响应
func UnittestPostJSON(uri string, param interface{}, router *gin.Engine) (*http.Response, []byte) {
	// 将参数转化为json比特流
	jsonByte, _ := json.Marshal(param)

	// 构造post请求，json数据以请求body的形式传递
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return result, body
}
