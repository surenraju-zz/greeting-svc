package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

//LogstashFormatter Set formatters and fields
type LogstashFormatter struct {
	logrus.Formatter
	logrus.Fields
}

var (
	logstashFieldMap = logrus.FieldMap{
		logrus.FieldKeyTime: "timestamp",
		logrus.FieldKeyMsg:  "message",
		logrus.FieldKeyFile: "codelocation",
	}
)

func LogGenerator(context map[string]interface{}) *logrus.Entry {
	return Generator(context, "log")
}

func EventLogGenerator(context map[string]interface{}) *logrus.Entry {
	return Generator(context, "event")
}

func Logger() *logrus.Entry {
	return Generator(make(map[string]interface{}), "log")
}

func EventLogger() *logrus.Entry {
	return Generator(make(map[string]interface{}), "event")
}

//Generator generates logs with given fields in given loglevel
func Generator(context map[string]interface{}, logtype string) *logrus.Entry {
	if len(logtype) == 0 {
		logtype = "log"
	}
	svc := os.Getenv("SERVICE_NAME")
	if len(svc) == 0 {
		svc = "cosmos-undefined-ms"
	}
	loglevel := os.Getenv("LOG_LEVEL")
	if len(loglevel) == 0 {
		loglevel = "info"
	}
	loglvl, err := logrus.ParseLevel(loglevel)
	if err != nil {
		loglvl = logrus.InfoLevel
	}
	log.SetLevel(loglvl)
	log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02T15:04:05.000-07:00", FieldMap: logstashFieldMap})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	fields := make(map[string]interface{})
	fields["type"] = logtype
	fields["app"] = svc
	for k, v := range context {
		fields[k] = v
	}
	contextLog := log.WithFields(fields)
	return contextLog
}
