// Snoopdroid 2
// Copyright (c) 2021 Claudio Guarnieri.
// Use of this software is governed by the MVT License 1.1 that can be found at
//   https://license.mvt.re/1.1/

package acquisition

import (
	"fmt"
	"os"
	"path/filepath"
)

func (a *Acquisition) DumpSys() error {
	fmt.Println("Extracting device diagnostic information. This might take a while...")

	out, err := a.ADB.Shell("dumpsys")
	if err != nil {
		return fmt.Errorf("Unable to run `adb shell dumpsys`: %s", err)
	}

	fileName := "dumpsys.txt"
	file, err := os.Create(filepath.Join(a.BasePath, fileName))
	if err != nil {
		return fmt.Errorf("Unable to create %s file: %s", fileName, err)
	}
	defer file.Close()

	file.WriteString(out)

	return nil
}