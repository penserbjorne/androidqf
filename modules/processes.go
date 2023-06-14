package modules

import (
	"fmt"
	"path/filepath"

	"github.com/botherder/androidqf/adb"
)

type Processes struct {
	StoragePath string
}

func NewProcesses() *Processes {
	return &Processes{}
}

func (p *Processes) Name() string {
	return "processes"
}

func (p *Processes) InitStorage(storagePath string) error {
	p.StoragePath = storagePath
	return nil
}

func (p *Processes) Run() error {
	fmt.Println("Collecting list of running processes...")

	out, err := adb.Client.Shell("ps -A")
	if err != nil {
		return fmt.Errorf("failed to run `adb shell ps -A`: %v", err)
	}

	return saveCommandOutput(filepath.Join(p.StoragePath, "processes.txt"), out)
}
