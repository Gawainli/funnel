package logutil

import (
	"net"
	"testing"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

func Test_logstr(t *testing.T) {
	log := logrus.New()
	conn, err := net.Dial("tcp", "192.168.0.90:1028")
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"@type": "mytestapp"}))

	if err != nil {
		log.Fatal(err)
	}
	log.Hooks.Add(hook)
	ctx := log.WithFields(logrus.Fields{
		"method": "main",
	})
	ctx.Info("Hello3World!!!!")
}

func Test_logredis(t *testing.T) {
	hookConfig := logredis.HookConfig{
		Host:     "192.168.0.98",
		Key:      "logstash-test-blist",
		Format:   "v1",
		App:      "my_app_name",
		Port:     6383,
		Hostname: "my_app_hostmame", // will be sent to field @source_host
		DB:       0,                 // optional
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}

	// when hook is injected succesfully, logs will be sent to redis server
	// logrus.Info("just some info 2 logging...")

	// we also support log.WithFields()
	logrus.WithFields(logrus.Fields{
		"type": "testlog"}).
		Info("additional fields type")

	// // If you want to disable writing to stdout, use setOutput
	// logrus.SetOutput(ioutil.Discard)
	// logrus.Info("This will only be sent to Redis")
}
