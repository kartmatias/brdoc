package brdoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Unidade     string `json:"unidade"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
}

func ConsultaCep(cep string) (*Result, error) {
	var scep bytes.Buffer
	for _, v := range cep {
		switch v {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			scep.WriteRune(v)
		}
	}
	resp, err := http.Get("https://viacep.com.br/ws/" + scep.String() + "/json/unicode/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("http status code %v", resp.StatusCode)
	}
	res := &Result{}
	scep.Reset()
	_, err = io.Copy(&scep, resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	err = json.Unmarshal(scep.Bytes(), res)
	if err != nil {
		return nil, err
	}
	if res.Cep == "" {
		return nil, fmt.Errorf("not found")
	}

	return res, nil
}

func NotFound(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == "not found"
}
