package handlers

import (
	"net/http"

	"github.com/wissio/go_final_project/internal/storage/sqlite"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	searchParam := r.URL.Query().Get("search")
	dateParam := ""

	tasks, err := storage.GetTasks(dateParam, searchParam, sqlite.TaskLimit)
	if err != nil {
		http.Error(w, "failed to get tasks", http.StatusInternalServerError)
		return
	}

	// TODO: Add actual response encoding here
	_ = tasks
}
