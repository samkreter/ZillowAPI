package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"github.com/spf13/viper"
	"os"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)


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

func getXMLFile() (*Searchresults, error) {
	// Open our xmlFile
	xmlFile, err := os.Open("example.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var s = new(Searchresults)
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(byteValue, s)
	if err != nil {
		fmt.Println(err)
	}
	
	return s, err
}

func main(){
	setUpConfig()
	s, err := getXMLFile()
	if err != nil{
		panic(err)
	}

	os.Remove("test.db")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
	  panic("failed to connect database")
	}
	defer db.Close()
  
	// Migrate the schema
	db.AutoMigrate(&Result{})

	db.Create(s.Response.Results[0])

	//fmt.Println(viper.GetString("zwsId"))
}

