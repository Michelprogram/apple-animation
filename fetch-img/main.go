package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	var err error

	err = cleanFolder()

	if err != nil{
		fmt.Printf("Problem during Cleaning folder : %v\n", err)
	}

	err = fetchImgMac()

	if err != nil{
		fmt.Printf("Problem during fetch : %v\n", err)
	}

	fmt.Println("All images collected")
}

func cleanFolder() error{

	var err error	

	err = os.RemoveAll("images")

	if err != nil{
		return err
	}

	err = os.Mkdir("images", 0777)
	
	if err != nil{
		return err
	}

	return nil
}

func fetchImgMac() error{

	var tempo string

	numbersOfImages := 87

	for i := 0; i < numbersOfImages; i++{


		if i / 10 == 0{
			tempo = fmt.Sprintf("0%d",i)
		} else {
			tempo = fmt.Sprintf("%d", i)
		}

		url := fmt.Sprintf("https://www.apple.com/105/media/us/macbook-pro-14-and-16/2021/a1c5d17e-d8e4-4fa8-b70a-bc61bd266412/anim/hero-specs//large/large_00%s.jpg",tempo)
		name := fmt.Sprintf("images/img_%s.jpg",tempo)

		fmt.Printf("Fetching img : %s\n", url[122:])

		//Get the response bytes from the url
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return errors.New("Received non 200 response code")
		}
		//Create a empty file
		file, err := os.Create(name)
		if err != nil {
			return err
		}
		defer file.Close()

		//Write the bytes to the fiel
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

	}

	return nil
	
}
