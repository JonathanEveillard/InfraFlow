package config

import (
	"os"
	"time"
)

type AppConfig struct {
	Server struct {
		ListenAddr string
	}

	Database struct {
		DatabaseUrl string
	}

	Proxmox struct {
		ProxmoxUrl      string
		ProxmoxToken    string
		ProxmoxSecret   string
		ProxmoxInsecure bool
	}

	Tofu struct {
		TofuPath string
	}

	SSH struct {
		SSHKeyPath string
	}

	INFO struct {
		PollInterval      time.Duration
		MaxConcurrentJobs int
		LogLevel          string
	}
}

func Load() AppConfig {
	newCLI := os.Getenv()
}
