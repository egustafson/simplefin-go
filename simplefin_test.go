package simplefin_test

import (
	"embed"
	"encoding/json"
	"io/fs"
	"testing"

	"github.com/egustafson/simplefin-go"
)

//go:embed testdata/*.json
var testdata embed.FS

func TestErrors(t *testing.T) {
	for _, testfile := range matchTestFiles(t, "testdata/error*.json") {
		t.Run(testfile, func(t *testing.T) {
			testLoadAccountsResp(t, testfile)
		})
	}
}

func TestAccounts(t *testing.T) {
	for _, testfile := range matchTestFiles(t, "testdata/account*.json") {
		t.Run(testfile, func(t *testing.T) {
			testLoadAccountsResp(t, testfile)
		})
	}
}

func testLoadAccountsResp(t *testing.T, testfile string) {
	data, err := testdata.ReadFile(testfile)
	if err != nil {
		t.Errorf("%s: %v", testfile, err)
		return
	}
	var resp simplefin.AccountsResp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		t.Errorf("%s: %v", testfile, err)
		return
	}
}

func matchTestFiles(t *testing.T, glob string) []string {
	matches, err := fs.Glob(testdata, glob)
	if err != nil {
		t.FailNow()
	}
	return matches
}
