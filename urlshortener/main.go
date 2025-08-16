package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
	_ "modernc.org/sqlite"
)

func main() {
	//secret,err := totp.Generate(totp.GenerateOpts{
	//	Issuer: randomString(10),
	//	AccountName: randomString(10),
	//})
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return
	//}
	secret, exists := os.LookupEnv("SECRET")
	if !exists {
		panic("SECRET env not set")
	}
	db,err := sql.Open("sqlite","urls.db")
	if err != nil{
		panic(err)
	}
	createTable(db) //not using &db cuz sql.Open alr returns a pointer
	mux := http.NewServeMux()
//	mux.HandleFunc("POST /", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Printf("request.Method: %v\n", request.Method)
//		fmt.Fprintf(writer, "%v", request.Method)
//	})

//	mux.HandleFunc("GET /digits", func(writer http.ResponseWriter, request *http.Request) {
//		digits, err := totp.GenerateCode(secret.Secret(),time.Now())	
//		if err != nil {
//			fmt.Printf("err: %v\n", err)
//			writer.WriteHeader(http.StatusInternalServerError)
//			fmt.Fprint(writer,"Something wrong happened jajajajaj")
//			return
//		}
//		fmt.Fprintf(writer,"%v",digits)
//	})
	mux.HandleFunc("POST /shorten", func(writer http.ResponseWriter, request *http.Request) {
		//request.URL.Query().Get("param")
		digitsReceived := request.Header.Get("totp")
		digits, err := totp.GenerateCode(secret, time.Now())
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(writer,"An error occurred, check logs")
			fmt.Printf("err: %v\n", err)
			return
		}
		if digits != digitsReceived {
			writer.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(writer,"fuk u")
			return
		}
		urlReceived := request.Header.Get("url")
		if urlReceived == "" {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(writer,"bro tf u doing pass the url")
			return
		}
		hash := sha256.New()
		_,err = hash.Write([]byte(urlReceived))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(writer,"An error occurred, check logs")
			fmt.Printf("err: %v\n", err)
			return
		}
		hashBytes := hash.Sum(nil)
		fmt.Printf("hashBytes: %x\n", hashBytes)
		fmt.Printf("string(hashBytes): %x\n", string(hashBytes))
		value := fmt.Sprintf("%x",hashBytes)[:8]
		query := "INSERT INTO urls VALUES (?, ?)"
		_,err = db.Exec(query, value, urlReceived)
		if err != nil {
			//writer.WriteHeader(http.StatusInternalServerError)
			//fmt.Fprintln(writer,"no se pudo jaja te jodiste")
			//fmt.Printf("%v",err)
			fmt.Printf("err.Error(): %v\n", err.Error())
		}
		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer,value)
	})

	mux.HandleFunc("GET /{url}", func(writer http.ResponseWriter, request *http.Request) {
		url := request.PathValue("url")
		//discord is so stupid that it only animates gifs if the link ends in .gif like bruh
		if i := strings.LastIndex(url, "."); i != -1{
			url = url[:i]
		}
		query := "SELECT url FROM urls WHERE hash = ? LIMIT 1"
		var urlToSend string
		err := db.QueryRow(query, url).Scan(&urlToSend)
		if err != nil{
			//writer.WriteHeader(http.StatusNotFound)
			http.NotFound(writer, request) //oooooo
			return
		}
		http.Redirect(writer, request, urlToSend, http.StatusFound)
		fmt.Printf("%v asjdhasjdhasjk",urlToSend)

	})

	mux.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "index.html")
	})

	fmt.Println("Starting API server")
	if err := http.ListenAndServe("0.0.0.0:6969", mux); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}


//func randomString(n int) string {
//    b := make([]byte, n)
//    _, err := rand.Read(b)
//    if err != nil {
//        panic(err)
//    }
//    return hex.EncodeToString(b)[:n]
//}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS urls (
		hash TEXT PRIMARY KEY,
		url TEXT NOT NULL
	)`
	if _,err := db.Exec(query); err != nil {
		panic(err)
	}
}
