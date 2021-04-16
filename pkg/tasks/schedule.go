package tasks

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// MAGIC_DELTA
const MAGIC_DELTA = float64(600) // seconds

//Things that satisfy this interface can be executed as a Task
type Task interface {
	When() time.Time
	End() time.Time
	Execute() error //A Task has to be able to be run
}

// Tasks that are sequentially executed
type SequentialTasks []Task

// Launch a bunch of sequentail tasks when their time is due
type Cron struct {
	ordered       SequentialTasks
	parentContext context.Context
}

// Create a new timed task
func NewCronTask(ctx context.Context, mis SequentialTasks) *Cron {
	return &Cron{
		ordered:       mis,
		parentContext: ctx,
	}
}

// Run a Crontask
func (c *Cron) Run() {

	var timer *time.Timer
	if len(c.ordered) == 0 {
		logrus.Info("Tasks are finished")
		return
	}

	// c.ordered is order in time.
	for _, task := range c.ordered {
		s := time.Now().Add(time.Second * time.Duration(MAGIC_DELTA))

		if s.After(task.When()) {

			if err := task.Execute(); err != nil {
				logrus.Warnf("Task: %v - ERROR - %s\n", task, err)
			}

			c.ordered = c.ordered[1:]
			continue
		}

		//start - now() == delta
		d := task.When().Sub(time.Now().Add(time.Second * time.Duration(MAGIC_DELTA)))

		logrus.Infof("TASK[ %+v ] - Timer Execution: %f seconds", task, d.Seconds())

		timer = time.NewTimer(d)
		break
	}

	if timer == nil {
		logrus.Warnf("Timer is nil\n")
		return
	}

	// block for the timer
	for {
		select {
		case <-c.parentContext.Done():
			timer.Stop()
			return
		case <-timer.C:
			timer.Stop()
			c.Run()
		}
	}
}
