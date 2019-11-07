package main

import (
	"fmt"
	"io/ioutil"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	object := Merchant{
		MerchantId: "merch1",
		Name:       "1234",
		Url:        "https://www.restaurantpublicwebsite.com",
	}

	fmt.Println("The previous Name is:", object.GetName())
	object.Name = "I renamed you"
	fmt.Println("The current Name is:", object.GetName())

	writeToFile("feeds.bin", &object)

	readObject := &Merchant{}
	readFromFile("feeds.bin", readObject)
	fmt.Println("Read the content:", readObject)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
		return err2
	}

	return nil
}
