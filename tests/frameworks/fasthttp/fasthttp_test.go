package fasthttp

import (
	"github.com/HongJaison/go-admin/tests/common"
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func TestFasthttp(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(newHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
