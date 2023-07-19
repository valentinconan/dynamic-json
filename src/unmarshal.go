package src

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func UnmarshalJSONFromInterface[T any](data interface{}) T {
	dataAsByte, err := json.Marshal(data)
	if err != nil {
		log.Panic("Unable to manipulate data")
	}
	var typed T
	var structuredData any
	if 24 == (reflect.ValueOf(typed).Kind()) {
		//if value is type of string
		structuredData = string(dataAsByte)
	} else {
		//otherwise, all other values will be Unmarshal as T type
		structuredData = UnmarshalJSON[T](dataAsByte)
	}

	return structuredData.(T)
}

func UnmarshalJSON[T any](bs []byte) T {

	var structuredData map[string]interface{}

	if err := json.Unmarshal(bs, &structuredData); err != nil {
		panic(err)
	}
	var genericData T

	structValue := reflect.ValueOf(&genericData).Elem()

	deepConvert(structuredData, structValue)

	return genericData
}

func deepConvert(structuredData map[string]interface{}, structValue reflect.Value) {

	//get structure ex : models.GlobalParam
	structType := structValue.Type()

	//get fields name Fields in order to deserialize unknown fields
	fields := reflect.Value{}
	if structValue.Kind() == 25 {
		//search fields only if current item is a struct
		fields = structValue.FieldByName("Fields")
	}

	//for each key as jsonkey
	for key, value := range structuredData {
		//identify the field from the struct using json Tag naming

		var realField reflect.StructField
		for i := 0; i < structType.NumField(); i++ {
			if structType.Field(i).Tag.Get("json") == key {
				realField = structType.Field(i)
			}
		}
		//From here, the real name of the field is in realField.Name
		//fieldName will be the name specified in the json Tag
		fieldName := realField.Tag.Get("json")
		if fieldName != "" {
			switch realField.Type.Kind() {
			case 21:
				//realField is a Map, set the value using reflection
				if realField.Type.Elem().Kind() == 24 {
					log.Printf("%+v\n", realField.Type.String())
					log.Printf("%+v\n", value)
					str := fmt.Sprintf("%v", value)
					_ = str
					structValue.FieldByName(realField.Name).Set(reflect.ValueOf(value))
				} else {
					structValue.FieldByName(realField.Name).Set(reflect.ValueOf(value))
				}
				break
			case 1: //realField is a Constante (bool)
				if value != nil {
					structValue.FieldByName(realField.Name).Set(reflect.ValueOf(value))
				}
				break
			case 6: //realField is a Constante (int64)
				if value != nil {
					structValue.FieldByName(realField.Name).Set(reflect.ValueOf(value))
				}
				break
			case 24: //realField is a string, set the value using reflection
				if value != nil {
					structValue.FieldByName(realField.Name).Set(reflect.ValueOf(value))
				}
				break
			default:

				//realField is another declared struct
				//Find recursively from the realField Type all subfields
				deepConvert(value.(map[string]interface{}), structValue.FieldByName(realField.Name))
				break
			}
		} else if structValue.Kind() == 25 {
			//if the current structValue is a real struct
			if fields.IsValid() {
				//if field exist and is valid, it might be necessary to initialise it
				fields.Set(reflect.ValueOf(make(map[string]interface{})))
			}
			//add all unknown data to the "fields" field
			fields.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
		}
	}
}
