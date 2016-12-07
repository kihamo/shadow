package config

import (
	"flag"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/kihamo/gotypes"
	"github.com/kihamo/shadow"
	"github.com/rakyll/globalconf"
)

const (
	FlagConfig = "config"
)

type Resource struct {
	mutex       sync.RWMutex
	application *shadow.Application
	config      *globalconf.GlobalConf
	values      map[string]interface{}
}

type Variable struct {
	Key   string
	Value interface{}
	Usage string
}

type ContextItemConfigurable interface {
	GetConfigVariables() []Variable
}

func (r *Resource) GetName() string {
	return "config"
}

func (r *Resource) Init(a *shadow.Application) (err error) {
	r.application = a

	config := flag.String(FlagConfig, "", "Config file which which override default config parameters")
	flag.Parse()

	opts := globalconf.Options{
		EnvPrefix: strings.ToUpper(strings.Replace(a.Name, " ", "_", -1)) + "_",
		Filename:  *config,
	}

	if r.config, err = globalconf.NewWithOptions(&opts); err != nil {
		return err
	}

	r.mutex.Lock()
	r.values = map[string]interface{}{}
	r.mutex.Unlock()

	return err
}

func (r *Resource) Run() error {
	r.Add("debug", false, "Debug mode")

	for _, resource := range r.application.GetResources() {
		if configurable, ok := resource.(ContextItemConfigurable); ok {
			for _, variable := range configurable.GetConfigVariables() {
				r.Add(variable.Key, variable.Value, variable.Usage)
			}
		}
	}

	for _, service := range r.application.GetServices() {
		if configurable, ok := service.(ContextItemConfigurable); ok {
			for _, variable := range configurable.GetConfigVariables() {
				r.Add(variable.Key, variable.Value, variable.Usage)
			}
		}
	}

	r.config.ParseAll()

	return nil
}

func (r *Resource) Add(key string, value interface{}, usage string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	switch value.(type) {
	case bool:
		r.values[key] = flag.Bool(key, value.(bool), usage)
	case int:
		r.values[key] = flag.Int(key, value.(int), usage)
	case int64:
		r.values[key] = flag.Int64(key, value.(int64), usage)
	case uint:
		r.values[key] = flag.Uint(key, value.(uint), usage)
	case uint64:
		r.values[key] = flag.Uint64(key, value.(uint64), usage)
	case float64:
		r.values[key] = flag.Float64(key, value.(float64), usage)
	case string:
		r.values[key] = flag.String(key, value.(string), usage)
	case time.Duration:
		r.values[key] = flag.Duration(key, value.(time.Duration), usage)
	}
}

func (r *Resource) Has(key string) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if _, ok := r.values[key]; ok {
		return true
	}

	return false
}

func (r *Resource) Get(key string) interface{} {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if val, ok := r.values[key]; ok {
		return reflect.Indirect(reflect.ValueOf(val)).Interface()
	}

	return nil
}

func (r *Resource) GetAll() map[string]interface{} {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.values
}

func (r *Resource) GetGlobalConf() *globalconf.GlobalConf {
	return r.config
}

func (r *Resource) GetBool(key string) bool {
	return r.GetBoolDefault(key, false)
}

func (r *Resource) GetBoolDefault(key string, value interface{}) bool {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if val, ok := r.values[key]; ok {
		return gotypes.ToBool(val)
	}

	return gotypes.ToBool(value)
}

func (r *Resource) GetInt(key string) int {
	return r.GetIntDefault(key, -1)
}

func (r *Resource) GetIntDefault(key string, value interface{}) int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if val, ok := r.values[key]; ok {
		return gotypes.ToInt(val)
	}

	return gotypes.ToInt(value)
}

func (r *Resource) GetInt64(key string) int64 {
	return r.GetInt64Default(key, -1)
}

func (r *Resource) GetInt64Default(key string, value interface{}) int64 {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if val, ok := r.values[key]; ok {
		return gotypes.ToInt64(val)
	}

	return gotypes.ToInt64(value)
}

func (r *Resource) GetUint(key string) uint {
	return r.GetUintDefault(key, 0)
}

func (r *Resource) GetUintDefault(key string, value interface{}) uint {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if val, ok := r.values[key]; ok {
		return gotypes.ToUint(val)
	}

	return gotypes.ToUint(value)
}

func (r *Resource) GetUint64(key string) uint64 {
	return r.GetUint64Default(key, 0)
}

func (r *Resource) GetUint64Default(key string, value interface{}) uint64 {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if val, ok := r.values[key]; ok {
		return gotypes.ToUint64(val)
	}

	return gotypes.ToUint64(value)
}

func (r *Resource) GetFloat64(key string) float64 {
	return r.GetFloat64Default(key, -1)
}

func (r *Resource) GetFloat64Default(key string, value interface{}) float64 {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if val, ok := r.values[key]; ok {
		return gotypes.ToFloat64(val)
	}

	return gotypes.ToFloat64(value)
}

func (r *Resource) GetString(key string) string {
	return r.GetStringDefault(key, "")
}

func (r *Resource) GetStringDefault(key string, value interface{}) string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if val, ok := r.values[key]; ok {
		return gotypes.ToString(val)
	}

	return gotypes.ToString(value)
}

func (r *Resource) GetDuration(key string) time.Duration {
	return r.GetDurationDefault(key, 0)
}

func (r *Resource) GetDurationDefault(key string, value time.Duration) time.Duration {
	if val := r.GetString(key); val != "" {
		if r, err := time.ParseDuration(val); err == nil {
			return r
		}
	}

	return value
}