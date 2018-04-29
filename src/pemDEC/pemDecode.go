package pemDEC

import (
	"encoding/pem"
	"io/ioutil"
	"fmt"
	"crypto/x509"
	"crypto/rand"
	"os"
)

func Jiami(file string, p string, createFile string) (error) {

	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ReadFile Error:", err.Error())
		return err
	}
	block, _ := pem.Decode(f)

	block, err = x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", block.Bytes, []byte(p), x509.PEMCipherAES256)
	if err != nil {
		fmt.Println("make block faild:", err.Error())
		return err
	}
	cf, err := os.Create(createFile)
	if err != nil {
		fmt.Println("create File faild:", err.Error())
		return err
	}
	defer cf.Close()
	pem.Encode(cf, block)
	return nil
}

func Jiemi(file string, p string, createFile string) error {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ReadFile Error:", err.Error())
		return err
	}
	block, _ := pem.Decode(f)

	byt, err := x509.DecryptPEMBlock(block,[]byte(p))
		if err != nil {
		fmt.Println("make block faild:", err.Error())
		return err
	}
	block=&pem.Block{
		Type:block.Type,
		Bytes:byt,
	}
	cf, err := os.Create(createFile)
	if err != nil {
		fmt.Println("create File faild:", err.Error())
		return err
	}
	defer cf.Close()
	pem.Encode(cf, block)
	return nil
}
