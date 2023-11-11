package alert

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

const ALERT_LOG = "ALERT"

type Monitor struct {
	Interval time.Duration
	Rules    []Rule
	alerts   []Alert
}

type Alert struct {
	Timestamp time.Time
	AlertName string
	Reviewer  string
	AlertData string
}

func NewAlertMonitoring(monitor *Monitor) {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	for {
		time.Sleep(monitor.Interval)
		monitor.checkFeatures(time.Now())
		monitor.DisplayAlerts()
	}
}

func (monitor *Monitor) checkFeatures(date time.Time) {
	for _, v := range monitor.Rules {
		if v.CreateAlert() {
			monitor.alerts = append(monitor.alerts, Alert{Timestamp: date, AlertName: v.GetErrorName(), Reviewer: v.GetReviewer(), AlertData: fmt.Sprintf("%v", 100-(v.GetErrorData()).(float32))})
		}
	}
}

func (monitor *Monitor) DisplayAlerts() {
	for _, v := range monitor.alerts {
		log.WithTime(v.Timestamp).Warningf("[%s] [%s] [%s] -> sent to [%s]", ALERT_LOG, strings.ToUpper(v.AlertName), v.AlertData, v.Reviewer)
	}
}

func (monitor *Monitor) GetAlerts() []Alert {
	return monitor.alerts
}
