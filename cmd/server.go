package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/wingelote/aisprid-alerting/internal/config"
	"github.com/wingelote/aisprid-alerting/pkg/alert"
	"github.com/wingelote/aisprid-alerting/pkg/stack"

	"github.com/gin-gonic/gin"
)

type Endpoints interface {
	SendDataMetric()
	GetAlertHistory()
}

type Metrics struct {
	Stack     string  `json:"stack" binding:"required"`
	Resources float32 `json:"resources" binding:"required"`
}

var mutex sync.Mutex

func NewServer(conf config.Config) {
	router := gin.Default()
	service := stack.NewStackService(&mutex)

	alertMonitoring := alert.Monitor{
		Interval: conf.Alerting.Interval,
		Rules: []alert.Rule{
			&alert.MaxUsageRule{
				MaxUsage:  conf.CPU.MaxUsage,
				ErrorName: conf.CPU.AlertName,
				GetData:   service.CPU.GetUsage,
				SendTo:    conf.CPU.SendTo,
			},
			&alert.MaxUsageRule{
				MaxUsage:  conf.Memory.MaxUsage,
				ErrorName: conf.Memory.AlertName,
				GetData:   service.Memory.GetUsage,
				SendTo:    conf.Memory.SendTo,
			},
		},
	}

	go alert.NewAlertMonitoring(&alertMonitoring)

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%v", conf.Server.Host, conf.Server.Port),
		Handler: router,
	}

	router.POST("/metrics", func(c *gin.Context) {
		var metrics Metrics

		c.ShouldBindJSON(&metrics)
		service.CPU.IncreaseResourcesUsage(metrics.Resources)
		c.JSON(http.StatusCreated, "")
	})

	router.GET("/alerts-history", func(c *gin.Context) {
		c.JSON(http.StatusOK, alertMonitoring.GetAlerts())
	})

	s.ListenAndServe()
}
