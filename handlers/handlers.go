package handlers

import (
	"blog/config"
	"github.com/go-redis/redis/v9"
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
	err2 := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == redis.Nil {
		config.Templates.ExecuteTemplate(w, "login.html", "unkown user")
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	} else if err2 != nil {
		config.Templates.ExecuteTemplate(w, "login.html", "invalid login")
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

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	err = config.Client.Set(config.Ctx, "user:"+username, hash, 0).Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}
	http.Redirect(w, r, "/login", 302)

}

func IndexGetHandler(w http.ResponseWriter, r *http.Request) {

	comments, err := config.Client.LRange(config.Ctx, "comments", 0, 100).Result()
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
	config.Client.LPush(config.Ctx, "comments", comment)

	http.Redirect(w, r, "/", 302)

}
