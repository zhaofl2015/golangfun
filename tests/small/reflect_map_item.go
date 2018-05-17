package main

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	m := make(map[string]interface{})
	m["shuz"] = []int{1, 2, 3}
	m["obj"] = bson.NewObjectId()
	m["map"] = make(map[string]string)

	//	fmt.Println(reflect.TypeOf(m["obj"]).Name() == "ObjectId")
	//	if bson.ObjectId.Valid(reflect.ValueOf(m["obj"])) {
	//		fmt.Println("hahaha")
	//	}

	//	fmt.Println(reflect.TypeOf(m["shuz"]).Name())
	//	fmt.Println(reflect.TypeOf(m["obj"]).Name())
	//	fmt.Println(reflect.TypeOf(m["map"]).Name())

	//	fmt.Println(reflect.ValueOf(m["shuz"]).Kind())
	//	fmt.Println(reflect.ValueOf(m["obj"]).Kind())
	//	fmt.Println(reflect.ValueOf(m["map"]).Kind())

	for k, v := range m {
		if reflect.TypeOf(v).Name() == "ObjectId" {
			fmt.Printf("%v is objectid\n", k)
		}

		switch reflect.ValueOf(v).Kind().String() {
		case "slice":
			fmt.Printf("%v is slice\n", k)
		case "map":
			fmt.Printf("%v is map\n", k)
		}
	}

	//	for k, v := range m {
	//		switch
	//		switch reflect.TypeOf(v) {
	//		case bson.ObjectId:
	//			fmt.Println("%s is objectid", k)
	//		case reflect.Map:
	//			fmt.Println("%s is map", k)
	//		case reflect.Slice:
	//		case reflect.Array:
	//			fmt.Println("%s is array")
	//		}
	//	}
}
