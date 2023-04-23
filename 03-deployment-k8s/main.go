// v2
// package main

// import (
// 	"io"
// 	"net/http"
// )

// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "[v2] Hello, Kubernetes!")
// }

//	func main() {
//		http.HandleFunc("/", hello)
//		http.ListenAndServe(":3000", nil)
//	}

// livenessProb
// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"
// )

// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "[v2] Hello, Kubernetes!")
// }

// func main() {
// 	started := time.Now()
// 	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
// 		duration := time.Since(started)
// 		if duration.Seconds() > 15 {
// 			w.WriteHeader(500)
// 			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
// 		} else {
// 			w.WriteHeader(200)
// 			w.Write([]byte("ok"))
// 		}
// 	})

//		http.HandleFunc("/", hello)
//		http.ListenAndServe(":3000", nil)
//	}

// bad
package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[v2] Hello, Kubernetes!")
}

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
