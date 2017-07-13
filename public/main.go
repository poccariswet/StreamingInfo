package main

import ()

func main() {
	var addr = flag.String("addr", ":8081", "webアドレス")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir("views"))))
	log.Println("Webサイトアドレス:", *addr)
	http.ListenAndServe(*addr, mux)
}
