package tests

import (
	"encoding/json"
	"github.com/revel/revel/testing"
	"github.com/tcurtil/go-fizz-buzz/app/models"
	"log"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestFizzBuzzEndPoint() {
	t.Get("/resetStats")
	t.AssertOk()

	t.Get("/fizzbuzz?limit=10&int1=2&int2=3&str1=f&str2=b")
	t.AssertEqual("application/json", t.Response.Header.Get("Content-Type"))

	var content []string
	println(t.ResponseBody)
	err := json.Unmarshal(t.ResponseBody, &content)
	if err != nil {
		log.Fatal("JSON decode error: ", err)
	}
	var delta = 1
	t.AssertEqual(10, len(content))
	t.AssertEqual("1", content[1-delta])
	t.AssertEqual("f", content[2-delta])
	t.AssertEqual("b", content[3-delta])
	t.AssertEqual("f", content[4-delta])
	t.AssertEqual("5", content[5-delta])
	t.AssertEqual("fb", content[6-delta])
	t.AssertEqual("7", content[7-delta])
	t.AssertEqual("f", content[8-delta])
	t.AssertEqual("b", content[9-delta])
	t.AssertEqual("f", content[10-delta])
}

func (t *AppTest) TestStatEndPoint() {
	t.Get("/resetStats")
	t.AssertOk()

	t.Get("/fizzbuzz?limit=10&int1=2&int2=3&str1=t&str2=s")
	t.AssertOk()

	t.Get("/stats")
	t.AssertOk()

	var content models.RequestStat
	err1 := json.Unmarshal(t.ResponseBody, &content)
	if err1 != nil {
		log.Fatal("JSON decode error: ", err1)
	}
	t.AssertEqual(1, content.HitCount)
	t.AssertEqual(10, content.Request.Limit)
	t.AssertEqual(2, content.Request.Int1)
	t.AssertEqual(3, content.Request.Int2)
	t.AssertEqual("t", content.Request.Str1)
	t.AssertEqual("s", content.Request.Str2)

	t.Get("/fizzbuzz?limit=11&int1=2&int2=3&str1=t&str2=s")
	t.AssertOk()

	t.Get("/fizzbuzz?limit=11&int1=2&int2=3&str1=t&str2=s")
	t.AssertOk()

	t.Get("/stats")
	t.AssertOk()

	err := json.Unmarshal(t.ResponseBody, &content)
	if err != nil {
		log.Fatal("JSON decode error: ", err)
	}
	t.AssertEqual(2, content.HitCount)
	t.AssertEqual(11, content.Request.Limit)
	t.AssertEqual(2, content.Request.Int1)
	t.AssertEqual(3, content.Request.Int2)
	t.AssertEqual("t", content.Request.Str1)
	t.AssertEqual("s", content.Request.Str2)
}

func (t *AppTest) After() {
	println("Tear down")
}
