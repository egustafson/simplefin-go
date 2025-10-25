package simplefin

// file client.go implements the (REST like) client behaviors

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// ErrBadAccessURL is returned when the access URL is invalid.
var ErrBadAccessURL = errors.New("the access URL is neither file:// nor https://")

// Claim takes a SimpleFIN (one time use) token and converts it into a reusable
// "Access URL", which is used by the SimpleFin client to connect to the server.
func Claim(token string) (accessURL string, err error) {

	urlBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	claimURL := string(urlBytes)
	if _, err := url.ParseRequestURI(claimURL); err != nil {
		return "", err
	}

	resp, err := http.Post(claimURL, "text/plain", strings.NewReader(""))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// simpleFin is the implementation of the SimpleFin interface.
type simpleFin struct {
	accessURL  string
	fileMode   bool
	httpClient *http.Client
	// apiCalls int // future use: track number of API calls made.  SimpleFIN servers
	// limit the number of calls allowed to 24 per day.
	// (https://beta-bridge.simplefin.org/info/developers)
}

// New constructs a SimpleFin interface from an access URL.  SimpleFin requires https
// connections.  The 'insecure' parameter can be set to true for testing.  The file://
// scheme is also supported for development.
func New(accessURL string, insecure ...bool) (SimpleFin, error) {
	sf := new(simpleFin)
	if strings.HasPrefix(accessURL, "file://") {
		sf.accessURL = accessURL
		sf.fileMode = true
		return sf, nil
	}
	if len(insecure) > 0 && insecure[0] {
		sf.httpClient = &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	} else {
		sf.httpClient = &http.Client{}
	}
	if strings.HasPrefix(accessURL, "https://") {
		sf.accessURL = accessURL
		return sf, nil
	}
	return nil, ErrBadAccessURL
}

// Info queries a SimpleFIN server and returns the supported protocol versions.
func (sf *simpleFin) Info() (InfoResp, error) {

	var body []byte
	if !sf.fileMode {
		infoURL := sf.accessURL + "/info"
		resp, err := sf.httpClient.Get(infoURL)
		if err != nil {
			return InfoResp{}, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return InfoResp{}, err
		}
	} else { // if sf.fileMode
		var err error
		body, err = loadLocalFile(sf.accessURL)
		if err != nil {
			return InfoResp{}, err
		}
	}

	var infoResp InfoResp
	err := json.Unmarshal(body, &infoResp)
	if err != nil {
		return InfoResp{}, err
	}
	return infoResp, nil
}

// Accounts queries a SimpleFIN server and returns a list of accounts with an
// optional list of transactions attributed to each account.
func (sf *simpleFin) Accounts(options ...AccountOption) (AccountsResp, error) {

	var body []byte
	if !sf.fileMode {
		ap := processAccountOptions(options)

		accountsURL := sf.accessURL + "/accounts" + ap.ToQueryString()
		resp, err := sf.httpClient.Get(accountsURL)
		if err != nil {
			return AccountsResp{}, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return AccountsResp{}, errors.New("http response to /accounts: " + resp.Status)
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return AccountsResp{}, err
		}
	} else { // if sf.fileMode
		var err error
		body, err = loadLocalFile(sf.accessURL)
		if err != nil {
			return AccountsResp{}, err
		}
	}

	var acctResp AccountsResp
	err := json.Unmarshal(body, &acctResp)
	if err != nil {
		return AccountsResp{}, err
	}
	return acctResp, nil
}

func loadLocalFile(accessURL string) ([]byte, error) {

	if !strings.HasPrefix(accessURL, "file://") {
		return nil, ErrBadAccessURL
	}

	filePath := strings.TrimPrefix(accessURL, "file://")
	return os.ReadFile(filePath)
}
