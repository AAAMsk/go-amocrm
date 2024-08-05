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

	accessToken  = `eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjBlMjNhYjQ3NTllMjEzNzc2YTI0NDlmMzlhNzEzZDcxNzk0YzgwYzNhYTE3OTZkYTVlNWVjMWUzYmZjYjdjN2IyNjA5ODU2OThkODY1Y2Q3In0.eyJhdWQiOiI4NmJmZTU1Ny05YTUyLTQ1YzItODYxMy00NTY5NDNkYzNjY2EiLCJqdGkiOiIwZTIzYWI0NzU5ZTIxMzc3NmEyNDQ5ZjM5YTcxM2Q3MTc5NGM4MGMzYWExNzk2ZGE1ZTVlYzFlM2JmY2I3YzdiMjYwOTg1Njk4ZDg2NWNkNyIsImlhdCI6MTcxODg3NzY0NSwibmJmIjoxNzE4ODc3NjQ1LCJleHAiOjE3MTg5NjQwNDUsInN1YiI6IjExMTI5MDQ2IiwiZ3JhbnRfdHlwZSI6IiIsImFjY291bnRfaWQiOjMxNDAxOTAyLCJiYXNlX2RvbWFpbiI6ImFtb2NybS5ydSIsInZlcnNpb24iOjEsInNjb3BlcyI6WyJjcm0iXSwiaGFzaF91dWlkIjoiNGI1NmFlOWYtMjY4YS00OTVlLWIwYzktY2MyZjQ1MDA1NTQxIn0.O6VYuL0J2aobYKEhZOEBjEu7WlISxkGA237AWECiWT7yv_iEyjuvXAeqG8tHSXkZu3SIpQ3XqLDg_ba4YnE3Ffzv3IyUX5zisn1QuVDy1aCagGIchsNCGIcT7jkhIpSClrt2uskIuj9GAkUsPTKpZSj_u-DkFCXvtgBeSp87R0nN7sH4585hep0nJ3PoZy1haomtldwoMaqnHg_PudeunD5P4wN8CVP2f93S6RcNOxLi1EwszJw4LHxp34ZRso9mTxB7Dn6LSLlF8p_Ll33_gDrFM7bySPhrDzT23rCtlccwTQuqIf-MHelD_Bv3MoLc95dmzhVkIQV9DWJ9i1WFoA`
	refreshToken = `def5020022e38ea4d5b4e1304e99695f98ca22d97ae30eb2449225ac3b0e002d37198891666357275e9c883a99b2b33834c5ffdc4b7dd4f7b6f530151834546de54d90712a51df95222907cf9ddb1317856c41452b33c5c708f7055dd09739958cdd719c6094ea77eaf9d5364bb3ec1e2a7cecfd470c71ca93a3229e0070e0c9aab0ade9ee9095cda36e9bbebb7c667f04d0068d583faa16d19561d29e7d517e19847380dac78f29e01a6cc15f2f9d1be2d7ea1ccb37029ac57127eb107daa2f3d87588be40cd30cf338b6199a08399af2a3a83f5b5316473c3d798bdcd273f5bc0b3185ed0b1532c1edf3eb05571617e6bd33c7f54bd831242454b2897c1d37d4b2e5b4ca9b93267020423003b78c797a0e5df578477f27ca6c37d0bc761d44c861cfcbcaea99eacc876c97acd7359405352477a9a8d065e0bc063cdf167e8d0bf6ff454697a971c479748955761f8abe0ab9af71a68e2a111bc40d403287cd38335fae30f47ac2cc93985baa1c23572e0655aded722cbae9ff20e8ecc69a46b64e764409d5dcdd628a6546f2d495c90cde55ef274658c8cd2e64b87d4024b2a16ccd042351553c63a2920b23ff27e935c5`
)

func main() {
	amoCRM := amocrm.NewAPI(clientID, clientSecret, redirectURI)
	amoCRM.SetOptions(domain, accessToken, true)

	p := amocrm.Params{
		Limit: `10`,
	}

	leads, _ := amoCRM.Get().Leads("", &p)

	f, _ := os.Create("leads.json")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	asJson, _ := json.MarshalIndent(leads, "", "\t")
	f.Write(asJson)
}
