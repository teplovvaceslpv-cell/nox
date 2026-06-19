package main
import "net/http"

func main(){
  mux := http.NewServeMux()
  mux.HandleFunc("/hello", getHello)

  lambda.Start(httpadapter.New(mux).ProxyWithContext)

func getHello(w mux.ResponseWriter, r *mux.Request){
  mux.ServeFile(w, r, "hello.html")

  
