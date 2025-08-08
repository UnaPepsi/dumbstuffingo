package main

import (
	"io"
	"log"
	"os"
	"strings"
)
func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("Error trying to get stdin info: %v",err.Error())
	}
	if info.Mode() & os.ModeNamedPipe != os.ModeNamedPipe {
		log.Fatal("Only pipe stdin supported")
	}
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error trying to read stdin: %v",err.Error())
	}
	data := string(bytes)
	var dataToFilter string
	if len(os.Args) > 1{
		dataToFilter = os.Args[1]
	}
	lines := strings.Split(data, "\n")
	builder := strings.Builder{}
	for _,line := range lines {
		if strings.Contains(line, dataToFilter){
			if _,err := builder.WriteString(line+"\n"); err != nil{
				log.Fatalf("Error ocurred creating string: %v",err.Error())
			}
		}
	}
	if _,err := os.Stdout.WriteString(builder.String()); err != nil{
		log.Fatalf("Error trying to print to stdout: %v",err.Error())
	}
}
