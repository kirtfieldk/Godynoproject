package dynowebportal

import(
	"net/http"
	"fmt"
)

func RunWebPortal(addr string) error{
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(addr, nil)
	
}
func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Server Started %s", r.RemoteAddr)
}