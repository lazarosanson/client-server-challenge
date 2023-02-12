package main

import (
	"context"
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
	io.Copy(os.Stdout, res.Body)
}
