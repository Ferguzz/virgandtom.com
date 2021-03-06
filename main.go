package virgandtom

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := `<!DOCTYPE html>
                 <html>
                     <body>
                     <div style="margin-right:auto;margin-left:auto;width:640px">
                         <img src="img/virgandtom.jpg" height="480px" width="640px">
                         </div>
                     </body>
                 </html>`
	fmt.Fprint(w, response)
}
