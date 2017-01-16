package main

import (
	"net/http"
	"basicGo/routes"
	"basicGo/config"
	"os"
	"fmt"
)
func main() {
	context:= config.Context{}
	errLogFilePath:= "errorLog"
	errorFile,err := os.OpenFile(errLogFilePath,os.O_APPEND|os.O_WRONLY,0600)
	if err!=nil {
		panic(err)
	}
	defer errorFile.Close()

	context.ErrorLogFile=errorFile
	fmt.Println(context)

	routes.HandleRequests(context)
	http.ListenAndServe(":8000",nil)
}
