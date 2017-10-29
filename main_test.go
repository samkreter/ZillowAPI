package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestXMLParsing(t *testing.T) {
	
	s,err := getXMLFile()
	if err != nil{
		t.Fatal("ERROR: ", err.Error())
	}

	assert.Equal(t,"Seattle", s.Response.Results[0].Address.City)
	assert.Equal(t,"http://www.zillow.com/east-queen-anne-seattle-wa/",strings.TrimSpace(s.Response.Results[0].LocalRealEstate[0].Links.ForSale))
	assert.Equal(t, "USD", s.Response.Results[0].Zestimate.Amount.Currency)
	assert.Equal(t,1024380, s.Response.Results[0].Zestimate.ValuationRange.Low.Low)


}