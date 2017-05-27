package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func geturl() string {

	apiurl := "https://api.eveonline.com/"
	_, keyID, vCode, _, _, _ := config()

	xmlurl := apiurl + "char/Notifications.xml.aspx?keyID=" + keyID + "&vCode=" + vCode
	//fmt.Println(xmlurl)
	return xmlurl
}

func vaterpaz() {

	type Letters struct {
		LetterID string `xml:"notificationID,attr"`
		Type     string `xml:"typeID,attr"`
	}

	type Eveapi2 struct {
		XMLName     xml.Name  `xml:"eveapi"`
		CurrentTime string    `xml:"currentTime"`
		Row         []Letters `xml:"result>rowset>row"`
		CachedUntil string    `xml:"cachedUntil"`
	}

	Url := geturl()
	res, err := http.Get(Url)
	if err != nil {
		fmt.Println(err)
		return
	}

	eve, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	v := Eveapi2{}

	err = xml.Unmarshal([]byte(eve), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return

	}

	//fmt.Printf("Rows: %v\n", v.Row)
	for i := 0; i < len(v.Row); i++ {
		//fmt.Printf("Rows: %v\n", v.Row[i].Type)
		if v.Row[i].Type == "75" {
			fmt.Println("FOUND on ID: %v", v.Row[i].LetterID)
			//POS
			_, _, _, onenotif, _, _ := config()
			if onenotif { //Если в конфиге проверка true то
				if readfile(v.Row[i].LetterID) == false { //если в файле известных нет номера письма
					achtung("POS")               //алерт
					writefile(v.Row[i].LetterID) //добавляем номер в список известных

				}

			} else {
				achtung("POS")
			}

		} else if v.Row[i].Type == "184" {
			fmt.Println("FOUND on ID: %v", v.Row[i].LetterID)
			//цитадель
			_, _, _, onenotif, _, _ := config()
			if onenotif { //Если в конфиге проверка true то
				if readfile(v.Row[i].LetterID) == false { //если в файле известных нет номера письма
					achtung("Citadel")           //алерт
					writefile(v.Row[i].LetterID) //добавляем номер в список известных

				}

			} else {
				achtung("Citadel")
			}
		}

	}

}
