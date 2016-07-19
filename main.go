package main

import (
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("You need to supply an ARN and hex value!")
		return
	}

	arn := os.Args[1]
	plhex := os.Args[2]
	fmt.Printf("Working with... \n\tarn = %s\n\thex = %s", arn, plhex)
	if arn == "" {
		fmt.Println("ARN missing; can't encrypt.")
		return
	}

	b, err := hex.DecodeString(plhex)
	if err != nil {
		fmt.Printf("Err decoding hex; maybe your value is bad? \nErr was %s\n", err)
		return
	}

	k := kms.New(session.New())
	r, err := k.Encrypt(
		&kms.EncryptInput{
			KeyId:     &arn,
			Plaintext: b,
		},
	)

	if err != nil {
		fmt.Printf("err encrypting: \n%s\n", err)
		return
	}

	fmt.Printf("encrypted hex value is %s\n", hex.EncodeToString(r.CiphertextBlob))

}
