package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"github.com/spf13/viper"
	"os"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// "bytes"
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

func getStations(body []byte) (*Searchresults, error){
	var s = new(Searchresults)
	err := xml.Unmarshal(body, &s)
	if (err != nil){
		fmt.Println("Bad broski: ",err)
	}
	return s, err
}

func getZillowSearchResult(address string, citystatezip string){
	u, err := url.Parse("http://www.zillow.com/webservice/GetSearchResults.html")
	checkErr(err)

	zwsId := viper.GetString("zwsId")

	q := u.Query()

	q.Add("zws-id",zwsId)
	q.Add("address", address)
	q.Add("citystatezip",citystatezip)
	u.RawQuery = q.Encode()
	fmt.Println(u)
}


func getXMLFile() (*Searchresults, error) {
	// Open our xmlFile
	xmlFile, err := os.Open("testData/example.xml")
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

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

func main(){
	setUpConfig()
	testing("sam","stillls")
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

