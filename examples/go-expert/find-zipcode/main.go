package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// http://viacep.com.br/ws/36025-275/json/

type ZipCode struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	State        string `json:"estado"`
}

func main() {
	for _, zipCode := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + zipCode + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "request error: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read response error: %v", err)
		}

		var data ZipCode
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "json parse error: %v\n", err)
		}
		fmt.Println(data)

		file, err := os.Create("zip-code.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "create file error: %v\n", err)
		}
		defer file.Close()
		err = json.NewEncoder(file).Encode(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "json encode error: %v\n", err)
		}
	}
}
