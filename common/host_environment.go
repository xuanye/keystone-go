package common

import "strings"

type HostEnvironment struct {
	EnvironmentName string
}

func (h *HostEnvironment) IsDevelopment() bool {
	return strings.EqualFold(h.EnvironmentName, "development")
}

func (h *HostEnvironment) IsProduction() bool {
	return strings.EqualFold(h.EnvironmentName, "production")
}
