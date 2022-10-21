package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil{
		log.Printf("error parsing template: %v", err)
		http.Error(w, "There was an error parsing home template.", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, "")
	if err != nil{
		log.Printf("error executing template: %v", err)
		http.Error(w, "There was an error executing home template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Contact Page</h1><p> To get in touch, email me at " +
	"<a href=\"mailto:anupghimire96@gmail.com\">anupghimire96@gmail.com</a>")
}

/* Practice using URLParam
func galleriesHandler(w http.ResponseWriter, r *http.Request){
	pageID := chi.URLParam(r, "galleries")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("<h1> Under construction... please visit later.</h1> %v", pageID)))
}
 */

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

	//r.Get("/{pageID}", galleriesHandler)
	
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found.", http.StatusNotFound)
	})

	fmt.Println("Starting server on :4500....")
	http.ListenAndServe(":4500", r)
}



