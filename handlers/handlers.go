package handlers

import (
	"blog/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	_ = config.Templates.ExecuteTemplate(w, "login.html", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	hash, err := config.Client.Get(config.Ctx, "user:"+username).Bytes()
	config.CheckError(err)
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	config.CheckError(err)

	session, _ := config.Mysession.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/index", 302)

}

func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	_ = config.Templates.ExecuteTemplate(w, "register.html", nil)
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	config.CheckError(err)
	config.Client.Set(config.Ctx, "user:"+username, hash, 0)
	http.Redirect(w, r, "/login", 302)

}

func IndexGetHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Mysession.Get(r, "session")
	_, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
		return
	} else {
		comments, err := config.Client.LRange(config.Ctx, "comments", 0, 100).Result()
		config.CheckError(err)
		_ = config.Templates.ExecuteTemplate(w, "index.html", comments)
	}
}

func IndexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	config.Client.LPush(config.Ctx, "comments", comment)
	http.Redirect(w, r, "/index", 302)
}
