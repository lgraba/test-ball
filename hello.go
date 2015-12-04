// hello.go
// A Hellow World HTTP serving example application
// Next up: Deploy a guestbook, fool

package hello

import (
	"fmt"
	"html/template"
	"net/http"
	// "appengine"
	// "appengine/user"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, guestbookForm)
}

const guestbookForm = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<title>Test Ballin' Guestbook</title>
	</head>
	<body>
		<form action="/sign" method="post">
			<div><textarea name="content" rows="3" cols="60"></textarea></div>
			<div><input type="submit" value="Sign It, Baby"></div>
		</form>
	</body>
</html>
`

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

const signTemplateHTML = `
<!DOCTYPE html>
<html lang="en">
	<head><meta charset="UTF-8" />
		<title>Test Ballin' Guestbook Entry</title>
	</head>
	<body>
		<p>You wrote:</p>
		<pre>{{.}}</pre>
	</body>
</html>
`