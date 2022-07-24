package handlers

import (
	"blog/config"
	"net/http"
)

func IndexGetHandler(w http.ResponseWriter, r *http.Request) {

	comments, err := config.Client.LRange(config.Ctx, "comments", 0, 100).Result()
	config.CheckError(err)

	_ = config.Templates.ExecuteTemplate(w, "index.html", comments)
}

func IndexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	config.Client.LPush(config.Ctx, "comments", comment)
	http.Redirect(w, r, "/", 302)
}
