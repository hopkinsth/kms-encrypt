package main

import (
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to supply an ARN and hex value!")
		return
	}

	arn := os.Args[0]
	plhex := os.Args[1]

	if arn == "" {
		fmt.Println("ARN missing; can't encrypt.")
		return
	}

	b, err := hex.DecodeString(plhex)
	if err != nil {
		fmt.Printf("Err decoding hex; maybe your value is bad? \n Err was %s\n", err)
	}

	k := kms.New(session.New())
	r, err := k.Encrypt(
		&kms.EncryptInput{
			Plaintext: b,
		},
	)

	if err != nil {
		fmt.Printf("err encrypting: \n%s\n", err)
	}

	fmt.Printf("encrypted hex value is %s\n", hex.EncodeToString(r.CiphertextBlob))

}
