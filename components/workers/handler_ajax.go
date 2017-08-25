package workers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/kihamo/go-workers/dispatcher"
	"github.com/kihamo/go-workers/task"
	"github.com/kihamo/go-workers/worker"
	"github.com/kihamo/shadow/components/dashboard"
)

// easyjson:json
type ajaxHandlerResponseSuccess struct {
	Result string `json:"result"`
}

// easyjson:json
type ajaxHandlerResponseTask struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Status    int64       `json:"status"`
	Priority  int64       `json:"priority"`
	Attempts  int64       `json:"attempts"`
	LastError interface{} `json:"last_error"`
	Created   time.Time   `json:"created"`
}

// easyjson:json
type ajaxHandlerResponseWorker struct {
	Id      string                   `json:"id"`
	Created time.Time                `json:"created"`
	Status  int64                    `json:"status"`
	Task    *ajaxHandlerResponseTask `json:"task"`
}

// easyjson:json
type ajaxHandlerResponseListener struct {
	Name              string     `json:"name"`
	CreatedAt         time.Time  `json:"created_at"`
	LastTaskSuccessAt *time.Time `json:"last_task_success_at"`
	LastTaskFailedAt  *time.Time `json:"last_task_failed_at"`
	CountTaskSuccess  uint64     `json:"count_task_success"`
	CountTaskFailed   uint64     `json:"count_task_failed"`
}

// easyjson:json
type ajaxHandlerResponse struct {
	Tasks []ajaxHandlerResponseTask `json:"tasks_wait"`

	Workers      []ajaxHandlerResponseWorker `json:"workers"`
	WorkersCount int                         `json:"workers_count"`
	WorkersWait  int                         `json:"workers_wait"`
	WorkersBusy  int                         `json:"workers_busy"`

	Listeners      []ajaxHandlerResponseListener `json:"listeners"`
	ListenersCount int                           `json:"listeners_count"`
}

type AjaxHandler struct {
	dashboard.Handler

	component *Component
}

func (h *AjaxHandler) fillResponseTask(t task.Tasker) *ajaxHandlerResponseTask {
	return &ajaxHandlerResponseTask{
		Id:        t.GetId(),
		Name:      t.GetName(),
		Status:    t.GetStatus(),
		Priority:  t.GetPriority(),
		Attempts:  t.GetAttempts(),
		LastError: t.GetLastError(),
		Created:   t.GetCreatedAt(),
	}
}

func (h *AjaxHandler) actionStats(w http.ResponseWriter, r *http.Request) {
	tasksList := []ajaxHandlerResponseTask{}

	for _, t := range h.component.dispatcher.GetTasks() {
		switch t.GetStatus() {
		case task.TaskStatusWait, task.TaskStatusRepeatWait:
			tasksList = append(tasksList, *h.fillResponseTask(t))
		}
	}

	workersList := []ajaxHandlerResponseWorker{}
	workersWait := 0
	workersBusy := 0

	for _, wrk := range h.component.dispatcher.GetWorkers() {
		switch wrk.GetStatus() {
		case worker.WorkerStatusWait:
			workersWait++
		case worker.WorkerStatusBusy:
			workersBusy++
		}

		workerInfo := ajaxHandlerResponseWorker{
			Id:      wrk.GetId(),
			Created: wrk.GetCreatedAt(),
			Status:  wrk.GetStatus(),
		}

		t := wrk.GetTask()
		if t != nil {
			workerInfo.Task = h.fillResponseTask(t)

			workersList = append([]ajaxHandlerResponseWorker{workerInfo}, workersList...)
		} else {
			workersList = append(workersList, workerInfo)
		}
	}

	listenersList := []ajaxHandlerResponseListener{}
	listenersCount := 0

	for _, l := range h.component.dispatcher.GetListeners() {
		listenersCount++
		item := l.(dispatcher.ListenerItem)

		listenersList = append(listenersList, ajaxHandlerResponseListener{
			Name:              item.GetName(),
			CreatedAt:         item.GetCreatedAt(),
			LastTaskSuccessAt: item.GetLastTaskSuccessAt(),
			LastTaskFailedAt:  item.GetLastTaskFailedAt(),
			CountTaskSuccess:  item.GetCountTaskSuccess(),
			CountTaskFailed:   item.GetCountTaskFailed(),
		})
	}

	stats := ajaxHandlerResponse{
		Tasks: tasksList,

		Workers:      workersList,
		WorkersCount: len(workersList),
		WorkersWait:  workersWait,
		WorkersBusy:  workersBusy,

		Listeners:      listenersList,
		ListenersCount: listenersCount,
	}

	dashboard.ResponseFromContext(r.Context()).SendJSON(stats)
}

func (h *AjaxHandler) actionListenersRemove(w http.ResponseWriter, r *http.Request) {
	listeners := h.component.dispatcher.GetListeners()
	checkId := r.FormValue("id")

	for _, listener := range listeners {
		if checkId == "" || listener.GetName() == checkId {
			if listener.GetName() != h.component.getDefaultListenerName() {
				h.component.RemoveListener(listener)
			}

			if checkId != "" {
				break
			}
		}
	}

	dashboard.ResponseFromContext(r.Context()).SendJSON(ajaxHandlerResponseSuccess{
		Result: "success",
	})
}

func (h *AjaxHandler) actionTaskRemove(w http.ResponseWriter, r *http.Request) {
	removeId := r.FormValue("id")

	if removeId != "" {
		h.component.RemoveTaskById(removeId)
	}

	dashboard.ResponseFromContext(r.Context()).SendJSON(ajaxHandlerResponseSuccess{
		Result: "success",
	})
}

func (h *AjaxHandler) actionWorkersReset(w http.ResponseWriter, r *http.Request) {
	workers := h.component.dispatcher.GetWorkers()
	checkId := r.FormValue("id")

	go func() {
		for _, worker := range workers {
			if checkId == "" || worker.GetId() == checkId {
				worker.Reset()
				h.component.logger.Infof("Reseted worker #%s", worker.GetId())

				if checkId != "" {
					break
				}
			}
		}
	}()

	dashboard.ResponseFromContext(r.Context()).SendJSON(ajaxHandlerResponseSuccess{
		Result: "success",
	})
}

func (h *AjaxHandler) actionWorkersKill(w http.ResponseWriter, r *http.Request) {
	workers := h.component.dispatcher.GetWorkers()
	checkId := r.FormValue("id")

	for _, worker := range workers {
		if checkId == "" || worker.GetId() == checkId {
			h.component.RemoveWorker(worker)

			if checkId != "" {
				break
			}
		}
	}

	dashboard.ResponseFromContext(r.Context()).SendJSON(ajaxHandlerResponseSuccess{
		Result: "success",
	})
}

func (h *AjaxHandler) actionWorkersAdd(w http.ResponseWriter, r *http.Request) {
	count := r.FormValue("count")
	if count != "" {
		if c, err := strconv.Atoi(count); err == nil {
			for i := 0; i < c; i++ {
				h.component.AddWorker()
			}
		}
	}

	dashboard.ResponseFromContext(r.Context()).SendJSON(ajaxHandlerResponseSuccess{
		Result: "success",
	})
}

func (h *AjaxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := dashboard.RequestFromContext(r.Context())

	switch r.URL.Query().Get("action") {
	case "stats":
		h.actionStats(w, r)

	case "listeners-remove":
		if request.IsPost() {
			h.actionListenersRemove(w, r)
		} else {
			h.MethodNotAllowed(w, r)
		}

	case "task-remove":
		if request.IsPost() {
			h.actionTaskRemove(w, r)
		} else {
			h.MethodNotAllowed(w, r)
		}

	case "workers-reset":
		if request.IsPost() {
			h.actionWorkersReset(w, r)
		} else {
			h.MethodNotAllowed(w, r)
		}

	case "workers-kill":
		if request.IsPost() {
			h.actionWorkersKill(w, r)
		} else {
			h.MethodNotAllowed(w, r)
		}

	case "workers-add":
		if request.IsPost() {
			h.actionWorkersAdd(w, r)
		} else {
			h.MethodNotAllowed(w, r)
		}

	default:
		h.NotFound(w, r)
	}
}
