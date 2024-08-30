package storage

import (
	"context"
	"fmt"
	"time"
)

type Manager struct {
	providers map[string]STORAGE
}

func NewManager() *Manager {
	return &Manager{
		providers: make(map[string]STORAGE),
	}
}

func (m *Manager) RegisterProvider(name string, provider STORAGE) {
	m.providers[name] = provider
}

func (m *Manager) GeneratePresignedUploadURL(ctx context.Context, providerName, bucket, key string, expireMinutes time.Duration) (string, string, error) {
	provider, ok := m.providers[providerName]
	if !ok {
		return "", "", fmt.Errorf("Provider %s not found", providerName)
	}
	return provider.GeneratePresignedUploadURL(ctx, bucket, key, expireMinutes)
}

func (m *Manager) CheckFileExists(ctx context.Context, providerName, bucket, key string) bool {
	provider, ok := m.providers[providerName]
	if !ok {
		return false
	}
	return provider.CheckFileExists(ctx, bucket, key)
}
