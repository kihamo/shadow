package workers

import (
	"sync"

	"github.com/kihamo/go-workers/dispatcher"
	"github.com/kihamo/go-workers/task"
	"github.com/kihamo/go-workers/worker"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/shadow/components/metrics"
)

const (
	ComponentName = "workers"
)

type Component struct {
	application shadow.Application
	config      *config.Component
	logger      logger.Logger

	dispatcher *dispatcher.Dispatcher
}

func (c *Component) GetName() string {
	return ComponentName
}

func (c *Component) GetVersion() string {
	return ComponentVersion
}

func (c *Component) GetDependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: dashboard.ComponentName,
		},
		{
			Name: logger.ComponentName,
		},
		{
			Name: metrics.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a
	c.config = a.GetComponent(config.ComponentName).(*config.Component)
	c.dispatcher = dispatcher.NewDispatcher()

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) (err error) {
	c.logger = logger.NewOrNop(c.GetName(), c.application)
	c.setLogListener(wg)

	for i := 1; i <= c.config.GetInt(ConfigCount); i++ {
		c.AddWorker()
	}

	go func() {
		defer wg.Done()
		c.dispatcher.Run()

		c.dispatcher.SetTickerExecuteTasksDuration(c.config.GetDuration(ConfigTickerExecuteTasksDuration))
		c.dispatcher.SetTickerNotifyListenersDuration(c.config.GetDuration(ConfigTickerNotifyListenersDuration))
	}()

	return nil
}

func (c *Component) getDefaultListenerName() string {
	return c.GetName() + ".all"
}

func (c *Component) setLogListener(wg *sync.WaitGroup) {
	listener := dispatcher.NewDefaultListener(c.getDefaultListenerName())
	c.dispatcher.AddListener(listener)

	// logger for finished tasks
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case t := <-listener.GetTaskDoneChannel():
				status := t.GetStatus()

				switch status {
				case task.TaskStatusWait:
					c.logger.Debug("Finished", c.getLogFieldsForTask(t, map[string]interface{}{"task.status": "wait"}))

					if metricTasksTotal != nil {
						metricTasksTotal.With("status", "wait").Inc()
					}
				case task.TaskStatusProcess:
					c.logger.Debug("Finished", c.getLogFieldsForTask(t, map[string]interface{}{"task.status": "process"}))

					if metricTasksTotal != nil {
						metricTasksTotal.With("status", "process").Inc()
					}
				case task.TaskStatusSuccess:
					c.logger.Debug("Success finished", c.getLogFieldsForTask(t, map[string]interface{}{"task.status": "success"}))

					if metricTasksTotal != nil {
						metricTasksTotal.With("status", "success").Inc()
					}
				case task.TaskStatusFail:
					c.logger.Error("Fail finished", c.getLogFieldsForTask(t, map[string]interface{}{"task.status": "fail"}))

					if metricTasksTotal != nil {
						metricTasksTotal.With("status", "fail").Inc()
					}
				case task.TaskStatusFailByTimeout:
					c.logger.Error("Fail by timeout finished", c.getLogFieldsForTask(t, map[string]interface{}{"task.status": "fail-by-timeout"}))

					if metricTasksTotal != nil {
						metricTasksTotal.With("status", "fail-by-timeout").Inc()
					}
				case task.TaskStatusKill:
					c.logger.Warn("Execute killed", c.getLogFieldsForTask(t, map[string]interface{}{"task.status": "kill"}))

					if metricTasksTotal != nil {
						metricTasksTotal.With("status", "kill").Inc()
					}
				case task.TaskStatusRepeatWait:
					c.logger.Debug("Repeat execute", c.getLogFieldsForTask(t, map[string]interface{}{"task.status": "repeat-wait"}))

					if metricTasksTotal != nil {
						metricTasksTotal.With("status", "repeat-wait").Inc()
					}
				default:
					c.logger.Warnf("Unknown task status %s", status, c.getLogFieldsForTask(t, map[string]interface{}{"task.status": "unknown"}))

					if metricTasksTotal != nil {
						metricTasksTotal.With("status", "unknown").Inc()
					}
				}
			}
		}
	}()
}

func (c *Component) AddTask(t task.Tasker) {
	c.dispatcher.AddTask(t)
	c.logger.Debug("Add task", c.getLogFieldsForTask(t, nil))

	if metricTasksTotal != nil {
		metricTasksTotal.Inc()
	}
}

func (c *Component) AddNamedTaskByFunc(n string, f task.TaskFunction, a ...interface{}) task.Tasker {
	t := c.dispatcher.AddNamedTaskByFunc(n, f, a...)
	c.logger.Debug("Add task", c.getLogFieldsForTask(t, nil))

	if metricTasksTotal != nil {
		metricTasksTotal.Inc()
	}

	return t
}

func (c *Component) AddTaskByFunc(f task.TaskFunction, a ...interface{}) task.Tasker {
	t := c.dispatcher.AddTaskByFunc(f, a...)
	c.logger.Debug("Add task", c.getLogFieldsForTask(t, nil))

	if metricTasksTotal != nil {
		metricTasksTotal.Inc()
	}

	return t
}

func (c *Component) AddTaskByPriorityAndFunc(p int64, f task.TaskFunction, a ...interface{}) task.Tasker {
	t := c.dispatcher.AddTaskByPriorityAndFunc(p, f, a...)
	c.logger.Debug("Add task", c.getLogFieldsForTask(t, nil), map[string]interface{}{"priority": p})

	if metricTasksTotal != nil {
		metricTasksTotal.Inc()
	}

	return t
}

func (c *Component) RemoveTask(t task.Tasker) {
	c.dispatcher.RemoveTask(t)
	c.logger.Debug("Removed task", map[string]interface{}{"task.id": t.GetId()})

	if metricTasksTotal != nil {
		metricTasksTotal.Dec()
	}
}

func (c *Component) RemoveTaskById(id string) {
	c.dispatcher.RemoveTaskById(id)
	c.logger.Debug("Removed task", map[string]interface{}{"task.id": id})

	if metricTasksTotal != nil {
		metricTasksTotal.Dec()
	}
}

func (c *Component) AddWorker() {
	w := c.dispatcher.AddWorker()
	c.logger.Debug("Added worker", map[string]interface{}{"worker.id": w.GetId()})

	if metricWorkersTotal != nil {
		metricWorkersTotal.Inc()
	}
}

func (c *Component) RemoveWorker(w worker.Worker) {
	c.dispatcher.RemoveWorker(w)
	c.logger.Debug("Removed worker", map[string]interface{}{"worker.id": w.GetId()})

	go func() {
		w.Kill()
		c.logger.Debug("Killed worker", map[string]interface{}{"worker.id": w.GetId()})
	}()

	if metricWorkersTotal != nil {
		metricWorkersTotal.Dec()
	}
}

func (c *Component) GetWorkers() []worker.Worker {
	return c.dispatcher.GetWorkers()
}

func (c *Component) getLogFieldsForTask(t task.Tasker, l map[string]interface{}) map[string]interface{} {
	fields := map[string]interface{}{
		"task.id":       t.GetId(),
		"task.function": t.GetFunctionName(),
		"task.priority": t.GetPriority(),
		"task.name":     t.GetName(),
		"task.duration": t.GetDuration().String(),
		"task.repeats":  t.GetRepeats(),
		"task.attemps":  t.GetAttempts(),
	}

	if lastError := t.GetLastError(); lastError != nil {
		fields["task.error"] = lastError
	}

	for k, v := range l {
		fields[k] = v
	}

	return fields
}

func (c *Component) AddListener(l dispatcher.Listener) {
	c.dispatcher.AddListener(l)
	c.logger.Debug("Added listener", map[string]interface{}{"listener.name": l.GetName()})

	if metricListenersTotal != nil {
		metricListenersTotal.Inc()
	}
}

func (c *Component) RemoveListener(l dispatcher.Listener) {
	c.dispatcher.RemoveListener(l)
	c.logger.Debug("Remove dlistener", map[string]interface{}{"listener.name": l.GetName()})

	if metricListenersTotal != nil {
		metricListenersTotal.Dec()
	}
}
