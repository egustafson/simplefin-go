package simplefin

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClaim_ValidToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("access-url-response"))
	}))
	defer server.Close()

	token := base64.StdEncoding.EncodeToString([]byte(server.URL))
	accessURL, err := Claim(token)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if accessURL != "access-url-response" {
		t.Errorf("expected 'access-url-response', got %q", accessURL)
	}
}

func TestClaim_InvalidBase64(t *testing.T) {
	_, err := Claim("not-base64!!")
	if err == nil {
		t.Error("expected error for invalid base64 token")
	}
}

func TestClaim_InvalidURL(t *testing.T) {
	badURL := base64.StdEncoding.EncodeToString([]byte("::not-a-url::"))
	_, err := Claim(badURL)
	if err == nil {
		t.Error("expected error for invalid URL")
	}
}

func TestNew_BadAccessURL(t *testing.T) {
	_, err := New("")
	if err != ErrBadAccessURL {
		t.Errorf("expected ErrBadAccessURL, got %v", err)
	}
}

func TestNewAndInfo(t *testing.T) {
	infoJSON := `{"versions":["1.0"]}`
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/info" {
			w.Write([]byte(infoJSON))
		} else {
			w.WriteHeader(404)
		}
	}))
	defer server.Close()

	sf, err := New(server.URL, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	info, err := sf.Info()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(info.Versions) != 1 || info.Versions[0] != "1.0" {
		t.Errorf("unexpected info response: %+v", info)
	}
}

func TestAccounts(t *testing.T) {
	accountsJSON := `{"accounts":[{"id":"a1"}]}`
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/accounts" {
			w.Write([]byte(accountsJSON))
		} else {
			w.WriteHeader(404)
		}
	}))
	defer server.Close()

	sf, err := New(server.URL, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	resp, err := sf.Accounts()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(resp.Accounts) != 1 || resp.Accounts[0].ID != "a1" {
		t.Errorf("unexpected accounts response: %+v", resp)
	}
}
