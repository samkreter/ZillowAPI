package main

import (
	"fmt"
	// "net/http"
	"encoding/xml"
	"io/ioutil"
	"github.com/spf13/viper"
	"os"
)


type ZillowResponse struct {
    ExecutionTime xml.Name `xml:"executionTime"`
}


func setUpConfig(){
	viper.SetConfigName(".config.dev") 
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func getStations(body []byte) (*Searchresults, error){
	var s = new(Searchresults)
	err := xml.Unmarshal(body, &s)
	if (err != nil){
		fmt.Println("Bad broski: ",err)
	}
	return s, err
}

// func getSearchResult(){
// 	res, err := http.Get("https://www.citibikenyc.com/stations/json")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	s, err := getStations([]byte(body))
// }

func getXMLFile(){
	// Open our xmlFile
	xmlFile, err := os.Open("e1.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	fmt.Println(string(byteValue))

	var s Searchresults
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(byteValue, &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("test: ",s.Request.Address)

}

func main(){
	setUpConfig()
	getXMLFile()

	fmt.Println(viper.GetString("zwsId"))
}


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
		Results []Result
	}	`xml:"results"`
}


type Result struct{
	Zpid 		string 		`xml:"zpid"`
	Links		Links		`xml:"links"` //TODO - FIGURE THIS OUT 
	Address		Address		`xml:"address"`
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
	Amount			string		`xml:"amount"`
	Currency		string		`xml:"currency,attr"`
	LastUpdated		string		`xml:"last-updated"` //TODO: change to date filed
	OneWeekChange	bool		`xml:"oneWeekChange"` 
	Deprecated		bool		`xml:"deprecated"`
	ValueChange		int			`xml:"valueChange"` 	
	Duration		int			`xml:"duration,attr"`
	Currency		string		`xml:"currency,attr"`
	valuationRange
	percentile

}

// the user struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

// a simple struct which contains all our
// social links
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}
