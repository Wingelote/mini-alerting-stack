package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/wingelote/aisprid-alerting/internal/config"
	"github.com/wingelote/aisprid-alerting/pkg/alert"

	"github.com/sirupsen/logrus"
)

var url string
var log *logrus.Logger

const ALERT_LOG = "Metric"

func NewClient(conf config.Config) {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	url = fmt.Sprintf("http://%s:%v", conf.Server.Host, conf.Server.Port)
}

func SendMetrics(stack string, qty float32) {
	client := &http.Client{}
	var buffer bytes.Buffer
	metrics := Metrics{Stack: stack, Resources: qty}
	err := json.NewEncoder(&buffer).Encode(metrics)
	if err != nil {
		log.Fatal(err)
	}
	request, error := http.NewRequest("POST", url+"/metrics", &buffer)
	if error != nil {
		logrus.Error(error)
		return
	}
	defer request.Body.Close()
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, error := client.Do(request)
	if error != nil {
		logrus.Error(error)
		return
	}

	if response.StatusCode == http.StatusCreated {
		log.WithTime(time.Now()).Warningf("[%s] %s -> %v", ALERT_LOG, metrics.Stack, metrics.Resources)
		return
	}
}

func GetAlertHistory() {
	var buffer bytes.Buffer
	client := &http.Client{}
	request, error := http.NewRequest("GET", url+"/alerts-history", &buffer)
	if error != nil {
		logrus.Error(error)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, error := client.Do(request)
	if error != nil {
		logrus.Error(error)
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		alerts := make([]alert.Alert, 0)
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(bodyBytes, &alerts)
		for _, v := range alerts {
			log.WithTime(v.Timestamp).Warningf("%s, %s (value %s)", strings.ToUpper(v.AlertName), v.Reviewer, v.AlertData)
		}
	}
}
