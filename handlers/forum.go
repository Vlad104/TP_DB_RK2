package handlers

import (
	"net/http"
	"io/ioutil"
	"github.com/Vlad104/TP_DB_RK2/database"
	"github.com/Vlad104/TP_DB_RK2/models"
	"github.com/gorilla/mux"
	"github.com/go-openapi/swag"
)

// /forum/create Создание форума
func CreateForum(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)	
	defer r.Body.Close()
	if err != nil {
		makeResponse(w, 500, []byte(err.Error()))
		return
	}	
	forum := &models.Forum{}
	err = forum.UnmarshalJSON(body)

	if err != nil {
		makeResponse(w, 500, []byte(err.Error()))
		return
	}

	result, err := database.CreateForumDB(forum)

	switch err {
	case nil:
		resp, _ := result.MarshalJSON()
		makeResponse(w, 201, resp)
	case database.UserNotFound:
		makeResponse(w, 404, []byte(makeErrorUser(forum.User)))
	case database.ForumIsExist:
		resp, _ := result.MarshalJSON()
		makeResponse(w, 409, resp)
	default:		
		makeResponse(w, 500, []byte(err.Error()))
	}
}

// /forum/{slug}/details Получение информации о форуме
func GetForum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	result, err := database.GetForumDB(slug)

	switch err {
	case nil:
		resp, _ := result.MarshalJSON()
		makeResponse(w, 200, resp)
	case database.ForumNotFound:
		makeResponse(w, 404, []byte(makeErrorForum(slug)))
	default:		
		makeResponse(w, 500, []byte(err.Error()))
	}
}

// /forum/{slug}/create Создание ветки
func CreateForumThread(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	body, err := ioutil.ReadAll(r.Body)	
	defer r.Body.Close()
	if err != nil {
		makeResponse(w, 500, []byte(err.Error()))
		return
	}	
	thread := &models.Thread{}
	err = thread.UnmarshalJSON(body)
	thread.Forum = slug // иначе не знаю как

	if err != nil {
		makeResponse(w, 500, []byte(err.Error()))
		return
	}

	result, err := database.CreateForumThreadDB(thread)

	switch err {
	case nil:
		resp, _ := result.MarshalJSON()
		makeResponse(w, 201, resp)
	case database.ForumOrAuthorNotFound:
		makeResponse(w, 404, []byte(makeErrorUser(slug)))
	case database.ThreadIsExist:
		resp, _ := result.MarshalJSON()
		makeResponse(w, 409, resp)
	default:		
		makeResponse(w, 500, []byte(err.Error()))
	}
}


// /forum/{slug}/threads Список ветвей обсужления форума
func GetForumThreads(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	queryParams := r.URL.Query()
	var limit, since, desc string
	if limit = queryParams.Get("limit"); limit == "" {
		limit = "1";
	}
	since = queryParams.Get("since");
	// if since = queryParams.Get("since"); since == "" {
	// 	since = "";
	// }
	if desc = queryParams.Get("desc"); desc == ""{
		desc = "false";
	}

	result, err := database.GetForumThreadsDB(slug, limit, since, desc)	
	switch err {
	case nil:
		resp, _ := swag.WriteJSON(result)
		makeResponse(w, 200, resp)
	case database.ForumNotFound:
		makeResponse(w, 404, []byte(makeErrorForum(slug)))
	default:		
		makeResponse(w, 500, []byte(err.Error()))
	}
}

// /forum/{slug}/users Пользователи данного форума
func GetForumUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]
	queryParams := r.URL.Query()
	var limit, since, desc string
	if limit = queryParams.Get("limit"); limit == "" {
		limit = "1";
	}
	since = queryParams.Get("since");
	// if since = queryParams.Get("since"); since == "" {
	// 	since = "";
	// }
	if desc = queryParams.Get("desc"); desc == ""{
		desc = "false";
	}

	result, err := database.GetForumUsersDB(slug, limit, since, desc)
	
	switch err {
	case nil:
		resp, _ := swag.WriteJSON(result) // можно через easyjson, но мне лень было
		makeResponse(w, 200, resp)
	case database.ForumNotFound:
		makeResponse(w, 404, []byte(makeErrorUser(slug)))
	default:		
		makeResponse(w, 500, []byte(err.Error()))
	}
}