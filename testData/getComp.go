type Comps struct {
	MapthishomeLinksCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>links>mapthishome"`
	Region	Region	`xml:"response>properties>principal>localRealEstate>region"`
	ForSale	string	`xml:"response>properties>principal>localRealEstate>region>links>forSale"`
	StreetAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>street"`
	CityAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>city"`
	LowValuationRangeZestimateCompComparablesPropertiesResponse	[]Low	`xml:"response>properties>comparables>comp>zestimate>valuationRange>low"`
	Longitude	string	`xml:"response>properties>principal>address>longitude"`
	ForSaleLinksRegionLocalRealEstateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>localRealEstate>region>links>forSale"`
	Graphsanddata	string	`xml:"response>properties>principal>links>graphsanddata"`
	HomedetailsLinksCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>links>homedetails"`
	PercentileZestimateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>zestimate>percentile"`
	Count	string	`xml:"request>count"`
	Homedetails	string	`xml:"response>properties>principal>links>homedetails"`
	GraphsanddataLinksCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>links>graphsanddata"`
	StateAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>state"`
	Comp	[]Comp	`xml:"response>properties>comparables>comp"`
	ZipcodeAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>zipcode"`
	Code	string	`xml:"message>code"`
	ForSaleByOwnerLinksRegionLocalRealEstateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>localRealEstate>region>links>forSaleByOwner"`
	Street	string	`xml:"response>properties>principal>address>street"`
	Zipcode	string	`xml:"response>properties>principal>address>zipcode"`
	Zpid	string	`xml:"response>properties>principal>zpid"`
	State	string	`xml:"response>properties>principal>address>state"`
	ZindexValue	string	`xml:"response>properties>principal>localRealEstate>region>zindexValue"`
	AmountZestimateCompComparablesPropertiesResponse	[]Amount	`xml:"response>properties>comparables>comp>zestimate>amount"`
	SchemaLocation	string	`xml:"schemaLocation,attr"`
	Comparables	string	`xml:"response>properties>principal>links>comparables"`
	City	string	`xml:"response>properties>principal>address>city"`
	ZpidCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>zpid"`
	Latitude	string	`xml:"response>properties>principal>address>latitude"`
	ZindexValueRegionLocalRealEstateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>localRealEstate>region>zindexValue"`
	Last-Updated	string	`xml:"response>properties>principal>zestimate>last-updated"`
	OneWeekChange	OneWeekChange	`xml:"response>properties>principal>zestimate>oneWeekChange"`
	LatitudeAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>latitude"`
	LongitudeAddressCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>address>longitude"`
	OverviewLinksRegionLocalRealEstateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>localRealEstate>region>links>overview"`
	Comps	string	`xml:"Comps,attr"`
	Xsi	string	`xml:"xsi,attr"`
	ZpidRequest	string	`xml:"request>zpid"`
	Text	string	`xml:"message>text"`
	ValueChange	ValueChange	`xml:"response>properties>principal>zestimate>valueChange"`
	High	High	`xml:"response>properties>principal>zestimate>valuationRange>high"`
	ForSaleByOwner	string	`xml:"response>properties>principal>localRealEstate>region>links>forSaleByOwner"`
	ValueChangeZestimateCompComparablesPropertiesResponse	[]ValueChange	`xml:"response>properties>comparables>comp>zestimate>valueChange"`
	Percentile	string	`xml:"response>properties>principal>zestimate>percentile"`
	Last-UpdatedZestimateCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>zestimate>last-updated"`
	ComparablesLinksCompComparablesPropertiesResponse	[]string	`xml:"response>properties>comparables>comp>links>comparables"`
	HighValuationRangeZestimateCompComparablesPropertiesResponse	[]High	`xml:"response>properties>comparables>comp>zestimate>valuationRange>high"`
	Mapthishome	string	`xml:"response>properties>principal>links>mapthishome"`
	RegionLocalRealEstateCompComparablesPropertiesResponse	[]Region	`xml:"response>properties>comparables>comp>localRealEstate>region"`
	Low	Low	`xml:"response>properties>principal>zestimate>valuationRange>low"`
	Amount	Amount	`xml:"response>properties>principal>zestimate>amount"`
	Overview	string	`xml:"response>properties>principal>localRealEstate>region>links>overview"`
	OneWeekChangeZestimateCompComparablesPropertiesResponse	[]OneWeekChange	`xml:"response>properties>comparables>comp>zestimate>oneWeekChange"`
}

type Low struct {
	Text	string	`xml:",chardata"`
	Currency	string	`xml:"currency,attr"`
}
type High struct {
	Text	string	`xml:",chardata"`
	Currency	string	`xml:"currency,attr"`
}
type Amount struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type RegionLocalRealEstateCompComparablesPropertiesResponse struct {
	Name	string	`xml:"name,attr"`
	Id	string	`xml:"id,attr"`
	Type	string	`xml:"type,attr"`
}
type HighValuationRangeZestimateCompComparablesPropertiesResponse struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type AmountZestimateCompComparablesPropertiesResponse struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type OneWeekChangeZestimateCompComparablesPropertiesResponse struct {
	Deprecated	string	`xml:"deprecated,attr"`
}
type ValueChange struct {
	Text	string	`xml:",chardata"`
	Duration	string	`xml:"duration,attr"`
	Currency	string	`xml:"currency,attr"`
}
type Region struct {
	Type	string	`xml:"type,attr"`
	Name	string	`xml:"name,attr"`
	Id	string	`xml:"id,attr"`
}
type Comp struct {
	Score	string	`xml:"score,attr"`
}
type ValueChangeZestimateCompComparablesPropertiesResponse struct {
	Text	string	`xml:",chardata"`
	Duration	string	`xml:"duration,attr"`
	Currency	string	`xml:"currency,attr"`
}
type LowValuationRangeZestimateCompComparablesPropertiesResponse struct {
	Currency	string	`xml:"currency,attr"`
	Text	string	`xml:",chardata"`
}
type OneWeekChange struct {
	Deprecated	string	`xml:"deprecated,attr"`
}

