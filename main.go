package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"github.com/spf13/viper"
	"os"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/url"
)

func setUpConfig(){
	viper.SetConfigName(".config.dev") 
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { 
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func getZillowSearchResult(address string, citystatezip string){
	u, err := url.Parse("http://www.zillow.com/webservice/GetSearchResults.html")
	checkErr(err, "URL PARSE ERROR")

	zwsId := viper.GetString("zwsId")

	q := u.Query()

	q.Add("zws-id",zwsId)
	q.Add("address", address)
	q.Add("citystatezip",citystatezip)

	u.RawQuery = q.Encode()
	fmt.Println(u)
}

func getXMLFile() (*Searchresults, error) {
	xmlFile, err := os.Open("testData/example.xml")
	checkErr(err,"File Operation Error:")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var s = new(Searchresults)
	err = xml.Unmarshal(byteValue, s)
	checkErr(err,"XML ERROR")
	
	return s, err
}

func checkErr(err error, msg string){
	if err != nil{
		panic(fmt.Errorf("%s %s \n", msg, err))
	}
}

func main(){
	setUpConfig()
	// s, err := getXMLFile()
	// checkErr(err)

	// os.Remove("test.db")
	// db, err := gorm.Open("sqlite3", "test.db")
	// checkErr(err)
	// defer db.Close()
  
	// // Migrate the schema
	// db.AutoMigrate(&Result{})

	// db.Create(s.Response.Results[0])

}

