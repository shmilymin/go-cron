package test

import (
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func TestCorn(t *testing.T) {
	sc := cron.New(cron.WithSeconds())
	sc.Start()
	defer sc.Stop()
	for i := 0; i < 10; i++ {
		spec := "*/5 * * * * ?"
		cornId, _ := sc.AddFunc(spec, func() {
			t.Logf("task i:%d", i)
		})
		t.Logf("spec:%s cornId:%d", spec, cornId)
	}
	time.Sleep(1 * time.Hour)
}
