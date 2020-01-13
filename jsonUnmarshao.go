package main

import "fmt"

import "encoding/json"

type Bird struct {
	Id      int         `json:"id"`
	Version int         `json:"version"`
	Detail  interface{} `json:"detail"`
}

type BirdInfoV1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type BirdInfoV2 struct {
	Address string `json:"address"`
}

func (b *Bird) UnmarshalJSON(data []byte) error {
	type Alias Bird
	bird := &struct{ *Alias }{Alias: (*Alias)(b)}
	if err := json.Unmarshal(data, &bird); err != nil {
		return err
	}

	dataMap := bird.Detail.(map[string]interface{})
	if bird.Version == 1 {
		name := dataMap["name"].(string)
		age := dataMap["age"].(float64)
		b.Detail = BirdInfoV1{Name: name, Age: int(age)}
	} else if bird.Version == 2 {
		address := dataMap["address"].(string)
		b.Detail = BirdInfoV2{Address: address}
	}
	return nil
}

func main() {

	// egle을 json으로 변환
	egle := Bird{Id: 100, Version: 1, Detail: BirdInfoV1{Name: "Egle", Age: 10}}
	egleStr, err := json.Marshal(egle)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(egleStr))

	/// egleStr을 egle로 변환
	var egleV1 Bird
	json.Unmarshal([]byte(egleStr), &egleV1)
	fmt.Println(egleV1)

	// egleV2를 json으로 변환
	egleV2 := Bird{Id: 101, Version: 2, Detail: BirdInfoV2{Address: "seoul dong-gu"}}
	egleV2Str, err := json.Marshal(egleV2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(egleV2Str))

	// egleV2Str을 egle로 변환
	var egleV2Again Bird
	json.Unmarshal([]byte(egleV2Str), &egleV2Again)
	fmt.Println(egleV2Again)
}
