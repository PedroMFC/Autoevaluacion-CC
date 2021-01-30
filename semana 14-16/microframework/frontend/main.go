package main

import (
   "io/ioutil"
   "log"
   "net/http"
)

func main() {
   // Cuidado con la IP es server porque es la que hemos puesto como servicio en el docker-compose (NO LOCALHOST!!!!!!)
   resp, err := http.Get("http://server:8080/ping") 
   if err != nil {
      log.Fatalln(err)
   }

//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   log.Printf(sb)
}