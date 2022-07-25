package routes

import (
	"blog/config"
	"blog/models"
	"net/http"
)

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	_ = config.Templates.ExecuteTemplate(w, "login.html", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.AuthenticateUser(username, password)
	if err != nil {
		switch err {
		case models.ErrUserNotFound:
			config.Templates.ExecuteTemplate(w, "login.html", "unkown user")
		case models.ErrInvalidLogin:
			config.Templates.ExecuteTemplate(w, "login.html", "invalid login")
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
		}
		return
	}
	session, _ := config.Mysession.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)

}

func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	_ = config.Templates.ExecuteTemplate(w, "register.html", nil)
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.RegisterUser(username, password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	http.Redirect(w, r, "/login", 302)

}

func IndexGetHandler(w http.ResponseWriter, r *http.Request) {

	comments, err := models.GetComments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	} else {
		_ = config.Templates.ExecuteTemplate(w, "index.html", comments)
	}

}

func IndexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	models.PostComments(comment)

	http.Redirect(w, r, "/", 302)

}
