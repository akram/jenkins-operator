package reason

import "fmt"

const (
	// OperatorSource defines that notification concerns operator
	OperatorSource Source = "operator"
	// KubernetesSource defines that notification concerns kubernetes
	KubernetesSource Source = "kubernetes"
	// HumanSource defines that notification concerns human
	HumanSource Source = "human"
)

// Reason is interface that let us know why operator sent notification.
type Reason interface {
	Short() []string
	Verbose() []string
	HasMessages() bool
}

// Undefined is base or untraceable reason.
type Undefined struct {
	source  Source
	short   []string
	verbose []string
}

// PodRestart defines the reason why Jenkins master pod restarted.
type PodRestart struct {
	Undefined
}

// DeploymentEvent informs that pod is being created.
type DeploymentEvent struct {
	Undefined
}

// ReconcileLoopFailed defines the reason why the reconcile loop failed.
type ReconcileLoopFailed struct {
	Undefined
}

// GroovyScriptExecutionFailed defines the reason why the groovy script execution failed.
type GroovyScriptExecutionFailed struct {
	Undefined
}

// BaseConfigurationFailed defines the reason why base configuration phase failed.
type BaseConfigurationFailed struct {
	Undefined
}

// BaseConfigurationComplete informs that base configuration is valid and complete.
type BaseConfigurationComplete struct {
	Undefined
}

// UserConfigurationFailed defines the reason why user configuration phase failed.
type UserConfigurationFailed struct {
	Undefined
}

// UserConfigurationComplete informs that user configuration is valid and complete.
type UserConfigurationComplete struct {
	Undefined
}

// NewUndefined returns new instance of Undefined.
func NewUndefined(source Source, short []string, verbose ...string) *Undefined {
	return &Undefined{source: source, short: short, verbose: checkIfVerboseEmpty(short, verbose)}
}

// NewPodRestart returns new instance of PodRestart.
func NewPodRestart(source Source, short []string, verbose ...string) *PodRestart {
	restartPodMessage := fmt.Sprintf("Jenkins master pod restarted by %s:", source)
	if len(short) == 1 {
		short = []string{fmt.Sprintf("%s %s", restartPodMessage, short[0])}
	} else if len(short) > 1 {
		short = append([]string{restartPodMessage}, short...)
	}

	if len(verbose) == 1 {
		verbose = []string{fmt.Sprintf("%s %s", restartPodMessage, verbose[0])}
	} else if len(verbose) > 1 {
		verbose = append([]string{restartPodMessage}, verbose...)
	}

	return &PodRestart{
		Undefined{
			source:  source,
			short:   short,
			verbose: checkIfVerboseEmpty(short, verbose),
		},
	}
}

// NewDeploymentEvent returns new instance of DeploymentEvent.
func NewDeploymentEvent(source Source, short []string, verbose ...string) *DeploymentEvent {
	return &DeploymentEvent{
		Undefined{
			source:  source,
			short:   short,
			verbose: checkIfVerboseEmpty(short, verbose),
		},
	}
}

// NewReconcileLoopFailed returns new instance of ReconcileLoopFailed.
func NewReconcileLoopFailed(source Source, short []string, verbose ...string) *ReconcileLoopFailed {
	return &ReconcileLoopFailed{
		Undefined{
			source:  source,
			short:   short,
			verbose: checkIfVerboseEmpty(short, verbose),
		},
	}
}

// NewGroovyScriptExecutionFailed returns new instance of GroovyScriptExecutionFailed.
func NewGroovyScriptExecutionFailed(source Source, short []string, verbose ...string) *GroovyScriptExecutionFailed {
	return &GroovyScriptExecutionFailed{
		Undefined{
			source:  source,
			short:   short,
			verbose: checkIfVerboseEmpty(short, verbose),
		},
	}
}

// NewBaseConfigurationFailed returns new instance of BaseConfigurationFailed.
func NewBaseConfigurationFailed(source Source, short []string, verbose ...string) *BaseConfigurationFailed {
	return &BaseConfigurationFailed{
		Undefined{
			source:  source,
			short:   short,
			verbose: checkIfVerboseEmpty(short, verbose),
		},
	}
}

// NewBaseConfigurationComplete returns new instance of BaseConfigurationComplete.
func NewBaseConfigurationComplete(source Source, short []string, verbose ...string) *BaseConfigurationComplete {
	return &BaseConfigurationComplete{
		Undefined{
			source:  source,
			short:   short,
			verbose: checkIfVerboseEmpty(short, verbose),
		},
	}
}

// NewUserConfigurationFailed returns new instance of UserConfigurationFailed.
func NewUserConfigurationFailed(source Source, short []string, verbose ...string) *UserConfigurationFailed {
	return &UserConfigurationFailed{
		Undefined{
			source:  source,
			short:   short,
			verbose: checkIfVerboseEmpty(short, verbose),
		},
	}
}

// NewUserConfigurationComplete returns new instance of UserConfigurationComplete.
func NewUserConfigurationComplete(source Source, short []string, verbose ...string) *UserConfigurationComplete {
	return &UserConfigurationComplete{
		Undefined{
			source:  source,
			short:   short,
			verbose: checkIfVerboseEmpty(short, verbose),
		},
	}
}

// Source is enum type that informs us what triggered notification.
type Source string

// Short is list of reasons.
func (p Undefined) Short() []string {
	return p.short
}

// Verbose is list of reasons with details.
func (p Undefined) Verbose() []string {
	return p.verbose
}

// HasMessages checks if there is any message.
func (p Undefined) HasMessages() bool {
	return len(p.short) > 0 || len(p.verbose) > 0
}

func checkIfVerboseEmpty(short []string, verbose []string) []string {
	if len(verbose) == 0 {
		return short
	}

	return verbose
}
