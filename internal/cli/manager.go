package cli

import (
	"os"
	"os/exec"
)

// ValeManager is a struct that manages a local installation of Vale.
type ValeManager struct {
	exe  string
	args []string
}

// NewValeManager creates a new ValeManager.
func NewValeManager(path string) (*ValeManager, error) {
	return &ValeManager{
		exe:  path,
		args: []string{"--output=JSON", "--no-exit"},
	}, nil
}

// Run lints the given text using Vale using the default configuration.
func (v *ValeManager) Lint(text, url, ext string) (string, error) {
	name := "*" + url + ext

	tmp, err := os.CreateTemp("", name)
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())

	_, err = tmp.WriteString(text)
	if err != nil {
		return "", err
	}

	return v.run([]string{tmp.Name()})
}

// Version returns the version of Vale.
func (v *ValeManager) Version() (string, error) {
	return v.run([]string{"--version"})
}

// Config returns the current configuration.
func (v *ValeManager) Config() (string, error) {
	return v.run([]string{"ls-config"})
}

func (v *ValeManager) run(args []string) (string, error) {
	cmd := exec.Command(v.exe, append(v.args, args...)...)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return string(result), err
	}

	return string(result), nil
}
