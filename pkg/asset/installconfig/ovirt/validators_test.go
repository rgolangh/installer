package ovirt

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/openshift/installer/pkg/types/ovirt"
)

var (
	okOvirtServer *httptest.Server
)

func init() {
	// mock ovirt server
	okOvirtServer = CreateMockOvirtServer(func(writer http.ResponseWriter, request *http.Request) {
		if strings.Contains(request.URL.Path, "ovirt-engine/services/pki-resource") {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("-----BEGIN CERTIFICATE-----\nFOO\n-----END CERTIFICATE-----\n;"))
			return
		}
		if strings.Contains(request.URL.Path, "/ovirt-engine/sso/oauth") {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("{}"))
			return
		}
	})
}

func Test_validateAuth(t *testing.T) {
	tests := []struct {
		url           string
		username      string
		password      string
		insecure      bool
		cafile        string
		expectSuccess bool
	}{{
		url:           okOvirtServer.URL,
		username:      "admin@internal",
		password:      "123",
		insecure:      false,
		cafile:        "",
		expectSuccess: true,
	},
		{
			url:           "https://nonexisting",
			username:      "foo",
			password:      "bar",
			insecure:      false,
			cafile:        "",
			expectSuccess: false,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			p := ovirt.Platform{
				URL:      test.url,
				Username: test.username,
				Password: test.password,
				Cafile:   test.cafile,
				Insecure: test.insecure,
			}

			validationFunc := Authenticated(&p)
			got := validationFunc(p.Password)
			assert.Equal(t, test.expectSuccess, got == nil, "got this %s", got)
			t.Log(got)
		})
	}
}

func CreateMockOvirtServer(handler http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(handler)

}
