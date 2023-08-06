/*
 * Copyright 2022 Holoinsight Project Authors. Licensed under Apache-2.0.
 */

package mysql

import (
	"encoding/json"
	"fmt"
	telegraf2 "github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs/mysql"
	"github.com/traas-stack/holoinsight-agent/pkg/collecttask"
	"github.com/traas-stack/holoinsight-agent/pkg/logger"
	"github.com/traas-stack/holoinsight-agent/pkg/plugin/api"
	"github.com/traas-stack/holoinsight-agent/pkg/plugin/input/standard/providers"
	"github.com/traas-stack/holoinsight-agent/pkg/telegraf"
)

type (
	Conf struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Port     int    `json:"port,omitempty"`
	}
)

func init() {
	providers.RegisterInputProvider("telegraf_mysql", func(task *collecttask.CollectTask) (api.Input, error) {
		conf := &Conf{}
		if err := json.Unmarshal(task.Config.Content, conf); err != nil {
			return nil, err
		}

		var telegrafInput telegraf2.Input
		ip := task.Target.GetIP()
		if ip != "" {
			if conf.Port <= 0 {
				conf.Port = 3306
			}
			server := fmt.Sprintf("%s:%s@tcp(%s:%d)/?tls=false&timeout=2s", conf.Username, conf.Password, task.Target.GetIP(), conf.Port)
			telegrafInput = &mysql.Mysql{
				Servers: []string{server},
				Log:     logger.ZapLogger.InfoS,
			}
		} else {
			telegrafInput = &mysql.Mysql{
				Log: logger.ZapLogger.InfoS,
			}
			if err := json.Unmarshal(task.Config.Content, telegrafInput); err != nil {
				return nil, err
			}
		}
		return telegraf.NewInputAdapter(telegrafInput), nil
	})
}
