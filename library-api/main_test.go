package main

import (
	"testing"

	"github.com/Joshua-Pok/library-api/internal/gql"
	"github.com/graphql-go/graphql"
)

func TestEnv(t *testing.T) {
	if 1+1 != 2 {
		t.Error("Math is Broken")
	}
}

func TestBasic(t *testing.T) {

	queryString := "{ hello }"
	schema, err := schema.NewSchema()
	if err != nil {
		t.Error("Error creating schema")
	}

	params := graphql.Params{
		Schema:        schema,
		RequestString: queryString,
	}

	res := graphql.Do(params)
	//result comes with its own errors field
	if len(res.Errors) > 0 {
		t.Errorf("Error executing query: %v", res.Errors)
	}

	data := res.Data.(map[string]interface{}) // cast to map

	if _, ok := data["hello"]; !ok {
		t.Error("Error: Result does not contain key Hello with value World")
	}

	if key, _ := data["hello"]; key != "World" {
		t.Error("Error: Result value is not World")
	}

}

func TestForBook(t *testing.T) {
	queryString := "{ book(id: \"1\"){ id title }}"

	schema, err := schema.NewSchema()
	if err != nil {
		t.Error("Error creating schema: ", err)
	}

	params := graphql.Params{
		Schema:        schema,
		RequestString: queryString,
	}

	res := graphql.Do(params)
	if len(res.Errors) > 0 {
		t.Errorf("Error executing query: %v", err)
	}

	data := res.Data.(map[string]interface{})
	book, ok := data["book"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected 'book' field to be an object")
	}
	id, ok := book["id"].(string)
	if id != "1" {
		t.Errorf("Expected ID to be 1")
	}

}

func TestForNonExistentBook(t *testing.T) {
	queryString := "{ book(id: \"9999\"){id title}}"

	schema, err := schema.NewSchema()
	if err != nil {
		t.Error("Error creating schema", err)
	}

	params := graphql.Params{
		Schema:        schema,
		RequestString: queryString,
	}

	res := graphql.Do(params)
	if len(res.Errors) > 0 {
		t.Errorf("Error executing query: %v", err)
	}

	data := res.Data.(map[string]interface{})
	_, ok := data["book"].(map[string]interface{})
	if ok {
		t.Fatal("Book should not exist")
	}
}
