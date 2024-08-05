package main

import (
	"encoding/json"
	"github.com/AAAMsk/go-amocrm"
	"log"
	"os"
)

const (
	clientID     = "86bfe557-9a52-45c2-8613-456943dc3cca"
	clientSecret = "K9fKhbkBWcWR3bNUHoYFYfo4UKCxEbf5vJmAM3quG57uZD2DKv9bpD70qvLl8Gr6"
	redirectURI  = "https://amocrm-interlayer.aanda.ru/oauth/amocrm/instance/redirect"
	domain       = "zabroniryi.amocrm.ru"

	accessToken  = `eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjNmMTMxZWJjOTlmY2IxNjU1MDFmNTUzNTVhYjc1ZWU1MTE2MTczZTFkMDY2M2ZkNjQ1ZDEwY2RiYjRjZTkzZjAwZjY4NTdkNmU0Zjg5ZDhmIn0.eyJhdWQiOiI4NmJmZTU1Ny05YTUyLTQ1YzItODYxMy00NTY5NDNkYzNjY2EiLCJqdGkiOiIzZjEzMWViYzk5ZmNiMTY1NTAxZjU1MzU1YWI3NWVlNTExNjE3M2UxZDA2NjNmZDY0NWQxMGNkYmI0Y2U5M2YwMGY2ODU3ZDZlNGY4OWQ4ZiIsImlhdCI6MTcyMjg2Mzc2NiwibmJmIjoxNzIyODYzNzY2LCJleHAiOjE3MjI5NTAxNjYsInN1YiI6IjExMTI5MDQ2IiwiZ3JhbnRfdHlwZSI6IiIsImFjY291bnRfaWQiOjMxNDAxOTAyLCJiYXNlX2RvbWFpbiI6ImFtb2NybS5ydSIsInZlcnNpb24iOjEsInNjb3BlcyI6WyJjcm0iXSwiaGFzaF91dWlkIjoiMThhNWE3ZDItZjg4OS00OTc5LWIyNmYtNDhhODJkZjkwYjdlIn0.OysViCn--t6FiBiOEOIqBnsQVuIffx72EoZLwnm4JSwPoVSi8uCYMz6nnEvagoGi2sL1D7D2yw4lLfcT0wW5hCloFxRG0GLMcrTedP2EdxCypgo1toYrcw-Adv_1LLYH-mawSP6S9nN7Aaz77ywB5fzq2qQufiV7r4TyFOCwnQzVSNGzRrAf1hsvkfp_uSZKsAKMiOir5h727owkbV8kyZAGRXl6PzP9YmPBrL6ivFxDOcr8F5hs478K3VSl0qOI9Gw-HJpvaITBBJvo0m4rsL0qP5GDH6i3Ao_4FLwDsAzvn5NxlDIkWfil7Eyy0VMMj7gVjADadq6LzQE7aeL_9Q`
	refreshToken = `def50200afe2730460e1b5584d867ead8dd8ba19c7cfc93f1bafa2c34becbf370db89c7bf68d3d31ec9900fc805e1a4acbd7620209a424d8345a6e14b39828e87745efdee04fde9db315a6f91afdc71db1144e1e8b666e15febbb0777db758043d35f2b245ca629176335b21f150f39cfe1a72fa43e6785d1683deafec62f2395f5512940d8f24868768ee83f426d1ab946031cceaa3995476b16225c434ace5685f5d58da4d0087ffbda2c53863260b83009fdbc7a59dc66999da4edeaae387f683f2a399fbf7ad9030d147c9cf60a41a77cde3abe7b1920246a0a7a7026255048b8d3b774f7472f3495e476cd572d92c634a01b6a44857c2425b8e9116d471c4d9ec04bb510ea9012f44cde10d0d6476a88505540d5251684921f93f4c448d35c8bb99127b3de5f77c6ce8d34c1254ec556297f1a0e19f01891a10258b432bf4ba0d453ab302206ab5200657b0782f40ad1026330cc72854fdf24a0fcf639ec66d3b686217664cac8482172e0a337ea81098336f54f411522844a581a6559f55372998afd9c94e4b5cdb56222fc4c59c29efadd0bfb0ca18317e4b7348b8d39b6547fe22510d92550bc5b745f697498f91`
)

func main() {
	amoCRM := amocrm.NewAPI(clientID, clientSecret, redirectURI)
	amoCRM.SetOptions(domain, accessToken, true)

	p := amocrm.Params{
		Limit: "250",
	}

	_, contact, _ := amoCRM.Get().GetContactsByCustomFields("74951059440", "zopa@mail.ru", &p)

	//_, contact, _ := amoCRM.Get().GetCompaniesByCustomFields("78512547404", "zopa@mail.ru", &p)

	f, _ := os.Create("company.json")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	asJson, _ := json.MarshalIndent(contact, "", "\t")
	f.Write(asJson)
}
