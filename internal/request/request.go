package request

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (Request, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return Request{}, fmt.Errorf("could not read from io reader")
	}
	parts := strings.Split(string(b), "\r\n")
	requestLine, err := parseRequestLine(parts[0])
	if err != nil {
		return Request{}, fmt.Errorf("could not create request line")
	}
	return Request{RequestLine: requestLine}, nil
}

func parseRequestLine(request string) (RequestLine, error) {
	parts := strings.Split(request, " ")
	if len(parts) != 3 {
		return RequestLine{}, fmt.Errorf("want 3 request parts, have %d", len(parts))
	}
	if !isMethodValid(parts[0]) {
		return RequestLine{}, fmt.Errorf("method signature invalid")
	}
	httpVersion, err := isHTTPVersionValid(parts[2])
	if err != nil {
		return RequestLine{}, fmt.Errorf("http version invalid")
	}
	return RequestLine{HttpVersion: httpVersion, RequestTarget: parts[1], Method: parts[0]}, nil
}

func isMethodValid(method string) bool {
	for _, char := range method {
		if !unicode.IsLetter(char) {
			return false
		}
		if !unicode.IsUpper(char) {
			return false
		}
	}
	return true
}

func isHTTPVersionValid(version string) (string, error) {
	parts := strings.Split(version, "/")
	if len(parts) != 2 {
		return "", fmt.Errorf("wrong length")
	}
	if parts[1] != "1.1" {
		return "", fmt.Errorf("wrong length")
	}
	return parts[1], nil
}
