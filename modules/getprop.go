package modules

import (
	"fmt"
	"path/filepath"

	"github.com/botherder/androidqf/adb"
)

type GetProp struct {
	StoragePath string
}

func NewGetProp() *GetProp {
	return &GetProp{}
}

func (g *GetProp) Name() string {
	return "getprop"
}

func (g *GetProp) InitStorage(storagePath string) error {
	g.StoragePath = storagePath
	return nil
}

func (g *GetProp) Run() error {
	fmt.Println("Collecting device properties...")

	out, err := adb.Client.Shell("getprop")
	if err != nil {
		return fmt.Errorf("failed to run `adb shell getprop`: %v", err)
	}

	return saveCommandOutput(filepath.Join(g.StoragePath, "getprop.txt"), out)
}
