package server

import (
	"fmt"
	"net/http"
	"github.com/unquabain/gibberish/config"
	"github.com/unquabain/gibberish/lexicon"
	"encoding/json"
)

type Server struct {
	lexicon *lexicon.Lexicon
}

type Gibberish struct {
	Status int `json:"status"`
	Error string `json:"error"`
	Text string `json:"text"`
}

func NewServer(lexicon *lexicon.Lexicon) Server {
	return Server  { lexicon }
}

func (this Server) Serve(port int) {
	http.Handle("/", http.FileServer(config.Web))
	http.Handle("/api/", this)

	fmt.Printf("Listening on http://localhost:%d\n^C to quit\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (this Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received URL %s\n", req.URL.Path)
	switch req.URL.Path {
	case "/api/gibberish":
		var response Gibberish
		text, err := this.lexicon.Evaluate()
		if err != nil {
			response.Status = 0
			response.Error = fmt.Sprintf("Could not evaluate: %v", err)
		} else {
			response.Status = 1
			response.Text = text
		}
		fmt.Printf("Evaluated text %s\n", text)

		m, err := json.Marshal(response)
		resp.Header().Add("Content-Type", "application/json")
		fmt.Fprint(resp, string(m))
	default:
		fmt.Println("Don't know what to do with that.")
	}
}

