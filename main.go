package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Message struct {
	Value   string
	Message string
}

func main() {

	encode := "wr zepmhexi ymj eqttgev cpuygt dtz mfaj wr jr wr zshzo gpkioc zshzo oguucig zshzo ywqhbelbwvlhkcgy fsi dtz hfs yjqq ymj wjxy ri ymj yjfr ymfy dtz bts ymj gleppirki"
	encodeArray := strings.Split(encode, " ")
	var message Message
	message = Message{Value: "", Message: " "}
	var data []byte
	var cont int = 0
	var flag bool = true
	for _, d := range encodeArray {
		flag = true
		cont = 0
		fmt.Println("Buscando -----------", d)
		for d != message.Value {
			cont++
			if flag {
				message.Value = d
				flag = false
			}
			response, err := http.Get("http://enigma.ml.com/enigma/encrypt/" + message.Value)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ = ioutil.ReadAll(response.Body)
			}
			err = json.Unmarshal(data, &message)
			fmt.Println(cont)
			fmt.Println(message.Value)
		}

	}

	/*response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}*/

}
