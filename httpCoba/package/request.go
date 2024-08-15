package hatetepe

import (
	"bufio"
	"encoding/json"
	"errors"

	// "errors"
	"fmt"
	"io"
	"net"
	"strings"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    io.Reader
}

func calculateContentLength(body io.Reader) (string, error) {
	if body == nil {
		return "0", nil
	}

	var buf [64]byte
	n, err := body.Read(buf[:])
	if err != nil && err != io.EOF {
		return "", err
	}

	return fmt.Sprintf("%d", n), nil
}

// TODO: Need handle error if the body json is nil
func (r *Request) ParseJSON(v interface{}) error {
	if r.Body == nil{
		return fmt.Errorf("request body is nil")
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func parseRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)

	// Read the request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	parts := strings.Fields(requestLine)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line")
	}

	req := &Request{
		Method:  parts[0],
		URL:     parts[1],
		Headers: make(map[string]string),
	}

	// Read headers
	for {
		line, err := reader.ReadString('\n')

		if err != nil || line == "\r\n" {
			break
		}

		parts := strings.SplitN(strings.TrimSpace(line), ":", 2)

		if len(parts) == 2 {
			req.Headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
		
	}

	length, err := calculateContentLength(reader)
	if err != nil{
		return req, err
	}
	if length == "0" {
		return req, errors.New("wokwok")
	}

	// if length, ok := req.Headers["Content-Length"]; ok && length == "0" {
	// 	return nil, fmt.Errorf("request body is empty")
	// }

	req.Body = reader
	return req, nil
}