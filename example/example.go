package main

import (
	t "github.com/fwhezfwhez/elastic-go-tool"
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:name`
}

func main() {
	//创建客户端
	client, ctx, err := t.GetClient()
	if err != nil {
		panic(err)
	}

	//创建索引
	indexName := "test_elastic"
	mapping := `
	{
		"settings":{
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
		"mappings":{
			"user":{
				"properties":{
					"name":{
						"type":"keyword"
					}
				}
			}
		}
	}`
	index := t.Index{Index: indexName, Mapping: mapping}
	err = t.CreateIndex(client, ctx, index)
	if err != nil {
		panic(err)
	}

	//插入数据
	document1 := t.Document{Index: "test_elastic", Type: "user", Id: "1", Body: `{"name":"ft1"}`}
	document2 := t.Document{Index: "test_elastic", Type: "user", Id: "2", Body: `{"name":"ft1"}`}
	document3 := t.Document{Index: "test_elastic", Type: "user", Id: "3", Body: `{"name":"ft1"}`}
	document4 := t.Document{Index: "test_elastic", Type: "user", Id: "4", Body: `{"name":"ft1"}`}

	document5 := t.Document{Index: "test_elastic", Type: "user", Id: "5", Body: User{Name: "ft2"}}
	//插入记录
	err = t.InsertDocument(client, ctx, document1)
	if err != nil {
		panic(err)
	}
	err = t.InsertDocument(client, ctx, document2)
	if err != nil {
		panic(err)
	}
	err = t.InsertDocument(client, ctx, document3)
	if err != nil {
		panic(err)
	}
	err = t.InsertDocument(client, ctx, document4)
	if err != nil {
		panic(err)
	}
	err = t.InsertDocument(client, ctx, document5)
	if err != nil {
		panic(err)
	}

	//通过id获取数据
	document := t.Document{Index: "test_elastic", Type: "user", Id: "1"}
	result, err := t.GetDocument(client, ctx, document)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

	//删除一条document
	err = t.DeleteDocument(client, ctx, document)
	if err != nil {
		panic(err)
	}

	//查询documents
	termQuery := &t.TermSearch{
		ElemType:   reflect.TypeOf(User{}),
		Query:      t.QueryStruct{Key: "name", Value: "ft1"},
		Index:      "test_elastic",
		Type:       "user",
		SortField:  "name",
		Asc:        true,
		StartIndex: 0,
		QuerySize:  5,
	}
	//执行Search
	var results []interface{}
	results, err = t.SearchDocuments(client, ctx, termQuery)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(results, len(results))
	}

	//删除索引
	index = t.Index{Index: "test_elastic"}

	err = t.DeleteIndex(client, ctx, index)
	if err != nil {
		panic(err)
	}
}


