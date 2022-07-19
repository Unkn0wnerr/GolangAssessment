package main

import (
	// "Question6/apperrors"
	"errors"
	"fmt"
	"net/http"
)

type RequestError struct {
	Status  int
	Message string
}

func (e *RequestError) Error() string {
	return e.Message
}

func request(url string) (string, error) {
	resp, _ := http.Get(url)
	if 400 < resp.StatusCode && resp.StatusCode < 600 {
		return "", &RequestError{
			Status:  resp.StatusCode,
			Message: fmt.Sprintf("status %v: %v", resp.StatusCode, http.StatusText(resp.StatusCode)),
		}
	}
	return fmt.Sprintf("status %v: %v", resp.StatusCode, http.StatusText(resp.StatusCode)), nil
}

func main() {

	response, err := request("https://www.google.com/apple")

	if err != nil {
		var reqErr *RequestError
		switch {
		case errors.As(err, &reqErr):
			fmt.Println(reqErr.Error())
		default:
			fmt.Printf("unexpected request error: %s\n", err)
		}
	}

	fmt.Println(response)

}
