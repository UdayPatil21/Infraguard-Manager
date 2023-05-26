package model

import (
	"time"

	"github.com/google/uuid"
)

type InstanceInfo struct {
	Agent_id  uuid.UUID `json:"agent_id" db:"agent_id"`
	Name      string    `json:"name" db:"name"`
	UserName  string    `json:"user_name" db:"user_name"`
	MachineID string    `json:"machine_id" db:"mahine_id"`
	PublicIP  string    `json:"public_ip" db:"public_ip"`
	HostName  string    `json:"host_name" db:"hostname"`
	OS        string    `json:"os_name" db:"os"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	Status    string    `json:"status" db:"status"`
}

type RunCommand struct {
	MachineID string `json:"machine_id" db:"machine_id" validate:"required"`
	Command   string `json:"command" validate:"required"`
}

type Executable struct {
	Script     []byte
	Permission string
}
