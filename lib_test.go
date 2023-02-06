package ugodict

import (
	"context"
	"fmt"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"testing"
)

// Loads the .env file in our repo - if you have one
func loadTestEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func TestDefineByTerm(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	term := os.Getenv("VAL_TERM")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewUrbanClient(hClient)
	res, err := c.DefineByTerm(context.Background(), term)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestDefineByID(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	definitionId, err := strconv.Atoi(os.Getenv("VAL_DEFID"))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewUrbanClient(hClient)
	res, err := c.DefineById(context.Background(), definitionId)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestRandomDefinition(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewUrbanClient(hClient)
	res, err := c.RandomDefinition(context.Background())
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestAutoComplete(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	term := os.Getenv("VAL_TERM")
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewUrbanClient(hClient)
	res, err := c.AutoComplete(context.Background(), term)
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}

func TestVote(t *testing.T) {
	// Load test env and values from env
	loadTestEnv()
	definitionId, err := strconv.Atoi(os.Getenv("VAL_DEFID"))
	if err != nil {
		t.Error(err)
		panic(err)
	}
	up := os.Getenv("VAL_DIRECTION_UP") == "true"
	hClient, err := gokhttp.TestHTTPClient()
	if err != nil {
		t.Error(err)
		panic(err)
	}

	c := NewUrbanClient(hClient)
	res, err := c.Vote(context.Background(), NewVoteRequestBody(definitionId, up))
	if err == nil {
		fmt.Println(spew.Sdump(res))
	} else {
		fmt.Println("err: ", err)
	}
}
