package main

import (
	"fmt"
	"os"

	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

func main() {

	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		fmt.Println("Utilización")
		fmt.Println(argsWithProg[0], "<JWT token en Base64>")
		fmt.Println("")
		os.Exit(1)
	}

	tokenToTest := argsWithProg[1]

	toValidate := map[string]string{}
	toValidate["aud"] = "api://default"
	toValidate["cid"] = "0oaclb5oh71DDfoSh4x6"

	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           "https://dev-175355.okta.com/oauth2/default",
		ClaimsToValidate: toValidate,
	}

	verifier := jwtVerifierSetup.New()
	verifier.SetLeeway(60) // seconds

	token, err := verifier.VerifyAccessToken(tokenToTest)
	if err != nil {
		fmt.Println("KO --> Error en verifier.VerifyAccessToken:", err)
	} else {
		fmt.Println("OK --> El token ha sido validado.")
		for k, v := range token.Claims {
			fmt.Printf("%s -> %v\n ", k, v)
		}
	}
}
