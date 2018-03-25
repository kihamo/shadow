package internal

import (
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/logger"
)

func (c *Component) ConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(
			logger.ConfigLevel,
			config.ValueTypeInt,
			logger.LevelInformational,
			"Log level in format RFC5424",
			true,
			"",
			[]string{config.ViewEnum},
			map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{logger.LevelEmergency, "Emergency"},
					{logger.LevelAlert, "Alert"},
					{logger.LevelCritical, "Critical"},
					{logger.LevelError, "Error"},
					{logger.LevelWarning, "Warning"},
					{logger.LevelNotice, "Notice"},
					{logger.LevelInformational, "Informational"},
					{logger.LevelDebug, "Debug"},
				},
			}),
		config.NewVariable(
			logger.ConfigFields,
			config.ValueTypeString,
			nil,
			"Fields in format field_name=field1_value,field2_name=field2_value",
			true,
			"",
			[]string{config.ViewTags},
			map[string]interface{}{
				config.ViewOptionTagsDefaultText: "add a field",
			}),
	}
}

func (c *Component) ConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher(logger.ComponentName, []string{logger.ConfigLevel}, c.watchLoggerLevel),
		config.NewWatcher(logger.ComponentName, []string{logger.ConfigFields}, c.watchLoggerFields),
	}
}

func (c *Component) watchLoggerLevel(_ string, newValue interface{}, _ interface{}) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	level := c.getXLogLevel()

	for key, _ := range c.loggers {
		c.loggers[key].(*loggerWrapper).setLevel(level)
	}
}

func (c *Component) watchLoggerFields(_ string, newValue interface{}, oldValue interface{}) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	removeFields := map[string]struct{}{}
	newFields := c.parseFields(newValue.(string))
	oldFields := c.parseFields(oldValue.(string))

	for k, _ := range oldFields {
		if _, ok := newFields[k]; !ok {
			removeFields[k] = struct{}{}
		}
	}

	globalFields := c.getFields()

	for key, _ := range c.loggers {
		l := c.loggers[key].(*loggerWrapper)

		existsFields := map[string]interface{}{}

		for k, v := range l.GetFields() {
			if _, ok := removeFields[k]; !ok {
				existsFields[k] = v
			}
		}

		for k, v := range globalFields {
			existsFields[k] = v
		}

		l.setFields(existsFields)
	}
}
