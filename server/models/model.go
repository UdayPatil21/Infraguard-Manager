package model

import "time"

type InstanceInfo struct {
	Name      any       `json:"name" db:"name"`
	UserName  string    `json:"user_name" db:"user_name"`
	MachineID string    `json:"machine_id" db:"mahine_id"`
	PublicIP  string    `json:"public_ip" db:"public_ip"`
	HostName  string    `json:"host_name" db:"hostname"`
	OS        string    `json:"os_name" db:"os"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type RunCommand struct {
	AgentID string `json:"agent_id" db:"agent_id" validate:"required"`
	Command string `json:"command" validate:"required"`
}

type Executable struct {
	Script     []byte
	Permission string
}
