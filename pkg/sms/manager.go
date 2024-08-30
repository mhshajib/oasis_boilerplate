package sms

import "fmt"

type Manager struct {
	providers map[string]SMS
}

func NewManager() *Manager {
	return &Manager{
		providers: make(map[string]SMS),
	}
}

func (m *Manager) RegisterProvider(name string, provider SMS) {
	m.providers[name] = provider
}

func (m *Manager) Send(providerName, to, message string) (string, error) {
	provider, ok := m.providers[providerName]
	if !ok {
		return "", fmt.Errorf("Provider %s not found", providerName)
	}
	return provider.Send(to, message)
}
