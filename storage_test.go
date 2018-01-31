package sqlite3utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func execSQLite(filename string, queries []string) {
	script, _ := filepath.Abs("./script/sqlite.rb")
	escape := regexp.MustCompile(`'`)
	for _, q := range queries {
		q = escape.ReplaceAllString(q, "\\'")
		//fmt.Print("Query> " + q)
		//out, err := exec.Command("ruby", script, filename, q).Output()
		err := exec.Command("ruby", script, filename, q).Run()
		//fmt.Print("Result> " + string(out))
		if err != nil {
			panic(err)
		}
	}
}

func rmSQLite(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		err := exec.Command("rm", filename).Run()
		if err != nil {
			panic(err)
		}
	}
}

func TestSimpleLoad(t *testing.T) {
	filename := "/tmp/test.db"
	rmSQLite(filename)

	execSQLite(filename, []string{
		"CREATE TABLE person(id integer, name text);",
		"INSERT INTO person VALUES (1, \"hoge\");",
		"INSERT INTO person VALUES (2, \"foo\");",
		"INSERT INTO person VALUES (3, \"bar\");",
	})

	pages, _ := Load(filename)
	assert.Equal(t, pages.Tables["person"].Entries[0].Datas[0].Value, "1", "Should be same")
	assert.Equal(t, pages.Tables["person"].Entries[0].Datas[1].Value, "hoge", "Should be same")
	assert.Equal(t, pages.Tables["person"].Entries[1].Datas[0].Value, "2", "Should be same")
	assert.Equal(t, pages.Tables["person"].Entries[1].Datas[1].Value, "foo", "Should be same")
	assert.Equal(t, pages.Tables["person"].Entries[2].Datas[0].Value, "3", "Should be same")
	assert.Equal(t, pages.Tables["person"].Entries[2].Datas[1].Value, "bar", "Should be same")

	rmSQLite(filename)
}