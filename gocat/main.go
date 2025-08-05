package main

import (
	"log"
	"os"
)
func main() {
	if len(os.Args) <= 1{
		log.Fatal("No arguments passed")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error trying to read file: %v",err.Error())
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil{
		log.Fatalf("Error trying to get file data size: %v",err.Error())
	}
	if fileInfo.IsDir(){
		log.Fatal("Path must be a file, got directory instead")
	}
	bytes := make([]byte, fileInfo.Size())
	_, err = file.Read(bytes)
	if err != nil {
		log.Fatalf("Error trying to get file size: %v",err.Error())
	}
	//var out *os.File = new(os.File) so this defines a pointer of os.File and new assigns it a location so it isnt nil
	var out *os.File
	if len(os.Args) >= 3 && os.Args[2] == ">"{
		out, err = os.Open(os.Args[3])
		if err != nil{
			log.Fatalf("Error trying to create file: %v",err.Error())
		}
	}else{
		//*out = *os.Stdout so this would be "The value of what out is pointing to = the value os.Stdout is pointing to
		//this comments do work but since theyre pointers...
		out = os.Stdout
		//simply changing the value out points to is much better
	}
	defer out.Close()
	if _, err := out.Write(bytes); err != nil{
		log.Fatalf("Error trying to write to output: %v",err.Error())
	}
}
