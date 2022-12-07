package model

type UserStatus struct {
	CurrentAccount string `json:"current_account"`
	SimulationPid  int    `json:"simulation_pid"`
	SimulationPort int    `json:"simulation_port"`
	Chain          string `json:"chain"`
	OnlineMode     bool   `json:"online_mode"`
}
