package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var ApiCotacao string = "http://localhost:8080/cotacao"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ApiCotacao, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		cotacao := string(body)
		SaveData(cotacao)
	} else {
		io.Copy(os.Stdout, res.Body)
	}

}

func SaveData(cotacao string) error {
	filename := "cotacao.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	_, err = file.Write([]byte(fmt.Sprintf("Dólar: %s", cotacao)))
	if err != nil {
		return err
	}

	file.Close()

	return nil
}
