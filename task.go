// The type definitions in this file were originally found in
// the eremetic project located: https://github.com/klarna/eremetic/blob/ae14694c1a44904cd2d533df7fc821aeb9b582ba/task.go
//
package main

// TaskState defines the valid task states.
type TaskState string

// Status represents the task status at a given time.
type Status struct {
	Time   int64     `json:"time"`
	Status TaskState `json:"status"`
}

// Volume is a mapping between ContainerPath and HostPath, to allow Docker
// to mount volumes.
type Volume struct {
	ContainerPath string `json:"container_path"`
	HostPath      string `json:"host_path"`
}

// Port defines a port mapping.
type Port struct {
	ContainerPort uint32 `json:"container_port"`
	HostPort      uint32 `json:"host_port"`
	Protocol      string `json:"protocol"`
}

// SlaveConstraint is a constraint that is validated for each slave when
// determining where to schedule a task.
type SlaveConstraint struct {
	AttributeName  string `json:"attribute_name"`
	AttributeValue string `json:"attribute_value"`
}

// URI holds meta-data for a sandbox resource.
type URI struct {
	URI        string `json:"uri"`
	Executable bool   `json:"executable"`
	Extract    bool   `json:"extract"`
	Cache      bool   `json:"cache"`
}

// Task defines the properties of a scheduled task.
type Task struct {
	TaskCPUs          float64           `json:"task_cpus"`
	TaskMem           float64           `json:"task_mem"`
	Command           string            `json:"command"`
	Args              []string          `json:"args"`
	User              string            `json:"user"`
	Environment       map[string]string `json:"env"`
	MaskedEnvironment map[string]string `json:"masked_env"`
	Image             string            `json:"image"`
	Volumes           []Volume          `json:"volumes"`
	Ports             []Port            `json:"ports"`
	Status            []Status          `json:"status"`
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	FrameworkID       string            `json:"framework_id"`
	SlaveID           string            `json:"slave_id"`
	SlaveConstraints  []SlaveConstraint `json:"slave_constraints"`
	Hostname          string            `json:"hostname"`
	Retry             int               `json:"retry"`
	CallbackURI       string            `json:"callback_uri"`
	SandboxPath       string            `json:"sandbox_path"`
	AgentIP           string            `json:"agent_ip"`
	AgentPort         int32             `json:"agent_port"`
	ForcePullImage    bool              `json:"force_pull_image"`
	FetchURIs         []URI             `json:"fetch"`
}
