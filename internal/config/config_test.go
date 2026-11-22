package config

import "testing"

func TestLoadValidPort(t *testing.T) {
	t.Setenv("APP_PORT", "8080")

	cfg, err := Load()

	if err != nil {
		t.Fatalf("expected no err, got: %v", err)
	}

	if cfg.AppPort != "8080" {
		t.Errorf("expected port 8080, got %s", cfg.AppPort)
	}
}

func TestLoadInvalidPort(t *testing.T) {
	tests := []struct {
		name string
		port string
	}{
		{name: "non numeric port", port: "abc"},
		{name: "port above maximum", port: "70000"},
		{name: "zero port", port: "0"},
		{name: "negative port", port: "-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("APP_PORT", tt.port)

			cfg, err := Load()

			if err == nil {
				t.Fatalf("expected and error, got nil")
			}

			if cfg != nil {
				t.Fatalf("expected nil config, got %+v", cfg)
			}
		})
	}
}

func TestLoadDefaultPort(t *testing.T) {
	t.Setenv("APP_PORT", "")

	cfg, err := Load()

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if cfg.AppPort != "8080" {
		t.Errorf("expected default port 8080, got %s", cfg.AppPort)
	}
}
