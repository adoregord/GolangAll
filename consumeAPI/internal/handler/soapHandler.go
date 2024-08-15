package handler

import (
	"bytes"
	"consume-api-try/internal/domain"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SoapHandler struct{}

type SoapHandlerInterface interface {
	ConsumeSOAP(c *gin.Context)
}

func NewSoapHandler() SoapHandlerInterface {
	return SoapHandler{}
}

func (h SoapHandler) ConsumeSOAP(c *gin.Context) {
	// json request for request body
	var request domain.MultiplyReq

	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	soapRequest := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <Multiply xmlns="http://tempuri.org/">
      <intA>%d</intA>
      <intB>%d</intB>
    </Multiply>
  </soap:Body>
</soap:Envelope>`, request.Angka1, request.Angka2)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var envelope domain.Envelope
	err = xml.Unmarshal(body, &envelope)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Multiply Result": envelope.Body.MultiplyResponse.MultiplyResult})
}
