package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
    // Load .env file.
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Loading .env file failed")
    }

    // Create data source json file.
    dir := "sample-files/"
    filePath := dir + "test.json"
    fmt.Println(filePath)
    file, err :=os.Create(filePath)
    if err != nil{
        log.Fatalf("Failed to create the json file: %s", err)
    }
    defer file.Close()

    s3RawDataBucket := os.Getenv("S3_RAW_DATA_BUCKET")
    fmt.Println(s3RawDataBucket)
}