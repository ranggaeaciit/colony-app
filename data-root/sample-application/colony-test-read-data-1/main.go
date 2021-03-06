package main

import (
	"fmt"
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"net/http"
	"os"
)

var (
	basePath   = func(dir string, err error) string { return dir }(os.Getwd())
	layoutFile = "_template.html"
)

func createResult(success bool, data interface{}, message string) map[string]interface{} {
	if !success {
		fmt.Println("ERROR! ", message)
	}

	return map[string]interface{}{
		"data":    data,
		"success": success,
		"message": message,
	}
}

type TestController struct {
}

func main() {
	server := new(knot.Server)
	server.Address = "localhost:8200"
	server.Register(new(TestController), "")
	server.Route("/", func(r *knot.WebContext) interface{} {
		http.Redirect(r.Writer, r.Request, "/test/index", 301)
		return true
	})
	server.Listen()
}

func (w *TestController) Index(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputTemplate
	r.Config.LayoutTemplate = layoutFile
	r.Config.ViewName = "index.html"

	return true
}

func (w *TestController) GetData(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson

	connectionInfo := &dbox.ConnectionInfo{"192.168.0.200:27017", "ecshell", "", "", nil}
	connection, err := dbox.NewConnection("mongo", connectionInfo)
	if err != nil {
		return createResult(false, nil, err.Error())
	}
	err = connection.Connect()
	if err != nil {
		return createResult(false, nil, err.Error())
	}

	c, err := connection.NewQuery().Select().From("WEISUserLog").Cursor(nil)
	if err != nil {
		return createResult(false, nil, err.Error())
	}

	data := []toolkit.M{}
	err = c.Fetch(&data, 0, true)
	if err != nil {
		return createResult(false, nil, err.Error())
	}

	return createResult(true, data, "")
}
