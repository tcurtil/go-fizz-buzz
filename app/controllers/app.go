package controllers

import (
	"github.com/revel/revel"
	"strconv"
	"sync"
	"github.com/tcurtil/go-fizz-buzz/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

// this should be stored in a database to avoid filling the memory (DoS) and also
// allow multi-node design
var requestCounter = struct{
	sync.RWMutex
	Store map[models.FizzBuzzRequest]int
}{Store: make(map[models.FizzBuzzRequest]int)}

func (c App) FizzBuzz(limit int, int1 int, str1 string, int2 int, str2 string) revel.Result {
	c.Validation.Required(limit)
	c.Validation.Required(int1)
	c.Validation.Required(str1)
	c.Validation.Required(int2)
	c.Validation.Required(str2)
	if c.Validation.HasErrors() {
		c.Response.Status = 400
		return c.RenderText("request error. Check all 5 parameters were provided : limit int1 int2 str1 str2")
	}
	var key = models.FizzBuzzRequest{limit, int1, int2, str1, str2}
	requestCounter.Lock()
	requestCounter.Store[key]++
	requestCounter.Unlock()
	listing := make([]string, limit)
	for i := 1; i <= limit; i++ {
		var item string
		if i % int1 == 0 && i % int2 == 0 {
			item = str1 + str2
		} else if i % int1 == 0 {
			item = str1
		} else if i % int2 == 0 {
			item = str2
		} else {
			item = strconv.Itoa(i)
		}
		listing[i-1] = item
	}
	c.Response.ContentType = "application/json"
	return c.RenderJSON(listing)
}

func (c App) Statistics() revel.Result {
	requestCounter.RLock()
	var mostPopularRequest = models.RequestStat{HitCount:0}
	for req, cnt := range requestCounter.Store {
		if cnt > mostPopularRequest.HitCount {
			mostPopularRequest.Request = req
			mostPopularRequest.HitCount = cnt
		}
	}
	requestCounter.RUnlock()

	c.Response.ContentType = "application/json"
	return c.RenderJSON(mostPopularRequest)
}

func (c App) ResetStats() revel.Result {
	requestCounter.Lock()
	requestCounter.Store = make(map[models.FizzBuzzRequest]int)
	requestCounter.Unlock()
	return c.RenderText("done")
}