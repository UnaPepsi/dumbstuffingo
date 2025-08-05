package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func main() {
	green := "\033[32m"
	normal := "\033[0m"
	boldWhite := "\033[1;37m"
	boldGreen := "\033[1;32m"
	var colorToUse *string
	showHidden := slices.Contains(os.Args,"-a")
	showSize := slices.Contains(os.Args,"-s")
	var item string
	entries, err := os.ReadDir("./")
	if err != nil{
		log.Fatalf("Couldn't list all items: %v",err.Error())
	}
	for _,entry := range entries {
		isHidden := fmt.Sprintf("%c",entry.Name()[0]) == "."
		if isHidden && !showHidden{
			continue
		}
		if entry.IsDir(){
			if isHidden{
				colorToUse = &boldGreen
			}else{
				colorToUse = &green
			}
		}else{
			if isHidden{
				colorToUse = &boldWhite
			}else{
				colorToUse = &normal
			}
		}
		item = entry.Name()
		if showSize{
			info, err := entry.Info()
			if err != nil{
				log.Fatalf("Couldn't get item info: %v",err.Error())
			}
			var size float64
			if info.IsDir(){ 
				err = filepath.Walk("./"+info.Name(), func(path string, walkInfo os.FileInfo, err error) error {
					if err != nil{
						return err
					}
					size += float64(walkInfo.Size()) / 1e+6
					return nil
				})
				if err != nil{
					log.Fatalf("Couldn't get item info: %v",err.Error())
				}
			}else{
				size = float64(info.Size()) / 1e+6 //MB
			}
			item = fmt.Sprintf("%v - %.2fMB",item,size)
		}
		fmt.Println(*colorToUse+item)
	}
	fmt.Print(normal)
}
