package executor

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.k6.io/k6/lib"
	"go.k6.io/k6/lib/executor"
	"go.k6.io/k6/lib/metrics"
	"go.k6.io/k6/stats"
	"go.k6.io/k6/ui/pb"
)

func init() {
	lib.RegisterExecutorConfigType("mycoolexecutor", func(name string, rawjson []byte) (lib.ExecutorConfig, error) {
		return &coolexecutorconfig{}, nil
	})
}

type coolexecutorconfig struct{}

func (c *coolexecutorconfig) Validate() []error {
	return nil
}

func (c *coolexecutorconfig) GetName() string {
	return "really cool"
}

func (c *coolexecutorconfig) GetType() string {
	return "some type"
}

func (c *coolexecutorconfig) GetStartTime() time.Duration {
	return 0
}

func (c *coolexecutorconfig) GetGracefulStop() time.Duration {
	return 0
}

func (c *coolexecutorconfig) IsDistributable() bool {
	return true
}

func (c *coolexecutorconfig) GetEnv() map[string]string {
	return nil
}

func (c *coolexecutorconfig) GetExec() string {
	return "default"
}

func (c *coolexecutorconfig) GetTags() map[string]string {
	return nil
}

func (c *coolexecutorconfig) GetExecutionRequirements(*lib.ExecutionTuple) []lib.ExecutionStep {
	return []lib.ExecutionStep{{PlannedVUs: 5}, {TimeOffset: time.Second}}
}

func (c *coolexecutorconfig) GetDescription(*lib.ExecutionTuple) string {
	return "a cool description"
}

func (c *coolexecutorconfig) NewExecutor(*lib.ExecutionState, *logrus.Entry) (lib.Executor, error) {
	return &coolexecutor{c: c}, nil
}

func (c *coolexecutorconfig) HasWork(*lib.ExecutionTuple) bool {
	return true
}

type coolexecutor struct {
	c *coolexecutorconfig
	executor.BaseExecutor
}

func (c *coolexecutor) GetConfig() lib.ExecutorConfig {
	return c.c // panics otherwise
}

func (c *coolexecutor) GetProgress() *pb.ProgressBar {
	return pb.New(pb.WithConstLeft("cool")) // panics otherwise
}

func (c *coolexecutor) Run(ctx context.Context, engineOut chan<- stats.SampleContainer, builtinMetrics *metrics.BuiltinMetrics) error {
	return nil
}
