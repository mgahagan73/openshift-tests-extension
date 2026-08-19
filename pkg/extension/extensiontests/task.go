package extensiontests

import "sync/atomic"

type SpecTask struct {
	fn func(spec ExtensionTestSpec)
}

func (t *SpecTask) Run(spec ExtensionTestSpec) {
	t.fn(spec)
}

type TestResultTask struct {
	fn func(result *ExtensionTestResult)
}

func (t *TestResultTask) Run(result *ExtensionTestResult) {
	t.fn(result)
}

// SpecMutatingTask wraps a function that receives a pointer to the spec,
// allowing mutation. Used by BeforeSpawn hooks to inject per-dispatch data
// (e.g. Env) before the child process is created.
type SpecMutatingTask struct {
	fn func(spec *ExtensionTestSpec)
}

func (t *SpecMutatingTask) Run(spec *ExtensionTestSpec) {
	t.fn(spec)
}

type OneTimeTask struct {
	fn       func()
	executed int32 // Atomic boolean to indicate whether the function has been run
}

func (t *OneTimeTask) Run() {
	// Ensure one-time tasks are only run once
	if atomic.CompareAndSwapInt32(&t.executed, 0, 1) {
		t.fn()
	}
}
