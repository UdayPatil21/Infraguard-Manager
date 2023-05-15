package model

type InstanceInfo struct {
	Name      any    `json:"name"`
	UserName  string `json:"user_name"`
	MachineID string `json:"machine_id"`
	PublicIP  string `json:"public_ip"`
	HostName  string `json:"host_name"`
	OS        string `json:"os_name"`
}
