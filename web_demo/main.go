package main


func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "hello golang")
}