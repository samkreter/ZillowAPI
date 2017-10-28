package main

import "encoding/xml"

type Searchresults struct {
	SearchResults 	xml.Name `xml:"SearchResults"`
    Request struct{
		Address			string	`xml:"address"`
		Citystatezip	string 	`xml:"citystatezip"`
	} 	`xml:"request"`
	Message struct{
		Text 	string 	`xml:"text"`
		Code 	int 	`xml:"code"`
	}	`xml:"message"`
	Response struct{
		Results 	[]Result	`xml:"results>result"`
	}	`xml:"response"`
}

type Result struct{
	Zpid 		string 		`xml:"zpid"`
	Links		Links		`xml:"links"` //TODO - FIGURE THIS OUT 
	Address		Address		`xml:"address"`
	Zestimate	Zestimate	`xml:"zestimate"`
}

type Links struct{
	Homedetails		string	`xml:"homedetails"`
	Graphsanddata	string	`xml:"graphsanddata"`
	Mapthishome		string	`xml:"mapthishome"`
	Comparables		string	`xml:"comparables"`
}

type Address struct{
	Street 		string	`xml:"street"`
	Zipcode		string	`xml:"zipcode"`
	City		string 	`xml:"city"`
	State		string 	`xml:"state"`
	Lat			string	`xml:"latitude"`
	Long		string	`xml:"longitude"`
}

type Zestimate struct{
	Amount	struct{
		Value			string			`xml:",chardata"`
		Currency		string			`xml:"currency,attr"`
	} `xml:"amount"`

	LastUpdated			string			`xml:"last-updated"` //TODO: change to date filed
	OneWeekChange		bool			`xml:"oneWeekChange"` 
	Deprecated			bool			`xml:"deprecated"`

	ValueChange struct{
		Value			int				`xml:",chardata"`
		Duration		int				`xml:"duration,attr"`
		Currency		string			`xml:"currency,attr"`
	} `xml:"valueChange"` 	

	ValuationRange		ValuationRange	`xml:"valuationRange"`
	Percentile			int				`xml:"percentile"`
}

type ValuationRange struct{
	Low struct{
		Low			int			`xml:",chardata"`
		Currency	string		`xml:"currency"`
	} `xml:"low"`
	High struct{
		High		int			`xml:",chardata"`
		Currency	string		`xml:"currency"`
	} `xml:"high"`
}
