package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"prints"

	"github.com/fatih/color"
)

type TokenRequestData struct {
	GENERATE_MACHINE_ID      string `json:"generate_machine_id"`
	FORMAT                   string `json:"format"`
	LOCALE                   string `json:"locale"`
	RETURN_SSL_RESOURCES     string `json:"return_ssl_resources"`
	CREDENTIALS_TYPE         string `json:"credentials_type"`
	PASSWORD                 string `json:"password"`
	GENERATE_SESSION_COOKIES string `json:"generate_session_cookies"`
	METHOD                   string `json:"method"`
	SIG                      string `json:"sig"`
	V                        string `json:"v"`
	APIKEY                   string `json:"apikey"`
	EMAIL                    string `json:"email"`
}

type TokenResponseData struct {
	ERROR_CODE int    `json:"error_code"`
	ERROR_MSG  string `json:"error_msg"`
	ERROR_DATA string `json:"error_data"`
}

func checkToken() (empty bool, token string) {
	prints.PrintPlus()
	color.New(color.FgYellow, color.Bold).Println("Starting GoFacebookDataExtractor...")
	fmt.Println()
	prints.PrintPlus()
	color.New(color.FgGreen, color.Bold).Println("Making sure token file exists...")
	if _, err := os.Stat("token/creds.token"); err == nil {
		//File found print status
		prints.PrintExclamation()
		color.New(color.FgRed, color.Bold).Println("Token file found...")
		prints.PrintPlus()
		color.New(color.FgGreen, color.Bold).Println("Checking if the file is empty...")
		//Check if file is empty
		fileContent, err := ioutil.ReadFile("token/creds.token")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if len(fileContent) > 0 {
			prints.PrintExclamation()
			color.New(color.FgRed, color.Bold).Println("Token file is not empty...")
			prints.PrintPlus()
			color.New(color.FgGreen, color.Bold).Printf("Token Found : ")
			color.New(color.FgGreen, color.Bold).Printf("%v\n", string(fileContent[:]))
			token = string(fileContent[:])
			return false, token
		} else {
			prints.PrintExclamation()
			color.New(color.FgRed, color.Bold).Println("Token file empty...")
			return true, ""
		}

	} else {
		//File not found print status
		prints.PrintExclamation()
		color.New(color.FgRed, color.Bold).Println("Token file not found...")
		prints.PrintPlus()
		color.New(color.FgGreen, color.Bold).Println("Creating it...")
		//Create file
		_, err := os.Create("token/creds.token")
		//If Create fails print log and terminate process
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		//File Successfully Created print status
		prints.PrintExclamation()
		color.New(color.FgRed, color.Bold).Println("File : token/creds.token successfully created !")
		return true, ""
	}
}

func AskForToken() bool {
	return true
}

func GetToken() (Token string) {
	var username string                              //Username Input
	var password string                              //Password Input
	API_SECRET := "62f8ce9f74b12f84c123cc23437a4a32" //Set the API Key, We will steal xHak5x one hehehe, just temporarily guys :)
	var JsonResponseData TokenResponseData
	//Ask for username
	fmt.Println()
	prints.PrintExclamation()
	color.New(color.FgRed, color.Bold).Println("Make sure to disable any VPN/Proxies !")
	prints.PrintExclamation()
	color.New(color.FgRed, color.Bold).Println("You can use a VPN/Proxy in the same country as the account you're using")
	prints.PrintPlus()
	color.New(color.FgGreen, color.Bold).Println("In order to get your account token you need to provide your login credentials")
	prints.PrintPlus()
	color.New(color.FgGreen, color.Bold).Print("Username : ")
	fmt.Scanf("%s", &username)

	//Ask for password
	prints.PrintPlus()
	color.New(color.FgGreen, color.Bold).Print("Password : ")
	fmt.Print("\033[8m")
	fmt.Scanf("%s", &password)
	fmt.Print("\033[28m")
	TokenData := TokenRequestData{
		"1",                                //Generate Machine ID
		"JSON",                             //Format
		"en_US",                            //Locale
		"0",                                //Return SSL Resources
		"password",                         //Credentials Type
		password,                           //Password
		"1",                                //Generate Session Cookies
		"auth.login",                       //Method
		"",                                 //Sig
		"1.0",                              //V
		"882a8490361da98702bf97a021ddc14d", //API Key
		username,                           //Email
	}

	//Set our Sig value in Token Data
	TokenData.SIG = "api_key=" + TokenData.APIKEY + "credentials_type=passwordemail=" + TokenData.EMAIL + "format=JSONgenerate_machine_id=1generate_session_cookies=1locale=en_USmethod=auth.loginpassword=" + TokenData.PASSWORD + "return_ssl_resources=0v=1.0" + API_SECRET
	hash := md5.Sum([]byte(TokenData.SIG))
	TokenData.SIG = hex.EncodeToString(hash[:])
	//Set other data in the request
	/*JsonTokenData, err := json.Marshal(TokenData)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}*/

	//Make request to the facebook api (https://api.facebook.com/restserver.php)
	resp, err := http.Get("https://api.facebook.com/restserver.php?" + "generate_machine_id=" + TokenData.GENERATE_MACHINE_ID + "&format=" + TokenData.FORMAT + "&locale=" + TokenData.LOCALE + "&return_ssl_resources=" + TokenData.RETURN_SSL_RESOURCES + "&credentials_type=" + TokenData.CREDENTIALS_TYPE + "&password=" + TokenData.PASSWORD + "&generate_session_cookies=" + TokenData.GENERATE_SESSION_COOKIES + "&method=" + TokenData.METHOD + "&sig=" + TokenData.SIG + "&v=" + TokenData.V + "&api_key=" + TokenData.APIKEY + "&email=" + TokenData.EMAIL)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	//Decode our json response
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&JsonResponseData)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//If the response code is different than 200 print error
	if JsonResponseData.ERROR_CODE == 368 {
		prints.PrintExclamation()
		color.New(color.FgRed, color.Bold).Println("It look's like the API blocked your IP Address")
		prints.PrintExclamation()
		color.New(color.FgRed, color.Bold).Println("Try again later or use another IP !")
	}

	if JsonResponseData.ERROR_CODE == 400 {
		prints.PrintExclamation()
		color.New(color.FgRed, color.Bold).Println("It look's like the credentials used were incorrect")
		prints.PrintExclamation()
		color.New(color.FgRed, color.Bold).Println("Make sure the credentials used were valid !")
	}

	//fmt.Println()                         //Print for debugging (TO REMOVE)
	//fmt.Println(string(JsonTokenData[:])) //Print for debugging (TO REMOVE)

	return ""
}

func main() {

	prints.PrintBanner()
	//Check if the file token/creds.token exists and that the file is not empty.
	isTokenEmpty, Token := checkToken()
	if isTokenEmpty {
		//If the token is empty ask for the user login and get token.
		Token := GetToken()
		fmt.Println(Token) //Print the token for debugging (TO REMOVE)
	} else {
		//If the token is not empty ask the user if he wants to use the current token. if not get a new token.
		var keepToken string
		//for with no instruction is a While loop in Go :)
		for {
			//From our module we made in src/prints we use the function PrintInterrogation
			fmt.Println()
			prints.PrintInterrogation()
			color.New(color.FgYellow, color.Bold).Print("Are you sure you want to use the following token for the data extraction ?\n")
			prints.PrintToken()
			color.New(color.FgYellow, color.Bold).Printf("%v ", Token)
			color.New(color.FgGreen, color.Bold).Print("[Y/n] : ")
			fmt.Scanf("%s", &keepToken)
			if keepToken == "y" || keepToken == "Y" || keepToken == "n" || keepToken == "N" {
				break
			} else {
				prints.PrintExclamation()
				color.New(color.FgRed, color.Bold).Print("Please enter a valid input (Y/n) !")
			}
		}
		if keepToken == "n" || keepToken == "N" {
			Token := GetToken()
			fmt.Println(Token)
		}
	}
}
