// Beego (http://beego.me/)
// @description beego is an open-source, high-performance web framework for the Go programming language.
// @link        http://github.com/DiogoDoreto/beego for the canonical source repository
// @license     http://github.com/DiogoDoreto/beego/blob/master/LICENSE
// @authors     astaxie

package beego

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/DiogoDoreto/beego/context"
)

func TestNamespaceGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/user", nil)
	w := httptest.NewRecorder()

	ns := NewNamespace("/v1")
	ns.Get("/user", func(ctx *context.Context) {
		ctx.Output.Body([]byte("v1_user"))
	})
	ns.ServeHTTP(w, r)
	if w.Body.String() != "v1_user" {
		t.Errorf("TestNamespaceGet can't run, get the response is " + w.Body.String())
	}
}

func TestNamespacePost(t *testing.T) {
	r, _ := http.NewRequest("POST", "/v1/user/123", nil)
	w := httptest.NewRecorder()

	ns := NewNamespace("/v1")
	ns.Post("/user/:id", func(ctx *context.Context) {
		ctx.Output.Body([]byte(ctx.Input.Param(":id")))
	})
	ns.ServeHTTP(w, r)
	if w.Body.String() != "123" {
		t.Errorf("TestNamespacePost can't run, get the response is " + w.Body.String())
	}
}

func TestNamespaceNest(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/admin/order", nil)
	w := httptest.NewRecorder()

	ns := NewNamespace("/v1")
	ns.Namespace(
		NewNamespace("/admin").
			Get("/order", func(ctx *context.Context) {
			ctx.Output.Body([]byte("order"))
		}),
	)
	ns.ServeHTTP(w, r)
	if w.Body.String() != "order" {
		t.Errorf("TestNamespaceNest can't run, get the response is " + w.Body.String())
	}
}

func TestNamespaceNestParam(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/admin/order/123", nil)
	w := httptest.NewRecorder()

	ns := NewNamespace("/v1")
	ns.Namespace(
		NewNamespace("/admin").
			Get("/order/:id", func(ctx *context.Context) {
			ctx.Output.Body([]byte(ctx.Input.Param(":id")))
		}),
	)
	ns.ServeHTTP(w, r)
	if w.Body.String() != "123" {
		t.Errorf("TestNamespaceNestParam can't run, get the response is " + w.Body.String())
	}
}

func TestNamespaceFilter(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/user/123", nil)
	w := httptest.NewRecorder()

	ns := NewNamespace("/v1")
	ns.Filter("before", func(ctx *context.Context) {
		ctx.Output.Body([]byte("this is Filter"))
	}).
		Get("/user/:id", func(ctx *context.Context) {
		ctx.Output.Body([]byte(ctx.Input.Param(":id")))
	})
	ns.ServeHTTP(w, r)
	if w.Body.String() != "this is Filter" {
		t.Errorf("TestNamespaceFilter can't run, get the response is " + w.Body.String())
	}
}

func TestNamespaceRouter(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/api/list", nil)
	w := httptest.NewRecorder()

	ns := NewNamespace("/v1")
	ns.Router("/api/list", &TestController{}, "*:List")
	ns.ServeHTTP(w, r)
	if w.Body.String() != "i am list" {
		t.Errorf("TestNamespaceRouter can't run, get the response is " + w.Body.String())
	}
}

func TestNamespaceAutoFunc(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/test/list", nil)
	w := httptest.NewRecorder()

	ns := NewNamespace("/v1")
	ns.AutoRouter(&TestController{})
	ns.ServeHTTP(w, r)
	if w.Body.String() != "i am list" {
		t.Errorf("user define func can't run")
	}
}

func TestNamespaceCond(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/test/list", nil)
	w := httptest.NewRecorder()

	ns := NewNamespace("/v1")
	ns.Cond(func(ctx *context.Context) bool {
		if ctx.Input.Domain() == "beego.me" {
			return true
		}
		return false
	}).
		AutoRouter(&TestController{})
	ns.ServeHTTP(w, r)
	if w.Code != 405 {
		t.Errorf("TestNamespaceCond can't run get the result " + strconv.Itoa(w.Code))
	}
}
