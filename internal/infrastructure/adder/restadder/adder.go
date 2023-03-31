package restadder

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Xsd     string   `xml:"xsd,attr"`
	Body    struct {
		Text        string `xml:",chardata"`
		AddResponse struct {
			Text      string `xml:",chardata"`
			Xmlns     string `xml:"xmlns,attr"`
			AddResult string `xml:"AddResult"`
		} `xml:"AddResponse"`
	} `xml:"Body"`
}

type Adder struct{}

func (a Adder) Do(ctx context.Context, x, y int) (int, error) {
	fmt.Println("REST adder")

	url := "http://www.dneonline.com/calculator.asmx"
	method := "POST"

	payload := strings.NewReader(
		fmt.Sprintf(
			`<?xml version="1.0" encoding="utf-8"?>
			  <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
			    <soap:Body>
				  <Add xmlns="http://tempuri.org/">
				    <intA>%d</intA>
				    <intB>%d</intB>
				  </Add>
			    </soap:Body>
			  </soap:Envelope>`,
			x, y,
		),
	)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return 0, fmt.Errorf("request error: %s", err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "http://tempuri.org/Add")

	res, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do error: %s", err)

	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("read error: %s", err)
	}

	env := &Envelope{}
	err = xml.Unmarshal(body, env)
	if err != nil {
		return 0, fmt.Errorf("unmarshall error: %s", err)
	}

	val := env.Body.AddResponse.AddResult
	atoi, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("atoi error: %s", err)

	}

	return atoi, nil
}
