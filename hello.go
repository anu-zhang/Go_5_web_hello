package main
import ( "io"
	"log"
	"net/http" )
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}
func main() {
	// 路径
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
