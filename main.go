package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Welcome to Lenslocked, 'tis a wondersite homioes!")
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Contact Page</h1><p> To get in touch, email me at " +
	"<a href=\"mailto:anupghimire96@gmail.com\">anupghimire96@gmail.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
		<ul>
		  <li><b>What is your price structure like?</b></li>
			<li>I charge the lowest market price which changes regularly, so it might come up around $350-$400 per hour.</li>
<p>
		  <li><b>How long does it take you to deliver your work?</b></li>
			<li>It will take at most two weeks to get everything done. I give 5% discount everyday I am late to deliver the work.</li>
<p>
		  <li><b>How do I contact support?</b></li>
			<li>Email us - <a href="mailto:anupghimire96@gmail.com">anupghimire96@gmail.com</a></li>
		</ul>
		`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found.", http.StatusNotFound)
	})

	fmt.Println("Starting server on :4500....")
	http.ListenAndServe(":4500", r)
}

